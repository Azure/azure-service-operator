---
title: containerregistry.azure.com/v1beta20210901
---
<h2 id="containerregistry.azure.com/v1beta20210901">containerregistry.azure.com/v1beta20210901</h2>
<div>
<p>Package v1beta20210901 contains API Schema definitions for the containerregistry v1beta20210901 API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="containerregistry.azure.com/v1beta20210901.EncryptionProperty">EncryptionProperty
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Registries_Spec">Registries_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/EncryptionProperty">https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/EncryptionProperty</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>keyVaultProperties</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.KeyVaultProperties">
KeyVaultProperties
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.EncryptionPropertyStatus">
EncryptionPropertyStatus
</a>
</em>
</td>
<td>
<p>Status: Indicates whether or not the encryption is enabled for container registry.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.EncryptionPropertyARM">EncryptionPropertyARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.RegistryPropertiesARM">RegistryPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/EncryptionProperty">https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/EncryptionProperty</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>keyVaultProperties</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.KeyVaultPropertiesARM">
KeyVaultPropertiesARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.EncryptionPropertyStatus">
EncryptionPropertyStatus
</a>
</em>
</td>
<td>
<p>Status: Indicates whether or not the encryption is enabled for container registry.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.EncryptionPropertyStatus">EncryptionPropertyStatus
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.EncryptionProperty">EncryptionProperty</a>, <a href="#containerregistry.azure.com/v1beta20210901.EncryptionPropertyARM">EncryptionPropertyARM</a>)
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
<tbody><tr><td><p>&#34;disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.EncryptionPropertyStatusStatus">EncryptionPropertyStatusStatus
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.EncryptionProperty_Status">EncryptionProperty_Status</a>, <a href="#containerregistry.azure.com/v1beta20210901.EncryptionProperty_StatusARM">EncryptionProperty_StatusARM</a>)
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
<tbody><tr><td><p>&#34;disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.EncryptionProperty_Status">EncryptionProperty_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Registry_Status">Registry_Status</a>)
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
<code>keyVaultProperties</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.KeyVaultProperties_Status">
KeyVaultProperties_Status
</a>
</em>
</td>
<td>
<p>KeyVaultProperties: Key vault properties.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.EncryptionPropertyStatusStatus">
EncryptionPropertyStatusStatus
</a>
</em>
</td>
<td>
<p>Status: Indicates whether or not the encryption is enabled for container registry.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.EncryptionProperty_StatusARM">EncryptionProperty_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.RegistryProperties_StatusARM">RegistryProperties_StatusARM</a>)
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
<code>keyVaultProperties</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.KeyVaultProperties_StatusARM">
KeyVaultProperties_StatusARM
</a>
</em>
</td>
<td>
<p>KeyVaultProperties: Key vault properties.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.EncryptionPropertyStatusStatus">
EncryptionPropertyStatusStatus
</a>
</em>
</td>
<td>
<p>Status: Indicates whether or not the encryption is enabled for container registry.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.ExportPolicy">ExportPolicy
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Policies">Policies</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/ExportPolicy">https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/ExportPolicy</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.ExportPolicyStatus">
ExportPolicyStatus
</a>
</em>
</td>
<td>
<p>Status: The value that indicates whether the policy is enabled or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.ExportPolicyARM">ExportPolicyARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.PoliciesARM">PoliciesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/ExportPolicy">https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/ExportPolicy</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.ExportPolicyStatus">
ExportPolicyStatus
</a>
</em>
</td>
<td>
<p>Status: The value that indicates whether the policy is enabled or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.ExportPolicyStatus">ExportPolicyStatus
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.ExportPolicy">ExportPolicy</a>, <a href="#containerregistry.azure.com/v1beta20210901.ExportPolicyARM">ExportPolicyARM</a>)
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
<tbody><tr><td><p>&#34;disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.ExportPolicyStatusStatus">ExportPolicyStatusStatus
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.ExportPolicy_Status">ExportPolicy_Status</a>, <a href="#containerregistry.azure.com/v1beta20210901.ExportPolicy_StatusARM">ExportPolicy_StatusARM</a>)
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
<tbody><tr><td><p>&#34;disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.ExportPolicy_Status">ExportPolicy_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Policies_Status">Policies_Status</a>)
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
<code>status</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.ExportPolicyStatusStatus">
ExportPolicyStatusStatus
</a>
</em>
</td>
<td>
<p>Status: The value that indicates whether the policy is enabled or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.ExportPolicy_StatusARM">ExportPolicy_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Policies_StatusARM">Policies_StatusARM</a>)
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
<code>status</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.ExportPolicyStatusStatus">
ExportPolicyStatusStatus
</a>
</em>
</td>
<td>
<p>Status: The value that indicates whether the policy is enabled or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.IPRule">IPRule
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.NetworkRuleSet">NetworkRuleSet</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/IPRule">https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/IPRule</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>action</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.IPRuleAction">
IPRuleAction
</a>
</em>
</td>
<td>
<p>Action: The action of IP ACL rule.</p>
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
<p>Value: Specifies the IP or IP range in CIDR format. Only IPV4 address is allowed.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.IPRuleARM">IPRuleARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.NetworkRuleSetARM">NetworkRuleSetARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/IPRule">https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/IPRule</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>action</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.IPRuleAction">
IPRuleAction
</a>
</em>
</td>
<td>
<p>Action: The action of IP ACL rule.</p>
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
<p>Value: Specifies the IP or IP range in CIDR format. Only IPV4 address is allowed.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.IPRuleAction">IPRuleAction
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.IPRule">IPRule</a>, <a href="#containerregistry.azure.com/v1beta20210901.IPRuleARM">IPRuleARM</a>)
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
</tr></tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.IPRuleStatusAction">IPRuleStatusAction
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.IPRule_Status">IPRule_Status</a>, <a href="#containerregistry.azure.com/v1beta20210901.IPRule_StatusARM">IPRule_StatusARM</a>)
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
</tr></tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.IPRule_Status">IPRule_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.NetworkRuleSet_Status">NetworkRuleSet_Status</a>)
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
<code>action</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.IPRuleStatusAction">
IPRuleStatusAction
</a>
</em>
</td>
<td>
<p>Action: The action of IP ACL rule.</p>
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
<p>Value: Specifies the IP or IP range in CIDR format. Only IPV4 address is allowed.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.IPRule_StatusARM">IPRule_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.NetworkRuleSet_StatusARM">NetworkRuleSet_StatusARM</a>)
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
<code>action</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.IPRuleStatusAction">
IPRuleStatusAction
</a>
</em>
</td>
<td>
<p>Action: The action of IP ACL rule.</p>
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
<p>Value: Specifies the IP or IP range in CIDR format. Only IPV4 address is allowed.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.IdentityProperties">IdentityProperties
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Registries_Spec">Registries_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/IdentityProperties">https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/IdentityProperties</a></p>
</div>
<table>
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
<p>PrincipalId: The principal ID of resource identity.</p>
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
<p>TenantId: The tenant ID of resource.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.IdentityPropertiesType">
IdentityPropertiesType
</a>
</em>
</td>
<td>
<p>Type: The identity type.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.UserIdentityProperties">
map[string]./api/containerregistry/v1beta20210901.UserIdentityProperties
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
<h3 id="containerregistry.azure.com/v1beta20210901.IdentityPropertiesARM">IdentityPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Registries_SpecARM">Registries_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/IdentityProperties">https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/IdentityProperties</a></p>
</div>
<table>
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
<p>PrincipalId: The principal ID of resource identity.</p>
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
<p>TenantId: The tenant ID of resource.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.IdentityPropertiesType">
IdentityPropertiesType
</a>
</em>
</td>
<td>
<p>Type: The identity type.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.UserIdentityPropertiesARM">
map[string]./api/containerregistry/v1beta20210901.UserIdentityPropertiesARM
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
<h3 id="containerregistry.azure.com/v1beta20210901.IdentityPropertiesStatusType">IdentityPropertiesStatusType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.IdentityProperties_Status">IdentityProperties_Status</a>, <a href="#containerregistry.azure.com/v1beta20210901.IdentityProperties_StatusARM">IdentityProperties_StatusARM</a>)
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
<h3 id="containerregistry.azure.com/v1beta20210901.IdentityPropertiesType">IdentityPropertiesType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.IdentityProperties">IdentityProperties</a>, <a href="#containerregistry.azure.com/v1beta20210901.IdentityPropertiesARM">IdentityPropertiesARM</a>)
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
<h3 id="containerregistry.azure.com/v1beta20210901.IdentityProperties_Status">IdentityProperties_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Registry_Status">Registry_Status</a>)
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
<code>principalId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrincipalId: The principal ID of resource identity.</p>
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
<p>TenantId: The tenant ID of resource.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.IdentityPropertiesStatusType">
IdentityPropertiesStatusType
</a>
</em>
</td>
<td>
<p>Type: The identity type.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.UserIdentityProperties_Status">
map[string]./api/containerregistry/v1beta20210901.UserIdentityProperties_Status
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
<h3 id="containerregistry.azure.com/v1beta20210901.IdentityProperties_StatusARM">IdentityProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Registry_StatusARM">Registry_StatusARM</a>)
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
<code>principalId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrincipalId: The principal ID of resource identity.</p>
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
<p>TenantId: The tenant ID of resource.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.IdentityPropertiesStatusType">
IdentityPropertiesStatusType
</a>
</em>
</td>
<td>
<p>Type: The identity type.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.UserIdentityProperties_StatusARM">
map[string]./api/containerregistry/v1beta20210901.UserIdentityProperties_StatusARM
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
<h3 id="containerregistry.azure.com/v1beta20210901.KeyVaultProperties">KeyVaultProperties
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.EncryptionProperty">EncryptionProperty</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/KeyVaultProperties">https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/KeyVaultProperties</a></p>
</div>
<table>
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
string
</em>
</td>
<td>
<p>Identity: The client id of the identity which will be used to access key vault.</p>
</td>
</tr>
<tr>
<td>
<code>keyIdentifier</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyIdentifier: Key vault uri to access the encryption key.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.KeyVaultPropertiesARM">KeyVaultPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.EncryptionPropertyARM">EncryptionPropertyARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/KeyVaultProperties">https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/KeyVaultProperties</a></p>
</div>
<table>
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
string
</em>
</td>
<td>
<p>Identity: The client id of the identity which will be used to access key vault.</p>
</td>
</tr>
<tr>
<td>
<code>keyIdentifier</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyIdentifier: Key vault uri to access the encryption key.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.KeyVaultProperties_Status">KeyVaultProperties_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.EncryptionProperty_Status">EncryptionProperty_Status</a>)
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
<code>identity</code><br/>
<em>
string
</em>
</td>
<td>
<p>Identity: The client id of the identity which will be used to access key vault.</p>
</td>
</tr>
<tr>
<td>
<code>keyIdentifier</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyIdentifier: Key vault uri to access the encryption key.</p>
</td>
</tr>
<tr>
<td>
<code>keyRotationEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>KeyRotationEnabled: Auto key rotation status for a CMK enabled registry.</p>
</td>
</tr>
<tr>
<td>
<code>lastKeyRotationTimestamp</code><br/>
<em>
string
</em>
</td>
<td>
<p>LastKeyRotationTimestamp: Timestamp of the last successful key rotation.</p>
</td>
</tr>
<tr>
<td>
<code>versionedKeyIdentifier</code><br/>
<em>
string
</em>
</td>
<td>
<p>VersionedKeyIdentifier: The fully qualified key identifier that includes the version of the key that is actually used
for encryption.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.KeyVaultProperties_StatusARM">KeyVaultProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.EncryptionProperty_StatusARM">EncryptionProperty_StatusARM</a>)
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
<code>identity</code><br/>
<em>
string
</em>
</td>
<td>
<p>Identity: The client id of the identity which will be used to access key vault.</p>
</td>
</tr>
<tr>
<td>
<code>keyIdentifier</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyIdentifier: Key vault uri to access the encryption key.</p>
</td>
</tr>
<tr>
<td>
<code>keyRotationEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>KeyRotationEnabled: Auto key rotation status for a CMK enabled registry.</p>
</td>
</tr>
<tr>
<td>
<code>lastKeyRotationTimestamp</code><br/>
<em>
string
</em>
</td>
<td>
<p>LastKeyRotationTimestamp: Timestamp of the last successful key rotation.</p>
</td>
</tr>
<tr>
<td>
<code>versionedKeyIdentifier</code><br/>
<em>
string
</em>
</td>
<td>
<p>VersionedKeyIdentifier: The fully qualified key identifier that includes the version of the key that is actually used
for encryption.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.NetworkRuleSet">NetworkRuleSet
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Registries_Spec">Registries_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/NetworkRuleSet">https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/NetworkRuleSet</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>defaultAction</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.NetworkRuleSetDefaultAction">
NetworkRuleSetDefaultAction
</a>
</em>
</td>
<td>
<p>DefaultAction: The default action of allow or deny when no other rules match.</p>
</td>
</tr>
<tr>
<td>
<code>ipRules</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.IPRule">
[]IPRule
</a>
</em>
</td>
<td>
<p>IpRules: The IP ACL rules.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.NetworkRuleSetARM">NetworkRuleSetARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.RegistryPropertiesARM">RegistryPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/NetworkRuleSet">https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/NetworkRuleSet</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>defaultAction</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.NetworkRuleSetDefaultAction">
NetworkRuleSetDefaultAction
</a>
</em>
</td>
<td>
<p>DefaultAction: The default action of allow or deny when no other rules match.</p>
</td>
</tr>
<tr>
<td>
<code>ipRules</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.IPRuleARM">
[]IPRuleARM
</a>
</em>
</td>
<td>
<p>IpRules: The IP ACL rules.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.NetworkRuleSetDefaultAction">NetworkRuleSetDefaultAction
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.NetworkRuleSet">NetworkRuleSet</a>, <a href="#containerregistry.azure.com/v1beta20210901.NetworkRuleSetARM">NetworkRuleSetARM</a>)
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
<h3 id="containerregistry.azure.com/v1beta20210901.NetworkRuleSetStatusDefaultAction">NetworkRuleSetStatusDefaultAction
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.NetworkRuleSet_Status">NetworkRuleSet_Status</a>, <a href="#containerregistry.azure.com/v1beta20210901.NetworkRuleSet_StatusARM">NetworkRuleSet_StatusARM</a>)
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
<h3 id="containerregistry.azure.com/v1beta20210901.NetworkRuleSet_Status">NetworkRuleSet_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Registry_Status">Registry_Status</a>)
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
<code>defaultAction</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.NetworkRuleSetStatusDefaultAction">
NetworkRuleSetStatusDefaultAction
</a>
</em>
</td>
<td>
<p>DefaultAction: The default action of allow or deny when no other rules match.</p>
</td>
</tr>
<tr>
<td>
<code>ipRules</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.IPRule_Status">
[]IPRule_Status
</a>
</em>
</td>
<td>
<p>IpRules: The IP ACL rules.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.NetworkRuleSet_StatusARM">NetworkRuleSet_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.RegistryProperties_StatusARM">RegistryProperties_StatusARM</a>)
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
<code>defaultAction</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.NetworkRuleSetStatusDefaultAction">
NetworkRuleSetStatusDefaultAction
</a>
</em>
</td>
<td>
<p>DefaultAction: The default action of allow or deny when no other rules match.</p>
</td>
</tr>
<tr>
<td>
<code>ipRules</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.IPRule_StatusARM">
[]IPRule_StatusARM
</a>
</em>
</td>
<td>
<p>IpRules: The IP ACL rules.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.Policies">Policies
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Registries_Spec">Registries_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/Policies">https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/Policies</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>exportPolicy</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.ExportPolicy">
ExportPolicy
</a>
</em>
</td>
<td>
<p>ExportPolicy: The export policy for a container registry.</p>
</td>
</tr>
<tr>
<td>
<code>quarantinePolicy</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.QuarantinePolicy">
QuarantinePolicy
</a>
</em>
</td>
<td>
<p>QuarantinePolicy: The quarantine policy for a container registry.</p>
</td>
</tr>
<tr>
<td>
<code>retentionPolicy</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RetentionPolicy">
RetentionPolicy
</a>
</em>
</td>
<td>
<p>RetentionPolicy: The retention policy for a container registry.</p>
</td>
</tr>
<tr>
<td>
<code>trustPolicy</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.TrustPolicy">
TrustPolicy
</a>
</em>
</td>
<td>
<p>TrustPolicy: The content trust policy for a container registry.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.PoliciesARM">PoliciesARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.RegistryPropertiesARM">RegistryPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/Policies">https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/Policies</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>exportPolicy</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.ExportPolicyARM">
ExportPolicyARM
</a>
</em>
</td>
<td>
<p>ExportPolicy: The export policy for a container registry.</p>
</td>
</tr>
<tr>
<td>
<code>quarantinePolicy</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.QuarantinePolicyARM">
QuarantinePolicyARM
</a>
</em>
</td>
<td>
<p>QuarantinePolicy: The quarantine policy for a container registry.</p>
</td>
</tr>
<tr>
<td>
<code>retentionPolicy</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RetentionPolicyARM">
RetentionPolicyARM
</a>
</em>
</td>
<td>
<p>RetentionPolicy: The retention policy for a container registry.</p>
</td>
</tr>
<tr>
<td>
<code>trustPolicy</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.TrustPolicyARM">
TrustPolicyARM
</a>
</em>
</td>
<td>
<p>TrustPolicy: The content trust policy for a container registry.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.Policies_Status">Policies_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Registry_Status">Registry_Status</a>)
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
<code>exportPolicy</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.ExportPolicy_Status">
ExportPolicy_Status
</a>
</em>
</td>
<td>
<p>ExportPolicy: The export policy for a container registry.</p>
</td>
</tr>
<tr>
<td>
<code>quarantinePolicy</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.QuarantinePolicy_Status">
QuarantinePolicy_Status
</a>
</em>
</td>
<td>
<p>QuarantinePolicy: The quarantine policy for a container registry.</p>
</td>
</tr>
<tr>
<td>
<code>retentionPolicy</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RetentionPolicy_Status">
RetentionPolicy_Status
</a>
</em>
</td>
<td>
<p>RetentionPolicy: The retention policy for a container registry.</p>
</td>
</tr>
<tr>
<td>
<code>trustPolicy</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.TrustPolicy_Status">
TrustPolicy_Status
</a>
</em>
</td>
<td>
<p>TrustPolicy: The content trust policy for a container registry.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.Policies_StatusARM">Policies_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.RegistryProperties_StatusARM">RegistryProperties_StatusARM</a>)
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
<code>exportPolicy</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.ExportPolicy_StatusARM">
ExportPolicy_StatusARM
</a>
</em>
</td>
<td>
<p>ExportPolicy: The export policy for a container registry.</p>
</td>
</tr>
<tr>
<td>
<code>quarantinePolicy</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.QuarantinePolicy_StatusARM">
QuarantinePolicy_StatusARM
</a>
</em>
</td>
<td>
<p>QuarantinePolicy: The quarantine policy for a container registry.</p>
</td>
</tr>
<tr>
<td>
<code>retentionPolicy</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RetentionPolicy_StatusARM">
RetentionPolicy_StatusARM
</a>
</em>
</td>
<td>
<p>RetentionPolicy: The retention policy for a container registry.</p>
</td>
</tr>
<tr>
<td>
<code>trustPolicy</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.TrustPolicy_StatusARM">
TrustPolicy_StatusARM
</a>
</em>
</td>
<td>
<p>TrustPolicy: The content trust policy for a container registry.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.PrivateEndpointConnection_Status_SubResourceEmbedded">PrivateEndpointConnection_Status_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Registry_Status">Registry_Status</a>)
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
<p>Id: The resource ID.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.SystemData_Status">
SystemData_Status
</a>
</em>
</td>
<td>
<p>SystemData: Metadata pertaining to creation and last modification of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.PrivateEndpointConnection_Status_SubResourceEmbeddedARM">PrivateEndpointConnection_Status_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.RegistryProperties_StatusARM">RegistryProperties_StatusARM</a>)
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
<p>Id: The resource ID.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.SystemData_StatusARM">
SystemData_StatusARM
</a>
</em>
</td>
<td>
<p>SystemData: Metadata pertaining to creation and last modification of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.QuarantinePolicy">QuarantinePolicy
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Policies">Policies</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/QuarantinePolicy">https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/QuarantinePolicy</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.QuarantinePolicyStatus">
QuarantinePolicyStatus
</a>
</em>
</td>
<td>
<p>Status: The value that indicates whether the policy is enabled or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.QuarantinePolicyARM">QuarantinePolicyARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.PoliciesARM">PoliciesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/QuarantinePolicy">https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/QuarantinePolicy</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.QuarantinePolicyStatus">
QuarantinePolicyStatus
</a>
</em>
</td>
<td>
<p>Status: The value that indicates whether the policy is enabled or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.QuarantinePolicyStatus">QuarantinePolicyStatus
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.QuarantinePolicy">QuarantinePolicy</a>, <a href="#containerregistry.azure.com/v1beta20210901.QuarantinePolicyARM">QuarantinePolicyARM</a>)
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
<tbody><tr><td><p>&#34;disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.QuarantinePolicyStatusStatus">QuarantinePolicyStatusStatus
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.QuarantinePolicy_Status">QuarantinePolicy_Status</a>, <a href="#containerregistry.azure.com/v1beta20210901.QuarantinePolicy_StatusARM">QuarantinePolicy_StatusARM</a>)
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
<tbody><tr><td><p>&#34;disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.QuarantinePolicy_Status">QuarantinePolicy_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Policies_Status">Policies_Status</a>)
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
<code>status</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.QuarantinePolicyStatusStatus">
QuarantinePolicyStatusStatus
</a>
</em>
</td>
<td>
<p>Status: The value that indicates whether the policy is enabled or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.QuarantinePolicy_StatusARM">QuarantinePolicy_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Policies_StatusARM">Policies_StatusARM</a>)
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
<code>status</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.QuarantinePolicyStatusStatus">
QuarantinePolicyStatusStatus
</a>
</em>
</td>
<td>
<p>Status: The value that indicates whether the policy is enabled or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.RegistriesSpecAPIVersion">RegistriesSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-09-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.Registries_Spec">Registries_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Registry">Registry</a>)
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
<code>adminUserEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AdminUserEnabled: The value that indicates whether the admin user is enabled.</p>
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
<code>dataEndpointEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DataEndpointEnabled: Enable a single data endpoint per region for serving data.</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.EncryptionProperty">
EncryptionProperty
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.IdentityProperties">
IdentityProperties
</a>
</em>
</td>
<td>
<p>Identity: Managed identity for the resource.</p>
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
<p>Location: The location of the resource. This cannot be changed after the resource is created.</p>
</td>
</tr>
<tr>
<td>
<code>networkRuleBypassOptions</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RegistryPropertiesNetworkRuleBypassOptions">
RegistryPropertiesNetworkRuleBypassOptions
</a>
</em>
</td>
<td>
<p>NetworkRuleBypassOptions: Whether to allow trusted Azure services to access a network restricted registry.</p>
</td>
</tr>
<tr>
<td>
<code>networkRuleSet</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.NetworkRuleSet">
NetworkRuleSet
</a>
</em>
</td>
<td>
<p>NetworkRuleSet: The network rule set for a container registry.</p>
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
<code>policies</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.Policies">
Policies
</a>
</em>
</td>
<td>
<p>Policies: The policies for a container registry.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RegistryPropertiesPublicNetworkAccess">
RegistryPropertiesPublicNetworkAccess
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Whether or not public network access is allowed for the container registry.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.Sku">
Sku
</a>
</em>
</td>
<td>
<p>Sku: The SKU of a container registry.</p>
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
<code>zoneRedundancy</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RegistryPropertiesZoneRedundancy">
RegistryPropertiesZoneRedundancy
</a>
</em>
</td>
<td>
<p>ZoneRedundancy: Whether or not zone redundancy is enabled for this container registry.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.Registries_SpecARM">Registries_SpecARM
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
<a href="#containerregistry.azure.com/v1beta20210901.IdentityPropertiesARM">
IdentityPropertiesARM
</a>
</em>
</td>
<td>
<p>Identity: Managed identity for the resource.</p>
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
<p>Location: The location of the resource. This cannot be changed after the resource is created.</p>
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
<p>Name: The name of the container registry.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RegistryPropertiesARM">
RegistryPropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: The properties of a container registry.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.SkuARM">
SkuARM
</a>
</em>
</td>
<td>
<p>Sku: The SKU of a container registry.</p>
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
<h3 id="containerregistry.azure.com/v1beta20210901.Registry">Registry
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/resourceDefinitions/registries">https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/resourceDefinitions/registries</a></p>
</div>
<table>
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
<a href="#containerregistry.azure.com/v1beta20210901.Registries_Spec">
Registries_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>adminUserEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AdminUserEnabled: The value that indicates whether the admin user is enabled.</p>
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
<code>dataEndpointEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DataEndpointEnabled: Enable a single data endpoint per region for serving data.</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.EncryptionProperty">
EncryptionProperty
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.IdentityProperties">
IdentityProperties
</a>
</em>
</td>
<td>
<p>Identity: Managed identity for the resource.</p>
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
<p>Location: The location of the resource. This cannot be changed after the resource is created.</p>
</td>
</tr>
<tr>
<td>
<code>networkRuleBypassOptions</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RegistryPropertiesNetworkRuleBypassOptions">
RegistryPropertiesNetworkRuleBypassOptions
</a>
</em>
</td>
<td>
<p>NetworkRuleBypassOptions: Whether to allow trusted Azure services to access a network restricted registry.</p>
</td>
</tr>
<tr>
<td>
<code>networkRuleSet</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.NetworkRuleSet">
NetworkRuleSet
</a>
</em>
</td>
<td>
<p>NetworkRuleSet: The network rule set for a container registry.</p>
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
<code>policies</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.Policies">
Policies
</a>
</em>
</td>
<td>
<p>Policies: The policies for a container registry.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RegistryPropertiesPublicNetworkAccess">
RegistryPropertiesPublicNetworkAccess
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Whether or not public network access is allowed for the container registry.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.Sku">
Sku
</a>
</em>
</td>
<td>
<p>Sku: The SKU of a container registry.</p>
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
<code>zoneRedundancy</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RegistryPropertiesZoneRedundancy">
RegistryPropertiesZoneRedundancy
</a>
</em>
</td>
<td>
<p>ZoneRedundancy: Whether or not zone redundancy is enabled for this container registry.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.Registry_Status">
Registry_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.RegistryPropertiesARM">RegistryPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Registries_SpecARM">Registries_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/RegistryProperties">https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/RegistryProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>adminUserEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AdminUserEnabled: The value that indicates whether the admin user is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>dataEndpointEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DataEndpointEnabled: Enable a single data endpoint per region for serving data.</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.EncryptionPropertyARM">
EncryptionPropertyARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>networkRuleBypassOptions</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RegistryPropertiesNetworkRuleBypassOptions">
RegistryPropertiesNetworkRuleBypassOptions
</a>
</em>
</td>
<td>
<p>NetworkRuleBypassOptions: Whether to allow trusted Azure services to access a network restricted registry.</p>
</td>
</tr>
<tr>
<td>
<code>networkRuleSet</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.NetworkRuleSetARM">
NetworkRuleSetARM
</a>
</em>
</td>
<td>
<p>NetworkRuleSet: The network rule set for a container registry.</p>
</td>
</tr>
<tr>
<td>
<code>policies</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.PoliciesARM">
PoliciesARM
</a>
</em>
</td>
<td>
<p>Policies: The policies for a container registry.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RegistryPropertiesPublicNetworkAccess">
RegistryPropertiesPublicNetworkAccess
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Whether or not public network access is allowed for the container registry.</p>
</td>
</tr>
<tr>
<td>
<code>zoneRedundancy</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RegistryPropertiesZoneRedundancy">
RegistryPropertiesZoneRedundancy
</a>
</em>
</td>
<td>
<p>ZoneRedundancy: Whether or not zone redundancy is enabled for this container registry.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.RegistryPropertiesNetworkRuleBypassOptions">RegistryPropertiesNetworkRuleBypassOptions
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Registries_Spec">Registries_Spec</a>, <a href="#containerregistry.azure.com/v1beta20210901.RegistryPropertiesARM">RegistryPropertiesARM</a>)
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
<h3 id="containerregistry.azure.com/v1beta20210901.RegistryPropertiesPublicNetworkAccess">RegistryPropertiesPublicNetworkAccess
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Registries_Spec">Registries_Spec</a>, <a href="#containerregistry.azure.com/v1beta20210901.RegistryPropertiesARM">RegistryPropertiesARM</a>)
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
<h3 id="containerregistry.azure.com/v1beta20210901.RegistryPropertiesStatusNetworkRuleBypassOptions">RegistryPropertiesStatusNetworkRuleBypassOptions
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.RegistryProperties_StatusARM">RegistryProperties_StatusARM</a>, <a href="#containerregistry.azure.com/v1beta20210901.Registry_Status">Registry_Status</a>)
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
<h3 id="containerregistry.azure.com/v1beta20210901.RegistryPropertiesStatusProvisioningState">RegistryPropertiesStatusProvisioningState
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.RegistryProperties_StatusARM">RegistryProperties_StatusARM</a>, <a href="#containerregistry.azure.com/v1beta20210901.Registry_Status">Registry_Status</a>)
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
<tbody><tr><td><p>&#34;Canceled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Creating&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Deleting&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Failed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Succeeded&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Updating&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.RegistryPropertiesStatusPublicNetworkAccess">RegistryPropertiesStatusPublicNetworkAccess
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.RegistryProperties_StatusARM">RegistryProperties_StatusARM</a>, <a href="#containerregistry.azure.com/v1beta20210901.Registry_Status">Registry_Status</a>)
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
<h3 id="containerregistry.azure.com/v1beta20210901.RegistryPropertiesStatusZoneRedundancy">RegistryPropertiesStatusZoneRedundancy
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.RegistryProperties_StatusARM">RegistryProperties_StatusARM</a>, <a href="#containerregistry.azure.com/v1beta20210901.Registry_Status">Registry_Status</a>)
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
<h3 id="containerregistry.azure.com/v1beta20210901.RegistryPropertiesZoneRedundancy">RegistryPropertiesZoneRedundancy
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Registries_Spec">Registries_Spec</a>, <a href="#containerregistry.azure.com/v1beta20210901.RegistryPropertiesARM">RegistryPropertiesARM</a>)
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
<h3 id="containerregistry.azure.com/v1beta20210901.RegistryProperties_StatusARM">RegistryProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Registry_StatusARM">Registry_StatusARM</a>)
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
<code>adminUserEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AdminUserEnabled: The value that indicates whether the admin user is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>creationDate</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreationDate: The creation date of the container registry in ISO8601 format.</p>
</td>
</tr>
<tr>
<td>
<code>dataEndpointEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DataEndpointEnabled: Enable a single data endpoint per region for serving data.</p>
</td>
</tr>
<tr>
<td>
<code>dataEndpointHostNames</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>DataEndpointHostNames: List of host names that will serve data when dataEndpointEnabled is true.</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.EncryptionProperty_StatusARM">
EncryptionProperty_StatusARM
</a>
</em>
</td>
<td>
<p>Encryption: The encryption settings of container registry.</p>
</td>
</tr>
<tr>
<td>
<code>loginServer</code><br/>
<em>
string
</em>
</td>
<td>
<p>LoginServer: The URL that can be used to log into the container registry.</p>
</td>
</tr>
<tr>
<td>
<code>networkRuleBypassOptions</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RegistryPropertiesStatusNetworkRuleBypassOptions">
RegistryPropertiesStatusNetworkRuleBypassOptions
</a>
</em>
</td>
<td>
<p>NetworkRuleBypassOptions: Whether to allow trusted Azure services to access a network restricted registry.</p>
</td>
</tr>
<tr>
<td>
<code>networkRuleSet</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.NetworkRuleSet_StatusARM">
NetworkRuleSet_StatusARM
</a>
</em>
</td>
<td>
<p>NetworkRuleSet: The network rule set for a container registry.</p>
</td>
</tr>
<tr>
<td>
<code>policies</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.Policies_StatusARM">
Policies_StatusARM
</a>
</em>
</td>
<td>
<p>Policies: The policies for a container registry.</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpointConnections</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.PrivateEndpointConnection_Status_SubResourceEmbeddedARM">
[]PrivateEndpointConnection_Status_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>PrivateEndpointConnections: List of private endpoint connections for a container registry.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RegistryPropertiesStatusProvisioningState">
RegistryPropertiesStatusProvisioningState
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the container registry at the time the operation was called.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RegistryPropertiesStatusPublicNetworkAccess">
RegistryPropertiesStatusPublicNetworkAccess
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Whether or not public network access is allowed for the container registry.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.Status_StatusARM">
Status_StatusARM
</a>
</em>
</td>
<td>
<p>Status: The status of the container registry at the time the operation was called.</p>
</td>
</tr>
<tr>
<td>
<code>zoneRedundancy</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RegistryPropertiesStatusZoneRedundancy">
RegistryPropertiesStatusZoneRedundancy
</a>
</em>
</td>
<td>
<p>ZoneRedundancy: Whether or not zone redundancy is enabled for this container registry</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.Registry_Status">Registry_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Registry">Registry</a>)
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
<code>adminUserEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AdminUserEnabled: The value that indicates whether the admin user is enabled.</p>
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
<code>creationDate</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreationDate: The creation date of the container registry in ISO8601 format.</p>
</td>
</tr>
<tr>
<td>
<code>dataEndpointEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DataEndpointEnabled: Enable a single data endpoint per region for serving data.</p>
</td>
</tr>
<tr>
<td>
<code>dataEndpointHostNames</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>DataEndpointHostNames: List of host names that will serve data when dataEndpointEnabled is true.</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.EncryptionProperty_Status">
EncryptionProperty_Status
</a>
</em>
</td>
<td>
<p>Encryption: The encryption settings of container registry.</p>
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
<p>Id: The resource ID.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.IdentityProperties_Status">
IdentityProperties_Status
</a>
</em>
</td>
<td>
<p>Identity: The identity of the container registry.</p>
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
<p>Location: The location of the resource. This cannot be changed after the resource is created.</p>
</td>
</tr>
<tr>
<td>
<code>loginServer</code><br/>
<em>
string
</em>
</td>
<td>
<p>LoginServer: The URL that can be used to log into the container registry.</p>
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
<p>Name: The name of the resource.</p>
</td>
</tr>
<tr>
<td>
<code>networkRuleBypassOptions</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RegistryPropertiesStatusNetworkRuleBypassOptions">
RegistryPropertiesStatusNetworkRuleBypassOptions
</a>
</em>
</td>
<td>
<p>NetworkRuleBypassOptions: Whether to allow trusted Azure services to access a network restricted registry.</p>
</td>
</tr>
<tr>
<td>
<code>networkRuleSet</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.NetworkRuleSet_Status">
NetworkRuleSet_Status
</a>
</em>
</td>
<td>
<p>NetworkRuleSet: The network rule set for a container registry.</p>
</td>
</tr>
<tr>
<td>
<code>policies</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.Policies_Status">
Policies_Status
</a>
</em>
</td>
<td>
<p>Policies: The policies for a container registry.</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpointConnections</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.PrivateEndpointConnection_Status_SubResourceEmbedded">
[]PrivateEndpointConnection_Status_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>PrivateEndpointConnections: List of private endpoint connections for a container registry.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RegistryPropertiesStatusProvisioningState">
RegistryPropertiesStatusProvisioningState
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the container registry at the time the operation was called.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RegistryPropertiesStatusPublicNetworkAccess">
RegistryPropertiesStatusPublicNetworkAccess
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Whether or not public network access is allowed for the container registry.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.Sku_Status">
Sku_Status
</a>
</em>
</td>
<td>
<p>Sku: The SKU of the container registry.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.Status_Status">
Status_Status
</a>
</em>
</td>
<td>
<p>Status: The status of the container registry at the time the operation was called.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.SystemData_Status">
SystemData_Status
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
<p>Tags: The tags of the resource.</p>
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
<p>Type: The type of the resource.</p>
</td>
</tr>
<tr>
<td>
<code>zoneRedundancy</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RegistryPropertiesStatusZoneRedundancy">
RegistryPropertiesStatusZoneRedundancy
</a>
</em>
</td>
<td>
<p>ZoneRedundancy: Whether or not zone redundancy is enabled for this container registry</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.Registry_StatusARM">Registry_StatusARM
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
<p>Id: The resource ID.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.IdentityProperties_StatusARM">
IdentityProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Identity: The identity of the container registry.</p>
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
<p>Location: The location of the resource. This cannot be changed after the resource is created.</p>
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
<p>Name: The name of the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RegistryProperties_StatusARM">
RegistryProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: The properties of the container registry.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.Sku_StatusARM">
Sku_StatusARM
</a>
</em>
</td>
<td>
<p>Sku: The SKU of the container registry.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.SystemData_StatusARM">
SystemData_StatusARM
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
<p>Tags: The tags of the resource.</p>
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
<p>Type: The type of the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.RetentionPolicy">RetentionPolicy
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Policies">Policies</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/RetentionPolicy">https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/RetentionPolicy</a></p>
</div>
<table>
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
<p>Days: The number of days to retain an untagged manifest after which it gets purged.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RetentionPolicyStatus">
RetentionPolicyStatus
</a>
</em>
</td>
<td>
<p>Status: The value that indicates whether the policy is enabled or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.RetentionPolicyARM">RetentionPolicyARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.PoliciesARM">PoliciesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/RetentionPolicy">https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/RetentionPolicy</a></p>
</div>
<table>
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
<p>Days: The number of days to retain an untagged manifest after which it gets purged.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RetentionPolicyStatus">
RetentionPolicyStatus
</a>
</em>
</td>
<td>
<p>Status: The value that indicates whether the policy is enabled or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.RetentionPolicyStatus">RetentionPolicyStatus
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.RetentionPolicy">RetentionPolicy</a>, <a href="#containerregistry.azure.com/v1beta20210901.RetentionPolicyARM">RetentionPolicyARM</a>)
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
<tbody><tr><td><p>&#34;disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.RetentionPolicyStatusStatus">RetentionPolicyStatusStatus
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.RetentionPolicy_Status">RetentionPolicy_Status</a>, <a href="#containerregistry.azure.com/v1beta20210901.RetentionPolicy_StatusARM">RetentionPolicy_StatusARM</a>)
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
<tbody><tr><td><p>&#34;disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.RetentionPolicy_Status">RetentionPolicy_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Policies_Status">Policies_Status</a>)
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
<code>days</code><br/>
<em>
int
</em>
</td>
<td>
<p>Days: The number of days to retain an untagged manifest after which it gets purged.</p>
</td>
</tr>
<tr>
<td>
<code>lastUpdatedTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>LastUpdatedTime: The timestamp when the policy was last updated.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RetentionPolicyStatusStatus">
RetentionPolicyStatusStatus
</a>
</em>
</td>
<td>
<p>Status: The value that indicates whether the policy is enabled or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.RetentionPolicy_StatusARM">RetentionPolicy_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Policies_StatusARM">Policies_StatusARM</a>)
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
<code>days</code><br/>
<em>
int
</em>
</td>
<td>
<p>Days: The number of days to retain an untagged manifest after which it gets purged.</p>
</td>
</tr>
<tr>
<td>
<code>lastUpdatedTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>LastUpdatedTime: The timestamp when the policy was last updated.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.RetentionPolicyStatusStatus">
RetentionPolicyStatusStatus
</a>
</em>
</td>
<td>
<p>Status: The value that indicates whether the policy is enabled or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.Sku">Sku
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Registries_Spec">Registries_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/Sku">https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/Sku</a></p>
</div>
<table>
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
<a href="#containerregistry.azure.com/v1beta20210901.SkuName">
SkuName
</a>
</em>
</td>
<td>
<p>Name: The SKU name of the container registry. Required for registry creation.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.SkuARM">SkuARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Registries_SpecARM">Registries_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/Sku">https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/Sku</a></p>
</div>
<table>
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
<a href="#containerregistry.azure.com/v1beta20210901.SkuName">
SkuName
</a>
</em>
</td>
<td>
<p>Name: The SKU name of the container registry. Required for registry creation.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.SkuName">SkuName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Sku">Sku</a>, <a href="#containerregistry.azure.com/v1beta20210901.SkuARM">SkuARM</a>)
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
</tr><tr><td><p>&#34;Classic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Premium&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.SkuStatusName">SkuStatusName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Sku_Status">Sku_Status</a>, <a href="#containerregistry.azure.com/v1beta20210901.Sku_StatusARM">Sku_StatusARM</a>)
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
</tr><tr><td><p>&#34;Classic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Premium&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.SkuStatusTier">SkuStatusTier
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Sku_Status">Sku_Status</a>, <a href="#containerregistry.azure.com/v1beta20210901.Sku_StatusARM">Sku_StatusARM</a>)
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
</tr><tr><td><p>&#34;Classic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Premium&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.Sku_Status">Sku_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Registry_Status">Registry_Status</a>)
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
<code>name</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.SkuStatusName">
SkuStatusName
</a>
</em>
</td>
<td>
<p>Name: The SKU name of the container registry. Required for registry creation.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.SkuStatusTier">
SkuStatusTier
</a>
</em>
</td>
<td>
<p>Tier: The SKU tier based on the SKU name.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.Sku_StatusARM">Sku_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Registry_StatusARM">Registry_StatusARM</a>)
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
<code>name</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.SkuStatusName">
SkuStatusName
</a>
</em>
</td>
<td>
<p>Name: The SKU name of the container registry. Required for registry creation.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.SkuStatusTier">
SkuStatusTier
</a>
</em>
</td>
<td>
<p>Tier: The SKU tier based on the SKU name.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.Status_Status">Status_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Registry_Status">Registry_Status</a>)
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
<code>displayStatus</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayStatus: The short label for the status.</p>
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
<p>Message: The detailed message for the status, including alerts and error messages.</p>
</td>
</tr>
<tr>
<td>
<code>timestamp</code><br/>
<em>
string
</em>
</td>
<td>
<p>Timestamp: The timestamp when the status was changed to the current value.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.Status_StatusARM">Status_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.RegistryProperties_StatusARM">RegistryProperties_StatusARM</a>)
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
<code>displayStatus</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayStatus: The short label for the status.</p>
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
<p>Message: The detailed message for the status, including alerts and error messages.</p>
</td>
</tr>
<tr>
<td>
<code>timestamp</code><br/>
<em>
string
</em>
</td>
<td>
<p>Timestamp: The timestamp when the status was changed to the current value.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.SystemDataStatusCreatedByType">SystemDataStatusCreatedByType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.SystemData_Status">SystemData_Status</a>, <a href="#containerregistry.azure.com/v1beta20210901.SystemData_StatusARM">SystemData_StatusARM</a>)
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
<h3 id="containerregistry.azure.com/v1beta20210901.SystemDataStatusLastModifiedByType">SystemDataStatusLastModifiedByType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.SystemData_Status">SystemData_Status</a>, <a href="#containerregistry.azure.com/v1beta20210901.SystemData_StatusARM">SystemData_StatusARM</a>)
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
<h3 id="containerregistry.azure.com/v1beta20210901.SystemData_Status">SystemData_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.PrivateEndpointConnection_Status_SubResourceEmbedded">PrivateEndpointConnection_Status_SubResourceEmbedded</a>, <a href="#containerregistry.azure.com/v1beta20210901.Registry_Status">Registry_Status</a>)
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
<a href="#containerregistry.azure.com/v1beta20210901.SystemDataStatusCreatedByType">
SystemDataStatusCreatedByType
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
<p>LastModifiedAt: The timestamp of resource modification (UTC).</p>
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
<a href="#containerregistry.azure.com/v1beta20210901.SystemDataStatusLastModifiedByType">
SystemDataStatusLastModifiedByType
</a>
</em>
</td>
<td>
<p>LastModifiedByType: The type of identity that last modified the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.SystemData_StatusARM">SystemData_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.PrivateEndpointConnection_Status_SubResourceEmbeddedARM">PrivateEndpointConnection_Status_SubResourceEmbeddedARM</a>, <a href="#containerregistry.azure.com/v1beta20210901.Registry_StatusARM">Registry_StatusARM</a>)
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
<a href="#containerregistry.azure.com/v1beta20210901.SystemDataStatusCreatedByType">
SystemDataStatusCreatedByType
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
<p>LastModifiedAt: The timestamp of resource modification (UTC).</p>
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
<a href="#containerregistry.azure.com/v1beta20210901.SystemDataStatusLastModifiedByType">
SystemDataStatusLastModifiedByType
</a>
</em>
</td>
<td>
<p>LastModifiedByType: The type of identity that last modified the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.TrustPolicy">TrustPolicy
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Policies">Policies</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/TrustPolicy">https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/TrustPolicy</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.TrustPolicyStatus">
TrustPolicyStatus
</a>
</em>
</td>
<td>
<p>Status: The value that indicates whether the policy is enabled or not.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.TrustPolicyType">
TrustPolicyType
</a>
</em>
</td>
<td>
<p>Type: The type of trust policy.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.TrustPolicyARM">TrustPolicyARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.PoliciesARM">PoliciesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/TrustPolicy">https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/TrustPolicy</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.TrustPolicyStatus">
TrustPolicyStatus
</a>
</em>
</td>
<td>
<p>Status: The value that indicates whether the policy is enabled or not.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.TrustPolicyType">
TrustPolicyType
</a>
</em>
</td>
<td>
<p>Type: The type of trust policy.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.TrustPolicyStatus">TrustPolicyStatus
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.TrustPolicy">TrustPolicy</a>, <a href="#containerregistry.azure.com/v1beta20210901.TrustPolicyARM">TrustPolicyARM</a>)
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
<tbody><tr><td><p>&#34;disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.TrustPolicyStatusStatus">TrustPolicyStatusStatus
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.TrustPolicy_Status">TrustPolicy_Status</a>, <a href="#containerregistry.azure.com/v1beta20210901.TrustPolicy_StatusARM">TrustPolicy_StatusARM</a>)
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
<tbody><tr><td><p>&#34;disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.TrustPolicyStatusType">TrustPolicyStatusType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.TrustPolicy_Status">TrustPolicy_Status</a>, <a href="#containerregistry.azure.com/v1beta20210901.TrustPolicy_StatusARM">TrustPolicy_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Notary&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.TrustPolicyType">TrustPolicyType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.TrustPolicy">TrustPolicy</a>, <a href="#containerregistry.azure.com/v1beta20210901.TrustPolicyARM">TrustPolicyARM</a>)
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
<tbody><tr><td><p>&#34;Notary&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.TrustPolicy_Status">TrustPolicy_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Policies_Status">Policies_Status</a>)
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
<code>status</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.TrustPolicyStatusStatus">
TrustPolicyStatusStatus
</a>
</em>
</td>
<td>
<p>Status: The value that indicates whether the policy is enabled or not.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.TrustPolicyStatusType">
TrustPolicyStatusType
</a>
</em>
</td>
<td>
<p>Type: The type of trust policy.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.TrustPolicy_StatusARM">TrustPolicy_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.Policies_StatusARM">Policies_StatusARM</a>)
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
<code>status</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.TrustPolicyStatusStatus">
TrustPolicyStatusStatus
</a>
</em>
</td>
<td>
<p>Status: The value that indicates whether the policy is enabled or not.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#containerregistry.azure.com/v1beta20210901.TrustPolicyStatusType">
TrustPolicyStatusType
</a>
</em>
</td>
<td>
<p>Type: The type of trust policy.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerregistry.azure.com/v1beta20210901.UserIdentityProperties">UserIdentityProperties
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.IdentityProperties">IdentityProperties</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/UserIdentityProperties">https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/UserIdentityProperties</a></p>
</div>
<table>
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
<h3 id="containerregistry.azure.com/v1beta20210901.UserIdentityPropertiesARM">UserIdentityPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.IdentityPropertiesARM">IdentityPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/UserIdentityProperties">https://schema.management.azure.com/schemas/2021-09-01/Microsoft.ContainerRegistry.json#/definitions/UserIdentityProperties</a></p>
</div>
<table>
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
<h3 id="containerregistry.azure.com/v1beta20210901.UserIdentityProperties_Status">UserIdentityProperties_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.IdentityProperties_Status">IdentityProperties_Status</a>)
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
<h3 id="containerregistry.azure.com/v1beta20210901.UserIdentityProperties_StatusARM">UserIdentityProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerregistry.azure.com/v1beta20210901.IdentityProperties_StatusARM">IdentityProperties_StatusARM</a>)
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
<hr/>
