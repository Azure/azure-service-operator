# Azure Service Operator (for Kubernetes)

[![Build Status](https://dev.azure.com/azure/azure-service-operator/_apis/build/status/Azure.azure-service-operator?branchName=master)](https://dev.azure.com/azure/azure-service-operator/_build/latest?definitionId=36&branchName=master)

> This project is experimental. Expect the API to change. It is not recommended for production environments.

## Introduction

Kubernetes offers the facility of extending it's API through the concept of 'Operators' ([Introducing Operators: Putting Operational Knowledge into Software](https://coreos.com/blog/introducing-operators.html)).

An Operator is an application-specific controller that extends the Kubernetes API to create, configure, and manage instances of complex stateful applications on behalf of a Kubernetes user. It builds upon the basic Kubernetes resource and controller concepts but includes domain or application-specific knowledge to automate common tasks.

This repository contains the resources and code to provision and deprovision different Azure services using a Kubernetes operator.

The Azure Operator comprises of:

- The Custom Resource Definitions (CRDs) for each of the Azure services that the Kubernetes user can provision
- The Kubernetes controller that watches for requests to create Custom Resources for these CRDs and creates them

The project was built using

[Kubebuilder](https://book.kubebuilder.io/)

## Azure Services supported

1. Resource Group
2. EventHub
3. [Azure SQL](/docs/azuresql/azuresql.md)

For information on how to build, test and run the operator, refer to the link below.
[Building, testing and running the operator](/docs/development.md)

## Testing

Testing the full project can be done in two ways:
* `make test` - Test against the Kubernetes integration testing framework.
* `make test-existing` - Test against an existing cluster. This is currently easiest to do run against a kind cluster setup.

Both of these invoke the Ginkgo test runner, running tests in parallel across 4 nodes by default.

### Focus tests

The ginkgo runner makes it simple to test single test cases, or packages in isolation.
* Rename the spec from `Describe` to `FDescribe`, or the individual test case from `It` to `FIt` excludes all other tests at the same level.
* Adding an additional parameter to ginkgo runner `--focus=REGEXP`.

See [https://onsi.github.io/ginkgo/#focused-specs]() for more details.


### Help

1. If the secret for the Eventhub in k8s gets deleted accidentally, the reconcile for the parent eventhub is triggered and secret gets created again.
2. If EventhubNamespace and Eventhub are deleted in Azure, then we need to delete the objects in k8s for the resources to be recreated. Reason being, if we apply the same manifest k8s does it recognise it as a change and the reconcile is not triggered. 

## Contributing

This project welcomes contributions and suggestions.  Most contributions require you to agree to a
Contributor License Agreement (CLA) declaring that you have the right to, and actually do, grant us
the rights to use your contribution. For details, visit https://cla.opensource.microsoft.com.

When you submit a pull request, a CLA bot will automatically determine whether you need to provide
a CLA and decorate the PR appropriately (e.g., status check, comment). Simply follow the instructions
provided by the bot. You will only need to do this once across all repos using our CLA.

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or
contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.
