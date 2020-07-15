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

func NewTestStruct(name string, properties ...string) StructDefinition {
	var ps []*PropertyDefinition
	for _, n := range properties {
		ps = append(ps, NewPropertyDefinition(PropertyName(n), n, StringType))
	}

	ref := NewTypeName(*NewLocalPackageReference("group", "2020-01-01"), name)
	definition := NewStructDefinition(ref, NewStructType().WithProperties(ps...))

	return *definition
}
