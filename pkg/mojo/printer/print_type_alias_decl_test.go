package printer

import (
	"testing"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/pkg/context"
	"github.com/mojo-lang/mojo/pkg/mojo/parser/syntax"
)

func getTypeAliasDecl(file *lang.SourceFile) *lang.TypeAliasDecl {
	if len(file.Statements) > 0 {
		statement := file.Statements[0]
		if decl := statement.GetDeclaration(); decl != nil {
			return decl.GetTypeAliasDecl()
		}
	}
	return nil
}

func parseTypeAliasDecl(t *testing.T, decl string) *lang.TypeAliasDecl {
	parser := &syntax.Parser{}
	file, err := parser.ParseString(context.Empty(), decl)
	assert.NoError(t, err)

	alias := getTypeAliasDecl(file)
	assert.NotNil(t, alias)
	return alias
}

func TestPrinter_PrintTypeAliasDecl(t *testing.T) {
	const typeDecl = `
// comment1
// comment2

// comment3
type Mailbox=Box  //< following document - 1
//< following document - 2
`
	const expect = `// comment1
// comment2

// comment3
type Mailbox = Box //< following document - 1
                   //< following document - 2
`

	decl := parseTypeAliasDecl(t, typeDecl)
	p := New(&Config{}).PrintTypeAliasDecl(context.Empty(), decl)
	assert.Equal(t, expect, p.Buffer.String())
}
