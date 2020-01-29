#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

bins="./hack/tools/bin"
cfssl="$bins/cfssl"
cfssljson="$bins/cfssljson"
k="$bins/kubectl"

$k delete csr/webhook.local --ignore-not-found=true

mkdir -p ./pki/certs

$cfssl genkey ./pki/webhook-csr.json | $cfssljson -bare ./pki/certs/tls

cat <<EOF | $k apply -f -
apiVersion: certificates.k8s.io/v1beta1
kind: CertificateSigningRequest
metadata:
  name: webhook.local
spec:
  request: $(cat ./pki/certs/tls.csr | base64 | tr -d '\n')
  usages:
  - digital signature
  - key encipherment
  - server auth
EOF

$k certificate approve webhook.local

#$k wait --for=condition=Issued csr/webhook.local --timeout=10s

sleep 2

cp ./pki/certs/tls-key.pem ./pki/certs/tls.key

$k get csr webhook.local -o jsonpath='{.status.certificate}' | base64 --decode > ./pki/certs/tls.crt