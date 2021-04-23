/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import "fmt"

type TypeFlag string

const (
	StorageFlag = TypeFlag("storage")
	ARMFlag     = TypeFlag("arm")
	OneOfFlag   = TypeFlag("oneof")
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

// ApplyTo applies the tag to the provided type
func (f TypeFlag) RemoveFrom(t Type) (Type, error) {
	visitor := MakeTypeVisitor()
	visitor.VisitFlaggedType = func(this *TypeVisitor, it *FlaggedType, ctx interface{}) (Type, error) {
		return it.WithoutFlag(f), nil
	}

	return visitor.Visit(t, nil)
}

// IsOn returns true if t is a flagged type that has this flag
func (f TypeFlag) IsOn(t Type) bool {
	if ft, ok := t.(*FlaggedType); ok {
		return ft.HasFlag(f)
	}

	return false
}
