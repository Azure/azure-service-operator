---
title: apimanagement.azure.com/v1api20230501preview
---
<h2 id="apimanagement.azure.com/v1api20230501preview">apimanagement.azure.com/v1api20230501preview</h2>
<div>
<p>Package v1api20230501preview contains API Schema definitions for the apimanagement v1api20230501preview API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="apimanagement.azure.com/v1api20230501preview.APIVersion">APIVersion
(<code>string</code> alias)</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;2023-05-01-preview&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AdditionalLocation">AdditionalLocation
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Spec">Service_Spec</a>)
</p>
<div>
<p>Description of an additional API Management resource location.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>disableGateway</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableGateway: Property only valid for an Api Management service deployed in multiple locations. This can be used to
disable the gateway in this additional location.</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: The location name of the additional region among Azure Data center regions.</p>
</td>
</tr>
<tr>
<td>
<code>natGatewayState</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AdditionalLocation_NatGatewayState">
AdditionalLocation_NatGatewayState
</a>
</em>
</td>
<td>
<p>NatGatewayState: Property can be used to enable NAT Gateway for this API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>publicIpAddressReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>PublicIpAddressReference: Public Standard SKU IP V4 based IP address to be associated with Virtual Network deployed
service in the location. Supported only for Premium SKU being deployed in Virtual Network.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceSkuProperties">
ApiManagementServiceSkuProperties
</a>
</em>
</td>
<td>
<p>Sku: SKU properties of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkConfiguration</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.VirtualNetworkConfiguration">
VirtualNetworkConfiguration
</a>
</em>
</td>
<td>
<p>VirtualNetworkConfiguration: Virtual network configuration for the location.</p>
</td>
</tr>
<tr>
<td>
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: A list of availability zones denoting where the resource needs to come from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AdditionalLocation_ARM">AdditionalLocation_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_ARM">ApiManagementServiceProperties_ARM</a>)
</p>
<div>
<p>Description of an additional API Management resource location.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>disableGateway</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableGateway: Property only valid for an Api Management service deployed in multiple locations. This can be used to
disable the gateway in this additional location.</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: The location name of the additional region among Azure Data center regions.</p>
</td>
</tr>
<tr>
<td>
<code>natGatewayState</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AdditionalLocation_NatGatewayState">
AdditionalLocation_NatGatewayState
</a>
</em>
</td>
<td>
<p>NatGatewayState: Property can be used to enable NAT Gateway for this API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>publicIpAddressId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceSkuProperties_ARM">
ApiManagementServiceSkuProperties_ARM
</a>
</em>
</td>
<td>
<p>Sku: SKU properties of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkConfiguration</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.VirtualNetworkConfiguration_ARM">
VirtualNetworkConfiguration_ARM
</a>
</em>
</td>
<td>
<p>VirtualNetworkConfiguration: Virtual network configuration for the location.</p>
</td>
</tr>
<tr>
<td>
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: A list of availability zones denoting where the resource needs to come from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AdditionalLocation_NatGatewayState">AdditionalLocation_NatGatewayState
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AdditionalLocation">AdditionalLocation</a>, <a href="#apimanagement.azure.com/v1api20230501preview.AdditionalLocation_ARM">AdditionalLocation_ARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AdditionalLocation_NatGatewayState_STATUS">AdditionalLocation_NatGatewayState_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AdditionalLocation_STATUS">AdditionalLocation_STATUS</a>, <a href="#apimanagement.azure.com/v1api20230501preview.AdditionalLocation_STATUS_ARM">AdditionalLocation_STATUS_ARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AdditionalLocation_PlatformVersion_STATUS">AdditionalLocation_PlatformVersion_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AdditionalLocation_STATUS">AdditionalLocation_STATUS</a>, <a href="#apimanagement.azure.com/v1api20230501preview.AdditionalLocation_STATUS_ARM">AdditionalLocation_STATUS_ARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;mtv1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;stv1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;stv2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;stv2.1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;undetermined&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AdditionalLocation_STATUS">AdditionalLocation_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_STATUS">Service_STATUS</a>)
</p>
<div>
<p>Description of an additional API Management resource location.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>disableGateway</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableGateway: Property only valid for an Api Management service deployed in multiple locations. This can be used to
disable the gateway in this additional location.</p>
</td>
</tr>
<tr>
<td>
<code>gatewayRegionalUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>GatewayRegionalUrl: Gateway URL of the API Management service in the Region.</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: The location name of the additional region among Azure Data center regions.</p>
</td>
</tr>
<tr>
<td>
<code>natGatewayState</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AdditionalLocation_NatGatewayState_STATUS">
AdditionalLocation_NatGatewayState_STATUS
</a>
</em>
</td>
<td>
<p>NatGatewayState: Property can be used to enable NAT Gateway for this API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>outboundPublicIPAddresses</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>OutboundPublicIPAddresses: Outbound public IPV4 address prefixes associated with NAT Gateway deployed service. Available
only for Premium SKU on stv2 platform.</p>
</td>
</tr>
<tr>
<td>
<code>platformVersion</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AdditionalLocation_PlatformVersion_STATUS">
AdditionalLocation_PlatformVersion_STATUS
</a>
</em>
</td>
<td>
<p>PlatformVersion: Compute Platform Version running the service.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAddresses</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>PrivateIPAddresses: Private Static Load Balanced IP addresses of the API Management service which is deployed in an
Internal Virtual Network in a particular additional location. Available only for Basic, Standard, Premium and Isolated
SKU.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddresses</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>PublicIPAddresses: Public Static Load Balanced IP addresses of the API Management service in the additional location.
Available only for Basic, Standard, Premium and Isolated SKU.</p>
</td>
</tr>
<tr>
<td>
<code>publicIpAddressId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PublicIpAddressId: Public Standard SKU IP V4 based IP address to be associated with Virtual Network deployed service in
the location. Supported only for Premium SKU being deployed in Virtual Network.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceSkuProperties_STATUS">
ApiManagementServiceSkuProperties_STATUS
</a>
</em>
</td>
<td>
<p>Sku: SKU properties of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkConfiguration</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.VirtualNetworkConfiguration_STATUS">
VirtualNetworkConfiguration_STATUS
</a>
</em>
</td>
<td>
<p>VirtualNetworkConfiguration: Virtual network configuration for the location.</p>
</td>
</tr>
<tr>
<td>
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: A list of availability zones denoting where the resource needs to come from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AdditionalLocation_STATUS_ARM">AdditionalLocation_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_STATUS_ARM">ApiManagementServiceProperties_STATUS_ARM</a>)
</p>
<div>
<p>Description of an additional API Management resource location.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>disableGateway</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableGateway: Property only valid for an Api Management service deployed in multiple locations. This can be used to
disable the gateway in this additional location.</p>
</td>
</tr>
<tr>
<td>
<code>gatewayRegionalUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>GatewayRegionalUrl: Gateway URL of the API Management service in the Region.</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: The location name of the additional region among Azure Data center regions.</p>
</td>
</tr>
<tr>
<td>
<code>natGatewayState</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AdditionalLocation_NatGatewayState_STATUS">
AdditionalLocation_NatGatewayState_STATUS
</a>
</em>
</td>
<td>
<p>NatGatewayState: Property can be used to enable NAT Gateway for this API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>outboundPublicIPAddresses</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>OutboundPublicIPAddresses: Outbound public IPV4 address prefixes associated with NAT Gateway deployed service. Available
only for Premium SKU on stv2 platform.</p>
</td>
</tr>
<tr>
<td>
<code>platformVersion</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AdditionalLocation_PlatformVersion_STATUS">
AdditionalLocation_PlatformVersion_STATUS
</a>
</em>
</td>
<td>
<p>PlatformVersion: Compute Platform Version running the service.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAddresses</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>PrivateIPAddresses: Private Static Load Balanced IP addresses of the API Management service which is deployed in an
Internal Virtual Network in a particular additional location. Available only for Basic, Standard, Premium and Isolated
SKU.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddresses</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>PublicIPAddresses: Public Static Load Balanced IP addresses of the API Management service in the additional location.
Available only for Basic, Standard, Premium and Isolated SKU.</p>
</td>
</tr>
<tr>
<td>
<code>publicIpAddressId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PublicIpAddressId: Public Standard SKU IP V4 based IP address to be associated with Virtual Network deployed service in
the location. Supported only for Premium SKU being deployed in Virtual Network.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceSkuProperties_STATUS_ARM">
ApiManagementServiceSkuProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Sku: SKU properties of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkConfiguration</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.VirtualNetworkConfiguration_STATUS_ARM">
VirtualNetworkConfiguration_STATUS_ARM
</a>
</em>
</td>
<td>
<p>VirtualNetworkConfiguration: Virtual network configuration for the location.</p>
</td>
</tr>
<tr>
<td>
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: A list of availability zones denoting where the resource needs to come from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Api">Api
</h3>
<div>
<p>Generator information:
- Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimapis.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ApiManagement/&#x200b;service/&#x200b;{serviceName}/&#x200b;apis/&#x200b;{apiId}</&#x200b;p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_Api_Spec">
Service_Api_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>apiVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>APIVersion: Indicates the version identifier of the API if the API is versioned</p>
</td>
</tr>
<tr>
<td>
<code>apiRevision</code><br/>
<em>
string
</em>
</td>
<td>
<p>ApiRevision: Describes the revision of the API. If no value is provided, default revision 1 is created</p>
</td>
</tr>
<tr>
<td>
<code>apiRevisionDescription</code><br/>
<em>
string
</em>
</td>
<td>
<p>ApiRevisionDescription: Description of the API Revision.</p>
</td>
</tr>
<tr>
<td>
<code>apiType</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_ApiType">
ApiCreateOrUpdateProperties_ApiType
</a>
</em>
</td>
<td>
<p>ApiType: Type of API to create.
* <code>http</code> creates a REST API
* <code>soap</code> creates a SOAP pass-through API
* <code>websocket</code> creates websocket API
* <code>graphql</code> creates GraphQL API.
New types can be added in the future.</p>
</td>
</tr>
<tr>
<td>
<code>apiVersionDescription</code><br/>
<em>
string
</em>
</td>
<td>
<p>ApiVersionDescription: Description of the API Version.</p>
</td>
</tr>
<tr>
<td>
<code>apiVersionSet</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractDetails">
ApiVersionSetContractDetails
</a>
</em>
</td>
<td>
<p>ApiVersionSet: Version set details</p>
</td>
</tr>
<tr>
<td>
<code>apiVersionSetReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>ApiVersionSetReference: A resource identifier for the related ApiVersionSet.</p>
</td>
</tr>
<tr>
<td>
<code>authenticationSettings</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthenticationSettingsContract">
AuthenticationSettingsContract
</a>
</em>
</td>
<td>
<p>AuthenticationSettings: Collection of authentication settings included into this API.</p>
</td>
</tr>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>contact</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiContactInformation">
ApiContactInformation
</a>
</em>
</td>
<td>
<p>Contact: Contact information for the API.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Description of the API. May include HTML formatting tags.</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: API name. Must be 1 to 300 characters long.</p>
</td>
</tr>
<tr>
<td>
<code>format</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_Format">
ApiCreateOrUpdateProperties_Format
</a>
</em>
</td>
<td>
<p>Format: Format of the Content in which the API is getting imported. New formats can be added in the future</p>
</td>
</tr>
<tr>
<td>
<code>isCurrent</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsCurrent: Indicates if API revision is current api revision.</p>
</td>
</tr>
<tr>
<td>
<code>license</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiLicenseInformation">
ApiLicenseInformation
</a>
</em>
</td>
<td>
<p>License: License information for the API.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a apimanagement.azure.com/Service resource</p>
</td>
</tr>
<tr>
<td>
<code>path</code><br/>
<em>
string
</em>
</td>
<td>
<p>Path: Relative URL uniquely identifying this API and all of its resource paths within the API Management service
instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public
URL for this API.</p>
</td>
</tr>
<tr>
<td>
<code>protocols</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_Protocols">
[]ApiCreateOrUpdateProperties_Protocols
</a>
</em>
</td>
<td>
<p>Protocols: Describes on which protocols the operations in this API can be invoked.</p>
</td>
</tr>
<tr>
<td>
<code>serviceUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServiceUrl: Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.</p>
</td>
</tr>
<tr>
<td>
<code>sourceApiReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>SourceApiReference: API identifier of the source API.</p>
</td>
</tr>
<tr>
<td>
<code>subscriptionKeyParameterNames</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.SubscriptionKeyParameterNamesContract">
SubscriptionKeyParameterNamesContract
</a>
</em>
</td>
<td>
<p>SubscriptionKeyParameterNames: Protocols over which API is made available.</p>
</td>
</tr>
<tr>
<td>
<code>subscriptionRequired</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SubscriptionRequired: Specifies whether an API or Product subscription is required for accessing the API.</p>
</td>
</tr>
<tr>
<td>
<code>termsOfServiceUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>TermsOfServiceUrl:  A URL to the Terms of Service for the API. MUST be in the format of a URL.</p>
</td>
</tr>
<tr>
<td>
<code>translateRequiredQueryParameters</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_TranslateRequiredQueryParameters">
ApiCreateOrUpdateProperties_TranslateRequiredQueryParameters
</a>
</em>
</td>
<td>
<p>TranslateRequiredQueryParameters: Strategy of translating required query parameters to template ones. By default has
value &lsquo;template&rsquo;. Possible values: &lsquo;template&rsquo;, &lsquo;query&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_Type">
ApiCreateOrUpdateProperties_Type
</a>
</em>
</td>
<td>
<p>Type: Type of API.</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Content value when Importing an API.</p>
</td>
</tr>
<tr>
<td>
<code>wsdlSelector</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_WsdlSelector">
ApiCreateOrUpdateProperties_WsdlSelector
</a>
</em>
</td>
<td>
<p>WsdlSelector: Criteria to limit import of WSDL to a subset of the document.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_Api_STATUS">
Service_Api_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiContactInformation">ApiContactInformation
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Api_Spec">Service_Api_Spec</a>)
</p>
<div>
<p>API contact information</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>email</code><br/>
<em>
string
</em>
</td>
<td>
<p>Email: The email address of the contact person/organization. MUST be in the format of an email address</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The identifying name of the contact person/organization</p>
</td>
</tr>
<tr>
<td>
<code>url</code><br/>
<em>
string
</em>
</td>
<td>
<p>Url: The URL pointing to the contact information. MUST be in the format of a URL</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiContactInformation_ARM">ApiContactInformation_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_ARM">ApiCreateOrUpdateProperties_ARM</a>)
</p>
<div>
<p>API contact information</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>email</code><br/>
<em>
string
</em>
</td>
<td>
<p>Email: The email address of the contact person/organization. MUST be in the format of an email address</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The identifying name of the contact person/organization</p>
</td>
</tr>
<tr>
<td>
<code>url</code><br/>
<em>
string
</em>
</td>
<td>
<p>Url: The URL pointing to the contact information. MUST be in the format of a URL</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiContactInformation_STATUS">ApiContactInformation_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Api_STATUS">Service_Api_STATUS</a>)
</p>
<div>
<p>API contact information</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>email</code><br/>
<em>
string
</em>
</td>
<td>
<p>Email: The email address of the contact person/organization. MUST be in the format of an email address</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The identifying name of the contact person/organization</p>
</td>
</tr>
<tr>
<td>
<code>url</code><br/>
<em>
string
</em>
</td>
<td>
<p>Url: The URL pointing to the contact information. MUST be in the format of a URL</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiContactInformation_STATUS_ARM">ApiContactInformation_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiContractProperties_STATUS_ARM">ApiContractProperties_STATUS_ARM</a>)
</p>
<div>
<p>API contact information</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>email</code><br/>
<em>
string
</em>
</td>
<td>
<p>Email: The email address of the contact person/organization. MUST be in the format of an email address</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The identifying name of the contact person/organization</p>
</td>
</tr>
<tr>
<td>
<code>url</code><br/>
<em>
string
</em>
</td>
<td>
<p>Url: The URL pointing to the contact information. MUST be in the format of a URL</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiContractProperties_Protocols_STATUS">ApiContractProperties_Protocols_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiContractProperties_STATUS_ARM">ApiContractProperties_STATUS_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Api_STATUS">Service_Api_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;http&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;https&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ws&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;wss&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiContractProperties_STATUS_ARM">ApiContractProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Api_STATUS_ARM">Service_Api_STATUS_ARM</a>)
</p>
<div>
<p>API Entity Properties</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>APIVersion: Indicates the version identifier of the API if the API is versioned</p>
</td>
</tr>
<tr>
<td>
<code>apiRevision</code><br/>
<em>
string
</em>
</td>
<td>
<p>ApiRevision: Describes the revision of the API. If no value is provided, default revision 1 is created</p>
</td>
</tr>
<tr>
<td>
<code>apiRevisionDescription</code><br/>
<em>
string
</em>
</td>
<td>
<p>ApiRevisionDescription: Description of the API Revision.</p>
</td>
</tr>
<tr>
<td>
<code>apiVersionDescription</code><br/>
<em>
string
</em>
</td>
<td>
<p>ApiVersionDescription: Description of the API Version.</p>
</td>
</tr>
<tr>
<td>
<code>apiVersionSet</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractDetails_STATUS_ARM">
ApiVersionSetContractDetails_STATUS_ARM
</a>
</em>
</td>
<td>
<p>ApiVersionSet: Version set details</p>
</td>
</tr>
<tr>
<td>
<code>apiVersionSetId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ApiVersionSetId: A resource identifier for the related ApiVersionSet.</p>
</td>
</tr>
<tr>
<td>
<code>authenticationSettings</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthenticationSettingsContract_STATUS_ARM">
AuthenticationSettingsContract_STATUS_ARM
</a>
</em>
</td>
<td>
<p>AuthenticationSettings: Collection of authentication settings included into this API.</p>
</td>
</tr>
<tr>
<td>
<code>contact</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiContactInformation_STATUS_ARM">
ApiContactInformation_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Contact: Contact information for the API.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Description of the API. May include HTML formatting tags.</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: API name. Must be 1 to 300 characters long.</p>
</td>
</tr>
<tr>
<td>
<code>isCurrent</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsCurrent: Indicates if API revision is current api revision.</p>
</td>
</tr>
<tr>
<td>
<code>isOnline</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsOnline: Indicates if API revision is accessible via the gateway.</p>
</td>
</tr>
<tr>
<td>
<code>license</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiLicenseInformation_STATUS_ARM">
ApiLicenseInformation_STATUS_ARM
</a>
</em>
</td>
<td>
<p>License: License information for the API.</p>
</td>
</tr>
<tr>
<td>
<code>path</code><br/>
<em>
string
</em>
</td>
<td>
<p>Path: Relative URL uniquely identifying this API and all of its resource paths within the API Management service
instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public
URL for this API.</p>
</td>
</tr>
<tr>
<td>
<code>protocols</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiContractProperties_Protocols_STATUS">
[]ApiContractProperties_Protocols_STATUS
</a>
</em>
</td>
<td>
<p>Protocols: Describes on which protocols the operations in this API can be invoked.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
string
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state</p>
</td>
</tr>
<tr>
<td>
<code>serviceUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServiceUrl: Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.</p>
</td>
</tr>
<tr>
<td>
<code>sourceApiId</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourceApiId: API identifier of the source API.</p>
</td>
</tr>
<tr>
<td>
<code>subscriptionKeyParameterNames</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.SubscriptionKeyParameterNamesContract_STATUS_ARM">
SubscriptionKeyParameterNamesContract_STATUS_ARM
</a>
</em>
</td>
<td>
<p>SubscriptionKeyParameterNames: Protocols over which API is made available.</p>
</td>
</tr>
<tr>
<td>
<code>subscriptionRequired</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SubscriptionRequired: Specifies whether an API or Product subscription is required for accessing the API.</p>
</td>
</tr>
<tr>
<td>
<code>termsOfServiceUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>TermsOfServiceUrl:  A URL to the Terms of Service for the API. MUST be in the format of a URL.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiContractProperties_Type_STATUS">
ApiContractProperties_Type_STATUS
</a>
</em>
</td>
<td>
<p>Type: Type of API.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiContractProperties_Type_STATUS">ApiContractProperties_Type_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiContractProperties_STATUS_ARM">ApiContractProperties_STATUS_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Api_STATUS">Service_Api_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;graphql&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;grpc&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;http&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;odata&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;soap&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;websocket&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_ARM">ApiCreateOrUpdateProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Api_Spec_ARM">Service_Api_Spec_ARM</a>)
</p>
<div>
<p>API Create or Update Properties.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>APIVersion: Indicates the version identifier of the API if the API is versioned</p>
</td>
</tr>
<tr>
<td>
<code>apiRevision</code><br/>
<em>
string
</em>
</td>
<td>
<p>ApiRevision: Describes the revision of the API. If no value is provided, default revision 1 is created</p>
</td>
</tr>
<tr>
<td>
<code>apiRevisionDescription</code><br/>
<em>
string
</em>
</td>
<td>
<p>ApiRevisionDescription: Description of the API Revision.</p>
</td>
</tr>
<tr>
<td>
<code>apiType</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_ApiType">
ApiCreateOrUpdateProperties_ApiType
</a>
</em>
</td>
<td>
<p>ApiType: Type of API to create.
* <code>http</code> creates a REST API
* <code>soap</code> creates a SOAP pass-through API
* <code>websocket</code> creates websocket API
* <code>graphql</code> creates GraphQL API.
New types can be added in the future.</p>
</td>
</tr>
<tr>
<td>
<code>apiVersionDescription</code><br/>
<em>
string
</em>
</td>
<td>
<p>ApiVersionDescription: Description of the API Version.</p>
</td>
</tr>
<tr>
<td>
<code>apiVersionSet</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractDetails_ARM">
ApiVersionSetContractDetails_ARM
</a>
</em>
</td>
<td>
<p>ApiVersionSet: Version set details</p>
</td>
</tr>
<tr>
<td>
<code>apiVersionSetId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>authenticationSettings</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthenticationSettingsContract_ARM">
AuthenticationSettingsContract_ARM
</a>
</em>
</td>
<td>
<p>AuthenticationSettings: Collection of authentication settings included into this API.</p>
</td>
</tr>
<tr>
<td>
<code>contact</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiContactInformation_ARM">
ApiContactInformation_ARM
</a>
</em>
</td>
<td>
<p>Contact: Contact information for the API.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Description of the API. May include HTML formatting tags.</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: API name. Must be 1 to 300 characters long.</p>
</td>
</tr>
<tr>
<td>
<code>format</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_Format">
ApiCreateOrUpdateProperties_Format
</a>
</em>
</td>
<td>
<p>Format: Format of the Content in which the API is getting imported. New formats can be added in the future</p>
</td>
</tr>
<tr>
<td>
<code>isCurrent</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsCurrent: Indicates if API revision is current api revision.</p>
</td>
</tr>
<tr>
<td>
<code>license</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiLicenseInformation_ARM">
ApiLicenseInformation_ARM
</a>
</em>
</td>
<td>
<p>License: License information for the API.</p>
</td>
</tr>
<tr>
<td>
<code>path</code><br/>
<em>
string
</em>
</td>
<td>
<p>Path: Relative URL uniquely identifying this API and all of its resource paths within the API Management service
instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public
URL for this API.</p>
</td>
</tr>
<tr>
<td>
<code>protocols</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_Protocols">
[]ApiCreateOrUpdateProperties_Protocols
</a>
</em>
</td>
<td>
<p>Protocols: Describes on which protocols the operations in this API can be invoked.</p>
</td>
</tr>
<tr>
<td>
<code>serviceUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServiceUrl: Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.</p>
</td>
</tr>
<tr>
<td>
<code>sourceApiId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>subscriptionKeyParameterNames</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.SubscriptionKeyParameterNamesContract_ARM">
SubscriptionKeyParameterNamesContract_ARM
</a>
</em>
</td>
<td>
<p>SubscriptionKeyParameterNames: Protocols over which API is made available.</p>
</td>
</tr>
<tr>
<td>
<code>subscriptionRequired</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SubscriptionRequired: Specifies whether an API or Product subscription is required for accessing the API.</p>
</td>
</tr>
<tr>
<td>
<code>termsOfServiceUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>TermsOfServiceUrl:  A URL to the Terms of Service for the API. MUST be in the format of a URL.</p>
</td>
</tr>
<tr>
<td>
<code>translateRequiredQueryParameters</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_TranslateRequiredQueryParameters">
ApiCreateOrUpdateProperties_TranslateRequiredQueryParameters
</a>
</em>
</td>
<td>
<p>TranslateRequiredQueryParameters: Strategy of translating required query parameters to template ones. By default has
value &lsquo;template&rsquo;. Possible values: &lsquo;template&rsquo;, &lsquo;query&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_Type">
ApiCreateOrUpdateProperties_Type
</a>
</em>
</td>
<td>
<p>Type: Type of API.</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Content value when Importing an API.</p>
</td>
</tr>
<tr>
<td>
<code>wsdlSelector</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_WsdlSelector_ARM">
ApiCreateOrUpdateProperties_WsdlSelector_ARM
</a>
</em>
</td>
<td>
<p>WsdlSelector: Criteria to limit import of WSDL to a subset of the document.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_ApiType">ApiCreateOrUpdateProperties_ApiType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_ARM">ApiCreateOrUpdateProperties_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Api_Spec">Service_Api_Spec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;graphql&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;grpc&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;http&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;odata&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;soap&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;websocket&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_Format">ApiCreateOrUpdateProperties_Format
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_ARM">ApiCreateOrUpdateProperties_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Api_Spec">Service_Api_Spec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;graphql-link&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;grpc&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;grpc-link&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;odata&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;odata-link&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;openapi&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;openapi&#43;json&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;openapi&#43;json-link&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;openapi-link&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;swagger-json&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;swagger-link-json&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;wadl-link-json&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;wadl-xml&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;wsdl&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;wsdl-link&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_Protocols">ApiCreateOrUpdateProperties_Protocols
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_ARM">ApiCreateOrUpdateProperties_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Api_Spec">Service_Api_Spec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;http&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;https&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ws&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;wss&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_TranslateRequiredQueryParameters">ApiCreateOrUpdateProperties_TranslateRequiredQueryParameters
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_ARM">ApiCreateOrUpdateProperties_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Api_Spec">Service_Api_Spec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;query&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;template&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_Type">ApiCreateOrUpdateProperties_Type
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_ARM">ApiCreateOrUpdateProperties_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Api_Spec">Service_Api_Spec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;graphql&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;grpc&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;http&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;odata&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;soap&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;websocket&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_WsdlSelector">ApiCreateOrUpdateProperties_WsdlSelector
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Api_Spec">Service_Api_Spec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>wsdlEndpointName</code><br/>
<em>
string
</em>
</td>
<td>
<p>WsdlEndpointName: Name of endpoint(port) to import from WSDL</p>
</td>
</tr>
<tr>
<td>
<code>wsdlServiceName</code><br/>
<em>
string
</em>
</td>
<td>
<p>WsdlServiceName: Name of service to import from WSDL</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_WsdlSelector_ARM">ApiCreateOrUpdateProperties_WsdlSelector_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_ARM">ApiCreateOrUpdateProperties_ARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>wsdlEndpointName</code><br/>
<em>
string
</em>
</td>
<td>
<p>WsdlEndpointName: Name of endpoint(port) to import from WSDL</p>
</td>
</tr>
<tr>
<td>
<code>wsdlServiceName</code><br/>
<em>
string
</em>
</td>
<td>
<p>WsdlServiceName: Name of service to import from WSDL</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiLicenseInformation">ApiLicenseInformation
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Api_Spec">Service_Api_Spec</a>)
</p>
<div>
<p>API license information</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The license name used for the API</p>
</td>
</tr>
<tr>
<td>
<code>url</code><br/>
<em>
string
</em>
</td>
<td>
<p>Url: A URL to the license used for the API. MUST be in the format of a URL</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiLicenseInformation_ARM">ApiLicenseInformation_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_ARM">ApiCreateOrUpdateProperties_ARM</a>)
</p>
<div>
<p>API license information</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The license name used for the API</p>
</td>
</tr>
<tr>
<td>
<code>url</code><br/>
<em>
string
</em>
</td>
<td>
<p>Url: A URL to the license used for the API. MUST be in the format of a URL</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiLicenseInformation_STATUS">ApiLicenseInformation_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Api_STATUS">Service_Api_STATUS</a>)
</p>
<div>
<p>API license information</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The license name used for the API</p>
</td>
</tr>
<tr>
<td>
<code>url</code><br/>
<em>
string
</em>
</td>
<td>
<p>Url: A URL to the license used for the API. MUST be in the format of a URL</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiLicenseInformation_STATUS_ARM">ApiLicenseInformation_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiContractProperties_STATUS_ARM">ApiContractProperties_STATUS_ARM</a>)
</p>
<div>
<p>API license information</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The license name used for the API</p>
</td>
</tr>
<tr>
<td>
<code>url</code><br/>
<em>
string
</em>
</td>
<td>
<p>Url: A URL to the license used for the API. MUST be in the format of a URL</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiManagementServiceIdentity">ApiManagementServiceIdentity
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Spec">Service_Spec</a>)
</p>
<div>
<p>Identity properties of the Api Management service resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceIdentity_Type">
ApiManagementServiceIdentity_Type
</a>
</em>
</td>
<td>
<p>Type: The type of identity used for the resource. The type &lsquo;SystemAssigned, UserAssigned&rsquo; includes both an implicitly
created identity and a set of user assigned identities. The type &lsquo;None&rsquo; will remove any identities from the service.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.UserAssignedIdentityDetails">
[]UserAssignedIdentityDetails
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: The list of user identities associated with the resource. The user identity
dictionary key references will be ARM resource ids in the form:
&lsquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;
providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}&rsquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiManagementServiceIdentity_ARM">ApiManagementServiceIdentity_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Spec_ARM">Service_Spec_ARM</a>)
</p>
<div>
<p>Identity properties of the Api Management service resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceIdentity_Type">
ApiManagementServiceIdentity_Type
</a>
</em>
</td>
<td>
<p>Type: The type of identity used for the resource. The type &lsquo;SystemAssigned, UserAssigned&rsquo; includes both an implicitly
created identity and a set of user assigned identities. The type &lsquo;None&rsquo; will remove any identities from the service.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.UserAssignedIdentityDetails_ARM">
map[string]./api/apimanagement/v1api20230501preview.UserAssignedIdentityDetails_ARM
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiManagementServiceIdentity_STATUS">ApiManagementServiceIdentity_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_STATUS">Service_STATUS</a>)
</p>
<div>
<p>Identity properties of the Api Management service resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>principalId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrincipalId: The principal id of the identity.</p>
</td>
</tr>
<tr>
<td>
<code>tenantId</code><br/>
<em>
string
</em>
</td>
<td>
<p>TenantId: The client tenant id of the identity.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceIdentity_Type_STATUS">
ApiManagementServiceIdentity_Type_STATUS
</a>
</em>
</td>
<td>
<p>Type: The type of identity used for the resource. The type &lsquo;SystemAssigned, UserAssigned&rsquo; includes both an implicitly
created identity and a set of user assigned identities. The type &lsquo;None&rsquo; will remove any identities from the service.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.UserIdentityProperties_STATUS">
map[string]./api/apimanagement/v1api20230501preview.UserIdentityProperties_STATUS
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: The list of user identities associated with the resource. The user identity
dictionary key references will be ARM resource ids in the form:
&lsquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;
providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}&rsquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiManagementServiceIdentity_STATUS_ARM">ApiManagementServiceIdentity_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_STATUS_ARM">Service_STATUS_ARM</a>)
</p>
<div>
<p>Identity properties of the Api Management service resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>principalId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrincipalId: The principal id of the identity.</p>
</td>
</tr>
<tr>
<td>
<code>tenantId</code><br/>
<em>
string
</em>
</td>
<td>
<p>TenantId: The client tenant id of the identity.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceIdentity_Type_STATUS">
ApiManagementServiceIdentity_Type_STATUS
</a>
</em>
</td>
<td>
<p>Type: The type of identity used for the resource. The type &lsquo;SystemAssigned, UserAssigned&rsquo; includes both an implicitly
created identity and a set of user assigned identities. The type &lsquo;None&rsquo; will remove any identities from the service.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.UserIdentityProperties_STATUS_ARM">
map[string]./api/apimanagement/v1api20230501preview.UserIdentityProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: The list of user identities associated with the resource. The user identity
dictionary key references will be ARM resource ids in the form:
&lsquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;
providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}&rsquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiManagementServiceIdentity_Type">ApiManagementServiceIdentity_Type
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceIdentity">ApiManagementServiceIdentity</a>, <a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceIdentity_ARM">ApiManagementServiceIdentity_ARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SystemAssigned&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SystemAssigned, UserAssigned&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UserAssigned&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiManagementServiceIdentity_Type_STATUS">ApiManagementServiceIdentity_Type_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceIdentity_STATUS">ApiManagementServiceIdentity_STATUS</a>, <a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceIdentity_STATUS_ARM">ApiManagementServiceIdentity_STATUS_ARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SystemAssigned&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SystemAssigned, UserAssigned&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UserAssigned&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_ARM">ApiManagementServiceProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Spec_ARM">Service_Spec_ARM</a>)
</p>
<div>
<p>Properties of an API Management service resource description.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>additionalLocations</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AdditionalLocation_ARM">
[]AdditionalLocation_ARM
</a>
</em>
</td>
<td>
<p>AdditionalLocations: Additional datacenter locations of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>apiVersionConstraint</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionConstraint_ARM">
ApiVersionConstraint_ARM
</a>
</em>
</td>
<td>
<p>ApiVersionConstraint: Control Plane Apis version constraint for the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>certificates</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CertificateConfiguration_ARM">
[]CertificateConfiguration_ARM
</a>
</em>
</td>
<td>
<p>Certificates: List of Certificates that need to be installed in the API Management service. Max supported certificates
that can be installed is 10.</p>
</td>
</tr>
<tr>
<td>
<code>configurationApi</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ConfigurationApi_ARM">
ConfigurationApi_ARM
</a>
</em>
</td>
<td>
<p>ConfigurationApi: Configuration API configuration of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>customProperties</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>CustomProperties: Custom properties of the API Management service.</br>Setting
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168</code> will disable the cipher
TLS_RSA_WITH_3DES_EDE_CBC_SHA for all TLS(1.0, 1.1 and 1.2).</br>Setting
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11</code> can be used to disable just TLS 1.1.</br>Setting
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10</code> can be used to disable TLS 1.0 on an API
Management service.</br>Setting <code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls11</code> can be
used to disable just TLS 1.1 for communications with backends.</br>Setting
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls10</code> can be used to disable TLS 1.0 for
communications with backends.</br>Setting <code>Microsoft.WindowsAzure.ApiManagement.Gateway.Protocols.Server.Http2</code> can be
used to enable HTTP2 protocol on an API Management service.</br>Not specifying any of these properties on PATCH
operation will reset omitted properties&rsquo; values to their defaults. For all the settings except Http2 the default value
is <code>True</code> if the service was created on or before April 1, 2018 and <code>False</code> otherwise. Http2 setting&rsquo;s default value is
<code>False</code>.</br></br>You can disable any of the following ciphers by using settings
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.[cipher_name]</code>: TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
TLS_RSA_WITH_AES_128_GCM_SHA256, TLS_RSA_WITH_AES_256_CBC_SHA256, TLS_RSA_WITH_AES_128_CBC_SHA256,
TLS_RSA_WITH_AES_256_CBC_SHA, TLS_RSA_WITH_AES_128_CBC_SHA. For example,
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TLS_RSA_WITH_AES_128_CBC_SHA256</code>:<code>false</code>. The default
value is <code>true</code> for them.</br> Note: The following ciphers can&rsquo;t be disabled since they are required by internal
platform components:
TLS_AES_256_GCM_SHA384,TLS_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384,TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384,TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256</p>
</td>
</tr>
<tr>
<td>
<code>developerPortalStatus</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_DeveloperPortalStatus">
ApiManagementServiceProperties_DeveloperPortalStatus
</a>
</em>
</td>
<td>
<p>DeveloperPortalStatus: Status of developer portal in this API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>disableGateway</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableGateway: Property only valid for an Api Management service deployed in multiple locations. This can be used to
disable the gateway in master region.</p>
</td>
</tr>
<tr>
<td>
<code>enableClientCertificate</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableClientCertificate: Property only meant to be used for Consumption SKU Service. This enforces a client certificate
to be presented on each request to the gateway. This also enables the ability to authenticate the certificate in the
policy on the gateway.</p>
</td>
</tr>
<tr>
<td>
<code>hostnameConfigurations</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_ARM">
[]HostnameConfiguration_ARM
</a>
</em>
</td>
<td>
<p>HostnameConfigurations: Custom hostname configuration of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>legacyPortalStatus</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_LegacyPortalStatus">
ApiManagementServiceProperties_LegacyPortalStatus
</a>
</em>
</td>
<td>
<p>LegacyPortalStatus: Status of legacy portal in the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>natGatewayState</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_NatGatewayState">
ApiManagementServiceProperties_NatGatewayState
</a>
</em>
</td>
<td>
<p>NatGatewayState: Property can be used to enable NAT Gateway for this API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>notificationSenderEmail</code><br/>
<em>
string
</em>
</td>
<td>
<p>NotificationSenderEmail: Email address from which the notification will be sent.</p>
</td>
</tr>
<tr>
<td>
<code>publicIpAddressId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_PublicNetworkAccess">
ApiManagementServiceProperties_PublicNetworkAccess
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Whether or not public endpoint access is allowed for this API Management service.  Value is
optional but if passed in, must be &lsquo;Enabled&rsquo; or &lsquo;Disabled&rsquo;. If &lsquo;Disabled&rsquo;, private endpoints are the exclusive access
method. Default value is &lsquo;Enabled&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>publisherEmail</code><br/>
<em>
string
</em>
</td>
<td>
<p>PublisherEmail: Publisher email.</p>
</td>
</tr>
<tr>
<td>
<code>publisherName</code><br/>
<em>
string
</em>
</td>
<td>
<p>PublisherName: Publisher name.</p>
</td>
</tr>
<tr>
<td>
<code>restore</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Restore: Undelete Api Management Service if it was previously soft-deleted. If this flag is specified and set to True
all other properties will be ignored.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkConfiguration</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.VirtualNetworkConfiguration_ARM">
VirtualNetworkConfiguration_ARM
</a>
</em>
</td>
<td>
<p>VirtualNetworkConfiguration: Virtual network configuration of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkType</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_VirtualNetworkType">
ApiManagementServiceProperties_VirtualNetworkType
</a>
</em>
</td>
<td>
<p>VirtualNetworkType: The type of VPN in which API Management service needs to be configured in. None (Default Value)
means the API Management service is not part of any Virtual Network, External means the API Management deployment is set
up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is
setup inside a Virtual Network having an Intranet Facing Endpoint only.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_DeveloperPortalStatus">ApiManagementServiceProperties_DeveloperPortalStatus
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_ARM">ApiManagementServiceProperties_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Spec">Service_Spec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_DeveloperPortalStatus_STATUS">ApiManagementServiceProperties_DeveloperPortalStatus_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_STATUS_ARM">ApiManagementServiceProperties_STATUS_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_STATUS">Service_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_LegacyPortalStatus">ApiManagementServiceProperties_LegacyPortalStatus
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_ARM">ApiManagementServiceProperties_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Spec">Service_Spec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_LegacyPortalStatus_STATUS">ApiManagementServiceProperties_LegacyPortalStatus_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_STATUS_ARM">ApiManagementServiceProperties_STATUS_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_STATUS">Service_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_NatGatewayState">ApiManagementServiceProperties_NatGatewayState
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_ARM">ApiManagementServiceProperties_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Spec">Service_Spec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_NatGatewayState_STATUS">ApiManagementServiceProperties_NatGatewayState_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_STATUS_ARM">ApiManagementServiceProperties_STATUS_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_STATUS">Service_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_PlatformVersion_STATUS">ApiManagementServiceProperties_PlatformVersion_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_STATUS_ARM">ApiManagementServiceProperties_STATUS_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_STATUS">Service_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;mtv1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;stv1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;stv2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;stv2.1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;undetermined&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_PublicNetworkAccess">ApiManagementServiceProperties_PublicNetworkAccess
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_ARM">ApiManagementServiceProperties_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Spec">Service_Spec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_PublicNetworkAccess_STATUS">ApiManagementServiceProperties_PublicNetworkAccess_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_STATUS_ARM">ApiManagementServiceProperties_STATUS_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_STATUS">Service_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_STATUS_ARM">ApiManagementServiceProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_STATUS_ARM">Service_STATUS_ARM</a>)
</p>
<div>
<p>Properties of an API Management service resource description.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>additionalLocations</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AdditionalLocation_STATUS_ARM">
[]AdditionalLocation_STATUS_ARM
</a>
</em>
</td>
<td>
<p>AdditionalLocations: Additional datacenter locations of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>apiVersionConstraint</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionConstraint_STATUS_ARM">
ApiVersionConstraint_STATUS_ARM
</a>
</em>
</td>
<td>
<p>ApiVersionConstraint: Control Plane Apis version constraint for the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>certificates</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CertificateConfiguration_STATUS_ARM">
[]CertificateConfiguration_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Certificates: List of Certificates that need to be installed in the API Management service. Max supported certificates
that can be installed is 10.</p>
</td>
</tr>
<tr>
<td>
<code>configurationApi</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ConfigurationApi_STATUS_ARM">
ConfigurationApi_STATUS_ARM
</a>
</em>
</td>
<td>
<p>ConfigurationApi: Configuration API configuration of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>createdAtUtc</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedAtUtc: Creation UTC date of the API Management service.The date conforms to the following format:
<code>yyyy-MM-ddTHH:mm:ssZ</code> as specified by the ISO 8601 standard.</p>
</td>
</tr>
<tr>
<td>
<code>customProperties</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>CustomProperties: Custom properties of the API Management service.</br>Setting
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168</code> will disable the cipher
TLS_RSA_WITH_3DES_EDE_CBC_SHA for all TLS(1.0, 1.1 and 1.2).</br>Setting
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11</code> can be used to disable just TLS 1.1.</br>Setting
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10</code> can be used to disable TLS 1.0 on an API
Management service.</br>Setting <code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls11</code> can be
used to disable just TLS 1.1 for communications with backends.</br>Setting
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls10</code> can be used to disable TLS 1.0 for
communications with backends.</br>Setting <code>Microsoft.WindowsAzure.ApiManagement.Gateway.Protocols.Server.Http2</code> can be
used to enable HTTP2 protocol on an API Management service.</br>Not specifying any of these properties on PATCH
operation will reset omitted properties&rsquo; values to their defaults. For all the settings except Http2 the default value
is <code>True</code> if the service was created on or before April 1, 2018 and <code>False</code> otherwise. Http2 setting&rsquo;s default value is
<code>False</code>.</br></br>You can disable any of the following ciphers by using settings
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.[cipher_name]</code>: TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
TLS_RSA_WITH_AES_128_GCM_SHA256, TLS_RSA_WITH_AES_256_CBC_SHA256, TLS_RSA_WITH_AES_128_CBC_SHA256,
TLS_RSA_WITH_AES_256_CBC_SHA, TLS_RSA_WITH_AES_128_CBC_SHA. For example,
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TLS_RSA_WITH_AES_128_CBC_SHA256</code>:<code>false</code>. The default
value is <code>true</code> for them.</br> Note: The following ciphers can&rsquo;t be disabled since they are required by internal
platform components:
TLS_AES_256_GCM_SHA384,TLS_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384,TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384,TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256</p>
</td>
</tr>
<tr>
<td>
<code>developerPortalStatus</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_DeveloperPortalStatus_STATUS">
ApiManagementServiceProperties_DeveloperPortalStatus_STATUS
</a>
</em>
</td>
<td>
<p>DeveloperPortalStatus: Status of developer portal in this API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>developerPortalUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>DeveloperPortalUrl: DEveloper Portal endpoint URL of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>disableGateway</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableGateway: Property only valid for an Api Management service deployed in multiple locations. This can be used to
disable the gateway in master region.</p>
</td>
</tr>
<tr>
<td>
<code>enableClientCertificate</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableClientCertificate: Property only meant to be used for Consumption SKU Service. This enforces a client certificate
to be presented on each request to the gateway. This also enables the ability to authenticate the certificate in the
policy on the gateway.</p>
</td>
</tr>
<tr>
<td>
<code>gatewayRegionalUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>GatewayRegionalUrl: Gateway URL of the API Management service in the Default Region.</p>
</td>
</tr>
<tr>
<td>
<code>gatewayUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>GatewayUrl: Gateway URL of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>hostnameConfigurations</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_STATUS_ARM">
[]HostnameConfiguration_STATUS_ARM
</a>
</em>
</td>
<td>
<p>HostnameConfigurations: Custom hostname configuration of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>legacyPortalStatus</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_LegacyPortalStatus_STATUS">
ApiManagementServiceProperties_LegacyPortalStatus_STATUS
</a>
</em>
</td>
<td>
<p>LegacyPortalStatus: Status of legacy portal in the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>managementApiUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>ManagementApiUrl: Management API endpoint URL of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>natGatewayState</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_NatGatewayState_STATUS">
ApiManagementServiceProperties_NatGatewayState_STATUS
</a>
</em>
</td>
<td>
<p>NatGatewayState: Property can be used to enable NAT Gateway for this API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>notificationSenderEmail</code><br/>
<em>
string
</em>
</td>
<td>
<p>NotificationSenderEmail: Email address from which the notification will be sent.</p>
</td>
</tr>
<tr>
<td>
<code>outboundPublicIPAddresses</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>OutboundPublicIPAddresses: Outbound public IPV4 address prefixes associated with NAT Gateway deployed service. Available
only for Premium SKU on stv2 platform.</p>
</td>
</tr>
<tr>
<td>
<code>platformVersion</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_PlatformVersion_STATUS">
ApiManagementServiceProperties_PlatformVersion_STATUS
</a>
</em>
</td>
<td>
<p>PlatformVersion: Compute Platform Version running the service in this location.</p>
</td>
</tr>
<tr>
<td>
<code>portalUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>PortalUrl: Publisher portal endpoint Url of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpointConnections</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.RemotePrivateEndpointConnectionWrapper_STATUS_ARM">
[]RemotePrivateEndpointConnectionWrapper_STATUS_ARM
</a>
</em>
</td>
<td>
<p>PrivateEndpointConnections: List of Private Endpoint Connections of this service.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAddresses</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>PrivateIPAddresses: Private Static Load Balanced IP addresses of the API Management service in Primary region which is
deployed in an Internal Virtual Network. Available only for Basic, Standard, Premium and Isolated SKU.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
string
</em>
</td>
<td>
<p>ProvisioningState: The current provisioning state of the API Management service which can be one of the following:
Created/Activating/Succeeded/Updating/Failed/Stopped/Terminating/TerminationFailed/Deleted.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddresses</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>PublicIPAddresses: Public Static Load Balanced IP addresses of the API Management service in Primary region. Available
only for Basic, Standard, Premium and Isolated SKU.</p>
</td>
</tr>
<tr>
<td>
<code>publicIpAddressId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PublicIpAddressId: Public Standard SKU IP V4 based IP address to be associated with Virtual Network deployed service in
the region. Supported only for Developer and Premium SKU being deployed in Virtual Network.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_PublicNetworkAccess_STATUS">
ApiManagementServiceProperties_PublicNetworkAccess_STATUS
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Whether or not public endpoint access is allowed for this API Management service.  Value is
optional but if passed in, must be &lsquo;Enabled&rsquo; or &lsquo;Disabled&rsquo;. If &lsquo;Disabled&rsquo;, private endpoints are the exclusive access
method. Default value is &lsquo;Enabled&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>publisherEmail</code><br/>
<em>
string
</em>
</td>
<td>
<p>PublisherEmail: Publisher email.</p>
</td>
</tr>
<tr>
<td>
<code>publisherName</code><br/>
<em>
string
</em>
</td>
<td>
<p>PublisherName: Publisher name.</p>
</td>
</tr>
<tr>
<td>
<code>restore</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Restore: Undelete Api Management Service if it was previously soft-deleted. If this flag is specified and set to True
all other properties will be ignored.</p>
</td>
</tr>
<tr>
<td>
<code>scmUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScmUrl: SCM endpoint URL of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>targetProvisioningState</code><br/>
<em>
string
</em>
</td>
<td>
<p>TargetProvisioningState: The provisioning state of the API Management service, which is targeted by the long running
operation started on the service.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkConfiguration</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.VirtualNetworkConfiguration_STATUS_ARM">
VirtualNetworkConfiguration_STATUS_ARM
</a>
</em>
</td>
<td>
<p>VirtualNetworkConfiguration: Virtual network configuration of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkType</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_VirtualNetworkType_STATUS">
ApiManagementServiceProperties_VirtualNetworkType_STATUS
</a>
</em>
</td>
<td>
<p>VirtualNetworkType: The type of VPN in which API Management service needs to be configured in. None (Default Value)
means the API Management service is not part of any Virtual Network, External means the API Management deployment is set
up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is
setup inside a Virtual Network having an Intranet Facing Endpoint only.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_VirtualNetworkType">ApiManagementServiceProperties_VirtualNetworkType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_ARM">ApiManagementServiceProperties_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Spec">Service_Spec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;External&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Internal&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_VirtualNetworkType_STATUS">ApiManagementServiceProperties_VirtualNetworkType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_STATUS_ARM">ApiManagementServiceProperties_STATUS_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_STATUS">Service_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;External&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Internal&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiManagementServiceSkuProperties">ApiManagementServiceSkuProperties
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AdditionalLocation">AdditionalLocation</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Spec">Service_Spec</a>)
</p>
<div>
<p>API Management service resource SKU properties.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>capacity</code><br/>
<em>
int
</em>
</td>
<td>
<p>Capacity: Capacity of the SKU (number of deployed units of the SKU). For Consumption SKU capacity must be specified as 0.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceSkuProperties_Name">
ApiManagementServiceSkuProperties_Name
</a>
</em>
</td>
<td>
<p>Name: Name of the Sku.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiManagementServiceSkuProperties_ARM">ApiManagementServiceSkuProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AdditionalLocation_ARM">AdditionalLocation_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Spec_ARM">Service_Spec_ARM</a>)
</p>
<div>
<p>API Management service resource SKU properties.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>capacity</code><br/>
<em>
int
</em>
</td>
<td>
<p>Capacity: Capacity of the SKU (number of deployed units of the SKU). For Consumption SKU capacity must be specified as 0.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceSkuProperties_Name">
ApiManagementServiceSkuProperties_Name
</a>
</em>
</td>
<td>
<p>Name: Name of the Sku.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiManagementServiceSkuProperties_Name">ApiManagementServiceSkuProperties_Name
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceSkuProperties">ApiManagementServiceSkuProperties</a>, <a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceSkuProperties_ARM">ApiManagementServiceSkuProperties_ARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Basic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;BasicV2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Consumption&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Developer&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Isolated&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Premium&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;StandardV2&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiManagementServiceSkuProperties_Name_STATUS">ApiManagementServiceSkuProperties_Name_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceSkuProperties_STATUS">ApiManagementServiceSkuProperties_STATUS</a>, <a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceSkuProperties_STATUS_ARM">ApiManagementServiceSkuProperties_STATUS_ARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Basic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;BasicV2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Consumption&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Developer&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Isolated&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Premium&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;StandardV2&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiManagementServiceSkuProperties_STATUS">ApiManagementServiceSkuProperties_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AdditionalLocation_STATUS">AdditionalLocation_STATUS</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_STATUS">Service_STATUS</a>)
</p>
<div>
<p>API Management service resource SKU properties.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>capacity</code><br/>
<em>
int
</em>
</td>
<td>
<p>Capacity: Capacity of the SKU (number of deployed units of the SKU). For Consumption SKU capacity must be specified as 0.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceSkuProperties_Name_STATUS">
ApiManagementServiceSkuProperties_Name_STATUS
</a>
</em>
</td>
<td>
<p>Name: Name of the Sku.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiManagementServiceSkuProperties_STATUS_ARM">ApiManagementServiceSkuProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AdditionalLocation_STATUS_ARM">AdditionalLocation_STATUS_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_STATUS_ARM">Service_STATUS_ARM</a>)
</p>
<div>
<p>API Management service resource SKU properties.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>capacity</code><br/>
<em>
int
</em>
</td>
<td>
<p>Capacity: Capacity of the SKU (number of deployed units of the SKU). For Consumption SKU capacity must be specified as 0.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceSkuProperties_Name_STATUS">
ApiManagementServiceSkuProperties_Name_STATUS
</a>
</em>
</td>
<td>
<p>Name: Name of the Sku.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiVersionConstraint">ApiVersionConstraint
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Spec">Service_Spec</a>)
</p>
<div>
<p>Control Plane Apis version constraint for the API Management service.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>minApiVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>MinApiVersion: Limit control plane API calls to API Management service with version equal to or newer than this value.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiVersionConstraint_ARM">ApiVersionConstraint_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_ARM">ApiManagementServiceProperties_ARM</a>)
</p>
<div>
<p>Control Plane Apis version constraint for the API Management service.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>minApiVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>MinApiVersion: Limit control plane API calls to API Management service with version equal to or newer than this value.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiVersionConstraint_STATUS">ApiVersionConstraint_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_STATUS">Service_STATUS</a>)
</p>
<div>
<p>Control Plane Apis version constraint for the API Management service.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>minApiVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>MinApiVersion: Limit control plane API calls to API Management service with version equal to or newer than this value.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiVersionConstraint_STATUS_ARM">ApiVersionConstraint_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_STATUS_ARM">ApiManagementServiceProperties_STATUS_ARM</a>)
</p>
<div>
<p>Control Plane Apis version constraint for the API Management service.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>minApiVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>MinApiVersion: Limit control plane API calls to API Management service with version equal to or newer than this value.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiVersionSet">ApiVersionSet
</h3>
<div>
<p>Generator information:
- Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimapiversionsets.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ApiManagement/&#x200b;service/&#x200b;{serviceName}/&#x200b;apiVersionSets/&#x200b;{versionSetId}</&#x200b;p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_ApiVersionSet_Spec">
Service_ApiVersionSet_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Description of API Version Set.</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: Name of API Version Set</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a apimanagement.azure.com/Service resource</p>
</td>
</tr>
<tr>
<td>
<code>versionHeaderName</code><br/>
<em>
string
</em>
</td>
<td>
<p>VersionHeaderName: Name of HTTP header parameter that indicates the API Version if versioningScheme is set to <code>header</code>.</p>
</td>
</tr>
<tr>
<td>
<code>versionQueryName</code><br/>
<em>
string
</em>
</td>
<td>
<p>VersionQueryName: Name of query parameter that indicates the API Version if versioningScheme is set to <code>query</code>.</p>
</td>
</tr>
<tr>
<td>
<code>versioningScheme</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractProperties_VersioningScheme">
ApiVersionSetContractProperties_VersioningScheme
</a>
</em>
</td>
<td>
<p>VersioningScheme: An value that determines where the API Version identifier will be located in a HTTP request.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_ApiVersionSet_STATUS">
Service_ApiVersionSet_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractDetails">ApiVersionSetContractDetails
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Api_Spec">Service_Api_Spec</a>)
</p>
<div>
<p>An API Version Set contains the common configuration for a set of API Versions relating</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Description of API Version Set.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The display Name of the API Version Set.</p>
</td>
</tr>
<tr>
<td>
<code>reference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>Reference: Identifier for existing API Version Set. Omit this value to create a new Version Set.</p>
</td>
</tr>
<tr>
<td>
<code>versionHeaderName</code><br/>
<em>
string
</em>
</td>
<td>
<p>VersionHeaderName: Name of HTTP header parameter that indicates the API Version if versioningScheme is set to <code>header</code>.</p>
</td>
</tr>
<tr>
<td>
<code>versionQueryName</code><br/>
<em>
string
</em>
</td>
<td>
<p>VersionQueryName: Name of query parameter that indicates the API Version if versioningScheme is set to <code>query</code>.</p>
</td>
</tr>
<tr>
<td>
<code>versioningScheme</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractDetails_VersioningScheme">
ApiVersionSetContractDetails_VersioningScheme
</a>
</em>
</td>
<td>
<p>VersioningScheme: An value that determines where the API Version identifier will be located in a HTTP request.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractDetails_ARM">ApiVersionSetContractDetails_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_ARM">ApiCreateOrUpdateProperties_ARM</a>)
</p>
<div>
<p>An API Version Set contains the common configuration for a set of API Versions relating</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Description of API Version Set.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The display Name of the API Version Set.</p>
</td>
</tr>
<tr>
<td>
<code>versionHeaderName</code><br/>
<em>
string
</em>
</td>
<td>
<p>VersionHeaderName: Name of HTTP header parameter that indicates the API Version if versioningScheme is set to <code>header</code>.</p>
</td>
</tr>
<tr>
<td>
<code>versionQueryName</code><br/>
<em>
string
</em>
</td>
<td>
<p>VersionQueryName: Name of query parameter that indicates the API Version if versioningScheme is set to <code>query</code>.</p>
</td>
</tr>
<tr>
<td>
<code>versioningScheme</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractDetails_VersioningScheme">
ApiVersionSetContractDetails_VersioningScheme
</a>
</em>
</td>
<td>
<p>VersioningScheme: An value that determines where the API Version identifier will be located in a HTTP request.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractDetails_STATUS">ApiVersionSetContractDetails_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Api_STATUS">Service_Api_STATUS</a>)
</p>
<div>
<p>An API Version Set contains the common configuration for a set of API Versions relating</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Description of API Version Set.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Identifier for existing API Version Set. Omit this value to create a new Version Set.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The display Name of the API Version Set.</p>
</td>
</tr>
<tr>
<td>
<code>versionHeaderName</code><br/>
<em>
string
</em>
</td>
<td>
<p>VersionHeaderName: Name of HTTP header parameter that indicates the API Version if versioningScheme is set to <code>header</code>.</p>
</td>
</tr>
<tr>
<td>
<code>versionQueryName</code><br/>
<em>
string
</em>
</td>
<td>
<p>VersionQueryName: Name of query parameter that indicates the API Version if versioningScheme is set to <code>query</code>.</p>
</td>
</tr>
<tr>
<td>
<code>versioningScheme</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractDetails_VersioningScheme_STATUS">
ApiVersionSetContractDetails_VersioningScheme_STATUS
</a>
</em>
</td>
<td>
<p>VersioningScheme: An value that determines where the API Version identifier will be located in a HTTP request.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractDetails_STATUS_ARM">ApiVersionSetContractDetails_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiContractProperties_STATUS_ARM">ApiContractProperties_STATUS_ARM</a>)
</p>
<div>
<p>An API Version Set contains the common configuration for a set of API Versions relating</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Description of API Version Set.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Identifier for existing API Version Set. Omit this value to create a new Version Set.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The display Name of the API Version Set.</p>
</td>
</tr>
<tr>
<td>
<code>versionHeaderName</code><br/>
<em>
string
</em>
</td>
<td>
<p>VersionHeaderName: Name of HTTP header parameter that indicates the API Version if versioningScheme is set to <code>header</code>.</p>
</td>
</tr>
<tr>
<td>
<code>versionQueryName</code><br/>
<em>
string
</em>
</td>
<td>
<p>VersionQueryName: Name of query parameter that indicates the API Version if versioningScheme is set to <code>query</code>.</p>
</td>
</tr>
<tr>
<td>
<code>versioningScheme</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractDetails_VersioningScheme_STATUS">
ApiVersionSetContractDetails_VersioningScheme_STATUS
</a>
</em>
</td>
<td>
<p>VersioningScheme: An value that determines where the API Version identifier will be located in a HTTP request.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractDetails_VersioningScheme">ApiVersionSetContractDetails_VersioningScheme
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractDetails">ApiVersionSetContractDetails</a>, <a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractDetails_ARM">ApiVersionSetContractDetails_ARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Header&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Query&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Segment&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractDetails_VersioningScheme_STATUS">ApiVersionSetContractDetails_VersioningScheme_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractDetails_STATUS">ApiVersionSetContractDetails_STATUS</a>, <a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractDetails_STATUS_ARM">ApiVersionSetContractDetails_STATUS_ARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Header&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Query&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Segment&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractProperties_ARM">ApiVersionSetContractProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_ApiVersionSet_Spec_ARM">Service_ApiVersionSet_Spec_ARM</a>)
</p>
<div>
<p>Properties of an API Version Set.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Description of API Version Set.</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: Name of API Version Set</p>
</td>
</tr>
<tr>
<td>
<code>versionHeaderName</code><br/>
<em>
string
</em>
</td>
<td>
<p>VersionHeaderName: Name of HTTP header parameter that indicates the API Version if versioningScheme is set to <code>header</code>.</p>
</td>
</tr>
<tr>
<td>
<code>versionQueryName</code><br/>
<em>
string
</em>
</td>
<td>
<p>VersionQueryName: Name of query parameter that indicates the API Version if versioningScheme is set to <code>query</code>.</p>
</td>
</tr>
<tr>
<td>
<code>versioningScheme</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractProperties_VersioningScheme">
ApiVersionSetContractProperties_VersioningScheme
</a>
</em>
</td>
<td>
<p>VersioningScheme: An value that determines where the API Version identifier will be located in a HTTP request.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractProperties_STATUS_ARM">ApiVersionSetContractProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_ApiVersionSet_STATUS_ARM">Service_ApiVersionSet_STATUS_ARM</a>)
</p>
<div>
<p>Properties of an API Version Set.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Description of API Version Set.</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: Name of API Version Set</p>
</td>
</tr>
<tr>
<td>
<code>versionHeaderName</code><br/>
<em>
string
</em>
</td>
<td>
<p>VersionHeaderName: Name of HTTP header parameter that indicates the API Version if versioningScheme is set to <code>header</code>.</p>
</td>
</tr>
<tr>
<td>
<code>versionQueryName</code><br/>
<em>
string
</em>
</td>
<td>
<p>VersionQueryName: Name of query parameter that indicates the API Version if versioningScheme is set to <code>query</code>.</p>
</td>
</tr>
<tr>
<td>
<code>versioningScheme</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractProperties_VersioningScheme_STATUS">
ApiVersionSetContractProperties_VersioningScheme_STATUS
</a>
</em>
</td>
<td>
<p>VersioningScheme: An value that determines where the API Version identifier will be located in a HTTP request.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractProperties_VersioningScheme">ApiVersionSetContractProperties_VersioningScheme
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractProperties_ARM">ApiVersionSetContractProperties_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_ApiVersionSet_Spec">Service_ApiVersionSet_Spec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Header&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Query&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Segment&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractProperties_VersioningScheme_STATUS">ApiVersionSetContractProperties_VersioningScheme_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractProperties_STATUS_ARM">ApiVersionSetContractProperties_STATUS_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_ApiVersionSet_STATUS">Service_ApiVersionSet_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Header&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Query&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Segment&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ArmIdWrapper_STATUS">ArmIdWrapper_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.RemotePrivateEndpointConnectionWrapper_STATUS">RemotePrivateEndpointConnectionWrapper_STATUS</a>)
</p>
<div>
<p>A wrapper for an ARM resource id</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ArmIdWrapper_STATUS_ARM">ArmIdWrapper_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.PrivateEndpointConnectionWrapperProperties_STATUS_ARM">PrivateEndpointConnectionWrapperProperties_STATUS_ARM</a>)
</p>
<div>
<p>A wrapper for an ARM resource id</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthenticationSettingsContract">AuthenticationSettingsContract
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Api_Spec">Service_Api_Spec</a>)
</p>
<div>
<p>API Authentication Settings.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>oAuth2</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.OAuth2AuthenticationSettingsContract">
OAuth2AuthenticationSettingsContract
</a>
</em>
</td>
<td>
<p>OAuth2: OAuth2 Authentication settings</p>
</td>
</tr>
<tr>
<td>
<code>oAuth2AuthenticationSettings</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.OAuth2AuthenticationSettingsContract">
[]OAuth2AuthenticationSettingsContract
</a>
</em>
</td>
<td>
<p>OAuth2AuthenticationSettings: Collection of OAuth2 authentication settings included into this API.</p>
</td>
</tr>
<tr>
<td>
<code>openid</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.OpenIdAuthenticationSettingsContract">
OpenIdAuthenticationSettingsContract
</a>
</em>
</td>
<td>
<p>Openid: OpenID Connect Authentication Settings</p>
</td>
</tr>
<tr>
<td>
<code>openidAuthenticationSettings</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.OpenIdAuthenticationSettingsContract">
[]OpenIdAuthenticationSettingsContract
</a>
</em>
</td>
<td>
<p>OpenidAuthenticationSettings: Collection of Open ID Connect authentication settings included into this API.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthenticationSettingsContract_ARM">AuthenticationSettingsContract_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_ARM">ApiCreateOrUpdateProperties_ARM</a>)
</p>
<div>
<p>API Authentication Settings.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>oAuth2</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.OAuth2AuthenticationSettingsContract_ARM">
OAuth2AuthenticationSettingsContract_ARM
</a>
</em>
</td>
<td>
<p>OAuth2: OAuth2 Authentication settings</p>
</td>
</tr>
<tr>
<td>
<code>oAuth2AuthenticationSettings</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.OAuth2AuthenticationSettingsContract_ARM">
[]OAuth2AuthenticationSettingsContract_ARM
</a>
</em>
</td>
<td>
<p>OAuth2AuthenticationSettings: Collection of OAuth2 authentication settings included into this API.</p>
</td>
</tr>
<tr>
<td>
<code>openid</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.OpenIdAuthenticationSettingsContract_ARM">
OpenIdAuthenticationSettingsContract_ARM
</a>
</em>
</td>
<td>
<p>Openid: OpenID Connect Authentication Settings</p>
</td>
</tr>
<tr>
<td>
<code>openidAuthenticationSettings</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.OpenIdAuthenticationSettingsContract_ARM">
[]OpenIdAuthenticationSettingsContract_ARM
</a>
</em>
</td>
<td>
<p>OpenidAuthenticationSettings: Collection of Open ID Connect authentication settings included into this API.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthenticationSettingsContract_STATUS">AuthenticationSettingsContract_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Api_STATUS">Service_Api_STATUS</a>)
</p>
<div>
<p>API Authentication Settings.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>oAuth2</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.OAuth2AuthenticationSettingsContract_STATUS">
OAuth2AuthenticationSettingsContract_STATUS
</a>
</em>
</td>
<td>
<p>OAuth2: OAuth2 Authentication settings</p>
</td>
</tr>
<tr>
<td>
<code>oAuth2AuthenticationSettings</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.OAuth2AuthenticationSettingsContract_STATUS">
[]OAuth2AuthenticationSettingsContract_STATUS
</a>
</em>
</td>
<td>
<p>OAuth2AuthenticationSettings: Collection of OAuth2 authentication settings included into this API.</p>
</td>
</tr>
<tr>
<td>
<code>openid</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.OpenIdAuthenticationSettingsContract_STATUS">
OpenIdAuthenticationSettingsContract_STATUS
</a>
</em>
</td>
<td>
<p>Openid: OpenID Connect Authentication Settings</p>
</td>
</tr>
<tr>
<td>
<code>openidAuthenticationSettings</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.OpenIdAuthenticationSettingsContract_STATUS">
[]OpenIdAuthenticationSettingsContract_STATUS
</a>
</em>
</td>
<td>
<p>OpenidAuthenticationSettings: Collection of Open ID Connect authentication settings included into this API.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthenticationSettingsContract_STATUS_ARM">AuthenticationSettingsContract_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiContractProperties_STATUS_ARM">ApiContractProperties_STATUS_ARM</a>)
</p>
<div>
<p>API Authentication Settings.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>oAuth2</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.OAuth2AuthenticationSettingsContract_STATUS_ARM">
OAuth2AuthenticationSettingsContract_STATUS_ARM
</a>
</em>
</td>
<td>
<p>OAuth2: OAuth2 Authentication settings</p>
</td>
</tr>
<tr>
<td>
<code>oAuth2AuthenticationSettings</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.OAuth2AuthenticationSettingsContract_STATUS_ARM">
[]OAuth2AuthenticationSettingsContract_STATUS_ARM
</a>
</em>
</td>
<td>
<p>OAuth2AuthenticationSettings: Collection of OAuth2 authentication settings included into this API.</p>
</td>
</tr>
<tr>
<td>
<code>openid</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.OpenIdAuthenticationSettingsContract_STATUS_ARM">
OpenIdAuthenticationSettingsContract_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Openid: OpenID Connect Authentication Settings</p>
</td>
</tr>
<tr>
<td>
<code>openidAuthenticationSettings</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.OpenIdAuthenticationSettingsContract_STATUS_ARM">
[]OpenIdAuthenticationSettingsContract_STATUS_ARM
</a>
</em>
</td>
<td>
<p>OpenidAuthenticationSettings: Collection of Open ID Connect authentication settings included into this API.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthorizationAccessPolicyContractProperties_ARM">AuthorizationAccessPolicyContractProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec_ARM">Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec_ARM</a>)
</p>
<div>
<p>Authorization Access Policy details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>appIds</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AppIds: The allowed Azure Active Directory Application IDs</p>
</td>
</tr>
<tr>
<td>
<code>objectId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ObjectId: The Object Id</p>
</td>
</tr>
<tr>
<td>
<code>tenantId</code><br/>
<em>
string
</em>
</td>
<td>
<p>TenantId: The Tenant Id</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthorizationAccessPolicyContractProperties_STATUS_ARM">AuthorizationAccessPolicyContractProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS_ARM">Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS_ARM</a>)
</p>
<div>
<p>Authorization Access Policy details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>appIds</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AppIds: The allowed Azure Active Directory Application IDs</p>
</td>
</tr>
<tr>
<td>
<code>objectId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ObjectId: The Object Id</p>
</td>
</tr>
<tr>
<td>
<code>tenantId</code><br/>
<em>
string
</em>
</td>
<td>
<p>TenantId: The Tenant Id</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthorizationContractProperties_ARM">AuthorizationContractProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProviders_Authorization_Spec_ARM">Service_AuthorizationProviders_Authorization_Spec_ARM</a>)
</p>
<div>
<p>Authorization details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>authorizationType</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationContractProperties_AuthorizationType">
AuthorizationContractProperties_AuthorizationType
</a>
</em>
</td>
<td>
<p>AuthorizationType: Authorization type options</p>
</td>
</tr>
<tr>
<td>
<code>oauth2grantType</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationContractProperties_Oauth2GrantType">
AuthorizationContractProperties_Oauth2GrantType
</a>
</em>
</td>
<td>
<p>Oauth2GrantType: OAuth2 grant type options</p>
</td>
</tr>
<tr>
<td>
<code>parameters</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Parameters: Authorization parameters</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthorizationContractProperties_AuthorizationType">AuthorizationContractProperties_AuthorizationType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationContractProperties_ARM">AuthorizationContractProperties_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProviders_Authorization_Spec">Service_AuthorizationProviders_Authorization_Spec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;OAuth2&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthorizationContractProperties_AuthorizationType_STATUS">AuthorizationContractProperties_AuthorizationType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationContractProperties_STATUS_ARM">AuthorizationContractProperties_STATUS_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProviders_Authorization_STATUS">Service_AuthorizationProviders_Authorization_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;OAuth2&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthorizationContractProperties_Oauth2GrantType">AuthorizationContractProperties_Oauth2GrantType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationContractProperties_ARM">AuthorizationContractProperties_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProviders_Authorization_Spec">Service_AuthorizationProviders_Authorization_Spec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;AuthorizationCode&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ClientCredentials&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthorizationContractProperties_Oauth2GrantType_STATUS">AuthorizationContractProperties_Oauth2GrantType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationContractProperties_STATUS_ARM">AuthorizationContractProperties_STATUS_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProviders_Authorization_STATUS">Service_AuthorizationProviders_Authorization_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;AuthorizationCode&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ClientCredentials&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthorizationContractProperties_STATUS_ARM">AuthorizationContractProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProviders_Authorization_STATUS_ARM">Service_AuthorizationProviders_Authorization_STATUS_ARM</a>)
</p>
<div>
<p>Authorization details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>authorizationType</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationContractProperties_AuthorizationType_STATUS">
AuthorizationContractProperties_AuthorizationType_STATUS
</a>
</em>
</td>
<td>
<p>AuthorizationType: Authorization type options</p>
</td>
</tr>
<tr>
<td>
<code>error</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationError_STATUS_ARM">
AuthorizationError_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Error: Authorization error details.</p>
</td>
</tr>
<tr>
<td>
<code>oauth2grantType</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationContractProperties_Oauth2GrantType_STATUS">
AuthorizationContractProperties_Oauth2GrantType_STATUS
</a>
</em>
</td>
<td>
<p>Oauth2GrantType: OAuth2 grant type options</p>
</td>
</tr>
<tr>
<td>
<code>parameters</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Parameters: Authorization parameters</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
string
</em>
</td>
<td>
<p>Status: Status of the Authorization</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthorizationError_STATUS">AuthorizationError_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProviders_Authorization_STATUS">Service_AuthorizationProviders_Authorization_STATUS</a>)
</p>
<div>
<p>Authorization error details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>code</code><br/>
<em>
string
</em>
</td>
<td>
<p>Code: Error code</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<p>Message: Error message</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthorizationError_STATUS_ARM">AuthorizationError_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationContractProperties_STATUS_ARM">AuthorizationContractProperties_STATUS_ARM</a>)
</p>
<div>
<p>Authorization error details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>code</code><br/>
<em>
string
</em>
</td>
<td>
<p>Code: Error code</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<p>Message: Error message</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthorizationProvider">AuthorizationProvider
</h3>
<div>
<p>Generator information:
- Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimauthorizationproviders.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ApiManagement/&#x200b;service/&#x200b;{serviceName}/&#x200b;authorizationProviders/&#x200b;{authorizationProviderId}</&#x200b;p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProvider_Spec">
Service_AuthorizationProvider_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: Authorization Provider name. Must be 1 to 300 characters long.</p>
</td>
</tr>
<tr>
<td>
<code>identityProvider</code><br/>
<em>
string
</em>
</td>
<td>
<p>IdentityProvider: Identity provider name. Must be 1 to 300 characters long.</p>
</td>
</tr>
<tr>
<td>
<code>oauth2</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationProviderOAuth2Settings">
AuthorizationProviderOAuth2Settings
</a>
</em>
</td>
<td>
<p>Oauth2: OAuth2 settings</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a apimanagement.azure.com/Service resource</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProvider_STATUS">
Service_AuthorizationProvider_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthorizationProviderContractProperties_ARM">AuthorizationProviderContractProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProvider_Spec_ARM">Service_AuthorizationProvider_Spec_ARM</a>)
</p>
<div>
<p>Authorization Provider details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: Authorization Provider name. Must be 1 to 300 characters long.</p>
</td>
</tr>
<tr>
<td>
<code>identityProvider</code><br/>
<em>
string
</em>
</td>
<td>
<p>IdentityProvider: Identity provider name. Must be 1 to 300 characters long.</p>
</td>
</tr>
<tr>
<td>
<code>oauth2</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationProviderOAuth2Settings_ARM">
AuthorizationProviderOAuth2Settings_ARM
</a>
</em>
</td>
<td>
<p>Oauth2: OAuth2 settings</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthorizationProviderContractProperties_STATUS_ARM">AuthorizationProviderContractProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProvider_STATUS_ARM">Service_AuthorizationProvider_STATUS_ARM</a>)
</p>
<div>
<p>Authorization Provider details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: Authorization Provider name. Must be 1 to 300 characters long.</p>
</td>
</tr>
<tr>
<td>
<code>identityProvider</code><br/>
<em>
string
</em>
</td>
<td>
<p>IdentityProvider: Identity provider name. Must be 1 to 300 characters long.</p>
</td>
</tr>
<tr>
<td>
<code>oauth2</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationProviderOAuth2Settings_STATUS_ARM">
AuthorizationProviderOAuth2Settings_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Oauth2: OAuth2 settings</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthorizationProviderOAuth2GrantTypes">AuthorizationProviderOAuth2GrantTypes
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationProviderOAuth2Settings">AuthorizationProviderOAuth2Settings</a>)
</p>
<div>
<p>Authorization Provider oauth2 grant types settings</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>authorizationCode</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretMapReference">
genruntime.SecretMapReference
</a>
</em>
</td>
<td>
<p>AuthorizationCode: OAuth2 authorization code grant parameters</p>
</td>
</tr>
<tr>
<td>
<code>clientCredentials</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretMapReference">
genruntime.SecretMapReference
</a>
</em>
</td>
<td>
<p>ClientCredentials: OAuth2 client credential grant parameters</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthorizationProviderOAuth2GrantTypes_ARM">AuthorizationProviderOAuth2GrantTypes_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationProviderOAuth2Settings_ARM">AuthorizationProviderOAuth2Settings_ARM</a>)
</p>
<div>
<p>Authorization Provider oauth2 grant types settings</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>authorizationCode</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>AuthorizationCode: OAuth2 authorization code grant parameters</p>
</td>
</tr>
<tr>
<td>
<code>clientCredentials</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>ClientCredentials: OAuth2 client credential grant parameters</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthorizationProviderOAuth2GrantTypes_STATUS">AuthorizationProviderOAuth2GrantTypes_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationProviderOAuth2Settings_STATUS">AuthorizationProviderOAuth2Settings_STATUS</a>)
</p>
<div>
<p>Authorization Provider oauth2 grant types settings</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>authorizationCode</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>AuthorizationCode: OAuth2 authorization code grant parameters</p>
</td>
</tr>
<tr>
<td>
<code>clientCredentials</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>ClientCredentials: OAuth2 client credential grant parameters</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthorizationProviderOAuth2GrantTypes_STATUS_ARM">AuthorizationProviderOAuth2GrantTypes_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationProviderOAuth2Settings_STATUS_ARM">AuthorizationProviderOAuth2Settings_STATUS_ARM</a>)
</p>
<div>
<p>Authorization Provider oauth2 grant types settings</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>authorizationCode</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>AuthorizationCode: OAuth2 authorization code grant parameters</p>
</td>
</tr>
<tr>
<td>
<code>clientCredentials</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>ClientCredentials: OAuth2 client credential grant parameters</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthorizationProviderOAuth2Settings">AuthorizationProviderOAuth2Settings
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProvider_Spec">Service_AuthorizationProvider_Spec</a>)
</p>
<div>
<p>OAuth2 settings details</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>grantTypes</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationProviderOAuth2GrantTypes">
AuthorizationProviderOAuth2GrantTypes
</a>
</em>
</td>
<td>
<p>GrantTypes: OAuth2 settings</p>
</td>
</tr>
<tr>
<td>
<code>redirectUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>RedirectUrl: Redirect URL to be set in the OAuth application.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthorizationProviderOAuth2Settings_ARM">AuthorizationProviderOAuth2Settings_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationProviderContractProperties_ARM">AuthorizationProviderContractProperties_ARM</a>)
</p>
<div>
<p>OAuth2 settings details</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>grantTypes</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationProviderOAuth2GrantTypes_ARM">
AuthorizationProviderOAuth2GrantTypes_ARM
</a>
</em>
</td>
<td>
<p>GrantTypes: OAuth2 settings</p>
</td>
</tr>
<tr>
<td>
<code>redirectUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>RedirectUrl: Redirect URL to be set in the OAuth application.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthorizationProviderOAuth2Settings_STATUS">AuthorizationProviderOAuth2Settings_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProvider_STATUS">Service_AuthorizationProvider_STATUS</a>)
</p>
<div>
<p>OAuth2 settings details</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>grantTypes</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationProviderOAuth2GrantTypes_STATUS">
AuthorizationProviderOAuth2GrantTypes_STATUS
</a>
</em>
</td>
<td>
<p>GrantTypes: OAuth2 settings</p>
</td>
</tr>
<tr>
<td>
<code>redirectUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>RedirectUrl: Redirect URL to be set in the OAuth application.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthorizationProviderOAuth2Settings_STATUS_ARM">AuthorizationProviderOAuth2Settings_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationProviderContractProperties_STATUS_ARM">AuthorizationProviderContractProperties_STATUS_ARM</a>)
</p>
<div>
<p>OAuth2 settings details</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>grantTypes</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationProviderOAuth2GrantTypes_STATUS_ARM">
AuthorizationProviderOAuth2GrantTypes_STATUS_ARM
</a>
</em>
</td>
<td>
<p>GrantTypes: OAuth2 settings</p>
</td>
</tr>
<tr>
<td>
<code>redirectUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>RedirectUrl: Redirect URL to be set in the OAuth application.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthorizationProvidersAuthorization">AuthorizationProvidersAuthorization
</h3>
<div>
<p>Generator information:
- Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimauthorizationproviders.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ApiManagement/&#x200b;service/&#x200b;{serviceName}/&#x200b;authorizationProviders/&#x200b;{authorizationProviderId}/&#x200b;authorizations/&#x200b;{authorizationId}</&#x200b;p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProviders_Authorization_Spec">
Service_AuthorizationProviders_Authorization_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>authorizationType</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationContractProperties_AuthorizationType">
AuthorizationContractProperties_AuthorizationType
</a>
</em>
</td>
<td>
<p>AuthorizationType: Authorization type options</p>
</td>
</tr>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>oauth2grantType</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationContractProperties_Oauth2GrantType">
AuthorizationContractProperties_Oauth2GrantType
</a>
</em>
</td>
<td>
<p>Oauth2GrantType: OAuth2 grant type options</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a apimanagement.azure.com/AuthorizationProvider resource</p>
</td>
</tr>
<tr>
<td>
<code>parameters</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretMapReference">
genruntime.SecretMapReference
</a>
</em>
</td>
<td>
<p>Parameters: Authorization parameters</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProviders_Authorization_STATUS">
Service_AuthorizationProviders_Authorization_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.AuthorizationProvidersAuthorizationsAccessPolicy">AuthorizationProvidersAuthorizationsAccessPolicy
</h3>
<div>
<p>Generator information:
- Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimauthorizationproviders.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ApiManagement/&#x200b;service/&#x200b;{serviceName}/&#x200b;authorizationProviders/&#x200b;{authorizationProviderId}/&#x200b;authorizations/&#x200b;{authorizationId}/&#x200b;accessPolicies/&#x200b;{authorizationAccessPolicyId}</&#x200b;p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec">
Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>appIds</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AppIds: The allowed Azure Active Directory Application IDs</p>
</td>
</tr>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>objectId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ObjectId: The Object Id</p>
</td>
</tr>
<tr>
<td>
<code>objectIdFromConfig</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ConfigMapReference">
genruntime.ConfigMapReference
</a>
</em>
</td>
<td>
<p>ObjectIdFromConfig: The Object Id</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a apimanagement.azure.com/AuthorizationProvidersAuthorization resource</p>
</td>
</tr>
<tr>
<td>
<code>tenantId</code><br/>
<em>
string
</em>
</td>
<td>
<p>TenantId: The Tenant Id</p>
</td>
</tr>
<tr>
<td>
<code>tenantIdFromConfig</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ConfigMapReference">
genruntime.ConfigMapReference
</a>
</em>
</td>
<td>
<p>TenantIdFromConfig: The Tenant Id</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS">
Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Backend">Backend
</h3>
<div>
<p>Generator information:
- Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimbackends.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ApiManagement/&#x200b;service/&#x200b;{serviceName}/&#x200b;backends/&#x200b;{backendId}</&#x200b;p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_Backend_Spec">
Service_Backend_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>circuitBreaker</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendCircuitBreaker">
BackendCircuitBreaker
</a>
</em>
</td>
<td>
<p>CircuitBreaker: Backend Circuit Breaker Configuration</p>
</td>
</tr>
<tr>
<td>
<code>credentials</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendCredentialsContract">
BackendCredentialsContract
</a>
</em>
</td>
<td>
<p>Credentials: Backend Credentials Contract Properties</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Backend Description.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a apimanagement.azure.com/Service resource</p>
</td>
</tr>
<tr>
<td>
<code>pool</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendPool">
BackendPool
</a>
</em>
</td>
<td>
<p>Pool: Backend pool information</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendProperties">
BackendProperties
</a>
</em>
</td>
<td>
<p>Properties: Backend Properties contract</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_Protocol">
BackendContractProperties_Protocol
</a>
</em>
</td>
<td>
<p>Protocol: Backend communication protocol.</p>
</td>
</tr>
<tr>
<td>
<code>proxy</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendProxyContract">
BackendProxyContract
</a>
</em>
</td>
<td>
<p>Proxy: Backend gateway Contract Properties</p>
</td>
</tr>
<tr>
<td>
<code>resourceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>ResourceReference: Management Uri of the Resource in External System. This URL can be the Arm Resource Id of Logic Apps,
Function Apps or API Apps.</p>
</td>
</tr>
<tr>
<td>
<code>title</code><br/>
<em>
string
</em>
</td>
<td>
<p>Title: Backend Title.</p>
</td>
</tr>
<tr>
<td>
<code>tls</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendTlsProperties">
BackendTlsProperties
</a>
</em>
</td>
<td>
<p>Tls: Backend TLS Properties</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_Type">
BackendContractProperties_Type
</a>
</em>
</td>
<td>
<p>Type: Type of the backend. A backend can be either Single or Pool.</p>
</td>
</tr>
<tr>
<td>
<code>url</code><br/>
<em>
string
</em>
</td>
<td>
<p>Url: Runtime Url of the Backend.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_Backend_STATUS">
Service_Backend_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendAuthorizationHeaderCredentials">BackendAuthorizationHeaderCredentials
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendCredentialsContract">BackendCredentialsContract</a>)
</p>
<div>
<p>Authorization header information.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>parameter</code><br/>
<em>
string
</em>
</td>
<td>
<p>Parameter: Authentication Parameter value.</p>
</td>
</tr>
<tr>
<td>
<code>scheme</code><br/>
<em>
string
</em>
</td>
<td>
<p>Scheme: Authentication Scheme name.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendAuthorizationHeaderCredentials_ARM">BackendAuthorizationHeaderCredentials_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendCredentialsContract_ARM">BackendCredentialsContract_ARM</a>)
</p>
<div>
<p>Authorization header information.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>parameter</code><br/>
<em>
string
</em>
</td>
<td>
<p>Parameter: Authentication Parameter value.</p>
</td>
</tr>
<tr>
<td>
<code>scheme</code><br/>
<em>
string
</em>
</td>
<td>
<p>Scheme: Authentication Scheme name.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendAuthorizationHeaderCredentials_STATUS">BackendAuthorizationHeaderCredentials_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendCredentialsContract_STATUS">BackendCredentialsContract_STATUS</a>)
</p>
<div>
<p>Authorization header information.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>parameter</code><br/>
<em>
string
</em>
</td>
<td>
<p>Parameter: Authentication Parameter value.</p>
</td>
</tr>
<tr>
<td>
<code>scheme</code><br/>
<em>
string
</em>
</td>
<td>
<p>Scheme: Authentication Scheme name.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendAuthorizationHeaderCredentials_STATUS_ARM">BackendAuthorizationHeaderCredentials_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendCredentialsContract_STATUS_ARM">BackendCredentialsContract_STATUS_ARM</a>)
</p>
<div>
<p>Authorization header information.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>parameter</code><br/>
<em>
string
</em>
</td>
<td>
<p>Parameter: Authentication Parameter value.</p>
</td>
</tr>
<tr>
<td>
<code>scheme</code><br/>
<em>
string
</em>
</td>
<td>
<p>Scheme: Authentication Scheme name.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendCircuitBreaker">BackendCircuitBreaker
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Backend_Spec">Service_Backend_Spec</a>)
</p>
<div>
<p>The configuration of the backend circuit breaker</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>rules</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CircuitBreakerRule">
[]CircuitBreakerRule
</a>
</em>
</td>
<td>
<p>Rules: The rules for tripping the backend.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendCircuitBreaker_ARM">BackendCircuitBreaker_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_ARM">BackendContractProperties_ARM</a>)
</p>
<div>
<p>The configuration of the backend circuit breaker</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>rules</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CircuitBreakerRule_ARM">
[]CircuitBreakerRule_ARM
</a>
</em>
</td>
<td>
<p>Rules: The rules for tripping the backend.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendCircuitBreaker_STATUS">BackendCircuitBreaker_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Backend_STATUS">Service_Backend_STATUS</a>)
</p>
<div>
<p>The configuration of the backend circuit breaker</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>rules</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CircuitBreakerRule_STATUS">
[]CircuitBreakerRule_STATUS
</a>
</em>
</td>
<td>
<p>Rules: The rules for tripping the backend.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendCircuitBreaker_STATUS_ARM">BackendCircuitBreaker_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_STATUS_ARM">BackendContractProperties_STATUS_ARM</a>)
</p>
<div>
<p>The configuration of the backend circuit breaker</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>rules</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CircuitBreakerRule_STATUS_ARM">
[]CircuitBreakerRule_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Rules: The rules for tripping the backend.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendContractProperties_ARM">BackendContractProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Backend_Spec_ARM">Service_Backend_Spec_ARM</a>)
</p>
<div>
<p>Parameters supplied to the Create Backend operation.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>circuitBreaker</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendCircuitBreaker_ARM">
BackendCircuitBreaker_ARM
</a>
</em>
</td>
<td>
<p>CircuitBreaker: Backend Circuit Breaker Configuration</p>
</td>
</tr>
<tr>
<td>
<code>credentials</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendCredentialsContract_ARM">
BackendCredentialsContract_ARM
</a>
</em>
</td>
<td>
<p>Credentials: Backend Credentials Contract Properties</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Backend Description.</p>
</td>
</tr>
<tr>
<td>
<code>pool</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendPool_ARM">
BackendPool_ARM
</a>
</em>
</td>
<td>
<p>Pool: Backend pool information</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendProperties_ARM">
BackendProperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: Backend Properties contract</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_Protocol">
BackendContractProperties_Protocol
</a>
</em>
</td>
<td>
<p>Protocol: Backend communication protocol.</p>
</td>
</tr>
<tr>
<td>
<code>proxy</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendProxyContract_ARM">
BackendProxyContract_ARM
</a>
</em>
</td>
<td>
<p>Proxy: Backend gateway Contract Properties</p>
</td>
</tr>
<tr>
<td>
<code>resourceId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>title</code><br/>
<em>
string
</em>
</td>
<td>
<p>Title: Backend Title.</p>
</td>
</tr>
<tr>
<td>
<code>tls</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendTlsProperties_ARM">
BackendTlsProperties_ARM
</a>
</em>
</td>
<td>
<p>Tls: Backend TLS Properties</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_Type">
BackendContractProperties_Type
</a>
</em>
</td>
<td>
<p>Type: Type of the backend. A backend can be either Single or Pool.</p>
</td>
</tr>
<tr>
<td>
<code>url</code><br/>
<em>
string
</em>
</td>
<td>
<p>Url: Runtime Url of the Backend.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendContractProperties_Protocol">BackendContractProperties_Protocol
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_ARM">BackendContractProperties_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Backend_Spec">Service_Backend_Spec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;http&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;soap&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendContractProperties_Protocol_STATUS">BackendContractProperties_Protocol_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_STATUS_ARM">BackendContractProperties_STATUS_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Backend_STATUS">Service_Backend_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;http&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;soap&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendContractProperties_STATUS_ARM">BackendContractProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Backend_STATUS_ARM">Service_Backend_STATUS_ARM</a>)
</p>
<div>
<p>Parameters supplied to the Create Backend operation.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>circuitBreaker</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendCircuitBreaker_STATUS_ARM">
BackendCircuitBreaker_STATUS_ARM
</a>
</em>
</td>
<td>
<p>CircuitBreaker: Backend Circuit Breaker Configuration</p>
</td>
</tr>
<tr>
<td>
<code>credentials</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendCredentialsContract_STATUS_ARM">
BackendCredentialsContract_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Credentials: Backend Credentials Contract Properties</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Backend Description.</p>
</td>
</tr>
<tr>
<td>
<code>pool</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendPool_STATUS_ARM">
BackendPool_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Pool: Backend pool information</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendProperties_STATUS_ARM">
BackendProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: Backend Properties contract</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_Protocol_STATUS">
BackendContractProperties_Protocol_STATUS
</a>
</em>
</td>
<td>
<p>Protocol: Backend communication protocol.</p>
</td>
</tr>
<tr>
<td>
<code>proxy</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendProxyContract_STATUS_ARM">
BackendProxyContract_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Proxy: Backend gateway Contract Properties</p>
</td>
</tr>
<tr>
<td>
<code>resourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ResourceId: Management Uri of the Resource in External System. This URL can be the Arm Resource Id of Logic Apps,
Function Apps or API Apps.</p>
</td>
</tr>
<tr>
<td>
<code>title</code><br/>
<em>
string
</em>
</td>
<td>
<p>Title: Backend Title.</p>
</td>
</tr>
<tr>
<td>
<code>tls</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendTlsProperties_STATUS_ARM">
BackendTlsProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Tls: Backend TLS Properties</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_Type_STATUS">
BackendContractProperties_Type_STATUS
</a>
</em>
</td>
<td>
<p>Type: Type of the backend. A backend can be either Single or Pool.</p>
</td>
</tr>
<tr>
<td>
<code>url</code><br/>
<em>
string
</em>
</td>
<td>
<p>Url: Runtime Url of the Backend.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendContractProperties_Type">BackendContractProperties_Type
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_ARM">BackendContractProperties_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Backend_Spec">Service_Backend_Spec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Pool&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Single&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendContractProperties_Type_STATUS">BackendContractProperties_Type_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_STATUS_ARM">BackendContractProperties_STATUS_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Backend_STATUS">Service_Backend_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Pool&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Single&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendCredentialsContract">BackendCredentialsContract
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Backend_Spec">Service_Backend_Spec</a>)
</p>
<div>
<p>Details of the Credentials used to connect to Backend.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>authorization</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendAuthorizationHeaderCredentials">
BackendAuthorizationHeaderCredentials
</a>
</em>
</td>
<td>
<p>Authorization: Authorization header authentication</p>
</td>
</tr>
<tr>
<td>
<code>certificate</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Certificate: List of Client Certificate Thumbprints. Will be ignored if certificatesIds are provided.</p>
</td>
</tr>
<tr>
<td>
<code>certificateIds</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>CertificateIds: List of Client Certificate Ids.</p>
</td>
</tr>
<tr>
<td>
<code>header</code><br/>
<em>
map[string][]string
</em>
</td>
<td>
<p>Header: Header Parameter description.</p>
</td>
</tr>
<tr>
<td>
<code>query</code><br/>
<em>
map[string][]string
</em>
</td>
<td>
<p>Query: Query Parameter description.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendCredentialsContract_ARM">BackendCredentialsContract_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_ARM">BackendContractProperties_ARM</a>)
</p>
<div>
<p>Details of the Credentials used to connect to Backend.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>authorization</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendAuthorizationHeaderCredentials_ARM">
BackendAuthorizationHeaderCredentials_ARM
</a>
</em>
</td>
<td>
<p>Authorization: Authorization header authentication</p>
</td>
</tr>
<tr>
<td>
<code>certificate</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Certificate: List of Client Certificate Thumbprints. Will be ignored if certificatesIds are provided.</p>
</td>
</tr>
<tr>
<td>
<code>certificateIds</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>CertificateIds: List of Client Certificate Ids.</p>
</td>
</tr>
<tr>
<td>
<code>header</code><br/>
<em>
map[string][]string
</em>
</td>
<td>
<p>Header: Header Parameter description.</p>
</td>
</tr>
<tr>
<td>
<code>query</code><br/>
<em>
map[string][]string
</em>
</td>
<td>
<p>Query: Query Parameter description.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendCredentialsContract_STATUS">BackendCredentialsContract_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Backend_STATUS">Service_Backend_STATUS</a>)
</p>
<div>
<p>Details of the Credentials used to connect to Backend.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>authorization</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendAuthorizationHeaderCredentials_STATUS">
BackendAuthorizationHeaderCredentials_STATUS
</a>
</em>
</td>
<td>
<p>Authorization: Authorization header authentication</p>
</td>
</tr>
<tr>
<td>
<code>certificate</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Certificate: List of Client Certificate Thumbprints. Will be ignored if certificatesIds are provided.</p>
</td>
</tr>
<tr>
<td>
<code>certificateIds</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>CertificateIds: List of Client Certificate Ids.</p>
</td>
</tr>
<tr>
<td>
<code>header</code><br/>
<em>
map[string][]string
</em>
</td>
<td>
<p>Header: Header Parameter description.</p>
</td>
</tr>
<tr>
<td>
<code>query</code><br/>
<em>
map[string][]string
</em>
</td>
<td>
<p>Query: Query Parameter description.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendCredentialsContract_STATUS_ARM">BackendCredentialsContract_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_STATUS_ARM">BackendContractProperties_STATUS_ARM</a>)
</p>
<div>
<p>Details of the Credentials used to connect to Backend.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>authorization</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendAuthorizationHeaderCredentials_STATUS_ARM">
BackendAuthorizationHeaderCredentials_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Authorization: Authorization header authentication</p>
</td>
</tr>
<tr>
<td>
<code>certificate</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Certificate: List of Client Certificate Thumbprints. Will be ignored if certificatesIds are provided.</p>
</td>
</tr>
<tr>
<td>
<code>certificateIds</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>CertificateIds: List of Client Certificate Ids.</p>
</td>
</tr>
<tr>
<td>
<code>header</code><br/>
<em>
map[string][]string
</em>
</td>
<td>
<p>Header: Header Parameter description.</p>
</td>
</tr>
<tr>
<td>
<code>query</code><br/>
<em>
map[string][]string
</em>
</td>
<td>
<p>Query: Query Parameter description.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendPool">BackendPool
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Backend_Spec">Service_Backend_Spec</a>)
</p>
<div>
<p>Backend pool information</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>services</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendPoolItem">
[]BackendPoolItem
</a>
</em>
</td>
<td>
<p>Services: The list of backend entities belonging to a pool.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendPoolItem">BackendPoolItem
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendPool">BackendPool</a>)
</p>
<div>
<p>Backend pool service information</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>reference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>Reference: The unique ARM id of the backend entity. The ARM id should refer to an already existing backend entity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendPoolItem_ARM">BackendPoolItem_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendPool_ARM">BackendPool_ARM</a>)
</p>
<div>
<p>Backend pool service information</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendPoolItem_STATUS">BackendPoolItem_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendPool_STATUS">BackendPool_STATUS</a>)
</p>
<div>
<p>Backend pool service information</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: The unique ARM id of the backend entity. The ARM id should refer to an already existing backend entity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendPoolItem_STATUS_ARM">BackendPoolItem_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendPool_STATUS_ARM">BackendPool_STATUS_ARM</a>)
</p>
<div>
<p>Backend pool service information</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: The unique ARM id of the backend entity. The ARM id should refer to an already existing backend entity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendPool_ARM">BackendPool_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_ARM">BackendContractProperties_ARM</a>)
</p>
<div>
<p>Backend pool information</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>services</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendPoolItem_ARM">
[]BackendPoolItem_ARM
</a>
</em>
</td>
<td>
<p>Services: The list of backend entities belonging to a pool.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendPool_STATUS">BackendPool_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Backend_STATUS">Service_Backend_STATUS</a>)
</p>
<div>
<p>Backend pool information</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>services</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendPoolItem_STATUS">
[]BackendPoolItem_STATUS
</a>
</em>
</td>
<td>
<p>Services: The list of backend entities belonging to a pool.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendPool_STATUS_ARM">BackendPool_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_STATUS_ARM">BackendContractProperties_STATUS_ARM</a>)
</p>
<div>
<p>Backend pool information</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>services</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendPoolItem_STATUS_ARM">
[]BackendPoolItem_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Services: The list of backend entities belonging to a pool.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendProperties">BackendProperties
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Backend_Spec">Service_Backend_Spec</a>)
</p>
<div>
<p>Properties specific to the Backend Type.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>serviceFabricCluster</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendServiceFabricClusterProperties">
BackendServiceFabricClusterProperties
</a>
</em>
</td>
<td>
<p>ServiceFabricCluster: Backend Service Fabric Cluster Properties</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendProperties_ARM">BackendProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_ARM">BackendContractProperties_ARM</a>)
</p>
<div>
<p>Properties specific to the Backend Type.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>serviceFabricCluster</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendServiceFabricClusterProperties_ARM">
BackendServiceFabricClusterProperties_ARM
</a>
</em>
</td>
<td>
<p>ServiceFabricCluster: Backend Service Fabric Cluster Properties</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendProperties_STATUS">BackendProperties_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Backend_STATUS">Service_Backend_STATUS</a>)
</p>
<div>
<p>Properties specific to the Backend Type.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>serviceFabricCluster</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendServiceFabricClusterProperties_STATUS">
BackendServiceFabricClusterProperties_STATUS
</a>
</em>
</td>
<td>
<p>ServiceFabricCluster: Backend Service Fabric Cluster Properties</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendProperties_STATUS_ARM">BackendProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_STATUS_ARM">BackendContractProperties_STATUS_ARM</a>)
</p>
<div>
<p>Properties specific to the Backend Type.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>serviceFabricCluster</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendServiceFabricClusterProperties_STATUS_ARM">
BackendServiceFabricClusterProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>ServiceFabricCluster: Backend Service Fabric Cluster Properties</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendProxyContract">BackendProxyContract
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Backend_Spec">Service_Backend_Spec</a>)
</p>
<div>
<p>Details of the Backend WebProxy Server to use in the Request to Backend.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>password</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretReference">
genruntime.SecretReference
</a>
</em>
</td>
<td>
<p>Password: Password to connect to the WebProxy Server</p>
</td>
</tr>
<tr>
<td>
<code>url</code><br/>
<em>
string
</em>
</td>
<td>
<p>Url: WebProxy Server AbsoluteUri property which includes the entire URI stored in the Uri instance, including all
fragments and query strings.</p>
</td>
</tr>
<tr>
<td>
<code>username</code><br/>
<em>
string
</em>
</td>
<td>
<p>Username: Username to connect to the WebProxy server</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendProxyContract_ARM">BackendProxyContract_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_ARM">BackendContractProperties_ARM</a>)
</p>
<div>
<p>Details of the Backend WebProxy Server to use in the Request to Backend.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>password</code><br/>
<em>
string
</em>
</td>
<td>
<p>Password: Password to connect to the WebProxy Server</p>
</td>
</tr>
<tr>
<td>
<code>url</code><br/>
<em>
string
</em>
</td>
<td>
<p>Url: WebProxy Server AbsoluteUri property which includes the entire URI stored in the Uri instance, including all
fragments and query strings.</p>
</td>
</tr>
<tr>
<td>
<code>username</code><br/>
<em>
string
</em>
</td>
<td>
<p>Username: Username to connect to the WebProxy server</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendProxyContract_STATUS">BackendProxyContract_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Backend_STATUS">Service_Backend_STATUS</a>)
</p>
<div>
<p>Details of the Backend WebProxy Server to use in the Request to Backend.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>url</code><br/>
<em>
string
</em>
</td>
<td>
<p>Url: WebProxy Server AbsoluteUri property which includes the entire URI stored in the Uri instance, including all
fragments and query strings.</p>
</td>
</tr>
<tr>
<td>
<code>username</code><br/>
<em>
string
</em>
</td>
<td>
<p>Username: Username to connect to the WebProxy server</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendProxyContract_STATUS_ARM">BackendProxyContract_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_STATUS_ARM">BackendContractProperties_STATUS_ARM</a>)
</p>
<div>
<p>Details of the Backend WebProxy Server to use in the Request to Backend.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>url</code><br/>
<em>
string
</em>
</td>
<td>
<p>Url: WebProxy Server AbsoluteUri property which includes the entire URI stored in the Uri instance, including all
fragments and query strings.</p>
</td>
</tr>
<tr>
<td>
<code>username</code><br/>
<em>
string
</em>
</td>
<td>
<p>Username: Username to connect to the WebProxy server</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendServiceFabricClusterProperties">BackendServiceFabricClusterProperties
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendProperties">BackendProperties</a>)
</p>
<div>
<p>Properties of the Service Fabric Type Backend.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>clientCertificateId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClientCertificateId: The client certificate id for the management endpoint.</p>
</td>
</tr>
<tr>
<td>
<code>clientCertificatethumbprint</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClientCertificatethumbprint: The client certificate thumbprint for the management endpoint. Will be ignored if
certificatesIds are provided</p>
</td>
</tr>
<tr>
<td>
<code>managementEndpoints</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>ManagementEndpoints: The cluster management endpoint.</p>
</td>
</tr>
<tr>
<td>
<code>maxPartitionResolutionRetries</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxPartitionResolutionRetries: Maximum number of retries while attempting resolve the partition.</p>
</td>
</tr>
<tr>
<td>
<code>serverCertificateThumbprints</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>ServerCertificateThumbprints: Thumbprints of certificates cluster management service uses for tls communication</p>
</td>
</tr>
<tr>
<td>
<code>serverX509Names</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.X509CertificateName">
[]X509CertificateName
</a>
</em>
</td>
<td>
<p>ServerX509Names: Server X509 Certificate Names Collection</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendServiceFabricClusterProperties_ARM">BackendServiceFabricClusterProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendProperties_ARM">BackendProperties_ARM</a>)
</p>
<div>
<p>Properties of the Service Fabric Type Backend.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>clientCertificateId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClientCertificateId: The client certificate id for the management endpoint.</p>
</td>
</tr>
<tr>
<td>
<code>clientCertificatethumbprint</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClientCertificatethumbprint: The client certificate thumbprint for the management endpoint. Will be ignored if
certificatesIds are provided</p>
</td>
</tr>
<tr>
<td>
<code>managementEndpoints</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>ManagementEndpoints: The cluster management endpoint.</p>
</td>
</tr>
<tr>
<td>
<code>maxPartitionResolutionRetries</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxPartitionResolutionRetries: Maximum number of retries while attempting resolve the partition.</p>
</td>
</tr>
<tr>
<td>
<code>serverCertificateThumbprints</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>ServerCertificateThumbprints: Thumbprints of certificates cluster management service uses for tls communication</p>
</td>
</tr>
<tr>
<td>
<code>serverX509Names</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.X509CertificateName_ARM">
[]X509CertificateName_ARM
</a>
</em>
</td>
<td>
<p>ServerX509Names: Server X509 Certificate Names Collection</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendServiceFabricClusterProperties_STATUS">BackendServiceFabricClusterProperties_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendProperties_STATUS">BackendProperties_STATUS</a>)
</p>
<div>
<p>Properties of the Service Fabric Type Backend.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>clientCertificateId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClientCertificateId: The client certificate id for the management endpoint.</p>
</td>
</tr>
<tr>
<td>
<code>clientCertificatethumbprint</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClientCertificatethumbprint: The client certificate thumbprint for the management endpoint. Will be ignored if
certificatesIds are provided</p>
</td>
</tr>
<tr>
<td>
<code>managementEndpoints</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>ManagementEndpoints: The cluster management endpoint.</p>
</td>
</tr>
<tr>
<td>
<code>maxPartitionResolutionRetries</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxPartitionResolutionRetries: Maximum number of retries while attempting resolve the partition.</p>
</td>
</tr>
<tr>
<td>
<code>serverCertificateThumbprints</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>ServerCertificateThumbprints: Thumbprints of certificates cluster management service uses for tls communication</p>
</td>
</tr>
<tr>
<td>
<code>serverX509Names</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.X509CertificateName_STATUS">
[]X509CertificateName_STATUS
</a>
</em>
</td>
<td>
<p>ServerX509Names: Server X509 Certificate Names Collection</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendServiceFabricClusterProperties_STATUS_ARM">BackendServiceFabricClusterProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendProperties_STATUS_ARM">BackendProperties_STATUS_ARM</a>)
</p>
<div>
<p>Properties of the Service Fabric Type Backend.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>clientCertificateId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClientCertificateId: The client certificate id for the management endpoint.</p>
</td>
</tr>
<tr>
<td>
<code>clientCertificatethumbprint</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClientCertificatethumbprint: The client certificate thumbprint for the management endpoint. Will be ignored if
certificatesIds are provided</p>
</td>
</tr>
<tr>
<td>
<code>managementEndpoints</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>ManagementEndpoints: The cluster management endpoint.</p>
</td>
</tr>
<tr>
<td>
<code>maxPartitionResolutionRetries</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxPartitionResolutionRetries: Maximum number of retries while attempting resolve the partition.</p>
</td>
</tr>
<tr>
<td>
<code>serverCertificateThumbprints</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>ServerCertificateThumbprints: Thumbprints of certificates cluster management service uses for tls communication</p>
</td>
</tr>
<tr>
<td>
<code>serverX509Names</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.X509CertificateName_STATUS_ARM">
[]X509CertificateName_STATUS_ARM
</a>
</em>
</td>
<td>
<p>ServerX509Names: Server X509 Certificate Names Collection</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendTlsProperties">BackendTlsProperties
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Backend_Spec">Service_Backend_Spec</a>)
</p>
<div>
<p>Properties controlling TLS Certificate Validation.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>validateCertificateChain</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ValidateCertificateChain: Flag indicating whether SSL certificate chain validation should be done when using self-signed
certificates for this backend host.</p>
</td>
</tr>
<tr>
<td>
<code>validateCertificateName</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ValidateCertificateName: Flag indicating whether SSL certificate name validation should be done when using self-signed
certificates for this backend host.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendTlsProperties_ARM">BackendTlsProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_ARM">BackendContractProperties_ARM</a>)
</p>
<div>
<p>Properties controlling TLS Certificate Validation.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>validateCertificateChain</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ValidateCertificateChain: Flag indicating whether SSL certificate chain validation should be done when using self-signed
certificates for this backend host.</p>
</td>
</tr>
<tr>
<td>
<code>validateCertificateName</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ValidateCertificateName: Flag indicating whether SSL certificate name validation should be done when using self-signed
certificates for this backend host.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendTlsProperties_STATUS">BackendTlsProperties_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Backend_STATUS">Service_Backend_STATUS</a>)
</p>
<div>
<p>Properties controlling TLS Certificate Validation.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>validateCertificateChain</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ValidateCertificateChain: Flag indicating whether SSL certificate chain validation should be done when using self-signed
certificates for this backend host.</p>
</td>
</tr>
<tr>
<td>
<code>validateCertificateName</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ValidateCertificateName: Flag indicating whether SSL certificate name validation should be done when using self-signed
certificates for this backend host.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BackendTlsProperties_STATUS_ARM">BackendTlsProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_STATUS_ARM">BackendContractProperties_STATUS_ARM</a>)
</p>
<div>
<p>Properties controlling TLS Certificate Validation.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>validateCertificateChain</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ValidateCertificateChain: Flag indicating whether SSL certificate chain validation should be done when using self-signed
certificates for this backend host.</p>
</td>
</tr>
<tr>
<td>
<code>validateCertificateName</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ValidateCertificateName: Flag indicating whether SSL certificate name validation should be done when using self-signed
certificates for this backend host.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BearerTokenSendingMethodsContract">BearerTokenSendingMethodsContract
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.OpenIdAuthenticationSettingsContract">OpenIdAuthenticationSettingsContract</a>, <a href="#apimanagement.azure.com/v1api20230501preview.OpenIdAuthenticationSettingsContract_ARM">OpenIdAuthenticationSettingsContract_ARM</a>)
</p>
<div>
<p>Form of an authorization grant, which the client uses to request the access token.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;authorizationHeader&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;query&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.BearerTokenSendingMethodsContract_STATUS">BearerTokenSendingMethodsContract_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.OpenIdAuthenticationSettingsContract_STATUS">OpenIdAuthenticationSettingsContract_STATUS</a>, <a href="#apimanagement.azure.com/v1api20230501preview.OpenIdAuthenticationSettingsContract_STATUS_ARM">OpenIdAuthenticationSettingsContract_STATUS_ARM</a>)
</p>
<div>
<p>Form of an authorization grant, which the client uses to request the access token.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;authorizationHeader&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;query&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.CertificateConfiguration">CertificateConfiguration
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Spec">Service_Spec</a>)
</p>
<div>
<p>Certificate configuration which consist of non-trusted intermediates and root certificates.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>certificate</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CertificateInformation">
CertificateInformation
</a>
</em>
</td>
<td>
<p>Certificate: Certificate information.</p>
</td>
</tr>
<tr>
<td>
<code>certificatePassword</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretReference">
genruntime.SecretReference
</a>
</em>
</td>
<td>
<p>CertificatePassword: Certificate Password.</p>
</td>
</tr>
<tr>
<td>
<code>encodedCertificate</code><br/>
<em>
string
</em>
</td>
<td>
<p>EncodedCertificate: Base64 Encoded certificate.</p>
</td>
</tr>
<tr>
<td>
<code>storeName</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CertificateConfiguration_StoreName">
CertificateConfiguration_StoreName
</a>
</em>
</td>
<td>
<p>StoreName: The System.Security.Cryptography.x509certificates.StoreName certificate store location. Only Root and
CertificateAuthority are valid locations.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.CertificateConfiguration_ARM">CertificateConfiguration_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_ARM">ApiManagementServiceProperties_ARM</a>)
</p>
<div>
<p>Certificate configuration which consist of non-trusted intermediates and root certificates.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>certificate</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CertificateInformation_ARM">
CertificateInformation_ARM
</a>
</em>
</td>
<td>
<p>Certificate: Certificate information.</p>
</td>
</tr>
<tr>
<td>
<code>certificatePassword</code><br/>
<em>
string
</em>
</td>
<td>
<p>CertificatePassword: Certificate Password.</p>
</td>
</tr>
<tr>
<td>
<code>encodedCertificate</code><br/>
<em>
string
</em>
</td>
<td>
<p>EncodedCertificate: Base64 Encoded certificate.</p>
</td>
</tr>
<tr>
<td>
<code>storeName</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CertificateConfiguration_StoreName">
CertificateConfiguration_StoreName
</a>
</em>
</td>
<td>
<p>StoreName: The System.Security.Cryptography.x509certificates.StoreName certificate store location. Only Root and
CertificateAuthority are valid locations.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.CertificateConfiguration_STATUS">CertificateConfiguration_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_STATUS">Service_STATUS</a>)
</p>
<div>
<p>Certificate configuration which consist of non-trusted intermediates and root certificates.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>certificate</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CertificateInformation_STATUS">
CertificateInformation_STATUS
</a>
</em>
</td>
<td>
<p>Certificate: Certificate information.</p>
</td>
</tr>
<tr>
<td>
<code>encodedCertificate</code><br/>
<em>
string
</em>
</td>
<td>
<p>EncodedCertificate: Base64 Encoded certificate.</p>
</td>
</tr>
<tr>
<td>
<code>storeName</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CertificateConfiguration_StoreName_STATUS">
CertificateConfiguration_StoreName_STATUS
</a>
</em>
</td>
<td>
<p>StoreName: The System.Security.Cryptography.x509certificates.StoreName certificate store location. Only Root and
CertificateAuthority are valid locations.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.CertificateConfiguration_STATUS_ARM">CertificateConfiguration_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_STATUS_ARM">ApiManagementServiceProperties_STATUS_ARM</a>)
</p>
<div>
<p>Certificate configuration which consist of non-trusted intermediates and root certificates.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>certificate</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CertificateInformation_STATUS_ARM">
CertificateInformation_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Certificate: Certificate information.</p>
</td>
</tr>
<tr>
<td>
<code>encodedCertificate</code><br/>
<em>
string
</em>
</td>
<td>
<p>EncodedCertificate: Base64 Encoded certificate.</p>
</td>
</tr>
<tr>
<td>
<code>storeName</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CertificateConfiguration_StoreName_STATUS">
CertificateConfiguration_StoreName_STATUS
</a>
</em>
</td>
<td>
<p>StoreName: The System.Security.Cryptography.x509certificates.StoreName certificate store location. Only Root and
CertificateAuthority are valid locations.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.CertificateConfiguration_StoreName">CertificateConfiguration_StoreName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.CertificateConfiguration">CertificateConfiguration</a>, <a href="#apimanagement.azure.com/v1api20230501preview.CertificateConfiguration_ARM">CertificateConfiguration_ARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;CertificateAuthority&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Root&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.CertificateConfiguration_StoreName_STATUS">CertificateConfiguration_StoreName_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.CertificateConfiguration_STATUS">CertificateConfiguration_STATUS</a>, <a href="#apimanagement.azure.com/v1api20230501preview.CertificateConfiguration_STATUS_ARM">CertificateConfiguration_STATUS_ARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;CertificateAuthority&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Root&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.CertificateInformation">CertificateInformation
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.CertificateConfiguration">CertificateConfiguration</a>, <a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration">HostnameConfiguration</a>)
</p>
<div>
<p>SSL certificate information.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>expiry</code><br/>
<em>
string
</em>
</td>
<td>
<p>Expiry: Expiration date of the certificate. The date conforms to the following format: <code>yyyy-MM-ddTHH:mm:ssZ</code> as
specified by the ISO 8601 standard.</p>
</td>
</tr>
<tr>
<td>
<code>expiryFromConfig</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ConfigMapReference">
genruntime.ConfigMapReference
</a>
</em>
</td>
<td>
<p>ExpiryFromConfig: Expiration date of the certificate. The date conforms to the following format: <code>yyyy-MM-ddTHH:mm:ssZ</code>
as specified by the ISO 8601 standard.</p>
</td>
</tr>
<tr>
<td>
<code>subject</code><br/>
<em>
string
</em>
</td>
<td>
<p>Subject: Subject of the certificate.</p>
</td>
</tr>
<tr>
<td>
<code>subjectFromConfig</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ConfigMapReference">
genruntime.ConfigMapReference
</a>
</em>
</td>
<td>
<p>SubjectFromConfig: Subject of the certificate.</p>
</td>
</tr>
<tr>
<td>
<code>thumbprint</code><br/>
<em>
string
</em>
</td>
<td>
<p>Thumbprint: Thumbprint of the certificate.</p>
</td>
</tr>
<tr>
<td>
<code>thumbprintFromConfig</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ConfigMapReference">
genruntime.ConfigMapReference
</a>
</em>
</td>
<td>
<p>ThumbprintFromConfig: Thumbprint of the certificate.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.CertificateInformation_ARM">CertificateInformation_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.CertificateConfiguration_ARM">CertificateConfiguration_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_ARM">HostnameConfiguration_ARM</a>)
</p>
<div>
<p>SSL certificate information.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>expiry</code><br/>
<em>
string
</em>
</td>
<td>
<p>Expiry: Expiration date of the certificate. The date conforms to the following format: <code>yyyy-MM-ddTHH:mm:ssZ</code> as
specified by the ISO 8601 standard.</p>
</td>
</tr>
<tr>
<td>
<code>subject</code><br/>
<em>
string
</em>
</td>
<td>
<p>Subject: Subject of the certificate.</p>
</td>
</tr>
<tr>
<td>
<code>thumbprint</code><br/>
<em>
string
</em>
</td>
<td>
<p>Thumbprint: Thumbprint of the certificate.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.CertificateInformation_STATUS">CertificateInformation_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.CertificateConfiguration_STATUS">CertificateConfiguration_STATUS</a>, <a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_STATUS">HostnameConfiguration_STATUS</a>)
</p>
<div>
<p>SSL certificate information.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>expiry</code><br/>
<em>
string
</em>
</td>
<td>
<p>Expiry: Expiration date of the certificate. The date conforms to the following format: <code>yyyy-MM-ddTHH:mm:ssZ</code> as
specified by the ISO 8601 standard.</p>
</td>
</tr>
<tr>
<td>
<code>subject</code><br/>
<em>
string
</em>
</td>
<td>
<p>Subject: Subject of the certificate.</p>
</td>
</tr>
<tr>
<td>
<code>thumbprint</code><br/>
<em>
string
</em>
</td>
<td>
<p>Thumbprint: Thumbprint of the certificate.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.CertificateInformation_STATUS_ARM">CertificateInformation_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.CertificateConfiguration_STATUS_ARM">CertificateConfiguration_STATUS_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_STATUS_ARM">HostnameConfiguration_STATUS_ARM</a>)
</p>
<div>
<p>SSL certificate information.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>expiry</code><br/>
<em>
string
</em>
</td>
<td>
<p>Expiry: Expiration date of the certificate. The date conforms to the following format: <code>yyyy-MM-ddTHH:mm:ssZ</code> as
specified by the ISO 8601 standard.</p>
</td>
</tr>
<tr>
<td>
<code>subject</code><br/>
<em>
string
</em>
</td>
<td>
<p>Subject: Subject of the certificate.</p>
</td>
</tr>
<tr>
<td>
<code>thumbprint</code><br/>
<em>
string
</em>
</td>
<td>
<p>Thumbprint: Thumbprint of the certificate.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.CircuitBreakerFailureCondition">CircuitBreakerFailureCondition
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.CircuitBreakerRule">CircuitBreakerRule</a>)
</p>
<div>
<p>The trip conditions of the circuit breaker</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>count</code><br/>
<em>
int
</em>
</td>
<td>
<p>Count: The threshold for opening the circuit.</p>
</td>
</tr>
<tr>
<td>
<code>errorReasons</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CircuitBreakerFailureCondition_ErrorReasons">
[]CircuitBreakerFailureCondition_ErrorReasons
</a>
</em>
</td>
<td>
<p>ErrorReasons: The error reasons which are considered as failure.</p>
</td>
</tr>
<tr>
<td>
<code>interval</code><br/>
<em>
string
</em>
</td>
<td>
<p>Interval: The interval during which the failures are counted.</p>
</td>
</tr>
<tr>
<td>
<code>percentage</code><br/>
<em>
int
</em>
</td>
<td>
<p>Percentage: The threshold for opening the circuit.</p>
</td>
</tr>
<tr>
<td>
<code>statusCodeRanges</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.FailureStatusCodeRange">
[]FailureStatusCodeRange
</a>
</em>
</td>
<td>
<p>StatusCodeRanges: The status code ranges which are considered as failure.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.CircuitBreakerFailureCondition_ARM">CircuitBreakerFailureCondition_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.CircuitBreakerRule_ARM">CircuitBreakerRule_ARM</a>)
</p>
<div>
<p>The trip conditions of the circuit breaker</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>count</code><br/>
<em>
int
</em>
</td>
<td>
<p>Count: The threshold for opening the circuit.</p>
</td>
</tr>
<tr>
<td>
<code>errorReasons</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CircuitBreakerFailureCondition_ErrorReasons">
[]CircuitBreakerFailureCondition_ErrorReasons
</a>
</em>
</td>
<td>
<p>ErrorReasons: The error reasons which are considered as failure.</p>
</td>
</tr>
<tr>
<td>
<code>interval</code><br/>
<em>
string
</em>
</td>
<td>
<p>Interval: The interval during which the failures are counted.</p>
</td>
</tr>
<tr>
<td>
<code>percentage</code><br/>
<em>
int
</em>
</td>
<td>
<p>Percentage: The threshold for opening the circuit.</p>
</td>
</tr>
<tr>
<td>
<code>statusCodeRanges</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.FailureStatusCodeRange_ARM">
[]FailureStatusCodeRange_ARM
</a>
</em>
</td>
<td>
<p>StatusCodeRanges: The status code ranges which are considered as failure.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.CircuitBreakerFailureCondition_ErrorReasons">CircuitBreakerFailureCondition_ErrorReasons
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.CircuitBreakerFailureCondition">CircuitBreakerFailureCondition</a>, <a href="#apimanagement.azure.com/v1api20230501preview.CircuitBreakerFailureCondition_ARM">CircuitBreakerFailureCondition_ARM</a>)
</p>
<div>
</div>
<h3 id="apimanagement.azure.com/v1api20230501preview.CircuitBreakerFailureCondition_ErrorReasons_STATUS">CircuitBreakerFailureCondition_ErrorReasons_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.CircuitBreakerFailureCondition_STATUS">CircuitBreakerFailureCondition_STATUS</a>, <a href="#apimanagement.azure.com/v1api20230501preview.CircuitBreakerFailureCondition_STATUS_ARM">CircuitBreakerFailureCondition_STATUS_ARM</a>)
</p>
<div>
</div>
<h3 id="apimanagement.azure.com/v1api20230501preview.CircuitBreakerFailureCondition_STATUS">CircuitBreakerFailureCondition_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.CircuitBreakerRule_STATUS">CircuitBreakerRule_STATUS</a>)
</p>
<div>
<p>The trip conditions of the circuit breaker</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>count</code><br/>
<em>
int
</em>
</td>
<td>
<p>Count: The threshold for opening the circuit.</p>
</td>
</tr>
<tr>
<td>
<code>errorReasons</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CircuitBreakerFailureCondition_ErrorReasons_STATUS">
[]CircuitBreakerFailureCondition_ErrorReasons_STATUS
</a>
</em>
</td>
<td>
<p>ErrorReasons: The error reasons which are considered as failure.</p>
</td>
</tr>
<tr>
<td>
<code>interval</code><br/>
<em>
string
</em>
</td>
<td>
<p>Interval: The interval during which the failures are counted.</p>
</td>
</tr>
<tr>
<td>
<code>percentage</code><br/>
<em>
int
</em>
</td>
<td>
<p>Percentage: The threshold for opening the circuit.</p>
</td>
</tr>
<tr>
<td>
<code>statusCodeRanges</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.FailureStatusCodeRange_STATUS">
[]FailureStatusCodeRange_STATUS
</a>
</em>
</td>
<td>
<p>StatusCodeRanges: The status code ranges which are considered as failure.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.CircuitBreakerFailureCondition_STATUS_ARM">CircuitBreakerFailureCondition_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.CircuitBreakerRule_STATUS_ARM">CircuitBreakerRule_STATUS_ARM</a>)
</p>
<div>
<p>The trip conditions of the circuit breaker</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>count</code><br/>
<em>
int
</em>
</td>
<td>
<p>Count: The threshold for opening the circuit.</p>
</td>
</tr>
<tr>
<td>
<code>errorReasons</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CircuitBreakerFailureCondition_ErrorReasons_STATUS">
[]CircuitBreakerFailureCondition_ErrorReasons_STATUS
</a>
</em>
</td>
<td>
<p>ErrorReasons: The error reasons which are considered as failure.</p>
</td>
</tr>
<tr>
<td>
<code>interval</code><br/>
<em>
string
</em>
</td>
<td>
<p>Interval: The interval during which the failures are counted.</p>
</td>
</tr>
<tr>
<td>
<code>percentage</code><br/>
<em>
int
</em>
</td>
<td>
<p>Percentage: The threshold for opening the circuit.</p>
</td>
</tr>
<tr>
<td>
<code>statusCodeRanges</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.FailureStatusCodeRange_STATUS_ARM">
[]FailureStatusCodeRange_STATUS_ARM
</a>
</em>
</td>
<td>
<p>StatusCodeRanges: The status code ranges which are considered as failure.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.CircuitBreakerRule">CircuitBreakerRule
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendCircuitBreaker">BackendCircuitBreaker</a>)
</p>
<div>
<p>Rule configuration to trip the backend.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>failureCondition</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CircuitBreakerFailureCondition">
CircuitBreakerFailureCondition
</a>
</em>
</td>
<td>
<p>FailureCondition: The conditions for tripping the circuit breaker.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The rule name.</p>
</td>
</tr>
<tr>
<td>
<code>tripDuration</code><br/>
<em>
string
</em>
</td>
<td>
<p>TripDuration: The duration for which the circuit will be tripped.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.CircuitBreakerRule_ARM">CircuitBreakerRule_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendCircuitBreaker_ARM">BackendCircuitBreaker_ARM</a>)
</p>
<div>
<p>Rule configuration to trip the backend.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>failureCondition</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CircuitBreakerFailureCondition_ARM">
CircuitBreakerFailureCondition_ARM
</a>
</em>
</td>
<td>
<p>FailureCondition: The conditions for tripping the circuit breaker.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The rule name.</p>
</td>
</tr>
<tr>
<td>
<code>tripDuration</code><br/>
<em>
string
</em>
</td>
<td>
<p>TripDuration: The duration for which the circuit will be tripped.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.CircuitBreakerRule_STATUS">CircuitBreakerRule_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendCircuitBreaker_STATUS">BackendCircuitBreaker_STATUS</a>)
</p>
<div>
<p>Rule configuration to trip the backend.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>failureCondition</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CircuitBreakerFailureCondition_STATUS">
CircuitBreakerFailureCondition_STATUS
</a>
</em>
</td>
<td>
<p>FailureCondition: The conditions for tripping the circuit breaker.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The rule name.</p>
</td>
</tr>
<tr>
<td>
<code>tripDuration</code><br/>
<em>
string
</em>
</td>
<td>
<p>TripDuration: The duration for which the circuit will be tripped.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.CircuitBreakerRule_STATUS_ARM">CircuitBreakerRule_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendCircuitBreaker_STATUS_ARM">BackendCircuitBreaker_STATUS_ARM</a>)
</p>
<div>
<p>Rule configuration to trip the backend.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>failureCondition</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CircuitBreakerFailureCondition_STATUS_ARM">
CircuitBreakerFailureCondition_STATUS_ARM
</a>
</em>
</td>
<td>
<p>FailureCondition: The conditions for tripping the circuit breaker.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The rule name.</p>
</td>
</tr>
<tr>
<td>
<code>tripDuration</code><br/>
<em>
string
</em>
</td>
<td>
<p>TripDuration: The duration for which the circuit will be tripped.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ConfigurationApi">ConfigurationApi
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Spec">Service_Spec</a>)
</p>
<div>
<p>Information regarding the Configuration API of the API Management service.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>legacyApi</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ConfigurationApi_LegacyApi">
ConfigurationApi_LegacyApi
</a>
</em>
</td>
<td>
<p>LegacyApi: Indication whether or not the legacy Configuration API (v1) should be exposed on the API Management service.
Value is optional but must be &lsquo;Enabled&rsquo; or &lsquo;Disabled&rsquo;. If &lsquo;Disabled&rsquo;, legacy Configuration API (v1) will not be
available for self-hosted gateways. Default value is &lsquo;Enabled&rsquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ConfigurationApi_ARM">ConfigurationApi_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_ARM">ApiManagementServiceProperties_ARM</a>)
</p>
<div>
<p>Information regarding the Configuration API of the API Management service.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>legacyApi</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ConfigurationApi_LegacyApi">
ConfigurationApi_LegacyApi
</a>
</em>
</td>
<td>
<p>LegacyApi: Indication whether or not the legacy Configuration API (v1) should be exposed on the API Management service.
Value is optional but must be &lsquo;Enabled&rsquo; or &lsquo;Disabled&rsquo;. If &lsquo;Disabled&rsquo;, legacy Configuration API (v1) will not be
available for self-hosted gateways. Default value is &lsquo;Enabled&rsquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ConfigurationApi_LegacyApi">ConfigurationApi_LegacyApi
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ConfigurationApi">ConfigurationApi</a>, <a href="#apimanagement.azure.com/v1api20230501preview.ConfigurationApi_ARM">ConfigurationApi_ARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ConfigurationApi_LegacyApi_STATUS">ConfigurationApi_LegacyApi_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ConfigurationApi_STATUS">ConfigurationApi_STATUS</a>, <a href="#apimanagement.azure.com/v1api20230501preview.ConfigurationApi_STATUS_ARM">ConfigurationApi_STATUS_ARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ConfigurationApi_STATUS">ConfigurationApi_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_STATUS">Service_STATUS</a>)
</p>
<div>
<p>Information regarding the Configuration API of the API Management service.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>legacyApi</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ConfigurationApi_LegacyApi_STATUS">
ConfigurationApi_LegacyApi_STATUS
</a>
</em>
</td>
<td>
<p>LegacyApi: Indication whether or not the legacy Configuration API (v1) should be exposed on the API Management service.
Value is optional but must be &lsquo;Enabled&rsquo; or &lsquo;Disabled&rsquo;. If &lsquo;Disabled&rsquo;, legacy Configuration API (v1) will not be
available for self-hosted gateways. Default value is &lsquo;Enabled&rsquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ConfigurationApi_STATUS_ARM">ConfigurationApi_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_STATUS_ARM">ApiManagementServiceProperties_STATUS_ARM</a>)
</p>
<div>
<p>Information regarding the Configuration API of the API Management service.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>legacyApi</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ConfigurationApi_LegacyApi_STATUS">
ConfigurationApi_LegacyApi_STATUS
</a>
</em>
</td>
<td>
<p>LegacyApi: Indication whether or not the legacy Configuration API (v1) should be exposed on the API Management service.
Value is optional but must be &lsquo;Enabled&rsquo; or &lsquo;Disabled&rsquo;. If &lsquo;Disabled&rsquo;, legacy Configuration API (v1) will not be
available for self-hosted gateways. Default value is &lsquo;Enabled&rsquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.FailureStatusCodeRange">FailureStatusCodeRange
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.CircuitBreakerFailureCondition">CircuitBreakerFailureCondition</a>)
</p>
<div>
<p>The failure http status code range</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>max</code><br/>
<em>
int
</em>
</td>
<td>
<p>Max: The maximum http status code.</p>
</td>
</tr>
<tr>
<td>
<code>min</code><br/>
<em>
int
</em>
</td>
<td>
<p>Min: The minimum http status code.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.FailureStatusCodeRange_ARM">FailureStatusCodeRange_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.CircuitBreakerFailureCondition_ARM">CircuitBreakerFailureCondition_ARM</a>)
</p>
<div>
<p>The failure http status code range</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>max</code><br/>
<em>
int
</em>
</td>
<td>
<p>Max: The maximum http status code.</p>
</td>
</tr>
<tr>
<td>
<code>min</code><br/>
<em>
int
</em>
</td>
<td>
<p>Min: The minimum http status code.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.FailureStatusCodeRange_STATUS">FailureStatusCodeRange_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.CircuitBreakerFailureCondition_STATUS">CircuitBreakerFailureCondition_STATUS</a>)
</p>
<div>
<p>The failure http status code range</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>max</code><br/>
<em>
int
</em>
</td>
<td>
<p>Max: The maximum http status code.</p>
</td>
</tr>
<tr>
<td>
<code>min</code><br/>
<em>
int
</em>
</td>
<td>
<p>Min: The minimum http status code.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.FailureStatusCodeRange_STATUS_ARM">FailureStatusCodeRange_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.CircuitBreakerFailureCondition_STATUS_ARM">CircuitBreakerFailureCondition_STATUS_ARM</a>)
</p>
<div>
<p>The failure http status code range</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>max</code><br/>
<em>
int
</em>
</td>
<td>
<p>Max: The maximum http status code.</p>
</td>
</tr>
<tr>
<td>
<code>min</code><br/>
<em>
int
</em>
</td>
<td>
<p>Min: The minimum http status code.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.HostnameConfiguration">HostnameConfiguration
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Spec">Service_Spec</a>)
</p>
<div>
<p>Custom hostname configuration.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>certificate</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CertificateInformation">
CertificateInformation
</a>
</em>
</td>
<td>
<p>Certificate: Certificate information.</p>
</td>
</tr>
<tr>
<td>
<code>certificatePassword</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretReference">
genruntime.SecretReference
</a>
</em>
</td>
<td>
<p>CertificatePassword: Certificate Password.</p>
</td>
</tr>
<tr>
<td>
<code>certificateSource</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_CertificateSource">
HostnameConfiguration_CertificateSource
</a>
</em>
</td>
<td>
<p>CertificateSource: Certificate Source.</p>
</td>
</tr>
<tr>
<td>
<code>certificateStatus</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_CertificateStatus">
HostnameConfiguration_CertificateStatus
</a>
</em>
</td>
<td>
<p>CertificateStatus: Certificate Status.</p>
</td>
</tr>
<tr>
<td>
<code>defaultSslBinding</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DefaultSslBinding: Specify true to setup the certificate associated with this Hostname as the Default SSL Certificate.
If a client does not send the SNI header, then this will be the certificate that will be challenged. The property is
useful if a service has multiple custom hostname enabled and it needs to decide on the default ssl certificate. The
setting only applied to gateway Hostname Type.</p>
</td>
</tr>
<tr>
<td>
<code>encodedCertificate</code><br/>
<em>
string
</em>
</td>
<td>
<p>EncodedCertificate: Base64 Encoded certificate.</p>
</td>
</tr>
<tr>
<td>
<code>hostName</code><br/>
<em>
string
</em>
</td>
<td>
<p>HostName: Hostname to configure on the Api Management service.</p>
</td>
</tr>
<tr>
<td>
<code>identityClientId</code><br/>
<em>
string
</em>
</td>
<td>
<p>IdentityClientId: System or User Assigned Managed identity clientId as generated by Azure AD, which has GET access to
the keyVault containing the SSL certificate.</p>
</td>
</tr>
<tr>
<td>
<code>identityClientIdFromConfig</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ConfigMapReference">
genruntime.ConfigMapReference
</a>
</em>
</td>
<td>
<p>IdentityClientIdFromConfig: System or User Assigned Managed identity clientId as generated by Azure AD, which has GET
access to the keyVault containing the SSL certificate.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultId</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVaultId: Url to the KeyVault Secret containing the Ssl Certificate. If absolute Url containing version is provided,
auto-update of ssl certificate will not work. This requires Api Management service to be configured with aka.ms/apimmsi.
The secret should be of type <em>application/x-pkcs12</em></p>
</td>
</tr>
<tr>
<td>
<code>negotiateClientCertificate</code><br/>
<em>
bool
</em>
</td>
<td>
<p>NegotiateClientCertificate: Specify true to always negotiate client certificate on the hostname. Default Value is false.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_Type">
HostnameConfiguration_Type
</a>
</em>
</td>
<td>
<p>Type: Hostname type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_ARM">HostnameConfiguration_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_ARM">ApiManagementServiceProperties_ARM</a>)
</p>
<div>
<p>Custom hostname configuration.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>certificate</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CertificateInformation_ARM">
CertificateInformation_ARM
</a>
</em>
</td>
<td>
<p>Certificate: Certificate information.</p>
</td>
</tr>
<tr>
<td>
<code>certificatePassword</code><br/>
<em>
string
</em>
</td>
<td>
<p>CertificatePassword: Certificate Password.</p>
</td>
</tr>
<tr>
<td>
<code>certificateSource</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_CertificateSource">
HostnameConfiguration_CertificateSource
</a>
</em>
</td>
<td>
<p>CertificateSource: Certificate Source.</p>
</td>
</tr>
<tr>
<td>
<code>certificateStatus</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_CertificateStatus">
HostnameConfiguration_CertificateStatus
</a>
</em>
</td>
<td>
<p>CertificateStatus: Certificate Status.</p>
</td>
</tr>
<tr>
<td>
<code>defaultSslBinding</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DefaultSslBinding: Specify true to setup the certificate associated with this Hostname as the Default SSL Certificate.
If a client does not send the SNI header, then this will be the certificate that will be challenged. The property is
useful if a service has multiple custom hostname enabled and it needs to decide on the default ssl certificate. The
setting only applied to gateway Hostname Type.</p>
</td>
</tr>
<tr>
<td>
<code>encodedCertificate</code><br/>
<em>
string
</em>
</td>
<td>
<p>EncodedCertificate: Base64 Encoded certificate.</p>
</td>
</tr>
<tr>
<td>
<code>hostName</code><br/>
<em>
string
</em>
</td>
<td>
<p>HostName: Hostname to configure on the Api Management service.</p>
</td>
</tr>
<tr>
<td>
<code>identityClientId</code><br/>
<em>
string
</em>
</td>
<td>
<p>IdentityClientId: System or User Assigned Managed identity clientId as generated by Azure AD, which has GET access to
the keyVault containing the SSL certificate.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultId</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVaultId: Url to the KeyVault Secret containing the Ssl Certificate. If absolute Url containing version is provided,
auto-update of ssl certificate will not work. This requires Api Management service to be configured with aka.ms/apimmsi.
The secret should be of type <em>application/x-pkcs12</em></p>
</td>
</tr>
<tr>
<td>
<code>negotiateClientCertificate</code><br/>
<em>
bool
</em>
</td>
<td>
<p>NegotiateClientCertificate: Specify true to always negotiate client certificate on the hostname. Default Value is false.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_Type">
HostnameConfiguration_Type
</a>
</em>
</td>
<td>
<p>Type: Hostname type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_CertificateSource">HostnameConfiguration_CertificateSource
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration">HostnameConfiguration</a>, <a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_ARM">HostnameConfiguration_ARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;BuiltIn&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Custom&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;KeyVault&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Managed&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_CertificateSource_STATUS">HostnameConfiguration_CertificateSource_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_STATUS">HostnameConfiguration_STATUS</a>, <a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_STATUS_ARM">HostnameConfiguration_STATUS_ARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;BuiltIn&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Custom&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;KeyVault&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Managed&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_CertificateStatus">HostnameConfiguration_CertificateStatus
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration">HostnameConfiguration</a>, <a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_ARM">HostnameConfiguration_ARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Completed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Failed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;InProgress&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_CertificateStatus_STATUS">HostnameConfiguration_CertificateStatus_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_STATUS">HostnameConfiguration_STATUS</a>, <a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_STATUS_ARM">HostnameConfiguration_STATUS_ARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Completed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Failed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;InProgress&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_STATUS">HostnameConfiguration_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_STATUS">Service_STATUS</a>)
</p>
<div>
<p>Custom hostname configuration.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>certificate</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CertificateInformation_STATUS">
CertificateInformation_STATUS
</a>
</em>
</td>
<td>
<p>Certificate: Certificate information.</p>
</td>
</tr>
<tr>
<td>
<code>certificateSource</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_CertificateSource_STATUS">
HostnameConfiguration_CertificateSource_STATUS
</a>
</em>
</td>
<td>
<p>CertificateSource: Certificate Source.</p>
</td>
</tr>
<tr>
<td>
<code>certificateStatus</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_CertificateStatus_STATUS">
HostnameConfiguration_CertificateStatus_STATUS
</a>
</em>
</td>
<td>
<p>CertificateStatus: Certificate Status.</p>
</td>
</tr>
<tr>
<td>
<code>defaultSslBinding</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DefaultSslBinding: Specify true to setup the certificate associated with this Hostname as the Default SSL Certificate.
If a client does not send the SNI header, then this will be the certificate that will be challenged. The property is
useful if a service has multiple custom hostname enabled and it needs to decide on the default ssl certificate. The
setting only applied to gateway Hostname Type.</p>
</td>
</tr>
<tr>
<td>
<code>encodedCertificate</code><br/>
<em>
string
</em>
</td>
<td>
<p>EncodedCertificate: Base64 Encoded certificate.</p>
</td>
</tr>
<tr>
<td>
<code>hostName</code><br/>
<em>
string
</em>
</td>
<td>
<p>HostName: Hostname to configure on the Api Management service.</p>
</td>
</tr>
<tr>
<td>
<code>identityClientId</code><br/>
<em>
string
</em>
</td>
<td>
<p>IdentityClientId: System or User Assigned Managed identity clientId as generated by Azure AD, which has GET access to
the keyVault containing the SSL certificate.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultId</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVaultId: Url to the KeyVault Secret containing the Ssl Certificate. If absolute Url containing version is provided,
auto-update of ssl certificate will not work. This requires Api Management service to be configured with aka.ms/apimmsi.
The secret should be of type <em>application/x-pkcs12</em></p>
</td>
</tr>
<tr>
<td>
<code>negotiateClientCertificate</code><br/>
<em>
bool
</em>
</td>
<td>
<p>NegotiateClientCertificate: Specify true to always negotiate client certificate on the hostname. Default Value is false.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_Type_STATUS">
HostnameConfiguration_Type_STATUS
</a>
</em>
</td>
<td>
<p>Type: Hostname type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_STATUS_ARM">HostnameConfiguration_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_STATUS_ARM">ApiManagementServiceProperties_STATUS_ARM</a>)
</p>
<div>
<p>Custom hostname configuration.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>certificate</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CertificateInformation_STATUS_ARM">
CertificateInformation_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Certificate: Certificate information.</p>
</td>
</tr>
<tr>
<td>
<code>certificateSource</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_CertificateSource_STATUS">
HostnameConfiguration_CertificateSource_STATUS
</a>
</em>
</td>
<td>
<p>CertificateSource: Certificate Source.</p>
</td>
</tr>
<tr>
<td>
<code>certificateStatus</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_CertificateStatus_STATUS">
HostnameConfiguration_CertificateStatus_STATUS
</a>
</em>
</td>
<td>
<p>CertificateStatus: Certificate Status.</p>
</td>
</tr>
<tr>
<td>
<code>defaultSslBinding</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DefaultSslBinding: Specify true to setup the certificate associated with this Hostname as the Default SSL Certificate.
If a client does not send the SNI header, then this will be the certificate that will be challenged. The property is
useful if a service has multiple custom hostname enabled and it needs to decide on the default ssl certificate. The
setting only applied to gateway Hostname Type.</p>
</td>
</tr>
<tr>
<td>
<code>encodedCertificate</code><br/>
<em>
string
</em>
</td>
<td>
<p>EncodedCertificate: Base64 Encoded certificate.</p>
</td>
</tr>
<tr>
<td>
<code>hostName</code><br/>
<em>
string
</em>
</td>
<td>
<p>HostName: Hostname to configure on the Api Management service.</p>
</td>
</tr>
<tr>
<td>
<code>identityClientId</code><br/>
<em>
string
</em>
</td>
<td>
<p>IdentityClientId: System or User Assigned Managed identity clientId as generated by Azure AD, which has GET access to
the keyVault containing the SSL certificate.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultId</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVaultId: Url to the KeyVault Secret containing the Ssl Certificate. If absolute Url containing version is provided,
auto-update of ssl certificate will not work. This requires Api Management service to be configured with aka.ms/apimmsi.
The secret should be of type <em>application/x-pkcs12</em></p>
</td>
</tr>
<tr>
<td>
<code>negotiateClientCertificate</code><br/>
<em>
bool
</em>
</td>
<td>
<p>NegotiateClientCertificate: Specify true to always negotiate client certificate on the hostname. Default Value is false.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_Type_STATUS">
HostnameConfiguration_Type_STATUS
</a>
</em>
</td>
<td>
<p>Type: Hostname type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_Type">HostnameConfiguration_Type
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration">HostnameConfiguration</a>, <a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_ARM">HostnameConfiguration_ARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;ConfigurationApi&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;DeveloperPortal&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Management&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Portal&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Proxy&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Scm&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_Type_STATUS">HostnameConfiguration_Type_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_STATUS">HostnameConfiguration_STATUS</a>, <a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_STATUS_ARM">HostnameConfiguration_STATUS_ARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;ConfigurationApi&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;DeveloperPortal&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Management&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Portal&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Proxy&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Scm&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.KeyVaultContractCreateProperties">KeyVaultContractCreateProperties
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_NamedValue_Spec">Service_NamedValue_Spec</a>)
</p>
<div>
<p>Create keyVault contract details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>identityClientId</code><br/>
<em>
string
</em>
</td>
<td>
<p>IdentityClientId: Null for SystemAssignedIdentity or Client Id for UserAssignedIdentity , which will be used to access
key vault secret.</p>
</td>
</tr>
<tr>
<td>
<code>identityClientIdFromConfig</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ConfigMapReference">
genruntime.ConfigMapReference
</a>
</em>
</td>
<td>
<p>IdentityClientIdFromConfig: Null for SystemAssignedIdentity or Client Id for UserAssignedIdentity , which will be used
to access key vault secret.</p>
</td>
</tr>
<tr>
<td>
<code>secretIdentifier</code><br/>
<em>
string
</em>
</td>
<td>
<p>SecretIdentifier: Key vault secret identifier for fetching secret. Providing a versioned secret will prevent
auto-refresh. This requires API Management service to be configured with aka.ms/apimmsi</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.KeyVaultContractCreateProperties_ARM">KeyVaultContractCreateProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.NamedValueCreateContractProperties_ARM">NamedValueCreateContractProperties_ARM</a>)
</p>
<div>
<p>Create keyVault contract details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>identityClientId</code><br/>
<em>
string
</em>
</td>
<td>
<p>IdentityClientId: Null for SystemAssignedIdentity or Client Id for UserAssignedIdentity , which will be used to access
key vault secret.</p>
</td>
</tr>
<tr>
<td>
<code>secretIdentifier</code><br/>
<em>
string
</em>
</td>
<td>
<p>SecretIdentifier: Key vault secret identifier for fetching secret. Providing a versioned secret will prevent
auto-refresh. This requires API Management service to be configured with aka.ms/apimmsi</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.KeyVaultContractProperties_STATUS">KeyVaultContractProperties_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_NamedValue_STATUS">Service_NamedValue_STATUS</a>)
</p>
<div>
<p>KeyVault contract details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>identityClientId</code><br/>
<em>
string
</em>
</td>
<td>
<p>IdentityClientId: Null for SystemAssignedIdentity or Client Id for UserAssignedIdentity , which will be used to access
key vault secret.</p>
</td>
</tr>
<tr>
<td>
<code>lastStatus</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.KeyVaultLastAccessStatusContractProperties_STATUS">
KeyVaultLastAccessStatusContractProperties_STATUS
</a>
</em>
</td>
<td>
<p>LastStatus: Last time sync and refresh status of secret from key vault.</p>
</td>
</tr>
<tr>
<td>
<code>secretIdentifier</code><br/>
<em>
string
</em>
</td>
<td>
<p>SecretIdentifier: Key vault secret identifier for fetching secret. Providing a versioned secret will prevent
auto-refresh. This requires API Management service to be configured with aka.ms/apimmsi</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.KeyVaultContractProperties_STATUS_ARM">KeyVaultContractProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.NamedValueContractProperties_STATUS_ARM">NamedValueContractProperties_STATUS_ARM</a>)
</p>
<div>
<p>KeyVault contract details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>identityClientId</code><br/>
<em>
string
</em>
</td>
<td>
<p>IdentityClientId: Null for SystemAssignedIdentity or Client Id for UserAssignedIdentity , which will be used to access
key vault secret.</p>
</td>
</tr>
<tr>
<td>
<code>lastStatus</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.KeyVaultLastAccessStatusContractProperties_STATUS_ARM">
KeyVaultLastAccessStatusContractProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>LastStatus: Last time sync and refresh status of secret from key vault.</p>
</td>
</tr>
<tr>
<td>
<code>secretIdentifier</code><br/>
<em>
string
</em>
</td>
<td>
<p>SecretIdentifier: Key vault secret identifier for fetching secret. Providing a versioned secret will prevent
auto-refresh. This requires API Management service to be configured with aka.ms/apimmsi</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.KeyVaultLastAccessStatusContractProperties_STATUS">KeyVaultLastAccessStatusContractProperties_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.KeyVaultContractProperties_STATUS">KeyVaultContractProperties_STATUS</a>)
</p>
<div>
<p>Issue contract Update Properties.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>code</code><br/>
<em>
string
</em>
</td>
<td>
<p>Code: Last status code for sync and refresh of secret from key vault.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<p>Message: Details of the error else empty.</p>
</td>
</tr>
<tr>
<td>
<code>timeStampUtc</code><br/>
<em>
string
</em>
</td>
<td>
<p>TimeStampUtc: Last time secret was accessed. The date conforms to the following format: <code>yyyy-MM-ddTHH:mm:ssZ</code> as
specified by the ISO 8601 standard.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.KeyVaultLastAccessStatusContractProperties_STATUS_ARM">KeyVaultLastAccessStatusContractProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.KeyVaultContractProperties_STATUS_ARM">KeyVaultContractProperties_STATUS_ARM</a>)
</p>
<div>
<p>Issue contract Update Properties.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>code</code><br/>
<em>
string
</em>
</td>
<td>
<p>Code: Last status code for sync and refresh of secret from key vault.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<p>Message: Details of the error else empty.</p>
</td>
</tr>
<tr>
<td>
<code>timeStampUtc</code><br/>
<em>
string
</em>
</td>
<td>
<p>TimeStampUtc: Last time secret was accessed. The date conforms to the following format: <code>yyyy-MM-ddTHH:mm:ssZ</code> as
specified by the ISO 8601 standard.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.NamedValue">NamedValue
</h3>
<div>
<p>Generator information:
- Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimnamedvalues.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ApiManagement/&#x200b;service/&#x200b;{serviceName}/&#x200b;namedValues/&#x200b;{namedValueId}</&#x200b;p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_NamedValue_Spec">
Service_NamedValue_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: Unique name of NamedValue. It may contain only letters, digits, period, dash, and underscore characters.</p>
</td>
</tr>
<tr>
<td>
<code>keyVault</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.KeyVaultContractCreateProperties">
KeyVaultContractCreateProperties
</a>
</em>
</td>
<td>
<p>KeyVault: KeyVault location details of the namedValue.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a apimanagement.azure.com/Service resource</p>
</td>
</tr>
<tr>
<td>
<code>secret</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Secret: Determines whether the value is a secret and should be encrypted or not. Default value is false.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Tags: Optional tags that when provided can be used to filter the NamedValue list.</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Value of the NamedValue. Can contain policy expressions. It may not be empty or consist only of whitespace. This
property will not be filled on &lsquo;GET&rsquo; operations! Use &lsquo;/listSecrets&rsquo; POST request to get the value.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_NamedValue_STATUS">
Service_NamedValue_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.NamedValueContractProperties_STATUS_ARM">NamedValueContractProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_NamedValue_STATUS_ARM">Service_NamedValue_STATUS_ARM</a>)
</p>
<div>
<p>NamedValue Contract properties.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: Unique name of NamedValue. It may contain only letters, digits, period, dash, and underscore characters.</p>
</td>
</tr>
<tr>
<td>
<code>keyVault</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.KeyVaultContractProperties_STATUS_ARM">
KeyVaultContractProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>KeyVault: KeyVault location details of the namedValue.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
string
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state</p>
</td>
</tr>
<tr>
<td>
<code>secret</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Secret: Determines whether the value is a secret and should be encrypted or not. Default value is false.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Tags: Optional tags that when provided can be used to filter the NamedValue list.</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Value of the NamedValue. Can contain policy expressions. It may not be empty or consist only of whitespace. This
property will not be filled on &lsquo;GET&rsquo; operations! Use &lsquo;/listSecrets&rsquo; POST request to get the value.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.NamedValueCreateContractProperties_ARM">NamedValueCreateContractProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_NamedValue_Spec_ARM">Service_NamedValue_Spec_ARM</a>)
</p>
<div>
<p>NamedValue Contract properties.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: Unique name of NamedValue. It may contain only letters, digits, period, dash, and underscore characters.</p>
</td>
</tr>
<tr>
<td>
<code>keyVault</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.KeyVaultContractCreateProperties_ARM">
KeyVaultContractCreateProperties_ARM
</a>
</em>
</td>
<td>
<p>KeyVault: KeyVault location details of the namedValue.</p>
</td>
</tr>
<tr>
<td>
<code>secret</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Secret: Determines whether the value is a secret and should be encrypted or not. Default value is false.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Tags: Optional tags that when provided can be used to filter the NamedValue list.</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Value of the NamedValue. Can contain policy expressions. It may not be empty or consist only of whitespace. This
property will not be filled on &lsquo;GET&rsquo; operations! Use &lsquo;/listSecrets&rsquo; POST request to get the value.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.OAuth2AuthenticationSettingsContract">OAuth2AuthenticationSettingsContract
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AuthenticationSettingsContract">AuthenticationSettingsContract</a>)
</p>
<div>
<p>API OAuth2 Authentication settings details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>authorizationServerId</code><br/>
<em>
string
</em>
</td>
<td>
<p>AuthorizationServerId: OAuth authorization server identifier.</p>
</td>
</tr>
<tr>
<td>
<code>scope</code><br/>
<em>
string
</em>
</td>
<td>
<p>Scope: operations scope.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.OAuth2AuthenticationSettingsContract_ARM">OAuth2AuthenticationSettingsContract_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AuthenticationSettingsContract_ARM">AuthenticationSettingsContract_ARM</a>)
</p>
<div>
<p>API OAuth2 Authentication settings details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>authorizationServerId</code><br/>
<em>
string
</em>
</td>
<td>
<p>AuthorizationServerId: OAuth authorization server identifier.</p>
</td>
</tr>
<tr>
<td>
<code>scope</code><br/>
<em>
string
</em>
</td>
<td>
<p>Scope: operations scope.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.OAuth2AuthenticationSettingsContract_STATUS">OAuth2AuthenticationSettingsContract_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AuthenticationSettingsContract_STATUS">AuthenticationSettingsContract_STATUS</a>)
</p>
<div>
<p>API OAuth2 Authentication settings details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>authorizationServerId</code><br/>
<em>
string
</em>
</td>
<td>
<p>AuthorizationServerId: OAuth authorization server identifier.</p>
</td>
</tr>
<tr>
<td>
<code>scope</code><br/>
<em>
string
</em>
</td>
<td>
<p>Scope: operations scope.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.OAuth2AuthenticationSettingsContract_STATUS_ARM">OAuth2AuthenticationSettingsContract_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AuthenticationSettingsContract_STATUS_ARM">AuthenticationSettingsContract_STATUS_ARM</a>)
</p>
<div>
<p>API OAuth2 Authentication settings details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>authorizationServerId</code><br/>
<em>
string
</em>
</td>
<td>
<p>AuthorizationServerId: OAuth authorization server identifier.</p>
</td>
</tr>
<tr>
<td>
<code>scope</code><br/>
<em>
string
</em>
</td>
<td>
<p>Scope: operations scope.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.OpenIdAuthenticationSettingsContract">OpenIdAuthenticationSettingsContract
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AuthenticationSettingsContract">AuthenticationSettingsContract</a>)
</p>
<div>
<p>API OAuth2 Authentication settings details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>bearerTokenSendingMethods</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BearerTokenSendingMethodsContract">
[]BearerTokenSendingMethodsContract
</a>
</em>
</td>
<td>
<p>BearerTokenSendingMethods: How to send token to the server.</p>
</td>
</tr>
<tr>
<td>
<code>openidProviderId</code><br/>
<em>
string
</em>
</td>
<td>
<p>OpenidProviderId: OAuth authorization server identifier.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.OpenIdAuthenticationSettingsContract_ARM">OpenIdAuthenticationSettingsContract_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AuthenticationSettingsContract_ARM">AuthenticationSettingsContract_ARM</a>)
</p>
<div>
<p>API OAuth2 Authentication settings details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>bearerTokenSendingMethods</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BearerTokenSendingMethodsContract">
[]BearerTokenSendingMethodsContract
</a>
</em>
</td>
<td>
<p>BearerTokenSendingMethods: How to send token to the server.</p>
</td>
</tr>
<tr>
<td>
<code>openidProviderId</code><br/>
<em>
string
</em>
</td>
<td>
<p>OpenidProviderId: OAuth authorization server identifier.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.OpenIdAuthenticationSettingsContract_STATUS">OpenIdAuthenticationSettingsContract_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AuthenticationSettingsContract_STATUS">AuthenticationSettingsContract_STATUS</a>)
</p>
<div>
<p>API OAuth2 Authentication settings details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>bearerTokenSendingMethods</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BearerTokenSendingMethodsContract_STATUS">
[]BearerTokenSendingMethodsContract_STATUS
</a>
</em>
</td>
<td>
<p>BearerTokenSendingMethods: How to send token to the server.</p>
</td>
</tr>
<tr>
<td>
<code>openidProviderId</code><br/>
<em>
string
</em>
</td>
<td>
<p>OpenidProviderId: OAuth authorization server identifier.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.OpenIdAuthenticationSettingsContract_STATUS_ARM">OpenIdAuthenticationSettingsContract_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AuthenticationSettingsContract_STATUS_ARM">AuthenticationSettingsContract_STATUS_ARM</a>)
</p>
<div>
<p>API OAuth2 Authentication settings details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>bearerTokenSendingMethods</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BearerTokenSendingMethodsContract_STATUS">
[]BearerTokenSendingMethodsContract_STATUS
</a>
</em>
</td>
<td>
<p>BearerTokenSendingMethods: How to send token to the server.</p>
</td>
</tr>
<tr>
<td>
<code>openidProviderId</code><br/>
<em>
string
</em>
</td>
<td>
<p>OpenidProviderId: OAuth authorization server identifier.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Policy">Policy
</h3>
<div>
<p>Generator information:
- Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimpolicies.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ApiManagement/&#x200b;service/&#x200b;{serviceName}/&#x200b;policies/&#x200b;{policyId}</&#x200b;p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_Policy_Spec">
Service_Policy_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>format</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.PolicyContractProperties_Format">
PolicyContractProperties_Format
</a>
</em>
</td>
<td>
<p>Format: Format of the policyContent.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a apimanagement.azure.com/Service resource</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Contents of the Policy as defined by the format.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_Policy_STATUS">
Service_Policy_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.PolicyContractProperties_ARM">PolicyContractProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Policy_Spec_ARM">Service_Policy_Spec_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Products_Policy_Spec_ARM">Service_Products_Policy_Spec_ARM</a>)
</p>
<div>
<p>Policy contract Properties.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>format</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.PolicyContractProperties_Format">
PolicyContractProperties_Format
</a>
</em>
</td>
<td>
<p>Format: Format of the policyContent.</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Contents of the Policy as defined by the format.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.PolicyContractProperties_Format">PolicyContractProperties_Format
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.PolicyContractProperties_ARM">PolicyContractProperties_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Policy_Spec">Service_Policy_Spec</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Products_Policy_Spec">Service_Products_Policy_Spec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;rawxml&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;rawxml-link&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;xml&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;xml-link&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.PolicyContractProperties_Format_STATUS">PolicyContractProperties_Format_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.PolicyContractProperties_STATUS_ARM">PolicyContractProperties_STATUS_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Policy_STATUS">Service_Policy_STATUS</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Products_Policy_STATUS">Service_Products_Policy_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;rawxml&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;rawxml-link&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;xml&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;xml-link&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.PolicyContractProperties_STATUS_ARM">PolicyContractProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Policy_STATUS_ARM">Service_Policy_STATUS_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Products_Policy_STATUS_ARM">Service_Products_Policy_STATUS_ARM</a>)
</p>
<div>
<p>Policy contract Properties.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>format</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.PolicyContractProperties_Format_STATUS">
PolicyContractProperties_Format_STATUS
</a>
</em>
</td>
<td>
<p>Format: Format of the policyContent.</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Contents of the Policy as defined by the format.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.PolicyFragment">PolicyFragment
</h3>
<div>
<p>Generator information:
- Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimpolicyfragments.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ApiManagement/&#x200b;service/&#x200b;{serviceName}/&#x200b;policyFragments/&#x200b;{id}</&#x200b;p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_PolicyFragment_Spec">
Service_PolicyFragment_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Policy fragment description.</p>
</td>
</tr>
<tr>
<td>
<code>format</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.PolicyFragmentContractProperties_Format">
PolicyFragmentContractProperties_Format
</a>
</em>
</td>
<td>
<p>Format: Format of the policy fragment content.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a apimanagement.azure.com/Service resource</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Contents of the policy fragment.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_PolicyFragment_STATUS">
Service_PolicyFragment_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.PolicyFragmentContractProperties_ARM">PolicyFragmentContractProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_PolicyFragment_Spec_ARM">Service_PolicyFragment_Spec_ARM</a>)
</p>
<div>
<p>Policy fragment contract properties.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Policy fragment description.</p>
</td>
</tr>
<tr>
<td>
<code>format</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.PolicyFragmentContractProperties_Format">
PolicyFragmentContractProperties_Format
</a>
</em>
</td>
<td>
<p>Format: Format of the policy fragment content.</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Contents of the policy fragment.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.PolicyFragmentContractProperties_Format">PolicyFragmentContractProperties_Format
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.PolicyFragmentContractProperties_ARM">PolicyFragmentContractProperties_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_PolicyFragment_Spec">Service_PolicyFragment_Spec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;rawxml&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;xml&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.PolicyFragmentContractProperties_Format_STATUS">PolicyFragmentContractProperties_Format_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.PolicyFragmentContractProperties_STATUS_ARM">PolicyFragmentContractProperties_STATUS_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_PolicyFragment_STATUS">Service_PolicyFragment_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;rawxml&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;xml&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.PolicyFragmentContractProperties_STATUS_ARM">PolicyFragmentContractProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_PolicyFragment_STATUS_ARM">Service_PolicyFragment_STATUS_ARM</a>)
</p>
<div>
<p>Policy fragment contract properties.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Policy fragment description.</p>
</td>
</tr>
<tr>
<td>
<code>format</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.PolicyFragmentContractProperties_Format_STATUS">
PolicyFragmentContractProperties_Format_STATUS
</a>
</em>
</td>
<td>
<p>Format: Format of the policy fragment content.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
string
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Contents of the policy fragment.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.PrivateEndpointConnectionWrapperProperties_STATUS_ARM">PrivateEndpointConnectionWrapperProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.RemotePrivateEndpointConnectionWrapper_STATUS_ARM">RemotePrivateEndpointConnectionWrapper_STATUS_ARM</a>)
</p>
<div>
<p>Properties of the PrivateEndpointConnectProperties.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>groupIds</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>GroupIds: All the Group ids.</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpoint</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ArmIdWrapper_STATUS_ARM">
ArmIdWrapper_STATUS_ARM
</a>
</em>
</td>
<td>
<p>PrivateEndpoint: The resource of private end point.</p>
</td>
</tr>
<tr>
<td>
<code>privateLinkServiceConnectionState</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.PrivateLinkServiceConnectionState_STATUS_ARM">
PrivateLinkServiceConnectionState_STATUS_ARM
</a>
</em>
</td>
<td>
<p>PrivateLinkServiceConnectionState: A collection of information about the state of the connection between service
consumer and provider.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
string
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the private endpoint connection resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.PrivateEndpointServiceConnectionStatus_STATUS">PrivateEndpointServiceConnectionStatus_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.PrivateLinkServiceConnectionState_STATUS">PrivateLinkServiceConnectionState_STATUS</a>, <a href="#apimanagement.azure.com/v1api20230501preview.PrivateLinkServiceConnectionState_STATUS_ARM">PrivateLinkServiceConnectionState_STATUS_ARM</a>)
</p>
<div>
<p>The private endpoint connection status.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Approved&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Pending&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Rejected&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.PrivateLinkServiceConnectionState_STATUS">PrivateLinkServiceConnectionState_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.RemotePrivateEndpointConnectionWrapper_STATUS">RemotePrivateEndpointConnectionWrapper_STATUS</a>)
</p>
<div>
<p>A collection of information about the state of the connection between service consumer and provider.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>actionsRequired</code><br/>
<em>
string
</em>
</td>
<td>
<p>ActionsRequired: A message indicating if changes on the service provider require any updates on the consumer.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: The reason for approval/rejection of the connection.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.PrivateEndpointServiceConnectionStatus_STATUS">
PrivateEndpointServiceConnectionStatus_STATUS
</a>
</em>
</td>
<td>
<p>Status: Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.PrivateLinkServiceConnectionState_STATUS_ARM">PrivateLinkServiceConnectionState_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.PrivateEndpointConnectionWrapperProperties_STATUS_ARM">PrivateEndpointConnectionWrapperProperties_STATUS_ARM</a>)
</p>
<div>
<p>A collection of information about the state of the connection between service consumer and provider.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>actionsRequired</code><br/>
<em>
string
</em>
</td>
<td>
<p>ActionsRequired: A message indicating if changes on the service provider require any updates on the consumer.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: The reason for approval/rejection of the connection.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.PrivateEndpointServiceConnectionStatus_STATUS">
PrivateEndpointServiceConnectionStatus_STATUS
</a>
</em>
</td>
<td>
<p>Status: Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Product">Product
</h3>
<div>
<p>Generator information:
- Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimproducts.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ApiManagement/&#x200b;service/&#x200b;{serviceName}/&#x200b;products/&#x200b;{productId}</&#x200b;p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_Product_Spec">
Service_Product_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>approvalRequired</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ApprovalRequired: whether subscription approval is required. If false, new subscriptions will be approved automatically
enabling developers to call the product’s APIs immediately after subscribing. If true, administrators must manually
approve the subscription before the developer can any of the product’s APIs. Can be present only if
subscriptionRequired property is present and has a value of false.</p>
</td>
</tr>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Product description. May include HTML formatting tags.</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: Product name.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a apimanagement.azure.com/Service resource</p>
</td>
</tr>
<tr>
<td>
<code>state</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ProductContractProperties_State">
ProductContractProperties_State
</a>
</em>
</td>
<td>
<p>State: whether product is published or not. Published products are discoverable by users of developer portal. Non
published products are visible only to administrators. Default state of Product is notPublished.</p>
</td>
</tr>
<tr>
<td>
<code>subscriptionRequired</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SubscriptionRequired: Whether a product subscription is required for accessing APIs included in this product. If true,
the product is referred to as &ldquo;protected&rdquo; and a valid subscription key is required for a request to an API included in
the product to succeed. If false, the product is referred to as &ldquo;open&rdquo; and requests to an API included in the product
can be made without a subscription key. If property is omitted when creating a new product it&rsquo;s value is assumed to be
true.</p>
</td>
</tr>
<tr>
<td>
<code>subscriptionsLimit</code><br/>
<em>
int
</em>
</td>
<td>
<p>SubscriptionsLimit: Whether the number of subscriptions a user can have to this product at the same time. Set to null or
omit to allow unlimited per user subscriptions. Can be present only if subscriptionRequired property is present and has
a value of false.</p>
</td>
</tr>
<tr>
<td>
<code>terms</code><br/>
<em>
string
</em>
</td>
<td>
<p>Terms: Product terms of use. Developers trying to subscribe to the product will be presented and required to accept
these terms before they can complete the subscription process.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_Product_STATUS">
Service_Product_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ProductApi">ProductApi
</h3>
<div>
<p>Generator information:
- Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimproducts.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ApiManagement/&#x200b;service/&#x200b;{serviceName}/&#x200b;products/&#x200b;{productId}/&#x200b;apis/&#x200b;{apiId}</&#x200b;p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_Products_Api_Spec">
Service_Products_Api_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a apimanagement.azure.com/Product resource</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_Products_Api_STATUS">
Service_Products_Api_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ProductContractProperties_ARM">ProductContractProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Product_Spec_ARM">Service_Product_Spec_ARM</a>)
</p>
<div>
<p>Product profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>approvalRequired</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ApprovalRequired: whether subscription approval is required. If false, new subscriptions will be approved automatically
enabling developers to call the product’s APIs immediately after subscribing. If true, administrators must manually
approve the subscription before the developer can any of the product’s APIs. Can be present only if
subscriptionRequired property is present and has a value of false.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Product description. May include HTML formatting tags.</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: Product name.</p>
</td>
</tr>
<tr>
<td>
<code>state</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ProductContractProperties_State">
ProductContractProperties_State
</a>
</em>
</td>
<td>
<p>State: whether product is published or not. Published products are discoverable by users of developer portal. Non
published products are visible only to administrators. Default state of Product is notPublished.</p>
</td>
</tr>
<tr>
<td>
<code>subscriptionRequired</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SubscriptionRequired: Whether a product subscription is required for accessing APIs included in this product. If true,
the product is referred to as &ldquo;protected&rdquo; and a valid subscription key is required for a request to an API included in
the product to succeed. If false, the product is referred to as &ldquo;open&rdquo; and requests to an API included in the product
can be made without a subscription key. If property is omitted when creating a new product it&rsquo;s value is assumed to be
true.</p>
</td>
</tr>
<tr>
<td>
<code>subscriptionsLimit</code><br/>
<em>
int
</em>
</td>
<td>
<p>SubscriptionsLimit: Whether the number of subscriptions a user can have to this product at the same time. Set to null or
omit to allow unlimited per user subscriptions. Can be present only if subscriptionRequired property is present and has
a value of false.</p>
</td>
</tr>
<tr>
<td>
<code>terms</code><br/>
<em>
string
</em>
</td>
<td>
<p>Terms: Product terms of use. Developers trying to subscribe to the product will be presented and required to accept
these terms before they can complete the subscription process.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ProductContractProperties_STATUS_ARM">ProductContractProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Product_STATUS_ARM">Service_Product_STATUS_ARM</a>)
</p>
<div>
<p>Product profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>approvalRequired</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ApprovalRequired: whether subscription approval is required. If false, new subscriptions will be approved automatically
enabling developers to call the product’s APIs immediately after subscribing. If true, administrators must manually
approve the subscription before the developer can any of the product’s APIs. Can be present only if
subscriptionRequired property is present and has a value of false.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Product description. May include HTML formatting tags.</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: Product name.</p>
</td>
</tr>
<tr>
<td>
<code>state</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ProductContractProperties_State_STATUS">
ProductContractProperties_State_STATUS
</a>
</em>
</td>
<td>
<p>State: whether product is published or not. Published products are discoverable by users of developer portal. Non
published products are visible only to administrators. Default state of Product is notPublished.</p>
</td>
</tr>
<tr>
<td>
<code>subscriptionRequired</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SubscriptionRequired: Whether a product subscription is required for accessing APIs included in this product. If true,
the product is referred to as &ldquo;protected&rdquo; and a valid subscription key is required for a request to an API included in
the product to succeed. If false, the product is referred to as &ldquo;open&rdquo; and requests to an API included in the product
can be made without a subscription key. If property is omitted when creating a new product it&rsquo;s value is assumed to be
true.</p>
</td>
</tr>
<tr>
<td>
<code>subscriptionsLimit</code><br/>
<em>
int
</em>
</td>
<td>
<p>SubscriptionsLimit: Whether the number of subscriptions a user can have to this product at the same time. Set to null or
omit to allow unlimited per user subscriptions. Can be present only if subscriptionRequired property is present and has
a value of false.</p>
</td>
</tr>
<tr>
<td>
<code>terms</code><br/>
<em>
string
</em>
</td>
<td>
<p>Terms: Product terms of use. Developers trying to subscribe to the product will be presented and required to accept
these terms before they can complete the subscription process.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ProductContractProperties_State">ProductContractProperties_State
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ProductContractProperties_ARM">ProductContractProperties_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Product_Spec">Service_Product_Spec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;notPublished&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;published&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ProductContractProperties_State_STATUS">ProductContractProperties_State_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ProductContractProperties_STATUS_ARM">ProductContractProperties_STATUS_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Product_STATUS">Service_Product_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;notPublished&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;published&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.ProductPolicy">ProductPolicy
</h3>
<div>
<p>Generator information:
- Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimproducts.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ApiManagement/&#x200b;service/&#x200b;{serviceName}/&#x200b;products/&#x200b;{productId}/&#x200b;policies/&#x200b;{policyId}</&#x200b;p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_Products_Policy_Spec">
Service_Products_Policy_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>format</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.PolicyContractProperties_Format">
PolicyContractProperties_Format
</a>
</em>
</td>
<td>
<p>Format: Format of the policyContent.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a apimanagement.azure.com/Product resource</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Contents of the Policy as defined by the format.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_Products_Policy_STATUS">
Service_Products_Policy_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.RemotePrivateEndpointConnectionWrapper_STATUS">RemotePrivateEndpointConnectionWrapper_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_STATUS">Service_STATUS</a>)
</p>
<div>
<p>Remote Private Endpoint Connection resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>groupIds</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>GroupIds: All the Group ids.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Private Endpoint connection resource id</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: Private Endpoint Connection Name</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpoint</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ArmIdWrapper_STATUS">
ArmIdWrapper_STATUS
</a>
</em>
</td>
<td>
<p>PrivateEndpoint: The resource of private end point.</p>
</td>
</tr>
<tr>
<td>
<code>privateLinkServiceConnectionState</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.PrivateLinkServiceConnectionState_STATUS">
PrivateLinkServiceConnectionState_STATUS
</a>
</em>
</td>
<td>
<p>PrivateLinkServiceConnectionState: A collection of information about the state of the connection between service
consumer and provider.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
string
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the private endpoint connection resource.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: Private Endpoint Connection Resource Type</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.RemotePrivateEndpointConnectionWrapper_STATUS_ARM">RemotePrivateEndpointConnectionWrapper_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_STATUS_ARM">ApiManagementServiceProperties_STATUS_ARM</a>)
</p>
<div>
<p>Remote Private Endpoint Connection resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Private Endpoint connection resource id</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: Private Endpoint Connection Name</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.PrivateEndpointConnectionWrapperProperties_STATUS_ARM">
PrivateEndpointConnectionWrapperProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: Resource properties.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: Private Endpoint Connection Resource Type</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service">Service
</h3>
<div>
<p>Generator information:
- Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimdeployment.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ApiManagement/&#x200b;service/&#x200b;{serviceName}</&#x200b;p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_Spec">
Service_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>additionalLocations</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AdditionalLocation">
[]AdditionalLocation
</a>
</em>
</td>
<td>
<p>AdditionalLocations: Additional datacenter locations of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>apiVersionConstraint</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionConstraint">
ApiVersionConstraint
</a>
</em>
</td>
<td>
<p>ApiVersionConstraint: Control Plane Apis version constraint for the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>certificates</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CertificateConfiguration">
[]CertificateConfiguration
</a>
</em>
</td>
<td>
<p>Certificates: List of Certificates that need to be installed in the API Management service. Max supported certificates
that can be installed is 10.</p>
</td>
</tr>
<tr>
<td>
<code>configurationApi</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ConfigurationApi">
ConfigurationApi
</a>
</em>
</td>
<td>
<p>ConfigurationApi: Configuration API configuration of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>customProperties</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>CustomProperties: Custom properties of the API Management service.</br>Setting
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168</code> will disable the cipher
TLS_RSA_WITH_3DES_EDE_CBC_SHA for all TLS(1.0, 1.1 and 1.2).</br>Setting
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11</code> can be used to disable just TLS 1.1.</br>Setting
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10</code> can be used to disable TLS 1.0 on an API
Management service.</br>Setting <code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls11</code> can be
used to disable just TLS 1.1 for communications with backends.</br>Setting
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls10</code> can be used to disable TLS 1.0 for
communications with backends.</br>Setting <code>Microsoft.WindowsAzure.ApiManagement.Gateway.Protocols.Server.Http2</code> can be
used to enable HTTP2 protocol on an API Management service.</br>Not specifying any of these properties on PATCH
operation will reset omitted properties&rsquo; values to their defaults. For all the settings except Http2 the default value
is <code>True</code> if the service was created on or before April 1, 2018 and <code>False</code> otherwise. Http2 setting&rsquo;s default value is
<code>False</code>.</br></br>You can disable any of the following ciphers by using settings
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.[cipher_name]</code>: TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
TLS_RSA_WITH_AES_128_GCM_SHA256, TLS_RSA_WITH_AES_256_CBC_SHA256, TLS_RSA_WITH_AES_128_CBC_SHA256,
TLS_RSA_WITH_AES_256_CBC_SHA, TLS_RSA_WITH_AES_128_CBC_SHA. For example,
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TLS_RSA_WITH_AES_128_CBC_SHA256</code>:<code>false</code>. The default
value is <code>true</code> for them.</br> Note: The following ciphers can&rsquo;t be disabled since they are required by internal
platform components:
TLS_AES_256_GCM_SHA384,TLS_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384,TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384,TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256</p>
</td>
</tr>
<tr>
<td>
<code>developerPortalStatus</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_DeveloperPortalStatus">
ApiManagementServiceProperties_DeveloperPortalStatus
</a>
</em>
</td>
<td>
<p>DeveloperPortalStatus: Status of developer portal in this API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>disableGateway</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableGateway: Property only valid for an Api Management service deployed in multiple locations. This can be used to
disable the gateway in master region.</p>
</td>
</tr>
<tr>
<td>
<code>enableClientCertificate</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableClientCertificate: Property only meant to be used for Consumption SKU Service. This enforces a client certificate
to be presented on each request to the gateway. This also enables the ability to authenticate the certificate in the
policy on the gateway.</p>
</td>
</tr>
<tr>
<td>
<code>hostnameConfigurations</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration">
[]HostnameConfiguration
</a>
</em>
</td>
<td>
<p>HostnameConfigurations: Custom hostname configuration of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceIdentity">
ApiManagementServiceIdentity
</a>
</em>
</td>
<td>
<p>Identity: Managed service identity of the Api Management service.</p>
</td>
</tr>
<tr>
<td>
<code>legacyPortalStatus</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_LegacyPortalStatus">
ApiManagementServiceProperties_LegacyPortalStatus
</a>
</em>
</td>
<td>
<p>LegacyPortalStatus: Status of legacy portal in the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: Resource location.</p>
</td>
</tr>
<tr>
<td>
<code>natGatewayState</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_NatGatewayState">
ApiManagementServiceProperties_NatGatewayState
</a>
</em>
</td>
<td>
<p>NatGatewayState: Property can be used to enable NAT Gateway for this API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>notificationSenderEmail</code><br/>
<em>
string
</em>
</td>
<td>
<p>NotificationSenderEmail: Email address from which the notification will be sent.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a resources.azure.com/ResourceGroup resource</p>
</td>
</tr>
<tr>
<td>
<code>publicIpAddressReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>PublicIpAddressReference: Public Standard SKU IP V4 based IP address to be associated with Virtual Network deployed
service in the region. Supported only for Developer and Premium SKU being deployed in Virtual Network.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_PublicNetworkAccess">
ApiManagementServiceProperties_PublicNetworkAccess
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Whether or not public endpoint access is allowed for this API Management service.  Value is
optional but if passed in, must be &lsquo;Enabled&rsquo; or &lsquo;Disabled&rsquo;. If &lsquo;Disabled&rsquo;, private endpoints are the exclusive access
method. Default value is &lsquo;Enabled&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>publisherEmail</code><br/>
<em>
string
</em>
</td>
<td>
<p>PublisherEmail: Publisher email.</p>
</td>
</tr>
<tr>
<td>
<code>publisherName</code><br/>
<em>
string
</em>
</td>
<td>
<p>PublisherName: Publisher name.</p>
</td>
</tr>
<tr>
<td>
<code>restore</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Restore: Undelete Api Management Service if it was previously soft-deleted. If this flag is specified and set to True
all other properties will be ignored.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceSkuProperties">
ApiManagementServiceSkuProperties
</a>
</em>
</td>
<td>
<p>Sku: SKU properties of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Resource tags.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkConfiguration</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.VirtualNetworkConfiguration">
VirtualNetworkConfiguration
</a>
</em>
</td>
<td>
<p>VirtualNetworkConfiguration: Virtual network configuration of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkType</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_VirtualNetworkType">
ApiManagementServiceProperties_VirtualNetworkType
</a>
</em>
</td>
<td>
<p>VirtualNetworkType: The type of VPN in which API Management service needs to be configured in. None (Default Value)
means the API Management service is not part of any Virtual Network, External means the API Management deployment is set
up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is
setup inside a Virtual Network having an Intranet Facing Endpoint only.</p>
</td>
</tr>
<tr>
<td>
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: A list of availability zones denoting where the resource needs to come from.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_STATUS">
Service_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_ApiVersionSet_STATUS">Service_ApiVersionSet_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionSet">ApiVersionSet</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#Condition">
[]genruntime/conditions.Condition
</a>
</em>
</td>
<td>
<p>Conditions: The observed state of the resource</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Description of API Version Set.</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: Name of API Version Set</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>versionHeaderName</code><br/>
<em>
string
</em>
</td>
<td>
<p>VersionHeaderName: Name of HTTP header parameter that indicates the API Version if versioningScheme is set to <code>header</code>.</p>
</td>
</tr>
<tr>
<td>
<code>versionQueryName</code><br/>
<em>
string
</em>
</td>
<td>
<p>VersionQueryName: Name of query parameter that indicates the API Version if versioningScheme is set to <code>query</code>.</p>
</td>
</tr>
<tr>
<td>
<code>versioningScheme</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractProperties_VersioningScheme_STATUS">
ApiVersionSetContractProperties_VersioningScheme_STATUS
</a>
</em>
</td>
<td>
<p>VersioningScheme: An value that determines where the API Version identifier will be located in a HTTP request.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_ApiVersionSet_STATUS_ARM">Service_ApiVersionSet_STATUS_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractProperties_STATUS_ARM">
ApiVersionSetContractProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: API VersionSet contract properties.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_ApiVersionSet_Spec">Service_ApiVersionSet_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionSet">ApiVersionSet</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Description of API Version Set.</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: Name of API Version Set</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a apimanagement.azure.com/Service resource</p>
</td>
</tr>
<tr>
<td>
<code>versionHeaderName</code><br/>
<em>
string
</em>
</td>
<td>
<p>VersionHeaderName: Name of HTTP header parameter that indicates the API Version if versioningScheme is set to <code>header</code>.</p>
</td>
</tr>
<tr>
<td>
<code>versionQueryName</code><br/>
<em>
string
</em>
</td>
<td>
<p>VersionQueryName: Name of query parameter that indicates the API Version if versioningScheme is set to <code>query</code>.</p>
</td>
</tr>
<tr>
<td>
<code>versioningScheme</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractProperties_VersioningScheme">
ApiVersionSetContractProperties_VersioningScheme
</a>
</em>
</td>
<td>
<p>VersioningScheme: An value that determines where the API Version identifier will be located in a HTTP request.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_ApiVersionSet_Spec_ARM">Service_ApiVersionSet_Spec_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractProperties_ARM">
ApiVersionSetContractProperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: API VersionSet contract properties.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Api_STATUS">Service_Api_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Api">Api</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>APIVersion: Indicates the version identifier of the API if the API is versioned</p>
</td>
</tr>
<tr>
<td>
<code>apiRevision</code><br/>
<em>
string
</em>
</td>
<td>
<p>ApiRevision: Describes the revision of the API. If no value is provided, default revision 1 is created</p>
</td>
</tr>
<tr>
<td>
<code>apiRevisionDescription</code><br/>
<em>
string
</em>
</td>
<td>
<p>ApiRevisionDescription: Description of the API Revision.</p>
</td>
</tr>
<tr>
<td>
<code>apiVersionDescription</code><br/>
<em>
string
</em>
</td>
<td>
<p>ApiVersionDescription: Description of the API Version.</p>
</td>
</tr>
<tr>
<td>
<code>apiVersionSet</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractDetails_STATUS">
ApiVersionSetContractDetails_STATUS
</a>
</em>
</td>
<td>
<p>ApiVersionSet: Version set details</p>
</td>
</tr>
<tr>
<td>
<code>apiVersionSetId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ApiVersionSetId: A resource identifier for the related ApiVersionSet.</p>
</td>
</tr>
<tr>
<td>
<code>authenticationSettings</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthenticationSettingsContract_STATUS">
AuthenticationSettingsContract_STATUS
</a>
</em>
</td>
<td>
<p>AuthenticationSettings: Collection of authentication settings included into this API.</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#Condition">
[]genruntime/conditions.Condition
</a>
</em>
</td>
<td>
<p>Conditions: The observed state of the resource</p>
</td>
</tr>
<tr>
<td>
<code>contact</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiContactInformation_STATUS">
ApiContactInformation_STATUS
</a>
</em>
</td>
<td>
<p>Contact: Contact information for the API.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Description of the API. May include HTML formatting tags.</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: API name. Must be 1 to 300 characters long.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>isCurrent</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsCurrent: Indicates if API revision is current api revision.</p>
</td>
</tr>
<tr>
<td>
<code>isOnline</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsOnline: Indicates if API revision is accessible via the gateway.</p>
</td>
</tr>
<tr>
<td>
<code>license</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiLicenseInformation_STATUS">
ApiLicenseInformation_STATUS
</a>
</em>
</td>
<td>
<p>License: License information for the API.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>path</code><br/>
<em>
string
</em>
</td>
<td>
<p>Path: Relative URL uniquely identifying this API and all of its resource paths within the API Management service
instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public
URL for this API.</p>
</td>
</tr>
<tr>
<td>
<code>properties_type</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiContractProperties_Type_STATUS">
ApiContractProperties_Type_STATUS
</a>
</em>
</td>
<td>
<p>PropertiesType: Type of API.</p>
</td>
</tr>
<tr>
<td>
<code>protocols</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiContractProperties_Protocols_STATUS">
[]ApiContractProperties_Protocols_STATUS
</a>
</em>
</td>
<td>
<p>Protocols: Describes on which protocols the operations in this API can be invoked.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
string
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state</p>
</td>
</tr>
<tr>
<td>
<code>serviceUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServiceUrl: Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.</p>
</td>
</tr>
<tr>
<td>
<code>sourceApiId</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourceApiId: API identifier of the source API.</p>
</td>
</tr>
<tr>
<td>
<code>subscriptionKeyParameterNames</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.SubscriptionKeyParameterNamesContract_STATUS">
SubscriptionKeyParameterNamesContract_STATUS
</a>
</em>
</td>
<td>
<p>SubscriptionKeyParameterNames: Protocols over which API is made available.</p>
</td>
</tr>
<tr>
<td>
<code>subscriptionRequired</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SubscriptionRequired: Specifies whether an API or Product subscription is required for accessing the API.</p>
</td>
</tr>
<tr>
<td>
<code>termsOfServiceUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>TermsOfServiceUrl:  A URL to the Terms of Service for the API. MUST be in the format of a URL.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Api_STATUS_ARM">Service_Api_STATUS_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiContractProperties_STATUS_ARM">
ApiContractProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: API entity contract properties.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Api_Spec">Service_Api_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Api">Api</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>APIVersion: Indicates the version identifier of the API if the API is versioned</p>
</td>
</tr>
<tr>
<td>
<code>apiRevision</code><br/>
<em>
string
</em>
</td>
<td>
<p>ApiRevision: Describes the revision of the API. If no value is provided, default revision 1 is created</p>
</td>
</tr>
<tr>
<td>
<code>apiRevisionDescription</code><br/>
<em>
string
</em>
</td>
<td>
<p>ApiRevisionDescription: Description of the API Revision.</p>
</td>
</tr>
<tr>
<td>
<code>apiType</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_ApiType">
ApiCreateOrUpdateProperties_ApiType
</a>
</em>
</td>
<td>
<p>ApiType: Type of API to create.
* <code>http</code> creates a REST API
* <code>soap</code> creates a SOAP pass-through API
* <code>websocket</code> creates websocket API
* <code>graphql</code> creates GraphQL API.
New types can be added in the future.</p>
</td>
</tr>
<tr>
<td>
<code>apiVersionDescription</code><br/>
<em>
string
</em>
</td>
<td>
<p>ApiVersionDescription: Description of the API Version.</p>
</td>
</tr>
<tr>
<td>
<code>apiVersionSet</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionSetContractDetails">
ApiVersionSetContractDetails
</a>
</em>
</td>
<td>
<p>ApiVersionSet: Version set details</p>
</td>
</tr>
<tr>
<td>
<code>apiVersionSetReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>ApiVersionSetReference: A resource identifier for the related ApiVersionSet.</p>
</td>
</tr>
<tr>
<td>
<code>authenticationSettings</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthenticationSettingsContract">
AuthenticationSettingsContract
</a>
</em>
</td>
<td>
<p>AuthenticationSettings: Collection of authentication settings included into this API.</p>
</td>
</tr>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>contact</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiContactInformation">
ApiContactInformation
</a>
</em>
</td>
<td>
<p>Contact: Contact information for the API.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Description of the API. May include HTML formatting tags.</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: API name. Must be 1 to 300 characters long.</p>
</td>
</tr>
<tr>
<td>
<code>format</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_Format">
ApiCreateOrUpdateProperties_Format
</a>
</em>
</td>
<td>
<p>Format: Format of the Content in which the API is getting imported. New formats can be added in the future</p>
</td>
</tr>
<tr>
<td>
<code>isCurrent</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsCurrent: Indicates if API revision is current api revision.</p>
</td>
</tr>
<tr>
<td>
<code>license</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiLicenseInformation">
ApiLicenseInformation
</a>
</em>
</td>
<td>
<p>License: License information for the API.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a apimanagement.azure.com/Service resource</p>
</td>
</tr>
<tr>
<td>
<code>path</code><br/>
<em>
string
</em>
</td>
<td>
<p>Path: Relative URL uniquely identifying this API and all of its resource paths within the API Management service
instance. It is appended to the API endpoint base URL specified during the service instance creation to form a public
URL for this API.</p>
</td>
</tr>
<tr>
<td>
<code>protocols</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_Protocols">
[]ApiCreateOrUpdateProperties_Protocols
</a>
</em>
</td>
<td>
<p>Protocols: Describes on which protocols the operations in this API can be invoked.</p>
</td>
</tr>
<tr>
<td>
<code>serviceUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServiceUrl: Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long.</p>
</td>
</tr>
<tr>
<td>
<code>sourceApiReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>SourceApiReference: API identifier of the source API.</p>
</td>
</tr>
<tr>
<td>
<code>subscriptionKeyParameterNames</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.SubscriptionKeyParameterNamesContract">
SubscriptionKeyParameterNamesContract
</a>
</em>
</td>
<td>
<p>SubscriptionKeyParameterNames: Protocols over which API is made available.</p>
</td>
</tr>
<tr>
<td>
<code>subscriptionRequired</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SubscriptionRequired: Specifies whether an API or Product subscription is required for accessing the API.</p>
</td>
</tr>
<tr>
<td>
<code>termsOfServiceUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>TermsOfServiceUrl:  A URL to the Terms of Service for the API. MUST be in the format of a URL.</p>
</td>
</tr>
<tr>
<td>
<code>translateRequiredQueryParameters</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_TranslateRequiredQueryParameters">
ApiCreateOrUpdateProperties_TranslateRequiredQueryParameters
</a>
</em>
</td>
<td>
<p>TranslateRequiredQueryParameters: Strategy of translating required query parameters to template ones. By default has
value &lsquo;template&rsquo;. Possible values: &lsquo;template&rsquo;, &lsquo;query&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_Type">
ApiCreateOrUpdateProperties_Type
</a>
</em>
</td>
<td>
<p>Type: Type of API.</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Content value when Importing an API.</p>
</td>
</tr>
<tr>
<td>
<code>wsdlSelector</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_WsdlSelector">
ApiCreateOrUpdateProperties_WsdlSelector
</a>
</em>
</td>
<td>
<p>WsdlSelector: Criteria to limit import of WSDL to a subset of the document.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Api_Spec_ARM">Service_Api_Spec_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_ARM">
ApiCreateOrUpdateProperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: API entity create of update properties.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProvider_STATUS">Service_AuthorizationProvider_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationProvider">AuthorizationProvider</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#Condition">
[]genruntime/conditions.Condition
</a>
</em>
</td>
<td>
<p>Conditions: The observed state of the resource</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: Authorization Provider name. Must be 1 to 300 characters long.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>identityProvider</code><br/>
<em>
string
</em>
</td>
<td>
<p>IdentityProvider: Identity provider name. Must be 1 to 300 characters long.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>oauth2</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationProviderOAuth2Settings_STATUS">
AuthorizationProviderOAuth2Settings_STATUS
</a>
</em>
</td>
<td>
<p>Oauth2: OAuth2 settings</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProvider_STATUS_ARM">Service_AuthorizationProvider_STATUS_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationProviderContractProperties_STATUS_ARM">
AuthorizationProviderContractProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the Authorization Provider Contract.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProvider_Spec">Service_AuthorizationProvider_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationProvider">AuthorizationProvider</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: Authorization Provider name. Must be 1 to 300 characters long.</p>
</td>
</tr>
<tr>
<td>
<code>identityProvider</code><br/>
<em>
string
</em>
</td>
<td>
<p>IdentityProvider: Identity provider name. Must be 1 to 300 characters long.</p>
</td>
</tr>
<tr>
<td>
<code>oauth2</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationProviderOAuth2Settings">
AuthorizationProviderOAuth2Settings
</a>
</em>
</td>
<td>
<p>Oauth2: OAuth2 settings</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a apimanagement.azure.com/Service resource</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProvider_Spec_ARM">Service_AuthorizationProvider_Spec_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationProviderContractProperties_ARM">
AuthorizationProviderContractProperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the Authorization Provider Contract.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProviders_Authorization_STATUS">Service_AuthorizationProviders_Authorization_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationProvidersAuthorization">AuthorizationProvidersAuthorization</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>authorizationType</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationContractProperties_AuthorizationType_STATUS">
AuthorizationContractProperties_AuthorizationType_STATUS
</a>
</em>
</td>
<td>
<p>AuthorizationType: Authorization type options</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#Condition">
[]genruntime/conditions.Condition
</a>
</em>
</td>
<td>
<p>Conditions: The observed state of the resource</p>
</td>
</tr>
<tr>
<td>
<code>error</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationError_STATUS">
AuthorizationError_STATUS
</a>
</em>
</td>
<td>
<p>Error: Authorization error details.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>oauth2grantType</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationContractProperties_Oauth2GrantType_STATUS">
AuthorizationContractProperties_Oauth2GrantType_STATUS
</a>
</em>
</td>
<td>
<p>Oauth2GrantType: OAuth2 grant type options</p>
</td>
</tr>
<tr>
<td>
<code>parameters</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Parameters: Authorization parameters</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
string
</em>
</td>
<td>
<p>Status: Status of the Authorization</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProviders_Authorization_STATUS_ARM">Service_AuthorizationProviders_Authorization_STATUS_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationContractProperties_STATUS_ARM">
AuthorizationContractProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the Authorization Contract.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProviders_Authorization_Spec">Service_AuthorizationProviders_Authorization_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationProvidersAuthorization">AuthorizationProvidersAuthorization</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>authorizationType</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationContractProperties_AuthorizationType">
AuthorizationContractProperties_AuthorizationType
</a>
</em>
</td>
<td>
<p>AuthorizationType: Authorization type options</p>
</td>
</tr>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>oauth2grantType</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationContractProperties_Oauth2GrantType">
AuthorizationContractProperties_Oauth2GrantType
</a>
</em>
</td>
<td>
<p>Oauth2GrantType: OAuth2 grant type options</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a apimanagement.azure.com/AuthorizationProvider resource</p>
</td>
</tr>
<tr>
<td>
<code>parameters</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretMapReference">
genruntime.SecretMapReference
</a>
</em>
</td>
<td>
<p>Parameters: Authorization parameters</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProviders_Authorization_Spec_ARM">Service_AuthorizationProviders_Authorization_Spec_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationContractProperties_ARM">
AuthorizationContractProperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the Authorization Contract.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS">Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationProvidersAuthorizationsAccessPolicy">AuthorizationProvidersAuthorizationsAccessPolicy</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>appIds</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AppIds: The allowed Azure Active Directory Application IDs</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#Condition">
[]genruntime/conditions.Condition
</a>
</em>
</td>
<td>
<p>Conditions: The observed state of the resource</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>objectId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ObjectId: The Object Id</p>
</td>
</tr>
<tr>
<td>
<code>tenantId</code><br/>
<em>
string
</em>
</td>
<td>
<p>TenantId: The Tenant Id</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS_ARM">Service_AuthorizationProviders_Authorizations_AccessPolicy_STATUS_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationAccessPolicyContractProperties_STATUS_ARM">
AuthorizationAccessPolicyContractProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the Authorization Contract.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec">Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationProvidersAuthorizationsAccessPolicy">AuthorizationProvidersAuthorizationsAccessPolicy</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>appIds</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AppIds: The allowed Azure Active Directory Application IDs</p>
</td>
</tr>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>objectId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ObjectId: The Object Id</p>
</td>
</tr>
<tr>
<td>
<code>objectIdFromConfig</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ConfigMapReference">
genruntime.ConfigMapReference
</a>
</em>
</td>
<td>
<p>ObjectIdFromConfig: The Object Id</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a apimanagement.azure.com/AuthorizationProvidersAuthorization resource</p>
</td>
</tr>
<tr>
<td>
<code>tenantId</code><br/>
<em>
string
</em>
</td>
<td>
<p>TenantId: The Tenant Id</p>
</td>
</tr>
<tr>
<td>
<code>tenantIdFromConfig</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ConfigMapReference">
genruntime.ConfigMapReference
</a>
</em>
</td>
<td>
<p>TenantIdFromConfig: The Tenant Id</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec_ARM">Service_AuthorizationProviders_Authorizations_AccessPolicy_Spec_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AuthorizationAccessPolicyContractProperties_ARM">
AuthorizationAccessPolicyContractProperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the Authorization Contract.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Backend_STATUS">Service_Backend_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Backend">Backend</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>circuitBreaker</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendCircuitBreaker_STATUS">
BackendCircuitBreaker_STATUS
</a>
</em>
</td>
<td>
<p>CircuitBreaker: Backend Circuit Breaker Configuration</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#Condition">
[]genruntime/conditions.Condition
</a>
</em>
</td>
<td>
<p>Conditions: The observed state of the resource</p>
</td>
</tr>
<tr>
<td>
<code>credentials</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendCredentialsContract_STATUS">
BackendCredentialsContract_STATUS
</a>
</em>
</td>
<td>
<p>Credentials: Backend Credentials Contract Properties</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Backend Description.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>pool</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendPool_STATUS">
BackendPool_STATUS
</a>
</em>
</td>
<td>
<p>Pool: Backend pool information</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendProperties_STATUS">
BackendProperties_STATUS
</a>
</em>
</td>
<td>
<p>Properties: Backend Properties contract</p>
</td>
</tr>
<tr>
<td>
<code>properties_type</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_Type_STATUS">
BackendContractProperties_Type_STATUS
</a>
</em>
</td>
<td>
<p>PropertiesType: Type of the backend. A backend can be either Single or Pool.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_Protocol_STATUS">
BackendContractProperties_Protocol_STATUS
</a>
</em>
</td>
<td>
<p>Protocol: Backend communication protocol.</p>
</td>
</tr>
<tr>
<td>
<code>proxy</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendProxyContract_STATUS">
BackendProxyContract_STATUS
</a>
</em>
</td>
<td>
<p>Proxy: Backend gateway Contract Properties</p>
</td>
</tr>
<tr>
<td>
<code>resourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ResourceId: Management Uri of the Resource in External System. This URL can be the Arm Resource Id of Logic Apps,
Function Apps or API Apps.</p>
</td>
</tr>
<tr>
<td>
<code>title</code><br/>
<em>
string
</em>
</td>
<td>
<p>Title: Backend Title.</p>
</td>
</tr>
<tr>
<td>
<code>tls</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendTlsProperties_STATUS">
BackendTlsProperties_STATUS
</a>
</em>
</td>
<td>
<p>Tls: Backend TLS Properties</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>url</code><br/>
<em>
string
</em>
</td>
<td>
<p>Url: Runtime Url of the Backend.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Backend_STATUS_ARM">Service_Backend_STATUS_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_STATUS_ARM">
BackendContractProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: Backend entity contract properties.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Backend_Spec">Service_Backend_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Backend">Backend</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>circuitBreaker</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendCircuitBreaker">
BackendCircuitBreaker
</a>
</em>
</td>
<td>
<p>CircuitBreaker: Backend Circuit Breaker Configuration</p>
</td>
</tr>
<tr>
<td>
<code>credentials</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendCredentialsContract">
BackendCredentialsContract
</a>
</em>
</td>
<td>
<p>Credentials: Backend Credentials Contract Properties</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Backend Description.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a apimanagement.azure.com/Service resource</p>
</td>
</tr>
<tr>
<td>
<code>pool</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendPool">
BackendPool
</a>
</em>
</td>
<td>
<p>Pool: Backend pool information</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendProperties">
BackendProperties
</a>
</em>
</td>
<td>
<p>Properties: Backend Properties contract</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_Protocol">
BackendContractProperties_Protocol
</a>
</em>
</td>
<td>
<p>Protocol: Backend communication protocol.</p>
</td>
</tr>
<tr>
<td>
<code>proxy</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendProxyContract">
BackendProxyContract
</a>
</em>
</td>
<td>
<p>Proxy: Backend gateway Contract Properties</p>
</td>
</tr>
<tr>
<td>
<code>resourceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>ResourceReference: Management Uri of the Resource in External System. This URL can be the Arm Resource Id of Logic Apps,
Function Apps or API Apps.</p>
</td>
</tr>
<tr>
<td>
<code>title</code><br/>
<em>
string
</em>
</td>
<td>
<p>Title: Backend Title.</p>
</td>
</tr>
<tr>
<td>
<code>tls</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendTlsProperties">
BackendTlsProperties
</a>
</em>
</td>
<td>
<p>Tls: Backend TLS Properties</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_Type">
BackendContractProperties_Type
</a>
</em>
</td>
<td>
<p>Type: Type of the backend. A backend can be either Single or Pool.</p>
</td>
</tr>
<tr>
<td>
<code>url</code><br/>
<em>
string
</em>
</td>
<td>
<p>Url: Runtime Url of the Backend.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Backend_Spec_ARM">Service_Backend_Spec_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.BackendContractProperties_ARM">
BackendContractProperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: Backend entity contract properties.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_NamedValue_STATUS">Service_NamedValue_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.NamedValue">NamedValue</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#Condition">
[]genruntime/conditions.Condition
</a>
</em>
</td>
<td>
<p>Conditions: The observed state of the resource</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: Unique name of NamedValue. It may contain only letters, digits, period, dash, and underscore characters.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>keyVault</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.KeyVaultContractProperties_STATUS">
KeyVaultContractProperties_STATUS
</a>
</em>
</td>
<td>
<p>KeyVault: KeyVault location details of the namedValue.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
string
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state</p>
</td>
</tr>
<tr>
<td>
<code>secret</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Secret: Determines whether the value is a secret and should be encrypted or not. Default value is false.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Tags: Optional tags that when provided can be used to filter the NamedValue list.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Value of the NamedValue. Can contain policy expressions. It may not be empty or consist only of whitespace. This
property will not be filled on &lsquo;GET&rsquo; operations! Use &lsquo;/listSecrets&rsquo; POST request to get the value.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_NamedValue_STATUS_ARM">Service_NamedValue_STATUS_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.NamedValueContractProperties_STATUS_ARM">
NamedValueContractProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: NamedValue entity contract properties.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_NamedValue_Spec">Service_NamedValue_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.NamedValue">NamedValue</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: Unique name of NamedValue. It may contain only letters, digits, period, dash, and underscore characters.</p>
</td>
</tr>
<tr>
<td>
<code>keyVault</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.KeyVaultContractCreateProperties">
KeyVaultContractCreateProperties
</a>
</em>
</td>
<td>
<p>KeyVault: KeyVault location details of the namedValue.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a apimanagement.azure.com/Service resource</p>
</td>
</tr>
<tr>
<td>
<code>secret</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Secret: Determines whether the value is a secret and should be encrypted or not. Default value is false.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Tags: Optional tags that when provided can be used to filter the NamedValue list.</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Value of the NamedValue. Can contain policy expressions. It may not be empty or consist only of whitespace. This
property will not be filled on &lsquo;GET&rsquo; operations! Use &lsquo;/listSecrets&rsquo; POST request to get the value.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_NamedValue_Spec_ARM">Service_NamedValue_Spec_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.NamedValueCreateContractProperties_ARM">
NamedValueCreateContractProperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: NamedValue entity contract properties for PUT operation.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_PolicyFragment_STATUS">Service_PolicyFragment_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.PolicyFragment">PolicyFragment</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#Condition">
[]genruntime/conditions.Condition
</a>
</em>
</td>
<td>
<p>Conditions: The observed state of the resource</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Policy fragment description.</p>
</td>
</tr>
<tr>
<td>
<code>format</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.PolicyFragmentContractProperties_Format_STATUS">
PolicyFragmentContractProperties_Format_STATUS
</a>
</em>
</td>
<td>
<p>Format: Format of the policy fragment content.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
string
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Contents of the policy fragment.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_PolicyFragment_STATUS_ARM">Service_PolicyFragment_STATUS_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.PolicyFragmentContractProperties_STATUS_ARM">
PolicyFragmentContractProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the Policy Fragment.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_PolicyFragment_Spec">Service_PolicyFragment_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.PolicyFragment">PolicyFragment</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Policy fragment description.</p>
</td>
</tr>
<tr>
<td>
<code>format</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.PolicyFragmentContractProperties_Format">
PolicyFragmentContractProperties_Format
</a>
</em>
</td>
<td>
<p>Format: Format of the policy fragment content.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a apimanagement.azure.com/Service resource</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Contents of the policy fragment.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_PolicyFragment_Spec_ARM">Service_PolicyFragment_Spec_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.PolicyFragmentContractProperties_ARM">
PolicyFragmentContractProperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the Policy Fragment.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Policy_STATUS">Service_Policy_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Policy">Policy</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#Condition">
[]genruntime/conditions.Condition
</a>
</em>
</td>
<td>
<p>Conditions: The observed state of the resource</p>
</td>
</tr>
<tr>
<td>
<code>format</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.PolicyContractProperties_Format_STATUS">
PolicyContractProperties_Format_STATUS
</a>
</em>
</td>
<td>
<p>Format: Format of the policyContent.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Contents of the Policy as defined by the format.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Policy_STATUS_ARM">Service_Policy_STATUS_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.PolicyContractProperties_STATUS_ARM">
PolicyContractProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the Policy.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Policy_Spec">Service_Policy_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Policy">Policy</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>format</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.PolicyContractProperties_Format">
PolicyContractProperties_Format
</a>
</em>
</td>
<td>
<p>Format: Format of the policyContent.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a apimanagement.azure.com/Service resource</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Contents of the Policy as defined by the format.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Policy_Spec_ARM">Service_Policy_Spec_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.PolicyContractProperties_ARM">
PolicyContractProperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the Policy.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Product_STATUS">Service_Product_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Product">Product</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>approvalRequired</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ApprovalRequired: whether subscription approval is required. If false, new subscriptions will be approved automatically
enabling developers to call the product’s APIs immediately after subscribing. If true, administrators must manually
approve the subscription before the developer can any of the product’s APIs. Can be present only if
subscriptionRequired property is present and has a value of false.</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#Condition">
[]genruntime/conditions.Condition
</a>
</em>
</td>
<td>
<p>Conditions: The observed state of the resource</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Product description. May include HTML formatting tags.</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: Product name.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>state</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ProductContractProperties_State_STATUS">
ProductContractProperties_State_STATUS
</a>
</em>
</td>
<td>
<p>State: whether product is published or not. Published products are discoverable by users of developer portal. Non
published products are visible only to administrators. Default state of Product is notPublished.</p>
</td>
</tr>
<tr>
<td>
<code>subscriptionRequired</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SubscriptionRequired: Whether a product subscription is required for accessing APIs included in this product. If true,
the product is referred to as &ldquo;protected&rdquo; and a valid subscription key is required for a request to an API included in
the product to succeed. If false, the product is referred to as &ldquo;open&rdquo; and requests to an API included in the product
can be made without a subscription key. If property is omitted when creating a new product it&rsquo;s value is assumed to be
true.</p>
</td>
</tr>
<tr>
<td>
<code>subscriptionsLimit</code><br/>
<em>
int
</em>
</td>
<td>
<p>SubscriptionsLimit: Whether the number of subscriptions a user can have to this product at the same time. Set to null or
omit to allow unlimited per user subscriptions. Can be present only if subscriptionRequired property is present and has
a value of false.</p>
</td>
</tr>
<tr>
<td>
<code>terms</code><br/>
<em>
string
</em>
</td>
<td>
<p>Terms: Product terms of use. Developers trying to subscribe to the product will be presented and required to accept
these terms before they can complete the subscription process.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Product_STATUS_ARM">Service_Product_STATUS_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ProductContractProperties_STATUS_ARM">
ProductContractProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: Product entity contract properties.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Product_Spec">Service_Product_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Product">Product</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>approvalRequired</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ApprovalRequired: whether subscription approval is required. If false, new subscriptions will be approved automatically
enabling developers to call the product’s APIs immediately after subscribing. If true, administrators must manually
approve the subscription before the developer can any of the product’s APIs. Can be present only if
subscriptionRequired property is present and has a value of false.</p>
</td>
</tr>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: Product description. May include HTML formatting tags.</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: Product name.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a apimanagement.azure.com/Service resource</p>
</td>
</tr>
<tr>
<td>
<code>state</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ProductContractProperties_State">
ProductContractProperties_State
</a>
</em>
</td>
<td>
<p>State: whether product is published or not. Published products are discoverable by users of developer portal. Non
published products are visible only to administrators. Default state of Product is notPublished.</p>
</td>
</tr>
<tr>
<td>
<code>subscriptionRequired</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SubscriptionRequired: Whether a product subscription is required for accessing APIs included in this product. If true,
the product is referred to as &ldquo;protected&rdquo; and a valid subscription key is required for a request to an API included in
the product to succeed. If false, the product is referred to as &ldquo;open&rdquo; and requests to an API included in the product
can be made without a subscription key. If property is omitted when creating a new product it&rsquo;s value is assumed to be
true.</p>
</td>
</tr>
<tr>
<td>
<code>subscriptionsLimit</code><br/>
<em>
int
</em>
</td>
<td>
<p>SubscriptionsLimit: Whether the number of subscriptions a user can have to this product at the same time. Set to null or
omit to allow unlimited per user subscriptions. Can be present only if subscriptionRequired property is present and has
a value of false.</p>
</td>
</tr>
<tr>
<td>
<code>terms</code><br/>
<em>
string
</em>
</td>
<td>
<p>Terms: Product terms of use. Developers trying to subscribe to the product will be presented and required to accept
these terms before they can complete the subscription process.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Product_Spec_ARM">Service_Product_Spec_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ProductContractProperties_ARM">
ProductContractProperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: Product entity contract properties.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Products_Api_STATUS">Service_Products_Api_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ProductApi">ProductApi</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#Condition">
[]genruntime/conditions.Condition
</a>
</em>
</td>
<td>
<p>Conditions: The observed state of the resource</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Products_Api_STATUS_ARM">Service_Products_Api_STATUS_ARM
</h3>
<div>
</div>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Products_Api_Spec">Service_Products_Api_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ProductApi">ProductApi</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a apimanagement.azure.com/Product resource</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Products_Api_Spec_ARM">Service_Products_Api_Spec_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Products_Policy_STATUS">Service_Products_Policy_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ProductPolicy">ProductPolicy</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#Condition">
[]genruntime/conditions.Condition
</a>
</em>
</td>
<td>
<p>Conditions: The observed state of the resource</p>
</td>
</tr>
<tr>
<td>
<code>format</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.PolicyContractProperties_Format_STATUS">
PolicyContractProperties_Format_STATUS
</a>
</em>
</td>
<td>
<p>Format: Format of the policyContent.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Contents of the Policy as defined by the format.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Products_Policy_STATUS_ARM">Service_Products_Policy_STATUS_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.PolicyContractProperties_STATUS_ARM">
PolicyContractProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the Policy.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Products_Policy_Spec">Service_Products_Policy_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ProductPolicy">ProductPolicy</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>format</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.PolicyContractProperties_Format">
PolicyContractProperties_Format
</a>
</em>
</td>
<td>
<p>Format: Format of the policyContent.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a apimanagement.azure.com/Product resource</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Contents of the Policy as defined by the format.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Products_Policy_Spec_ARM">Service_Products_Policy_Spec_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.PolicyContractProperties_ARM">
PolicyContractProperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the Policy.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_STATUS">Service_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service">Service</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>additionalLocations</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AdditionalLocation_STATUS">
[]AdditionalLocation_STATUS
</a>
</em>
</td>
<td>
<p>AdditionalLocations: Additional datacenter locations of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>apiVersionConstraint</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionConstraint_STATUS">
ApiVersionConstraint_STATUS
</a>
</em>
</td>
<td>
<p>ApiVersionConstraint: Control Plane Apis version constraint for the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>certificates</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CertificateConfiguration_STATUS">
[]CertificateConfiguration_STATUS
</a>
</em>
</td>
<td>
<p>Certificates: List of Certificates that need to be installed in the API Management service. Max supported certificates
that can be installed is 10.</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#Condition">
[]genruntime/conditions.Condition
</a>
</em>
</td>
<td>
<p>Conditions: The observed state of the resource</p>
</td>
</tr>
<tr>
<td>
<code>configurationApi</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ConfigurationApi_STATUS">
ConfigurationApi_STATUS
</a>
</em>
</td>
<td>
<p>ConfigurationApi: Configuration API configuration of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>createdAtUtc</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedAtUtc: Creation UTC date of the API Management service.The date conforms to the following format:
<code>yyyy-MM-ddTHH:mm:ssZ</code> as specified by the ISO 8601 standard.</p>
</td>
</tr>
<tr>
<td>
<code>customProperties</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>CustomProperties: Custom properties of the API Management service.</br>Setting
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168</code> will disable the cipher
TLS_RSA_WITH_3DES_EDE_CBC_SHA for all TLS(1.0, 1.1 and 1.2).</br>Setting
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11</code> can be used to disable just TLS 1.1.</br>Setting
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10</code> can be used to disable TLS 1.0 on an API
Management service.</br>Setting <code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls11</code> can be
used to disable just TLS 1.1 for communications with backends.</br>Setting
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls10</code> can be used to disable TLS 1.0 for
communications with backends.</br>Setting <code>Microsoft.WindowsAzure.ApiManagement.Gateway.Protocols.Server.Http2</code> can be
used to enable HTTP2 protocol on an API Management service.</br>Not specifying any of these properties on PATCH
operation will reset omitted properties&rsquo; values to their defaults. For all the settings except Http2 the default value
is <code>True</code> if the service was created on or before April 1, 2018 and <code>False</code> otherwise. Http2 setting&rsquo;s default value is
<code>False</code>.</br></br>You can disable any of the following ciphers by using settings
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.[cipher_name]</code>: TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
TLS_RSA_WITH_AES_128_GCM_SHA256, TLS_RSA_WITH_AES_256_CBC_SHA256, TLS_RSA_WITH_AES_128_CBC_SHA256,
TLS_RSA_WITH_AES_256_CBC_SHA, TLS_RSA_WITH_AES_128_CBC_SHA. For example,
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TLS_RSA_WITH_AES_128_CBC_SHA256</code>:<code>false</code>. The default
value is <code>true</code> for them.</br> Note: The following ciphers can&rsquo;t be disabled since they are required by internal
platform components:
TLS_AES_256_GCM_SHA384,TLS_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384,TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384,TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256</p>
</td>
</tr>
<tr>
<td>
<code>developerPortalStatus</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_DeveloperPortalStatus_STATUS">
ApiManagementServiceProperties_DeveloperPortalStatus_STATUS
</a>
</em>
</td>
<td>
<p>DeveloperPortalStatus: Status of developer portal in this API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>developerPortalUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>DeveloperPortalUrl: DEveloper Portal endpoint URL of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>disableGateway</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableGateway: Property only valid for an Api Management service deployed in multiple locations. This can be used to
disable the gateway in master region.</p>
</td>
</tr>
<tr>
<td>
<code>enableClientCertificate</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableClientCertificate: Property only meant to be used for Consumption SKU Service. This enforces a client certificate
to be presented on each request to the gateway. This also enables the ability to authenticate the certificate in the
policy on the gateway.</p>
</td>
</tr>
<tr>
<td>
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: ETag of the resource.</p>
</td>
</tr>
<tr>
<td>
<code>gatewayRegionalUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>GatewayRegionalUrl: Gateway URL of the API Management service in the Default Region.</p>
</td>
</tr>
<tr>
<td>
<code>gatewayUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>GatewayUrl: Gateway URL of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>hostnameConfigurations</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration_STATUS">
[]HostnameConfiguration_STATUS
</a>
</em>
</td>
<td>
<p>HostnameConfigurations: Custom hostname configuration of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Resource ID.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceIdentity_STATUS">
ApiManagementServiceIdentity_STATUS
</a>
</em>
</td>
<td>
<p>Identity: Managed service identity of the Api Management service.</p>
</td>
</tr>
<tr>
<td>
<code>legacyPortalStatus</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_LegacyPortalStatus_STATUS">
ApiManagementServiceProperties_LegacyPortalStatus_STATUS
</a>
</em>
</td>
<td>
<p>LegacyPortalStatus: Status of legacy portal in the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: Resource location.</p>
</td>
</tr>
<tr>
<td>
<code>managementApiUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>ManagementApiUrl: Management API endpoint URL of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: Resource name.</p>
</td>
</tr>
<tr>
<td>
<code>natGatewayState</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_NatGatewayState_STATUS">
ApiManagementServiceProperties_NatGatewayState_STATUS
</a>
</em>
</td>
<td>
<p>NatGatewayState: Property can be used to enable NAT Gateway for this API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>notificationSenderEmail</code><br/>
<em>
string
</em>
</td>
<td>
<p>NotificationSenderEmail: Email address from which the notification will be sent.</p>
</td>
</tr>
<tr>
<td>
<code>outboundPublicIPAddresses</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>OutboundPublicIPAddresses: Outbound public IPV4 address prefixes associated with NAT Gateway deployed service. Available
only for Premium SKU on stv2 platform.</p>
</td>
</tr>
<tr>
<td>
<code>platformVersion</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_PlatformVersion_STATUS">
ApiManagementServiceProperties_PlatformVersion_STATUS
</a>
</em>
</td>
<td>
<p>PlatformVersion: Compute Platform Version running the service in this location.</p>
</td>
</tr>
<tr>
<td>
<code>portalUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>PortalUrl: Publisher portal endpoint Url of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpointConnections</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.RemotePrivateEndpointConnectionWrapper_STATUS">
[]RemotePrivateEndpointConnectionWrapper_STATUS
</a>
</em>
</td>
<td>
<p>PrivateEndpointConnections: List of Private Endpoint Connections of this service.</p>
</td>
</tr>
<tr>
<td>
<code>privateIPAddresses</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>PrivateIPAddresses: Private Static Load Balanced IP addresses of the API Management service in Primary region which is
deployed in an Internal Virtual Network. Available only for Basic, Standard, Premium and Isolated SKU.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
string
</em>
</td>
<td>
<p>ProvisioningState: The current provisioning state of the API Management service which can be one of the following:
Created/Activating/Succeeded/Updating/Failed/Stopped/Terminating/TerminationFailed/Deleted.</p>
</td>
</tr>
<tr>
<td>
<code>publicIPAddresses</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>PublicIPAddresses: Public Static Load Balanced IP addresses of the API Management service in Primary region. Available
only for Basic, Standard, Premium and Isolated SKU.</p>
</td>
</tr>
<tr>
<td>
<code>publicIpAddressId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PublicIpAddressId: Public Standard SKU IP V4 based IP address to be associated with Virtual Network deployed service in
the region. Supported only for Developer and Premium SKU being deployed in Virtual Network.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_PublicNetworkAccess_STATUS">
ApiManagementServiceProperties_PublicNetworkAccess_STATUS
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Whether or not public endpoint access is allowed for this API Management service.  Value is
optional but if passed in, must be &lsquo;Enabled&rsquo; or &lsquo;Disabled&rsquo;. If &lsquo;Disabled&rsquo;, private endpoints are the exclusive access
method. Default value is &lsquo;Enabled&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>publisherEmail</code><br/>
<em>
string
</em>
</td>
<td>
<p>PublisherEmail: Publisher email.</p>
</td>
</tr>
<tr>
<td>
<code>publisherName</code><br/>
<em>
string
</em>
</td>
<td>
<p>PublisherName: Publisher name.</p>
</td>
</tr>
<tr>
<td>
<code>restore</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Restore: Undelete Api Management Service if it was previously soft-deleted. If this flag is specified and set to True
all other properties will be ignored.</p>
</td>
</tr>
<tr>
<td>
<code>scmUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScmUrl: SCM endpoint URL of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceSkuProperties_STATUS">
ApiManagementServiceSkuProperties_STATUS
</a>
</em>
</td>
<td>
<p>Sku: SKU properties of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.SystemData_STATUS">
SystemData_STATUS
</a>
</em>
</td>
<td>
<p>SystemData: Metadata pertaining to creation and last modification of the resource.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Resource tags.</p>
</td>
</tr>
<tr>
<td>
<code>targetProvisioningState</code><br/>
<em>
string
</em>
</td>
<td>
<p>TargetProvisioningState: The provisioning state of the API Management service, which is targeted by the long running
operation started on the service.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: Resource type for API Management resource is set to Microsoft.ApiManagement.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkConfiguration</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.VirtualNetworkConfiguration_STATUS">
VirtualNetworkConfiguration_STATUS
</a>
</em>
</td>
<td>
<p>VirtualNetworkConfiguration: Virtual network configuration of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkType</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_VirtualNetworkType_STATUS">
ApiManagementServiceProperties_VirtualNetworkType_STATUS
</a>
</em>
</td>
<td>
<p>VirtualNetworkType: The type of VPN in which API Management service needs to be configured in. None (Default Value)
means the API Management service is not part of any Virtual Network, External means the API Management deployment is set
up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is
setup inside a Virtual Network having an Intranet Facing Endpoint only.</p>
</td>
</tr>
<tr>
<td>
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: A list of availability zones denoting where the resource needs to come from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_STATUS_ARM">Service_STATUS_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: ETag of the resource.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Resource ID.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceIdentity_STATUS_ARM">
ApiManagementServiceIdentity_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Identity: Managed service identity of the Api Management service.</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: Resource location.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: Resource name.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_STATUS_ARM">
ApiManagementServiceProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceSkuProperties_STATUS_ARM">
ApiManagementServiceSkuProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Sku: SKU properties of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.SystemData_STATUS_ARM">
SystemData_STATUS_ARM
</a>
</em>
</td>
<td>
<p>SystemData: Metadata pertaining to creation and last modification of the resource.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Resource tags.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: Resource type for API Management resource is set to Microsoft.ApiManagement.</p>
</td>
</tr>
<tr>
<td>
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: A list of availability zones denoting where the resource needs to come from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Spec">Service_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service">Service</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>additionalLocations</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.AdditionalLocation">
[]AdditionalLocation
</a>
</em>
</td>
<td>
<p>AdditionalLocations: Additional datacenter locations of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>apiVersionConstraint</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiVersionConstraint">
ApiVersionConstraint
</a>
</em>
</td>
<td>
<p>ApiVersionConstraint: Control Plane Apis version constraint for the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>certificates</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.CertificateConfiguration">
[]CertificateConfiguration
</a>
</em>
</td>
<td>
<p>Certificates: List of Certificates that need to be installed in the API Management service. Max supported certificates
that can be installed is 10.</p>
</td>
</tr>
<tr>
<td>
<code>configurationApi</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ConfigurationApi">
ConfigurationApi
</a>
</em>
</td>
<td>
<p>ConfigurationApi: Configuration API configuration of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>customProperties</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>CustomProperties: Custom properties of the API Management service.</br>Setting
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168</code> will disable the cipher
TLS_RSA_WITH_3DES_EDE_CBC_SHA for all TLS(1.0, 1.1 and 1.2).</br>Setting
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11</code> can be used to disable just TLS 1.1.</br>Setting
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10</code> can be used to disable TLS 1.0 on an API
Management service.</br>Setting <code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls11</code> can be
used to disable just TLS 1.1 for communications with backends.</br>Setting
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls10</code> can be used to disable TLS 1.0 for
communications with backends.</br>Setting <code>Microsoft.WindowsAzure.ApiManagement.Gateway.Protocols.Server.Http2</code> can be
used to enable HTTP2 protocol on an API Management service.</br>Not specifying any of these properties on PATCH
operation will reset omitted properties&rsquo; values to their defaults. For all the settings except Http2 the default value
is <code>True</code> if the service was created on or before April 1, 2018 and <code>False</code> otherwise. Http2 setting&rsquo;s default value is
<code>False</code>.</br></br>You can disable any of the following ciphers by using settings
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.[cipher_name]</code>: TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
TLS_RSA_WITH_AES_128_GCM_SHA256, TLS_RSA_WITH_AES_256_CBC_SHA256, TLS_RSA_WITH_AES_128_CBC_SHA256,
TLS_RSA_WITH_AES_256_CBC_SHA, TLS_RSA_WITH_AES_128_CBC_SHA. For example,
<code>Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TLS_RSA_WITH_AES_128_CBC_SHA256</code>:<code>false</code>. The default
value is <code>true</code> for them.</br> Note: The following ciphers can&rsquo;t be disabled since they are required by internal
platform components:
TLS_AES_256_GCM_SHA384,TLS_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384,TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384,TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256</p>
</td>
</tr>
<tr>
<td>
<code>developerPortalStatus</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_DeveloperPortalStatus">
ApiManagementServiceProperties_DeveloperPortalStatus
</a>
</em>
</td>
<td>
<p>DeveloperPortalStatus: Status of developer portal in this API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>disableGateway</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableGateway: Property only valid for an Api Management service deployed in multiple locations. This can be used to
disable the gateway in master region.</p>
</td>
</tr>
<tr>
<td>
<code>enableClientCertificate</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableClientCertificate: Property only meant to be used for Consumption SKU Service. This enforces a client certificate
to be presented on each request to the gateway. This also enables the ability to authenticate the certificate in the
policy on the gateway.</p>
</td>
</tr>
<tr>
<td>
<code>hostnameConfigurations</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.HostnameConfiguration">
[]HostnameConfiguration
</a>
</em>
</td>
<td>
<p>HostnameConfigurations: Custom hostname configuration of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceIdentity">
ApiManagementServiceIdentity
</a>
</em>
</td>
<td>
<p>Identity: Managed service identity of the Api Management service.</p>
</td>
</tr>
<tr>
<td>
<code>legacyPortalStatus</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_LegacyPortalStatus">
ApiManagementServiceProperties_LegacyPortalStatus
</a>
</em>
</td>
<td>
<p>LegacyPortalStatus: Status of legacy portal in the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: Resource location.</p>
</td>
</tr>
<tr>
<td>
<code>natGatewayState</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_NatGatewayState">
ApiManagementServiceProperties_NatGatewayState
</a>
</em>
</td>
<td>
<p>NatGatewayState: Property can be used to enable NAT Gateway for this API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>notificationSenderEmail</code><br/>
<em>
string
</em>
</td>
<td>
<p>NotificationSenderEmail: Email address from which the notification will be sent.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a resources.azure.com/ResourceGroup resource</p>
</td>
</tr>
<tr>
<td>
<code>publicIpAddressReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>PublicIpAddressReference: Public Standard SKU IP V4 based IP address to be associated with Virtual Network deployed
service in the region. Supported only for Developer and Premium SKU being deployed in Virtual Network.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_PublicNetworkAccess">
ApiManagementServiceProperties_PublicNetworkAccess
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Whether or not public endpoint access is allowed for this API Management service.  Value is
optional but if passed in, must be &lsquo;Enabled&rsquo; or &lsquo;Disabled&rsquo;. If &lsquo;Disabled&rsquo;, private endpoints are the exclusive access
method. Default value is &lsquo;Enabled&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>publisherEmail</code><br/>
<em>
string
</em>
</td>
<td>
<p>PublisherEmail: Publisher email.</p>
</td>
</tr>
<tr>
<td>
<code>publisherName</code><br/>
<em>
string
</em>
</td>
<td>
<p>PublisherName: Publisher name.</p>
</td>
</tr>
<tr>
<td>
<code>restore</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Restore: Undelete Api Management Service if it was previously soft-deleted. If this flag is specified and set to True
all other properties will be ignored.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceSkuProperties">
ApiManagementServiceSkuProperties
</a>
</em>
</td>
<td>
<p>Sku: SKU properties of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Resource tags.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkConfiguration</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.VirtualNetworkConfiguration">
VirtualNetworkConfiguration
</a>
</em>
</td>
<td>
<p>VirtualNetworkConfiguration: Virtual network configuration of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkType</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_VirtualNetworkType">
ApiManagementServiceProperties_VirtualNetworkType
</a>
</em>
</td>
<td>
<p>VirtualNetworkType: The type of VPN in which API Management service needs to be configured in. None (Default Value)
means the API Management service is not part of any Virtual Network, External means the API Management deployment is set
up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is
setup inside a Virtual Network having an Intranet Facing Endpoint only.</p>
</td>
</tr>
<tr>
<td>
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: A list of availability zones denoting where the resource needs to come from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Spec_ARM">Service_Spec_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceIdentity_ARM">
ApiManagementServiceIdentity_ARM
</a>
</em>
</td>
<td>
<p>Identity: Managed service identity of the Api Management service.</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: Resource location.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_ARM">
ApiManagementServiceProperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceSkuProperties_ARM">
ApiManagementServiceSkuProperties_ARM
</a>
</em>
</td>
<td>
<p>Sku: SKU properties of the API Management service.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Resource tags.</p>
</td>
</tr>
<tr>
<td>
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: A list of availability zones denoting where the resource needs to come from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Subscription_STATUS">Service_Subscription_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Subscription">Subscription</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allowTracing</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowTracing: Determines whether tracing is enabled</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#Condition">
[]genruntime/conditions.Condition
</a>
</em>
</td>
<td>
<p>Conditions: The observed state of the resource</p>
</td>
</tr>
<tr>
<td>
<code>createdDate</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedDate: Subscription creation date. The date conforms to the following format: <code>yyyy-MM-ddTHH:mm:ssZ</code> as specified
by the ISO 8601 standard.</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: The name of the subscription, or null if the subscription has no name.</p>
</td>
</tr>
<tr>
<td>
<code>endDate</code><br/>
<em>
string
</em>
</td>
<td>
<p>EndDate: Date when subscription was cancelled or expired. The setting is for audit purposes only and the subscription is
not automatically cancelled. The subscription lifecycle can be managed by using the <code>state</code> property. The date conforms
to the following format: <code>yyyy-MM-ddTHH:mm:ssZ</code> as specified by the ISO 8601 standard.</p>
</td>
</tr>
<tr>
<td>
<code>expirationDate</code><br/>
<em>
string
</em>
</td>
<td>
<p>ExpirationDate: Subscription expiration date. The setting is for audit purposes only and the subscription is not
automatically expired. The subscription lifecycle can be managed by using the <code>state</code> property. The date conforms to the
following format: <code>yyyy-MM-ddTHH:mm:ssZ</code> as specified by the ISO 8601 standard.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>notificationDate</code><br/>
<em>
string
</em>
</td>
<td>
<p>NotificationDate: Upcoming subscription expiration notification date. The date conforms to the following format:
<code>yyyy-MM-ddTHH:mm:ssZ</code> as specified by the ISO 8601 standard.</p>
</td>
</tr>
<tr>
<td>
<code>ownerId</code><br/>
<em>
string
</em>
</td>
<td>
<p>OwnerId: The user resource identifier of the subscription owner. The value is a valid relative URL in the format of
/users/{userId} where {userId} is a user identifier.</p>
</td>
</tr>
<tr>
<td>
<code>scope</code><br/>
<em>
string
</em>
</td>
<td>
<p>Scope: Scope like /products/{productId} or /apis or /apis/{apiId}.</p>
</td>
</tr>
<tr>
<td>
<code>startDate</code><br/>
<em>
string
</em>
</td>
<td>
<p>StartDate: Subscription activation date. The setting is for audit purposes only and the subscription is not
automatically activated. The subscription lifecycle can be managed by using the <code>state</code> property. The date conforms to
the following format: <code>yyyy-MM-ddTHH:mm:ssZ</code> as specified by the ISO 8601 standard.</p>
</td>
</tr>
<tr>
<td>
<code>state</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.SubscriptionContractProperties_State_STATUS">
SubscriptionContractProperties_State_STATUS
</a>
</em>
</td>
<td>
<p>State: Subscription state. Possible states are * active – the subscription is active, * suspended – the subscription
is blocked, and the subscriber cannot call any APIs of the product, * submitted – the subscription request has been
made by the developer, but has not yet been approved or rejected, * rejected – the subscription request has been
denied by an administrator, * cancelled – the subscription has been cancelled by the developer or administrator, *
expired – the subscription reached its expiration date and was deactivated.</p>
</td>
</tr>
<tr>
<td>
<code>stateComment</code><br/>
<em>
string
</em>
</td>
<td>
<p>StateComment: Optional subscription comment added by an administrator when the state is changed to the &lsquo;rejected&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Subscription_STATUS_ARM">Service_Subscription_STATUS_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.SubscriptionContractProperties_STATUS_ARM">
SubscriptionContractProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: Subscription contract properties.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Subscription_Spec">Service_Subscription_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Subscription">Subscription</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allowTracing</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowTracing: Determines whether tracing can be enabled</p>
</td>
</tr>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: Subscription name.</p>
</td>
</tr>
<tr>
<td>
<code>operatorSpec</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.SubscriptionOperatorSpec">
SubscriptionOperatorSpec
</a>
</em>
</td>
<td>
<p>OperatorSpec: The specification for configuring operator behavior. This field is interpreted by the operator and not
passed directly to Azure</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a apimanagement.azure.com/Service resource</p>
</td>
</tr>
<tr>
<td>
<code>ownerReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>OwnerReference: User (user id path) for whom subscription is being created in form /users/{userId}</p>
</td>
</tr>
<tr>
<td>
<code>primaryKey</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretReference">
genruntime.SecretReference
</a>
</em>
</td>
<td>
<p>PrimaryKey: Primary subscription key. If not specified during request key will be generated automatically.</p>
</td>
</tr>
<tr>
<td>
<code>scope</code><br/>
<em>
string
</em>
</td>
<td>
<p>Scope: Scope like /products/{productId} or /apis or /apis/{apiId}.</p>
</td>
</tr>
<tr>
<td>
<code>secondaryKey</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretReference">
genruntime.SecretReference
</a>
</em>
</td>
<td>
<p>SecondaryKey: Secondary subscription key. If not specified during request key will be generated automatically.</p>
</td>
</tr>
<tr>
<td>
<code>state</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.SubscriptionCreateParameterProperties_State">
SubscriptionCreateParameterProperties_State
</a>
</em>
</td>
<td>
<p>State: Initial subscription state. If no value is specified, subscription is created with Submitted state. Possible
states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber
cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has
not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, *
cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription
reached its expiration date and was deactivated.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Service_Subscription_Spec_ARM">Service_Subscription_Spec_ARM
</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.SubscriptionCreateParameterProperties_ARM">
SubscriptionCreateParameterProperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: Subscription contract properties.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.Subscription">Subscription
</h3>
<div>
<p>Generator information:
- Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimsubscriptions.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ApiManagement/&#x200b;service/&#x200b;{serviceName}/&#x200b;subscriptions/&#x200b;{sid}</&#x200b;p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_Subscription_Spec">
Service_Subscription_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>allowTracing</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowTracing: Determines whether tracing can be enabled</p>
</td>
</tr>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: Subscription name.</p>
</td>
</tr>
<tr>
<td>
<code>operatorSpec</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.SubscriptionOperatorSpec">
SubscriptionOperatorSpec
</a>
</em>
</td>
<td>
<p>OperatorSpec: The specification for configuring operator behavior. This field is interpreted by the operator and not
passed directly to Azure</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a apimanagement.azure.com/Service resource</p>
</td>
</tr>
<tr>
<td>
<code>ownerReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>OwnerReference: User (user id path) for whom subscription is being created in form /users/{userId}</p>
</td>
</tr>
<tr>
<td>
<code>primaryKey</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretReference">
genruntime.SecretReference
</a>
</em>
</td>
<td>
<p>PrimaryKey: Primary subscription key. If not specified during request key will be generated automatically.</p>
</td>
</tr>
<tr>
<td>
<code>scope</code><br/>
<em>
string
</em>
</td>
<td>
<p>Scope: Scope like /products/{productId} or /apis or /apis/{apiId}.</p>
</td>
</tr>
<tr>
<td>
<code>secondaryKey</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretReference">
genruntime.SecretReference
</a>
</em>
</td>
<td>
<p>SecondaryKey: Secondary subscription key. If not specified during request key will be generated automatically.</p>
</td>
</tr>
<tr>
<td>
<code>state</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.SubscriptionCreateParameterProperties_State">
SubscriptionCreateParameterProperties_State
</a>
</em>
</td>
<td>
<p>State: Initial subscription state. If no value is specified, subscription is created with Submitted state. Possible
states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber
cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has
not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, *
cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription
reached its expiration date and was deactivated.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.Service_Subscription_STATUS">
Service_Subscription_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.SubscriptionContractProperties_STATUS_ARM">SubscriptionContractProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Subscription_STATUS_ARM">Service_Subscription_STATUS_ARM</a>)
</p>
<div>
<p>Subscription details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allowTracing</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowTracing: Determines whether tracing is enabled</p>
</td>
</tr>
<tr>
<td>
<code>createdDate</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedDate: Subscription creation date. The date conforms to the following format: <code>yyyy-MM-ddTHH:mm:ssZ</code> as specified
by the ISO 8601 standard.</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: The name of the subscription, or null if the subscription has no name.</p>
</td>
</tr>
<tr>
<td>
<code>endDate</code><br/>
<em>
string
</em>
</td>
<td>
<p>EndDate: Date when subscription was cancelled or expired. The setting is for audit purposes only and the subscription is
not automatically cancelled. The subscription lifecycle can be managed by using the <code>state</code> property. The date conforms
to the following format: <code>yyyy-MM-ddTHH:mm:ssZ</code> as specified by the ISO 8601 standard.</p>
</td>
</tr>
<tr>
<td>
<code>expirationDate</code><br/>
<em>
string
</em>
</td>
<td>
<p>ExpirationDate: Subscription expiration date. The setting is for audit purposes only and the subscription is not
automatically expired. The subscription lifecycle can be managed by using the <code>state</code> property. The date conforms to the
following format: <code>yyyy-MM-ddTHH:mm:ssZ</code> as specified by the ISO 8601 standard.</p>
</td>
</tr>
<tr>
<td>
<code>notificationDate</code><br/>
<em>
string
</em>
</td>
<td>
<p>NotificationDate: Upcoming subscription expiration notification date. The date conforms to the following format:
<code>yyyy-MM-ddTHH:mm:ssZ</code> as specified by the ISO 8601 standard.</p>
</td>
</tr>
<tr>
<td>
<code>ownerId</code><br/>
<em>
string
</em>
</td>
<td>
<p>OwnerId: The user resource identifier of the subscription owner. The value is a valid relative URL in the format of
/users/{userId} where {userId} is a user identifier.</p>
</td>
</tr>
<tr>
<td>
<code>scope</code><br/>
<em>
string
</em>
</td>
<td>
<p>Scope: Scope like /products/{productId} or /apis or /apis/{apiId}.</p>
</td>
</tr>
<tr>
<td>
<code>startDate</code><br/>
<em>
string
</em>
</td>
<td>
<p>StartDate: Subscription activation date. The setting is for audit purposes only and the subscription is not
automatically activated. The subscription lifecycle can be managed by using the <code>state</code> property. The date conforms to
the following format: <code>yyyy-MM-ddTHH:mm:ssZ</code> as specified by the ISO 8601 standard.</p>
</td>
</tr>
<tr>
<td>
<code>state</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.SubscriptionContractProperties_State_STATUS">
SubscriptionContractProperties_State_STATUS
</a>
</em>
</td>
<td>
<p>State: Subscription state. Possible states are * active – the subscription is active, * suspended – the subscription
is blocked, and the subscriber cannot call any APIs of the product, * submitted – the subscription request has been
made by the developer, but has not yet been approved or rejected, * rejected – the subscription request has been
denied by an administrator, * cancelled – the subscription has been cancelled by the developer or administrator, *
expired – the subscription reached its expiration date and was deactivated.</p>
</td>
</tr>
<tr>
<td>
<code>stateComment</code><br/>
<em>
string
</em>
</td>
<td>
<p>StateComment: Optional subscription comment added by an administrator when the state is changed to the &lsquo;rejected&rsquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.SubscriptionContractProperties_State_STATUS">SubscriptionContractProperties_State_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Subscription_STATUS">Service_Subscription_STATUS</a>, <a href="#apimanagement.azure.com/v1api20230501preview.SubscriptionContractProperties_STATUS_ARM">SubscriptionContractProperties_STATUS_ARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;active&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;cancelled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;expired&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;rejected&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;submitted&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;suspended&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.SubscriptionCreateParameterProperties_ARM">SubscriptionCreateParameterProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Subscription_Spec_ARM">Service_Subscription_Spec_ARM</a>)
</p>
<div>
<p>Parameters supplied to the Create subscription operation.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allowTracing</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowTracing: Determines whether tracing can be enabled</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: Subscription name.</p>
</td>
</tr>
<tr>
<td>
<code>ownerId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>primaryKey</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrimaryKey: Primary subscription key. If not specified during request key will be generated automatically.</p>
</td>
</tr>
<tr>
<td>
<code>scope</code><br/>
<em>
string
</em>
</td>
<td>
<p>Scope: Scope like /products/{productId} or /apis or /apis/{apiId}.</p>
</td>
</tr>
<tr>
<td>
<code>secondaryKey</code><br/>
<em>
string
</em>
</td>
<td>
<p>SecondaryKey: Secondary subscription key. If not specified during request key will be generated automatically.</p>
</td>
</tr>
<tr>
<td>
<code>state</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.SubscriptionCreateParameterProperties_State">
SubscriptionCreateParameterProperties_State
</a>
</em>
</td>
<td>
<p>State: Initial subscription state. If no value is specified, subscription is created with Submitted state. Possible
states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber
cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has
not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, *
cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription
reached its expiration date and was deactivated.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.SubscriptionCreateParameterProperties_State">SubscriptionCreateParameterProperties_State
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Subscription_Spec">Service_Subscription_Spec</a>, <a href="#apimanagement.azure.com/v1api20230501preview.SubscriptionCreateParameterProperties_ARM">SubscriptionCreateParameterProperties_ARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;active&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;cancelled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;expired&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;rejected&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;submitted&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;suspended&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.SubscriptionKeyParameterNamesContract">SubscriptionKeyParameterNamesContract
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Api_Spec">Service_Api_Spec</a>)
</p>
<div>
<p>Subscription key parameter names details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>header</code><br/>
<em>
string
</em>
</td>
<td>
<p>Header: Subscription key header name.</p>
</td>
</tr>
<tr>
<td>
<code>query</code><br/>
<em>
string
</em>
</td>
<td>
<p>Query: Subscription key query string parameter name.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.SubscriptionKeyParameterNamesContract_ARM">SubscriptionKeyParameterNamesContract_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiCreateOrUpdateProperties_ARM">ApiCreateOrUpdateProperties_ARM</a>)
</p>
<div>
<p>Subscription key parameter names details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>header</code><br/>
<em>
string
</em>
</td>
<td>
<p>Header: Subscription key header name.</p>
</td>
</tr>
<tr>
<td>
<code>query</code><br/>
<em>
string
</em>
</td>
<td>
<p>Query: Subscription key query string parameter name.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.SubscriptionKeyParameterNamesContract_STATUS">SubscriptionKeyParameterNamesContract_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Api_STATUS">Service_Api_STATUS</a>)
</p>
<div>
<p>Subscription key parameter names details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>header</code><br/>
<em>
string
</em>
</td>
<td>
<p>Header: Subscription key header name.</p>
</td>
</tr>
<tr>
<td>
<code>query</code><br/>
<em>
string
</em>
</td>
<td>
<p>Query: Subscription key query string parameter name.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.SubscriptionKeyParameterNamesContract_STATUS_ARM">SubscriptionKeyParameterNamesContract_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiContractProperties_STATUS_ARM">ApiContractProperties_STATUS_ARM</a>)
</p>
<div>
<p>Subscription key parameter names details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>header</code><br/>
<em>
string
</em>
</td>
<td>
<p>Header: Subscription key header name.</p>
</td>
</tr>
<tr>
<td>
<code>query</code><br/>
<em>
string
</em>
</td>
<td>
<p>Query: Subscription key query string parameter name.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.SubscriptionOperatorSecrets">SubscriptionOperatorSecrets
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.SubscriptionOperatorSpec">SubscriptionOperatorSpec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>primaryKey</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretDestination">
genruntime.SecretDestination
</a>
</em>
</td>
<td>
<p>PrimaryKey: indicates where the PrimaryKey secret should be placed. If omitted, the secret will not be retrieved from
Azure.</p>
</td>
</tr>
<tr>
<td>
<code>secondaryKey</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretDestination">
genruntime.SecretDestination
</a>
</em>
</td>
<td>
<p>SecondaryKey: indicates where the SecondaryKey secret should be placed. If omitted, the secret will not be retrieved
from Azure.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.SubscriptionOperatorSpec">SubscriptionOperatorSpec
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_Subscription_Spec">Service_Subscription_Spec</a>)
</p>
<div>
<p>Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>secrets</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.SubscriptionOperatorSecrets">
SubscriptionOperatorSecrets
</a>
</em>
</td>
<td>
<p>Secrets: configures where to place Azure generated secrets.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.SystemData_CreatedByType_STATUS">SystemData_CreatedByType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.SystemData_STATUS">SystemData_STATUS</a>, <a href="#apimanagement.azure.com/v1api20230501preview.SystemData_STATUS_ARM">SystemData_STATUS_ARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Application&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Key&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ManagedIdentity&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;User&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.SystemData_LastModifiedByType_STATUS">SystemData_LastModifiedByType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.SystemData_STATUS">SystemData_STATUS</a>, <a href="#apimanagement.azure.com/v1api20230501preview.SystemData_STATUS_ARM">SystemData_STATUS_ARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Application&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Key&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ManagedIdentity&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;User&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.SystemData_STATUS">SystemData_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_STATUS">Service_STATUS</a>)
</p>
<div>
<p>Metadata pertaining to creation and last modification of the resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>createdAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedAt: The timestamp of resource creation (UTC).</p>
</td>
</tr>
<tr>
<td>
<code>createdBy</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedBy: The identity that created the resource.</p>
</td>
</tr>
<tr>
<td>
<code>createdByType</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.SystemData_CreatedByType_STATUS">
SystemData_CreatedByType_STATUS
</a>
</em>
</td>
<td>
<p>CreatedByType: The type of identity that created the resource.</p>
</td>
</tr>
<tr>
<td>
<code>lastModifiedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>LastModifiedAt: The timestamp of resource last modification (UTC)</p>
</td>
</tr>
<tr>
<td>
<code>lastModifiedBy</code><br/>
<em>
string
</em>
</td>
<td>
<p>LastModifiedBy: The identity that last modified the resource.</p>
</td>
</tr>
<tr>
<td>
<code>lastModifiedByType</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.SystemData_LastModifiedByType_STATUS">
SystemData_LastModifiedByType_STATUS
</a>
</em>
</td>
<td>
<p>LastModifiedByType: The type of identity that last modified the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.SystemData_STATUS_ARM">SystemData_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.Service_STATUS_ARM">Service_STATUS_ARM</a>)
</p>
<div>
<p>Metadata pertaining to creation and last modification of the resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>createdAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedAt: The timestamp of resource creation (UTC).</p>
</td>
</tr>
<tr>
<td>
<code>createdBy</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedBy: The identity that created the resource.</p>
</td>
</tr>
<tr>
<td>
<code>createdByType</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.SystemData_CreatedByType_STATUS">
SystemData_CreatedByType_STATUS
</a>
</em>
</td>
<td>
<p>CreatedByType: The type of identity that created the resource.</p>
</td>
</tr>
<tr>
<td>
<code>lastModifiedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>LastModifiedAt: The timestamp of resource last modification (UTC)</p>
</td>
</tr>
<tr>
<td>
<code>lastModifiedBy</code><br/>
<em>
string
</em>
</td>
<td>
<p>LastModifiedBy: The identity that last modified the resource.</p>
</td>
</tr>
<tr>
<td>
<code>lastModifiedByType</code><br/>
<em>
<a href="#apimanagement.azure.com/v1api20230501preview.SystemData_LastModifiedByType_STATUS">
SystemData_LastModifiedByType_STATUS
</a>
</em>
</td>
<td>
<p>LastModifiedByType: The type of identity that last modified the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.UserAssignedIdentityDetails">UserAssignedIdentityDetails
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceIdentity">ApiManagementServiceIdentity</a>)
</p>
<div>
<p>Information about the user assigned identity for the resource</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>reference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.UserAssignedIdentityDetails_ARM">UserAssignedIdentityDetails_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceIdentity_ARM">ApiManagementServiceIdentity_ARM</a>)
</p>
<div>
<p>Information about the user assigned identity for the resource</p>
</div>
<h3 id="apimanagement.azure.com/v1api20230501preview.UserIdentityProperties_STATUS">UserIdentityProperties_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceIdentity_STATUS">ApiManagementServiceIdentity_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>clientId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClientId: The client id of user assigned identity.</p>
</td>
</tr>
<tr>
<td>
<code>principalId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrincipalId: The principal id of user assigned identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.UserIdentityProperties_STATUS_ARM">UserIdentityProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceIdentity_STATUS_ARM">ApiManagementServiceIdentity_STATUS_ARM</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>clientId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClientId: The client id of user assigned identity.</p>
</td>
</tr>
<tr>
<td>
<code>principalId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrincipalId: The principal id of user assigned identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.VirtualNetworkConfiguration">VirtualNetworkConfiguration
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AdditionalLocation">AdditionalLocation</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_Spec">Service_Spec</a>)
</p>
<div>
<p>Configuration of a virtual network to which API Management service is deployed.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>subnetResourceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>SubnetResourceReference: The full resource ID of a subnet in a virtual network to deploy the API Management service in.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.VirtualNetworkConfiguration_ARM">VirtualNetworkConfiguration_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AdditionalLocation_ARM">AdditionalLocation_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_ARM">ApiManagementServiceProperties_ARM</a>)
</p>
<div>
<p>Configuration of a virtual network to which API Management service is deployed.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>subnetResourceId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.VirtualNetworkConfiguration_STATUS">VirtualNetworkConfiguration_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AdditionalLocation_STATUS">AdditionalLocation_STATUS</a>, <a href="#apimanagement.azure.com/v1api20230501preview.Service_STATUS">Service_STATUS</a>)
</p>
<div>
<p>Configuration of a virtual network to which API Management service is deployed.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>subnetResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>SubnetResourceId: The full resource ID of a subnet in a virtual network to deploy the API Management service in.</p>
</td>
</tr>
<tr>
<td>
<code>subnetname</code><br/>
<em>
string
</em>
</td>
<td>
<p>Subnetname: The name of the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>vnetid</code><br/>
<em>
string
</em>
</td>
<td>
<p>Vnetid: The virtual network ID. This is typically a GUID. Expect a null GUID by default.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.VirtualNetworkConfiguration_STATUS_ARM">VirtualNetworkConfiguration_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.AdditionalLocation_STATUS_ARM">AdditionalLocation_STATUS_ARM</a>, <a href="#apimanagement.azure.com/v1api20230501preview.ApiManagementServiceProperties_STATUS_ARM">ApiManagementServiceProperties_STATUS_ARM</a>)
</p>
<div>
<p>Configuration of a virtual network to which API Management service is deployed.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>subnetResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>SubnetResourceId: The full resource ID of a subnet in a virtual network to deploy the API Management service in.</p>
</td>
</tr>
<tr>
<td>
<code>subnetname</code><br/>
<em>
string
</em>
</td>
<td>
<p>Subnetname: The name of the subnet.</p>
</td>
</tr>
<tr>
<td>
<code>vnetid</code><br/>
<em>
string
</em>
</td>
<td>
<p>Vnetid: The virtual network ID. This is typically a GUID. Expect a null GUID by default.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.X509CertificateName">X509CertificateName
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendServiceFabricClusterProperties">BackendServiceFabricClusterProperties</a>)
</p>
<div>
<p>Properties of server X509Names.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>issuerCertificateThumbprint</code><br/>
<em>
string
</em>
</td>
<td>
<p>IssuerCertificateThumbprint: Thumbprint for the Issuer of the Certificate.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: Common Name of the Certificate.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.X509CertificateName_ARM">X509CertificateName_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendServiceFabricClusterProperties_ARM">BackendServiceFabricClusterProperties_ARM</a>)
</p>
<div>
<p>Properties of server X509Names.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>issuerCertificateThumbprint</code><br/>
<em>
string
</em>
</td>
<td>
<p>IssuerCertificateThumbprint: Thumbprint for the Issuer of the Certificate.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: Common Name of the Certificate.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.X509CertificateName_STATUS">X509CertificateName_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendServiceFabricClusterProperties_STATUS">BackendServiceFabricClusterProperties_STATUS</a>)
</p>
<div>
<p>Properties of server X509Names.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>issuerCertificateThumbprint</code><br/>
<em>
string
</em>
</td>
<td>
<p>IssuerCertificateThumbprint: Thumbprint for the Issuer of the Certificate.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: Common Name of the Certificate.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="apimanagement.azure.com/v1api20230501preview.X509CertificateName_STATUS_ARM">X509CertificateName_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#apimanagement.azure.com/v1api20230501preview.BackendServiceFabricClusterProperties_STATUS_ARM">BackendServiceFabricClusterProperties_STATUS_ARM</a>)
</p>
<div>
<p>Properties of server X509Names.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>issuerCertificateThumbprint</code><br/>
<em>
string
</em>
</td>
<td>
<p>IssuerCertificateThumbprint: Thumbprint for the Issuer of the Certificate.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: Common Name of the Certificate.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
