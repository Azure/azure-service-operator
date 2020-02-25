// +build all keyvaultkey

package controllers

import (
	"context"
	"net/http"
	"testing"
	"time"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	resourcemanagerkeyvaults "github.com/Azure/azure-service-operator/pkg/resourcemanager/keyvaults"
	"github.com/stretchr/testify/assert"

	//apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestKeyvaultKeyControllerHappyPath(t *testing.T) {
	t.Parallel()
	defer PanicRecover()
	ctx := context.Background()
	assert := assert.New(t)

	keyVaultName := "t-kv-dev-" + helpers.RandomString(10)
	keyVaultKeyName := "t-kv-dev-" + helpers.RandomString(10)
	const poll = time.Second * 10

	keyVaultLocation := tc.resourceGroupLocation

	allPermissions := []string{"get", "list", "set", "delete", "recover", "backup", "restore", "create"}
	accessPolicies := []azurev1alpha1.AccessPolicyEntry{
		{
			TenantID: config.TenantID(),
			ObjectID: config.ClientID(),
			Permissions: &azurev1alpha1.Permissions{
				Keys:         &allPermissions,
				Secrets:      &allPermissions,
				Certificates: &allPermissions,
				Storage:      &allPermissions,
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
		},
	}

	// create key
	EnsureInstance(ctx, t, tc, keyVaultKey)

	kvopsclient := resourcemanagerkeyvaults.NewOpsClient(keyVaultName)

	assert.Eventually(func() bool {
		kvault, _ := tc.keyVaultManager.GetVault(ctx, tc.resourceGroupName, keyVaultInstance.Name)
		vUrl := *kvault.Properties.VaultURI
		_, err := kvopsclient.GetKey(ctx, vUrl, keyVaultKeyName, "")

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
