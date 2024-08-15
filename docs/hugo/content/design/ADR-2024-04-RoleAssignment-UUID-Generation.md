---
title: '2024-04: RoleAssignments UUID Generation'
---

## Context

In ASO, we made several improvements to make RoleAssignment user-friendly by auto-generating their AzureName (which must be a UUID) using UUIDv5 with a seed string based on group, kind, namespace, and name. 
The aim for above approach was to:
- include the namespace and name to ensure no two RoleAssignments in the same cluster can end up with the same UUID.
- include the group and kind to ensure that different kinds of resources get different UUIDs. This isn't entirely required by Azure, but it makes sense to avoid collisions between two resources of different types even if they have the same namespace and name.
- include the owner group, kind, and name to avoid collisions between resources with the same name in different clusters that actually point to different Azure resources.

In scenarios where users have multiple ASO instances across different clusters, resources within the identical namespaces, and each cluster has resources with identical names pointing to distinct Azure resources, there is no native support for this configuration. To work around this limitation, users must manually assign a unique UUID to each RoleAssignment. Alternatively, setting the AzureName property to a random UUID by default can prevent the issue, but it may lead to complications if a RoleAssignment is ever deleted and then recreated (or migrated between clusters), as the old UUID would become orphaned.

Issue: [RoleAssignment UUID clashes](https://github.com/Azure/azure-service-operator/issues/3637)

### Option 1: Adding a new property

Add a new property on RoleAssignment spec to control this behaviour. 

To use random uuid: 
```yaml
uuidGeneration: random
```

To use deterministic uuid(default)
```yaml
uuidGeneration: stable
```

**Pro**: Backward compatible
**Pro**: Ease of use

**Con**: Need to implement infrastructure to add a new property
**Con**: Users will have to manage the AzureName by themselves while exporting/importing

### Option 2: Using annotations

We can use annotation below to control the name generation behaviour. 
In this case, If user does not specify any annotation and does not specify AzureName, ASO sets the default annotation below and follows the default behaviour of stable UUID generation.  

To use random uuid:
```yaml
serviceoperator.azure.com/uuid-generation: random 
```

To use deterministic uuid(default)
```yaml
serviceoperator.azure.com/uuid-generation: stable
```

**Pro**: Backward compatible
**Pro**: Ease of use

**Con**: Users will have to manage the AzureName by themselves while exporting/importing
**Con**: Pushes a crucial part of the resource definition into an annotation; the spec is no longer a complete definition of the resource.
**Con**: Annotations are far more easily modified by other tooling, which may have unexpected flow on effects.

### Option 3: Using subscription ID as seed

Adding subscriptionId as a seed string would make the resource names distinct across subscriptions

**Pro**: User does not have to do anything

**Con**: Moving resource with same name between older to newer ASO version would require a workaround
**Con**: Webhooks don't have information about subscriptionID

### Option 4: Use custom OperatorSpec property

Adding an OperatorSpec property to the RoleAssignment resource. Which would enable users to choose the generation type within operatorSpec while resource creation.

``` yaml
  operatorSpec:
    generationType: random/default
```

**Pro**: Backward compatible
**Pro**: Ease of use
**Pro**: Decoupled from annotations and resource spec

**Con**: Users will have to manage the AzureName by themselves while exporting/importing if they use Random generation

## Decision

Recommendation: Option 4 - Using OperatorSpec properties

We retain how the resource shape looks like and suggest using OperatorSpec property for the users running into the edge case. 

## Status

Proposed.

## Consequences

TBC

## Experience Report

TBC

## References

None