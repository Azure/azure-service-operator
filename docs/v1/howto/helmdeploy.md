## Deploy Azure Service Operator using Helm charts

### Installing

#### Before you begin

This Helm chart contains certificates that depend on [cert-manager](https://cert-manager.io/docs/installation/kubernetes/) to be installed. Install cert-manager:

```sh
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v0.12.0/cert-manager.yaml
```

You can use the below command to check if the cert manager pods are ready. The cert manager pods should be running before proceeding to the next step.

```console
kubectl rollout status -n cert-manager deploy/cert-manager-webhook
```

Next, install the latest (3.x+) [Helm](https://helm.sh/docs/intro/install/) on your machine.

Add the helm repo:
```sh
helm repo add azureserviceoperator https://raw.githubusercontent.com/Azure/azure-service-operator/main/charts
```

### Configure Azure Service Operator

Copy the `values.yaml` [file](../../charts/azure-service-operator/values.yaml) and fill in the requisite values in the following steps.

First, set the following variables to your Azure Tenant ID and Subscription ID:
```yaml
azureTenantID: 00000000-0000-0000-0000-000000000000
azureSubscriptionID: 00000000-0000-0000-0000-000000000000
```

You can find these values by running the following:
```sh
az account show
```

#### Authentication

Next, choose one of the following authentication methods, and set its appropriate variables.

##### Service Principal

Once you have created a service principal, set the following variables to your client ID and secret values:
```yaml
azureClientID: 00000000-0000-0000-0000-000000000000
azureClientSecret: 00000000-0000-0000-0000-000000000000
```

##### Managed Identity

Follow the instructions [here](./managedidentity.md) to create an identity and assign it the correct permissions.

Set the following Helm Chart values:
```yaml
azureUseMI: True
installAadPodIdentity: True
azureClientID: 00000000-0000-0000-0000-000000000000
aad-pod-identity:
    azureIdentities:
        azureIdentity:
            resourceID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/<rg>/providers/Microsoft.ManagedIdentity/userAssignedIdentities/<identity>"
            clientID: "00000000-0000-0000-0000-000000000000"
```

#### Store secrets in Azure KeyVault (optional)

By default, secrets will be stored as Kubernetes secrets. If you wish to store them in KeyVault instead, follow the instructions [here](./deploy.md).

Then, set the following chart value to your KeyVault name:
```
azureOperatorKeyvault: OperatorSecretKeyVault
```

#### Install Chart

Finally, install the chart with your added values. The chart can be installed by using a values file or environment variables.
```
helm upgrade --install aso azureserviceoperator/azure-service-operator -n azureoperator-system --create-namespace -f values.yaml
```

### Configuration

The following table lists the configurable parameters of the azure-service-operator chart and its dependency chart, aad-pod-identity.

| Parameter                  | Description              | Default              |
|:---------------------------|:-------------------------|:---------------------|
| `azureSubscriptionID`  | Azure Subscription ID | `` |
| `azureTenantID`  | Azure Tenant ID | `` |
| `azureClientID`  | Azure Service Principal Client ID | `` |
| `azureClientSecret`  | Azure Service Principal Client Secret | `` |
| `azureUseMI`  | Set to True if using Managed Identity for authentication | `False` |
| `installAadPodIdentity` | Set to `True` to install [aad-pod-identity](https://github.com/Azure/aad-pod-identity), or leave false if you already have it installed. You will have to create the `AzureIdentity` and `AzureIdentityBinding` yourself if you already have it installed | `False` |
| `azureOperatorKeyvault`  | Set this value with the name of your Azure Key Vault resource if you prefer to store secrets in Key Vault rather than as Kubernetes secrets (default) | `` |
| `image.repository`  | Image repository | `mcr.microsoft.com/k8s/azureserviceoperator:latest` |
| `cloudEnvironment`  | Set the cloud environment, possible values include: AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud, AzureGermanCloud | `AzurePublicCloud` |
| `createNamespace`  | Set to True if you would like the namespace autocreated, otherwise False if you have an existing namespace. If using an existing namespace, the `namespace` field must also be updated | `True` |
| `namespace`  | Configure a custom namespace to deploy the operator into | `azureoperator-system` |
| `aad-pod-identity.azureIdentity.resourceID`  | The resource ID for your managed identity | `` |
| `aad-pod-identity.azureIdentity.clientID`  | The client ID for your managed identity | `` |
