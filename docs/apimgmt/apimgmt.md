# API Management Operator

The API Management suite is made up of the following operators:
* API Management Service
* API Management API

Learn more about Azure API Management [here](https://docs.microsoft.com/en-us/azure/api-management/api-management-key-concepts).

### API Management Service

The API Management Service deploys an API Management instance into a specified resource group at the specified location. It also provides the option to link to an Application Insights instance for logging, and to place the API Management instance in a specified Virtual Network.

Here is a [sample YAML](/config/samples/azure_v1alpha1_apimservice.yaml) to provision an API Management Service.

The spec consists of the following fields:

#### Required Fields
* `Tier` Specify the tier of the service. Options include: 'basic', 'standard', and 'premium'.
* `PublisherName` Specify the name of the publisher.
* `PublisherEmail` Specify the email of the publisher.

#### Optional Fields
* `VnetType` Specify the type of vnet for the service. Options include: 'none', 'internal', and 'external'. If selecting either 'internal' or 'external', make sure to set the `Tier` to 'premium'.
* `VnetResourceGroup` Resource group of the Virtual Network
* `VnetName` Name of the Virtual Network
* `VnetSubnetName` Name of the Virtual Network Subnet
* `AppInsightsResourceGroup` Resource group of the Application Insights instance
* `AppInsightsName` Name of the Application Insights instance

### API Management API

The API Management API Operator creates an API with the specified properties in the specified API Management service.

Here is a [sample YAML](/config/samples/azure_v1alpha1_apimgmtapi.yaml) to provision an API Management API.

The spec consists of the following fields:

* `apiService` The name of the API Management service to manage
* `apiId` Specify an ID for the API
* `properties`
   * `apiRevision` Describes the Revision of the API. If no value is provided, default revision 1 is created
   * `apiRevisionDescription` Description of the API Revision
   * `apiVersion`  Indicates the Version identifier of the API if the API is versioned
   * `apiVersionDescription` Description of the API Version
   * `apiVersionSet`  API Version Set contains the common configuration for a set of API versions
      * `id` Identifier for existing API Version Set. Omit this value to create a new Version Set.
      * `name` The display Name of the API Version Set.
      * `description` Description of API Version Set.
   * `apiVersionSetID` A resource identifier for the related ApiVersionSet
   * `description` Description of the API 
   * `displayName` Display name for the API. Must be 1 to 300 characters long
   * `format` Format of the Content in which the API is getting imported. Possible values include: 'WadlXML', 'WadlLinkJSON', 'SwaggerJSON', 'SwaggerLinkJSON', 'Wsdl', 'WsdlLink', 'Openapi', 'Openapijson', 'OpenapiLink'
   * `isCurrent` Indicate if this API revision is the current revision
   * `isOnline` Indicate if this API is accessible via the gateway
   * `path` Path for the API
   * `protocols` Describes on which protocols the operations in this API can be invoked. Possible values are: 'http' or 'https'
   * `serviceURL` Absolute URL of the backend service implementing this API. Cannot be more than 2000 characters long
   * `sourceAPIID` API identifier of the source API
   * `subscriptionRequired` Specify whether an API or Product subscription is required for accessing the API

## Deploy, view and delete resources

You can follow the steps [here](/docs/customresource.md) to deploy, view and delete resources.
