# Explicit Reconciliation Policy Design

## Problem

The branch currently records reconciliation policy information on `context.Context` so post-reconciliation extensions can decide whether they may modify Azure resources. This preserves policy semantics, but it hides a required dependency: callers and implementations compile without acknowledging that policy affects their behavior.

The refactor will preserve the branch's policy-resolution behavior while replacing context-based transport with an explicit parameter.

## Goals

- Make reconciliation policy a visible dependency of every post-reconciliation check.
- Preserve support for checking both the reconciled resource and another resource, such as its owner.
- Resolve object, namespace, and operator policy precedence in one place.
- Preserve fail-closed handling when a namespace cannot be read.
- Remove context accessors and their implicit fallback behavior.

## Non-goals

- Changing reconcile-policy precedence or accepted annotation values.
- Changing delete reconciliation, which does not invoke post-reconciliation checks.
- Refactoring unrelated reconciler interfaces or extension points.
- Modifying generated API types.

## Shared Policy Value

Move `ReconcilePolicies` to `pkg/common/annotations`, alongside `ReconcilePolicyValue`. It remains an immutable value with three fields:

- `Effective`: the resolved policy for the resource being reconciled.
- `Inherited`: the policy an unannotated resource in the same namespace would receive.
- `Default`: the operator-configured fallback for an invalid non-empty annotation.

Add `ForAnnotation(annotation string) ReconcilePolicyValue` to the value. An empty annotation returns `Inherited`; a non-empty annotation is parsed against `Default`. This keeps secondary-resource policy resolution centralized without requiring context or duplicating precedence rules in extensions.

Move `ParseReconcilePolicy` from `internal/reconcilers` into `pkg/common/annotations` as part of this change. Both `ForAnnotation` and `GenericReconciler.mergeReconcilePolicy` use that shared parser, avoiding duplicated rules and an import cycle between `genruntime` and the internal reconciler packages.

The zero value is not assigned an implicit `manage` meaning. Reconciliation and tests must provide the resolved value explicitly.

## Resolution and Data Flow

`GenericReconciler.mergeReconcilePolicy` remains the only code that reads the reconciled object's annotation, its namespace annotation, and operator configuration. It returns `annotations.ReconcilePolicies`.

The existing branch behavior remains:

- The operator's empty configured default is normalized to `manage`.
- The reconciled object's effective policy follows object, namespace, then operator precedence.
- An invalid annotation falls back to the operator's default.
- `Inherited` is `skip` when the namespace cannot be read.

The generic reconciler resolves policies once per create/update reconcile:

1. If `Effective.AllowsModify()` is true, pass the policies to `genruntime.Reconciler.CreateOrUpdate`.
2. Otherwise, pass the same policies to `genruntime.Reconciler.UpdateStatus`.
3. Do not add a policy parameter to `Delete`, because that path does not invoke a post-reconciliation check.

The ARM reconciler threads the value unchanged through its create/update success handling and status-only success handling to `postReconciliationCheck`.

## Interface Changes

Add `annotations.ReconcilePolicies` as an explicit parameter to:

- `genruntime.Reconciler.CreateOrUpdate`
- `genruntime.Reconciler.UpdateStatus`
- `extensions.PostReconciliationChecker.PostReconcileCheck`
- `extensions.PostReconcileCheckFunc`
- The ARM reconciler's internal success and post-check functions where needed

The checker factory passes the exact value received by the ARM reconciler to the extension. `alwaysSucceed` and `next` carry the same parameter so future checker chains cannot lose or replace the policy dependency.

All non-ARM reconcilers implement the revised `genruntime.Reconciler` interface and may ignore the parameter. All existing post-reconciliation extensions update their signatures; extensions that do not act on Azure may also ignore it.

The context-based `WithReconcilePolicies`, `ReconcilePolicyFromContext`, and `ReconcilePolicyForAnnotation` functions and their context key are removed.

## Extension Usage

A post-reconciliation extension uses:

- `policies.Effective` to decide whether it may modify the resource being reconciled.
- `policies.ForAnnotation(other.GetAnnotations()[annotations.ReconcilePolicy])` to decide whether it may modify another resource governed by the same namespace and operator configuration.

Extensions must still establish that another resource belongs to the same operator before applying this operator's policy configuration to it.

## Error Handling

Policy resolution continues to report invalid annotations through the existing reconciler logging behavior while returning the configured fallback policy. Namespace read failures continue to log and fail closed for `Inherited`.

Removing the context fallback intentionally eliminates success-shaped behavior when a caller forgets to propagate policies. Missing propagation becomes a compile-time interface error rather than silently behaving as `manage`.

## Testing

Retain and adapt the branch's table tests for:

- Object, namespace, and operator precedence.
- Invalid annotation fallback.
- Unreadable namespace behavior.
- Fail-closed inherited policy.

Replace the context accessor tests with table tests for `ReconcilePolicies.ForAnnotation`.

Add focused propagation tests proving:

- The generic reconciler passes the resolved policies to `CreateOrUpdate` when modification is allowed.
- The generic reconciler passes the same resolved policies to `UpdateStatus` when modification is blocked.
- The post-reconciliation checker factory passes the exact policies value to an extension and through `next`.

Compilation updates cover every `genruntime.Reconciler` implementation and every existing `PostReconciliationChecker` implementation.

## Documentation

Update the post-reconciliation extension documentation to direct authors to the explicit `policies` parameter instead of context helpers. Keep the guidance that checks run on the status-only path and must respect policy before invoking ARM actions.
