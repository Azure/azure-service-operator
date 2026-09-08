/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	servicelinker "github.com/Azure/azure-service-operator/v2/api/servicelinker/v20240401"
	web "github.com/Azure/azure-service-operator/v2/api/web/v1api20220301"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_ServiceLinker_Linker_CRUD_v20240401(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	tc.AzureRegion = to.Ptr("westus3")
	serverFarm := newServerFarm(tc, rg, *tc.AzureRegion)
	site := &web.Site{
		ObjectMeta: tc.MakeObjectMeta("site"),
		Spec: web.Site_Spec{
			Enabled:             to.Ptr(true),
			Location:            tc.AzureRegion,
			Owner:               testcommon.AsOwner(rg),
			ServerFarmReference: tc.MakeReferenceFromResource(serverFarm),
		},
	}
	account := newStorageAccount(tc, rg)
	linker := &servicelinker.Linker{
		ObjectMeta: tc.MakeObjectMeta("linker"),
		Spec: servicelinker.Linker_Spec{
			AzureName: tc.NoSpaceNamer.GenerateName("linker"),
			AuthInfo: &servicelinker.AuthInfoBase{
				SystemAssignedIdentity: &servicelinker.SystemAssignedIdentityAuthInfo{
					AuthType: to.Ptr(servicelinker.SystemAssignedIdentityAuthInfo_AuthType_SystemAssignedIdentity),
				},
			},
			ClientType: to.Ptr(servicelinker.ClientType_Dotnet),
			Owner:      tc.AsExtensionOwner(site),
			TargetService: &servicelinker.TargetServiceBase{
				AzureResource: &servicelinker.AzureResource{
					Reference: tc.MakeReferenceFromResource(account),
					Type:      to.Ptr(servicelinker.AzureResource_Type_AzureResource),
				},
			},
		},
	}

	tc.CreateResourcesAndWait(serverFarm, site, account, linker)

	tc.Expect(linker.Status.Id).ToNot(BeNil())
	tc.Expect(linker.Status.ClientType).To(Equal(to.Ptr(servicelinker.ClientType_STATUS_Dotnet)))
	tc.Expect(linker.Status.TargetService).ToNot(BeNil())
	tc.Expect(linker.Status.TargetService.AzureResource).ToNot(BeNil())

	armID := *linker.Status.Id
	old := linker.DeepCopy()
	linker.Spec.ClientType = to.Ptr(servicelinker.ClientType_None)
	tc.PatchResourceAndWait(old, linker)
	tc.Expect(linker.Status.ClientType).To(Equal(to.Ptr(servicelinker.ClientType_STATUS_None)))

	tc.DeleteResourceAndWait(linker)

	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armID,
		string(servicelinker.APIVersion_Value),
	)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
