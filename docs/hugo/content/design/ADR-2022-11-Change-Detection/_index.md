---
title: '2022-11: Change Detection'
---

## Context

In ASO v2 up to at least the v2.15 release, we reconcile each resource by performing a PUT operation, relying on the Azure Resource Manager (ARM) to do goal state comparison and to only update the resource if it has changed.

While this approach works, some customers are running into ARM throttling limits with moderate numbers of resources. This occurs despite ARM throttling limits being lifted substantially since ASO went GA.

There are also a limited number of resources (AKS AgentPools being one) where a PUT operation is not strictly idempotent. In these cases, doing a PUT when no change is needed may have detectable side effects, even if minor.

The current ASO default reconciliation interval is 1 hour. This is a compromise that means we will detect and correct resource drift more slowly than we might like, and has proven insufficient to prevent throttling for larger deployments.

It has been previously suggested (see [#1491](https://github.com/Azure/azure-service-operator/issues/1491)) that we should do our own goal state comparison and only perform a PUT if the resource has changed. This would allow us to reconcile more frequently, as GET requests have a significantly higher throttle limit. According to [_Understand how Azure Resource Manager throttles requests_](https://docs.azure.cn/en-us/azure-resource-manager/management/request-limits-and-throttling), the default write limit is just 1,200 per hour, while the default read limit is 12,000 per hour.

Performing a GET followed by a PUT will also better align with user expectations—we've had multiple issues raised where users have been surprised that we perform a PUT for every reconciliation.

While it sounds straightforward to compare a _Status_ retrieved from Azure with the _Spec_ held by ASO, there are several complicating factors. (Credit to @matthchr for much of the wording in these descriptions, lifted from earlier discussions.)

### Issue: Read-only fields

Some fields are set by Azure and cannot be changed by the user. These fields _should_ be marked as read-only in the Swagger specification, resulting in their omission from our generated Spec types. However, not all such fields are properly marked.

**Potential mitigations**:

* We already attempt to prune these from our spec types by manually listing them in our configuration file, and we create PRs to update the upstream Swagger when practical. Currently, the impact of missing one of these is minor, but if/when we move to GET+PUT, this will result in change being detected every time we reconcile, resulting in unnecessary PUTs. This isn't ideal, but it's no worse than our current baseline.
* We already have generated code to convert between Status and Spec types (added for `asoctl import azure resource`), so we can use those conversions to create a `spec` that represents the resource as it exists in Azure.

### Issue: Default values

Some fields are optional in the Swagger specification but required by Azure. In these cases, Azure will set a default value if the user does not provide one. In a naive implementation, we'll see these fields as differences between the _Spec_ we have and the _Status_ we retrieve from Azure. Complicating this further, some resources may use different default values depending on other fields. (We know this is the case for AKS, for example.)

While we can tell the difference in a Spec between a field that was never set by the user and a field that was set to an explicit value (because all our Spec properties are pointers), this doesn't mean we can assume any explicit value in a Status is a default. For example, if a third party makes a change via the Azure portal, we want to detect this as a change and revert it.

A variation of value defaulting may occur with arrays and maps, where some services will add a default value to the collection if the user does not specify it. Conceptually, you can write `["a"]` and get back `["default", "a"]` when you read it back.

**Potential mitigation**: Keep track of the changes made by Azure when we originally PUT the resource, and use those as another input to our change detection process. For most resources, the current state of the resource is returned in the response to either the PUT or the long-running-operation (LRO) polling operation, so we can capture these changes at that point. For others, we may need to perform an initial GET after the PUT to capture these changes, despite the inherent race condition of doing so.

### Issue: Read-After-Write Consistency

Azure is not guaranteed to return the exact same resource that was PUT.

* Some services normalize identifier fields (converting them to all lowercase)
* Some fields accept a wider range of values than they return.

For example, the CosmosDB API will return the region `West US` when the resource specified `westus`.

**Potential mitigations**:

* Generally speaking, Azure aims for resource providers to be "case-preserving, case-insensitive", but this isn't uniformly applied. We should compare those string fields in a case-insensitive manner.
* Region names and virtual machine SKUs are two known special cases where we may need to implement special case handling.
* Capturing the changes (as detailed above) would also work for this issue.

### Issue: Array ordering

Some resources have arrays of items. These arrays may not be guaranteed to be returned in the same order as they were specified.

**Potential mitigations**:

* In the common case where the items in an array have a known identifier (e.g., when they are sub-resources with an `id` field), use that identifier to match them up. Where the items in an array do not have a known identifier, compare them by index. Many resources include `x-ms-identifiers` in their Swagger specification to indicate which fields are identifiers; we may be able to leverage this.
* Add specific configuration for array properties to specify how they should be compared — whether by an id field or by index. Consistent with our similar configuration used elsewhere, we'd require every array property to be annotated to ensure we don't do the wrong thing by default. This would be very verbose.

### Issue: Changes to test recordings

When a resource changes from PUT-only to GET+PUT, we'll need to re-record the test results for that resource. To avoid having to re-record all the tests in one go, we probably want some way to migrate to the new approach in a controlled fashion.

**Potential mitigation**: We could have a configuration flag that allows us to switch between the two approaches for a specific resource. We'd make GET+PUT the default, explicitly configure all our existing resources as PUT-only, and migrate in a controlled fashion, prioritizing those resources where we know we have the most issues with the PUT-only approach.

This will also allow us to leave problematic resources using only PUT if we find that it's difficult or cumbersome to use GET+PUT because of issues with the comparison logic.

### Issue: Violation of the DRY principle

The really unfortunate aspect — and the reason we actually like the approach of always PUT-ing (with the obvious caveat that it causes throttling if done frequently) — is that all the logic to determine if there has been a change (what fields have what defaults, whether the field is stable or not, etc.) already exists in the Resource Provider PUT path.

With that functionality present, however, we don't need our change detection to be perfect, we just need it to be good enough that we don't end up doing unnecessary PUTs most of the time.

### Issue: Risk of bad change detection

If we get the change detection wrong, we could end up in a situation where we're using PUT to reconcile a resource every 15 minutes, but we're not actually making any changes to it.

This would be problematic for several reasons, but the most obvious is that we'd still be hitting the ARM throttling limits for no reason.

This is something that we can get _good enough_, but it will take iteration to get actually correct. Practically speaking, it's hard to know what we're doing well or poorly with diffs because, for example, with the normalization case, if you look at something like MySQLServer, it has maybe 30 user-settable fields. If one of those fields normalizes certain values when set, how would we ever notice?

For properties we use in our recorded tests, we'll see issues when we go to re-record them, but for properties we don't use in our tests, we'll never know.

This points to a need to log which fields are detected as changed, so that we can review the ASO logs and identify when we're getting the diff comparison wrong.

### Issue: Storage of known-safe differences

If we are going to track the differences between resource-as-PUT and resource-as-created-in-Azure, we need to store those differences somewhere. For most resources, this will be a very small delta, but we need to allow for the possibility that it could be large.

**Possible mitigations**:

* Store as an annotation on the resource. This is simple but has a distinct size limit (256KiB for _all annotations combined_). This is probably sufficient for most resources, but we need to be aware of the limit and decide what to do if we exceed it.

* Capture a hash of the diff between resource-as-PUT and resource-as-created-in-Azure and store that as an annotation. This will have a fixed size but means we won't have the actual diff itself. Any change at all will result in a different hash, so we'll likely be performing a PUT to reconcile more often, and we won't be able to log the actual differences for troubleshooting purposes.

* Add `operatorStatus` to the status of the resource and use that to capture a map of the changes found between resource-as-PUT and resource-as-created-in-Azure. This avoids the size limit of annotations and makes differences easy to inspect and log.

### Issue: Performance

What's the CPU overhead of doing a GET and diff on every reconciliation compared to just doing a PUT?

We'll be adding another network round trip (with the associated bandwidth and latency costs) for every reconciliation, plus the CPU cost of doing the diff itself. We need to ensure that this doesn't cause us to exceed our resource limits in the controller pod, especially when reconciling large numbers of resources.

### Issue: ARM References

References to other resources can be specified in a number of ways - as an `ARM ID` in Azure, by specifying Group/Kind/Name of the resource within the cluster, or (for some resources) by specifying a well known name.

When we do a GET to retrieve the current state of a resource, any references will be returned as `ARM IDs`. We need to ensure that our diffing logic can handle this and correctly identify when a reference has or has not changed.

## Decision

Recommendation: Implement differencing as follows:

To detect changes:

* Perform a GET to retrieve the current state of the resource from Azure.
* Convert the returned Status to a Spec using our existing generated conversion code.
* Compare the original Spec (the desired state) with the converted Spec (the actual state) to identify current differences.
* Reconcile that set of differences with a set of _known differences_ that we have previously recorded.
* If any new differences are found, log them (so we know why the PUT was triggered) and perform a PUT to reconcile the resource to the desired state.

When creating or updating a resource:

* After performing a PUT (or completing the LRO polling), capture the returned Status on the resource (this already happens).
* Convert that Status to a Spec using our existing generated conversion code.
* Compare the original Spec (the desired state) with the converted Spec (the actual state) to identify _known differences_ for the resource, capturing any normalizations or default values applied by Azure.
* Store those _known differences_ in the operatorStatus on the resource.

To support this functionality:

* Update our code generator to add a `DiffWith()` method to every spec type that accumulates a list of the differences found.
* Add helper methods to a package `genruntime/diff` to help with common comparison tasks.
* Introduce `operatorStatus` on each top level status type, with a `knownDifferences` field to capture the differences found.

Migrate to the new approach by:
  
* Making GET+PUT the default reconciliation approach for all resources.
* Add configuration to allow PUT-only reconciliation to be specified for selected resources.
* Opt-in all existing resources to use PUT-only so their behaviour doesn't change until we can validate the new approach.

### Outstanding questions

* Can we come up with a better name than `knownDifferences`?

## Status

Proposed

## Consequences

TBC

## Experience Report

TBC

## References

TBC
