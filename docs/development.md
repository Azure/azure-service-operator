# K8s Infra Development

## Prerequisites for Getting Started
This project expects you have the following installed:
- [Golang 1.13+](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/install/)
- [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest) 
- [JQ](https://stedolan.github.io/jq/)
- [optional: Tilt](https://docs.tilt.dev/install.html)

## Tilt up - with live reload
When you want a quick development loop with fast reloading of changes, use the following command. 
This will start [tilt](https://tilt.dev), a fast reloading build loop.

```bash
make tilt-up
```

### What does this command do?
- Installs the build / local dev tools needed (kind, kubectl, controller-tools) into `./hack/tools/bin`
  - Always use the bins in `./hack/tools/bin` to interact with the project. Otherwise, nuances in behavior between versions will lead to pain.
- Creates a service principal using Azure CLI and jq
- Builds a Kind cluster
  - adds cert-manager to the cluster
  - injects Azure service principal secrets for the controller
- Iterative Dev Loop
  - Builds the controller
  - Generates the manifests
  - Builds a docker image of the controller
  - Pushes the docker image of the controller to a local registry
  - Installs the manifests into the Kind cluster

## Run Tests and Deploy using [Kind](https://kind.sigs.k8s.io)
```bash
make deploy-kind
```

### What does this command do?
- Installs the build / local dev tools needed (kind, kubectl, controller-tools) into `./hack/tools/bin`
  - Always use the bins in `./hack/tools/bin` to interact with the project. Otherwise, nuances in behavior between versions will lead to pain.
- Creates a service principal using Azure CLI and jq
- Builds the controller
- Runs unit tests
- Builds a Kind cluster
  - adds cert-manager to the cluster
  - injects Azure service principal secrets for the controller
- Generates the manifests
- Builds a docker image of the controller
- Pushes the docker image of the controller to a local registry
- Installs the manifests into the Kind cluster

## The controller is running, then what?
You can now interact with the cluster and deploy Azure resources via Kubernetes CRDs.
```bash
export KUBECONFIG=$HOME/.kube/kind-k8sinfra
export PATH=$HOME/hack/tools/bin:$PATH
kubectl get pod -A                                                                    # see all the pods
kubectl logs pod/${the_controller_pod_name} -c manager --namespace k8s-infra-system   # get the controller logs
kubectl apply -f config/samples/microsoft.resources_v20191001_resourcegroup.yaml      # create a resource group
```

### Reset your local cluster
If you find you have completely messed up your local Kind cluster, just wipe it out and start fresh.
```bash
make kind-reset
```