// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package appinsights

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/appinsights/mgmt/2015-05-01/insights"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/go-autorest/autorest"
)

// ApplicationInsightsManager manages Azure Application Insights service components
type ApplicationInsightsManager interface {
	CreateAppInsights(
		ctx context.Context,
		resourceGroupName string,
		kind string,
		applicationType string,
		location string,
		resourceName string) (insights.ApplicationInsightsComponent, error)
	DeleteAppInsights(ctx context.Context, resourceGroupName string, resourceName string) (autorest.Response, error)
	GetAppInsights(ctx context.Context, resourceGroupName string, resourceName string) (insights.ApplicationInsightsComponent, error)

	// ARM Client
	resourcemanager.ARMClient
}
