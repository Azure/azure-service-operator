// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all appinsights

package controllers

import (
	"context"
	"strings"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestAppInsightsController(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	rgName := tc.resourceGroupName
	rgLocation := tc.resourceGroupLocation
	appInsightsName := GenerateTestResourceName("appinsights")

	// Create an instance of Azure AppInsights
	appInsightsInstance := &azurev1alpha1.AppInsights{
		ObjectMeta: metav1.ObjectMeta{
			Name:      appInsightsName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AppInsightsSpec{
			Kind:            "web",
			Location:        rgLocation,
			ResourceGroup:   rgName,
			ApplicationType: "other",
		},
	}

	err := tc.k8sClient.Create(ctx, appInsightsInstance)
	assert.Equal(nil, err, "create appinsights record in k8s")

	appInsightsNamespacedName := types.NamespacedName{Name: appInsightsName, Namespace: "default"}

	// Wait for the AppInsights instance to be provisioned
	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, appInsightsNamespacedName, appInsightsInstance)
		return strings.Contains(appInsightsInstance.Status.Message, successMsg)
	}, tc.timeout, tc.retry, "awaiting appinsights instance creation")

	// Delete the service
	err = tc.k8sClient.Delete(ctx, appInsightsInstance)
	assert.Equal(nil, err, "deleting appinsights in k8s")

	// Wait for the AppInsights instance to be deleted
	assert.Eventually(func() bool {
		err := tc.k8sClient.Get(ctx, appInsightsNamespacedName, appInsightsInstance)
		return err != nil
	}, tc.timeout, tc.retry, "awaiting appInsightsInstance deletion")
}
