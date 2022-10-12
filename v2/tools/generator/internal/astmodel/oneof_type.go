/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"sort"
	"strings"

	"github.com/dave/dst"
	"github.com/pkg/errors"
)

// OneOfType represents something that can be any
// one of a number of selected types
type OneOfType struct {
	// invariants:
	// - all types are unique (enforced by TypeSet)
	// - length > 1
	// - no nested OneOfs (aside from indirectly via TypeName)
	types TypeSet
}

var _ Type = &OneOfType{}

// BuildOneOfType is a smart constructor for a OneOfType,
// maintaining the invariants. If only one unique type
// is passed, the result will be that type, not a OneOf.
func BuildOneOfType(types ...Type) Type {
	uniqueTypes := MakeTypeSet()
	for _, t := range types {
		if oneOf, ok := t.(*OneOfType); ok {
			oneOf.types.ForEach(func(t Type, _ int) {
				uniqueTypes.Add(t)
			})
		} else {
			uniqueTypes.Add(t)
		}
	}

	if uniqueTypes.Len() == 1 {
		var result Type
		uniqueTypes.ForEach(func(t Type, _ int) {
			result = t
		})

		return result
	}

	return &OneOfType{uniqueTypes}
}

// Types returns what types the OneOf can be.
// Exposed as ReadonlyTypeSet so caller can't break invariants.
func (oneOf *OneOfType) Types() ReadonlyTypeSet {
	return oneOf.types
}

// References returns any type referenced by the OneOf types
func (oneOf *OneOfType) References() TypeNameSet {
	result := NewTypeNameSet()
	oneOf.types.ForEach(func(t Type, _ int) {
		result = SetUnion(result, t.References())
	})

	return result
}

var oneOfPanicMsg = "OneOfType should have been replaced by generation time by 'convertAllOfAndOneOf' phase"

// AsType always panics; OneOf cannot be represented by the Go AST and must be
// lowered to an object type
func (oneOf OneOfType) AsType(_ *CodeGenerationContext) dst.Expr {
	panic(errors.New(oneOfPanicMsg))
}

// AsDeclarations always panics; OneOf cannot be represented by the Go AST and must be
// lowered to an object type
func (oneOf OneOfType) AsDeclarations(_ *CodeGenerationContext, _ DeclarationContext) []dst.Decl {
	panic(errors.New(oneOfPanicMsg))
}

// AsZero always panics; OneOf cannot be represented by the Go AST and must be
// lowered to an object type
func (oneOf OneOfType) AsZero(definitions TypeDefinitionSet, ctx *CodeGenerationContext) dst.Expr {
	panic(errors.New(oneOfPanicMsg))
}

// RequiredPackageReferences returns the union of the required imports of all the oneOf types
func (oneOf OneOfType) RequiredPackageReferences() *PackageReferenceSet {
	panic(errors.New(oneOfPanicMsg))
}

// Equals returns true if the other Type is a OneOfType that contains
// the same set of types
func (oneOf *OneOfType) Equals(t Type, overrides EqualityOverrides) bool {
	if oneOf == t {
		return true // short-circuit
	}

	other, ok := t.(*OneOfType)
	if !ok {
		return false
	}

	return oneOf.types.Equals(other.types, overrides)
}

// String implements fmt.Stringer
func (oneOf *OneOfType) String() string {
	var subStrings []string
	oneOf.types.ForEach(func(t Type, _ int) {
		subStrings = append(subStrings, t.String())
	})

	sort.Slice(subStrings, func(i, j int) bool {
		return subStrings[i] < subStrings[j]
	})

	return fmt.Sprintf("(oneOf: %s)", strings.Join(subStrings, ", "))
}

// WriteDebugDescription adds a description of the current type to the passed builder
// builder receives the full description, including nested types
// definitions is a dictionary for resolving named types
func (oneOf *OneOfType) WriteDebugDescription(builder *strings.Builder, currentPackage PackageReference) {
	if oneOf == nil {
		builder.WriteString("<nilOneOf>")
		return
	}

	builder.WriteString("OneOf[")
	oneOf.types.ForEach(func(t Type, ix int) {
		if ix > 0 {
			builder.WriteString("|")
		}
		t.WriteDebugDescription(builder, currentPackage)
	})
	builder.WriteString("]")
}

// AsOneOfType unwraps any wrappers around the provided type and returns either the underlying OneOfType and true,
// or nil and false.
func AsOneOfType(t Type) (*OneOfType, bool) {
	if one, ok := t.(*OneOfType); ok {
		return one, true
	}

	if wrapper, ok := t.(MetaType); ok {
		return AsOneOfType(wrapper.Unwrap())
	}

	return nil, false
}
