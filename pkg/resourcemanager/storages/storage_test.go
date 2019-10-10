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
	"time"

	apiv1 "github.com/Azure/azure-service-operator/api/v1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/go-autorest/autorest/to"

	"github.com/Azure/azure-service-operator/pkg/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Storage Account", func() {

	const timeout = time.Second * 180

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

			storageAccountName := "tdevsa" + helpers.RandomString(10)
			storageLocation := config.DefaultLocation()
			storageManagers := tc.StorageManagers

			var err error

			_, err = storageManagers.Storage.CreateStorage(context.Background(), tc.ResourceGroupName, storageAccountName, storageLocation, apiv1.StorageSku{
				Name: "Standard_LRS",
			}, "Storage", map[string]*string{}, "", to.BoolPtr(false))

			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := storageManagers.Storage.GetStorage(context.Background(), tc.ResourceGroupName, storageAccountName)
				return result.Response.StatusCode == http.StatusOK
			}, timeout,
			).Should(BeTrue())

			_, err = storageManagers.Storage.DeleteStorage(context.Background(), tc.ResourceGroupName, storageAccountName)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := storageManagers.Storage.GetStorage(context.Background(), tc.ResourceGroupName, storageAccountName)
				return result.Response.StatusCode == http.StatusNotFound
			}, timeout,
			).Should(BeTrue())
		})

	})

	// Context("Create and Delete Data Lakes", func() {
	// 	It("should create and delete data lakes in azure", func() {

	// 		storageAccountName := "tdevsa" + helpers.RandomString(10)
	// 		storageLocation := config.DefaultLocation()
	// 		datalakeEnabled := true
	// 		storageManagers := tc.StorageManagers

	// 		var err error

	// 		_, err = storageManagers.Storage.CreateStorage(context.Background(), tc.ResourceGroupName, storageAccountName, storageLocation, apiv1.StorageSku{
	// 			Name: "Standard_LRS",
	// 		}, "Storage", map[string]*string{}, "", to.BoolPtr(false), to.BoolPtr(datalakeEnabled))

	// 		Expect(err).NotTo(HaveOccurred())

	// 		Eventually(func() bool {
	// 			result, _ := storageManagers.Storage.GetStorage(context.Background(), tc.ResourceGroupName, storageAccountName)
	// 			return result.Response.StatusCode == http.StatusOK
	// 		}, timeout,
	// 		).Should(BeTrue())

	// 		_, err = storageManagers.Storage.DeleteStorage(context.Background(), tc.ResourceGroupName, storageAccountName)
	// 		Expect(err).NotTo(HaveOccurred())

	// 		Eventually(func() bool {
	// 			result, _ := storageManagers.Storage.GetStorage(context.Background(), tc.ResourceGroupName, storageAccountName)
	// 			return result.Response.StatusCode == http.StatusNotFound
	// 		}, timeout,
	// 		).Should(BeTrue())
	// 	})

	// })
})
