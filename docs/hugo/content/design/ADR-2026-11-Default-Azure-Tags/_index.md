---
title: '2026-11: Default Azure Tags'
toc_hide: true
---

## Context

ASO records ownership only on the Kubernetes side, using the `serviceoperator.azure.com/*` annotations and labels described in the [annotations guide]({{< relref "annotations" >}}). Nothing is written to the Azure resource itself. If someone finds a resource in the portal, in Azure Resource Graph, in Cost Management, or in an audit log, they can't tell that ASO manages it, or which cluster, namespace, and custom resource it came from. To find out, they have to search Kubernetes state on every cluster that might own it.

[#5713](https://github.com/Azure/azure-service-operator/issues/5713) asks for an opt-in way to add owner metadata as tags on the Azure resources ASO manages. Users want this for three things:

* **Cost attribution**: splitting Azure spend by cluster, team, or namespace.
* **Audit and compliance**: finding the owner of a resource that appears in an access review or a security finding.
* **Incident response**: tracing a misbehaving or leaked resource back to the Kubernetes object that owns it.

Users can already set `spec.tags` on each resource, and Kustomize or Helm can add static tags. These tools can't easily add values that depend on where the resource ends up, such as its namespace or name, and every manifest would need the same boilerplate.

### Prior art

* [AWS Controllers for Kubernetes](https://aws-controllers-k8s.github.io/community/docs/user-docs/ack-tags/) adds `services.k8s.aws/controller-version` and `services.k8s.aws/namespace` by default. The set can be configured through `resourceTags`, and `Spec.Tags` takes precedence.
* [Config Connector](https://cloud.google.com/config-connector/docs/concepts/resources) adds a `managed-by-cnrm` label and copies `metadata.labels` onto the GCP resource.
* Crossplane's Azure providers add `crossplane-kind`, `crossplane-name` and `crossplane-providerconfig` tags.

The Kubernetes-style keys used by AWS Controllers for Kubernetes contain `/`, which Azure doesn't allow in tag names (see below).

### Relevant ASO behavior

* ASO builds the request it sends to Azure in `ConvertToARMResourceImpl`, by calling the generated `ConvertToARM()` on the spec. Resource extensions can then adjust the request through `extensions.ARMResourceModifier` before it's sent with a `PUT`.
* ASO sends a full `PUT` of the desired state on each reconcile (see [Change Detection]({{< relref "ADR-2022-11-Change-Detection" >}})). Most resource providers remove tags that aren't in the request and reapply the tags that are.
* Most resources expose tags as a top-level `Tags map[string]string` property, but not all of them do:
  * Some resources don't support tags at all.
  * Some have a property named `tags` that isn't an Azure tag, for example `[]string` or `*string`.
* Some resources don't go through the normal `PUT` path at all:
  * The `azuresql`, `mysql` and `postgresql` reconcilers manage database users through the database's own protocol.
  * The `entra` reconciler manages Entra ID objects, which don't have Azure tags.
* KeyVault keys use the normal `PUT` path only to create a key. Azure ignores that `PUT` once the key exists, so after each successful reconcile an extension updates the key through the Key Vault data-plane `UpdateKey` API ([vault_key_extensions.go](https://github.com/Azure/azure-service-operator/blob/main/v2/api/keyvault/customizations/vault_key_extensions.go)). When `spec.tags` isn't empty and differs from the key's tags in Azure, the extension replaces the key's tags with `spec.tags`. Otherwise it leaves them alone.
* ASO already has a CEL evaluator (`v2/internal/util/cel`) used for `operatorSpec.configMapExpressions` and `secretExpressions`. It includes program caching and metrics.
* Operator-wide settings are read from environment variables populated from the `aso-controller-settings` secret, and are checked at startup by `config.Values.Validate()`. An invalid value stops the operator from starting.
* `AZURE_USER_AGENT_SUFFIX` already lets an installation identify itself in Azure request logs.

### Azure tag constraints

These constraints come from [Use tags to organize your Azure resources](https://learn.microsoft.com/azure/azure-resource-manager/management/tag-resources): the limits from its [Limitations](https://learn.microsoft.com/azure/azure-resource-manager/management/tag-resources#limitations) section, and the casing and plain-text notes from [Tag usage and recommendations](https://learn.microsoft.com/azure/azure-resource-manager/management/tag-resources#tag-usage-and-recommendations):

* Tag names can't contain `<`, `>`, `%`, `&`, `\`, `?` or `/`.
* Some services add their own restrictions on tag names:
  * Azure DNS zones don't support spaces, parentheses, special characters or Unicode characters, or names that start with a digit.
  * Traffic Manager doesn't support spaces, `#` or `:`, or names that start with a digit.
  * Azure Front Door doesn't support `#` or `:`.
* Tag names can be up to 512 characters, but only 128 on storage accounts. Tag values can be up to 256 characters.
* Most resources can have up to 50 tags. Azure Automation, Azure CDN, Azure Public DNS and Azure Private DNS (zones and A records), and Log Analytics saved searches support only **15**.
* Tag names are case-insensitive for operations, but a resource provider may keep the casing it was given, and cost reports show that casing. Tag values are case-sensitive.
* Tags are stored as plain text and appear in cost reports, deployment history, exported templates and monitoring logs. Azure warns against putting sensitive values in tags, and also warns that non-English tag values can break virtual machine metadata decoding through the Instance Metadata Service (IMDS).
* Some resource types have their own rules about when tags can be updated. For example, VM extension tags can only be updated while the VM is running.

These are Azure's published values at the time of writing, and they may change. The implementation will keep the ones ASO relies on as named constants (see [Validation](#validation)).

### Azure features relevant to this design

* **`cm-resource-parent` tag.** Cost Management [groups related resources](https://learn.microsoft.com/azure/cost-management-billing/costs/group-filter#group-related-resources-in-the-resources-view) under a parent when a resource has a `cm-resource-parent` tag whose value is the parent's resource ID. ASO already works out each resource's owner hierarchy, so it can supply this value.
* **Tags API.** [`Tags - Update At Scope`](https://learn.microsoft.com/rest/api/resources/tags/update-at-scope) (`PATCH {scope}/providers/Microsoft.Resources/tags/default`) supports `Merge`, `Replace` and `Delete` operations without sending the whole resource. `Delete` can remove tags by *name/value pair*, so it removes a tag only if it still has the expected value. The caller needs write access to the resource or the `Tag Contributor` role.
* **Tag inheritance.** [Azure Policy](https://learn.microsoft.com/azure/azure-resource-manager/management/tag-policies) can copy tags from a resource group or subscription onto the resources inside it. [Cost Management tag inheritance](https://learn.microsoft.com/azure/cost-management-billing/costs/enable-tag-inheritance) does the same for cost records only.
* **Azure Service Groups.** Service Groups are a newer container that can group resources across subscriptions by membership. They were [announced in public preview in August 2025](https://microsoft.github.io/azurekorea/Announcing-Public-Preview-for-Azure-Service-Groups/), and their SDKs were still in beta in 2026.

## Requirements

* **Opt-in.** Existing installations see no change, and no resources are modified, unless the feature is configured. After the feature is turned off, ASO only finishes removing what it applied while it was on (see [Update and lifecycle semantics](#update-and-lifecycle-semantics)).
* **Configurable.** Users choose which tags are applied. ASO doesn't hard-code a fixed set, because different users need different information. For example, a `dev` cluster and a `prod` cluster have to be distinguishable.
* **Dynamic values.** Tag values can depend on the Kubernetes resource being reconciled, for example its namespace, name, or owner.
* **User intent wins by default.** A tag that the user sets in `spec.tags` is never overwritten. Users can also turn off individual default tags on a single resource.
* **Safe everywhere.** If a configuration is accepted, the tag names and the number of tags must be valid on every resource type ASO supports.
* **Stable.** If nothing about a resource has changed, reconciling it again must produce the same tags. The feature must not create a change on every reconcile.
* **Clear errors.** Configuration mistakes are reported as early as possible, at a level that matches how much of the operator they affect.
* **No secret leakage.** Expressions must not be able to read secrets or arbitrary spec values.
* **Compatible with future change detection.** The injected tags are part of the desired state, so a future comparison between desired and actual state includes them. An Azure Policy `modify` effect can overwrite a default tag on every write, so change detection must not treat that difference as drift that triggers a `PUT` each time. The comparison of default tags should be configurable (the mechanism is left to the change-detection design), and a tag listed in `skip-default-tags` isn't part of the desired state, so it's never compared.

## Options

### Should ASO do this at all?

#### Option 0: No. Use admission-time mutation

Users can already add `spec.tags` dynamically by mutating resources when they are admitted to the cluster:

* [Kyverno](https://kyverno.io) mutate rules can use variables such as `{{request.object.metadata.namespace}}`.
* Kubernetes' built-in [MutatingAdmissionPolicy](https://kubernetes.io/docs/reference/access-authn-authz/mutating-admission-policy/) uses CEL. It has been beta since Kubernetes 1.34, and [KEP-3962](https://www.kubernetes.dev/resources/keps/3962/) targets stable for 1.36.

This would match the [Templating]({{< relref "ADR-2026-05-Templating" >}}) decision, which leaves this kind of value composition to the wider Kubernetes ecosystem.

**Pros:**

* No change to ASO.
* The tags are visible in `spec`.

**Cons:**

* A policy that adds `spec.tags` to a kind without that field fails or is pruned by the schema. Users would therefore need an allowlist of the hundreds of ASO kinds that have a top-level `tags` field, and would have to keep it up to date with every ASO release.
* Changing the user's `spec` creates a difference that GitOps tools (Argo CD, Flux) report or revert, unless each tool is configured to ignore it.
* Admission policies can't know resource-specific Azure limits (the 15-tag services, DNS naming rules) or ASO's owner hierarchy (for `cm-resource-parent`).
* Mutation only happens on admission, so resources that already exist are tagged only after their next update.
* Every user has to build and maintain the same policies.

### Where the configuration lives

#### Option A: Operator-wide environment variable (JSON)

A new setting in `aso-controller-settings`, exposed as a Helm value, contains a JSON object that maps tag names to expressions.

**Pros:**

* Follows the existing pattern for operator settings. It is documented in one place and checked at startup.
* JSON avoids the escaping problems of `key=value,key=value`, because values can legitimately contain `,` and `=`.
* Helm can render the value from a YAML map with `toJson`.

**Cons:**

* Changing it requires a restart, as with every other setting.
* The value isn't secret but lives in a Secret, like the other operator settings.

#### Option B: A dedicated ConfigMap, watched at runtime

**Pros:**

* Changes take effect without a restart.

**Cons:**

* Adds a new watch, new RBAC rules, and a second configuration mechanism.
* A change has to trigger a re-reconcile, and so a `PUT`, of every resource in the cluster. That load on Azure happens at an unpredictable time instead of at a planned restart.

#### Option C: Per-namespace or per-resource-group configuration

This could live in `operatorSpec`, or be read from namespace labels.

**Pros:**

* Fine-grained control. Platform teams often record a cost center as a label on the namespace.

**Cons:**

* Needs precedence rules between the operator, namespace, owner and resource levels.
* Reading namespace labels means ASO has to watch namespaces. A label change on a namespace would then re-reconcile, and so `PUT`, every resource in that namespace.
* Azure Policy and Cost Management inheritance already cover the resource-group case.

### How values are expressed

#### Option 1: Hard-coded tags with on/off switches

**Pros:**

* Simplest, with nothing for users to get wrong.

**Cons:**

* Doesn't meet the *Configurable* requirement. Two clusters with the same namespace and resource names would produce identical tags.

#### Option 2: Placeholder substitution (for example `${namespace}`)

**Pros:**

* Easy to read.

**Cons:**

* Means inventing, documenting and maintaining a new template language. Users would soon ask for defaults, conditionals and string functions.

#### Option 3: CEL expressions

**Pros:**

* ASO already uses CEL, already documents it for `operatorSpec` expressions, and already has an evaluator with caching and metrics.
* Literals, field access, defaults and string functions all use the same syntax.
* Expressions are type-checked when they're compiled.

**Cons:**

* Literal values need CEL quoting (`"'prod-eastus'"`).
* More powerful than most users need.

### Where tags are injected

#### Option X: Generated interface on ARM spec types

The code generator implements a new interface, for example `genruntime.ARMTaggable` with `GetTags()` and `SetTags()`, on every ARM spec type that has a top-level `Tags map[string]string` property.

**Pros:**

* Driven by the schema and type-safe. Resources that don't support tags, and fields named `tags` that aren't Azure tags, are excluded automatically.
* Extensions with their own update path, such as KeyVault keys, can use the same merge code later. KeyVault keys are excluded in phase 1 (see [Coverage](#coverage)).

**Cons:**

* Requires a change to the generator and a regeneration of the code.
* Resources that keep tags somewhere other than a top-level `tags` property are silently excluded.

#### Option Y: Patch the serialized JSON

**Pros:**

* No generator changes.

**Cons:**

* Can't tell whether a resource supports tags, so some resource providers would reject or silently drop them.

#### Option Z: Tags API for every resource

Use `Tags - Update At Scope` with `Merge` after each successful `PUT`.

**Pros:**

* Works even for resource types whose ASO schema doesn't have a `tags` field.

**Cons:**

* Each reconcile needs a second write.
* The `PUT` (which removes tags that aren't in the request) and the merge then race, so tags briefly disappear and reappear on every reconcile.
* Some resource types, such as IP Groups and Firewall Policies, don't support tag updates through `PATCH` ([limitations](https://learn.microsoft.com/azure/azure-resource-manager/management/tag-resources#limitations)).

### What happens when a resource is detached

#### Option D1: Leave the tags in place

**Pros:**

* No extra calls to Azure.

**Cons:**

* The resource keeps ownership tags that point to an object that no longer exists. That is exactly the misleading information the audit use case is trying to avoid.

#### Option D2: Remove them with the Tags API, using a record of what was applied

After each successful `PUT`, ASO records the default tags it sent in an annotation, `serviceoperator.azure.com/applied-default-tags`. On a delete with `detach-on-delete`, ASO calls `Tags - Update At Scope` with `Delete`, passing the recorded name/value pairs.

**Pros:**

* Removes stale tags, even if the labels or other values the expressions read have changed since the last reconcile, because nothing is recalculated at delete time.
* Only tags that ASO added as defaults are recorded. Tags that the user set in `spec.tags` or opted out of with `skip-default-tags` are never removed, even when their value matches a default.
* Because the delete matches on name *and* value, it doesn't remove a recorded tag whose value has since changed in Azure, for example because of another installation, a user or an Azure Policy `modify` effect.

**Cons:**

* One more Azure call during detach, and one more operator-maintained annotation.
* Can fail on resource types that don't support tag updates through `PATCH`.
* A tag set outside ASO to exactly the recorded name and value can't be told apart from ASO's tag, so it's removed.

## Decision

Don't adopt Option 0. Combine **Option A** (operator-wide setting), **Option 3** (CEL) and **Option X** (generated interface), and remove tags on detach with **Option D2** as a follow-up phase. Option C is deferred, and is listed under future work.

### Core decision

Accepting this ADR means accepting the following. The rest of this section describes the behavior in detail.

* Default tags are opt-in and configured operator-wide (Option A) in `AZURE_DEFAULT_TAGS`, as a JSON map of tag names to CEL expressions (Option 3). The expressions use a fixed environment that excludes `spec`, `status` and secrets.
* Startup validation accepts only tag names and tag counts that are valid on every resource type.
* Tags the user sets in `spec.tags` win, and `skip-default-tags` opts a resource out of individual defaults.
* Default tags are merged into the request sent to Azure through a generated `ARMTaggable` interface on ARM spec types (Option X). The Kubernetes `spec` is never modified.
* In phase 2, `detach-on-delete` removes the default tags ASO applied (Option D2).

[Phase 1 implementation details](#phase-1-implementation-details) describes how ASO keeps tags and its record of them accurate across long-running `PUT`s and restarts. That mechanism is to be validated by a prototype, and it can change without revisiting this decision.

{{% alert title="Warning" color="warning" %}}
Default tags are information, not a security boundary. Anyone who can edit an ASO resource can set, change or remove its default tags: by setting the same name in `spec.tags`, by listing it in `skip-default-tags`, or by changing the labels and annotations an expression reads. Don't use default tags, especially those derived from labels or annotations, in Azure ABAC conditions, in Azure Policy rules that allow or deny based on tag values, or in any other access-control decision. Organizations that need trusted tag values should set or enforce them with Azure Policy (`deny` or `modify`), which applies to every client. An operator-enforced mode isn't part of this decision (see [Future work](#future-work)).
{{% /alert %}}

### Configuration

A new operator setting, `AZURE_DEFAULT_TAGS`, contains a JSON object that maps each tag name to a CEL expression. The Helm chart exposes it as `azureDefaultTags`, a YAML map rendered with `toJson`. The feature is off when the setting is unset or empty.

```yaml
# Helm values
azureDefaultTags:
  aso-managed-by: "'azure-service-operator'"
  aso-cluster: "'prod-eastus'"
  aso-namespace: "self.metadata.namespace"
  aso-name: "self.metadata.name"
  aso-kind: "self.kind"
  aso-team: "'team' in self.metadata.labels ? self.metadata.labels['team'] : ''"
  cm-resource-parent: "owner.armType != 'Microsoft.Resources/resourceGroups' ? owner.armId : ''"
```

ASO doesn't reserve a prefix. The documentation recommends the `aso-` convention shown above.

ASO doesn't try to detect cluster identity, because there is no reliable cluster name in Kubernetes. Users set it as a literal.

### Expression environment

Every resource kind uses the same small CEL environment, so expressions are compiled and type-checked once at startup instead of once per resource type:

| Variable                    | Type                  | Notes                                                                                                                                       |
| --------------------------- | --------------------- | ------------------------------------------------------------------------------------------------------------------------------------------- |
| `self.apiVersion`           | `string`              | Group and version, for example `resources.azure.com/v1api20200601`.                                                                         |
| `self.kind`                 | `string`              |                                                                                                                                             |
| `self.metadata.name`        | `string`              |                                                                                                                                             |
| `self.metadata.namespace`   | `string`              |                                                                                                                                             |
| `self.metadata.uid`         | `string`              | Changes if the object is recreated or restored from a backup, for example by Velero.                                                        |
| `self.metadata.labels`      | `map(string, string)` | Excludes `serviceoperator.azure.com/*`, which ASO maintains itself. This includes `last-reconciled-version`, the ASO version (see the FAQ). |
| `self.metadata.annotations` | `map(string, string)` | Excludes `serviceoperator.azure.com/*` and `kubectl.kubernetes.io/last-applied-configuration`, which change constantly or are too large.    |
| `owner.armId`               | `string`              | Resource ID of the immediate owner, taken from the owner hierarchy ASO already works out. Empty when the owner is a subscription or tenant. |
| `owner.armType`             | `string`              | Azure resource type of the owner, for example `Microsoft.Sql/servers`. Empty when `owner.armId` is empty.                                   |
| `operator.namespace`        | `string`              | The namespace ASO is installed in. Useful when several ASO installations share a subscription.                                              |

`spec`, `status` and secrets are deliberately left out. Tag values are widely visible, and keeping `spec` out keeps the environment the same for every resource kind. The ASO version is also left out on purpose, which is why `serviceoperator.azure.com/*` labels are excluded as well as annotations (see the FAQ).

Every expression must evaluate to a `string`. If an expression returns the empty string, ASO omits that tag. Users can use this to apply a tag conditionally, or to roll the feature out gradually, for example `self.metadata.namespace in ['team-a'] ? 'team-a' : ''`.

### Validation

Validation at startup is designed so that any configuration it accepts produces tag names and tag counts that are valid on every resource type:

* Each tag name must match `^[A-Za-z][A-Za-z0-9_-]{0,127}$`. That means it starts with a letter, uses only ASCII letters, digits, `_` and `-`, and is at most 128 characters (the storage account limit). This fits the DNS, Traffic Manager and Front Door rules. The pattern is ASO's own rule, derived from the [Azure limitations](https://learn.microsoft.com/azure/azure-resource-manager/management/tag-resources#limitations).
* No two names may differ only in case.
* At most **15** tags can be configured, the lowest per-resource limit in Azure ([limitations](https://learn.microsoft.com/azure/azure-resource-manager/management/tag-resources#limitations)).

Validation at reconcile time checks each value that ASO will send, after the opt-out and override steps in [Merge semantics](#merge-semantics):

* At most 256 characters ([limitations](https://learn.microsoft.com/azure/azure-resource-manager/management/tag-resources#limitations)).
* Printable ASCII only, to avoid the [IMDS decoding problem](https://learn.microsoft.com/azure/azure-resource-manager/management/tag-resources#tag-usage-and-recommendations).

The implementation will keep the 15-tag cap, the 128-character name limit, the 256-character value limit and the name pattern built from them as named constants, each with a comment linking to its source, so they can be reviewed when Azure's limits change.

### Merge semantics

When ASO reconciles a resource whose ARM spec implements `ARMTaggable`:

1. If the resource has `serviceoperator.azure.com/skip-default-tags`, the default tags it lists are dropped. The value is either `*` (drop all of them) or a comma-separated list of tag names, compared case-insensitively.
2. ASO drops every remaining default tag whose name matches a key in the user's `spec.tags`, compared **case-insensitively**. The user's key and value win, and the default's expression isn't evaluated. An expression that would fail on this resource therefore can't block a tag the user set explicitly.
3. ASO evaluates the remaining expressions against the resource, validates the values (see [Validation](#validation)), and adds the non-empty results to the tags in the request.
4. The merge runs only when ASO builds the request for a `PUT`. It runs after `ConvertToARM()` and before `ARMResourceModifier` extensions run, so extensions still have the final say.

   `Claim()` also runs this conversion code, through `AddInitialResourceState()`, but only to work out the resource ID before the finalizer is added. That happens before ASO checks `reconcile-policy`, so it also runs for `reconcile-policy: skip` resources. That conversion builds the request without default tags (for example, through an option on `ConvertResourceToARMResource()` that `BeginCreateOrUpdateResource()` sets and `AddInitialResourceState()` doesn't). It doesn't evaluate, validate or record default tags, so a default-tag expression can't block claiming or read-only adoption. A resource whose `reconcile-policy` doesn't allow changes never reaches the `PUT`, so its default tags are never evaluated.

   Future change detection must ask for the default tags explicitly when it builds the desired state, so they're part of the comparison.

The merge only changes the request sent to Azure. The Kubernetes `spec.tags` is never modified, so GitOps tools don't fight with ASO. `status.tags` comes from Azure, so it shows the merged result, which is how users see what was applied.

### Phase 1 implementation details

{{% alert title="Note" %}}
This section describes a proposed mechanism, not part of the [core decision](#core-decision). It covers the operator-written annotations (`applied-default-tags`, `pending-default-tags` and `pending-default-tags-hash`), the hash and its versioned format, the configuration fingerprint, and the metadata-change predicate. Other sections that refer to these follow the same rule. A prototype in phase 1 will validate them, and they may change based on what it finds.
{{% /alert %}}

Whatever form the mechanism takes, it must:

* Record only the default tags ASO actually sent, so that phase 2 detach cleanup also works for resources tagged in phase 1.
* Update Azure tags when expression-visible metadata or `skip-default-tags` changes, without waiting for `AZURE_SYNC_PERIOD`, including during a long-running `PUT` and with `AZURE_SYNC_PERIOD=never`.
* Apply a configuration change, or turning the feature off, across a restart to resources whose long-running `PUT` was in progress.
* Never trigger a reconcile from ASO's own metadata writes.
* Leave installations that never enable the feature, and kinds that don't implement `ARMTaggable`, behaving exactly as today.

#### Record of applied default tags

When a `PUT` that includes default tags completes successfully, ASO writes the tags it sent to the annotation `serviceoperator.azure.com/applied-default-tags`, as a JSON object that maps tag names to values, for example `{"aso-cluster":"prod-eastus","aso-namespace":"team-a"}`. The record holds only the defaults that were actually sent. Defaults dropped by `skip-default-tags`, overridden in `spec.tags`, or that evaluated to the empty string are left out. When a `PUT` succeeds without a long-running poller, the record is set to the default tags it sent, and a `PUT` that sent none removes the record. For a long-running `PUT`, the pending pair says what was sent. When the poller reports success and a pending pair is present, ASO promotes `pending-default-tags` to `applied-default-tags`, and an empty `{}` removes the record. When no pending pair is present, the `PUT` sent no default tags, so ASO removes the record. This is the case for the feature-off follow-up `PUT`. Promotion happens before any requeue, and the record is removed only after a successful `PUT`, so detach cleanup stays accurate while a failing follow-up `PUT` retries. Promotion and removal run whether or not the feature is currently on, and act on the annotations already on the resource, so turning the feature off clears records written while it was on and does nothing to resources that never had one. Phase 1 already writes the record, so resources tagged before phase 2 ships can still be cleaned up. The record is written only after a `PUT`.

When `reconcile-policy` doesn't allow changes, ASO still refreshes status from Azure through the same code that handles a successful `PUT`, but no `PUT` is sent and the record is left as it is. That code runs in `WatchResource` mode for the refresh and `ManageResource` mode after a `PUT`, so the record logic must run only in `ManageResource` mode.

The record describes the default tags sent in the last successful `PUT`. A default is later dropped from the request if the feature is turned off (see [Update and lifecycle semantics](#update-and-lifecycle-semantics)), if the tag is added to `skip-default-tags`, or if its expression now evaluates to the empty string. Because `PUT` replaces the resource's tags, most resource providers remove a dropped default at that point. Detach cleanup therefore only covers tags sent in the last successful `PUT`. For the few resource providers that keep tags omitted from a `PUT`, a dropped default stays on the resource and has to be removed by hand.

The record is built from the exact default tags in the request, never recalculated when the `PUT` completes, so a label that changes while a long-running `PUT` is in progress can't put values into the record that were never sent. For a long-running `PUT`, ASO stores the sent default tags in `serviceoperator.azure.com/pending-default-tags`, and a hash of the metadata and default-tags configuration they were computed from in `serviceoperator.azure.com/pending-default-tags-hash`, alongside the poller resume token. ASO uses the hash to find metadata changes made while the `PUT` was in progress, and configuration changes or disablement across a restart (see [Reconcile triggers](#reconcile-triggers)). ASO persists the record in the same commit of the resource's metadata that already persists its other state annotations, such as `latest-reconciled-generation` and the poller resume token. A `PUT` that fails, or that Azure rejects, leaves the existing record untouched. If the commit fails, the reconcile fails and is retried: ASO re-sends the `PUT`, or resumes the poller if its token was already persisted, and writes the record when it completes. Until it does, detach cleanup can act on the previous record or on no record. This is safe, because the `Delete` only removes tags whose name and value both match, but it can leave a tag behind, in line with detach cleanup being best effort.

The record is an annotation, not a status field, for three reasons:

* ASO already keeps its own reconcile state in `serviceoperator.azure.com/` annotations, such as `poller-resume-token` and `latest-reconciled-generation`.
* A status field would need a change to every generated status type. ASO also rebuilds status from Azure's response, so the field would show Azure's value rather than what ASO sent.
* Annotations under `serviceoperator.azure.com/` are excluded from the expression environment, so the record can't feed back into tag values.

The record holds at most 15 entries, so it stays well within the Kubernetes annotation size limit.

#### Reconcile triggers

Today a resource is reconciled when its `metadata.generation` changes (a spec change), when its `serviceoperator.azure.com/reconcile-policy` or `serviceoperator.azure.com/credential-from` annotation changes (see `makeStandardPredicate()` in `v2/internal/controllers/controller_resources.go`), and once every `AZURE_SYNC_PERIOD`. Changing any other label or annotation doesn't trigger a reconcile, and with `AZURE_SYNC_PERIOD=never` nothing else does.

When `AZURE_DEFAULT_TAGS` is set, phase 1 adds a predicate to the registrations of resource kinds whose ARM spec implements `ARMTaggable`. It also triggers a reconcile when an update changes either of these:

* `metadata.labels` or `metadata.annotations`, compared after removing the keys that the [expression environment](#expression-environment) excludes from each map (`serviceoperator.azure.com/*` for labels; `serviceoperator.azure.com/*` and `kubectl.kubernetes.io/last-applied-configuration` for annotations).
* The `serviceoperator.azure.com/skip-default-tags` annotation.

All the metadata ASO maintains on a resource is under `serviceoperator.azure.com/`. That includes the `last-reconciled-version` and owner labels, and the `resource-id`, `poller-resume-token`, `latest-reconciled-generation`, `applied-default-tags`, `pending-default-tags` and `pending-default-tags-hash` annotations. The predicate ignores all of it except `skip-default-tags`, which only users set, so ASO's own commits can't trigger another reconcile. The other inputs can't change through metadata. `metadata.name`, `metadata.namespace` and `metadata.uid` are fixed for an existing object. `owner` comes from `spec.owner`, so it changes with the generation. `operator.namespace` and `AZURE_DEFAULT_TAGS` only change with a restart, which reconciles every resource. A restart during a long-running `PUT` is covered by the configuration fingerprint described below. `makeStandardPredicate()` is shared by every generated resource and by the hand-written registrations, such as the MySQL, PostgreSQL and Azure SQL `User` controllers and the Entra controllers, so the new predicate isn't added there. `augmentWithPredicate()` adds it alongside `makeStandardPredicate()` only for generated registrations whose ARM spec implements `ARMTaggable`. Kinds excluded in the generator configuration, such as `VaultKey`, kinds without a top-level `tags` property, and the hand-written controllers keep today's predicate. Changing their metadata therefore triggers no Azure or database operation. When the feature is off, the predicate isn't added and reconcile triggers stay as they are today, and only the check after a long-running `PUT` still runs (see below).

A trigger can arrive while a long-running `PUT` is in progress. While the poller resume token is set, `DetermineCreateOrUpdateAction()` sends every reconcile to `MonitorResourceCreation()`, which only polls the operation and never sends a new `PUT`. ASO already handles a spec change in this window: `BeginCreateOrUpdateResource()` stores the generation it sent in `latest-reconciled-generation`, and when the poller finishes, `resultBasedOnGenerationCount()` requeues the resource if the generation has changed since. Label and annotation changes don't change the generation. Without more work, a relabel during a long-running `PUT` would trigger a reconcile that only polls, and with `AZURE_SYNC_PERIOD=never` the Azure tags would stay stale until something else triggered a reconcile.

When `AZURE_DEFAULT_TAGS` is set, phase 1 handles these inputs the same way. When ASO sends a `PUT`, it computes a hash of everything the sent default tags were derived from, except the spec, which `latest-reconciled-generation` already covers. The hash has two parts. The first is the inputs this predicate watches: the filtered `metadata.labels` and `metadata.annotations` described above, and `skip-default-tags`. The second is a fingerprint of the effective default-tags configuration, meaning the parsed and validated `AZURE_DEFAULT_TAGS` name-to-expression map plus `operator.namespace`. Both change only with a restart. ASO stores the hash in `serviceoperator.azure.com/pending-default-tags-hash`, next to the sent tags in `serviceoperator.azure.com/pending-default-tags` and the poller resume token (see [Record of applied default tags](#record-of-applied-default-tags)). The hash is stored under `serviceoperator.azure.com/`, so writing it can't trigger the predicate. While the feature is on, every long-running `PUT` on an `ARMTaggable` kind stores the pending pair, even when it sends no default tags, for example with `skip-default-tags: *` or when every expression evaluates to the empty string. In that case `pending-default-tags` is `{}`.

The hash is built from parsed values, not raw strings, so key order and whitespace between JSON tokens in the configuration don't matter. ASO trims each expression when it parses the setting, and both compiles and hashes the trimmed form. Tag names keep the casing ASO sends. Whitespace inside an expression is kept, because it can be part of a CEL string literal. Each map is serialized as JSON with keys sorted by byte value and no insignificant whitespace. The hash is the SHA-256 of that serialization, stored with a version prefix such as `v1:<hex>`, so that a future format change counts as a difference instead of a false match. A change with no effect, such as reformatting inside an expression, can still change the hash. That costs at most one extra `PUT`. Missing a real change would leave tags stale. Computing the fingerprint can't fail, because the configuration was validated at startup.

When the poller finishes, ASO hashes the current inputs and the current configuration. If the result differs from the stored hash, ASO requeues the resource, and the next reconcile sends a new `PUT` with the current values. The requeue clears the pending pair, so the follow-up `PUT` starts with none: with the feature off it writes none, and with the feature on it writes a fresh one. The check therefore can't repeat. While the feature is on, a long-running `PUT` on an `ARMTaggable` kind with no stored hash, for example one started before the feature was enabled, counts as a difference. ASO compares inputs instead of re-evaluating the expressions, so polling never evaluates expressions and can't fail with `DefaultTagsInvalid`.

When the feature is off, ASO still runs this check for a resource with outstanding pending state (`pending-default-tags` or `pending-default-tags-hash`) and treats that state as a difference: pending state is only written while the feature is on, so it means the in-flight `PUT` may have sent default tags the current configuration no longer wants. ASO requeues the resource, as above, and the next reconcile sends a `PUT` without default tags. When that `PUT` succeeds, the record is removed as described in [Record of applied default tags](#record-of-applied-default-tags). A resource without pending state is handled exactly as today, so installations that never enabled the feature see no change.

The predicate, the pending pair, the completion check and the record apply only to resource kinds whose ARM spec implements `ARMTaggable`. Other kinds never get default tags, never get pending state written, aren't reconciled on metadata changes, skip the completion check, and behave exactly as today.

If the poller reports a failure, ASO clears `serviceoperator.azure.com/pending-default-tags` and `serviceoperator.azure.com/pending-default-tags-hash` together with the poller resume token. The hash check runs only on success. After a failure, `handleCreateOrUpdateFailed()` classifies the error. A failure with severity `Warning` is retried with the usual backoff, and the retry sends a fresh `PUT` with current values, so a metadata or configuration change made during the failed `PUT` is still applied. A failure with severity `Error`, such as an Azure `400 Bad Request`, is terminal. `failureResult()` in `v2/internal/util/interval/interval_calculator.go` blocks the requeue whatever result the caller returns, so ASO doesn't retry it. A metadata or configuration change made during that `PUT` was consumed by a reconcile that only polled, and it isn't applied until the next trigger: a later spec change, a change to the metadata this predicate watches, or an operator restart. `AZURE_SYNC_PERIOD` doesn't requeue a resource after a terminal failure. This is an exception to the lifecycle guarantee in [Update and lifecycle semantics](#update-and-lifecycle-semantics), and it is the same gap that already exists for a spec change made during a long-running `PUT` that fails terminally, since `resultBasedOnGenerationCount()` is also suppressed then. Phase 1 doesn't close it, because a terminal failure already leaves the resource `Ready=False` with severity `Error` until the user changes something, and that change triggers a fresh `PUT` with the current values. Closing the gap for spec and metadata changes together, for example by downgrading a terminal failure to `Warning` when the generation or the hash has changed since the `PUT`, is listed in [Future work](#future-work). On success, the pending pair is cleared with the token once the remaining success steps complete. If one of those steps fails, the token and the pending state are kept and the next reconcile retries that step. The exception is a NotFound when refreshing status in manage mode (`ManageResource`), which already clears the token to avoid getting stuck. The pending pair is cleared with it. The in-flight `PUT` may have applied the pending tags in Azure, so the record can lag the tags actually applied. The record then isn't updated until the next successful `PUT`, so detach cleanup is best effort here as well. Clearing the pending pair doesn't depend on the feature being on.

### Update and lifecycle semantics

* **Each reconcile.** Because ASO sends a full `PUT` each time, default tags are recalculated and reapplied. Tags that are deleted or edited in the portal are restored, the same as `spec.tags`.
* **A source value changes**, for example because the resource is relabeled, recreated in another namespace, or adopted by another cluster. A change to a label, an annotation or `skip-default-tags` triggers a reconcile (see [Reconcile triggers](#reconcile-triggers)). A recreated or newly adopted object is reconciled because it's new to the controller. If the change arrives while a long-running `PUT` is in progress, ASO sends a new `PUT` once it finishes, or retries if it fails with severity `Warning`. Either way, the tag is updated without waiting for `AZURE_SYNC_PERIOD`, even when it's set to `never`. The exception is a long-running `PUT` that fails with severity `Error`: ASO doesn't retry it, so the change isn't applied until the next spec or watched-metadata change or operator restart. Spec changes already behave this way today (mechanism: see [Phase 1 implementation details](#phase-1-implementation-details)).
* **The feature is turned off.** ASO stops sending the tags. The restart reconciles every resource, and most resource providers remove the tags on that `PUT`. If a long-running `PUT` that sent default tags was in progress at the restart, ASO resumes the poller and, when it succeeds, sends a follow-up `PUT` without them (see [Reconcile triggers](#reconcile-triggers)), even with `AZURE_SYNC_PERIOD=never`. The record is removed when that `PUT` succeeds (see [Record of applied default tags](#record-of-applied-default-tags); mechanism: see [Phase 1 implementation details](#phase-1-implementation-details)).
* **`reconcile-policy: skip`.** The resource isn't touched, so its tags may become stale. That is expected for `skip`. ASO doesn't evaluate default tags for the resource, including when it first claims it, and doesn't change its `applied-default-tags` record. If the policy changes to `skip` while a long-running `PUT` is in progress, the `skip` status refresh clears the poller resume token, as it does today, and ASO clears the pending pair with it without changing the record. The in-flight `PUT` may still complete in Azure, so the record can lag the tags actually applied, and `detach-on-delete` may then leave those tags behind, in line with detach cleanup being best effort (see [Record of applied default tags](#record-of-applied-default-tags)). When changes are allowed again, ASO sends a fresh `PUT` with the current configuration, so no fingerprint check is needed (mechanism: see [Phase 1 implementation details](#phase-1-implementation-details)).
* **`detach-on-delete` (phase 2).** ASO reads the [record of applied default tags](#record-of-applied-default-tags) and drops any name that the user has since added to `spec.tags` (compared case-insensitively) or to `skip-default-tags`. Those changes were made after the last successful `PUT`, so the record doesn't reflect them yet. The record itself never contains overridden or skipped names. It then calls `Tags - Update At Scope` with `Delete` and the remaining recorded name/value pairs, so Azure removes a tag only if it still has the value ASO last wrote. Expressions aren't evaluated at delete time, so a label that changed since the last reconcile can't cause the wrong value to be targeted. Cleanup uses the record whether or not the feature is currently on. This is best effort. If there's no record, there's nothing to remove. If the call fails, ASO records a Warning event and still removes the finalizer, so detaching is never blocked.
* **Changing the setting** requires a restart. ASO already reconciles every resource when it starts, but the first enablement makes a *real* change to every resource. If a long-running `PUT` was in progress at the restart, the configuration fingerprint in its pending hash makes ASO send a follow-up `PUT` with the new configuration once it finishes (mechanism: see [Phase 1 implementation details](#phase-1-implementation-details)). See [Edge cases](#edge-cases) for how to roll it out.

### Error handling

| Failure                                                                                                    | When detected | Result                                                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------- | ------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| The setting isn't valid JSON                                                                               | Startup       | `Validate()` fails and the operator doesn't start. The error is logged, the same as for any other invalid setting.                                                        |
| A tag name fails the pattern, has a case-insensitive duplicate, or more than 15 tags are configured        | Startup       | The operator doesn't start.                                                                                                                                               |
| An expression doesn't compile, refers to an unknown variable, or doesn't return a `string`                 | Startup       | The operator doesn't start.                                                                                                                                               |
| An expression fails at runtime, for example a map key is missing                                           | Reconcile     | The resource gets `Ready=False` with severity `Warning` and reason `DefaultTagsInvalid`, and the message names the tag. Reconciliation is retried with the usual backoff. |
| A value is longer than 256 characters or isn't printable ASCII                                             | Reconcile     | Same as a runtime expression failure. Values aren't truncated or rewritten, because that would make tags quietly unreliable for audit.                                    |
| Azure rejects the merged tags, for example because there are too many or a service-specific rule is broken | Reconcile     | Handled by the existing Azure error path, so the error appears on the resource's `Ready` condition.                                                                       |
| Removing tags on detach fails                                                                              | Delete        | A Warning event is recorded and the detach continues.                                                                                                                     |

Startup errors stop the operator because a bad configuration would otherwise fail every resource. Runtime errors are reported on the individual resource because they depend on its data. Failing the reconcile, instead of silently leaving the tag off, means that if a resource reconciled successfully, users know it carries every configured tag. Only defaults that survive `skip-default-tags` and the `spec.tags` override are evaluated (see [Merge semantics](#merge-semantics)). A default that the user has overridden or opted out of therefore can't cause `DefaultTagsInvalid`.

### Edge cases

#### Tag counts and names

* **User tags plus default tags exceed the resource's limit.** This is most likely on the 15-tag services. Azure rejects the request and the error appears on the resource. The fix is to opt out per key with `skip-default-tags`.
* **The same tag name with different casing.** When a user's `spec.tags` key matches a default tag name except for casing, the user's key and value win. Azure may keep whichever casing it received first, so cost reports can show more than one casing across a fleet. This is documented.
* **Resource-specific update rules**, for example VM extensions that can only be retagged while the VM is running. These surface as ordinary Azure errors. The documentation notes that turning on default tags can expose them on resources whose tags never changed before.

#### Values

* **Values that change often.** ASO-maintained annotations (the poller resume token, the latest reconciled generation, the pending default tags and their hash) and labels (the last reconciled ASO version) would change a tag on every reconcile or every upgrade. Everything under `serviceoperator.azure.com/` is removed from both maps. The documentation warns against basing tags on other annotations that change often, such as those set by Argo CD. With the feature on, changing any label or annotation that expressions can see triggers a reconcile, and so a `PUT`, even if no expression reads it (see [Reconcile triggers](#reconcile-triggers); mechanism: see [Phase 1 implementation details](#phase-1-implementation-details)). Tools that update such metadata often therefore cause a `PUT` each time.
* **`owner.armId` longer than 256 characters** in deep hierarchies. This fails validation and is reported as `DefaultTagsInvalid`. Users can guard it with `size(owner.armId) <= 256 ? owner.armId : ''`.

#### Rollout and ownership

* **The first enablement changes every resource at once.** This can cause Azure throttling and long-running operations, for example on AKS, API Management and Cosmos DB. An Azure Policy that restricts tag values could also reject every update at once. The documentation recommends rolling out gradually with a namespace condition in the expressions, and checking existing `deny` policies before enabling.
* **Two installations manage the same resource**, for example during a cluster migration. Each `PUT` overwrites the other's identity tags. This is visible in `status.tags` and the Activity Log. Warning about it automatically is future work.
* **Multi-tenant clusters.** Labels and annotations are controlled by whoever can edit the object, and in multi-tenant setups (see [Secure Workload Identity Multitenancy]({{< relref "ADR-2026-07-Secure-Workload-Identity-Multitenancy" >}})) that is the namespace tenant. Tags built from labels can therefore be set by tenants. The documentation recommends basing audit tags on `self.metadata.namespace`, `operator.namespace` and literals. These inputs can't be changed through labels or annotations, but a resource editor can still override or skip the tag with `spec.tags` or `skip-default-tags`, so it isn't an access-control signal (see the warning under [Core decision](#core-decision) and [Future work](#future-work)). Tags built from labels or annotations must not be used for access control.

#### Lifecycle

* **`asoctl import` copies default tags into `spec.tags`.** Because user tags win, the imported values would never change again, for example an old namespace. This is documented, and an `asoctl` option to exclude tag names is future work.
* **Resources adopted with tags that aren't in `spec`.** These tags are removed by the `PUT`, which is existing behavior. This feature doesn't change it, but users turning the feature on may expect the tags to be merged. The documentation states this explicitly.

#### Coverage

* **Resources that don't use the normal `PUT` path.** The SQL, MySQL and PostgreSQL user reconcilers and the Entra reconciler don't get default tags, because database users and Entra objects have no Azure tags.
* **KeyVault keys are excluded in phase 1**, even though their ARM spec has a top-level `tags` property. Azure applies the `PUT` only when it creates a key, and the extension then replaces the key's tags with `spec.tags` whenever `spec.tags` is set (see [Relevant ASO behavior](#relevant-aso-behavior)). Injected defaults would therefore be removed right after creation, or never updated, while the record said they were applied. The shared injection point wouldn't skip keys on its own, so the generator configuration marks `VaultKey` (in `v2/azure-arm/keyvault.yaml`, alongside its existing `$importable: false`) so that its ARM spec doesn't implement `ARMTaggable`. The merge then never runs for keys, and ASO never writes an `applied-default-tags` record for them. Keys can opt in once the extension merges default tags into its `UpdateKey` call and records them (see [Future work](#future-work)).
* **Resources that keep tags somewhere other than a top-level `tags` property** are skipped. The generator change includes a report of resources that have a property named `tags` in any other shape, so the gaps are known and documented.

### FAQ

This FAQ also covers objections we expect reviewers to raise.

**Q: Doesn't this contradict the [Templating]({{< relref "ADR-2026-05-Templating" >}}) ADR? Why not use Kyverno or MutatingAdmissionPolicy?**

A: The Templating ADR rejects expressions *in resource properties* because they weaken schema validation. This feature leaves every property alone and adds operator-level metadata to the request sent to Azure, which is closer to how ASO already sets its own annotations. Admission mutation was considered (Option 0) and rejected because:

* It needs a per-kind allowlist that changes with every release.
* It conflicts with GitOps tools.
* It can't follow Azure's per-service limits or use ASO's owner hierarchy.

Users who prefer admission mutation can still use it. Tags it puts in `spec.tags` simply win over the defaults.

**Q: Isn't CEL too much? Could a simpler syntax work?**

A: CEL is already a dependency, is already documented for `operatorSpec` expressions, and is compiled at startup, so most mistakes are found before any resource is touched. If reviewers prefer a smaller first step, the same configuration format can start with literals and a fixed set of variables, and grow into full CEL later without breaking existing configurations.

**Q: If user tags win, can't anyone spoof ownership? What about tag-based access control or policy?**

A: Yes. Default tags are information, not enforcement (see the warning under [Core decision](#core-decision)). Organizations that rely on tags for Azure ABAC conditions or Azure Policy should enforce them with Azure Policy (`deny` or `modify`), which applies to every client, not only ASO. A per-key "enforced" mode is listed under [Future work](#future-work). It isn't part of this decision.

**Q: The generator change regenerates every resource. Is it worth it?**

A: The interface is small and only affects the types used to build Azure requests. Storage types and conversions don't change. The alternative of patching the serialized request (Option Y) can't tell which resources support tags.

**Q: Won't this create a support burden from Azure tag errors?**

A: Startup validation (the safe name pattern and the 15-tag cap) removes most configuration mistakes that would only show up on some resource types. Any remaining errors come from Azure and appear on the affected resource with Azure's own message.

**Q: Why operator-wide only? Platform teams keep cost centers on namespace labels.**

A: Supporting namespace labels needs a namespace watch, and a label change would re-reconcile every resource in that namespace. That's worth doing, but separately (see future work). Literals and per-namespace conditions in expressions cover many cases today.

**Q: Why is this in the `aso-controller-settings` Secret, and why does it need a restart?**

A: For consistency. Every other operator setting works this way. A restart is a planned point at which every resource is reconciled anyway. A watched ConfigMap (Option B) would trigger the same load at an unpredictable time.

**Q: Why not reserve a prefix such as `serviceoperator.azure.com/`?**

A: `/` isn't allowed in Azure tag names, and users asked for configurable tags. A recommended convention is enough.

**Q: Why isn't the ASO version available?**

A: Using it would change the tags on every resource at every upgrade, which means a real update to every resource in the fleet, including slow operations on AKS clusters and API Management. `AZURE_USER_AGENT_SUFFIX` and the Activity Log already show which installation made a change. For the same reason, `self.metadata.labels` excludes `serviceoperator.azure.com/*` labels. Otherwise the `serviceoperator.azure.com/last-reconciled-version` label, which ASO sets to its own version on every reconcile, would bring the version back in indirectly.

**Q: Does this change how ASO interacts with Azure Policy?**

A: No. Tags added by a `modify` policy are applied to ASO's `PUT` as they are today. Default tags can help satisfy `deny` policies that require certain tags. A `modify` policy that overwrites a default tag wins in Azure, and `status.tags` shows the policy's value. Future change detection would otherwise see that difference after every `PUT` and treat it as permanent drift. Comparing default tags is therefore expected to be configurable (see [Requirements](#requirements)), and users can list the tag in `skip-default-tags` on affected resources, which removes it from the desired state. On detach, the policy's value doesn't match the recorded value, so ASO leaves that tag alone.

**Q: Why fail the reconcile when a tag can't be evaluated, instead of skipping that tag?**

A: The audit use case depends on knowing that a resource that reconciled successfully has every configured tag. Expressions can provide their own fallback (`... ? ... : 'unknown'`) when a missing value is acceptable. Setting the tag in `spec.tags`, or listing it in `skip-default-tags`, also avoids the failure, because that default's expression is then never evaluated for the resource.

## Status

Proposed. Accepting this ADR accepts the [core decision](#core-decision). The [Phase 1 implementation details](#phase-1-implementation-details) are confirmed by a prototype during phase 1.

## Consequences

### Phase 1

The predicate, the follow-up reconcile after a long-running `PUT`, the operator-written annotations and their tests below implement [Phase 1 implementation details](#phase-1-implementation-details), and may change with the prototype.

* A generator change adds the `ARMTaggable` interface to ARM spec types that have a top-level `Tags map[string]string`, except resources excluded in the generator configuration (in phase 1, only KeyVault `VaultKey`; see [Coverage](#coverage)), followed by a regeneration. The change also produces a report of resources whose `tags` property has a different shape.
* A new operator setting, `AZURE_DEFAULT_TAGS`, plus the `azureDefaultTags` Helm value and wiring through `aso-controller-settings`.
* The CEL package gains a fixed "tag environment" (with `self`, `owner` and `operator`) alongside the existing per-type environment.
* A predicate that reconciles a resource when its expression-visible labels or annotations, or its `skip-default-tags` annotation, change (see [Reconcile triggers](#reconcile-triggers)). `augmentWithPredicate()` adds it alongside `makeStandardPredicate()` only when `AZURE_DEFAULT_TAGS` is set and only for generated registrations whose ARM spec implements `ARMTaggable`. `makeStandardPredicate()` itself is unchanged, so excluded kinds and the hand-written database `User` and Entra controllers keep today's triggers.
* A follow-up reconcile after a long-running `PUT` when expression-visible metadata or `skip-default-tags` changed while it was in progress, or when `AZURE_DEFAULT_TAGS` or `operator.namespace` changed or the feature was turned off across a restart. `BeginCreateOrUpdateResource()` stores the sent default tags in `serviceoperator.azure.com/pending-default-tags` and a hash of the watched inputs and a canonical fingerprint of the default-tags configuration in `serviceoperator.azure.com/pending-default-tags-hash`, next to the poller resume token. When the poller reports success, a check on the success path of `MonitorResourceCreation()`, next to the generation check in `resultBasedOnGenerationCount()`, also requeues the resource if the current hash differs, if no hash was stored while the feature is on, or if the feature is off and pending state is outstanding. The check does not run when the poller reports failure. A failure with severity `Warning` is retried with a fresh `PUT`, and a failure with severity `Error` isn't retried, as for spec changes today (see [Reconcile triggers](#reconcile-triggers), which also covers when the pending pair is cleared).
* Four new annotations: `serviceoperator.azure.com/skip-default-tags` (set by users), `serviceoperator.azure.com/applied-default-tags` (written by ASO, see [Record of applied default tags](#record-of-applied-default-tags)), and the internal `serviceoperator.azure.com/pending-default-tags` and `serviceoperator.azure.com/pending-default-tags-hash` (written by ASO during a long-running `PUT`, see [Reconcile triggers](#reconcile-triggers)), plus a new condition reason, `DefaultTagsInvalid`.
* Documentation updates to [ASO controller settings options]({{< relref "aso-controller-settings-options" >}}) and the [annotations guide]({{< relref "annotations" >}}). They cover:
  * Worked examples, including `cm-resource-parent`.
  * Gradual rollout and checking Azure Policy before enabling.
  * The multi-tenant guidance.
  * The `asoctl import` warning.
  * The list of resources that don't get default tags.
  * `applied-default-tags`, `pending-default-tags` and `pending-default-tags-hash` listed under "Annotations written by the operator".
* Tests:
  * Table-driven unit tests for configuration parsing and validation, case-insensitive merging (including that overridden and skipped defaults aren't evaluated), `skip-default-tags`, the contents of the `applied-default-tags` record, the reconcile predicate (it fires on expression-visible label and annotation changes and on `skip-default-tags`, and ignores `serviceoperator.azure.com/*` changes such as `last-reconciled-version` and `applied-default-tags`), predicate registration (with the feature on, only generated `ARMTaggable` registrations get the new predicate, and `VaultKey`, a kind without top-level `tags` and the hand-written database `User` and Entra registrations don't fire on a label change), the claim conversion (it doesn't evaluate default tags, so a missing-label expression doesn't fail claiming a `reconcile-policy: skip` resource), the `skip`-policy status refresh (it doesn't change the record), and the configuration fingerprint (it doesn't change when keys are reordered, when JSON whitespace changes, or when leading/trailing whitespace is added to an expression, and it changes when a tag name, an expression or `operator.namespace` changes).
  * Reconciler unit tests for metadata that changes during a long-running `PUT`. In the main case, ASO starts a `PUT` and stores the poller resume token, and then the `team` label changes. The next reconcile only polls, and the poller reports success. ASO then requeues the resource, and the following reconcile sends a `PUT` with the new tag value. The same sequence also requeues when only `skip-default-tags` changes, or when no hash was stored while the feature is on. It doesn't requeue when nothing changed, or when only `serviceoperator.azure.com/*` metadata such as `last-reconciled-version` changed. They also check that the pending pair is cleared when the poller reports failure and after a successful completion, and kept when a later success step fails. Two failure scenarios change the `team` label while the `PUT` is in progress. When the poller then reports a failure with severity `Warning`, ASO retries, and the retry `PUT` carries the new tag value. When it reports a failure with severity `Error`, ASO doesn't requeue, matching today's behavior for a spec change, and a later change to a watched label triggers a `PUT` with the current values. Two restart scenarios run the same sequence, with ASO restarting after the poller resume token is stored. In the first, the restart is with a changed `AZURE_DEFAULT_TAGS`: when the poller reports success, ASO requeues, and the next `PUT` carries the new tags. In the second, the restart is with the feature turned off: when the poller reports success, ASO promotes the pending tags, requeues, sends a `PUT` without default tags, and then removes `applied-default-tags` and clears the pending pair. With the feature off and no pending state, a completed long-running `PUT` doesn't requeue. With the feature on, a long-running `PUT` on a non-`ARMTaggable` kind doesn't requeue. With the feature on, a long-running `PUT` that sends zero default tags stores the pending pair and doesn't requeue once it completes.
  * Generator golden-file tests for the interface, including that `VaultKey` doesn't implement `ARMTaggable`.
  * Envtest recordings with the feature enabled for a resource group, a DNS zone (to confirm the safe name pattern and the 15-tag limit), and a child resource that uses `cm-resource-parent`.

### Phase 2

* Removing tags on `detach-on-delete` through the Tags API's `Delete`, using the name/value pairs recorded in `serviceoperator.azure.com/applied-default-tags`, minus any names now overridden in `spec.tags` or listed in `skip-default-tags`.

### Future work

These items are not decided by this ADR:

* A Warning event when a resource's identity tags in Azure show that another installation is managing it.
* An `asoctl import` option that excludes tag names from imported specs.
* Namespace labels in the expression environment (`namespace.labels`), with a namespace watch.
* A per-key "enforced" mode, where the operator value wins over `spec.tags` and `skip-default-tags`. This needs its own design and security review. Even an enforced tag can only be trusted if its expression reads only inputs resource editors can't set, such as literals, `self.metadata.namespace` and `operator.namespace`. Label-derived tags would stay unsuitable for access control.
* Default tags for KeyVault keys, once the extension merges them into its data-plane `UpdateKey` call and records them.
* Retrying a long-running `PUT` that fails with severity `Error` when its inputs changed while it was in progress, for example by downgrading the failure to `Warning` when the generation or `pending-default-tags-hash` no longer matches. This would close the same gap for spec changes and metadata changes, and changes existing behavior for every resource, so it needs its own review.
* Azure Service Groups, once generally available, as another way to group resources by cluster.

## Experience Report

TBC

## References

* [#5713: Feature request: optionally tag Azure resources with owner metadata](https://github.com/Azure/azure-service-operator/issues/5713)
* [Use tags to organize your Azure resources (limitations)](https://learn.microsoft.com/azure/azure-resource-manager/management/tag-resources#limitations)
* [Tag support for Azure resources](https://learn.microsoft.com/azure/azure-resource-manager/management/tag-support)
* [Tags - Update At Scope REST API](https://learn.microsoft.com/rest/api/resources/tags/update-at-scope)
* [Azure Policy tag policies and inheritance](https://learn.microsoft.com/azure/azure-resource-manager/management/tag-policies)
* [Cost Management tag inheritance](https://learn.microsoft.com/azure/cost-management-billing/costs/enable-tag-inheritance)
* [Cost Management: group related resources (`cm-resource-parent`)](https://learn.microsoft.com/azure/cost-management-billing/costs/group-filter#group-related-resources-in-the-resources-view)
* [Announcing Public Preview for Azure Service Groups](https://microsoft.github.io/azurekorea/Announcing-Public-Preview-for-Azure-Service-Groups/)
* [Kubernetes MutatingAdmissionPolicy](https://kubernetes.io/docs/reference/access-authn-authz/mutating-admission-policy/)
* [KEP-3962: Mutating Admission Policies](https://www.kubernetes.dev/resources/keps/3962/)
* [Kyverno](https://kyverno.io)
* [AWS Controllers for Kubernetes resource tags](https://aws-controllers-k8s.github.io/community/docs/user-docs/ack-tags/)
* [Config Connector resources](https://cloud.google.com/config-connector/docs/concepts/resources)
* [Change Detection]({{< relref "ADR-2022-11-Change-Detection" >}})
* [Templating]({{< relref "ADR-2026-05-Templating" >}})
* [Dynamic Export to Secret or ConfigMap]({{< relref "ADR-2024-07-Dynamic-Export-To-Secret-Or-ConfigMap" >}}) (existing CEL usage)
* [Secure Workload Identity Multitenancy]({{< relref "ADR-2026-07-Secure-Workload-Identity-Multitenancy" >}})
