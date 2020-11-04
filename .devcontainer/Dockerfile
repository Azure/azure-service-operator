# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.137.0/containers/go/.devcontainer/base.Dockerfile
ARG VARIANT="1.15"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# APT dependencies
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update \
    && apt-get -y install --no-install-recommends bash-completion software-properties-common lsb-release \
    # install az-cli
    && curl -sL https://aka.ms/InstallAzureCLIDeb | bash -

# install docker
# - not yet needed?
# RUN curl -fsSL https://get.docker.com | sh -
# RUN usermod -aG docker vscode

# Go dependencies
ENV GO111MODULE=on
RUN mkdir -p /tmp/gotools \
    && export GOPATH=/tmp/gotools \
    && export GOCACHE=/tmp/gotools/cache \
    && go get \
        # go tools for vscode are preinstalled by base image (see first comment in Dockefile)
        # golangci-lint is provided by base image
        github.com/jandelgado/gcov2lcov@v1.0.2 \
        github.com/mitchellh/gox@v1.0.1 \
        k8s.io/code-generator/cmd/conversion-gen@v0.18.2 \
        sigs.k8s.io/controller-tools/cmd/controller-gen@v0.4.0 \
        sigs.k8s.io/kind@v0.9.0 \
        sigs.k8s.io/kustomize/kustomize/v3@v3.8.6 \
    && mv /tmp/gotools/bin/* /usr/local/bin/ \
    && rm -rf /tmp/gotools

# install kubectl
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.19.0/bin/linux/amd64/kubectl \
    && chmod +x ./kubectl \
    && mv ./kubectl /usr/local/bin/kubectl

# install go-task (as task runner)
RUN curl -sL "https://github.com/go-task/task/releases/download/v3.0.0/task_linux_amd64.tar.gz" | tar xz -C /usr/local/bin task
RUN curl -sL "https://raw.githubusercontent.com/go-task/task/v3.0.0/completion/bash/task.bash" > "/home/vscode/.oh-my-bash/completions/task.completion.sh"

# install kubebuilder
RUN os=$(go env GOOS) \
    && arch=$(go env GOARCH) \
    && curl -L https://go.kubebuilder.io/dl/2.3.1/${os}/${arch} | tar -xz -C /tmp/ \
    && mv /tmp/kubebuilder_2.3.1_${os}_${arch} /usr/local/kubebuilder
ENV PATH=$PATH:/usr/local/kubebuilder/bin

# add bash customizations
# note that the base image includes oh-my-bash, we are enabling plugins here
RUN sed -i '/^plugins=/a kubectl\ngolang' "/home/vscode/.bashrc"
RUN sed -i '/^completions=/a kubectl\ngo\ntask' "/home/vscode/.bashrc"
# make kubectl completions work with 'k' alias
RUN echo 'complete -F __start_kubectl k' >> "/home/vscode/.bashrc"

ENV KIND_CLUSTER_NAME=k8sinfra
