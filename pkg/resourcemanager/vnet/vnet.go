// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package vnet

import (
	"context"
	"fmt"
	"net"

	vnetwork "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/telemetry"
)

// AzureVNetManager is the struct that the manager functions hang off
type AzureVNetManager struct {
	Creds     config.Credentials
	Telemetry telemetry.Telemetry
}

// getVNetClient returns a new instance of an VirtualNetwork client
func getVNetClient(creds config.Credentials) (vnetwork.VirtualNetworksClient, error) {
	client := vnetwork.NewVirtualNetworksClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer(creds)
	if err != nil {
		client = vnetwork.VirtualNetworksClient{}
	} else {
		client.Authorizer = a
		client.AddToUserAgent(config.UserAgent())
	}
	return client, err
}

// CreateVNet creates VNets
func (m *AzureVNetManager) CreateVNet(ctx context.Context, location string, resourceGroupName string, resourceName string, addressSpace string, subnets []azurev1alpha1.VNetSubnets) (vnetwork.VirtualNetwork, error) {
	client, err := getVNetClient(m.Creds)
	if err != nil {
		return vnetwork.VirtualNetwork{}, err
	}

	var subnetsToAdd []vnetwork.Subnet
	for i := 0; i < len(subnets); i++ {
		eps := []vnetwork.ServiceEndpointPropertiesFormat{}
		for _, ep := range subnets[i].ServiceEndpoints {
			eps = append(eps, vnetwork.ServiceEndpointPropertiesFormat{
				Service: to.StringPtr(ep),
			})
		}

		s := vnetwork.Subnet{
			Name: &subnets[i].SubnetName,
			SubnetPropertiesFormat: &vnetwork.SubnetPropertiesFormat{
				AddressPrefix:    &subnets[i].SubnetAddressPrefix,
				ServiceEndpoints: &eps,
			},
		}

		subnetsToAdd = append(subnetsToAdd, s)
	}

	var future vnetwork.VirtualNetworksCreateOrUpdateFuture
	future, err = client.CreateOrUpdate(
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
func (m *AzureVNetManager) DeleteVNet(ctx context.Context, resourceGroupName string, resourceName string) (autorest.Response, error) {
	client, err := getVNetClient(m.Creds)
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

// GetVNet gets a VNet
func (m *AzureVNetManager) GetVNet(ctx context.Context, resourceGroupName string, resourceName string) (vNet vnetwork.VirtualNetwork, err error) {
	client, err := getVNetClient(m.Creds)
	if err != nil {
		return vnetwork.VirtualNetwork{}, err
	}

	return client.Get(ctx, resourceGroupName, resourceName, "")
}

func (m *AzureVNetManager) GetAvailableIP(ctx context.Context, resourceGroup, vnet, subnet string) (string, error) {
	client, err := getVNetClient(m.Creds)
	if err != nil {
		return "", err
	}

	sclient := NewAzureSubnetManager(m.Creds)

	sub, err := sclient.Get(ctx, resourceGroup, vnet, subnet)
	if err != nil {
		return "", err
	}

	if sub.SubnetPropertiesFormat == nil {
		return "", fmt.Errorf("could not find subnet '%s'", subnet)
	}
	prefix := *sub.AddressPrefix
	ip, _, err := net.ParseCIDR(prefix)
	if err != nil {
		return "", err
	}

	result, err := client.CheckIPAddressAvailability(ctx, resourceGroup, vnet, ip.String())
	if err != nil {
		return "", err
	}

	if result.AvailableIPAddresses == nil || len(*result.AvailableIPAddresses) == 0 {
		return "", fmt.Errorf("No available IP addresses in vnet %s", vnet)
	}

	return (*result.AvailableIPAddresses)[0], nil

}
