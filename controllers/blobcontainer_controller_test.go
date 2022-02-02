// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || blobcontainer
// +build all blobcontainer

package controllers

import (
	"context"
	"testing"

	s "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/Azure/go-autorest/autorest/to"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/api/v1alpha2"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestBlobContainerControlleNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	var rgLocation string
	var saName string
	var containerAccessLevel s.PublicAccess

	// Add any setup steps that needs to be executed before each test
	rgLocation = tc.resourceGroupLocation
	saName = GenerateAlphaNumTestResourceName("blobsa")
	containerAccessLevel = s.PublicAccessContainer

	blobContainerName := GenerateTestResourceNameWithRandom("bc", 10)
	resourceGroupName := GenerateTestResourceNameWithRandom("rg", 10)

	// Create Storage account
	saInstance := &v1alpha1.StorageAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      saName,
			Namespace: "default",
		},
		Spec: v1alpha1.StorageAccountSpec{
			Location:      tc.resourceGroupLocation,
			ResourceGroup: tc.resourceGroupName,
			Sku: v1alpha1.StorageAccountSku{
				Name: "Standard_RAGRS",
			},
			Kind:                   "StorageV2",
			AccessTier:             "Hot",
			EnableHTTPSTrafficOnly: to.BoolPtr(true),
		},
	}
	EnsureInstance(ctx, t, tc, saInstance)

	blobContainerInstance := &v1alpha2.BlobContainer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      blobContainerName,
			Namespace: "default",
		},
		Spec: v1alpha2.BlobContainerSpec{
			Location:      rgLocation,
			ResourceGroup: resourceGroupName,
			AccountName:   saName,
			AccessLevel:   containerAccessLevel,
		},
	}
	EnsureInstanceWithResult(ctx, t, tc, blobContainerInstance, errhelp.ResourceGroupNotFoundErrorCode, false)

	EnsureDelete(ctx, t, tc, blobContainerInstance)
}

func TestTestBlobContainerControllerNoStorageAccount(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	var rgLocation string
	var rgName string
	var containerAccessLevel s.PublicAccess

	// Add any setup steps that needs to be executed before each test
	rgLocation = tc.resourceGroupLocation
	rgName = tc.resourceGroupName
	containerAccessLevel = s.PublicAccessContainer

	blobContainerName := GenerateTestResourceNameWithRandom("bc", 10)
	storageAccountName := helpers.FillWithRandom(GenerateAlphaNumTestResourceName("sa"), 24)

	blobContainerInstance := &v1alpha2.BlobContainer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      blobContainerName,
			Namespace: "default",
		},
		Spec: v1alpha2.BlobContainerSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
			AccountName:   storageAccountName,
			AccessLevel:   containerAccessLevel,
		},
	}
	EnsureInstanceWithResult(ctx, t, tc, blobContainerInstance, errhelp.ParentNotFoundErrorCode, false)

	EnsureDelete(ctx, t, tc, blobContainerInstance)
}
