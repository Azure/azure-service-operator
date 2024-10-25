---
title: "Tutorials"
linkTitle: "Tutorials"
weight: 20
menu:
  main:
    weight: 20
cascade:
- type: docs
- render: always
description: Tutorials for using Azure Service Operator
---

{{% cardpane %}}
{{% card header="CosmosDB to-do List"%}}

[Create a to-do list application](https://github.com/Azure-Samples/azure-service-operator-samples/tree/master/cosmos-todo-list) backed by CosmosDB.

The CosmosDB is hosted in Azure but created easily via `kubectl` and Azure Service Operator.

{{% /card %}}
{{% card header="CosmosDB to-do List with Managed Identity"%}}

[Create a to-do list application](https://github.com/Azure-Samples/azure-service-operator-samples/tree/master/cosmos-todo-list-mi) backed by CosmosDB using Managed Identity and Workload Identity.

{{% /card %}}
{{% /cardpane %}}

{{% cardpane %}}
{{% card header="PostgreSQL Votes"%}}

[Create a to-do list application](https://github.com/Azure-Samples/azure-service-operator-samples/tree/master/azure-votes-postgresql) backed by PostgreSQL.

The CosmosDB is hosted in Azure but created easily via `kubectl` and Azure Service Operator.

{{% /card %}}
{{% card header="Redis Votes"%}}

[Create a simple voting application](https://github.com/Azure-Samples/azure-service-operator-samples/tree/master/azure-votes-redis) backed by Azure Redis.

{{% /card %}}
{{% /cardpane %}}
