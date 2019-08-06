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
	resoucegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
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
			messageRetentionInDays := int32(7)
			partitionCount := int32(1)

			var err error

			_, err = resoucegroupsresourcemanager.CreateGroup(context.Background(), resourceGroupName, resourcegroupLocation)
			Expect(err).NotTo(HaveOccurred())

			time.Sleep(30 * time.Second)

			_, err = CreateNamespaceAndWait(context.Background(), resourceGroupName, eventhubNamespaceName, namespaceLocation)
			Expect(err).NotTo(HaveOccurred())

			time.Sleep(30 * time.Second)

			_, err = CreateHub(context.Background(), resourceGroupName, eventhubNamespaceName, eventhubName, messageRetentionInDays, partitionCount)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := GetHub(context.Background(), resourceGroupName, eventhubNamespaceName, eventhubName)
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
			_, err = CreateOrUpdateAuthorizationRule(context.Background(), resourceGroupName, eventhubNamespaceName, eventhubName, authorizationRuleName, parameters)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				result, _ := ListKeys(context.Background(), resourceGroupName, eventhubNamespaceName, eventhubName, authorizationRuleName)
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
