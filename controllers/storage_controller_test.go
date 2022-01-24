// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || appinsights
// +build all appinsights

package controllers

import (
	"context"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
)

func TestStorageAccountController(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	rgName := tc.resourceGroupName
	rgLocation := tc.resourceGroupLocation
	storageName := "storageacct" + helpers.RandomString(6)

	instance := &v1alpha1.StorageAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      storageName,
			Namespace: "default",
		},
		Spec: v1alpha1.StorageAccountSpec{
			Kind:          "BlobStorage",
			Location:      rgLocation,
			ResourceGroup: rgName,
			Sku: v1alpha1.StorageAccountSku{
				Name: "Standard_LRS",
			},
			AccessTier:             "Hot",
			EnableHTTPSTrafficOnly: to.BoolPtr(true),
		},
	}

	EnsureInstance(ctx, t, tc, instance)

	EnsureDelete(ctx, t, tc, instance)
}
