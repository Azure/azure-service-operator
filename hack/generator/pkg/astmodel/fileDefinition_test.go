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

	packageName := "demo"
	person := NewTestStruct("Person", "fullName", "knownAs", "familyName")
	file := NewFileDefinition(packageName, &person)

	g.Expect(file.packageName).To(Equal(packageName))
	g.Expect(file.structs).To(HaveLen(1))
}

func NewTestStruct(name string, fields ...string) StructDefinition {
	var fs []*FieldDefinition
	for _, n := range fields {
		fs = append(fs, NewFieldDefinition(n, n, StringType))
	}

	definition := NewStructDefinition(name, "2020-01-01", fs...)

	return *definition
}
