#!/usr/bin/env bash

# This artifact use the temp disk (SSD) for docker and optionally containerd if a temp disk is available
# This was copied from https://devdiv.visualstudio.com/_git/XlabImageFactory?path=%2Fartifacts%2Flinux-docker-use-ssd%2Flinux-docker-use-ssd.sh
# We should consider using a 1ES managed image so we can just use that script directly.

set -e

# check if temp disk is mounted at /mnt and if not exit early
if ! mountpoint -q /mnt; then
    echo "INFO: No temp disk mounted at /mnt, skipping artifact"
    exit 0
fi

# check if docker is installed
if ! command -v docker &> /dev/null; then
    echo "ERROR: Docker is not installed" 1>&2
    exit 1
fi

# check if jq is installed
if ! command -v jq &> /dev/null; then
    echo "ERROR: jq is not installed" 1>&2
    exit 1
fi

# Print the current disk usage, useful for diagnosing what's gone wrong if
# the disk fills up
df -h

# Uncomment the below to show more disk stats (including what files are large).
# Commented out as it can take a long time to run and slows CI down
# du -h --threshold=1G --max-depth=5 / 2>&1 | grep -v 'denied' | sort -hr

CONTAINERD="false"
while [[ $# -gt 0 ]]; do
    case $1 in
    --containerd)
        CONTAINERD="$2"
        echo -e "INFO: containerd: ${CONTAINERD}"
        shift
        shift
        ;;
    *)
      echo "Error: Unknown argument \"$1\"" 1>&2
      exit 1
      ;;
    esac
done

if [[ "${CONTAINERD,,}" == "true" ]]; then
    CONTAINERD_CONFIG="/etc/containerd/config.toml"
    CONTAINERD_ROOT="/mnt/containerd"

    if ! command -v containerd &> /dev/null; then
        echo "ERROR: containerd is not installed, skipping artifact" 1>&2
        exit 1
    fi

    mkdir -p "${CONTAINERD_ROOT}"

    existing_config=$(if [ -f "${CONTAINERD_CONFIG}" ]; then cat ${CONTAINERD_CONFIG}; else echo ""; fi)

    # Gets all lines of the config except the one that starts with root and adds the root line to the top of the file
    echo "${existing_config}" | grep -v "^root" | sed "1i root = \"${CONTAINERD_ROOT}\"" > ${CONTAINERD_CONFIG}

    systemctl restart containerd
fi

DAEMON_JSON="/etc/docker/daemon.json"
DOCKER_ROOT="/mnt/docker"

if [ ! -f "${DAEMON_JSON}" ]; then
    echo "INFO: Creating ${DAEMON_JSON}"
    echo "{}" > "${DAEMON_JSON}"
fi

mkdir -p "${DOCKER_ROOT}"
new_config=$(jq --arg root "${DOCKER_ROOT}" '. += { "data-root": $root }' "${DAEMON_JSON}")

echo "INFO: Updated docker daemon config:"
echo "${new_config}"
echo "${new_config}" > "${DAEMON_JSON}"

echo "INFO: Restarting docker"
systemctl restart docker

docker info
