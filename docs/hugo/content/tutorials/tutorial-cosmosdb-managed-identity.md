---
title: "Tutorial: CosmosDB to-do List with Managed Identity"
---

Follow the [guided example](https://github.com/Azure-Samples/azure-service-operator-samples/tree/master/cosmos-todo-list-mi)
to create a to-do list application backed by CosmosDB.
The CosmosDB is hosted in Azure but created easily via `kubectl` and Azure Service Operator!

It also uses Managed Identity + Workload Identity, so the identity of the application and its permissions can
be managed through YAML with Azure Service Operator.
