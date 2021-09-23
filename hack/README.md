# Azure Service Operator v2

## Project Status
This project is an alpha. We follow the [Kubernetes definition of alpha](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/#feature-stages).

## Why use Azure Service Operator v2?
- **K8s Native:** we provide CRDs and Golang API structures to deploy and manage Azure resources through Kubernetes.
- **Azure Native:** our CRDs understand Azure resource lifecycle and model it using K8s garbage collection via ownership references.
- **Cloud Scale:** we generate K8s CRDs from Azure Resource Manager schemas to move as fast as Azure.
- **Async Reconciliation:** we don't block on resource creation.

## What resources does ASO v2 support?

See the list of supported resources [here](/v2/api/resources.md).

Sample YAMLs for creating each of these resources can be found in the [samples directory](/hack/generated/config/samples).

## Getting Started
### Prerequisites
1. A Kubernetes cluster (at least version 1.16) [created and running](https://kubernetes.io/docs/tutorials/kubernetes-basics/create-cluster/). You can check your cluster version with `kubectl version`. If you want to try it out quickly, spin up a local cluster using [Kind](https://kind.sigs.k8s.io).
2. An Azure Subscription to provision resources into.
3. An Azure Service Principal for the operator to use, or the [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/?view=azure-cli-latest) to create one. How to create a Service Principal is covered in [installation](#installation).

### Installation
1. Install [cert-manager](https://cert-manager.io/docs/installation/kubernetes/) on the cluster using the following command.

    ```bash
    kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.1.0/cert-manager.yaml
    ```

2. Create an Azure Service Principal. You'll need this to grant Azure Service Operator permissions to create resources in your subscription.
   
   First, set the following environment variables to your Azure Tenant ID and Subscription ID with your values:
   ```yaml
   AZURE_TENANT_ID=<your-tenant-id-goes-here>
   AZURE_SUBSCRIPTION_ID=<your-subscription-id-goes-here>
   ```

   You can find these values by using the Azure CLI: `az account show`

   Next, create a service principal with Contributor permissions for your subscription.

   ```bash
   az ad sp create-for-rbac -n "azure-service-operator" --role contributor \
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
3. Download [the latest **v2+** release](https://github.com/Azure/azure-service-operator/releases) of Azure Service Operator and install it into your cluster.
   ```bash
   kubectl apply -f azureserviceoperator_v2.0.0-alpha.0.yaml
   ```
4. Create the Azure Service Operator v2 secret. This secret contains the identity that Azure Service Operator will run as. Make sure that you have the 4 environment variables from step 2 set before running this command:
   ```bash
   cat <<EOF | kubectl apply -f -
   apiVersion: v1
   kind: Secret
   metadata:
     name: aso-controller-settings
     namespace: azureoperator-system
   stringData:
     AZURE_SUBSCRIPTION_ID: "$AZURE_SUBSCRIPTION_ID"
     AZURE_TENANT_ID: "$AZURE_TENANT_ID"
     AZURE_CLIENT_ID: "$AZURE_CLIENT_ID"
     AZURE_CLIENT_SECRET: "$AZURE_CLIENT_SECRET"
   EOF
   ```

### Usage

Once the controller has been installed in your cluster, you should be able to run the following:

```bash
$ kubectl get pods -n azureoperator-system
NAME                                                READY   STATUS    RESTARTS   AGE
azureoperator-controller-manager-5b4bfc59df-lfpqf   2/2     Running   0          24s

# check out the logs for the running controller
$ kubectl logs -n azureoperator-system azureoperator-controller-manager-5b4bfc59df-lfpqf manager 

# let's create an Azure ResourceGroup in westcentralus with the name "aso-sample-rg"
cat <<EOF | kubectl apply -f -
apiVersion: microsoft.resources.azure.com/v1alpha1api20200601
kind: ResourceGroup
metadata:
  name: aso-sample-rg
  namespace: default
spec:
  location: westcentralus
EOF
# resourcegroup.microsoft.resources.azure.com/aso-sample-rg created

# let's see what the ResourceGroup resource looks like
$ kubectl describe resourcegroups/aso-sample-rg
Name:         aso-sample-rg
Namespace:    default
Labels:       <none>
Annotations:  resource-id.azure.com: /subscriptions/82acd5bb-4206-47d4-9c12-a65db028483d/resourceGroups/aso-sample-rg
              resource-sig.azure.com: 1e3a37c42f6beadbe23d53cf0d271f02d2805d6e295a7e13d5f07bda1fc5b800
API Version:  microsoft.resources.azure.com/v1alpha1api20200601
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
  Id:                      /subscriptions/82acd5bb-4206-47d4-9c12-a65db028483d/resourceGroups/aso-sample-rg
  Location:                westcentralus
  Name:                    aso-sample-rg
  Provisioning State:      Succeeded
Events:
  Type    Reason             Age   From                     Message
  ----    ------             ----  ----                     -------
  Normal  BeginDeployment    32s   ResourceGroupController  Created new deployment to Azure with ID "/subscriptions/82acd5bb-4206-47d4-9c12-a65db028483d/providers/Microsoft.Resources/deployments/k8s_1629763146_19a8f8c2-046e-11ec-8e54-3eec50af7c79"
  Normal  MonitorDeployment  32s   ResourceGroupController  Monitoring Azure deployment ID="/subscriptions/82acd5bb-4206-47d4-9c12-a65db028483d/providers/Microsoft.Resources/deployments/k8s_1629763146_19a8f8c2-046e-11ec-8e54-3eec50af7c79", state="Accepted"
  Normal  MonitorDeployment  27s   ResourceGroupController  Monitoring Azure deployment ID="/subscriptions/82acd5bb-4206-47d4-9c12-a65db028483d/providers/Microsoft.Resources/deployments/k8s_1629763146_19a8f8c2-046e-11ec-8e54-3eec50af7c79", state="Succeeded"


# delete the ResourceGroup
$ kubectl delete resourcegroups/aso-sample-rg
# resourcegroup.microsoft.resources.azure.com "aso-sample-rg" deleted
```

For samples of additional resources, see the [resource samples directory](./generated/config/samples).

### Tearing it down
**Warning: if you `kubectl delete` an Azure resource, it will delete the Azure resource. This can
be dangerous if you were to do this with an existing resource group which contains resources you do
not wish to be deleted.**

If you want to delete the resources you've created, just `kubectl delete` each of the Azure
resources.

As for deleting controller components, just `kubectl delete -f` the release manifests you created
to get started. For example, creating and deleting cert-manager.
```bash
# remove the cert-manager components
kubectl delete -f https://github.com/jetstack/cert-manager/releases/download/v1.1.0/cert-manager.yaml
```

## How to contribute
To get started developing or contributing to the project, follow the instructions in the [contributing guide](./CONTRIBUTING.md).
