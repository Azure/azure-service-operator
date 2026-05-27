
---
title: '2026-05: Templating'
toc_hide: true
---

## Context

ARM templates support expressions in any property, allowing values to be composed dynamically at deployment time. This flexibility is a long-standing feature of the ARM ecosystem, and users coming to Azure Service Operator (ASO) from an ARM background sometimes ask whether ASO will offer the same capability for Kubernetes custom resources (e.g. [#2950](https://github.com/Azure/azure-service-operator/issues/2950))

In a Kubernetes context, templating is a more complex proposition. CRDs are validated by the API server at admission time using OpenAPI schemas, and the broader Kubernetes ecosystem (kubectl, GitOps tools, policy engines, IDE tooling) relies on those schemas being honest about the shape and type of each property. Embedding expressions inside arbitrary properties conflicts with these expectations, and runs counter to the standing prohibition on polymorphic properties in Kubernetes APIs - even the polymorphic properties we already have are now considered a mistake.

Templating is also already well served elsewhere in the Kubernetes ecosystem. Tools like [KRO](https://kro.run), [Helm](https://helm.sh), and [Kustomize](https://kustomize.io) exist specifically to address it, and users routinely combine them with operators like ASO.

## Requirements

Any templating story for ASO would need to:

- Preserve the strong, per-property validation that Kubernetes users expect from a CRD.
- Provide clear, immediate feedback when a user submits an invalid resource, rather than deferring errors to status conditions.
- Avoid conflicting with expressions that are intended for evaluation by ARM itself.
- Remain maintainable as the number of generated resources and properties continues to grow.

## Options

### Option 1: Adopt ARM-style expressions on every property

Allow any property on any ASO resource to accept an expression string, which ASO would evaluate before sending the payload to ARM. In practice this would mean re-typing every property as a string and stripping out validation from the generated model, so the controller could walk the resource looking for expressions to substitute.

**Pros:**

- Familiar to users with an ARM background.
- Maximally flexible - any property can be computed.

**Cons:**

- Violates the Kubernetes convention against polymorphic properties. Every property would effectively need to be typed as a string, removing the validation that makes CRDs useful.
- Minor errors would not be caught on submission. Users would have to know to inspect resource conditions to discover problems, which is contrary to the immediate feedback Kubernetes users expect everywhere else. Crucially, ASO has historically gone to great lengths to feel Kubernetes native in use, and this would be a significant step away from that.
- ASO cannot reliably distinguish an expression intended for ARM from one intended for ASO. We might intercept expressions ARM was meant to evaluate, or forward malformed expressions to ARM. Either case would surface unhelpful errors.
- Tooling that relies on CRD schemas (policy engines, validators, IDEs) would lose most of its value against ASO resources.

### Option 2: Add a parallel "expression" property for every field

Mirror the pattern we already use for config maps and secrets: keep the existing strongly-typed property and add a sibling property that accepts an expression. Existing validation is preserved, but the API is extended dramatically.

For example, we currenty have the suffixes `FromConfigMap` and `FromSecret` to indicate that a property is sourcing its value from a config map or secret. We could introduce a new suffix like `FromExpression` to indicate that a property is sourcing its value from an ASO-evaluated expression.

**Pros:**

- Preserves existing validation for users who do not need templating.
- Keeps expressions clearly separated from literal values.

**Cons:**

- Roughly doubles the number of properties on every object and sub-object, purely to serve the minority of scenarios where an expression is warranted.
- Significantly increases CRD size, which is already a [known constraint]({{< relref "ADR-2023-02-Helm-Chart-Size-Limitations" >}}).
- Adds substantial complexity to the code generator, conversion functions, and documentation, with ongoing maintenance cost.
- Does not solve the underlying user problem any better than existing OSS tools already do.
- Is not pay-to-play as all users will need to navigate the expanded surface area even if they don't need templating.

### Option 3: Rely on the OSS ecosystem

Treat templating as an orthogonal concern and direct users to existing Kubernetes-native composition and templating tools when they need it.

**Pros:**

- Keeps ASO focused on its core responsibility: faithfully projecting Azure resources into Kubernetes.
- Preserves strong, per-property validation and immediate feedback for ASO resources.
- Users can pick the templating tool that best fits their workflow, rather than being tied to whatever ASO ships.
- No additional surface area in the generator, the CRDs, or the controller.

**Cons:**

- Users wanting an ARM-style experience must learn a second tool.
- ASO does not own the end-to-end templating experience, so we cannot guarantee how well it integrates.

## Decision

ASO will not provide templating out of the box. We will rely on the existing OSS ecosystem - [KRO](https://kro.run) in particular, alongside Helm, Kustomize, and similar tools - to cover scenarios that need value composition or expression evaluation.

This keeps ASO aligned with Kubernetes conventions around strongly-typed, schema-validated resources, and avoids a large, ongoing investment in a feature that is already well served elsewhere.

## Status

Proposed.

## Consequences

- ASO documentation will point users at OSS templating and composition tools rather than offering a built-in equivalent.
- We will not accept feature requests for ARM-style expressions on ASO properties, and we will close existing ones with a reference to this ADR.
- The code generator and CRDs remain focused on faithfully representing Azure resources, with no parallel "expression" surface to maintain.

## Experience Report

TBC

## References

- [KRO](https://kro.run)
- [Helm](https://helm.sh)
- [Kustomize](https://kustomize.io)
- [Helm Chart Size Limitations]({{< relref "ADR-2023-02-Helm-Chart-Size-Limitations" >}})
