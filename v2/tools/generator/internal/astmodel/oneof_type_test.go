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
	t.Parallel()
	g := NewGomegaWithT(t)

	oneType := StringType
	result := BuildOneOfType(oneType)

	g.Expect(result).To(BeIdenticalTo(oneType))
}

func TestOneOfIdenticalTypesReturnsThatType(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	oneType := StringType
	result := BuildOneOfType(oneType, oneType)

	g.Expect(result).To(BeIdenticalTo(oneType))
}

func TestOneOfFlattensNestedOneOfs(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	result := BuildOneOfType(BoolType, BuildOneOfType(StringType, IntType))

	expected := BuildOneOfType(BoolType, StringType, IntType)

	g.Expect(result).To(Equal(expected))
}

func TestOneOfEqualityDoesNotCareAboutOrder(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	x := BuildOneOfType(StringType, BoolType)
	y := BuildOneOfType(BoolType, StringType)

	g.Expect(TypeEquals(x, y)).To(BeTrue())
	g.Expect(TypeEquals(y, x)).To(BeTrue())
}

func TestOneOfMustHaveAllTypesToBeEqual(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	x := BuildOneOfType(StringType, BoolType, FloatType)
	y := BuildOneOfType(BoolType, StringType)

	g.Expect(TypeEquals(x, y)).To(BeFalse())
	g.Expect(TypeEquals(y, x)).To(BeFalse())
}

func TestOneOfsWithDifferentTypesAreNotEqual(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	x := BuildOneOfType(StringType, FloatType)
	y := BuildOneOfType(BoolType, StringType)

	g.Expect(TypeEquals(x, y)).To(BeFalse())
	g.Expect(TypeEquals(y, x)).To(BeFalse())
}

var expectedOneOfPanic = "OneOfType should have been replaced by generation time by 'convertAllOfAndOneOf' phase"

func TestOneOfAsTypePanics(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	x := OneOfType{}
	g.Expect(func() {
		x.AsType(&CodeGenerationContext{})
	}).To(PanicWith(MatchError(expectedOneOfPanic)))
}

func TestOneOfAsDeclarationsPanics(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	x := OneOfType{}
	g.Expect(func() {
		x.AsDeclarations(&CodeGenerationContext{}, DeclarationContext{})
	}).To(PanicWith(MatchError(expectedOneOfPanic)))
}

func TestOneOfRequiredImportsPanics(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	x := OneOfType{}
	g.Expect(func() {
		x.RequiredPackageReferences()
	}).To(PanicWith(MatchError(expectedOneOfPanic)))
}
