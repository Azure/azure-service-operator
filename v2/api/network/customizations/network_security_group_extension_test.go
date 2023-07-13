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

func Test_FuzzySetNetworkSecurityGroupsSecurityRules(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	nsg := &network.NetworkSecurityGroup_Spec_ARM{
		Location: to.Ptr("westus"),
		Name:     "my-nsg",
	}

	rule := &network.NetworkSecurityGroups_SecurityRule_Spec_ARM{
		Name: "myrule",
		Properties: &network.SecurityRulePropertiesFormat_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM{
			Access:                   to.Ptr(network.SecurityRuleAccess_Allow),
			DestinationAddressPrefix: to.Ptr("*"),
			DestinationPortRange:     to.Ptr("46-56"),
			Direction:                to.Ptr(network.SecurityRuleDirection_Inbound),
			Priority:                 to.Ptr(123),
			Protocol:                 to.Ptr(network.SecurityRulePropertiesFormat_Protocol_Tcp),
			SourceAddressPrefix:      to.Ptr("*"),
			SourcePortRange:          to.Ptr("23-45"),
		},
	}

	err := fuzzySetResources(nsg, []genruntime.ARMResourceSpec{rule}, "SecurityRules")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(nsg.Location).To(Equal(to.Ptr("westus")))
	g.Expect(nsg.Properties).ToNot(BeNil())
	g.Expect(nsg.Properties.SecurityRules).To(HaveLen(1))
	g.Expect(nsg.Properties.SecurityRules[0].Properties).ToNot(BeNil())
	g.Expect(nsg.Properties.SecurityRules[0].Name).To(Equal(to.Ptr("myrule")))
	g.Expect(nsg.Properties.SecurityRules[0].Properties.DestinationAddressPrefix).ToNot(BeNil())
	g.Expect(nsg.Properties.SecurityRules[0].Properties.DestinationPortRange).ToNot(BeNil())
}
