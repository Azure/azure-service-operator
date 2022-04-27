---
title: batch.azure.com/v1beta20210101
---
<h2 id="batch.azure.com/v1beta20210101">batch.azure.com/v1beta20210101</h2>
<div>
<p>Package v1beta20210101 contains API Schema definitions for the batch v1beta20210101 API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="batch.azure.com/v1beta20210101.AutoStorageBaseProperties">AutoStorageBaseProperties
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccounts_Spec">BatchAccounts_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/AutoStorageBaseProperties">https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/AutoStorageBaseProperties</a></p>
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
<code>storageAccountReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>StorageAccountReference: The resource ID of the storage account to be used for auto-storage account.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.AutoStorageBasePropertiesARM">AutoStorageBasePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccountCreatePropertiesARM">BatchAccountCreatePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/AutoStorageBaseProperties">https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/AutoStorageBaseProperties</a></p>
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
<code>storageAccountId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.AutoStorageProperties_Status">AutoStorageProperties_Status
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccount_Status">BatchAccount_Status</a>)
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
<code>lastKeySync</code><br/>
<em>
string
</em>
</td>
<td>
<p>LastKeySync: The UTC time at which storage keys were last synchronized with the Batch account.</p>
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
<p>StorageAccountId: The resource ID of the storage account to be used for auto-storage account.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.AutoStorageProperties_StatusARM">AutoStorageProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccountProperties_StatusARM">BatchAccountProperties_StatusARM</a>)
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
<code>lastKeySync</code><br/>
<em>
string
</em>
</td>
<td>
<p>LastKeySync: The UTC time at which storage keys were last synchronized with the Batch account.</p>
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
<p>StorageAccountId: The resource ID of the storage account to be used for auto-storage account.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.BatchAccount">BatchAccount
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/resourceDefinitions/batchAccounts">https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/resourceDefinitions/batchAccounts</a></p>
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
<a href="#batch.azure.com/v1beta20210101.BatchAccounts_Spec">
BatchAccounts_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>autoStorage</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.AutoStorageBaseProperties">
AutoStorageBaseProperties
</a>
</em>
</td>
<td>
<p>AutoStorage: The properties related to the auto-storage account.</p>
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
<code>encryption</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.EncryptionProperties">
EncryptionProperties
</a>
</em>
</td>
<td>
<p>Encryption: Configures how customer data is encrypted inside the Batch account. By default, accounts are encrypted using
a Microsoft managed key. For additional control, a customer-managed key can be used instead.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.BatchAccountIdentity">
BatchAccountIdentity
</a>
</em>
</td>
<td>
<p>Identity: The identity of the Batch account, if configured. This is only used when the user specifies
&lsquo;Microsoft.KeyVault&rsquo; as their Batch account encryption configuration.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultReference</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.KeyVaultReference">
KeyVaultReference
</a>
</em>
</td>
<td>
<p>KeyVaultReference: Identifies the Azure key vault associated with a Batch account.</p>
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
<p>Location: The region in which to create the account.</p>
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
<code>poolAllocationMode</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.BatchAccountCreatePropertiesPoolAllocationMode">
BatchAccountCreatePropertiesPoolAllocationMode
</a>
</em>
</td>
<td>
<p>PoolAllocationMode: The pool allocation mode also affects how clients may authenticate to the Batch Service API. If the
mode is BatchService, clients may authenticate using access keys or Azure Active Directory. If the mode is
UserSubscription, clients must use Azure Active Directory. The default is BatchService.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.BatchAccountCreatePropertiesPublicNetworkAccess">
BatchAccountCreatePropertiesPublicNetworkAccess
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: If not specified, the default value is &lsquo;enabled&rsquo;.</p>
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
<p>Tags: The user-specified tags associated with the account.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.BatchAccount_Status">
BatchAccount_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.BatchAccountCreatePropertiesARM">BatchAccountCreatePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccounts_SpecARM">BatchAccounts_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/BatchAccountCreateProperties">https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/BatchAccountCreateProperties</a></p>
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
<code>autoStorage</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.AutoStorageBasePropertiesARM">
AutoStorageBasePropertiesARM
</a>
</em>
</td>
<td>
<p>AutoStorage: The properties related to the auto-storage account.</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.EncryptionPropertiesARM">
EncryptionPropertiesARM
</a>
</em>
</td>
<td>
<p>Encryption: Configures how customer data is encrypted inside the Batch account. By default, accounts are encrypted using
a Microsoft managed key. For additional control, a customer-managed key can be used instead.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultReference</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.KeyVaultReferenceARM">
KeyVaultReferenceARM
</a>
</em>
</td>
<td>
<p>KeyVaultReference: Identifies the Azure key vault associated with a Batch account.</p>
</td>
</tr>
<tr>
<td>
<code>poolAllocationMode</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.BatchAccountCreatePropertiesPoolAllocationMode">
BatchAccountCreatePropertiesPoolAllocationMode
</a>
</em>
</td>
<td>
<p>PoolAllocationMode: The pool allocation mode also affects how clients may authenticate to the Batch Service API. If the
mode is BatchService, clients may authenticate using access keys or Azure Active Directory. If the mode is
UserSubscription, clients must use Azure Active Directory. The default is BatchService.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.BatchAccountCreatePropertiesPublicNetworkAccess">
BatchAccountCreatePropertiesPublicNetworkAccess
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: If not specified, the default value is &lsquo;enabled&rsquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.BatchAccountCreatePropertiesPoolAllocationMode">BatchAccountCreatePropertiesPoolAllocationMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccountCreatePropertiesARM">BatchAccountCreatePropertiesARM</a>, <a href="#batch.azure.com/v1beta20210101.BatchAccounts_Spec">BatchAccounts_Spec</a>)
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
<tbody><tr><td><p>&#34;BatchService&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UserSubscription&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.BatchAccountCreatePropertiesPublicNetworkAccess">BatchAccountCreatePropertiesPublicNetworkAccess
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccountCreatePropertiesARM">BatchAccountCreatePropertiesARM</a>, <a href="#batch.azure.com/v1beta20210101.BatchAccounts_Spec">BatchAccounts_Spec</a>)
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
<h3 id="batch.azure.com/v1beta20210101.BatchAccountIdentity">BatchAccountIdentity
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccounts_Spec">BatchAccounts_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/BatchAccountIdentity">https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/BatchAccountIdentity</a></p>
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
<code>type</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.BatchAccountIdentityType">
BatchAccountIdentityType
</a>
</em>
</td>
<td>
<p>Type: The type of identity used for the Batch account.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.BatchAccountIdentityARM">BatchAccountIdentityARM
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccounts_SpecARM">BatchAccounts_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/BatchAccountIdentity">https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/BatchAccountIdentity</a></p>
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
<code>type</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.BatchAccountIdentityType">
BatchAccountIdentityType
</a>
</em>
</td>
<td>
<p>Type: The type of identity used for the Batch account.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.BatchAccountIdentityStatusType">BatchAccountIdentityStatusType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccountIdentity_Status">BatchAccountIdentity_Status</a>, <a href="#batch.azure.com/v1beta20210101.BatchAccountIdentity_StatusARM">BatchAccountIdentity_StatusARM</a>)
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
</tr><tr><td><p>&#34;UserAssigned&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.BatchAccountIdentityType">BatchAccountIdentityType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccountIdentity">BatchAccountIdentity</a>, <a href="#batch.azure.com/v1beta20210101.BatchAccountIdentityARM">BatchAccountIdentityARM</a>)
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
</tr><tr><td><p>&#34;UserAssigned&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.BatchAccountIdentity_Status">BatchAccountIdentity_Status
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccount_Status">BatchAccount_Status</a>)
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
<p>PrincipalId: The principal id of the Batch account. This property will only be provided for a system assigned identity.</p>
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
<p>TenantId: The tenant id associated with the Batch account. This property will only be provided for a system assigned
identity.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.BatchAccountIdentityStatusType">
BatchAccountIdentityStatusType
</a>
</em>
</td>
<td>
<p>Type: The type of identity used for the Batch account.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.BatchAccountIdentity_Status_UserAssignedIdentities">
map[string]./api/batch/v1beta20210101.BatchAccountIdentity_Status_UserAssignedIdentities
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: The list of user identities associated with the Batch account. The user identity dictionary key
references will be ARM resource ids in the form:
&lsquo;/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}&rsquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.BatchAccountIdentity_StatusARM">BatchAccountIdentity_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccount_StatusARM">BatchAccount_StatusARM</a>)
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
<p>PrincipalId: The principal id of the Batch account. This property will only be provided for a system assigned identity.</p>
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
<p>TenantId: The tenant id associated with the Batch account. This property will only be provided for a system assigned
identity.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.BatchAccountIdentityStatusType">
BatchAccountIdentityStatusType
</a>
</em>
</td>
<td>
<p>Type: The type of identity used for the Batch account.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.BatchAccountIdentity_Status_UserAssignedIdentitiesARM">
map[string]./api/batch/v1beta20210101.BatchAccountIdentity_Status_UserAssignedIdentitiesARM
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: The list of user identities associated with the Batch account. The user identity dictionary key
references will be ARM resource ids in the form:
&lsquo;/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}&rsquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.BatchAccountIdentity_Status_UserAssignedIdentities">BatchAccountIdentity_Status_UserAssignedIdentities
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccountIdentity_Status">BatchAccountIdentity_Status</a>)
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
<h3 id="batch.azure.com/v1beta20210101.BatchAccountIdentity_Status_UserAssignedIdentitiesARM">BatchAccountIdentity_Status_UserAssignedIdentitiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccountIdentity_StatusARM">BatchAccountIdentity_StatusARM</a>)
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
<h3 id="batch.azure.com/v1beta20210101.BatchAccountPropertiesStatusProvisioningState">BatchAccountPropertiesStatusProvisioningState
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccountProperties_StatusARM">BatchAccountProperties_StatusARM</a>, <a href="#batch.azure.com/v1beta20210101.BatchAccount_Status">BatchAccount_Status</a>)
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
<tbody><tr><td><p>&#34;Cancelled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Creating&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Deleting&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Failed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Invalid&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Succeeded&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.BatchAccountProperties_StatusARM">BatchAccountProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccount_StatusARM">BatchAccount_StatusARM</a>)
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
<code>accountEndpoint</code><br/>
<em>
string
</em>
</td>
<td>
<p>AccountEndpoint: The account endpoint used to interact with the Batch service.</p>
</td>
</tr>
<tr>
<td>
<code>activeJobAndJobScheduleQuota</code><br/>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>autoStorage</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.AutoStorageProperties_StatusARM">
AutoStorageProperties_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>dedicatedCoreQuota</code><br/>
<em>
int
</em>
</td>
<td>
<p>DedicatedCoreQuota: For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription
so this value is not returned.</p>
</td>
</tr>
<tr>
<td>
<code>dedicatedCoreQuotaPerVMFamily</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.VirtualMachineFamilyCoreQuota_StatusARM">
[]VirtualMachineFamilyCoreQuota_StatusARM
</a>
</em>
</td>
<td>
<p>DedicatedCoreQuotaPerVMFamily: A list of the dedicated core quota per Virtual Machine family for the Batch account. For
accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not
returned.</p>
</td>
</tr>
<tr>
<td>
<code>dedicatedCoreQuotaPerVMFamilyEnforced</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DedicatedCoreQuotaPerVMFamilyEnforced: Batch is transitioning its core quota system for dedicated cores to be enforced
per Virtual Machine family. During this transitional phase, the dedicated core quota per Virtual Machine family may not
yet be enforced. If this flag is false, dedicated core quota is enforced via the old dedicatedCoreQuota property on the
account and does not consider Virtual Machine family. If this flag is true, dedicated core quota is enforced via the
dedicatedCoreQuotaPerVMFamily property on the account, and the old dedicatedCoreQuota does not apply.</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.EncryptionProperties_StatusARM">
EncryptionProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Encryption: Configures how customer data is encrypted inside the Batch account. By default, accounts are encrypted using
a Microsoft managed key. For additional control, a customer-managed key can be used instead.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultReference</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.KeyVaultReference_StatusARM">
KeyVaultReference_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>lowPriorityCoreQuota</code><br/>
<em>
int
</em>
</td>
<td>
<p>LowPriorityCoreQuota: For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription
so this value is not returned.</p>
</td>
</tr>
<tr>
<td>
<code>poolAllocationMode</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.PoolAllocationMode_Status">
PoolAllocationMode_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>poolQuota</code><br/>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>privateEndpointConnections</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.PrivateEndpointConnection_StatusARM">
[]PrivateEndpointConnection_StatusARM
</a>
</em>
</td>
<td>
<p>PrivateEndpointConnections: List of private endpoint connections associated with the Batch account</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.BatchAccountPropertiesStatusProvisioningState">
BatchAccountPropertiesStatusProvisioningState
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioned state of the resource</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.PublicNetworkAccessType_Status">
PublicNetworkAccessType_Status
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: If not specified, the default value is &lsquo;enabled&rsquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.BatchAccount_Status">BatchAccount_Status
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccount">BatchAccount</a>)
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
<code>accountEndpoint</code><br/>
<em>
string
</em>
</td>
<td>
<p>AccountEndpoint: The account endpoint used to interact with the Batch service.</p>
</td>
</tr>
<tr>
<td>
<code>activeJobAndJobScheduleQuota</code><br/>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>autoStorage</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.AutoStorageProperties_Status">
AutoStorageProperties_Status
</a>
</em>
</td>
<td>
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
<code>dedicatedCoreQuota</code><br/>
<em>
int
</em>
</td>
<td>
<p>DedicatedCoreQuota: For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription
so this value is not returned.</p>
</td>
</tr>
<tr>
<td>
<code>dedicatedCoreQuotaPerVMFamily</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.VirtualMachineFamilyCoreQuota_Status">
[]VirtualMachineFamilyCoreQuota_Status
</a>
</em>
</td>
<td>
<p>DedicatedCoreQuotaPerVMFamily: A list of the dedicated core quota per Virtual Machine family for the Batch account. For
accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription so this value is not
returned.</p>
</td>
</tr>
<tr>
<td>
<code>dedicatedCoreQuotaPerVMFamilyEnforced</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DedicatedCoreQuotaPerVMFamilyEnforced: Batch is transitioning its core quota system for dedicated cores to be enforced
per Virtual Machine family. During this transitional phase, the dedicated core quota per Virtual Machine family may not
yet be enforced. If this flag is false, dedicated core quota is enforced via the old dedicatedCoreQuota property on the
account and does not consider Virtual Machine family. If this flag is true, dedicated core quota is enforced via the
dedicatedCoreQuotaPerVMFamily property on the account, and the old dedicatedCoreQuota does not apply.</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.EncryptionProperties_Status">
EncryptionProperties_Status
</a>
</em>
</td>
<td>
<p>Encryption: Configures how customer data is encrypted inside the Batch account. By default, accounts are encrypted using
a Microsoft managed key. For additional control, a customer-managed key can be used instead.</p>
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
<p>Id: The ID of the resource.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.BatchAccountIdentity_Status">
BatchAccountIdentity_Status
</a>
</em>
</td>
<td>
<p>Identity: The identity of the Batch account.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultReference</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.KeyVaultReference_Status">
KeyVaultReference_Status
</a>
</em>
</td>
<td>
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
<p>Location: The location of the resource.</p>
</td>
</tr>
<tr>
<td>
<code>lowPriorityCoreQuota</code><br/>
<em>
int
</em>
</td>
<td>
<p>LowPriorityCoreQuota: For accounts with PoolAllocationMode set to UserSubscription, quota is managed on the subscription
so this value is not returned.</p>
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
<code>poolAllocationMode</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.PoolAllocationMode_Status">
PoolAllocationMode_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>poolQuota</code><br/>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>privateEndpointConnections</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.PrivateEndpointConnection_Status">
[]PrivateEndpointConnection_Status
</a>
</em>
</td>
<td>
<p>PrivateEndpointConnections: List of private endpoint connections associated with the Batch account</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.BatchAccountPropertiesStatusProvisioningState">
BatchAccountPropertiesStatusProvisioningState
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioned state of the resource</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.PublicNetworkAccessType_Status">
PublicNetworkAccessType_Status
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: If not specified, the default value is &lsquo;enabled&rsquo;.</p>
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
<h3 id="batch.azure.com/v1beta20210101.BatchAccount_StatusARM">BatchAccount_StatusARM
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
<p>Id: The ID of the resource.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.BatchAccountIdentity_StatusARM">
BatchAccountIdentity_StatusARM
</a>
</em>
</td>
<td>
<p>Identity: The identity of the Batch account.</p>
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
<p>Location: The location of the resource.</p>
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
<a href="#batch.azure.com/v1beta20210101.BatchAccountProperties_StatusARM">
BatchAccountProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: The properties associated with the account.</p>
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
<h3 id="batch.azure.com/v1beta20210101.BatchAccountsSpecAPIVersion">BatchAccountsSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-01-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.BatchAccounts_Spec">BatchAccounts_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccount">BatchAccount</a>)
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
<code>autoStorage</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.AutoStorageBaseProperties">
AutoStorageBaseProperties
</a>
</em>
</td>
<td>
<p>AutoStorage: The properties related to the auto-storage account.</p>
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
<code>encryption</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.EncryptionProperties">
EncryptionProperties
</a>
</em>
</td>
<td>
<p>Encryption: Configures how customer data is encrypted inside the Batch account. By default, accounts are encrypted using
a Microsoft managed key. For additional control, a customer-managed key can be used instead.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.BatchAccountIdentity">
BatchAccountIdentity
</a>
</em>
</td>
<td>
<p>Identity: The identity of the Batch account, if configured. This is only used when the user specifies
&lsquo;Microsoft.KeyVault&rsquo; as their Batch account encryption configuration.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultReference</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.KeyVaultReference">
KeyVaultReference
</a>
</em>
</td>
<td>
<p>KeyVaultReference: Identifies the Azure key vault associated with a Batch account.</p>
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
<p>Location: The region in which to create the account.</p>
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
<code>poolAllocationMode</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.BatchAccountCreatePropertiesPoolAllocationMode">
BatchAccountCreatePropertiesPoolAllocationMode
</a>
</em>
</td>
<td>
<p>PoolAllocationMode: The pool allocation mode also affects how clients may authenticate to the Batch Service API. If the
mode is BatchService, clients may authenticate using access keys or Azure Active Directory. If the mode is
UserSubscription, clients must use Azure Active Directory. The default is BatchService.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.BatchAccountCreatePropertiesPublicNetworkAccess">
BatchAccountCreatePropertiesPublicNetworkAccess
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: If not specified, the default value is &lsquo;enabled&rsquo;.</p>
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
<p>Tags: The user-specified tags associated with the account.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.BatchAccounts_SpecARM">BatchAccounts_SpecARM
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
<a href="#batch.azure.com/v1beta20210101.BatchAccountIdentityARM">
BatchAccountIdentityARM
</a>
</em>
</td>
<td>
<p>Identity: The identity of the Batch account, if configured. This is only used when the user specifies
&lsquo;Microsoft.KeyVault&rsquo; as their Batch account encryption configuration.</p>
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
<p>Location: The region in which to create the account.</p>
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
<p>Name: A name for the Batch account which must be unique within the region. Batch account names must be between 3 and 24
characters in length and must use only numbers and lowercase letters. This name is used as part of the DNS name that is
used to access the Batch service in the region in which the account is created. For example:
<a href="http://accountname.region.batch.azure.com/">http://accountname.region.batch.azure.com/</a>.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.BatchAccountCreatePropertiesARM">
BatchAccountCreatePropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: The properties of a Batch account.</p>
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
<p>Tags: The user-specified tags associated with the account.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.EncryptionProperties">EncryptionProperties
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccounts_Spec">BatchAccounts_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/EncryptionProperties">https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/EncryptionProperties</a></p>
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
<code>keySource</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.EncryptionPropertiesKeySource">
EncryptionPropertiesKeySource
</a>
</em>
</td>
<td>
<p>KeySource: Type of the key source.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultProperties</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.KeyVaultProperties">
KeyVaultProperties
</a>
</em>
</td>
<td>
<p>KeyVaultProperties: KeyVault configuration when using an encryption KeySource of Microsoft.KeyVault.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.EncryptionPropertiesARM">EncryptionPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccountCreatePropertiesARM">BatchAccountCreatePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/EncryptionProperties">https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/EncryptionProperties</a></p>
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
<code>keySource</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.EncryptionPropertiesKeySource">
EncryptionPropertiesKeySource
</a>
</em>
</td>
<td>
<p>KeySource: Type of the key source.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultProperties</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.KeyVaultPropertiesARM">
KeyVaultPropertiesARM
</a>
</em>
</td>
<td>
<p>KeyVaultProperties: KeyVault configuration when using an encryption KeySource of Microsoft.KeyVault.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.EncryptionPropertiesKeySource">EncryptionPropertiesKeySource
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.EncryptionProperties">EncryptionProperties</a>, <a href="#batch.azure.com/v1beta20210101.EncryptionPropertiesARM">EncryptionPropertiesARM</a>)
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
<tbody><tr><td><p>&#34;Microsoft.Batch&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Microsoft.KeyVault&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.EncryptionPropertiesStatusKeySource">EncryptionPropertiesStatusKeySource
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.EncryptionProperties_Status">EncryptionProperties_Status</a>, <a href="#batch.azure.com/v1beta20210101.EncryptionProperties_StatusARM">EncryptionProperties_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Microsoft.Batch&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Microsoft.KeyVault&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.EncryptionProperties_Status">EncryptionProperties_Status
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccount_Status">BatchAccount_Status</a>)
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
<code>keySource</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.EncryptionPropertiesStatusKeySource">
EncryptionPropertiesStatusKeySource
</a>
</em>
</td>
<td>
<p>KeySource: Type of the key source.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultProperties</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.KeyVaultProperties_Status">
KeyVaultProperties_Status
</a>
</em>
</td>
<td>
<p>KeyVaultProperties: Additional details when using Microsoft.KeyVault</p>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.EncryptionProperties_StatusARM">EncryptionProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccountProperties_StatusARM">BatchAccountProperties_StatusARM</a>)
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
<code>keySource</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.EncryptionPropertiesStatusKeySource">
EncryptionPropertiesStatusKeySource
</a>
</em>
</td>
<td>
<p>KeySource: Type of the key source.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultProperties</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.KeyVaultProperties_StatusARM">
KeyVaultProperties_StatusARM
</a>
</em>
</td>
<td>
<p>KeyVaultProperties: Additional details when using Microsoft.KeyVault</p>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.KeyVaultProperties">KeyVaultProperties
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.EncryptionProperties">EncryptionProperties</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/KeyVaultProperties">https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/KeyVaultProperties</a></p>
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
<code>keyIdentifier</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyIdentifier: Full path to the versioned secret. Example
<a href="https://mykeyvault.vault.azure.net/keys/testkey/6e34a81fef704045975661e297a4c053">https://mykeyvault.vault.azure.net/keys/testkey/6e34a81fef704045975661e297a4c053</a>. To be usable the following
prerequisites must be met:
The Batch Account has a System Assigned identity
The account identity has been granted Key/Get, Key/Unwrap and Key/Wrap permissions
The KeyVault has soft-delete and purge protection enabled</p>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.KeyVaultPropertiesARM">KeyVaultPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.EncryptionPropertiesARM">EncryptionPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/KeyVaultProperties">https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/KeyVaultProperties</a></p>
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
<code>keyIdentifier</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyIdentifier: Full path to the versioned secret. Example
<a href="https://mykeyvault.vault.azure.net/keys/testkey/6e34a81fef704045975661e297a4c053">https://mykeyvault.vault.azure.net/keys/testkey/6e34a81fef704045975661e297a4c053</a>. To be usable the following
prerequisites must be met:
The Batch Account has a System Assigned identity
The account identity has been granted Key/Get, Key/Unwrap and Key/Wrap permissions
The KeyVault has soft-delete and purge protection enabled</p>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.KeyVaultProperties_Status">KeyVaultProperties_Status
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.EncryptionProperties_Status">EncryptionProperties_Status</a>)
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
<code>keyIdentifier</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyIdentifier: Full path to the versioned secret. Example
<a href="https://mykeyvault.vault.azure.net/keys/testkey/6e34a81fef704045975661e297a4c053">https://mykeyvault.vault.azure.net/keys/testkey/6e34a81fef704045975661e297a4c053</a>. To be usable the following
prerequisites must be met:
The Batch Account has a System Assigned identity
The account identity has been granted Key/Get, Key/Unwrap and Key/Wrap permissions
The KeyVault has soft-delete and purge protection enabled</p>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.KeyVaultProperties_StatusARM">KeyVaultProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.EncryptionProperties_StatusARM">EncryptionProperties_StatusARM</a>)
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
<code>keyIdentifier</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyIdentifier: Full path to the versioned secret. Example
<a href="https://mykeyvault.vault.azure.net/keys/testkey/6e34a81fef704045975661e297a4c053">https://mykeyvault.vault.azure.net/keys/testkey/6e34a81fef704045975661e297a4c053</a>. To be usable the following
prerequisites must be met:
The Batch Account has a System Assigned identity
The account identity has been granted Key/Get, Key/Unwrap and Key/Wrap permissions
The KeyVault has soft-delete and purge protection enabled</p>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.KeyVaultReference">KeyVaultReference
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccounts_Spec">BatchAccounts_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/KeyVaultReference">https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/KeyVaultReference</a></p>
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
<code>reference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>Reference: The resource ID of the Azure key vault associated with the Batch account.</p>
</td>
</tr>
<tr>
<td>
<code>url</code><br/>
<em>
string
</em>
</td>
<td>
<p>Url: The URL of the Azure key vault associated with the Batch account.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.KeyVaultReferenceARM">KeyVaultReferenceARM
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccountCreatePropertiesARM">BatchAccountCreatePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/KeyVaultReference">https://schema.management.azure.com/schemas/2021-01-01/Microsoft.Batch.json#/definitions/KeyVaultReference</a></p>
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
<code>url</code><br/>
<em>
string
</em>
</td>
<td>
<p>Url: The URL of the Azure key vault associated with the Batch account.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.KeyVaultReference_Status">KeyVaultReference_Status
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccount_Status">BatchAccount_Status</a>)
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
<p>Id: The resource ID of the Azure key vault associated with the Batch account.</p>
</td>
</tr>
<tr>
<td>
<code>url</code><br/>
<em>
string
</em>
</td>
<td>
<p>Url: The URL of the Azure key vault associated with the Batch account.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.KeyVaultReference_StatusARM">KeyVaultReference_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccountProperties_StatusARM">BatchAccountProperties_StatusARM</a>)
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
<p>Id: The resource ID of the Azure key vault associated with the Batch account.</p>
</td>
</tr>
<tr>
<td>
<code>url</code><br/>
<em>
string
</em>
</td>
<td>
<p>Url: The URL of the Azure key vault associated with the Batch account.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.PoolAllocationMode_Status">PoolAllocationMode_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccountProperties_StatusARM">BatchAccountProperties_StatusARM</a>, <a href="#batch.azure.com/v1beta20210101.BatchAccount_Status">BatchAccount_Status</a>)
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
<tbody><tr><td><p>&#34;BatchService&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UserSubscription&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.PrivateEndpointConnectionPropertiesStatusProvisioningState">PrivateEndpointConnectionPropertiesStatusProvisioningState
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.PrivateEndpointConnectionProperties_StatusARM">PrivateEndpointConnectionProperties_StatusARM</a>, <a href="#batch.azure.com/v1beta20210101.PrivateEndpointConnection_Status">PrivateEndpointConnection_Status</a>)
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
<tbody><tr><td><p>&#34;Failed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Succeeded&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Updating&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.PrivateEndpointConnectionProperties_StatusARM">PrivateEndpointConnectionProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.PrivateEndpointConnection_StatusARM">PrivateEndpointConnection_StatusARM</a>)
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
<a href="#batch.azure.com/v1beta20210101.PrivateEndpoint_StatusARM">
PrivateEndpoint_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>privateLinkServiceConnectionState</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.PrivateLinkServiceConnectionState_StatusARM">
PrivateLinkServiceConnectionState_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.PrivateEndpointConnectionPropertiesStatusProvisioningState">
PrivateEndpointConnectionPropertiesStatusProvisioningState
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.PrivateEndpointConnection_Status">PrivateEndpointConnection_Status
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccount_Status">BatchAccount_Status</a>)
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
<p>Etag: The ETag of the resource, used for concurrency statements.</p>
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
<p>Id: The ID of the resource.</p>
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
<code>privateEndpoint</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.PrivateEndpoint_Status">
PrivateEndpoint_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>privateLinkServiceConnectionState</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.PrivateLinkServiceConnectionState_Status">
PrivateLinkServiceConnectionState_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.PrivateEndpointConnectionPropertiesStatusProvisioningState">
PrivateEndpointConnectionPropertiesStatusProvisioningState
</a>
</em>
</td>
<td>
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
<h3 id="batch.azure.com/v1beta20210101.PrivateEndpointConnection_StatusARM">PrivateEndpointConnection_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccountProperties_StatusARM">BatchAccountProperties_StatusARM</a>)
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
<p>Etag: The ETag of the resource, used for concurrency statements.</p>
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
<p>Id: The ID of the resource.</p>
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
<a href="#batch.azure.com/v1beta20210101.PrivateEndpointConnectionProperties_StatusARM">
PrivateEndpointConnectionProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: The properties associated with the private endpoint connection.</p>
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
<h3 id="batch.azure.com/v1beta20210101.PrivateEndpoint_Status">PrivateEndpoint_Status
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.PrivateEndpointConnection_Status">PrivateEndpointConnection_Status</a>)
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
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.PrivateEndpoint_StatusARM">PrivateEndpoint_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.PrivateEndpointConnectionProperties_StatusARM">PrivateEndpointConnectionProperties_StatusARM</a>)
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
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.PrivateLinkServiceConnectionState_Status">PrivateLinkServiceConnectionState_Status
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.PrivateEndpointConnection_Status">PrivateEndpointConnection_Status</a>)
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
<code>actionRequired</code><br/>
<em>
string
</em>
</td>
<td>
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
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.PrivateLinkServiceConnectionStatus_Status">
PrivateLinkServiceConnectionStatus_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.PrivateLinkServiceConnectionState_StatusARM">PrivateLinkServiceConnectionState_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.PrivateEndpointConnectionProperties_StatusARM">PrivateEndpointConnectionProperties_StatusARM</a>)
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
<code>actionRequired</code><br/>
<em>
string
</em>
</td>
<td>
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
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#batch.azure.com/v1beta20210101.PrivateLinkServiceConnectionStatus_Status">
PrivateLinkServiceConnectionStatus_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.PrivateLinkServiceConnectionStatus_Status">PrivateLinkServiceConnectionStatus_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.PrivateLinkServiceConnectionState_Status">PrivateLinkServiceConnectionState_Status</a>, <a href="#batch.azure.com/v1beta20210101.PrivateLinkServiceConnectionState_StatusARM">PrivateLinkServiceConnectionState_StatusARM</a>)
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
<h3 id="batch.azure.com/v1beta20210101.PublicNetworkAccessType_Status">PublicNetworkAccessType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccountProperties_StatusARM">BatchAccountProperties_StatusARM</a>, <a href="#batch.azure.com/v1beta20210101.BatchAccount_Status">BatchAccount_Status</a>)
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
<h3 id="batch.azure.com/v1beta20210101.VirtualMachineFamilyCoreQuota_Status">VirtualMachineFamilyCoreQuota_Status
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccount_Status">BatchAccount_Status</a>)
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
<code>coreQuota</code><br/>
<em>
int
</em>
</td>
<td>
<p>CoreQuota: The core quota for the VM family for the Batch account.</p>
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
<p>Name: The Virtual Machine family name.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="batch.azure.com/v1beta20210101.VirtualMachineFamilyCoreQuota_StatusARM">VirtualMachineFamilyCoreQuota_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#batch.azure.com/v1beta20210101.BatchAccountProperties_StatusARM">BatchAccountProperties_StatusARM</a>)
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
<code>coreQuota</code><br/>
<em>
int
</em>
</td>
<td>
<p>CoreQuota: The core quota for the VM family for the Batch account.</p>
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
<p>Name: The Virtual Machine family name.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
