---
title: 2024-05 OperatorSpec Custom Properties
---

## Context

When generating the AzureName for a RoleAssignment, we currently use a deterministic algorithm to generate a GUID, allowing RoleAssignment resources to be moved between instances of Azure Service Operator (ASO) without changes. This approach doesn't handle every scenario, and we have a bug report (see [#3637](https://github.com/Azure/azure-service-operator/issues/3637)) from a user who needs different behaviour.

This spotlights a need we've previously identified informally, where we want to be able to provide additional resource configuration over and above the shape defined by the Azure REST API.

Last time this came up, we wanted to make our KeyVault resources more flexible, making them truly goal-seeking. We solved this by extending the definition of `createMode` to include two additional values (`createOrRecover` and `purgeThenCreate`).

Redefining an existing property isn't always going to give us what we need, and in particular isn't an option for RoleAssignments.

We're expecting the need for this kind of additional configuration to come up again, and a solution that caters for future needs as well as the current one would be ideal.

Confining any additional properties to the `operatorSpec` keeps them contained, helps to avoid conflicts with the Azure REST API, and makes it clear that these properties are specific to ASO and not a part of the Azure resource shape.

### Option 1: Per-property configuration

Extend our configuration to allow the need for a new operator spec property to be declared against a specific resource property.

``` yaml
  authorization:
    2020-08-01-preview:
      RoleAssignment:
        AzureName:
          $operatorSpecProperty:
              name: generationType
              type: string
```

* PRO: Makes it clear the relationship between the new operator property and the resource property it configures.
* CON: Only allows one operator spec property per resource property.
* CON: Not a great fit when there are multiple properties influenced by a single operatorSpec property.

### Option 2: Per-resource configuration

Extend our configuration to allow the need for a new operator spec property to be declared against a specific resource.

``` yaml
  authorization:
    2020-08-01-preview:
      RoleAssignment:
        $operatorSpecProperties:
          - name: generationType
            type: string
```

* PRO: Allows multiple operator spec properties to be declared against a single resource.
* PRO: Allows for operator spec properties that aren't specific to a single property or group of properties
* CON: Doesn't make it clear which resource properties are influenced by the operator spec property (but we can provide documentation to clarify this).

### Option 3: Add new properties via a Type Transformation

Modify our type transformation system to allow new properties to be added to the operator spec.

``` yaml
  - group: authorization
    name: RoleAssignmentOperatorSpec
    property: generationType
    add: true
    target:
      type: string
    because: "We need to be able to configure the generation type for RoleAssignments"
```

* PRO: Allows for multiple operator spec properties to be declared against a single resource.
* PRO: Allows for operator spec properties that aren't specific to a single property or group of properties
* CON: Doesn't make it clear which resource properties are influenced by the operator spec property.
* CON: Type transformations are applied before the spec/status type split, and before any operator spec is introduced, so we'd need to run these transformations at a different point in our pipeline
* CON: If the operator spec doesn't already exist (for other reasons), we'd need to create it in order to add properties to it.

## Decision

Recommendation: Option 2: Per-resource configuration

## Status

## Consequences

## Experience Report

## References
