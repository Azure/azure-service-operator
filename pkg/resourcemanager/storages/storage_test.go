// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package storages

import (
	"context"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/go-autorest/autorest/to"

	s "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Storage Account", func() {

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

	Context("Create and Delete Storage Accounts", func() {
		It("should create and delete storage account in azure", func() {

			defer GinkgoRecover()

			storageAccountName := "tdevsa" + helpers.RandomString(10)
			storageLocation := config.DefaultLocation()
			storageManagers := tc.StorageManagers

			ctx := context.Background()

			var err error

			_, err = storageManagers.Storage.CreateStorage(ctx, tc.ResourceGroupName, storageAccountName, storageLocation, azurev1alpha1.StorageSku{
				Name: "Standard_LRS",
			}, "Storage", map[string]*string{}, "", to.BoolPtr(false), to.BoolPtr(false))
			//Expect(err).NotTo(HaveOccurred())

			Eventually(func() s.ProvisioningState {
				result, _ := storageManagers.Storage.GetStorage(ctx, tc.ResourceGroupName, storageAccountName)
				return result.ProvisioningState
			}, tc.timeout, tc.retryInterval,
			).Should(Equal(s.Succeeded))

			_, err = storageManagers.Storage.DeleteStorage(context.Background(), tc.ResourceGroupName, storageAccountName)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				_, err := storageManagers.Storage.GetStorage(context.Background(), tc.ResourceGroupName, storageAccountName)
				return err != nil
			}, tc.timeout, tc.retryInterval,
			).Should(BeTrue())
		})

	})

	Context("Create and Delete Data Lakes", func() {
		It("should create and delete data lakes in azure", func() {

			storageAccountName := "tdevsa" + helpers.RandomString(10)
			storageLocation := config.DefaultLocation()
			datalakeEnabled := true
			storageManagers := tc.StorageManagers

			var err error
			ctx := context.Background()

			_, err = storageManagers.Storage.CreateStorage(ctx, tc.ResourceGroupName, storageAccountName, storageLocation, azurev1alpha1.StorageSku{
				Name: "Standard_LRS",
			}, "StorageV2", map[string]*string{}, "", to.BoolPtr(true), to.BoolPtr(datalakeEnabled))

			Eventually(func() s.ProvisioningState {
				result, _ := storageManagers.Storage.GetStorage(ctx, tc.ResourceGroupName, storageAccountName)
				return result.ProvisioningState
			}, tc.timeout, tc.retryInterval,
			).Should(Equal(s.Succeeded))

			_, err = storageManagers.Storage.DeleteStorage(ctx, tc.ResourceGroupName, storageAccountName)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				_, err := storageManagers.Storage.GetStorage(ctx, tc.ResourceGroupName, storageAccountName)
				return err != nil
			}, tc.timeout, tc.retryInterval,
			).Should(BeTrue())
		})
	})
})
