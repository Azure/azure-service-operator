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

package eventhubs

import (
	"context"
	"time"

	resoucegroupsresourcemanager "Telstra.Dx.AzureOperator/resourcemanager/resourcegroups"

	helpers "Telstra.Dx.AzureOperator/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Namespace", func() {

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
		It("should create and delete hubs in azure", func() {

			resourceGroupName := "t-rg-dev-eh-" + helpers.RandomString(10)
			resourcegroupLocation := "westus"
			eventhubNamespaceName := "t-ns-dev-eh-" + helpers.RandomString(10)
			namespaceLocation := "westus"
			eventhubName := "t-eh-" + helpers.RandomString(10)
			var err error

			_, err = resoucegroupsresourcemanager.CreateGroup(context.Background(), resourceGroupName, resourcegroupLocation)
			Expect(err).NotTo(HaveOccurred())

			time.Sleep(30 * time.Second)

			_, err = CreateNamespace(context.Background(), resourceGroupName, eventhubNamespaceName, namespaceLocation)
			Expect(err).NotTo(HaveOccurred())

			time.Sleep(30 * time.Second)

			_, err = CreateHub(context.Background(), resourceGroupName, eventhubNamespaceName, eventhubName)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := GetHub(context.Background(), resourceGroupName, eventhubNamespaceName, eventhubName)
				return result.Response.StatusCode == 200
			}, timeout,
			).Should(BeTrue())

			_, err = DeleteHub(context.Background(), resourceGroupName, eventhubNamespaceName, eventhubName)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := GetHub(context.Background(), resourceGroupName, eventhubNamespaceName, eventhubName)
				return result.Response.StatusCode == 404
			}, timeout,
			).Should(BeTrue())

			time.Sleep(30 * time.Second)

			_, err = resoucegroupsresourcemanager.DeleteGroup(context.Background(), resourceGroupName)
			Expect(err).NotTo(HaveOccurred())

		})

	})
})
