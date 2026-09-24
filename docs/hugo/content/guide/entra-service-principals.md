---
title: Entra Service Principals
---

An Entra `ServicePrincipal` is the tenant-local instance of an application. Its
`appId` (application/client ID) is globally stable, whereas its `entraID`
(object ID) varies by tenant. Use `spec.appId` to adopt a service principal
without first discovering its tenant-local object ID:

```yaml
apiVersion: entra.azure.com/v1
kind: ServicePrincipal
metadata:
  name: cassandra-service-principal
spec:
  appId: a232010e-820c-4083-83bb-3ace5fc29d0b
  operatorSpec:
    creationMode: AdoptOnly
    configmaps:
      entraID:
        name: cassandra-service-principal
        key: objectId
```

`AdoptOnly` reports an error if the service principal is not present; it does
not create one. The default mode, `AdoptOrCreate`, creates a service principal
for the specified application ID if none exists. `AlwaysCreate` skips adoption.
The resolved object ID is available in `status.entraID` and, when configured,
in the specified ConfigMap for use by other resources.
Deleting a Kubernetes resource that adopted an existing service principal does
not delete the service principal in Entra.

ASO looks up a principal by `appId` first. If it exists, ASO adopts it
regardless of its display name. If no principal matches the GUID, ASO looks
up `spec.displayName` and adopts the sole match, as it does for Entra
Applications. Multiple name matches cause an error instead of creating
another principal. If both fields are specified and the name matches a
different `appId`, reconciliation also errors instead of adopting the wrong
principal or creating a duplicate.

An existing principal can be adopted by `displayName` alone, but Graph
requires `appId` to **create** a new one; a name-only resource with no match
reports an error. For example, a name-only adoption can use:

```yaml
apiVersion: entra.azure.com/v1
kind: ServicePrincipal
metadata:
  name: existing-service-principal
spec:
  displayName: Existing Service Principal
  operatorSpec:
    creationMode: AdoptOnly
```

`spec.displayName` can update a principal created by ASO,
but adopted principals are never modified, even when a display name is
specified.
