// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package apimshared

import (
	"context"
	"fmt"
	"strings"

	apim "github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2019-01-01/apimanagement"
	insights "github.com/Azure/azure-sdk-for-go/services/appinsights/mgmt/2015-05-01/insights"
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

// GetAPIMgmtLoggerClient returns a new instance of an VirtualNetwork client
func GetAPIMgmtLoggerClient() (apim.LoggerClient, error) {
	client := apim.NewLoggerClient(config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		client = apim.LoggerClient{}
	} else {
		client.Authorizer = a
		client.AddToUserAgent(config.UserAgent())
	}
	return client, err
}

// GetInsightsClient retrieves a client
func GetInsightsClient() (insights.ComponentsClient, error) {
	client := insights.NewComponentsClient(config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		client = insights.ComponentsClient{}
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

// APIMgmtSvcStatus check to see if the API Mgmt Svc has been activated, returns "true" if it has been activated
func APIMgmtSvcStatus(ctx context.Context, resourceGroupName string, resourceName string) (exists bool, result bool, resourceID *string, err error) {
	resource, err := GetAPIMgmtSvc(
		ctx,
		resourceGroupName,
		resourceName,
	)
	if err != nil {
		return false, false, nil, err
	} else if resource.Name == nil {
		return false, false, nil, nil
	}

	result = false
	err = fmt.Errorf("Could not evaluate provisioning state of API Mgmt Service: %s, %s", resourceGroupName, resourceName)
	if resource.ServiceProperties != nil && resource.ServiceProperties.ProvisioningState != nil {
		if strings.EqualFold(*resource.ServiceProperties.ProvisioningState, "succeeded") {
			result = true
		}
		err = nil
	}

	return true, result, resource.ID, err
}

// GetSubnetConfigurationByName gets a VNet by name
func GetSubnetConfigurationByName(ctx context.Context, resourceGroupName string, resourceName string, subnetName string) (apim.VirtualNetworkConfiguration, error) {
	client, err := GetVNetClient()
	if err != nil {
		return apim.VirtualNetworkConfiguration{}, err
	}

	// get the vnet
	vnetNetwork, err := client.Get(
		ctx,
		resourceGroupName,
		resourceName,
		"",
	)
	if err != nil {
		return apim.VirtualNetworkConfiguration{}, err
	}

	// look for the correct subnet
	correctSubnet := vnet.Subnet{}
	for i := 0; i < len(*vnetNetwork.VirtualNetworkPropertiesFormat.Subnets); i++ {
		subnetCheck := (*vnetNetwork.VirtualNetworkPropertiesFormat.Subnets)[i]
		if strings.EqualFold(*subnetCheck.Name, subnetName) {
			correctSubnet = subnetCheck
			break
		}
	}
	if correctSubnet.Name == nil {
		return apim.VirtualNetworkConfiguration{}, fmt.Errorf("Subnet was not found: %s", subnetName)
	}

	// the subnet was found, return it
	result := apim.VirtualNetworkConfiguration{
		Vnetid:           vnetNetwork.ID,
		Subnetname:       &subnetName,
		SubnetResourceID: correctSubnet.ID,
	}

	return result, nil
}

// CheckAPIMgmtSvcName checks to see if the APIM service name is available
func CheckAPIMgmtSvcName(ctx context.Context, resourceName string) (available bool, err error) {
	client, err := GetAPIMgmtSvcClient()
	if err != nil {
		return false, err
	}

	result, err := client.CheckNameAvailability(
		ctx,
		apim.ServiceCheckNameAvailabilityParameters{
			Name: &resourceName,
		},
	)

	nameAvailable := false
	if result.NameAvailable != nil {
		nameAvailable = *result.NameAvailable
	}

	return nameAvailable, err
}

// GetAppInstanceIDByName retrieves an app insight by name
func GetAppInstanceIDByName(ctx context.Context, resourceGroup string, resourceName string) (insights.ApplicationInsightsComponent, error) {
	client, err := GetInsightsClient()
	if err != nil {
		return insights.ApplicationInsightsComponent{}, err
	}

	return client.Get(
		ctx,
		resourceGroup,
		resourceName,
	)
}

// GetAPIMClient returns a pointer to an API Management client
func GetAPIMClient() apim.APIClient {
	apimClient := apim.NewAPIClient(config.SubscriptionID())

	a, err := iam.GetResourceManagementAuthorizer()
	apimClient.Authorizer = a
	apimClient.AddToUserAgent(config.UserAgent())

	if err != nil {
		panic(err)
	}

	return apimClient
}
