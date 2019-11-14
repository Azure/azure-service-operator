# Debugging

The following guide details how to debug the Azure Service Operator locally using the power of DevContainers and Kind.

Running the debugger in VSCode from wihtin a DevContainer, you'll be able to interact with the operator just as you would if it was running within kubernetes using `kubectl`.

![debugging in vscode](images/debugging.gif)

## Prerequests

* DevContainer is up and running.
* The operator is deployed and running within a local k8s cluster.

## Step-by-step guide

Before we start, verify your local k8s cluster is running and the operator is deployed using `kubectl get pods -n azureoperator-system`.

If it is not, run `make set-kindcluster` to spin up a new local k8s cluster using Kind with the operator deployed.

For this example we'll be working with:

* [controllers/async_controller.go](../controllers/async_controller.go)
* [config/samples/azure_v1alpha1_resourcegroup.yaml](../config/samples/azure_v1alpha1_resourcegroup.yaml)

If you're not familar with how ResourceGroup works, spend some time reviewing `resourcegroup_controller.go`, more specificly the `Reconcile` func found on this [this](../controllers/resourcegroup_controller.go#L41) line. 

1. Set your breakpoints. Place our breakpoint anywhere within the`Reconcile` func.
2. From your menu bar, click `Debug`-> `Start Debugging` (or simply hit `F5`).
3. From your console panel, click the `DEBUG CONSOLE` tab and verify the debugger is running. You should see something like this: INSERT IMAGE
4. Now click on the `TERMINAL` tab and enter `kubectl apply -f config/samples/azure_v1alpha1_resourcegroup.yaml`.

If you've done everything right you should see your breakpoint hit.

Happy debugging!
