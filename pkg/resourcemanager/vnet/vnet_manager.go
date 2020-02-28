/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
