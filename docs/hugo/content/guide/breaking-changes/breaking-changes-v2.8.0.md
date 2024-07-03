---
title: "v2.8.0 Breaking Changes"
linkTitle: "v2.8.0"
weight: -30  # This should be 5 lower than the previous breaking change document
---

## ManagedCluster property .spec.properties.windowsProfile.adminPassword is now a SecretReference rather than a string

The `spec.properties.windowsProfile.adminPassword` on `ManagedCluster` has been changed from a string to a 
`SecretReference`. 

We try to avoid breaking changes, but in this case, allowing raw passwords in the spec is a security 
problem and as such we've decided to make a break to correct this issue.

**Action required:** If the `containerservice.azure.com/ManagedCluster` resource is used in your cluster 
**and** the `spec.properties.windowsProfile.adminPassword` property is set, do the 
following before upgrading ASO:

1. Annotate the resource with `serviceoperator.azure.com/reconcile-policy: skip` to prevent ASO from trying 
   to reconcile the resource while you are upgrading.
2. Save the current YAML for the resource using `kubectl` if you don't have it elsewhere.
3. Create a kubernetes secret containing the value from `spec.properties.windowsProfile.adminPassword`.
4. Edit the resource YAML in step 2, and add a secret key and name reference. 
   Example [here](https://github.com/Azure/azure-service-operator/blob/main/v2/samples/apimanagement/v1api20230501preview/v1api20230501preview_authorizationprovider.yaml#L12).
5. Delete the resource from your cluster using `kubectl delete`. Your Azure resource will be left untouched because 
   of the `reconcile-policy` annotation you added above.
6. [Upgrade ASO]( {{< relref "upgrading" >}} ) in your cluster.
7. Apply the updated YAML to your cluster using `kubectl apply`. If any errors occur, address them.
8. If the `reconcile-policy` annotation is still present, remove it from the resource.

## Use "never" rather than "" to prevent syncing for AZURE_SYNC_PERIOD

The documentation always said that an `AZURE_SYNC_PERIOD` of `""` meant to use the default value (`15m`), but
in actuality in the code `""` meant never sync. This corrects the behavior to be as documented. After this release,
the behavior is as follows:

| AZURE_SYNC_PERIOD | Meaning                  |
| ----------------- | ------------------------ |
| omitted/not set   | Use default value (15m)  |
| ""                | Use default value (15m)  |
| "never"           | Do not sync              |

If you intend to prevent ASO from periodically syncing with Azure, set `AZURE_SYNC_PERIOD` to `"never"`.
For more information, see #3965.

## Upcoming breaking changes

### v20230202preview ManagedCluster will be removed, due to underlying Azure API version deprecation

Other versions will continue to be supported. We recommend you move to use a different CRD version to avoid
seeing errors from Azure due to the 2023-02-02-preview API being deprecated.
