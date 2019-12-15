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

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

	s "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("BlobContainer Controller", func() {

	var rgLocation string
	var rgName string
	var saName string
	var containerAccessLevel s.PublicAccess

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
		rgLocation = tc.resourceGroupLocation
		rgName = tc.resourceGroupName
		saName = tc.storageAccountName
		containerAccessLevel = s.PublicAccessContainer
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
	})

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	Context("Create and Delete", func() {
		It("should fail to create a blob container if the resource group doesn't exist", func() {
			blobContainerName := "bc-" + helpers.RandomString(10)
			resourceGroupName := "rg-" + helpers.RandomString(10)

			var err error

			blobContainerInstance := &azurev1alpha1.BlobContainer{
				ObjectMeta: metav1.ObjectMeta{
					Name:      blobContainerName,
					Namespace: "default",
				},
				Spec: azurev1alpha1.BlobContainerSpec{
					Location:      rgLocation,
					ResourceGroup: resourceGroupName,
					AccountName:   saName,
					AccessLevel:   containerAccessLevel,
				},
			}

			err = tc.k8sClient.Create(context.Background(), blobContainerInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			blobContainerNamespacedName := types.NamespacedName{Name: blobContainerName, Namespace: "default"}
			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), blobContainerNamespacedName, blobContainerInstance)
				return blobContainerInstance.Status.Provisioned
			}, tc.timeout, tc.retry,
			).Should(BeFalse())
		})

		It("should fail to create a blob container if the storage account doesn't exist", func() {
			blobContainerName := "bc-" + helpers.RandomString(10)
			storageAccountName := "sa-" + helpers.RandomString(10)

			var err error

			blobContainerInstance := &azurev1alpha1.BlobContainer{
				ObjectMeta: metav1.ObjectMeta{
					Name:      blobContainerName,
					Namespace: "default",
				},
				Spec: azurev1alpha1.BlobContainerSpec{
					Location:      rgLocation,
					ResourceGroup: rgName,
					AccountName:   storageAccountName,
					AccessLevel:   containerAccessLevel,
				},
			}

			err = tc.k8sClient.Create(context.Background(), blobContainerInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			blobContainerNamespacedName := types.NamespacedName{Name: blobContainerName, Namespace: "default"}
			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), blobContainerNamespacedName, blobContainerInstance)
				return blobContainerInstance.Status.Provisioned
			}, tc.timeout, tc.retry,
			).Should(BeFalse())
		})

		It("should create and delete a blob container if the resource group and storage account exist", func() {
			blobContainerName := "bc-" + helpers.RandomString(10)

			var err error

			blobContainerInstance := &azurev1alpha1.BlobContainer{
				ObjectMeta: metav1.ObjectMeta{
					Name:      blobContainerName,
					Namespace: "default",
				},
				Spec: azurev1alpha1.BlobContainerSpec{
					Location:      rgLocation,
					ResourceGroup: rgName,
					AccountName:   saName,
					AccessLevel:   containerAccessLevel,
				},
			}

			err = tc.k8sClient.Create(context.Background(), blobContainerInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			blobContainerNamespacedName := types.NamespacedName{Name: blobContainerName, Namespace: "default"}

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), blobContainerNamespacedName, blobContainerInstance)
				return blobContainerInstance.HasFinalizer(blobContainerFinalizerName)
			}, tc.timeout, tc.retry,
			).Should(BeTrue())

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), blobContainerNamespacedName, blobContainerInstance)
				return blobContainerInstance.Status.Provisioned
			}, tc.timeout, tc.retry,
			).Should(BeTrue())

			err = tc.k8sClient.Delete(context.Background(), blobContainerInstance)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), blobContainerNamespacedName, blobContainerInstance)
				return blobContainerInstance.IsBeingDeleted()
			}, tc.timeout, tc.retry,
			).Should(BeTrue())
		})
	})
})
