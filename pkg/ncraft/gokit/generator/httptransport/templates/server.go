package templates

// ServerDecodeTemplate is the templates for generating the server-side decoding
// function for a particular Binding.
var ServerDecodeTemplate = `
{{- with $binding := . -}}
	// DecodeHTTP{{$binding.Label}}Request is a transport/http.DecodeRequestFunc that
	// decodes a JSON-encoded {{ToLower $binding.Parent.Name}} request from the HTTP request
	// body. Primarily useful in a server.
	func DecodeHTTP{{$binding.Label}}Request(_ context.Context, r *http.Request) (interface{}, error) {
		var req {{GoPackageName $binding.Parent.Request.Name}}.{{GoName $binding.Parent.Request.Name}}

		// to support gzip input
		var reader io.ReadCloser
		var err error 
		switch r.Header.Get("Content-Encoding") {
		case "gzip":
			reader, err = gzip.NewReader(r.Body)
			defer reader.Close()
			if err != nil {
				return nil, nhttp.WrapError(err, 400, "failed to read the gzip content")
			}
		default:
			reader = r.Body
		}

		buf, err := io.ReadAll(reader)
		if err != nil {
			return nil, nhttp.WrapError(err, 400, "cannot read body of http request")
		}
		if len(buf) > 0 {
			{{- if and $binding.Body $binding.Body.IsRawStyle}}
			req.{{GoName $binding.Body.Name}} = string(buf)
			{{else}}
			{{- if not $binding.Body}}
			if err = jsoniter.ConfigFastest.Unmarshal(buf, &req); err != nil {
			{{else}}
			{{- if not $binding.Body.IsScalar}}
			req.{{GoName $binding.Body.Name}} = {{if $binding.Body.IsMap}}make({{$binding.Body.GetGoTypeName}}){{else if $binding.Body.IsArray}}{{$binding.Body.GetGoTypeName}}{}{{else}}&{{$binding.Body.GetGoTypeName}}{}{{end}}
			{{end -}}
			if err = jsoniter.ConfigFastest.Unmarshal(buf, {{if or (or $binding.Body.IsMap $binding.Body.IsArray) $binding.Body.IsScalar}}&{{end}}req.{{GoName $binding.Body.Field.Name}}); err != nil {
			{{end -}}
				const size = 8196
				if len(buf) > size {
					buf = buf[:size]
				}
				return nil, nhttp.WrapError(err,
					http.StatusBadRequest,
					fmt.Sprintf("request body '%s': cannot parse non-json request body", buf),
				)
			}
			{{end -}}
		}

		pathParams := mux.Vars(r)
		_ = pathParams

		queryParams := core.NewUrlQueryFrom(r.URL.Query())
		_ = queryParams

		parsedQueryParams := make(map[string]bool)
		_ = parsedQueryParams

		{{range $param := $binding.Parameters}}
			{{if ne $param.Location "body"}}
				{{$param.Go.ParamUnmarshaler}}
			{{end}}
		{{end}}
		return &req, nil
	}
{{- end -}}
`

// WARNING: Changing the contents of these strings, even a little bit, will cause tests
// to fail. So don't change them purely because you think they look a little
// funny.

var ServerTemplate = `// Code generated by chaosmojo. DO NOT EDIT.
// Rerunning chaosmojo will overwrite this file.
// Version: {{.Version}}
// Version Date: {{.VersionDate}}

package svc

// This file provides server-side bindings for the HTTP transport.
// It utilizes the transport/http.Server.

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/chaos-io/chaos/logs"
	pagination "github.com/chaos-io/gokit/pagination"
	"github.com/chaos-io/gokit/tracing"
	nhttp "github.com/chaos-io/gokit/transport/http"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/json-iterator/go"
	mjhttp "github.com/mojo-lang/http/go/pkg/mojo/http"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/trace"

    {{$corePackage := "github.com/mojo-lang/core/go/pkg/mojo/core"}}
    "{{$corePackage}}"
    {{range $i := .Go.ImportedTypePaths}}
	{{if ne $i $corePackage}}"{{$i}}"{{end}}
	{{- end}}

	// this service api
	pb "{{.Go.ApiImportPath -}}"
)

const contentType = "application/json; charset=utf-8"

var (
	_ = fmt.Sprint
	_ = bytes.Compare
	_ = strconv.Atoi
	_ = httptransport.NewServer
	_ = ioutil.NopCloser
	_ = pb.New{{.Interface.BaredName}}Client
	_ = io.Copy
	_ = errors.Wrap
    _ = mjhttp.UnmarshalQueryParam
)
{{if .HasImported}}
var ({{range $msg := .ImportedMessages}}
	_ = {{$msg.Go.PackageName}}.{{$msg.Name}}{}
{{- end}}{{range $enum := .ImportedEnums}}
	_ = {{$enum.Go.PackageName}}.{{$enum.Name}}(0)
{{- end}}){{end}}

var cfg *nhttp.Config

func init()  {
	cfg = nhttp.NewConfig()
}

// RegisterHttpHandler register a set of endpoints available on predefined paths to the router.
func RegisterHttpHandler(router *mux.Router, endpoints Endpoints, tracer trace.Tracer, logger log.Logger)  {
	{{- if .Interface.Methods}}
		serverOptions := []httptransport.ServerOption{
			httptransport.ServerBefore(headersToContext, queryToContext),
			httptransport.ServerErrorEncoder(errorEncoder),
			httptransport.ServerErrorLogger(logger),
			httptransport.ServerAfter(httptransport.SetContentType(contentType)),
		}
	{{- end }}

	addTracerOption := func(methodName string) []httptransport.ServerOption {
	    if tracer != nil {
	        return append(serverOptions, httptransport.ServerBefore(tracing.HTTPToContext(tracer, methodName)))
	    }
	    return serverOptions
	}

	{{range $method := .Interface.Methods}}
		{{range $binding := $method.Bindings}}
			router.Methods("{{$binding.Verb | ToUpper}}").Path("{{$binding.Path}}").Handler(
				httptransport.NewServer(
					endpoints.{{ToCamel $method.Name}}Endpoint,
					DecodeHTTP{{$binding.Label}}Request,{{if $binding.GetResponseBody}}
					EncodeHTTP{{ToCamel $method.Name}}Response,{{else}}
					EncodeHTTPGenericResponse,{{end}}
					addTracerOption("{{$method.Name}}")...,
			))
		{{- end}}
	{{- end}}
}

// ErrorEncoder writes the error to the ResponseWriter, by default a content
// type of application/json, a body of json with key "error" and the value
// error.Error(), and a status code of 500. If the error implements Headerer,
// the provided headers will be applied to the response. If the error
// implements json.Marshaler, and the marshaling succeeds, the JSON encoded
// form of the error will be used. If the error implements StatusCoder, the
// provided StatusCode will be used instead of 500.
func errorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	e, ok := err.(*core.Error)
	if !ok {
		e = core.NewErrorFrom(500, err.Error())
	}

	w.Header().Set("Content-Type", contentType)
	if headerer, ok := err.(httptransport.Headerer); ok {
		for k := range headerer.Headers() {
			w.Header().Set(k, headerer.Headers().Get(k))
		}
	}

	var body []byte
	code := http.StatusInternalServerError

	if enveloped := nhttp.IsEnvelopeStyle(ctx, cfg.GetStyle()); enveloped {
		envelope := &nhttp.EnvelopedResponse{}
		envelope.Error = e

		if cfg.GetEnvelop().MappingCode {
			if sc, ok := err.(httptransport.StatusCoder); ok {
				code = sc.StatusCode()
			}
		} else {
			code = http.StatusOK
		}

		var response interface{}
		if cfg.GetEnvelop().ErrorWrapped {
			response = envelope.ToErrorWrapped()
		} else {
			response = envelope
		}

		jsonBody, marshalErr := jsoniter.ConfigFastest.Marshal(response)
		if marshalErr != nil {
			logs.Warnw("failed to marshal the error response to json", "error", marshalErr)
		} else {
			body = jsonBody
		}
	} else {
		if sc, ok := err.(httptransport.StatusCoder); ok {
			code = sc.StatusCode()
		}
		if marshaler, ok := err.(json.Marshaler); ok {
			if jsonBody, marshalErr := marshaler.MarshalJSON(); marshalErr == nil {
				body = jsonBody
			}
		}

		if jsonBody, marshalErr := jsoniter.ConfigFastest.Marshal(e); marshalErr == nil {
			body = jsonBody
		}
	}

	w.WriteHeader(code)
	w.Write(body)
}

// Server Decode
{{range $method := .Interface.Methods}}
	{{range $binding := $method.Bindings}}
		{{$binding.Extensions.ServerDecode}}
	{{end}}
{{end}}

{{range $method := .Interface.Methods}}
    {{range $binding := $method.Bindings}}
	    {{if $binding.GetResponseBody}}
func EncodeHTTP{{ToCamel $method.Name}}Response(_ context.Context, w http.ResponseWriter, response interface{}) error {
	r := response.(*{{GoPackageName $method.Response.Body.Name}}.{{$method.Response.Body.Name}})
    
	{{range $h, $v := $method.Response.Headers}}
	w.Header().Set("{{$h}}", "{{$v}}")
	{{end}}

	cnt, err := w.Write([]byte(r.{{$method.Response.Body.Name}}))
	if err != nil {
		return err
	}
	responseCnt := len(r.{{$method.Response.Body.Name}})
	if cnt != responseCnt {
		return errors.Errorf("wrong to write response content expect %d, but %d written", responseCnt, cnt)
	}

	return nil
}
        {{end}}
    {{end}}
{{end}}

// EncodeHTTPGenericResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func EncodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if writer, ok := response.(nhttp.ResponseWriter); ok {
		return writer.WriteHttpResponse(ctx, w)
	}

	if reflect.ValueOf(response).IsNil() {
		response = nil
	}
	if _, ok := response.(*core.Null); ok {
		response = nil
	}

	enveloped := nhttp.IsEnvelopeStyle(ctx, cfg.GetStyle())
	if enveloped {
		code := core.NewErrorCode(200)
		message := "OK"
		if sc := cfg.GetEnvelop().SuccessCode; len(sc) > 0 {
			if c, err := core.ParseErrorCode(sc); err != nil {
				logs.Warnw("failed to parse the user setting success code, will use \"200\" indeed.", "code", sc, "error", err)
			} else {
				code = c
			}
		}
		if msg := cfg.GetEnvelop().SuccessMessage; len(msg) > 0 {
			message = msg
		}

		totalCount := int32(0)
		nextPageToken := ""
		if p, ok := response.(pagination.Paginater); ok {
			totalCount = p.GetTotalCount()
			nextPageToken = p.GetNextPageToken()
		}

		response = &nhttp.EnvelopedResponse {
			Error: &core.Error{
				Code:    code,
				Message: message,
			},
			TotalCount: totalCount,
			NextPageToken: nextPageToken,
			Data: response,
		}
	}

	if response == nil {
		return nil
	}

	if p, ok := response.(pagination.Paginater); ok && !enveloped {
		total := p.GetTotalCount()
		if total > 0 {
			w.Header().Set("X-Total-Count", strconv.Itoa(int(total)))
		}

		next := p.GetNextPageToken()
		if len(next) > 0 {
			path, _ := ctx.Value("http-request-path").(string)
			if len(path) == 0 {
				path = "/?next-page-token=" + next
			} else {
				url, _ := core.ParseUrl(path)
				url.Query.Add("next-page-token", next)
				path = url.Format()
			}
			w.Header().Set("Link", fmt.Sprintf("<%s>; rel=\"next\"", path))
		}
	}

	return nhttp.NewResponseJsonWriter(response).WriteHttpResponse(ctx, w)
}

// Helper functions

func headersToContext(ctx context.Context, r *http.Request) context.Context {
	for k, _ := range r.Header {
		// The key is added both in http format (k) which has had
		// http.CanonicalHeaderKey called on it in transport as well as the
		// strings.ToLower which is the grpc metadata format of the key so
		// that it can be accessed in either format
		ctx = context.WithValue(ctx, k, r.Header.Get(k))
		ctx = context.WithValue(ctx, strings.ToLower(k), r.Header.Get(k))
	}

	// add the access key to context
	accessKey := r.URL.Query().Get("access_key")
	if len(accessKey) > 0{
		ctx = context.WithValue(ctx, "access_key", accessKey)
	}

	// Tune specific change.
	// also add the request url
	ctx = context.WithValue(ctx, "http-request-path", r.URL.Path)
	ctx = context.WithValue(ctx, "transport", "HTTPJSON")

	return ctx
}

func queryToContext(ctx context.Context, r *http.Request) context.Context {
	check := func(values []string) bool{
		for _, value := range values {
			if value == "true" {
				return true
			}
		}
		return false
	}
	for key, values := range r.URL.Query() {
		switch key {
		case "envelope":
			if check(values) {
				ctx = context.WithValue(ctx, "envelope", true)
			}
		}
	}
	return ctx
}
`
