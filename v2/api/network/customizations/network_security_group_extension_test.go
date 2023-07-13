/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package customizations

import (
	"testing"

	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_FuzzySetNetworkSecurityGroupsSecurityRules(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	nsg := &network.NetworkSecurityGroup_Spec_ARM{
		Location: to.Ptr("westus"),
		Name:     "my-nsg",
	}

	// Note that many of these fields are readonly and will not be present on the PUT
	rawNSG := map[string]any{
		"name":     "mynsg",
		"id":       "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myrg/providers/Microsoft.Network/networkSecurityGroups/mynsg",
		"etag":     "W/\"ae3ff66d-e583-4ef9-a83c-a2e00029e2ed\"",
		"type":     "Microsoft.Network/networkSecurityGroups",
		"location": "westus",
		"properties": map[string]any{
			"provisioningState": "Succeeded",
			"resourceGuid":      "de9b1f0d-1911-4a5f-9902-74d4c14018d7",
			"securityRules": []any{
				map[string]any{
					"name": "myrule",
					"id":   "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myrg/providers/Microsoft.Network/networkSecurityGroups/mynsg/securityRules/myrule",
					"etag": "W/\"ae3ff66d-e583-4ef9-a83c-a2e00029e2ed\"",
					"properties": map[string]any{
						"provisioningState":        "Succeeded",
						"access":                   "Allow",
						"destinationAddressPrefix": "*",
						"destinationPortRange":     "46-56",
						"direction":                "Inbound",
						"priority":                 123,
						"protocol":                 "Tcp",
						"sourceAddressPrefix":      "*",
						"sourcePortRange":          "23-45",
					},
					"type": "Microsoft.Network/networkSecurityGroups/securityRules",
				},
			},
		},
	}

	azureSecurityRules, err := getRawChildCollection(rawNSG, "securityRules")
	g.Expect(err).ToNot(HaveOccurred())

	err = setChildCollection(nsg, azureSecurityRules, "SecurityRules")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(nsg.Location).To(Equal(to.Ptr("westus")))
	g.Expect(nsg.Properties).ToNot(BeNil())
	g.Expect(nsg.Properties.SecurityRules).To(HaveLen(1))
	g.Expect(nsg.Properties.SecurityRules[0].Properties).ToNot(BeNil())
	g.Expect(nsg.Properties.SecurityRules[0].Name).To(Equal(to.Ptr("myrule")))
	g.Expect(nsg.Properties.SecurityRules[0].Properties.DestinationAddressPrefix).ToNot(BeNil())
	g.Expect(nsg.Properties.SecurityRules[0].Properties.DestinationPortRange).ToNot(BeNil())
}
