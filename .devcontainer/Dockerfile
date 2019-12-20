#-------------------------------------------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See https://go.microsoft.com/fwlink/?linkid=2090316 for license information.
#-------------------------------------------------------------------------------------------------------------

FROM golang:1.13.5-stretch

# Avoid warnings by switching to noninteractive
ENV DEBIAN_FRONTEND=noninteractive

# Configure apt, install packages and tools
RUN apt-get update \
    && apt-get -y install --no-install-recommends apt-utils dialog nano bash-completion 2>&1 \
    #
    # Verify git, process tools, lsb-release (common in install instructions for CLIs) installed
    && apt-get -y install git iproute2 procps lsb-release \
    # Install docker
    && apt-get install -y apt-transport-https ca-certificates curl gnupg-agent software-properties-common lsb-release \
    # Install icu-devtools (require for vscode liveshare)
    && apt-get install -y icu-devtools \ 
    # Install Azure cli
    && curl -sL https://aka.ms/InstallAzureCLIDeb | bash - \
    # Clean up
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

# Install docker
RUN bash -c "cd /tmp && curl -fsSLO https://download.docker.com/linux/static/stable/x86_64/docker-19.03.5.tgz && tar --strip-components=1 -xvzf docker-19.03.5.tgz -C /usr/local/bin"

# Enable go modules
ENV GO111MODULE=on

# Install K8s Tools
RUN \
    # Install Helm
    curl -s https://raw.githubusercontent.com/helm/helm/master/scripts/get | DESIRED_VERSION=v3.0.1 bash - \
    # Install kubectl
    && curl -sSL -o /usr/local/bin/kubectl https://storage.googleapis.com/kubernetes-release/release/v1.17.0/bin/linux/amd64/kubectl \
    && chmod +x /usr/local/bin/kubectl 

# Install Go tools
RUN \
    # Delve for debugging
    go get github.com/go-delve/delve/cmd/dlv@v1.3.2 \
    # Go language server
    && go get golang.org/x/tools/gopls@v0.2.1 \
    # Go symbols and outline for go to symbol support and test support 
    && go get github.com/acroca/go-symbols@v0.1.1 && go get github.com/ramya-rao-a/go-outline@7182a932836a71948db4a81991a494751eccfe77 \
    # GolangCI-lint
    && curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(go env GOPATH)/bin v1.21.0 \
    # Go releaser 
    && curl -sfL https://install.goreleaser.com/github.com/goreleaser/goreleaser.sh | sh -s -- "v0.123.3"\
    && rm -rf /go/src/ && rm -rf /go/pkg

# Install other stuff using scripts because it's nicer
COPY ./scripts/kubebuilder.sh .
RUN bash -f ./kubebuilder.sh

COPY ./scripts/kustomize.sh .
RUN bash -f ./kustomize.sh

COPY ./scripts/install-kind.sh .
RUN bash -f ./install-kind.sh

COPY ./scripts/configbash.sh .
RUN bash -f ./configbash.sh

# Build the path
ENV PATH="/usr/local/kubebuilder/bin:${PATH}"
ENV PATH="/root/.kubectx:${PATH}"

# Misc config
ENV KIND_CLUSTER_NAME=k8sinfra
ENV KUBE_EDITOR="nano"