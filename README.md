# Azure K8s Infra (Experimental)
Kubernetes-native infrastructure provider for Azure at scale

## What is Azure K8s Infra?
Azure K8s Infra is a Kubernetes controller, API, and Custom Resource Definition (CRD) code generator, 
which enables K8s users to build any Azure resource.

### What makes K8s Infra different?
- **K8s Native:** we provide strong structured schema and Golang API structures
- **Azure Native:** our CRDs understand Azure resource lifecycle and model it using K8s garbage collection via ownership references
- **Cloud Scale:** we generate K8s CRDs from Azure Resource Manager schemas to move as fast as Azure
- **Async Reconciliation:** we don't block on resource creation

## Project Status
This project is experimental and is not supported.

### Project Goals
- Create a scalable process for exposing Azure resources natively in Kubernetes
- Provide a prototype infrastructure reconciler for [Cluster API Provider Azure](https://github.com/kubernetes-sigs/cluster-api-provider-azure)
- Be a research and development space for improving Azure user experience in Kubernetes

### Project Non-Goals
- Long term support

## Getting Started
To get started you are going to need a cluster to deploy k8s-infra. You can use any 1.16.0+ K8s 
cluster. To get going quicker, just spin up a local cluster using [Kind](https://kind.sigs.k8s.io).

To get started using the Azure k8s-infra infrastructure provider, visit our [releases](https://github.com/Azure/k8s-infra/releases),
and follow the instructions for the latest release.

Once the controller has been installed in your cluster, you should be able to run the following:
```bash
$ k get pod -A
NAMESPACE            NAME                                             READY   STATUS    RESTARTS   AGE
k8s-infra-system     k8s-infra-controller-manager-b98bc664-6h6sv      2/2     Running   0          7m15s

# check out the logs for the running controller
$ k logs pod/k8s-infra-controller-manager-b98bc664-6h6sv -n k8s-infra-system -c manager

# let's create an Azure ResourceGroup in westus2 with the name "foo-2019"
cat <<EOF | kubectl apply -f -
apiVersion: microsoft.resources.infra.azure.com/v20191001
kind: ResourceGroup
metadata:
  name: foo-2019
spec:
  location: westus2
EOF
# resourcegroup.microsoft.resources.infra.azure.com/foo-2019 created

# let's see what the ResourceGroup resource looks like
$ k describe resourcegroups/foo-2019
Name:         foo-2019
Namespace:    default
Labels:       <none>
Annotations:  resource-sig.infra.azure.com: 56da597b51d66c934c510b22e183a69994a6654bbae92d88fdb46f962272220a
API Version:  microsoft.resources.infra.azure.com/v20191001
Kind:         ResourceGroup
Metadata:
  Creation Timestamp:  2020-05-13T22:55:56Z
  Finalizers:
    infra.azure.com/finalizer
  Generation:  1
Spec:
  Location:  westus2
Status:
  Id:                  /subscriptions/8ec81d93-7715-4cec-8dcf-e07583d8a24a/resourceGroups/foo-2019
  Provisioning State:  Succeeded
Events:
  Type    Reason                    Age                From               Message
  ----    ------                    ----               ----               -------
  Normal  ResourceHasChanged        106s               ResourceGroupCtrl  resource in state "" has changed and spec will be applied to Azure
  Normal  ResourceStateNonTerminal  101s               ResourceGroupCtrl  resource in state "", asking Azure for updated state
  Normal  ResourceStateNonTerminal  98s (x2 over 99s)  ResourceGroupCtrl  resource in state "Accepted", asking Azure for updated state
  Normal  ResourceStateNonTerminal  96s (x3 over 99s)  ResourceGroupCtrl  resource in state "Running", asking Azure for updated state
  Normal  ResourceHasNotChanged

# delete the ResourceGroup
$ k delete resourcegroups/foo-2019
# resourcegroup.microsoft.resources.infra.azure.com "foo-2019" deleted
```

For samples of additional resources, see the [resource samples directory in the repo](https://github.com/Azure/k8s-infra/tree/master/config/samples).
We currently provide only a subset of resources. As work on the [code generator](https://github.com/Azure/k8s-infra/tree/master/hack/generator)
progresses, the resources provided will expand considerably.

### Tearing it down
**Warning: if you `kubectl delete` an Azure resource, it will delete the Azure resource. This can
be dangerous if you were to do this with an existing resource group which contains resources you do
not wish to be deleted.**

If you want to delete the resources you've created, just `kubectl delete` each of the Azure 
resources.

As for deleting controller components, just `kubectl delete -f` the release manifests you created 
to get started. For example, creating and deleting cert-manager.
```bash
# apply the cert-manager components
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v0.13.0/cert-manager.yaml

# remove the cert-manager components
kubectl delete -f https://github.com/jetstack/cert-manager/releases/download/v0.13.0/cert-manager.yaml
```

## Developing
To get started developing in the project, follow the instructions in the [developer readme](./docs/development.md).

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
