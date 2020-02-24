// +build all keyvaultkey

package controllers

import (
	"context"
	"net/http"
	"testing"
	"time"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/stretchr/testify/assert"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestKeyvaultKeyControllerHappyPath(t *testing.T) {
	t.Parallel()
	defer PanicRecover()
	ctx := context.Background()
	assert := assert.New(t)
	var err error

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
	EnsureInstance(ctx, t, tc, keyVaultInstance)

	// Prep query for get
	keyVaultNamespacedName := types.NamespacedName{Name: keyVaultName, Namespace: "default"}

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
