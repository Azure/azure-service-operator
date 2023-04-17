---
title: '2023-02: Adoption Policy'
---

## Context

Resources which exist in Azure prior to being created in Kubernetes are "adopted" by the operator when the resource
is created in Kubernetes. Today this adoption process is automatic and implicit. As a user there isn't an easy way
to tell if a resource you're about to `kubectl apply` already exists in Azure (will be adopted) or does not exist in
Azure (will be created).

We already have support for [reconcile-policy](https://github.com/Azure/azure-service-operator/pull/2060) which
is related to adoption. Some of `reconcile-policy`'s primary use cases are for when the user does not want ASO to
fully manage an adopted resource.

The following are principles we should follow related to resource adoption:

1. **The operator should adopt resources by default:** This is standard goal-seeking/desired-state behavior. 
   If you create a resource in Kubernetes that says: "I expect there to be a Virtual Machine that looks like this
   at location X" and that Virtual Machine already exists, then we are in the desired state.
   The alternative would be to error if the resource already exists, which goes against the Kubernetes model.

2. **Users must be able to instruct the operator to stop making changes to the backing Azure resource:**
   This is enabled by `reconcile-policy`.

   This is useful for the following scenarios:
   - The operator is doing something wrong and the user needs a way to stop the operator from acting.
   - Pausing the operator temporarily while they make some manual fix via the Azure portal.
   - Creating a resource in ASO purely to get ownership references to work. The resource in question is not managed
     by ASO at all. It was created and managed some other way (portal, terraform, az-cli, etc).

3. **Users must be able to instruct the operator to skip deletion of the backing resource:**
   This is useful for scenarios such as moving ownership of resources from ASO in cluster A to ASO in cluster B,
   or as a safety net to protect stateful resources such as databases from accidental deletion.
   This is enabled by `reconcile-policy`.

### Open question: what should the default `reconcile-policy` be for resources which are adopted?

There are two main options.

#### Option 1: The default `reconcile-policy` for adopted resources is `manage`

`manage` is also the default `reconcile-policy` for resources ASO creates, so there is no difference between
creation and adoption with this option.
This means ASO automatically adopts, updates, and will delete backing Azure resources as normal unless the user
explicitly sets `reconcile-policy`.

**Advantages:**
- It's easy to explain what the behavior is. Simplistic users don't have to understand `reconcile-policy` (they won't
  ever see it unless they set it themselves).
- Matches ARM/BICEP, raw REST, and SDK behavior, as well as some of our competitors (Crossplane and Google Config Connector)
  as spelled out below.
- Is the current default

**Disadvantages:**
- Users have no good way to know their resource existed previously. They may `kubectl delete` it without realizing that
  other things depend on this resource. Deletion (and some mutation) is destructive in Azure so lost data, etc cannot 
  easily be recovered.

**Thoughts/comments:**

This sort-of matches the behavior of ARM templates and SDKs. If you call the `DELETE` API the resource is gone.

Some discussion of this topic has happened over in [ACK](https://github.com/aws-controllers-k8s/community/issues/82).

[Crossplane](https://github.com/crossplane/crossplane/issues/2590) doesn't support any notion of `reconcile-policy` at all.
`kubectl delete` always deletes the backing resource no matter what.

[GCP](https://cloud.google.com/config-connector/docs/how-to/managing-deleting-resources#creating_a_resource) does it this way:
> When you create a resource, Config Connector creates the resource if it doesn't exist. If a Google Cloud resource 
> already exists with the same name, then Config Connector acquires the resource and manages it.

GCP also [supports](https://cloud.google.com/config-connector/docs/how-to/managing-deleting-resources) a 
`reconcile-policy`-like annotation, so doing it like this would basically be doing it exactly like GCP.

Kubernetes has a similar concept for `StorageClasses`/`PersistentVolumes` called 
[reclaimPolicy](https://kubernetes.io/docs/concepts/storage/storage-classes/#reclaim-policy), whose default
is `Delete` for dynamically provisioned PVs, but `Retain` for manually created PVs.

#### Option 2: The default `reconcile-policy` for adopted resource is `abandon-on-delete`.

ASO detects that a resource previously existed and sets `reconcile-policy` automatically to `abandon-on-delete`.

**Advantages:**
- Defaults to a more conservative `reconcile-policy` to protect users from accidentally deleting things they may not
  have wanted to delete.

**Disadvantages:**
- Higher complexity. Users will be exposed to `reconcile-policy` as set by the operator (when a resource is adopted)
  so need to understand it.
- Bias towards safety may cause resource user thinks was deleted to continue existing (and billing) in Azure. Not a data
  loss though.
- `reconcile-policy` would have split ownership between user and operator, which means deciding if
  we should trigger reconciliation events on its change is tricky (we don't want our defaulting of it to trigger an event).

**Thoughts/comments:**

There's some discussion of a similar feature over in [ACK](https://github.com/aws-controllers-k8s/community/pull/1148#discussion_r794605821).
Jay Pipes had an interesting comment:
> In the case of ACK and the user experience of resource deletion, I posit that the behaviour of our interface has 
> no surprises and is actually hard to misuse. It doesn't have surprises because our interface to delete a resource 
> is to call kubectl delete on that resource. This is as non-surprising as it gets. It is hard to misuse because 
> calling kubectl delete is an explicit action the user has to take and it's not a confusing verb or semantic. 
> "delete" means to remove the resource.

It's important to note that ACK does adoption differently than ASO though (see 
[example](https://github.com/aws-controllers-k8s/sagemaker-controller#60-adopt-resources)), so adoption is much more
explicit for them. We don't want to have such explicit adoption though as covered in 
principle #1: The operator should adopt resources by default

CAPZ does something sort-of like this using Azure `tags`, but only for a select set of resources which they support
"BYO-resource" for. Things like VNET. Since some child resources don't support tags they apply the parent tag to
all of the children automatically, which makes sense for their use-case but doesn't make as much sense for ASO.

### Open question: Should the adoption behavior be configurable at the operator (global) or per-resource level?

Some users (such as CAPZ) may want adoption semantics different than the defaults we choose. Even on 
[the PR that added this design proposal](https://github.com/Azure/azure-service-operator/pull/2067) different users disagreed
on the behavior they wanted.

#### Option 3: No configuration at the operator or per-resource level

**Advantages:**
- Simple - value of this cannot be understated. The more configuration users _have_ to understand the harder the operator
  becomes to adopt or understand.

**Disadvantages:**
- Fails to cater to some legitimate use-cases.

#### Option 4: Allow configuration at the operator or per-resource level

This option would just be configuring the behavior of the operator in case it found an existing resource
in Azure. The behavior if ASO created the resource is unchanged.

The 3 scenarios here are, if the resource already exists in Azure when ASO goes to create it:
1. Fully adopt the resource, mutate it if its configuration is different in ASO than in Azure, and delete 
   it when `kubectl delete` is run.
2. Partially adopt the resource. Mutations of the resource would be performed by ASO but `kubectl delete` will only
   delete the resource in Kubernetes and _not_ in Azure.
3. Use the existing resource and don't make any changes to it to make it match the ASO shape, and do not delete it
   when `kubectl delete` is run.

This is discussed in [#2703](https://github.com/Azure/azure-service-operator/issues/2703)

##### Option 4.1: adoption-policy annotation

A new property `ADOPTION_POLICY` in the operator global secret, and an annotation on each
resource `serviceoperator.azure.com/adoption-policy` would allow users to control this behavior.
As with our other configurations, the more specific configuration would win.

Possible `adoption-policy`'s are:
- `adopt-if-exists`: Corresponding to scenario 1 above.
- `detach-on-delete-if-exists`: Corresponding to scenario 2 above.
- `skip-if-exists`: Corresponding to scenario 3 above.

This would exist alongside the existing `reconcile-policy` annotation. The most restrictive of the two options would
be used.
The full matrix of possibilities (where there is overlap) is:

| `reconcile-policy` | `adoption-policy`          | Resource already exists in Azure? | Behavior                                 |
| ------------------ | -------------------------- | --------------------------------- | ---------------------------------------- |
| Unspecified        | Unspecified                | Yes                               | Fully manages the resource¹              |
| Unspecified        | skip-if-exists             | Yes                               | Skips managing the resource              |
| Unspecified        | skip-if-exists             | No                                | Fully manages the resource               |
| detach-on-delete   | skip-if-exists             | Yes                               | Skips managing the resource              |
| detach-on-delete   | skip-if-exists             | No                                | Manages the resource, detaches on delete |
| skip               | skip-if-exists             | Yes                               | Skips managing the resource              |
| skip               | skip-if-exists             | No                                | Skips managing the resource              |
| Unspecified        | detach-on-delete-if-exists | Yes                               | Manages the resource, detaches on delete |
| Unspecified        | detach-on-delete-if-exists | No                                | Fully manages the resource               |
| detach-on-delete   | detach-on-delete-if-exists | Yes                               | Manages the resource, detaches on delete |
| detach-on-delete   | detach-on-delete-if-exists | No                                | Manages the resource, detaches on delete |
| skip               | detach-on-delete-if-exists | Yes                               | Skips managing the resource              |
| skip               | detach-on-delete-if-exists | No                                | Skips managing the resource              |

¹: Actual behavior depends on earlier sections of this design

**Advantages:**
- ???

**Disadvantages:**
- Users need to understand this complicated table to figure out what is actually happening to their resource, it's not
  immediately obvious based on the annotations.
- There's no way to tell which option was actually chosen for users.
- We need to somehow remember which option we chose so we don't think that a resource we created needs to be adopted after
  a crash or pod restart. That further increases the number of annotations used by this feature...

##### Option 4.2: reconcile-policy-if-exists annotation

A new property `RECONCILE_POLICY_IF_EXISTS` in the operator global secret, and an annotation on each
resource `serviceoperator.azure.com/reconcile-policy-if-exists` would allow users to control this behavior.

Possible `reconcile-policy-if-exists`'s are the same as the existing `reconcile-policy`
- `manage`: Corresponding to scenario 1 above.
- `detach-on-delete`: Corresponding to scenario 2 above.
- `skip`: Corresponding to scenario 3 above.

The implementation of `reconcile-policy-if-exists` would basically be a delayed defaulting of `reconcile-policy` applied
when reconciliation ran for the first time. The operator would check if the resource already exists and write the corresponding 
`reconcile-policy` annotation to the resource if the resource existed.

**Advantages:**
- Uses a concept we already have.

**Disadvantages:**
- Possibility of collision between user specified `reconcile-policy` and the `reconcile-policy` written by the operator because
  of `reconcile-policy-if-exists`. 
  - We can mitigate this issue by rejecting some configurations with a webhook, although we have to be careful how we craft
    that webhook because the operator itself needs to be able to set both. The easiest way to accomplish this would be with
    a webhook that rejects requests that include both  `reconcile-policy` and `reconcile-policy-if-exists` annotations 
    if the finalizer has not been added yet. Once the finalizer has been set, changing `reconcile-policy-if-exists` has no 
    effect, and all changes to the policy should flow via the `reconcile-policy` annotation.
  - An alternative mitigation would be that `reconcile-policy-if-exists` always overwrites `reconcile-policy` if the 
    resource exists. This has the advantage of allowing configurations like `reconcile-policy: skip` and 
    `reconcile-policy-if-exists: manage`, where a resource wouldn't be created if it didn't exist, but _would_ be adopted
    and managed if it did already exist. This would also allow `reconcile-policy: detach-on-delete` and 
    `reconcile-policy-if-exists: skip`. Possible downsides here include us modifying user-specified annotations. It's 
    possible that this will cause problems with GitOps workflows as the users checked-in spec will have 
    `reconcile-policy: <value>`, but that may have been overwritten by the operator dynamically. Subsequent applications
    of the resource would likely overwrite `reconcile-policy` back to its original value.
- Excess events generated by us setting `reconcile-policy`, which is an annotation we watch for updates. We can deal with this: 
  - If the `reconcile-policy-if-exists` is `skip` this is safe as we don't have to trigger reconciliation if the 
    only modification to the resource is that the `reconcile-policy` is being set to `skip`. This helps avoid a situation
    where we reconcile once, issue a GET to Azure, see that the resource exists and update the `reconcile-policy` annotation
    to be `skip`, which then triggers another reconciliation and another GET to Azure.
  - If the `reconcile-policy-if-exists` is `detach-on-delete` we're safe as those events only matter for delete 
    so also don't trigger reconciliation.

#### An implementation note on fault tolerance

If we choose either flavor of [option 4](#option-4-allow-configuration-at-the-operator-or-per-resource-level) we need to
ensure that we are fault-tolerant. We must avoid a situation where the operator:

- Checks if the resource in Azure exists and determines it does not
- Creates the resource in Azure
- Crashes or is restarted and does not ever persist to etcd the fact the resource didn't already exist
- Restarts and thinks that the resource already existed, triggering the adoption logic on a resource ASO itself just created

This means there must be an etcd write after checking if the resource exists, before we actually issue a PUT to Azure.
The existing `Claim` step in the operator reconcile path already caters to exactly this problem. 

Unfortunately because ARM resources don't universally support `If-Match` headers + `etag` we are open to a race where
the resource is created between the GET and the following PUT. There's nothing we can do about that without optimistic 
concurrency support on the server side.

## Decision

I propose we choose [option 1: Default `reconcile-policy` is `manage`](#option-1-the-default-reconcile-policy-for-adopted-resources-is-manage).
As mentioned in the discussion of that topic, it's the default we currently have and so is non-breaking to adopt, and is how
ARM templates/BICEP, raw REST and the SDKs work today already. Extra protection against adoption/updating a resource that
already exists can be achieved with some of those technologies, but it's not the default and requires the user to opt-in
to it.

Alongside that, I propose we choose [option 4.2: `reconcile-policy-if-exists` annotation](#option-42-reconcile-policy-if-exists-annotation)
to allow users to configure this behavior at the operator level or at the per-resource level. This is the other side of the
coin, if we're going to default to adopting we should give users who do _not_ want that as the default a way to configure
the behavior they do want.

As part of option 4.2, we should also implement one of the mitigations for the issue of `reconcile-policy-if-exists`
colliding with `reconcile-policy`. The obvious choice is the webhook rejecting overlap. While this limits certain scenarios
such as `reconcile-policy-if-exists: manage` + `reconcile-policy: skip`, there's no known use-case for these scenarios.
The advantage of rejecting things now is that we can always open it up and be less restrictive in the future if there is
customer need. If instead we allowed both to be set then it's hard to take it away.

## Status

Proposed.

## Consequences

TBC.

## Experience Report

TBC.

## References
