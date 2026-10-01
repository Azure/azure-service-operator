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


## VaultKey

`VaultKey` (`Microsoft.KeyVault/vaults/keys`) manages a cryptographic key inside an existing
`Vault`.

### Generation-only, by design

Azure Key Vault always generates the key material itself when a `VaultKey` is created; there is
no field on this resource for supplying key material, because the underlying ARM API has none. As
a result:

* This resource **cannot be used to import key material**. Although ASO uses the Key Vault data
  plane to update and delete keys (see below), it never imports anything, and there is no field on
  this resource through which key material could be supplied.
* `spec.properties.keyOps` may not include `import`; the validating webhook rejects it. A key whose
  operation is `import` is a *key exchange key*: its only purpose is to wrap other key material
  during a bring-your-own-key (BYOK) transfer into the vault, a flow this resource does not
  support. (Whether an identity may import material into a vault is governed by that identity's
  `import` permission, not by this key operation, so this is a scope decision rather than a
  security control.)
* Private key material never transits the Kubernetes cluster, the `VaultKey` spec, or any
  Kubernetes `Secret`.

### Adopting an existing key

If a live key with the requested name already exists in the vault, the `VaultKey` resource adopts
it and manages it from then on - but only when the generation-time properties in the spec (`kty`,
`keySize`, `curveName`) match the real key. A key's type, size and curve are fixed when its
material is generated, so a spec that disagrees with them describes a *different* key; rather than
reporting `Ready` for a key that doesn't match its spec, ASO blocks the reconcile with a warning
condition explaining the mismatch and retries periodically. By then the resource's status has been
populated from the existing key, which locks the generation-time properties in the spec, so the
remedies are to delete and recreate the resource with matching properties (deleting it leaves a key
it never adopted untouched, whatever `deleteMode` says), or to remove the key in Azure.
Properties the spec leaves unset match anything.

Keys that back a certificate are managed by Key Vault itself and reject updates, so they cannot be
adopted; the same warning condition explains why, and the remedy is to remove the certificate that
owns the key or to recreate the resource with a different name.

Because adoption is by name, ASO cannot guarantee that every key managed by a `VaultKey` resource
was actually generated by ASO - it may have pre-existed in the vault. If you need provenance
assurance, restrict who/what is permitted to create keys in the target vault using Azure RBAC.

### Updates flow through the data plane

The ARM API for `Microsoft.KeyVault/vaults/keys` supports only create-if-not-exist and read, so
ASO applies changes to the mutable key properties through the Key Vault **data plane** after each
reconcile: `keyOps`, `attributes` (`enabled`, `exp`, `nbf`), `release_policy`, `rotationPolicy`
and `tags`. Only properties the spec actually sets are managed; anything the spec leaves unset
keeps its value in Azure. An empty `keyOps` list or `tags` map counts as unset, so there is no way
to clear them through the spec.

`rotationPolicy` is managed as a whole once the spec sets it: lifetime actions removed from the
spec are removed from the key. The notify action Key Vault adds by itself to a policy that
declares none is left alone.

Two things remain locked down by the validating webhook:

* `kty`, `keySize` and `curveName` are immutable once the resource's status reflects a key in
  Azure - they are fixed at key generation time under any mechanism, ARM or data plane. Delete and
  recreate the resource to change them.
* `spec.properties.attributes.exportable: true` and the `import` keyOp are unconditionally
  rejected, preserving the generation-only contract above.

A release policy that Key Vault has marked immutable cannot be changed by anyone; if the spec
disagrees with it, ASO reports a `Ready` condition with severity `Error` until the spec matches.

Under a `serviceoperator.azure.com/reconcile-policy` of `skip`, ASO writes nothing to the key.

A data-plane update made by ASO becomes visible in the resource's `status` on the *following*
reconcile, since status is refreshed just before the update is applied.

### createMode

Like `Vault` itself (see above), `VaultKey` supports goal-seeking around name collisions with
**soft-deleted** keys via `spec.operatorSpec.createMode`:

* `default` - create the key; if a soft-deleted key blocks the name, the Azure error is surfaced
  on the resource.
* `recover` - recover the soft-deleted key; the reconcile is blocked when no soft-deleted key
  with that name exists.
* `createOrRecover` - recover a soft-deleted key when one exists, otherwise create a new key.
* `purgeThenCreate` - **permanently purge** the soft-deleted key, then generate a brand-new one.
  **Dangerous**: the old key material is unrecoverable, and anything encrypted under it can no
  longer be decrypted. A vault with purge protection refuses the purge.

ASO refuses to recover a soft-deleted key whose `kty`, `keySize` or `curveName` don't match the
spec - recovery would resurrect key material the spec doesn't describe. The reconcile is blocked
with a warning until the spec is corrected (no live key has been observed, so its generation-time
properties are not yet locked), the soft-deleted key is purged, or `purgeThenCreate` is chosen.

### Deletion

What happens to the key in Azure when the `VaultKey` resource is deleted from Kubernetes is
controlled by `spec.operatorSpec.deleteMode`:

* `detach` (the **default**) - the key is left untouched in the vault; only the Kubernetes
  resource goes away. ASO never destroys key material unless explicitly asked to.
* `disable` - the `enabled` attribute of the key's current version is set to `false` and the key
  is left in the vault.
* `delete` - the key is **soft-deleted** via the data plane. It remains recoverable for the
  vault's soft-delete retention period, after which the service purges it; purge protection only
  prevents it from being purged *early*. ASO never purges a key on deletion.

A key this resource never adopted - because its generation-time properties differ from the spec,
or because it backs a certificate - is left untouched whatever `deleteMode` says.

If Key Vault refuses the read that precedes it, or the `disable` or `delete` operation itself -
usually because ASO's identity lacks
the permission for it, or because the vault's firewall or private endpoint blocks ASO - deletion is
blocked and the resource's `Ready` condition names the permission the mode needs. Grant access, or
change `deleteMode` to `detach`, to let the resource go.

The standard `serviceoperator.azure.com/reconcile-policy` annotation still takes precedence: if
`detach-on-delete` or `skip` is in effect - from the object itself, from an annotation on its
**namespace**, or operator-wide via the `DEFAULT_RECONCILE_POLICY` setting (see
[annotations]( {{< relref "annotations" >}} ) and
[configuration options]( {{< relref "aso-controller-settings-options" >}} )) - ASO bypasses
deletion handling entirely and leaves the key untouched, regardless of `deleteMode`. If you rely
on `deleteMode: delete` or `disable`, make sure those resources live under the default `manage`
policy.

### Required permissions for ASO's identity

`VaultKey` uses both planes, so ASO's identity needs control-plane *and* data-plane access - but
only to keys, and only for the operations you actually use. Grant exactly what's needed and
nothing more.

**Control plane** (Azure RBAC `Actions`), always required:

* `Microsoft.KeyVault/vaults/read` - to look up the vault's data-plane URI before the key exists
* `Microsoft.KeyVault/vaults/keys/read`
* `Microsoft.KeyVault/vaults/keys/write`

**Data plane**: for vaults using Azure RBAC (`enableRbacAuthorization: true`), grant
`DataActions`; for vaults using access policies, grant the equivalent key permissions:

| ASO operation                                         | RBAC `DataActions`                                                                          | Access-policy key permissions            |
|-------------------------------------------------------|---------------------------------------------------------------------------------------------|------------------------------------------|
| Read / adopt / verify (always required)               | `Microsoft.KeyVault/vaults/keys/read`                                                        | `get`                                    |
| Update mutable properties, `deleteMode: disable`      | `Microsoft.KeyVault/vaults/keys/update/action`                                               | `update`                                 |
| Manage `rotationPolicy` (if spec sets it)             | `Microsoft.KeyVault/vaults/keyrotationpolicies/read` and `.../keyrotationpolicies/write`     | `getrotationpolicy`, `setrotationpolicy` |
| `deleteMode: delete`                                  | `Microsoft.KeyVault/vaults/keys/delete`                                                      | `delete`                                 |
| `createMode: recover`/`createOrRecover`               | `Microsoft.KeyVault/vaults/keys/recover/action`                                              | `recover`                                |
| `createMode: purgeThenCreate`                         | `Microsoft.KeyVault/vaults/keys/purge/action`                                                | `purge`                                  |

Looking for a soft-deleted key (any `createMode` other than `default`) is covered by the read
permission.

Do **not** grant import, export, release, backup/restore, or cryptographic operation
(encrypt/decrypt/sign/verify/wrap/unwrap) permissions - ASO never needs them, and granting them
expands the blast radius of a compromised ASO identity beyond what's required.

The built-in "Key Vault Crypto Officer" role includes far more than the rows above (including
purge and all cryptographic operations) and should be avoided for ASO's identity; define a custom
role restricted to the operations you use instead.

## Further Reading

* [Azure Key Vault soft-delete overview](https://learn.microsoft.com/en-us/azure/key-vault/general/soft-delete-overview)
* [Configure key rotation in Azure Key Vault](https://learn.microsoft.com/en-us/azure/key-vault/keys/how-to-configure-key-rotation)
* [Azure Key Vault data-plane permissions](https://learn.microsoft.com/en-us/azure/role-based-access-control/permissions/security#microsoftkeyvault)
