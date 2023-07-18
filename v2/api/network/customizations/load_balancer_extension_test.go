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

func Test_FuzzySetLoadBalancers(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	lb := &network.LoadBalancer_Spec_ARM{
		Location: to.Ptr("westus"),
		Name:     "mylb",
		Properties: &network.LoadBalancerPropertiesFormat_ARM{
			InboundNatRules: []network.InboundNatRule_LoadBalancer_SubResourceEmbedded_ARM{
				{
					Name: to.Ptr("myrule"),
					Properties: &network.InboundNatRulePropertiesFormat_ARM{
						BackendPort:  to.Ptr(80),
						FrontendPort: to.Ptr(90),
					},
				},
			},
		},
	}

	// Note that many of these fields are readonly and will not be present on the PUT
	rawLBWithNatRules := map[string]any{
		"name":     "mylb",
		"id":       "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myrg/providers/Microsoft.Network/loadBalancers/mylb",
		"etag":     "W/\"e764086f-6986-403d-a7e5-7e8d99301921\"",
		"type":     "Microsoft.Network/loadBalancers",
		"location": "westus",
		"properties": map[string]any{
			"provisioningState": "Succeeded",
			"resourceGuid":      "a54785bf-72a4-4c89-ae93-fef560df2b53",
			"inboundNatRules": []any{
				map[string]any{
					"name": "myrule1",
					"id":   "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myrg/providers/Microsoft.Network/loadBalancers/mylb/inboundNatRules/myrule1",
					"etag": "W/\"e764086f-6986-403d-a7e5-7e8d99301921\"",
					"type": "Microsoft.Network/loadBalancers/inboundNatRules",
					"properties": map[string]any{
						"backendPort":  22,
						"frontendPort": 23,
					},
				},
			},
		},
	}

	azureInboundNatRules, err := getRawChildCollection(rawLBWithNatRules, "inboundNatRules")
	g.Expect(err).ToNot(HaveOccurred())

	err = setInboundNatRules(lb, azureInboundNatRules)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(lb.Location).To(Equal(to.Ptr("westus")))
	g.Expect(lb.Properties).ToNot(BeNil())
	g.Expect(lb.Properties.InboundNatRules).To(HaveLen(2))

	rule0 := lb.Properties.InboundNatRules[0]
	rule1 := lb.Properties.InboundNatRules[1]

	g.Expect(rule0.Properties).ToNot(BeNil())
	g.Expect(rule0.Name).To(Equal(to.Ptr("myrule")))
	g.Expect(rule0.Properties.BackendPort).ToNot(BeNil())
	g.Expect(rule0.Properties.FrontendPort).ToNot(BeNil())
	g.Expect(rule1.Properties).ToNot(BeNil())
	g.Expect(rule1.Name).To(Equal(to.Ptr("myrule1")))
	g.Expect(rule1.Properties.BackendPort).ToNot(BeNil())
	g.Expect(rule1.Properties.FrontendPort).ToNot(BeNil())
}

func Test_FuzzySetLoadBalancers_NatRulesMerged(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	lb := &network.LoadBalancer_Spec_ARM{
		Location: to.Ptr("westus"),
		Name:     "mylb",
		Properties: &network.LoadBalancerPropertiesFormat_ARM{
			InboundNatRules: []network.InboundNatRule_LoadBalancer_SubResourceEmbedded_ARM{
				{
					Name: to.Ptr("myrule"),
					Properties: &network.InboundNatRulePropertiesFormat_ARM{
						BackendPort:  to.Ptr(80),
						FrontendPort: to.Ptr(90),
					},
				},
				{
					Name: to.Ptr("myrule1"),
					Properties: &network.InboundNatRulePropertiesFormat_ARM{
						BackendPort:  to.Ptr(22),
						FrontendPort: to.Ptr(23),
					},
				},
			},
		},
	}

	// Note that many of these fields are readonly and will not be present on the PUT
	rawLBWithNatRules := map[string]any{
		"name":     "mylb",
		"id":       "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myrg/providers/Microsoft.Network/loadBalancers/mylb",
		"etag":     "W/\"e764086f-6986-403d-a7e5-7e8d99301921\"",
		"type":     "Microsoft.Network/loadBalancers",
		"location": "westus",
		"properties": map[string]any{
			"provisioningState": "Succeeded",
			"resourceGuid":      "a54785bf-72a4-4c89-ae93-fef560df2b53",
			"inboundNatRules": []any{
				map[string]any{
					"name": "myrule1",
					"id":   "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myrg/providers/Microsoft.Network/loadBalancers/mylb/inboundNatRules/myrule1",
					"etag": "W/\"e764086f-6986-403d-a7e5-7e8d99301921\"",
					"type": "Microsoft.Network/loadBalancers/inboundNatRules",
					"properties": map[string]any{
						"backendPort":  22,
						"frontendPort": 23,
					},
				},
			},
		},
	}

	azureInboundNatRules, err := getRawChildCollection(rawLBWithNatRules, "inboundNatRules")
	g.Expect(err).ToNot(HaveOccurred())

	err = setInboundNatRules(lb, azureInboundNatRules)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(lb.Location).To(Equal(to.Ptr("westus")))
	g.Expect(lb.Properties).ToNot(BeNil())
	g.Expect(lb.Properties.InboundNatRules).To(HaveLen(2))

	rule0 := lb.Properties.InboundNatRules[0]
	rule1 := lb.Properties.InboundNatRules[1]

	g.Expect(rule0.Properties).ToNot(BeNil())
	g.Expect(rule0.Name).To(Equal(to.Ptr("myrule")))
	g.Expect(rule0.Properties.BackendPort).ToNot(BeNil())
	g.Expect(rule0.Properties.FrontendPort).ToNot(BeNil())
	g.Expect(rule1.Properties).ToNot(BeNil())
	g.Expect(rule1.Name).To(Equal(to.Ptr("myrule1")))
	g.Expect(rule1.Properties.BackendPort).ToNot(BeNil())
	g.Expect(rule1.Properties.FrontendPort).ToNot(BeNil())
}
