---
title: Configuring OperatorSpec
weight: 1 # This is the default weight if you just want to be ordered alphabetically
---

The `operatorSpec` field lets you control how ASO handles a resource—independent of what gets sent to Azure. Use it to export values from your resources to `ConfigMaps` or `Secrets`, or to configure resource-specific operator behaviour.

Every ASO resource includes an `operatorSpec` field in its spec. ASO interprets these fields directly during reconciliation rather than passing them to Azure.

## Available properties

All resources support these properties by default:

- `configMapExpressions`: Export values from the resource's spec or status to a `ConfigMap`. See the [ConfigMaps documentation](../configmaps/) for details.
- `secretExpressions`: Export values to a `Secret`. See the [Secrets documentation](../secrets/) for details.

Many resources include the additional properties `configMaps` and `secrets`, based on configuration provided to our code generator. These are pre-defined export destinations for specific Azure-generated values.

- On a very few resources, there are additional properties. For example, [`RoleAssignment`](https://azure.github.io/azure-service-operator/reference/authorization/v1api20220401/#RoleAssignmentOperatorSpec) uses `namingConvention` to configure how the GUID for the assignment name is generated.

ASO only retrieves properties that you explicitly configure in `operatorSpec`. If you don't specify an export destination, the value won't be fetched.
