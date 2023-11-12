---
title: containerservice.azure.com/v1api20230315preview
---
<h2 id="containerservice.azure.com/v1api20230315preview">containerservice.azure.com/v1api20230315preview</h2>
<div>
<p>Package v1api20230315preview contains API Schema definitions for the containerservice v1api20230315preview API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="containerservice.azure.com/v1api20230315preview.APIVersion">APIVersion
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
<tbody><tr><td><p>&#34;2023-03-15-preview&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.ErrorAdditionalInfo_STATUS">ErrorAdditionalInfo_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.ErrorDetail_STATUS">ErrorDetail_STATUS</a>, <a href="#containerservice.azure.com/v1api20230315preview.ErrorDetail_STATUS_Unrolled">ErrorDetail_STATUS_Unrolled</a>)
</p>
<div>
<p>The resource management error additional info.</p>
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
<code>info</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
map[string]k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1.JSON
</a>
</em>
</td>
<td>
<p>Info: The additional info.</p>
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
<p>Type: The additional info type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.ErrorAdditionalInfo_STATUS_ARM">ErrorAdditionalInfo_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.ErrorDetail_STATUS_ARM">ErrorDetail_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230315preview.ErrorDetail_STATUS_Unrolled_ARM">ErrorDetail_STATUS_Unrolled_ARM</a>)
</p>
<div>
<p>The resource management error additional info.</p>
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
<code>info</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
map[string]k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1.JSON
</a>
</em>
</td>
<td>
<p>Info: The additional info.</p>
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
<p>Type: The additional info type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.ErrorDetail_STATUS">ErrorDetail_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.UpdateStatus_STATUS">UpdateStatus_STATUS</a>)
</p>
<div>
<p>The error detail.</p>
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
<code>additionalInfo</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.ErrorAdditionalInfo_STATUS">
[]ErrorAdditionalInfo_STATUS
</a>
</em>
</td>
<td>
<p>AdditionalInfo: The error additional info.</p>
</td>
</tr>
<tr>
<td>
<code>code</code><br/>
<em>
string
</em>
</td>
<td>
<p>Code: The error code.</p>
</td>
</tr>
<tr>
<td>
<code>details</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.ErrorDetail_STATUS_Unrolled">
[]ErrorDetail_STATUS_Unrolled
</a>
</em>
</td>
<td>
<p>Details: The error details.</p>
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
<p>Message: The error message.</p>
</td>
</tr>
<tr>
<td>
<code>target</code><br/>
<em>
string
</em>
</td>
<td>
<p>Target: The error target.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.ErrorDetail_STATUS_ARM">ErrorDetail_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.UpdateStatus_STATUS_ARM">UpdateStatus_STATUS_ARM</a>)
</p>
<div>
<p>The error detail.</p>
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
<code>additionalInfo</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.ErrorAdditionalInfo_STATUS_ARM">
[]ErrorAdditionalInfo_STATUS_ARM
</a>
</em>
</td>
<td>
<p>AdditionalInfo: The error additional info.</p>
</td>
</tr>
<tr>
<td>
<code>code</code><br/>
<em>
string
</em>
</td>
<td>
<p>Code: The error code.</p>
</td>
</tr>
<tr>
<td>
<code>details</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.ErrorDetail_STATUS_Unrolled_ARM">
[]ErrorDetail_STATUS_Unrolled_ARM
</a>
</em>
</td>
<td>
<p>Details: The error details.</p>
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
<p>Message: The error message.</p>
</td>
</tr>
<tr>
<td>
<code>target</code><br/>
<em>
string
</em>
</td>
<td>
<p>Target: The error target.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.ErrorDetail_STATUS_Unrolled">ErrorDetail_STATUS_Unrolled
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.ErrorDetail_STATUS">ErrorDetail_STATUS</a>)
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
<code>additionalInfo</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.ErrorAdditionalInfo_STATUS">
[]ErrorAdditionalInfo_STATUS
</a>
</em>
</td>
<td>
<p>AdditionalInfo: The error additional info.</p>
</td>
</tr>
<tr>
<td>
<code>code</code><br/>
<em>
string
</em>
</td>
<td>
<p>Code: The error code.</p>
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
<p>Message: The error message.</p>
</td>
</tr>
<tr>
<td>
<code>target</code><br/>
<em>
string
</em>
</td>
<td>
<p>Target: The error target.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.ErrorDetail_STATUS_Unrolled_ARM">ErrorDetail_STATUS_Unrolled_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.ErrorDetail_STATUS_ARM">ErrorDetail_STATUS_ARM</a>)
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
<code>additionalInfo</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.ErrorAdditionalInfo_STATUS_ARM">
[]ErrorAdditionalInfo_STATUS_ARM
</a>
</em>
</td>
<td>
<p>AdditionalInfo: The error additional info.</p>
</td>
</tr>
<tr>
<td>
<code>code</code><br/>
<em>
string
</em>
</td>
<td>
<p>Code: The error code.</p>
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
<p>Message: The error message.</p>
</td>
</tr>
<tr>
<td>
<code>target</code><br/>
<em>
string
</em>
</td>
<td>
<p>Target: The error target.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.Fleet">Fleet
</h3>
<div>
<p>Generator information:
- Generated from: /containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-03-15-preview/fleets.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ContainerService/&#x200b;fleets/&#x200b;{fleetName}</&#x200b;p>
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
<a href="#containerservice.azure.com/v1api20230315preview.Fleet_Spec">
Fleet_Spec
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
<code>hubProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.FleetHubProfile">
FleetHubProfile
</a>
</em>
</td>
<td>
<p>HubProfile: The FleetHubProfile configures the Fleet&rsquo;s hub.</p>
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
<p>Location: The geo-location where the resource lives</p>
</td>
</tr>
<tr>
<td>
<code>operatorSpec</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.FleetOperatorSpec">
FleetOperatorSpec
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
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Resource tags.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.Fleet_STATUS">
Fleet_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.FleetHubProfile">FleetHubProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.Fleet_Spec">Fleet_Spec</a>)
</p>
<div>
<p>The FleetHubProfile configures the fleet hub.</p>
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
<code>dnsPrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>DnsPrefix: DNS prefix used to create the FQDN for the Fleet hub.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.FleetHubProfile_ARM">FleetHubProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.FleetProperties_ARM">FleetProperties_ARM</a>)
</p>
<div>
<p>The FleetHubProfile configures the fleet hub.</p>
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
<code>dnsPrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>DnsPrefix: DNS prefix used to create the FQDN for the Fleet hub.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.FleetHubProfile_STATUS">FleetHubProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.Fleet_STATUS">Fleet_STATUS</a>)
</p>
<div>
<p>The FleetHubProfile configures the fleet hub.</p>
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
<code>dnsPrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>DnsPrefix: DNS prefix used to create the FQDN for the Fleet hub.</p>
</td>
</tr>
<tr>
<td>
<code>fqdn</code><br/>
<em>
string
</em>
</td>
<td>
<p>Fqdn: The FQDN of the Fleet hub.</p>
</td>
</tr>
<tr>
<td>
<code>kubernetesVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>KubernetesVersion: The Kubernetes version of the Fleet hub.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.FleetHubProfile_STATUS_ARM">FleetHubProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.FleetProperties_STATUS_ARM">FleetProperties_STATUS_ARM</a>)
</p>
<div>
<p>The FleetHubProfile configures the fleet hub.</p>
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
<code>dnsPrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>DnsPrefix: DNS prefix used to create the FQDN for the Fleet hub.</p>
</td>
</tr>
<tr>
<td>
<code>fqdn</code><br/>
<em>
string
</em>
</td>
<td>
<p>Fqdn: The FQDN of the Fleet hub.</p>
</td>
</tr>
<tr>
<td>
<code>kubernetesVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>KubernetesVersion: The Kubernetes version of the Fleet hub.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.FleetMemberProperties_ARM">FleetMemberProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.Fleets_Member_Spec_ARM">Fleets_Member_Spec_ARM</a>)
</p>
<div>
<p>A member of the Fleet. It contains a reference to an existing Kubernetes cluster on Azure.</p>
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
<code>clusterResourceId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>group</code><br/>
<em>
string
</em>
</td>
<td>
<p>Group: The group this member belongs to for multi-cluster update management.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.FleetMemberProperties_STATUS_ARM">FleetMemberProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.Fleets_Member_STATUS_ARM">Fleets_Member_STATUS_ARM</a>)
</p>
<div>
<p>A member of the Fleet. It contains a reference to an existing Kubernetes cluster on Azure.</p>
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
<code>clusterResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClusterResourceId: The ARM resource id of the cluster that joins the Fleet. Must be a valid Azure resource id. e.g.:
&lsquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ContainerService/&#x200b;managedClusters/&#x200b;{clusterName}&rsquo;.</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>group</code><br/>
<em>
string
</em>
</td>
<td>
<p>Group: The group this member belongs to for multi-cluster update management.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.FleetMemberProvisioningState_STATUS">
FleetMemberProvisioningState_STATUS
</a>
</em>
</td>
<td>
<p>ProvisioningState: The status of the last operation.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.FleetMemberProvisioningState_STATUS">FleetMemberProvisioningState_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.FleetMemberProperties_STATUS_ARM">FleetMemberProperties_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230315preview.Fleets_Member_STATUS">Fleets_Member_STATUS</a>)
</p>
<div>
<p>The provisioning state of the last accepted operation.</p>
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
</tr><tr><td><p>&#34;Failed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Joining&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Leaving&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Succeeded&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Updating&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.FleetOperatorSecrets">FleetOperatorSecrets
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.FleetOperatorSpec">FleetOperatorSpec</a>)
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
<code>userCredentials</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretDestination">
genruntime.SecretDestination
</a>
</em>
</td>
<td>
<p>UserCredentials: indicates where the UserCredentials secret should be placed. If omitted, the secret will not be
retrieved from Azure.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.FleetOperatorSpec">FleetOperatorSpec
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.Fleet_Spec">Fleet_Spec</a>)
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
<code>secrets</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.FleetOperatorSecrets">
FleetOperatorSecrets
</a>
</em>
</td>
<td>
<p>Secrets: configures where to place Azure generated secrets.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.FleetProperties_ARM">FleetProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.Fleet_Spec_ARM">Fleet_Spec_ARM</a>)
</p>
<div>
<p>Fleet properties.</p>
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
<code>hubProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.FleetHubProfile_ARM">
FleetHubProfile_ARM
</a>
</em>
</td>
<td>
<p>HubProfile: The FleetHubProfile configures the Fleet&rsquo;s hub.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.FleetProperties_STATUS_ARM">FleetProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.Fleet_STATUS_ARM">Fleet_STATUS_ARM</a>)
</p>
<div>
<p>Fleet properties.</p>
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
<code>hubProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.FleetHubProfile_STATUS_ARM">
FleetHubProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>HubProfile: The FleetHubProfile configures the Fleet&rsquo;s hub.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.FleetProvisioningState_STATUS">
FleetProvisioningState_STATUS
</a>
</em>
</td>
<td>
<p>ProvisioningState: The status of the last operation.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.FleetProvisioningState_STATUS">FleetProvisioningState_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.FleetProperties_STATUS_ARM">FleetProperties_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230315preview.Fleet_STATUS">Fleet_STATUS</a>)
</p>
<div>
<p>The provisioning state of the last accepted operation.</p>
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
<h3 id="containerservice.azure.com/v1api20230315preview.Fleet_STATUS">Fleet_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.Fleet">Fleet</a>)
</p>
<div>
<p>The Fleet resource.</p>
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
<code>eTag</code><br/>
<em>
string
</em>
</td>
<td>
<p>ETag: If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.
Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in
the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header
fields.</p>
</td>
</tr>
<tr>
<td>
<code>hubProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.FleetHubProfile_STATUS">
FleetHubProfile_STATUS
</a>
</em>
</td>
<td>
<p>HubProfile: The FleetHubProfile configures the Fleet&rsquo;s hub.</p>
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
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: The geo-location where the resource lives</p>
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
<code>provisioningState</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.FleetProvisioningState_STATUS">
FleetProvisioningState_STATUS
</a>
</em>
</td>
<td>
<p>ProvisioningState: The status of the last operation.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.SystemData_STATUS">
SystemData_STATUS
</a>
</em>
</td>
<td>
<p>SystemData: Azure Resource Manager metadata containing createdBy and modifiedBy information.</p>
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
<p>Tags: Resource tags.</p>
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
<h3 id="containerservice.azure.com/v1api20230315preview.Fleet_STATUS_ARM">Fleet_STATUS_ARM
</h3>
<div>
<p>The Fleet resource.</p>
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
<code>eTag</code><br/>
<em>
string
</em>
</td>
<td>
<p>ETag: If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.
Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in
the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header
fields.</p>
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
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: The geo-location where the resource lives</p>
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
<a href="#containerservice.azure.com/v1api20230315preview.FleetProperties_STATUS_ARM">
FleetProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: The resource-specific properties for this resource.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.SystemData_STATUS_ARM">
SystemData_STATUS_ARM
</a>
</em>
</td>
<td>
<p>SystemData: Azure Resource Manager metadata containing createdBy and modifiedBy information.</p>
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
<p>Tags: Resource tags.</p>
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
<h3 id="containerservice.azure.com/v1api20230315preview.Fleet_Spec">Fleet_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.Fleet">Fleet</a>)
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
<code>hubProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.FleetHubProfile">
FleetHubProfile
</a>
</em>
</td>
<td>
<p>HubProfile: The FleetHubProfile configures the Fleet&rsquo;s hub.</p>
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
<p>Location: The geo-location where the resource lives</p>
</td>
</tr>
<tr>
<td>
<code>operatorSpec</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.FleetOperatorSpec">
FleetOperatorSpec
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
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Resource tags.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.Fleet_Spec_ARM">Fleet_Spec_ARM
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
<p>Location: The geo-location where the resource lives</p>
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
<a href="#containerservice.azure.com/v1api20230315preview.FleetProperties_ARM">
FleetProperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: The resource-specific properties for this resource.</p>
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
<p>Tags: Resource tags.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.FleetsMember">FleetsMember
</h3>
<div>
<p>Generator information:
- Generated from: /containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-03-15-preview/fleets.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ContainerService/&#x200b;fleets/&#x200b;{fleetName}/&#x200b;members/&#x200b;{fleetMemberName}</&#x200b;p>
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
<a href="#containerservice.azure.com/v1api20230315preview.Fleets_Member_Spec">
Fleets_Member_Spec
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
<code>clusterResourceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>ClusterResourceReference: The ARM resource id of the cluster that joins the Fleet. Must be a valid Azure resource id.
e.g.:
&lsquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ContainerService/&#x200b;managedClusters/&#x200b;{clusterName}&rsquo;.</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>group</code><br/>
<em>
string
</em>
</td>
<td>
<p>Group: The group this member belongs to for multi-cluster update management.</p>
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
reference to a containerservice.azure.com/Fleet resource</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.Fleets_Member_STATUS">
Fleets_Member_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.FleetsUpdateRun">FleetsUpdateRun
</h3>
<div>
<p>Generator information:
- Generated from: /containerservice/resource-manager/Microsoft.ContainerService/fleet/preview/2023-03-15-preview/fleets.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ContainerService/&#x200b;fleets/&#x200b;{fleetName}/&#x200b;updateRuns/&#x200b;{updateRunName}</&#x200b;p>
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
<a href="#containerservice.azure.com/v1api20230315preview.Fleets_UpdateRun_Spec">
Fleets_UpdateRun_Spec
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
<code>managedClusterUpdate</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.ManagedClusterUpdate">
ManagedClusterUpdate
</a>
</em>
</td>
<td>
<p>ManagedClusterUpdate: The update to be applied to all clusters in the UpdateRun. The managedClusterUpdate can be
modified until the run is started.</p>
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
reference to a containerservice.azure.com/Fleet resource</p>
</td>
</tr>
<tr>
<td>
<code>strategy</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateRunStrategy">
UpdateRunStrategy
</a>
</em>
</td>
<td>
<p>Strategy: The strategy defines the order in which the clusters will be updated.
If not set, all members will be updated sequentially. The UpdateRun status will show a single UpdateStage and a single
UpdateGroup targeting all members.
The strategy of the UpdateRun can be modified until the run is started.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.Fleets_UpdateRun_STATUS">
Fleets_UpdateRun_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.Fleets_Member_STATUS">Fleets_Member_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.FleetsMember">FleetsMember</a>)
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
<code>clusterResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClusterResourceId: The ARM resource id of the cluster that joins the Fleet. Must be a valid Azure resource id. e.g.:
&lsquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ContainerService/&#x200b;managedClusters/&#x200b;{clusterName}&rsquo;.</&#x200b;p>
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
<code>eTag</code><br/>
<em>
string
</em>
</td>
<td>
<p>ETag: If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.
Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in
the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header
fields.</p>
</td>
</tr>
<tr>
<td>
<code>group</code><br/>
<em>
string
</em>
</td>
<td>
<p>Group: The group this member belongs to for multi-cluster update management.</p>
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
<code>provisioningState</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.FleetMemberProvisioningState_STATUS">
FleetMemberProvisioningState_STATUS
</a>
</em>
</td>
<td>
<p>ProvisioningState: The status of the last operation.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.SystemData_STATUS">
SystemData_STATUS
</a>
</em>
</td>
<td>
<p>SystemData: Azure Resource Manager metadata containing createdBy and modifiedBy information.</p>
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
<h3 id="containerservice.azure.com/v1api20230315preview.Fleets_Member_STATUS_ARM">Fleets_Member_STATUS_ARM
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
<code>eTag</code><br/>
<em>
string
</em>
</td>
<td>
<p>ETag: If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.
Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in
the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header
fields.</p>
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
<a href="#containerservice.azure.com/v1api20230315preview.FleetMemberProperties_STATUS_ARM">
FleetMemberProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: The resource-specific properties for this resource.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.SystemData_STATUS_ARM">
SystemData_STATUS_ARM
</a>
</em>
</td>
<td>
<p>SystemData: Azure Resource Manager metadata containing createdBy and modifiedBy information.</p>
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
<h3 id="containerservice.azure.com/v1api20230315preview.Fleets_Member_Spec">Fleets_Member_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.FleetsMember">FleetsMember</a>)
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
<code>clusterResourceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>ClusterResourceReference: The ARM resource id of the cluster that joins the Fleet. Must be a valid Azure resource id.
e.g.:
&lsquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ContainerService/&#x200b;managedClusters/&#x200b;{clusterName}&rsquo;.</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>group</code><br/>
<em>
string
</em>
</td>
<td>
<p>Group: The group this member belongs to for multi-cluster update management.</p>
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
reference to a containerservice.azure.com/Fleet resource</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.Fleets_Member_Spec_ARM">Fleets_Member_Spec_ARM
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
<a href="#containerservice.azure.com/v1api20230315preview.FleetMemberProperties_ARM">
FleetMemberProperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: The resource-specific properties for this resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.Fleets_UpdateRun_STATUS">Fleets_UpdateRun_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.FleetsUpdateRun">FleetsUpdateRun</a>)
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
<code>eTag</code><br/>
<em>
string
</em>
</td>
<td>
<p>ETag: If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.
Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in
the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header
fields.</p>
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
<code>managedClusterUpdate</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.ManagedClusterUpdate_STATUS">
ManagedClusterUpdate_STATUS
</a>
</em>
</td>
<td>
<p>ManagedClusterUpdate: The update to be applied to all clusters in the UpdateRun. The managedClusterUpdate can be
modified until the run is started.</p>
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
<code>provisioningState</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateRunProvisioningState_STATUS">
UpdateRunProvisioningState_STATUS
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the UpdateRun resource.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateRunStatus_STATUS">
UpdateRunStatus_STATUS
</a>
</em>
</td>
<td>
<p>Status: The status of the UpdateRun.</p>
</td>
</tr>
<tr>
<td>
<code>strategy</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateRunStrategy_STATUS">
UpdateRunStrategy_STATUS
</a>
</em>
</td>
<td>
<p>Strategy: The strategy defines the order in which the clusters will be updated.
If not set, all members will be updated sequentially. The UpdateRun status will show a single UpdateStage and a single
UpdateGroup targeting all members.
The strategy of the UpdateRun can be modified until the run is started.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.SystemData_STATUS">
SystemData_STATUS
</a>
</em>
</td>
<td>
<p>SystemData: Azure Resource Manager metadata containing createdBy and modifiedBy information.</p>
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
<h3 id="containerservice.azure.com/v1api20230315preview.Fleets_UpdateRun_STATUS_ARM">Fleets_UpdateRun_STATUS_ARM
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
<code>eTag</code><br/>
<em>
string
</em>
</td>
<td>
<p>ETag: If eTag is provided in the response body, it may also be provided as a header per the normal etag convention.
Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in
the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header
fields.</p>
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
<a href="#containerservice.azure.com/v1api20230315preview.UpdateRunProperties_STATUS_ARM">
UpdateRunProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: The resource-specific properties for this resource.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.SystemData_STATUS_ARM">
SystemData_STATUS_ARM
</a>
</em>
</td>
<td>
<p>SystemData: Azure Resource Manager metadata containing createdBy and modifiedBy information.</p>
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
<h3 id="containerservice.azure.com/v1api20230315preview.Fleets_UpdateRun_Spec">Fleets_UpdateRun_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.FleetsUpdateRun">FleetsUpdateRun</a>)
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
<code>managedClusterUpdate</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.ManagedClusterUpdate">
ManagedClusterUpdate
</a>
</em>
</td>
<td>
<p>ManagedClusterUpdate: The update to be applied to all clusters in the UpdateRun. The managedClusterUpdate can be
modified until the run is started.</p>
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
reference to a containerservice.azure.com/Fleet resource</p>
</td>
</tr>
<tr>
<td>
<code>strategy</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateRunStrategy">
UpdateRunStrategy
</a>
</em>
</td>
<td>
<p>Strategy: The strategy defines the order in which the clusters will be updated.
If not set, all members will be updated sequentially. The UpdateRun status will show a single UpdateStage and a single
UpdateGroup targeting all members.
The strategy of the UpdateRun can be modified until the run is started.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.Fleets_UpdateRun_Spec_ARM">Fleets_UpdateRun_Spec_ARM
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
<a href="#containerservice.azure.com/v1api20230315preview.UpdateRunProperties_ARM">
UpdateRunProperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: The resource-specific properties for this resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.ManagedClusterUpdate">ManagedClusterUpdate
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.Fleets_UpdateRun_Spec">Fleets_UpdateRun_Spec</a>)
</p>
<div>
<p>The update to be applied to the ManagedClusters.</p>
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
<code>upgrade</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.ManagedClusterUpgradeSpec">
ManagedClusterUpgradeSpec
</a>
</em>
</td>
<td>
<p>Upgrade: The upgrade to apply to the ManagedClusters.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.ManagedClusterUpdate_ARM">ManagedClusterUpdate_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.UpdateRunProperties_ARM">UpdateRunProperties_ARM</a>)
</p>
<div>
<p>The update to be applied to the ManagedClusters.</p>
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
<code>upgrade</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.ManagedClusterUpgradeSpec_ARM">
ManagedClusterUpgradeSpec_ARM
</a>
</em>
</td>
<td>
<p>Upgrade: The upgrade to apply to the ManagedClusters.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.ManagedClusterUpdate_STATUS">ManagedClusterUpdate_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.Fleets_UpdateRun_STATUS">Fleets_UpdateRun_STATUS</a>)
</p>
<div>
<p>The update to be applied to the ManagedClusters.</p>
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
<code>upgrade</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.ManagedClusterUpgradeSpec_STATUS">
ManagedClusterUpgradeSpec_STATUS
</a>
</em>
</td>
<td>
<p>Upgrade: The upgrade to apply to the ManagedClusters.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.ManagedClusterUpdate_STATUS_ARM">ManagedClusterUpdate_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.UpdateRunProperties_STATUS_ARM">UpdateRunProperties_STATUS_ARM</a>)
</p>
<div>
<p>The update to be applied to the ManagedClusters.</p>
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
<code>upgrade</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.ManagedClusterUpgradeSpec_STATUS_ARM">
ManagedClusterUpgradeSpec_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Upgrade: The upgrade to apply to the ManagedClusters.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.ManagedClusterUpgradeSpec">ManagedClusterUpgradeSpec
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.ManagedClusterUpdate">ManagedClusterUpdate</a>)
</p>
<div>
<p>The upgrade to apply to a ManagedCluster.</p>
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
<code>kubernetesVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>KubernetesVersion: The Kubernetes version to upgrade the member clusters to.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.ManagedClusterUpgradeType">
ManagedClusterUpgradeType
</a>
</em>
</td>
<td>
<p>Type: The upgrade type.
Full requires the KubernetesVersion property to be set.
NodeImageOnly requires the KubernetesVersion property not to be set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.ManagedClusterUpgradeSpec_ARM">ManagedClusterUpgradeSpec_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.ManagedClusterUpdate_ARM">ManagedClusterUpdate_ARM</a>)
</p>
<div>
<p>The upgrade to apply to a ManagedCluster.</p>
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
<code>kubernetesVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>KubernetesVersion: The Kubernetes version to upgrade the member clusters to.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.ManagedClusterUpgradeType">
ManagedClusterUpgradeType
</a>
</em>
</td>
<td>
<p>Type: The upgrade type.
Full requires the KubernetesVersion property to be set.
NodeImageOnly requires the KubernetesVersion property not to be set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.ManagedClusterUpgradeSpec_STATUS">ManagedClusterUpgradeSpec_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.ManagedClusterUpdate_STATUS">ManagedClusterUpdate_STATUS</a>)
</p>
<div>
<p>The upgrade to apply to a ManagedCluster.</p>
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
<code>kubernetesVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>KubernetesVersion: The Kubernetes version to upgrade the member clusters to.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.ManagedClusterUpgradeType_STATUS">
ManagedClusterUpgradeType_STATUS
</a>
</em>
</td>
<td>
<p>Type: The upgrade type.
Full requires the KubernetesVersion property to be set.
NodeImageOnly requires the KubernetesVersion property not to be set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.ManagedClusterUpgradeSpec_STATUS_ARM">ManagedClusterUpgradeSpec_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.ManagedClusterUpdate_STATUS_ARM">ManagedClusterUpdate_STATUS_ARM</a>)
</p>
<div>
<p>The upgrade to apply to a ManagedCluster.</p>
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
<code>kubernetesVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>KubernetesVersion: The Kubernetes version to upgrade the member clusters to.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.ManagedClusterUpgradeType_STATUS">
ManagedClusterUpgradeType_STATUS
</a>
</em>
</td>
<td>
<p>Type: The upgrade type.
Full requires the KubernetesVersion property to be set.
NodeImageOnly requires the KubernetesVersion property not to be set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.ManagedClusterUpgradeType">ManagedClusterUpgradeType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.ManagedClusterUpgradeSpec">ManagedClusterUpgradeSpec</a>, <a href="#containerservice.azure.com/v1api20230315preview.ManagedClusterUpgradeSpec_ARM">ManagedClusterUpgradeSpec_ARM</a>)
</p>
<div>
<p>The type of upgrade to perform when targeting ManagedClusters.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Full&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;NodeImageOnly&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.ManagedClusterUpgradeType_STATUS">ManagedClusterUpgradeType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.ManagedClusterUpgradeSpec_STATUS">ManagedClusterUpgradeSpec_STATUS</a>, <a href="#containerservice.azure.com/v1api20230315preview.ManagedClusterUpgradeSpec_STATUS_ARM">ManagedClusterUpgradeSpec_STATUS_ARM</a>)
</p>
<div>
<p>The type of upgrade to perform when targeting ManagedClusters.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Full&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;NodeImageOnly&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.MemberUpdateStatus_STATUS">MemberUpdateStatus_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.UpdateGroupStatus_STATUS">UpdateGroupStatus_STATUS</a>)
</p>
<div>
<p>The status of a member update operation.</p>
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
<code>clusterResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClusterResourceId: The Azure resource id of the target Kubernetes cluster.</p>
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
<p>Name: The name of the FleetMember.</p>
</td>
</tr>
<tr>
<td>
<code>operationId</code><br/>
<em>
string
</em>
</td>
<td>
<p>OperationId: The operation resource id of the latest attempt to perform the operation.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateStatus_STATUS">
UpdateStatus_STATUS
</a>
</em>
</td>
<td>
<p>Status: The status of the MemberUpdate operation.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.MemberUpdateStatus_STATUS_ARM">MemberUpdateStatus_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.UpdateGroupStatus_STATUS_ARM">UpdateGroupStatus_STATUS_ARM</a>)
</p>
<div>
<p>The status of a member update operation.</p>
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
<code>clusterResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClusterResourceId: The Azure resource id of the target Kubernetes cluster.</p>
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
<p>Name: The name of the FleetMember.</p>
</td>
</tr>
<tr>
<td>
<code>operationId</code><br/>
<em>
string
</em>
</td>
<td>
<p>OperationId: The operation resource id of the latest attempt to perform the operation.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateStatus_STATUS_ARM">
UpdateStatus_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Status: The status of the MemberUpdate operation.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.SystemData_CreatedByType_STATUS">SystemData_CreatedByType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.SystemData_STATUS">SystemData_STATUS</a>, <a href="#containerservice.azure.com/v1api20230315preview.SystemData_STATUS_ARM">SystemData_STATUS_ARM</a>)
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
<h3 id="containerservice.azure.com/v1api20230315preview.SystemData_LastModifiedByType_STATUS">SystemData_LastModifiedByType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.SystemData_STATUS">SystemData_STATUS</a>, <a href="#containerservice.azure.com/v1api20230315preview.SystemData_STATUS_ARM">SystemData_STATUS_ARM</a>)
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
<h3 id="containerservice.azure.com/v1api20230315preview.SystemData_STATUS">SystemData_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.Fleet_STATUS">Fleet_STATUS</a>, <a href="#containerservice.azure.com/v1api20230315preview.Fleets_Member_STATUS">Fleets_Member_STATUS</a>, <a href="#containerservice.azure.com/v1api20230315preview.Fleets_UpdateRun_STATUS">Fleets_UpdateRun_STATUS</a>)
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
<a href="#containerservice.azure.com/v1api20230315preview.SystemData_CreatedByType_STATUS">
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
<a href="#containerservice.azure.com/v1api20230315preview.SystemData_LastModifiedByType_STATUS">
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
<h3 id="containerservice.azure.com/v1api20230315preview.SystemData_STATUS_ARM">SystemData_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.Fleet_STATUS_ARM">Fleet_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230315preview.Fleets_Member_STATUS_ARM">Fleets_Member_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230315preview.Fleets_UpdateRun_STATUS_ARM">Fleets_UpdateRun_STATUS_ARM</a>)
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
<a href="#containerservice.azure.com/v1api20230315preview.SystemData_CreatedByType_STATUS">
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
<a href="#containerservice.azure.com/v1api20230315preview.SystemData_LastModifiedByType_STATUS">
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
<h3 id="containerservice.azure.com/v1api20230315preview.UpdateGroup">UpdateGroup
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.UpdateStage">UpdateStage</a>)
</p>
<div>
<p>A group to be updated.</p>
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
<p>Name: Name of the group.
It must match a group name of an existing fleet member.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.UpdateGroupStatus_STATUS">UpdateGroupStatus_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.UpdateStageStatus_STATUS">UpdateStageStatus_STATUS</a>)
</p>
<div>
<p>The status of a UpdateGroup.</p>
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
<code>members</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.MemberUpdateStatus_STATUS">
[]MemberUpdateStatus_STATUS
</a>
</em>
</td>
<td>
<p>Members: The list of member this UpdateGroup updates.</p>
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
<p>Name: The name of the UpdateGroup.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateStatus_STATUS">
UpdateStatus_STATUS
</a>
</em>
</td>
<td>
<p>Status: The status of the UpdateGroup.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.UpdateGroupStatus_STATUS_ARM">UpdateGroupStatus_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.UpdateStageStatus_STATUS_ARM">UpdateStageStatus_STATUS_ARM</a>)
</p>
<div>
<p>The status of a UpdateGroup.</p>
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
<code>members</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.MemberUpdateStatus_STATUS_ARM">
[]MemberUpdateStatus_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Members: The list of member this UpdateGroup updates.</p>
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
<p>Name: The name of the UpdateGroup.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateStatus_STATUS_ARM">
UpdateStatus_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Status: The status of the UpdateGroup.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.UpdateGroup_ARM">UpdateGroup_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.UpdateStage_ARM">UpdateStage_ARM</a>)
</p>
<div>
<p>A group to be updated.</p>
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
<p>Name: Name of the group.
It must match a group name of an existing fleet member.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.UpdateGroup_STATUS">UpdateGroup_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.UpdateStage_STATUS">UpdateStage_STATUS</a>)
</p>
<div>
<p>A group to be updated.</p>
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
<p>Name: Name of the group.
It must match a group name of an existing fleet member.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.UpdateGroup_STATUS_ARM">UpdateGroup_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.UpdateStage_STATUS_ARM">UpdateStage_STATUS_ARM</a>)
</p>
<div>
<p>A group to be updated.</p>
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
<p>Name: Name of the group.
It must match a group name of an existing fleet member.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.UpdateRunProperties_ARM">UpdateRunProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.Fleets_UpdateRun_Spec_ARM">Fleets_UpdateRun_Spec_ARM</a>)
</p>
<div>
<p>The properties of the UpdateRun.</p>
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
<code>managedClusterUpdate</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.ManagedClusterUpdate_ARM">
ManagedClusterUpdate_ARM
</a>
</em>
</td>
<td>
<p>ManagedClusterUpdate: The update to be applied to all clusters in the UpdateRun. The managedClusterUpdate can be
modified until the run is started.</p>
</td>
</tr>
<tr>
<td>
<code>strategy</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateRunStrategy_ARM">
UpdateRunStrategy_ARM
</a>
</em>
</td>
<td>
<p>Strategy: The strategy defines the order in which the clusters will be updated.
If not set, all members will be updated sequentially. The UpdateRun status will show a single UpdateStage and a single
UpdateGroup targeting all members.
The strategy of the UpdateRun can be modified until the run is started.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.UpdateRunProperties_STATUS_ARM">UpdateRunProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.Fleets_UpdateRun_STATUS_ARM">Fleets_UpdateRun_STATUS_ARM</a>)
</p>
<div>
<p>The properties of the UpdateRun.</p>
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
<code>managedClusterUpdate</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.ManagedClusterUpdate_STATUS_ARM">
ManagedClusterUpdate_STATUS_ARM
</a>
</em>
</td>
<td>
<p>ManagedClusterUpdate: The update to be applied to all clusters in the UpdateRun. The managedClusterUpdate can be
modified until the run is started.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateRunProvisioningState_STATUS">
UpdateRunProvisioningState_STATUS
</a>
</em>
</td>
<td>
<p>ProvisioningState: The provisioning state of the UpdateRun resource.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateRunStatus_STATUS_ARM">
UpdateRunStatus_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Status: The status of the UpdateRun.</p>
</td>
</tr>
<tr>
<td>
<code>strategy</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateRunStrategy_STATUS_ARM">
UpdateRunStrategy_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Strategy: The strategy defines the order in which the clusters will be updated.
If not set, all members will be updated sequentially. The UpdateRun status will show a single UpdateStage and a single
UpdateGroup targeting all members.
The strategy of the UpdateRun can be modified until the run is started.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.UpdateRunProvisioningState_STATUS">UpdateRunProvisioningState_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.Fleets_UpdateRun_STATUS">Fleets_UpdateRun_STATUS</a>, <a href="#containerservice.azure.com/v1api20230315preview.UpdateRunProperties_STATUS_ARM">UpdateRunProperties_STATUS_ARM</a>)
</p>
<div>
<p>The provisioning state of the UpdateRun resource.</p>
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
</tr><tr><td><p>&#34;Failed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Succeeded&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.UpdateRunStatus_STATUS">UpdateRunStatus_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.Fleets_UpdateRun_STATUS">Fleets_UpdateRun_STATUS</a>)
</p>
<div>
<p>The status of a UpdateRun.</p>
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
<code>stages</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateStageStatus_STATUS">
[]UpdateStageStatus_STATUS
</a>
</em>
</td>
<td>
<p>Stages: The stages composing an update run. Stages are run sequentially withing an UpdateRun.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateStatus_STATUS">
UpdateStatus_STATUS
</a>
</em>
</td>
<td>
<p>Status: The status of the UpdateRun.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.UpdateRunStatus_STATUS_ARM">UpdateRunStatus_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.UpdateRunProperties_STATUS_ARM">UpdateRunProperties_STATUS_ARM</a>)
</p>
<div>
<p>The status of a UpdateRun.</p>
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
<code>stages</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateStageStatus_STATUS_ARM">
[]UpdateStageStatus_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Stages: The stages composing an update run. Stages are run sequentially withing an UpdateRun.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateStatus_STATUS_ARM">
UpdateStatus_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Status: The status of the UpdateRun.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.UpdateRunStrategy">UpdateRunStrategy
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.Fleets_UpdateRun_Spec">Fleets_UpdateRun_Spec</a>)
</p>
<div>
<p>Defines the update sequence of the clusters via stages and groups.
Stages within a run are executed sequentially one
after another.
Groups within a stage are executed in parallel.
Member clusters within a group are updated sequentially
one after another.
A valid strategy contains no duplicate groups within or across stages.</p>
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
<code>stages</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateStage">
[]UpdateStage
</a>
</em>
</td>
<td>
<p>Stages: The list of stages that compose this update run. Min size: 1.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.UpdateRunStrategy_ARM">UpdateRunStrategy_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.UpdateRunProperties_ARM">UpdateRunProperties_ARM</a>)
</p>
<div>
<p>Defines the update sequence of the clusters via stages and groups.
Stages within a run are executed sequentially one
after another.
Groups within a stage are executed in parallel.
Member clusters within a group are updated sequentially
one after another.
A valid strategy contains no duplicate groups within or across stages.</p>
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
<code>stages</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateStage_ARM">
[]UpdateStage_ARM
</a>
</em>
</td>
<td>
<p>Stages: The list of stages that compose this update run. Min size: 1.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.UpdateRunStrategy_STATUS">UpdateRunStrategy_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.Fleets_UpdateRun_STATUS">Fleets_UpdateRun_STATUS</a>)
</p>
<div>
<p>Defines the update sequence of the clusters via stages and groups.
Stages within a run are executed sequentially one
after another.
Groups within a stage are executed in parallel.
Member clusters within a group are updated sequentially
one after another.
A valid strategy contains no duplicate groups within or across stages.</p>
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
<code>stages</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateStage_STATUS">
[]UpdateStage_STATUS
</a>
</em>
</td>
<td>
<p>Stages: The list of stages that compose this update run. Min size: 1.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.UpdateRunStrategy_STATUS_ARM">UpdateRunStrategy_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.UpdateRunProperties_STATUS_ARM">UpdateRunProperties_STATUS_ARM</a>)
</p>
<div>
<p>Defines the update sequence of the clusters via stages and groups.
Stages within a run are executed sequentially one
after another.
Groups within a stage are executed in parallel.
Member clusters within a group are updated sequentially
one after another.
A valid strategy contains no duplicate groups within or across stages.</p>
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
<code>stages</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateStage_STATUS_ARM">
[]UpdateStage_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Stages: The list of stages that compose this update run. Min size: 1.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.UpdateStage">UpdateStage
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.UpdateRunStrategy">UpdateRunStrategy</a>)
</p>
<div>
<p>Defines a stage which contains the groups to update and the steps to take (e.g., wait for a time period) before starting
the next stage.</p>
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
<code>afterStageWaitInSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>AfterStageWaitInSeconds: The time in seconds to wait at the end of this stage before starting the next one. Defaults to
0 seconds if unspecified.</p>
</td>
</tr>
<tr>
<td>
<code>groups</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateGroup">
[]UpdateGroup
</a>
</em>
</td>
<td>
<p>Groups: Defines the groups to be executed in parallel in this stage. Duplicate groups are not allowed. Min size: 1.</p>
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
<p>Name: The name of the stage. Must be unique within the UpdateRun.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.UpdateStageStatus_STATUS">UpdateStageStatus_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.UpdateRunStatus_STATUS">UpdateRunStatus_STATUS</a>)
</p>
<div>
<p>The status of a UpdateStage.</p>
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
<code>afterStageWaitStatus</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.WaitStatus_STATUS">
WaitStatus_STATUS
</a>
</em>
</td>
<td>
<p>AfterStageWaitStatus: The status of the wait period configured on the UpdateStage.</p>
</td>
</tr>
<tr>
<td>
<code>groups</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateGroupStatus_STATUS">
[]UpdateGroupStatus_STATUS
</a>
</em>
</td>
<td>
<p>Groups: The list of groups to be updated as part of this UpdateStage.</p>
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
<p>Name: The name of the UpdateStage.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateStatus_STATUS">
UpdateStatus_STATUS
</a>
</em>
</td>
<td>
<p>Status: The status of the UpdateStage.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.UpdateStageStatus_STATUS_ARM">UpdateStageStatus_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.UpdateRunStatus_STATUS_ARM">UpdateRunStatus_STATUS_ARM</a>)
</p>
<div>
<p>The status of a UpdateStage.</p>
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
<code>afterStageWaitStatus</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.WaitStatus_STATUS_ARM">
WaitStatus_STATUS_ARM
</a>
</em>
</td>
<td>
<p>AfterStageWaitStatus: The status of the wait period configured on the UpdateStage.</p>
</td>
</tr>
<tr>
<td>
<code>groups</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateGroupStatus_STATUS_ARM">
[]UpdateGroupStatus_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Groups: The list of groups to be updated as part of this UpdateStage.</p>
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
<p>Name: The name of the UpdateStage.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateStatus_STATUS_ARM">
UpdateStatus_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Status: The status of the UpdateStage.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.UpdateStage_ARM">UpdateStage_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.UpdateRunStrategy_ARM">UpdateRunStrategy_ARM</a>)
</p>
<div>
<p>Defines a stage which contains the groups to update and the steps to take (e.g., wait for a time period) before starting
the next stage.</p>
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
<code>afterStageWaitInSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>AfterStageWaitInSeconds: The time in seconds to wait at the end of this stage before starting the next one. Defaults to
0 seconds if unspecified.</p>
</td>
</tr>
<tr>
<td>
<code>groups</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateGroup_ARM">
[]UpdateGroup_ARM
</a>
</em>
</td>
<td>
<p>Groups: Defines the groups to be executed in parallel in this stage. Duplicate groups are not allowed. Min size: 1.</p>
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
<p>Name: The name of the stage. Must be unique within the UpdateRun.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.UpdateStage_STATUS">UpdateStage_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.UpdateRunStrategy_STATUS">UpdateRunStrategy_STATUS</a>)
</p>
<div>
<p>Defines a stage which contains the groups to update and the steps to take (e.g., wait for a time period) before starting
the next stage.</p>
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
<code>afterStageWaitInSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>AfterStageWaitInSeconds: The time in seconds to wait at the end of this stage before starting the next one. Defaults to
0 seconds if unspecified.</p>
</td>
</tr>
<tr>
<td>
<code>groups</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateGroup_STATUS">
[]UpdateGroup_STATUS
</a>
</em>
</td>
<td>
<p>Groups: Defines the groups to be executed in parallel in this stage. Duplicate groups are not allowed. Min size: 1.</p>
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
<p>Name: The name of the stage. Must be unique within the UpdateRun.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.UpdateStage_STATUS_ARM">UpdateStage_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.UpdateRunStrategy_STATUS_ARM">UpdateRunStrategy_STATUS_ARM</a>)
</p>
<div>
<p>Defines a stage which contains the groups to update and the steps to take (e.g., wait for a time period) before starting
the next stage.</p>
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
<code>afterStageWaitInSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>AfterStageWaitInSeconds: The time in seconds to wait at the end of this stage before starting the next one. Defaults to
0 seconds if unspecified.</p>
</td>
</tr>
<tr>
<td>
<code>groups</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateGroup_STATUS_ARM">
[]UpdateGroup_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Groups: Defines the groups to be executed in parallel in this stage. Duplicate groups are not allowed. Min size: 1.</p>
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
<p>Name: The name of the stage. Must be unique within the UpdateRun.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.UpdateState_STATUS">UpdateState_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.UpdateStatus_STATUS">UpdateStatus_STATUS</a>, <a href="#containerservice.azure.com/v1api20230315preview.UpdateStatus_STATUS_ARM">UpdateStatus_STATUS_ARM</a>)
</p>
<div>
<p>The state of the UpdateRun, UpdateStage, UpdateGroup, or MemberUpdate.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Completed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Failed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;NotStarted&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Running&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Stopped&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Stopping&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.UpdateStatus_STATUS">UpdateStatus_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.MemberUpdateStatus_STATUS">MemberUpdateStatus_STATUS</a>, <a href="#containerservice.azure.com/v1api20230315preview.UpdateGroupStatus_STATUS">UpdateGroupStatus_STATUS</a>, <a href="#containerservice.azure.com/v1api20230315preview.UpdateRunStatus_STATUS">UpdateRunStatus_STATUS</a>, <a href="#containerservice.azure.com/v1api20230315preview.UpdateStageStatus_STATUS">UpdateStageStatus_STATUS</a>, <a href="#containerservice.azure.com/v1api20230315preview.WaitStatus_STATUS">WaitStatus_STATUS</a>)
</p>
<div>
<p>The status for an operation or group of operations.</p>
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
<code>completedTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>CompletedTime: The time the operation or group was completed.</p>
</td>
</tr>
<tr>
<td>
<code>error</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.ErrorDetail_STATUS">
ErrorDetail_STATUS
</a>
</em>
</td>
<td>
<p>Error: The error details when a failure is encountered.</p>
</td>
</tr>
<tr>
<td>
<code>startTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>StartTime: The time the operation or group was started.</p>
</td>
</tr>
<tr>
<td>
<code>state</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateState_STATUS">
UpdateState_STATUS
</a>
</em>
</td>
<td>
<p>State: The State of the operation or group.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.UpdateStatus_STATUS_ARM">UpdateStatus_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.MemberUpdateStatus_STATUS_ARM">MemberUpdateStatus_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230315preview.UpdateGroupStatus_STATUS_ARM">UpdateGroupStatus_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230315preview.UpdateRunStatus_STATUS_ARM">UpdateRunStatus_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230315preview.UpdateStageStatus_STATUS_ARM">UpdateStageStatus_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230315preview.WaitStatus_STATUS_ARM">WaitStatus_STATUS_ARM</a>)
</p>
<div>
<p>The status for an operation or group of operations.</p>
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
<code>completedTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>CompletedTime: The time the operation or group was completed.</p>
</td>
</tr>
<tr>
<td>
<code>error</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.ErrorDetail_STATUS_ARM">
ErrorDetail_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Error: The error details when a failure is encountered.</p>
</td>
</tr>
<tr>
<td>
<code>startTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>StartTime: The time the operation or group was started.</p>
</td>
</tr>
<tr>
<td>
<code>state</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230315preview.UpdateState_STATUS">
UpdateState_STATUS
</a>
</em>
</td>
<td>
<p>State: The State of the operation or group.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.WaitStatus_STATUS">WaitStatus_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.UpdateStageStatus_STATUS">UpdateStageStatus_STATUS</a>)
</p>
<div>
<p>The status of the wait duration.</p>
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
<a href="#containerservice.azure.com/v1api20230315preview.UpdateStatus_STATUS">
UpdateStatus_STATUS
</a>
</em>
</td>
<td>
<p>Status: The status of the wait duration.</p>
</td>
</tr>
<tr>
<td>
<code>waitDurationInSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>WaitDurationInSeconds: The wait duration configured in seconds.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230315preview.WaitStatus_STATUS_ARM">WaitStatus_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230315preview.UpdateStageStatus_STATUS_ARM">UpdateStageStatus_STATUS_ARM</a>)
</p>
<div>
<p>The status of the wait duration.</p>
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
<a href="#containerservice.azure.com/v1api20230315preview.UpdateStatus_STATUS_ARM">
UpdateStatus_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Status: The status of the wait duration.</p>
</td>
</tr>
<tr>
<td>
<code>waitDurationInSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>WaitDurationInSeconds: The wait duration configured in seconds.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
