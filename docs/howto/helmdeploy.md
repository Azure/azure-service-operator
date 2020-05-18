## Deploy using Helm charts

### Installing

#### Pre-Install

This Helm chart contains certificates that depend on [cert-manager](https://cert-manager.io/docs/installation/kubernetes/) to be installed. Install cert-manager prior to installing chart:

```
kubectl create namespace cert-manager
kubectl label namespace cert-manager cert-manager.io/disable-validation=true
kubectl apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v0.12.0/cert-manager.yaml
```

You can use the below command to check if the cert manager pods are ready. The cert manager pods should be running before proceeding to the next step.

```
kubectl rollout status -n cert-manager deploy/cert-manager-webhook
```

#### Helm

Install the latest (3.x+) [Helm](https://helm.sh/docs/intro/install/) on your machine.

Add the helm repo:
```console
helm repo add azureserviceoperator https://raw.githubusercontent.com/Azure/azure-service-operator/master/charts --username <github username> --password <github personal access token>
```

### Getting Started

Copy the `values.yaml` file from this directory and in the following steps fill in the requisite values.

First, set the following variables to your Azure Tenant ID and Subscription ID:
```
azureTenantID: 00000000-0000-0000-0000-000000000000
azureSubscriptionID: 00000000-0000-0000-0000-000000000000
```

#### Authentication

Next, choose one of the following authentication methods, and set its appropriate variables.

##### Service Principal

Set the following variables to your client ID and secret values:
```
azureClientID: 00000000-0000-0000-0000-000000000000
azureClientSecret: 00000000-0000-0000-0000-000000000000
```

##### Managed Identity

Set the following Helm Chart values:
```
azureUseMI: True
aad-pod-identity:
    azureIdentity:
        resourceID: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/<rg>/providers/Microsoft.ManagedIdentity/userAssignedIdentities/<identity>"
        clientID: "00000000-0000-0000-0000-000000000000"
```

Follow the instructions [here](../../docs/deploy.md) to create an identity and assign it the correct permissions.

#### Keyvault

By default, secrets will be stored as Kubernetes secrets. If you wish to store them in KeyVault instead, follow the instructions [here](../../docs/deploy.md).

Then, set the following chart value to your KeyVault name:
```
azureOperatorKeyvault: OperatorSecretKeyVault
```

#### Install Chart

If you are deploying into an already created namespace, be sure to set the following variable to false:
```
createNamespace: False
```

and specify the namespace name:
```
namespace: your-namespace
```

Finally, install the chart with your added values. The chart can be installed by using a values file or environment variables.
```
helm upgrade --install aso azureserviceoperator/azure-service-operator -f values.yaml
```

```
helm upgrade --install aso azureserviceoperator/azure-service-operator \
    --set azureSubscriptionID=$AZURE_SUBSCRIPTION_ID \
    --set azureTenantID=$AZURE_TENANT_ID \
    --set azureClientID=$AZURE_CLIENT_ID \
    --set azureClientSecret=$AZURE_CLIENT_SECRET \
    --set azureUseMI=$AZURE_USE_MI \
    --set azureOperatorKeyvault=$AZURE_OPERATOR_KEYVAULT \
    --set createNamespace=False
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
| `azureOperatorKeyvault`  | Set this value with the name of your Azure Key Vault resource if you prefer to store secrets in Key Vault rather than as Kubernetes secrets (default) | `` |
| `image.repository`  | Image repository | `mcr.microsoft.com/k8s/azure-service-operator:0.0.20258` |
| `cloudEnvironment`  | Set the cloud environment, possible values include: AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud, AzureGermanCloud | `AzurePublicCloud` |
| `createNamespace`  | Set to True if you would like the namespace autocreated, otherwise False if you have an existing namespace. If using an existing namespace, the `namespace` field must also be updated | `True` |
| `namespace`  | Configure a custom namespace to deploy the operator into | `azureoperator-system` |
| `aad-pod-identity.azureIdentity.resourceID`  | The resource ID for your managed identity | `` |
| `aad-pod-identity.azureIdentity.clientID`  | The client ID for your managed identity | `` |
