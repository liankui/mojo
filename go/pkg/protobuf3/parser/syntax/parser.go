package syntax

import (
	"context"
	"io/fs"
	"path"
	"regexp"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/mojo-lang/core/go/pkg/logs"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"github.com/mojo-lang/lang/go/pkg/mojo/lang"

	"github.com/mojo-lang/mojo/go/pkg/util"
)

type Parser struct {
}

func New(options core.Options) *Parser {
	return &Parser{}
}

func (p Parser) ParseString(mojo string) (*lang.SourceFile, error) {
	input := antlr.NewInputStream(mojo)
	return p.ParseStream("", input)
}

var protoMismatched = regexp.MustCompile(`.*mismatched.+proto2.+proto3.+`)

func (p Parser) ParseStream(fileName string, input antlr.CharStream) (*lang.SourceFile, error) {
	lexer := NewProtobuf3Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	errorListener := util.NewErrorListener(fileName, false)

	parser := NewProtobuf3Parser(stream)
	parser.AddErrorListener(errorListener)
	parser.BuildParseTrees = true

	tree := parser.Proto()
	if sourceFile, ok := tree.Accept(NewProtoVisitor()).(*lang.SourceFile); ok {
		if len(errorListener.Errors) == 0 {
			return sourceFile, nil
		}
	}

	if len(errorListener.Errors) > 0 {
		msg := errorListener.Errors[0].Error()
		if protoMismatched.MatchString(msg) {
			return nil, &ProtoError{}
		}

		return nil, logs.NewErrorw("failed to parse proto3 file", "file", fileName, "error", errorListener.Errors.Error())
	}

	return nil, logs.NewErrorw("failed to parse proto3 file", "file", fileName)
}

func (p Parser) ParseFile(ctx context.Context, fileName string, fileSys fs.FS) (*lang.SourceFile, error) {
	if bytes, err := fs.ReadFile(fileSys, fileName); err != nil {
		return nil, err
	} else {
		if sourceFile, err := p.ParseStream(fileName, antlr.NewInputStream(string(bytes))); err != nil {
			return nil, err
		} else {
			sourceFile.Name = path.Base(fileName)
			sourceFile.FullName = fileName

			return sourceFile, nil
		}
	}
}
