---
title: '2024-04: RoleAssignments UUID Generation'
---

## Context

In ASO, we have made several improvements on making the RoleAssignment's name user-friendly by auto-generating the UUID using UUIDv5 with a seed string based on the group, kind, namespace and name. 
The aim for above approach was to:
- include the namespace and name to ensure no two RoleAssignments in the same cluster can end up with the same UUID.
- include the group and kind to ensure that different kinds of resources get different UUIDs. This isn't entirely required by Azure, but it makes sense to avoid collisions between two resources of different types even if they have the same namespace and name.
- include the owner group, kind, and name to avoid collisions between resources with the same name in different clusters that actually point to different Azure resources.

However, the case where users have multiple ASO instances in multiple clusters with resources in the same namespace, in each cluster having the same name and pointing to different Azure resource is not supported without the user manually giving each RoleAssignment its own UUID.
The above issue can be avoided by defaulting the AzureName property to a random UUID, however it will cause issues if the RoleAssignment is ever deleted and recreated (or migrated between clusters) as the old UUID will be orphaned.

Issue: [RoleAssignment UUID clashes](https://github.com/Azure/azure-service-operator/issues/3637)

### Option 1: Adding a new property

Add a new property on RoleAssignment spec to control this behaviour. 

To use random uuid: 
```yaml
uuidGeneration: Random
```

To use deterministic uuid(default)
```yaml
uuidGeneration: Default
```

**Pro**: Backward compatible
**Pro**: Ease of use

**Con**: Need to implement infrastructure to add a new property
**Con**: Users will have to manage the AzureName by themselves while exporting/importing

### Option 2: Using annotations

We can use annotation below to control the name generation behaviour

To use random uuid:
```yaml
serviceoperator.azure.com/uuid-generation: random 
```

To use deterministic uuid(default)
```yaml
serviceoperator.azure.com/uuid-generation: Default
```

**Pro**: Backward compatible
**Pro**: Ease of use

**Con**: Users will have to manage the AzureName by themselves while exporting/importing

### Option 3: Using subscription ID as seed

Adding subscriptionId as a seed string would make the resource names distinct across subscriptions

**Pro**: User does not have to do anything

**Con**: Moving resource with same name between multiple ASO instances would require a workaround
**Con**: Webhooks don't have information about subscriptionID

## Status

Proposed.

## Consequences

TBC

## Experience Report

TBC

## References

None