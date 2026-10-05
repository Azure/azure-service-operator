---
title: '2026-04: Optional Secrets'
toc_hide: true
---

## Context

We have explicit configuration to control whether a property is considered a secret or not - we can specify `$isSecret`
as either `true` or `false`.

Specifying `$isSecret: true` has been required because the Azure REST API specs don't always call out which properties
are secrets; our config allows us to address omissions.

We've also used some heuristics to detect properties which might be secrets, requiring an explicit decision to be made.
Supporting `$isSecret: false` allows us to configure false positives in these cases.

To date, this has covered all of our needs - but we now have an interesting case (identified in
[#5095](https://github.com/Azure/azure-service-operator/issues/5095)) where a property is _sometimes_ a secret but not
always.

For some users, a webhook service URI is a plain text field with nothing sensitive.

But for others, the URI necessarily contains credentials (a password, api key, or similar) and they want to treat it as
a secret, pulling it from a secret at runtime.

## Requirements

To support this, we need more flexibility than a simple true/false answer to the question "is this property a secret?" -
we need the ability to handle a property which is an optional secret, i.e. it can be treated as a secret if the user
chooses to do so, but it doesn't have to be.

For properties already marked as secrets, we want to continue treating them as secrets, preserving the existing
generated CRD structure exactly as it is. No breaking changes.

We want to keep configuration, code generation, and CRD use as simple as possible, but are willing to accept some
complexity if it leads to a cleaner overall result, expecially if it's simpler for end users.

## Options

The problem falls into two natural parts:

* How do we represent this in our configuration?
* How do we transform the property in our code generator?

## Configuration Option 1: Additional Property

Add new custom configuration for this situation, something mutually exclusive with `$isSecret`, say `$isOptionalSecret:
true/false`.

### Pros

* Minimal Change

### Cons

* Adds complexity to the configuration, and to the code which processes it.
* Doesn't really capture the idea that this is a single property with three states (secret, not secret, or optional
secret) - instead it has two properties which are mutually exclusive, which is a bit awkward.

## Configuration Option 2: Trivalued Boolean

Keep the existing true/false property but allow it to take a third value, say `optional`, to indicate that the property
is an optional secret.

### Pros

* Simpler configuration - just one property to indicate the secret status of a property, with three possible values.
* More directly captures the idea that this is a single property with three states.

### Cons

* More significant change to the configuration and code, as we need to support a new value for an existing property.
* Significantly less intuitive for users who are used to thinking of `$isSecret` as a boolean - they now need to
understand that it can take a non-boolean value.

## Configuration Option 3: Change to new Enum

Replace the existing `$isSecret` boolean with a new enum property, say `$importSecretMode` with values `Required`,
`Never`, and `Optional`.

### Pros

* Cleanest solution - we have a new property that clearly indicates the three possible states without overloading the
meaning of an existing property.
* Avoids confusion for users who are used to thinking of `$isSecret` as a boolean, as we have a new property with a
clear name that indicates its purpose.
* Well defined in scope, potentially achievable with Copilot.

### Cons

* Most significant change to the configuration and code, as we need to replace an existing property with a new one and
update all references to it.
* Requires users to update their configuration to use the new property, which is a breaking change to any work currently
in progress.

## Configuration Option 4: Change to new Enum, with backward compatibility

Replace the existing `$isSecret` boolean with a new enum property, say `$importSecretMode` with values `Required`,
`Never`, and `Optional`. Handle `$isSecret` as a legacy property for backward compatibility, mapping `true` to
`Required` and `false` to `Never`.

### Pros

* Provides a clear path forward with the new enum property while maintaining backward compatibility for existing
configurations.
* Allows users to transition to the new configuration at their own pace without breaking existing work.
* Clean solution that clearly indicates the three possible states without overloading the meaning of an existing
property
* Avoids confusion for users who are used to thinking of `$isSecret` as a boolean, as we have a new property with a
clear name that indicates its purpose.

### Cons

* Still requires a significant change to the configuration and code to introduce the new enum property and handle the
legacy `$isSecret` property.
* Adds some complexity to the code to support both the new and legacy properties, though this is mitigated by the clear
mapping between them.
* Users will need to update their configuration to use the new property eventually, which is a breaking change to any
work currently in progress, though they can continue using the old property in the meantime.

## Generation Option 1: Parallel Properties

As we've done for properties that can be optionally populated from ConfigMaps, we could have two properties in the
generated code - one for the plain text value and one for the secret reference. They'd be mutually exclusive, and the
user would populate one or the other.

For consistency with the way we've handled properties that can be optionally populated from ConfigMaps, we could use the
suffix `Secret` for the secret reference property

### Pros

* Consistent with existing patterns for optional ConfigMap properties.

### Cons

* Adds complexity to the generated code, as we now have two properties for each optional secret, and we need to ensure
they are mutually exclusive.

## Generation Option 2: Single Property with Union Type

We could have a single property in the generated code that can accept either a plain text value or a secret reference,
using a union type to allow for both possibilities.

### Pros

* Simplifies the generated code by having a single property for each optional secret.
* Reduces the risk of user error by eliminating the need to ensure mutual exclusivity between two properties.

### Cons

* Less consistent with existing patterns for optional ConfigMap properties.
* Contrary to Kubernetes conventions, which explicitly recommend avoiding union/polymorphic properties.


## Decision

Configuration Option 4: Change to new Enum, with backward compatibility.

Generation Option 1: Parallel Properties.

## Status

Proposed.

## Consequences

TBC.

## Experience Report

TBC.

## References

None.