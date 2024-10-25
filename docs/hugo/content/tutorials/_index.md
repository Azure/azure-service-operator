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

Follow the [guided example](https://github.com/Azure-Samples/azure-service-operator-samples/tree/master/cosmos-todo-list)
to create a to-do list application backed by CosmosDB.

The CosmosDB is hosted in Azure but created easily via `kubectl` and Azure Service Operator.

{{% /card %}}
{{% card header="CosmosDB to-do List with Managed Identity"%}}

Follow the [guided example](https://github.com/Azure-Samples/azure-service-operator-samples/tree/master/cosmos-todo-list-mi)
to create a to-do list application backed by CosmosDB using Managed Identity and Workload Identity.

{{% /card %}}
{{% /cardpane %}}

{{% cardpane %}}
{{% card header="PostgreSQL Votes"%}}

Follow the [guided example](https://github.com/Azure-Samples/azure-service-operator-samples/tree/master/cosmos-todo-list)
to create a to-do list application backed by CosmosDB.

The CosmosDB is hosted in Azure but created easily via `kubectl` and Azure Service Operator.

{{% /card %}}
{{% card header="Redis Votes"%}}

Follow the [guided example](https://github.com/Azure-Samples/azure-service-operator-samples/tree/master/azure-votes-redis)
to create a simple voting application backed by Azure Redis.

{{% /card %}}
{{% /cardpane %}}
