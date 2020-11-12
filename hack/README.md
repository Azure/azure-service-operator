# Azure Service Operator

## Developer setup (with VS Code)
This repository contains a [devcontainer](https://code.visualstudio.com/docs/remote/containers) configuration that can be used in conjunction with VS Code to set up an environment with all the required tools preinstalled. 

This is the recommended setup, especially if you are using Windows as your development platform.

If you want to use this:

0. Make sure you have installed [the prerequisites to use Docker](https://code.visualstudio.com/docs/remote/containers#_system-requirements), including WSL if on Windows.
1. Install VS Code and the [Remote Development](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack) extension (check installation instructions there).
2. Run the VS Code command (with `Ctrl-Shift-P`): `Remote Containers: Clone Repository in Container Volume...`
  
   **Note**: in Windows, it is important to clone directly into a container instead of cloning first and then loading that with the `Remote Containers` extension, as the tooling performs a lot of file I/O, and if this is performed against a volume mounted in WSL then it is unusably slow.

   To complete the clone:
   1. Select "`GitHub`".
   2. Search for "`Azure/k8s-infra`".
   3. Choose either of the following options about where to create the volume.
   4. The window will reload and run the `Dockerfile` setup. The first time, this will take some minutes to complete as it installs all dependencies.

3. To validate everything is working correctly, you can open a terminal in VS Code and run `task -l`. This will show a list of all `task` commands. Running `task` by itself (or `task default`) will run quick local pre-checkin tests and validation.

## Without VS Code

### Dockerfile

The same `Dockerfile` that the VS Code `devcontainer` extension uses can also be used outside of VS Code; it is stored in the root `.devcontainer` directory and can be used to create a development container with all the tooling preinstalled:

```console
$ docker build $(git rev-parse --show-toplevel)/.devcontainer -t k8sinfradev:latest
… image will be created …

$ # After that you can start a terminal in the development container with:
$ docker run -v $(git rev-parse --show-toplevel):/go/src -w /go/src -u $(id -u ${USER}):$(id -g ${USER}) -it k8sinfradev:latest
```

It is not recommended to mount the source like this on Windows (WSL2) as the cross-VM file operations are very slow.

### ./dev.sh

If you are using Linux, instead of using VS Code you can run the `dev.sh` script in the root of the repository. This will install all required tooling into the `hack/tools` directory and then start a new shell with the `PATH` updated to use it.

## Running integration tests

Basic use: run `task controller:test-integration-envtest`.

### Record/replay

The task `controller:test-integration-envtest` runs the tests in a record/replay mode by default, so that it does not touch any live Azure resources. (This uses the [go-vcr](https://github.com/dnaeon/go-vcr) library.) If you change the controller or other code in such a way that the required requests/responses from ARM change, you will need to update the recordings. 

To do this, delete the recordings for the failing tests (under `{test-dir}/recordings/{test-name}.yml`), and re-run `controller:test-integration-envtest`. If the test passes, a new recording will be saved, which you can commit to include with your change. All authentication and subscription information is removed from the recording.

To run the test and produce a new recording you will also need to have set the required authentication environment variables for an Azure Service Principal: `AZURE_SUBSCRIPTION_ID`, `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and `AZURE_CLIENT_SECRET`. This Service Principal will need access to the subscription to create and delete resources.

If you need to create a new Azure Service Principal, run the following commands:

```console
$ az login
… follow the instructions …
$ az account set --subscription {the subscription ID you would like to use}
Creating a role assignment under the scope of "/subscriptions/{subscription ID you chose}"
…
$ az ad sp create-for-rbac --role contributor --name {the name you would like to use}
{
  "appId": "…",
  "displayName": "{name you chose}",
  "name": "{name you chose}",
  "password": "…",
  "tenant": "…"
}
```
The output contains `appId` (`AZURE_CLIENT_ID`), `password` (`AZURE_CLIENT_SECRET`), and `tenant` (`AZURE_TENANT_ID`). Store these somewhere safe as the password cannot be viewed again, only reset. The Service Principal will be created as a “contributor” to your subscription which means it can create and delete resources, so **ensure you keep the secrets secure**.

### Running live tests

If you want to skip all recordings and run all tests directly against live Azure resources, you can use the `controller:test-integration-envtest-live` task. This will also require you to set the authentication environment variables, as detailed above.
