---
title: '2022-11: ASO v1 Migration and Resource Import'
---

## Context

We have a significant number of customers successfully using ASO v1 whom we want to migrate to ASO v2. This requires both feature parity (so that users don't lose functionality when they switch) and a straightforward migration path (if it's too hard to migrate, they won't).

Discussions so far have identified two possible solutions - *version integration* and *resource import*.
### Version Integration

We would add new hand-written resources into ASO v2 with the same names and shapes as the resources supported by ASO v1. Users would be able to migrate to ASO v2 simply by removing ASO v1 from their cluster and installing ASO v2 in its place.

Since many of these resources are already supported by ASO v2, albeit with shapes much more closely aligned with the ARM originals, we would need to implement forward and backward conversions between the two shapes, preserving the Kubernetes version compatibility guarantees we have today.

* Pro: Very simple migration for our customers.
* Con: Would require a lot of work to implement and test the conversions.
* Con: Any resource features that don't map exactly to a property in the ARM shape would be difficult to retain.
* Con: The investment would benefit only a small number of customers.
* Con: The conversions would need to be maintained, potentially indefinitely.
* Con: It's unlikely we'd achieve 100% fidelity, requiring customers to change their YAML files anyway.

### Resource Migration

Create a tool that allowed any supported Azure resource to be to be exported as a YAML file in the shape expected by ASO v2. Users would then be able to modify the YAML file as necessary and apply it to their cluster.

The most likely form of this tool would be as a command-line utility for users to run locally. Idiomatically for the Kubernetes ecosystem, we'd call this `asoctl`.

* Pro: Straightforward migration for ASO v1 customers (though, not zero touch).
* Pro: Would benefit all customers wanting to migrate existing Azure resources, not just those using ASO v1.
* Pro: Could also be used to snapshot existing resources (e.g. to capture a hand configured resource as YAML for reuse).
* Con: We'd need to extend the code generator to provide the required support (but much of what we need already exists).

## Decision

Create a new command line tool, `asoctl` to house multiple functions, starting with resource import. The tool will follow the usual convention of supporting a variety of verbs, each activating a different mode.

To export a single existing resource, we'd use `export resource` as the verb. This would take the resource ID as a parameter and output the YAML file to stdout. The user would then be able to redirect the output to a file and modify it as necessary.

For example, to export the YAML for an existing Virtual Network, the user would run:

``` bash
$ asoctl export resource http://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1
```

As an extension, we could also support `export resourcegroup` to generate a single YAML file containing all the resources in a resource group. This would be useful for users who want to snapshot an existing resource group as YAML for reuse.

``` bash
$ asoctl export resourcegroup http://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1
```
### Data Flow

The process of generating the YAML for a given resource can be simplified to the following flow:

![Data Flow](./images/adr-2022-11-import-flow.png)

1. The user specified URL is used to ***GET*** the resource from Azure in its native ARM format.
2. We applly a ***conversion*** this to a Status object for the relevant ASO custom resource.
3. We then ***update*** a blank Spec to give us a custom resource ready for export (this has the shape we want).
4. The spec is ***exported*** as YAML.

Only the conversion from Status to Spec is missing from our existing code generation pipeline, and we believe the existing property assignment code can be largely reused.

That said, we're aware of some challenges. 

**Defaulting**: Some services will fill in missing values, so the YAML we emit will be more explicit than a handwritten resource. Sometimes defaults are documented in Swagger but more often they're not, and sometimes they're dynamic (like in AKS)

**Ordering**: When a resource contains a list of items, we may in some cases need to do an ordinal comparision (based on index) and in other cases a nominal comparison (based on name). We may be to be able to detect which is required, but it's likely we'll need configuration too.

## Status

Proposed.

## Consequences

TBC.

## Experience Report

TBC.

## References

TBC.
