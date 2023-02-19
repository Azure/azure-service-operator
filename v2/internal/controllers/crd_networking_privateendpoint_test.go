/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1beta20220701"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1beta20210401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_Networking_PrivateEndpoint_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	saKind := storage.StorageAccount_Kind_Spec_BlobStorage
	sa := newStorageAccount(tc, rg)
	sa.Spec.Kind = &saKind

	vnet := newVMVirtualNetwork(tc, testcommon.AsOwner(rg))
	subnet := newVMSubnet(tc, testcommon.AsOwner(vnet))

	tc.CreateResourcesAndWait(sa, vnet, subnet)

	tc.Expect(sa.Status.Id).ToNot(BeNil())

	endpoint := &network.PrivateEndpoint{
		ObjectMeta: tc.MakeObjectMeta("endpoint"),
		Spec: network.PrivateEndpoint_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			PrivateLinkServiceConnections: []network.PrivateLinkServiceConnection{
				{
					Name:                        to.StringPtr("testEndpoint"),
					PrivateLinkServiceReference: tc.MakeReferenceFromResource(sa),
					GroupIds:                    []string{"blob"}, // TODO: This is a bit weird that user has to figure out the group ID(s).
				},
			},
			Subnet: &network.Subnet_PrivateEndpoint_SubResourceEmbedded{
				Reference: tc.MakeReferenceFromResource(subnet),
			},
		},
	}

	tc.CreateResourceAndWait(endpoint)

	tc.Expect(endpoint.Status.Id).ToNot(BeNil())
	armId := *endpoint.Status.Id

	old := endpoint.DeepCopy()
	key := "foo"
	endpoint.Spec.Tags = map[string]string{key: "bar"}

	tc.PatchResourceAndWait(old, endpoint)
	tc.Expect(endpoint.Status.Tags).To(HaveKey(key))

	tc.DeleteResourceAndWait(endpoint)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())

}
