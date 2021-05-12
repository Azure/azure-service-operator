# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.137.0/containers/go/.devcontainer/base.Dockerfile

# This is pinned to a particular version of go:
FROM mcr.microsoft.com/vscode/devcontainers/go:0-1.16

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

COPY install-dependencies.sh .
RUN ./install-dependencies.sh devcontainer && rm install-dependencies.sh

# Add kubebuilder to PATH
ENV PATH=$PATH:/usr/local/kubebuilder/bin

# Add further bash customizations
# note that the base image includes oh-my-bash, we are enabling plugins here
# TODO: restore oh-my-bash? it was removed in base image.
# RUN sed -i '/^plugins=/a kubectl\ngolang' "/home/vscode/.bashrc"
# RUN sed -i '/^completions=/a kubectl\ngo\ntask' "/home/vscode/.bashrc"

# Make kubectl completions work with 'k' alias
RUN echo 'complete -F __start_kubectl k' >> "/home/vscode/.bashrc"

# Setup go-task completions
RUN curl -sL "https://raw.githubusercontent.com/go-task/task/v3.0.0/completion/bash/task.bash" > "/home/vscode/.task.completion.sh" \
    && echo 'source /home/vscode/.task.completion.sh' >> /home/vscode/.bashrc

ENV KIND_CLUSTER_NAME=k8sinfra
