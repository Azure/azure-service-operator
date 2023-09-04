---
title: 2023-09 Complex Properties that Skip Versions
---

## Context

We have a long standing issue (originally documented in [#1776](https://github.com/Azure/azure-service-operator/issues/1776)) where our conversions between versions will break if a complex (object valued) property is reintroduced in a later version after being removed in an earlier one.

**TL;DR:** The crux of the problem is that we can currently end up with two different versions (shapes) of the same property serialized in a property bag on the same intermediate version.

To clarify, consider the following examples, demonstrating the current behaviour, and then how that causes a problem. Following that, we'll discuss some constraints on possible solutions, and some approaches we might take.

### Version 4

Consider a CRM system containing details of people. In **v3** of the system, we capture each person's residental address, but in **v4** we have dropped that property.

<!-- yuml.me
[Person v3|FullName string;FamilyName string; KnownAs string;ResidentialAddress]
[Person v3]--[Address v3|Label string]
[Person v4|FullName string;FamilyName string; KnownAs string]
-->

{{< figure src="version-4.png" >}}

When we convert from **v3** to **v4** (**v3** -> **v4**) the `ResidentialAddress` property gets serialized into the `PropertyBag` on the **v4** Person. Note that the bag contains a **v3** `Address` in a serialized form.

In the other direction, from **v4** to **v3**, works fine, as we can deserialize the **v3** `Address` from the `PropertyBag`.

Conversion back and forward between versions **v3** and **v4** works fine.

### Version 5

In **v5**, the `ResidentialAddress` is reintroduced, but with a different shape. Instead of being the single field `Label`, it now has multiple fields.

<!-- yuml.me 
[Person v4|FullName string;FamilyName string; KnownAs string|PropertyBag PropertyBag]
[Person v5|FullName string;FamilyName string; KnownAs string;ResidentialAddress|PropertyBag PropertyBag]
[Person v5]--[Address v5|Street string;Suburb string;City string|PropertyBag PropertyBag]
-->

{{< figure src="version-5.png" >}}

When we convert from **v5** to **v4**, again the `ResidentialAddress` property gets serialized into the `PropertyBag` on the **v4** Person. Note that this time the bag contains a **v5** `Address` in a serialized form, **not** a **v3** `Address`.

In the other direction, from **v4** to **v5**, works fine, as we can deserialize the **v5** `Address` from the property bag.

Conversion back and forward between versions **v4** and **v5** works fine.

### The Problem

Observe how we can end up with two different variants of **v4** `Person`. We can end up with one where the property bag contains a **v3** `Address` or we can end up with one where the property bag contains a **v5** `Address`.

Here is where we run into problems.

<!-- yuml.me
[Person v3|FullName string;FamilyName string; KnownAs string;ResidentialAddress]
[Person v3]--[Address v3|Label string]
[Person v4|FullName string;FamilyName string; KnownAs string|PropertyBag PropertyBag]
[Person v5|FullName string;FamilyName string; KnownAs string;ResidentialAddress|PropertyBag PropertyBag]
[Person v5]--[Address v5|Street string;Suburb string;City string|PropertyBag PropertyBag]
-->

{{< figure src="problem.png" >}}

When we convert from **v3** to **v4** to **v5** we end up stuck part way through. 

After our first conversion step, we have a **v4** `Person` that contains a **v3** `Address` in the property bag. 

Conversion to the **v5** `Person` will fail when we try to deserialize a **v5** `Address` from the **v3** version in the property bag. We can't populate a **v3** `Address` from a serialized **v5* `Address`, the two are not compatible, and the conversion will fail. This will cause the operator to fail.

In the other direction, from **v5** to **v4** to **v3** we end up with a **v3** `Person` that contains a **v5** `Address` in the property bag, and a similar problem occurs. We can't populate a **v5** `Address` from a serialized **v3** `Address`, The two are not compatible, and the conversion will fail.


### Complications

#### We can't see the future

We don't know what will happen with future versions - so we may already have users consuming ASO with **v3** and/or **v4** resources prior to the release of **v5**. In fact, given that it's highly unusual for us to introduce multiple versions of a resource in a single release, this is almost a certainty.

However we choose to fix this problem, the introduction of a new resource version (**v5** in our example) has to be backward compatible with resources that already exist prior to that release. We can't make this a breaking change.

#### Multi-generational skips

The example shown here has the new property returning after just one version; we need to handle the case where the property returns after a hiatis of several versions.

#### Type changes

This example shows the shape of the complex type changing between versions. 

Consider how things would change if the `ResidentialAddress` on **v3** `Person` was a simple string - instead of the property bag containing an object to deserialize, the **v5** `Address` would find a simple string.

### Proposed Solution

Ensure that complex types stored in property bags have a consistent shape by introducing a dedicated type for that purpose.

<!-- yuml.me
[Person v3|FullName string;FamilyName string; KnownAs string;ResidentialAddress]
[Person v3]--[Address v3|Label string]
[Person v4|FullName string;FamilyName string; KnownAs string|PropertyBag PropertyBag]-.-[Address v4|Label string {bg:cornsilk}]
[Person v5|FullName string;FamilyName string; KnownAs string;ResidentialAddress|PropertyBag PropertyBag] 
[Person v5]--[Address v5|Street string; Suburb string; City string|PropertyBag PropertyBag]
-->

{{< figure src="solution.png" >}}

We replicate the structure of the version _before_ when synthesizing the intermediate type.

When we convert from **v3** to **v4**, we first convert the **v3** `Address` into a **v4** `Address`  before storing it in the property bag.

Then, when converting from **v4** to **v5**, we rehydrate a **v4** `Address` from the property bag and then convert that to a **v5** `Address`.

The process for conversion from **v5** to **v4** to **v3** is the reverse. We convert the **v5** `Address` into a **v4** `Address` before storing it in the property bag, and then rehydrate a **v4** `Address` from the property bag and convert that to a **v3** `Address`.

While this might seem to just _kick the can down the road_, our conversion framework already has support for customization of the conversions between adjacent versions of a type, allowing the mismatch between **v4** and **v5** of `Address` to be handled by augmenting the conversion. 

#### Backward Compatibility

With the shape of the introduced type being the same as the shape of the type in the version _before_ the skip, deserialization will _just work_. 


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