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

var _ = Describe("ConsumerGroup", func() {

	const timeout = time.Second * 240
	var rgName string
	var eventhubNamespaceName string
	var eventhubName string
	var namespaceLocation string
	var messageRetentionInDays int32
	var partitionCount int32

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
		rgName = tc.ResourceGroupName
		eventhubNamespaceName = "t-ns-dev-eh-" + helpers.RandomString(10)
		namespaceLocation = tc.ResourcegroupLocation
		eventhubName = "t-eh-dev-ehs-" + helpers.RandomString(10)
		messageRetentionInDays = int32(7)
		partitionCount = int32(2)

		_, _ = CreateNamespaceAndWait(context.Background(), rgName, eventhubNamespaceName, namespaceLocation)
		_, _ = CreateHub(context.Background(), rgName, eventhubNamespaceName, eventhubName, messageRetentionInDays, partitionCount, nil)
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
	})

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	Context("Create and Delete", func() {
		It("should create and delete consumer groups in azure", func() {

			consumerGroupName := "t-cg-" + helpers.RandomString(10)

			var err error

			_, err = CreateConsumerGroup(context.Background(), rgName, eventhubNamespaceName, eventhubName, consumerGroupName)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := GetConsumerGroup(context.Background(), rgName, eventhubNamespaceName, eventhubName, consumerGroupName)
				return result.Response.StatusCode == 200
			}, timeout,
			).Should(BeTrue())

			_, err = DeleteConsumerGroup(context.Background(), rgName, eventhubNamespaceName, eventhubName, consumerGroupName)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := GetConsumerGroup(context.Background(), rgName, eventhubNamespaceName, eventhubName, consumerGroupName)
				return result.Response.StatusCode == 404
			}, timeout,
			).Should(BeTrue())

		})

	})
})
