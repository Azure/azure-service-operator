/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

var (
	pkg                = MakeExternalPackageReference("foo")
	alphaDefinition    = createTestDefinition("alpha", StringType)
	betaDefinition     = createTestDefinition("beta", StringType)
	gammaDefinition    = createTestDefinition("gamma", StringType)
	deltaDefinition    = createTestDefinition("delta", StringType)
	deltaIntDefinition = createTestDefinition("delta", IntType)
)

/*
 * Add() Tests
 */

func Test_TypesAdd_GivenType_ModifiesSet(t *testing.T) {
	g := NewGomegaWithT(t)

	types := make(Types)
	types.Add(alphaDefinition)

	g.Expect(types).To(ContainElement(alphaDefinition))
}

func Test_TypesAdd_GivenTypeAlreadyPresent_Panics(t *testing.T) {
	g := NewGomegaWithT(t)

	types := make(Types)
	types.Add(alphaDefinition)

	g.Expect(func() { types.Add(alphaDefinition) }).To(Panic())
}

/*
 * AddAll() Tests
 */

func Test_TypesAddAll_GivenTypes_ModifiesSet(t *testing.T) {
	g := NewGomegaWithT(t)

	types := createTestTypes(alphaDefinition, betaDefinition)
	types.AddAll(gammaDefinition, deltaDefinition)

	g.Expect(types).To(ContainElements(gammaDefinition, deltaDefinition))
}

func Test_TypesAddAll_GivenOverlappingTypes_Panics(t *testing.T) {
	g := NewGomegaWithT(t)

	types := createTestTypes(alphaDefinition, betaDefinition)
	g.Expect(func() { types.AddAll(betaDefinition, deltaDefinition) }).To(Panic())
}

/*
 * AddTypes() Tests
 */

func Test_TypesAddTypes_GivenTypes_ModifiesSet(t *testing.T) {
	g := NewGomegaWithT(t)

	types := createTestTypes(alphaDefinition, betaDefinition)
	otherTypes := createTestTypes(gammaDefinition, deltaDefinition)

	types.AddTypes(otherTypes)

	g.Expect(types).To(ContainElements(gammaDefinition, deltaDefinition))
}

func Test_TypesAddTypes_GivenOverlappingTypes_Panics(t *testing.T) {
	g := NewGomegaWithT(t)

	types := createTestTypes(alphaDefinition, betaDefinition)
	otherTypes := createTestTypes(betaDefinition, deltaDefinition)

	g.Expect(func() { types.AddTypes(otherTypes) }).To(Panic())
}

/*
 * Where() Tests
 */

func Test_TypesWhere_GivenPredicate_ReturnsExpectedSet(t *testing.T) {
	g := NewGomegaWithT(t)

	types := createTestTypes(alphaDefinition, betaDefinition, gammaDefinition, deltaDefinition).
		Where(func(def TypeDefinition) bool {
			return len(def.name.name) == 5
		})

	g.Expect(types).To(ContainElements(alphaDefinition, gammaDefinition, deltaDefinition))
	g.Expect(types).NotTo(ContainElement(betaDefinition))
}

/*
 * Except() tests
 */

func Test_TypesExcept_GivenEmptySet_ReturnsExpectedSet(t *testing.T) {
	g := NewGomegaWithT(t)

	types := createTestTypes(alphaDefinition, betaDefinition, gammaDefinition, deltaDefinition)
	empty := make(Types)
	set := types.Except(empty)

	g.Expect(len(set)).To(Equal(len(types)))
	g.Expect(set).To(ContainElement(alphaDefinition))
	g.Expect(set).To(ContainElement(betaDefinition))
	g.Expect(set).To(ContainElement(gammaDefinition))
	g.Expect(set).To(ContainElement(deltaDefinition))
}

func Test_TypesExcept_GivenSelf_ReturnsEmptySet(t *testing.T) {
	g := NewGomegaWithT(t)

	types := createTestTypes(alphaDefinition, betaDefinition, gammaDefinition, deltaDefinition)
	set := types.Except(types)

	g.Expect(len(set)).To(Equal(0))
}

func Test_TypesExcept_GivenSubset_ReturnsExpectedSet(t *testing.T) {
	g := NewGomegaWithT(t)

	types := createTestTypes(alphaDefinition, betaDefinition, gammaDefinition, deltaDefinition)
	subset := createTestTypes(alphaDefinition, betaDefinition)
	set := types.Except(subset)

	g.Expect(len(set)).To(Equal(2))
	g.Expect(set).NotTo(ContainElement(alphaDefinition))
	g.Expect(set).NotTo(ContainElement(betaDefinition))
	g.Expect(set).To(ContainElement(gammaDefinition))
	g.Expect(set).To(ContainElement(deltaDefinition))
}

/*
 * Overlay() tests
 */

func Test_TypesOverlayWith_GivenDisjointSets_ReturnsUnionSet(t *testing.T) {
	g := NewGomegaWithT(t)
	left := createTestTypes(alphaDefinition, betaDefinition)
	right := createTestTypes(gammaDefinition, deltaDefinition)

	set := left.OverlayWith(right)

	g.Expect(len(set)).To(Equal(4))
	g.Expect(set).To(ContainElement(alphaDefinition))
	g.Expect(set).To(ContainElement(betaDefinition))
	g.Expect(set).To(ContainElement(gammaDefinition))
	g.Expect(set).To(ContainElement(deltaDefinition))
}

func Test_TypesOverlayWith_GivenOverlappingSets_PrefersTypeInOverlay(t *testing.T) {
	g := NewGomegaWithT(t)
	left := createTestTypes(alphaDefinition, deltaDefinition)
	right := createTestTypes(gammaDefinition, deltaIntDefinition)

	set := left.OverlayWith(right)

	g.Expect(len(set)).To(Equal(3))
	g.Expect(set).To(ContainElement(alphaDefinition))
	g.Expect(set).To(ContainElement(gammaDefinition))
	g.Expect(set).To(ContainElement(deltaIntDefinition))
	g.Expect(set).NotTo(ContainElement(deltaDefinition))
}

/*
 * Utility functions
 */

func createTestDefinition(name string, underlyingType Type) TypeDefinition {
	n := MakeTypeName(pkg, name)
	return MakeTypeDefinition(n, underlyingType)
}

func createTestTypes(defs ...TypeDefinition) Types {
	result := make(Types)
	for _, d := range defs {
		result.Add(d)
	}

	return result
}
