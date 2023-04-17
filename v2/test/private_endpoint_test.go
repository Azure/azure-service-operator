/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package test

import (
	"testing"

	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	networkvnet "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20220701"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20210401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// This test cannot be run in record/replay mode because the state it looks for at the beginning (Ready = false with warning)
// is not "stable" (the reconciler keeps retrying). Since it keeps retrying there isn't a deterministic number of
// retry attempts it makes which means a recording test may run out of recorded retries.
func Test_Networking_PrivateEndpoint_WithoutAutoApproval_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	saKind := storage.StorageAccount_Kind_Spec_BlobStorage
	sa := newStorageAccount(tc, rg)
	sa.Spec.Kind = &saKind

	vnet := newVMVirtualNetwork(tc, testcommon.AsOwner(rg))
	subnet := newVMSubnet(tc, testcommon.AsOwner(vnet))
	endpoint := &network.PrivateEndpoint{
		ObjectMeta: tc.MakeObjectMeta("endpoint"),
		Spec: network.PrivateEndpoint_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			ManualPrivateLinkServiceConnections: []network.PrivateLinkServiceConnection{
				{
					Name:                        to.Ptr("testEndpoint"),
					PrivateLinkServiceReference: tc.MakeReferenceFromResource(sa),
					GroupIds:                    []string{"blob"},
				},
			},
			Subnet: &network.Subnet_PrivateEndpoint_SubResourceEmbedded{
				Reference: tc.MakeReferenceFromResource(subnet),
			},
		},
	}

	tc.CreateResourcesAndWait(sa, vnet, subnet)
	tc.Expect(sa.Status.Id).ToNot(BeNil())

	tc.CreateResourceAndWaitForState(endpoint, metav1.ConditionFalse, conditions.ConditionSeverityWarning)
	tc.Expect(endpoint.Status.Id).ToNot(BeNil())
	armId := *endpoint.Status.Id

	tc.DeleteResourceAndWait(endpoint)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func newVMVirtualNetwork(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference) *networkvnet.VirtualNetwork {
	return &networkvnet.VirtualNetwork{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vn")),
		Spec: networkvnet.VirtualNetwork_Spec{
			Owner:    owner,
			Location: tc.AzureRegion,
			AddressSpace: &networkvnet.AddressSpace{
				AddressPrefixes: []string{"10.0.0.0/16"},
			},
		},
	}
}

func newVMSubnet(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference) *networkvnet.VirtualNetworksSubnet {
	return &networkvnet.VirtualNetworksSubnet{
		ObjectMeta: tc.MakeObjectMeta("subnet"),
		Spec: networkvnet.VirtualNetworks_Subnet_Spec{
			Owner:         owner,
			AddressPrefix: to.Ptr("10.0.0.0/24"),
		},
	}
}
