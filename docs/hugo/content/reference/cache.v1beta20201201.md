---
title: cache.azure.com/v1beta20201201
---
<h2 id="cache.azure.com/v1beta20201201">cache.azure.com/v1beta20201201</h2>
<div>
<p>Package v1beta20201201 contains API Schema definitions for the cache v1beta20201201 API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="cache.azure.com/v1beta20201201.PrivateEndpointConnection_Status_SubResourceEmbedded">PrivateEndpointConnection_Status_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisResource_Status">RedisResource_Status</a>)
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
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.PrivateEndpointConnection_Status_SubResourceEmbeddedARM">PrivateEndpointConnection_Status_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisProperties_StatusARM">RedisProperties_StatusARM</a>)
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
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.Redis">Redis
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/resourceDefinitions/redis">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/resourceDefinitions/redis</a></p>
</div>
<table>
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
<a href="#cache.azure.com/v1beta20201201.Redis_Spec">
Redis_Spec
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
<code>enableNonSslPort</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableNonSslPort: Specifies whether the non-ssl Redis server port (6379) is enabled.</p>
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
<code>minimumTlsVersion</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisCreatePropertiesMinimumTlsVersion">
RedisCreatePropertiesMinimumTlsVersion
</a>
</em>
</td>
<td>
<p>MinimumTlsVersion: Optional: requires clients to use a specified TLS version (or higher) to connect (e,g, &lsquo;1.0&rsquo;, &lsquo;1.1&rsquo;,
&lsquo;1.2&rsquo;).</p>
</td>
</tr>
<tr>
<td>
<code>operatorSpec</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisOperatorSpec">
RedisOperatorSpec
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
<code>publicNetworkAccess</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisCreatePropertiesPublicNetworkAccess">
RedisCreatePropertiesPublicNetworkAccess
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Whether or not public endpoint access is allowed for this cache.  Value is optional but if passed
in, must be &lsquo;Enabled&rsquo; or &lsquo;Disabled&rsquo;. If &lsquo;Disabled&rsquo;, private endpoints are the exclusive access method. Default value is
&lsquo;Enabled&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>redisConfiguration</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>RedisConfiguration: All Redis Settings. Few possible keys:
rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value
etc.</p>
</td>
</tr>
<tr>
<td>
<code>redisVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>RedisVersion: Redis version. Only major version will be used in PUT/PATCH request with current valid values: (4, 6)</p>
</td>
</tr>
<tr>
<td>
<code>replicasPerMaster</code><br/>
<em>
int
</em>
</td>
<td>
<p>ReplicasPerMaster: The number of replicas to be created per primary.</p>
</td>
</tr>
<tr>
<td>
<code>replicasPerPrimary</code><br/>
<em>
int
</em>
</td>
<td>
<p>ReplicasPerPrimary: The number of replicas to be created per primary.</p>
</td>
</tr>
<tr>
<td>
<code>shardCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>ShardCount: The number of shards to be created on a Premium Cluster Cache.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.Sku">
Sku
</a>
</em>
</td>
<td>
<p>Sku: SKU parameters supplied to the create Redis operation.</p>
</td>
</tr>
<tr>
<td>
<code>staticIP</code><br/>
<em>
string
</em>
</td>
<td>
<p>StaticIP: Static IP address. Optionally, may be specified when deploying a Redis cache inside an existing Azure Virtual
Network; auto assigned by default.</p>
</td>
</tr>
<tr>
<td>
<code>subnetReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>SubnetReference: The full resource ID of a subnet in a virtual network to deploy the Redis cache in. Example format:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;Microsoft.{Network|ClassicNetwork}/&#x200b;VirtualNetworks/&#x200b;vnet1/&#x200b;subnets/&#x200b;subnet1</&#x200b;p>
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
<code>tenantSettings</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>TenantSettings: A dictionary of tenant settings</p>
</td>
</tr>
<tr>
<td>
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: A list of availability zones denoting where the resource needs to come from.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisResource_Status">
RedisResource_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.RedisCreatePropertiesARM">RedisCreatePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.Redis_SpecARM">Redis_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/definitions/RedisCreateProperties">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/definitions/RedisCreateProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>enableNonSslPort</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableNonSslPort: Specifies whether the non-ssl Redis server port (6379) is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>minimumTlsVersion</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisCreatePropertiesMinimumTlsVersion">
RedisCreatePropertiesMinimumTlsVersion
</a>
</em>
</td>
<td>
<p>MinimumTlsVersion: Optional: requires clients to use a specified TLS version (or higher) to connect (e,g, &lsquo;1.0&rsquo;, &lsquo;1.1&rsquo;,
&lsquo;1.2&rsquo;).</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisCreatePropertiesPublicNetworkAccess">
RedisCreatePropertiesPublicNetworkAccess
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Whether or not public endpoint access is allowed for this cache.  Value is optional but if passed
in, must be &lsquo;Enabled&rsquo; or &lsquo;Disabled&rsquo;. If &lsquo;Disabled&rsquo;, private endpoints are the exclusive access method. Default value is
&lsquo;Enabled&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>redisConfiguration</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>RedisConfiguration: All Redis Settings. Few possible keys:
rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value
etc.</p>
</td>
</tr>
<tr>
<td>
<code>redisVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>RedisVersion: Redis version. Only major version will be used in PUT/PATCH request with current valid values: (4, 6)</p>
</td>
</tr>
<tr>
<td>
<code>replicasPerMaster</code><br/>
<em>
int
</em>
</td>
<td>
<p>ReplicasPerMaster: The number of replicas to be created per primary.</p>
</td>
</tr>
<tr>
<td>
<code>replicasPerPrimary</code><br/>
<em>
int
</em>
</td>
<td>
<p>ReplicasPerPrimary: The number of replicas to be created per primary.</p>
</td>
</tr>
<tr>
<td>
<code>shardCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>ShardCount: The number of shards to be created on a Premium Cluster Cache.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.SkuARM">
SkuARM
</a>
</em>
</td>
<td>
<p>Sku: SKU parameters supplied to the create Redis operation.</p>
</td>
</tr>
<tr>
<td>
<code>staticIP</code><br/>
<em>
string
</em>
</td>
<td>
<p>StaticIP: Static IP address. Optionally, may be specified when deploying a Redis cache inside an existing Azure Virtual
Network; auto assigned by default.</p>
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
</td>
</tr>
<tr>
<td>
<code>tenantSettings</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>TenantSettings: A dictionary of tenant settings</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.RedisCreatePropertiesMinimumTlsVersion">RedisCreatePropertiesMinimumTlsVersion
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisCreatePropertiesARM">RedisCreatePropertiesARM</a>, <a href="#cache.azure.com/v1beta20201201.Redis_Spec">Redis_Spec</a>)
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
<tbody><tr><td><p>&#34;1.0&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;1.1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;1.2&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.RedisCreatePropertiesPublicNetworkAccess">RedisCreatePropertiesPublicNetworkAccess
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisCreatePropertiesARM">RedisCreatePropertiesARM</a>, <a href="#cache.azure.com/v1beta20201201.Redis_Spec">Redis_Spec</a>)
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
<h3 id="cache.azure.com/v1beta20201201.RedisFirewallRule">RedisFirewallRule
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/resourceDefinitions/redis_firewallRules">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/resourceDefinitions/redis_firewallRules</a></p>
</div>
<table>
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
<a href="#cache.azure.com/v1beta20201201.RedisFirewallRules_Spec">
RedisFirewallRules_Spec
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
<code>endIP</code><br/>
<em>
string
</em>
</td>
<td>
<p>EndIP: highest IP address included in the range</p>
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
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a cache.azure.com/Redis resource</p>
</td>
</tr>
<tr>
<td>
<code>startIP</code><br/>
<em>
string
</em>
</td>
<td>
<p>StartIP: lowest IP address included in the range</p>
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
<a href="#cache.azure.com/v1beta20201201.RedisFirewallRule_Status">
RedisFirewallRule_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.RedisFirewallRulePropertiesARM">RedisFirewallRulePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisFirewallRules_SpecARM">RedisFirewallRules_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/definitions/RedisFirewallRuleProperties">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/definitions/RedisFirewallRuleProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>endIP</code><br/>
<em>
string
</em>
</td>
<td>
<p>EndIP: highest IP address included in the range</p>
</td>
</tr>
<tr>
<td>
<code>startIP</code><br/>
<em>
string
</em>
</td>
<td>
<p>StartIP: lowest IP address included in the range</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.RedisFirewallRuleProperties_StatusARM">RedisFirewallRuleProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisFirewallRule_StatusARM">RedisFirewallRule_StatusARM</a>)
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
<code>endIP</code><br/>
<em>
string
</em>
</td>
<td>
<p>EndIP: highest IP address included in the range</p>
</td>
</tr>
<tr>
<td>
<code>startIP</code><br/>
<em>
string
</em>
</td>
<td>
<p>StartIP: lowest IP address included in the range</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.RedisFirewallRule_Status">RedisFirewallRule_Status
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisFirewallRule">RedisFirewallRule</a>)
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
<code>endIP</code><br/>
<em>
string
</em>
</td>
<td>
<p>EndIP: highest IP address included in the range</p>
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
<code>startIP</code><br/>
<em>
string
</em>
</td>
<td>
<p>StartIP: lowest IP address included in the range</p>
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
<h3 id="cache.azure.com/v1beta20201201.RedisFirewallRule_StatusARM">RedisFirewallRule_StatusARM
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
<a href="#cache.azure.com/v1beta20201201.RedisFirewallRuleProperties_StatusARM">
RedisFirewallRuleProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: redis cache firewall rule properties</p>
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
<h3 id="cache.azure.com/v1beta20201201.RedisFirewallRulesSpecAPIVersion">RedisFirewallRulesSpecAPIVersion
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
<tbody><tr><td><p>&#34;2020-12-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.RedisFirewallRules_Spec">RedisFirewallRules_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisFirewallRule">RedisFirewallRule</a>)
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
<code>endIP</code><br/>
<em>
string
</em>
</td>
<td>
<p>EndIP: highest IP address included in the range</p>
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
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a cache.azure.com/Redis resource</p>
</td>
</tr>
<tr>
<td>
<code>startIP</code><br/>
<em>
string
</em>
</td>
<td>
<p>StartIP: lowest IP address included in the range</p>
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
<h3 id="cache.azure.com/v1beta20201201.RedisFirewallRules_SpecARM">RedisFirewallRules_SpecARM
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
<p>Name: The name of the firewall rule.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisFirewallRulePropertiesARM">
RedisFirewallRulePropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Specifies a range of IP addresses permitted to connect to the cache</p>
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
<h3 id="cache.azure.com/v1beta20201201.RedisInstanceDetails_Status">RedisInstanceDetails_Status
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisResource_Status">RedisResource_Status</a>)
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
<code>isMaster</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsMaster: Specifies whether the instance is a primary node.</p>
</td>
</tr>
<tr>
<td>
<code>isPrimary</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsPrimary: Specifies whether the instance is a primary node.</p>
</td>
</tr>
<tr>
<td>
<code>nonSslPort</code><br/>
<em>
int
</em>
</td>
<td>
<p>NonSslPort: If enableNonSslPort is true, provides Redis instance Non-SSL port.</p>
</td>
</tr>
<tr>
<td>
<code>shardId</code><br/>
<em>
int
</em>
</td>
<td>
<p>ShardId: If clustering is enabled, the Shard ID of Redis Instance</p>
</td>
</tr>
<tr>
<td>
<code>sslPort</code><br/>
<em>
int
</em>
</td>
<td>
<p>SslPort: Redis instance SSL port.</p>
</td>
</tr>
<tr>
<td>
<code>zone</code><br/>
<em>
string
</em>
</td>
<td>
<p>Zone: If the Cache uses availability zones, specifies availability zone where this instance is located.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.RedisInstanceDetails_StatusARM">RedisInstanceDetails_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisProperties_StatusARM">RedisProperties_StatusARM</a>)
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
<code>isMaster</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsMaster: Specifies whether the instance is a primary node.</p>
</td>
</tr>
<tr>
<td>
<code>isPrimary</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsPrimary: Specifies whether the instance is a primary node.</p>
</td>
</tr>
<tr>
<td>
<code>nonSslPort</code><br/>
<em>
int
</em>
</td>
<td>
<p>NonSslPort: If enableNonSslPort is true, provides Redis instance Non-SSL port.</p>
</td>
</tr>
<tr>
<td>
<code>shardId</code><br/>
<em>
int
</em>
</td>
<td>
<p>ShardId: If clustering is enabled, the Shard ID of Redis Instance</p>
</td>
</tr>
<tr>
<td>
<code>sslPort</code><br/>
<em>
int
</em>
</td>
<td>
<p>SslPort: Redis instance SSL port.</p>
</td>
</tr>
<tr>
<td>
<code>zone</code><br/>
<em>
string
</em>
</td>
<td>
<p>Zone: If the Cache uses availability zones, specifies availability zone where this instance is located.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.RedisLinkedServer">RedisLinkedServer
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/resourceDefinitions/redis_linkedServers">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/resourceDefinitions/redis_linkedServers</a></p>
</div>
<table>
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
<a href="#cache.azure.com/v1beta20201201.RedisLinkedServers_Spec">
RedisLinkedServers_Spec
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
<code>linkedRedisCacheLocation</code><br/>
<em>
string
</em>
</td>
<td>
<p>LinkedRedisCacheLocation: Location of the linked redis cache.</p>
</td>
</tr>
<tr>
<td>
<code>linkedRedisCacheReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>LinkedRedisCacheReference: Fully qualified resourceId of the linked redis cache.</p>
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
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a cache.azure.com/Redis resource</p>
</td>
</tr>
<tr>
<td>
<code>serverRole</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisLinkedServerCreatePropertiesServerRole">
RedisLinkedServerCreatePropertiesServerRole
</a>
</em>
</td>
<td>
<p>ServerRole: Role of the linked server.</p>
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
<a href="#cache.azure.com/v1beta20201201.RedisLinkedServerWithProperties_Status">
RedisLinkedServerWithProperties_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.RedisLinkedServerCreatePropertiesARM">RedisLinkedServerCreatePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisLinkedServers_SpecARM">RedisLinkedServers_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/definitions/RedisLinkedServerCreateProperties">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/definitions/RedisLinkedServerCreateProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>linkedRedisCacheId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>linkedRedisCacheLocation</code><br/>
<em>
string
</em>
</td>
<td>
<p>LinkedRedisCacheLocation: Location of the linked redis cache.</p>
</td>
</tr>
<tr>
<td>
<code>serverRole</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisLinkedServerCreatePropertiesServerRole">
RedisLinkedServerCreatePropertiesServerRole
</a>
</em>
</td>
<td>
<p>ServerRole: Role of the linked server.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.RedisLinkedServerCreatePropertiesServerRole">RedisLinkedServerCreatePropertiesServerRole
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisLinkedServerCreatePropertiesARM">RedisLinkedServerCreatePropertiesARM</a>, <a href="#cache.azure.com/v1beta20201201.RedisLinkedServers_Spec">RedisLinkedServers_Spec</a>)
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
<tbody><tr><td><p>&#34;Primary&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Secondary&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.RedisLinkedServerPropertiesStatusServerRole">RedisLinkedServerPropertiesStatusServerRole
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisLinkedServerProperties_StatusARM">RedisLinkedServerProperties_StatusARM</a>, <a href="#cache.azure.com/v1beta20201201.RedisLinkedServerWithProperties_Status">RedisLinkedServerWithProperties_Status</a>)
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
<tbody><tr><td><p>&#34;Primary&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Secondary&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.RedisLinkedServerProperties_StatusARM">RedisLinkedServerProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisLinkedServerWithProperties_StatusARM">RedisLinkedServerWithProperties_StatusARM</a>)
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
<code>linkedRedisCacheId</code><br/>
<em>
string
</em>
</td>
<td>
<p>LinkedRedisCacheId: Fully qualified resourceId of the linked redis cache.</p>
</td>
</tr>
<tr>
<td>
<code>linkedRedisCacheLocation</code><br/>
<em>
string
</em>
</td>
<td>
<p>LinkedRedisCacheLocation: Location of the linked redis cache.</p>
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
<p>ProvisioningState: Terminal state of the link between primary and secondary redis cache.</p>
</td>
</tr>
<tr>
<td>
<code>serverRole</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisLinkedServerPropertiesStatusServerRole">
RedisLinkedServerPropertiesStatusServerRole
</a>
</em>
</td>
<td>
<p>ServerRole: Role of the linked server.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.RedisLinkedServerWithProperties_Status">RedisLinkedServerWithProperties_Status
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisLinkedServer">RedisLinkedServer</a>)
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
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>linkedRedisCacheId</code><br/>
<em>
string
</em>
</td>
<td>
<p>LinkedRedisCacheId: Fully qualified resourceId of the linked redis cache.</p>
</td>
</tr>
<tr>
<td>
<code>linkedRedisCacheLocation</code><br/>
<em>
string
</em>
</td>
<td>
<p>LinkedRedisCacheLocation: Location of the linked redis cache.</p>
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
string
</em>
</td>
<td>
<p>ProvisioningState: Terminal state of the link between primary and secondary redis cache.</p>
</td>
</tr>
<tr>
<td>
<code>serverRole</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisLinkedServerPropertiesStatusServerRole">
RedisLinkedServerPropertiesStatusServerRole
</a>
</em>
</td>
<td>
<p>ServerRole: Role of the linked server.</p>
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
<h3 id="cache.azure.com/v1beta20201201.RedisLinkedServerWithProperties_StatusARM">RedisLinkedServerWithProperties_StatusARM
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
<a href="#cache.azure.com/v1beta20201201.RedisLinkedServerProperties_StatusARM">
RedisLinkedServerProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the linked server.</p>
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
<h3 id="cache.azure.com/v1beta20201201.RedisLinkedServer_Status">RedisLinkedServer_Status
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisResource_Status">RedisResource_Status</a>)
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
<p>Id: Linked server Id.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.RedisLinkedServer_StatusARM">RedisLinkedServer_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisProperties_StatusARM">RedisProperties_StatusARM</a>)
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
<p>Id: Linked server Id.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.RedisLinkedServersSpecAPIVersion">RedisLinkedServersSpecAPIVersion
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
<tbody><tr><td><p>&#34;2020-12-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.RedisLinkedServers_Spec">RedisLinkedServers_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisLinkedServer">RedisLinkedServer</a>)
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
<code>linkedRedisCacheLocation</code><br/>
<em>
string
</em>
</td>
<td>
<p>LinkedRedisCacheLocation: Location of the linked redis cache.</p>
</td>
</tr>
<tr>
<td>
<code>linkedRedisCacheReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>LinkedRedisCacheReference: Fully qualified resourceId of the linked redis cache.</p>
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
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a cache.azure.com/Redis resource</p>
</td>
</tr>
<tr>
<td>
<code>serverRole</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisLinkedServerCreatePropertiesServerRole">
RedisLinkedServerCreatePropertiesServerRole
</a>
</em>
</td>
<td>
<p>ServerRole: Role of the linked server.</p>
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
<h3 id="cache.azure.com/v1beta20201201.RedisLinkedServers_SpecARM">RedisLinkedServers_SpecARM
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
<p>Name: The name of the linked server that is being added to the Redis cache.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisLinkedServerCreatePropertiesARM">
RedisLinkedServerCreatePropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Create properties for a linked server</p>
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
<h3 id="cache.azure.com/v1beta20201201.RedisOperatorSecrets">RedisOperatorSecrets
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisOperatorSpec">RedisOperatorSpec</a>)
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
<code>hostName</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretDestination">
genruntime.SecretDestination
</a>
</em>
</td>
<td>
<p>HostName: indicates where the HostName secret should be placed. If omitted, the secret will not be retrieved from Azure.</p>
</td>
</tr>
<tr>
<td>
<code>port</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretDestination">
genruntime.SecretDestination
</a>
</em>
</td>
<td>
<p>Port: indicates where the Port secret should be placed. If omitted, the secret will not be retrieved from Azure.</p>
</td>
</tr>
<tr>
<td>
<code>primaryKey</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretDestination">
genruntime.SecretDestination
</a>
</em>
</td>
<td>
<p>PrimaryKey: indicates where the PrimaryKey secret should be placed. If omitted, the secret will not be retrieved from
Azure.</p>
</td>
</tr>
<tr>
<td>
<code>sslPort</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretDestination">
genruntime.SecretDestination
</a>
</em>
</td>
<td>
<p>SSLPort: indicates where the SSLPort secret should be placed. If omitted, the secret will not be retrieved from Azure.</p>
</td>
</tr>
<tr>
<td>
<code>secondaryKey</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretDestination">
genruntime.SecretDestination
</a>
</em>
</td>
<td>
<p>SecondaryKey: indicates where the SecondaryKey secret should be placed. If omitted, the secret will not be retrieved
from Azure.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.RedisOperatorSpec">RedisOperatorSpec
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.Redis_Spec">Redis_Spec</a>)
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
<a href="#cache.azure.com/v1beta20201201.RedisOperatorSecrets">
RedisOperatorSecrets
</a>
</em>
</td>
<td>
<p>Secrets: configures where to place Azure generated secrets.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.RedisPatchSchedule">RedisPatchSchedule
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/resourceDefinitions/redis_patchSchedules">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/resourceDefinitions/redis_patchSchedules</a></p>
</div>
<table>
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
<a href="#cache.azure.com/v1beta20201201.RedisPatchSchedules_Spec">
RedisPatchSchedules_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
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
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a cache.azure.com/Redis resource</p>
</td>
</tr>
<tr>
<td>
<code>scheduleEntries</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.ScheduleEntry">
[]ScheduleEntry
</a>
</em>
</td>
<td>
<p>ScheduleEntries: List of patch schedules for a Redis cache.</p>
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
<a href="#cache.azure.com/v1beta20201201.RedisPatchSchedule_Status">
RedisPatchSchedule_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.RedisPatchSchedule_Status">RedisPatchSchedule_Status
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisPatchSchedule">RedisPatchSchedule</a>)
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
<code>scheduleEntries</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.ScheduleEntry_Status">
[]ScheduleEntry_Status
</a>
</em>
</td>
<td>
<p>ScheduleEntries: List of patch schedules for a Redis cache.</p>
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
<h3 id="cache.azure.com/v1beta20201201.RedisPatchSchedule_StatusARM">RedisPatchSchedule_StatusARM
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
<a href="#cache.azure.com/v1beta20201201.ScheduleEntries_StatusARM">
ScheduleEntries_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: List of patch schedules for a Redis cache.</p>
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
<h3 id="cache.azure.com/v1beta20201201.RedisPatchSchedulesSpecAPIVersion">RedisPatchSchedulesSpecAPIVersion
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
<tbody><tr><td><p>&#34;2020-12-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.RedisPatchSchedules_Spec">RedisPatchSchedules_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisPatchSchedule">RedisPatchSchedule</a>)
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
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a cache.azure.com/Redis resource</p>
</td>
</tr>
<tr>
<td>
<code>scheduleEntries</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.ScheduleEntry">
[]ScheduleEntry
</a>
</em>
</td>
<td>
<p>ScheduleEntries: List of patch schedules for a Redis cache.</p>
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
<h3 id="cache.azure.com/v1beta20201201.RedisPatchSchedules_SpecARM">RedisPatchSchedules_SpecARM
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
<p>Name: Default string modeled as parameter for auto generation to work correctly.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.ScheduleEntriesARM">
ScheduleEntriesARM
</a>
</em>
</td>
<td>
<p>Properties: List of patch schedules for a Redis cache.</p>
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
<h3 id="cache.azure.com/v1beta20201201.RedisPropertiesStatusMinimumTlsVersion">RedisPropertiesStatusMinimumTlsVersion
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisProperties_StatusARM">RedisProperties_StatusARM</a>, <a href="#cache.azure.com/v1beta20201201.RedisResource_Status">RedisResource_Status</a>)
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
<tbody><tr><td><p>&#34;1.0&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;1.1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;1.2&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.RedisPropertiesStatusProvisioningState">RedisPropertiesStatusProvisioningState
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisProperties_StatusARM">RedisProperties_StatusARM</a>, <a href="#cache.azure.com/v1beta20201201.RedisResource_Status">RedisResource_Status</a>)
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
</tr><tr><td><p>&#34;Disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Failed&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Linking&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Provisioning&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;RecoveringScaleFailure&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Scaling&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Succeeded&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Unlinking&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Unprovisioning&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Updating&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.RedisPropertiesStatusPublicNetworkAccess">RedisPropertiesStatusPublicNetworkAccess
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisProperties_StatusARM">RedisProperties_StatusARM</a>, <a href="#cache.azure.com/v1beta20201201.RedisResource_Status">RedisResource_Status</a>)
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
<h3 id="cache.azure.com/v1beta20201201.RedisProperties_StatusARM">RedisProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisResource_StatusARM">RedisResource_StatusARM</a>)
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
<code>enableNonSslPort</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableNonSslPort: Specifies whether the non-ssl Redis server port (6379) is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>hostName</code><br/>
<em>
string
</em>
</td>
<td>
<p>HostName: Redis host name.</p>
</td>
</tr>
<tr>
<td>
<code>instances</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisInstanceDetails_StatusARM">
[]RedisInstanceDetails_StatusARM
</a>
</em>
</td>
<td>
<p>Instances: List of the Redis instances associated with the cache</p>
</td>
</tr>
<tr>
<td>
<code>linkedServers</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisLinkedServer_StatusARM">
[]RedisLinkedServer_StatusARM
</a>
</em>
</td>
<td>
<p>LinkedServers: List of the linked servers associated with the cache</p>
</td>
</tr>
<tr>
<td>
<code>minimumTlsVersion</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisPropertiesStatusMinimumTlsVersion">
RedisPropertiesStatusMinimumTlsVersion
</a>
</em>
</td>
<td>
<p>MinimumTlsVersion: Optional: requires clients to use a specified TLS version (or higher) to connect (e,g, &lsquo;1.0&rsquo;, &lsquo;1.1&rsquo;,
&lsquo;1.2&rsquo;)</p>
</td>
</tr>
<tr>
<td>
<code>port</code><br/>
<em>
int
</em>
</td>
<td>
<p>Port: Redis non-SSL port.</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpointConnections</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.PrivateEndpointConnection_Status_SubResourceEmbeddedARM">
[]PrivateEndpointConnection_Status_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>PrivateEndpointConnections: List of private endpoint connection associated with the specified redis cache</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisPropertiesStatusProvisioningState">
RedisPropertiesStatusProvisioningState
</a>
</em>
</td>
<td>
<p>ProvisioningState: Redis instance provisioning status.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisPropertiesStatusPublicNetworkAccess">
RedisPropertiesStatusPublicNetworkAccess
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Whether or not public endpoint access is allowed for this cache.  Value is optional but if passed
in, must be &lsquo;Enabled&rsquo; or &lsquo;Disabled&rsquo;. If &lsquo;Disabled&rsquo;, private endpoints are the exclusive access method. Default value is
&lsquo;Enabled&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>redisConfiguration</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>RedisConfiguration: All Redis Settings. Few possible keys:
rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value
etc.</p>
</td>
</tr>
<tr>
<td>
<code>redisVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>RedisVersion: Redis version. Only major version will be used in PUT/PATCH request with current valid values: (4, 6)</p>
</td>
</tr>
<tr>
<td>
<code>replicasPerMaster</code><br/>
<em>
int
</em>
</td>
<td>
<p>ReplicasPerMaster: The number of replicas to be created per primary.</p>
</td>
</tr>
<tr>
<td>
<code>replicasPerPrimary</code><br/>
<em>
int
</em>
</td>
<td>
<p>ReplicasPerPrimary: The number of replicas to be created per primary.</p>
</td>
</tr>
<tr>
<td>
<code>shardCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>ShardCount: The number of shards to be created on a Premium Cluster Cache.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.Sku_StatusARM">
Sku_StatusARM
</a>
</em>
</td>
<td>
<p>Sku: The SKU of the Redis cache to deploy.</p>
</td>
</tr>
<tr>
<td>
<code>sslPort</code><br/>
<em>
int
</em>
</td>
<td>
<p>SslPort: Redis SSL port.</p>
</td>
</tr>
<tr>
<td>
<code>staticIP</code><br/>
<em>
string
</em>
</td>
<td>
<p>StaticIP: Static IP address. Optionally, may be specified when deploying a Redis cache inside an existing Azure Virtual
Network; auto assigned by default.</p>
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
<p>SubnetId: The full resource ID of a subnet in a virtual network to deploy the Redis cache in. Example format:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;Microsoft.{Network|ClassicNetwork}/&#x200b;VirtualNetworks/&#x200b;vnet1/&#x200b;subnets/&#x200b;subnet1</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>tenantSettings</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>TenantSettings: A dictionary of tenant settings</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.RedisResource_Status">RedisResource_Status
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.Redis">Redis</a>)
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
<code>enableNonSslPort</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableNonSslPort: Specifies whether the non-ssl Redis server port (6379) is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>hostName</code><br/>
<em>
string
</em>
</td>
<td>
<p>HostName: Redis host name.</p>
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
<code>instances</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisInstanceDetails_Status">
[]RedisInstanceDetails_Status
</a>
</em>
</td>
<td>
<p>Instances: List of the Redis instances associated with the cache</p>
</td>
</tr>
<tr>
<td>
<code>linkedServers</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisLinkedServer_Status">
[]RedisLinkedServer_Status
</a>
</em>
</td>
<td>
<p>LinkedServers: List of the linked servers associated with the cache</p>
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
<code>minimumTlsVersion</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisPropertiesStatusMinimumTlsVersion">
RedisPropertiesStatusMinimumTlsVersion
</a>
</em>
</td>
<td>
<p>MinimumTlsVersion: Optional: requires clients to use a specified TLS version (or higher) to connect (e,g, &lsquo;1.0&rsquo;, &lsquo;1.1&rsquo;,
&lsquo;1.2&rsquo;)</p>
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
<code>port</code><br/>
<em>
int
</em>
</td>
<td>
<p>Port: Redis non-SSL port.</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpointConnections</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.PrivateEndpointConnection_Status_SubResourceEmbedded">
[]PrivateEndpointConnection_Status_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>PrivateEndpointConnections: List of private endpoint connection associated with the specified redis cache</p>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisPropertiesStatusProvisioningState">
RedisPropertiesStatusProvisioningState
</a>
</em>
</td>
<td>
<p>ProvisioningState: Redis instance provisioning status.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisPropertiesStatusPublicNetworkAccess">
RedisPropertiesStatusPublicNetworkAccess
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Whether or not public endpoint access is allowed for this cache.  Value is optional but if passed
in, must be &lsquo;Enabled&rsquo; or &lsquo;Disabled&rsquo;. If &lsquo;Disabled&rsquo;, private endpoints are the exclusive access method. Default value is
&lsquo;Enabled&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>redisConfiguration</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>RedisConfiguration: All Redis Settings. Few possible keys:
rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value
etc.</p>
</td>
</tr>
<tr>
<td>
<code>redisVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>RedisVersion: Redis version. Only major version will be used in PUT/PATCH request with current valid values: (4, 6)</p>
</td>
</tr>
<tr>
<td>
<code>replicasPerMaster</code><br/>
<em>
int
</em>
</td>
<td>
<p>ReplicasPerMaster: The number of replicas to be created per primary.</p>
</td>
</tr>
<tr>
<td>
<code>replicasPerPrimary</code><br/>
<em>
int
</em>
</td>
<td>
<p>ReplicasPerPrimary: The number of replicas to be created per primary.</p>
</td>
</tr>
<tr>
<td>
<code>shardCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>ShardCount: The number of shards to be created on a Premium Cluster Cache.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.Sku_Status">
Sku_Status
</a>
</em>
</td>
<td>
<p>Sku: The SKU of the Redis cache to deploy.</p>
</td>
</tr>
<tr>
<td>
<code>sslPort</code><br/>
<em>
int
</em>
</td>
<td>
<p>SslPort: Redis SSL port.</p>
</td>
</tr>
<tr>
<td>
<code>staticIP</code><br/>
<em>
string
</em>
</td>
<td>
<p>StaticIP: Static IP address. Optionally, may be specified when deploying a Redis cache inside an existing Azure Virtual
Network; auto assigned by default.</p>
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
<p>SubnetId: The full resource ID of a subnet in a virtual network to deploy the Redis cache in. Example format:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;Microsoft.{Network|ClassicNetwork}/&#x200b;VirtualNetworks/&#x200b;vnet1/&#x200b;subnets/&#x200b;subnet1</&#x200b;p>
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
<code>tenantSettings</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>TenantSettings: A dictionary of tenant settings</p>
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
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: A list of availability zones denoting where the resource needs to come from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.RedisResource_StatusARM">RedisResource_StatusARM
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
<a href="#cache.azure.com/v1beta20201201.RedisProperties_StatusARM">
RedisProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Redis cache properties.</p>
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
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: A list of availability zones denoting where the resource needs to come from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.RedisSpecAPIVersion">RedisSpecAPIVersion
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
<tbody><tr><td><p>&#34;2020-12-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.Redis_Spec">Redis_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.Redis">Redis</a>)
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
<code>enableNonSslPort</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableNonSslPort: Specifies whether the non-ssl Redis server port (6379) is enabled.</p>
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
<code>minimumTlsVersion</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisCreatePropertiesMinimumTlsVersion">
RedisCreatePropertiesMinimumTlsVersion
</a>
</em>
</td>
<td>
<p>MinimumTlsVersion: Optional: requires clients to use a specified TLS version (or higher) to connect (e,g, &lsquo;1.0&rsquo;, &lsquo;1.1&rsquo;,
&lsquo;1.2&rsquo;).</p>
</td>
</tr>
<tr>
<td>
<code>operatorSpec</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisOperatorSpec">
RedisOperatorSpec
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
<code>publicNetworkAccess</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisCreatePropertiesPublicNetworkAccess">
RedisCreatePropertiesPublicNetworkAccess
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Whether or not public endpoint access is allowed for this cache.  Value is optional but if passed
in, must be &lsquo;Enabled&rsquo; or &lsquo;Disabled&rsquo;. If &lsquo;Disabled&rsquo;, private endpoints are the exclusive access method. Default value is
&lsquo;Enabled&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>redisConfiguration</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>RedisConfiguration: All Redis Settings. Few possible keys:
rdb-backup-enabled,rdb-storage-connection-string,rdb-backup-frequency,maxmemory-delta,maxmemory-policy,notify-keyspace-events,maxmemory-samples,slowlog-log-slower-than,slowlog-max-len,list-max-ziplist-entries,list-max-ziplist-value,hash-max-ziplist-entries,hash-max-ziplist-value,set-max-intset-entries,zset-max-ziplist-entries,zset-max-ziplist-value
etc.</p>
</td>
</tr>
<tr>
<td>
<code>redisVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>RedisVersion: Redis version. Only major version will be used in PUT/PATCH request with current valid values: (4, 6)</p>
</td>
</tr>
<tr>
<td>
<code>replicasPerMaster</code><br/>
<em>
int
</em>
</td>
<td>
<p>ReplicasPerMaster: The number of replicas to be created per primary.</p>
</td>
</tr>
<tr>
<td>
<code>replicasPerPrimary</code><br/>
<em>
int
</em>
</td>
<td>
<p>ReplicasPerPrimary: The number of replicas to be created per primary.</p>
</td>
</tr>
<tr>
<td>
<code>shardCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>ShardCount: The number of shards to be created on a Premium Cluster Cache.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.Sku">
Sku
</a>
</em>
</td>
<td>
<p>Sku: SKU parameters supplied to the create Redis operation.</p>
</td>
</tr>
<tr>
<td>
<code>staticIP</code><br/>
<em>
string
</em>
</td>
<td>
<p>StaticIP: Static IP address. Optionally, may be specified when deploying a Redis cache inside an existing Azure Virtual
Network; auto assigned by default.</p>
</td>
</tr>
<tr>
<td>
<code>subnetReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>SubnetReference: The full resource ID of a subnet in a virtual network to deploy the Redis cache in. Example format:
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;Microsoft.{Network|ClassicNetwork}/&#x200b;VirtualNetworks/&#x200b;vnet1/&#x200b;subnets/&#x200b;subnet1</&#x200b;p>
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
<code>tenantSettings</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>TenantSettings: A dictionary of tenant settings</p>
</td>
</tr>
<tr>
<td>
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: A list of availability zones denoting where the resource needs to come from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.Redis_SpecARM">Redis_SpecARM
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
<p>Name: The name of the Redis cache.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.RedisCreatePropertiesARM">
RedisCreatePropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties supplied to Create Redis operation.</p>
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
<code>zones</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Zones: A list of availability zones denoting where the resource needs to come from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.ScheduleEntriesARM">ScheduleEntriesARM
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisPatchSchedules_SpecARM">RedisPatchSchedules_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/definitions/ScheduleEntries">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/definitions/ScheduleEntries</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>scheduleEntries</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.ScheduleEntryARM">
[]ScheduleEntryARM
</a>
</em>
</td>
<td>
<p>ScheduleEntries: List of patch schedules for a Redis cache.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.ScheduleEntries_StatusARM">ScheduleEntries_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisPatchSchedule_StatusARM">RedisPatchSchedule_StatusARM</a>)
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
<code>scheduleEntries</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.ScheduleEntry_StatusARM">
[]ScheduleEntry_StatusARM
</a>
</em>
</td>
<td>
<p>ScheduleEntries: List of patch schedules for a Redis cache.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.ScheduleEntry">ScheduleEntry
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisPatchSchedules_Spec">RedisPatchSchedules_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/definitions/ScheduleEntry">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/definitions/ScheduleEntry</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dayOfWeek</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.ScheduleEntryDayOfWeek">
ScheduleEntryDayOfWeek
</a>
</em>
</td>
<td>
<p>DayOfWeek: Day of the week when a cache can be patched.</p>
</td>
</tr>
<tr>
<td>
<code>maintenanceWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaintenanceWindow: ISO8601 timespan specifying how much time cache patching can take.</p>
</td>
</tr>
<tr>
<td>
<code>startHourUtc</code><br/>
<em>
int
</em>
</td>
<td>
<p>StartHourUtc: Start hour after which cache patching can start.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.ScheduleEntryARM">ScheduleEntryARM
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.ScheduleEntriesARM">ScheduleEntriesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/definitions/ScheduleEntry">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/definitions/ScheduleEntry</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>dayOfWeek</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.ScheduleEntryDayOfWeek">
ScheduleEntryDayOfWeek
</a>
</em>
</td>
<td>
<p>DayOfWeek: Day of the week when a cache can be patched.</p>
</td>
</tr>
<tr>
<td>
<code>maintenanceWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaintenanceWindow: ISO8601 timespan specifying how much time cache patching can take.</p>
</td>
</tr>
<tr>
<td>
<code>startHourUtc</code><br/>
<em>
int
</em>
</td>
<td>
<p>StartHourUtc: Start hour after which cache patching can start.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.ScheduleEntryDayOfWeek">ScheduleEntryDayOfWeek
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.ScheduleEntry">ScheduleEntry</a>, <a href="#cache.azure.com/v1beta20201201.ScheduleEntryARM">ScheduleEntryARM</a>)
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
<tbody><tr><td><p>&#34;Everyday&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Friday&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Monday&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Saturday&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Sunday&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Thursday&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Tuesday&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Wednesday&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Weekend&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.ScheduleEntryStatusDayOfWeek">ScheduleEntryStatusDayOfWeek
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.ScheduleEntry_Status">ScheduleEntry_Status</a>, <a href="#cache.azure.com/v1beta20201201.ScheduleEntry_StatusARM">ScheduleEntry_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Everyday&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Friday&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Monday&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Saturday&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Sunday&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Thursday&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Tuesday&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Wednesday&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Weekend&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.ScheduleEntry_Status">ScheduleEntry_Status
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisPatchSchedule_Status">RedisPatchSchedule_Status</a>)
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
<code>dayOfWeek</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.ScheduleEntryStatusDayOfWeek">
ScheduleEntryStatusDayOfWeek
</a>
</em>
</td>
<td>
<p>DayOfWeek: Day of the week when a cache can be patched.</p>
</td>
</tr>
<tr>
<td>
<code>maintenanceWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaintenanceWindow: ISO8601 timespan specifying how much time cache patching can take.</p>
</td>
</tr>
<tr>
<td>
<code>startHourUtc</code><br/>
<em>
int
</em>
</td>
<td>
<p>StartHourUtc: Start hour after which cache patching can start.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.ScheduleEntry_StatusARM">ScheduleEntry_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.ScheduleEntries_StatusARM">ScheduleEntries_StatusARM</a>)
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
<code>dayOfWeek</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.ScheduleEntryStatusDayOfWeek">
ScheduleEntryStatusDayOfWeek
</a>
</em>
</td>
<td>
<p>DayOfWeek: Day of the week when a cache can be patched.</p>
</td>
</tr>
<tr>
<td>
<code>maintenanceWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>MaintenanceWindow: ISO8601 timespan specifying how much time cache patching can take.</p>
</td>
</tr>
<tr>
<td>
<code>startHourUtc</code><br/>
<em>
int
</em>
</td>
<td>
<p>StartHourUtc: Start hour after which cache patching can start.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.Sku">Sku
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.Redis_Spec">Redis_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/definitions/Sku">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/definitions/Sku</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>capacity</code><br/>
<em>
int
</em>
</td>
<td>
<p>Capacity: The size of the Redis cache to deploy. Valid values: for C (Basic/Standard) family (0, 1, 2, 3, 4, 5, 6), for
P (Premium) family (1, 2, 3, 4).</p>
</td>
</tr>
<tr>
<td>
<code>family</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.SkuFamily">
SkuFamily
</a>
</em>
</td>
<td>
<p>Family: The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.SkuName">
SkuName
</a>
</em>
</td>
<td>
<p>Name: The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.SkuARM">SkuARM
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisCreatePropertiesARM">RedisCreatePropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/definitions/Sku">https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/definitions/Sku</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>capacity</code><br/>
<em>
int
</em>
</td>
<td>
<p>Capacity: The size of the Redis cache to deploy. Valid values: for C (Basic/Standard) family (0, 1, 2, 3, 4, 5, 6), for
P (Premium) family (1, 2, 3, 4).</p>
</td>
</tr>
<tr>
<td>
<code>family</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.SkuFamily">
SkuFamily
</a>
</em>
</td>
<td>
<p>Family: The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.SkuName">
SkuName
</a>
</em>
</td>
<td>
<p>Name: The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.SkuFamily">SkuFamily
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.Sku">Sku</a>, <a href="#cache.azure.com/v1beta20201201.SkuARM">SkuARM</a>)
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
<tbody><tr><td><p>&#34;C&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.SkuName">SkuName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.Sku">Sku</a>, <a href="#cache.azure.com/v1beta20201201.SkuARM">SkuARM</a>)
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
</tr><tr><td><p>&#34;Premium&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.SkuStatusFamily">SkuStatusFamily
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.Sku_Status">Sku_Status</a>, <a href="#cache.azure.com/v1beta20201201.Sku_StatusARM">Sku_StatusARM</a>)
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
<tbody><tr><td><p>&#34;C&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.SkuStatusName">SkuStatusName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.Sku_Status">Sku_Status</a>, <a href="#cache.azure.com/v1beta20201201.Sku_StatusARM">Sku_StatusARM</a>)
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
</tr><tr><td><p>&#34;Premium&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Standard&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.Sku_Status">Sku_Status
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisResource_Status">RedisResource_Status</a>)
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
<code>capacity</code><br/>
<em>
int
</em>
</td>
<td>
<p>Capacity: The size of the Redis cache to deploy. Valid values: for C (Basic/Standard) family (0, 1, 2, 3, 4, 5, 6), for
P (Premium) family (1, 2, 3, 4).</p>
</td>
</tr>
<tr>
<td>
<code>family</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.SkuStatusFamily">
SkuStatusFamily
</a>
</em>
</td>
<td>
<p>Family: The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.SkuStatusName">
SkuStatusName
</a>
</em>
</td>
<td>
<p>Name: The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="cache.azure.com/v1beta20201201.Sku_StatusARM">Sku_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#cache.azure.com/v1beta20201201.RedisProperties_StatusARM">RedisProperties_StatusARM</a>)
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
<code>capacity</code><br/>
<em>
int
</em>
</td>
<td>
<p>Capacity: The size of the Redis cache to deploy. Valid values: for C (Basic/Standard) family (0, 1, 2, 3, 4, 5, 6), for
P (Premium) family (1, 2, 3, 4).</p>
</td>
</tr>
<tr>
<td>
<code>family</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.SkuStatusFamily">
SkuStatusFamily
</a>
</em>
</td>
<td>
<p>Family: The SKU family to use. Valid values: (C, P). (C = Basic/Standard, P = Premium).</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#cache.azure.com/v1beta20201201.SkuStatusName">
SkuStatusName
</a>
</em>
</td>
<td>
<p>Name: The type of Redis cache to deploy. Valid values: (Basic, Standard, Premium)</p>
</td>
</tr>
</tbody>
</table>
<hr/>
