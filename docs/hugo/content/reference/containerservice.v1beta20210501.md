---
---
<h2 id="containerservice.azure.com/v1beta20210501">containerservice.azure.com/v1beta20210501</h2>
<div>
<p>Package v1beta20210501 contains API Schema definitions for the containerservice v1beta20210501 API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="containerservice.azure.com/v1beta20210501.AgentPoolMode_Status">AgentPoolMode_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.AgentPool_Status">AgentPool_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileProperties_StatusARM">ManagedClusterAgentPoolProfileProperties_StatusARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_Status">ManagedClusterAgentPoolProfile_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_StatusARM">ManagedClusterAgentPoolProfile_StatusARM</a>)
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
<tbody><tr><td><p>&#34;System&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;User&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.AgentPoolType_Status">AgentPoolType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.AgentPool_Status">AgentPool_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileProperties_StatusARM">ManagedClusterAgentPoolProfileProperties_StatusARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_Status">ManagedClusterAgentPoolProfile_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_StatusARM">ManagedClusterAgentPoolProfile_StatusARM</a>)
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
<tbody><tr><td><p>&#34;AvailabilitySet&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VirtualMachineScaleSets&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.AgentPoolUpgradeSettings">AgentPoolUpgradeSettings
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClustersAgentPools_Spec">ManagedClustersAgentPools_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/AgentPoolUpgradeSettings">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/AgentPoolUpgradeSettings</a></p>
</div>
<table>
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
<h3 id="containerservice.azure.com/v1beta20210501.AgentPoolUpgradeSettingsARM">AgentPoolUpgradeSettingsARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileARM">ManagedClusterAgentPoolProfileARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesARM">ManagedClusterAgentPoolProfilePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/AgentPoolUpgradeSettings">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/AgentPoolUpgradeSettings</a></p>
</div>
<table>
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
<h3 id="containerservice.azure.com/v1beta20210501.AgentPoolUpgradeSettings_Status">AgentPoolUpgradeSettings_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.AgentPool_Status">AgentPool_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_Status">ManagedClusterAgentPoolProfile_Status</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.AgentPoolUpgradeSettings_StatusARM">AgentPoolUpgradeSettings_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileProperties_StatusARM">ManagedClusterAgentPoolProfileProperties_StatusARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_StatusARM">ManagedClusterAgentPoolProfile_StatusARM</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.AgentPool_Status">AgentPool_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClustersAgentPool">ManagedClustersAgentPool</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.GPUInstanceProfile_Status">
GPUInstanceProfile_Status
</a>
</em>
</td>
<td>
<p>GpuInstanceProfile: GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.</p>
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
<a href="#containerservice.azure.com/v1beta20210501.KubeletConfig_Status">
KubeletConfig_Status
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
<a href="#containerservice.azure.com/v1beta20210501.KubeletDiskType_Status">
KubeletDiskType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>linuxOSConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.LinuxOSConfig_Status">
LinuxOSConfig_Status
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
<a href="#containerservice.azure.com/v1beta20210501.AgentPoolMode_Status">
AgentPoolMode_Status
</a>
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
<p>Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.</p>
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
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}</p>
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
<p>OrchestratorVersion: As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes
version. The node pool version must have the same major version as the control plane. The node pool minor version must
be within two minor versions of the control plane version. The node pool version cannot be greater than the control
plane version. For more information see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool">upgrading a node
pool</a>.</p>
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
<a href="#containerservice.azure.com/v1beta20210501.OSDiskType_Status">
OSDiskType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osSKU</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.OSSKU_Status">
OSSKU_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.OSType_Status">
OSType_Status
</a>
</em>
</td>
<td>
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
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}</p>
</td>
</tr>
<tr>
<td>
<code>powerState</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.PowerState_Status">
PowerState_Status
</a>
</em>
</td>
<td>
<p>PowerState: Describes whether the Agent Pool is Running or Stopped</p>
</td>
</tr>
<tr>
<td>
<code>properties_type</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.AgentPoolType_Status">
AgentPoolType_Status
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
<code>scaleSetEvictionPolicy</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ScaleSetEvictionPolicy_Status">
ScaleSetEvictionPolicy_Status
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
<a href="#containerservice.azure.com/v1beta20210501.ScaleSetPriority_Status">
ScaleSetPriority_Status
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
<a href="#containerservice.azure.com/v1beta20210501.AgentPoolUpgradeSettings_Status">
AgentPoolUpgradeSettings_Status
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
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.AgentPool_StatusARM">AgentPool_StatusARM
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileProperties_StatusARM">
ManagedClusterAgentPoolProfileProperties_StatusARM
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
<h3 id="containerservice.azure.com/v1beta20210501.Componentsqit0Etschemasmanagedclusterpropertiespropertiesidentityprofileadditionalproperties">Componentsqit0Etschemasmanagedclusterpropertiespropertiesidentityprofileadditionalproperties
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusters_Spec">ManagedClusters_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/Componentsqit0etschemasmanagedclusterpropertiespropertiesidentityprofileadditionalproperties">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/Componentsqit0etschemasmanagedclusterpropertiespropertiesidentityprofileadditionalproperties</a></p>
</div>
<table>
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
<h3 id="containerservice.azure.com/v1beta20210501.Componentsqit0EtschemasmanagedclusterpropertiespropertiesidentityprofileadditionalpropertiesARM">Componentsqit0EtschemasmanagedclusterpropertiespropertiesidentityprofileadditionalpropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesARM">ManagedClusterPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/Componentsqit0etschemasmanagedclusterpropertiespropertiesidentityprofileadditionalproperties">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/Componentsqit0etschemasmanagedclusterpropertiespropertiesidentityprofileadditionalproperties</a></p>
</div>
<table>
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
<h3 id="containerservice.azure.com/v1beta20210501.ContainerServiceLinuxProfile">ContainerServiceLinuxProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusters_Spec">ManagedClusters_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ContainerServiceLinuxProfile">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ContainerServiceLinuxProfile</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceSshConfiguration">
ContainerServiceSshConfiguration
</a>
</em>
</td>
<td>
<p>Ssh: SSH configuration for Linux-based VMs running on Azure.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ContainerServiceLinuxProfileARM">ContainerServiceLinuxProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesARM">ManagedClusterPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ContainerServiceLinuxProfile">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ContainerServiceLinuxProfile</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceSshConfigurationARM">
ContainerServiceSshConfigurationARM
</a>
</em>
</td>
<td>
<p>Ssh: SSH configuration for Linux-based VMs running on Azure.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ContainerServiceLinuxProfile_Status">ContainerServiceLinuxProfile_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedCluster_Status">ManagedCluster_Status</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceSshConfiguration_Status">
ContainerServiceSshConfiguration_Status
</a>
</em>
</td>
<td>
<p>Ssh: The SSH configuration for Linux-based VMs running on Azure.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ContainerServiceLinuxProfile_StatusARM">ContainerServiceLinuxProfile_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterProperties_StatusARM">ManagedClusterProperties_StatusARM</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceSshConfiguration_StatusARM">
ContainerServiceSshConfiguration_StatusARM
</a>
</em>
</td>
<td>
<p>Ssh: The SSH configuration for Linux-based VMs running on Azure.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfile">ContainerServiceNetworkProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusters_Spec">ManagedClusters_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ContainerServiceNetworkProfile">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ContainerServiceNetworkProfile</a></p>
</div>
<table>
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
<code>loadBalancerProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile">
ManagedClusterLoadBalancerProfile
</a>
</em>
</td>
<td>
<p>LoadBalancerProfile: Profile of the managed cluster load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerSku</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileLoadBalancerSku">
ContainerServiceNetworkProfileLoadBalancerSku
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
<code>networkMode</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileNetworkMode">
ContainerServiceNetworkProfileNetworkMode
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
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileNetworkPlugin">
ContainerServiceNetworkProfileNetworkPlugin
</a>
</em>
</td>
<td>
<p>NetworkPlugin: Network plugin used for building the Kubernetes network.</p>
</td>
</tr>
<tr>
<td>
<code>networkPolicy</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileNetworkPolicy">
ContainerServiceNetworkProfileNetworkPolicy
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
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileOutboundType">
ContainerServiceNetworkProfileOutboundType
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
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileARM">ContainerServiceNetworkProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesARM">ManagedClusterPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ContainerServiceNetworkProfile">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ContainerServiceNetworkProfile</a></p>
</div>
<table>
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
<code>loadBalancerProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfileARM">
ManagedClusterLoadBalancerProfileARM
</a>
</em>
</td>
<td>
<p>LoadBalancerProfile: Profile of the managed cluster load balancer.</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerSku</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileLoadBalancerSku">
ContainerServiceNetworkProfileLoadBalancerSku
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
<code>networkMode</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileNetworkMode">
ContainerServiceNetworkProfileNetworkMode
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
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileNetworkPlugin">
ContainerServiceNetworkProfileNetworkPlugin
</a>
</em>
</td>
<td>
<p>NetworkPlugin: Network plugin used for building the Kubernetes network.</p>
</td>
</tr>
<tr>
<td>
<code>networkPolicy</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileNetworkPolicy">
ContainerServiceNetworkProfileNetworkPolicy
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
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileOutboundType">
ContainerServiceNetworkProfileOutboundType
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
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileLoadBalancerSku">ContainerServiceNetworkProfileLoadBalancerSku
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfile">ContainerServiceNetworkProfile</a>, <a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileARM">ContainerServiceNetworkProfileARM</a>)
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
<tbody><tr><td><p>&#34;basic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileNetworkMode">ContainerServiceNetworkProfileNetworkMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfile">ContainerServiceNetworkProfile</a>, <a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileARM">ContainerServiceNetworkProfileARM</a>)
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
<tbody><tr><td><p>&#34;bridge&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;transparent&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileNetworkPlugin">ContainerServiceNetworkProfileNetworkPlugin
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfile">ContainerServiceNetworkProfile</a>, <a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileARM">ContainerServiceNetworkProfileARM</a>)
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
<tbody><tr><td><p>&#34;azure&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;kubenet&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileNetworkPolicy">ContainerServiceNetworkProfileNetworkPolicy
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfile">ContainerServiceNetworkProfile</a>, <a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileARM">ContainerServiceNetworkProfileARM</a>)
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
<tbody><tr><td><p>&#34;azure&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;calico&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileOutboundType">ContainerServiceNetworkProfileOutboundType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfile">ContainerServiceNetworkProfile</a>, <a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileARM">ContainerServiceNetworkProfileARM</a>)
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
</tr><tr><td><p>&#34;userDefinedRouting&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileStatusLoadBalancerSku">ContainerServiceNetworkProfileStatusLoadBalancerSku
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfile_Status">ContainerServiceNetworkProfile_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfile_StatusARM">ContainerServiceNetworkProfile_StatusARM</a>)
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
<tbody><tr><td><p>&#34;basic&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileStatusNetworkMode">ContainerServiceNetworkProfileStatusNetworkMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfile_Status">ContainerServiceNetworkProfile_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfile_StatusARM">ContainerServiceNetworkProfile_StatusARM</a>)
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
<tbody><tr><td><p>&#34;bridge&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;transparent&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileStatusNetworkPlugin">ContainerServiceNetworkProfileStatusNetworkPlugin
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfile_Status">ContainerServiceNetworkProfile_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfile_StatusARM">ContainerServiceNetworkProfile_StatusARM</a>)
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
<tbody><tr><td><p>&#34;azure&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;kubenet&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileStatusNetworkPolicy">ContainerServiceNetworkProfileStatusNetworkPolicy
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfile_Status">ContainerServiceNetworkProfile_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfile_StatusARM">ContainerServiceNetworkProfile_StatusARM</a>)
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
<tbody><tr><td><p>&#34;azure&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;calico&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileStatusOutboundType">ContainerServiceNetworkProfileStatusOutboundType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfile_Status">ContainerServiceNetworkProfile_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfile_StatusARM">ContainerServiceNetworkProfile_StatusARM</a>)
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
</tr><tr><td><p>&#34;userDefinedRouting&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfile_Status">ContainerServiceNetworkProfile_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedCluster_Status">ManagedCluster_Status</a>)
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
<code>loadBalancerProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_Status">
ManagedClusterLoadBalancerProfile_Status
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
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileStatusLoadBalancerSku">
ContainerServiceNetworkProfileStatusLoadBalancerSku
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
<code>networkMode</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileStatusNetworkMode">
ContainerServiceNetworkProfileStatusNetworkMode
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
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileStatusNetworkPlugin">
ContainerServiceNetworkProfileStatusNetworkPlugin
</a>
</em>
</td>
<td>
<p>NetworkPlugin: Network plugin used for building the Kubernetes network.</p>
</td>
</tr>
<tr>
<td>
<code>networkPolicy</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileStatusNetworkPolicy">
ContainerServiceNetworkProfileStatusNetworkPolicy
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
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileStatusOutboundType">
ContainerServiceNetworkProfileStatusOutboundType
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
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfile_StatusARM">ContainerServiceNetworkProfile_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterProperties_StatusARM">ManagedClusterProperties_StatusARM</a>)
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
<code>loadBalancerProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_StatusARM">
ManagedClusterLoadBalancerProfile_StatusARM
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
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileStatusLoadBalancerSku">
ContainerServiceNetworkProfileStatusLoadBalancerSku
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
<code>networkMode</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileStatusNetworkMode">
ContainerServiceNetworkProfileStatusNetworkMode
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
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileStatusNetworkPlugin">
ContainerServiceNetworkProfileStatusNetworkPlugin
</a>
</em>
</td>
<td>
<p>NetworkPlugin: Network plugin used for building the Kubernetes network.</p>
</td>
</tr>
<tr>
<td>
<code>networkPolicy</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileStatusNetworkPolicy">
ContainerServiceNetworkProfileStatusNetworkPolicy
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
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileStatusOutboundType">
ContainerServiceNetworkProfileStatusOutboundType
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
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ContainerServiceSshConfiguration">ContainerServiceSshConfiguration
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ContainerServiceLinuxProfile">ContainerServiceLinuxProfile</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ContainerServiceSshConfiguration">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ContainerServiceSshConfiguration</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceSshPublicKey">
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
<h3 id="containerservice.azure.com/v1beta20210501.ContainerServiceSshConfigurationARM">ContainerServiceSshConfigurationARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ContainerServiceLinuxProfileARM">ContainerServiceLinuxProfileARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ContainerServiceSshConfiguration">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ContainerServiceSshConfiguration</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceSshPublicKeyARM">
[]ContainerServiceSshPublicKeyARM
</a>
</em>
</td>
<td>
<p>PublicKeys: The list of SSH public keys used to authenticate with Linux-based VMs. A maximum of 1 key may be specified.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ContainerServiceSshConfiguration_Status">ContainerServiceSshConfiguration_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ContainerServiceLinuxProfile_Status">ContainerServiceLinuxProfile_Status</a>)
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
<code>publicKeys</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceSshPublicKey_Status">
[]ContainerServiceSshPublicKey_Status
</a>
</em>
</td>
<td>
<p>PublicKeys: The list of SSH public keys used to authenticate with Linux-based VMs. A maximum of 1 key may be specified.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ContainerServiceSshConfiguration_StatusARM">ContainerServiceSshConfiguration_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ContainerServiceLinuxProfile_StatusARM">ContainerServiceLinuxProfile_StatusARM</a>)
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
<code>publicKeys</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceSshPublicKey_StatusARM">
[]ContainerServiceSshPublicKey_StatusARM
</a>
</em>
</td>
<td>
<p>PublicKeys: The list of SSH public keys used to authenticate with Linux-based VMs. A maximum of 1 key may be specified.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ContainerServiceSshPublicKey">ContainerServiceSshPublicKey
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ContainerServiceSshConfiguration">ContainerServiceSshConfiguration</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ContainerServiceSshPublicKey">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ContainerServiceSshPublicKey</a></p>
</div>
<table>
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
<h3 id="containerservice.azure.com/v1beta20210501.ContainerServiceSshPublicKeyARM">ContainerServiceSshPublicKeyARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ContainerServiceSshConfigurationARM">ContainerServiceSshConfigurationARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ContainerServiceSshPublicKey">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ContainerServiceSshPublicKey</a></p>
</div>
<table>
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
<h3 id="containerservice.azure.com/v1beta20210501.ContainerServiceSshPublicKey_Status">ContainerServiceSshPublicKey_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ContainerServiceSshConfiguration_Status">ContainerServiceSshConfiguration_Status</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.ContainerServiceSshPublicKey_StatusARM">ContainerServiceSshPublicKey_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ContainerServiceSshConfiguration_StatusARM">ContainerServiceSshConfiguration_StatusARM</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.ExtendedLocation">ExtendedLocation
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusters_Spec">ManagedClusters_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ExtendedLocation">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ExtendedLocation</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.ExtendedLocationType">
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
<h3 id="containerservice.azure.com/v1beta20210501.ExtendedLocationARM">ExtendedLocationARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusters_SpecARM">ManagedClusters_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ExtendedLocation">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ExtendedLocation</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.ExtendedLocationType">
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
<h3 id="containerservice.azure.com/v1beta20210501.ExtendedLocationType">ExtendedLocationType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ExtendedLocation">ExtendedLocation</a>, <a href="#containerservice.azure.com/v1beta20210501.ExtendedLocationARM">ExtendedLocationARM</a>)
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
<tbody><tr><td><p>&#34;EdgeZone&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ExtendedLocationType_Status">ExtendedLocationType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ExtendedLocation_Status">ExtendedLocation_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ExtendedLocation_StatusARM">ExtendedLocation_StatusARM</a>)
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
<tbody><tr><td><p>&#34;EdgeZone&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ExtendedLocation_Status">ExtendedLocation_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedCluster_Status">ManagedCluster_Status</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.ExtendedLocationType_Status">
ExtendedLocationType_Status
</a>
</em>
</td>
<td>
<p>Type: The type of the extended location.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ExtendedLocation_StatusARM">ExtendedLocation_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedCluster_StatusARM">ManagedCluster_StatusARM</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.ExtendedLocationType_Status">
ExtendedLocationType_Status
</a>
</em>
</td>
<td>
<p>Type: The type of the extended location.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.GPUInstanceProfile_Status">GPUInstanceProfile_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.AgentPool_Status">AgentPool_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileProperties_StatusARM">ManagedClusterAgentPoolProfileProperties_StatusARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_Status">ManagedClusterAgentPoolProfile_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_StatusARM">ManagedClusterAgentPoolProfile_StatusARM</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.KubeletConfig">KubeletConfig
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClustersAgentPools_Spec">ManagedClustersAgentPools_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/KubeletConfig">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/KubeletConfig</a></p>
</div>
<table>
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
<h3 id="containerservice.azure.com/v1beta20210501.KubeletConfigARM">KubeletConfigARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileARM">ManagedClusterAgentPoolProfileARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesARM">ManagedClusterAgentPoolProfilePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/KubeletConfig">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/KubeletConfig</a></p>
</div>
<table>
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
<h3 id="containerservice.azure.com/v1beta20210501.KubeletConfig_Status">KubeletConfig_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.AgentPool_Status">AgentPool_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_Status">ManagedClusterAgentPoolProfile_Status</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.KubeletConfig_StatusARM">KubeletConfig_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileProperties_StatusARM">ManagedClusterAgentPoolProfileProperties_StatusARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_StatusARM">ManagedClusterAgentPoolProfile_StatusARM</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.KubeletDiskType_Status">KubeletDiskType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.AgentPool_Status">AgentPool_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileProperties_StatusARM">ManagedClusterAgentPoolProfileProperties_StatusARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_Status">ManagedClusterAgentPoolProfile_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_StatusARM">ManagedClusterAgentPoolProfile_StatusARM</a>)
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
<tbody><tr><td><p>&#34;OS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Temporary&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.LinuxOSConfig">LinuxOSConfig
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClustersAgentPools_Spec">ManagedClustersAgentPools_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/LinuxOSConfig">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/LinuxOSConfig</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.SysctlConfig">
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
<h3 id="containerservice.azure.com/v1beta20210501.LinuxOSConfigARM">LinuxOSConfigARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileARM">ManagedClusterAgentPoolProfileARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesARM">ManagedClusterAgentPoolProfilePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/LinuxOSConfig">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/LinuxOSConfig</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.SysctlConfigARM">
SysctlConfigARM
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
<h3 id="containerservice.azure.com/v1beta20210501.LinuxOSConfig_Status">LinuxOSConfig_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.AgentPool_Status">AgentPool_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_Status">ManagedClusterAgentPoolProfile_Status</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.SysctlConfig_Status">
SysctlConfig_Status
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
<h3 id="containerservice.azure.com/v1beta20210501.LinuxOSConfig_StatusARM">LinuxOSConfig_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileProperties_StatusARM">ManagedClusterAgentPoolProfileProperties_StatusARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_StatusARM">ManagedClusterAgentPoolProfile_StatusARM</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.SysctlConfig_StatusARM">
SysctlConfig_StatusARM
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedCluster">ManagedCluster
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/resourceDefinitions/managedClusters">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/resourceDefinitions/managedClusters</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusters_Spec">
ManagedClusters_Spec
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAADProfile">
ManagedClusterAADProfile
</a>
</em>
</td>
<td>
<p>AadProfile: For more details see <a href="https://docs.microsoft.com/azure/aks/managed-aad">managed AAD on AKS</a>.</p>
</td>
</tr>
<tr>
<td>
<code>addonProfiles</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAddonProfile">
map[string]./api/containerservice/v1beta20210501.ManagedClusterAddonProfile
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile">
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAPIServerAccessProfile">
ManagedClusterAPIServerAccessProfile
</a>
</em>
</td>
<td>
<p>ApiServerAccessProfile: Access profile for managed cluster API server.</p>
</td>
</tr>
<tr>
<td>
<code>autoScalerProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesAutoScalerProfile">
ManagedClusterPropertiesAutoScalerProfile
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAutoUpgradeProfile">
ManagedClusterAutoUpgradeProfile
</a>
</em>
</td>
<td>
<p>AutoUpgradeProfile: Auto upgrade profile for a managed cluster.</p>
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
<code>diskEncryptionSetIDReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>DiskEncryptionSetIDReference: This is of the form:
&lsquo;/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{encryptionSetName}&rsquo;</p>
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
<code>enablePodSecurityPolicy</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePodSecurityPolicy: (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set
for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.</p>
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
<a href="#containerservice.azure.com/v1beta20210501.ExtendedLocation">
ExtendedLocation
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The complex type of the extended location.</p>
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
<code>httpProxyConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterHTTPProxyConfig">
ManagedClusterHTTPProxyConfig
</a>
</em>
</td>
<td>
<p>HttpProxyConfig: Cluster HTTP proxy configuration.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterIdentity">
ManagedClusterIdentity
</a>
</em>
</td>
<td>
<p>Identity: Identity for the managed cluster.</p>
</td>
</tr>
<tr>
<td>
<code>identityProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.Componentsqit0Etschemasmanagedclusterpropertiespropertiesidentityprofileadditionalproperties">
map[string]./api/containerservice/v1beta20210501.Componentsqit0Etschemasmanagedclusterpropertiespropertiesidentityprofileadditionalproperties
</a>
</em>
</td>
<td>
<p>IdentityProfile: Identities associated with the cluster.</p>
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
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceLinuxProfile">
ContainerServiceLinuxProfile
</a>
</em>
</td>
<td>
<p>LinuxProfile: Profile for Linux VMs in the container service cluster.</p>
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
<code>networkProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfile">
ContainerServiceNetworkProfile
</a>
</em>
</td>
<td>
<p>NetworkProfile: Profile of network configuration.</p>
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProfile">
ManagedClusterPodIdentityProfile
</a>
</em>
</td>
<td>
<p>PodIdentityProfile: See <a href="https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity">use AAD pod identity</a> for more
details on pod identity integration.</p>
</td>
</tr>
<tr>
<td>
<code>privateLinkResources</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.PrivateLinkResource">
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
<code>servicePrincipalProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterServicePrincipalProfile">
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterSKU">
ManagedClusterSKU
</a>
</em>
</td>
<td>
<p>Sku: The SKU of a Managed Cluster.</p>
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
<code>windowsProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterWindowsProfile">
ManagedClusterWindowsProfile
</a>
</em>
</td>
<td>
<p>WindowsProfile: Profile for Windows VMs in the managed cluster.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedCluster_Status">
ManagedCluster_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAADProfile">ManagedClusterAADProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusters_Spec">ManagedClusters_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterAADProfile">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterAADProfile</a></p>
</div>
<table>
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
<p>ClientAppID: The client AAD application ID.</p>
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
<p>ServerAppID: The server AAD application ID.</p>
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
<p>ServerAppSecret: The server AAD application secret.</p>
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAADProfileARM">ManagedClusterAADProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesARM">ManagedClusterPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterAADProfile">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterAADProfile</a></p>
</div>
<table>
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
<p>ClientAppID: The client AAD application ID.</p>
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
<p>ServerAppID: The server AAD application ID.</p>
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
<p>ServerAppSecret: The server AAD application secret.</p>
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAADProfile_Status">ManagedClusterAADProfile_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedCluster_Status">ManagedCluster_Status</a>)
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
<p>ClientAppID: The client AAD application ID.</p>
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
<p>ServerAppID: The server AAD application ID.</p>
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
<p>ServerAppSecret: The server AAD application secret.</p>
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAADProfile_StatusARM">ManagedClusterAADProfile_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterProperties_StatusARM">ManagedClusterProperties_StatusARM</a>)
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
<p>ClientAppID: The client AAD application ID.</p>
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
<p>ServerAppID: The server AAD application ID.</p>
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
<p>ServerAppSecret: The server AAD application secret.</p>
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAPIServerAccessProfile">ManagedClusterAPIServerAccessProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusters_Spec">ManagedClusters_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterAPIServerAccessProfile">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterAPIServerAccessProfile</a></p>
</div>
<table>
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
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAPIServerAccessProfileARM">ManagedClusterAPIServerAccessProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesARM">ManagedClusterPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterAPIServerAccessProfile">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterAPIServerAccessProfile</a></p>
</div>
<table>
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
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAPIServerAccessProfile_Status">ManagedClusterAPIServerAccessProfile_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedCluster_Status">ManagedCluster_Status</a>)
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
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAPIServerAccessProfile_StatusARM">ManagedClusterAPIServerAccessProfile_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterProperties_StatusARM">ManagedClusterProperties_StatusARM</a>)
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
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAddonProfile">ManagedClusterAddonProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusters_Spec">ManagedClusters_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterAddonProfile">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterAddonProfile</a></p>
</div>
<table>
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAddonProfileARM">ManagedClusterAddonProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesARM">ManagedClusterPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterAddonProfile">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterAddonProfile</a></p>
</div>
<table>
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusters_Spec">ManagedClusters_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterAgentPoolProfile">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterAgentPoolProfile</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileGpuInstanceProfile">
ManagedClusterAgentPoolProfileGpuInstanceProfile
</a>
</em>
</td>
<td>
<p>GpuInstanceProfile: GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.KubeletConfig">
KubeletConfig
</a>
</em>
</td>
<td>
<p>KubeletConfig: See <a href="https://docs.microsoft.com/azure/aks/custom-node-configuration">AKS custom node configuration</a> for
more details.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileKubeletDiskType">
ManagedClusterAgentPoolProfileKubeletDiskType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>linuxOSConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.LinuxOSConfig">
LinuxOSConfig
</a>
</em>
</td>
<td>
<p>LinuxOSConfig: See <a href="https://docs.microsoft.com/azure/aks/custom-node-configuration">AKS custom node configuration</a> for
more details.</p>
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileMode">
ManagedClusterAgentPoolProfileMode
</a>
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
<p>Name: Windows agent pool names must be 6 characters or less.</p>
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
<code>nodePublicIPPrefixIDReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>NodePublicIPPrefixIDReference: This is of the form:
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}</p>
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
<p>OrchestratorVersion: As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes
version. The node pool version must have the same major version as the control plane. The node pool minor version must
be within two minor versions of the control plane version. The node pool version cannot be greater than the control
plane version. For more information see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool">upgrading a node
pool</a>.</p>
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
<p>OsDiskSizeGB: OS Disk Size in GB to be used to specify the disk size for every machine in the master/agent pool. If you
specify 0, it will apply the default osDisk size according to the vmSize specified.</p>
</td>
</tr>
<tr>
<td>
<code>osDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileOsDiskType">
ManagedClusterAgentPoolProfileOsDiskType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osSKU</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileOsSKU">
ManagedClusterAgentPoolProfileOsSKU
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileOsType">
ManagedClusterAgentPoolProfileOsType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>podSubnetIDReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>PodSubnetIDReference: If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more
details). This is of the form:
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}</p>
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
<code>scaleSetEvictionPolicy</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileScaleSetEvictionPolicy">
ManagedClusterAgentPoolProfileScaleSetEvictionPolicy
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileScaleSetPriority">
ManagedClusterAgentPoolProfileScaleSetPriority
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileType">
ManagedClusterAgentPoolProfileType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>upgradeSettings</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.AgentPoolUpgradeSettings">
AgentPoolUpgradeSettings
</a>
</em>
</td>
<td>
<p>UpgradeSettings: Settings for upgrading an agentpool</p>
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
<code>vnetSubnetIDReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>VnetSubnetIDReference: If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is
specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form:
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileARM">ManagedClusterAgentPoolProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesARM">ManagedClusterPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterAgentPoolProfile">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterAgentPoolProfile</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileGpuInstanceProfile">
ManagedClusterAgentPoolProfileGpuInstanceProfile
</a>
</em>
</td>
<td>
<p>GpuInstanceProfile: GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.KubeletConfigARM">
KubeletConfigARM
</a>
</em>
</td>
<td>
<p>KubeletConfig: See <a href="https://docs.microsoft.com/azure/aks/custom-node-configuration">AKS custom node configuration</a> for
more details.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileKubeletDiskType">
ManagedClusterAgentPoolProfileKubeletDiskType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>linuxOSConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.LinuxOSConfigARM">
LinuxOSConfigARM
</a>
</em>
</td>
<td>
<p>LinuxOSConfig: See <a href="https://docs.microsoft.com/azure/aks/custom-node-configuration">AKS custom node configuration</a> for
more details.</p>
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileMode">
ManagedClusterAgentPoolProfileMode
</a>
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
<p>Name: Windows agent pool names must be 6 characters or less.</p>
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
<p>OrchestratorVersion: As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes
version. The node pool version must have the same major version as the control plane. The node pool minor version must
be within two minor versions of the control plane version. The node pool version cannot be greater than the control
plane version. For more information see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool">upgrading a node
pool</a>.</p>
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
<p>OsDiskSizeGB: OS Disk Size in GB to be used to specify the disk size for every machine in the master/agent pool. If you
specify 0, it will apply the default osDisk size according to the vmSize specified.</p>
</td>
</tr>
<tr>
<td>
<code>osDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileOsDiskType">
ManagedClusterAgentPoolProfileOsDiskType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osSKU</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileOsSKU">
ManagedClusterAgentPoolProfileOsSKU
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileOsType">
ManagedClusterAgentPoolProfileOsType
</a>
</em>
</td>
<td>
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
<code>scaleSetEvictionPolicy</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileScaleSetEvictionPolicy">
ManagedClusterAgentPoolProfileScaleSetEvictionPolicy
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileScaleSetPriority">
ManagedClusterAgentPoolProfileScaleSetPriority
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileType">
ManagedClusterAgentPoolProfileType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>upgradeSettings</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.AgentPoolUpgradeSettingsARM">
AgentPoolUpgradeSettingsARM
</a>
</em>
</td>
<td>
<p>UpgradeSettings: Settings for upgrading an agentpool</p>
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
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileGpuInstanceProfile">ManagedClusterAgentPoolProfileGpuInstanceProfile
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileARM">ManagedClusterAgentPoolProfileARM</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileKubeletDiskType">ManagedClusterAgentPoolProfileKubeletDiskType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileARM">ManagedClusterAgentPoolProfileARM</a>)
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
<tbody><tr><td><p>&#34;OS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Temporary&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileMode">ManagedClusterAgentPoolProfileMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileARM">ManagedClusterAgentPoolProfileARM</a>)
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
<tbody><tr><td><p>&#34;System&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;User&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileOsDiskType">ManagedClusterAgentPoolProfileOsDiskType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileARM">ManagedClusterAgentPoolProfileARM</a>)
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
<tbody><tr><td><p>&#34;Ephemeral&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Managed&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileOsSKU">ManagedClusterAgentPoolProfileOsSKU
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileARM">ManagedClusterAgentPoolProfileARM</a>)
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
<tbody><tr><td><p>&#34;CBLMariner&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Ubuntu&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileOsType">ManagedClusterAgentPoolProfileOsType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileARM">ManagedClusterAgentPoolProfileARM</a>)
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
<tbody><tr><td><p>&#34;Linux&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Windows&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesARM">ManagedClusterAgentPoolProfilePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClustersAgentPools_SpecARM">ManagedClustersAgentPools_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterAgentPoolProfileProperties">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterAgentPoolProfileProperties</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesGpuInstanceProfile">
ManagedClusterAgentPoolProfilePropertiesGpuInstanceProfile
</a>
</em>
</td>
<td>
<p>GpuInstanceProfile: GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.KubeletConfigARM">
KubeletConfigARM
</a>
</em>
</td>
<td>
<p>KubeletConfig: See <a href="https://docs.microsoft.com/azure/aks/custom-node-configuration">AKS custom node configuration</a> for
more details.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesKubeletDiskType">
ManagedClusterAgentPoolProfilePropertiesKubeletDiskType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>linuxOSConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.LinuxOSConfigARM">
LinuxOSConfigARM
</a>
</em>
</td>
<td>
<p>LinuxOSConfig: See <a href="https://docs.microsoft.com/azure/aks/custom-node-configuration">AKS custom node configuration</a> for
more details.</p>
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesMode">
ManagedClusterAgentPoolProfilePropertiesMode
</a>
</em>
</td>
<td>
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
<p>OrchestratorVersion: As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes
version. The node pool version must have the same major version as the control plane. The node pool minor version must
be within two minor versions of the control plane version. The node pool version cannot be greater than the control
plane version. For more information see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool">upgrading a node
pool</a>.</p>
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
<p>OsDiskSizeGB: OS Disk Size in GB to be used to specify the disk size for every machine in the master/agent pool. If you
specify 0, it will apply the default osDisk size according to the vmSize specified.</p>
</td>
</tr>
<tr>
<td>
<code>osDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesOsDiskType">
ManagedClusterAgentPoolProfilePropertiesOsDiskType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osSKU</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesOsSKU">
ManagedClusterAgentPoolProfilePropertiesOsSKU
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesOsType">
ManagedClusterAgentPoolProfilePropertiesOsType
</a>
</em>
</td>
<td>
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
<code>scaleSetEvictionPolicy</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesScaleSetEvictionPolicy">
ManagedClusterAgentPoolProfilePropertiesScaleSetEvictionPolicy
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesScaleSetPriority">
ManagedClusterAgentPoolProfilePropertiesScaleSetPriority
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesType">
ManagedClusterAgentPoolProfilePropertiesType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>upgradeSettings</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.AgentPoolUpgradeSettingsARM">
AgentPoolUpgradeSettingsARM
</a>
</em>
</td>
<td>
<p>UpgradeSettings: Settings for upgrading an agentpool</p>
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
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesGpuInstanceProfile">ManagedClusterAgentPoolProfilePropertiesGpuInstanceProfile
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesARM">ManagedClusterAgentPoolProfilePropertiesARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClustersAgentPools_Spec">ManagedClustersAgentPools_Spec</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesKubeletDiskType">ManagedClusterAgentPoolProfilePropertiesKubeletDiskType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesARM">ManagedClusterAgentPoolProfilePropertiesARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClustersAgentPools_Spec">ManagedClustersAgentPools_Spec</a>)
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
<tbody><tr><td><p>&#34;OS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Temporary&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesMode">ManagedClusterAgentPoolProfilePropertiesMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesARM">ManagedClusterAgentPoolProfilePropertiesARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClustersAgentPools_Spec">ManagedClustersAgentPools_Spec</a>)
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
<tbody><tr><td><p>&#34;System&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;User&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesOsDiskType">ManagedClusterAgentPoolProfilePropertiesOsDiskType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesARM">ManagedClusterAgentPoolProfilePropertiesARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClustersAgentPools_Spec">ManagedClustersAgentPools_Spec</a>)
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
<tbody><tr><td><p>&#34;Ephemeral&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Managed&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesOsSKU">ManagedClusterAgentPoolProfilePropertiesOsSKU
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesARM">ManagedClusterAgentPoolProfilePropertiesARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClustersAgentPools_Spec">ManagedClustersAgentPools_Spec</a>)
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
<tbody><tr><td><p>&#34;CBLMariner&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Ubuntu&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesOsType">ManagedClusterAgentPoolProfilePropertiesOsType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesARM">ManagedClusterAgentPoolProfilePropertiesARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClustersAgentPools_Spec">ManagedClustersAgentPools_Spec</a>)
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
<tbody><tr><td><p>&#34;Linux&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Windows&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesScaleSetEvictionPolicy">ManagedClusterAgentPoolProfilePropertiesScaleSetEvictionPolicy
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesARM">ManagedClusterAgentPoolProfilePropertiesARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClustersAgentPools_Spec">ManagedClustersAgentPools_Spec</a>)
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
<tbody><tr><td><p>&#34;Deallocate&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Delete&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesScaleSetPriority">ManagedClusterAgentPoolProfilePropertiesScaleSetPriority
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesARM">ManagedClusterAgentPoolProfilePropertiesARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClustersAgentPools_Spec">ManagedClustersAgentPools_Spec</a>)
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
<tbody><tr><td><p>&#34;Regular&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Spot&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesType">ManagedClusterAgentPoolProfilePropertiesType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesARM">ManagedClusterAgentPoolProfilePropertiesARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClustersAgentPools_Spec">ManagedClustersAgentPools_Spec</a>)
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
<tbody><tr><td><p>&#34;AvailabilitySet&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VirtualMachineScaleSets&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileProperties_StatusARM">ManagedClusterAgentPoolProfileProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.AgentPool_StatusARM">AgentPool_StatusARM</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.GPUInstanceProfile_Status">
GPUInstanceProfile_Status
</a>
</em>
</td>
<td>
<p>GpuInstanceProfile: GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.KubeletConfig_StatusARM">
KubeletConfig_StatusARM
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
<a href="#containerservice.azure.com/v1beta20210501.KubeletDiskType_Status">
KubeletDiskType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>linuxOSConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.LinuxOSConfig_StatusARM">
LinuxOSConfig_StatusARM
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
<a href="#containerservice.azure.com/v1beta20210501.AgentPoolMode_Status">
AgentPoolMode_Status
</a>
</em>
</td>
<td>
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
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}</p>
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
<p>OrchestratorVersion: As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes
version. The node pool version must have the same major version as the control plane. The node pool minor version must
be within two minor versions of the control plane version. The node pool version cannot be greater than the control
plane version. For more information see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool">upgrading a node
pool</a>.</p>
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
<a href="#containerservice.azure.com/v1beta20210501.OSDiskType_Status">
OSDiskType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osSKU</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.OSSKU_Status">
OSSKU_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.OSType_Status">
OSType_Status
</a>
</em>
</td>
<td>
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
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}</p>
</td>
</tr>
<tr>
<td>
<code>powerState</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.PowerState_StatusARM">
PowerState_StatusARM
</a>
</em>
</td>
<td>
<p>PowerState: Describes whether the Agent Pool is Running or Stopped</p>
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
<code>scaleSetEvictionPolicy</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ScaleSetEvictionPolicy_Status">
ScaleSetEvictionPolicy_Status
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
<a href="#containerservice.azure.com/v1beta20210501.ScaleSetPriority_Status">
ScaleSetPriority_Status
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
<a href="#containerservice.azure.com/v1beta20210501.AgentPoolType_Status">
AgentPoolType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>upgradeSettings</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.AgentPoolUpgradeSettings_StatusARM">
AgentPoolUpgradeSettings_StatusARM
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
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileScaleSetEvictionPolicy">ManagedClusterAgentPoolProfileScaleSetEvictionPolicy
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileARM">ManagedClusterAgentPoolProfileARM</a>)
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
<tbody><tr><td><p>&#34;Deallocate&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Delete&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileScaleSetPriority">ManagedClusterAgentPoolProfileScaleSetPriority
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileARM">ManagedClusterAgentPoolProfileARM</a>)
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
<tbody><tr><td><p>&#34;Regular&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Spot&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileType">ManagedClusterAgentPoolProfileType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile">ManagedClusterAgentPoolProfile</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileARM">ManagedClusterAgentPoolProfileARM</a>)
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
<tbody><tr><td><p>&#34;AvailabilitySet&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;VirtualMachineScaleSets&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_Status">ManagedClusterAgentPoolProfile_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedCluster_Status">ManagedCluster_Status</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.GPUInstanceProfile_Status">
GPUInstanceProfile_Status
</a>
</em>
</td>
<td>
<p>GpuInstanceProfile: GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.KubeletConfig_Status">
KubeletConfig_Status
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
<a href="#containerservice.azure.com/v1beta20210501.KubeletDiskType_Status">
KubeletDiskType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>linuxOSConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.LinuxOSConfig_Status">
LinuxOSConfig_Status
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
<a href="#containerservice.azure.com/v1beta20210501.AgentPoolMode_Status">
AgentPoolMode_Status
</a>
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
<p>Name: Windows agent pool names must be 6 characters or less.</p>
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
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}</p>
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
<p>OrchestratorVersion: As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes
version. The node pool version must have the same major version as the control plane. The node pool minor version must
be within two minor versions of the control plane version. The node pool version cannot be greater than the control
plane version. For more information see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool">upgrading a node
pool</a>.</p>
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
<a href="#containerservice.azure.com/v1beta20210501.OSDiskType_Status">
OSDiskType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osSKU</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.OSSKU_Status">
OSSKU_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.OSType_Status">
OSType_Status
</a>
</em>
</td>
<td>
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
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}</p>
</td>
</tr>
<tr>
<td>
<code>powerState</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.PowerState_Status">
PowerState_Status
</a>
</em>
</td>
<td>
<p>PowerState: Describes whether the Agent Pool is Running or Stopped</p>
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
<code>scaleSetEvictionPolicy</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ScaleSetEvictionPolicy_Status">
ScaleSetEvictionPolicy_Status
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
<a href="#containerservice.azure.com/v1beta20210501.ScaleSetPriority_Status">
ScaleSetPriority_Status
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
<a href="#containerservice.azure.com/v1beta20210501.AgentPoolType_Status">
AgentPoolType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>upgradeSettings</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.AgentPoolUpgradeSettings_Status">
AgentPoolUpgradeSettings_Status
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
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_StatusARM">ManagedClusterAgentPoolProfile_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterProperties_StatusARM">ManagedClusterProperties_StatusARM</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.GPUInstanceProfile_Status">
GPUInstanceProfile_Status
</a>
</em>
</td>
<td>
<p>GpuInstanceProfile: GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.KubeletConfig_StatusARM">
KubeletConfig_StatusARM
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
<a href="#containerservice.azure.com/v1beta20210501.KubeletDiskType_Status">
KubeletDiskType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>linuxOSConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.LinuxOSConfig_StatusARM">
LinuxOSConfig_StatusARM
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
<a href="#containerservice.azure.com/v1beta20210501.AgentPoolMode_Status">
AgentPoolMode_Status
</a>
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
<p>Name: Windows agent pool names must be 6 characters or less.</p>
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
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}</p>
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
<p>OrchestratorVersion: As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes
version. The node pool version must have the same major version as the control plane. The node pool minor version must
be within two minor versions of the control plane version. The node pool version cannot be greater than the control
plane version. For more information see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool">upgrading a node
pool</a>.</p>
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
<a href="#containerservice.azure.com/v1beta20210501.OSDiskType_Status">
OSDiskType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osSKU</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.OSSKU_Status">
OSSKU_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.OSType_Status">
OSType_Status
</a>
</em>
</td>
<td>
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
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}</p>
</td>
</tr>
<tr>
<td>
<code>powerState</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.PowerState_StatusARM">
PowerState_StatusARM
</a>
</em>
</td>
<td>
<p>PowerState: Describes whether the Agent Pool is Running or Stopped</p>
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
<code>scaleSetEvictionPolicy</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ScaleSetEvictionPolicy_Status">
ScaleSetEvictionPolicy_Status
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
<a href="#containerservice.azure.com/v1beta20210501.ScaleSetPriority_Status">
ScaleSetPriority_Status
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
<a href="#containerservice.azure.com/v1beta20210501.AgentPoolType_Status">
AgentPoolType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>upgradeSettings</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.AgentPoolUpgradeSettings_StatusARM">
AgentPoolUpgradeSettings_StatusARM
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
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAutoUpgradeProfile">ManagedClusterAutoUpgradeProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusters_Spec">ManagedClusters_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterAutoUpgradeProfile">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterAutoUpgradeProfile</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>upgradeChannel</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAutoUpgradeProfileUpgradeChannel">
ManagedClusterAutoUpgradeProfileUpgradeChannel
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAutoUpgradeProfileARM">ManagedClusterAutoUpgradeProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesARM">ManagedClusterPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterAutoUpgradeProfile">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterAutoUpgradeProfile</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>upgradeChannel</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAutoUpgradeProfileUpgradeChannel">
ManagedClusterAutoUpgradeProfileUpgradeChannel
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAutoUpgradeProfileStatusUpgradeChannel">ManagedClusterAutoUpgradeProfileStatusUpgradeChannel
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAutoUpgradeProfile_Status">ManagedClusterAutoUpgradeProfile_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAutoUpgradeProfile_StatusARM">ManagedClusterAutoUpgradeProfile_StatusARM</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAutoUpgradeProfileUpgradeChannel">ManagedClusterAutoUpgradeProfileUpgradeChannel
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAutoUpgradeProfile">ManagedClusterAutoUpgradeProfile</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAutoUpgradeProfileARM">ManagedClusterAutoUpgradeProfileARM</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAutoUpgradeProfile_Status">ManagedClusterAutoUpgradeProfile_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedCluster_Status">ManagedCluster_Status</a>)
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
<code>upgradeChannel</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAutoUpgradeProfileStatusUpgradeChannel">
ManagedClusterAutoUpgradeProfileStatusUpgradeChannel
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterAutoUpgradeProfile_StatusARM">ManagedClusterAutoUpgradeProfile_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterProperties_StatusARM">ManagedClusterProperties_StatusARM</a>)
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
<code>upgradeChannel</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAutoUpgradeProfileStatusUpgradeChannel">
ManagedClusterAutoUpgradeProfileStatusUpgradeChannel
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterHTTPProxyConfig">ManagedClusterHTTPProxyConfig
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusters_Spec">ManagedClusters_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterHTTPProxyConfig">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterHTTPProxyConfig</a></p>
</div>
<table>
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterHTTPProxyConfigARM">ManagedClusterHTTPProxyConfigARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesARM">ManagedClusterPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterHTTPProxyConfig">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterHTTPProxyConfig</a></p>
</div>
<table>
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterHTTPProxyConfig_Status">ManagedClusterHTTPProxyConfig_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedCluster_Status">ManagedCluster_Status</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterHTTPProxyConfig_StatusARM">ManagedClusterHTTPProxyConfig_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterProperties_StatusARM">ManagedClusterProperties_StatusARM</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterIdentity">ManagedClusterIdentity
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusters_Spec">ManagedClusters_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterIdentity">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterIdentity</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterIdentityType">
ManagedClusterIdentityType
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
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
map[string]k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1.JSON
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: The keys must be ARM resource IDs in the form:
&lsquo;/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}&rsquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterIdentityARM">ManagedClusterIdentityARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusters_SpecARM">ManagedClusters_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterIdentity">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterIdentity</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterIdentityType">
ManagedClusterIdentityType
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
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
map[string]k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1.JSON
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: The keys must be ARM resource IDs in the form:
&lsquo;/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}&rsquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterIdentityStatusType">ManagedClusterIdentityStatusType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterIdentity_Status">ManagedClusterIdentity_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterIdentity_StatusARM">ManagedClusterIdentity_StatusARM</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterIdentityType">ManagedClusterIdentityType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterIdentity">ManagedClusterIdentity</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterIdentityARM">ManagedClusterIdentityARM</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterIdentity_Status">ManagedClusterIdentity_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedCluster_Status">ManagedCluster_Status</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterIdentityStatusType">
ManagedClusterIdentityStatusType
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterIdentity_Status_UserAssignedIdentities">
map[string]./api/containerservice/v1beta20210501.ManagedClusterIdentity_Status_UserAssignedIdentities
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: The keys must be ARM resource IDs in the form:
&lsquo;/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}&rsquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterIdentity_StatusARM">ManagedClusterIdentity_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedCluster_StatusARM">ManagedCluster_StatusARM</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterIdentityStatusType">
ManagedClusterIdentityStatusType
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterIdentity_Status_UserAssignedIdentitiesARM">
map[string]./api/containerservice/v1beta20210501.ManagedClusterIdentity_Status_UserAssignedIdentitiesARM
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: The keys must be ARM resource IDs in the form:
&lsquo;/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}&rsquo;.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterIdentity_Status_UserAssignedIdentities">ManagedClusterIdentity_Status_UserAssignedIdentities
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterIdentity_Status">ManagedClusterIdentity_Status</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterIdentity_Status_UserAssignedIdentitiesARM">ManagedClusterIdentity_Status_UserAssignedIdentitiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterIdentity_StatusARM">ManagedClusterIdentity_StatusARM</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile">ManagedClusterLoadBalancerProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfile">ContainerServiceNetworkProfile</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterLoadBalancerProfile">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterLoadBalancerProfile</a></p>
</div>
<table>
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
<code>effectiveOutboundIPs</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ResourceReference">
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfileManagedOutboundIPs">
ManagedClusterLoadBalancerProfileManagedOutboundIPs
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfileOutboundIPPrefixes">
ManagedClusterLoadBalancerProfileOutboundIPPrefixes
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfileOutboundIPs">
ManagedClusterLoadBalancerProfileOutboundIPs
</a>
</em>
</td>
<td>
<p>OutboundIPs: Desired outbound IP resources for the cluster load balancer.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfileARM">ManagedClusterLoadBalancerProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileARM">ContainerServiceNetworkProfileARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterLoadBalancerProfile">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterLoadBalancerProfile</a></p>
</div>
<table>
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
<code>effectiveOutboundIPs</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ResourceReferenceARM">
[]ResourceReferenceARM
</a>
</em>
</td>
<td>
<p>EffectiveOutboundIPs: The effective outbound IP resources of the cluster load balancer.</p>
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfileManagedOutboundIPsARM">
ManagedClusterLoadBalancerProfileManagedOutboundIPsARM
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfileOutboundIPPrefixesARM">
ManagedClusterLoadBalancerProfileOutboundIPPrefixesARM
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfileOutboundIPsARM">
ManagedClusterLoadBalancerProfileOutboundIPsARM
</a>
</em>
</td>
<td>
<p>OutboundIPs: Desired outbound IP resources for the cluster load balancer.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfileManagedOutboundIPs">ManagedClusterLoadBalancerProfileManagedOutboundIPs
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile">ManagedClusterLoadBalancerProfile</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterLoadBalancerProfileManagedOutboundIPs">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterLoadBalancerProfileManagedOutboundIPs</a></p>
</div>
<table>
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
<p>Count: The desired number of outbound IPs created/managed by Azure for the cluster load balancer. Allowed values must be
in the range of 1 to 100 (inclusive). The default value is 1.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfileManagedOutboundIPsARM">ManagedClusterLoadBalancerProfileManagedOutboundIPsARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfileARM">ManagedClusterLoadBalancerProfileARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterLoadBalancerProfileManagedOutboundIPs">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterLoadBalancerProfileManagedOutboundIPs</a></p>
</div>
<table>
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
<p>Count: The desired number of outbound IPs created/managed by Azure for the cluster load balancer. Allowed values must be
in the range of 1 to 100 (inclusive). The default value is 1.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfileOutboundIPPrefixes">ManagedClusterLoadBalancerProfileOutboundIPPrefixes
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile">ManagedClusterLoadBalancerProfile</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterLoadBalancerProfileOutboundIPPrefixes">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterLoadBalancerProfileOutboundIPPrefixes</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.ResourceReference">
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfileOutboundIPPrefixesARM">ManagedClusterLoadBalancerProfileOutboundIPPrefixesARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfileARM">ManagedClusterLoadBalancerProfileARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterLoadBalancerProfileOutboundIPPrefixes">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterLoadBalancerProfileOutboundIPPrefixes</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.ResourceReferenceARM">
[]ResourceReferenceARM
</a>
</em>
</td>
<td>
<p>PublicIPPrefixes: A list of public IP prefix resources.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfileOutboundIPs">ManagedClusterLoadBalancerProfileOutboundIPs
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile">ManagedClusterLoadBalancerProfile</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterLoadBalancerProfileOutboundIPs">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterLoadBalancerProfileOutboundIPs</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.ResourceReference">
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfileOutboundIPsARM">ManagedClusterLoadBalancerProfileOutboundIPsARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfileARM">ManagedClusterLoadBalancerProfileARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterLoadBalancerProfileOutboundIPs">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterLoadBalancerProfileOutboundIPs</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.ResourceReferenceARM">
[]ResourceReferenceARM
</a>
</em>
</td>
<td>
<p>PublicIPs: A list of public IP resources.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_Status">ManagedClusterLoadBalancerProfile_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfile_Status">ContainerServiceNetworkProfile_Status</a>)
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
<code>effectiveOutboundIPs</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ResourceReference_Status">
[]ResourceReference_Status
</a>
</em>
</td>
<td>
<p>EffectiveOutboundIPs: The effective outbound IP resources of the cluster load balancer.</p>
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_Status_ManagedOutboundIPs">
ManagedClusterLoadBalancerProfile_Status_ManagedOutboundIPs
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_Status_OutboundIPPrefixes">
ManagedClusterLoadBalancerProfile_Status_OutboundIPPrefixes
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_Status_OutboundIPs">
ManagedClusterLoadBalancerProfile_Status_OutboundIPs
</a>
</em>
</td>
<td>
<p>OutboundIPs: Desired outbound IP resources for the cluster load balancer.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_StatusARM">ManagedClusterLoadBalancerProfile_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfile_StatusARM">ContainerServiceNetworkProfile_StatusARM</a>)
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
<code>effectiveOutboundIPs</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ResourceReference_StatusARM">
[]ResourceReference_StatusARM
</a>
</em>
</td>
<td>
<p>EffectiveOutboundIPs: The effective outbound IP resources of the cluster load balancer.</p>
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_Status_ManagedOutboundIPsARM">
ManagedClusterLoadBalancerProfile_Status_ManagedOutboundIPsARM
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_Status_OutboundIPPrefixesARM">
ManagedClusterLoadBalancerProfile_Status_OutboundIPPrefixesARM
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_Status_OutboundIPsARM">
ManagedClusterLoadBalancerProfile_Status_OutboundIPsARM
</a>
</em>
</td>
<td>
<p>OutboundIPs: Desired outbound IP resources for the cluster load balancer.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_Status_ManagedOutboundIPs">ManagedClusterLoadBalancerProfile_Status_ManagedOutboundIPs
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_Status">ManagedClusterLoadBalancerProfile_Status</a>)
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
<p>Count: The desired number of outbound IPs created/managed by Azure for the cluster load balancer. Allowed values must be
in the range of 1 to 100 (inclusive). The default value is 1.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_Status_ManagedOutboundIPsARM">ManagedClusterLoadBalancerProfile_Status_ManagedOutboundIPsARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_StatusARM">ManagedClusterLoadBalancerProfile_StatusARM</a>)
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
<p>Count: The desired number of outbound IPs created/managed by Azure for the cluster load balancer. Allowed values must be
in the range of 1 to 100 (inclusive). The default value is 1.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_Status_OutboundIPPrefixes">ManagedClusterLoadBalancerProfile_Status_OutboundIPPrefixes
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_Status">ManagedClusterLoadBalancerProfile_Status</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.ResourceReference_Status">
[]ResourceReference_Status
</a>
</em>
</td>
<td>
<p>PublicIPPrefixes: A list of public IP prefix resources.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_Status_OutboundIPPrefixesARM">ManagedClusterLoadBalancerProfile_Status_OutboundIPPrefixesARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_StatusARM">ManagedClusterLoadBalancerProfile_StatusARM</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.ResourceReference_StatusARM">
[]ResourceReference_StatusARM
</a>
</em>
</td>
<td>
<p>PublicIPPrefixes: A list of public IP prefix resources.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_Status_OutboundIPs">ManagedClusterLoadBalancerProfile_Status_OutboundIPs
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_Status">ManagedClusterLoadBalancerProfile_Status</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.ResourceReference_Status">
[]ResourceReference_Status
</a>
</em>
</td>
<td>
<p>PublicIPs: A list of public IP resources.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_Status_OutboundIPsARM">ManagedClusterLoadBalancerProfile_Status_OutboundIPsARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_StatusARM">ManagedClusterLoadBalancerProfile_StatusARM</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.ResourceReference_StatusARM">
[]ResourceReference_StatusARM
</a>
</em>
</td>
<td>
<p>PublicIPs: A list of public IP resources.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentity">ManagedClusterPodIdentity
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProfile">ManagedClusterPodIdentityProfile</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterPodIdentity">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterPodIdentity</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.UserAssignedIdentity">
UserAssignedIdentity
</a>
</em>
</td>
<td>
<p>Identity: Details about a user assigned identity.</p>
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityARM">ManagedClusterPodIdentityARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProfileARM">ManagedClusterPodIdentityProfileARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterPodIdentity">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterPodIdentity</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.UserAssignedIdentityARM">
UserAssignedIdentityARM
</a>
</em>
</td>
<td>
<p>Identity: Details about a user assigned identity.</p>
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityException">ManagedClusterPodIdentityException
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProfile">ManagedClusterPodIdentityProfile</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterPodIdentityException">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterPodIdentityException</a></p>
</div>
<table>
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityExceptionARM">ManagedClusterPodIdentityExceptionARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProfileARM">ManagedClusterPodIdentityProfileARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterPodIdentityException">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterPodIdentityException</a></p>
</div>
<table>
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityException_Status">ManagedClusterPodIdentityException_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProfile_Status">ManagedClusterPodIdentityProfile_Status</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityException_StatusARM">ManagedClusterPodIdentityException_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProfile_StatusARM">ManagedClusterPodIdentityProfile_StatusARM</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProfile">ManagedClusterPodIdentityProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusters_Spec">ManagedClusters_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterPodIdentityProfile">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterPodIdentityProfile</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentity">
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityException">
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProfileARM">ManagedClusterPodIdentityProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesARM">ManagedClusterPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterPodIdentityProfile">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterPodIdentityProfile</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityARM">
[]ManagedClusterPodIdentityARM
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityExceptionARM">
[]ManagedClusterPodIdentityExceptionARM
</a>
</em>
</td>
<td>
<p>UserAssignedIdentityExceptions: The pod identity exceptions to allow.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProfile_Status">ManagedClusterPodIdentityProfile_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedCluster_Status">ManagedCluster_Status</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentity_Status">
[]ManagedClusterPodIdentity_Status
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityException_Status">
[]ManagedClusterPodIdentityException_Status
</a>
</em>
</td>
<td>
<p>UserAssignedIdentityExceptions: The pod identity exceptions to allow.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProfile_StatusARM">ManagedClusterPodIdentityProfile_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterProperties_StatusARM">ManagedClusterProperties_StatusARM</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentity_StatusARM">
[]ManagedClusterPodIdentity_StatusARM
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityException_StatusARM">
[]ManagedClusterPodIdentityException_StatusARM
</a>
</em>
</td>
<td>
<p>UserAssignedIdentityExceptions: The pod identity exceptions to allow.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProvisioningErrorBody_Status">ManagedClusterPodIdentityProvisioningErrorBody_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProvisioningError_Status">ManagedClusterPodIdentityProvisioningError_Status</a>)
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
<code>details</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProvisioningErrorBody_Status_Unrolled">
[]ManagedClusterPodIdentityProvisioningErrorBody_Status_Unrolled
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProvisioningErrorBody_StatusARM">ManagedClusterPodIdentityProvisioningErrorBody_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProvisioningError_StatusARM">ManagedClusterPodIdentityProvisioningError_StatusARM</a>)
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
<code>details</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProvisioningErrorBody_Status_UnrolledARM">
[]ManagedClusterPodIdentityProvisioningErrorBody_Status_UnrolledARM
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProvisioningErrorBody_Status_Unrolled">ManagedClusterPodIdentityProvisioningErrorBody_Status_Unrolled
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProvisioningErrorBody_Status">ManagedClusterPodIdentityProvisioningErrorBody_Status</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProvisioningErrorBody_Status_UnrolledARM">ManagedClusterPodIdentityProvisioningErrorBody_Status_UnrolledARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProvisioningErrorBody_StatusARM">ManagedClusterPodIdentityProvisioningErrorBody_StatusARM</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProvisioningError_Status">ManagedClusterPodIdentityProvisioningError_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentity_Status_ProvisioningInfo">ManagedClusterPodIdentity_Status_ProvisioningInfo</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProvisioningErrorBody_Status">
ManagedClusterPodIdentityProvisioningErrorBody_Status
</a>
</em>
</td>
<td>
<p>Error: Details about the error.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProvisioningError_StatusARM">ManagedClusterPodIdentityProvisioningError_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentity_Status_ProvisioningInfoARM">ManagedClusterPodIdentity_Status_ProvisioningInfoARM</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProvisioningErrorBody_StatusARM">
ManagedClusterPodIdentityProvisioningErrorBody_StatusARM
</a>
</em>
</td>
<td>
<p>Error: Details about the error.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityStatusProvisioningState">ManagedClusterPodIdentityStatusProvisioningState
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentity_Status">ManagedClusterPodIdentity_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentity_StatusARM">ManagedClusterPodIdentity_StatusARM</a>)
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
</tr><tr><td><p>&#34;Deleting&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Failed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Updating&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentity_Status">ManagedClusterPodIdentity_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProfile_Status">ManagedClusterPodIdentityProfile_Status</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.UserAssignedIdentity_Status">
UserAssignedIdentity_Status
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentity_Status_ProvisioningInfo">
ManagedClusterPodIdentity_Status_ProvisioningInfo
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityStatusProvisioningState">
ManagedClusterPodIdentityStatusProvisioningState
</a>
</em>
</td>
<td>
<p>ProvisioningState: The current provisioning state of the pod identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentity_StatusARM">ManagedClusterPodIdentity_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProfile_StatusARM">ManagedClusterPodIdentityProfile_StatusARM</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.UserAssignedIdentity_StatusARM">
UserAssignedIdentity_StatusARM
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentity_Status_ProvisioningInfoARM">
ManagedClusterPodIdentity_Status_ProvisioningInfoARM
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityStatusProvisioningState">
ManagedClusterPodIdentityStatusProvisioningState
</a>
</em>
</td>
<td>
<p>ProvisioningState: The current provisioning state of the pod identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentity_Status_ProvisioningInfo">ManagedClusterPodIdentity_Status_ProvisioningInfo
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentity_Status">ManagedClusterPodIdentity_Status</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProvisioningError_Status">
ManagedClusterPodIdentityProvisioningError_Status
</a>
</em>
</td>
<td>
<p>Error: Pod identity assignment error (if any).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentity_Status_ProvisioningInfoARM">ManagedClusterPodIdentity_Status_ProvisioningInfoARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentity_StatusARM">ManagedClusterPodIdentity_StatusARM</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProvisioningError_StatusARM">
ManagedClusterPodIdentityProvisioningError_StatusARM
</a>
</em>
</td>
<td>
<p>Error: Pod identity assignment error (if any).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesARM">ManagedClusterPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusters_SpecARM">ManagedClusters_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterProperties">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterProperties</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAADProfileARM">
ManagedClusterAADProfileARM
</a>
</em>
</td>
<td>
<p>AadProfile: For more details see <a href="https://docs.microsoft.com/azure/aks/managed-aad">managed AAD on AKS</a>.</p>
</td>
</tr>
<tr>
<td>
<code>addonProfiles</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAddonProfileARM">
map[string]./api/containerservice/v1beta20210501.ManagedClusterAddonProfileARM
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileARM">
[]ManagedClusterAgentPoolProfileARM
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAPIServerAccessProfileARM">
ManagedClusterAPIServerAccessProfileARM
</a>
</em>
</td>
<td>
<p>ApiServerAccessProfile: Access profile for managed cluster API server.</p>
</td>
</tr>
<tr>
<td>
<code>autoScalerProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesAutoScalerProfileARM">
ManagedClusterPropertiesAutoScalerProfileARM
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAutoUpgradeProfileARM">
ManagedClusterAutoUpgradeProfileARM
</a>
</em>
</td>
<td>
<p>AutoUpgradeProfile: Auto upgrade profile for a managed cluster.</p>
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
<code>enablePodSecurityPolicy</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePodSecurityPolicy: (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set
for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.</p>
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
<code>httpProxyConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterHTTPProxyConfigARM">
ManagedClusterHTTPProxyConfigARM
</a>
</em>
</td>
<td>
<p>HttpProxyConfig: Cluster HTTP proxy configuration.</p>
</td>
</tr>
<tr>
<td>
<code>identityProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.Componentsqit0EtschemasmanagedclusterpropertiespropertiesidentityprofileadditionalpropertiesARM">
map[string]./api/containerservice/v1beta20210501.Componentsqit0EtschemasmanagedclusterpropertiespropertiesidentityprofileadditionalpropertiesARM
</a>
</em>
</td>
<td>
<p>IdentityProfile: Identities associated with the cluster.</p>
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
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceLinuxProfileARM">
ContainerServiceLinuxProfileARM
</a>
</em>
</td>
<td>
<p>LinuxProfile: Profile for Linux VMs in the container service cluster.</p>
</td>
</tr>
<tr>
<td>
<code>networkProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfileARM">
ContainerServiceNetworkProfileARM
</a>
</em>
</td>
<td>
<p>NetworkProfile: Profile of network configuration.</p>
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
<code>podIdentityProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProfileARM">
ManagedClusterPodIdentityProfileARM
</a>
</em>
</td>
<td>
<p>PodIdentityProfile: See <a href="https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity">use AAD pod identity</a> for more
details on pod identity integration.</p>
</td>
</tr>
<tr>
<td>
<code>privateLinkResources</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.PrivateLinkResourceARM">
[]PrivateLinkResourceARM
</a>
</em>
</td>
<td>
<p>PrivateLinkResources: Private link resources associated with the cluster.</p>
</td>
</tr>
<tr>
<td>
<code>servicePrincipalProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterServicePrincipalProfileARM">
ManagedClusterServicePrincipalProfileARM
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
<code>windowsProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterWindowsProfileARM">
ManagedClusterWindowsProfileARM
</a>
</em>
</td>
<td>
<p>WindowsProfile: Profile for Windows VMs in the managed cluster.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesAutoScalerProfile">ManagedClusterPropertiesAutoScalerProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusters_Spec">ManagedClusters_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterPropertiesAutoScalerProfile">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterPropertiesAutoScalerProfile</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesAutoScalerProfileExpander">
ManagedClusterPropertiesAutoScalerProfileExpander
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesAutoScalerProfileARM">ManagedClusterPropertiesAutoScalerProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesARM">ManagedClusterPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterPropertiesAutoScalerProfile">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterPropertiesAutoScalerProfile</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesAutoScalerProfileExpander">
ManagedClusterPropertiesAutoScalerProfileExpander
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesAutoScalerProfileExpander">ManagedClusterPropertiesAutoScalerProfileExpander
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesAutoScalerProfile">ManagedClusterPropertiesAutoScalerProfile</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesAutoScalerProfileARM">ManagedClusterPropertiesAutoScalerProfileARM</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesStatusAutoScalerProfileExpander">ManagedClusterPropertiesStatusAutoScalerProfileExpander
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterProperties_Status_AutoScalerProfile">ManagedClusterProperties_Status_AutoScalerProfile</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterProperties_Status_AutoScalerProfileARM">ManagedClusterProperties_Status_AutoScalerProfileARM</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterProperties_StatusARM">ManagedClusterProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedCluster_StatusARM">ManagedCluster_StatusARM</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAADProfile_StatusARM">
ManagedClusterAADProfile_StatusARM
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
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
Kubernetes v1.JSON
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_StatusARM">
[]ManagedClusterAgentPoolProfile_StatusARM
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAPIServerAccessProfile_StatusARM">
ManagedClusterAPIServerAccessProfile_StatusARM
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterProperties_Status_AutoScalerProfileARM">
ManagedClusterProperties_Status_AutoScalerProfileARM
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAutoUpgradeProfile_StatusARM">
ManagedClusterAutoUpgradeProfile_StatusARM
</a>
</em>
</td>
<td>
<p>AutoUpgradeProfile: The auto upgrade configuration.</p>
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
&lsquo;/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{encryptionSetName}&rsquo;</p>
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
<code>enablePodSecurityPolicy</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePodSecurityPolicy: (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set
for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.</p>
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
<code>httpProxyConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterHTTPProxyConfig_StatusARM">
ManagedClusterHTTPProxyConfig_StatusARM
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
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
Kubernetes v1.JSON
</a>
</em>
</td>
<td>
<p>IdentityProfile: Identities associated with the cluster.</p>
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
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceLinuxProfile_StatusARM">
ContainerServiceLinuxProfile_StatusARM
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
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfile_StatusARM">
ContainerServiceNetworkProfile_StatusARM
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
<code>podIdentityProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProfile_StatusARM">
ManagedClusterPodIdentityProfile_StatusARM
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
<a href="#containerservice.azure.com/v1beta20210501.PowerState_StatusARM">
PowerState_StatusARM
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
<a href="#containerservice.azure.com/v1beta20210501.PrivateLinkResource_StatusARM">
[]PrivateLinkResource_StatusARM
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
<code>servicePrincipalProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterServicePrincipalProfile_StatusARM">
ManagedClusterServicePrincipalProfile_StatusARM
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
<code>windowsProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterWindowsProfile_StatusARM">
ManagedClusterWindowsProfile_StatusARM
</a>
</em>
</td>
<td>
<p>WindowsProfile: The profile for Windows VMs in the Managed Cluster.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterProperties_Status_AutoScalerProfile">ManagedClusterProperties_Status_AutoScalerProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedCluster_Status">ManagedCluster_Status</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesStatusAutoScalerProfileExpander">
ManagedClusterPropertiesStatusAutoScalerProfileExpander
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterProperties_Status_AutoScalerProfileARM">ManagedClusterProperties_Status_AutoScalerProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterProperties_StatusARM">ManagedClusterProperties_StatusARM</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesStatusAutoScalerProfileExpander">
ManagedClusterPropertiesStatusAutoScalerProfileExpander
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterSKU">ManagedClusterSKU
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusters_Spec">ManagedClusters_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterSKU">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterSKU</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterSKUName">
ManagedClusterSKUName
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterSKUTier">
ManagedClusterSKUTier
</a>
</em>
</td>
<td>
<p>Tier: If not specified, the default is &lsquo;Free&rsquo;. See <a href="https://docs.microsoft.com/azure/aks/uptime-sla">uptime SLA</a> for
more details.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterSKUARM">ManagedClusterSKUARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusters_SpecARM">ManagedClusters_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterSKU">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterSKU</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterSKUName">
ManagedClusterSKUName
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterSKUTier">
ManagedClusterSKUTier
</a>
</em>
</td>
<td>
<p>Tier: If not specified, the default is &lsquo;Free&rsquo;. See <a href="https://docs.microsoft.com/azure/aks/uptime-sla">uptime SLA</a> for
more details.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterSKUName">ManagedClusterSKUName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterSKU">ManagedClusterSKU</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterSKUARM">ManagedClusterSKUARM</a>)
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
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterSKUStatusName">ManagedClusterSKUStatusName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterSKU_Status">ManagedClusterSKU_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterSKU_StatusARM">ManagedClusterSKU_StatusARM</a>)
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
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterSKUStatusTier">ManagedClusterSKUStatusTier
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterSKU_Status">ManagedClusterSKU_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterSKU_StatusARM">ManagedClusterSKU_StatusARM</a>)
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
</tr><tr><td><p>&#34;Paid&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterSKUTier">ManagedClusterSKUTier
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterSKU">ManagedClusterSKU</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterSKUARM">ManagedClusterSKUARM</a>)
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
</tr><tr><td><p>&#34;Paid&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterSKU_Status">ManagedClusterSKU_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedCluster_Status">ManagedCluster_Status</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterSKUStatusName">
ManagedClusterSKUStatusName
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterSKUStatusTier">
ManagedClusterSKUStatusTier
</a>
</em>
</td>
<td>
<p>Tier: If not specified, the default is &lsquo;Free&rsquo;. See <a href="https://docs.microsoft.com/azure/aks/uptime-sla">uptime SLA</a> for
more details.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterSKU_StatusARM">ManagedClusterSKU_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedCluster_StatusARM">ManagedCluster_StatusARM</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterSKUStatusName">
ManagedClusterSKUStatusName
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterSKUStatusTier">
ManagedClusterSKUStatusTier
</a>
</em>
</td>
<td>
<p>Tier: If not specified, the default is &lsquo;Free&rsquo;. See <a href="https://docs.microsoft.com/azure/aks/uptime-sla">uptime SLA</a> for
more details.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterServicePrincipalProfile">ManagedClusterServicePrincipalProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusters_Spec">ManagedClusters_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterServicePrincipalProfile">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterServicePrincipalProfile</a></p>
</div>
<table>
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterServicePrincipalProfileARM">ManagedClusterServicePrincipalProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesARM">ManagedClusterPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterServicePrincipalProfile">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterServicePrincipalProfile</a></p>
</div>
<table>
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterServicePrincipalProfile_Status">ManagedClusterServicePrincipalProfile_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedCluster_Status">ManagedCluster_Status</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterServicePrincipalProfile_StatusARM">ManagedClusterServicePrincipalProfile_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterProperties_StatusARM">ManagedClusterProperties_StatusARM</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterWindowsProfile">ManagedClusterWindowsProfile
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusters_Spec">ManagedClusters_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterWindowsProfile">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterWindowsProfile</a></p>
</div>
<table>
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
<code>licenseType</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterWindowsProfileLicenseType">
ManagedClusterWindowsProfileLicenseType
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterWindowsProfileARM">ManagedClusterWindowsProfileARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesARM">ManagedClusterPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterWindowsProfile">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ManagedClusterWindowsProfile</a></p>
</div>
<table>
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
<code>licenseType</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterWindowsProfileLicenseType">
ManagedClusterWindowsProfileLicenseType
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterWindowsProfileLicenseType">ManagedClusterWindowsProfileLicenseType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterWindowsProfile">ManagedClusterWindowsProfile</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterWindowsProfileARM">ManagedClusterWindowsProfileARM</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterWindowsProfileStatusLicenseType">ManagedClusterWindowsProfileStatusLicenseType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterWindowsProfile_Status">ManagedClusterWindowsProfile_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterWindowsProfile_StatusARM">ManagedClusterWindowsProfile_StatusARM</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterWindowsProfile_Status">ManagedClusterWindowsProfile_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedCluster_Status">ManagedCluster_Status</a>)
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
<code>licenseType</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterWindowsProfileStatusLicenseType">
ManagedClusterWindowsProfileStatusLicenseType
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusterWindowsProfile_StatusARM">ManagedClusterWindowsProfile_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterProperties_StatusARM">ManagedClusterProperties_StatusARM</a>)
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
<code>licenseType</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterWindowsProfileStatusLicenseType">
ManagedClusterWindowsProfileStatusLicenseType
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedCluster_Status">ManagedCluster_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedCluster">ManagedCluster</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAADProfile_Status">
ManagedClusterAADProfile_Status
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
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
Kubernetes v1.JSON
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_Status">
[]ManagedClusterAgentPoolProfile_Status
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAPIServerAccessProfile_Status">
ManagedClusterAPIServerAccessProfile_Status
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterProperties_Status_AutoScalerProfile">
ManagedClusterProperties_Status_AutoScalerProfile
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAutoUpgradeProfile_Status">
ManagedClusterAutoUpgradeProfile_Status
</a>
</em>
</td>
<td>
<p>AutoUpgradeProfile: The auto upgrade configuration.</p>
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
&lsquo;/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{encryptionSetName}&rsquo;</p>
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
<code>enablePodSecurityPolicy</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePodSecurityPolicy: (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set
for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.</p>
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
<a href="#containerservice.azure.com/v1beta20210501.ExtendedLocation_Status">
ExtendedLocation_Status
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
<code>httpProxyConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterHTTPProxyConfig_Status">
ManagedClusterHTTPProxyConfig_Status
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
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterIdentity_Status">
ManagedClusterIdentity_Status
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
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
Kubernetes v1.JSON
</a>
</em>
</td>
<td>
<p>IdentityProfile: Identities associated with the cluster.</p>
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
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceLinuxProfile_Status">
ContainerServiceLinuxProfile_Status
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
<p>Location: Resource location</p>
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
<p>Name: Resource name</p>
</td>
</tr>
<tr>
<td>
<code>networkProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfile_Status">
ContainerServiceNetworkProfile_Status
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
<code>podIdentityProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProfile_Status">
ManagedClusterPodIdentityProfile_Status
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
<a href="#containerservice.azure.com/v1beta20210501.PowerState_Status">
PowerState_Status
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
<a href="#containerservice.azure.com/v1beta20210501.PrivateLinkResource_Status">
[]PrivateLinkResource_Status
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
<code>servicePrincipalProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterServicePrincipalProfile_Status">
ManagedClusterServicePrincipalProfile_Status
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterSKU_Status">
ManagedClusterSKU_Status
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
<p>Tags: Resource tags</p>
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
<code>windowsProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterWindowsProfile_Status">
ManagedClusterWindowsProfile_Status
</a>
</em>
</td>
<td>
<p>WindowsProfile: The profile for Windows VMs in the Managed Cluster.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedCluster_StatusARM">ManagedCluster_StatusARM
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
<a href="#containerservice.azure.com/v1beta20210501.ExtendedLocation_StatusARM">
ExtendedLocation_StatusARM
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
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterIdentity_StatusARM">
ManagedClusterIdentity_StatusARM
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
<p>Location: Resource location</p>
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
<p>Name: Resource name</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterProperties_StatusARM">
ManagedClusterProperties_StatusARM
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterSKU_StatusARM">
ManagedClusterSKU_StatusARM
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
<p>Tags: Resource tags</p>
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
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClustersAgentPool">ManagedClustersAgentPool
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/resourceDefinitions/managedClusters_agentPools">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/resourceDefinitions/managedClusters_agentPools</a></p>
</div>
<table>
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClustersAgentPools_Spec">
ManagedClustersAgentPools_Spec
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesGpuInstanceProfile">
ManagedClusterAgentPoolProfilePropertiesGpuInstanceProfile
</a>
</em>
</td>
<td>
<p>GpuInstanceProfile: GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.KubeletConfig">
KubeletConfig
</a>
</em>
</td>
<td>
<p>KubeletConfig: See <a href="https://docs.microsoft.com/azure/aks/custom-node-configuration">AKS custom node configuration</a> for
more details.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesKubeletDiskType">
ManagedClusterAgentPoolProfilePropertiesKubeletDiskType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>linuxOSConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.LinuxOSConfig">
LinuxOSConfig
</a>
</em>
</td>
<td>
<p>LinuxOSConfig: See <a href="https://docs.microsoft.com/azure/aks/custom-node-configuration">AKS custom node configuration</a> for
more details.</p>
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesMode">
ManagedClusterAgentPoolProfilePropertiesMode
</a>
</em>
</td>
<td>
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
<code>nodePublicIPPrefixIDReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>NodePublicIPPrefixIDReference: This is of the form:
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}</p>
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
<p>OrchestratorVersion: As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes
version. The node pool version must have the same major version as the control plane. The node pool minor version must
be within two minor versions of the control plane version. The node pool version cannot be greater than the control
plane version. For more information see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool">upgrading a node
pool</a>.</p>
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
<p>OsDiskSizeGB: OS Disk Size in GB to be used to specify the disk size for every machine in the master/agent pool. If you
specify 0, it will apply the default osDisk size according to the vmSize specified.</p>
</td>
</tr>
<tr>
<td>
<code>osDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesOsDiskType">
ManagedClusterAgentPoolProfilePropertiesOsDiskType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osSKU</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesOsSKU">
ManagedClusterAgentPoolProfilePropertiesOsSKU
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesOsType">
ManagedClusterAgentPoolProfilePropertiesOsType
</a>
</em>
</td>
<td>
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
<code>podSubnetIDReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>PodSubnetIDReference: If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more
details). This is of the form:
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}</p>
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
<code>scaleSetEvictionPolicy</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesScaleSetEvictionPolicy">
ManagedClusterAgentPoolProfilePropertiesScaleSetEvictionPolicy
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesScaleSetPriority">
ManagedClusterAgentPoolProfilePropertiesScaleSetPriority
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesType">
ManagedClusterAgentPoolProfilePropertiesType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>upgradeSettings</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.AgentPoolUpgradeSettings">
AgentPoolUpgradeSettings
</a>
</em>
</td>
<td>
<p>UpgradeSettings: Settings for upgrading an agentpool</p>
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
<code>vnetSubnetIDReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>VnetSubnetIDReference: If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is
specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form:
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.AgentPool_Status">
AgentPool_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClustersAgentPoolsSpecAPIVersion">ManagedClustersAgentPoolsSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-05-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClustersAgentPools_Spec">ManagedClustersAgentPools_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClustersAgentPool">ManagedClustersAgentPool</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesGpuInstanceProfile">
ManagedClusterAgentPoolProfilePropertiesGpuInstanceProfile
</a>
</em>
</td>
<td>
<p>GpuInstanceProfile: GPUInstanceProfile to be used to specify GPU MIG instance profile for supported GPU VM SKU.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.KubeletConfig">
KubeletConfig
</a>
</em>
</td>
<td>
<p>KubeletConfig: See <a href="https://docs.microsoft.com/azure/aks/custom-node-configuration">AKS custom node configuration</a> for
more details.</p>
</td>
</tr>
<tr>
<td>
<code>kubeletDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesKubeletDiskType">
ManagedClusterAgentPoolProfilePropertiesKubeletDiskType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>linuxOSConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.LinuxOSConfig">
LinuxOSConfig
</a>
</em>
</td>
<td>
<p>LinuxOSConfig: See <a href="https://docs.microsoft.com/azure/aks/custom-node-configuration">AKS custom node configuration</a> for
more details.</p>
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesMode">
ManagedClusterAgentPoolProfilePropertiesMode
</a>
</em>
</td>
<td>
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
<code>nodePublicIPPrefixIDReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>NodePublicIPPrefixIDReference: This is of the form:
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPPrefixes/{publicIPPrefixName}</p>
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
<p>OrchestratorVersion: As a best practice, you should upgrade all node pools in an AKS cluster to the same Kubernetes
version. The node pool version must have the same major version as the control plane. The node pool minor version must
be within two minor versions of the control plane version. The node pool version cannot be greater than the control
plane version. For more information see <a href="https://docs.microsoft.com/azure/aks/use-multiple-node-pools#upgrade-a-node-pool">upgrading a node
pool</a>.</p>
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
<p>OsDiskSizeGB: OS Disk Size in GB to be used to specify the disk size for every machine in the master/agent pool. If you
specify 0, it will apply the default osDisk size according to the vmSize specified.</p>
</td>
</tr>
<tr>
<td>
<code>osDiskType</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesOsDiskType">
ManagedClusterAgentPoolProfilePropertiesOsDiskType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osSKU</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesOsSKU">
ManagedClusterAgentPoolProfilePropertiesOsSKU
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>osType</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesOsType">
ManagedClusterAgentPoolProfilePropertiesOsType
</a>
</em>
</td>
<td>
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
<code>podSubnetIDReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>PodSubnetIDReference: If omitted, pod IPs are statically assigned on the node subnet (see vnetSubnetID for more
details). This is of the form:
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}</p>
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
<code>scaleSetEvictionPolicy</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesScaleSetEvictionPolicy">
ManagedClusterAgentPoolProfilePropertiesScaleSetEvictionPolicy
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesScaleSetPriority">
ManagedClusterAgentPoolProfilePropertiesScaleSetPriority
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesType">
ManagedClusterAgentPoolProfilePropertiesType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>upgradeSettings</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.AgentPoolUpgradeSettings">
AgentPoolUpgradeSettings
</a>
</em>
</td>
<td>
<p>UpgradeSettings: Settings for upgrading an agentpool</p>
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
<code>vnetSubnetIDReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>VnetSubnetIDReference: If this is not specified, a VNET and subnet will be generated and used. If no podSubnetID is
specified, this applies to nodes and pods, otherwise it applies to just nodes. This is of the form:
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClustersAgentPools_SpecARM">ManagedClustersAgentPools_SpecARM
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
<p>Name: The name of the agent pool.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfilePropertiesARM">
ManagedClusterAgentPoolProfilePropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties for the container service agent pool profile.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClustersSpecAPIVersion">ManagedClustersSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-05-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusters_Spec">ManagedClusters_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedCluster">ManagedCluster</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAADProfile">
ManagedClusterAADProfile
</a>
</em>
</td>
<td>
<p>AadProfile: For more details see <a href="https://docs.microsoft.com/azure/aks/managed-aad">managed AAD on AKS</a>.</p>
</td>
</tr>
<tr>
<td>
<code>addonProfiles</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAddonProfile">
map[string]./api/containerservice/v1beta20210501.ManagedClusterAddonProfile
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile">
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAPIServerAccessProfile">
ManagedClusterAPIServerAccessProfile
</a>
</em>
</td>
<td>
<p>ApiServerAccessProfile: Access profile for managed cluster API server.</p>
</td>
</tr>
<tr>
<td>
<code>autoScalerProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesAutoScalerProfile">
ManagedClusterPropertiesAutoScalerProfile
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAutoUpgradeProfile">
ManagedClusterAutoUpgradeProfile
</a>
</em>
</td>
<td>
<p>AutoUpgradeProfile: Auto upgrade profile for a managed cluster.</p>
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
<code>diskEncryptionSetIDReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>DiskEncryptionSetIDReference: This is of the form:
&lsquo;/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{encryptionSetName}&rsquo;</p>
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
<code>enablePodSecurityPolicy</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePodSecurityPolicy: (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set
for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.</p>
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
<a href="#containerservice.azure.com/v1beta20210501.ExtendedLocation">
ExtendedLocation
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The complex type of the extended location.</p>
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
<code>httpProxyConfig</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterHTTPProxyConfig">
ManagedClusterHTTPProxyConfig
</a>
</em>
</td>
<td>
<p>HttpProxyConfig: Cluster HTTP proxy configuration.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterIdentity">
ManagedClusterIdentity
</a>
</em>
</td>
<td>
<p>Identity: Identity for the managed cluster.</p>
</td>
</tr>
<tr>
<td>
<code>identityProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.Componentsqit0Etschemasmanagedclusterpropertiespropertiesidentityprofileadditionalproperties">
map[string]./api/containerservice/v1beta20210501.Componentsqit0Etschemasmanagedclusterpropertiespropertiesidentityprofileadditionalproperties
</a>
</em>
</td>
<td>
<p>IdentityProfile: Identities associated with the cluster.</p>
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
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceLinuxProfile">
ContainerServiceLinuxProfile
</a>
</em>
</td>
<td>
<p>LinuxProfile: Profile for Linux VMs in the container service cluster.</p>
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
<code>networkProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ContainerServiceNetworkProfile">
ContainerServiceNetworkProfile
</a>
</em>
</td>
<td>
<p>NetworkProfile: Profile of network configuration.</p>
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityProfile">
ManagedClusterPodIdentityProfile
</a>
</em>
</td>
<td>
<p>PodIdentityProfile: See <a href="https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity">use AAD pod identity</a> for more
details on pod identity integration.</p>
</td>
</tr>
<tr>
<td>
<code>privateLinkResources</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.PrivateLinkResource">
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
<code>servicePrincipalProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterServicePrincipalProfile">
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
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterSKU">
ManagedClusterSKU
</a>
</em>
</td>
<td>
<p>Sku: The SKU of a Managed Cluster.</p>
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
<code>windowsProfile</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterWindowsProfile">
ManagedClusterWindowsProfile
</a>
</em>
</td>
<td>
<p>WindowsProfile: Profile for Windows VMs in the managed cluster.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ManagedClusters_SpecARM">ManagedClusters_SpecARM
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
<a href="#containerservice.azure.com/v1beta20210501.ExtendedLocationARM">
ExtendedLocationARM
</a>
</em>
</td>
<td>
<p>ExtendedLocation: The complex type of the extended location.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterIdentityARM">
ManagedClusterIdentityARM
</a>
</em>
</td>
<td>
<p>Identity: Identity for the managed cluster.</p>
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
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the managed cluster resource.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesARM">
ManagedClusterPropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the managed cluster.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#containerservice.azure.com/v1beta20210501.ManagedClusterSKUARM">
ManagedClusterSKUARM
</a>
</em>
</td>
<td>
<p>Sku: The SKU of a Managed Cluster.</p>
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
<h3 id="containerservice.azure.com/v1beta20210501.OSDiskType_Status">OSDiskType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.AgentPool_Status">AgentPool_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileProperties_StatusARM">ManagedClusterAgentPoolProfileProperties_StatusARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_Status">ManagedClusterAgentPoolProfile_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_StatusARM">ManagedClusterAgentPoolProfile_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Ephemeral&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Managed&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.OSSKU_Status">OSSKU_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.AgentPool_Status">AgentPool_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileProperties_StatusARM">ManagedClusterAgentPoolProfileProperties_StatusARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_Status">ManagedClusterAgentPoolProfile_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_StatusARM">ManagedClusterAgentPoolProfile_StatusARM</a>)
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
<tbody><tr><td><p>&#34;CBLMariner&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Ubuntu&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.OSType_Status">OSType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.AgentPool_Status">AgentPool_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileProperties_StatusARM">ManagedClusterAgentPoolProfileProperties_StatusARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_Status">ManagedClusterAgentPoolProfile_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_StatusARM">ManagedClusterAgentPoolProfile_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Linux&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Windows&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.PowerStateStatusCode">PowerStateStatusCode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.PowerState_Status">PowerState_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.PowerState_StatusARM">PowerState_StatusARM</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.PowerState_Status">PowerState_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.AgentPool_Status">AgentPool_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_Status">ManagedClusterAgentPoolProfile_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedCluster_Status">ManagedCluster_Status</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.PowerStateStatusCode">
PowerStateStatusCode
</a>
</em>
</td>
<td>
<p>Code: Tells whether the cluster is Running or Stopped</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.PowerState_StatusARM">PowerState_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileProperties_StatusARM">ManagedClusterAgentPoolProfileProperties_StatusARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_StatusARM">ManagedClusterAgentPoolProfile_StatusARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterProperties_StatusARM">ManagedClusterProperties_StatusARM</a>)
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
<a href="#containerservice.azure.com/v1beta20210501.PowerStateStatusCode">
PowerStateStatusCode
</a>
</em>
</td>
<td>
<p>Code: Tells whether the cluster is Running or Stopped</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.PrivateLinkResource">PrivateLinkResource
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusters_Spec">ManagedClusters_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/PrivateLinkResource">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/PrivateLinkResource</a></p>
</div>
<table>
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
<h3 id="containerservice.azure.com/v1beta20210501.PrivateLinkResourceARM">PrivateLinkResourceARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPropertiesARM">ManagedClusterPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/PrivateLinkResource">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/PrivateLinkResource</a></p>
</div>
<table>
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
<h3 id="containerservice.azure.com/v1beta20210501.PrivateLinkResource_Status">PrivateLinkResource_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedCluster_Status">ManagedCluster_Status</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.PrivateLinkResource_StatusARM">PrivateLinkResource_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterProperties_StatusARM">ManagedClusterProperties_StatusARM</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.ResourceReference">ResourceReference
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile">ManagedClusterLoadBalancerProfile</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfileOutboundIPPrefixes">ManagedClusterLoadBalancerProfileOutboundIPPrefixes</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfileOutboundIPs">ManagedClusterLoadBalancerProfileOutboundIPs</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ResourceReference">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ResourceReference</a></p>
</div>
<table>
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
<h3 id="containerservice.azure.com/v1beta20210501.ResourceReferenceARM">ResourceReferenceARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfileARM">ManagedClusterLoadBalancerProfileARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfileOutboundIPPrefixesARM">ManagedClusterLoadBalancerProfileOutboundIPPrefixesARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfileOutboundIPsARM">ManagedClusterLoadBalancerProfileOutboundIPsARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ResourceReference">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/ResourceReference</a></p>
</div>
<table>
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
<h3 id="containerservice.azure.com/v1beta20210501.ResourceReference_Status">ResourceReference_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_Status">ManagedClusterLoadBalancerProfile_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_Status_OutboundIPPrefixes">ManagedClusterLoadBalancerProfile_Status_OutboundIPPrefixes</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_Status_OutboundIPs">ManagedClusterLoadBalancerProfile_Status_OutboundIPs</a>)
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
<p>Id: The fully qualified Azure resource id.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ResourceReference_StatusARM">ResourceReference_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_StatusARM">ManagedClusterLoadBalancerProfile_StatusARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_Status_OutboundIPPrefixesARM">ManagedClusterLoadBalancerProfile_Status_OutboundIPPrefixesARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterLoadBalancerProfile_Status_OutboundIPsARM">ManagedClusterLoadBalancerProfile_Status_OutboundIPsARM</a>)
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
<p>Id: The fully qualified Azure resource id.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ScaleSetEvictionPolicy_Status">ScaleSetEvictionPolicy_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.AgentPool_Status">AgentPool_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileProperties_StatusARM">ManagedClusterAgentPoolProfileProperties_StatusARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_Status">ManagedClusterAgentPoolProfile_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_StatusARM">ManagedClusterAgentPoolProfile_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Deallocate&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Delete&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.ScaleSetPriority_Status">ScaleSetPriority_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.AgentPool_Status">AgentPool_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfileProperties_StatusARM">ManagedClusterAgentPoolProfileProperties_StatusARM</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_Status">ManagedClusterAgentPoolProfile_Status</a>, <a href="#containerservice.azure.com/v1beta20210501.ManagedClusterAgentPoolProfile_StatusARM">ManagedClusterAgentPoolProfile_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Regular&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Spot&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="containerservice.azure.com/v1beta20210501.SysctlConfig">SysctlConfig
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.LinuxOSConfig">LinuxOSConfig</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/SysctlConfig">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/SysctlConfig</a></p>
</div>
<table>
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
<h3 id="containerservice.azure.com/v1beta20210501.SysctlConfigARM">SysctlConfigARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.LinuxOSConfigARM">LinuxOSConfigARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/SysctlConfig">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/SysctlConfig</a></p>
</div>
<table>
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
<h3 id="containerservice.azure.com/v1beta20210501.SysctlConfig_Status">SysctlConfig_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.LinuxOSConfig_Status">LinuxOSConfig_Status</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.SysctlConfig_StatusARM">SysctlConfig_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.LinuxOSConfig_StatusARM">LinuxOSConfig_StatusARM</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.UserAssignedIdentity">UserAssignedIdentity
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentity">ManagedClusterPodIdentity</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/UserAssignedIdentity">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/UserAssignedIdentity</a></p>
</div>
<table>
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
<h3 id="containerservice.azure.com/v1beta20210501.UserAssignedIdentityARM">UserAssignedIdentityARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentityARM">ManagedClusterPodIdentityARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/UserAssignedIdentity">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.ContainerService.json#/definitions/UserAssignedIdentity</a></p>
</div>
<table>
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
<h3 id="containerservice.azure.com/v1beta20210501.UserAssignedIdentity_Status">UserAssignedIdentity_Status
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentity_Status">ManagedClusterPodIdentity_Status</a>)
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
<h3 id="containerservice.azure.com/v1beta20210501.UserAssignedIdentity_StatusARM">UserAssignedIdentity_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#containerservice.azure.com/v1beta20210501.ManagedClusterPodIdentity_StatusARM">ManagedClusterPodIdentity_StatusARM</a>)
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
<hr/>
