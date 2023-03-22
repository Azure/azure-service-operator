---
title: Authentication in Azure Service Operator v2
linktitle: Authentication
---

Azure Service Operator supports four different styles of authentication today. Each of these options can be used either as a global credential applied to all resources created by the operator (as shown below), or as a per-resource or per-namespace credential as documented in [single-operator-multitenancy](https://azure.github.io/azure-service-operator/introduction/multitenant-deployment/#single-operator-multitenancy-default).

1. Service Principal using Client Secret
2. Service Principal using Client Certificate
3. [Azure-Workload-Identity](https://github.com/Azure/azure-workload-identity) authentication (OIDC + Managed Identity or Service Principal)
4. [Deprecated] aad-pod-identity authentication (Managed Identity)


## Service Principal using Client Secret

### Prerequisites
1. An existing Azure Service Principal.

To use Service Principal authentication with **client secret**, specify an `aso-controller-settings` secret with the `AZURE_CLIENT_ID` and `AZURE_CLIENT_SECRET` keys set.

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

## Service Principal using Client Certificate

### Prerequisites
1. An existing Azure Service Principal.
2. Certificate in ASCII format such as PEM, CER, or DER.

To use Service Principal authentication via client certificate, specify an `aso-controller-settings` secret with the `AZURE_CLIENT_ID`, `AZURE_CLIENT_CERTIFICATE` and `AZURE_CLIENT_CERTIFICATE_PASSWORD`(optional) keys set.

For more information about Service Principals, see [creating an Azure Service Principal using the Azure CLI](https://learn.microsoft.com/cli/azure/create-an-azure-service-principal-azure-cli#certificate-based-authentication).
The `AZURE_CLIENT_ID` is sometimes also called the App ID. The `AZURE_CLIENT_CERTIFICATE` is the "certificate" returned by the command in the previously linked documentation.

Use the following Bash script to set the environment variables for the `aso-controller-settings` secret:
```bash
export AZURE_CLIENT_ID="00000000-0000-0000-0000-00000000000"          # The client ID (sometimes called App Id) of the Service Principal.
export AZURE_SUBSCRIPTION_ID="00000000-0000-0000-0000-00000000000"    # The Azure Subscription ID the identity is in.
export AZURE_TENANT_ID="00000000-0000-0000-0000-00000000000"          # The Azure AAD Tenant the identity/subscription is associated with.
export AZURE_CLIENT_CERTIFICATE=`cat path/to/certFile.pem`                    # The client certificate of the Service Principal.
export AZURE_CLIENT_CERTIFICATE_PASSWORD="myPrivateKeyValue"          # The private key for the above certificate (optional)
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
 AZURE_CLIENT_CERTIFICATE: "$AZURE_CLIENT_CERTIFICATE"
 AZURE_CLIENT_CERTIFICATE_PASSWORD: "$AZURE_CLIENT_CERTIFICATE_PASSWORD"
EOF
```


## Azure Workload Identity

### Prerequisites
1. An existing Azure Service Principal or Managed Identity. The setup is the same regardless of which you choose.
2. The [Azure CLI](https://docs.microsoft.com/cli/azure/install-azure-cli).

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

##### Updating an existing deployment

If you installed ASO manually, you can update the existing `Secret` to use Workload Identity authentication.

Update the `aso-controller-settings` secret to have string data `USE_WORKLOAD_IDENTITY_AUTH: "true"`

Ensure that the `aso-controller-settings` secret has the key `USE_WORKLOAD_IDENTITY_AUTH` set to `true` and restart the ASO pod. 

##### New Deployment

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
 USE_WORKLOAD_IDENTITY_AUTH: "true"
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

#### Manual Deploy

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

#### Helm Chart Deploy

```bash
helm repo add aso2 https://raw.githubusercontent.com/Azure/azure-service-operator/main/v2/charts
helm repo update

helm upgrade --install --devel aso2 aso2/azure-service-operator \
     --create-namespace \
     --namespace=azureserviceoperator-system \
     --set azureSubscriptionID=$AZURE_SUBSCRIPTION_ID \
     --set aadPodIdentity.enable=true \
     --set aadPodIdentity.azureManagedIdentityResourceId=${IDENTITY_RESOURCE_ID} \
     --set azureClientID=${IDENTITY_CLIENT_ID}
```

## Using a credential for ASO with reduced permissions

Most examples of installing ASO suggest using an identity that has Contributor access to the Subscription.
Such broadly scoped access is _**not**_ required to run the operator, it's just the easiest way to set things up and so
is often used in examples. 

Here we discuss a few approaches to restricting the access that the ASO identity has.

### Contributor access at a reduced scope

In some scenarios you may know that ASO will be operating on only certain resource groups. For example, if you have two
teams working in a cluster with ASO, "Marketing" and "Product". If you know ahead of time that the only resource groups these
teams will use are `marketing-dev`, `marketing-prod`, `product-dev` and `product-prod`, you can pre-create the following resources:

* The 4 resource groups.
* An identity for use by ASO, assigned Contributor permission to each of the 4 resource groups.

Now install ASO following the instructions above. Once ASO is installed, adopt the existing resource groups in ASO by applying a YAML in the cluster
for each of the 4 resource groups. Now users of ASO are free to create resources in those resource groups as normal. If they attempt
to use ASO to create a new resource group `foo` they will be rejected by Azure as the ASO identity doesn't have permission to create
arbitrary resource groups in the Subscription.

### Reduced access at Subscription scope

In other scenarios, you may want to give out reduced access at the Subscription scope. This can be done by determining what resources 
in ASO you will be using and creating a custom role scoped to just those permissions. 

For example if you want to use ASO to manage ResourceGroups, Redis and MySQL flexible servers, but _not_ anything else, you could 
[create a role](https://learn.microsoft.com/azure/role-based-access-control/custom-roles-cli) similar to `aso-operator.json`:

```json
{
  "Name": "ASO Operator",
  "IsCustom": true,
  "Description": "Role with access to perform only the operations which we allow ASO to perform",
  "Actions": [
    "Microsoft.Resources/subscriptions/resourceGroups/*",
    "Microsoft.Cache/*",
    "Microsoft.DBforMySQL/*"
  ],
  "NotActions": [
  ],
  "AssignableScopes": [
    "/subscriptions/{subscriptionId1}",
    "/subscriptions/{subscriptionId2}"
  ]
}
```
Then use the az cli to create that custom role definition: `az role definition create --role-definition ~/aso-operator.json`

See [list of resource provider operations](https://learn.microsoft.com/en-us/azure/role-based-access-control/resource-provider-operations) for a comprehensive
list of operations. 

**Note: We strongly recommend giving full permissions to the resource types ASO will be managing.** ASO needs `read`, `write`, and `delete` permissions to a resource
to fully manage it. In some cases, it also needs `action` permissions. It's recommended to give `*` permissions for a given resource type which ASO 
will be managing.

Once the role has been created, assign it to the ASO identity and install the operator as normal. It will be able to
manage any resource of the types you allowed. Any other resources will fail to be created/updated/deleted with a permissions error.
