/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	containerregistry "github.com/Azure/azure-service-operator/v2/api/containerregistry/v1alpha1api20210901"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_ContainerRegistry_Registry_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	publicNetworkAccess := containerregistry.RegistryPropertiesPublicNetworkAccessEnabled
	zoneRedundancy := containerregistry.RegistryPropertiesZoneRedundancyDisabled
	adminUserEnabled := false
	name := tc.NoSpaceNamer.GenerateName("registry")

	// Create a ContainerRegistry
	acct := containerregistry.Registry{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: containerregistry.Registries_Spec{
			AdminUserEnabled:    &adminUserEnabled,
			AzureName:           name,
			Location:            tc.AzureRegion,
			Owner:               testcommon.AsOwner(rg),
			PublicNetworkAccess: &publicNetworkAccess,
			Sku: containerregistry.Sku{
				Name: containerregistry.SkuNameBasic,
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
