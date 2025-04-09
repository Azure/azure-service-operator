---
title: Before you begin adding a new code-generated resource
linktitle: Before you begin
weight: 10
---

## Determine a resource to add

The first step in adding a new code-generated resource is to determine which resource you want to add. 

Any ARM resource can be generated. If you're not sure if the resource you are interested in is an ARM resource, check if it is defined in a `resource-manager` folder in the [Azure REST API specs](https://github.com/Azure/azure-rest-api-specs/tree/main/specification) repo.

If it is, it's an ARM resource.

To unambigously identify the resource you want to add, you need to know its _**group**_, _**version**_, and _**kind**_ (aka GVK).

We work out these by reviewing the [Azure REST API specs](https://github.com/Azure/azure-rest-api-specs/tree/main/specification) repo. We will walk through an example of adding Azure Synapse Workspace.

### Kind

Also known as the _name_ of the resource, you usually know this going in. In our example above, the name of the resource is `workspace`. 

If you're not sure, one approach is to look in the Swagger/OpenAPI specification file for the service and find the documented PUT for the resource you're interested in. The resource name will be the
[second to last section of the URL](https://github.com/Azure/azure-rest-api-specs/blob/main/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/workspace.json#L71).

Another approach is to search for the ARM documentation of the resource. Including the search terms `Azure` `ARM` and `Bicep` will help you find the right documentation. For example, we can search for [`Azure ARM Bicep Synapse Workspaces`](https://www.google.com/search?q=azure+arm+bicep+synapse+workspaces&oq=azure+arm+bicep+synapse+workspaces) and (at least at the time of writing) the first result is the [Azure Synapse Workspace](https://learn.microsoft.com/en-us/azure/templates/microsoft.synapse/workspaces?pivots=deployment-language-arm-template) resource.

**Note**: ARM almost exclusively uses plural names for resources. ASO uses singular forms, aligning with the usual practices in Kubernetes. For example, the Azure Synapse Workspace resource is called `workspaces` in ARM and `workspace` in ASO.

### Group

This is named after the Azure service, for example `resources` or `documentdb`.

In our example, this is `synapse` (from `Microsoft.Synapse`, the provider documented in the resource URL).

### Version

The version is the `api-version` of the resource. This is usually a date, sometimes with a `-preview` suffix. 

In our example entry from above, this is `2021-06-01`.

**Note**: In most cases there will be multiple API versions for a given resource, you can see this in the
[Synapse](https://github.com/Azure/azure-rest-api-specs/tree/main/specification/synapse/resource-manager/Microsoft.Synapse/stable)
folder for example, there are 4 API versions as of April 2023. It is _strongly_ recommended that you use the latest
available non-preview `api-version` when choosing the version of the resource to add.

## Development Environment

With the resource identified, you should ensure your [development environment]({{< relref "developer-setup" >}}) is set up to add the resource.

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

``` bash
git fetch --all --tags
```

and then check again.

_**Your builds will fail if you don't have the tags!**_

## Initialize the submodule

ASO references the [`azure-rest-api-specs`](https://github.com/Azure/azure-rest-api-specs) repo for the official Azure API definitions for Azure Resource Providers. _**If this submodule is missing, the code generator will not run**_.

From the root of your ASO repo clone, initialize the submodule with

``` bash
git submodule init
```

and then checkout the submodule with

``` bash
git submodule update
```

**Tip**: If you find you need to update the submodule to a later version as a part of creating your new resource, please create a separate PR for that. These updates often pull in documentation changes to existing resources and keeping those changes separate makes it easier to review and merge the updates.

----

Now that you have the resource identified and your development environment set up, we can begin by [running the code generator]({{< relref "run-the-code-generator" >}}) to generate the code for the new resource.