// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package apimservice

import (
	"context"

	apim "github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2019-01-01/apimanagement"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	telemetry "github.com/Azure/azure-service-operator/pkg/telemetry"
	ctrl "sigs.k8s.io/controller-runtime"
)

// NewAzureAPIMgmtServiceManager creates a new instance of AzureAPIMgmtServiceManager
func NewAzureAPIMgmtServiceManager() *AzureAPIMgmtServiceManager {
	return &AzureAPIMgmtServiceManager{
		Telemetry: telemetry.InitializePrometheusDefault(
			ctrl.Log.WithName("controllers").WithName("ApimService"),
			"ApimService",
		),
	}
}

// APIMgmtServiceManager manages Azure Application Insights service components
type APIMgmtServiceManager interface {
	CreateAPIMgmtSvc(ctx context.Context,
		tier string,
		location string,
		resourceGroupName string,
		resourceName string,
		publisherName string,
		publisherEmail string) (*apim.ServiceResource, error)

	DeleteAPIMgmtSvc(ctx context.Context,
		resourceGroupName string,
		resourceName string) (*apim.ServiceResource, error)

	APIMgmtSvcStatus(ctx context.Context,
		resourceGroupName string,
		resourceName string) (exists bool, result bool, resourceID *string, err error)

	SetVNetForAPIMgmtSvc(ctx context.Context,
		resourceGroupName string,
		resourceName string,
		vnetType string,
		vnetResourceGroupName string,
		vnetResourceName string,
		subnetName string) (err error, updated bool)

	SetAppInsightsForAPIMgmtSvc(ctx context.Context,
		resourceGroupName string,
		resourceName string,
		appInsightsResourceGroup string,
		appInsightsName string) error

	CheckAPIMgmtSvcName(ctx context.Context, resourceName string) (available bool, err error)

	// also embed async client methods
	resourcemanager.ARMClient
}
