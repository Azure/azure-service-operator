---
title: keyvault.azure.com/v1beta20210401preview
---
<h2 id="keyvault.azure.com/v1beta20210401preview">keyvault.azure.com/v1beta20210401preview</h2>
<div>
<p>Package v1beta20210401preview contains API Schema definitions for the keyvault v1beta20210401preview API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="keyvault.azure.com/v1beta20210401preview.APIVersion">APIVersion
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
<h3 id="keyvault.azure.com/v1beta20210401preview.AccessPolicyEntry">AccessPolicyEntry
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.VaultProperties">VaultProperties</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/AccessPolicyEntry">https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/AccessPolicyEntry</a></p>
</div>
<table>
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
<a href="#keyvault.azure.com/v1beta20210401preview.Permissions">
Permissions
</a>
</em>
</td>
<td>
<p>Permissions: Permissions the identity has for keys, secrets, certificates and storage.</p>
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
<h3 id="keyvault.azure.com/v1beta20210401preview.AccessPolicyEntryARM">AccessPolicyEntryARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.VaultPropertiesARM">VaultPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/AccessPolicyEntry">https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/AccessPolicyEntry</a></p>
</div>
<table>
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
<a href="#keyvault.azure.com/v1beta20210401preview.PermissionsARM">
PermissionsARM
</a>
</em>
</td>
<td>
<p>Permissions: Permissions the identity has for keys, secrets, certificates and storage.</p>
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
<h3 id="keyvault.azure.com/v1beta20210401preview.AccessPolicyEntry_Status">AccessPolicyEntry_Status
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.VaultProperties_Status">VaultProperties_Status</a>)
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
<a href="#keyvault.azure.com/v1beta20210401preview.Permissions_Status">
Permissions_Status
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
<h3 id="keyvault.azure.com/v1beta20210401preview.AccessPolicyEntry_StatusARM">AccessPolicyEntry_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.VaultProperties_StatusARM">VaultProperties_StatusARM</a>)
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
<a href="#keyvault.azure.com/v1beta20210401preview.Permissions_StatusARM">
Permissions_StatusARM
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
<h3 id="keyvault.azure.com/v1beta20210401preview.IPRule">IPRule
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSet">NetworkRuleSet</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/IPRule">https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/IPRule</a></p>
</div>
<table>
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
<h3 id="keyvault.azure.com/v1beta20210401preview.IPRuleARM">IPRuleARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSetARM">NetworkRuleSetARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/IPRule">https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/IPRule</a></p>
</div>
<table>
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
<h3 id="keyvault.azure.com/v1beta20210401preview.IPRule_Status">IPRule_Status
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSet_Status">NetworkRuleSet_Status</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.IPRule_StatusARM">IPRule_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSet_StatusARM">NetworkRuleSet_StatusARM</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.IdentityType_Status">IdentityType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.SystemData_Status">SystemData_Status</a>, <a href="#keyvault.azure.com/v1beta20210401preview.SystemData_StatusARM">SystemData_StatusARM</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.NetworkRuleSet">NetworkRuleSet
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.VaultProperties">VaultProperties</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/NetworkRuleSet">https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/NetworkRuleSet</a></p>
</div>
<table>
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
<a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSetBypass">
NetworkRuleSetBypass
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
<a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSetDefaultAction">
NetworkRuleSetDefaultAction
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
<a href="#keyvault.azure.com/v1beta20210401preview.IPRule">
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
<a href="#keyvault.azure.com/v1beta20210401preview.VirtualNetworkRule">
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
<h3 id="keyvault.azure.com/v1beta20210401preview.NetworkRuleSetARM">NetworkRuleSetARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.VaultPropertiesARM">VaultPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/NetworkRuleSet">https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/NetworkRuleSet</a></p>
</div>
<table>
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
<a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSetBypass">
NetworkRuleSetBypass
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
<a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSetDefaultAction">
NetworkRuleSetDefaultAction
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
<a href="#keyvault.azure.com/v1beta20210401preview.IPRuleARM">
[]IPRuleARM
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
<a href="#keyvault.azure.com/v1beta20210401preview.VirtualNetworkRuleARM">
[]VirtualNetworkRuleARM
</a>
</em>
</td>
<td>
<p>VirtualNetworkRules: The list of virtual network rules.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1beta20210401preview.NetworkRuleSetBypass">NetworkRuleSetBypass
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSet">NetworkRuleSet</a>, <a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSetARM">NetworkRuleSetARM</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.NetworkRuleSetDefaultAction">NetworkRuleSetDefaultAction
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSet">NetworkRuleSet</a>, <a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSetARM">NetworkRuleSetARM</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.NetworkRuleSetStatusBypass">NetworkRuleSetStatusBypass
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSet_Status">NetworkRuleSet_Status</a>, <a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSet_StatusARM">NetworkRuleSet_StatusARM</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.NetworkRuleSetStatusDefaultAction">NetworkRuleSetStatusDefaultAction
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSet_Status">NetworkRuleSet_Status</a>, <a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSet_StatusARM">NetworkRuleSet_StatusARM</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.NetworkRuleSet_Status">NetworkRuleSet_Status
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.VaultProperties_Status">VaultProperties_Status</a>)
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
<code>bypass</code><br/>
<em>
<a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSetStatusBypass">
NetworkRuleSetStatusBypass
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
<a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSetStatusDefaultAction">
NetworkRuleSetStatusDefaultAction
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
<a href="#keyvault.azure.com/v1beta20210401preview.IPRule_Status">
[]IPRule_Status
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
<a href="#keyvault.azure.com/v1beta20210401preview.VirtualNetworkRule_Status">
[]VirtualNetworkRule_Status
</a>
</em>
</td>
<td>
<p>VirtualNetworkRules: The list of virtual network rules.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1beta20210401preview.NetworkRuleSet_StatusARM">NetworkRuleSet_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.VaultProperties_StatusARM">VaultProperties_StatusARM</a>)
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
<code>bypass</code><br/>
<em>
<a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSetStatusBypass">
NetworkRuleSetStatusBypass
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
<a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSetStatusDefaultAction">
NetworkRuleSetStatusDefaultAction
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
<a href="#keyvault.azure.com/v1beta20210401preview.IPRule_StatusARM">
[]IPRule_StatusARM
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
<a href="#keyvault.azure.com/v1beta20210401preview.VirtualNetworkRule_StatusARM">
[]VirtualNetworkRule_StatusARM
</a>
</em>
</td>
<td>
<p>VirtualNetworkRules: The list of virtual network rules.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1beta20210401preview.Permissions">Permissions
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.AccessPolicyEntry">AccessPolicyEntry</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/Permissions">https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/Permissions</a></p>
</div>
<table>
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
<a href="#keyvault.azure.com/v1beta20210401preview.PermissionsCertificates">
[]PermissionsCertificates
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
<a href="#keyvault.azure.com/v1beta20210401preview.PermissionsKeys">
[]PermissionsKeys
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
<a href="#keyvault.azure.com/v1beta20210401preview.PermissionsSecrets">
[]PermissionsSecrets
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
<a href="#keyvault.azure.com/v1beta20210401preview.PermissionsStorage">
[]PermissionsStorage
</a>
</em>
</td>
<td>
<p>Storage: Permissions to storage accounts</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1beta20210401preview.PermissionsARM">PermissionsARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.AccessPolicyEntryARM">AccessPolicyEntryARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/Permissions">https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/Permissions</a></p>
</div>
<table>
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
<a href="#keyvault.azure.com/v1beta20210401preview.PermissionsCertificates">
[]PermissionsCertificates
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
<a href="#keyvault.azure.com/v1beta20210401preview.PermissionsKeys">
[]PermissionsKeys
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
<a href="#keyvault.azure.com/v1beta20210401preview.PermissionsSecrets">
[]PermissionsSecrets
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
<a href="#keyvault.azure.com/v1beta20210401preview.PermissionsStorage">
[]PermissionsStorage
</a>
</em>
</td>
<td>
<p>Storage: Permissions to storage accounts</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1beta20210401preview.PermissionsCertificates">PermissionsCertificates
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.Permissions">Permissions</a>, <a href="#keyvault.azure.com/v1beta20210401preview.PermissionsARM">PermissionsARM</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.PermissionsKeys">PermissionsKeys
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.Permissions">Permissions</a>, <a href="#keyvault.azure.com/v1beta20210401preview.PermissionsARM">PermissionsARM</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.PermissionsSecrets">PermissionsSecrets
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.Permissions">Permissions</a>, <a href="#keyvault.azure.com/v1beta20210401preview.PermissionsARM">PermissionsARM</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.PermissionsStatusCertificates">PermissionsStatusCertificates
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.Permissions_Status">Permissions_Status</a>, <a href="#keyvault.azure.com/v1beta20210401preview.Permissions_StatusARM">Permissions_StatusARM</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.PermissionsStatusKeys">PermissionsStatusKeys
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.Permissions_Status">Permissions_Status</a>, <a href="#keyvault.azure.com/v1beta20210401preview.Permissions_StatusARM">Permissions_StatusARM</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.PermissionsStatusSecrets">PermissionsStatusSecrets
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.Permissions_Status">Permissions_Status</a>, <a href="#keyvault.azure.com/v1beta20210401preview.Permissions_StatusARM">Permissions_StatusARM</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.PermissionsStatusStorage">PermissionsStatusStorage
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.Permissions_Status">Permissions_Status</a>, <a href="#keyvault.azure.com/v1beta20210401preview.Permissions_StatusARM">Permissions_StatusARM</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.PermissionsStorage">PermissionsStorage
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.Permissions">Permissions</a>, <a href="#keyvault.azure.com/v1beta20210401preview.PermissionsARM">PermissionsARM</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.Permissions_Status">Permissions_Status
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.AccessPolicyEntry_Status">AccessPolicyEntry_Status</a>)
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
<code>certificates</code><br/>
<em>
<a href="#keyvault.azure.com/v1beta20210401preview.PermissionsStatusCertificates">
[]PermissionsStatusCertificates
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
<a href="#keyvault.azure.com/v1beta20210401preview.PermissionsStatusKeys">
[]PermissionsStatusKeys
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
<a href="#keyvault.azure.com/v1beta20210401preview.PermissionsStatusSecrets">
[]PermissionsStatusSecrets
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
<a href="#keyvault.azure.com/v1beta20210401preview.PermissionsStatusStorage">
[]PermissionsStatusStorage
</a>
</em>
</td>
<td>
<p>Storage: Permissions to storage accounts</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1beta20210401preview.Permissions_StatusARM">Permissions_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.AccessPolicyEntry_StatusARM">AccessPolicyEntry_StatusARM</a>)
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
<code>certificates</code><br/>
<em>
<a href="#keyvault.azure.com/v1beta20210401preview.PermissionsStatusCertificates">
[]PermissionsStatusCertificates
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
<a href="#keyvault.azure.com/v1beta20210401preview.PermissionsStatusKeys">
[]PermissionsStatusKeys
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
<a href="#keyvault.azure.com/v1beta20210401preview.PermissionsStatusSecrets">
[]PermissionsStatusSecrets
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
<a href="#keyvault.azure.com/v1beta20210401preview.PermissionsStatusStorage">
[]PermissionsStatusStorage
</a>
</em>
</td>
<td>
<p>Storage: Permissions to storage accounts</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1beta20210401preview.PrivateEndpointConnectionItem_Status">PrivateEndpointConnectionItem_Status
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.VaultProperties_Status">VaultProperties_Status</a>)
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
<a href="#keyvault.azure.com/v1beta20210401preview.PrivateEndpoint_Status">
PrivateEndpoint_Status
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
<a href="#keyvault.azure.com/v1beta20210401preview.PrivateLinkServiceConnectionState_Status">
PrivateLinkServiceConnectionState_Status
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
<a href="#keyvault.azure.com/v1beta20210401preview.PrivateEndpointConnectionProvisioningState_Status">
PrivateEndpointConnectionProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: Provisioning state of the private endpoint connection.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1beta20210401preview.PrivateEndpointConnectionItem_StatusARM">PrivateEndpointConnectionItem_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.VaultProperties_StatusARM">VaultProperties_StatusARM</a>)
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
<a href="#keyvault.azure.com/v1beta20210401preview.PrivateEndpointConnectionProperties_StatusARM">
PrivateEndpointConnectionProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Private endpoint connection properties.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1beta20210401preview.PrivateEndpointConnectionProperties_StatusARM">PrivateEndpointConnectionProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.PrivateEndpointConnectionItem_StatusARM">PrivateEndpointConnectionItem_StatusARM</a>)
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
<code>privateEndpoint</code><br/>
<em>
<a href="#keyvault.azure.com/v1beta20210401preview.PrivateEndpoint_StatusARM">
PrivateEndpoint_StatusARM
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
<a href="#keyvault.azure.com/v1beta20210401preview.PrivateLinkServiceConnectionState_StatusARM">
PrivateLinkServiceConnectionState_StatusARM
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
<a href="#keyvault.azure.com/v1beta20210401preview.PrivateEndpointConnectionProvisioningState_Status">
PrivateEndpointConnectionProvisioningState_Status
</a>
</em>
</td>
<td>
<p>ProvisioningState: Provisioning state of the private endpoint connection.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1beta20210401preview.PrivateEndpointConnectionProvisioningState_Status">PrivateEndpointConnectionProvisioningState_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.PrivateEndpointConnectionItem_Status">PrivateEndpointConnectionItem_Status</a>, <a href="#keyvault.azure.com/v1beta20210401preview.PrivateEndpointConnectionProperties_StatusARM">PrivateEndpointConnectionProperties_StatusARM</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.PrivateEndpointServiceConnectionStatus_Status">PrivateEndpointServiceConnectionStatus_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.PrivateLinkServiceConnectionState_Status">PrivateLinkServiceConnectionState_Status</a>, <a href="#keyvault.azure.com/v1beta20210401preview.PrivateLinkServiceConnectionState_StatusARM">PrivateLinkServiceConnectionState_StatusARM</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.PrivateEndpoint_Status">PrivateEndpoint_Status
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.PrivateEndpointConnectionItem_Status">PrivateEndpointConnectionItem_Status</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.PrivateEndpoint_StatusARM">PrivateEndpoint_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.PrivateEndpointConnectionProperties_StatusARM">PrivateEndpointConnectionProperties_StatusARM</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.PrivateLinkServiceConnectionStateStatusActionsRequired">PrivateLinkServiceConnectionStateStatusActionsRequired
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.PrivateLinkServiceConnectionState_Status">PrivateLinkServiceConnectionState_Status</a>, <a href="#keyvault.azure.com/v1beta20210401preview.PrivateLinkServiceConnectionState_StatusARM">PrivateLinkServiceConnectionState_StatusARM</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.PrivateLinkServiceConnectionState_Status">PrivateLinkServiceConnectionState_Status
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.PrivateEndpointConnectionItem_Status">PrivateEndpointConnectionItem_Status</a>)
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
<code>actionsRequired</code><br/>
<em>
<a href="#keyvault.azure.com/v1beta20210401preview.PrivateLinkServiceConnectionStateStatusActionsRequired">
PrivateLinkServiceConnectionStateStatusActionsRequired
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
<a href="#keyvault.azure.com/v1beta20210401preview.PrivateEndpointServiceConnectionStatus_Status">
PrivateEndpointServiceConnectionStatus_Status
</a>
</em>
</td>
<td>
<p>Status: Indicates whether the connection has been approved, rejected or removed by the key vault owner.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1beta20210401preview.PrivateLinkServiceConnectionState_StatusARM">PrivateLinkServiceConnectionState_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.PrivateEndpointConnectionProperties_StatusARM">PrivateEndpointConnectionProperties_StatusARM</a>)
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
<code>actionsRequired</code><br/>
<em>
<a href="#keyvault.azure.com/v1beta20210401preview.PrivateLinkServiceConnectionStateStatusActionsRequired">
PrivateLinkServiceConnectionStateStatusActionsRequired
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
<a href="#keyvault.azure.com/v1beta20210401preview.PrivateEndpointServiceConnectionStatus_Status">
PrivateEndpointServiceConnectionStatus_Status
</a>
</em>
</td>
<td>
<p>Status: Indicates whether the connection has been approved, rejected or removed by the key vault owner.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1beta20210401preview.Sku">Sku
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.VaultProperties">VaultProperties</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/Sku">https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/Sku</a></p>
</div>
<table>
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
<a href="#keyvault.azure.com/v1beta20210401preview.SkuFamily">
SkuFamily
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
<a href="#keyvault.azure.com/v1beta20210401preview.SkuName">
SkuName
</a>
</em>
</td>
<td>
<p>Name: SKU name to specify whether the key vault is a standard vault or a premium vault.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1beta20210401preview.SkuARM">SkuARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.VaultPropertiesARM">VaultPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/Sku">https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/Sku</a></p>
</div>
<table>
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
<a href="#keyvault.azure.com/v1beta20210401preview.SkuFamily">
SkuFamily
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
<a href="#keyvault.azure.com/v1beta20210401preview.SkuName">
SkuName
</a>
</em>
</td>
<td>
<p>Name: SKU name to specify whether the key vault is a standard vault or a premium vault.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1beta20210401preview.SkuFamily">SkuFamily
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.Sku">Sku</a>, <a href="#keyvault.azure.com/v1beta20210401preview.SkuARM">SkuARM</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.SkuName">SkuName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.Sku">Sku</a>, <a href="#keyvault.azure.com/v1beta20210401preview.SkuARM">SkuARM</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.SkuStatusFamily">SkuStatusFamily
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.Sku_Status">Sku_Status</a>, <a href="#keyvault.azure.com/v1beta20210401preview.Sku_StatusARM">Sku_StatusARM</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.SkuStatusName">SkuStatusName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.Sku_Status">Sku_Status</a>, <a href="#keyvault.azure.com/v1beta20210401preview.Sku_StatusARM">Sku_StatusARM</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.Sku_Status">Sku_Status
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.VaultProperties_Status">VaultProperties_Status</a>)
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
<code>family</code><br/>
<em>
<a href="#keyvault.azure.com/v1beta20210401preview.SkuStatusFamily">
SkuStatusFamily
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
<a href="#keyvault.azure.com/v1beta20210401preview.SkuStatusName">
SkuStatusName
</a>
</em>
</td>
<td>
<p>Name: SKU name to specify whether the key vault is a standard vault or a premium vault.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1beta20210401preview.Sku_StatusARM">Sku_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.VaultProperties_StatusARM">VaultProperties_StatusARM</a>)
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
<code>family</code><br/>
<em>
<a href="#keyvault.azure.com/v1beta20210401preview.SkuStatusFamily">
SkuStatusFamily
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
<a href="#keyvault.azure.com/v1beta20210401preview.SkuStatusName">
SkuStatusName
</a>
</em>
</td>
<td>
<p>Name: SKU name to specify whether the key vault is a standard vault or a premium vault.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1beta20210401preview.SystemData_Status">SystemData_Status
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.Vault_Status">Vault_Status</a>)
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
<a href="#keyvault.azure.com/v1beta20210401preview.IdentityType_Status">
IdentityType_Status
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
<a href="#keyvault.azure.com/v1beta20210401preview.IdentityType_Status">
IdentityType_Status
</a>
</em>
</td>
<td>
<p>LastModifiedByType: The type of identity that last modified the key vault resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1beta20210401preview.SystemData_StatusARM">SystemData_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.Vault_StatusARM">Vault_StatusARM</a>)
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
<a href="#keyvault.azure.com/v1beta20210401preview.IdentityType_Status">
IdentityType_Status
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
<a href="#keyvault.azure.com/v1beta20210401preview.IdentityType_Status">
IdentityType_Status
</a>
</em>
</td>
<td>
<p>LastModifiedByType: The type of identity that last modified the key vault resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1beta20210401preview.Vault">Vault
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/resourceDefinitions/vaults">https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/resourceDefinitions/vaults</a></p>
</div>
<table>
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
<a href="#keyvault.azure.com/v1beta20210401preview.Vaults_Spec">
Vaults_Spec
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
<a href="#keyvault.azure.com/v1beta20210401preview.VaultProperties">
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
<a href="#keyvault.azure.com/v1beta20210401preview.Vault_Status">
Vault_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="keyvault.azure.com/v1beta20210401preview.VaultProperties">VaultProperties
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.Vaults_Spec">Vaults_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/VaultProperties">https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/VaultProperties</a></p>
</div>
<table>
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
<a href="#keyvault.azure.com/v1beta20210401preview.AccessPolicyEntry">
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
<a href="#keyvault.azure.com/v1beta20210401preview.VaultPropertiesCreateMode">
VaultPropertiesCreateMode
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
<a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSet">
NetworkRuleSet
</a>
</em>
</td>
<td>
<p>NetworkAcls: A set of rules governing the network accessibility of a vault.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#keyvault.azure.com/v1beta20210401preview.VaultPropertiesProvisioningState">
VaultPropertiesProvisioningState
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
<a href="#keyvault.azure.com/v1beta20210401preview.Sku">
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
<h3 id="keyvault.azure.com/v1beta20210401preview.VaultPropertiesARM">VaultPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.Vaults_SpecARM">Vaults_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/VaultProperties">https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/VaultProperties</a></p>
</div>
<table>
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
<a href="#keyvault.azure.com/v1beta20210401preview.AccessPolicyEntryARM">
[]AccessPolicyEntryARM
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
<a href="#keyvault.azure.com/v1beta20210401preview.VaultPropertiesCreateMode">
VaultPropertiesCreateMode
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
<a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSetARM">
NetworkRuleSetARM
</a>
</em>
</td>
<td>
<p>NetworkAcls: A set of rules governing the network accessibility of a vault.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#keyvault.azure.com/v1beta20210401preview.VaultPropertiesProvisioningState">
VaultPropertiesProvisioningState
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
<a href="#keyvault.azure.com/v1beta20210401preview.SkuARM">
SkuARM
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
<h3 id="keyvault.azure.com/v1beta20210401preview.VaultPropertiesCreateMode">VaultPropertiesCreateMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.VaultProperties">VaultProperties</a>, <a href="#keyvault.azure.com/v1beta20210401preview.VaultPropertiesARM">VaultPropertiesARM</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.VaultPropertiesProvisioningState">VaultPropertiesProvisioningState
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.VaultProperties">VaultProperties</a>, <a href="#keyvault.azure.com/v1beta20210401preview.VaultPropertiesARM">VaultPropertiesARM</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.VaultPropertiesStatusCreateMode">VaultPropertiesStatusCreateMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.VaultProperties_Status">VaultProperties_Status</a>, <a href="#keyvault.azure.com/v1beta20210401preview.VaultProperties_StatusARM">VaultProperties_StatusARM</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.VaultPropertiesStatusProvisioningState">VaultPropertiesStatusProvisioningState
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.VaultProperties_Status">VaultProperties_Status</a>, <a href="#keyvault.azure.com/v1beta20210401preview.VaultProperties_StatusARM">VaultProperties_StatusARM</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.VaultProperties_Status">VaultProperties_Status
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.Vault_Status">Vault_Status</a>)
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
<code>accessPolicies</code><br/>
<em>
<a href="#keyvault.azure.com/v1beta20210401preview.AccessPolicyEntry_Status">
[]AccessPolicyEntry_Status
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
<a href="#keyvault.azure.com/v1beta20210401preview.VaultPropertiesStatusCreateMode">
VaultPropertiesStatusCreateMode
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
<a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSet_Status">
NetworkRuleSet_Status
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
<a href="#keyvault.azure.com/v1beta20210401preview.PrivateEndpointConnectionItem_Status">
[]PrivateEndpointConnectionItem_Status
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
<a href="#keyvault.azure.com/v1beta20210401preview.VaultPropertiesStatusProvisioningState">
VaultPropertiesStatusProvisioningState
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
<a href="#keyvault.azure.com/v1beta20210401preview.Sku_Status">
Sku_Status
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
<h3 id="keyvault.azure.com/v1beta20210401preview.VaultProperties_StatusARM">VaultProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.Vault_StatusARM">Vault_StatusARM</a>)
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
<code>accessPolicies</code><br/>
<em>
<a href="#keyvault.azure.com/v1beta20210401preview.AccessPolicyEntry_StatusARM">
[]AccessPolicyEntry_StatusARM
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
<a href="#keyvault.azure.com/v1beta20210401preview.VaultPropertiesStatusCreateMode">
VaultPropertiesStatusCreateMode
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
<a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSet_StatusARM">
NetworkRuleSet_StatusARM
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
<a href="#keyvault.azure.com/v1beta20210401preview.PrivateEndpointConnectionItem_StatusARM">
[]PrivateEndpointConnectionItem_StatusARM
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
<a href="#keyvault.azure.com/v1beta20210401preview.VaultPropertiesStatusProvisioningState">
VaultPropertiesStatusProvisioningState
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
<a href="#keyvault.azure.com/v1beta20210401preview.Sku_StatusARM">
Sku_StatusARM
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
<h3 id="keyvault.azure.com/v1beta20210401preview.Vault_Status">Vault_Status
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.Vault">Vault</a>)
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
<a href="#keyvault.azure.com/v1beta20210401preview.VaultProperties_Status">
VaultProperties_Status
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
<a href="#keyvault.azure.com/v1beta20210401preview.SystemData_Status">
SystemData_Status
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
<h3 id="keyvault.azure.com/v1beta20210401preview.Vault_StatusARM">Vault_StatusARM
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
<a href="#keyvault.azure.com/v1beta20210401preview.VaultProperties_StatusARM">
VaultProperties_StatusARM
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
<a href="#keyvault.azure.com/v1beta20210401preview.SystemData_StatusARM">
SystemData_StatusARM
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
<h3 id="keyvault.azure.com/v1beta20210401preview.Vaults_Spec">Vaults_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.Vault">Vault</a>)
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
<a href="#keyvault.azure.com/v1beta20210401preview.VaultProperties">
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
<h3 id="keyvault.azure.com/v1beta20210401preview.Vaults_SpecARM">Vaults_SpecARM
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
<p>Name: Name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#keyvault.azure.com/v1beta20210401preview.VaultPropertiesARM">
VaultPropertiesARM
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
<h3 id="keyvault.azure.com/v1beta20210401preview.VirtualNetworkRule">VirtualNetworkRule
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSet">NetworkRuleSet</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/VirtualNetworkRule">https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/VirtualNetworkRule</a></p>
</div>
<table>
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
<h3 id="keyvault.azure.com/v1beta20210401preview.VirtualNetworkRuleARM">VirtualNetworkRuleARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSetARM">NetworkRuleSetARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/VirtualNetworkRule">https://schema.management.azure.com/schemas/2021-04-01-preview/Microsoft.KeyVault.json#/definitions/VirtualNetworkRule</a></p>
</div>
<table>
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
<h3 id="keyvault.azure.com/v1beta20210401preview.VirtualNetworkRule_Status">VirtualNetworkRule_Status
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSet_Status">NetworkRuleSet_Status</a>)
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
<h3 id="keyvault.azure.com/v1beta20210401preview.VirtualNetworkRule_StatusARM">VirtualNetworkRule_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#keyvault.azure.com/v1beta20210401preview.NetworkRuleSet_StatusARM">NetworkRuleSet_StatusARM</a>)
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
