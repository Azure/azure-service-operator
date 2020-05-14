#-------------------------------------------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See https://go.microsoft.com/fwlink/?linkid=2090316 for license information.
#-------------------------------------------------------------------------------------------------------------

FROM golang:1.13.10-buster

# Configure apt, install packages and cli tools
RUN export DEBIAN_FRONTEND=noninteractive \
    && apt-get update \
    && apt-get -y install --no-install-recommends apt-transport-https ca-certificates curl gnupg2 lsb-release \
    # 
    # Install Kubebuilder
    && curl -sL https://go.kubebuilder.io/dl/2.3.1/$(go env GOOS)/$(go env GOARCH) | tar -xz --strip-components=2 -C /usr/local/bin \
    #
    # Install Helm
    && curl -s https://get.helm.sh/helm-v3.2.1-linux-amd64.tar.gz | tar -zxv --strip-components=1 -C /usr/local/bin linux-amd64/helm \
    #
    # Install Docker CLI package source
    && curl -fsSL https://download.docker.com/linux/$(lsb_release -is | tr '[:upper:]' '[:lower:]')/gpg | (OUT=$(apt-key add - 2>&1) || echo $OUT) \
    && echo "deb [arch=amd64] https://download.docker.com/linux/$(lsb_release -is | tr '[:upper:]' '[:lower:]') $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list \
    #
    # Add the Azure CLI package source
    && curl -sL https://packages.microsoft.com/keys/microsoft.asc | apt-key add - 2>/dev/null \
    && echo "deb [arch=amd64] https://packages.microsoft.com/repos/azure-cli/ $(lsb_release -cs) main" > /etc/apt/sources.list.d/azure-cli.list \
    #
    # Install Packages
    && apt-get update \
    && apt-get install -y --no-install-recommends \
        docker-ce-cli \
        azure-cli \
    #
    # Clean up
    && apt-get autoremove -y \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

# Install Go development tools
RUN mkdir -p /tmp/gotools \
    && cd /tmp/gotools \
    #
    # Build Go tools w/module support
    && GOPATH=/tmp/gotools GO111MODULE=on go get -v \
        golang.org/x/tools/gopls@latest \
        honnef.co/go/tools/...@latest \
        golang.org/x/tools/cmd/gorename@latest \
        golang.org/x/tools/cmd/goimports@latest \
        golang.org/x/tools/cmd/guru@latest \
        golang.org/x/lint/golint@latest \
        github.com/mdempsky/gocode@latest \
        github.com/cweill/gotests/...@latest \
        github.com/haya14busa/goplay/cmd/goplay@latest \
        github.com/sqs/goreturns@latest \
        github.com/josharian/impl@latest \
        github.com/davidrjenni/reftools/cmd/fillstruct@latest \
        github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest  \
        github.com/ramya-rao-a/go-outline@latest  \
        github.com/acroca/go-symbols@latest  \
        github.com/godoctor/godoctor@latest  \
        github.com/rogpeppe/godef@latest  \
        github.com/zmb3/gogetdoc@latest \
        github.com/fatih/gomodifytags@latest  \
        github.com/mgechev/revive@latest  \
        github.com/go-delve/delve/cmd/dlv@latest 2>&1 \
    #
    # Build Go tools w/o module support
    && GOPATH=/tmp/gotools go get -v github.com/alecthomas/gometalinter 2>&1 \
    #
    # Build gocode-gomod
    && GOPATH=/tmp/gotools go get -x -d github.com/stamblerre/gocode 2>&1 \
    && GOPATH=/tmp/gotools go build -o gocode-gomod github.com/stamblerre/gocode \
    #
    # Install Go tools
    && mv /tmp/gotools/bin/* /usr/local/bin/ \
    && mv gocode-gomod /usr/local/bin/ \
    #
    # Install golangci-lint
    && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/local/bin 2>&1 \
    #
    # Clean up
    && rm -rf /tmp/gotools

# Copy shares from localhost into container when new shells are created
RUN echo '\n\
        if [ -d "/usr/local/share/kube-localhost" ]; then\n\
            mkdir -p $HOME/.kube\n\
            cp -r /usr/local/share/kube-localhost/* $HOME/.kube\n\
            chown -R $(id -u) $HOME/.kube\n\
            sed -i -e "s/localhost/host.docker.internal/g" $HOME/.kube/config\n\
        \n\
            if [ -d "/usr/local/share/minikube-localhost" ]; then\n\
                mkdir -p $HOME/.minikube\n\
                cp -r /usr/local/share/minikube-localhost/ca.crt $HOME/.minikube\n\
                cp -r /usr/local/share/minikube-localhost/client.crt $HOME/.minikube\n\
                cp -r /usr/local/share/minikube-localhost/client.key $HOME/.minikube\n\
                chown -R $(id -u) $HOME/.minikube\n\
                sed -i -r "s|(\s*certificate-authority:\s).*|\\1$HOME\/.minikube\/ca.crt|g" $HOME/.kube/config\n\
                sed -i -r "s|(\s*client-certificate:\s).*|\\1$HOME\/.minikube\/client.crt|g" $HOME/.kube/config\n\
                sed -i -r "s|(\s*client-key:\s).*|\\1$HOME\/.minikube\/client.key|g" $HOME/.kube/config\n\
            fi\n\
        fi\n\
        if [ -d "/usr/local/share/azure-localhost" ]; then\n\
            mkdir -p $HOME/.azure\n\
            cp -r /usr/local/share/azure-localhost/* $HOME/.azure\n\
            chown -R $(id -u) $HOME/.azure\n\
        fi' | tee -a /root/.bashrc /root/.zshrc /home/${USERNAME}/.bashrc >> /home/${USERNAME}/.zshrc
