# Deploying the Azure Service Operator

The easiest way to deploy the Azure Service Operator is to use Helm charts. Follow instructions [here](helmdeploy.md).
However, if you're interested and building and deploying the operators from source, follow the instructions laid out here.

## Build the operator

1. Clone the repository.

2. Make sure the environment variable `GO111MODULE` is set to `on`.

    ```bash
    export GO111MODULE=on
    ```

3. Log in to your container registry, set the IMG env var, and then build and push the image.

    ```shell
    docker login
    IMG=<container_registry>/<username>/<image_name>:<tag> make build-and-push
    ```

## Deploy the operator

> [!NOTE]
> This step assumes you already have a Kubernetes cluster. See prerequisites [here](prereqs.md) for information on creating a Kubernetes cluster.

### Set up the cluster

Create the namespace you want to deploy the operator to.

> [!NOTE]
> The scripts currently are configured to deploy to the ```azureoperator-system``` namespace

```shell
kubectl create namespace azureoperator-system
```

Next, install [Cert Manager](https://docs.cert-manager.io/en/latest/getting-started/install/kubernetes.html).

```shell
make install-cert-manager
```

### Secrets storage
   
You have the option to use either of the below for storing secrets like connection strings and SQL server username that result from the resource provisioning.

*Kubernetes secrets* -
    this is the default. Secrets will be stored as Kubernetes secrets by default.

*Azure Key Vault* -
    iff you want to use Azure Key Vault to store the secrets, you should also additionally do the steps below.

#### Set up Azure Key Vault for Secret Storage (optional)
Create an Azure Key Vault to use to store secrets.

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

Set the environment variable 'AZURE_OPERATOR_KEYVAULT' to indicate you want to use Azure Key Vault for secrets.

```shell
export AZURE_OPERATOR_KEYVAULT=OperatorSecretKeyVault
```

### Authentication

You can choose to use either Service Principals or Managed Identity for authentication.

*Service Principal authentication* - if you choose to use Service Principal authentication, set these environment variables.

```shell
export AZURE_CLIENT_ID=xxxxxxx
export AZURE_CLIENT_SECRET=aaaaaaa
```

*Managed Identity authentication* - if you choose to use Managed Identity, set the below environment variable and then perform the steps listed [here](managedidentity.md).

```shell
export AZURE_USE_MI=1
```

> [!NOTE]
> Use only one of the above.


Set the ```azureoperatorsettings``` secret, and set the following environment variables `AZURE_TENANT_ID`, `AZURE_SUBSCRIPTION_ID`, `REQUEUE_AFTER`.

```shell
export AZURE_TENANT_ID=xxxxxxx
export AZURE_SUBSCRIPTION_ID=aaaaaaa
export REQUEUE_AFTER=30
```

From the same terminal, run the below command.

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

### Deploy the operator to the Kubernetes cluster

```shell
make deploy
```

Check that the operator is deployed to the cluster using the following commands.

```shell
kubectl get pods -n azureoperator-system
```

You can view the logs from the operator using the following command. The `podname` is the name of the pod in the output from `kubectl get pods -n azureoperator-system`, `manager` is the name of the container inside the pod.

```shell
kubectl logs <podname> -c manager -n azureoperator-system
```
### Prometheus metrics
If you would like to view the Prometheus metrics from the operator, you can redirect port 8080 to the local machine using the following commands:

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
