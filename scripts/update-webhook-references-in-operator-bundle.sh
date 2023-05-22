#!/usr/bin/env bash
set -euo pipefail

# Find all of the CRDs that have conversion webhooks defined
# (filtering out the --- separators yq puts between them).
query='select(.spec.conversion.webhook.clientConfig.service.namespace == "azureserviceoperator-system") | filename'
webhook_crds=$(yq eval "$query" bundle/manifests/*.azure.com_*.yaml | grep -v -e "---")

# Remove the cert-manager annotation and conversion details from CRDs
# with conversion webhooks - OLM will set up the conversion structure
# for the webhooks.
update='del(.spec.conversion) | del(.metadata.annotations["cert-manager.io/inject-ca-from"])'
for fname in $webhook_crds; do
    yq -i eval "$update" "$fname"
done
