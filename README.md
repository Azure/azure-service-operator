# Azure Service Operator (for Kubernetes)

[![Build Status](https://dev.azure.com/azure/azure-service-operator/_apis/build/status/Azure.azure-service-operator?branchName=master)](https://dev.azure.com/azure/azure-service-operator/_build/latest?definitionId=36&branchName=master)

> This project is experimental. The API is expected to change (while adhering to semantic versioning). It is not recommended for production environments.

The Azure Service Operator helps you provision Azure resources and connect your applications to them from within Kubernetes.

## Overview

The Azure Operator comprises of:

- The Custom Resource Definitions (CRDs) for each of the Azure services that the Kubernetes user can provision
- The Kubernetes controller that watches for requests to create Custom Resources for these CRDs and creates them

The project was built using [Kubebuilder](https://book.kubebuilder.io/).

For more details on the control flow of the Azure Service operator, refer to the link below

[Azure Service Operator control flow](/docs/design/controlflow.md)

## Azure Services supported

- [Resource Group](/docs/services/resourcegroup/resourcegroup.md)
- [EventHub](/docs/services/eventhub/eventhub.md)
- [Azure SQL](/docs/services/azuresql/azuresql.md)
- [Azure Database for PostgreSQL](/docs/services/postgresql/postgresql.md)
- [Azure Database for MySQL](/docs/services/mysql/mysql.md)
- [Azure Keyvault](/docs/services/keyvault/keyvault.md)
- [Azure Rediscache](/docs/services/rediscache/rediscache.md)
- [Storage Account](/docs/services/storage/storageaccount.md)
- [Blob container](/docs/services/storage/blobcontainer.md)
- [Virtual Network](/docs/services/virtualnetwork/virtualnetwork.md)
- [Application Insights](/docs/services/appinsights/appinsights.md)
- [API Management](/docs/services/apimgmt/apimgmt.md)
- [Cosmos DB](/docs/services/cosmosdb/cosmosdb.md)
- [Virtual Machine](/docs/services/virtualmachine/virtualmachine.md)
- [Virtual Machine Scale Set](/docs/services/vmscaleset/vmscaleset.md)

## Getting started

This project maintains [releases of the Azure Service Operator](https://github.com/Azure/azure-service-operator/releases) that you can deploy via a [configurable Helm chart](/charts/azure-service-operator).

For detailed instructions on getting started, go [here](docs/howto/contents.md).

Please see the [FAQ][docs/faq.md] for answers to commonly asked questions about the Azure Service Operator

## Contributing

The [contribution guide][contribution-guide] covers everything you need to know about how you can contribute to Azure Service Operators. The [developer guide][developer-guide] will help you onboard as a developer.

## Support

Azure Service Operator is an open source project that is [**not** covered by the Microsoft Azure support policy](https://support.microsoft.com/en-us/help/2941892/support-for-linux-and-open-source-technology-in-azure). [Please search open issues here](https://github.com/Azure/azure-service-operator/issues), and if your issue isn't already represented please [open a new one](https://github.com/Azure/azure-service-operator/issues/new/choose). The Azure Service Operator project maintainers will respond to the best of their abilities.

## Code of conduct

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/). For more information, see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq) or contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.

[contribution-guide]: CONTRIBUTING.md
[developer-guide]: docs/howto/contents.md
[FAQ]: docs/faq.md