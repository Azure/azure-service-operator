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
	"net/http"
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
	var consumerGroupManager ConsumerGroupManager

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
		rgName = tc.ResourceGroupName
		eventhubNamespaceName = helpers.GenerateName("consumer-eventhub-namespace")
		namespaceLocation = tc.ResourceGroupLocation
		eventhubName = helpers.GenerateName("consumer-eventhub")
		messageRetentionInDays = int32(7)
		partitionCount = int32(2)
		consumerGroupManager = tc.EventHubManagers.ConsumerGroup

		_, _ = tc.EventHubManagers.EventHubNamespace.CreateNamespaceAndWait(context.Background(), rgName, eventhubNamespaceName, namespaceLocation)
		_, _ = tc.EventHubManagers.EventHub.CreateHub(context.Background(), rgName, eventhubNamespaceName, eventhubName, messageRetentionInDays, partitionCount, nil)
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

			defer GinkgoRecover()

			consumerGroupName := helpers.GenerateName("eventhub-consumer-group")

			var err error

			_, err = consumerGroupManager.CreateConsumerGroup(context.Background(), rgName, eventhubNamespaceName, eventhubName, consumerGroupName)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := consumerGroupManager.GetConsumerGroup(context.Background(), rgName, eventhubNamespaceName, eventhubName, consumerGroupName)
				return result.Response.StatusCode == http.StatusOK
			}, timeout,
			).Should(BeTrue())

			_, err = consumerGroupManager.DeleteConsumerGroup(context.Background(), rgName, eventhubNamespaceName, eventhubName, consumerGroupName)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := consumerGroupManager.GetConsumerGroup(context.Background(), rgName, eventhubNamespaceName, eventhubName, consumerGroupName)
				return result.Response.StatusCode == http.StatusNotFound
			}, timeout,
			).Should(BeTrue())

		})

	})
})
