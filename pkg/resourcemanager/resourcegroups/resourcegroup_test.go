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

package resourcegroups

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
)

var _ = Describe("ResourceGroups", func() {

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
		It("should create and delete resource group in azure", func() {
			const timeout = time.Second * 240

			resourcegroupName := "t-rg-" + helpers.RandomString(10)
			resourcegroupLocation := "westus"
			var err error

			_, err = CreateGroup(context.Background(), resourcegroupName, resourcegroupLocation)
			Expect(err).NotTo(HaveOccurred())

			time.Sleep(40 * time.Second)

			Eventually(func() bool {
				result, _ := CheckExistence(context.Background(), resourcegroupName)

				return result.Response.StatusCode == 204
			}, timeout,
			).Should(BeTrue())

			_, err = DeleteGroup(context.Background(), resourcegroupName)
			Expect(err).NotTo(HaveOccurred())

			time.Sleep(30 * time.Second)

			Eventually(func() bool {
				result, _ := CheckExistence(context.Background(), resourcegroupName)

				return result.Response.StatusCode == 404
			}, timeout,
			).Should(BeTrue())

			//DeleteAllGroupsWithPrefix(context.Background(), "t-rg-")
		})

	})
})
