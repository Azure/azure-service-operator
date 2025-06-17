---

title: insights.azure.com/v1api20180501preview

linktitle: v1api20180501preview
-------------------------------

APIVersion{#APIVersion}
-----------------------

| Value                | Description |
|----------------------|-------------|
| "2018-05-01-preview" |             |

Webtest{#Webtest}
-----------------

Generator information: - Generated from: /applicationinsights/resource-manager/Microsoft.Insights/preview/2018-05-01-preview/webTests_API.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Insights/&ZeroWidthSpace;webtests/&ZeroWidthSpace;{webTestName}

Used by: [WebtestList](#WebtestList).

| Property                                                                                | Description | Type                                                          |
|-----------------------------------------------------------------------------------------|-------------|---------------------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta)     |             |                                                               |
| [metav1.ObjectMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ObjectMeta) |             |                                                               |
| spec                                                                                    |             | [Webtest_Spec](#Webtest_Spec)<br/><small>Optional</small>     |
| status                                                                                  |             | [Webtest_STATUS](#Webtest_STATUS)<br/><small>Optional</small> |

### Webtest_Spec {#Webtest_Spec}

| Property           | Description                                                                                                                                                                                                                                                                                  | Type                                                                                                                                                                 |
|--------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| azureName          | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                                   |
| Configuration      | An XML configuration specification for a WebTest.                                                                                                                                                                                                                                            | [WebTestProperties_Configuration](#WebTestProperties_Configuration)<br/><small>Optional</small>                                                                      |
| Description        | User defined description for this WebTest.                                                                                                                                                                                                                                                   | string<br/><small>Optional</small>                                                                                                                                   |
| Enabled            | Is the test actively being monitored.                                                                                                                                                                                                                                                        | bool<br/><small>Optional</small>                                                                                                                                     |
| Frequency          | Interval in seconds between test runs for this WebTest. Default value is 300.                                                                                                                                                                                                                | int<br/><small>Optional</small>                                                                                                                                      |
| Kind               | The kind of web test this is, valid choices are ping, multistep, basic, and standard.                                                                                                                                                                                                        | [WebTestProperties_Kind](#WebTestProperties_Kind)<br/><small>Required</small>                                                                                        |
| location           | Resource location                                                                                                                                                                                                                                                                            | string<br/><small>Required</small>                                                                                                                                   |
| Locations          | A list of where to physically run the tests from to give global coverage for accessibility of your application.                                                                                                                                                                              | [WebTestGeolocation[]](#WebTestGeolocation)<br/><small>Required</small>                                                                                              |
| Name               | User defined name if this WebTest.                                                                                                                                                                                                                                                           | string<br/><small>Required</small>                                                                                                                                   |
| operatorSpec       | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                              | [WebtestOperatorSpec](#WebtestOperatorSpec)<br/><small>Optional</small>                                                                                              |
| owner              | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a resources.azure.com/ResourceGroup resource | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |
| Request            | The collection of request properties                                                                                                                                                                                                                                                         | [WebTestProperties_Request](#WebTestProperties_Request)<br/><small>Optional</small>                                                                                  |
| RetryEnabled       | Allow for retries should this WebTest fail.                                                                                                                                                                                                                                                  | bool<br/><small>Optional</small>                                                                                                                                     |
| SyntheticMonitorId | Unique ID of this WebTest. This is typically the same value as the Name field.                                                                                                                                                                                                               | string<br/><small>Required</small>                                                                                                                                   |
| tags               | Resource tags                                                                                                                                                                                                                                                                                | map[string]string<br/><small>Optional</small>                                                                                                                        |
| Timeout            | Seconds until this WebTest will timeout and fail. Default value is 30.                                                                                                                                                                                                                       | int<br/><small>Optional</small>                                                                                                                                      |
| ValidationRules    | The collection of validation rule properties                                                                                                                                                                                                                                                 | [WebTestProperties_ValidationRules](#WebTestProperties_ValidationRules)<br/><small>Optional</small>                                                                  |

### Webtest_STATUS{#Webtest_STATUS}

| Property           | Description                                                                                                                                                                                                                                   | Type                                                                                                                                                    |
|--------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| conditions         | The observed state of the resource                                                                                                                                                                                                            | [conditions.Condition[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions#Condition)<br/><small>Optional</small> |
| Configuration      | An XML configuration specification for a WebTest.                                                                                                                                                                                             | [WebTestProperties_Configuration_STATUS](#WebTestProperties_Configuration_STATUS)<br/><small>Optional</small>                                           |
| Description        | User defined description for this WebTest.                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                      |
| Enabled            | Is the test actively being monitored.                                                                                                                                                                                                         | bool<br/><small>Optional</small>                                                                                                                        |
| Frequency          | Interval in seconds between test runs for this WebTest. Default value is 300.                                                                                                                                                                 | int<br/><small>Optional</small>                                                                                                                         |
| id                 | Azure resource Id                                                                                                                                                                                                                             | string<br/><small>Optional</small>                                                                                                                      |
| Kind               | The kind of web test this is, valid choices are ping, multistep, basic, and standard.                                                                                                                                                         | [WebTestProperties_Kind_STATUS](#WebTestProperties_Kind_STATUS)<br/><small>Optional</small>                                                             |
| location           | Resource location                                                                                                                                                                                                                             | string<br/><small>Optional</small>                                                                                                                      |
| Locations          | A list of where to physically run the tests from to give global coverage for accessibility of your application.                                                                                                                               | [WebTestGeolocation_STATUS[]](#WebTestGeolocation_STATUS)<br/><small>Optional</small>                                                                   |
| name               | Azure resource name                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                      |
| properties_name    | User defined name if this WebTest.                                                                                                                                                                                                            | string<br/><small>Optional</small>                                                                                                                      |
| provisioningState  | Current state of this component, whether or not is has been provisioned within the resource group it is defined. Users cannot change this value but are able to read from it. Values will include Succeeded, Deploying, Canceled, and Failed. | string<br/><small>Optional</small>                                                                                                                      |
| Request            | The collection of request properties                                                                                                                                                                                                          | [WebTestProperties_Request_STATUS](#WebTestProperties_Request_STATUS)<br/><small>Optional</small>                                                       |
| RetryEnabled       | Allow for retries should this WebTest fail.                                                                                                                                                                                                   | bool<br/><small>Optional</small>                                                                                                                        |
| SyntheticMonitorId | Unique ID of this WebTest. This is typically the same value as the Name field.                                                                                                                                                                | string<br/><small>Optional</small>                                                                                                                      |
| tags               | Resource tags                                                                                                                                                                                                                                 | map[string]string<br/><small>Optional</small>                                                                                                           |
| Timeout            | Seconds until this WebTest will timeout and fail. Default value is 30.                                                                                                                                                                        | int<br/><small>Optional</small>                                                                                                                         |
| type               | Azure resource type                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                      |
| ValidationRules    | The collection of validation rule properties                                                                                                                                                                                                  | [WebTestProperties_ValidationRules_STATUS](#WebTestProperties_ValidationRules_STATUS)<br/><small>Optional</small>                                       |

WebtestList{#WebtestList}
-------------------------

Generator information: - Generated from: /applicationinsights/resource-manager/Microsoft.Insights/preview/2018-05-01-preview/webTests_API.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Insights/&ZeroWidthSpace;webtests/&ZeroWidthSpace;{webTestName}

| Property                                                                            | Description | Type                                              |
|-------------------------------------------------------------------------------------|-------------|---------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta) |             |                                                   |
| [metav1.ListMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ListMeta) |             |                                                   |
| items                                                                               |             | [Webtest[]](#Webtest)<br/><small>Optional</small> |

Webtest_Spec{#Webtest_Spec}
---------------------------

Used by: [Webtest](#Webtest).

| Property           | Description                                                                                                                                                                                                                                                                                  | Type                                                                                                                                                                 |
|--------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| azureName          | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                                   |
| Configuration      | An XML configuration specification for a WebTest.                                                                                                                                                                                                                                            | [WebTestProperties_Configuration](#WebTestProperties_Configuration)<br/><small>Optional</small>                                                                      |
| Description        | User defined description for this WebTest.                                                                                                                                                                                                                                                   | string<br/><small>Optional</small>                                                                                                                                   |
| Enabled            | Is the test actively being monitored.                                                                                                                                                                                                                                                        | bool<br/><small>Optional</small>                                                                                                                                     |
| Frequency          | Interval in seconds between test runs for this WebTest. Default value is 300.                                                                                                                                                                                                                | int<br/><small>Optional</small>                                                                                                                                      |
| Kind               | The kind of web test this is, valid choices are ping, multistep, basic, and standard.                                                                                                                                                                                                        | [WebTestProperties_Kind](#WebTestProperties_Kind)<br/><small>Required</small>                                                                                        |
| location           | Resource location                                                                                                                                                                                                                                                                            | string<br/><small>Required</small>                                                                                                                                   |
| Locations          | A list of where to physically run the tests from to give global coverage for accessibility of your application.                                                                                                                                                                              | [WebTestGeolocation[]](#WebTestGeolocation)<br/><small>Required</small>                                                                                              |
| Name               | User defined name if this WebTest.                                                                                                                                                                                                                                                           | string<br/><small>Required</small>                                                                                                                                   |
| operatorSpec       | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                              | [WebtestOperatorSpec](#WebtestOperatorSpec)<br/><small>Optional</small>                                                                                              |
| owner              | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a resources.azure.com/ResourceGroup resource | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |
| Request            | The collection of request properties                                                                                                                                                                                                                                                         | [WebTestProperties_Request](#WebTestProperties_Request)<br/><small>Optional</small>                                                                                  |
| RetryEnabled       | Allow for retries should this WebTest fail.                                                                                                                                                                                                                                                  | bool<br/><small>Optional</small>                                                                                                                                     |
| SyntheticMonitorId | Unique ID of this WebTest. This is typically the same value as the Name field.                                                                                                                                                                                                               | string<br/><small>Required</small>                                                                                                                                   |
| tags               | Resource tags                                                                                                                                                                                                                                                                                | map[string]string<br/><small>Optional</small>                                                                                                                        |
| Timeout            | Seconds until this WebTest will timeout and fail. Default value is 30.                                                                                                                                                                                                                       | int<br/><small>Optional</small>                                                                                                                                      |
| ValidationRules    | The collection of validation rule properties                                                                                                                                                                                                                                                 | [WebTestProperties_ValidationRules](#WebTestProperties_ValidationRules)<br/><small>Optional</small>                                                                  |

Webtest_STATUS{#Webtest_STATUS}
-------------------------------

Used by: [Webtest](#Webtest).

| Property           | Description                                                                                                                                                                                                                                   | Type                                                                                                                                                    |
|--------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| conditions         | The observed state of the resource                                                                                                                                                                                                            | [conditions.Condition[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions#Condition)<br/><small>Optional</small> |
| Configuration      | An XML configuration specification for a WebTest.                                                                                                                                                                                             | [WebTestProperties_Configuration_STATUS](#WebTestProperties_Configuration_STATUS)<br/><small>Optional</small>                                           |
| Description        | User defined description for this WebTest.                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                      |
| Enabled            | Is the test actively being monitored.                                                                                                                                                                                                         | bool<br/><small>Optional</small>                                                                                                                        |
| Frequency          | Interval in seconds between test runs for this WebTest. Default value is 300.                                                                                                                                                                 | int<br/><small>Optional</small>                                                                                                                         |
| id                 | Azure resource Id                                                                                                                                                                                                                             | string<br/><small>Optional</small>                                                                                                                      |
| Kind               | The kind of web test this is, valid choices are ping, multistep, basic, and standard.                                                                                                                                                         | [WebTestProperties_Kind_STATUS](#WebTestProperties_Kind_STATUS)<br/><small>Optional</small>                                                             |
| location           | Resource location                                                                                                                                                                                                                             | string<br/><small>Optional</small>                                                                                                                      |
| Locations          | A list of where to physically run the tests from to give global coverage for accessibility of your application.                                                                                                                               | [WebTestGeolocation_STATUS[]](#WebTestGeolocation_STATUS)<br/><small>Optional</small>                                                                   |
| name               | Azure resource name                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                      |
| properties_name    | User defined name if this WebTest.                                                                                                                                                                                                            | string<br/><small>Optional</small>                                                                                                                      |
| provisioningState  | Current state of this component, whether or not is has been provisioned within the resource group it is defined. Users cannot change this value but are able to read from it. Values will include Succeeded, Deploying, Canceled, and Failed. | string<br/><small>Optional</small>                                                                                                                      |
| Request            | The collection of request properties                                                                                                                                                                                                          | [WebTestProperties_Request_STATUS](#WebTestProperties_Request_STATUS)<br/><small>Optional</small>                                                       |
| RetryEnabled       | Allow for retries should this WebTest fail.                                                                                                                                                                                                   | bool<br/><small>Optional</small>                                                                                                                        |
| SyntheticMonitorId | Unique ID of this WebTest. This is typically the same value as the Name field.                                                                                                                                                                | string<br/><small>Optional</small>                                                                                                                      |
| tags               | Resource tags                                                                                                                                                                                                                                 | map[string]string<br/><small>Optional</small>                                                                                                           |
| Timeout            | Seconds until this WebTest will timeout and fail. Default value is 30.                                                                                                                                                                        | int<br/><small>Optional</small>                                                                                                                         |
| type               | Azure resource type                                                                                                                                                                                                                           | string<br/><small>Optional</small>                                                                                                                      |
| ValidationRules    | The collection of validation rule properties                                                                                                                                                                                                  | [WebTestProperties_ValidationRules_STATUS](#WebTestProperties_ValidationRules_STATUS)<br/><small>Optional</small>                                       |

WebTestGeolocation{#WebTestGeolocation}
---------------------------------------

Geo-physical location to run a WebTest from. You must specify one or more locations for the test to run from.

Used by: [Webtest_Spec](#Webtest_Spec).

| Property | Description                              | Type                               |
|----------|------------------------------------------|------------------------------------|
| Id       | Location ID for the WebTest to run from. | string<br/><small>Optional</small> |

WebTestGeolocation_STATUS{#WebTestGeolocation_STATUS}
-----------------------------------------------------

Geo-physical location to run a WebTest from. You must specify one or more locations for the test to run from.

Used by: [Webtest_STATUS](#Webtest_STATUS).

| Property | Description                              | Type                               |
|----------|------------------------------------------|------------------------------------|
| Id       | Location ID for the WebTest to run from. | string<br/><small>Optional</small> |

WebtestOperatorSpec{#WebtestOperatorSpec}
-----------------------------------------

Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure

Used by: [Webtest_Spec](#Webtest_Spec).

| Property             | Description                                                                                   | Type                                                                                                                                                                |
|----------------------|-----------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| configMapExpressions | configures where to place operator written dynamic ConfigMaps (created with CEL expressions). | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |
| secretExpressions    | configures where to place operator written dynamic secrets (created with CEL expressions).    | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |

WebTestProperties_Configuration{#WebTestProperties_Configuration}
-----------------------------------------------------------------

Used by: [Webtest_Spec](#Webtest_Spec).

| Property | Description                                                       | Type                               |
|----------|-------------------------------------------------------------------|------------------------------------|
| WebTest  | The XML specification of a WebTest to run against an application. | string<br/><small>Optional</small> |

WebTestProperties_Configuration_STATUS{#WebTestProperties_Configuration_STATUS}
-------------------------------------------------------------------------------

Used by: [Webtest_STATUS](#Webtest_STATUS).

| Property | Description                                                       | Type                               |
|----------|-------------------------------------------------------------------|------------------------------------|
| WebTest  | The XML specification of a WebTest to run against an application. | string<br/><small>Optional</small> |

WebTestProperties_Kind{#WebTestProperties_Kind}
-----------------------------------------------

Used by: [Webtest_Spec](#Webtest_Spec).

| Value       | Description |
|-------------|-------------|
| "basic"     |             |
| "multistep" |             |
| "ping"      |             |
| "standard"  |             |

WebTestProperties_Kind_STATUS{#WebTestProperties_Kind_STATUS}
-------------------------------------------------------------

Used by: [Webtest_STATUS](#Webtest_STATUS).

| Value       | Description |
|-------------|-------------|
| "basic"     |             |
| "multistep" |             |
| "ping"      |             |
| "standard"  |             |

WebTestProperties_Request{#WebTestProperties_Request}
-----------------------------------------------------

Used by: [Webtest_Spec](#Webtest_Spec).

| Property               | Description                                                  | Type                                                      |
|------------------------|--------------------------------------------------------------|-----------------------------------------------------------|
| FollowRedirects        | Follow redirects for this web test.                          | bool<br/><small>Optional</small>                          |
| Headers                | List of headers and their values to add to the WebTest call. | [HeaderField[]](#HeaderField)<br/><small>Optional</small> |
| HttpVerb               | Http verb to use for this web test.                          | string<br/><small>Optional</small>                        |
| ParseDependentRequests | Parse Dependent request for this WebTest.                    | bool<br/><small>Optional</small>                          |
| RequestBody            | Base64 encoded string body to send with this web test.       | string<br/><small>Optional</small>                        |
| RequestUrl             | Url location to test.                                        | string<br/><small>Optional</small>                        |

WebTestProperties_Request_STATUS{#WebTestProperties_Request_STATUS}
-------------------------------------------------------------------

Used by: [Webtest_STATUS](#Webtest_STATUS).

| Property               | Description                                                  | Type                                                                    |
|------------------------|--------------------------------------------------------------|-------------------------------------------------------------------------|
| FollowRedirects        | Follow redirects for this web test.                          | bool<br/><small>Optional</small>                                        |
| Headers                | List of headers and their values to add to the WebTest call. | [HeaderField_STATUS[]](#HeaderField_STATUS)<br/><small>Optional</small> |
| HttpVerb               | Http verb to use for this web test.                          | string<br/><small>Optional</small>                                      |
| ParseDependentRequests | Parse Dependent request for this WebTest.                    | bool<br/><small>Optional</small>                                        |
| RequestBody            | Base64 encoded string body to send with this web test.       | string<br/><small>Optional</small>                                      |
| RequestUrl             | Url location to test.                                        | string<br/><small>Optional</small>                                      |

WebTestProperties_ValidationRules{#WebTestProperties_ValidationRules}
---------------------------------------------------------------------

Used by: [Webtest_Spec](#Webtest_Spec).

| Property                      | Description                                                                                                                                   | Type                                                                                                                                    |
|-------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------|
| ContentValidation             | The collection of content validation properties                                                                                               | [WebTestProperties_ValidationRules_ContentValidation](#WebTestProperties_ValidationRules_ContentValidation)<br/><small>Optional</small> |
| ExpectedHttpStatusCode        | Validate that the WebTest returns the http status code provided.                                                                              | int<br/><small>Optional</small>                                                                                                         |
| IgnoreHttpsStatusCode         | When set, validation will ignore the status code.                                                                                             | bool<br/><small>Optional</small>                                                                                                        |
| SSLCertRemainingLifetimeCheck | A number of days to check still remain before the the existing SSL cert expires. Value must be positive and the SSLCheck must be set to true. | int<br/><small>Optional</small>                                                                                                         |
| SSLCheck                      | Checks to see if the SSL cert is still valid.                                                                                                 | bool<br/><small>Optional</small>                                                                                                        |

WebTestProperties_ValidationRules_STATUS{#WebTestProperties_ValidationRules_STATUS}
-----------------------------------------------------------------------------------

Used by: [Webtest_STATUS](#Webtest_STATUS).

| Property                      | Description                                                                                                                                   | Type                                                                                                                                                  |
|-------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------|
| ContentValidation             | The collection of content validation properties                                                                                               | [WebTestProperties_ValidationRules_ContentValidation_STATUS](#WebTestProperties_ValidationRules_ContentValidation_STATUS)<br/><small>Optional</small> |
| ExpectedHttpStatusCode        | Validate that the WebTest returns the http status code provided.                                                                              | int<br/><small>Optional</small>                                                                                                                       |
| IgnoreHttpsStatusCode         | When set, validation will ignore the status code.                                                                                             | bool<br/><small>Optional</small>                                                                                                                      |
| SSLCertRemainingLifetimeCheck | A number of days to check still remain before the the existing SSL cert expires. Value must be positive and the SSLCheck must be set to true. | int<br/><small>Optional</small>                                                                                                                       |
| SSLCheck                      | Checks to see if the SSL cert is still valid.                                                                                                 | bool<br/><small>Optional</small>                                                                                                                      |

HeaderField{#HeaderField}
-------------------------

A header to add to the WebTest.

Used by: [WebTestProperties_Request](#WebTestProperties_Request).

| Property | Description              | Type                               |
|----------|--------------------------|------------------------------------|
| key      | The name of the header.  | string<br/><small>Optional</small> |
| value    | The value of the header. | string<br/><small>Optional</small> |

HeaderField_STATUS{#HeaderField_STATUS}
---------------------------------------

A header to add to the WebTest.

Used by: [WebTestProperties_Request_STATUS](#WebTestProperties_Request_STATUS).

| Property | Description              | Type                               |
|----------|--------------------------|------------------------------------|
| key      | The name of the header.  | string<br/><small>Optional</small> |
| value    | The value of the header. | string<br/><small>Optional</small> |

WebTestProperties_ValidationRules_ContentValidation{#WebTestProperties_ValidationRules_ContentValidation}
---------------------------------------------------------------------------------------------------------

Used by: [WebTestProperties_ValidationRules](#WebTestProperties_ValidationRules).

| Property        | Description                                                                                                                         | Type                               |
|-----------------|-------------------------------------------------------------------------------------------------------------------------------------|------------------------------------|
| ContentMatch    | Content to look for in the return of the WebTest. Must not be null or empty.                                                        | string<br/><small>Optional</small> |
| IgnoreCase      | When set, this value makes the ContentMatch validation case insensitive.                                                            | bool<br/><small>Optional</small>   |
| PassIfTextFound | When true, validation will pass if there is a match for the ContentMatch string. If false, validation will fail if there is a match | bool<br/><small>Optional</small>   |

WebTestProperties_ValidationRules_ContentValidation_STATUS{#WebTestProperties_ValidationRules_ContentValidation_STATUS}
-----------------------------------------------------------------------------------------------------------------------

Used by: [WebTestProperties_ValidationRules_STATUS](#WebTestProperties_ValidationRules_STATUS).

| Property        | Description                                                                                                                         | Type                               |
|-----------------|-------------------------------------------------------------------------------------------------------------------------------------|------------------------------------|
| ContentMatch    | Content to look for in the return of the WebTest. Must not be null or empty.                                                        | string<br/><small>Optional</small> |
| IgnoreCase      | When set, this value makes the ContentMatch validation case insensitive.                                                            | bool<br/><small>Optional</small>   |
| PassIfTextFound | When true, validation will pass if there is a match for the ContentMatch string. If false, validation will fail if there is a match | bool<br/><small>Optional</small>   |
