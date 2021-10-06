/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestOneOfOneTypeReturnsThatType(t *testing.T) {
	g := NewGomegaWithT(t)

	oneType := StringType
	result := BuildOneOfType(oneType)

	g.Expect(result).To(BeIdenticalTo(oneType))
}

func TestOneOfIdenticalTypesReturnsThatType(t *testing.T) {
	g := NewGomegaWithT(t)

	oneType := StringType
	result := BuildOneOfType(oneType, oneType)

	g.Expect(result).To(BeIdenticalTo(oneType))
}

func TestOneOfFlattensNestedOneOfs(t *testing.T) {
	g := NewGomegaWithT(t)

	result := BuildOneOfType(BoolType, BuildOneOfType(StringType, IntType))

	expected := BuildOneOfType(BoolType, StringType, IntType)

	g.Expect(result).To(Equal(expected))
}

func TestOneOfEqualityDoesNotCareAboutOrder(t *testing.T) {
	g := NewGomegaWithT(t)

	x := BuildOneOfType(StringType, BoolType)
	y := BuildOneOfType(BoolType, StringType)

	g.Expect(x.Equals(y)).To(BeTrue())
	g.Expect(y.Equals(x)).To(BeTrue())
}

func TestOneOfMustHaveAllTypesToBeEqual(t *testing.T) {
	g := NewGomegaWithT(t)

	x := BuildOneOfType(StringType, BoolType, FloatType)
	y := BuildOneOfType(BoolType, StringType)

	g.Expect(x.Equals(y)).To(BeFalse())
	g.Expect(y.Equals(x)).To(BeFalse())
}

func TestOneOfsWithDifferentTypesAreNotEqual(t *testing.T) {
	g := NewGomegaWithT(t)

	x := BuildOneOfType(StringType, FloatType)
	y := BuildOneOfType(BoolType, StringType)

	g.Expect(x.Equals(y)).To(BeFalse())
	g.Expect(y.Equals(x)).To(BeFalse())
}

var expectedOneOfPanic = "OneOfType should have been replaced by generation time by 'convertAllOfAndOneOf' phase"

func TestOneOfAsTypePanics(t *testing.T) {
	g := NewGomegaWithT(t)

	x := OneOfType{}
	g.Expect(func() {
		x.AsType(&CodeGenerationContext{})
	}).To(PanicWith(MatchError(expectedOneOfPanic)))
}

func TestOneOfAsDeclarationsPanics(t *testing.T) {
	g := NewGomegaWithT(t)

	x := OneOfType{}
	g.Expect(func() {
		x.AsDeclarations(&CodeGenerationContext{}, DeclarationContext{})
	}).To(PanicWith(MatchError(expectedOneOfPanic)))
}

func TestOneOfRequiredImportsPanics(t *testing.T) {
	g := NewGomegaWithT(t)

	x := OneOfType{}
	g.Expect(func() {
		x.RequiredPackageReferences()
	}).To(PanicWith(MatchError(expectedOneOfPanic)))
}
