#! /bin/bash 
set -e
set -x

# download kustomize
curl -o /usr/local/kubebuilder/bin/kustomize -sL "https://go.kubebuilder.io/kustomize/$(go env GOOS)/$(go env GOARCH)"
# set permission
chmod a+x /usr/local/kubebuilder/bin/kustomize