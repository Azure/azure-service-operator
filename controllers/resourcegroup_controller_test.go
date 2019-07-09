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

package controllers

import (
	"context"
	"time"

	azurev1 "Telstra.Dx.AzureOperator/api/v1"
	helpers "Telstra.Dx.AzureOperator/helpers"
	. "github.com/onsi/ginkgo"

	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("ResourceGroup Controller", func() {

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
		It("should create and delete resource groups", func() {
			resourceGroupName := "t-rg-dev-" + helpers.RandomString(10)

			var err error

			// Create the Resourcegroup object and expect the Reconcile to be created
			resourceGroupInstance := &azurev1.ResourceGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      resourceGroupName,
					Namespace: "default",
				},
				Spec: azurev1.ResourceGroupSpec{
					Location: "westus",
				},
			}

			err = k8sClient.Create(context.Background(), resourceGroupInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			resourceGroupNamespacedName := types.NamespacedName{Name: resourceGroupName, Namespace: "default"}

			time.Sleep(30 * time.Second)

			k8sClient.Delete(context.Background(), resourceGroupInstance)
			Eventually(func() bool {
				_ = k8sClient.Get(context.Background(), resourceGroupNamespacedName, resourceGroupInstance)
				return resourceGroupInstance.IsBeingDeleted()
			}, timeout,
			).Should(BeTrue())

		})
	})
})
