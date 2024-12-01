---
title: insights.azure.com/v1api20210501preview
---
<h2 id="insights.azure.com/v1api20210501preview">insights.azure.com/v1api20210501preview</h2>
<div>
<p>Package v1api20210501preview contains API Schema definitions for the insights v1api20210501preview API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="insights.azure.com/v1api20210501preview.APIVersion">APIVersion
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
<tbody><tr><td><p>&#34;2021-05-01-preview&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1api20210501preview.DiagnosticSetting">DiagnosticSetting
</h3>
<div>
<p>Generator information:
- Generated from: /monitor/resource-manager/Microsoft.Insights/preview/2021-05-01-preview/diagnosticsSettings_API.json
- ARM URI: /{resourceUri}/providers/Microsoft.Insights/diagnosticSettings/{name}</p>
</div>
<table>
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
<a href="#insights.azure.com/v1api20210501preview.DiagnosticSetting_Spec">
DiagnosticSetting_Spec
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
<code>eventHubAuthorizationRuleReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>EventHubAuthorizationRuleReference: The resource Id for the event hub authorization rule.</p>
</td>
</tr>
<tr>
<td>
<code>eventHubName</code><br/>
<em>
string
</em>
</td>
<td>
<p>EventHubName: The name of the event hub. If none is specified, the default event hub will be selected.</p>
</td>
</tr>
<tr>
<td>
<code>logAnalyticsDestinationType</code><br/>
<em>
string
</em>
</td>
<td>
<p>LogAnalyticsDestinationType: A string indicating whether the export to Log Analytics should use the default destination
type, i.e. AzureDiagnostics, or use a destination type constructed as follows: <normalized service identity>_<normalized
category name>. Possible values are: Dedicated and null (null is default.)</p>
</td>
</tr>
<tr>
<td>
<code>logs</code><br/>
<em>
<a href="#insights.azure.com/v1api20210501preview.LogSettings">
[]LogSettings
</a>
</em>
</td>
<td>
<p>Logs: The list of logs settings.</p>
</td>
</tr>
<tr>
<td>
<code>marketplacePartnerReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>MarketplacePartnerReference: The full ARM resource ID of the Marketplace resource to which you would like to send
Diagnostic Logs.</p>
</td>
</tr>
<tr>
<td>
<code>metrics</code><br/>
<em>
<a href="#insights.azure.com/v1api20210501preview.MetricSettings">
[]MetricSettings
</a>
</em>
</td>
<td>
<p>Metrics: The list of metric settings.</p>
</td>
</tr>
<tr>
<td>
<code>operatorSpec</code><br/>
<em>
<a href="#insights.azure.com/v1api20210501preview.DiagnosticSettingOperatorSpec">
DiagnosticSettingOperatorSpec
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
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ArbitraryOwnerReference">
genruntime.ArbitraryOwnerReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. This resource is an
extension resource, which means that any other Azure resource can be its owner.</p>
</td>
</tr>
<tr>
<td>
<code>serviceBusRuleId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServiceBusRuleId: The service bus rule Id of the diagnostic setting. This is here to maintain backwards compatibility.</p>
</td>
</tr>
<tr>
<td>
<code>storageAccountReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>StorageAccountReference: The resource ID of the storage account to which you would like to send Diagnostic Logs.</p>
</td>
</tr>
<tr>
<td>
<code>workspaceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>WorkspaceReference: The full ARM resource ID of the Log Analytics workspace to which you would like to send Diagnostic
Logs. Example:
/&#x200b;subscriptions/&#x200b;4b9e8510-67ab-4e9a-95a9-e2f1e570ea9c/&#x200b;resourceGroups/&#x200b;insights-integration/&#x200b;providers/&#x200b;Microsoft.OperationalInsights/&#x200b;workspaces/&#x200b;viruela2</&#x200b;p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#insights.azure.com/v1api20210501preview.DiagnosticSetting_STATUS">
DiagnosticSetting_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20210501preview.DiagnosticSettingOperatorSpec">DiagnosticSettingOperatorSpec
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20210501preview.DiagnosticSetting_Spec">DiagnosticSetting_Spec</a>)
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
<h3 id="insights.azure.com/v1api20210501preview.DiagnosticSetting_STATUS">DiagnosticSetting_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20210501preview.DiagnosticSetting">DiagnosticSetting</a>)
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
<code>eventHubAuthorizationRuleId</code><br/>
<em>
string
</em>
</td>
<td>
<p>EventHubAuthorizationRuleId: The resource Id for the event hub authorization rule.</p>
</td>
</tr>
<tr>
<td>
<code>eventHubName</code><br/>
<em>
string
</em>
</td>
<td>
<p>EventHubName: The name of the event hub. If none is specified, the default event hub will be selected.</p>
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
<code>logAnalyticsDestinationType</code><br/>
<em>
string
</em>
</td>
<td>
<p>LogAnalyticsDestinationType: A string indicating whether the export to Log Analytics should use the default destination
type, i.e. AzureDiagnostics, or use a destination type constructed as follows: <normalized service identity>_<normalized
category name>. Possible values are: Dedicated and null (null is default.)</p>
</td>
</tr>
<tr>
<td>
<code>logs</code><br/>
<em>
<a href="#insights.azure.com/v1api20210501preview.LogSettings_STATUS">
[]LogSettings_STATUS
</a>
</em>
</td>
<td>
<p>Logs: The list of logs settings.</p>
</td>
</tr>
<tr>
<td>
<code>marketplacePartnerId</code><br/>
<em>
string
</em>
</td>
<td>
<p>MarketplacePartnerId: The full ARM resource ID of the Marketplace resource to which you would like to send Diagnostic
Logs.</p>
</td>
</tr>
<tr>
<td>
<code>metrics</code><br/>
<em>
<a href="#insights.azure.com/v1api20210501preview.MetricSettings_STATUS">
[]MetricSettings_STATUS
</a>
</em>
</td>
<td>
<p>Metrics: The list of metric settings.</p>
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
<code>serviceBusRuleId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServiceBusRuleId: The service bus rule Id of the diagnostic setting. This is here to maintain backwards compatibility.</p>
</td>
</tr>
<tr>
<td>
<code>storageAccountId</code><br/>
<em>
string
</em>
</td>
<td>
<p>StorageAccountId: The resource ID of the storage account to which you would like to send Diagnostic Logs.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#insights.azure.com/v1api20210501preview.SystemData_STATUS">
SystemData_STATUS
</a>
</em>
</td>
<td>
<p>SystemData: The system metadata related to this resource.</p>
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
<code>workspaceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>WorkspaceId: The full ARM resource ID of the Log Analytics workspace to which you would like to send Diagnostic Logs.
Example:
/&#x200b;subscriptions/&#x200b;4b9e8510-67ab-4e9a-95a9-e2f1e570ea9c/&#x200b;resourceGroups/&#x200b;insights-integration/&#x200b;providers/&#x200b;Microsoft.OperationalInsights/&#x200b;workspaces/&#x200b;viruela2</&#x200b;p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20210501preview.DiagnosticSetting_Spec">DiagnosticSetting_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20210501preview.DiagnosticSetting">DiagnosticSetting</a>)
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
<code>eventHubAuthorizationRuleReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>EventHubAuthorizationRuleReference: The resource Id for the event hub authorization rule.</p>
</td>
</tr>
<tr>
<td>
<code>eventHubName</code><br/>
<em>
string
</em>
</td>
<td>
<p>EventHubName: The name of the event hub. If none is specified, the default event hub will be selected.</p>
</td>
</tr>
<tr>
<td>
<code>logAnalyticsDestinationType</code><br/>
<em>
string
</em>
</td>
<td>
<p>LogAnalyticsDestinationType: A string indicating whether the export to Log Analytics should use the default destination
type, i.e. AzureDiagnostics, or use a destination type constructed as follows: <normalized service identity>_<normalized
category name>. Possible values are: Dedicated and null (null is default.)</p>
</td>
</tr>
<tr>
<td>
<code>logs</code><br/>
<em>
<a href="#insights.azure.com/v1api20210501preview.LogSettings">
[]LogSettings
</a>
</em>
</td>
<td>
<p>Logs: The list of logs settings.</p>
</td>
</tr>
<tr>
<td>
<code>marketplacePartnerReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>MarketplacePartnerReference: The full ARM resource ID of the Marketplace resource to which you would like to send
Diagnostic Logs.</p>
</td>
</tr>
<tr>
<td>
<code>metrics</code><br/>
<em>
<a href="#insights.azure.com/v1api20210501preview.MetricSettings">
[]MetricSettings
</a>
</em>
</td>
<td>
<p>Metrics: The list of metric settings.</p>
</td>
</tr>
<tr>
<td>
<code>operatorSpec</code><br/>
<em>
<a href="#insights.azure.com/v1api20210501preview.DiagnosticSettingOperatorSpec">
DiagnosticSettingOperatorSpec
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
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ArbitraryOwnerReference">
genruntime.ArbitraryOwnerReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. This resource is an
extension resource, which means that any other Azure resource can be its owner.</p>
</td>
</tr>
<tr>
<td>
<code>serviceBusRuleId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServiceBusRuleId: The service bus rule Id of the diagnostic setting. This is here to maintain backwards compatibility.</p>
</td>
</tr>
<tr>
<td>
<code>storageAccountReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>StorageAccountReference: The resource ID of the storage account to which you would like to send Diagnostic Logs.</p>
</td>
</tr>
<tr>
<td>
<code>workspaceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>WorkspaceReference: The full ARM resource ID of the Log Analytics workspace to which you would like to send Diagnostic
Logs. Example:
/&#x200b;subscriptions/&#x200b;4b9e8510-67ab-4e9a-95a9-e2f1e570ea9c/&#x200b;resourceGroups/&#x200b;insights-integration/&#x200b;providers/&#x200b;Microsoft.OperationalInsights/&#x200b;workspaces/&#x200b;viruela2</&#x200b;p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20210501preview.LogSettings">LogSettings
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20210501preview.DiagnosticSetting_Spec">DiagnosticSetting_Spec</a>)
</p>
<div>
<p>Part of MultiTenantDiagnosticSettings. Specifies the settings for a particular log.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>category</code><br/>
<em>
string
</em>
</td>
<td>
<p>Category: Name of a Diagnostic Log category for a resource type this setting is applied to. To obtain the list of
Diagnostic Log categories for a resource, first perform a GET diagnostic settings operation.</p>
</td>
</tr>
<tr>
<td>
<code>categoryGroup</code><br/>
<em>
string
</em>
</td>
<td>
<p>CategoryGroup: Name of a Diagnostic Log category group for a resource type this setting is applied to. To obtain the
list of Diagnostic Log categories for a resource, first perform a GET diagnostic settings operation.</p>
</td>
</tr>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: a value indicating whether this log is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>retentionPolicy</code><br/>
<em>
<a href="#insights.azure.com/v1api20210501preview.RetentionPolicy">
RetentionPolicy
</a>
</em>
</td>
<td>
<p>RetentionPolicy: the retention policy for this log.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20210501preview.LogSettings_STATUS">LogSettings_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20210501preview.DiagnosticSetting_STATUS">DiagnosticSetting_STATUS</a>)
</p>
<div>
<p>Part of MultiTenantDiagnosticSettings. Specifies the settings for a particular log.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>category</code><br/>
<em>
string
</em>
</td>
<td>
<p>Category: Name of a Diagnostic Log category for a resource type this setting is applied to. To obtain the list of
Diagnostic Log categories for a resource, first perform a GET diagnostic settings operation.</p>
</td>
</tr>
<tr>
<td>
<code>categoryGroup</code><br/>
<em>
string
</em>
</td>
<td>
<p>CategoryGroup: Name of a Diagnostic Log category group for a resource type this setting is applied to. To obtain the
list of Diagnostic Log categories for a resource, first perform a GET diagnostic settings operation.</p>
</td>
</tr>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: a value indicating whether this log is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>retentionPolicy</code><br/>
<em>
<a href="#insights.azure.com/v1api20210501preview.RetentionPolicy_STATUS">
RetentionPolicy_STATUS
</a>
</em>
</td>
<td>
<p>RetentionPolicy: the retention policy for this log.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20210501preview.MetricSettings">MetricSettings
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20210501preview.DiagnosticSetting_Spec">DiagnosticSetting_Spec</a>)
</p>
<div>
<p>Part of MultiTenantDiagnosticSettings. Specifies the settings for a particular metric.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>category</code><br/>
<em>
string
</em>
</td>
<td>
<p>Category: Name of a Diagnostic Metric category for a resource type this setting is applied to. To obtain the list of
Diagnostic metric categories for a resource, first perform a GET diagnostic settings operation.</p>
</td>
</tr>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: a value indicating whether this category is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>retentionPolicy</code><br/>
<em>
<a href="#insights.azure.com/v1api20210501preview.RetentionPolicy">
RetentionPolicy
</a>
</em>
</td>
<td>
<p>RetentionPolicy: the retention policy for this category.</p>
</td>
</tr>
<tr>
<td>
<code>timeGrain</code><br/>
<em>
string
</em>
</td>
<td>
<p>TimeGrain: the timegrain of the metric in ISO8601 format.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20210501preview.MetricSettings_STATUS">MetricSettings_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20210501preview.DiagnosticSetting_STATUS">DiagnosticSetting_STATUS</a>)
</p>
<div>
<p>Part of MultiTenantDiagnosticSettings. Specifies the settings for a particular metric.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>category</code><br/>
<em>
string
</em>
</td>
<td>
<p>Category: Name of a Diagnostic Metric category for a resource type this setting is applied to. To obtain the list of
Diagnostic metric categories for a resource, first perform a GET diagnostic settings operation.</p>
</td>
</tr>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: a value indicating whether this category is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>retentionPolicy</code><br/>
<em>
<a href="#insights.azure.com/v1api20210501preview.RetentionPolicy_STATUS">
RetentionPolicy_STATUS
</a>
</em>
</td>
<td>
<p>RetentionPolicy: the retention policy for this category.</p>
</td>
</tr>
<tr>
<td>
<code>timeGrain</code><br/>
<em>
string
</em>
</td>
<td>
<p>TimeGrain: the timegrain of the metric in ISO8601 format.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20210501preview.RetentionPolicy">RetentionPolicy
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20210501preview.LogSettings">LogSettings</a>, <a href="#insights.azure.com/v1api20210501preview.MetricSettings">MetricSettings</a>)
</p>
<div>
<p>Specifies the retention policy for the log.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>days</code><br/>
<em>
int
</em>
</td>
<td>
<p>Days: the number of days for the retention in days. A value of 0 will retain the events indefinitely.</p>
</td>
</tr>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: a value indicating whether the retention policy is enabled.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20210501preview.RetentionPolicy_STATUS">RetentionPolicy_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20210501preview.LogSettings_STATUS">LogSettings_STATUS</a>, <a href="#insights.azure.com/v1api20210501preview.MetricSettings_STATUS">MetricSettings_STATUS</a>)
</p>
<div>
<p>Specifies the retention policy for the log.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>days</code><br/>
<em>
int
</em>
</td>
<td>
<p>Days: the number of days for the retention in days. A value of 0 will retain the events indefinitely.</p>
</td>
</tr>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: a value indicating whether the retention policy is enabled.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20210501preview.SystemData_CreatedByType_STATUS">SystemData_CreatedByType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20210501preview.SystemData_STATUS">SystemData_STATUS</a>)
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
<h3 id="insights.azure.com/v1api20210501preview.SystemData_LastModifiedByType_STATUS">SystemData_LastModifiedByType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20210501preview.SystemData_STATUS">SystemData_STATUS</a>)
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
<h3 id="insights.azure.com/v1api20210501preview.SystemData_STATUS">SystemData_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20210501preview.DiagnosticSetting_STATUS">DiagnosticSetting_STATUS</a>)
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
<a href="#insights.azure.com/v1api20210501preview.SystemData_CreatedByType_STATUS">
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
<a href="#insights.azure.com/v1api20210501preview.SystemData_LastModifiedByType_STATUS">
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
<hr/>
