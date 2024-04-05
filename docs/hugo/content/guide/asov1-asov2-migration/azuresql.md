---
title: Azure SQL
---
## Secrets

ASOv1 `AzureSqlServer` creates a Kubernetes secret associated with the SQL server:

```
kubectl get secrets -n ns1

azuresqlserver-azuresql-migration-sample-1   Opaque   5      22h
```

This secret has the following 5 keys:

| Key                      | Source          | ASOv2 equivalent                                                    |
|--------------------------|-----------------|---------------------------------------------------------------------|
| azureSqlServerName       | User            | None (see `.spec.operatorSpec.configMaps.fullyQualifiedDomainName`) |
| fullyQualifiedServerName | Azure           | `.spec.operatorSpec.configMaps.fullyQualifiedDomainName`            |
| fullyQualifiedUsername   | ASOv1 generated | None                                                                |
| password                 | ASOv1 generated | `.spec.administratorPassword`                                       |
| username                 | ASOv1 generated | `.spec.administratorLogin`                                          |

Unlike ASOv1, ASOv2 does not automatically generate any usernames or passwords. It is up to you to manage and generate 
usernames and passwords. You can continue using the ASOv1 generated username and password when you switch from ASOv1 to ASOv2,
as shown below.

Example ASOv2 YAML snippet for exporting FQDN to configmap:
```yaml
spec:
  operatorSpec:
    configMaps:
      fullyQualifiedDomainName:
        name: azuresqlserver-azuresql-migration-sample-1-asov2
        key: fullyQualifiedServerName
```

Example ASOv2 YAML snippet for configuring the server username and password:

ASOv2 doesn't classify the administrator account name as a secret, so it will be automatically
included in the YAML after running `asoctl`. 

We recommend making a _new_ secret containing just the server password. You can produce this secret by copying 
the value from the ASOv1 secret `password` key into a new secret. In this example we'd create a new secret named 
`azuresqlserver-azuresql-migration-sample-admin-pw` containing the `password` key with value from the ASOv1 secret
`azuresqlserver-azuresql-migration-sample-1` `password` key.
```yaml
spec:
  administratorLogin: myusername  # This value will be automatically populated by asoctl
  administratorLoginPassword: 
    name: azuresqlserver-azuresql-migration-sample-admin-pw
    key: password
```

Once you've applied the above, make sure to update your applications to depend on the new configmap
written by ASOv2 (`azuresqlserver-azuresql-migration-sample-1-asov2`), and the new secret written by you 
(`azuresqlserver-azuresql-migration-sample-admin-pw`).
