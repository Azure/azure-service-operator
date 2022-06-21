/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

import (
	"fmt"
	"strings"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// PropertyBagMemberType is a wrapper type for properties that are stored-in/recalled-from a PropertyBag
// Using this custom type allows us to cleanly define endpoints that read-from/write-to the property bag without needing
// to introduce special cases through the conversions package. Instead, code generation of the conversion steps now
// naturally falls out of the usual flow - the endpoints contain definitions, and handlers for those definitions pick up the context
// and generate the required code.
type PropertyBagMemberType struct {
	// element is the wrapped type of the property
	element astmodel.Type
}

// PropertyBagMemberType must implement Type
var _ astmodel.Type = &PropertyBagMemberType{}

// PropertyBagMemberType must implement MetaType
var _ astmodel.MetaType = &PropertyBagMemberType{}

// NewPropertyBagMemberType returns a new PropertyBagMemberType, ready for use by the conversion framework
// A bag item can't be optional, so if we're given an optional type we just unwrap it
func NewPropertyBagMemberType(element astmodel.Type) *PropertyBagMemberType {
	e := element
	if opt, ok := astmodel.AsOptionalType(e); ok {
		// Unwrap the optional type
		e = opt.Element()
	}

	return &PropertyBagMemberType{
		element: e,
	}
}

// Element returns the type contained by the bag item
func (b *PropertyBagMemberType) Element() astmodel.Type {
	return b.element
}

// RequiredPackageReferences returns a set of packages imports required by the element we wrap
func (b *PropertyBagMemberType) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	return b.element.RequiredPackageReferences()
}

// References returns the names of all definitions that the element type we wrap references.
func (b *PropertyBagMemberType) References() astmodel.TypeNameSet {
	return b.element.References()
}

// AsType renders as our contained type
func (b *PropertyBagMemberType) AsType(ctx *astmodel.CodeGenerationContext) dst.Expr {
	return b.element.AsType(ctx)
}

// AsDeclarations panics because this is a metatype that will never be rendered
func (b *PropertyBagMemberType) AsDeclarations(_ *astmodel.CodeGenerationContext, _ astmodel.DeclarationContext) []dst.Decl {
	panic("should never try to render a PropertyBagMemberType as declarations")
}

// AsZero renders an expression for the "zero" value of the type
func (b *PropertyBagMemberType) AsZero(definitions astmodel.TypeDefinitionSet, ctx *astmodel.CodeGenerationContext) dst.Expr {
	return b.element.AsZero(definitions, ctx)
}

// Equals returns true if the passed type is a PropertyBagMemberType with the same element this one, false otherwise
func (b *PropertyBagMemberType) Equals(t astmodel.Type, override astmodel.EqualityOverrides) bool {
	other, ok := t.(*PropertyBagMemberType)
	return ok && other.element.Equals(b.element, override)
}

// String returns a string representing the type
func (b *PropertyBagMemberType) String() string {
	return fmt.Sprintf("PropertyBagMemberType(%s)", b.element.String())
}

// WriteDebugDescription adds a description of the current type to the passed builder
// builder receives the full description, including nested definitions
// definitions is a dictionary for resolving named definitions
func (b *PropertyBagMemberType) WriteDebugDescription(builder *strings.Builder, currentPackage astmodel.PackageReference) {
	builder.WriteString("Bag[")
	b.element.WriteDebugDescription(builder, currentPackage)
	builder.WriteString("]")
}

// Unwrap returns the type contained within the wrapper type
func (b *PropertyBagMemberType) Unwrap() astmodel.Type {
	return b.element
}

// AsPropertyBagMemberType unwraps any wrappers around the provided type and returns either the underlying PropertyBagMemberType and true,
// or nil and false.
func AsPropertyBagMemberType(aType astmodel.Type) (*PropertyBagMemberType, bool) {
	if obj, ok := aType.(*PropertyBagMemberType); ok {
		return obj, true
	}

	if wrapper, ok := aType.(astmodel.MetaType); ok {
		return AsPropertyBagMemberType(wrapper.Unwrap())
	}

	return nil, false
}
