#!/bin/bash

# -e immediate exit on error
# -u treat unset variables as an error
set -eu

# This may be run in two modes:
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
# have the devcontainer script pass the argument 
# `devcontainer`
#
# Other available arguments
#
# -v --verbose          : Generate more logging
# -s --skip-installed   : Skip anything that's already installed
#

VERBOSE=false
SKIP=false
DEVCONTAINER=false

while [[ $# -gt 0 ]]; do 
  case $1 in 
    -v | --verbose)
      VERBOSE=true
      shift
      ;;
    -s | --skip-installed)
      SKIP=true
      shift
      ;;
    -* | --*)
      echo "Unknown option $1"
      exit 1
      ;;
    devcontainer)
      DEVCONTAINER=true
      shift
      ;;
    *)
      echo "Unknown parameter $1"
      exit 1
      ;;
  esac
done

write-verbose() {
    if [ "$VERBOSE" = true ]; then
      echo "[VER] $1"
    fi
}

write-info() {
      echo "[INF] $1"
}

write-error() {
    echo "[ERR] $1"
}

# Configure behaviour for devcontainer mode or not

if [ "$DEVCONTAINER" = true ]; then 
    TOOL_DEST=/usr/local/bin
    KUBEBUILDER_DEST=/usr/local/kubebuilder
else
    TOOL_DEST=$(git rev-parse --show-toplevel)/hack/tools
    mkdir -p "$TOOL_DEST"
    KUBEBUILDER_DEST="$TOOL_DEST/kubebuilder"
fi

# Ensure we have the right version of GO

if ! command -v go > /dev/null 2>&1; then
    write-error "Go must be installed manually; see https://golang.org/doc/install"
    exit 1
fi

GOVER=$(go version)
write-info "Go version: ${GOVER[*]}"

GOVERREQUIRED="go1.18.*"
GOVERACTUAL=$(go version | { read _ _ ver _; echo $ver; })
if ! [[ "$GOVERACTUAL" =~ $GOVERREQUIRED ]]; then
    write-error "Go must be version $GOVERREQUIRED, not $GOVERACTUAL; see : https://golang.org/doc/install"
    exit 1
fi

# Ensure we have AZ

if ! command -v az > /dev/null 2>&1; then
    write-error "Azure CLI must be installed manually: https://docs.microsoft.com/en-us/cli/azure/install-azure-cli"
    exit 1
fi

write-verbose "Installing tools to $TOOL_DEST"

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

write-verbose "Installing Go tools…"

# go tools for vscode are preinstalled by base image (see first comment in Dockerfile)

# should-install() is a helper function for deciding whether 
# a given installation is necessary
should-install() {
    if [ "$SKIP" == true ] && [ -f "$1" ]; then 
        # We can skip installation
        return 1
    fi

    # Installation is needed
    return 0
}

# go-install() is a helper function to trigger `go install` 
go-install() {
    write-verbose "Checking for $GOBIN/$1"
    if should-install "$GOBIN/$1"; then 
        write-info "Installing $1"
        shift # Discard the command name so we can pass the remaining arguments to GO
        go install $@
    fi
}

go-install conversion-gen k8s.io/code-generator/cmd/conversion-gen@v0.22.2 
go-install controller-gen sigs.k8s.io/controller-tools/cmd/controller-gen@v0.8.0 
go-install kind sigs.k8s.io/kind@v0.11.1 
go-install kustomize sigs.k8s.io/kustomize/kustomize/v4@v4.5.2 

# for docs site
go-install hugo -tags extended github.com/gohugoio/hugo@v0.88.1
go-install htmltest github.com/wjdp/htmltest@v0.15.0

# for api docs 
# TODO: Replace this with the new release tag.
go-install gen-crd-api-reference-docs github.com/ahmetb/gen-crd-api-reference-docs@v0.3.1-0.20220223025230-af7c5e0048a3

# Install golangci-lint
if [ "$DEVCONTAINER" != true ]; then 
    write-verbose "Checking for $TOOL_DEST/golangci-lint"
    if should-install "$TOOL_DEST/golangci-lint"; then 
        write-info "Installing golangci-lint"
        # golangci-lint is provided by base image if in devcontainer
        # this command copied from there
        curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$TOOL_DEST" v1.46.2 2>&1
    fi
fi

# Install Task
write-verbose "Checking for $TOOL_DEST/go-task"
if should-install "$TOOL_DEST/task"; then 
    write-info "Installing go-task"
    curl -sL "https://github.com/go-task/task/releases/download/v3.12.1/task_linux_amd64.tar.gz" | tar xz -C "$TOOL_DEST" task
fi

# Install binaries for envtest
os=$(go env GOOS)
arch=$(go env GOARCH)
K8S_VERSION=1.23.3
write-verbose "Checking for envtest binaries in $KUBEBUILDER_DEST/bin/kubebuilder"
if should-install "$KUBEBUILDER_DEST/bin/kubebuilder"; then 
    write-info "Installing envtest binaries (kubectl, etcd, kube-apiserver) for ${K8S_VERSION} ($os $arch)…"
    curl -sSLo envtest-bins.tar.gz "https://go.kubebuilder.io/test-tools/${K8S_VERSION}/${os}/${arch}"
    mkdir -p "$KUBEBUILDER_DEST"
    tar -C "$KUBEBUILDER_DEST" --strip-components=1 -zvxf envtest-bins.tar.gz
    rm envtest-bins.tar.gz
fi

# Install helm
write-verbose "Checking for $TOOL_DEST/helm"
if should-install "$TOOL_DEST/helm"; then 
    write-info "Installing helm…"
    curl -sL "https://get.helm.sh/helm-v3.8.0-linux-amd64.tar.gz" | tar -C "$TOOL_DEST" --strip-components=1 -xz linux-amd64/helm
fi

# Install yq
yq_version=v4.13.0
yq_binary=yq_linux_amd64
write-verbose "Checking for $TOOL_DEST/yq"
if should-install "$TOOL_DEST/yq"; then 
    write-info "Installing yq…"
    rm -f "$TOOL_DEST/yq" # remove yq in case we're forcing the install
    wget "https://github.com/mikefarah/yq/releases/download/${yq_version}/${yq_binary}.tar.gz" -O - | tar -xz -C "$TOOL_DEST" && mv "$TOOL_DEST/$yq_binary" "$TOOL_DEST/yq"
fi

# Install cmctl, used to wait for cert manager installation during some tests cases
write-verbose "Checking for $TOOL_DEST/cmctl"
if should-install "$TOOL_DEST/cmctl"; then 
    write-info "Installing cmctl-${os}_${arch}…"
    curl -L "https://github.com/jetstack/cert-manager/releases/latest/download/cmctl-${os}-${arch}.tar.gz" | tar -xz -C "$TOOL_DEST"
fi

if [ "$VERBOSE" == true ]; then 
    echo "Installed tools: $(ls "$TOOL_DEST")"
fi

if [ "$DEVCONTAINER" == true ]; then 
    write-info "Setting up k8s webhook certificates"

    mkdir -p /tmp/k8s-webhook-server/serving-certs
    openssl genrsa 2048 > tls.key
    openssl req -new -x509 -nodes -sha256 -days 3650 -key tls.key -subj '/' -out tls.crt
    mv tls.key tls.crt /tmp/k8s-webhook-server/serving-certs
fi
