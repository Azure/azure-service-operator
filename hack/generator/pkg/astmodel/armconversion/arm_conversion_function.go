/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package armconversion

import (
	"fmt"

	"github.com/dave/dst"

	"github.com/Azure/k8s-infra/hack/generator/pkg/astmodel"
)

type ConversionDirection string

const (
	ConversionDirectionToARM   = ConversionDirection("ToARM")
	ConversionDirectionFromARM = ConversionDirection("FromARM")
)

// ARMConversionFunction represents an ARM conversion function for converting between a Kubernetes resource
// and an ARM resource.
type ARMConversionFunction struct {
	name        string
	armTypeName astmodel.TypeName
	armType     *astmodel.ObjectType
	idFactory   astmodel.IdentifierFactory
	direction   ConversionDirection
	isSpecType  bool
}

var _ astmodel.Function = &ARMConversionFunction{}

func (c *ARMConversionFunction) Name() string {
	return c.name
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
func (c *ARMConversionFunction) AsFunc(codeGenerationContext *astmodel.CodeGenerationContext, receiver astmodel.TypeName) *dst.FuncDecl {
	switch c.direction {
	case ConversionDirectionToARM:
		return c.asConvertToARMFunc(codeGenerationContext, receiver, c.Name())
	case ConversionDirectionFromARM:
		return c.asConvertFromARMFunc(codeGenerationContext, receiver, c.Name())
	default:
		panic(fmt.Sprintf("Unknown conversion direction %s", c.direction))
	}
}

func (c *ARMConversionFunction) asConvertToARMFunc(
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string) *dst.FuncDecl {

	builder := newConvertToARMFunctionBuilder(
		c,
		codeGenerationContext,
		receiver,
		methodName)

	return builder.functionDeclaration()
}

func (c *ARMConversionFunction) asConvertFromARMFunc(
	codeGenerationContext *astmodel.CodeGenerationContext,
	receiver astmodel.TypeName,
	methodName string) *dst.FuncDecl {

	builder := newConvertFromARMFunctionBuilder(
		c,
		codeGenerationContext,
		receiver,
		methodName)
	return builder.functionDeclaration()
}

// Equals determines if this function is equal to the passed in function
func (c *ARMConversionFunction) Equals(other astmodel.Function) bool {
	// TODO: Equality on functions is currently awkward because we can't easily pass
	// TODO: a reference to the object the function is on to the function (since both
	// TODO: are immutable and it's impossible to have two immutable objects with
	// TODO: references to each other). Really this equality is always in the context
	// TODO: of comparing objects (since you can't have a free-floating function
	// TODO: with a receiver), and as such we just need to compare things that aren't
	// TODO: the receiver type.

	if o, ok := other.(*ARMConversionFunction); ok {
		return c.armType.Equals(o.armType) &&
			c.armTypeName.Equals(o.armTypeName) &&
			c.direction == o.direction
	}

	return false
}
