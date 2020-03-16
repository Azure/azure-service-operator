// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all adlsgen2

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/stretchr/testify/assert"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestADLSFilesystemControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	assert := assert.New(t)
	ctx := context.Background()

	var saName string = tc.storageAccountName

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

	fileSystemName := "adls-filesystem-" + helpers.RandomString(10)
	resouceGroupName := "rg-" + helpers.RandomString(10)

	var err error

	fileSystemInstance := &azurev1alpha1.AzureDataLakeGen2FileSystem{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fileSystemName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureDataLakeGen2FileSystemSpec{
			StorageAccountName: saName,
			ResourceGroupName:  resouceGroupName,
		},
	}

	err = tc.k8sClient.Create(ctx, fileSystemInstance)
	assert.Equal(nil, err, "create filesystem instance in k8s")

	// @todo update this to check something other than status.Provisioned
	fileSystemNamespacedName := types.NamespacedName{Name: fileSystemName, Namespace: "default"}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, fileSystemNamespacedName, fileSystemInstance)
		return fileSystemInstance.Status.Provisioned == false
	}, tc.timeout, tc.retry, "wait for filesystem to have false for provisioned bool")

	// Delete should still appear successful
	err = tc.k8sClient.Delete(ctx, fileSystemInstance)
	assert.Equal(nil, err, "delete filesystem instance in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, fileSystemNamespacedName, fileSystemInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for filesystem to be gone")

}

func TestADLSFilesystemControllerNoStorageAccount(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	assert := assert.New(t)
	ctx := context.Background()

	var rgName string = tc.resourceGroupName
	fileSystemName := "adls-filesystem-" + helpers.RandomString(10)
	storageAccountName := "sa-" + helpers.RandomString(10)

	var err error

	fileSystemInstance := &azurev1alpha1.AzureDataLakeGen2FileSystem{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fileSystemName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureDataLakeGen2FileSystemSpec{
			StorageAccountName: storageAccountName,
			ResourceGroupName:  rgName,
		},
	}

	err = tc.k8sClient.Create(ctx, fileSystemInstance)
	assert.Equal(nil, err, "create filesystem instance in k8s")

	fileSystemNamespacedName := types.NamespacedName{Name: fileSystemName, Namespace: "default"}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, fileSystemNamespacedName, fileSystemInstance)
		return fileSystemInstance.Status.Provisioned == false
	}, tc.timeout, tc.retry, "wait for filesystem to have false for provisioned bool")

	// Delete should still appear successful
	err = tc.k8sClient.Delete(context.Background(), fileSystemInstance)
	assert.Equal(nil, err, "delete filesystem instance in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, fileSystemNamespacedName, fileSystemInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for filesystem to be gone")
}

// func TestADLSFilesystemControllerHappyPath(t *testing.T) {
// 	t.Parallel()
// 	RegisterTestingT(t)
// 	defer PanicRecover()

// 	var rgName string = tc.resourceGroupName
// 	var saName string = tc.storageAccountName

// 	fileSystemName := "adls-filesystem-" + helpers.RandomString(10)

// 	var err error

// 	fileSystemInstance := &azurev1alpha1.AzureDataLakeGen2FileSystem{
// 		ObjectMeta: metav1.ObjectMeta{
// 			Name:      fileSystemName,
// 			Namespace: "default",
// 		},
// 		Spec: azurev1alpha1.AzureDataLakeGen2FileSystemSpec{
// 			StorageAccountName: saName,
// 			ResourceGroupName:  rgName,
// 		},
// 	}

// 	err = tc.k8sClient.Create(context.Background(), fileSystemInstance)
// 	Expect(apierrors.IsInvalid(err)).To(Equal(false))
// 	Expect(err).NotTo(HaveOccurred())

// 	fileSystemNamespacedName := types.NamespacedName{Name: fileSystemName, Namespace: "default"}

// 	Eventually(func() bool {
// 		_ = tc.k8sClient.Get(context.Background(), fileSystemNamespacedName, fileSystemInstance)
// 		return fileSystemInstance.HasFinalizer(fileSystemFinalizerName)
// 	}, tc.timeout, tc.retry,
// 	).Should(BeTrue())

// 	Eventually(func() bool {
// 		_ = tc.k8sClient.Get(context.Background(), fileSystemNamespacedName, fileSystemInstance)
// 		// return fileSystemInstance.Status.Provisioned
// 		// this is wrong @todo fix
// 		return fileSystemInstance.IsSubmitted()
// 	}, tc.timeout, tc.retry,
// 	).Should(BeTrue())

// 	err = tc.k8sClient.Delete(context.Background(), fileSystemInstance)
// 	Expect(err).NotTo(HaveOccurred())

// 	Eventually(func() bool {
// 		err = tc.k8sClient.Get(context.Background(), fileSystemNamespacedName, fileSystemInstance)
// 		//return apierrors.IsNotFound(err)
// 		return fileSystemInstance.IsBeingDeleted()
// 	}, tc.timeout, tc.retry,
// 	).Should(BeTrue())

// }
