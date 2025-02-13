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

## Contact us

If you've got a question, a problem, a request, or just want to chat, here are two ways to get in touch:

* Log an issue on our [GitHub repository](https://github.com/azure/azure-service-operator).
* Join us over on the [Kubernetes Slack](https://kubernetes.slack.com) in the [#azure-service-operator](https://kubernetes.slack.com/archives/C046DEVLAQM) channel. (Don't have an account? [Sign up here](https://slack.k8s.io/).)

## Getting Started

### Prerequisites

1. A Kubernetes cluster (at least version 1.16) [created and running](https://kubernetes.io/docs/tutorials/kubernetes-basics/create-cluster/). You can check your cluster version with `kubectl version`. If you want to try it out quickly, spin up a local cluster using [Kind](https://kind.sigs.k8s.io).
2. An Azure Subscription to provision resources into.
3. An Azure Service Principal for the operator to use, or the [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/?view=azure-cli-latest) to create one. How to create a Service Principal is covered in [installation](#installation).
   See the [Azure Workload Identity]( {{< relref "credential-format#managed-identity-via-workload-identity" >}} ) setup for how to use managed identity instead. We recommend using workload identity in production.

### Installation

#### Install cert-manager on the cluster

See [cert-manager](https://cert-manager.io/docs/installation/kubernetes/) if you'd like to learn more about the project.

``` bash
$ kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.14.1/cert-manager.yaml
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

#### Install the latest **v2+** Helm chart

The latest v2+ Helm chart can be found [here](https://github.com/Azure/azure-service-operator/tree/main/v2/charts)

{{< tabpane text=true left=true >}}
{{% tab header="**Shell**:" disabled=true /%}}
{{% tab header="bash" %}}
``` bash
$ helm repo add aso2 https://raw.githubusercontent.com/Azure/azure-service-operator/main/v2/charts
$ helm upgrade --install aso2 aso2/azure-service-operator \
    --create-namespace \
    --namespace=azureserviceoperator-system \
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
    --set crdPattern=resources.azure.com/*;containerservice.azure.com/*;keyvault.azure.com/*;managedidentity.azure.com/*;eventhub.azure.com/*
```
{{% /tab %}}
{{% tab header="CMD" %}}
``` cmd
C:\> helm repo add aso2 https://raw.githubusercontent.com/Azure/azure-service-operator/main/v2/charts
C:\> helm upgrade --install aso2 aso2/azure-service-operator ^
    --create-namespace ^
    --namespace=azureserviceoperator-system ^
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

{{% alert title="Note" %}}
When installing ASO in an AKS cluster, see [AKS Best Practices]({{< relref "best-practices#aks-best-practices" >}}).
{{% /alert %}}

Alternatively you can install from the [release YAML directly](https://azure.github.io/azure-service-operator/guide/installing-from-yaml/).

#### Create a Managed Identity or Service Principal

This identity or service principal will be used by ASO to authenticate with Azure.
You'll need this to grant the identity or Service Principal permissions to create resources in your subscription.

{{% alert title="Note" %}}
We show steps for using a Service Principal below, as it's easiest to get started with, but recommend using a 
Managed Identity with [Azure Workload Identity]( {{< relref "credential-format#managed-identity-via-workload-identity" >}} ) for
use-cases other than testing.

See [Security best practices]({{< relref "security" >}}) for the full list of security best practices.
{{% /alert %}}

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

#### Create the Azure Service Operator namespaced secret

The secret must be named `aso-credential` and be created in the namespace you'd like to create ASO resources in.

{{% alert title="Note" %}}
To learn about the ASO global secret, which works across all namespaces, or per-resource secrets, see
[authentication documentation](https://azure.github.io/azure-service-operator/guide/authentication/).
{{% /alert %}}

{{< tabpane text=true left=true >}}
{{% tab header="**Shell**:" disabled=true /%}}
{{% tab header="bash" %}}
```bash
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
 name: aso-credential
 namespace: default
stringData:
 AZURE_SUBSCRIPTION_ID: "$AZURE_SUBSCRIPTION_ID"
 AZURE_TENANT_ID: "$AZURE_TENANT_ID"
 AZURE_CLIENT_ID: "$AZURE_CLIENT_ID"
 AZURE_CLIENT_SECRET: "$AZURE_CLIENT_SECRET"
EOF
```
{{% /tab %}}

{{% tab header="PowerShell" %}}
```powershell
@"
apiVersion: v1
kind: Secret
metadata:
 name: aso-credential
 namespace: default
stringData:
 AZURE_SUBSCRIPTION_ID: "$AZURE_SUBSCRIPTION_ID"
 AZURE_TENANT_ID: "$AZURE_TENANT_ID"
 AZURE_CLIENT_ID: "$AZURE_CLIENT_ID"
 AZURE_CLIENT_SECRET: "$AZURE_CLIENT_SECRET"
"@ | kubectl apply -f -
```
{{% /tab %}}

{{% tab header="CMD" %}}

Create a file named `secret.yaml` with the following content. Replace each of the variables such as
`%AZURE_SUBSCRIPTION_ID%` with the subscription ID, `%AZURE_CLIENT_SECRET%` with the client secret and so on.

```
apiVersion: v1
kind: Secret
metadata:
 name: aso-credential
 namespace: default
stringData:
 AZURE_SUBSCRIPTION_ID: "%AZURE_SUBSCRIPTION_ID%"
 AZURE_TENANT_ID: "%AZURE_TENANT_ID%"
 AZURE_CLIENT_ID: "%AZURE_CLIENT_ID%"
 AZURE_CLIENT_SECRET: "%AZURE_CLIENT_SECRET%"
```

Then run: `kubectl apply -f secret.yaml`

{{% /tab %}}
{{< /tabpane >}}

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

Let's create an Azure ResourceGroup in westcentralus with the name "aso-sample-rg". Create a file called `rg.yaml` 
with the following contents:

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
$ kubectl describe resourcegroups/aso-sample-rg -n default
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
$ kubectl delete -f https://github.com/jetstack/cert-manager/releases/download/v1.14.4/cert-manager.yaml
```

## How to contribute

To get started developing or contributing to the project, follow the instructions in the [contributing guide](https://azure.github.io/azure-service-operator/contributing/).
