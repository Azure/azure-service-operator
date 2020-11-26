/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"go/token"
	"sort"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astbuilder"
	ast "github.com/dave/dst"
	"k8s.io/klog/v2"
)

// EnumType represents a set of mutually exclusive predefined options
type EnumType struct {
	// BaseType is the underlying type used to define the values
	baseType *PrimitiveType
	// Options is the set of all unique values
	options []EnumValue
}

// EnumType must implement the Type interface correctly
var _ Type = (*EnumType)(nil)

// NewEnumType defines a new enumeration including the legal values
func NewEnumType(baseType *PrimitiveType, options []EnumValue) *EnumType {
	if baseType == nil {
		panic("baseType must be provided")
	}

	sort.Slice(options, func(left int, right int) bool {
		return options[left].Identifier < options[right].Identifier
	})

	return &EnumType{baseType: baseType, options: options}
}

// AsDeclarations converts the EnumType to a series of Go AST Decls
func (enum *EnumType) AsDeclarations(codeGenerationContext *CodeGenerationContext, name TypeName, description []string) []ast.Decl {
	result := []ast.Decl{enum.createBaseDeclaration(codeGenerationContext, name, description)}

	var specs []ast.Spec
	for _, v := range enum.options {
		s := enum.createValueDeclaration(name, v)
		specs = append(specs, s)
	}

	if len(specs) > 0 {
		declaration := &ast.GenDecl{
			Tok:   token.CONST,
			Specs: specs,
		}

		result = append(result, declaration)
	}

	return result
}

func (enum *EnumType) createBaseDeclaration(
	codeGenerationContext *CodeGenerationContext,
	name TypeName,
	description []string) ast.Decl {

	typeSpecification := &ast.TypeSpec{
		Name: ast.NewIdent(name.Name()),
		Type: enum.baseType.AsType(codeGenerationContext),
	}

	declaration := &ast.GenDecl{
		Decs: ast.GenDeclDecorations{
			NodeDecs: ast.NodeDecs{
				Before: ast.EmptyLine,
			},
		},
		Tok: token.TYPE,
		Specs: []ast.Spec{
			typeSpecification,
		},
	}

	astbuilder.AddWrappedComments(&declaration.Decs.Start, description, 120)

	validationComment := GenerateKubebuilderComment(enum.CreateValidation())
	astbuilder.AddComment(&declaration.Decs.Start, validationComment)

	return declaration
}

func (enum *EnumType) createValueDeclaration(name TypeName, value EnumValue) ast.Spec {

	valueSpec := &ast.ValueSpec{
		Names: []*ast.Ident{ast.NewIdent(GetEnumValueId(name.name, value))},
		Values: []ast.Expr{
			astbuilder.CallFunc(name.Name(), astbuilder.TextLiteral(value.Value)),
		},
	}

	return valueSpec
}

// AsType implements Type for EnumType
func (enum *EnumType) AsType(codeGenerationContext *CodeGenerationContext) ast.Expr {
	// this should "never" happen as we name all enums; warn about it if it does
	klog.Warning("Emitting unnamed enum, something’s awry")
	return enum.baseType.AsType(codeGenerationContext)
}

// References returns any types the underlying type refers to.
func (enum *EnumType) References() TypeNameSet {
	return enum.baseType.References()
}

// Equals will return true if the supplied type has the same base type and options
func (enum *EnumType) Equals(t Type) bool {
	if e, ok := t.(*EnumType); ok {
		if !enum.baseType.Equals(e.baseType) {
			return false
		}

		if len(enum.options) != len(e.options) {
			// Different number of properties, not equal
			return false
		}

		for i := range enum.options {
			if !enum.options[i].Equals(&e.options[i]) {
				return false
			}
		}

		// All options match, equal
		return true
	}

	return false
}

// RequiredPackageReferences indicates that Enums never need additional imports
func (enum *EnumType) RequiredPackageReferences() *PackageReferenceSet {
	return NewPackageReferenceSet()
}

// Options returns all the enum options
// A copy of the slice is returned to preserve immutability
func (enum *EnumType) Options() []EnumValue {
	return append(enum.options[:0:0], enum.options...)
}

// CreateValidation creates the validation annotation for this Enum
func (enum *EnumType) CreateValidation() Validation {
	var values []interface{}
	for _, opt := range enum.Options() {
		values = append(values, opt.Value)
	}

	return ValidateEnum(values)
}

// BaseType returns the base type of the enum
func (enum *EnumType) BaseType() *PrimitiveType {
	return enum.baseType
}

func GetEnumValueId(name string, value EnumValue) string {
	return name + value.Identifier
}

// String implements fmt.Stringer
func (enum *EnumType) String() string {
	return fmt.Sprintf("(enum: %s)", enum.baseType.String())
}
