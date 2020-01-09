# azure-service-operator

[Azure Service Operator](https://github.com/azure/azure-service-operator) enables users to provision and deprovision various Azure services using a Kubernetes Operator.

## Installing

### Pre-Install

This Helm chart contains certificates that depend on [cert-manager](https://cert-manager.io/docs/installation/kubernetes/) to be installed. Install cert-manager prior to installing chart:

```
kubectl create namespace cert-manager
kubectl label namespace cert-manager certmanager.k8s.io/disable-validation=true
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v0.9.0/cert-manager.yaml
```

### Helm

Install [Helm](https://helm.sh/docs/intro/install/) on your machine.

Add the helm repo:
```console
helm repo add azureserviceoperator https://azureserviceoperator.azurecr.io/helm/v1/repo
```

## Getting Started

Set the following variables to your Azure Tenant ID and Subscription ID:
```
azureTenantID: 00000000-0000-0000-0000-000000000000
azureSubscriptionID: 00000000-0000-0000-0000-000000000000
```

### Authentication

Choose one of the following authentication methods:

#### Service Principal

Set the following variables to your client ID and secret values:
```
azureClientID: 00000000-0000-0000-0000-000000000000
azureClientSecret: 00000000-0000-0000-0000-000000000000
```

#### Managed Identity

Set the following Helm Chart Value:
```
azureUseMI: "1"
```

Install the aad-pod-identity project:
```
kubectl apply -f https://raw.githubusercontent.com/Azure/aad-pod-identity/master/deploy/infra/deployment.yaml
```

1. Create an identity that will be able to manage resources

    ```shell
    # Create an identity to give to our operator-manager that will be used to authorize creation of resources in our
    # subscription. This could be restricted to a resource group by changing the scope on the "Contributor" role below
    az identity create -g <resourcegroup> -n aso-manager-identity -o json
    ```

    This command will have an output like the below

    ```shell
    {
    "clientId": "288a7d63-ab78-442e-89ee-2a353fb990ab",
    "clientSecretUrl": "https://control-westus.identity.azure.net/subscriptions/7060bca0-7a3c-44bd-b54c-4bb1e9facfac/resourcegroups/resourcegroup-operators/providers/Microsoft.ManagedIdentity/userAssignedIdentities/aso-manager-identity/credentials?tid=72f988bf-86f1-41af-91ab-2d7cd011db47&oid=ddcc0726-c3cd-49b2-9a4b-68a4a33bdc1d&aid=288a7d63-ab78-442e-89ee-2a353fb990ab",
    "id": "/subscriptions/7060bca0-7a3c-44bd-b54c-4bb1e9facfac/resourcegroups/resourcegroup-operators/providers/Microsoft.ManagedIdentity/userAssignedIdentities/aso-manager-identity",
    "location": "westus",
    "name": "aso-manager-identity",
    "principalId": "ddcc0726-c3cd-49b2-9a4b-68a4a33bdc1d",
    "resourceGroup": "resourcegroup-operators",
    "tags": {},
    "tenantId": "72f988bf-86f1-41af-91ab-2d7cd011db47",
    "type": "Microsoft.ManagedIdentity/userAssignedIdentities"
    }
    ```

    ```shell
    # Give the AKS SP control over managing the identity we just created
    az role assignment create --role "Managed Identity Operator" --assignee <AKS Service Principal ID> --scope <Managed Identity ID Path>
    ```

    In the above case, this will look like below:

    ```shell
    az role assignment create --role "Managed Identity Operator" --assignee <AKS Service Principal ID> --scope "/subscriptions/7060bca0-7a3c-44bd-b54c-4bb1e9facfac/resourcegroups/resourcegroup-operators/providers/Microsoft.ManagedIdentity/userAssignedIdentities/aso-manager-identity"
    ```

    ```shell
    # Give our aso-manager-identity authorization to provision resources in our subscription
    az role assignment create --role "Contributor" --assignee <Managed Identity Principal ID> --scope <Subscription ID Path>
    ```

    In the above case, this will look like below:

    ```shell
    az role assignment create --role "Contributor" --assignee "288a7d63-ab78-442e-89ee-2a353fb990ab"  --scope "/subscriptions/7060bca0-7a3c-44bd-b54c-4bb1e9facfac"
    ```

2. Create and apply the AzureIdentity and Binding manifests

    ```yaml
    apiVersion: "aadpodidentity.k8s.io/v1"
    kind: AzureIdentity
    metadata:
      name: <a-idname>
    spec:
      type: 0
      ResourceID: /subscriptions/<subid>/resourcegroups/<resourcegroup>/providers/Microsoft.ManagedIdentity/userAssignedIdentities/<name>
      ClientID: <clientId>
    ```

    ```yaml
    apiVersion: "aadpodidentity.k8s.io/v1"
    kind: AzureIdentityBinding
    metadata:
      name: aso-identity-binding
    spec:
      AzureIdentity: <a-idname>
      Selector: aso_manager_binding
    ```

### Keyvault

By default, secrets will be stored as Kubernetes secrets. If you wish to store them in KeyVault instead:

Create a Key Vault to use to store secrets

    ```shell
    az keyvault create --name "OperatorSecretKeyVault" --resource-group "resourceGroup-operators" --location "West US"
    ```

    Add appropriate Key Vault access policies to allow the service principal access to this Key Vault

    ```shell
    az keyvault set-policy --name "OperatorSecretKeyVault" --spn <AZURE_CLIENT_ID> --secret-permissions get list delete set
    ```

    If you use Managed Identity instead of Service Principal, use the Client ID of the Managed Identity instead in the above command.

    ```shell
    az keyvault set-policy --name "OperatorSecretKeyVault" --spn <MANAGEDIDENTITY_CLIENT_ID> --secret-permissions get list delete set
    ```

Then, set the following chart value to your KeyVault name:
```
azureOperatorKeyvault: OperatorSecretKeyVault
```

### Install Chart

Finally, install the chart with your added values.

```
helm install azureserviceoperator/azure-service-operator -f values.yaml
```

## Configuration

The following table lists the configurable parameters of the azure-service-operator chart and its dependency chart, aad-pod-identity.

| Parameter                  | Description              | Default              |
|:---------------------------|:-------------------------|:---------------------|
| `azureSubscriptionID`  | Azure Subscription ID | `` |
| `azureTenantID`  | Azure Tenant ID | `` |
| `azureClientID`  | Azure Service Principal Client ID | `` |
| `azureClientSecret`  | Azure Service Principal Client Secret | `` |
| `azureUseMI`  | Set to 1 if using Managed Identity for authentication | `` |
| `azureOperatorKeyvault`  | Set this value with the name of your Azure Key Vault resource if you prefer to store secrets in Key Vault rather than as Kubernetes secrets (default) | `` |
| `image.repository`  | Image repository | `mcr.microsoft.com/k8s/azure-service-operator:0.0.9150` |