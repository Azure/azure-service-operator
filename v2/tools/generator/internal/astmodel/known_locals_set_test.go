/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_KnownLocalsSet_CreatesLocal(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	locals := NewKnownLocalsSet(NewIdentifierFactory())

	ident := "myVar"
	g.Expect(locals.CreateLocal(ident)).To(Equal("myVar"))
	g.Expect(locals.CreateLocal(ident)).To(Equal("myVar1"))
	g.Expect(locals.CreateLocal(ident)).To(Equal("myVar2"))
	g.Expect(locals.CreateLocal(ident)).To(Equal("myVar3"))
}

func Test_KnownLocalsSet_CreatesLocalCaseInsensitive(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	locals := NewKnownLocalsSet(NewIdentifierFactory())

	g.Expect(locals.CreateLocal("Student")).To(Equal("student"))
	g.Expect(locals.CreateLocal("Student")).To(Equal("student1"))
	g.Expect(locals.CreateLocal("student")).To(Equal("student2"))
	g.Expect(locals.CreateLocal("student")).To(Equal("student3"))

	g.Expect(locals.CreateLocal("KnownAs")).To(Equal("knownAs"))
}

func Test_KnownLocalsSet_CreatesLocalWithSuffix(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	locals := NewKnownLocalsSet(NewIdentifierFactory())

	ident := "person"
	suffixes := []string{"item", "Map", "List"}
	g.Expect(locals.CreateLocal(ident, suffixes...)).To(Equal("personItem"))
	g.Expect(locals.CreateLocal(ident, suffixes...)).To(Equal("personMap"))
	g.Expect(locals.CreateLocal(ident, suffixes...)).To(Equal("personList"))
	g.Expect(locals.CreateLocal(ident, suffixes...)).To(Equal("personItem1"))
	g.Expect(locals.CreateLocal(ident, suffixes...)).To(Equal("personMap1"))
	g.Expect(locals.CreateLocal(ident, suffixes...)).To(Equal("personList1"))
	g.Expect(locals.CreateLocal(ident, suffixes...)).To(Equal("personItem2"))
	g.Expect(locals.CreateLocal(ident, suffixes...)).To(Equal("personMap2"))
	g.Expect(locals.CreateLocal(ident, suffixes...)).To(Equal("personList2"))
}

func Test_KnownLocalsSet_CreatesLocalWithSuffixAlreadyHasSuffix(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	locals := NewKnownLocalsSet(NewIdentifierFactory())

	ident := "theItem"
	suffixes := []string{"item"}
	g.Expect(locals.CreateLocal(ident, suffixes...)).To(Equal("theItem"))
	g.Expect(locals.CreateLocal(ident, suffixes...)).To(Equal("theItem1"))
	g.Expect(locals.CreateLocal(ident, suffixes...)).To(Equal("theItem2"))
}

func Test_KnownLocalsSet_CreatesLocalWithMultipleSuffixesAlreadyHasSuffix(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	locals := NewKnownLocalsSet(NewIdentifierFactory())

	ident := "theItem"
	suffixes := []string{"item", "element"}
	// TODO: Is this the behavior we want?
	g.Expect(locals.CreateLocal(ident, suffixes...)).To(Equal("theItem"))
	g.Expect(locals.CreateLocal(ident, suffixes...)).To(Equal("theItemElement"))
	g.Expect(locals.CreateLocal(ident, suffixes...)).To(Equal("theItem1"))
	g.Expect(locals.CreateLocal(ident, suffixes...)).To(Equal("theItemElement1"))
}

func Test_KnownLocalsSet_CreateSingularLocal_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	locals := NewKnownLocalsSet(NewIdentifierFactory())

	g.Expect(locals.CreateSingularLocal("people")).To(Equal("person"))
	g.Expect(locals.CreateSingularLocal("people")).To(Equal("person1"))

	g.Expect(locals.CreateSingularLocal("knownAs")).To(Equal("knownAs"))
	g.Expect(locals.CreateSingularLocal("knownAs")).To(Equal("knownAs1"))
}

func Test_KnownLocalsSet_CreatePluralLocal_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	locals := NewKnownLocalsSet(NewIdentifierFactory())

	g.Expect(locals.CreatePluralLocal("person")).To(Equal("people"))
	g.Expect(locals.CreatePluralLocal("person")).To(Equal("people1"))

	g.Expect(locals.CreatePluralLocal("knownAs")).To(Equal("knownAs"))
	g.Expect(locals.CreatePluralLocal("knownAs")).To(Equal("knownAs1"))
}
