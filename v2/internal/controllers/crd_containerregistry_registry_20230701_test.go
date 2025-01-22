/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	containerregistry "github.com/Azure/azure-service-operator/v2/api/containerregistry/v1api20230701"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_ContainerRegistry_Registry_20230701_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create a ContainerRegistry

	name := tc.NoSpaceNamer.GenerateName("registry")
	zoneRedundancy := containerregistry.RegistryProperties_ZoneRedundancy_Disabled
	registry := &containerregistry.Registry{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: containerregistry.Registry_Spec{
			AdminUserEnabled:    to.Ptr(false),
			AzureName:           name,
			Location:            tc.AzureRegion,
			Owner:               testcommon.AsOwner(rg),
			PublicNetworkAccess: to.Ptr(containerregistry.RegistryProperties_PublicNetworkAccess_Enabled),
			Sku: &containerregistry.Sku{
				Name: to.Ptr(containerregistry.Sku_Name_Premium),
			},
			ZoneRedundancy: &zoneRedundancy,
		},
	}

	// Create a Replication

	replication := &containerregistry.RegistryReplication{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: containerregistry.RegistryReplication_Spec{
			Location:              to.Ptr("australiaeast"), // Has to be different from tc.AzureRegion
			Owner:                 testcommon.AsOwner(registry),
			RegionEndpointEnabled: to.Ptr(false),
			ZoneRedundancy:        to.Ptr(containerregistry.ReplicationProperties_ZoneRedundancy_Disabled),
		},
	}

	tc.CreateResourcesAndWait(registry, replication)
	defer tc.DeleteResourcesAndWait(registry, replication)

	// Perform some assertions on the resources we just created
	tc.Expect(registry.Status.Id).ToNot(BeNil())

	// Update the registry to ensure that works
	tc.T.Log("updating tags on registry")
	old := registry.DeepCopy()
	registry.Spec.Tags = map[string]string{
		"scratchcard": "lanyard",
	}

	tc.PatchResourceAndWait(old, registry)
	tc.Expect(registry.Status.Tags).To(HaveKey("scratchcard"))

	// Update the replication to make sure that works

	tc.T.Log("updating tags on replication")
	oldReplication := replication.DeepCopy()
	replication.Spec.Tags = map[string]string{
		"scratchcard": "lanyard",
	}

	tc.PatchResourceAndWait(oldReplication, replication)
	tc.Expect(replication.Status.Tags).To(HaveKey("scratchcard"))
}
