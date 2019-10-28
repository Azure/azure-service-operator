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
	"github.com/Azure/azure-service-operator/pkg/helpers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("ADLS Filesystem Controller", func() {

	var rgName string
	var saName string

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
		rgName = tc.resourceGroupName
		saName = tc.storageAccountName
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
	})

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	Context("Create and Delete", func() {
		// It("should fail to create a file system if the resource group doesn't exist", func() {
		// 	fileSystemName := "adls-filesystem-" + helpers.RandomString(10)
		// 	resouceGroupName := "rg-" + helpers.RandomString(10)

		// 	var err error

		// 	// Create the FileSystem object and expect the Reconcile to be created
		// 	fileSystemInstance := &azurev1alpha1.AzureDataLakeGen2FileSystem{ //////
		// 		ObjectMeta: metav1.ObjectMeta{
		// 			Name:      fileSystemName,
		// 			Namespace: "default",
		// 		},
		// 		Spec: azurev1alpha1.AzureDataLakeGen2FileSystemSpec{
		// 			StorageAccountName: saName,
		// 			ResourceGroupName:  resouceGroupName,
		// 		},
		// 	}

		// 	err = tc.k8sClient.Create(context.Background(), fileSystemInstance)
		// 	Expect(apierrors.IsInvalid(err)).To(Equal(false))
		// 	Expect(err).NotTo(HaveOccurred())

		// 	fileSystemNamespacedName := types.NamespacedName{Name: fileSystemName, Namespace: "default"}
		// 	Eventually(func() bool {
		// 		_ = tc.k8sClient.Get(context.Background(), fileSystemNamespacedName, fileSystemInstance)
		// 		testthing := fileSystemInstance.IsSubmitted()
		// 		return testthing
		// 	}, tc.timeout,
		// 	).Should(BeFalse())
		// })

		It("should create a file system if the resource group exists", func() {
			fileSystemName := "adls-filesystem-" + helpers.RandomString(10)

			var err error

			// Create the FileSystem object and expect the Reconcile to be created
			fileSystemInstance := &azurev1alpha1.AzureDataLakeGen2FileSystem{ //////
				ObjectMeta: metav1.ObjectMeta{
					Name:      fileSystemName,
					Namespace: "default",
				},
				Spec: azurev1alpha1.AzureDataLakeGen2FileSystemSpec{
					StorageAccountName: saName,
					ResourceGroupName:  rgName,
				},
			}

			err = tc.k8sClient.Create(context.Background(), fileSystemInstance)
			Expect(apierrors.IsInvalid(err)).To(Equal(false))
			Expect(err).NotTo(HaveOccurred())

			fileSystemNamespacedName := types.NamespacedName{Name: fileSystemName, Namespace: "default"}

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), fileSystemNamespacedName, fileSystemInstance)
				hasfinalizer := fileSystemInstance.HasFinalizer(fileSystemFinalizerName)
				return hasfinalizer
			}, tc.timeout,// 100,
			).Should(BeTrue())

			Eventually(func() bool {
				_ = tc.k8sClient.Get(context.Background(), fileSystemNamespacedName, fileSystemInstance)
				testthing := fileSystemInstance.IsSubmitted()
				return testthing
			}, tc.timeout,// 100,
			).Should(BeTrue())

		})

		// err = tc.k8sClient.Delete(context.Background(), fileSystemInstance)
		// Expect(err).NotTo(HaveOccurred())

		// Eventually(func() bool {
		// 	_ = tc.k8sClient.Get(context.Background(), fileSystemNamespacedName, fileSystemInstance)
		// 	hello123 := fileSystemInstance.IsBeingDeleted()
		// 	return hello123
		// }, tc.timeout,
		// ).Should(BeTrue())

	})
})
