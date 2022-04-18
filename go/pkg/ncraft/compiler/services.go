package compiler

import (
    "fmt"
    "sort"
    "strings"
    "text/template"

    "github.com/mojo-lang/core/go/pkg/logs"
    "github.com/mojo-lang/core/go/pkg/mojo/core"
    "github.com/mojo-lang/core/go/pkg/mojo/core/strcase"
    "github.com/mojo-lang/http/go/pkg/mojo/http"
    "github.com/mojo-lang/lang/go/pkg/mojo/lang"
    "github.com/mojo-lang/mojo/go/pkg/mojo/compiler/transformer"
    "github.com/mojo-lang/mojo/go/pkg/mojo/context"
    "github.com/mojo-lang/mojo/go/pkg/mojo/plugin"
    "github.com/mojo-lang/mojo/go/pkg/ncraft/data"
    apicompiler "github.com/mojo-lang/mojo/go/pkg/openapi/compiler"
    "github.com/mojo-lang/mojo/go/pkg/protobuf/precompiler"
    "github.com/mojo-lang/openapi/go/pkg/mojo/openapi"
    "github.com/mojo-lang/protobuf/go/pkg/mojo/protobuf"
)

type Services struct {
    Data []*data.Service

    Interfaces []*data.Interface
}

func CompilePackage(ctx context.Context, pkg *lang.Package) ([]*data.Service, error) {
    services := &Services{}
    if err := plugin.CompilePackage(services, ctx, pkg); err != nil {
        return nil, err
    }

    for _, s := range services.Data {
        s.AllInterfaces = services.Interfaces
    }

    return services.Data, nil
}

func (s *Services) CompileInterface(ctx context.Context, decl *lang.InterfaceDecl) error {
    if len(decl.GenericParameters) > 0 {
        return nil
    }

    thisCtx := context.WithType(ctx, decl)
    pkg := context.Package(ctx)

    pkgName := lang.GetGoPackageName(decl.PackageName)
    service := &data.Service{
        PackageName:     pkg.Name,
        PackageFullName: pkg.FullName,
        Interface: &data.Interface{
            Decl:       decl,
            Name:       decl.Name,
            BaredName:  data.GetInterfaceBaredName(decl.Name),
            ServerName: data.GetInterfaceServerName(decl.Name),
            Methods:    nil,
        },
        FuncMap: template.FuncMap{
            "ToLower":                    strings.ToLower,
            "ToUpper":                    strings.ToUpper,
            "GoName":                     GoName,
            "ToSnake":                    strcase.ToSnake,
            "ToKebab":                    strcase.ToKebab,
            "ToCamel":                    strcase.ToCamel,
            "ToLowerCamel":               strcase.ToLowerCamel,
            "Plural":                     transformer.Plural,
            "GoPackageName":              GoPackageName,
            "GoFieldArrayElementType":    GoFieldArrayElementType,
            "GoIsArrayElementStringType": GoIsArrayElementStringType,
        },
    }
    service.Go.PackageName = pkgName

    if path, ok := GoPackageImport(ctx, decl.PackageName).(string); ok {
        service.Go.ApiImportPath = path
    }

    for _, d := range decl.Type.Methods {
        if err := s.CompileMethod(thisCtx, d, service); err != nil {
            return err
        }
    }

    service.Go.ImportedTypePaths = unifyStringArray(service.Go.ImportedTypePaths)

    s.Data = append(s.Data, service)
    s.Interfaces = append(s.Interfaces, service.Interface)

    return nil
}

func unifyStringArray(array []string) []string {
    index := make(map[string]bool)
    for _, item := range array {
        if _, ok := index[item]; !ok {
            index[item] = true
        }
    }
    var result []string
    for key, _ := range index {
        result = append(result, key)
    }
    sort.Strings(result)
    return result
}

func (s *Services) CompileMethod(ctx context.Context, decl *lang.FunctionDecl, service *data.Service) error {
    thisCtx := context.WithType(ctx, decl)
    m := &data.Method{
        Decl:     decl,
        Name:     decl.Name,
        Bindings: nil,
    }
    registerType := func(t *lang.NominalType) {
        if !t.IsScalar() && !t.IsMapType() && !t.IsArrayType() && !t.IsUnionType() && (len(t.PackageName) > 0 && t.PackageName != service.PackageFullName) {
            goPackageName := lang.GetGoPackageName(t.PackageName)
            if RegisterMessageGoPackageName(t.Name, goPackageName) {
                return
            }

            if len(t.GenericArguments) == 0 {
                if t.TypeDeclaration != nil && t.TypeDeclaration.GetEnumDecl() != nil {
                    service.ImportedEnums = append(service.ImportedEnums, &data.Enum{
                        Decl:        t.GetTypeDeclaration().GetEnumDecl(),
                        Name:        t.Name,
                        PackageName: t.PackageName,
                        Go: data.GoEnum{
                            PackageName: goPackageName,
                        },
                    })
                } else {
                    service.ImportedMessages = append(service.ImportedMessages, &data.Message{
                        Decl:        t.GetTypeDeclaration().GetStructDecl(),
                        Name:        t.Name,
                        PackageName: t.PackageName,
                        Go: data.GoMessage{
                            PackageName: goPackageName,
                        },
                    })
                }

                if path, ok := GoPackageImport(ctx, t.PackageName).(string); ok {
                    service.Go.ImportedTypePaths = append(service.Go.ImportedTypePaths, path)
                }
            }
        }
    }

    req, resp, err := precompiler.CompileMethod(ctx, decl)
    if err != nil {
        return err
    }
    if m.Request, err = s.CompileMessage(thisCtx, req); err != nil {
        return err
    }
    for _, field := range req.Type.Fields {
        t := field.Type
        if t.IsMapType() {
            for _, gt := range t.GenericArguments {
                registerType(gt)
            }
        } else if !t.IsScalar() {
            registerType(t)
        }
    }

    result := decl.GetSignature().GetResultType()
    if result.GetTypeDeclaration().GetStructDecl() != resp {
        result = &lang.NominalType{
            PackageName:     resp.PackageName,
            Name:            resp.Name,
            TypeDeclaration: lang.NewStructTypeDeclaration(resp),
        }
    }
    if m.Response, err = s.CompileMessage(thisCtx, resp); err != nil {
        return err
    }

    var httpResponse *data.HTTPResponse
    if result != nil {
        if v, e := result.GetStringAttribute(http.BodyAttributeFullName); e == nil && len(v) > 0 {
            httpResponse = &data.HTTPResponse{
                Headers: make(map[string]string),
            }

            for _, attribute := range result.Attributes {
                if attribute.IsSameName(http.HeaderAttributeFullName) {
                    if len(attribute.Arguments) == 2 {
                        key, _ := attribute.Arguments[0].GetString()
                        value, _ := attribute.Arguments[1].GetString()
                        if len(key) > 0 && len(value) > 0 {
                            httpResponse.Headers[key] = value
                        }
                    }
                }
            }

            for _, field := range m.Response.Fields {
                if strings.ToLower(v) == strings.ToLower(field.Name) {
                    httpResponse.Body = field
                }
            }
        }

        registerType(result)
    }

    //if len(decl.PackageName) > 0 && decl.PackageName != service.PackageFullName {
    //    goPackageName := lang.GetGoPackageName(decl.PackageName)
    //    if !RegisterMessageGoPackageName(decl.Name, goPackageName) && result != nil && len(result.GenericArguments) == 0 {
    //        if result.TypeDeclaration != nil && result.TypeDeclaration.GetEnumDecl() != nil {
    //            service.ImportedEnums = append(service.ImportedEnums, &data.Enum{
    //                Decl:        t.GetTypeDeclaration().GetEnumDecl(),
    //                Name:        t.Name,
    //                PackageName: t.PackageName,
    //                Go: data.GoEnum{
    //                    PackageName: goPackageName,
    //                },
    //            })
    //
    //            service.ImportedEnums = append(service.ImportedEnums, &data.Enum{
    //                Name: decl.Name,
    //            })
    //        } else {
    //            service.ImportedMessages = append(service.ImportedMessages, &data.Message{Name: decl.Name})
    //        }
    //
    //        if path, ok := GoPackageImport(ctx, decl.PackageName).(string); ok {
    //            service.Go.ImportedTypePaths = append(service.Go.ImportedTypePaths, path)
    //        }
    //    }
    //}

    index := 0
    for _, attribute := range decl.Attributes {
        if attribute.PackageName != "http" {
            continue
        }

        switch attribute.Name {
        case http.GetAttributeName, http.PostAttributeName, http.PutAttributeName, http.PatchAttributeName,
            http.DeleteAttributeName, http.OptionsAttributeName, http.HeadAttributeName, http.TraceAttributeName:
            if len(attribute.Arguments) > 0 {
                if value, err := attribute.Arguments[0].GetString(); err == nil && len(value) > 0 {
                    thisCtx = WithDataMethod(context.WithValues(ctx, "index", index), m)
                    binding, err := s.compileBindings(thisCtx, attribute.Name, value, decl)
                    if err != nil {
                        continue
                    }

                    sort.Sort(binding.Parameters)
                    
                    m.Bindings = append(m.Bindings, binding)
                    index++
                }
            }
        }
    }

    service.Interface.Methods = append(service.Interface.Methods, m)
    return nil
}

func (s *Services) CompileMessage(ctx context.Context, decl *lang.StructDecl) (*data.Message, error) {
    msg := &data.Message{
        Decl:        decl,
        PackageName: decl.PackageName,
        Name:        decl.Name,
    }

    decl.EachField(func(f *lang.ValueDecl) error {
        t := f.GetType()
        fieldType := &data.FieldType{}

        switch t.GetFullName() {
        case core.ArrayTypeFullName:
            if len(t.GenericArguments) == 1 {
                fieldType.Name = t.GenericArguments[0].Name
                fieldType.IsArray = true
            }
        case core.MapTypeFullName:
            if len(t.GenericArguments) == 2 {
                fieldType.KeyType = &data.FieldType{
                    Name: t.GenericArguments[0].Name,
                }
                fieldType.Name = t.GenericArguments[1].Name
                fieldType.IsMap = true
            }
        default:
            fieldType.Name = t.Name
        }

        msg.Fields = append(msg.Fields, &data.Field{
            Name: f.Name,
            Type: fieldType,
        })

        return nil
    })

    return msg, nil
}

func (s *Services) compileBindings(ctx context.Context, methodName string, path string, method *lang.FunctionDecl) (*data.HTTPBinding, error) {
    dm := DataMethod(ctx)
    p, pathParams := apicompiler.CompilePath(path)

    index := 0
    if value, ok := ctx.Value("index").(int); ok {
        index = value
    }

    binding := &data.HTTPBinding{
        Verb:       methodName,
        Path:       p,
        Label:      strcase.ToCamel(method.Name) + data.EnglishNumber(index),
        Extensions: make(map[string]interface{}),
        Parent:     dm,
    }

    if len(method.Signature.GetParameters()) == 1 {
        decl := method.Signature.ParameterDecl(0)
        if b, _ := decl.GetBoolAttribute(protobuf.MethodRequestTypeAttributeName); b {
            if structDecl := decl.Type.GetTypeDeclaration().GetStructDecl(); structDecl != nil {
                structDecl.EachField(func(decl *lang.ValueDecl) error {
                    return s.compileBindingParameter(ctx, decl, pathParams, nil, binding)
                })
                return binding, nil
            } else {
                return nil, logs.NewErrorw("invalid request type", "method", method.FullName, "http-bind", methodName+":"+path)
            }
        }
    }

    for _, param := range method.Signature.GetParameters() {
        if err := s.compileBindingParameter(ctx, param, pathParams, nil, binding); err != nil {
            return nil, err
        }
    }

    if binding.Body == nil && http.CanCarryBody(binding.Verb) {
        for _, parameter := range binding.Parameters {
            goTypeName := parameter.GetGoTypeName()
            if parameter.Location == "query" && len(GoPackageName(goTypeName)) > 0 &&
                !(strings.HasPrefix(goTypeName, "map[") || strings.HasPrefix(goTypeName, "[]")) {
                binding.Body = parameter
                break
            }
        }
    }

    return binding, nil
}

func (s *Services) compileBindingParameter(ctx context.Context, decl *lang.ValueDecl, pathParams map[string]bool, enclosing *data.HTTPParameter, binding *data.HTTPBinding) error {
    compiler := &apicompiler.NominalTypeCompiler{}

    schema, err := compiler.Compile(decl.Type)
    if err != nil {
        return err
    }

    param := &data.HTTPParameter{
        Name: decl.Name,

        Field: s.compileBindingField(ctx, schema, context.Components(compiler.Context).GetSchemas()),
    }

    if param.Field == nil {
        return logs.NewErrorw("failed to compile the binding field")
    }

    param.Field.Name = decl.Name

    if enclosing != nil {
        param.Field.Enclosing = enclosing.Field
        param.Field.FullName = enclosing.Field.FullName + "." + decl.Name
    } else {
        param.Field.FullName = decl.Name
    }
    param.FullName = strcase.ToSnake(strings.ReplaceAll(param.Field.FullName, ".", "_"))

    if v, e := lang.GetStringAttribute(decl.Type.Attributes, core.ExplodedAttributeFullName); e == nil {
        param.Field.Exploded = true
        param.Field.ExplodePrefix = v
        if len(v) == 0 {
            param.Field.ExplodePrefix = decl.Name
        }

        structDecl := decl.Type.GetTypeDeclaration().GetStructDecl()
        if len(decl.Type.GenericArguments) == 0 && structDecl != nil {
            err = structDecl.EachField(func(decl *lang.ValueDecl) error {
                if err = s.compileBindingParameter(ctx, decl, pathParams, param, binding); err != nil {
                    return err
                }
                return nil
            })
            if err != nil {
                return err
            }
        }
    }

    param.Go.LocalName = fmt.Sprintf("%s%s", strcase.ToLowerCamel(decl.Name), strcase.ToCamel(binding.Parent.Name))

    //TODO update sync with openapi, add request body
    if pathParams[param.Field.Name] || pathParams[param.Field.FullName] {
        param.Location = "path"
    } else {
        param.Location = "query"
    }
    binding.Parameters = append(binding.Parameters, param)

    if fieldType := param.Field.GetType(); fieldType != nil && fieldType.Message != nil {
        for pathParam := range pathParams {
            if strings.Contains(pathParam, ".") {
                for _, field := range fieldType.Message.Fields {
                    fullName := param.Field.Name + "." + field.Name
                    if pathParam != fullName {
                        fullName = param.Field.Name + "." + strcase.ToSnake(field.Name)
                    }

                    if pathParam == fullName {
                        parameter := &data.HTTPParameter{
                            Name:     field.Name,
                            FullName: HttpPathParameterName(fullName),
                            Field: &data.Field{
                                Name:      field.Name,
                                FullName:  fullName,
                                Type:      field.Type,
                                Enclosing: param.Field,
                                Decl:      field.Decl,
                            },
                            Location: "path",
                        }
                        parameter.Go.LocalName = fmt.Sprintf("%s%s", strcase.ToLowerCamel(parameter.Name), strcase.ToCamel(binding.Parent.Name))
                        binding.Parameters = append(binding.Parameters, parameter)
                    }
                }
            }
        }
    }
    return nil
}

func (s *Services) compileBindingField(ctx context.Context, schema *openapi.Schema, index map[string]*openapi.Schema) *data.Field {
    if schema == nil {
        // error
        return nil
    }

    field := &data.Field{}
    switch schema.Type {
    case openapi.Schema_TYPE_BOOLEAN:
        field.Type = &data.FieldType{
            Name: "Bool",
            Go: data.GoFieldType{
                Name:       "bool",
                IsBaseType: true,
                IsPointer:  false,
            },
        }
    case openapi.Schema_TYPE_INTEGER:
        if len(schema.Format) > 0 {
            switch schema.Format {
            case "Int8", "Int16":
                field.Type = &data.FieldType{
                    Name: "Int32",
                }
            default:
                field.Type = &data.FieldType{
                    Name: schema.Format,
                }
            }
            field.Type.Go = data.GoFieldType{
                Name:       strings.ToLower(field.Type.Name),
                IsBaseType: true,
                IsPointer:  false,
            }
        } else {
            field.Type = &data.FieldType{
                Name: "Int64",
                Go: data.GoFieldType{
                    Name:       "int64",
                    IsBaseType: true,
                },
            }
        }
    case openapi.Schema_TYPE_NUMBER:
        field.Type = &data.FieldType{
            Name: "Float64",
            Go: data.GoFieldType{
                Name:       "float64",
                IsBaseType: true,
            },
        }
    case openapi.Schema_TYPE_STRING:
        if len(schema.Enum) > 0 {
            field.Type = &data.FieldType{
                PackageName: lang.GetTypePackageName(schema.Title),
                Name:        lang.GetTypeTypeName(schema.Title),
                IsEnum:      true,
                Go: data.GoFieldType{
                    Name: getGoTypeName(schema.Title),
                },
            }
        } else {
            field.Type = &data.FieldType{
                Name: "String",
                Go: data.GoFieldType{
                    Name:       "string",
                    IsBaseType: true,
                },
            }
        }
    case openapi.Schema_TYPE_ARRAY:
        sch := schema.Items.GetSchemaOf(index)
        f := s.compileBindingField(ctx, sch, index)
        if f != nil {
            field.Type = &data.FieldType{
                Name:    f.Type.Name,
                Enum:    f.Type.Enum,
                Message: f.Type.Message,
                IsArray: true,
            }

            field.Type.ElementGo = f.Type.Go
            if f.Type.Go.IsPointer {
                field.Type.Go.Name = "[]*" + f.Type.Go.Name
            } else {
                field.Type.Go.Name = "[]" + f.Type.Go.Name
            }
        }
    case openapi.Schema_TYPE_OBJECT:
        if schema.AdditionalProperties != nil { // map
            typ := s.compileBindingField(ctx, schema.GetAdditionalProperties().GetSchema(), index).GetType()
            field.Type = &data.FieldType{
                PackageName: typ.PackageName,
                Name:        typ.Name,
                Enum:        typ.Enum,
                Message:     typ.Message,
                KeyType: &data.FieldType{
                    Name: "String",
                    Go: data.GoFieldType{
                        Name:       "string",
                        IsBaseType: true,
                    },
                },
                IsMap: true,
                Go: data.GoFieldType{
                    IsBaseType: false,
                },
                ElementGo: typ.Go,
            }
            if field.Type.Go.IsPointer {
                field.Type.Go.Name = "map[string]*" + typ.Go.Name
            } else {
                field.Type.Go.Name = "map[string]" + typ.Go.Name
            }
        } else {
            field.Type = &data.FieldType{
                PackageName: lang.GetTypePackageName(schema.Title),
                Name:        lang.GetTypeTypeName(schema.Title),
                Message: &data.Message{
                    PackageName: lang.GetTypePackageName(schema.Title),
                    Name:        lang.GetTypeTypeName(schema.Title),
                },
                Go: data.GoFieldType{
                    Name:       getGoTypeName(schema.Title),
                    IsBaseType: false,
                    IsPointer:  true,
                },
            }

            for name, item := range schema.Properties {
                f := s.compileBindingField(ctx, item.GetSchemaOf(index), index)
                f.Name = strcase.ToSnake(name)
                field.Type.Message.Fields = append(field.Type.Message.Fields, f)
            }
        }
    }

    return field
}

func getGoTypeName(title string) string {
    pkg := lang.GetTypeGoPackageName(title)
    name := lang.GetTypeGoTypeName(title)
    if len(pkg) > 0 {
        return strings.Join([]string{pkg, name}, ".")
    }
    return name
}