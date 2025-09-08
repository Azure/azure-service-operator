/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import "fmt"

type TypeFlag string

const (
	StorageFlag       = TypeFlag("storage")
	ARMFlag           = TypeFlag("arm")
	OneOfFlag         = TypeFlag("oneof")
	DoNotPrune        = TypeFlag("donotprune")
	WellknownFlag     = TypeFlag("wellknown")  // Indicates a resource reference that should support well-known names
	CompatibilityFlag = TypeFlag("compatible") // Indicates a resource reference that needs to maintain backwards compatibility
)

var _ fmt.Stringer = TypeFlag("")

// String renders the tag as a string
func (f TypeFlag) String() string {
	return string(f)
}

// ApplyTo applies the tag to the provided type
func (f TypeFlag) ApplyTo(t Type) *FlaggedType {
	return NewFlaggedType(t, f)
}

// RemoveFrom applies the tag to the provided type
func (f TypeFlag) RemoveFrom(t Type) (Type, error) {
	removeFlag := func(it *FlaggedType) Type {
		return it.WithoutFlag(f)
	}

	visitor := TypeVisitorBuilder[any]{
		VisitFlaggedType: removeFlag,
	}.Build()

	return visitor.Visit(t, nil)
}

// IsOn returns true if t is or contains a flagged type that has this flag
func (f TypeFlag) IsOn(theType Type) bool {
	switch t := theType.(type) {
	case *FlaggedType:
		return t.HasFlag(f)

	case MetaType:
		return f.IsOn(t.Unwrap())
	default:
		return false
	}
}
