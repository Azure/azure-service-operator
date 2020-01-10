// +build all keyvault

package controllers

import (
	"context"
	"net/http"
	"testing"
	"time"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestKeyvaultController(t *testing.T) {
	t.Parallel()
	RegisterTestingT(t)
	defer PanicRecover()
	ctx := context.Background()

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
			Location:          keyVaultLocation,
			ResourceGroupName: tc.resourceGroupName,
		},
	}

	// Create the Keyvault object and expect the Reconcile to be created
	err := tc.k8sClient.Create(ctx, keyVaultInstance)
	Expect(apierrors.IsInvalid(err)).To(Equal(false))
	Expect(err).NotTo(HaveOccurred())

	// Prep query for get
	keyVaultNamespacedName := types.NamespacedName{Name: keyVaultName, Namespace: "default"}

	// Wait until key vault is provisioned
	Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, keyVaultNamespacedName, keyVaultInstance)
		//log.Print(keyVaultInstance.Status)
		return keyVaultInstance.Status.Provisioned
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	// verify key vault exists in Azure
	Eventually(func() bool {
		result, _ := tc.keyVaultManager.GetVault(ctx, tc.resourceGroupName, keyVaultInstance.Name)
		return result.Response.StatusCode == http.StatusOK
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	// delete key vault
	err = tc.k8sClient.Delete(ctx, keyVaultInstance)
	Expect(err).NotTo(HaveOccurred())

	// verify key vault is gone from kubernetes
	Eventually(func() bool {
		err := tc.k8sClient.Get(ctx, keyVaultNamespacedName, keyVaultInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

	// confirm key vault is gone from Azure
	Eventually(func() bool {
		result, _ := tc.keyVaultManager.GetVault(ctx, tc.resourceGroupName, keyVaultInstance.Name)
		return result.Response.StatusCode == http.StatusNotFound
	}, tc.timeout, tc.retry,
	).Should(BeTrue())

}
