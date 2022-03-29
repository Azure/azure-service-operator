<h2 id="dbforpostgresql.azure.com/v1alpha1api20210601">dbforpostgresql.azure.com/v1alpha1api20210601</h2>
<div>
<p>Package v1alpha1api20210601 contains API Schema definitions for the dbforpostgresql v1alpha1api20210601 API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.Backup">Backup
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServers_Spec">FlexibleServers_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/Backup">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/Backup</a></p>
</div>
<table>
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.BackupGeoRedundantBackup">
BackupGeoRedundantBackup
</a>
</em>
</td>
<td>
<p>GeoRedundantBackup: A value indicating whether Geo-Redundant backup is enabled on the server.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.BackupARM">BackupARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerPropertiesARM">ServerPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/Backup">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/Backup</a></p>
</div>
<table>
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.BackupGeoRedundantBackup">
BackupGeoRedundantBackup
</a>
</em>
</td>
<td>
<p>GeoRedundantBackup: A value indicating whether Geo-Redundant backup is enabled on the server.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.BackupGeoRedundantBackup">BackupGeoRedundantBackup
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Backup">Backup</a>, <a href="#dbforpostgresql.azure.com/v1alpha1api20210601.BackupARM">BackupARM</a>)
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.BackupStatusGeoRedundantBackup">BackupStatusGeoRedundantBackup
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Backup_Status">Backup_Status</a>, <a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Backup_StatusARM">Backup_StatusARM</a>)
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.Backup_Status">Backup_Status
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Server_Status">Server_Status</a>)
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
<p>EarliestRestoreDate: The earliest restore point time (ISO8601 format) for server.</p>
</td>
</tr>
<tr>
<td>
<code>geoRedundantBackup</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.BackupStatusGeoRedundantBackup">
BackupStatusGeoRedundantBackup
</a>
</em>
</td>
<td>
<p>GeoRedundantBackup: A value indicating whether Geo-Redundant backup is enabled on the server.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.Backup_StatusARM">Backup_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerProperties_StatusARM">ServerProperties_StatusARM</a>)
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
<p>EarliestRestoreDate: The earliest restore point time (ISO8601 format) for server.</p>
</td>
</tr>
<tr>
<td>
<code>geoRedundantBackup</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.BackupStatusGeoRedundantBackup">
BackupStatusGeoRedundantBackup
</a>
</em>
</td>
<td>
<p>GeoRedundantBackup: A value indicating whether Geo-Redundant backup is enabled on the server.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.ConfigurationPropertiesARM">ConfigurationPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServersConfigurations_SpecARM">FlexibleServersConfigurations_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/ConfigurationProperties">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/ConfigurationProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>source</code><br/>
<em>
string
</em>
</td>
<td>
<p>Source: Source of the configuration.</p>
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
<p>Value: Value of the configuration.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.ConfigurationPropertiesStatusDataType">ConfigurationPropertiesStatusDataType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ConfigurationProperties_StatusARM">ConfigurationProperties_StatusARM</a>, <a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Configuration_Status">Configuration_Status</a>)
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
<tbody><tr><td><p>&#34;Boolean&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enumeration&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Integer&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Numeric&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.ConfigurationProperties_StatusARM">ConfigurationProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Configuration_StatusARM">Configuration_StatusARM</a>)
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
<code>allowedValues</code><br/>
<em>
string
</em>
</td>
<td>
<p>AllowedValues: Allowed values of the configuration.</p>
</td>
</tr>
<tr>
<td>
<code>dataType</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ConfigurationPropertiesStatusDataType">
ConfigurationPropertiesStatusDataType
</a>
</em>
</td>
<td>
<p>DataType: Data type of the configuration.</p>
</td>
</tr>
<tr>
<td>
<code>defaultValue</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultValue: Default value of the configuration.</p>
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
<p>Description: Description of the configuration.</p>
</td>
</tr>
<tr>
<td>
<code>documentationLink</code><br/>
<em>
string
</em>
</td>
<td>
<p>DocumentationLink: Configuration documentation link.</p>
</td>
</tr>
<tr>
<td>
<code>isConfigPendingRestart</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsConfigPendingRestart: Configuration is pending restart or not.</p>
</td>
</tr>
<tr>
<td>
<code>isDynamicConfig</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsDynamicConfig: Configuration dynamic or static.</p>
</td>
</tr>
<tr>
<td>
<code>isReadOnly</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsReadOnly: Configuration read-only or not.</p>
</td>
</tr>
<tr>
<td>
<code>source</code><br/>
<em>
string
</em>
</td>
<td>
<p>Source: Source of the configuration.</p>
</td>
</tr>
<tr>
<td>
<code>unit</code><br/>
<em>
string
</em>
</td>
<td>
<p>Unit: Configuration unit.</p>
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
<p>Value: Value of the configuration.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.Configuration_Status">Configuration_Status
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServersConfiguration">FlexibleServersConfiguration</a>)
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
<code>allowedValues</code><br/>
<em>
string
</em>
</td>
<td>
<p>AllowedValues: Allowed values of the configuration.</p>
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
<code>dataType</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ConfigurationPropertiesStatusDataType">
ConfigurationPropertiesStatusDataType
</a>
</em>
</td>
<td>
<p>DataType: Data type of the configuration.</p>
</td>
</tr>
<tr>
<td>
<code>defaultValue</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultValue: Default value of the configuration.</p>
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
<p>Description: Description of the configuration.</p>
</td>
</tr>
<tr>
<td>
<code>documentationLink</code><br/>
<em>
string
</em>
</td>
<td>
<p>DocumentationLink: Configuration documentation link.</p>
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
<code>isConfigPendingRestart</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsConfigPendingRestart: Configuration is pending restart or not.</p>
</td>
</tr>
<tr>
<td>
<code>isDynamicConfig</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsDynamicConfig: Configuration dynamic or static.</p>
</td>
</tr>
<tr>
<td>
<code>isReadOnly</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsReadOnly: Configuration read-only or not.</p>
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
<code>source</code><br/>
<em>
string
</em>
</td>
<td>
<p>Source: Source of the configuration.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.SystemData_Status">
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
<tr>
<td>
<code>unit</code><br/>
<em>
string
</em>
</td>
<td>
<p>Unit: Configuration unit.</p>
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
<p>Value: Value of the configuration.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.Configuration_StatusARM">Configuration_StatusARM
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ConfigurationProperties_StatusARM">
ConfigurationProperties_StatusARM
</a>
</em>
</td>
<td>
<p>Properties: The properties of a configuration.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.SystemData_StatusARM">
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.DatabasePropertiesARM">DatabasePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServersDatabases_SpecARM">FlexibleServersDatabases_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/DatabaseProperties">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/DatabaseProperties</a></p>
</div>
<table>
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.DatabaseProperties_StatusARM">DatabaseProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Database_StatusARM">Database_StatusARM</a>)
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.Database_Status">Database_Status
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServersDatabase">FlexibleServersDatabase</a>)
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.SystemData_Status">
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.Database_StatusARM">Database_StatusARM
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.DatabaseProperties_StatusARM">
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.SystemData_StatusARM">
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.FirewallRulePropertiesARM">FirewallRulePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServersFirewallRules_SpecARM">FlexibleServersFirewallRules_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/FirewallRuleProperties">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/FirewallRuleProperties</a></p>
</div>
<table>
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.FirewallRuleProperties_StatusARM">FirewallRuleProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FirewallRule_StatusARM">FirewallRule_StatusARM</a>)
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.FirewallRule_Status">FirewallRule_Status
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServersFirewallRule">FlexibleServersFirewallRule</a>)
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.SystemData_Status">
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.FirewallRule_StatusARM">FirewallRule_StatusARM
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FirewallRuleProperties_StatusARM">
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.SystemData_StatusARM">
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServer">FlexibleServer
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/resourceDefinitions/flexibleServers">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/resourceDefinitions/flexibleServers</a></p>
</div>
<table>
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServers_Spec">
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
<p>AdministratorLoginPassword: The administrator login password (required for server creation).</p>
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
<p>AvailabilityZone: availability zone information of the server.</p>
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Backup">
Backup
</a>
</em>
</td>
<td>
<p>Backup: Backup properties of a server</p>
</td>
</tr>
<tr>
<td>
<code>createMode</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerPropertiesCreateMode">
ServerPropertiesCreateMode
</a>
</em>
</td>
<td>
<p>CreateMode: The mode to create a new PostgreSQL server.</p>
</td>
</tr>
<tr>
<td>
<code>highAvailability</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.HighAvailability">
HighAvailability
</a>
</em>
</td>
<td>
<p>HighAvailability: High availability properties of a server</p>
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.MaintenanceWindow">
MaintenanceWindow
</a>
</em>
</td>
<td>
<p>MaintenanceWindow: Maintenance window properties of a server.</p>
</td>
</tr>
<tr>
<td>
<code>network</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Network">
Network
</a>
</em>
</td>
<td>
<p>Network: Network properties of a server</p>
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
<code>pointInTimeUTC</code><br/>
<em>
string
</em>
</td>
<td>
<p>PointInTimeUTC: Restore point creation time (ISO8601 format), specifying the time to restore from. It&rsquo;s required when
&lsquo;createMode&rsquo; is &lsquo;PointInTimeRestore&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Sku">
Sku
</a>
</em>
</td>
<td>
<p>Sku: Sku information related properties of a server.</p>
</td>
</tr>
<tr>
<td>
<code>sourceServerResourceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>SourceServerResourceReference: The source server resource ID to restore from. It&rsquo;s required when &lsquo;createMode&rsquo; is
&lsquo;PointInTimeRestore&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>storage</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Storage">
Storage
</a>
</em>
</td>
<td>
<p>Storage: Storage properties of a server</p>
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerPropertiesVersion">
ServerPropertiesVersion
</a>
</em>
</td>
<td>
<p>Version: PostgreSQL Server version.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Server_Status">
Server_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServersConfiguration">FlexibleServersConfiguration
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/resourceDefinitions/flexibleServers_configurations">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/resourceDefinitions/flexibleServers_configurations</a></p>
</div>
<table>
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServersConfigurations_Spec">
FlexibleServersConfigurations_Spec
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
reference to a dbforpostgresql.azure.com/FlexibleServer resource</p>
</td>
</tr>
<tr>
<td>
<code>source</code><br/>
<em>
string
</em>
</td>
<td>
<p>Source: Source of the configuration.</p>
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
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Value of the configuration.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Configuration_Status">
Configuration_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServersConfigurationsSpecAPIVersion">FlexibleServersConfigurationsSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-06-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServersConfigurations_Spec">FlexibleServersConfigurations_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServersConfiguration">FlexibleServersConfiguration</a>)
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
reference to a dbforpostgresql.azure.com/FlexibleServer resource</p>
</td>
</tr>
<tr>
<td>
<code>source</code><br/>
<em>
string
</em>
</td>
<td>
<p>Source: Source of the configuration.</p>
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
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value: Value of the configuration.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServersConfigurations_SpecARM">FlexibleServersConfigurations_SpecARM
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
<p>Name: The name of the server configuration.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ConfigurationPropertiesARM">
ConfigurationPropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: The properties of a configuration.</p>
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServersDatabase">FlexibleServersDatabase
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/resourceDefinitions/flexibleServers_databases">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/resourceDefinitions/flexibleServers_databases</a></p>
</div>
<table>
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServersDatabases_Spec">
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
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a dbforpostgresql.azure.com/FlexibleServer resource</p>
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Database_Status">
Database_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServersDatabasesSpecAPIVersion">FlexibleServersDatabasesSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-06-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServersDatabases_Spec">FlexibleServersDatabases_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServersDatabase">FlexibleServersDatabase</a>)
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
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a dbforpostgresql.azure.com/FlexibleServer resource</p>
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServersDatabases_SpecARM">FlexibleServersDatabases_SpecARM
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.DatabasePropertiesARM">
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServersFirewallRule">FlexibleServersFirewallRule
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/resourceDefinitions/flexibleServers_firewallRules">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/resourceDefinitions/flexibleServers_firewallRules</a></p>
</div>
<table>
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServersFirewallRules_Spec">
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
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a dbforpostgresql.azure.com/FlexibleServer resource</p>
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FirewallRule_Status">
FirewallRule_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServersFirewallRulesSpecAPIVersion">FlexibleServersFirewallRulesSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-06-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServersFirewallRules_Spec">FlexibleServersFirewallRules_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServersFirewallRule">FlexibleServersFirewallRule</a>)
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
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a dbforpostgresql.azure.com/FlexibleServer resource</p>
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServersFirewallRules_SpecARM">FlexibleServersFirewallRules_SpecARM
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FirewallRulePropertiesARM">
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServersSpecAPIVersion">FlexibleServersSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-06-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServers_Spec">FlexibleServers_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServer">FlexibleServer</a>)
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
<p>AdministratorLoginPassword: The administrator login password (required for server creation).</p>
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
<p>AvailabilityZone: availability zone information of the server.</p>
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Backup">
Backup
</a>
</em>
</td>
<td>
<p>Backup: Backup properties of a server</p>
</td>
</tr>
<tr>
<td>
<code>createMode</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerPropertiesCreateMode">
ServerPropertiesCreateMode
</a>
</em>
</td>
<td>
<p>CreateMode: The mode to create a new PostgreSQL server.</p>
</td>
</tr>
<tr>
<td>
<code>highAvailability</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.HighAvailability">
HighAvailability
</a>
</em>
</td>
<td>
<p>HighAvailability: High availability properties of a server</p>
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.MaintenanceWindow">
MaintenanceWindow
</a>
</em>
</td>
<td>
<p>MaintenanceWindow: Maintenance window properties of a server.</p>
</td>
</tr>
<tr>
<td>
<code>network</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Network">
Network
</a>
</em>
</td>
<td>
<p>Network: Network properties of a server</p>
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
<code>pointInTimeUTC</code><br/>
<em>
string
</em>
</td>
<td>
<p>PointInTimeUTC: Restore point creation time (ISO8601 format), specifying the time to restore from. It&rsquo;s required when
&lsquo;createMode&rsquo; is &lsquo;PointInTimeRestore&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Sku">
Sku
</a>
</em>
</td>
<td>
<p>Sku: Sku information related properties of a server.</p>
</td>
</tr>
<tr>
<td>
<code>sourceServerResourceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>SourceServerResourceReference: The source server resource ID to restore from. It&rsquo;s required when &lsquo;createMode&rsquo; is
&lsquo;PointInTimeRestore&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>storage</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Storage">
Storage
</a>
</em>
</td>
<td>
<p>Storage: Storage properties of a server</p>
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerPropertiesVersion">
ServerPropertiesVersion
</a>
</em>
</td>
<td>
<p>Version: PostgreSQL Server version.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServers_SpecARM">FlexibleServers_SpecARM
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerPropertiesARM">
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.SkuARM">
SkuARM
</a>
</em>
</td>
<td>
<p>Sku: Sku information related properties of a server.</p>
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.HighAvailability">HighAvailability
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServers_Spec">FlexibleServers_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/HighAvailability">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/HighAvailability</a></p>
</div>
<table>
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.HighAvailabilityMode">
HighAvailabilityMode
</a>
</em>
</td>
<td>
<p>Mode: The HA mode for the server.</p>
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
<p>StandbyAvailabilityZone: availability zone information of the standby.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.HighAvailabilityARM">HighAvailabilityARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerPropertiesARM">ServerPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/HighAvailability">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/HighAvailability</a></p>
</div>
<table>
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.HighAvailabilityMode">
HighAvailabilityMode
</a>
</em>
</td>
<td>
<p>Mode: The HA mode for the server.</p>
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
<p>StandbyAvailabilityZone: availability zone information of the standby.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.HighAvailabilityMode">HighAvailabilityMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.HighAvailability">HighAvailability</a>, <a href="#dbforpostgresql.azure.com/v1alpha1api20210601.HighAvailabilityARM">HighAvailabilityARM</a>)
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
</tr><tr><td><p>&#34;ZoneRedundant&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.HighAvailabilityStatusMode">HighAvailabilityStatusMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.HighAvailability_Status">HighAvailability_Status</a>, <a href="#dbforpostgresql.azure.com/v1alpha1api20210601.HighAvailability_StatusARM">HighAvailability_StatusARM</a>)
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
</tr><tr><td><p>&#34;ZoneRedundant&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.HighAvailabilityStatusState">HighAvailabilityStatusState
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.HighAvailability_Status">HighAvailability_Status</a>, <a href="#dbforpostgresql.azure.com/v1alpha1api20210601.HighAvailability_StatusARM">HighAvailability_StatusARM</a>)
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
</tr><tr><td><p>&#34;ReplicatingData&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.HighAvailability_Status">HighAvailability_Status
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Server_Status">Server_Status</a>)
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.HighAvailabilityStatusMode">
HighAvailabilityStatusMode
</a>
</em>
</td>
<td>
<p>Mode: The HA mode for the server.</p>
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
<p>StandbyAvailabilityZone: availability zone information of the standby.</p>
</td>
</tr>
<tr>
<td>
<code>state</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.HighAvailabilityStatusState">
HighAvailabilityStatusState
</a>
</em>
</td>
<td>
<p>State: A state of a HA server that is visible to user.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.HighAvailability_StatusARM">HighAvailability_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerProperties_StatusARM">ServerProperties_StatusARM</a>)
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.HighAvailabilityStatusMode">
HighAvailabilityStatusMode
</a>
</em>
</td>
<td>
<p>Mode: The HA mode for the server.</p>
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
<p>StandbyAvailabilityZone: availability zone information of the standby.</p>
</td>
</tr>
<tr>
<td>
<code>state</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.HighAvailabilityStatusState">
HighAvailabilityStatusState
</a>
</em>
</td>
<td>
<p>State: A state of a HA server that is visible to user.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.MaintenanceWindow">MaintenanceWindow
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServers_Spec">FlexibleServers_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/MaintenanceWindow">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/MaintenanceWindow</a></p>
</div>
<table>
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.MaintenanceWindowARM">MaintenanceWindowARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerPropertiesARM">ServerPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/MaintenanceWindow">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/MaintenanceWindow</a></p>
</div>
<table>
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.MaintenanceWindow_Status">MaintenanceWindow_Status
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Server_Status">Server_Status</a>)
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.MaintenanceWindow_StatusARM">MaintenanceWindow_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerProperties_StatusARM">ServerProperties_StatusARM</a>)
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.Network">Network
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServers_Spec">FlexibleServers_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/Network">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/Network</a></p>
</div>
<table>
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
<p>DelegatedSubnetResourceReference: delegated subnet arm resource id.</p>
</td>
</tr>
<tr>
<td>
<code>privateDnsZoneArmResourceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>PrivateDnsZoneArmResourceReference: private dns zone arm resource id.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.NetworkARM">NetworkARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerPropertiesARM">ServerPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/Network">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/Network</a></p>
</div>
<table>
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
<code>privateDnsZoneArmResourceId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.NetworkStatusPublicNetworkAccess">NetworkStatusPublicNetworkAccess
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Network_Status">Network_Status</a>, <a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Network_StatusARM">Network_StatusARM</a>)
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.Network_Status">Network_Status
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Server_Status">Server_Status</a>)
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
<p>DelegatedSubnetResourceId: delegated subnet arm resource id.</p>
</td>
</tr>
<tr>
<td>
<code>privateDnsZoneArmResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateDnsZoneArmResourceId: private dns zone arm resource id.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.NetworkStatusPublicNetworkAccess">
NetworkStatusPublicNetworkAccess
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: public network access is enabled or not</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.Network_StatusARM">Network_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerProperties_StatusARM">ServerProperties_StatusARM</a>)
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
<p>DelegatedSubnetResourceId: delegated subnet arm resource id.</p>
</td>
</tr>
<tr>
<td>
<code>privateDnsZoneArmResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrivateDnsZoneArmResourceId: private dns zone arm resource id.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.NetworkStatusPublicNetworkAccess">
NetworkStatusPublicNetworkAccess
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: public network access is enabled or not</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.ServerPropertiesARM">ServerPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServers_SpecARM">FlexibleServers_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/ServerProperties">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/ServerProperties</a></p>
</div>
<table>
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
<p>AdministratorLoginPassword: The administrator login password (required for server creation).</p>
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
<p>AvailabilityZone: availability zone information of the server.</p>
</td>
</tr>
<tr>
<td>
<code>backup</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.BackupARM">
BackupARM
</a>
</em>
</td>
<td>
<p>Backup: Backup properties of a server</p>
</td>
</tr>
<tr>
<td>
<code>createMode</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerPropertiesCreateMode">
ServerPropertiesCreateMode
</a>
</em>
</td>
<td>
<p>CreateMode: The mode to create a new PostgreSQL server.</p>
</td>
</tr>
<tr>
<td>
<code>highAvailability</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.HighAvailabilityARM">
HighAvailabilityARM
</a>
</em>
</td>
<td>
<p>HighAvailability: High availability properties of a server</p>
</td>
</tr>
<tr>
<td>
<code>maintenanceWindow</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.MaintenanceWindowARM">
MaintenanceWindowARM
</a>
</em>
</td>
<td>
<p>MaintenanceWindow: Maintenance window properties of a server.</p>
</td>
</tr>
<tr>
<td>
<code>network</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.NetworkARM">
NetworkARM
</a>
</em>
</td>
<td>
<p>Network: Network properties of a server</p>
</td>
</tr>
<tr>
<td>
<code>pointInTimeUTC</code><br/>
<em>
string
</em>
</td>
<td>
<p>PointInTimeUTC: Restore point creation time (ISO8601 format), specifying the time to restore from. It&rsquo;s required when
&lsquo;createMode&rsquo; is &lsquo;PointInTimeRestore&rsquo;.</p>
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
</td>
</tr>
<tr>
<td>
<code>storage</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.StorageARM">
StorageARM
</a>
</em>
</td>
<td>
<p>Storage: Storage properties of a server</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerPropertiesVersion">
ServerPropertiesVersion
</a>
</em>
</td>
<td>
<p>Version: PostgreSQL Server version.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.ServerPropertiesCreateMode">ServerPropertiesCreateMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServers_Spec">FlexibleServers_Spec</a>, <a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerPropertiesARM">ServerPropertiesARM</a>)
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
<tbody><tr><td><p>&#34;Create&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Default&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;PointInTimeRestore&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Update&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.ServerPropertiesStatusCreateMode">ServerPropertiesStatusCreateMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerProperties_StatusARM">ServerProperties_StatusARM</a>, <a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Server_Status">Server_Status</a>)
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
<tbody><tr><td><p>&#34;Create&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Default&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;PointInTimeRestore&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Update&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.ServerPropertiesStatusState">ServerPropertiesStatusState
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerProperties_StatusARM">ServerProperties_StatusARM</a>, <a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Server_Status">Server_Status</a>)
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.ServerPropertiesVersion">ServerPropertiesVersion
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServers_Spec">FlexibleServers_Spec</a>, <a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerPropertiesARM">ServerPropertiesARM</a>)
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
<tbody><tr><td><p>&#34;11&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;12&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;13&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.ServerProperties_StatusARM">ServerProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Server_StatusARM">Server_StatusARM</a>)
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
<p>AvailabilityZone: availability zone information of the server.</p>
</td>
</tr>
<tr>
<td>
<code>backup</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Backup_StatusARM">
Backup_StatusARM
</a>
</em>
</td>
<td>
<p>Backup: Backup properties of a server.</p>
</td>
</tr>
<tr>
<td>
<code>createMode</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerPropertiesStatusCreateMode">
ServerPropertiesStatusCreateMode
</a>
</em>
</td>
<td>
<p>CreateMode: The mode to create a new PostgreSQL server.</p>
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.HighAvailability_StatusARM">
HighAvailability_StatusARM
</a>
</em>
</td>
<td>
<p>HighAvailability: High availability properties of a server.</p>
</td>
</tr>
<tr>
<td>
<code>maintenanceWindow</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.MaintenanceWindow_StatusARM">
MaintenanceWindow_StatusARM
</a>
</em>
</td>
<td>
<p>MaintenanceWindow: Maintenance window properties of a server.</p>
</td>
</tr>
<tr>
<td>
<code>minorVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>MinorVersion: The minor version of the server.</p>
</td>
</tr>
<tr>
<td>
<code>network</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Network_StatusARM">
Network_StatusARM
</a>
</em>
</td>
<td>
<p>Network: Network properties of a server.</p>
</td>
</tr>
<tr>
<td>
<code>pointInTimeUTC</code><br/>
<em>
string
</em>
</td>
<td>
<p>PointInTimeUTC: Restore point creation time (ISO8601 format), specifying the time to restore from. It&rsquo;s required when
&lsquo;createMode&rsquo; is &lsquo;PointInTimeRestore&rsquo;.</p>
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
<p>SourceServerResourceId: The source server resource ID to restore from. It&rsquo;s required when &lsquo;createMode&rsquo; is
&lsquo;PointInTimeRestore&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>state</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerPropertiesStatusState">
ServerPropertiesStatusState
</a>
</em>
</td>
<td>
<p>State: A state of a server that is visible to user.</p>
</td>
</tr>
<tr>
<td>
<code>storage</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Storage_StatusARM">
Storage_StatusARM
</a>
</em>
</td>
<td>
<p>Storage: Storage properties of a server.</p>
</td>
</tr>
<tr>
<td>
<code>version</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerVersion_Status">
ServerVersion_Status
</a>
</em>
</td>
<td>
<p>Version: PostgreSQL Server version.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.ServerVersion_Status">ServerVersion_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerProperties_StatusARM">ServerProperties_StatusARM</a>, <a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Server_Status">Server_Status</a>)
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
<tbody><tr><td><p>&#34;11&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;12&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;13&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.Server_Status">Server_Status
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServer">FlexibleServer</a>)
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
<p>AvailabilityZone: availability zone information of the server.</p>
</td>
</tr>
<tr>
<td>
<code>backup</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Backup_Status">
Backup_Status
</a>
</em>
</td>
<td>
<p>Backup: Backup properties of a server.</p>
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerPropertiesStatusCreateMode">
ServerPropertiesStatusCreateMode
</a>
</em>
</td>
<td>
<p>CreateMode: The mode to create a new PostgreSQL server.</p>
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.HighAvailability_Status">
HighAvailability_Status
</a>
</em>
</td>
<td>
<p>HighAvailability: High availability properties of a server.</p>
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.MaintenanceWindow_Status">
MaintenanceWindow_Status
</a>
</em>
</td>
<td>
<p>MaintenanceWindow: Maintenance window properties of a server.</p>
</td>
</tr>
<tr>
<td>
<code>minorVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>MinorVersion: The minor version of the server.</p>
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Network_Status">
Network_Status
</a>
</em>
</td>
<td>
<p>Network: Network properties of a server.</p>
</td>
</tr>
<tr>
<td>
<code>pointInTimeUTC</code><br/>
<em>
string
</em>
</td>
<td>
<p>PointInTimeUTC: Restore point creation time (ISO8601 format), specifying the time to restore from. It&rsquo;s required when
&lsquo;createMode&rsquo; is &lsquo;PointInTimeRestore&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Sku_Status">
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
<p>SourceServerResourceId: The source server resource ID to restore from. It&rsquo;s required when &lsquo;createMode&rsquo; is
&lsquo;PointInTimeRestore&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>state</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerPropertiesStatusState">
ServerPropertiesStatusState
</a>
</em>
</td>
<td>
<p>State: A state of a server that is visible to user.</p>
</td>
</tr>
<tr>
<td>
<code>storage</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Storage_Status">
Storage_Status
</a>
</em>
</td>
<td>
<p>Storage: Storage properties of a server.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.SystemData_Status">
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerVersion_Status">
ServerVersion_Status
</a>
</em>
</td>
<td>
<p>Version: PostgreSQL Server version.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.Server_StatusARM">Server_StatusARM
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerProperties_StatusARM">
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Sku_StatusARM">
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.SystemData_StatusARM">
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.Sku">Sku
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServers_Spec">FlexibleServers_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/Sku">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/Sku</a></p>
</div>
<table>
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
<p>Name: The name of the sku, typically, tier + family + cores, e.g. Standard_D4s_v3.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.SkuTier">
SkuTier
</a>
</em>
</td>
<td>
<p>Tier: The tier of the particular SKU, e.g. Burstable.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.SkuARM">SkuARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServers_SpecARM">FlexibleServers_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/Sku">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/Sku</a></p>
</div>
<table>
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
<p>Name: The name of the sku, typically, tier + family + cores, e.g. Standard_D4s_v3.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.SkuTier">
SkuTier
</a>
</em>
</td>
<td>
<p>Tier: The tier of the particular SKU, e.g. Burstable.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.SkuStatusTier">SkuStatusTier
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Sku_Status">Sku_Status</a>, <a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Sku_StatusARM">Sku_StatusARM</a>)
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.SkuTier">SkuTier
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Sku">Sku</a>, <a href="#dbforpostgresql.azure.com/v1alpha1api20210601.SkuARM">SkuARM</a>)
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.Sku_Status">Sku_Status
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Server_Status">Server_Status</a>)
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
<p>Name: The name of the sku, typically, tier + family + cores, e.g. Standard_D4s_v3.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.SkuStatusTier">
SkuStatusTier
</a>
</em>
</td>
<td>
<p>Tier: The tier of the particular SKU, e.g. Burstable.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.Sku_StatusARM">Sku_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Server_StatusARM">Server_StatusARM</a>)
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
<p>Name: The name of the sku, typically, tier + family + cores, e.g. Standard_D4s_v3.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.SkuStatusTier">
SkuStatusTier
</a>
</em>
</td>
<td>
<p>Tier: The tier of the particular SKU, e.g. Burstable.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.Storage">Storage
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FlexibleServers_Spec">FlexibleServers_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/Storage">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/Storage</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>storageSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
<p>StorageSizeGB: Max storage allowed for a server.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.StorageARM">StorageARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerPropertiesARM">ServerPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/Storage">https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/definitions/Storage</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>storageSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
<p>StorageSizeGB: Max storage allowed for a server.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.Storage_Status">Storage_Status
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Server_Status">Server_Status</a>)
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
<code>storageSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
<p>StorageSizeGB: Max storage allowed for a server.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.Storage_StatusARM">Storage_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.ServerProperties_StatusARM">ServerProperties_StatusARM</a>)
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
<code>storageSizeGB</code><br/>
<em>
int
</em>
</td>
<td>
<p>StorageSizeGB: Max storage allowed for a server.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.SystemDataStatusCreatedByType">SystemDataStatusCreatedByType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.SystemData_Status">SystemData_Status</a>, <a href="#dbforpostgresql.azure.com/v1alpha1api20210601.SystemData_StatusARM">SystemData_StatusARM</a>)
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.SystemDataStatusLastModifiedByType">SystemDataStatusLastModifiedByType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.SystemData_Status">SystemData_Status</a>, <a href="#dbforpostgresql.azure.com/v1alpha1api20210601.SystemData_StatusARM">SystemData_StatusARM</a>)
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.SystemData_Status">SystemData_Status
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Configuration_Status">Configuration_Status</a>, <a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Database_Status">Database_Status</a>, <a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FirewallRule_Status">FirewallRule_Status</a>, <a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Server_Status">Server_Status</a>)
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.SystemDataStatusCreatedByType">
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.SystemDataStatusLastModifiedByType">
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
<h3 id="dbforpostgresql.azure.com/v1alpha1api20210601.SystemData_StatusARM">SystemData_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Configuration_StatusARM">Configuration_StatusARM</a>, <a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Database_StatusARM">Database_StatusARM</a>, <a href="#dbforpostgresql.azure.com/v1alpha1api20210601.FirewallRule_StatusARM">FirewallRule_StatusARM</a>, <a href="#dbforpostgresql.azure.com/v1alpha1api20210601.Server_StatusARM">Server_StatusARM</a>)
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.SystemDataStatusCreatedByType">
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
<a href="#dbforpostgresql.azure.com/v1alpha1api20210601.SystemDataStatusLastModifiedByType">
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
