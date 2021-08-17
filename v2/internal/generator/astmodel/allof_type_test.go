/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestAllOfOneTypeReturnsThatType(t *testing.T) {
	g := NewGomegaWithT(t)

	oneType := StringType
	result := BuildAllOfType(oneType)

	g.Expect(result).To(BeIdenticalTo(oneType))
}

func TestAllOfIdenticalTypesReturnsThatType(t *testing.T) {
	g := NewGomegaWithT(t)

	oneType := StringType
	result := BuildAllOfType(oneType, oneType)

	g.Expect(result).To(BeIdenticalTo(oneType))
}

func TestAllOfFlattensNestedAllOfs(t *testing.T) {
	g := NewGomegaWithT(t)

	result := BuildAllOfType(BoolType, BuildAllOfType(StringType, IntType))

	expected := BuildAllOfType(BoolType, StringType, IntType)

	g.Expect(result).To(Equal(expected))
}

func TestAllOfOneOfBecomesOneOfAllOf(t *testing.T) {
	g := NewGomegaWithT(t)

	oneOf := BuildOneOfType(StringType, FloatType)
	result := BuildAllOfType(BoolType, IntType, oneOf)

	expected := BuildOneOfType(
		BuildAllOfType(BoolType, IntType, StringType),
		BuildAllOfType(BoolType, IntType, FloatType),
	)

	g.Expect(result).To(Equal(expected))
}

func TestAllOfEqualityDoesNotCareAboutOrder(t *testing.T) {
	g := NewGomegaWithT(t)

	x := BuildAllOfType(StringType, BoolType)
	y := BuildAllOfType(BoolType, StringType)

	g.Expect(x.Equals(y)).To(BeTrue())
	g.Expect(y.Equals(x)).To(BeTrue())
}

func TestAllOfMustHaveAllTypesToBeEqual(t *testing.T) {
	g := NewGomegaWithT(t)

	x := BuildAllOfType(StringType, BoolType, FloatType)
	y := BuildAllOfType(BoolType, StringType)

	g.Expect(x.Equals(y)).To(BeFalse())
	g.Expect(y.Equals(x)).To(BeFalse())
}

func TestAllOfsWithDifferentTypesAreNotEqual(t *testing.T) {
	g := NewGomegaWithT(t)

	x := BuildAllOfType(StringType, FloatType)
	y := BuildAllOfType(BoolType, StringType)

	g.Expect(x.Equals(y)).To(BeFalse())
	g.Expect(y.Equals(x)).To(BeFalse())
}

var expectedAllOfPanic = "AllOfType should have been replaced by generation time by 'convertAllOfAndOneOf' phase"

func TestAllOfAsTypePanics(t *testing.T) {
	g := NewGomegaWithT(t)

	x := AllOfType{}
	g.Expect(func() {
		x.AsType(&CodeGenerationContext{})
	}).To(PanicWith(MatchError(expectedAllOfPanic)))
}

func TestAllOfAsDeclarationsPanics(t *testing.T) {
	g := NewGomegaWithT(t)

	x := AllOfType{}
	g.Expect(func() {
		x.AsDeclarations(&CodeGenerationContext{}, DeclarationContext{})
	}).To(PanicWith(MatchError(expectedAllOfPanic)))
}

func TestAllOfRequiredImportsPanics(t *testing.T) {
	g := NewGomegaWithT(t)

	x := AllOfType{}
	g.Expect(func() {
		x.RequiredPackageReferences()
	}).To(PanicWith(MatchError(expectedAllOfPanic)))
}
