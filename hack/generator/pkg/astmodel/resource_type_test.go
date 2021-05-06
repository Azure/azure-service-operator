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

/*
 * WithTestCase() tests
 */

func TestResourceType_WithTestCase_ReturnsExpectedInstance(t *testing.T) {
	g := NewGomegaWithT(t)

	spec := NewObjectType()
	status := NewObjectType()
	name := "assertStuff"
	fake := NewFakeTestCase(name)
	base := NewResourceType(spec, status)

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

	spec := NewObjectType()
	status := NewObjectType()
	name := "assertStuff"

	fake := NewFakeFunction(name)
	base := NewResourceType(spec, status)

	resource := base.WithFunction(fake)

	g.Expect(base).NotTo(Equal(resource)) // Ensure the original wasn't modified
	g.Expect(resource.functions).To(HaveLen(1))
	g.Expect(resource.functions[name]).To(Equal(fake))

}
