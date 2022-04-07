#!/bin/sh

set -eu

# This script must run in two modes:
#
# - When being used to set up a devcontainer.
#   In this mode the code is not checked out yet,
#   and we can install the tools globally.
#
# - When being used to install tools locally.
#   In this mode the code is already checked out,
#   and we do not want to pollute the user’s system.
#
# To distinguish between these modes we will
# have the devcontainer script pass an argument:

if [ "$1" = "devcontainer" ]; then 
    TOOL_DEST=/usr/local/bin
    KUBEBUILDER_DEST=/usr/local/kubebuilder
else
    TOOL_DEST=$(git rev-parse --show-toplevel)/hack/tools
    mkdir -p "$TOOL_DEST"
    KUBEBUILDER_DEST="$TOOL_DEST/kubebuilder"
fi

if ! command -v go > /dev/null 2>&1; then
    echo "Go must be installed manually: https://golang.org/doc/install"
    exit 1
fi

if ! command -v az > /dev/null 2>&1; then
    echo "Azure CLI must be installed manually: https://docs.microsoft.com/en-us/cli/azure/install-azure-cli"
    exit 1
fi

echo "Installing tools to $TOOL_DEST"

# Install Go tools
TMPDIR=$(mktemp -d)
clean() { 
    chmod +w -R "$TMPDIR"
    rm -rf "$TMPDIR"
}
trap clean EXIT

export GOBIN=$TOOL_DEST
export GOPATH=$TMPDIR
export GOCACHE=$TMPDIR/cache
export GO111MODULE=on

echo "Installing Go tools…"

# go tools for vscode are preinstalled by base image (see first comment in Dockerfile)
go install k8s.io/code-generator/cmd/conversion-gen@v0.22.2 
go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.8.0 
go install sigs.k8s.io/kind@v0.11.1 
go install sigs.k8s.io/kustomize/kustomize/v4@v4.5.2 

# for docs site
go install -tags extended github.com/gohugoio/hugo@v0.88.1
go install github.com/wjdp/htmltest@v0.15.0

# for api docs 
# TODO: Replace this with the new release tag.
go install github.com/ahmetb/gen-crd-api-reference-docs@v0.3.1-0.20220223025230-af7c5e0048a3

if [ "$1" != "devcontainer" ]; then 
    echo "Installing golangci-lint…"
    # golangci-lint is provided by base image if in devcontainer
    # this command copied from there
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$TOOL_DEST" v1.45.2 2>&1
fi

# Install go-task (task runner)
echo "Installing go-task…"
curl -sL "https://github.com/go-task/task/releases/download/v3.7.0/task_linux_amd64.tar.gz" | tar xz -C "$TOOL_DEST" task

# Install binaries for envtest
os=$(go env GOOS)
arch=$(go env GOARCH)
K8S_VERSION=1.23.3
echo "Installing envtest binaries (kubectl, etcd, kube-apiserver) for ${K8S_VERSION} ($os $arch)…"
curl -sSLo envtest-bins.tar.gz "https://go.kubebuilder.io/test-tools/${K8S_VERSION}/${os}/${arch}"
mkdir -p "/usr/local/kubebuilder"
tar -C /usr/local/kubebuilder --strip-components=1 -zvxf envtest-bins.tar.gz
rm envtest-bins.tar.gz

# Install helm
echo "Installing helm…"
curl -sL "https://get.helm.sh/helm-v3.8.0-linux-amd64.tar.gz" | tar -C "$TOOL_DEST" --strip-components=1 -xz linux-amd64/helm

# Install yq
echo "Installing yq…"
yq_version=v4.13.0
yq_binary=yq_linux_amd64
wget "https://github.com/mikefarah/yq/releases/download/${yq_version}/${yq_binary}.tar.gz" -O - | tar -xz -C "$TOOL_DEST" && mv "$TOOL_DEST/$yq_binary" "$TOOL_DEST/yq"

# Install cmctl, used to wait for cert manager installation during some tests cases
echo "Installing cmctl-${os}_${arch}…"
curl -L "https://github.com/jetstack/cert-manager/releases/latest/download/cmctl-${os}-${arch}.tar.gz" | tar -xz -C "$TOOL_DEST"

echo "Installed tools: $(ls "$TOOL_DEST")"


if [ "$1" = "devcontainer" ]; then 
    echo "Setting up k8s webhook certificates"

    mkdir -p /tmp/k8s-webhook-server/serving-certs
    openssl genrsa 2048 > tls.key
    openssl req -new -x509 -nodes -sha256 -days 3650 -key tls.key -subj '/' -out tls.crt
    mv tls.key tls.crt /tmp/k8s-webhook-server/serving-certs
fi
