---
---
<h2 id="insights.azure.com/v1alpha1api20200202">insights.azure.com/v1alpha1api20200202</h2>
<div>
<p>Package v1alpha1api20200202 contains API Schema definitions for the insights v1alpha1api20200202 API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesARM">ApplicationInsightsComponentPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.Components_SpecARM">Components_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-02-02/Microsoft.Insights.Application.json#/definitions/ApplicationInsightsComponentProperties">https://schema.management.azure.com/schemas/2020-02-02/Microsoft.Insights.Application.json#/definitions/ApplicationInsightsComponentProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Application_Type</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesApplicationType">
ApplicationInsightsComponentPropertiesApplicationType
</a>
</em>
</td>
<td>
<p>ApplicationType: Type of application being monitored.</p>
</td>
</tr>
<tr>
<td>
<code>DisableIpMasking</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableIpMasking: Disable IP masking.</p>
</td>
</tr>
<tr>
<td>
<code>DisableLocalAuth</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableLocalAuth: Disable Non-AAD based Auth.</p>
</td>
</tr>
<tr>
<td>
<code>Flow_Type</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesFlowType">
ApplicationInsightsComponentPropertiesFlowType
</a>
</em>
</td>
<td>
<p>FlowType: Used by the Application Insights system to determine what kind of flow this component was created by. This is
to be set to &lsquo;Bluefield&rsquo; when creating/updating a component via the REST API.</p>
</td>
</tr>
<tr>
<td>
<code>ForceCustomerStorageForProfiler</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ForceCustomerStorageForProfiler: Force users to create their own storage account for profiler and debugger.</p>
</td>
</tr>
<tr>
<td>
<code>HockeyAppId</code><br/>
<em>
string
</em>
</td>
<td>
<p>HockeyAppId: The unique application ID created when a new application is added to HockeyApp, used for communications
with HockeyApp.</p>
</td>
</tr>
<tr>
<td>
<code>ImmediatePurgeDataOn30Days</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ImmediatePurgeDataOn30Days: Purge data immediately after 30 days.</p>
</td>
</tr>
<tr>
<td>
<code>IngestionMode</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesIngestionMode">
ApplicationInsightsComponentPropertiesIngestionMode
</a>
</em>
</td>
<td>
<p>IngestionMode: Indicates the flow of the ingestion.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForIngestion</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesPublicNetworkAccessForIngestion">
ApplicationInsightsComponentPropertiesPublicNetworkAccessForIngestion
</a>
</em>
</td>
<td>
<p>PublicNetworkAccessForIngestion: The network access type for accessing Application Insights ingestion.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForQuery</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesPublicNetworkAccessForQuery">
ApplicationInsightsComponentPropertiesPublicNetworkAccessForQuery
</a>
</em>
</td>
<td>
<p>PublicNetworkAccessForQuery: The network access type for accessing Application Insights query.</p>
</td>
</tr>
<tr>
<td>
<code>Request_Source</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesRequestSource">
ApplicationInsightsComponentPropertiesRequestSource
</a>
</em>
</td>
<td>
<p>RequestSource: Describes what tool created this Application Insights component. Customers using this API should set this
to the default &lsquo;rest&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>RetentionInDays</code><br/>
<em>
int
</em>
</td>
<td>
<p>RetentionInDays: Retention period in days.</p>
</td>
</tr>
<tr>
<td>
<code>SamplingPercentage</code><br/>
<em>
float64
</em>
</td>
<td>
<p>SamplingPercentage: Percentage of the data produced by the application being monitored that is being sampled for
Application Insights telemetry.</p>
</td>
</tr>
<tr>
<td>
<code>workspaceResourceId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesApplicationType">ApplicationInsightsComponentPropertiesApplicationType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesARM">ApplicationInsightsComponentPropertiesARM</a>, <a href="#insights.azure.com/v1alpha1api20200202.Components_Spec">Components_Spec</a>)
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
<tbody><tr><td><p>&#34;other&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;web&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesFlowType">ApplicationInsightsComponentPropertiesFlowType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesARM">ApplicationInsightsComponentPropertiesARM</a>, <a href="#insights.azure.com/v1alpha1api20200202.Components_Spec">Components_Spec</a>)
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
<tbody><tr><td><p>&#34;Bluefield&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesIngestionMode">ApplicationInsightsComponentPropertiesIngestionMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesARM">ApplicationInsightsComponentPropertiesARM</a>, <a href="#insights.azure.com/v1alpha1api20200202.Components_Spec">Components_Spec</a>)
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
<tbody><tr><td><p>&#34;ApplicationInsights&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ApplicationInsightsWithDiagnosticSettings&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;LogAnalytics&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesPublicNetworkAccessForIngestion">ApplicationInsightsComponentPropertiesPublicNetworkAccessForIngestion
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesARM">ApplicationInsightsComponentPropertiesARM</a>, <a href="#insights.azure.com/v1alpha1api20200202.Components_Spec">Components_Spec</a>)
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
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesPublicNetworkAccessForQuery">ApplicationInsightsComponentPropertiesPublicNetworkAccessForQuery
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesARM">ApplicationInsightsComponentPropertiesARM</a>, <a href="#insights.azure.com/v1alpha1api20200202.Components_Spec">Components_Spec</a>)
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
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesRequestSource">ApplicationInsightsComponentPropertiesRequestSource
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesARM">ApplicationInsightsComponentPropertiesARM</a>, <a href="#insights.azure.com/v1alpha1api20200202.Components_Spec">Components_Spec</a>)
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
<tbody><tr><td><p>&#34;rest&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesStatusApplicationType">ApplicationInsightsComponentPropertiesStatusApplicationType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentProperties_StatusARM">ApplicationInsightsComponentProperties_StatusARM</a>, <a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponent_Status">ApplicationInsightsComponent_Status</a>)
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
<tbody><tr><td><p>&#34;other&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;web&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesStatusFlowType">ApplicationInsightsComponentPropertiesStatusFlowType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentProperties_StatusARM">ApplicationInsightsComponentProperties_StatusARM</a>, <a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponent_Status">ApplicationInsightsComponent_Status</a>)
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
<tbody><tr><td><p>&#34;Bluefield&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesStatusIngestionMode">ApplicationInsightsComponentPropertiesStatusIngestionMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentProperties_StatusARM">ApplicationInsightsComponentProperties_StatusARM</a>, <a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponent_Status">ApplicationInsightsComponent_Status</a>)
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
<tbody><tr><td><p>&#34;ApplicationInsights&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ApplicationInsightsWithDiagnosticSettings&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;LogAnalytics&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesStatusRequestSource">ApplicationInsightsComponentPropertiesStatusRequestSource
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentProperties_StatusARM">ApplicationInsightsComponentProperties_StatusARM</a>, <a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponent_Status">ApplicationInsightsComponent_Status</a>)
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
<tbody><tr><td><p>&#34;rest&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentProperties_StatusARM">ApplicationInsightsComponentProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponent_StatusARM">ApplicationInsightsComponent_StatusARM</a>)
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
<code>AppId</code><br/>
<em>
string
</em>
</td>
<td>
<p>AppId: Application Insights Unique ID for your Application.</p>
</td>
</tr>
<tr>
<td>
<code>ApplicationId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ApplicationId: The unique ID of your application. This field mirrors the &lsquo;Name&rsquo; field and cannot be changed.</p>
</td>
</tr>
<tr>
<td>
<code>Application_Type</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesStatusApplicationType">
ApplicationInsightsComponentPropertiesStatusApplicationType
</a>
</em>
</td>
<td>
<p>ApplicationType: Type of application being monitored.</p>
</td>
</tr>
<tr>
<td>
<code>ConnectionString</code><br/>
<em>
string
</em>
</td>
<td>
<p>ConnectionString: Application Insights component connection string.</p>
</td>
</tr>
<tr>
<td>
<code>CreationDate</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreationDate: Creation Date for the Application Insights component, in ISO 8601 format.</p>
</td>
</tr>
<tr>
<td>
<code>DisableIpMasking</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableIpMasking: Disable IP masking.</p>
</td>
</tr>
<tr>
<td>
<code>DisableLocalAuth</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableLocalAuth: Disable Non-AAD based Auth.</p>
</td>
</tr>
<tr>
<td>
<code>Flow_Type</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesStatusFlowType">
ApplicationInsightsComponentPropertiesStatusFlowType
</a>
</em>
</td>
<td>
<p>FlowType: Used by the Application Insights system to determine what kind of flow this component was created by. This is
to be set to &lsquo;Bluefield&rsquo; when creating/updating a component via the REST API.</p>
</td>
</tr>
<tr>
<td>
<code>ForceCustomerStorageForProfiler</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ForceCustomerStorageForProfiler: Force users to create their own storage account for profiler and debugger.</p>
</td>
</tr>
<tr>
<td>
<code>HockeyAppId</code><br/>
<em>
string
</em>
</td>
<td>
<p>HockeyAppId: The unique application ID created when a new application is added to HockeyApp, used for communications
with HockeyApp.</p>
</td>
</tr>
<tr>
<td>
<code>HockeyAppToken</code><br/>
<em>
string
</em>
</td>
<td>
<p>HockeyAppToken: Token used to authenticate communications with between Application Insights and HockeyApp.</p>
</td>
</tr>
<tr>
<td>
<code>ImmediatePurgeDataOn30Days</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ImmediatePurgeDataOn30Days: Purge data immediately after 30 days.</p>
</td>
</tr>
<tr>
<td>
<code>IngestionMode</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesStatusIngestionMode">
ApplicationInsightsComponentPropertiesStatusIngestionMode
</a>
</em>
</td>
<td>
<p>IngestionMode: Indicates the flow of the ingestion.</p>
</td>
</tr>
<tr>
<td>
<code>InstrumentationKey</code><br/>
<em>
string
</em>
</td>
<td>
<p>InstrumentationKey: Application Insights Instrumentation key. A read-only value that applications can use to identify
the destination for all telemetry sent to Azure Application Insights. This value will be supplied upon construction of
each new Application Insights component.</p>
</td>
</tr>
<tr>
<td>
<code>LaMigrationDate</code><br/>
<em>
string
</em>
</td>
<td>
<p>LaMigrationDate: The date which the component got migrated to LA, in ISO 8601 format.</p>
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
<p>Name: Application name.</p>
</td>
</tr>
<tr>
<td>
<code>PrivateLinkScopedResources</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.PrivateLinkScopedResource_StatusARM">
[]PrivateLinkScopedResource_StatusARM
</a>
</em>
</td>
<td>
<p>PrivateLinkScopedResources: List of linked private link scope resources.</p>
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
<p>ProvisioningState: Current state of this component: whether or not is has been provisioned within the resource group it
is defined. Users cannot change this value but are able to read from it. Values will include Succeeded, Deploying,
Canceled, and Failed.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForIngestion</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.PublicNetworkAccessType_Status">
PublicNetworkAccessType_Status
</a>
</em>
</td>
<td>
<p>PublicNetworkAccessForIngestion: The network access type for accessing Application Insights ingestion.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForQuery</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.PublicNetworkAccessType_Status">
PublicNetworkAccessType_Status
</a>
</em>
</td>
<td>
<p>PublicNetworkAccessForQuery: The network access type for accessing Application Insights query.</p>
</td>
</tr>
<tr>
<td>
<code>Request_Source</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesStatusRequestSource">
ApplicationInsightsComponentPropertiesStatusRequestSource
</a>
</em>
</td>
<td>
<p>RequestSource: Describes what tool created this Application Insights component. Customers using this API should set this
to the default &lsquo;rest&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>RetentionInDays</code><br/>
<em>
int
</em>
</td>
<td>
<p>RetentionInDays: Retention period in days.</p>
</td>
</tr>
<tr>
<td>
<code>SamplingPercentage</code><br/>
<em>
float64
</em>
</td>
<td>
<p>SamplingPercentage: Percentage of the data produced by the application being monitored that is being sampled for
Application Insights telemetry.</p>
</td>
</tr>
<tr>
<td>
<code>TenantId</code><br/>
<em>
string
</em>
</td>
<td>
<p>TenantId: Azure Tenant Id.</p>
</td>
</tr>
<tr>
<td>
<code>WorkspaceResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>WorkspaceResourceId: Resource Id of the log analytics workspace which the data will be ingested to. This property is
required to create an application with this API version. Applications from older versions will not have this property.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponent_Status">ApplicationInsightsComponent_Status
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.Component">Component</a>)
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
<code>AppId</code><br/>
<em>
string
</em>
</td>
<td>
<p>AppId: Application Insights Unique ID for your Application.</p>
</td>
</tr>
<tr>
<td>
<code>ApplicationId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ApplicationId: The unique ID of your application. This field mirrors the &lsquo;Name&rsquo; field and cannot be changed.</p>
</td>
</tr>
<tr>
<td>
<code>Application_Type</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesStatusApplicationType">
ApplicationInsightsComponentPropertiesStatusApplicationType
</a>
</em>
</td>
<td>
<p>ApplicationType: Type of application being monitored.</p>
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
<code>ConnectionString</code><br/>
<em>
string
</em>
</td>
<td>
<p>ConnectionString: Application Insights component connection string.</p>
</td>
</tr>
<tr>
<td>
<code>CreationDate</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreationDate: Creation Date for the Application Insights component, in ISO 8601 format.</p>
</td>
</tr>
<tr>
<td>
<code>DisableIpMasking</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableIpMasking: Disable IP masking.</p>
</td>
</tr>
<tr>
<td>
<code>DisableLocalAuth</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableLocalAuth: Disable Non-AAD based Auth.</p>
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
<p>Etag: Resource etag</p>
</td>
</tr>
<tr>
<td>
<code>Flow_Type</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesStatusFlowType">
ApplicationInsightsComponentPropertiesStatusFlowType
</a>
</em>
</td>
<td>
<p>FlowType: Used by the Application Insights system to determine what kind of flow this component was created by. This is
to be set to &lsquo;Bluefield&rsquo; when creating/updating a component via the REST API.</p>
</td>
</tr>
<tr>
<td>
<code>ForceCustomerStorageForProfiler</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ForceCustomerStorageForProfiler: Force users to create their own storage account for profiler and debugger.</p>
</td>
</tr>
<tr>
<td>
<code>HockeyAppId</code><br/>
<em>
string
</em>
</td>
<td>
<p>HockeyAppId: The unique application ID created when a new application is added to HockeyApp, used for communications
with HockeyApp.</p>
</td>
</tr>
<tr>
<td>
<code>HockeyAppToken</code><br/>
<em>
string
</em>
</td>
<td>
<p>HockeyAppToken: Token used to authenticate communications with between Application Insights and HockeyApp.</p>
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
<code>ImmediatePurgeDataOn30Days</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ImmediatePurgeDataOn30Days: Purge data immediately after 30 days.</p>
</td>
</tr>
<tr>
<td>
<code>IngestionMode</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesStatusIngestionMode">
ApplicationInsightsComponentPropertiesStatusIngestionMode
</a>
</em>
</td>
<td>
<p>IngestionMode: Indicates the flow of the ingestion.</p>
</td>
</tr>
<tr>
<td>
<code>InstrumentationKey</code><br/>
<em>
string
</em>
</td>
<td>
<p>InstrumentationKey: Application Insights Instrumentation key. A read-only value that applications can use to identify
the destination for all telemetry sent to Azure Application Insights. This value will be supplied upon construction of
each new Application Insights component.</p>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
string
</em>
</td>
<td>
<p>Kind: The kind of application that this component refers to, used to customize UI. This value is a freeform string,
values should typically be one of the following: web, ios, other, store, java, phone.</p>
</td>
</tr>
<tr>
<td>
<code>LaMigrationDate</code><br/>
<em>
string
</em>
</td>
<td>
<p>LaMigrationDate: The date which the component got migrated to LA, in ISO 8601 format.</p>
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
<code>PrivateLinkScopedResources</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.PrivateLinkScopedResource_Status">
[]PrivateLinkScopedResource_Status
</a>
</em>
</td>
<td>
<p>PrivateLinkScopedResources: List of linked private link scope resources.</p>
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
<p>PropertiesName: Application name.</p>
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
<p>ProvisioningState: Current state of this component: whether or not is has been provisioned within the resource group it
is defined. Users cannot change this value but are able to read from it. Values will include Succeeded, Deploying,
Canceled, and Failed.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForIngestion</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.PublicNetworkAccessType_Status">
PublicNetworkAccessType_Status
</a>
</em>
</td>
<td>
<p>PublicNetworkAccessForIngestion: The network access type for accessing Application Insights ingestion.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForQuery</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.PublicNetworkAccessType_Status">
PublicNetworkAccessType_Status
</a>
</em>
</td>
<td>
<p>PublicNetworkAccessForQuery: The network access type for accessing Application Insights query.</p>
</td>
</tr>
<tr>
<td>
<code>Request_Source</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesStatusRequestSource">
ApplicationInsightsComponentPropertiesStatusRequestSource
</a>
</em>
</td>
<td>
<p>RequestSource: Describes what tool created this Application Insights component. Customers using this API should set this
to the default &lsquo;rest&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>RetentionInDays</code><br/>
<em>
int
</em>
</td>
<td>
<p>RetentionInDays: Retention period in days.</p>
</td>
</tr>
<tr>
<td>
<code>SamplingPercentage</code><br/>
<em>
float64
</em>
</td>
<td>
<p>SamplingPercentage: Percentage of the data produced by the application being monitored that is being sampled for
Application Insights telemetry.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
Kubernetes v1.JSON
</a>
</em>
</td>
<td>
<p>Tags: Resource tags</p>
</td>
</tr>
<tr>
<td>
<code>TenantId</code><br/>
<em>
string
</em>
</td>
<td>
<p>TenantId: Azure Tenant Id.</p>
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
<code>WorkspaceResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>WorkspaceResourceId: Resource Id of the log analytics workspace which the data will be ingested to. This property is
required to create an application with this API version. Applications from older versions will not have this property.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponent_StatusARM">ApplicationInsightsComponent_StatusARM
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
<p>Etag: Resource etag</p>
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
<code>kind</code><br/>
<em>
string
</em>
</td>
<td>
<p>Kind: The kind of application that this component refers to, used to customize UI. This value is a freeform string,
values should typically be one of the following: web, ios, other, store, java, phone.</p>
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
<code>properties</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentProperties_StatusARM">
ApplicationInsightsComponentProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Properties that define an Application Insights component resource.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
Kubernetes v1.JSON
</a>
</em>
</td>
<td>
<p>Tags: Resource tags</p>
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
</tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.Component">Component
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-02-02/Microsoft.Insights.Application.json#/resourceDefinitions/components">https://schema.management.azure.com/schemas/2020-02-02/Microsoft.Insights.Application.json#/resourceDefinitions/components</a></p>
</div>
<table>
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
<a href="#insights.azure.com/v1alpha1api20200202.Components_Spec">
Components_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>Application_Type</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesApplicationType">
ApplicationInsightsComponentPropertiesApplicationType
</a>
</em>
</td>
<td>
<p>ApplicationType: Type of application being monitored.</p>
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
<code>DisableIpMasking</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableIpMasking: Disable IP masking.</p>
</td>
</tr>
<tr>
<td>
<code>DisableLocalAuth</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableLocalAuth: Disable Non-AAD based Auth.</p>
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
<p>Etag: Resource etag</p>
</td>
</tr>
<tr>
<td>
<code>Flow_Type</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesFlowType">
ApplicationInsightsComponentPropertiesFlowType
</a>
</em>
</td>
<td>
<p>FlowType: Used by the Application Insights system to determine what kind of flow this component was created by. This is
to be set to &lsquo;Bluefield&rsquo; when creating/updating a component via the REST API.</p>
</td>
</tr>
<tr>
<td>
<code>ForceCustomerStorageForProfiler</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ForceCustomerStorageForProfiler: Force users to create their own storage account for profiler and debugger.</p>
</td>
</tr>
<tr>
<td>
<code>HockeyAppId</code><br/>
<em>
string
</em>
</td>
<td>
<p>HockeyAppId: The unique application ID created when a new application is added to HockeyApp, used for communications
with HockeyApp.</p>
</td>
</tr>
<tr>
<td>
<code>ImmediatePurgeDataOn30Days</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ImmediatePurgeDataOn30Days: Purge data immediately after 30 days.</p>
</td>
</tr>
<tr>
<td>
<code>IngestionMode</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesIngestionMode">
ApplicationInsightsComponentPropertiesIngestionMode
</a>
</em>
</td>
<td>
<p>IngestionMode: Indicates the flow of the ingestion.</p>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
string
</em>
</td>
<td>
<p>Kind: The kind of application that this component refers to, used to customize UI. This value is a freeform string,
values should typically be one of the following: web, ios, other, store, java, phone.</p>
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
<p>Location: Location to deploy resource to</p>
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
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForIngestion</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesPublicNetworkAccessForIngestion">
ApplicationInsightsComponentPropertiesPublicNetworkAccessForIngestion
</a>
</em>
</td>
<td>
<p>PublicNetworkAccessForIngestion: The network access type for accessing Application Insights ingestion.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForQuery</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesPublicNetworkAccessForQuery">
ApplicationInsightsComponentPropertiesPublicNetworkAccessForQuery
</a>
</em>
</td>
<td>
<p>PublicNetworkAccessForQuery: The network access type for accessing Application Insights query.</p>
</td>
</tr>
<tr>
<td>
<code>Request_Source</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesRequestSource">
ApplicationInsightsComponentPropertiesRequestSource
</a>
</em>
</td>
<td>
<p>RequestSource: Describes what tool created this Application Insights component. Customers using this API should set this
to the default &lsquo;rest&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>RetentionInDays</code><br/>
<em>
int
</em>
</td>
<td>
<p>RetentionInDays: Retention period in days.</p>
</td>
</tr>
<tr>
<td>
<code>SamplingPercentage</code><br/>
<em>
float64
</em>
</td>
<td>
<p>SamplingPercentage: Percentage of the data produced by the application being monitored that is being sampled for
Application Insights telemetry.</p>
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
<p>Tags: Name-value pairs to add to the resource</p>
</td>
</tr>
<tr>
<td>
<code>workspaceResourceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>WorkspaceResourceReference: Resource Id of the log analytics workspace which the data will be ingested to. This property
is required to create an application with this API version. Applications from older versions will not have this property.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponent_Status">
ApplicationInsightsComponent_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.ComponentsSpecAPIVersion">ComponentsSpecAPIVersion
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
<tbody><tr><td><p>&#34;2020-02-02&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.Components_Spec">Components_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.Component">Component</a>)
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
<code>Application_Type</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesApplicationType">
ApplicationInsightsComponentPropertiesApplicationType
</a>
</em>
</td>
<td>
<p>ApplicationType: Type of application being monitored.</p>
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
<code>DisableIpMasking</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableIpMasking: Disable IP masking.</p>
</td>
</tr>
<tr>
<td>
<code>DisableLocalAuth</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableLocalAuth: Disable Non-AAD based Auth.</p>
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
<p>Etag: Resource etag</p>
</td>
</tr>
<tr>
<td>
<code>Flow_Type</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesFlowType">
ApplicationInsightsComponentPropertiesFlowType
</a>
</em>
</td>
<td>
<p>FlowType: Used by the Application Insights system to determine what kind of flow this component was created by. This is
to be set to &lsquo;Bluefield&rsquo; when creating/updating a component via the REST API.</p>
</td>
</tr>
<tr>
<td>
<code>ForceCustomerStorageForProfiler</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ForceCustomerStorageForProfiler: Force users to create their own storage account for profiler and debugger.</p>
</td>
</tr>
<tr>
<td>
<code>HockeyAppId</code><br/>
<em>
string
</em>
</td>
<td>
<p>HockeyAppId: The unique application ID created when a new application is added to HockeyApp, used for communications
with HockeyApp.</p>
</td>
</tr>
<tr>
<td>
<code>ImmediatePurgeDataOn30Days</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ImmediatePurgeDataOn30Days: Purge data immediately after 30 days.</p>
</td>
</tr>
<tr>
<td>
<code>IngestionMode</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesIngestionMode">
ApplicationInsightsComponentPropertiesIngestionMode
</a>
</em>
</td>
<td>
<p>IngestionMode: Indicates the flow of the ingestion.</p>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
string
</em>
</td>
<td>
<p>Kind: The kind of application that this component refers to, used to customize UI. This value is a freeform string,
values should typically be one of the following: web, ios, other, store, java, phone.</p>
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
<p>Location: Location to deploy resource to</p>
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
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForIngestion</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesPublicNetworkAccessForIngestion">
ApplicationInsightsComponentPropertiesPublicNetworkAccessForIngestion
</a>
</em>
</td>
<td>
<p>PublicNetworkAccessForIngestion: The network access type for accessing Application Insights ingestion.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForQuery</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesPublicNetworkAccessForQuery">
ApplicationInsightsComponentPropertiesPublicNetworkAccessForQuery
</a>
</em>
</td>
<td>
<p>PublicNetworkAccessForQuery: The network access type for accessing Application Insights query.</p>
</td>
</tr>
<tr>
<td>
<code>Request_Source</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesRequestSource">
ApplicationInsightsComponentPropertiesRequestSource
</a>
</em>
</td>
<td>
<p>RequestSource: Describes what tool created this Application Insights component. Customers using this API should set this
to the default &lsquo;rest&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>RetentionInDays</code><br/>
<em>
int
</em>
</td>
<td>
<p>RetentionInDays: Retention period in days.</p>
</td>
</tr>
<tr>
<td>
<code>SamplingPercentage</code><br/>
<em>
float64
</em>
</td>
<td>
<p>SamplingPercentage: Percentage of the data produced by the application being monitored that is being sampled for
Application Insights telemetry.</p>
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
<p>Tags: Name-value pairs to add to the resource</p>
</td>
</tr>
<tr>
<td>
<code>workspaceResourceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>WorkspaceResourceReference: Resource Id of the log analytics workspace which the data will be ingested to. This property
is required to create an application with this API version. Applications from older versions will not have this property.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.Components_SpecARM">Components_SpecARM
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
<p>Etag: Resource etag</p>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
string
</em>
</td>
<td>
<p>Kind: The kind of application that this component refers to, used to customize UI. This value is a freeform string,
values should typically be one of the following: web, ios, other, store, java, phone.</p>
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
<p>Location: Location to deploy resource to</p>
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
<p>Name: The name of the Application Insights component resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesARM">
ApplicationInsightsComponentPropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties that define an Application Insights component resource.</p>
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
<p>Tags: Name-value pairs to add to the resource</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.PrivateLinkScopedResource_Status">PrivateLinkScopedResource_Status
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponent_Status">ApplicationInsightsComponent_Status</a>)
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
<code>ResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ResourceId: The full resource Id of the private link scope resource.</p>
</td>
</tr>
<tr>
<td>
<code>ScopeId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScopeId: The private link scope unique Identifier.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.PrivateLinkScopedResource_StatusARM">PrivateLinkScopedResource_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentProperties_StatusARM">ApplicationInsightsComponentProperties_StatusARM</a>)
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
<code>ResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ResourceId: The full resource Id of the private link scope resource.</p>
</td>
</tr>
<tr>
<td>
<code>ScopeId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScopeId: The private link scope unique Identifier.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.PublicNetworkAccessType_Status">PublicNetworkAccessType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentProperties_StatusARM">ApplicationInsightsComponentProperties_StatusARM</a>, <a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponent_Status">ApplicationInsightsComponent_Status</a>)
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
<hr/>
