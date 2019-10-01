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
	"github.com/Azure/azure-service-operator/pkg/helpers"

	. "github.com/onsi/ginkgo"

	. "github.com/onsi/ginkgo"

	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("EventHubNamespace Controller", func() {

	var rgName string
	var rgLocation string

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
		rgName = tc.resourceGroupName
		rgLocation = tc.resourceGroupLocation
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
	})

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.
	Context("Create and Delete", func() {

		It("should fail to create eventhubnamespace if resourcegroup doesn't exist", func() {

			resourceGroupName := "t-rg-dev-controller"
			eventhubNamespaceName := "t-ns-dev-eh-" + helpers.RandomString(10)

			// Create the EventHubNamespace object and expect the Reconcile to be created
			eventhubNamespaceInstance := &azurev1alpha1.EventhubNamespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:      eventhubNamespaceName,
					Namespace: "default",
				},
				Spec: azurev1.EventhubNamespaceSpec{
					Location:      rgLocation,
					ResourceGroup: resourceGroupName,
				},
			}

			err := tc.k8sClient.Create(context.Background(), eventhubNamespaceInstance)
			Expect(err).NotTo(HaveOccurred())

			eventhubNamespacedName := types.NamespacedName{Name: eventhubNamespaceName, Namespace: "default"}

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubNamespaceInstance)
				return eventhubNamespaceInstance.IsSubmitted()
			}, tc.timeout,
			).Should(BeFalse())
		})

		It("should create and delete namespace in k8s", func() {

			eventhubNamespaceName := "t-ns-dev-eh-" + helpers.RandomString(10)

			var err error

			// Create the Eventhub namespace object and expect the Reconcile to be created
			eventhubNamespaceInstance := &azurev1alpha1.EventhubNamespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:      eventhubNamespaceName,
					Namespace: "default",
				},
				Spec: azurev1.EventhubNamespaceSpec{
					Location:      rgLocation,
					ResourceGroup: rgName,
				},
			}

			err = tc.k8sClient.Create(context.Background(), eventhubNamespaceInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			eventhubNamespacedName := types.NamespacedName{Name: eventhubNamespaceName, Namespace: "default"}

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubNamespaceInstance)
				return eventhubNamespaceInstance.HasFinalizer(eventhubNamespaceFinalizerName)
			}, tc.timeout,
			).Should(BeTrue())

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubNamespaceInstance)
				return eventhubNamespaceInstance.IsSubmitted()
			}, tc.timeout,
			).Should(BeTrue())

			err = tc.k8sClient.Delete(context.Background(), eventhubNamespaceInstance)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubNamespaceInstance)
				return eventhubNamespaceInstance.IsBeingDeleted()
			}, tc.timeout,
			).Should(BeTrue())

		})
	})
})
