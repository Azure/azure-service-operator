/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestOneOfEqualityDoesNotCareAboutOrder(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	x := NewOneOfType("x", StringType, BoolType)
	y := NewOneOfType("y", BoolType, StringType)

	g.Expect(TypeEquals(x, y)).To(BeTrue())
	g.Expect(TypeEquals(y, x)).To(BeTrue())
}

func TestOneOfMustHaveAllTypesToBeEqual(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	x := NewOneOfType("x", StringType, BoolType, FloatType)
	y := NewOneOfType("y", BoolType, StringType)

	g.Expect(TypeEquals(x, y)).To(BeFalse())
	g.Expect(TypeEquals(y, x)).To(BeFalse())
}

func TestOneOfsWithDifferentTypesAreNotEqual(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	x := NewOneOfType("x", StringType, FloatType)
	y := NewOneOfType("y", BoolType, StringType)

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
