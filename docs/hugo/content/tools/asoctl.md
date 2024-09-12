---
title: "asoctl"
linkTitle: "asoctl"
weight: 10
layout: single
cascade:
- type: docs
- render: always
description: Azure Service Operator Controller
---

``` bash
$ asoctl
asoctl provides a cmdline interface for working with Azure Service Operator

Usage:
  asoctl [command]

Available Commands:
  clean       Clean Custom Resource Definitions (CRDs) prior to upgrade
  completion  Generate the autocompletion script for the specified shell
  export      Exports an ASO YAML file from a template
  help        Help about any command
  import      Imports ARM resources to YAML files containing ASO custom resource definitions
  version     Display version information

Flags:
      --quiet     Silence most logging
      --verbose   Enable verbose logging
  -h, --help      help for asoctl

Use "asoctl [command] --help" for more information about a command.
```

## Installation

{{< tabpane text=true left=true >}}
{{% tab header="**OS**:" disabled=true /%}}
{{% tab header="Linux (incl WSL)" %}}

Download for AMD64:

``` bash
$ curl -L https://github.com/Azure/azure-service-operator/releases/latest/download/asoctl-linux-amd64.gz -o asoctl.gz
$ gunzip asoctl.gz
```

Download for ARM64:

``` bash
$ curl -L https://github.com/Azure/azure-service-operator/releases/latest/download/asoctl-linux-arm64.gz -o asoctl.gz
$ gunzip asoctl.gz
```

Install:

``` bash
$ sudo install -o root -g root -m 0755 asoctl /usr/local/bin/asoctl`
```

{{% /tab %}}
{{% tab header="macOS" %}}

Download for AMD64:

``` bash
$ curl -L https://github.com/Azure/azure-service-operator/releases/latest/download/asoctl-darwin-amd64.gz -o asoctl.gz
$ gunzip asoctl.gz
```

Download for ARM64:

``` bash
$ curl -L https://github.com/Azure/azure-service-operator/releases/latest/download/asoctl-darwin-arm64.gz -o asoctl.gz
$ gunzip asoctl.gz
```

Make the binary executable:

``` bash
$ chmod +x ./asoctl
```

Move the binary to your PATH:

``` bash
$ sudo mv ./asoctl /usr/local/bin/asoctl
```

{{% /tab %}}
{{% tab header="Windows" %}}

Using a PowerShell command prompt, download the latest release to your current directory:

``` Powershell
$ curl.exe -L https://github.com/Azure/azure-service-operator/releases/latest/download/asoctl-windows-amd64.zip -o asoctl.zip
$ Expand-Archive asoctl.zip . -Force
```

Append that directory to your `PATH` if desired.

{{% /tab %}}
{{< /tabpane >}}

## Export template

The `export template` commands helps render a YAML file from the raw ASO published with each release.

```bash
Template creates a YAML file from the specified ASO yaml template

Usage:
  asoctl export template [--source <string> |--version <string>] [--crd-pattern <string>|--raw] [flags]

Examples:
asoctl export template --version v2.6.0 --crd-pattern "resources.azure.com/*" --crd-pattern "containerservice.azure.com/*"

With combined crd-pattern:
asoctl export template --version v2.6.0 --crd-pattern "resources.azure.com/*;containerservice.azure.com/*;keyvault.azure.com/*;managedidentity.azure.com/*;eventhub.azure.com/*"

With remote source:
asoctl export template --source https://github.com/Azure/azure-service-operator/releases/download/v2.6.0/azureserviceoperator_v2.6.0.yaml --crd-pattern "resources.azure.com/*" --crd-pattern "containerservice.azure.com/*"

With local source:
asoctl export template --source ~/Downloads/azureserviceoperator_v2.6.0.yaml --crd-pattern "resources.azure.com/*" --crd-pattern "containerservice.azure.com/*"

Raw:
asoctl export template --version v2.6.0  --raw

With kubectl:
asoctl export template --version v2.6.0 --crd-pattern "resources.azure.com/*;containerservice.azure.com/*;keyvault.azure.com/*;managedidentity.azure.com/*;eventhub.azure.com/*" | kubectl apply -f -

Flags:
  -p, --crd-pattern strings   What new CRDs to install. Existing ASO CRDs in the cluster will always be upgraded even if crdPattern is empty. See https://azure.github.io/azure-service-operator/guide/crd-management/ for more details.
  -h, --help                  help for template
      --raw                   Export the YAML without any variable replacements
  -s, --source string         File or URL path to the ASO YAML template. Use this if you've customized the base ASO YAML locally or are using a base YAML other than the one hosted at https://github.com/Azure/azure-service-operator/tags
  -v, --version string        Release version to use.

Global Flags:
      --quiet     Silence most logging
      --verbose   Enable verbose logging
```

### --crd-pattern

From v2.10, `asoctl` does proactive checking of the `--crd-pattern` parameter to see how many CRDs are selected.

A warning is displayed if a pattern does not match any CRDs. This is nonfatal - the export of the template will proceed - as it's possible the pattern will match a CRD included in a newer version of ASO.

## Clean CRDs

This command can be used to prepare ASOv2 `v1alpha1api`(deprecated in v2.0.0) CustomResources and CustomResourceDefinitions for ASO `v2.0.0` release. 
It ensures that any ASOv2 `v1alpha1api` deprecated version resources that may have been stored in etcd get migrated to `v1beta` version before upgrading ASO to `v2.0.0`. 

```bash
$ asoctl clean crds --help
Clean deprecated CRD versions from cluster

Usage:
  asoctl clean crds [flags]

Flags:
      --dry-run   
  -h, --help      help for clean

Global Flags:
      --verbose   Enable verbose logging
```

`--dry-run` flag outputs about CRDs and CRs to be updated and **does not** modify any CRD and CRs.

### Steps for migration using `asoctl clean crds`:

**Prerequisite:** Ensure the current ASO v2 version in your cluster is `beta.5`.

Run the migration tool:
``` bash
$ asoctl clean crds
```

Once that's successfully run, you can upgrade ASO to `v2.0.0`

Using `asoctl clean crds` is an important step if `v1alpha1api` resources have ever been present in the cluster. If not used correctly, operator pod will produce log error messages:

```
"msg"="failed to apply CRDs" "error"="failed to apply CRD storageaccountsqueueservicesqueues.storage.azure.com: CustomResourceDefinition.apiextensions.k8s.io \"storageaccountsqueueservicesqueues.storage.azure.com\" is invalid: status.storedVersions[0]: Invalid value: \"v1alpha1api20210401storage\": must appear in spec.versions" 
```

### Example Output

```bash
$ asoctl clean crds
...
INF Starting cleanup crd-name=resourcegroups.resources.azure.com
INF Migration finished resource-count=1
INF Updated CRD status storedVersions crd-name=resourcegroups.resources.azure.com storedVersions=["v1beta20200601"]
INF Starting cleanup crd-name=roleassignments.authorization.azure.com
INF Migration finished resource-count=0
INF Updated CRD status storedVersions crd-name=roleassignments.authorization.azure.com storedVersions=["v1beta20200801previewstorage"]
INF Nothing to update crd-name=routetables.network.azure.com
INF Nothing to update crd-name=routetablesroutes.network.azure.com
INF Update finished crd-count=67
```


## Import Azure Resource

When you have an existing Azure resource that needs to be managed by ASO, you can use the `import` command to generate a YAML file that can be used to create a new ASO resource. This is useful when:

* You want to migrate an existing Azure resource from ASO v1 to ASO v2.
* You have an existing Azure resource and want ASO to manage it.
* You have a hand-configured Azure resource and you want to replicate it for development, staging, or production use.

``` bash
$ asoctl import azure-resource --help
Imports ARM resources as Custom Resources.

This command requires you to authenticate with Azure using an identity which has access to the resource(s) you would
like to import. The following authentication modes are supported:

Az-login token: az login and then use asoctl.
Managed Identity: Set the AZURE_CLIENT_ID environment variable and run on a machine with access to the managed identity.
Service Principal: Set the AZURE_SUBSCRIPTION_ID, AZURE_TENANT_ID, AZURE_CLIENT_ID, and AZURE_CLIENT_SECRET environment variables,

The following environment variables can be used to configure which cloud to use with asoctl:

AZURE_RESOURCE_MANAGER_ENDPOINT: The Azure Resource Manager endpoint. 
If not specified, the default is the Public cloud resource manager endpoint.
See https://docs.microsoft.com/cli/azure/manage-clouds-azure-cli#list-available-clouds for details
about how to find available resource manager endpoints for your cloud. Note that the resource manager
endpoint is referred to as "resourceManager" in the Azure CLI.

AZURE_RESOURCE_MANAGER_AUDIENCE: The Azure Resource Manager AAD audience.
If not specified, the default is the Public cloud resource manager audience https://management.core.windows.net/.
See https://docs.microsoft.com/cli/azure/manage-clouds-azure-cli#list-available-clouds for details
about how to find available resource manager audiences for your cloud. Note that the resource manager
audience is referred to as "activeDirectoryResourceId" in the Azure CLI.

AZURE_AUTHORITY_HOST: The URL of the AAD authority.
If not specified, the default
is the AAD URL for the public cloud: https://login.microsoftonline.com/. See
https://docs.microsoft.com/azure/active-directory/develop/authentication-national-cloud

Usage:
  asoctl import azure-resource <ARM/ID/of/resource> [flags]

Flags:
  -a, --annotation strings     Add the specified annotations to the imported resources. Multiple comma-separated annotations can be specified (--annotation example.com/myannotation=foo,example.com/myannotation2=bar) or the --annotation (-a) argument can be used multiple times (-a example.com/myannotation=foo -a example.com/myannotation2=bar)
  -h, --help                   help for azure-resource
  -l, --label strings          Add the specified labels to the imported resources. Multiple comma-separated labels can be specified (--label example.com/mylabel=foo,example.com/mylabel2=bar) or the --label (-l) argument can be used multiple times (-l example.com/mylabel=foo -l example.com/mylabel2=bar)
  -n, --namespace string       Write the imported resources to the specified namespace
  -o, --output string          Write ARM resource CRDs to a single file
  -f, --output-folder string   Write ARM resource CRDs to individual files in a folder
  -w, --workers int            Specify the number of parallel workers to use when importing resources (default 4)
  
Global Flags:
      --quiet     Silence most logging
      --verbose   Enable verbose logging

```

The `asoctl import azure-resource` command will accept any number of ARM resource Ids. 

Each ARM resource will be scanned for supported child or extension resources, and all the results combined together into a single YAML file. If `asoctl` encounters a resource type that it doesn't support, details will be logged. 

### Example: Importing a PostgreSQL Server

To import the configuration of an existing PostgreSQL server, we'd run the following command:

``` bash
$ asoctl import azure-resource /subscriptions/[redacted]/resourceGroups/aso-rg/providers/Microsoft.DBforPostgreSQL/flexibleServers/aso-pg --output aso.yaml
```
* The parameter is the full ARM ID to the server; one way to find this is via the Azure Portal.
* Multiple ARM IDs are permitted.
* Here, the subscription ID has been redacted; in practice, you'd see your own subscription list

While `asoctl import azure-resource` runs, you'll see progress shown dynamically as child and extension resources are found and imported:

![Screenshot showing asoctl import azure-resource running](../images/asoctl-import-progress-bars.png)

Once finished, you'll see a list of all the imported resources, and the file they were written to:

``` bash
14:43:34 INF Imported kind=FlexibleServer.dbforpostgresql.azure.com name=aso-pg
14:43:37 INF Imported kind=FlexibleServersFirewallRule.dbforpostgresql.azure.com name=AllowAllAzureServicesAndResourcesWithinAzureIps_2023-3-30_15-26-24
14:44:06 INF Imported kind=FlexibleServersConfiguration.dbforpostgresql.azure.com name=shared_preload_libraries
14:44:15 INF Imported kind=FlexibleServersConfiguration.dbforpostgresql.azure.com name=vacuum_cost_page_miss
14:44:18 INF Imported kind=FlexibleServersDatabase.dbforpostgresql.azure.com name=aso-demo
14:44:25 INF Writing to file path=aso.yaml
Import Azure Resources  [======================================================================================================================================================================] 100 %
```

The import found a total of five resources:

* 1 x PostgreSQL Flexible Server (the one originally identified by ARM Id)
* 1 x Firewall Rule
* 2 x Configuration settings
* 1 x Database

Not all resources found need to imported; to see more information about those resources, and why they were omitted, use the `--verbose` command line flag.

One situation where this happens is with the configuration of a PostgreSQL server, where most of the values are system defaults that don't need to be explicitly specified.

``` bash
14:47:56 DBG Skipped because="system-defaults don't need to be imported" kind=FlexibleServersConfiguration.dbforpostgresql.azure.com name=synchronous_commit
14:47:56 DBG Skipped because="system-defaults don't need to be imported" kind=FlexibleServersConfiguration.dbforpostgresql.azure.com name=tcp_keepalives_idle
14:47:57 DBG Skipped because="system-defaults don't need to be imported" kind=FlexibleServersConfiguration.dbforpostgresql.azure.com name=tcp_keepalives_count
14:47:57 DBG Skipped because="system-defaults don't need to be imported" kind=FlexibleServersConfiguration.dbforpostgresql.azure.com name=tcp_keepalives_interval
14:47:57 DBG Skipped because="system-defaults don't need to be imported" kind=FlexibleServersConfiguration.dbforpostgresql.azure.com name=temp_buffers
14:47:57 DBG Skipped because="system-defaults don't need to be imported" kind=FlexibleServersConfiguration.dbforpostgresql.azure.com name=temp_tablespaces
```

Another is when role assignments have been inherited by the resource:

``` bash
14:48:03 DBG Skipped because="role assignment is inherited" kind=RoleAssignment.authorization.azure.com name=fe65ff58-[redacted]
14:48:03 DBG Skipped because="role assignment is inherited" kind=RoleAssignment.authorization.azure.com name=61681810-[redacted]
14:48:03 DBG Skipped because="role assignment is inherited" kind=RoleAssignment.authorization.azure.com name=06c013cb-[redacted]
14:48:03 DBG Skipped because="role assignment is inherited" kind=RoleAssignment.authorization.azure.com name=37047941-[redacted]
14:48:03 DBG Skipped because="role assignment is inherited" kind=RoleAssignment.authorization.azure.com name=b0b9667d-[redacted]
14:48:03 DBG Skipped because="role assignment is inherited" kind=RoleAssignment.authorization.azure.com name=88a998ec-[redacted]
```

