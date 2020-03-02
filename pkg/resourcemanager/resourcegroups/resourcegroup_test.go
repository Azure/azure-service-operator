// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package resourcegroups

import (
	"context"
	"net/http"
	"time"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/pkg/helpers"
)

var _ = Describe("ResourceGroups", func() {

	const timeout = time.Second * 240

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
	})

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	Context("Create and Delete", func() {
		It("should create and delete resource group in azure", func() {
			const timeout = time.Second * 240

			resourcegroupName := "t-rg-" + helpers.RandomString(10)
			resourcegroupLocation := config.DefaultLocation()
			resourceGroupManager := NewAzureResourceGroupManager()
			var err error

			_, err = resourceGroupManager.CreateGroup(context.Background(), resourcegroupName, resourcegroupLocation)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := resourceGroupManager.CheckExistence(context.Background(), resourcegroupName)

				return result.Response.StatusCode == http.StatusNoContent
			}, timeout,
			).Should(BeTrue())

			_, err = resourceGroupManager.DeleteGroup(context.Background(), resourcegroupName)
			if err != nil {
				azerr := errhelp.NewAzureErrorAzureError(err)
				if azerr.Type == errhelp.AsyncOpIncompleteError {
					err = nil
				}
			}
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := GetGroup(context.Background(), resourcegroupName)
				return result.Response.StatusCode == http.StatusNotFound || *result.Properties.ProvisioningState == "Deleting"
			}, timeout,
			).Should(BeTrue())

		})

	})
})
