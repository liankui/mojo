package syntax

import (
	"testing"

	"github.com/mojo-lang/lang/go/pkg/mojo/lang"
	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/pkg/context"
)

func TestNominalTypeVisitor_VisitArrayType(t *testing.T) {
	const arrayType = `type Val{ val: [Int] }`

	parser := &Parser{}
	file, err := parser.ParseString(context.Empty(), arrayType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Array", nominalType.Name)
	assert.Equal(t, 1, len(nominalType.GenericArguments))
	assert.Equal(t, "Int", nominalType.GenericArguments[0].Name)
}

func TestNominalTypeVisitor_VisitArrayType2(t *testing.T) {
	const arrayType = `type Val{ val: [mojo.lang.StructType] }`

	parser := &Parser{}
	file, err := parser.ParseString(context.Empty(), arrayType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Array", nominalType.Name)
	assert.Equal(t, 1, len(nominalType.GenericArguments))
	assert.Equal(t, "mojo.lang", nominalType.GenericArguments[0].PackageName)
	assert.Equal(t, "StructType", nominalType.GenericArguments[0].Name)
}

func TestNominalTypeVisitor_VisitPrimeType_PackageName(t *testing.T) {
	const primeType = `type Val{ val: mojo.core.Int }`

	parser := &Parser{}
	file, err := parser.ParseString(context.Empty(), primeType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Int", nominalType.Name)
	assert.Equal(t, "mojo.core", nominalType.PackageName)
}

func TestNominalTypeVisitor_VisitType_EnclosingName(t *testing.T) {
	const primeType = `type Val{ val: mojo.core.Url.Path }`

	parser := &Parser{}
	file, err := parser.ParseString(context.Empty(), primeType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Path", nominalType.Name)
	assert.Equal(t, "Url", nominalType.Enclosing.Name)
	assert.Equal(t, "mojo.core", nominalType.PackageName)
}

func TestNominalTypeVisitor_VisitType_EnclosingName2(t *testing.T) {
	const primeType = `type Val{ val: [mojo.core.Url.Path] }`

	parser := &Parser{}
	file, err := parser.ParseString(context.Empty(), primeType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Array", nominalType.Name)
	assert.Equal(t, "Path", nominalType.GenericArguments[0].Name)
	assert.Equal(t, "Url", nominalType.GenericArguments[0].Enclosing.Name)
	assert.Equal(t, "mojo.core", nominalType.GenericArguments[0].PackageName)
}

func TestNominalTypeVisitor_VisitMapType(t *testing.T) {
	const dictType = `type Val{ val: {String: Int} }`

	parser := &Parser{}
	file, err := parser.ParseString(context.Empty(), dictType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Map", nominalType.Name)
	assert.Equal(t, 2, len(nominalType.GenericArguments))
	assert.Equal(t, "String", nominalType.GenericArguments[0].Name)
	assert.Equal(t, "Int", nominalType.GenericArguments[1].Name)
}

func TestNominalTypeVisitor_VisitMapType2(t *testing.T) {
	const dictType = `type Val{ val: {String @label("k"): Int @label("v")} }`

	parser := &Parser{}
	file, err := parser.ParseString(context.Empty(), dictType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Map", nominalType.Name)
	assert.Equal(t, 2, len(nominalType.GenericArguments))
	assert.Equal(t, "String", nominalType.GenericArguments[0].Name)
	assert.Equal(t, "Int", nominalType.GenericArguments[1].Name)
	assert.Equal(t, 1, len(nominalType.GenericArguments[0].Attributes))
	assert.Equal(t, "label", nominalType.GenericArguments[0].Attributes[0].Name)
	kv, _ := nominalType.GenericArguments[0].Attributes[0].GetString()
	assert.Equal(t, "k", kv)

	assert.Equal(t, 1, len(nominalType.GenericArguments[1].Attributes))
	assert.Equal(t, "label", nominalType.GenericArguments[1].Attributes[0].Name)
	kv, _ = nominalType.GenericArguments[1].Attributes[0].GetString()
	assert.Equal(t, "v", kv)
}

func TestNominalTypeVisitor_VisitMapType3(t *testing.T) {
	const dictType = `type Val{ val: {String: Int @label("v")} }`

	parser := &Parser{}
	file, err := parser.ParseString(context.Empty(), dictType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Map", nominalType.Name)
	assert.Equal(t, 2, len(nominalType.GenericArguments))
	assert.Equal(t, "String", nominalType.GenericArguments[0].Name)
	assert.Equal(t, "Int", nominalType.GenericArguments[1].Name)
	assert.Equal(t, 1, len(nominalType.GenericArguments[1].Attributes))
	assert.Equal(t, "label", nominalType.GenericArguments[1].Attributes[0].Name)
}

func TestNominalTypeVisitor_VisitUnion(t *testing.T) {
	const unionType = `type Val{ val: String | Int | [String] }`

	parser := &Parser{}
	file, err := parser.ParseString(context.Empty(), unionType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Union", nominalType.Name)
	assert.Equal(t, 3, len(nominalType.GenericArguments))
	assert.Equal(t, "String", nominalType.GenericArguments[0].Name)
	assert.Equal(t, "Int", nominalType.GenericArguments[1].Name)
	assert.Equal(t, "Array", nominalType.GenericArguments[2].Name)
}

func TestNominalTypeVisitor_VisitUnion2(t *testing.T) {
	const unionType = `type Val{ val: String @1 //< string doc
                                    | Int
                                    | [String] @3 //< strings doc
}`

	parser := &Parser{}
	file, err := parser.ParseString(context.Empty(), unionType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Union", nominalType.Name)
	assert.Equal(t, 3, len(nominalType.GenericArguments))
	assert.Equal(t, "String", nominalType.GenericArguments[0].Name)
	assert.Equal(t, "string doc", nominalType.GenericArguments[0].Document.Lines[0].Content)
	assert.Equal(t, "Int", nominalType.GenericArguments[1].Name)
	assert.Equal(t, "Array", nominalType.GenericArguments[2].Name)
	assert.Equal(t, "strings doc", nominalType.GenericArguments[2].Document.Lines[0].Content)
}

func TestNominalTypeVisitor_VisitTupleType(t *testing.T) {
	const tupleType = `type Val{ val: (String, Int) }`

	parser := &Parser{}
	file, err := parser.ParseString(context.Empty(), tupleType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Tuple", nominalType.Name)
	assert.Equal(t, 2, len(nominalType.GenericArguments))
	assert.Equal(t, "String", nominalType.GenericArguments[0].Name)
	assert.Equal(t, "Int", nominalType.GenericArguments[1].Name)
}

func TestNominalTypeVisitor_VisitTupleTypeWithLabel(t *testing.T) {
	const tupleType = `type Val{ val: (str: String, integer: Int) }`

	parser := &Parser{}
	file, err := parser.ParseString(context.Empty(), tupleType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Tuple", nominalType.Name)
	assert.Equal(t, 2, len(nominalType.GenericArguments))
	assert.Equal(t, "String", nominalType.GenericArguments[0].Name)
	assert.Equal(t, "Int", nominalType.GenericArguments[1].Name)

	label, _ := lang.GetStringAttribute(nominalType.GenericArguments[0].Attributes, "label")
	assert.Equal(t, "str", label)

	label, _ = lang.GetStringAttribute(nominalType.GenericArguments[1].Attributes, "label")
	assert.Equal(t, "integer", label)
}

func TestNominalTypeVisitor_VisitQuestionType(t *testing.T) {
	const questionType = `type Val{ val: Int? }`

	parser := &Parser{}
	file, err := parser.ParseString(context.Empty(), questionType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Int", nominalType.Name)
	v, _ := lang.GetBoolAttribute(nominalType.Attributes, "optional")
	assert.True(t, v)
}

func TestNominalTypeVisitor_VisitBangType(t *testing.T) {
	const bangType = `type Val{ val: Int! }`

	parser := &Parser{}
	file, err := parser.ParseString(context.Empty(), bangType)

	assert.NoError(t, err)
	nominalType := getNominalType(file)
	assert.NotNil(t, nominalType)
	assert.Equal(t, "Int", nominalType.Name)
	v, _ := lang.GetBoolAttribute(nominalType.Attributes, "required")
	assert.True(t, v)
}

func getNominalType(file *lang.SourceFile) *lang.NominalType {
	if len(file.Statements) > 0 {
		statement := file.Statements[0]
		if decl := statement.GetDeclaration(); decl != nil {
			if structDecl := decl.GetStructDecl(); structDecl != nil {
				if t := structDecl.Type; t != nil {
					if len(t.Fields) > 0 {
						if field := t.Fields[0]; field != nil {
							return field.Type
						}
					}
				}
			}
		}
	}
	return nil
}
