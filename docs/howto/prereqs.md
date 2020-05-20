# Development and Deployment Dependencies

1. [GoLang](https://golang.org/dl/) is installed.
2. [Docker](https://docs.docker.com/install/) is installed and running.
3. [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) Command-line tool is installed.
4. A Kubernetes cluster. Can be any of the following but an AKS cluster is recommended for testing RBAC and multiple CRD versions (as webhooks are required)
    - Local: [Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube)
    - [Kind](https://github.com/kubernetes-sigs/kind), or,
    - [Docker for desktop](https://blog.docker.com/2018/07/kubernetes-is-now-available-in-docker-desktop-stable-channel/).
    - Or in the cloud: [Azure Kubernetes Service](https://azure.microsoft.com/en-us/services/kubernetes-service/)

5. [Kubebuilder](https://book.kubebuilder.io/) (notes for installing on OSX: [here](/docs/kubebuilder.md))

6. [Kustomize](https://github.com/kubernetes-sigs/kustomize) (This may be installed via `make install-kustomize`)

7. Create an identity for running the Operator as - Managed Identity or Service Principal.

```
az ad sp create-for-rbac --name ServicePrincipalName
```

```
az identity create -g ResourceGroup -n ManagedIdentityName --subscription SubscriptionId
```

Then make sure this identity has rights assigned to provision resources in your Azure Subscription and Resource Group.

8. Basic commands to check if you have an active Kubernetes cluster:

    ```shell
        kubectl config get-contexts
        kubectl cluster-info
        kubectl version
        kubectl get pods -n kube-system
    ```

    If you use an Azure Kubernetes Service (AKS) cluster, you can connect to it by using the following command:

    ```shell
        az aks get-credentials -g <ResourceGroup> -n <AKSClusterName>
    ```
