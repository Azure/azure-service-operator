#-------------------------------------------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See https://go.microsoft.com/fwlink/?linkid=2090316 for license information.
#-------------------------------------------------------------------------------------------------------------

FROM golang:1.12.5

# Avoid warnings by switching to noninteractive
ENV DEBIAN_FRONTEND=noninteractive

# Configure apt, install packages and tools
RUN apt-get update \
    && apt-get -y install --no-install-recommends apt-utils 2>&1 \
    #
    # Verify git, process tools, lsb-release (common in install instructions for CLIs) installed
    && apt-get -y install git procps lsb-release \
    #
    # Install gocode-gomod
    && go get -x -d github.com/stamblerre/gocode 2>&1 \
    && go build -o gocode-gomod github.com/stamblerre/gocode \
    && mv gocode-gomod $GOPATH/bin/ \
    #
    # Install Go tools
    && go get -u -v \
        github.com/mdempsky/gocode \
        github.com/uudashr/gopkgs/cmd/gopkgs \
        github.com/ramya-rao-a/go-outline \
        github.com/acroca/go-symbols \
        github.com/godoctor/godoctor \
        golang.org/x/tools/cmd/guru \
        golang.org/x/tools/cmd/gorename \
        github.com/rogpeppe/godef \
        github.com/zmb3/gogetdoc \
        github.com/haya14busa/goplay/cmd/goplay \
        github.com/sqs/goreturns \
        github.com/josharian/impl \
        github.com/davidrjenni/reftools/cmd/fillstruct \
        github.com/fatih/gomodifytags \
        github.com/cweill/gotests/... \
        golang.org/x/tools/cmd/goimports \
        golang.org/x/lint/golint \
        golang.org/x/tools/cmd/gopls \
        github.com/alecthomas/gometalinter \
        honnef.co/go/tools/... \
        github.com/golangci/golangci-lint/cmd/golangci-lint \
        github.com/mgechev/revive \
        github.com/derekparker/delve/cmd/dlv 2>&1 \
    #
    # Clean up
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

RUN apt-get update \
    #
    # Install Docker CE CLI
    && apt-get install -y apt-transport-https ca-certificates curl gnupg-agent software-properties-common lsb-release \
    && curl -fsSL https://download.docker.com/linux/$(lsb_release -is | tr '[:upper:]' '[:lower:]')/gpg | apt-key add - 2>/dev/null \
    && add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/$(lsb_release -is | tr '[:upper:]' '[:lower:]') $(lsb_release -cs) stable" \
    && apt-get update \
    && apt-get install -y docker-ce-cli \
    #
    # Install kubectl
    && curl -sSL -o /usr/local/bin/kubectl https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl \
    && chmod +x /usr/local/bin/kubectl \
    #
    # Install Helm
    && curl -s https://raw.githubusercontent.com/helm/helm/master/scripts/get | bash -

RUN os=$(go env GOOS)  \
&& arch=$(go env GOARCH)  \
# download kubebuilder and extract it to tmp
&& curl -sL https://go.kubebuilder.io/dl/2.0.0-rc.0/${os}/${arch} | tar -xz -C /tmp/   \
# move to a long-term location and put it on your path
# (you'll need to set the KUBEBUILDER_ASSETS env var if you put it somewhere else)
&& mv /tmp/kubebuilder_2.0.0-rc.0_${os}_${arch} /usr/local/kubebuilder   \
# && export PATH=$PATH:/usr/local/kubebuilder/bin \
&& chmod +x /usr/local/kubebuilder/bin 



ENV PATH="/usr/local/kubebuilder/bin:${PATH}"

ENV GO111MODULE=on


# Set the default shell to bash instead of sh
ENV AZURE_CLIENT_ID=""
ENV AZURE_CLIENT_SECRET=""
ENV AZURE_SUBSCRIPTION_ID=""
ENV AZURE_TENANT_ID=""
ENV KUBECONFIG=/root/.kube/config

ENV SHELL /bin/bash

RUN echo '\n\
if [ "$SYNC_LOCALHOST_KUBECONFIG" == "true" ]; then\n\
    mkdir -p $HOME/.kube\n\
    cp -r $HOME/.kube-localhost/* $HOME/.kube\n\
    sed -i -e "s/localhost/host.docker.internal/g" $HOME/.kube/config\n\
\n\
    if [ -d "$HOME/.minikube-localhost" ]; then\n\
        mkdir -p $HOME/.minikube\n\
        cp -r $HOME/.minikube-localhost/ca.crt $HOME/.minikube\n\
        sed -i -r "s|(\s*certificate-authority:\s).*|\\1$HOME\/.minikube\/ca.crt|g" $HOME/.kube/config\n\
        cp -r $HOME/.minikube-localhost/client.crt $HOME/.minikube\n\
        sed -i -r "s|(\s*client-certificate:\s).*|\\1$HOME\/.minikube\/client.crt|g" $HOME/.kube/config\n\
        cp -r $HOME/.minikube-localhost/client.key $HOME/.minikube\n\
        sed -i -r "s|(\s*client-key:\s).*|\\1$HOME\/.minikube\/client.key|g" $HOME/.kube/config\n\
    fi\n\
fi' \
>> $HOME/.bashrc
