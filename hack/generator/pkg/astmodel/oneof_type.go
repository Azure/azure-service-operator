/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"sort"
	"strings"

	ast "github.com/dave/dst"
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

// MakeOneOfType is a smart constructor for a  OneOfType,
// maintaining the invariants. If only one unique type
// is passed, the result will be that type, not a OneOf.
func MakeOneOfType(types ...Type) Type {
	uniqueTypes := MakeTypeSet()
	for _, t := range types {
		if oneOf, ok := t.(OneOfType); ok {
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

	return OneOfType{uniqueTypes}
}

var _ Type = OneOfType{}

// Types returns what types the OneOf can be.
// Exposed as ReadonlyTypeSet so caller can't break invariants.
func (oneOf OneOfType) Types() ReadonlyTypeSet {
	return oneOf.types
}

// References returns any type referenced by the OneOf types
func (oneOf OneOfType) References() TypeNameSet {
	var result TypeNameSet

	oneOf.types.ForEach(func(t Type, _ int) {
		result = SetUnion(result, t.References())
	})

	return result
}

// AsType always panics; AllOf cannot be represented by the Go AST and must be
// lowered to an object type
func (oneOf OneOfType) AsType(_ *CodeGenerationContext) ast.Expr {
	panic("should have been replaced by generation time by 'convertAllOfAndOneOf' phase")
}

// AsDeclarations always panics; AllOf cannot be represented by the Go AST and must be
// lowered to an object type
func (oneOf OneOfType) AsDeclarations(_ *CodeGenerationContext, _ TypeName, _ []string) []ast.Decl {
	panic("should have been replaced by generation time by 'convertAllOfAndOneOf' phase")
}

// RequiredPackageReferences returns the union of the required imports of all the oneOf types
func (oneOf OneOfType) RequiredPackageReferences() *PackageReferenceSet {
	panic("should have been replaced by generation time by 'convertAllOfAndOneOf' phase")
}

// Equals returns true if the other Type is a OneOfType that contains
// the same set of types
func (oneOf OneOfType) Equals(t Type) bool {
	other, ok := t.(OneOfType)
	if !ok {
		return false
	}

	return oneOf.types.Equals(other.types)
}

// String implements fmt.Stringer
func (oneOf OneOfType) String() string {
	var subStrings []string
	oneOf.types.ForEach(func(t Type, _ int) {
		subStrings = append(subStrings, t.String())
	})

	sort.Slice(subStrings, func(i, j int) bool {
		return subStrings[i] < subStrings[j]
	})

	return fmt.Sprintf("(oneOf: %s)", strings.Join(subStrings, ", "))
}
