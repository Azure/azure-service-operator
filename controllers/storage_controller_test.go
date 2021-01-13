// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all appinsights

package controllers

import (
	"context"
	"fmt"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

func TestStorageAccountController(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	assert := assert.New(t)

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

	// Make sure the secret is created
	var keyName string
	if tc.secretClient.GetSecretNamingVersion() == secrets.SecretNamingV1 {
		keyName = fmt.Sprintf("storageaccount-%s-%s", instance.Spec.ResourceGroup, instance.Name)
	} else {
		keyName = instance.Name
	}
	secretKey := secrets.SecretKey{Name: keyName, Namespace: instance.Namespace, Kind: "storageaccount"}
	secret, err := tc.secretClient.Get(ctx, secretKey)
	assert.NoError(err)

	assert.NotEmpty(string(secret["StorageAccountName"]))
	assert.NotEmpty(string(secret["connectionString0"]))
	assert.NotEmpty(string(secret["key0"]))
	assert.NotEmpty(string(secret["connectionString1"]))
	assert.NotEmpty(string(secret["key1"]))

	EnsureDelete(ctx, t, tc, instance)
}
