/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20240301"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Network_PublicIPAddress_creationNotBlockedByPublicIPPrefix_20240301(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	prefix := &network.PublicIPPrefix{
		ObjectMeta: tc.MakeObjectMetaWithName("prefix"),
		Spec: network.PublicIPPrefix_Spec{
			Location:               tc.AzureRegion,
			Owner:                  testcommon.AsOwner(rg),
			PrefixLength:           to.Ptr(29),
			PublicIPAddressVersion: to.Ptr(network.IPVersion_IPv4),
			Sku: &network.PublicIPPrefixSku{
				Name: to.Ptr(network.PublicIPPrefixSku_Name_Standard),
				Tier: to.Ptr(network.PublicIPPrefixSku_Tier_Regional),
			},
		},
	}

	ip1 := &network.PublicIPAddress{
		ObjectMeta: tc.MakeObjectMetaWithName("ip1"),
		Spec: network.PublicIPAddress_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &network.PublicIPAddressSku{
				Name: to.Ptr(network.PublicIPAddressSku_Name_Standard),
			},
			PublicIPAllocationMethod: to.Ptr(network.IPAllocationMethod_Static),
			PublicIPAddressVersion:   to.Ptr(network.IPVersion_IPv4),
			PublicIPPrefix: &network.SubResource{
				Reference: tc.MakeReferenceFromResource(prefix),
			},
		},
	}

	ip2 := &network.PublicIPAddress{
		ObjectMeta: tc.MakeObjectMetaWithName("ip2"),
		Spec: network.PublicIPAddress_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &network.PublicIPAddressSku{
				Name: to.Ptr(network.PublicIPAddressSku_Name_Standard),
			},
			PublicIPAllocationMethod: to.Ptr(network.IPAllocationMethod_Static),
			PublicIPAddressVersion:   to.Ptr(network.IPVersion_IPv4),
			PublicIPPrefix: &network.SubResource{
				Reference: tc.MakeReferenceFromResource(prefix),
			},
		},
	}

	ip3 := &network.PublicIPAddress{
		ObjectMeta: tc.MakeObjectMetaWithName("ip3"),
		Spec: network.PublicIPAddress_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &network.PublicIPAddressSku{
				Name: to.Ptr(network.PublicIPAddressSku_Name_Standard),
			},
			PublicIPAllocationMethod: to.Ptr(network.IPAllocationMethod_Static),
			PublicIPAddressVersion:   to.Ptr(network.IPVersion_IPv4),
			PublicIPPrefix: &network.SubResource{
				Reference: tc.MakeReferenceFromResource(prefix),
			},
		},
	}

	tc.CreateResourcesAndWait(prefix, ip1, ip2, ip3)
}
