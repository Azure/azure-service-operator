/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package perf_test

import (
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/client"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

// ResourceGroupFactory returns a ResourceFactory that creates a single ResourceGroup per invocation.
func ResourceGroupFactory() ResourceFactory {
	return func(tc *testcommon.KubePerTestContext, index int) []client.Object {
		rg := &resources.ResourceGroup{
			ObjectMeta: tc.MakeObjectMeta(fmt.Sprintf("perfrg-%d", index)),
			Spec: resources.ResourceGroup_Spec{
				Location: tc.AzureRegion,
				Tags:     testcommon.CreateTestResourceGroupDefaultTags(),
			},
		}
		return []client.Object{rg}
	}
}

// VirtualNetworkFactory returns a ResourceFactory that creates a ResourceGroup and a VirtualNetwork
// owned by that ResourceGroup per invocation. The resources are returned in dependency order
// (ResourceGroup first, VirtualNetwork second).
func VirtualNetworkFactory() ResourceFactory {
	return func(tc *testcommon.KubePerTestContext, index int) []client.Object {
		rg := &resources.ResourceGroup{
			ObjectMeta: tc.MakeObjectMeta(fmt.Sprintf("perfvnetrg-%d", index)),
			Spec: resources.ResourceGroup_Spec{
				Location: tc.AzureRegion,
				Tags:     testcommon.CreateTestResourceGroupDefaultTags(),
			},
		}

		vnet := &network.VirtualNetwork{
			ObjectMeta: tc.MakeObjectMeta(fmt.Sprintf("perfvnet-%d", index)),
			Spec: network.VirtualNetwork_Spec{
				Owner:    testcommon.AsOwner(rg),
				Location: tc.AzureRegion,
				AddressSpace: &network.AddressSpace{
					AddressPrefixes: []string{fmt.Sprintf("10.%d.0.0/16", index%256)},
				},
			},
		}

		return []client.Object{rg, vnet}
	}
}
