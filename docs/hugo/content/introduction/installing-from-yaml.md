---
title: "Installation: From YAML"
---

### Prerequisites
1. You have installed Cert Manager as per the [installation instructions](https://azure.github.io/azure-service-operator/#installation) up to the "install from Helm" step.
2. You have the `AZURE_SUBSCRIPTION_ID`, `AZURE_TENANT_ID`, `AZURE_CLIENT_ID` and `AZURE_CLIENT_SECRET` environment variables set from the
   [installation instructions](https://azure.github.io/azure-service-operator/#installation).

### Installation

1. Install [the latest **v2+** release](https://github.com/Azure/azure-service-operator/releases) of Azure Service Operator.
   ```bash
   kubectl apply --server-side=true -f https://github.com/Azure/azure-service-operator/releases/download/v2.0.0-beta.0/azureserviceoperator_v2.0.0-beta.0.yaml
   ```
2. Create the Azure Service Operator v2 secret. This secret contains the identity that Azure Service Operator will run as. 
   Make sure that you have the 4 environment variables from the [Helm installation instructions](https://azure.github.io/azure-service-operator/#installation) set.
   To learn more about other authentication options, see the [authentication documentation](https://azure.github.io/azure-service-operator/introduction/authentication/):
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