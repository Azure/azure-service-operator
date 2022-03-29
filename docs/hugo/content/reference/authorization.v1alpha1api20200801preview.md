<h2 id="authorization.azure.com/v1alpha1api20200801preview">authorization.azure.com/v1alpha1api20200801preview</h2>
<div>
<p>Package v1alpha1api20200801preview contains API Schema definitions for the authorization v1alpha1api20200801preview API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="authorization.azure.com/v1alpha1api20200801preview.RoleAssignment">RoleAssignment
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-08-01-preview/Microsoft.Authorization.Authz.json#/unknown_resourceDefinitions/roleAssignments">https://schema.management.azure.com/schemas/2020-08-01-preview/Microsoft.Authorization.Authz.json#/unknown_resourceDefinitions/roleAssignments</a></p>
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
<a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignments_Spec">
RoleAssignments_Spec
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
<code>condition</code><br/>
<em>
string
</em>
</td>
<td>
<p>Condition: The conditions on the role assignment. This limits the resources it can be assigned to. e.g.:
@Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase
&lsquo;foo_storage_container&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>conditionVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>ConditionVersion: Version of the condition. Currently accepted value is &lsquo;2.0&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>delegatedManagedIdentityResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>DelegatedManagedIdentityResourceId: Id of the delegated managed identity resource</p>
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
<p>Description: Description of role assignment</p>
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
<code>principalId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrincipalId: The principal ID.</p>
</td>
</tr>
<tr>
<td>
<code>principalType</code><br/>
<em>
<a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentPropertiesPrincipalType">
RoleAssignmentPropertiesPrincipalType
</a>
</em>
</td>
<td>
<p>PrincipalType: The principal type of the assigned principal ID.</p>
</td>
</tr>
<tr>
<td>
<code>roleDefinitionReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>RoleDefinitionReference: The role definition ID.</p>
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
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignment_Status">
RoleAssignment_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentPropertiesARM">RoleAssignmentPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignments_SpecARM">RoleAssignments_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-08-01-preview/Microsoft.Authorization.Authz.json#/definitions/RoleAssignmentProperties">https://schema.management.azure.com/schemas/2020-08-01-preview/Microsoft.Authorization.Authz.json#/definitions/RoleAssignmentProperties</a></p>
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
<code>condition</code><br/>
<em>
string
</em>
</td>
<td>
<p>Condition: The conditions on the role assignment. This limits the resources it can be assigned to. e.g.:
@Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase
&lsquo;foo_storage_container&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>conditionVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>ConditionVersion: Version of the condition. Currently accepted value is &lsquo;2.0&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>delegatedManagedIdentityResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>DelegatedManagedIdentityResourceId: Id of the delegated managed identity resource</p>
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
<p>Description: Description of role assignment</p>
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
<p>PrincipalId: The principal ID.</p>
</td>
</tr>
<tr>
<td>
<code>principalType</code><br/>
<em>
<a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentPropertiesPrincipalType">
RoleAssignmentPropertiesPrincipalType
</a>
</em>
</td>
<td>
<p>PrincipalType: The principal type of the assigned principal ID.</p>
</td>
</tr>
<tr>
<td>
<code>roleDefinitionId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentPropertiesPrincipalType">RoleAssignmentPropertiesPrincipalType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentPropertiesARM">RoleAssignmentPropertiesARM</a>, <a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignments_Spec">RoleAssignments_Spec</a>)
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
<tbody><tr><td><p>&#34;ForeignGroup&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Group&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ServicePrincipal&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;User&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentPropertiesStatusPrincipalType">RoleAssignmentPropertiesStatusPrincipalType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentProperties_StatusARM">RoleAssignmentProperties_StatusARM</a>, <a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignment_Status">RoleAssignment_Status</a>)
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
<tbody><tr><td><p>&#34;ForeignGroup&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Group&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ServicePrincipal&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;User&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentProperties_StatusARM">RoleAssignmentProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignment_StatusARM">RoleAssignment_StatusARM</a>)
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
<code>condition</code><br/>
<em>
string
</em>
</td>
<td>
<p>Condition: The conditions on the role assignment. This limits the resources it can be assigned to. e.g.:
@Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase
&lsquo;foo_storage_container&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>conditionVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>ConditionVersion: Version of the condition. Currently accepted value is &lsquo;2.0&rsquo;</p>
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
<p>CreatedBy: Id of the user who created the assignment</p>
</td>
</tr>
<tr>
<td>
<code>createdOn</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedOn: Time it was created</p>
</td>
</tr>
<tr>
<td>
<code>delegatedManagedIdentityResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>DelegatedManagedIdentityResourceId: Id of the delegated managed identity resource</p>
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
<p>Description: Description of role assignment</p>
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
<p>PrincipalId: The principal ID.</p>
</td>
</tr>
<tr>
<td>
<code>principalType</code><br/>
<em>
<a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentPropertiesStatusPrincipalType">
RoleAssignmentPropertiesStatusPrincipalType
</a>
</em>
</td>
<td>
<p>PrincipalType: The principal type of the assigned principal ID.</p>
</td>
</tr>
<tr>
<td>
<code>roleDefinitionId</code><br/>
<em>
string
</em>
</td>
<td>
<p>RoleDefinitionId: The role definition ID.</p>
</td>
</tr>
<tr>
<td>
<code>scope</code><br/>
<em>
string
</em>
</td>
<td>
<p>Scope: The role assignment scope.</p>
</td>
</tr>
<tr>
<td>
<code>updatedBy</code><br/>
<em>
string
</em>
</td>
<td>
<p>UpdatedBy: Id of the user who updated the assignment</p>
</td>
</tr>
<tr>
<td>
<code>updatedOn</code><br/>
<em>
string
</em>
</td>
<td>
<p>UpdatedOn: Time it was updated</p>
</td>
</tr>
</tbody>
</table>
<h3 id="authorization.azure.com/v1alpha1api20200801preview.RoleAssignment_Status">RoleAssignment_Status
</h3>
<p>
(<em>Appears on:</em><a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignment">RoleAssignment</a>)
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
<code>condition</code><br/>
<em>
string
</em>
</td>
<td>
<p>Condition: The conditions on the role assignment. This limits the resources it can be assigned to. e.g.:
@Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase
&lsquo;foo_storage_container&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>conditionVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>ConditionVersion: Version of the condition. Currently accepted value is &lsquo;2.0&rsquo;</p>
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
<code>createdBy</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedBy: Id of the user who created the assignment</p>
</td>
</tr>
<tr>
<td>
<code>createdOn</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedOn: Time it was created</p>
</td>
</tr>
<tr>
<td>
<code>delegatedManagedIdentityResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>DelegatedManagedIdentityResourceId: Id of the delegated managed identity resource</p>
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
<p>Description: Description of role assignment</p>
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
<p>Id: The role assignment ID.</p>
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
<p>Name: The role assignment name.</p>
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
<p>PrincipalId: The principal ID.</p>
</td>
</tr>
<tr>
<td>
<code>principalType</code><br/>
<em>
<a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentPropertiesStatusPrincipalType">
RoleAssignmentPropertiesStatusPrincipalType
</a>
</em>
</td>
<td>
<p>PrincipalType: The principal type of the assigned principal ID.</p>
</td>
</tr>
<tr>
<td>
<code>roleDefinitionId</code><br/>
<em>
string
</em>
</td>
<td>
<p>RoleDefinitionId: The role definition ID.</p>
</td>
</tr>
<tr>
<td>
<code>scope</code><br/>
<em>
string
</em>
</td>
<td>
<p>Scope: The role assignment scope.</p>
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
<p>Type: The role assignment type.</p>
</td>
</tr>
<tr>
<td>
<code>updatedBy</code><br/>
<em>
string
</em>
</td>
<td>
<p>UpdatedBy: Id of the user who updated the assignment</p>
</td>
</tr>
<tr>
<td>
<code>updatedOn</code><br/>
<em>
string
</em>
</td>
<td>
<p>UpdatedOn: Time it was updated</p>
</td>
</tr>
</tbody>
</table>
<h3 id="authorization.azure.com/v1alpha1api20200801preview.RoleAssignment_StatusARM">RoleAssignment_StatusARM
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
<p>Id: The role assignment ID.</p>
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
<p>Name: The role assignment name.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentProperties_StatusARM">
RoleAssignmentProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Role assignment properties.</p>
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
<p>Type: The role assignment type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentsSpecAPIVersion">RoleAssignmentsSpecAPIVersion
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
<tbody><tr><td><p>&#34;2020-08-01-preview&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="authorization.azure.com/v1alpha1api20200801preview.RoleAssignments_Spec">RoleAssignments_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignment">RoleAssignment</a>)
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
<code>condition</code><br/>
<em>
string
</em>
</td>
<td>
<p>Condition: The conditions on the role assignment. This limits the resources it can be assigned to. e.g.:
@Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase
&lsquo;foo_storage_container&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>conditionVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>ConditionVersion: Version of the condition. Currently accepted value is &lsquo;2.0&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>delegatedManagedIdentityResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>DelegatedManagedIdentityResourceId: Id of the delegated managed identity resource</p>
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
<p>Description: Description of role assignment</p>
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
<code>principalId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrincipalId: The principal ID.</p>
</td>
</tr>
<tr>
<td>
<code>principalType</code><br/>
<em>
<a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentPropertiesPrincipalType">
RoleAssignmentPropertiesPrincipalType
</a>
</em>
</td>
<td>
<p>PrincipalType: The principal type of the assigned principal ID.</p>
</td>
</tr>
<tr>
<td>
<code>roleDefinitionReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>RoleDefinitionReference: The role definition ID.</p>
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
<h3 id="authorization.azure.com/v1alpha1api20200801preview.RoleAssignments_SpecARM">RoleAssignments_SpecARM
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
<p>Name: The name of the role assignment. It can be any valid GUID.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#authorization.azure.com/v1alpha1api20200801preview.RoleAssignmentPropertiesARM">
RoleAssignmentPropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Role assignment properties.</p>
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
<hr/>
