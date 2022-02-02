// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || keyvault
// +build all keyvault

package controllers

import (
	"context"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/gofrs/uuid"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	kvsecrets "github.com/Azure/azure-service-operator/pkg/secrets/keyvault"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestKeyvaultControllerHappyPath(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	keyVaultName := GenerateTestResourceNameWithRandom("kv", 6)
	const poll = time.Second * 10

	keyVaultLocation := tc.resourceGroupLocation

	// Declare KeyVault object
	keyVaultInstance := &azurev1alpha1.KeyVault{
		ObjectMeta: metav1.ObjectMeta{
			Name:      keyVaultName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.KeyVaultSpec{
			Location:         keyVaultLocation,
			ResourceGroup:    tc.resourceGroupName,
			EnableSoftDelete: true,
		},
	}

	// Create the Keyvault object and expect the Reconcile to be created
	EnsureInstance(ctx, t, tc, keyVaultInstance)

	// verify key vault exists in Azure
	assert.Eventually(func() bool {
		result, _ := tc.keyVaultManager.GetVault(ctx, tc.resourceGroupName, keyVaultInstance.Name)
		return result.Response.StatusCode == http.StatusOK
	}, tc.timeout, tc.retry, "wait for keyVaultInstance to be ready in azure")

	EnsureDelete(ctx, t, tc, keyVaultInstance)

	assert.Eventually(func() bool {
		result, _ := tc.keyVaultManager.GetVault(ctx, tc.resourceGroupName, keyVaultInstance.Name)
		return result.Response.StatusCode == http.StatusNotFound
	}, tc.timeout, tc.retry, "wait for keyVaultInstance to be gone from azure")

}

func TestKeyvaultControllerWithAccessPolicies(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	keyVaultName := GenerateTestResourceNameWithRandom("kv", 6)
	const poll = time.Second * 10
	keyVaultLocation := tc.resourceGroupLocation
	accessPolicies := []azurev1alpha1.AccessPolicyEntry{
		{
			TenantID: config.GlobalCredentials().TenantID(),
			ClientID: config.GlobalCredentials().ClientID(),

			Permissions: &azurev1alpha1.Permissions{
				Keys: &[]string{
					"get",
					"list",
				},
				Secrets: &[]string{
					"get",
					"list",
					"set",
				},
				Certificates: &[]string{
					"get",
					"list",
				},
				Storage: &[]string{
					"get",
					"list",
				},
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

	EnsureInstance(ctx, t, tc, keyVaultInstance)

	// verify key vault exists in Azure
	assert.Eventually(func() bool {
		result, _ := tc.keyVaultManager.GetVault(ctx, tc.resourceGroupName, keyVaultInstance.Name)
		return result.Response.StatusCode == http.StatusOK
	}, tc.timeout, tc.retry, "wait for keyVaultInstance to be ready in azure")

	//Add code to set secret and get secret from this keyvault using secretclient

	keyvaultSecretClient := kvsecrets.New(
		keyVaultName,
		config.GlobalCredentials(),
		config.SecretNamingVersion(),
		config.PurgeDeletedKeyVaultSecrets(),
		config.RecoverSoftDeletedKeyVaultSecrets())
	secretName := "test-key"
	key := secrets.SecretKey{Name: secretName, Namespace: "default", Kind: "test"}
	datanew := map[string][]byte{
		"test1": []byte("test2"),
		"test2": []byte("test3"),
	}
	err := keyvaultSecretClient.Upsert(ctx, key, datanew)
	assert.Equal(nil, err, "expect secret to be inserted into keyvault")

	_, err = keyvaultSecretClient.Get(ctx, key)
	assert.Equal(nil, err, "checking if secret is present in keyvault")

	EnsureDelete(ctx, t, tc, keyVaultInstance)

	assert.Eventually(func() bool {
		result, _ := tc.keyVaultManager.GetVault(ctx, tc.resourceGroupName, keyVaultInstance.Name)
		return result.Response.StatusCode == http.StatusNotFound
	}, tc.timeout, tc.retry, "wait for keyVaultInstance to be gone from azure")
}

func TestKeyvaultControllerWithLimitedAccessPoliciesAndUpdate(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	keyVaultName := GenerateTestResourceNameWithRandom("kv", 6)
	const poll = time.Second * 10

	// Declare KeyVault object
	keyVaultInstance := &azurev1alpha1.KeyVault{
		ObjectMeta: metav1.ObjectMeta{
			Name:      keyVaultName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.KeyVaultSpec{
			Location:      tc.resourceGroupLocation,
			ResourceGroup: tc.resourceGroupName,
			AccessPolicies: &[]azurev1alpha1.AccessPolicyEntry{
				{
					TenantID: config.GlobalCredentials().TenantID(),
					ClientID: config.GlobalCredentials().ClientID(),
					Permissions: &azurev1alpha1.Permissions{
						Secrets: &[]string{"backup"},
					},
				},
			},
		},
	}

	EnsureInstance(ctx, t, tc, keyVaultInstance)

	// verify key vault exists in Azure
	assert.Eventually(func() bool {
		result, _ := tc.keyVaultManager.GetVault(ctx, tc.resourceGroupName, keyVaultInstance.Name)
		return result.Response.StatusCode == http.StatusOK
	}, tc.timeout, tc.retry, "wait for keyVaultInstance to be ready in azure")
	//Add code to set secret and get secret from this keyvault using secretclient

	keyvaultSecretClient := kvsecrets.New(
		keyVaultName,
		config.GlobalCredentials(),
		config.SecretNamingVersion(),
		config.PurgeDeletedKeyVaultSecrets(),
		config.RecoverSoftDeletedKeyVaultSecrets())
	key := secrets.SecretKey{Name: "test-key", Namespace: "default", Kind: "test"}
	datanew := map[string][]byte{
		"test1": []byte("test2"),
		"test2": []byte("test3"),
	}

	err := keyvaultSecretClient.Upsert(ctx, key, datanew)
	assert.NotEqual(nil, err, "should not be able to insert secrets")

	_, err = keyvaultSecretClient.Get(ctx, key)
	assert.NotEqual(nil, err, "should not be able to get secrets")

	names := types.NamespacedName{Name: keyVaultName, Namespace: "default"}

	err = tc.k8sClient.Get(ctx, names, keyVaultInstance)
	assert.Equal(nil, err, "get keyvault in k8s")

	originalHash := keyVaultInstance.Status.SpecHash
	updatedPermissions := []string{"get", "list", "set"}
	l := *keyVaultInstance.Spec.AccessPolicies
	l[0].Permissions.Secrets = &updatedPermissions

	err = tc.k8sClient.Update(ctx, keyVaultInstance)
	assert.Equal(nil, err, "updating keyvault in k8s")

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, names, keyVaultInstance)
		return originalHash != keyVaultInstance.Status.SpecHash
	}, tc.timeout, tc.retry, "wait for keyVaultInstance to be updated")

	assert.Eventually(func() bool {
		err := keyvaultSecretClient.Upsert(ctx, key, datanew)
		return err == nil
	}, tc.timeoutFast, tc.retry, "secret should get inserted")

	assert.Eventually(func() bool {
		_, err = keyvaultSecretClient.Get(ctx, key)
		return err == nil
	}, tc.timeoutFast, tc.retry, "should now be able to get secret")

	EnsureDelete(ctx, t, tc, keyVaultInstance)

	assert.Eventually(func() bool {
		result, _ := tc.keyVaultManager.GetVault(ctx, tc.resourceGroupName, keyVaultInstance.Name)
		return result.Response.StatusCode == http.StatusNotFound
	}, tc.timeout, tc.retry, "wait for keyVaultInstance to be gone from azure")

}

func TestKeyvaultControllerInvalidName(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	keyVaultName := "k"

	keyVaultLocation := tc.resourceGroupLocation

	// Declare KeyVault object
	keyVaultInstance := &azurev1alpha1.KeyVault{
		ObjectMeta: metav1.ObjectMeta{
			Name:      keyVaultName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.KeyVaultSpec{
			Location:      keyVaultLocation,
			ResourceGroup: tc.resourceGroupName,
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, keyVaultInstance, "validation failed: parameter=vaultName", false)

	EnsureDelete(ctx, t, tc, keyVaultInstance)
}

func TestKeyvaultControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	keyVaultName := helpers.FillWithRandom(GenerateTestResourceName("kv"), 24)

	keyVaultLocation := tc.resourceGroupLocation

	// Declare KeyVault object
	keyVaultInstance := &azurev1alpha1.KeyVault{
		ObjectMeta: metav1.ObjectMeta{
			Name:      keyVaultName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.KeyVaultSpec{
			Location:      keyVaultLocation,
			ResourceGroup: "fakerg",
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, keyVaultInstance, errhelp.ResourceGroupNotFoundErrorCode, false)

	EnsureDelete(ctx, t, tc, keyVaultInstance)

}

func TestKeyvaultControllerWithVirtualNetworkRulesAndUpdate(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	keyVaultName := GenerateTestResourceNameWithRandom("kv", 6)
	const poll = time.Second * 10
	keyVaultLocation := tc.resourceGroupLocation
	accessPolicies := []azurev1alpha1.AccessPolicyEntry{
		{
			TenantID: config.GlobalCredentials().TenantID(),
			ClientID: config.GlobalCredentials().ClientID(),

			Permissions: &azurev1alpha1.Permissions{
				Secrets: &[]string{
					"get",
					"list",
					"set",
				},
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

	EnsureInstance(ctx, t, tc, keyVaultInstance)

	// verify key vault exists in Azure
	assert.Eventually(func() bool {
		result, _ := tc.keyVaultManager.GetVault(ctx, tc.resourceGroupName, keyVaultInstance.Name)
		return result.Response.StatusCode == http.StatusOK
	}, tc.timeout, tc.retry, "wait for keyVaultInstance to be ready in azure")

	keyvaultSecretClient := kvsecrets.New(
		keyVaultName,
		config.GlobalCredentials(),
		config.SecretNamingVersion(),
		config.PurgeDeletedKeyVaultSecrets(),
		config.RecoverSoftDeletedKeyVaultSecrets())
	secretName := "test-key"
	key := secrets.SecretKey{Name: secretName, Namespace: "default", Kind: "test"}
	datanew := map[string][]byte{
		"test1": []byte("test2"),
		"test2": []byte("test3"),
	}
	assert.Eventually(func() bool {
		err := keyvaultSecretClient.Upsert(ctx, key, datanew)
		return err == nil
	}, tc.timeoutFast, tc.retry, "wait for secret to be inserted into keyvault")

	_, err := keyvaultSecretClient.Get(ctx, key)
	assert.Equal(nil, err, "checking if secret is present in keyvault")

	names := types.NamespacedName{Name: keyVaultName, Namespace: "default"}
	err = tc.k8sClient.Get(ctx, names, keyVaultInstance)
	assert.Equal(nil, err, "get keyvault in k8s")

	originalHash := keyVaultInstance.Status.SpecHash
	networkPolicy := azurev1alpha1.NetworkRuleSet{
		Bypass:        "None",
		DefaultAction: "Deny",
	}
	keyVaultInstance.Spec.NetworkPolicies = &networkPolicy

	err = tc.k8sClient.Update(ctx, keyVaultInstance)
	assert.Equal(nil, err, "updating keyvault in k8s")

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, names, keyVaultInstance)
		return originalHash != keyVaultInstance.Status.SpecHash
	}, tc.timeout, tc.retry, "wait for keyVaultInstance to be updated")

	assert.Eventually(func() bool {
		_, err = keyvaultSecretClient.Get(ctx, key)
		return err != nil
	}, tc.timeoutFast, tc.retry, "wait for secret to be inaccessible")

	EnsureDelete(ctx, t, tc, keyVaultInstance)

	assert.Eventually(func() bool {
		result, _ := tc.keyVaultManager.GetVault(ctx, tc.resourceGroupName, keyVaultInstance.Name)
		return result.Response.StatusCode == http.StatusNotFound
	}, tc.timeout, tc.retry, "wait for keyVaultInstance to be gone from azure")
}

func TestKeyvaultControllerBadAccessPolicy(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	keyVaultName := GenerateTestResourceNameWithRandom("kv", 6)
	keyVaultLocation := tc.resourceGroupLocation

	// Declare KeyVault object
	accessPolicies := []azurev1alpha1.AccessPolicyEntry{
		{
			TenantID: config.GlobalCredentials().TenantID(),
			ClientID: uuid.Nil.String(),

			Permissions: &azurev1alpha1.Permissions{
				Keys: &[]string{
					"get",
					"list",
				},
			},
		}}

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

	assert := assert.New(t)

	err := tc.k8sClient.Create(ctx, keyVaultInstance)
	assert.NoError(err)
	namespacedName := types.NamespacedName{Name: keyVaultInstance.GetName(), Namespace: keyVaultInstance.GetNamespace()}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, namespacedName, keyVaultInstance)
		return strings.Contains(keyVaultInstance.Status.Message, "Authorization_RequestDenied")
	}, tc.timeout, tc.retry, "wait for message to be updated indicating auth error")

	EnsureDelete(ctx, t, tc, keyVaultInstance)
}
