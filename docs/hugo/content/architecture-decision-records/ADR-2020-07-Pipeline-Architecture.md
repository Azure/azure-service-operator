---
title: '2020-07: Pipeline Architecture'
---

# Pipeline architecture

## Context

As the complexity of code generation grew, and with a partial view of the complexity yet to come, it became clear that we were facing a challenge. Each additional required function was making integration and testing progressively more difficult.

We also have the requirement to support a parallel pipeline that generates code appropriate for CrossPlane integration, and have a strong desire for an approach that minimizes the amount of ongoing maintenance.

Introducing a formal pipeline architecture, where each function forms a distinct pipeline stage, solves these problems.

* Each pipeline stage can (in theory, at least) be tested independently, improving our confidence in the functionality

* Dependencies between stages can be explicitly declared, allowing the validity of the pipeline to be checked

* A single factory method can create both required pipeline variants

## Decision

Adopted.

## Status

Pipeline introduced in [PR#171](https://github.com/Azure/k8s-infra/pull/171).

Stage prerequisites introduced in [PR#366](https://github.com/Azure/k8s-infra/pull/366) and postrequisites in [PR#1541](https://github.com/Azure/azure-service-operator/pull/1541).

## Consequences

Individual pipeline stages are independently tested, but some require considerable set up to create the expected initial state required for execution. The addition of mini-pipelines in [PR#1649](https://github.com/Azure/azure-service-operator/pull/1649) partially mitigates this.

## References

[Pipeline (software)](https://en.wikipedia.org/wiki/Pipeline_(software)) on Wikipedia
