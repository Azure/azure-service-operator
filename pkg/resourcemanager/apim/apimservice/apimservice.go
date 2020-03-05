// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package apimservice

import (
	"context"
	"fmt"
	"strings"

	apim "github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2019-01-01/apimanagement"
	apimshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/apim/apimshared"
	telemetry "github.com/Azure/azure-service-operator/pkg/telemetry"
	"github.com/Azure/go-autorest/autorest/to"
)

type AzureAPIMgmtServiceManager struct {
	Telemetry telemetry.Telemetry
}

// CreateAPIMgmtSvc creates a new API Mgmt Svc
func (_ *AzureAPIMgmtServiceManager) CreateAPIMgmtSvc(ctx context.Context, tier string, location string, resourceGroupName string, resourceName string, publisherName string, publisherEmail string) (*apim.ServiceResource, error) {
	client, err := apimshared.GetAPIMgmtSvcClient()
	if err != nil {
		return nil, err
	}

	svcProp := apim.ServiceProperties{
		PublisherEmail: &publisherEmail,
		PublisherName:  &publisherName,
	}

	// translate tier / sku type
	skuTypeTier := apim.SkuTypeBasic
	if strings.EqualFold(tier, "standard") {
		skuTypeTier = apim.SkuTypeStandard
	} else if strings.EqualFold(tier, "premium") {
		skuTypeTier = apim.SkuTypePremium
	}
	sku := apim.ServiceSkuProperties{
		Name: skuTypeTier,
	}

	future, err := client.CreateOrUpdate(
		ctx,
		resourceGroupName,
		resourceName,
		apim.ServiceResource{
			Location:          to.StringPtr(location),
			ServiceProperties: &svcProp,
			Sku:               &sku,
		},
	)
	if err != nil {
		return nil, err
	}

	ret, err := future.Result(client)
	return &ret, err
}

// DeleteAPIMgmtSvc deletes an instance of an API Mgmt Svc
func (_ *AzureAPIMgmtServiceManager) DeleteAPIMgmtSvc(ctx context.Context, resourceGroupName string, resourceName string) (*apim.ServiceResource, error) {
	client, err := apimshared.GetAPIMgmtSvcClient()
	if err != nil {
		return nil, err
	}

	future, err := client.Delete(ctx, resourceGroupName, resourceName)
	if err != nil {
		return nil, err
	}

	ret, err := future.Result(client)
	return &ret, err
}

// APIMgmtSvcStatus checks to see if the API Mgmt Svc has been activated
func (_ *AzureAPIMgmtServiceManager) APIMgmtSvcStatus(ctx context.Context, resourceGroupName string, resourceName string) (exists bool, result bool, resourceID *string, err error) {
	return apimshared.APIMgmtSvcStatus(ctx, resourceGroupName, resourceName)
}

// SetVNetForAPIMgmtSvc sets the VNet for an API Mgmt Svc by name (only if it hasn't been previously set)
func (g *AzureAPIMgmtServiceManager) SetVNetForAPIMgmtSvc(ctx context.Context, resourceGroupName string, resourceName string, vnetType string, vnetResourceGroupName string, vnetResourceName string, subnetName string) (err error, updated bool) {

	// check to make sure that the API Mgmt Svc has been activated
	exists, activated, _, err := g.APIMgmtSvcStatus(ctx, resourceGroupName, resourceName)
	if !exists || !activated || err != nil {
		return fmt.Errorf("API Mgmt Service hasn't been created or activated yet: %s, %s", resourceGroupName, resourceName), false
	}

	// translate vnet type
	var vnetTypeConverted apim.VirtualNetworkType
	if strings.EqualFold(vnetType, "external") {
		vnetTypeConverted = apim.VirtualNetworkTypeExternal
	} else if strings.EqualFold(vnetType, "internal") {
		vnetTypeConverted = apim.VirtualNetworkTypeInternal
	} else {
		return nil, false
	}

	// get the subnet configuration
	subnetConfig, err := apimshared.GetSubnetConfigurationByName(ctx, vnetResourceGroupName, vnetResourceName, subnetName)
	if err != nil {
		return err, false
	}

	client, err := apimshared.GetAPIMgmtSvcClient()
	if err != nil {
		return err, false
	}

	// check to make sure that the VPN hasn't already been added to the APIM svc
	apimsvc, err := apimshared.GetAPIMgmtSvc(ctx, resourceGroupName, resourceName)
	if err != nil {
		return err, false
	}
	if apimsvc.ServiceProperties.VirtualNetworkConfiguration != nil {
		if apimsvc.ServiceProperties.VirtualNetworkConfiguration.SubnetResourceID != nil {

			// extract the subnet name from subnetResourceId
			vals := strings.Split(*(*apimsvc.ServiceProperties.VirtualNetworkConfiguration).SubnetResourceID, "/")
			currentSubnet := vals[len(vals)-1]

			if strings.EqualFold(currentSubnet, subnetName) {

				// this is not an error, but the vnet was already set - so exit successfully
				return nil, false
			}
		}
	}

	_, err = client.Update(
		ctx,
		resourceGroupName,
		resourceName,
		apim.ServiceUpdateParameters{
			ServiceUpdateProperties: &apim.ServiceUpdateProperties{
				VirtualNetworkType:          vnetTypeConverted,
				VirtualNetworkConfiguration: &subnetConfig,
			},
		},
	)

	return err, true
}

// SetAppInsightsForAPIMgmtSvc sets the app insight instance to use with the service
func (g *AzureAPIMgmtServiceManager) SetAppInsightsForAPIMgmtSvc(ctx context.Context, resourceGroupName string, resourceName string, appInsightsResourceGroup string, appInsightsName string) error {

	// check to make sure app insight exists
	insight, err := apimshared.GetAppInstanceIDByName(ctx, appInsightsResourceGroup, appInsightsName)
	if err != nil {
		return err
	} else if insight.ID == nil || insight.InstrumentationKey == nil {
		return fmt.Errorf("could not find App Insight %s, %s", appInsightsResourceGroup, appInsightsName)
	}

	// check to make sure that the API Mgmt Svc has been activated
	exists, activated, _, err := g.APIMgmtSvcStatus(ctx, resourceGroupName, resourceName)
	if !exists || !activated || err != nil {
		return fmt.Errorf("API Mgmt Service hasn't been created or activated yet: %s, %s", resourceGroupName, resourceName)
	}

	// get the etag for apim service
	apimsvc, err := apimshared.GetAPIMgmtSvc(
		ctx,
		resourceGroupName,
		resourceName,
	)
	if err != nil {
		return err
	} else if apimsvc.Etag == nil {
		return fmt.Errorf("could not find API Mgmt Service %s, %s", resourceGroupName, resourceName)
	}

	loggerClient, err := apimshared.GetAPIMgmtLoggerClient()
	if err != nil {
		return err
	}

	// check logger config and see if we actually need to update the logger
	contract, err := loggerClient.Get(ctx,
		resourceGroupName,
		resourceName,
		appInsightsName)
	if err != nil {
		return nil
	} else if contract.LoggerContractProperties != nil &&
		contract.LoggerContractProperties.Description != nil &&
		strings.EqualFold(*contract.LoggerContractProperties.Description, appInsightsName) {
		return nil
	}

	credentials := make(map[string]*string)
	credentials["instrumentationKey"] = insight.InstrumentationKey
	_, err = loggerClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		resourceName,
		appInsightsName,
		apim.LoggerContract{
			LoggerContractProperties: &apim.LoggerContractProperties{
				LoggerType:  "ApplicationInsights",
				Description: &appInsightsName,
				ResourceID:  insight.ID,
				Credentials: credentials,
			},
		},
		*apimsvc.Etag,
	)

	return err
}

// CheckAPIMgmtSvcName checks to see if the APIM service name is available
func (g *AzureAPIMgmtServiceManager) CheckAPIMgmtSvcName(ctx context.Context, resourceName string) (available bool, err error) {
	return apimshared.CheckAPIMgmtSvcName(ctx, resourceName)
}
