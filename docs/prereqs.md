# Development and Decployment Dependencies


1. [GoLang](https://golang.org/dl/) is installed.
2. [Docker](https://docs.docker.com/install/) is installed and running.
3. [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) Command-line tool is installed.
4. A Kubernetes cluster.
    - Local: [Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/), [Kind](https://github.com/kubernetes-sigs/kind), or [Docker for desktop](https://blog.docker.com/2018/07/kubernetes-is-now-available-in-docker-desktop-stable-channel/).
    - Or in the cloud: [Azure Kubernetes Service](https://azure.microsoft.com/en-au/services/kubernetes-service/)

5. [Kubebuilder](https://book.kubebuilder.io/) (notes for installing on OSX: [here](/docs/kubebuilder.md))

6. [Kustomize](https://github.com/kubernetes-sigs/kustomize) (This may be installed via `make install-kustomize`)

7. Create a Service Principal
    If you don't have a Service Principal create one from Azure CLI:

    ```bash
    az ad sp create-for-rbac --role Contributor
    ```

    Then make sure this service principal has rights assigned to provision resources in your Azure Subscription and resource group.


8. Set the environment variables `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, `AZURE_CLIENT_SECRET`, `AZURE_SUBSCRIPTION_ID`, `REQUEUE_AFTER`.

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

Basic commands to check if you have an active Kubernetes cluster:

```shell
    kubectl config get-contexts
    kubectl cluster-info
    kubectl version
    kubectl get pods -n kube-system
```
