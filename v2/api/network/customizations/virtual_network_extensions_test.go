/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package customizations

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/arbitrary"
	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_FuzzySetSubnets(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	vnet := &network.VirtualNetwork_Spec_ARM{
		Location: to.Ptr("westus"),
		Properties: &network.VirtualNetworkPropertiesFormat_ARM{
			EnableDdosProtection: to.Ptr(true),
		},
	}

	// Note that many of these fields are readonly and will not be present on the PUT
	rawVNETWithSubnets := map[string]any{
		"name":     "asotest-vn-wghxbo",
		"id":       "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/asotest-rg-kdebnj/providers/Microsoft.Network/virtualNetworks/asotest-vn-wghxbo",
		"etag":     "W/\"e764086f-6986-403d-a7e5-7e8d99301921\"",
		"type":     "Microsoft.Network/virtualNetworks",
		"location": "westus2",
		"properties": map[string]any{
			"provisioningState": "Succeeded",
			"resourceGuid":      "ec941818-8662-4b6b-b282-4089dc5df651",
			"addressSpace": map[string]any{
				"addressPrefixes": []any{
					"10.0.0.0/16",
				},
			},
			"subnets": []any{
				map[string]any{
					"name": "asotest-subnet-occbrr",
					"id":   "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/asotest-rg-kdebnj/providers/Microsoft.Network/virtualNetworks/asotest-vn-wghxbo/subnets/asotest-subnet-occbrr",
					"etag": "W/\"e764086f-6986-403d-a7e5-7e8d99301921\"",
					"properties": map[string]any{
						"provisioningState": "Succeeded",
						"addressPrefix":     "10.0.0.0/24",
						"natGateway": map[string]any{
							"id": "/this/is/a/test",
						},
						"ipConfigurations": []any{
							map[string]any{
								"id": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/asotest-rg-kdebnj/providers/Microsoft.Network/networkInterfaces/asotest-nic-pvpkfk/ipConfigurations/ipconfig1",
							},
						},
						"privateEndpointNetworkPolicies":    "Disabled",
						"privateLinkServiceNetworkPolicies": "Enabled",
					},
					"type": "Microsoft.Network/virtualNetworks/subnets",
				},
			},
			"virtualNetworkPeerings": []any{},
			"enableDdosProtection":   false,
		},
	}

	azureSubnets, err := getRawChildCollection(rawVNETWithSubnets, "subnets")
	g.Expect(err).ToNot(HaveOccurred())

	err = setChildCollection(vnet, azureSubnets, "Subnets")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(vnet.Location).To(Equal(to.Ptr("westus")))
	g.Expect(vnet.Properties).ToNot(BeNil())
	g.Expect(vnet.Properties.EnableDdosProtection).ToNot(BeNil())
	g.Expect(*vnet.Properties.EnableDdosProtection).To(Equal(true))
	g.Expect(vnet.Properties.Subnets).To(HaveLen(1))
	g.Expect(vnet.Properties.Subnets[0].Properties).ToNot(BeNil())
	g.Expect(vnet.Properties.Subnets[0].Name).To(Equal(to.Ptr("asotest-subnet-occbrr")))
	g.Expect(vnet.Properties.Subnets[0].Properties.AddressPrefix).To(Equal(to.Ptr("10.0.0.0/24")))
	g.Expect(vnet.Properties.Subnets[0].Properties.NatGateway).ToNot(BeNil())
	g.Expect(vnet.Properties.Subnets[0].Properties.NatGateway.Id).To(Equal(to.Ptr("/this/is/a/test")))
}

func Test_FuzzySetSubnet(t *testing.T) {
	t.Parallel()

	embeddedType := reflect.TypeOf(network.Subnet_VirtualNetwork_SubResourceEmbedded_ARM{})
	properties := gopter.NewProperties(nil)
	arbitraries := arbitrary.DefaultArbitraries()

	properties.Property(
		"all subnet types can be converted between non-embedded and embedded",
		arbitraries.ForAll(
			func(subnet *network.VirtualNetworks_Subnet_Spec_ARM) (bool, error) {
				val := reflect.New(embeddedType)

				bytes, err := json.Marshal(subnet)
				if err != nil {
					return false, err
				}

				raw := make(map[string]any)
				err = json.Unmarshal(bytes, &raw)
				if err != nil {
					return false, err
				}

				err = fuzzySetResource(raw, val)
				return err == nil, err
			}))

	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}
