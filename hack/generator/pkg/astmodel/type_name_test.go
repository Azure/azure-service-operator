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
	cases := []struct {
		name     string
		expected string
	}{
		{"Account", "Account"},
		{"Accounts", "Account"},
		{"Batch", "Batch"},
		{"Batches", "Batch"},
		{"ImportServices", "ImportService"},
		{"Exportservices", "Exportservice"},
		{"AzureRedis", "AzureRedis"},
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
