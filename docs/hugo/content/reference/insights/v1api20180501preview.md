---
title: insights.azure.com/v1api20180501preview
---
<h2 id="insights.azure.com/v1api20180501preview">insights.azure.com/v1api20180501preview</h2>
<div>
<p>Package v1api20180501preview contains API Schema definitions for the insights v1api20180501preview API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="insights.azure.com/v1api20180501preview.APIVersion">APIVersion
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
<tbody><tr><td><p>&#34;2018-05-01-preview&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1api20180501preview.HeaderField">HeaderField
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20180501preview.WebTestProperties_Request">WebTestProperties_Request</a>)
</p>
<div>
<p>A header to add to the WebTest.</p>
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
<code>key</code><br/>
<em>
string
</em>
</td>
<td>
<p>Key: The name of the header.</p>
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
<p>Value: The value of the header.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20180501preview.HeaderField_STATUS">HeaderField_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20180501preview.WebTestProperties_Request_STATUS">WebTestProperties_Request_STATUS</a>)
</p>
<div>
<p>A header to add to the WebTest.</p>
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
<code>key</code><br/>
<em>
string
</em>
</td>
<td>
<p>Key: The name of the header.</p>
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
<p>Value: The value of the header.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20180501preview.WebTestGeolocation">WebTestGeolocation
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20180501preview.Webtest_Spec">Webtest_Spec</a>)
</p>
<div>
<p>Geo-physical location to run a WebTest from. You must specify one or more locations for the test to run from.</p>
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
<code>Id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Location ID for the WebTest to run from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20180501preview.WebTestGeolocation_STATUS">WebTestGeolocation_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20180501preview.Webtest_STATUS">Webtest_STATUS</a>)
</p>
<div>
<p>Geo-physical location to run a WebTest from. You must specify one or more locations for the test to run from.</p>
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
<code>Id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Location ID for the WebTest to run from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20180501preview.WebTestProperties_Configuration">WebTestProperties_Configuration
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20180501preview.Webtest_Spec">Webtest_Spec</a>)
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
<code>WebTest</code><br/>
<em>
string
</em>
</td>
<td>
<p>WebTest: The XML specification of a WebTest to run against an application.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20180501preview.WebTestProperties_Configuration_STATUS">WebTestProperties_Configuration_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20180501preview.Webtest_STATUS">Webtest_STATUS</a>)
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
<code>WebTest</code><br/>
<em>
string
</em>
</td>
<td>
<p>WebTest: The XML specification of a WebTest to run against an application.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20180501preview.WebTestProperties_Kind">WebTestProperties_Kind
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20180501preview.Webtest_Spec">Webtest_Spec</a>)
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
<tbody><tr><td><p>&#34;basic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;multistep&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ping&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1api20180501preview.WebTestProperties_Kind_STATUS">WebTestProperties_Kind_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20180501preview.Webtest_STATUS">Webtest_STATUS</a>)
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
<tbody><tr><td><p>&#34;basic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;multistep&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ping&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1api20180501preview.WebTestProperties_Request">WebTestProperties_Request
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20180501preview.Webtest_Spec">Webtest_Spec</a>)
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
<code>FollowRedirects</code><br/>
<em>
bool
</em>
</td>
<td>
<p>FollowRedirects: Follow redirects for this web test.</p>
</td>
</tr>
<tr>
<td>
<code>Headers</code><br/>
<em>
<a href="#insights.azure.com/v1api20180501preview.HeaderField">
[]HeaderField
</a>
</em>
</td>
<td>
<p>Headers: List of headers and their values to add to the WebTest call.</p>
</td>
</tr>
<tr>
<td>
<code>HttpVerb</code><br/>
<em>
string
</em>
</td>
<td>
<p>HttpVerb: Http verb to use for this web test.</p>
</td>
</tr>
<tr>
<td>
<code>ParseDependentRequests</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ParseDependentRequests: Parse Dependent request for this WebTest.</p>
</td>
</tr>
<tr>
<td>
<code>RequestBody</code><br/>
<em>
string
</em>
</td>
<td>
<p>RequestBody: Base64 encoded string body to send with this web test.</p>
</td>
</tr>
<tr>
<td>
<code>RequestUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>RequestUrl: Url location to test.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20180501preview.WebTestProperties_Request_STATUS">WebTestProperties_Request_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20180501preview.Webtest_STATUS">Webtest_STATUS</a>)
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
<code>FollowRedirects</code><br/>
<em>
bool
</em>
</td>
<td>
<p>FollowRedirects: Follow redirects for this web test.</p>
</td>
</tr>
<tr>
<td>
<code>Headers</code><br/>
<em>
<a href="#insights.azure.com/v1api20180501preview.HeaderField_STATUS">
[]HeaderField_STATUS
</a>
</em>
</td>
<td>
<p>Headers: List of headers and their values to add to the WebTest call.</p>
</td>
</tr>
<tr>
<td>
<code>HttpVerb</code><br/>
<em>
string
</em>
</td>
<td>
<p>HttpVerb: Http verb to use for this web test.</p>
</td>
</tr>
<tr>
<td>
<code>ParseDependentRequests</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ParseDependentRequests: Parse Dependent request for this WebTest.</p>
</td>
</tr>
<tr>
<td>
<code>RequestBody</code><br/>
<em>
string
</em>
</td>
<td>
<p>RequestBody: Base64 encoded string body to send with this web test.</p>
</td>
</tr>
<tr>
<td>
<code>RequestUrl</code><br/>
<em>
string
</em>
</td>
<td>
<p>RequestUrl: Url location to test.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20180501preview.WebTestProperties_ValidationRules">WebTestProperties_ValidationRules
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20180501preview.Webtest_Spec">Webtest_Spec</a>)
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
<code>ContentValidation</code><br/>
<em>
<a href="#insights.azure.com/v1api20180501preview.WebTestProperties_ValidationRules_ContentValidation">
WebTestProperties_ValidationRules_ContentValidation
</a>
</em>
</td>
<td>
<p>ContentValidation: The collection of content validation properties</p>
</td>
</tr>
<tr>
<td>
<code>ExpectedHttpStatusCode</code><br/>
<em>
int
</em>
</td>
<td>
<p>ExpectedHttpStatusCode: Validate that the WebTest returns the http status code provided.</p>
</td>
</tr>
<tr>
<td>
<code>IgnoreHttpsStatusCode</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IgnoreHttpsStatusCode: When set, validation will ignore the status code.</p>
</td>
</tr>
<tr>
<td>
<code>SSLCertRemainingLifetimeCheck</code><br/>
<em>
int
</em>
</td>
<td>
<p>SSLCertRemainingLifetimeCheck: A number of days to check still remain before the the existing SSL cert expires.  Value
must be positive and the SSLCheck must be set to true.</p>
</td>
</tr>
<tr>
<td>
<code>SSLCheck</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SSLCheck: Checks to see if the SSL cert is still valid.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20180501preview.WebTestProperties_ValidationRules_ContentValidation">WebTestProperties_ValidationRules_ContentValidation
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20180501preview.WebTestProperties_ValidationRules">WebTestProperties_ValidationRules</a>)
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
<code>ContentMatch</code><br/>
<em>
string
</em>
</td>
<td>
<p>ContentMatch: Content to look for in the return of the WebTest.  Must not be null or empty.</p>
</td>
</tr>
<tr>
<td>
<code>IgnoreCase</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IgnoreCase: When set, this value makes the ContentMatch validation case insensitive.</p>
</td>
</tr>
<tr>
<td>
<code>PassIfTextFound</code><br/>
<em>
bool
</em>
</td>
<td>
<p>PassIfTextFound: When true, validation will pass if there is a match for the ContentMatch string.  If false, validation
will fail if there is a match</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20180501preview.WebTestProperties_ValidationRules_ContentValidation_STATUS">WebTestProperties_ValidationRules_ContentValidation_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20180501preview.WebTestProperties_ValidationRules_STATUS">WebTestProperties_ValidationRules_STATUS</a>)
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
<code>ContentMatch</code><br/>
<em>
string
</em>
</td>
<td>
<p>ContentMatch: Content to look for in the return of the WebTest.  Must not be null or empty.</p>
</td>
</tr>
<tr>
<td>
<code>IgnoreCase</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IgnoreCase: When set, this value makes the ContentMatch validation case insensitive.</p>
</td>
</tr>
<tr>
<td>
<code>PassIfTextFound</code><br/>
<em>
bool
</em>
</td>
<td>
<p>PassIfTextFound: When true, validation will pass if there is a match for the ContentMatch string.  If false, validation
will fail if there is a match</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20180501preview.WebTestProperties_ValidationRules_STATUS">WebTestProperties_ValidationRules_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20180501preview.Webtest_STATUS">Webtest_STATUS</a>)
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
<code>ContentValidation</code><br/>
<em>
<a href="#insights.azure.com/v1api20180501preview.WebTestProperties_ValidationRules_ContentValidation_STATUS">
WebTestProperties_ValidationRules_ContentValidation_STATUS
</a>
</em>
</td>
<td>
<p>ContentValidation: The collection of content validation properties</p>
</td>
</tr>
<tr>
<td>
<code>ExpectedHttpStatusCode</code><br/>
<em>
int
</em>
</td>
<td>
<p>ExpectedHttpStatusCode: Validate that the WebTest returns the http status code provided.</p>
</td>
</tr>
<tr>
<td>
<code>IgnoreHttpsStatusCode</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IgnoreHttpsStatusCode: When set, validation will ignore the status code.</p>
</td>
</tr>
<tr>
<td>
<code>SSLCertRemainingLifetimeCheck</code><br/>
<em>
int
</em>
</td>
<td>
<p>SSLCertRemainingLifetimeCheck: A number of days to check still remain before the the existing SSL cert expires.  Value
must be positive and the SSLCheck must be set to true.</p>
</td>
</tr>
<tr>
<td>
<code>SSLCheck</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SSLCheck: Checks to see if the SSL cert is still valid.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20180501preview.Webtest">Webtest
</h3>
<div>
<p>Generator information:
- Generated from: /applicationinsights/resource-manager/Microsoft.Insights/preview/2018-05-01-preview/webTests_API.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Insights/&#x200b;webtests/&#x200b;{webTestName}</&#x200b;p>
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
<a href="#insights.azure.com/v1api20180501preview.Webtest_Spec">
Webtest_Spec
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
<code>Configuration</code><br/>
<em>
<a href="#insights.azure.com/v1api20180501preview.WebTestProperties_Configuration">
WebTestProperties_Configuration
</a>
</em>
</td>
<td>
<p>Configuration: An XML configuration specification for a WebTest.</p>
</td>
</tr>
<tr>
<td>
<code>Description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: User defined description for this WebTest.</p>
</td>
</tr>
<tr>
<td>
<code>Enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Is the test actively being monitored.</p>
</td>
</tr>
<tr>
<td>
<code>Frequency</code><br/>
<em>
int
</em>
</td>
<td>
<p>Frequency: Interval in seconds between test runs for this WebTest. Default value is 300.</p>
</td>
</tr>
<tr>
<td>
<code>Kind</code><br/>
<em>
<a href="#insights.azure.com/v1api20180501preview.WebTestProperties_Kind">
WebTestProperties_Kind
</a>
</em>
</td>
<td>
<p>Kind: The kind of web test this is, valid choices are ping, multistep, basic, and standard.</p>
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
<p>Location: Resource location</p>
</td>
</tr>
<tr>
<td>
<code>Locations</code><br/>
<em>
<a href="#insights.azure.com/v1api20180501preview.WebTestGeolocation">
[]WebTestGeolocation
</a>
</em>
</td>
<td>
<p>Locations: A list of where to physically run the tests from to give global coverage for accessibility of your
application.</p>
</td>
</tr>
<tr>
<td>
<code>Name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: User defined name if this WebTest.</p>
</td>
</tr>
<tr>
<td>
<code>operatorSpec</code><br/>
<em>
<a href="#insights.azure.com/v1api20180501preview.WebtestOperatorSpec">
WebtestOperatorSpec
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
reference to a resources.azure.com/ResourceGroup resource</p>
</td>
</tr>
<tr>
<td>
<code>Request</code><br/>
<em>
<a href="#insights.azure.com/v1api20180501preview.WebTestProperties_Request">
WebTestProperties_Request
</a>
</em>
</td>
<td>
<p>Request: The collection of request properties</p>
</td>
</tr>
<tr>
<td>
<code>RetryEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RetryEnabled: Allow for retries should this WebTest fail.</p>
</td>
</tr>
<tr>
<td>
<code>SyntheticMonitorId</code><br/>
<em>
string
</em>
</td>
<td>
<p>SyntheticMonitorId: Unique ID of this WebTest. This is typically the same value as the Name field.</p>
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
<p>Tags: Resource tags</p>
</td>
</tr>
<tr>
<td>
<code>Timeout</code><br/>
<em>
int
</em>
</td>
<td>
<p>Timeout: Seconds until this WebTest will timeout and fail. Default value is 30.</p>
</td>
</tr>
<tr>
<td>
<code>ValidationRules</code><br/>
<em>
<a href="#insights.azure.com/v1api20180501preview.WebTestProperties_ValidationRules">
WebTestProperties_ValidationRules
</a>
</em>
</td>
<td>
<p>ValidationRules: The collection of validation rule properties</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#insights.azure.com/v1api20180501preview.Webtest_STATUS">
Webtest_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20180501preview.WebtestOperatorSpec">WebtestOperatorSpec
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20180501preview.Webtest_Spec">Webtest_Spec</a>)
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
<code>configMapExpressions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#DestinationExpression">
[]genruntime/core.DestinationExpression
</a>
</em>
</td>
<td>
<p>ConfigMapExpressions: configures where to place operator written dynamic ConfigMaps (created with CEL expressions).</p>
</td>
</tr>
<tr>
<td>
<code>secretExpressions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#DestinationExpression">
[]genruntime/core.DestinationExpression
</a>
</em>
</td>
<td>
<p>SecretExpressions: configures where to place operator written dynamic secrets (created with CEL expressions).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20180501preview.Webtest_STATUS">Webtest_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20180501preview.Webtest">Webtest</a>)
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
<code>Configuration</code><br/>
<em>
<a href="#insights.azure.com/v1api20180501preview.WebTestProperties_Configuration_STATUS">
WebTestProperties_Configuration_STATUS
</a>
</em>
</td>
<td>
<p>Configuration: An XML configuration specification for a WebTest.</p>
</td>
</tr>
<tr>
<td>
<code>Description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: User defined description for this WebTest.</p>
</td>
</tr>
<tr>
<td>
<code>Enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Is the test actively being monitored.</p>
</td>
</tr>
<tr>
<td>
<code>Frequency</code><br/>
<em>
int
</em>
</td>
<td>
<p>Frequency: Interval in seconds between test runs for this WebTest. Default value is 300.</p>
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
<p>Id: Azure resource Id</p>
</td>
</tr>
<tr>
<td>
<code>Kind</code><br/>
<em>
<a href="#insights.azure.com/v1api20180501preview.WebTestProperties_Kind_STATUS">
WebTestProperties_Kind_STATUS
</a>
</em>
</td>
<td>
<p>Kind: The kind of web test this is, valid choices are ping, multistep, basic, and standard.</p>
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
<p>Location: Resource location</p>
</td>
</tr>
<tr>
<td>
<code>Locations</code><br/>
<em>
<a href="#insights.azure.com/v1api20180501preview.WebTestGeolocation_STATUS">
[]WebTestGeolocation_STATUS
</a>
</em>
</td>
<td>
<p>Locations: A list of where to physically run the tests from to give global coverage for accessibility of your
application.</p>
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
<p>Name: Azure resource name</p>
</td>
</tr>
<tr>
<td>
<code>properties_name</code><br/>
<em>
string
</em>
</td>
<td>
<p>PropertiesName: User defined name if this WebTest.</p>
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
<p>ProvisioningState: Current state of this component, whether or not is has been provisioned within the resource group it
is defined. Users cannot change this value but are able to read from it. Values will include Succeeded, Deploying,
Canceled, and Failed.</p>
</td>
</tr>
<tr>
<td>
<code>Request</code><br/>
<em>
<a href="#insights.azure.com/v1api20180501preview.WebTestProperties_Request_STATUS">
WebTestProperties_Request_STATUS
</a>
</em>
</td>
<td>
<p>Request: The collection of request properties</p>
</td>
</tr>
<tr>
<td>
<code>RetryEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RetryEnabled: Allow for retries should this WebTest fail.</p>
</td>
</tr>
<tr>
<td>
<code>SyntheticMonitorId</code><br/>
<em>
string
</em>
</td>
<td>
<p>SyntheticMonitorId: Unique ID of this WebTest. This is typically the same value as the Name field.</p>
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
<p>Tags: Resource tags</p>
</td>
</tr>
<tr>
<td>
<code>Timeout</code><br/>
<em>
int
</em>
</td>
<td>
<p>Timeout: Seconds until this WebTest will timeout and fail. Default value is 30.</p>
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
<p>Type: Azure resource type</p>
</td>
</tr>
<tr>
<td>
<code>ValidationRules</code><br/>
<em>
<a href="#insights.azure.com/v1api20180501preview.WebTestProperties_ValidationRules_STATUS">
WebTestProperties_ValidationRules_STATUS
</a>
</em>
</td>
<td>
<p>ValidationRules: The collection of validation rule properties</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20180501preview.Webtest_Spec">Webtest_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20180501preview.Webtest">Webtest</a>)
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
<code>Configuration</code><br/>
<em>
<a href="#insights.azure.com/v1api20180501preview.WebTestProperties_Configuration">
WebTestProperties_Configuration
</a>
</em>
</td>
<td>
<p>Configuration: An XML configuration specification for a WebTest.</p>
</td>
</tr>
<tr>
<td>
<code>Description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: User defined description for this WebTest.</p>
</td>
</tr>
<tr>
<td>
<code>Enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Is the test actively being monitored.</p>
</td>
</tr>
<tr>
<td>
<code>Frequency</code><br/>
<em>
int
</em>
</td>
<td>
<p>Frequency: Interval in seconds between test runs for this WebTest. Default value is 300.</p>
</td>
</tr>
<tr>
<td>
<code>Kind</code><br/>
<em>
<a href="#insights.azure.com/v1api20180501preview.WebTestProperties_Kind">
WebTestProperties_Kind
</a>
</em>
</td>
<td>
<p>Kind: The kind of web test this is, valid choices are ping, multistep, basic, and standard.</p>
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
<p>Location: Resource location</p>
</td>
</tr>
<tr>
<td>
<code>Locations</code><br/>
<em>
<a href="#insights.azure.com/v1api20180501preview.WebTestGeolocation">
[]WebTestGeolocation
</a>
</em>
</td>
<td>
<p>Locations: A list of where to physically run the tests from to give global coverage for accessibility of your
application.</p>
</td>
</tr>
<tr>
<td>
<code>Name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: User defined name if this WebTest.</p>
</td>
</tr>
<tr>
<td>
<code>operatorSpec</code><br/>
<em>
<a href="#insights.azure.com/v1api20180501preview.WebtestOperatorSpec">
WebtestOperatorSpec
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
reference to a resources.azure.com/ResourceGroup resource</p>
</td>
</tr>
<tr>
<td>
<code>Request</code><br/>
<em>
<a href="#insights.azure.com/v1api20180501preview.WebTestProperties_Request">
WebTestProperties_Request
</a>
</em>
</td>
<td>
<p>Request: The collection of request properties</p>
</td>
</tr>
<tr>
<td>
<code>RetryEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RetryEnabled: Allow for retries should this WebTest fail.</p>
</td>
</tr>
<tr>
<td>
<code>SyntheticMonitorId</code><br/>
<em>
string
</em>
</td>
<td>
<p>SyntheticMonitorId: Unique ID of this WebTest. This is typically the same value as the Name field.</p>
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
<p>Tags: Resource tags</p>
</td>
</tr>
<tr>
<td>
<code>Timeout</code><br/>
<em>
int
</em>
</td>
<td>
<p>Timeout: Seconds until this WebTest will timeout and fail. Default value is 30.</p>
</td>
</tr>
<tr>
<td>
<code>ValidationRules</code><br/>
<em>
<a href="#insights.azure.com/v1api20180501preview.WebTestProperties_ValidationRules">
WebTestProperties_ValidationRules
</a>
</em>
</td>
<td>
<p>ValidationRules: The collection of validation rule properties</p>
</td>
</tr>
</tbody>
</table>
<hr/>
