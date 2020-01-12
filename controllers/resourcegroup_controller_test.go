// +build all resourcegroup

package controllers

import (
	"context"
	"net/http"
	"strings"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/stretchr/testify/assert"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestResourceGroupControllerHappyPath(t *testing.T) {
	t.Parallel()
	defer PanicRecover()
	ctx := context.Background()
	assert := assert.New(t)

	resourceGroupName := "t-rg-dev-" + helpers.RandomString(10)

	var err error

	// Create the ResourceGroup object and expect the Reconcile to be created
	resourceGroupInstance := &azurev1alpha1.ResourceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      resourceGroupName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.ResourceGroupSpec{
			Location: tc.resourceGroupLocation,
		},
	}

	// create rg
	err = tc.k8sClient.Create(ctx, resourceGroupInstance)
	assert.Equal(false, apierrors.IsInvalid(err), "create db resource")
	assert.Equal(nil, err, "create rg in k8s")

	resourceGroupNamespacedName := types.NamespacedName{Name: resourceGroupName, Namespace: "default"}

	// make sure rg has a finalizer
	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, resourceGroupNamespacedName, resourceGroupInstance)
		return resourceGroupInstance.HasFinalizer(finalizerName)
	}, tc.timeout, tc.retry, "wait for finlizer on rg")

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, resourceGroupNamespacedName, resourceGroupInstance)
		return strings.Contains(resourceGroupInstance.Status.Message, "successfully provisioned")
	}, tc.timeout, tc.retry, "wait for finlizer on rg")

	// verify rg exists in azure

	assert.Eventually(func() bool {
		_, err := tc.resourceGroupManager.CheckExistence(ctx, resourceGroupName)
		return err == nil
	}, tc.timeout, tc.retry, "wait for resourceGroupInstance to be gone from k8s")

	// delete rg
	err = tc.k8sClient.Delete(ctx, resourceGroupInstance)
	assert.Equal(nil, err, "delete rg in k8s")

	// verify rg is being deleted

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, resourceGroupNamespacedName, resourceGroupInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for resourceGroupInstance to be gone from k8s")

	assert.Eventually(func() bool {
		result, _ := tc.resourceGroupManager.CheckExistence(ctx, resourceGroupName)
		if result.Response == nil {
			return false
		}
		return result.Response.StatusCode == http.StatusNotFound
	}, tc.timeout, tc.retry, "wait for resourceGroupInstance to be gone from k8s")

}
