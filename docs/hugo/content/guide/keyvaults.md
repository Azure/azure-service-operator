---
title: "KeyVaults"
linktitle: "KeyVaults"
weight: 1 # This is the default weight if you just want to be ordered alphabetically
---

The standard options for `createMode` when creating a new KeyVault assume you know whether there is an existing soft-deleted KeyVault already present.

* `default` - create a new KeyVault. Deployment will fail if an existing soft-deleted KeyVault has the same name.
* `recover` - recover an existing soft-deleted KeyVault. Deployment will fail if no soft-deleted KeyVault is found.

These options works perfectly well when deploying via ARM or Bicep as those deployments are single use.

However, with a goal seeking system that expects to maintain resources over time, these options don't work so well (see _Motivation_, below), so in Azure Service Operator (ASO) v2.4.0 we are introducing two additional options:

* `createOrRecover` - create a new key KeyVault. If an existing soft-deleted KeyVault has the same name, recover and reuse it.

* `purgeThenCreate` - If an existing soft-deleted KeyVault has the same name, purge it first; then create a brand new one. **Dangerous**

{{% alert title="Warning" color="warning" %}}
Purging a KeyVault is a permanent action from which there is no recovery.
{{% /alert %}}

## Motivation

One of the core value propositions of ASO is its active management of the deployed resources. If they drift from the desired configuration, ASO will act to bring them back into the required state.

KeyVaults are a key piece of infrastructure, containing secrets and credentials required for an application to operate. If a KeyVault is deleted (whether accidentally, or maliciously), it's desirable for ASO to be able to automatically recover so that normal operation of the application can be restored. These new options for `createMode` allow that to happen.

## Recommendations

For a **production** KeyVault (one where you care very much about the contents of the KeyVault being retained), use `createOrRecover`. Consider also setting `enablePurgeProtection` to **true** (preventing a malicious actor from nuking your secrets), and setting `enableSoftDelete` to **true** (while this is the default, signalling your intent may be desirable).

For a **test** or **staging** KeyVault (one where you're testing deployment and do not care about preserving the secrets within the KeyVault), use `purgeThenCreate`. For the purge to work, you'll need to grant ASO additional permissions.

{{% alert title="Warning" color="warning" %}}
Configuring an ASO KeyVault with `purgeThenCreate` will result in any _soft-delete_ action being swiftly promoted to a non-recoverable _hard-delete_. Do not do this with any KeyVault containing valuable secrets.
{{% /alert %}}


## Further Reading

* [Azure Key Vault soft-delete overview](https://learn.microsoft.com/en-us/azure/key-vault/general/soft-delete-overview)
