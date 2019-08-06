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

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	resoucegroupsresourcemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/resourcegroups"
	. "github.com/onsi/ginkgo"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("EventHub Controller", func() {

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
		// It("should validate eventhubnamespaces exist before creating eventhubs", func() {

		// 	resourceGroupName := "t-rg-dev-eh-" + helpers.RandomString(10)
		// 	eventhubNamespaceName := "t-ns-dev-eh-" + helpers.RandomString(10)
		// 	eventhubName := "t-eh-" + helpers.RandomString(10)

		// 	// Create the EventHub object and expect the Reconcile to be created
		// 	eventhubInstance := &azurev1.Eventhub{
		// 		ObjectMeta: metav1.ObjectMeta{
		// 			Name:      eventhubName,
		// 			Namespace: "default",
		// 		},
		// 		Spec: azurev1.EventhubSpec{
		// 			Location:      "westus",
		// 			Namespace:     eventhubNamespaceName,
		// 			ResourceGroup: resourceGroupName,
		// 			Properties: azurev1.EventhubProperties{
		// 				MessageRetentionInDays: 7,
		// 				PartitionCount:         1,
		// 			},
		// 		},
		// 	}

		// 	k8sClient.Create(context.Background(), eventhubInstance)

		// 	time.Sleep(60 * time.Second)

		// 	eventhubNamespacedName := types.NamespacedName{Name: eventhubName, Namespace: "default"}

		// 	Eventually(func() bool {
		// 		_ = k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubInstance)
		// 		return eventhubInstance.IsSubmitted()
		// 	}, timeout,
		// 	).Should(BeFalse())

		// })

		It("should create and delete eventhubs", func() {

			resourceGroupName := "t-rg-dev-eh-" + helpers.RandomString(10)
			eventhubNamespaceName := "t-ns-dev-eh-" + helpers.RandomString(10)
			eventhubName := "t-eh-" + helpers.RandomString(10)

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

			time.Sleep(30 * time.Second)
			// The instance object may not be a valid object because it might be missing some required fields.
			// Please modify the instance object by adding required fields and then remove the following if statement.
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			eventhubNamespacedName := types.NamespacedName{Name: eventhubName, Namespace: "default"}

			Eventually(func() bool {
				_ = k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubInstance)
				return eventhubInstance.HasFinalizer(eventhubFinalizerName)
			}, timeout,
			).Should(BeTrue())

			Eventually(func() bool {
				_ = k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubInstance)
				return eventhubInstance.IsSubmitted()
			}, timeout,
			).Should(BeTrue())

			time.Sleep(2 * time.Second)

			//create secret in k8s
			csecret := &v1.Secret{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Secret",
					APIVersion: "apps/v1beta1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      eventhubName,
					Namespace: "default",
				},
				Data: map[string][]byte{
					"primaryconnectionstring":   []byte("primaryConnectionValue"),
					"secondaryconnectionstring": []byte("secondaryConnectionValue"),
					"primaryKey":                []byte("primaryKeyValue"),
					"secondaryKey":              []byte("secondaryKeyValue"),
					"sharedaccesskey":           []byte("sharedAccessKeyValue"),
					"eventhubnamespace":         []byte(eventhubInstance.Namespace),
				},
				Type: "Opaque",
			}

			err = k8sClient.Create(context.Background(), csecret)
			Expect(err).NotTo(HaveOccurred())

			time.Sleep(2 * time.Second)

			//get secret from k8s
			secret := &v1.Secret{}
			err = k8sClient.Get(context.Background(), types.NamespacedName{Name: eventhubName, Namespace: eventhubInstance.Namespace}, secret)
			Expect(err).NotTo(HaveOccurred())
			Expect(secret.Data).To(Equal(csecret.Data))
			Expect(secret.ObjectMeta).To(Equal(csecret.ObjectMeta))

			time.Sleep(2 * time.Second)

			k8sClient.Delete(context.Background(), eventhubInstance)
			Eventually(func() bool {
				_ = k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubInstance)
				return eventhubInstance.IsBeingDeleted()
			}, timeout,
			).Should(BeTrue())

			time.Sleep(2 * time.Second)

			_, err = resoucegroupsresourcemanager.DeleteGroup(context.Background(), resourceGroupName)
			Expect(err).NotTo(HaveOccurred())

		})
	})
})
