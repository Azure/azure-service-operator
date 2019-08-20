// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package apimgmt

import (
	"strings"

	api "github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2019-01-01/apimanagement"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest/to"
)

// returns a valid instance of an API client
func getAPIClient() (api.APIClient, error) {
	apiClient := api.NewAPIClient(config.SubscriptionID())
	var err error
	a, err := iam.GetResourceManagementAuthorizer()
	if err == nil {
		apiClient.Authorizer = a
		apiClient.AddToUserAgent(config.UserAgent())
	}
	return apiClient, err
}

// returns a new instance of an API Svc client
func getAPISvcClient() (api.ServiceClient, error) {
	serviceClient := api.NewServiceClient(config.SubscriptionID())
	var err error
	a, err := iam.GetResourceManagementAuthorizer()
	if err == nil {
		serviceClient.Authorizer = a
		serviceClient.AddToUserAgent(config.UserAgent())
	}
	return serviceClient, err
}

// CreateAPIMgmtSvcImpl creates an instance of an API Management Service. Returns "true" if successful,
// wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2019-01-01/apimanagement#ServiceClient.CreateOrUpdate
func (sdk GoSDKClient) CreateAPIMgmtSvcImpl() (result bool, err error) {
	serviceClient, err := getAPISvcClient()
	if err != nil {
		return false, err
	}

	svcProp := api.ServiceProperties{
		PublisherEmail: &sdk.Email,
		PublisherName:  &sdk.Name,
	}
	sku := api.ServiceSkuProperties{
		Name: api.SkuTypeBasic,
	}
	_, err = serviceClient.CreateOrUpdate(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServiceName,
		api.ServiceResource{
			Location:          to.StringPtr(config.Location()),
			ServiceProperties: &svcProp,
			Sku:               &sku,
		},
	)

	result = (err == nil)
	return result, err
}

// CreateOrUpdateAPIImpl creates an API endpoint on an API Management Service. Returns "true" if successful,
// wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2019-01-01/apimanagement#APIClient.CreateOrUpdate
func (sdk GoSDKClient) CreateOrUpdateAPIImpl(apiid string, properties APIProperties, ifMatch string) (result bool, err error) {
	apiClient, err := getAPIClient()
	if err != nil {
		return false, err
	}

	APIParameter := APIPropertiesToAPICreateOrUpdateParameter(properties)

	future, err := apiClient.CreateOrUpdate(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServiceName,
		apiid,
		APIParameter,
		ifMatch,
	)

	err = future.WaitForCompletionRef(sdk.Ctx, apiClient.Client)

	// TODO: Amanda add the code to create "Blank", "WADL", and "OpenAPI" api endpoints here. You may need to alter APIProperties.go. Also, add unit tests
	// Note: Type is indicated as a field in APIProperties struct

	result = (err == nil)
	return result, err
}

// DeleteAPIImpl deletes an API endpoint on an API Management Service. Returns "true" if successful,
// wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2019-01-01/apimanagement#APIClient.Delete
func (sdk GoSDKClient) DeleteAPIImpl(apiid string) (result bool, err error) {
	// apiClient, err := getAPIClient()

	// TODO: Amanda add the code to delete "Blank", "WADL", and "OpenAPI" api endpoints here. You may need to alter APIProperties.go. Also, add unit tests

	return true, err
}

// DeleteAPIMgmtSvcImpl deletes an instance of an API Management Service. Returns "true" if successful,
// wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2019-01-01/apimanagement#ServiceClient.Delete
func (sdk GoSDKClient) DeleteAPIMgmtSvcImpl() (result bool, err error) {
	serviceClient, err := getAPISvcClient()
	if err != nil {
		return false, err
	}

	_, err = serviceClient.Delete(
		sdk.Ctx,
		sdk.ResourceGroupName,
		sdk.ServiceName,
	)

	result = (err == nil)
	return result, err
}

// IsAPIMgmtSvcActivatedImpl check to see if the API Mgmt Svc has been activated, returns "true" if it has been.
func (sdk GoSDKClient) IsAPIMgmtSvcActivatedImpl() (result bool, err error) {
	serviceClient, err := getAPISvcClient()
	if err != nil {
		return false, err
	}

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
