# Building and testing the Operator

## Running the operator

### Prerequisites And Assumptions

1. [GoLang](https://golang.org/dl/) is installed.
2. [Docker](https://docs.docker.com/install/) is installed and running.
3. [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) Command-line tool is installed.
4. You have access to a Kubernetes cluster.
    - It can be a local hosted Cluster like
    [Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/),
    [Kind](https://github.com/kubernetes-sigs/kind), or,
    [Docker for desktop](https://blog.docker.com/2018/07/kubernetes-is-now-available-in-docker-desktop-stable-channel/) installed locally with RBAC enabled.
    - If you opt for Azure Kubernetes Service ([AKS](https://azure.microsoft.com/en-au/services/kubernetes-service/)), you can set the current context to your cluster using the following command:

    `az aks get-credentials --resource-group $RESOURCEGROUP_NAME --name $CLUSTER_NAME`

5. Install [Kubebuilder](https://book.kubebuilder.io/), following the linked installation instructions.

If you are running on MacOS, look at the notes [here](/docs/kubebuilder.md) for issues you may encounter.

6. [Kustomize](https://github.com/kubernetes-sigs/kustomize) is also required. This must be installed via `make install-kustomize` (see section below).

Basic commands to check if you have an active Kubernetes cluster:

```shell
    kubectl config get-contexts
    kubectl cluster-info
    kubectl version
    kubectl get pods -n kube-system
```

### Running the operator locally for testing

1. Clone the repository into the following folder `<GOPATH>/src/github.com/Azure`.

2. Make sure the environment variable `GO111MODULE` is set to `on`.
    ```bash
    export GO111MODULE=on
    ```

3. Create a Service Principal
    If you don't have a Service Principal create one from Azure CLI:

    ```bash
    az ad sp create-for-rbac --role Contributor
    ```

    Then make sure this service principal has rights assigned to provision resources in your Azure Subscription and resource group.
  
4. Set the environment variables `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, `AZURE_CLIENT_SECRET`, `AZURE_SUBSCRIPTION_ID`, `REQUEUE_AFTER`.

    ```shell
        export AZURE_TENANT_ID=xxxxxxx
        export AZURE_CLIENT_ID=yyyyyyy
        export AZURE_CLIENT_SECRET=zzzzzz
        export AZURE_SUBSCRIPTION_ID=aaaaaaa
        export REQUEUE_AFTER=30
    ```

    If running on Windows, the environment variables should not have quotes and should be set like below.

    ```shell
        set AZURE_TENANT_ID=xxxxxxx
        set AZURE_CLIENT_ID=yyyyyyy
        set AZURE_CLIENT_SECRET=zzzzzz
        set AZURE_SUBSCRIPTION_ID=aaaaaaa
        set REQUEUE_AFTER=30
    ```
    VSCode or `make run` should be run from the same session/command/terminal window where the environment variables are set.

5. Install [kustomize](https://github.com/kubernetes-sigs/kustomize) using `make install-kustomize`.

6. Install test certificates using `make generate-test-certs`.

7. Install the CRDs defined in the config/crd/bases folder using the command
    ```make install```
    You will see output as below

    ```shell
        kubectl apply -f config/crd/bases
        customresourcedefinition.apiextensions.k8s.io/eventhubnamespaces.azure.microsoft.com created
        customresourcedefinition.apiextensions.k8s.io/eventhubs.azure.microsoft.com created
        customresourcedefinition.apiextensions.k8s.io/resourcegroups.azure.microsoft.com created
        customresourcedefinition.apiextensions.k8s.io/sqldatabases.azure.microsoft.com created
        customresourcedefinition.apiextensions.k8s.io/sqlfirewallrules.azure.microsoft.com created
        customresourcedefinition.apiextensions.k8s.io/sqlservers.azure.microsoft.com configured
    ```

8. Run the operator locally using 
   ```make run```
   This will cause the operator to run and "watch" for events on this terminal. You will need to open a new terminal to trigger the creation of a custom resource.

   You will see something like this on the terminal window indicating that the controller is running.

   ```shell
    go fmt ./...
    go vet ./...
    go run ./main.go
    2019-09-24T12:18:10.419-0600	INFO	controller-runtime.metrics	metrics server is starting to listen	{"addr": ":8080"}

    2019-09-24T12:18:10.419-0600	INFO	controller-runtime.controller	Starting EventSource	{"controller": "eventhub", "source": "kind source: /, Kind="}
    2019-09-24T12:18:10.419-0600	INFO	controller-runtime.controller	Starting EventSource	{"controller": "resourcegroup", "source": "kind source: /, Kind="}
    2019-09-24T12:18:10.419-0600	INFO	controller-runtime.controller	Starting EventSource	{"controller": "eventhubnamespace", "source": "kind source: /, Kind="}
    2019-09-24T12:18:10.420-0600	INFO	controller-runtime.builder	Registering a mutating webhook	{"GVK": "azure.microsoft.com/v1, Kind=EventhubNamespace", "path": "/mutate-azure-microsoft-com-v1-eventhubnamespace"}
    2019-09-24T12:18:10.420-0600	INFO	controller-runtime.builder	Registering a validating webhook	{"GVK": "azure.microsoft.com/v1, Kind=EventhubNamespace", "path": "/validate-azure-microsoft-com-v1-eventhubnamespace"}
    2019-09-24T12:18:10.420-0600	INFO	controller-runtime.builder	Registering a mutating webhook	{"GVK": "azure.microsoft.com/v1, Kind=Eventhub", "path": "/mutate-azure-microsoft-com-v1-eventhub"}
    2019-09-24T12:18:10.424-0600	INFO	controller-runtime.builder	Registering a validating webhook	{"GVK": "azure.microsoft.com/v1, Kind=Eventhub", "path": "/validate-azure-microsoft-com-v1-eventhub"}
    2019-09-24T12:18:10.424-0600	INFO	controller-runtime.controller	Starting EventSource	{"controller": "sqlserver", "source": "kind source: /, Kind="}
    2019-09-24T12:18:10.424-0600	INFO	controller-runtime.controller	Starting EventSource	{"controller": "sqldatabase", "source": "kind source: /, Kind="}
    2019-09-24T12:18:10.424-0600	INFO	controller-runtime.controller	Starting EventSource	{"controller": "sqlfirewallrule", "source": "kind source: /, Kind="}
    2019-09-24T12:18:10.424-0600	INFO	setup	starting manager
    2019-09-24T12:18:10.424-0600	INFO	controller-runtime.manager	starting metrics server	{"path": "/metrics"}
    2019-09-24T12:18:10.526-0600	INFO	controller-runtime.controller	Starting Controller	{"controller": "eventhub"}
    2019-09-24T12:18:10.526-0600	INFO	controller-runtime.controller	Starting Controller	{"controller": "resourcegroup"}
    2019-09-24T12:18:10.526-0600	INFO	controller-runtime.controller	Starting Controller	{"controller": "sqldatabase"}
    2019-09-24T12:18:10.526-0600	INFO	controller-runtime.controller	Starting Controller	{"controller": "sqlfirewallrule"}
    2019-09-24T12:18:10.526-0600	INFO	controller-runtime.controller	Starting Controller	{"controller": "eventhubnamespace"}
    2019-09-24T12:18:10.526-0600	INFO	controller-runtime.controller	Starting Controller	{"controller": "sqlserver"}
    2019-09-24T12:18:10.527-0600	INFO	controller-runtime.certwatcher	Updated current TLS certiface
    2019-09-24T12:18:10.527-0600	INFO	controller-runtime.certwatcher	Starting certificate watcher
    2019-09-24T12:18:10.626-0600	INFO	controller-runtime.controller	Starting workers	{"controller": "sqldatabase", "worker count": 1}
    2019-09-24T12:18:10.626-0600	INFO	controller-runtime.controller	Starting workers	{"controller": "resourcegroup", "worker count": 1}
    2019-09-24T12:18:10.626-0600	INFO	controller-runtime.controller	Starting workers	{"controller": "sqlfirewallrule", "worker count": 1}
    2019-09-24T12:18:10.626-0600	INFO	controller-runtime.controller	Starting workers	{"controller": "eventhub", "worker count": 1}
    2019-09-24T12:18:10.626-0600	INFO	controller-runtime.controller	Starting workers	{"controller": "sqlserver", "worker count": 1}
    2019-09-24T12:18:10.626-0600	INFO	controller-runtime.controller	Starting workers	{"controller": "eventhubnamespace", "worker count": 1}
   ```

9. Open a new terminal window. Trigger the creation of a custom resource using      kubectl and the sample YAML file provided.
   For instance, you would use the following command to create a SQL server: 
   ```shell
   kubectl apply -f config/samples/azure_v1_sqlserver.yaml
   sqlserver.azure.microsoft.com/sqlserver-sample created
   ```

10. You should see logs on the other terminal from the operator when this custom     resource is being created.
    ``` shell
    2019-09-24T12:27:12.450-0600	DEBUG	controller-runtime.manager.events	Normal	{"object": {"kind":"SqlServer","namespace":"default","name":"sqlserver-sample1","uid":"ed3774af-def8-11e9-90c4-025000000001","apiVersion":"azure.microsoft.com/v1","resourceVersion":"194427"}, "reason": "Updated", "message": "finalizer sqlserver.finalizers.azure.com added"}
    2019-09-24T12:27:12.450-0600	DEBUG	controller-runtime.manager.events	Normal	{"object": {"kind":"SqlServer","namespace":"default","name":"sqlserver-sample1","uid":"ed3774af-def8-11e9-90c4-025000000001","apiVersion":"azure.microsoft.com/v1","resourceVersion":"194427"}, "reason": "Submitting", "message": "starting resource reconciliation"}
    2019-09-24T12:27:17.129-0600	INFO	controllers.SqlServer	mutating secret bundle
    2019-09-24T12:27:17.144-0600	INFO	controllers.SqlServer	waiting for provision to take effect	{"sqlserver": "default/sqlserver-sample1"}
    2019-09-24T12:27:17.350-0600	DEBUG	controller-runtime.manager.events	Normal	{"object": {"kind":"SqlServer","namespace":"default","name":"sqlserver-sample1","uid":"ed3774af-def8-11e9-90c4-025000000001","apiVersion":"azure.microsoft.com/v1","resourceVersion":"194437"}, "reason": "Checking", "message": "instance in NotReady state"}
    2019-09-24T12:27:17.359-0600	INFO	controllers.SqlServer	Got ignorable error	{"sqlserver": "default/sqlserver-sample1", "type": "ResourceNotFound"}
    2019-09-24T12:27:17.564-0600	DEBUG	controller-runtime.manager.events	Normal	{"object": {"kind":"SqlServer","namespace":"default","name":"sqlserver-sample1","uid":"ed3774af-def8-11e9-90c4-025000000001","apiVersion":"azure.microsoft.com/v1","resourceVersion":"194439"}, "reason": "Checking", "message": "instance in NotReady state"}
    2019-09-24T12:27:17.570-0600	INFO	controllers.SqlServer	Got ignorable error	{"sqlserver": "default/sqlserver-sample1", "type": "ResourceNotFound"}
    2019-09-24T12:28:17.805-0600	DEBUG	controller-runtime.manager.events	Normal	{"object": {"kind":"SqlServer","namespace":"default","name":"sqlserver-sample1","uid":"ed3774af-def8-11e9-90c4-025000000001","apiVersion":"azure.microsoft.com/v1","resourceVersion":"194439"}, "reason": "Checking", "message": "instance in Ready state"}
    2019-09-24T12:28:19.010-0600	DEBUG	controller-runtime.controller	Successfully Reconciled	{"controller": "sqlserver", "request": "default/sqlserver-sample1"}
    2019-09-24T12:28:19.010-0600	DEBUG	controller-runtime.manager.events	Normal	{"object": {"kind":"SqlServer","namespace":"default","name":"sqlserver-sample1","uid":"ed3774af-def8-11e9-90c4-025000000001","apiVersion":"azure.microsoft.com/v1","resourceVersion":"194518"}, "reason": "Provisioned", "message": "sqlserver sqlserver-sample1 provisioned "}
    2019-09-24T12:28:19.202-0600	DEBUG	controller-runtime.manager.events	Normal	{"object": {"kind":"SqlServer","namespace":"default","name":"sqlserver-sample1","uid":"ed3774af-def8-11e9-90c4-025000000001","apiVersion":"azure.microsoft.com/v1","resourceVersion":"194518"}, "reason": "Checking", "message": "instance in Ready state"}
    2019-09-24T12:28:20.331-0600	DEBUG	controller-runtime.controller	Successfully Reconciled	{"controller": "sqlserver", "request": "default/sqlserver-sample1"}
    2019-09-24T12:28:20.331-0600	DEBUG	controller-runtime.manager.events	Normal	{"object": {"kind":"SqlServer","namespace":"default","name":"sqlserver-sample1","uid":"ed3774af-def8-11e9-90c4-025000000001","apiVersion":"azure.microsoft.com/v1","resourceVersion":"194518"}, "reason": "Provisioned", "message": "sqlserver sqlserver-sample1 provisioned "}```






If you're using VSCode with [Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) extensions installed, you can quickly have your environment set up and ready to go, with everything you need to get started.

1. Open this project in VSCode.
2. Inside the folder `.devcontainer`, create a file called `.env` and using the following template, copy your environment variable details.

    ```txt
    AZURE_CLIENT_ID=

    AZURE_CLIENT_SECRET=

    AZURE_SUBSCRIPTION_ID=

    AZURE_TENANT_ID=
    ```

3. Open the Command Pallet (`Command+Shift+P` on MacOS or `CTRL+Shift+P` on Windows), type `Remote-Containers: Open Folder in Container...`, select the ```Azure-service-operator``` folder and hit enter.

4. VSCode will relaunch and start building our development container. This will install all the necessary dependencies required for you to begin developing.

5. Once the container has finished building, you can now start testing your Azure Service Operator within your own local kubernetes environment via the terminal inside VSCode.

**Note**: after the DevContainer has finished building, the kind cluster will start initialising and installing the Azure Service Operator in the background. This will take some time before it is available.

To see when the kind cluster is ready, use `docker ps -a` to list your running containers, look for `IMAGE` with the name `azure-service-operator_devcontainer_docker-in-docker...`. Using that image's `CONTAINER ID`, use `docker logs -f CONTAINER ID` to view the logs from the container setting up your cluster.

6. Use ```kubectl apply``` with the sample YAML files to create custom resources for testing. 
For eg., use ```kubectl apply -f config/samples/azure_v1_sqlserver.yaml``` from the terminal to create a SQL server using the operator. 
```kubectl describe SqlServer``` would show the events that indicate if the resource is created or being created.

### Deploying the operator on a Kubernetes cluster

1. Create your Kubernetes cluster
   (i) If you use Kind, install [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

    ```shell
        GO111MODULE="on" go get sigs.k8s.io/kind@v0.4.0 && kind create cluster
        kind create cluster
        export KUBECONFIG="$(kind get kubeconfig-path --name="kind")"
        kubectl cluster-info
        IMG="docker.io/yourimage:tag" make docker-build
        kind load docker-image docker.io/yourimage:tag --loglevel "trace"
        make deploy
    ```

    (ii) If you use Docker for Desktop, enable [Kubernetes](https://docs.docker.com/docker-for-mac/#kubernetes). For Windows users, you can follow [these](https://blog.docker.com/2018/01/docker-windows-desktop-now-kubernetes/) steps

    (iii) If you use Azure Kubernetes Service, create a cluster from the Azure portal or using [Azure CLI](https://docs.microsoft.com/en-us/azure/aks/kubernetes-walkthrough#create-aks-cluster)

    ```shell
    az aks create -g <ResourceGroupName> -n <AKSClusterName>
    ```

    You could also create the AKS cluster from the Azure portal. Once you create the  cluster, use the following command to set it as the current cluster.
   
    ```shell
    az aks get-credentials -g <ResourceGroupName> -n <AKSClusterName>
    ```

2. Set up the Cluster

   If you are using Kind:

    ```shell
    make set-kindcluster
    ```

    If you are not using Kind, it's a manual process, as follows:

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

    If running on Windows, the environment variables should not have quotes and should be set like below.

    ```shell
        set AZURE_TENANT_ID=xxxxxxx
        set AZURE_CLIENT_ID=yyyyyyy
        set AZURE_CLIENT_SECRET=zzzzzz
        set AZURE_SUBSCRIPTION_ID=aaaaaaa
        set REQUEUE_AFTER=30
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

8. Build the image and push it to docker hub.

    ```shell
    docker login
    IMG=<dockerhubusername>/<imagename>:<tag> make docker-build
    IMG=<dockerhubusername>/<imagename>:<tag> make docker-push
    ```

9. Give the default service account on the azureoperator-system namespace "cluster-admin" access to the "default" namespace where the operator deploys the resources by default.

**Note** If you deploy resources to a different namespace than default then you will need to provide permission to that namespace too.

    ```shell
    kubectl create clusterrolebinding default-admin-binding --clusterrole=cluster-admin --user=system:serviceaccount:azureoperator-system:default --namespace=default
    ```
10. Install the Custom Resource Definitions (CRDs) in the configured Kubernetes cluster

    Use  `make install`

11. Deploy the operator to the Kubernetes cluster

    ```shell
    IMG=<dockerhubusername>/<imagename>:<tag> make deploy
    ```

12. Check that the operator is deployed to the cluster using the following commands.
    ```shell
    kubectl get namespaces
    kubectl get pods -n azureoperator-system
    ```

13. You can view the logs from the operator using the following command. The `podname` is the name of the pod in the output from `kubectl get pods -n azureoperator-system`, `manager` is the name of the container inside the pod.

    ```shell
    kubectl logs <podname> -c manager -n azureoperator-system
    ```



This repository is generated by [Kubebuilder](https://book.kubebuilder.io/).

To Extend the operator `github.com/Azure/azure-service-operator`:

1. Run `go mod download` to download dependencies. It doesn't show any progress bar and takes a while to download all of dependencies.
2. Update `api\v1\eventhub_types.go`.
3. Regenerate CRD `make manifests`.
4. Install updated CRD `make install`
5. Generate code `make generate`
6. Update operator `controller\eventhub_controller.go`
7. Update tests and run `make test`
8. Deploy `make deploy`

If you make changes to the operator and want to update the deployment without recreating the cluster (when testing locally), you can use the `make update` to update your Azure Operator pod. If you need to rebuild the docker image without cache, use `make ARGS="--no-cache" update`.

## Developing

### Using VSCode with Remote-Containers extension

If you're using VSCode with [Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) extensions installed, you can quickly have your environment set up and ready to go, with everything you need to get started.

1. Open this project in VSCode.
2. Inside the folder `.devcontainer`, create a file called `.env` and using the following template, copy your environment variable details.

    ```txt
    AZURE_CLIENT_ID=

    AZURE_CLIENT_SECRET=

    AZURE_SUBSCRIPTION_ID=

    AZURE_TENANT_ID=
    ```

3. Open the Command Pallet (`Command+Shift+P` on MacOS or `CTRL+Shift+P` on Windows), type `Remote-Containers: Open Folder in Container...`, select the ```Azure-service-operator``` folder and hit enter.

4. VSCode will relaunch and start building our development container. This will install all the necessary dependencies required for you to begin developing.

5. Once the container has finished building, you can now start testing your Azure Service Operator within your own local kubernetes environment via the terminal inside VSCode.

**Note**: after the DevContainer has finished building, the kind cluster will start initialising and installing the Azure Service Operator in the background. This will take some time before it is available.

To see when the kind cluster is ready, use `docker ps -a` to list your running containers, look for `IMAGE` with the name `azure-service-operator_devcontainer_docker-in-docker...`. Using that image's `CONTAINER ID`, use `docker logs -f CONTAINER ID` to view the logs from the container setting up your cluster.

6. Use ```kubectl apply``` with the sample YAML files to create custom resources for testing. 
For eg., use ```kubectl apply -f config/samples/azure_v1_sqlserver.yaml``` from the terminal to create a SQL server using the operator. 
```kubectl describe SqlServer``` would show the events that indicate if the resource is created or being created.

## Add a New Custom Resource

If you want to create a new custom resource for a new Azure service, you will need to follow the following steps:

1. Add a New API using the following Kubebuilder command.

```shell
kubebuilder create api --group service --version v1alpha1 --kind <Azure-Service-Name>
```

For instance, this is how I would create a new API/scaffold for a Redis operator

```shell
kubebuilder create api --group azure --version v1 --kind RedisCache
```

Refer to [kubebuilder's doc](https://book.kubebuilder.io/cronjob-tutorial/new-api.html)

2. Define the Spec and the Status for the new `Kind` in the appropriate types.go file under `api\v1`

3. Implement the Reconcile logic for the controller in the appropriate controller file under `controllers` directory.


**Notes:**

- Don't forget to add `// +kubebuilder:subresource:status` if we want a status subresource. This is recommended as it makes updates to status more efficient (without having to update the whole instance)

- Run `make manifests` if you find the property you add doesn't work.

- Finalizers are recommended as they make deleting a resource more robust. Refer to [Using Finalizers](https://book.kubebuilder.io/reference/using-finalizers.html) for more information