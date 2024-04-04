---
title: PostgreSQL
---

ASOv1 supports PostgreSQL single server, which is 
[being deprecated](https://azure.microsoft.com/updates/azure-database-for-postgresql-single-server-will-be-retired-migrate-to-flexible-server-by-28-march-2025/).

ASOv2 does _not_ support PostgreSQL single server, but it does support PostgreSQL Flexible Server 
(the single server replacement).

Migration from ASOv1 to ASOv2 for these resources is a two-step process:
1. Migrate from Single Server to Flexible Server.
2. Import the Flexible Server into ASOv2.

### Migrate from Single Server to Flexible Server

Annotate your ASOv1 PostgreSQL single servers in Kubernetes with the `skipreconcile=true` annotation, 
and then follow the 
[migration guide](https://learn.microsoft.com/azure/postgresql/migrate/migration-service/concepts-migration-service-postgresql).

Once the migration has been complete, use [asoctl](../../../tools/asoctl/) to import the new Flexible Server resource from 
Azure following the standard process defined in the [ASOv1 to ASOv2 migration guide](../)

