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
	"k8s.io/klog/v2"
)

// AllOfType represents something that is the union
// of all the given types
type AllOfType struct {
	// invariants:
	// - all types are unique (enforced by TypeSet)
	// - length > 1
	// - no nested AllOfs (aside from indirectly via TypeName)
	// - we also transform an allOf with a single oneOf inside
	//   to a oneOf with allOfs inside (see below)
	types TypeSet
}

var _ Type = &AllOfType{}

func NewAllOfType(types ...Type) *AllOfType {
	return &AllOfType{
		types: MakeTypeSet(types...),
	}
}

// BuildAllOfType is a smart constructor for AllOfType,
// maintaining the invariants. If only one unique type
// is passed, then the result will be that type, not an AllOf.
func BuildAllOfType(types ...Type) Type {
	uniqueTypes := MakeTypeSet()
	for _, t := range types {
		if allOf, ok := t.(*AllOfType); ok {
			allOf.types.ForEach(func(t Type, _ int) {
				uniqueTypes.Add(t)
			})
		} else {
			uniqueTypes.Add(t)
		}
	}

	if t, ok := uniqueTypes.Single(); ok {
		return t
	}

	// see if there are any OneOfs inside
	var oneOfs []*OneOfType
	var notOneOfs []Type
	uniqueTypes.ForEach(func(t Type, _ int) {
		if oneOf, ok := t.(*OneOfType); ok {
			oneOfs = append(oneOfs, oneOf)
		} else {
			notOneOfs = append(notOneOfs, t)
		}
	})

	if len(oneOfs) == 1 {
		// we want to push AllOf down so that:
		// 		allOf { x, y, oneOf { a, b } }
		// becomes
		//		oneOf { allOf { x, y, a }, allOf { x, y, b } }
		// the latter is much easier to deal with in code
		// as we can deal with each case separately instead of mixing
		// "outer" and "inner" properties

		var ts []Type
		onlyOneOf := oneOfs[0]
		onlyOneOf.types.ForEach(func(t Type, _ int) {
			ts = append(ts, BuildAllOfType(append(notOneOfs, t)...))
		})

		return onlyOneOf.WithTypes(ts)
	} else if len(oneOfs) > 1 {
		// emit a warning if this ever comes up
		// (it doesn't at the moment)
		klog.Warningf("More than one oneOf inside allOf")
	}

	// 0 oneOf (nothing to do) or >1 oneOf (too hard)
	return &AllOfType{uniqueTypes}
}

// Types returns what types the AllOf can be.
// Exposed as ReadonlyTypeSet so caller can't break invariants.
func (allOf *AllOfType) Types() ReadonlyTypeSet {
	return allOf.types
}

// References returns any type referenced by the AllOf types
func (allOf *AllOfType) References() TypeNameSet {
	result := NewTypeNameSet()
	allOf.types.ForEach(func(t Type, _ int) {
		result = SetUnion(result, t.References())
	})

	return result
}

var allOfPanicMsg = "AllOfType should have been replaced by generation time by 'convertAllOfAndOneOf' phase"

// AsType always panics; AllOf cannot be represented by the Go AST and must be
// lowered to an object type
func (allOf *AllOfType) AsType(_ *CodeGenerationContext) dst.Expr {
	panic(errors.New(allOfPanicMsg))
}

// AsDeclarations always panics; AllOf cannot be represented by the Go AST and must be
// lowered to an object type
func (allOf *AllOfType) AsDeclarations(_ *CodeGenerationContext, _ DeclarationContext) []dst.Decl {
	panic(errors.New(allOfPanicMsg))
}

// AsZero always panics; AllOf cannot be represented by the Go AST and must be
// lowered to an object type
func (allOf *AllOfType) AsZero(_ TypeDefinitionSet, _ *CodeGenerationContext) dst.Expr {
	panic(errors.New(allOfPanicMsg))
}

// RequiredPackageReferences always panics; AllOf cannot be represented by the Go AST and must be
// lowered to an object type
func (allOf *AllOfType) RequiredPackageReferences() *PackageReferenceSet {
	panic(errors.New(allOfPanicMsg))
}

// Equals returns true if the other Type is a AllOf that contains
// the same set of types
func (allOf *AllOfType) Equals(t Type, overrides EqualityOverrides) bool {
	if allOf == t {
		return true // short-circuit
	}

	other, ok := t.(*AllOfType)
	if !ok {
		return false
	}

	return allOf.types.Equals(other.types, overrides)
}

// String implements fmt.Stringer
func (allOf *AllOfType) String() string {
	var subStrings []string
	allOf.types.ForEach(func(t Type, _ int) {
		subStrings = append(subStrings, t.String())
	})

	sort.Slice(subStrings, func(i, j int) bool {
		return subStrings[i] < subStrings[j]
	})

	return fmt.Sprintf("(allOf: %s)", strings.Join(subStrings, ", "))
}

// WriteDebugDescription adds a description of the current AnyOf type to the passed builder.
// builder receives the full description, including nested types.
// definitions is a dictionary for resolving named types.
func (allOf *AllOfType) WriteDebugDescription(builder *strings.Builder, currentPackage PackageReference) {
	if allOf == nil {
		builder.WriteString("<nilAllOf>")
		return
	}

	builder.WriteString("AllOf[")
	allOf.types.ForEach(func(t Type, ix int) {
		if ix > 0 {
			builder.WriteString("|")
		}
		t.WriteDebugDescription(builder, currentPackage)
	})
	builder.WriteString("]")
}

// AsAllOfType unwraps any wrappers around the provided type and returns either the underlying AllOfType and true,
// or nil and false.
func AsAllOfType(t Type) (*AllOfType, bool) {
	if all, ok := t.(*AllOfType); ok {
		return all, true
	}

	if wrapper, ok := t.(MetaType); ok {
		return AsAllOfType(wrapper.Unwrap())
	}

	return nil, false
}
