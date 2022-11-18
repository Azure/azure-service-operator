---
title: '2022-11: Resource Import and ASO v1 Migration'
---

## Context

We have a significant number of customers successfully using ASO v1 whom we want to migrate to ASO v2. This requires both feature parity (so that users don't lose functionality when they switch) and a straightforward migration path (if it's too hard to migrate, they won't).

Discussions so far have identified two possible solutions - *version integration* and *resource import*.

### Version Integration

We would add new hand-written resources into ASO v2 with the same names and shapes as the resources supported by ASO v1. Users would be able to migrate to ASO v2 simply by removing ASO v1 from their cluster and installing ASO v2 in its place.

Since many of these resources are already supported by ASO v2, albeit with shapes much more closely aligned with the ARM originals, we would need to implement forward and backward conversions between the two shapes, preserving the Kubernetes version compatibility guarantees we have today.

* Pro: Very simple migration for our customers.
* Con: Would require a lot of work to implement the conversions.
* Con: The investment would benefit only a small number of customers.
* Con: The conversions would need to be maintained, potentially indefinitely.
* Con: It's unlikely we'd achieve 100% fidelity, requiring customers to change their YAML files anyway.

### Resource Migration

Instead of trying to bring forward the ASO v1 YAML resource definitions, we'd create a tool that allowed any supported Azure resource to be to be exported as a YAML file in the shape expected by ASO v2. Users would then be able to modify the YAML file as necessary and apply it to their cluster.

The most likely form of this tool would be as a commandline utility for users to run locally.

For example, to export the YAML for an existing Virtual Network, the user would run:

``` bash
$ asoctl import resource http://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1
```

This would output a YAML file to stdout, which the user could then redirect to a file.

Extensions might allow the user to specify an entire resource group:

``` bash
$ asoctl import resource-group http://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1
```

The output would be a single YAML file containing all the resources in the resource group.

* Pro: Straightforward migration for ASO v1 customers (though, not zero touch).
* Pro: Would benefit all customers wanting to migrate existing Azure resources, not just those using ASO v1.
* Con: We'd need to extend the code generator to provide the required support (but much of what we need already exists).

## Decision

Create a new command line tool, `asoctl` to house multiple functions, starting with resource import.

### Data Flow

The process of generating the YAML for a given resource can be simplified to the following flow:

![Data Flow](./images/2022-11-resource-import-data-flow.png)

1. The user specified URL is used to GET the resource from Azure in its native ARM format.
2. We convert this to a status object for the relevant ASO custom resource.
3. We then convert this to a spec object ready for export (this has the shape we want).
4. The spec is exported as YAML.

Only the conversion from status to spec is missing from our existing code generation pipeline, and we believe the existing property assignment code can be largely reused.

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
