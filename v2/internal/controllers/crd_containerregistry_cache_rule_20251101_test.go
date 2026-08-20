/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	registry "github.com/Azure/azure-service-operator/v2/api/containerregistry/v1api20230701"
	containerregistry "github.com/Azure/azure-service-operator/v2/api/containerregistry/v20251101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_ContainerRegistry_CacheRule_20251101_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	registryName := tc.NoSpaceNamer.GenerateName("registry")
	zoneRedundancy := registry.RegistryProperties_ZoneRedundancy_Disabled
	acr := &registry.Registry{
		ObjectMeta: tc.MakeObjectMetaWithName(registryName),
		Spec: registry.Registry_Spec{
			AzureName: registryName,
			Location:  tc.AzureRegion,
			Owner:     testcommon.AsOwner(rg),
			Sku: &registry.Sku{
				Name: to.Ptr(registry.Sku_Name_Premium),
			},
			ZoneRedundancy: &zoneRedundancy,
		},
	}

	cacheRule := &containerregistry.RegistryCacheRule{
		ObjectMeta: tc.MakeObjectMeta("cache-rule"),
		Spec: containerregistry.RegistryCacheRule_Spec{
			Owner:            testcommon.AsOwner(acr),
			SourceRepository: to.Ptr("mcr.microsoft.com/azuredocs/aci-helloworld"),
			TargetRepository: to.Ptr("cached-mcr/aci-helloworld"),
		},
	}

	tc.CreateResourcesAndWait(acr, cacheRule)
	defer tc.DeleteResourcesAndWait(cacheRule, acr)

	tc.Expect(cacheRule.Status.Id).ToNot(BeNil())
}
