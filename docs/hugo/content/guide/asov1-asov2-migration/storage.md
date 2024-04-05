---
title: Storage Account
---
## Secrets

ASOv1 `StorageAccounts` save a Kubernetes secret:

```
kubectl get secrets -n ns1

storageaccount-cutoverteststorage1           Opaque   5      17h
```

This secret has the following 5 keys:

| Key                | Source | ASOv2 equivalent                              |
|--------------------|--------|-----------------------------------------------|
| StorageAccountName | User   | None                                          |
| connectionString0  | Azure  | None (see `.spec.operatorSpec.secrets.key1`)  |
| key0               | Azure  | `.spec.operatorSpec.secrets.key1`             |
| connectionString1  | Azure  | None (see `.spec.operatorSpec.secrets.key2`)  |
| key1               | Azure  | `.spec.operatorSpec.secrets.key2`             |

Instead of full connection strings, ASOv2 exposes individual endpoints such as `blobEndpoint`, which you can use to 
craft a connection string.

Example ASOv2 YAML snippet:
```yaml
spec:
  operatorSpec:
    secrets:
      blobEndpoint:
        name: storageaccount-cutoverteststorage1-asov2
        key: blobEndpoint
      key1:
        name: storageaccount-cutoverteststorage1-asov2
        key: key0  # Matches the name the ASOv1 generated secret used
```

Once you've applied the above, make sure to update your applications to depend on the new secret
written by ASOv2.

TODO: How to work around connection string issues?
1. init container? https://stackoverflow.com/questions/77748779/how-to-prepare-environment-variables-with-init-containers-in-a-kubernetes-deploy
2. Support https://github.com/Azure/azure-service-operator/issues/3446?
3. Modify the command of the container itself to run a script first?
