// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all blobcontainer

package controllers

import (
	"context"
	"strings"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"

	s "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestBlobContainerControlleNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover()
	ctx := context.Background()
	assert := assert.New(t)

	var rgLocation string
	var saName string
	var containerAccessLevel s.PublicAccess

	// Add any setup steps that needs to be executed before each test
	rgLocation = tc.resourceGroupLocation
	saName = tc.storageAccountName
	containerAccessLevel = s.PublicAccessContainer

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

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

	err = tc.k8sClient.Create(ctx, blobContainerInstance)
	assert.Equal(nil, err, "create blobcontainer in k8s")

	blobContainerNamespacedName := types.NamespacedName{Name: blobContainerName, Namespace: "default"}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, blobContainerNamespacedName, blobContainerInstance)
		return strings.Contains(blobContainerInstance.Status.Message, errhelp.ResourceGroupNotFoundErrorCode)
	}, tc.timeout, tc.retry, "wait for blob to have rg not found error")

	err = tc.k8sClient.Delete(ctx, blobContainerInstance)
	assert.Equal(nil, err, "delete blobcontainer in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, blobContainerNamespacedName, blobContainerInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for blob to be not found")

}

func TestTestBlobContainerControllerNoStorageAccount(t *testing.T) {
	t.Parallel()
	defer PanicRecover()
	ctx := context.Background()
	assert := assert.New(t)

	var rgLocation string
	var rgName string
	var containerAccessLevel s.PublicAccess

	// Add any setup steps that needs to be executed before each test
	rgLocation = tc.resourceGroupLocation
	rgName = tc.resourceGroupName
	containerAccessLevel = s.PublicAccessContainer

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

	err = tc.k8sClient.Create(ctx, blobContainerInstance)
	assert.Equal(nil, err, "create blob container in k8s")

	blobContainerNamespacedName := types.NamespacedName{Name: blobContainerName, Namespace: "default"}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, blobContainerNamespacedName, blobContainerInstance)
		return strings.Contains(blobContainerInstance.Status.Message, errhelp.ParentNotFoundErrorCode)
	}, tc.timeout, tc.retry, "wait for blob to have parent not found error")

	err = tc.k8sClient.Delete(ctx, blobContainerInstance)
	assert.Equal(nil, err, "delete blob container in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, blobContainerNamespacedName, blobContainerInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for blob to be not found")

}

func TestTestBlobContainerControllerHappyPath(t *testing.T) {
	t.Parallel()
	defer PanicRecover()
	ctx := context.Background()
	assert := assert.New(t)

	var rgLocation string
	var rgName string
	var saName string
	var containerAccessLevel s.PublicAccess

	// Add any setup steps that needs to be executed before each test
	rgLocation = tc.resourceGroupLocation
	rgName = tc.resourceGroupName
	saName = tc.storageAccountName
	containerAccessLevel = s.PublicAccessContainer

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.

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

	err = tc.k8sClient.Create(ctx, blobContainerInstance)
	assert.Equal(nil, err, "create blobcontainer in k8s")

	blobContainerNamespacedName := types.NamespacedName{Name: blobContainerName, Namespace: "default"}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, blobContainerNamespacedName, blobContainerInstance)
		return blobContainerInstance.HasFinalizer(blobContainerFinalizerName)
	}, tc.timeout, tc.retry, "wait for blob to have finalizer")

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, blobContainerNamespacedName, blobContainerInstance)
		return blobContainerInstance.Status.Provisioned
	}, tc.timeout, tc.retry, "wait for blob to be provisioned")

	err = tc.k8sClient.Delete(ctx, blobContainerInstance)
	assert.Equal(nil, err, "delete blob container in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, blobContainerNamespacedName, blobContainerInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for blob to be not found")

}
