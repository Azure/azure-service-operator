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
	pkg                = makeTestLocalPackageReference("foo", "2020-01-01")
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
	t.Parallel()
	g := NewGomegaWithT(t)

	types := make(TypeDefinitionSet)
	types.Add(alphaDefinition)

	g.Expect(types).To(ContainElement(alphaDefinition))
}

func Test_TypesAdd_GivenTypeAlreadyPresent_Panics(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	types := make(TypeDefinitionSet)
	types.Add(alphaDefinition)

	g.Expect(func() { types.Add(alphaDefinition) }).To(Panic())
}

/*
 * AddAll() Tests
 */

func Test_TypesAddAll_GivenTypes_ModifiesSet(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	types := createTestTypes(alphaDefinition, betaDefinition)
	types.AddAll(gammaDefinition, deltaDefinition)

	g.Expect(types).To(ContainElements(gammaDefinition, deltaDefinition))
}

func Test_TypesAddAll_GivenOverlappingTypes_Panics(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	types := createTestTypes(alphaDefinition, betaDefinition)
	g.Expect(func() { types.AddAll(betaDefinition, deltaDefinition) }).To(Panic())
}

/*
 * AddTypes() Tests
 */

func Test_TypesAddTypes_GivenTypes_ModifiesSet(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	types := createTestTypes(alphaDefinition, betaDefinition)
	otherTypes := createTestTypes(gammaDefinition, deltaDefinition)

	types.AddTypes(otherTypes)

	g.Expect(types).To(ContainElements(gammaDefinition, deltaDefinition))
}

func Test_TypesAddTypes_GivenOverlappingTypes_Panics(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	types := createTestTypes(alphaDefinition, betaDefinition)
	otherTypes := createTestTypes(betaDefinition, deltaDefinition)

	g.Expect(func() { types.AddTypes(otherTypes) }).To(Panic())
}

/*
 * Where() Tests
 */

func Test_TypesWhere_GivenPredicate_ReturnsExpectedSet(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
	g := NewGomegaWithT(t)

	types := createTestTypes(alphaDefinition, betaDefinition, gammaDefinition, deltaDefinition)
	empty := make(TypeDefinitionSet)
	set := types.Except(empty)

	g.Expect(len(set)).To(Equal(len(types)))
	g.Expect(set).To(ContainElement(alphaDefinition))
	g.Expect(set).To(ContainElement(betaDefinition))
	g.Expect(set).To(ContainElement(gammaDefinition))
	g.Expect(set).To(ContainElement(deltaDefinition))
}

func Test_TypesExcept_GivenSelf_ReturnsEmptySet(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	types := createTestTypes(alphaDefinition, betaDefinition, gammaDefinition, deltaDefinition)
	set := types.Except(types)

	g.Expect(len(set)).To(Equal(0))
}

func Test_TypesExcept_GivenSubset_ReturnsExpectedSet(t *testing.T) {
	t.Parallel()
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
 * Intersect() tests
 */

func Test_TypesIntersect_GivenEmptySet_ReturnsExpectedSet(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	types := createTestTypes(alphaDefinition, betaDefinition, gammaDefinition, deltaDefinition)
	empty := make(TypeDefinitionSet)
	set := types.Intersect(empty)

	g.Expect(len(set)).To(Equal(0))
}

func Test_TypesIntersect_GivenSelf_ReturnsSelf(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	types := createTestTypes(alphaDefinition, betaDefinition, gammaDefinition, deltaDefinition)
	set := types.Intersect(types)

	g.Expect(len(set)).To(Equal(len(types)))
	g.Expect(set).To(ContainElement(alphaDefinition))
	g.Expect(set).To(ContainElement(betaDefinition))
	g.Expect(set).To(ContainElement(gammaDefinition))
	g.Expect(set).To(ContainElement(deltaDefinition))
}

func Test_TypesIntersect_GivenSubset_ReturnsExpectedSet(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	types := createTestTypes(alphaDefinition, betaDefinition, gammaDefinition, deltaDefinition)
	subset := createTestTypes(alphaDefinition, betaDefinition)
	set := types.Intersect(subset)

	g.Expect(len(set)).To(Equal(2))
	g.Expect(set).To(ContainElement(alphaDefinition))
	g.Expect(set).To(ContainElement(betaDefinition))
	g.Expect(set).NotTo(ContainElement(gammaDefinition))
	g.Expect(set).NotTo(ContainElement(deltaDefinition))
}

/*
 * Overlay() tests
 */

func Test_TypesOverlayWith_GivenDisjointSets_ReturnsUnionSet(t *testing.T) {
	t.Parallel()
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
	t.Parallel()
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
 * FindSpecTypes() tests
 */

func TestFindSpecTypes(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Define a test resource
	spec := createTestSpec("Person", fullNameProperty, knownAsProperty)
	status := createTestStatus("Person")
	resource := createTestResource("Person", spec, status)

	types := make(TypeDefinitionSet)
	types.AddAll(resource, status, spec)

	specs := FindSpecTypes(types)

	g.Expect(specs).To(HaveLen(1))
	g.Expect(specs.Contains(spec.Name())).To(BeTrue())
}

/*
 * FindStatusTypes() tests
 */

func TestFindStatusTypes(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Define a test resource
	spec := createTestSpec("Person", fullNameProperty, knownAsProperty)
	status := createTestStatus("Person")
	resource := createTestResource("Person", spec, status)

	types := make(TypeDefinitionSet)
	types.AddAll(resource, status, spec)

	statuses := FindStatusTypes(types)

	g.Expect(statuses).To(HaveLen(1))
	g.Expect(statuses.Contains(status.Name())).To(BeTrue())
}

/*
 * FindSpecConnectedTypes() tests
 */

func TestFindSpecConnectedTypes(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Define a test resource
	nameInfo := NewTestObject("NameInfo", fullNameProperty, knownAsProperty)
	nameInfoProperty := NewPropertyDefinition("NameInfo", "nameInfo", nameInfo.Name())
	spec := createTestSpec("PersonAugmented", nameInfoProperty)

	nameInfoStatus := NewTestObject("NameInfoStatus", fullNameProperty, knownAsProperty)
	nameInfoStatusProperty := NewPropertyDefinition("NameInfo", "nameInfo", nameInfoStatus.Name())
	status := createTestStatus("PersonAugmented", nameInfoStatusProperty)

	resource := createTestResource("PersonAugmented", spec, status)

	types := make(TypeDefinitionSet)
	types.AddAll(resource, status, spec, nameInfo, nameInfoStatus)

	specs, err := FindSpecConnectedTypes(types)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(specs).To(HaveLen(2))
	g.Expect(specs.Contains(spec.Name())).To(BeTrue())
	g.Expect(specs.Contains(nameInfo.Name())).To(BeTrue())
}

/*
 * FindStatusConnectedTypes() tests
 */

func TestFindStatusConnectedTypes(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Define a test resource
	nameInfo := NewTestObject("NameInfo", fullNameProperty, knownAsProperty)
	nameInfoProperty := NewPropertyDefinition("NameInfo", "nameInfo", nameInfo.Name())
	spec := createTestSpec("PersonAugmented", nameInfoProperty)

	nameInfoStatus := NewTestObject("NameInfoStatus", fullNameProperty, knownAsProperty)
	nameInfoStatusProperty := NewPropertyDefinition("NameInfo", "nameInfo", nameInfoStatus.Name())
	status := createTestStatus("PersonAugmented", nameInfoStatusProperty)

	resource := createTestResource("PersonAugmented", spec, status)

	types := make(TypeDefinitionSet)
	types.AddAll(resource, status, spec, nameInfo, nameInfoStatus)

	statuses, err := FindStatusConnectedTypes(types)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(statuses).To(HaveLen(2))
	g.Expect(statuses.Contains(status.Name())).To(BeTrue())
	g.Expect(statuses.Contains(nameInfoStatus.Name())).To(BeTrue())
}

/*
 * ResolveResourceSpecAndStatus() tests
 */

func TestResolveResourceSpecAndStatus(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Define a test resource
	spec := createTestSpec("Person", fullNameProperty, knownAsProperty)
	status := createTestStatus("Person")
	resource := createTestResource("Person", spec, status)

	types := make(TypeDefinitionSet)
	types.AddAll(resource, status, spec)

	resolved, err := types.ResolveResourceSpecAndStatus(resource)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(resolved.SpecDef.Name()).To(Equal(spec.Name()))
	g.Expect(resolved.SpecDef.Type()).To(Equal(spec.Type()))
	g.Expect(resolved.SpecType).To(Equal(spec.Type()))
	g.Expect(resolved.StatusDef.Name()).To(Equal(status.Name()))
	g.Expect(resolved.StatusDef.Type()).To(Equal(status.Type()))
	g.Expect(resolved.StatusType).To(Equal(status.Type()))
	g.Expect(resolved.ResourceDef.Name()).To(Equal(resource.Name()))
	g.Expect(resolved.ResourceDef.Type()).To(Equal(resource.Type()))
	g.Expect(resolved.ResourceType).To(Equal(resource.Type()))
}

/*
 * Utility functions
 */

func createTestDefinition(name string, underlyingType Type) TypeDefinition {
	n := MakeTypeName(pkg, name)
	return MakeTypeDefinition(n, underlyingType)
}

func createTestTypes(defs ...TypeDefinition) TypeDefinitionSet {
	result := make(TypeDefinitionSet)
	for _, d := range defs {
		result.Add(d)
	}

	return result
}

// CreateTestResource makes a resource for testing
func createTestResource(
	name string,
	spec TypeDefinition,
	status TypeDefinition) TypeDefinition {

	resourceType := NewResourceType(spec.Name(), status.Name())
	return MakeTypeDefinition(MakeTypeName(pkg, name), resourceType)
}

// createTestSpec makes a spec for testing
func createTestSpec(
	name string,
	properties ...*PropertyDefinition) TypeDefinition {
	specName := MakeTypeName(pkg, name+"_Spec")
	return MakeTypeDefinition(
		specName,
		NewObjectType().WithProperties(properties...))
}

// createTestStatus makes a status for testing
func createTestStatus(
	name string,
	properties ...*PropertyDefinition) TypeDefinition {
	statusName := MakeTypeName(pkg, name+"_Status")
	return MakeTypeDefinition(
		statusName,
		NewObjectType().WithProperties(properties...))
}
