# Azure Operator (for Kubernetes)

[![Build Status](https://dev.azure.com/azure/azure-service-operator/_apis/build/status/Azure.azure-service-operator?branchName=master)](https://dev.azure.com/azure/azure-service-operator/_build/latest?definitionId=36&branchName=master)

> This project is experimental. Expect the API to change. It is not recommended for production environments.

## Introduction

Kubernetes offers the facility of extending it's API through the concept of 'Operators' ([Introducing Operators: Putting Operational Knowledge into Software](https://coreos.com/blog/introducing-operators.html)). This repository contains the resources and code to provision a Resource group and Azure Event Hub using Kubernetes operator.

The Azure Operator comprises of:

- The golang application is a Kubernetes controller that watches Customer Resource Definitions (CRDs) that define a Resource Group and Event Hub

The project was built using

1. [Kubebuilder](https://book.kubebuilder.io/)

## Building and Running from Source

### Prerequisites And Assumptions

1. You have GoLang installed.
2. [Docker](https://docs.docker.com/install/) is installed and running.
3. You have the kubectl command line (kubectl CLI) installed.
4. You have access to a Kubernetes cluster.
    - It can be a local hosted Cluster like
    [Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/),
    [Kind](https://github.com/kubernetes-sigs/kind) or Docker for desktop installed locally with RBAC enabled.
    - If you opt for Azure Kubernetes Service ([AKS](https://azure.microsoft.com/en-au/services/kubernetes-service/)), you can use:
    `az aks get-credentials --resource-group $RG_NAME --name $Cluster_NAME`
    - Kubectl: Client version 1.14 Server Version 1.12

    **Note:** it is recommended to use [Kind](https://github.com/kubernetes-sigs/kind) as it is needed for testing Webhooks.
5. Install [Kubebuilder](https://book.kubebuilder.io/), following the linked installation instructions.
6. [Kustomize](https://github.com/kubernetes-sigs/kustomize) is also needed. This must be installed via `make install-kustomize` (see section below).

Basic commands to check your cluster

```shell
    kubectl config get-contexts
    kubectl cluster-info
    kubectl version
    kubectl get pods -n kube-system
```

### Quick Start

If you're using VSCode with [Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) extensions installed, you quickly have you're environment set up and ready to go with everything you need to get started.

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

**Note**: after the DevContainer has finished building, the kind cluster will initialising and installing the Azure Service Operator in the background. This will take some time before it is available.

To see when the kind cluster is ready, use `docker ps -a` to list your running containers, look for `IMAGE` with the name `azure-service-operator_devcontainer_docker-in-docker...`. Using that image's `CONTAINER ID`, use `docker logs -f CONTAINER ID` to view the logs from the container setting up your cluster.

### Getting started

1. Clone the repository from the following folder `<GOPATH>/src/github.com/Azure`.

2. Make sure the environment variable `GO111MODULE=on` is set.

3. Update the values in `azure_v1_eventhub.yaml` to reflect the resource group and event hub you want to provision

4. Install [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

    ```shell
        GO111MODULE="on" go get sigs.k8s.io/kind@v0.4.0 && kind create cluster
        kind create cluster
        export KUBECONFIG="$(kind get kubeconfig-path --name="kind")"
        kubectl cluster-info
        IMG="docker.io/yourimage:tag" make docker-build
        kind load docker-image docker.io/yourimage:tag --loglevel "trace"
        make deploy
    ```

5. Create a Service Principal
    If you don't have a Service Principal create one from the Azure CLI:

    ```bash
    az ad sp create-for-rbac --role Contributor
    ```

    Then make sure this service principal has rights assigned to provision resources on your Azure account.
  
6. Set the environment variables `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, `AZURE_CLIENT_SECRET`, `AZURE_SUBSCRIPTION_ID`, `REQUEUE_AFTER`.

    If you are running it on Windows the environment variables should not have quotes.

    It should be set in this way:
    `SET  AZURE_TENANT_ID=11xxxx-xxx-xxx-xxx-xxxxx`
    and the VSCode should be run from the same session/command window

7. Set up the Cluster

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

8. Install [kustomize](https://github.com/kubernetes-sigs/kustomize) using `make install-kustomize`.

9. Install the azure_v1_eventhub CRD in the configured Kubernetes cluster folder ~/.kube/config,

    run `kubectl apply -f config/crd/bases` or `make install`

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
