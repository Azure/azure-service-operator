# Deploying ASO v2 in multi-tenant mode

The current release version of the ASO deployment YAML installs the operator in single-tenant mode:
it's deployed in the `azureserviceoperator-system` namespace with one set of Azure credentials, and manages resources in any namespace in the cluster.
That single operator deployment handles webhooks fired when Azure resources are changed.

ASO can also be deployed in a multi-tenant configuration that enables using separate credentials to manage resources in different Kubernetes namespaces.

Running the operator in multi-tenant mode requires one deployment to handle webhooks (since webhook configurations are cluster-level resources) and then a deployment for each tenant, each with its own credentials and set of namespaces that it watches for Azure resources.
To deploy the operator in multi-tenant mode the release YAML has been split into two parts:

1. Cluster-wide resources:
   * Custom resource definitions for the Azure resources
   * Cluster roles for managing those resources
   * The `azureserviceoperator-system` namespace containing the deployment and service to run the operator.
     The webhook deployment of the operator is configured to run in webhook-only mode.
     It doesn't require any Azure credentials, since this operator never talks to ARM, only to the kubernetes API server.
   * Webhook configuration referring to that service.

2. Per-tenant resources:
   * A namespace containing the deployment to run the tenant operator, configured for watchers-only mode.
   * The `aso-controller-settings` secret defining the Azure credentials that should be used, and the set of namespaces that this operator will watch for Azure resources.
   * A cluster role binding enabling the per-tenant operator's service account to manage the Azure resources.

## Example files
Examples of the deployment YAML files are available on the release page for ASO v2 releases from [v2.0.0-alpha.5](https://github.com/Azure/azure-service-operator/releases/tag/v2.0.0-alpha.5).
The cluster-wide file `multitenant-cluster_v2.0.0-alpha.5.yaml` can be used as-is (the webhook operator namespace is fixed as `azureserviceoperator-system`),
but the namespaces and cluster role binding in the per-tenant file `multitenant-tenant_v2.0.0-alpha.5.yaml` will need to be updated from `tenant1` to the desired name for each tenant.

## Per-tenant configuration
Create the `aso-controller-settings` secret as described in the [authentication docs](https://azure.github.io/azure-service-operator/introduction/authentication/),
but create the secret in the tenant namespace and add an extra target namespaces key to it:
```
export TENANT_NAMESPACE="<tenant namespace>"
export AZURE_SUBSCRIPTION_ID="<subscription id>"
export AZURE_TENANT_ID="<tenant id>"
export AZURE_CLIENT_ID="<client id>"
export AZURE_CLIENT_SECRET="<client secret>"
export AZURE_TARGET_NAMESPACES="<comma-separated namespace names>"
kubectl create namespace "$TENANT_NAMESPACE"
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Secret
metadata:
  name: aso-controller-settings
  namespace: $TENANT_NAMESPACE
stringData:
  AZURE_SUBSCRIPTION_ID: "$AZURE_SUBSCRIPTION_ID"
  AZURE_TENANT_ID: "$AZURE_TENANT_ID"
  AZURE_CLIENT_ID: "$AZURE_CLIENT_ID"
  AZURE_CLIENT_SECRET: "$AZURE_CLIENT_SECRET"
  AZURE_TARGET_NAMESPACES: "$AZURE_TARGET_NAMESPACES"
EOF
```

## Role handling
The multi-tenant deployment example files have a cluster role granting access to the Azure resource types whichever namespaces they are created in,
and then a binding to that cluster role for the service account in each tenant-operator namespace.
Each cluster role binding is named for the specific tenant so they don't collide and can be managed separately.
This is convenient since it means that you don't need to permit access for Azure resources in each of the target namespaces individually,
but for some usage scenarios it might be too permissive.

In those cases the `azureserviceoperator-manager-role` should be changed from a `ClusterRole` into `Role`s in each of the target namespaces (where the Azure resources will be created, rather than where the operator pods run),
and a `RoleBinding` should be created in that namespace linking the `Role` to the service account for the tenant operator that will be managing Azure resources in this target namespace.

## Updgrading
When upgrading to a newer version of ASO the cluster-wide resources (CRDs, cluster roles) and the webhook operator deployment must be upgraded before upgrading the tenant operators.
They can be upgraded by applying the new YAML for that type.
