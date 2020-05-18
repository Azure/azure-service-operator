# Using Managed Identity for authentication in the operator

Perform these steps to use Managed Identity instead of Service Principal authentication in the operator.

1. Create a managed identity in a specified resource group within your subscription.
2. Give the Service Principal associated with the AKS cluster control over managing the identity we just created.
3. Give the Managed Identity we just created authorization to provision resources in our subscription

You can accomplish all of the above 3 steps using the script [here](./../../cli/main.go)

Run the script as follows:

```
go run cli/main.go createmi \
  -g resourcegroup_name \
  -i managed_identity_name \
  -p aks_cluster_service_principal_id \
  --subscription sub_id
```

*resourcegroup_name* = Resource Group where you want to create the Managed Identity
*managed_identity_name* = Name of the Managed Identity to create
*sub_id* = Subscription Id
*aks_cluster_service_principal_id* = Principal Id of your AKS cluster's identity

For an AKS cluster with Service Principal, this is:
```az aks show -g <AKSResourceGroup> -n <AKSClusterName> --query servicePrincipalProfile.clientId -otsv```

for an AKS cluster with Managed Identity, this is:
```az aks show -g <AKSResourceGroup> -n <AKSClusterName> --query identityProfile.kubeletidentity.clientId -otsv```

You will see the script run some commands and then will output the next steps to perform as below.

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

4. Install [aad-pod-identity](https://github.com/Azure/aad-pod-identity#1-create-the-deployment)

```shell
make install-aad-pod-identity
```

5. Apply the AzureIdentity and Binding manifests. For this, you can just copy over the command from the above script  as follows

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
