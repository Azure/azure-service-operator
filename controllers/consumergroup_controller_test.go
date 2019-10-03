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

	"github.com/Azure/azure-service-operator/pkg/helpers"

	. "github.com/onsi/ginkgo"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("ConsumerGroup Controller", func() {

	var rgName string
	var ehnName string
	var ehName string

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
		rgName = tc.resourceGroupName
		ehnName = tc.eventhubNamespaceName
		ehName = tc.eventhubName
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

			consumerGroupName := "t-cg-" + helpers.RandomString(10)

			var err error

			// Create the consumer group object and expect the Reconcile to be created
			consumerGroupInstance := &azurev1.ConsumerGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      consumerGroupName,
					Namespace: "default",
				},
				Spec: azurev1.ConsumerGroupSpec{
					NamespaceName:     ehnName,
					ResourceGroupName: rgName,
					EventhubName:      ehName,
				},
			}

			err = tc.k8sClient.Create(context.Background(), consumerGroupInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			consumerGroupNamespacedName := types.NamespacedName{Name: consumerGroupName, Namespace: "default"}

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), consumerGroupNamespacedName, consumerGroupInstance)
				return consumerGroupInstance.HasFinalizer(consumerGroupFinalizerName)
			}, tc.timeout,
			).Should(BeTrue())

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), consumerGroupNamespacedName, consumerGroupInstance)
				return consumerGroupInstance.IsSubmitted()
			}, tc.timeout,
			).Should(BeTrue())

			err = tc.k8sClient.Delete(context.Background(), consumerGroupInstance)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), consumerGroupNamespacedName, consumerGroupInstance)
				return consumerGroupInstance.IsBeingDeleted()
			}, tc.timeout,
			).Should(BeTrue())

		})
	})
})
