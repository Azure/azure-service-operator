/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vnet

import (
	"context"

	vnetwork "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
)

// AzureVNetManager is the struct that the manager functions hang off
type AzureVNetManager struct{}

// getVNetClient returns a new instance of an VirtualNetwork client
func getVNetClient() (vnetwork.VirtualNetworksClient, error) {
	client := vnetwork.NewVirtualNetworksClient(config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		client = vnetwork.VirtualNetworksClient{}
	} else {
		client.Authorizer = a
		client.AddToUserAgent(config.UserAgent())
	}
	return client, err
}

// getSubnetsClient returns a new instance of an VirtualNetwork client
func getSubnetsClient() (vnetwork.SubnetsClient, error) {
	client := vnetwork.NewSubnetsClient(config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		client = vnetwork.SubnetsClient{}
	} else {
		client.Authorizer = a
		client.AddToUserAgent(config.UserAgent())
	}
	return client, err
}

// CreateVNet creates VNets
func (_ *AzureAPIMgmtServiceManager) CreateVNet(ctx context.Context, resourceGroupName string, resourceName string) (future, error) {
	client, err := apimshared.getVNetClient()
	if err != nil {
		return apim.ServiceResource{}, err
	}

}

// CreateSubnet creates subnets
func (_ *AzureAPIMgmtServiceManager) CreateSubnet(ctx context.Context, resourceGroupName string, resourceName string, subnetName string, subnetAddressPrefix string) (future, error) {
	client, err := apimshared.getSubnetsClient()
	if err != nil {
		return apim.ServiceResource{}, err
	}

}

// DeleteVNet deletes a VNet
func (_ *AzureAPIMgmtServiceManager) DeleteVNet(ctx context.Context, resourceGroupName string, resourceName string) (future, error) {
	client, err := apimshared.getVNetClient()
	if err != nil {
		return apim.ServiceResource{}, err
	}

}
