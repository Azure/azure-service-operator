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
	// ReconcilePolicyManage instructs the operator to manage the resource in question.
	// This includes issuing PUTs to update it and DELETE's to delete it if deleted in Kuberentes.
	// This is the default policy when no policy is specified.
	ReconcilePolicyManage = ReconcilePolicy("manage")

	// ReconcilePolicySkip instructs the operator to skip all reconciliation actions. This includes creating
	// the resource.
	ReconcilePolicySkip = ReconcilePolicy("skip")

	// ReconcilePolicyDetachOnDelete instructs the operator to skip deletion of resources in Azure. This allows
	// deletion of the resource in Kubernetes to go through but does not delete the underlying Azure resource.
	ReconcilePolicyDetachOnDelete = ReconcilePolicy("detach-on-delete")
)

// ParseReconcilePolicy parses the provided reconcile policy.
func ParseReconcilePolicy(policy string) (ReconcilePolicy, error) {
	switch policy {
	case "":
		return ReconcilePolicyManage, nil
	case string(ReconcilePolicyManage):
		return ReconcilePolicyManage, nil
	case string(ReconcilePolicySkip):
		return ReconcilePolicySkip, nil
	case string(ReconcilePolicyDetachOnDelete):
		return ReconcilePolicyDetachOnDelete, nil
	default:
		// Defaulting to manage.
		return ReconcilePolicyManage, errors.Errorf("%q is not a known reconcile policy", policy)
	}
}

// ShouldDelete determines if the policy allows deletion of the backing Azure resource
func (r ReconcilePolicy) ShouldDelete() bool {
	return r == ReconcilePolicyManage
}

// ShouldModify determines if the policy allows modification of the backing Azure resource
func (r ReconcilePolicy) ShouldModify() bool {
	return r == ReconcilePolicyManage || r == ReconcilePolicyDetachOnDelete
}
