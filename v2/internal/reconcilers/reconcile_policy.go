/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package reconcilers

import (
	"reflect"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/common"
)

type ReconcilePolicy string

// ParseReconcilePolicy parses the provided reconcile policy.
func ParseReconcilePolicy(policy string) (ReconcilePolicy, error) {
	switch policy {
	case "":
		return common.ReconcilePolicyManage, nil
	case string(common.ReconcilePolicyManage):
		return common.ReconcilePolicyManage, nil
	case string(common.ReconcilePolicySkip):
		return common.ReconcilePolicySkip, nil
	case string(common.ReconcilePolicyDetachOnDelete):
		return common.ReconcilePolicyDetachOnDelete, nil
	default:
		// Defaulting to manage.
		return common.ReconcilePolicyManage, errors.Errorf("%q is not a known reconcile policy", policy)
	}
}

// AllowsDelete determines if the policy allows deletion of the backing Azure resource
func (r ReconcilePolicy) AllowsDelete() bool {
	return r == common.ReconcilePolicyManage
}

// AllowsModify determines if the policy allows modification of the backing Azure resource
func (r ReconcilePolicy) AllowsModify() bool {
	return r == common.ReconcilePolicyManage || r == common.ReconcilePolicyDetachOnDelete
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
	return oldStr == string(common.ReconcilePolicySkip) || newStr == string(common.ReconcilePolicySkip)
}
