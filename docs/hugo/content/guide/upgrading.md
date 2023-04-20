---
title: "Upgrading"
weight: -3
---

## Before you upgrade

Ensure that you have carefully reviewed the upgrade instructions included in the [release notes](https://github.com/Azure/azure-service-operator/releases) for the release you are upgrading to. If you are upgrading multiple versions, check the release notes for each version. 

### Caution

<!-- Our replacementPatterns don't handle `.` in the filename well.
     Manually checked by @theunrepentantgeek -->
<!-- markdown-link-check-disable-next-line -->
If upgrading to **v2.0.0**, carefully review [v2.0.0 Breaking Changes](../breaking-changes/breaking-changes-v2.0.0), especially the section on using [`asoctl clean crds`](./../../tools/asoctl#clean-crds).

<!-- Our replacementPatterns don't handle `.` in the filename well.
     Manually checked by @theunrepentantgeek -->
<!-- markdown-link-check-disable-next-line -->
If upgrading to **v2.0.0-beta.4**, carefully review [v2.0.0-beta.4 Breaking Changes](../breaking-changes/breaking-changes-v2.0.0-beta.4), you may need to make some minor changes to your resources.

## Recommended upgrade pattern

We recommend that you upgrade ASO one minor version at a time, and that you plan to upgrade to the latest version of ASO so that you get benefit from the latest bug fixes and features.

| Old version      | New version     | Recommended |
| ---------------- | --------------- | :---------: |
| `v2.0.0-beta.0`  | `v2.0.0-beta.1` |      ✔      |
| `v2.0.0-alpha.6` | `v2.0.0-beta.1` |      ❌      |

If you need to upgrade Azure Service Operator more than 1 minor version, we recommend that you do so one minor version at a time, giving your cluster time to stabilize between upgrades.

For example `v2.0.0-alpha.6` to `v2.0.0-beta.1` should be accomplished with two upgrades:
 - `v2.0.0-alpha.6` to `v2.0.0-beta.0`
 - `v2.0.0-beta.0` to `v2.0.0-beta.1`

## Upgrading

We recommend that you upgrade ASO using the same tool you installed it with. 

**Note:** The instructions below all assume you're upgrading from the previous version (N-1) to the latest version (vN). Please follow 
the [Recommended upgrade pattern](#recommended-upgrade-pattern) for upgrading multiple versions. 

### kubectl apply

The operator can be upgraded simply by running the same command you used to install it: 

```bash
kubectl apply --server-side=true -f https://github.com/Azure/azure-service-operator/releases/download/v2.0.0/azureserviceoperator_v2.0.0.yaml
```

### Helm

```bash
helm repo add aso2 https://raw.githubusercontent.com/Azure/azure-service-operator/main/v2/charts
helm repo update
helm upgrade --devel --version v2.0.0 aso2 aso2/azure-service-operator \ 
        --namespace=azureserviceoperator-system \
        --set azureSubscriptionID=$AZURE_SUBSCRIPTION_ID \
        --set azureTenantID=$AZURE_TENANT_ID \
        --set azureClientID=$AZURE_CLIENT_ID \
        --set azureClientSecret=$AZURE_CLIENT_SECRET
```

## Supported Versions

We provide best effort support via GitHub for ASO, focusing on the latest version and the previous version. If you're running an older version, we may ask you to upgrade to the latest release before we can help you.

## Downgrading

If you encounter an issue after an upgrade of ASO that can't be addressed by any other means, you may want to downgrade to the prior version. This isn't an officially supported scenario, we recommend rolling forward to a newer version if at all possible, but we'll try to provide some guidance here.

Summary: If the new version of ASO didn't change the stored version of your resource, then you should be able to downgrade without issue. If the new version of ASO did change the stored version of your resource, then you will need to manually fix things up.

To determine whether the new version of ASO changed the stored version of your resource, consult both the [release notes](https://github.com/Azure/azure-service-operator/releases) for that version and the list of [supported resources](../../reference/).

Each ASO resource is stored within the cluster using a canonical storage (or hub) version, regardless of the actual version used for interaction with the cluster or with Azure. These stored versions are based on the latest Stable Azure API version of that resource. (See [Resource Versioning](../../design/versioning/) for all the background on this design decision.)

When a new version of ASO adds support for a later version of an already supported resource (say, adding support for the latest release of Azure Redis, `v1api20220601`, in addition to the existing support for `v1api20211201`), then the new version of ASO will use the newer API version of the resource as the storage version. In this situation, downgrading ASO to the prior version will orphan the resource, as the older version of ASO is unaware of the newer storage version and is unable to interact with it. This means that any resources created with, or modified by, the newer version of ASO will be unable to be managed by the older version of ASO. They will also be inaccessible via `kubectl` or any other tool.

These orphaned resources can be recovered by reapplying the resource (using `kubectl apply`) and overwriting the resource with a new revision after the downgrade. Alternatively, if you are going to be re-upgrading to ASO soon, you can leave those resources in place. The older version of ASO will ignore them, and any newer version should pick them up and start reconciliation automatically.
These orphaned resources can be recovered by reapplying the resource (using `kubectl apply`) and overwriting the resource with a new version after the downgrade. Alternatively, if you are going to be re-upgrading to ASO soon, you can leave those resources in place. The older version of ASO will ignore them, and any newer version should pick them up and start reconciliation automatically.

When a new version of ASO does not change the stored version of a resource (say, adding support for an old version of Azure Redis, `v1api20201201`, in addition fo the existing support for `v1api20211201`) then the new version of ASO will be using the same API version of the resource as the prior release. In this situation, downgrading ASO to the prior version should work without issue.

Given that it's unusual for a release of ASO to upgrade all resources, it's likely that you'll find both of the above scenarios in your cluster, depending on the variety of ASO supported resource types in play.

### Preparation for downgrading

For each ASO managed resource in your cluster

* Check the [ASO release notes](https://github.com/Azure/azure-service-operator/releases) and the list of [supported resources](../../reference/) to identify whether the ASO upgrade added a new version for that resource type.
* Ensure you have the required YAML definition to recreate the resource in your cluster. If necessary, download the YAML definition for the resource using `kubectl get`.

