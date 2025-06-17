---

title: insights.azure.com/v1api20240101preview

linktitle: v1api20240101preview
-------------------------------

APIVersion{#APIVersion}
-----------------------

| Value                | Description |
|----------------------|-------------|
| "2024-01-01-preview" |             |

ScheduledQueryRule{#ScheduledQueryRule}
---------------------------------------

Generator information: - Generated from: /monitor/resource-manager/Microsoft.Insights/preview/2024-01-01-preview/scheduledQueryRule_API.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Insights/&ZeroWidthSpace;scheduledQueryRules/&ZeroWidthSpace;{ruleName}

Used by: [ScheduledQueryRuleList](#ScheduledQueryRuleList).

| Property                                                                                | Description | Type                                                                                |
|-----------------------------------------------------------------------------------------|-------------|-------------------------------------------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta)     |             |                                                                                     |
| [metav1.ObjectMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ObjectMeta) |             |                                                                                     |
| spec                                                                                    |             | [ScheduledQueryRule_Spec](#ScheduledQueryRule_Spec)<br/><small>Optional</small>     |
| status                                                                                  |             | [ScheduledQueryRule_STATUS](#ScheduledQueryRule_STATUS)<br/><small>Optional</small> |

### ScheduledQueryRule_Spec {#ScheduledQueryRule_Spec}

| Property                              | Description                                                                                                                                                                                                                                                                                                                                                       | Type                                                                                                                                                                 |
|---------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| actions                               | Actions to invoke when the alert fires.                                                                                                                                                                                                                                                                                                                           | [Actions](#Actions)<br/><small>Optional</small>                                                                                                                      |
| autoMitigate                          | The flag that indicates whether the alert should be automatically resolved or not. The default is true. Relevant only for rules of the kind LogAlert.                                                                                                                                                                                                             | bool<br/><small>Optional</small>                                                                                                                                     |
| azureName                             | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                                   |
| checkWorkspaceAlertsStorageConfigured | The flag which indicates whether this scheduled query rule should be stored in the customer's storage. The default is false. Relevant only for rules of the kind LogAlert.                                                                                                                                                                                        | bool<br/><small>Optional</small>                                                                                                                                     |
| criteria                              | The rule criteria that defines the conditions of the scheduled query rule.                                                                                                                                                                                                                                                                                        | [ScheduledQueryRuleCriteria](#ScheduledQueryRuleCriteria)<br/><small>Optional</small>                                                                                |
| description                           | The description of the scheduled query rule.                                                                                                                                                                                                                                                                                                                      | string<br/><small>Optional</small>                                                                                                                                   |
| displayName                           | The display name of the alert rule                                                                                                                                                                                                                                                                                                                                | string<br/><small>Optional</small>                                                                                                                                   |
| enabled                               | The flag which indicates whether this scheduled query rule is enabled. Value should be true or false                                                                                                                                                                                                                                                              | bool<br/><small>Optional</small>                                                                                                                                     |
| evaluationFrequency                   | How often the scheduled query rule is evaluated represented in ISO 8601 duration format. Relevant and required only for rules of the kind LogAlert.                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                                   |
| identity                              | The identity of the resource.                                                                                                                                                                                                                                                                                                                                     | [Identity](#Identity)<br/><small>Optional</small>                                                                                                                    |
| kind                                  | Indicates the type of scheduled query rule. The default is LogAlert.                                                                                                                                                                                                                                                                                              | [ScheduledQueryRule_Kind_Spec](#ScheduledQueryRule_Kind_Spec)<br/><small>Optional</small>                                                                            |
| location                              | The geo-location where the resource lives                                                                                                                                                                                                                                                                                                                         | string<br/><small>Required</small>                                                                                                                                   |
| muteActionsDuration                   | Mute actions for the chosen period of time (in ISO 8601 duration format) after the alert is fired. Relevant only for rules of the kind LogAlert.                                                                                                                                                                                                                  | string<br/><small>Optional</small>                                                                                                                                   |
| operatorSpec                          | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                                                                                                   | [ScheduledQueryRuleOperatorSpec](#ScheduledQueryRuleOperatorSpec)<br/><small>Optional</small>                                                                        |
| overrideQueryTimeRange                | If specified then overrides the query time range (default is WindowSize*NumberOfEvaluationPeriods). Relevant only for rules of the kind LogAlert.                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                                                   |
| owner                                 | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a resources.azure.com/ResourceGroup resource                                                                      | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |
| resolveConfiguration                  | Defines the configuration for resolving fired alerts. Relevant only for rules of the kind LogAlert.                                                                                                                                                                                                                                                               | [RuleResolveConfiguration](#RuleResolveConfiguration)<br/><small>Optional</small>                                                                                    |
| scopesReferences                      | The list of resource id's that this scheduled query rule is scoped to.                                                                                                                                                                                                                                                                                            | [genruntime.ResourceReference[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>         |
| severity                              | Severity of the alert. Should be an integer between [0-4]. Value of 0 is severest. Relevant and required only for rules of the kind LogAlert.                                                                                                                                                                                                                     | [ScheduledQueryRuleProperties_Severity](#ScheduledQueryRuleProperties_Severity)<br/><small>Optional</small>                                                          |
| skipQueryValidation                   | The flag which indicates whether the provided query should be validated or not. The default is false. Relevant only for rules of the kind LogAlert.                                                                                                                                                                                                               | bool<br/><small>Optional</small>                                                                                                                                     |
| tags                                  | Resource tags.                                                                                                                                                                                                                                                                                                                                                    | map[string]string<br/><small>Optional</small>                                                                                                                        |
| targetResourceTypes                   | List of resource type of the target resource(s) on which the alert is created/updated. For example if the scope is a resource group and targetResourceTypes is Microsoft.Compute/virtualMachines, then a different alert will be fired for each virtual machine in the resource group which meet the alert criteria. Relevant only for rules of the kind LogAlert | string[]<br/><small>Optional</small>                                                                                                                                 |
| windowSize                            | The period of time (in ISO 8601 duration format) on which the Alert query will be executed (bin size). Relevant and required only for rules of the kind LogAlert.                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                                                   |

### ScheduledQueryRule_STATUS{#ScheduledQueryRule_STATUS}

| Property                              | Description                                                                                                                                                                                                                                                                                                                                                                                          | Type                                                                                                                                                    |
|---------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| actions                               | Actions to invoke when the alert fires.                                                                                                                                                                                                                                                                                                                                                              | [Actions_STATUS](#Actions_STATUS)<br/><small>Optional</small>                                                                                           |
| autoMitigate                          | The flag that indicates whether the alert should be automatically resolved or not. The default is true. Relevant only for rules of the kind LogAlert.                                                                                                                                                                                                                                                | bool<br/><small>Optional</small>                                                                                                                        |
| checkWorkspaceAlertsStorageConfigured | The flag which indicates whether this scheduled query rule should be stored in the customer's storage. The default is false. Relevant only for rules of the kind LogAlert.                                                                                                                                                                                                                           | bool<br/><small>Optional</small>                                                                                                                        |
| conditions                            | The observed state of the resource                                                                                                                                                                                                                                                                                                                                                                   | [conditions.Condition[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions#Condition)<br/><small>Optional</small> |
| createdWithApiVersion                 | The api-version used when creating this alert rule                                                                                                                                                                                                                                                                                                                                                   | string<br/><small>Optional</small>                                                                                                                      |
| criteria                              | The rule criteria that defines the conditions of the scheduled query rule.                                                                                                                                                                                                                                                                                                                           | [ScheduledQueryRuleCriteria_STATUS](#ScheduledQueryRuleCriteria_STATUS)<br/><small>Optional</small>                                                     |
| description                           | The description of the scheduled query rule.                                                                                                                                                                                                                                                                                                                                                         | string<br/><small>Optional</small>                                                                                                                      |
| displayName                           | The display name of the alert rule                                                                                                                                                                                                                                                                                                                                                                   | string<br/><small>Optional</small>                                                                                                                      |
| enabled                               | The flag which indicates whether this scheduled query rule is enabled. Value should be true or false                                                                                                                                                                                                                                                                                                 | bool<br/><small>Optional</small>                                                                                                                        |
| etag                                  | The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention. Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields. | string<br/><small>Optional</small>                                                                                                                      |
| evaluationFrequency                   | How often the scheduled query rule is evaluated represented in ISO 8601 duration format. Relevant and required only for rules of the kind LogAlert.                                                                                                                                                                                                                                                  | string<br/><small>Optional</small>                                                                                                                      |
| id                                    | Fully qualified resource ID for the resource. Ex - /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;{resourceProviderNamespace}/&ZeroWidthSpace;{resourceType}/&ZeroWidthSpace;{resourceName}                                                                            | string<br/><small>Optional</small>                                                                                                                      |
| identity                              | The identity of the resource.                                                                                                                                                                                                                                                                                                                                                                        | [Identity_STATUS](#Identity_STATUS)<br/><small>Optional</small>                                                                                         |
| isLegacyLogAnalyticsRule              | True if alert rule is legacy Log Analytic rule                                                                                                                                                                                                                                                                                                                                                       | bool<br/><small>Optional</small>                                                                                                                        |
| isWorkspaceAlertsStorageConfigured    | The flag which indicates whether this scheduled query rule has been configured to be stored in the customer's storage. The default is false.                                                                                                                                                                                                                                                         | bool<br/><small>Optional</small>                                                                                                                        |
| kind                                  | Indicates the type of scheduled query rule. The default is LogAlert.                                                                                                                                                                                                                                                                                                                                 | [ScheduledQueryRule_Kind_STATUS](#ScheduledQueryRule_Kind_STATUS)<br/><small>Optional</small>                                                           |
| location                              | The geo-location where the resource lives                                                                                                                                                                                                                                                                                                                                                            | string<br/><small>Optional</small>                                                                                                                      |
| muteActionsDuration                   | Mute actions for the chosen period of time (in ISO 8601 duration format) after the alert is fired. Relevant only for rules of the kind LogAlert.                                                                                                                                                                                                                                                     | string<br/><small>Optional</small>                                                                                                                      |
| name                                  | The name of the resource                                                                                                                                                                                                                                                                                                                                                                             | string<br/><small>Optional</small>                                                                                                                      |
| overrideQueryTimeRange                | If specified then overrides the query time range (default is WindowSize*NumberOfEvaluationPeriods). Relevant only for rules of the kind LogAlert.                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                      |
| resolveConfiguration                  | Defines the configuration for resolving fired alerts. Relevant only for rules of the kind LogAlert.                                                                                                                                                                                                                                                                                                  | [RuleResolveConfiguration_STATUS](#RuleResolveConfiguration_STATUS)<br/><small>Optional</small>                                                         |
| scopes                                | The list of resource id's that this scheduled query rule is scoped to.                                                                                                                                                                                                                                                                                                                               | string[]<br/><small>Optional</small>                                                                                                                    |
| severity                              | Severity of the alert. Should be an integer between [0-4]. Value of 0 is severest. Relevant and required only for rules of the kind LogAlert.                                                                                                                                                                                                                                                        | [ScheduledQueryRuleProperties_Severity_STATUS](#ScheduledQueryRuleProperties_Severity_STATUS)<br/><small>Optional</small>                               |
| skipQueryValidation                   | The flag which indicates whether the provided query should be validated or not. The default is false. Relevant only for rules of the kind LogAlert.                                                                                                                                                                                                                                                  | bool<br/><small>Optional</small>                                                                                                                        |
| systemData                            | SystemData of ScheduledQueryRule.                                                                                                                                                                                                                                                                                                                                                                    | [SystemData_STATUS](#SystemData_STATUS)<br/><small>Optional</small>                                                                                     |
| tags                                  | Resource tags.                                                                                                                                                                                                                                                                                                                                                                                       | map[string]string<br/><small>Optional</small>                                                                                                           |
| targetResourceTypes                   | List of resource type of the target resource(s) on which the alert is created/updated. For example if the scope is a resource group and targetResourceTypes is Microsoft.Compute/virtualMachines, then a different alert will be fired for each virtual machine in the resource group which meet the alert criteria. Relevant only for rules of the kind LogAlert                                    | string[]<br/><small>Optional</small>                                                                                                                    |
| type                                  | The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"                                                                                                                                                                                                                                                                                            | string<br/><small>Optional</small>                                                                                                                      |
| windowSize                            | The period of time (in ISO 8601 duration format) on which the Alert query will be executed (bin size). Relevant and required only for rules of the kind LogAlert.                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                      |

ScheduledQueryRuleList{#ScheduledQueryRuleList}
-----------------------------------------------

Generator information: - Generated from: /monitor/resource-manager/Microsoft.Insights/preview/2024-01-01-preview/scheduledQueryRule_API.json - ARM URI: /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.Insights/&ZeroWidthSpace;scheduledQueryRules/&ZeroWidthSpace;{ruleName}

| Property                                                                            | Description | Type                                                                    |
|-------------------------------------------------------------------------------------|-------------|-------------------------------------------------------------------------|
| [metav1.TypeMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#TypeMeta) |             |                                                                         |
| [metav1.ListMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ListMeta) |             |                                                                         |
| items                                                                               |             | [ScheduledQueryRule[]](#ScheduledQueryRule)<br/><small>Optional</small> |

ScheduledQueryRule_Spec{#ScheduledQueryRule_Spec}
-------------------------------------------------

Used by: [ScheduledQueryRule](#ScheduledQueryRule).

| Property                              | Description                                                                                                                                                                                                                                                                                                                                                       | Type                                                                                                                                                                 |
|---------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| actions                               | Actions to invoke when the alert fires.                                                                                                                                                                                                                                                                                                                           | [Actions](#Actions)<br/><small>Optional</small>                                                                                                                      |
| autoMitigate                          | The flag that indicates whether the alert should be automatically resolved or not. The default is true. Relevant only for rules of the kind LogAlert.                                                                                                                                                                                                             | bool<br/><small>Optional</small>                                                                                                                                     |
| azureName                             | The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it doesn't have to be.                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                                   |
| checkWorkspaceAlertsStorageConfigured | The flag which indicates whether this scheduled query rule should be stored in the customer's storage. The default is false. Relevant only for rules of the kind LogAlert.                                                                                                                                                                                        | bool<br/><small>Optional</small>                                                                                                                                     |
| criteria                              | The rule criteria that defines the conditions of the scheduled query rule.                                                                                                                                                                                                                                                                                        | [ScheduledQueryRuleCriteria](#ScheduledQueryRuleCriteria)<br/><small>Optional</small>                                                                                |
| description                           | The description of the scheduled query rule.                                                                                                                                                                                                                                                                                                                      | string<br/><small>Optional</small>                                                                                                                                   |
| displayName                           | The display name of the alert rule                                                                                                                                                                                                                                                                                                                                | string<br/><small>Optional</small>                                                                                                                                   |
| enabled                               | The flag which indicates whether this scheduled query rule is enabled. Value should be true or false                                                                                                                                                                                                                                                              | bool<br/><small>Optional</small>                                                                                                                                     |
| evaluationFrequency                   | How often the scheduled query rule is evaluated represented in ISO 8601 duration format. Relevant and required only for rules of the kind LogAlert.                                                                                                                                                                                                               | string<br/><small>Optional</small>                                                                                                                                   |
| identity                              | The identity of the resource.                                                                                                                                                                                                                                                                                                                                     | [Identity](#Identity)<br/><small>Optional</small>                                                                                                                    |
| kind                                  | Indicates the type of scheduled query rule. The default is LogAlert.                                                                                                                                                                                                                                                                                              | [ScheduledQueryRule_Kind_Spec](#ScheduledQueryRule_Kind_Spec)<br/><small>Optional</small>                                                                            |
| location                              | The geo-location where the resource lives                                                                                                                                                                                                                                                                                                                         | string<br/><small>Required</small>                                                                                                                                   |
| muteActionsDuration                   | Mute actions for the chosen period of time (in ISO 8601 duration format) after the alert is fired. Relevant only for rules of the kind LogAlert.                                                                                                                                                                                                                  | string<br/><small>Optional</small>                                                                                                                                   |
| operatorSpec                          | The specification for configuring operator behavior. This field is interpreted by the operator and not passed directly to Azure                                                                                                                                                                                                                                   | [ScheduledQueryRuleOperatorSpec](#ScheduledQueryRuleOperatorSpec)<br/><small>Optional</small>                                                                        |
| overrideQueryTimeRange                | If specified then overrides the query time range (default is WindowSize*NumberOfEvaluationPeriods). Relevant only for rules of the kind LogAlert.                                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                                                   |
| owner                                 | The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a reference to a resources.azure.com/ResourceGroup resource                                                                      | [genruntime.KnownResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference)<br/><small>Required</small> |
| resolveConfiguration                  | Defines the configuration for resolving fired alerts. Relevant only for rules of the kind LogAlert.                                                                                                                                                                                                                                                               | [RuleResolveConfiguration](#RuleResolveConfiguration)<br/><small>Optional</small>                                                                                    |
| scopesReferences                      | The list of resource id's that this scheduled query rule is scoped to.                                                                                                                                                                                                                                                                                            | [genruntime.ResourceReference[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small>         |
| severity                              | Severity of the alert. Should be an integer between [0-4]. Value of 0 is severest. Relevant and required only for rules of the kind LogAlert.                                                                                                                                                                                                                     | [ScheduledQueryRuleProperties_Severity](#ScheduledQueryRuleProperties_Severity)<br/><small>Optional</small>                                                          |
| skipQueryValidation                   | The flag which indicates whether the provided query should be validated or not. The default is false. Relevant only for rules of the kind LogAlert.                                                                                                                                                                                                               | bool<br/><small>Optional</small>                                                                                                                                     |
| tags                                  | Resource tags.                                                                                                                                                                                                                                                                                                                                                    | map[string]string<br/><small>Optional</small>                                                                                                                        |
| targetResourceTypes                   | List of resource type of the target resource(s) on which the alert is created/updated. For example if the scope is a resource group and targetResourceTypes is Microsoft.Compute/virtualMachines, then a different alert will be fired for each virtual machine in the resource group which meet the alert criteria. Relevant only for rules of the kind LogAlert | string[]<br/><small>Optional</small>                                                                                                                                 |
| windowSize                            | The period of time (in ISO 8601 duration format) on which the Alert query will be executed (bin size). Relevant and required only for rules of the kind LogAlert.                                                                                                                                                                                                 | string<br/><small>Optional</small>                                                                                                                                   |

ScheduledQueryRule_STATUS{#ScheduledQueryRule_STATUS}
-----------------------------------------------------

Used by: [ScheduledQueryRule](#ScheduledQueryRule).

| Property                              | Description                                                                                                                                                                                                                                                                                                                                                                                          | Type                                                                                                                                                    |
|---------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| actions                               | Actions to invoke when the alert fires.                                                                                                                                                                                                                                                                                                                                                              | [Actions_STATUS](#Actions_STATUS)<br/><small>Optional</small>                                                                                           |
| autoMitigate                          | The flag that indicates whether the alert should be automatically resolved or not. The default is true. Relevant only for rules of the kind LogAlert.                                                                                                                                                                                                                                                | bool<br/><small>Optional</small>                                                                                                                        |
| checkWorkspaceAlertsStorageConfigured | The flag which indicates whether this scheduled query rule should be stored in the customer's storage. The default is false. Relevant only for rules of the kind LogAlert.                                                                                                                                                                                                                           | bool<br/><small>Optional</small>                                                                                                                        |
| conditions                            | The observed state of the resource                                                                                                                                                                                                                                                                                                                                                                   | [conditions.Condition[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions#Condition)<br/><small>Optional</small> |
| createdWithApiVersion                 | The api-version used when creating this alert rule                                                                                                                                                                                                                                                                                                                                                   | string<br/><small>Optional</small>                                                                                                                      |
| criteria                              | The rule criteria that defines the conditions of the scheduled query rule.                                                                                                                                                                                                                                                                                                                           | [ScheduledQueryRuleCriteria_STATUS](#ScheduledQueryRuleCriteria_STATUS)<br/><small>Optional</small>                                                     |
| description                           | The description of the scheduled query rule.                                                                                                                                                                                                                                                                                                                                                         | string<br/><small>Optional</small>                                                                                                                      |
| displayName                           | The display name of the alert rule                                                                                                                                                                                                                                                                                                                                                                   | string<br/><small>Optional</small>                                                                                                                      |
| enabled                               | The flag which indicates whether this scheduled query rule is enabled. Value should be true or false                                                                                                                                                                                                                                                                                                 | bool<br/><small>Optional</small>                                                                                                                        |
| etag                                  | The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal etag convention. Entity tags are used for comparing two or more entities from the same requested resource. HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and If-Range (section 14.27) header fields. | string<br/><small>Optional</small>                                                                                                                      |
| evaluationFrequency                   | How often the scheduled query rule is evaluated represented in ISO 8601 duration format. Relevant and required only for rules of the kind LogAlert.                                                                                                                                                                                                                                                  | string<br/><small>Optional</small>                                                                                                                      |
| id                                    | Fully qualified resource ID for the resource. Ex - /&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;{resourceProviderNamespace}/&ZeroWidthSpace;{resourceType}/&ZeroWidthSpace;{resourceName}                                                                            | string<br/><small>Optional</small>                                                                                                                      |
| identity                              | The identity of the resource.                                                                                                                                                                                                                                                                                                                                                                        | [Identity_STATUS](#Identity_STATUS)<br/><small>Optional</small>                                                                                         |
| isLegacyLogAnalyticsRule              | True if alert rule is legacy Log Analytic rule                                                                                                                                                                                                                                                                                                                                                       | bool<br/><small>Optional</small>                                                                                                                        |
| isWorkspaceAlertsStorageConfigured    | The flag which indicates whether this scheduled query rule has been configured to be stored in the customer's storage. The default is false.                                                                                                                                                                                                                                                         | bool<br/><small>Optional</small>                                                                                                                        |
| kind                                  | Indicates the type of scheduled query rule. The default is LogAlert.                                                                                                                                                                                                                                                                                                                                 | [ScheduledQueryRule_Kind_STATUS](#ScheduledQueryRule_Kind_STATUS)<br/><small>Optional</small>                                                           |
| location                              | The geo-location where the resource lives                                                                                                                                                                                                                                                                                                                                                            | string<br/><small>Optional</small>                                                                                                                      |
| muteActionsDuration                   | Mute actions for the chosen period of time (in ISO 8601 duration format) after the alert is fired. Relevant only for rules of the kind LogAlert.                                                                                                                                                                                                                                                     | string<br/><small>Optional</small>                                                                                                                      |
| name                                  | The name of the resource                                                                                                                                                                                                                                                                                                                                                                             | string<br/><small>Optional</small>                                                                                                                      |
| overrideQueryTimeRange                | If specified then overrides the query time range (default is WindowSize*NumberOfEvaluationPeriods). Relevant only for rules of the kind LogAlert.                                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                      |
| resolveConfiguration                  | Defines the configuration for resolving fired alerts. Relevant only for rules of the kind LogAlert.                                                                                                                                                                                                                                                                                                  | [RuleResolveConfiguration_STATUS](#RuleResolveConfiguration_STATUS)<br/><small>Optional</small>                                                         |
| scopes                                | The list of resource id's that this scheduled query rule is scoped to.                                                                                                                                                                                                                                                                                                                               | string[]<br/><small>Optional</small>                                                                                                                    |
| severity                              | Severity of the alert. Should be an integer between [0-4]. Value of 0 is severest. Relevant and required only for rules of the kind LogAlert.                                                                                                                                                                                                                                                        | [ScheduledQueryRuleProperties_Severity_STATUS](#ScheduledQueryRuleProperties_Severity_STATUS)<br/><small>Optional</small>                               |
| skipQueryValidation                   | The flag which indicates whether the provided query should be validated or not. The default is false. Relevant only for rules of the kind LogAlert.                                                                                                                                                                                                                                                  | bool<br/><small>Optional</small>                                                                                                                        |
| systemData                            | SystemData of ScheduledQueryRule.                                                                                                                                                                                                                                                                                                                                                                    | [SystemData_STATUS](#SystemData_STATUS)<br/><small>Optional</small>                                                                                     |
| tags                                  | Resource tags.                                                                                                                                                                                                                                                                                                                                                                                       | map[string]string<br/><small>Optional</small>                                                                                                           |
| targetResourceTypes                   | List of resource type of the target resource(s) on which the alert is created/updated. For example if the scope is a resource group and targetResourceTypes is Microsoft.Compute/virtualMachines, then a different alert will be fired for each virtual machine in the resource group which meet the alert criteria. Relevant only for rules of the kind LogAlert                                    | string[]<br/><small>Optional</small>                                                                                                                    |
| type                                  | The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"                                                                                                                                                                                                                                                                                            | string<br/><small>Optional</small>                                                                                                                      |
| windowSize                            | The period of time (in ISO 8601 duration format) on which the Alert query will be executed (bin size). Relevant and required only for rules of the kind LogAlert.                                                                                                                                                                                                                                    | string<br/><small>Optional</small>                                                                                                                      |

Actions{#Actions}
-----------------

Actions to invoke when the alert fires.

Used by: [ScheduledQueryRule_Spec](#ScheduledQueryRule_Spec).

| Property               | Description                                               | Type                                                                                                                                                         |
|------------------------|-----------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------|
| actionGroupsReferences | Action Group resource Ids to invoke when the alert fires. | [genruntime.ResourceReference[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| actionProperties       | The properties of an action properties.                   | map[string]string<br/><small>Optional</small>                                                                                                                |
| customProperties       | The properties of an alert payload.                       | map[string]string<br/><small>Optional</small>                                                                                                                |

Actions_STATUS{#Actions_STATUS}
-------------------------------

Actions to invoke when the alert fires.

Used by: [ScheduledQueryRule_STATUS](#ScheduledQueryRule_STATUS).

| Property         | Description                                               | Type                                          |
|------------------|-----------------------------------------------------------|-----------------------------------------------|
| actionGroups     | Action Group resource Ids to invoke when the alert fires. | string[]<br/><small>Optional</small>          |
| actionProperties | The properties of an action properties.                   | map[string]string<br/><small>Optional</small> |
| customProperties | The properties of an alert payload.                       | map[string]string<br/><small>Optional</small> |

Identity{#Identity}
-------------------

Identity for the resource.

Used by: [ScheduledQueryRule_Spec](#ScheduledQueryRule_Spec).

| Property               | Description                                                                                                                                                                                                                                                                                                                                                                                                                 | Type                                                                                      |
|------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------|
| type                   | Type of managed service identity.                                                                                                                                                                                                                                                                                                                                                                                           | [Identity_Type](#Identity_Type)<br/><small>Required</small>                               |
| userAssignedIdentities | The list of user identities associated with the resource. The user identity dictionary key references will be ARM resource ids in the form: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ManagedIdentity/&ZeroWidthSpace;userAssignedIdentities/&ZeroWidthSpace;{identityName}'. | [UserAssignedIdentityDetails[]](#UserAssignedIdentityDetails)<br/><small>Optional</small> |

Identity_STATUS{#Identity_STATUS}
---------------------------------

Identity for the resource.

Used by: [ScheduledQueryRule_STATUS](#ScheduledQueryRule_STATUS).

| Property               | Description                                                                                                                                                                                                                                                                                                                                                                                                                 | Type                                                                                                   |
|------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------|
| principalId            | The principal ID of resource identity.                                                                                                                                                                                                                                                                                                                                                                                      | string<br/><small>Optional</small>                                                                     |
| tenantId               | The tenant ID of resource.                                                                                                                                                                                                                                                                                                                                                                                                  | string<br/><small>Optional</small>                                                                     |
| type                   | Type of managed service identity.                                                                                                                                                                                                                                                                                                                                                                                           | [Identity_Type_STATUS](#Identity_Type_STATUS)<br/><small>Optional</small>                              |
| userAssignedIdentities | The list of user identities associated with the resource. The user identity dictionary key references will be ARM resource ids in the form: '/&ZeroWidthSpace;subscriptions/&ZeroWidthSpace;{subscriptionId}/&ZeroWidthSpace;resourceGroups/&ZeroWidthSpace;{resourceGroupName}/&ZeroWidthSpace;providers/&ZeroWidthSpace;Microsoft.ManagedIdentity/&ZeroWidthSpace;userAssignedIdentities/&ZeroWidthSpace;{identityName}'. | [map[string]UserIdentityProperties_STATUS](#UserIdentityProperties_STATUS)<br/><small>Optional</small> |

RuleResolveConfiguration{#RuleResolveConfiguration}
---------------------------------------------------

TBD. Relevant only for rules of the kind LogAlert.

Used by: [ScheduledQueryRule_Spec](#ScheduledQueryRule_Spec).

| Property      | Description                                                                                                                            | Type                               |
|---------------|----------------------------------------------------------------------------------------------------------------------------------------|------------------------------------|
| autoResolved  | The flag that indicates whether or not to auto resolve a fired alert.                                                                  | bool<br/><small>Optional</small>   |
| timeToResolve | The duration a rule must evaluate as healthy before the fired alert is automatically resolved represented in ISO 8601 duration format. | string<br/><small>Optional</small> |

RuleResolveConfiguration_STATUS{#RuleResolveConfiguration_STATUS}
-----------------------------------------------------------------

TBD. Relevant only for rules of the kind LogAlert.

Used by: [ScheduledQueryRule_STATUS](#ScheduledQueryRule_STATUS).

| Property      | Description                                                                                                                            | Type                               |
|---------------|----------------------------------------------------------------------------------------------------------------------------------------|------------------------------------|
| autoResolved  | The flag that indicates whether or not to auto resolve a fired alert.                                                                  | bool<br/><small>Optional</small>   |
| timeToResolve | The duration a rule must evaluate as healthy before the fired alert is automatically resolved represented in ISO 8601 duration format. | string<br/><small>Optional</small> |

ScheduledQueryRule_Kind_Spec{#ScheduledQueryRule_Kind_Spec}
-----------------------------------------------------------

Used by: [ScheduledQueryRule_Spec](#ScheduledQueryRule_Spec).

| Value           | Description |
|-----------------|-------------|
| "EventLogAlert" |             |
| "LogAlert"      |             |
| "LogToMetric"   |             |

ScheduledQueryRule_Kind_STATUS{#ScheduledQueryRule_Kind_STATUS}
---------------------------------------------------------------

Used by: [ScheduledQueryRule_STATUS](#ScheduledQueryRule_STATUS).

| Value           | Description |
|-----------------|-------------|
| "EventLogAlert" |             |
| "LogAlert"      |             |
| "LogToMetric"   |             |

ScheduledQueryRuleCriteria{#ScheduledQueryRuleCriteria}
-------------------------------------------------------

The rule criteria that defines the conditions of the scheduled query rule.

Used by: [ScheduledQueryRule_Spec](#ScheduledQueryRule_Spec).

| Property | Description                                                   | Type                                                  |
|----------|---------------------------------------------------------------|-------------------------------------------------------|
| allOf    | A list of conditions to evaluate against the specified scopes | [Condition[]](#Condition)<br/><small>Optional</small> |

ScheduledQueryRuleCriteria_STATUS{#ScheduledQueryRuleCriteria_STATUS}
---------------------------------------------------------------------

The rule criteria that defines the conditions of the scheduled query rule.

Used by: [ScheduledQueryRule_STATUS](#ScheduledQueryRule_STATUS).

| Property | Description                                                   | Type                                                                |
|----------|---------------------------------------------------------------|---------------------------------------------------------------------|
| allOf    | A list of conditions to evaluate against the specified scopes | [Condition_STATUS[]](#Condition_STATUS)<br/><small>Optional</small> |

ScheduledQueryRuleOperatorSpec{#ScheduledQueryRuleOperatorSpec}
---------------------------------------------------------------

Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure

Used by: [ScheduledQueryRule_Spec](#ScheduledQueryRule_Spec).

| Property             | Description                                                                                   | Type                                                                                                                                                                |
|----------------------|-----------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| configMapExpressions | configures where to place operator written dynamic ConfigMaps (created with CEL expressions). | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |
| secretExpressions    | configures where to place operator written dynamic secrets (created with CEL expressions).    | [core.DestinationExpression[]](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime/core#DestinationExpression)<br/><small>Optional</small> |

ScheduledQueryRuleProperties_Severity{#ScheduledQueryRuleProperties_Severity}
-----------------------------------------------------------------------------

Used by: [ScheduledQueryRule_Spec](#ScheduledQueryRule_Spec).

| Value | Description |
|-------|-------------|
| 0     |             |
| 1     |             |
| 2     |             |
| 3     |             |
| 4     |             |

ScheduledQueryRuleProperties_Severity_STATUS{#ScheduledQueryRuleProperties_Severity_STATUS}
-------------------------------------------------------------------------------------------

Used by: [ScheduledQueryRule_STATUS](#ScheduledQueryRule_STATUS).

| Value | Description |
|-------|-------------|
| 0     |             |
| 1     |             |
| 2     |             |
| 3     |             |
| 4     |             |

SystemData_STATUS{#SystemData_STATUS}
-------------------------------------

Metadata pertaining to creation and last modification of the resource.

Used by: [ScheduledQueryRule_STATUS](#ScheduledQueryRule_STATUS).

| Property           | Description                                           | Type                                                                                                      |
|--------------------|-------------------------------------------------------|-----------------------------------------------------------------------------------------------------------|
| createdAt          | The timestamp of resource creation (UTC).             | string<br/><small>Optional</small>                                                                        |
| createdBy          | The identity that created the resource.               | string<br/><small>Optional</small>                                                                        |
| createdByType      | The type of identity that created the resource.       | [SystemData_CreatedByType_STATUS](#SystemData_CreatedByType_STATUS)<br/><small>Optional</small>           |
| lastModifiedAt     | The timestamp of resource last modification (UTC)     | string<br/><small>Optional</small>                                                                        |
| lastModifiedBy     | The identity that last modified the resource.         | string<br/><small>Optional</small>                                                                        |
| lastModifiedByType | The type of identity that last modified the resource. | [SystemData_LastModifiedByType_STATUS](#SystemData_LastModifiedByType_STATUS)<br/><small>Optional</small> |

Condition{#Condition}
---------------------

A condition of the scheduled query rule.

Used by: [ScheduledQueryRuleCriteria](#ScheduledQueryRuleCriteria).

| Property                  | Description                                                                                                                                                                                                                                                 | Type                                                                                                                                                       |
|---------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| alertSensitivity          | The extent of deviation required to trigger an alert. Allowed values are `Low`, `Medium` and `High`. This will affect how tight the threshold is to the metric series pattern. Relevant and required only for dynamic threshold rules of the kind LogAlert. | string<br/><small>Optional</small>                                                                                                                         |
| criterionType             | Specifies the type of threshold criteria                                                                                                                                                                                                                    | [Condition_CriterionType](#Condition_CriterionType)<br/><small>Optional</small>                                                                            |
| dimensions                | List of Dimensions conditions                                                                                                                                                                                                                               | [Dimension[]](#Dimension)<br/><small>Optional</small>                                                                                                      |
| failingPeriods            | The minimum number of violations required within the selected lookback time window required to raise an alert. Relevant only for rules of the kind LogAlert.                                                                                                | [Condition_FailingPeriods](#Condition_FailingPeriods)<br/><small>Optional</small>                                                                          |
| ignoreDataBefore          | Use this option to set the date from which to start learning the metric historical data and calculate the dynamic thresholds (in ISO8601 format). Relevant only for dynamic threshold rules of the kind LogAlert.                                           | string<br/><small>Optional</small>                                                                                                                         |
| metricMeasureColumn       | The column containing the metric measure number. Relevant only for rules of the kind LogAlert.                                                                                                                                                              | string<br/><small>Optional</small>                                                                                                                         |
| metricName                | The name of the metric to be sent. Relevant and required only for rules of the kind LogToMetric.                                                                                                                                                            | string<br/><small>Optional</small>                                                                                                                         |
| operator                  | The criteria operator. Relevant and required only for rules of the kind LogAlert.                                                                                                                                                                           | [Condition_Operator](#Condition_Operator)<br/><small>Optional</small>                                                                                      |
| query                     | Log query alert                                                                                                                                                                                                                                             | string<br/><small>Optional</small>                                                                                                                         |
| resourceIdColumnReference | The column containing the resource id. The content of the column must be a uri formatted as resource id. Relevant only for rules of the kind LogAlert.                                                                                                      | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |
| threshold                 | the criteria threshold value that activates the alert. Relevant and required only for static threshold rules of the kind LogAlert.                                                                                                                          | float64<br/><small>Optional</small>                                                                                                                        |
| timeAggregation           | Aggregation type. Relevant and required only for rules of the kind LogAlert.                                                                                                                                                                                | [Condition_TimeAggregation](#Condition_TimeAggregation)<br/><small>Optional</small>                                                                        |

Condition_STATUS{#Condition_STATUS}
-----------------------------------

A condition of the scheduled query rule.

Used by: [ScheduledQueryRuleCriteria_STATUS](#ScheduledQueryRuleCriteria_STATUS).

| Property            | Description                                                                                                                                                                                                                                                 | Type                                                                                              |
|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------|
| alertSensitivity    | The extent of deviation required to trigger an alert. Allowed values are `Low`, `Medium` and `High`. This will affect how tight the threshold is to the metric series pattern. Relevant and required only for dynamic threshold rules of the kind LogAlert. | string<br/><small>Optional</small>                                                                |
| criterionType       | Specifies the type of threshold criteria                                                                                                                                                                                                                    | [Condition_CriterionType_STATUS](#Condition_CriterionType_STATUS)<br/><small>Optional</small>     |
| dimensions          | List of Dimensions conditions                                                                                                                                                                                                                               | [Dimension_STATUS[]](#Dimension_STATUS)<br/><small>Optional</small>                               |
| failingPeriods      | The minimum number of violations required within the selected lookback time window required to raise an alert. Relevant only for rules of the kind LogAlert.                                                                                                | [Condition_FailingPeriods_STATUS](#Condition_FailingPeriods_STATUS)<br/><small>Optional</small>   |
| ignoreDataBefore    | Use this option to set the date from which to start learning the metric historical data and calculate the dynamic thresholds (in ISO8601 format). Relevant only for dynamic threshold rules of the kind LogAlert.                                           | string<br/><small>Optional</small>                                                                |
| metricMeasureColumn | The column containing the metric measure number. Relevant only for rules of the kind LogAlert.                                                                                                                                                              | string<br/><small>Optional</small>                                                                |
| metricName          | The name of the metric to be sent. Relevant and required only for rules of the kind LogToMetric.                                                                                                                                                            | string<br/><small>Optional</small>                                                                |
| operator            | The criteria operator. Relevant and required only for rules of the kind LogAlert.                                                                                                                                                                           | [Condition_Operator_STATUS](#Condition_Operator_STATUS)<br/><small>Optional</small>               |
| query               | Log query alert                                                                                                                                                                                                                                             | string<br/><small>Optional</small>                                                                |
| resourceIdColumn    | The column containing the resource id. The content of the column must be a uri formatted as resource id. Relevant only for rules of the kind LogAlert.                                                                                                      | string<br/><small>Optional</small>                                                                |
| threshold           | the criteria threshold value that activates the alert. Relevant and required only for static threshold rules of the kind LogAlert.                                                                                                                          | float64<br/><small>Optional</small>                                                               |
| timeAggregation     | Aggregation type. Relevant and required only for rules of the kind LogAlert.                                                                                                                                                                                | [Condition_TimeAggregation_STATUS](#Condition_TimeAggregation_STATUS)<br/><small>Optional</small> |

Identity_Type{#Identity_Type}
-----------------------------

Used by: [Identity](#Identity).

| Value            | Description |
|------------------|-------------|
| "None"           |             |
| "SystemAssigned" |             |
| "UserAssigned"   |             |

Identity_Type_STATUS{#Identity_Type_STATUS}
-------------------------------------------

Used by: [Identity_STATUS](#Identity_STATUS).

| Value            | Description |
|------------------|-------------|
| "None"           |             |
| "SystemAssigned" |             |
| "UserAssigned"   |             |

SystemData_CreatedByType_STATUS{#SystemData_CreatedByType_STATUS}
-----------------------------------------------------------------

Used by: [SystemData_STATUS](#SystemData_STATUS).

| Value             | Description |
|-------------------|-------------|
| "Application"     |             |
| "Key"             |             |
| "ManagedIdentity" |             |
| "User"            |             |

SystemData_LastModifiedByType_STATUS{#SystemData_LastModifiedByType_STATUS}
---------------------------------------------------------------------------

Used by: [SystemData_STATUS](#SystemData_STATUS).

| Value             | Description |
|-------------------|-------------|
| "Application"     |             |
| "Key"             |             |
| "ManagedIdentity" |             |
| "User"            |             |

UserAssignedIdentityDetails{#UserAssignedIdentityDetails}
---------------------------------------------------------

Information about the user assigned identity for the resource

Used by: [Identity](#Identity).

| Property  | Description | Type                                                                                                                                                       |
|-----------|-------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|
| reference |             | [genruntime.ResourceReference](https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference)<br/><small>Optional</small> |

UserIdentityProperties_STATUS{#UserIdentityProperties_STATUS}
-------------------------------------------------------------

User assigned identity properties.

Used by: [Identity_STATUS](#Identity_STATUS).

| Property    | Description                                 | Type                               |
|-------------|---------------------------------------------|------------------------------------|
| clientId    | The client id of user assigned identity.    | string<br/><small>Optional</small> |
| principalId | The principal id of user assigned identity. | string<br/><small>Optional</small> |

Condition_CriterionType{#Condition_CriterionType}
-------------------------------------------------

Used by: [Condition](#Condition).

| Value                       | Description |
|-----------------------------|-------------|
| "DynamicThresholdCriterion" |             |
| "StaticThresholdCriterion"  |             |

Condition_CriterionType_STATUS{#Condition_CriterionType_STATUS}
---------------------------------------------------------------

Used by: [Condition_STATUS](#Condition_STATUS).

| Value                       | Description |
|-----------------------------|-------------|
| "DynamicThresholdCriterion" |             |
| "StaticThresholdCriterion"  |             |

Condition_FailingPeriods{#Condition_FailingPeriods}
---------------------------------------------------

Used by: [Condition](#Condition).

| Property                  | Description                                                                                                                                                                                         | Type                            |
|---------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------|
| minFailingPeriodsToAlert  | The number of violations to trigger an alert. Should be smaller or equal to numberOfEvaluationPeriods. Default value is 1                                                                           | int<br/><small>Optional</small> |
| numberOfEvaluationPeriods | The number of aggregated lookback points. The lookback time window is calculated based on the aggregation granularity (windowSize) and the selected number of aggregated points. Default value is 1 | int<br/><small>Optional</small> |

Condition_FailingPeriods_STATUS{#Condition_FailingPeriods_STATUS}
-----------------------------------------------------------------

Used by: [Condition_STATUS](#Condition_STATUS).

| Property                  | Description                                                                                                                                                                                         | Type                            |
|---------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------|
| minFailingPeriodsToAlert  | The number of violations to trigger an alert. Should be smaller or equal to numberOfEvaluationPeriods. Default value is 1                                                                           | int<br/><small>Optional</small> |
| numberOfEvaluationPeriods | The number of aggregated lookback points. The lookback time window is calculated based on the aggregation granularity (windowSize) and the selected number of aggregated points. Default value is 1 | int<br/><small>Optional</small> |

Condition_Operator{#Condition_Operator}
---------------------------------------

Used by: [Condition](#Condition).

| Value                | Description |
|----------------------|-------------|
| "Equals"             |             |
| "GreaterOrLessThan"  |             |
| "GreaterThan"        |             |
| "GreaterThanOrEqual" |             |
| "LessThan"           |             |
| "LessThanOrEqual"    |             |

Condition_Operator_STATUS{#Condition_Operator_STATUS}
-----------------------------------------------------

Used by: [Condition_STATUS](#Condition_STATUS).

| Value                | Description |
|----------------------|-------------|
| "Equals"             |             |
| "GreaterOrLessThan"  |             |
| "GreaterThan"        |             |
| "GreaterThanOrEqual" |             |
| "LessThan"           |             |
| "LessThanOrEqual"    |             |

Condition_TimeAggregation{#Condition_TimeAggregation}
-----------------------------------------------------

Used by: [Condition](#Condition).

| Value     | Description |
|-----------|-------------|
| "Average" |             |
| "Count"   |             |
| "Maximum" |             |
| "Minimum" |             |
| "Total"   |             |

Condition_TimeAggregation_STATUS{#Condition_TimeAggregation_STATUS}
-------------------------------------------------------------------

Used by: [Condition_STATUS](#Condition_STATUS).

| Value     | Description |
|-----------|-------------|
| "Average" |             |
| "Count"   |             |
| "Maximum" |             |
| "Minimum" |             |
| "Total"   |             |

Dimension{#Dimension}
---------------------

Dimension splitting and filtering definition

Used by: [Condition](#Condition).

| Property | Description                   | Type                                                                  |
|----------|-------------------------------|-----------------------------------------------------------------------|
| name     | Name of the dimension         | string<br/><small>Required</small>                                    |
| operator | Operator for dimension values | [Dimension_Operator](#Dimension_Operator)<br/><small>Required</small> |
| values   | List of dimension values      | string[]<br/><small>Required</small>                                  |

Dimension_STATUS{#Dimension_STATUS}
-----------------------------------

Dimension splitting and filtering definition

Used by: [Condition_STATUS](#Condition_STATUS).

| Property | Description                   | Type                                                                                |
|----------|-------------------------------|-------------------------------------------------------------------------------------|
| name     | Name of the dimension         | string<br/><small>Optional</small>                                                  |
| operator | Operator for dimension values | [Dimension_Operator_STATUS](#Dimension_Operator_STATUS)<br/><small>Optional</small> |
| values   | List of dimension values      | string[]<br/><small>Optional</small>                                                |

Dimension_Operator{#Dimension_Operator}
---------------------------------------

Used by: [Dimension](#Dimension).

| Value     | Description |
|-----------|-------------|
| "Exclude" |             |
| "Include" |             |

Dimension_Operator_STATUS{#Dimension_Operator_STATUS}
-----------------------------------------------------

Used by: [Dimension_STATUS](#Dimension_STATUS).

| Value     | Description |
|-----------|-------------|
| "Exclude" |             |
| "Include" |             |
