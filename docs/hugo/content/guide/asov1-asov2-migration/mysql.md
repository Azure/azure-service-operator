---
title: MySQL
---

ASOv1 supports MySQL single server, which is 
[being deprecated](https://learn.microsoft.com/azure/mysql/single-server/whats-happening-to-mysql-single-server).

ASOv2 does _not_ support MySQL single server, but it does support MySQL Flexible Server (the single server replacement).

Migration from ASOv1 to ASOv2 for these resources is a two-step process:
1. Migrate from Single Server to Flexible Server.
2. Import the Flexible Server into ASOv2.

### Migrate from Single Server to Flexible Server

Annotate your ASOv1 MySQL single servers in Kubernetes with the `skipreconcile=true` annotation, and then 
follow the [migration guide](https://learn.microsoft.com/azure/mysql/single-server/whats-happening-to-mysql-single-server#migrate-from-single-server-to-flexible-server).

Once the migration has been complete, use [asoctl](../../../tools/asoctl/) to import the new Flexible Server resource from 
Azure following the standard process defined in the [ASOv1 to ASOv2 migration guide](../)

## Secrets

ASOv1 `MySQLServer` saves a Kubernetes secret:

```
kubectl get secrets -n ns1

mysqlserver-mysqlserver-migration           Opaque   5      17h
```

This secret has the following 5 keys:

| Key                       | Source          | ASOv2 equivalent                                                         |
|---------------------------|-----------------|--------------------------------------------------------------------------|
| fullyQualifiedServerName  | User            | `.spec.operatorSpec.secrets.fullyQualifiedDomainName`                    |
| fullyQualifiedUsername    | User            | None                                                                     |
| mySqlServerName           | User            | None (see `.spec.operatorSpec.secrets.fullyQualifiedDomainName`)         |
| password                  | ASOv1 generated | `.spec.administratorLoginPassword`                                       |
| username                  | ASOv1 generated | `.spec.operatorSpec.configMaps.administratorLogin` (added in ASO v2.7.0) |
