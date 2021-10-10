# Azure Service Operator (for Kubernetes)
[![Go Report Card](https://goreportcard.com/badge/github.com/Azure/azure-service-operator)](https://goreportcard.com/report/github.com/Azure/azure-service-operator)
[![Build Status](https://dev.azure.com/azure/azure-service-operator/_apis/build/status/Azure.azure-service-operator?branchName=master)](https://dev.azure.com/azure/azure-service-operator/_build/latest?definitionId=36&branchName=master)
![v2 Status](https://github.com/azure/azure-service-operator/actions/workflows/master-validation.yml/badge.svg?branch=master)

> Note: The API is expected to change (while adhering to semantic versioning). Alpha and Beta resources are generally not recommended for production environments. Alpha, Beta, and Stable mean roughly the same for this project as they do for [all Kubernetes features](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/#feature-stages).

## What is it?
**Azure Service Operator** (ASO) helps you provision Azure resources and connect your applications to them from within Kubernetes.

If you want to use Azure resources but would prefer to manage those resources using Kubernetes tooling and primitives (for example `kubectl apply`), then Azure Service Operator might be for you.

## Overview

The Azure Service Operator consists of:

- The Custom Resource Definitions (CRDs) for each of the Azure services a Kubernetes user can provision.
- The Kubernetes controller that manages the Azure resources represented by the user specified Custom Resources. The controller attempts to synchronize the desired state in the user specified Custom Resource with the actual state of that resource in Azure, creating it if it doesn't exist, updating it if it has been changed, or deleting it.

## Versions of Azure Service Operator
There are two major versions of Azure Service Operator: v1 and v2. Consult the below table and descriptions to learn more about which you should use.

> Note: ASO v1 and v2 are two totally independent operators. Each has its own unique set of CRDs and controllers. They can be deployed side by side in the same cluster.

| ASO Version | Lifecycle stage | Development status                | Installation options                                                                                                                                                           |
| ----------- | --------------- | --------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| v1          | Beta            | Bug and security fixes primarily. | [Helm chart](/charts), [OperatorHub](https://operatorhub.io/operator/azure-service-operator) or [GitHub release 1.x](https://github.com/Azure/azure-service-operator/releases) |
| v2          | Alpha           | Under active development.         | [GitHub release 2.x](https://github.com/Azure/azure-service-operator/releases)                                                                                                 |

### ASO v1
Azure Service Operator v1 is no longer under active development. Bug and security fixes are still made. Some features may be added if the scope is small and the impact is large, but we are winding down investment into ASO v1. If you are already using ASO v1 a migration path/tool will be provided to eventually move ASO v1 resources to ASO v2. In the meantime you can continue using ASO v1 as you have been.

[Learn more about Azure Service Operator v1](/docs/v1/README.md)

### ASO v2
Azure Service Operator v2 was built based on the lessons learned from ASO v1, with the following improvements:

* Supports code-generated CRDs based on [Azure OpenAPI specifications](https://github.com/Azure/azure-rest-api-specs). This enables us to quickly add new resources as they are requested.
* More powerful `Status`. You can view the actual state of the resource in Azure through ASO v2, which enables you to see server-side applied defaults and more easily debug issues.
* Dedicated storage versions. This enables faster (and less error prone) support for new Azure API versions, even if there were significant changes in resource shape.
* Uniformity. ASO v2 resources are very uniform due to their code-generated nature. 
* Clearer resource states. The state a resource is in is exposed via a [Ready condition](/docs/v2/design/resource-states.md). 

[Learn more about Azure Service Operator v2](/hack/README.md)

## Contributing

The [contribution guide](CONTRIBUTING.md) covers everything you need to know about how you can contribute to Azure Service Operators.

## Support

Please search open issues [here](https://github.com/Azure/azure-service-operator/issues). If your issue isn't already represented, please [open a new one](https://github.com/Azure/azure-service-operator/issues/new/choose). The Azure Service Operator project maintainers will respond to the best of their abilities.

For more information, see [SUPPORT.md](SUPPORT.md).
## Code of conduct

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/). For more information, see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq) or contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.
