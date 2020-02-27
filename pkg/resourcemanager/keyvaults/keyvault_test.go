/*
Copyright 2019 microsoft.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package keyvaults

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2018-02-14/keyvault"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

	ctx = context.Background()

	resp, err := parseAccessPolicy(&entry, ctx)
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

	ctx = context.Background()

	resp, err := parseAccessPolicy(&entry, ctx)
	assert.True(t, err == nil)
	assert.True(t, cmp.Equal(resp, out))
}

var _ = Describe("KeyVault Resource Manager test", func() {

	var rgName string
	var location string
	var keyvaultName string
	var keyVaultManager KeyVaultManager

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
		rgName = tc.ResourceGroupName
		location = tc.ResourceGroupLocation
		keyVaultManager = tc.keyvaultManager
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
	})

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	Context("Create and Delete", func() {
		It("should create and delete Key Vaults in azure", func() {

			defer GinkgoRecover()

			keyvaultName = "t-dev-kv-" + helpers.RandomString(10)

			tags := map[string]*string{
				"tag1": to.StringPtr("value1"),
				"tag2": to.StringPtr("value2"),
			}

			kv := v1alpha1.KeyVault{
				ObjectMeta: metav1.ObjectMeta{
					Name: keyvaultName,
				},
				Spec: v1alpha1.KeyVaultSpec{
					ResourceGroup: rgName,
					Location:      location,
				},
			}

			// Create Key Vault instance
			Eventually(func() bool {
				_, err := keyVaultManager.CreateVault(
					ctx,
					&kv,
					tags,
				)
				if err != nil {
					fmt.Println(err.Error())
					if !errhelp.IsAsynchronousOperationNotComplete(err) {
						fmt.Println("error occured")
						return false
					}
				}
				return true
			}, tc.timeout, tc.retryInterval,
			).Should(BeTrue())

			Eventually(func() bool {
				_, err := keyVaultManager.GetVault(ctx, rgName, keyvaultName)
				if err == nil {
					return true
				}
				return false
			}, tc.timeout, tc.retryInterval,
			).Should(BeTrue())

			// Delete KeyVault instance
			Eventually(func() bool {
				_, err := keyVaultManager.DeleteVault(ctx, rgName, keyvaultName)
				if err != nil {
					fmt.Println(err.Error())
					if !errhelp.IsAsynchronousOperationNotComplete(err) {
						fmt.Println("error occured")
						return false
					}
				}
				return err == nil
			}, tc.timeout, tc.retryInterval,
			).Should(BeTrue())

			Eventually(func() bool {
				_, err := keyVaultManager.GetVault(ctx, rgName, keyvaultName)
				if err == nil {
					return true
				}
				return false
			}, tc.timeout, tc.retryInterval,
			).Should(BeFalse())
		})

	})
})
