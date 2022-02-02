// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package testing

import (
	"context"
	"testing"

	keyvault "github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2018-02-14/keyvault"
	uuid "github.com/gofrs/uuid"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"

	v1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	azurekeyvault "github.com/Azure/azure-service-operator/pkg/resourcemanager/keyvaults"
)

func TestParseAccessPoliciesInvalid(t *testing.T) {
	entry := v1alpha1.AccessPolicyEntry{
		TenantID: "00000000-0000-0000-0000-000000000000",
		Permissions: &v1alpha1.Permissions{
			Keys: &[]string{
				"create",
			},
			Secrets: &[]string{
				"set",
				"badpermission",
			},
		},
	}

	ctx := context.Background()

	resp, err := azurekeyvault.ParseAccessPolicy(ctx, config.GlobalCredentials(), &entry)
	assert.True(t, err != nil)
	assert.True(t, cmp.Equal(resp, keyvault.AccessPolicyEntry{}))
}

func TestParseAccessPolicies(t *testing.T) {
	entry := v1alpha1.AccessPolicyEntry{
		TenantID: "00000000-0000-0000-0000-000000000000",
		Permissions: &v1alpha1.Permissions{
			Keys: &[]string{
				"create",
			},
			Secrets: &[]string{
				"set",
				"get",
			},
			Certificates: &[]string{
				"list",
				"get",
			},
			Storage: &[]string{
				"list",
				"get",
			},
		},
	}

	id, err := uuid.FromString("00000000-0000-0000-0000-000000000000")
	assert.True(t, err == nil)

	out := keyvault.AccessPolicyEntry{
		TenantID: &id,
		Permissions: &keyvault.Permissions{
			Keys: &[]keyvault.KeyPermissions{
				keyvault.KeyPermissionsCreate,
			},
			Secrets: &[]keyvault.SecretPermissions{
				keyvault.SecretPermissionsSet,
				keyvault.SecretPermissionsGet,
			},
			Certificates: &[]keyvault.CertificatePermissions{
				keyvault.List,
				keyvault.Get,
			},
			Storage: &[]keyvault.StoragePermissions{
				keyvault.StoragePermissionsList,
				keyvault.StoragePermissionsGet,
			},
		},
	}

	ctx := context.Background()

	resp, err := azurekeyvault.ParseAccessPolicy(ctx, config.GlobalCredentials(), &entry)
	assert.True(t, err == nil)
	assert.True(t, cmp.Equal(resp, out))
}

func TestParseNetworkPolicy(t *testing.T) {
	ip1 := "191.10.1.0/24"
	ip2 := "191.10.2.0/24"
	vnetRule1 := "/subscriptions/<subid>/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet1"
	vnetRule2 := "/subscriptions/<subid>/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet2"
	entry := v1alpha1.NetworkRuleSet{
		Bypass:        "None",
		DefaultAction: "Deny",
		IPRules: &[]string{
			ip1,
			ip2,
		},
		VirtualNetworkRules: &[]string{
			vnetRule1,
			vnetRule2,
		},
	}

	out := keyvault.NetworkRuleSet{
		Bypass:        keyvault.None,
		DefaultAction: keyvault.Deny,
		IPRules: &[]keyvault.IPRule{
			{
				Value: &ip1,
			},
			{
				Value: &ip2,
			},
		},
		VirtualNetworkRules: &[]keyvault.VirtualNetworkRule{
			{
				ID: &vnetRule1,
			},
			{
				ID: &vnetRule2,
			},
		},
	}

	resp := azurekeyvault.ParseNetworkPolicy(&entry)

	assert.True(t, cmp.Equal(resp, out))
}
