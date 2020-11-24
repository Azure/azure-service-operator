/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import "fmt"

type TypeFlag string

const (
	StorageFlag = TypeFlag("storage")
	ArmFlag     = TypeFlag("arm")
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

// IsOn returns true if t is a flagged type that has this flag
func (f TypeFlag) IsOn(t Type) bool {
	if ft, ok := t.(*FlaggedType); ok {
		return ft.HasFlag(f)
	}

	return false
}
