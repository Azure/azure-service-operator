/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package armconversion

import (
	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// ARMConversionFunction represents an ARM conversion function for converting between a Kubernetes resource
// and an ARM resource.
type ARMConversionFunction struct {
	armTypeName astmodel.TypeName
	armType     *astmodel.ObjectType
	idFactory   astmodel.IdentifierFactory
	typeKind    TypeKind
}

type ConvertToARMFunction struct {
	ARMConversionFunction
}

type PopulateFromARMFunction struct {
	ARMConversionFunction
}

var (
	_ astmodel.Function = &ConvertToARMFunction{}
	_ astmodel.Function = &PopulateFromARMFunction{}
)

func (c *ConvertToARMFunction) Name() string {
	return "ConvertToARM"
}

func (c *PopulateFromARMFunction) Name() string {
	return "PopulateFromARM"
}

// RequiredPackageReferences returns the imports required for this conversion function
func (c *ARMConversionFunction) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	// Because this interface is attached to an object, by definition that object will specify
	// its own required imports. We don't want to call the objects required imports here
	// because then we're in infinite recursion (object delegates to us, we delegate back to it)

	// We need these because we're going to be constructing/casting to the types
	// of the properties in the ARM object, so we need to import those.
	result := astmodel.NewPackageReferenceSet(
		astmodel.GenRuntimeReference,
		astmodel.GitHubErrorsReference,
		astmodel.MakeExternalPackageReference("fmt"))
	result.Merge(c.armType.RequiredPackageReferences())
	return result
}

// References returns the set of types to which this function refers.
// SHOULD include any types which this function references but its receiver doesn't.
// SHOULD NOT include the receiver of this function.
func (c *ARMConversionFunction) References() astmodel.TypeNameSet[astmodel.TypeName] {
	return c.armType.References()
}

// AsFunc returns the function as a Go AST
func (c *ConvertToARMFunction) AsFunc(
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
) *dst.FuncDecl {
	builder := newConvertToARMFunctionBuilder(
		&c.ARMConversionFunction,
		codeGenerationContext,
		receiver,
		c.Name())

	decl, err := builder.functionDeclaration()
	if err != nil {
		// TODO: This will become an error return when we refactor the conversion functions for issue #2971
		panic(err)
	}

	return decl
}

func (c *PopulateFromARMFunction) AsFunc(
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
) *dst.FuncDecl {
	builder := newConvertFromARMFunctionBuilder(
		&c.ARMConversionFunction,
		codeGenerationContext,
		receiver,
		c.Name())

	decl, err := builder.functionDeclaration()
	if err != nil {
		// TODO: This will become an error return when we refactor the conversion functions for issue #2971
		panic(err)
	}

	return decl
}

// Equals determines if this function is equal to the passed in function
func (c *ConvertToARMFunction) Equals(other astmodel.Function, overrides astmodel.EqualityOverrides) bool {
	// TODO: Equality on functions is currently awkward because we can't easily pass
	// TODO: a reference to the object the function is on to the function (since both
	// TODO: are immutable and it's impossible to have two immutable objects with
	// TODO: references to each other). Really this equality is always in the context
	// TODO: of comparing objects (since you can't have a free-floating function
	// TODO: with a receiver), and as such we just need to compare things that aren't
	// TODO: the receiver type.

	if o, ok := other.(*ConvertToARMFunction); ok {
		return c.armType.Equals(o.armType, overrides) &&
			c.armTypeName.Equals(o.armTypeName, overrides)
	}

	return false
}

// Equals determines if this function is equal to the passed in function
func (c *PopulateFromARMFunction) Equals(other astmodel.Function, overrides astmodel.EqualityOverrides) bool {
	// TODO: Equality on functions is currently awkward because we can't easily pass
	// TODO: a reference to the object the function is on to the function (since both
	// TODO: are immutable and it's impossible to have two immutable objects with
	// TODO: references to each other). Really this equality is always in the context
	// TODO: of comparing objects (since you can't have a free-floating function
	// TODO: with a receiver), and as such we just need to compare things that aren't
	// TODO: the receiver type.

	if o, ok := other.(*PopulateFromARMFunction); ok {
		return c.armType.Equals(o.armType, overrides) &&
			c.armTypeName.Equals(o.armTypeName, overrides)
	}

	return false
}
