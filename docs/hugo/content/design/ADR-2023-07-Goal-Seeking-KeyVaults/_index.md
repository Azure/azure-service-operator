---
title: 2023-07 Goal Seeking KeyVaults
---

## Context

Azure KeyVaults are soft deleted, to protect against accidental deletion. If you delete a KeyVault, you can't create a new KeyVault with the same name unless you explicit purge or recover the KeyVault.

In the context of an ARM request to create the KeyVault, where a request is made and that request either succeeds or fails, this behaviour makes perfect sense.

However, in the context of a goal seeking system like the Azure Service Operator (ASO), this is a problem. It's undesirable for resources to enter permanently degraded states that can't be recovered by the system. It's also undesirable for users to have to manually intervene to recover from such a state.

### Known Scenarios

Known scenarios where this is a problem include:

#### Production KeyVault Deletion

A KeyVault being managed by ASO is deleted by an outside agent, impacting applications dependent on the KeyVault. Functionality may be simply degraded (with applications continuing to use secrets previously retrieved), or entirely broken (if the application is restarted and attempts to retrieve secrets from the KeyVault).

The KeyVault may have been removed by a user in error, by an automated process, or by a malicious actor.

Currently, the resource will end up in an unready state that ASO cannot automatically resolve. When ASO attempts to reconcile the resource, the reconciliation will fail with an error that the requested KeyVault name is unavailable, and the resource will remain in an unready state.

Ideally, ASO would be able to automatically take the appropriate remedial action, restoring the functionality of the application running on the cluster.

#### Staging Deployment

It's typical for an application in development to be deployed into a staging environment for testing. In order to test the deployment process, this is often done by completely deleting the prior release and deploying everything from scratch.

If cleanup of the prior release into Staging has soft-deleted the KeyVault, the new deployment will fail with an error that the requested KeyVault name is unavailable.

As a consequence, users currently need to separately purge or recover the KeyVault before the new deployment will successfully provision the resources required by the application. They may do this manually, or by running a separate script.

Ideally, staging deployments driven by ASO would work end to end without additional intervention.

### Current Properties

The [KeyVault](https://azure.github.io/azure-service-operator/reference/keyvault/v1api20210401preview/#keyvault.azure.com/v1api20210401preview.VaultProperties) resource has the following properties relevant to this discussion.

#### createMode

The vault’s create mode to indicate whether the vault need to be recovered or not.

If set to `default` (or if not specified), KeyVault creation will only succeed if there is no existing soft-deleted KeyVault with the same name. If there is such a soft-deleted KeyVault, creation of the KeyVault will fail.

If set to `recover`, KeyVault creation will only succeed if there is an existing soft-deleted KeyVault with the same name. If there is no such KeyVault, creation of the KeyVault will fail.

There is no `purge` option, so it's not possible to purge a KeyVault and create a replacement in a single operation. We can reasonably infer that the Product Group did this deliberately, to prevent accidental purging of a KeyVault.

#### enablePurgeProtection

Specify whether protection against purge is enabled for this vault. 

Setting this property to **true** activates protection against purge for this vault and its content - only the Key Vault service may initiate a hard, irrecoverable deletion. The setting is effective only if soft delete is also enabled. Enabling this functionality is irreversible - that is, the property does not accept **false** as its value.
  
#### enableSoftDelete

Specify whether the ‘soft delete’ functionality is enabled for this key vault. 

If it’s not set to any value (**true** or **false**) when creating new key vault, it will be set to true by default. Once set to true, it cannot be reverted to false.

### Option 1: Reinterpretation of createMode

We reinterpret/redefine the permitted values of `createMode` to reflect the goal seeking nature of ASO:

`block` - Create a new KeyVault if there is no existing KeyVault with the same name. If there is an existing KeyVault with the same name, block creation. (No change to current behaviour.)

`recover` - Create a new KeyVault if there is no existing KeyVault with the same name. If there is an existing KeyVault with the same name, recover it. 

The default value would remain be `block`, maintaining the existing behaviour and requiring users to opt-in if they want automatic recovery.

* PRO: Allows users to opt-in to the new behaviour, and to continue to use the existing behaviour if they prefer.
* PRO: Can be achieved by using our existing supported extension points.
* CON: Changing the meaning of an existing property might be confusing.

### Option 2: Reterpretation with Purge

As for **Option 1**, but we include an option to purge the KeyVault if it exists.

* PRO: Allows users to configure all possible behaviours.
* CON: Dangerous, as we could accidentally purge a KeyVault that we weren't intended to purge.
* CON: The KeyVault Program Group doesn't provide this, likely for good reason.

If someone soft-deleted a KeyVault (whether maliciously or accidentally), automatically purging that KeyVault would make any existing secrets unrecoverable, compounding any problems. To suggest *this would be a bad thing* would be somewhat of an understatement.

### Option 3: Sidestep the Problem

If we attempt to create a KeyVault and encounter a name collision, we could automatically select a different AzureName for the KeyVault and publish that as a ConfigMap. The application could then be configured to use the new name.

* CON: May be surprising to users
* CON: If the KeyVault was deleted by a malicious actor, we'd respond by starting afresh with a blank KeyVault, resulting in a longer outage than necessary.
* CON: Only applications running in the cluster would be able to use the new KeyVault name. Applications running outside the cluster would continue to use the old name.
* CON: Requires that applications running in the cluster load their KeyVault name from a configmap, rather than hard coding it via other config.


### Option 4: Automatic Recovery

Remove the `createMode` property entirely and automatically recover any soft-deleted KeyVaults that we encounter.

This is a natural extension of ASO as a maintainer of the declared goal state. If the goal state is to have a KeyVault, and there is a soft-deleted KeyVault, then a natural action is to recover the soft-deleted KeyVault.

* PRO: We can remove the property using existing generator config (though it would be a breaking change.)
* PRO: We can implement the behaviour using existing extension points.
* PRO: Likely to be the behaviour wanted by most users.
* CON: No way to opt-out

### Option 5: New Property

Remove the existing `createMode` property and create a new property to configure the behaviour, say `reconciliationMode`:

`block` - Create a new KeyVault if there is no existing KeyVault with the same name. If there is an existing KeyVault with the same name, block creation.

`recover` - Create a new KeyVault if there is no existing KeyVault with the same name. If there is an existing KeyVault with the same name, recover it. 

We could specify `recover` as the default.

* PRO: We're not redefining `createMode` so less possiblity for confusion there.
* PRO: A new property can have exactly the semantics we want.
* CON: Users may be surprised that `createMode` is not supported.
* CON: We don't have a general purpose technique for introducing new custom properties and would need to add one.


### Option 6: New Property with Purge

As for **Option 6**, but we include an option to purge the KeyVault if it exists.

* PRO: Allows users to configure all possible behaviours.
* CON: Dangerous, as we could accidentally purge a KeyVault that we weren't intended to purge.
* CON: The KeyVault Program Group doesn't provide this, likely for good reason.

If someone soft-deleted a KeyVault (whether maliciously or accidentally), automatically purging that KeyVault would make any existing secrets unrecoverable, compounding any problems. To suggest *this would be a bad thing* would be somewhat of an understatement.


## Decision

TBC

## Status

TBC

## Consequences

TBC

## Experience Report

TBC

## References

TBC

