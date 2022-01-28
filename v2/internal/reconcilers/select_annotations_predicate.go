/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package reconcilers

import (
	"reflect"

	"github.com/go-logr/logr"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	. "github.com/Azure/azure-service-operator/v2/internal/logging"
)

// ARMReconcilerAnnotationChangedPredicate creates a predicate that emits events when annotations
// interesting to the generic ARM reconciler are changed
func ARMReconcilerAnnotationChangedPredicate(log logr.Logger) predicate.Predicate {
	return MakeSelectAnnotationChangedPredicate(log, ReconcilePolicyAnnotation)
}

// MakeSelectAnnotationChangedPredicate creates a selectAnnotationChangedPredicate watching for
// changes to select annotations
func MakeSelectAnnotationChangedPredicate(log logr.Logger, annotations ...string) predicate.Predicate {
	annotationsMap := make(map[string]struct{})
	for _, annotation := range annotations {
		annotationsMap[annotation] = struct{}{}
	}

	return selectAnnotationChangedPredicate{
		log:         log,
		annotations: annotationsMap,
	}
}

type selectAnnotationChangedPredicate struct {
	predicate.Funcs

	log         logr.Logger
	annotations map[string]struct{}
}

var _ predicate.Predicate = selectAnnotationChangedPredicate{}

// Update implements UpdateEvent filter for annotation changes.
func (p selectAnnotationChangedPredicate) Update(e event.UpdateEvent) bool {
	if e.ObjectOld == nil {
		p.log.V(Debug).Error(nil, "Update event has no old object to update", "event", e)
		return false
	}
	if e.ObjectNew == nil {
		p.log.V(Debug).Error(nil, "Update event has no new object for update", "event", e)
		return false
	}

	newAnnotations := e.ObjectNew.GetAnnotations()
	oldAnnotations := e.ObjectOld.GetAnnotations()

	selectNew := make(map[string]string)
	selectOld := make(map[string]string)

	for k, v := range newAnnotations {
		if _, ok := p.annotations[k]; ok {
			selectNew[k] = v
		}
	}

	for k, v := range oldAnnotations {
		if _, ok := p.annotations[k]; ok {
			selectOld[k] = v
		}
	}

	return !reflect.DeepEqual(selectNew, selectOld)
}
