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

	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
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
		It("should create and delete namespace in azure", func() {

			resourceGroupName := "t-rg-dev-rm-eh"
			eventhubNamespaceName := "t-ns-dev-eh-" + helpers.RandomString(10)
			namespaceLocation := resourcegroupLocation

			var err error

			_, err = CreateNamespaceAndWait(context.Background(), resourceGroupName, eventhubNamespaceName, namespaceLocation)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := GetNamespace(context.Background(), resourceGroupName, eventhubNamespaceName)
				return result.Response.StatusCode == 200
			}, timeout,
			).Should(BeTrue())

			_, err = DeleteNamespace(context.Background(), resourceGroupName, eventhubNamespaceName)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := GetNamespace(context.Background(), resourceGroupName, eventhubNamespaceName)
				return result.Response.StatusCode == 404
			}, timeout,
			).Should(BeTrue())

		})

	})
})
