/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package reconcilers

import (
	"reflect"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/common/reconciler"
)

// ParseReconcilePolicy parses the provided reconcile policy.
func ParseReconcilePolicy(policy string) (reconciler.ReconcilePolicy, error) {
	switch policy {
	case "":
		return reconciler.ReconcilePolicyManage, nil
	case string(reconciler.ReconcilePolicyManage):
		return reconciler.ReconcilePolicyManage, nil
	case string(reconciler.ReconcilePolicySkip):
		return reconciler.ReconcilePolicySkip, nil
	case string(reconciler.ReconcilePolicyDetachOnDelete):
		return reconciler.ReconcilePolicyDetachOnDelete, nil
	default:
		// Defaulting to manage.
		return reconciler.ReconcilePolicyManage, errors.Errorf("%q is not a known reconcile policy", policy)
	}
}

// HasReconcilePolicyAnnotationChanged returns true if the reconcile-policy annotation has
// changed in a way that needs to trigger a reconcile.
func HasReconcilePolicyAnnotationChanged(old *string, new *string) bool {
	equal := reflect.DeepEqual(old, new)
	if equal {
		// If the annotations are equal there's been no change
		return false
	}

	oldStr := to.Value(old)
	newStr := to.Value(new)

	// We only care about transitions to or from ReconcilePolicySkip. We don't need to
	// trigger an event if ReconcilePolicyDetachOnDelete is added or removed, as that annotation
	// only applies on delete (which we will always run reconcile on).
	return oldStr == string(reconciler.ReconcilePolicySkip) || newStr == string(reconciler.ReconcilePolicySkip)
}
