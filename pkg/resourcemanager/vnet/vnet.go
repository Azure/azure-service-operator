// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package vnet

import (
	"context"

	vnetwork "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	telemetry "github.com/Azure/azure-service-operator/pkg/telemetry"
	"github.com/Azure/go-autorest/autorest"
)

// AzureVNetManager is the struct that the manager functions hang off
type AzureVNetManager struct {
	Telemetry telemetry.Telemetry
}

// getVNetClient returns a new instance of an VirtualNetwork client
func getVNetClient() (vnetwork.VirtualNetworksClient, error) {
	client := vnetwork.NewVirtualNetworksClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		client = vnetwork.VirtualNetworksClient{}
	} else {
		client.Authorizer = a
		client.AddToUserAgent(config.UserAgent())
	}
	return client, err
}

// CreateVNet creates VNets
func (_ *AzureVNetManager) CreateVNet(ctx context.Context, location string, resourceGroupName string, resourceName string, addressSpace string, subnets []azurev1alpha1.VNetSubnets) (vnetwork.VirtualNetwork, error) {
	client, err := getVNetClient()
	if err != nil {
		return vnetwork.VirtualNetwork{}, err
	}

	var subnetsToAdd []vnetwork.Subnet
	for i := 0; i < len(subnets); i++ {
		subnetsToAdd = append(
			subnetsToAdd,
			vnetwork.Subnet{
				Name: &subnets[i].SubnetName,
				SubnetPropertiesFormat: &vnetwork.SubnetPropertiesFormat{
					AddressPrefix: &subnets[i].SubnetAddressPrefix,
				},
			},
		)
	}

	future, err := client.CreateOrUpdate(
		ctx,
		resourceGroupName,
		resourceName,
		vnetwork.VirtualNetwork{
			Location: &location,
			VirtualNetworkPropertiesFormat: &vnetwork.VirtualNetworkPropertiesFormat{
				AddressSpace: &vnetwork.AddressSpace{
					AddressPrefixes: &[]string{addressSpace},
				},
				Subnets: &subnetsToAdd,
			},
		},
	)
	if err != nil {
		return vnetwork.VirtualNetwork{}, err
	}

	return future.Result(client)
}

// DeleteVNet deletes a VNet
func (_ *AzureVNetManager) DeleteVNet(ctx context.Context, resourceGroupName string, resourceName string) (autorest.Response, error) {
	client, err := getVNetClient()
	if err != nil {
		return autorest.Response{}, err
	}

	future, err := client.Delete(
		ctx,
		resourceGroupName,
		resourceName,
	)
	if err != nil {
		return autorest.Response{}, err
	}

	return future.Result(client)
}

// VNetExists checks to see if a VNet exists
func (v *AzureVNetManager) VNetExists(ctx context.Context, resourceGroupName string, resourceName string) (vNet vnetwork.VirtualNetwork, err error) {
	client, err := getVNetClient()
	if err != nil {
		return vnetwork.VirtualNetwork{}, err
	}

	return client.Get(ctx, resourceGroupName, resourceName, "")
}
