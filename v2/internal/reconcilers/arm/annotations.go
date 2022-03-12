/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package arm

import (
	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	"github.com/Azure/azure-service-operator/v2/internal/util/annotations"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

const (
	PollerResumeTokenAnnotation = "serviceoperator.azure.com/poller-resume-token"
	PollerResumeIDAnnotation    = "serviceoperator.azure.com/poller-resume-id"
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

// GetPollerResumeToken returns a poller ID and the poller token
func GetPollerResumeToken(obj genruntime.MetaObject) (string, string, bool) {
	token, hasResumeToken := obj.GetAnnotations()[PollerResumeTokenAnnotation]
	id, hasResumeID := obj.GetAnnotations()[PollerResumeIDAnnotation]

	return id, token, hasResumeToken && hasResumeID
}

func SetPollerResumeToken(obj genruntime.MetaObject, id string, token string) {
	genruntime.AddAnnotation(obj, PollerResumeTokenAnnotation, token)
	genruntime.AddAnnotation(obj, PollerResumeIDAnnotation, id)
}

// ClearPollerResumeToken clears the poller resume token and ID annotations
func ClearPollerResumeToken(obj genruntime.MetaObject) {
	genruntime.RemoveAnnotation(obj, PollerResumeTokenAnnotation)
	genruntime.RemoveAnnotation(obj, PollerResumeIDAnnotation)
}

// GetReconcilePolicy gets the reconcile policy from the ReconcilePolicyAnnotation
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
