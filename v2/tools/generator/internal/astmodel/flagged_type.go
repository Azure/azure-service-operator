/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"strings"

	"github.com/dave/dst"

	"github.com/Azure/azure-service-operator/v2/internal/set"
)

type FlaggedType struct {
	element Type
	flags   set.Set[TypeFlag]
}

var _ Type = &FlaggedType{}

var _ MetaType = &FlaggedType{}

// NewFlaggedType applies flags to an existing type and returns a wrapper
func NewFlaggedType(t Type, flags ...TypeFlag) *FlaggedType {
	result := &FlaggedType{
		element: t,
		flags:   set.Make[TypeFlag](),
	}

	if ft, ok := t.(*FlaggedType); ok {
		// It's flagged type, so unwrap to avoid nesting
		result.element = ft.element
		for f := range ft.flags {
			result.flags.Add(f)
		}
	}

	for _, f := range flags {
		result.flags.Add(f)
	}

	return result
}

// Element returns the underlying type that is flagged
func (ft *FlaggedType) Element() Type {
	return ft.element
}

// HasFlag tests to see if this flagged type has the specified flag
func (ft *FlaggedType) HasFlag(flag TypeFlag) bool {
	return ft.flags.Contains(flag)
}

// WithFlag returns a new FlaggedType with the specified flag added
func (ft *FlaggedType) WithFlag(flag TypeFlag) *FlaggedType {
	return NewFlaggedType(ft, flag)
}

// WithoutFlag returns a new FlaggedType with the specified flag excluded
// If the flag is not present, the existing FlaggedType is returned unmodified
// If the last flag is removed, the underlying type is returned
func (ft *FlaggedType) WithoutFlag(flag TypeFlag) Type {
	if !ft.HasFlag(flag) {
		return ft
	}

	if len(ft.flags) == 1 {
		return ft.element
	}

	flags := ft.flags.Copy()
	flags.Remove(flag)

	return &FlaggedType{
		element: ft.element,
		flags:   flags,
	}
}

// WithElement returns the flagged type with the same flags but a different element
func (ft *FlaggedType) WithElement(t Type) *FlaggedType {
	if TypeEquals(ft.element, t) {
		return ft // short-circuit
	}

	return NewFlaggedType(t, ft.flags.Values()...)
}

// RequiredPackageReferences returns a set of packages imports required by this type
func (ft *FlaggedType) RequiredPackageReferences() *PackageReferenceSet {
	return ft.element.RequiredPackageReferences()
}

// References returns the names of all types that this type
// references.
func (ft *FlaggedType) References() TypeNameSet {
	return ft.element.References()
}

// AsType renders as a Go abstract syntax tree for a type
// (yes this says ast.Expr but that is what the Go 'dst' package uses for types)
func (ft *FlaggedType) AsType(ctx *CodeGenerationContext) dst.Expr {
	return ft.element.AsType(ctx)
}

// AsDeclarations renders as a Go abstract syntax tree for a declaration
func (ft *FlaggedType) AsDeclarations(ctx *CodeGenerationContext, declContext DeclarationContext) []dst.Decl {
	return ft.element.AsDeclarations(ctx, declContext)
}

// AsZero renders an expression for the "zero" value of the type
// by delegating to the wrapped type
func (ft *FlaggedType) AsZero(definitions TypeDefinitionSet, ctx *CodeGenerationContext) dst.Expr {
	return ft.element.AsZero(definitions, ctx)
}

// Equals returns true if the passed type is the same as this one, false otherwise
func (ft *FlaggedType) Equals(t Type, overrides EqualityOverrides) bool {
	if ft == t {
		return true // short-circuit
	}

	other, ok := t.(*FlaggedType)
	if !ok {
		return false
	}

	if !set.AreEqual(ft.flags, other.flags) {
		return false
	}

	return ft.element.Equals(other.element, overrides)
}

// String returns the underlying flagged type plus the flags present
// For this to be useful, make sure all Types have a printable version for debugging/user info.
// This doesn't need to be a full representation of the type.
func (ft *FlaggedType) String() string {
	var result strings.Builder
	result.WriteString(ft.element.String())

	if len(ft.flags) > 0 {
		flags := set.AsSortedSlice(ft.flags)

		result.WriteRune('[')
		for i, f := range flags {
			if i > 0 {
				result.WriteRune('|')
			}
			result.WriteString(f.String())
		}
		result.WriteRune(']')
	}

	return result.String()
}

// Unwrap returns the type contained within the wrapper type
func (ft *FlaggedType) Unwrap() Type {
	return ft.element
}

// WriteDebugDescription adds a description of the current type to the passed builder
// builder receives the full description, including nested types.
// definitions is a dictionary for resolving named types.
func (ft *FlaggedType) WriteDebugDescription(builder *strings.Builder, definitions TypeDefinitionSet) {
	if ft == nil {
		builder.WriteString("<nilFlagged>")
		return
	}

	ft.element.WriteDebugDescription(builder, definitions)
	for f := range ft.flags {
		builder.WriteString("[Flag:")
		builder.WriteString(f.String())
		builder.WriteString("]")
	}
}
