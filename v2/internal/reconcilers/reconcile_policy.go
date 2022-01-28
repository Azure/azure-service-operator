/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package reconcilers

import "github.com/pkg/errors"

// ReconcilePolicyAnnotation describes the reconcile policy for the resource in question.
// A reconcile policy describes what action (if any) the operator is allowed to take when
// reconciling the resource.
// If no reconcile policy is specified, the default is "run"
const ReconcilePolicyAnnotation = "serviceoperator.azure.com/reconcile-policy"

type ReconcilePolicy string

const (
	// ReconcilePolicyRun instructs the operator to run all normal reconciliation actions.
	// This is the default policy when no policy is specified.
	ReconcilePolicyRun = ReconcilePolicy("run")

	// ReconcilePolicySkip instructs the operator to skip all reconciliation actions. This includes creating
	// the resource.
	ReconcilePolicySkip = ReconcilePolicy("skip")

	// ReconcilePolicySkipDelete instructs the operator to skip deletion of resources in Azure. This allows
	// deletion of the resource in Kubernetes to go through
	ReconcilePolicySkipDelete = ReconcilePolicy("skip-delete") // TODO: Google config connector calls this deletePolicy: abandon
)

// ParseReconcilePolicy parses the provided reconcile policy.
func ParseReconcilePolicy(policy string) (ReconcilePolicy, error) {
	switch policy {
	case "":
		return ReconcilePolicyRun, nil
	case string(ReconcilePolicyRun):
		return ReconcilePolicyRun, nil
	case string(ReconcilePolicySkip):
		return ReconcilePolicySkip, nil
	case string(ReconcilePolicySkipDelete):
		return ReconcilePolicySkipDelete, nil
	default:
		// Defaulting to skip. The user is attempting to configure policy but has done it wrong,
		// if we default to Run we may inadvertently modify their object
		return ReconcilePolicySkip, errors.Errorf("%q is not a known reconcile policy", policy)
	}
}

// ShouldDelete determines if the policy allows deletion of the backing Azure resource
func (r ReconcilePolicy) ShouldDelete() bool {
	return r == ReconcilePolicyRun
}

// ShouldModify determines if the policy allows modification of the backing Azure resource
func (r ReconcilePolicy) ShouldModify() bool {
	return r == ReconcilePolicyRun || r == ReconcilePolicySkipDelete
}
