---
title: EventHub
---
## Secrets

ASOv1 `EventHubs` save a Kubernetes secret:

```
kubectl get secrets -n ns1

eventhub-eventhub-sample-1                   Opaque   7      22h
```

This secret has the following 7 keys:

| Key                       | Source | ASOv2 equivalent                                                                                                                                                                                                                                                                           |
|---------------------------|--------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| eventhubName              | User   | None                                                                                                                                                                                                                                                                                       |
| eventhubNamespace         | User   | None                                                                                                                                                                                                                                                                                       |
| primaryConnectionString   | Azure  | `.spec.operatorSpec.secrets.primaryConnectionString` (added in ASO v2.7.0)                                                                                                                                                                                                                 |
| primaryKey                | Azure  | `.spec.operatorSpec.secrets.primaryKey` (added in ASO v2.7.0)                                                                                                                                                                                                                              |
| secondaryConnectionString | Azure  | `.spec.operatorSpec.secrets.secondaryConnectionString` (added in ASO v2.7.0)                                                                                                                                                                                                               |
| secondaryKey              | Azure  | `.spec.operatorSpec.secrets.secondaryKey` (added in ASO v2.7.0)                                                                                                                                                                                                                            |
| sharedaccessKey           | Azure  | None, see ASOv2 [NamespacesAuthorizationRules](https://azure.github.io/azure-service-operator/reference/eventhub/v1api20211101/#eventhub.azure.com/v1api20211101.NamespacesAuthorizationRule) for exporting keys for any authorization rule except the default `RootManageSharedAccessKey` |
