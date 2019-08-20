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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("ConsumerGroup Controller", func() {

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
		It("should create and delete consumer groups", func() {

			resourceGroupName := "t-rg-dev-eh-" + helpers.RandomString(10)
			eventhubNamespaceName := "t-ns-dev-eh-" + helpers.RandomString(10)
			eventhubName := "t-eh-" + helpers.RandomString(10)
			consumerGroupName := "t-cg-" + helpers.RandomString(10)

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

			time.Sleep(30 * time.Second)

			// Create the Eventhub namespace object and expect the Reconcile to be created
			eventhubNamespaceInstance := &azurev1.EventhubNamespace{
				ObjectMeta: metav1.ObjectMeta{
					Name:      eventhubNamespaceName,
					Namespace: "default",
				},
				Spec: azurev1.EventhubNamespaceSpec{
					Location:      "westus",
					ResourceGroup: resourceGroupName,
				},
			}

			err = k8sClient.Create(context.Background(), eventhubNamespaceInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			time.Sleep(30 * time.Second)

			// Create the EventHub object and expect the Reconcile to be created
			eventhubInstance := &azurev1.Eventhub{
				ObjectMeta: metav1.ObjectMeta{
					Name:      eventhubName,
					Namespace: "default",
				},
				Spec: azurev1.EventhubSpec{
					Location:      "westus",
					Namespace:     eventhubNamespaceName,
					ResourceGroup: resourceGroupName,
					Properties: azurev1.EventhubProperties{
						MessageRetentionInDays: 7,
						PartitionCount:         1,
					},
				},
			}

			err = k8sClient.Create(context.Background(), eventhubInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			// Create the consumer group object and expect the Reconcile to be created
			consumerGroupInstance := &azurev1.ConsumerGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      consumerGroupName,
					Namespace: "default",
				},
				Spec: azurev1.ConsumerGroupSpec{
					NamespaceName:     eventhubNamespaceName,
					ResourceGroupName: resourceGroupName,
					EventhubName:      eventhubName,
				},
			}

			err = k8sClient.Create(context.Background(), consumerGroupInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			consumerGroupNamespacedName := types.NamespacedName{Name: consumerGroupName, Namespace: "default"}

			Eventually(func() bool {
				_ = k8sClient.Get(context.Background(), consumerGroupNamespacedName, consumerGroupInstance)
				return consumerGroupInstance.HasFinalizer(consumerGroupFinalizerName)
			}, timeout,
			).Should(BeTrue())

			Eventually(func() bool {
				_ = k8sClient.Get(context.Background(), consumerGroupNamespacedName, consumerGroupInstance)
				return consumerGroupInstance.IsSubmitted()
			}, timeout,
			).Should(BeTrue())

			k8sClient.Delete(context.Background(), consumerGroupInstance)
			Eventually(func() bool {
				_ = k8sClient.Get(context.Background(), consumerGroupNamespacedName, consumerGroupInstance)
				return consumerGroupInstance.IsBeingDeleted()
			}, timeout,
			).Should(BeTrue())

		})
	})
})
