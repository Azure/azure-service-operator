---
title: '2024-10 OneOf Resouces'
toc_hide: false
---

## Context

One of the core features of OpenAPI (aka Swagger) specifications is the ability to handle polymorphism by requiring a part of the structure to match _one-of_ several possible schemas.

Azure Service Operator (ASO) has good support for use of OneOf definitions, but during import of Kusto resources we've run into a problem where our approach for handling them is conflicting with other decisions we made elsewhere.

### Background: OneOf types

While the OneOf structure is a really useful and expressive way to design APIs, there's no direct equivilent in the Go type system, or in the way Kubernetes CRDs may be defined (the later likely being driven by the former).

We therefore handle these by creating an intermediate layer that represents the range of options, and then flatten things when constructing the payload for submission to ARM.

To illustrate, imagine our API has support for a number of different roles that a person might take on. The Swagger definition might capture the `Role` as a one-of selecting from the four available roles (Student, Tutor, Teacher, and Marker), with an associated `RoleProperties` for shared properties that are present in all roles.

``` mermaid
classDiagram

class Role {
  <<oneof>>
}

class RoleProperties {
    StartDate date
    FinishDate optional<date>
}

Role *-- RoleProperties

class StudentProperties {
    Class string
    Grade string
}


class TutorProperties {
    Class string
    Timeslot string
}

class TeacherProperties {
    Subject string
    Tenure bool
}

class MarkerProperties {
    Exam string
    Strictness string
}

Role --> StudentProperties : Student
Role --> TutorProperties : Tutor
Role --> TeacherProperties : Teacher
Role --> MarkerProperties : Marker
```

To represent this as a CRD, we create an object structure where the top level `Role` has four mutually exclusive properties, one for each option, with all the properties pushed down to the leaves.

``` mermaid
classDiagram

class Role {
  Student StudentProperties
  Tutor TutorProperties
  Teacher TeacherProperties
  Marker MarkerProperties
}

class StudentProperties {
    Class string
    Grade string
    StartDate date
    FinishDate optional<date>
}


class TutorProperties {
    Class string
    Timeslot string
    StartDate date
    FinishDate optional<date>
}

class TeacherProperties {
    Subject string
    Tenure bool
    StartDate date
    FinishDate optional<date>
}

class MarkerProperties {
    Exam string
    Strictness string
    StartDate date
    FinishDate optional<date>
}

Role *-- StudentProperties : Student
Role *-- TutorProperties : Tutor
Role *-- TeacherProperties : Teacher
Role *-- MarkerProperties : Marker
```

### Background: Resource Structure

When we generate the ARM types for submission to Azure Resource Manager, we require each resource to have a `Name` property.

If that name is missing, we get the following error:

```
 error generating code: 
 failed to execute stage 38: 
 Create types for interaction with ARM: 
 unable to create arm resource spec definition for resource <id>: 
 resource spec doesn't have "Name" property"
```

### The Problem

When importing resources from `Microsoft.Kusto` we have, for the first time, a resource using a _One-Of_ as a part of it's definition right at the root. There are two flavours of the `Database` resource, one for a _read/write_ database, and another for a _following_ database.

``` mermaid
classDiagram

class ClustersDatabase_Spec {
    <<oneof>>
    Name string
}

class ReadWriteDatabase
class ReadOnlyFollowingDatabase

ClustersDatabase_Spec --> ReadWriteDatabase : ReadWrite
ClustersDatabase_Spec --> ReadOnlyFollowingDatabase : ReadOnlyFollowing
```

When the OneOf is rendered according to our current rules, we end up with this structure:

``` mermaid
classDiagram

class ClustersDatabase_Spec {
    <<oneof>>
    ReadWrite ReadWriteDatabase
    ReadOnlyFollowing ReadOnlyFollowingDatabase
}

class ReadWriteDatabase {
    Name string
}

class ReadOnlyFollowingDatabase {
    Name string
}

ClustersDatabase_Spec *-- ReadWriteDatabase
ClustersDatabase_Spec *-- ReadOnlyFollowingDatabase
```

This is correct according to our rules for OneOf, but it doesn't work for ARM Spec generation due to the lack of `Name`.

#### Other factors

We also have an `AzureName` property for many our custom resources, to allow the name in Azure to differ from the name in the cluster, due to different naming rules in each environment.

## Option 1: Do Nothing

Accept that the current decisions mean that we have to decline to support the `Microsoft.Kusto` resources, and any future resources that use this pattern.

### Pros

* Easy

### Cons

* Distasteful to decline to support a resource due to a technical limitation, especially since we have a customer ask for this resource.
* No guarantee that Kusto will be the only affected resource.
* Risk that resources we already support will use this pattern in a new API version, rendering us unable to upgrade.

## Option 2: Loosen the rules on Name

For the specific case where the top level `spec` is a one-of, permit the `Name` to be omitted from the top level as long as it's present on all of the leaf types (guaranteeing that a name will always be present).

### Pros

* Avoids making already complex one-of handling more complex.

### Cons

* May be confusing to users.
* Changes the existing rules, potentially breaking other code.
* Requires changes to the code generator to handle implementation of `AzureName()` and `SetAzureName()` methods on one-of resource types.

### Questions

* How much of the existing operation of the code generator relies on the presence of `Name` at the top level on the ARM spec?
* Ditto for the controller?

## Option 3: Change the rules for all OneOf Types

At the moment, the only permitted properties at the root level of a one-of are the mutually exclusive properties that represent the available options. We could loosen this rule to permit other properties, those in common to all leaf types, to be present at the root level.

### Pros

* Conceptually simpler

### Cons

* Changing already complex one-of handling to make it more complex.
* Requires changing the way we serialize/deserialize one-of types in non-trivial ways.
* Potentially large blast-radius if our changes impact on one-of types we've already generated and released.

## Option 4: Special case Name for root OneOf Types

Preserve the existing rules for one-of types, but special case the root level of a resource spec to permit `Name` be specified alongside the one-of properties.

### Pros

* Limits the scope of impact of the change.

### Cons

* Need to change the generation of our JSON marshalling code to handle this special case.
* Special casing is always a bit of a code smell.

### Questions

* Are there other properties that might be required at the root level of a resource spec in the future?
* Do we special case `Name` by itself, or do just apply different rules for root one-of objects?

## Option 5: Split the resources

Split the one-of resource into multiple variants, each representing one altnerative.

For example, for `kusto` we'd replace `Database` with `ReadWriteDatabase` and `ReadOnlyFollowingDatabase`, two resources that happened to use the same ARM URL but have different properties.

### Pros

* Conceptually simple

### Cons

* Choosing good names for the split resources may be difficult to code
* Increases the cognitive distance between ARM API and CRD structure, making it harder to understand
* Issues with ownership of child resources

## Decision

Proposed: Option 2: Loosen the rules on Name

## Status

Discussion.

## Consequences

TBC

## Experience Report

TBC

## References

None
