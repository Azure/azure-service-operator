---
title: '2023-03: Post Reconciliation Extensions'
---

## Context

The current behaviour of Azure Service Operator (ASO) is to set the `Ready` status on a resource by default after a successful creation of a resource. 

For some resources, this is not sufficient as there may be additional steps required after initial creation of the resource. We need the ability to check the status of a created resource post reconciliation and set a condition accordingly. 

An example of this is `PrivateEndpoints` with `ConnectionState` of `Approved`, `Rejected` or `Pending`.

We need a way to extend the operator to handle the above case, an extension point to be called post reconciliation allowing handwritten code to set a custom condition on the resource based on the status we've received from ARM.

## Decision

We'll define a new extension called `PostReconciliationChecker` in similar way to our `PreReconciliationChecker` extension point. 

The extension point will receive the following parameters:

- The current resource, with a status freshly updated from Azure.
- A Kubernetes client allowing for Cluster operations.
- An ARM client allowing for ARM operations.
- A logger to allow for tracing of what the extension did.
- A context to allow cancellation of long-running operations.

The return will be one of three possibilities:

- `Success` as `PostReconcileResult` if the post reconcile check is met. Nothing to do, the existing reconciler behaviour applies here.  
- `Failure` as `PostReconcileResult` if the post reconcile check is not met. In this case, `NewReadyConditionImpactingError` would be set on the resource by the reconciler.
- `error` if something went wrong.

## Status

TBC

## Consequences

TBC

## Experience Report

TBC

## References