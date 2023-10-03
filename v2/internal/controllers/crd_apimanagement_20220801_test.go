/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	apim "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_ApiManagement_20220801_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	sku := apim.ApiManagementServiceSkuProperties{
		Capacity: to.Ptr(1),
		Name:     to.Ptr(apim.ApiManagementServiceSkuProperties_Name_Developer),
	}

	// Create an APIM instance
	service := apim.Service{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("apim")),
		Spec: apim.Service_Spec{
			Location:       tc.AzureRegion,
			Owner:          testcommon.AsOwner(rg),
			PublisherEmail: to.Ptr("ASO@testing.com"),
			PublisherName:  to.Ptr("ASOTesting"),
			Sku:            &sku,
		},
	}

	tc.CreateResourceAndWait(&service)

	tc.Expect(service.Status.Id).ToNot(BeNil())

	tc.Expect(service.Status.Id).ToNot(BeNil())

	// Update the service to ensure that works
	tc.T.Log("updating tags on apim")
	old := service.DeepCopy()
	service.Spec.Tags = map[string]string{"scratchcard": "lanyard"}
	tc.PatchResourceAndWait(old, &service)
	tc.Expect(service.Status.Tags).To(HaveKey("scratchcard"))
}
