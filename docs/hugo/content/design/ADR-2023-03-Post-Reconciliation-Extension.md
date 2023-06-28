---
title: '2023-03: Post Reconciliation Extensions'
---

## Context

The current behaviour of Azure Service Operator (ASO) is to set the `Ready` status on a resource by default after a successful creation of a resource. 

For some resources, this is not sufficient as there may be additional steps required after initial creation of the resource. We need the ability to check the status of a created resource post reconciliation and set a condition accordingly. 

An example of this is `PrivateEndpoints` with `ConnectionState` of `Approved`, `Rejected` or `Pending`. We don't want a `PrivateEndpoint` to go to condition `Ready` until the connection is approved.

We need a way to extend the operator to handle the above case, an extension point to be called post reconciliation allowing handwritten code to set a custom condition on the resource based on the status we've received from ARM.

## Decision

We'll define a new extension called `PostReconciliationChecker` in similar way to our `PreReconciliationChecker` extension point to handle post reconciliation operations on a resource. 

The extension point will receive the following parameters:

- The current resource, with a status freshly updated from Azure.
- A Kubernetes client allowing for Cluster operations.
- An ARM client allowing for ARM operations.
- A logger to allow for tracing of what the extension did.
- A context to allow cancellation of long-running operations.

The return will be one of three possibilities:

- `Success` as `PostReconcileCheckResult` if the post reconcile check is met. Nothing to do, the existing reconciler behaviour applies here.  
- `Failure` as `PostReconcileCheckResult` if the post reconcile check is not met. In this case, `NewReadyConditionImpactingError` would be set on the resource by the reconciler.
- `error` if something went wrong.

## Example

To allow post reconciliation check for PrivateEndpoints, we'll introduce a PostReconciliationChecker extension point.

We create the required extension interface in the package genruntime/extensions:

``` go
// postreconciliation_checker.go

package extensions

type PostReconciliationChecker interface {
	// PostReconcileCheck does a post-reconcile check to see if the resource is in a state to set 'Ready' condition.
	// ARM resources should implement this if they need to defer the Ready condition until later.
	// Returns PostReconcileCheckResultSuccess if the reconciliation is successful.
	// Returns PostReconcileCheckResultFailure and a human-readable reason if the reconciliation should put a condition on resource.
	// ctx is the current operation context.
	// obj is the resource about to be reconciled. The resource's State will be freshly updated.
	// kubeClient allows access to the cluster for any required queries.
	// armClient allows access to ARM for any required queries.
	// log is the logger for the current operation.
	// next is the next (nested) implementation to call.
	PostReconcileCheck(
		ctx context.Context,
		obj genruntime.MetaObject,
		owner genruntime.MetaObject,
		kubeClient kubeclient.Client,
		armClient *genericarmclient.GenericClient,
		log logr.Logger,
		next PostReconcileCheckFunc,
	) (PostReconcileCheckResult, error)
}

```

The type `PostReconcileCheckResult` contains the information of the result with a message and conditions from `PostReconcileCheck` method to indicate the reconciler on the action to be taken.

``` go
// postreconciliation_checker.go

package extensions

type PostReconcileCheckResult struct {
	action   postReconcileCheckResultType
	severity conditions.ConditionSeverity
	reason   conditions.Reason
	message  string
}

```

To implement the above, we need to write an extension manually on `PrivateEndpoint` type which implements the above `PostReconciliationChecker` interface:

``` go

// private_endpoint_extensions.go

package customization

var _ extensions.PostReconciliationChecker = &PrivateEndpointExtension{}

func (extension *PrivateEndpointExtension) PostReconcileCheck(
	_ context.Context,
	obj genruntime.MetaObject,
	_ genruntime.MetaObject,
	_ kubeclient.Client,
	_ *genericarmclient.GenericClient,
	_ logr.Logger
	_ PostReconcileCheckFunc,) (extensions.PostReconcileCheckResult, error) {

	if endpoint, ok := obj.(*network.PrivateEndpoint); ok && endpoint.Status.PrivateLinkServiceConnections != nil {

		for _, connection := range endpoint.Status.PrivateLinkServiceConnections {
			if *connection.PrivateLinkServiceConnectionState.Status != "Approved" {
				// Returns 'conditions.NewReadyConditionImpactingError' error
				return extensions.PostReconcileCheckResultFailure(
					fmt.Sprintf(
						"Private connection '%s' to the endpoint requires approval %q",
						*connection.Id,
						*connection.PrivateLinkServiceConnectionState)), nil
			}
		}
	}

	return extensions.PostReconcileCheckResultSuccess(), nil
}

```

## Status

TBC

## Consequences

TBC

## Experience Report

As described in [Reconciliation Extensions]( {{< relref "ADR-2022-12-Reconciliation-Extensions" >}} ), we've found that providing a `kubeclient.Client` to the extension is not ideal, and we're replacing it with a `resolver.Resolver` that allows for easy mapping between a `genruntime.Reference` and a GVK.

## References

* [ADR 2022-12: Reconciliation Extensions]( {{< relref "ADR-2022-12-Reconciliation-Extensions" >}} )
* [#3105 - Modify Pre/Post-Reconciliation Extensions to provide a Resolver](https://github.com/Azure/azure-service-operator/pull/3105)
