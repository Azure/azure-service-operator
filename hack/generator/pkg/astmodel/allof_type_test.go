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
	result := MakeAllOfType(oneType)

	g.Expect(result).To(BeIdenticalTo(oneType))
}

func TestAllOfIdenticalTypesReturnsThatType(t *testing.T) {
	g := NewGomegaWithT(t)

	oneType := StringType
	result := MakeAllOfType(oneType, oneType)

	g.Expect(result).To(BeIdenticalTo(oneType))
}

func TestAllOfFlattensNestedAllOfs(t *testing.T) {
	g := NewGomegaWithT(t)

	result := MakeAllOfType(BoolType, MakeAllOfType(StringType, IntType))

	expected := MakeAllOfType(BoolType, StringType, IntType)

	g.Expect(result).To(Equal(expected))
}

func TestAllOfOneOfBecomesOneOfAllOf(t *testing.T) {
	g := NewGomegaWithT(t)

	oneOf := MakeOneOfType(StringType, FloatType)
	result := MakeAllOfType(BoolType, IntType, oneOf)

	expected := MakeOneOfType(
		MakeAllOfType(BoolType, IntType, StringType),
		MakeAllOfType(BoolType, IntType, FloatType),
	)

	g.Expect(result).To(Equal(expected))
}

func TestAllOfEqualityDoesNotCareAboutOrder(t *testing.T) {
	g := NewGomegaWithT(t)

	x := MakeAllOfType(StringType, BoolType)
	y := MakeAllOfType(BoolType, StringType)

	g.Expect(x.Equals(y)).To(BeTrue())
	g.Expect(y.Equals(x)).To(BeTrue())
}

func TestAllOfMustHaveAllTypesToBeEqual(t *testing.T) {
	g := NewGomegaWithT(t)

	x := MakeAllOfType(StringType, BoolType, FloatType)
	y := MakeAllOfType(BoolType, StringType)

	g.Expect(x.Equals(y)).To(BeFalse())
	g.Expect(y.Equals(x)).To(BeFalse())
}

func TestAllOfsWithDifferentTypesAreNotEqual(t *testing.T) {
	g := NewGomegaWithT(t)

	x := MakeAllOfType(StringType, FloatType)
	y := MakeAllOfType(BoolType, StringType)

	g.Expect(x.Equals(y)).To(BeFalse())
	g.Expect(y.Equals(x)).To(BeFalse())
}

func TestAllOfAsTypePanics(t *testing.T) {
	g := NewGomegaWithT(t)

	x := AllOfType{}
	g.Expect(func() {
		x.AsType(&CodeGenerationContext{})
	}).To(PanicWith("should have been replaced by generation time by 'convertAllOfAndOneOf' phase"))
}

func TestAllOfAsDeclarationsPanics(t *testing.T) {
	g := NewGomegaWithT(t)

	x := AllOfType{}
	g.Expect(func() {
		x.AsDeclarations(&CodeGenerationContext{}, TypeName{}, []string{})
	}).To(PanicWith("should have been replaced by generation time by 'convertAllOfAndOneOf' phase"))
}

func TestAllOfRequiredImportsPanics(t *testing.T) {
	g := NewGomegaWithT(t)

	x := AllOfType{}
	g.Expect(func() {
		x.RequiredPackageReferences()
	}).To(PanicWith("should have been replaced by generation time by 'convertAllOfAndOneOf' phase"))
}
