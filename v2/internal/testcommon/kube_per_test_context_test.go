/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestSanitiseSample(t *testing.T) {
	cases := []struct {
		name     string
		in       string
		expected string
	}{
		{
			"Subscriptions in URLs",
			"armId: /subscriptions/2caff979-f2e6-4629-9f02-e883e6a0d652/resourceGroups/aso-sample-rg/providers/Microsoft.Network/loadBalancers/sampleloadbalancer/inboundNatPools/samplenatpool",
			"armId: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-sample-rg/providers/Microsoft.Network/loadBalancers/sampleloadbalancer/inboundNatPools/samplenatpool",
		},
		{
			"Subscriptions in URLs, preserve letter case",
			"armId: /Subscriptions/2caff979-f2e6-4629-9f02-e883e6a0d652/resourceGroups/aso-sample-rg/providers/Microsoft.Network/loadBalancers/sampleloadbalancer/inboundNatPools/samplenatpool",
			"armId: /Subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-sample-rg/providers/Microsoft.Network/loadBalancers/sampleloadbalancer/inboundNatPools/samplenatpool",
		},
		{
			"Change asotest- naming to aso-sample-, removing random parts",
			"asotest-rg-asdf",
			"aso-sample-rg",
		},
		{
			"Remove azurename entirely",
			"spec:\n  azurename: asotest-image-hkobbe \n  size: demo",
			"spec:\n  size: demo",
		},
		{
			"Remove azurename entirely, case insensitive",
			"spec:\n  azureName: asotest-image-hkobbe \n  size: demo",
			"spec:\n  size: demo",
		},
	}

	t.Parallel()
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			actual := sanitiseSample(c.in)
			g.Expect(actual).To(Equal(c.expected))
		})
	}
}
