// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

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
		eventhubNamespaceName = "t-ns-dev-eh-" + helpers.RandomString(10)
		namespaceLocation = tc.ResourceGroupLocation
		eventhubName = "t-eh-dev-ehs-" + helpers.RandomString(10)
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

			consumerGroupName := "t-cg-" + helpers.RandomString(10)

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
