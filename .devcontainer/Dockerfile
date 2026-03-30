# See here for image contents: https://github.com/devcontainers/images/blob/main/src/go/.devcontainer/Dockerfile

# Compound versioning - v2 of the devcontainer, v1.26 of Go. 
# Should be automatically kept up to date by dependabot.
FROM mcr.microsoft.com/devcontainers/go:2-1.26

# https://docs.docker.com/engine/reference/builder/#automatic-platform-args-in-the-global-scope

ARG TARGETARCH

# APT dependencies
# including docker-cli for docker-from-docker pattern
# and az-cli for interacting with azure resources
ENV DEBIAN_FRONTEND=noninteractive
RUN curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg \
    && echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/debian $(lsb_release -cs) stable" > /etc/apt/sources.list.d/docker.list \
    && apt-get update \
    && apt-get -y install --no-install-recommends bash-completion lsb-release graphviz zip nodejs npm python3-pip docker-ce-cli docker-compose-plugin gnuplot\
    # install az-cli
    && curl -sL https://aka.ms/InstallAzureCLIDeb | bash - \
    # Temporary fix to avoid regression in v2.77
    && apt-get remove azure-cli --assume-yes \
    && apt-get -y install azure-cli=2.76.0-1~bookworm \
    && rm -rf /var/lib/apt/lists/*

COPY .devcontainer/install-dependencies.sh .
RUN ./install-dependencies.sh devcontainer && rm install-dependencies.sh

# Pre-download Go modules for faster builds
# Some of these go.mod files have replace directives that point to local directories, so we need to be careful to 
# maintain the same relative positions as we copy them into the container.
# v2 first (largest, most stable, needed by replace directives)
COPY v2/go.mod v2/go.sum /tmp/mod-download/v2/
# asoctl (has replace => ../../ which resolves to v2/)
COPY v2/cmd/asoctl/go.mod v2/cmd/asoctl/go.sum /tmp/mod-download/v2/cmd/asoctl/
# generator (has replace => ../../ which resolves to v2/)
COPY v2/tools/generator/go.mod v2/tools/generator/go.sum /tmp/mod-download/v2/tools/generator/
# mangle-test-json (standalone, no replace directives)
COPY v2/tools/mangle-test-json/go.mod v2/tools/mangle-test-json/go.sum /tmp/mod-download/v2/tools/mangle-test-json/
# Now do all the downloads
RUN find /tmp/mod-download -name go.mod -execdir pwd \; -execdir go mod download \;

# Clean up temporary module download directory
RUN rm -rf /tmp/mod-download

# Setup envtest
# NB: if you change this, dev.sh also likely needs updating, also need to update the env below
RUN setup-envtest use 1.35.0 --bin-dir /usr/local/envtest/bin
# Set the KUBEBUILDER_ASSETS variable. Ideally we could do source <(setup-envtest use --bin-dir /usr/local/envtest/bin -i -p env)
# but there's no way to dynamically set an env variable for all container users.
# Many usages of this container are from docker exec task <x>, where no shell is invoked and the entrypoint is not run
# (entrypoint is only run on start, not on exec). Due to that, the following approaches do not work:
# - ~/.bashrc - only works for one user in a shell but we must support -u $(id -u ${USER}):$(id -g ${USER}) which means the container could run as more than 1 user
# - /etc/profile or /etc/profile.d - only works for one user in a login shell
ENV KUBEBUILDER_ASSETS=/usr/local/envtest/bin/k8s/1.35.0-linux-${TARGETARCH}
ENV PATH=$KUBEBUILDER_ASSETS:$PATH

# Make kubectl completions work with 'k' alias
RUN echo 'alias k=kubectl' >> "/etc/bash.bashrc"
RUN echo 'complete -F __start_kubectl k' >> "/etc/bash.bashrc"
RUN echo 'source <(kubectl completion bash)' >> "/etc/bash.bashrc"

# Setup go-task completions
RUN curl -sL "https://raw.githubusercontent.com/go-task/task/v3.2.0/completion/bash/task.bash" > "/etc/.task.completion.sh" \
    && echo 'source /etc/.task.completion.sh' >> "/etc/bash.bashrc"

ENV KIND_CLUSTER_NAME=aso

# Add user to docker group for docker socket access
RUN groupadd -f docker && usermod -aG docker vscode

CMD ["sleep", "infinity"]
