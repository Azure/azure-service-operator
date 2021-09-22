/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

//TODO (theunrepentantgeek): MOAR TESTS

var (
	emptySpec   = NewObjectType()
	emptyStatus = NewObjectType()
)

/*
 * WithTestCase() tests
 */

func TestResourceType_WithTestCase_ReturnsExpectedInstance(t *testing.T) {
	g := NewGomegaWithT(t)

	name := "assertStuff"
	fake := NewFakeTestCase(name)
	base := NewResourceType(emptySpec, emptyStatus)

	resource := base.WithTestCase(fake)

	g.Expect(base).NotTo(Equal(resource)) // Ensure the original wasn't modified
	g.Expect(resource.testcases).To(HaveLen(1))
	g.Expect(resource.testcases[name]).To(Equal(fake))
}

/*
 * WithFunction() tests
 */

func TestResourceType_WithFunction_ReturnsExpectedInstance(t *testing.T) {
	g := NewGomegaWithT(t)

	name := "assertStuff"

	fake := NewFakeFunction(name)
	base := NewResourceType(emptySpec, emptyStatus)

	resource := base.WithFunction(fake)

	g.Expect(base).NotTo(Equal(resource)) // Ensure the original wasn't modified
	g.Expect(resource.functions).To(HaveLen(1))
	g.Expect(resource.functions[name]).To(Equal(fake))

}

/*
 * Properties() tests
 */

func TestResourceType_Properties_ReturnsExpectedCount(t *testing.T) {
	g := NewGomegaWithT(t)
	base := NewResourceType(emptySpec, emptyStatus)
	g.Expect(base.Properties()).To(HaveLen(2))
}

/*
 * Property() tests
 */

func TestResourceType_Property_ForStatus_ReturnsProperty(t *testing.T) {
	g := NewGomegaWithT(t)
	base := NewResourceType(emptySpec, emptyStatus)
	prop, ok := base.Property("Spec")
	g.Expect(ok).To(BeTrue())
	g.Expect(prop).NotTo(BeNil())
}

func TestResourceType_Property_ForSpec_ReturnsProperty(t *testing.T) {
	g := NewGomegaWithT(t)
	base := NewResourceType(emptySpec, emptyStatus)
	prop, ok := base.Property("Spec")
	g.Expect(ok).To(BeTrue())
	g.Expect(prop).NotTo(BeNil())
}

/*
 * WithProperty() tests
 */

func TestResourceType_WithProperty_HasExpectedLength(t *testing.T) {
	g := NewGomegaWithT(t)
	base := NewResourceType(emptySpec, emptyStatus)
	ownerProp := NewPropertyDefinition("Owner", "owner", StringType)
	resource := base.WithProperty(ownerProp)
	g.Expect(resource.Properties()).To(HaveLen(3))
}

func TestResourceType_WithProperty_IncludesProperty(t *testing.T) {
	g := NewGomegaWithT(t)
	base := NewResourceType(emptySpec, emptyStatus)
	ownerProp := NewPropertyDefinition("Owner", "owner", StringType)
	resource := base.WithProperty(ownerProp)
	prop, ok := resource.Property("Owner")
	g.Expect(ok).To(BeTrue())
	g.Expect(prop).NotTo(BeNil())
}

func TestResourceType_WithProperty_OverridingSpec_Panics(t *testing.T) {
	g := NewGomegaWithT(t)
	base := NewResourceType(emptySpec, emptyStatus)
	specProp := NewPropertyDefinition("Spec", "spec", StringType)
	g.Expect(func() { base.WithProperty(specProp) }).To(Panic())
}

func TestResourceType_WithProperty_OverridingStatus_Panics(t *testing.T) {
	g := NewGomegaWithT(t)
	base := NewResourceType(emptySpec, emptyStatus)
	statusProp := NewPropertyDefinition("Status", "status", StringType)
	g.Expect(func() { base.WithProperty(statusProp) }).To(Panic())
}

/*
 * WithoutProperty() tests
 */

func TestResourceType_WithoutProperty_ExcludesProperty(t *testing.T) {
	g := NewGomegaWithT(t)
	ownerProp := NewPropertyDefinition("Owner", "owner", StringType)
	base := NewResourceType(emptySpec, emptyStatus).WithProperty(ownerProp)
	resource := base.WithoutProperty("Owner")
	prop, ok := resource.Property("Owner")
	g.Expect(ok).To(BeFalse())
	g.Expect(prop).To(BeNil())
}

func TestResourceType_WithoutSpecProperty_Panics(t *testing.T) {
	g := NewGomegaWithT(t)
	base := NewResourceType(emptySpec, emptyStatus)
	g.Expect(func() { base.WithoutProperty("Spec") }).To(Panic())
}

func TestResourceType_WithoutStatusProperty_Panics(t *testing.T) {
	g := NewGomegaWithT(t)
	base := NewResourceType(emptySpec, emptyStatus)
	g.Expect(func() { base.WithoutProperty("Status") }).To(Panic())
}

/*
 * WithAnnotation() tests
 */

func TestResourceType_WithAnnotation_ReturnsExpectedInstance(t *testing.T) {
	g := NewGomegaWithT(t)

	annotation := "kubebuilder:annotation:whatever"
	base := NewResourceType(emptySpec, emptyStatus)
	updated := base.WithAnnotation(annotation)

	g.Expect(base).NotTo(Equal(updated)) // Ensure the original wasn't modified
	g.Expect(base.Equals(updated)).To(BeFalse())
	g.Expect(updated.annotations).To(HaveLen(1))
	g.Expect(updated.annotations[0]).To(Equal(annotation))
}
