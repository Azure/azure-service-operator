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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("EventHub Controller", func() {
	const timeout = time.Second * 240

	var rgName string
	var ehnName string

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
		rgName = resourceGroupName
		ehnName = eventhubNamespaceName
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
	})

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.
	Context("Create and Delete", func() {
		It("should validate eventhubnamespaces exist before creating eventhubs", func() {

			eventhubName := "t-eh-" + helpers.RandomString(10)

			// Create the EventHub object and expect the Reconcile to be created
			eventhubInstance := &azurev1.Eventhub{
				ObjectMeta: metav1.ObjectMeta{
					Name:      eventhubName,
					Namespace: "default",
				},
				Spec: azurev1.EventhubSpec{
					Location:      "westus",
					Namespace:     "t-ns-dev-eh-" + helpers.RandomString(10),
					ResourceGroup: "t-rg-dev-eh-" + helpers.RandomString(10),
					Properties: azurev1.EventhubProperties{
						MessageRetentionInDays: 7,
						PartitionCount:         1,
					},
				},
			}

			k8sClient.Create(context.Background(), eventhubInstance)

			eventhubNamespacedName := types.NamespacedName{Name: eventhubName, Namespace: "default"}

			Eventually(func() bool {
				_ = k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubInstance)
				return eventhubInstance.IsSubmitted()
			}, timeout,
			).Should(BeFalse())
		})

		It("should create and delete eventhubs", func() {

			eventhubName := "t-eh-" + helpers.RandomString(10)

			var err error

			// Create the EventHub object and expect the Reconcile to be created
			eventhubInstance := &azurev1.Eventhub{
				ObjectMeta: metav1.ObjectMeta{
					Name:      eventhubName,
					Namespace: "default",
				},
				Spec: azurev1.EventhubSpec{
					Location:      "westus",
					Namespace:     ehnName,
					ResourceGroup: rgName,
					Properties: azurev1.EventhubProperties{
						MessageRetentionInDays: 7,
						PartitionCount:         1,
					},
					AuthorizationRule: azurev1.EventhubAuthorizationRule{
						Name:   "RootManageSharedAccessKey",
						Rights: []string{"Listen"},
					},
				},
			}

			err = k8sClient.Create(context.Background(), eventhubInstance)
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
					"eventhubName":              []byte(eventhubName),
				},
				Type: "Opaque",
			}

			err = k8sClient.Create(context.Background(), csecret)
			Expect(err).NotTo(HaveOccurred())

			//get secret from k8s
			secret := &v1.Secret{}
			err = k8sClient.Get(context.Background(), types.NamespacedName{Name: eventhubName, Namespace: eventhubInstance.Namespace}, secret)
			Expect(err).NotTo(HaveOccurred())
			Expect(secret.Data).To(Equal(csecret.Data))
			Expect(secret.ObjectMeta).To(Equal(csecret.ObjectMeta))

			k8sClient.Delete(context.Background(), eventhubInstance)
			Eventually(func() bool {
				_ = k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubInstance)
				return eventhubInstance.IsBeingDeleted()
			}, timeout,
			).Should(BeTrue())

		})

		It("should create and delete eventhubs with custom secret name", func() {

			eventhubName := "t-eh-" + helpers.RandomString(10)
			secretName := "secret-" + eventhubName

			var err error

			// Create the EventHub object and expect the Reconcile to be created
			eventhubInstance := &azurev1.Eventhub{
				ObjectMeta: metav1.ObjectMeta{
					Name:      eventhubName,
					Namespace: "default",
				},
				Spec: azurev1.EventhubSpec{
					Location:      "westus",
					Namespace:     ehnName,
					ResourceGroup: rgName,
					Properties: azurev1.EventhubProperties{
						MessageRetentionInDays: 7,
						PartitionCount:         1,
					},
					AuthorizationRule: azurev1.EventhubAuthorizationRule{
						Name:   "RootManageSharedAccessKey",
						Rights: []string{"Listen"},
					},
					SecretName: secretName,
				},
			}

			err = k8sClient.Create(context.Background(), eventhubInstance)
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

			//create secret in k8s
			csecret := &v1.Secret{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Secret",
					APIVersion: "apps/v1beta1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      secretName,
					Namespace: "default",
				},
				Data: map[string][]byte{
					"primaryconnectionstring":   []byte("primaryConnectionValue"),
					"secondaryconnectionstring": []byte("secondaryConnectionValue"),
					"primaryKey":                []byte("primaryKeyValue"),
					"secondaryKey":              []byte("secondaryKeyValue"),
					"sharedaccesskey":           []byte("sharedAccessKeyValue"),
					"eventhubnamespace":         []byte(eventhubInstance.Namespace),
					"eventhubName":              []byte(eventhubName),
				},
				Type: "Opaque",
			}

			err = k8sClient.Create(context.Background(), csecret)
			Expect(err).NotTo(HaveOccurred())

			//get secret from k8s
			secret := &v1.Secret{}
			err = k8sClient.Get(context.Background(), types.NamespacedName{Name: secretName, Namespace: eventhubInstance.Namespace}, secret)
			Expect(err).NotTo(HaveOccurred())
			Expect(secret.Data).To(Equal(csecret.Data))
			Expect(secret.ObjectMeta).To(Equal(csecret.ObjectMeta))

			k8sClient.Delete(context.Background(), eventhubInstance)
			Eventually(func() bool {
				_ = k8sClient.Get(context.Background(), eventhubNamespacedName, eventhubInstance)
				return eventhubInstance.IsBeingDeleted()
			}, timeout,
			).Should(BeTrue())

		})
	})
})
