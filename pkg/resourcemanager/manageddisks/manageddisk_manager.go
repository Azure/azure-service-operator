// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package manageddisks

import (
	"context"

	compute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	telemetry "github.com/Azure/azure-service-operator/pkg/telemetry"
	"github.com/Azure/go-autorest/autorest"
	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
)

// NewAzureManagedDiskManager creates a new instance of AzureManagedDiskManager
func NewAzureManagedDiskManager(log logr.Logger) *AzureManagedDiskManager {
	return &AzureManagedDiskManager{
		Telemetry: *telemetry.InitializeTelemetryDefault(
			"ManagedDisk",
			ctrl.Log.WithName("controllers").WithName("ManagedDisk"),
		),
	}
}

// ManagedDiskManager manages ManagedDisk service components
type ManagedDiskManager interface {
	CreateManagedDisk(ctx context.Context,
		location string,
		resourceGroupName string,
		resourceName string,
		createOption string,
		diskSizeGB int) (compute.DisksClient, error)

	DeleteManagedDisk(ctx context.Context,
		resourceGroupName string,
		resourceName string) (autorest.Response, error)

	ManagedDiskExists(ctx context.Context,
		resourceGroupName string,
		resourceName string) (bool, error)

	// also embed async client methods
	resourcemanager.ARMClient
}
