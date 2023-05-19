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

type mockPackageReference struct {
	packagePath string
}

func (m mockPackageReference) PackageName() string {
	// Implement PackageName method as needed
	return ""
}

func (m mockPackageReference) PackagePath() string {
	return m.packagePath
}

func (m mockPackageReference) Equals(ref PackageReference) bool {
	// Implement Equals method as needed
	return m.packagePath == ref.PackagePath()
}

func (m mockPackageReference) String() string {
	// Implement String method as needed
	return ""
}

func (m mockPackageReference) IsPreview() bool {
	// Implement IsPreview method as needed
	return false
}

func (m mockPackageReference) TryGroupVersion() (string, string, bool) {
	// Implement TryGroupVersion method as needed
	return "", "", false
}

func (m mockPackageReference) GroupVersion() (string, string) {
	// Implement GroupVersion method as needed
	return "", ""
}

func TestSortTypeName(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Helper function to create PackageReference
	makePackageReference := func(packagePath string) PackageReference {
		return mockPackageReference{packagePath: packagePath}
	}

	// Test cases
	testCases := []struct {
		left     TypeName
		right    TypeName
		expected bool
	}{
		{
			left: TypeName{
				PackageReference: makePackageReference("github.com/Azure/azure-service-operator/v2"),
				name:             "TypeA",
			},
			right: TypeName{
				PackageReference: makePackageReference("github.com/Azure/azure-service-operator/v3"),
				name:             "TypeB",
			},
			expected: true,
		},
		{
			left: TypeName{
				PackageReference: makePackageReference("github.com/Azure/azure-service-operator/v3"),
				name:             "TypeA",
			},
			right: TypeName{
				PackageReference: makePackageReference("github.com/Azure/azure-service-operator/v2"),
				name:             "TypeB",
			},
			expected: false,
		},
		{
			left: TypeName{
				PackageReference: makePackageReference("github.com/Azure/azure-service-operator/v2"),
				name:             "TypeA",
			},
			right: TypeName{
				PackageReference: makePackageReference("github.com/Azure/azure-service-operator/v2"),
				name:             "TypeB",
			},
			expected: true,
		},
		{
			left: TypeName{
				PackageReference: makePackageReference("github.com/Azure/azure-service-operator/v2"),
				name:             "TypeB",
			},
			right: TypeName{
				PackageReference: makePackageReference("github.com/Azure/azure-service-operator/v2"),
				name:             "TypeA",
			},
			expected: false,
		},
		{
			left: TypeName{
				PackageReference: makePackageReference("github.com/Azure/azure-service-operator/v2"),
				name:             "TypeA",
			},
			right: TypeName{
				PackageReference: makePackageReference("github.com/Azure/azure-service-operator/v2"),
				name:             "TypeA",
			},
			expected: false,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		result := SortTypeName(tc.left, tc.right)
		g.Expect(result).To(Equal(tc.expected))
	}
}
