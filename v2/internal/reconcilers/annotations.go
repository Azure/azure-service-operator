/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package reconcilers

import (
	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	"github.com/Azure/azure-service-operator/v2/internal/util/annotations"
)

// ARMReconcilerAnnotationChangedPredicate creates a predicate that emits events when annotations
// interesting to the generic ARM reconciler are changed
func ARMReconcilerAnnotationChangedPredicate(log logr.Logger) predicate.Predicate {
	return annotations.MakeSelectAnnotationChangedPredicate(
		log,
		map[string]annotations.HasAnnotationChanged{
			ReconcilePolicyAnnotation: HasReconcilePolicyAnnotationChanged,
		})
}
