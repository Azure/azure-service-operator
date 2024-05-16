---
title: Running a Development Version
---

## Locally

If you would like to try something out but do not want to write an integration test, you can run your local version of the
operator locally in a [kind](https://kind.sigs.k8s.io) cluster.

Before launching `kind`, make sure that your shell has the `AZURE_SUBSCRIPTION_ID` and `AZURE_TENANT_ID`
environment variables set. See [testing](../testing/#recordreplay) for more details about them.

Once you've set the environment variables above, create a `kind` cluster:
- `task controller:kind-create-with-workload-identity`.

You can use `kubectl` to interact with the local `kind` cluster.

When you're done with the local cluster, tear it down with `task controller:kind-delete`.

## On AKS

Sometimes running in `kind` does not suffice and a real cluster is needed. The `task controller:aks-create-helm-install`
will perform the following actions:
- Create an AKS cluster named `{{.HOSTNAME}}-aso-aks` in a resource group `{{.HOSTNAME}}-aso-rg`. These resources
  are created in the subscription set in the `AZURE_SUBSCRIPTION_ID` environment variable.
  - By default, the cluster is created in `westus3`, but that can be overridden by specifying the `LOCATION` variable to
    the `task` command like so: `task controller:aks-create-helm-install LOCATION=mylocation`
- Create an ACR in the `{{.HOSTNAME}}-aso-rg` associated with the AKS cluster.
- Install `cert-manager` into the cluster (required for ASO).
- Build and push your local container image into the ACR.
- Install ASO into the cluster, using the ACR image as the source for the controller pod.
