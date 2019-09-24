# Azure Service Operator (for Kubernetes)

[![Build Status](https://dev.azure.com/azure/azure-service-operator/_apis/build/status/Azure.azure-service-operator?branchName=master)](https://dev.azure.com/azure/azure-service-operator/_build/latest?definitionId=36&branchName=master)

> This project is experimental. Expect the API to change. It is not recommended for production environments.

## Introduction

Kubernetes offers the facility of extending it's API through the concept of 'Operators' ([Introducing Operators: Putting Operational Knowledge into Software](https://coreos.com/blog/introducing-operators.html)). 

An Operator is an application-specific controller that extends the Kubernetes API to create, configure, and manage instances of complex stateful applications on behalf of a Kubernetes user. It builds upon the basic Kubernetes resource and controller concepts but includes domain or application-specific knowledge to automate common tasks.

This repository contains the resources and code to provision and deprovision different Azure services using a Kubernetes operator.

The Azure Operator comprises of:

- The Custom Resource Definitions (CRDs) for each of the Azure services that the Kubernetes user can provision
- The Kubernetes controller that watches for requests to create Custom Resources for these CRDs and creates them

The project was built using

[Kubebuilder](https://book.kubebuilder.io/)

## Azure Services supported
1. Resource Group 
2. EventHub
3. Azure SQL

## Building and Running from Source

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
6. [Kustomize](https://github.com/kubernetes-sigs/kustomize) is also required. This must be installed via `make install-kustomize` (see section below).

Basic commands to check if you have an active Kubernetes cluster:

```shell
    kubectl config get-contexts
    kubectl cluster-info
    kubectl version
    kubectl get pods -n kube-system
```

### Quick start - Using VSCode with Remote-Containers extension 

If you're using VSCode with [Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) extensions installed, you can quickly have your environment set up and ready to go, with everything you need to get started.

1. Open this project in VSCode.
2. Inside `.devcontainer`, create a file called `.env` and using the following template, copy your Service Principal's details.

    ```txt
    AZURE_CLIENT_ID=

    AZURE_CLIENT_SECRET=

    AZURE_SUBSCRIPTION_ID=

    AZURE_TENANT_ID=
    ```

3. Open the Command Pallet (`Command+Shift+P` on MacOS or `CTRL+Shift+P` on Windows), type `Remote-Containers: Open Folder in Container...` and hit enter.
4. VSCode will relaunch and start building our development container. This will install all the necessary dependencies required for you to begin developing.
5. Once the container has finished building, you can now start testing your Azure Service Operator within your own local kubernetes environment.

**Note**: after the DevContainer has finished building, the kind cluster will start initialising and installing the Azure Service Operator in the background. This will take some time before it is available.

To see when the kind cluster is ready, use `docker ps -a` to list your running containers, look for `IMAGE` with the name `azure-service-operator_devcontainer_docker-in-docker...`. Using that image's `CONTAINER ID`, use `docker logs -f CONTAINER ID` to view the logs from the container setting up your cluster.


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

    Then make sure this service principal has rights assigned to provision resources in your Azure Subscription.
  
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

   You will see something like this on the terminal windows indicating that the controller is running.

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
   ```bash
   kubectl apply -f config/samples/azure_v1_sqlserver.yaml
   sqlserver.azure.microsoft.com/sqlserver-sample created
   ```

10. You should see logs on the other terminal from the operator when this custom     resource is being created.
    ``` bash
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
    2019-09-24T12:28:20.331-0600	DEBUG	controller-runtime.manager.events	Normal	{"object": {"kind":"SqlServer","namespace":"default","name":"sqlserver-sample1","uid":"ed3774af-def8-11e9-90c4-025000000001","apiVersion":"azure.microsoft.com/v1","resourceVersion":"194518"}, "reason": "Provisioned", "message": "sqlserver sqlserver-sample1 provisioned "}
```

### Deploying the operator on a Kubernetes cluster

3. Get your Kubernetes cluster setup:
   
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


6. Set up the Cluster

   If you are using Kind:

    ```shell
    make set-kindcluster
    ```

    If you are not using Kind, it's a manual process, as follows:

    a. Create the namespace

    ```shell
    kubectl create namespace azureoperator-system
    ```

    b. Set the azureoperatorsettings secret

    ```shell
    kubectl --namespace azureoperator-system \
        create secret generic azureoperatorsettings \
        --from-literal=AZURE_CLIENT_ID="$AZURE_CLIENT_ID" \
        --from-literal=AZURE_CLIENT_SECRET="$AZURE_CLIENT_SECRET" \
        --from-literal=AZURE_SUBSCRIPTION_ID="$AZURE_SUBSCRIPTION_ID" \
        --from-literal=AZURE_TENANT_ID="$AZURE_TENANT_ID"
    ```

    c. [Cert Manager](https://docs.cert-manager.io/en/latest/getting-started/install/kubernetes.html)

    ```shell
    kubectl get secret webhook-server-cert -n azureoperator-system -o yaml > certs.txt
    ```

    you can use `https://inbrowser.tools/` and extract `ca.crt`, `tls.crt` and `tls.key`



8. Install the azure_v1_eventhub CRD in the configured Kubernetes cluster folder ~/.kube/config,

    run `kubectl apply -f config/crd/bases` or `make install`

## Add support for a new Azure service

TODO: Fill out

## How to extend the operator and build your own images

### Updating the Azure operator

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

## Contributing

This project welcomes contributions and suggestions.  Most contributions require you to agree to a
Contributor License Agreement (CLA) declaring that you have the right to, and actually do, grant us
the rights to use your contribution. For details, visit https://cla.opensource.microsoft.com.

When you submit a pull request, a CLA bot will automatically determine whether you need to provide
a CLA and decorate the PR appropriately (e.g., status check, comment). Simply follow the instructions
provided by the bot. You will only need to do this once across all repos using our CLA.

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or
contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.
