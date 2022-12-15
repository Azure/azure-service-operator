---
title: "Breaking Changes"
linkTitle: "Breaking Changes"
weight: 20
menu:
  main:
    weight: 20
layout: single
cascade:
- type: docs
- render: always
---

We go to great lengths to avoid breaking changes as much as possible. However, they do occasionally occur. This page lists the breaking changes that have occurred in the past.

## Breaking Changes in beta.4

In the 'beta.4' release of Azure Service Operator (ASO) we are pivoting to using the Azure Swagger API Specifications as our sole source of truth for our code generator. This change brings with it a significant improvement in fidelity - the code we generate is now much closer to what the Azure Swagger API Specifications describe.

Unfortunately, this change brings with it a few breaking changes. 

We expect that most users will find their resources are unaffected by these changes. However, if you are using a resource that is affected, you will need to take action to migrate your resources to the new format.

The impact of these breaking changes falls into two categories, *immediate migration required* and *upgrade when modified*, distinguished by when you need to take remedial action.

### Immediate migration required

As a part of upgrading ASO in your cluster, you will need to modify these resources as described. In all cases, the changes are straightforward and can be done with a simple text editor.

The process to follow is as follows:

1. Annotate the resource with `serviceoperator.azure.com/reconcile-policy: skip` to prevent ASO from trying to reconcile the resource while you are upgrading.
2. Download the current YAML for the resource using `kubectl` if you don't have it elsewhere.
3. Delete the resource from your cluster using `kubectl delete`. Your Azure resource will be left untouched because of the `reconcile-policy` annotation you added above.
4. [Upgrade ASO](https://azure.github.io/azure-service-operator/introduction/upgrading/) in your cluster.
5. Modify the YAML for your resource to address the breaking changes noted below.
6. Apply the updated YAML to your cluster using `kubectl apply`. If any errors occur, address them.
7. If the `reconcile-policy` annotation is still present, remove the it from the resource.

If you are updating multiple resources, you may want to apply the `reconcile-policy: skip` annotation to all the resources up front, and then remove it only once you've successfully applied all the changes to all the resources.

### Upgrade when modified

You can upgrade ASO in your cluster and the resource will continue to operate normally. If you modify the resource in the future you will need to make the indicated changes at that point.

## Types of breaking changes

| Type                | Description                                                                                                                                                                                                                                                                                                                                                            | Impact                                           |
| :------------------ | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | :----------------------------------------------- |
| Discriminator       | The discriminator value for polymorphic types has been changed to match the name of the property used to specify that option. Previously we were synthesizing a name based on other factors, resulting in longer property names that did not appear in the ARM/Bicep documentation.<br/><br/>*Example*:<br/>`DeliveryRuleCacheExpiration` renamed to `CacheExpiration` | Immediate migration                              |
| Enumeration         | Properties that previously had a base type but no enumeration values have been updated to include the enumeration values.<br/><br/>*Example*:<br/>`KubernetesCluster.KubernetesVersion`.                                                                                                                                                                               | Immediate migration for malformed resources only |
| Inlined             | Objects that were incorrectly generated as nested properties have been inlined.<br/><br/>*Example:<br/>Properties found in `DeadLetterDestination.Properties` have been promoted to `DeadLetterDestination`.                                                                                                                                                           | Immediate migration                              |
| Reference Detection | Id fields now correctly identified as references which now allow for linking to a resource in Kubernetes instead of only in Azure.<br/><br/>*Example:*<br/>Property `VirtualMachineProfile.NetworkProfile.Id` has been changed to `VirtualMachineProfile.NetworkProfile.Reference`.                                                                                    | Immediate migration                              |
| Status Only         | Status properties that cannot be set by the end user on a Spec that were included in the Spec in error.<br/><br/>*Example*:<br/>`Identity.UserAssignedIdentities`                                                                                                                                                                                                      | Upgrade when modified                            |
| Subresource         | Sub-resources that were incorrectly inlined into the parent resource have been moved to a separate resource.<br/><br/>*Example*:<br/>`VirtualNetworkGateway.VirtualNetworkExtendedLocation`                                                                                                                                                                            | Immediate migration                              |
| Unused              | Properties that previously included on Spec but actually had no function have been removed.<br/><br/>*Examples*:<br/>`Location` on child resources that inherit their actual location from their parent resource.<br/>`Tags` on resources that don't support tags.                                                                                                     | Upgrade when modified                            |
| Validation          | Validation rules have been tightened, or added to properties that previously had no validation rules.<br/><br/>*Example:*<br/> The property `BatchAccount.AzureName` now has a more restrictive regular expression.                                                                                                                                                    | Upgrade when modified                            |


## Authorization

### RoleAssignment

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |


## Batch

### BatchAccount

| Property  | Change     | Reason     |
| :-------- | :--------- | :--------- |
| AzureName | Restricted | Validation |

## Cache v2020-12-01

### RedisFirewallRule

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |

### RedisLinkedServer

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |

### RedisPatchSchedule

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |

## Cache v2021-03-01

### RedisEnterpriseDatabase

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |

## CDN v2021-06-01

### ProfilesEndpoint

| Property Path                                                       | Change                                | Reason        |
| ------------------------------------------------------------------- | ------------------------------------- | ------------- |
| DeliveryPolicy.Rules.Actions.DeliveryRuleCacheExpiration            | Renamed to CacheExpiration            | Discriminator |
| DeliveryPolicy.Rules.Actions.DeliveryRuleCacheKeyQueryString        | Renamed to CacheKeyQueryString        | Discriminator |
| DeliveryPolicy.Rules.Actions.DeliveryRuleRequestHeader              | Renamed to ModifyRequestHeader        | Discriminator |
| DeliveryPolicy.Rules.Actions.DeliveryRuleResponseHeader             | Renamed to ModifyResponseHeader       | Discriminator |
| DeliveryPolicy.Rules.Actions.DeliveryRuleRouteConfigurationOverride | Renamed to RouteConfigurationOverride | Discriminator |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleClientPort              | Renamed to ClientPort                 | Discriminator |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleCookies                 | Renamed to Cookies                    | Discriminator |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleHostName                | Renamed to HostName                   | Discriminator |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleHttpVersion             | Renamed to HttpVersion                | Discriminator |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleIsDevice                | Renamed to IsDevice                   | Discriminator |
| DeliveryPolicy.Rules.Conditions.DeliveryRulePostArgs                | Renamed to PostArgs                   | Discriminator |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleQueryString             | Renamed to QueryString                | Discriminator |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleRemoteAddress           | Renamed to RemoteAddress              | Discriminator |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleRequestBody             | Renamed to RequestBody                | Discriminator |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleRequestHeader           | Renamed to RequestHeader              | Discriminator |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleRequestMethod           | Renamed to RequestMethod              | Discriminator |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleRequestScheme           | Renamed to RequestScheme              | Discriminator |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleRequestUri              | Renamed to RequestUri                 | Discriminator |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleServerPort              | Renamed to ServerPort                 | Discriminator |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleSslProtocol             | Renamed to SslProtocol                | Discriminator |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleUrlFileExtension        | Renamed to UrlFileExtension           | Discriminator |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleUrlFileName             | Renamed to UrlFileName                | Discriminator |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleUrlPath                 | Renamed to UrlPath                    | Discriminator |


## Compute v2020-12-01

### VirtualMachineScaleSet

| Property Path                                                                           | Change               | Reason              |
| --------------------------------------------------------------------------------------- | -------------------- | ------------------- |
| VirtualMachineProfile.NetworkProfile.NetworkInterfaceConfigurations.Id                  | Renamed to Reference | Reference Detection |
| VirtualMachineProfile.NetworkProfile.NetworkInterfaceConfigurations.IpConfigurations.Id | Renamed to Reference | Reference Detection |


## Compute v2022-03-01

### VirtualMachineScaleSet

| Property Path                                            | Change               | Reason              |
| -------------------------------------------------------- | -------------------- | ------------------- |
| VirtualMachineProfile.NetworkProfile.Id                  | Renamed to Reference | Reference Detection |
| VirtualMachineProfile.NetworkProfile.IpConfigurations.Id | Renamed to Reference | Reference Detection |

## Container Service

### ManagedCluster

| Property Path                   | Change  | Reason      |
| ------------------------------- | ------- | ----------- |
| Identity.UserAssignedIdentities | Removed | Status Only |

### ManagedClustersAgentPool

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |

## DbForMariaDB v2018-06-01

### Configuration

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |

### Database

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |

### Server

| Property Path                               | Change                        | Reason        |
| :------------------------------------------ | :---------------------------- | :------------ |
| Properties.ServerPropertiesForDefaultCreate | Renamed to Default            | Discriminator |
| Properties.ServerPropertiesForGeoRestore    | Renamed to GeoRestore         | Discriminator |
| Properties.ServerPropertiesForReplica       | Renamed to Replica            | Discriminator |
| Properties.ServerPropertiesForRestore       | Renamed to PointInTimeRestore | Discriminator |


## DbForMySql v2021-05-01

### FlexibleServersDatabase

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |

### FlexibleServersFirewallRule

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |

## DbForPostgreSQL v2021-06-01

### FlexibleServersConfiguration

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |

### FlexibleServersDatabase

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |

### FlexibleServersFirewallRule

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |

## EventGrid v2020-06-01

### Domain

| Property Path                             | Change          | Reason        |
| ----------------------------------------- | --------------- | ------------- |
| InputSchemaMapping.InputSchemaMappingType | Removed         | Discriminator |
| InputSchemaMapping.Properties             | Renamed to Json | Discriminator |

### DomainsTopic

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |

### EventSubscription

| Property Path                           | Change             | Reason  |
| :-------------------------------------- | :----------------- | :------ |
| DeadLetterDestination.Properties        | Promoted to parent | Inlined |
| Destination.AzureFunction.Properties    | Promoted to parent | Inlined |
| Destination.EventHub.Properties         | Promoted to parent | Inlined |
| Destination.HybridConnection.Properties | Promoted to parent | Inlined |
| Destination.ServiceBusQueue.Properties  | Promoted to parent | Inlined |
| Destination.ServiceBusTopic.Properties  | Promoted to parent | Inlined |
| Destination.StorageQueue.Properties     | Promoted to parent | Inlined |
| Destination.WebHook.Properties          | Promoted to parent | Inlined |
| Location                                | Removed            | Unused  |
| Tags                                    | Removed            | Unused  |

## EventHub v2021-11-01

### Namespace

| Property                   | Change  | Reason      |
| -------------------------- | ------- | ----------- |
| PrivateEndpointConnections | Removed | Status Only |

### NamespacesAuthorizationRule

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |

### NamespacesEventhub

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |

### NamespacesEventhubsAuthorizationRule

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |

### NamespacesEventhubsConsumerGroup

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |

## MachineLearningServices v2021-07-01

### Workspace

| Property Path                   | Change  | Reason      |
| ------------------------------- | ------- | ----------- |
| Identity.UserAssignedIdentities | Removed | Status Only |

### WorkspacesCompute

| Property Path                   | Change  | Reason      |
| ------------------------------- | ------- | ----------- |
| Identity.UserAssignedIdentities | Removed | Status Only |

### WorkspacesConnection

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |


## Network v2020-11-01

### LoadBalancer

| Property Path                | Change  | Reason |
| :--------------------------- | :------ | :----- |
| BackendAddressPools.Location | Removed | Unused |

### NetworkSecurityGroupsSecurityRule

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |

### VirtualNetworkGateway

| Property                       | Change                      | Reason        |
| :----------------------------- | :-------------------------- | :------------ |
| VirtualNetworkExtendedLocation | Renamed to ExtendedLocation | Discriminator |

## Operational Insights v2021-06-01

### Workspace

| Property Path                | Change     | Reason     |
| :--------------------------- | :--------- | :--------- |
| Sku.CapacityReservationLevel | Restricted | Validation |

## ServiceBus v2021-10-01

### NamespacesQueue

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |

### NamespacesTopic

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |


### NamespacesTopicsSubscription

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |

### NamespacesTopicsSubscriptionsRule

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |

## SignalRService v2021-11-01

| Property Path                   | Change  | Reason      |
| ------------------------------- | ------- | ----------- |
| Identity.UserAssignedIdentities | Removed | Status Only |


## Storage v2021-04-01

### StorageAccountsBlobService

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |

### StorageAccountsBlobServicesContainer

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |

### StorageAccountsManagementPolicy

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Tags     | Removed | Unused |

## StorageAccountsQueueService

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |

## StorageAccountsQueueServicesQueue

| Property | Change  | Reason |
| :------- | :------ | :----- |
| Location | Removed | Unused |
| Tags     | Removed | Unused |


## Web v2022-03-1

### Site

| Property Path                   | Change  | Reason Only |
| ------------------------------- | ------- | ----------- |
| Identity.UserAssignedIdentities | Removed | Status Only |
