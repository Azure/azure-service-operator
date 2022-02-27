---
---
<h2 id="eventhub.azure.com/v1alpha1api20211101">eventhub.azure.com/v1alpha1api20211101</h2>
<div>
<p>Package v1alpha1api20211101 contains API Schema definitions for the eventhub v1alpha1api20211101 API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="eventhub.azure.com/v1alpha1api20211101.AuthorizationRulePropertiesARM">AuthorizationRulePropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesAuthorizationRules_SpecARM">NamespacesAuthorizationRules_SpecARM</a>, <a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubsAuthorizationRules_SpecARM">NamespacesEventhubsAuthorizationRules_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/AuthorizationRuleProperties">https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/AuthorizationRuleProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>rights</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.AuthorizationRulePropertiesRights">
[]AuthorizationRulePropertiesRights
</a>
</em>
</td>
<td>
<p>Rights: The rights associated with the rule.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.AuthorizationRulePropertiesRights">AuthorizationRulePropertiesRights
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.AuthorizationRulePropertiesARM">AuthorizationRulePropertiesARM</a>, <a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesAuthorizationRules_Spec">NamespacesAuthorizationRules_Spec</a>, <a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubsAuthorizationRules_Spec">NamespacesEventhubsAuthorizationRules_Spec</a>)
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
<tbody><tr><td><p>&#34;Listen&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Manage&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Send&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.AuthorizationRuleStatusPropertiesRights">AuthorizationRuleStatusPropertiesRights
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.AuthorizationRule_Status">AuthorizationRule_Status</a>, <a href="#eventhub.azure.com/v1alpha1api20211101.AuthorizationRule_Status_PropertiesARM">AuthorizationRule_Status_PropertiesARM</a>)
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
<tbody><tr><td><p>&#34;Listen&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Manage&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Send&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.AuthorizationRule_Status">AuthorizationRule_Status
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesAuthorizationRule">NamespacesAuthorizationRule</a>, <a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubsAuthorizationRule">NamespacesEventhubsAuthorizationRule</a>)
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
[]github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions.Condition
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
<code>rights</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.AuthorizationRuleStatusPropertiesRights">
[]AuthorizationRuleStatusPropertiesRights
</a>
</em>
</td>
<td>
<p>Rights: The rights associated with the rule.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.SystemData_Status">
SystemData_Status
</a>
</em>
</td>
<td>
<p>SystemData: The system meta data relating to this resource.</p>
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
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.EventHub/Namespaces&rdquo; or &ldquo;Microsoft.EventHub/Namespaces/EventHubs&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.AuthorizationRule_StatusARM">AuthorizationRule_StatusARM
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
<a href="#eventhub.azure.com/v1alpha1api20211101.AuthorizationRule_Status_PropertiesARM">
AuthorizationRule_Status_PropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties supplied to create or update AuthorizationRule</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.SystemData_StatusARM">
SystemData_StatusARM
</a>
</em>
</td>
<td>
<p>SystemData: The system meta data relating to this resource.</p>
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
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.EventHub/Namespaces&rdquo; or &ldquo;Microsoft.EventHub/Namespaces/EventHubs&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.AuthorizationRule_Status_PropertiesARM">AuthorizationRule_Status_PropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.AuthorizationRule_StatusARM">AuthorizationRule_StatusARM</a>)
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
<code>rights</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.AuthorizationRuleStatusPropertiesRights">
[]AuthorizationRuleStatusPropertiesRights
</a>
</em>
</td>
<td>
<p>Rights: The rights associated with the rule.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.CaptureDescriptionStatusEncoding">CaptureDescriptionStatusEncoding
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.CaptureDescription_Status">CaptureDescription_Status</a>, <a href="#eventhub.azure.com/v1alpha1api20211101.CaptureDescription_StatusARM">CaptureDescription_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Avro&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;AvroDeflate&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.CaptureDescription_Status">CaptureDescription_Status
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Eventhub_Status">Eventhub_Status</a>)
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
<code>destination</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.Destination_Status">
Destination_Status
</a>
</em>
</td>
<td>
<p>Destination: Properties of Destination where capture will be stored. (Storage Account, Blob Names)</p>
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
<p>Enabled: A value that indicates whether capture description is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>encoding</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.CaptureDescriptionStatusEncoding">
CaptureDescriptionStatusEncoding
</a>
</em>
</td>
<td>
<p>Encoding: Enumerates the possible values for the encoding format of capture description. Note: &lsquo;AvroDeflate&rsquo; will be
deprecated in New API Version</p>
</td>
</tr>
<tr>
<td>
<code>intervalInSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>IntervalInSeconds: The time window allows you to set the frequency with which the capture to Azure Blobs will happen,
value should between 60 to 900 seconds</p>
</td>
</tr>
<tr>
<td>
<code>sizeLimitInBytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>SizeLimitInBytes: The size window defines the amount of data built up in your Event Hub before an capture operation,
value should be between 10485760 to 524288000 bytes</p>
</td>
</tr>
<tr>
<td>
<code>skipEmptyArchives</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SkipEmptyArchives: A value that indicates whether to Skip Empty Archives</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.CaptureDescription_StatusARM">CaptureDescription_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Eventhub_Status_PropertiesARM">Eventhub_Status_PropertiesARM</a>)
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
<code>destination</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.Destination_StatusARM">
Destination_StatusARM
</a>
</em>
</td>
<td>
<p>Destination: Properties of Destination where capture will be stored. (Storage Account, Blob Names)</p>
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
<p>Enabled: A value that indicates whether capture description is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>encoding</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.CaptureDescriptionStatusEncoding">
CaptureDescriptionStatusEncoding
</a>
</em>
</td>
<td>
<p>Encoding: Enumerates the possible values for the encoding format of capture description. Note: &lsquo;AvroDeflate&rsquo; will be
deprecated in New API Version</p>
</td>
</tr>
<tr>
<td>
<code>intervalInSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>IntervalInSeconds: The time window allows you to set the frequency with which the capture to Azure Blobs will happen,
value should between 60 to 900 seconds</p>
</td>
</tr>
<tr>
<td>
<code>sizeLimitInBytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>SizeLimitInBytes: The size window defines the amount of data built up in your Event Hub before an capture operation,
value should be between 10485760 to 524288000 bytes</p>
</td>
</tr>
<tr>
<td>
<code>skipEmptyArchives</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SkipEmptyArchives: A value that indicates whether to Skip Empty Archives</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.ConsumerGroupPropertiesARM">ConsumerGroupPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubsConsumergroups_SpecARM">NamespacesEventhubsConsumergroups_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/ConsumerGroupProperties">https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/ConsumerGroupProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>userMetadata</code><br/>
<em>
string
</em>
</td>
<td>
<p>UserMetadata: User Metadata is a placeholder to store user-defined string data with maximum length 1024. e.g. it can be
used to store descriptive data, such as list of teams and their contact information also user-defined configuration
settings can be stored.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.ConsumerGroup_Status">ConsumerGroup_Status
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubsConsumerGroup">NamespacesEventhubsConsumerGroup</a>)
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
[]github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions.Condition
</em>
</td>
<td>
<p>Conditions: The observed state of the resource</p>
</td>
</tr>
<tr>
<td>
<code>createdAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedAt: Exact time the message was created.</p>
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
<a href="#eventhub.azure.com/v1alpha1api20211101.SystemData_Status">
SystemData_Status
</a>
</em>
</td>
<td>
<p>SystemData: The system meta data relating to this resource.</p>
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
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.EventHub/Namespaces&rdquo; or &ldquo;Microsoft.EventHub/Namespaces/EventHubs&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>updatedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>UpdatedAt: The exact time the message was updated.</p>
</td>
</tr>
<tr>
<td>
<code>userMetadata</code><br/>
<em>
string
</em>
</td>
<td>
<p>UserMetadata: User Metadata is a placeholder to store user-defined string data with maximum length 1024. e.g. it can be
used to store descriptive data, such as list of teams and their contact information also user-defined configuration
settings can be stored.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.ConsumerGroup_StatusARM">ConsumerGroup_StatusARM
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
<a href="#eventhub.azure.com/v1alpha1api20211101.ConsumerGroup_Status_PropertiesARM">
ConsumerGroup_Status_PropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Single item in List or Get Consumer group operation</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.SystemData_StatusARM">
SystemData_StatusARM
</a>
</em>
</td>
<td>
<p>SystemData: The system meta data relating to this resource.</p>
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
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.EventHub/Namespaces&rdquo; or &ldquo;Microsoft.EventHub/Namespaces/EventHubs&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.ConsumerGroup_Status_PropertiesARM">ConsumerGroup_Status_PropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.ConsumerGroup_StatusARM">ConsumerGroup_StatusARM</a>)
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
<p>CreatedAt: Exact time the message was created.</p>
</td>
</tr>
<tr>
<td>
<code>updatedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>UpdatedAt: The exact time the message was updated.</p>
</td>
</tr>
<tr>
<td>
<code>userMetadata</code><br/>
<em>
string
</em>
</td>
<td>
<p>UserMetadata: User Metadata is a placeholder to store user-defined string data with maximum length 1024. e.g. it can be
used to store descriptive data, such as list of teams and their contact information also user-defined configuration
settings can be stored.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.DestinationPropertiesARM">DestinationPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubs_Spec_Properties_CaptureDescription_DestinationARM">NamespacesEventhubs_Spec_Properties_CaptureDescription_DestinationARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/DestinationProperties">https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/DestinationProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>archiveNameFormat</code><br/>
<em>
string
</em>
</td>
<td>
<p>ArchiveNameFormat: Blob naming convention for archive, e.g.
{Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}. Here all the parameters
(Namespace,EventHub .. etc) are mandatory irrespective of order</p>
</td>
</tr>
<tr>
<td>
<code>blobContainer</code><br/>
<em>
string
</em>
</td>
<td>
<p>BlobContainer: Blob container Name</p>
</td>
</tr>
<tr>
<td>
<code>dataLakeAccountName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DataLakeAccountName: The Azure Data Lake Store name for the captured events</p>
</td>
</tr>
<tr>
<td>
<code>dataLakeFolderPath</code><br/>
<em>
string
</em>
</td>
<td>
<p>DataLakeFolderPath: The destination folder path for the captured events</p>
</td>
</tr>
<tr>
<td>
<code>dataLakeSubscriptionId</code><br/>
<em>
string
</em>
</td>
<td>
<p>DataLakeSubscriptionId: Subscription Id of Azure Data Lake Store</p>
</td>
</tr>
<tr>
<td>
<code>storageAccountResourceId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.Destination_Status">Destination_Status
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.CaptureDescription_Status">CaptureDescription_Status</a>)
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
<code>archiveNameFormat</code><br/>
<em>
string
</em>
</td>
<td>
<p>ArchiveNameFormat: Blob naming convention for archive, e.g.
{Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}. Here all the parameters
(Namespace,EventHub .. etc) are mandatory irrespective of order</p>
</td>
</tr>
<tr>
<td>
<code>blobContainer</code><br/>
<em>
string
</em>
</td>
<td>
<p>BlobContainer: Blob container Name</p>
</td>
</tr>
<tr>
<td>
<code>dataLakeAccountName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DataLakeAccountName: The Azure Data Lake Store name for the captured events</p>
</td>
</tr>
<tr>
<td>
<code>dataLakeFolderPath</code><br/>
<em>
string
</em>
</td>
<td>
<p>DataLakeFolderPath: The destination folder path for the captured events</p>
</td>
</tr>
<tr>
<td>
<code>dataLakeSubscriptionId</code><br/>
<em>
string
</em>
</td>
<td>
<p>DataLakeSubscriptionId: Subscription Id of Azure Data Lake Store</p>
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
<p>Name: Name for capture destination</p>
</td>
</tr>
<tr>
<td>
<code>storageAccountResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>StorageAccountResourceId: Resource id of the storage account to be used to create the blobs</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.Destination_StatusARM">Destination_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.CaptureDescription_StatusARM">CaptureDescription_StatusARM</a>)
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
<p>Name: Name for capture destination</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.Destination_Status_PropertiesARM">
Destination_Status_PropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties describing the storage account, blob container and archive name format for capture destination</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.Destination_Status_PropertiesARM">Destination_Status_PropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Destination_StatusARM">Destination_StatusARM</a>)
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
<code>archiveNameFormat</code><br/>
<em>
string
</em>
</td>
<td>
<p>ArchiveNameFormat: Blob naming convention for archive, e.g.
{Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}. Here all the parameters
(Namespace,EventHub .. etc) are mandatory irrespective of order</p>
</td>
</tr>
<tr>
<td>
<code>blobContainer</code><br/>
<em>
string
</em>
</td>
<td>
<p>BlobContainer: Blob container Name</p>
</td>
</tr>
<tr>
<td>
<code>dataLakeAccountName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DataLakeAccountName: The Azure Data Lake Store name for the captured events</p>
</td>
</tr>
<tr>
<td>
<code>dataLakeFolderPath</code><br/>
<em>
string
</em>
</td>
<td>
<p>DataLakeFolderPath: The destination folder path for the captured events</p>
</td>
</tr>
<tr>
<td>
<code>dataLakeSubscriptionId</code><br/>
<em>
string
</em>
</td>
<td>
<p>DataLakeSubscriptionId: Subscription Id of Azure Data Lake Store</p>
</td>
</tr>
<tr>
<td>
<code>storageAccountResourceId</code><br/>
<em>
string
</em>
</td>
<td>
<p>StorageAccountResourceId: Resource id of the storage account to be used to create the blobs</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.EHNamespace_Status">EHNamespace_Status
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Namespace">Namespace</a>)
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
<code>alternateName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AlternateName: Alternate name specified when alias and namespace names are same.</p>
</td>
</tr>
<tr>
<td>
<code>clusterArmId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClusterArmId: Cluster ARM ID of the Namespace.</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code><br/>
<em>
[]github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions.Condition
</em>
</td>
<td>
<p>Conditions: The observed state of the resource</p>
</td>
</tr>
<tr>
<td>
<code>createdAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedAt: The time the Namespace was created.</p>
</td>
</tr>
<tr>
<td>
<code>disableLocalAuth</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableLocalAuth: This property disables SAS authentication for the Event Hubs namespace.</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.Encryption_Status">
Encryption_Status
</a>
</em>
</td>
<td>
<p>Encryption: Properties of BYOK Encryption description</p>
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
<code>identity</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.Identity_Status">
Identity_Status
</a>
</em>
</td>
<td>
<p>Identity: Properties of BYOK Identity description</p>
</td>
</tr>
<tr>
<td>
<code>isAutoInflateEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsAutoInflateEnabled: Value that indicates whether AutoInflate is enabled for eventhub namespace.</p>
</td>
</tr>
<tr>
<td>
<code>kafkaEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>KafkaEnabled: Value that indicates whether Kafka is enabled for eventhub namespace.</p>
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
<p>Location: Resource location.</p>
</td>
</tr>
<tr>
<td>
<code>maximumThroughputUnits</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaximumThroughputUnits: Upper limit of throughput units when AutoInflate is enabled, value should be within 0 to 20
throughput units. ( &lsquo;0&rsquo; if AutoInflateEnabled = true)</p>
</td>
</tr>
<tr>
<td>
<code>metricId</code><br/>
<em>
string
</em>
</td>
<td>
<p>MetricId: Identifier for Azure Insights metrics.</p>
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
<code>privateEndpointConnections</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.PrivateEndpointConnection_Status_SubResourceEmbedded">
[]PrivateEndpointConnection_Status_SubResourceEmbedded
</a>
</em>
</td>
<td>
<p>PrivateEndpointConnections: List of private endpoint connections.</p>
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
<p>ProvisioningState: Provisioning state of the Namespace.</p>
</td>
</tr>
<tr>
<td>
<code>serviceBusEndpoint</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServiceBusEndpoint: Endpoint you can use to perform Service Bus operations.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.Sku_Status">
Sku_Status
</a>
</em>
</td>
<td>
<p>Sku: Properties of sku resource</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
string
</em>
</td>
<td>
<p>Status: Status of the Namespace.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.SystemData_Status">
SystemData_Status
</a>
</em>
</td>
<td>
<p>SystemData: The system meta data relating to this resource.</p>
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
<code>updatedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>UpdatedAt: The time the Namespace was updated.</p>
</td>
</tr>
<tr>
<td>
<code>zoneRedundant</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ZoneRedundant: Enabling this property creates a Standard Event Hubs Namespace in regions supported availability zones.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.EHNamespace_StatusARM">EHNamespace_StatusARM
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
<code>identity</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.Identity_StatusARM">
Identity_StatusARM
</a>
</em>
</td>
<td>
<p>Identity: Properties of BYOK Identity description</p>
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
<p>Location: Resource location.</p>
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
<a href="#eventhub.azure.com/v1alpha1api20211101.EHNamespace_Status_PropertiesARM">
EHNamespace_Status_PropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Namespace properties supplied for create namespace operation.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.Sku_StatusARM">
Sku_StatusARM
</a>
</em>
</td>
<td>
<p>Sku: Properties of sku resource</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.SystemData_StatusARM">
SystemData_StatusARM
</a>
</em>
</td>
<td>
<p>SystemData: The system meta data relating to this resource.</p>
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
<h3 id="eventhub.azure.com/v1alpha1api20211101.EHNamespace_Status_PropertiesARM">EHNamespace_Status_PropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.EHNamespace_StatusARM">EHNamespace_StatusARM</a>)
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
<code>alternateName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AlternateName: Alternate name specified when alias and namespace names are same.</p>
</td>
</tr>
<tr>
<td>
<code>clusterArmId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClusterArmId: Cluster ARM ID of the Namespace.</p>
</td>
</tr>
<tr>
<td>
<code>createdAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedAt: The time the Namespace was created.</p>
</td>
</tr>
<tr>
<td>
<code>disableLocalAuth</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableLocalAuth: This property disables SAS authentication for the Event Hubs namespace.</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.Encryption_StatusARM">
Encryption_StatusARM
</a>
</em>
</td>
<td>
<p>Encryption: Properties of BYOK Encryption description</p>
</td>
</tr>
<tr>
<td>
<code>isAutoInflateEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsAutoInflateEnabled: Value that indicates whether AutoInflate is enabled for eventhub namespace.</p>
</td>
</tr>
<tr>
<td>
<code>kafkaEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>KafkaEnabled: Value that indicates whether Kafka is enabled for eventhub namespace.</p>
</td>
</tr>
<tr>
<td>
<code>maximumThroughputUnits</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaximumThroughputUnits: Upper limit of throughput units when AutoInflate is enabled, value should be within 0 to 20
throughput units. ( &lsquo;0&rsquo; if AutoInflateEnabled = true)</p>
</td>
</tr>
<tr>
<td>
<code>metricId</code><br/>
<em>
string
</em>
</td>
<td>
<p>MetricId: Identifier for Azure Insights metrics.</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpointConnections</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.PrivateEndpointConnection_Status_SubResourceEmbeddedARM">
[]PrivateEndpointConnection_Status_SubResourceEmbeddedARM
</a>
</em>
</td>
<td>
<p>PrivateEndpointConnections: List of private endpoint connections.</p>
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
<p>ProvisioningState: Provisioning state of the Namespace.</p>
</td>
</tr>
<tr>
<td>
<code>serviceBusEndpoint</code><br/>
<em>
string
</em>
</td>
<td>
<p>ServiceBusEndpoint: Endpoint you can use to perform Service Bus operations.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
string
</em>
</td>
<td>
<p>Status: Status of the Namespace.</p>
</td>
</tr>
<tr>
<td>
<code>updatedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>UpdatedAt: The time the Namespace was updated.</p>
</td>
</tr>
<tr>
<td>
<code>zoneRedundant</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ZoneRedundant: Enabling this property creates a Standard Event Hubs Namespace in regions supported availability zones.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.Encryption">Encryption
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Namespaces_Spec">Namespaces_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/Encryption">https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/Encryption</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>keySource</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.EncryptionKeySource">
EncryptionKeySource
</a>
</em>
</td>
<td>
<p>KeySource: Enumerates the possible value of keySource for Encryption.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultProperties</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.KeyVaultProperties">
[]KeyVaultProperties
</a>
</em>
</td>
<td>
<p>KeyVaultProperties: Properties of KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>requireInfrastructureEncryption</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequireInfrastructureEncryption: Enable Infrastructure Encryption (Double Encryption)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.EncryptionARM">EncryptionARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Namespaces_Spec_PropertiesARM">Namespaces_Spec_PropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/Encryption">https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/Encryption</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>keySource</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.EncryptionKeySource">
EncryptionKeySource
</a>
</em>
</td>
<td>
<p>KeySource: Enumerates the possible value of keySource for Encryption.</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultProperties</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.KeyVaultPropertiesARM">
[]KeyVaultPropertiesARM
</a>
</em>
</td>
<td>
<p>KeyVaultProperties: Properties of KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>requireInfrastructureEncryption</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequireInfrastructureEncryption: Enable Infrastructure Encryption (Double Encryption)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.EncryptionKeySource">EncryptionKeySource
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Encryption">Encryption</a>, <a href="#eventhub.azure.com/v1alpha1api20211101.EncryptionARM">EncryptionARM</a>)
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
<tbody><tr><td><p>&#34;Microsoft.KeyVault&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.EncryptionStatusKeySource">EncryptionStatusKeySource
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Encryption_Status">Encryption_Status</a>, <a href="#eventhub.azure.com/v1alpha1api20211101.Encryption_StatusARM">Encryption_StatusARM</a>)
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
<tbody><tr><td><p>&#34;Microsoft.KeyVault&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.Encryption_Status">Encryption_Status
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.EHNamespace_Status">EHNamespace_Status</a>)
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
<code>keySource</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.EncryptionStatusKeySource">
EncryptionStatusKeySource
</a>
</em>
</td>
<td>
<p>KeySource: Enumerates the possible value of keySource for Encryption</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultProperties</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.KeyVaultProperties_Status">
[]KeyVaultProperties_Status
</a>
</em>
</td>
<td>
<p>KeyVaultProperties: Properties of KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>requireInfrastructureEncryption</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequireInfrastructureEncryption: Enable Infrastructure Encryption (Double Encryption)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.Encryption_StatusARM">Encryption_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.EHNamespace_Status_PropertiesARM">EHNamespace_Status_PropertiesARM</a>)
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
<code>keySource</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.EncryptionStatusKeySource">
EncryptionStatusKeySource
</a>
</em>
</td>
<td>
<p>KeySource: Enumerates the possible value of keySource for Encryption</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultProperties</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.KeyVaultProperties_StatusARM">
[]KeyVaultProperties_StatusARM
</a>
</em>
</td>
<td>
<p>KeyVaultProperties: Properties of KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>requireInfrastructureEncryption</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequireInfrastructureEncryption: Enable Infrastructure Encryption (Double Encryption)</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.EventhubStatusPropertiesStatus">EventhubStatusPropertiesStatus
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Eventhub_Status">Eventhub_Status</a>, <a href="#eventhub.azure.com/v1alpha1api20211101.Eventhub_Status_PropertiesARM">Eventhub_Status_PropertiesARM</a>)
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
</tr><tr><td><p>&#34;Creating&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Deleting&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ReceiveDisabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Renaming&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Restoring&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SendDisabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Unknown&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.Eventhub_Status">Eventhub_Status
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhub">NamespacesEventhub</a>)
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
<code>captureDescription</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.CaptureDescription_Status">
CaptureDescription_Status
</a>
</em>
</td>
<td>
<p>CaptureDescription: Properties of capture description</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code><br/>
<em>
[]github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions.Condition
</em>
</td>
<td>
<p>Conditions: The observed state of the resource</p>
</td>
</tr>
<tr>
<td>
<code>createdAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedAt: Exact time the Event Hub was created.</p>
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
<code>messageRetentionInDays</code><br/>
<em>
int
</em>
</td>
<td>
<p>MessageRetentionInDays: Number of days to retain the events for this Event Hub, value should be 1 to 7 days</p>
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
<code>partitionCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>PartitionCount: Number of partitions created for the Event Hub, allowed values are from 1 to 32 partitions.</p>
</td>
</tr>
<tr>
<td>
<code>partitionIds</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>PartitionIds: Current number of shards on the Event Hub.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.EventhubStatusPropertiesStatus">
EventhubStatusPropertiesStatus
</a>
</em>
</td>
<td>
<p>Status: Enumerates the possible values for the status of the Event Hub.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.SystemData_Status">
SystemData_Status
</a>
</em>
</td>
<td>
<p>SystemData: The system meta data relating to this resource.</p>
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
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.EventHub/Namespaces&rdquo; or &ldquo;Microsoft.EventHub/Namespaces/EventHubs&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>updatedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>UpdatedAt: The exact time the message was updated.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.Eventhub_StatusARM">Eventhub_StatusARM
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
<a href="#eventhub.azure.com/v1alpha1api20211101.Eventhub_Status_PropertiesARM">
Eventhub_Status_PropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties supplied to the Create Or Update Event Hub operation.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.SystemData_StatusARM">
SystemData_StatusARM
</a>
</em>
</td>
<td>
<p>SystemData: The system meta data relating to this resource.</p>
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
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.EventHub/Namespaces&rdquo; or &ldquo;Microsoft.EventHub/Namespaces/EventHubs&rdquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.Eventhub_Status_PropertiesARM">Eventhub_Status_PropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Eventhub_StatusARM">Eventhub_StatusARM</a>)
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
<code>captureDescription</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.CaptureDescription_StatusARM">
CaptureDescription_StatusARM
</a>
</em>
</td>
<td>
<p>CaptureDescription: Properties of capture description</p>
</td>
</tr>
<tr>
<td>
<code>createdAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedAt: Exact time the Event Hub was created.</p>
</td>
</tr>
<tr>
<td>
<code>messageRetentionInDays</code><br/>
<em>
int
</em>
</td>
<td>
<p>MessageRetentionInDays: Number of days to retain the events for this Event Hub, value should be 1 to 7 days</p>
</td>
</tr>
<tr>
<td>
<code>partitionCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>PartitionCount: Number of partitions created for the Event Hub, allowed values are from 1 to 32 partitions.</p>
</td>
</tr>
<tr>
<td>
<code>partitionIds</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>PartitionIds: Current number of shards on the Event Hub.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.EventhubStatusPropertiesStatus">
EventhubStatusPropertiesStatus
</a>
</em>
</td>
<td>
<p>Status: Enumerates the possible values for the status of the Event Hub.</p>
</td>
</tr>
<tr>
<td>
<code>updatedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>UpdatedAt: The exact time the message was updated.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.Identity">Identity
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Namespaces_Spec">Namespaces_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/Identity">https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/Identity</a></p>
</div>
<table>
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
<a href="#eventhub.azure.com/v1alpha1api20211101.IdentityType">
IdentityType
</a>
</em>
</td>
<td>
<p>Type: Type of managed service identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.IdentityARM">IdentityARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Namespaces_SpecARM">Namespaces_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/Identity">https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/Identity</a></p>
</div>
<table>
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
<a href="#eventhub.azure.com/v1alpha1api20211101.IdentityType">
IdentityType
</a>
</em>
</td>
<td>
<p>Type: Type of managed service identity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.IdentityStatusType">IdentityStatusType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Identity_Status">Identity_Status</a>, <a href="#eventhub.azure.com/v1alpha1api20211101.Identity_StatusARM">Identity_StatusARM</a>)
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
</tr><tr><td><p>&#34;SystemAssigned, UserAssigned&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UserAssigned&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.IdentityType">IdentityType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Identity">Identity</a>, <a href="#eventhub.azure.com/v1alpha1api20211101.IdentityARM">IdentityARM</a>)
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
</tr><tr><td><p>&#34;SystemAssigned, UserAssigned&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UserAssigned&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.Identity_Status">Identity_Status
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.EHNamespace_Status">EHNamespace_Status</a>)
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
<p>PrincipalId: ObjectId from the KeyVault</p>
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
<p>TenantId: TenantId from the KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.IdentityStatusType">
IdentityStatusType
</a>
</em>
</td>
<td>
<p>Type: Type of managed service identity.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.UserAssignedIdentity_Status">
map[string]./api/eventhub/v1alpha1api20211101.UserAssignedIdentity_Status
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: Properties for User Assigned Identities</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.Identity_StatusARM">Identity_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.EHNamespace_StatusARM">EHNamespace_StatusARM</a>)
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
<p>PrincipalId: ObjectId from the KeyVault</p>
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
<p>TenantId: TenantId from the KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.IdentityStatusType">
IdentityStatusType
</a>
</em>
</td>
<td>
<p>Type: Type of managed service identity.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.UserAssignedIdentity_StatusARM">
map[string]./api/eventhub/v1alpha1api20211101.UserAssignedIdentity_StatusARM
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: Properties for User Assigned Identities</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.KeyVaultProperties">KeyVaultProperties
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Encryption">Encryption</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/KeyVaultProperties">https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/KeyVaultProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.UserAssignedIdentityProperties">
UserAssignedIdentityProperties
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>keyName</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyName: Name of the Key from KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVaultUri: Uri of KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>keyVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVersion: Key Version</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.KeyVaultPropertiesARM">KeyVaultPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.EncryptionARM">EncryptionARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/KeyVaultProperties">https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/KeyVaultProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.UserAssignedIdentityPropertiesARM">
UserAssignedIdentityPropertiesARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>keyName</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyName: Name of the Key from KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVaultUri: Uri of KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>keyVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVersion: Key Version</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.KeyVaultProperties_Status">KeyVaultProperties_Status
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Encryption_Status">Encryption_Status</a>)
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
<code>identity</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.UserAssignedIdentityProperties_Status">
UserAssignedIdentityProperties_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>keyName</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyName: Name of the Key from KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVaultUri: Uri of KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>keyVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVersion: Key Version</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.KeyVaultProperties_StatusARM">KeyVaultProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Encryption_StatusARM">Encryption_StatusARM</a>)
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
<code>identity</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.UserAssignedIdentityProperties_StatusARM">
UserAssignedIdentityProperties_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>keyName</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyName: Name of the Key from KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>keyVaultUri</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVaultUri: Uri of KeyVault</p>
</td>
</tr>
<tr>
<td>
<code>keyVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>KeyVersion: Key Version</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.Namespace">Namespace
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/resourceDefinitions/namespaces">https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/resourceDefinitions/namespaces</a></p>
</div>
<table>
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
<a href="#eventhub.azure.com/v1alpha1api20211101.Namespaces_Spec">
Namespaces_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>alternateName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AlternateName: Alternate name specified when alias and namespace names are same.</p>
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
<code>clusterArmReference</code><br/>
<em>
github.com/Azure/azure-service-operator/v2/pkg/genruntime.ResourceReference
</em>
</td>
<td>
<p>ClusterArmReference: Cluster ARM ID of the Namespace.</p>
</td>
</tr>
<tr>
<td>
<code>disableLocalAuth</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableLocalAuth: This property disables SAS authentication for the Event Hubs namespace.</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.Encryption">
Encryption
</a>
</em>
</td>
<td>
<p>Encryption: Properties to configure Encryption</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.Identity">
Identity
</a>
</em>
</td>
<td>
<p>Identity: Properties to configure Identity for Bring your Own Keys</p>
</td>
</tr>
<tr>
<td>
<code>isAutoInflateEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsAutoInflateEnabled: Value that indicates whether AutoInflate is enabled for eventhub namespace.</p>
</td>
</tr>
<tr>
<td>
<code>kafkaEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>KafkaEnabled: Value that indicates whether Kafka is enabled for eventhub namespace.</p>
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
<code>maximumThroughputUnits</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaximumThroughputUnits: Upper limit of throughput units when AutoInflate is enabled, value should be within 0 to 20
throughput units. ( &lsquo;0&rsquo; if AutoInflateEnabled = true)</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
github.com/Azure/azure-service-operator/v2/pkg/genruntime.KnownResourceReference
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>privateEndpointConnections</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.Namespaces_Spec_Properties_PrivateEndpointConnections">
[]Namespaces_Spec_Properties_PrivateEndpointConnections
</a>
</em>
</td>
<td>
<p>PrivateEndpointConnections: List of private endpoint connections.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.Sku">
Sku
</a>
</em>
</td>
<td>
<p>Sku: SKU parameters supplied to the create namespace operation</p>
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
<code>zoneRedundant</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ZoneRedundant: Enabling this property creates a Standard Event Hubs Namespace in regions supported availability zones.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.EHNamespace_Status">
EHNamespace_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.NamespacesAuthorizationRule">NamespacesAuthorizationRule
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/resourceDefinitions/namespaces_authorizationRules">https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/resourceDefinitions/namespaces_authorizationRules</a></p>
</div>
<table>
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
<a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesAuthorizationRules_Spec">
NamespacesAuthorizationRules_Spec
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
github.com/Azure/azure-service-operator/v2/pkg/genruntime.KnownResourceReference
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>rights</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.AuthorizationRulePropertiesRights">
[]AuthorizationRulePropertiesRights
</a>
</em>
</td>
<td>
<p>Rights: The rights associated with the rule.</p>
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
<a href="#eventhub.azure.com/v1alpha1api20211101.AuthorizationRule_Status">
AuthorizationRule_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.NamespacesAuthorizationRulesSpecAPIVersion">NamespacesAuthorizationRulesSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-11-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.NamespacesAuthorizationRules_Spec">NamespacesAuthorizationRules_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesAuthorizationRule">NamespacesAuthorizationRule</a>)
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
github.com/Azure/azure-service-operator/v2/pkg/genruntime.KnownResourceReference
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>rights</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.AuthorizationRulePropertiesRights">
[]AuthorizationRulePropertiesRights
</a>
</em>
</td>
<td>
<p>Rights: The rights associated with the rule.</p>
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
<h3 id="eventhub.azure.com/v1alpha1api20211101.NamespacesAuthorizationRules_SpecARM">NamespacesAuthorizationRules_SpecARM
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
<p>Name: The authorization rule name.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.AuthorizationRulePropertiesARM">
AuthorizationRulePropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties supplied to create or update AuthorizationRule</p>
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
<h3 id="eventhub.azure.com/v1alpha1api20211101.NamespacesEventhub">NamespacesEventhub
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/resourceDefinitions/namespaces_eventhubs">https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/resourceDefinitions/namespaces_eventhubs</a></p>
</div>
<table>
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
<a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubs_Spec">
NamespacesEventhubs_Spec
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
<code>captureDescription</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubs_Spec_Properties_CaptureDescription">
NamespacesEventhubs_Spec_Properties_CaptureDescription
</a>
</em>
</td>
<td>
<p>CaptureDescription: Properties to configure capture description for eventhub</p>
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
<code>messageRetentionInDays</code><br/>
<em>
int
</em>
</td>
<td>
<p>MessageRetentionInDays: Number of days to retain the events for this Event Hub, value should be 1 to 7 days</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
github.com/Azure/azure-service-operator/v2/pkg/genruntime.KnownResourceReference
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>partitionCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>PartitionCount: Number of partitions created for the Event Hub, allowed values are from 1 to 32 partitions.</p>
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
<a href="#eventhub.azure.com/v1alpha1api20211101.Eventhub_Status">
Eventhub_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubsAuthorizationRule">NamespacesEventhubsAuthorizationRule
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/resourceDefinitions/namespaces_eventhubs_authorizationRules">https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/resourceDefinitions/namespaces_eventhubs_authorizationRules</a></p>
</div>
<table>
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
<a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubsAuthorizationRules_Spec">
NamespacesEventhubsAuthorizationRules_Spec
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
github.com/Azure/azure-service-operator/v2/pkg/genruntime.KnownResourceReference
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>rights</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.AuthorizationRulePropertiesRights">
[]AuthorizationRulePropertiesRights
</a>
</em>
</td>
<td>
<p>Rights: The rights associated with the rule.</p>
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
<a href="#eventhub.azure.com/v1alpha1api20211101.AuthorizationRule_Status">
AuthorizationRule_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubsAuthorizationRulesSpecAPIVersion">NamespacesEventhubsAuthorizationRulesSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-11-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubsAuthorizationRules_Spec">NamespacesEventhubsAuthorizationRules_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubsAuthorizationRule">NamespacesEventhubsAuthorizationRule</a>)
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
github.com/Azure/azure-service-operator/v2/pkg/genruntime.KnownResourceReference
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>rights</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.AuthorizationRulePropertiesRights">
[]AuthorizationRulePropertiesRights
</a>
</em>
</td>
<td>
<p>Rights: The rights associated with the rule.</p>
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
<h3 id="eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubsAuthorizationRules_SpecARM">NamespacesEventhubsAuthorizationRules_SpecARM
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
<p>Name: The authorization rule name.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.AuthorizationRulePropertiesARM">
AuthorizationRulePropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties supplied to create or update AuthorizationRule</p>
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
<h3 id="eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubsConsumerGroup">NamespacesEventhubsConsumerGroup
</h3>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/resourceDefinitions/namespaces_eventhubs_consumergroups">https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/resourceDefinitions/namespaces_eventhubs_consumergroups</a></p>
</div>
<table>
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
<a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubsConsumergroups_Spec">
NamespacesEventhubsConsumergroups_Spec
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
github.com/Azure/azure-service-operator/v2/pkg/genruntime.KnownResourceReference
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
<tr>
<td>
<code>userMetadata</code><br/>
<em>
string
</em>
</td>
<td>
<p>UserMetadata: User Metadata is a placeholder to store user-defined string data with maximum length 1024. e.g. it can be
used to store descriptive data, such as list of teams and their contact information also user-defined configuration
settings can be stored.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.ConsumerGroup_Status">
ConsumerGroup_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubsConsumergroupsSpecAPIVersion">NamespacesEventhubsConsumergroupsSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-11-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubsConsumergroups_Spec">NamespacesEventhubsConsumergroups_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubsConsumerGroup">NamespacesEventhubsConsumerGroup</a>)
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
github.com/Azure/azure-service-operator/v2/pkg/genruntime.KnownResourceReference
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
<tr>
<td>
<code>userMetadata</code><br/>
<em>
string
</em>
</td>
<td>
<p>UserMetadata: User Metadata is a placeholder to store user-defined string data with maximum length 1024. e.g. it can be
used to store descriptive data, such as list of teams and their contact information also user-defined configuration
settings can be stored.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubsConsumergroups_SpecARM">NamespacesEventhubsConsumergroups_SpecARM
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
<p>Name: The consumer group name</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.ConsumerGroupPropertiesARM">
ConsumerGroupPropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Single item in List or Get Consumer group operation</p>
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
<h3 id="eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubsSpecAPIVersion">NamespacesEventhubsSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-11-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubsSpecPropertiesCaptureDescriptionEncoding">NamespacesEventhubsSpecPropertiesCaptureDescriptionEncoding
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubs_Spec_Properties_CaptureDescription">NamespacesEventhubs_Spec_Properties_CaptureDescription</a>, <a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubs_Spec_Properties_CaptureDescriptionARM">NamespacesEventhubs_Spec_Properties_CaptureDescriptionARM</a>)
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
<tbody><tr><td><p>&#34;Avro&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;AvroDeflate&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubs_Spec">NamespacesEventhubs_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhub">NamespacesEventhub</a>)
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
<code>captureDescription</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubs_Spec_Properties_CaptureDescription">
NamespacesEventhubs_Spec_Properties_CaptureDescription
</a>
</em>
</td>
<td>
<p>CaptureDescription: Properties to configure capture description for eventhub</p>
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
<code>messageRetentionInDays</code><br/>
<em>
int
</em>
</td>
<td>
<p>MessageRetentionInDays: Number of days to retain the events for this Event Hub, value should be 1 to 7 days</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
github.com/Azure/azure-service-operator/v2/pkg/genruntime.KnownResourceReference
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>partitionCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>PartitionCount: Number of partitions created for the Event Hub, allowed values are from 1 to 32 partitions.</p>
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
<h3 id="eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubs_SpecARM">NamespacesEventhubs_SpecARM
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
<p>Name: Name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubs_Spec_PropertiesARM">
NamespacesEventhubs_Spec_PropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties supplied to the Create Or Update Event Hub operation.</p>
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
<h3 id="eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubs_Spec_PropertiesARM">NamespacesEventhubs_Spec_PropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubs_SpecARM">NamespacesEventhubs_SpecARM</a>)
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
<code>captureDescription</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubs_Spec_Properties_CaptureDescriptionARM">
NamespacesEventhubs_Spec_Properties_CaptureDescriptionARM
</a>
</em>
</td>
<td>
<p>CaptureDescription: Properties to configure capture description for eventhub</p>
</td>
</tr>
<tr>
<td>
<code>messageRetentionInDays</code><br/>
<em>
int
</em>
</td>
<td>
<p>MessageRetentionInDays: Number of days to retain the events for this Event Hub, value should be 1 to 7 days</p>
</td>
</tr>
<tr>
<td>
<code>partitionCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>PartitionCount: Number of partitions created for the Event Hub, allowed values are from 1 to 32 partitions.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubs_Spec_Properties_CaptureDescription">NamespacesEventhubs_Spec_Properties_CaptureDescription
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubs_Spec">NamespacesEventhubs_Spec</a>)
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
<code>destination</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination">
NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination
</a>
</em>
</td>
<td>
<p>Destination: Capture storage details for capture description</p>
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
<p>Enabled: A value that indicates whether capture description is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>encoding</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubsSpecPropertiesCaptureDescriptionEncoding">
NamespacesEventhubsSpecPropertiesCaptureDescriptionEncoding
</a>
</em>
</td>
<td>
<p>Encoding: Enumerates the possible values for the encoding format of capture description. Note: &lsquo;AvroDeflate&rsquo; will be
deprecated in New API Version.</p>
</td>
</tr>
<tr>
<td>
<code>intervalInSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>IntervalInSeconds: The time window allows you to set the frequency with which the capture to Azure Blobs will happen,
value should between 60 to 900 seconds</p>
</td>
</tr>
<tr>
<td>
<code>sizeLimitInBytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>SizeLimitInBytes: The size window defines the amount of data built up in your Event Hub before an capture operation,
value should be between 10485760 to 524288000 bytes</p>
</td>
</tr>
<tr>
<td>
<code>skipEmptyArchives</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SkipEmptyArchives: A value that indicates whether to Skip Empty Archives</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubs_Spec_Properties_CaptureDescriptionARM">NamespacesEventhubs_Spec_Properties_CaptureDescriptionARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubs_Spec_PropertiesARM">NamespacesEventhubs_Spec_PropertiesARM</a>)
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
<code>destination</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubs_Spec_Properties_CaptureDescription_DestinationARM">
NamespacesEventhubs_Spec_Properties_CaptureDescription_DestinationARM
</a>
</em>
</td>
<td>
<p>Destination: Capture storage details for capture description</p>
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
<p>Enabled: A value that indicates whether capture description is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>encoding</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubsSpecPropertiesCaptureDescriptionEncoding">
NamespacesEventhubsSpecPropertiesCaptureDescriptionEncoding
</a>
</em>
</td>
<td>
<p>Encoding: Enumerates the possible values for the encoding format of capture description. Note: &lsquo;AvroDeflate&rsquo; will be
deprecated in New API Version.</p>
</td>
</tr>
<tr>
<td>
<code>intervalInSeconds</code><br/>
<em>
int
</em>
</td>
<td>
<p>IntervalInSeconds: The time window allows you to set the frequency with which the capture to Azure Blobs will happen,
value should between 60 to 900 seconds</p>
</td>
</tr>
<tr>
<td>
<code>sizeLimitInBytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>SizeLimitInBytes: The size window defines the amount of data built up in your Event Hub before an capture operation,
value should be between 10485760 to 524288000 bytes</p>
</td>
</tr>
<tr>
<td>
<code>skipEmptyArchives</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SkipEmptyArchives: A value that indicates whether to Skip Empty Archives</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination">NamespacesEventhubs_Spec_Properties_CaptureDescription_Destination
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubs_Spec_Properties_CaptureDescription">NamespacesEventhubs_Spec_Properties_CaptureDescription</a>)
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
<code>archiveNameFormat</code><br/>
<em>
string
</em>
</td>
<td>
<p>ArchiveNameFormat: Blob naming convention for archive, e.g.
{Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}. Here all the parameters
(Namespace,EventHub .. etc) are mandatory irrespective of order</p>
</td>
</tr>
<tr>
<td>
<code>blobContainer</code><br/>
<em>
string
</em>
</td>
<td>
<p>BlobContainer: Blob container Name</p>
</td>
</tr>
<tr>
<td>
<code>dataLakeAccountName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DataLakeAccountName: The Azure Data Lake Store name for the captured events</p>
</td>
</tr>
<tr>
<td>
<code>dataLakeFolderPath</code><br/>
<em>
string
</em>
</td>
<td>
<p>DataLakeFolderPath: The destination folder path for the captured events</p>
</td>
</tr>
<tr>
<td>
<code>dataLakeSubscriptionId</code><br/>
<em>
string
</em>
</td>
<td>
<p>DataLakeSubscriptionId: Subscription Id of Azure Data Lake Store</p>
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
<p>Name: Name for capture destination</p>
</td>
</tr>
<tr>
<td>
<code>storageAccountResourceReference</code><br/>
<em>
github.com/Azure/azure-service-operator/v2/pkg/genruntime.ResourceReference
</em>
</td>
<td>
<p>StorageAccountResourceReference: Resource id of the storage account to be used to create the blobs</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubs_Spec_Properties_CaptureDescription_DestinationARM">NamespacesEventhubs_Spec_Properties_CaptureDescription_DestinationARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.NamespacesEventhubs_Spec_Properties_CaptureDescriptionARM">NamespacesEventhubs_Spec_Properties_CaptureDescriptionARM</a>)
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
<p>Name: Name for capture destination</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.DestinationPropertiesARM">
DestinationPropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties describing the storage account, blob container and archive name format for capture destination</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.NamespacesSpecAPIVersion">NamespacesSpecAPIVersion
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
<tbody><tr><td><p>&#34;2021-11-01&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.Namespaces_Spec">Namespaces_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Namespace">Namespace</a>)
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
<code>alternateName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AlternateName: Alternate name specified when alias and namespace names are same.</p>
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
<code>clusterArmReference</code><br/>
<em>
github.com/Azure/azure-service-operator/v2/pkg/genruntime.ResourceReference
</em>
</td>
<td>
<p>ClusterArmReference: Cluster ARM ID of the Namespace.</p>
</td>
</tr>
<tr>
<td>
<code>disableLocalAuth</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableLocalAuth: This property disables SAS authentication for the Event Hubs namespace.</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.Encryption">
Encryption
</a>
</em>
</td>
<td>
<p>Encryption: Properties to configure Encryption</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.Identity">
Identity
</a>
</em>
</td>
<td>
<p>Identity: Properties to configure Identity for Bring your Own Keys</p>
</td>
</tr>
<tr>
<td>
<code>isAutoInflateEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsAutoInflateEnabled: Value that indicates whether AutoInflate is enabled for eventhub namespace.</p>
</td>
</tr>
<tr>
<td>
<code>kafkaEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>KafkaEnabled: Value that indicates whether Kafka is enabled for eventhub namespace.</p>
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
<code>maximumThroughputUnits</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaximumThroughputUnits: Upper limit of throughput units when AutoInflate is enabled, value should be within 0 to 20
throughput units. ( &lsquo;0&rsquo; if AutoInflateEnabled = true)</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
github.com/Azure/azure-service-operator/v2/pkg/genruntime.KnownResourceReference
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>privateEndpointConnections</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.Namespaces_Spec_Properties_PrivateEndpointConnections">
[]Namespaces_Spec_Properties_PrivateEndpointConnections
</a>
</em>
</td>
<td>
<p>PrivateEndpointConnections: List of private endpoint connections.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.Sku">
Sku
</a>
</em>
</td>
<td>
<p>Sku: SKU parameters supplied to the create namespace operation</p>
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
<code>zoneRedundant</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ZoneRedundant: Enabling this property creates a Standard Event Hubs Namespace in regions supported availability zones.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.Namespaces_SpecARM">Namespaces_SpecARM
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
<code>identity</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.IdentityARM">
IdentityARM
</a>
</em>
</td>
<td>
<p>Identity: Properties to configure Identity for Bring your Own Keys</p>
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
<p>Name: Name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.Namespaces_Spec_PropertiesARM">
Namespaces_Spec_PropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Namespace properties supplied for create namespace operation.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.SkuARM">
SkuARM
</a>
</em>
</td>
<td>
<p>Sku: SKU parameters supplied to the create namespace operation</p>
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
<h3 id="eventhub.azure.com/v1alpha1api20211101.Namespaces_Spec_PropertiesARM">Namespaces_Spec_PropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Namespaces_SpecARM">Namespaces_SpecARM</a>)
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
<code>alternateName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AlternateName: Alternate name specified when alias and namespace names are same.</p>
</td>
</tr>
<tr>
<td>
<code>clusterArmId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>disableLocalAuth</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DisableLocalAuth: This property disables SAS authentication for the Event Hubs namespace.</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.EncryptionARM">
EncryptionARM
</a>
</em>
</td>
<td>
<p>Encryption: Properties to configure Encryption</p>
</td>
</tr>
<tr>
<td>
<code>isAutoInflateEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsAutoInflateEnabled: Value that indicates whether AutoInflate is enabled for eventhub namespace.</p>
</td>
</tr>
<tr>
<td>
<code>kafkaEnabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>KafkaEnabled: Value that indicates whether Kafka is enabled for eventhub namespace.</p>
</td>
</tr>
<tr>
<td>
<code>maximumThroughputUnits</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaximumThroughputUnits: Upper limit of throughput units when AutoInflate is enabled, value should be within 0 to 20
throughput units. ( &lsquo;0&rsquo; if AutoInflateEnabled = true)</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpointConnections</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.Namespaces_Spec_Properties_PrivateEndpointConnectionsARM">
[]Namespaces_Spec_Properties_PrivateEndpointConnectionsARM
</a>
</em>
</td>
<td>
<p>PrivateEndpointConnections: List of private endpoint connections.</p>
</td>
</tr>
<tr>
<td>
<code>zoneRedundant</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ZoneRedundant: Enabling this property creates a Standard Event Hubs Namespace in regions supported availability zones.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.Namespaces_Spec_Properties_PrivateEndpointConnections">Namespaces_Spec_Properties_PrivateEndpointConnections
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Namespaces_Spec">Namespaces_Spec</a>)
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
<code>privateEndpoint</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.PrivateEndpoint">
PrivateEndpoint
</a>
</em>
</td>
<td>
<p>PrivateEndpoint: PrivateEndpoint information.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.Namespaces_Spec_Properties_PrivateEndpointConnectionsARM">Namespaces_Spec_Properties_PrivateEndpointConnectionsARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Namespaces_Spec_PropertiesARM">Namespaces_Spec_PropertiesARM</a>)
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
<code>properties</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.PrivateEndpointConnectionPropertiesARM">
PrivateEndpointConnectionPropertiesARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the private endpoint connection resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.PrivateEndpoint">PrivateEndpoint
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Namespaces_Spec_Properties_PrivateEndpointConnections">Namespaces_Spec_Properties_PrivateEndpointConnections</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/PrivateEndpoint">https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/PrivateEndpoint</a></p>
</div>
<table>
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
github.com/Azure/azure-service-operator/v2/pkg/genruntime.ResourceReference
</em>
</td>
<td>
<p>Reference: The ARM identifier for Private Endpoint.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.PrivateEndpointARM">PrivateEndpointARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.PrivateEndpointConnectionPropertiesARM">PrivateEndpointConnectionPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/PrivateEndpoint">https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/PrivateEndpoint</a></p>
</div>
<table>
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
<h3 id="eventhub.azure.com/v1alpha1api20211101.PrivateEndpointConnectionPropertiesARM">PrivateEndpointConnectionPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Namespaces_Spec_Properties_PrivateEndpointConnectionsARM">Namespaces_Spec_Properties_PrivateEndpointConnectionsARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/PrivateEndpointConnectionProperties">https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/PrivateEndpointConnectionProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>privateEndpoint</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.PrivateEndpointARM">
PrivateEndpointARM
</a>
</em>
</td>
<td>
<p>PrivateEndpoint: PrivateEndpoint information.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.PrivateEndpointConnection_Status_SubResourceEmbedded">PrivateEndpointConnection_Status_SubResourceEmbedded
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.EHNamespace_Status">EHNamespace_Status</a>)
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
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.SystemData_Status">
SystemData_Status
</a>
</em>
</td>
<td>
<p>SystemData: The system meta data relating to this resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.PrivateEndpointConnection_Status_SubResourceEmbeddedARM">PrivateEndpointConnection_Status_SubResourceEmbeddedARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.EHNamespace_Status_PropertiesARM">EHNamespace_Status_PropertiesARM</a>)
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
/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.SystemData_StatusARM">
SystemData_StatusARM
</a>
</em>
</td>
<td>
<p>SystemData: The system meta data relating to this resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.Sku">Sku
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Namespaces_Spec">Namespaces_Spec</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/Sku">https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/Sku</a></p>
</div>
<table>
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
<p>Capacity: The Event Hubs throughput units for Basic or Standard tiers, where value should be 0 to 20 throughput units.
The Event Hubs premium units for Premium tier, where value should be 0 to 10 premium units.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.SkuName">
SkuName
</a>
</em>
</td>
<td>
<p>Name: Name of this SKU.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.SkuTier">
SkuTier
</a>
</em>
</td>
<td>
<p>Tier: The billing tier of this particular SKU.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.SkuARM">SkuARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Namespaces_SpecARM">Namespaces_SpecARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/Sku">https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/Sku</a></p>
</div>
<table>
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
<p>Capacity: The Event Hubs throughput units for Basic or Standard tiers, where value should be 0 to 20 throughput units.
The Event Hubs premium units for Premium tier, where value should be 0 to 10 premium units.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.SkuName">
SkuName
</a>
</em>
</td>
<td>
<p>Name: Name of this SKU.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.SkuTier">
SkuTier
</a>
</em>
</td>
<td>
<p>Tier: The billing tier of this particular SKU.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.SkuName">SkuName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Sku">Sku</a>, <a href="#eventhub.azure.com/v1alpha1api20211101.SkuARM">SkuARM</a>)
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
<h3 id="eventhub.azure.com/v1alpha1api20211101.SkuStatusName">SkuStatusName
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Sku_Status">Sku_Status</a>, <a href="#eventhub.azure.com/v1alpha1api20211101.Sku_StatusARM">Sku_StatusARM</a>)
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
<h3 id="eventhub.azure.com/v1alpha1api20211101.SkuStatusTier">SkuStatusTier
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Sku_Status">Sku_Status</a>, <a href="#eventhub.azure.com/v1alpha1api20211101.Sku_StatusARM">Sku_StatusARM</a>)
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
<h3 id="eventhub.azure.com/v1alpha1api20211101.SkuTier">SkuTier
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Sku">Sku</a>, <a href="#eventhub.azure.com/v1alpha1api20211101.SkuARM">SkuARM</a>)
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
<h3 id="eventhub.azure.com/v1alpha1api20211101.Sku_Status">Sku_Status
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.EHNamespace_Status">EHNamespace_Status</a>)
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
<p>Capacity: The Event Hubs throughput units for Basic or Standard tiers, where value should be 0 to 20 throughput units.
The Event Hubs premium units for Premium tier, where value should be 0 to 10 premium units.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.SkuStatusName">
SkuStatusName
</a>
</em>
</td>
<td>
<p>Name: Name of this SKU.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.SkuStatusTier">
SkuStatusTier
</a>
</em>
</td>
<td>
<p>Tier: The billing tier of this particular SKU.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.Sku_StatusARM">Sku_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.EHNamespace_StatusARM">EHNamespace_StatusARM</a>)
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
<p>Capacity: The Event Hubs throughput units for Basic or Standard tiers, where value should be 0 to 20 throughput units.
The Event Hubs premium units for Premium tier, where value should be 0 to 10 premium units.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.SkuStatusName">
SkuStatusName
</a>
</em>
</td>
<td>
<p>Name: Name of this SKU.</p>
</td>
</tr>
<tr>
<td>
<code>tier</code><br/>
<em>
<a href="#eventhub.azure.com/v1alpha1api20211101.SkuStatusTier">
SkuStatusTier
</a>
</em>
</td>
<td>
<p>Tier: The billing tier of this particular SKU.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.SystemDataStatusCreatedByType">SystemDataStatusCreatedByType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.SystemData_Status">SystemData_Status</a>, <a href="#eventhub.azure.com/v1alpha1api20211101.SystemData_StatusARM">SystemData_StatusARM</a>)
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
<h3 id="eventhub.azure.com/v1alpha1api20211101.SystemDataStatusLastModifiedByType">SystemDataStatusLastModifiedByType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.SystemData_Status">SystemData_Status</a>, <a href="#eventhub.azure.com/v1alpha1api20211101.SystemData_StatusARM">SystemData_StatusARM</a>)
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
<h3 id="eventhub.azure.com/v1alpha1api20211101.SystemData_Status">SystemData_Status
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.AuthorizationRule_Status">AuthorizationRule_Status</a>, <a href="#eventhub.azure.com/v1alpha1api20211101.ConsumerGroup_Status">ConsumerGroup_Status</a>, <a href="#eventhub.azure.com/v1alpha1api20211101.EHNamespace_Status">EHNamespace_Status</a>, <a href="#eventhub.azure.com/v1alpha1api20211101.Eventhub_Status">Eventhub_Status</a>, <a href="#eventhub.azure.com/v1alpha1api20211101.PrivateEndpointConnection_Status_SubResourceEmbedded">PrivateEndpointConnection_Status_SubResourceEmbedded</a>)
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
<a href="#eventhub.azure.com/v1alpha1api20211101.SystemDataStatusCreatedByType">
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
<p>LastModifiedAt: The type of identity that last modified the resource.</p>
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
<a href="#eventhub.azure.com/v1alpha1api20211101.SystemDataStatusLastModifiedByType">
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
<h3 id="eventhub.azure.com/v1alpha1api20211101.SystemData_StatusARM">SystemData_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.AuthorizationRule_StatusARM">AuthorizationRule_StatusARM</a>, <a href="#eventhub.azure.com/v1alpha1api20211101.ConsumerGroup_StatusARM">ConsumerGroup_StatusARM</a>, <a href="#eventhub.azure.com/v1alpha1api20211101.EHNamespace_StatusARM">EHNamespace_StatusARM</a>, <a href="#eventhub.azure.com/v1alpha1api20211101.Eventhub_StatusARM">Eventhub_StatusARM</a>, <a href="#eventhub.azure.com/v1alpha1api20211101.PrivateEndpointConnection_Status_SubResourceEmbeddedARM">PrivateEndpointConnection_Status_SubResourceEmbeddedARM</a>)
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
<a href="#eventhub.azure.com/v1alpha1api20211101.SystemDataStatusCreatedByType">
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
<p>LastModifiedAt: The type of identity that last modified the resource.</p>
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
<a href="#eventhub.azure.com/v1alpha1api20211101.SystemDataStatusLastModifiedByType">
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
<h3 id="eventhub.azure.com/v1alpha1api20211101.UserAssignedIdentityProperties">UserAssignedIdentityProperties
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.KeyVaultProperties">KeyVaultProperties</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/UserAssignedIdentityProperties">https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/UserAssignedIdentityProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>userAssignedIdentityReference</code><br/>
<em>
github.com/Azure/azure-service-operator/v2/pkg/genruntime.ResourceReference
</em>
</td>
<td>
<p>UserAssignedIdentityReference: ARM ID of user Identity selected for encryption</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.UserAssignedIdentityPropertiesARM">UserAssignedIdentityPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.KeyVaultPropertiesARM">KeyVaultPropertiesARM</a>)
</p>
<div>
<p>Generated from: <a href="https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/UserAssignedIdentityProperties">https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/definitions/UserAssignedIdentityProperties</a></p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>userAssignedIdentity</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.UserAssignedIdentityProperties_Status">UserAssignedIdentityProperties_Status
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.KeyVaultProperties_Status">KeyVaultProperties_Status</a>)
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
<code>userAssignedIdentity</code><br/>
<em>
string
</em>
</td>
<td>
<p>UserAssignedIdentity: ARM ID of user Identity selected for encryption</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.UserAssignedIdentityProperties_StatusARM">UserAssignedIdentityProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.KeyVaultProperties_StatusARM">KeyVaultProperties_StatusARM</a>)
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
<code>userAssignedIdentity</code><br/>
<em>
string
</em>
</td>
<td>
<p>UserAssignedIdentity: ARM ID of user Identity selected for encryption</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.UserAssignedIdentity_Status">UserAssignedIdentity_Status
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Identity_Status">Identity_Status</a>)
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
<p>ClientId: Client Id of user assigned identity</p>
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
<p>PrincipalId: Principal Id of user assigned identity</p>
</td>
</tr>
</tbody>
</table>
<h3 id="eventhub.azure.com/v1alpha1api20211101.UserAssignedIdentity_StatusARM">UserAssignedIdentity_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#eventhub.azure.com/v1alpha1api20211101.Identity_StatusARM">Identity_StatusARM</a>)
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
<p>ClientId: Client Id of user assigned identity</p>
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
<p>PrincipalId: Principal Id of user assigned identity</p>
</td>
</tr>
</tbody>
</table>
<hr/>
