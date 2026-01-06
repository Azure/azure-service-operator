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

#
# This file includes documentation in lines prefixed `#doc#`.
# These lines are extracted by running `task doc:dependencies` from the root folder.
#
# Each depencency should be documented by a single line with the following content:
#
# | <name> | <version> | <details> |
#
# Details should include at minimum a link to the originating website. 
# Be sure to use include the `#doc#` prefix on each line.
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
    BUILDX_DEST=/usr/lib/docker/cli-plugins
else
    TOOL_DEST=$(git rev-parse --show-toplevel)/hack/tools
    mkdir -p "$TOOL_DEST"
    KUBEBUILDER_DEST="$TOOL_DEST/kubebuilder"
    BUILDX_DEST=$HOME/.docker/cli-plugins
fi

# Ensure we have the right version of GO

#doc# | Go | 1.23 | https://golang.org/doc/install #
if ! command -v go > /dev/null 2>&1; then
    write-error "Go must be installed manually; see https://golang.org/doc/install"
    exit 1
fi

GOVER=$(go version)
write-info "Go version: ${GOVER[*]}"

GOVERREGEX=".*go1.([0-9]+).([0-9]+).*"
GOVERREQUIRED="go1.24.*"
GOVERACTUAL=$(go version | { read _ _ ver _; echo "$ver"; })

if ! [[ $GOVERACTUAL =~ $GOVERREGEX ]]; then
    write-error "Unexpected Go version format: $GOVERACTUAL"
    exit 1
fi

GOMINORVER="${BASH_REMATCH[1]}"
GOMINORREQUIRED=24

# We allow for Go versions above the min version, but prevent versions below. This is safe given Go's back-compat guarantees
if ! [[ $GOMINORVER -ge $GOMINORREQUIRED ]]; then
    write-error "Go must be version 1.$GOMINORREQUIRED, not $GOVERACTUAL; see : https://golang.org/doc/install"
    exit 1
fi

# Define os and arch
os=$(go env GOOS)
arch=$(go env GOARCH)

# Ensure we have AZ

#doc# | AZ | latest | https://docs.microsoft.com/en-us/cli/azure/install-azure-cli |
if ! command -v az > /dev/null 2>&1; then
    write-error "Azure CLI must be installed manually: https://docs.microsoft.com/en-us/cli/azure/install-azure-cli"
    exit 1
fi

#doc# | Pip3 | latest | https://pip.pypa.io/en/stable/installation/ |
if ! command -v pip3 > /dev/null 2>&1; then
    write-error "Pip3 must be installed manually: https://pip.pypa.io/en/stable/installation/"
    exit 1
fi

write-verbose "Installing tools to $TOOL_DEST"

# Install Go tools
TMPDIR=$(mktemp -d)
clean() { 
    # Macos wants different flag order
    if [[ ${os} == "darwin" ]]; then
        chmod -R +w "$TMPDIR"
    else
        chmod +w -R "$TMPDIR"
    fi
    rm -rf "$TMPDIR"
}
trap clean EXIT

export GOBIN=$TOOL_DEST
export GOPATH=$TMPDIR
export GOCACHE=$TMPDIR/cache
export GO111MODULE=on

write-verbose "Installing Go tools..."

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

#doc# | conversion-gen | v0.34.1 | https://pkg.go.dev/k8s.io/code-generator/cmd/conversion-gen |
go-install conversion-gen k8s.io/code-generator/cmd/conversion-gen@v0.34.1

#doc# | controller-gen | v0.19.0 | https://book.kubebuilder.io/reference/controller-gen |
go-install controller-gen sigs.k8s.io/controller-tools/cmd/controller-gen@v0.19.0

#doc# | kind | v0.30.0 | https://kind.sigs.k8s.io/ |
go-install kind sigs.k8s.io/kind@v0.30.0

#doc# | kustomize | v4.5.7 | https://kustomize.io/ |
go-install kustomize sigs.k8s.io/kustomize/kustomize/v4@v4.5.7

# for docs site

#doc# | hugo | v0.152.2 | https://gohugo.io/ |
go-install hugo -tags extended github.com/gohugoio/hugo@v0.152.2

#doc# | htmltest | latest | https://github.com/wjdp/htmltest (but see https://github.com/theunrepentantgeek/htmltest for our custom build )
# Restore this to github.com/wjdp/htmltest@v?? once PR#215 is merged with the feature we need
go-install htmltest github.com/theunrepentantgeek/htmltest@latest

#doc# | crddoc | ce2ddd3 | https://github.com/theunrepentantgeek/crddoc |
go-install crddoc github.com/theunrepentantgeek/crddoc@ce2ddd3a5f8c148c9d3eeaa0280291875046d5b7

# Install envtest tooling
#doc# | setup-envtest | v0.22.4 | https://book.kubebuilder.io/reference/envtest.html |
write-verbose "Checking for $TOOL_DEST/setup-envtest"
if should-install "$TOOL_DEST/setup-envtest"; then
    write-info "Installing setup-envtest"
    curl -sL "https://github.com/kubernetes-sigs/controller-runtime/releases/download/v0.22.4/setup-envtest-${os}-${arch}" --output "$TOOL_DEST/setup-envtest"
    chmod +x "$TOOL_DEST/setup-envtest"
fi

# Stricter GO formatting
#doc# | gofumpt | latest | https://pkg.go.dev/mvdan.cc/gofumpt |
go-install gofumpt mvdan.cc/gofumpt@latest

# Install golangci-lint
#doc# | golangci-lint | 2.6.1 | https://github.com/golangci/golangci-lint |
write-verbose "Checking for $TOOL_DEST/golangci-lint"
if should-install "$TOOL_DEST/golangci-lint"; then
    write-info "Installing golangci-lint"
    # golangci-lint is provided by base image if in devcontainer
    # this command copied from there
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$TOOL_DEST" v2.6.1 2>&1
fi

# Install Task
#doc# | Task | v3.44.0 | https://taskfile.dev/ |
# not using 3.44.1+ due to https://github.com/go-task/task/issues/2355
write-verbose "Checking for $TOOL_DEST/go-task"
if should-install "$TOOL_DEST/task"; then 
    write-info "Installing go-task"
    curl -sL "https://github.com/go-task/task/releases/download/v3.44.0/task_${os}_${arch}.tar.gz" | tar xz -C "$TOOL_DEST" task
fi

# Install Trivy
#doc# | Trivy | v0.67.2 | https://trivy.dev/ |
write-verbose "Checking for $TOOL_DEST/trivy"
if should-install "$TOOL_DEST/trivy"; then
    write-info "Installing trivy"
    # This guys decided to use different naming conventions for os(go env GOOS) and arch(go env GOARCH) despite trivy is 98.6% written in Go
    # This fixes macos arm64 architechture. Every other os/arch is named differently. Consider adding a workaround of your own ¯\_(ツ)_/¯
    if [[ ${os} == "darwin" ]] && [[ ${arch} == "arm64" ]]; then
        curl -sL "https://github.com/aquasecurity/trivy/releases/download/v0.67.2/trivy_0.67.2_macOS-ARM64.tar.gz" | tar xz -C "$TOOL_DEST" trivy
    else
        curl -sL "https://github.com/aquasecurity/trivy/releases/download/v0.67.2/trivy_0.67.2_Linux-64bit.tar.gz" | tar xz -C "$TOOL_DEST" trivy
    fi
fi

# Install helm
#doc# | Helm | v3.19.0 | https://helm.sh/ |
write-verbose "Checking for $TOOL_DEST/helm"
if should-install "$TOOL_DEST/helm"; then
    write-info "Installing helm..."
    curl -sL "https://get.helm.sh/helm-v3.19.0-${os}-${arch}.tar.gz" | tar -C "$TOOL_DEST" --strip-components=1 -xz ${os}-${arch}/helm
fi

# Install yq
#doc# | YQ | v4.48.1 | https://github.com/mikefarah/yq/ |
yq_version=v4.48.1
yq_binary=yq_${os}_${arch}
write-verbose "Checking for $TOOL_DEST/yq"
if should-install "$TOOL_DEST/yq"; then 
    write-info "Installing yq..."
    rm -f "$TOOL_DEST/yq" # remove yq in case we're forcing the install
    wget "https://github.com/mikefarah/yq/releases/download/${yq_version}/${yq_binary}.tar.gz" -O - | tar -xz -C "$TOOL_DEST" && mv "$TOOL_DEST/$yq_binary" "$TOOL_DEST/yq"
fi

# Install cmctl, used to wait for cert manager installation during some tests cases
#doc# | cmctl | latest | https://cert-manager.io/docs/reference/cmctl |
write-verbose "Checking for $TOOL_DEST/cmctl"
if should-install "$TOOL_DEST/cmctl"; then 
    write-info "Installing cmctl-${os}_${arch}..."
    curl -fL "https://github.com/cert-manager/cmctl/releases/latest/download/cmctl_${os}_${arch}" --output "$TOOL_DEST/cmctl"
    chmod +x "$TOOL_DEST/cmctl"
fi

write-verbose "Checking for $BUILDX_DEST/docker-buildx"
#doc# | BuildX | v0.29.1 | https://github.com/docker/buildx |
if should-install "$BUILDX_DEST/docker-buildx"; then
    write-info "Installing buildx-${os}_${arch} to $BUILDX_DEST ..."
    if ! test -f $BUILDX_DEST; then
        mkdir -p "$BUILDX_DEST"
    fi
    if ! test -f $BUILDX_DEST/docker-buildx; then
        curl  -o "$BUILDX_DEST/docker-buildx" -L "https://github.com/docker/buildx/releases/download/v0.29.1/buildx-v0.29.1.${os}-${arch}"
        chmod +x "$BUILDX_DEST/docker-buildx"
    fi
fi

# Install azwi
#doc# | AZWI | v1.5.1 | https://github.com/Azure/azure-workload-identity |
write-verbose "Checking for $TOOL_DEST/azwi"
if should-install "$TOOL_DEST/azwi"; then
    write-info "Installing azwi..."
    curl -sL "https://github.com/Azure/azure-workload-identity/releases/download/v1.5.1/azwi-v1.5.1-${os}-${arch}.tar.gz" | tar xz -C "$TOOL_DEST" azwi
fi

# Ensure tooling for Hugo is available
#doc# | PostCSS | latest | https://postcss.org/ |
write-verbose "Checking for /usr/bin/postcss"
if ! which postcss  > /dev/null 2>&1; then 
    write-info "Installing postcss"
    npm config set fund false --location=global
    npm install --global postcss postcss-cli autoprefixer
fi

if [ "$VERBOSE" == true ]; then 
    echo "Installed tools: $(ls "$TOOL_DEST")"
fi

if [ "$DEVCONTAINER" == true ]; then

    # Webhook Certs
    write-info "Setting up k8s webhook certificates"
    mkdir -p /tmp/k8s-webhook-server/serving-certs
    openssl genrsa 2048 > tls.key
    openssl req -new -x509 -nodes -sha256 -days 3650 -key tls.key -subj '/' -out tls.crt
    mv tls.key tls.crt /tmp/k8s-webhook-server/serving-certs

    # Git Permissions
    # Workaround for issue where /workspace has different owner because checkout happens outside the container
    git config --global --add safe.directory /workspace
fi

write-info "Setting up python virtual environment"

# Install python virtualenv
pip3 install virtualenv --break-system-packages
