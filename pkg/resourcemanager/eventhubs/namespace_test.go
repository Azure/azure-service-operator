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

var _ = Describe("Namespace", func() {

	const timeout = time.Second * 240
	var eventHubNamespaceManager EventHubNamespaceManager

	var rgName string
	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
		rgName = tc.ResourceGroupName
		eventHubNamespaceManager = tc.EventHubManagers.EventHubNamespace
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

			defer GinkgoRecover()

			eventhubNamespaceName := "t-ns-dev-eh-" + helpers.RandomString(10)
			namespaceLocation := tc.ResourceGroupLocation

			var err error

			_, err = eventHubNamespaceManager.CreateNamespaceAndWait(context.Background(), rgName, eventhubNamespaceName, namespaceLocation)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := eventHubNamespaceManager.GetNamespace(context.Background(), rgName, eventhubNamespaceName)
				return result.Response.StatusCode == http.StatusOK
			}, timeout,
			).Should(BeTrue())

			_, err = eventHubNamespaceManager.DeleteNamespace(context.Background(), rgName, eventhubNamespaceName)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := eventHubNamespaceManager.GetNamespace(context.Background(), rgName, eventhubNamespaceName)
				return result.Response.StatusCode == http.StatusNotFound || *result.ProvisioningState == "Deleting"
			}, timeout,
			).Should(BeTrue())

		})

	})
})
