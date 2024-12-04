/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importresources

import (
	"testing"

	. "github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/Azure/azure-service-operator/v2/api"
)

func Test_selectVersionFromGK_givenGK_returnsExpectedVersion(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		group           string
		kind            string
		expectedVersion string
		expectedError   string
	}{
		"Batch Account": {
			group:           "batch.azure.com",
			kind:            "BatchAccount",
			expectedVersion: "v1api20210101",
		},
		"Managed Cluster": {
			group:           "containerservice.azure.com",
			kind:            "ManagedCluster",
			expectedVersion: "v1api20240901",
		},
		"Coffee isn't supported": {
			group:         "coffee.azure.com",
			kind:          "Latte",
			expectedError: "no known versions for Group coffee.azure.com, Kind Latte",
		},
	}

	factory := newImportFactory(api.CreateScheme())

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			gk := schema.GroupKind{
				Group: c.group,
				Kind:  c.kind,
			}

			gvk, err := factory.selectVersionFromGK(gk)

			// This check will start failing if/when a newer version of the resource being tested is added.
			// This is expected, just update the test to reflect the new version.
			if c.expectedVersion != "" {
				g.Expect(err).ToNot(HaveOccurred())
				g.Expect(gvk.Version).To(Equal(c.expectedVersion))
			}

			if c.expectedError != "" {
				g.Expect(err).To(HaveOccurred())
				g.Expect(err.Error()).To(ContainSubstring(c.expectedError))
			}
		})
	}
}

func Test_createBlankObjectFromGVK_GivenGVK_returnsExpectedInstance(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		group   string
		kind    string
		version string
	}{
		"Batch Account": {
			group:   "batch.azure.com",
			kind:    "BatchAccount",
			version: "v1api20210101",
		},
		"Managed Cluster": {
			group:   "containerservice.azure.com",
			kind:    "ManagedCluster",
			version: "v1api20240901",
		},
	}

	factory := newImportFactory(api.CreateScheme())

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			gvk := schema.GroupVersionKind{
				Group:   c.group,
				Kind:    c.kind,
				Version: c.version,
			}

			obj, err := factory.createBlankObjectFromGVK(gvk)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(obj.GetObjectKind().GroupVersionKind()).To(Equal(gvk))
		})
	}
}
