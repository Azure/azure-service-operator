// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package nic

import (
	"context"
	"strings"

	vnetwork "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
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

func MakeResourceId(subscriptionId string, resourceGroupName string, provider string, resourceType string, resourceName string, subResourceType string, subResourceName string) string {
	// Sample 1: /subscriptions/88fd8cb2-8248-499e-9a2d-4929a4b0133c/resourceGroups/resourcegroup-azure-operators/providers/Microsoft.Network/publicIPAddresses/azurepublicipaddress-sample-3
	// Sample 2: /subscriptions/88fd8cb2-8248-499e-9a2d-4929a4b0133c/resourceGroups/resourcegroup-azure-operators/providers/Microsoft.Network/virtualNetworks/vnet-sample-hpf-1/subnets/test2
	segments := []string{
		"subscriptions",
		subscriptionId,
		"resourceGroups",
		resourceGroupName,
		"providers",
		provider,
		resourceType,
		resourceName,
		subResourceType,
		subResourceName,
	}

	result := "/" + strings.Join(segments, "/")

	return result
}

func (m *AzureNetworkInterfaceClient) CreateNetworkInterface(ctx context.Context, location string, resourceGroupName string, resourceName string, vnetName string, subnetName string, publicIPAddressName string) (future vnetwork.InterfacesCreateOrUpdateFuture, err error) {

	client := getNetworkInterfaceClient()

	subnetIDInput := MakeResourceId(
		client.SubscriptionID,
		resourceGroupName,
		"Microsoft.Network",
		"virtualNetworks",
		vnetName,
		"subnets",
		subnetName,
	)

	publicIPAddressIDInput := MakeResourceId(
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
