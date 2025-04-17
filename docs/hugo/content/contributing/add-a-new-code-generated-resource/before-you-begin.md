---
title: Before you begin adding a new code-generated resource
linktitle: Before you begin
weight: 10
---

## Determine a resource to add

The first step in adding a new code-generated resource is to determine which resource you want to add.

Any ARM resource can be generated, provided you know the Azure type of the resource, and the version of the API you want to use.

To unambiguously identify the resource you want to add, you need to know the Azure type, and the version of the API you want to use. We will walk through an example of adding Azure Synapse Workspace.
### Find the Azure type of the resource

The Azure type of a resource consists of a resource provider and a name. For example, the Azure type of a Synapse Workspace is `Microsoft.Synapse/workspaces` - the resource provider is `Microsoft.Synapse` and the name is `workspaces`.

If you're not sure of the Azure type for the resource you want to add, one approach is to search for an ARM or Bicep template that deploys the resource. These templates state the Azure type and version of the resource near the top.

E.g. for an ARM template:

``` json
{
  "type": "Microsoft.Synapse/workspaces",
  "apiVersion": "2021-06-01",
  ...
}
```

or for a Bicep template:

``` bicep
resource synapse 'Microsoft.Synapse/workspaces@2021-06-01' = {
    ...
}
```

An alternative is to search for the `ARM` or `Bicep` documentation of the resource. Including the search terms `Azure` `ARM` and `Bicep` in your search will help you find the right documentation. For example, we can search for [`Azure ARM Bicep Synapse Workspaces`](https://www.google.com/search?q=azure+arm+bicep+synapse+workspaces&oq=azure+arm+bicep+synapse+workspaces) and (at least at the time of writing) the first result is [Microsoft.Synapse workspaces](https://learn.microsoft.com/en-us/azure/templates/microsoft.synapse/workspaces?pivots=deployment-language-arm-template).

Another approach is to check if it is defined in a `resource-manager` folder in the [Azure REST API specs](https://github.com/Azure/azure-rest-api-specs/tree/main/specification) repo. (Though, sometimes this can be difficult because of the deep folder structure.) This repo is the authorative source for Azure resources and versions, and is used by the Azure CLI, Azure SDKs, and other tools.

### Select a version

The Azure version of a resource is known as `api-version`, and is usually a date, sometimes with a `-preview` suffix. For example, available versions for a Synapse Workspace are `2021-06-01`, `2021-04-01-preview` and `2020-12-01`.

It is _strongly_ recommended that you use the latest available non-preview `api-version` when choosing the version of the resource to add. This is usually the first version listed in the documentation. Sometimes, a feature you want to have is only available in a preview version, in which case you can use that version.

In the Azure documentation, available versions are shown by the `API Versions` selector; in the `azure-rest-api-specs` repo, they found as sibling folders under either `stable` or `preview` parent directories.

### Define the GVK

Kubernetes resources are identified by their _**group**_, _**version**_, and _**kind**_ (aka GVK).

For the resource you're adding, the GVK is derived from the Azure type and version, as follows:

* **group**: Based on the resource provider, with the `Microsoft.` prefix removed, and lowercased. For example, `Microsoft.Synapse` becomes `synapse`.
* **version**: The constant prefix `v1api` followed by the Azure `api-version`, with non-alphanumeric characters removed. For example, `2021-06-01` becomes `v1api20210601` and `2021-04-01-preview` becomes `v1api20210401preview`.
* **kind**: The Azure type, with the resource provider prefix removed and the plural form converted to singular. For example, `Microsoft.Synapse/workspaces` becomes `workspace`.

To illustrate, here are some examples of GVKs for common resources:

| Azure Type                                 | API Version        | Group            | Version              | Kind           |
| ------------------------------------------ | ------------------ | ---------------- | -------------------- | -------------- |
| Microsoft.Synapse/workspaces               | 2021-06-01         | synapse          | v1api20210601        | workspace      |
| Microsoft.Storage/storageAccounts          | 2021-04-01         | storage          | v1api20210401        | storageaccount |
| Microsoft.Storage/storageAccounts          | 2021-04-01-preview | storage          | v1api20210401preview | storageaccount |
| Microsoft.ContainerService/managedClusters | 2021-03-01         | containerservice | v1api20210301        | managedcluster |

## Development Environment

With the resource identified, you should ensure your development environment is set up to add the resource. (See our [developer-setup]({{< relref "developer-setup" >}}) docs for more detailed information.)

While it _**is**_ possible to do this using just `go`, you end up doing a lot of manual work that we've automated using `task` and other tools.

The recommended approaches are

* Use our .devcontainer, either on [Linux]({{< relref "developer-setup" >}}#dev-container-with-vs-code-on-linux) or [using WSL on Windows]({{< relref "developer-setup" >}}#dev-container-with-vs-code-on-windows)
* Install the tools locally on your machine, using the CLI under either [MacOS]({{< relref "developer-setup" >}}#cli-on-macos) or [Linux]({{< relref "developer-setup" >}}#cli-on-linux).

The [Linux CLI]({{< relref "developer-setup" >}}#cli-on-linux) instructions also work for WSL on Windows.

## Check for Repository Tags

Once you've forked and cloned the ASO repository, make sure you have cloned the tags by running this command:

```bash
git tag --list 'v2*'
```

If you don't see see a list of tags, pull them from your upstream repo

```bash
git fetch --all --tags
```

and then check again.

_**Your builds will fail if you don't have the tags!**_

## Initialize the submodule

ASO references the [`azure-rest-api-specs`](https://github.com/Azure/azure-rest-api-specs) repo for the official Azure API definitions for Azure Resource Providers. _**If this submodule is missing, the code generator will not run**_.

From the root of your ASO repo clone, initialize the submodule with

```bash
git submodule init
```

and then checkout the submodule with

```bash
git submodule update
```

{{% alert title="Note" %}}
If you find you need to update the `azure-rest-api-specs` submodule to a later version as a part of creating your new resource, please create a separate PR for that. These updates often pull in documentation changes to existing resources and keeping those changes separate makes it easier to review and merge the updates.
{{% /alert %}}

----

Now that you have the resource identified and your development environment set up, we can begin by [running the code generator]({{< relref "run-the-code-generator" >}}) to generate the code for the new resource.
