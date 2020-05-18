package astmodel

import (
	"go/ast"
	"go/token"
	"sort"
)

// EnumDefinition generates the full definition of an enumeration
type EnumDefinition struct {
	EnumType
}

var _ Definition = (*EnumDefinition)(nil)

// FileNameHint returns a desired name for this enum if it goes into a standalone file
func (enum *EnumDefinition) FileNameHint() string {
	return enum.name
}

// Reference returns the unique name to use for specifying this enumeration
func (enum *EnumDefinition) Reference() *DefinitionName {
	return &enum.DefinitionName
}

// Type returns the underlying EnumerationType for this enum
func (enum *EnumDefinition) Type() Type {
	return &enum.EnumType
}

// AsDeclarations generates the Go code representing this definition
func (enum *EnumDefinition) AsDeclarations() []ast.Decl {
	var specs []ast.Spec
	for _, v := range enum.Options {
		s := enum.createValueDeclaration(v)
		specs = append(specs, s)
	}

	declaration := &ast.GenDecl{
		Tok:   token.CONST,
		Doc:   &ast.CommentGroup{},
		Specs: specs,
	}

	result := []ast.Decl{
		enum.createBaseDeclaration(),
		declaration}

	return result
}

// Tidy does cleanup to ensure deterministic code generation
func (enum *EnumDefinition) Tidy() {
	sort.Slice(enum.Options, func(left int, right int) bool {
		return enum.Options[left].Identifier < enum.Options[right].Identifier
	})
}

func (enum *EnumDefinition) createBaseDeclaration() ast.Decl {
	var identifier *ast.Ident
	identifier = ast.NewIdent(enum.name)

	typeSpecification := &ast.TypeSpec{
		Name: identifier,
		Type: enum.BaseType.AsType(),
	}

	declaration := &ast.GenDecl{
		Tok: token.TYPE,
		Doc: &ast.CommentGroup{},
		Specs: []ast.Spec{
			typeSpecification,
		},
	}

	return declaration
}

func (enum *EnumDefinition) createValueDeclaration(value EnumValue) ast.Spec {

	var enumIdentifier *ast.Ident
	enumIdentifier = ast.NewIdent(enum.name)

	valueIdentifier := ast.NewIdent(enum.Name() + value.Identifier)
	valueLiteral := ast.BasicLit{
		Kind:  token.STRING,
		Value: value.Value,
	}

	valueSpec := &ast.ValueSpec{
		Names: []*ast.Ident{valueIdentifier},
		Values: []ast.Expr{
			&ast.CallExpr{
				Fun:  enumIdentifier,
				Args: []ast.Expr{&valueLiteral},
			},
		},
	}

	return valueSpec
}
