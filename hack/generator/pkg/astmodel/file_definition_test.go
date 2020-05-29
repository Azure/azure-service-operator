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
	file := NewFileDefinition(&person.Name().PackageReference, &person)

	g.Expect(*file.packageReference).To(Equal(person.Name().PackageReference))
	g.Expect(file.definitions).To(HaveLen(1))
}

func NewTestStruct(name string, fields ...string) StructDefinition {
	var fs []*FieldDefinition
	for _, n := range fields {
		fs = append(fs, NewFieldDefinition(FieldName(n), n, StringType))
	}

	ref := NewTypeName(*NewLocalPackageReference("group", "2020-01-01"), name)
	definition := NewStructDefinition(ref, NewStructType(fs...), false)

	return *definition
}
