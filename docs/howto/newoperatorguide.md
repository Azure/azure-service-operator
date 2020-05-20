## Step-by-Step for developing a new operator using the generic controller

This project utilizes a generic async controller that separates out the Kubernetes controller logic from the corresponding Azure service interactions. There are a few extra steps (listed below) you will need to follow while using kubebuilder to generate a new operator.

1. Clone the repository to your computer. Run `go mod download` to download dependencies. This may take several minutes.
2. Ensure you have `Kubebuilder` installed on your computer.
3. From the root folder of the cloned repository, run the following command:

```
kubebuilder create api --group azure --version v1alpha1 --kind <AzureNewType>

# e.g. kubebuilder create api --group azure --version v1alpha1 --kind AzureSQLServer
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
    Also add the Kubebuilder directives to designate short names for the CRD and to add additional columns to the output of `kubectl get`.

    ```go
    // +kubebuilder:object:root=true
    // +kubebuilder:subresource:status   <--- This is the line we need to add

    // +kubebuilder:resource:shortName=[shortnamehere],path=[crdpath]
    // +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
    // +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"

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

    **Note** Please make sure that you make all creation and deletion calls to Azure are asynchronous.

    a. The guidance is to add a folder under `pkg/resourcemanager` for the new resource, like for instance, `pkg/resourcemanager/newresource`

    b. Use subfolders under this folder if you have subresources. For instance, for PostgreSQL we have separate sub folders for `server`, `database` and `firewallrule`

    c. Add these files under this folder
        azurenewtype_manager.go
        azurenewtype_reconcile.go
        azurenewtype.go

    d. The `azurenewtype_manager.go` file would implement the interface that includes the ARMClient as follows, and in addition have any other functions you need for Create, Delete, and Get of the resource.

    ```go
    type AzureNewTypeManager interface {
        // Functions for Create, Delete and Get for the resource as needed

        // also embed async client methods
        resourcemanager.ARMClient
    }

    ```

    The `azurenewtype.go` file defines a struct that implements the `AzureNewTypeManager` interface. It has the definitions of the Create, Delete and Get functions included in the interface in the `azurenewtype_manager.go`

    Here is an example of what this struct looks like.

    **Note** Don't add a logger to this struct. Return all errors from this file to the controller so we can log it there.

    ```go
        type AzureNewTypeClient struct {}

        func NewAzureNewTypeClient() *AzureNewTypeClient {
            return &AzureNewTypeClient{}
        }
    ```

    e. The `azurenewtype_reconcile.go` file implements the  following functions in the `ARMClient` interface:
       -  `Ensure`, `Delete`, `GetParents`, `GetStatus`
       - It would also have a `convert` function to convert the runtime object into the appropriate type

    Some key points to note:
    (i) The Ensure and Delete functions return as the first return value, a bool which indicates if the resource was found in Azure. So Ensure() if successful would return `true` and Delete() if successful would return `false`

    (ii) On successful provisioning in `Ensure()`, 
    - set `instance.Status.Message` to the constant `SuccessMsg` in the `resourcemanager` package to be consistent across all controllers.
    - set `instance.Status.ResourceID` to the full Azure Resource ID of the resource
    - set `instance.Status.Provisioned` to `true` and `instance.Status.Provisioning` to `false`

    ```go
    func (p *AzureNewTypeClient) Ensure(ctx context.Context, obj runtime.Object) (found bool, err error) {
            instance, err := p.convert(obj)
            if err != nil {
                return true, err
            }
            // Add logic to idempotently create the resource

            // successful return
            return true, nil
        }

        func (p *AzureNewTypeClient) Delete(ctx context.Context, obj runtime.Object) (found bool, err error) {
            instance, err := p.convert(obj)
            if err != nil {
                return true, err
            }
            // Add logic to idempotently delete the resource

            // successful return
            return false, nil
        }
    ```

    (ii) The `GetParents()` function returns the Azure Resource Manager (ARM) hierarchy of the resource. The order here matters - the immediate hierarchical resource should be returned first. For instance, for an Azure SQL database, the first parent should be Azure SQL server followed by the Resource Group.

    An example is shown below:

    ```go
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
    ```

    (iii) The `GetStatus()` is a boilerplate function, use the below function and alter to use the struct you attach the function to.

    ```go
        func (p *AzureNewTypeClient) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
            instance, err := g.convert(obj)
            if err != nil {
                return nil, err
            }
            return &instance.Status, nil
        }
    ```

    (iv) The `convert()` function looks like the below, use the correct type based on the controller you are implementing.
  
    ```go
        func (p *AzureNewTypeClient) convert(obj runtime.Object) (*v1alpha1.AzureNewType, error) {
            local, ok := obj.(*v1alpha1.AzureNewType)
            if !ok {
                return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
            }
            return local, nil
        }
    ```

9. *Tests*

    Add the following tests for your new operator:
    (i) Unit test for the new type: You will add this as a file named `azurenewtype_types_test.go` under `api/v1alpha`. Refer to the existing tests in the repository to author this for your new type.

    (ii) Controller tests: This will be a file named `azurenewresourcetype_controller_test.go` under the `controllers` folder. Refer to the other controller tests in the repo for how to write this test. This test validates the new controller to make sure the resource is created and deleted in Kubernetes effectively.

10. *Instantiating the reconciler*

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

11. Install the new CRD and generate the manifests needed using the following commands. This is required in order to generate canonical resource definitions (manifests as errors about a DeepCopyObject() method missing).

    ```
    make generate
    make manifests
    make install
    ```

    Run controller tests using `make test-existing-controllers` and deploy using `make deploy`

    If you make changes to the operator and want to update the deployment without recreating the cluster (when testing locally), you can use the `make update` to update your Azure Operator pod. If you need to rebuild the docker image without cache, use `make ARGS="--no-cache" update`

12. Update the Helm Chart to include the new CRD.
    Run:
    ```
    make helm-chart-manifests
    make delete-helm-gen-manifests
    ```

    This will generate the manifests into the Helm Chart directory, and repackage them into a new Helm Chart tar.gz file. Add the newly modified files to the PR, which should be the following:
    ```
    charts/azure-service-operator-0.1.0.tgz
    charts/index.yaml
    ```

**Notes:**

- Run `make manifests` if you find the property you add doesn't work.

- Finalizers are recommended as they make deleting a resource more robust. Refer to [Using Finalizers](https://book.kubebuilder.io/reference/using-finalizers.html) for more information
