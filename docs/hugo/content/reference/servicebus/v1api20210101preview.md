---
title: servicebus.azure.com/v1api20210101preview
---
<h2 id="servicebus.azure.com/v1api20210101preview">servicebus.azure.com/v1api20210101preview</h2>
<div>
<p>Package v1api20210101preview contains API Schema definitions for the servicebus v1api20210101preview API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="servicebus.azure.com/v1api20210101preview.APIVersion">APIVersion
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
<tbody><tr><td><p>&#34;2021-01-01-preview&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Action">Action
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscriptions_Rule_Spec">Namespaces_Topics_Subscriptions_Rule_Spec</a>)
</p>
<div>
<p>Represents the filter actions which are allowed for the transformation of a message that have been matched by a filter
expression.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>compatibilityLevel</code><br/>
<em>
int
</em>
</td>
<td>
<p>CompatibilityLevel: This property is reserved for future use. An integer value showing the compatibility level,
currently hard-coded to 20.</p>
</td>
</tr>
<tr>
<td>
<code>requiresPreprocessing</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresPreprocessing: Value that indicates whether the rule action requires preprocessing.</p>
</td>
</tr>
<tr>
<td>
<code>sqlExpression</code><br/>
<em>
string
</em>
</td>
<td>
<p>SqlExpression: SQL expression. e.g. MyProperty=&lsquo;ABC&rsquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Action_ARM">Action_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Ruleproperties_ARM">Ruleproperties_ARM</a>)
</p>
<div>
<p>Represents the filter actions which are allowed for the transformation of a message that have been matched by a filter
expression.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>compatibilityLevel</code><br/>
<em>
int
</em>
</td>
<td>
<p>CompatibilityLevel: This property is reserved for future use. An integer value showing the compatibility level,
currently hard-coded to 20.</p>
</td>
</tr>
<tr>
<td>
<code>requiresPreprocessing</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresPreprocessing: Value that indicates whether the rule action requires preprocessing.</p>
</td>
</tr>
<tr>
<td>
<code>sqlExpression</code><br/>
<em>
string
</em>
</td>
<td>
<p>SqlExpression: SQL expression. e.g. MyProperty=&lsquo;ABC&rsquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Action_STATUS">Action_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscriptions_Rule_STATUS">Namespaces_Topics_Subscriptions_Rule_STATUS</a>)
</p>
<div>
<p>Represents the filter actions which are allowed for the transformation of a message that have been matched by a filter
expression.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>compatibilityLevel</code><br/>
<em>
int
</em>
</td>
<td>
<p>CompatibilityLevel: This property is reserved for future use. An integer value showing the compatibility level,
currently hard-coded to 20.</p>
</td>
</tr>
<tr>
<td>
<code>requiresPreprocessing</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresPreprocessing: Value that indicates whether the rule action requires preprocessing.</p>
</td>
</tr>
<tr>
<td>
<code>sqlExpression</code><br/>
<em>
string
</em>
</td>
<td>
<p>SqlExpression: SQL expression. e.g. MyProperty=&lsquo;ABC&rsquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Action_STATUS_ARM">Action_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Ruleproperties_STATUS_ARM">Ruleproperties_STATUS_ARM</a>)
</p>
<div>
<p>Represents the filter actions which are allowed for the transformation of a message that have been matched by a filter
expression.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>compatibilityLevel</code><br/>
<em>
int
</em>
</td>
<td>
<p>CompatibilityLevel: This property is reserved for future use. An integer value showing the compatibility level,
currently hard-coded to 20.</p>
</td>
</tr>
<tr>
<td>
<code>requiresPreprocessing</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresPreprocessing: Value that indicates whether the rule action requires preprocessing.</p>
</td>
</tr>
<tr>
<td>
<code>sqlExpression</code><br/>
<em>
string
</em>
</td>
<td>
<p>SqlExpression: SQL expression. e.g. MyProperty=&lsquo;ABC&rsquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.CorrelationFilter">CorrelationFilter
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscriptions_Rule_Spec">Namespaces_Topics_Subscriptions_Rule_Spec</a>)
</p>
<div>
<p>Represents the correlation filter expression.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>contentType</code><br/>
<em>
string
</em>
</td>
<td>
<p>ContentType: Content type of the message.</p>
</td>
</tr>
<tr>
<td>
<code>correlationId</code><br/>
<em>
string
</em>
</td>
<td>
<p>CorrelationId: Identifier of the correlation.</p>
</td>
</tr>
<tr>
<td>
<code>label</code><br/>
<em>
string
</em>
</td>
<td>
<p>Label: Application specific label.</p>
</td>
</tr>
<tr>
<td>
<code>messageId</code><br/>
<em>
string
</em>
</td>
<td>
<p>MessageId: Identifier of the message.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Properties: dictionary object for custom filters</p>
</td>
</tr>
<tr>
<td>
<code>replyTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ReplyTo: Address of the queue to reply to.</p>
</td>
</tr>
<tr>
<td>
<code>replyToSessionId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ReplyToSessionId: Session identifier to reply to.</p>
</td>
</tr>
<tr>
<td>
<code>requiresPreprocessing</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresPreprocessing: Value that indicates whether the rule action requires preprocessing.</p>
</td>
</tr>
<tr>
<td>
<code>sessionId</code><br/>
<em>
string
</em>
</td>
<td>
<p>SessionId: Session identifier.</p>
</td>
</tr>
<tr>
<td>
<code>to</code><br/>
<em>
string
</em>
</td>
<td>
<p>To: Address to send to.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.CorrelationFilter_ARM">CorrelationFilter_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Ruleproperties_ARM">Ruleproperties_ARM</a>)
</p>
<div>
<p>Represents the correlation filter expression.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>contentType</code><br/>
<em>
string
</em>
</td>
<td>
<p>ContentType: Content type of the message.</p>
</td>
</tr>
<tr>
<td>
<code>correlationId</code><br/>
<em>
string
</em>
</td>
<td>
<p>CorrelationId: Identifier of the correlation.</p>
</td>
</tr>
<tr>
<td>
<code>label</code><br/>
<em>
string
</em>
</td>
<td>
<p>Label: Application specific label.</p>
</td>
</tr>
<tr>
<td>
<code>messageId</code><br/>
<em>
string
</em>
</td>
<td>
<p>MessageId: Identifier of the message.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Properties: dictionary object for custom filters</p>
</td>
</tr>
<tr>
<td>
<code>replyTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ReplyTo: Address of the queue to reply to.</p>
</td>
</tr>
<tr>
<td>
<code>replyToSessionId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ReplyToSessionId: Session identifier to reply to.</p>
</td>
</tr>
<tr>
<td>
<code>requiresPreprocessing</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresPreprocessing: Value that indicates whether the rule action requires preprocessing.</p>
</td>
</tr>
<tr>
<td>
<code>sessionId</code><br/>
<em>
string
</em>
</td>
<td>
<p>SessionId: Session identifier.</p>
</td>
</tr>
<tr>
<td>
<code>to</code><br/>
<em>
string
</em>
</td>
<td>
<p>To: Address to send to.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.CorrelationFilter_STATUS">CorrelationFilter_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscriptions_Rule_STATUS">Namespaces_Topics_Subscriptions_Rule_STATUS</a>)
</p>
<div>
<p>Represents the correlation filter expression.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>contentType</code><br/>
<em>
string
</em>
</td>
<td>
<p>ContentType: Content type of the message.</p>
</td>
</tr>
<tr>
<td>
<code>correlationId</code><br/>
<em>
string
</em>
</td>
<td>
<p>CorrelationId: Identifier of the correlation.</p>
</td>
</tr>
<tr>
<td>
<code>label</code><br/>
<em>
string
</em>
</td>
<td>
<p>Label: Application specific label.</p>
</td>
</tr>
<tr>
<td>
<code>messageId</code><br/>
<em>
string
</em>
</td>
<td>
<p>MessageId: Identifier of the message.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Properties: dictionary object for custom filters</p>
</td>
</tr>
<tr>
<td>
<code>replyTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ReplyTo: Address of the queue to reply to.</p>
</td>
</tr>
<tr>
<td>
<code>replyToSessionId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ReplyToSessionId: Session identifier to reply to.</p>
</td>
</tr>
<tr>
<td>
<code>requiresPreprocessing</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresPreprocessing: Value that indicates whether the rule action requires preprocessing.</p>
</td>
</tr>
<tr>
<td>
<code>sessionId</code><br/>
<em>
string
</em>
</td>
<td>
<p>SessionId: Session identifier.</p>
</td>
</tr>
<tr>
<td>
<code>to</code><br/>
<em>
string
</em>
</td>
<td>
<p>To: Address to send to.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.CorrelationFilter_STATUS_ARM">CorrelationFilter_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Ruleproperties_STATUS_ARM">Ruleproperties_STATUS_ARM</a>)
</p>
<div>
<p>Represents the correlation filter expression.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>contentType</code><br/>
<em>
string
</em>
</td>
<td>
<p>ContentType: Content type of the message.</p>
</td>
</tr>
<tr>
<td>
<code>correlationId</code><br/>
<em>
string
</em>
</td>
<td>
<p>CorrelationId: Identifier of the correlation.</p>
</td>
</tr>
<tr>
<td>
<code>label</code><br/>
<em>
string
</em>
</td>
<td>
<p>Label: Application specific label.</p>
</td>
</tr>
<tr>
<td>
<code>messageId</code><br/>
<em>
string
</em>
</td>
<td>
<p>MessageId: Identifier of the message.</p>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Properties: dictionary object for custom filters</p>
</td>
</tr>
<tr>
<td>
<code>replyTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ReplyTo: Address of the queue to reply to.</p>
</td>
</tr>
<tr>
<td>
<code>replyToSessionId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ReplyToSessionId: Session identifier to reply to.</p>
</td>
</tr>
<tr>
<td>
<code>requiresPreprocessing</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresPreprocessing: Value that indicates whether the rule action requires preprocessing.</p>
</td>
</tr>
<tr>
<td>
<code>sessionId</code><br/>
<em>
string
</em>
</td>
<td>
<p>SessionId: Session identifier.</p>
</td>
</tr>
<tr>
<td>
<code>to</code><br/>
<em>
string
</em>
</td>
<td>
<p>To: Address to send to.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.DictionaryValue_STATUS">DictionaryValue_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Identity_STATUS">Identity_STATUS</a>)
</p>
<div>
<p>Recognized Dictionary value.</p>
</div>
<table>
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
<h3 id="servicebus.azure.com/v1api20210101preview.DictionaryValue_STATUS_ARM">DictionaryValue_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Identity_STATUS_ARM">Identity_STATUS_ARM</a>)
</p>
<div>
<p>Recognized Dictionary value.</p>
</div>
<table>
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
<h3 id="servicebus.azure.com/v1api20210101preview.Encryption">Encryption
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespace_Spec">Namespace_Spec</a>)
</p>
<div>
<p>Properties to configure Encryption</p>
</div>
<table>
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
<a href="#servicebus.azure.com/v1api20210101preview.Encryption_KeySource">
Encryption_KeySource
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
<a href="#servicebus.azure.com/v1api20210101preview.KeyVaultProperties">
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
<h3 id="servicebus.azure.com/v1api20210101preview.Encryption_ARM">Encryption_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.SBNamespaceProperties_ARM">SBNamespaceProperties_ARM</a>)
</p>
<div>
<p>Properties to configure Encryption</p>
</div>
<table>
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
<a href="#servicebus.azure.com/v1api20210101preview.Encryption_KeySource">
Encryption_KeySource
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
<a href="#servicebus.azure.com/v1api20210101preview.KeyVaultProperties_ARM">
[]KeyVaultProperties_ARM
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
<h3 id="servicebus.azure.com/v1api20210101preview.Encryption_KeySource">Encryption_KeySource
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Encryption">Encryption</a>, <a href="#servicebus.azure.com/v1api20210101preview.Encryption_ARM">Encryption_ARM</a>)
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
<h3 id="servicebus.azure.com/v1api20210101preview.Encryption_KeySource_STATUS">Encryption_KeySource_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Encryption_STATUS">Encryption_STATUS</a>, <a href="#servicebus.azure.com/v1api20210101preview.Encryption_STATUS_ARM">Encryption_STATUS_ARM</a>)
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
<h3 id="servicebus.azure.com/v1api20210101preview.Encryption_STATUS">Encryption_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespace_STATUS">Namespace_STATUS</a>)
</p>
<div>
<p>Properties to configure Encryption</p>
</div>
<table>
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
<a href="#servicebus.azure.com/v1api20210101preview.Encryption_KeySource_STATUS">
Encryption_KeySource_STATUS
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
<a href="#servicebus.azure.com/v1api20210101preview.KeyVaultProperties_STATUS">
[]KeyVaultProperties_STATUS
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
<h3 id="servicebus.azure.com/v1api20210101preview.Encryption_STATUS_ARM">Encryption_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.SBNamespaceProperties_STATUS_ARM">SBNamespaceProperties_STATUS_ARM</a>)
</p>
<div>
<p>Properties to configure Encryption</p>
</div>
<table>
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
<a href="#servicebus.azure.com/v1api20210101preview.Encryption_KeySource_STATUS">
Encryption_KeySource_STATUS
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
<a href="#servicebus.azure.com/v1api20210101preview.KeyVaultProperties_STATUS_ARM">
[]KeyVaultProperties_STATUS_ARM
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
<h3 id="servicebus.azure.com/v1api20210101preview.EntityStatus_STATUS">EntityStatus_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Queue_STATUS">Namespaces_Queue_STATUS</a>, <a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topic_STATUS">Namespaces_Topic_STATUS</a>, <a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscription_STATUS">Namespaces_Topics_Subscription_STATUS</a>, <a href="#servicebus.azure.com/v1api20210101preview.SBQueueProperties_STATUS_ARM">SBQueueProperties_STATUS_ARM</a>, <a href="#servicebus.azure.com/v1api20210101preview.SBSubscriptionProperties_STATUS_ARM">SBSubscriptionProperties_STATUS_ARM</a>, <a href="#servicebus.azure.com/v1api20210101preview.SBTopicProperties_STATUS_ARM">SBTopicProperties_STATUS_ARM</a>)
</p>
<div>
<p>Entity status.</p>
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
<h3 id="servicebus.azure.com/v1api20210101preview.FilterType">FilterType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscriptions_Rule_Spec">Namespaces_Topics_Subscriptions_Rule_Spec</a>, <a href="#servicebus.azure.com/v1api20210101preview.Ruleproperties_ARM">Ruleproperties_ARM</a>)
</p>
<div>
<p>Rule filter types</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;CorrelationFilter&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SqlFilter&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.FilterType_STATUS">FilterType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscriptions_Rule_STATUS">Namespaces_Topics_Subscriptions_Rule_STATUS</a>, <a href="#servicebus.azure.com/v1api20210101preview.Ruleproperties_STATUS_ARM">Ruleproperties_STATUS_ARM</a>)
</p>
<div>
<p>Rule filter types</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;CorrelationFilter&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SqlFilter&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Identity">Identity
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespace_Spec">Namespace_Spec</a>)
</p>
<div>
<p>Properties to configure User Assigned Identities for Bring your Own Keys</p>
</div>
<table>
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
<a href="#servicebus.azure.com/v1api20210101preview.Identity_Type">
Identity_Type
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
<a href="#servicebus.azure.com/v1api20210101preview.UserAssignedIdentityDetails">
[]UserAssignedIdentityDetails
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: Properties for User Assigned Identities</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Identity_ARM">Identity_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespace_Spec_ARM">Namespace_Spec_ARM</a>)
</p>
<div>
<p>Properties to configure User Assigned Identities for Bring your Own Keys</p>
</div>
<table>
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
<a href="#servicebus.azure.com/v1api20210101preview.Identity_Type">
Identity_Type
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
<a href="#servicebus.azure.com/v1api20210101preview.UserAssignedIdentityDetails_ARM">
map[string]./api/servicebus/v1api20210101preview.UserAssignedIdentityDetails_ARM
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Identity_STATUS">Identity_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespace_STATUS">Namespace_STATUS</a>)
</p>
<div>
<p>Properties to configure User Assigned Identities for Bring your Own Keys</p>
</div>
<table>
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
<a href="#servicebus.azure.com/v1api20210101preview.Identity_Type_STATUS">
Identity_Type_STATUS
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
<a href="#servicebus.azure.com/v1api20210101preview.DictionaryValue_STATUS">
map[string]./api/servicebus/v1api20210101preview.DictionaryValue_STATUS
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: Properties for User Assigned Identities</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Identity_STATUS_ARM">Identity_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespace_STATUS_ARM">Namespace_STATUS_ARM</a>)
</p>
<div>
<p>Properties to configure User Assigned Identities for Bring your Own Keys</p>
</div>
<table>
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
<a href="#servicebus.azure.com/v1api20210101preview.Identity_Type_STATUS">
Identity_Type_STATUS
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
<a href="#servicebus.azure.com/v1api20210101preview.DictionaryValue_STATUS_ARM">
map[string]./api/servicebus/v1api20210101preview.DictionaryValue_STATUS_ARM
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: Properties for User Assigned Identities</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Identity_Type">Identity_Type
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Identity">Identity</a>, <a href="#servicebus.azure.com/v1api20210101preview.Identity_ARM">Identity_ARM</a>)
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
<h3 id="servicebus.azure.com/v1api20210101preview.Identity_Type_STATUS">Identity_Type_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Identity_STATUS">Identity_STATUS</a>, <a href="#servicebus.azure.com/v1api20210101preview.Identity_STATUS_ARM">Identity_STATUS_ARM</a>)
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
<h3 id="servicebus.azure.com/v1api20210101preview.KeyVaultProperties">KeyVaultProperties
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Encryption">Encryption</a>)
</p>
<div>
<p>Properties to configure keyVault Properties</p>
</div>
<table>
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
<a href="#servicebus.azure.com/v1api20210101preview.UserAssignedIdentityProperties">
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
<p>KeyVersion: Version of KeyVault</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.KeyVaultProperties_ARM">KeyVaultProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Encryption_ARM">Encryption_ARM</a>)
</p>
<div>
<p>Properties to configure keyVault Properties</p>
</div>
<table>
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
<a href="#servicebus.azure.com/v1api20210101preview.UserAssignedIdentityProperties_ARM">
UserAssignedIdentityProperties_ARM
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
<p>KeyVersion: Version of KeyVault</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.KeyVaultProperties_STATUS">KeyVaultProperties_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Encryption_STATUS">Encryption_STATUS</a>)
</p>
<div>
<p>Properties to configure keyVault Properties</p>
</div>
<table>
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
<a href="#servicebus.azure.com/v1api20210101preview.UserAssignedIdentityProperties_STATUS">
UserAssignedIdentityProperties_STATUS
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
<p>KeyVersion: Version of KeyVault</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.KeyVaultProperties_STATUS_ARM">KeyVaultProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Encryption_STATUS_ARM">Encryption_STATUS_ARM</a>)
</p>
<div>
<p>Properties to configure keyVault Properties</p>
</div>
<table>
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
<a href="#servicebus.azure.com/v1api20210101preview.UserAssignedIdentityProperties_STATUS_ARM">
UserAssignedIdentityProperties_STATUS_ARM
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
<p>KeyVersion: Version of KeyVault</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.MessageCountDetails_STATUS">MessageCountDetails_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Queue_STATUS">Namespaces_Queue_STATUS</a>, <a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topic_STATUS">Namespaces_Topic_STATUS</a>, <a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscription_STATUS">Namespaces_Topics_Subscription_STATUS</a>)
</p>
<div>
<p>Message Count Details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>activeMessageCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>ActiveMessageCount: Number of active messages in the queue, topic, or subscription.</p>
</td>
</tr>
<tr>
<td>
<code>deadLetterMessageCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>DeadLetterMessageCount: Number of messages that are dead lettered.</p>
</td>
</tr>
<tr>
<td>
<code>scheduledMessageCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>ScheduledMessageCount: Number of scheduled messages.</p>
</td>
</tr>
<tr>
<td>
<code>transferDeadLetterMessageCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>TransferDeadLetterMessageCount: Number of messages transferred into dead letters.</p>
</td>
</tr>
<tr>
<td>
<code>transferMessageCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>TransferMessageCount: Number of messages transferred to another queue, topic, or subscription.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.MessageCountDetails_STATUS_ARM">MessageCountDetails_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.SBQueueProperties_STATUS_ARM">SBQueueProperties_STATUS_ARM</a>, <a href="#servicebus.azure.com/v1api20210101preview.SBSubscriptionProperties_STATUS_ARM">SBSubscriptionProperties_STATUS_ARM</a>, <a href="#servicebus.azure.com/v1api20210101preview.SBTopicProperties_STATUS_ARM">SBTopicProperties_STATUS_ARM</a>)
</p>
<div>
<p>Message Count Details.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>activeMessageCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>ActiveMessageCount: Number of active messages in the queue, topic, or subscription.</p>
</td>
</tr>
<tr>
<td>
<code>deadLetterMessageCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>DeadLetterMessageCount: Number of messages that are dead lettered.</p>
</td>
</tr>
<tr>
<td>
<code>scheduledMessageCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>ScheduledMessageCount: Number of scheduled messages.</p>
</td>
</tr>
<tr>
<td>
<code>transferDeadLetterMessageCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>TransferDeadLetterMessageCount: Number of messages transferred into dead letters.</p>
</td>
</tr>
<tr>
<td>
<code>transferMessageCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>TransferMessageCount: Number of messages transferred to another queue, topic, or subscription.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Namespace">Namespace
</h3>
<div>
<p>Generator information:
- Generated from: /servicebus/resource-manager/Microsoft.ServiceBus/preview/2021-01-01-preview/namespace-preview.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ServiceBus/&#x200b;namespaces/&#x200b;{namespaceName}</&#x200b;p>
</div>
<table>
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
<a href="#servicebus.azure.com/v1api20210101preview.Namespace_Spec">
Namespace_Spec
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
<code>encryption</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.Encryption">
Encryption
</a>
</em>
</td>
<td>
<p>Encryption: Properties of BYOK Encryption description</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.Identity">
Identity
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
<p>Location: The Geo-location where the resource lives</p>
</td>
</tr>
<tr>
<td>
<code>operatorSpec</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.NamespaceOperatorSpec">
NamespaceOperatorSpec
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
<code>sku</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.SBSku">
SBSku
</a>
</em>
</td>
<td>
<p>Sku: Properties of SKU</p>
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
<code>zoneRedundant</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ZoneRedundant: Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.Namespace_STATUS">
Namespace_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.NamespaceOperatorSecrets">NamespaceOperatorSecrets
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.NamespaceOperatorSpec">NamespaceOperatorSpec</a>)
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
<code>endpoint</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#SecretDestination">
genruntime.SecretDestination
</a>
</em>
</td>
<td>
<p>Endpoint: indicates where the Endpoint secret should be placed. If omitted, the secret will not be retrieved from Azure.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.NamespaceOperatorSpec">NamespaceOperatorSpec
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespace_Spec">Namespace_Spec</a>)
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
<a href="#servicebus.azure.com/v1api20210101preview.NamespaceOperatorSecrets">
NamespaceOperatorSecrets
</a>
</em>
</td>
<td>
<p>Secrets: configures where to place Azure generated secrets.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Namespace_STATUS">Namespace_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespace">Namespace</a>)
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
<code>createdAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedAt: The time the namespace was created</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.Encryption_STATUS">
Encryption_STATUS
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
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.Identity_STATUS">
Identity_STATUS
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
<p>Location: The Geo-location where the resource lives</p>
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
<p>MetricId: Identifier for Azure Insights metrics</p>
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
<code>privateEndpointConnections</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.PrivateEndpointConnection_STATUS">
[]PrivateEndpointConnection_STATUS
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
<p>ProvisioningState: Provisioning state of the namespace.</p>
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
<a href="#servicebus.azure.com/v1api20210101preview.SBSku_STATUS">
SBSku_STATUS
</a>
</em>
</td>
<td>
<p>Sku: Properties of SKU</p>
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
<p>Status: Status of the namespace.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.SystemData_STATUS">
SystemData_STATUS
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
<code>updatedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>UpdatedAt: The time the namespace was updated.</p>
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
<p>ZoneRedundant: Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Namespace_STATUS_ARM">Namespace_STATUS_ARM
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
<p>Id: Resource Id</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.Identity_STATUS_ARM">
Identity_STATUS_ARM
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
<p>Location: The Geo-location where the resource lives</p>
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
<a href="#servicebus.azure.com/v1api20210101preview.SBNamespaceProperties_STATUS_ARM">
SBNamespaceProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the namespace.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.SBSku_STATUS_ARM">
SBSku_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Sku: Properties of SKU</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.SystemData_STATUS_ARM">
SystemData_STATUS_ARM
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
<h3 id="servicebus.azure.com/v1api20210101preview.Namespace_Spec">Namespace_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespace">Namespace</a>)
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
<code>encryption</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.Encryption">
Encryption
</a>
</em>
</td>
<td>
<p>Encryption: Properties of BYOK Encryption description</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.Identity">
Identity
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
<p>Location: The Geo-location where the resource lives</p>
</td>
</tr>
<tr>
<td>
<code>operatorSpec</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.NamespaceOperatorSpec">
NamespaceOperatorSpec
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
<code>sku</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.SBSku">
SBSku
</a>
</em>
</td>
<td>
<p>Sku: Properties of SKU</p>
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
<code>zoneRedundant</code><br/>
<em>
bool
</em>
</td>
<td>
<p>ZoneRedundant: Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Namespace_Spec_ARM">Namespace_Spec_ARM
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
<a href="#servicebus.azure.com/v1api20210101preview.Identity_ARM">
Identity_ARM
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
<p>Location: The Geo-location where the resource lives</p>
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
<a href="#servicebus.azure.com/v1api20210101preview.SBNamespaceProperties_ARM">
SBNamespaceProperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of the namespace.</p>
</td>
</tr>
<tr>
<td>
<code>sku</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.SBSku_ARM">
SBSku_ARM
</a>
</em>
</td>
<td>
<p>Sku: Properties of SKU</p>
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
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.NamespacesQueue">NamespacesQueue
</h3>
<div>
<p>Generator information:
- Generated from: /servicebus/resource-manager/Microsoft.ServiceBus/preview/2021-01-01-preview/Queue.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ServiceBus/&#x200b;namespaces/&#x200b;{namespaceName}/&#x200b;queues/&#x200b;{queueName}</&#x200b;p>
</div>
<table>
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
<a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Queue_Spec">
Namespaces_Queue_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>autoDeleteOnIdle</code><br/>
<em>
string
</em>
</td>
<td>
<p>AutoDeleteOnIdle: ISO 8061 timeSpan idle interval after which the queue is automatically deleted. The minimum duration
is 5 minutes.</p>
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
<code>deadLetteringOnMessageExpiration</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DeadLetteringOnMessageExpiration: A value that indicates whether this queue has dead letter support when a message
expires.</p>
</td>
</tr>
<tr>
<td>
<code>defaultMessageTimeToLive</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultMessageTimeToLive: ISO 8601 default message timespan to live value. This is the duration after which the message
expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
set on a message itself.</p>
</td>
</tr>
<tr>
<td>
<code>duplicateDetectionHistoryTimeWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>DuplicateDetectionHistoryTimeWindow: ISO 8601 timeSpan structure that defines the duration of the duplicate detection
history. The default value is 10 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>enableBatchedOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableExpress</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableExpress: A value that indicates whether Express Entities are enabled. An express queue holds a message in memory
temporarily before writing it to persistent storage.</p>
</td>
</tr>
<tr>
<td>
<code>enablePartitioning</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePartitioning: A value that indicates whether the queue is to be partitioned across multiple message brokers.</p>
</td>
</tr>
<tr>
<td>
<code>forwardDeadLetteredMessagesTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardDeadLetteredMessagesTo: Queue/Topic name to forward the Dead Letter message</p>
</td>
</tr>
<tr>
<td>
<code>forwardTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardTo: Queue/Topic name to forward the messages</p>
</td>
</tr>
<tr>
<td>
<code>lockDuration</code><br/>
<em>
string
</em>
</td>
<td>
<p>LockDuration: ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for
other receivers. The maximum value for LockDuration is 5 minutes; the default value is 1 minute.</p>
</td>
</tr>
<tr>
<td>
<code>maxDeliveryCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxDeliveryCount: The maximum delivery count. A message is automatically deadlettered after this number of deliveries.
default value is 10.</p>
</td>
</tr>
<tr>
<td>
<code>maxSizeInMegabytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxSizeInMegabytes: The maximum size of the queue in megabytes, which is the size of memory allocated for the queue.
Default is 1024.</p>
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
reference to a servicebus.azure.com/Namespace resource</p>
</td>
</tr>
<tr>
<td>
<code>requiresDuplicateDetection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresDuplicateDetection: A value indicating if this queue requires duplicate detection.</p>
</td>
</tr>
<tr>
<td>
<code>requiresSession</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresSession: A value that indicates whether the queue supports the concept of sessions.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Queue_STATUS">
Namespaces_Queue_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.NamespacesTopic">NamespacesTopic
</h3>
<div>
<p>Generator information:
- Generated from: /servicebus/resource-manager/Microsoft.ServiceBus/preview/2021-01-01-preview/topics.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ServiceBus/&#x200b;namespaces/&#x200b;{namespaceName}/&#x200b;topics/&#x200b;{topicName}</&#x200b;p>
</div>
<table>
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
<a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topic_Spec">
Namespaces_Topic_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>autoDeleteOnIdle</code><br/>
<em>
string
</em>
</td>
<td>
<p>AutoDeleteOnIdle: ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration
is 5 minutes.</p>
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
<code>defaultMessageTimeToLive</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultMessageTimeToLive: ISO 8601 Default message timespan to live value. This is the duration after which the message
expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
set on a message itself.</p>
</td>
</tr>
<tr>
<td>
<code>duplicateDetectionHistoryTimeWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>DuplicateDetectionHistoryTimeWindow: ISO8601 timespan structure that defines the duration of the duplicate detection
history. The default value is 10 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>enableBatchedOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableExpress</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableExpress: Value that indicates whether Express Entities are enabled. An express topic holds a message in memory
temporarily before writing it to persistent storage.</p>
</td>
</tr>
<tr>
<td>
<code>enablePartitioning</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePartitioning: Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>maxSizeInMegabytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxSizeInMegabytes: Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic.
Default is 1024.</p>
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
reference to a servicebus.azure.com/Namespace resource</p>
</td>
</tr>
<tr>
<td>
<code>requiresDuplicateDetection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresDuplicateDetection: Value indicating if this topic requires duplicate detection.</p>
</td>
</tr>
<tr>
<td>
<code>supportOrdering</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SupportOrdering: Value that indicates whether the topic supports ordering.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topic_STATUS">
Namespaces_Topic_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.NamespacesTopicsSubscription">NamespacesTopicsSubscription
</h3>
<div>
<p>Generator information:
- Generated from: /servicebus/resource-manager/Microsoft.ServiceBus/preview/2021-01-01-preview/subscriptions.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ServiceBus/&#x200b;namespaces/&#x200b;{namespaceName}/&#x200b;topics/&#x200b;{topicName}/&#x200b;subscriptions/&#x200b;{subscriptionName}</&#x200b;p>
</div>
<table>
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
<a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscription_Spec">
Namespaces_Topics_Subscription_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>autoDeleteOnIdle</code><br/>
<em>
string
</em>
</td>
<td>
<p>AutoDeleteOnIdle: ISO 8061 timeSpan idle interval after which the topic is automatically deleted. The minimum duration
is 5 minutes.</p>
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
<code>deadLetteringOnFilterEvaluationExceptions</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DeadLetteringOnFilterEvaluationExceptions: Value that indicates whether a subscription has dead letter support on filter
evaluation exceptions.</p>
</td>
</tr>
<tr>
<td>
<code>deadLetteringOnMessageExpiration</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DeadLetteringOnMessageExpiration: Value that indicates whether a subscription has dead letter support when a message
expires.</p>
</td>
</tr>
<tr>
<td>
<code>defaultMessageTimeToLive</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultMessageTimeToLive: ISO 8061 Default message timespan to live value. This is the duration after which the message
expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
set on a message itself.</p>
</td>
</tr>
<tr>
<td>
<code>duplicateDetectionHistoryTimeWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>DuplicateDetectionHistoryTimeWindow: ISO 8601 timeSpan structure that defines the duration of the duplicate detection
history. The default value is 10 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>enableBatchedOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.</p>
</td>
</tr>
<tr>
<td>
<code>forwardDeadLetteredMessagesTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardDeadLetteredMessagesTo: Queue/Topic name to forward the Dead Letter message</p>
</td>
</tr>
<tr>
<td>
<code>forwardTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardTo: Queue/Topic name to forward the messages</p>
</td>
</tr>
<tr>
<td>
<code>lockDuration</code><br/>
<em>
string
</em>
</td>
<td>
<p>LockDuration: ISO 8061 lock duration timespan for the subscription. The default value is 1 minute.</p>
</td>
</tr>
<tr>
<td>
<code>maxDeliveryCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxDeliveryCount: Number of maximum deliveries.</p>
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
reference to a servicebus.azure.com/NamespacesTopic resource</p>
</td>
</tr>
<tr>
<td>
<code>requiresSession</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresSession: Value indicating if a subscription supports the concept of sessions.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscription_STATUS">
Namespaces_Topics_Subscription_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.NamespacesTopicsSubscriptionsRule">NamespacesTopicsSubscriptionsRule
</h3>
<div>
<p>Generator information:
- Generated from: /servicebus/resource-manager/Microsoft.ServiceBus/preview/2021-01-01-preview/Rules.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ServiceBus/&#x200b;namespaces/&#x200b;{namespaceName}/&#x200b;topics/&#x200b;{topicName}/&#x200b;subscriptions/&#x200b;{subscriptionName}/&#x200b;rules/&#x200b;{ruleName}</&#x200b;p>
</div>
<table>
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
<a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscriptions_Rule_Spec">
Namespaces_Topics_Subscriptions_Rule_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>action</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.Action">
Action
</a>
</em>
</td>
<td>
<p>Action: Represents the filter actions which are allowed for the transformation of a message that have been matched by a
filter expression.</p>
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
<code>correlationFilter</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.CorrelationFilter">
CorrelationFilter
</a>
</em>
</td>
<td>
<p>CorrelationFilter: Properties of correlationFilter</p>
</td>
</tr>
<tr>
<td>
<code>filterType</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.FilterType">
FilterType
</a>
</em>
</td>
<td>
<p>FilterType: Filter type that is evaluated against a BrokeredMessage.</p>
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
reference to a servicebus.azure.com/NamespacesTopicsSubscription resource</p>
</td>
</tr>
<tr>
<td>
<code>sqlFilter</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.SqlFilter">
SqlFilter
</a>
</em>
</td>
<td>
<p>SqlFilter: Properties of sqlFilter</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscriptions_Rule_STATUS">
Namespaces_Topics_Subscriptions_Rule_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Namespaces_Queue_STATUS">Namespaces_Queue_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.NamespacesQueue">NamespacesQueue</a>)
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
<code>accessedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>AccessedAt: Last time a message was sent, or the last time there was a receive request to this queue.</p>
</td>
</tr>
<tr>
<td>
<code>autoDeleteOnIdle</code><br/>
<em>
string
</em>
</td>
<td>
<p>AutoDeleteOnIdle: ISO 8061 timeSpan idle interval after which the queue is automatically deleted. The minimum duration
is 5 minutes.</p>
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
<code>countDetails</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.MessageCountDetails_STATUS">
MessageCountDetails_STATUS
</a>
</em>
</td>
<td>
<p>CountDetails: Message Count Details.</p>
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
<p>CreatedAt: The exact time the message was created.</p>
</td>
</tr>
<tr>
<td>
<code>deadLetteringOnMessageExpiration</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DeadLetteringOnMessageExpiration: A value that indicates whether this queue has dead letter support when a message
expires.</p>
</td>
</tr>
<tr>
<td>
<code>defaultMessageTimeToLive</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultMessageTimeToLive: ISO 8601 default message timespan to live value. This is the duration after which the message
expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
set on a message itself.</p>
</td>
</tr>
<tr>
<td>
<code>duplicateDetectionHistoryTimeWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>DuplicateDetectionHistoryTimeWindow: ISO 8601 timeSpan structure that defines the duration of the duplicate detection
history. The default value is 10 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>enableBatchedOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableExpress</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableExpress: A value that indicates whether Express Entities are enabled. An express queue holds a message in memory
temporarily before writing it to persistent storage.</p>
</td>
</tr>
<tr>
<td>
<code>enablePartitioning</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePartitioning: A value that indicates whether the queue is to be partitioned across multiple message brokers.</p>
</td>
</tr>
<tr>
<td>
<code>forwardDeadLetteredMessagesTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardDeadLetteredMessagesTo: Queue/Topic name to forward the Dead Letter message</p>
</td>
</tr>
<tr>
<td>
<code>forwardTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardTo: Queue/Topic name to forward the messages</p>
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
<code>lockDuration</code><br/>
<em>
string
</em>
</td>
<td>
<p>LockDuration: ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for
other receivers. The maximum value for LockDuration is 5 minutes; the default value is 1 minute.</p>
</td>
</tr>
<tr>
<td>
<code>maxDeliveryCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxDeliveryCount: The maximum delivery count. A message is automatically deadlettered after this number of deliveries.
default value is 10.</p>
</td>
</tr>
<tr>
<td>
<code>maxSizeInMegabytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxSizeInMegabytes: The maximum size of the queue in megabytes, which is the size of memory allocated for the queue.
Default is 1024.</p>
</td>
</tr>
<tr>
<td>
<code>messageCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MessageCount: The number of messages in the queue.</p>
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
<code>requiresDuplicateDetection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresDuplicateDetection: A value indicating if this queue requires duplicate detection.</p>
</td>
</tr>
<tr>
<td>
<code>requiresSession</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresSession: A value that indicates whether the queue supports the concept of sessions.</p>
</td>
</tr>
<tr>
<td>
<code>sizeInBytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>SizeInBytes: The size of the queue, in bytes.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.EntityStatus_STATUS">
EntityStatus_STATUS
</a>
</em>
</td>
<td>
<p>Status: Enumerates the possible values for the status of a messaging entity.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.SystemData_STATUS">
SystemData_STATUS
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
<p>Type: Resource type</p>
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
<h3 id="servicebus.azure.com/v1api20210101preview.Namespaces_Queue_STATUS_ARM">Namespaces_Queue_STATUS_ARM
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
<p>Id: Resource Id</p>
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
<a href="#servicebus.azure.com/v1api20210101preview.SBQueueProperties_STATUS_ARM">
SBQueueProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: Queue Properties</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.SystemData_STATUS_ARM">
SystemData_STATUS_ARM
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
<p>Type: Resource type</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Namespaces_Queue_Spec">Namespaces_Queue_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.NamespacesQueue">NamespacesQueue</a>)
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
<code>autoDeleteOnIdle</code><br/>
<em>
string
</em>
</td>
<td>
<p>AutoDeleteOnIdle: ISO 8061 timeSpan idle interval after which the queue is automatically deleted. The minimum duration
is 5 minutes.</p>
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
<code>deadLetteringOnMessageExpiration</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DeadLetteringOnMessageExpiration: A value that indicates whether this queue has dead letter support when a message
expires.</p>
</td>
</tr>
<tr>
<td>
<code>defaultMessageTimeToLive</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultMessageTimeToLive: ISO 8601 default message timespan to live value. This is the duration after which the message
expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
set on a message itself.</p>
</td>
</tr>
<tr>
<td>
<code>duplicateDetectionHistoryTimeWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>DuplicateDetectionHistoryTimeWindow: ISO 8601 timeSpan structure that defines the duration of the duplicate detection
history. The default value is 10 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>enableBatchedOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableExpress</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableExpress: A value that indicates whether Express Entities are enabled. An express queue holds a message in memory
temporarily before writing it to persistent storage.</p>
</td>
</tr>
<tr>
<td>
<code>enablePartitioning</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePartitioning: A value that indicates whether the queue is to be partitioned across multiple message brokers.</p>
</td>
</tr>
<tr>
<td>
<code>forwardDeadLetteredMessagesTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardDeadLetteredMessagesTo: Queue/Topic name to forward the Dead Letter message</p>
</td>
</tr>
<tr>
<td>
<code>forwardTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardTo: Queue/Topic name to forward the messages</p>
</td>
</tr>
<tr>
<td>
<code>lockDuration</code><br/>
<em>
string
</em>
</td>
<td>
<p>LockDuration: ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for
other receivers. The maximum value for LockDuration is 5 minutes; the default value is 1 minute.</p>
</td>
</tr>
<tr>
<td>
<code>maxDeliveryCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxDeliveryCount: The maximum delivery count. A message is automatically deadlettered after this number of deliveries.
default value is 10.</p>
</td>
</tr>
<tr>
<td>
<code>maxSizeInMegabytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxSizeInMegabytes: The maximum size of the queue in megabytes, which is the size of memory allocated for the queue.
Default is 1024.</p>
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
reference to a servicebus.azure.com/Namespace resource</p>
</td>
</tr>
<tr>
<td>
<code>requiresDuplicateDetection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresDuplicateDetection: A value indicating if this queue requires duplicate detection.</p>
</td>
</tr>
<tr>
<td>
<code>requiresSession</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresSession: A value that indicates whether the queue supports the concept of sessions.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Namespaces_Queue_Spec_ARM">Namespaces_Queue_Spec_ARM
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
<a href="#servicebus.azure.com/v1api20210101preview.SBQueueProperties_ARM">
SBQueueProperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: Queue Properties</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Namespaces_Topic_STATUS">Namespaces_Topic_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.NamespacesTopic">NamespacesTopic</a>)
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
<code>accessedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>AccessedAt: Last time the message was sent, or a request was received, for this topic.</p>
</td>
</tr>
<tr>
<td>
<code>autoDeleteOnIdle</code><br/>
<em>
string
</em>
</td>
<td>
<p>AutoDeleteOnIdle: ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration
is 5 minutes.</p>
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
<code>countDetails</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.MessageCountDetails_STATUS">
MessageCountDetails_STATUS
</a>
</em>
</td>
<td>
<p>CountDetails: Message count details</p>
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
<code>defaultMessageTimeToLive</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultMessageTimeToLive: ISO 8601 Default message timespan to live value. This is the duration after which the message
expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
set on a message itself.</p>
</td>
</tr>
<tr>
<td>
<code>duplicateDetectionHistoryTimeWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>DuplicateDetectionHistoryTimeWindow: ISO8601 timespan structure that defines the duration of the duplicate detection
history. The default value is 10 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>enableBatchedOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableExpress</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableExpress: Value that indicates whether Express Entities are enabled. An express topic holds a message in memory
temporarily before writing it to persistent storage.</p>
</td>
</tr>
<tr>
<td>
<code>enablePartitioning</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePartitioning: Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.</p>
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
<code>maxSizeInMegabytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxSizeInMegabytes: Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic.
Default is 1024.</p>
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
<code>requiresDuplicateDetection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresDuplicateDetection: Value indicating if this topic requires duplicate detection.</p>
</td>
</tr>
<tr>
<td>
<code>sizeInBytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>SizeInBytes: Size of the topic, in bytes.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.EntityStatus_STATUS">
EntityStatus_STATUS
</a>
</em>
</td>
<td>
<p>Status: Enumerates the possible values for the status of a messaging entity.</p>
</td>
</tr>
<tr>
<td>
<code>subscriptionCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>SubscriptionCount: Number of subscriptions.</p>
</td>
</tr>
<tr>
<td>
<code>supportOrdering</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SupportOrdering: Value that indicates whether the topic supports ordering.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.SystemData_STATUS">
SystemData_STATUS
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
<p>Type: Resource type</p>
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
<h3 id="servicebus.azure.com/v1api20210101preview.Namespaces_Topic_STATUS_ARM">Namespaces_Topic_STATUS_ARM
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
<p>Id: Resource Id</p>
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
<a href="#servicebus.azure.com/v1api20210101preview.SBTopicProperties_STATUS_ARM">
SBTopicProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of topic resource.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.SystemData_STATUS_ARM">
SystemData_STATUS_ARM
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
<p>Type: Resource type</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Namespaces_Topic_Spec">Namespaces_Topic_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.NamespacesTopic">NamespacesTopic</a>)
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
<code>autoDeleteOnIdle</code><br/>
<em>
string
</em>
</td>
<td>
<p>AutoDeleteOnIdle: ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration
is 5 minutes.</p>
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
<code>defaultMessageTimeToLive</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultMessageTimeToLive: ISO 8601 Default message timespan to live value. This is the duration after which the message
expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
set on a message itself.</p>
</td>
</tr>
<tr>
<td>
<code>duplicateDetectionHistoryTimeWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>DuplicateDetectionHistoryTimeWindow: ISO8601 timespan structure that defines the duration of the duplicate detection
history. The default value is 10 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>enableBatchedOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableExpress</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableExpress: Value that indicates whether Express Entities are enabled. An express topic holds a message in memory
temporarily before writing it to persistent storage.</p>
</td>
</tr>
<tr>
<td>
<code>enablePartitioning</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePartitioning: Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>maxSizeInMegabytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxSizeInMegabytes: Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic.
Default is 1024.</p>
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
reference to a servicebus.azure.com/Namespace resource</p>
</td>
</tr>
<tr>
<td>
<code>requiresDuplicateDetection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresDuplicateDetection: Value indicating if this topic requires duplicate detection.</p>
</td>
</tr>
<tr>
<td>
<code>supportOrdering</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SupportOrdering: Value that indicates whether the topic supports ordering.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Namespaces_Topic_Spec_ARM">Namespaces_Topic_Spec_ARM
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
<a href="#servicebus.azure.com/v1api20210101preview.SBTopicProperties_ARM">
SBTopicProperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of topic resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscription_STATUS">Namespaces_Topics_Subscription_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.NamespacesTopicsSubscription">NamespacesTopicsSubscription</a>)
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
<code>accessedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>AccessedAt: Last time there was a receive request to this subscription.</p>
</td>
</tr>
<tr>
<td>
<code>autoDeleteOnIdle</code><br/>
<em>
string
</em>
</td>
<td>
<p>AutoDeleteOnIdle: ISO 8061 timeSpan idle interval after which the topic is automatically deleted. The minimum duration
is 5 minutes.</p>
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
<code>countDetails</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.MessageCountDetails_STATUS">
MessageCountDetails_STATUS
</a>
</em>
</td>
<td>
<p>CountDetails: Message count details</p>
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
<code>deadLetteringOnFilterEvaluationExceptions</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DeadLetteringOnFilterEvaluationExceptions: Value that indicates whether a subscription has dead letter support on filter
evaluation exceptions.</p>
</td>
</tr>
<tr>
<td>
<code>deadLetteringOnMessageExpiration</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DeadLetteringOnMessageExpiration: Value that indicates whether a subscription has dead letter support when a message
expires.</p>
</td>
</tr>
<tr>
<td>
<code>defaultMessageTimeToLive</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultMessageTimeToLive: ISO 8061 Default message timespan to live value. This is the duration after which the message
expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
set on a message itself.</p>
</td>
</tr>
<tr>
<td>
<code>duplicateDetectionHistoryTimeWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>DuplicateDetectionHistoryTimeWindow: ISO 8601 timeSpan structure that defines the duration of the duplicate detection
history. The default value is 10 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>enableBatchedOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.</p>
</td>
</tr>
<tr>
<td>
<code>forwardDeadLetteredMessagesTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardDeadLetteredMessagesTo: Queue/Topic name to forward the Dead Letter message</p>
</td>
</tr>
<tr>
<td>
<code>forwardTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardTo: Queue/Topic name to forward the messages</p>
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
<code>lockDuration</code><br/>
<em>
string
</em>
</td>
<td>
<p>LockDuration: ISO 8061 lock duration timespan for the subscription. The default value is 1 minute.</p>
</td>
</tr>
<tr>
<td>
<code>maxDeliveryCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxDeliveryCount: Number of maximum deliveries.</p>
</td>
</tr>
<tr>
<td>
<code>messageCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MessageCount: Number of messages.</p>
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
<code>requiresSession</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresSession: Value indicating if a subscription supports the concept of sessions.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.EntityStatus_STATUS">
EntityStatus_STATUS
</a>
</em>
</td>
<td>
<p>Status: Enumerates the possible values for the status of a messaging entity.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.SystemData_STATUS">
SystemData_STATUS
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
<p>Type: Resource type</p>
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
<h3 id="servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscription_STATUS_ARM">Namespaces_Topics_Subscription_STATUS_ARM
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
<p>Id: Resource Id</p>
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
<a href="#servicebus.azure.com/v1api20210101preview.SBSubscriptionProperties_STATUS_ARM">
SBSubscriptionProperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of subscriptions resource.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.SystemData_STATUS_ARM">
SystemData_STATUS_ARM
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
<p>Type: Resource type</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscription_Spec">Namespaces_Topics_Subscription_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.NamespacesTopicsSubscription">NamespacesTopicsSubscription</a>)
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
<code>autoDeleteOnIdle</code><br/>
<em>
string
</em>
</td>
<td>
<p>AutoDeleteOnIdle: ISO 8061 timeSpan idle interval after which the topic is automatically deleted. The minimum duration
is 5 minutes.</p>
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
<code>deadLetteringOnFilterEvaluationExceptions</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DeadLetteringOnFilterEvaluationExceptions: Value that indicates whether a subscription has dead letter support on filter
evaluation exceptions.</p>
</td>
</tr>
<tr>
<td>
<code>deadLetteringOnMessageExpiration</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DeadLetteringOnMessageExpiration: Value that indicates whether a subscription has dead letter support when a message
expires.</p>
</td>
</tr>
<tr>
<td>
<code>defaultMessageTimeToLive</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultMessageTimeToLive: ISO 8061 Default message timespan to live value. This is the duration after which the message
expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
set on a message itself.</p>
</td>
</tr>
<tr>
<td>
<code>duplicateDetectionHistoryTimeWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>DuplicateDetectionHistoryTimeWindow: ISO 8601 timeSpan structure that defines the duration of the duplicate detection
history. The default value is 10 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>enableBatchedOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.</p>
</td>
</tr>
<tr>
<td>
<code>forwardDeadLetteredMessagesTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardDeadLetteredMessagesTo: Queue/Topic name to forward the Dead Letter message</p>
</td>
</tr>
<tr>
<td>
<code>forwardTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardTo: Queue/Topic name to forward the messages</p>
</td>
</tr>
<tr>
<td>
<code>lockDuration</code><br/>
<em>
string
</em>
</td>
<td>
<p>LockDuration: ISO 8061 lock duration timespan for the subscription. The default value is 1 minute.</p>
</td>
</tr>
<tr>
<td>
<code>maxDeliveryCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxDeliveryCount: Number of maximum deliveries.</p>
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
reference to a servicebus.azure.com/NamespacesTopic resource</p>
</td>
</tr>
<tr>
<td>
<code>requiresSession</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresSession: Value indicating if a subscription supports the concept of sessions.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscription_Spec_ARM">Namespaces_Topics_Subscription_Spec_ARM
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
<a href="#servicebus.azure.com/v1api20210101preview.SBSubscriptionProperties_ARM">
SBSubscriptionProperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of subscriptions resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscriptions_Rule_STATUS">Namespaces_Topics_Subscriptions_Rule_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.NamespacesTopicsSubscriptionsRule">NamespacesTopicsSubscriptionsRule</a>)
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
<code>action</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.Action_STATUS">
Action_STATUS
</a>
</em>
</td>
<td>
<p>Action: Represents the filter actions which are allowed for the transformation of a message that have been matched by a
filter expression.</p>
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
<code>correlationFilter</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.CorrelationFilter_STATUS">
CorrelationFilter_STATUS
</a>
</em>
</td>
<td>
<p>CorrelationFilter: Properties of correlationFilter</p>
</td>
</tr>
<tr>
<td>
<code>filterType</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.FilterType_STATUS">
FilterType_STATUS
</a>
</em>
</td>
<td>
<p>FilterType: Filter type that is evaluated against a BrokeredMessage.</p>
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
<code>sqlFilter</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.SqlFilter_STATUS">
SqlFilter_STATUS
</a>
</em>
</td>
<td>
<p>SqlFilter: Properties of sqlFilter</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.SystemData_STATUS">
SystemData_STATUS
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
<p>Type: Resource type</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscriptions_Rule_STATUS_ARM">Namespaces_Topics_Subscriptions_Rule_STATUS_ARM
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
<p>Id: Resource Id</p>
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
<a href="#servicebus.azure.com/v1api20210101preview.Ruleproperties_STATUS_ARM">
Ruleproperties_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of Rule resource</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.SystemData_STATUS_ARM">
SystemData_STATUS_ARM
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
<p>Type: Resource type</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscriptions_Rule_Spec">Namespaces_Topics_Subscriptions_Rule_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.NamespacesTopicsSubscriptionsRule">NamespacesTopicsSubscriptionsRule</a>)
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
<code>action</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.Action">
Action
</a>
</em>
</td>
<td>
<p>Action: Represents the filter actions which are allowed for the transformation of a message that have been matched by a
filter expression.</p>
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
<code>correlationFilter</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.CorrelationFilter">
CorrelationFilter
</a>
</em>
</td>
<td>
<p>CorrelationFilter: Properties of correlationFilter</p>
</td>
</tr>
<tr>
<td>
<code>filterType</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.FilterType">
FilterType
</a>
</em>
</td>
<td>
<p>FilterType: Filter type that is evaluated against a BrokeredMessage.</p>
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
reference to a servicebus.azure.com/NamespacesTopicsSubscription resource</p>
</td>
</tr>
<tr>
<td>
<code>sqlFilter</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.SqlFilter">
SqlFilter
</a>
</em>
</td>
<td>
<p>SqlFilter: Properties of sqlFilter</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscriptions_Rule_Spec_ARM">Namespaces_Topics_Subscriptions_Rule_Spec_ARM
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
<a href="#servicebus.azure.com/v1api20210101preview.Ruleproperties_ARM">
Ruleproperties_ARM
</a>
</em>
</td>
<td>
<p>Properties: Properties of Rule resource</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.PrivateEndpointConnection_STATUS">PrivateEndpointConnection_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespace_STATUS">Namespace_STATUS</a>)
</p>
<div>
<p>Properties of the PrivateEndpointConnection.</p>
</div>
<table>
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
<p>Id: Resource Id</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.PrivateEndpointConnection_STATUS_ARM">PrivateEndpointConnection_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.SBNamespaceProperties_STATUS_ARM">SBNamespaceProperties_STATUS_ARM</a>)
</p>
<div>
<p>Properties of the PrivateEndpointConnection.</p>
</div>
<table>
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
<p>Id: Resource Id</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Ruleproperties_ARM">Ruleproperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscriptions_Rule_Spec_ARM">Namespaces_Topics_Subscriptions_Rule_Spec_ARM</a>)
</p>
<div>
<p>Description of Rule Resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>action</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.Action_ARM">
Action_ARM
</a>
</em>
</td>
<td>
<p>Action: Represents the filter actions which are allowed for the transformation of a message that have been matched by a
filter expression.</p>
</td>
</tr>
<tr>
<td>
<code>correlationFilter</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.CorrelationFilter_ARM">
CorrelationFilter_ARM
</a>
</em>
</td>
<td>
<p>CorrelationFilter: Properties of correlationFilter</p>
</td>
</tr>
<tr>
<td>
<code>filterType</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.FilterType">
FilterType
</a>
</em>
</td>
<td>
<p>FilterType: Filter type that is evaluated against a BrokeredMessage.</p>
</td>
</tr>
<tr>
<td>
<code>sqlFilter</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.SqlFilter_ARM">
SqlFilter_ARM
</a>
</em>
</td>
<td>
<p>SqlFilter: Properties of sqlFilter</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.Ruleproperties_STATUS_ARM">Ruleproperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscriptions_Rule_STATUS_ARM">Namespaces_Topics_Subscriptions_Rule_STATUS_ARM</a>)
</p>
<div>
<p>Description of Rule Resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>action</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.Action_STATUS_ARM">
Action_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Action: Represents the filter actions which are allowed for the transformation of a message that have been matched by a
filter expression.</p>
</td>
</tr>
<tr>
<td>
<code>correlationFilter</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.CorrelationFilter_STATUS_ARM">
CorrelationFilter_STATUS_ARM
</a>
</em>
</td>
<td>
<p>CorrelationFilter: Properties of correlationFilter</p>
</td>
</tr>
<tr>
<td>
<code>filterType</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.FilterType_STATUS">
FilterType_STATUS
</a>
</em>
</td>
<td>
<p>FilterType: Filter type that is evaluated against a BrokeredMessage.</p>
</td>
</tr>
<tr>
<td>
<code>sqlFilter</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.SqlFilter_STATUS_ARM">
SqlFilter_STATUS_ARM
</a>
</em>
</td>
<td>
<p>SqlFilter: Properties of sqlFilter</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.SBNamespaceProperties_ARM">SBNamespaceProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespace_Spec_ARM">Namespace_Spec_ARM</a>)
</p>
<div>
<p>Properties of the namespace.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.Encryption_ARM">
Encryption_ARM
</a>
</em>
</td>
<td>
<p>Encryption: Properties of BYOK Encryption description</p>
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
<p>ZoneRedundant: Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.SBNamespaceProperties_STATUS_ARM">SBNamespaceProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespace_STATUS_ARM">Namespace_STATUS_ARM</a>)
</p>
<div>
<p>Properties of the namespace.</p>
</div>
<table>
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
<p>CreatedAt: The time the namespace was created</p>
</td>
</tr>
<tr>
<td>
<code>encryption</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.Encryption_STATUS_ARM">
Encryption_STATUS_ARM
</a>
</em>
</td>
<td>
<p>Encryption: Properties of BYOK Encryption description</p>
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
<p>MetricId: Identifier for Azure Insights metrics</p>
</td>
</tr>
<tr>
<td>
<code>privateEndpointConnections</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.PrivateEndpointConnection_STATUS_ARM">
[]PrivateEndpointConnection_STATUS_ARM
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
<p>ProvisioningState: Provisioning state of the namespace.</p>
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
<p>Status: Status of the namespace.</p>
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
<p>UpdatedAt: The time the namespace was updated.</p>
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
<p>ZoneRedundant: Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.SBQueueProperties_ARM">SBQueueProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Queue_Spec_ARM">Namespaces_Queue_Spec_ARM</a>)
</p>
<div>
<p>The Queue Properties definition.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>autoDeleteOnIdle</code><br/>
<em>
string
</em>
</td>
<td>
<p>AutoDeleteOnIdle: ISO 8061 timeSpan idle interval after which the queue is automatically deleted. The minimum duration
is 5 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>deadLetteringOnMessageExpiration</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DeadLetteringOnMessageExpiration: A value that indicates whether this queue has dead letter support when a message
expires.</p>
</td>
</tr>
<tr>
<td>
<code>defaultMessageTimeToLive</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultMessageTimeToLive: ISO 8601 default message timespan to live value. This is the duration after which the message
expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
set on a message itself.</p>
</td>
</tr>
<tr>
<td>
<code>duplicateDetectionHistoryTimeWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>DuplicateDetectionHistoryTimeWindow: ISO 8601 timeSpan structure that defines the duration of the duplicate detection
history. The default value is 10 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>enableBatchedOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableExpress</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableExpress: A value that indicates whether Express Entities are enabled. An express queue holds a message in memory
temporarily before writing it to persistent storage.</p>
</td>
</tr>
<tr>
<td>
<code>enablePartitioning</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePartitioning: A value that indicates whether the queue is to be partitioned across multiple message brokers.</p>
</td>
</tr>
<tr>
<td>
<code>forwardDeadLetteredMessagesTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardDeadLetteredMessagesTo: Queue/Topic name to forward the Dead Letter message</p>
</td>
</tr>
<tr>
<td>
<code>forwardTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardTo: Queue/Topic name to forward the messages</p>
</td>
</tr>
<tr>
<td>
<code>lockDuration</code><br/>
<em>
string
</em>
</td>
<td>
<p>LockDuration: ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for
other receivers. The maximum value for LockDuration is 5 minutes; the default value is 1 minute.</p>
</td>
</tr>
<tr>
<td>
<code>maxDeliveryCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxDeliveryCount: The maximum delivery count. A message is automatically deadlettered after this number of deliveries.
default value is 10.</p>
</td>
</tr>
<tr>
<td>
<code>maxSizeInMegabytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxSizeInMegabytes: The maximum size of the queue in megabytes, which is the size of memory allocated for the queue.
Default is 1024.</p>
</td>
</tr>
<tr>
<td>
<code>requiresDuplicateDetection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresDuplicateDetection: A value indicating if this queue requires duplicate detection.</p>
</td>
</tr>
<tr>
<td>
<code>requiresSession</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresSession: A value that indicates whether the queue supports the concept of sessions.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.SBQueueProperties_STATUS_ARM">SBQueueProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Queue_STATUS_ARM">Namespaces_Queue_STATUS_ARM</a>)
</p>
<div>
<p>The Queue Properties definition.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>accessedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>AccessedAt: Last time a message was sent, or the last time there was a receive request to this queue.</p>
</td>
</tr>
<tr>
<td>
<code>autoDeleteOnIdle</code><br/>
<em>
string
</em>
</td>
<td>
<p>AutoDeleteOnIdle: ISO 8061 timeSpan idle interval after which the queue is automatically deleted. The minimum duration
is 5 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>countDetails</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.MessageCountDetails_STATUS_ARM">
MessageCountDetails_STATUS_ARM
</a>
</em>
</td>
<td>
<p>CountDetails: Message Count Details.</p>
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
<p>CreatedAt: The exact time the message was created.</p>
</td>
</tr>
<tr>
<td>
<code>deadLetteringOnMessageExpiration</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DeadLetteringOnMessageExpiration: A value that indicates whether this queue has dead letter support when a message
expires.</p>
</td>
</tr>
<tr>
<td>
<code>defaultMessageTimeToLive</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultMessageTimeToLive: ISO 8601 default message timespan to live value. This is the duration after which the message
expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
set on a message itself.</p>
</td>
</tr>
<tr>
<td>
<code>duplicateDetectionHistoryTimeWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>DuplicateDetectionHistoryTimeWindow: ISO 8601 timeSpan structure that defines the duration of the duplicate detection
history. The default value is 10 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>enableBatchedOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableExpress</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableExpress: A value that indicates whether Express Entities are enabled. An express queue holds a message in memory
temporarily before writing it to persistent storage.</p>
</td>
</tr>
<tr>
<td>
<code>enablePartitioning</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePartitioning: A value that indicates whether the queue is to be partitioned across multiple message brokers.</p>
</td>
</tr>
<tr>
<td>
<code>forwardDeadLetteredMessagesTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardDeadLetteredMessagesTo: Queue/Topic name to forward the Dead Letter message</p>
</td>
</tr>
<tr>
<td>
<code>forwardTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardTo: Queue/Topic name to forward the messages</p>
</td>
</tr>
<tr>
<td>
<code>lockDuration</code><br/>
<em>
string
</em>
</td>
<td>
<p>LockDuration: ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for
other receivers. The maximum value for LockDuration is 5 minutes; the default value is 1 minute.</p>
</td>
</tr>
<tr>
<td>
<code>maxDeliveryCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxDeliveryCount: The maximum delivery count. A message is automatically deadlettered after this number of deliveries.
default value is 10.</p>
</td>
</tr>
<tr>
<td>
<code>maxSizeInMegabytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxSizeInMegabytes: The maximum size of the queue in megabytes, which is the size of memory allocated for the queue.
Default is 1024.</p>
</td>
</tr>
<tr>
<td>
<code>messageCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MessageCount: The number of messages in the queue.</p>
</td>
</tr>
<tr>
<td>
<code>requiresDuplicateDetection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresDuplicateDetection: A value indicating if this queue requires duplicate detection.</p>
</td>
</tr>
<tr>
<td>
<code>requiresSession</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresSession: A value that indicates whether the queue supports the concept of sessions.</p>
</td>
</tr>
<tr>
<td>
<code>sizeInBytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>SizeInBytes: The size of the queue, in bytes.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.EntityStatus_STATUS">
EntityStatus_STATUS
</a>
</em>
</td>
<td>
<p>Status: Enumerates the possible values for the status of a messaging entity.</p>
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
<h3 id="servicebus.azure.com/v1api20210101preview.SBSku">SBSku
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespace_Spec">Namespace_Spec</a>)
</p>
<div>
<p>SKU of the namespace.</p>
</div>
<table>
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
<p>Capacity: The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.SBSku_Name">
SBSku_Name
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
<a href="#servicebus.azure.com/v1api20210101preview.SBSku_Tier">
SBSku_Tier
</a>
</em>
</td>
<td>
<p>Tier: The billing tier of this particular SKU.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.SBSku_ARM">SBSku_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespace_Spec_ARM">Namespace_Spec_ARM</a>)
</p>
<div>
<p>SKU of the namespace.</p>
</div>
<table>
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
<p>Capacity: The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.SBSku_Name">
SBSku_Name
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
<a href="#servicebus.azure.com/v1api20210101preview.SBSku_Tier">
SBSku_Tier
</a>
</em>
</td>
<td>
<p>Tier: The billing tier of this particular SKU.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.SBSku_Name">SBSku_Name
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.SBSku">SBSku</a>, <a href="#servicebus.azure.com/v1api20210101preview.SBSku_ARM">SBSku_ARM</a>)
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
<h3 id="servicebus.azure.com/v1api20210101preview.SBSku_Name_STATUS">SBSku_Name_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.SBSku_STATUS">SBSku_STATUS</a>, <a href="#servicebus.azure.com/v1api20210101preview.SBSku_STATUS_ARM">SBSku_STATUS_ARM</a>)
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
<h3 id="servicebus.azure.com/v1api20210101preview.SBSku_STATUS">SBSku_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespace_STATUS">Namespace_STATUS</a>)
</p>
<div>
<p>SKU of the namespace.</p>
</div>
<table>
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
<p>Capacity: The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.SBSku_Name_STATUS">
SBSku_Name_STATUS
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
<a href="#servicebus.azure.com/v1api20210101preview.SBSku_Tier_STATUS">
SBSku_Tier_STATUS
</a>
</em>
</td>
<td>
<p>Tier: The billing tier of this particular SKU.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.SBSku_STATUS_ARM">SBSku_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespace_STATUS_ARM">Namespace_STATUS_ARM</a>)
</p>
<div>
<p>SKU of the namespace.</p>
</div>
<table>
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
<p>Capacity: The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.SBSku_Name_STATUS">
SBSku_Name_STATUS
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
<a href="#servicebus.azure.com/v1api20210101preview.SBSku_Tier_STATUS">
SBSku_Tier_STATUS
</a>
</em>
</td>
<td>
<p>Tier: The billing tier of this particular SKU.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.SBSku_Tier">SBSku_Tier
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.SBSku">SBSku</a>, <a href="#servicebus.azure.com/v1api20210101preview.SBSku_ARM">SBSku_ARM</a>)
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
<h3 id="servicebus.azure.com/v1api20210101preview.SBSku_Tier_STATUS">SBSku_Tier_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.SBSku_STATUS">SBSku_STATUS</a>, <a href="#servicebus.azure.com/v1api20210101preview.SBSku_STATUS_ARM">SBSku_STATUS_ARM</a>)
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
<h3 id="servicebus.azure.com/v1api20210101preview.SBSubscriptionProperties_ARM">SBSubscriptionProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscription_Spec_ARM">Namespaces_Topics_Subscription_Spec_ARM</a>)
</p>
<div>
<p>Description of Subscription Resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>autoDeleteOnIdle</code><br/>
<em>
string
</em>
</td>
<td>
<p>AutoDeleteOnIdle: ISO 8061 timeSpan idle interval after which the topic is automatically deleted. The minimum duration
is 5 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>deadLetteringOnFilterEvaluationExceptions</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DeadLetteringOnFilterEvaluationExceptions: Value that indicates whether a subscription has dead letter support on filter
evaluation exceptions.</p>
</td>
</tr>
<tr>
<td>
<code>deadLetteringOnMessageExpiration</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DeadLetteringOnMessageExpiration: Value that indicates whether a subscription has dead letter support when a message
expires.</p>
</td>
</tr>
<tr>
<td>
<code>defaultMessageTimeToLive</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultMessageTimeToLive: ISO 8061 Default message timespan to live value. This is the duration after which the message
expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
set on a message itself.</p>
</td>
</tr>
<tr>
<td>
<code>duplicateDetectionHistoryTimeWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>DuplicateDetectionHistoryTimeWindow: ISO 8601 timeSpan structure that defines the duration of the duplicate detection
history. The default value is 10 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>enableBatchedOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.</p>
</td>
</tr>
<tr>
<td>
<code>forwardDeadLetteredMessagesTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardDeadLetteredMessagesTo: Queue/Topic name to forward the Dead Letter message</p>
</td>
</tr>
<tr>
<td>
<code>forwardTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardTo: Queue/Topic name to forward the messages</p>
</td>
</tr>
<tr>
<td>
<code>lockDuration</code><br/>
<em>
string
</em>
</td>
<td>
<p>LockDuration: ISO 8061 lock duration timespan for the subscription. The default value is 1 minute.</p>
</td>
</tr>
<tr>
<td>
<code>maxDeliveryCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxDeliveryCount: Number of maximum deliveries.</p>
</td>
</tr>
<tr>
<td>
<code>requiresSession</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresSession: Value indicating if a subscription supports the concept of sessions.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.SBSubscriptionProperties_STATUS_ARM">SBSubscriptionProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscription_STATUS_ARM">Namespaces_Topics_Subscription_STATUS_ARM</a>)
</p>
<div>
<p>Description of Subscription Resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>accessedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>AccessedAt: Last time there was a receive request to this subscription.</p>
</td>
</tr>
<tr>
<td>
<code>autoDeleteOnIdle</code><br/>
<em>
string
</em>
</td>
<td>
<p>AutoDeleteOnIdle: ISO 8061 timeSpan idle interval after which the topic is automatically deleted. The minimum duration
is 5 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>countDetails</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.MessageCountDetails_STATUS_ARM">
MessageCountDetails_STATUS_ARM
</a>
</em>
</td>
<td>
<p>CountDetails: Message count details</p>
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
<code>deadLetteringOnFilterEvaluationExceptions</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DeadLetteringOnFilterEvaluationExceptions: Value that indicates whether a subscription has dead letter support on filter
evaluation exceptions.</p>
</td>
</tr>
<tr>
<td>
<code>deadLetteringOnMessageExpiration</code><br/>
<em>
bool
</em>
</td>
<td>
<p>DeadLetteringOnMessageExpiration: Value that indicates whether a subscription has dead letter support when a message
expires.</p>
</td>
</tr>
<tr>
<td>
<code>defaultMessageTimeToLive</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultMessageTimeToLive: ISO 8061 Default message timespan to live value. This is the duration after which the message
expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
set on a message itself.</p>
</td>
</tr>
<tr>
<td>
<code>duplicateDetectionHistoryTimeWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>DuplicateDetectionHistoryTimeWindow: ISO 8601 timeSpan structure that defines the duration of the duplicate detection
history. The default value is 10 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>enableBatchedOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.</p>
</td>
</tr>
<tr>
<td>
<code>forwardDeadLetteredMessagesTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardDeadLetteredMessagesTo: Queue/Topic name to forward the Dead Letter message</p>
</td>
</tr>
<tr>
<td>
<code>forwardTo</code><br/>
<em>
string
</em>
</td>
<td>
<p>ForwardTo: Queue/Topic name to forward the messages</p>
</td>
</tr>
<tr>
<td>
<code>lockDuration</code><br/>
<em>
string
</em>
</td>
<td>
<p>LockDuration: ISO 8061 lock duration timespan for the subscription. The default value is 1 minute.</p>
</td>
</tr>
<tr>
<td>
<code>maxDeliveryCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxDeliveryCount: Number of maximum deliveries.</p>
</td>
</tr>
<tr>
<td>
<code>messageCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>MessageCount: Number of messages.</p>
</td>
</tr>
<tr>
<td>
<code>requiresSession</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresSession: Value indicating if a subscription supports the concept of sessions.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.EntityStatus_STATUS">
EntityStatus_STATUS
</a>
</em>
</td>
<td>
<p>Status: Enumerates the possible values for the status of a messaging entity.</p>
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
<h3 id="servicebus.azure.com/v1api20210101preview.SBTopicProperties_ARM">SBTopicProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topic_Spec_ARM">Namespaces_Topic_Spec_ARM</a>)
</p>
<div>
<p>The Topic Properties definition.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>autoDeleteOnIdle</code><br/>
<em>
string
</em>
</td>
<td>
<p>AutoDeleteOnIdle: ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration
is 5 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>defaultMessageTimeToLive</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultMessageTimeToLive: ISO 8601 Default message timespan to live value. This is the duration after which the message
expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
set on a message itself.</p>
</td>
</tr>
<tr>
<td>
<code>duplicateDetectionHistoryTimeWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>DuplicateDetectionHistoryTimeWindow: ISO8601 timespan structure that defines the duration of the duplicate detection
history. The default value is 10 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>enableBatchedOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableExpress</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableExpress: Value that indicates whether Express Entities are enabled. An express topic holds a message in memory
temporarily before writing it to persistent storage.</p>
</td>
</tr>
<tr>
<td>
<code>enablePartitioning</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePartitioning: Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>maxSizeInMegabytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxSizeInMegabytes: Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic.
Default is 1024.</p>
</td>
</tr>
<tr>
<td>
<code>requiresDuplicateDetection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresDuplicateDetection: Value indicating if this topic requires duplicate detection.</p>
</td>
</tr>
<tr>
<td>
<code>supportOrdering</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SupportOrdering: Value that indicates whether the topic supports ordering.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.SBTopicProperties_STATUS_ARM">SBTopicProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topic_STATUS_ARM">Namespaces_Topic_STATUS_ARM</a>)
</p>
<div>
<p>The Topic Properties definition.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>accessedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>AccessedAt: Last time the message was sent, or a request was received, for this topic.</p>
</td>
</tr>
<tr>
<td>
<code>autoDeleteOnIdle</code><br/>
<em>
string
</em>
</td>
<td>
<p>AutoDeleteOnIdle: ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration
is 5 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>countDetails</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.MessageCountDetails_STATUS_ARM">
MessageCountDetails_STATUS_ARM
</a>
</em>
</td>
<td>
<p>CountDetails: Message count details</p>
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
<code>defaultMessageTimeToLive</code><br/>
<em>
string
</em>
</td>
<td>
<p>DefaultMessageTimeToLive: ISO 8601 Default message timespan to live value. This is the duration after which the message
expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
set on a message itself.</p>
</td>
</tr>
<tr>
<td>
<code>duplicateDetectionHistoryTimeWindow</code><br/>
<em>
string
</em>
</td>
<td>
<p>DuplicateDetectionHistoryTimeWindow: ISO8601 timespan structure that defines the duration of the duplicate detection
history. The default value is 10 minutes.</p>
</td>
</tr>
<tr>
<td>
<code>enableBatchedOperations</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.</p>
</td>
</tr>
<tr>
<td>
<code>enableExpress</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnableExpress: Value that indicates whether Express Entities are enabled. An express topic holds a message in memory
temporarily before writing it to persistent storage.</p>
</td>
</tr>
<tr>
<td>
<code>enablePartitioning</code><br/>
<em>
bool
</em>
</td>
<td>
<p>EnablePartitioning: Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.</p>
</td>
</tr>
<tr>
<td>
<code>maxSizeInMegabytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>MaxSizeInMegabytes: Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic.
Default is 1024.</p>
</td>
</tr>
<tr>
<td>
<code>requiresDuplicateDetection</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresDuplicateDetection: Value indicating if this topic requires duplicate detection.</p>
</td>
</tr>
<tr>
<td>
<code>sizeInBytes</code><br/>
<em>
int
</em>
</td>
<td>
<p>SizeInBytes: Size of the topic, in bytes.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#servicebus.azure.com/v1api20210101preview.EntityStatus_STATUS">
EntityStatus_STATUS
</a>
</em>
</td>
<td>
<p>Status: Enumerates the possible values for the status of a messaging entity.</p>
</td>
</tr>
<tr>
<td>
<code>subscriptionCount</code><br/>
<em>
int
</em>
</td>
<td>
<p>SubscriptionCount: Number of subscriptions.</p>
</td>
</tr>
<tr>
<td>
<code>supportOrdering</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SupportOrdering: Value that indicates whether the topic supports ordering.</p>
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
<h3 id="servicebus.azure.com/v1api20210101preview.SqlFilter">SqlFilter
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscriptions_Rule_Spec">Namespaces_Topics_Subscriptions_Rule_Spec</a>)
</p>
<div>
<p>Represents a filter which is a composition of an expression and an action that is executed in the pub/sub pipeline.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>compatibilityLevel</code><br/>
<em>
int
</em>
</td>
<td>
<p>CompatibilityLevel: This property is reserved for future use. An integer value showing the compatibility level,
currently hard-coded to 20.</p>
</td>
</tr>
<tr>
<td>
<code>requiresPreprocessing</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresPreprocessing: Value that indicates whether the rule action requires preprocessing.</p>
</td>
</tr>
<tr>
<td>
<code>sqlExpression</code><br/>
<em>
string
</em>
</td>
<td>
<p>SqlExpression: The SQL expression. e.g. MyProperty=&lsquo;ABC&rsquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.SqlFilter_ARM">SqlFilter_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Ruleproperties_ARM">Ruleproperties_ARM</a>)
</p>
<div>
<p>Represents a filter which is a composition of an expression and an action that is executed in the pub/sub pipeline.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>compatibilityLevel</code><br/>
<em>
int
</em>
</td>
<td>
<p>CompatibilityLevel: This property is reserved for future use. An integer value showing the compatibility level,
currently hard-coded to 20.</p>
</td>
</tr>
<tr>
<td>
<code>requiresPreprocessing</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresPreprocessing: Value that indicates whether the rule action requires preprocessing.</p>
</td>
</tr>
<tr>
<td>
<code>sqlExpression</code><br/>
<em>
string
</em>
</td>
<td>
<p>SqlExpression: The SQL expression. e.g. MyProperty=&lsquo;ABC&rsquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.SqlFilter_STATUS">SqlFilter_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscriptions_Rule_STATUS">Namespaces_Topics_Subscriptions_Rule_STATUS</a>)
</p>
<div>
<p>Represents a filter which is a composition of an expression and an action that is executed in the pub/sub pipeline.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>compatibilityLevel</code><br/>
<em>
int
</em>
</td>
<td>
<p>CompatibilityLevel: This property is reserved for future use. An integer value showing the compatibility level,
currently hard-coded to 20.</p>
</td>
</tr>
<tr>
<td>
<code>requiresPreprocessing</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresPreprocessing: Value that indicates whether the rule action requires preprocessing.</p>
</td>
</tr>
<tr>
<td>
<code>sqlExpression</code><br/>
<em>
string
</em>
</td>
<td>
<p>SqlExpression: The SQL expression. e.g. MyProperty=&lsquo;ABC&rsquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.SqlFilter_STATUS_ARM">SqlFilter_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Ruleproperties_STATUS_ARM">Ruleproperties_STATUS_ARM</a>)
</p>
<div>
<p>Represents a filter which is a composition of an expression and an action that is executed in the pub/sub pipeline.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>compatibilityLevel</code><br/>
<em>
int
</em>
</td>
<td>
<p>CompatibilityLevel: This property is reserved for future use. An integer value showing the compatibility level,
currently hard-coded to 20.</p>
</td>
</tr>
<tr>
<td>
<code>requiresPreprocessing</code><br/>
<em>
bool
</em>
</td>
<td>
<p>RequiresPreprocessing: Value that indicates whether the rule action requires preprocessing.</p>
</td>
</tr>
<tr>
<td>
<code>sqlExpression</code><br/>
<em>
string
</em>
</td>
<td>
<p>SqlExpression: The SQL expression. e.g. MyProperty=&lsquo;ABC&rsquo;</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.SystemData_CreatedByType_STATUS">SystemData_CreatedByType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.SystemData_STATUS">SystemData_STATUS</a>, <a href="#servicebus.azure.com/v1api20210101preview.SystemData_STATUS_ARM">SystemData_STATUS_ARM</a>)
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
<h3 id="servicebus.azure.com/v1api20210101preview.SystemData_LastModifiedByType_STATUS">SystemData_LastModifiedByType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.SystemData_STATUS">SystemData_STATUS</a>, <a href="#servicebus.azure.com/v1api20210101preview.SystemData_STATUS_ARM">SystemData_STATUS_ARM</a>)
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
<h3 id="servicebus.azure.com/v1api20210101preview.SystemData_STATUS">SystemData_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespace_STATUS">Namespace_STATUS</a>, <a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Queue_STATUS">Namespaces_Queue_STATUS</a>, <a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topic_STATUS">Namespaces_Topic_STATUS</a>, <a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscription_STATUS">Namespaces_Topics_Subscription_STATUS</a>, <a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscriptions_Rule_STATUS">Namespaces_Topics_Subscriptions_Rule_STATUS</a>)
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
<a href="#servicebus.azure.com/v1api20210101preview.SystemData_CreatedByType_STATUS">
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
<a href="#servicebus.azure.com/v1api20210101preview.SystemData_LastModifiedByType_STATUS">
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
<h3 id="servicebus.azure.com/v1api20210101preview.SystemData_STATUS_ARM">SystemData_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Namespace_STATUS_ARM">Namespace_STATUS_ARM</a>, <a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Queue_STATUS_ARM">Namespaces_Queue_STATUS_ARM</a>, <a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topic_STATUS_ARM">Namespaces_Topic_STATUS_ARM</a>, <a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscription_STATUS_ARM">Namespaces_Topics_Subscription_STATUS_ARM</a>, <a href="#servicebus.azure.com/v1api20210101preview.Namespaces_Topics_Subscriptions_Rule_STATUS_ARM">Namespaces_Topics_Subscriptions_Rule_STATUS_ARM</a>)
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
<a href="#servicebus.azure.com/v1api20210101preview.SystemData_CreatedByType_STATUS">
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
<a href="#servicebus.azure.com/v1api20210101preview.SystemData_LastModifiedByType_STATUS">
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
<h3 id="servicebus.azure.com/v1api20210101preview.UserAssignedIdentityDetails">UserAssignedIdentityDetails
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Identity">Identity</a>)
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
<h3 id="servicebus.azure.com/v1api20210101preview.UserAssignedIdentityDetails_ARM">UserAssignedIdentityDetails_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.Identity_ARM">Identity_ARM</a>)
</p>
<div>
<p>Information about the user assigned identity for the resource</p>
</div>
<h3 id="servicebus.azure.com/v1api20210101preview.UserAssignedIdentityProperties">UserAssignedIdentityProperties
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.KeyVaultProperties">KeyVaultProperties</a>)
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
<code>userAssignedIdentityReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>UserAssignedIdentityReference: ARM ID of user Identity selected for encryption</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.UserAssignedIdentityProperties_ARM">UserAssignedIdentityProperties_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.KeyVaultProperties_ARM">KeyVaultProperties_ARM</a>)
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
</td>
</tr>
</tbody>
</table>
<h3 id="servicebus.azure.com/v1api20210101preview.UserAssignedIdentityProperties_STATUS">UserAssignedIdentityProperties_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.KeyVaultProperties_STATUS">KeyVaultProperties_STATUS</a>)
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
<h3 id="servicebus.azure.com/v1api20210101preview.UserAssignedIdentityProperties_STATUS_ARM">UserAssignedIdentityProperties_STATUS_ARM
</h3>
<p>
(<em>Appears on:</em><a href="#servicebus.azure.com/v1api20210101preview.KeyVaultProperties_STATUS_ARM">KeyVaultProperties_STATUS_ARM</a>)
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
<hr/>
