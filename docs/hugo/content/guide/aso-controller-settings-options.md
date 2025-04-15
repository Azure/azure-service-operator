---
title: Configuring ASO
linktitle: Configuring ASO
weight: 1 # This is the default weight if you just want to be ordered alphabetically
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

**[Allowed scopes]( {{< relref "authentication#credential-scope" >}} )**: Global, namespace, or per-resource

This may be set to empty string to configure no global credential.

### AZURE_TENANT_ID

Azure tenantID the operator will use for ARM communication if
no [more specific]( {{< relref "authentication/credential-scope" >}} )
credential is specified at the per-resource or per-namespace scope.

**Format:** `GUID`

**Example:** `00000000-0000-0000-0000-000000000000`

**Required**: True

**[Allowed scopes]( {{< relref "authentication#credential-scope" >}} )**: Global, namespace, or per-resource

This may be set to empty string to configure no global credential.

### AZURE_CLIENT_ID

Azure clientID the operator will use for ARM communication if
no [more specific]( {{< relref "authentication/credential-scope" >}} )
credential is specified at the per-resource or per-namespace scope.

**Format:** `GUID`

**Example:** `00000000-0000-0000-0000-000000000000`

**Required**: True

**[Allowed scopes]( {{< relref "authentication#credential-scope" >}} )**: Global, namespace, or per-resource

This may be set to empty string to configure no global credential.

### AZURE_CLIENT_SECRET

The secret associated with the client to use if
no [more specific]( {{< relref "authentication/credential-scope" >}} )
credential is specified at the per-resource or per-namespace scope.

**Format:** `String`

**Required**: False

**[Allowed scopes]( {{< relref "authentication#credential-scope" >}} )**: Global, namespace, or per-resource

### AZURE_CLIENT_CERTIFICATE

AzureClientCertificate is a PEM or PKCS12 certificate string including the private key for 
Azure Credential Authentication.
If the certificate is password protected,  use `AZURE_CLIENT_CERTIFICATE_PASSWORD` for the password.

**Format:** `String`

**Required**: False

**[Allowed scopes]( {{< relref "authentication#credential-scope" >}} )**: Global, namespace, or per-resource

### AZURE_CLIENT_CERTIFICATE_PASSWORD

The password used to protect the `AZURE_CLIENT_CERTIFICATE`.

**Format:** `String`

**Required**: False

**[Allowed scopes]( {{< relref "authentication#credential-scope" >}} )**: Global, namespace, or per-resource

### AZURE_SYNC_PERIOD

AZURE_SYNC_PERIOD is the frequency at which resources are re-reconciled with Azure when
there have been no triggering changes in the Kubernetes resources.
This sync exists to detect and correct changes that happened in Azure that Kubernetes is not aware about.
BE VERY CAREFUL setting this value low - even a modest number of resources can cause
subscription level throttling if they are re-synced frequently. If nil or empty (`""`), sync period defaults to `1h`.

Specify the special value `"never"` to stop syncing.

**Format:** `duration string`

**Example:** `"1h"`, `"15m"`, or `"60s"`. See [ParseDuration](https://pkg.go.dev/time#ParseDuration) for more details.

**Required**: False

**[Allowed scopes]( {{< relref "authentication#credential-scope" >}} )**: Global

### AZURE_OPERATOR_MODE

AZURE_OPERATOR_MODE determines whether the operator should run _watchers_, _webhooks_ or _both_ (default). An empty string, or any unrecognized value, means _both_.

**Format:** `webhooks|watchers|both`

**Examples:** `"webhooks"`, `"watchers"` or `"both"`

**Required**: False

**[Allowed scopes]( {{< relref "authentication#credential-scope" >}} )**: Global

### AZURE_TARGET_NAMESPACES

AZURE_TARGET_NAMESPACES lists the namespaces the operator will watch for Azure resources (if the mode includes running watchers). 
If it's empty the operator will watch all namespaces.

Spaces after `,`'s and at the start and end of the string are ignored.

**Format:** `string` (comma-separated namespace names)

**Example:** `ns1,ns2`

**Required**: False

**[Allowed scopes]( {{< relref "authentication#credential-scope" >}} )**: Global

### USE_WORKLOAD_IDENTITY_AUTH

USE_WORKLOAD_IDENTITY_AUTH boolean is used to determine if we're using Workload Identity authentication for global credential.

**Format:** `true|false`

**Example:** `"true"` or `"false"`

**Required**: False

**[Allowed scopes]( {{< relref "authentication#credential-scope" >}} )**: Global

### AZURE_AUTHORITY_HOST

AZURE_AUTHORITY_HOST is the URL of the AAD authority. If not specified, the default is the AAD URL for the public cloud: `https://login.microsoftonline.com/`. 
See https://docs.microsoft.com/azure/active-directory/develop/authentication-national-cloud

**Format:** `string`

**Example:** `"https://login.chinacloudapi.cn"`

**Required**: False

**[Allowed scopes]( {{< relref "authentication#credential-scope" >}} )**: Global

### AZURE_RESOURCE_MANAGER_ENDPOINT

AZURE_RESOURCE_MANAGER_ENDPOINT is the Azure Resource Manager endpoint. If not specified, the default is the Public cloud resource manager endpoint.
See https://docs.microsoft.com/cli/azure/manage-clouds-azure-cli#list-available-clouds for details about how to find available resource manager endpoints for your cloud. 
Note that the resource manager endpoint is referred to as "resourceManager" in the Azure CLI.

**Format:** `string`

**Example:** `"https://management.chinacloudapi.cn"`

**Required**: False

**[Allowed scopes]( {{< relref "authentication#credential-scope" >}} )**: Global

### AZURE_RESOURCE_MANAGER_AUDIENCE

AZURE_RESOURCE_MANAGER_AUDIENCE is the Azure Resource Manager AAD audience. If not specified, the default is the Public cloud resource manager audience `https://management.core.windows.net/`.
See https://docs.microsoft.com/cli/azure/manage-clouds-azure-cli#list-available-clouds for details about how to find available resource manager audiences for your cloud. 
Note that the resource manager audience is referred to as "activeDirectoryResourceId" in the Azure CLI.

**Format:** `string`

**Example:** `"https://management.core.chinacloudapi.cn/"`

**Required**: False

**[Allowed scopes]( {{< relref "authentication#credential-scope" >}} )**: Global

### AZURE_ADDITIONAL_TENANTS

The list of (comma-separated) additional tenants the operator can authenticate with.
This is required when performing cross-tenant authentication. See the
[Entra documentation](https://learn.microsoft.com/entra/external-id/cross-tenant-access-overview) for more details.

**Format:** `string` (comma-separated tenant GUIDs - spaces are allowed)

**Example:** `00000000-0000-0000-0000-000000000001,00000000-0000-0000-0000-000000000002`

**Required**: False

**[Allowed scopes]( {{< relref "authentication#credential-scope" >}} )**: Global, namespace, or per-resource

### AZURE_USER_AGENT_SUFFIX

AZURE_USER_AGENT_SUFFIX is appended to the default User-Agent for Azure HTTP clients.

**Format:** `string`

**Example:** `"my-user-agent"`

**Required**: False

**[Allowed scopes]( {{< relref "authentication#credential-scope" >}} )**: Global

### ENTRA_APP_ID

Required if you want to use Entra resources, ENTRA_APP_ID identifies ASO to Microsoft Entra.

To create an ENTRA_APP_ID of your own, see [Register an application in Microsoft Entra ID](https://learn.microsoft.com/en-us/entra/identity-platform/quickstart-register-app).

Format: `string` (GUID)

Example: `00000000-0000-0000-0000-000000000000`

Required: True if using Entra resources, otherwise False


### MAX_CONCURRENT_RECONCILES

MAX_CONCURRENT_RECONCILES is the number of threads/goroutines dedicated to reconciling each resource type.
If not specified, the default is 1.

IMPORTANT: Having MAX_CONCURRENT_RECONCILES set to N does not mean that ASO is limited to N interactions with
Azure at any given time, because the control loop yields to another resource while it is not actively issuing HTTP
calls to Azure. Any single resource only blocks the control-loop for its resource-type for as long as it takes to issue
an HTTP call to Azure, view the result, and make a decision. In most cases the time taken to perform these actions
(and thus how long the loop is blocked and preventing other resources from being acted upon) is a few hundred
milliseconds to at most a second or two. In a typical 60s period, many hundreds or even thousands of resources
can be managed with this set to 1.

MAX_CONCURRENT_RECONCILES applies to every registered resource type being watched/managed by ASO.

**Format:** `int`

**Example:** `2`

**Required**: False

**[Allowed scopes]( {{< relref "authentication#credential-scope" >}} )**: Global

### RATE_LIMIT_MODE

RateLimitMode configures the internal rate-limiting mode.

 * **disabled** (default): No ASO-controlled rate-limiting occurs. ASO will attempt to communicate with Azure and
   kube-apiserver as much as needed based on load. It will back off based on throttling from
   either kube-apiserver or Azure, but will not artificially limit its throughput.
 * **bucket**: Uses a token-bucket algorithm to rate-limit reconciliations. Note that this limits how often
   the operator performs a reconciliation, but not every reconciliation triggers a call to kube-apiserver
   or Azure (though many do). Since this controls reconciles it can be used to coarsely control throughput
   and CPU usage of the operator, as well as the number of requests that the operator issues to Azure.
   Keep in mind that the Azure throttling limits (defined at
   https://learn.microsoft.com/en-us/azure/azure-resource-manager/management/request-limits-and-throttling)
   differentiate between request types. Since a given reconcile for a resource may result in polling (a GET) or
   modification (a PUT) it's not possible to entirely avoid Azure throttling by tuning these bucket limits.
   
   We don't recommend enabling this mode by default.

   If enabling this mode, we strongly recommend doing some experimentation to tune these values to something to
   works for your specific need.

**Format:** `disabled|bucket`

**Example:** `disabled` or `bucket`

**Required**: False

**[Allowed scopes]( {{< relref "authentication#credential-scope" >}} )**: Global

### RATE_LIMIT_QPS

RATE_LIMIT_QPS is the rate (per second) that the bucket is refilled. 
This value only has an effect if RATE_LIMIT_MODE is 'bucket'.

**Format:** `float`

**Example:** `5`

**Required**: False

**[Allowed scopes]( {{< relref "authentication#credential-scope" >}} )**: Global

### RATE_LIMIT_BUCKET_SIZE

RATE_LIMIT_BUCKET_SIZE is the size of the bucket. This value only has an effect if RATE_LIMIT_MODE is 'bucket'.

**Format:** `int`

**Example:** `200`

**Required**: False

**[Allowed scopes]( {{< relref "authentication#credential-scope" >}} )**: Global

### DEFAULT_RECONCILE_POLICY

DEFAULT_RECONCILE_POLICY specifies the reconcile strategy to be used by the operator. If not specified, it is set to 'manage'.

**Format:** `string`

**Example:** `detach-on-delete`

**Required**: False

**[Allowed scopes]( {{< relref "authentication#credential-scope" >}} )**: Global
