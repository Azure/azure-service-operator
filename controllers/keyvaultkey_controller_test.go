// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || keyvaultkey
// +build all keyvaultkey

package controllers

import (
	"context"
	"net/http"
	"testing"
	"time"

	kvops "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	resourcemanagerkeyvaults "github.com/Azure/azure-service-operator/pkg/resourcemanager/keyvaults"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestKeyvaultKeyControllerHappyPath(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	keyVaultName := helpers.FillWithRandom(GenerateTestResourceName("kv"), 24)
	keyVaultKeyName := GenerateTestResourceNameWithRandom("kvk-dev", 10)
	const poll = time.Second * 10

	keyVaultLocation := tc.resourceGroupLocation

	keyPermissions := []string{"get", "list", "update", "delete", "recover", "backup", "restore", "create", "import"}
	accessPolicies := []azurev1alpha1.AccessPolicyEntry{
		{
			TenantID: config.GlobalCredentials().TenantID(),
			ClientID: config.GlobalCredentials().ClientID(),
			Permissions: &azurev1alpha1.Permissions{
				Keys: &keyPermissions,
			},
		}}

	// Declare KeyVault object
	keyVaultInstance := &azurev1alpha1.KeyVault{
		ObjectMeta: metav1.ObjectMeta{
			Name:      keyVaultName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.KeyVaultSpec{
			Location:       keyVaultLocation,
			ResourceGroup:  tc.resourceGroupName,
			AccessPolicies: &accessPolicies,
		},
	}

	// Create the Keyvault object and expect the Reconcile to be created
	EnsureInstance(ctx, t, tc, keyVaultInstance)

	// Prep query for get
	//keyVaultNamespacedName := types.NamespacedName{Name: keyVaultName, Namespace: "default"}

	// verify key vault exists in Azure
	assert.Eventually(func() bool {
		result, _ := tc.keyVaultManager.GetVault(ctx, tc.resourceGroupName, keyVaultInstance.Name)
		return result.Response.StatusCode == http.StatusOK
	}, tc.timeout, tc.retry, "wait for keyVaultInstance to be ready in azure")

	keyVaultKey := &azurev1alpha1.KeyVaultKey{
		ObjectMeta: metav1.ObjectMeta{
			Name:      keyVaultKeyName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.KeyVaultKeySpec{
			Location:      keyVaultLocation,
			ResourceGroup: tc.resourceGroupName,
			KeyVault:      keyVaultName,
			KeySize:       4096,
			Type:          kvops.RSA,
			Operations:    kvops.PossibleJSONWebKeyOperationValues(),
		},
	}

	// create key
	EnsureInstance(ctx, t, tc, keyVaultKey)

	kvopsclient := resourcemanagerkeyvaults.NewOpsClient(config.GlobalCredentials(), keyVaultName)

	assert.Eventually(func() bool {
		kvault, err := tc.keyVaultManager.GetVault(ctx, tc.resourceGroupName, keyVaultInstance.Name)
		if err != nil {
			return err == nil
		}
		vUrl := *kvault.Properties.VaultURI
		_, err = kvopsclient.GetKey(ctx, vUrl, keyVaultKeyName, "")

		return err == nil
	}, tc.timeout, tc.retry, "wait for keyVaultkey to be ready in azure")

	// delete key vault key
	EnsureDelete(ctx, t, tc, keyVaultKey)

	assert.Eventually(func() bool {
		kvault, _ := tc.keyVaultManager.GetVault(ctx, tc.resourceGroupName, keyVaultInstance.Name)
		vUrl := *kvault.Properties.VaultURI
		_, err := kvopsclient.GetKey(ctx, vUrl, keyVaultKeyName, "")

		return err != nil
	}, tc.timeout, tc.retry, "wait for keyVaultkey to be gone from azure")

	// delete key vault
	EnsureDelete(ctx, t, tc, keyVaultInstance)

}
