---
title: dbforpostgresql.azure.com/v1api20230601preview
---
<h2 id="dbforpostgresql.azure.com/v1api20230601preview">dbforpostgresql.azure.com/v1api20230601preview</h2>
<div>
<p>Package v1api20230601preview contains API Schema definitions for the dbforpostgresql v1api20230601preview API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.APIVersion">APIVersion
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
<tbody><tr><td><p>&#34;2023-06-01-preview&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.AuthConfig">AuthConfig
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_Spec">FlexibleServer_Spec</a>)
</p>
<div>
<p>Authentication configuration properties of a server</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>activeDirectoryAuth</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.AuthConfig_ActiveDirectoryAuth">
AuthConfig_ActiveDirectoryAuth
</a>
</em>
</td>
<td>
<p>ActiveDirectoryAuth: If Enabled, Azure Active Directory authentication is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>passwordAuth</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.AuthConfig_PasswordAuth">
AuthConfig_PasswordAuth
</a>
</em>
</td>
<td>
<p>PasswordAuth: If Enabled, Password authentication is enabled.</p>
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
<p>TenantId: Tenant id of the server.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.AuthConfig_ActiveDirectoryAuth">AuthConfig_ActiveDirectoryAuth
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.AuthConfig">AuthConfig</a>)
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.AuthConfig_ActiveDirectoryAuth_STATUS">AuthConfig_ActiveDirectoryAuth_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.AuthConfig_STATUS">AuthConfig_STATUS</a>)
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.AuthConfig_PasswordAuth">AuthConfig_PasswordAuth
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.AuthConfig">AuthConfig</a>)
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.AuthConfig_PasswordAuth_STATUS">AuthConfig_PasswordAuth_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.AuthConfig_STATUS">AuthConfig_STATUS</a>)
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.AuthConfig_STATUS">AuthConfig_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_STATUS">FlexibleServer_STATUS</a>)
</p>
<div>
<p>Authentication configuration properties of a server</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>activeDirectoryAuth</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.AuthConfig_ActiveDirectoryAuth_STATUS">
AuthConfig_ActiveDirectoryAuth_STATUS
</a>
</em>
</td>
<td>
<p>ActiveDirectoryAuth: If Enabled, Azure Active Directory authentication is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>passwordAuth</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.AuthConfig_PasswordAuth_STATUS">
AuthConfig_PasswordAuth_STATUS
</a>
</em>
</td>
<td>
<p>PasswordAuth: If Enabled, Password authentication is enabled.</p>
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
<p>TenantId: Tenant id of the server.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Backup">Backup
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_Spec">FlexibleServer_Spec</a>)
</p>
<div>
<p>Backup properties of a server</p>
</div>
<table>
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Backup_GeoRedundantBackup">
Backup_GeoRedundantBackup
</a>
</em>
</td>
<td>
<p>GeoRedundantBackup: A value indicating whether Geo-Redundant backup is enabled on the server.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Backup_GeoRedundantBackup">Backup_GeoRedundantBackup
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.Backup">Backup</a>)
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Backup_GeoRedundantBackup_STATUS">Backup_GeoRedundantBackup_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.Backup_STATUS">Backup_STATUS</a>)
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Backup_STATUS">Backup_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_STATUS">FlexibleServer_STATUS</a>)
</p>
<div>
<p>Backup properties of a server</p>
</div>
<table>
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Backup_GeoRedundantBackup_STATUS">
Backup_GeoRedundantBackup_STATUS
</a>
</em>
</td>
<td>
<p>GeoRedundantBackup: A value indicating whether Geo-Redundant backup is enabled on the server.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.ConfigurationProperties_DataType_STATUS">ConfigurationProperties_DataType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersConfiguration_STATUS">FlexibleServersConfiguration_STATUS</a>)
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.DataEncryption">DataEncryption
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_Spec">FlexibleServer_Spec</a>)
</p>
<div>
<p>Data encryption properties of a server</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>geoBackupEncryptionKeyStatus</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.DataEncryption_GeoBackupEncryptionKeyStatus">
DataEncryption_GeoBackupEncryptionKeyStatus
</a>
</em>
</td>
<td>
<p>GeoBackupEncryptionKeyStatus: Geo-backup encryption key status for Data encryption enabled server.</p>
</td>
</tr>
<tr>
<td>
<code>geoBackupKeyURI</code><br/>
<em>
string
</em>
</td>
<td>
<p>GeoBackupKeyURI: URI for the key in keyvault for data encryption for geo-backup of server.</p>
</td>
</tr>
<tr>
<td>
<code>geoBackupKeyURIFromConfig</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ConfigMapReference">
genruntime.ConfigMapReference
</a>
</em>
</td>
<td>
<p>GeoBackupKeyURIFromConfig: URI for the key in keyvault for data encryption for geo-backup of server.</p>
</td>
</tr>
<tr>
<td>
<code>geoBackupUserAssignedIdentityReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>GeoBackupUserAssignedIdentityReference: Resource Id for the User assigned identity to be used for data encryption for
geo-backup of server.</p>
</td>
</tr>
<tr>
<td>
<code>primaryEncryptionKeyStatus</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.DataEncryption_PrimaryEncryptionKeyStatus">
DataEncryption_PrimaryEncryptionKeyStatus
</a>
</em>
</td>
<td>
<p>PrimaryEncryptionKeyStatus: Primary encryption key status for Data encryption enabled server.</p>
</td>
</tr>
<tr>
<td>
<code>primaryKeyURI</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrimaryKeyURI: URI for the key in keyvault for data encryption of the primary server.</p>
</td>
</tr>
<tr>
<td>
<code>primaryKeyURIFromConfig</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ConfigMapReference">
genruntime.ConfigMapReference
</a>
</em>
</td>
<td>
<p>PrimaryKeyURIFromConfig: URI for the key in keyvault for data encryption of the primary server.</p>
</td>
</tr>
<tr>
<td>
<code>primaryUserAssignedIdentityReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>PrimaryUserAssignedIdentityReference: Resource Id for the User assigned identity to be used for data encryption of the
primary server.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.DataEncryption_Type">
DataEncryption_Type
</a>
</em>
</td>
<td>
<p>Type: Data encryption type to depict if it is System Managed vs Azure Key vault.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.DataEncryption_GeoBackupEncryptionKeyStatus">DataEncryption_GeoBackupEncryptionKeyStatus
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.DataEncryption">DataEncryption</a>)
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
<tbody><tr><td><p>&#34;Invalid&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Valid&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.DataEncryption_GeoBackupEncryptionKeyStatus_STATUS">DataEncryption_GeoBackupEncryptionKeyStatus_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.DataEncryption_STATUS">DataEncryption_STATUS</a>)
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
<tbody><tr><td><p>&#34;Invalid&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Valid&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.DataEncryption_PrimaryEncryptionKeyStatus">DataEncryption_PrimaryEncryptionKeyStatus
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.DataEncryption">DataEncryption</a>)
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
<tbody><tr><td><p>&#34;Invalid&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Valid&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.DataEncryption_PrimaryEncryptionKeyStatus_STATUS">DataEncryption_PrimaryEncryptionKeyStatus_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.DataEncryption_STATUS">DataEncryption_STATUS</a>)
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
<tbody><tr><td><p>&#34;Invalid&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Valid&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.DataEncryption_STATUS">DataEncryption_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_STATUS">FlexibleServer_STATUS</a>)
</p>
<div>
<p>Data encryption properties of a server</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>geoBackupEncryptionKeyStatus</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.DataEncryption_GeoBackupEncryptionKeyStatus_STATUS">
DataEncryption_GeoBackupEncryptionKeyStatus_STATUS
</a>
</em>
</td>
<td>
<p>GeoBackupEncryptionKeyStatus: Geo-backup encryption key status for Data encryption enabled server.</p>
</td>
</tr>
<tr>
<td>
<code>geoBackupKeyURI</code><br/>
<em>
string
</em>
</td>
<td>
<p>GeoBackupKeyURI: URI for the key in keyvault for data encryption for geo-backup of server.</p>
</td>
</tr>
<tr>
<td>
<code>geoBackupUserAssignedIdentityId</code><br/>
<em>
string
</em>
</td>
<td>
<p>GeoBackupUserAssignedIdentityId: Resource Id for the User assigned identity to be used for data encryption for
geo-backup of server.</p>
</td>
</tr>
<tr>
<td>
<code>primaryEncryptionKeyStatus</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.DataEncryption_PrimaryEncryptionKeyStatus_STATUS">
DataEncryption_PrimaryEncryptionKeyStatus_STATUS
</a>
</em>
</td>
<td>
<p>PrimaryEncryptionKeyStatus: Primary encryption key status for Data encryption enabled server.</p>
</td>
</tr>
<tr>
<td>
<code>primaryKeyURI</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrimaryKeyURI: URI for the key in keyvault for data encryption of the primary server.</p>
</td>
</tr>
<tr>
<td>
<code>primaryUserAssignedIdentityId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrimaryUserAssignedIdentityId: Resource Id for the User assigned identity to be used for data encryption of the primary
server.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.DataEncryption_Type_STATUS">
DataEncryption_Type_STATUS
</a>
</em>
</td>
<td>
<p>Type: Data encryption type to depict if it is System Managed vs Azure Key vault.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.DataEncryption_Type">DataEncryption_Type
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.DataEncryption">DataEncryption</a>)
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
<tbody><tr><td><p>&#34;AzureKeyVault&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SystemManaged&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.DataEncryption_Type_STATUS">DataEncryption_Type_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.DataEncryption_STATUS">DataEncryption_STATUS</a>)
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
<tbody><tr><td><p>&#34;AzureKeyVault&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SystemManaged&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer">FlexibleServer
</h3>
<div>
<p>Generator information:
- Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/FlexibleServers.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.DBforPostgreSQL/&#x200b;flexibleServers/&#x200b;{serverName}</&#x200b;p>
</div>
<table>
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_Spec">
FlexibleServer_Spec
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
<code>authConfig</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.AuthConfig">
AuthConfig
</a>
</em>
</td>
<td>
<p>AuthConfig: AuthConfig properties of a server.</p>
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Backup">
Backup
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.ServerProperties_CreateMode">
ServerProperties_CreateMode
</a>
</em>
</td>
<td>
<p>CreateMode: The mode to create a new PostgreSQL server.</p>
</td>
</tr>
<tr>
<td>
<code>dataEncryption</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.DataEncryption">
DataEncryption
</a>
</em>
</td>
<td>
<p>DataEncryption: Data encryption properties of a server.</p>
</td>
</tr>
<tr>
<td>
<code>highAvailability</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.HighAvailability">
HighAvailability
</a>
</em>
</td>
<td>
<p>HighAvailability: High availability properties of a server.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.UserAssignedIdentity">
UserAssignedIdentity
</a>
</em>
</td>
<td>
<p>Identity: Describes the identity of the application.</p>
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.MaintenanceWindow">
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Network">
Network
</a>
</em>
</td>
<td>
<p>Network: Network properties of a server. This Network property is required to be passed only in case you want the server
to be Private access server.</p>
</td>
</tr>
<tr>
<td>
<code>operatorSpec</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServerOperatorSpec">
FlexibleServerOperatorSpec
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
<code>pointInTimeUTC</code><br/>
<em>
string
</em>
</td>
<td>
<p>PointInTimeUTC: Restore point creation time (ISO8601 format), specifying the time to restore from. It&rsquo;s required when
&lsquo;createMode&rsquo; is &lsquo;PointInTimeRestore&rsquo; or &lsquo;GeoRestore&rsquo; or &lsquo;ReviveDropped&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>replica</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Replica">
Replica
</a>
</em>
</td>
<td>
<p>Replica: Replica properties of a server. These Replica properties are required to be passed only in case you want to
Promote a server.</p>
</td>
</tr>
<tr>
<td>
<code>replicationRole</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.ReplicationRole">
ReplicationRole
</a>
</em>
</td>
<td>
<p>ReplicationRole: Replication role of the server</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Sku">
Sku
</a>
</em>
</td>
<td>
<p>Sku: The SKU (pricing tier) of the server.</p>
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
&lsquo;PointInTimeRestore&rsquo; or &lsquo;GeoRestore&rsquo; or &lsquo;Replica&rsquo; or &lsquo;ReviveDropped&rsquo;. This property is returned only for Replica server</p>
</td>
</tr>
<tr>
<td>
<code>storage</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Storage">
Storage
</a>
</em>
</td>
<td>
<p>Storage: Storage properties of a server.</p>
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
<code>version</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.ServerVersion">
ServerVersion
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_STATUS">
FlexibleServer_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.FlexibleServerOperatorConfigMaps">FlexibleServerOperatorConfigMaps
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServerOperatorSpec">FlexibleServerOperatorSpec</a>)
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
<code>fullyQualifiedDomainName</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ConfigMapDestination">
genruntime.ConfigMapDestination
</a>
</em>
</td>
<td>
<p>FullyQualifiedDomainName: indicates where the FullyQualifiedDomainName config map should be placed. If omitted, no
config map will be created.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.FlexibleServerOperatorSecrets">FlexibleServerOperatorSecrets
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServerOperatorSpec">FlexibleServerOperatorSpec</a>)
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
<code>fullyQualifiedDomainName</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretDestination">
genruntime.SecretDestination
</a>
</em>
</td>
<td>
<p>FullyQualifiedDomainName: indicates where the FullyQualifiedDomainName secret should be placed. If omitted, the secret
will not be retrieved from Azure.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.FlexibleServerOperatorSpec">FlexibleServerOperatorSpec
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_Spec">FlexibleServer_Spec</a>)
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
<code>configMapExpressions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#DestinationExpression">
[]genruntime/core.DestinationExpression
</a>
</em>
</td>
<td>
<p>ConfigMapExpressions: configures where to place operator written dynamic ConfigMaps (created with CEL expressions).</p>
</td>
</tr>
<tr>
<td>
<code>configMaps</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServerOperatorConfigMaps">
FlexibleServerOperatorConfigMaps
</a>
</em>
</td>
<td>
<p>ConfigMaps: configures where to place operator written ConfigMaps.</p>
</td>
</tr>
<tr>
<td>
<code>secretExpressions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#DestinationExpression">
[]genruntime/core.DestinationExpression
</a>
</em>
</td>
<td>
<p>SecretExpressions: configures where to place operator written dynamic secrets (created with CEL expressions).</p>
</td>
</tr>
<tr>
<td>
<code>secrets</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServerOperatorSecrets">
FlexibleServerOperatorSecrets
</a>
</em>
</td>
<td>
<p>Secrets: configures where to place Azure generated secrets.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_STATUS">FlexibleServer_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer">FlexibleServer</a>)
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
<code>authConfig</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.AuthConfig_STATUS">
AuthConfig_STATUS
</a>
</em>
</td>
<td>
<p>AuthConfig: AuthConfig properties of a server.</p>
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Backup_STATUS">
Backup_STATUS
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.ServerProperties_CreateMode_STATUS">
ServerProperties_CreateMode_STATUS
</a>
</em>
</td>
<td>
<p>CreateMode: The mode to create a new PostgreSQL server.</p>
</td>
</tr>
<tr>
<td>
<code>dataEncryption</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.DataEncryption_STATUS">
DataEncryption_STATUS
</a>
</em>
</td>
<td>
<p>DataEncryption: Data encryption properties of a server.</p>
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.HighAvailability_STATUS">
HighAvailability_STATUS
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
<p>Id: Fully qualified resource ID for the resource. E.g.
&ldquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}&rdquo;</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.UserAssignedIdentity_STATUS">
UserAssignedIdentity_STATUS
</a>
</em>
</td>
<td>
<p>Identity: Describes the identity of the application.</p>
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.MaintenanceWindow_STATUS">
MaintenanceWindow_STATUS
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Network_STATUS">
Network_STATUS
</a>
</em>
</td>
<td>
<p>Network: Network properties of a server. This Network property is required to be passed only in case you want the server
to be Private access server.</p>
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
&lsquo;createMode&rsquo; is &lsquo;PointInTimeRestore&rsquo; or &lsquo;GeoRestore&rsquo; or &lsquo;ReviveDropped&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpointConnections</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.PrivateEndpointConnection_STATUS">
[]PrivateEndpointConnection_STATUS
</a>
</em>
</td>
<td>
<p>PrivateEndpointConnections: List of private endpoint connections associated with the specified resource.</p>
</td>
</tr>
<tr>
<td>
<code>replica</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Replica_STATUS">
Replica_STATUS
</a>
</em>
</td>
<td>
<p>Replica: Replica properties of a server. These Replica properties are required to be passed only in case you want to
Promote a server.</p>
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
<p>ReplicaCapacity: Replicas allowed for a server.</p>
</td>
</tr>
<tr>
<td>
<code>replicationRole</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.ReplicationRole_STATUS">
ReplicationRole_STATUS
</a>
</em>
</td>
<td>
<p>ReplicationRole: Replication role of the server</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Sku_STATUS">
Sku_STATUS
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
&lsquo;PointInTimeRestore&rsquo; or &lsquo;GeoRestore&rsquo; or &lsquo;Replica&rsquo; or &lsquo;ReviveDropped&rsquo;. This property is returned only for Replica server</p>
</td>
</tr>
<tr>
<td>
<code>state</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.ServerProperties_State_STATUS">
ServerProperties_State_STATUS
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Storage_STATUS">
Storage_STATUS
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.SystemData_STATUS">
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
<code>version</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.ServerVersion_STATUS">
ServerVersion_STATUS
</a>
</em>
</td>
<td>
<p>Version: PostgreSQL Server version.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_Spec">FlexibleServer_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer">FlexibleServer</a>)
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
<code>authConfig</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.AuthConfig">
AuthConfig
</a>
</em>
</td>
<td>
<p>AuthConfig: AuthConfig properties of a server.</p>
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Backup">
Backup
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.ServerProperties_CreateMode">
ServerProperties_CreateMode
</a>
</em>
</td>
<td>
<p>CreateMode: The mode to create a new PostgreSQL server.</p>
</td>
</tr>
<tr>
<td>
<code>dataEncryption</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.DataEncryption">
DataEncryption
</a>
</em>
</td>
<td>
<p>DataEncryption: Data encryption properties of a server.</p>
</td>
</tr>
<tr>
<td>
<code>highAvailability</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.HighAvailability">
HighAvailability
</a>
</em>
</td>
<td>
<p>HighAvailability: High availability properties of a server.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.UserAssignedIdentity">
UserAssignedIdentity
</a>
</em>
</td>
<td>
<p>Identity: Describes the identity of the application.</p>
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.MaintenanceWindow">
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Network">
Network
</a>
</em>
</td>
<td>
<p>Network: Network properties of a server. This Network property is required to be passed only in case you want the server
to be Private access server.</p>
</td>
</tr>
<tr>
<td>
<code>operatorSpec</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServerOperatorSpec">
FlexibleServerOperatorSpec
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
<code>pointInTimeUTC</code><br/>
<em>
string
</em>
</td>
<td>
<p>PointInTimeUTC: Restore point creation time (ISO8601 format), specifying the time to restore from. It&rsquo;s required when
&lsquo;createMode&rsquo; is &lsquo;PointInTimeRestore&rsquo; or &lsquo;GeoRestore&rsquo; or &lsquo;ReviveDropped&rsquo;.</p>
</td>
</tr>
<tr>
<td>
<code>replica</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Replica">
Replica
</a>
</em>
</td>
<td>
<p>Replica: Replica properties of a server. These Replica properties are required to be passed only in case you want to
Promote a server.</p>
</td>
</tr>
<tr>
<td>
<code>replicationRole</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.ReplicationRole">
ReplicationRole
</a>
</em>
</td>
<td>
<p>ReplicationRole: Replication role of the server</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Sku">
Sku
</a>
</em>
</td>
<td>
<p>Sku: The SKU (pricing tier) of the server.</p>
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
&lsquo;PointInTimeRestore&rsquo; or &lsquo;GeoRestore&rsquo; or &lsquo;Replica&rsquo; or &lsquo;ReviveDropped&rsquo;. This property is returned only for Replica server</p>
</td>
</tr>
<tr>
<td>
<code>storage</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Storage">
Storage
</a>
</em>
</td>
<td>
<p>Storage: Storage properties of a server.</p>
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
<code>version</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.ServerVersion">
ServerVersion
</a>
</em>
</td>
<td>
<p>Version: PostgreSQL Server version.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersConfiguration">FlexibleServersConfiguration
</h3>
<div>
<p>Generator information:
- Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/Configuration.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.DBforPostgreSQL/&#x200b;flexibleServers/&#x200b;{serverName}/&#x200b;configurations/&#x200b;{configurationName}</&#x200b;p>
</div>
<table>
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersConfiguration_Spec">
FlexibleServersConfiguration_Spec
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
<code>operatorSpec</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersConfigurationOperatorSpec">
FlexibleServersConfigurationOperatorSpec
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersConfiguration_STATUS">
FlexibleServersConfiguration_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersConfigurationOperatorSpec">FlexibleServersConfigurationOperatorSpec
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersConfiguration_Spec">FlexibleServersConfiguration_Spec</a>)
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
<code>configMapExpressions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#DestinationExpression">
[]genruntime/core.DestinationExpression
</a>
</em>
</td>
<td>
<p>ConfigMapExpressions: configures where to place operator written dynamic ConfigMaps (created with CEL expressions).</p>
</td>
</tr>
<tr>
<td>
<code>secretExpressions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#DestinationExpression">
[]genruntime/core.DestinationExpression
</a>
</em>
</td>
<td>
<p>SecretExpressions: configures where to place operator written dynamic secrets (created with CEL expressions).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersConfiguration_STATUS">FlexibleServersConfiguration_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersConfiguration">FlexibleServersConfiguration</a>)
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.ConfigurationProperties_DataType_STATUS">
ConfigurationProperties_DataType_STATUS
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
<p>Id: Fully qualified resource ID for the resource. E.g.
&ldquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}&rdquo;</&#x200b;p>
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.SystemData_STATUS">
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersConfiguration_Spec">FlexibleServersConfiguration_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersConfiguration">FlexibleServersConfiguration</a>)
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
<code>operatorSpec</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersConfigurationOperatorSpec">
FlexibleServersConfigurationOperatorSpec
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersDatabase">FlexibleServersDatabase
</h3>
<div>
<p>Generator information:
- Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/Databases.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.DBforPostgreSQL/&#x200b;flexibleServers/&#x200b;{serverName}/&#x200b;databases/&#x200b;{databaseName}</&#x200b;p>
</div>
<table>
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersDatabase_Spec">
FlexibleServersDatabase_Spec
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
<code>operatorSpec</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersDatabaseOperatorSpec">
FlexibleServersDatabaseOperatorSpec
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
reference to a dbforpostgresql.azure.com/FlexibleServer resource</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersDatabase_STATUS">
FlexibleServersDatabase_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersDatabaseOperatorSpec">FlexibleServersDatabaseOperatorSpec
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersDatabase_Spec">FlexibleServersDatabase_Spec</a>)
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
<code>configMapExpressions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#DestinationExpression">
[]genruntime/core.DestinationExpression
</a>
</em>
</td>
<td>
<p>ConfigMapExpressions: configures where to place operator written dynamic ConfigMaps (created with CEL expressions).</p>
</td>
</tr>
<tr>
<td>
<code>secretExpressions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#DestinationExpression">
[]genruntime/core.DestinationExpression
</a>
</em>
</td>
<td>
<p>SecretExpressions: configures where to place operator written dynamic secrets (created with CEL expressions).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersDatabase_STATUS">FlexibleServersDatabase_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersDatabase">FlexibleServersDatabase</a>)
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
<p>Id: Fully qualified resource ID for the resource. E.g.
&ldquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}&rdquo;</&#x200b;p>
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.SystemData_STATUS">
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersDatabase_Spec">FlexibleServersDatabase_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersDatabase">FlexibleServersDatabase</a>)
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
<code>operatorSpec</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersDatabaseOperatorSpec">
FlexibleServersDatabaseOperatorSpec
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
reference to a dbforpostgresql.azure.com/FlexibleServer resource</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersFirewallRule">FlexibleServersFirewallRule
</h3>
<div>
<p>Generator information:
- Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/FirewallRules.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.DBforPostgreSQL/&#x200b;flexibleServers/&#x200b;{serverName}/&#x200b;firewallRules/&#x200b;{firewallRuleName}</&#x200b;p>
</div>
<table>
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersFirewallRule_Spec">
FlexibleServersFirewallRule_Spec
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
<code>operatorSpec</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersFirewallRuleOperatorSpec">
FlexibleServersFirewallRuleOperatorSpec
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
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersFirewallRule_STATUS">
FlexibleServersFirewallRule_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersFirewallRuleOperatorSpec">FlexibleServersFirewallRuleOperatorSpec
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersFirewallRule_Spec">FlexibleServersFirewallRule_Spec</a>)
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
<code>configMapExpressions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#DestinationExpression">
[]genruntime/core.DestinationExpression
</a>
</em>
</td>
<td>
<p>ConfigMapExpressions: configures where to place operator written dynamic ConfigMaps (created with CEL expressions).</p>
</td>
</tr>
<tr>
<td>
<code>secretExpressions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#DestinationExpression">
[]genruntime/core.DestinationExpression
</a>
</em>
</td>
<td>
<p>SecretExpressions: configures where to place operator written dynamic secrets (created with CEL expressions).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersFirewallRule_STATUS">FlexibleServersFirewallRule_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersFirewallRule">FlexibleServersFirewallRule</a>)
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
<p>Id: Fully qualified resource ID for the resource. E.g.
&ldquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}&rdquo;</&#x200b;p>
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.SystemData_STATUS">
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersFirewallRule_Spec">FlexibleServersFirewallRule_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersFirewallRule">FlexibleServersFirewallRule</a>)
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
<code>operatorSpec</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersFirewallRuleOperatorSpec">
FlexibleServersFirewallRuleOperatorSpec
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
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.HighAvailability">HighAvailability
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_Spec">FlexibleServer_Spec</a>)
</p>
<div>
<p>High availability properties of a server</p>
</div>
<table>
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.HighAvailability_Mode">
HighAvailability_Mode
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.HighAvailability_Mode">HighAvailability_Mode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.HighAvailability">HighAvailability</a>)
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.HighAvailability_Mode_STATUS">HighAvailability_Mode_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.HighAvailability_STATUS">HighAvailability_STATUS</a>)
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.HighAvailability_STATUS">HighAvailability_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_STATUS">FlexibleServer_STATUS</a>)
</p>
<div>
<p>High availability properties of a server</p>
</div>
<table>
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.HighAvailability_Mode_STATUS">
HighAvailability_Mode_STATUS
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.HighAvailability_State_STATUS">
HighAvailability_State_STATUS
</a>
</em>
</td>
<td>
<p>State: A state of a HA server that is visible to user.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.HighAvailability_State_STATUS">HighAvailability_State_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.HighAvailability_STATUS">HighAvailability_STATUS</a>)
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.MaintenanceWindow">MaintenanceWindow
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_Spec">FlexibleServer_Spec</a>)
</p>
<div>
<p>Maintenance window properties of a server.</p>
</div>
<table>
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.MaintenanceWindow_STATUS">MaintenanceWindow_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_STATUS">FlexibleServer_STATUS</a>)
</p>
<div>
<p>Maintenance window properties of a server.</p>
</div>
<table>
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Network">Network
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_Spec">FlexibleServer_Spec</a>)
</p>
<div>
<p>Network properties of a server.</p>
</div>
<table>
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
<p>DelegatedSubnetResourceReference: Delegated subnet arm resource id. This is required to be passed during create, in case
we want the server to be VNET injected, i.e. Private access server. During update, pass this only if we want to update
the value for Private DNS zone.</p>
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
<p>PrivateDnsZoneArmResourceReference: Private dns zone arm resource id. This is required to be passed during create, in
case we want the server to be VNET injected, i.e. Private access server. During update, pass this only if we want to
update the value for Private DNS zone.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Network_PublicNetworkAccess">
Network_PublicNetworkAccess
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: public network access is enabled or not</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Network_PublicNetworkAccess">Network_PublicNetworkAccess
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.Network">Network</a>)
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Network_PublicNetworkAccess_STATUS">Network_PublicNetworkAccess_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.Network_STATUS">Network_STATUS</a>)
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Network_STATUS">Network_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_STATUS">FlexibleServer_STATUS</a>)
</p>
<div>
<p>Network properties of a server.</p>
</div>
<table>
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
<p>DelegatedSubnetResourceId: Delegated subnet arm resource id. This is required to be passed during create, in case we
want the server to be VNET injected, i.e. Private access server. During update, pass this only if we want to update the
value for Private DNS zone.</p>
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
<p>PrivateDnsZoneArmResourceId: Private dns zone arm resource id. This is required to be passed during create, in case we
want the server to be VNET injected, i.e. Private access server. During update, pass this only if we want to update the
value for Private DNS zone.</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccess</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Network_PublicNetworkAccess_STATUS">
Network_PublicNetworkAccess_STATUS
</a>
</em>
</td>
<td>
<p>PublicNetworkAccess: public network access is enabled or not</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.PrivateEndpointConnection_STATUS">PrivateEndpointConnection_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_STATUS">FlexibleServer_STATUS</a>)
</p>
<div>
<p>The private endpoint connection resource.</p>
</div>
<table>
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
<p>Id: Fully qualified resource ID for the resource. E.g.
&ldquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}&rdquo;</&#x200b;p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Replica">Replica
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_Spec">FlexibleServer_Spec</a>)
</p>
<div>
<p>Replica properties of a server</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>promoteMode</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Replica_PromoteMode">
Replica_PromoteMode
</a>
</em>
</td>
<td>
<p>PromoteMode: Sets the promote mode for a replica server. This is a write only property.</p>
</td>
</tr>
<tr>
<td>
<code>promoteOption</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Replica_PromoteOption">
Replica_PromoteOption
</a>
</em>
</td>
<td>
<p>PromoteOption: Sets the promote options for a replica server. This is a write only property.</p>
</td>
</tr>
<tr>
<td>
<code>role</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.ReplicationRole">
ReplicationRole
</a>
</em>
</td>
<td>
<p>Role: Used to indicate role of the server in replication set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Replica_PromoteMode">Replica_PromoteMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.Replica">Replica</a>)
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
<tbody><tr><td><p>&#34;standalone&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;switchover&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Replica_PromoteMode_STATUS">Replica_PromoteMode_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.Replica_STATUS">Replica_STATUS</a>)
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
<tbody><tr><td><p>&#34;standalone&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;switchover&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Replica_PromoteOption">Replica_PromoteOption
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.Replica">Replica</a>)
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
<tbody><tr><td><p>&#34;forced&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;planned&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Replica_PromoteOption_STATUS">Replica_PromoteOption_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.Replica_STATUS">Replica_STATUS</a>)
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
<tbody><tr><td><p>&#34;forced&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;planned&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Replica_ReplicationState_STATUS">Replica_ReplicationState_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.Replica_STATUS">Replica_STATUS</a>)
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
<tbody><tr><td><p>&#34;Active&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Broken&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Catchup&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Provisioning&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Reconfiguring&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Updating&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Replica_STATUS">Replica_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_STATUS">FlexibleServer_STATUS</a>)
</p>
<div>
<p>Replica properties of a server</p>
</div>
<table>
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
<p>Capacity: Replicas allowed for a server.</p>
</td>
</tr>
<tr>
<td>
<code>promoteMode</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Replica_PromoteMode_STATUS">
Replica_PromoteMode_STATUS
</a>
</em>
</td>
<td>
<p>PromoteMode: Sets the promote mode for a replica server. This is a write only property.</p>
</td>
</tr>
<tr>
<td>
<code>promoteOption</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Replica_PromoteOption_STATUS">
Replica_PromoteOption_STATUS
</a>
</em>
</td>
<td>
<p>PromoteOption: Sets the promote options for a replica server. This is a write only property.</p>
</td>
</tr>
<tr>
<td>
<code>replicationState</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Replica_ReplicationState_STATUS">
Replica_ReplicationState_STATUS
</a>
</em>
</td>
<td>
<p>ReplicationState: Gets the replication state of a replica server. This property is returned only for replicas api call.
Supported values are Active, Catchup, Provisioning, Updating, Broken, Reconfiguring</p>
</td>
</tr>
<tr>
<td>
<code>role</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.ReplicationRole_STATUS">
ReplicationRole_STATUS
</a>
</em>
</td>
<td>
<p>Role: Used to indicate role of the server in replication set.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.ReplicationRole">ReplicationRole
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_Spec">FlexibleServer_Spec</a>, <a href="#dbforpostgresql.azure.com/v1api20230601preview.Replica">Replica</a>)
</p>
<div>
<p>Used to indicate role of the server in replication set.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;AsyncReplica&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GeoAsyncReplica&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Primary&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.ReplicationRole_STATUS">ReplicationRole_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_STATUS">FlexibleServer_STATUS</a>, <a href="#dbforpostgresql.azure.com/v1api20230601preview.Replica_STATUS">Replica_STATUS</a>)
</p>
<div>
<p>Used to indicate role of the server in replication set.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;AsyncReplica&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GeoAsyncReplica&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Primary&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.ServerProperties_CreateMode">ServerProperties_CreateMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_Spec">FlexibleServer_Spec</a>)
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
</tr><tr><td><p>&#34;GeoRestore&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;PointInTimeRestore&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Replica&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ReviveDropped&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Update&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.ServerProperties_CreateMode_STATUS">ServerProperties_CreateMode_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_STATUS">FlexibleServer_STATUS</a>)
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
</tr><tr><td><p>&#34;GeoRestore&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;PointInTimeRestore&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Replica&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ReviveDropped&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Update&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.ServerProperties_State_STATUS">ServerProperties_State_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_STATUS">FlexibleServer_STATUS</a>)
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.ServerVersion">ServerVersion
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_Spec">FlexibleServer_Spec</a>)
</p>
<div>
<p>The version of a server.</p>
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
</tr><tr><td><p>&#34;14&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;15&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;16&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.ServerVersion_STATUS">ServerVersion_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_STATUS">FlexibleServer_STATUS</a>)
</p>
<div>
<p>The version of a server.</p>
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
</tr><tr><td><p>&#34;14&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;15&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;16&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Sku">Sku
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_Spec">FlexibleServer_Spec</a>)
</p>
<div>
<p>Sku information related properties of a server.</p>
</div>
<table>
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Sku_Tier">
Sku_Tier
</a>
</em>
</td>
<td>
<p>Tier: The tier of the particular SKU, e.g. Burstable.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Sku_STATUS">Sku_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_STATUS">FlexibleServer_STATUS</a>)
</p>
<div>
<p>Sku information related properties of a server.</p>
</div>
<table>
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Sku_Tier_STATUS">
Sku_Tier_STATUS
</a>
</em>
</td>
<td>
<p>Tier: The tier of the particular SKU, e.g. Burstable.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Sku_Tier">Sku_Tier
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.Sku">Sku</a>)
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Sku_Tier_STATUS">Sku_Tier_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.Sku_STATUS">Sku_STATUS</a>)
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Storage">Storage
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_Spec">FlexibleServer_Spec</a>)
</p>
<div>
<p>Storage properties of a server</p>
</div>
<table>
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Storage_AutoGrow">
Storage_AutoGrow
</a>
</em>
</td>
<td>
<p>AutoGrow: Flag to enable / disable Storage Auto grow for flexible server.</p>
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
<p>Iops: Storage tier IOPS quantity. This property is required to be set for storage Type PremiumV2_LRS</p>
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
<p>StorageSizeGB: Max storage allowed for a server.</p>
</td>
</tr>
<tr>
<td>
<code>throughput</code><br/>
<em>
int
</em>
</td>
<td>
<p>Throughput: Storage throughput for the server. This is required to be set for storage Type PremiumV2_LRS</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Storage_Tier">
Storage_Tier
</a>
</em>
</td>
<td>
<p>Tier: Name of storage tier for IOPS.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Storage_Type">
Storage_Type
</a>
</em>
</td>
<td>
<p>Type: Storage type for the server. Allowed values are Premium_LRS and PremiumV2_LRS, and default is Premium_LRS if not
specified</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Storage_AutoGrow">Storage_AutoGrow
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.Storage">Storage</a>)
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Storage_AutoGrow_STATUS">Storage_AutoGrow_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.Storage_STATUS">Storage_STATUS</a>)
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Storage_STATUS">Storage_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_STATUS">FlexibleServer_STATUS</a>)
</p>
<div>
<p>Storage properties of a server</p>
</div>
<table>
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Storage_AutoGrow_STATUS">
Storage_AutoGrow_STATUS
</a>
</em>
</td>
<td>
<p>AutoGrow: Flag to enable / disable Storage Auto grow for flexible server.</p>
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
<p>Iops: Storage tier IOPS quantity. This property is required to be set for storage Type PremiumV2_LRS</p>
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
<p>StorageSizeGB: Max storage allowed for a server.</p>
</td>
</tr>
<tr>
<td>
<code>throughput</code><br/>
<em>
int
</em>
</td>
<td>
<p>Throughput: Storage throughput for the server. This is required to be set for storage Type PremiumV2_LRS</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Storage_Tier_STATUS">
Storage_Tier_STATUS
</a>
</em>
</td>
<td>
<p>Tier: Name of storage tier for IOPS.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.Storage_Type_STATUS">
Storage_Type_STATUS
</a>
</em>
</td>
<td>
<p>Type: Storage type for the server. Allowed values are Premium_LRS and PremiumV2_LRS, and default is Premium_LRS if not
specified</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Storage_Tier">Storage_Tier
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.Storage">Storage</a>)
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
<tbody><tr><td><p>&#34;P1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P10&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P15&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P20&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P30&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P40&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P50&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P6&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P60&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P70&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P80&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Storage_Tier_STATUS">Storage_Tier_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.Storage_STATUS">Storage_STATUS</a>)
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
<tbody><tr><td><p>&#34;P1&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P10&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P15&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P2&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P20&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P3&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P30&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P4&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P40&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P50&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P6&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P60&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P70&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;P80&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Storage_Type">Storage_Type
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.Storage">Storage</a>)
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
<tbody><tr><td><p>&#34;PremiumV2_LRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Premium_LRS&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.Storage_Type_STATUS">Storage_Type_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.Storage_STATUS">Storage_STATUS</a>)
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
<tbody><tr><td><p>&#34;PremiumV2_LRS&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Premium_LRS&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.SystemData_CreatedByType_STATUS">SystemData_CreatedByType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.SystemData_STATUS">SystemData_STATUS</a>)
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.SystemData_LastModifiedByType_STATUS">SystemData_LastModifiedByType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.SystemData_STATUS">SystemData_STATUS</a>)
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.SystemData_STATUS">SystemData_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_STATUS">FlexibleServer_STATUS</a>, <a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersConfiguration_STATUS">FlexibleServersConfiguration_STATUS</a>, <a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersDatabase_STATUS">FlexibleServersDatabase_STATUS</a>, <a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServersFirewallRule_STATUS">FlexibleServersFirewallRule_STATUS</a>)
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.SystemData_CreatedByType_STATUS">
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.SystemData_LastModifiedByType_STATUS">
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.UserAssignedIdentity">UserAssignedIdentity
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_Spec">FlexibleServer_Spec</a>)
</p>
<div>
<p>Information describing the identities associated with this application.</p>
</div>
<table>
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
<a href="#dbforpostgresql.azure.com/v1api20230601preview.UserAssignedIdentity_Type">
UserAssignedIdentity_Type
</a>
</em>
</td>
<td>
<p>Type: the types of identities associated with this resource; currently restricted to &lsquo;None and UserAssigned&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.UserAssignedIdentityDetails">
[]UserAssignedIdentityDetails
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: represents user assigned identities map.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.UserAssignedIdentityDetails">UserAssignedIdentityDetails
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.UserAssignedIdentity">UserAssignedIdentity</a>)
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
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.UserAssignedIdentity_STATUS">UserAssignedIdentity_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.FlexibleServer_STATUS">FlexibleServer_STATUS</a>)
</p>
<div>
<p>Information describing the identities associated with this application.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>tenantId</code><br/>
<em>
string
</em>
</td>
<td>
<p>TenantId: Tenant id of the server.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.UserAssignedIdentity_Type_STATUS">
UserAssignedIdentity_Type_STATUS
</a>
</em>
</td>
<td>
<p>Type: the types of identities associated with this resource; currently restricted to &lsquo;None and UserAssigned&rsquo;</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#dbforpostgresql.azure.com/v1api20230601preview.UserIdentity_STATUS">
map[string]./api/dbforpostgresql/v1api20230601preview.UserIdentity_STATUS
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: represents user assigned identities map.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.UserAssignedIdentity_Type">UserAssignedIdentity_Type
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.UserAssignedIdentity">UserAssignedIdentity</a>)
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
</tr><tr><td><p>&#34;UserAssigned&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.UserAssignedIdentity_Type_STATUS">UserAssignedIdentity_Type_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.UserAssignedIdentity_STATUS">UserAssignedIdentity_STATUS</a>)
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
</tr><tr><td><p>&#34;UserAssigned&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="dbforpostgresql.azure.com/v1api20230601preview.UserIdentity_STATUS">UserIdentity_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#dbforpostgresql.azure.com/v1api20230601preview.UserAssignedIdentity_STATUS">UserAssignedIdentity_STATUS</a>)
</p>
<div>
<p>Describes a single user-assigned identity associated with the application.</p>
</div>
<table>
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
<p>ClientId: the client identifier of the Service Principal which this identity represents.</p>
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
<p>PrincipalId: the object identifier of the Service Principal which this identity represents.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
