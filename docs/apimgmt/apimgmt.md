# API Management Operator

The API Management suite is made up of the following operators:
* API Management Service
* API Management API

Learn more about Azure API Management [here](https://docs.microsoft.com/en-us/azure/api-management/api-management-key-concepts).

### API Management Service

The API Management Service deploys an API Management instance into a specified resource group at the specified location.

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

The API Management API Operator manages your API Management service, and deploys an instance into a specified resource group at the specified location.

The spec consists of the following fields:

* `apiService` The name of the API Management service to manage
* `apiId` Specify an ID for the API
* `properties`
   *  `isCurrent` Indicate if this API revision is the current revision
   *  `isOnline` Indicate if this API is accessible via the gateway
   *  `displayName` Display name for the API
   *  `description` Description of the API 
   *  `apiVersionDescription` Description of the API Version
   *  `path` Path for the API
   *  `protocols` Protocols for the service. Possible values are: 'http' or 'https'
   *  `subscriptionRequired` Specify whether an API or Product subscription is required for accessing the API

## Deploy, view and delete resources

You can follow the steps [here](/docs/customresource.md) to deploy, view and delete resources.
