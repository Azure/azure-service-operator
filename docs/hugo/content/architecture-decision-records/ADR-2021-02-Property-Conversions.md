---
title: '2021-02: Property Conversions'
---

## Context

To facilitate the use of dedicated storage variants for persistence of our custom resources, we need to codegen conversion routines that will copy all properties defined on one version of a resource to another version.

Given the way resources evolve from version to version, these we need to support a wide range of conversions between types that are similar, but not identical.

For example, looking primitive types (such as **string**, **int**, and **bool**):

A **primitive type can become optional** in a later version when a suitable default is provided, or when an alternative becomes available.

An **optional primitive type can become mandatory** in a later version when deprecation of other properties leaves only one way to do things.

A **primitive type can be replaced by an enumeration** when validation for a property (such as VM SKU) is tightened up to avoid problems.

A **primitive type can be replaced by an alias** when additional validation constraints (e.g., limiting bounds or a regular expression) are added to address problems.

These transformations (and others) can occur in combination with each other (e.g., **a primitive type replaced by an optional enumeration**) and with other constructs, such as maps, arrays, resources, and complex object types.

Other constraints apply, such as the need to **clone optional values during copying** lest the two objects become coupled with changes to one being visible on the other.

When implementation began, it very quickly became apparent that independently addressing every possible conversion requirement would be difficult to impossible given the combinatorial explosion of possibilities.

## Decision

Use a recursive algorithm to generate the required conversion by composing simpler conversion steps together.

For example, given a need to copy an optional string to an optional enumeration, the process works as follows:

Original problem: `*string -> *Sku` (where Sku is based on a string)

The handler `assignFromOptional` knows how to handle the optionality of `*string`, reducing the size of the problem. A recursive call is made to solve the new problem.

Reduced problem: `string -> *Sku`

The handler `assignToOptional` knows how to handle the optionality of `*Sku`, reducing the size of the problem further. A recursive call is made to solve the new problem.

Reduced problem #2: `string -> Sku`

The handler `assignToAliasedPrimitive` recognizes that `Sku` is an enumeration based on `string` and reduces the problem. A recursive call is made to solve the new problem.

Reduced problem #3: `string -> string`

Now `assignPrimitiveFromPrimitive` can handle the reduced problem, generating a simple assignment to copy the value across:

``` go
destination = source
```

Working backwards, the handler `assignToAliasedPrimitive` injects the required cast to the enumeration

``` go
destination = Sku(source)
```

Then, `assignToOptional` injects a local variable and takes its address

``` go
sku := Sku(source)
destination = &sku
```

Finally, `assignFromOptional` injects a check to see if we have a value to assign in the first place, assigning a suitable zero value if we don't:

``` go
if source != nil {
    sku := Sku(source)
    destination := &sku
} else {
    destination := ""
}
```

## Status

Successfully implemented. First commit in [PR# #378](https://github.com/Azure/k8s-infra/pull/378)

The full list of implemented conversions can be found in `property_conversions.go`.

## Consequences

It required some finessing of the conversion code to generate high quality conversions; early results were functional but not comparable with handwritten efforts.

As we’ve encountered new cases where new transformations are required, the list of conversions has been extended with additional handlers, including:

•	Support for enumerations ([PR #392](https://github.com/Azure/k8s-infra/pull/392))
•	Conversion of complex object types ([PR #395](https://github.com/Azure/k8s-infra/pull/395))
•	Support for aliases of otherwise supported types ([PR #433](https://github.com/Azure/k8s-infra/pull/433))
•	Ability to read values from functions ([PR #1545](https://github.com/Azure/azure-service-operator/pull/1545))
•	Support for JSON properties ([PR #1574](https://github.com/Azure/azure-service-operator/pull/1545))
•	Property Bags for storing/recalling unknown properties ([PR# 1682](https://github.com/Azure/azure-service-operator/pull/1682))

## Experience Report

TBC

## References

TBC
