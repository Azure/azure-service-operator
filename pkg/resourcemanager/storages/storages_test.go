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
	apiv1 "github.com/Azure/azure-service-operator/api/v1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/go-autorest/autorest/to"
	"time"

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

	Context("Create and Delete", func() {
		It("should create and delete storage account in azure", func() {

			storageAccountName := "tdevsa" + helpers.RandomString(10)
			storageLocation := config.DefaultLocation()

			var err error

			_, err = CreateStorage(context.Background(), resourceGroupName, storageAccountName, storageLocation, apiv1.StorageSku{
				Name: "Standard_LRS",
			}, "Storage", map[string]*string{}, "", to.BoolPtr(false))

			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := GetStorage(context.Background(), resourceGroupName, storageAccountName)
				return result.Response.StatusCode == 200
			}, timeout,
			).Should(BeTrue())

			_, err = DeleteStorage(context.Background(), resourceGroupName, storageAccountName)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := GetStorage(context.Background(), resourceGroupName, storageAccountName)
				return result.Response.StatusCode == 404
			}, timeout,
			).Should(BeTrue())
		})

	})
})
