---
title: dbformysql.azure.com/v1beta1
---
<h2 id="dbformysql.azure.com/v1beta1">dbformysql.azure.com/v1beta1</h2>
<div>
<p>Package v1beta contains hand crafted API Schema definitions for the dbformysql v1beta API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="dbformysql.azure.com/v1beta1.LocalUserSpec">LocalUserSpec
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1beta1.UserSpec">UserSpec</a>)
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
<code>serverAdminUsername</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServerAdminUsername is the user name of the Server administrator</p>
</td>
</tr>
<tr>
<td>
<code>serverAdminPassword</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretReference">
genruntime.SecretReference
</a>
</em>
</td>
<td>
<p>ServerAdminPassword is a reference to a secret containing the servers administrator password</p>
</td>
</tr>
<tr>
<td>
<code>password</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretReference">
genruntime.SecretReference
</a>
</em>
</td>
<td>
<p>Password is the password to use for the user</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1beta1.User">User
</h3>
<div>
<p>User is a MySQL user</p>
</div>
<table>
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
<a href="#dbformysql.azure.com/v1beta1.UserSpec">
UserSpec
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
reference to a dbformysql.azure.com/FlexibleServer resource</p>
</td>
</tr>
<tr>
<td>
<code>hostname</code><br/>
<em>
string
</em>
</td>
<td>
<p>Hostname is the host the user will connect from. If omitted, the default is to allow connection from any hostname.</p>
</td>
</tr>
<tr>
<td>
<code>privileges</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>The server-level roles assigned to the user.
Privileges include the following: RELOAD, PROCESS, SHOW
DATABASES, REPLICATION SLAVE, REPLICATION CLIENT, CREATE USER</p>
</td>
</tr>
<tr>
<td>
<code>databasePrivileges</code><br/>
<em>
map[string][]string
</em>
</td>
<td>
<p>The database-level roles assigned to the user (keyed by
database name). Privileges include the following: SELECT,
INSERT, UPDATE, DELETE, CREATE, DROP, REFERENCES, INDEX,
ALTER, CREATE TEMPORARY TABLES, LOCK TABLES, EXECUTE, CREATE
VIEW, SHOW VIEW, CREATE ROUTINE, ALTER ROUTINE, EVENT, TRIGGER</p>
</td>
</tr>
<tr>
<td>
<code>localUser</code><br/>
<em>
<a href="#dbformysql.azure.com/v1beta1.LocalUserSpec">
LocalUserSpec
</a>
</em>
</td>
<td>
<p>TODO: Note this is required right now but will move to be optional in the future when we have AAD support
LocalUser contains details for creating a standard (non-aad) MySQL User</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#dbformysql.azure.com/v1beta1.UserStatus">
UserStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1beta1.UserSpec">UserSpec
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1beta1.User">User</a>)
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
reference to a dbformysql.azure.com/FlexibleServer resource</p>
</td>
</tr>
<tr>
<td>
<code>hostname</code><br/>
<em>
string
</em>
</td>
<td>
<p>Hostname is the host the user will connect from. If omitted, the default is to allow connection from any hostname.</p>
</td>
</tr>
<tr>
<td>
<code>privileges</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>The server-level roles assigned to the user.
Privileges include the following: RELOAD, PROCESS, SHOW
DATABASES, REPLICATION SLAVE, REPLICATION CLIENT, CREATE USER</p>
</td>
</tr>
<tr>
<td>
<code>databasePrivileges</code><br/>
<em>
map[string][]string
</em>
</td>
<td>
<p>The database-level roles assigned to the user (keyed by
database name). Privileges include the following: SELECT,
INSERT, UPDATE, DELETE, CREATE, DROP, REFERENCES, INDEX,
ALTER, CREATE TEMPORARY TABLES, LOCK TABLES, EXECUTE, CREATE
VIEW, SHOW VIEW, CREATE ROUTINE, ALTER ROUTINE, EVENT, TRIGGER</p>
</td>
</tr>
<tr>
<td>
<code>localUser</code><br/>
<em>
<a href="#dbformysql.azure.com/v1beta1.LocalUserSpec">
LocalUserSpec
</a>
</em>
</td>
<td>
<p>TODO: Note this is required right now but will move to be optional in the future when we have AAD support
LocalUser contains details for creating a standard (non-aad) MySQL User</p>
</td>
</tr>
</tbody>
</table>
<h3 id="dbformysql.azure.com/v1beta1.UserStatus">UserStatus
</h3>
<p>
(<em>Appears on:</em><a href="#dbformysql.azure.com/v1beta1.User">User</a>)
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
</tbody>
</table>
<hr/>
