/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	network20201101 "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20220701"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20210401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// Example is from https://learn.microsoft.com/en-us/azure/private-link/create-private-endpoint-bicep?tabs=CLI#review-the-bicep-file
// We've replaced SQLServers with StorageAccount in the test.
func Test_Networking_PrivateEndpoint_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	saKind := storage.StorageAccount_Kind_Spec_BlobStorage
	sa := newStorageAccount(tc, rg)
	sa.Spec.Kind = &saKind

	vnet := newVMVirtualNetwork(tc, testcommon.AsOwner(rg))
	subnet := newVMSubnet(tc, testcommon.AsOwner(vnet))
	endpoint := newPrivateEndpoint(tc, rg, sa, subnet)

	tc.CreateResourcesAndWait(sa, vnet, subnet, endpoint)
	tc.Expect(sa.Status.Id).ToNot(BeNil())
	tc.Expect(endpoint.Status.Id).ToNot(BeNil())
	armId := *endpoint.Status.Id

	old := endpoint.DeepCopy()
	key := "foo"
	endpoint.Spec.Tags = map[string]string{key: "bar"}

	tc.PatchResourceAndWait(old, endpoint)
	tc.Expect(endpoint.Status.Tags).To(HaveKey(key))

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Test_DNSZoneGroup_CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				PrivateEndpoint_DNSZoneGroup_CRUD(tc, vnet, endpoint, rg)
			},
		},
	)

	tc.DeleteResourceAndWait(endpoint)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func Test_Networking_PrivateEndpoint_WithoutAutoApproval_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	saKind := storage.StorageAccount_Kind_Spec_BlobStorage
	sa := newStorageAccount(tc, rg)
	sa.Spec.Kind = &saKind

	vnet := newVMVirtualNetwork(tc, testcommon.AsOwner(rg))
	subnet := newVMSubnet(tc, testcommon.AsOwner(vnet))
	endpoint := newPrivateEndpoint(tc, rg, sa, subnet)

	// We need to use ManualPrivateLinkServiceConnection here which is not auto-approved.
	endpoint.Spec.PrivateLinkServiceConnections = []network.PrivateLinkServiceConnection{}
	endpoint.Spec.ManualPrivateLinkServiceConnections = []network.PrivateLinkServiceConnection{
		{
			Name:                        to.Ptr("testEndpoint"),
			PrivateLinkServiceReference: tc.MakeReferenceFromResource(sa),
			GroupIds:                    []string{"blob"},
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

func PrivateEndpoint_DNSZoneGroup_CRUD(tc *testcommon.KubePerTestContext, vnet *network20201101.VirtualNetwork, endpoint *network.PrivateEndpoint, rg *resources.ResourceGroup) {
	zone := newPrivateDNSZone(tc, "privatelink.blob.core.windows.net", rg)
	vnetLink := newVirtualNetworkLink(tc, zone, vnet)

	tc.CreateResourcesAndWait(zone, vnetLink)

	dnsZoneGroup := &network.PrivateEndpointsPrivateDnsZoneGroup{
		ObjectMeta: tc.MakeObjectMeta("dnszonegroup"),
		Spec: network.PrivateEndpoints_PrivateDnsZoneGroup_Spec{
			Owner: testcommon.AsOwner(endpoint),
			PrivateDnsZoneConfigs: []network.PrivateDnsZoneConfig{
				{
					Name:                    to.Ptr("config"),
					PrivateDnsZoneReference: tc.MakeReferenceFromResource(zone),
				},
			},
		},
	}

	tc.CreateResourceAndWait(dnsZoneGroup)

	tc.Expect(dnsZoneGroup.Status.Id).ToNot(BeNil())

	tc.DeleteResource(dnsZoneGroup)
}

func newPrivateEndpoint(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, sa *storage.StorageAccount, subnet *network20201101.VirtualNetworksSubnet) *network.PrivateEndpoint {
	endpoint := &network.PrivateEndpoint{
		ObjectMeta: tc.MakeObjectMeta("endpoint"),
		Spec: network.PrivateEndpoint_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			PrivateLinkServiceConnections: []network.PrivateLinkServiceConnection{
				{
					Name:                        to.Ptr("testEndpoint"),
					PrivateLinkServiceReference: tc.MakeReferenceFromResource(sa),
					GroupIds:                    []string{"blob"}, // TODO: This is a bit weird that user has to figure out the group ID(s).
				},
			},
			Subnet: &network.Subnet_PrivateEndpoint_SubResourceEmbedded{
				Reference: tc.MakeReferenceFromResource(subnet),
			},
		},
	}
	return endpoint
}
