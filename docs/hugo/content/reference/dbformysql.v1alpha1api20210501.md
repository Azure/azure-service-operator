---
---
<h2 id="dbformysql.azure.com/v1alpha1api20210501">dbformysql.azure.com/v1alpha1api20210501</h2>
<div>
<p>Package v1alpha1api20210501 contains API Schema definitions for the dbformysql v1alpha1api20210501 API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.Backup">Backup
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.FlexibleServers_Spec">FlexibleServers_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/Backup">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/Backup</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>backupRetentionDays</code><br/>
<em>
int
</em>
</td>
<td>
<p>BackupRetentionDays: Backup retention days for the server.</p>
</td>
</tr>
<tr>
<td>
<code>geoRedundantBackup</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.BackupGeoRedundantBackup">
BackupGeoRedundantBackup
</a>
</em>
</td>
<td>
<p>GeoRedundantBackup: Whether or not geo redundant backup is enabled.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.BackupARM">BackupARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesARM">ServerPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/Backup">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/Backup</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>backupRetentionDays</code><br/>
<em>
int
</em>
</td>
<td>
<p>BackupRetentionDays: Backup retention days for the server.</p>
</td>
</tr>
<tr>
<td>
<code>geoRedundantBackup</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.BackupGeoRedundantBackup">
BackupGeoRedundantBackup
</a>
</em>
</td>
<td>
<p>GeoRedundantBackup: Whether or not geo redundant backup is enabled.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.BackupGeoRedundantBackup">BackupGeoRedundantBackup
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.Backup">Backup</a>, <a href="#dbformysql.azure.com/v1alpha1api20210501.BackupARM">BackupARM</a>)
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
<h3 id="dbformysql.azure.com/v1alpha1api20210501.Backup_Status">Backup_Status
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.Server_Status">Server_Status</a>)
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
<code>backupRetentionDays</code><br/>
<em>
int
</em>
</td>
<td>
<p>BackupRetentionDays: Backup retention days for the server.</p>
</td>
</tr>
<tr>
<td>
<code>earliestRestoreDate</code><br/>
<em>
string
</em>
</td>
<td>
<p>EarliestRestoreDate: Earliest restore point creation time (ISO8601 format)</p>
</td>
</tr>
<tr>
<td>
<code>geoRedundantBackup</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.EnableStatusEnum_Status">
EnableStatusEnum_Status
</a>
</em>
</td>
<td>
<p>GeoRedundantBackup: Whether or not geo redundant backup is enabled.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.Backup_StatusARM">Backup_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.ServerProperties_StatusARM">ServerProperties_StatusARM</a>)
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
<code>backupRetentionDays</code><br/>
<em>
int
</em>
</td>
<td>
<p>BackupRetentionDays: Backup retention days for the server.</p>
</td>
</tr>
<tr>
<td>
<code>earliestRestoreDate</code><br/>
<em>
string
</em>
</td>
<td>
<p>EarliestRestoreDate: Earliest restore point creation time (ISO8601 format)</p>
</td>
</tr>
<tr>
<td>
<code>geoRedundantBackup</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.EnableStatusEnum_Status">
EnableStatusEnum_Status
</a>
</em>
</td>
<td>
<p>GeoRedundantBackup: Whether or not geo redundant backup is enabled.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.DatabasePropertiesARM">DatabasePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.FlexibleServersDatabases_SpecARM">FlexibleServersDatabases_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/DatabaseProperties">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/DatabaseProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>charset</code><br/>
<em>
string
</em>
</td>
<td>
<p>Charset: The charset of the database.</p>
</td>
</tr>
<tr>
<td>
<code>collation</code><br/>
<em>
string
</em>
</td>
<td>
<p>Collation: The collation of the database.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.DatabaseProperties_StatusARM">DatabaseProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.Database_StatusARM">Database_StatusARM</a>)
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
<code>charset</code><br/>
<em>
string
</em>
</td>
<td>
<p>Charset: The charset of the database.</p>
</td>
</tr>
<tr>
<td>
<code>collation</code><br/>
<em>
string
</em>
</td>
<td>
<p>Collation: The collation of the database.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.Database_Status">Database_Status
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.FlexibleServersDatabase">FlexibleServersDatabase</a>)
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
<code>charset</code><br/>
<em>
string
</em>
</td>
<td>
<p>Charset: The charset of the database.</p>
</td>
</tr>
<tr>
<td>
<code>collation</code><br/>
<em>
string
</em>
</td>
<td>
<p>Collation: The collation of the database.</p>
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
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}</p>
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
<code>systemData</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.SystemData_Status">
SystemData_Status
</a>
</em>
</td>
<td>
<p>SystemData: The system metadata relating to this resource.</p>
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
<h3 id="dbformysql.azure.com/v1alpha1api20210501.Database_StatusARM">Database_StatusARM
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
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}</p>
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
<a href="#dbformysql.azure.com/v1alpha1api20210501.DatabaseProperties_StatusARM">
DatabaseProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: The properties of a database.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.SystemData_StatusARM">
SystemData_StatusARM
</a>
</em>
</td>
<td>
<p>SystemData: The system metadata relating to this resource.</p>
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
<h3 id="dbformysql.azure.com/v1alpha1api20210501.EnableStatusEnum_Status">EnableStatusEnum_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.Backup_Status">Backup_Status</a>, <a href="#dbformysql.azure.com/v1alpha1api20210501.Backup_StatusARM">Backup_StatusARM</a>, <a href="#dbformysql.azure.com/v1alpha1api20210501.Network_Status">Network_Status</a>, <a href="#dbformysql.azure.com/v1alpha1api20210501.Network_StatusARM">Network_StatusARM</a>, <a href="#dbformysql.azure.com/v1alpha1api20210501.Storage_Status">Storage_Status</a>, <a href="#dbformysql.azure.com/v1alpha1api20210501.Storage_StatusARM">Storage_StatusARM</a>)
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
<h3 id="dbformysql.azure.com/v1alpha1api20210501.FirewallRulePropertiesARM">FirewallRulePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.FlexibleServersFirewallRules_SpecARM">FlexibleServersFirewallRules_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/FirewallRuleProperties">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/FirewallRuleProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>endIpAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>EndIpAddress: The end IP address of the server firewall rule. Must be IPv4 format.</p>
</td>
</tr>
<tr>
<td>
<code>startIpAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>StartIpAddress: The start IP address of the server firewall rule. Must be IPv4 format.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.FirewallRuleProperties_StatusARM">FirewallRuleProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.FirewallRule_StatusARM">FirewallRule_StatusARM</a>)
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
<code>endIpAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>EndIpAddress: The end IP address of the server firewall rule. Must be IPv4 format.</p>
</td>
</tr>
<tr>
<td>
<code>startIpAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>StartIpAddress: The start IP address of the server firewall rule. Must be IPv4 format.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.FirewallRule_Status">FirewallRule_Status
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.FlexibleServersFirewallRule">FlexibleServersFirewallRule</a>)
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
<code>endIpAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>EndIpAddress: The end IP address of the server firewall rule. Must be IPv4 format.</p>
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
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}</p>
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
<code>startIpAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>StartIpAddress: The start IP address of the server firewall rule. Must be IPv4 format.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.SystemData_Status">
SystemData_Status
</a>
</em>
</td>
<td>
<p>SystemData: The system metadata relating to this resource.</p>
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
<h3 id="dbformysql.azure.com/v1alpha1api20210501.FirewallRule_StatusARM">FirewallRule_StatusARM
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
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}</p>
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
<a href="#dbformysql.azure.com/v1alpha1api20210501.FirewallRuleProperties_StatusARM">
FirewallRuleProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: The properties of a firewall rule.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.SystemData_StatusARM">
SystemData_StatusARM
</a>
</em>
</td>
<td>
<p>SystemData: The system metadata relating to this resource.</p>
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
<h3 id="dbformysql.azure.com/v1alpha1api20210501.FlexibleServer">FlexibleServer
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/resourceDefinitions/flexibleServers">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/resourceDefinitions/flexibleServers</a></p>
</div>
<table>
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
<a href="#dbformysql.azure.com/v1alpha1api20210501.FlexibleServers_Spec">
FlexibleServers_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>administratorLogin</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdministratorLogin: The administrator&rsquo;s login name of a server. Can only be specified when the server is being created
(and is required for creation).</p>
</td>
</tr>
<tr>
<td>
<code>administratorLoginPassword</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretReference">
genruntime.SecretReference
</a>
</em>
</td>
<td>
<p>AdministratorLoginPassword: The password of the administrator login (required for server creation).</p>
</td>
</tr>
<tr>
<td>
<code>availabilityZone</code><br/>
<em>
string
</em>
</td>
<td>
<p>AvailabilityZone: availability Zone information of the server.</p>
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
<code>backup</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.Backup">
Backup
</a>
</em>
</td>
<td>
<p>Backup: Storage Profile properties of a server</p>
</td>
</tr>
<tr>
<td>
<code>createMode</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesCreateMode">
ServerPropertiesCreateMode
</a>
</em>
</td>
<td>
<p>CreateMode: The mode to create a new MySQL server.</p>
</td>
</tr>
<tr>
<td>
<code>highAvailability</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.HighAvailability">
HighAvailability
</a>
</em>
</td>
<td>
<p>HighAvailability: Network related properties of a server</p>
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
<code>maintenanceWindow</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.MaintenanceWindow">
MaintenanceWindow
</a>
</em>
</td>
<td>
<p>MaintenanceWindow: Maintenance window of a server.</p>
</td>
</tr>
<tr>
<td>
<code>network</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.Network">
Network
</a>
</em>
</td>
<td>
<p>Network: Network related properties of a server</p>
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
</td>
</tr>
<tr>
<td>
<code>replicationRole</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesReplicationRole">
ServerPropertiesReplicationRole
</a>
</em>
</td>
<td>
<p>ReplicationRole: The replication role.</p>
</td>
</tr>
<tr>
<td>
<code>restorePointInTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>RestorePointInTime: Restore point creation time (ISO8601 format), specifying the time to restore from.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.Sku">
Sku
</a>
</em>
</td>
<td>
<p>Sku: Billing information related properties of a server.</p>
</td>
</tr>
<tr>
<td>
<code>sourceServerResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourceServerResourceId: The source MySQL server id.</p>
</td>
</tr>
<tr>
<td>
<code>storage</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.Storage">
Storage
</a>
</em>
</td>
<td>
<p>Storage: Storage Profile properties of a server</p>
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
<code>version</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesVersion">
ServerPropertiesVersion
</a>
</em>
</td>
<td>
<p>Version: Server version.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.Server_Status">
Server_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.FlexibleServersDatabase">FlexibleServersDatabase
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/resourceDefinitions/flexibleServers_databases">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/resourceDefinitions/flexibleServers_databases</a></p>
</div>
<table>
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
<a href="#dbformysql.azure.com/v1alpha1api20210501.FlexibleServersDatabases_Spec">
FlexibleServersDatabases_Spec
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
<code>charset</code><br/>
<em>
string
</em>
</td>
<td>
<p>Charset: The charset of the database.</p>
</td>
</tr>
<tr>
<td>
<code>collation</code><br/>
<em>
string
</em>
</td>
<td>
<p>Collation: The collation of the database.</p>
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
<a href="#dbformysql.azure.com/v1alpha1api20210501.Database_Status">
Database_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.FlexibleServersDatabasesSpecAPIVersion">FlexibleServersDatabasesSpecAPIVersion
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
<h3 id="dbformysql.azure.com/v1alpha1api20210501.FlexibleServersDatabases_Spec">FlexibleServersDatabases_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.FlexibleServersDatabase">FlexibleServersDatabase</a>)
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
<code>charset</code><br/>
<em>
string
</em>
</td>
<td>
<p>Charset: The charset of the database.</p>
</td>
</tr>
<tr>
<td>
<code>collation</code><br/>
<em>
string
</em>
</td>
<td>
<p>Collation: The collation of the database.</p>
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
<h3 id="dbformysql.azure.com/v1alpha1api20210501.FlexibleServersDatabases_SpecARM">FlexibleServersDatabases_SpecARM
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
<p>Name: The name of the database.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.DatabasePropertiesARM">
DatabasePropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: The properties of a database.</p>
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
<h3 id="dbformysql.azure.com/v1alpha1api20210501.FlexibleServersFirewallRule">FlexibleServersFirewallRule
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/resourceDefinitions/flexibleServers_firewallRules">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/resourceDefinitions/flexibleServers_firewallRules</a></p>
</div>
<table>
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
<a href="#dbformysql.azure.com/v1alpha1api20210501.FlexibleServersFirewallRules_Spec">
FlexibleServersFirewallRules_Spec
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
<code>endIpAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>EndIpAddress: The end IP address of the server firewall rule. Must be IPv4 format.</p>
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
</td>
</tr>
<tr>
<td>
<code>startIpAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>StartIpAddress: The start IP address of the server firewall rule. Must be IPv4 format.</p>
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
<a href="#dbformysql.azure.com/v1alpha1api20210501.FirewallRule_Status">
FirewallRule_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.FlexibleServersFirewallRulesSpecAPIVersion">FlexibleServersFirewallRulesSpecAPIVersion
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
<h3 id="dbformysql.azure.com/v1alpha1api20210501.FlexibleServersFirewallRules_Spec">FlexibleServersFirewallRules_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.FlexibleServersFirewallRule">FlexibleServersFirewallRule</a>)
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
<code>endIpAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>EndIpAddress: The end IP address of the server firewall rule. Must be IPv4 format.</p>
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
</td>
</tr>
<tr>
<td>
<code>startIpAddress</code><br/>
<em>
string
</em>
</td>
<td>
<p>StartIpAddress: The start IP address of the server firewall rule. Must be IPv4 format.</p>
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
<h3 id="dbformysql.azure.com/v1alpha1api20210501.FlexibleServersFirewallRules_SpecARM">FlexibleServersFirewallRules_SpecARM
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
<p>Name: The name of the server firewall rule.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.FirewallRulePropertiesARM">
FirewallRulePropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: The properties of a server firewall rule.</p>
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
<h3 id="dbformysql.azure.com/v1alpha1api20210501.FlexibleServersSpecAPIVersion">FlexibleServersSpecAPIVersion
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
<h3 id="dbformysql.azure.com/v1alpha1api20210501.FlexibleServers_Spec">FlexibleServers_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.FlexibleServer">FlexibleServer</a>)
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
<code>administratorLogin</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdministratorLogin: The administrator&rsquo;s login name of a server. Can only be specified when the server is being created
(and is required for creation).</p>
</td>
</tr>
<tr>
<td>
<code>administratorLoginPassword</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretReference">
genruntime.SecretReference
</a>
</em>
</td>
<td>
<p>AdministratorLoginPassword: The password of the administrator login (required for server creation).</p>
</td>
</tr>
<tr>
<td>
<code>availabilityZone</code><br/>
<em>
string
</em>
</td>
<td>
<p>AvailabilityZone: availability Zone information of the server.</p>
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
<code>backup</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.Backup">
Backup
</a>
</em>
</td>
<td>
<p>Backup: Storage Profile properties of a server</p>
</td>
</tr>
<tr>
<td>
<code>createMode</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesCreateMode">
ServerPropertiesCreateMode
</a>
</em>
</td>
<td>
<p>CreateMode: The mode to create a new MySQL server.</p>
</td>
</tr>
<tr>
<td>
<code>highAvailability</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.HighAvailability">
HighAvailability
</a>
</em>
</td>
<td>
<p>HighAvailability: Network related properties of a server</p>
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
<code>maintenanceWindow</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.MaintenanceWindow">
MaintenanceWindow
</a>
</em>
</td>
<td>
<p>MaintenanceWindow: Maintenance window of a server.</p>
</td>
</tr>
<tr>
<td>
<code>network</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.Network">
Network
</a>
</em>
</td>
<td>
<p>Network: Network related properties of a server</p>
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
</td>
</tr>
<tr>
<td>
<code>replicationRole</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesReplicationRole">
ServerPropertiesReplicationRole
</a>
</em>
</td>
<td>
<p>ReplicationRole: The replication role.</p>
</td>
</tr>
<tr>
<td>
<code>restorePointInTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>RestorePointInTime: Restore point creation time (ISO8601 format), specifying the time to restore from.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.Sku">
Sku
</a>
</em>
</td>
<td>
<p>Sku: Billing information related properties of a server.</p>
</td>
</tr>
<tr>
<td>
<code>sourceServerResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourceServerResourceId: The source MySQL server id.</p>
</td>
</tr>
<tr>
<td>
<code>storage</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.Storage">
Storage
</a>
</em>
</td>
<td>
<p>Storage: Storage Profile properties of a server</p>
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
<code>version</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesVersion">
ServerPropertiesVersion
</a>
</em>
</td>
<td>
<p>Version: Server version.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.FlexibleServers_SpecARM">FlexibleServers_SpecARM
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
<p>Name: The name of the server.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesARM">
ServerPropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: The properties of a server.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.SkuARM">
SkuARM
</a>
</em>
</td>
<td>
<p>Sku: Billing information related properties of a server.</p>
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
<h3 id="dbformysql.azure.com/v1alpha1api20210501.HighAvailability">HighAvailability
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.FlexibleServers_Spec">FlexibleServers_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/HighAvailability">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/HighAvailability</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.HighAvailabilityMode">
HighAvailabilityMode
</a>
</em>
</td>
<td>
<p>Mode: High availability mode for a server.</p>
</td>
</tr>
<tr>
<td>
<code>standbyAvailabilityZone</code><br/>
<em>
string
</em>
</td>
<td>
<p>StandbyAvailabilityZone: Availability zone of the standby server.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.HighAvailabilityARM">HighAvailabilityARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesARM">ServerPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/HighAvailability">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/HighAvailability</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>mode</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.HighAvailabilityMode">
HighAvailabilityMode
</a>
</em>
</td>
<td>
<p>Mode: High availability mode for a server.</p>
</td>
</tr>
<tr>
<td>
<code>standbyAvailabilityZone</code><br/>
<em>
string
</em>
</td>
<td>
<p>StandbyAvailabilityZone: Availability zone of the standby server.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.HighAvailabilityMode">HighAvailabilityMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.HighAvailability">HighAvailability</a>, <a href="#dbformysql.azure.com/v1alpha1api20210501.HighAvailabilityARM">HighAvailabilityARM</a>)
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
</tr><tr><td><p>&#34;SameZone&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ZoneRedundant&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.HighAvailabilityStatusMode">HighAvailabilityStatusMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.HighAvailability_Status">HighAvailability_Status</a>, <a href="#dbformysql.azure.com/v1alpha1api20210501.HighAvailability_StatusARM">HighAvailability_StatusARM</a>)
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
</tr><tr><td><p>&#34;SameZone&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ZoneRedundant&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.HighAvailabilityStatusState">HighAvailabilityStatusState
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.HighAvailability_Status">HighAvailability_Status</a>, <a href="#dbformysql.azure.com/v1alpha1api20210501.HighAvailability_StatusARM">HighAvailability_StatusARM</a>)
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
<tbody><tr><td><p>&#34;CreatingStandby&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;FailingOver&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Healthy&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;NotEnabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;RemovingStandby&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.HighAvailability_Status">HighAvailability_Status
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.Server_Status">Server_Status</a>)
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
<code>mode</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.HighAvailabilityStatusMode">
HighAvailabilityStatusMode
</a>
</em>
</td>
<td>
<p>Mode: High availability mode for a server.</p>
</td>
</tr>
<tr>
<td>
<code>standbyAvailabilityZone</code><br/>
<em>
string
</em>
</td>
<td>
<p>StandbyAvailabilityZone: Availability zone of the standby server.</p>
</td>
</tr>
<tr>
<td>
<code>state</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.HighAvailabilityStatusState">
HighAvailabilityStatusState
</a>
</em>
</td>
<td>
<p>State: The state of server high availability.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.HighAvailability_StatusARM">HighAvailability_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.ServerProperties_StatusARM">ServerProperties_StatusARM</a>)
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
<code>mode</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.HighAvailabilityStatusMode">
HighAvailabilityStatusMode
</a>
</em>
</td>
<td>
<p>Mode: High availability mode for a server.</p>
</td>
</tr>
<tr>
<td>
<code>standbyAvailabilityZone</code><br/>
<em>
string
</em>
</td>
<td>
<p>StandbyAvailabilityZone: Availability zone of the standby server.</p>
</td>
</tr>
<tr>
<td>
<code>state</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.HighAvailabilityStatusState">
HighAvailabilityStatusState
</a>
</em>
</td>
<td>
<p>State: The state of server high availability.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.MaintenanceWindow">MaintenanceWindow
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.FlexibleServers_Spec">FlexibleServers_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/MaintenanceWindow">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/MaintenanceWindow</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>customWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>CustomWindow: indicates whether custom window is enabled or disabled</p>
</td>
</tr>
<tr>
<td>
<code>dayOfWeek</code><br/>
<em>
int
</em>
</td>
<td>
<p>DayOfWeek: day of week for maintenance window</p>
</td>
</tr>
<tr>
<td>
<code>startHour</code><br/>
<em>
int
</em>
</td>
<td>
<p>StartHour: start hour for maintenance window</p>
</td>
</tr>
<tr>
<td>
<code>startMinute</code><br/>
<em>
int
</em>
</td>
<td>
<p>StartMinute: start minute for maintenance window</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.MaintenanceWindowARM">MaintenanceWindowARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesARM">ServerPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/MaintenanceWindow">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/MaintenanceWindow</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>customWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>CustomWindow: indicates whether custom window is enabled or disabled</p>
</td>
</tr>
<tr>
<td>
<code>dayOfWeek</code><br/>
<em>
int
</em>
</td>
<td>
<p>DayOfWeek: day of week for maintenance window</p>
</td>
</tr>
<tr>
<td>
<code>startHour</code><br/>
<em>
int
</em>
</td>
<td>
<p>StartHour: start hour for maintenance window</p>
</td>
</tr>
<tr>
<td>
<code>startMinute</code><br/>
<em>
int
</em>
</td>
<td>
<p>StartMinute: start minute for maintenance window</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.MaintenanceWindow_Status">MaintenanceWindow_Status
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.Server_Status">Server_Status</a>)
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
<code>customWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>CustomWindow: indicates whether custom window is enabled or disabled</p>
</td>
</tr>
<tr>
<td>
<code>dayOfWeek</code><br/>
<em>
int
</em>
</td>
<td>
<p>DayOfWeek: day of week for maintenance window</p>
</td>
</tr>
<tr>
<td>
<code>startHour</code><br/>
<em>
int
</em>
</td>
<td>
<p>StartHour: start hour for maintenance window</p>
</td>
</tr>
<tr>
<td>
<code>startMinute</code><br/>
<em>
int
</em>
</td>
<td>
<p>StartMinute: start minute for maintenance window</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.MaintenanceWindow_StatusARM">MaintenanceWindow_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.ServerProperties_StatusARM">ServerProperties_StatusARM</a>)
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
<code>customWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>CustomWindow: indicates whether custom window is enabled or disabled</p>
</td>
</tr>
<tr>
<td>
<code>dayOfWeek</code><br/>
<em>
int
</em>
</td>
<td>
<p>DayOfWeek: day of week for maintenance window</p>
</td>
</tr>
<tr>
<td>
<code>startHour</code><br/>
<em>
int
</em>
</td>
<td>
<p>StartHour: start hour for maintenance window</p>
</td>
</tr>
<tr>
<td>
<code>startMinute</code><br/>
<em>
int
</em>
</td>
<td>
<p>StartMinute: start minute for maintenance window</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.Network">Network
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.FlexibleServers_Spec">FlexibleServers_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/Network">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/Network</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>delegatedSubnetResourceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>DelegatedSubnetResourceReference: Delegated subnet resource id used to setup vnet for a server.</p>
</td>
</tr>
<tr>
<td>
<code>privateDnsZoneResourceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>PrivateDnsZoneResourceReference: Private DNS zone resource id.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.NetworkARM">NetworkARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesARM">ServerPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/Network">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/Network</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>delegatedSubnetResourceId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>privateDnsZoneResourceId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.Network_Status">Network_Status
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.Server_Status">Server_Status</a>)
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
<code>delegatedSubnetResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>DelegatedSubnetResourceId: Delegated subnet resource id used to setup vnet for a server.</p>
</td>
</tr>
<tr>
<td>
<code>privateDnsZoneResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateDnsZoneResourceId: Private DNS zone resource id.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.EnableStatusEnum_Status">
EnableStatusEnum_Status
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Whether or not public network access is allowed for this server. Value is &lsquo;Disabled&rsquo; when server
has VNet integration.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.Network_StatusARM">Network_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.ServerProperties_StatusARM">ServerProperties_StatusARM</a>)
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
<code>delegatedSubnetResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>DelegatedSubnetResourceId: Delegated subnet resource id used to setup vnet for a server.</p>
</td>
</tr>
<tr>
<td>
<code>privateDnsZoneResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateDnsZoneResourceId: Private DNS zone resource id.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.EnableStatusEnum_Status">
EnableStatusEnum_Status
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: Whether or not public network access is allowed for this server. Value is &lsquo;Disabled&rsquo; when server
has VNet integration.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.ReplicationRole_Status">ReplicationRole_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.ServerProperties_StatusARM">ServerProperties_StatusARM</a>, <a href="#dbformysql.azure.com/v1alpha1api20210501.Server_Status">Server_Status</a>)
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
</tr><tr><td><p>&#34;Replica&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Source&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesARM">ServerPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.FlexibleServers_SpecARM">FlexibleServers_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/ServerProperties">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/ServerProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>administratorLogin</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdministratorLogin: The administrator&rsquo;s login name of a server. Can only be specified when the server is being created
(and is required for creation).</p>
</td>
</tr>
<tr>
<td>
<code>administratorLoginPassword</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdministratorLoginPassword: The password of the administrator login (required for server creation).</p>
</td>
</tr>
<tr>
<td>
<code>availabilityZone</code><br/>
<em>
string
</em>
</td>
<td>
<p>AvailabilityZone: availability Zone information of the server.</p>
</td>
</tr>
<tr>
<td>
<code>backup</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.BackupARM">
BackupARM
</a>
</em>
</td>
<td>
<p>Backup: Storage Profile properties of a server</p>
</td>
</tr>
<tr>
<td>
<code>createMode</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesCreateMode">
ServerPropertiesCreateMode
</a>
</em>
</td>
<td>
<p>CreateMode: The mode to create a new MySQL server.</p>
</td>
</tr>
<tr>
<td>
<code>highAvailability</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.HighAvailabilityARM">
HighAvailabilityARM
</a>
</em>
</td>
<td>
<p>HighAvailability: Network related properties of a server</p>
</td>
</tr>
<tr>
<td>
<code>maintenanceWindow</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.MaintenanceWindowARM">
MaintenanceWindowARM
</a>
</em>
</td>
<td>
<p>MaintenanceWindow: Maintenance window of a server.</p>
</td>
</tr>
<tr>
<td>
<code>network</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.NetworkARM">
NetworkARM
</a>
</em>
</td>
<td>
<p>Network: Network related properties of a server</p>
</td>
</tr>
<tr>
<td>
<code>replicationRole</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesReplicationRole">
ServerPropertiesReplicationRole
</a>
</em>
</td>
<td>
<p>ReplicationRole: The replication role.</p>
</td>
</tr>
<tr>
<td>
<code>restorePointInTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>RestorePointInTime: Restore point creation time (ISO8601 format), specifying the time to restore from.</p>
</td>
</tr>
<tr>
<td>
<code>sourceServerResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourceServerResourceId: The source MySQL server id.</p>
</td>
</tr>
<tr>
<td>
<code>storage</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.StorageARM">
StorageARM
</a>
</em>
</td>
<td>
<p>Storage: Storage Profile properties of a server</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesVersion">
ServerPropertiesVersion
</a>
</em>
</td>
<td>
<p>Version: Server version.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesCreateMode">ServerPropertiesCreateMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.FlexibleServers_Spec">FlexibleServers_Spec</a>, <a href="#dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesARM">ServerPropertiesARM</a>)
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
<tbody><tr><td><p>&#34;Default&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GeoRestore&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;PointInTimeRestore&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Replica&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesReplicationRole">ServerPropertiesReplicationRole
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.FlexibleServers_Spec">FlexibleServers_Spec</a>, <a href="#dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesARM">ServerPropertiesARM</a>)
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
</tr><tr><td><p>&#34;Replica&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Source&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesStatusCreateMode">ServerPropertiesStatusCreateMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.ServerProperties_StatusARM">ServerProperties_StatusARM</a>, <a href="#dbformysql.azure.com/v1alpha1api20210501.Server_Status">Server_Status</a>)
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
<tbody><tr><td><p>&#34;Default&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GeoRestore&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;PointInTimeRestore&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Replica&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesStatusState">ServerPropertiesStatusState
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.ServerProperties_StatusARM">ServerProperties_StatusARM</a>, <a href="#dbformysql.azure.com/v1alpha1api20210501.Server_Status">Server_Status</a>)
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
</tr><tr><td><p>&#34;Dropping&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Ready&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Starting&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Stopped&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Stopping&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Updating&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesVersion">ServerPropertiesVersion
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.FlexibleServers_Spec">FlexibleServers_Spec</a>, <a href="#dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesARM">ServerPropertiesARM</a>)
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
<tbody><tr><td><p>&#34;5.7&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;8.0.21&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.ServerProperties_StatusARM">ServerProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.Server_StatusARM">Server_StatusARM</a>)
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
<code>administratorLogin</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdministratorLogin: The administrator&rsquo;s login name of a server. Can only be specified when the server is being created
(and is required for creation).</p>
</td>
</tr>
<tr>
<td>
<code>availabilityZone</code><br/>
<em>
string
</em>
</td>
<td>
<p>AvailabilityZone: availability Zone information of the server.</p>
</td>
</tr>
<tr>
<td>
<code>backup</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.Backup_StatusARM">
Backup_StatusARM
</a>
</em>
</td>
<td>
<p>Backup: Backup related properties of a server.</p>
</td>
</tr>
<tr>
<td>
<code>createMode</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesStatusCreateMode">
ServerPropertiesStatusCreateMode
</a>
</em>
</td>
<td>
<p>CreateMode: The mode to create a new MySQL server.</p>
</td>
</tr>
<tr>
<td>
<code>fullyQualifiedDomainName</code><br/>
<em>
string
</em>
</td>
<td>
<p>FullyQualifiedDomainName: The fully qualified domain name of a server.</p>
</td>
</tr>
<tr>
<td>
<code>highAvailability</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.HighAvailability_StatusARM">
HighAvailability_StatusARM
</a>
</em>
</td>
<td>
<p>HighAvailability: High availability related properties of a server.</p>
</td>
</tr>
<tr>
<td>
<code>maintenanceWindow</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.MaintenanceWindow_StatusARM">
MaintenanceWindow_StatusARM
</a>
</em>
</td>
<td>
<p>MaintenanceWindow: Maintenance window of a server.</p>
</td>
</tr>
<tr>
<td>
<code>network</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.Network_StatusARM">
Network_StatusARM
</a>
</em>
</td>
<td>
<p>Network: Network related properties of a server.</p>
</td>
</tr>
<tr>
<td>
<code>replicaCapacity</code><br/>
<em>
int
</em>
</td>
<td>
<p>ReplicaCapacity: The maximum number of replicas that a primary server can have.</p>
</td>
</tr>
<tr>
<td>
<code>replicationRole</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.ReplicationRole_Status">
ReplicationRole_Status
</a>
</em>
</td>
<td>
<p>ReplicationRole: The replication role.</p>
</td>
</tr>
<tr>
<td>
<code>restorePointInTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>RestorePointInTime: Restore point creation time (ISO8601 format), specifying the time to restore from.</p>
</td>
</tr>
<tr>
<td>
<code>sourceServerResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourceServerResourceId: The source MySQL server id.</p>
</td>
</tr>
<tr>
<td>
<code>state</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesStatusState">
ServerPropertiesStatusState
</a>
</em>
</td>
<td>
<p>State: The state of a server.</p>
</td>
</tr>
<tr>
<td>
<code>storage</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.Storage_StatusARM">
Storage_StatusARM
</a>
</em>
</td>
<td>
<p>Storage: Storage related properties of a server.</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.ServerVersion_Status">
ServerVersion_Status
</a>
</em>
</td>
<td>
<p>Version: Server version.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.ServerVersion_Status">ServerVersion_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.ServerProperties_StatusARM">ServerProperties_StatusARM</a>, <a href="#dbformysql.azure.com/v1alpha1api20210501.Server_Status">Server_Status</a>)
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
<tbody><tr><td><p>&#34;5.7&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;8.0.21&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.Server_Status">Server_Status
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.FlexibleServer">FlexibleServer</a>)
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
<code>administratorLogin</code><br/>
<em>
string
</em>
</td>
<td>
<p>AdministratorLogin: The administrator&rsquo;s login name of a server. Can only be specified when the server is being created
(and is required for creation).</p>
</td>
</tr>
<tr>
<td>
<code>availabilityZone</code><br/>
<em>
string
</em>
</td>
<td>
<p>AvailabilityZone: availability Zone information of the server.</p>
</td>
</tr>
<tr>
<td>
<code>backup</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.Backup_Status">
Backup_Status
</a>
</em>
</td>
<td>
<p>Backup: Backup related properties of a server.</p>
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
<code>createMode</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesStatusCreateMode">
ServerPropertiesStatusCreateMode
</a>
</em>
</td>
<td>
<p>CreateMode: The mode to create a new MySQL server.</p>
</td>
</tr>
<tr>
<td>
<code>fullyQualifiedDomainName</code><br/>
<em>
string
</em>
</td>
<td>
<p>FullyQualifiedDomainName: The fully qualified domain name of a server.</p>
</td>
</tr>
<tr>
<td>
<code>highAvailability</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.HighAvailability_Status">
HighAvailability_Status
</a>
</em>
</td>
<td>
<p>HighAvailability: High availability related properties of a server.</p>
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
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}</p>
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
<code>maintenanceWindow</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.MaintenanceWindow_Status">
MaintenanceWindow_Status
</a>
</em>
</td>
<td>
<p>MaintenanceWindow: Maintenance window of a server.</p>
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
<code>network</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.Network_Status">
Network_Status
</a>
</em>
</td>
<td>
<p>Network: Network related properties of a server.</p>
</td>
</tr>
<tr>
<td>
<code>replicaCapacity</code><br/>
<em>
int
</em>
</td>
<td>
<p>ReplicaCapacity: The maximum number of replicas that a primary server can have.</p>
</td>
</tr>
<tr>
<td>
<code>replicationRole</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.ReplicationRole_Status">
ReplicationRole_Status
</a>
</em>
</td>
<td>
<p>ReplicationRole: The replication role.</p>
</td>
</tr>
<tr>
<td>
<code>restorePointInTime</code><br/>
<em>
string
</em>
</td>
<td>
<p>RestorePointInTime: Restore point creation time (ISO8601 format), specifying the time to restore from.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.Sku_Status">
Sku_Status
</a>
</em>
</td>
<td>
<p>Sku: The SKU (pricing tier) of the server.</p>
</td>
</tr>
<tr>
<td>
<code>sourceServerResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>SourceServerResourceId: The source MySQL server id.</p>
</td>
</tr>
<tr>
<td>
<code>state</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesStatusState">
ServerPropertiesStatusState
</a>
</em>
</td>
<td>
<p>State: The state of a server.</p>
</td>
</tr>
<tr>
<td>
<code>storage</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.Storage_Status">
Storage_Status
</a>
</em>
</td>
<td>
<p>Storage: Storage related properties of a server.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.SystemData_Status">
SystemData_Status
</a>
</em>
</td>
<td>
<p>SystemData: The system metadata relating to this resource.</p>
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
<code>version</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.ServerVersion_Status">
ServerVersion_Status
</a>
</em>
</td>
<td>
<p>Version: Server version.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.Server_StatusARM">Server_StatusARM
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
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}</p>
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
<a href="#dbformysql.azure.com/v1alpha1api20210501.ServerProperties_StatusARM">
ServerProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the server.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.Sku_StatusARM">
Sku_StatusARM
</a>
</em>
</td>
<td>
<p>Sku: The SKU (pricing tier) of the server.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.SystemData_StatusARM">
SystemData_StatusARM
</a>
</em>
</td>
<td>
<p>SystemData: The system metadata relating to this resource.</p>
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
<h3 id="dbformysql.azure.com/v1alpha1api20210501.Sku">Sku
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.FlexibleServers_Spec">FlexibleServers_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/Sku">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/Sku</a></p>
</div>
<table>
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
<p>Name: The name of the sku, e.g. Standard_D32s_v3.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.SkuTier">
SkuTier
</a>
</em>
</td>
<td>
<p>Tier: The tier of the particular SKU, e.g. GeneralPurpose.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.SkuARM">SkuARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.FlexibleServers_SpecARM">FlexibleServers_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/Sku">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/Sku</a></p>
</div>
<table>
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
<p>Name: The name of the sku, e.g. Standard_D32s_v3.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.SkuTier">
SkuTier
</a>
</em>
</td>
<td>
<p>Tier: The tier of the particular SKU, e.g. GeneralPurpose.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.SkuStatusTier">SkuStatusTier
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.Sku_Status">Sku_Status</a>, <a href="#dbformysql.azure.com/v1alpha1api20210501.Sku_StatusARM">Sku_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Burstable&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GeneralPurpose&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;MemoryOptimized&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.SkuTier">SkuTier
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.Sku">Sku</a>, <a href="#dbformysql.azure.com/v1alpha1api20210501.SkuARM">SkuARM</a>)
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
<tbody><tr><td><p>&#34;Burstable&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GeneralPurpose&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;MemoryOptimized&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.Sku_Status">Sku_Status
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.Server_Status">Server_Status</a>)
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
<p>Name: The name of the sku, e.g. Standard_D32s_v3.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.SkuStatusTier">
SkuStatusTier
</a>
</em>
</td>
<td>
<p>Tier: The tier of the particular SKU, e.g. GeneralPurpose.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.Sku_StatusARM">Sku_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.Server_StatusARM">Server_StatusARM</a>)
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
<p>Name: The name of the sku, e.g. Standard_D32s_v3.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.SkuStatusTier">
SkuStatusTier
</a>
</em>
</td>
<td>
<p>Tier: The tier of the particular SKU, e.g. GeneralPurpose.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.Storage">Storage
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.FlexibleServers_Spec">FlexibleServers_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/Storage">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/Storage</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>autoGrow</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.StorageAutoGrow">
StorageAutoGrow
</a>
</em>
</td>
<td>
<p>AutoGrow: Enable Storage Auto Grow or not.</p>
</td>
</tr>
<tr>
<td>
<code>iops</code><br/>
<em>
int
</em>
</td>
<td>
<p>Iops: Storage IOPS for a server.</p>
</td>
</tr>
<tr>
<td>
<code>storageSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
<p>StorageSizeGB: Max storage size allowed for a server.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.StorageARM">StorageARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.ServerPropertiesARM">ServerPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/Storage">https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/Storage</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>autoGrow</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.StorageAutoGrow">
StorageAutoGrow
</a>
</em>
</td>
<td>
<p>AutoGrow: Enable Storage Auto Grow or not.</p>
</td>
</tr>
<tr>
<td>
<code>iops</code><br/>
<em>
int
</em>
</td>
<td>
<p>Iops: Storage IOPS for a server.</p>
</td>
</tr>
<tr>
<td>
<code>storageSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
<p>StorageSizeGB: Max storage size allowed for a server.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.StorageAutoGrow">StorageAutoGrow
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.Storage">Storage</a>, <a href="#dbformysql.azure.com/v1alpha1api20210501.StorageARM">StorageARM</a>)
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
<h3 id="dbformysql.azure.com/v1alpha1api20210501.Storage_Status">Storage_Status
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.Server_Status">Server_Status</a>)
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
<code>autoGrow</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.EnableStatusEnum_Status">
EnableStatusEnum_Status
</a>
</em>
</td>
<td>
<p>AutoGrow: Enable Storage Auto Grow or not.</p>
</td>
</tr>
<tr>
<td>
<code>iops</code><br/>
<em>
int
</em>
</td>
<td>
<p>Iops: Storage IOPS for a server.</p>
</td>
</tr>
<tr>
<td>
<code>storageSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
<p>StorageSizeGB: Max storage size allowed for a server.</p>
</td>
</tr>
<tr>
<td>
<code>storageSku</code><br/>
<em>
string
</em>
</td>
<td>
<p>StorageSku: The sku name of the server storage.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.Storage_StatusARM">Storage_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.ServerProperties_StatusARM">ServerProperties_StatusARM</a>)
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
<code>autoGrow</code><br/>
<em>
<a href="#dbformysql.azure.com/v1alpha1api20210501.EnableStatusEnum_Status">
EnableStatusEnum_Status
</a>
</em>
</td>
<td>
<p>AutoGrow: Enable Storage Auto Grow or not.</p>
</td>
</tr>
<tr>
<td>
<code>iops</code><br/>
<em>
int
</em>
</td>
<td>
<p>Iops: Storage IOPS for a server.</p>
</td>
</tr>
<tr>
<td>
<code>storageSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
<p>StorageSizeGB: Max storage size allowed for a server.</p>
</td>
</tr>
<tr>
<td>
<code>storageSku</code><br/>
<em>
string
</em>
</td>
<td>
<p>StorageSku: The sku name of the server storage.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1alpha1api20210501.SystemDataStatusCreatedByType">SystemDataStatusCreatedByType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.SystemData_Status">SystemData_Status</a>, <a href="#dbformysql.azure.com/v1alpha1api20210501.SystemData_StatusARM">SystemData_StatusARM</a>)
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
<h3 id="dbformysql.azure.com/v1alpha1api20210501.SystemDataStatusLastModifiedByType">SystemDataStatusLastModifiedByType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.SystemData_Status">SystemData_Status</a>, <a href="#dbformysql.azure.com/v1alpha1api20210501.SystemData_StatusARM">SystemData_StatusARM</a>)
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
<h3 id="dbformysql.azure.com/v1alpha1api20210501.SystemData_Status">SystemData_Status
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.Database_Status">Database_Status</a>, <a href="#dbformysql.azure.com/v1alpha1api20210501.FirewallRule_Status">FirewallRule_Status</a>, <a href="#dbformysql.azure.com/v1alpha1api20210501.Server_Status">Server_Status</a>)
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
<a href="#dbformysql.azure.com/v1alpha1api20210501.SystemDataStatusCreatedByType">
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
<a href="#dbformysql.azure.com/v1alpha1api20210501.SystemDataStatusLastModifiedByType">
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
<h3 id="dbformysql.azure.com/v1alpha1api20210501.SystemData_StatusARM">SystemData_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1alpha1api20210501.Database_StatusARM">Database_StatusARM</a>, <a href="#dbformysql.azure.com/v1alpha1api20210501.FirewallRule_StatusARM">FirewallRule_StatusARM</a>, <a href="#dbformysql.azure.com/v1alpha1api20210501.Server_StatusARM">Server_StatusARM</a>)
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
<a href="#dbformysql.azure.com/v1alpha1api20210501.SystemDataStatusCreatedByType">
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
<a href="#dbformysql.azure.com/v1alpha1api20210501.SystemDataStatusLastModifiedByType">
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
<hr/>
