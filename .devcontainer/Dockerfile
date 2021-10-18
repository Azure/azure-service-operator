# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.137.0/containers/go/.devcontainer/base.Dockerfile

# This is pinned to a particular version of go:
FROM mcr.microsoft.com/vscode/devcontainers/go:0-1.18

# APT dependencies
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update \
    && apt-get -y install --no-install-recommends bash-completion software-properties-common lsb-release \
    # install az-cli
    && curl -sL https://aka.ms/InstallAzureCLIDeb | bash -

# For Hugo we need a newer version of NodeJS than the v12.22.12 available via the standard sources
RUN curl -fsSL https://deb.nodesource.com/setup_18.x -o /tmp/nodesource-setup.sh \
    && sudo bash /tmp/nodesource-setup.sh \
    && apt-get -y install nodejs

# install docker
# - not yet needed?
# RUN curl -fsSL https://get.docker.com | sh -
# RUN usermod -aG docker vscode

COPY install-dependencies.sh .
RUN ./install-dependencies.sh devcontainer && rm install-dependencies.sh

# Setup envtest
# NB: if you change this, dev.sh also likely needs updating
USER vscode
# - install binaries (these go to HOME which is why we switch user)
RUN setup-envtest use 1.23.x
# - add setup to .bashrc; first the KUBEBUILDER_ASSETS variable
RUN echo 'source <(setup-envtest use -i -p env)' >> "/home/vscode/.bashrc"
# - also add KUBEBUILDER_ASSETS to path
RUN echo 'export PATH=$KUBEBUILDER_ASSETS:$PATH' >> "/home/vscode/.bashrc"
USER root

# Make kubectl completions work with 'k' alias
RUN echo 'alias k=kubectl' >> "/home/vscode/.bashrc"
RUN echo 'complete -F __start_kubectl k' >> "/home/vscode/.bashrc"
RUN echo 'source <(kubectl completion bash)' >> "/home/vscode/.bashrc"

# Setup go-task completions
RUN curl -sL "https://raw.githubusercontent.com/go-task/task/v3.0.0/completion/bash/task.bash" > "/home/vscode/.task.completion.sh" \
    && echo 'source /home/vscode/.task.completion.sh' >> /home/vscode/.bashrc

ENV KIND_CLUSTER_NAME=aso

# install docker, from: https://github.com/microsoft/vscode-dev-containers/blob/main/script-library/docs/docker.md
COPY library-scripts/docker-debian.sh /tmp/library-scripts/

# these are all the default values except for the last one (USE_MOBY) which is what we want to set.
# currently packages.microsoft.com does not have Moby packages for Debian Bullseye, which the
# devcontainer image is based on: https://github.com/microsoft/vscode-dev-containers/issues/1008
RUN bash /tmp/library-scripts/docker-debian.sh true /var/run/docker-host.sock /var/run/docker.sock automatic false

ENTRYPOINT ["/usr/local/share/docker-init.sh"]
CMD ["sleep", "infinity"]
