---
title: Authentication in Azure Service Operator v2
---
# Authentication in Azure Service Operator v2

Azure Service Operator supports two different styles of authentication today.

1. Service Principal
2. aad-pod-identity authentication (managed identity)

## Service Principal

### Prerequisites
1. An existing Azure Service Principal.

To use Service Principal authentication, specify an `aso-controller-settings` secret with `AZURE_CLIENT_ID` and `AZURE_CLIENT_SECRET` set.

* `AZURE_CLIENT_ID` must be set to the Service Principal client ID. This will be a GUID.
* `AZURE_CLIENT_SECRET` must be set to the Service Principal client secret.

For more information about Service Principals, see [creating an Azure Service Principal using the Azure CLI](https://docs.microsoft.com/cli/azure/create-an-azure-service-principal-azure-cli#password-based-authentication).
The `AZURE_CLIENT_ID` is sometimes also called the App ID. The `AZURE_CLIENT_SECRET` is the "password" returned by the command in the previously linked documentation.

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

## Managed Identity (aad-pod-identity)

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
