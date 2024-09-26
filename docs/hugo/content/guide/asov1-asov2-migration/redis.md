---
title: Redis
---
## Secrets

ASOv1 `RedisCaches` save a Kubernetes secret:

```
kubectl get secrets -n ns1

rediscache-rediscache-sample-1               Opaque   2      9s
```

This secret has the following 2 keys:

| Key               | Source | ASOv2 equivalent                          |
|-------------------|--------|-------------------------------------------|
| primaryKey        | Azure  | `.spec.operatorSpec.secrets.primaryKey`   |
| secondaryKey      | Azure  | `.spec.operatorSpec.secrets.secondaryKey` |

Example ASOv2 YAML snippet:
```yaml
spec:
  operatorSpec:
    secrets:
      primaryKey:
        name: rediscache-rediscache-sample-1-asov2
        key: primaryKey
      secondaryKey:
        name: rediscache-rediscache-sample-1-asov2
        key: secondaryKey
```

Once you've applied the above, make sure to update your applications to depend on the new secret
written by ASOv2.
