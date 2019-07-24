# Development

## Prerequisites

To get started you will need:

* a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster, e.g. [Azure Kubernetes Service](https://docs.microsoft.com/en-us/azure/aks/kubernetes-walkthrough).
* [kubebuilder](https://book.kubebuilder.io/quick-start.html#installation)
* [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

## Deploy Operator and Test

### Test it locally

1. Create Cluster.

    ```
    kind create cluster
    export KUBECONFIG="$(kind get kubeconfig-path --name="kind")"
    kubectl cluster-info
    ```

1. Install CRDs.

    ```
    make install
    ```

1. Run Controller.

    Setup the environment variables:

    ```
    export CLOUD_NAME=AzurePublicCloud
    export TENANT_ID=
    export SUBSCRIPTION_ID=
    export CLIENT_ID=
    export CLIENT_SECRET=
    ```

    Run your controller (this will run in the foreground, so switch to a new terminal if you want to leave it running):

    ```
    make run
    ```

    Refer to [kubebuilder's doc](https://book.kubebuilder.io/quick-start.html#test-it-out-locally).

1. Create a Custom Resource.

    Create your CR (make sure to edit them first to specify the fields). Example:

    ```
    kubectl apply -f examples/service/v1alpha1/storage.yaml
    ```

### Test it on a remote cluster

1. Create Cluster.

    ```
    az aks create -g <RG-NAME> -n <AKS-NAME>
    az aks get-credentials -g <RG-NAME> -n <AKS-NAME>
    kubectl cluster-info
    ```

1. Install CRDs.

    ```
    make install
    ```

1. Build and Push the image.

    ```
    IMG=<IMAGE-NAME-WITH-ORG> make build-and-push
    ```

    Update kustomize image patch file `config/default/manager_image_patch.yaml` for manager resource manually.

1. Run Controller.

    Update `config/manager/manager.yaml` with your service principal.

    ```
    make deploy
    ```

1. Create a Custom Resource.

    Create your CR (make sure to edit them first to specify the fields). Example:

    ```
    kubectl apply -f examples/service/v1alpha1/storage.yaml
    ```

## Add a New Custom Resource

### 1. Add a New API

```
kubebuilder create api --group service --version v1alpha1 --kind <Azure-Service-Name>
```

Refer to [kubebuilder's doc](https://book.kubebuilder.io/cronjob-tutorial/new-api.html)

### 2. Design an API

1. Try to create the specific Azure service, and download the template in the `Review+Create` step.
2. Upload the template to a storage account. For now, we can use the storage account `azureserviceoperator`.
3. Based on the template, we can figure out what the `Spec` should be like.
4. The `Status` should contain the resource group name, which can be used to delete the resource.

Refer to [kubebuilder's doc](https://book.kubebuilder.io/cronjob-tutorial/api-design.html)

Note:

* Don't forget to add `// +kubebuilder:subresource:status` if we want a status subresource.

* Run `make manifests` if you find the property you add doesn't work.

### 3. Delete external resource

[Using Finalizers](https://book.kubebuilder.io/reference/using-finalizers.html)
