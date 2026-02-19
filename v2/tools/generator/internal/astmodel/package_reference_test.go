/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestSortPackageReferencesByPathAndVersion(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	v1 := makeTestLocalPackageReference("picnic", "v1")
	v2alpha := makeTestLocalPackageReference("picnic", "v2alpha")
	v2beta := makeTestLocalPackageReference("picnic", "v2beta")
	v2 := makeTestLocalPackageReference("picnic", "v2")
	v3 := makeTestLocalPackageReference("picnic", "v3")
	v10 := makeTestLocalPackageReference("picnic", "v10")

	references := []PackageReference{v10, v3, v2, v1, v2beta, v2alpha}
	expected := []PackageReference{v1, v2alpha, v2beta, v2, v3, v10}

	SortPackageReferencesByPathAndVersion(references)
	g.Expect(references).To(Equal(expected))
}

func TestComparePackageReferencesByPathAndVersion(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		left     string
		right    string
		expected int
	}{
		// Equality tests
		"grobo v1 is equal to self": {
			left:     "grobo/v1",
			right:    "grobo/v1",
			expected: 0,
		},
		"grobo v2 is equal to self": {
			left:     "grobo/v2",
			right:    "grobo/v2",
			expected: 0,
		},
		// Different length tests
		"Shorter is less than longer": {
			left:     "abc",
			right:    "abcdef",
			expected: -1,
		},
		"Longer is not less than shorter": {
			left:     "abcdef",
			right:    "abc",
			expected: 1,
		},
		// Numeric tests
		"Number equal to self": {
			left:     "42",
			right:    "42",
			expected: 0,
		},
		"grobo v1 is less than grobo v2": {
			left:     "grobo/v1",
			right:    "grobo/v2",
			expected: -1,
		},
		"grobo v2 is not less than grobo v1": {
			left:     "grobo/v2",
			right:    "grobo/v1",
			expected: 1,
		},
		// Prerelease tests
		"Alpha comes before beta": {
			left:     "1alpha",
			right:    "1beta",
			expected: -1,
		},
		"Alpha comes before aardvark": {
			left:     "1alpha",
			right:    "1aardvark",
			expected: -1,
		},
		"Alpha comes before GA": {
			left:     "1alpha",
			right:    "1",
			expected: -1,
		},
		"Beta comes before preview": {
			left:     "1beta",
			right:    "1preview",
			expected: -1,
		},
		"Beta comes before baboon": {
			left:     "1beta",
			right:    "1baboon",
			expected: -1,
		},
		"Beta comes before GA": {
			left:     "1beta",
			right:    "1",
			expected: -1,
		},
		"Preview comes before storage": {
			left:     "1preview",
			right:    "1storage",
			expected: -1,
		},
		"Preview comes before grape": {
			left:     "1preview",
			right:    "1grape",
			expected: -1,
		},
		"Preview comes before GA": {
			left:     "1preview",
			right:    "1",
			expected: -1,
		},
		// Version tests
		"v2.0.0 comes before v2.1.0": {
			left:     "v2.0.0",
			right:    "v2.1.0",
			expected: -1,
		},
		"v2.1.0 comes before v2.9.0": {
			left:     "v2.1.0",
			right:    "v2.9.0",
			expected: -1,
		},
		"v2.9.0 comes before v2.10.0": {
			left:     "v2.9.0",
			right:    "v2.10.0",
			expected: -1,
		},
		// Consistency tests, based on observed weirdness
		"Storage subpackage comes after base": {
			left:     "github.com/Azure/azure-service-operator/testing/person/v20200101",
			right:    "github.com/Azure/azure-service-operator/testing/person/v20200101/storage",
			expected: -1,
		},
		"Storage subpackage comes after base, reversed": {
			left:     "github.com/Azure/azure-service-operator/testing/person/v20200101/storage",
			right:    "github.com/Azure/azure-service-operator/testing/person/v20200101",
			expected: 1,
		},
		"Storage subpackage comes after preview base": {
			left:     "github.com/Azure/azure-service-operator/testing/person/v20211231preview",
			right:    "github.com/Azure/azure-service-operator/testing/person/v20211231preview/storage",
			expected: -1,
		},
		"Storage subpackage comes after preview base, reversed": {
			left:     "github.com/Azure/azure-service-operator/testing/person/v20211231preview/storage",
			right:    "github.com/Azure/azure-service-operator/testing/person/v20211231preview",
			expected: 1,
		},
		"Storage subpackage of non-preview comes after preview base": {
			left:     "github.com/Azure/azure-service-operator/testing/person/v20211231/storage",
			right:    "github.com/Azure/azure-service-operator/testing/person/v20211231preview",
			expected: 1,
		},
		"Storage subpackage of non-preview comes after preview base, reversed": {
			left:     "github.com/Azure/azure-service-operator/testing/person/v20211231preview",
			right:    "github.com/Azure/azure-service-operator/testing/person/v20211231/storage",
			expected: -1,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			comparer := versionComparer{
				left:  []rune(c.left),
				right: []rune(c.right),
			}

			g.Expect(comparer.Compare()).To(Equal(c.expected))
		})
	}
}

func TestVersionComparerCompareNumeric(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		left     string
		right    string
		expected int
	}{
		"Number equal to self": {
			left:     "42",
			right:    "42",
			expected: 0,
		},
		"Leading zeros on left-hand value don't matter": {
			left:     "042",
			right:    "42",
			expected: 0,
		},
		"Leading zeros on right-hand value don't matter": {
			left:     "42",
			right:    "042",
			expected: 0,
		},
		"Longer numbers are greater": {
			left:     "142",
			right:    "42",
			expected: 1,
		},
		"Shorter numbers are lesser": {
			left:     "42",
			right:    "942",
			expected: -1,
		},
		"Really really long equal numbers are equal": {
			left:     "1099511627776",
			right:    "1099511627776",
			expected: 0,
		},
		"Really really long non-equal numbers are comparable": {
			left:     "10995116277763",
			right:    "10995116277769",
			expected: -1,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			comparer := versionComparer{
				left:  []rune(c.left),
				right: []rune(c.right),
			}

			g.Expect(comparer.compareNumeric()).To(Equal(c.expected))
		})
	}
}

func TestContainsPreviewVersionLabel(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		version  string
		expected bool
	}{
		"Non-preview": {
			version:  "v20200201",
			expected: false,
		},
		"Alpha": {
			version:  "v20200201alpha",
			expected: true,
		},
		"Beta": {
			version:  "v20200201beta",
			expected: true,
		},
		"Preview": {
			version:  "v20200201preview",
			expected: true,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			isPreview := ContainsPreviewVersionLabel(c.version)
			g.Expect(isPreview).To(Equal(c.expected))
		})
	}
}
