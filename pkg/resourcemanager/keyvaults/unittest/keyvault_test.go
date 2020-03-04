// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package testing

import (
	"context"
	"testing"

	keyvault "github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2018-02-14/keyvault"
	v1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	azurekeyvault "github.com/Azure/azure-service-operator/pkg/resourcemanager/keyvaults"
	"github.com/google/go-cmp/cmp"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
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

	resp, err := azurekeyvault.ParseAccessPolicy(&entry, ctx)
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

	resp, err := azurekeyvault.ParseAccessPolicy(&entry, ctx)
	assert.True(t, err == nil)
	assert.True(t, cmp.Equal(resp, out))
}
