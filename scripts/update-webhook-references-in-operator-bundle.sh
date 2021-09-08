#!/usr/bin/env bash
set -euo pipefail

# Find all of the CRDs that have conversion webhooks defined
# (filtering out the --- separators yq puts between them).
query='select(.spec.conversion.webhook.clientConfig.service.namespace == "azureoperator-system") | filename'
webhook_crds=$(yq eval "$query" bundle/manifests/azure.microsoft.com_*.yaml | grep -v -e "---")

# Update the service details to point at
# operators/azureoperator-controller-manager-service and remove the
# cert-manager annotation.
update='.spec.conversion.webhook.clientConfig.service.namespace = "operators" | .spec.conversion.webhook.clientConfig.service.name = "azureoperator-controller-manager-service" | del(.metadata.annotations["cert-manager.io/inject-ca-from"])'
for fname in $webhook_crds; do
    yq -i eval "$update" "$fname"
done
