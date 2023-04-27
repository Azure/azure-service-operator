---
title: '2023-04: Preventing Resource Drift'
---

## Context

Some Azure Resource Providers do not quite have full replace semantics for a PUT.

Instead, they have PATCH-like behavior on PUT which effectively boils down to: They will replace sub-objects/sub-collections if the object/collection is in the JSON payload, but they will leave untouched omitted properties or collections.

To illustrate, assume you have a Person object with a Name and a list of Addresses.

* If you PUT a Person object with a new Name and _no addresses_, the new name will be persisted, but the addresses will remain _unchanged_.
* If you PUT a Person object with a new Name and an _empty list_ of addresses, the new name will be persisted, and the addresses collection will be cleared.

This is a problem for ASO because we use `omitempty` on slice/map properties, which means that if you remove the last item from a slice/map, the property will be entirely omitted from the JSON payload.

In principle, an ASO custom resource (CR) should be the goal-state for the resource and should be a complete definition of the desired state. Whether the collection is empty or missing entirely should not matter, we should interpret that as "clear the collection".

### Option 1: Optional Properties

Tack an extra pointer onto every slice/map property across all generated types.

Instead of:

``` go
MyProperty []Type 'json:"myProperty,omitempty"'
```

we would generate:

``` go
MyProperty *[]Type 'json:"myProperty,omitempty"'
```

* PRO: Behaves exactly like we want from a serialization perspective.
* CON: Horrible to use in code because you must dereference the property to range over it (or really do much of anything with it).
* CON: Breaking change for any consumers of our object model.
* CON: Modifies our Spec and Status object definitions for an issue that is specific to ARM

### Option 2: Remove `omitempty` 

Remove omitempty from slice/map properties across all generated types.

``` go
MyProperty []Type 'json:"myProperty"'. 
```

This means that the slice or map will always be serialized, even if empty.

* PRO: Behaves exactly like we want from a serialization perspective.
* PRO: Not a breaking change for consumers of our object model.
* CON: Empty collections will always be present when an ASO CR is read from the cluster, even if the user did not specify them at all.
* CON: If we extend this to all properties, all properties will be present when an ASO CR is read from the cluster, even if the user did not specify them at all.
* CON: Modifies our Spec and Status object definitions for an issue that is specific to ARM
* QUERY: It is possible that some Azure services treat `myProperty: null` as different than omitting the property entirely, although we do not actually know of any cases where that is true.

### Option 3: Custom JSON Serialization

Use a custom JSON serialization library that supports an `omitnil` equivalent, such as [jettison](https://github.com/wI2L/jettison).

* PRO: Behaves exactly like we want from a serialization perspective.
* PRO: Not a breaking change for consumers of our object model.
* CON: Breaks CRD YAML deserialization for the ASO controller (as the [sigs.k8s.io/yaml]( https://github.com/kubernetes-sigs/yaml) library we use works with standard JSON serialization attributes)
* CON: Also breaks YAML serialization for `asoctl` (for the same reason)
* CON: Modifies our Spec and Status object definitions for an issue that is specific to ARM

### Option 4: Remove `omitempty` for selected ARM types only

As for Option 2, above, we apply the change only selected ARM Spec types, leaving the core ASO Spec and Status types unmodified.

Selection occurs through our existing `ObjectModelConfiguration` allowing us to explicitly target collections (and other properties) that we know are problematic.

* PRO: Behaves exactly like we want from a serialization perspective.
* PRO: Limits the scope of the change to only ARM types. These are not directly used by consumers of our object model.
* PRO: Limits the scope of the change to only properties that we know are problematic, reducing any potential for unintended consequences.
* PRO: No change to the YAML serialization for ASO Spec and Status types, so no visible change to our users.
* CON: We can only configure properties we know have this problem. We may miss some, requiring a new release to fix affected users.

## Decision

Adopt Option 4, as it is a minimal change that addresses the issue without impacting on current users.

## Status

Proposed.

## Consequences

TBC

## Experience Report

TBC

## References

One of our users encountered this problem with the `virtualNetworkRules` and `ipRules` properties on Azure Storage accounts and reported it in issue [#2914](https://github.com/Azure/azure-service-operator/issues/2914).

@Matthchr's response was:

> Azure Storage accounts don't quite have PUT == full replace semantics. They have a PATCH-like behavior on PUT which effectively boils down to: They will replace sub-objects/sub-collections if the collection is in the JSON payload, but they will leave untouched properties which were omitted.
> 
> When you remove virtualNetworkRules or ipRules, what actually happens on the backend is those properties are nil and don't get serialized. This falls afoul of item 1 above because since those properties were not included in the JSON payload storage handles them with PATCH semantics and doesn't clear/remove the contents of the collection.
> 
> You can see some evidence of this behavior in the az storage account network-rule remove implementation, where if you read closely, removing the last item from the list of rules will serialize an empty collection to storage.

> It's this empty collection which is the key to tell storage: "I really want you to clear this property please". I've confirmed this by contacting the Storage team as well.

> It would be nice (if not ideal) if you could just specify an empty collection ([]) for the virtualNetworkRules property to work around this issue... but you can't right now because ASO tacks Go's omitempty onto properties, and unfortunately doing that means that it's impossible for Go to serialize [] to the JSON payload. It will omit the property entirely if the collection is nil or [].

Credit where it is due; I've liberally reused Matt's wording while writing this ADR.



