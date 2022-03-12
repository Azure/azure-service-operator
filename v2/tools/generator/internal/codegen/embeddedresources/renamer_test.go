/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package embeddedresources

import (
	"testing"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"

	. "github.com/onsi/gomega"
)

var (
	exampleTypeFlag   = astmodel.TypeFlag("flag")
	resourceTypeName  = newTestName("Resource")
	resourceTypeName2 = newTestName("Resource2")
)

func newTestName(name string) astmodel.TypeName {
	return astmodel.MakeTypeName(test.MakeLocalPackageReference("group", "2020-01-01"), name)
}

func newTestObject(name astmodel.TypeName, fields ...*astmodel.PropertyDefinition) astmodel.TypeDefinition {
	return astmodel.MakeTypeDefinition(name, astmodel.NewObjectType().WithProperties(fields...))
}

func typesWithSubresourceTypeNoOriginalNameUsage() astmodel.TypeDefinitionSet {
	result := make(astmodel.TypeDefinitionSet)

	suffix := "TestSuffix"

	originalTypeName := newTestName("T1")
	modifiedTypeName := embeddedResourceTypeName{original: originalTypeName, context: resourceTypeName.Name(), suffix: suffix, count: 0}.ToTypeName()
	modifiedObject := newTestObject(modifiedTypeName)
	result.Add(modifiedObject.WithType(exampleTypeFlag.ApplyTo(modifiedObject.Type())))

	prop := astmodel.NewPropertyDefinition(
		"prop1",
		"prop1",
		modifiedTypeName)
	resource := newTestObject(resourceTypeName, prop)
	result.Add(resource)

	return result
}

func typesWithSubresourceTypeOriginalNameUsage() astmodel.TypeDefinitionSet {
	result := make(astmodel.TypeDefinitionSet)

	suffix := "TestSuffix"

	originalTypeName := newTestName("T1")
	modifiedTypeName := embeddedResourceTypeName{original: originalTypeName, context: resourceTypeName.Name(), suffix: suffix, count: 0}.ToTypeName()
	modifiedObject := newTestObject(modifiedTypeName)
	result.Add(modifiedObject.WithType(exampleTypeFlag.ApplyTo(modifiedObject.Type())))
	result.Add(newTestObject(originalTypeName))

	prop1 := astmodel.NewPropertyDefinition(
		"prop1",
		"prop1",
		modifiedTypeName)
	prop2 := astmodel.NewPropertyDefinition(
		"prop2",
		"prop2",
		originalTypeName)

	resource := newTestObject(resourceTypeName, prop1, prop2)
	result.Add(resource)

	return result
}

func typesWithSubresourceTypeMultipleUsageContextsOneResource() astmodel.TypeDefinitionSet {
	result := make(astmodel.TypeDefinitionSet)

	suffix := "TestSuffix"

	originalTypeName := newTestName("T1")
	modifiedTypeName1 := embeddedResourceTypeName{original: originalTypeName, context: resourceTypeName.Name(), suffix: suffix, count: 0}.ToTypeName()
	modifiedObject1 := newTestObject(modifiedTypeName1)
	result.Add(modifiedObject1.WithType(exampleTypeFlag.ApplyTo(modifiedObject1.Type())))

	modifiedTypeName2 := embeddedResourceTypeName{original: originalTypeName, context: resourceTypeName.Name(), suffix: suffix, count: 1}.ToTypeName()
	modifiedObject2 := newTestObject(modifiedTypeName2)
	result.Add(modifiedObject2.WithType(exampleTypeFlag.ApplyTo(modifiedObject2.Type())))

	// result.Add(newTestObject(originalTypeName))

	prop1 := astmodel.NewPropertyDefinition(
		"prop1",
		"prop1",
		modifiedTypeName1)
	prop2 := astmodel.NewPropertyDefinition(
		"prop2",
		"prop2",
		modifiedTypeName2)

	resource := newTestObject(resourceTypeName, prop1, prop2)
	result.Add(resource)

	return result
}

func typesWithSubresourceTypeMultipleResourcesOneUsageContextEach() astmodel.TypeDefinitionSet {
	result := make(astmodel.TypeDefinitionSet)

	suffix := "TestSuffix"

	originalTypeName := newTestName("T1")
	modifiedTypeName1 := embeddedResourceTypeName{original: originalTypeName, context: resourceTypeName.Name(), suffix: suffix, count: 0}.ToTypeName()
	modifiedObject1 := newTestObject(modifiedTypeName1)
	result.Add(modifiedObject1.WithType(exampleTypeFlag.ApplyTo(modifiedObject1.Type())))

	modifiedTypeName2 := embeddedResourceTypeName{original: originalTypeName, context: resourceTypeName2.Name(), suffix: suffix, count: 0}.ToTypeName()
	modifiedObject2 := newTestObject(modifiedTypeName2)
	result.Add(modifiedObject2.WithType(exampleTypeFlag.ApplyTo(modifiedObject2.Type())))

	prop1 := astmodel.NewPropertyDefinition(
		"prop1",
		"prop1",
		modifiedTypeName1)
	prop2 := astmodel.NewPropertyDefinition(
		"prop2",
		"prop2",
		modifiedTypeName2)

	resource := newTestObject(resourceTypeName, prop1, prop2)
	result.Add(resource)

	return result
}

func TestCleanupTypeNames_TypeWithNoOriginalName_UpdatedNameCollapsed(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	expectedUpdatedTypeName := newTestName("T1")

	updatedTypes, err := simplifyTypeNames(typesWithSubresourceTypeNoOriginalNameUsage(), exampleTypeFlag)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(updatedTypes)).To(Equal(2))

	ot, ok := astmodel.AsObjectType(updatedTypes[resourceTypeName].Type())
	g.Expect(ok).To(BeTrue())

	property, ok := ot.Property("prop1")
	g.Expect(ok).To(BeTrue())

	propertyTypeName, ok := astmodel.AsTypeName(property.PropertyType())
	g.Expect(ok).To(BeTrue())

	g.Expect(astmodel.TypeEquals(propertyTypeName, expectedUpdatedTypeName)).To(BeTrue())
}

func TestCleanupTypeNames_TypeWithOriginalNameExists_UpdatedNamePartiallyCollapsed(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	expectedUpdatedTypeName := newTestName("T1_TestSuffix")
	expectedOriginalTypeName := newTestName("T1")

	updatedTypes, err := simplifyTypeNames(typesWithSubresourceTypeOriginalNameUsage(), exampleTypeFlag)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(updatedTypes)).To(Equal(3))

	ot, ok := astmodel.AsObjectType(updatedTypes[resourceTypeName].Type())
	g.Expect(ok).To(BeTrue())

	prop1, ok := ot.Property("prop1")
	g.Expect(ok).To(BeTrue())
	prop1TypeName, ok := astmodel.AsTypeName(prop1.PropertyType())
	g.Expect(ok).To(BeTrue())
	g.Expect(astmodel.TypeEquals(prop1TypeName, expectedUpdatedTypeName)).To(BeTrue())

	prop2, ok := ot.Property("prop2")
	g.Expect(ok).To(BeTrue())
	prop2TypeName, ok := astmodel.AsTypeName(prop2.PropertyType())
	g.Expect(ok).To(BeTrue())
	g.Expect(astmodel.TypeEquals(prop2TypeName, expectedOriginalTypeName)).To(BeTrue())
}

func TestCleanupTypeNames_UpdatedNamesAreAllForSameResource_UpdatedNamesStrippedOfResourceContext(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	expectedUpdatedTypeName1 := newTestName("T1_TestSuffix")
	expectedUpdatedTypeName2 := newTestName("T1_TestSuffix_1")

	updatedTypes, err := simplifyTypeNames(typesWithSubresourceTypeMultipleUsageContextsOneResource(), exampleTypeFlag)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(updatedTypes)).To(Equal(3))

	ot, ok := astmodel.AsObjectType(updatedTypes[resourceTypeName].Type())
	g.Expect(ok).To(BeTrue())

	prop1, ok := ot.Property("prop1")
	g.Expect(ok).To(BeTrue())
	prop1TypeName, ok := astmodel.AsTypeName(prop1.PropertyType())
	g.Expect(ok).To(BeTrue())
	g.Expect(astmodel.TypeEquals(prop1TypeName, expectedUpdatedTypeName1)).To(BeTrue())

	prop2, ok := ot.Property("prop2")
	g.Expect(ok).To(BeTrue())
	prop2TypeName, ok := astmodel.AsTypeName(prop2.PropertyType())
	g.Expect(ok).To(BeTrue())
	g.Expect(astmodel.TypeEquals(prop2TypeName, expectedUpdatedTypeName2)).To(BeTrue())
}

func TestCleanupTypeNames_UpdatedNamesAreEachForDifferentResource_UpdatedNamesStrippedOfCount(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	expectedUpdatedTypeName1 := newTestName("T1_Resource_TestSuffix")
	expectedUpdatedTypeName2 := newTestName("T1_Resource2_TestSuffix")

	updatedTypes, err := simplifyTypeNames(typesWithSubresourceTypeMultipleResourcesOneUsageContextEach(), exampleTypeFlag)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(len(updatedTypes)).To(Equal(3))

	ot, ok := astmodel.AsObjectType(updatedTypes[resourceTypeName].Type())
	g.Expect(ok).To(BeTrue())

	prop1, ok := ot.Property("prop1")
	g.Expect(ok).To(BeTrue())
	prop1TypeName, ok := astmodel.AsTypeName(prop1.PropertyType())
	g.Expect(ok).To(BeTrue())
	g.Expect(astmodel.TypeEquals(prop1TypeName, expectedUpdatedTypeName1)).To(BeTrue())

	prop2, ok := ot.Property("prop2")
	g.Expect(ok).To(BeTrue())
	prop2TypeName, ok := astmodel.AsTypeName(prop2.PropertyType())
	g.Expect(ok).To(BeTrue())
	g.Expect(astmodel.TypeEquals(prop2TypeName, expectedUpdatedTypeName2)).To(BeTrue())
}
