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
	"time"

	eventhubsmanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/eventhubs"

	. "github.com/onsi/ginkgo"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("EventHub Controller", func() {

	const timeout = time.Second * 60

	var rgName string
	var rgLocation string
	var ehnName string
	var saName string
	var bcName string

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
		rgName = tc.ResourceGroupName
		rgLocation = tc.ResourcegroupLocation
		ehnName = tc.EventhubNamespaceName
		saName = tc.StorageAccountName
		bcName = tc.BlobContainerName
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

			tc.K8sClient.Create(context.Background(), eventhubInstance)

			eventhubNamespacedName := types.NamespacedName{Name: eventhubName, Namespace: "default"}

			Eventually(func() bool {
				_ = tc.K8sClient.Get(context.Background(), eventhubNamespacedName, eventhubInstance)
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

			err = tc.K8sClient.Create(context.Background(), eventhubInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			eventhubNamespacedName := types.NamespacedName{Name: eventhubName, Namespace: "default"}

			Eventually(func() bool {
				_ = tc.K8sClient.Get(context.Background(), eventhubNamespacedName, eventhubInstance)
				return eventhubInstance.HasFinalizer(eventhubFinalizerName)
			}, timeout,
			).Should(BeTrue())

			Eventually(func() bool {
				_ = tc.K8sClient.Get(context.Background(), eventhubNamespacedName, eventhubInstance)
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

			err = tc.K8sClient.Create(context.Background(), csecret)
			Expect(err).NotTo(HaveOccurred())

			//get secret from k8s
			secret := &v1.Secret{}
			err = tc.K8sClient.Get(context.Background(), types.NamespacedName{Name: eventhubName, Namespace: eventhubInstance.Namespace}, secret)
			Expect(err).NotTo(HaveOccurred())
			Expect(secret.Data).To(Equal(csecret.Data))
			Expect(secret.ObjectMeta).To(Equal(csecret.ObjectMeta))

			tc.K8sClient.Delete(context.Background(), eventhubInstance)
			Eventually(func() bool {
				_ = tc.K8sClient.Get(context.Background(), eventhubNamespacedName, eventhubInstance)
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
					Location:      rgLocation,
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

			err = tc.K8sClient.Create(context.Background(), eventhubInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			eventhubNamespacedName := types.NamespacedName{Name: eventhubName, Namespace: "default"}

			Eventually(func() bool {
				_ = tc.K8sClient.Get(context.Background(), eventhubNamespacedName, eventhubInstance)
				return eventhubInstance.HasFinalizer(eventhubFinalizerName)
			}, timeout,
			).Should(BeTrue())

			Eventually(func() bool {
				_ = tc.K8sClient.Get(context.Background(), eventhubNamespacedName, eventhubInstance)
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

			err = tc.K8sClient.Create(context.Background(), csecret)
			Expect(err).NotTo(HaveOccurred())

			//get secret from k8s
			secret := v1.Secret{}
			Eventually(func() bool {
				err = tc.K8sClient.Get(context.Background(), types.NamespacedName{Name: secretName, Namespace: eventhubInstance.Namespace}, &secret)
				if err != nil {
					return false
				}
				Expect(secret.Data).To(Equal(csecret.Data))
				Expect(secret.ObjectMeta).To(Equal(csecret.ObjectMeta))
				return true
			}, 60).Should(BeTrue())

			tc.K8sClient.Delete(context.Background(), eventhubInstance)
			Eventually(func() bool {
				_ = tc.K8sClient.Get(context.Background(), eventhubNamespacedName, eventhubInstance)
				return eventhubInstance.IsBeingDeleted()
			}, timeout,
			).Should(BeTrue())

		})

		It("should create and delete event hubs with capture", func() {

			eventHubName := "t-eh-" + helpers.RandomString(10)

			var err error

			// Create the EventHub object and expect the Reconcile to be created
			eventHubInstance := &azurev1.Eventhub{
				ObjectMeta: metav1.ObjectMeta{
					Name:      eventHubName,
					Namespace: "default",
				},
				Spec: azurev1.EventhubSpec{
					Location:      rgLocation,
					Namespace:     ehnName,
					ResourceGroup: rgName,
					Properties: azurev1.EventhubProperties{
						MessageRetentionInDays: 7,
						PartitionCount:         1,
						CaptureDescription: azurev1.CaptureDescription{
							Destination: azurev1.Destination{
								ArchiveNameFormat: "{Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}",
								BlobContainer:     bcName,
								Name:              "EventHubArchive.AzureBlockBlob",
								StorageAccount: azurev1.StorageAccount{
									ResourceGroup: rgName,
									AccountName:   saName,
								},
							},
							Enabled:           true,
							SizeLimitInBytes:  524288000,
							IntervalInSeconds: 300,
						},
					},
				},
			}

			err = tc.K8sClient.Create(context.Background(), eventHubInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			eventHubNamespacedName := types.NamespacedName{Name: eventHubName, Namespace: "default"}

			Eventually(func() bool {
				_ = tc.K8sClient.Get(context.Background(), eventHubNamespacedName, eventHubInstance)
				return eventHubInstance.HasFinalizer(eventhubFinalizerName)
			}, timeout,
			).Should(BeTrue())

			Eventually(func() bool {
				_ = tc.K8sClient.Get(context.Background(), eventHubNamespacedName, eventHubInstance)
				return eventHubInstance.IsSubmitted()
			}, timeout,
			).Should(BeTrue())

			Eventually(func() bool {
				hub, _ := eventhubsmanager.GetHub(context.Background(), rgName, ehnName, eventHubName)
				if hub.Properties == nil || hub.CaptureDescription == nil || hub.CaptureDescription.Enabled == nil {
					return false
				}
				return *hub.CaptureDescription.Enabled
			}, timeout,
			).Should(BeTrue())

			tc.K8sClient.Delete(context.Background(), eventHubInstance)
			Eventually(func() bool {
				_ = tc.K8sClient.Get(context.Background(), eventHubNamespacedName, eventHubInstance)
				return eventHubInstance.IsBeingDeleted()
			}, timeout,
			).Should(BeTrue())

		})
	})
})
