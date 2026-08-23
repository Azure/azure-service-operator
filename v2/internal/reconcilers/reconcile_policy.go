/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package reconcilers

import (
	"reflect"

	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
)

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
	return oldStr == string(annotations.ReconcilePolicySkip) || newStr == string(annotations.ReconcilePolicySkip)
}
