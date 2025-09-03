# Library Scripts


## Docker-in-Docker

The script `docker-in-docker.sh` is [copied from GitHub](https://github.com/devcontainers/features/tree/main/src/docker-in-docker) and used here to preinstall things on the container image.

To update it to the latest version, run the following command from the `.devcontainer/library-scripts` folder:

```bash
curl --output docker-in-docker.sh https://raw.githubusercontent.com/devcontainers/features/refs/heads/main/src/docker-in-docker/install.sh
```

