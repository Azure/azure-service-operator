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
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"

	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	storagemanager "github.com/Azure/azure-service-operator/pkg/resourcemanager/storage"

	"time"

	. "github.com/onsi/ginkgo"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("EventHub Controller (with capture)", func() {

	const timeout = time.Second * 240
	var storageAccountName = "tsadeveh" + helpers.RandomString(10)
	var blobContainerName = "t-bc-dev-eh-" + helpers.RandomString(10)
	var location = config.DefaultLocation()

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
		_, _ = storagemanager.CreateStorageAccountAndWait(context.Background(), resourceGroupName, storageAccountName, "Storage", location)
		_, _ = storagemanager.CreateBlobContainer(context.Background(), resourceGroupName, storageAccountName, blobContainerName)

	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
		_, _ = storagemanager.DeleteBlobContainer(context.Background(), resourceGroupName, storageAccountName, blobContainerName)
		_, _ = storagemanager.DeleteStorageAccount(context.Background(), resourceGroupName, storageAccountName)
	})

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.
	Context("Create and Delete", func() {
		It("should create and delete event hubs with capture", func() {

			resourceGroupName = "t-rg-dev-controller"
			eventhubNamespaceName = "t-ns-dev-eh-ns"
			eventHubName := "t-eh-" + helpers.RandomString(10)

			var err error

			// Create the EventHub object and expect the Reconcile to be created
			eventHubInstance := &azurev1.Eventhub{
				ObjectMeta: metav1.ObjectMeta{
					Name:      eventHubName,
					Namespace: "default",
				},
				Spec: azurev1.EventhubSpec{
					Location:      location,
					Namespace:     eventhubNamespaceName,
					ResourceGroup: resourceGroupName,
					Properties: azurev1.EventhubProperties{
						MessageRetentionInDays: 7,
						PartitionCount:         1,
						CaptureDescription: azurev1.CaptureDescription{
							Destination: azurev1.Destination{
								ArchiveNameFormat: "{Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}",
								BlobContainer:     blobContainerName,
								Name:              "EventHubArchive.AzureBlockBlob",
								StorageAccount: azurev1.StorageAccount{
									ResourceGroup: resourceGroupName,
									AccountName:   storageAccountName,
								},
							},
							Enabled:           true,
							SizeLimitInBytes:  524288000,
							IntervalInSeconds: 300,
						},
					},
				},
			}

			err = k8sClient.Create(context.Background(), eventHubInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			eventHubNamespacedName := types.NamespacedName{Name: eventHubName, Namespace: "default"}

			Eventually(func() bool {
				_ = k8sClient.Get(context.Background(), eventHubNamespacedName, eventHubInstance)
				return eventHubInstance.HasFinalizer(eventhubFinalizerName)
			}, timeout,
			).Should(BeTrue())

			Eventually(func() bool {
				_ = k8sClient.Get(context.Background(), eventHubNamespacedName, eventHubInstance)
				return eventHubInstance.IsSubmitted()
			}, timeout,
			).Should(BeTrue())

			//create secret in k8s
			csecret := &v1.Secret{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Secret",
					APIVersion: "apps/v1beta1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      eventHubName,
					Namespace: "default",
				},
				Data: map[string][]byte{
					"primaryconnectionstring":   []byte("primaryConnectionValue"),
					"secondaryconnectionstring": []byte("secondaryConnectionValue"),
					"primaryKey":                []byte("primaryKeyValue"),
					"secondaryKey":              []byte("secondaryKeyValue"),
					"sharedaccesskey":           []byte("sharedAccessKeyValue"),
					"eventHubnamespace":         []byte(eventHubInstance.Namespace),
				},
				Type: "Opaque",
			}

			err = k8sClient.Create(context.Background(), csecret)
			Expect(err).NotTo(HaveOccurred())

			//get secret from k8s
			secret := &v1.Secret{}
			err = k8sClient.Get(context.Background(), types.NamespacedName{Name: eventHubName, Namespace: eventHubInstance.Namespace}, secret)
			Expect(err).NotTo(HaveOccurred())
			Expect(secret.Data).To(Equal(csecret.Data))
			Expect(secret.ObjectMeta).To(Equal(csecret.ObjectMeta))

			k8sClient.Delete(context.Background(), eventHubInstance)
			Eventually(func() bool {
				_ = k8sClient.Get(context.Background(), eventHubNamespacedName, eventHubInstance)
				return eventHubInstance.IsBeingDeleted()
			}, timeout,
			).Should(BeTrue())

		})
	})
})
