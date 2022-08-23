---
title: Authentication in Azure Service Operator v2
linktitle: Authentication
---

Azure Service Operator supports three different styles of authentication today.

1. Service Principal
2. [azure-workload-identity](https://github.com/Azure/azure-workload-identity) authentication (OIDC + Managed Identity or Service Principal)
3. [Deprecated] aad-pod-identity authentication (Managed Identity)

## Service Principal

### Prerequisites
1. An existing Azure Service Principal.

To use Service Principal authentication, specify an `aso-controller-settings` secret with the `AZURE_CLIENT_ID` and `AZURE_CLIENT_SECRET` keys set.

For more information about Service Principals, see [creating an Azure Service Principal using the Azure CLI](https://docs.microsoft.com/cli/azure/create-an-azure-service-principal-azure-cli#password-based-authentication).
The `AZURE_CLIENT_ID` is sometimes also called the App ID. The `AZURE_CLIENT_SECRET` is the "password" returned by the command in the previously linked documentation.

Use the following Bash script to set the environment variables for the `aso-controller-settings` secret:
```bash
export AZURE_CLIENT_ID="00000000-0000-0000-0000-00000000000"       # The client ID (sometimes called App Id) of the Service Principal.
export AZURE_CLIENT_SECRET="00000000-0000-0000-0000-00000000000"   # The client secret of the Service Principal.
export AZURE_SUBSCRIPTION_ID="00000000-0000-0000-0000-00000000000" # The Azure Subscription ID the identity is in.
export AZURE_TENANT_ID="00000000-0000-0000-0000-00000000000"       # The Azure AAD Tenant the identity/subscription is associated with.
```

Create the `aso-controller-settings` secret:
```bash
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
 name: aso-controller-settings
 namespace: azureserviceoperator-system
stringData:
 AZURE_SUBSCRIPTION_ID: "$AZURE_SUBSCRIPTION_ID"
 AZURE_TENANT_ID: "$AZURE_TENANT_ID"
 AZURE_CLIENT_ID: "$AZURE_CLIENT_ID"
 AZURE_CLIENT_SECRET: "$AZURE_CLIENT_SECRET"
EOF
```

## Azure Workload Identity

### Prerequisites
1. An existing Azure Service Principal or Managed Identity. The setup is the same regardless of which you choose.
2. The [Azure CLI](https://docs.microsoft.com/cli/azure/install-azure-cli).
3. The [Azure Workload Identity](https://github.com/Azure/azure-workload-identity) webhook installed into your cluster. See [Azure Workload Identity installation](https://azure.github.io/azure-workload-identity/docs/installation.html). Note that setup involves two steps, getting your clusters OIDC issuer URL and installing the [Azure Workload Identity mutating webhook](https://azure.github.io/azure-workload-identity/docs/installation/mutating-admission-webhook.html).

Use the following Bash script to set the environment variables required for the below commands:
```bash
export AZURE_CLIENT_ID="00000000-0000-0000-0000-00000000000"       # The client ID (sometimes called App Id) of the Service Principal, or the Client ID of the Managed Identity with which you are using Workload Identity.
export AZURE_SUBSCRIPTION_ID="00000000-0000-0000-0000-00000000000" # The Azure Subscription ID the identity is in.
export AZURE_TENANT_ID="00000000-0000-0000-0000-00000000000"       # The Azure AAD Tenant the identity/subscription is associated with.
export SERVICE_ACCOUNT_ISSUER="https://oidc.prod-aks.azure.com/00000000-0000-0000-0000-00000000000/" # The OIDC endpoint for your cluster in this example AKS
```

### Get the Application Object ID of the identity

For Managed Identity: `export APPLICATION_OBJECT_ID=$(az resource show --id /subscriptions/${AZURE_SUBSCRIPTION_ID}/resourceGroups/{my resource group name}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{my managed identity name} --query "properties.principalId" -o tsv | tr -d '[:space:]')`

### Configure trust

Establish trust between your OIDC issuer URL and the backing Service Principal or Managed Identity. See [how it works](https://docs.microsoft.com/en-us/azure/active-directory/develop/workload-identity-federation#how-it-works) for details.

#### Service Principal

```bash
export APPLICATION_OBJECT_ID="$(az ad app show --id ${AZURE_CLIENT_ID} --query id -otsv)"

cat <<EOF > params.json
{
  "name": "aso-federated-credential",
  "issuer": "${SERVICE_ACCOUNT_ISSUER}",
  "subject": "system:serviceaccount:azureserviceoperator-system:azureserviceoperator-default",
  "description": "Kubernetes service account federated credential",
  "audiences": [
    "api://AzureADTokenExchange"
  ]
}
EOF

az ad app federated-credential create --id ${APPLICATION_OBJECT_ID} --parameters @params.json
```

#### Managed Identity

Set the following additional environment variables:
```bash
export MI_RESOURCE_GROUP="my-rg"  # The resource group containing the managed identity that will be used by ASO
export MI_NAME="my-mi"            # The name of the managed identity that will be used by ASO
```

```bash
cat <<EOF > body.json
{
  "name": "aso-federated-credential",
  "type":"Microsoft.ManagedIdentity/userAssignedIdentities/federatedIdentityCredentials",
  "properties": {
    "issuer":"${SERVICE_ACCOUNT_ISSUER}",
    "subject":"system:serviceaccount:azureserviceoperator-system:azureserviceoperator-default",
    "audiences": [
      "api://AzureADTokenExchange"
    ]
  }
}
EOF

az rest --method put --url /subscriptions/${AZURE_SUBSCRIPTION_ID}/resourcegroups/${MI_RESOURCE_GROUP}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/${MI_NAME}/federatedIdentityCredentials/aso-federated-credential?api-version=2022-01-31-preview --body @body.json
```

### Set ASO to use the Service Account token

#### Helm
When installing ASO via Helm, set `useWorkloadIdentityAuth` to `true`.

```bash
helm upgrade --install --devel aso2 aso2/azure-service-operator \
        --create-namespace \
        --namespace=azureserviceoperator-system \
        --set azureSubscriptionID=$AZURE_SUBSCRIPTION_ID \
        --set azureTenantID=$AZURE_TENANT_ID \
        --set azureClientID=$AZURE_CLIENT_ID \
        --set useWorkloadIdentityAuth=true
```

#### Kubectl

If you installed ASO manually, you can update the existing `ServiceAccount` to use Workload Identity.

Update the ASO service account to use Workload Identity
```bash
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ServiceAccount
metadata:
  annotations:
    azure.workload.identity/client-id: ${AZURE_CLIENT_ID}
  labels:
    azure.workload.identity/use: "true"
  name: azureserviceoperator-default
  namespace: azureserviceoperator-system
EOF
```

Update the existing ASO deployment and ensure that the controller pod has the following annotation:
`azure.workload.identity/inject-proxy-sidecar: "true"`

Create the `aso-controller-settings` secret:
```bash
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
 name: aso-controller-settings
 namespace: azureserviceoperator-system
stringData:
 AZURE_SUBSCRIPTION_ID: "$AZURE_SUBSCRIPTION_ID"
 AZURE_TENANT_ID: "$AZURE_TENANT_ID"
 AZURE_CLIENT_ID: "$AZURE_CLIENT_ID"
EOF
```

## [Deprecated] Managed Identity (aad-pod-identity)

> **This authentication mechanism still works but is deprecated. See [Azure Workload Identity](#azure-workload-identity) for the new way**

### Prerequisites
1. An existing Azure Managed Identity.
2. [aad-pod-identity](https://github.com/Azure/aad-pod-identity) installed into your cluster. If you are running ASO on an Azure Kubernetes Service (AKS) cluster, you can instead use the
   [integrated aad-pod-identity](https://docs.microsoft.com/azure/aks/use-azure-ad-pod-identity).

First, set the following environment variables:
```bash
export IDENTITY_RESOURCE_GROUP="myrg"                              # The resource group containing the managed identity.
export IDENTITY_NAME="myidentity"                                  # The name of the identity.
export AZURE_SUBSCRIPTION_ID="00000000-0000-0000-0000-00000000000" # The Azure Subscription ID the identity is in.
export AZURE_TENANT_ID="00000000-0000-0000-0000-00000000000"       # The Azure AAD Tenant the identity/subscription is associated with.
```

Use the `az cli` to get some more details about the identity to use:
```bash
export IDENTITY_CLIENT_ID="$(az identity show -g ${IDENTITY_RESOURCE_GROUP} -n ${IDENTITY_NAME} --query clientId -otsv)"
export IDENTITY_RESOURCE_ID="$(az identity show -g ${IDENTITY_RESOURCE_GROUP} -n ${IDENTITY_NAME} --query id -otsv)"
```

Deploy an `AzureIdentity`:
```bash
cat <<EOF | kubectl apply -f -
apiVersion: "aadpodidentity.k8s.io/v1"
kind: AzureIdentity
metadata:
  name: aso-identity
  namespace: azureserviceoperator-system
spec:
  type: 0
  resourceID: ${IDENTITY_RESOURCE_ID}
  clientID: ${IDENTITY_CLIENT_ID}
EOF
```

Deploy an `AzureIdentityBinding` to bind this identity to the Azure Service Operator manager pod:
```bash
cat <<EOF | kubectl apply -f -
apiVersion: "aadpodidentity.k8s.io/v1"
kind: AzureIdentityBinding
metadata:
  name: aso-identity-binding
  namespace: azureserviceoperator-system
spec:
  azureIdentity: aso-identity
  selector: aso-manager-binding
EOF
```

Deploy the `aso-controller-settings` secret, configured to use the identity:
```bash
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
 name: aso-controller-settings
 namespace: azureserviceoperator-system
stringData:
 AZURE_SUBSCRIPTION_ID: "$AZURE_SUBSCRIPTION_ID"
 AZURE_TENANT_ID: "$AZURE_TENANT_ID"
 AZURE_CLIENT_ID: "$IDENTITY_CLIENT_ID"
EOF
```
