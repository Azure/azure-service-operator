/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestSingular_GivesExpectedResults(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		expected string
	}{
		{"Account", "Account"},
		{"Accounts", "Account"},
		{"Address", "Address"},
		{"Addresses", "Address"},
		{"Batch", "Batch"},
		{"Batches", "Batch"},
		{"ImportServices", "ImportService"},
		{"Exportservices", "Exportservice"},
		{"AzureRedis", "AzureRedis"},
		{"Aliases", "Alias"},
		{"AdoptedFoxes", "AdoptedFox"},
	}

	ref := makeTestLocalPackageReference("Demo", "v2010")

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			name := MakeTypeName(ref, c.name)
			result := name.Singular()
			g.Expect(result.name).To(Equal(c.expected))
		})
	}
}

func TestPlural_GivesExpectedResults(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		expected string
	}{
		{"Account", "Accounts"},
		{"Accounts", "Accounts"},
		{"Batch", "Batches"},
		{"Batches", "Batches"},
		{"ImportService", "ImportServices"},
		{"Exportservice", "Exportservices"},
		{"AzureRedis", "AzureRedis"},
	}

	ref := makeTestLocalPackageReference("Demo", "v2010")

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			name := MakeTypeName(ref, c.name)
			result := name.Plural()
			g.Expect(result.name).To(Equal(c.expected))
		})
	}
}

func TestTypeName_IsEmpty(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	ref := makeTestLocalPackageReference("Demo", "v2010")
	blank := TypeName{}
	name := MakeTypeName(ref, "Person")

	g.Expect(blank.IsEmpty()).To(BeTrue())
	g.Expect(name.IsEmpty()).To(BeFalse())
}
