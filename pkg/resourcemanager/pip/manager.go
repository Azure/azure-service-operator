// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package pip

import (
	"context"

	vnetwork "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

type PublicIPAddressManager interface {
	CreatePublicIPAddress(ctx context.Context,
		location string,
		resourceGroupName string,
		resourceName string,
		publicIPAllocationMethod string,
		idleTimeoutInMinutes int,
		publicIPAddressVersion string,
		skuName string) (vnetwork.PublicIPAddress, error)

	DeletePublicIPAddress(ctx context.Context,
		resourceName string,
		resourceGroupName string) (string, error)

	GetPublicIPAddress(ctx context.Context,
		resourceGroupName string,
		resourceName string) (vnetwork.PublicIPAddress, error)

	// also embed async client methods
	resourcemanager.ARMClient
}
