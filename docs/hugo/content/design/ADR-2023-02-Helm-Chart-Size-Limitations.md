---
title: '2023-01: Helm Chart Size Limitations and Workarounds'
---

## Context

Helm v3 stores chart state in a Kubernetes secret by default. Kubernetes secrets are limited to a max size based on 
etcd's configuration. The default etcd configuration limits secrets to a maximum size of 1MB.

Our chart secret is fast approaching 1MB. In fact, 
[outstanding PRs](https://github.com/Azure/azure-service-operator/pull/2727) already cause us to surpass the 1MB limit.

Once this limit is reached, users will be unable to install our chart. When this happens, the error looks like:
> Error: create: failed to create: Secret "sh.helm.release.v1.asov2.v1" is invalid: data: Too long: must have at most 1048576 bytes

Digging in a bit, the Helm secret contains compressed data. This means that our chart itself can consume >1MB as long 
as its compressed size is <1MB. Here's a breakdown of our chart size right now:

| Scenario                         | Chart size | Helm secret size | Space savings |
|----------------------------------|------------|------------------|---------------|
| v2.0.0-beta.4                    | 9917 KB    | 732 KB           | 93%           |
| v2.0.0-beta.4 (all docs removed) | 8600 KB    | 536 KB           | 94%           |

Unfortunately, Helm stores the chart _and all subcharts_ state in a single secret, which means that splitting the chart
doesn't prevent us from bypassing this limit if we include the split-charts as subcharts.

Helm supports alternative storage schemes such as a [SQL backend](https://helm.sh/docs/topics/advanced/#sql-storage-backend)
that bypasses this limit, but we can't reasonably expect all users to use that.

**We need to do something to overcome the 1MB max Helm chart limit**

### Splitting the Helm Chart

We could split the Helm chart into a number of separate charts, 1 for the operator like we have now, 
and N charts containing just CRDs. The main question here is how to decide which CRDs go into which chart? There
are a few options:

1. Based on popularity/usage, with multiple groups together to limit the total number of charts
2. By RP/group, with a chart per group. For example `networking`, `compute`, and `resources` would all be different charts.

Regardless of which Helm split option we choose, we likely need to support 
[#1433](https://github.com/Azure/azure-service-operator/issues/1433).

The requirements for the chosen chart breakdown would need to be future-proof. Adding new versions of resources just makes
that CRD bigger, it doesn't add a new CRD, so we need to be careful to leave ourselves lots of room to grow in the future.

I think that the best option for chart split is by RP/group. It ends up with a clean structure that matches SDKs. Other 
possible splits seem to risk some charts becoming too large, or a confusing experience for customers. Imagine "what chart
contains RoleAssignment again?"

**Pros**:
- It's clear which resource goes where, and even for the large RPs like `compute` and `networking` there should be plenty
  of space in each chart to grow and add new resources/versions, at least in the near term.

**Cons**:
- There's no way with Helm out of the box to get an all-in-one installation, which is the current default now.
  There are tools like [Helmfile](https://github.com/helmfile/helmfile) that could enable an all-in-one story.
- There are a lot of RPs, which translates into a lot of Helm charts. The number of charts might get a bit overwhelming.
- It's possible we run into this problem again with one of the huge RPs like `compute` or `networking` a few years down
  the line if we keep adding new versions and new resources.

### Avoiding Helm (for installing CRDs)

Helm has pretty poor CRD support. In fact, we install CRDs via Helm by treating them like normal resources. We
have to do this because otherwise Helm never upgrades them at all. We could just stop using Helm to install CRDs entirely.

There are a few options for this:
1. We have users install CRDs manually via `kubectl`.
2. We update `asoctl` to support installing CRDs for users.
3. We embed the CRDs into the operator pod and have the operator pod manage them on startup.

For all of the above options, we would still maintain a Helm chart for the operator itself (the deployment, service account, etc).

Of the above options, I believe embedding the CRDs into the operator pod is the best option. Here are some pros/cons for it.

**Pros**:
- Helm wasn't really getting us a lot for managing CRDs anyway, and embedding the CRDs into the operator itself
  means that we can guarantee the right versions are installed for the corresponding operator version (unlike all
  the chart-splitting options above)
- We're able to continue to offer an easy-to-use Helm chart that provides the entire ASO feature set.
- Because it's integrated into the operator we could add safeguards and capabilities to this "CRD reconciliation" 
  for both deprecation or uninstall cases over and above what we could easily do in Helm.
- More easily enables [#1433](https://github.com/Azure/azure-service-operator/issues/1433) as compared to a single 
  monolithic Helm chart. We could offer configuration options either via a ConfigMap or a CRD that _is_ included in the
  Helm chart. That single included CRD would be something like `InstalledResourceDefinitions` and support specifying what groups
  or individual resources would be installed.

**Cons**:
- CRD installation would _probably_ be a required pre-step and if it failed the operator would stop running, which means
  existing resources wouldn't be reconciled until the deployment was rolled back. We could possibly mitigate this by making
  CRD installation failure nonfatal, but that has other complications.
- Operator needs permission to CRUD CRDs, which it previously didn't have.
- Uninstalling the Helm chart will uninstall the operator pod but will not uninstall the CRDs. This is somewhat mitigated
  by the fact that uninstalling CRDs is very dangerous and/or impossible if there are existing resources of that CRD type
  even in the current ASO chart.

Another possible option is using `asoctl`. Since that involves writing pretty much the same code as would be required to 
manage the CRDs from the operator, and would preclude installing ASO as an addon in AKS easily (since they want everything
Helm deployable), it seems to make sense to manage the CRDs from the operator rather than from `asoctl`.

### Digging in
We will introduce a new CRD `InstalledResourceDefinitions.serviceoperator.azure.com`, with the following structure:
```go
// +kubebuilder:object:root=true
type InstalledResourceDefinitionsList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items           []InstalledResourceDefinitions `json:"items"`
}

type InstalledResourceDefinitionsSpec struct {
    // +kubebuilder:validation:Required
    // Pattern defines the set of CRDs which should be installed.
    // Currently only "*" (all resources) is supported.
    // Future expansion will allow "<group>.*" to install all resources from a group (e.g. "network.*")
    // as well as "<group>.<kind> to install a specific resource only (e.g. "network.virtualnetworks") 
    Pattern []string `json:"pattern,omitempty"`
}

type InstalledResourceDefinitionsStatus struct {
    //Conditions: The observed state of the resource
    Conditions []conditions.Condition `json:"conditions,omitempty"`
}
```

This CRD _will_ be included in the Helm chart and will eventually allow users to configure which resources to install.
We will watch a magically named "aso-installed-resources" instance of this CRD and install CRDs according to the instructions.

### How should the `InstalledResourceDefinitions` CRD be consumed?

#### From within the existing container

**Pros:**
- Infrastructure is already there to watch CRDs, controller-runtime `Manager` is set up, etc.
- When the list of installed CRDs changes, the controller container probably has to restart to get the updated set of
  resources to watch and create clients + caches for them. This would be easier to orchestrate from within the same
  container as we could just coordinate a graceful exit.

**Cons:**
- When the container launches, it goes through the process of starting up watches on ALL the CRDs it's looking for.
  That includes all the standard ASO resources + the new `InstalledResourceDefinitions` CRD. The operator will fail
  to find the expected ASO resources and error out.
- `v2/cmd/controller/main.go` is already quite complicated and doing this there as well would add more complexity.

#### From another container within the operator pod

**Pros:**
- Separation of concerns from reconciliation of Azure resources.
- Avoids further complication of the already somewhat complicated `v2/cmd/controller/main.go`

**Cons:**
- Probably have to produce another docker image, including required image publishing infra work. Could be bypassed
  by just having a `mode` cmdline flag on `v2/cmd/controller/main.go` to run in CRD reconciliation mode.
- Harder to coordinate restart of operator container to refresh clients/caches.
- The main operator container still needs to read the `aso-installed-resources` CR to figure out which resources it 
  should watch. We could instead just avoid watch errors but that seems worse than knowing what we should be watching and
  erroring if we're unable to do so for some reason.

#### At operator pod start only

**Pros:**
- Easy

**Cons:**
- Not a true control-loop -- user would have to restart the pod manually after updating the set of installed CRDs for
  them to be picked up.

## Decision

- We will embed the CRDs in the operator and have a new controller manage their installation.
- We will manage the CRD installation and reconciliation process from within the existing ASO container.
- If the `aso-installed-resources` resource is missing, the operator pod will exit with an error. This is to ensure that
  we have a place to report status of CRD installations during upgrade and to ensure that we are driving to the expected goal
  state for the CRDs after an upgrade.
- The operator will operate on whatever ASO CRDs it finds as long as they are the expected version. If the CRD is not the
  expected version those resources will not be registered to watch (should trigger webhooks/etc to fail for them as well, need to confirm)
- The operator will add a label to the resources it adds, to make querying them easier.
- The operator will compare the existing CRD against its local copy and ensure that the *Spec* matches. I have a prototype 
  of this diffing which seems to work. We will ignore annotation differences and probably ignore every label except the 
  one we add above.
- The operator will tolerate the situation where the meta-resource reconciliation isn't done yet when it launches.

## Status
TBD

## Consequences
TBD 

## Reference

[Slack conversation about Helm limits](https://kubernetes.slack.com/archives/C0NH30761/p1677110151888539)

[Helmfile](https://github.com/helmfile/helmfile)

[Installing Google config connector doesn't support Helm](https://cloud.google.com/config-connector/docs/how-to/install-other-kubernetes)

[Crossplane supports Helm but only for the core, providers and their CRDs are installed another way](https://docs.crossplane.io/v1.11/getting-started/provider-aws/)