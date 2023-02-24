/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_Networking_RouteTable_CRUD(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	routeTable := &network.RouteTable{
		ObjectMeta: tc.MakeObjectMeta("routetable"),
		Spec: network.RouteTable_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
		},
	}

	tc.CreateResourceAndWait(routeTable)

	// It should be created in Kubernetes
	g.Expect(routeTable.Status.Id).ToNot(BeNil())
	armId := *routeTable.Status.Id

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Routes CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				Routes_CRUD(tc, routeTable)
			},
		},
	)

	tc.DeleteResourceAndWait(routeTable)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(network.APIVersion_Value))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(retryAfter).To(BeZero())
	g.Expect(exists).To(BeFalse())
}

func Routes_CRUD(tc *testcommon.KubePerTestContext, routeTable *network.RouteTable) {
	nextHopType := network.RouteNextHopType_VirtualAppliance
	ipv6Route := &network.RouteTablesRoute{
		ObjectMeta: tc.MakeObjectMeta("ipv6route"),
		Spec: network.RouteTables_Route_Spec{
			Owner:            testcommon.AsOwner(routeTable),
			AddressPrefix:    to.StringPtr("cab:cab::/96"),
			NextHopType:      &nextHopType,
			NextHopIpAddress: to.StringPtr("ace:cab:deca:f00d::1"),
		},
	}

	ipv4Route := &network.RouteTablesRoute{
		ObjectMeta: tc.MakeObjectMeta("ipv4route"),
		Spec: network.RouteTables_Route_Spec{
			Owner:            testcommon.AsOwner(routeTable),
			AddressPrefix:    to.StringPtr("Storage"),
			NextHopType:      &nextHopType,
			NextHopIpAddress: to.StringPtr("10.0.100.4"),
		},
	}

	tc.CreateResourcesAndWait(ipv4Route, ipv6Route)

	tc.Expect(ipv4Route.Status.Id).ToNot(BeNil())
	armId := *ipv4Route.Status.Id
	tc.Expect(ipv4Route.Status.AddressPrefix).ToNot(BeNil())
	tc.Expect(*ipv4Route.Status.AddressPrefix).To(Equal("Storage"))
	tc.Expect(ipv6Route.Status.Id).ToNot(BeNil())
	tc.Expect(ipv6Route.Status.AddressPrefix).ToNot(BeNil())
	tc.Expect(*ipv6Route.Status.AddressPrefix).To(Equal("cab:cab::/96"))

	// Update the subnet
	old := ipv4Route.DeepCopy()
	ipv4Route.Spec.NextHopIpAddress = to.StringPtr("10.0.100.5")
	tc.Patch(old, ipv4Route)

	objectKey := client.ObjectKeyFromObject(ipv4Route)

	// ensure state got updated in Azure
	tc.Eventually(func() *string {
		updated := &network.RouteTablesRoute{}
		tc.GetResource(objectKey, updated)
		return updated.Status.NextHopIpAddress
	}).Should(Equal(to.StringPtr("10.0.100.5")))

	tc.DeleteResourcesAndWait(ipv4Route, ipv6Route)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func Test_Networking_Route_CreatedThenRouteTableUpdated_RouteStillExists(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	routeTable := &network.RouteTable{
		ObjectMeta: tc.MakeObjectMeta("routetable"),
		Spec: network.RouteTable_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
		},
	}

	nextHopType := network.RouteNextHopType_VirtualAppliance
	ipv4Route := &network.RouteTablesRoute{
		ObjectMeta: tc.MakeObjectMeta("ipv4route"),
		Spec: network.RouteTables_Route_Spec{
			Owner:            testcommon.AsOwner(routeTable),
			AddressPrefix:    to.StringPtr("Storage"),
			NextHopType:      &nextHopType,
			NextHopIpAddress: to.StringPtr("10.0.100.4"),
		},
	}

	tc.CreateResourceAndWait(routeTable)
	tc.CreateResourcesAndWait(ipv4Route)

	// It should be created in Kubernetes
	tc.Expect(ipv4Route.Status.Id).ToNot(BeNil())
	armId := *ipv4Route.Status.Id

	// Now update the RouteTable
	old := routeTable.DeepCopy()
	routeTable.Spec.Tags = map[string]string{
		"taters": "boil 'em, mash 'em, stick 'em in a stew",
	}
	tc.PatchResourceAndWait(old, routeTable)

	// Now ensure that the Route still exists
	exists, _, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeTrue())
}
