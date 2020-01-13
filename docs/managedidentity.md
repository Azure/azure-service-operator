# Using Managed Identity for authentication in the operator

Perform these steps to use Managed Identity instead of Service Principal authentication in the operator.

1. Create an identity that will be able to manage resources in the specified resource group.

    ```shell
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

2. Give the Service Principal associated with the AKS cluster control over managing the identity we just created.

    ```shell
    az role assignment create --role "Managed Identity Operator" --assignee <AKS Service Principal ID> --scope <Managed Identity ID Path>

    az role assignment create --role "Managed Identity Operator" --assignee <AKS Service Principal ID> --scope "/subscriptions/7060bca0-7a3c-44bd-b54c-4bb1e9facfac/resourcegroups/resourcegroup-operators/providers/Microsoft.ManagedIdentity/userAssignedIdentities/aso-manager-identity"
    ```

3. Give the Managed Identity we just created authorization to provision resources in our subscription

    ```shell
    az role assignment create --role "Contributor" --assignee <Managed Identity clientID> --scope <Subscription ID Path>

    az role assignment create --role "Contributor" --assignee "288a7d63-ab78-442e-89ee-2a353fb990ab"  --scope "/subscriptions/7060bca0-7a3c-44bd-b54c-4bb1e9facfac"
    ```

4. Install [aad-pod-identity](https://github.com/Azure/aad-pod-identity#1-create-the-deployment)

    ```shell
    make install-aad-pod-identity
    ```

5. Create and apply the AzureIdentity and Binding manifests

    ```yaml
    apiVersion: "aadpodidentity.k8s.io/v1"
    kind: AzureIdentity
    metadata:
        name: <a-idname>
    spec:
        type: 0
        ResourceID: /subscriptions/<subid>/resourcegroups/<resourcegroup>/providers/Microsoft.ManagedIdentity/userAssignedIdentities/<name>
        ClientID: <Managed Identity clientId>
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
