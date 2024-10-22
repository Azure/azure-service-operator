/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	insights "github.com/Azure/azure-service-operator/v2/api/insights/v1api20210501preview"
	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20210401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Insights_DiagnosticSetting_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	rg_name := rg.Status.Name

	// vnet := VMVirtualNetwork(tc, testcommon.AsOwner(rg))

	vnet := &network.VirtualNetwork{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vn")),
		Spec: network.VirtualNetwork_Spec{
			Owner:    testcommon.AsOwner(rg),
			Location: tc.AzureRegion,
			AddressSpace: &network.AddressSpace{
				AddressPrefixes: []string{"10.0.0.0/16"},
			},
		},
	}

	tc.CreateResourceAndWait(vnet)

	vnet_name := vnet.Status.Name

	accessTier := storage.StorageAccountPropertiesCreateParameters_AccessTier_Hot
	kind := storage.StorageAccount_Kind_Spec_StorageV2
	sku := storage.SkuName_Standard_LRS
	acct := &storage.StorageAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("stor")),
		Spec: storage.StorageAccount_Spec{
			Location:              tc.AzureRegion,
			AllowBlobPublicAccess: to.Ptr(false),
			Owner:                 testcommon.AsOwner(rg),
			Kind:                  &kind,
			Sku: &storage.Sku{
				Name: &sku,
			},
			// TODO: They mark this property as optional but actually it is required
			AccessTier: &accessTier,
		},
	}

	tc.CreateResourceAndWait(acct)

	subARMID := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/virtualNetworks/%s", tc.AzureSubscription, *rg_name, *vnet_name)

	vnet_diagnosticSettings := &insights.DiagnosticSetting{
		ObjectMeta: tc.MakeObjectMeta("diagnosticsetting"),
		Spec: insights.DiagnosticSetting_Spec{
			Owner: &genruntime.ArbitraryOwnerReference{
				ARMID: subARMID,
			},
			Logs: []insights.LogSettings{
				{
					CategoryGroup: to.Ptr("allLogs"),
					Enabled:       to.Ptr(true),
				},
			},
			StorageAccountReference: tc.MakeReferenceFromResource(acct),
		},
	}

	tc.CreateResourceAndWait(vnet_diagnosticSettings)

	tc.Expect(vnet_diagnosticSettings.Status.Id).ToNot(BeNil())
	armId := *vnet_diagnosticSettings.Status.Id

	tc.DeleteResourceAndWait(vnet_diagnosticSettings)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(insights.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
