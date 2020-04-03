# Azure Service Operator (for Kubernetes)

[![Build Status](https://dev.azure.com/azure/azure-service-operator/_apis/build/status/Azure.azure-service-operator?branchName=master)](https://dev.azure.com/azure/azure-service-operator/_build/latest?definitionId=36&branchName=master)

> This project is experimental. Expect the API to change. It is not recommended for production environments.

## Introduction

An Operator is an application-specific controller that extends the Kubernetes API to create, configure, and manage instances of complex stateful applications on behalf of a Kubernetes user. It builds upon the basic Kubernetes resource and controller concepts but includes domain or application-specific knowledge to automate common tasks.

(For more details about operators, we recommend [Introducing Operators: Putting Operational Knowledge into Software](https://coreos.com/blog/introducing-operators.html)).

This repository contains the resources and code to provision and deprovision different Azure services using a Kubernetes operator.

The Azure Operator comprises of:

- The Custom Resource Definitions (CRDs) for each of the Azure services that the Kubernetes user can provision
- The Kubernetes controller that watches for requests to create Custom Resources for these CRDs and creates them

The project was built using [Kubebuilder](https://book.kubebuilder.io/).

For more details on the control flow of the Azure Service operator, refer to the link below

[Azure Service Operator control flow](/docs/controlflow.md)

## Install the operator

This project maintains [releases of the Azure Service Operator](https://github.com/Azure/azure-service-operator/releases) that you can deploy via a [configurable Helm chart](./charts/azure-service-operator).

## Azure Services supported

1. [Resource Group](/docs/resourcegroup/resourcegroup.md)
2. [EventHub](/docs/eventhub/eventhub.md)
3. [Azure SQL](/docs/azuresql/azuresql.md)
4. [Azure Keyvault](/docs/keyvault/keyvault.md)
5. [Azure Rediscache](/docs/rediscache/rediscache.md)
6. [Storage Account](/docs/storage/storageaccount.md)
7. [Blob container](/docs/storage/blobcontainer.md)
8. [Azure Database for PostgreSQL](/docs/postgresql/postgresql.md)
9. [Virtual Network](/docs/virtualnetwork/virtualnetwork.md)
10.[Application Insights](/docs/appinsights/appinsights.md)
11.[API Management](/docs/apimgmt/apimgmt.md)
12.[Cosmos DB](/docs/cosmosdb/cosmosdb.md)

For more information on deploying, troubleshooting & deleting resources, refer to [this](/docs/customresource.md) link

## Building the operators

You can also build, test, and run the operator from source by [following these instructions](/docs/contents.md).

## Contributing

This project welcomes contributions and suggestions.  Most contributions require you to agree to a
Contributor License Agreement (CLA) declaring that you have the right to, and actually do, grant us
the rights to use your contribution. For details, visit https://cla.opensource.microsoft.com.

When you submit a pull request, a CLA bot will automatically determine whether you need to provide
a CLA and decorate the PR appropriately (e.g., status check, comment). Simply follow the instructions
provided by the bot. You will only need to do this once across all repos using our CLA.

For more specific information on the GIT workflow and guidelines to follow, check [here](docs/contributionguidelines.md).

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or
contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.
