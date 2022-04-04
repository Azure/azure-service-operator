#!/usr/bin/env bash
set -euxo pipefail

version="$1"
target="multitenant-tenant_${version}.yaml"
source="$2"

tenant="tenant1"
tenant_namespace="$tenant-system"


# Collect the few documents we want in the per-tenant yaml.
namespace='.kind == "Namespace"'
leader_election_role='(.kind == "Role" and .metadata.name == "azureserviceoperator-leader-election-role")'
leader_election_binding='(.kind == "RoleBinding" and .metadata.name == "azureserviceoperator-leader-election-rolebinding")'
manager_role_binding='(.kind == "ClusterRoleBinding" and .metadata.name == "azureserviceoperator-manager-rolebinding")'
deployment='(.kind == "Deployment" and .metadata.name == "azureserviceoperator-controller-manager")'
query="select($namespace or $leader_election_role or $leader_election_binding or $manager_role_binding or $deployment)"

yq eval "$query" "$source" > "$target"

function inplace_edit() {
    command=$1
    yq eval -i "$command" "$target"
}

# Create the tenant namespace.
inplace_edit "(select($namespace).metadata.name) = \"$tenant_namespace\""

# Put the other resources into that namespace (except for the cluster
# role binding since that's not namespaced).
inplace_edit "(select($leader_election_role or $leader_election_binding or $deployment).metadata.namespace) = \"$tenant_namespace\""

# Update the subject namespaces for the bindings so they refer to the
# service account in the tenant namespace
inplace_edit "(select($leader_election_binding or $manager_role_binding).subjects[0].namespace) = \"$tenant_namespace\""

# Rename the cluster role binding so bindings for different tenants
# can coexist.
inplace_edit "(select($manager_role_binding).metadata.name) = \"azureserviceoperator-manager-rolebinding-$tenant\""

# Changes to the deployment:
# * Remove the webserver cert volume and mount.
inplace_edit "del(select($deployment) | .spec.template.spec.volumes)"

manager_container="select($deployment) | .spec.template.spec.containers[] | select(.name == \"manager\")"

inplace_edit "del(${manager_container}.volumeMounts)"

# * Remove the rbac-proxy container.
inplace_edit "del(select($deployment) | .spec.template.spec.containers[] | select(.name == \"kube-rbac-proxy\"))"

# * Remove the webhook port.
inplace_edit "del(${manager_container} | .ports[] | select(.name == \"webhook-server\"))"

# * Set the operator-mode env var to "watchers".
new_env_val='{"name": "AZURE_OPERATOR_MODE", "value": "watchers"}'
inplace_edit "(${manager_container} | .env[] | select(.name == \"AZURE_OPERATOR_MODE\")) = $new_env_val"
