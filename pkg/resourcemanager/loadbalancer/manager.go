// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package loadbalancer

import (
	"context"

	network "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

type LoadBalancerManager interface {
	CreateLoadBalancer(ctx context.Context,
		location string,
		resourceGroupName string,
		resourceName string,
		publicIPAddressName string,
		backEndPoolName string,
		inboundNatPoolName string,
		frontEndPortRangeStart int,
		frontEndPortRangeEnd int,
		backEndPort int) (network.LoadBalancer, error)

	DeleteLoadBalancer(ctx context.Context,
		resourceName string,
		resourceGroupName string) (string, error)

	GetLoadBalancer(ctx context.Context,
		resourceGroupName string,
		resourceName string) (network.LoadBalancer, error)

	// also embed async client methods
	resourcemanager.ARMClient
}
