---
title: Passing values directly to ASO through the operatorSpec
weight: 1 # This is the default weight if you just want to be ordered alphabetically
---

There may be instances where a user may want to configure where to place certain properties that will be used by the operator. Utilizing `OperatorSpec` allows users to specify where to place specific properties. Fields in the `OperatorSpec` struct will be interpreted directly by the operator instead of being passed to Azure. 

Each resource has an `OperatorSpec` that allows a user to configure additional behavior which ASO will apply as it reconciles the resource. Values from the spec or status of the resource can be exported to `configmaps` or `secrets` by populating the `configMapExpressions` and/or `secretExpressions` accordingly. 

Unless the field is specifically configured to request the property, ASO will not retrieve it.

## How to configure OperatorSpec in ASO

You can consider fields under the `operatorSpec` field under certain resources with values that you wish to populate.

At the moment, some common properties the `operatorSpec` supports include `configMapExpressions`, `secretExpressions`, `secrets`, and `configMaps`. Every resource has `configMapExpressions` and `secretExpressions` by default, while `configMaps` and `secrets` will only appear if they are configured to do so via the generator. Some resources, such as the [`RoleAssignmentOperatorSpec`](https://azure.github.io/azure-service-operator/reference/authorization/v1api20220401/#RoleAssignmentOperatorSpec), has the `namingConvention` field that is used to configure how the GUID used for the assignment name is generated. 

- `configMapExpressions` allows uers to configuire where to place operator written dynamic ConfigMaps. You can learn more about `configMap` and walk through a sample [here](../configmaps/). 
- `secretExpressions` configures where to place operator written dynamic secrets, while `secrets` configure where to place Azure generated secrets. You can learn more about `secretExpressions`/`secrets` and walk through an example [here](../secrets/).
