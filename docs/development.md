# Running the Azure Service Operator in development

## Running the operator locally

1. Clone the repository into the following folder `<GOPATH>/src/github.com/Azure`.

2. Make sure the environment variable `GO111MODULE` is set to `on`.

    ```bash
    export GO111MODULE=on
    ```

3. Install test certificates using `make generate-test-certs`.

4. Install the CRDs defined in the config/crd/bases folder using the command
    `make install`
    You will see output as below.

    ```shell
        kubectl apply -f config/crd/bases
        customresourcedefinition.apiextensions.k8s.io/eventhubnamespaces.azure.microsoft.com created
        customresourcedefinition.apiextensions.k8s.io/eventhubs.azure.microsoft.com created
        customresourcedefinition.apiextensions.k8s.io/resourcegroups.azure.microsoft.com created
        customresourcedefinition.apiextensions.k8s.io/sqldatabases.azure.microsoft.com created
        customresourcedefinition.apiextensions.k8s.io/sqlfirewallrules.azure.microsoft.com created
        customresourcedefinition.apiextensions.k8s.io/sqlservers.azure.microsoft.com configured
    ```

5. Run the operator locally using
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

6. Open a new terminal window. Trigger the creation of a custom resource using      kubectl and the sample YAML file provided.
   For instance, you would use the following command to create a SQL server:

   ```shell
   kubectl apply -f config/samples/azure_v1_sqlserver.yaml
   sqlserver.azure.microsoft.com/sqlserver-sample created
   ```

7. You should see logs on the other terminal from the operator when this custom     resource is being created.

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

## Developing using VSCode with Remote - Containers

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

6. Use `kubectl apply` with the sample YAML files to create custom resources for testing.
For eg., use  `kubectl apply -f config/samples/azure_v1_sqlserver.yaml` from the terminal to create a SQL server using the operator. 
`kubectl describe SqlServer` would show the events that indicate if the resource is created or being created.

## Development info

This repository is generated by [Kubebuilder](https://book.kubebuilder.io/).

To Extend the operator `github.com/Azure/azure-service-operator`:

1. Run `go mod download` to download dependencies. It doesn't show any progress bar and takes a while to download all of dependencies.
2. Update `api\v1\<service>_types.go`.
3. Regenerate CRD `make manifests`.
4. Install updated CRD `make install`
5. Generate code `make generate`
6. Update operator `controller\<service>_controller.go`
7. Update tests and run `make test`
8. Deploy `make deploy`

If you make changes to the operator and want to update the deployment without recreating the cluster (when testing locally), you can use the `make update` to update your Azure Operator pod. If you need to rebuild the docker image without cache, use `make ARGS="--no-cache" update`.

**Notes:**

- Don't forget to add `// +kubebuilder:subresource:status` if we want a status subresource. This is recommended as it makes updates to status more efficient (without having to update the whole instance)

- Run `make manifests` if you find the property you add doesn't work.

- Finalizers are recommended as they make deleting a resource more robust. Refer to [Using Finalizers](https://book.kubebuilder.io/reference/using-finalizers.html) for more information
