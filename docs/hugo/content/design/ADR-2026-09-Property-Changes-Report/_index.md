---
title: '2026-09: Property Changes Report'
toc_hide: true
---

## Context

For full fidelity of conversions between different versions of a custom resource (CR) the Azure Service Operator code generator supports renaming between versions. This allows types that are renamed between versions to be linked, and properties that are renamed between versions to be linked.

This mapping has been constructed by hand based on a detailed review of the `structure.txt` files during new resource imports.

In [#3864](https://github.com/Azure/azure-service-operator/issues/3864) we raised the idea of a specific report to make this process easier, more straightforward, and less error-prone.

## Requirements

For each version package containing resources, we will generate a `property-changes.md` file alongside the existing `structure.txt` file. Reports are generated for both API and `storage` packages, but not for `arm`, `webhook`, `compat`, or other subpackages.

Each report compares a resource and its dependent types with the "next" version, as determined by the version-conversion graph already constructed by the code generator.

Each file includes every _resource_ defined in the package.

For the _hub_ version of each resource there is no "next" one, and no report will be generated.

Before generating reports, all previously generated `*-changes.md` files will be deleted using the same generated-file cleanup mechanism used for `structure.txt`. This prevents reports for removed resources, or packages containing only hub resources, from being left behind.

### Included Types

Each report includes:

* The resource.
* Its spec and status types.
* The recursive closure of all types referenced by the spec and status.

The closure is calculated for both versions and combined so that newly introduced and retired types are included. A type rename configured in the generator links the two type definitions; otherwise, types are matched by name.

### Report Format

The report will start with a legend explaining the statuses used. Two summary tables follow: one listing resources, and one listing the dependent object types referenced by those resources. Following will be a series of differential tables, showing the changes present in each type.

#### Summary Table format

* Both spec and status types are included in the Objects summary.
* Each package version has its own column. Subpackages use the parent version and subpackage name, such as `v1api20260801/storage`.
* If resources in the current package transition to different package versions, the table has a sparse column for each target version.
* Resources are listed alphabetically in the Resources summary.
* Referenced objects are listed alphabetically in the Objects summary.
* Type renames already specified in generator config are respected.
* Status and notes columns give a key and explanation for easy scanning.
* Sorting is alphabetical by this.Name if present, next.Name if not.

Sample summary table:

| v20240101   | v20260101    |  Status  | Notes |
| :---------- | :----------- | :------: | :---- |
| Person      | Person       | Identical |       |
| Address     | PostalAddress | Renamed |       |
|             | CensusData   |   New    |       |
| Demographics |             | Retired  |       |
| Person_Spec | Person_Spec  | Modified |       |
| Person_Status | Person_Status | Identical |    |

Status meanings

* New - introduced in the newer version.
* Renamed - name changed between versions.
* Retired - does not exist in the newer version.
* Extended - the only changes are properties added in the newer version.
* Modified - a property was retired, renamed, or had its displayed type changed.
* Identical - no properties changed.

Multiple statuses are comma-separated. For example, `Renamed, Modified` is shown if a type has both a new name and different properties.

If no other status applies, `Identical` is shown.

Configured type renames are treated as equivalent when comparing property types. A property is not marked as modified if its displayed type changed only because the referenced type was renamed.

#### Differential Table format

A differential table is generated for any type (resource or object) flagged as Modified. The table lists both types side by side with properties aligned. Each such table is preceded by a heading, and they're listed in the order of the Summary Table.

Property types use the same display format as the existing `structure.txt` reports.

For the `Person_Spec` type above, we might see:

| v20240101 |         | v20260101  |               | Status    | Notes                               |
| :-------- | :------ | :--------- | :------------ | :-----: | :---------------------------------- |
| BirthDate | date    | BirthDate  | date          | Identical |                                     |
| FirstName | string  |            |               | Retired |                                     |
| LastName  | string  | FamilyName | string        | Renamed |                                     |
|           |         | LegalName  | string        |   New   |                                     |
| Nickname  | string  | KnownAs    | string        | Renamed |                                     |
| Residence | Address | Residence  | PostalAddress | Identical | Address renamed to PostalAddress.   |

* Sorting is alphabetical by this.Name if present, next.Name if not.
* Package versions are used as headers.
* Property renames already specified in generator config are respected.
* A status column gives a key for easy scanning.
* Multiple statuses are comma-separated, as in the summary table.
* A notes column explains differences that do not represent modifications, such as a configured rename of a referenced type.

When both _retired_ and _new_ properties exist (and thus we have a potential opportunity for a rename), use the notes column to propose the closest match for each retirement. Reuse the existing TypoAdvisor for this, omitting suggestions below a minimum similarity threshold.

Configured type renames follow conversion-graph semantics: an exact name match in the next package takes precedence over `$nameInNextVersion`.

## Status

Proposed.

## Consequences

## Experience Report

## References
