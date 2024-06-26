---
title: Annotations understood by the operator
linktitle: Annotations
weight: 1 # This is the default weight if you just want to be ordered alphabetically
---

## Annotations specified by the user

Note that unless otherwise specified, allowed values are _case sensitive_ and should be provided in lower case.

### `serviceoperator.azure.com/reconcile-policy`

Specifies the reconcile policy to use. Allowed values are:

- `manage`: The operator performs all actions as normal. This is the default if no annotation is specified.
- `skip`: All modification actions on the backing Azure resource are skipped. GETs are still issued to ensure that the resource
exists. If the resource doesn't exist, the `Ready` condition will show a `Warning` state with the details of the missing resource
until the resource is created. If the resource is deleted in Kubernetes, it is _not_ deleted in Azure. In REST API terminology, 
PUT and DELETE are skipped while GET is allowed.
- `detach-on-delete`: Modifications are pushed to the backing Azure resource, but if the resource is deleted in Kubernetes, 
it is _not_ deleted in Azure. In REST API terminology, PUT and GET are allowed while DELETE is skipped.
    
Unknown values default to `manage`.

### `serviceoperator.azure.com/credential-from`

Instructs the operator to read the credential for the resource from the specified secret. 
This credential supersedes any global or namespace scoped credentials the operator has configured

Allow values are:
- The name of any secret in the same namespace as the resource. Secrets from other namespaces cannot be referenced.

See [authentication]( {{< relref "authentication#credential-scope" >}} ) for more details.

## Annotations written by the operator

These annotations are written by the operator for its own internal use. Their existence and usage may change in the future.
We recommend users avoid depending upon these annotations:

1. `serviceoperator.azure.com/resource-id`: The ARM resource ID.
2. `serviceoperator.azure.com/poller-resume-token`: JSON encoded token for polling long running operation.
3. `serviceoperator.azure.com/poller-resume-id`: ID describing the poller to use.
