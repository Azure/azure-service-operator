---
title: '2025-01: Improving ARM References'
toc_hide: true
---

## Context

Many Azure Resource Manager (ARM) resources have references to other resources. When working with ARM Templates or Bison, these are almost universally handled as simple strings containing a fully qualified ARM reference.

As a part of importing resources into Azure Service Operator (ASO), we modify these properties into formal resource references allowing them to contain either an ARM resource Id, or the GVK of a resource within the cluster.

We modify a referencing property as follows:

* The suffix `Reference` is added to the name (selected suffixes such as `Id` are also removed).
* The type of the property is changed from **string** to `genruntime.ResourceReference`

For example, `keyVaultId: string` becomes `keyVaultReference: genruntime.ResourceReference`.

Not every such property is unambiguously annotated in the OpenAPI specifications we read, so our code generator applies a number of heuristics to identify these properties. We end up with three groups of properties:

* Those that are definitely ARM references.
* Those that our heuristics suggest are ARM references.
* All other properties.

For those properties in the second group, we require manual configuration to indicate whether they are ARM references or not. This is done by adding an `$armReference` modifier into our configuration file. The code generator will fail with an error any property identified by our heuristics is not explicitly configured.

We can also use the configuration for any property in the third group that's actually an ARM reference too. This is useful when we have a property that's not identified by our heuristics, but we know it's an ARM reference.

For example, in this configuration fragment we can see confirmation that `ActionGroupsInformation.GroupIds` _**is**_ an ARM Reference, while `Detector.Id` _**is not**_:

``` yaml
  alertsmanagement:
    2021-04-01:
      SmartDetectorAlertRule:
        $export: true
        $supportedFrom: v2.11.0
      AlertRuleProperties:
        Scope:
          $armReference: true
      ActionGroupsInformation:
        GroupIds:
          $armReference: true
      Detector:
        Id:
          $armReference: false
```

### Missed ARM references

The major failure mode of this approach is that we can sometimes (fortunately rarely) miss identifying a property as an ARM reference. When this happens, we end up with a property that is a **string** when it should be a `genruntime.ResourceReference`. This isn't a problem if we catch the problem before releasing the resource, but if the resource is already released, changing the type and name of the property is a breaking change.

In [#3772](https://github.com/Azure/azure-service-operator/issues/3772) (_Bug: MySQL Flexible Server property sourceServerResourceId has wrong type_), @theunrepentantgeek made this comment:

> Trying out some improvements to our heuristic and finding other properties we missed:
>
> | Group | Version | Kind | Property |
> | -- | -- | -- | -- |
> | authorization | v1api20200801preview | RoleAssignmentProperties | DelegatedManagedIdentityResourceId |
> | authorization | v1api20220401 | RoleAssignmentProperties | DelegatedManagedIdentityResourceId |
> | dbformysql | v1api20210501 | ServerProperties | SourceServerResourceId |
> | machinelearningservices | v1api20210701 | IdentityForCmk | UserAssignedIdentity |
> | machinelearningservices | v1api20210701 | KeyVaultProperties | KeyVaultArmId |

For some of these, we were able to _patch_ the problem by importing a newer version of the resource and adding conversion code to translate between the old and new versions. Users of the newer version of the resource would then see the correct property types, while users of the old version are stuck with the incorrect types.

However, we cannot apply this workaround when there is no newer version of the resource available - as has occurred with the `authorization` group, where the latest version is v1api20220401.

Other examples

* `MetricAlert.ActionGroupId` - [#3643](https://github.com/Azure/azure-service-operator/issues/3643) _Improvement: support action group and metric linking_
* `ServerProperties.SourceServeResourceId` - see [#3829](https://github.com/Azure/azure-service-operator/issues/3829) _Update dbformysql to latest version_

### Aliases

There are rare properties that accept more than just ARM IDs - meaning that a direct transformation to a `ResourceReference` isn't sufficient.

One case of this is discussed in [#4531](https://github.com/Azure/azure-service-operator/issues/4531) (_ResourceReference should support Alias_) - in the case of `PrivateLinkServiceConnection`, the value `privateLinkServiceId` that's used to identify a `PrivateLinkService` may be _either_ a regular ARM ID or an alias.

We currently have no way to represent this in ASO, as validation on the `armId` property requires values to confirm to the ARM ID format. (Loosening validation to allow aliases would delay discovery of invalid values for the majority of cases where aliases are not permitted.)

### Friendly names

In [#3642](https://github.com/Azure/azure-service-operator/issues/3642) (_Improvement: support for the friendly names of the builtin roles_) we have a request to allow friendly names like `Contributor` instead of requiring users to look up the specific GUID needed and constructing the ARM ID by hand.

This is similar to _Alias support_ above, except that we would be resolving the name to an ARM ID within ASO itself, instead of leaving that to ARM.

Again, we currently have no way to represent this in ASO.

### Option 1: Break existing users

When we identify a property that should be treated as an ARM reference, but for which we missed setting `$armReference: true` in our configuration, update the configuration and run the code generator to _apply the change even if that resource has already been released_.

* PRO: Simple to do
* PRO: Familiar, as we already do this with unreleased resources
* PRO: Solves the _Missed ARM References_ problem
* CON: Breaking change for any existing users, as both the name and type of the property will change.
* CON: Does not solve the _Alias support_ or _Friendly names_ problems.

The breaking change is nasty, as an upgrade of the ASO operator would result in a cluster resource that's silently updated incorrectly, plus any attempt to repair the resource would fail until the YAML was modified to the new structure.

### Option 2: Allow preserving the original property

Under normal operation (`$armReference: true`) the existing property is removed, and replaced with a new property for the reference.

For example, we transform

``` go
type PrometheusRuleGroupAction struct {
    ActionGroupId    string
    ActionProperties map[string]string
}
```

to

``` go
type PrometheusRuleGroupAction struct {
    ActionGroupReference *genruntime.ResourceReference
    ActionProperties     map[string]string
}
```

When configured with `$armReference: false`, we leave the property as is.

We could add a third option to specify the existing property is to be retained, making the result of the transformation look like this:

``` go
type PrometheusRuleGroupAction struct {
    ActionGroupId        string
    ActionGroupReference *genruntime.ResourceReference
    ActionProperties     map[string]string
}
```

We already have precendent for this, with the `*FromConfig` properties that we inject for selected properties.

Validation would need to ensure the two properties are mutually exclusive.

When constructing the payload to send to ARM, we would need to select between the two properties, using whichever is set. We have an exisiting resource extension point we can use to implement this.

Adding a third option to a boolean flag is a very bad idea (though we've seen it done), so we'd need to rename `$armReference` to something else for clarity. Fortunately, we control `azure_arm.yaml` so this is achievable. (We don't believe any users have their own configuration file, but we should call this out as a breaking in our release notes anyway.)

One possible name for flag would be `$isReference` with values `arm`, `side-by-side` and `no`, though this phrasing is quite awkward.

* PRO: Relatively straightforward
* PRO: We already have precedent for side-by-side properties
* PRO: Solves all three scenarios.
* CON: Problematic to have _two_ properties that handle _three_ use-cases (e.g. use `privateLinkServiceReference` for GVK or ARMId, and `privateLinkServiceId` for alias).
* CON: No control over the naming of things

### Option 3: Custom reference types

Instead of always replacing the property with a `genruntime.ResourceReference`, allow selection between a set of standard reference types.

For the vast majority of reference properties, we'd continue to use the existing `genruntime.ResourceReference`.

To support the _Alias_ and _Friendly names_ cases, introduce a new reference type, say `genruntime.AliasResourceReference`, with the following structure:

``` go
type AliasResourceReference struct {
   Group     string
   Kind      string
   Namespace string
   Name      string

    ARMID string

    Alias string
}
```

The code generator would need to be updated to support this new value.

Finally, when the property isn't an reference at all, but just a simple string, we'd have a value that means _leave this alone_. This is the case where we'd currently us `$armReference: false`.

Again, our current `$armReference` flag would need to be renamed to something more general. Perhaps `$referenceType` with possible values `arm`, `armOrAlias`, and `other`.

* PRO: Great YAML structure for end users of resources.
* PRO: Extensible for the (admittedly unlikely) scenario where we need another different kind.
* PRO: Cleanly handles both the _Alias_ and _Friendly names_ scenarios.
* CON: Slightly greater implementation complexity.
* CON: Does not solve the _Missed ARM References_ problem.

### Option 4: Multivalued custom reference types

Similar to Option 3, but allowing set combinations of values to be specified for the reference type.

Possible values would be `arm`, `alias` and `other`, plus the combinations `arm+alias` and `arm+other`. Other combinations would be invalid.

``` yaml
MetricAlertAction:
    ActionGroupId:
        $referenceType: arm+other
```

Returning to the `PrometheusRuleGroupAction` example from above, we'd specify `$referenceType: arm+other` to get:

``` go
type PrometheusRuleGroupAction struct {
    ActionGroupId        string
    ActionGroupReference *genruntime.ResourceReference
    ActionProperties     map[string]string
}
```

This handles the _Missed ARM References_ problem by allowing us to retain the existing property while introducing the corrected property side-by-side.

In the `RoleAssignment` case, we'd specify `$referenceType: arm+alias` to get:

``` go
type RoleAssignment_Spec struct {
    // ... elided ...
    roleDefinitionReference *genruntime.AliasResourceReference
}
```

* PRO: Great YAML structure for end users.
* PRO: Good experience for configuration of our code generator.
* PRO: Extensible for the (admittedly unlikely) scenario where we need another different kind.
* PRO: Solves all three scenarios.
* CON: Slightly greater implementation complexity.


## Decision

Reccommendation: Option 4

Variation: We may want to YAML list syntax for clarity:

``` yaml
MetricAlertAction:
    ActionGroupId:
        $referenceType: 
          - arm
          - other
```

or 

``` yaml
MetricAlertAction:
    ActionGroupId:
        $referenceType: [arm, other]
```

The major downside of this is that we'd make the usual case (with just a single option) more verbose, but we could mitigate that by having two mutually exclusive options, `$referenceType` (a string) and `$referenceTypes` (a list).

* We'll want to look into where we resolve IDs and make sure that implementing this as a per-resource extension is right.
* The value `other` is a bit vague, we should try to find a better name that covers both _not a reference_ and _a reference we don't provide special support for_. Suggestions so far include `raw` and `unmodified`.

## Status

Proposed.

## Consequences

TBC

## Experience Report

TBC

## References

None
