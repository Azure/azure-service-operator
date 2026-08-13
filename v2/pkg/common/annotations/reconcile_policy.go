// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package annotations

import "github.com/rotisserie/eris"

// ReconcilePolicy describes the reconcile policy for the resource in question.
// A reconcile policy describes what action (if any) the operator is allowed to take when
// reconciling the resource.
// If no reconcile policy is specified, the default is "manage"
const ReconcilePolicy = "serviceoperator.azure.com/reconcile-policy"

type ReconcilePolicyValue string

const (
	// ReconcilePolicyManage instructs the operator to manage the resource in question.
	// This includes issuing PUTs to update it and DELETE's to delete it from Azure if deleted in Kubernetes.
	// This is the default policy when no policy is specified.
	ReconcilePolicyManage = ReconcilePolicyValue("manage")

	// ReconcilePolicySkip instructs the operator to skip all reconciliation actions. This includes creating
	// the resource.
	ReconcilePolicySkip = ReconcilePolicyValue("skip")

	// ReconcilePolicyDetachOnDelete instructs the operator to skip deletion of resources in Azure. This allows
	// deletion of the resource in Kubernetes to go through but does not delete the underlying Azure resource.
	ReconcilePolicyDetachOnDelete = ReconcilePolicyValue("detach-on-delete")
)

// ReconcilePolicies are the reconcile policies in effect during a reconcile, none of which a resource
// extension can work out on its own. Read them with ReconcilePolicyFromContext or
// ReconcilePolicyForAnnotation rather than combining them.
type ReconcilePolicies struct {
	// Effective is the policy for the resource being reconciled, resolved from its own annotation, its
	// namespace, and the operator's configuration.
	Effective ReconcilePolicyValue

	// Inherited is the policy a resource in this namespace carrying no annotation of its own would get.
	Inherited ReconcilePolicyValue

	// Default is the policy the operator is configured with, which an unusable annotation falls back to.
	Default ReconcilePolicyValue
}

// ForAnnotation returns the reconcile policy for a resource other than the one being reconciled, given
// that resource's own reconcile-policy annotation.
func (r ReconcilePolicies) ForAnnotation(annotation string) ReconcilePolicyValue {
	if annotation == "" {
		return r.Inherited
	}

	policy, _ := ParseReconcilePolicy(annotation, r.Default)
	return policy
}

// ParseReconcilePolicy parses provided reconcile policy.
// defaultPolicyValue is read from DEFAULT_RECONCILE_POLICY env variable or set to 'manage' when missing
func ParseReconcilePolicy(policy string, defaultReconcilePolicy ReconcilePolicyValue) (ReconcilePolicyValue, error) {
	// policy is read from CR annotation, if it's empty it being read from defaultReconcilePolicy
	switch policy {
	case "":
		return defaultReconcilePolicy, nil
	case string(ReconcilePolicyManage):
		return ReconcilePolicyManage, nil
	case string(ReconcilePolicySkip):
		return ReconcilePolicySkip, nil
	case string(ReconcilePolicyDetachOnDelete):
		return ReconcilePolicyDetachOnDelete, nil
	default:
		return defaultReconcilePolicy, eris.Errorf("%q is not a known reconcile policy", policy)
	}
}

// AllowsDelete determines if the policy allows deletion of the backing Azure resource
func (r ReconcilePolicyValue) AllowsDelete() bool {
	return r == ReconcilePolicyManage
}

// AllowsModify determines if the policy allows modification of the backing Azure resource
func (r ReconcilePolicyValue) AllowsModify() bool {
	return r == ReconcilePolicyManage || r == ReconcilePolicyDetachOnDelete
}
