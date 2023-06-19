/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	search "github.com/Azure/azure-service-operator/v2/api/search/v1api20220901"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Search_SearchService_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	service := &search.SearchService{
		ObjectMeta: tc.MakeObjectMeta("search"),
		Spec: search.SearchService_Spec{
			EncryptionWithCmk: &search.EncryptionWithCmk{
				Enforcement: to.Ptr(search.EncryptionWithCmk_Enforcement_Disabled),
			},
			HostingMode: to.Ptr(search.SearchServiceProperties_HostingMode_Default),
			Identity: &search.Identity{
				Type: to.Ptr(search.Identity_Type_SystemAssigned),
			},
			Location:            tc.AzureRegion,
			Owner:               testcommon.AsOwner(rg),
			PartitionCount:      to.Ptr(1),
			PublicNetworkAccess: to.Ptr(search.SearchServiceProperties_PublicNetworkAccess_Disabled),
			ReplicaCount:        to.Ptr(1),
			Sku:                 &search.Sku{Name: to.Ptr(search.Sku_Name_Standard)},
		},
	}

	tc.CreateResourceAndWait(service)
	tc.Expect(service.Status.Id).ToNot(BeNil())
	armId := *service.Status.Id

	tc.RunSubtests(
		testcommon.Subtest{
			Name: "WriteSearchServiceSecrets",
			Test: func(tc *testcommon.KubePerTestContext) {
				SearchService_WriteSecrets(tc, service)
			},
		})

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

func SearchService_WriteSecrets(tc *testcommon.KubePerTestContext, service *search.SearchService) {
	old := service.DeepCopy()
	searchKeysSecret := "searchkeyssecret"
	service.Spec.OperatorSpec = &search.SearchServiceOperatorSpec{
		Secrets: &search.SearchServiceOperatorSecrets{
			AdminPrimaryKey:   &genruntime.SecretDestination{Name: searchKeysSecret, Key: "adminPrimaryKey"},
			AdminSecondaryKey: &genruntime.SecretDestination{Name: searchKeysSecret, Key: "adminSecondaryKey"},
			QueryKey:          &genruntime.SecretDestination{Name: searchKeysSecret, Key: "queryKey"},
		},
	}

	tc.PatchResourceAndWait(old, service)
	tc.ExpectSecretHasKeys(
		searchKeysSecret,
		"adminPrimaryKey",
		"adminSecondaryKey",
		"queryKey")
}
