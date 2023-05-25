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
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
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

	rule := &network.LoadBalancers_InboundNatRule_Spec_ARM{
		Name: "myrule1",
		Properties: &network.InboundNatRulePropertiesFormat_ARM{
			BackendPort:  to.Ptr(22),
			FrontendPort: to.Ptr(23),
		},
	}

	err := fuzzySetInboundNatRules(lb, []genruntime.ARMResourceSpec{rule})
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(lb.Location).To(Equal(to.Ptr("westus")))
	g.Expect(lb.Properties).ToNot(BeNil())
	g.Expect(lb.Properties.InboundNatRules).To(HaveLen(2))
	g.Expect(lb.Properties.InboundNatRules[0].Properties).ToNot(BeNil())
	g.Expect(lb.Properties.InboundNatRules[0].Name).To(Equal(to.Ptr("myrule")))
	g.Expect(lb.Properties.InboundNatRules[0].Properties.BackendPort).ToNot(BeNil())
	g.Expect(lb.Properties.InboundNatRules[0].Properties.FrontendPort).ToNot(BeNil())
	g.Expect(lb.Properties.InboundNatRules[1].Properties).ToNot(BeNil())
	g.Expect(lb.Properties.InboundNatRules[1].Name).To(Equal(to.Ptr("myrule1")))
	g.Expect(lb.Properties.InboundNatRules[1].Properties.BackendPort).ToNot(BeNil())
	g.Expect(lb.Properties.InboundNatRules[1].Properties.FrontendPort).ToNot(BeNil())
}
