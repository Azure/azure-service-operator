#!/usr/bin/env bash
set -euxo pipefail

crd='.kind == "CustomResourceDefinition"'
namespace='.kind == "Namespace"'
cluster_role='.kind == "ClusterRole"'
leader_election_role='(.kind == "Role" and .metadata.name == "azureserviceoperator-leader-election-role")'
leader_election_binding='(.kind == "RoleBinding" and .metadata.name == "azureserviceoperator-leader-election-rolebinding")'
proxy_binding='(.kind == "ClusterRoleBinding" and .metadata.name == "azureserviceoperator-proxy-rolebinding")'
service='.kind == "Service"'
deployment='.kind == "Deployment"'
certificate='.kind == "Certificate"'
issuer='.kind == "Issuer"'
webhooks='(.kind == "MutatingWebhookConfiguration" or .kind == "ValidatingWebhookConfiguration")'
query="select($crd or $namespace or $cluster_role or $leader_election_role or $leader_election_binding or $service or $deployment or $certificate or $issuer or $webhooks or $proxy_binding)"

version="$1"
target="multitenant-cluster_${version}.yaml"
source="$2"

yq eval "$query" "$source" > "$target"
