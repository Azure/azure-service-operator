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

	model "github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Eventhub", func() {

	const timeout = time.Second * 240
	var rgName string
	var eventhubNamespaceName string
	var namespaceLocation string
	var eventHubManager EventHubManager
	var eventHubNamespaceManager EventHubNamespaceManager

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test

		rgName = tc.ResourceGroupName
		eventhubNamespaceName = "t-ns-dev-eh-" + helpers.RandomString(10)
		namespaceLocation = tc.ResourceGroupLocation
		eventHubManager = tc.EventHubManagers.EventHub
		eventHubNamespaceManager = tc.EventHubManagers.EventHubNamespace

		_, _ = eventHubNamespaceManager.CreateNamespaceAndWait(context.Background(), tc.ResourceGroupName, eventhubNamespaceName, namespaceLocation)
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

			defer GinkgoRecover()

			eventhubName := "t-eh-" + helpers.RandomString(10)
			messageRetentionInDays := int32(7)
			partitionCount := int32(2)

			var err error

			// TODO: add test for Capture
			_, err = eventHubManager.CreateHub(context.Background(), rgName, eventhubNamespaceName, eventhubName, messageRetentionInDays, partitionCount, nil)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := eventHubManager.GetHub(context.Background(), rgName, eventhubNamespaceName, eventhubName)
				return result.Response.StatusCode == http.StatusOK
			}, timeout,
			).Should(BeTrue())

			authorizationRuleName := "t-rootmanagedsharedaccesskey"
			accessRights := []model.AccessRights{"Listen", "Manage", "Send"}
			parameters := model.AuthorizationRule{
				AuthorizationRuleProperties: &model.AuthorizationRuleProperties{
					Rights: &accessRights,
				},
			}

			_, err = eventHubManager.CreateOrUpdateAuthorizationRule(context.Background(), rgName, eventhubNamespaceName, eventhubName, authorizationRuleName, parameters)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := eventHubManager.ListKeys(context.Background(), rgName, eventhubNamespaceName, eventhubName, authorizationRuleName)
				return result.Response.StatusCode == http.StatusOK
			}, timeout,
			).Should(BeTrue())

			_, err = eventHubManager.DeleteHub(context.Background(), rgName, eventhubNamespaceName, eventhubName)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := eventHubManager.GetHub(context.Background(), rgName, eventhubNamespaceName, eventhubName)
				return result.Response.StatusCode == http.StatusNotFound
			}, timeout,
			).Should(BeTrue())

		})

	})
})
