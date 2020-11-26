/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	. "github.com/onsi/gomega"
	"testing"
)

//TODO (theunrepentantgeek): MOAR TESTS

/*
 * WithTestCase() tests
 */

func TestResourceType_WithTestCase_ReturnsExpectedObject(t *testing.T) {
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
