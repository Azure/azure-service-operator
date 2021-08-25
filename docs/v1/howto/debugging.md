# Debugging

The following guide details how to debug the Azure Service Operator locally using the power of DevContainers and Kind.

Running the debugger in VSCode from within a DevContainer, you'll be able to interact with the operator just as you would if it was running within kubernetes using `kubectl`.

![debugging in vscode](images/debugging.gif)

## Prerequisites

* DevContainer is up and running.
* The operator is deployed and running within a local k8s cluster.

## Step-by-step guide

Before we start, verify your local k8s cluster is running and the operator is deployed using `kubectl get pods -n azureoperator-system`.

If it is not, run `make set-kindcluster` to spin up a new local k8s cluster using Kind with the operator deployed.

For this example we'll be working with:

* [controllers/async_controller.go](../controllers/async_controller.go)
* [config/samples/azure_v1alpha1_resourcegroup.yaml](../config/samples/azure_v1alpha1_resourcegroup.yaml)

If you're not familiar with how ResourceGroup works, spend some time reviewing `resourcegroup_controller.go`, more specifically the `Reconcile` func which can be found on [this](../controllers/resourcegroup_controller.go#L41) line.

1. Set your breakpoints. Place our breakpoint anywhere within the `Reconcile` func.
2. Create a folder in the root of the project called `.vscode` and create a new file in that folder called `launch.json`.
3. Copy and paste the following inside `launch.json`.

    ```json
    {
        "version": "0.2.0",
        "configurations": [
            {
                "name": "Debug",
                "type": "go",
                "request": "launch",
                "mode": "auto",
                "program": "${workspaceFolder}/main.go",
                "env": {},
                "args": []
            }
        ]
    }
    ```

4. From your menu bar, click `Debug`-> `Start Debugging` (or simply hit `F5`).
5. From your console panel, click the `DEBUG CONSOLE` tab and verify the debugger is running. You should see something like this: INSERT IMAGE
6. Now click on the `TERMINAL` tab and enter `kubectl apply -f config/samples/azure_v1alpha1_resourcegroup.yaml`.

If you've done everything right, you should see your breakpoint hit.

Happy debugging!
