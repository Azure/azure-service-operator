# Breaking Changes in beta.4

In the 'beta.4' release of Azure Service Operator (ASO) we are pivoting to using the Azure Swagger API Specifications as our sole source of truth for our code generator.

This change brings with in a significant improvement in fidelity - the code we generate is now much closer to what the Azure Swagger API Specifications describe.

Unfortunately, this change brings with it a few breaking changes, listed here.

These changes fall into a few categories:

**Unused**: Properties that were previously generated but actually had no function have been removed.

Examples include:

* `Location` on subresources that inherit their actual location from their parent resource
* `Tags` on resources that don't support tags

If your resource Spec was using one of these properties, no action is required; the new release of ASO will silently ignore them.

**Status Only**: Properties that cannot set by the end user on a resource spec. Typically these are actually Status properties included in the Spec in error.

Examples include:

* Identity.UserAssignedIdentities

If your resource Spec was using one of these properties, no action is required; the new release of ASO will silently ignore them.

**Subresources**: In some cases we weren't correctly identifying sub-resources and breaking them out into separate CRDs, leaving those properties inline in the parent resource.

Examples include:

* VirtualNetworkGateway.VirtualNetworkExtendedLocation

If your resource was inlining one of these, you'll need to modify the YAML for your resource to replace the existing properties with a reference to a new resource, and to list that resource elsewhere in the same YAML file.

**Enumerations**: In a few cases we previously had only a base type for a property, without any enumeration values to restrict, even though there were only a few legitimate values to provide. The allowed Kubernetes resource could be created with any value at all even though creation of the Azure resource would fail.

If your resource was using an invalid value for one of these properties, you'll need to modify the YAML for your resource to use a valid value. 

Note that any resource successfully created in Azure will be fine.

**Fixed value fields**: Fields that required a specific value to be provided, with no option; these are similar to Status Only fields.

**Inlined objects**: We're now correctly inlining some objects that were previously generated as nested properties.

If your resource is using one of these, you'll need to modify your resource to promote the property to the parent object.

**Discriminator use**: We now use the discriminator value for polymorphic types to name the property used to specify that option. Previously we were synthesizing a name based on other factors, resulting in longer property names that did not appear in the ARM/Bicep documentation.

If your resource is using one of these, you'll need to change the name of the parent property as detailed below.

## Authorization

### RoleAssignment

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |


## Cache v2020-12-01

### RedisFirewallRule

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |

### RedisLinkedServer

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |

### RedisPatchSchedule

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |

## Cache v2021-03-01

### RedisEnterpriseDatabase

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |


## CDN v2021-06-01

### ProfilesEndpoint

| Property                                                            | Change (Reason)                                       | Impact           |
| ------------------------------------------------------------------- | ----------------------------------------------------- | ---------------- |
| DeliveryPolicy.Rules.Actions.DeliveryRuleCacheExpiration            | Renamed to CacheExpiration (discriminator)            | Update your YAML |
| DeliveryPolicy.Rules.Actions.DeliveryRuleCacheKeyQueryString        | Renamed to CacheKeyQueryString (discriminator)        | Update your YAML |
| DeliveryPolicy.Rules.Actions.DeliveryRuleRequestHeader              | Renamed to ModifyRequestHeader (discriminator)        | Update your YAML |
| DeliveryPolicy.Rules.Actions.DeliveryRuleResponseHeader             | Renamed to ModifyResponseHeader (discriminator)       | Update your YAML |
| DeliveryPolicy.Rules.Actions.DeliveryRuleRouteConfigurationOverride | Renamed to RouteConfigurationOverride (discriminator) | Update your YAML |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleClientPort              | Renamed to ClientPort (discriminator)                 | Update your YAML |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleCookies                 | Renamed to Cookies (discriminator)                    | Update your YAML |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleHostName                | Renamed to HostName (discriminator)                   | Update your YAML |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleHttpVersion             | Renamed to HttpVersion (discriminator)                | Update your YAML |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleIsDevice                | Renamed to IsDevice (discriminator)                   | Update your YAML |
| DeliveryPolicy.Rules.Conditions.DeliveryRulePostArgs                | Renamed to PostArgs (discriminator)                   | Update your YAML |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleQueryString             | Renamed to QueryString (discriminator)                | Update your YAML |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleRemoteAddress           | Renamed to RemoteAddress (discriminator)              | Update your YAML |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleRequestBody             | Renamed to RequestBody (discriminator)                | Update your YAML |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleRequestHeader           | Renamed to RequestHeader (discriminator)              | Update your YAML |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleRequestMethod           | Renamed to RequestMethod (discriminator)              | Update your YAML |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleRequestScheme           | Renamed to RequestScheme (discriminator)              | Update your YAML |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleRequestUri              | Renamed to RequestUri (discriminator)                 | Update your YAML |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleServerPort              | Renamed to ServerPort (discriminator)                 | Update your YAML |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleSslProtocol             | Renamed to SslProtocol (discriminator)                | Update your YAML |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleUrlFileExtension        | Renamed to UrlFileExtension (discriminator)           | Update your YAML |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleUrlFileName             | Renamed to UrlFileName (discriminator)                | Update your YAML |
| DeliveryPolicy.Rules.Conditions.DeliveryRuleUrlPath                 | Renamed to UrlPath (discriminator)                    | Update your YAML |


## Compute v2020-12-01

### VirtualMachineScaleSet

| Property                                                 | Change (Reason) | Impact |
| -------------------------------------------------------- | --------------- | ------ |
| VirtualMachineProfile.NetworkProfile.Id                  | Removed (??)    | ??     |
| VirtualMachineProfile.NetworkProfile.IpConfigurations.Id | Removed (??)    | ??     |


## Container Service

### ManagedCluster

| Property                        | Change (Reason)       | Impact |
| ------------------------------- | --------------------- | ------ |
| Identity.UserAssignedIdentities | Removed (Status Only) | None   |

### ManagedClustersAgentPool

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Tags     | Removed (Unused) | None   |

## DbForMariaDB v2018-06-01

### Configuration

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |

### Server

| Property                                    | Change (Reason)                               | Impact                                |
| :------------------------------------------ | :-------------------------------------------- | :------------------------------------ |
| Properties.ServerPropertiesForDefaultCreate | Renamed to Default (discriminator)            | Update your YAML to use the new name. |
| Properties.ServerPropertiesForGeoRestore    | Renamed to GeoRestore (discriminator)         | Update your YAML to use the new name. |
| Properties.ServerPropertiesForReplica       | Renamed to Replica (discriminator)            | Update your YAML to use the new name. |
| Properties.ServerPropertiesForRestore       | Renamed to PointInTimeRestore (discriminator) | Update your YAML to use the new name. |

### Database

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |


## DbForMySql v2021-05-01

### FlexibleServersDatabase

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |

### FlexibleServersFirewallRule

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |

## DbForPostgreSQL v2021-06-01

### FlexibleServersConfiguration

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |

### FlexibleServersDatabase

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |

### FlexibleServersFirewallRule

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |

## EventGrid v2020-06-01

### Domain

| Property                                  | Change (Reason)       | Impact           |
| ----------------------------------------- | --------------------- | ---------------- |
| InputSchemaMapping.InputSchemaMappingType | Removed (Fixed value) | Remove from YAML |
| InputSchemaMapping.Properties             | Renamed to Json (??)  | Rename in YAML   |

### DomainsTopic

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |

### EventSubscription

| Property                                | Change (Reason)   | Impact      |
| :-------------------------------------- | :---------------- | :---------- |
| DeadLetterDestination.Properties        | Removed (Inlined) | Modify YAML |
| Destination.AzureFunction.Properties    | Removed (Inlined) | Modify YAML |
| Destination.EventHub.Properties         | Removed (Inlined) | Modify YAML |
| Destination.HybridConnection.Properties | Removed (Inlined) | Modify YAML |
| Destination.ServiceBusQueue.Properties  | Removed (Inlined) | Modify YAML |
| Destination.ServiceBusTopic.Properties  | Removed (Inlined) | Modify YAML |
| Destination.StorageQueue.Properties     | Removed (Inlined) | Modify YAML |
| Destination.WebHook.Properties          | Removed (Inlined) | Modify YAML |
| Location                                | Removed (Unused)  | None        |
| Tags                                    | Removed (Unused)  | None        |

## EventHub v2021-11-01

### Namespace

| Property                   | Change (Reason)       | Impact |
| -------------------------- | --------------------- | ------ |
| PrivateEndpointConnections | Removed (Status only) | None   |

### NamespacesAuthorizationRule

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |

### NamespacesEventhub

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |

### NamespacesEventhubsAuthorizationRule

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |

### NamespacesEventhubsConsumerGroup

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |


## MachineLearningServices v2021-07-01

### Workspace

| Property                        | Change (Reason)       | Impact |
| ------------------------------- | --------------------- | ------ |
| Identity.UserAssignedIdentities | Removed (Status Only) | None   |

### WorkspacesCompute

| Property                        | Change (Reason)       | Impact |
| ------------------------------- | --------------------- | ------ |
| Identity.UserAssignedIdentities | Removed (Status Only) | None   |

### WorkspacesConnection

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |


## Network v2018-09-01

### PrivateDnsZone

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |

## Network v2020-11-01

### LoadBalancer

| Property                     | Change (Reason)  | Impact |
| :--------------------------- | :--------------- | :----- |
| BackendAddressPools.Location | Removed (Unused) | None   |

### NetworkSecurityGroupsSecurityRule

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |

### VirtualNetworkGateway

| Property                       | Change (Reason)           | Impact |
| :----------------------------- | :------------------------ | :----- |
| GatewayType                    | Option 'HyperNet` removed | #2631  |
| VirtualNetworkExtendedLocation | Removed (Subresource)     | ??     |

## Operational Insights v2021-06-01

### Workspace

| Property                      | Change (Reason) | Impact |
| :---------------------------- | :-------------- | :----- |
| Features.AdditionalProperties | ??              |        |

## ServiceBus v2021-10-01

### NamespacesQueue

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |

### NamespacesTopic

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |

### NamespacesTopicsSubscription

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |

### NamespacesTopicsSubscriptionsRule

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |

### VirtualNetworksSubnet

| Property                           | Change (Reason)            | Impact |
| :--------------------------------- | :------------------------- | :----- |
| ApplicationGatewayIpConfigurations | Removed (Status Only)      | None   |
| PrivateEndpointNetworkPolicies     | Type changed (Enumeration) | None   |
| PrivateLinkServiceNetworkPolicies  | Type changed (Enumeration) | None   |

### VirtualNetworksVirtualNetworkPeering

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |


## SignalRService v2021-11-01

| Property                        | Change (Reason)       | Impact |
| ------------------------------- | --------------------- | ------ |
| Identity.UserAssignedIdentities | Removed (Status Only) | None   |


## Storage v2021-04-01

### StorageAccountsBlobService

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |


### StorageAccountsManagementPolicy

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Tags     | Removed (Unused) | None   |

StorageAccountsQueueService

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |

StorageAccountsQueueServicesQueue

| Property | Change (Reason)  | Impact |
| :------- | :--------------- | :----- |
| Location | Removed (Unused) | None   |
| Tags     | Removed (Unused) | None   |


## Web v2022-03-1

### Site

| Property                        | Change (Reason)       | Impact |
| ------------------------------- | --------------------- | ------ |
| Identity.UserAssignedIdentities | Removed (Status Only) | None   |
