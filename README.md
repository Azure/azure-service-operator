# Azure Service Operator

The Azure Service Operator allows you to manage Azure resources using Kubernetes [Custom Resource Definitions (CRD)](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/).

## Prerequisites

* a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster, e.g. [Azure Kubernetes Service](https://docs.microsoft.com/en-us/azure/aks/kubernetes-walkthrough).
* [kubebuilder](https://book.kubebuilder.io/quick-start.html#installation)
* [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

## Getting Started

To walkthrough the demo, you'll need an AKS cluster because the service needs an external IP address.

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

1. Run Controller.

    Update `config/default/manager_auth_proxy_patch.yaml` with your service principal.

    ```
    make deploy
    ```

1. Run the demo.

    ```
    kubectl apply -f examples/demo/
    ```

1. Test the demo.

    To monitor progress, use the `kubectl get service` command with the `--watch` argument.

    ```
    kubectl get service azure-vote-front --watch
    ```

    Initially the `EXTERNAL-IP` for the `azure-vote-front` service is shown as pending.

    ```
    NAME               TYPE           CLUSTER-IP   EXTERNAL-IP   PORT(S)        AGE
    azure-vote-front   LoadBalancer   10.0.37.27   <pending>     80:30572/TCP   6s
    ```

    After the `EXTERNAL-IP` address changes from `pending` to an actual public IP address, open a web browser to the external IP address of your service.

# Contributing

This project welcomes contributions and suggestions.  Most contributions require you to agree to a
Contributor License Agreement (CLA) declaring that you have the right to, and actually do, grant us
the rights to use your contribution. For details, visit https://cla.opensource.microsoft.com.

When you submit a pull request, a CLA bot will automatically determine whether you need to provide
a CLA and decorate the PR appropriately (e.g., status check, comment). Simply follow the instructions
provided by the bot. You will only need to do this once across all repos using our CLA.

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/).
For more information see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq/) or
contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.
