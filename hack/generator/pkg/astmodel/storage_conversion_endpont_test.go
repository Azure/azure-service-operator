/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestKnownLocalsSetCreateLocal(t *testing.T) {
	g := NewGomegaWithT(t)

	idFactory := NewIdentifierFactory()
	knownLocals := NewKnownLocalsSet(idFactory)

	g.Expect(knownLocals.createLocal("person", "Item", "Map", "List")).To(Equal("personItem"))
	g.Expect(knownLocals.createLocal("person", "Item", "Map", "List")).To(Equal("personMap"))
	g.Expect(knownLocals.createLocal("person", "Item", "Map", "List")).To(Equal("personList"))
	g.Expect(knownLocals.createLocal("person", "Item", "Map", "List")).To(Equal("personItem1"))
	g.Expect(knownLocals.createLocal("person", "Item", "Map", "List")).To(Equal("personMap1"))
	g.Expect(knownLocals.createLocal("person", "Item", "Map", "List")).To(Equal("personList1"))
	g.Expect(knownLocals.createLocal("person", "Item", "Map", "List")).To(Equal("personItem2"))
	g.Expect(knownLocals.createLocal("person", "Item", "Map", "List")).To(Equal("personMap2"))
	g.Expect(knownLocals.createLocal("person", "Item", "Map", "List")).To(Equal("personList2"))

	// Case insensitivity
	g.Expect(knownLocals.createLocal("Student")).To(Equal("student"))
	g.Expect(knownLocals.createLocal("Student")).To(Equal("student1"))
	g.Expect(knownLocals.createLocal("student")).To(Equal("student2"))
	g.Expect(knownLocals.createLocal("student")).To(Equal("student3"))
}
