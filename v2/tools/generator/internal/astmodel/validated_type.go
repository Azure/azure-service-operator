/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"

	"github.com/dave/dst"
	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/internal/set"
)

type ArrayValidations struct {
	MaxItems *int64
	MinItems *int64
}

func (av ArrayValidations) Equals(other Validations) bool {
	o, ok := other.(ArrayValidations)
	if !ok {
		return false
	}

	return equalOptionalInt64s(av.MaxItems, o.MaxItems) &&
		equalOptionalInt64s(av.MinItems, o.MinItems)
}

func (av ArrayValidations) ToKubeBuilderValidations() []KubeBuilderValidation {
	var result []KubeBuilderValidation
	if av.MaxItems != nil {
		result = append(result, MakeMaxItemsValidation(*av.MaxItems))
	}

	if av.MinItems != nil {
		result = append(result, MakeMinItemsValidation(*av.MinItems))
	}

	return result
}

func (av ArrayValidations) MergeWith(other Validations) (Validations, error) {
	o, ok := other.(ArrayValidations)
	if !ok {
		return nil, eris.New(fmt.Sprintf("cannot merge different validation types %T and %T", av, other))
	}

	result := ArrayValidations{}

	// Merge MaxItems by taking the more restrictive (smaller) value
	result.MaxItems = min(av.MaxItems, o.MaxItems)
	// Merge MinItems by taking the more restrictive (larger) value
	result.MinItems = max(av.MinItems, o.MinItems)
	return result, nil
}

type StringValidations struct {
	MaxLength *int64
	MinLength *int64
	Patterns  []*regexp.Regexp
}

func (sv StringValidations) Equals(other Validations) bool {
	o, ok := other.(StringValidations)
	if !ok {
		return false
	}

	return equalOptionalInt64s(sv.MaxLength, o.MaxLength) &&
		equalOptionalInt64s(sv.MinLength, o.MinLength) &&
		equalRegexpSlices(sv.Patterns, o.Patterns)
}

func (sv StringValidations) ToKubeBuilderValidations() []KubeBuilderValidation {
	result := make([]KubeBuilderValidation, 0, len(sv.Patterns)+2)
	if sv.MaxLength != nil {
		result = append(result, MakeMaxLengthValidation(*sv.MaxLength))
	}

	if sv.MinLength != nil {
		result = append(result, MakeMinLengthValidation(*sv.MinLength))
	}

	for _, pattern := range sv.Patterns {
		result = append(result, MakePatternValidation(pattern))
	}

	return result
}

func (sv StringValidations) MergeWith(other Validations) (Validations, error) {
	o, ok := other.(StringValidations)
	if !ok {
		return nil, eris.New(fmt.Sprintf("cannot merge different validation types %T and %T", sv, other))
	}

	result := StringValidations{}

	// Merge MaxLength by taking the more restrictive (smaller) value
	result.MaxLength = min(sv.MaxLength, o.MaxLength)
	// Merge MinLength by taking the more restrictive (larger) value
	result.MinLength = max(sv.MinLength, o.MinLength)
	// Merge Patterns by concatenating the slices, keeping only unique patterns
	// TODO: It may be possible for two conflicting patterns to make something impossible to set... ideally we would fix that upstream in the REST API specification though
	result.Patterns = appendUniquePatterns(sv.Patterns, o.Patterns)
	return result, nil
}

type NumberValidations struct {
	// TODO: update to use doubles once newer version of controller-gen is released
	Maximum          *big.Rat
	Minimum          *big.Rat
	ExclusiveMaximum bool
	ExclusiveMinimum bool
	MultipleOf       *big.Rat
}

func (nv NumberValidations) Equals(other Validations) bool {
	o, ok := other.(NumberValidations)
	if !ok {
		return false
	}

	return nv.ExclusiveMaximum == o.ExclusiveMaximum &&
		nv.ExclusiveMinimum == o.ExclusiveMinimum &&
		equalOptionalBigRats(nv.Maximum, o.Maximum) &&
		equalOptionalBigRats(nv.Minimum, o.Minimum) &&
		equalOptionalBigRats(nv.MultipleOf, o.MultipleOf)
}

func (nv NumberValidations) ToKubeBuilderValidations() []KubeBuilderValidation {
	var result []KubeBuilderValidation

	if nv.Maximum != nil {
		result = append(result, MakeMaximumValidation(nv.Maximum))
	}

	if nv.ExclusiveMaximum {
		result = append(result, MakeExclusiveMaxiumValidation())
	}

	if nv.Minimum != nil {
		result = append(result, MaxMinimumValidation(nv.Minimum))
	}

	if nv.ExclusiveMinimum {
		result = append(result, MakeExclusiveMinimumValidation())
	}

	if nv.MultipleOf != nil {
		result = append(result, MakeMultipleOfValidation(nv.MultipleOf))
	}

	return result
}

func (nv NumberValidations) MergeWith(other Validations) (Validations, error) {
	o, ok := other.(NumberValidations)
	if !ok {
		return nil, eris.New(fmt.Sprintf("cannot merge different validation types %T and %T", nv, other))
	}

	result := NumberValidations{}

	// Merge Maximum by taking the more restrictive (smaller) value
	result.Maximum = minRat(nv.Maximum, o.Maximum)
	// Merge Minimum by taking the more restrictive (larger) value
	result.Minimum = maxRat(nv.Minimum, o.Minimum)
	// Merge ExclusiveMaximum by ||ing the two values
	result.ExclusiveMaximum = nv.ExclusiveMaximum || o.ExclusiveMaximum
	// Merge ExclusiveMinimum by ||ing the two values
	result.ExclusiveMinimum = nv.ExclusiveMinimum || o.ExclusiveMinimum
	// Merge MultipleOf by taking the more restrictive (larger) value
	// TODO: Possibly we should use the GCD of the two values?
	result.MultipleOf = maxRat(nv.MultipleOf, o.MultipleOf)
	return result, nil
}

type Validations interface {
	Equals(other Validations) bool
	ToKubeBuilderValidations() []KubeBuilderValidation
	MergeWith(other Validations) (Validations, error)
}

// ValidatedType is used for schema validation attributes
type ValidatedType struct {
	validations Validations
	element     Type
}

var _ Type = &ValidatedType{}

var _ MetaType = &ValidatedType{}

func NewValidatedType(element Type, validations Validations) *ValidatedType {
	return &ValidatedType{element: element, validations: validations}
}

func (v *ValidatedType) ElementType() Type {
	return v.element
}

func (v *ValidatedType) WithElement(t Type) Type {
	if TypeEquals(v.element, t) {
		return v
	}

	result := *v
	result.element = t
	return &result
}

func (v *ValidatedType) Validations() Validations {
	return v.validations
}

func (v *ValidatedType) WithType(newElement Type) *ValidatedType {
	result := *v
	result.element = newElement
	return &result
}

func (v *ValidatedType) AsDeclarations(
	codeGenerationContext *CodeGenerationContext,
	declContext DeclarationContext,
) ([]dst.Decl, error) {
	declContext.Validations = append(declContext.Validations, v.validations.ToKubeBuilderValidations()...)
	return v.ElementType().AsDeclarations(codeGenerationContext, declContext)
}

// AsType panics because validated types should always be named
func (v *ValidatedType) AsTypeExpr(codeGenerationContext *CodeGenerationContext) (dst.Expr, error) {
	panic("Should never happen: validated types must either be named (handled by 'name types for CRDs' pipeline stage) or be directly under properties (handled by PropertyDefinition.AsField)")
}

// AsZero returns the zero for our underlying type
func (v *ValidatedType) AsZero(definitions TypeDefinitionSet, ctx *CodeGenerationContext) dst.Expr {
	return v.element.AsZero(definitions, ctx)
}

func (v *ValidatedType) References() TypeNameSet {
	return v.element.References()
}

func (v *ValidatedType) RequiredPackageReferences() *PackageReferenceSet {
	return v.element.RequiredPackageReferences()
}

func (v *ValidatedType) String() string {
	return fmt.Sprintf("Validated(%s)", v.element.String())
}

func (v *ValidatedType) Equals(t Type, overrides EqualityOverrides) bool {
	if v == t {
		return true
	}

	other, ok := t.(*ValidatedType)
	if !ok {
		return false
	}

	return v.validations.Equals(other.validations) &&
		v.element.Equals(other.element, overrides)
}

func equalOptionalInt64s(left *int64, right *int64) bool {
	if left != nil {
		return right != nil && *left == *right
	}

	return right == nil
}

func equalOptionalBigRats(left *big.Rat, right *big.Rat) bool {
	if left != nil {
		return right != nil && left.Cmp(right) == 0
	}

	return right == nil
}

func equalRegexpSlices(left []*regexp.Regexp, right []*regexp.Regexp) bool {
	if len(left) != len(right) {
		return false
	}

	for ix := range left {
		if left[ix] != right[ix] &&
			left[ix].String() != right[ix].String() {
			return false
		}
	}

	return true
}

// Unwrap returns the type contained within the validated type
func (v *ValidatedType) Unwrap() Type {
	return v.element
}

// WriteDebugDescription adds a description of the current type to the passed builder
// builder receives the full description, including nested types
// definitions is a dictionary for resolving named types
func (v *ValidatedType) WriteDebugDescription(builder *strings.Builder, currentPackage InternalPackageReference) {
	builder.WriteString("Validated[")
	if v.element != nil {
		v.element.WriteDebugDescription(builder, currentPackage)
	} else {
		builder.WriteString("<nilType>")
	}

	builder.WriteString("]")
}

func min(a, b *int64) *int64 {
	if a != nil && b != nil {
		if *a < *b {
			return a
		} else {
			return b
		}
	} else if a != nil {
		return a
	} else if b != nil {
		return b
	}

	return nil
}

func max(a, b *int64) *int64 {
	if a != nil && b != nil {
		if *a > *b {
			return a
		} else {
			return b
		}
	} else if a != nil {
		return a
	} else if b != nil {
		return b
	}

	return nil
}

func minRat(a, b *big.Rat) *big.Rat {
	if a != nil && b != nil {
		if a.Cmp(b) < 0 {
			return a
		} else {
			return b
		}
	} else if a != nil {
		return a
	} else if b != nil {
		return b
	}

	return nil
}

func maxRat(a, b *big.Rat) *big.Rat {
	if a != nil && b != nil {
		if a.Cmp(b) > 0 {
			return a
		} else {
			return b
		}
	} else if a != nil {
		return a
	} else if b != nil {
		return b
	}

	return nil
}

// appendUniquePatterns combines two pattern slices, keeping only unique patterns
// based on their string representation.
func appendUniquePatterns(existing, new []*regexp.Regexp) []*regexp.Regexp {
	if len(new) == 0 {
		return existing
	}
	if len(existing) == 0 {
		return new
	}

	// Build a set of existing pattern strings
	seen := set.Make[string]()
	for _, p := range existing {
		seen.Add(p.String())
	}

	// Start with existing patterns
	result := existing
	for _, p := range new {
		patternStr := p.String()
		if seen.Contains(patternStr) {
			continue
		}
		seen.Add(patternStr)
		result = append(result, p)
	}

	return result
}
