---
title: containerservice.azure.com/v1api20230202preview
---
<h2 id="containerservice.azure.com/v1api20230202preview">containerservice.azure.com/v1api20230202preview</h2>
<div>
<p>Package v1api20230202preview contains API Schema definitions for the containerservice v1api20230202preview API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="containerservice.azure.com/v1api20230202preview.APIVersion">APIVersion
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
<tbody><tr><td><p>&#34;2023-02-02-preview&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.AgentPoolMode">AgentPoolMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_ARM">ManagedClusterAgentPoolProfileProperties_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_ARM">ManagedClusterAgentPoolProfile_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_Spec">ManagedClusters_AgentPool_Spec</a>)
</p>
<div>
<p>A cluster must have at least one &lsquo;System&rsquo; Agent Pool at all times. For additional information on agent pool restrictions
and best practices, see: <a href="https://docs.microsoft.com/azure/aks/use-system-pools">https://docs.microsoft.com/azure/aks/use-system-pools</a></p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;System&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;User&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.AgentPoolMode_STATUS">AgentPoolMode_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_STATUS_ARM">ManagedClusterAgentPoolProfileProperties_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS">ManagedClusterAgentPoolProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS_ARM">ManagedClusterAgentPoolProfile_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_STATUS">ManagedClusters_AgentPool_STATUS</a>)
</p>
<div>
<p>A cluster must have at least one &lsquo;System&rsquo; Agent Pool at all times. For additional information on agent pool restrictions
and best practices, see: <a href="https://docs.microsoft.com/azure/aks/use-system-pools">https://docs.microsoft.com/azure/aks/use-system-pools</a></p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;System&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;User&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.AgentPoolNetworkProfile">AgentPoolNetworkProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_Spec">ManagedClusters_AgentPool_Spec</a>)
</p>
<div>
<p>Network settings of an agent pool.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allowedHostPorts</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PortRange">
[]PortRange
</a>
</em>
</td>
<td>
<p>AllowedHostPorts: The port ranges that are allowed to access. The specified ranges are allowed to overlap.</p>
</td>
</tr>
<tr>
<td>
<code>applicationSecurityGroupsReferences</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
[]genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>ApplicationSecurityGroupsReferences: The IDs of the application security groups which agent pool will associate when
created.</p>
</td>
</tr>
<tr>
<td>
<code>nodePublicIPTags</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.IPTag">
[]IPTag
</a>
</em>
</td>
<td>
<p>NodePublicIPTags: IPTags of instance-level public IPs.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.AgentPoolNetworkProfile_ARM">AgentPoolNetworkProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_ARM">ManagedClusterAgentPoolProfileProperties_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_ARM">ManagedClusterAgentPoolProfile_ARM</a>)
</p>
<div>
<p>Network settings of an agent pool.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allowedHostPorts</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PortRange_ARM">
[]PortRange_ARM
</a>
</em>
</td>
<td>
<p>AllowedHostPorts: The port ranges that are allowed to access. The specified ranges are allowed to overlap.</p>
</td>
</tr>
<tr>
<td>
<code>applicationSecurityGroups</code><br/>
<em>
[]string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>nodePublicIPTags</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.IPTag_ARM">
[]IPTag_ARM
</a>
</em>
</td>
<td>
<p>NodePublicIPTags: IPTags of instance-level public IPs.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.AgentPoolNetworkProfile_STATUS">AgentPoolNetworkProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS">ManagedClusterAgentPoolProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_STATUS">ManagedClusters_AgentPool_STATUS</a>)
</p>
<div>
<p>Network settings of an agent pool.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allowedHostPorts</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PortRange_STATUS">
[]PortRange_STATUS
</a>
</em>
</td>
<td>
<p>AllowedHostPorts: The port ranges that are allowed to access. The specified ranges are allowed to overlap.</p>
</td>
</tr>
<tr>
<td>
<code>applicationSecurityGroups</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>ApplicationSecurityGroups: The IDs of the application security groups which agent pool will associate when created.</p>
</td>
</tr>
<tr>
<td>
<code>nodePublicIPTags</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.IPTag_STATUS">
[]IPTag_STATUS
</a>
</em>
</td>
<td>
<p>NodePublicIPTags: IPTags of instance-level public IPs.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.AgentPoolNetworkProfile_STATUS_ARM">AgentPoolNetworkProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_STATUS_ARM">ManagedClusterAgentPoolProfileProperties_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS_ARM">ManagedClusterAgentPoolProfile_STATUS_ARM</a>)
</p>
<div>
<p>Network settings of an agent pool.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allowedHostPorts</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PortRange_STATUS_ARM">
[]PortRange_STATUS_ARM
</a>
</em>
</td>
<td>
<p>AllowedHostPorts: The port ranges that are allowed to access. The specified ranges are allowed to overlap.</p>
</td>
</tr>
<tr>
<td>
<code>applicationSecurityGroups</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>ApplicationSecurityGroups: The IDs of the application security groups which agent pool will associate when created.</p>
</td>
</tr>
<tr>
<td>
<code>nodePublicIPTags</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.IPTag_STATUS_ARM">
[]IPTag_STATUS_ARM
</a>
</em>
</td>
<td>
<p>NodePublicIPTags: IPTags of instance-level public IPs.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.AgentPoolType">AgentPoolType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_ARM">ManagedClusterAgentPoolProfileProperties_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_ARM">ManagedClusterAgentPoolProfile_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_Spec">ManagedClusters_AgentPool_Spec</a>)
</p>
<div>
<p>The type of Agent Pool.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;AvailabilitySet&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VirtualMachineScaleSets&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.AgentPoolType_STATUS">AgentPoolType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_STATUS_ARM">ManagedClusterAgentPoolProfileProperties_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS">ManagedClusterAgentPoolProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS_ARM">ManagedClusterAgentPoolProfile_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_STATUS">ManagedClusters_AgentPool_STATUS</a>)
</p>
<div>
<p>The type of Agent Pool.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;AvailabilitySet&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VirtualMachineScaleSets&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.AgentPoolUpgradeSettings">AgentPoolUpgradeSettings
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_Spec">ManagedClusters_AgentPool_Spec</a>)
</p>
<div>
<p>Settings for upgrading an agentpool</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>maxSurge</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaxSurge: This can either be set to an integer (e.g. &lsquo;5&rsquo;) or a percentage (e.g. &lsquo;50%&rsquo;). If a percentage is specified, it
is the percentage of the total agent pool size at the time of the upgrade. For percentages, fractional nodes are rounded
up. If not specified, the default is 1. For more information, including best practices, see:
<a href="https://docs.microsoft.com/azure/aks/upgrade-cluster#customize-node-surge-upgrade">https://docs.microsoft.com/azure/aks/upgrade-cluster#customize-node-surge-upgrade</a></p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.AgentPoolUpgradeSettings_ARM">AgentPoolUpgradeSettings_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_ARM">ManagedClusterAgentPoolProfileProperties_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_ARM">ManagedClusterAgentPoolProfile_ARM</a>)
</p>
<div>
<p>Settings for upgrading an agentpool</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>maxSurge</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaxSurge: This can either be set to an integer (e.g. &lsquo;5&rsquo;) or a percentage (e.g. &lsquo;50%&rsquo;). If a percentage is specified, it
is the percentage of the total agent pool size at the time of the upgrade. For percentages, fractional nodes are rounded
up. If not specified, the default is 1. For more information, including best practices, see:
<a href="https://docs.microsoft.com/azure/aks/upgrade-cluster#customize-node-surge-upgrade">https://docs.microsoft.com/azure/aks/upgrade-cluster#customize-node-surge-upgrade</a></p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.AgentPoolUpgradeSettings_STATUS">AgentPoolUpgradeSettings_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS">ManagedClusterAgentPoolProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_STATUS">ManagedClusters_AgentPool_STATUS</a>)
</p>
<div>
<p>Settings for upgrading an agentpool</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>maxSurge</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaxSurge: This can either be set to an integer (e.g. &lsquo;5&rsquo;) or a percentage (e.g. &lsquo;50%&rsquo;). If a percentage is specified, it
is the percentage of the total agent pool size at the time of the upgrade. For percentages, fractional nodes are rounded
up. If not specified, the default is 1. For more information, including best practices, see:
<a href="https://docs.microsoft.com/azure/aks/upgrade-cluster#customize-node-surge-upgrade">https://docs.microsoft.com/azure/aks/upgrade-cluster#customize-node-surge-upgrade</a></p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.AgentPoolUpgradeSettings_STATUS_ARM">AgentPoolUpgradeSettings_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_STATUS_ARM">ManagedClusterAgentPoolProfileProperties_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS_ARM">ManagedClusterAgentPoolProfile_STATUS_ARM</a>)
</p>
<div>
<p>Settings for upgrading an agentpool</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>maxSurge</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaxSurge: This can either be set to an integer (e.g. &lsquo;5&rsquo;) or a percentage (e.g. &lsquo;50%&rsquo;). If a percentage is specified, it
is the percentage of the total agent pool size at the time of the upgrade. For percentages, fractional nodes are rounded
up. If not specified, the default is 1. For more information, including best practices, see:
<a href="https://docs.microsoft.com/azure/aks/upgrade-cluster#customize-node-surge-upgrade">https://docs.microsoft.com/azure/aks/upgrade-cluster#customize-node-surge-upgrade</a></p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.AgentPoolWindowsProfile">AgentPoolWindowsProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_Spec">ManagedClusters_AgentPool_Spec</a>)
</p>
<div>
<p>The Windows agent pool&rsquo;s specific profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>disableOutboundNat</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableOutboundNat: The default value is false. Outbound NAT can only be disabled if the cluster outboundType is NAT
Gateway and the Windows agent pool does not have node public IP enabled.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.AgentPoolWindowsProfile_ARM">AgentPoolWindowsProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_ARM">ManagedClusterAgentPoolProfileProperties_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_ARM">ManagedClusterAgentPoolProfile_ARM</a>)
</p>
<div>
<p>The Windows agent pool&rsquo;s specific profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>disableOutboundNat</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableOutboundNat: The default value is false. Outbound NAT can only be disabled if the cluster outboundType is NAT
Gateway and the Windows agent pool does not have node public IP enabled.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.AgentPoolWindowsProfile_STATUS">AgentPoolWindowsProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS">ManagedClusterAgentPoolProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_STATUS">ManagedClusters_AgentPool_STATUS</a>)
</p>
<div>
<p>The Windows agent pool&rsquo;s specific profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>disableOutboundNat</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableOutboundNat: The default value is false. Outbound NAT can only be disabled if the cluster outboundType is NAT
Gateway and the Windows agent pool does not have node public IP enabled.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.AgentPoolWindowsProfile_STATUS_ARM">AgentPoolWindowsProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_STATUS_ARM">ManagedClusterAgentPoolProfileProperties_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS_ARM">ManagedClusterAgentPoolProfile_STATUS_ARM</a>)
</p>
<div>
<p>The Windows agent pool&rsquo;s specific profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>disableOutboundNat</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableOutboundNat: The default value is false. Outbound NAT can only be disabled if the cluster outboundType is NAT
Gateway and the Windows agent pool does not have node public IP enabled.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.AzureKeyVaultKms">AzureKeyVaultKms
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile">ManagedClusterSecurityProfile</a>)
</p>
<div>
<p>Azure Key Vault key management service settings for the security profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable Azure Key Vault key management service. The default is false.</p>
</td>
</tr>
<tr>
<td>
<code>keyId</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyId: Identifier of Azure Key Vault key. See <a href="https://docs.microsoft.com/en-us/azure/key-vault/general/about-keys-secrets-certificates#vault-name-and-object-name">key identifier
format</a>
for more details. When Azure Key Vault key management service is enabled, this field is required and must be a valid key
identifier. When Azure Key Vault key management service is disabled, leave the field empty.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultNetworkAccess</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AzureKeyVaultKms_KeyVaultNetworkAccess">
AzureKeyVaultKms_KeyVaultNetworkAccess
</a>
</em>
</td>
<td>
<p>KeyVaultNetworkAccess: Network access of key vault. The possible values are <code>Public</code> and <code>Private</code>. <code>Public</code> means the
key vault allows public access from all networks. <code>Private</code> means the key vault disables public access and enables
private link. The default value is <code>Public</code>.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultResourceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>KeyVaultResourceReference: Resource ID of key vault. When keyVaultNetworkAccess is <code>Private</code>, this field is required and
must be a valid resource ID. When keyVaultNetworkAccess is <code>Public</code>, leave the field empty.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.AzureKeyVaultKms_ARM">AzureKeyVaultKms_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile_ARM">ManagedClusterSecurityProfile_ARM</a>)
</p>
<div>
<p>Azure Key Vault key management service settings for the security profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable Azure Key Vault key management service. The default is false.</p>
</td>
</tr>
<tr>
<td>
<code>keyId</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyId: Identifier of Azure Key Vault key. See <a href="https://docs.microsoft.com/en-us/azure/key-vault/general/about-keys-secrets-certificates#vault-name-and-object-name">key identifier
format</a>
for more details. When Azure Key Vault key management service is enabled, this field is required and must be a valid key
identifier. When Azure Key Vault key management service is disabled, leave the field empty.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultNetworkAccess</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AzureKeyVaultKms_KeyVaultNetworkAccess">
AzureKeyVaultKms_KeyVaultNetworkAccess
</a>
</em>
</td>
<td>
<p>KeyVaultNetworkAccess: Network access of key vault. The possible values are <code>Public</code> and <code>Private</code>. <code>Public</code> means the
key vault allows public access from all networks. <code>Private</code> means the key vault disables public access and enables
private link. The default value is <code>Public</code>.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultResourceId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.AzureKeyVaultKms_KeyVaultNetworkAccess">AzureKeyVaultKms_KeyVaultNetworkAccess
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.AzureKeyVaultKms">AzureKeyVaultKms</a>, <a href="#containerservice.azure.com/v1api20230202preview.AzureKeyVaultKms_ARM">AzureKeyVaultKms_ARM</a>)
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
<tbody><tr><td><p>&#34;Private&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Public&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.AzureKeyVaultKms_KeyVaultNetworkAccess_STATUS">AzureKeyVaultKms_KeyVaultNetworkAccess_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.AzureKeyVaultKms_STATUS">AzureKeyVaultKms_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.AzureKeyVaultKms_STATUS_ARM">AzureKeyVaultKms_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;Private&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Public&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.AzureKeyVaultKms_STATUS">AzureKeyVaultKms_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile_STATUS">ManagedClusterSecurityProfile_STATUS</a>)
</p>
<div>
<p>Azure Key Vault key management service settings for the security profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable Azure Key Vault key management service. The default is false.</p>
</td>
</tr>
<tr>
<td>
<code>keyId</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyId: Identifier of Azure Key Vault key. See <a href="https://docs.microsoft.com/en-us/azure/key-vault/general/about-keys-secrets-certificates#vault-name-and-object-name">key identifier
format</a>
for more details. When Azure Key Vault key management service is enabled, this field is required and must be a valid key
identifier. When Azure Key Vault key management service is disabled, leave the field empty.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultNetworkAccess</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AzureKeyVaultKms_KeyVaultNetworkAccess_STATUS">
AzureKeyVaultKms_KeyVaultNetworkAccess_STATUS
</a>
</em>
</td>
<td>
<p>KeyVaultNetworkAccess: Network access of key vault. The possible values are <code>Public</code> and <code>Private</code>. <code>Public</code> means the
key vault allows public access from all networks. <code>Private</code> means the key vault disables public access and enables
private link. The default value is <code>Public</code>.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVaultResourceId: Resource ID of key vault. When keyVaultNetworkAccess is <code>Private</code>, this field is required and must
be a valid resource ID. When keyVaultNetworkAccess is <code>Public</code>, leave the field empty.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.AzureKeyVaultKms_STATUS_ARM">AzureKeyVaultKms_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile_STATUS_ARM">ManagedClusterSecurityProfile_STATUS_ARM</a>)
</p>
<div>
<p>Azure Key Vault key management service settings for the security profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable Azure Key Vault key management service. The default is false.</p>
</td>
</tr>
<tr>
<td>
<code>keyId</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyId: Identifier of Azure Key Vault key. See <a href="https://docs.microsoft.com/en-us/azure/key-vault/general/about-keys-secrets-certificates#vault-name-and-object-name">key identifier
format</a>
for more details. When Azure Key Vault key management service is enabled, this field is required and must be a valid key
identifier. When Azure Key Vault key management service is disabled, leave the field empty.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultNetworkAccess</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AzureKeyVaultKms_KeyVaultNetworkAccess_STATUS">
AzureKeyVaultKms_KeyVaultNetworkAccess_STATUS
</a>
</em>
</td>
<td>
<p>KeyVaultNetworkAccess: Network access of key vault. The possible values are <code>Public</code> and <code>Private</code>. <code>Public</code> means the
key vault allows public access from all networks. <code>Private</code> means the key vault disables public access and enables
private link. The default value is <code>Public</code>.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVaultResourceId: Resource ID of key vault. When keyVaultNetworkAccess is <code>Private</code>, this field is required and must
be a valid resource ID. When keyVaultNetworkAccess is <code>Public</code>, leave the field empty.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ClusterUpgradeSettings">ClusterUpgradeSettings
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
</p>
<div>
<p>Settings for upgrading a cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>overrideSettings</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.UpgradeOverrideSettings">
UpgradeOverrideSettings
</a>
</em>
</td>
<td>
<p>OverrideSettings: Settings for overrides.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ClusterUpgradeSettings_ARM">ClusterUpgradeSettings_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM</a>)
</p>
<div>
<p>Settings for upgrading a cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>overrideSettings</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.UpgradeOverrideSettings_ARM">
UpgradeOverrideSettings_ARM
</a>
</em>
</td>
<td>
<p>OverrideSettings: Settings for overrides.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ClusterUpgradeSettings_STATUS">ClusterUpgradeSettings_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
</p>
<div>
<p>Settings for upgrading a cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>overrideSettings</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.UpgradeOverrideSettings_STATUS">
UpgradeOverrideSettings_STATUS
</a>
</em>
</td>
<td>
<p>OverrideSettings: Settings for overrides.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ClusterUpgradeSettings_STATUS_ARM">ClusterUpgradeSettings_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>)
</p>
<div>
<p>Settings for upgrading a cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>overrideSettings</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.UpgradeOverrideSettings_STATUS_ARM">
UpgradeOverrideSettings_STATUS_ARM
</a>
</em>
</td>
<td>
<p>OverrideSettings: Settings for overrides.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceLinuxProfile">ContainerServiceLinuxProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
</p>
<div>
<p>Profile for Linux VMs in the container service cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>adminUsername</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdminUsername: The administrator username to use for Linux VMs.</p>
</td>
</tr>
<tr>
<td>
<code>ssh</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceSshConfiguration">
ContainerServiceSshConfiguration
</a>
</em>
</td>
<td>
<p>Ssh: The SSH configuration for Linux-based VMs running on Azure.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceLinuxProfile_ARM">ContainerServiceLinuxProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM</a>)
</p>
<div>
<p>Profile for Linux VMs in the container service cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>adminUsername</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdminUsername: The administrator username to use for Linux VMs.</p>
</td>
</tr>
<tr>
<td>
<code>ssh</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceSshConfiguration_ARM">
ContainerServiceSshConfiguration_ARM
</a>
</em>
</td>
<td>
<p>Ssh: The SSH configuration for Linux-based VMs running on Azure.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceLinuxProfile_STATUS">ContainerServiceLinuxProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
</p>
<div>
<p>Profile for Linux VMs in the container service cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>adminUsername</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdminUsername: The administrator username to use for Linux VMs.</p>
</td>
</tr>
<tr>
<td>
<code>ssh</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceSshConfiguration_STATUS">
ContainerServiceSshConfiguration_STATUS
</a>
</em>
</td>
<td>
<p>Ssh: The SSH configuration for Linux-based VMs running on Azure.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceLinuxProfile_STATUS_ARM">ContainerServiceLinuxProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>)
</p>
<div>
<p>Profile for Linux VMs in the container service cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>adminUsername</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdminUsername: The administrator username to use for Linux VMs.</p>
</td>
</tr>
<tr>
<td>
<code>ssh</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceSshConfiguration_STATUS_ARM">
ContainerServiceSshConfiguration_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Ssh: The SSH configuration for Linux-based VMs running on Azure.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile">ContainerServiceNetworkProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
</p>
<div>
<p>Profile of network configuration.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dnsServiceIP</code><br/>
<em>
string
</em>
</td>
<td>
<p>DnsServiceIP: An IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address
range specified in serviceCidr.</p>
</td>
</tr>
<tr>
<td>
<code>dockerBridgeCidr</code><br/>
<em>
string
</em>
</td>
<td>
<p>DockerBridgeCidr: A CIDR notation IP range assigned to the Docker bridge network. It must not overlap with any Subnet IP
ranges or the Kubernetes service address range.</p>
</td>
</tr>
<tr>
<td>
<code>ipFamilies</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_IpFamilies">
[]ContainerServiceNetworkProfile_IpFamilies
</a>
</em>
</td>
<td>
<p>IpFamilies: IP families are used to determine single-stack or dual-stack clusters. For single-stack, the expected value
is IPv4. For dual-stack, the expected values are IPv4 and IPv6.</p>
</td>
</tr>
<tr>
<td>
<code>kubeProxyConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig">
ContainerServiceNetworkProfile_KubeProxyConfig
</a>
</em>
</td>
<td>
<p>KubeProxyConfig: Holds configuration customizations for kube-proxy. Any values not defined will use the kube-proxy
defaulting behavior. See <a href="https://v">https://v</a><version>.docs.kubernetes.io/docs/reference/command-line-tools-reference/kube-proxy/
where <version> is represented by a <major version>-<minor version> string. Kubernetes version 1.23 would be &lsquo;1-23&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile">
ManagedClusterLoadBalancerProfile
</a>
</em>
</td>
<td>
<p>LoadBalancerProfile: Profile of the cluster load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerSku</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.LoadBalancerSku">
LoadBalancerSku
</a>
</em>
</td>
<td>
<p>LoadBalancerSku: The default is &lsquo;standard&rsquo;. See <a href="https://docs.microsoft.com/azure/load-balancer/skus">Azure Load Balancer
SKUs</a> for more information about the differences between load
balancer SKUs.</p>
</td>
</tr>
<tr>
<td>
<code>natGatewayProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterNATGatewayProfile">
ManagedClusterNATGatewayProfile
</a>
</em>
</td>
<td>
<p>NatGatewayProfile: Profile of the cluster NAT gateway.</p>
</td>
</tr>
<tr>
<td>
<code>networkDataplane</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.NetworkDataplane">
NetworkDataplane
</a>
</em>
</td>
<td>
<p>NetworkDataplane: Network dataplane used in the Kubernetes cluster.</p>
</td>
</tr>
<tr>
<td>
<code>networkMode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.NetworkMode">
NetworkMode
</a>
</em>
</td>
<td>
<p>NetworkMode: This cannot be specified if networkPlugin is anything other than &lsquo;azure&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>networkPlugin</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.NetworkPlugin">
NetworkPlugin
</a>
</em>
</td>
<td>
<p>NetworkPlugin: Network plugin used for building the Kubernetes network.</p>
</td>
</tr>
<tr>
<td>
<code>networkPluginMode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.NetworkPluginMode">
NetworkPluginMode
</a>
</em>
</td>
<td>
<p>NetworkPluginMode: Network plugin mode used for building the Kubernetes network.</p>
</td>
</tr>
<tr>
<td>
<code>networkPolicy</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.NetworkPolicy">
NetworkPolicy
</a>
</em>
</td>
<td>
<p>NetworkPolicy: Network policy used for building the Kubernetes network.</p>
</td>
</tr>
<tr>
<td>
<code>outboundType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_OutboundType">
ContainerServiceNetworkProfile_OutboundType
</a>
</em>
</td>
<td>
<p>OutboundType: This can only be set at cluster creation time and cannot be changed later. For more information see
<a href="https://docs.microsoft.com/azure/aks/egress-outboundtype">egress outbound type</a>.</p>
</td>
</tr>
<tr>
<td>
<code>podCidr</code><br/>
<em>
string
</em>
</td>
<td>
<p>PodCidr: A CIDR notation IP range from which to assign pod IPs when kubenet is used.</p>
</td>
</tr>
<tr>
<td>
<code>podCidrs</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>PodCidrs: One IPv4 CIDR is expected for single-stack networking. Two CIDRs, one for each IP family (IPv4/IPv6), is
expected for dual-stack networking.</p>
</td>
</tr>
<tr>
<td>
<code>serviceCidr</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServiceCidr: A CIDR notation IP range from which to assign service cluster IPs. It must not overlap with any Subnet IP
ranges.</p>
</td>
</tr>
<tr>
<td>
<code>serviceCidrs</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>ServiceCidrs: One IPv4 CIDR is expected for single-stack networking. Two CIDRs, one for each IP family (IPv4/IPv6), is
expected for dual-stack networking. They must not overlap with any Subnet IP ranges.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_ARM">ContainerServiceNetworkProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM</a>)
</p>
<div>
<p>Profile of network configuration.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dnsServiceIP</code><br/>
<em>
string
</em>
</td>
<td>
<p>DnsServiceIP: An IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address
range specified in serviceCidr.</p>
</td>
</tr>
<tr>
<td>
<code>dockerBridgeCidr</code><br/>
<em>
string
</em>
</td>
<td>
<p>DockerBridgeCidr: A CIDR notation IP range assigned to the Docker bridge network. It must not overlap with any Subnet IP
ranges or the Kubernetes service address range.</p>
</td>
</tr>
<tr>
<td>
<code>ipFamilies</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_IpFamilies">
[]ContainerServiceNetworkProfile_IpFamilies
</a>
</em>
</td>
<td>
<p>IpFamilies: IP families are used to determine single-stack or dual-stack clusters. For single-stack, the expected value
is IPv4. For dual-stack, the expected values are IPv4 and IPv6.</p>
</td>
</tr>
<tr>
<td>
<code>kubeProxyConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_ARM">
ContainerServiceNetworkProfile_KubeProxyConfig_ARM
</a>
</em>
</td>
<td>
<p>KubeProxyConfig: Holds configuration customizations for kube-proxy. Any values not defined will use the kube-proxy
defaulting behavior. See <a href="https://v">https://v</a><version>.docs.kubernetes.io/docs/reference/command-line-tools-reference/kube-proxy/
where <version> is represented by a <major version>-<minor version> string. Kubernetes version 1.23 would be &lsquo;1-23&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_ARM">
ManagedClusterLoadBalancerProfile_ARM
</a>
</em>
</td>
<td>
<p>LoadBalancerProfile: Profile of the cluster load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerSku</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.LoadBalancerSku">
LoadBalancerSku
</a>
</em>
</td>
<td>
<p>LoadBalancerSku: The default is &lsquo;standard&rsquo;. See <a href="https://docs.microsoft.com/azure/load-balancer/skus">Azure Load Balancer
SKUs</a> for more information about the differences between load
balancer SKUs.</p>
</td>
</tr>
<tr>
<td>
<code>natGatewayProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterNATGatewayProfile_ARM">
ManagedClusterNATGatewayProfile_ARM
</a>
</em>
</td>
<td>
<p>NatGatewayProfile: Profile of the cluster NAT gateway.</p>
</td>
</tr>
<tr>
<td>
<code>networkDataplane</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.NetworkDataplane">
NetworkDataplane
</a>
</em>
</td>
<td>
<p>NetworkDataplane: Network dataplane used in the Kubernetes cluster.</p>
</td>
</tr>
<tr>
<td>
<code>networkMode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.NetworkMode">
NetworkMode
</a>
</em>
</td>
<td>
<p>NetworkMode: This cannot be specified if networkPlugin is anything other than &lsquo;azure&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>networkPlugin</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.NetworkPlugin">
NetworkPlugin
</a>
</em>
</td>
<td>
<p>NetworkPlugin: Network plugin used for building the Kubernetes network.</p>
</td>
</tr>
<tr>
<td>
<code>networkPluginMode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.NetworkPluginMode">
NetworkPluginMode
</a>
</em>
</td>
<td>
<p>NetworkPluginMode: Network plugin mode used for building the Kubernetes network.</p>
</td>
</tr>
<tr>
<td>
<code>networkPolicy</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.NetworkPolicy">
NetworkPolicy
</a>
</em>
</td>
<td>
<p>NetworkPolicy: Network policy used for building the Kubernetes network.</p>
</td>
</tr>
<tr>
<td>
<code>outboundType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_OutboundType">
ContainerServiceNetworkProfile_OutboundType
</a>
</em>
</td>
<td>
<p>OutboundType: This can only be set at cluster creation time and cannot be changed later. For more information see
<a href="https://docs.microsoft.com/azure/aks/egress-outboundtype">egress outbound type</a>.</p>
</td>
</tr>
<tr>
<td>
<code>podCidr</code><br/>
<em>
string
</em>
</td>
<td>
<p>PodCidr: A CIDR notation IP range from which to assign pod IPs when kubenet is used.</p>
</td>
</tr>
<tr>
<td>
<code>podCidrs</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>PodCidrs: One IPv4 CIDR is expected for single-stack networking. Two CIDRs, one for each IP family (IPv4/IPv6), is
expected for dual-stack networking.</p>
</td>
</tr>
<tr>
<td>
<code>serviceCidr</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServiceCidr: A CIDR notation IP range from which to assign service cluster IPs. It must not overlap with any Subnet IP
ranges.</p>
</td>
</tr>
<tr>
<td>
<code>serviceCidrs</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>ServiceCidrs: One IPv4 CIDR is expected for single-stack networking. Two CIDRs, one for each IP family (IPv4/IPv6), is
expected for dual-stack networking. They must not overlap with any Subnet IP ranges.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_IpFamilies">ContainerServiceNetworkProfile_IpFamilies
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile">ContainerServiceNetworkProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_ARM">ContainerServiceNetworkProfile_ARM</a>)
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
<tbody><tr><td><p>&#34;IPv4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;IPv6&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_IpFamilies_STATUS">ContainerServiceNetworkProfile_IpFamilies_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_STATUS">ContainerServiceNetworkProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_STATUS_ARM">ContainerServiceNetworkProfile_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;IPv4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;IPv6&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig">ContainerServiceNetworkProfile_KubeProxyConfig
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile">ContainerServiceNetworkProfile</a>)
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
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable on kube-proxy on the cluster (if no &lsquo;kubeProxyConfig&rsquo; exists, kube-proxy is enabled in AKS by
default without these customizations).</p>
</td>
</tr>
<tr>
<td>
<code>ipvsConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig">
ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig
</a>
</em>
</td>
<td>
<p>IpvsConfig: Holds configuration customizations for IPVS. May only be specified if &lsquo;mode&rsquo; is set to &lsquo;IPVS&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_Mode">
ContainerServiceNetworkProfile_KubeProxyConfig_Mode
</a>
</em>
</td>
<td>
<p>Mode: Specify which proxy mode to use (&lsquo;IPTABLES&rsquo; or &lsquo;IPVS&rsquo;)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_ARM">ContainerServiceNetworkProfile_KubeProxyConfig_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_ARM">ContainerServiceNetworkProfile_ARM</a>)
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
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable on kube-proxy on the cluster (if no &lsquo;kubeProxyConfig&rsquo; exists, kube-proxy is enabled in AKS by
default without these customizations).</p>
</td>
</tr>
<tr>
<td>
<code>ipvsConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_ARM">
ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_ARM
</a>
</em>
</td>
<td>
<p>IpvsConfig: Holds configuration customizations for IPVS. May only be specified if &lsquo;mode&rsquo; is set to &lsquo;IPVS&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_Mode">
ContainerServiceNetworkProfile_KubeProxyConfig_Mode
</a>
</em>
</td>
<td>
<p>Mode: Specify which proxy mode to use (&lsquo;IPTABLES&rsquo; or &lsquo;IPVS&rsquo;)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig">ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig">ContainerServiceNetworkProfile_KubeProxyConfig</a>)
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
<code>scheduler</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler">
ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler
</a>
</em>
</td>
<td>
<p>Scheduler: IPVS scheduler, for more information please see <a href="http://www.linuxvirtualserver.org/docs/scheduling.html">http://www.linuxvirtualserver.org/docs/scheduling.html</a>.</p>
</td>
</tr>
<tr>
<td>
<code>tcpFinTimeoutSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>TcpFinTimeoutSeconds: The timeout value used for IPVS TCP sessions after receiving a FIN in seconds. Must be a positive
integer value.</p>
</td>
</tr>
<tr>
<td>
<code>tcpTimeoutSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>TcpTimeoutSeconds: The timeout value used for idle IPVS TCP sessions in seconds. Must be a positive integer value.</p>
</td>
</tr>
<tr>
<td>
<code>udpTimeoutSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>UdpTimeoutSeconds: The timeout value used for IPVS UDP packets in seconds. Must be a positive integer value.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_ARM">ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_ARM">ContainerServiceNetworkProfile_KubeProxyConfig_ARM</a>)
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
<code>scheduler</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler">
ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler
</a>
</em>
</td>
<td>
<p>Scheduler: IPVS scheduler, for more information please see <a href="http://www.linuxvirtualserver.org/docs/scheduling.html">http://www.linuxvirtualserver.org/docs/scheduling.html</a>.</p>
</td>
</tr>
<tr>
<td>
<code>tcpFinTimeoutSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>TcpFinTimeoutSeconds: The timeout value used for IPVS TCP sessions after receiving a FIN in seconds. Must be a positive
integer value.</p>
</td>
</tr>
<tr>
<td>
<code>tcpTimeoutSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>TcpTimeoutSeconds: The timeout value used for idle IPVS TCP sessions in seconds. Must be a positive integer value.</p>
</td>
</tr>
<tr>
<td>
<code>udpTimeoutSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>UdpTimeoutSeconds: The timeout value used for IPVS UDP packets in seconds. Must be a positive integer value.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_STATUS">ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_STATUS">ContainerServiceNetworkProfile_KubeProxyConfig_STATUS</a>)
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
<code>scheduler</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler_STATUS">
ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler_STATUS
</a>
</em>
</td>
<td>
<p>Scheduler: IPVS scheduler, for more information please see <a href="http://www.linuxvirtualserver.org/docs/scheduling.html">http://www.linuxvirtualserver.org/docs/scheduling.html</a>.</p>
</td>
</tr>
<tr>
<td>
<code>tcpFinTimeoutSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>TcpFinTimeoutSeconds: The timeout value used for IPVS TCP sessions after receiving a FIN in seconds. Must be a positive
integer value.</p>
</td>
</tr>
<tr>
<td>
<code>tcpTimeoutSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>TcpTimeoutSeconds: The timeout value used for idle IPVS TCP sessions in seconds. Must be a positive integer value.</p>
</td>
</tr>
<tr>
<td>
<code>udpTimeoutSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>UdpTimeoutSeconds: The timeout value used for IPVS UDP packets in seconds. Must be a positive integer value.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_STATUS_ARM">ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_STATUS_ARM">ContainerServiceNetworkProfile_KubeProxyConfig_STATUS_ARM</a>)
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
<code>scheduler</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler_STATUS">
ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler_STATUS
</a>
</em>
</td>
<td>
<p>Scheduler: IPVS scheduler, for more information please see <a href="http://www.linuxvirtualserver.org/docs/scheduling.html">http://www.linuxvirtualserver.org/docs/scheduling.html</a>.</p>
</td>
</tr>
<tr>
<td>
<code>tcpFinTimeoutSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>TcpFinTimeoutSeconds: The timeout value used for IPVS TCP sessions after receiving a FIN in seconds. Must be a positive
integer value.</p>
</td>
</tr>
<tr>
<td>
<code>tcpTimeoutSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>TcpTimeoutSeconds: The timeout value used for idle IPVS TCP sessions in seconds. Must be a positive integer value.</p>
</td>
</tr>
<tr>
<td>
<code>udpTimeoutSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>UdpTimeoutSeconds: The timeout value used for IPVS UDP packets in seconds. Must be a positive integer value.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler">ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig">ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig</a>, <a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_ARM">ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_ARM</a>)
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
<tbody><tr><td><p>&#34;LeastConnection&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;RoundRobin&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler_STATUS">ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_Scheduler_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_STATUS">ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_STATUS_ARM">ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;LeastConnection&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;RoundRobin&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_Mode">ContainerServiceNetworkProfile_KubeProxyConfig_Mode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig">ContainerServiceNetworkProfile_KubeProxyConfig</a>, <a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_ARM">ContainerServiceNetworkProfile_KubeProxyConfig_ARM</a>)
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
<tbody><tr><td><p>&#34;IPTABLES&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;IPVS&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_Mode_STATUS">ContainerServiceNetworkProfile_KubeProxyConfig_Mode_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_STATUS">ContainerServiceNetworkProfile_KubeProxyConfig_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_STATUS_ARM">ContainerServiceNetworkProfile_KubeProxyConfig_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;IPTABLES&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;IPVS&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_STATUS">ContainerServiceNetworkProfile_KubeProxyConfig_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_STATUS">ContainerServiceNetworkProfile_STATUS</a>)
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
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable on kube-proxy on the cluster (if no &lsquo;kubeProxyConfig&rsquo; exists, kube-proxy is enabled in AKS by
default without these customizations).</p>
</td>
</tr>
<tr>
<td>
<code>ipvsConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_STATUS">
ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_STATUS
</a>
</em>
</td>
<td>
<p>IpvsConfig: Holds configuration customizations for IPVS. May only be specified if &lsquo;mode&rsquo; is set to &lsquo;IPVS&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_Mode_STATUS">
ContainerServiceNetworkProfile_KubeProxyConfig_Mode_STATUS
</a>
</em>
</td>
<td>
<p>Mode: Specify which proxy mode to use (&lsquo;IPTABLES&rsquo; or &lsquo;IPVS&rsquo;)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_STATUS_ARM">ContainerServiceNetworkProfile_KubeProxyConfig_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_STATUS_ARM">ContainerServiceNetworkProfile_STATUS_ARM</a>)
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
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable on kube-proxy on the cluster (if no &lsquo;kubeProxyConfig&rsquo; exists, kube-proxy is enabled in AKS by
default without these customizations).</p>
</td>
</tr>
<tr>
<td>
<code>ipvsConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_STATUS_ARM">
ContainerServiceNetworkProfile_KubeProxyConfig_IpvsConfig_STATUS_ARM
</a>
</em>
</td>
<td>
<p>IpvsConfig: Holds configuration customizations for IPVS. May only be specified if &lsquo;mode&rsquo; is set to &lsquo;IPVS&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_Mode_STATUS">
ContainerServiceNetworkProfile_KubeProxyConfig_Mode_STATUS
</a>
</em>
</td>
<td>
<p>Mode: Specify which proxy mode to use (&lsquo;IPTABLES&rsquo; or &lsquo;IPVS&rsquo;)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_OutboundType">ContainerServiceNetworkProfile_OutboundType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile">ContainerServiceNetworkProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_ARM">ContainerServiceNetworkProfile_ARM</a>)
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
<tbody><tr><td><p>&#34;loadBalancer&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;managedNATGateway&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;userAssignedNATGateway&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;userDefinedRouting&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_OutboundType_STATUS">ContainerServiceNetworkProfile_OutboundType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_STATUS">ContainerServiceNetworkProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_STATUS_ARM">ContainerServiceNetworkProfile_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;loadBalancer&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;managedNATGateway&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;userAssignedNATGateway&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;userDefinedRouting&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_STATUS">ContainerServiceNetworkProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
</p>
<div>
<p>Profile of network configuration.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dnsServiceIP</code><br/>
<em>
string
</em>
</td>
<td>
<p>DnsServiceIP: An IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address
range specified in serviceCidr.</p>
</td>
</tr>
<tr>
<td>
<code>dockerBridgeCidr</code><br/>
<em>
string
</em>
</td>
<td>
<p>DockerBridgeCidr: A CIDR notation IP range assigned to the Docker bridge network. It must not overlap with any Subnet IP
ranges or the Kubernetes service address range.</p>
</td>
</tr>
<tr>
<td>
<code>ipFamilies</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_IpFamilies_STATUS">
[]ContainerServiceNetworkProfile_IpFamilies_STATUS
</a>
</em>
</td>
<td>
<p>IpFamilies: IP families are used to determine single-stack or dual-stack clusters. For single-stack, the expected value
is IPv4. For dual-stack, the expected values are IPv4 and IPv6.</p>
</td>
</tr>
<tr>
<td>
<code>kubeProxyConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_STATUS">
ContainerServiceNetworkProfile_KubeProxyConfig_STATUS
</a>
</em>
</td>
<td>
<p>KubeProxyConfig: Holds configuration customizations for kube-proxy. Any values not defined will use the kube-proxy
defaulting behavior. See <a href="https://v">https://v</a><version>.docs.kubernetes.io/docs/reference/command-line-tools-reference/kube-proxy/
where <version> is represented by a <major version>-<minor version> string. Kubernetes version 1.23 would be &lsquo;1-23&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_STATUS">
ManagedClusterLoadBalancerProfile_STATUS
</a>
</em>
</td>
<td>
<p>LoadBalancerProfile: Profile of the cluster load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerSku</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.LoadBalancerSku_STATUS">
LoadBalancerSku_STATUS
</a>
</em>
</td>
<td>
<p>LoadBalancerSku: The default is &lsquo;standard&rsquo;. See <a href="https://docs.microsoft.com/azure/load-balancer/skus">Azure Load Balancer
SKUs</a> for more information about the differences between load
balancer SKUs.</p>
</td>
</tr>
<tr>
<td>
<code>natGatewayProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterNATGatewayProfile_STATUS">
ManagedClusterNATGatewayProfile_STATUS
</a>
</em>
</td>
<td>
<p>NatGatewayProfile: Profile of the cluster NAT gateway.</p>
</td>
</tr>
<tr>
<td>
<code>networkDataplane</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.NetworkDataplane_STATUS">
NetworkDataplane_STATUS
</a>
</em>
</td>
<td>
<p>NetworkDataplane: Network dataplane used in the Kubernetes cluster.</p>
</td>
</tr>
<tr>
<td>
<code>networkMode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.NetworkMode_STATUS">
NetworkMode_STATUS
</a>
</em>
</td>
<td>
<p>NetworkMode: This cannot be specified if networkPlugin is anything other than &lsquo;azure&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>networkPlugin</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.NetworkPlugin_STATUS">
NetworkPlugin_STATUS
</a>
</em>
</td>
<td>
<p>NetworkPlugin: Network plugin used for building the Kubernetes network.</p>
</td>
</tr>
<tr>
<td>
<code>networkPluginMode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.NetworkPluginMode_STATUS">
NetworkPluginMode_STATUS
</a>
</em>
</td>
<td>
<p>NetworkPluginMode: Network plugin mode used for building the Kubernetes network.</p>
</td>
</tr>
<tr>
<td>
<code>networkPolicy</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.NetworkPolicy_STATUS">
NetworkPolicy_STATUS
</a>
</em>
</td>
<td>
<p>NetworkPolicy: Network policy used for building the Kubernetes network.</p>
</td>
</tr>
<tr>
<td>
<code>outboundType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_OutboundType_STATUS">
ContainerServiceNetworkProfile_OutboundType_STATUS
</a>
</em>
</td>
<td>
<p>OutboundType: This can only be set at cluster creation time and cannot be changed later. For more information see
<a href="https://docs.microsoft.com/azure/aks/egress-outboundtype">egress outbound type</a>.</p>
</td>
</tr>
<tr>
<td>
<code>podCidr</code><br/>
<em>
string
</em>
</td>
<td>
<p>PodCidr: A CIDR notation IP range from which to assign pod IPs when kubenet is used.</p>
</td>
</tr>
<tr>
<td>
<code>podCidrs</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>PodCidrs: One IPv4 CIDR is expected for single-stack networking. Two CIDRs, one for each IP family (IPv4/IPv6), is
expected for dual-stack networking.</p>
</td>
</tr>
<tr>
<td>
<code>serviceCidr</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServiceCidr: A CIDR notation IP range from which to assign service cluster IPs. It must not overlap with any Subnet IP
ranges.</p>
</td>
</tr>
<tr>
<td>
<code>serviceCidrs</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>ServiceCidrs: One IPv4 CIDR is expected for single-stack networking. Two CIDRs, one for each IP family (IPv4/IPv6), is
expected for dual-stack networking. They must not overlap with any Subnet IP ranges.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_STATUS_ARM">ContainerServiceNetworkProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>)
</p>
<div>
<p>Profile of network configuration.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dnsServiceIP</code><br/>
<em>
string
</em>
</td>
<td>
<p>DnsServiceIP: An IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address
range specified in serviceCidr.</p>
</td>
</tr>
<tr>
<td>
<code>dockerBridgeCidr</code><br/>
<em>
string
</em>
</td>
<td>
<p>DockerBridgeCidr: A CIDR notation IP range assigned to the Docker bridge network. It must not overlap with any Subnet IP
ranges or the Kubernetes service address range.</p>
</td>
</tr>
<tr>
<td>
<code>ipFamilies</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_IpFamilies_STATUS">
[]ContainerServiceNetworkProfile_IpFamilies_STATUS
</a>
</em>
</td>
<td>
<p>IpFamilies: IP families are used to determine single-stack or dual-stack clusters. For single-stack, the expected value
is IPv4. For dual-stack, the expected values are IPv4 and IPv6.</p>
</td>
</tr>
<tr>
<td>
<code>kubeProxyConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_KubeProxyConfig_STATUS_ARM">
ContainerServiceNetworkProfile_KubeProxyConfig_STATUS_ARM
</a>
</em>
</td>
<td>
<p>KubeProxyConfig: Holds configuration customizations for kube-proxy. Any values not defined will use the kube-proxy
defaulting behavior. See <a href="https://v">https://v</a><version>.docs.kubernetes.io/docs/reference/command-line-tools-reference/kube-proxy/
where <version> is represented by a <major version>-<minor version> string. Kubernetes version 1.23 would be &lsquo;1-23&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_STATUS_ARM">
ManagedClusterLoadBalancerProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>LoadBalancerProfile: Profile of the cluster load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerSku</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.LoadBalancerSku_STATUS">
LoadBalancerSku_STATUS
</a>
</em>
</td>
<td>
<p>LoadBalancerSku: The default is &lsquo;standard&rsquo;. See <a href="https://docs.microsoft.com/azure/load-balancer/skus">Azure Load Balancer
SKUs</a> for more information about the differences between load
balancer SKUs.</p>
</td>
</tr>
<tr>
<td>
<code>natGatewayProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterNATGatewayProfile_STATUS_ARM">
ManagedClusterNATGatewayProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>NatGatewayProfile: Profile of the cluster NAT gateway.</p>
</td>
</tr>
<tr>
<td>
<code>networkDataplane</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.NetworkDataplane_STATUS">
NetworkDataplane_STATUS
</a>
</em>
</td>
<td>
<p>NetworkDataplane: Network dataplane used in the Kubernetes cluster.</p>
</td>
</tr>
<tr>
<td>
<code>networkMode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.NetworkMode_STATUS">
NetworkMode_STATUS
</a>
</em>
</td>
<td>
<p>NetworkMode: This cannot be specified if networkPlugin is anything other than &lsquo;azure&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>networkPlugin</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.NetworkPlugin_STATUS">
NetworkPlugin_STATUS
</a>
</em>
</td>
<td>
<p>NetworkPlugin: Network plugin used for building the Kubernetes network.</p>
</td>
</tr>
<tr>
<td>
<code>networkPluginMode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.NetworkPluginMode_STATUS">
NetworkPluginMode_STATUS
</a>
</em>
</td>
<td>
<p>NetworkPluginMode: Network plugin mode used for building the Kubernetes network.</p>
</td>
</tr>
<tr>
<td>
<code>networkPolicy</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.NetworkPolicy_STATUS">
NetworkPolicy_STATUS
</a>
</em>
</td>
<td>
<p>NetworkPolicy: Network policy used for building the Kubernetes network.</p>
</td>
</tr>
<tr>
<td>
<code>outboundType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_OutboundType_STATUS">
ContainerServiceNetworkProfile_OutboundType_STATUS
</a>
</em>
</td>
<td>
<p>OutboundType: This can only be set at cluster creation time and cannot be changed later. For more information see
<a href="https://docs.microsoft.com/azure/aks/egress-outboundtype">egress outbound type</a>.</p>
</td>
</tr>
<tr>
<td>
<code>podCidr</code><br/>
<em>
string
</em>
</td>
<td>
<p>PodCidr: A CIDR notation IP range from which to assign pod IPs when kubenet is used.</p>
</td>
</tr>
<tr>
<td>
<code>podCidrs</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>PodCidrs: One IPv4 CIDR is expected for single-stack networking. Two CIDRs, one for each IP family (IPv4/IPv6), is
expected for dual-stack networking.</p>
</td>
</tr>
<tr>
<td>
<code>serviceCidr</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServiceCidr: A CIDR notation IP range from which to assign service cluster IPs. It must not overlap with any Subnet IP
ranges.</p>
</td>
</tr>
<tr>
<td>
<code>serviceCidrs</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>ServiceCidrs: One IPv4 CIDR is expected for single-stack networking. Two CIDRs, one for each IP family (IPv4/IPv6), is
expected for dual-stack networking. They must not overlap with any Subnet IP ranges.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceOSDisk">ContainerServiceOSDisk
(<code>int</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_ARM">ManagedClusterAgentPoolProfileProperties_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_ARM">ManagedClusterAgentPoolProfile_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_Spec">ManagedClusters_AgentPool_Spec</a>)
</p>
<div>
</div>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceSshConfiguration">ContainerServiceSshConfiguration
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceLinuxProfile">ContainerServiceLinuxProfile</a>)
</p>
<div>
<p>SSH configuration for Linux-based VMs running on Azure.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>publicKeys</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceSshPublicKey">
[]ContainerServiceSshPublicKey
</a>
</em>
</td>
<td>
<p>PublicKeys: The list of SSH public keys used to authenticate with Linux-based VMs. A maximum of 1 key may be specified.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceSshConfiguration_ARM">ContainerServiceSshConfiguration_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceLinuxProfile_ARM">ContainerServiceLinuxProfile_ARM</a>)
</p>
<div>
<p>SSH configuration for Linux-based VMs running on Azure.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>publicKeys</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceSshPublicKey_ARM">
[]ContainerServiceSshPublicKey_ARM
</a>
</em>
</td>
<td>
<p>PublicKeys: The list of SSH public keys used to authenticate with Linux-based VMs. A maximum of 1 key may be specified.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceSshConfiguration_STATUS">ContainerServiceSshConfiguration_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceLinuxProfile_STATUS">ContainerServiceLinuxProfile_STATUS</a>)
</p>
<div>
<p>SSH configuration for Linux-based VMs running on Azure.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>publicKeys</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceSshPublicKey_STATUS">
[]ContainerServiceSshPublicKey_STATUS
</a>
</em>
</td>
<td>
<p>PublicKeys: The list of SSH public keys used to authenticate with Linux-based VMs. A maximum of 1 key may be specified.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceSshConfiguration_STATUS_ARM">ContainerServiceSshConfiguration_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceLinuxProfile_STATUS_ARM">ContainerServiceLinuxProfile_STATUS_ARM</a>)
</p>
<div>
<p>SSH configuration for Linux-based VMs running on Azure.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>publicKeys</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceSshPublicKey_STATUS_ARM">
[]ContainerServiceSshPublicKey_STATUS_ARM
</a>
</em>
</td>
<td>
<p>PublicKeys: The list of SSH public keys used to authenticate with Linux-based VMs. A maximum of 1 key may be specified.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceSshPublicKey">ContainerServiceSshPublicKey
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceSshConfiguration">ContainerServiceSshConfiguration</a>)
</p>
<div>
<p>Contains information about SSH certificate public key data.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>keyData</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyData: Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or
without headers.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceSshPublicKey_ARM">ContainerServiceSshPublicKey_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceSshConfiguration_ARM">ContainerServiceSshConfiguration_ARM</a>)
</p>
<div>
<p>Contains information about SSH certificate public key data.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>keyData</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyData: Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or
without headers.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceSshPublicKey_STATUS">ContainerServiceSshPublicKey_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceSshConfiguration_STATUS">ContainerServiceSshConfiguration_STATUS</a>)
</p>
<div>
<p>Contains information about SSH certificate public key data.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>keyData</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyData: Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or
without headers.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ContainerServiceSshPublicKey_STATUS_ARM">ContainerServiceSshPublicKey_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceSshConfiguration_STATUS_ARM">ContainerServiceSshConfiguration_STATUS_ARM</a>)
</p>
<div>
<p>Contains information about SSH certificate public key data.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>keyData</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyData: Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or
without headers.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ControlPlaneUpgradeOverride">ControlPlaneUpgradeOverride
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.UpgradeOverrideSettings">UpgradeOverrideSettings</a>, <a href="#containerservice.azure.com/v1api20230202preview.UpgradeOverrideSettings_ARM">UpgradeOverrideSettings_ARM</a>)
</p>
<div>
<p>The list of control plane upgrade override settings.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;IgnoreKubernetesDeprecations&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ControlPlaneUpgradeOverride_STATUS">ControlPlaneUpgradeOverride_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.UpgradeOverrideSettings_STATUS">UpgradeOverrideSettings_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.UpgradeOverrideSettings_STATUS_ARM">UpgradeOverrideSettings_STATUS_ARM</a>)
</p>
<div>
<p>The list of control plane upgrade override settings.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;IgnoreKubernetesDeprecations&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.CreationData">CreationData
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_Spec">ManagedClusters_AgentPool_Spec</a>)
</p>
<div>
<p>Data used when creating a target resource from a source resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>sourceResourceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>SourceResourceReference: This is the ARM ID of the source object to be used to create the target object.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.CreationData_ARM">CreationData_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_ARM">ManagedClusterAgentPoolProfileProperties_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_ARM">ManagedClusterAgentPoolProfile_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM</a>)
</p>
<div>
<p>Data used when creating a target resource from a source resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>sourceResourceId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.CreationData_STATUS">CreationData_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS">ManagedClusterAgentPoolProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_STATUS">ManagedClusters_AgentPool_STATUS</a>)
</p>
<div>
<p>Data used when creating a target resource from a source resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>sourceResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourceResourceId: This is the ARM ID of the source object to be used to create the target object.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.CreationData_STATUS_ARM">CreationData_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_STATUS_ARM">ManagedClusterAgentPoolProfileProperties_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS_ARM">ManagedClusterAgentPoolProfile_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>)
</p>
<div>
<p>Data used when creating a target resource from a source resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>sourceResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourceResourceId: This is the ARM ID of the source object to be used to create the target object.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ExtendedLocation">ExtendedLocation
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
</p>
<div>
<p>The complex type of the extended location.</p>
</div>
<table>
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
<p>Name: The name of the extended location.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ExtendedLocationType">
ExtendedLocationType
</a>
</em>
</td>
<td>
<p>Type: The type of the extended location.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ExtendedLocationType">ExtendedLocationType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ExtendedLocation">ExtendedLocation</a>, <a href="#containerservice.azure.com/v1api20230202preview.ExtendedLocation_ARM">ExtendedLocation_ARM</a>)
</p>
<div>
<p>The type of extendedLocation.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;EdgeZone&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ExtendedLocationType_STATUS">ExtendedLocationType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ExtendedLocation_STATUS">ExtendedLocation_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ExtendedLocation_STATUS_ARM">ExtendedLocation_STATUS_ARM</a>)
</p>
<div>
<p>The type of extendedLocation.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;EdgeZone&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ExtendedLocation_ARM">ExtendedLocation_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec_ARM">ManagedCluster_Spec_ARM</a>)
</p>
<div>
<p>The complex type of the extended location.</p>
</div>
<table>
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
<p>Name: The name of the extended location.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ExtendedLocationType">
ExtendedLocationType
</a>
</em>
</td>
<td>
<p>Type: The type of the extended location.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ExtendedLocation_STATUS">ExtendedLocation_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
</p>
<div>
<p>The complex type of the extended location.</p>
</div>
<table>
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
<p>Name: The name of the extended location.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ExtendedLocationType_STATUS">
ExtendedLocationType_STATUS
</a>
</em>
</td>
<td>
<p>Type: The type of the extended location.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ExtendedLocation_STATUS_ARM">ExtendedLocation_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS_ARM">ManagedCluster_STATUS_ARM</a>)
</p>
<div>
<p>The complex type of the extended location.</p>
</div>
<table>
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
<p>Name: The name of the extended location.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ExtendedLocationType_STATUS">
ExtendedLocationType_STATUS
</a>
</em>
</td>
<td>
<p>Type: The type of the extended location.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.GPUInstanceProfile">GPUInstanceProfile
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_ARM">ManagedClusterAgentPoolProfileProperties_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_ARM">ManagedClusterAgentPoolProfile_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_Spec">ManagedClusters_AgentPool_Spec</a>)
</p>
<div>
<p>GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;MIG1g&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;MIG2g&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;MIG3g&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;MIG4g&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;MIG7g&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.GPUInstanceProfile_STATUS">GPUInstanceProfile_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_STATUS_ARM">ManagedClusterAgentPoolProfileProperties_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS">ManagedClusterAgentPoolProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS_ARM">ManagedClusterAgentPoolProfile_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_STATUS">ManagedClusters_AgentPool_STATUS</a>)
</p>
<div>
<p>GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;MIG1g&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;MIG2g&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;MIG3g&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;MIG4g&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;MIG7g&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.GuardrailsProfile">GuardrailsProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
</p>
<div>
<p>The Guardrails profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>excludedNamespaces</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>ExcludedNamespaces: List of namespaces excluded from guardrails checks</p>
</td>
</tr>
<tr>
<td>
<code>level</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.GuardrailsProfile_Level">
GuardrailsProfile_Level
</a>
</em>
</td>
<td>
<p>Level: The guardrails level to be used. By default, Guardrails is enabled for all namespaces except those that AKS
excludes via systemExcludedNamespaces</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br/>
<em>
string
</em>
</td>
<td>
<p>Version: The version of constraints to use</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.GuardrailsProfile_ARM">GuardrailsProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM</a>)
</p>
<div>
<p>The Guardrails profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>excludedNamespaces</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>ExcludedNamespaces: List of namespaces excluded from guardrails checks</p>
</td>
</tr>
<tr>
<td>
<code>level</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.GuardrailsProfile_Level">
GuardrailsProfile_Level
</a>
</em>
</td>
<td>
<p>Level: The guardrails level to be used. By default, Guardrails is enabled for all namespaces except those that AKS
excludes via systemExcludedNamespaces</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br/>
<em>
string
</em>
</td>
<td>
<p>Version: The version of constraints to use</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.GuardrailsProfile_Level">GuardrailsProfile_Level
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.GuardrailsProfile">GuardrailsProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.GuardrailsProfile_ARM">GuardrailsProfile_ARM</a>)
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
<tbody><tr><td><p>&#34;Enforcement&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Off&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Warning&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.GuardrailsProfile_Level_STATUS">GuardrailsProfile_Level_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.GuardrailsProfile_STATUS">GuardrailsProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.GuardrailsProfile_STATUS_ARM">GuardrailsProfile_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;Enforcement&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Off&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Warning&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.GuardrailsProfile_STATUS">GuardrailsProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
</p>
<div>
<p>The Guardrails profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>excludedNamespaces</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>ExcludedNamespaces: List of namespaces excluded from guardrails checks</p>
</td>
</tr>
<tr>
<td>
<code>level</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.GuardrailsProfile_Level_STATUS">
GuardrailsProfile_Level_STATUS
</a>
</em>
</td>
<td>
<p>Level: The guardrails level to be used. By default, Guardrails is enabled for all namespaces except those that AKS
excludes via systemExcludedNamespaces</p>
</td>
</tr>
<tr>
<td>
<code>systemExcludedNamespaces</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>SystemExcludedNamespaces: List of namespaces specified by AKS to be excluded from Guardrails</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br/>
<em>
string
</em>
</td>
<td>
<p>Version: The version of constraints to use</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.GuardrailsProfile_STATUS_ARM">GuardrailsProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>)
</p>
<div>
<p>The Guardrails profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>excludedNamespaces</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>ExcludedNamespaces: List of namespaces excluded from guardrails checks</p>
</td>
</tr>
<tr>
<td>
<code>level</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.GuardrailsProfile_Level_STATUS">
GuardrailsProfile_Level_STATUS
</a>
</em>
</td>
<td>
<p>Level: The guardrails level to be used. By default, Guardrails is enabled for all namespaces except those that AKS
excludes via systemExcludedNamespaces</p>
</td>
</tr>
<tr>
<td>
<code>systemExcludedNamespaces</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>SystemExcludedNamespaces: List of namespaces specified by AKS to be excluded from Guardrails</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br/>
<em>
string
</em>
</td>
<td>
<p>Version: The version of constraints to use</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.IPTag">IPTag
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.AgentPoolNetworkProfile">AgentPoolNetworkProfile</a>)
</p>
<div>
<p>Contains the IPTag associated with the object.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ipTagType</code><br/>
<em>
string
</em>
</td>
<td>
<p>IpTagType: The IP tag type. Example: RoutingPreference.</p>
</td>
</tr>
<tr>
<td>
<code>tag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Tag: The value of the IP tag associated with the public IP. Example: Internet.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.IPTag_ARM">IPTag_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.AgentPoolNetworkProfile_ARM">AgentPoolNetworkProfile_ARM</a>)
</p>
<div>
<p>Contains the IPTag associated with the object.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ipTagType</code><br/>
<em>
string
</em>
</td>
<td>
<p>IpTagType: The IP tag type. Example: RoutingPreference.</p>
</td>
</tr>
<tr>
<td>
<code>tag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Tag: The value of the IP tag associated with the public IP. Example: Internet.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.IPTag_STATUS">IPTag_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.AgentPoolNetworkProfile_STATUS">AgentPoolNetworkProfile_STATUS</a>)
</p>
<div>
<p>Contains the IPTag associated with the object.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ipTagType</code><br/>
<em>
string
</em>
</td>
<td>
<p>IpTagType: The IP tag type. Example: RoutingPreference.</p>
</td>
</tr>
<tr>
<td>
<code>tag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Tag: The value of the IP tag associated with the public IP. Example: Internet.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.IPTag_STATUS_ARM">IPTag_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.AgentPoolNetworkProfile_STATUS_ARM">AgentPoolNetworkProfile_STATUS_ARM</a>)
</p>
<div>
<p>Contains the IPTag associated with the object.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ipTagType</code><br/>
<em>
string
</em>
</td>
<td>
<p>IpTagType: The IP tag type. Example: RoutingPreference.</p>
</td>
</tr>
<tr>
<td>
<code>tag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Tag: The value of the IP tag associated with the public IP. Example: Internet.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.IstioComponents">IstioComponents
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.IstioServiceMesh">IstioServiceMesh</a>)
</p>
<div>
<p>Istio components configuration.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ingressGateways</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.IstioIngressGateway">
[]IstioIngressGateway
</a>
</em>
</td>
<td>
<p>IngressGateways: Istio ingress gateways.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.IstioComponents_ARM">IstioComponents_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.IstioServiceMesh_ARM">IstioServiceMesh_ARM</a>)
</p>
<div>
<p>Istio components configuration.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ingressGateways</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.IstioIngressGateway_ARM">
[]IstioIngressGateway_ARM
</a>
</em>
</td>
<td>
<p>IngressGateways: Istio ingress gateways.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.IstioComponents_STATUS">IstioComponents_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.IstioServiceMesh_STATUS">IstioServiceMesh_STATUS</a>)
</p>
<div>
<p>Istio components configuration.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ingressGateways</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.IstioIngressGateway_STATUS">
[]IstioIngressGateway_STATUS
</a>
</em>
</td>
<td>
<p>IngressGateways: Istio ingress gateways.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.IstioComponents_STATUS_ARM">IstioComponents_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.IstioServiceMesh_STATUS_ARM">IstioServiceMesh_STATUS_ARM</a>)
</p>
<div>
<p>Istio components configuration.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ingressGateways</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.IstioIngressGateway_STATUS_ARM">
[]IstioIngressGateway_STATUS_ARM
</a>
</em>
</td>
<td>
<p>IngressGateways: Istio ingress gateways.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.IstioIngressGateway">IstioIngressGateway
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.IstioComponents">IstioComponents</a>)
</p>
<div>
<p>Istio ingress gateway configuration. For now, we support up to one external ingress gateway named
<code>aks-istio-ingressgateway-external</code> and one internal ingress gateway named <code>aks-istio-ingressgateway-internal</code>.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable the ingress gateway.</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.IstioIngressGateway_Mode">
IstioIngressGateway_Mode
</a>
</em>
</td>
<td>
<p>Mode: Mode of an ingress gateway.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.IstioIngressGateway_ARM">IstioIngressGateway_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.IstioComponents_ARM">IstioComponents_ARM</a>)
</p>
<div>
<p>Istio ingress gateway configuration. For now, we support up to one external ingress gateway named
<code>aks-istio-ingressgateway-external</code> and one internal ingress gateway named <code>aks-istio-ingressgateway-internal</code>.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable the ingress gateway.</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.IstioIngressGateway_Mode">
IstioIngressGateway_Mode
</a>
</em>
</td>
<td>
<p>Mode: Mode of an ingress gateway.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.IstioIngressGateway_Mode">IstioIngressGateway_Mode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.IstioIngressGateway">IstioIngressGateway</a>, <a href="#containerservice.azure.com/v1api20230202preview.IstioIngressGateway_ARM">IstioIngressGateway_ARM</a>)
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
<tbody><tr><td><p>&#34;External&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Internal&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.IstioIngressGateway_Mode_STATUS">IstioIngressGateway_Mode_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.IstioIngressGateway_STATUS">IstioIngressGateway_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.IstioIngressGateway_STATUS_ARM">IstioIngressGateway_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;External&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Internal&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.IstioIngressGateway_STATUS">IstioIngressGateway_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.IstioComponents_STATUS">IstioComponents_STATUS</a>)
</p>
<div>
<p>Istio ingress gateway configuration. For now, we support up to one external ingress gateway named
<code>aks-istio-ingressgateway-external</code> and one internal ingress gateway named <code>aks-istio-ingressgateway-internal</code>.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable the ingress gateway.</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.IstioIngressGateway_Mode_STATUS">
IstioIngressGateway_Mode_STATUS
</a>
</em>
</td>
<td>
<p>Mode: Mode of an ingress gateway.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.IstioIngressGateway_STATUS_ARM">IstioIngressGateway_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.IstioComponents_STATUS_ARM">IstioComponents_STATUS_ARM</a>)
</p>
<div>
<p>Istio ingress gateway configuration. For now, we support up to one external ingress gateway named
<code>aks-istio-ingressgateway-external</code> and one internal ingress gateway named <code>aks-istio-ingressgateway-internal</code>.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable the ingress gateway.</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.IstioIngressGateway_Mode_STATUS">
IstioIngressGateway_Mode_STATUS
</a>
</em>
</td>
<td>
<p>Mode: Mode of an ingress gateway.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.IstioServiceMesh">IstioServiceMesh
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ServiceMeshProfile">ServiceMeshProfile</a>)
</p>
<div>
<p>Istio service mesh configuration.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>components</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.IstioComponents">
IstioComponents
</a>
</em>
</td>
<td>
<p>Components: Istio components configuration.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.IstioServiceMesh_ARM">IstioServiceMesh_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ServiceMeshProfile_ARM">ServiceMeshProfile_ARM</a>)
</p>
<div>
<p>Istio service mesh configuration.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>components</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.IstioComponents_ARM">
IstioComponents_ARM
</a>
</em>
</td>
<td>
<p>Components: Istio components configuration.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.IstioServiceMesh_STATUS">IstioServiceMesh_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ServiceMeshProfile_STATUS">ServiceMeshProfile_STATUS</a>)
</p>
<div>
<p>Istio service mesh configuration.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>components</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.IstioComponents_STATUS">
IstioComponents_STATUS
</a>
</em>
</td>
<td>
<p>Components: Istio components configuration.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.IstioServiceMesh_STATUS_ARM">IstioServiceMesh_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ServiceMeshProfile_STATUS_ARM">ServiceMeshProfile_STATUS_ARM</a>)
</p>
<div>
<p>Istio service mesh configuration.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>components</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.IstioComponents_STATUS_ARM">
IstioComponents_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Components: Istio components configuration.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.KubeletConfig">KubeletConfig
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_Spec">ManagedClusters_AgentPool_Spec</a>)
</p>
<div>
<p>See <a href="https://docs.microsoft.com/azure/aks/custom-node-configuration">AKS custom node configuration</a> for more details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allowedUnsafeSysctls</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AllowedUnsafeSysctls: Allowed list of unsafe sysctls or unsafe sysctl patterns (ending in <code>*</code>).</p>
</td>
</tr>
<tr>
<td>
<code>containerLogMaxFiles</code><br/>
<em>
int
</em>
</td>
<td>
<p>ContainerLogMaxFiles: The maximum number of container log files that can be present for a container. The number must be
≥ 2.</p>
</td>
</tr>
<tr>
<td>
<code>containerLogMaxSizeMB</code><br/>
<em>
int
</em>
</td>
<td>
<p>ContainerLogMaxSizeMB: The maximum size (e.g. 10Mi) of container log file before it is rotated.</p>
</td>
</tr>
<tr>
<td>
<code>cpuCfsQuota</code><br/>
<em>
bool
</em>
</td>
<td>
<p>CpuCfsQuota: The default is true.</p>
</td>
</tr>
<tr>
<td>
<code>cpuCfsQuotaPeriod</code><br/>
<em>
string
</em>
</td>
<td>
<p>CpuCfsQuotaPeriod: The default is &lsquo;100ms.&rsquo; Valid values are a sequence of decimal numbers with an optional fraction and
a unit suffix. For example: &lsquo;300ms&rsquo;, &lsquo;2h45m&rsquo;. Supported units are &lsquo;ns&rsquo;, &lsquo;us&rsquo;, &lsquo;ms&rsquo;, &rsquo;s&rsquo;, &rsquo;m&rsquo;, and &lsquo;h&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>cpuManagerPolicy</code><br/>
<em>
string
</em>
</td>
<td>
<p>CpuManagerPolicy: The default is &lsquo;none&rsquo;. See <a href="https://kubernetes.io/docs/tasks/administer-cluster/cpu-management-policies/#cpu-management-policies">Kubernetes CPU management
policies</a> for more
information. Allowed values are &lsquo;none&rsquo; and &lsquo;static&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>failSwapOn</code><br/>
<em>
bool
</em>
</td>
<td>
<p>FailSwapOn: If set to true it will make the Kubelet fail to start if swap is enabled on the node.</p>
</td>
</tr>
<tr>
<td>
<code>imageGcHighThreshold</code><br/>
<em>
int
</em>
</td>
<td>
<p>ImageGcHighThreshold: To disable image garbage collection, set to 100. The default is 85%</p>
</td>
</tr>
<tr>
<td>
<code>imageGcLowThreshold</code><br/>
<em>
int
</em>
</td>
<td>
<p>ImageGcLowThreshold: This cannot be set higher than imageGcHighThreshold. The default is 80%</p>
</td>
</tr>
<tr>
<td>
<code>podMaxPids</code><br/>
<em>
int
</em>
</td>
<td>
<p>PodMaxPids: The maximum number of processes per pod.</p>
</td>
</tr>
<tr>
<td>
<code>topologyManagerPolicy</code><br/>
<em>
string
</em>
</td>
<td>
<p>TopologyManagerPolicy: For more information see <a href="https://kubernetes.io/docs/tasks/administer-cluster/topology-manager">Kubernetes Topology
Manager</a>. The default is &lsquo;none&rsquo;. Allowed values
are &lsquo;none&rsquo;, &lsquo;best-effort&rsquo;, &lsquo;restricted&rsquo;, and &lsquo;single-numa-node&rsquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.KubeletConfig_ARM">KubeletConfig_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_ARM">ManagedClusterAgentPoolProfileProperties_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_ARM">ManagedClusterAgentPoolProfile_ARM</a>)
</p>
<div>
<p>See <a href="https://docs.microsoft.com/azure/aks/custom-node-configuration">AKS custom node configuration</a> for more details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allowedUnsafeSysctls</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AllowedUnsafeSysctls: Allowed list of unsafe sysctls or unsafe sysctl patterns (ending in <code>*</code>).</p>
</td>
</tr>
<tr>
<td>
<code>containerLogMaxFiles</code><br/>
<em>
int
</em>
</td>
<td>
<p>ContainerLogMaxFiles: The maximum number of container log files that can be present for a container. The number must be
≥ 2.</p>
</td>
</tr>
<tr>
<td>
<code>containerLogMaxSizeMB</code><br/>
<em>
int
</em>
</td>
<td>
<p>ContainerLogMaxSizeMB: The maximum size (e.g. 10Mi) of container log file before it is rotated.</p>
</td>
</tr>
<tr>
<td>
<code>cpuCfsQuota</code><br/>
<em>
bool
</em>
</td>
<td>
<p>CpuCfsQuota: The default is true.</p>
</td>
</tr>
<tr>
<td>
<code>cpuCfsQuotaPeriod</code><br/>
<em>
string
</em>
</td>
<td>
<p>CpuCfsQuotaPeriod: The default is &lsquo;100ms.&rsquo; Valid values are a sequence of decimal numbers with an optional fraction and
a unit suffix. For example: &lsquo;300ms&rsquo;, &lsquo;2h45m&rsquo;. Supported units are &lsquo;ns&rsquo;, &lsquo;us&rsquo;, &lsquo;ms&rsquo;, &rsquo;s&rsquo;, &rsquo;m&rsquo;, and &lsquo;h&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>cpuManagerPolicy</code><br/>
<em>
string
</em>
</td>
<td>
<p>CpuManagerPolicy: The default is &lsquo;none&rsquo;. See <a href="https://kubernetes.io/docs/tasks/administer-cluster/cpu-management-policies/#cpu-management-policies">Kubernetes CPU management
policies</a> for more
information. Allowed values are &lsquo;none&rsquo; and &lsquo;static&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>failSwapOn</code><br/>
<em>
bool
</em>
</td>
<td>
<p>FailSwapOn: If set to true it will make the Kubelet fail to start if swap is enabled on the node.</p>
</td>
</tr>
<tr>
<td>
<code>imageGcHighThreshold</code><br/>
<em>
int
</em>
</td>
<td>
<p>ImageGcHighThreshold: To disable image garbage collection, set to 100. The default is 85%</p>
</td>
</tr>
<tr>
<td>
<code>imageGcLowThreshold</code><br/>
<em>
int
</em>
</td>
<td>
<p>ImageGcLowThreshold: This cannot be set higher than imageGcHighThreshold. The default is 80%</p>
</td>
</tr>
<tr>
<td>
<code>podMaxPids</code><br/>
<em>
int
</em>
</td>
<td>
<p>PodMaxPids: The maximum number of processes per pod.</p>
</td>
</tr>
<tr>
<td>
<code>topologyManagerPolicy</code><br/>
<em>
string
</em>
</td>
<td>
<p>TopologyManagerPolicy: For more information see <a href="https://kubernetes.io/docs/tasks/administer-cluster/topology-manager">Kubernetes Topology
Manager</a>. The default is &lsquo;none&rsquo;. Allowed values
are &lsquo;none&rsquo;, &lsquo;best-effort&rsquo;, &lsquo;restricted&rsquo;, and &lsquo;single-numa-node&rsquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.KubeletConfig_STATUS">KubeletConfig_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS">ManagedClusterAgentPoolProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_STATUS">ManagedClusters_AgentPool_STATUS</a>)
</p>
<div>
<p>See <a href="https://docs.microsoft.com/azure/aks/custom-node-configuration">AKS custom node configuration</a> for more details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allowedUnsafeSysctls</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AllowedUnsafeSysctls: Allowed list of unsafe sysctls or unsafe sysctl patterns (ending in <code>*</code>).</p>
</td>
</tr>
<tr>
<td>
<code>containerLogMaxFiles</code><br/>
<em>
int
</em>
</td>
<td>
<p>ContainerLogMaxFiles: The maximum number of container log files that can be present for a container. The number must be
≥ 2.</p>
</td>
</tr>
<tr>
<td>
<code>containerLogMaxSizeMB</code><br/>
<em>
int
</em>
</td>
<td>
<p>ContainerLogMaxSizeMB: The maximum size (e.g. 10Mi) of container log file before it is rotated.</p>
</td>
</tr>
<tr>
<td>
<code>cpuCfsQuota</code><br/>
<em>
bool
</em>
</td>
<td>
<p>CpuCfsQuota: The default is true.</p>
</td>
</tr>
<tr>
<td>
<code>cpuCfsQuotaPeriod</code><br/>
<em>
string
</em>
</td>
<td>
<p>CpuCfsQuotaPeriod: The default is &lsquo;100ms.&rsquo; Valid values are a sequence of decimal numbers with an optional fraction and
a unit suffix. For example: &lsquo;300ms&rsquo;, &lsquo;2h45m&rsquo;. Supported units are &lsquo;ns&rsquo;, &lsquo;us&rsquo;, &lsquo;ms&rsquo;, &rsquo;s&rsquo;, &rsquo;m&rsquo;, and &lsquo;h&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>cpuManagerPolicy</code><br/>
<em>
string
</em>
</td>
<td>
<p>CpuManagerPolicy: The default is &lsquo;none&rsquo;. See <a href="https://kubernetes.io/docs/tasks/administer-cluster/cpu-management-policies/#cpu-management-policies">Kubernetes CPU management
policies</a> for more
information. Allowed values are &lsquo;none&rsquo; and &lsquo;static&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>failSwapOn</code><br/>
<em>
bool
</em>
</td>
<td>
<p>FailSwapOn: If set to true it will make the Kubelet fail to start if swap is enabled on the node.</p>
</td>
</tr>
<tr>
<td>
<code>imageGcHighThreshold</code><br/>
<em>
int
</em>
</td>
<td>
<p>ImageGcHighThreshold: To disable image garbage collection, set to 100. The default is 85%</p>
</td>
</tr>
<tr>
<td>
<code>imageGcLowThreshold</code><br/>
<em>
int
</em>
</td>
<td>
<p>ImageGcLowThreshold: This cannot be set higher than imageGcHighThreshold. The default is 80%</p>
</td>
</tr>
<tr>
<td>
<code>podMaxPids</code><br/>
<em>
int
</em>
</td>
<td>
<p>PodMaxPids: The maximum number of processes per pod.</p>
</td>
</tr>
<tr>
<td>
<code>topologyManagerPolicy</code><br/>
<em>
string
</em>
</td>
<td>
<p>TopologyManagerPolicy: For more information see <a href="https://kubernetes.io/docs/tasks/administer-cluster/topology-manager">Kubernetes Topology
Manager</a>. The default is &lsquo;none&rsquo;. Allowed values
are &lsquo;none&rsquo;, &lsquo;best-effort&rsquo;, &lsquo;restricted&rsquo;, and &lsquo;single-numa-node&rsquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.KubeletConfig_STATUS_ARM">KubeletConfig_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_STATUS_ARM">ManagedClusterAgentPoolProfileProperties_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS_ARM">ManagedClusterAgentPoolProfile_STATUS_ARM</a>)
</p>
<div>
<p>See <a href="https://docs.microsoft.com/azure/aks/custom-node-configuration">AKS custom node configuration</a> for more details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allowedUnsafeSysctls</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AllowedUnsafeSysctls: Allowed list of unsafe sysctls or unsafe sysctl patterns (ending in <code>*</code>).</p>
</td>
</tr>
<tr>
<td>
<code>containerLogMaxFiles</code><br/>
<em>
int
</em>
</td>
<td>
<p>ContainerLogMaxFiles: The maximum number of container log files that can be present for a container. The number must be
≥ 2.</p>
</td>
</tr>
<tr>
<td>
<code>containerLogMaxSizeMB</code><br/>
<em>
int
</em>
</td>
<td>
<p>ContainerLogMaxSizeMB: The maximum size (e.g. 10Mi) of container log file before it is rotated.</p>
</td>
</tr>
<tr>
<td>
<code>cpuCfsQuota</code><br/>
<em>
bool
</em>
</td>
<td>
<p>CpuCfsQuota: The default is true.</p>
</td>
</tr>
<tr>
<td>
<code>cpuCfsQuotaPeriod</code><br/>
<em>
string
</em>
</td>
<td>
<p>CpuCfsQuotaPeriod: The default is &lsquo;100ms.&rsquo; Valid values are a sequence of decimal numbers with an optional fraction and
a unit suffix. For example: &lsquo;300ms&rsquo;, &lsquo;2h45m&rsquo;. Supported units are &lsquo;ns&rsquo;, &lsquo;us&rsquo;, &lsquo;ms&rsquo;, &rsquo;s&rsquo;, &rsquo;m&rsquo;, and &lsquo;h&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>cpuManagerPolicy</code><br/>
<em>
string
</em>
</td>
<td>
<p>CpuManagerPolicy: The default is &lsquo;none&rsquo;. See <a href="https://kubernetes.io/docs/tasks/administer-cluster/cpu-management-policies/#cpu-management-policies">Kubernetes CPU management
policies</a> for more
information. Allowed values are &lsquo;none&rsquo; and &lsquo;static&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>failSwapOn</code><br/>
<em>
bool
</em>
</td>
<td>
<p>FailSwapOn: If set to true it will make the Kubelet fail to start if swap is enabled on the node.</p>
</td>
</tr>
<tr>
<td>
<code>imageGcHighThreshold</code><br/>
<em>
int
</em>
</td>
<td>
<p>ImageGcHighThreshold: To disable image garbage collection, set to 100. The default is 85%</p>
</td>
</tr>
<tr>
<td>
<code>imageGcLowThreshold</code><br/>
<em>
int
</em>
</td>
<td>
<p>ImageGcLowThreshold: This cannot be set higher than imageGcHighThreshold. The default is 80%</p>
</td>
</tr>
<tr>
<td>
<code>podMaxPids</code><br/>
<em>
int
</em>
</td>
<td>
<p>PodMaxPids: The maximum number of processes per pod.</p>
</td>
</tr>
<tr>
<td>
<code>topologyManagerPolicy</code><br/>
<em>
string
</em>
</td>
<td>
<p>TopologyManagerPolicy: For more information see <a href="https://kubernetes.io/docs/tasks/administer-cluster/topology-manager">Kubernetes Topology
Manager</a>. The default is &lsquo;none&rsquo;. Allowed values
are &lsquo;none&rsquo;, &lsquo;best-effort&rsquo;, &lsquo;restricted&rsquo;, and &lsquo;single-numa-node&rsquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.KubeletDiskType">KubeletDiskType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_ARM">ManagedClusterAgentPoolProfileProperties_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_ARM">ManagedClusterAgentPoolProfile_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_Spec">ManagedClusters_AgentPool_Spec</a>)
</p>
<div>
<p>Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;OS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Temporary&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.KubeletDiskType_STATUS">KubeletDiskType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_STATUS_ARM">ManagedClusterAgentPoolProfileProperties_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS">ManagedClusterAgentPoolProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS_ARM">ManagedClusterAgentPoolProfile_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_STATUS">ManagedClusters_AgentPool_STATUS</a>)
</p>
<div>
<p>Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral storage.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;OS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Temporary&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.LinuxOSConfig">LinuxOSConfig
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_Spec">ManagedClusters_AgentPool_Spec</a>)
</p>
<div>
<p>See <a href="https://docs.microsoft.com/azure/aks/custom-node-configuration">AKS custom node configuration</a> for more details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>swapFileSizeMB</code><br/>
<em>
int
</em>
</td>
<td>
<p>SwapFileSizeMB: The size in MB of a swap file that will be created on each node.</p>
</td>
</tr>
<tr>
<td>
<code>sysctls</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.SysctlConfig">
SysctlConfig
</a>
</em>
</td>
<td>
<p>Sysctls: Sysctl settings for Linux agent nodes.</p>
</td>
</tr>
<tr>
<td>
<code>transparentHugePageDefrag</code><br/>
<em>
string
</em>
</td>
<td>
<p>TransparentHugePageDefrag: Valid values are &lsquo;always&rsquo;, &lsquo;defer&rsquo;, &lsquo;defer+madvise&rsquo;, &lsquo;madvise&rsquo; and &lsquo;never&rsquo;. The default is
&lsquo;madvise&rsquo;. For more information see <a href="https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge">Transparent
Hugepages</a>.</p>
</td>
</tr>
<tr>
<td>
<code>transparentHugePageEnabled</code><br/>
<em>
string
</em>
</td>
<td>
<p>TransparentHugePageEnabled: Valid values are &lsquo;always&rsquo;, &lsquo;madvise&rsquo;, and &lsquo;never&rsquo;. The default is &lsquo;always&rsquo;. For more
information see <a href="https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge">Transparent
Hugepages</a>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.LinuxOSConfig_ARM">LinuxOSConfig_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_ARM">ManagedClusterAgentPoolProfileProperties_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_ARM">ManagedClusterAgentPoolProfile_ARM</a>)
</p>
<div>
<p>See <a href="https://docs.microsoft.com/azure/aks/custom-node-configuration">AKS custom node configuration</a> for more details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>swapFileSizeMB</code><br/>
<em>
int
</em>
</td>
<td>
<p>SwapFileSizeMB: The size in MB of a swap file that will be created on each node.</p>
</td>
</tr>
<tr>
<td>
<code>sysctls</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.SysctlConfig_ARM">
SysctlConfig_ARM
</a>
</em>
</td>
<td>
<p>Sysctls: Sysctl settings for Linux agent nodes.</p>
</td>
</tr>
<tr>
<td>
<code>transparentHugePageDefrag</code><br/>
<em>
string
</em>
</td>
<td>
<p>TransparentHugePageDefrag: Valid values are &lsquo;always&rsquo;, &lsquo;defer&rsquo;, &lsquo;defer+madvise&rsquo;, &lsquo;madvise&rsquo; and &lsquo;never&rsquo;. The default is
&lsquo;madvise&rsquo;. For more information see <a href="https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge">Transparent
Hugepages</a>.</p>
</td>
</tr>
<tr>
<td>
<code>transparentHugePageEnabled</code><br/>
<em>
string
</em>
</td>
<td>
<p>TransparentHugePageEnabled: Valid values are &lsquo;always&rsquo;, &lsquo;madvise&rsquo;, and &lsquo;never&rsquo;. The default is &lsquo;always&rsquo;. For more
information see <a href="https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge">Transparent
Hugepages</a>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.LinuxOSConfig_STATUS">LinuxOSConfig_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS">ManagedClusterAgentPoolProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_STATUS">ManagedClusters_AgentPool_STATUS</a>)
</p>
<div>
<p>See <a href="https://docs.microsoft.com/azure/aks/custom-node-configuration">AKS custom node configuration</a> for more details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>swapFileSizeMB</code><br/>
<em>
int
</em>
</td>
<td>
<p>SwapFileSizeMB: The size in MB of a swap file that will be created on each node.</p>
</td>
</tr>
<tr>
<td>
<code>sysctls</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.SysctlConfig_STATUS">
SysctlConfig_STATUS
</a>
</em>
</td>
<td>
<p>Sysctls: Sysctl settings for Linux agent nodes.</p>
</td>
</tr>
<tr>
<td>
<code>transparentHugePageDefrag</code><br/>
<em>
string
</em>
</td>
<td>
<p>TransparentHugePageDefrag: Valid values are &lsquo;always&rsquo;, &lsquo;defer&rsquo;, &lsquo;defer+madvise&rsquo;, &lsquo;madvise&rsquo; and &lsquo;never&rsquo;. The default is
&lsquo;madvise&rsquo;. For more information see <a href="https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge">Transparent
Hugepages</a>.</p>
</td>
</tr>
<tr>
<td>
<code>transparentHugePageEnabled</code><br/>
<em>
string
</em>
</td>
<td>
<p>TransparentHugePageEnabled: Valid values are &lsquo;always&rsquo;, &lsquo;madvise&rsquo;, and &lsquo;never&rsquo;. The default is &lsquo;always&rsquo;. For more
information see <a href="https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge">Transparent
Hugepages</a>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.LinuxOSConfig_STATUS_ARM">LinuxOSConfig_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_STATUS_ARM">ManagedClusterAgentPoolProfileProperties_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS_ARM">ManagedClusterAgentPoolProfile_STATUS_ARM</a>)
</p>
<div>
<p>See <a href="https://docs.microsoft.com/azure/aks/custom-node-configuration">AKS custom node configuration</a> for more details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>swapFileSizeMB</code><br/>
<em>
int
</em>
</td>
<td>
<p>SwapFileSizeMB: The size in MB of a swap file that will be created on each node.</p>
</td>
</tr>
<tr>
<td>
<code>sysctls</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.SysctlConfig_STATUS_ARM">
SysctlConfig_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Sysctls: Sysctl settings for Linux agent nodes.</p>
</td>
</tr>
<tr>
<td>
<code>transparentHugePageDefrag</code><br/>
<em>
string
</em>
</td>
<td>
<p>TransparentHugePageDefrag: Valid values are &lsquo;always&rsquo;, &lsquo;defer&rsquo;, &lsquo;defer+madvise&rsquo;, &lsquo;madvise&rsquo; and &lsquo;never&rsquo;. The default is
&lsquo;madvise&rsquo;. For more information see <a href="https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge">Transparent
Hugepages</a>.</p>
</td>
</tr>
<tr>
<td>
<code>transparentHugePageEnabled</code><br/>
<em>
string
</em>
</td>
<td>
<p>TransparentHugePageEnabled: Valid values are &lsquo;always&rsquo;, &lsquo;madvise&rsquo;, and &lsquo;never&rsquo;. The default is &lsquo;always&rsquo;. For more
information see <a href="https://www.kernel.org/doc/html/latest/admin-guide/mm/transhuge.html#admin-guide-transhuge">Transparent
Hugepages</a>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.LoadBalancerSku">LoadBalancerSku
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile">ContainerServiceNetworkProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_ARM">ContainerServiceNetworkProfile_ARM</a>)
</p>
<div>
<p>The default is &lsquo;standard&rsquo;. See <a href="https://docs.microsoft.com/azure/load-balancer/skus">Azure Load Balancer SKUs</a> for more
information about the differences between load balancer SKUs.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;basic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.LoadBalancerSku_STATUS">LoadBalancerSku_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_STATUS">ContainerServiceNetworkProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_STATUS_ARM">ContainerServiceNetworkProfile_STATUS_ARM</a>)
</p>
<div>
<p>The default is &lsquo;standard&rsquo;. See <a href="https://docs.microsoft.com/azure/load-balancer/skus">Azure Load Balancer SKUs</a> for more
information about the differences between load balancer SKUs.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;basic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedCluster">ManagedCluster
</h3>
<div>
<p>Generator information:
- Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-02-02-preview/managedClusters.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ContainerService/&#x200b;managedClusters/&#x200b;{resourceName}</&#x200b;p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">
ManagedCluster_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>aadProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAADProfile">
ManagedClusterAADProfile
</a>
</em>
</td>
<td>
<p>AadProfile: The Azure Active Directory configuration.</p>
</td>
</tr>
<tr>
<td>
<code>addonProfiles</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAddonProfile">
map[string]./api/containerservice/v1api20230202preview.ManagedClusterAddonProfile
</a>
</em>
</td>
<td>
<p>AddonProfiles: The profile of managed cluster add-on.</p>
</td>
</tr>
<tr>
<td>
<code>agentPoolProfiles</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile">
[]ManagedClusterAgentPoolProfile
</a>
</em>
</td>
<td>
<p>AgentPoolProfiles: The agent pool properties.</p>
</td>
</tr>
<tr>
<td>
<code>apiServerAccessProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAPIServerAccessProfile">
ManagedClusterAPIServerAccessProfile
</a>
</em>
</td>
<td>
<p>ApiServerAccessProfile: The access profile for managed cluster API server.</p>
</td>
</tr>
<tr>
<td>
<code>autoScalerProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_AutoScalerProfile">
ManagedClusterProperties_AutoScalerProfile
</a>
</em>
</td>
<td>
<p>AutoScalerProfile: Parameters to be applied to the cluster-autoscaler when enabled</p>
</td>
</tr>
<tr>
<td>
<code>autoUpgradeProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile">
ManagedClusterAutoUpgradeProfile
</a>
</em>
</td>
<td>
<p>AutoUpgradeProfile: The auto upgrade configuration.</p>
</td>
</tr>
<tr>
<td>
<code>azureMonitorProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfile">
ManagedClusterAzureMonitorProfile
</a>
</em>
</td>
<td>
<p>AzureMonitorProfile: Prometheus addon profile for the container service cluster</p>
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
<code>creationData</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.CreationData">
CreationData
</a>
</em>
</td>
<td>
<p>CreationData: CreationData to be used to specify the source Snapshot ID if the cluster will be created/upgraded using a
snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>disableLocalAccounts</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableLocalAccounts: If set to true, getting static credentials will be disabled for this cluster. This must only be
used on Managed Clusters that are AAD enabled. For more details see <a href="https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview">disable local
accounts</a>.</p>
</td>
</tr>
<tr>
<td>
<code>diskEncryptionSetReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>DiskEncryptionSetReference: This is of the form:
&lsquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Compute/&#x200b;diskEncryptionSets/&#x200b;{encryptionSetName}&rsquo;</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>dnsPrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>DnsPrefix: This cannot be updated once the Managed Cluster has been created.</p>
</td>
</tr>
<tr>
<td>
<code>enableNamespaceResources</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableNamespaceResources: The default value is false. It can be enabled/disabled on creation and updating of the managed
cluster. See <a href="https://aka.ms/NamespaceARMResource">https://aka.ms/NamespaceARMResource</a> for more details on Namespace as
a ARM Resource.</p>
</td>
</tr>
<tr>
<td>
<code>enablePodSecurityPolicy</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePodSecurityPolicy: (DEPRECATED) Whether to enable Kubernetes pod security policy (preview). PodSecurityPolicy was
deprecated in Kubernetes v1.21, and removed from Kubernetes in v1.25. Learn more at <a href="https://aka.ms/k8s/psp">https://aka.ms/k8s/psp</a> and
<a href="https://aka.ms/aks/psp">https://aka.ms/aks/psp</a>.</p>
</td>
</tr>
<tr>
<td>
<code>enableRBAC</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableRBAC: Whether to enable Kubernetes Role-Based Access Control.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ExtendedLocation">
ExtendedLocation
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the Virtual Machine.</p>
</td>
</tr>
<tr>
<td>
<code>fqdnSubdomain</code><br/>
<em>
string
</em>
</td>
<td>
<p>FqdnSubdomain: This cannot be updated once the Managed Cluster has been created.</p>
</td>
</tr>
<tr>
<td>
<code>guardrailsProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.GuardrailsProfile">
GuardrailsProfile
</a>
</em>
</td>
<td>
<p>GuardrailsProfile: The guardrails profile holds all the guardrails information for a given cluster</p>
</td>
</tr>
<tr>
<td>
<code>httpProxyConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterHTTPProxyConfig">
ManagedClusterHTTPProxyConfig
</a>
</em>
</td>
<td>
<p>HttpProxyConfig: Configurations for provisioning the cluster with HTTP proxy servers.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity">
ManagedClusterIdentity
</a>
</em>
</td>
<td>
<p>Identity: The identity of the managed cluster, if configured.</p>
</td>
</tr>
<tr>
<td>
<code>identityProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.UserAssignedIdentity">
map[string]./api/containerservice/v1api20230202preview.UserAssignedIdentity
</a>
</em>
</td>
<td>
<p>IdentityProfile: Identities associated with the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>ingressProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIngressProfile">
ManagedClusterIngressProfile
</a>
</em>
</td>
<td>
<p>IngressProfile: Ingress profile for the managed cluster.</p>
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
<p>KubernetesVersion: When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades
must be performed sequentially by major version number. For example, upgrades between 1.14.x -&gt; 1.15.x or 1.15.x -&gt;
1.16.x are allowed, however 1.14.x -&gt; 1.16.x is not allowed. See <a href="https://docs.microsoft.com/azure/aks/upgrade-cluster">upgrading an AKS
cluster</a> for more details.</p>
</td>
</tr>
<tr>
<td>
<code>linuxProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceLinuxProfile">
ContainerServiceLinuxProfile
</a>
</em>
</td>
<td>
<p>LinuxProfile: The profile for Linux VMs in the Managed Cluster.</p>
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
<code>networkProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile">
ContainerServiceNetworkProfile
</a>
</em>
</td>
<td>
<p>NetworkProfile: The network configuration profile.</p>
</td>
</tr>
<tr>
<td>
<code>nodeResourceGroup</code><br/>
<em>
string
</em>
</td>
<td>
<p>NodeResourceGroup: The name of the resource group containing agent pool nodes.</p>
</td>
</tr>
<tr>
<td>
<code>nodeResourceGroupProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterNodeResourceGroupProfile">
ManagedClusterNodeResourceGroupProfile
</a>
</em>
</td>
<td>
<p>NodeResourceGroupProfile: The node resource group configuration profile.</p>
</td>
</tr>
<tr>
<td>
<code>oidcIssuerProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterOIDCIssuerProfile">
ManagedClusterOIDCIssuerProfile
</a>
</em>
</td>
<td>
<p>OidcIssuerProfile: The OIDC issuer profile of the Managed Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>operatorSpec</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterOperatorSpec">
ManagedClusterOperatorSpec
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
<code>podIdentityProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProfile">
ManagedClusterPodIdentityProfile
</a>
</em>
</td>
<td>
<p>PodIdentityProfile: See <a href="https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity">use AAD pod identity</a> for more
details on AAD pod identity integration.</p>
</td>
</tr>
<tr>
<td>
<code>privateLinkResources</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PrivateLinkResource">
[]PrivateLinkResource
</a>
</em>
</td>
<td>
<p>PrivateLinkResources: Private link resources associated with the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_PublicNetworkAccess">
ManagedClusterProperties_PublicNetworkAccess
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Allow or deny public network access for AKS</p>
</td>
</tr>
<tr>
<td>
<code>securityProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile">
ManagedClusterSecurityProfile
</a>
</em>
</td>
<td>
<p>SecurityProfile: Security profile for the managed cluster.</p>
</td>
</tr>
<tr>
<td>
<code>serviceMeshProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ServiceMeshProfile">
ServiceMeshProfile
</a>
</em>
</td>
<td>
<p>ServiceMeshProfile: Service mesh profile for a managed cluster.</p>
</td>
</tr>
<tr>
<td>
<code>servicePrincipalProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterServicePrincipalProfile">
ManagedClusterServicePrincipalProfile
</a>
</em>
</td>
<td>
<p>ServicePrincipalProfile: Information about a service principal identity for the cluster to use for manipulating Azure
APIs.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSKU">
ManagedClusterSKU
</a>
</em>
</td>
<td>
<p>Sku: The managed cluster SKU.</p>
</td>
</tr>
<tr>
<td>
<code>storageProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfile">
ManagedClusterStorageProfile
</a>
</em>
</td>
<td>
<p>StorageProfile: Storage profile for the managed cluster.</p>
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
<code>upgradeSettings</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ClusterUpgradeSettings">
ClusterUpgradeSettings
</a>
</em>
</td>
<td>
<p>UpgradeSettings: Settings for upgrading a cluster.</p>
</td>
</tr>
<tr>
<td>
<code>windowsProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWindowsProfile">
ManagedClusterWindowsProfile
</a>
</em>
</td>
<td>
<p>WindowsProfile: The profile for Windows VMs in the Managed Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>workloadAutoScalerProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfile">
ManagedClusterWorkloadAutoScalerProfile
</a>
</em>
</td>
<td>
<p>WorkloadAutoScalerProfile: Workload Auto-scaler profile for the managed cluster.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">
ManagedCluster_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAADProfile">ManagedClusterAADProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
</p>
<div>
<p>For more details see <a href="https://docs.microsoft.com/azure/aks/managed-aad">managed AAD on AKS</a>.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>adminGroupObjectIDs</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AdminGroupObjectIDs: The list of AAD group object IDs that will have admin role of the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>clientAppID</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClientAppID: (DEPRECATED) The client AAD application ID. Learn more at <a href="https://aka.ms/aks/aad-legacy">https://aka.ms/aks/aad-legacy</a>.</p>
</td>
</tr>
<tr>
<td>
<code>enableAzureRBAC</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAzureRBAC: Whether to enable Azure RBAC for Kubernetes authorization.</p>
</td>
</tr>
<tr>
<td>
<code>managed</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Managed: Whether to enable managed AAD.</p>
</td>
</tr>
<tr>
<td>
<code>serverAppID</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServerAppID: (DEPRECATED) The server AAD application ID. Learn more at <a href="https://aka.ms/aks/aad-legacy">https://aka.ms/aks/aad-legacy</a>.</p>
</td>
</tr>
<tr>
<td>
<code>serverAppSecret</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServerAppSecret: (DEPRECATED) The server AAD application secret. Learn more at <a href="https://aka.ms/aks/aad-legacy">https://aka.ms/aks/aad-legacy</a>.</p>
</td>
</tr>
<tr>
<td>
<code>tenantID</code><br/>
<em>
string
</em>
</td>
<td>
<p>TenantID: The AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment
subscription.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAADProfile_ARM">ManagedClusterAADProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM</a>)
</p>
<div>
<p>For more details see <a href="https://docs.microsoft.com/azure/aks/managed-aad">managed AAD on AKS</a>.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>adminGroupObjectIDs</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AdminGroupObjectIDs: The list of AAD group object IDs that will have admin role of the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>clientAppID</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClientAppID: (DEPRECATED) The client AAD application ID. Learn more at <a href="https://aka.ms/aks/aad-legacy">https://aka.ms/aks/aad-legacy</a>.</p>
</td>
</tr>
<tr>
<td>
<code>enableAzureRBAC</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAzureRBAC: Whether to enable Azure RBAC for Kubernetes authorization.</p>
</td>
</tr>
<tr>
<td>
<code>managed</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Managed: Whether to enable managed AAD.</p>
</td>
</tr>
<tr>
<td>
<code>serverAppID</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServerAppID: (DEPRECATED) The server AAD application ID. Learn more at <a href="https://aka.ms/aks/aad-legacy">https://aka.ms/aks/aad-legacy</a>.</p>
</td>
</tr>
<tr>
<td>
<code>serverAppSecret</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServerAppSecret: (DEPRECATED) The server AAD application secret. Learn more at <a href="https://aka.ms/aks/aad-legacy">https://aka.ms/aks/aad-legacy</a>.</p>
</td>
</tr>
<tr>
<td>
<code>tenantID</code><br/>
<em>
string
</em>
</td>
<td>
<p>TenantID: The AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment
subscription.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAADProfile_STATUS">ManagedClusterAADProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
</p>
<div>
<p>For more details see <a href="https://docs.microsoft.com/azure/aks/managed-aad">managed AAD on AKS</a>.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>adminGroupObjectIDs</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AdminGroupObjectIDs: The list of AAD group object IDs that will have admin role of the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>clientAppID</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClientAppID: (DEPRECATED) The client AAD application ID. Learn more at <a href="https://aka.ms/aks/aad-legacy">https://aka.ms/aks/aad-legacy</a>.</p>
</td>
</tr>
<tr>
<td>
<code>enableAzureRBAC</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAzureRBAC: Whether to enable Azure RBAC for Kubernetes authorization.</p>
</td>
</tr>
<tr>
<td>
<code>managed</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Managed: Whether to enable managed AAD.</p>
</td>
</tr>
<tr>
<td>
<code>serverAppID</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServerAppID: (DEPRECATED) The server AAD application ID. Learn more at <a href="https://aka.ms/aks/aad-legacy">https://aka.ms/aks/aad-legacy</a>.</p>
</td>
</tr>
<tr>
<td>
<code>serverAppSecret</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServerAppSecret: (DEPRECATED) The server AAD application secret. Learn more at <a href="https://aka.ms/aks/aad-legacy">https://aka.ms/aks/aad-legacy</a>.</p>
</td>
</tr>
<tr>
<td>
<code>tenantID</code><br/>
<em>
string
</em>
</td>
<td>
<p>TenantID: The AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment
subscription.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAADProfile_STATUS_ARM">ManagedClusterAADProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>)
</p>
<div>
<p>For more details see <a href="https://docs.microsoft.com/azure/aks/managed-aad">managed AAD on AKS</a>.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>adminGroupObjectIDs</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AdminGroupObjectIDs: The list of AAD group object IDs that will have admin role of the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>clientAppID</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClientAppID: (DEPRECATED) The client AAD application ID. Learn more at <a href="https://aka.ms/aks/aad-legacy">https://aka.ms/aks/aad-legacy</a>.</p>
</td>
</tr>
<tr>
<td>
<code>enableAzureRBAC</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAzureRBAC: Whether to enable Azure RBAC for Kubernetes authorization.</p>
</td>
</tr>
<tr>
<td>
<code>managed</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Managed: Whether to enable managed AAD.</p>
</td>
</tr>
<tr>
<td>
<code>serverAppID</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServerAppID: (DEPRECATED) The server AAD application ID. Learn more at <a href="https://aka.ms/aks/aad-legacy">https://aka.ms/aks/aad-legacy</a>.</p>
</td>
</tr>
<tr>
<td>
<code>serverAppSecret</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServerAppSecret: (DEPRECATED) The server AAD application secret. Learn more at <a href="https://aka.ms/aks/aad-legacy">https://aka.ms/aks/aad-legacy</a>.</p>
</td>
</tr>
<tr>
<td>
<code>tenantID</code><br/>
<em>
string
</em>
</td>
<td>
<p>TenantID: The AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment
subscription.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAPIServerAccessProfile">ManagedClusterAPIServerAccessProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
</p>
<div>
<p>Access profile for managed cluster API server.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>authorizedIPRanges</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AuthorizedIPRanges: IP ranges are specified in CIDR format, e.g. 137.117.106.<sup>88</sup>&frasl;<sub>29</sub>. This feature is not compatible with
clusters that use Public IP Per Node, or clusters that are using a Basic Load Balancer. For more information see <a href="https://docs.microsoft.com/azure/aks/api-server-authorized-ip-ranges">API
server authorized IP ranges</a>.</p>
</td>
</tr>
<tr>
<td>
<code>disableRunCommand</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableRunCommand: Whether to disable run command for the cluster or not.</p>
</td>
</tr>
<tr>
<td>
<code>enablePrivateCluster</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePrivateCluster: For more details, see <a href="https://docs.microsoft.com/azure/aks/private-clusters">Creating a private AKS
cluster</a>.</p>
</td>
</tr>
<tr>
<td>
<code>enablePrivateClusterPublicFQDN</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePrivateClusterPublicFQDN: Whether to create additional public FQDN for private cluster or not.</p>
</td>
</tr>
<tr>
<td>
<code>enableVnetIntegration</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableVnetIntegration: Whether to enable apiserver vnet integration for the cluster or not.</p>
</td>
</tr>
<tr>
<td>
<code>privateDNSZone</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateDNSZone: The default is System. For more details see <a href="https://docs.microsoft.com/azure/aks/private-clusters#configure-private-dns-zone">configure private DNS
zone</a>. Allowed values are &lsquo;system&rsquo; and
&lsquo;none&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>subnetId</code><br/>
<em>
string
</em>
</td>
<td>
<p>SubnetId: It is required when: 1. creating a new cluster with BYO Vnet; 2. updating an existing cluster to enable
apiserver vnet integration.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAPIServerAccessProfile_ARM">ManagedClusterAPIServerAccessProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM</a>)
</p>
<div>
<p>Access profile for managed cluster API server.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>authorizedIPRanges</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AuthorizedIPRanges: IP ranges are specified in CIDR format, e.g. 137.117.106.<sup>88</sup>&frasl;<sub>29</sub>. This feature is not compatible with
clusters that use Public IP Per Node, or clusters that are using a Basic Load Balancer. For more information see <a href="https://docs.microsoft.com/azure/aks/api-server-authorized-ip-ranges">API
server authorized IP ranges</a>.</p>
</td>
</tr>
<tr>
<td>
<code>disableRunCommand</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableRunCommand: Whether to disable run command for the cluster or not.</p>
</td>
</tr>
<tr>
<td>
<code>enablePrivateCluster</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePrivateCluster: For more details, see <a href="https://docs.microsoft.com/azure/aks/private-clusters">Creating a private AKS
cluster</a>.</p>
</td>
</tr>
<tr>
<td>
<code>enablePrivateClusterPublicFQDN</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePrivateClusterPublicFQDN: Whether to create additional public FQDN for private cluster or not.</p>
</td>
</tr>
<tr>
<td>
<code>enableVnetIntegration</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableVnetIntegration: Whether to enable apiserver vnet integration for the cluster or not.</p>
</td>
</tr>
<tr>
<td>
<code>privateDNSZone</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateDNSZone: The default is System. For more details see <a href="https://docs.microsoft.com/azure/aks/private-clusters#configure-private-dns-zone">configure private DNS
zone</a>. Allowed values are &lsquo;system&rsquo; and
&lsquo;none&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>subnetId</code><br/>
<em>
string
</em>
</td>
<td>
<p>SubnetId: It is required when: 1. creating a new cluster with BYO Vnet; 2. updating an existing cluster to enable
apiserver vnet integration.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAPIServerAccessProfile_STATUS">ManagedClusterAPIServerAccessProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
</p>
<div>
<p>Access profile for managed cluster API server.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>authorizedIPRanges</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AuthorizedIPRanges: IP ranges are specified in CIDR format, e.g. 137.117.106.<sup>88</sup>&frasl;<sub>29</sub>. This feature is not compatible with
clusters that use Public IP Per Node, or clusters that are using a Basic Load Balancer. For more information see <a href="https://docs.microsoft.com/azure/aks/api-server-authorized-ip-ranges">API
server authorized IP ranges</a>.</p>
</td>
</tr>
<tr>
<td>
<code>disableRunCommand</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableRunCommand: Whether to disable run command for the cluster or not.</p>
</td>
</tr>
<tr>
<td>
<code>enablePrivateCluster</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePrivateCluster: For more details, see <a href="https://docs.microsoft.com/azure/aks/private-clusters">Creating a private AKS
cluster</a>.</p>
</td>
</tr>
<tr>
<td>
<code>enablePrivateClusterPublicFQDN</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePrivateClusterPublicFQDN: Whether to create additional public FQDN for private cluster or not.</p>
</td>
</tr>
<tr>
<td>
<code>enableVnetIntegration</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableVnetIntegration: Whether to enable apiserver vnet integration for the cluster or not.</p>
</td>
</tr>
<tr>
<td>
<code>privateDNSZone</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateDNSZone: The default is System. For more details see <a href="https://docs.microsoft.com/azure/aks/private-clusters#configure-private-dns-zone">configure private DNS
zone</a>. Allowed values are &lsquo;system&rsquo; and
&lsquo;none&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>subnetId</code><br/>
<em>
string
</em>
</td>
<td>
<p>SubnetId: It is required when: 1. creating a new cluster with BYO Vnet; 2. updating an existing cluster to enable
apiserver vnet integration.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAPIServerAccessProfile_STATUS_ARM">ManagedClusterAPIServerAccessProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>)
</p>
<div>
<p>Access profile for managed cluster API server.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>authorizedIPRanges</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AuthorizedIPRanges: IP ranges are specified in CIDR format, e.g. 137.117.106.<sup>88</sup>&frasl;<sub>29</sub>. This feature is not compatible with
clusters that use Public IP Per Node, or clusters that are using a Basic Load Balancer. For more information see <a href="https://docs.microsoft.com/azure/aks/api-server-authorized-ip-ranges">API
server authorized IP ranges</a>.</p>
</td>
</tr>
<tr>
<td>
<code>disableRunCommand</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableRunCommand: Whether to disable run command for the cluster or not.</p>
</td>
</tr>
<tr>
<td>
<code>enablePrivateCluster</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePrivateCluster: For more details, see <a href="https://docs.microsoft.com/azure/aks/private-clusters">Creating a private AKS
cluster</a>.</p>
</td>
</tr>
<tr>
<td>
<code>enablePrivateClusterPublicFQDN</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePrivateClusterPublicFQDN: Whether to create additional public FQDN for private cluster or not.</p>
</td>
</tr>
<tr>
<td>
<code>enableVnetIntegration</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableVnetIntegration: Whether to enable apiserver vnet integration for the cluster or not.</p>
</td>
</tr>
<tr>
<td>
<code>privateDNSZone</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateDNSZone: The default is System. For more details see <a href="https://docs.microsoft.com/azure/aks/private-clusters#configure-private-dns-zone">configure private DNS
zone</a>. Allowed values are &lsquo;system&rsquo; and
&lsquo;none&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>subnetId</code><br/>
<em>
string
</em>
</td>
<td>
<p>SubnetId: It is required when: 1. creating a new cluster with BYO Vnet; 2. updating an existing cluster to enable
apiserver vnet integration.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAddonProfile">ManagedClusterAddonProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
</p>
<div>
<p>A Kubernetes add-on profile for a managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>config</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Config: Key-value pairs for configuring an add-on.</p>
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
<p>Enabled: Whether the add-on is enabled or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAddonProfile_ARM">ManagedClusterAddonProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM</a>)
</p>
<div>
<p>A Kubernetes add-on profile for a managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>config</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Config: Key-value pairs for configuring an add-on.</p>
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
<p>Enabled: Whether the add-on is enabled or not.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAddonProfile_STATUS">ManagedClusterAddonProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
</p>
<div>
<p>A Kubernetes add-on profile for a managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>config</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Config: Key-value pairs for configuring an add-on.</p>
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
<p>Enabled: Whether the add-on is enabled or not.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.UserAssignedIdentity_STATUS">
UserAssignedIdentity_STATUS
</a>
</em>
</td>
<td>
<p>Identity: Information of user assigned identity used by this add-on.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAddonProfile_STATUS_ARM">ManagedClusterAddonProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>)
</p>
<div>
<p>A Kubernetes add-on profile for a managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>config</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Config: Key-value pairs for configuring an add-on.</p>
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
<p>Enabled: Whether the add-on is enabled or not.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.UserAssignedIdentity_STATUS_ARM">
UserAssignedIdentity_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Identity: Information of user assigned identity used by this add-on.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
</p>
<div>
<p>Profile for the container service agent pool.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>availabilityZones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AvailabilityZones: The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType
property is &lsquo;VirtualMachineScaleSets&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>capacityReservationGroupID</code><br/>
<em>
string
</em>
</td>
<td>
<p>CapacityReservationGroupID: AKS will associate the specified agent pool with the Capacity Reservation Group.</p>
</td>
</tr>
<tr>
<td>
<code>count</code><br/>
<em>
int
</em>
</td>
<td>
<p>Count: Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive)
for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.</p>
</td>
</tr>
<tr>
<td>
<code>creationData</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.CreationData">
CreationData
</a>
</em>
</td>
<td>
<p>CreationData: CreationData to be used to specify the source Snapshot ID if the node pool will be created/upgraded using
a snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>enableAutoScaling</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAutoScaling: Whether to enable auto-scaler</p>
</td>
</tr>
<tr>
<td>
<code>enableCustomCATrust</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableCustomCATrust: When set to true, AKS adds a label to the node indicating that the feature is enabled and deploys a
daemonset along with host services to sync custom certificate authorities from user-provided list of base64 encoded
certificates into node trust stores. Defaults to false.</p>
</td>
</tr>
<tr>
<td>
<code>enableEncryptionAtHost</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableEncryptionAtHost: This is only supported on certain VM sizes and in certain Azure regions. For more information,
see: <a href="https://docs.microsoft.com/azure/aks/enable-host-encryption">https://docs.microsoft.com/azure/aks/enable-host-encryption</a></p>
</td>
</tr>
<tr>
<td>
<code>enableFIPS</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableFIPS: See <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview">Add a FIPS-enabled node
pool</a> for more
details.</p>
</td>
</tr>
<tr>
<td>
<code>enableNodePublicIP</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableNodePublicIP: Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses.
A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine
to minimize hops. For more information see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools">assigning a public IP per
node</a>. The
default is false.</p>
</td>
</tr>
<tr>
<td>
<code>enableUltraSSD</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableUltraSSD: Whether to enable UltraSSD</p>
</td>
</tr>
<tr>
<td>
<code>gpuInstanceProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.GPUInstanceProfile">
GPUInstanceProfile
</a>
</em>
</td>
<td>
<p>GpuInstanceProfile: GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.</p>
</td>
</tr>
<tr>
<td>
<code>hostGroupReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>HostGroupReference: This is of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Compute/&#x200b;hostGroups/&#x200b;{hostGroupName}.
For more information see <a href="https://docs.microsoft.com/azure/virtual-machines/dedicated-hosts">Azure dedicated hosts</a>.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.KubeletConfig">
KubeletConfig
</a>
</em>
</td>
<td>
<p>KubeletConfig: The Kubelet configuration on the agent pool nodes.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.KubeletDiskType">
KubeletDiskType
</a>
</em>
</td>
<td>
<p>KubeletDiskType: Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral
storage.</p>
</td>
</tr>
<tr>
<td>
<code>linuxOSConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.LinuxOSConfig">
LinuxOSConfig
</a>
</em>
</td>
<td>
<p>LinuxOSConfig: The OS configuration of Linux agent nodes.</p>
</td>
</tr>
<tr>
<td>
<code>maxCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxCount: The maximum number of nodes for auto-scaling</p>
</td>
</tr>
<tr>
<td>
<code>maxPods</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxPods: The maximum number of pods that can run on a node.</p>
</td>
</tr>
<tr>
<td>
<code>messageOfTheDay</code><br/>
<em>
string
</em>
</td>
<td>
<p>MessageOfTheDay: A base64-encoded string which will be written to /etc/motd after decoding. This allows customization of
the message of the day for Linux nodes. It must not be specified for Windows nodes. It must be a static string (i.e.,
will be printed raw and not be executed as a script).</p>
</td>
</tr>
<tr>
<td>
<code>minCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MinCount: The minimum number of nodes for auto-scaling</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolMode">
AgentPoolMode
</a>
</em>
</td>
<td>
<p>Mode: A cluster must have at least one &lsquo;System&rsquo; Agent Pool at all times. For additional information on agent pool
restrictions  and best practices, see: <a href="https://docs.microsoft.com/azure/aks/use-system-pools">https://docs.microsoft.com/azure/aks/use-system-pools</a></p>
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
<p>Name: Windows agent pool names must be 6 characters or less.</p>
</td>
</tr>
<tr>
<td>
<code>networkProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolNetworkProfile">
AgentPoolNetworkProfile
</a>
</em>
</td>
<td>
<p>NetworkProfile: Network-related settings of an agent pool.</p>
</td>
</tr>
<tr>
<td>
<code>nodeLabels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>NodeLabels: The node labels to be persisted across all nodes in agent pool.</p>
</td>
</tr>
<tr>
<td>
<code>nodePublicIPPrefixReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>NodePublicIPPrefixReference: This is of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Network/&#x200b;publicIPPrefixes/&#x200b;{publicIPPrefixName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>nodeTaints</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>NodeTaints: The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.</p>
</td>
</tr>
<tr>
<td>
<code>orchestratorVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>OrchestratorVersion: Both patch version <major.minor.patch> and <major.minor> are supported. When <major.minor> is
specified, the latest supported patch version is chosen automatically. Updating the agent pool with the same
<major.minor> once it has been created will not trigger an upgrade, even if a newer patch version is available. As a
best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version
must have the same major version as the control plane. The node pool minor version must be within two minor versions of
the control plane version. The node pool version cannot be greater than the control plane version. For more information
see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool">upgrading a node pool</a>.</p>
</td>
</tr>
<tr>
<td>
<code>osDiskSizeGB</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceOSDisk">
ContainerServiceOSDisk
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSDiskType">
OSDiskType
</a>
</em>
</td>
<td>
<p>OsDiskType: The default is &lsquo;Ephemeral&rsquo; if the VM supports it and has a cache disk larger than the requested
OSDiskSizeGB. Otherwise,  defaults to &lsquo;Managed&rsquo;. May not be changed after creation. For more information see <a href="https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os">Ephemeral
OS</a>.</p>
</td>
</tr>
<tr>
<td>
<code>osSKU</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSSKU">
OSSKU
</a>
</em>
</td>
<td>
<p>OsSKU: Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or
Windows2019 if  OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is
deprecated.</p>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSType">
OSType
</a>
</em>
</td>
<td>
<p>OsType: The operating system type. The default is Linux.</p>
</td>
</tr>
<tr>
<td>
<code>podSubnetReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>PodSubnetReference: If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details).
This is of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Network/&#x200b;virtualNetworks/&#x200b;{virtualNetworkName}/&#x200b;subnets/&#x200b;{subnetName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>powerState</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PowerState">
PowerState
</a>
</em>
</td>
<td>
<p>PowerState: When an Agent Pool is first created it is initially Running. The Agent Pool can be stopped by setting this
field to Stopped. A stopped Agent Pool stops all of its VMs and does not accrue billing charges. An Agent Pool can only
be stopped if it is Running and provisioning state is Succeeded</p>
</td>
</tr>
<tr>
<td>
<code>proximityPlacementGroupReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>ProximityPlacementGroupReference: The ID for Proximity Placement Group.</p>
</td>
</tr>
<tr>
<td>
<code>scaleDownMode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleDownMode">
ScaleDownMode
</a>
</em>
</td>
<td>
<p>ScaleDownMode: This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.</p>
</td>
</tr>
<tr>
<td>
<code>scaleSetEvictionPolicy</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleSetEvictionPolicy">
ScaleSetEvictionPolicy
</a>
</em>
</td>
<td>
<p>ScaleSetEvictionPolicy: This cannot be specified unless the scaleSetPriority is &lsquo;Spot&rsquo;. If not specified, the default is
&lsquo;Delete&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>scaleSetPriority</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleSetPriority">
ScaleSetPriority
</a>
</em>
</td>
<td>
<p>ScaleSetPriority: The Virtual Machine Scale Set priority. If not specified, the default is &lsquo;Regular&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>spotMaxPrice</code><br/>
<em>
float64
</em>
</td>
<td>
<p>SpotMaxPrice: Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any
on-demand price. For more details on spot pricing, see <a href="https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing">spot VMs
pricing</a></p>
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
<p>Tags: The tags to be persisted on the agent pool virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolType">
AgentPoolType
</a>
</em>
</td>
<td>
<p>Type: The type of Agent Pool.</p>
</td>
</tr>
<tr>
<td>
<code>upgradeSettings</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolUpgradeSettings">
AgentPoolUpgradeSettings
</a>
</em>
</td>
<td>
<p>UpgradeSettings: Settings for upgrading the agentpool</p>
</td>
</tr>
<tr>
<td>
<code>vmSize</code><br/>
<em>
string
</em>
</td>
<td>
<p>VmSize: VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods
might fail to run correctly. For more details on restricted VM sizes, see:
<a href="https://docs.microsoft.com/azure/aks/quotas-skus-regions">https://docs.microsoft.com/azure/aks/quotas-skus-regions</a></p>
</td>
</tr>
<tr>
<td>
<code>vnetSubnetReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>VnetSubnetReference: If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is
specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Network/&#x200b;virtualNetworks/&#x200b;{virtualNetworkName}/&#x200b;subnets/&#x200b;{subnetName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>windowsProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolWindowsProfile">
AgentPoolWindowsProfile
</a>
</em>
</td>
<td>
<p>WindowsProfile: The Windows agent pool&rsquo;s specific profile.</p>
</td>
</tr>
<tr>
<td>
<code>workloadRuntime</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.WorkloadRuntime">
WorkloadRuntime
</a>
</em>
</td>
<td>
<p>WorkloadRuntime: Determines the type of workload a node can run.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_ARM">ManagedClusterAgentPoolProfileProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_Spec_ARM">ManagedClusters_AgentPool_Spec_ARM</a>)
</p>
<div>
<p>Properties for the container service agent pool profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>availabilityZones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AvailabilityZones: The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType
property is &lsquo;VirtualMachineScaleSets&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>capacityReservationGroupID</code><br/>
<em>
string
</em>
</td>
<td>
<p>CapacityReservationGroupID: AKS will associate the specified agent pool with the Capacity Reservation Group.</p>
</td>
</tr>
<tr>
<td>
<code>count</code><br/>
<em>
int
</em>
</td>
<td>
<p>Count: Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive)
for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.</p>
</td>
</tr>
<tr>
<td>
<code>creationData</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.CreationData_ARM">
CreationData_ARM
</a>
</em>
</td>
<td>
<p>CreationData: CreationData to be used to specify the source Snapshot ID if the node pool will be created/upgraded using
a snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>enableAutoScaling</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAutoScaling: Whether to enable auto-scaler</p>
</td>
</tr>
<tr>
<td>
<code>enableCustomCATrust</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableCustomCATrust: When set to true, AKS adds a label to the node indicating that the feature is enabled and deploys a
daemonset along with host services to sync custom certificate authorities from user-provided list of base64 encoded
certificates into node trust stores. Defaults to false.</p>
</td>
</tr>
<tr>
<td>
<code>enableEncryptionAtHost</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableEncryptionAtHost: This is only supported on certain VM sizes and in certain Azure regions. For more information,
see: <a href="https://docs.microsoft.com/azure/aks/enable-host-encryption">https://docs.microsoft.com/azure/aks/enable-host-encryption</a></p>
</td>
</tr>
<tr>
<td>
<code>enableFIPS</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableFIPS: See <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview">Add a FIPS-enabled node
pool</a> for more
details.</p>
</td>
</tr>
<tr>
<td>
<code>enableNodePublicIP</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableNodePublicIP: Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses.
A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine
to minimize hops. For more information see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools">assigning a public IP per
node</a>. The
default is false.</p>
</td>
</tr>
<tr>
<td>
<code>enableUltraSSD</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableUltraSSD: Whether to enable UltraSSD</p>
</td>
</tr>
<tr>
<td>
<code>gpuInstanceProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.GPUInstanceProfile">
GPUInstanceProfile
</a>
</em>
</td>
<td>
<p>GpuInstanceProfile: GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.</p>
</td>
</tr>
<tr>
<td>
<code>hostGroupID</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>kubeletConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.KubeletConfig_ARM">
KubeletConfig_ARM
</a>
</em>
</td>
<td>
<p>KubeletConfig: The Kubelet configuration on the agent pool nodes.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.KubeletDiskType">
KubeletDiskType
</a>
</em>
</td>
<td>
<p>KubeletDiskType: Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral
storage.</p>
</td>
</tr>
<tr>
<td>
<code>linuxOSConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.LinuxOSConfig_ARM">
LinuxOSConfig_ARM
</a>
</em>
</td>
<td>
<p>LinuxOSConfig: The OS configuration of Linux agent nodes.</p>
</td>
</tr>
<tr>
<td>
<code>maxCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxCount: The maximum number of nodes for auto-scaling</p>
</td>
</tr>
<tr>
<td>
<code>maxPods</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxPods: The maximum number of pods that can run on a node.</p>
</td>
</tr>
<tr>
<td>
<code>messageOfTheDay</code><br/>
<em>
string
</em>
</td>
<td>
<p>MessageOfTheDay: A base64-encoded string which will be written to /etc/motd after decoding. This allows customization of
the message of the day for Linux nodes. It must not be specified for Windows nodes. It must be a static string (i.e.,
will be printed raw and not be executed as a script).</p>
</td>
</tr>
<tr>
<td>
<code>minCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MinCount: The minimum number of nodes for auto-scaling</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolMode">
AgentPoolMode
</a>
</em>
</td>
<td>
<p>Mode: A cluster must have at least one &lsquo;System&rsquo; Agent Pool at all times. For additional information on agent pool
restrictions  and best practices, see: <a href="https://docs.microsoft.com/azure/aks/use-system-pools">https://docs.microsoft.com/azure/aks/use-system-pools</a></p>
</td>
</tr>
<tr>
<td>
<code>networkProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolNetworkProfile_ARM">
AgentPoolNetworkProfile_ARM
</a>
</em>
</td>
<td>
<p>NetworkProfile: Network-related settings of an agent pool.</p>
</td>
</tr>
<tr>
<td>
<code>nodeLabels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>NodeLabels: The node labels to be persisted across all nodes in agent pool.</p>
</td>
</tr>
<tr>
<td>
<code>nodePublicIPPrefixID</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>nodeTaints</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>NodeTaints: The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.</p>
</td>
</tr>
<tr>
<td>
<code>orchestratorVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>OrchestratorVersion: Both patch version <major.minor.patch> and <major.minor> are supported. When <major.minor> is
specified, the latest supported patch version is chosen automatically. Updating the agent pool with the same
<major.minor> once it has been created will not trigger an upgrade, even if a newer patch version is available. As a
best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version
must have the same major version as the control plane. The node pool minor version must be within two minor versions of
the control plane version. The node pool version cannot be greater than the control plane version. For more information
see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool">upgrading a node pool</a>.</p>
</td>
</tr>
<tr>
<td>
<code>osDiskSizeGB</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceOSDisk">
ContainerServiceOSDisk
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSDiskType">
OSDiskType
</a>
</em>
</td>
<td>
<p>OsDiskType: The default is &lsquo;Ephemeral&rsquo; if the VM supports it and has a cache disk larger than the requested
OSDiskSizeGB. Otherwise,  defaults to &lsquo;Managed&rsquo;. May not be changed after creation. For more information see <a href="https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os">Ephemeral
OS</a>.</p>
</td>
</tr>
<tr>
<td>
<code>osSKU</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSSKU">
OSSKU
</a>
</em>
</td>
<td>
<p>OsSKU: Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or
Windows2019 if  OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is
deprecated.</p>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSType">
OSType
</a>
</em>
</td>
<td>
<p>OsType: The operating system type. The default is Linux.</p>
</td>
</tr>
<tr>
<td>
<code>podSubnetID</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>powerState</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PowerState_ARM">
PowerState_ARM
</a>
</em>
</td>
<td>
<p>PowerState: When an Agent Pool is first created it is initially Running. The Agent Pool can be stopped by setting this
field to Stopped. A stopped Agent Pool stops all of its VMs and does not accrue billing charges. An Agent Pool can only
be stopped if it is Running and provisioning state is Succeeded</p>
</td>
</tr>
<tr>
<td>
<code>proximityPlacementGroupID</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>scaleDownMode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleDownMode">
ScaleDownMode
</a>
</em>
</td>
<td>
<p>ScaleDownMode: This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.</p>
</td>
</tr>
<tr>
<td>
<code>scaleSetEvictionPolicy</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleSetEvictionPolicy">
ScaleSetEvictionPolicy
</a>
</em>
</td>
<td>
<p>ScaleSetEvictionPolicy: This cannot be specified unless the scaleSetPriority is &lsquo;Spot&rsquo;. If not specified, the default is
&lsquo;Delete&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>scaleSetPriority</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleSetPriority">
ScaleSetPriority
</a>
</em>
</td>
<td>
<p>ScaleSetPriority: The Virtual Machine Scale Set priority. If not specified, the default is &lsquo;Regular&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>spotMaxPrice</code><br/>
<em>
float64
</em>
</td>
<td>
<p>SpotMaxPrice: Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any
on-demand price. For more details on spot pricing, see <a href="https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing">spot VMs
pricing</a></p>
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
<p>Tags: The tags to be persisted on the agent pool virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolType">
AgentPoolType
</a>
</em>
</td>
<td>
<p>Type: The type of Agent Pool.</p>
</td>
</tr>
<tr>
<td>
<code>upgradeSettings</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolUpgradeSettings_ARM">
AgentPoolUpgradeSettings_ARM
</a>
</em>
</td>
<td>
<p>UpgradeSettings: Settings for upgrading the agentpool</p>
</td>
</tr>
<tr>
<td>
<code>vmSize</code><br/>
<em>
string
</em>
</td>
<td>
<p>VmSize: VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods
might fail to run correctly. For more details on restricted VM sizes, see:
<a href="https://docs.microsoft.com/azure/aks/quotas-skus-regions">https://docs.microsoft.com/azure/aks/quotas-skus-regions</a></p>
</td>
</tr>
<tr>
<td>
<code>vnetSubnetID</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>windowsProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolWindowsProfile_ARM">
AgentPoolWindowsProfile_ARM
</a>
</em>
</td>
<td>
<p>WindowsProfile: The Windows agent pool&rsquo;s specific profile.</p>
</td>
</tr>
<tr>
<td>
<code>workloadRuntime</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.WorkloadRuntime">
WorkloadRuntime
</a>
</em>
</td>
<td>
<p>WorkloadRuntime: Determines the type of workload a node can run.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_STATUS_ARM">ManagedClusterAgentPoolProfileProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_STATUS_ARM">ManagedClusters_AgentPool_STATUS_ARM</a>)
</p>
<div>
<p>Properties for the container service agent pool profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>availabilityZones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AvailabilityZones: The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType
property is &lsquo;VirtualMachineScaleSets&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>capacityReservationGroupID</code><br/>
<em>
string
</em>
</td>
<td>
<p>CapacityReservationGroupID: AKS will associate the specified agent pool with the Capacity Reservation Group.</p>
</td>
</tr>
<tr>
<td>
<code>count</code><br/>
<em>
int
</em>
</td>
<td>
<p>Count: Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive)
for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.</p>
</td>
</tr>
<tr>
<td>
<code>creationData</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.CreationData_STATUS_ARM">
CreationData_STATUS_ARM
</a>
</em>
</td>
<td>
<p>CreationData: CreationData to be used to specify the source Snapshot ID if the node pool will be created/upgraded using
a snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>currentOrchestratorVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>CurrentOrchestratorVersion: If orchestratorVersion was a fully specified version <major.minor.patch>, this field will be
exactly equal to it. If orchestratorVersion was <major.minor>, this field will contain the full <major.minor.patch>
version being used.</p>
</td>
</tr>
<tr>
<td>
<code>enableAutoScaling</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAutoScaling: Whether to enable auto-scaler</p>
</td>
</tr>
<tr>
<td>
<code>enableCustomCATrust</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableCustomCATrust: When set to true, AKS adds a label to the node indicating that the feature is enabled and deploys a
daemonset along with host services to sync custom certificate authorities from user-provided list of base64 encoded
certificates into node trust stores. Defaults to false.</p>
</td>
</tr>
<tr>
<td>
<code>enableEncryptionAtHost</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableEncryptionAtHost: This is only supported on certain VM sizes and in certain Azure regions. For more information,
see: <a href="https://docs.microsoft.com/azure/aks/enable-host-encryption">https://docs.microsoft.com/azure/aks/enable-host-encryption</a></p>
</td>
</tr>
<tr>
<td>
<code>enableFIPS</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableFIPS: See <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview">Add a FIPS-enabled node
pool</a> for more
details.</p>
</td>
</tr>
<tr>
<td>
<code>enableNodePublicIP</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableNodePublicIP: Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses.
A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine
to minimize hops. For more information see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools">assigning a public IP per
node</a>. The
default is false.</p>
</td>
</tr>
<tr>
<td>
<code>enableUltraSSD</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableUltraSSD: Whether to enable UltraSSD</p>
</td>
</tr>
<tr>
<td>
<code>gpuInstanceProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.GPUInstanceProfile_STATUS">
GPUInstanceProfile_STATUS
</a>
</em>
</td>
<td>
<p>GpuInstanceProfile: GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.</p>
</td>
</tr>
<tr>
<td>
<code>hostGroupID</code><br/>
<em>
string
</em>
</td>
<td>
<p>HostGroupID: This is of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Compute/&#x200b;hostGroups/&#x200b;{hostGroupName}.
For more information see <a href="https://docs.microsoft.com/azure/virtual-machines/dedicated-hosts">Azure dedicated hosts</a>.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.KubeletConfig_STATUS_ARM">
KubeletConfig_STATUS_ARM
</a>
</em>
</td>
<td>
<p>KubeletConfig: The Kubelet configuration on the agent pool nodes.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.KubeletDiskType_STATUS">
KubeletDiskType_STATUS
</a>
</em>
</td>
<td>
<p>KubeletDiskType: Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral
storage.</p>
</td>
</tr>
<tr>
<td>
<code>linuxOSConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.LinuxOSConfig_STATUS_ARM">
LinuxOSConfig_STATUS_ARM
</a>
</em>
</td>
<td>
<p>LinuxOSConfig: The OS configuration of Linux agent nodes.</p>
</td>
</tr>
<tr>
<td>
<code>maxCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxCount: The maximum number of nodes for auto-scaling</p>
</td>
</tr>
<tr>
<td>
<code>maxPods</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxPods: The maximum number of pods that can run on a node.</p>
</td>
</tr>
<tr>
<td>
<code>messageOfTheDay</code><br/>
<em>
string
</em>
</td>
<td>
<p>MessageOfTheDay: A base64-encoded string which will be written to /etc/motd after decoding. This allows customization of
the message of the day for Linux nodes. It must not be specified for Windows nodes. It must be a static string (i.e.,
will be printed raw and not be executed as a script).</p>
</td>
</tr>
<tr>
<td>
<code>minCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MinCount: The minimum number of nodes for auto-scaling</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolMode_STATUS">
AgentPoolMode_STATUS
</a>
</em>
</td>
<td>
<p>Mode: A cluster must have at least one &lsquo;System&rsquo; Agent Pool at all times. For additional information on agent pool
restrictions  and best practices, see: <a href="https://docs.microsoft.com/azure/aks/use-system-pools">https://docs.microsoft.com/azure/aks/use-system-pools</a></p>
</td>
</tr>
<tr>
<td>
<code>networkProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolNetworkProfile_STATUS_ARM">
AgentPoolNetworkProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>NetworkProfile: Network-related settings of an agent pool.</p>
</td>
</tr>
<tr>
<td>
<code>nodeImageVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>NodeImageVersion: The version of node image</p>
</td>
</tr>
<tr>
<td>
<code>nodeLabels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>NodeLabels: The node labels to be persisted across all nodes in agent pool.</p>
</td>
</tr>
<tr>
<td>
<code>nodePublicIPPrefixID</code><br/>
<em>
string
</em>
</td>
<td>
<p>NodePublicIPPrefixID: This is of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Network/&#x200b;publicIPPrefixes/&#x200b;{publicIPPrefixName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>nodeTaints</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>NodeTaints: The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.</p>
</td>
</tr>
<tr>
<td>
<code>orchestratorVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>OrchestratorVersion: Both patch version <major.minor.patch> and <major.minor> are supported. When <major.minor> is
specified, the latest supported patch version is chosen automatically. Updating the agent pool with the same
<major.minor> once it has been created will not trigger an upgrade, even if a newer patch version is available. As a
best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version
must have the same major version as the control plane. The node pool minor version must be within two minor versions of
the control plane version. The node pool version cannot be greater than the control plane version. For more information
see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool">upgrading a node pool</a>.</p>
</td>
</tr>
<tr>
<td>
<code>osDiskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSDiskType_STATUS">
OSDiskType_STATUS
</a>
</em>
</td>
<td>
<p>OsDiskType: The default is &lsquo;Ephemeral&rsquo; if the VM supports it and has a cache disk larger than the requested
OSDiskSizeGB. Otherwise,  defaults to &lsquo;Managed&rsquo;. May not be changed after creation. For more information see <a href="https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os">Ephemeral
OS</a>.</p>
</td>
</tr>
<tr>
<td>
<code>osSKU</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSSKU_STATUS">
OSSKU_STATUS
</a>
</em>
</td>
<td>
<p>OsSKU: Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or
Windows2019 if  OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is
deprecated.</p>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSType_STATUS">
OSType_STATUS
</a>
</em>
</td>
<td>
<p>OsType: The operating system type. The default is Linux.</p>
</td>
</tr>
<tr>
<td>
<code>podSubnetID</code><br/>
<em>
string
</em>
</td>
<td>
<p>PodSubnetID: If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is
of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Network/&#x200b;virtualNetworks/&#x200b;{virtualNetworkName}/&#x200b;subnets/&#x200b;{subnetName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>powerState</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PowerState_STATUS_ARM">
PowerState_STATUS_ARM
</a>
</em>
</td>
<td>
<p>PowerState: When an Agent Pool is first created it is initially Running. The Agent Pool can be stopped by setting this
field to Stopped. A stopped Agent Pool stops all of its VMs and does not accrue billing charges. An Agent Pool can only
be stopped if it is Running and provisioning state is Succeeded</p>
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
<p>ProvisioningState: The current deployment or provisioning state.</p>
</td>
</tr>
<tr>
<td>
<code>proximityPlacementGroupID</code><br/>
<em>
string
</em>
</td>
<td>
<p>ProximityPlacementGroupID: The ID for Proximity Placement Group.</p>
</td>
</tr>
<tr>
<td>
<code>scaleDownMode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleDownMode_STATUS">
ScaleDownMode_STATUS
</a>
</em>
</td>
<td>
<p>ScaleDownMode: This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.</p>
</td>
</tr>
<tr>
<td>
<code>scaleSetEvictionPolicy</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleSetEvictionPolicy_STATUS">
ScaleSetEvictionPolicy_STATUS
</a>
</em>
</td>
<td>
<p>ScaleSetEvictionPolicy: This cannot be specified unless the scaleSetPriority is &lsquo;Spot&rsquo;. If not specified, the default is
&lsquo;Delete&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>scaleSetPriority</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleSetPriority_STATUS">
ScaleSetPriority_STATUS
</a>
</em>
</td>
<td>
<p>ScaleSetPriority: The Virtual Machine Scale Set priority. If not specified, the default is &lsquo;Regular&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>spotMaxPrice</code><br/>
<em>
float64
</em>
</td>
<td>
<p>SpotMaxPrice: Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any
on-demand price. For more details on spot pricing, see <a href="https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing">spot VMs
pricing</a></p>
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
<p>Tags: The tags to be persisted on the agent pool virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolType_STATUS">
AgentPoolType_STATUS
</a>
</em>
</td>
<td>
<p>Type: The type of Agent Pool.</p>
</td>
</tr>
<tr>
<td>
<code>upgradeSettings</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolUpgradeSettings_STATUS_ARM">
AgentPoolUpgradeSettings_STATUS_ARM
</a>
</em>
</td>
<td>
<p>UpgradeSettings: Settings for upgrading the agentpool</p>
</td>
</tr>
<tr>
<td>
<code>vmSize</code><br/>
<em>
string
</em>
</td>
<td>
<p>VmSize: VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods
might fail to run correctly. For more details on restricted VM sizes, see:
<a href="https://docs.microsoft.com/azure/aks/quotas-skus-regions">https://docs.microsoft.com/azure/aks/quotas-skus-regions</a></p>
</td>
</tr>
<tr>
<td>
<code>vnetSubnetID</code><br/>
<em>
string
</em>
</td>
<td>
<p>VnetSubnetID: If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified,
this applies to nodes and pods, otherwise it applies to just nodes. This is of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Network/&#x200b;virtualNetworks/&#x200b;{virtualNetworkName}/&#x200b;subnets/&#x200b;{subnetName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>windowsProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolWindowsProfile_STATUS_ARM">
AgentPoolWindowsProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>WindowsProfile: The Windows agent pool&rsquo;s specific profile.</p>
</td>
</tr>
<tr>
<td>
<code>workloadRuntime</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.WorkloadRuntime_STATUS">
WorkloadRuntime_STATUS
</a>
</em>
</td>
<td>
<p>WorkloadRuntime: Determines the type of workload a node can run.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_ARM">ManagedClusterAgentPoolProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM</a>)
</p>
<div>
<p>Profile for the container service agent pool.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>availabilityZones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AvailabilityZones: The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType
property is &lsquo;VirtualMachineScaleSets&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>capacityReservationGroupID</code><br/>
<em>
string
</em>
</td>
<td>
<p>CapacityReservationGroupID: AKS will associate the specified agent pool with the Capacity Reservation Group.</p>
</td>
</tr>
<tr>
<td>
<code>count</code><br/>
<em>
int
</em>
</td>
<td>
<p>Count: Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive)
for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.</p>
</td>
</tr>
<tr>
<td>
<code>creationData</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.CreationData_ARM">
CreationData_ARM
</a>
</em>
</td>
<td>
<p>CreationData: CreationData to be used to specify the source Snapshot ID if the node pool will be created/upgraded using
a snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>enableAutoScaling</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAutoScaling: Whether to enable auto-scaler</p>
</td>
</tr>
<tr>
<td>
<code>enableCustomCATrust</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableCustomCATrust: When set to true, AKS adds a label to the node indicating that the feature is enabled and deploys a
daemonset along with host services to sync custom certificate authorities from user-provided list of base64 encoded
certificates into node trust stores. Defaults to false.</p>
</td>
</tr>
<tr>
<td>
<code>enableEncryptionAtHost</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableEncryptionAtHost: This is only supported on certain VM sizes and in certain Azure regions. For more information,
see: <a href="https://docs.microsoft.com/azure/aks/enable-host-encryption">https://docs.microsoft.com/azure/aks/enable-host-encryption</a></p>
</td>
</tr>
<tr>
<td>
<code>enableFIPS</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableFIPS: See <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview">Add a FIPS-enabled node
pool</a> for more
details.</p>
</td>
</tr>
<tr>
<td>
<code>enableNodePublicIP</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableNodePublicIP: Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses.
A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine
to minimize hops. For more information see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools">assigning a public IP per
node</a>. The
default is false.</p>
</td>
</tr>
<tr>
<td>
<code>enableUltraSSD</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableUltraSSD: Whether to enable UltraSSD</p>
</td>
</tr>
<tr>
<td>
<code>gpuInstanceProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.GPUInstanceProfile">
GPUInstanceProfile
</a>
</em>
</td>
<td>
<p>GpuInstanceProfile: GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.</p>
</td>
</tr>
<tr>
<td>
<code>hostGroupID</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>kubeletConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.KubeletConfig_ARM">
KubeletConfig_ARM
</a>
</em>
</td>
<td>
<p>KubeletConfig: The Kubelet configuration on the agent pool nodes.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.KubeletDiskType">
KubeletDiskType
</a>
</em>
</td>
<td>
<p>KubeletDiskType: Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral
storage.</p>
</td>
</tr>
<tr>
<td>
<code>linuxOSConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.LinuxOSConfig_ARM">
LinuxOSConfig_ARM
</a>
</em>
</td>
<td>
<p>LinuxOSConfig: The OS configuration of Linux agent nodes.</p>
</td>
</tr>
<tr>
<td>
<code>maxCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxCount: The maximum number of nodes for auto-scaling</p>
</td>
</tr>
<tr>
<td>
<code>maxPods</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxPods: The maximum number of pods that can run on a node.</p>
</td>
</tr>
<tr>
<td>
<code>messageOfTheDay</code><br/>
<em>
string
</em>
</td>
<td>
<p>MessageOfTheDay: A base64-encoded string which will be written to /etc/motd after decoding. This allows customization of
the message of the day for Linux nodes. It must not be specified for Windows nodes. It must be a static string (i.e.,
will be printed raw and not be executed as a script).</p>
</td>
</tr>
<tr>
<td>
<code>minCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MinCount: The minimum number of nodes for auto-scaling</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolMode">
AgentPoolMode
</a>
</em>
</td>
<td>
<p>Mode: A cluster must have at least one &lsquo;System&rsquo; Agent Pool at all times. For additional information on agent pool
restrictions  and best practices, see: <a href="https://docs.microsoft.com/azure/aks/use-system-pools">https://docs.microsoft.com/azure/aks/use-system-pools</a></p>
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
<p>Name: Windows agent pool names must be 6 characters or less.</p>
</td>
</tr>
<tr>
<td>
<code>networkProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolNetworkProfile_ARM">
AgentPoolNetworkProfile_ARM
</a>
</em>
</td>
<td>
<p>NetworkProfile: Network-related settings of an agent pool.</p>
</td>
</tr>
<tr>
<td>
<code>nodeLabels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>NodeLabels: The node labels to be persisted across all nodes in agent pool.</p>
</td>
</tr>
<tr>
<td>
<code>nodePublicIPPrefixID</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>nodeTaints</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>NodeTaints: The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.</p>
</td>
</tr>
<tr>
<td>
<code>orchestratorVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>OrchestratorVersion: Both patch version <major.minor.patch> and <major.minor> are supported. When <major.minor> is
specified, the latest supported patch version is chosen automatically. Updating the agent pool with the same
<major.minor> once it has been created will not trigger an upgrade, even if a newer patch version is available. As a
best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version
must have the same major version as the control plane. The node pool minor version must be within two minor versions of
the control plane version. The node pool version cannot be greater than the control plane version. For more information
see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool">upgrading a node pool</a>.</p>
</td>
</tr>
<tr>
<td>
<code>osDiskSizeGB</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceOSDisk">
ContainerServiceOSDisk
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSDiskType">
OSDiskType
</a>
</em>
</td>
<td>
<p>OsDiskType: The default is &lsquo;Ephemeral&rsquo; if the VM supports it and has a cache disk larger than the requested
OSDiskSizeGB. Otherwise,  defaults to &lsquo;Managed&rsquo;. May not be changed after creation. For more information see <a href="https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os">Ephemeral
OS</a>.</p>
</td>
</tr>
<tr>
<td>
<code>osSKU</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSSKU">
OSSKU
</a>
</em>
</td>
<td>
<p>OsSKU: Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or
Windows2019 if  OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is
deprecated.</p>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSType">
OSType
</a>
</em>
</td>
<td>
<p>OsType: The operating system type. The default is Linux.</p>
</td>
</tr>
<tr>
<td>
<code>podSubnetID</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>powerState</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PowerState_ARM">
PowerState_ARM
</a>
</em>
</td>
<td>
<p>PowerState: When an Agent Pool is first created it is initially Running. The Agent Pool can be stopped by setting this
field to Stopped. A stopped Agent Pool stops all of its VMs and does not accrue billing charges. An Agent Pool can only
be stopped if it is Running and provisioning state is Succeeded</p>
</td>
</tr>
<tr>
<td>
<code>proximityPlacementGroupID</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>scaleDownMode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleDownMode">
ScaleDownMode
</a>
</em>
</td>
<td>
<p>ScaleDownMode: This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.</p>
</td>
</tr>
<tr>
<td>
<code>scaleSetEvictionPolicy</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleSetEvictionPolicy">
ScaleSetEvictionPolicy
</a>
</em>
</td>
<td>
<p>ScaleSetEvictionPolicy: This cannot be specified unless the scaleSetPriority is &lsquo;Spot&rsquo;. If not specified, the default is
&lsquo;Delete&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>scaleSetPriority</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleSetPriority">
ScaleSetPriority
</a>
</em>
</td>
<td>
<p>ScaleSetPriority: The Virtual Machine Scale Set priority. If not specified, the default is &lsquo;Regular&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>spotMaxPrice</code><br/>
<em>
float64
</em>
</td>
<td>
<p>SpotMaxPrice: Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any
on-demand price. For more details on spot pricing, see <a href="https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing">spot VMs
pricing</a></p>
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
<p>Tags: The tags to be persisted on the agent pool virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolType">
AgentPoolType
</a>
</em>
</td>
<td>
<p>Type: The type of Agent Pool.</p>
</td>
</tr>
<tr>
<td>
<code>upgradeSettings</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolUpgradeSettings_ARM">
AgentPoolUpgradeSettings_ARM
</a>
</em>
</td>
<td>
<p>UpgradeSettings: Settings for upgrading the agentpool</p>
</td>
</tr>
<tr>
<td>
<code>vmSize</code><br/>
<em>
string
</em>
</td>
<td>
<p>VmSize: VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods
might fail to run correctly. For more details on restricted VM sizes, see:
<a href="https://docs.microsoft.com/azure/aks/quotas-skus-regions">https://docs.microsoft.com/azure/aks/quotas-skus-regions</a></p>
</td>
</tr>
<tr>
<td>
<code>vnetSubnetID</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>windowsProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolWindowsProfile_ARM">
AgentPoolWindowsProfile_ARM
</a>
</em>
</td>
<td>
<p>WindowsProfile: The Windows agent pool&rsquo;s specific profile.</p>
</td>
</tr>
<tr>
<td>
<code>workloadRuntime</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.WorkloadRuntime">
WorkloadRuntime
</a>
</em>
</td>
<td>
<p>WorkloadRuntime: Determines the type of workload a node can run.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS">ManagedClusterAgentPoolProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
</p>
<div>
<p>Profile for the container service agent pool.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>availabilityZones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AvailabilityZones: The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType
property is &lsquo;VirtualMachineScaleSets&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>capacityReservationGroupID</code><br/>
<em>
string
</em>
</td>
<td>
<p>CapacityReservationGroupID: AKS will associate the specified agent pool with the Capacity Reservation Group.</p>
</td>
</tr>
<tr>
<td>
<code>count</code><br/>
<em>
int
</em>
</td>
<td>
<p>Count: Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive)
for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.</p>
</td>
</tr>
<tr>
<td>
<code>creationData</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.CreationData_STATUS">
CreationData_STATUS
</a>
</em>
</td>
<td>
<p>CreationData: CreationData to be used to specify the source Snapshot ID if the node pool will be created/upgraded using
a snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>currentOrchestratorVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>CurrentOrchestratorVersion: If orchestratorVersion was a fully specified version <major.minor.patch>, this field will be
exactly equal to it. If orchestratorVersion was <major.minor>, this field will contain the full <major.minor.patch>
version being used.</p>
</td>
</tr>
<tr>
<td>
<code>enableAutoScaling</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAutoScaling: Whether to enable auto-scaler</p>
</td>
</tr>
<tr>
<td>
<code>enableCustomCATrust</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableCustomCATrust: When set to true, AKS adds a label to the node indicating that the feature is enabled and deploys a
daemonset along with host services to sync custom certificate authorities from user-provided list of base64 encoded
certificates into node trust stores. Defaults to false.</p>
</td>
</tr>
<tr>
<td>
<code>enableEncryptionAtHost</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableEncryptionAtHost: This is only supported on certain VM sizes and in certain Azure regions. For more information,
see: <a href="https://docs.microsoft.com/azure/aks/enable-host-encryption">https://docs.microsoft.com/azure/aks/enable-host-encryption</a></p>
</td>
</tr>
<tr>
<td>
<code>enableFIPS</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableFIPS: See <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview">Add a FIPS-enabled node
pool</a> for more
details.</p>
</td>
</tr>
<tr>
<td>
<code>enableNodePublicIP</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableNodePublicIP: Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses.
A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine
to minimize hops. For more information see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools">assigning a public IP per
node</a>. The
default is false.</p>
</td>
</tr>
<tr>
<td>
<code>enableUltraSSD</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableUltraSSD: Whether to enable UltraSSD</p>
</td>
</tr>
<tr>
<td>
<code>gpuInstanceProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.GPUInstanceProfile_STATUS">
GPUInstanceProfile_STATUS
</a>
</em>
</td>
<td>
<p>GpuInstanceProfile: GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.</p>
</td>
</tr>
<tr>
<td>
<code>hostGroupID</code><br/>
<em>
string
</em>
</td>
<td>
<p>HostGroupID: This is of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Compute/&#x200b;hostGroups/&#x200b;{hostGroupName}.
For more information see <a href="https://docs.microsoft.com/azure/virtual-machines/dedicated-hosts">Azure dedicated hosts</a>.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.KubeletConfig_STATUS">
KubeletConfig_STATUS
</a>
</em>
</td>
<td>
<p>KubeletConfig: The Kubelet configuration on the agent pool nodes.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.KubeletDiskType_STATUS">
KubeletDiskType_STATUS
</a>
</em>
</td>
<td>
<p>KubeletDiskType: Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral
storage.</p>
</td>
</tr>
<tr>
<td>
<code>linuxOSConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.LinuxOSConfig_STATUS">
LinuxOSConfig_STATUS
</a>
</em>
</td>
<td>
<p>LinuxOSConfig: The OS configuration of Linux agent nodes.</p>
</td>
</tr>
<tr>
<td>
<code>maxCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxCount: The maximum number of nodes for auto-scaling</p>
</td>
</tr>
<tr>
<td>
<code>maxPods</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxPods: The maximum number of pods that can run on a node.</p>
</td>
</tr>
<tr>
<td>
<code>messageOfTheDay</code><br/>
<em>
string
</em>
</td>
<td>
<p>MessageOfTheDay: A base64-encoded string which will be written to /etc/motd after decoding. This allows customization of
the message of the day for Linux nodes. It must not be specified for Windows nodes. It must be a static string (i.e.,
will be printed raw and not be executed as a script).</p>
</td>
</tr>
<tr>
<td>
<code>minCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MinCount: The minimum number of nodes for auto-scaling</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolMode_STATUS">
AgentPoolMode_STATUS
</a>
</em>
</td>
<td>
<p>Mode: A cluster must have at least one &lsquo;System&rsquo; Agent Pool at all times. For additional information on agent pool
restrictions  and best practices, see: <a href="https://docs.microsoft.com/azure/aks/use-system-pools">https://docs.microsoft.com/azure/aks/use-system-pools</a></p>
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
<p>Name: Windows agent pool names must be 6 characters or less.</p>
</td>
</tr>
<tr>
<td>
<code>networkProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolNetworkProfile_STATUS">
AgentPoolNetworkProfile_STATUS
</a>
</em>
</td>
<td>
<p>NetworkProfile: Network-related settings of an agent pool.</p>
</td>
</tr>
<tr>
<td>
<code>nodeImageVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>NodeImageVersion: The version of node image</p>
</td>
</tr>
<tr>
<td>
<code>nodeLabels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>NodeLabels: The node labels to be persisted across all nodes in agent pool.</p>
</td>
</tr>
<tr>
<td>
<code>nodePublicIPPrefixID</code><br/>
<em>
string
</em>
</td>
<td>
<p>NodePublicIPPrefixID: This is of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Network/&#x200b;publicIPPrefixes/&#x200b;{publicIPPrefixName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>nodeTaints</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>NodeTaints: The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.</p>
</td>
</tr>
<tr>
<td>
<code>orchestratorVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>OrchestratorVersion: Both patch version <major.minor.patch> and <major.minor> are supported. When <major.minor> is
specified, the latest supported patch version is chosen automatically. Updating the agent pool with the same
<major.minor> once it has been created will not trigger an upgrade, even if a newer patch version is available. As a
best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version
must have the same major version as the control plane. The node pool minor version must be within two minor versions of
the control plane version. The node pool version cannot be greater than the control plane version. For more information
see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool">upgrading a node pool</a>.</p>
</td>
</tr>
<tr>
<td>
<code>osDiskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSDiskType_STATUS">
OSDiskType_STATUS
</a>
</em>
</td>
<td>
<p>OsDiskType: The default is &lsquo;Ephemeral&rsquo; if the VM supports it and has a cache disk larger than the requested
OSDiskSizeGB. Otherwise,  defaults to &lsquo;Managed&rsquo;. May not be changed after creation. For more information see <a href="https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os">Ephemeral
OS</a>.</p>
</td>
</tr>
<tr>
<td>
<code>osSKU</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSSKU_STATUS">
OSSKU_STATUS
</a>
</em>
</td>
<td>
<p>OsSKU: Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or
Windows2019 if  OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is
deprecated.</p>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSType_STATUS">
OSType_STATUS
</a>
</em>
</td>
<td>
<p>OsType: The operating system type. The default is Linux.</p>
</td>
</tr>
<tr>
<td>
<code>podSubnetID</code><br/>
<em>
string
</em>
</td>
<td>
<p>PodSubnetID: If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is
of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Network/&#x200b;virtualNetworks/&#x200b;{virtualNetworkName}/&#x200b;subnets/&#x200b;{subnetName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>powerState</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PowerState_STATUS">
PowerState_STATUS
</a>
</em>
</td>
<td>
<p>PowerState: When an Agent Pool is first created it is initially Running. The Agent Pool can be stopped by setting this
field to Stopped. A stopped Agent Pool stops all of its VMs and does not accrue billing charges. An Agent Pool can only
be stopped if it is Running and provisioning state is Succeeded</p>
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
<p>ProvisioningState: The current deployment or provisioning state.</p>
</td>
</tr>
<tr>
<td>
<code>proximityPlacementGroupID</code><br/>
<em>
string
</em>
</td>
<td>
<p>ProximityPlacementGroupID: The ID for Proximity Placement Group.</p>
</td>
</tr>
<tr>
<td>
<code>scaleDownMode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleDownMode_STATUS">
ScaleDownMode_STATUS
</a>
</em>
</td>
<td>
<p>ScaleDownMode: This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.</p>
</td>
</tr>
<tr>
<td>
<code>scaleSetEvictionPolicy</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleSetEvictionPolicy_STATUS">
ScaleSetEvictionPolicy_STATUS
</a>
</em>
</td>
<td>
<p>ScaleSetEvictionPolicy: This cannot be specified unless the scaleSetPriority is &lsquo;Spot&rsquo;. If not specified, the default is
&lsquo;Delete&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>scaleSetPriority</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleSetPriority_STATUS">
ScaleSetPriority_STATUS
</a>
</em>
</td>
<td>
<p>ScaleSetPriority: The Virtual Machine Scale Set priority. If not specified, the default is &lsquo;Regular&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>spotMaxPrice</code><br/>
<em>
float64
</em>
</td>
<td>
<p>SpotMaxPrice: Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any
on-demand price. For more details on spot pricing, see <a href="https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing">spot VMs
pricing</a></p>
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
<p>Tags: The tags to be persisted on the agent pool virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolType_STATUS">
AgentPoolType_STATUS
</a>
</em>
</td>
<td>
<p>Type: The type of Agent Pool.</p>
</td>
</tr>
<tr>
<td>
<code>upgradeSettings</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolUpgradeSettings_STATUS">
AgentPoolUpgradeSettings_STATUS
</a>
</em>
</td>
<td>
<p>UpgradeSettings: Settings for upgrading the agentpool</p>
</td>
</tr>
<tr>
<td>
<code>vmSize</code><br/>
<em>
string
</em>
</td>
<td>
<p>VmSize: VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods
might fail to run correctly. For more details on restricted VM sizes, see:
<a href="https://docs.microsoft.com/azure/aks/quotas-skus-regions">https://docs.microsoft.com/azure/aks/quotas-skus-regions</a></p>
</td>
</tr>
<tr>
<td>
<code>vnetSubnetID</code><br/>
<em>
string
</em>
</td>
<td>
<p>VnetSubnetID: If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified,
this applies to nodes and pods, otherwise it applies to just nodes. This is of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Network/&#x200b;virtualNetworks/&#x200b;{virtualNetworkName}/&#x200b;subnets/&#x200b;{subnetName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>windowsProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolWindowsProfile_STATUS">
AgentPoolWindowsProfile_STATUS
</a>
</em>
</td>
<td>
<p>WindowsProfile: The Windows agent pool&rsquo;s specific profile.</p>
</td>
</tr>
<tr>
<td>
<code>workloadRuntime</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.WorkloadRuntime_STATUS">
WorkloadRuntime_STATUS
</a>
</em>
</td>
<td>
<p>WorkloadRuntime: Determines the type of workload a node can run.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS_ARM">ManagedClusterAgentPoolProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>)
</p>
<div>
<p>Profile for the container service agent pool.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>availabilityZones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AvailabilityZones: The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType
property is &lsquo;VirtualMachineScaleSets&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>capacityReservationGroupID</code><br/>
<em>
string
</em>
</td>
<td>
<p>CapacityReservationGroupID: AKS will associate the specified agent pool with the Capacity Reservation Group.</p>
</td>
</tr>
<tr>
<td>
<code>count</code><br/>
<em>
int
</em>
</td>
<td>
<p>Count: Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive)
for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.</p>
</td>
</tr>
<tr>
<td>
<code>creationData</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.CreationData_STATUS_ARM">
CreationData_STATUS_ARM
</a>
</em>
</td>
<td>
<p>CreationData: CreationData to be used to specify the source Snapshot ID if the node pool will be created/upgraded using
a snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>currentOrchestratorVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>CurrentOrchestratorVersion: If orchestratorVersion was a fully specified version <major.minor.patch>, this field will be
exactly equal to it. If orchestratorVersion was <major.minor>, this field will contain the full <major.minor.patch>
version being used.</p>
</td>
</tr>
<tr>
<td>
<code>enableAutoScaling</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAutoScaling: Whether to enable auto-scaler</p>
</td>
</tr>
<tr>
<td>
<code>enableCustomCATrust</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableCustomCATrust: When set to true, AKS adds a label to the node indicating that the feature is enabled and deploys a
daemonset along with host services to sync custom certificate authorities from user-provided list of base64 encoded
certificates into node trust stores. Defaults to false.</p>
</td>
</tr>
<tr>
<td>
<code>enableEncryptionAtHost</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableEncryptionAtHost: This is only supported on certain VM sizes and in certain Azure regions. For more information,
see: <a href="https://docs.microsoft.com/azure/aks/enable-host-encryption">https://docs.microsoft.com/azure/aks/enable-host-encryption</a></p>
</td>
</tr>
<tr>
<td>
<code>enableFIPS</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableFIPS: See <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview">Add a FIPS-enabled node
pool</a> for more
details.</p>
</td>
</tr>
<tr>
<td>
<code>enableNodePublicIP</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableNodePublicIP: Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses.
A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine
to minimize hops. For more information see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools">assigning a public IP per
node</a>. The
default is false.</p>
</td>
</tr>
<tr>
<td>
<code>enableUltraSSD</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableUltraSSD: Whether to enable UltraSSD</p>
</td>
</tr>
<tr>
<td>
<code>gpuInstanceProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.GPUInstanceProfile_STATUS">
GPUInstanceProfile_STATUS
</a>
</em>
</td>
<td>
<p>GpuInstanceProfile: GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.</p>
</td>
</tr>
<tr>
<td>
<code>hostGroupID</code><br/>
<em>
string
</em>
</td>
<td>
<p>HostGroupID: This is of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Compute/&#x200b;hostGroups/&#x200b;{hostGroupName}.
For more information see <a href="https://docs.microsoft.com/azure/virtual-machines/dedicated-hosts">Azure dedicated hosts</a>.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.KubeletConfig_STATUS_ARM">
KubeletConfig_STATUS_ARM
</a>
</em>
</td>
<td>
<p>KubeletConfig: The Kubelet configuration on the agent pool nodes.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.KubeletDiskType_STATUS">
KubeletDiskType_STATUS
</a>
</em>
</td>
<td>
<p>KubeletDiskType: Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral
storage.</p>
</td>
</tr>
<tr>
<td>
<code>linuxOSConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.LinuxOSConfig_STATUS_ARM">
LinuxOSConfig_STATUS_ARM
</a>
</em>
</td>
<td>
<p>LinuxOSConfig: The OS configuration of Linux agent nodes.</p>
</td>
</tr>
<tr>
<td>
<code>maxCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxCount: The maximum number of nodes for auto-scaling</p>
</td>
</tr>
<tr>
<td>
<code>maxPods</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxPods: The maximum number of pods that can run on a node.</p>
</td>
</tr>
<tr>
<td>
<code>messageOfTheDay</code><br/>
<em>
string
</em>
</td>
<td>
<p>MessageOfTheDay: A base64-encoded string which will be written to /etc/motd after decoding. This allows customization of
the message of the day for Linux nodes. It must not be specified for Windows nodes. It must be a static string (i.e.,
will be printed raw and not be executed as a script).</p>
</td>
</tr>
<tr>
<td>
<code>minCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MinCount: The minimum number of nodes for auto-scaling</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolMode_STATUS">
AgentPoolMode_STATUS
</a>
</em>
</td>
<td>
<p>Mode: A cluster must have at least one &lsquo;System&rsquo; Agent Pool at all times. For additional information on agent pool
restrictions  and best practices, see: <a href="https://docs.microsoft.com/azure/aks/use-system-pools">https://docs.microsoft.com/azure/aks/use-system-pools</a></p>
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
<p>Name: Windows agent pool names must be 6 characters or less.</p>
</td>
</tr>
<tr>
<td>
<code>networkProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolNetworkProfile_STATUS_ARM">
AgentPoolNetworkProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>NetworkProfile: Network-related settings of an agent pool.</p>
</td>
</tr>
<tr>
<td>
<code>nodeImageVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>NodeImageVersion: The version of node image</p>
</td>
</tr>
<tr>
<td>
<code>nodeLabels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>NodeLabels: The node labels to be persisted across all nodes in agent pool.</p>
</td>
</tr>
<tr>
<td>
<code>nodePublicIPPrefixID</code><br/>
<em>
string
</em>
</td>
<td>
<p>NodePublicIPPrefixID: This is of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Network/&#x200b;publicIPPrefixes/&#x200b;{publicIPPrefixName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>nodeTaints</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>NodeTaints: The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.</p>
</td>
</tr>
<tr>
<td>
<code>orchestratorVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>OrchestratorVersion: Both patch version <major.minor.patch> and <major.minor> are supported. When <major.minor> is
specified, the latest supported patch version is chosen automatically. Updating the agent pool with the same
<major.minor> once it has been created will not trigger an upgrade, even if a newer patch version is available. As a
best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version
must have the same major version as the control plane. The node pool minor version must be within two minor versions of
the control plane version. The node pool version cannot be greater than the control plane version. For more information
see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool">upgrading a node pool</a>.</p>
</td>
</tr>
<tr>
<td>
<code>osDiskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSDiskType_STATUS">
OSDiskType_STATUS
</a>
</em>
</td>
<td>
<p>OsDiskType: The default is &lsquo;Ephemeral&rsquo; if the VM supports it and has a cache disk larger than the requested
OSDiskSizeGB. Otherwise,  defaults to &lsquo;Managed&rsquo;. May not be changed after creation. For more information see <a href="https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os">Ephemeral
OS</a>.</p>
</td>
</tr>
<tr>
<td>
<code>osSKU</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSSKU_STATUS">
OSSKU_STATUS
</a>
</em>
</td>
<td>
<p>OsSKU: Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or
Windows2019 if  OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is
deprecated.</p>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSType_STATUS">
OSType_STATUS
</a>
</em>
</td>
<td>
<p>OsType: The operating system type. The default is Linux.</p>
</td>
</tr>
<tr>
<td>
<code>podSubnetID</code><br/>
<em>
string
</em>
</td>
<td>
<p>PodSubnetID: If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is
of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Network/&#x200b;virtualNetworks/&#x200b;{virtualNetworkName}/&#x200b;subnets/&#x200b;{subnetName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>powerState</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PowerState_STATUS_ARM">
PowerState_STATUS_ARM
</a>
</em>
</td>
<td>
<p>PowerState: When an Agent Pool is first created it is initially Running. The Agent Pool can be stopped by setting this
field to Stopped. A stopped Agent Pool stops all of its VMs and does not accrue billing charges. An Agent Pool can only
be stopped if it is Running and provisioning state is Succeeded</p>
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
<p>ProvisioningState: The current deployment or provisioning state.</p>
</td>
</tr>
<tr>
<td>
<code>proximityPlacementGroupID</code><br/>
<em>
string
</em>
</td>
<td>
<p>ProximityPlacementGroupID: The ID for Proximity Placement Group.</p>
</td>
</tr>
<tr>
<td>
<code>scaleDownMode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleDownMode_STATUS">
ScaleDownMode_STATUS
</a>
</em>
</td>
<td>
<p>ScaleDownMode: This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.</p>
</td>
</tr>
<tr>
<td>
<code>scaleSetEvictionPolicy</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleSetEvictionPolicy_STATUS">
ScaleSetEvictionPolicy_STATUS
</a>
</em>
</td>
<td>
<p>ScaleSetEvictionPolicy: This cannot be specified unless the scaleSetPriority is &lsquo;Spot&rsquo;. If not specified, the default is
&lsquo;Delete&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>scaleSetPriority</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleSetPriority_STATUS">
ScaleSetPriority_STATUS
</a>
</em>
</td>
<td>
<p>ScaleSetPriority: The Virtual Machine Scale Set priority. If not specified, the default is &lsquo;Regular&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>spotMaxPrice</code><br/>
<em>
float64
</em>
</td>
<td>
<p>SpotMaxPrice: Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any
on-demand price. For more details on spot pricing, see <a href="https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing">spot VMs
pricing</a></p>
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
<p>Tags: The tags to be persisted on the agent pool virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolType_STATUS">
AgentPoolType_STATUS
</a>
</em>
</td>
<td>
<p>Type: The type of Agent Pool.</p>
</td>
</tr>
<tr>
<td>
<code>upgradeSettings</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolUpgradeSettings_STATUS_ARM">
AgentPoolUpgradeSettings_STATUS_ARM
</a>
</em>
</td>
<td>
<p>UpgradeSettings: Settings for upgrading the agentpool</p>
</td>
</tr>
<tr>
<td>
<code>vmSize</code><br/>
<em>
string
</em>
</td>
<td>
<p>VmSize: VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods
might fail to run correctly. For more details on restricted VM sizes, see:
<a href="https://docs.microsoft.com/azure/aks/quotas-skus-regions">https://docs.microsoft.com/azure/aks/quotas-skus-regions</a></p>
</td>
</tr>
<tr>
<td>
<code>vnetSubnetID</code><br/>
<em>
string
</em>
</td>
<td>
<p>VnetSubnetID: If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified,
this applies to nodes and pods, otherwise it applies to just nodes. This is of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Network/&#x200b;virtualNetworks/&#x200b;{virtualNetworkName}/&#x200b;subnets/&#x200b;{subnetName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>windowsProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolWindowsProfile_STATUS_ARM">
AgentPoolWindowsProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>WindowsProfile: The Windows agent pool&rsquo;s specific profile.</p>
</td>
</tr>
<tr>
<td>
<code>workloadRuntime</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.WorkloadRuntime_STATUS">
WorkloadRuntime_STATUS
</a>
</em>
</td>
<td>
<p>WorkloadRuntime: Determines the type of workload a node can run.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile">ManagedClusterAutoUpgradeProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
</p>
<div>
<p>Auto upgrade profile for a managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>nodeOSUpgradeChannel</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel">
ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel
</a>
</em>
</td>
<td>
<p>NodeOSUpgradeChannel: The default is Unmanaged, but may change to either NodeImage or SecurityPatch at GA.</p>
</td>
</tr>
<tr>
<td>
<code>upgradeChannel</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile_UpgradeChannel">
ManagedClusterAutoUpgradeProfile_UpgradeChannel
</a>
</em>
</td>
<td>
<p>UpgradeChannel: For more information see <a href="https://docs.microsoft.com/azure/aks/upgrade-cluster#set-auto-upgrade-channel">setting the AKS cluster auto-upgrade
channel</a>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile_ARM">ManagedClusterAutoUpgradeProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM</a>)
</p>
<div>
<p>Auto upgrade profile for a managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>nodeOSUpgradeChannel</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel">
ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel
</a>
</em>
</td>
<td>
<p>NodeOSUpgradeChannel: The default is Unmanaged, but may change to either NodeImage or SecurityPatch at GA.</p>
</td>
</tr>
<tr>
<td>
<code>upgradeChannel</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile_UpgradeChannel">
ManagedClusterAutoUpgradeProfile_UpgradeChannel
</a>
</em>
</td>
<td>
<p>UpgradeChannel: For more information see <a href="https://docs.microsoft.com/azure/aks/upgrade-cluster#set-auto-upgrade-channel">setting the AKS cluster auto-upgrade
channel</a>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel">ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile">ManagedClusterAutoUpgradeProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile_ARM">ManagedClusterAutoUpgradeProfile_ARM</a>)
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
<tbody><tr><td><p>&#34;NodeImage&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SecurityPatch&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Unmanaged&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel_STATUS">ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile_STATUS">ManagedClusterAutoUpgradeProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile_STATUS_ARM">ManagedClusterAutoUpgradeProfile_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;NodeImage&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SecurityPatch&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Unmanaged&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile_STATUS">ManagedClusterAutoUpgradeProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
</p>
<div>
<p>Auto upgrade profile for a managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>nodeOSUpgradeChannel</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel_STATUS">
ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel_STATUS
</a>
</em>
</td>
<td>
<p>NodeOSUpgradeChannel: The default is Unmanaged, but may change to either NodeImage or SecurityPatch at GA.</p>
</td>
</tr>
<tr>
<td>
<code>upgradeChannel</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile_UpgradeChannel_STATUS">
ManagedClusterAutoUpgradeProfile_UpgradeChannel_STATUS
</a>
</em>
</td>
<td>
<p>UpgradeChannel: For more information see <a href="https://docs.microsoft.com/azure/aks/upgrade-cluster#set-auto-upgrade-channel">setting the AKS cluster auto-upgrade
channel</a>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile_STATUS_ARM">ManagedClusterAutoUpgradeProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>)
</p>
<div>
<p>Auto upgrade profile for a managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>nodeOSUpgradeChannel</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel_STATUS">
ManagedClusterAutoUpgradeProfile_NodeOSUpgradeChannel_STATUS
</a>
</em>
</td>
<td>
<p>NodeOSUpgradeChannel: The default is Unmanaged, but may change to either NodeImage or SecurityPatch at GA.</p>
</td>
</tr>
<tr>
<td>
<code>upgradeChannel</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile_UpgradeChannel_STATUS">
ManagedClusterAutoUpgradeProfile_UpgradeChannel_STATUS
</a>
</em>
</td>
<td>
<p>UpgradeChannel: For more information see <a href="https://docs.microsoft.com/azure/aks/upgrade-cluster#set-auto-upgrade-channel">setting the AKS cluster auto-upgrade
channel</a>.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile_UpgradeChannel">ManagedClusterAutoUpgradeProfile_UpgradeChannel
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile">ManagedClusterAutoUpgradeProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile_ARM">ManagedClusterAutoUpgradeProfile_ARM</a>)
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
<tbody><tr><td><p>&#34;node-image&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;none&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;patch&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;rapid&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;stable&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile_UpgradeChannel_STATUS">ManagedClusterAutoUpgradeProfile_UpgradeChannel_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile_STATUS">ManagedClusterAutoUpgradeProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile_STATUS_ARM">ManagedClusterAutoUpgradeProfile_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;node-image&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;none&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;patch&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;rapid&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;stable&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfile">ManagedClusterAzureMonitorProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
</p>
<div>
<p>Prometheus addon profile for the container service cluster</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metrics</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfileMetrics">
ManagedClusterAzureMonitorProfileMetrics
</a>
</em>
</td>
<td>
<p>Metrics: Metrics profile for the prometheus service addon</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfileKubeStateMetrics">ManagedClusterAzureMonitorProfileKubeStateMetrics
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfileMetrics">ManagedClusterAzureMonitorProfileMetrics</a>)
</p>
<div>
<p>Kube State Metrics for prometheus addon profile for the container service cluster</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metricAnnotationsAllowList</code><br/>
<em>
string
</em>
</td>
<td>
<p>MetricAnnotationsAllowList: Comma-separated list of additional Kubernetes label keys that will be used in the resource&rsquo;s
labels metric.</p>
</td>
</tr>
<tr>
<td>
<code>metricLabelsAllowlist</code><br/>
<em>
string
</em>
</td>
<td>
<p>MetricLabelsAllowlist: Comma-separated list of Kubernetes annotations keys that will be used in the resource&rsquo;s labels
metric.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfileKubeStateMetrics_ARM">ManagedClusterAzureMonitorProfileKubeStateMetrics_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfileMetrics_ARM">ManagedClusterAzureMonitorProfileMetrics_ARM</a>)
</p>
<div>
<p>Kube State Metrics for prometheus addon profile for the container service cluster</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metricAnnotationsAllowList</code><br/>
<em>
string
</em>
</td>
<td>
<p>MetricAnnotationsAllowList: Comma-separated list of additional Kubernetes label keys that will be used in the resource&rsquo;s
labels metric.</p>
</td>
</tr>
<tr>
<td>
<code>metricLabelsAllowlist</code><br/>
<em>
string
</em>
</td>
<td>
<p>MetricLabelsAllowlist: Comma-separated list of Kubernetes annotations keys that will be used in the resource&rsquo;s labels
metric.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfileKubeStateMetrics_STATUS">ManagedClusterAzureMonitorProfileKubeStateMetrics_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfileMetrics_STATUS">ManagedClusterAzureMonitorProfileMetrics_STATUS</a>)
</p>
<div>
<p>Kube State Metrics for prometheus addon profile for the container service cluster</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metricAnnotationsAllowList</code><br/>
<em>
string
</em>
</td>
<td>
<p>MetricAnnotationsAllowList: Comma-separated list of additional Kubernetes label keys that will be used in the resource&rsquo;s
labels metric.</p>
</td>
</tr>
<tr>
<td>
<code>metricLabelsAllowlist</code><br/>
<em>
string
</em>
</td>
<td>
<p>MetricLabelsAllowlist: Comma-separated list of Kubernetes annotations keys that will be used in the resource&rsquo;s labels
metric.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfileKubeStateMetrics_STATUS_ARM">ManagedClusterAzureMonitorProfileKubeStateMetrics_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfileMetrics_STATUS_ARM">ManagedClusterAzureMonitorProfileMetrics_STATUS_ARM</a>)
</p>
<div>
<p>Kube State Metrics for prometheus addon profile for the container service cluster</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metricAnnotationsAllowList</code><br/>
<em>
string
</em>
</td>
<td>
<p>MetricAnnotationsAllowList: Comma-separated list of additional Kubernetes label keys that will be used in the resource&rsquo;s
labels metric.</p>
</td>
</tr>
<tr>
<td>
<code>metricLabelsAllowlist</code><br/>
<em>
string
</em>
</td>
<td>
<p>MetricLabelsAllowlist: Comma-separated list of Kubernetes annotations keys that will be used in the resource&rsquo;s labels
metric.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfileMetrics">ManagedClusterAzureMonitorProfileMetrics
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfile">ManagedClusterAzureMonitorProfile</a>)
</p>
<div>
<p>Metrics profile for the prometheus service addon</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable the Prometheus collector</p>
</td>
</tr>
<tr>
<td>
<code>kubeStateMetrics</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfileKubeStateMetrics">
ManagedClusterAzureMonitorProfileKubeStateMetrics
</a>
</em>
</td>
<td>
<p>KubeStateMetrics: Kube State Metrics for prometheus addon profile for the container service cluster</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfileMetrics_ARM">ManagedClusterAzureMonitorProfileMetrics_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfile_ARM">ManagedClusterAzureMonitorProfile_ARM</a>)
</p>
<div>
<p>Metrics profile for the prometheus service addon</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable the Prometheus collector</p>
</td>
</tr>
<tr>
<td>
<code>kubeStateMetrics</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfileKubeStateMetrics_ARM">
ManagedClusterAzureMonitorProfileKubeStateMetrics_ARM
</a>
</em>
</td>
<td>
<p>KubeStateMetrics: Kube State Metrics for prometheus addon profile for the container service cluster</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfileMetrics_STATUS">ManagedClusterAzureMonitorProfileMetrics_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfile_STATUS">ManagedClusterAzureMonitorProfile_STATUS</a>)
</p>
<div>
<p>Metrics profile for the prometheus service addon</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable the Prometheus collector</p>
</td>
</tr>
<tr>
<td>
<code>kubeStateMetrics</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfileKubeStateMetrics_STATUS">
ManagedClusterAzureMonitorProfileKubeStateMetrics_STATUS
</a>
</em>
</td>
<td>
<p>KubeStateMetrics: Kube State Metrics for prometheus addon profile for the container service cluster</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfileMetrics_STATUS_ARM">ManagedClusterAzureMonitorProfileMetrics_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfile_STATUS_ARM">ManagedClusterAzureMonitorProfile_STATUS_ARM</a>)
</p>
<div>
<p>Metrics profile for the prometheus service addon</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable the Prometheus collector</p>
</td>
</tr>
<tr>
<td>
<code>kubeStateMetrics</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfileKubeStateMetrics_STATUS_ARM">
ManagedClusterAzureMonitorProfileKubeStateMetrics_STATUS_ARM
</a>
</em>
</td>
<td>
<p>KubeStateMetrics: Kube State Metrics for prometheus addon profile for the container service cluster</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfile_ARM">ManagedClusterAzureMonitorProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM</a>)
</p>
<div>
<p>Prometheus addon profile for the container service cluster</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metrics</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfileMetrics_ARM">
ManagedClusterAzureMonitorProfileMetrics_ARM
</a>
</em>
</td>
<td>
<p>Metrics: Metrics profile for the prometheus service addon</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfile_STATUS">ManagedClusterAzureMonitorProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
</p>
<div>
<p>Prometheus addon profile for the container service cluster</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metrics</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfileMetrics_STATUS">
ManagedClusterAzureMonitorProfileMetrics_STATUS
</a>
</em>
</td>
<td>
<p>Metrics: Metrics profile for the prometheus service addon</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfile_STATUS_ARM">ManagedClusterAzureMonitorProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>)
</p>
<div>
<p>Prometheus addon profile for the container service cluster</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metrics</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfileMetrics_STATUS_ARM">
ManagedClusterAzureMonitorProfileMetrics_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Metrics: Metrics profile for the prometheus service addon</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterHTTPProxyConfig">ManagedClusterHTTPProxyConfig
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
</p>
<div>
<p>Cluster HTTP proxy configuration.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>httpProxy</code><br/>
<em>
string
</em>
</td>
<td>
<p>HttpProxy: The HTTP proxy server endpoint to use.</p>
</td>
</tr>
<tr>
<td>
<code>httpsProxy</code><br/>
<em>
string
</em>
</td>
<td>
<p>HttpsProxy: The HTTPS proxy server endpoint to use.</p>
</td>
</tr>
<tr>
<td>
<code>noProxy</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>NoProxy: The endpoints that should not go through proxy.</p>
</td>
</tr>
<tr>
<td>
<code>trustedCa</code><br/>
<em>
string
</em>
</td>
<td>
<p>TrustedCa: Alternative CA cert to use for connecting to proxy servers.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterHTTPProxyConfig_ARM">ManagedClusterHTTPProxyConfig_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM</a>)
</p>
<div>
<p>Cluster HTTP proxy configuration.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>httpProxy</code><br/>
<em>
string
</em>
</td>
<td>
<p>HttpProxy: The HTTP proxy server endpoint to use.</p>
</td>
</tr>
<tr>
<td>
<code>httpsProxy</code><br/>
<em>
string
</em>
</td>
<td>
<p>HttpsProxy: The HTTPS proxy server endpoint to use.</p>
</td>
</tr>
<tr>
<td>
<code>noProxy</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>NoProxy: The endpoints that should not go through proxy.</p>
</td>
</tr>
<tr>
<td>
<code>trustedCa</code><br/>
<em>
string
</em>
</td>
<td>
<p>TrustedCa: Alternative CA cert to use for connecting to proxy servers.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterHTTPProxyConfig_STATUS">ManagedClusterHTTPProxyConfig_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
</p>
<div>
<p>Cluster HTTP proxy configuration.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>effectiveNoProxy</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>EffectiveNoProxy: A read-only list of all endpoints for which traffic should not be sent to the proxy. This list is a
superset of noProxy and values injected by AKS.</p>
</td>
</tr>
<tr>
<td>
<code>httpProxy</code><br/>
<em>
string
</em>
</td>
<td>
<p>HttpProxy: The HTTP proxy server endpoint to use.</p>
</td>
</tr>
<tr>
<td>
<code>httpsProxy</code><br/>
<em>
string
</em>
</td>
<td>
<p>HttpsProxy: The HTTPS proxy server endpoint to use.</p>
</td>
</tr>
<tr>
<td>
<code>noProxy</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>NoProxy: The endpoints that should not go through proxy.</p>
</td>
</tr>
<tr>
<td>
<code>trustedCa</code><br/>
<em>
string
</em>
</td>
<td>
<p>TrustedCa: Alternative CA cert to use for connecting to proxy servers.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterHTTPProxyConfig_STATUS_ARM">ManagedClusterHTTPProxyConfig_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>)
</p>
<div>
<p>Cluster HTTP proxy configuration.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>effectiveNoProxy</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>EffectiveNoProxy: A read-only list of all endpoints for which traffic should not be sent to the proxy. This list is a
superset of noProxy and values injected by AKS.</p>
</td>
</tr>
<tr>
<td>
<code>httpProxy</code><br/>
<em>
string
</em>
</td>
<td>
<p>HttpProxy: The HTTP proxy server endpoint to use.</p>
</td>
</tr>
<tr>
<td>
<code>httpsProxy</code><br/>
<em>
string
</em>
</td>
<td>
<p>HttpsProxy: The HTTPS proxy server endpoint to use.</p>
</td>
</tr>
<tr>
<td>
<code>noProxy</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>NoProxy: The endpoints that should not go through proxy.</p>
</td>
</tr>
<tr>
<td>
<code>trustedCa</code><br/>
<em>
string
</em>
</td>
<td>
<p>TrustedCa: Alternative CA cert to use for connecting to proxy servers.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity">ManagedClusterIdentity
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
</p>
<div>
<p>Identity for the managed cluster.</p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity_Type">
ManagedClusterIdentity_Type
</a>
</em>
</td>
<td>
<p>Type: For more information see <a href="https://docs.microsoft.com/azure/aks/use-managed-identity">use managed identities in
AKS</a>.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.UserAssignedIdentityDetails">
[]UserAssignedIdentityDetails
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: The keys must be ARM resource IDs in the form:
&lsquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ManagedIdentity/&#x200b;userAssignedIdentities/&#x200b;{identityName}&rsquo;.</&#x200b;p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity_ARM">ManagedClusterIdentity_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec_ARM">ManagedCluster_Spec_ARM</a>)
</p>
<div>
<p>Identity for the managed cluster.</p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity_Type">
ManagedClusterIdentity_Type
</a>
</em>
</td>
<td>
<p>Type: For more information see <a href="https://docs.microsoft.com/azure/aks/use-managed-identity">use managed identities in
AKS</a>.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.UserAssignedIdentityDetails_ARM">
map[string]./api/containerservice/v1api20230202preview.UserAssignedIdentityDetails_ARM
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity_STATUS">ManagedClusterIdentity_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
</p>
<div>
<p>Identity for the managed cluster.</p>
</div>
<table>
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
<p>PrincipalId: The principal id of the system assigned identity which is used by master components.</p>
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
<p>TenantId: The tenant id of the system assigned identity which is used by master components.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity_Type_STATUS">
ManagedClusterIdentity_Type_STATUS
</a>
</em>
</td>
<td>
<p>Type: For more information see <a href="https://docs.microsoft.com/azure/aks/use-managed-identity">use managed identities in
AKS</a>.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity_UserAssignedIdentities_STATUS">
map[string]./api/containerservice/v1api20230202preview.ManagedClusterIdentity_UserAssignedIdentities_STATUS
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: The keys must be ARM resource IDs in the form:
&lsquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ManagedIdentity/&#x200b;userAssignedIdentities/&#x200b;{identityName}&rsquo;.</&#x200b;p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity_STATUS_ARM">ManagedClusterIdentity_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS_ARM">ManagedCluster_STATUS_ARM</a>)
</p>
<div>
<p>Identity for the managed cluster.</p>
</div>
<table>
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
<p>PrincipalId: The principal id of the system assigned identity which is used by master components.</p>
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
<p>TenantId: The tenant id of the system assigned identity which is used by master components.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity_Type_STATUS">
ManagedClusterIdentity_Type_STATUS
</a>
</em>
</td>
<td>
<p>Type: For more information see <a href="https://docs.microsoft.com/azure/aks/use-managed-identity">use managed identities in
AKS</a>.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity_UserAssignedIdentities_STATUS_ARM">
map[string]./api/containerservice/v1api20230202preview.ManagedClusterIdentity_UserAssignedIdentities_STATUS_ARM
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: The keys must be ARM resource IDs in the form:
&lsquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ManagedIdentity/&#x200b;userAssignedIdentities/&#x200b;{identityName}&rsquo;.</&#x200b;p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity_Type">ManagedClusterIdentity_Type
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity">ManagedClusterIdentity</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity_ARM">ManagedClusterIdentity_ARM</a>)
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
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity_Type_STATUS">ManagedClusterIdentity_Type_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity_STATUS">ManagedClusterIdentity_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity_STATUS_ARM">ManagedClusterIdentity_STATUS_ARM</a>)
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
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity_UserAssignedIdentities_STATUS">ManagedClusterIdentity_UserAssignedIdentities_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity_STATUS">ManagedClusterIdentity_STATUS</a>)
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
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity_UserAssignedIdentities_STATUS_ARM">ManagedClusterIdentity_UserAssignedIdentities_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity_STATUS_ARM">ManagedClusterIdentity_STATUS_ARM</a>)
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
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterIngressProfile">ManagedClusterIngressProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
</p>
<div>
<p>Ingress profile for the container service cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>webAppRouting</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIngressProfileWebAppRouting">
ManagedClusterIngressProfileWebAppRouting
</a>
</em>
</td>
<td>
<p>WebAppRouting: Web App Routing settings for the ingress profile.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterIngressProfileWebAppRouting">ManagedClusterIngressProfileWebAppRouting
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIngressProfile">ManagedClusterIngressProfile</a>)
</p>
<div>
<p>Web App Routing settings for the ingress profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dnsZoneResourceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>DnsZoneResourceReference: Resource ID of the DNS Zone to be associated with the web app. Used only when Web App Routing
is enabled.</p>
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
<p>Enabled: Whether to enable Web App Routing.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterIngressProfileWebAppRouting_ARM">ManagedClusterIngressProfileWebAppRouting_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIngressProfile_ARM">ManagedClusterIngressProfile_ARM</a>)
</p>
<div>
<p>Web App Routing settings for the ingress profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dnsZoneResourceId</code><br/>
<em>
string
</em>
</td>
<td>
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
<p>Enabled: Whether to enable Web App Routing.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterIngressProfileWebAppRouting_STATUS">ManagedClusterIngressProfileWebAppRouting_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIngressProfile_STATUS">ManagedClusterIngressProfile_STATUS</a>)
</p>
<div>
<p>Web App Routing settings for the ingress profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dnsZoneResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>DnsZoneResourceId: Resource ID of the DNS Zone to be associated with the web app. Used only when Web App Routing is
enabled.</p>
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
<p>Enabled: Whether to enable Web App Routing.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.UserAssignedIdentity_STATUS">
UserAssignedIdentity_STATUS
</a>
</em>
</td>
<td>
<p>Identity: Managed identity of the Web Application Routing add-on. This is the identity that should be granted
permissions, for example, to manage the associated Azure DNS resource and get certificates from Azure Key Vault. See
<a href="https://learn.microsoft.com/en-us/azure/aks/web-app-routing?tabs=with-osm">this overview of the add-on</a> for more
instructions.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterIngressProfileWebAppRouting_STATUS_ARM">ManagedClusterIngressProfileWebAppRouting_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIngressProfile_STATUS_ARM">ManagedClusterIngressProfile_STATUS_ARM</a>)
</p>
<div>
<p>Web App Routing settings for the ingress profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dnsZoneResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>DnsZoneResourceId: Resource ID of the DNS Zone to be associated with the web app. Used only when Web App Routing is
enabled.</p>
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
<p>Enabled: Whether to enable Web App Routing.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.UserAssignedIdentity_STATUS_ARM">
UserAssignedIdentity_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Identity: Managed identity of the Web Application Routing add-on. This is the identity that should be granted
permissions, for example, to manage the associated Azure DNS resource and get certificates from Azure Key Vault. See
<a href="https://learn.microsoft.com/en-us/azure/aks/web-app-routing?tabs=with-osm">this overview of the add-on</a> for more
instructions.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterIngressProfile_ARM">ManagedClusterIngressProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM</a>)
</p>
<div>
<p>Ingress profile for the container service cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>webAppRouting</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIngressProfileWebAppRouting_ARM">
ManagedClusterIngressProfileWebAppRouting_ARM
</a>
</em>
</td>
<td>
<p>WebAppRouting: Web App Routing settings for the ingress profile.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterIngressProfile_STATUS">ManagedClusterIngressProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
</p>
<div>
<p>Ingress profile for the container service cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>webAppRouting</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIngressProfileWebAppRouting_STATUS">
ManagedClusterIngressProfileWebAppRouting_STATUS
</a>
</em>
</td>
<td>
<p>WebAppRouting: Web App Routing settings for the ingress profile.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterIngressProfile_STATUS_ARM">ManagedClusterIngressProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>)
</p>
<div>
<p>Ingress profile for the container service cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>webAppRouting</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIngressProfileWebAppRouting_STATUS_ARM">
ManagedClusterIngressProfileWebAppRouting_STATUS_ARM
</a>
</em>
</td>
<td>
<p>WebAppRouting: Web App Routing settings for the ingress profile.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile">ManagedClusterLoadBalancerProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile">ContainerServiceNetworkProfile</a>)
</p>
<div>
<p>Profile of the managed cluster load balancer.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allocatedOutboundPorts</code><br/>
<em>
int
</em>
</td>
<td>
<p>AllocatedOutboundPorts: The desired number of allocated SNAT ports per VM. Allowed values are in the range of 0 to 64000
(inclusive). The default value is 0 which results in Azure dynamically allocating ports.</p>
</td>
</tr>
<tr>
<td>
<code>backendPoolType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_BackendPoolType">
ManagedClusterLoadBalancerProfile_BackendPoolType
</a>
</em>
</td>
<td>
<p>BackendPoolType: The type of the managed inbound Load Balancer BackendPool.</p>
</td>
</tr>
<tr>
<td>
<code>effectiveOutboundIPs</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ResourceReference">
[]ResourceReference
</a>
</em>
</td>
<td>
<p>EffectiveOutboundIPs: The effective outbound IP resources of the cluster load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>enableMultipleStandardLoadBalancers</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableMultipleStandardLoadBalancers: Enable multiple standard load balancers per AKS cluster or not.</p>
</td>
</tr>
<tr>
<td>
<code>idleTimeoutInMinutes</code><br/>
<em>
int
</em>
</td>
<td>
<p>IdleTimeoutInMinutes: Desired outbound flow idle timeout in minutes. Allowed values are in the range of 4 to 120
(inclusive). The default value is 30 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>managedOutboundIPs</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_ManagedOutboundIPs">
ManagedClusterLoadBalancerProfile_ManagedOutboundIPs
</a>
</em>
</td>
<td>
<p>ManagedOutboundIPs: Desired managed outbound IPs for the cluster load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>outboundIPPrefixes</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_OutboundIPPrefixes">
ManagedClusterLoadBalancerProfile_OutboundIPPrefixes
</a>
</em>
</td>
<td>
<p>OutboundIPPrefixes: Desired outbound IP Prefix resources for the cluster load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>outboundIPs</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_OutboundIPs">
ManagedClusterLoadBalancerProfile_OutboundIPs
</a>
</em>
</td>
<td>
<p>OutboundIPs: Desired outbound IP resources for the cluster load balancer.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_ARM">ManagedClusterLoadBalancerProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_ARM">ContainerServiceNetworkProfile_ARM</a>)
</p>
<div>
<p>Profile of the managed cluster load balancer.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allocatedOutboundPorts</code><br/>
<em>
int
</em>
</td>
<td>
<p>AllocatedOutboundPorts: The desired number of allocated SNAT ports per VM. Allowed values are in the range of 0 to 64000
(inclusive). The default value is 0 which results in Azure dynamically allocating ports.</p>
</td>
</tr>
<tr>
<td>
<code>backendPoolType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_BackendPoolType">
ManagedClusterLoadBalancerProfile_BackendPoolType
</a>
</em>
</td>
<td>
<p>BackendPoolType: The type of the managed inbound Load Balancer BackendPool.</p>
</td>
</tr>
<tr>
<td>
<code>effectiveOutboundIPs</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ResourceReference_ARM">
[]ResourceReference_ARM
</a>
</em>
</td>
<td>
<p>EffectiveOutboundIPs: The effective outbound IP resources of the cluster load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>enableMultipleStandardLoadBalancers</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableMultipleStandardLoadBalancers: Enable multiple standard load balancers per AKS cluster or not.</p>
</td>
</tr>
<tr>
<td>
<code>idleTimeoutInMinutes</code><br/>
<em>
int
</em>
</td>
<td>
<p>IdleTimeoutInMinutes: Desired outbound flow idle timeout in minutes. Allowed values are in the range of 4 to 120
(inclusive). The default value is 30 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>managedOutboundIPs</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_ARM">
ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_ARM
</a>
</em>
</td>
<td>
<p>ManagedOutboundIPs: Desired managed outbound IPs for the cluster load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>outboundIPPrefixes</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_ARM">
ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_ARM
</a>
</em>
</td>
<td>
<p>OutboundIPPrefixes: Desired outbound IP Prefix resources for the cluster load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>outboundIPs</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_OutboundIPs_ARM">
ManagedClusterLoadBalancerProfile_OutboundIPs_ARM
</a>
</em>
</td>
<td>
<p>OutboundIPs: Desired outbound IP resources for the cluster load balancer.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_BackendPoolType">ManagedClusterLoadBalancerProfile_BackendPoolType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile">ManagedClusterLoadBalancerProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_ARM">ManagedClusterLoadBalancerProfile_ARM</a>)
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
<tbody><tr><td><p>&#34;NodeIP&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;NodeIPConfiguration&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_BackendPoolType_STATUS">ManagedClusterLoadBalancerProfile_BackendPoolType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_STATUS">ManagedClusterLoadBalancerProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_STATUS_ARM">ManagedClusterLoadBalancerProfile_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;NodeIP&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;NodeIPConfiguration&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_ManagedOutboundIPs">ManagedClusterLoadBalancerProfile_ManagedOutboundIPs
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile">ManagedClusterLoadBalancerProfile</a>)
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
<code>count</code><br/>
<em>
int
</em>
</td>
<td>
<p>Count: The desired number of IPv4 outbound IPs created/managed by Azure for the cluster load balancer. Allowed values
must be in the range of 1 to 100 (inclusive). The default value is 1.</p>
</td>
</tr>
<tr>
<td>
<code>countIPv6</code><br/>
<em>
int
</em>
</td>
<td>
<p>CountIPv6: The desired number of IPv6 outbound IPs created/managed by Azure for the cluster load balancer. Allowed
values must be in the range of 1 to 100 (inclusive). The default value is 0 for single-stack and 1 for dual-stack.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_ARM">ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_ARM">ManagedClusterLoadBalancerProfile_ARM</a>)
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
<code>count</code><br/>
<em>
int
</em>
</td>
<td>
<p>Count: The desired number of IPv4 outbound IPs created/managed by Azure for the cluster load balancer. Allowed values
must be in the range of 1 to 100 (inclusive). The default value is 1.</p>
</td>
</tr>
<tr>
<td>
<code>countIPv6</code><br/>
<em>
int
</em>
</td>
<td>
<p>CountIPv6: The desired number of IPv6 outbound IPs created/managed by Azure for the cluster load balancer. Allowed
values must be in the range of 1 to 100 (inclusive). The default value is 0 for single-stack and 1 for dual-stack.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS">ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_STATUS">ManagedClusterLoadBalancerProfile_STATUS</a>)
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
<code>count</code><br/>
<em>
int
</em>
</td>
<td>
<p>Count: The desired number of IPv4 outbound IPs created/managed by Azure for the cluster load balancer. Allowed values
must be in the range of 1 to 100 (inclusive). The default value is 1.</p>
</td>
</tr>
<tr>
<td>
<code>countIPv6</code><br/>
<em>
int
</em>
</td>
<td>
<p>CountIPv6: The desired number of IPv6 outbound IPs created/managed by Azure for the cluster load balancer. Allowed
values must be in the range of 1 to 100 (inclusive). The default value is 0 for single-stack and 1 for dual-stack.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS_ARM">ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_STATUS_ARM">ManagedClusterLoadBalancerProfile_STATUS_ARM</a>)
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
<code>count</code><br/>
<em>
int
</em>
</td>
<td>
<p>Count: The desired number of IPv4 outbound IPs created/managed by Azure for the cluster load balancer. Allowed values
must be in the range of 1 to 100 (inclusive). The default value is 1.</p>
</td>
</tr>
<tr>
<td>
<code>countIPv6</code><br/>
<em>
int
</em>
</td>
<td>
<p>CountIPv6: The desired number of IPv6 outbound IPs created/managed by Azure for the cluster load balancer. Allowed
values must be in the range of 1 to 100 (inclusive). The default value is 0 for single-stack and 1 for dual-stack.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_OutboundIPPrefixes">ManagedClusterLoadBalancerProfile_OutboundIPPrefixes
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile">ManagedClusterLoadBalancerProfile</a>)
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
<code>publicIPPrefixes</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ResourceReference">
[]ResourceReference
</a>
</em>
</td>
<td>
<p>PublicIPPrefixes: A list of public IP prefix resources.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_ARM">ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_ARM">ManagedClusterLoadBalancerProfile_ARM</a>)
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
<code>publicIPPrefixes</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ResourceReference_ARM">
[]ResourceReference_ARM
</a>
</em>
</td>
<td>
<p>PublicIPPrefixes: A list of public IP prefix resources.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS">ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_STATUS">ManagedClusterLoadBalancerProfile_STATUS</a>)
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
<code>publicIPPrefixes</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ResourceReference_STATUS">
[]ResourceReference_STATUS
</a>
</em>
</td>
<td>
<p>PublicIPPrefixes: A list of public IP prefix resources.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS_ARM">ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_STATUS_ARM">ManagedClusterLoadBalancerProfile_STATUS_ARM</a>)
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
<code>publicIPPrefixes</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ResourceReference_STATUS_ARM">
[]ResourceReference_STATUS_ARM
</a>
</em>
</td>
<td>
<p>PublicIPPrefixes: A list of public IP prefix resources.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_OutboundIPs">ManagedClusterLoadBalancerProfile_OutboundIPs
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile">ManagedClusterLoadBalancerProfile</a>)
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
<code>publicIPs</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ResourceReference">
[]ResourceReference
</a>
</em>
</td>
<td>
<p>PublicIPs: A list of public IP resources.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_OutboundIPs_ARM">ManagedClusterLoadBalancerProfile_OutboundIPs_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_ARM">ManagedClusterLoadBalancerProfile_ARM</a>)
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
<code>publicIPs</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ResourceReference_ARM">
[]ResourceReference_ARM
</a>
</em>
</td>
<td>
<p>PublicIPs: A list of public IP resources.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS">ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_STATUS">ManagedClusterLoadBalancerProfile_STATUS</a>)
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
<code>publicIPs</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ResourceReference_STATUS">
[]ResourceReference_STATUS
</a>
</em>
</td>
<td>
<p>PublicIPs: A list of public IP resources.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS_ARM">ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_STATUS_ARM">ManagedClusterLoadBalancerProfile_STATUS_ARM</a>)
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
<code>publicIPs</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ResourceReference_STATUS_ARM">
[]ResourceReference_STATUS_ARM
</a>
</em>
</td>
<td>
<p>PublicIPs: A list of public IP resources.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_STATUS">ManagedClusterLoadBalancerProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_STATUS">ContainerServiceNetworkProfile_STATUS</a>)
</p>
<div>
<p>Profile of the managed cluster load balancer.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allocatedOutboundPorts</code><br/>
<em>
int
</em>
</td>
<td>
<p>AllocatedOutboundPorts: The desired number of allocated SNAT ports per VM. Allowed values are in the range of 0 to 64000
(inclusive). The default value is 0 which results in Azure dynamically allocating ports.</p>
</td>
</tr>
<tr>
<td>
<code>backendPoolType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_BackendPoolType_STATUS">
ManagedClusterLoadBalancerProfile_BackendPoolType_STATUS
</a>
</em>
</td>
<td>
<p>BackendPoolType: The type of the managed inbound Load Balancer BackendPool.</p>
</td>
</tr>
<tr>
<td>
<code>effectiveOutboundIPs</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ResourceReference_STATUS">
[]ResourceReference_STATUS
</a>
</em>
</td>
<td>
<p>EffectiveOutboundIPs: The effective outbound IP resources of the cluster load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>enableMultipleStandardLoadBalancers</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableMultipleStandardLoadBalancers: Enable multiple standard load balancers per AKS cluster or not.</p>
</td>
</tr>
<tr>
<td>
<code>idleTimeoutInMinutes</code><br/>
<em>
int
</em>
</td>
<td>
<p>IdleTimeoutInMinutes: Desired outbound flow idle timeout in minutes. Allowed values are in the range of 4 to 120
(inclusive). The default value is 30 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>managedOutboundIPs</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS">
ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS
</a>
</em>
</td>
<td>
<p>ManagedOutboundIPs: Desired managed outbound IPs for the cluster load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>outboundIPPrefixes</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS">
ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS
</a>
</em>
</td>
<td>
<p>OutboundIPPrefixes: Desired outbound IP Prefix resources for the cluster load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>outboundIPs</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS">
ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS
</a>
</em>
</td>
<td>
<p>OutboundIPs: Desired outbound IP resources for the cluster load balancer.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_STATUS_ARM">ManagedClusterLoadBalancerProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_STATUS_ARM">ContainerServiceNetworkProfile_STATUS_ARM</a>)
</p>
<div>
<p>Profile of the managed cluster load balancer.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allocatedOutboundPorts</code><br/>
<em>
int
</em>
</td>
<td>
<p>AllocatedOutboundPorts: The desired number of allocated SNAT ports per VM. Allowed values are in the range of 0 to 64000
(inclusive). The default value is 0 which results in Azure dynamically allocating ports.</p>
</td>
</tr>
<tr>
<td>
<code>backendPoolType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_BackendPoolType_STATUS">
ManagedClusterLoadBalancerProfile_BackendPoolType_STATUS
</a>
</em>
</td>
<td>
<p>BackendPoolType: The type of the managed inbound Load Balancer BackendPool.</p>
</td>
</tr>
<tr>
<td>
<code>effectiveOutboundIPs</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ResourceReference_STATUS_ARM">
[]ResourceReference_STATUS_ARM
</a>
</em>
</td>
<td>
<p>EffectiveOutboundIPs: The effective outbound IP resources of the cluster load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>enableMultipleStandardLoadBalancers</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableMultipleStandardLoadBalancers: Enable multiple standard load balancers per AKS cluster or not.</p>
</td>
</tr>
<tr>
<td>
<code>idleTimeoutInMinutes</code><br/>
<em>
int
</em>
</td>
<td>
<p>IdleTimeoutInMinutes: Desired outbound flow idle timeout in minutes. Allowed values are in the range of 4 to 120
(inclusive). The default value is 30 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>managedOutboundIPs</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS_ARM">
ManagedClusterLoadBalancerProfile_ManagedOutboundIPs_STATUS_ARM
</a>
</em>
</td>
<td>
<p>ManagedOutboundIPs: Desired managed outbound IPs for the cluster load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>outboundIPPrefixes</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS_ARM">
ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS_ARM
</a>
</em>
</td>
<td>
<p>OutboundIPPrefixes: Desired outbound IP Prefix resources for the cluster load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>outboundIPs</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS_ARM">
ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS_ARM
</a>
</em>
</td>
<td>
<p>OutboundIPs: Desired outbound IP resources for the cluster load balancer.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterManagedOutboundIPProfile">ManagedClusterManagedOutboundIPProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterNATGatewayProfile">ManagedClusterNATGatewayProfile</a>)
</p>
<div>
<p>Profile of the managed outbound IP resources of the managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>count</code><br/>
<em>
int
</em>
</td>
<td>
<p>Count: The desired number of outbound IPs created/managed by Azure. Allowed values must be in the range of 1 to 16
(inclusive). The default value is 1.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterManagedOutboundIPProfile_ARM">ManagedClusterManagedOutboundIPProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterNATGatewayProfile_ARM">ManagedClusterNATGatewayProfile_ARM</a>)
</p>
<div>
<p>Profile of the managed outbound IP resources of the managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>count</code><br/>
<em>
int
</em>
</td>
<td>
<p>Count: The desired number of outbound IPs created/managed by Azure. Allowed values must be in the range of 1 to 16
(inclusive). The default value is 1.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterManagedOutboundIPProfile_STATUS">ManagedClusterManagedOutboundIPProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterNATGatewayProfile_STATUS">ManagedClusterNATGatewayProfile_STATUS</a>)
</p>
<div>
<p>Profile of the managed outbound IP resources of the managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>count</code><br/>
<em>
int
</em>
</td>
<td>
<p>Count: The desired number of outbound IPs created/managed by Azure. Allowed values must be in the range of 1 to 16
(inclusive). The default value is 1.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterManagedOutboundIPProfile_STATUS_ARM">ManagedClusterManagedOutboundIPProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterNATGatewayProfile_STATUS_ARM">ManagedClusterNATGatewayProfile_STATUS_ARM</a>)
</p>
<div>
<p>Profile of the managed outbound IP resources of the managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>count</code><br/>
<em>
int
</em>
</td>
<td>
<p>Count: The desired number of outbound IPs created/managed by Azure. Allowed values must be in the range of 1 to 16
(inclusive). The default value is 1.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterNATGatewayProfile">ManagedClusterNATGatewayProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile">ContainerServiceNetworkProfile</a>)
</p>
<div>
<p>Profile of the managed cluster NAT gateway.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>effectiveOutboundIPs</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ResourceReference">
[]ResourceReference
</a>
</em>
</td>
<td>
<p>EffectiveOutboundIPs: The effective outbound IP resources of the cluster NAT gateway.</p>
</td>
</tr>
<tr>
<td>
<code>idleTimeoutInMinutes</code><br/>
<em>
int
</em>
</td>
<td>
<p>IdleTimeoutInMinutes: Desired outbound flow idle timeout in minutes. Allowed values are in the range of 4 to 120
(inclusive). The default value is 4 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>managedOutboundIPProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterManagedOutboundIPProfile">
ManagedClusterManagedOutboundIPProfile
</a>
</em>
</td>
<td>
<p>ManagedOutboundIPProfile: Profile of the managed outbound IP resources of the cluster NAT gateway.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterNATGatewayProfile_ARM">ManagedClusterNATGatewayProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_ARM">ContainerServiceNetworkProfile_ARM</a>)
</p>
<div>
<p>Profile of the managed cluster NAT gateway.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>effectiveOutboundIPs</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ResourceReference_ARM">
[]ResourceReference_ARM
</a>
</em>
</td>
<td>
<p>EffectiveOutboundIPs: The effective outbound IP resources of the cluster NAT gateway.</p>
</td>
</tr>
<tr>
<td>
<code>idleTimeoutInMinutes</code><br/>
<em>
int
</em>
</td>
<td>
<p>IdleTimeoutInMinutes: Desired outbound flow idle timeout in minutes. Allowed values are in the range of 4 to 120
(inclusive). The default value is 4 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>managedOutboundIPProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterManagedOutboundIPProfile_ARM">
ManagedClusterManagedOutboundIPProfile_ARM
</a>
</em>
</td>
<td>
<p>ManagedOutboundIPProfile: Profile of the managed outbound IP resources of the cluster NAT gateway.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterNATGatewayProfile_STATUS">ManagedClusterNATGatewayProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_STATUS">ContainerServiceNetworkProfile_STATUS</a>)
</p>
<div>
<p>Profile of the managed cluster NAT gateway.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>effectiveOutboundIPs</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ResourceReference_STATUS">
[]ResourceReference_STATUS
</a>
</em>
</td>
<td>
<p>EffectiveOutboundIPs: The effective outbound IP resources of the cluster NAT gateway.</p>
</td>
</tr>
<tr>
<td>
<code>idleTimeoutInMinutes</code><br/>
<em>
int
</em>
</td>
<td>
<p>IdleTimeoutInMinutes: Desired outbound flow idle timeout in minutes. Allowed values are in the range of 4 to 120
(inclusive). The default value is 4 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>managedOutboundIPProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterManagedOutboundIPProfile_STATUS">
ManagedClusterManagedOutboundIPProfile_STATUS
</a>
</em>
</td>
<td>
<p>ManagedOutboundIPProfile: Profile of the managed outbound IP resources of the cluster NAT gateway.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterNATGatewayProfile_STATUS_ARM">ManagedClusterNATGatewayProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_STATUS_ARM">ContainerServiceNetworkProfile_STATUS_ARM</a>)
</p>
<div>
<p>Profile of the managed cluster NAT gateway.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>effectiveOutboundIPs</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ResourceReference_STATUS_ARM">
[]ResourceReference_STATUS_ARM
</a>
</em>
</td>
<td>
<p>EffectiveOutboundIPs: The effective outbound IP resources of the cluster NAT gateway.</p>
</td>
</tr>
<tr>
<td>
<code>idleTimeoutInMinutes</code><br/>
<em>
int
</em>
</td>
<td>
<p>IdleTimeoutInMinutes: Desired outbound flow idle timeout in minutes. Allowed values are in the range of 4 to 120
(inclusive). The default value is 4 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>managedOutboundIPProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterManagedOutboundIPProfile_STATUS_ARM">
ManagedClusterManagedOutboundIPProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>ManagedOutboundIPProfile: Profile of the managed outbound IP resources of the cluster NAT gateway.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterNodeResourceGroupProfile">ManagedClusterNodeResourceGroupProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
</p>
<div>
<p>Node resource group lockdown profile for a managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>restrictionLevel</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterNodeResourceGroupProfile_RestrictionLevel">
ManagedClusterNodeResourceGroupProfile_RestrictionLevel
</a>
</em>
</td>
<td>
<p>RestrictionLevel: The restriction level applied to the cluster&rsquo;s node resource group</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterNodeResourceGroupProfile_ARM">ManagedClusterNodeResourceGroupProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM</a>)
</p>
<div>
<p>Node resource group lockdown profile for a managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>restrictionLevel</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterNodeResourceGroupProfile_RestrictionLevel">
ManagedClusterNodeResourceGroupProfile_RestrictionLevel
</a>
</em>
</td>
<td>
<p>RestrictionLevel: The restriction level applied to the cluster&rsquo;s node resource group</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterNodeResourceGroupProfile_RestrictionLevel">ManagedClusterNodeResourceGroupProfile_RestrictionLevel
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterNodeResourceGroupProfile">ManagedClusterNodeResourceGroupProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterNodeResourceGroupProfile_ARM">ManagedClusterNodeResourceGroupProfile_ARM</a>)
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
<tbody><tr><td><p>&#34;ReadOnly&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Unrestricted&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterNodeResourceGroupProfile_RestrictionLevel_STATUS">ManagedClusterNodeResourceGroupProfile_RestrictionLevel_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterNodeResourceGroupProfile_STATUS">ManagedClusterNodeResourceGroupProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterNodeResourceGroupProfile_STATUS_ARM">ManagedClusterNodeResourceGroupProfile_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;ReadOnly&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Unrestricted&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterNodeResourceGroupProfile_STATUS">ManagedClusterNodeResourceGroupProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
</p>
<div>
<p>Node resource group lockdown profile for a managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>restrictionLevel</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterNodeResourceGroupProfile_RestrictionLevel_STATUS">
ManagedClusterNodeResourceGroupProfile_RestrictionLevel_STATUS
</a>
</em>
</td>
<td>
<p>RestrictionLevel: The restriction level applied to the cluster&rsquo;s node resource group</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterNodeResourceGroupProfile_STATUS_ARM">ManagedClusterNodeResourceGroupProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>)
</p>
<div>
<p>Node resource group lockdown profile for a managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>restrictionLevel</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterNodeResourceGroupProfile_RestrictionLevel_STATUS">
ManagedClusterNodeResourceGroupProfile_RestrictionLevel_STATUS
</a>
</em>
</td>
<td>
<p>RestrictionLevel: The restriction level applied to the cluster&rsquo;s node resource group</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterOIDCIssuerProfile">ManagedClusterOIDCIssuerProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
</p>
<div>
<p>The OIDC issuer profile of the Managed Cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether the OIDC issuer is enabled.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterOIDCIssuerProfile_ARM">ManagedClusterOIDCIssuerProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM</a>)
</p>
<div>
<p>The OIDC issuer profile of the Managed Cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether the OIDC issuer is enabled.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterOIDCIssuerProfile_STATUS">ManagedClusterOIDCIssuerProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
</p>
<div>
<p>The OIDC issuer profile of the Managed Cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether the OIDC issuer is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>issuerURL</code><br/>
<em>
string
</em>
</td>
<td>
<p>IssuerURL: The OIDC issuer url of the Managed Cluster.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterOIDCIssuerProfile_STATUS_ARM">ManagedClusterOIDCIssuerProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>)
</p>
<div>
<p>The OIDC issuer profile of the Managed Cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether the OIDC issuer is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>issuerURL</code><br/>
<em>
string
</em>
</td>
<td>
<p>IssuerURL: The OIDC issuer url of the Managed Cluster.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterOperatorSecrets">ManagedClusterOperatorSecrets
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterOperatorSpec">ManagedClusterOperatorSpec</a>)
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
<code>adminCredentials</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretDestination">
genruntime.SecretDestination
</a>
</em>
</td>
<td>
<p>AdminCredentials: indicates where the AdminCredentials secret should be placed. If omitted, the secret will not be
retrieved from Azure.</p>
</td>
</tr>
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
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterOperatorSpec">ManagedClusterOperatorSpec
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
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
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterOperatorSecrets">
ManagedClusterOperatorSecrets
</a>
</em>
</td>
<td>
<p>Secrets: configures where to place Azure generated secrets.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentity">ManagedClusterPodIdentity
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProfile">ManagedClusterPodIdentityProfile</a>)
</p>
<div>
<p>Details about the pod identity assigned to the Managed Cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>bindingSelector</code><br/>
<em>
string
</em>
</td>
<td>
<p>BindingSelector: The binding selector to use for the AzureIdentityBinding resource.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.UserAssignedIdentity">
UserAssignedIdentity
</a>
</em>
</td>
<td>
<p>Identity: The user assigned identity details.</p>
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
<p>Name: The name of the pod identity.</p>
</td>
</tr>
<tr>
<td>
<code>namespace</code><br/>
<em>
string
</em>
</td>
<td>
<p>Namespace: The namespace of the pod identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityException">ManagedClusterPodIdentityException
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProfile">ManagedClusterPodIdentityProfile</a>)
</p>
<div>
<p>See <a href="https://azure.github.io/aad-pod-identity/docs/configure/application_exception/">disable AAD Pod Identity for a specific
Pod/Application</a> for more details.</p>
</div>
<table>
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
<p>Name: The name of the pod identity exception.</p>
</td>
</tr>
<tr>
<td>
<code>namespace</code><br/>
<em>
string
</em>
</td>
<td>
<p>Namespace: The namespace of the pod identity exception.</p>
</td>
</tr>
<tr>
<td>
<code>podLabels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>PodLabels: The pod labels to match.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityException_ARM">ManagedClusterPodIdentityException_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProfile_ARM">ManagedClusterPodIdentityProfile_ARM</a>)
</p>
<div>
<p>See <a href="https://azure.github.io/aad-pod-identity/docs/configure/application_exception/">disable AAD Pod Identity for a specific
Pod/Application</a> for more details.</p>
</div>
<table>
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
<p>Name: The name of the pod identity exception.</p>
</td>
</tr>
<tr>
<td>
<code>namespace</code><br/>
<em>
string
</em>
</td>
<td>
<p>Namespace: The namespace of the pod identity exception.</p>
</td>
</tr>
<tr>
<td>
<code>podLabels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>PodLabels: The pod labels to match.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityException_STATUS">ManagedClusterPodIdentityException_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProfile_STATUS">ManagedClusterPodIdentityProfile_STATUS</a>)
</p>
<div>
<p>See <a href="https://azure.github.io/aad-pod-identity/docs/configure/application_exception/">disable AAD Pod Identity for a specific
Pod/Application</a> for more details.</p>
</div>
<table>
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
<p>Name: The name of the pod identity exception.</p>
</td>
</tr>
<tr>
<td>
<code>namespace</code><br/>
<em>
string
</em>
</td>
<td>
<p>Namespace: The namespace of the pod identity exception.</p>
</td>
</tr>
<tr>
<td>
<code>podLabels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>PodLabels: The pod labels to match.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityException_STATUS_ARM">ManagedClusterPodIdentityException_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProfile_STATUS_ARM">ManagedClusterPodIdentityProfile_STATUS_ARM</a>)
</p>
<div>
<p>See <a href="https://azure.github.io/aad-pod-identity/docs/configure/application_exception/">disable AAD Pod Identity for a specific
Pod/Application</a> for more details.</p>
</div>
<table>
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
<p>Name: The name of the pod identity exception.</p>
</td>
</tr>
<tr>
<td>
<code>namespace</code><br/>
<em>
string
</em>
</td>
<td>
<p>Namespace: The namespace of the pod identity exception.</p>
</td>
</tr>
<tr>
<td>
<code>podLabels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>PodLabels: The pod labels to match.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProfile">ManagedClusterPodIdentityProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
</p>
<div>
<p>See <a href="https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity">use AAD pod identity</a> for more details on pod
identity integration.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allowNetworkPluginKubenet</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowNetworkPluginKubenet: Running in Kubenet is disabled by default due to the security related nature of AAD Pod
Identity and the risks of IP spoofing. See <a href="https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity#using-kubenet-network-plugin-with-azure-active-directory-pod-managed-identities">using Kubenet network plugin with AAD Pod
Identity</a>
for more information.</p>
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
<p>Enabled: Whether the pod identity addon is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentity">
[]ManagedClusterPodIdentity
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: The pod identities to use in the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentityExceptions</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityException">
[]ManagedClusterPodIdentityException
</a>
</em>
</td>
<td>
<p>UserAssignedIdentityExceptions: The pod identity exceptions to allow.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProfile_ARM">ManagedClusterPodIdentityProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM</a>)
</p>
<div>
<p>See <a href="https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity">use AAD pod identity</a> for more details on pod
identity integration.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allowNetworkPluginKubenet</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowNetworkPluginKubenet: Running in Kubenet is disabled by default due to the security related nature of AAD Pod
Identity and the risks of IP spoofing. See <a href="https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity#using-kubenet-network-plugin-with-azure-active-directory-pod-managed-identities">using Kubenet network plugin with AAD Pod
Identity</a>
for more information.</p>
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
<p>Enabled: Whether the pod identity addon is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentity_ARM">
[]ManagedClusterPodIdentity_ARM
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: The pod identities to use in the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentityExceptions</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityException_ARM">
[]ManagedClusterPodIdentityException_ARM
</a>
</em>
</td>
<td>
<p>UserAssignedIdentityExceptions: The pod identity exceptions to allow.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProfile_STATUS">ManagedClusterPodIdentityProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
</p>
<div>
<p>See <a href="https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity">use AAD pod identity</a> for more details on pod
identity integration.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allowNetworkPluginKubenet</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowNetworkPluginKubenet: Running in Kubenet is disabled by default due to the security related nature of AAD Pod
Identity and the risks of IP spoofing. See <a href="https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity#using-kubenet-network-plugin-with-azure-active-directory-pod-managed-identities">using Kubenet network plugin with AAD Pod
Identity</a>
for more information.</p>
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
<p>Enabled: Whether the pod identity addon is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentity_STATUS">
[]ManagedClusterPodIdentity_STATUS
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: The pod identities to use in the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentityExceptions</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityException_STATUS">
[]ManagedClusterPodIdentityException_STATUS
</a>
</em>
</td>
<td>
<p>UserAssignedIdentityExceptions: The pod identity exceptions to allow.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProfile_STATUS_ARM">ManagedClusterPodIdentityProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>)
</p>
<div>
<p>See <a href="https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity">use AAD pod identity</a> for more details on pod
identity integration.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allowNetworkPluginKubenet</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AllowNetworkPluginKubenet: Running in Kubenet is disabled by default due to the security related nature of AAD Pod
Identity and the risks of IP spoofing. See <a href="https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity#using-kubenet-network-plugin-with-azure-active-directory-pod-managed-identities">using Kubenet network plugin with AAD Pod
Identity</a>
for more information.</p>
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
<p>Enabled: Whether the pod identity addon is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentity_STATUS_ARM">
[]ManagedClusterPodIdentity_STATUS_ARM
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: The pod identities to use in the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentityExceptions</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityException_STATUS_ARM">
[]ManagedClusterPodIdentityException_STATUS_ARM
</a>
</em>
</td>
<td>
<p>UserAssignedIdentityExceptions: The pod identity exceptions to allow.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProvisioningErrorBody_STATUS">ManagedClusterPodIdentityProvisioningErrorBody_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProvisioningError_STATUS">ManagedClusterPodIdentityProvisioningError_STATUS</a>)
</p>
<div>
<p>An error response from the pod identity provisioning.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>code</code><br/>
<em>
string
</em>
</td>
<td>
<p>Code: An identifier for the error. Codes are invariant and are intended to be consumed programmatically.</p>
</td>
</tr>
<tr>
<td>
<code>details</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled">
[]ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled
</a>
</em>
</td>
<td>
<p>Details: A list of additional details about the error.</p>
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
<p>Message: A message describing the error, intended to be suitable for display in a user interface.</p>
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
<p>Target: The target of the particular error. For example, the name of the property in error.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProvisioningErrorBody_STATUS_ARM">ManagedClusterPodIdentityProvisioningErrorBody_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProvisioningError_STATUS_ARM">ManagedClusterPodIdentityProvisioningError_STATUS_ARM</a>)
</p>
<div>
<p>An error response from the pod identity provisioning.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>code</code><br/>
<em>
string
</em>
</td>
<td>
<p>Code: An identifier for the error. Codes are invariant and are intended to be consumed programmatically.</p>
</td>
</tr>
<tr>
<td>
<code>details</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled_ARM">
[]ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled_ARM
</a>
</em>
</td>
<td>
<p>Details: A list of additional details about the error.</p>
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
<p>Message: A message describing the error, intended to be suitable for display in a user interface.</p>
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
<p>Target: The target of the particular error. For example, the name of the property in error.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled">ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProvisioningErrorBody_STATUS">ManagedClusterPodIdentityProvisioningErrorBody_STATUS</a>)
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
<code>code</code><br/>
<em>
string
</em>
</td>
<td>
<p>Code: An identifier for the error. Codes are invariant and are intended to be consumed programmatically.</p>
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
<p>Message: A message describing the error, intended to be suitable for display in a user interface.</p>
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
<p>Target: The target of the particular error. For example, the name of the property in error.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled_ARM">ManagedClusterPodIdentityProvisioningErrorBody_STATUS_Unrolled_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProvisioningErrorBody_STATUS_ARM">ManagedClusterPodIdentityProvisioningErrorBody_STATUS_ARM</a>)
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
<code>code</code><br/>
<em>
string
</em>
</td>
<td>
<p>Code: An identifier for the error. Codes are invariant and are intended to be consumed programmatically.</p>
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
<p>Message: A message describing the error, intended to be suitable for display in a user interface.</p>
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
<p>Target: The target of the particular error. For example, the name of the property in error.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProvisioningError_STATUS">ManagedClusterPodIdentityProvisioningError_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentity_ProvisioningInfo_STATUS">ManagedClusterPodIdentity_ProvisioningInfo_STATUS</a>)
</p>
<div>
<p>An error response from the pod identity provisioning.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>error</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProvisioningErrorBody_STATUS">
ManagedClusterPodIdentityProvisioningErrorBody_STATUS
</a>
</em>
</td>
<td>
<p>Error: Details about the error.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProvisioningError_STATUS_ARM">ManagedClusterPodIdentityProvisioningError_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentity_ProvisioningInfo_STATUS_ARM">ManagedClusterPodIdentity_ProvisioningInfo_STATUS_ARM</a>)
</p>
<div>
<p>An error response from the pod identity provisioning.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>error</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProvisioningErrorBody_STATUS_ARM">
ManagedClusterPodIdentityProvisioningErrorBody_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Error: Details about the error.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentity_ARM">ManagedClusterPodIdentity_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProfile_ARM">ManagedClusterPodIdentityProfile_ARM</a>)
</p>
<div>
<p>Details about the pod identity assigned to the Managed Cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>bindingSelector</code><br/>
<em>
string
</em>
</td>
<td>
<p>BindingSelector: The binding selector to use for the AzureIdentityBinding resource.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.UserAssignedIdentity_ARM">
UserAssignedIdentity_ARM
</a>
</em>
</td>
<td>
<p>Identity: The user assigned identity details.</p>
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
<p>Name: The name of the pod identity.</p>
</td>
</tr>
<tr>
<td>
<code>namespace</code><br/>
<em>
string
</em>
</td>
<td>
<p>Namespace: The namespace of the pod identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentity_ProvisioningInfo_STATUS">ManagedClusterPodIdentity_ProvisioningInfo_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentity_STATUS">ManagedClusterPodIdentity_STATUS</a>)
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
<code>error</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProvisioningError_STATUS">
ManagedClusterPodIdentityProvisioningError_STATUS
</a>
</em>
</td>
<td>
<p>Error: Pod identity assignment error (if any).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentity_ProvisioningInfo_STATUS_ARM">ManagedClusterPodIdentity_ProvisioningInfo_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentity_STATUS_ARM">ManagedClusterPodIdentity_STATUS_ARM</a>)
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
<code>error</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProvisioningError_STATUS_ARM">
ManagedClusterPodIdentityProvisioningError_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Error: Pod identity assignment error (if any).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentity_ProvisioningState_STATUS">ManagedClusterPodIdentity_ProvisioningState_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentity_STATUS">ManagedClusterPodIdentity_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentity_STATUS_ARM">ManagedClusterPodIdentity_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;Assigned&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Canceled&#34;</p></td>
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
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentity_STATUS">ManagedClusterPodIdentity_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProfile_STATUS">ManagedClusterPodIdentityProfile_STATUS</a>)
</p>
<div>
<p>Details about the pod identity assigned to the Managed Cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>bindingSelector</code><br/>
<em>
string
</em>
</td>
<td>
<p>BindingSelector: The binding selector to use for the AzureIdentityBinding resource.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.UserAssignedIdentity_STATUS">
UserAssignedIdentity_STATUS
</a>
</em>
</td>
<td>
<p>Identity: The user assigned identity details.</p>
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
<p>Name: The name of the pod identity.</p>
</td>
</tr>
<tr>
<td>
<code>namespace</code><br/>
<em>
string
</em>
</td>
<td>
<p>Namespace: The namespace of the pod identity.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningInfo</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentity_ProvisioningInfo_STATUS">
ManagedClusterPodIdentity_ProvisioningInfo_STATUS
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
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentity_ProvisioningState_STATUS">
ManagedClusterPodIdentity_ProvisioningState_STATUS
</a>
</em>
</td>
<td>
<p>ProvisioningState: The current provisioning state of the pod identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentity_STATUS_ARM">ManagedClusterPodIdentity_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProfile_STATUS_ARM">ManagedClusterPodIdentityProfile_STATUS_ARM</a>)
</p>
<div>
<p>Details about the pod identity assigned to the Managed Cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>bindingSelector</code><br/>
<em>
string
</em>
</td>
<td>
<p>BindingSelector: The binding selector to use for the AzureIdentityBinding resource.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.UserAssignedIdentity_STATUS_ARM">
UserAssignedIdentity_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Identity: The user assigned identity details.</p>
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
<p>Name: The name of the pod identity.</p>
</td>
</tr>
<tr>
<td>
<code>namespace</code><br/>
<em>
string
</em>
</td>
<td>
<p>Namespace: The namespace of the pod identity.</p>
</td>
</tr>
<tr>
<td>
<code>provisioningInfo</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentity_ProvisioningInfo_STATUS_ARM">
ManagedClusterPodIdentity_ProvisioningInfo_STATUS_ARM
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
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentity_ProvisioningState_STATUS">
ManagedClusterPodIdentity_ProvisioningState_STATUS
</a>
</em>
</td>
<td>
<p>ProvisioningState: The current provisioning state of the pod identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec_ARM">ManagedCluster_Spec_ARM</a>)
</p>
<div>
<p>Properties of the managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>aadProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAADProfile_ARM">
ManagedClusterAADProfile_ARM
</a>
</em>
</td>
<td>
<p>AadProfile: The Azure Active Directory configuration.</p>
</td>
</tr>
<tr>
<td>
<code>addonProfiles</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAddonProfile_ARM">
map[string]./api/containerservice/v1api20230202preview.ManagedClusterAddonProfile_ARM
</a>
</em>
</td>
<td>
<p>AddonProfiles: The profile of managed cluster add-on.</p>
</td>
</tr>
<tr>
<td>
<code>agentPoolProfiles</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_ARM">
[]ManagedClusterAgentPoolProfile_ARM
</a>
</em>
</td>
<td>
<p>AgentPoolProfiles: The agent pool properties.</p>
</td>
</tr>
<tr>
<td>
<code>apiServerAccessProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAPIServerAccessProfile_ARM">
ManagedClusterAPIServerAccessProfile_ARM
</a>
</em>
</td>
<td>
<p>ApiServerAccessProfile: The access profile for managed cluster API server.</p>
</td>
</tr>
<tr>
<td>
<code>autoScalerProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_AutoScalerProfile_ARM">
ManagedClusterProperties_AutoScalerProfile_ARM
</a>
</em>
</td>
<td>
<p>AutoScalerProfile: Parameters to be applied to the cluster-autoscaler when enabled</p>
</td>
</tr>
<tr>
<td>
<code>autoUpgradeProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile_ARM">
ManagedClusterAutoUpgradeProfile_ARM
</a>
</em>
</td>
<td>
<p>AutoUpgradeProfile: The auto upgrade configuration.</p>
</td>
</tr>
<tr>
<td>
<code>azureMonitorProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfile_ARM">
ManagedClusterAzureMonitorProfile_ARM
</a>
</em>
</td>
<td>
<p>AzureMonitorProfile: Prometheus addon profile for the container service cluster</p>
</td>
</tr>
<tr>
<td>
<code>creationData</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.CreationData_ARM">
CreationData_ARM
</a>
</em>
</td>
<td>
<p>CreationData: CreationData to be used to specify the source Snapshot ID if the cluster will be created/upgraded using a
snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>disableLocalAccounts</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableLocalAccounts: If set to true, getting static credentials will be disabled for this cluster. This must only be
used on Managed Clusters that are AAD enabled. For more details see <a href="https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview">disable local
accounts</a>.</p>
</td>
</tr>
<tr>
<td>
<code>diskEncryptionSetID</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>dnsPrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>DnsPrefix: This cannot be updated once the Managed Cluster has been created.</p>
</td>
</tr>
<tr>
<td>
<code>enableNamespaceResources</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableNamespaceResources: The default value is false. It can be enabled/disabled on creation and updating of the managed
cluster. See <a href="https://aka.ms/NamespaceARMResource">https://aka.ms/NamespaceARMResource</a> for more details on Namespace as
a ARM Resource.</p>
</td>
</tr>
<tr>
<td>
<code>enablePodSecurityPolicy</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePodSecurityPolicy: (DEPRECATED) Whether to enable Kubernetes pod security policy (preview). PodSecurityPolicy was
deprecated in Kubernetes v1.21, and removed from Kubernetes in v1.25. Learn more at <a href="https://aka.ms/k8s/psp">https://aka.ms/k8s/psp</a> and
<a href="https://aka.ms/aks/psp">https://aka.ms/aks/psp</a>.</p>
</td>
</tr>
<tr>
<td>
<code>enableRBAC</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableRBAC: Whether to enable Kubernetes Role-Based Access Control.</p>
</td>
</tr>
<tr>
<td>
<code>fqdnSubdomain</code><br/>
<em>
string
</em>
</td>
<td>
<p>FqdnSubdomain: This cannot be updated once the Managed Cluster has been created.</p>
</td>
</tr>
<tr>
<td>
<code>guardrailsProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.GuardrailsProfile_ARM">
GuardrailsProfile_ARM
</a>
</em>
</td>
<td>
<p>GuardrailsProfile: The guardrails profile holds all the guardrails information for a given cluster</p>
</td>
</tr>
<tr>
<td>
<code>httpProxyConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterHTTPProxyConfig_ARM">
ManagedClusterHTTPProxyConfig_ARM
</a>
</em>
</td>
<td>
<p>HttpProxyConfig: Configurations for provisioning the cluster with HTTP proxy servers.</p>
</td>
</tr>
<tr>
<td>
<code>identityProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.UserAssignedIdentity_ARM">
map[string]./api/containerservice/v1api20230202preview.UserAssignedIdentity_ARM
</a>
</em>
</td>
<td>
<p>IdentityProfile: Identities associated with the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>ingressProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIngressProfile_ARM">
ManagedClusterIngressProfile_ARM
</a>
</em>
</td>
<td>
<p>IngressProfile: Ingress profile for the managed cluster.</p>
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
<p>KubernetesVersion: When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades
must be performed sequentially by major version number. For example, upgrades between 1.14.x -&gt; 1.15.x or 1.15.x -&gt;
1.16.x are allowed, however 1.14.x -&gt; 1.16.x is not allowed. See <a href="https://docs.microsoft.com/azure/aks/upgrade-cluster">upgrading an AKS
cluster</a> for more details.</p>
</td>
</tr>
<tr>
<td>
<code>linuxProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceLinuxProfile_ARM">
ContainerServiceLinuxProfile_ARM
</a>
</em>
</td>
<td>
<p>LinuxProfile: The profile for Linux VMs in the Managed Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>networkProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_ARM">
ContainerServiceNetworkProfile_ARM
</a>
</em>
</td>
<td>
<p>NetworkProfile: The network configuration profile.</p>
</td>
</tr>
<tr>
<td>
<code>nodeResourceGroup</code><br/>
<em>
string
</em>
</td>
<td>
<p>NodeResourceGroup: The name of the resource group containing agent pool nodes.</p>
</td>
</tr>
<tr>
<td>
<code>nodeResourceGroupProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterNodeResourceGroupProfile_ARM">
ManagedClusterNodeResourceGroupProfile_ARM
</a>
</em>
</td>
<td>
<p>NodeResourceGroupProfile: The node resource group configuration profile.</p>
</td>
</tr>
<tr>
<td>
<code>oidcIssuerProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterOIDCIssuerProfile_ARM">
ManagedClusterOIDCIssuerProfile_ARM
</a>
</em>
</td>
<td>
<p>OidcIssuerProfile: The OIDC issuer profile of the Managed Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>podIdentityProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProfile_ARM">
ManagedClusterPodIdentityProfile_ARM
</a>
</em>
</td>
<td>
<p>PodIdentityProfile: See <a href="https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity">use AAD pod identity</a> for more
details on AAD pod identity integration.</p>
</td>
</tr>
<tr>
<td>
<code>privateLinkResources</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PrivateLinkResource_ARM">
[]PrivateLinkResource_ARM
</a>
</em>
</td>
<td>
<p>PrivateLinkResources: Private link resources associated with the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_PublicNetworkAccess">
ManagedClusterProperties_PublicNetworkAccess
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Allow or deny public network access for AKS</p>
</td>
</tr>
<tr>
<td>
<code>securityProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile_ARM">
ManagedClusterSecurityProfile_ARM
</a>
</em>
</td>
<td>
<p>SecurityProfile: Security profile for the managed cluster.</p>
</td>
</tr>
<tr>
<td>
<code>serviceMeshProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ServiceMeshProfile_ARM">
ServiceMeshProfile_ARM
</a>
</em>
</td>
<td>
<p>ServiceMeshProfile: Service mesh profile for a managed cluster.</p>
</td>
</tr>
<tr>
<td>
<code>servicePrincipalProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterServicePrincipalProfile_ARM">
ManagedClusterServicePrincipalProfile_ARM
</a>
</em>
</td>
<td>
<p>ServicePrincipalProfile: Information about a service principal identity for the cluster to use for manipulating Azure
APIs.</p>
</td>
</tr>
<tr>
<td>
<code>storageProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfile_ARM">
ManagedClusterStorageProfile_ARM
</a>
</em>
</td>
<td>
<p>StorageProfile: Storage profile for the managed cluster.</p>
</td>
</tr>
<tr>
<td>
<code>upgradeSettings</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ClusterUpgradeSettings_ARM">
ClusterUpgradeSettings_ARM
</a>
</em>
</td>
<td>
<p>UpgradeSettings: Settings for upgrading a cluster.</p>
</td>
</tr>
<tr>
<td>
<code>windowsProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWindowsProfile_ARM">
ManagedClusterWindowsProfile_ARM
</a>
</em>
</td>
<td>
<p>WindowsProfile: The profile for Windows VMs in the Managed Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>workloadAutoScalerProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfile_ARM">
ManagedClusterWorkloadAutoScalerProfile_ARM
</a>
</em>
</td>
<td>
<p>WorkloadAutoScalerProfile: Workload Auto-scaler profile for the managed cluster.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_AutoScalerProfile">ManagedClusterProperties_AutoScalerProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
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
<code>balance-similar-node-groups</code><br/>
<em>
string
</em>
</td>
<td>
<p>BalanceSimilarNodeGroups: Valid values are &lsquo;true&rsquo; and &lsquo;false&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>expander</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_AutoScalerProfile_Expander">
ManagedClusterProperties_AutoScalerProfile_Expander
</a>
</em>
</td>
<td>
<p>Expander: If not specified, the default is &lsquo;random&rsquo;. See
<a href="https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders">expanders</a> for more
information.</p>
</td>
</tr>
<tr>
<td>
<code>max-empty-bulk-delete</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaxEmptyBulkDelete: The default is 10.</p>
</td>
</tr>
<tr>
<td>
<code>max-graceful-termination-sec</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaxGracefulTerminationSec: The default is 600.</p>
</td>
</tr>
<tr>
<td>
<code>max-node-provision-time</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaxNodeProvisionTime: The default is &lsquo;15m&rsquo;. Values must be an integer followed by an &rsquo;m&rsquo;. No unit of time other than
minutes (m) is supported.</p>
</td>
</tr>
<tr>
<td>
<code>max-total-unready-percentage</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaxTotalUnreadyPercentage: The default is 45. The maximum is 100 and the minimum is 0.</p>
</td>
</tr>
<tr>
<td>
<code>new-pod-scale-up-delay</code><br/>
<em>
string
</em>
</td>
<td>
<p>NewPodScaleUpDelay: For scenarios like burst/batch scale where you don&rsquo;t want CA to act before the kubernetes scheduler
could schedule all the pods, you can tell CA to ignore unscheduled pods before they&rsquo;re a certain age. The default is
&lsquo;0s&rsquo;. Values must be an integer followed by a unit (&rsquo;s&rsquo; for seconds, &rsquo;m&rsquo; for minutes, &lsquo;h&rsquo; for hours, etc).</p>
</td>
</tr>
<tr>
<td>
<code>ok-total-unready-count</code><br/>
<em>
string
</em>
</td>
<td>
<p>OkTotalUnreadyCount: This must be an integer. The default is 3.</p>
</td>
</tr>
<tr>
<td>
<code>scale-down-delay-after-add</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScaleDownDelayAfterAdd: The default is &lsquo;10m&rsquo;. Values must be an integer followed by an &rsquo;m&rsquo;. No unit of time other than
minutes (m) is supported.</p>
</td>
</tr>
<tr>
<td>
<code>scale-down-delay-after-delete</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScaleDownDelayAfterDelete: The default is the scan-interval. Values must be an integer followed by an &rsquo;m&rsquo;. No unit of
time other than minutes (m) is supported.</p>
</td>
</tr>
<tr>
<td>
<code>scale-down-delay-after-failure</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScaleDownDelayAfterFailure: The default is &lsquo;3m&rsquo;. Values must be an integer followed by an &rsquo;m&rsquo;. No unit of time other
than minutes (m) is supported.</p>
</td>
</tr>
<tr>
<td>
<code>scale-down-unneeded-time</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScaleDownUnneededTime: The default is &lsquo;10m&rsquo;. Values must be an integer followed by an &rsquo;m&rsquo;. No unit of time other than
minutes (m) is supported.</p>
</td>
</tr>
<tr>
<td>
<code>scale-down-unready-time</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScaleDownUnreadyTime: The default is &lsquo;20m&rsquo;. Values must be an integer followed by an &rsquo;m&rsquo;. No unit of time other than
minutes (m) is supported.</p>
</td>
</tr>
<tr>
<td>
<code>scale-down-utilization-threshold</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScaleDownUtilizationThreshold: The default is &lsquo;0.5&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>scan-interval</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScanInterval: The default is &lsquo;10&rsquo;. Values must be an integer number of seconds.</p>
</td>
</tr>
<tr>
<td>
<code>skip-nodes-with-local-storage</code><br/>
<em>
string
</em>
</td>
<td>
<p>SkipNodesWithLocalStorage: The default is true.</p>
</td>
</tr>
<tr>
<td>
<code>skip-nodes-with-system-pods</code><br/>
<em>
string
</em>
</td>
<td>
<p>SkipNodesWithSystemPods: The default is true.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_AutoScalerProfile_ARM">ManagedClusterProperties_AutoScalerProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM</a>)
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
<code>balance-similar-node-groups</code><br/>
<em>
string
</em>
</td>
<td>
<p>BalanceSimilarNodeGroups: Valid values are &lsquo;true&rsquo; and &lsquo;false&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>expander</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_AutoScalerProfile_Expander">
ManagedClusterProperties_AutoScalerProfile_Expander
</a>
</em>
</td>
<td>
<p>Expander: If not specified, the default is &lsquo;random&rsquo;. See
<a href="https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders">expanders</a> for more
information.</p>
</td>
</tr>
<tr>
<td>
<code>max-empty-bulk-delete</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaxEmptyBulkDelete: The default is 10.</p>
</td>
</tr>
<tr>
<td>
<code>max-graceful-termination-sec</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaxGracefulTerminationSec: The default is 600.</p>
</td>
</tr>
<tr>
<td>
<code>max-node-provision-time</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaxNodeProvisionTime: The default is &lsquo;15m&rsquo;. Values must be an integer followed by an &rsquo;m&rsquo;. No unit of time other than
minutes (m) is supported.</p>
</td>
</tr>
<tr>
<td>
<code>max-total-unready-percentage</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaxTotalUnreadyPercentage: The default is 45. The maximum is 100 and the minimum is 0.</p>
</td>
</tr>
<tr>
<td>
<code>new-pod-scale-up-delay</code><br/>
<em>
string
</em>
</td>
<td>
<p>NewPodScaleUpDelay: For scenarios like burst/batch scale where you don&rsquo;t want CA to act before the kubernetes scheduler
could schedule all the pods, you can tell CA to ignore unscheduled pods before they&rsquo;re a certain age. The default is
&lsquo;0s&rsquo;. Values must be an integer followed by a unit (&rsquo;s&rsquo; for seconds, &rsquo;m&rsquo; for minutes, &lsquo;h&rsquo; for hours, etc).</p>
</td>
</tr>
<tr>
<td>
<code>ok-total-unready-count</code><br/>
<em>
string
</em>
</td>
<td>
<p>OkTotalUnreadyCount: This must be an integer. The default is 3.</p>
</td>
</tr>
<tr>
<td>
<code>scale-down-delay-after-add</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScaleDownDelayAfterAdd: The default is &lsquo;10m&rsquo;. Values must be an integer followed by an &rsquo;m&rsquo;. No unit of time other than
minutes (m) is supported.</p>
</td>
</tr>
<tr>
<td>
<code>scale-down-delay-after-delete</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScaleDownDelayAfterDelete: The default is the scan-interval. Values must be an integer followed by an &rsquo;m&rsquo;. No unit of
time other than minutes (m) is supported.</p>
</td>
</tr>
<tr>
<td>
<code>scale-down-delay-after-failure</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScaleDownDelayAfterFailure: The default is &lsquo;3m&rsquo;. Values must be an integer followed by an &rsquo;m&rsquo;. No unit of time other
than minutes (m) is supported.</p>
</td>
</tr>
<tr>
<td>
<code>scale-down-unneeded-time</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScaleDownUnneededTime: The default is &lsquo;10m&rsquo;. Values must be an integer followed by an &rsquo;m&rsquo;. No unit of time other than
minutes (m) is supported.</p>
</td>
</tr>
<tr>
<td>
<code>scale-down-unready-time</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScaleDownUnreadyTime: The default is &lsquo;20m&rsquo;. Values must be an integer followed by an &rsquo;m&rsquo;. No unit of time other than
minutes (m) is supported.</p>
</td>
</tr>
<tr>
<td>
<code>scale-down-utilization-threshold</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScaleDownUtilizationThreshold: The default is &lsquo;0.5&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>scan-interval</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScanInterval: The default is &lsquo;10&rsquo;. Values must be an integer number of seconds.</p>
</td>
</tr>
<tr>
<td>
<code>skip-nodes-with-local-storage</code><br/>
<em>
string
</em>
</td>
<td>
<p>SkipNodesWithLocalStorage: The default is true.</p>
</td>
</tr>
<tr>
<td>
<code>skip-nodes-with-system-pods</code><br/>
<em>
string
</em>
</td>
<td>
<p>SkipNodesWithSystemPods: The default is true.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_AutoScalerProfile_Expander">ManagedClusterProperties_AutoScalerProfile_Expander
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_AutoScalerProfile">ManagedClusterProperties_AutoScalerProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_AutoScalerProfile_ARM">ManagedClusterProperties_AutoScalerProfile_ARM</a>)
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
<tbody><tr><td><p>&#34;least-waste&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;most-pods&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;priority&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;random&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_AutoScalerProfile_Expander_STATUS">ManagedClusterProperties_AutoScalerProfile_Expander_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_AutoScalerProfile_STATUS">ManagedClusterProperties_AutoScalerProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_AutoScalerProfile_STATUS_ARM">ManagedClusterProperties_AutoScalerProfile_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;least-waste&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;most-pods&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;priority&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;random&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_AutoScalerProfile_STATUS">ManagedClusterProperties_AutoScalerProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
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
<code>balance-similar-node-groups</code><br/>
<em>
string
</em>
</td>
<td>
<p>BalanceSimilarNodeGroups: Valid values are &lsquo;true&rsquo; and &lsquo;false&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>expander</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_AutoScalerProfile_Expander_STATUS">
ManagedClusterProperties_AutoScalerProfile_Expander_STATUS
</a>
</em>
</td>
<td>
<p>Expander: If not specified, the default is &lsquo;random&rsquo;. See
<a href="https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders">expanders</a> for more
information.</p>
</td>
</tr>
<tr>
<td>
<code>max-empty-bulk-delete</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaxEmptyBulkDelete: The default is 10.</p>
</td>
</tr>
<tr>
<td>
<code>max-graceful-termination-sec</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaxGracefulTerminationSec: The default is 600.</p>
</td>
</tr>
<tr>
<td>
<code>max-node-provision-time</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaxNodeProvisionTime: The default is &lsquo;15m&rsquo;. Values must be an integer followed by an &rsquo;m&rsquo;. No unit of time other than
minutes (m) is supported.</p>
</td>
</tr>
<tr>
<td>
<code>max-total-unready-percentage</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaxTotalUnreadyPercentage: The default is 45. The maximum is 100 and the minimum is 0.</p>
</td>
</tr>
<tr>
<td>
<code>new-pod-scale-up-delay</code><br/>
<em>
string
</em>
</td>
<td>
<p>NewPodScaleUpDelay: For scenarios like burst/batch scale where you don&rsquo;t want CA to act before the kubernetes scheduler
could schedule all the pods, you can tell CA to ignore unscheduled pods before they&rsquo;re a certain age. The default is
&lsquo;0s&rsquo;. Values must be an integer followed by a unit (&rsquo;s&rsquo; for seconds, &rsquo;m&rsquo; for minutes, &lsquo;h&rsquo; for hours, etc).</p>
</td>
</tr>
<tr>
<td>
<code>ok-total-unready-count</code><br/>
<em>
string
</em>
</td>
<td>
<p>OkTotalUnreadyCount: This must be an integer. The default is 3.</p>
</td>
</tr>
<tr>
<td>
<code>scale-down-delay-after-add</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScaleDownDelayAfterAdd: The default is &lsquo;10m&rsquo;. Values must be an integer followed by an &rsquo;m&rsquo;. No unit of time other than
minutes (m) is supported.</p>
</td>
</tr>
<tr>
<td>
<code>scale-down-delay-after-delete</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScaleDownDelayAfterDelete: The default is the scan-interval. Values must be an integer followed by an &rsquo;m&rsquo;. No unit of
time other than minutes (m) is supported.</p>
</td>
</tr>
<tr>
<td>
<code>scale-down-delay-after-failure</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScaleDownDelayAfterFailure: The default is &lsquo;3m&rsquo;. Values must be an integer followed by an &rsquo;m&rsquo;. No unit of time other
than minutes (m) is supported.</p>
</td>
</tr>
<tr>
<td>
<code>scale-down-unneeded-time</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScaleDownUnneededTime: The default is &lsquo;10m&rsquo;. Values must be an integer followed by an &rsquo;m&rsquo;. No unit of time other than
minutes (m) is supported.</p>
</td>
</tr>
<tr>
<td>
<code>scale-down-unready-time</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScaleDownUnreadyTime: The default is &lsquo;20m&rsquo;. Values must be an integer followed by an &rsquo;m&rsquo;. No unit of time other than
minutes (m) is supported.</p>
</td>
</tr>
<tr>
<td>
<code>scale-down-utilization-threshold</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScaleDownUtilizationThreshold: The default is &lsquo;0.5&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>scan-interval</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScanInterval: The default is &lsquo;10&rsquo;. Values must be an integer number of seconds.</p>
</td>
</tr>
<tr>
<td>
<code>skip-nodes-with-local-storage</code><br/>
<em>
string
</em>
</td>
<td>
<p>SkipNodesWithLocalStorage: The default is true.</p>
</td>
</tr>
<tr>
<td>
<code>skip-nodes-with-system-pods</code><br/>
<em>
string
</em>
</td>
<td>
<p>SkipNodesWithSystemPods: The default is true.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_AutoScalerProfile_STATUS_ARM">ManagedClusterProperties_AutoScalerProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>)
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
<code>balance-similar-node-groups</code><br/>
<em>
string
</em>
</td>
<td>
<p>BalanceSimilarNodeGroups: Valid values are &lsquo;true&rsquo; and &lsquo;false&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>expander</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_AutoScalerProfile_Expander_STATUS">
ManagedClusterProperties_AutoScalerProfile_Expander_STATUS
</a>
</em>
</td>
<td>
<p>Expander: If not specified, the default is &lsquo;random&rsquo;. See
<a href="https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders">expanders</a> for more
information.</p>
</td>
</tr>
<tr>
<td>
<code>max-empty-bulk-delete</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaxEmptyBulkDelete: The default is 10.</p>
</td>
</tr>
<tr>
<td>
<code>max-graceful-termination-sec</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaxGracefulTerminationSec: The default is 600.</p>
</td>
</tr>
<tr>
<td>
<code>max-node-provision-time</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaxNodeProvisionTime: The default is &lsquo;15m&rsquo;. Values must be an integer followed by an &rsquo;m&rsquo;. No unit of time other than
minutes (m) is supported.</p>
</td>
</tr>
<tr>
<td>
<code>max-total-unready-percentage</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaxTotalUnreadyPercentage: The default is 45. The maximum is 100 and the minimum is 0.</p>
</td>
</tr>
<tr>
<td>
<code>new-pod-scale-up-delay</code><br/>
<em>
string
</em>
</td>
<td>
<p>NewPodScaleUpDelay: For scenarios like burst/batch scale where you don&rsquo;t want CA to act before the kubernetes scheduler
could schedule all the pods, you can tell CA to ignore unscheduled pods before they&rsquo;re a certain age. The default is
&lsquo;0s&rsquo;. Values must be an integer followed by a unit (&rsquo;s&rsquo; for seconds, &rsquo;m&rsquo; for minutes, &lsquo;h&rsquo; for hours, etc).</p>
</td>
</tr>
<tr>
<td>
<code>ok-total-unready-count</code><br/>
<em>
string
</em>
</td>
<td>
<p>OkTotalUnreadyCount: This must be an integer. The default is 3.</p>
</td>
</tr>
<tr>
<td>
<code>scale-down-delay-after-add</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScaleDownDelayAfterAdd: The default is &lsquo;10m&rsquo;. Values must be an integer followed by an &rsquo;m&rsquo;. No unit of time other than
minutes (m) is supported.</p>
</td>
</tr>
<tr>
<td>
<code>scale-down-delay-after-delete</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScaleDownDelayAfterDelete: The default is the scan-interval. Values must be an integer followed by an &rsquo;m&rsquo;. No unit of
time other than minutes (m) is supported.</p>
</td>
</tr>
<tr>
<td>
<code>scale-down-delay-after-failure</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScaleDownDelayAfterFailure: The default is &lsquo;3m&rsquo;. Values must be an integer followed by an &rsquo;m&rsquo;. No unit of time other
than minutes (m) is supported.</p>
</td>
</tr>
<tr>
<td>
<code>scale-down-unneeded-time</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScaleDownUnneededTime: The default is &lsquo;10m&rsquo;. Values must be an integer followed by an &rsquo;m&rsquo;. No unit of time other than
minutes (m) is supported.</p>
</td>
</tr>
<tr>
<td>
<code>scale-down-unready-time</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScaleDownUnreadyTime: The default is &lsquo;20m&rsquo;. Values must be an integer followed by an &rsquo;m&rsquo;. No unit of time other than
minutes (m) is supported.</p>
</td>
</tr>
<tr>
<td>
<code>scale-down-utilization-threshold</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScaleDownUtilizationThreshold: The default is &lsquo;0.5&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>scan-interval</code><br/>
<em>
string
</em>
</td>
<td>
<p>ScanInterval: The default is &lsquo;10&rsquo;. Values must be an integer number of seconds.</p>
</td>
</tr>
<tr>
<td>
<code>skip-nodes-with-local-storage</code><br/>
<em>
string
</em>
</td>
<td>
<p>SkipNodesWithLocalStorage: The default is true.</p>
</td>
</tr>
<tr>
<td>
<code>skip-nodes-with-system-pods</code><br/>
<em>
string
</em>
</td>
<td>
<p>SkipNodesWithSystemPods: The default is true.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_PublicNetworkAccess">ManagedClusterProperties_PublicNetworkAccess
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
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
</tr><tr><td><p>&#34;SecuredByPerimeter&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_PublicNetworkAccess_STATUS">ManagedClusterProperties_PublicNetworkAccess_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
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
</tr><tr><td><p>&#34;SecuredByPerimeter&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS_ARM">ManagedCluster_STATUS_ARM</a>)
</p>
<div>
<p>Properties of the managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>aadProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAADProfile_STATUS_ARM">
ManagedClusterAADProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>AadProfile: The Azure Active Directory configuration.</p>
</td>
</tr>
<tr>
<td>
<code>addonProfiles</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAddonProfile_STATUS_ARM">
map[string]./api/containerservice/v1api20230202preview.ManagedClusterAddonProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>AddonProfiles: The profile of managed cluster add-on.</p>
</td>
</tr>
<tr>
<td>
<code>agentPoolProfiles</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS_ARM">
[]ManagedClusterAgentPoolProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>AgentPoolProfiles: The agent pool properties.</p>
</td>
</tr>
<tr>
<td>
<code>apiServerAccessProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAPIServerAccessProfile_STATUS_ARM">
ManagedClusterAPIServerAccessProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>ApiServerAccessProfile: The access profile for managed cluster API server.</p>
</td>
</tr>
<tr>
<td>
<code>autoScalerProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_AutoScalerProfile_STATUS_ARM">
ManagedClusterProperties_AutoScalerProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>AutoScalerProfile: Parameters to be applied to the cluster-autoscaler when enabled</p>
</td>
</tr>
<tr>
<td>
<code>autoUpgradeProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile_STATUS_ARM">
ManagedClusterAutoUpgradeProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>AutoUpgradeProfile: The auto upgrade configuration.</p>
</td>
</tr>
<tr>
<td>
<code>azureMonitorProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfile_STATUS_ARM">
ManagedClusterAzureMonitorProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>AzureMonitorProfile: Prometheus addon profile for the container service cluster</p>
</td>
</tr>
<tr>
<td>
<code>azurePortalFQDN</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzurePortalFQDN: The Azure Portal requires certain Cross-Origin Resource Sharing (CORS) headers to be sent in some
responses, which Kubernetes APIServer doesn&rsquo;t handle by default. This special FQDN supports CORS, allowing the Azure
Portal to function properly.</p>
</td>
</tr>
<tr>
<td>
<code>creationData</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.CreationData_STATUS_ARM">
CreationData_STATUS_ARM
</a>
</em>
</td>
<td>
<p>CreationData: CreationData to be used to specify the source Snapshot ID if the cluster will be created/upgraded using a
snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>currentKubernetesVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>CurrentKubernetesVersion: The version of Kubernetes the Managed Cluster is running.</p>
</td>
</tr>
<tr>
<td>
<code>disableLocalAccounts</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableLocalAccounts: If set to true, getting static credentials will be disabled for this cluster. This must only be
used on Managed Clusters that are AAD enabled. For more details see <a href="https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview">disable local
accounts</a>.</p>
</td>
</tr>
<tr>
<td>
<code>diskEncryptionSetID</code><br/>
<em>
string
</em>
</td>
<td>
<p>DiskEncryptionSetID: This is of the form:
&lsquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Compute/&#x200b;diskEncryptionSets/&#x200b;{encryptionSetName}&rsquo;</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>dnsPrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>DnsPrefix: This cannot be updated once the Managed Cluster has been created.</p>
</td>
</tr>
<tr>
<td>
<code>enableNamespaceResources</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableNamespaceResources: The default value is false. It can be enabled/disabled on creation and updating of the managed
cluster. See <a href="https://aka.ms/NamespaceARMResource">https://aka.ms/NamespaceARMResource</a> for more details on Namespace as
a ARM Resource.</p>
</td>
</tr>
<tr>
<td>
<code>enablePodSecurityPolicy</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePodSecurityPolicy: (DEPRECATED) Whether to enable Kubernetes pod security policy (preview). PodSecurityPolicy was
deprecated in Kubernetes v1.21, and removed from Kubernetes in v1.25. Learn more at <a href="https://aka.ms/k8s/psp">https://aka.ms/k8s/psp</a> and
<a href="https://aka.ms/aks/psp">https://aka.ms/aks/psp</a>.</p>
</td>
</tr>
<tr>
<td>
<code>enableRBAC</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableRBAC: Whether to enable Kubernetes Role-Based Access Control.</p>
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
<p>Fqdn: The FQDN of the master pool.</p>
</td>
</tr>
<tr>
<td>
<code>fqdnSubdomain</code><br/>
<em>
string
</em>
</td>
<td>
<p>FqdnSubdomain: This cannot be updated once the Managed Cluster has been created.</p>
</td>
</tr>
<tr>
<td>
<code>guardrailsProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.GuardrailsProfile_STATUS_ARM">
GuardrailsProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>GuardrailsProfile: The guardrails profile holds all the guardrails information for a given cluster</p>
</td>
</tr>
<tr>
<td>
<code>httpProxyConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterHTTPProxyConfig_STATUS_ARM">
ManagedClusterHTTPProxyConfig_STATUS_ARM
</a>
</em>
</td>
<td>
<p>HttpProxyConfig: Configurations for provisioning the cluster with HTTP proxy servers.</p>
</td>
</tr>
<tr>
<td>
<code>identityProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.UserAssignedIdentity_STATUS_ARM">
map[string]./api/containerservice/v1api20230202preview.UserAssignedIdentity_STATUS_ARM
</a>
</em>
</td>
<td>
<p>IdentityProfile: Identities associated with the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>ingressProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIngressProfile_STATUS_ARM">
ManagedClusterIngressProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>IngressProfile: Ingress profile for the managed cluster.</p>
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
<p>KubernetesVersion: When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades
must be performed sequentially by major version number. For example, upgrades between 1.14.x -&gt; 1.15.x or 1.15.x -&gt;
1.16.x are allowed, however 1.14.x -&gt; 1.16.x is not allowed. See <a href="https://docs.microsoft.com/azure/aks/upgrade-cluster">upgrading an AKS
cluster</a> for more details.</p>
</td>
</tr>
<tr>
<td>
<code>linuxProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceLinuxProfile_STATUS_ARM">
ContainerServiceLinuxProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>LinuxProfile: The profile for Linux VMs in the Managed Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>maxAgentPools</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxAgentPools: The max number of agent pools for the managed cluster.</p>
</td>
</tr>
<tr>
<td>
<code>networkProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_STATUS_ARM">
ContainerServiceNetworkProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>NetworkProfile: The network configuration profile.</p>
</td>
</tr>
<tr>
<td>
<code>nodeResourceGroup</code><br/>
<em>
string
</em>
</td>
<td>
<p>NodeResourceGroup: The name of the resource group containing agent pool nodes.</p>
</td>
</tr>
<tr>
<td>
<code>nodeResourceGroupProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterNodeResourceGroupProfile_STATUS_ARM">
ManagedClusterNodeResourceGroupProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>NodeResourceGroupProfile: The node resource group configuration profile.</p>
</td>
</tr>
<tr>
<td>
<code>oidcIssuerProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterOIDCIssuerProfile_STATUS_ARM">
ManagedClusterOIDCIssuerProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>OidcIssuerProfile: The OIDC issuer profile of the Managed Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>podIdentityProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProfile_STATUS_ARM">
ManagedClusterPodIdentityProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>PodIdentityProfile: See <a href="https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity">use AAD pod identity</a> for more
details on AAD pod identity integration.</p>
</td>
</tr>
<tr>
<td>
<code>powerState</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PowerState_STATUS_ARM">
PowerState_STATUS_ARM
</a>
</em>
</td>
<td>
<p>PowerState: The Power State of the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>privateFQDN</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateFQDN: The FQDN of private cluster.</p>
</td>
</tr>
<tr>
<td>
<code>privateLinkResources</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PrivateLinkResource_STATUS_ARM">
[]PrivateLinkResource_STATUS_ARM
</a>
</em>
</td>
<td>
<p>PrivateLinkResources: Private link resources associated with the cluster.</p>
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
<p>ProvisioningState: The current provisioning state.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_PublicNetworkAccess_STATUS">
ManagedClusterProperties_PublicNetworkAccess_STATUS
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Allow or deny public network access for AKS</p>
</td>
</tr>
<tr>
<td>
<code>securityProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile_STATUS_ARM">
ManagedClusterSecurityProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>SecurityProfile: Security profile for the managed cluster.</p>
</td>
</tr>
<tr>
<td>
<code>serviceMeshProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ServiceMeshProfile_STATUS_ARM">
ServiceMeshProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>ServiceMeshProfile: Service mesh profile for a managed cluster.</p>
</td>
</tr>
<tr>
<td>
<code>servicePrincipalProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterServicePrincipalProfile_STATUS_ARM">
ManagedClusterServicePrincipalProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>ServicePrincipalProfile: Information about a service principal identity for the cluster to use for manipulating Azure
APIs.</p>
</td>
</tr>
<tr>
<td>
<code>storageProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfile_STATUS_ARM">
ManagedClusterStorageProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>StorageProfile: Storage profile for the managed cluster.</p>
</td>
</tr>
<tr>
<td>
<code>upgradeSettings</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ClusterUpgradeSettings_STATUS_ARM">
ClusterUpgradeSettings_STATUS_ARM
</a>
</em>
</td>
<td>
<p>UpgradeSettings: Settings for upgrading a cluster.</p>
</td>
</tr>
<tr>
<td>
<code>windowsProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWindowsProfile_STATUS_ARM">
ManagedClusterWindowsProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>WindowsProfile: The profile for Windows VMs in the Managed Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>workloadAutoScalerProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfile_STATUS_ARM">
ManagedClusterWorkloadAutoScalerProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>WorkloadAutoScalerProfile: Workload Auto-scaler profile for the managed cluster.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSKU">ManagedClusterSKU
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
</p>
<div>
<p>The SKU of a Managed Cluster.</p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSKU_Name">
ManagedClusterSKU_Name
</a>
</em>
</td>
<td>
<p>Name: The name of a managed cluster SKU.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSKU_Tier">
ManagedClusterSKU_Tier
</a>
</em>
</td>
<td>
<p>Tier: If not specified, the default is &lsquo;Free&rsquo;. See <a href="https://learn.microsoft.com/azure/aks/free-standard-pricing-tiers">AKS Pricing
Tier</a> for more details.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSKU_ARM">ManagedClusterSKU_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec_ARM">ManagedCluster_Spec_ARM</a>)
</p>
<div>
<p>The SKU of a Managed Cluster.</p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSKU_Name">
ManagedClusterSKU_Name
</a>
</em>
</td>
<td>
<p>Name: The name of a managed cluster SKU.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSKU_Tier">
ManagedClusterSKU_Tier
</a>
</em>
</td>
<td>
<p>Tier: If not specified, the default is &lsquo;Free&rsquo;. See <a href="https://learn.microsoft.com/azure/aks/free-standard-pricing-tiers">AKS Pricing
Tier</a> for more details.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSKU_Name">ManagedClusterSKU_Name
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSKU">ManagedClusterSKU</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSKU_ARM">ManagedClusterSKU_ARM</a>)
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
<tbody><tr><td><p>&#34;Base&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSKU_Name_STATUS">ManagedClusterSKU_Name_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSKU_STATUS">ManagedClusterSKU_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSKU_STATUS_ARM">ManagedClusterSKU_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;Base&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSKU_STATUS">ManagedClusterSKU_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
</p>
<div>
<p>The SKU of a Managed Cluster.</p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSKU_Name_STATUS">
ManagedClusterSKU_Name_STATUS
</a>
</em>
</td>
<td>
<p>Name: The name of a managed cluster SKU.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSKU_Tier_STATUS">
ManagedClusterSKU_Tier_STATUS
</a>
</em>
</td>
<td>
<p>Tier: If not specified, the default is &lsquo;Free&rsquo;. See <a href="https://learn.microsoft.com/azure/aks/free-standard-pricing-tiers">AKS Pricing
Tier</a> for more details.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSKU_STATUS_ARM">ManagedClusterSKU_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS_ARM">ManagedCluster_STATUS_ARM</a>)
</p>
<div>
<p>The SKU of a Managed Cluster.</p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSKU_Name_STATUS">
ManagedClusterSKU_Name_STATUS
</a>
</em>
</td>
<td>
<p>Name: The name of a managed cluster SKU.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSKU_Tier_STATUS">
ManagedClusterSKU_Tier_STATUS
</a>
</em>
</td>
<td>
<p>Tier: If not specified, the default is &lsquo;Free&rsquo;. See <a href="https://learn.microsoft.com/azure/aks/free-standard-pricing-tiers">AKS Pricing
Tier</a> for more details.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSKU_Tier">ManagedClusterSKU_Tier
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSKU">ManagedClusterSKU</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSKU_ARM">ManagedClusterSKU_ARM</a>)
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
<tbody><tr><td><p>&#34;Free&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSKU_Tier_STATUS">ManagedClusterSKU_Tier_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSKU_STATUS">ManagedClusterSKU_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSKU_STATUS_ARM">ManagedClusterSKU_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;Free&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile">ManagedClusterSecurityProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
</p>
<div>
<p>Security profile for the container service cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>azureKeyVaultKms</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AzureKeyVaultKms">
AzureKeyVaultKms
</a>
</em>
</td>
<td>
<p>AzureKeyVaultKms: Azure Key Vault <a href="https://kubernetes.io/docs/tasks/administer-cluster/kms-provider/">key management
service</a> settings for the security profile.</p>
</td>
</tr>
<tr>
<td>
<code>customCATrustCertificates</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileCustomCATrustCertificates">
ManagedClusterSecurityProfileCustomCATrustCertificates
</a>
</em>
</td>
<td>
<p>CustomCATrustCertificates: A list of up to 10 base64 encoded CAs that will be added to the trust store on nodes with the
Custom CA Trust feature enabled. For more information see <a href="https://learn.microsoft.com/en-us/azure/aks/custom-certificate-authority">Custom CA Trust
Certificates</a></p>
</td>
</tr>
<tr>
<td>
<code>defender</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileDefender">
ManagedClusterSecurityProfileDefender
</a>
</em>
</td>
<td>
<p>Defender: Microsoft Defender settings for the security profile.</p>
</td>
</tr>
<tr>
<td>
<code>imageCleaner</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileImageCleaner">
ManagedClusterSecurityProfileImageCleaner
</a>
</em>
</td>
<td>
<p>ImageCleaner: Image Cleaner settings for the security profile.</p>
</td>
</tr>
<tr>
<td>
<code>nodeRestriction</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileNodeRestriction">
ManagedClusterSecurityProfileNodeRestriction
</a>
</em>
</td>
<td>
<p>NodeRestriction: <a href="https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/#noderestriction">Node
Restriction</a> settings
for the security profile.</p>
</td>
</tr>
<tr>
<td>
<code>workloadIdentity</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileWorkloadIdentity">
ManagedClusterSecurityProfileWorkloadIdentity
</a>
</em>
</td>
<td>
<p>WorkloadIdentity: Workload identity settings for the security profile. Workload identity enables Kubernetes applications
to access Azure cloud resources securely with Azure AD. See <a href="https://aka.ms/aks/wi">https://aka.ms/aks/wi</a> for more details.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileCustomCATrustCertificates">ManagedClusterSecurityProfileCustomCATrustCertificates
(<code>[]string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile">ManagedClusterSecurityProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile_ARM">ManagedClusterSecurityProfile_ARM</a>)
</p>
<div>
</div>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileDefender">ManagedClusterSecurityProfileDefender
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile">ManagedClusterSecurityProfile</a>)
</p>
<div>
<p>Microsoft Defender settings for the security profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>logAnalyticsWorkspaceResourceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>LogAnalyticsWorkspaceResourceReference: Resource ID of the Log Analytics workspace to be associated with Microsoft
Defender. When Microsoft Defender is enabled, this field is required and must be a valid workspace resource ID. When
Microsoft Defender is disabled, leave the field empty.</p>
</td>
</tr>
<tr>
<td>
<code>securityMonitoring</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileDefenderSecurityMonitoring">
ManagedClusterSecurityProfileDefenderSecurityMonitoring
</a>
</em>
</td>
<td>
<p>SecurityMonitoring: Microsoft Defender threat detection for Cloud settings for the security profile.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileDefenderSecurityMonitoring">ManagedClusterSecurityProfileDefenderSecurityMonitoring
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileDefender">ManagedClusterSecurityProfileDefender</a>)
</p>
<div>
<p>Microsoft Defender settings for the security profile threat detection.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable Defender threat detection</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileDefenderSecurityMonitoring_ARM">ManagedClusterSecurityProfileDefenderSecurityMonitoring_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileDefender_ARM">ManagedClusterSecurityProfileDefender_ARM</a>)
</p>
<div>
<p>Microsoft Defender settings for the security profile threat detection.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable Defender threat detection</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileDefenderSecurityMonitoring_STATUS">ManagedClusterSecurityProfileDefenderSecurityMonitoring_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileDefender_STATUS">ManagedClusterSecurityProfileDefender_STATUS</a>)
</p>
<div>
<p>Microsoft Defender settings for the security profile threat detection.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable Defender threat detection</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileDefenderSecurityMonitoring_STATUS_ARM">ManagedClusterSecurityProfileDefenderSecurityMonitoring_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileDefender_STATUS_ARM">ManagedClusterSecurityProfileDefender_STATUS_ARM</a>)
</p>
<div>
<p>Microsoft Defender settings for the security profile threat detection.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable Defender threat detection</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileDefender_ARM">ManagedClusterSecurityProfileDefender_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile_ARM">ManagedClusterSecurityProfile_ARM</a>)
</p>
<div>
<p>Microsoft Defender settings for the security profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>logAnalyticsWorkspaceResourceId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>securityMonitoring</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileDefenderSecurityMonitoring_ARM">
ManagedClusterSecurityProfileDefenderSecurityMonitoring_ARM
</a>
</em>
</td>
<td>
<p>SecurityMonitoring: Microsoft Defender threat detection for Cloud settings for the security profile.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileDefender_STATUS">ManagedClusterSecurityProfileDefender_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile_STATUS">ManagedClusterSecurityProfile_STATUS</a>)
</p>
<div>
<p>Microsoft Defender settings for the security profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>logAnalyticsWorkspaceResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>LogAnalyticsWorkspaceResourceId: Resource ID of the Log Analytics workspace to be associated with Microsoft Defender.
When Microsoft Defender is enabled, this field is required and must be a valid workspace resource ID. When Microsoft
Defender is disabled, leave the field empty.</p>
</td>
</tr>
<tr>
<td>
<code>securityMonitoring</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileDefenderSecurityMonitoring_STATUS">
ManagedClusterSecurityProfileDefenderSecurityMonitoring_STATUS
</a>
</em>
</td>
<td>
<p>SecurityMonitoring: Microsoft Defender threat detection for Cloud settings for the security profile.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileDefender_STATUS_ARM">ManagedClusterSecurityProfileDefender_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile_STATUS_ARM">ManagedClusterSecurityProfile_STATUS_ARM</a>)
</p>
<div>
<p>Microsoft Defender settings for the security profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>logAnalyticsWorkspaceResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>LogAnalyticsWorkspaceResourceId: Resource ID of the Log Analytics workspace to be associated with Microsoft Defender.
When Microsoft Defender is enabled, this field is required and must be a valid workspace resource ID. When Microsoft
Defender is disabled, leave the field empty.</p>
</td>
</tr>
<tr>
<td>
<code>securityMonitoring</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileDefenderSecurityMonitoring_STATUS_ARM">
ManagedClusterSecurityProfileDefenderSecurityMonitoring_STATUS_ARM
</a>
</em>
</td>
<td>
<p>SecurityMonitoring: Microsoft Defender threat detection for Cloud settings for the security profile.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileImageCleaner">ManagedClusterSecurityProfileImageCleaner
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile">ManagedClusterSecurityProfile</a>)
</p>
<div>
<p>Image Cleaner removes unused images from nodes, freeing up disk space and helping to reduce attack surface area. Here
are settings for the security profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable Image Cleaner on AKS cluster.</p>
</td>
</tr>
<tr>
<td>
<code>intervalHours</code><br/>
<em>
int
</em>
</td>
<td>
<p>IntervalHours: Image Cleaner scanning interval in hours.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileImageCleaner_ARM">ManagedClusterSecurityProfileImageCleaner_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile_ARM">ManagedClusterSecurityProfile_ARM</a>)
</p>
<div>
<p>Image Cleaner removes unused images from nodes, freeing up disk space and helping to reduce attack surface area. Here
are settings for the security profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable Image Cleaner on AKS cluster.</p>
</td>
</tr>
<tr>
<td>
<code>intervalHours</code><br/>
<em>
int
</em>
</td>
<td>
<p>IntervalHours: Image Cleaner scanning interval in hours.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileImageCleaner_STATUS">ManagedClusterSecurityProfileImageCleaner_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile_STATUS">ManagedClusterSecurityProfile_STATUS</a>)
</p>
<div>
<p>Image Cleaner removes unused images from nodes, freeing up disk space and helping to reduce attack surface area. Here
are settings for the security profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable Image Cleaner on AKS cluster.</p>
</td>
</tr>
<tr>
<td>
<code>intervalHours</code><br/>
<em>
int
</em>
</td>
<td>
<p>IntervalHours: Image Cleaner scanning interval in hours.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileImageCleaner_STATUS_ARM">ManagedClusterSecurityProfileImageCleaner_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile_STATUS_ARM">ManagedClusterSecurityProfile_STATUS_ARM</a>)
</p>
<div>
<p>Image Cleaner removes unused images from nodes, freeing up disk space and helping to reduce attack surface area. Here
are settings for the security profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable Image Cleaner on AKS cluster.</p>
</td>
</tr>
<tr>
<td>
<code>intervalHours</code><br/>
<em>
int
</em>
</td>
<td>
<p>IntervalHours: Image Cleaner scanning interval in hours.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileNodeRestriction">ManagedClusterSecurityProfileNodeRestriction
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile">ManagedClusterSecurityProfile</a>)
</p>
<div>
<p>Node Restriction settings for the security profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable Node Restriction</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileNodeRestriction_ARM">ManagedClusterSecurityProfileNodeRestriction_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile_ARM">ManagedClusterSecurityProfile_ARM</a>)
</p>
<div>
<p>Node Restriction settings for the security profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable Node Restriction</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileNodeRestriction_STATUS">ManagedClusterSecurityProfileNodeRestriction_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile_STATUS">ManagedClusterSecurityProfile_STATUS</a>)
</p>
<div>
<p>Node Restriction settings for the security profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable Node Restriction</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileNodeRestriction_STATUS_ARM">ManagedClusterSecurityProfileNodeRestriction_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile_STATUS_ARM">ManagedClusterSecurityProfile_STATUS_ARM</a>)
</p>
<div>
<p>Node Restriction settings for the security profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable Node Restriction</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileWorkloadIdentity">ManagedClusterSecurityProfileWorkloadIdentity
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile">ManagedClusterSecurityProfile</a>)
</p>
<div>
<p>Workload identity settings for the security profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable workload identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileWorkloadIdentity_ARM">ManagedClusterSecurityProfileWorkloadIdentity_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile_ARM">ManagedClusterSecurityProfile_ARM</a>)
</p>
<div>
<p>Workload identity settings for the security profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable workload identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileWorkloadIdentity_STATUS">ManagedClusterSecurityProfileWorkloadIdentity_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile_STATUS">ManagedClusterSecurityProfile_STATUS</a>)
</p>
<div>
<p>Workload identity settings for the security profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable workload identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileWorkloadIdentity_STATUS_ARM">ManagedClusterSecurityProfileWorkloadIdentity_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile_STATUS_ARM">ManagedClusterSecurityProfile_STATUS_ARM</a>)
</p>
<div>
<p>Workload identity settings for the security profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable workload identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile_ARM">ManagedClusterSecurityProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM</a>)
</p>
<div>
<p>Security profile for the container service cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>azureKeyVaultKms</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AzureKeyVaultKms_ARM">
AzureKeyVaultKms_ARM
</a>
</em>
</td>
<td>
<p>AzureKeyVaultKms: Azure Key Vault <a href="https://kubernetes.io/docs/tasks/administer-cluster/kms-provider/">key management
service</a> settings for the security profile.</p>
</td>
</tr>
<tr>
<td>
<code>customCATrustCertificates</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileCustomCATrustCertificates">
ManagedClusterSecurityProfileCustomCATrustCertificates
</a>
</em>
</td>
<td>
<p>CustomCATrustCertificates: A list of up to 10 base64 encoded CAs that will be added to the trust store on nodes with the
Custom CA Trust feature enabled. For more information see <a href="https://learn.microsoft.com/en-us/azure/aks/custom-certificate-authority">Custom CA Trust
Certificates</a></p>
</td>
</tr>
<tr>
<td>
<code>defender</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileDefender_ARM">
ManagedClusterSecurityProfileDefender_ARM
</a>
</em>
</td>
<td>
<p>Defender: Microsoft Defender settings for the security profile.</p>
</td>
</tr>
<tr>
<td>
<code>imageCleaner</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileImageCleaner_ARM">
ManagedClusterSecurityProfileImageCleaner_ARM
</a>
</em>
</td>
<td>
<p>ImageCleaner: Image Cleaner settings for the security profile.</p>
</td>
</tr>
<tr>
<td>
<code>nodeRestriction</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileNodeRestriction_ARM">
ManagedClusterSecurityProfileNodeRestriction_ARM
</a>
</em>
</td>
<td>
<p>NodeRestriction: <a href="https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/#noderestriction">Node
Restriction</a> settings
for the security profile.</p>
</td>
</tr>
<tr>
<td>
<code>workloadIdentity</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileWorkloadIdentity_ARM">
ManagedClusterSecurityProfileWorkloadIdentity_ARM
</a>
</em>
</td>
<td>
<p>WorkloadIdentity: Workload identity settings for the security profile. Workload identity enables Kubernetes applications
to access Azure cloud resources securely with Azure AD. See <a href="https://aka.ms/aks/wi">https://aka.ms/aks/wi</a> for more details.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile_STATUS">ManagedClusterSecurityProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
</p>
<div>
<p>Security profile for the container service cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>azureKeyVaultKms</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AzureKeyVaultKms_STATUS">
AzureKeyVaultKms_STATUS
</a>
</em>
</td>
<td>
<p>AzureKeyVaultKms: Azure Key Vault <a href="https://kubernetes.io/docs/tasks/administer-cluster/kms-provider/">key management
service</a> settings for the security profile.</p>
</td>
</tr>
<tr>
<td>
<code>customCATrustCertificates</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>CustomCATrustCertificates: A list of up to 10 base64 encoded CAs that will be added to the trust store on nodes with the
Custom CA Trust feature enabled. For more information see <a href="https://learn.microsoft.com/en-us/azure/aks/custom-certificate-authority">Custom CA Trust
Certificates</a></p>
</td>
</tr>
<tr>
<td>
<code>defender</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileDefender_STATUS">
ManagedClusterSecurityProfileDefender_STATUS
</a>
</em>
</td>
<td>
<p>Defender: Microsoft Defender settings for the security profile.</p>
</td>
</tr>
<tr>
<td>
<code>imageCleaner</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileImageCleaner_STATUS">
ManagedClusterSecurityProfileImageCleaner_STATUS
</a>
</em>
</td>
<td>
<p>ImageCleaner: Image Cleaner settings for the security profile.</p>
</td>
</tr>
<tr>
<td>
<code>nodeRestriction</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileNodeRestriction_STATUS">
ManagedClusterSecurityProfileNodeRestriction_STATUS
</a>
</em>
</td>
<td>
<p>NodeRestriction: <a href="https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/#noderestriction">Node
Restriction</a> settings
for the security profile.</p>
</td>
</tr>
<tr>
<td>
<code>workloadIdentity</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileWorkloadIdentity_STATUS">
ManagedClusterSecurityProfileWorkloadIdentity_STATUS
</a>
</em>
</td>
<td>
<p>WorkloadIdentity: Workload identity settings for the security profile. Workload identity enables Kubernetes applications
to access Azure cloud resources securely with Azure AD. See <a href="https://aka.ms/aks/wi">https://aka.ms/aks/wi</a> for more details.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile_STATUS_ARM">ManagedClusterSecurityProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>)
</p>
<div>
<p>Security profile for the container service cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>azureKeyVaultKms</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AzureKeyVaultKms_STATUS_ARM">
AzureKeyVaultKms_STATUS_ARM
</a>
</em>
</td>
<td>
<p>AzureKeyVaultKms: Azure Key Vault <a href="https://kubernetes.io/docs/tasks/administer-cluster/kms-provider/">key management
service</a> settings for the security profile.</p>
</td>
</tr>
<tr>
<td>
<code>customCATrustCertificates</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>CustomCATrustCertificates: A list of up to 10 base64 encoded CAs that will be added to the trust store on nodes with the
Custom CA Trust feature enabled. For more information see <a href="https://learn.microsoft.com/en-us/azure/aks/custom-certificate-authority">Custom CA Trust
Certificates</a></p>
</td>
</tr>
<tr>
<td>
<code>defender</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileDefender_STATUS_ARM">
ManagedClusterSecurityProfileDefender_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Defender: Microsoft Defender settings for the security profile.</p>
</td>
</tr>
<tr>
<td>
<code>imageCleaner</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileImageCleaner_STATUS_ARM">
ManagedClusterSecurityProfileImageCleaner_STATUS_ARM
</a>
</em>
</td>
<td>
<p>ImageCleaner: Image Cleaner settings for the security profile.</p>
</td>
</tr>
<tr>
<td>
<code>nodeRestriction</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileNodeRestriction_STATUS_ARM">
ManagedClusterSecurityProfileNodeRestriction_STATUS_ARM
</a>
</em>
</td>
<td>
<p>NodeRestriction: <a href="https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/#noderestriction">Node
Restriction</a> settings
for the security profile.</p>
</td>
</tr>
<tr>
<td>
<code>workloadIdentity</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfileWorkloadIdentity_STATUS_ARM">
ManagedClusterSecurityProfileWorkloadIdentity_STATUS_ARM
</a>
</em>
</td>
<td>
<p>WorkloadIdentity: Workload identity settings for the security profile. Workload identity enables Kubernetes applications
to access Azure cloud resources securely with Azure AD. See <a href="https://aka.ms/aks/wi">https://aka.ms/aks/wi</a> for more details.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterServicePrincipalProfile">ManagedClusterServicePrincipalProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
</p>
<div>
<p>Information about a service principal identity for the cluster to use for manipulating Azure APIs.</p>
</div>
<table>
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
<p>ClientId: The ID for the service principal.</p>
</td>
</tr>
<tr>
<td>
<code>secret</code><br/>
<em>
string
</em>
</td>
<td>
<p>Secret: The secret password associated with the service principal in plain text.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterServicePrincipalProfile_ARM">ManagedClusterServicePrincipalProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM</a>)
</p>
<div>
<p>Information about a service principal identity for the cluster to use for manipulating Azure APIs.</p>
</div>
<table>
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
<p>ClientId: The ID for the service principal.</p>
</td>
</tr>
<tr>
<td>
<code>secret</code><br/>
<em>
string
</em>
</td>
<td>
<p>Secret: The secret password associated with the service principal in plain text.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterServicePrincipalProfile_STATUS">ManagedClusterServicePrincipalProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
</p>
<div>
<p>Information about a service principal identity for the cluster to use for manipulating Azure APIs.</p>
</div>
<table>
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
<p>ClientId: The ID for the service principal.</p>
</td>
</tr>
<tr>
<td>
<code>secret</code><br/>
<em>
string
</em>
</td>
<td>
<p>Secret: The secret password associated with the service principal in plain text.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterServicePrincipalProfile_STATUS_ARM">ManagedClusterServicePrincipalProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>)
</p>
<div>
<p>Information about a service principal identity for the cluster to use for manipulating Azure APIs.</p>
</div>
<table>
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
<p>ClientId: The ID for the service principal.</p>
</td>
</tr>
<tr>
<td>
<code>secret</code><br/>
<em>
string
</em>
</td>
<td>
<p>Secret: The secret password associated with the service principal in plain text.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfile">ManagedClusterStorageProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
</p>
<div>
<p>Storage profile for the container service cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>blobCSIDriver</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileBlobCSIDriver">
ManagedClusterStorageProfileBlobCSIDriver
</a>
</em>
</td>
<td>
<p>BlobCSIDriver: AzureBlob CSI Driver settings for the storage profile.</p>
</td>
</tr>
<tr>
<td>
<code>diskCSIDriver</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileDiskCSIDriver">
ManagedClusterStorageProfileDiskCSIDriver
</a>
</em>
</td>
<td>
<p>DiskCSIDriver: AzureDisk CSI Driver settings for the storage profile.</p>
</td>
</tr>
<tr>
<td>
<code>fileCSIDriver</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileFileCSIDriver">
ManagedClusterStorageProfileFileCSIDriver
</a>
</em>
</td>
<td>
<p>FileCSIDriver: AzureFile CSI Driver settings for the storage profile.</p>
</td>
</tr>
<tr>
<td>
<code>snapshotController</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileSnapshotController">
ManagedClusterStorageProfileSnapshotController
</a>
</em>
</td>
<td>
<p>SnapshotController: Snapshot Controller settings for the storage profile.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileBlobCSIDriver">ManagedClusterStorageProfileBlobCSIDriver
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfile">ManagedClusterStorageProfile</a>)
</p>
<div>
<p>AzureBlob CSI Driver settings for the storage profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable AzureBlob CSI Driver. The default value is false.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileBlobCSIDriver_ARM">ManagedClusterStorageProfileBlobCSIDriver_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfile_ARM">ManagedClusterStorageProfile_ARM</a>)
</p>
<div>
<p>AzureBlob CSI Driver settings for the storage profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable AzureBlob CSI Driver. The default value is false.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileBlobCSIDriver_STATUS">ManagedClusterStorageProfileBlobCSIDriver_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfile_STATUS">ManagedClusterStorageProfile_STATUS</a>)
</p>
<div>
<p>AzureBlob CSI Driver settings for the storage profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable AzureBlob CSI Driver. The default value is false.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileBlobCSIDriver_STATUS_ARM">ManagedClusterStorageProfileBlobCSIDriver_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfile_STATUS_ARM">ManagedClusterStorageProfile_STATUS_ARM</a>)
</p>
<div>
<p>AzureBlob CSI Driver settings for the storage profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable AzureBlob CSI Driver. The default value is false.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileDiskCSIDriver">ManagedClusterStorageProfileDiskCSIDriver
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfile">ManagedClusterStorageProfile</a>)
</p>
<div>
<p>AzureDisk CSI Driver settings for the storage profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable AzureDisk CSI Driver. The default value is true.</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br/>
<em>
string
</em>
</td>
<td>
<p>Version: The version of AzureDisk CSI Driver. The default value is v1.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileDiskCSIDriver_ARM">ManagedClusterStorageProfileDiskCSIDriver_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfile_ARM">ManagedClusterStorageProfile_ARM</a>)
</p>
<div>
<p>AzureDisk CSI Driver settings for the storage profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable AzureDisk CSI Driver. The default value is true.</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br/>
<em>
string
</em>
</td>
<td>
<p>Version: The version of AzureDisk CSI Driver. The default value is v1.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileDiskCSIDriver_STATUS">ManagedClusterStorageProfileDiskCSIDriver_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfile_STATUS">ManagedClusterStorageProfile_STATUS</a>)
</p>
<div>
<p>AzureDisk CSI Driver settings for the storage profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable AzureDisk CSI Driver. The default value is true.</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br/>
<em>
string
</em>
</td>
<td>
<p>Version: The version of AzureDisk CSI Driver. The default value is v1.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileDiskCSIDriver_STATUS_ARM">ManagedClusterStorageProfileDiskCSIDriver_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfile_STATUS_ARM">ManagedClusterStorageProfile_STATUS_ARM</a>)
</p>
<div>
<p>AzureDisk CSI Driver settings for the storage profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable AzureDisk CSI Driver. The default value is true.</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br/>
<em>
string
</em>
</td>
<td>
<p>Version: The version of AzureDisk CSI Driver. The default value is v1.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileFileCSIDriver">ManagedClusterStorageProfileFileCSIDriver
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfile">ManagedClusterStorageProfile</a>)
</p>
<div>
<p>AzureFile CSI Driver settings for the storage profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable AzureFile CSI Driver. The default value is true.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileFileCSIDriver_ARM">ManagedClusterStorageProfileFileCSIDriver_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfile_ARM">ManagedClusterStorageProfile_ARM</a>)
</p>
<div>
<p>AzureFile CSI Driver settings for the storage profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable AzureFile CSI Driver. The default value is true.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileFileCSIDriver_STATUS">ManagedClusterStorageProfileFileCSIDriver_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfile_STATUS">ManagedClusterStorageProfile_STATUS</a>)
</p>
<div>
<p>AzureFile CSI Driver settings for the storage profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable AzureFile CSI Driver. The default value is true.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileFileCSIDriver_STATUS_ARM">ManagedClusterStorageProfileFileCSIDriver_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfile_STATUS_ARM">ManagedClusterStorageProfile_STATUS_ARM</a>)
</p>
<div>
<p>AzureFile CSI Driver settings for the storage profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable AzureFile CSI Driver. The default value is true.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileSnapshotController">ManagedClusterStorageProfileSnapshotController
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfile">ManagedClusterStorageProfile</a>)
</p>
<div>
<p>Snapshot Controller settings for the storage profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable Snapshot Controller. The default value is true.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileSnapshotController_ARM">ManagedClusterStorageProfileSnapshotController_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfile_ARM">ManagedClusterStorageProfile_ARM</a>)
</p>
<div>
<p>Snapshot Controller settings for the storage profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable Snapshot Controller. The default value is true.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileSnapshotController_STATUS">ManagedClusterStorageProfileSnapshotController_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfile_STATUS">ManagedClusterStorageProfile_STATUS</a>)
</p>
<div>
<p>Snapshot Controller settings for the storage profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable Snapshot Controller. The default value is true.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileSnapshotController_STATUS_ARM">ManagedClusterStorageProfileSnapshotController_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfile_STATUS_ARM">ManagedClusterStorageProfile_STATUS_ARM</a>)
</p>
<div>
<p>Snapshot Controller settings for the storage profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable Snapshot Controller. The default value is true.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfile_ARM">ManagedClusterStorageProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM</a>)
</p>
<div>
<p>Storage profile for the container service cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>blobCSIDriver</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileBlobCSIDriver_ARM">
ManagedClusterStorageProfileBlobCSIDriver_ARM
</a>
</em>
</td>
<td>
<p>BlobCSIDriver: AzureBlob CSI Driver settings for the storage profile.</p>
</td>
</tr>
<tr>
<td>
<code>diskCSIDriver</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileDiskCSIDriver_ARM">
ManagedClusterStorageProfileDiskCSIDriver_ARM
</a>
</em>
</td>
<td>
<p>DiskCSIDriver: AzureDisk CSI Driver settings for the storage profile.</p>
</td>
</tr>
<tr>
<td>
<code>fileCSIDriver</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileFileCSIDriver_ARM">
ManagedClusterStorageProfileFileCSIDriver_ARM
</a>
</em>
</td>
<td>
<p>FileCSIDriver: AzureFile CSI Driver settings for the storage profile.</p>
</td>
</tr>
<tr>
<td>
<code>snapshotController</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileSnapshotController_ARM">
ManagedClusterStorageProfileSnapshotController_ARM
</a>
</em>
</td>
<td>
<p>SnapshotController: Snapshot Controller settings for the storage profile.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfile_STATUS">ManagedClusterStorageProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
</p>
<div>
<p>Storage profile for the container service cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>blobCSIDriver</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileBlobCSIDriver_STATUS">
ManagedClusterStorageProfileBlobCSIDriver_STATUS
</a>
</em>
</td>
<td>
<p>BlobCSIDriver: AzureBlob CSI Driver settings for the storage profile.</p>
</td>
</tr>
<tr>
<td>
<code>diskCSIDriver</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileDiskCSIDriver_STATUS">
ManagedClusterStorageProfileDiskCSIDriver_STATUS
</a>
</em>
</td>
<td>
<p>DiskCSIDriver: AzureDisk CSI Driver settings for the storage profile.</p>
</td>
</tr>
<tr>
<td>
<code>fileCSIDriver</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileFileCSIDriver_STATUS">
ManagedClusterStorageProfileFileCSIDriver_STATUS
</a>
</em>
</td>
<td>
<p>FileCSIDriver: AzureFile CSI Driver settings for the storage profile.</p>
</td>
</tr>
<tr>
<td>
<code>snapshotController</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileSnapshotController_STATUS">
ManagedClusterStorageProfileSnapshotController_STATUS
</a>
</em>
</td>
<td>
<p>SnapshotController: Snapshot Controller settings for the storage profile.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfile_STATUS_ARM">ManagedClusterStorageProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>)
</p>
<div>
<p>Storage profile for the container service cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>blobCSIDriver</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileBlobCSIDriver_STATUS_ARM">
ManagedClusterStorageProfileBlobCSIDriver_STATUS_ARM
</a>
</em>
</td>
<td>
<p>BlobCSIDriver: AzureBlob CSI Driver settings for the storage profile.</p>
</td>
</tr>
<tr>
<td>
<code>diskCSIDriver</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileDiskCSIDriver_STATUS_ARM">
ManagedClusterStorageProfileDiskCSIDriver_STATUS_ARM
</a>
</em>
</td>
<td>
<p>DiskCSIDriver: AzureDisk CSI Driver settings for the storage profile.</p>
</td>
</tr>
<tr>
<td>
<code>fileCSIDriver</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileFileCSIDriver_STATUS_ARM">
ManagedClusterStorageProfileFileCSIDriver_STATUS_ARM
</a>
</em>
</td>
<td>
<p>FileCSIDriver: AzureFile CSI Driver settings for the storage profile.</p>
</td>
</tr>
<tr>
<td>
<code>snapshotController</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfileSnapshotController_STATUS_ARM">
ManagedClusterStorageProfileSnapshotController_STATUS_ARM
</a>
</em>
</td>
<td>
<p>SnapshotController: Snapshot Controller settings for the storage profile.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterWindowsProfile">ManagedClusterWindowsProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
</p>
<div>
<p>Profile for Windows VMs in the managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>adminPassword</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdminPassword: Specifies the password of the administrator account.
Minimum-length: 8 characters
Max-length: 123 characters
Complexity requirements: 3 out of 4 conditions below need to be fulfilled
Has lower characters
Has upper characters
Has a digit
Has a special character (Regex match [\W_])
Disallowed values: &ldquo;abc@123&rdquo;, &ldquo;P@$$w0rd&rdquo;, &ldquo;P@ssw0rd&rdquo;, &ldquo;P@ssword123&rdquo;, &ldquo;Pa$$word&rdquo;, &ldquo;pass@word1&rdquo;, &ldquo;Password!&rdquo;, &ldquo;Password1&rdquo;,
&ldquo;Password22&rdquo;, &ldquo;iloveyou!&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>adminUsername</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdminUsername: Specifies the name of the administrator account.
Restriction: Cannot end in &ldquo;.&rdquo;
Disallowed values: &ldquo;administrator&rdquo;, &ldquo;admin&rdquo;, &ldquo;user&rdquo;, &ldquo;user1&rdquo;, &ldquo;test&rdquo;, &ldquo;user2&rdquo;, &ldquo;test1&rdquo;, &ldquo;user3&rdquo;, &ldquo;admin1&rdquo;, &ldquo;1&rdquo;, &ldquo;123&rdquo;,
&ldquo;a&rdquo;, &ldquo;actuser&rdquo;, &ldquo;adm&rdquo;, &ldquo;admin2&rdquo;, &ldquo;aspnet&rdquo;, &ldquo;backup&rdquo;, &ldquo;console&rdquo;, &ldquo;david&rdquo;, &ldquo;guest&rdquo;, &ldquo;john&rdquo;, &ldquo;owner&rdquo;, &ldquo;root&rdquo;, &ldquo;server&rdquo;,
&ldquo;sql&rdquo;, &ldquo;support&rdquo;, &ldquo;support_388945a0&rdquo;, &ldquo;sys&rdquo;, &ldquo;test2&rdquo;, &ldquo;test3&rdquo;, &ldquo;user4&rdquo;, &ldquo;user5&rdquo;.
Minimum-length: 1 character
Max-length: 20 characters</p>
</td>
</tr>
<tr>
<td>
<code>enableCSIProxy</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableCSIProxy: For more details on CSI proxy, see the <a href="https://github.com/kubernetes-csi/csi-proxy">CSI proxy GitHub
repo</a>.</p>
</td>
</tr>
<tr>
<td>
<code>gmsaProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.WindowsGmsaProfile">
WindowsGmsaProfile
</a>
</em>
</td>
<td>
<p>GmsaProfile: The Windows gMSA Profile in the Managed Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>licenseType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWindowsProfile_LicenseType">
ManagedClusterWindowsProfile_LicenseType
</a>
</em>
</td>
<td>
<p>LicenseType: The license type to use for Windows VMs. See <a href="https://azure.microsoft.com/pricing/hybrid-benefit/faq/">Azure Hybrid User
Benefits</a> for more details.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterWindowsProfile_ARM">ManagedClusterWindowsProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM</a>)
</p>
<div>
<p>Profile for Windows VMs in the managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>adminPassword</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdminPassword: Specifies the password of the administrator account.
Minimum-length: 8 characters
Max-length: 123 characters
Complexity requirements: 3 out of 4 conditions below need to be fulfilled
Has lower characters
Has upper characters
Has a digit
Has a special character (Regex match [\W_])
Disallowed values: &ldquo;abc@123&rdquo;, &ldquo;P@$$w0rd&rdquo;, &ldquo;P@ssw0rd&rdquo;, &ldquo;P@ssword123&rdquo;, &ldquo;Pa$$word&rdquo;, &ldquo;pass@word1&rdquo;, &ldquo;Password!&rdquo;, &ldquo;Password1&rdquo;,
&ldquo;Password22&rdquo;, &ldquo;iloveyou!&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>adminUsername</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdminUsername: Specifies the name of the administrator account.
Restriction: Cannot end in &ldquo;.&rdquo;
Disallowed values: &ldquo;administrator&rdquo;, &ldquo;admin&rdquo;, &ldquo;user&rdquo;, &ldquo;user1&rdquo;, &ldquo;test&rdquo;, &ldquo;user2&rdquo;, &ldquo;test1&rdquo;, &ldquo;user3&rdquo;, &ldquo;admin1&rdquo;, &ldquo;1&rdquo;, &ldquo;123&rdquo;,
&ldquo;a&rdquo;, &ldquo;actuser&rdquo;, &ldquo;adm&rdquo;, &ldquo;admin2&rdquo;, &ldquo;aspnet&rdquo;, &ldquo;backup&rdquo;, &ldquo;console&rdquo;, &ldquo;david&rdquo;, &ldquo;guest&rdquo;, &ldquo;john&rdquo;, &ldquo;owner&rdquo;, &ldquo;root&rdquo;, &ldquo;server&rdquo;,
&ldquo;sql&rdquo;, &ldquo;support&rdquo;, &ldquo;support_388945a0&rdquo;, &ldquo;sys&rdquo;, &ldquo;test2&rdquo;, &ldquo;test3&rdquo;, &ldquo;user4&rdquo;, &ldquo;user5&rdquo;.
Minimum-length: 1 character
Max-length: 20 characters</p>
</td>
</tr>
<tr>
<td>
<code>enableCSIProxy</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableCSIProxy: For more details on CSI proxy, see the <a href="https://github.com/kubernetes-csi/csi-proxy">CSI proxy GitHub
repo</a>.</p>
</td>
</tr>
<tr>
<td>
<code>gmsaProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.WindowsGmsaProfile_ARM">
WindowsGmsaProfile_ARM
</a>
</em>
</td>
<td>
<p>GmsaProfile: The Windows gMSA Profile in the Managed Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>licenseType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWindowsProfile_LicenseType">
ManagedClusterWindowsProfile_LicenseType
</a>
</em>
</td>
<td>
<p>LicenseType: The license type to use for Windows VMs. See <a href="https://azure.microsoft.com/pricing/hybrid-benefit/faq/">Azure Hybrid User
Benefits</a> for more details.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterWindowsProfile_LicenseType">ManagedClusterWindowsProfile_LicenseType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWindowsProfile">ManagedClusterWindowsProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWindowsProfile_ARM">ManagedClusterWindowsProfile_ARM</a>)
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
</tr><tr><td><p>&#34;Windows_Server&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterWindowsProfile_LicenseType_STATUS">ManagedClusterWindowsProfile_LicenseType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWindowsProfile_STATUS">ManagedClusterWindowsProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWindowsProfile_STATUS_ARM">ManagedClusterWindowsProfile_STATUS_ARM</a>)
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
</tr><tr><td><p>&#34;Windows_Server&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterWindowsProfile_STATUS">ManagedClusterWindowsProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
</p>
<div>
<p>Profile for Windows VMs in the managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>adminPassword</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdminPassword: Specifies the password of the administrator account.
Minimum-length: 8 characters
Max-length: 123 characters
Complexity requirements: 3 out of 4 conditions below need to be fulfilled
Has lower characters
Has upper characters
Has a digit
Has a special character (Regex match [\W_])
Disallowed values: &ldquo;abc@123&rdquo;, &ldquo;P@$$w0rd&rdquo;, &ldquo;P@ssw0rd&rdquo;, &ldquo;P@ssword123&rdquo;, &ldquo;Pa$$word&rdquo;, &ldquo;pass@word1&rdquo;, &ldquo;Password!&rdquo;, &ldquo;Password1&rdquo;,
&ldquo;Password22&rdquo;, &ldquo;iloveyou!&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>adminUsername</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdminUsername: Specifies the name of the administrator account.
Restriction: Cannot end in &ldquo;.&rdquo;
Disallowed values: &ldquo;administrator&rdquo;, &ldquo;admin&rdquo;, &ldquo;user&rdquo;, &ldquo;user1&rdquo;, &ldquo;test&rdquo;, &ldquo;user2&rdquo;, &ldquo;test1&rdquo;, &ldquo;user3&rdquo;, &ldquo;admin1&rdquo;, &ldquo;1&rdquo;, &ldquo;123&rdquo;,
&ldquo;a&rdquo;, &ldquo;actuser&rdquo;, &ldquo;adm&rdquo;, &ldquo;admin2&rdquo;, &ldquo;aspnet&rdquo;, &ldquo;backup&rdquo;, &ldquo;console&rdquo;, &ldquo;david&rdquo;, &ldquo;guest&rdquo;, &ldquo;john&rdquo;, &ldquo;owner&rdquo;, &ldquo;root&rdquo;, &ldquo;server&rdquo;,
&ldquo;sql&rdquo;, &ldquo;support&rdquo;, &ldquo;support_388945a0&rdquo;, &ldquo;sys&rdquo;, &ldquo;test2&rdquo;, &ldquo;test3&rdquo;, &ldquo;user4&rdquo;, &ldquo;user5&rdquo;.
Minimum-length: 1 character
Max-length: 20 characters</p>
</td>
</tr>
<tr>
<td>
<code>enableCSIProxy</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableCSIProxy: For more details on CSI proxy, see the <a href="https://github.com/kubernetes-csi/csi-proxy">CSI proxy GitHub
repo</a>.</p>
</td>
</tr>
<tr>
<td>
<code>gmsaProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.WindowsGmsaProfile_STATUS">
WindowsGmsaProfile_STATUS
</a>
</em>
</td>
<td>
<p>GmsaProfile: The Windows gMSA Profile in the Managed Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>licenseType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWindowsProfile_LicenseType_STATUS">
ManagedClusterWindowsProfile_LicenseType_STATUS
</a>
</em>
</td>
<td>
<p>LicenseType: The license type to use for Windows VMs. See <a href="https://azure.microsoft.com/pricing/hybrid-benefit/faq/">Azure Hybrid User
Benefits</a> for more details.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterWindowsProfile_STATUS_ARM">ManagedClusterWindowsProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>)
</p>
<div>
<p>Profile for Windows VMs in the managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>adminPassword</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdminPassword: Specifies the password of the administrator account.
Minimum-length: 8 characters
Max-length: 123 characters
Complexity requirements: 3 out of 4 conditions below need to be fulfilled
Has lower characters
Has upper characters
Has a digit
Has a special character (Regex match [\W_])
Disallowed values: &ldquo;abc@123&rdquo;, &ldquo;P@$$w0rd&rdquo;, &ldquo;P@ssw0rd&rdquo;, &ldquo;P@ssword123&rdquo;, &ldquo;Pa$$word&rdquo;, &ldquo;pass@word1&rdquo;, &ldquo;Password!&rdquo;, &ldquo;Password1&rdquo;,
&ldquo;Password22&rdquo;, &ldquo;iloveyou!&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>adminUsername</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdminUsername: Specifies the name of the administrator account.
Restriction: Cannot end in &ldquo;.&rdquo;
Disallowed values: &ldquo;administrator&rdquo;, &ldquo;admin&rdquo;, &ldquo;user&rdquo;, &ldquo;user1&rdquo;, &ldquo;test&rdquo;, &ldquo;user2&rdquo;, &ldquo;test1&rdquo;, &ldquo;user3&rdquo;, &ldquo;admin1&rdquo;, &ldquo;1&rdquo;, &ldquo;123&rdquo;,
&ldquo;a&rdquo;, &ldquo;actuser&rdquo;, &ldquo;adm&rdquo;, &ldquo;admin2&rdquo;, &ldquo;aspnet&rdquo;, &ldquo;backup&rdquo;, &ldquo;console&rdquo;, &ldquo;david&rdquo;, &ldquo;guest&rdquo;, &ldquo;john&rdquo;, &ldquo;owner&rdquo;, &ldquo;root&rdquo;, &ldquo;server&rdquo;,
&ldquo;sql&rdquo;, &ldquo;support&rdquo;, &ldquo;support_388945a0&rdquo;, &ldquo;sys&rdquo;, &ldquo;test2&rdquo;, &ldquo;test3&rdquo;, &ldquo;user4&rdquo;, &ldquo;user5&rdquo;.
Minimum-length: 1 character
Max-length: 20 characters</p>
</td>
</tr>
<tr>
<td>
<code>enableCSIProxy</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableCSIProxy: For more details on CSI proxy, see the <a href="https://github.com/kubernetes-csi/csi-proxy">CSI proxy GitHub
repo</a>.</p>
</td>
</tr>
<tr>
<td>
<code>gmsaProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.WindowsGmsaProfile_STATUS_ARM">
WindowsGmsaProfile_STATUS_ARM
</a>
</em>
</td>
<td>
<p>GmsaProfile: The Windows gMSA Profile in the Managed Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>licenseType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWindowsProfile_LicenseType_STATUS">
ManagedClusterWindowsProfile_LicenseType_STATUS
</a>
</em>
</td>
<td>
<p>LicenseType: The license type to use for Windows VMs. See <a href="https://azure.microsoft.com/pricing/hybrid-benefit/faq/">Azure Hybrid User
Benefits</a> for more details.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfile">ManagedClusterWorkloadAutoScalerProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
</p>
<div>
<p>Workload Auto-scaler profile for the managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>keda</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileKeda">
ManagedClusterWorkloadAutoScalerProfileKeda
</a>
</em>
</td>
<td>
<p>Keda: KEDA (Kubernetes Event-driven Autoscaling) settings for the workload auto-scaler profile.</p>
</td>
</tr>
<tr>
<td>
<code>verticalPodAutoscaler</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler">
ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileKeda">ManagedClusterWorkloadAutoScalerProfileKeda
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfile">ManagedClusterWorkloadAutoScalerProfile</a>)
</p>
<div>
<p>KEDA (Kubernetes Event-driven Autoscaling) settings for the workload auto-scaler profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable KEDA.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileKeda_ARM">ManagedClusterWorkloadAutoScalerProfileKeda_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfile_ARM">ManagedClusterWorkloadAutoScalerProfile_ARM</a>)
</p>
<div>
<p>KEDA (Kubernetes Event-driven Autoscaling) settings for the workload auto-scaler profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable KEDA.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileKeda_STATUS">ManagedClusterWorkloadAutoScalerProfileKeda_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfile_STATUS">ManagedClusterWorkloadAutoScalerProfile_STATUS</a>)
</p>
<div>
<p>KEDA (Kubernetes Event-driven Autoscaling) settings for the workload auto-scaler profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable KEDA.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileKeda_STATUS_ARM">ManagedClusterWorkloadAutoScalerProfileKeda_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfile_STATUS_ARM">ManagedClusterWorkloadAutoScalerProfile_STATUS_ARM</a>)
</p>
<div>
<p>KEDA (Kubernetes Event-driven Autoscaling) settings for the workload auto-scaler profile.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: Whether to enable KEDA.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler">ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfile">ManagedClusterWorkloadAutoScalerProfile</a>)
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
<code>controlledValues</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_ControlledValues">
ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_ControlledValues
</a>
</em>
</td>
<td>
<p>ControlledValues: Controls which resource value autoscaler will change. Default value is RequestsAndLimits.</p>
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
<p>Enabled: Whether to enable VPA. Default value is false.</p>
</td>
</tr>
<tr>
<td>
<code>updateMode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_UpdateMode">
ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_UpdateMode
</a>
</em>
</td>
<td>
<p>UpdateMode: Each update mode level is a superset of the lower levels. Off<Initial<Recreate&lt;=Auto. For example: if
UpdateMode is Initial, it means VPA sets the recommended resources in the VerticalPodAutoscaler Custom Resource (from
UpdateMode Off) and also assigns resources on pod creation (from Initial). The default value is Off.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_ARM">ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfile_ARM">ManagedClusterWorkloadAutoScalerProfile_ARM</a>)
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
<code>controlledValues</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_ControlledValues">
ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_ControlledValues
</a>
</em>
</td>
<td>
<p>ControlledValues: Controls which resource value autoscaler will change. Default value is RequestsAndLimits.</p>
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
<p>Enabled: Whether to enable VPA. Default value is false.</p>
</td>
</tr>
<tr>
<td>
<code>updateMode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_UpdateMode">
ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_UpdateMode
</a>
</em>
</td>
<td>
<p>UpdateMode: Each update mode level is a superset of the lower levels. Off<Initial<Recreate&lt;=Auto. For example: if
UpdateMode is Initial, it means VPA sets the recommended resources in the VerticalPodAutoscaler Custom Resource (from
UpdateMode Off) and also assigns resources on pod creation (from Initial). The default value is Off.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_ControlledValues">ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_ControlledValues
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler">ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_ARM">ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_ARM</a>)
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
<tbody><tr><td><p>&#34;RequestsAndLimits&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;RequestsOnly&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_ControlledValues_STATUS">ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_ControlledValues_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS">ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS_ARM">ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;RequestsAndLimits&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;RequestsOnly&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS">ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfile_STATUS">ManagedClusterWorkloadAutoScalerProfile_STATUS</a>)
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
<code>controlledValues</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_ControlledValues_STATUS">
ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_ControlledValues_STATUS
</a>
</em>
</td>
<td>
<p>ControlledValues: Controls which resource value autoscaler will change. Default value is RequestsAndLimits.</p>
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
<p>Enabled: Whether to enable VPA. Default value is false.</p>
</td>
</tr>
<tr>
<td>
<code>updateMode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_UpdateMode_STATUS">
ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_UpdateMode_STATUS
</a>
</em>
</td>
<td>
<p>UpdateMode: Each update mode level is a superset of the lower levels. Off<Initial<Recreate&lt;=Auto. For example: if
UpdateMode is Initial, it means VPA sets the recommended resources in the VerticalPodAutoscaler Custom Resource (from
UpdateMode Off) and also assigns resources on pod creation (from Initial). The default value is Off.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS_ARM">ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfile_STATUS_ARM">ManagedClusterWorkloadAutoScalerProfile_STATUS_ARM</a>)
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
<code>controlledValues</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_ControlledValues_STATUS">
ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_ControlledValues_STATUS
</a>
</em>
</td>
<td>
<p>ControlledValues: Controls which resource value autoscaler will change. Default value is RequestsAndLimits.</p>
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
<p>Enabled: Whether to enable VPA. Default value is false.</p>
</td>
</tr>
<tr>
<td>
<code>updateMode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_UpdateMode_STATUS">
ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_UpdateMode_STATUS
</a>
</em>
</td>
<td>
<p>UpdateMode: Each update mode level is a superset of the lower levels. Off<Initial<Recreate&lt;=Auto. For example: if
UpdateMode is Initial, it means VPA sets the recommended resources in the VerticalPodAutoscaler Custom Resource (from
UpdateMode Off) and also assigns resources on pod creation (from Initial). The default value is Off.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_UpdateMode">ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_UpdateMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler">ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_ARM">ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_ARM</a>)
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
<tbody><tr><td><p>&#34;Auto&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Initial&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Off&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Recreate&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_UpdateMode_STATUS">ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_UpdateMode_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS">ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS_ARM">ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;Auto&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Initial&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Off&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Recreate&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfile_ARM">ManagedClusterWorkloadAutoScalerProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM</a>)
</p>
<div>
<p>Workload Auto-scaler profile for the managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>keda</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileKeda_ARM">
ManagedClusterWorkloadAutoScalerProfileKeda_ARM
</a>
</em>
</td>
<td>
<p>Keda: KEDA (Kubernetes Event-driven Autoscaling) settings for the workload auto-scaler profile.</p>
</td>
</tr>
<tr>
<td>
<code>verticalPodAutoscaler</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_ARM">
ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_ARM
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfile_STATUS">ManagedClusterWorkloadAutoScalerProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
</p>
<div>
<p>Workload Auto-scaler profile for the managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>keda</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileKeda_STATUS">
ManagedClusterWorkloadAutoScalerProfileKeda_STATUS
</a>
</em>
</td>
<td>
<p>Keda: KEDA (Kubernetes Event-driven Autoscaling) settings for the workload auto-scaler profile.</p>
</td>
</tr>
<tr>
<td>
<code>verticalPodAutoscaler</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS">
ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfile_STATUS_ARM">ManagedClusterWorkloadAutoScalerProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>)
</p>
<div>
<p>Workload Auto-scaler profile for the managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>keda</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileKeda_STATUS_ARM">
ManagedClusterWorkloadAutoScalerProfileKeda_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Keda: KEDA (Kubernetes Event-driven Autoscaling) settings for the workload auto-scaler profile.</p>
</td>
</tr>
<tr>
<td>
<code>verticalPodAutoscaler</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS_ARM">
ManagedClusterWorkloadAutoScalerProfileVerticalPodAutoscaler_STATUS_ARM
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster">ManagedCluster</a>)
</p>
<div>
<p>Managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>aadProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAADProfile_STATUS">
ManagedClusterAADProfile_STATUS
</a>
</em>
</td>
<td>
<p>AadProfile: The Azure Active Directory configuration.</p>
</td>
</tr>
<tr>
<td>
<code>addonProfiles</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAddonProfile_STATUS">
map[string]./api/containerservice/v1api20230202preview.ManagedClusterAddonProfile_STATUS
</a>
</em>
</td>
<td>
<p>AddonProfiles: The profile of managed cluster add-on.</p>
</td>
</tr>
<tr>
<td>
<code>agentPoolProfiles</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS">
[]ManagedClusterAgentPoolProfile_STATUS
</a>
</em>
</td>
<td>
<p>AgentPoolProfiles: The agent pool properties.</p>
</td>
</tr>
<tr>
<td>
<code>apiServerAccessProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAPIServerAccessProfile_STATUS">
ManagedClusterAPIServerAccessProfile_STATUS
</a>
</em>
</td>
<td>
<p>ApiServerAccessProfile: The access profile for managed cluster API server.</p>
</td>
</tr>
<tr>
<td>
<code>autoScalerProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_AutoScalerProfile_STATUS">
ManagedClusterProperties_AutoScalerProfile_STATUS
</a>
</em>
</td>
<td>
<p>AutoScalerProfile: Parameters to be applied to the cluster-autoscaler when enabled</p>
</td>
</tr>
<tr>
<td>
<code>autoUpgradeProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile_STATUS">
ManagedClusterAutoUpgradeProfile_STATUS
</a>
</em>
</td>
<td>
<p>AutoUpgradeProfile: The auto upgrade configuration.</p>
</td>
</tr>
<tr>
<td>
<code>azureMonitorProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfile_STATUS">
ManagedClusterAzureMonitorProfile_STATUS
</a>
</em>
</td>
<td>
<p>AzureMonitorProfile: Prometheus addon profile for the container service cluster</p>
</td>
</tr>
<tr>
<td>
<code>azurePortalFQDN</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzurePortalFQDN: The Azure Portal requires certain Cross-Origin Resource Sharing (CORS) headers to be sent in some
responses, which Kubernetes APIServer doesn&rsquo;t handle by default. This special FQDN supports CORS, allowing the Azure
Portal to function properly.</p>
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
<code>creationData</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.CreationData_STATUS">
CreationData_STATUS
</a>
</em>
</td>
<td>
<p>CreationData: CreationData to be used to specify the source Snapshot ID if the cluster will be created/upgraded using a
snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>currentKubernetesVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>CurrentKubernetesVersion: The version of Kubernetes the Managed Cluster is running.</p>
</td>
</tr>
<tr>
<td>
<code>disableLocalAccounts</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableLocalAccounts: If set to true, getting static credentials will be disabled for this cluster. This must only be
used on Managed Clusters that are AAD enabled. For more details see <a href="https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview">disable local
accounts</a>.</p>
</td>
</tr>
<tr>
<td>
<code>diskEncryptionSetID</code><br/>
<em>
string
</em>
</td>
<td>
<p>DiskEncryptionSetID: This is of the form:
&lsquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Compute/&#x200b;diskEncryptionSets/&#x200b;{encryptionSetName}&rsquo;</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>dnsPrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>DnsPrefix: This cannot be updated once the Managed Cluster has been created.</p>
</td>
</tr>
<tr>
<td>
<code>enableNamespaceResources</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableNamespaceResources: The default value is false. It can be enabled/disabled on creation and updating of the managed
cluster. See <a href="https://aka.ms/NamespaceARMResource">https://aka.ms/NamespaceARMResource</a> for more details on Namespace as
a ARM Resource.</p>
</td>
</tr>
<tr>
<td>
<code>enablePodSecurityPolicy</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePodSecurityPolicy: (DEPRECATED) Whether to enable Kubernetes pod security policy (preview). PodSecurityPolicy was
deprecated in Kubernetes v1.21, and removed from Kubernetes in v1.25. Learn more at <a href="https://aka.ms/k8s/psp">https://aka.ms/k8s/psp</a> and
<a href="https://aka.ms/aks/psp">https://aka.ms/aks/psp</a>.</p>
</td>
</tr>
<tr>
<td>
<code>enableRBAC</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableRBAC: Whether to enable Kubernetes Role-Based Access Control.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ExtendedLocation_STATUS">
ExtendedLocation_STATUS
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the Virtual Machine.</p>
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
<p>Fqdn: The FQDN of the master pool.</p>
</td>
</tr>
<tr>
<td>
<code>fqdnSubdomain</code><br/>
<em>
string
</em>
</td>
<td>
<p>FqdnSubdomain: This cannot be updated once the Managed Cluster has been created.</p>
</td>
</tr>
<tr>
<td>
<code>guardrailsProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.GuardrailsProfile_STATUS">
GuardrailsProfile_STATUS
</a>
</em>
</td>
<td>
<p>GuardrailsProfile: The guardrails profile holds all the guardrails information for a given cluster</p>
</td>
</tr>
<tr>
<td>
<code>httpProxyConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterHTTPProxyConfig_STATUS">
ManagedClusterHTTPProxyConfig_STATUS
</a>
</em>
</td>
<td>
<p>HttpProxyConfig: Configurations for provisioning the cluster with HTTP proxy servers.</p>
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
<code>identity</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity_STATUS">
ManagedClusterIdentity_STATUS
</a>
</em>
</td>
<td>
<p>Identity: The identity of the managed cluster, if configured.</p>
</td>
</tr>
<tr>
<td>
<code>identityProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.UserAssignedIdentity_STATUS">
map[string]./api/containerservice/v1api20230202preview.UserAssignedIdentity_STATUS
</a>
</em>
</td>
<td>
<p>IdentityProfile: Identities associated with the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>ingressProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIngressProfile_STATUS">
ManagedClusterIngressProfile_STATUS
</a>
</em>
</td>
<td>
<p>IngressProfile: Ingress profile for the managed cluster.</p>
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
<p>KubernetesVersion: When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades
must be performed sequentially by major version number. For example, upgrades between 1.14.x -&gt; 1.15.x or 1.15.x -&gt;
1.16.x are allowed, however 1.14.x -&gt; 1.16.x is not allowed. See <a href="https://docs.microsoft.com/azure/aks/upgrade-cluster">upgrading an AKS
cluster</a> for more details.</p>
</td>
</tr>
<tr>
<td>
<code>linuxProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceLinuxProfile_STATUS">
ContainerServiceLinuxProfile_STATUS
</a>
</em>
</td>
<td>
<p>LinuxProfile: The profile for Linux VMs in the Managed Cluster.</p>
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
<code>maxAgentPools</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxAgentPools: The max number of agent pools for the managed cluster.</p>
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
<code>networkProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_STATUS">
ContainerServiceNetworkProfile_STATUS
</a>
</em>
</td>
<td>
<p>NetworkProfile: The network configuration profile.</p>
</td>
</tr>
<tr>
<td>
<code>nodeResourceGroup</code><br/>
<em>
string
</em>
</td>
<td>
<p>NodeResourceGroup: The name of the resource group containing agent pool nodes.</p>
</td>
</tr>
<tr>
<td>
<code>nodeResourceGroupProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterNodeResourceGroupProfile_STATUS">
ManagedClusterNodeResourceGroupProfile_STATUS
</a>
</em>
</td>
<td>
<p>NodeResourceGroupProfile: The node resource group configuration profile.</p>
</td>
</tr>
<tr>
<td>
<code>oidcIssuerProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterOIDCIssuerProfile_STATUS">
ManagedClusterOIDCIssuerProfile_STATUS
</a>
</em>
</td>
<td>
<p>OidcIssuerProfile: The OIDC issuer profile of the Managed Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>podIdentityProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProfile_STATUS">
ManagedClusterPodIdentityProfile_STATUS
</a>
</em>
</td>
<td>
<p>PodIdentityProfile: See <a href="https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity">use AAD pod identity</a> for more
details on AAD pod identity integration.</p>
</td>
</tr>
<tr>
<td>
<code>powerState</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PowerState_STATUS">
PowerState_STATUS
</a>
</em>
</td>
<td>
<p>PowerState: The Power State of the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>privateFQDN</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateFQDN: The FQDN of private cluster.</p>
</td>
</tr>
<tr>
<td>
<code>privateLinkResources</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PrivateLinkResource_STATUS">
[]PrivateLinkResource_STATUS
</a>
</em>
</td>
<td>
<p>PrivateLinkResources: Private link resources associated with the cluster.</p>
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
<p>ProvisioningState: The current provisioning state.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_PublicNetworkAccess_STATUS">
ManagedClusterProperties_PublicNetworkAccess_STATUS
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Allow or deny public network access for AKS</p>
</td>
</tr>
<tr>
<td>
<code>securityProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile_STATUS">
ManagedClusterSecurityProfile_STATUS
</a>
</em>
</td>
<td>
<p>SecurityProfile: Security profile for the managed cluster.</p>
</td>
</tr>
<tr>
<td>
<code>serviceMeshProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ServiceMeshProfile_STATUS">
ServiceMeshProfile_STATUS
</a>
</em>
</td>
<td>
<p>ServiceMeshProfile: Service mesh profile for a managed cluster.</p>
</td>
</tr>
<tr>
<td>
<code>servicePrincipalProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterServicePrincipalProfile_STATUS">
ManagedClusterServicePrincipalProfile_STATUS
</a>
</em>
</td>
<td>
<p>ServicePrincipalProfile: Information about a service principal identity for the cluster to use for manipulating Azure
APIs.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSKU_STATUS">
ManagedClusterSKU_STATUS
</a>
</em>
</td>
<td>
<p>Sku: The managed cluster SKU.</p>
</td>
</tr>
<tr>
<td>
<code>storageProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfile_STATUS">
ManagedClusterStorageProfile_STATUS
</a>
</em>
</td>
<td>
<p>StorageProfile: Storage profile for the managed cluster.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.SystemData_STATUS">
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
<tr>
<td>
<code>upgradeSettings</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ClusterUpgradeSettings_STATUS">
ClusterUpgradeSettings_STATUS
</a>
</em>
</td>
<td>
<p>UpgradeSettings: Settings for upgrading a cluster.</p>
</td>
</tr>
<tr>
<td>
<code>windowsProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWindowsProfile_STATUS">
ManagedClusterWindowsProfile_STATUS
</a>
</em>
</td>
<td>
<p>WindowsProfile: The profile for Windows VMs in the Managed Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>workloadAutoScalerProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfile_STATUS">
ManagedClusterWorkloadAutoScalerProfile_STATUS
</a>
</em>
</td>
<td>
<p>WorkloadAutoScalerProfile: Workload Auto-scaler profile for the managed cluster.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS_ARM">ManagedCluster_STATUS_ARM
</h3>
<div>
<p>Managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ExtendedLocation_STATUS_ARM">
ExtendedLocation_STATUS_ARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the Virtual Machine.</p>
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
<code>identity</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity_STATUS_ARM">
ManagedClusterIdentity_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Identity: The identity of the managed cluster, if configured.</p>
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
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">
ManagedClusterProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of a managed cluster.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSKU_STATUS_ARM">
ManagedClusterSKU_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Sku: The managed cluster SKU.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.SystemData_STATUS_ARM">
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
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster">ManagedCluster</a>)
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
<code>aadProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAADProfile">
ManagedClusterAADProfile
</a>
</em>
</td>
<td>
<p>AadProfile: The Azure Active Directory configuration.</p>
</td>
</tr>
<tr>
<td>
<code>addonProfiles</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAddonProfile">
map[string]./api/containerservice/v1api20230202preview.ManagedClusterAddonProfile
</a>
</em>
</td>
<td>
<p>AddonProfiles: The profile of managed cluster add-on.</p>
</td>
</tr>
<tr>
<td>
<code>agentPoolProfiles</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile">
[]ManagedClusterAgentPoolProfile
</a>
</em>
</td>
<td>
<p>AgentPoolProfiles: The agent pool properties.</p>
</td>
</tr>
<tr>
<td>
<code>apiServerAccessProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAPIServerAccessProfile">
ManagedClusterAPIServerAccessProfile
</a>
</em>
</td>
<td>
<p>ApiServerAccessProfile: The access profile for managed cluster API server.</p>
</td>
</tr>
<tr>
<td>
<code>autoScalerProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_AutoScalerProfile">
ManagedClusterProperties_AutoScalerProfile
</a>
</em>
</td>
<td>
<p>AutoScalerProfile: Parameters to be applied to the cluster-autoscaler when enabled</p>
</td>
</tr>
<tr>
<td>
<code>autoUpgradeProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAutoUpgradeProfile">
ManagedClusterAutoUpgradeProfile
</a>
</em>
</td>
<td>
<p>AutoUpgradeProfile: The auto upgrade configuration.</p>
</td>
</tr>
<tr>
<td>
<code>azureMonitorProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAzureMonitorProfile">
ManagedClusterAzureMonitorProfile
</a>
</em>
</td>
<td>
<p>AzureMonitorProfile: Prometheus addon profile for the container service cluster</p>
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
<code>creationData</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.CreationData">
CreationData
</a>
</em>
</td>
<td>
<p>CreationData: CreationData to be used to specify the source Snapshot ID if the cluster will be created/upgraded using a
snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>disableLocalAccounts</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableLocalAccounts: If set to true, getting static credentials will be disabled for this cluster. This must only be
used on Managed Clusters that are AAD enabled. For more details see <a href="https://docs.microsoft.com/azure/aks/managed-aad#disable-local-accounts-preview">disable local
accounts</a>.</p>
</td>
</tr>
<tr>
<td>
<code>diskEncryptionSetReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>DiskEncryptionSetReference: This is of the form:
&lsquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Compute/&#x200b;diskEncryptionSets/&#x200b;{encryptionSetName}&rsquo;</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>dnsPrefix</code><br/>
<em>
string
</em>
</td>
<td>
<p>DnsPrefix: This cannot be updated once the Managed Cluster has been created.</p>
</td>
</tr>
<tr>
<td>
<code>enableNamespaceResources</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableNamespaceResources: The default value is false. It can be enabled/disabled on creation and updating of the managed
cluster. See <a href="https://aka.ms/NamespaceARMResource">https://aka.ms/NamespaceARMResource</a> for more details on Namespace as
a ARM Resource.</p>
</td>
</tr>
<tr>
<td>
<code>enablePodSecurityPolicy</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePodSecurityPolicy: (DEPRECATED) Whether to enable Kubernetes pod security policy (preview). PodSecurityPolicy was
deprecated in Kubernetes v1.21, and removed from Kubernetes in v1.25. Learn more at <a href="https://aka.ms/k8s/psp">https://aka.ms/k8s/psp</a> and
<a href="https://aka.ms/aks/psp">https://aka.ms/aks/psp</a>.</p>
</td>
</tr>
<tr>
<td>
<code>enableRBAC</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableRBAC: Whether to enable Kubernetes Role-Based Access Control.</p>
</td>
</tr>
<tr>
<td>
<code>extendedLocation</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ExtendedLocation">
ExtendedLocation
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the Virtual Machine.</p>
</td>
</tr>
<tr>
<td>
<code>fqdnSubdomain</code><br/>
<em>
string
</em>
</td>
<td>
<p>FqdnSubdomain: This cannot be updated once the Managed Cluster has been created.</p>
</td>
</tr>
<tr>
<td>
<code>guardrailsProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.GuardrailsProfile">
GuardrailsProfile
</a>
</em>
</td>
<td>
<p>GuardrailsProfile: The guardrails profile holds all the guardrails information for a given cluster</p>
</td>
</tr>
<tr>
<td>
<code>httpProxyConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterHTTPProxyConfig">
ManagedClusterHTTPProxyConfig
</a>
</em>
</td>
<td>
<p>HttpProxyConfig: Configurations for provisioning the cluster with HTTP proxy servers.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity">
ManagedClusterIdentity
</a>
</em>
</td>
<td>
<p>Identity: The identity of the managed cluster, if configured.</p>
</td>
</tr>
<tr>
<td>
<code>identityProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.UserAssignedIdentity">
map[string]./api/containerservice/v1api20230202preview.UserAssignedIdentity
</a>
</em>
</td>
<td>
<p>IdentityProfile: Identities associated with the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>ingressProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIngressProfile">
ManagedClusterIngressProfile
</a>
</em>
</td>
<td>
<p>IngressProfile: Ingress profile for the managed cluster.</p>
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
<p>KubernetesVersion: When you upgrade a supported AKS cluster, Kubernetes minor versions cannot be skipped. All upgrades
must be performed sequentially by major version number. For example, upgrades between 1.14.x -&gt; 1.15.x or 1.15.x -&gt;
1.16.x are allowed, however 1.14.x -&gt; 1.16.x is not allowed. See <a href="https://docs.microsoft.com/azure/aks/upgrade-cluster">upgrading an AKS
cluster</a> for more details.</p>
</td>
</tr>
<tr>
<td>
<code>linuxProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceLinuxProfile">
ContainerServiceLinuxProfile
</a>
</em>
</td>
<td>
<p>LinuxProfile: The profile for Linux VMs in the Managed Cluster.</p>
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
<code>networkProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile">
ContainerServiceNetworkProfile
</a>
</em>
</td>
<td>
<p>NetworkProfile: The network configuration profile.</p>
</td>
</tr>
<tr>
<td>
<code>nodeResourceGroup</code><br/>
<em>
string
</em>
</td>
<td>
<p>NodeResourceGroup: The name of the resource group containing agent pool nodes.</p>
</td>
</tr>
<tr>
<td>
<code>nodeResourceGroupProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterNodeResourceGroupProfile">
ManagedClusterNodeResourceGroupProfile
</a>
</em>
</td>
<td>
<p>NodeResourceGroupProfile: The node resource group configuration profile.</p>
</td>
</tr>
<tr>
<td>
<code>oidcIssuerProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterOIDCIssuerProfile">
ManagedClusterOIDCIssuerProfile
</a>
</em>
</td>
<td>
<p>OidcIssuerProfile: The OIDC issuer profile of the Managed Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>operatorSpec</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterOperatorSpec">
ManagedClusterOperatorSpec
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
<code>podIdentityProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentityProfile">
ManagedClusterPodIdentityProfile
</a>
</em>
</td>
<td>
<p>PodIdentityProfile: See <a href="https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity">use AAD pod identity</a> for more
details on AAD pod identity integration.</p>
</td>
</tr>
<tr>
<td>
<code>privateLinkResources</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PrivateLinkResource">
[]PrivateLinkResource
</a>
</em>
</td>
<td>
<p>PrivateLinkResources: Private link resources associated with the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_PublicNetworkAccess">
ManagedClusterProperties_PublicNetworkAccess
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Allow or deny public network access for AKS</p>
</td>
</tr>
<tr>
<td>
<code>securityProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSecurityProfile">
ManagedClusterSecurityProfile
</a>
</em>
</td>
<td>
<p>SecurityProfile: Security profile for the managed cluster.</p>
</td>
</tr>
<tr>
<td>
<code>serviceMeshProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ServiceMeshProfile">
ServiceMeshProfile
</a>
</em>
</td>
<td>
<p>ServiceMeshProfile: Service mesh profile for a managed cluster.</p>
</td>
</tr>
<tr>
<td>
<code>servicePrincipalProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterServicePrincipalProfile">
ManagedClusterServicePrincipalProfile
</a>
</em>
</td>
<td>
<p>ServicePrincipalProfile: Information about a service principal identity for the cluster to use for manipulating Azure
APIs.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSKU">
ManagedClusterSKU
</a>
</em>
</td>
<td>
<p>Sku: The managed cluster SKU.</p>
</td>
</tr>
<tr>
<td>
<code>storageProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterStorageProfile">
ManagedClusterStorageProfile
</a>
</em>
</td>
<td>
<p>StorageProfile: Storage profile for the managed cluster.</p>
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
<code>upgradeSettings</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ClusterUpgradeSettings">
ClusterUpgradeSettings
</a>
</em>
</td>
<td>
<p>UpgradeSettings: Settings for upgrading a cluster.</p>
</td>
</tr>
<tr>
<td>
<code>windowsProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWindowsProfile">
ManagedClusterWindowsProfile
</a>
</em>
</td>
<td>
<p>WindowsProfile: The profile for Windows VMs in the Managed Cluster.</p>
</td>
</tr>
<tr>
<td>
<code>workloadAutoScalerProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWorkloadAutoScalerProfile">
ManagedClusterWorkloadAutoScalerProfile
</a>
</em>
</td>
<td>
<p>WorkloadAutoScalerProfile: Workload Auto-scaler profile for the managed cluster.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec_ARM">ManagedCluster_Spec_ARM
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
<code>extendedLocation</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ExtendedLocation_ARM">
ExtendedLocation_ARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The extended location of the Virtual Machine.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity_ARM">
ManagedClusterIdentity_ARM
</a>
</em>
</td>
<td>
<p>Identity: The identity of the managed cluster, if configured.</p>
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
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">
ManagedClusterProperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of a managed cluster.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterSKU_ARM">
ManagedClusterSKU_ARM
</a>
</em>
</td>
<td>
<p>Sku: The managed cluster SKU.</p>
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
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClustersAgentPool">ManagedClustersAgentPool
</h3>
<div>
<p>Generator information:
- Generated from: /containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-02-02-preview/managedClusters.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ContainerService/&#x200b;managedClusters/&#x200b;{resourceName}/&#x200b;agentPools/&#x200b;{agentPoolName}</&#x200b;p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_Spec">
ManagedClusters_AgentPool_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>availabilityZones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AvailabilityZones: The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType
property is &lsquo;VirtualMachineScaleSets&rsquo;.</p>
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
<code>capacityReservationGroupID</code><br/>
<em>
string
</em>
</td>
<td>
<p>CapacityReservationGroupID: AKS will associate the specified agent pool with the Capacity Reservation Group.</p>
</td>
</tr>
<tr>
<td>
<code>count</code><br/>
<em>
int
</em>
</td>
<td>
<p>Count: Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive)
for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.</p>
</td>
</tr>
<tr>
<td>
<code>creationData</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.CreationData">
CreationData
</a>
</em>
</td>
<td>
<p>CreationData: CreationData to be used to specify the source Snapshot ID if the node pool will be created/upgraded using
a snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>enableAutoScaling</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAutoScaling: Whether to enable auto-scaler</p>
</td>
</tr>
<tr>
<td>
<code>enableCustomCATrust</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableCustomCATrust: When set to true, AKS adds a label to the node indicating that the feature is enabled and deploys a
daemonset along with host services to sync custom certificate authorities from user-provided list of base64 encoded
certificates into node trust stores. Defaults to false.</p>
</td>
</tr>
<tr>
<td>
<code>enableEncryptionAtHost</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableEncryptionAtHost: This is only supported on certain VM sizes and in certain Azure regions. For more information,
see: <a href="https://docs.microsoft.com/azure/aks/enable-host-encryption">https://docs.microsoft.com/azure/aks/enable-host-encryption</a></p>
</td>
</tr>
<tr>
<td>
<code>enableFIPS</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableFIPS: See <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview">Add a FIPS-enabled node
pool</a> for more
details.</p>
</td>
</tr>
<tr>
<td>
<code>enableNodePublicIP</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableNodePublicIP: Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses.
A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine
to minimize hops. For more information see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools">assigning a public IP per
node</a>. The
default is false.</p>
</td>
</tr>
<tr>
<td>
<code>enableUltraSSD</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableUltraSSD: Whether to enable UltraSSD</p>
</td>
</tr>
<tr>
<td>
<code>gpuInstanceProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.GPUInstanceProfile">
GPUInstanceProfile
</a>
</em>
</td>
<td>
<p>GpuInstanceProfile: GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.</p>
</td>
</tr>
<tr>
<td>
<code>hostGroupReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>HostGroupReference: This is of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Compute/&#x200b;hostGroups/&#x200b;{hostGroupName}.
For more information see <a href="https://docs.microsoft.com/azure/virtual-machines/dedicated-hosts">Azure dedicated hosts</a>.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.KubeletConfig">
KubeletConfig
</a>
</em>
</td>
<td>
<p>KubeletConfig: The Kubelet configuration on the agent pool nodes.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.KubeletDiskType">
KubeletDiskType
</a>
</em>
</td>
<td>
<p>KubeletDiskType: Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral
storage.</p>
</td>
</tr>
<tr>
<td>
<code>linuxOSConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.LinuxOSConfig">
LinuxOSConfig
</a>
</em>
</td>
<td>
<p>LinuxOSConfig: The OS configuration of Linux agent nodes.</p>
</td>
</tr>
<tr>
<td>
<code>maxCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxCount: The maximum number of nodes for auto-scaling</p>
</td>
</tr>
<tr>
<td>
<code>maxPods</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxPods: The maximum number of pods that can run on a node.</p>
</td>
</tr>
<tr>
<td>
<code>messageOfTheDay</code><br/>
<em>
string
</em>
</td>
<td>
<p>MessageOfTheDay: A base64-encoded string which will be written to /etc/motd after decoding. This allows customization of
the message of the day for Linux nodes. It must not be specified for Windows nodes. It must be a static string (i.e.,
will be printed raw and not be executed as a script).</p>
</td>
</tr>
<tr>
<td>
<code>minCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MinCount: The minimum number of nodes for auto-scaling</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolMode">
AgentPoolMode
</a>
</em>
</td>
<td>
<p>Mode: A cluster must have at least one &lsquo;System&rsquo; Agent Pool at all times. For additional information on agent pool
restrictions  and best practices, see: <a href="https://docs.microsoft.com/azure/aks/use-system-pools">https://docs.microsoft.com/azure/aks/use-system-pools</a></p>
</td>
</tr>
<tr>
<td>
<code>networkProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolNetworkProfile">
AgentPoolNetworkProfile
</a>
</em>
</td>
<td>
<p>NetworkProfile: Network-related settings of an agent pool.</p>
</td>
</tr>
<tr>
<td>
<code>nodeLabels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>NodeLabels: The node labels to be persisted across all nodes in agent pool.</p>
</td>
</tr>
<tr>
<td>
<code>nodePublicIPPrefixReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>NodePublicIPPrefixReference: This is of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Network/&#x200b;publicIPPrefixes/&#x200b;{publicIPPrefixName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>nodeTaints</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>NodeTaints: The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.</p>
</td>
</tr>
<tr>
<td>
<code>orchestratorVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>OrchestratorVersion: Both patch version <major.minor.patch> and <major.minor> are supported. When <major.minor> is
specified, the latest supported patch version is chosen automatically. Updating the agent pool with the same
<major.minor> once it has been created will not trigger an upgrade, even if a newer patch version is available. As a
best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version
must have the same major version as the control plane. The node pool minor version must be within two minor versions of
the control plane version. The node pool version cannot be greater than the control plane version. For more information
see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool">upgrading a node pool</a>.</p>
</td>
</tr>
<tr>
<td>
<code>osDiskSizeGB</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceOSDisk">
ContainerServiceOSDisk
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSDiskType">
OSDiskType
</a>
</em>
</td>
<td>
<p>OsDiskType: The default is &lsquo;Ephemeral&rsquo; if the VM supports it and has a cache disk larger than the requested
OSDiskSizeGB. Otherwise,  defaults to &lsquo;Managed&rsquo;. May not be changed after creation. For more information see <a href="https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os">Ephemeral
OS</a>.</p>
</td>
</tr>
<tr>
<td>
<code>osSKU</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSSKU">
OSSKU
</a>
</em>
</td>
<td>
<p>OsSKU: Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or
Windows2019 if  OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is
deprecated.</p>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSType">
OSType
</a>
</em>
</td>
<td>
<p>OsType: The operating system type. The default is Linux.</p>
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
reference to a containerservice.azure.com/ManagedCluster resource</p>
</td>
</tr>
<tr>
<td>
<code>podSubnetReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>PodSubnetReference: If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details).
This is of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Network/&#x200b;virtualNetworks/&#x200b;{virtualNetworkName}/&#x200b;subnets/&#x200b;{subnetName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>powerState</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PowerState">
PowerState
</a>
</em>
</td>
<td>
<p>PowerState: When an Agent Pool is first created it is initially Running. The Agent Pool can be stopped by setting this
field to Stopped. A stopped Agent Pool stops all of its VMs and does not accrue billing charges. An Agent Pool can only
be stopped if it is Running and provisioning state is Succeeded</p>
</td>
</tr>
<tr>
<td>
<code>proximityPlacementGroupReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>ProximityPlacementGroupReference: The ID for Proximity Placement Group.</p>
</td>
</tr>
<tr>
<td>
<code>scaleDownMode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleDownMode">
ScaleDownMode
</a>
</em>
</td>
<td>
<p>ScaleDownMode: This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.</p>
</td>
</tr>
<tr>
<td>
<code>scaleSetEvictionPolicy</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleSetEvictionPolicy">
ScaleSetEvictionPolicy
</a>
</em>
</td>
<td>
<p>ScaleSetEvictionPolicy: This cannot be specified unless the scaleSetPriority is &lsquo;Spot&rsquo;. If not specified, the default is
&lsquo;Delete&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>scaleSetPriority</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleSetPriority">
ScaleSetPriority
</a>
</em>
</td>
<td>
<p>ScaleSetPriority: The Virtual Machine Scale Set priority. If not specified, the default is &lsquo;Regular&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>spotMaxPrice</code><br/>
<em>
float64
</em>
</td>
<td>
<p>SpotMaxPrice: Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any
on-demand price. For more details on spot pricing, see <a href="https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing">spot VMs
pricing</a></p>
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
<p>Tags: The tags to be persisted on the agent pool virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolType">
AgentPoolType
</a>
</em>
</td>
<td>
<p>Type: The type of Agent Pool.</p>
</td>
</tr>
<tr>
<td>
<code>upgradeSettings</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolUpgradeSettings">
AgentPoolUpgradeSettings
</a>
</em>
</td>
<td>
<p>UpgradeSettings: Settings for upgrading the agentpool</p>
</td>
</tr>
<tr>
<td>
<code>vmSize</code><br/>
<em>
string
</em>
</td>
<td>
<p>VmSize: VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods
might fail to run correctly. For more details on restricted VM sizes, see:
<a href="https://docs.microsoft.com/azure/aks/quotas-skus-regions">https://docs.microsoft.com/azure/aks/quotas-skus-regions</a></p>
</td>
</tr>
<tr>
<td>
<code>vnetSubnetReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>VnetSubnetReference: If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is
specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Network/&#x200b;virtualNetworks/&#x200b;{virtualNetworkName}/&#x200b;subnets/&#x200b;{subnetName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>windowsProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolWindowsProfile">
AgentPoolWindowsProfile
</a>
</em>
</td>
<td>
<p>WindowsProfile: The Windows agent pool&rsquo;s specific profile.</p>
</td>
</tr>
<tr>
<td>
<code>workloadRuntime</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.WorkloadRuntime">
WorkloadRuntime
</a>
</em>
</td>
<td>
<p>WorkloadRuntime: Determines the type of workload a node can run.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_STATUS">
ManagedClusters_AgentPool_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_STATUS">ManagedClusters_AgentPool_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClustersAgentPool">ManagedClustersAgentPool</a>)
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
<code>availabilityZones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AvailabilityZones: The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType
property is &lsquo;VirtualMachineScaleSets&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>capacityReservationGroupID</code><br/>
<em>
string
</em>
</td>
<td>
<p>CapacityReservationGroupID: AKS will associate the specified agent pool with the Capacity Reservation Group.</p>
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
<code>count</code><br/>
<em>
int
</em>
</td>
<td>
<p>Count: Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive)
for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.</p>
</td>
</tr>
<tr>
<td>
<code>creationData</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.CreationData_STATUS">
CreationData_STATUS
</a>
</em>
</td>
<td>
<p>CreationData: CreationData to be used to specify the source Snapshot ID if the node pool will be created/upgraded using
a snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>currentOrchestratorVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>CurrentOrchestratorVersion: If orchestratorVersion was a fully specified version <major.minor.patch>, this field will be
exactly equal to it. If orchestratorVersion was <major.minor>, this field will contain the full <major.minor.patch>
version being used.</p>
</td>
</tr>
<tr>
<td>
<code>enableAutoScaling</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAutoScaling: Whether to enable auto-scaler</p>
</td>
</tr>
<tr>
<td>
<code>enableCustomCATrust</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableCustomCATrust: When set to true, AKS adds a label to the node indicating that the feature is enabled and deploys a
daemonset along with host services to sync custom certificate authorities from user-provided list of base64 encoded
certificates into node trust stores. Defaults to false.</p>
</td>
</tr>
<tr>
<td>
<code>enableEncryptionAtHost</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableEncryptionAtHost: This is only supported on certain VM sizes and in certain Azure regions. For more information,
see: <a href="https://docs.microsoft.com/azure/aks/enable-host-encryption">https://docs.microsoft.com/azure/aks/enable-host-encryption</a></p>
</td>
</tr>
<tr>
<td>
<code>enableFIPS</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableFIPS: See <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview">Add a FIPS-enabled node
pool</a> for more
details.</p>
</td>
</tr>
<tr>
<td>
<code>enableNodePublicIP</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableNodePublicIP: Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses.
A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine
to minimize hops. For more information see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools">assigning a public IP per
node</a>. The
default is false.</p>
</td>
</tr>
<tr>
<td>
<code>enableUltraSSD</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableUltraSSD: Whether to enable UltraSSD</p>
</td>
</tr>
<tr>
<td>
<code>gpuInstanceProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.GPUInstanceProfile_STATUS">
GPUInstanceProfile_STATUS
</a>
</em>
</td>
<td>
<p>GpuInstanceProfile: GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.</p>
</td>
</tr>
<tr>
<td>
<code>hostGroupID</code><br/>
<em>
string
</em>
</td>
<td>
<p>HostGroupID: This is of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Compute/&#x200b;hostGroups/&#x200b;{hostGroupName}.
For more information see <a href="https://docs.microsoft.com/azure/virtual-machines/dedicated-hosts">Azure dedicated hosts</a>.</p>
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
<p>Id: Resource ID.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.KubeletConfig_STATUS">
KubeletConfig_STATUS
</a>
</em>
</td>
<td>
<p>KubeletConfig: The Kubelet configuration on the agent pool nodes.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.KubeletDiskType_STATUS">
KubeletDiskType_STATUS
</a>
</em>
</td>
<td>
<p>KubeletDiskType: Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral
storage.</p>
</td>
</tr>
<tr>
<td>
<code>linuxOSConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.LinuxOSConfig_STATUS">
LinuxOSConfig_STATUS
</a>
</em>
</td>
<td>
<p>LinuxOSConfig: The OS configuration of Linux agent nodes.</p>
</td>
</tr>
<tr>
<td>
<code>maxCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxCount: The maximum number of nodes for auto-scaling</p>
</td>
</tr>
<tr>
<td>
<code>maxPods</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxPods: The maximum number of pods that can run on a node.</p>
</td>
</tr>
<tr>
<td>
<code>messageOfTheDay</code><br/>
<em>
string
</em>
</td>
<td>
<p>MessageOfTheDay: A base64-encoded string which will be written to /etc/motd after decoding. This allows customization of
the message of the day for Linux nodes. It must not be specified for Windows nodes. It must be a static string (i.e.,
will be printed raw and not be executed as a script).</p>
</td>
</tr>
<tr>
<td>
<code>minCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MinCount: The minimum number of nodes for auto-scaling</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolMode_STATUS">
AgentPoolMode_STATUS
</a>
</em>
</td>
<td>
<p>Mode: A cluster must have at least one &lsquo;System&rsquo; Agent Pool at all times. For additional information on agent pool
restrictions  and best practices, see: <a href="https://docs.microsoft.com/azure/aks/use-system-pools">https://docs.microsoft.com/azure/aks/use-system-pools</a></p>
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>networkProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolNetworkProfile_STATUS">
AgentPoolNetworkProfile_STATUS
</a>
</em>
</td>
<td>
<p>NetworkProfile: Network-related settings of an agent pool.</p>
</td>
</tr>
<tr>
<td>
<code>nodeImageVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>NodeImageVersion: The version of node image</p>
</td>
</tr>
<tr>
<td>
<code>nodeLabels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>NodeLabels: The node labels to be persisted across all nodes in agent pool.</p>
</td>
</tr>
<tr>
<td>
<code>nodePublicIPPrefixID</code><br/>
<em>
string
</em>
</td>
<td>
<p>NodePublicIPPrefixID: This is of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Network/&#x200b;publicIPPrefixes/&#x200b;{publicIPPrefixName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>nodeTaints</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>NodeTaints: The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.</p>
</td>
</tr>
<tr>
<td>
<code>orchestratorVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>OrchestratorVersion: Both patch version <major.minor.patch> and <major.minor> are supported. When <major.minor> is
specified, the latest supported patch version is chosen automatically. Updating the agent pool with the same
<major.minor> once it has been created will not trigger an upgrade, even if a newer patch version is available. As a
best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version
must have the same major version as the control plane. The node pool minor version must be within two minor versions of
the control plane version. The node pool version cannot be greater than the control plane version. For more information
see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool">upgrading a node pool</a>.</p>
</td>
</tr>
<tr>
<td>
<code>osDiskSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSDiskType_STATUS">
OSDiskType_STATUS
</a>
</em>
</td>
<td>
<p>OsDiskType: The default is &lsquo;Ephemeral&rsquo; if the VM supports it and has a cache disk larger than the requested
OSDiskSizeGB. Otherwise,  defaults to &lsquo;Managed&rsquo;. May not be changed after creation. For more information see <a href="https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os">Ephemeral
OS</a>.</p>
</td>
</tr>
<tr>
<td>
<code>osSKU</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSSKU_STATUS">
OSSKU_STATUS
</a>
</em>
</td>
<td>
<p>OsSKU: Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or
Windows2019 if  OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is
deprecated.</p>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSType_STATUS">
OSType_STATUS
</a>
</em>
</td>
<td>
<p>OsType: The operating system type. The default is Linux.</p>
</td>
</tr>
<tr>
<td>
<code>podSubnetID</code><br/>
<em>
string
</em>
</td>
<td>
<p>PodSubnetID: If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details). This is
of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Network/&#x200b;virtualNetworks/&#x200b;{virtualNetworkName}/&#x200b;subnets/&#x200b;{subnetName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>powerState</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PowerState_STATUS">
PowerState_STATUS
</a>
</em>
</td>
<td>
<p>PowerState: When an Agent Pool is first created it is initially Running. The Agent Pool can be stopped by setting this
field to Stopped. A stopped Agent Pool stops all of its VMs and does not accrue billing charges. An Agent Pool can only
be stopped if it is Running and provisioning state is Succeeded</p>
</td>
</tr>
<tr>
<td>
<code>properties_type</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolType_STATUS">
AgentPoolType_STATUS
</a>
</em>
</td>
<td>
<p>PropertiesType: The type of Agent Pool.</p>
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
<p>ProvisioningState: The current deployment or provisioning state.</p>
</td>
</tr>
<tr>
<td>
<code>proximityPlacementGroupID</code><br/>
<em>
string
</em>
</td>
<td>
<p>ProximityPlacementGroupID: The ID for Proximity Placement Group.</p>
</td>
</tr>
<tr>
<td>
<code>scaleDownMode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleDownMode_STATUS">
ScaleDownMode_STATUS
</a>
</em>
</td>
<td>
<p>ScaleDownMode: This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.</p>
</td>
</tr>
<tr>
<td>
<code>scaleSetEvictionPolicy</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleSetEvictionPolicy_STATUS">
ScaleSetEvictionPolicy_STATUS
</a>
</em>
</td>
<td>
<p>ScaleSetEvictionPolicy: This cannot be specified unless the scaleSetPriority is &lsquo;Spot&rsquo;. If not specified, the default is
&lsquo;Delete&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>scaleSetPriority</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleSetPriority_STATUS">
ScaleSetPriority_STATUS
</a>
</em>
</td>
<td>
<p>ScaleSetPriority: The Virtual Machine Scale Set priority. If not specified, the default is &lsquo;Regular&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>spotMaxPrice</code><br/>
<em>
float64
</em>
</td>
<td>
<p>SpotMaxPrice: Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any
on-demand price. For more details on spot pricing, see <a href="https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing">spot VMs
pricing</a></p>
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
<p>Tags: The tags to be persisted on the agent pool virtual machine scale set.</p>
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
<p>Type: Resource type</p>
</td>
</tr>
<tr>
<td>
<code>upgradeSettings</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolUpgradeSettings_STATUS">
AgentPoolUpgradeSettings_STATUS
</a>
</em>
</td>
<td>
<p>UpgradeSettings: Settings for upgrading the agentpool</p>
</td>
</tr>
<tr>
<td>
<code>vmSize</code><br/>
<em>
string
</em>
</td>
<td>
<p>VmSize: VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods
might fail to run correctly. For more details on restricted VM sizes, see:
<a href="https://docs.microsoft.com/azure/aks/quotas-skus-regions">https://docs.microsoft.com/azure/aks/quotas-skus-regions</a></p>
</td>
</tr>
<tr>
<td>
<code>vnetSubnetID</code><br/>
<em>
string
</em>
</td>
<td>
<p>VnetSubnetID: If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is specified,
this applies to nodes and pods, otherwise it applies to just nodes. This is of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Network/&#x200b;virtualNetworks/&#x200b;{virtualNetworkName}/&#x200b;subnets/&#x200b;{subnetName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>windowsProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolWindowsProfile_STATUS">
AgentPoolWindowsProfile_STATUS
</a>
</em>
</td>
<td>
<p>WindowsProfile: The Windows agent pool&rsquo;s specific profile.</p>
</td>
</tr>
<tr>
<td>
<code>workloadRuntime</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.WorkloadRuntime_STATUS">
WorkloadRuntime_STATUS
</a>
</em>
</td>
<td>
<p>WorkloadRuntime: Determines the type of workload a node can run.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_STATUS_ARM">ManagedClusters_AgentPool_STATUS_ARM
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
<p>Id: Resource ID.</p>
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_STATUS_ARM">
ManagedClusterAgentPoolProfileProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of an agent pool.</p>
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
<p>Type: Resource type</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_Spec">ManagedClusters_AgentPool_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClustersAgentPool">ManagedClustersAgentPool</a>)
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
<code>availabilityZones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>AvailabilityZones: The list of Availability zones to use for nodes. This can only be specified if the AgentPoolType
property is &lsquo;VirtualMachineScaleSets&rsquo;.</p>
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
<code>capacityReservationGroupID</code><br/>
<em>
string
</em>
</td>
<td>
<p>CapacityReservationGroupID: AKS will associate the specified agent pool with the Capacity Reservation Group.</p>
</td>
</tr>
<tr>
<td>
<code>count</code><br/>
<em>
int
</em>
</td>
<td>
<p>Count: Number of agents (VMs) to host docker containers. Allowed values must be in the range of 0 to 1000 (inclusive)
for user pools and in the range of 1 to 1000 (inclusive) for system pools. The default value is 1.</p>
</td>
</tr>
<tr>
<td>
<code>creationData</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.CreationData">
CreationData
</a>
</em>
</td>
<td>
<p>CreationData: CreationData to be used to specify the source Snapshot ID if the node pool will be created/upgraded using
a snapshot.</p>
</td>
</tr>
<tr>
<td>
<code>enableAutoScaling</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableAutoScaling: Whether to enable auto-scaler</p>
</td>
</tr>
<tr>
<td>
<code>enableCustomCATrust</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableCustomCATrust: When set to true, AKS adds a label to the node indicating that the feature is enabled and deploys a
daemonset along with host services to sync custom certificate authorities from user-provided list of base64 encoded
certificates into node trust stores. Defaults to false.</p>
</td>
</tr>
<tr>
<td>
<code>enableEncryptionAtHost</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableEncryptionAtHost: This is only supported on certain VM sizes and in certain Azure regions. For more information,
see: <a href="https://docs.microsoft.com/azure/aks/enable-host-encryption">https://docs.microsoft.com/azure/aks/enable-host-encryption</a></p>
</td>
</tr>
<tr>
<td>
<code>enableFIPS</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableFIPS: See <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#add-a-fips-enabled-node-pool-preview">Add a FIPS-enabled node
pool</a> for more
details.</p>
</td>
</tr>
<tr>
<td>
<code>enableNodePublicIP</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableNodePublicIP: Some scenarios may require nodes in a node pool to receive their own dedicated public IP addresses.
A common scenario is for gaming workloads, where a console needs to make a direct connection to a cloud virtual machine
to minimize hops. For more information see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#assign-a-public-ip-per-node-for-your-node-pools">assigning a public IP per
node</a>. The
default is false.</p>
</td>
</tr>
<tr>
<td>
<code>enableUltraSSD</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableUltraSSD: Whether to enable UltraSSD</p>
</td>
</tr>
<tr>
<td>
<code>gpuInstanceProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.GPUInstanceProfile">
GPUInstanceProfile
</a>
</em>
</td>
<td>
<p>GpuInstanceProfile: GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.</p>
</td>
</tr>
<tr>
<td>
<code>hostGroupReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>HostGroupReference: This is of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Compute/&#x200b;hostGroups/&#x200b;{hostGroupName}.
For more information see <a href="https://docs.microsoft.com/azure/virtual-machines/dedicated-hosts">Azure dedicated hosts</a>.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.KubeletConfig">
KubeletConfig
</a>
</em>
</td>
<td>
<p>KubeletConfig: The Kubelet configuration on the agent pool nodes.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.KubeletDiskType">
KubeletDiskType
</a>
</em>
</td>
<td>
<p>KubeletDiskType: Determines the placement of emptyDir volumes, container runtime data root, and Kubelet ephemeral
storage.</p>
</td>
</tr>
<tr>
<td>
<code>linuxOSConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.LinuxOSConfig">
LinuxOSConfig
</a>
</em>
</td>
<td>
<p>LinuxOSConfig: The OS configuration of Linux agent nodes.</p>
</td>
</tr>
<tr>
<td>
<code>maxCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxCount: The maximum number of nodes for auto-scaling</p>
</td>
</tr>
<tr>
<td>
<code>maxPods</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxPods: The maximum number of pods that can run on a node.</p>
</td>
</tr>
<tr>
<td>
<code>messageOfTheDay</code><br/>
<em>
string
</em>
</td>
<td>
<p>MessageOfTheDay: A base64-encoded string which will be written to /etc/motd after decoding. This allows customization of
the message of the day for Linux nodes. It must not be specified for Windows nodes. It must be a static string (i.e.,
will be printed raw and not be executed as a script).</p>
</td>
</tr>
<tr>
<td>
<code>minCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MinCount: The minimum number of nodes for auto-scaling</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolMode">
AgentPoolMode
</a>
</em>
</td>
<td>
<p>Mode: A cluster must have at least one &lsquo;System&rsquo; Agent Pool at all times. For additional information on agent pool
restrictions  and best practices, see: <a href="https://docs.microsoft.com/azure/aks/use-system-pools">https://docs.microsoft.com/azure/aks/use-system-pools</a></p>
</td>
</tr>
<tr>
<td>
<code>networkProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolNetworkProfile">
AgentPoolNetworkProfile
</a>
</em>
</td>
<td>
<p>NetworkProfile: Network-related settings of an agent pool.</p>
</td>
</tr>
<tr>
<td>
<code>nodeLabels</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>NodeLabels: The node labels to be persisted across all nodes in agent pool.</p>
</td>
</tr>
<tr>
<td>
<code>nodePublicIPPrefixReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>NodePublicIPPrefixReference: This is of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Network/&#x200b;publicIPPrefixes/&#x200b;{publicIPPrefixName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>nodeTaints</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>NodeTaints: The taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.</p>
</td>
</tr>
<tr>
<td>
<code>orchestratorVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>OrchestratorVersion: Both patch version <major.minor.patch> and <major.minor> are supported. When <major.minor> is
specified, the latest supported patch version is chosen automatically. Updating the agent pool with the same
<major.minor> once it has been created will not trigger an upgrade, even if a newer patch version is available. As a
best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes version. The node pool version
must have the same major version as the control plane. The node pool minor version must be within two minor versions of
the control plane version. The node pool version cannot be greater than the control plane version. For more information
see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool">upgrading a node pool</a>.</p>
</td>
</tr>
<tr>
<td>
<code>osDiskSizeGB</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceOSDisk">
ContainerServiceOSDisk
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSDiskType">
OSDiskType
</a>
</em>
</td>
<td>
<p>OsDiskType: The default is &lsquo;Ephemeral&rsquo; if the VM supports it and has a cache disk larger than the requested
OSDiskSizeGB. Otherwise,  defaults to &lsquo;Managed&rsquo;. May not be changed after creation. For more information see <a href="https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os">Ephemeral
OS</a>.</p>
</td>
</tr>
<tr>
<td>
<code>osSKU</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSSKU">
OSSKU
</a>
</em>
</td>
<td>
<p>OsSKU: Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or
Windows2019 if  OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is
deprecated.</p>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.OSType">
OSType
</a>
</em>
</td>
<td>
<p>OsType: The operating system type. The default is Linux.</p>
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
reference to a containerservice.azure.com/ManagedCluster resource</p>
</td>
</tr>
<tr>
<td>
<code>podSubnetReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>PodSubnetReference: If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more details).
This is of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Network/&#x200b;virtualNetworks/&#x200b;{virtualNetworkName}/&#x200b;subnets/&#x200b;{subnetName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>powerState</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PowerState">
PowerState
</a>
</em>
</td>
<td>
<p>PowerState: When an Agent Pool is first created it is initially Running. The Agent Pool can be stopped by setting this
field to Stopped. A stopped Agent Pool stops all of its VMs and does not accrue billing charges. An Agent Pool can only
be stopped if it is Running and provisioning state is Succeeded</p>
</td>
</tr>
<tr>
<td>
<code>proximityPlacementGroupReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>ProximityPlacementGroupReference: The ID for Proximity Placement Group.</p>
</td>
</tr>
<tr>
<td>
<code>scaleDownMode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleDownMode">
ScaleDownMode
</a>
</em>
</td>
<td>
<p>ScaleDownMode: This also effects the cluster autoscaler behavior. If not specified, it defaults to Delete.</p>
</td>
</tr>
<tr>
<td>
<code>scaleSetEvictionPolicy</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleSetEvictionPolicy">
ScaleSetEvictionPolicy
</a>
</em>
</td>
<td>
<p>ScaleSetEvictionPolicy: This cannot be specified unless the scaleSetPriority is &lsquo;Spot&rsquo;. If not specified, the default is
&lsquo;Delete&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>scaleSetPriority</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ScaleSetPriority">
ScaleSetPriority
</a>
</em>
</td>
<td>
<p>ScaleSetPriority: The Virtual Machine Scale Set priority. If not specified, the default is &lsquo;Regular&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>spotMaxPrice</code><br/>
<em>
float64
</em>
</td>
<td>
<p>SpotMaxPrice: Possible values are any decimal value greater than zero or -1 which indicates the willingness to pay any
on-demand price. For more details on spot pricing, see <a href="https://docs.microsoft.com/azure/virtual-machines/spot-vms#pricing">spot VMs
pricing</a></p>
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
<p>Tags: The tags to be persisted on the agent pool virtual machine scale set.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolType">
AgentPoolType
</a>
</em>
</td>
<td>
<p>Type: The type of Agent Pool.</p>
</td>
</tr>
<tr>
<td>
<code>upgradeSettings</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolUpgradeSettings">
AgentPoolUpgradeSettings
</a>
</em>
</td>
<td>
<p>UpgradeSettings: Settings for upgrading the agentpool</p>
</td>
</tr>
<tr>
<td>
<code>vmSize</code><br/>
<em>
string
</em>
</td>
<td>
<p>VmSize: VM size availability varies by region. If a node contains insufficient compute resources (memory, cpu, etc) pods
might fail to run correctly. For more details on restricted VM sizes, see:
<a href="https://docs.microsoft.com/azure/aks/quotas-skus-regions">https://docs.microsoft.com/azure/aks/quotas-skus-regions</a></p>
</td>
</tr>
<tr>
<td>
<code>vnetSubnetReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>VnetSubnetReference: If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is
specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Network/&#x200b;virtualNetworks/&#x200b;{virtualNetworkName}/&#x200b;subnets/&#x200b;{subnetName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>windowsProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.AgentPoolWindowsProfile">
AgentPoolWindowsProfile
</a>
</em>
</td>
<td>
<p>WindowsProfile: The Windows agent pool&rsquo;s specific profile.</p>
</td>
</tr>
<tr>
<td>
<code>workloadRuntime</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.WorkloadRuntime">
WorkloadRuntime
</a>
</em>
</td>
<td>
<p>WorkloadRuntime: Determines the type of workload a node can run.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_Spec_ARM">ManagedClusters_AgentPool_Spec_ARM
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
<a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_ARM">
ManagedClusterAgentPoolProfileProperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of an agent pool.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.NetworkDataplane">NetworkDataplane
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile">ContainerServiceNetworkProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_ARM">ContainerServiceNetworkProfile_ARM</a>)
</p>
<div>
<p>Network dataplane used in the Kubernetes cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;azure&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;cilium&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.NetworkDataplane_STATUS">NetworkDataplane_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_STATUS">ContainerServiceNetworkProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_STATUS_ARM">ContainerServiceNetworkProfile_STATUS_ARM</a>)
</p>
<div>
<p>Network dataplane used in the Kubernetes cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;azure&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;cilium&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.NetworkMode">NetworkMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile">ContainerServiceNetworkProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_ARM">ContainerServiceNetworkProfile_ARM</a>)
</p>
<div>
<p>This cannot be specified if networkPlugin is anything other than &lsquo;azure&rsquo;.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;bridge&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;transparent&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.NetworkMode_STATUS">NetworkMode_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_STATUS">ContainerServiceNetworkProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_STATUS_ARM">ContainerServiceNetworkProfile_STATUS_ARM</a>)
</p>
<div>
<p>This cannot be specified if networkPlugin is anything other than &lsquo;azure&rsquo;.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;bridge&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;transparent&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.NetworkPlugin">NetworkPlugin
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile">ContainerServiceNetworkProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_ARM">ContainerServiceNetworkProfile_ARM</a>)
</p>
<div>
<p>Network plugin used for building the Kubernetes network.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;azure&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;kubenet&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;none&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.NetworkPluginMode">NetworkPluginMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile">ContainerServiceNetworkProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_ARM">ContainerServiceNetworkProfile_ARM</a>)
</p>
<div>
<p>The mode the network plugin should use.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Overlay&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.NetworkPluginMode_STATUS">NetworkPluginMode_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_STATUS">ContainerServiceNetworkProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_STATUS_ARM">ContainerServiceNetworkProfile_STATUS_ARM</a>)
</p>
<div>
<p>The mode the network plugin should use.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Overlay&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.NetworkPlugin_STATUS">NetworkPlugin_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_STATUS">ContainerServiceNetworkProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_STATUS_ARM">ContainerServiceNetworkProfile_STATUS_ARM</a>)
</p>
<div>
<p>Network plugin used for building the Kubernetes network.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;azure&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;kubenet&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;none&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.NetworkPolicy">NetworkPolicy
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile">ContainerServiceNetworkProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_ARM">ContainerServiceNetworkProfile_ARM</a>)
</p>
<div>
<p>Network policy used for building the Kubernetes network.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;azure&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;calico&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;cilium&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.NetworkPolicy_STATUS">NetworkPolicy_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_STATUS">ContainerServiceNetworkProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ContainerServiceNetworkProfile_STATUS_ARM">ContainerServiceNetworkProfile_STATUS_ARM</a>)
</p>
<div>
<p>Network policy used for building the Kubernetes network.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;azure&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;calico&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;cilium&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.OSDiskType">OSDiskType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_ARM">ManagedClusterAgentPoolProfileProperties_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_ARM">ManagedClusterAgentPoolProfile_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_Spec">ManagedClusters_AgentPool_Spec</a>)
</p>
<div>
<p>The default is &lsquo;Ephemeral&rsquo; if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise,
defaults to &lsquo;Managed&rsquo;. May not be changed after creation. For more information see <a href="https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os">Ephemeral
OS</a>.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Ephemeral&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Managed&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.OSDiskType_STATUS">OSDiskType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_STATUS_ARM">ManagedClusterAgentPoolProfileProperties_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS">ManagedClusterAgentPoolProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS_ARM">ManagedClusterAgentPoolProfile_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_STATUS">ManagedClusters_AgentPool_STATUS</a>)
</p>
<div>
<p>The default is &lsquo;Ephemeral&rsquo; if the VM supports it and has a cache disk larger than the requested OSDiskSizeGB. Otherwise,
defaults to &lsquo;Managed&rsquo;. May not be changed after creation. For more information see <a href="https://docs.microsoft.com/azure/aks/cluster-configuration#ephemeral-os">Ephemeral
OS</a>.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Ephemeral&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Managed&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.OSSKU">OSSKU
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_ARM">ManagedClusterAgentPoolProfileProperties_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_ARM">ManagedClusterAgentPoolProfile_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_Spec">ManagedClusters_AgentPool_Spec</a>)
</p>
<div>
<p>Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or Windows2019 if
OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is deprecated.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;CBLMariner&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Mariner&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Ubuntu&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Windows2019&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Windows2022&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.OSSKU_STATUS">OSSKU_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_STATUS_ARM">ManagedClusterAgentPoolProfileProperties_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS">ManagedClusterAgentPoolProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS_ARM">ManagedClusterAgentPoolProfile_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_STATUS">ManagedClusters_AgentPool_STATUS</a>)
</p>
<div>
<p>Specifies the OS SKU used by the agent pool. If not specified, the default is Ubuntu if OSType=Linux or Windows2019 if
OSType=Windows. And the default Windows OSSKU will be changed to Windows2022 after Windows2019 is deprecated.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;CBLMariner&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Mariner&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Ubuntu&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Windows2019&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Windows2022&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.OSType">OSType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_ARM">ManagedClusterAgentPoolProfileProperties_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_ARM">ManagedClusterAgentPoolProfile_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_Spec">ManagedClusters_AgentPool_Spec</a>)
</p>
<div>
<p>The operating system type. The default is Linux.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Linux&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Windows&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.OSType_STATUS">OSType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_STATUS_ARM">ManagedClusterAgentPoolProfileProperties_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS">ManagedClusterAgentPoolProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS_ARM">ManagedClusterAgentPoolProfile_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_STATUS">ManagedClusters_AgentPool_STATUS</a>)
</p>
<div>
<p>The operating system type. The default is Linux.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Linux&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Windows&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.PortRange">PortRange
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.AgentPoolNetworkProfile">AgentPoolNetworkProfile</a>)
</p>
<div>
<p>The port range.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>portEnd</code><br/>
<em>
int
</em>
</td>
<td>
<p>PortEnd: The maximum port that is included in the range. It should be ranged from 1 to 65535, and be greater than or
equal to portStart.</p>
</td>
</tr>
<tr>
<td>
<code>portStart</code><br/>
<em>
int
</em>
</td>
<td>
<p>PortStart: The minimum port that is included in the range. It should be ranged from 1 to 65535, and be less than or
equal to portEnd.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PortRange_Protocol">
PortRange_Protocol
</a>
</em>
</td>
<td>
<p>Protocol: The network protocol of the port.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.PortRange_ARM">PortRange_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.AgentPoolNetworkProfile_ARM">AgentPoolNetworkProfile_ARM</a>)
</p>
<div>
<p>The port range.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>portEnd</code><br/>
<em>
int
</em>
</td>
<td>
<p>PortEnd: The maximum port that is included in the range. It should be ranged from 1 to 65535, and be greater than or
equal to portStart.</p>
</td>
</tr>
<tr>
<td>
<code>portStart</code><br/>
<em>
int
</em>
</td>
<td>
<p>PortStart: The minimum port that is included in the range. It should be ranged from 1 to 65535, and be less than or
equal to portEnd.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PortRange_Protocol">
PortRange_Protocol
</a>
</em>
</td>
<td>
<p>Protocol: The network protocol of the port.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.PortRange_Protocol">PortRange_Protocol
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.PortRange">PortRange</a>, <a href="#containerservice.azure.com/v1api20230202preview.PortRange_ARM">PortRange_ARM</a>)
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
<tbody><tr><td><p>&#34;TCP&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UDP&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.PortRange_Protocol_STATUS">PortRange_Protocol_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.PortRange_STATUS">PortRange_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.PortRange_STATUS_ARM">PortRange_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;TCP&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UDP&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.PortRange_STATUS">PortRange_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.AgentPoolNetworkProfile_STATUS">AgentPoolNetworkProfile_STATUS</a>)
</p>
<div>
<p>The port range.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>portEnd</code><br/>
<em>
int
</em>
</td>
<td>
<p>PortEnd: The maximum port that is included in the range. It should be ranged from 1 to 65535, and be greater than or
equal to portStart.</p>
</td>
</tr>
<tr>
<td>
<code>portStart</code><br/>
<em>
int
</em>
</td>
<td>
<p>PortStart: The minimum port that is included in the range. It should be ranged from 1 to 65535, and be less than or
equal to portEnd.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PortRange_Protocol_STATUS">
PortRange_Protocol_STATUS
</a>
</em>
</td>
<td>
<p>Protocol: The network protocol of the port.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.PortRange_STATUS_ARM">PortRange_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.AgentPoolNetworkProfile_STATUS_ARM">AgentPoolNetworkProfile_STATUS_ARM</a>)
</p>
<div>
<p>The port range.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>portEnd</code><br/>
<em>
int
</em>
</td>
<td>
<p>PortEnd: The maximum port that is included in the range. It should be ranged from 1 to 65535, and be greater than or
equal to portStart.</p>
</td>
</tr>
<tr>
<td>
<code>portStart</code><br/>
<em>
int
</em>
</td>
<td>
<p>PortStart: The minimum port that is included in the range. It should be ranged from 1 to 65535, and be less than or
equal to portEnd.</p>
</td>
</tr>
<tr>
<td>
<code>protocol</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PortRange_Protocol_STATUS">
PortRange_Protocol_STATUS
</a>
</em>
</td>
<td>
<p>Protocol: The network protocol of the port.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.PowerState">PowerState
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_Spec">ManagedClusters_AgentPool_Spec</a>)
</p>
<div>
<p>Describes the Power State of the cluster</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>code</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PowerState_Code">
PowerState_Code
</a>
</em>
</td>
<td>
<p>Code: Tells whether the cluster is Running or Stopped</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.PowerState_ARM">PowerState_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_ARM">ManagedClusterAgentPoolProfileProperties_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_ARM">ManagedClusterAgentPoolProfile_ARM</a>)
</p>
<div>
<p>Describes the Power State of the cluster</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>code</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PowerState_Code">
PowerState_Code
</a>
</em>
</td>
<td>
<p>Code: Tells whether the cluster is Running or Stopped</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.PowerState_Code">PowerState_Code
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.PowerState">PowerState</a>, <a href="#containerservice.azure.com/v1api20230202preview.PowerState_ARM">PowerState_ARM</a>)
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
<tbody><tr><td><p>&#34;Running&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Stopped&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.PowerState_Code_STATUS">PowerState_Code_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.PowerState_STATUS">PowerState_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.PowerState_STATUS_ARM">PowerState_STATUS_ARM</a>)
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
<tbody><tr><td><p>&#34;Running&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Stopped&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.PowerState_STATUS">PowerState_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS">ManagedClusterAgentPoolProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_STATUS">ManagedClusters_AgentPool_STATUS</a>)
</p>
<div>
<p>Describes the Power State of the cluster</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>code</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PowerState_Code_STATUS">
PowerState_Code_STATUS
</a>
</em>
</td>
<td>
<p>Code: Tells whether the cluster is Running or Stopped</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.PowerState_STATUS_ARM">PowerState_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_STATUS_ARM">ManagedClusterAgentPoolProfileProperties_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS_ARM">ManagedClusterAgentPoolProfile_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>)
</p>
<div>
<p>Describes the Power State of the cluster</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>code</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.PowerState_Code_STATUS">
PowerState_Code_STATUS
</a>
</em>
</td>
<td>
<p>Code: Tells whether the cluster is Running or Stopped</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.PrivateLinkResource">PrivateLinkResource
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
</p>
<div>
<p>A private link resource</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>groupId</code><br/>
<em>
string
</em>
</td>
<td>
<p>GroupId: The group ID of the resource.</p>
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
<p>Name: The name of the private link resource.</p>
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
<p>Reference: The ID of the private link resource.</p>
</td>
</tr>
<tr>
<td>
<code>requiredMembers</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>RequiredMembers: The RequiredMembers of the resource</p>
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
<p>Type: The resource type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.PrivateLinkResource_ARM">PrivateLinkResource_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM</a>)
</p>
<div>
<p>A private link resource</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>groupId</code><br/>
<em>
string
</em>
</td>
<td>
<p>GroupId: The group ID of the resource.</p>
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
<p>Name: The name of the private link resource.</p>
</td>
</tr>
<tr>
<td>
<code>requiredMembers</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>RequiredMembers: The RequiredMembers of the resource</p>
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
<p>Type: The resource type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.PrivateLinkResource_STATUS">PrivateLinkResource_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
</p>
<div>
<p>A private link resource</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>groupId</code><br/>
<em>
string
</em>
</td>
<td>
<p>GroupId: The group ID of the resource.</p>
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
<p>Id: The ID of the private link resource.</p>
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
<p>Name: The name of the private link resource.</p>
</td>
</tr>
<tr>
<td>
<code>privateLinkServiceID</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateLinkServiceID: The private link service ID of the resource, this field is exposed only to NRP internally.</p>
</td>
</tr>
<tr>
<td>
<code>requiredMembers</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>RequiredMembers: The RequiredMembers of the resource</p>
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
<p>Type: The resource type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.PrivateLinkResource_STATUS_ARM">PrivateLinkResource_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>)
</p>
<div>
<p>A private link resource</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>groupId</code><br/>
<em>
string
</em>
</td>
<td>
<p>GroupId: The group ID of the resource.</p>
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
<p>Id: The ID of the private link resource.</p>
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
<p>Name: The name of the private link resource.</p>
</td>
</tr>
<tr>
<td>
<code>privateLinkServiceID</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateLinkServiceID: The private link service ID of the resource, this field is exposed only to NRP internally.</p>
</td>
</tr>
<tr>
<td>
<code>requiredMembers</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>RequiredMembers: The RequiredMembers of the resource</p>
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
<p>Type: The resource type.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ResourceReference">ResourceReference
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile">ManagedClusterLoadBalancerProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_OutboundIPPrefixes">ManagedClusterLoadBalancerProfile_OutboundIPPrefixes</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_OutboundIPs">ManagedClusterLoadBalancerProfile_OutboundIPs</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterNATGatewayProfile">ManagedClusterNATGatewayProfile</a>)
</p>
<div>
<p>A reference to an Azure resource.</p>
</div>
<table>
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
<p>Reference: The fully qualified Azure resource id.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ResourceReference_ARM">ResourceReference_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_ARM">ManagedClusterLoadBalancerProfile_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_ARM">ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_OutboundIPs_ARM">ManagedClusterLoadBalancerProfile_OutboundIPs_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterNATGatewayProfile_ARM">ManagedClusterNATGatewayProfile_ARM</a>)
</p>
<div>
<p>A reference to an Azure resource.</p>
</div>
<table>
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
<h3 id="containerservice.azure.com/v1api20230202preview.ResourceReference_STATUS">ResourceReference_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS">ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS">ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_STATUS">ManagedClusterLoadBalancerProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterNATGatewayProfile_STATUS">ManagedClusterNATGatewayProfile_STATUS</a>)
</p>
<div>
<p>A reference to an Azure resource.</p>
</div>
<table>
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
<p>Id: The fully qualified Azure resource id.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ResourceReference_STATUS_ARM">ResourceReference_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS_ARM">ManagedClusterLoadBalancerProfile_OutboundIPPrefixes_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS_ARM">ManagedClusterLoadBalancerProfile_OutboundIPs_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterLoadBalancerProfile_STATUS_ARM">ManagedClusterLoadBalancerProfile_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterNATGatewayProfile_STATUS_ARM">ManagedClusterNATGatewayProfile_STATUS_ARM</a>)
</p>
<div>
<p>A reference to an Azure resource.</p>
</div>
<table>
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
<p>Id: The fully qualified Azure resource id.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ScaleDownMode">ScaleDownMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_ARM">ManagedClusterAgentPoolProfileProperties_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_ARM">ManagedClusterAgentPoolProfile_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_Spec">ManagedClusters_AgentPool_Spec</a>)
</p>
<div>
<p>Describes how VMs are added to or removed from Agent Pools. See <a href="https://docs.microsoft.com/azure/virtual-machines/states-billing">billing
states</a>.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Deallocate&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Delete&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ScaleDownMode_STATUS">ScaleDownMode_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_STATUS_ARM">ManagedClusterAgentPoolProfileProperties_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS">ManagedClusterAgentPoolProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS_ARM">ManagedClusterAgentPoolProfile_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_STATUS">ManagedClusters_AgentPool_STATUS</a>)
</p>
<div>
<p>Describes how VMs are added to or removed from Agent Pools. See <a href="https://docs.microsoft.com/azure/virtual-machines/states-billing">billing
states</a>.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Deallocate&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Delete&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ScaleSetEvictionPolicy">ScaleSetEvictionPolicy
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_ARM">ManagedClusterAgentPoolProfileProperties_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_ARM">ManagedClusterAgentPoolProfile_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_Spec">ManagedClusters_AgentPool_Spec</a>)
</p>
<div>
<p>The eviction policy specifies what to do with the VM when it is evicted. The default is Delete. For more information
about eviction see <a href="https://docs.microsoft.com/azure/virtual-machines/spot-vms">spot VMs</a></p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Deallocate&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Delete&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ScaleSetEvictionPolicy_STATUS">ScaleSetEvictionPolicy_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_STATUS_ARM">ManagedClusterAgentPoolProfileProperties_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS">ManagedClusterAgentPoolProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS_ARM">ManagedClusterAgentPoolProfile_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_STATUS">ManagedClusters_AgentPool_STATUS</a>)
</p>
<div>
<p>The eviction policy specifies what to do with the VM when it is evicted. The default is Delete. For more information
about eviction see <a href="https://docs.microsoft.com/azure/virtual-machines/spot-vms">spot VMs</a></p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Deallocate&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Delete&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ScaleSetPriority">ScaleSetPriority
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_ARM">ManagedClusterAgentPoolProfileProperties_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_ARM">ManagedClusterAgentPoolProfile_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_Spec">ManagedClusters_AgentPool_Spec</a>)
</p>
<div>
<p>The Virtual Machine Scale Set priority.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Regular&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Spot&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ScaleSetPriority_STATUS">ScaleSetPriority_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_STATUS_ARM">ManagedClusterAgentPoolProfileProperties_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS">ManagedClusterAgentPoolProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS_ARM">ManagedClusterAgentPoolProfile_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_STATUS">ManagedClusters_AgentPool_STATUS</a>)
</p>
<div>
<p>The Virtual Machine Scale Set priority.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Regular&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Spot&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ServiceMeshProfile">ServiceMeshProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
</p>
<div>
<p>Service mesh profile for a managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>istio</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.IstioServiceMesh">
IstioServiceMesh
</a>
</em>
</td>
<td>
<p>Istio: Istio service mesh configuration.</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ServiceMeshProfile_Mode">
ServiceMeshProfile_Mode
</a>
</em>
</td>
<td>
<p>Mode: Mode of the service mesh.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ServiceMeshProfile_ARM">ServiceMeshProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM</a>)
</p>
<div>
<p>Service mesh profile for a managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>istio</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.IstioServiceMesh_ARM">
IstioServiceMesh_ARM
</a>
</em>
</td>
<td>
<p>Istio: Istio service mesh configuration.</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ServiceMeshProfile_Mode">
ServiceMeshProfile_Mode
</a>
</em>
</td>
<td>
<p>Mode: Mode of the service mesh.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ServiceMeshProfile_Mode">ServiceMeshProfile_Mode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ServiceMeshProfile">ServiceMeshProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ServiceMeshProfile_ARM">ServiceMeshProfile_ARM</a>)
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
</tr><tr><td><p>&#34;Istio&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ServiceMeshProfile_Mode_STATUS">ServiceMeshProfile_Mode_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ServiceMeshProfile_STATUS">ServiceMeshProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ServiceMeshProfile_STATUS_ARM">ServiceMeshProfile_STATUS_ARM</a>)
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
</tr><tr><td><p>&#34;Istio&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ServiceMeshProfile_STATUS">ServiceMeshProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
</p>
<div>
<p>Service mesh profile for a managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>istio</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.IstioServiceMesh_STATUS">
IstioServiceMesh_STATUS
</a>
</em>
</td>
<td>
<p>Istio: Istio service mesh configuration.</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ServiceMeshProfile_Mode_STATUS">
ServiceMeshProfile_Mode_STATUS
</a>
</em>
</td>
<td>
<p>Mode: Mode of the service mesh.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.ServiceMeshProfile_STATUS_ARM">ServiceMeshProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>)
</p>
<div>
<p>Service mesh profile for a managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>istio</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.IstioServiceMesh_STATUS_ARM">
IstioServiceMesh_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Istio: Istio service mesh configuration.</p>
</td>
</tr>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ServiceMeshProfile_Mode_STATUS">
ServiceMeshProfile_Mode_STATUS
</a>
</em>
</td>
<td>
<p>Mode: Mode of the service mesh.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.SysctlConfig">SysctlConfig
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.LinuxOSConfig">LinuxOSConfig</a>)
</p>
<div>
<p>Sysctl settings for Linux agent nodes.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>fsAioMaxNr</code><br/>
<em>
int
</em>
</td>
<td>
<p>FsAioMaxNr: Sysctl setting fs.aio-max-nr.</p>
</td>
</tr>
<tr>
<td>
<code>fsFileMax</code><br/>
<em>
int
</em>
</td>
<td>
<p>FsFileMax: Sysctl setting fs.file-max.</p>
</td>
</tr>
<tr>
<td>
<code>fsInotifyMaxUserWatches</code><br/>
<em>
int
</em>
</td>
<td>
<p>FsInotifyMaxUserWatches: Sysctl setting fs.inotify.max_user_watches.</p>
</td>
</tr>
<tr>
<td>
<code>fsNrOpen</code><br/>
<em>
int
</em>
</td>
<td>
<p>FsNrOpen: Sysctl setting fs.nr_open.</p>
</td>
</tr>
<tr>
<td>
<code>kernelThreadsMax</code><br/>
<em>
int
</em>
</td>
<td>
<p>KernelThreadsMax: Sysctl setting kernel.threads-max.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreNetdevMaxBacklog</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreNetdevMaxBacklog: Sysctl setting net.core.netdev_max_backlog.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreOptmemMax</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreOptmemMax: Sysctl setting net.core.optmem_max.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreRmemDefault</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreRmemDefault: Sysctl setting net.core.rmem_default.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreRmemMax</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreRmemMax: Sysctl setting net.core.rmem_max.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreSomaxconn</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreSomaxconn: Sysctl setting net.core.somaxconn.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreWmemDefault</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreWmemDefault: Sysctl setting net.core.wmem_default.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreWmemMax</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreWmemMax: Sysctl setting net.core.wmem_max.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4IpLocalPortRange</code><br/>
<em>
string
</em>
</td>
<td>
<p>NetIpv4IpLocalPortRange: Sysctl setting net.ipv4.ip_local_port_range.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4NeighDefaultGcThresh1</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4NeighDefaultGcThresh1: Sysctl setting net.ipv4.neigh.default.gc_thresh1.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4NeighDefaultGcThresh2</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4NeighDefaultGcThresh2: Sysctl setting net.ipv4.neigh.default.gc_thresh2.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4NeighDefaultGcThresh3</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4NeighDefaultGcThresh3: Sysctl setting net.ipv4.neigh.default.gc_thresh3.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpFinTimeout</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4TcpFinTimeout: Sysctl setting net.ipv4.tcp_fin_timeout.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpKeepaliveProbes</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4TcpKeepaliveProbes: Sysctl setting net.ipv4.tcp_keepalive_probes.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpKeepaliveTime</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4TcpKeepaliveTime: Sysctl setting net.ipv4.tcp_keepalive_time.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpMaxSynBacklog</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4TcpMaxSynBacklog: Sysctl setting net.ipv4.tcp_max_syn_backlog.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpMaxTwBuckets</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4TcpMaxTwBuckets: Sysctl setting net.ipv4.tcp_max_tw_buckets.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpTwReuse</code><br/>
<em>
bool
</em>
</td>
<td>
<p>NetIpv4TcpTwReuse: Sysctl setting net.ipv4.tcp_tw_reuse.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpkeepaliveIntvl</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4TcpkeepaliveIntvl: Sysctl setting net.ipv4.tcp_keepalive_intvl.</p>
</td>
</tr>
<tr>
<td>
<code>netNetfilterNfConntrackBuckets</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetNetfilterNfConntrackBuckets: Sysctl setting net.netfilter.nf_conntrack_buckets.</p>
</td>
</tr>
<tr>
<td>
<code>netNetfilterNfConntrackMax</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetNetfilterNfConntrackMax: Sysctl setting net.netfilter.nf_conntrack_max.</p>
</td>
</tr>
<tr>
<td>
<code>vmMaxMapCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>VmMaxMapCount: Sysctl setting vm.max_map_count.</p>
</td>
</tr>
<tr>
<td>
<code>vmSwappiness</code><br/>
<em>
int
</em>
</td>
<td>
<p>VmSwappiness: Sysctl setting vm.swappiness.</p>
</td>
</tr>
<tr>
<td>
<code>vmVfsCachePressure</code><br/>
<em>
int
</em>
</td>
<td>
<p>VmVfsCachePressure: Sysctl setting vm.vfs_cache_pressure.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.SysctlConfig_ARM">SysctlConfig_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.LinuxOSConfig_ARM">LinuxOSConfig_ARM</a>)
</p>
<div>
<p>Sysctl settings for Linux agent nodes.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>fsAioMaxNr</code><br/>
<em>
int
</em>
</td>
<td>
<p>FsAioMaxNr: Sysctl setting fs.aio-max-nr.</p>
</td>
</tr>
<tr>
<td>
<code>fsFileMax</code><br/>
<em>
int
</em>
</td>
<td>
<p>FsFileMax: Sysctl setting fs.file-max.</p>
</td>
</tr>
<tr>
<td>
<code>fsInotifyMaxUserWatches</code><br/>
<em>
int
</em>
</td>
<td>
<p>FsInotifyMaxUserWatches: Sysctl setting fs.inotify.max_user_watches.</p>
</td>
</tr>
<tr>
<td>
<code>fsNrOpen</code><br/>
<em>
int
</em>
</td>
<td>
<p>FsNrOpen: Sysctl setting fs.nr_open.</p>
</td>
</tr>
<tr>
<td>
<code>kernelThreadsMax</code><br/>
<em>
int
</em>
</td>
<td>
<p>KernelThreadsMax: Sysctl setting kernel.threads-max.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreNetdevMaxBacklog</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreNetdevMaxBacklog: Sysctl setting net.core.netdev_max_backlog.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreOptmemMax</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreOptmemMax: Sysctl setting net.core.optmem_max.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreRmemDefault</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreRmemDefault: Sysctl setting net.core.rmem_default.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreRmemMax</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreRmemMax: Sysctl setting net.core.rmem_max.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreSomaxconn</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreSomaxconn: Sysctl setting net.core.somaxconn.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreWmemDefault</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreWmemDefault: Sysctl setting net.core.wmem_default.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreWmemMax</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreWmemMax: Sysctl setting net.core.wmem_max.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4IpLocalPortRange</code><br/>
<em>
string
</em>
</td>
<td>
<p>NetIpv4IpLocalPortRange: Sysctl setting net.ipv4.ip_local_port_range.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4NeighDefaultGcThresh1</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4NeighDefaultGcThresh1: Sysctl setting net.ipv4.neigh.default.gc_thresh1.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4NeighDefaultGcThresh2</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4NeighDefaultGcThresh2: Sysctl setting net.ipv4.neigh.default.gc_thresh2.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4NeighDefaultGcThresh3</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4NeighDefaultGcThresh3: Sysctl setting net.ipv4.neigh.default.gc_thresh3.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpFinTimeout</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4TcpFinTimeout: Sysctl setting net.ipv4.tcp_fin_timeout.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpKeepaliveProbes</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4TcpKeepaliveProbes: Sysctl setting net.ipv4.tcp_keepalive_probes.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpKeepaliveTime</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4TcpKeepaliveTime: Sysctl setting net.ipv4.tcp_keepalive_time.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpMaxSynBacklog</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4TcpMaxSynBacklog: Sysctl setting net.ipv4.tcp_max_syn_backlog.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpMaxTwBuckets</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4TcpMaxTwBuckets: Sysctl setting net.ipv4.tcp_max_tw_buckets.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpTwReuse</code><br/>
<em>
bool
</em>
</td>
<td>
<p>NetIpv4TcpTwReuse: Sysctl setting net.ipv4.tcp_tw_reuse.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpkeepaliveIntvl</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4TcpkeepaliveIntvl: Sysctl setting net.ipv4.tcp_keepalive_intvl.</p>
</td>
</tr>
<tr>
<td>
<code>netNetfilterNfConntrackBuckets</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetNetfilterNfConntrackBuckets: Sysctl setting net.netfilter.nf_conntrack_buckets.</p>
</td>
</tr>
<tr>
<td>
<code>netNetfilterNfConntrackMax</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetNetfilterNfConntrackMax: Sysctl setting net.netfilter.nf_conntrack_max.</p>
</td>
</tr>
<tr>
<td>
<code>vmMaxMapCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>VmMaxMapCount: Sysctl setting vm.max_map_count.</p>
</td>
</tr>
<tr>
<td>
<code>vmSwappiness</code><br/>
<em>
int
</em>
</td>
<td>
<p>VmSwappiness: Sysctl setting vm.swappiness.</p>
</td>
</tr>
<tr>
<td>
<code>vmVfsCachePressure</code><br/>
<em>
int
</em>
</td>
<td>
<p>VmVfsCachePressure: Sysctl setting vm.vfs_cache_pressure.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.SysctlConfig_STATUS">SysctlConfig_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.LinuxOSConfig_STATUS">LinuxOSConfig_STATUS</a>)
</p>
<div>
<p>Sysctl settings for Linux agent nodes.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>fsAioMaxNr</code><br/>
<em>
int
</em>
</td>
<td>
<p>FsAioMaxNr: Sysctl setting fs.aio-max-nr.</p>
</td>
</tr>
<tr>
<td>
<code>fsFileMax</code><br/>
<em>
int
</em>
</td>
<td>
<p>FsFileMax: Sysctl setting fs.file-max.</p>
</td>
</tr>
<tr>
<td>
<code>fsInotifyMaxUserWatches</code><br/>
<em>
int
</em>
</td>
<td>
<p>FsInotifyMaxUserWatches: Sysctl setting fs.inotify.max_user_watches.</p>
</td>
</tr>
<tr>
<td>
<code>fsNrOpen</code><br/>
<em>
int
</em>
</td>
<td>
<p>FsNrOpen: Sysctl setting fs.nr_open.</p>
</td>
</tr>
<tr>
<td>
<code>kernelThreadsMax</code><br/>
<em>
int
</em>
</td>
<td>
<p>KernelThreadsMax: Sysctl setting kernel.threads-max.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreNetdevMaxBacklog</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreNetdevMaxBacklog: Sysctl setting net.core.netdev_max_backlog.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreOptmemMax</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreOptmemMax: Sysctl setting net.core.optmem_max.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreRmemDefault</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreRmemDefault: Sysctl setting net.core.rmem_default.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreRmemMax</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreRmemMax: Sysctl setting net.core.rmem_max.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreSomaxconn</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreSomaxconn: Sysctl setting net.core.somaxconn.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreWmemDefault</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreWmemDefault: Sysctl setting net.core.wmem_default.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreWmemMax</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreWmemMax: Sysctl setting net.core.wmem_max.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4IpLocalPortRange</code><br/>
<em>
string
</em>
</td>
<td>
<p>NetIpv4IpLocalPortRange: Sysctl setting net.ipv4.ip_local_port_range.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4NeighDefaultGcThresh1</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4NeighDefaultGcThresh1: Sysctl setting net.ipv4.neigh.default.gc_thresh1.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4NeighDefaultGcThresh2</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4NeighDefaultGcThresh2: Sysctl setting net.ipv4.neigh.default.gc_thresh2.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4NeighDefaultGcThresh3</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4NeighDefaultGcThresh3: Sysctl setting net.ipv4.neigh.default.gc_thresh3.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpFinTimeout</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4TcpFinTimeout: Sysctl setting net.ipv4.tcp_fin_timeout.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpKeepaliveProbes</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4TcpKeepaliveProbes: Sysctl setting net.ipv4.tcp_keepalive_probes.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpKeepaliveTime</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4TcpKeepaliveTime: Sysctl setting net.ipv4.tcp_keepalive_time.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpMaxSynBacklog</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4TcpMaxSynBacklog: Sysctl setting net.ipv4.tcp_max_syn_backlog.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpMaxTwBuckets</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4TcpMaxTwBuckets: Sysctl setting net.ipv4.tcp_max_tw_buckets.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpTwReuse</code><br/>
<em>
bool
</em>
</td>
<td>
<p>NetIpv4TcpTwReuse: Sysctl setting net.ipv4.tcp_tw_reuse.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpkeepaliveIntvl</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4TcpkeepaliveIntvl: Sysctl setting net.ipv4.tcp_keepalive_intvl.</p>
</td>
</tr>
<tr>
<td>
<code>netNetfilterNfConntrackBuckets</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetNetfilterNfConntrackBuckets: Sysctl setting net.netfilter.nf_conntrack_buckets.</p>
</td>
</tr>
<tr>
<td>
<code>netNetfilterNfConntrackMax</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetNetfilterNfConntrackMax: Sysctl setting net.netfilter.nf_conntrack_max.</p>
</td>
</tr>
<tr>
<td>
<code>vmMaxMapCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>VmMaxMapCount: Sysctl setting vm.max_map_count.</p>
</td>
</tr>
<tr>
<td>
<code>vmSwappiness</code><br/>
<em>
int
</em>
</td>
<td>
<p>VmSwappiness: Sysctl setting vm.swappiness.</p>
</td>
</tr>
<tr>
<td>
<code>vmVfsCachePressure</code><br/>
<em>
int
</em>
</td>
<td>
<p>VmVfsCachePressure: Sysctl setting vm.vfs_cache_pressure.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.SysctlConfig_STATUS_ARM">SysctlConfig_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.LinuxOSConfig_STATUS_ARM">LinuxOSConfig_STATUS_ARM</a>)
</p>
<div>
<p>Sysctl settings for Linux agent nodes.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>fsAioMaxNr</code><br/>
<em>
int
</em>
</td>
<td>
<p>FsAioMaxNr: Sysctl setting fs.aio-max-nr.</p>
</td>
</tr>
<tr>
<td>
<code>fsFileMax</code><br/>
<em>
int
</em>
</td>
<td>
<p>FsFileMax: Sysctl setting fs.file-max.</p>
</td>
</tr>
<tr>
<td>
<code>fsInotifyMaxUserWatches</code><br/>
<em>
int
</em>
</td>
<td>
<p>FsInotifyMaxUserWatches: Sysctl setting fs.inotify.max_user_watches.</p>
</td>
</tr>
<tr>
<td>
<code>fsNrOpen</code><br/>
<em>
int
</em>
</td>
<td>
<p>FsNrOpen: Sysctl setting fs.nr_open.</p>
</td>
</tr>
<tr>
<td>
<code>kernelThreadsMax</code><br/>
<em>
int
</em>
</td>
<td>
<p>KernelThreadsMax: Sysctl setting kernel.threads-max.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreNetdevMaxBacklog</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreNetdevMaxBacklog: Sysctl setting net.core.netdev_max_backlog.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreOptmemMax</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreOptmemMax: Sysctl setting net.core.optmem_max.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreRmemDefault</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreRmemDefault: Sysctl setting net.core.rmem_default.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreRmemMax</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreRmemMax: Sysctl setting net.core.rmem_max.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreSomaxconn</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreSomaxconn: Sysctl setting net.core.somaxconn.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreWmemDefault</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreWmemDefault: Sysctl setting net.core.wmem_default.</p>
</td>
</tr>
<tr>
<td>
<code>netCoreWmemMax</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetCoreWmemMax: Sysctl setting net.core.wmem_max.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4IpLocalPortRange</code><br/>
<em>
string
</em>
</td>
<td>
<p>NetIpv4IpLocalPortRange: Sysctl setting net.ipv4.ip_local_port_range.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4NeighDefaultGcThresh1</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4NeighDefaultGcThresh1: Sysctl setting net.ipv4.neigh.default.gc_thresh1.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4NeighDefaultGcThresh2</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4NeighDefaultGcThresh2: Sysctl setting net.ipv4.neigh.default.gc_thresh2.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4NeighDefaultGcThresh3</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4NeighDefaultGcThresh3: Sysctl setting net.ipv4.neigh.default.gc_thresh3.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpFinTimeout</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4TcpFinTimeout: Sysctl setting net.ipv4.tcp_fin_timeout.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpKeepaliveProbes</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4TcpKeepaliveProbes: Sysctl setting net.ipv4.tcp_keepalive_probes.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpKeepaliveTime</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4TcpKeepaliveTime: Sysctl setting net.ipv4.tcp_keepalive_time.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpMaxSynBacklog</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4TcpMaxSynBacklog: Sysctl setting net.ipv4.tcp_max_syn_backlog.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpMaxTwBuckets</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4TcpMaxTwBuckets: Sysctl setting net.ipv4.tcp_max_tw_buckets.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpTwReuse</code><br/>
<em>
bool
</em>
</td>
<td>
<p>NetIpv4TcpTwReuse: Sysctl setting net.ipv4.tcp_tw_reuse.</p>
</td>
</tr>
<tr>
<td>
<code>netIpv4TcpkeepaliveIntvl</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetIpv4TcpkeepaliveIntvl: Sysctl setting net.ipv4.tcp_keepalive_intvl.</p>
</td>
</tr>
<tr>
<td>
<code>netNetfilterNfConntrackBuckets</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetNetfilterNfConntrackBuckets: Sysctl setting net.netfilter.nf_conntrack_buckets.</p>
</td>
</tr>
<tr>
<td>
<code>netNetfilterNfConntrackMax</code><br/>
<em>
int
</em>
</td>
<td>
<p>NetNetfilterNfConntrackMax: Sysctl setting net.netfilter.nf_conntrack_max.</p>
</td>
</tr>
<tr>
<td>
<code>vmMaxMapCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>VmMaxMapCount: Sysctl setting vm.max_map_count.</p>
</td>
</tr>
<tr>
<td>
<code>vmSwappiness</code><br/>
<em>
int
</em>
</td>
<td>
<p>VmSwappiness: Sysctl setting vm.swappiness.</p>
</td>
</tr>
<tr>
<td>
<code>vmVfsCachePressure</code><br/>
<em>
int
</em>
</td>
<td>
<p>VmVfsCachePressure: Sysctl setting vm.vfs_cache_pressure.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.SystemData_CreatedByType_STATUS">SystemData_CreatedByType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.SystemData_STATUS">SystemData_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.SystemData_STATUS_ARM">SystemData_STATUS_ARM</a>)
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
<h3 id="containerservice.azure.com/v1api20230202preview.SystemData_LastModifiedByType_STATUS">SystemData_LastModifiedByType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.SystemData_STATUS">SystemData_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.SystemData_STATUS_ARM">SystemData_STATUS_ARM</a>)
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
<h3 id="containerservice.azure.com/v1api20230202preview.SystemData_STATUS">SystemData_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
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
<a href="#containerservice.azure.com/v1api20230202preview.SystemData_CreatedByType_STATUS">
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
<a href="#containerservice.azure.com/v1api20230202preview.SystemData_LastModifiedByType_STATUS">
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
<h3 id="containerservice.azure.com/v1api20230202preview.SystemData_STATUS_ARM">SystemData_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS_ARM">ManagedCluster_STATUS_ARM</a>)
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
<a href="#containerservice.azure.com/v1api20230202preview.SystemData_CreatedByType_STATUS">
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
<a href="#containerservice.azure.com/v1api20230202preview.SystemData_LastModifiedByType_STATUS">
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
<h3 id="containerservice.azure.com/v1api20230202preview.UpgradeOverrideSettings">UpgradeOverrideSettings
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ClusterUpgradeSettings">ClusterUpgradeSettings</a>)
</p>
<div>
<p>Settings for overrides when upgrading a cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>controlPlaneOverrides</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ControlPlaneUpgradeOverride">
[]ControlPlaneUpgradeOverride
</a>
</em>
</td>
<td>
<p>ControlPlaneOverrides: List of upgrade overrides when upgrading a cluster&rsquo;s control plane.</p>
</td>
</tr>
<tr>
<td>
<code>until</code><br/>
<em>
string
</em>
</td>
<td>
<p>Until: Until when the overrides are effective. Note that this only matches the start time of an upgrade, and the
effectiveness won&rsquo;t change once an upgrade starts even if the <code>until</code> expires as upgrade proceeds. This field is not set
by default. It must be set for the overrides to take effect.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.UpgradeOverrideSettings_ARM">UpgradeOverrideSettings_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ClusterUpgradeSettings_ARM">ClusterUpgradeSettings_ARM</a>)
</p>
<div>
<p>Settings for overrides when upgrading a cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>controlPlaneOverrides</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ControlPlaneUpgradeOverride">
[]ControlPlaneUpgradeOverride
</a>
</em>
</td>
<td>
<p>ControlPlaneOverrides: List of upgrade overrides when upgrading a cluster&rsquo;s control plane.</p>
</td>
</tr>
<tr>
<td>
<code>until</code><br/>
<em>
string
</em>
</td>
<td>
<p>Until: Until when the overrides are effective. Note that this only matches the start time of an upgrade, and the
effectiveness won&rsquo;t change once an upgrade starts even if the <code>until</code> expires as upgrade proceeds. This field is not set
by default. It must be set for the overrides to take effect.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.UpgradeOverrideSettings_STATUS">UpgradeOverrideSettings_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ClusterUpgradeSettings_STATUS">ClusterUpgradeSettings_STATUS</a>)
</p>
<div>
<p>Settings for overrides when upgrading a cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>controlPlaneOverrides</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ControlPlaneUpgradeOverride_STATUS">
[]ControlPlaneUpgradeOverride_STATUS
</a>
</em>
</td>
<td>
<p>ControlPlaneOverrides: List of upgrade overrides when upgrading a cluster&rsquo;s control plane.</p>
</td>
</tr>
<tr>
<td>
<code>until</code><br/>
<em>
string
</em>
</td>
<td>
<p>Until: Until when the overrides are effective. Note that this only matches the start time of an upgrade, and the
effectiveness won&rsquo;t change once an upgrade starts even if the <code>until</code> expires as upgrade proceeds. This field is not set
by default. It must be set for the overrides to take effect.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.UpgradeOverrideSettings_STATUS_ARM">UpgradeOverrideSettings_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ClusterUpgradeSettings_STATUS_ARM">ClusterUpgradeSettings_STATUS_ARM</a>)
</p>
<div>
<p>Settings for overrides when upgrading a cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>controlPlaneOverrides</code><br/>
<em>
<a href="#containerservice.azure.com/v1api20230202preview.ControlPlaneUpgradeOverride_STATUS">
[]ControlPlaneUpgradeOverride_STATUS
</a>
</em>
</td>
<td>
<p>ControlPlaneOverrides: List of upgrade overrides when upgrading a cluster&rsquo;s control plane.</p>
</td>
</tr>
<tr>
<td>
<code>until</code><br/>
<em>
string
</em>
</td>
<td>
<p>Until: Until when the overrides are effective. Note that this only matches the start time of an upgrade, and the
effectiveness won&rsquo;t change once an upgrade starts even if the <code>until</code> expires as upgrade proceeds. This field is not set
by default. It must be set for the overrides to take effect.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.UserAssignedIdentity">UserAssignedIdentity
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentity">ManagedClusterPodIdentity</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_Spec">ManagedCluster_Spec</a>)
</p>
<div>
<p>Details about a user assigned identity.</p>
</div>
<table>
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
<p>ClientId: The client ID of the user assigned identity.</p>
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
<p>ObjectId: The object ID of the user assigned identity.</p>
</td>
</tr>
<tr>
<td>
<code>resourceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>ResourceReference: The resource ID of the user assigned identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.UserAssignedIdentityDetails">UserAssignedIdentityDetails
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity">ManagedClusterIdentity</a>)
</p>
<div>
<p>Information about the user assigned identity for the resource</p>
</div>
<table>
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
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.UserAssignedIdentityDetails_ARM">UserAssignedIdentityDetails_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIdentity_ARM">ManagedClusterIdentity_ARM</a>)
</p>
<div>
<p>Information about the user assigned identity for the resource</p>
</div>
<h3 id="containerservice.azure.com/v1api20230202preview.UserAssignedIdentity_ARM">UserAssignedIdentity_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentity_ARM">ManagedClusterPodIdentity_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_ARM">ManagedClusterProperties_ARM</a>)
</p>
<div>
<p>Details about a user assigned identity.</p>
</div>
<table>
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
<p>ClientId: The client ID of the user assigned identity.</p>
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
<p>ObjectId: The object ID of the user assigned identity.</p>
</td>
</tr>
<tr>
<td>
<code>resourceId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.UserAssignedIdentity_STATUS">UserAssignedIdentity_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAddonProfile_STATUS">ManagedClusterAddonProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIngressProfileWebAppRouting_STATUS">ManagedClusterIngressProfileWebAppRouting_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentity_STATUS">ManagedClusterPodIdentity_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedCluster_STATUS">ManagedCluster_STATUS</a>)
</p>
<div>
<p>Details about a user assigned identity.</p>
</div>
<table>
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
<p>ClientId: The client ID of the user assigned identity.</p>
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
<p>ObjectId: The object ID of the user assigned identity.</p>
</td>
</tr>
<tr>
<td>
<code>resourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ResourceId: The resource ID of the user assigned identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.UserAssignedIdentity_STATUS_ARM">UserAssignedIdentity_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAddonProfile_STATUS_ARM">ManagedClusterAddonProfile_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterIngressProfileWebAppRouting_STATUS_ARM">ManagedClusterIngressProfileWebAppRouting_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterPodIdentity_STATUS_ARM">ManagedClusterPodIdentity_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterProperties_STATUS_ARM">ManagedClusterProperties_STATUS_ARM</a>)
</p>
<div>
<p>Details about a user assigned identity.</p>
</div>
<table>
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
<p>ClientId: The client ID of the user assigned identity.</p>
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
<p>ObjectId: The object ID of the user assigned identity.</p>
</td>
</tr>
<tr>
<td>
<code>resourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ResourceId: The resource ID of the user assigned identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.WindowsGmsaProfile">WindowsGmsaProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWindowsProfile">ManagedClusterWindowsProfile</a>)
</p>
<div>
<p>Windows gMSA Profile in the managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dnsServer</code><br/>
<em>
string
</em>
</td>
<td>
<p>DnsServer: Specifies the DNS server for Windows gMSA.
Set it to empty if you have configured the DNS server in the vnet which is used to create the managed cluster.</p>
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
<p>Enabled: Specifies whether to enable Windows gMSA in the managed cluster.</p>
</td>
</tr>
<tr>
<td>
<code>rootDomainName</code><br/>
<em>
string
</em>
</td>
<td>
<p>RootDomainName: Specifies the root domain name for Windows gMSA.
Set it to empty if you have configured the DNS server in the vnet which is used to create the managed cluster.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.WindowsGmsaProfile_ARM">WindowsGmsaProfile_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWindowsProfile_ARM">ManagedClusterWindowsProfile_ARM</a>)
</p>
<div>
<p>Windows gMSA Profile in the managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dnsServer</code><br/>
<em>
string
</em>
</td>
<td>
<p>DnsServer: Specifies the DNS server for Windows gMSA.
Set it to empty if you have configured the DNS server in the vnet which is used to create the managed cluster.</p>
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
<p>Enabled: Specifies whether to enable Windows gMSA in the managed cluster.</p>
</td>
</tr>
<tr>
<td>
<code>rootDomainName</code><br/>
<em>
string
</em>
</td>
<td>
<p>RootDomainName: Specifies the root domain name for Windows gMSA.
Set it to empty if you have configured the DNS server in the vnet which is used to create the managed cluster.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.WindowsGmsaProfile_STATUS">WindowsGmsaProfile_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWindowsProfile_STATUS">ManagedClusterWindowsProfile_STATUS</a>)
</p>
<div>
<p>Windows gMSA Profile in the managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dnsServer</code><br/>
<em>
string
</em>
</td>
<td>
<p>DnsServer: Specifies the DNS server for Windows gMSA.
Set it to empty if you have configured the DNS server in the vnet which is used to create the managed cluster.</p>
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
<p>Enabled: Specifies whether to enable Windows gMSA in the managed cluster.</p>
</td>
</tr>
<tr>
<td>
<code>rootDomainName</code><br/>
<em>
string
</em>
</td>
<td>
<p>RootDomainName: Specifies the root domain name for Windows gMSA.
Set it to empty if you have configured the DNS server in the vnet which is used to create the managed cluster.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.WindowsGmsaProfile_STATUS_ARM">WindowsGmsaProfile_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterWindowsProfile_STATUS_ARM">ManagedClusterWindowsProfile_STATUS_ARM</a>)
</p>
<div>
<p>Windows gMSA Profile in the managed cluster.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dnsServer</code><br/>
<em>
string
</em>
</td>
<td>
<p>DnsServer: Specifies the DNS server for Windows gMSA.
Set it to empty if you have configured the DNS server in the vnet which is used to create the managed cluster.</p>
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
<p>Enabled: Specifies whether to enable Windows gMSA in the managed cluster.</p>
</td>
</tr>
<tr>
<td>
<code>rootDomainName</code><br/>
<em>
string
</em>
</td>
<td>
<p>RootDomainName: Specifies the root domain name for Windows gMSA.
Set it to empty if you have configured the DNS server in the vnet which is used to create the managed cluster.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.WorkloadRuntime">WorkloadRuntime
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_ARM">ManagedClusterAgentPoolProfileProperties_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_ARM">ManagedClusterAgentPoolProfile_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_Spec">ManagedClusters_AgentPool_Spec</a>)
</p>
<div>
<p>Determines the type of workload a node can run.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;KataMshvVmIsolation&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;OCIContainer&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;WasmWasi&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1api20230202preview.WorkloadRuntime_STATUS">WorkloadRuntime_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfileProperties_STATUS_ARM">ManagedClusterAgentPoolProfileProperties_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS">ManagedClusterAgentPoolProfile_STATUS</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusterAgentPoolProfile_STATUS_ARM">ManagedClusterAgentPoolProfile_STATUS_ARM</a>, <a href="#containerservice.azure.com/v1api20230202preview.ManagedClusters_AgentPool_STATUS">ManagedClusters_AgentPool_STATUS</a>)
</p>
<div>
<p>Determines the type of workload a node can run.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;KataMshvVmIsolation&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;OCIContainer&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;WasmWasi&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<hr/>
