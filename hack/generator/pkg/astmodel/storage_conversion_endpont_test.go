/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestKnownLocalsSetCreateLocal(t *testing.T) {
	g := NewGomegaWithT(t)

	idFactory := NewIdentifierFactory()
	knownLocals := NewKnownLocalsSet(idFactory)

	g.Expect(knownLocals.createLocal("person")).To(Equal("person"))
	g.Expect(knownLocals.createLocal("person")).To(Equal("person1"))
	g.Expect(knownLocals.createLocal("person")).To(Equal("person2"))

	// Case insensitivity
	g.Expect(knownLocals.createLocal("Student")).To(Equal("student"))
	g.Expect(knownLocals.createLocal("Student")).To(Equal("student1"))
	g.Expect(knownLocals.createLocal("student")).To(Equal("student2"))
	g.Expect(knownLocals.createLocal("student")).To(Equal("student3"))
}
