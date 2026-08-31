// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package annotations

import (
	"github.com/rotisserie/eris"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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

// ResolvedReconcilePolicies are the resolved reconcile policies in effect during a reconcile.
// Read them with Effective or ForResource() rather than combining them.
type ResolvedReconcilePolicies struct {
	// Effective is the policy for the resource being reconciled, resolved from its own annotation, Namespace,
	// and Global.
	Effective ReconcilePolicyValue

	// NamespacePolicy is the policy for resources in this namespace carrying no annotation of their own, resolved
	// from the namespace annotation or Global.
	NamespacePolicy ReconcilePolicyValue

	// NamespaceName is the namespace from which NamespacePolicy was resolved.
	NamespaceName string

	// Global is the policy resolved from the operator configuration or the built-in default. An unusable
	// annotation falls back to this policy.
	Global ReconcilePolicyValue
}

// ForResource returns the reconcile policy for a resource other than the one being reconciled.
// The resource must be in NamespaceName because Namespace was resolved specifically for that namespace.
func (r ResolvedReconcilePolicies) ForResource(resource metav1.Object) (ReconcilePolicyValue, error) {
	if resource.GetNamespace() != r.NamespaceName {
		return "", eris.Errorf(
			"expected resource in namespace %q, but it was in %q",
			r.NamespaceName,
			resource.GetNamespace(),
		)
	}

	annotation := resource.GetAnnotations()[ReconcilePolicy]
	if annotation == "" {
		return r.NamespacePolicy, nil
	}

	policy, _ := ParseReconcilePolicy(annotation, r.Global)
	return policy, nil
}

// ParseReconcilePolicy parses provided reconcile policy, will fallback if the value is missing.
func ParseReconcilePolicy(
	policy string,
	defaultReconcilePolicy ReconcilePolicyValue,
) (ReconcilePolicyValue, error) {
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
