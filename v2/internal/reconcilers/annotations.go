/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package reconcilers

import (
	"reflect"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	"github.com/Azure/azure-service-operator/v2/internal/util/annotations"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

const PerResourceSecretAnnotation = "serviceoperator.azure.com/credential-from"

// ARMReconcilerAnnotationChangedPredicate creates a predicate that emits events when annotations
// interesting to the generic ARM reconciler are changed
func ARMReconcilerAnnotationChangedPredicate(log logr.Logger) predicate.Predicate {
	return annotations.MakeSelectAnnotationChangedPredicate(
		log,
		map[string]annotations.HasAnnotationChanged{
			ReconcilePolicyAnnotation: HasReconcilePolicyAnnotationChanged,
		})
}

// ARMPerResourceSecretAnnotationChangedPredicate creates a predicate that emits events when annotations
// interesting to the generic ARM reconciler are changed
func ARMPerResourceSecretAnnotationChangedPredicate(log logr.Logger) predicate.Predicate {
	return annotations.MakeSelectAnnotationChangedPredicate(
		log,
		map[string]annotations.HasAnnotationChanged{
			PerResourceSecretAnnotation: HasAnnotationChanged,
		})
}

// GetReconcilePolicy gets the reconcile-policy from the ReconcilePolicyAnnotation
func GetReconcilePolicy(obj genruntime.MetaObject, log logr.Logger) ReconcilePolicy {
	policyStr := obj.GetAnnotations()[ReconcilePolicyAnnotation]
	policy, err := ParseReconcilePolicy(policyStr)
	if err != nil {
		log.Error(
			err,
			"failed to get reconcile policy. Applying default policy instead",
			"chosenPolicy", policy,
			"policyAnnotation", policyStr)
	}

	return policy
}

// HasAnnotationChanged returns true if the annotation has changed
func HasAnnotationChanged(old *string, new *string) bool {
	return !reflect.DeepEqual(old, new)
}
