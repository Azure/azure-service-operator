---
title: '2022-11: Change Detection'
---

## Context

In ASO v2 up to at least the `beta.4` release, we reconcile each resource by doing a PUT, relying on the Azure Resource Manager (ARM) to do the goal state comparison and only update the resource if it has changed. 

While this works, customers are already running up against ARM throttling with moderate numbers of resources Typically, ARM throttles PUT requests for a given subscription to just 1200 per hour. With a reconcile period of 15m, ASO users are hitting this limit with just 300 active resources. 

In the `beta.4` release, we're lifting the default to 1 hour, but this is a compromise. It means we will detect and correct resource drift more slowly, and it's still not going to be enough for larger deployments. The change only impacts on steady state as well; we'll still hit throttling during initial deployments and when errors occur.

It's been previously suggested (see #1491) that we should do our own goal state comparison and only do a PUT if the resource has changed. This would allow us to reconcile more frequently, as GET requests have a significantly higher throttle limit.

While it sounds straightforward to compare a Status retrieved from Azure with the Spec held by ASO, there are a few complicating factors. (Credit to @matthchr for much of the wording in these descriptions, lifted from earlier discussions.)

### Readonly fields

Some fields are set by Azure and cannot be changed by the user. These fields *should* be marked as read-only in the Swagger, resulting in their omission from our generated Spec types. However, not all such fields are properly marked. 

Potential mitigation: We will likely have to use our configuration file to explicitly omit these fields from the comparison.

### Default values

Some fields are optional in the Swagger, but are required by Azure. In these cases, Azure will set a default value if the user does not provide one. In a Naive implementation, we'll see a difference between the Spec we have and the Status we retrieve from Azure. Complicating this, some resources may use a different default value depending on other fields. 

For example, GitHub CoPilot suggests the `publicNetworkAccess` field on a Storage Account defaults to `Enabled` if the `networkRuleSet` field is not set, but defaults to `Disabled` if it is set. 

*TODO: Verify this is true.*

While we can tell the different between a field that was never set by the user, and a field that we set to an explicit value (because all our Spec properties are pointers), this doesn't mean we can assume any explicit value is a default. For example, if a third party makes a change via the Azure portal, we want to detect this as a change and and revert it.

### Read-After-Write Consistency

Azure is not guaranteed to return the exact same resource that was PUT. For example, the `etag` field may change, or the `id` field may be returned with a different casing. Some resource providers also *normalize* field values, for example the CosmosDB API will return `West US` when the resource specified `westus`.

Potential mitigation: Generally speaking, Azure is expected to be "case-preserving, case-insensitive", so we should compare all string fields in a case insensitive manner. Region names and virtual machines SKUS are two known special cases where we may need to do some special case handling.

### Array ordering

Some resources have arrays of sub-resources. These arrays may not be guaranteed to be returned in the same order as they were specified. 

For example, GitHub CoPilot suggests the `networkAcls` field on a Storage Account has an array of `bypass` values. These are returned in a different order than they were specified.

*TODO: Verify this is true.*

Potential mitigation: Where the items in an array have a known identifier (easy for nested resources), use that identifier to match them up. Where the items in an array do not have a known identifier, compare them by index.

### Changes to test recordings

When a resource changes from PUT only to GET+PUT, we'll need to re-record the test results for that resource. To avoid having to re-record all the tests in one go, we probably want to have some way to migrate to the new approach in a controlled fashion. 

### Violation of the DRY principle

The really unfortunate thing, and the reason we actually like just the approach of always PUT-ing (with the obvious caveat that it causes throttling if you do it a lot) is that all of the logic to determine if there has been a change (what fields have what defaults, whether the field is stable or not, etc) all of that exists in the Resource Provider PUT path.

### Risk of bad change detection

If we get the change detection wrong, we could end up in a situation where we're reconciling a resource every 15m, but we're not actually making any changes to it. 

This would be bad for a number of reasons, but the most obvious is that we'd still be hitting the ARM throttling limits for no reason.

This is something that we can get *good enough* but that is gonna take iteration to get actually correct. Practically speaking it's hard to know what we're doing good/bad with diffs because for example with the normalization case if you look at something like MySQLServer, it has maybe 30 user-settable fields. If 1 of those fields normalizes certain values when set how are we ever gonna notice?

For properties we use in our recorded tests, we'll see issues when we go to re-record them, but for properties we don't use in our tests, we'll never know. 

This points to a need to log which fields are detected as changed, so that we can review the ASO logs and identify when we're getting the diff comparison wrong.

## Decision

TBC

## Status

Proposed

## Consequences

TBC

## Experience Report

TBC

## References

TBC
