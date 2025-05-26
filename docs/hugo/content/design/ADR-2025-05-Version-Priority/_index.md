---
title: '2025-05: Resources and Version Priority'
toc_hide: true
---

## Context

As initially identified in [#4147](https://github.com/Azure/azure-service-operator/issues/4147), we have a problem with the way we version resources.

If a specific resource isn't specified, Kubernetes will automatically select one for use through a process called [_version priority_](https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definition-versioning/#version-priority). 

Unfortunately, the way we are currently constructing our versions isn't playing well with this algorithm.

When ASO constructs a resource version number, it uses a constant prefix followed by the Azure api-version of the resource. 

Beta releases of ASO used the constant prefix `v1beta` giving rise to resource versions such as:

* `v1beta20210201`
* `v1beta20210201preview`

These are compliant with how Kubernetes defines versions, and the system would correctly select the most recent version as the default version.

When we went GA, we changed the prefix to `v1api` (removing the `beta` tag), giving resource versions such as:

* `v1api20210201`
* `v1api20240201` 
* `v1api20210201preview`

While `beta` is in the list of identifiers permitted in Kubernetes versions, `api` is not, so we've inadvertently ended up with version numbers that are not compliant with Kubernetes rules.

The Kubernetes _version priority_ algorithm gives priority to compliant versions, and sorts the remaining versions alphabetically. 

With all `v1api` versions considered non-compliant, they're sorted alphabetically and then the first one selected - resulting in automatic selection of the *oldest version*, not the newest.

## Requirements

Including the Azure API Version as a component of an ASO Resource Version gives our users the ability to select which Azure API version they use. This can be critical when behaviour changes between versions. It also helps to avoid any nasty surprises, as might happen when a new version gets adopted without warning.

We originally chose the `v1` prefix to give us the ability to increment the version if we needed to make a breaking change to our generated resources; we haven't needed to do this, and we can potentially drop this requirement.

## Option 1: Do nothing

Leave things the way they are.

### Pros

* Simple to achieve

### Cons

* We've already had a couple cases where version priority caused issues for users, so we know this problem isn't going away.
* As we add more resource versions, the scope for this occurring can only increase.

## Option 2: Cut over to versioning format

Change all generated API versions to be compliant with Kubernetes version rules.
### Pros

* Version priority will work as users expect.

### Cons

* Would be breaking for all existing users.

## Option 3: Migrate to new versioning format over time

Introduce a new version format as from a particular release of ASO, leaving existing resources with the existing format. 

### Pros

* Not a breaking change for our users.
* As we introduce new resource versions, the number of resources where version priority does something untoward will decrease.

### Cons

* Still allows the problem to occur (but prevents it from getting worse)

## Version format

A simple change to the format we use for resource versions would be to simplify the prefix used to just `v`, giving resource versions like this:

* `v20210201`
* `v20240201` 
* `v20210201preview`

Unfortunately, multi-part versions aren't supported for Kubernetes custom resources, so we *can't* lean into the date-based scheme used for Azure api-versions and use `.` separators as follows:

* `v2021.02.01`
* `v2024.02.01` 
* `v2021.02.01.preview`

# Status

Recommendation: Option 3 with the proposed version change.

