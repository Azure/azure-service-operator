// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all storage

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	config "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/go-autorest/autorest/to"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestStorageControllerHappyPathWithoutNetworkRule(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	StorageAccountName := GenerateAlphaNumTestResourceName("sadev")
	// Create the ResourceGroup object and expect the Reconcile to be created
	saInstance := &azurev1alpha1.Storage{
		ObjectMeta: metav1.ObjectMeta{
			Name:      StorageAccountName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.StorageSpec{
			Location:      tc.resourceGroupLocation,
			ResourceGroup: tc.resourceGroupName,
			Sku: azurev1alpha1.StorageSku{
				Name: "Standard_RAGRS",
			},
			Kind:                   "StorageV2",
			AccessTier:             "Hot",
			EnableHTTPSTrafficOnly: to.BoolPtr(true),
		},
	}
	// create rg
	EnsureInstance(ctx, t, tc, saInstance)
	// delete rg
	EnsureDelete(ctx, t, tc, saInstance)
}
