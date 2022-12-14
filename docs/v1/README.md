# Azure Service Operator (for Kubernetes) v1

> **⚠️ We strongly recommend new users consider [ASO v2]((https://azure.github.io/azure-service-operator/)) instead of ASO v1**
>
> Azure Service Operator v1 is no longer under active development.

The project was built using [Kubebuilder](https://book.kubebuilder.io/).

Curious to see how it all works? Check out our [control flow diagrams](/docs/v1/howto/controlflow.md).

## Supported Azure Services

- [Resource Group](/docs/v1/services/resourcegroup/resourcegroup.md)
- [Event Hubs](/docs/v1/services/eventhub/eventhub.md)
- [Azure SQL](/docs/v1/services/azuresql/azuresql.md)
- [Azure Database for PostgreSQL](/docs/v1/services/postgresql/postgresql.md)
- [Azure Database for MySQL](/docs/v1/services/mysql/mysql.md)
- [Azure Key Vault](/docs/v1/services/keyvault/keyvault.md)
- [Azure Cache for Redis](/docs/v1/services/rediscache/rediscache.md)
- [Storage Account](/docs/v1/services/storage/storageaccount.md)
- [Blob Storage](/docs/v1/services/storage/blobcontainer.md)
- [Virtual Network](/docs/v1/services/virtualnetwork/virtualnetwork.md)
- [Application Insights](/docs/v1/services/appinsights/appinsights.md)
- [API Management](/docs/v1/services/apimgmt/apimgmt.md)
- [Cosmos DB](/docs/v1/services/cosmosdb/cosmosdb.md)
- [Virtual Machine](/docs/v1/services/virtualmachine/virtualmachine.md)
- [Virtual Machine Scale Set](/docs/v1/services/vmscaleset/vmscaleset.md)

> **Deprecation notice**: Azure Database for MySQL - Single Server is on the retirement path and is [scheduled for retirement by September 16, 2024](https://learn.microsoft.com/en-us/azure/mysql/single-server/whats-happening-to-mysql-single-server).   
> Existing instances can be migrated to Azure Database for MySQL - Flexible Server using the Azure Database migration Service.   
> Azure Database for MySQL - Flexible Server is [fully supported in ASO v2](https://azure.github.io/azure-service-operator/reference/#dbformysql).

## Quickstart

![Deploying ASO](/docs/v1/images/asodeploy.gif)

Ready to quickly deploy the latest version of Azure Service Operator on your Kubernetes cluster and start exploring? Follow these steps.

0. Before starting, you must have a Kubernetes cluster (at least version 1.16) [created and running](https://kubernetes.io/docs/tutorials/kubernetes-basics/create-cluster/). Check your connection and version with:

   ```console
   $ kubectl version
   Client Version: version.Info{Major:"1", Minor:"19", GitVersion:"v1.19.2", GitCommit:"f5743093fd1c663cb0cbc89748f730662345d44d", GitTreeState:"clean", BuildDate:"2020-09-16T13:41:02Z", GoVersion:"go1.15", Compiler:"gc", Platform:"linux/amd64"}
   Server Version: version.Info{Major:"1", Minor:"18", GitVersion:"v1.18.2", GitCommit:"52c56ce7a8272c798dbc29846288d7cd9fbae032", GitTreeState:"clean", BuildDate:"2020-04-30T20:19:45Z", GoVersion:"go1.13.9", Compiler:"gc", Platform:"linux/amd64"}
    ```
   You'll also need to have the [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/?view=azure-cli-latest) installed (>= 2.13.0).

1. Install [cert-manager](https://cert-manager.io/docs/installation/kubernetes/) on the cluster using the following command.

    ```sh
    kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v0.12.0/cert-manager.yaml
    ```
   **Note:** if you are using a more recent version of cert-manager you will have to add `--set certManagerResourcesAPIVersion=cert-manager.io/v1` to the Helm command specified below.

2.  Install [Helm](https://helm.sh/docs/intro/install/), and add the Helm repo for Azure Service Operator. Please note that the instructions here use Helm 3.

    ```sh
    helm repo add aso https://raw.githubusercontent.com/Azure/azure-service-operator/main/charts
    ```
3. Create an Azure Service Principal. You'll need this to grant Azure Service Operator permissions to create resources in your subscription.
   For more information about other forms of authentication supported by ASO, see [the authentication section of the deployment documentation](./docs/v1/howto/deploy.md#Authentication).

   First, set the following environment variables to your Azure Tenant ID and Subscription ID with your values:
    ```yaml
    AZURE_TENANT_ID=<your-tenant-id-goes-here>
    AZURE_SUBSCRIPTION_ID=<your-subscription-id-goes-here>
    ```

   You can find these values by using the [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/?view=azure-cli-latest):
    ```sh
    az account show
    ```
   Next, we'll create a service principal with Contributor permissions for your subscription, so ASO can create resources in your subscription on your behalf. Note that the [ServicePrincipal](https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-azure-cli) you pass to the command below needs to have access to create resources in your subscription. If you'd like to use Managed Identity for authorization instead, check out instructions [here](/docs/v1/howto/managedidentity.md).

    ```sh
    az ad sp create-for-rbac -n "azure-service-operator" --role contributor \
        --scopes /subscriptions/$AZURE_SUBSCRIPTION_ID
    ```

   This should give you output like the following:
    ```sh
    "appId": "xxxxxxxxxx",
    "displayName": "azure-service-operator",
    "name": "http://azure-service-operator",
    "password": "xxxxxxxxxxx",
    "tenant": "xxxxxxxxxxxxx"
    ```

   Once you have created a service principal, set the following variables to your app ID and password values:
    ```sh 
    AZURE_CLIENT_ID=<your-client-id> # This is the appID from the service principal we created.
    AZURE_CLIENT_SECRET=<your-client-secret> # This is the password from the service principal we created.
    ```

4. Install the Azure Service Operator on your cluster using Helm.

    ```sh
    helm upgrade --install aso aso/azure-service-operator \
            --create-namespace \
            --namespace=azureoperator-system \
            --set azureSubscriptionID=$AZURE_SUBSCRIPTION_ID \
            --set azureTenantID=$AZURE_TENANT_ID \
            --set azureClientID=$AZURE_CLIENT_ID \
            --set azureClientSecret=$AZURE_CLIENT_SECRET
    ```

   If you would like to install an older version you can list the available versions:
    ```sh
    helm search repo aso --versions
    ```

   You should now see the Azure service operator pods running in your cluster, like the below.

    ```console
    $ kubectl get pods -n azureoperator-system
    NAME                                                READY   STATUS    RESTARTS   AGE
    azureoperator-controller-manager-7dd75bbd97-mk4s9   2/2     Running   0          35s
    ```

To deploy an Azure service through the operator, check out the set of [supported Azure services](#supported-azure-services) and the sample YAML files in the `config/samples` [folder](config/samples) to create the resources using the following command.

```sh
kubectl apply -f <YAML file>
```

## About the project

This project maintains [releases of the Azure Service Operator](https://github.com/Azure/azure-service-operator/releases) that you can deploy via a [configurable Helm chart](/docs/v1/howto/helmdeploy.md).

Please see the [FAQ](docs/faq.md) for answers to commonly asked questions about the Azure Service Operator.

Have more questions? Feel free to consult our documentation [here](/docs/v1/howto/contents.md).
