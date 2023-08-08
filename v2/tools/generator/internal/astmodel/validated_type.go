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
)

type ArrayValidations struct {
	MaxItems    *int64
	MinItems    *int64
	UniqueItems bool
	/*
		maxContains *int
		minContains *int
	*/
}

func (av ArrayValidations) Equals(other Validations) bool {
	o, ok := other.(ArrayValidations)
	if !ok {
		return false
	}

	return equalOptionalInt64s(av.MaxItems, o.MaxItems) &&
		equalOptionalInt64s(av.MinItems, o.MinItems) &&
		av.UniqueItems == o.UniqueItems
}

func (av ArrayValidations) ToKubeBuilderValidations() []KubeBuilderValidation {
	var result []KubeBuilderValidation
	if av.MaxItems != nil {
		result = append(result, MakeMaxItemsValidation(*av.MaxItems))
	}

	if av.MinItems != nil {
		result = append(result, MakeMinItemsValidation(*av.MinItems))
	}

	if av.UniqueItems {
		result = append(result, MakeUniqueItemsValidation())
	}

	return result
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

type Validations interface {
	Equals(other Validations) bool
	ToKubeBuilderValidations() []KubeBuilderValidation
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

func (v *ValidatedType) AsDeclarations(c *CodeGenerationContext, declContext DeclarationContext) []dst.Decl {
	declContext.Validations = append(declContext.Validations, v.validations.ToKubeBuilderValidations()...)
	return v.ElementType().AsDeclarations(c, declContext)
}

// AsType panics because validated types should always be named
func (v *ValidatedType) AsType(_ *CodeGenerationContext) dst.Expr {
	panic("Should never happen: validated types must either be named (handled by 'name types for CRDs' pipeline stage) or be directly under properties (handled by PropertyDefinition.AsField)")
}

// AsZero returns the zero for our underlying type
func (v *ValidatedType) AsZero(definitions TypeDefinitionSet, ctx *CodeGenerationContext) dst.Expr {
	return v.element.AsZero(definitions, ctx)
}

func (v *ValidatedType) References() TypeNameSet[TypeName] {
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
func (v ValidatedType) Unwrap() Type {
	return v.element
}

// WriteDebugDescription adds a description of the current type to the passed builder
// builder receives the full description, including nested types
// definitions is a dictionary for resolving named types
func (v ValidatedType) WriteDebugDescription(builder *strings.Builder, currentPackage PackageReference) {
	builder.WriteString("Validated[")
	if v.element != nil {
		v.element.WriteDebugDescription(builder, currentPackage)
	} else {
		builder.WriteString("<nilType>")
	}

	builder.WriteString("]")
}
