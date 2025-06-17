---
title: Credential Format
linktitle: Credential Format
---

Azure Service Operator supports four different styles of authentication today.

Each section below dives into one of these authentication options, including examples for how to set it up and
use it at the different [credential scopes]( {{< relref "credential-scope" >}} ).

## Allowed credential fields

These fields are common across all [credential scopes]( {{< relref "credential-scope" >}} ).

- [AZURE_SUBSCRIPTION_ID]( {{< relref "aso-controller-settings-options" >}}/#azure_subscription_id)
- [AZURE_TENANT_ID]( {{< relref "aso-controller-settings-options" >}}/#azure_tenant_id)
- [AZURE_CLIENT_ID]( {{< relref "aso-controller-settings-options" >}}/#azure_client_id)
- [AZURE_CLIENT_SECRET]( {{< relref "aso-controller-settings-options" >}}/#azure_client_secret)
- [AZURE_CLIENT_CERTIFICATE]( {{< relref "aso-controller-settings-options" >}}/#azure_client_certificate)
- [AZURE_CLIENT_CERTIFICATE_PASSWORD]( {{< relref "aso-controller-settings-options" >}}/#azure_client_certificate_password)
- [AZURE_ADDITIONAL_TENANTS]( {{< relref "aso-controller-settings-options" >}}/#azure_additional_tenants)
- [ENTRA_APP_ID]( {{< relref "aso-controller-settings-options" >}}/#entra_app_id)

Note that the global credential scope has fields that can be set in addition to the fields documented above.

## Managed Identity (via workload identity)

See [Azure Workload Identity](https://github.com/Azure/azure-workload-identity) for details about the workload identity project.

**Managed Identity (via workload identity) is the recommended authentication mode for production use-cases**.

### Prerequisites

1. An existing Azure Service Principal or Managed Identity. The setup is the same regardless of which you choose.
2. The [Azure CLI](https://docs.microsoft.com/cli/azure/install-azure-cli).
3. An OIDC endpoint associated with your cluster. See [how to enable OIDC on AKS](https://learn.microsoft.com/en-us/azure/aks/use-oidc-issuer).

Use the following Bash commands to set the environment variables containing the workload identity secret (customize with your values):

```bash
export AZURE_CLIENT_ID="00000000-0000-0000-0000-00000000000"       # The client ID (sometimes called App Id) of the Service Principal, or the Client ID of the Managed Identity with which you are using Workload Identity.
export AZURE_SUBSCRIPTION_ID="00000000-0000-0000-0000-00000000000" # The Azure Subscription ID the identity is in.
export AZURE_TENANT_ID="00000000-0000-0000-0000-00000000000"       # The Azure AAD Tenant the identity/subscription is associated with.
export SERVICE_ACCOUNT_ISSUER="https://oidc.prod-aks.azure.com/00000000-0000-0000-0000-00000000000/" # The OIDC endpoint for your cluster in this example AKS
```

### Configure trust

Establish trust between your OIDC issuer URL and the backing Service Principal or Managed Identity. See [how it works](https://docs.microsoft.com/en-us/azure/active-directory/develop/workload-identity-federation#how-it-works) for details.

{{< tabpane text=true left=true >}}
{{% tab header="**Kind**:" disabled=true /%}}
{{% tab header="Managed Identity" %}}

Set the following additional environment variables:

```bash
export MI_RESOURCE_GROUP="my-rg"  # The resource group containing the managed identity that will be used by ASO
export MI_NAME="my-mi"            # The name of the managed identity that will be used by ASO
```

Create the Federated Identity Credential registering your service account with AAD:

```bash
az identity federated-credential create --name aso-federated-credential --identity-name ${MI_NAME} --resource-group ${MI_RESOURCE_GROUP} --issuer ${SERVICE_ACCOUNT_ISSUER} --subject "system:serviceaccount:azureserviceoperator-system:azureserviceoperator-default" --audiences "api://AzureADTokenExchange"
```

{{% /tab %}}
{{% tab header="Service Principal" %}}

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

{{% /tab %}}
{{< /tabpane >}}

### Create the secret

{{< tabpane text=true left=true >}}
{{% tab header="**Scope**:" disabled=true /%}}
{{% tab header="Global" %}}

If installing ASO for the first time, you can pass these values via Helm arguments:

```bash
helm upgrade --install --devel aso2 aso2/azure-service-operator \
        --create-namespace \
        --namespace=azureserviceoperator-system \
        --set azureSubscriptionID=$AZURE_SUBSCRIPTION_ID \
        --set azureTenantID=$AZURE_TENANT_ID \
        --set azureClientID=$AZURE_CLIENT_ID \
        --set useWorkloadIdentityAuth=true \
        --set crdPattern='resources.azure.com/*;containerservice.azure.com/*;keyvault.azure.com/*;managedidentity.azure.com/*;eventhub.azure.com/*'
```

See [CRD management]( {{< relref "crd-management" >}} ) for more details about `crdPattern`.

Otherwise, create or update the `aso-controller-settings` secret:

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
 USE_WORKLOAD_IDENTITY_AUTH: "true"
EOF
```

**Note:** The `aso-controller-settings` secret contains more configuration than just the global credential.
If ASO was already installed on your cluster and you are updating the `aso-controller-settings` secret, ensure that
[other values]( {{< relref "aso-controller-settings-options" >}} ) in that secret are not being overwritten.

{{% /tab %}}
{{% tab header="Namespace" %}}

Create the `aso-credential` secret in your namespace:

```bash
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
 name: aso-credential
 namespace: my-namespace
stringData:
 AZURE_SUBSCRIPTION_ID: "$AZURE_SUBSCRIPTION_ID"
 AZURE_TENANT_ID:    "$AZURE_TENANT_ID"
 AZURE_CLIENT_ID:    "$AZURE_CLIENT_ID"
EOF
```

**Note:** Each credential (both namespaced and per-resource) you create must have a trust relationship between your OIDC
issuer URL and the backing Service Principal or Managed Identity. See [how to configure trust](#configure-trust) for more details.

{{% /tab %}}
{{% tab header="Resource" %}}

Create a per-resource secret. We'll use `my-resource-secret`:

```bash
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
 name: my-resource-secret
 namespace: my-namespace
stringData:
 AZURE_SUBSCRIPTION_ID: "$AZURE_SUBSCRIPTION_ID"
 AZURE_TENANT_ID:    "$AZURE_TENANT_ID"
 AZURE_CLIENT_ID:    "$AZURE_CLIENT_ID"
EOF
```

Create the ASO resource referring to `my-resource-secret`. We show a `ResourceGroup` here, but any ASO resource will work.

```bash
cat <<EOF | kubectl apply -f -
apiVersion: resources.azure.com/v1api20200601
kind: ResourceGroup
metadata:
  name: aso-sample-rg
  namespace: default
  annotations:
    serviceoperator.azure.com/credential-from: my-resource-secret
spec:
  location: westcentralus
EOF
```

**Note:** Each credential (both namespaced and per-resource) you create must have a trust relationship between your OIDC
issuer URL and the backing Service Principal or Managed Identity. See [how to configure trust](#configure-trust) for more details.

{{% /tab %}}
{{< /tabpane >}}

## Service Principal using a Client Secret

### Prerequisites

1. An existing Azure Service Principal.

To use Service Principal authentication with **client secret**, create a secret with the `AZURE_CLIENT_ID` and `AZURE_CLIENT_SECRET` keys set.

For more information about Service Principals, see [Work with Azure service principal using the Azure CLI](https://learn.microsoft.com/en-us/cli/azure/azure-cli-sp-tutorial-1).
The `AZURE_CLIENT_ID` is sometimes also called the App ID. The `AZURE_CLIENT_SECRET` is the "password" returned by the command in the previously linked documentation.

Use the following Bash commands to set the environment variables containing the service principal secret (customize with your values):

```bash
export AZURE_CLIENT_ID="00000000-0000-0000-0000-00000000000"       # The client ID (sometimes called App Id) of the Service Principal.
export AZURE_CLIENT_SECRET="00000000-0000-0000-0000-00000000000"   # The client secret of the Service Principal.
export AZURE_SUBSCRIPTION_ID="00000000-0000-0000-0000-00000000000" # The Azure Subscription ID the identity is in.
export AZURE_TENANT_ID="00000000-0000-0000-0000-00000000000"       # The Azure AAD Tenant the identity/subscription is associated with.
```

### Create the secret

{{< tabpane text=true left=true >}}
{{% tab header="**Scope**:" disabled=true /%}}
{{% tab header="Global" %}}

If installing ASO for the first time, you can pass these values via Helm arguments:

```bash
helm upgrade --install --devel aso2 aso2/azure-service-operator \
        --create-namespace \
        --namespace=azureserviceoperator-system \
        --set azureSubscriptionID=$AZURE_SUBSCRIPTION_ID \
        --set azureTenantID=$AZURE_TENANT_ID \
        --set azureClientID=$AZURE_CLIENT_ID \
        --set azureClientSecret=$AZURE_CLIENT_SECRET \
        --set crdPattern='resources.azure.com/*;containerservice.azure.com/*;keyvault.azure.com/*;managedidentity.azure.com/*;eventhub.azure.com/*'
```

See [CRD management]( {{< relref "crd-management" >}} ) for more details about `crdPattern`.

Otherwise, create or update the `aso-controller-settings` secret:

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

**Note:** The `aso-controller-settings` secret contains more configuration than just the global credential.
If ASO was already installed on your cluster and you are updating the `aso-controller-settings` secret, ensure that
[other values]( {{< relref "aso-controller-settings-options" >}} ) in that secret are not being overwritten.

{{% /tab %}}
{{% tab header="Namespace" %}}

Create the `aso-credential` secret in your namespace:

```bash
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
 name: aso-credential
 namespace: my-namespace
stringData:
 AZURE_SUBSCRIPTION_ID: "$AZURE_SUBSCRIPTION_ID"
 AZURE_TENANT_ID: "$AZURE_TENANT_ID"
 AZURE_CLIENT_ID: "$AZURE_CLIENT_ID"
 AZURE_CLIENT_SECRET: "$AZURE_CLIENT_SECRET"
EOF
```

{{% /tab %}}
{{% tab header="Resource" %}}

Create a per-resource secret. We'll use `my-resource-secret`:

```bash
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
 name: my-resource-secret
 namespace: my-namespace
stringData:
 AZURE_SUBSCRIPTION_ID: "$AZURE_SUBSCRIPTION_ID"
 AZURE_TENANT_ID: "$AZURE_TENANT_ID"
 AZURE_CLIENT_ID: "$AZURE_CLIENT_ID"
 AZURE_CLIENT_SECRET: "$AZURE_CLIENT_SECRET"
EOF
```

Create the ASO resource referring to `my-resource-secret`. We show a `ResourceGroup` here, but any ASO resource will work.

```bash
cat <<EOF | kubectl apply -f -
apiVersion: resources.azure.com/v1api20200601
kind: ResourceGroup
metadata:
  name: aso-sample-rg
  namespace: default
  annotation:
    serviceoperator.azure.com/credential-from: my-resource-secret
spec:
  location: westcentralus
EOF
```

{{% /tab %}}
{{< /tabpane >}}

## Service Principal using a Client Certificate

### Prerequisites

1. An existing Azure Service Principal.
2. X.509 certificate in ASCII format such as PEM, CER, or DER.

To use Service Principal authentication via client certificate, create a secret with the `AZURE_CLIENT_ID`, `AZURE_CLIENT_CERTIFICATE` and `AZURE_CLIENT_CERTIFICATE_PASSWORD`(optional) keys set.

For more information about creating Service Principals with certificate, see [creating an Azure Service Principal using certificate based authentication](https://learn.microsoft.com/en-us/cli/azure/azure-cli-sp-tutorial-3?tabs=concepts).
The `AZURE_CLIENT_ID` is sometimes also called the App ID. The `AZURE_CLIENT_CERTIFICATE` is the _certificate_ returned by the command in the previously linked documentation.

Use the following Bash commands to set the environment variables containing the service principal certificate secret (customize with your values):

```bash
export AZURE_CLIENT_ID="00000000-0000-0000-0000-00000000000"          # The client ID (sometimes called App Id) of the Service Principal.
export AZURE_SUBSCRIPTION_ID="00000000-0000-0000-0000-00000000000"    # The Azure Subscription ID the identity is in.
export AZURE_TENANT_ID="00000000-0000-0000-0000-00000000000"          # The Azure AAD Tenant the identity/subscription is associated with.
export AZURE_CLIENT_CERTIFICATE=`cat path/to/certFile.pem`                    # The client certificate of the Service Principal.
export AZURE_CLIENT_CERTIFICATE_PASSWORD="myPrivateKeyValue"          # The private key for the above certificate (optional)
```

### Create the secret

{{< tabpane text=true left=true >}}
{{% tab header="**Scope**:" disabled=true /%}}
{{% tab header="Global" %}}

If installing ASO for the first time, you can pass these values via Helm arguments:

```bash
helm upgrade --install --devel aso2 aso2/azure-service-operator \
        --create-namespace \
        --namespace=azureserviceoperator-system \
        --set azureSubscriptionID=$AZURE_SUBSCRIPTION_ID \
        --set azureTenantID=$AZURE_TENANT_ID \
        --set azureClientID=$AZURE_CLIENT_ID \
        --set azureClientCertificatePassword=$AZURE_CLIENT_CERTIFICATE_PASSWORD \
        --set crdPattern='resources.azure.com/*;containerservice.azure.com/*;keyvault.azure.com/*;managedidentity.azure.com/*;eventhub.azure.com/*'
```

See [CRD management]( {{< relref "crd-management" >}} ) for more details about `crdPattern`.

Otherwise, create or update the `aso-controller-settings` secret:

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
 AZURE_CLIENT_CERTIFICATE: "$AZURE_CLIENT_CERTIFICATE"
 AZURE_CLIENT_CERTIFICATE_PASSWORD: "$AZURE_CLIENT_CERTIFICATE_PASSWORD"
EOF
```

**Note:** The `aso-controller-settings` secret contains more configuration than just the global credential.
If ASO was already installed on your cluster and you are updating the `aso-controller-settings` secret, ensure that
[other values]( {{< relref "aso-controller-settings-options" >}} ) in that secret are not being overwritten.

{{% /tab %}}
{{% tab header="Namespace" %}}

Create the `aso-credential` secret in your namespace:

```bash
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
 name: aso-credential
 namespace: my-namespace
stringData:
 AZURE_SUBSCRIPTION_ID: "$AZURE_SUBSCRIPTION_ID"
 AZURE_TENANT_ID: "$AZURE_TENANT_ID"
 AZURE_CLIENT_ID: "$AZURE_CLIENT_ID"
 AZURE_CLIENT_CERTIFICATE: "$AZURE_CLIENT_CERTIFICATE"
 AZURE_CLIENT_CERTIFICATE_PASSWORD: "$AZURE_CLIENT_CERTIFICATE_PASSWORD"
EOF
```

{{% /tab %}}
{{% tab header="Resource" %}}

Create a per-resource secret. We'll use `my-resource-secret`:

```bash
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
 name: my-resource-secret
 namespace: my-namespace
stringData:
 AZURE_SUBSCRIPTION_ID: "$AZURE_SUBSCRIPTION_ID"
 AZURE_TENANT_ID: "$AZURE_TENANT_ID"
 AZURE_CLIENT_ID: "$AZURE_CLIENT_ID"
 AZURE_CLIENT_CERTIFICATE: "$AZURE_CLIENT_CERTIFICATE"
 AZURE_CLIENT_CERTIFICATE_PASSWORD: "$AZURE_CLIENT_CERTIFICATE_PASSWORD"
EOF
```

Create the ASO resource referring to `my-resource-secret`. We show a `ResourceGroup` here, but any ASO resource will work.

```bash
cat <<EOF | kubectl apply -f -
apiVersion: resources.azure.com/v1api20200601
kind: ResourceGroup
metadata:
  name: aso-sample-rg
  namespace: default
  annotation:
    serviceoperator.azure.com/credential-from: my-resource-secret
spec:
  location: westcentralus
EOF
```

{{% /tab %}}
{{< /tabpane >}}

## Managed Identity (via IMDS on Azure infrastructure)

### Prerequisites

1. An existing Azure Managed Identity.
2. ASO running on Azure infrastructure (such as an AKS cluster) with the Managed Identity assigned to that infrastructure.

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

### Create the secret

{{< tabpane text=true left=true >}}
{{% tab header="**Scope**:" disabled=true /%}}
{{% tab header="Global" %}}

If installing ASO for the first time, you can pass these values via Helm arguments:

```bash
helm upgrade --install --devel aso2 aso2/azure-service-operator \
     --create-namespace \
     --namespace=azureserviceoperator-system \
     --set azureSubscriptionID=$AZURE_SUBSCRIPTION_ID \
     --set azureClientID=${IDENTITY_CLIENT_ID} \
     --set crdPattern='resources.azure.com/*;containerservice.azure.com/*;keyvault.azure.com/*;managedidentity.azure.com/*;eventhub.azure.com/*'
```

See [CRD management]( {{< relref "crd-management" >}} ) for more details about `crdPattern`.

Create or update the `aso-controller-settings` secret:

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

**Note:** The `aso-controller-settings` secret contains more configuration than just the global credential.
If ASO was already installed on your cluster and you are updating the `aso-controller-settings` secret, ensure that
[other values]( {{< relref "aso-controller-settings-options" >}} ) in that secret are not being overwritten.

{{% /tab %}}
{{% tab header="Namespace" %}}

Create the `aso-credential` secret in your namespace:

```bash
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
 name: aso-credential
 namespace: my-namespace
stringData:
 AZURE_SUBSCRIPTION_ID: "$AZURE_SUBSCRIPTION_ID"
 AZURE_TENANT_ID:       "$AZURE_TENANT_ID"
 AZURE_CLIENT_ID:       "$IDENTITY_CLIENT_ID"
 AUTH_MODE:             "podidentity"
EOF
```

{{% /tab %}}
{{% tab header="Resource" %}}

Create a per-resource secret. We'll use `my-resource-secret`:

```bash
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
 name: my-resource-secret
 namespace: my-namespace
stringData:
 AZURE_SUBSCRIPTION_ID: "$AZURE_SUBSCRIPTION_ID"
 AZURE_TENANT_ID:       "$AZURE_TENANT_ID"
 AZURE_CLIENT_ID:       "$IDENTITY_CLIENT_ID"
 AUTH_MODE:             "podidentity"
EOF
```

Create the ASO resource referring to `my-resource-secret`. We show a `ResourceGroup` here, but any ASO resource will work.

```bash
cat <<EOF | kubectl apply -f -
apiVersion: resources.azure.com/v1api20200601
kind: ResourceGroup
metadata:
  name: aso-sample-rg
  namespace: default
  annotations:
    serviceoperator.azure.com/credential-from: my-resource-secret
spec:
  location: westcentralus
EOF
```

{{% /tab %}}
{{< /tabpane >}}

## [Deprecated] Managed Identity (aad-pod-identity)

> **This authentication mechanism still works but is deprecated. See [Managed Identity (via workload identity)](#managed-identity-via-workload-identity) for the new way**

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

### Create the secret

{{< tabpane text=true left=true >}}
{{% tab header="**Scope**:" disabled=true /%}}
{{% tab header="Global" %}}

If installing ASO for the first time, you can pass these values via Helm arguments:

```bash
helm upgrade --install --devel aso2 aso2/azure-service-operator \
     --create-namespace \
     --namespace=azureserviceoperator-system \
     --set azureSubscriptionID=$AZURE_SUBSCRIPTION_ID \
     --set aadPodIdentity.enable=true \
     --set aadPodIdentity.azureManagedIdentityResourceId=${IDENTITY_RESOURCE_ID} \
     --set azureClientID=${IDENTITY_CLIENT_ID} \
     --set crdPattern='resources.azure.com/*;containerservice.azure.com/*;keyvault.azure.com/*;managedidentity.azure.com/*;eventhub.azure.com/*'
```

See [CRD management]( {{< relref "crd-management" >}} ) for more details about `crdPattern`.

Otherwise, if deploying manually:

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

Create or update the `aso-controller-settings` secret:

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

**Note:** The `aso-controller-settings` secret contains more configuration than just the global credential.
If ASO was already installed on your cluster and you are updating the `aso-controller-settings` secret, ensure that
[other values]( {{< relref "aso-controller-settings-options" >}} ) in that secret are not being overwritten.

{{% /tab %}}
{{% tab header="Namespace" %}}

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

Create the `aso-credential` secret in your namespace:

```bash
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
 name: aso-credential
 namespace: my-namespace
stringData:
 AZURE_SUBSCRIPTION_ID: "$AZURE_SUBSCRIPTION_ID"
 AZURE_TENANT_ID:       "$AZURE_TENANT_ID"
 AZURE_CLIENT_ID:       "$IDENTITY_CLIENT_ID"
 AUTH_MODE:             "podidentity"
EOF
```

{{% /tab %}}
{{% tab header="Resource" %}}

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

Create a per-resource secret. We'll use `my-resource-secret`:

```bash
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
 name: my-resource-secret
 namespace: my-namespace
stringData:
 AZURE_SUBSCRIPTION_ID: "$AZURE_SUBSCRIPTION_ID"
 AZURE_TENANT_ID:       "$AZURE_TENANT_ID"
 AZURE_CLIENT_ID:       "$IDENTITY_CLIENT_ID"
 AUTH_MODE:             "podidentity"
EOF
```

Create the ASO resource referring to `my-resource-secret`. We show a `ResourceGroup` here, but any ASO resource will work.

```bash
cat <<EOF | kubectl apply -f -
apiVersion: resources.azure.com/v1api20200601
kind: ResourceGroup
metadata:
  name: aso-sample-rg
  namespace: default
  annotations:
    serviceoperator.azure.com/credential-from: my-resource-secret
spec:
  location: westcentralus
EOF
```

## User Assigned Identity Credentials

### General
{{% alert title="Warning" color="warning" %}}
This option is only available for 1st party Microsoft applications who have access to the msi data-plane.
{{% /alert %}}

This authentication type is similar to user assigned managed identity authentication combined with client certificate
authentication. As a 1st party Microsoft application, one has access to pull a user assigned managed identity's backing
certificate information from the MSI data plane. Using this data, a user can authenticate to Azure Cloud.

### Prerequisites
A JSON file with information from the user assigned managed identity. It should be in this format:
```json
        {
            "client_id": "0998...",
            "client_secret": "MIIKUA...",
            "client_secret_url": "https://control...",
            "tenant_id": "93b...",
            "object_id": "ae...",
            "resource_id": "/subscriptions/...",
            "authentication_endpoint": "https://login.microsoftonline.com/",
            "mtls_authentication_endpoint": "https://login.microsoftonline.com/",
            "not_before": "2025-02-07T13:29:00Z",
            "not_after": "2025-05-08T13:29:00Z",
            "renew_after": "2025-03-25T13:29:00Z",
            "cannot_renew_after": "2025-08-06T13:29:00Z"
        }
```

Note, the client secret should be a base64 encoded certificate.

The steps to get this information from the MSI data plane are as follows:
1. Make an unauthenticated GET or POST (no Authorization request headers) on the x-ms-identity-url received from ARM to get the token authority and, on older api versions, resource.
2. Get an Access Token from Azure AD using your Resource Provider applicationId and Certificate. The applicationId should match the one you added to your manifest. The response should give you an access token.
3. Perform a GET or POST to MSI on the same URL from earlier to get the Credentials using this bearer token.

The only required environment variable is AZURE_USER_ASSIGNED_IDENTITY_CREDENTIALS.
```bash
export AZURE_USER_ASSIGNED_IDENTITY_CREDENTIALS="/path/to/credentials-file.json" # The file path to the msi data plane credentials in a JSON file format.
```

It is expected this JSON file is available in a volume on the pod needing to authenticate with Azure cloud with this
authentication method. For example, if the credentials were stored as a secret in Azure Key Vault, one could use the 
Secrets CSI Driver and a SecretProviderClass custom resource to automatically have the file mounted into a volume on a 
pod.

{{% alert title="Warning" color="warning" %}}
No option is exposed to configure additional deployment volumes in the Helm chart because it is reserved for use by Microsoft 1st party access and not intended for 3rd party use.
{{% /alert %}}

The cloud configuration defaults to Azure Public. If you want to configure the cloud configuration, you will need to 
export at least AZURE_AUTHORITY_HOST. Both AZURE_RESOURCE_MANAGER_AUDIENCE and AZURE_RESOURCE_MANAGER_ENDPOINT must be 
specified if you wish to set values for these two fields.
(see https://learn.microsoft.com/en-us/cli/azure/manage-clouds-azure-cli#list-available-clouds for more info):
```bash
export AZURE_AUTHORITY_HOST="https://login.microsoftonline.com/" # The URL of the Entra authority.
export AZURE_RESOURCE_MANAGER_AUDIENCE=""                        # The Azure Resource Manager Entra audience.
export AZURE_RESOURCE_MANAGER_ENDPOINT=""                        # The Azure Resource Manager endpoint.
```

Create or update the `aso-controller-settings` secret:

```bash
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
 name: aso-controller-settings
 namespace: azureserviceoperator-system
stringData:
 AZURE_USER_ASSIGNED_IDENTITY_CREDENTIALS: "$AZURE_USER_ASSIGNED_IDENTITY_CREDENTIALS"
 AZURE_AUTHORITY_HOST: "$AZURE_AUTHORITY_HOST"
 AZURE_RESOURCE_MANAGER_AUDIENCE: "$AZURE_RESOURCE_MANAGER_AUDIENCE"
 AZURE_RESOURCE_MANAGER_ENDPOINT: "$AZURE_RESOURCE_MANAGER_ENDPOINT"
EOF
```

**Note:** The `aso-controller-settings` secret contains more configuration than just the global credential.
If ASO was already installed on your cluster and you are updating the `aso-controller-settings` secret, ensure that
[other values]( {{< relref "aso-controller-settings-options" >}} ) in that secret are not being overwritten.

{{% /tab %}}
{{< /tabpane >}}
