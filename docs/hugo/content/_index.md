---
title: Azure Service Operator v2
type: docs
description: "Manage your Azure resources from within your Kubernetes cluster."
---

<img src="https://azure.github.io/azure-service-operator/favicons/favicon-128.png" style="float:left; margin: -8px 8px 8px 0px;"/>Azure Service Operator (ASO) allows you to deploy and maintain a wide variety of Azure Resources using the Kubernetes tooling you already know and use.

Instead of deploying and managing your Azure resources separately from your Kubernetes application, ASO allows you to manage them together, automatically configuring your application as needed. For example, ASO can set up your [Redis Cache](https://azure.github.io/azure-service-operator/reference/cache/) or [PostgreSQL database server](https://azure.github.io/azure-service-operator/reference/dbforpostgresql/) and then configure your Kubernetes application to use them.

## Project Status

This project is stable. We follow the [Kubernetes definition of stable](https://kubernetes.io/docs/reference/command-line-tools-reference/feature-gates/#feature-stages).

## Why use Azure Service Operator v2?

- **K8s Native:** we provide CRDs and Golang API structures to deploy and manage Azure resources through Kubernetes.
- **Azure Native:** our CRDs understand Azure resource lifecycle and model it using K8s garbage collection via ownership references.
- **Cloud Scale:** we generate K8s CRDs from Azure Resource Manager schemas to move as fast as Azure.
- **Async Reconciliation:** we don't block on resource creation.

## What resources does ASO v2 support?

ASO supports more than 150 different Azure resources, with more added every release. See the full list of [supported resources](https://azure.github.io/azure-service-operator/reference/).

## Getting Started

### Prerequisites

1. A Kubernetes cluster (at least version 1.16) [created and running](https://kubernetes.io/docs/tutorials/kubernetes-basics/create-cluster/). You can check your cluster version with `kubectl version`. If you want to try it out quickly, spin up a local cluster using [Kind](https://kind.sigs.k8s.io).
2. An Azure Subscription to provision resources into.
3. An Azure Service Principal for the operator to use, or the [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/?view=azure-cli-latest) to create one. How to create a Service Principal is covered in [installation](#installation).
   See the [Azure Workload Identity](https://azure.github.io/azure-service-operator/guide/authentication/credential-format/#azure-workload-identity) setup for how to use managed identity instead. We recommend using workload identity in production.

### Installation

1. Install [cert-manager](https://cert-manager.io/docs/installation/kubernetes/) on the cluster using the following command.

``` bash
$ kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.12.1/cert-manager.yaml
```

Check that the cert-manager pods have started successfully before continuing.

``` bash
$ kubectl get pods -n cert-manager
NAME                                      READY   STATUS    RESTARTS   AGE
cert-manager-5597cff495-lmphj             1/1     Running   0          1m
cert-manager-cainjector-bd5f9c764-gvxm4   1/1     Running   0          1m
cert-manager-webhook-c4b5687dc-x66bj      1/1     Running   0          1m
```

(Alternatively, you can wait for cert-manager to be ready with `cmctl check api --wait=2m` - see the [cert-manager documentation](https://cert-manager.io/docs/usage/cmctl/) for more information about `cmctl`.)

2. Create an Azure Service Principal. You'll need this to grant Azure Service Operator permissions to create resources in your subscription.

   First, set the following environment variables to your Azure Tenant ID and Subscription ID with your values:

{{< tabpane text=true left=true >}}
{{% tab header="**Shell**:" disabled=true /%}}
{{% tab header="bash" %}}
``` bash
$ export AZURE_TENANT_ID=<your-tenant-id-goes-here>
$ export AZURE_SUBSCRIPTION_ID=<your-subscription-id-goes-here>
```
{{% /tab %}}
{{% tab header="PowerShell" %}}
``` powershell
PS> $AZURE_TENANT_ID=<your-tenant-id-goes-here>
PS> $AZURE_SUBSCRIPTION_ID=<your-subscription-id-goes-here>
```
{{% /tab %}}
{{% tab header="CMD" %}}
``` cmd
C:\> SET AZURE_TENANT_ID=<your-tenant-id-goes-here>
C:\> SET AZURE_SUBSCRIPTION_ID=<your-subscription-id-goes-here>
```
{{% /tab %}}
{{< /tabpane >}}

   You can find these values by using the Azure CLI: `az account show`

   Next, create a service principal with Contributor permissions for your subscription.

   You can optionally use a service principal with a more restricted permission set
   (for example contributor to just a Resource Group), but that will restrict what you can
   do with ASO. See [using reduced permissions](https://azure.github.io/azure-service-operator/guide/authentication/reducing-access/#using-a-credential-for-aso-with-reduced-permissions) for more details.

{{< tabpane text=true left=true >}}
{{% tab header="**Shell**:" disabled=true /%}}
{{% tab header="bash" %}}
``` bash
$ az ad sp create-for-rbac -n azure-service-operator --role contributor \
    --scopes /subscriptions/$AZURE_SUBSCRIPTION_ID
```
{{% /tab %}}
{{% tab header="PowerShell" %}}
``` powershell
PS> az ad sp create-for-rbac -n azure-service-operator --role contributor `
    --scopes /subscriptions/$AZURE_SUBSCRIPTION_ID
```
{{% /tab %}}
{{% tab header="CMD" %}}
``` cmd
C:\> az ad sp create-for-rbac -n azure-service-operator --role contributor ^
    --scopes /subscriptions/%AZURE_SUBSCRIPTION_ID%
```
{{% /tab %}}
{{< /tabpane >}}

This should give you output including the following:

```bash
"appId": "xxxxxxxxxx",
"displayName": "azure-service-operator",
"name": "http://azure-service-operator",
"password": "xxxxxxxxxxx",
"tenant": "xxxxxxxxxxxxx"
```

Once you have created a service principal, set the following variables to your app ID and password values:

{{< tabpane text=true left=true >}}
{{% tab header="**Shell**:" disabled=true /%}}
{{% tab header="bash" %}}
``` bash
$ export AZURE_CLIENT_ID=<your-client-id>         # This is the appID from the service principal we created.
$ export AZURE_CLIENT_SECRET=<your-client-secret> # This is the password from the service principal we created.
```
{{% /tab %}}
{{% tab header="PowerShell" %}}
``` powershell
PS> $AZURE_CLIENT_ID=<your-client-id>         # This is the appID from the service principal we created.
PS> $AZURE_CLIENT_SECRET=<your-client-secret> # This is the password from the service principal we created.
```
{{% /tab %}}
{{% tab header="CMD" %}}
``` cmd
C:\> :: This is the appID from the service principal we created.
C:\> SET AZURE_CLIENT_ID=<your-client-id>         
C:\> :: This is the password from the service principal we created.
C:\> SET AZURE_CLIENT_SECRET=<your-client-secret> 
```
{{% /tab %}}
{{< /tabpane >}}

3. Install [the latest **v2+** Helm chart](https://github.com/Azure/azure-service-operator/tree/main/v2/charts):

{{< tabpane text=true left=true >}}
{{% tab header="**Shell**:" disabled=true /%}}
{{% tab header="bash" %}}
``` bash
$ helm repo add aso2 https://raw.githubusercontent.com/Azure/azure-service-operator/main/v2/charts
$ helm upgrade --install aso2 aso2/azure-service-operator \
    --create-namespace \
    --namespace=azureserviceoperator-system \
    --set azureSubscriptionID=$AZURE_SUBSCRIPTION_ID \
    --set azureTenantID=$AZURE_TENANT_ID \
    --set azureClientID=$AZURE_CLIENT_ID \
    --set azureClientSecret=$AZURE_CLIENT_SECRET \
    --set crdPattern='resources.azure.com/*;containerservice.azure.com/*;keyvault.azure.com/*;managedidentity.azure.com/*;eventhub.azure.com/*'
```
Note: **bash** requires the value for `crdPattern` to be quoted with `'` to avoid expansion of the wildcards.
{{% /tab %}}
{{% tab header="PowerShell" %}}
``` powershell
PS> helm repo add aso2 https://raw.githubusercontent.com/Azure/azure-service-operator/main/v2/charts
PS> helm upgrade --install aso2 aso2/azure-service-operator `
    --create-namespace `
    --namespace=azureserviceoperator-system `
    --set azureSubscriptionID=$AZURE_SUBSCRIPTION_ID `
    --set azureTenantID=$AZURE_TENANT_ID `
    --set azureClientID=$AZURE_CLIENT_ID `
    --set azureClientSecret=$AZURE_CLIENT_SECRET `
    --set crdPattern=resources.azure.com/*;containerservice.azure.com/*;keyvault.azure.com/*;managedidentity.azure.com/*;eventhub.azure.com/*
```
{{% /tab %}}
{{% tab header="CMD" %}}
``` cmd
C:\> helm repo add aso2 https://raw.githubusercontent.com/Azure/azure-service-operator/main/v2/charts
C:\> helm upgrade --install aso2 aso2/azure-service-operator ^
    --create-namespace ^
    --namespace=azureserviceoperator-system ^
    --set azureSubscriptionID=%AZURE_SUBSCRIPTION_ID% ^
    --set azureTenantID=%AZURE_TENANT_ID% ^
    --set azureClientID=%AZURE_CLIENT_ID% ^
    --set azureClientSecret=%AZURE_CLIENT_SECRET% ^
    --set crdPattern=resources.azure.com/*;containerservice.azure.com/*;keyvault.azure.com/*;managedidentity.azure.com/*;eventhub.azure.com/*
```
{{% /tab %}}
{{< /tabpane >}}

{{% alert title="Warning" color="warning" %}}
Make sure to set the `crdPattern` variable to include the CRDs you are interested in using.
You can use `--set crdPattern=*` to install all the CRDs, but be aware of the
[limits of the Kubernetes you are running](https://github.com/Azure/azure-service-operator/issues/2920). `*` is **not**
recommended on AKS Free-tier clusters.
See [CRD management](https://azure.github.io/azure-service-operator/guide/crd-management/) for more details.
{{% /alert %}}

   Alternatively you can install from the [release YAML directly](https://azure.github.io/azure-service-operator/guide/installing-from-yaml/).

   To learn more about other authentication options, see the [authentication documentation](https://azure.github.io/azure-service-operator/guide/authentication/).

### Usage

Once the controller has been installed in your cluster, you should be able to see the ASO pod running.

``` bash
$ kubectl get pods -n azureserviceoperator-system
NAME                                                READY   STATUS    RESTARTS   AGE
azureserviceoperator-controller-manager-5b4bfc59df-lfpqf   2/2     Running   0          24s
```

To view the logs for the running ASO controller, take note of the pod name shown above and then run the following command.

``` bash
$ kubectl logs -n azureserviceoperator-system azureserviceoperator-controller-manager-5b4bfc59df-lfpqf manager 
```

Let's create an Azure ResourceGroup in westcentralus with the name "aso-sample-rg". Create a file called `rg.yaml` with the following contents:

```yaml
apiVersion: resources.azure.com/v1api20200601
kind: ResourceGroup
metadata:
  name: aso-sample-rg
  namespace: default
spec:
  location: westcentralus
````

Then apply the file to your cluster:

``` bash
$ kubectl apply -f rg.yaml
# resourcegroup.resources.azure.com/aso-sample-rg created
```

Once the resource group has been created, we can see what it looks like.

``` bash
$ kubectl describe resourcegroups/aso-sample-rg
```

The output will be similar to this:
``` yaml
Name:         aso-sample-rg
Namespace:    default
Labels:       <none>
Annotations:  serviceoperator.azure.com/operator-namespace: azureserviceoperator-system
              serviceoperator.azure.com/resource-id: /subscriptions/82acd5bb-4206-47d4-9c12-a65db028483d/resourceGroups/aso-sample-rg
API Version:  resources.azure.com/v1beta20200601
Kind:         ResourceGroup
Metadata:
  Creation Timestamp:  2023-08-31T01:25:50Z
  Finalizers:
    serviceoperator.azure.com/finalizer
  Generation:        1
  Resource Version:  3198
  UID:               70e2fef1-8c43-452f-8260-ffe5a73470fb
Spec:
  Azure Name:  aso-sample-rg
  Location:    westcentralus
Status:
  Conditions:
    Last Transition Time:  2023-08-31T01:25:58Z
    Observed Generation:   1
    Reason:                Succeeded
    Status:                True
    Type:                  Ready
  Id:                      /subscriptions/82acd5bb-4206-47d4-9c12-a65db028483d/resourceGroups/aso-sample-rg
  Location:                westcentralus
  Name:                    aso-sample-rg
  Properties:
    Provisioning State:  Succeeded
  Type:                  Microsoft.Resources/resourceGroups
Events:
  Type    Reason               Age                From                     Message
  ----    ------               ----               ----                     -------
  Normal  CredentialFrom       42s (x2 over 42s)  ResourceGroupController  Using credential from "azureserviceoperator-system/aso-controller-settings"
  Normal  BeginCreateOrUpdate  35s                ResourceGroupController  Successfully sent resource to Azure with ID "/subscriptions/82acd5bb-4206-47d4-9c12-a65db028483d/resourceGroups/aso-sample-rg"
```

We can delete the resource group from the cluster. This will also delete it from Azure.

``` bash
$ kubectl delete resourcegroups/aso-sample-rg
# resourcegroup.resources.azure.com "aso-sample-rg" deleted
```

For samples of additional resources, see the [resource samples directory](https://github.com/Azure/azure-service-operator/tree/main/v2/samples).

### Tearing it down

{{% alert title="Warning" color="warning" %}}
if you `kubectl delete` an ASO resource from your cluster, ASO will delete the Azure resource. This can
be dangerous if you were to do this with an existing resource group which contains resources you do
not wish to be deleted.
{{% /alert %}}

If you want to delete the resources you've created, just `kubectl delete` each of the Azure
resources.

If you want to delete the cluster resource without affecting the Azure resource, apply the annotation `serviceoperator.azure.com/reconcile-policy: skip` first.

As for deleting controller components, just `kubectl delete -f` the release manifests you created
to get started. For example, creating and deleting cert-manager.

``` bash
# remove the cert-manager components
$ kubectl delete -f https://github.com/jetstack/cert-manager/releases/download/v1.12.1/cert-manager.yaml
```

## How to contribute

To get started developing or contributing to the project, follow the instructions in the [contributing guide](https://azure.github.io/azure-service-operator/contributing/).
