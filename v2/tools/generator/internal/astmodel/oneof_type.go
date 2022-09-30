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

// OneOfType represents something that can be any one of a number of selected types
// We have three forms:
// Completed - one that has been fully constructed and is ready to use
// Base - defines the base type and has a list of valid discriminator values
// Option - defines one of the potential types and has a discriminator value and a name to identify the relevant Base

type OneOfType struct {
	kind                  OneOfKind           // What kind of OneOf do we have?
	name                  string              // Name of the OneOf
	discriminatorProperty *PropertyDefinition // Defines the discriminator
	types                 TypeSet             // Set of all possible types
	discriminatorValue    string              // Value used to identify this option
	baseName              string              // Name of the expected base type
}

var _ Type = &OneOfType{}

// BuildOneOfType is a smart constructor for a OneOfType,
// maintaining the invariants. If only one unique type
// is passed, the result will be that type, not a OneOf.
func BuildOneOfType(types ...Type) Type {
	uniqueTypes := flattenTypesForOneOf(types)
	if t, ok := uniqueTypes.Only(); ok {
		return t
	}

	return &OneOfType{
		types: uniqueTypes,
	}
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
