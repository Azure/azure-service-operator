---
title: keyvault.azure.com/v1api20210401preview
---
<h2 id="keyvault.azure.com/v1api20210401preview">keyvault.azure.com/v1api20210401preview</h2>
<div>
<p>Package v1api20210401preview contains API Schema definitions for the keyvault v1api20210401preview API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="keyvault.azure.com/v1api20210401preview.APIVersion">APIVersion
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
<tbody><tr><td><p>&#34;2021-04-01-preview&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.AccessPolicyEntry">AccessPolicyEntry
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.VaultProperties">VaultProperties</a>)
</p>
<div>
<p>An identity that have access to the key vault. All identities in the array must use the same tenant ID as the key
vault&rsquo;s tenant ID.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>applicationId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ApplicationId:  Application ID of the client making request on behalf of a principal</p>
</td>
</tr>
<tr>
<td>
<code>applicationIdFromConfig</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ConfigMapReference">
genruntime.ConfigMapReference
</a>
</em>
</td>
<td>
<p>ApplicationIdFromConfig:  Application ID of the client making request on behalf of a principal</p>
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
<p>ObjectId: The object ID of a user, service principal or security group in the Azure Active Directory tenant for the
vault. The object ID must be unique for the list of access policies.</p>
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
<p>ObjectIdFromConfig: The object ID of a user, service principal or security group in the Azure Active Directory tenant
for the vault. The object ID must be unique for the list of access policies.</p>
</td>
</tr>
<tr>
<td>
<code>permissions</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Permissions">
Permissions
</a>
</em>
</td>
<td>
<p>Permissions: Permissions the identity has for keys, secrets and certificates.</p>
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
<p>TenantId: The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.</p>
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
<p>TenantIdFromConfig: The Azure Active Directory tenant ID that should be used for authenticating requests to the key
vault.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.AccessPolicyEntry_ARM">AccessPolicyEntry_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_ARM">VaultProperties_ARM</a>)
</p>
<div>
<p>An identity that have access to the key vault. All identities in the array must use the same tenant ID as the key
vault&rsquo;s tenant ID.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>applicationId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ApplicationId:  Application ID of the client making request on behalf of a principal</p>
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
<p>ObjectId: The object ID of a user, service principal or security group in the Azure Active Directory tenant for the
vault. The object ID must be unique for the list of access policies.</p>
</td>
</tr>
<tr>
<td>
<code>permissions</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Permissions_ARM">
Permissions_ARM
</a>
</em>
</td>
<td>
<p>Permissions: Permissions the identity has for keys, secrets and certificates.</p>
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
<p>TenantId: The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.AccessPolicyEntry_STATUS">AccessPolicyEntry_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_STATUS">VaultProperties_STATUS</a>)
</p>
<div>
<p>An identity that have access to the key vault. All identities in the array must use the same tenant ID as the key
vault&rsquo;s tenant ID.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>applicationId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ApplicationId:  Application ID of the client making request on behalf of a principal</p>
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
<p>ObjectId: The object ID of a user, service principal or security group in the Azure Active Directory tenant for the
vault. The object ID must be unique for the list of access policies.</p>
</td>
</tr>
<tr>
<td>
<code>permissions</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Permissions_STATUS">
Permissions_STATUS
</a>
</em>
</td>
<td>
<p>Permissions: Permissions the identity has for keys, secrets and certificates.</p>
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
<p>TenantId: The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.AccessPolicyEntry_STATUS_ARM">AccessPolicyEntry_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_STATUS_ARM">VaultProperties_STATUS_ARM</a>)
</p>
<div>
<p>An identity that have access to the key vault. All identities in the array must use the same tenant ID as the key
vault&rsquo;s tenant ID.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>applicationId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ApplicationId:  Application ID of the client making request on behalf of a principal</p>
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
<p>ObjectId: The object ID of a user, service principal or security group in the Azure Active Directory tenant for the
vault. The object ID must be unique for the list of access policies.</p>
</td>
</tr>
<tr>
<td>
<code>permissions</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Permissions_STATUS_ARM">
Permissions_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Permissions: Permissions the identity has for keys, secrets and certificates.</p>
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
<p>TenantId: The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.IPRule">IPRule
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet">NetworkRuleSet</a>)
</p>
<div>
<p>A rule governing the accessibility of a vault from a specific ip address or ip range.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: An IPv4 address range in CIDR notation, such as &lsquo;124.56.78.91&rsquo; (simple IP address) or &lsquo;124.56.78.0/24&rsquo; (all
addresses that start with 124.56.78).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.IPRule_ARM">IPRule_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet_ARM">NetworkRuleSet_ARM</a>)
</p>
<div>
<p>A rule governing the accessibility of a vault from a specific ip address or ip range.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: An IPv4 address range in CIDR notation, such as &lsquo;124.56.78.91&rsquo; (simple IP address) or &lsquo;124.56.78.0/24&rsquo; (all
addresses that start with 124.56.78).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.IPRule_STATUS">IPRule_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet_STATUS">NetworkRuleSet_STATUS</a>)
</p>
<div>
<p>A rule governing the accessibility of a vault from a specific ip address or ip range.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: An IPv4 address range in CIDR notation, such as &lsquo;124.56.78.91&rsquo; (simple IP address) or &lsquo;124.56.78.0/24&rsquo; (all
addresses that start with 124.56.78).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.IPRule_STATUS_ARM">IPRule_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet_STATUS_ARM">NetworkRuleSet_STATUS_ARM</a>)
</p>
<div>
<p>A rule governing the accessibility of a vault from a specific ip address or ip range.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: An IPv4 address range in CIDR notation, such as &lsquo;124.56.78.91&rsquo; (simple IP address) or &lsquo;124.56.78.0/24&rsquo; (all
addresses that start with 124.56.78).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.IdentityType_STATUS">IdentityType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.SystemData_STATUS">SystemData_STATUS</a>, <a href="#keyvault.azure.com/v1api20210401preview.SystemData_STATUS_ARM">SystemData_STATUS_ARM</a>)
</p>
<div>
<p>The type of identity.</p>
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
<h3 id="keyvault.azure.com/v1api20210401preview.NetworkRuleSet">NetworkRuleSet
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.VaultProperties">VaultProperties</a>)
</p>
<div>
<p>A set of rules governing the network accessibility of a vault.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>bypass</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet_Bypass">
NetworkRuleSet_Bypass
</a>
</em>
</td>
<td>
<p>Bypass: Tells what traffic can bypass network rules. This can be &lsquo;AzureServices&rsquo; or &lsquo;None&rsquo;.  If not specified the
default is &lsquo;AzureServices&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>defaultAction</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet_DefaultAction">
NetworkRuleSet_DefaultAction
</a>
</em>
</td>
<td>
<p>DefaultAction: The default action when no rule from ipRules and from virtualNetworkRules match. This is only used after
the bypass property has been evaluated.</p>
</td>
</tr>
<tr>
<td>
<code>ipRules</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.IPRule">
[]IPRule
</a>
</em>
</td>
<td>
<p>IpRules: The list of IP address rules.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkRules</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.VirtualNetworkRule">
[]VirtualNetworkRule
</a>
</em>
</td>
<td>
<p>VirtualNetworkRules: The list of virtual network rules.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.NetworkRuleSet_ARM">NetworkRuleSet_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_ARM">VaultProperties_ARM</a>)
</p>
<div>
<p>A set of rules governing the network accessibility of a vault.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>bypass</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet_Bypass">
NetworkRuleSet_Bypass
</a>
</em>
</td>
<td>
<p>Bypass: Tells what traffic can bypass network rules. This can be &lsquo;AzureServices&rsquo; or &lsquo;None&rsquo;.  If not specified the
default is &lsquo;AzureServices&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>defaultAction</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet_DefaultAction">
NetworkRuleSet_DefaultAction
</a>
</em>
</td>
<td>
<p>DefaultAction: The default action when no rule from ipRules and from virtualNetworkRules match. This is only used after
the bypass property has been evaluated.</p>
</td>
</tr>
<tr>
<td>
<code>ipRules</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.IPRule_ARM">
[]IPRule_ARM
</a>
</em>
</td>
<td>
<p>IpRules: The list of IP address rules.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkRules</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.VirtualNetworkRule_ARM">
[]VirtualNetworkRule_ARM
</a>
</em>
</td>
<td>
<p>VirtualNetworkRules: The list of virtual network rules.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.NetworkRuleSet_Bypass">NetworkRuleSet_Bypass
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet">NetworkRuleSet</a>, <a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet_ARM">NetworkRuleSet_ARM</a>)
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
<tbody><tr><td><p>&#34;AzureServices&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.NetworkRuleSet_Bypass_STATUS">NetworkRuleSet_Bypass_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet_STATUS">NetworkRuleSet_STATUS</a>, <a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet_STATUS_ARM">NetworkRuleSet_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;AzureServices&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.NetworkRuleSet_DefaultAction">NetworkRuleSet_DefaultAction
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet">NetworkRuleSet</a>, <a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet_ARM">NetworkRuleSet_ARM</a>)
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
<tbody><tr><td><p>&#34;Allow&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Deny&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.NetworkRuleSet_DefaultAction_STATUS">NetworkRuleSet_DefaultAction_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet_STATUS">NetworkRuleSet_STATUS</a>, <a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet_STATUS_ARM">NetworkRuleSet_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;Allow&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Deny&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.NetworkRuleSet_STATUS">NetworkRuleSet_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_STATUS">VaultProperties_STATUS</a>)
</p>
<div>
<p>A set of rules governing the network accessibility of a vault.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>bypass</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet_Bypass_STATUS">
NetworkRuleSet_Bypass_STATUS
</a>
</em>
</td>
<td>
<p>Bypass: Tells what traffic can bypass network rules. This can be &lsquo;AzureServices&rsquo; or &lsquo;None&rsquo;.  If not specified the
default is &lsquo;AzureServices&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>defaultAction</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet_DefaultAction_STATUS">
NetworkRuleSet_DefaultAction_STATUS
</a>
</em>
</td>
<td>
<p>DefaultAction: The default action when no rule from ipRules and from virtualNetworkRules match. This is only used after
the bypass property has been evaluated.</p>
</td>
</tr>
<tr>
<td>
<code>ipRules</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.IPRule_STATUS">
[]IPRule_STATUS
</a>
</em>
</td>
<td>
<p>IpRules: The list of IP address rules.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkRules</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.VirtualNetworkRule_STATUS">
[]VirtualNetworkRule_STATUS
</a>
</em>
</td>
<td>
<p>VirtualNetworkRules: The list of virtual network rules.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.NetworkRuleSet_STATUS_ARM">NetworkRuleSet_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_STATUS_ARM">VaultProperties_STATUS_ARM</a>)
</p>
<div>
<p>A set of rules governing the network accessibility of a vault.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>bypass</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet_Bypass_STATUS">
NetworkRuleSet_Bypass_STATUS
</a>
</em>
</td>
<td>
<p>Bypass: Tells what traffic can bypass network rules. This can be &lsquo;AzureServices&rsquo; or &lsquo;None&rsquo;.  If not specified the
default is &lsquo;AzureServices&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>defaultAction</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet_DefaultAction_STATUS">
NetworkRuleSet_DefaultAction_STATUS
</a>
</em>
</td>
<td>
<p>DefaultAction: The default action when no rule from ipRules and from virtualNetworkRules match. This is only used after
the bypass property has been evaluated.</p>
</td>
</tr>
<tr>
<td>
<code>ipRules</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.IPRule_STATUS_ARM">
[]IPRule_STATUS_ARM
</a>
</em>
</td>
<td>
<p>IpRules: The list of IP address rules.</p>
</td>
</tr>
<tr>
<td>
<code>virtualNetworkRules</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.VirtualNetworkRule_STATUS_ARM">
[]VirtualNetworkRule_STATUS_ARM
</a>
</em>
</td>
<td>
<p>VirtualNetworkRules: The list of virtual network rules.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.Permissions">Permissions
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.AccessPolicyEntry">AccessPolicyEntry</a>)
</p>
<div>
<p>Permissions the identity has for keys, secrets, certificates and storage.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>certificates</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Permissions_Certificates">
[]Permissions_Certificates
</a>
</em>
</td>
<td>
<p>Certificates: Permissions to certificates</p>
</td>
</tr>
<tr>
<td>
<code>keys</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Permissions_Keys">
[]Permissions_Keys
</a>
</em>
</td>
<td>
<p>Keys: Permissions to keys</p>
</td>
</tr>
<tr>
<td>
<code>secrets</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Permissions_Secrets">
[]Permissions_Secrets
</a>
</em>
</td>
<td>
<p>Secrets: Permissions to secrets</p>
</td>
</tr>
<tr>
<td>
<code>storage</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Permissions_Storage">
[]Permissions_Storage
</a>
</em>
</td>
<td>
<p>Storage: Permissions to storage accounts</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.Permissions_ARM">Permissions_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.AccessPolicyEntry_ARM">AccessPolicyEntry_ARM</a>)
</p>
<div>
<p>Permissions the identity has for keys, secrets, certificates and storage.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>certificates</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Permissions_Certificates">
[]Permissions_Certificates
</a>
</em>
</td>
<td>
<p>Certificates: Permissions to certificates</p>
</td>
</tr>
<tr>
<td>
<code>keys</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Permissions_Keys">
[]Permissions_Keys
</a>
</em>
</td>
<td>
<p>Keys: Permissions to keys</p>
</td>
</tr>
<tr>
<td>
<code>secrets</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Permissions_Secrets">
[]Permissions_Secrets
</a>
</em>
</td>
<td>
<p>Secrets: Permissions to secrets</p>
</td>
</tr>
<tr>
<td>
<code>storage</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Permissions_Storage">
[]Permissions_Storage
</a>
</em>
</td>
<td>
<p>Storage: Permissions to storage accounts</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.Permissions_Certificates">Permissions_Certificates
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.Permissions">Permissions</a>, <a href="#keyvault.azure.com/v1api20210401preview.Permissions_ARM">Permissions_ARM</a>)
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
<tbody><tr><td><p>&#34;backup&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;create&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;delete&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;deleteissuers&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;get&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;getissuers&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;import&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;list&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;listissuers&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;managecontacts&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;manageissuers&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;purge&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;recover&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;restore&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;setissuers&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;update&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.Permissions_Certificates_STATUS">Permissions_Certificates_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.Permissions_STATUS">Permissions_STATUS</a>, <a href="#keyvault.azure.com/v1api20210401preview.Permissions_STATUS_ARM">Permissions_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;backup&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;create&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;delete&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;deleteissuers&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;get&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;getissuers&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;import&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;list&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;listissuers&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;managecontacts&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;manageissuers&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;purge&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;recover&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;restore&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;setissuers&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;update&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.Permissions_Keys">Permissions_Keys
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.Permissions">Permissions</a>, <a href="#keyvault.azure.com/v1api20210401preview.Permissions_ARM">Permissions_ARM</a>)
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
<tbody><tr><td><p>&#34;backup&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;create&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;decrypt&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;delete&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;encrypt&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;get&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;import&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;list&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;purge&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;recover&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;release&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;restore&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;sign&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;unwrapKey&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;update&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;verify&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;wrapKey&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.Permissions_Keys_STATUS">Permissions_Keys_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.Permissions_STATUS">Permissions_STATUS</a>, <a href="#keyvault.azure.com/v1api20210401preview.Permissions_STATUS_ARM">Permissions_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;backup&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;create&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;decrypt&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;delete&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;encrypt&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;get&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;import&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;list&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;purge&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;recover&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;release&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;restore&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;sign&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;unwrapKey&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;update&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;verify&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;wrapKey&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.Permissions_STATUS">Permissions_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.AccessPolicyEntry_STATUS">AccessPolicyEntry_STATUS</a>)
</p>
<div>
<p>Permissions the identity has for keys, secrets, certificates and storage.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>certificates</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Permissions_Certificates_STATUS">
[]Permissions_Certificates_STATUS
</a>
</em>
</td>
<td>
<p>Certificates: Permissions to certificates</p>
</td>
</tr>
<tr>
<td>
<code>keys</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Permissions_Keys_STATUS">
[]Permissions_Keys_STATUS
</a>
</em>
</td>
<td>
<p>Keys: Permissions to keys</p>
</td>
</tr>
<tr>
<td>
<code>secrets</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Permissions_Secrets_STATUS">
[]Permissions_Secrets_STATUS
</a>
</em>
</td>
<td>
<p>Secrets: Permissions to secrets</p>
</td>
</tr>
<tr>
<td>
<code>storage</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Permissions_Storage_STATUS">
[]Permissions_Storage_STATUS
</a>
</em>
</td>
<td>
<p>Storage: Permissions to storage accounts</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.Permissions_STATUS_ARM">Permissions_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.AccessPolicyEntry_STATUS_ARM">AccessPolicyEntry_STATUS_ARM</a>)
</p>
<div>
<p>Permissions the identity has for keys, secrets, certificates and storage.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>certificates</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Permissions_Certificates_STATUS">
[]Permissions_Certificates_STATUS
</a>
</em>
</td>
<td>
<p>Certificates: Permissions to certificates</p>
</td>
</tr>
<tr>
<td>
<code>keys</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Permissions_Keys_STATUS">
[]Permissions_Keys_STATUS
</a>
</em>
</td>
<td>
<p>Keys: Permissions to keys</p>
</td>
</tr>
<tr>
<td>
<code>secrets</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Permissions_Secrets_STATUS">
[]Permissions_Secrets_STATUS
</a>
</em>
</td>
<td>
<p>Secrets: Permissions to secrets</p>
</td>
</tr>
<tr>
<td>
<code>storage</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Permissions_Storage_STATUS">
[]Permissions_Storage_STATUS
</a>
</em>
</td>
<td>
<p>Storage: Permissions to storage accounts</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.Permissions_Secrets">Permissions_Secrets
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.Permissions">Permissions</a>, <a href="#keyvault.azure.com/v1api20210401preview.Permissions_ARM">Permissions_ARM</a>)
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
<tbody><tr><td><p>&#34;backup&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;delete&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;get&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;list&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;purge&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;recover&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;restore&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;set&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.Permissions_Secrets_STATUS">Permissions_Secrets_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.Permissions_STATUS">Permissions_STATUS</a>, <a href="#keyvault.azure.com/v1api20210401preview.Permissions_STATUS_ARM">Permissions_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;backup&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;delete&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;get&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;list&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;purge&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;recover&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;restore&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;set&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.Permissions_Storage">Permissions_Storage
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.Permissions">Permissions</a>, <a href="#keyvault.azure.com/v1api20210401preview.Permissions_ARM">Permissions_ARM</a>)
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
<tbody><tr><td><p>&#34;backup&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;delete&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;deletesas&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;get&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;getsas&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;list&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;listsas&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;purge&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;recover&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;regeneratekey&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;restore&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;set&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;setsas&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;update&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.Permissions_Storage_STATUS">Permissions_Storage_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.Permissions_STATUS">Permissions_STATUS</a>, <a href="#keyvault.azure.com/v1api20210401preview.Permissions_STATUS_ARM">Permissions_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;backup&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;delete&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;deletesas&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;get&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;getsas&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;list&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;listsas&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;purge&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;recover&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;regeneratekey&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;restore&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;set&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;setsas&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;update&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.PrivateEndpointConnectionItem_STATUS">PrivateEndpointConnectionItem_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_STATUS">VaultProperties_STATUS</a>)
</p>
<div>
<p>Private endpoint connection item.</p>
</div>
<table>
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
<p>Etag: Modified whenever there is a change in the state of private endpoint connection.</p>
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
<p>Id: Id of private endpoint connection.</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpoint</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.PrivateEndpoint_STATUS">
PrivateEndpoint_STATUS
</a>
</em>
</td>
<td>
<p>PrivateEndpoint: Properties of the private endpoint object.</p>
</td>
</tr>
<tr>
<td>
<code>privateLinkServiceConnectionState</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.PrivateLinkServiceConnectionState_STATUS">
PrivateLinkServiceConnectionState_STATUS
</a>
</em>
</td>
<td>
<p>PrivateLinkServiceConnectionState: Approval state of the private link connection.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.PrivateEndpointConnectionProvisioningState_STATUS">
PrivateEndpointConnectionProvisioningState_STATUS
</a>
</em>
</td>
<td>
<p>ProvisioningState: Provisioning state of the private endpoint connection.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.PrivateEndpointConnectionItem_STATUS_ARM">PrivateEndpointConnectionItem_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_STATUS_ARM">VaultProperties_STATUS_ARM</a>)
</p>
<div>
<p>Private endpoint connection item.</p>
</div>
<table>
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
<p>Etag: Modified whenever there is a change in the state of private endpoint connection.</p>
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
<p>Id: Id of private endpoint connection.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.PrivateEndpointConnectionProperties_STATUS_ARM">
PrivateEndpointConnectionProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: Private endpoint connection properties.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.PrivateEndpointConnectionProperties_STATUS_ARM">PrivateEndpointConnectionProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.PrivateEndpointConnectionItem_STATUS_ARM">PrivateEndpointConnectionItem_STATUS_ARM</a>)
</p>
<div>
<p>Properties of the private endpoint connection resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>privateEndpoint</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.PrivateEndpoint_STATUS_ARM">
PrivateEndpoint_STATUS_ARM
</a>
</em>
</td>
<td>
<p>PrivateEndpoint: Properties of the private endpoint object.</p>
</td>
</tr>
<tr>
<td>
<code>privateLinkServiceConnectionState</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.PrivateLinkServiceConnectionState_STATUS_ARM">
PrivateLinkServiceConnectionState_STATUS_ARM
</a>
</em>
</td>
<td>
<p>PrivateLinkServiceConnectionState: Approval state of the private link connection.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.PrivateEndpointConnectionProvisioningState_STATUS">
PrivateEndpointConnectionProvisioningState_STATUS
</a>
</em>
</td>
<td>
<p>ProvisioningState: Provisioning state of the private endpoint connection.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.PrivateEndpointConnectionProvisioningState_STATUS">PrivateEndpointConnectionProvisioningState_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.PrivateEndpointConnectionItem_STATUS">PrivateEndpointConnectionItem_STATUS</a>, <a href="#keyvault.azure.com/v1api20210401preview.PrivateEndpointConnectionProperties_STATUS_ARM">PrivateEndpointConnectionProperties_STATUS_ARM</a>)
</p>
<div>
<p>The current provisioning state.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Creating&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Deleting&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Disconnected&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Failed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Succeeded&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Updating&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.PrivateEndpointServiceConnectionStatus_STATUS">PrivateEndpointServiceConnectionStatus_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.PrivateLinkServiceConnectionState_STATUS">PrivateLinkServiceConnectionState_STATUS</a>, <a href="#keyvault.azure.com/v1api20210401preview.PrivateLinkServiceConnectionState_STATUS_ARM">PrivateLinkServiceConnectionState_STATUS_ARM</a>)
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
</tr><tr><td><p>&#34;Disconnected&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Pending&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Rejected&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.PrivateEndpoint_STATUS">PrivateEndpoint_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.PrivateEndpointConnectionItem_STATUS">PrivateEndpointConnectionItem_STATUS</a>)
</p>
<div>
<p>Private endpoint object properties.</p>
</div>
<table>
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
<p>Id: Full identifier of the private endpoint resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.PrivateEndpoint_STATUS_ARM">PrivateEndpoint_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.PrivateEndpointConnectionProperties_STATUS_ARM">PrivateEndpointConnectionProperties_STATUS_ARM</a>)
</p>
<div>
<p>Private endpoint object properties.</p>
</div>
<table>
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
<p>Id: Full identifier of the private endpoint resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.PrivateLinkServiceConnectionState_ActionsRequired_STATUS">PrivateLinkServiceConnectionState_ActionsRequired_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.PrivateLinkServiceConnectionState_STATUS">PrivateLinkServiceConnectionState_STATUS</a>, <a href="#keyvault.azure.com/v1api20210401preview.PrivateLinkServiceConnectionState_STATUS_ARM">PrivateLinkServiceConnectionState_STATUS_ARM</a>)
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
</tr></tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.PrivateLinkServiceConnectionState_STATUS">PrivateLinkServiceConnectionState_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.PrivateEndpointConnectionItem_STATUS">PrivateEndpointConnectionItem_STATUS</a>)
</p>
<div>
<p>An object that represents the approval state of the private link connection.</p>
</div>
<table>
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
<a href="#keyvault.azure.com/v1api20210401preview.PrivateLinkServiceConnectionState_ActionsRequired_STATUS">
PrivateLinkServiceConnectionState_ActionsRequired_STATUS
</a>
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
<p>Description: The reason for approval or rejection.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.PrivateEndpointServiceConnectionStatus_STATUS">
PrivateEndpointServiceConnectionStatus_STATUS
</a>
</em>
</td>
<td>
<p>Status: Indicates whether the connection has been approved, rejected or removed by the key vault owner.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.PrivateLinkServiceConnectionState_STATUS_ARM">PrivateLinkServiceConnectionState_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.PrivateEndpointConnectionProperties_STATUS_ARM">PrivateEndpointConnectionProperties_STATUS_ARM</a>)
</p>
<div>
<p>An object that represents the approval state of the private link connection.</p>
</div>
<table>
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
<a href="#keyvault.azure.com/v1api20210401preview.PrivateLinkServiceConnectionState_ActionsRequired_STATUS">
PrivateLinkServiceConnectionState_ActionsRequired_STATUS
</a>
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
<p>Description: The reason for approval or rejection.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.PrivateEndpointServiceConnectionStatus_STATUS">
PrivateEndpointServiceConnectionStatus_STATUS
</a>
</em>
</td>
<td>
<p>Status: Indicates whether the connection has been approved, rejected or removed by the key vault owner.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.Sku">Sku
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.VaultProperties">VaultProperties</a>)
</p>
<div>
<p>SKU details</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>family</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Sku_Family">
Sku_Family
</a>
</em>
</td>
<td>
<p>Family: SKU family name</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Sku_Name">
Sku_Name
</a>
</em>
</td>
<td>
<p>Name: SKU name to specify whether the key vault is a standard vault or a premium vault.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.Sku_ARM">Sku_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_ARM">VaultProperties_ARM</a>)
</p>
<div>
<p>SKU details</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>family</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Sku_Family">
Sku_Family
</a>
</em>
</td>
<td>
<p>Family: SKU family name</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Sku_Name">
Sku_Name
</a>
</em>
</td>
<td>
<p>Name: SKU name to specify whether the key vault is a standard vault or a premium vault.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.Sku_Family">Sku_Family
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.Sku">Sku</a>, <a href="#keyvault.azure.com/v1api20210401preview.Sku_ARM">Sku_ARM</a>)
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
<tbody><tr><td><p>&#34;A&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.Sku_Family_STATUS">Sku_Family_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.Sku_STATUS">Sku_STATUS</a>, <a href="#keyvault.azure.com/v1api20210401preview.Sku_STATUS_ARM">Sku_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;A&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.Sku_Name">Sku_Name
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.Sku">Sku</a>, <a href="#keyvault.azure.com/v1api20210401preview.Sku_ARM">Sku_ARM</a>)
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
<tbody><tr><td><p>&#34;premium&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.Sku_Name_STATUS">Sku_Name_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.Sku_STATUS">Sku_STATUS</a>, <a href="#keyvault.azure.com/v1api20210401preview.Sku_STATUS_ARM">Sku_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;premium&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.Sku_STATUS">Sku_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_STATUS">VaultProperties_STATUS</a>)
</p>
<div>
<p>SKU details</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>family</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Sku_Family_STATUS">
Sku_Family_STATUS
</a>
</em>
</td>
<td>
<p>Family: SKU family name</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Sku_Name_STATUS">
Sku_Name_STATUS
</a>
</em>
</td>
<td>
<p>Name: SKU name to specify whether the key vault is a standard vault or a premium vault.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.Sku_STATUS_ARM">Sku_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_STATUS_ARM">VaultProperties_STATUS_ARM</a>)
</p>
<div>
<p>SKU details</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>family</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Sku_Family_STATUS">
Sku_Family_STATUS
</a>
</em>
</td>
<td>
<p>Family: SKU family name</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Sku_Name_STATUS">
Sku_Name_STATUS
</a>
</em>
</td>
<td>
<p>Name: SKU name to specify whether the key vault is a standard vault or a premium vault.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.SystemData_STATUS">SystemData_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.Vault_STATUS">Vault_STATUS</a>)
</p>
<div>
<p>Metadata pertaining to creation and last modification of the key vault resource.</p>
</div>
<table>
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
<p>CreatedAt: The timestamp of the key vault resource creation (UTC).</p>
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
<p>CreatedBy: The identity that created the key vault resource.</p>
</td>
</tr>
<tr>
<td>
<code>createdByType</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.IdentityType_STATUS">
IdentityType_STATUS
</a>
</em>
</td>
<td>
<p>CreatedByType: The type of identity that created the key vault resource.</p>
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
<p>LastModifiedAt: The timestamp of the key vault resource last modification (UTC).</p>
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
<p>LastModifiedBy: The identity that last modified the key vault resource.</p>
</td>
</tr>
<tr>
<td>
<code>lastModifiedByType</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.IdentityType_STATUS">
IdentityType_STATUS
</a>
</em>
</td>
<td>
<p>LastModifiedByType: The type of identity that last modified the key vault resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.SystemData_STATUS_ARM">SystemData_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.Vault_STATUS_ARM">Vault_STATUS_ARM</a>)
</p>
<div>
<p>Metadata pertaining to creation and last modification of the key vault resource.</p>
</div>
<table>
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
<p>CreatedAt: The timestamp of the key vault resource creation (UTC).</p>
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
<p>CreatedBy: The identity that created the key vault resource.</p>
</td>
</tr>
<tr>
<td>
<code>createdByType</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.IdentityType_STATUS">
IdentityType_STATUS
</a>
</em>
</td>
<td>
<p>CreatedByType: The type of identity that created the key vault resource.</p>
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
<p>LastModifiedAt: The timestamp of the key vault resource last modification (UTC).</p>
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
<p>LastModifiedBy: The identity that last modified the key vault resource.</p>
</td>
</tr>
<tr>
<td>
<code>lastModifiedByType</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.IdentityType_STATUS">
IdentityType_STATUS
</a>
</em>
</td>
<td>
<p>LastModifiedByType: The type of identity that last modified the key vault resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.Vault">Vault
</h3>
<div>
<p>Generator information:
- Generated from: /keyvault/resource-manager/Microsoft.KeyVault/preview/2021-04-01-preview/keyvault.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.KeyVault/&#x200b;vaults/&#x200b;{vaultName}</&#x200b;p>
</div>
<table>
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
<a href="#keyvault.azure.com/v1api20210401preview.Vault_Spec">
Vault_Spec
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
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: The supported Azure location where the key vault should be created.</p>
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
<code>properties</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.VaultProperties">
VaultProperties
</a>
</em>
</td>
<td>
<p>Properties: Properties of the vault</p>
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
<p>Tags: The tags that will be assigned to the key vault.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Vault_STATUS">
Vault_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.VaultProperties">VaultProperties
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.Vault_Spec">Vault_Spec</a>)
</p>
<div>
<p>Properties of the vault</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>accessPolicies</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.AccessPolicyEntry">
[]AccessPolicyEntry
</a>
</em>
</td>
<td>
<p>AccessPolicies: An array of 0 to 1024 identities that have access to the key vault. All identities in the array must use
the same tenant ID as the key vault&rsquo;s tenant ID. When <code>createMode</code> is set to <code>recover</code>, access policies are not
required. Otherwise, access policies are required.</p>
</td>
</tr>
<tr>
<td>
<code>createMode</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_CreateMode">
VaultProperties_CreateMode
</a>
</em>
</td>
<td>
<p>CreateMode: The vault&rsquo;s create mode to indicate whether the vault need to be recovered or not.</p>
</td>
</tr>
<tr>
<td>
<code>enablePurgeProtection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePurgeProtection: Property specifying whether protection against purge is enabled for this vault. Setting this
property to true activates protection against purge for this vault and its content - only the Key Vault service may
initiate a hard, irrecoverable deletion. The setting is effective only if soft delete is also enabled. Enabling this
functionality is irreversible - that is, the property does not accept false as its value.</p>
</td>
</tr>
<tr>
<td>
<code>enableRbacAuthorization</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableRbacAuthorization: Property that controls how data actions are authorized. When true, the key vault will use Role
Based Access Control (RBAC) for authorization of data actions, and the access policies specified in vault properties
will be  ignored. When false, the key vault will use the access policies specified in vault properties, and any policy
stored on Azure Resource Manager will be ignored. If null or not specified, the vault is created with the default value
of false. Note that management actions are always authorized with RBAC.</p>
</td>
</tr>
<tr>
<td>
<code>enableSoftDelete</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableSoftDelete: Property to specify whether the &lsquo;soft delete&rsquo; functionality is enabled for this key vault. If it&rsquo;s not
set to any value(true or false) when creating new key vault, it will be set to true by default. Once set to true, it
cannot be reverted to false.</p>
</td>
</tr>
<tr>
<td>
<code>enabledForDeployment</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnabledForDeployment: Property to specify whether Azure Virtual Machines are permitted to retrieve certificates stored
as secrets from the key vault.</p>
</td>
</tr>
<tr>
<td>
<code>enabledForDiskEncryption</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnabledForDiskEncryption: Property to specify whether Azure Disk Encryption is permitted to retrieve secrets from the
vault and unwrap keys.</p>
</td>
</tr>
<tr>
<td>
<code>enabledForTemplateDeployment</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnabledForTemplateDeployment: Property to specify whether Azure Resource Manager is permitted to retrieve secrets from
the key vault.</p>
</td>
</tr>
<tr>
<td>
<code>networkAcls</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet">
NetworkRuleSet
</a>
</em>
</td>
<td>
<p>NetworkAcls: Rules governing the accessibility of the key vault from specific network locations.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_ProvisioningState">
VaultProperties_ProvisioningState
</a>
</em>
</td>
<td>
<p>ProvisioningState: Provisioning state of the vault.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Sku">
Sku
</a>
</em>
</td>
<td>
<p>Sku: SKU details</p>
</td>
</tr>
<tr>
<td>
<code>softDeleteRetentionInDays</code><br/>
<em>
int
</em>
</td>
<td>
<p>SoftDeleteRetentionInDays: softDelete data retention days. It accepts &gt;=7 and &lt;=90.</p>
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
<p>TenantId: The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.</p>
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
<p>TenantIdFromConfig: The Azure Active Directory tenant ID that should be used for authenticating requests to the key
vault.</p>
</td>
</tr>
<tr>
<td>
<code>vaultUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>VaultUri: The URI of the vault for performing operations on keys and secrets.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.VaultProperties_ARM">VaultProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.Vault_Spec_ARM">Vault_Spec_ARM</a>)
</p>
<div>
<p>Properties of the vault</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>accessPolicies</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.AccessPolicyEntry_ARM">
[]AccessPolicyEntry_ARM
</a>
</em>
</td>
<td>
<p>AccessPolicies: An array of 0 to 1024 identities that have access to the key vault. All identities in the array must use
the same tenant ID as the key vault&rsquo;s tenant ID. When <code>createMode</code> is set to <code>recover</code>, access policies are not
required. Otherwise, access policies are required.</p>
</td>
</tr>
<tr>
<td>
<code>createMode</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_CreateMode">
VaultProperties_CreateMode
</a>
</em>
</td>
<td>
<p>CreateMode: The vault&rsquo;s create mode to indicate whether the vault need to be recovered or not.</p>
</td>
</tr>
<tr>
<td>
<code>enablePurgeProtection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePurgeProtection: Property specifying whether protection against purge is enabled for this vault. Setting this
property to true activates protection against purge for this vault and its content - only the Key Vault service may
initiate a hard, irrecoverable deletion. The setting is effective only if soft delete is also enabled. Enabling this
functionality is irreversible - that is, the property does not accept false as its value.</p>
</td>
</tr>
<tr>
<td>
<code>enableRbacAuthorization</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableRbacAuthorization: Property that controls how data actions are authorized. When true, the key vault will use Role
Based Access Control (RBAC) for authorization of data actions, and the access policies specified in vault properties
will be  ignored. When false, the key vault will use the access policies specified in vault properties, and any policy
stored on Azure Resource Manager will be ignored. If null or not specified, the vault is created with the default value
of false. Note that management actions are always authorized with RBAC.</p>
</td>
</tr>
<tr>
<td>
<code>enableSoftDelete</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableSoftDelete: Property to specify whether the &lsquo;soft delete&rsquo; functionality is enabled for this key vault. If it&rsquo;s not
set to any value(true or false) when creating new key vault, it will be set to true by default. Once set to true, it
cannot be reverted to false.</p>
</td>
</tr>
<tr>
<td>
<code>enabledForDeployment</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnabledForDeployment: Property to specify whether Azure Virtual Machines are permitted to retrieve certificates stored
as secrets from the key vault.</p>
</td>
</tr>
<tr>
<td>
<code>enabledForDiskEncryption</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnabledForDiskEncryption: Property to specify whether Azure Disk Encryption is permitted to retrieve secrets from the
vault and unwrap keys.</p>
</td>
</tr>
<tr>
<td>
<code>enabledForTemplateDeployment</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnabledForTemplateDeployment: Property to specify whether Azure Resource Manager is permitted to retrieve secrets from
the key vault.</p>
</td>
</tr>
<tr>
<td>
<code>networkAcls</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet_ARM">
NetworkRuleSet_ARM
</a>
</em>
</td>
<td>
<p>NetworkAcls: Rules governing the accessibility of the key vault from specific network locations.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_ProvisioningState">
VaultProperties_ProvisioningState
</a>
</em>
</td>
<td>
<p>ProvisioningState: Provisioning state of the vault.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Sku_ARM">
Sku_ARM
</a>
</em>
</td>
<td>
<p>Sku: SKU details</p>
</td>
</tr>
<tr>
<td>
<code>softDeleteRetentionInDays</code><br/>
<em>
int
</em>
</td>
<td>
<p>SoftDeleteRetentionInDays: softDelete data retention days. It accepts &gt;=7 and &lt;=90.</p>
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
<p>TenantId: The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.</p>
</td>
</tr>
<tr>
<td>
<code>vaultUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>VaultUri: The URI of the vault for performing operations on keys and secrets.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.VaultProperties_CreateMode">VaultProperties_CreateMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.VaultProperties">VaultProperties</a>, <a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_ARM">VaultProperties_ARM</a>)
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
<tbody><tr><td><p>&#34;default&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;recover&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.VaultProperties_CreateMode_STATUS">VaultProperties_CreateMode_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_STATUS">VaultProperties_STATUS</a>, <a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_STATUS_ARM">VaultProperties_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;default&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;recover&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.VaultProperties_ProvisioningState">VaultProperties_ProvisioningState
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.VaultProperties">VaultProperties</a>, <a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_ARM">VaultProperties_ARM</a>)
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
<tbody><tr><td><p>&#34;RegisteringDns&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Succeeded&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.VaultProperties_ProvisioningState_STATUS">VaultProperties_ProvisioningState_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_STATUS">VaultProperties_STATUS</a>, <a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_STATUS_ARM">VaultProperties_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;RegisteringDns&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Succeeded&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.VaultProperties_STATUS">VaultProperties_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.Vault_STATUS">Vault_STATUS</a>)
</p>
<div>
<p>Properties of the vault</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>accessPolicies</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.AccessPolicyEntry_STATUS">
[]AccessPolicyEntry_STATUS
</a>
</em>
</td>
<td>
<p>AccessPolicies: An array of 0 to 1024 identities that have access to the key vault. All identities in the array must use
the same tenant ID as the key vault&rsquo;s tenant ID. When <code>createMode</code> is set to <code>recover</code>, access policies are not
required. Otherwise, access policies are required.</p>
</td>
</tr>
<tr>
<td>
<code>createMode</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_CreateMode_STATUS">
VaultProperties_CreateMode_STATUS
</a>
</em>
</td>
<td>
<p>CreateMode: The vault&rsquo;s create mode to indicate whether the vault need to be recovered or not.</p>
</td>
</tr>
<tr>
<td>
<code>enablePurgeProtection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePurgeProtection: Property specifying whether protection against purge is enabled for this vault. Setting this
property to true activates protection against purge for this vault and its content - only the Key Vault service may
initiate a hard, irrecoverable deletion. The setting is effective only if soft delete is also enabled. Enabling this
functionality is irreversible - that is, the property does not accept false as its value.</p>
</td>
</tr>
<tr>
<td>
<code>enableRbacAuthorization</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableRbacAuthorization: Property that controls how data actions are authorized. When true, the key vault will use Role
Based Access Control (RBAC) for authorization of data actions, and the access policies specified in vault properties
will be  ignored. When false, the key vault will use the access policies specified in vault properties, and any policy
stored on Azure Resource Manager will be ignored. If null or not specified, the vault is created with the default value
of false. Note that management actions are always authorized with RBAC.</p>
</td>
</tr>
<tr>
<td>
<code>enableSoftDelete</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableSoftDelete: Property to specify whether the &lsquo;soft delete&rsquo; functionality is enabled for this key vault. If it&rsquo;s not
set to any value(true or false) when creating new key vault, it will be set to true by default. Once set to true, it
cannot be reverted to false.</p>
</td>
</tr>
<tr>
<td>
<code>enabledForDeployment</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnabledForDeployment: Property to specify whether Azure Virtual Machines are permitted to retrieve certificates stored
as secrets from the key vault.</p>
</td>
</tr>
<tr>
<td>
<code>enabledForDiskEncryption</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnabledForDiskEncryption: Property to specify whether Azure Disk Encryption is permitted to retrieve secrets from the
vault and unwrap keys.</p>
</td>
</tr>
<tr>
<td>
<code>enabledForTemplateDeployment</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnabledForTemplateDeployment: Property to specify whether Azure Resource Manager is permitted to retrieve secrets from
the key vault.</p>
</td>
</tr>
<tr>
<td>
<code>hsmPoolResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>HsmPoolResourceId: The resource id of HSM Pool.</p>
</td>
</tr>
<tr>
<td>
<code>networkAcls</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet_STATUS">
NetworkRuleSet_STATUS
</a>
</em>
</td>
<td>
<p>NetworkAcls: Rules governing the accessibility of the key vault from specific network locations.</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpointConnections</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.PrivateEndpointConnectionItem_STATUS">
[]PrivateEndpointConnectionItem_STATUS
</a>
</em>
</td>
<td>
<p>PrivateEndpointConnections: List of private endpoint connections associated with the key vault.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_ProvisioningState_STATUS">
VaultProperties_ProvisioningState_STATUS
</a>
</em>
</td>
<td>
<p>ProvisioningState: Provisioning state of the vault.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Sku_STATUS">
Sku_STATUS
</a>
</em>
</td>
<td>
<p>Sku: SKU details</p>
</td>
</tr>
<tr>
<td>
<code>softDeleteRetentionInDays</code><br/>
<em>
int
</em>
</td>
<td>
<p>SoftDeleteRetentionInDays: softDelete data retention days. It accepts &gt;=7 and &lt;=90.</p>
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
<p>TenantId: The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.</p>
</td>
</tr>
<tr>
<td>
<code>vaultUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>VaultUri: The URI of the vault for performing operations on keys and secrets.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.VaultProperties_STATUS_ARM">VaultProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.Vault_STATUS_ARM">Vault_STATUS_ARM</a>)
</p>
<div>
<p>Properties of the vault</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>accessPolicies</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.AccessPolicyEntry_STATUS_ARM">
[]AccessPolicyEntry_STATUS_ARM
</a>
</em>
</td>
<td>
<p>AccessPolicies: An array of 0 to 1024 identities that have access to the key vault. All identities in the array must use
the same tenant ID as the key vault&rsquo;s tenant ID. When <code>createMode</code> is set to <code>recover</code>, access policies are not
required. Otherwise, access policies are required.</p>
</td>
</tr>
<tr>
<td>
<code>createMode</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_CreateMode_STATUS">
VaultProperties_CreateMode_STATUS
</a>
</em>
</td>
<td>
<p>CreateMode: The vault&rsquo;s create mode to indicate whether the vault need to be recovered or not.</p>
</td>
</tr>
<tr>
<td>
<code>enablePurgeProtection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePurgeProtection: Property specifying whether protection against purge is enabled for this vault. Setting this
property to true activates protection against purge for this vault and its content - only the Key Vault service may
initiate a hard, irrecoverable deletion. The setting is effective only if soft delete is also enabled. Enabling this
functionality is irreversible - that is, the property does not accept false as its value.</p>
</td>
</tr>
<tr>
<td>
<code>enableRbacAuthorization</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableRbacAuthorization: Property that controls how data actions are authorized. When true, the key vault will use Role
Based Access Control (RBAC) for authorization of data actions, and the access policies specified in vault properties
will be  ignored. When false, the key vault will use the access policies specified in vault properties, and any policy
stored on Azure Resource Manager will be ignored. If null or not specified, the vault is created with the default value
of false. Note that management actions are always authorized with RBAC.</p>
</td>
</tr>
<tr>
<td>
<code>enableSoftDelete</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableSoftDelete: Property to specify whether the &lsquo;soft delete&rsquo; functionality is enabled for this key vault. If it&rsquo;s not
set to any value(true or false) when creating new key vault, it will be set to true by default. Once set to true, it
cannot be reverted to false.</p>
</td>
</tr>
<tr>
<td>
<code>enabledForDeployment</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnabledForDeployment: Property to specify whether Azure Virtual Machines are permitted to retrieve certificates stored
as secrets from the key vault.</p>
</td>
</tr>
<tr>
<td>
<code>enabledForDiskEncryption</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnabledForDiskEncryption: Property to specify whether Azure Disk Encryption is permitted to retrieve secrets from the
vault and unwrap keys.</p>
</td>
</tr>
<tr>
<td>
<code>enabledForTemplateDeployment</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnabledForTemplateDeployment: Property to specify whether Azure Resource Manager is permitted to retrieve secrets from
the key vault.</p>
</td>
</tr>
<tr>
<td>
<code>hsmPoolResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>HsmPoolResourceId: The resource id of HSM Pool.</p>
</td>
</tr>
<tr>
<td>
<code>networkAcls</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet_STATUS_ARM">
NetworkRuleSet_STATUS_ARM
</a>
</em>
</td>
<td>
<p>NetworkAcls: Rules governing the accessibility of the key vault from specific network locations.</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpointConnections</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.PrivateEndpointConnectionItem_STATUS_ARM">
[]PrivateEndpointConnectionItem_STATUS_ARM
</a>
</em>
</td>
<td>
<p>PrivateEndpointConnections: List of private endpoint connections associated with the key vault.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_ProvisioningState_STATUS">
VaultProperties_ProvisioningState_STATUS
</a>
</em>
</td>
<td>
<p>ProvisioningState: Provisioning state of the vault.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.Sku_STATUS_ARM">
Sku_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Sku: SKU details</p>
</td>
</tr>
<tr>
<td>
<code>softDeleteRetentionInDays</code><br/>
<em>
int
</em>
</td>
<td>
<p>SoftDeleteRetentionInDays: softDelete data retention days. It accepts &gt;=7 and &lt;=90.</p>
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
<p>TenantId: The Azure Active Directory tenant ID that should be used for authenticating requests to the key vault.</p>
</td>
</tr>
<tr>
<td>
<code>vaultUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>VaultUri: The URI of the vault for performing operations on keys and secrets.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.Vault_STATUS">Vault_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.Vault">Vault</a>)
</p>
<div>
<p>Resource information with extended details.</p>
</div>
<table>
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
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified identifier of the key vault resource.</p>
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
<p>Location: Azure location of the key vault resource.</p>
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
<p>Name: Name of the key vault resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_STATUS">
VaultProperties_STATUS
</a>
</em>
</td>
<td>
<p>Properties: Properties of the vault</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.SystemData_STATUS">
SystemData_STATUS
</a>
</em>
</td>
<td>
<p>SystemData: System metadata for the key vault.</p>
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
<p>Tags: Tags assigned to the key vault resource.</p>
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
<p>Type: Resource type of the key vault resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.Vault_STATUS_ARM">Vault_STATUS_ARM
</h3>
<div>
<p>Resource information with extended details.</p>
</div>
<table>
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
<p>Id: Fully qualified identifier of the key vault resource.</p>
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
<p>Location: Azure location of the key vault resource.</p>
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
<p>Name: Name of the key vault resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_STATUS_ARM">
VaultProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the vault</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.SystemData_STATUS_ARM">
SystemData_STATUS_ARM
</a>
</em>
</td>
<td>
<p>SystemData: System metadata for the key vault.</p>
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
<p>Tags: Tags assigned to the key vault resource.</p>
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
<p>Type: Resource type of the key vault resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.Vault_Spec">Vault_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.Vault">Vault</a>)
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
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: The supported Azure location where the key vault should be created.</p>
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
<code>properties</code><br/>
<em>
<a href="#keyvault.azure.com/v1api20210401preview.VaultProperties">
VaultProperties
</a>
</em>
</td>
<td>
<p>Properties: Properties of the vault</p>
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
<p>Tags: The tags that will be assigned to the key vault.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.Vault_Spec_ARM">Vault_Spec_ARM
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
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: The supported Azure location where the key vault should be created.</p>
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
<a href="#keyvault.azure.com/v1api20210401preview.VaultProperties_ARM">
VaultProperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the vault</p>
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
<p>Tags: The tags that will be assigned to the key vault.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.VirtualNetworkRule">VirtualNetworkRule
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet">NetworkRuleSet</a>)
</p>
<div>
<p>A rule governing the accessibility of a vault from a specific virtual network.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ignoreMissingVnetServiceEndpoint</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IgnoreMissingVnetServiceEndpoint: Property to specify whether NRP will ignore the check if parent subnet has
serviceEndpoints configured.</p>
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
<p>Reference: Full resource id of a vnet subnet, such as
&lsquo;/&#x200b;subscriptions/&#x200b;subid/&#x200b;resourceGroups/&#x200b;rg1/&#x200b;providers/&#x200b;Microsoft.Network/&#x200b;virtualNetworks/&#x200b;test-vnet/&#x200b;subnets/&#x200b;subnet1&rsquo;.</&#x200b;p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.VirtualNetworkRule_ARM">VirtualNetworkRule_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet_ARM">NetworkRuleSet_ARM</a>)
</p>
<div>
<p>A rule governing the accessibility of a vault from a specific virtual network.</p>
</div>
<table>
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
<tr>
<td>
<code>ignoreMissingVnetServiceEndpoint</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IgnoreMissingVnetServiceEndpoint: Property to specify whether NRP will ignore the check if parent subnet has
serviceEndpoints configured.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.VirtualNetworkRule_STATUS">VirtualNetworkRule_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet_STATUS">NetworkRuleSet_STATUS</a>)
</p>
<div>
<p>A rule governing the accessibility of a vault from a specific virtual network.</p>
</div>
<table>
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
<p>Id: Full resource id of a vnet subnet, such as
&lsquo;/&#x200b;subscriptions/&#x200b;subid/&#x200b;resourceGroups/&#x200b;rg1/&#x200b;providers/&#x200b;Microsoft.Network/&#x200b;virtualNetworks/&#x200b;test-vnet/&#x200b;subnets/&#x200b;subnet1&rsquo;.</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>ignoreMissingVnetServiceEndpoint</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IgnoreMissingVnetServiceEndpoint: Property to specify whether NRP will ignore the check if parent subnet has
serviceEndpoints configured.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1api20210401preview.VirtualNetworkRule_STATUS_ARM">VirtualNetworkRule_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1api20210401preview.NetworkRuleSet_STATUS_ARM">NetworkRuleSet_STATUS_ARM</a>)
</p>
<div>
<p>A rule governing the accessibility of a vault from a specific virtual network.</p>
</div>
<table>
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
<p>Id: Full resource id of a vnet subnet, such as
&lsquo;/&#x200b;subscriptions/&#x200b;subid/&#x200b;resourceGroups/&#x200b;rg1/&#x200b;providers/&#x200b;Microsoft.Network/&#x200b;virtualNetworks/&#x200b;test-vnet/&#x200b;subnets/&#x200b;subnet1&rsquo;.</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>ignoreMissingVnetServiceEndpoint</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IgnoreMissingVnetServiceEndpoint: Property to specify whether NRP will ignore the check if parent subnet has
serviceEndpoints configured.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
