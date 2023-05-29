---
title: Azure Service Operator v2
type: docs
description: "Manage your Azure resources from within your Kubernetes cluster."
---
## Project Status
This project is stable. We follow the [Kubernetes definition of stable](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/#feature-stages).

## Why use Azure Service Operator v2?
- **K8s Native:** we provide CRDs and Golang API structures to deploy and manage Azure resources through Kubernetes.
- **Azure Native:** our CRDs understand Azure resource lifecycle and model it using K8s garbage collection via ownership references.
- **Cloud Scale:** we generate K8s CRDs from Azure Resource Manager schemas to move as fast as Azure.
- **Async Reconciliation:** we don't block on resource creation.

## What resources does ASO v2 support?

See the list of supported resources [here](https://azure.github.io/azure-service-operator/reference/).

## Getting Started
### Prerequisites
1. A Kubernetes cluster (at least version 1.16) [created and running](https://kubernetes.io/docs/tutorials/kubernetes-basics/create-cluster/). You can check your cluster version with `kubectl version`. If you want to try it out quickly, spin up a local cluster using [Kind](https://kind.sigs.k8s.io).
2. An Azure Subscription to provision resources into.
3. An Azure Service Principal for the operator to use, or the [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/?view=azure-cli-latest) to create one. How to create a Service Principal is covered in [installation](#installation).

### Installation
1. Install [cert-manager](https://cert-manager.io/docs/installation/kubernetes/) on the cluster using the following command.

    ```bash
    kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.8.2/cert-manager.yaml
    ```
   Check that the cert-manager pods have started successfully before continuing.

   ```bash
   $ kubectl get pods -n cert-manager
   NAME                                      READY   STATUS    RESTARTS   AGE
   cert-manager-5597cff495-lmphj             1/1     Running   0          1m
   cert-manager-cainjector-bd5f9c764-gvxm4   1/1     Running   0          1m
   cert-manager-webhook-c4b5687dc-x66bj      1/1     Running   0          1m
   ```

   (Alternatively, you can wait for cert-manager to be ready with `cmctl check api --wait=2m` - see the [cert-manager documentation](https://cert-manager.io/docs/usage/cmctl/) for more information about `cmctl`.)

2. Create an Azure Service Principal. You'll need this to grant Azure Service Operator permissions to create resources in your subscription.

   First, set the following environment variables to your Azure Tenant ID and Subscription ID with your values:
   ```yaml
   AZURE_TENANT_ID=<your-tenant-id-goes-here>
   AZURE_SUBSCRIPTION_ID=<your-subscription-id-goes-here>
   ```

   You can find these values by using the Azure CLI: `az account show`

   Next, create a service principal with Contributor permissions for your subscription.
   
   You can optionally use a service principal with a more restricted permission set 
   (for example contributor to just a Resource Group), but that will restrict what you can
   do with ASO. See [using reduced permissions](https://azure.github.io/azure-service-operator/guide/authentication/reducing-access#using-a-credential-for-aso-with-reduced-permissions) for more details.

   ```bash
   az ad sp create-for-rbac -n azure-service-operator --role contributor \
       --scopes /subscriptions/$AZURE_SUBSCRIPTION_ID
   ```

   This should give you output like the following:
   ```bash
   "appId": "xxxxxxxxxx",
   "displayName": "azure-service-operator",
   "name": "http://azure-service-operator",
   "password": "xxxxxxxxxxx",
   "tenant": "xxxxxxxxxxxxx"
   ```

   Once you have created a service principal, set the following variables to your app ID and password values:
   ```bash 
   AZURE_CLIENT_ID=<your-client-id> # This is the appID from the service principal we created.
   AZURE_CLIENT_SECRET=<your-client-secret> # This is the password from the service principal we created.
   ```

3. Install [the latest **v2+** Helm chart](https://github.com/Azure/azure-service-operator/tree/main/charts):
   
   ```
   helm repo add aso2 https://raw.githubusercontent.com/Azure/azure-service-operator/main/v2/charts
   ```

   ```
   helm upgrade --install --devel aso2 aso2/azure-service-operator \
        --create-namespace \
        --namespace=azureserviceoperator-system \
        --set azureSubscriptionID=$AZURE_SUBSCRIPTION_ID \
        --set azureTenantID=$AZURE_TENANT_ID \
        --set azureClientID=$AZURE_CLIENT_ID \
        --set azureClientSecret=$AZURE_CLIENT_SECRET \
        --set crdPattern='resources.azure.com/*;containerservice.azure.com/*;keyvault.azure.com/*;managedidentity.azure.com/*;eventhub.azure.com/*'
   ```

   > **Warning:** Make sure to set the `crdPattern` variable to include the CRDs you are interested in using. 
   > You can use `--set crdPattern=*` to install all the CRDs, but be aware of the 
   > [limits of the Kubernetes you are running](https://github.com/Azure/azure-service-operator/issues/2920). `*` is **not**
   > recommended on AKS Free-tier clusters.
   > 
   > See [CRD management](https://azure.github.io/azure-service-operator/guide/crd-management/) for more details.

   Alternatively you can install from the [release YAML directly](https://azure.github.io/azure-service-operator/guide/installing-from-yaml).

   To learn more about other authentication options, see the [authentication documentation](https://azure.github.io/azure-service-operator/guide/authentication/).

### Usage

Once the controller has been installed in your cluster, you should be able to run the following:

```bash
$ kubectl get pods -n azureserviceoperator-system
NAME                                                READY   STATUS    RESTARTS   AGE
azureserviceoperator-controller-manager-5b4bfc59df-lfpqf   2/2     Running   0          24s

# check out the logs for the running controller
$ kubectl logs -n azureserviceoperator-system azureserviceoperator-controller-manager-5b4bfc59df-lfpqf manager 

# let's create an Azure ResourceGroup in westcentralus with the name "aso-sample-rg"
cat <<EOF | kubectl apply -f -
apiVersion: resources.azure.com/v1alpha1api20200601
kind: ResourceGroup
metadata:
  name: aso-sample-rg
  namespace: default
spec:
  location: westcentralus
EOF
# resourcegroup.resources.azure.com/aso-sample-rg created

# let's see what the ResourceGroup resource looks like
$ kubectl describe resourcegroups/aso-sample-rg
Name:         aso-sample-rg
Namespace:    default
Labels:       <none>
Annotations:  resource-id.azure.com: /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-sample-rg
              resource-sig.azure.com: 1e3a37c42f6beadbe23d53cf0d271f02d2805d6e295a7e13d5f07bda1fc5b800
API Version:  resources.azure.com/v1alpha1api20200601
Kind:         ResourceGroup
Metadata:
  Creation Timestamp:  2021-08-23T23:59:06Z
  Finalizers:
    serviceoperator.azure.com/finalizer
  Generation:  1
Spec:
  Azure Name:  aso-sample-rg
  Location:    westcentralus
Status:
  Conditions:
    Last Transition Time:  2021-08-23T23:59:13Z
    Reason:                Succeeded
    Status:                True
    Type:                  Ready
  Id:                      /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/aso-sample-rg
  Location:                westcentralus
  Name:                    aso-sample-rg
  Provisioning State:      Succeeded
Events:
  Type    Reason             Age   From                     Message
  ----    ------             ----  ----                     -------
  Normal  BeginDeployment    32s   ResourceGroupController  Created new deployment to Azure with ID "/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Resources/deployments/k8s_1629763146_19a8f8c2-046e-11ec-8e54-3eec50af7c79"
  Normal  MonitorDeployment  32s   ResourceGroupController  Monitoring Azure deployment ID="/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Resources/deployments/k8s_1629763146_19a8f8c2-046e-11ec-8e54-3eec50af7c79", state="Accepted"
  Normal  MonitorDeployment  27s   ResourceGroupController  Monitoring Azure deployment ID="/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Resources/deployments/k8s_1629763146_19a8f8c2-046e-11ec-8e54-3eec50af7c79", state="Succeeded"


# delete the ResourceGroup
$ kubectl delete resourcegroups/aso-sample-rg
# resourcegroup.resources.azure.com "aso-sample-rg" deleted
```

For samples of additional resources, see the [resource samples directory](https://github.com/Azure/azure-service-operator/tree/main/v2/samples).

### Tearing it down
**Warning: if you `kubectl delete` an Azure resource, it will delete the Azure resource. This can
be dangerous if you were to do this with an existing resource group which contains resources you do
not wish to be deleted.**

If you want to delete the resources you've created, just `kubectl delete` each of the Azure
resources.

If you want to delete the cluster resource without affecting the Azure resource, apply the annotation `serviceoperator.azure.com/reconcile-policy: skip` first.

As for deleting controller components, just `kubectl delete -f` the release manifests you created
to get started. For example, creating and deleting cert-manager.
```bash
# remove the cert-manager components
kubectl delete -f https://github.com/jetstack/cert-manager/releases/download/v1.8.2/cert-manager.yaml
```

## How to contribute
To get started developing or contributing to the project, follow the instructions in the [contributing guide](https://azure.github.io/azure-service-operator/contributing/).
