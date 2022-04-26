---
title: Committing Generated Code
---
# Committing Generated Code

## Context

We’re aware that common practice avoids committing generated code into source control, treating it as an ephemeral resource that can be regenerated on demand.

Even so, we are concluding that we want to check in our generated code.

### Upstream changes

We do not fully control the source definitions from which we are generating our object model and we are concerned about upstream changes that change our generated output in unexpected ways.

While such changes are very unlikely, given the controls that ARM has in place around breaking API changes, exceptions are made, and mistakes do happen.

As a part of the Kubernetes ecosystem of tooling, we need to adhere to strong guidelines for backward compatibility; inadvertently releasing an incompatible version of ASO would be a significant issue. It’s likely that recovering from such a release would require a second breaking release (to restore compatibility with the original) which would itself cause customer pain.

Having the generated code committed into our repository gives us an opportunity to detect and address such changes before a new release of ASO.

It is also useful to have the assurance of knowing exactly what generated code was included in a particular release of ASO. Our ability to reproduce and debug a reported issue would be much reduced if an upstream modification (resulting in a change to our generated code) meant that we could not reproduce the exact code being run by a customer.

### Generator changes

With so much code being generated, we are starting to become concerned that a simple change to the generator might have unexpected consequences, perhaps due to reuse of a function definition, for example.

Having the generated code committed into our repository would give an opportunity to review those changes in advance, helping to catch any issues well before any code is merged into our repository. 

We admit that reviewing all the generated code is likely to be infeasible, but we can at least benefit from random sampling of the changes, and by reviewing at least one resource in depth.

### Generator awareness

We want the project to be accessible to outside developers.

Requiring a new developer to successfully run the code generator before being able to build the operator itself would be a significant barrier to anyone trying to get an understanding of what we’ve built.

Having the generated code committed into our repository means that a full build of the service operator can be performed without needing to even be aware that the generator exists, reducing the effort required for someone to get started.

## Decision

We will commit the generated code in a separate folder.

## Status

Adopted.

First commit of our generated code was in [PR #1741](https://github.com/Azure/azure-service-operator/pull/1741).

## Consequences

We have observed unexpected consequences of generator changes and been able to correct course before creating a PR.

Some PRs have had a wide reaching impact, resulting in very large PRs with lots of modifications to generated code. To facilitate review, some pieces of work have been split into two PRs: One with a fully tested new feature in disabled form, allowing detailed review of the change, and another to enable the feature and regenerate affected code.

## Experience Report

TBC.

## References

TBC.
