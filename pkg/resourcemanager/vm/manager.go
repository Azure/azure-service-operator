// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package vm

import (
	"context"

	compute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

type VirtualMachineManager interface {
	CreateVirtualMachine(ctx context.Context,
		location string,
		resourceGroupName string,
		resourceName string,
		vmSize string,
		osType string,
		adminUserName string,
		adminPassword string,
		sshPublicKeyData string,
		networkInterfaceName string,
		platformImageURN string) (compute.VirtualMachine, error)

	DeleteVirtualMachine(ctx context.Context,
		resourceName string,
		resourceGroupName string) (string, error)

	GetVirtualMachine(ctx context.Context,
		resourceGroupName string,
		resourceName string) (compute.VirtualMachine, error)

	// also embed async client methods
	resourcemanager.ARMClient
}
