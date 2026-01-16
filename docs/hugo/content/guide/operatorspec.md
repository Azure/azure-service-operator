---
title: Passing values directly to ASO through the operatorSpec
weight: 1 # This is the default weight if you just want to be ordered alphabetically
---

There may be instances where a user may want to configure where to place certain properties that will be used by the operator. Utilizing `OpeartorSpec` allows users to specify where to place specific properties. Fields in the `OperatorSpec` struct willbe interpreted directly by the operator instead of being passed to Azure. 

Unless the field is specifically configured to request the property, ASO will not retrieve it.

## How to configure OperatorSpec in ASO

You can consider fields under the `operatorSpec` field under certain resources with values that you wish to populate.

At the moment, the `operatorSpec` supports `configMapExpressions`, `secretExpressions`, and `secrets`.

- `configMapExpressions` allows uers to configuire where to place operator written dynamic ConfigMaps. You can learn more about `configMap` and walk through a sample [here](../configmaps/). 
- `secretExpressions` configures where to place operator written dynamic secrets, while `secrets` configure where to place Azure generated secrets. You can learn more about `secretExpressions`/`secrets` and walk through an example [here](../secrets/).
