/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	managedidentity2018 "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20181130"
	search "github.com/Azure/azure-service-operator/v2/api/search/v1api20210401preview"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Search_SearchService_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	mi := &managedidentity2018.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity2018.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
		},
	}

	tc.CreateResourceAndWait(mi)
	tc.Expect(mi.Status.Id).ToNot(BeNil())

	service := &search.SearchService{
		ObjectMeta: tc.MakeObjectMeta("search"),
		Spec: search.SearchService_Spec{
			EncryptionWithCmk: &search.EncryptionWithCmk{
				Enforcement: to.Ptr(search.EncryptionWithCmk_Enforcement_Disabled),
			},
			HostingMode: to.Ptr(search.SearchServiceProperties_HostingMode_Default),
			Identity: &search.Identity{
				Type: to.Ptr(search.Identity_Type_SystemAssignedUserAssigned),
				UserAssignedIdentities: []search.UserAssignedIdentityDetails{
					{
						Reference: genruntime.ResourceReference{
							ARMID: *mi.Status.Id,
						},
					},
				},
			},
			Location: tc.AzureRegion,
			NetworkRuleSet: &search.NetworkRuleSet{
				Bypass: to.Ptr(search.NetworkRuleSet_Bypass_None),
			},
			Owner:               testcommon.AsOwner(rg),
			PartitionCount:      to.Ptr(1),
			PublicNetworkAccess: to.Ptr(search.SearchServiceProperties_PublicNetworkAccess_Disabled),
			ReplicaCount:        to.Ptr(1),
			SemanticSearch:      to.Ptr(search.SemanticSearch_Free),
			Sku:                 &search.Sku{Name: to.Ptr(search.Sku_Name_Standard)},
		},
	}

	tc.CreateResourceAndWait(service)
	tc.Expect(service.Status.Id).ToNot(BeNil())
	armId := *service.Status.Id

	old := service.DeepCopy()
	key := "foo"
	service.Spec.Tags = map[string]string{key: "bar"}

	tc.PatchResourceAndWait(old, service)
	tc.Expect(service.Status.Tags).To(HaveKey(key))

	tc.DeleteResourceAndWait(service)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(search.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
