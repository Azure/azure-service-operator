// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package disk

import (
	"context"

	compute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

type DiskManager interface {
	CreateDisk(ctx context.Context,
		location string,
		resourceGroupName string,
		createOption string,
		diskSizeGB int) (compute.Disk, error)

	DeleteDisk(ctx context.Context,
		resourceName string,
		resourceGroupName string) (string, error)

	GetDisk(ctx context.Context,
		resourceGroupName string,
		resourceName string) (compute.Disk, error)

	// also embed async client methods
	resourcemanager.ARMClient
}
