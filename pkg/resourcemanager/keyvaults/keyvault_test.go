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
	"fmt"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
