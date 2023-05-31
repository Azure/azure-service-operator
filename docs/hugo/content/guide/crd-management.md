---
title: CRD management in ASO
weight: -3
---

## How does CRD management work in Azure Service Operator?

ASO performs CRD management differently than many traditional operators. You might notice if you install the
[Helm chart]( {{< relref "installing-from-helm" >}} ) or the [YAML]( {{< relref "installing-from-yaml" >}} ) that there are NO CRDs installed by default.

This is for two reasons:

1. ASO supports so many CRDs that it exceeds the maximum Helm chart size when CRDs are included, so the CRDs cannot be 
   included in the Helm chart.
2. ASO supports so many CRDs that it can cause problems if they're all installed onto a Kubernetes cluster with a low
   memory limit. See [#2920](https://github.com/Azure/azure-service-operator/issues/2920) for details.

For the above reasons, by default each version of the ASO operator pod embeds the CRDs inside the pod and installs them when
the pod launches. 

**No CRDs are installed by default.** New CRDs will only be installed if they are specified in the 
`--crd-pattern` cmdline argument to the `manager` container. `--crd-pattern` is a semicolon delimited string of CRD
patterns to apply. Each pattern is a case-insensitive glob (supports `*`, `?` globbing).

**Existing CRDs are always upgraded.** If a CRD already exists in the cluster (from a previous version of ASO or 
a manual CRD installation), it will always be updated to the latest version of the CRD. Note that upgrading a CRD doesn't
upgrade existing resources, it just gives access to new versions of that CRD. For example `resources.azure.com/ResourceGroup`
might be upgraded to support a new API version `v1api20230601`.

## Helm

### Automatic CRD installation (recommended)

Use the `crdPattern` value in `values.yaml` to specify what new CRDs to install. Existing CRDs will always be upgraded
even if `crdPattern` is empty.

Example: `resources.azure.com/*;containerservice.azure.com/*;keyvault.azure.com/*;managedidentity.azure.com/*;eventhub.azure.com/*`

> **Note:** If you are specifying the Helm values from a shell with `--set crdPattern=...` you must wrap the crdPattern
> argument in single quotes `'` if it contains wildcard characters the shell would expand. For example:
> `--set crdPattern='resources.azure.com/*;containerservice.azure.com/*'`

> **⚠️ Warning:** A `crdPattern` of `*` can be used to install all resources, but is not recommended.
> See [best practices](#best-practices) for more details.

### Manual CRD installation

If you do not want to grant the operator permission to create + update CRDs, you can instead manage the CRDs yourself.
To do this you can use a YAML file containing all the ASO custom resource definitions that we include with each release. 
This file contains _only_ the ASO CRDs.
You can then filter to whichever subset of CRDs you want or apply them all. 
See for example [azureserviceoperator_customresourcedefinitions_v2.0.0.yaml](https://github.com/Azure/azure-service-operator/releases/download/v2.0.0/azureserviceoperator_customresourcedefinitions_v2.0.0.yaml)

The process to do this for a fresh install is:

1. Install the CRDs you want to use. These _must_ be the CRDs for operator version you install in the next step.
2. Install the operator Helm chart. Specify `installCRDs: false` as part of `values.yaml` when installing.
3. Ensure the pod launches correctly.

The process for upgrading the operator is similar:

1. Upgrade the CRDs you are using. These _must_ be the CRDs for operator version you are upgrading to.
2. Upgrade the operator Helm chart. Specify `installCRDs: false` as part of `values.yaml` when upgrading.
3. Ensure the pod launches correctly.

## YAML

### Automatic CRD installation (recommended)

Unlike Helm there are no parameters built into the raw YAML. You must use a tool such as `kustomize` or `sed` to
set the `--crd-pattern` command-line argument of the operator pod to include the CRDs you wish to install.

Example: `cat operator.yaml | sed 's_--crd-pattern=.*_--crd-pattern=resources.azure.com/*;containerservice.azure.com/*_g' | kubectl apply -`

Alternatively you can use `kustomize` with a configuration like:
```yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

resources:
  - https://github.com/Azure/azure-service-operator/releases/download/v2.1.0/azureserviceoperator_v2.1.0.yaml

patches:
  - patch: |-
      - op: add
        path: /spec/template/spec/containers/0/args/-
        value: --crd-pattern=resources.azure.com/*;containerservice.azure.com/*;keyvault.azure.com/*;managedidentity.azure.com/*;eventhub.azure.com/*
    target:
      kind: Deployment
```

> **⚠️ Warning:** A `--crd-pattern` of `*` can be used to install all resources, but is not recommended.
> See [best practices](#best-practices) for more details.

### Manual CRD installation

If you do not want to grant the operator permission to create + update CRDs, you can instead manage the CRDs yourself.
To do this you can use a YAML file containing all the ASO custom resource definitions that we include with each release.
This file contains _only_ the ASO CRDs.
You can then filter to whichever subset of CRDs you want or apply them all.
See for example [azureserviceoperator_customresourcedefinitions_v2.0.0.yaml](https://github.com/Azure/azure-service-operator/releases/download/v2.0.0/azureserviceoperator_customresourcedefinitions_v2.0.0.yaml)

The process to do this for a fresh install is:

1. Install the CRDs you want to use. These _must_ be the CRDs for operator version you install in the next step.
2. Install the operator YAML.
3. Ensure the pod launches correctly.

The process for upgrading the operator is similar:

1. Upgrade the CRDs you are using. These _must_ be the CRDs for operator version you are upgrading to.
2. Upgrade the operator YAML.
3. Ensure the pod launches correctly.

## Uninstalling CRDs

ASO CRD automation will **never** under any circumstances uninstall CRDs. In fact, it doesn't have delete CRD permissions.
This is a safety feature because removal of CRDs can be dangerous.

The recommended workflow for removing CRDs you are no longer using is:

1. Delete or otherwise remove all instances of the CRD in question. For example if you want to remove the 
   `keyvault.azure.com/vault` CRD, first delete all instances of that CRD.
   * If you want to retain the `vault` resources in Azure, apply a [reconcile-policy]( {{< relref "annotations#serviceoperatorazurecomreconcile-policy" >}} )
     of `skip` prior to issuing `kubectl delete`.
2. Ensure that the operator pod does not have a pattern matching the CRD you would like to remove in `--crd-patterns`.
   If it does, the CRD may be re-installed by the operator pod after you've removed it.
3. Stop the operator pod.
4. Delete the CRD using `kubectl delete`.
5. Start the operator pod.

## Best Practices

* Do not use `crdPattern=*` on AKS free tier clusters. It installs so many CRDs that it can overload kube-apiserver and cause
  restarts. See [#2920](https://github.com/Azure/azure-service-operator/issues/2920) for details. Instead, specify
  a `crdPattern` containing only the groups (or individual CRDs) you intend to use.
* We strongly recommend including entire groups such as `dbformysql.azure.com/*` in `crdPattern`. Individual CRDs such as
  `dbformysql.azure.com/FlexibleServer` may be specified as well, but there are often other resources in the 
  group which pair together to enable other scenarios, such as `dbformysql.azure.com/FlexibleServersFirewallRules`.
  It's generally easier to just include the whole group.
