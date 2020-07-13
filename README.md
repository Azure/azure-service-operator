# Azure Service Operator (for Kubernetes)

[![Build Status](https://dev.azure.com/azure/azure-service-operator/_apis/build/status/Azure.azure-service-operator?branchName=master)](https://dev.azure.com/azure/azure-service-operator/_build/latest?definitionId=36&branchName=master)

> Note: The API is expected to change (while adhering to semantic versioning). Alpha and Beta resources are generally not recommended for production environments.

The **Azure Service Operator** helps you provision Azure resources and connect your applications to them from within Kubernetes.

## Overview

The Azure Service Operator comprises of:

- The Custom Resource Definitions (CRDs) for each of the Azure services a Kubernetes user can provision.
- The Kubernetes controller that watches for requests to create Custom Resources for each of these CRDs and creates them.

The project was built using [Kubebuilder](https://book.kubebuilder.io/).

Curious to see how it all works? Check out our [control flow diagrams](/docs/howto/controlflow.md).

## Supported Azure Services

- [Resource Group](/docs/services/resourcegroup/resourcegroup.md)
- [Event Hubs](/docs/services/eventhub/eventhub.md)
- [Azure SQL](/docs/services/azuresql/azuresql.md)
- [Azure Database for PostgreSQL](/docs/services/postgresql/postgresql.md)
- [Azure Database for MySQL](/docs/services/mysql/mysql.md)
- [Azure Key Vault](/docs/services/keyvault/keyvault.md)
- [Azure Cache for Redis](/docs/services/rediscache/rediscache.md)
- [Storage Account](/docs/services/storage/storageaccount.md)
- [Blob Storage](/docs/services/storage/blobcontainer.md)
- [Virtual Network](/docs/services/virtualnetwork/virtualnetwork.md)
- [Application Insights](/docs/services/appinsights/appinsights.md)
- [API Management](/docs/services/apimgmt/apimgmt.md)
- [Cosmos DB](/docs/services/cosmosdb/cosmosdb.md)
- [Virtual Machine](/docs/services/virtualmachine/virtualmachine.md)
- [Virtual Machine Scale Set](/docs/services/vmscaleset/vmscaleset.md)

## Quickstart

![Deploying ASO](/docs/images/asodeploy.gif)

Ready to quickly deploy the latest version of Azure Service Operator on your Kubernetes cluster and start exploring? Follow these steps.


1. Install [cert-manager](https://cert-manager.io/docs/installation/kubernetes/) on the cluster using the following command.

    ```sh
    kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v0.12.0/cert-manager.yaml
    ```
2.  Install [Helm](https://helm.sh/docs/intro/install/), and add the Helm repo for Azure Service Operator. Please note that the instructions here use Helm 3.

    ```sh
    helm repo add azureserviceoperator https://raw.githubusercontent.com/Azure/azure-service-operator/master/charts
    ```

3. Install the Azure Service Operator on your cluster using the following helm install command.

    Note that the [ServicePrincipal](https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-azure-cli) you pass to the command below needs to have access to create resources in your subscription. If you'd like to use Managed Identity for authorization instead, check out instructions [here](docs/howto/managedidentity.md)

    ```sh
    helm upgrade --install aso https://github.com/Azure/azure-service-operator/raw/master/charts/azure-service-operator-0.1.0.tgz \
            --create-namespace \
            --namespace=azureoperator-system \
            --set azureSubscriptionID=$AZURE_SUBSCRIPTION_ID \
            --set azureTenantID=$AZURE_TENANT_ID \
            --set azureClientID=$AZURE_CLIENT_ID \
            --set azureClientSecret=$AZURE_CLIENT_SECRET \
            --set image.repository="mcr.microsoft.com/k8s/azureserviceoperator:latest"
    ```

    You should now see the Azure service operator pods running in your cluster, like the below.

    ```console
    $ kubectl get pods -n azureoperator-system
    NAME                                                READY   STATUS    RESTARTS   AGE
    azureoperator-controller-manager-7dd75bbd97-mk4s9   2/2     Running   0          35s
    ```

To deploy an Azure service through the operator, check out the set of [supported Azure services](#supported-azure-services) and the sample YAML files in the `config/samples` [folder](../../config/samples) to create the resources using the following command.

```sh
kubectl apply -f <YAML file>
```

## About the project

This project maintains [releases of the Azure Service Operator](https://github.com/Azure/azure-service-operator/releases) that you can deploy via a [configurable Helm chart](docs/howto/helmdeploy.md).

Please see the [FAQ](docs/faq.md) for answers to commonly asked questions about the Azure Service Operator.

Have more questions? Feel free to consult our documentation [here](docs/howto/contents.md).

## Contributing

The [contribution guide](CONTRIBUTING.md) covers everything you need to know about how you can contribute to Azure Service Operators. The [developer guide](docs/howto/contents.md#developing-azure-service-operator) will help you onboard as a developer.

## Support

Azure Service Operator is an open source project that is **not** covered by the [Microsoft Azure support policy](https://support.microsoft.com/en-us/help/2941892/support-for-linux-and-open-source-technology-in-azure). Please search open issues [here](https://github.com/Azure/azure-service-operator/issues). If your issue isn't already represented, please [open a new one](https://github.com/Azure/azure-service-operator/issues/new/choose). The Azure Service Operator project maintainers will respond to the best of their abilities.

## Code of conduct

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/). For more information, see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq) or contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.
