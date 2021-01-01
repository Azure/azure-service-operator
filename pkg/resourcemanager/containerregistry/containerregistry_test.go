// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package containerregistry

import (
	"fmt"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("KeyVault Resource Manager test", func() {

	var rgName string
	var location string
	var containerRegistryName string
	var azureContainerRegistryManager azureContainerRegistryManager

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
		rgName = tc.ResourceGroupName
		location = tc.ResourceGroupLocation
		azureContainerRegistryManager = tc.azureContainerRegistryManager
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
	})

	// Add Tests for OpenAPI validation (or additional CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	Context("Create and Delete", func() {
		It("should create and delete Container Registry in azure", func() {

			defer GinkgoRecover()

			containerRegistryName = "tdevacr" + helpers.RandomString(10)

			acr := v1alpha1.AzureContainerRegistry{
				ObjectMeta: metav1.ObjectMeta{
					Name: containerRegistryName,
				},
				Spec: v1alpha1.AzureContainerRegistrySpec{
					ResourceGroup:    rgName,
					Location:         location,
					AdminUserEnabled: true,
					Sku:              v1alpha1.ContainerRegistrySku{Name: "Standard"},
				},
			}

			// Create Container Registry
			Eventually(func() bool {
				_, err := azureContainerRegistryManager.CreateRegistry(
					ctx,
					&acr)

				if err != nil {
					fmt.Println(err.Error())
					ignore := []string{
						errhelp.AsyncOpIncompleteError,
					}
					azerr := errhelp.NewAzureError(err)
					if !helpers.ContainsString(ignore, azerr.Type) {
						fmt.Println("error occurred")
						return false
					}
				}
				return true
			}, tc.timeout, tc.retryInterval,
			).Should(BeTrue())

			Eventually(func() bool {
				_, err := azureContainerRegistryManager.GetRegistry(ctx, rgName, containerRegistryName)
				if err == nil {
					return true
				}
				return false
			}, tc.timeout, tc.retryInterval,
			).Should(BeTrue())

			// Delete KeyVault instance
			Eventually(func() bool {
				_, err := azureContainerRegistryManager.DeleteRegistry(ctx, rgName, containerRegistryName)
				if err != nil {
					fmt.Println(err.Error())
					ignore := []string{
						errhelp.AsyncOpIncompleteError,
					}
					azerr := errhelp.NewAzureError(err)
					if !helpers.ContainsString(ignore, azerr.Type) {
						fmt.Println("error occurred")
						return false
					}
				}
				return err == nil
			}, tc.timeout, tc.retryInterval,
			).Should(BeTrue())

			Eventually(func() bool {
				_, err := azureContainerRegistryManager.DeleteRegistry(ctx, rgName, containerRegistryName)
				if err == nil {
					return true
				}
				return false
			}, tc.timeout, tc.retryInterval,
			).Should(BeFalse())
		})

	})
})
