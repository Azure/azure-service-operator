// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package keyvaults

import (
	"fmt"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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

			sku := v1alpha1.KeyVaultSku{
				Name: "Standard",
			}

			// Create Key Vault instance
			Eventually(func() bool {
				_, err := keyVaultManager.CreateVault(
					ctx,
					&kv,
					sku,
					tags,
				)
				if err != nil {
					fmt.Println(err.Error())
					ignore := []string{
						errhelp.AsyncOpIncompleteError,
					}
					azerr := errhelp.NewAzureErrorAzureError(err)
					if !helpers.ContainsString(ignore, azerr.Type) {
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
					ignore := []string{
						errhelp.AsyncOpIncompleteError,
					}
					azerr := errhelp.NewAzureErrorAzureError(err)
					if !helpers.ContainsString(ignore, azerr.Type) {
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
