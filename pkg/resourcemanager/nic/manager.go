// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package nic

import (
	"context"

	network "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

type NetworkInterfaceManager interface {
	CreateNetworkInterface(ctx context.Context,
		location string,
		resourceGroupName string,
		resourceName string,
		vnetName string,
		subnetName string,
		publicIPAddressName string) (network.Interface, error)

	DeleteNetworkInterface(ctx context.Context,
		resourceName string,
		resourceGroupName string) (string, error)

	GetNetworkInterface(ctx context.Context,
		resourceGroupName string,
		resourceName string) (network.Interface, error)

	// also embed async client methods
	resourcemanager.ARMClient
}
