---
title: "Installation: From YAML"
weight: -4
---
## Prerequisites
1. You have installed Cert Manager as per the [installation instructions](../../#installation) up to the "install from Helm" step.
2. You have followed the [instructions for creating a Managed Identity or Service Principal](../../#create-a-managed-identity-or-service-principal)
   and set the appropriate environment variables.

## Installation (operator)

1. Download and install the latest version of [asoctl]({{<relref "asoctl#installation">}}).
2. Install [the latest **v2+** release](https://github.com/Azure/azure-service-operator/releases) of Azure Service Operator.
   ```bash
   asoctl export template --version v2.6.0 --crd-pattern "<your pattern>" | kubectl apply -f -
   ```
   
   When specifying `--crd-pattern`, ensure you choose only the CRDs you need, for example: 
   `--crd-pattern "resources.azure.com/*;containerservice.azure.com/*;keyvault.azure.com/*;managedidentity.azure.com/*;eventhub.azure.com/*"`. 
   For more information about what `--crd-pattern` means, see [CRD management in ASO]({{<relref "crd-management">}}).

3. Create the Azure Service Operator v2 global secret. This secret contains the identity that Azure Service Operator will run as. 
   Make sure that you have the 4 environment variables from the 
   [create a service principal step of the Helm instructions](../../#create-a-managed-identity-or-service-principal) set.
   ```bash
   cat <<EOF | kubectl apply -f -
   apiVersion: v1
   kind: Secret
   metadata:
     name: aso-controller-settings
     namespace: azureserviceoperator-system
   stringData:
     AZURE_SUBSCRIPTION_ID: "$AZURE_SUBSCRIPTION_ID"
     AZURE_TENANT_ID: "$AZURE_TENANT_ID"
     AZURE_CLIENT_ID: "$AZURE_CLIENT_ID"
     AZURE_CLIENT_SECRET: "$AZURE_CLIENT_SECRET"
   EOF
   ```

   Note that unlike the Helm installation instructions, we use the ASO global secret here. To learn more about different
   secret scopes and formats, see the [authentication documentation](../authentication/).

## Installation (crds)

The operator manages its own CRDs by default, based on the `--crd-pattern` specified during installation.

{{% alert title="Warning" color="warning" %}}
If you specified `--crd-pattern` when installing ASO, you don't need to do any of the below.
{{% /alert %}}

If you don't want the operator to manage the CRDs itself, you can install the latest version of the CRDs yourself:

```bash
kubectl apply --server-side=true -f https://github.com/Azure/azure-service-operator/releases/download/v2.6.0/azureserviceoperator_customresourcedefinitions_v2.6.0.yaml
```

{{% alert title="Warning" color="warning" %}}
The azureserviceoperator_customresourcedefinitions_v2.6.0.yaml file contains _all_ the supported CRDs. We recommend filtering
it locally to only the CRDs you want.
{{% /alert %}}

### Troubleshooting

#### Metadata too long

If you omit the `--server-side=true` flag from the `kubectl apply` command, you will see an error like the following:

``` 
CustomResourceDefinition.apiextensions.k8s.io "storageaccounts.storage.azure.com" is invalid:
metadata.annotations: Too long: must have at most 262144 bytes
```

Why does this happen? ASO CRDs are a complete representation of the Azure Resource surface area, including documentation. 
This is tremendously useful - but also means they are quite large.
