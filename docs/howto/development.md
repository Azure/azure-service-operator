# Building the Azure Service Operator from source

## Running the operator locally

> [!NOTE]
> Using this method you can only install resources having the the latest version of a CRD, as there are no conversion webhooks running to convert other versions. This is the reason you need to setup test certificates using the steps below too. Use this method only to try out the operator to get familiar.

1. Clone the repository into the following folder `<GOPATH>/src/github.com/Azure`.

2. Make sure the environment variable `GO111MODULE` is set to `on`.

    ```bash
    export GO111MODULE=on
    ```

3. Make sure all the below environment variables have been set appropriately. The possible values for `AZURE_CLOUD_ENV` are AzurePubliCloud (default if unspecified), AzureChinaCloud, AzureUSGovernmentCloud, AzureGermanCloud.

    ```bash
    export AZURE_TENANT_ID=xxxx
    export AZURE_SUBSCRIPTION_ID=xxxx
    export AZURE_CLIENT_ID=xxxx
    export AZURE_CLIENT_SECRET=xxxx
    export AZURE_CLOUD_ENV=AzurePublicCloud
    ```

4. Install test certificates using `make generate-test-certs`.

5. Install the CRDs defined in the config/crd/bases folder using the command
    `make install`
    You will see output as below.

    ```
    kubectl apply -f config/crd/bases
    customresourcedefinition.apiextensions.k8s.io/eventhubnamespaces.azure.microsoft.com created
    customresourcedefinition.apiextensions.k8s.io/eventhubs.azure.microsoft.com created
    customresourcedefinition.apiextensions.k8s.io/resourcegroups.azure.microsoft.com created
    customresourcedefinition.apiextensions.k8s.io/azuresqldatabases.azure.microsoft.com created
    customresourcedefinition.apiextensions.k8s.io/azuresqlfirewallrules.azure.microsoft.com created
    customresourcedefinition.apiextensions.k8s.io/azuresqlservers.azure.microsoft.com configured
    ```

6. Copy the generated certificates to the correct location by first attempting to run the operator locally using
   ```make run```
   This will cause the operator to run (briefly) and will display an error (as well the correct location to copy generated certificates to). You will see something similar to this in the terminal window indicating that the controller started running and then subsequently failed:

    ```
    go fmt ./...
    go vet ./...
    go run ./main.go

    ...

    2020-05-13T17:04:26.587-0700	INFO	controller-runtime.builder	skip registering a validating webhook, admission.Validator interface is not implemented	{"GVK": "azure.microsoft.com/v1alpha1, Kind=AzureSqlFirewallRule"}
    2020-05-13T17:04:26.587-0700	INFO	controller-runtime.builder	conversion webhook enabled	{"object": {"metadata":{"creationTimestamp":null},"spec":{"server":""},"status":{}}}
    2020-05-13T17:04:26.587-0700	INFO	controller-runtime.builder	skip registering a mutating webhook, admission.Defaulter interface is not implemented	{"GVK": "azure.microsoft.com/v1alpha1, Kind=AzureSqlFailoverGroup"}
    2020-05-13T17:04:26.587-0700	INFO	controller-runtime.builder	skip registering a validating webhook, admission.Validator interface is not implemented	{"GVK": "azure.microsoft.com/v1alpha1, Kind=AzureSqlFailoverGroup"}
    2020-05-13T17:04:26.587-0700	INFO	controller-runtime.builder	conversion webhook enabled	{"object": {"metadata":{"creationTimestamp":null},"spec":{"location":"","server":"","failoverpolicy":"","failovergraceperiod":0,"secondaryserver":"","secondaryserverresourcegroup":"","databaselist":null},"status":{}}}
    2020-05-13T17:04:26.587-0700	INFO	controller-runtime.builder	skip registering a mutating webhook, admission.Defaulter interface is not implemented	{"GVK": "azure.microsoft.com/v1alpha1, Kind=BlobContainer"}
    2020-05-13T17:04:26.587-0700	INFO	controller-runtime.builder	skip registering a validating webhook, admission.Validator interface is not implemented	{"GVK": "azure.microsoft.com/v1alpha1, Kind=BlobContainer"}
    2020-05-13T17:04:26.587-0700	INFO	controller-runtime.builder	conversion webhook enabled	{"object": {"metadata":{"creationTimestamp":null},"spec":{"location":""},"status":{}}}
    2020-05-13T17:04:26.587-0700	INFO	setup	starting manager
    2020-05-13T17:04:26.587-0700	INFO	controller-runtime.webhook.webhooks	starting webhook server
    2020-05-13T17:04:26.587-0700	INFO	controller-runtime.manager	starting metrics server	{"path": "/metrics"}
    2020-05-13T17:04:26.588-0700	DEBUG	controller-runtime.manager	non-leader-election runnable finished	{"runnable type": "*webhook.Server"}
    2020-05-13T17:04:26.588-0700	INFO	controller-runtime.controller	Starting EventSource	{"controller": "keyvault", "source": "kind source: /, Kind="}
    2020-05-13T17:04:26.588-0700	ERROR	setup	problem running manager	{"error": "open /var/folders/1_/d50yhxgd357cfl0yjdclppfc0000gn/T/k8s-webhook-server/serving-certs/tls.crt: no such file or directory"}
    github.com/go-logr/zapr.(*zapLogger).Error
        /Users/username/Code/go/pkg/mod/github.com/go-logr/zapr@v0.1.0/zapr.go:128
    main.main
        /Users/username/Code/go/src/github.com/WilliamMortlMicrosoft/azure-service-operator/main.go:781
    runtime.main
        /usr/local/Cellar/go/1.14.2_1/libexec/src/runtime/proc.go:203
    exit status 1
    make: *** [run] Error 1`
    ```

    Look for the listed directory given in the *problem running manager* log entry, above that directory is: "/var/folders/1_/d50yhxgd357cfl0yjdclppfc0000gn/T/". Next, run the following command `cp -r /tmp/k8s-webhook-server /var/folders/1_/d50yhxgd357cfl0yjdclppfc0000gn/T/`.

    Please note, copying the generated certificates into the this folder only needs to be done only once (or when the machine is rebooted or the local Kubernetes cluster is cleared). You do not need to do this step every time.

7. Run the operator locally using
   ```make run```
   This will cause the operator to run and "watch" for events on this terminal. You will need to open a new terminal to trigger the creation of a custom resource.

   You will see something like this on the terminal window indicating that the controller is running.

   ```
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
    2019-09-24T12:18:10.424-0600	INFO	controller-runtime.controller	Starting EventSource	{"controller": "azuresqlserver", "source": "kind source: /, Kind="}
    2019-09-24T12:18:10.424-0600	INFO	controller-runtime.controller	Starting EventSource	{"controller": "azuresqldatabase", "source": "kind source: /, Kind="}
    2019-09-24T12:18:10.424-0600	INFO	controller-runtime.controller	Starting EventSource	{"controller": "azuresqlfirewallrule", "source": "kind source: /, Kind="}
    2019-09-24T12:18:10.424-0600	INFO	setup	starting manager
    2019-09-24T12:18:10.424-0600	INFO	controller-runtime.manager	starting metrics server	{"path": "/metrics"}
    2019-09-24T12:18:10.526-0600	INFO	controller-runtime.controller	Starting Controller	{"controller": "eventhub"}
    2019-09-24T12:18:10.526-0600	INFO	controller-runtime.controller	Starting Controller	{"controller": "resourcegroup"}
    2019-09-24T12:18:10.526-0600	INFO	controller-runtime.controller	Starting Controller	{"controller": "azuresqldatabase"}
    2019-09-24T12:18:10.526-0600	INFO	controller-runtime.controller	Starting Controller	{"controller": "azuresqlfirewallrule"}
    2019-09-24T12:18:10.526-0600	INFO	controller-runtime.controller	Starting Controller	{"controller": "eventhubnamespace"}
    2019-09-24T12:18:10.526-0600	INFO	controller-runtime.controller	Starting Controller	{"controller": "azuresqlserver"}
    2019-09-24T12:18:10.527-0600	INFO	controller-runtime.certwatcher	Updated current TLS certiface
    2019-09-24T12:18:10.527-0600	INFO	controller-runtime.certwatcher	Starting certificate watcher
    2019-09-24T12:18:10.626-0600	INFO	controller-runtime.controller	Starting workers	{"controller": "azuresqldatabase", "worker count": 1}
    2019-09-24T12:18:10.626-0600	INFO	controller-runtime.controller	Starting workers	{"controller": "resourcegroup", "worker count": 1}
    2019-09-24T12:18:10.626-0600	INFO	controller-runtime.controller	Starting workers	{"controller": "azuresqlfirewallrule", "worker count": 1}
    2019-09-24T12:18:10.626-0600	INFO	controller-runtime.controller	Starting workers	{"controller": "eventhub", "worker count": 1}
    2019-09-24T12:18:10.626-0600	INFO	controller-runtime.controller	Starting workers	{"controller": "azuresqlserver", "worker count": 1}
    2019-09-24T12:18:10.626-0600	INFO	controller-runtime.controller	Starting workers	{"controller": "eventhubnamespace", "worker count": 1}
   ```

8. Open a new terminal window. Trigger the creation of a custom resource using  kubectl and the sample YAML file provided.
   For instance, you would use the following command to create a SQL server:

   ```shell
   kubectl apply -f config/samples/azure_v1alpha1_azuresqlserver.yaml
   azuresqlserver.azure.microsoft.com/sqlserver-sample created
   ```

9. You should see logs on the other terminal from the operator when this custom resource is being created.

    ```
    2019-09-24T12:27:12.450-0600	DEBUG	controller-runtime.manager.events	Normal	{"object": {"kind":"AzureSqlServer","namespace":"default","name":"sqlserver-sample1","uid":"ed3774af-def8-11e9-90c4-025000000001","apiVersion":"azure.microsoft.com/v1alpha1","resourceVersion":"194427"}, "reason": "Updated", "message": "finalizer azuresqlserver.finalizers.azure.com added"}
    2019-09-24T12:27:12.450-0600	DEBUG	controller-runtime.manager.events	Normal	{"object": {"kind":"AzureSqlServer","namespace":"default","name":"sqlserver-sample1","uid":"ed3774af-def8-11e9-90c4-025000000001","apiVersion":"azure.microsoft.com/v1alpha1","resourceVersion":"194427"}, "reason": "Submitting", "message": "starting resource reconciliation"}
    2019-09-24T12:27:17.129-0600	INFO	controllers.AzureSqlServer	mutating secret bundle
    2019-09-24T12:27:17.144-0600	INFO	controllers.AzureSqlServer	waiting for provision to take effect	{"azuresqlserver": "default/sqlserver-sample1"}
    2019-09-24T12:27:17.350-0600	DEBUG	controller-runtime.manager.events	Normal	{"object": {"kind":"AzureSqlServer","namespace":"default","name":"sqlserver-sample1","uid":"ed3774af-def8-11e9-90c4-025000000001","apiVersion":"azure.microsoft.com/v1alpha1","resourceVersion":"194437"}, "reason": "Checking", "message": "instance in NotReady state"}
    2019-09-24T12:27:17.359-0600	INFO	controllers.AzureSqlServer	Got ignorable error	{"azuresqlserver": "default/sqlserver-sample1", "type": "ResourceNotFound"}
    2019-09-24T12:27:17.564-0600	DEBUG	controller-runtime.manager.events	Normal	{"object": {"kind":"AzureSqlServer","namespace":"default","name":"sqlserver-sample1","uid":"ed3774af-def8-11e9-90c4-025000000001","apiVersion":"azure.microsoft.com/v1alpha1","resourceVersion":"194439"}, "reason": "Checking", "message": "instance in NotReady state"}
    2019-09-24T12:27:17.570-0600	INFO	controllers.AzureSqlServer	Got ignorable error	{"azuresqlserver": "default/sqlserver-sample1", "type": "ResourceNotFound"}
    2019-09-24T12:28:17.805-0600	DEBUG	controller-runtime.manager.events	Normal	{"object": {"kind":"AzureSqlServer","namespace":"default","name":"sqlserver-sample1","uid":"ed3774af-def8-11e9-90c4-025000000001","apiVersion":"azure.microsoft.com/v1alpha1","resourceVersion":"194439"}, "reason": "Checking", "message": "instance in Ready state"}
    2019-09-24T12:28:19.010-0600	DEBUG	controller-runtime.controller	Successfully Reconciled	{"controller": "sqlserver", "request": "default/sqlserver-sample1"}
    2019-09-24T12:28:19.010-0600	DEBUG	controller-runtime.manager.events	Normal	{"object": {"kind":"AzureSqlServer","namespace":"default","name":"sqlserver-sample1","uid":"ed3774af-def8-11e9-90c4-025000000001","apiVersion":"azure.microsoft.com/v1alpha1","resourceVersion":"194518"}, "reason": "Provisioned", "message": "sqlserver sqlserver-sample1 provisioned "}
    2019-09-24T12:28:19.202-0600	DEBUG	controller-runtime.manager.events	Normal	{"object": {"kind":"AzureSqlServer","namespace":"default","name":"sqlserver-sample1","uid":"ed3774af-def8-11e9-90c4-025000000001","apiVersion":"azure.microsoft.com/v1alpha1","resourceVersion":"194518"}, "reason": "Checking", "message": "instance in Ready state"}
    2019-09-24T12:28:20.331-0600	DEBUG	controller-runtime.controller	Successfully Reconciled	{"controller": "azuresqlserver", "request": "default/sqlserver-sample1"}
    2019-09-24T12:28:20.331-0600	DEBUG	controller-runtime.manager.events	Normal	{"object": {"kind":"AzureSqlServer","namespace":"default","name":"sqlserver-sample1","uid":"ed3774af-def8-11e9-90c4-025000000001","apiVersion":"azure.microsoft.com/v1alpha1","resourceVersion":"194518"}, "reason": "Provisioned", "message": "sqlserver sqlserver-sample1 provisioned "}
    ```

10. Once the operator is running locally, in order to view debugging (Prometheus-based) metrics for the Azure operator, open a web browser and navigate to the [Metrics Endpoint](http://127.0.0.1:8080/metrics).

If you are using the Kubernetes cluster that comes with "Docker Desktop" and would like to view the Prometheus metrics, you can redirect port 8080 to the local machine using the following command:

```
kubectl port-forward deployment/controller-manager 8080
```

Then, open a web browser and navigate to the [Metrics Endpoint](http://127.0.0.1:8080/metrics).

## Developing using VSCode with Remote - Containers

If you're using VSCode with [Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) extensions installed, you can quickly have your environment set up and ready to go, with everything you need to get started.

1. Clone the repository into the following folder `~/go/src/github.com/Azure`.
2. Make sure the environment variable `GO111MODULE` is set to `on`.

    ```
    export GO111MODULE=on
    ```

3. Install test certificates using `make generate-test-certs`.
4. Open this folder `~/go/src/github.com/Azure` in VSCode.
5. Inside the folder `.devcontainer`, create a file called `.env` and using the following template, copy your environment variable details.

    ```txt
    AZURE_CLIENT_ID=

    AZURE_CLIENT_SECRET=

    AZURE_SUBSCRIPTION_ID=

    AZURE_TENANT_ID=
    ```

6. Open the Command Pallet (`Command+Shift+P` on MacOS or `CTRL+Shift+P` on Windows), type `Remote-Containers: Open Folder in Container...`, select the ```Azure-service-operator``` folder and hit enter.
7. VSCode will relaunch and start building our development container. This will install all the necessary dependencies required for you to begin developing.
8. Once the container has finished building, you can now start testing your Azure Service Operator within your own local kubernetes environment via the terminal inside VSCode.

**Note**: after the DevContainer has finished building, the Kind cluster will start initializing and installing the Azure Service Operator in the background. This will take some time before it is available.

To see when the Kind cluster is ready, use `docker ps -a` to list your running containers, look for `IMAGE` with the name `azure-service-operator_devcontainer_docker-in-docker...`. Using that image's `CONTAINER ID`, use `docker logs -f CONTAINER ID` to view the logs from the container setting up your cluster.
9. Use `kubectl apply` with the sample YAML files to create custom resources for testing.
For eg., use  `kubectl apply -f config/samples/azure_v1alpha1_azuresqlserver.yaml` from the terminal to create a SQL server using the operator.
`kubectl describe SqlServer` would show the events that indicate if the resource is created or being created.

## Want to contribute a new operator?

If you are ready to contribute a new operator to the Azure Service Operators project, please follow the step-by-step guide [here](newoperatorguide.md).