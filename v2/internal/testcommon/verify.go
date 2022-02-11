/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"

	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

type Verify struct {
	kubeClient client.Client
}

func NewVerify(c client.Client) *Verify {
	return &Verify{
		kubeClient: c,
	}
}

// HasState verifies that the provisioning state of the resource is in the target state.
func (e *Verify) HasState(
	ctx context.Context,
	obj client.Object,
	desiredState metav1.ConditionStatus,
	desiredSeverity conditions.ConditionSeverity,
	oldGeneration int64) (bool, error) {

	key := client.ObjectKeyFromObject(obj)

	// In order to ensure that "old state" is cleared out from obj, we need to:
	// 1. construct a newObj of the same type.
	// 2. deserialize the kubeclient into newObj.
	// 3. Use DeepCopyInto to copy newObj into obj (this ensures that nils from newObj make it onto obj, which
	//    doesn't happen normally with JSON deserialization).
	newObj, err := genruntime.NewObjectFromExemplar(obj, e.kubeClient.Scheme())
	if err != nil {
		return false, err
	}

	err = e.kubeClient.Get(ctx, key, newObj)
	if err != nil {
		return false, err
	}

	// Assign the new stuff to obj by calling DeepCopyInto.
	// This is relatively "safe" because DeepCopyObject() is a method required
	// to reconcile a resource at all, and it internally uses DeepCopyInto. Places where
	// this could fail are on types that don't use deepcopy-gen
	reflecthelpers.DeepCopyInto(newObj, obj)

	conditioner, ok := obj.(conditions.Conditioner)
	if !ok {
		return false, errors.Errorf("result of get was not conditions.Conditioner, was: %T", obj)
	}

	ready, ok := conditions.GetCondition(conditioner, conditions.ConditionTypeReady)
	if !ok {
		return false, nil
	}

	inGoalState := ready.Status == desiredState && ready.Severity == desiredSeverity
	hasObservedGenerationChanged := oldGeneration != ready.ObservedGeneration
	return inGoalState && hasObservedGenerationChanged, nil
}

// Deleted verifies that the object specified has been deleted
func (e *Verify) Deleted(ctx context.Context, obj client.Object) (bool, error) {
	key := client.ObjectKeyFromObject(obj)

	// Note that obj won't be modified if it's already deleted, so
	// could be "stale" after this call
	err := e.kubeClient.Get(ctx, key, obj)
	if apierrors.IsNotFound(err) {
		return true, nil
	}
	if err != nil {
		return false, err
	}

	return false, nil
}

// AllDeleted verifies that all of the specified objects are deleted
func (e *Verify) AllDeleted(ctx context.Context, objs []client.Object) (bool, error) {
	for _, obj := range objs {
		// It's possible that this is horribly inefficient. Should be good enough for now though
		deleted, err := e.Deleted(ctx, obj)
		if err != nil {
			return false, err
		}
		if !deleted {
			return false, nil
		}
	}

	return true, nil
}
