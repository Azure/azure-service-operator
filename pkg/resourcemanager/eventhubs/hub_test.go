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

	model "github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("Eventhub", func() {

	const timeout = time.Second * 240
	var rgName string
	var eventhubNamespaceName string
	var namespaceLocation string

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test

		rgName = resourceGroupName
		eventhubNamespaceName = "t-ns-dev-eh-" + helpers.RandomString(10)
		namespaceLocation = resourcegroupLocation

		_, _ = CreateNamespaceAndWait(context.Background(), resourceGroupName, eventhubNamespaceName, namespaceLocation)
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

			eventhubName := "t-eh-" + helpers.RandomString(10)
			messageRetentionInDays := int32(7)
			partitionCount := int32(1)

			var err error

			// TODO: add test for Capture
			_, err = CreateHub(context.Background(), rgName, eventhubNamespaceName, eventhubName, messageRetentionInDays, partitionCount, nil)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := GetHub(context.Background(), rgName, eventhubNamespaceName, eventhubName)
				return result.Response.StatusCode == 200
			}, timeout,
			).Should(BeTrue())

			authorizationRuleName := "t-rootmanagedsharedaccesskey"
			accessRights := []model.AccessRights{"Listen", "Manage", "Send"}
			parameters := model.AuthorizationRule{
				AuthorizationRuleProperties: &model.AuthorizationRuleProperties{
					Rights: &accessRights,
				},
			}

			_, err = CreateOrUpdateAuthorizationRule(context.Background(), rgName, eventhubNamespaceName, eventhubName, authorizationRuleName, parameters)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := ListKeys(context.Background(), rgName, eventhubNamespaceName, eventhubName, authorizationRuleName)
				return result.Response.StatusCode == 200
			}, timeout,
			).Should(BeTrue())

			_, err = DeleteHub(context.Background(), rgName, eventhubNamespaceName, eventhubName)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := GetHub(context.Background(), rgName, eventhubNamespaceName, eventhubName)
				return result.Response.StatusCode == 404
			}, timeout,
			).Should(BeTrue())

		})

	})
})
