---
title: "v2.4.0 Breaking Changes"
linkTitle: "v2.4.0"
weight: -20  # This should be 5 lower than the previous breaking change document
---

## Beta CRD versions have been removed

**You cannot successfully upgrade to v2.4.0 until you have followed the migration steps below:**

1. If you have never installed a `beta` version of ASOv2, or used an `beta` version of a CRD, **no action is required**.
   If you're not sure, run the below steps anyway. They will not do anything if beta was never used.
2. Ensure the cluster is running the `v2.3.0` version of ASO.
3. Install the `v2.4.0` version of the `asoctl` tool by following the [installation instructions]( {{< relref "asoctl#installation" >}})
4. Run `asoctl clean crds`. This will migrate your resources to a newer version if needed and remove the `beta` CRD
   versions from the CRD `Status.StoredVersions` collection. For more details on this command see the
   [asoctl clean crds documentation]( {{< relref "asoctl#clean-crds" >}} )

Once you have successfully executed the above steps, you can upgrade to `v2.4.0`. 

When upgrading to `v2.4.0` the following error in the operator pod logs means that you need to run the above steps:
```
E1113 22:24:48.837454       1 setup.go:193] "msg"="failed to apply CRDs" "error"="failed to apply CRD flexibleserversdatabases.dbforpostgresql.azure.com: CustomResourceDefinition.apiextensions.k8s.io \"flexibleserversdatabases.dbforpostgresql.azure.com\" is invalid: status.storedVersions[0]: Invalid value: \"v1beta20210601storage\": must appear in spec.versions" "logger"="setup" 
```
