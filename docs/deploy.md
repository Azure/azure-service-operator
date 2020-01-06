# Building and deploying the Azure Service Operator

## Build the operator

1. Clone the repository.

2. Make sure the environment variable `GO111MODULE` is set to `on`.

    ```bash
    export GO111MODULE=on
    ```

3. Build the image and push it to docker hub.

    ```shell
    docker login
    IMG=<dockerhubusername>/<imagename>:<tag> make build-and-push
    ```

## Deploy the operator

**Note** You should already have a Kuberenetes cluster prerequisites [here](prereqs.md) for information on creating a Kubernetes cluster.

1. Set up the Cluster

    a. Create the namespace you want to deploy the operator to.

    **Note** The scripts currently are configured to deploy to the ```azureoperator-system``` namespace

    ```shell
    kubectl create namespace azureoperator-system
    ```

    b. Set the ```azureoperatorsettings``` secret.

    First, set the following environment variables `AZURE_TENANT_ID`, `AZURE_SUBSCRIPTION_ID`, `REQUEUE_AFTER`.

    ```shell
        export AZURE_TENANT_ID=xxxxxxx
        export AZURE_SUBSCRIPTION_ID=aaaaaaa
        export REQUEUE_AFTER=30
    ```

    c. **Authentication** If you would like to use a Service Principals (not recommended) for authentication, set these variables.

    ```shell
        export AZURE_CLIENT_ID=xxxxxxx
        export AZURE_CLIENT_SECRET=aaaaaaa
    ```

    If you would like to use Managed Identity (recommended) instead of service principals, set these variables instead.

    ```shell
        export AZURE_USE_MI=1
    ```

    Note: Use only one of the above.

    d. **Secrets storage** If you want to use Azure Key Vault to store the secrets like connection strings and SQL server username that result from the resource provisioning, you should also additionally do the steps below.
    If you do not perform these steps, the secrets will be stored as kube secrets by default.ß

    Create a Key Vault to use to store secrets

    ```shell
    az keyvault create --name "OperatorSecretKeyVault" --resource-group "resourceGroup-operators" --location "West US"
    ```

    Add appropriate Key Vault access policies to allow the service principal access to this Key Vault

    ```shell
    az keyvault set-policy --name "OperatorSecretKeyVault" --spn <AZURE_CLIENT_ID> --secret-permissions get list delete set
    ```

    If you use Managed Identity instead of Service Principal, use the Client ID of the Managed Identity instead in the above command.

    ```shell
    az keyvault set-policy --name "OperatorSecretKeyVault" --spn <MANAGEDIDENTITY_CLIENT_ID> --secret-permissions get list delete set
    ```

    Set the additional environment variable 'AZURE_OPERATOR_KEYVAULT' to indicate you want to use Key Vault for secrets.

    ```shell
    export AZURE_OPERATOR_KEYVAULT=OperatorSecretKeyVault
    ```

    e. From the same terminal, run the below command. **Note** The variables used here will depend on the environment variables you have set based on the authentication and secret storage choices made above. You will need to modify the below command accordingly.

    For instance, the below command assumes you have chosen to use Managed Identity for authentication and Key Vault for storing secrets.

    ```shell
    kubectl --namespace azureoperator-system \
        create secret generic azureoperatorsettings \
        --from-literal=AZURE_SUBSCRIPTION_ID="$AZURE_SUBSCRIPTION_ID" \
        --from-literal=AZURE_TENANT_ID="$AZURE_TENANT_ID" \
        --from-literal=AZURE_CLIENT_ID="$AZURE_CLIENT_ID" \
        --from-literal=AZURE_CLIENT_SECRET="$AZURE_CLIENT_SECRET" \
        --from-literal=AZURE_USE_MI="$AZURE_USE_MI" \
        --from-literal=AZURE_OPERATOR_KEYVAULT="$AZURE_OPERATOR_KEYVAULT" \
    ```

    f. Install [Cert Manager](https://docs.cert-manager.io/en/latest/getting-started/install/kubernetes.html)

    ```shell
    make install-cert-manager
    ```

2. Create an identity that will be able to manage resources

    ```shell
    # Create an identity to give to our operator-manager that will be used to authorize creation of resources in our
    # subscription. This could be restricted to a resource group by changing the scope on the "Contributor" role below
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

    ```shell
    # Give the AKS SP control over managing the identity we just created
    az role assignment create --role "Managed Identity Operator" --assignee <AKS Service Principal ID> --scope <Managed Identity ID Path>
    ```

    In the above case, this will look like below:

    ```shell
    az role assignment create --role "Managed Identity Operator" --assignee <AKS Service Principal ID> --scope "/subscriptions/7060bca0-7a3c-44bd-b54c-4bb1e9facfac/resourcegroups/resourcegroup-operators/providers/Microsoft.ManagedIdentity/userAssignedIdentities/aso-manager-identity"
    ```

    ```shell
    # Give our aso-manager-identity authorization to provision resources in our subscription
    az role assignment create --role "Contributor" --assignee <Managed Identity Principal ID> --scope <Subscription ID Path>
    ```

    In the above case, this will look like below:

    ```shell
    az role assignment create --role "Contributor" --assignee "288a7d63-ab78-442e-89ee-2a353fb990ab"  --scope "/subscriptions/7060bca0-7a3c-44bd-b54c-4bb1e9facfac"
    ```

    h. Install [aad-pod-identity](https://github.com/Azure/aad-pod-identity#1-create-the-deployment)

    ```shell
    make install-aad-pod-identity
    ```

3. Create and apply the AzureIdentity and Binding manifests

    ```yaml
    apiVersion: "aadpodidentity.k8s.io/v1"
    kind: AzureIdentity
    metadata:
      name: <a-idname>
    spec:
      type: 0
      ResourceID: /subscriptions/<subid>/resourcegroups/<resourcegroup>/providers/Microsoft.ManagedIdentity/userAssignedIdentities/<name>
      ClientID: <clientId>
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

4. Deploy the operator to the Kubernetes cluster

    ```shell
    make deploy
    ```

5. Check that the operator is deployed to the cluster using the following commands.

    ```shell
    kubectl get pods -n azureoperator-system
    ```

6. You can view the logs from the operator using the following command. The `podname` is the name of the pod in the output from `kubectl get pods -n azureoperator-system`, `manager` is the name of the container inside the pod.

    ```shell
    kubectl logs <podname> -c manager -n azureoperator-system
    ```

7. If you would like to view the Prometheus metrics from the operator, you can redirect port 8080 to the local machine using the following command:

   Get the deployment using the following command

   ```shell
   kubectl get deployment -n azureoperator-system
   ```

   You'll see output like the below.

   ```shell
   NAME                               READY   UP-TO-DATE   AVAILABLE   AGE
   azureoperator-controller-manager   1/1     1            1           2d1h
   ```

   Use the deployment name in the command as below

    ```shell
    kubectl port-forward deployment/<deployment name> -n <namespace> 8080
    ```

    So we would use the following command here

    ```shell
    kubectl port-forward deployment/azureoperator-controller-manager -n azureoperator-system 8080
    ```

    You can now browse to `http://localhost:8080/metrics` from the browser to view the metrics.