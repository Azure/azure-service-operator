// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || appinsights
// +build all appinsights

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestAppInsightsApiKeyController(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	rgName := tc.resourceGroupName
	rgLocation := tc.resourceGroupLocation
	appInsightsName := GenerateTestResourceName("appinsights2")
	appInsightsKeyName := GenerateTestResourceName("insightskey")

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

	apiKey := &azurev1alpha1.AppInsightsApiKey{
		ObjectMeta: metav1.ObjectMeta{
			Name:      appInsightsKeyName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AppInsightsApiKeySpec{
			AppInsights:      appInsightsName,
			ResourceGroup:    rgName,
			ReadTelemetry:    true,
			WriteAnnotations: true,
		},
	}

	EnsureInstance(ctx, t, tc, apiKey)
	EnsureDelete(ctx, t, tc, apiKey)
	EnsureDelete(ctx, t, tc, appInsightsInstance)
}
