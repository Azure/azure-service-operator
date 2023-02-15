/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package importing

import (
	"github.com/Azure/azure-service-operator/v2/api"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"testing"
)

func Test_ResourceImporterFactory_GroupKindFromArmId(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name          string
		armId         string
		expectedGroup string
		expectedKind  string
	}{
		{
			name:          "resource group",
			armId:         "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg",
			expectedGroup: "resources.azure.com",
			expectedKind:  "ResourceGroup",
		},
		{
			name:          "storage account",
			armId:         "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg/providers/Microsoft.Storage/storageAccounts/asostorage",
			expectedGroup: "storage.azure.com",
			expectedKind:  "StorageAccount",
		},
		{
			name:          "cosmos account",
			armId:         "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg/providers/Microsoft.DocumentDB/databaseAccounts/aso-cosmos",
			expectedGroup: "documentdb.azure.com",
			expectedKind:  "DatabaseAccount",
		},
		{
			name: "VMSS Scale Set	",
			armId:         "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg/providers/Microsoft.Compute/virtualMachineScaleSets/aso-scaleset",
			expectedGroup: "compute.azure.com",
			expectedKind:  "VirtualMachineScaleSet",
		},
		{
			name:          "VMSS Scale Set",
			armId:         "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg/providers/Microsoft.ContainerService/managedClusters/aso-cluster",
			expectedGroup: "containerservice.azure.com",
			expectedKind:  "ManagedCluster",
		},
		//{
		//	"VMSS Instance",
		//	"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg/providers/Microsoft.Compute/virtualMachineScaleSets/aso-scaleset/virtualMachines/0",
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

			factory := newResourceImporterFactory(nil)
			gk, err := factory.groupKindFromArmId(c.armId)
			g.Expect(err).To(BeNil())
			g.Expect(gk.Group).To(Equal(c.expectedGroup))
			g.Expect(gk.Kind).To(Equal(c.expectedKind))
		})
	}
}

func Test_ResourceImporterFactory_SelectLatestVersion(t *testing.T) {
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

			importer := newResourceImporterFactory(nil)
			gk := schema.GroupKind{Group: c.group, Kind: c.kind}
			gvk := importer.selectLatestVersion(gk, c.knownVersions)
			g.Expect(gvk.Group).To(Equal(gk.Group))
			g.Expect(gvk.Kind).To(Equal(gk.Kind))
			g.Expect(gvk.Version).To(Equal(c.expectedVersion))
		})
	}
}

func Test_ResourceImporterFactory_GroupVersionKindFromArmId(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name            string
		armId           string
		expectedGroup   string
		expectedKind    string
		expectedVersion string
	}{
		{
			name:            "resource group",
			armId:           "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg",
			expectedGroup:   "resources.azure.com",
			expectedKind:    "ResourceGroup",
			expectedVersion: "v1beta20200601",
		},
		{
			name:            "storage account",
			armId:           "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg/providers/Microsoft.Storage/storageAccounts/aso-storage",
			expectedGroup:   "storage.azure.com",
			expectedKind:    "StorageAccount",
			expectedVersion: "v1beta20210401",
		},
		{
			name:            "managed cluster",
			armId:           "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg/providers/Microsoft.ContainerService/managedClusters/aso-cluster",
			expectedGroup:   "containerservice.azure.com",
			expectedKind:    "ManagedCluster",
			expectedVersion: "v1beta20210501",
		},
	}

	factory := newResourceImporterFactory(api.CreateScheme())
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			g := NewGomegaWithT(t)

			gvk, err := factory.groupVersionKindFromArmId(c.armId)
			g.Expect(err).To(BeNil())
			g.Expect(gvk.Group).To(Equal(c.expectedGroup))
			g.Expect(gvk.Kind).To(Equal(c.expectedKind))
			g.Expect(gvk.Version).To(Equal(c.expectedVersion))
		})
	}
}

func Test_ResourceImporterFactory_CreateForArmId(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name          string
		armId         string
		expectedError string
	}{
		{
			name:          "resource group",
			armId:         "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg",
			expectedError: "",
		},
		{
			name:          "storage account",
			armId:         "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg/providers/Microsoft.Storage/storageAccounts/aso-storage",
			expectedError: "",
		},
		{
			name:          "managed cluster",
			armId:         "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-rg/providers/Microsoft.ContainerService/managedClusters/aso-cluster",
			expectedError: "",
		},
		{
			name:          "invalid arm id",
			armId:         "invalid",
			expectedError: "invalid resource ID",
		},
	}

	factory := newResourceImporterFactory(api.CreateScheme())
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			g := NewGomegaWithT(t)

			_, err := factory.CreateForArmId(c.armId)
			if c.expectedError == "" {
				g.Expect(err).To(BeNil())
			} else {
				g.Expect(err).ToNot(BeNil())
				g.Expect(err.Error()).To(ContainSubstring(c.expectedError))
			}
		})
	}
}
