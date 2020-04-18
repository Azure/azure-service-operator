// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package nic

import (
	"context"

	vnetwork "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"k8s.io/apimachinery/pkg/runtime"
)

type AzureNetworkInterfaceClient struct {
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func NewAzureNetworkInterfaceClient(secretclient secrets.SecretClient, scheme *runtime.Scheme) *AzureNetworkInterfaceClient {
	return &AzureNetworkInterfaceClient{
		SecretClient: secretclient,
		Scheme:       scheme,
	}
}

func getNetworkInterfaceClient() vnetwork.InterfacesClient {
	nicClient := vnetwork.NewInterfacesClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	nicClient.Authorizer = a
	nicClient.AddToUserAgent(config.UserAgent())
	return nicClient
}

func (m *AzureNetworkInterfaceClient) CreateNetworkInterface(ctx context.Context, location string, resourceGroupName string, resourceName string, vnetName string, subnetName string, publicIPAddressName string) (future vnetwork.InterfacesCreateOrUpdateFuture, err error) {

	client := getNetworkInterfaceClient()

	subnetIDInput := helpers.MakeResourceID(
		client.SubscriptionID,
		resourceGroupName,
		"Microsoft.Network",
		"virtualNetworks",
		vnetName,
		"subnets",
		subnetName,
	)

	publicIPAddressIDInput := helpers.MakeResourceID(
		client.SubscriptionID,
		resourceGroupName,
		"Microsoft.Network",
		"publicIPAddresses",
		publicIPAddressName,
		"",
		"",
	)

	var ipConfigsToAdd []vnetwork.InterfaceIPConfiguration
	ipConfigsToAdd = append(
		ipConfigsToAdd,
		vnetwork.InterfaceIPConfiguration{
			Name: &resourceName,
			InterfaceIPConfigurationPropertiesFormat: &vnetwork.InterfaceIPConfigurationPropertiesFormat{
				Subnet: &vnetwork.Subnet{
					ID: &subnetIDInput,
				},
				PublicIPAddress: &vnetwork.PublicIPAddress{
					ID: &publicIPAddressIDInput,
				},
			},
		},
	)

	future, err = client.CreateOrUpdate(
		ctx,
		resourceGroupName,
		resourceName,
		vnetwork.Interface{
			Location: &location,
			InterfacePropertiesFormat: &vnetwork.InterfacePropertiesFormat{
				IPConfigurations: &ipConfigsToAdd,
			},
		},
	)

	return future, err
}

func (m *AzureNetworkInterfaceClient) DeleteNetworkInterface(ctx context.Context, nicName string, resourcegroup string) (status string, err error) {

	client := getNetworkInterfaceClient()

	_, err = client.Get(ctx, resourcegroup, nicName, "")
	if err == nil { // nic present, so go ahead and delete
		future, err := client.Delete(ctx, resourcegroup, nicName)
		return future.Status(), err
	}
	// nic not present so return success anyway
	return "nic not present", nil

}

func (m *AzureNetworkInterfaceClient) GetNetworkInterface(ctx context.Context, resourcegroup string, nicName string) (nic vnetwork.Interface, err error) {

	client := getNetworkInterfaceClient()

	return client.Get(ctx, resourcegroup, nicName, "")
}
