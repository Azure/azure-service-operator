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

package apimshared

import (
	"context"
	"fmt"
	"strings"

	apim "github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2019-01-01/apimanagement"
	vnet "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-09-01/network"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
)

// GetAPIMgmtSvcClient returns a new instance of an API Svc client
func GetAPIMgmtSvcClient() (apim.ServiceClient, error) {
	client := apim.NewServiceClient(config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		client = apim.ServiceClient{}
	} else {
		client.Authorizer = a
		client.AddToUserAgent(config.UserAgent())
	}
	return client, err
}

// GetVNetClient returns a new instance of an VirtualNetwork client
func GetVNetClient() (vnet.VirtualNetworksClient, error) {
	client := vnet.NewVirtualNetworksClient(config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		client = vnet.VirtualNetworksClient{}
	} else {
		client.Authorizer = a
		client.AddToUserAgent(config.UserAgent())
	}
	return client, err
}

// GetAPIMgmtSvc returns an instance of an APIM service
func GetAPIMgmtSvc(ctx context.Context, resourceGroupName string, resourceName string) (apim.ServiceResource, error) {
	client, err := GetAPIMgmtSvcClient()
	if err != nil {
		return apim.ServiceResource{}, err
	}

	return client.Get(
		ctx,
		resourceGroupName,
		resourceName,
	)
}

// IsAPIMgmtSvcActivated check to see if the API Mgmt Svc has been activated, returns "true" if it has been activated
func IsAPIMgmtSvcActivated(ctx context.Context, resourceGroupName string, resourceName string) (result bool, err error) {
	resource, err := GetAPIMgmtSvc(
		ctx,
		resourceGroupName,
		resourceName,
	)
	if err != nil {
		return false, err
	}

	result = false
	err = fmt.Errorf("Could not evaluate provisioning state of API Mgmt Service: %s, %s", resourceGroupName, resourceName)
	if resource.ServiceProperties != nil && resource.ServiceProperties.ProvisioningState != nil {
		result = !strings.EqualFold(*resource.ServiceProperties.ProvisioningState, "activating")
		err = nil
	}

	return result, err
}

// GetVNetConfigurationByName gets a VNet by name
func GetVNetConfigurationByName(ctx context.Context, resourceGroupName string, resourceName string, subnetName string) (apim.VirtualNetworkConfiguration, error) {
	client, err := GetVNetClient()
	if err != nil {
		return apim.VirtualNetworkConfiguration{}, err
	}

	vnetNetwork, err := client.Get(
		ctx,
		resourceGroupName,
		resourceName,
		"",
	)
	if err != nil {
		return apim.VirtualNetworkConfiguration{}, err
	}

	// find the correct subnet
	correctSubnet := vnet.Subnet{}
	for i := 0; i < len(*vnetNetwork.VirtualNetworkPropertiesFormat.Subnets); i++ {
		subnetCheck := (*vnetNetwork.VirtualNetworkPropertiesFormat.Subnets)[i]
		if strings.EqualFold(*subnetCheck.Name, subnetName) {
			correctSubnet = subnetCheck
		}
	}
	if correctSubnet.Name == nil {
		return apim.VirtualNetworkConfiguration{}, fmt.Errorf("Subnet not found: %s", subnetName)
	}

	result := apim.VirtualNetworkConfiguration{
		Vnetid:           vnetNetwork.ID,
		Subnetname:       &subnetName,
		SubnetResourceID: correctSubnet.ID,
	}

	return result, nil
}
