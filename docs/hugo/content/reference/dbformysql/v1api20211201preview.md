---
title: dbformysql.azure.com/v1api20211201preview
---
<h2 id="dbformysql.azure.com/v1api20211201preview">dbformysql.azure.com/v1api20211201preview</h2>
<div>
<p>Package v1api20211201preview contains API Schema definitions for the dbformysql v1api20211201preview API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="dbformysql.azure.com/v1api20211201preview.APIVersion">APIVersion
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
<tbody><tr><td><p>&#34;2021-12-01-preview&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbformysql.azure.com/v1api20211201preview.AdministratorProperties_ARM">AdministratorProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1api20211201preview.FlexibleServers_Administrator_Spec_ARM">FlexibleServers_Administrator_Spec_ARM</a>)
</p>
<div>
<p>The properties of an administrator.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>administratorType</code><br/>
<em>
<a href="#dbformysql.azure.com/v1api20211201preview.AdministratorProperties_AdministratorType">
AdministratorProperties_AdministratorType
</a>
</em>
</td>
<td>
<p>AdministratorType: Type of the sever administrator.</p>
</td>
</tr>
<tr>
<td>
<code>identityResourceId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>login</code><br/>
<em>
string
</em>
</td>
<td>
<p>Login: Login name of the server administrator.</p>
</td>
</tr>
<tr>
<td>
<code>sid</code><br/>
<em>
string
</em>
</td>
<td>
<p>Sid: SID (object ID) of the server administrator.</p>
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
<p>TenantId: Tenant ID of the administrator.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1api20211201preview.AdministratorProperties_AdministratorType">AdministratorProperties_AdministratorType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1api20211201preview.AdministratorProperties_ARM">AdministratorProperties_ARM</a>, <a href="#dbformysql.azure.com/v1api20211201preview.FlexibleServers_Administrator_Spec">FlexibleServers_Administrator_Spec</a>)
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
<tbody><tr><td><p>&#34;ActiveDirectory&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbformysql.azure.com/v1api20211201preview.AdministratorProperties_AdministratorType_STATUS">AdministratorProperties_AdministratorType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1api20211201preview.AdministratorProperties_STATUS_ARM">AdministratorProperties_STATUS_ARM</a>, <a href="#dbformysql.azure.com/v1api20211201preview.FlexibleServers_Administrator_STATUS">FlexibleServers_Administrator_STATUS</a>)
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
<tbody><tr><td><p>&#34;ActiveDirectory&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbformysql.azure.com/v1api20211201preview.AdministratorProperties_STATUS_ARM">AdministratorProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1api20211201preview.FlexibleServers_Administrator_STATUS_ARM">FlexibleServers_Administrator_STATUS_ARM</a>)
</p>
<div>
<p>The properties of an administrator.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>administratorType</code><br/>
<em>
<a href="#dbformysql.azure.com/v1api20211201preview.AdministratorProperties_AdministratorType_STATUS">
AdministratorProperties_AdministratorType_STATUS
</a>
</em>
</td>
<td>
<p>AdministratorType: Type of the sever administrator.</p>
</td>
</tr>
<tr>
<td>
<code>identityResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>IdentityResourceId: The resource id of the identity used for AAD Authentication.</p>
</td>
</tr>
<tr>
<td>
<code>login</code><br/>
<em>
string
</em>
</td>
<td>
<p>Login: Login name of the server administrator.</p>
</td>
</tr>
<tr>
<td>
<code>sid</code><br/>
<em>
string
</em>
</td>
<td>
<p>Sid: SID (object ID) of the server administrator.</p>
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
<p>TenantId: Tenant ID of the administrator.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1api20211201preview.FlexibleServersAdministrator">FlexibleServersAdministrator
</h3>
<div>
<p>Generator information:
- Generated from: /mysql/resource-manager/Microsoft.DBforMySQL/AAD/preview/2021-12-01-preview/AzureADAdministrator.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.DBforMySQL/&#x200b;flexibleServers/&#x200b;{serverName}/&#x200b;administrators/&#x200b;{administratorName}</&#x200b;p>
</div>
<table>
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
<a href="#dbformysql.azure.com/v1api20211201preview.FlexibleServers_Administrator_Spec">
FlexibleServers_Administrator_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>administratorType</code><br/>
<em>
<a href="#dbformysql.azure.com/v1api20211201preview.AdministratorProperties_AdministratorType">
AdministratorProperties_AdministratorType
</a>
</em>
</td>
<td>
<p>AdministratorType: Type of the sever administrator.</p>
</td>
</tr>
<tr>
<td>
<code>identityResourceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>IdentityResourceReference: The resource id of the identity used for AAD Authentication.</p>
</td>
</tr>
<tr>
<td>
<code>login</code><br/>
<em>
string
</em>
</td>
<td>
<p>Login: Login name of the server administrator.</p>
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
reference to a dbformysql.azure.com/FlexibleServer resource</p>
</td>
</tr>
<tr>
<td>
<code>sid</code><br/>
<em>
string
</em>
</td>
<td>
<p>Sid: SID (object ID) of the server administrator.</p>
</td>
</tr>
<tr>
<td>
<code>sidFromConfig</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ConfigMapReference">
genruntime.ConfigMapReference
</a>
</em>
</td>
<td>
<p>SidFromConfig: SID (object ID) of the server administrator.</p>
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
<p>TenantId: Tenant ID of the administrator.</p>
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
<p>TenantIdFromConfig: Tenant ID of the administrator.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#dbformysql.azure.com/v1api20211201preview.FlexibleServers_Administrator_STATUS">
FlexibleServers_Administrator_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1api20211201preview.FlexibleServers_Administrator_STATUS">FlexibleServers_Administrator_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1api20211201preview.FlexibleServersAdministrator">FlexibleServersAdministrator</a>)
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
<code>administratorType</code><br/>
<em>
<a href="#dbformysql.azure.com/v1api20211201preview.AdministratorProperties_AdministratorType_STATUS">
AdministratorProperties_AdministratorType_STATUS
</a>
</em>
</td>
<td>
<p>AdministratorType: Type of the sever administrator.</p>
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
<code>identityResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>IdentityResourceId: The resource id of the identity used for AAD Authentication.</p>
</td>
</tr>
<tr>
<td>
<code>login</code><br/>
<em>
string
</em>
</td>
<td>
<p>Login: Login name of the server administrator.</p>
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
<code>sid</code><br/>
<em>
string
</em>
</td>
<td>
<p>Sid: SID (object ID) of the server administrator.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#dbformysql.azure.com/v1api20211201preview.SystemData_STATUS">
SystemData_STATUS
</a>
</em>
</td>
<td>
<p>SystemData: The system metadata relating to this resource.</p>
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
<p>TenantId: Tenant ID of the administrator.</p>
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
<h3 id="dbformysql.azure.com/v1api20211201preview.FlexibleServers_Administrator_STATUS_ARM">FlexibleServers_Administrator_STATUS_ARM
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
<a href="#dbformysql.azure.com/v1api20211201preview.AdministratorProperties_STATUS_ARM">
AdministratorProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: The properties of an administrator.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#dbformysql.azure.com/v1api20211201preview.SystemData_STATUS_ARM">
SystemData_STATUS_ARM
</a>
</em>
</td>
<td>
<p>SystemData: The system metadata relating to this resource.</p>
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
<h3 id="dbformysql.azure.com/v1api20211201preview.FlexibleServers_Administrator_Spec">FlexibleServers_Administrator_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1api20211201preview.FlexibleServersAdministrator">FlexibleServersAdministrator</a>)
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
<code>administratorType</code><br/>
<em>
<a href="#dbformysql.azure.com/v1api20211201preview.AdministratorProperties_AdministratorType">
AdministratorProperties_AdministratorType
</a>
</em>
</td>
<td>
<p>AdministratorType: Type of the sever administrator.</p>
</td>
</tr>
<tr>
<td>
<code>identityResourceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>IdentityResourceReference: The resource id of the identity used for AAD Authentication.</p>
</td>
</tr>
<tr>
<td>
<code>login</code><br/>
<em>
string
</em>
</td>
<td>
<p>Login: Login name of the server administrator.</p>
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
reference to a dbformysql.azure.com/FlexibleServer resource</p>
</td>
</tr>
<tr>
<td>
<code>sid</code><br/>
<em>
string
</em>
</td>
<td>
<p>Sid: SID (object ID) of the server administrator.</p>
</td>
</tr>
<tr>
<td>
<code>sidFromConfig</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ConfigMapReference">
genruntime.ConfigMapReference
</a>
</em>
</td>
<td>
<p>SidFromConfig: SID (object ID) of the server administrator.</p>
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
<p>TenantId: Tenant ID of the administrator.</p>
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
<p>TenantIdFromConfig: Tenant ID of the administrator.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1api20211201preview.FlexibleServers_Administrator_Spec_ARM">FlexibleServers_Administrator_Spec_ARM
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
<a href="#dbformysql.azure.com/v1api20211201preview.AdministratorProperties_ARM">
AdministratorProperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: The properties of an administrator.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1api20211201preview.SystemData_CreatedByType_STATUS">SystemData_CreatedByType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1api20211201preview.SystemData_STATUS">SystemData_STATUS</a>, <a href="#dbformysql.azure.com/v1api20211201preview.SystemData_STATUS_ARM">SystemData_STATUS_ARM</a>)
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
<h3 id="dbformysql.azure.com/v1api20211201preview.SystemData_LastModifiedByType_STATUS">SystemData_LastModifiedByType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1api20211201preview.SystemData_STATUS">SystemData_STATUS</a>, <a href="#dbformysql.azure.com/v1api20211201preview.SystemData_STATUS_ARM">SystemData_STATUS_ARM</a>)
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
<h3 id="dbformysql.azure.com/v1api20211201preview.SystemData_STATUS">SystemData_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1api20211201preview.FlexibleServers_Administrator_STATUS">FlexibleServers_Administrator_STATUS</a>)
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
<a href="#dbformysql.azure.com/v1api20211201preview.SystemData_CreatedByType_STATUS">
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
<a href="#dbformysql.azure.com/v1api20211201preview.SystemData_LastModifiedByType_STATUS">
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
<h3 id="dbformysql.azure.com/v1api20211201preview.SystemData_STATUS_ARM">SystemData_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1api20211201preview.FlexibleServers_Administrator_STATUS_ARM">FlexibleServers_Administrator_STATUS_ARM</a>)
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
<a href="#dbformysql.azure.com/v1api20211201preview.SystemData_CreatedByType_STATUS">
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
<a href="#dbformysql.azure.com/v1api20211201preview.SystemData_LastModifiedByType_STATUS">
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
