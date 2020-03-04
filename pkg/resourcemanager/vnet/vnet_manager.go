// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package vnet

import (
	"context"

	vnetwork "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	telemetry "github.com/Azure/azure-service-operator/pkg/telemetry"
	"github.com/Azure/go-autorest/autorest"
	"github.com/go-logr/logr"
	ctrl "sigs.k8s.io/controller-runtime"
)

// NewAzureVNetManager creates a new instance of AzureVNetManager
func NewAzureVNetManager(log logr.Logger) *AzureVNetManager {
	return &AzureVNetManager{
		Telemetry: *telemetry.InitializeTelemetryDefault(
			"VNet",
			ctrl.Log.WithName("controllers").WithName("VNet"),
		),
	}
}

// VNetManager manages VNet service components
type VNetManager interface {
	CreateVNet(ctx context.Context,
		location string,
		resourceGroupName string,
		resourceName string,
		addressSpace string,
		subnets []azurev1alpha1.VNetSubnets) (vnetwork.VirtualNetwork, error)

	DeleteVNet(ctx context.Context,
		resourceGroupName string,
		resourceName string) (autorest.Response, error)

	VNetExists(ctx context.Context,
		resourceGroupName string,
		resourceName string) (bool, error)

	// also embed async client methods
	resourcemanager.ARMClient
}
