# Use managed identities for authentication in the operator

## Before you begin
This guide assumes that you have an existing Azure Kubernetes Service cluster. Instructions for other Kubernetes clusters are in progress. To run the provided script, you will need [Go](https://golang.org/) installed.

In this guide, we'll show you how to use managed identities instead of service principal authentication when using Azure Service Operator.

1. Create a managed identity in a specified resource group within your subscription. This is the resource group your Azure resources will be created in.
2. Give the Service Principal associated with the AKS cluster control over managing the identity we just created.
3. Give the Managed Identity we just created authorization to provision resources in our subscription.

## Create a managed identity
We've provided a script [here](./../../cli/main.go) to accomplish the steps of creating a managed identity and giving it the right permissions. Just ensure you have the following values when you run it.

|         |         |
|---------|---------|
|**resourcegroup_name**               |   Resource group where you want to create the managed identity      |
|**managed_identity_name**            | Name of the managed identity you'd like to create
|**sub_id**                           | Azure Subscription ID
|**aks_cluster_service_principal_id** | Principal ID of your AKS cluster's identity 

To find the principal ID of your AKS cluster, you can run the following commands:

- For an AKS cluster with service principal, this is:
 ```az aks show -g <AKSResourceGroup> -n <AKSClusterName> --query servicePrincipalProfile.clientId -otsv```

- For an AKS cluster with managed identity, this is: ```az aks show -g <AKSResourceGroup> -n <AKSClusterName> --query identityProfile.kubeletidentity.clientId -otsv```

Run the script as follows:

```sh
go run cli/main.go createmi \
  -g resourcegroup_name \
  -i managed_identity_name \
  -p aks_cluster_service_principal_id \
  --subscription subscription_id
```

## Install Azure Active Directory Pod Identity
After creating the managed identity you've specified, the script will then install Azure Active Directory (AAD) pod identity. This ensures that Azure Service Operator has the permissions of the identity that we just created.

```bash
Install AAD Pod Identity:
make install-aad-pod-identity

cat <<EOF | kubectl apply -f -
apiVersion: "aadpodidentity.k8s.io/v1"
kind: AzureIdentity
metadata:
  name: aso-managed-id1
  namespace: azureoperator-system
spec:
  type: 0
  resourceID: /subscriptions/7060bca0-7a3c-44bd-b54c-4bb1e9facfac/resourcegroups/my-rg-test/providers/Microsoft.ManagedIdentity/userAssignedIdentities/aso-managed-id1
  clientID: 5bb69783-06a1-4277-82b7-97820b357503
---
apiVersion: "aadpodidentity.k8s.io/v1"
kind: AzureIdentityBinding
metadata:
  name: aso-identity-binding
  namespace: azureoperator-system
spec:
  azureIdentity: aso-managed-id1
  selector: aso_manager_binding
EOF
```

Let's break down exactly what the script is doing.

1. Install [aad-pod-identity](https://github.com/Azure/aad-pod-identity#1-create-the-deployment).

  ```shell
  make install-aad-pod-identity
  ```


2. Apply the AzureIdentity and Binding manifests. This binds the identity to the Azure Service Operator.

  ```bash
  ➜  azure-service-operator git:(master) ✗ cat <<EOF | kubectl apply -f -
  apiVersion: "aadpodidentity.k8s.io/v1"
  kind: AzureIdentity
  metadata:
    name: aso-managed-id1
    namespace: azureoperator-system
  spec:
    type: 0
    resourceID: /subscriptions/7060bca0-7a3c-44bd-b54c-4bb1e9facfac/resourcegroups/my-rg-test/providers/Microsoft.ManagedIdentity/userAssignedIdentities/aso-managed-id1
    clientID: 5bb69783-06a1-4277-82b7-97820b357503
  ---
  apiVersion: "aadpodidentity.k8s.io/v1"
  kind: AzureIdentityBinding
  metadata:
    name: aso-identity-binding
    namespace: azureoperator-system
  spec:
    azureIdentity: aso-managed-id1
    selector: aso_manager_binding
  EOF
  ```

  Now that you've completed the above steps, the Azure Service Operator you've deployed now uses a managed identity for authentication. Use the following command to confirm the above steps worked.

  ```
  TBD
