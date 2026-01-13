---
title: Passing values directly to ASO through the operatorSpec
weight: 1 # This is the default weight if you just want to be ordered alphabetically
---

There may be instances where a user may want to configure where to place certain properties that will be used by the operator. Utilizing `OpeartorSpec` allows users to specify where to place specific properties. Fields in the `OperatorSpec` struct willbe interpreted directly by the operator instead of being passed to Azure. 

Unless the field is specifically configured to request the property, ASO will not retrieve it.

## How to configure OperatorSpec in ASO

