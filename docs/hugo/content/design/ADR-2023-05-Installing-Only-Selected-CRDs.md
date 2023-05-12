---
title: '2023-05: Installing only selected CRDs'
---

## Context

Outstanding issues:
1. [#2920: kube-apiserver persistently high memory usage with large number of CRDs](https://github.com/Azure/azure-service-operator/issues/2920).
2. [#1433: Allow users to control which CRDs are installed](https://github.com/Azure/azure-service-operator/issues/1433)

CRDs consume kube-apiserver resources even if no instances of the CRDs are ever created. For more details, see 
[#2920](https://github.com/Azure/azure-service-operator/issues/2920).


We have had a longstanding item ([#1433](https://github.com/Azure/azure-service-operator/issues/1433)) to allow users 
to be more selective about what CRDs to install. As the number of CRDs installed by ASO increases, 
this becomes more critical in order to mitigate the kube-apiserver memory usage problem.

### Requirements

Before getting into different options for solving this problem we should go over the requirements.

1. Reduction in default number of CRDs installed.
    * The goal is that the default ASO configuration should fit into a minimal memory environment
      (such as a small `kind` cluster or a `Free` tier AKS cluster).
2. Seamless upgrade experience for existing users.
    * If the user has installed a previous version of ASO and upgrade to a new version, whatever 
      resources they had installed should continue to function. The previous default was "everything".
3. Seamless upgrade experience going forward.
    * If a user selects a subset of CRDs when initially installing ASO, they should get that exact 
      subset automatically during upgrade unless they say otherwise.
    * No "oops I forgot to select any CRDs, so none were installed"
    * No "oops I forgot to filter CRDs, so they were all installed"
    * Exactly the CRDs the user originally specified should be upgraded, and only those CRDs.
4. No CRDs uninstalled ever.
    * CRD uninstallation is dangerous. If users want to remove CRDs they can do it manually. 
    * TODO: A reduction in the list of CRDs to install should result in only the new list CRDs of CRDs getting new versions?
      Or maybe it needs to actually upgrade those versions too even though they weren't specified, to maintain a consistent
      storage version?

### Question 1: What to install by default?

As mentioned above, a key goal is a reduction in the default number of CRDs installed. That means we somehow need to 
choose which CRDs we will install and which CRDs we will not.

#### Answer 1a: Handpicked default

We handpick a set of resources we feel are "critical" and install just those by default

Pros:
* Has at least a chance of working for some user-scenarios out of the box.
* Ensures that resources which are extremely likely to be needed (such as Microsoft.Resources/resourceGroup) are included 
  by default.

Cons:
* Arbitrary.
* May result in higher kube-apiserver memory consumption than is strictly needed for users as it is likely to include some
  resources which they may not use.

#### Answer 1b: Nothing by default

We install nothing by default. The operator may treat "no CRDs" as an error case and exit or it may
treat it as a standard case and run but do nothing.

Pros:
* Users choose exactly what they need.

Cons:
* The operator is nonfunctional out of the box, which is a poor new-user experience.

#### Answer 1c: Mandatory configuration

Selection of CRDs to install is required and the operator pod will not launch without it.

Can also use [Helm required values](https://stackoverflow.com/questions/53866196/how-best-to-say-a-value-is-required-in-a-helm-chart)
to ensure that the value is required at the Helm chart level, although making it required may run afoul of requirement #3.

Pros:
* Users choose exactly what they need.

Cons:
* The operator is nonfunctional out of the box, but at least it's clearly nonfunctional and has a good error why.

### Question 2: How do users specify what CRDs to install?

#### Answer 2a: Using the existing command line parameter

Continue using the existing `--crd-pattern` command line parameter on the operator pod

Pros:
* Mostly exists today

Cons:
* Command-line isn't preserved between upgrades, which may fail requirement #3 (seamless upgrade experience).

#### Answer 2b: Use a ConfigMap

A `ConfigMap` would persist between upgrades and so in theory could be used to remember what was installed previously
without the user needing to re-specify it.

One challenge is that in the Helm context if the `ConfigMap` is created by Helm, Helm will overwrite it in
subsequent upgrades. This could be dealt with by using a
[pre-install hook](https://stackoverflow.com/questions/55503893/helm-patch-default-service-account). We would create and
patch the `ConfigMap` in the pre-install hook.

One possible ConfigMap structure:
```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: managed-resources
  namespace: azureserviceoperator-system
data:
  "all": ""  # magic key meaning "*" which isn't supported in keys
  microsoft.resources: "*"
  microsoft.storage: "StorageAccount,StorageAccountsBlobService"
```

The following `patch` commands could be run either directly by the user or by our pre-install hook:

* Addition of new resources: `kubectl patch configmap managed-resources -n azureserviceoperator-system --patch '{"data":{"microsoft.storage":"*"}}'`
* Removal of old resources: `kubectl patch configmap managed-resources -n azureserviceoperator-system --type=json -p='[{"op": "remove", "path": "/data/microsoft.resources"}]'`

If nothing is specified as part of the upgrade, the set of supported resources will not change.

If the `ConfigMap` doesn't exist when we go to **upgrade**, we can assume we're coming from an older ASO version that didn't
support the `ConfigMap` and default to creating one that signals to install all resources.

Pros:
* Meets all the requirements.
* Works for both Helm and raw-YAML situations. Helm as discussed above and raw YAML by having users.

Cons:
* Complicated to implement.
* Not the behavior that Helm users would necessarily expect for Helm values.
* Upgrade scenario from versions prior to this feature will still have all CRDs and thus still have the memory usage problem.

Open questions:
* Do we even need to support a Helm value for configuring this? Technically speaking the only reason we would need to is
  to enable ASO as an Extension or Addon, as there everything must be done through Helm. For users of Helm itself we likely
  can just tell them to run a `kubectl` command afterwards to create the `ConfigMap`. That does make the Helm install
  experience worse though.
* We may be able to create the `ConfigMap` as empty given Helms use of
  [3 way strategic merge](https://helm.sh/docs/faq/changes_since_helm2/#improved-upgrade-strategy-3-way-strategic-merge-patches)
  and then patch it afterwards.
* Is it OK that _providers_ would be patchable, but you would need to include all the resources when patching something
  like `microsoft.storage: "StorageAccount,StorageAccountsBlobService"`? We could just disallow individual resource selection
  for now, although that might eventually become a problem for certain large providers (`Microsoft.Network` comes to mind).

See also Helm [post rendering](https://helm.sh/docs/topics/advanced/#post-rendering), which we could also consider.
pre-install hook (or post-install) seems easier though.

#### Answer 2c: Using a CRD

We investigated doing this originally for [Helm chart size limitations](../ADR-2023-02-Helm-Chart-Size-Limitations)

Pros:
* More structured than a `ConfigMap`.

Cons:
* Chicken and egg problem of using the CRD that's part of the operator to configure the operator behavior.
  Helm has a very hard time creating a CRD _and_ an instance of the CRD in the same deployment.
* May need to implement collection strategic merge behavior that already exists for `ConfigMap`.

### Question 3: How do we ensure consistent configuration as ASO is upgraded

#### Answer 3a: Require the user to be consistent when upgrading

The user would need re-specify what resources they want every time they upgrade

Pros:
* Easy for us

Cons:
* Very easy to make a mistake and accidentally leave a resource off they want.

#### Answer 3b: Use a configuration technique that persists across upgrades

Pros:
* Satisfies the requirement.

Cons:
* Would augment/complicate the approach we already have
* More complex for us to implement, especially with Helm. Helm owns any resource it creates, 
  and deletes it if not specified on the next upgrade. See [Using a ConfigMap](#answer-2b-use-a-configmap) for some 
  workarounds.

**Variant: Store the configuration independently** 

For example user could specify configuration in a non-durable way (Helm variables that pass to cmdline) and then we could
merge this configuration into a "store" such as a `ConfigMap`. This differs from the above pattern only in that the user
would not interact with this `ConfigMap` directly, it would be managed by ASO.

It has similar pros/cons as above.

#### Answer 3c: Query existing CRDs and always support those

Combined with some form of non-persistent storage such as [command line](#answer-2a-using-the-existing-command-line-parameter)
we would query all existing CRDs and install/upgrade them, while adding whatever new CRDs were specified in 
the non-persistent storage.

Pros:
* Automatically supports seamless upgrades.
* Installed resources are always upgrade to the latest version, even if user has forgotten about them.
* Could be built onto existing cmdline configuration.
* Hard for users to get wrong.
* Single source of truth.

Cons:
* ???

## Decision

### Answer to [question #1: What to install by default?](#question-1-what-to-install-by-default)

A combination of  [Answer 1b: Nothing by default](#answer-1b-nothing-by-default) and 
[Answer 1c: Mandatory configuration](#answer-1c-mandatory-configuration).

Specifically, the operator pod should exit with an error if there are no CRDs configured in `--crd-pattern` _and_ no
existing CRDs in the cluster. If a `--crd-pattern` is specified, or if there are already existing ASO CRDs in the cluster
then the pod will launch without an error, install/upgrade the required CRDs and then start watches on those CRDs.

### Answer to [question #2: How do users specify what CRDs to install?](#question-2-how-do-users-specify-what-crds-to-install)

[Answer 2a: Using the existing command line parameter](#answer-2a-using-the-existing-command-line-parameter)

This parameter will come to mean: "New CRDs to install to the cluster".

### Answer to [question #3: How do we ensure consistent configuration as ASO is upgraded?](#question-3-how-do-we-ensure-consistent-configuration-as-aso-is-upgraded)

[Answer 3c: Query existing CRDs and always support those](#answer-3c-query-existing-crds-and-always-support-those).

Note that users may manage their own CRDs. In Helm, they can do this by setting `installCRDs` in values.yaml to `false`.
When `installCRDs` is set to `false`, the operator only has CRD `read,list,watch` permissions and cannot update CRDs. We
need to ensure that we cater to this case. Doing so doesn't really fundamentally change the algorithm at pod start, which 
will still be:

* Pod launches and reads the `--crd-pattern` variable, containing new CRDs to install.
* Pod reads the expected CRDs from disk.
* Pod reads the currently installed CRDs in the cluster.
* Set of CRDs to install is determined by comparing expected CRDs (`--crd-pattern` + already installed CRDs) with
  the CRDs from disk. Any CRDs with differences (or which don't exist) are candidates for installation/upgrade.
* Operator installs the CRDs (will fail if `--installCRDs` is false). We may need to improve the error here to ensure
  users understand it.

### Open questions and their answers

#### Should we do "graph-based" inclusions of resources?

If a user imports `microsoft.storage/storageAccountsBlobService` should we automatically determine that they 
also want `microsoft.storage/storageAccounts` since that's the `owner` type?

Conclusion: No, users must list all resources they want explicitly. We strongly recommend using full provider imports such as 
`microsoft.networking/*` rather than tightly scoping to specific resources.

## Status

Proposed.

## Consequences

TBC

## Experience Report

TBC

## References

None