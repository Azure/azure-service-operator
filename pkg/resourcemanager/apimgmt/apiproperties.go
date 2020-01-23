// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package apimgmt

import (
	api "github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2019-01-01/apimanagement"
)

// APIProperties contains values needed for adding / updating API endpoints for API Mgmt Svc's,
// wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2019-01-01/apimanagement#APICreateOrUpdateProperties
type APIProperties struct {

	// REQUIRED FOR ALL 3 - BLANK, OPENAPI, AND WADL

	// Format - Format of the Content in which the API is getting imported. Possible values include: 'WadlXML', 'WadlLinkJSON', 'SwaggerJSON', 'SwaggerLinkJSON', 'Wsdl', 'WsdlLink', 'Openapi', 'Openapijson', 'OpenapiLink'
	Format api.ContentFormat

	// DisplayName - API name. Must be 1 to 300 characters long.
	DisplayName *string

	// REQUIRED FOR OPENAPI + WADL ONLY

	// Value - Content value when Importing an API (OpenAPI or WADL Specification?).
	Value *string

	// REQUIRED FOR OPENAPI ONLY

	// Protocols - Describes on which protocols the operations in this API can be invoked (HTTP, HTTPS, BOTH?)
	Protocols *[]api.Protocol

	// OTHER FIELDS WE DISCOVERED WERE REQUIRED DURING TESTING

	// Path - Relative URL uniquely identifying this API and all of its resource paths within the API Management service instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public URL for this API.
	Path *string

	// APIID is the API revision identifier. Must be unique in the current API Management service instance.
	APIID string

	// IFMatchis the eTag of the Entity. ETag should match the current entity state from the header response of the GET request or it should be * for unconditional update.
	IFMatch string

	// OPTIONAL BUT NICE TO HAVE

	// Description - Description of the API. May include HTML formatting tags.
	// Description *string `json:"description,omitempty"`

	// ServiceURL - Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.
	// ServiceURL *string `json:"serviceUrl,omitempty"`

	// SubscriptionRequired - Specifies whether an API or Product subscription is required for accessing the API.
	// SubscriptionRequired *bool `json:"subscriptionRequired,omitempty"`

	// AuthenticationSettings - Collection of authentication settings included into this API.
	// AuthenticationSettings *AuthenticationSettingsContract `json:"authenticationSettings,omitempty"`

	// TODO: What's the difference between this and "protocols"?  Which one is the required field at OpenAPI setup?
	// SubscriptionKeyParameterNames - Protocols over which API is made available.
	// SubscriptionKeyParameterNames *SubscriptionKeyParameterNamesContract `json:"subscriptionKeyParameterNames,omitempty"`
}

// APIPropertiesToAPICreateOrUpdateParameter translates APIProperties to APICreateOrUpdateParameter and underlying APICreateOrUpdateProperties
func APIPropertiesToAPICreateOrUpdateParameter(parameter APIProperties) (result api.APICreateOrUpdateParameter) {
	var createOrUpdateParameter api.APICreateOrUpdateParameter

	var apiCreateOrUpdateProperties api.APICreateOrUpdateProperties

	// Change to fields that I've added in the struct above
	apiCreateOrUpdateProperties.Format = parameter.Format
	apiCreateOrUpdateProperties.DisplayName = parameter.DisplayName
	apiCreateOrUpdateProperties.Value = parameter.Value
	apiCreateOrUpdateProperties.Protocols = parameter.Protocols
	apiCreateOrUpdateProperties.Path = parameter.Path

	createOrUpdateParameter = api.APICreateOrUpdateParameter{
		APICreateOrUpdateProperties: &apiCreateOrUpdateProperties}

	return createOrUpdateParameter
}
