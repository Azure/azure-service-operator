# Setting up a Managed Identity

## Before you begin
This guide assumes that you have an existing Azure Kubernetes Service cluster and the [Azure CLI](https://github.com/Azure/azure-cli) installed locally.

In this guide we'll show you how to create a Managed Identity and configure it so that it can be used with Azure Service Operator. To do this we will:

1. Create a Managed Identity in a specified resource group within your subscription. This is the resource group your Azure resources will be created in.
2. Give the Service Principal associated with the AKS cluster control over managing the identity we just created.
3. Give the Managed Identity we just created authorization to provision resources in our subscription.

## Create a Managed Identity
Run the command `az identity create -g <ResourceGroup> -n <ManagedIdentity> --subscription <Subscription> -o json` to create a new Managed Identity in the specified subscription.

## Assign AKS Cluster control over the Managed Identity
First we need to get the AKS Cluster identity `ClientID`

- For an AKS cluster with service principal: 
 ```az aks show -g <AKSResourceGroup> -n <AKSClusterName> --query servicePrincipalProfile.clientId -otsv```

- For an AKS cluster with managed identity: 
  ```az aks show -g <AKSResourceGroup> -n <AKSClusterName> --query identityProfile.kubeletidentity.clientId -otsv```

Then assign permission for AKS to place this managed identity onto the nodes in your node pools:
```az role assignment create --role "Managed Identity Operator" --assignee <ServicePrincipal> --scope "<ManagedIdentityID>"```

## Give the Managed Identity permission to create and manage resources in your subscription
`az role assignment create --role "Contributor" --assignee <ManagedIdentityClientID> --scope /subscriptions/<Subscription>`

## Next steps
Now that you've created a managed identity, you need to assign that identity to ASO. This is done using [aad-pod-identity](https://github.com/Azure/aad-pod-identity). 
For details on how to install aad-pod-identity, see [deploying ASO via Helm](./helmdeploy.md) or [deploying ASO from source](./deploy.md).
