/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"testing"

	. "github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

func Test_ImportableResource_SelectLatestVersion(t *testing.T) {
	t.Parallel()

	createKnownVersions := func(group string, versions ...string) []schema.GroupVersion {
		result := make([]schema.GroupVersion, 0, len(versions))
		for _, v := range versions {
			gv := schema.GroupVersion{Group: group, Version: v}
			result = append(result, gv)
		}

		return result
	}

	cases := []struct {
		name            string
		group           string
		kind            string
		knownVersions   []schema.GroupVersion
		expectedVersion string
	}{
		{
			name:            "Single available version, get that version",
			group:           "resources.azure.com",
			kind:            "ResourceGroup",
			knownVersions:   createKnownVersions("resources.azure.com", "v1beta20200601"),
			expectedVersion: "v1beta20200601",
		},
		{
			name:            "Multiple versions available, get the latest",
			group:           "storage.azure.com",
			kind:            "StorageAccount",
			knownVersions:   createKnownVersions("storage.azure.com", "v1beta20200601", "v1beta20220101"),
			expectedVersion: "v1beta20220101",
		},
		{
			name:            "Latest version is preview, get latest stable",
			group:           "containerservice.azure.com",
			kind:            "ManagedCluster",
			knownVersions:   createKnownVersions("containerservice.azure.com", "v1beta20200601", "v1beta20220101preview"),
			expectedVersion: "v1beta20200601",
		},
		{
			name:            "All versions are preview, get latest preview",
			group:           "dbforpostgresql.azure.com",
			kind:            "FlexibleServer",
			knownVersions:   createKnownVersions("dbforpostgresql.azure.com", "v1beta20200601preview", "v1beta20220101preview"),
			expectedVersion: "v1beta20220101preview",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			g := NewGomegaWithT(t)

			importer := importableResource{}
			
			gk := schema.GroupKind{Group: c.group, Kind: c.kind}
			gvk := importer.selectLatestVersion(gk, c.knownVersions)
			g.Expect(gvk.Group).To(Equal(gk.Group))
			g.Expect(gvk.Kind).To(Equal(gk.Kind))
			g.Expect(gvk.Version).To(Equal(c.expectedVersion))
		})
	}
}
