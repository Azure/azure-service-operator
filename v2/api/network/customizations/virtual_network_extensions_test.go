/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package customizations

import (
	"os"
	"reflect"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/arbitrary"
	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_FuzzySetSubnets(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	vnet := &network.VirtualNetworks_SpecARM{
		Location: to.StringPtr("westus"),
		Properties: &network.VirtualNetworks_Spec_PropertiesARM{
			EnableDdosProtection: to.BoolPtr(true),
		},
	}

	subnet := &network.VirtualNetworks_Subnets_SpecARM{
		Name: "mysubnet",
		Properties: &network.VirtualNetworks_Subnets_Spec_PropertiesARM{
			AddressPrefix: to.StringPtr("1.2.3.4"),
			NatGateway: &network.SubResourceARM{
				Id: to.StringPtr("/this/is/a/test"),
			},
		},
	}

	err := fuzzySetSubnets(vnet, []genruntime.ARMResourceSpec{subnet})
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(vnet.Location).To(Equal(to.StringPtr("westus")))
	g.Expect(vnet.Properties).ToNot(BeNil())
	g.Expect(vnet.Properties.EnableDdosProtection).ToNot(BeNil())
	g.Expect(*vnet.Properties.EnableDdosProtection).To(Equal(true))
	g.Expect(vnet.Properties.Subnets).To(HaveLen(1))
	g.Expect(vnet.Properties.Subnets[0].Properties).ToNot(BeNil())
	g.Expect(vnet.Properties.Subnets[0].Name).To(Equal(to.StringPtr("mysubnet")))
	g.Expect(vnet.Properties.Subnets[0].Properties.AddressPrefix).To(Equal(to.StringPtr("1.2.3.4")))
	g.Expect(vnet.Properties.Subnets[0].Properties.NatGateway).ToNot(BeNil())
	g.Expect(vnet.Properties.Subnets[0].Properties.NatGateway.Id).To(Equal(to.StringPtr("/this/is/a/test")))
}

func Test_FuzzySetSubnet(t *testing.T) {
	t.Parallel()

	embeddedType := reflect.TypeOf(network.VirtualNetworks_Spec_Properties_SubnetsARM{})
	properties := gopter.NewProperties(nil)
	arbitraries := arbitrary.DefaultArbitraries()

	properties.Property(
		"all subnet types can be converted between non-embedded and embedded",
		arbitraries.ForAll(
			func(subnet *network.VirtualNetworks_Subnets_SpecARM) bool {
				val := reflect.New(embeddedType)
				err := fuzzySetSubnet(subnet, val)

				// This ensures that the self-check that fuzzySetSubnet does did not fail
				return err == nil
			}))

	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}
