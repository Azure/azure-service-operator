/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"go/token"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astbuilder"
	ast "github.com/dave/dst"
	"github.com/pkg/errors"
)

// TypeDefinition is a name paired with a type
type TypeDefinition struct {
	name        TypeName
	description []string
	theType     Type
}

func MakeTypeDefinition(name TypeName, theType Type) TypeDefinition {
	return TypeDefinition{name: name, theType: theType}
}

// Name returns the name being associated with the type
func (def TypeDefinition) Name() TypeName {
	return def.name
}

// Type returns the type being associated with the name
func (def TypeDefinition) Type() Type {
	return def.theType
}

// Description returns the description to be attached to this type definition (as a comment)
// We return a new slice to preserve immutability
func (def TypeDefinition) Description() []string {
	var result []string
	result = append(result, def.description...)
	return result
}

func (def TypeDefinition) References() TypeNameSet {
	return def.theType.References()
}

// WithDescription replaces the description of the definition with a new one (if any)
func (def TypeDefinition) WithDescription(desc []string) TypeDefinition {
	var d []string
	def.description = append(d, desc...)
	return def
}

// WithType returns an updated TypeDefinition with the specified type
func (def TypeDefinition) WithType(t Type) TypeDefinition {
	result := def
	result.theType = t
	return result
}

// WithName returns an updated TypeDefinition with the specified name
func (def TypeDefinition) WithName(typeName TypeName) TypeDefinition {
	result := def
	result.name = typeName
	return result
}

func (def TypeDefinition) AsDeclarations(codeGenerationContext *CodeGenerationContext) []ast.Decl {
	declContext := DeclarationContext{
		Name:        def.name,
		Description: def.description,
	}

	return def.theType.AsDeclarations(codeGenerationContext, declContext)
}

// AsSimpleDeclarations is a helper for types that only require a simple name/alias to be defined
func AsSimpleDeclarations(
	codeGenerationContext *CodeGenerationContext,
	declContext DeclarationContext,
	theType Type) []ast.Decl {

	var docComments ast.Decorations
	if len(declContext.Description) > 0 {
		astbuilder.AddWrappedComments(&docComments, declContext.Description, 120)
	}

	AddValidationComments(&docComments, declContext.Validations)

	result := &ast.GenDecl{
		Decs: ast.GenDeclDecorations{
			NodeDecs: ast.NodeDecs{
				Start:  docComments,
				Before: ast.EmptyLine,
			},
		},
		Tok: token.TYPE,
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: ast.NewIdent(declContext.Name.Name()),
				Type: theType.AsType(codeGenerationContext),
			},
		},
	}

	return []ast.Decl{result}
}

// RequiredImports returns a list of packages required by this type
func (def TypeDefinition) RequiredPackageReferences() *PackageReferenceSet {
	return def.theType.RequiredPackageReferences()
}

func (def TypeDefinition) HasTestCases() bool {
	switch d := def.theType.(type) {
	case *ObjectType:
		return d.HasTestCases()
	case *ResourceType:
		return d.HasTestCases()
	}

	return false
}

// FileNameHint returns what a file that contains this name (if any) should be called
// this is not always used as we often combine multiple definitions into one file
func FileNameHint(name TypeName) string {
	return transformToSnakeCase(name.name)
}

// ApplyObjectTransformation applies a specific transformation to the ObjectType contained by this
// definition, returning a new definition
// If the definition does not contain an object, an error will be returned
func (def TypeDefinition) ApplyObjectTransformation(transform func(*ObjectType) (Type, error)) (TypeDefinition, error) {
	// We use a TypeVisitor to allow automatic handling of wrapper types (such as ArmType and StorageType)
	visited := false
	visitor := MakeTypeVisitor()
	visitor.VisitObjectType = func(_ *TypeVisitor, ot *ObjectType, _ interface{}) (Type, error) {
		rt, err := transform(ot)
		if err != nil {
			return nil, err
		}
		visited = true
		return rt, nil
	}

	newType, err := visitor.Visit(def.theType, nil)
	if err != nil {
		return TypeDefinition{}, errors.Wrapf(err, "transformation of %v failed", def.name)
	}

	if !visited {
		return TypeDefinition{}, errors.Errorf("transformation was not applied to %v (expected object type, found %v)", def.name, def.theType)
	}

	result := def.WithType(newType)
	return result, nil
}

// ApplyObjectTransformations applies multiple transformations to the ObjectType contained by this
// definition, returning a new definition.
// If the definition does not contain an object, an error will be returned
// The transformations are constrained to return ObjectType results to allow them to be chained together.
func (def TypeDefinition) ApplyObjectTransformations(transforms ...func(*ObjectType) (*ObjectType, error)) (TypeDefinition, error) {
	return def.ApplyObjectTransformation(func(objectType *ObjectType) (Type, error) {
		result := objectType
		for i, transform := range transforms {
			rt, err := transform(result)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to apply object transformation %d", i)
			}

			result = rt
		}

		return result, nil
	})
}
