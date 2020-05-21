// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package vmext

import (
	"context"

	compute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

type VirtualMachineExtensionManager interface {
	CreateVirtualMachineExtension(ctx context.Context,
		location string,
		resourceGroupName string,
		vmName string,
		extName string,
		autoUpgradeMinorVersion bool,
		forceUpdateTag string,
		publisher string,
		typeName string,
		typeHandlerVersion string,
		settings string,
		protectedSettings string) (compute.VirtualMachineExtension, error)

	DeleteVirtualMachineExtension(ctx context.Context,
		extName string,
		vmName string,
		resourceGroupName string) (string, error)

	GetVirtualMachineExtension(ctx context.Context,
		resourceGroupName string,
		vmName string,
		extName string) (compute.VirtualMachineExtension, error)

	// also embed async client methods
	resourcemanager.ARMClient
}
