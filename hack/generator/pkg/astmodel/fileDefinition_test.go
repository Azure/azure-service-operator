/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_NewFileDefinition_GivenValues_InitializesFields(t *testing.T) {
	g := NewGomegaWithT(t)

	person := NewTestStruct("Person", "fullName", "knownAs", "familyName")
	file := NewFileDefinition(&person)

	g.Expect(file.PackageReference).To(Equal(person.PackageReference))
	g.Expect(file.structs).To(HaveLen(1))
}

func NewTestStruct(name string, fields ...string) StructDefinition {
	var fs []*FieldDefinition
	for _, n := range fields {
		fs = append(fs, NewFieldDefinition(n, n, StringType))
	}

	ref := NewStructReference(name, "group", "2020-01-01")
	definition := NewStructDefinition(ref, fs...)

	return *definition
}
