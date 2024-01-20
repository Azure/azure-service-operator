---
title: Developer Setup
linktitle: Developer Setup
weight: -10
cascade:
- type: docs
- render: always
description: "How to set up your developer environment for Azure Service Operator v2"
---

We support a number of different approaches to ASO development.

- [Dev Container with VS Code on Linux](#dev-container-with-vs-code-on-linux)
- [Dev Container with VS Code on Windows](#dev-container-with-vs-code-on-windows)
- [Docker on Linux](#docker-on-linux)
- [CLI on Linux](#cli-on-linux)
- [CLI on MacOS](#cli-on-macos)

Each of these is described in a different section below. See also the [troubleshooting](#troubleshooting-repo-health) sections below for help with common problems.

## Dev Container with VS Code on Linux 

Use these steps if you've checked out the ASO code into a Linux environment (_including_ WSL 2 on Windows). We've found this to be the best performing option.

The ASO repository contains a [devcontainer](https://code.visualstudio.com/docs/remote/containers) configuration that can be used in conjunction with VS Code to set up an environment with all the required tools preinstalled.

0. Make sure you have installed [the prerequisites to use Docker](https://code.visualstudio.com/docs/remote/containers#_system-requirements), including [WSL](https://learn.microsoft.com/en-us/windows/wsl/install) if on Windows. 
1. Install VS Code and the [Remote Development](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack) extension (check installation instructions there).
2. Ensure you have ASO checked out, and that your [repo is healthy](#troubleshooting-repo-health).
3. Open VS Code by running `code .` from the root folder of the ASO repo.
4. VS Code will automatically notice the `.devcontainer` folder and ask you if you want to open it; press _Reopen in Container_ when prompted.
   * If you miss this, type `Control-Shift-P` and run the command `Dev Containers: Reopen in Container`.
5. The window will reload and run the `Dockerfile` setup. The first time, this will take some minutes to complete as it installs all dependencies. Later runs will be faster due to caching.
6. To validate everything is working correctly,  open a terminal in VS Code and run `task -l`. This will show a list of all `task` commands. Running `task` by itself (or `task default`) will run quick local pre-checkin tests and validation.

## Dev Container with VS Code on Windows

If you're working on Windows and _**not**_ using WSL 2, this is the recommended setup.

The ASO repository contains a [devcontainer](https://code.visualstudio.com/docs/remote/containers) configuration that can be used in conjunction with VS Code to set up an environment with all the required tools preinstalled.

0. Make sure you have installed [the prerequisites to use Docker](https://code.visualstudio.com/docs/remote/containers#_system-requirements), including [WSL](https://learn.microsoft.com/en-us/windows/wsl/install) if on Windows. 
1. Install VS Code and the [Remote Development](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack) extension (check installation instructions there).
2. Open VS Code
3. Run the VS Code command (with `Ctrl-Shift-P`): `Dev Containers: Clone Repository in Container Volume...`

   **Note**: in Windows, it is important to clone directly into a container instead of cloning first and then loading that with the `Remote Containers` extension, as the tooling performs a lot of file I/O, and if this is performed against a Windows folder mounted into the devcontainer then it is unusably slow.

4. To complete the clone:
    1. Select "`GitHub`".
    2. Search the azure-service-operator repo.
       * If you have direct access to the ASO repo, search for `Azure/azure-service-operator`
       * Otherwise search for `<your-user-name>/azure-service-operator`
    3. If asked, choose either of the following options about where to create the volume.
    4. The window will reload and run the `Dockerfile` setup to build the devcontainer. The first time, this will take some minutes to complete as it installs all dependencies. Later runs will be faster due to caching.
5. Once the repository is cloned in your local
   * Make sure you have access to tags (by running `git tag --list 'v2*'`).
   * Set up our submodule by running both `git submodule init` and `git submodule update`.
   * See [Check your repo health](#troubleshooting-repo-health) for more information.
6. To validate everything is working correctly,  open a terminal in VS Code and run `task -l`. This will show a list of all `task` commands. Running `task` by itself (or `task default`) will run quick local pre-checkin tests and validation.

## Docker on Linux

The same `Dockerfile` that the VS Code `devcontainer` extension uses can also be used outside of VS Code; it is stored in the root `.devcontainer` directory and can be used to create a development container with all the tooling preinstalled:

```console
$ docker build $(git rev-parse --show-toplevel)/.devcontainer -t asodev:latest
… image will be created …

$ # After that you can start a terminal in the development container with:
$ docker run --env-file ~/work/envs.env --env HOSTROOT=$(git rev-parse --show-toplevel) -v $(git rev-parse --show-toplevel):/go/src -w /go/src -u $(id -u ${USER}):$(id -g ${USER}) --group-add $(stat -c '%g' /var/run/docker.sock) -v /var/run/docker.sock:/var/run/docker.sock --network=host  -it asodev:latest /bin/bash
```

Note: If you mount the source like this from a Windows folder, performance will be poor as file operations between the container and Windows are very slow.

## CLI on Linux

If you are using Linux, instead of using VS Code you can run the `dev.sh` script in the root of the repository. This will install all required tooling into the `hack/tools` directory and then start a new shell with the `PATH` updated to use it.

## CLI on MacOS

Development of ASO on MacOS is also possible.

You can either use the VS Code devcontainer approach (recommended) which installs all the tools into a container, or you can install the tools directly on your Mac. In case of the latter, you'll need to install the following tools manually running: `.devcontainer/install-dependencies.sh`.

This creates `hack/tools` and downloads all the required tools into it based on the architecture(arm64 or amd64) of your machine.

If you have an ARM based Mac, you'll also need to install [Rosetta](https://support.apple.com/en-nz/HT211861).

## Troubleshooting: Repo health

A simple cloning of the ASO repo is not enough to successfully run a build. You must also ensure:

* You have access to git tags.  
  The build scripts depend on tags in order to create a version number.

* The `azure-rest-api-specs` submodule has been cloned.  
  The specifications in this repo are the input used by the code generator.

### Git tag access 

Some tools default to shallow cloning of repos, omitting tags. The build process depends on tags being present and will fail if they're missing. 

To check if you have any tags:

``` bash
$ git tag --list 'v2*'
v2.0.0
# ... elided ...
v2.1.0
```

If you see a list of tags (as shown above), then you're good to go.

Otherwise, pull tags from your upstream repo and check again:

``` bash
$ git fetch --all --tags
Fetching origin
$ git tag --list 'v2*'
v2.0.0
# ... elided ...
v2.1.0
```

If the list is still empty, add the ASO repo as a direct upstream reference and pull the tags from there.

``` bash
$ git remote add azure https://github.com/Azure/azure-service-operator.git
$ git fetch azure
$ git tag --list 'v2*'
v2.0.0
# ... elided ...
v2.1.0
```

You should now have all the tags.

### Submodule cloning

ASO references the `azure-rest-api-specs` repo containing API definitions for Azure Resource Providers. If this submodule is missing, the code generator will not run.

From the root of your ASO repo clone, run

``` bash
$ git submodule init
Submodule 'azure-rest-api-specs' (https://github.com/Azure/azure-rest-api-specs) registered for path 'v2/specs/azure-rest-api-specs'

$ git submodule update
Cloning into '/workspaces/azure-service-operator/v2/specs/azure-rest-api-specs'...
remote: Enumerating objects: 15, done.
... elided ...
Submodule path 'v2/specs/azure-rest-api-specs': checked out '35f8372749cbe9b1f69bde0eee3b0bc63c49c2c2'
```

The exact git hash that's checked out will likely be different.

## Troubleshooting: Docker on Windows

We've observed that Docker on Windows performs significantly better when using WSL as the back end. Check your settings in Docker Desktop and ensure the option _Use the WSL 2 based engine_ is turned on.

![Docker Desktop General Settings showing WSL 2 option](../images/docker-desktop-wsl.png)

## Troubleshooting: Visual Studio Code

If the commands to open a dev container don't work, check to see where you have Visual Studio Code installed. Open a PowerShell window and run this command:

``` powershell
PS> get-command code
CommandType     Name        Version    Source
-----------     ----        -------    ------
Application     code.cmd    0.0.0.0    C:\Users\<user>\AppData\Local\Programs\Microso...
```

If your installation is a global one (in `C:\Program Files` or `C:\Program Files (x86)`) then you may find some extensions don't work properly. Install another copy of Visual Studio Code using the standard installer; this will be a local installation into your `AppData` folder and will take priority.

