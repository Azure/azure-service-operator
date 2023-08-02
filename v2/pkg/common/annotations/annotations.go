// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package annotations

const PerResourceSecretAnnotation = "serviceoperator.azure.com/credential-from"

// ReconcilePolicyAnnotation describes the reconcile policy for the resource in question.
// A reconcile policy describes what action (if any) the operator is allowed to take when
// reconciling the resource.
// If no reconcile policy is specified, the default is "run"
const ReconcilePolicyAnnotation = "serviceoperator.azure.com/reconcile-policy"
