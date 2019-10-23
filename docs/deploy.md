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

    First, set the following environment variables `AZURE_TENANT_ID`, `AZURE_SUBSCRIPTION_ID`, `AZURE_USE_MSI`, `REQUEUE_AFTER`.

    ```shell
        export AZURE_TENANT_ID=xxxxxxx
        export AZURE_SUBSCRIPTION_ID=aaaaaaa
        export AZURE_USE_MSI=1
        export REQUEUE_AFTER=30
    ```

    From the same terminal, run the below command.

    ```shell
    kubectl --namespace azureoperator-system \
        create secret generic azureoperatorsettings \
        --from-literal=AZURE_SUBSCRIPTION_ID="$AZURE_SUBSCRIPTION_ID" \
        --from-literal=AZURE_TENANT_ID="$AZURE_TENANT_ID" \
        --from-literal=AZURE_USE_MSI="$AZURE_USE_MSI"
    ```

    c. Install [Cert Manager](https://docs.cert-manager.io/en/latest/getting-started/install/kubernetes.html)

    ```shell
    make install-cert-manager
    ```

    e. Create an identity that will be able to manage resources

    ```shell
    # Create an identity to give to our operator-manager that will be used to authorize creation of resources in our
    # subscription. This could be restricted to a resource group by changing the scope on the "Contributor" role below
    az identity create -g <resourcegroup> -n aso-manager-identity -o json

    # Give the AKS SP control over managing the identity we just created
    az role assignment create --role "Managed Identity Operator" --assignee <AKS Service Principal ID> --scope <Managed Identity ID Path>

    # Give our aso-manager-identity authorization to provision resources in our subscription
    az role assignment create --role "Contributor" --assignee <Managed Identity Principal ID> --scope <Subscription ID Path>
    ```

    d. Install [aad-pod-identity](https://github.com/Azure/aad-pod-identity#1-create-the-deployment)

    ```shell
    make install-aad-pod-identity
    ```

3. Give the default service account on the azureoperator-system namespace "cluster-admin" access to the namespace where the operator deploys the resources by default.

    **Note** Not recommended for production.

    ```shell
    kubectl create clusterrolebinding cluster-admin-aso \
      --clusterrole=cluster-admin \
      --user=system:serviceaccount:azureoperator-system:default
    ```

4. Create and apply the AzureIdentity and Binding manifests

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
      Selector: cluster_identity_binding
    ```

5. Deploy the operator to the Kubernetes cluster

    ```shell
    make deploy
    ```

6. Check that the operator is deployed to the cluster using the following commands.

    ```shell
    kubectl get pods -n azureoperator-system
    ```

7. You can view the logs from the operator using the following command. The `podname` is the name of the pod in the output from `kubectl get pods -n azureoperator-system`, `manager` is the name of the container inside the pod.

    ```shell
    kubectl logs <podname> -c manager -n azureoperator-system
    ```
