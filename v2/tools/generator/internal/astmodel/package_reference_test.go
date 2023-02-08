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

	cases := []struct {
		name     string
		left     string
		right    string
		expected int
	}{
		// Equality tests
		{"grobo v1 is equal to self", "grobo/v1", "grobo/v1", 0},
		{"grobo v2 is equal to self", "grobo/v2", "grobo/v2", 0},
		// Different length tests
		{"Shorter is less than longer", "abc", "abcdef", -1},
		{"Longer is not less than shorter", "abcdef", "abc", 1},
		// Numeric tests
		{"Number equal to self", "42", "42", 0},
		{"grobo v1 is less than grobo v2", "grobo/v1", "grobo/v2", -1},
		{"grobo v2 is not less than grobo v1", "grobo/v2", "grobo/v1", 1},
		// Prerelease tests
		{"Alpha comes before beta", "1alpha", "1beta", -1},
		{"Alpha comes before aardvark", "1alpha", "1aardvark", -1},
		{"Alpha comes before GA", "1alpha", "1", -1},
		{"Beta comes before preview", "1beta", "1preview", -1},
		{"Beta comes before baboon", "1beta", "1baboon", -1},
		{"Beta comes before GA", "1beta", "1", -1},
		{"Preview comes before storage", "1preview", "1storage", -1},
		{"Preview comes before grape", "1preview", "1grape", -1},
		{"Preview comes before GA", "1preview", "1", -1},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
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

	cases := []struct {
		name     string
		left     string
		right    string
		expected int
	}{
		{"Number equal to self", "42", "42", 0},
		{"Leading zeros on left-hand value don't matter", "042", "42", 0},
		{"Leading zeros on right-hand value don't matter", "42", "042", 0},
		{"Longer numbers are greater", "142", "42", 1},
		{"Shorter numbers are lesser", "42", "942", -1},
		{"Really really long equal numbers are equal", "1099511627776", "1099511627776", 0},
		{"Really really long non-equal numbers are comparable", "10995116277763", "10995116277769", -1},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
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

	cases := []struct {
		name     string
		version  string
		expected bool
	}{
		{"Non-preview", "v20200201", false},
		{"Alpha", "v20200201alpha", true},
		{"Beta", "v20200201beta", true},
		{"Preview", "v20200201preview", true},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			isPreview := ContainsPreviewVersionLabel(c.version)
			g.Expect(isPreview).To(Equal(c.expected))
		})
	}
}
