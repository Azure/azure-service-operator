// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type ApiVersionSet_STATUS struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: API VersionSet contract properties.
	Properties *ApiVersionSetContractProperties_STATUS `json:"properties,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// Properties of an API Version Set.
type ApiVersionSetContractProperties_STATUS struct {
	// Description: Description of API Version Set.
	Description *string `json:"description,omitempty"`

	// DisplayName: Name of API Version Set
	DisplayName *string `json:"displayName,omitempty"`

	// VersionHeaderName: Name of HTTP header parameter that indicates the API Version if versioningScheme is set to `header`.
	VersionHeaderName *string `json:"versionHeaderName,omitempty"`

	// VersionQueryName: Name of query parameter that indicates the API Version if versioningScheme is set to `query`.
	VersionQueryName *string `json:"versionQueryName,omitempty"`

	// VersioningScheme: An value that determines where the API Version identifier will be located in a HTTP request.
	VersioningScheme *ApiVersionSetContractProperties_VersioningScheme_STATUS `json:"versioningScheme,omitempty"`
}

type ApiVersionSetContractProperties_VersioningScheme_STATUS string

const (
	ApiVersionSetContractProperties_VersioningScheme_STATUS_Header  = ApiVersionSetContractProperties_VersioningScheme_STATUS("Header")
	ApiVersionSetContractProperties_VersioningScheme_STATUS_Query   = ApiVersionSetContractProperties_VersioningScheme_STATUS("Query")
	ApiVersionSetContractProperties_VersioningScheme_STATUS_Segment = ApiVersionSetContractProperties_VersioningScheme_STATUS("Segment")
)

// Mapping from string to ApiVersionSetContractProperties_VersioningScheme_STATUS
var apiVersionSetContractProperties_VersioningScheme_STATUS_Values = map[string]ApiVersionSetContractProperties_VersioningScheme_STATUS{
	"header":  ApiVersionSetContractProperties_VersioningScheme_STATUS_Header,
	"query":   ApiVersionSetContractProperties_VersioningScheme_STATUS_Query,
	"segment": ApiVersionSetContractProperties_VersioningScheme_STATUS_Segment,
}