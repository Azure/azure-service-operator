---
title: Configuring ASO
linktitle: Configuring ASO
---

Configuration of ASO is done primarily through a secret in the `azureserviceoperator-system` 
namespace called `aso-controller-settings`. This secret contains both details about the (optional) global credential
as well as other operator pod options. 

The supported options are:


### AZURE_SUBSCRIPTION_ID

Azure subscription the operator will use for ARM communication if 
no [more specific]( {{< relref "authentication/credential-scope" >}} )
credential is specified at the per-resource or per-namespace scope.

**Format:** `GUID`

**Example:** `00000000-0000-0000-0000-000000000000`

**Required**: True

This may be set to empty string to configure no global credential.

### AZURE_TENANT_ID

Azure tenantID the operator will use for ARM communication if
no [more specific]( {{< relref "authentication/credential-scope" >}} )
credential is specified at the per-resource or per-namespace scope.

**Format:** `GUID`

**Example:** `00000000-0000-0000-0000-000000000000`

**Required**: True

This may be set to empty string to configure no global credential.

### AZURE_CLIENT_ID

Azure clientID the operator will use for ARM communication if
no [more specific]( {{< relref "authentication/credential-scope" >}} )
credential is specified at the per-resource or per-namespace scope.

**Format:** `GUID`

**Example:** `00000000-0000-0000-0000-000000000000`

**Required**: True

This may be set to empty string to configure no global credential.

### AZURE_CLIENT_SECRET

The secret associated with the client to use if
no [more specific]( {{< relref "authentication/credential-scope" >}} )
credential is specified at the per-resource or per-namespace scope.

**Format:** `String`

**Required**: False

### AZURE_SYNC_PERIOD

AZURE_SYNC_PERIOD is the frequency at which resources are re-reconciled with Azure when
there have been no triggering changes in the Kubernetes resources.
This sync exists to detect and correct changes that happened in Azure that Kubernetes is not aware about.
BE VERY CAREFUL setting this value low - even a modest number of resources can cause
subscription level throttling if they are re-synced frequently. If nil or empty (`""`), sync period defaults to `1h`.

**Format:** `duration string`

**Example:** `"1h"`, `"15m"`, or `"60s"`. See [ParseDuration](https://pkg.go.dev/time#ParseDuration) for more details.

**Required**: False

### AZURE_OPERATOR_MODE

AZURE_OPERATOR_MODE determines whether the operator should run _watchers_, _webhooks_ or _both_ (default). An empty string, or any unrecognized value, means _both_.

**Format:** `webhooks|watchers|both`

**Examples:** `"webhooks"`, `"watchers"` or `"both"`

**Required**: False

### AZURE_TARGET_NAMESPACES

AZURE_TARGET_NAMESPACES lists the namespaces the operator will watch for Azure resources (if the mode includes running watchers). 
If it's empty the operator will watch all namespaces.

Spaces after `,`'s and at the start and end of the string are ignored.

**Format:** `string` (comma-separated namespace names)

**Example:** `ns1,ns2`

**Required**: False

### USE_WORKLOAD_IDENTITY_AUTH

USE_WORKLOAD_IDENTITY_AUTH boolean is used to determine if we're using Workload Identity authentication for global credential.

**Format:** `true|false`

**Example:** `"true"` or `"false"`

**Required**: False

### AZURE_AUTHORITY_HOST

AZURE_AUTHORITY_HOST is the URL of the AAD authority. If not specified, the default is the AAD URL for the public cloud: `https://login.microsoftonline.com/`. 
See https://docs.microsoft.com/azure/active-directory/develop/authentication-national-cloud

**Format:** `string`

**Example:** `"https://login.chinacloudapi.cn"`

**Required**: False

### AZURE_RESOURCE_MANAGER_ENDPOINT

AZURE_RESOURCE_MANAGER_ENDPOINT is the Azure Resource Manager endpoint. If not specified, the default is the Public cloud resource manager endpoint.
See https://docs.microsoft.com/cli/azure/manage-clouds-azure-cli#list-available-clouds for details about how to find available resource manager endpoints for your cloud. 
Note that the resource manager endpoint is referred to as "resourceManager" in the Azure CLI.

**Format:** `string`

**Example:** `"https://management.chinacloudapi.cn"`

**Required**: False

### AZURE_RESOURCE_MANAGER_AUDIENCE

AZURE_RESOURCE_MANAGER_AUDIENCE is the Azure Resource Manager AAD audience. If not specified, the default is the Public cloud resource manager audience `https://management.core.windows.net/`.
See https://docs.microsoft.com/cli/azure/manage-clouds-azure-cli#list-available-clouds for details about how to find available resource manager audiences for your cloud. 
Note that the resource manager audience is referred to as "activeDirectoryResourceId" in the Azure CLI.

**Format:** `string`

**Example:** `"https://management.core.chinacloudapi.cn/"`

**Required**: False
