/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package conversions

import (
	"fmt"
	"strings"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/astmodel"
)

// BagItemType is a wrapper type for properties that are stored-in/recalled-from a PropertyBag
type BagItemType struct {
	// element is the wrapped type of the property
	element astmodel.Type
}

// BagItemType must implement Type
var _ astmodel.Type = &BagItemType{}

// BagItemType must implement MetaType
var _ astmodel.MetaType = &BagItemType{}

// NewBagItemType returns a new BagItemType, ready for use by the conversion framework
// A bag item can't be optional, so if we're given an optional type we just unwrap it
func NewBagItemType(element astmodel.Type) *BagItemType {

	e := element
	if opt, ok := astmodel.AsOptionalType(e); ok {
		e = opt.Element()
	}

	return &BagItemType{
		element: e,
	}
}

// Element returns the type contained by the bag item
func (b *BagItemType) Element() astmodel.Type {
	return b.element
}

// RequiredPackageReferences returns a set of packages imports required by the element we wrap
func (b *BagItemType) RequiredPackageReferences() *astmodel.PackageReferenceSet {
	return b.element.RequiredPackageReferences()
}

// References returns the names of all types that the element type we wrap references.
func (b *BagItemType) References() astmodel.TypeNameSet {
	return b.element.References()
}

// AsType renders as our contained type
func (b *BagItemType) AsType(ctx *astmodel.CodeGenerationContext) dst.Expr {
	return b.element.AsType(ctx)
}

// AsDeclarations panics because this is a metatype that will never be rendered
func (b *BagItemType) AsDeclarations(_ *astmodel.CodeGenerationContext, _ astmodel.DeclarationContext) []dst.Decl {
	panic("should never try to render a BagItemType as declarations")
}

// AsZero renders an expression for the "zero" value of the type
func (b *BagItemType) AsZero(types astmodel.Types, ctx *astmodel.CodeGenerationContext) dst.Expr {
	return b.element.AsZero(types, ctx)
}

// Equals returns true if the passed type is a BagItemType with the same element this one, false otherwise
func (b *BagItemType) Equals(t astmodel.Type) bool {
	other, ok := t.(*BagItemType)
	return ok && other.element.Equals(b.element)
}

// String returns a string representing the type
func (b *BagItemType) String() string {
	return fmt.Sprintf("BagItemType(%s)", b.element.String())
}

// WriteDebugDescription adds a description of the current type to the passed builder
// builder receives the full description, including nested types
// types is a dictionary for resolving named types
func (b *BagItemType) WriteDebugDescription(builder *strings.Builder, types astmodel.Types) {
	builder.WriteString("Optional[")
	b.element.WriteDebugDescription(builder, types)
	builder.WriteString("]")
}

// Unwrap returns the type contained within the wrapper type
func (b *BagItemType) Unwrap() astmodel.Type {
	return b.element
}

// AsBagItemType unwraps any wrappers around the provided type and returns either the underlying BagItemType and true,
// or nil and false.
func AsBagItemType(aType astmodel.Type) (*BagItemType, bool) {
	if obj, ok := aType.(*BagItemType); ok {
		return obj, true
	}

	if wrapper, ok := aType.(astmodel.MetaType); ok {
		return AsBagItemType(wrapper.Unwrap())
	}

	return nil, false
}
