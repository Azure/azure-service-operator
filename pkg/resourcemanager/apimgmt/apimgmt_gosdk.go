// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package apimgmt

import (
	"net/http"
	"strings"

	api "github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2019-01-01/apimanagement"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

// returns a valid instance of an API client
func getAPIClient() api.APIClient {
	apiClient := api.NewAPIClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	apiClient.Authorizer = a
	apiClient.AddToUserAgent(config.UserAgent())
	return apiClient
}

// returns a new instance of an API Svc client
func getAPISvcClient() api.ServiceClient {
	serviceClient := api.NewServiceClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	serviceClient.Authorizer = a
	serviceClient.AddToUserAgent(config.UserAgent())
	return serviceClient
}

// CreateAPIMgmtSvc creates an instance of an API Management service
// wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2019-01-01/apimanagement#ServiceClient.CreateOrUpdate
func (sdk GoSDKClient) CreateAPIMgmtSvc() (result api.ServiceResource, err error) {
	serviceClient := getAPISvcClient()

	// create the service
	svcProp := api.ServiceProperties{
		PublisherEmail: &sdk.Email,
		PublisherName:  &sdk.Name,
	}
	sku := api.ServiceSkuProperties{
		Name: api.SkuTypeBasic,
	}
	future, err := serviceClient.CreateOrUpdate(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServiceName,
		api.ServiceResource{
			Location:          to.StringPtr(config.Location()),
			ServiceProperties: &svcProp,
			Sku:               &sku,
		},
	)
	if err != nil {
		return result, err
	}

	return future.Result(serviceClient)
}

// CreateOrUpdateAPI creates an API endpoint on an API Management service
// wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2019-01-01/apimanagement#APIClient.CreateOrUpdate
func (sdk GoSDKClient) CreateOrUpdateAPI(properties APIProperties) (result api.APIContract, err error) {
	apiClient := getAPIClient()

	apiParameter := APIPropertiesToAPICreateOrUpdateParameter(properties)
	future, err := apiClient.CreateOrUpdate(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServiceName,
		properties.APIID,
		apiParameter,
		properties.IFMatch,
	)
	if err != nil {
		return result, err
	}

	return future.Result(apiClient)
}

// DeleteAPI deletes an API endpoint on an API Management service
// wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2019-01-01/apimanagement#APIClient.Delete
func (sdk GoSDKClient) DeleteAPI(apiid string, ifMatch string) (result autorest.Response, err error) {
	apiClient := getAPIClient()

	// check to see if the api endpoint exists, if it does then short-circuit
	_, err = apiClient.Get(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServiceName,
		apiid,
	)
	if err != nil {
		result = autorest.Response{
			Response: &http.Response{
				StatusCode: 200,
			},
		}
		return result, nil
	}

	return apiClient.Delete(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServiceName,
		apiid,
		ifMatch,
		to.BoolPtr(true),
	)
}

// DeleteAPIMgmtSvc deletes an instance of an API Management service
// wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2019-01-01/apimanagement#ServiceClient.Delete
func (sdk GoSDKClient) DeleteAPIMgmtSvc() (result api.ServiceResource, err error) {
	serviceClient := getAPISvcClient()

	// check to see if the api mgmt svc exists, if it does then short-circuit
	_, err = serviceClient.Get(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServiceName,
	)
	if err != nil {
		return result, nil
	}

	future, err := serviceClient.Delete(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServiceName,
	)
	if err != nil {
		return result, err
	}

	return future.Result(serviceClient)
}

// IsAPIMgmtSvcActivated check to see if the API Mgmt Svc has been activated, returns "true" if it has been
// activated. CALL THIS FUNCTION BEFORE CREATING THE SERVICE in the reconcile loop.
func (sdk GoSDKClient) IsAPIMgmtSvcActivated() (result bool, err error) {
	serviceClient := getAPISvcClient()

	resource, err := serviceClient.Get(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServiceName,
	)
	if err != nil {
		return false, err
	}

	result = false
	if strings.Compare(*resource.ServiceProperties.ProvisioningState, "Activating") != 0 {
		result = true
	}

	return result, err
}
