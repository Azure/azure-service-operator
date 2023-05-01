---
title: managedidentity.azure.com/v1api20220131preview
---
<h2 id="managedidentity.azure.com/v1api20220131preview">managedidentity.azure.com/v1api20220131preview</h2>
<div>
<p>Package v1api20220131preview contains API Schema definitions for the managedidentity v1api20220131preview API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="managedidentity.azure.com/v1api20220131preview.APIVersion">APIVersion
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
<tbody><tr><td><p>&#34;2022-01-31-preview&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="managedidentity.azure.com/v1api20220131preview.FederatedIdentityCredential">FederatedIdentityCredential
</h3>
<div>
<p>Generator information:
- Generated from: /msi/resource-manager/Microsoft.ManagedIdentity/preview/2022-01-31-preview/ManagedIdentity.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ManagedIdentity/&#x200b;userAssignedIdentities/&#x200b;{resourceName}/&#x200b;federatedIdentityCredentials/&#x200b;{federatedIdentityCredentialResourceName}</&#x200b;p>
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
<a href="#managedidentity.azure.com/v1api20220131preview.UserAssignedIdentities_FederatedIdentityCredential_Spec">
UserAssignedIdentities_FederatedIdentityCredential_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>audiences</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Audiences: The list of audiences that can appear in the issued token.</p>
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
<code>issuer</code><br/>
<em>
string
</em>
</td>
<td>
<p>Issuer: The URL of the issuer to be trusted.</p>
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
reference to a managedidentity.azure.com/UserAssignedIdentity resource</p>
</td>
</tr>
<tr>
<td>
<code>subject</code><br/>
<em>
string
</em>
</td>
<td>
<p>Subject: The identifier of the external identity.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#managedidentity.azure.com/v1api20220131preview.UserAssignedIdentities_FederatedIdentityCredential_STATUS">
UserAssignedIdentities_FederatedIdentityCredential_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="managedidentity.azure.com/v1api20220131preview.FederatedIdentityCredentialProperties_ARM">FederatedIdentityCredentialProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#managedidentity.azure.com/v1api20220131preview.UserAssignedIdentities_FederatedIdentityCredential_Spec_ARM">UserAssignedIdentities_FederatedIdentityCredential_Spec_ARM</a>)
</p>
<div>
<p>The properties associated with a federated identity credential.</p>
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
<code>audiences</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Audiences: The list of audiences that can appear in the issued token.</p>
</td>
</tr>
<tr>
<td>
<code>issuer</code><br/>
<em>
string
</em>
</td>
<td>
<p>Issuer: The URL of the issuer to be trusted.</p>
</td>
</tr>
<tr>
<td>
<code>subject</code><br/>
<em>
string
</em>
</td>
<td>
<p>Subject: The identifier of the external identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="managedidentity.azure.com/v1api20220131preview.FederatedIdentityCredentialProperties_STATUS_ARM">FederatedIdentityCredentialProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#managedidentity.azure.com/v1api20220131preview.UserAssignedIdentities_FederatedIdentityCredential_STATUS_ARM">UserAssignedIdentities_FederatedIdentityCredential_STATUS_ARM</a>)
</p>
<div>
<p>The properties associated with a federated identity credential.</p>
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
<code>audiences</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Audiences: The list of audiences that can appear in the issued token.</p>
</td>
</tr>
<tr>
<td>
<code>issuer</code><br/>
<em>
string
</em>
</td>
<td>
<p>Issuer: The URL of the issuer to be trusted.</p>
</td>
</tr>
<tr>
<td>
<code>subject</code><br/>
<em>
string
</em>
</td>
<td>
<p>Subject: The identifier of the external identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="managedidentity.azure.com/v1api20220131preview.UserAssignedIdentities_FederatedIdentityCredential_STATUS">UserAssignedIdentities_FederatedIdentityCredential_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#managedidentity.azure.com/v1api20220131preview.FederatedIdentityCredential">FederatedIdentityCredential</a>)
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
<code>audiences</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Audiences: The list of audiences that can appear in the issued token.</p>
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
<code>issuer</code><br/>
<em>
string
</em>
</td>
<td>
<p>Issuer: The URL of the issuer to be trusted.</p>
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
<code>subject</code><br/>
<em>
string
</em>
</td>
<td>
<p>Subject: The identifier of the external identity.</p>
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
<h3 id="managedidentity.azure.com/v1api20220131preview.UserAssignedIdentities_FederatedIdentityCredential_STATUS_ARM">UserAssignedIdentities_FederatedIdentityCredential_STATUS_ARM
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
<a href="#managedidentity.azure.com/v1api20220131preview.FederatedIdentityCredentialProperties_STATUS_ARM">
FederatedIdentityCredentialProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: The properties associated with the federated identity credential.</p>
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
<h3 id="managedidentity.azure.com/v1api20220131preview.UserAssignedIdentities_FederatedIdentityCredential_Spec">UserAssignedIdentities_FederatedIdentityCredential_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#managedidentity.azure.com/v1api20220131preview.FederatedIdentityCredential">FederatedIdentityCredential</a>)
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
<code>audiences</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Audiences: The list of audiences that can appear in the issued token.</p>
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
<code>issuer</code><br/>
<em>
string
</em>
</td>
<td>
<p>Issuer: The URL of the issuer to be trusted.</p>
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
reference to a managedidentity.azure.com/UserAssignedIdentity resource</p>
</td>
</tr>
<tr>
<td>
<code>subject</code><br/>
<em>
string
</em>
</td>
<td>
<p>Subject: The identifier of the external identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="managedidentity.azure.com/v1api20220131preview.UserAssignedIdentities_FederatedIdentityCredential_Spec_ARM">UserAssignedIdentities_FederatedIdentityCredential_Spec_ARM
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
<a href="#managedidentity.azure.com/v1api20220131preview.FederatedIdentityCredentialProperties_ARM">
FederatedIdentityCredentialProperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: The properties associated with the federated identity credential.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
