#!/usr/bin/env bash
set -euxo pipefail

version="$1"
target="multitenant-cluster_${version}.yaml"
source="$2"

# Collect all of the documents we want in the cluster deployment yaml.
crd='.kind == "CustomResourceDefinition"'
namespace='.kind == "Namespace"'
cluster_role='.kind == "ClusterRole"'
leader_election_role='(.kind == "Role" and .metadata.name == "azureserviceoperator-leader-election-role")'
leader_election_binding='(.kind == "RoleBinding" and .metadata.name == "azureserviceoperator-leader-election-rolebinding")'
proxy_binding='(.kind == "ClusterRoleBinding" and .metadata.name == "azureserviceoperator-proxy-rolebinding")'
service='.kind == "Service"'
deployment='(.kind == "Deployment" and .metadata.name == "azureserviceoperator-controller-manager")'
certificate='.kind == "Certificate"'
issuer='.kind == "Issuer"'
webhooks='(.kind == "MutatingWebhookConfiguration" or .kind == "ValidatingWebhookConfiguration")'
query="select($crd or $namespace or $cluster_role or $leader_election_role or $leader_election_binding or $service or $deployment or $certificate or $issuer or $webhooks or $proxy_binding)"

yq eval "$query" "$source" > "$target"


# Edit the deployment to run the ASO binary in webhooks-only mode.
# Remove the aadpodidbinding label - this is only needed for communicating to ARM
yq eval -i "del(select($deployment) | .spec.template.metadata.labels.aadpodidbinding)" "$target"

# Edit the deployment to turn off leader election for the webhook only pod, as the webhooks don't wait for leader election anyway.
yq eval -i "del(select($deployment) | .spec.template.spec.containers[] | select(.name == \"manager\").args[] | select(. == \"--enable-leader-election\"))" "$target"

# Change the manager container env vars - the webhook server only
# needs pod namespace, operator mode and subscription id (which isn't
# used).

# This is awkward but I can't see a better way to express it in yq -
# it corresponds to this yaml:
# - name: POD_NAMESPACE
#   valueFrom:
#     fieldRef:
#       fieldPath: metadata.namespace
# - name: AZURE_OPERATOR_MODE
#   value: webhooks
#   # This is unfortunately required although we won't use it.
# - name: AZURE_SUBSCRIPTION_ID
#   value: none
pod_namespace='{"name": "POD_NAMESPACE", "valueFrom": {"fieldRef": {"fieldPath": "metadata.namespace"}}}'
operator_mode='{"name": "AZURE_OPERATOR_MODE", "value": "webhooks"}'
subscription_id='{"name": "AZURE_SUBSCRIPTION_ID", "value": "none"}'
new_env_vars="[$pod_namespace, $operator_mode, $subscription_id]"
yq eval -i "(select($deployment) | .spec.template.spec.containers[] | select(.name == \"manager\").env) = $new_env_vars" "$target"
