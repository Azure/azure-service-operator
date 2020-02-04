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

    ```
    kubectl apply -f config/crd/bases
    customresourcedefinition.apiextensions.k8s.io/eventhubnamespaces.azure.microsoft.com created
    customresourcedefinition.apiextensions.k8s.io/eventhubs.azure.microsoft.com created
    customresourcedefinition.apiextensions.k8s.io/resourcegroups.azure.microsoft.com created
    customresourcedefinition.apiextensions.k8s.io/azuresqldatabases.azure.microsoft.com created
    customresourcedefinition.apiextensions.k8s.io/azuresqlfirewallrules.azure.microsoft.com created
    customresourcedefinition.apiextensions.k8s.io/azuresqlservers.azure.microsoft.com configured
    ```

5. Run the operator locally using
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

6. Open a new terminal window. Trigger the creation of a custom resource using      kubectl and the sample YAML file provided.
   For instance, you would use the following command to create a SQL server:

   ```shell
   kubectl apply -f config/samples/azure_v1alpha1_azuresqlserver.yaml
   azuresqlserver.azure.microsoft.com/sqlserver-sample created
   ```

7. You should see logs on the other terminal from the operator when this custom     resource is being created.

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

8. Once the operator is running locally, in order to view debugging (Prometheus-based) metrics for the Azure operator, open a web browser and navigate to the [Metrics Endpoint](http://127.0.0.1:8080/metrics).

If you are using the Kubernetes cluster that comes with "Docker Desktop" and would like to view the Prometheus metrics, you can redirect port 8080 to the local machine using the following command:

```
kubectl port-forward deployment/controller-manager 8080
```

Then, open a web browser and navigate to the [Metrics Endpoint](http://127.0.0.1:8080/metrics).

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
For eg., use  `kubectl apply -f config/samples/azure_v1alpha1_azuresqlserver.yaml` from the terminal to create a SQL server using the operator. 
`kubectl describe SqlServer` would show the events that indicate if the resource is created or being created.

## Step-by-Step for developing a new operator using the generic controller

This project utilizes a generic async controller that separates out the Kubernetes controller logic from the corresponding Azure service interactions. There are a few extra steps (listed below) you will need to follow while using kubebuilder to generate a new operator.

1. Clone the repository to your computer. Run `go mod download` to download dependencies. This may take several minutes.
2. Ensure you have `Kubebuilder` installed on your computer.
3. From the root folder of the cloned repository, run the following command

```
kubebuilder create api --group azure --version v1alpha1 --kind <AzureNewType>

kubebuilder create api --group azure --version v1alpha1 --kind AzureSQLServer
```

4. When you run this command, you will be prompted with two questions. Answer yes to both of these.

This will create

- a CRD types file in /api/v1alpha1
- a controller file in controllers/{resource_type}_controller.go
- and other needed scaffolding.

```
➜  azure-service-operator git:(testbranch) ✗ kubebuilder create api --group azure --version v1alpha1 --kind AzureNewType
Create Resource [y/n]
y
Create Controller [y/n]
y
Writing scaffold for you to edit...
api/v1alpha1/azurenewtype_types.go
controllers/azurenewtype_controller.go
Running make...
make: kind: Command not found
/bin/sh: kind: command not found
go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.2.0
go: finding sigs.k8s.io/controller-tools/cmd/controller-gen v0.2.0
go: finding sigs.k8s.io/controller-tools/cmd v0.2.0
go: finding sigs.k8s.io v0.2.0
/Users/jananiv/go/bin/controller-gen "crd:trivialVersions=true" rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases
/Users/jananiv/go/bin/controller-gen object:headerFile=./hack/boilerplate.go.txt paths=./api/...
go fmt ./...
go vet ./...
go build -o bin/manager main.go
```

5. *Updating the types file*

    Open the `AzureNewType_types.go` file that was generated under `api/v1alpha1`.
    a. Replace the generated `Status` struct with the shared `ASOStatus` from `/api/{version}/aso_types.go` as below

    ```go
    type AzureNewType struct {
        metav1.TypeMeta   `json:",inline"`
        metav1.ObjectMeta `json:"metadata,omitempty"`
        
        Spec   AzureNewTypeSpec `json:"spec,omitempty"`
        Status ASOStatus       `json:"status,omitempty"`
        }
    ```

    b. Define the fields for the `AzureNewTypeSpec` struct. This will be all the parameters that you will need to create a resource of type `AzureNewType`.

    For instance, if I need to create a resource of type ResourceGroup, I need a Subscription ID, Name, and a Location.

    The `Subscription ID` is something we configure through config for the entire operator, so it can be omitted.
    The `Name` for every resource is going to be the Kubernetes `Metadata.Name` we pass in the manifest, so omit that from the Spec as well.
    In this case you would only  add `Location` of type `string` to the Spec.

    Here is an example of how this might look. Note that we use camel case for the names of the fields in the json specification.

    ```go
    type AzureNewTypeSpec struct {
        Location        string `json:"location"`
        ResourceGroup   string `json:"resourceGroup,omitempty"`
        }
    ```

    You can refer to the other types in this repo as examples to do this.

    c. Add the Kubebuilder directive to indicate that you want to treat `Status` as a subresource of the CRD. This allows for more efficient updates and the generic controller assumes that this is enabled for all the types.

    ```go
    // +kubebuilder:object:root=true
    // +kubebuilder:subresource:status   <--- This is the line we need to add

    // AzureNewType is the Schema for the azurenewtypes API
    type AzureNewType struct {
        metav1.TypeMeta   `json:",inline"`
        metav1.ObjectMeta `json:"metadata,omitempty"`

        Spec   AzureNewTypeSpec   `json:"spec,omitempty"`
        Status AzureNewTypeStatus `json:"status,omitempty"`
    }
    ```

6. *Updating the generated controller file*

    Open the `azurenewtype_controller.go` file generated under the `controllers` directory.
    a. Update the `AzureNewTypeReconciler` struct by removing the fields in there and replacing it with the below.

    ```go
    // AzureNewTypeReconciler reconciles a AzureNewType object
    type AzureNewTypeReconciler struct {
        Reconciler *AsyncReconciler
    }
    ```

    b. Update the `Reconcile` function code in this file to the following:

    ```go
    // +kubebuilder:rbac:groups=azure.microsoft.com,resources=azurenewtypes,verbs=get;list;watch;create;update;patch;delete
    // +kubebuilder:rbac:groups=azure.microsoft.com,resources=azurenewtypes/status,verbs=get;update;patch

    func (r *AzureNewTypeReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
        return r.Reconciler.Reconcile(req, &azurev1alpha1.AzureNewType{})
    }
    ```

7. *Updating the sample YAML file*

    Open the `azure_v1alpha1_azurenewtype.yaml` file under `config/samples`.
    Update the spec portion of this file with the fields that you added to the Spec in step 5b above with corresponding values.

    Refer to the other sample YAML files in the repo for how to do this correctly.

8. *Implementing the actual resource creation/deletion logic*

    This is the service specific piece that differs from one Azure service to another, and the crux of the service provisioning/deletion logic used by the controller.

    a. The guidance is to add a folder under `pkg/resourcemanager` for the new resource, like for instance, `pkg/resourcemanager/newresource`

    b. Use subfolders under this folder if you have subresources. For instance, for PostgreSQL we have separate sub folders for `server`, `database` and `firewallrule`

    c. Add two files under this folder
        azurenewtype_manager.go
        azurenewtype_client.go

    d. The azurenewtype_manager.go file would implement the ARMClient as follows,and in addition have any other functions you need for Create, Delete, and Get of the resource.

    ```go
    type AzureNewTypeManager interface {
        // Functions for Create, Delete and Get for the resource as needed

        // also embed async client methods
        resourcemanager.ARMClient
    }

    ```

    e. The `azurenewtype_client.go` file defines a struct that implements the `AzureNewTypeManager` interface.
       - This struct would have the following functions attached to it from the ARMClient interface - `Ensure`, `Delete`, `GetParents`.
       - It would also have a `convert` function to convert the runtime object into the appropriate type
       - It would also have the other functions defined in the `AzureNewTypeManager` interface

    Some key points to note:
    (i) The Ensure and Delete functions return as the first return value, a bool which indicates if the resource was found in Azure. So Ensure() if successful would return `true` and Delete() if successful would return `false`
    (ii) On successful provisioning in `Ensure()`, set instance.Status.Message to `successfully provisioned` to be consistent across all controllers. (There is a constant called `SuccessMsg` in the `resourcemanager` package that you can use for this to be consistent)
    (ii) The GetParents() function returns the Azure Resource Manager (ARM) hierarchy of the resource. The order here matters - the immediate hierarchical resource should be returned first. For instance, for an Azure SQL database, the first parent should be Azure SQL server followed by the Resource Group.

    An example is shown below:

    ```go
        type AzureNewTypeClient struct {
        Telemetry telemetry.PrometheusTelemetry
        }

        func NewAzureNewTypeClient(log logr.Logger) *AzureNewTypeClient {
            return &AzureNewTypeClient{
                Telemetry: telemetry.InitializePrometheusDefault(
                    log,
                    "NewType",
                ),
            }
        }
        ...
        ...
        func (p *AzureNewTypeClient) Ensure(ctx context.Context, obj runtime.Object) (found bool, err error) {
            instance, err := p.convert(obj)
            if err != nil {
                return true, err
            }
            // Add logic to idempotently create the resource

            return true, nil
        }

        func (p *AzureNewTypeClient) Delete(ctx context.Context, obj runtime.Object) (found bool, err error) {
            instance, err := p.convert(obj)
            if err != nil {
                return true, err
            }
            // Add logic to idempotently delete the resource

            return false, nil
        }

        func (p *AzureNewTypeClient) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

            instance, err := p.convert(obj)
            if err != nil {
                return nil, err
            }

            // Update this based on the resource to add other parents

            return []resourcemanager.KubeParent{
                {
                    Key: types.NamespacedName{
                        Namespace: instance.Namespace,
                        Name:      instance.Spec.ResourceGroup,
                    },
                    Target: &azurev1alpha1.ResourceGroup{},
                },
            }, nil

        }

        func (p *AzureNewTypeClient) convert(obj runtime.Object) (*v1alpha1.AzureNewType, error) {
            local, ok := obj.(*v1alpha1.AzureNewType)
            if !ok {
                return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
            }
            return local, nil
        }

        // Add other Create, Delete, Get methods for the resource as needed

        ```

9. *Tests*

    Add the following tests for your new operator:
    (i) Unit test for the new type: You will add this as a file named `azurenewtype_types_test.go` under `api/v1alpha`. Refer to the existing tests in the repository to author this for your new type.
    (ii) Functional tests for the resource manager: These will be test files under `pkg/resourcemanager/newresource`. You can refer to the tests under `pkg/resourcemanage/psql/server` as reference for writing these. These tests help validate the resource creation and deletion in Azure.
    Make sure to add this folder with the tests to the `make test-existing` target in the `Makefile`
    *Tip*: You can run `go test` in this folder to test the functions before testing with the controller end-to-end.
    (iii) Controller tests: This will be a file named `azurenewresourcetype_controller_test.go` under the `controllers` folder. Refer to the other controller tests in the repo for how to write this test. This test validates the new controller to make sure the resource is created and deleted in Kubernetes effectively.

10. *Mocks for the controller test*

    In order to run the Controller tests (9.(iii) above) without having to communicate with Azure, you will generate a mock of the `AzureNewResourceTypeManager` in `pkg/resourcemanager/mocks`. You can refer to the mocks under `pkg/resourcemanager/mock/psql` for how to write this mock.

11. *Instantiating the reconciler*

    The last step to tie everything together is to ensure that your new controller's reconciler is instantiated in both the `main.go` and the `suite_test.go` (under `controllers` folder) files.

    ```go
    main.go

    if err = (&controllers.AzureNewResourceTypeReconciler{
            Reconciler: &controllers.AsyncReconciler{
                Client:      mgr.GetClient(),
                AzureClient: <var of type AzureNewResourceTypeManager>,
                Telemetry: telemetry.InitializePrometheusDefault(
                    ctrl.Log.WithName("controllers").WithName("AzureNewResourceType"),
                    "AzureNewResourceType",
                ),
                Recorder: mgr.GetEventRecorderFor("AzureNewResourceType-controller"),
                Scheme:   scheme,
            },
        }).SetupWithManager(mgr); err != nil {
            setupLog.Error(err, "unable to create controller", "controller", "AzureNewResourceType")
            os.Exit(1)
        }
    ```

    ```go
    suite_test.go

    err = (&AzureNewResourceTypeReconciler{
        Reconciler: &AsyncReconciler{
            Client:      k8sManager.GetClient(),
            AzureClient: <var of type AzureNewResourceTypeManager>,
            Telemetry: telemetry.InitializePrometheusDefault(
                ctrl.Log.WithName("controllers").WithName("AzureNewResourceType"),
                "AzureNewResourceType",
                ),
            Recorder: k8sManager.GetEventRecorderFor("AzureNewResourceType-controller"),
            Scheme:   k8sManager.GetScheme(),
        },
    }).SetupWithManager(k8sManager)
    Expect(err).ToNot(HaveOccurred())
    ```

12. Install the new CRD and generate the manifests needed using the following commands. This is required in order to generate canonical resource definitions (manifests as errors about a DeepCopyObject() method missing).

    ```
    make manifests
    make install
    ```

    Run tests using `make test` and deploy using `make deploy`

    If you make changes to the operator and want to update the deployment without recreating the cluster (when testing locally), you can use the `make update` to update your Azure Operator pod. If you need to rebuild the docker image without cache, use `make ARGS="--no-cache" update`

**Notes:**

- Run `make manifests` if you find the property you add doesn't work.

- Finalizers are recommended as they make deleting a resource more robust. Refer to [Using Finalizers](https://book.kubebuilder.io/reference/using-finalizers.html) for more information
