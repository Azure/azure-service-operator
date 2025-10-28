---
title: '2025-10: The Long Tail of Property Conversion'
toc_hide: true
---

## Context

We are successfully generating conversion code for inter-version transformation of resources across a wide range of resources, but are now running into rare occassions where the existing framework doesn't fully handle the required conversion.

The most recent example affects preview versions of [`ManagedClustersAgentPool`]({{< relref "containerservice/v1api20240402preview/_index.md" >}}#ManagedClustersAgentPool) (a child resource of [`ManagedCluster`]({{< relref "containerservice/v1api20240402preview/_index.md" >}}#ManagedCluster)) where the cardinatlity of [ScaleProfile]({{< relref "containerservice/v1api20240402preview/_index.md" >}}#ScaleProfile) has changed from plural (`AutoScaleProfile[]`) to singular (`AutoScaleProfile`).

In the past, as each case of incompatibility has arisen, we've enhanced the code generator to handle the new case, taking the stance that it's better to handle these cases once (by enhancing the generator) than to handle them many times (by writing custom conversion code for each resource that needs it).

However, have we reached the point of diminishing returns? Each new case we handle requires more complexity in the generator, and the frequency of new cases is now very low.

### Option 1: Continue to extend the generator

As new cases that aren't yet handled arise, enhance the code generator to handle the required conviersions automatically. This is the status quo approach.

#### Pros

* Once a case is handled, it is handled for all resources, so we don't need to write custom conversion code for each resource that needs it.

#### Cons

* Each new case adds complexity to the generator, making it harder to maintain and reason about.
* The frequency of new cases is low, so the cost of enhancing the generator may outweigh the benefits.

### Option 2: Add an opt-out mechanism for custom conversion

When a resource requires a property conversion that the generator doesn't handle, allow configuration to specify that custom conversion code will be provided for that specific property. The generator will then skip generating conversion code for that property, allowing the custom code to take precedence.

#### Pros

* Avoids adding complexity to the generator for rare cases.
* Allows for targeted custom conversion code only where needed.
* We already have the `augment*` hooks for custom conversion code, so this approach builds on existing mechanisms.

#### Cons

* Requires writing custom conversion code for each resource that needs it, which could lead to more code to maintain.
* May lead to inconsistencies if different resources implement custom conversion differently (though, policy differences between different situations may make this an advantage).


## Decision

Add an opt-out mechanism for custom conversion for specific properties when the generator doesn't handle the required conversion.

Assess each case as it arises, and if it is rare enough that enhancing the generator isn't justified, use the opt-out mechanism to provide custom conversion code for that property. For more common cases, continue to enhance the generator as we have in the past.

### Suggested Configuration

Enhance the existing `objectModelConfiguration` structure to allow specifying `$conversionStrategy` at the property level. Permissible values would be:

* `auto` - Default behavior, generate conversion code automatically.
* `custom` - Skip generating conversion code for this specific property only, but generate a type assertion to force implementation of the appropriate `augment*` hook (this will help to avoid silent failures where a developer forgets to implement the custom conversion code).
* `none` - Skip generating conversion code for this property without forcing implementation of the `augment*` hook (useful for properties that don't require conversion at all).

## Status

Proposed.

## Consequences

## Experience Report

## References

* [Resource Versioning]({{< relref "design/versioning/_index.md" >}})
