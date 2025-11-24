// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package crdmanagement_test

import (
	"testing"

	. "github.com/onsi/gomega"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	"github.com/Azure/azure-service-operator/v2/internal/crdmanagement"
)

func Test_GetDeprecatedStorageVersions(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name            string
		storedVersions  []string
		expectedKept    []string
		expectedRemoved []string
	}{
		{
			name:            "Single version, nothing deprecated",
			storedVersions:  []string{"v1api20240101"},
			expectedKept:    []string{"v1api20240101"},
			expectedRemoved: nil,
		},
		{
			name: "Three versions, newest is vYYYYMMDD-style",
			storedVersions: []string{
				"v1api20230101preview",
				"v1api20240101",
				"v20250101",
			},
			expectedKept:    []string{"v20250101"},
			expectedRemoved: []string{"v1api20230101preview", "v1api20240101"},
		},
		{
			name: "Three versions, newest is v1apiYYYYMMDD-style",
			storedVersions: []string{
				"v20230101preview",
				"v20240101preview",
				"v1api20250101",
			},
			expectedKept:    []string{"v1api20250101"},
			expectedRemoved: []string{"v20230101preview", "v20240101preview"},
		},
		{
			name: "Two versions, preview newer than non-preview",
			storedVersions: []string{
				"v1api20240101",
				"v1api20250101preview",
			},
			expectedKept:    []string{"v1api20240101"},
			expectedRemoved: []string{"v1api20250101preview"},
		},
		{
			name:            "Empty versions",
			storedVersions:  []string{},
			expectedKept:    []string{},
			expectedRemoved: nil,
		},
		{
			name: "Multiple versions same date different types",
			storedVersions: []string{
				"v20240101",
				"v1api20240101",
				"v20240101preview",
			},
			expectedKept:    []string{"v20240101"},
			expectedRemoved: []string{"v1api20240101", "v20240101preview"},
		},
		{
			name: "Versions with storage suffix",
			storedVersions: []string{
				"v1api20230101storage",
				"v1api20240101storage",
				"v1api20250101storage",
			},
			expectedKept:    []string{"v1api20250101storage"},
			expectedRemoved: []string{"v1api20230101storage", "v1api20240101storage"},
		},
		{
			name: "All preview versions, keep newest preview",
			storedVersions: []string{
				"v1api20230101preview",
				"v1api20240101preview",
				"v1api20250101preview",
			},
			expectedKept:    []string{"v1api20250101preview"},
			expectedRemoved: []string{"v1api20230101preview", "v1api20240101preview"},
		},
		{
			name: "Non-matching version pattern",
			storedVersions: []string{
				"v1beta1",
				"v1",
				"v2",
			},
			expectedKept:    []string{"v2"},
			expectedRemoved: []string{"v1beta1", "v1"},
		},
		{
			name: "Prefer v over v1api with same date",
			storedVersions: []string{
				"v1api20240101",
				"v20240101",
			},
			expectedKept:    []string{"v20240101"},
			expectedRemoved: []string{"v1api20240101"},
		},
		{
			name: "Prefer v over v1api with same date and storage suffix",
			storedVersions: []string{
				"v1api20240101storage",
				"v20240101storage",
				"v1api20230101storage",
			},
			expectedKept:    []string{"v20240101storage"},
			expectedRemoved: []string{"v1api20240101storage", "v1api20230101storage"},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			g := NewGomegaWithT(t)
			crd := makeBasicCRDWithStoredVersions("test", c.storedVersions)
			kept, removed := crdmanagement.GetDeprecatedStorageVersions(crd)
			g.Expect(kept).To(Equal(c.expectedKept))
			g.Expect(removed).To(Equal(c.expectedRemoved))
		})
	}
}

func Test_CompareVersions(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		a        string
		b        string
		expected int
	}{
		{
			name:     "Equal versions",
			a:        "v1api20240101",
			b:        "v1api20240101",
			expected: 0,
		},
		{
			name:     "Newer date wins (a > b)",
			a:        "v1api20250101",
			b:        "v1api20240101",
			expected: 1,
		},
		{
			name:     "Newer date wins (a < b)",
			a:        "v1api20240101",
			b:        "v1api20250101",
			expected: -1,
		},
		{
			name:     "Non-preview beats preview same date (a > b)",
			a:        "v1api20240101",
			b:        "v1api20240101preview",
			expected: 1,
		},
		{
			name:     "Non-preview beats preview same date (a < b)",
			a:        "v1api20240101preview",
			b:        "v1api20240101",
			expected: -1,
		},
		{
			name:     "Non-preview beats preview even if older date (a > b)",
			a:        "v1api20240101",
			b:        "v1api20250101preview",
			expected: 1,
		},
		{
			name:     "Non-preview beats preview even if older date (a < b)",
			a:        "v1api20250101preview",
			b:        "v1api20240101",
			expected: -1,
		},
		{
			name:     "Storage suffix ignored (equal)",
			a:        "v1api20240101storage",
			b:        "v1api20240101",
			expected: 0,
		},
		{
			name:     "Storage suffix ignored (a > b)",
			a:        "v1api20250101storage",
			b:        "v1api20240101storage",
			expected: 1,
		},
		{
			name:     "V without api (a > b)",
			a:        "v20250101",
			b:        "v20240101",
			expected: 1,
		},
		{
			name:     "V without api (a < b)",
			a:        "v20240101",
			b:        "v20250101",
			expected: -1,
		},
		{
			name:     "Mixed v and v1api different dates (a > b)",
			a:        "v1api20250101",
			b:        "v20240101",
			expected: 1,
		},
		{
			name:     "Mixed v and v1api different dates (a < b)",
			a:        "v20240101",
			b:        "v1api20250101",
			expected: -1,
		},
		{
			name:     "Prefer v over v1api same date (a < b)",
			a:        "v1api20240101",
			b:        "v20240101",
			expected: -1,
		},
		{
			name:     "Prefer v over v1api same date (a > b)",
			a:        "v20240101",
			b:        "v1api20240101",
			expected: 1,
		},
		{
			name:     "Prefer v over v1api same date with storage (a < b)",
			a:        "v1api20240101storage",
			b:        "v20240101storage",
			expected: -1,
		},
		{
			name:     "Prefer v over v1api same date with storage (a > b)",
			a:        "v20240101storage",
			b:        "v1api20240101storage",
			expected: 1,
		},
		{
			name:     "Prefer v over v1api same date both preview (a < b)",
			a:        "v1api20240101preview",
			b:        "v20240101preview",
			expected: -1,
		},
		{
			name:     "Prefer v over v1api same date both preview (a > b)",
			a:        "v20240101preview",
			b:        "v1api20240101preview",
			expected: 1,
		},
		{
			name:     "Preview versions (a > b)",
			a:        "v1api20250101preview",
			b:        "v1api20240101preview",
			expected: 1,
		},
		{
			name:     "Preview versions (a < b)",
			a:        "v1api20240101preview",
			b:        "v1api20250101preview",
			expected: -1,
		},
		{
			name:     "Invalid pattern falls back to string comparison (a > b)",
			a:        "v1beta1",
			b:        "v1alpha1",
			expected: 1,
		},
		{
			name:     "Invalid pattern falls back to string comparison (a < b)",
			a:        "v1",
			b:        "v2",
			expected: -1,
		},
		{
			name:     "One invalid pattern (a < b)",
			a:        "v1api20240101",
			b:        "v1beta1",
			expected: -1,
		},
		{
			name:     "One invalid pattern (a > b)",
			a:        "v1beta1",
			b:        "v1api20240101",
			expected: 1,
		},
		{
			name:     "Equal invalid patterns",
			a:        "v1beta1",
			b:        "v1beta1",
			expected: 0,
		},
		{
			name:     "Empty strings",
			a:        "",
			b:        "",
			expected: 0,
		},
		{
			name:     "Empty string (a > b)",
			a:        "v1api20240101",
			b:        "",
			expected: 1,
		},
		{
			name:     "Empty string (a < b)",
			a:        "",
			b:        "v1api20240101",
			expected: -1,
		},
		{
			name:     "Preview and storage mixed (a > b)",
			a:        "v1api20240101storage",
			b:        "v1api20250101previewstorage",
			expected: 1,
		},
		{
			name:     "Preview and storage mixed (a < b)",
			a:        "v1api20250101previewstorage",
			b:        "v1api20240101storage",
			expected: -1,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			g := NewGomegaWithT(t)
			result := crdmanagement.CompareVersions(c.a, c.b)
			g.Expect(result).To(Equal(c.expected))
		})
	}
}

/*
 * Helper functions
 */

func makeBasicCRDWithStoredVersions(name string, storedVersions []string) apiextensions.CustomResourceDefinition {
	crd := makeBasicCRD(name)
	crd.Status.StoredVersions = storedVersions
	return crd
}
