---
title: "Upgrading"
---

## Before you upgrade

Ensure that you have carefully reviewed the upgrade instructions included in the [release notes](https://github.com/Azure/azure-service-operator/releases) for the release
you are upgrading to.

## Recommended upgrade pattern

We recommend that you upgrade ASO one minor version at a time. 

| Old version      | New version     |  Recommended  |
|------------------|-----------------|:-------------:|
| `v2.0.0-beta.0`  | `v2.0.0-beta.1` |       ✔       |
| `v2.0.0-alpha.6` | `v2.0.0-beta.1` |       ❌       |

If you need to upgrade Azure Service Operator more than 1 minor version, we recommend that you do so one minor version at a time.
For example `v2.0.0-alpha.6` to `v2.0.0-beta.1` should be accomplished with two upgrades:
 - `v2.0.0-alpha.6` to `v2.0.0-beta.0`
 - `v2.0.0-beta.0` to `v2.0.0-beta.1`

## Upgrading

We recommend that you upgrade ASO using the same tool you installed it with. 

**Note:** The instructions below all assume you're upgrading from the previous version to the latest version. Please follow 
the [Recommended upgrade pattern](#recommended-upgrade-pattern) for upgrading multiple versions. 

### kubectl apply

The operator can be upgraded simply by running the same command you used to install it: 

```bash
kubectl apply --server-side=true -f https://github.com/Azure/azure-service-operator/releases/download/v2.0.0-beta.0/azureserviceoperator_v2.0.0-beta.1.yaml
```

### Helm

```bash
helm repo add aso2 https://raw.githubusercontent.com/Azure/azure-service-operator/main/v2/charts
helm repo update
helm upgrade --devel --version v2.0.0-beta.0 aso2 aso2/azure-service-operator \ 
        --namespace=azureserviceoperator-system \
        --set azureSubscriptionID=$AZURE_SUBSCRIPTION_ID \
        --set azureTenantID=$AZURE_TENANT_ID \
        --set azureClientID=$AZURE_CLIENT_ID \
        --set azureClientSecret=$AZURE_CLIENT_SECRET
```