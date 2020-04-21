// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all keyvault

package controllers

import (
	"context"
	"net/http"
	"strings"
	"testing"
	"time"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	kvsecrets "github.com/Azure/azure-service-operator/pkg/secrets/keyvault"

	"github.com/stretchr/testify/assert"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
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
			TenantID: config.TenantID(),
			ClientID: config.ClientID(),

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

	keyvaultSecretClient := kvsecrets.New(keyVaultName)
	secretName := "test-key"
	key := types.NamespacedName{Name: secretName, Namespace: "default"}
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
	keyVaultLocation := tc.resourceGroupLocation
	limitedPermissions := []string{"backup"}

	accessPolicies := azurev1alpha1.AccessPolicyEntry{
		TenantID: config.TenantID(),
		ClientID: config.ClientID(),
		Permissions: &azurev1alpha1.Permissions{
			Secrets: &limitedPermissions,
		},
	}

	// Declare KeyVault object
	keyVaultInstance := &azurev1alpha1.KeyVault{
		ObjectMeta: metav1.ObjectMeta{
			Name:      keyVaultName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.KeyVaultSpec{
			Location:       keyVaultLocation,
			ResourceGroup:  tc.resourceGroupName,
			AccessPolicies: &[]azurev1alpha1.AccessPolicyEntry{accessPolicies},
		},
	}

	EnsureInstance(ctx, t, tc, keyVaultInstance)

	// verify key vault exists in Azure
	assert.Eventually(func() bool {
		result, _ := tc.keyVaultManager.GetVault(ctx, tc.resourceGroupName, keyVaultInstance.Name)
		return result.Response.StatusCode == http.StatusOK
	}, tc.timeout, tc.retry, "wait for keyVaultInstance to be ready in azure")
	//Add code to set secret and get secret from this keyvault using secretclient

	keyvaultSecretClient := kvsecrets.New(keyVaultName)
	secretName := "test-key"
	key := types.NamespacedName{Name: secretName, Namespace: "default"}
	datanew := map[string][]byte{
		"test1": []byte("test2"),
		"test2": []byte("test3"),
	}
	err := keyvaultSecretClient.Upsert(ctx, key, datanew)
	assert.NotEqual(nil, err, "expect secret to not be inserted into keyvault")

	_, err = keyvaultSecretClient.Get(ctx, key)
	assert.NotEqual(nil, err, "should not be able to get secrets")

	updatedPermissions := []string{"get", "list", "set"}
	accessPolicies.Permissions.Secrets = &updatedPermissions

	names := types.NamespacedName{Name: keyVaultName, Namespace: "default"}

	retInstance := &azurev1alpha1.KeyVault{}
	err = tc.k8sClient.Get(ctx, names, retInstance)
	assert.Equal(nil, err, "get keyvault in k8s")
	originalHash := retInstance.Status.SpecHash
	retInstance.Spec.AccessPolicies = &[]azurev1alpha1.AccessPolicyEntry{accessPolicies}

	err = tc.k8sClient.Update(ctx, retInstance)
	assert.Equal(nil, err, "updating keyvault in k8s")

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, names, retInstance)
		return originalHash != retInstance.Status.SpecHash
	}, tc.timeout, tc.retry, "wait for keyVaultInstance to be updated")

	err = keyvaultSecretClient.Upsert(ctx, key, datanew)
	assert.Equal(nil, err, "expect secret to be inserted into keyvault after update")

	_, err = keyvaultSecretClient.Get(ctx, key)
	assert.Equal(nil, err, "should be able to get secrets after update")

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
	assert := assert.New(t)

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

	EnsureInstanceWithResult(ctx, t, tc, keyVaultInstance, errhelp.AccountNameInvalid, false)

	EnsureDelete(ctx, t, tc, keyVaultInstance)
}

func TestKeyvaultControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

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
			TenantID: config.TenantID(),
			ClientID: config.ClientID(),

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

	keyvaultSecretClient := kvsecrets.New(keyVaultName)
	secretName := "test-key"
	key := types.NamespacedName{Name: secretName, Namespace: "default"}
	datanew := map[string][]byte{
		"test1": []byte("test2"),
		"test2": []byte("test3"),
	}
	err := keyvaultSecretClient.Upsert(ctx, key, datanew)
	assert.Equal(nil, err, "expect secret to be inserted into keyvault")

	_, err = keyvaultSecretClient.Get(ctx, key)
	assert.Equal(nil, err, "checking if secret is present in keyvault")

	names := types.NamespacedName{Name: keyVaultName, Namespace: "default"}
	retInstance := &azurev1alpha1.KeyVault{}
	err = tc.k8sClient.Get(ctx, names, retInstance)

	assert.Equal(nil, err, "get keyvault in k8s")
	originalHash := retInstance.Status.SpecHash
	networkPolicy := azurev1alpha1.NetworkRuleSet{
		Bypass:        "None",
		DefaultAction: "Deny",
	}
	retInstance.Spec.NetworkPolicies = &networkPolicy

	err = tc.k8sClient.Update(ctx, retInstance)
	assert.Equal(nil, err, "updating keyvault in k8s")

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, names, retInstance)
		return originalHash != retInstance.Status.SpecHash
	}, tc.timeout, tc.retry, "wait for keyVaultInstance to be updated")

	_, err = keyvaultSecretClient.Get(ctx, key)
	assert.NotEqual(nil, err, "should not be able to get secrets after updating network rules")

	EnsureDelete(ctx, t, tc, keyVaultInstance)

	assert.Eventually(func() bool {
		result, _ := tc.keyVaultManager.GetVault(ctx, tc.resourceGroupName, keyVaultInstance.Name)
		return result.Response.StatusCode == http.StatusNotFound
	}, tc.timeout, tc.retry, "wait for keyVaultInstance to be gone from azure")
}
