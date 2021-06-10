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

func TestResourceType_WithProperty_ContainsProperty(t *testing.T) {
	g := NewGomegaWithT(t)
	base := NewResourceType(emptySpec, emptyStatus)
	ownerProp := NewPropertyDefinition("Owner", "owner", StringType)
	resource := base.WithProperty(ownerProp)
	prop, ok := resource.Property("Owner")
	g.Expect(ok).To(BeTrue())
	g.Expect(prop).NotTo(BeNil())
}
