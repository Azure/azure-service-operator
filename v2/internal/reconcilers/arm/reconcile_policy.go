/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package arm

import (
	"reflect"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/pkg/errors"
)

// ReconcilePolicyAnnotation describes the reconcile policy for the resource in question.
// A reconcile policy describes what action (if any) the operator is allowed to take when
// reconciling the resource.
// If no reconcile policy is specified, the default is "run"
const ReconcilePolicyAnnotation = "serviceoperator.azure.com/reconcile-policy"

type ReconcilePolicy string

const (
	// ReconcilePolicyManage instructs the operator to manage the resource in question.
	// This includes issuing PUTs to update it and DELETE's to delete it from Azure if deleted in Kuberentes.
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

// AllowsDelete determines if the policy allows deletion of the backing Azure resource
func (r ReconcilePolicy) AllowsDelete() bool {
	return r == ReconcilePolicyManage
}

// AllowsModify determines if the policy allows modification of the backing Azure resource
func (r ReconcilePolicy) AllowsModify() bool {
	return r == ReconcilePolicyManage || r == ReconcilePolicyDetachOnDelete
}

// HasReconcilePolicyAnnotationChanged returns true if the reconcile-policy annotation has
// changed in a way that needs to trigger a reconcile.
func HasReconcilePolicyAnnotationChanged(old *string, new *string) bool {
	equal := reflect.DeepEqual(old, new)
	if equal {
		// If the annotations are equal there's been no change
		return false
	}

	oldStr := to.String(old)
	newStr := to.String(new)

	// We only care about transitions to or from ReconcilePolicySkip. We don't need to
	// trigger an event if ReconcilePolicyDetachOnDelete is added or removed, as that annotation
	// only applies on delete (which we will always run reconcile on).
	return oldStr == string(ReconcilePolicySkip) || newStr == string(ReconcilePolicySkip)
}
