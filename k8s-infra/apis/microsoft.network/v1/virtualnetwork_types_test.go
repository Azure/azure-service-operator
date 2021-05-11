/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v1

import (
	"encoding/json"
	"testing"

	"github.com/onsi/gomega"

	v1 "github.com/Azure/k8s-infra/apis/core/v1"
)

func TestVirtualNetwork_Marshalling(t *testing.T) {
	vnet := VirtualNetwork{
		Spec: VirtualNetworkSpec{
			APIVersion: "api",
			Location:   "westus2",
			Tags:       nil,
			Properties: &VirtualNetworkSpecProperties{
				AddressSpace: &AddressSpaceSpec{
					AddressPrefixes: []string{
						"10.0.0.0/16",
					},
				},
				SubnetRefs: []v1.KnownTypeReference{
					{
						Name: "test-1",
					},
				},
				EnableVMProtection: false,
			},
		},
	}

	bits, err := json.Marshal(vnet)
	g := gomega.NewGomegaWithT(t)
	g.Expect(err).ToNot(gomega.HaveOccurred())
	js := string(bits)
	g.Expect(js).ToNot(gomega.ContainSubstring("bgpCommunities"))
	g.Expect(js).ToNot(gomega.ContainSubstring("dhcpOptions"))
	g.Expect(js).ToNot(gomega.ContainSubstring("enableVMProtection"))
	g.Expect(js).To(gomega.ContainSubstring("subnetRefs"))
	g.Expect(js).To(gomega.ContainSubstring("addressSpace"))
}
