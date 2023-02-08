/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package azureresource

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"testing"

	. "github.com/onsi/gomega"
)

func TestParseGroupKind(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name          string
		url           string
		expectedGroup string
		expectedKind  string
	}{
		{
			name:          "resource group",
			url:           "https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg",
			expectedGroup: "resources.azure.com",
			expectedKind:  "ResourceGroup",
		},
		{
			name:          "storage account",
			url:           "https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg/providers/Microsoft.Storage/storageAccounts/asostorage",
			expectedGroup: "storage.azure.com",
			expectedKind:  "StorageAccount",
		},
		{
			name:          "cosmos account",
			url:           "https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg/providers/Microsoft.DocumentDB/databaseAccounts/aso-cosmos",
			expectedGroup: "documentdb.azure.com",
			expectedKind:  "DatabaseAccount",
		},
		{
			name: "VMSS Scale Set	",
			url:           "https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg/providers/Microsoft.Compute/virtualMachineScaleSets/aso-scaleset",
			expectedGroup: "compute.azure.com",
			expectedKind:  "VirtualMachineScaleSet",
		},
		{
			name:          "VMSS Scale Set",
			url:           "https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg/providers/Microsoft.ContainerService/managedClusters/aso-cluster",
			expectedGroup: "containerservice.azure.com",
			expectedKind:  "ManagedCluster",
		},
		//{
		//	"VMSS Instance",
		//	"https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg/providers/Microsoft.Compute/virtualMachineScaleSets/aso-scaleset/virtualMachines/0",
		//	"storage",
		//	"storageaccount",
		//},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			c := c
			t.Parallel()

			g := NewGomegaWithT(t)

			importer := NewImporter()
			gk, err := importer.parseGroupKind(c.url)
			g.Expect(err).To(BeNil())
			g.Expect(gk.Group).To(Equal(c.expectedGroup))
			g.Expect(gk.Kind).To(Equal(c.expectedKind))
		})
	}
}

func TestSelectVersionKind(t *testing.T) {
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

	/*
	 * This test is a little brittle because the expected versions given above are going to change when new versions
	 * of those resources are released. Not sure of a better way to test this, so leaving this reminder for anyone
	 * who comes along and wonders why this test is failing.
	 */

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			g := NewGomegaWithT(t)

			importer := NewImporter()
			gk := schema.GroupKind{Group: c.group, Kind: c.kind}
			gvk := importer.selectVersion(gk, c.knownVersions)
			g.Expect(gvk.Group).To(Equal(gk.Group))
			g.Expect(gvk.Kind).To(Equal(gk.Kind))
			g.Expect(gvk.Version).To(Equal(c.expectedVersion))
		})
	}
}
