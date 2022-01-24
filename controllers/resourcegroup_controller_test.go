// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || resourcegroup
// +build all resourcegroup

package controllers

import (
	"context"
	"net/http"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestResourceGroupControllerHappyPath(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	resourceGroupName := GenerateTestResourceNameWithRandom("rg-dev", 10)

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
	EnsureInstance(ctx, t, tc, resourceGroupInstance)

	// verify rg exists in azure
	assert.Eventually(func() bool {
		_, err := tc.resourceGroupManager.CheckExistence(ctx, resourceGroupName)
		return err == nil
	}, tc.timeout, tc.retry, "wait for resourceGroupInstance to exist in azure")

	// delete rg
	EnsureDelete(ctx, t, tc, resourceGroupInstance)

	assert.Eventually(func() bool {
		result, _ := tc.resourceGroupManager.CheckExistence(ctx, resourceGroupName)
		if result.Response == nil {
			return false
		}
		return result.Response.StatusCode == http.StatusNotFound
	}, tc.timeout, tc.retry, "wait for resourceGroupInstance to be gone from azure")

}
