# Azure K8s Infra Spike
This project is a spike to see if we can programatically generate Custom Resource Definitions (CRDs) for Azure resource types in Kubenetes so that these custom resources can be used to apply and reconcile state into Azure.

## Getting started
This project expects you have [Golang 1.13+](https://golang.org/doc/install), [Docker](https://docs.docker.com/install/), [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest) and [JQ](https://stedolan.github.io/jq/) installed.

### Run Tests and Deploy
The [Makefile](./Makefile) should install all of the required tools needed for running the project. To get up and running, you should need to run one command.
```bash
make deploy-kind
```
#### What does this do?
- Installs all of the build / local dev tools needed (kind, kubectl, controller tools) into `./hack/tools/bin`
  - Always use the bins in `./hack/tools/bin` to interact with the project. Otherwise, nuances in behavior between versions will lead to pain.
- Creates a service principal using Azure CLI and jq
- Builds the controller
- Runs unit tests
- Builds a Kind cluster
  - adds cert-manager to the cluster
  - injects Azure service principal secrets for the controller
- Generates the manifests
- Builds a docker image of the controller
- Pushes the docker image of the controller
- Installs manifests into the Kind cluster
- $$$$

#### Once I'm deployed, then what?
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
