# Azure Service Operator Documentation

## Beginner's guide to Azure Service Operator

- [Deploy with Helm (recommended)](helmdeploy.md)
    - Deploy Azure Service Operator using a configurable Helm chart.
- [Build locally & deploy without Helm](deploy.md)
    - Build Azure Service Operator locally and deploy to a Kubernetes cluster.
- [Provision a resource using Azure Service Operator](resourceprovision.md) Learn to use Azure Service Operator to provision different Azure resources.
- [Authentication using managed identities](managedidentity.md)
    - Use a managed identity for authentication of Azure Service Operator instead of a Service Principal.
- [Secret management](secrets.md) Learn how secrets are managed in Azure Service Operator and choose the best option for you.
- [How it works](controlflow.md) Learn the control flow behind deploying the operator and provision resources through the operator.

## Developing Azure Service Operator

- [Prerequisites](prereqs.md):
    - Find the list of tools and dependencies needed to develop and deploy the Azure Service Operator
- [Developing](development.md)
    - Find information for running the operator locally.
- [Testing](testing.md)
    - Find information about how to run and author tests.
- [Debugging](debugging.md)
    - Debug your Azure Service Operator deployment.
- [New Operator Guide](newoperatorguide.md)
    - Create a new operator to be used with Azure Service Operator.
- [Kubebuilder tips & tricks](kubebuilder.md) Install and get started with Kubebuilder
