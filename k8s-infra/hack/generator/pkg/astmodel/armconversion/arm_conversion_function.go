/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package armconversion

import (
	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
	"github.com/dave/dst"
)

// ARMConversionFunction represents an ARM conversion function for converting between a Kubernetes resource
// and an ARM resource.
type ARMConversionFunction struct {
	armTypeName astmodel.TypeName
	armType     *astmodel.ObjectType
	idFactory   astmodel.IdentifierFactory
	isSpecType  bool
}

type ConvertToARMFunction struct {
	ARMConversionFunction
}

type PopulateFromARMFunction struct {
	ARMConversionFunction
}

var _ astmodel.Function = &ConvertToARMFunction{}
var _ astmodel.Function = &PopulateFromARMFunction{}

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
		astmodel.MakeExternalPackageReference("fmt"))
	result.Merge(c.armType.RequiredPackageReferences())
	return result
}

// References this type has to the given type
func (c *ARMConversionFunction) References() astmodel.TypeNameSet {
	return c.armType.References()
}

// AsFunc returns the function as a Go AST
func (c *ConvertToARMFunction) AsFunc(codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName) *dst.FuncDecl {
	builder := newConvertToARMFunctionBuilder(
		&c.ARMConversionFunction,
		codeGenerationContext,
		receiver,
		c.Name())

	return builder.functionDeclaration()
}

func (c *PopulateFromARMFunction) AsFunc(codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName) *dst.FuncDecl {
	builder := newConvertFromARMFunctionBuilder(
		&c.ARMConversionFunction,
		codeGenerationContext,
		receiver,
		c.Name())

	return builder.functionDeclaration()
}

// Equals determines if this function is equal to the passed in function
func (c *ConvertToARMFunction) Equals(other astmodel.Function) bool {
	// TODO: Equality on functions is currently awkward because we can't easily pass
	// TODO: a reference to the object the function is on to the function (since both
	// TODO: are immutable and it's impossible to have two immutable objects with
	// TODO: references to each other). Really this equality is always in the context
	// TODO: of comparing objects (since you can't have a free-floating function
	// TODO: with a receiver), and as such we just need to compare things that aren't
	// TODO: the receiver type.

	if o, ok := other.(*ConvertToARMFunction); ok {
		return c.armType.Equals(o.armType) && c.armTypeName.Equals(o.armTypeName)
	}

	return false
}

// Equals determines if this function is equal to the passed in function
func (c *PopulateFromARMFunction) Equals(other astmodel.Function) bool {
	// TODO: Equality on functions is currently awkward because we can't easily pass
	// TODO: a reference to the object the function is on to the function (since both
	// TODO: are immutable and it's impossible to have two immutable objects with
	// TODO: references to each other). Really this equality is always in the context
	// TODO: of comparing objects (since you can't have a free-floating function
	// TODO: with a receiver), and as such we just need to compare things that aren't
	// TODO: the receiver type.

	if o, ok := other.(*PopulateFromARMFunction); ok {
		return c.armType.Equals(o.armType) && c.armTypeName.Equals(o.armTypeName)
	}

	return false
}
