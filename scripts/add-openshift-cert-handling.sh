#!/usr/bin/env bash
set -euo pipefail


# Add ` service.beta.openshift.io/inject-cabundle: "true"` annotation
# for CRDs requiring conversion.
for fname in $(grep cert-manager.io bundle/manifests/* -l); do
    sed -i '/cert-manager.io\/inject-ca-from/a\    service.beta.openshift.io/inject-cabundle: "true"' $fname
done
cat <<EOF > bundle/manifests/service.yaml
kind: Service
metadata:
  name: azureoperator-webhook-service
  annotations:
    service.alpha.openshift.io/serving-cert-secret-name: webhook-server-cert
spec:
  ports:
    - port: 443
      protocol: TCP
      targetPort: 9443
  selector:
    control-plane: controller-manager
  sessionAffinity: None
  type: ClusterIP
EOF
