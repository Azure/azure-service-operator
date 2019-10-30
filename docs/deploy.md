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

    First, set the following environment variables `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, `AZURE_CLIENT_SECRET`, `AZURE_SUBSCRIPTION_ID`, `REQUEUE_AFTER`.

    ```shell
        export AZURE_TENANT_ID=xxxxxxx
        export AZURE_CLIENT_ID=yyyyyyy
        export AZURE_CLIENT_SECRET=zzzzzz
        export AZURE_SUBSCRIPTION_ID=aaaaaaa
        export REQUEUE_AFTER=30
    ```

    From the same terminal, run the below command.

    ```shell
    kubectl --namespace azureoperator-system \
        create secret generic azureoperatorsettings \
        --from-literal=AZURE_CLIENT_ID="$AZURE_CLIENT_ID" \
        --from-literal=AZURE_CLIENT_SECRET="$AZURE_CLIENT_SECRET" \
        --from-literal=AZURE_SUBSCRIPTION_ID="$AZURE_SUBSCRIPTION_ID" \
        --from-literal=AZURE_TENANT_ID="$AZURE_TENANT_ID"
    ```

    c. Install [Cert Manager](https://docs.cert-manager.io/en/latest/getting-started/install/kubernetes.html)

    ```shell
    make install-cert-manager
    ```

3. Give the default service account on the azureoperator-system namespace "cluster-admin" access to the namespace where the operator deploys the resources by default.

    **Note** Not recommended for production.

    ```shell
    kubectl create clusterrolebinding cluster-admin-aso \
      --clusterrole=cluster-admin \
      --user=system:serviceaccount:azureoperator-system:default
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

7. In order to view the Prometheus metrics, you can redirect port 8080 to the local machine using the following command:

    ```shell
    kubectl port-forward deployment/controller-manager 8080
    ```

    Then, open a web browser and navigate to the [Metrics Endpoint](http://127.0.0.1:8080/metrics).
