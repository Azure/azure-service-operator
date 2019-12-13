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

package storages

import (
	"context"
	"net/http"

	s "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Blob Container", func() {

	var storageAccountName = "tsadevsc" + helpers.RandomString(10)
	ctx := context.Background()

	BeforeEach(func() {
		storageLocation := config.DefaultLocation()
		// Add any setup steps that needs to be executed before each test
		_, _ = tc.StorageManagers.Storage.CreateStorage(ctx, tc.ResourceGroupName, storageAccountName, storageLocation, azurev1alpha1.StorageSku{
			Name: "Standard_LRS",
		}, "Storage", map[string]*string{}, "", nil, nil)

		Eventually(func() s.ProvisioningState {
			result, _ := tc.StorageManagers.Storage.GetStorage(ctx, tc.ResourceGroupName, storageAccountName)
			return result.ProvisioningState
		}, tc.timeout, tc.retryInterval,
		).Should(Equal(s.Succeeded))
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
		_, err := tc.StorageManagers.Storage.DeleteStorage(ctx, tc.ResourceGroupName, storageAccountName)
		Expect(err).NotTo(HaveOccurred())
	})

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	Context("Create and Delete Blob Containers", func() {
		It("should create and delete blob container in azure", func() {

			var err error
			defer GinkgoRecover()

			containerName := "t-dev-bc-" + helpers.RandomString(10)

			// accessLevel - Specifies whether data in the container may be accessed publicly and the level of access.
			// Possible values include: 'PublicAccessContainer', 'PublicAccessBlob', 'PublicAccessNone'
			accessLevel := s.PublicAccessContainer


			_, err = tc.StorageManagers.BlobContainer.CreateBlobContainer(ctx, tc.ResourceGroupName, storageAccountName, containerName, accessLevel)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := tc.StorageManagers.BlobContainer.GetBlobContainer(ctx, tc.ResourceGroupName, storageAccountName, containerName)
				return result.Response.StatusCode == http.StatusOK
			}, tc.timeout, tc.retryInterval,
			).Should(BeTrue())

			_, err = tc.StorageManagers.BlobContainer.DeleteBlobContainer(ctx, tc.ResourceGroupName, storageAccountName, containerName)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := tc.StorageManagers.BlobContainer.GetBlobContainer(ctx, tc.ResourceGroupName, storageAccountName, containerName)
				return result.Response.StatusCode == http.StatusNotFound
			}, tc.timeout, tc.retryInterval,
			).Should(BeTrue())

		})

	})
})
