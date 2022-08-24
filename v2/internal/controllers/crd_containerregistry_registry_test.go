/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	containerregistry "github.com/Azure/azure-service-operator/v2/api/containerregistry/v1beta20210901"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_ContainerRegistry_Registry_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	publicNetworkAccess := containerregistry.RegistryProperties_PublicNetworkAccess_Enabled
	zoneRedundancy := containerregistry.RegistryProperties_ZoneRedundancy_Disabled
	adminUserEnabled := false
	name := tc.NoSpaceNamer.GenerateName("registry")

	// Create a ContainerRegistry
	skuName := containerregistry.Sku_Name_Basic
	acct := containerregistry.Registry{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: containerregistry.Registries_Spec{
			AdminUserEnabled:    &adminUserEnabled,
			AzureName:           name,
			Location:            tc.AzureRegion,
			Owner:               testcommon.AsOwner(rg),
			PublicNetworkAccess: &publicNetworkAccess,
			Sku: &containerregistry.Sku{
				Name: &skuName,
			},
			ZoneRedundancy: &zoneRedundancy,
		},
	}

	tc.CreateResourcesAndWait(&acct)
	defer tc.DeleteResourcesAndWait(&acct)

	// Perform some assertions on the resources we just created
	tc.Expect(acct.Status.Id).ToNot(BeNil())

	// Update the account to ensure that works
	tc.T.Log("updating tags on account")
	old := acct.DeepCopy()
	acct.Spec.Tags = map[string]string{"scratchcard": "lanyard"}
	tc.PatchResourceAndWait(old, &acct)
	tc.Expect(acct.Status.Tags).To(HaveKey("scratchcard"))
}
