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
	Telemetry telemetry.PrometheusTelemetry
}

// CreateAPIMgmtSvc creates a new API Mgmt Svc
func (_ *AzureAPIMgmtServiceManager) CreateAPIMgmtSvc(ctx context.Context, location string, resourceGroupName string, resourceName string, publisherName string, publisherEmail string) (apim.ServiceResource, error) {
	client, err := apimshared.GetAPIMgmtSvcClient()
	if err != nil {
		return apim.ServiceResource{}, err
	}

	svcProp := apim.ServiceProperties{
		PublisherEmail: &publisherEmail,
		PublisherName:  &publisherName,
	}
	sku := apim.ServiceSkuProperties{
		Name: apim.SkuTypeBasic,
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
		return apim.ServiceResource{}, err
	}

	return future.Result(client)
}

// DeleteAPIMgmtSvc deletes an instance of an API Mgmt Svc
func (_ *AzureAPIMgmtServiceManager) DeleteAPIMgmtSvc(ctx context.Context, resourceGroupName string, resourceName string) (apim.ServiceResource, error) {
	client, err := apimshared.GetAPIMgmtSvcClient()
	if err != nil {
		return apim.ServiceResource{}, err
	}

	result, err := client.Delete(ctx, resourceGroupName, resourceName)
	if err != nil {
		return apim.ServiceResource{}, err
	}

	return result.Result(client)
}

// APIMgmtSvcStatus checks to see if the API Mgmt Svc has been activated
func (_ *AzureAPIMgmtServiceManager) APIMgmtSvcStatus(ctx context.Context, resourceGroupName string, resourceName string) (exists bool, result bool, err error) {
	return apimshared.APIMgmtSvcStatus(ctx, resourceGroupName, resourceName)
}

// SetVNetForAPIMgmtSvc sets the VNet for an API Mgmt Svc by name
func (g *AzureAPIMgmtServiceManager) SetVNetForAPIMgmtSvc(ctx context.Context, resourceGroupName string, resourceName string, vnetType string, vnetResourceGroupName string, vnetResourceName string, subnetName string) error {

	// check to make sure that the API Mgmt Svc has been activated
	exists, activated, err := g.APIMgmtSvcStatus(ctx, resourceGroupName, resourceName)
	if !exists || !activated || err != nil {
		return fmt.Errorf("API Mgmt Service hasn't been created or activated yet: %s, %s", resourceGroupName, resourceName)
	}

	// translate vnet type
	var vnetTypeConverted apim.VirtualNetworkType
	if strings.EqualFold(vnetType, "external") {
		vnetTypeConverted = apim.VirtualNetworkTypeExternal
	} else if strings.EqualFold(vnetType, "internal") {
		vnetTypeConverted = apim.VirtualNetworkTypeInternal
	} else {
		return nil
	}

	// get the subnet configuration
	subnetConfig, err := apimshared.GetSubnetConfigurationByName(ctx, vnetResourceGroupName, vnetResourceName, subnetName)
	if err != nil {
		return err
	}

	client, err := apimshared.GetAPIMgmtSvcClient()
	if err != nil {
		return err
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

	return err
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
	exists, activated, err := g.APIMgmtSvcStatus(ctx, resourceGroupName, resourceName)
	if !exists || !activated || err != nil {
		return fmt.Errorf("API Mgmt Service hasn't been created or activated yet: %s, %s", resourceGroupName, resourceName)
	}

	// get the etag for apim service
	apimSvc, err := apimshared.GetAPIMgmtSvc(
		ctx,
		resourceGroupName,
		resourceName,
	)
	if err != nil {
		return err
	} else if apimSvc.Etag == nil {
		return fmt.Errorf("could not find API Mgmt Service %s, %s", resourceGroupName, resourceName)
	}

	client, err := apimshared.GetAPIMgmtLoggerClient()
	if err != nil {
		return err
	}

	var credentials map[string]*string
	credentials["instrumentationKey"] = insight.InstrumentationKey
	_, err = client.CreateOrUpdate(
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
		*apimSvc.Etag,
	)

	return err
}

// CheckAPIMgmtSvcName checks to see if the APIM service name is available
func (g *AzureAPIMgmtServiceManager) CheckAPIMgmtSvcName(ctx context.Context, resourceName string) (available bool, err error) {
	return apimshared.CheckAPIMgmtSvcName(ctx, resourceName)
}
