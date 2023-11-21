---
title: 2023-09 Complex Properties that Skip Versions
---

## Context

We have a long standing issue (originally documented in [#1776](https://github.com/Azure/azure-service-operator/issues/1776)) where our conversions between versions will break if a complex (object valued) property is reintroduced with a different shape in a later version after being removed in an earlier one.

**TL;DR:** The crux of the problem is that we can currently end up with two different versions (shapes) of the same property serialized in a property bag on the same intermediate version.

To clarify, consider the following examples, demonstrating the current behaviour, and then how that causes a problem. Following that, we'll discuss some constraints on possible solutions, and some approaches we might take.

### Conversion between v3 and v4

Consider a CRM system containing details of people. In **v3** of the system, we capture each person's residental address, but in **v4** we have dropped that property.

<!-- yuml.me

[v3.Person|FullName string;FamilyName string; KnownAs string; ResidentialAddress v3.Address|Properties PropertyBag]
[v3.Person]->[v3.Address|Label multilineString|Properties PropertyBag]
[v4.Person|FullName string;FamilyName string; KnownAs string|Properties PropertyBag]

-->

{{< figure src="./conversion-v3-v4-class.png" >}}

When we convert from **v3** to **v4** (**v3** -> **v4**) the `ResidentialAddress` property gets serialized into the `PropertyBag` on the **v4** Person. Note that the bag contains a **v3** `Address` in a serialized form.

To illustrate this, consider this concrete example:

<!-- yuml.me

[v3.Person|FullName: Michael Theodore Mouse;FamilyName: Mouse; KnownAs: Mickey; ResidentialAddress v3.Address|Properties PropertyBag]
[v3.Person]->[v3.Address|Label:\n1313 S. Harbor Blvd\nAnaheim\nCA 92803\nUSA|Properties PropertyBag]
[v4.Person|FullName: Michael Theodore Mouse;FamilyName: Mouse; KnownAs: Mickey|Properties PropertyBag]
[v4.Person]-.-[ResidentialAddress {bg:cornsilk}|Label:\n'1313 S. Harbor Blvd',\n'Anaheim',\n'CA 92803',\n'USA']

-->

{{< figure src="./conversion-v3-v4-instance.png" >}}

There is nowhere in the **v4** `Person` to store Mickey's residential address, so it gets safely stashed away in the `PropertyBag`.

In the other direction, from **v4** to **v3**, conversion works fine, as we can take the `ResidentialAddress` value from the property bag, deserialize the **v3** `Address`, and set things up as they were before.

Conversion back and forward between versions **v3** and **v4** works fine.

### Conversion between v4 and v5

In **v5**, the `ResidentialAddress` is reintroduced, but with a different shape. Instead of being the single field `Label`, it now has multiple fields.

<!-- yuml.me 

[v4.Person|FullName string;FamilyName string; KnownAs string|Properties PropertyBag]
[v5.Person|FullName string;FamilyName string; KnownAs string;ResidentialAddress|Properties PropertyBag]
[v5.Person]--[v5.Address|Street string; Suburb string; City string; Country string|Properties PropertyBag]

-->

{{< figure src="./conversion-v4-v5-class.png" >}}

When we convert from **v5** to **v4**, again the `ResidentialAddress` property gets serialized into the `PropertyBag` on the **v4** Person. 

Again, it's useful to see a concrete example:

<!-- yuml.me

[v4.Person|FullName: Michael Theodore Mouse;FamilyName: Mouse; KnownAs: Mickey|Properties PropertyBag]
[v4.Person]-.-[ResidentialAddress {bg:cornsilk}|Street: '1313 S. Harbor Blvd'\nSuburb: ''\nCity: 'Anaheim, CA 92803'\nCountry: 'USA']
[v5.Person|FullName: Michael Theodore Mouse;FamilyName: Mouse; KnownAs: Mickey; ResidentialAddress v3.Address|Properties PropertyBag]
[v5.Person]->[v5.Address|Street: 1313 S. Harbor Blvd\nSuburb:\nCity: Anaheim, CA 92803\nCountry: USA|Properties PropertyBag]

-->

{{< figure src="./conversion-v4-v5-instance.png" >}}

As before, there is nowhere in the **v4** `Person` to store Mickey's residential address, so it gets safely stashed away in the `PropertyBag`.

However, note that this time the bag contains a **v5** `Address` in a serialized form, **not** a **v3** `Address`.

In the other direction, from **v4** to **v5**, works fine, as we can deserialize the **v5** `Address` from the property bag.

Conversion back and forward between versions **v4** and **v5** works fine.

### The Problem

Observe how we can end up with two different variants of **v4** `Person`. In one case, we have a **v4** `Person` where the property bag contains a **v3** `Address`; in the other case, we have a **v4** `Person` the property bag contains a **v5** `Address`:

<!-- yuml.me 

[v4.Person|FullName: Michael Theodore Mouse;FamilyName: Mouse; KnownAs: Mickey|Properties PropertyBag]
[v4.Person]-.-[v3.ResidentialAddress {bg:cornsilk}|Label:\n'1313 S. Harbor Blvd',\n'Anaheim',\n'CA 92803',\n'USA']
[v4.Person]-.-[v5.ResidentialAddress {bg:cornsilk}|Street: '1313 S. Harbor Blvd'\nSuburb: ''\nCity: 'Anaheim, CA 92803'\nCountry: 'USA']

-->

{{< figure src="./conflict-v3-v5.png" >}}


Round trips between adjacent versions work fine, but we run into problems.

When we convert from **v3** to **v4** to **v5** we end up stuck part-way through.

After our first conversion step, we have a **v4** `Person` that contains a **v3** `Address` in the property bag. 

Conversion to the **v5** `Person` will fail when we try to deserialize a **v5** `Address` from the **v3** version in the property bag. 

We can't populate a **v3** `Address` from a serialized **v5* `Address`, the two are not compatible, and the conversion will fail. 

This will cause the operator to fail at runtime. 

In the opposite direction, from **v5** to **v4** to **v3**, we end up with a **v4** `Person` that contains a **v5** `Address` in the property bag, and a similar problem occurs. We can't populate a **v3** `Address` from a serialized **v5** `Address`, The two are not compatible, and the conversion will fail.


### Complications

#### We can't see the future

We don't know what will happen with future versions - so we may already have users consuming ASO with **v3** and/or **v4** resources prior to the release of **v5**. In fact, given that it's highly unusual for us to introduce multiple versions of a resource in a single release, this is almost a certainty.

However we choose to fix this problem, the introduction of a new resource version (**v5** in our example) has to be backward compatible with resources that already exist prior to that release. We can't make this a breaking change.

#### Multi-generational skips

The example shown here has the new property returning after just one version; we need to handle the case where the property returns after a hiatis of several versions.

#### Other PropertyBag use

The skipping-property problem is an edge case that emerges from the way we use the `PropertyBag` to maintain the conversion contract required by Kubernetes. 

When converting a resource to an earlier version, properties introduced in that later version get stored in the `PropertyBag` on the earlier version to ensure they're not lost.

Similarly, when converting a resource to a later version, properties removed from that later version get stored in the `PropertyBag` of the later version to ensure they are not lost.

### Option 1: Introduce Intermediate Type

Ensure that complex types stored in property bags have a consistent shape by introducing a dedicated type for that purpose. Regardless of whether the property bag is populated by an earlier version (e.g. a **v3** `Address`) or a later version (e.g. a **v5** `Address`), the shape of the object stored property bag will be the same (**v4**).

<!-- yuml.me
[v3.Person|FullName string;FamilyName string; KnownAs string;ResidentialAddress]
[v3.Person]--[v3.Address|Label string]
[v4.Person|FullName string;FamilyName string; KnownAs string|PropertyBag PropertyBag]..[v4.Address|Label string {bg:cornsilk}]
[v5.Person|FullName string;FamilyName string; KnownAs string;ResidentialAddress|PropertyBag PropertyBag] 
[v5.Person]--[v5.Address|Street string; Suburb string; City string|PropertyBag PropertyBag]
-->

{{< figure src="./option-1.png" >}}

We create the shape of the intermediate type by replicating the structure of the version _before_ (**v3**), essentially just copying it forward to the new version (**v4**).

When we convert from **v3** to **v4**, we first convert the **v3** `Address` into a **v4** `Address` before storing it in the property bag.

Then, when converting from **v4** to **v5**, we rehydrate a **v4** `Address` from the property bag and then convert that to a **v5** `Address`.

The process for conversion from **v5** to **v4** to **v3** is the reverse. We convert the **v5** `Address` into a **v4** `Address` before storing it in the property bag, and then rehydrate a **v4** `Address` from the property bag and convert that to a **v3** `Address`.

While this might seem to just _kick the can down the road_, our conversion framework already has support for customization of the conversions between adjacent versions of a type, allowing the mismatch between **v4** and **v5** of `Address` to be handled by augmenting the conversion.

Pros:
* Addresses the central inconsistency problem, ensuring we always serialize the same shape into intermediate property bags
* Preserves the serialization format for existing resource versions (e.g. **v3** and **v4**)

Cons:
* Changes the object model of released resources, a breaking change for consumers of the object model
* When there are multi-generational skips, other questions arise: Do we introduce the new type into every intermediate version, only the last, or both first and last? Each approach has advantages and drawbacks.

Fatal:
* The **v4** package already has a reference to **v5**. The Go compiler's rules prohibiting circular package dependencies preclude adding a reference to **v4** from **v5**.

### Option 2: Introduce Conversion Type

As for Option 1, but instead of introducing the new type into **v4** where it can't be referenced by **v5**, use a dedicated **v5/compat** subpackage.

{{< figure src="./option-2.png" >}}

<!-- yuml.me
[v3.Person|FullName string;FamilyName string; KnownAs string;ResidentialAddress]
[v3.Person]--[v3.Address|Label string]
[v4.Person|FullName string;FamilyName string; KnownAs string|PropertyBag PropertyBag]..[v4.Address|Label string {bg:cornsilk}]
[v5.Person|FullName string;FamilyName string; KnownAs string;ResidentialAddress|PropertyBag PropertyBag] 
[v5.Person]--[v5.Address|Street string; Suburb string; City string|PropertyBag PropertyBag]
-->

The invariant we need to maintain here is that the **v5/compat** version of `Address` needs to have the same exact shape as the **v3** version of `Address`.

Pros:
* Addresses the central inconsistency problem, ensuring we always serialize the same shape into intermediate property bags
* Preserves the serialization format for existing resource versions (e.g. **v3** and **v4**)
* Does not change the object model for released code
* Isolates types introduced for compatibility purposes in a dedicated subpackage, avoiding the problem of circular dependencies.

## Decision

TBC

## Status

TBC

## Consequences

TBC

## Experience Report

TBC

## References

TBC
``