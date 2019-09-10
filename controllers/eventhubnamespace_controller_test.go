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

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"

	"time"

	. "github.com/onsi/ginkgo"

	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("EventHubNamespace Controller", func() {

	const timeout = time.Second * 240
	var rgName string
	var rgLocation string

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
		rgName = tc.ResourceGroupName
		rgLocation = tc.ResourceGroupLocation
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
	})

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.
	Context("Create and Delete", func() {

		It("should validate eventhubnamespace name is valid", func() {

			resourceGroupName := "t-rg-dev-eh-" + helpers.RandomString(10)
			eventhubNamespaceName := "t-ns"

			// Create the EventHubNamespace object and expect the Reconcile to be created
			eventhubNamespaceInstance := &azurev1.EventhubNamespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:      eventhubNamespaceName,
					Namespace: "default",
				},
				Spec: azurev1.EventhubNamespaceSpec{
					Location:      rgLocation,
					ResourceGroup: resourceGroupName,
				},
			}

			tc.K8sClient.Create(context.Background(), eventhubNamespaceInstance)

			eventhubNamespacedName := types.NamespacedName{Name: eventhubNamespaceName, Namespace: "default"}

			Eventually(func() bool {
				_ = tc.K8sClient.Get(context.Background(), eventhubNamespacedName, eventhubNamespaceInstance)
				return eventhubNamespaceInstance.IsSubmitted()
			}, timeout,
			).Should(BeFalse())

		})

		It("should validate resourcegroup exist before creating eventhubnamespaces", func() {

			resourceGroupName := "t-rg-dev-eh-" + helpers.RandomString(10)
			eventhubNamespaceName := "t-ns-dev-eh-" + helpers.RandomString(10)

			// Create the EventHubNamespace object and expect the Reconcile to be created
			eventhubNamespaceInstance := &azurev1.EventhubNamespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:      eventhubNamespaceName,
					Namespace: "default",
				},
				Spec: azurev1.EventhubNamespaceSpec{
					Location:      rgLocation,
					ResourceGroup: resourceGroupName,
				},
			}

			tc.K8sClient.Create(context.Background(), eventhubNamespaceInstance)

			eventhubNamespacedName := types.NamespacedName{Name: eventhubNamespaceName, Namespace: "default"}

			Eventually(func() bool {
				_ = tc.K8sClient.Get(context.Background(), eventhubNamespacedName, eventhubNamespaceInstance)
				return eventhubNamespaceInstance.IsSubmitted()
			}, timeout,
			).Should(BeFalse())

		})

		It("should create and delete namespace in k8s", func() {

			eventhubNamespaceName := "t-ns-dev-eh-" + helpers.RandomString(10)

			var err error

			// Create the Eventhub namespace object and expect the Reconcile to be created
			eventhubNamespaceInstance := &azurev1.EventhubNamespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:      eventhubNamespaceName,
					Namespace: "default",
				},
				Spec: azurev1.EventhubNamespaceSpec{
					Location:      rgLocation,
					ResourceGroup: rgName,
				},
			}

			err = tc.K8sClient.Create(context.Background(), eventhubNamespaceInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			eventhubNamespacedName := types.NamespacedName{Name: eventhubNamespaceName, Namespace: "default"}

			Eventually(func() bool {
				_ = tc.K8sClient.Get(context.Background(), eventhubNamespacedName, eventhubNamespaceInstance)
				return eventhubNamespaceInstance.HasFinalizer(eventhubNamespaceFinalizerName)
			}, timeout,
			).Should(BeTrue())

			Eventually(func() bool {
				_ = tc.K8sClient.Get(context.Background(), eventhubNamespacedName, eventhubNamespaceInstance)
				return eventhubNamespaceInstance.IsSubmitted()
			}, timeout,
			).Should(BeTrue())

			tc.K8sClient.Delete(context.Background(), eventhubNamespaceInstance)
			Eventually(func() bool {
				_ = tc.K8sClient.Get(context.Background(), eventhubNamespacedName, eventhubNamespaceInstance)
				return eventhubNamespaceInstance.IsBeingDeleted()
			}, timeout,
			).Should(BeTrue())

		})
	})
})
