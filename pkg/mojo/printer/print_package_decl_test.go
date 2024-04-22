package printer

import (
	"testing"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/pkg/context"
	"github.com/mojo-lang/mojo/pkg/mojo/parser/syntax"
)

func getPackageDecl(file *lang.SourceFile) *lang.PackageDecl {
	if len(file.Statements) > 0 {
		statement := file.Statements[0]
		if decl := statement.GetDeclaration(); decl != nil {
			return decl.GetPackageDecl()
		}
	}
	return nil
}

func parsePackageDecl(t *testing.T, str string) *lang.PackageDecl {
	parser := &syntax.Parser{}
	file, err := parser.ParseString(context.Empty(), str)
	assert.NoError(t, err)

	decl := getPackageDecl(file)
	assert.NotNil(t, decl)
	return decl
}

func TestPrinter_PrintPackageDecl(t *testing.T) {
	const typeDecl = `
///
package mojo.lang {
version: '0.1.0'
license: 'Apache'
authors: [{
        author: 'Frankee'
        email: 'frankee.zhou@gmail.com'
        organization: 'mojolang.org'
}]
dependencies: {
    'mojo.core': {path: '../core', version: '^0.1'}
    'mojo.document': {path: '../document', version: '^0.1'}
}

repository: 'https://github.com/mojo-lang/lang'
}
`
	const expect = `/// 
package mojo.lang {
    version: "0.1.0"
    license: "Apache"
    authors: [{
        author: "Frankee"
        email: "frankee.zhou@gmail.com"
        organization: "mojolang.org"
    }]
    dependencies: {
        "mojo.core": {
            path: "../core"
            version: "^0.1"
        }
        "mojo.document": {
            path: "../document"
            version: "^0.1"
        }
    }
    repository: "https://github.com/mojo-lang/lang"
}`

	decl := parsePackageDecl(t, typeDecl)
	p := New(&Config{}).PrintPackageDecl(context.Empty(), decl)
	assert.Equal(t, expect, p.Buffer.String())
}
