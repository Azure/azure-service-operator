// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type Api_STATUS struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: API entity contract properties.
	Properties *ApiContractProperties_STATUS `json:"properties,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// API Entity Properties
type ApiContractProperties_STATUS struct {
	// APIVersion: Indicates the version identifier of the API if the API is versioned
	APIVersion *string `json:"apiVersion,omitempty"`

	// ApiRevision: Describes the revision of the API. If no value is provided, default revision 1 is created
	ApiRevision *string `json:"apiRevision,omitempty"`

	// ApiRevisionDescription: Description of the API Revision.
	ApiRevisionDescription *string `json:"apiRevisionDescription,omitempty"`

	// ApiVersionDescription: Description of the API Version.
	ApiVersionDescription *string `json:"apiVersionDescription,omitempty"`

	// ApiVersionSet: Version set details
	ApiVersionSet *ApiVersionSetContractDetails_STATUS `json:"apiVersionSet,omitempty"`

	// ApiVersionSetId: A resource identifier for the related ApiVersionSet.
	ApiVersionSetId *string `json:"apiVersionSetId,omitempty"`

	// AuthenticationSettings: Collection of authentication settings included into this API.
	AuthenticationSettings *AuthenticationSettingsContract_STATUS `json:"authenticationSettings,omitempty"`

	// Contact: Contact information for the API.
	Contact *ApiContactInformation_STATUS `json:"contact,omitempty"`

	// Description: Description of the API. May include HTML formatting tags.
	Description *string `json:"description,omitempty"`

	// DisplayName: API name. Must be 1 to 300 characters long.
	DisplayName *string `json:"displayName,omitempty"`

	// IsCurrent: Indicates if API revision is current api revision.
	IsCurrent *bool `json:"isCurrent,omitempty"`

	// IsOnline: Indicates if API revision is accessible via the gateway.
	IsOnline *bool `json:"isOnline,omitempty"`

	// License: License information for the API.
	License *ApiLicenseInformation_STATUS `json:"license,omitempty"`

	// Path: Relative URL uniquely identifying this API and all of its resource paths within the API Management service
	// instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public
	// URL for this API.
	Path *string `json:"path,omitempty"`

	// Protocols: Describes on which protocols the operations in this API can be invoked.
	Protocols []ApiContractProperties_Protocols_STATUS `json:"protocols,omitempty"`

	// ProvisioningState: The provisioning state
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// ServiceUrl: Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.
	ServiceUrl *string `json:"serviceUrl,omitempty"`

	// SourceApiId: API identifier of the source API.
	SourceApiId *string `json:"sourceApiId,omitempty"`

	// SubscriptionKeyParameterNames: Protocols over which API is made available.
	SubscriptionKeyParameterNames *SubscriptionKeyParameterNamesContract_STATUS `json:"subscriptionKeyParameterNames,omitempty"`

	// SubscriptionRequired: Specifies whether an API or Product subscription is required for accessing the API.
	SubscriptionRequired *bool `json:"subscriptionRequired,omitempty"`

	// TermsOfServiceUrl:  A URL to the Terms of Service for the API. MUST be in the format of a URL.
	TermsOfServiceUrl *string `json:"termsOfServiceUrl,omitempty"`

	// Type: Type of API.
	Type *ApiContractProperties_Type_STATUS `json:"type,omitempty"`
}

// API contact information
type ApiContactInformation_STATUS struct {
	// Email: The email address of the contact person/organization. MUST be in the format of an email address
	Email *string `json:"email,omitempty"`

	// Name: The identifying name of the contact person/organization
	Name *string `json:"name,omitempty"`

	// Url: The URL pointing to the contact information. MUST be in the format of a URL
	Url *string `json:"url,omitempty"`
}

type ApiContractProperties_Protocols_STATUS string

const (
	ApiContractProperties_Protocols_STATUS_Http  = ApiContractProperties_Protocols_STATUS("http")
	ApiContractProperties_Protocols_STATUS_Https = ApiContractProperties_Protocols_STATUS("https")
	ApiContractProperties_Protocols_STATUS_Ws    = ApiContractProperties_Protocols_STATUS("ws")
	ApiContractProperties_Protocols_STATUS_Wss   = ApiContractProperties_Protocols_STATUS("wss")
)

// Mapping from string to ApiContractProperties_Protocols_STATUS
var apiContractProperties_Protocols_STATUS_Values = map[string]ApiContractProperties_Protocols_STATUS{
	"http":  ApiContractProperties_Protocols_STATUS_Http,
	"https": ApiContractProperties_Protocols_STATUS_Https,
	"ws":    ApiContractProperties_Protocols_STATUS_Ws,
	"wss":   ApiContractProperties_Protocols_STATUS_Wss,
}

type ApiContractProperties_Type_STATUS string

const (
	ApiContractProperties_Type_STATUS_Graphql   = ApiContractProperties_Type_STATUS("graphql")
	ApiContractProperties_Type_STATUS_Grpc      = ApiContractProperties_Type_STATUS("grpc")
	ApiContractProperties_Type_STATUS_Http      = ApiContractProperties_Type_STATUS("http")
	ApiContractProperties_Type_STATUS_Odata     = ApiContractProperties_Type_STATUS("odata")
	ApiContractProperties_Type_STATUS_Soap      = ApiContractProperties_Type_STATUS("soap")
	ApiContractProperties_Type_STATUS_Websocket = ApiContractProperties_Type_STATUS("websocket")
)

// Mapping from string to ApiContractProperties_Type_STATUS
var apiContractProperties_Type_STATUS_Values = map[string]ApiContractProperties_Type_STATUS{
	"graphql":   ApiContractProperties_Type_STATUS_Graphql,
	"grpc":      ApiContractProperties_Type_STATUS_Grpc,
	"http":      ApiContractProperties_Type_STATUS_Http,
	"odata":     ApiContractProperties_Type_STATUS_Odata,
	"soap":      ApiContractProperties_Type_STATUS_Soap,
	"websocket": ApiContractProperties_Type_STATUS_Websocket,
}

// API license information
type ApiLicenseInformation_STATUS struct {
	// Name: The license name used for the API
	Name *string `json:"name,omitempty"`

	// Url: A URL to the license used for the API. MUST be in the format of a URL
	Url *string `json:"url,omitempty"`
}

// An API Version Set contains the common configuration for a set of API Versions relating
type ApiVersionSetContractDetails_STATUS struct {
	// Description: Description of API Version Set.
	Description *string `json:"description,omitempty"`

	// Id: Identifier for existing API Version Set. Omit this value to create a new Version Set.
	Id *string `json:"id,omitempty"`

	// Name: The display Name of the API Version Set.
	Name *string `json:"name,omitempty"`

	// VersionHeaderName: Name of HTTP header parameter that indicates the API Version if versioningScheme is set to `header`.
	VersionHeaderName *string `json:"versionHeaderName,omitempty"`

	// VersionQueryName: Name of query parameter that indicates the API Version if versioningScheme is set to `query`.
	VersionQueryName *string `json:"versionQueryName,omitempty"`

	// VersioningScheme: An value that determines where the API Version identifier will be located in a HTTP request.
	VersioningScheme *ApiVersionSetContractDetails_VersioningScheme_STATUS `json:"versioningScheme,omitempty"`
}

// API Authentication Settings.
type AuthenticationSettingsContract_STATUS struct {
	// OAuth2: OAuth2 Authentication settings
	OAuth2 *OAuth2AuthenticationSettingsContract_STATUS `json:"oAuth2,omitempty"`

	// OAuth2AuthenticationSettings: Collection of OAuth2 authentication settings included into this API.
	OAuth2AuthenticationSettings []OAuth2AuthenticationSettingsContract_STATUS `json:"oAuth2AuthenticationSettings,omitempty"`

	// Openid: OpenID Connect Authentication Settings
	Openid *OpenIdAuthenticationSettingsContract_STATUS `json:"openid,omitempty"`

	// OpenidAuthenticationSettings: Collection of Open ID Connect authentication settings included into this API.
	OpenidAuthenticationSettings []OpenIdAuthenticationSettingsContract_STATUS `json:"openidAuthenticationSettings,omitempty"`
}

// Subscription key parameter names details.
type SubscriptionKeyParameterNamesContract_STATUS struct {
	// Header: Subscription key header name.
	Header *string `json:"header,omitempty"`

	// Query: Subscription key query string parameter name.
	Query *string `json:"query,omitempty"`
}

type ApiVersionSetContractDetails_VersioningScheme_STATUS string

const (
	ApiVersionSetContractDetails_VersioningScheme_STATUS_Header  = ApiVersionSetContractDetails_VersioningScheme_STATUS("Header")
	ApiVersionSetContractDetails_VersioningScheme_STATUS_Query   = ApiVersionSetContractDetails_VersioningScheme_STATUS("Query")
	ApiVersionSetContractDetails_VersioningScheme_STATUS_Segment = ApiVersionSetContractDetails_VersioningScheme_STATUS("Segment")
)

// Mapping from string to ApiVersionSetContractDetails_VersioningScheme_STATUS
var apiVersionSetContractDetails_VersioningScheme_STATUS_Values = map[string]ApiVersionSetContractDetails_VersioningScheme_STATUS{
	"header":  ApiVersionSetContractDetails_VersioningScheme_STATUS_Header,
	"query":   ApiVersionSetContractDetails_VersioningScheme_STATUS_Query,
	"segment": ApiVersionSetContractDetails_VersioningScheme_STATUS_Segment,
}

// API OAuth2 Authentication settings details.
type OAuth2AuthenticationSettingsContract_STATUS struct {
	// AuthorizationServerId: OAuth authorization server identifier.
	AuthorizationServerId *string `json:"authorizationServerId,omitempty"`

	// Scope: operations scope.
	Scope *string `json:"scope,omitempty"`
}

// API OAuth2 Authentication settings details.
type OpenIdAuthenticationSettingsContract_STATUS struct {
	// BearerTokenSendingMethods: How to send token to the server.
	BearerTokenSendingMethods []BearerTokenSendingMethodsContract_STATUS `json:"bearerTokenSendingMethods,omitempty"`

	// OpenidProviderId: OAuth authorization server identifier.
	OpenidProviderId *string `json:"openidProviderId,omitempty"`
}

// Form of an authorization grant, which the client uses to request the access token.
type BearerTokenSendingMethodsContract_STATUS string

const (
	BearerTokenSendingMethodsContract_STATUS_AuthorizationHeader = BearerTokenSendingMethodsContract_STATUS("authorizationHeader")
	BearerTokenSendingMethodsContract_STATUS_Query               = BearerTokenSendingMethodsContract_STATUS("query")
)

// Mapping from string to BearerTokenSendingMethodsContract_STATUS
var bearerTokenSendingMethodsContract_STATUS_Values = map[string]BearerTokenSendingMethodsContract_STATUS{
	"authorizationheader": BearerTokenSendingMethodsContract_STATUS_AuthorizationHeader,
	"query":               BearerTokenSendingMethodsContract_STATUS_Query,
}
