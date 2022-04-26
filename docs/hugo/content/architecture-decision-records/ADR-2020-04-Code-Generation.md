---
title: '2020-04: Why Code Generation?'
---

# Why Code Generation?

## Context

The Azure Service Operator CSE team had successfully developed support for a handful of ARM (Azure Resource Manager) based custom resources over the course of eighteen months or so.

Their experience yielded several useful lessons.

* Working closely with their customers, they only built out resource support for those features required by the current customer base. While this successfully reduced the amount of work, each additional customer brought new requirements and a need to revisit/enhance existing resources.

* Adding support for a brand-new resource was often needed by new adopters of ASO and required a significant effort. Extending ASO to have broad support across all (or nearly all) Azure ARM resources would require a major investment

* A separate controller for each resource type introduced inconsistencies in behavior

Multiple people concurrently had the idea of using code generation based on ARM JSON Schema to address these problems:

* By default, a code generator would mirror the the entire ARM specification for a resource in Kubernetes, providing full feature coverage with consistent naming.

* The overhead of adding a new resource would be much reduced, potentially limited to as little as testing
  
* Using a single generic controller for all resources would ensure consistent behavior

## Decision

A proof of concept by [David Justice](https://github.com/devigned) proved the concept sufficiently that we had confidence the idea would work.

## Status

Adopted, based on the proof of concept code in the [Azure/k8s-infra](https://github.com/Azure/k8s-infra/tree/master/hack/generator) repository.

## Consequences

We think it's working.

## Experience Report

The ARM JSON Schema definitions weren’t sufficient in isolation as they don’t contain all of the information required. Use of Azure Swagger specifications to augment our status types was introduced in [PR #205]( https://github.com/Azure/k8s-infra/pull/205).

Migration of the code generator into the ASO repo occurred in [PR #1427](https://github.com/Azure/azure-service-operator/pull/1427).

## References

None.
