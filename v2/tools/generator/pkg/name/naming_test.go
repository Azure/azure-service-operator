/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package name

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_Singularize(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		text     string
		expected string
	}{
		{"Empty", "", ""},
		{"Person", "Person", "Person"},
		{"Persons", "Persons", "Person"},
		{"People", "People", "Person"},
		{"Peoples", "Peoples", "People"},
		{"Redis", "Redis", "Redis"},
		{"redis", "redis", "redis"},
		{"Service", "Service", "Service"},
		{"Services", "Services", "Service"},
		{"FIPS", "FIPS", "FIPS"},
		{"ID", "ID", "ID"},
		{"IDs", "IDs", "ID"},
		{"SSH", "SSH", "SSH"},
		{"SubnetID", "subnetID", "subnetID"},
		{"SubnetIDs", "subnetIDs", "subnetID"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			actual := Singularize(c.text)
			g.Expect(actual).To(Equal(c.expected))
		})
	}
}

func Test_Pluralize(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		text     string
		expected string
	}{
		{"Empty", "", ""},
		{"Person", "Person", "People"},
		{"Persons", "Persons", "Persons"},
		{"People", "People", "People"},
		{"Peoples", "Peoples", "Peoples"},
		{"Redis", "Redis", "Redis"},
		{"redis", "redis", "redis"},
		{"Service", "Service", "Services"},
		{"Services", "Services", "Services"},
		{"FIPS", "FIPS", "FIPS"},
		{"ID", "ID", "IDs"},
		{"IDs", "IDs", "IDs"},
		{"SSH", "SSH", "SSHs"},
		{"SubnetID", "subnetID", "subnetIDs"},
		{"SubnetIDs", "subnetIDs", "subnetIDs"},
		{"SubnetId", "subnetId", "subnetIds"},
		{"SubnetIds", "subnetIds", "subnetIds"},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			actual := Pluralize(c.text)
			g.Expect(actual).To(Equal(c.expected))
		})
	}
}
