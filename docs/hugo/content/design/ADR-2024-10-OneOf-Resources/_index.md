---
title: '2024-10 OneOf Resouces'
toc_hide: true
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

### Background

## Requirements

<!-- 
## Options (optional) 

### Option 1: Foo

**Pros:**

- Pro 1

**Cons:**

- Con 1

### Option 2: Bar

**Pros:**

- Pro 1

**Cons:**

- Con 1

-->

## Decision

<!--
### FAQ 

Q: Q1

A: A1
-->

## Status

Proposed.

## Consequences

TBC

## Experience Report

TBC

## References

None
