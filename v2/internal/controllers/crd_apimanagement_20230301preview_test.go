/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	apim "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20230301preview"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_ApiManagement_20230301preview_CRUD(t *testing.T) {
	t.Parallel()

	if *isLive {
		t.Skip("skipping test in live mode as it takes a very long time to provision an APIM service (1+h)")
	}

	tc := globalTestContext.ForTest(t)

	// We don't want to delete the resource group at the end of the test as APIM
	// takes a long time to provision. We'll clean it up manually
	rg := tc.NewTestResourceGroup()
	tc.CreateResourceAndWaitWithoutCleanup(rg)

	skuName := apim.ApiManagementServiceSkuProperties_Name("BasicV2")

	// There will be a New v2 SKU released 5/10/2023 which will have a much quicker start up
	// time. Move to that when it's available (BasicV2 or StandardV2 SKU)
	sku := apim.ApiManagementServiceSkuProperties{
		Capacity: to.Ptr(1),
		Name:     to.Ptr(skuName),
	}

	// Create an APIM instance. APIM has a soft delete feature; if you find that you
	// hit this problem add the `restore`` back in to resurrect it
	service := apim.Service{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("apim")),
		Spec: apim.Service_Spec{
			Location:       tc.AzureRegion,
			Owner:          testcommon.AsOwner(rg),
			PublisherEmail: to.Ptr("ASO@testing.com"),
			PublisherName:  to.Ptr("ASOTesting"),
			Sku:            &sku,
			// Restore:        to.Ptr(true),
		},
	}

	// TODO: When you are debugging, you can use this to create the APIM service once and not delete it
	tc.CreateResourceAndWaitWithoutCleanup(&service)
	// tc.CreateResourceAndWait(&service)

	tc.Expect(service.Status.Id).ToNot(BeNil())

	// Update the service to ensure that works
	tc.T.Log("updating tags on apim")
	old := service.DeepCopy()
	service.Spec.Tags = map[string]string{"scratchcard": "lanyard"}
	tc.PatchResourceAndWait(old, &service)
	tc.Expect(service.Status.Tags).To(HaveKey("scratchcard"))
}