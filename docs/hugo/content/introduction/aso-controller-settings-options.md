---
title: ASO-Controller-Settings Options
linktitle: ASO-Controller-Settings Options
---


### AZURE_SUBSCRIPTION_ID

Azure subscription the operator will use for ARM communication. 

**Specified As:** `00000000-0000-0000-0000-000000000000`

### AZURE_TENANT_ID

Azure tenantID the operator will use for ARM communication.

**Specified As:** `00000000-0000-0000-0000-000000000000`


### AZURE_CLIENT_ID

Azure clientID the operator will use for ARM communication.

**Specified As:** `00000000-0000-0000-0000-000000000000`


### AZURE_CLIENT_SECRET

The secret associated with the client to use.

**Specified As:** `String`

### AZURE_SYNC_PERIOD

AZURE_SYNC_PERIOD is the frequency at which resources are re-reconciled with Azure when
there have been no triggering changes in the Kubernetes resources.
This sync exists to detect and correct changes that happened in Azure that Kubernetes is not aware about.
BE VERY CAREFUL setting this value low - even a modest number of resources can cause
subscription level throttling if they are re-synced frequently. If nil, sync period defaults to `1h`.

**Specified As:** `"1h"`, `"15m"`, or `"60s"`. See https://pkg.go.dev/time#ParseDuration for more details.

### AZURE_OPERATOR_MODE

AZURE_OPERATOR_MODE determines whether the operator should run watchers, webhooks or both(default).

**Specified As:** `"webhooks"`, `"watchers"` or `"both"`

### AZURE_TARGET_NAMESPACES

AZURE_TARGET_NAMESPACES lists the namespaces the operator will watch for Azure resources (if the mode includes running watchers). 
If it's empty the operator will watch all namespaces.

**Specified As:** `Comma-separated namespace names`


### USE_WORKLOAD_IDENTITY_AUTH

USE_WORKLOAD_IDENTITY_AUTH boolean is used to determine if we're using Workload Identity authentication for global credential.

**Specified As:** `"true"` or `"false"`

### AZURE_AUTHORITY_HOST

AZURE_AUTHORITY_HOST is the URL of the AAD authority. If not specified, the default is the AAD URL for the public cloud: `https://login.microsoftonline.com/`. 
See https://docs.microsoft.com/azure/active-directory/develop/authentication-national-cloud

### AZURE_RESOURCE_MANAGER_ENDPOINT

AZURE_RESOURCE_MANAGER_ENDPOINT is the Azure Resource Manager endpoint. If not specified, the default is the Public cloud resource manager endpoint.
See https://docs.microsoft.com/cli/azure/manage-clouds-azure-cli#list-available-clouds for details about how to find available resource manager endpoints for your cloud. 
Note that the resource manager endpoint is referred to as "resourceManager" in the Azure CLI.

### AZURE_RESOURCE_MANAGER_AUDIENCE

AZURE_RESOURCE_MANAGER_AUDIENCE is the Azure Resource Manager AAD audience. If not specified, the default is the Public cloud resource manager audience `https://management.core.windows.net/`.
See https://docs.microsoft.com/cli/azure/manage-clouds-azure-cli#list-available-clouds for details about how to find available resource manager audiences for your cloud. 
Note that the resource manager audience is referred to as "activeDirectoryResourceId" in the Azure CLI.
