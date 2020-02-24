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
	defer PanicRecover()
	ctx := context.Background()
	assert := assert.New(t)

	keyVaultName := "t-kv-dev-" + helpers.RandomString(10)
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
	err := tc.k8sClient.Create(ctx, keyVaultInstance)
	assert.Equal(nil, err, "create keyvault in k8s")

	// Prep query for get
	keyVaultNamespacedName := types.NamespacedName{Name: keyVaultName, Namespace: "default"}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, keyVaultNamespacedName, keyVaultInstance)
		return helpers.HasFinalizer(keyVaultInstance, finalizerName)
	}, tc.timeout, tc.retry, "wait for keyvault to have finalizer")

	// Wait until key vault is provisioned

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, keyVaultNamespacedName, keyVaultInstance)
		return strings.Contains(keyVaultInstance.Status.Message, successMsg)
	}, tc.timeout, tc.retry, "wait for keyVaultInstance to be ready in k8s")

	// verify key vault exists in Azure
	assert.Eventually(func() bool {
		result, _ := tc.keyVaultManager.GetVault(ctx, tc.resourceGroupName, keyVaultInstance.Name)
		return result.Response.StatusCode == http.StatusOK
	}, tc.timeout, tc.retry, "wait for keyVaultInstance to be ready in azure")

	// delete key vault
	err = tc.k8sClient.Delete(ctx, keyVaultInstance)
	assert.Equal(nil, err, "delete keyvault in k8s")

	// verify key vault is gone from kubernetes

	assert.Eventually(func() bool {
		err := tc.k8sClient.Get(ctx, keyVaultNamespacedName, keyVaultInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for keyVaultInstance to be gone from k8s")

	assert.Eventually(func() bool {
		result, _ := tc.keyVaultManager.GetVault(ctx, tc.resourceGroupName, keyVaultInstance.Name)
		return result.Response.StatusCode == http.StatusNotFound
	}, tc.timeout, tc.retry, "wait for keyVaultInstance to be gone from azure")

}

func TestKeyvaultControllerWithAccessPolicies(t *testing.T) {
	t.Parallel()
	defer PanicRecover()
	ctx := context.Background()
	assert := assert.New(t)
	
	keyVaultName := "t-kv-dev-" + helpers.RandomString(10)
	const poll = time.Second * 10
	keyVaultLocation := tc.resourceGroupLocation
	allPermissions := []string{"get", "list", "set", "delete", "recover", "backup", "restore"}
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
	err := tc.k8sClient.Create(ctx, keyVaultInstance)
	assert.Equal(nil, err, "create keyvault in k8s")

	// Prep query for get
	keyVaultNamespacedName := types.NamespacedName{Name: keyVaultName, Namespace: "default"}
  
	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, keyVaultNamespacedName, keyVaultInstance)
		return helpers.HasFinalizer(keyVaultInstance, finalizerName)
	}, tc.timeout, tc.retry, "wait for keyvault to have finalizer")


	// Wait until key vault is provisioned

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, keyVaultNamespacedName, keyVaultInstance)
		return strings.Contains(keyVaultInstance.Status.Message, successMsg)
	}, tc.timeout, tc.retry, "wait for keyVaultInstance to be ready in k8s")

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
	err = keyvaultSecretClient.Upsert(ctx, key, datanew)
	assert.Equal(nil, err, "expect secret to be inserted into keyvault")

	_, err = keyvaultSecretClient.Get(ctx, key)
	assert.Equal(nil, err, "checking if secret is present in keyvault")

	// delete key vault
	err = tc.k8sClient.Delete(ctx, keyVaultInstance)
	assert.Equal(nil, err, "delete keyvault in k8s")

	// verify key vault is gone from kubernetes
	assert.Eventually(func() bool {
		err := tc.k8sClient.Get(ctx, keyVaultNamespacedName, keyVaultInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for keyVaultInstance to be gone from k8s")

	assert.Eventually(func() bool {
		result, _ := tc.keyVaultManager.GetVault(ctx, tc.resourceGroupName, keyVaultInstance.Name)
		return result.Response.StatusCode == http.StatusNotFound
	}, tc.timeout, tc.retry, "wait for keyVaultInstance to be gone from azure")

}


func TestKeyvaultControllerInvalidName(t *testing.T) {
	t.Parallel()
	defer PanicRecover()
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

	// Create the Keyvault object and expect the Reconcile to be created
	err := tc.k8sClient.Create(ctx, keyVaultInstance)
	assert.Equal(nil, err, "create keyvault in k8s")

	// Prep query for get
	keyVaultNamespacedName := types.NamespacedName{Name: keyVaultName, Namespace: "default"}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, keyVaultNamespacedName, keyVaultInstance)
		return helpers.HasFinalizer(keyVaultInstance, finalizerName)
	}, tc.timeout, tc.retry, "wait for keyvault to have finalizer")

	// Verify you get the invalid name error

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, keyVaultNamespacedName, keyVaultInstance)
		return strings.Contains(keyVaultInstance.Status.Message, errhelp.AccountNameInvalid)
	}, tc.timeout, tc.retry, "wait for invalid account name error")

	// delete key vault
	err = tc.k8sClient.Delete(ctx, keyVaultInstance)
	assert.Equal(nil, err, "delete keyvault in k8s")

	// verify key vault is gone from kubernetes

	assert.Eventually(func() bool {
		err := tc.k8sClient.Get(ctx, keyVaultNamespacedName, keyVaultInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for keyVaultInstance to be gone from k8s")

}

func TestKeyvaultControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover()
	ctx := context.Background()
	assert := assert.New(t)

	keyVaultName := "t-kv-dev-" + helpers.RandomString(10)

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

	// Create the Keyvault object and expect the Reconcile to be created
	err := tc.k8sClient.Create(ctx, keyVaultInstance)
	assert.Equal(nil, err, "create keyvault in k8s")

	// Prep query for get
	keyVaultNamespacedName := types.NamespacedName{Name: keyVaultName, Namespace: "default"}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, keyVaultNamespacedName, keyVaultInstance)
		return helpers.HasFinalizer(keyVaultInstance, finalizerName)
	}, tc.timeout, tc.retry, "wait for keyvault to have finalizer")

	// Verify you get the resource group not found error

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, keyVaultNamespacedName, keyVaultInstance)
		return strings.Contains(keyVaultInstance.Status.Message, errhelp.ResourceGroupNotFoundErrorCode)
	}, tc.timeout, tc.retry, "wait for ResourceGroupNotFound error")

	// delete key vault
	err = tc.k8sClient.Delete(ctx, keyVaultInstance)
	assert.Equal(nil, err, "delete keyvault in k8s")

	// verify key vault is gone from kubernetes

	assert.Eventually(func() bool {
		err := tc.k8sClient.Get(ctx, keyVaultNamespacedName, keyVaultInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for keyVaultInstance to be gone from k8s")

}
