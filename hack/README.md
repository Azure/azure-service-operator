# Azure Service Operator

## Developer setup (with VS Code)
This repository contains a [devcontainer](https://code.visualstudio.com/docs/remote/containers) configuration that can be used in conjunction with VS Code to set up an environment with all the required tools preinstalled. 

If you want to use this:

0. Make sure you have installed [the prerequisites to use Docker](https://code.visualstudio.com/docs/remote/containers#_system-requirements), including WSL if on Windows.
1. Install VS Code and the [Remote Development](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack) extension (check installation instructions there).
2. Run the VS Code command (with `Ctrl-Shift-P`): `Remote Containers: Clone Repository in Container Volume...`
  
   **Note**: in Windows, it is important to clone directly into a container instead of cloning first and then loading that with the `Remote Containers` extension, as the tooling performs a lot of file I/O and if this is performed against a volume mounted in WSL, then it is unusably slow.

   To complete the clone:
   1. Select "`GitHub`".
   2. Search for "`Azure/k8s-infra`".
   3. Choose either of the following options about where to create the volume.
   4. The window will reload and run the `Dockerfile` setup. The first time, this will take some minutes to complete as it installs all dependencies.

3. To validate everything is working correctly, you can open a terminal in VS Code and run `task -l`. This will show a list of all `task` commands. Running `task` by itself (or `task default`) will run quick local pre-checkin tests and validation.

### Without VS Code

The same `Dockerfile` can also be used if you are not using VS Code; it is stored in the root `.devcontainer` directory and can be used to create a development container with all the tooling preinstalled.

## Running integration tests

Run `task generated:test-integration-envtest`.

TODO: instructions on record/replay usage, etc.
