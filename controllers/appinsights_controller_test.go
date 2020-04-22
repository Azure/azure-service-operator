// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all appinsights

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestAppInsightsController(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

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

	EnsureInstance(ctx, t, tc, appInsightsInstance)

	EnsureDelete(ctx, t, tc, appInsightsInstance)
}
