/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package reconcilers

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	. "github.com/Azure/azure-service-operator/v2/internal/logging"
	"github.com/Azure/azure-service-operator/v2/internal/ownerutil"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

const GenericControllerFinalizer = "serviceoperator.azure.com/finalizer"

// LogObj logs the obj
func LogObj(log logr.Logger, note string, obj genruntime.MetaObject) {
	if log.V(Debug).Enabled() {
		// This could technically select annotations from other Azure operators, but for now that's ok.
		// In the future when we no longer use annotations as heavily as we do now we can remove this or
		// scope it to a finite set of annotations.
		ourAnnotations := make(map[string]string)
		for key, value := range obj.GetAnnotations() {
			if strings.HasSuffix(key, ".azure.com") {
				ourAnnotations[key] = value
			}
		}

		keysAndValues := []interface{}{
			"kind", obj.GetObjectKind(),
			"resourceVersion", obj.GetResourceVersion(),
			"generation", obj.GetGeneration(),
			"uid", obj.GetUID(),
			"ownerReferences", obj.GetOwnerReferences(),
			"creationTimestamp", obj.GetCreationTimestamp(),
			"finalizers", obj.GetFinalizers(),
			"annotations", ourAnnotations,
			// Use fmt here to ensure the output uses the String() method, which log.Info doesn't seem to do by default
			"conditions", fmt.Sprintf("%s", obj.GetConditions()),
		}

		if armObj, ok := obj.(genruntime.ARMMetaObject); ok {
			keysAndValues = append(keysAndValues, "owner", armObj.Owner())
		}

		// Log just what we're interested in. We avoid logging the whole obj
		// due to possible risk of disclosing secrets or other data that is "private" and users may
		// not want in logs.
		log.V(Debug).Info(note, keysAndValues...)
	}
}

type ARMOwnedResourceReconcilerCommon struct {
	ReconcilerCommon
	ResourceResolver *resolver.Resolver
}

// IsOwnerReady returns true if the owner is ready or if there is no owner required
func (r *ARMOwnedResourceReconcilerCommon) IsOwnerReady(ctx context.Context, log logr.Logger, obj genruntime.ARMOwnedMetaObject) (bool, error) {
	_, err := r.ResourceResolver.ResolveOwner(ctx, obj)
	if err != nil {
		var typedErr *resolver.ReferenceNotFound
		if errors.As(err, &typedErr) {
			log.V(Info).Info("Owner does not yet exist", "NamespacedName", typedErr.NamespacedName)
			return false, nil
		}

		return false, errors.Wrap(err, "failed to get owner")
	}

	return true, nil
}

func (r *ARMOwnedResourceReconcilerCommon) ApplyOwnership(ctx context.Context, log logr.Logger, obj genruntime.ARMOwnedMetaObject) error {
	owner, err := r.ResourceResolver.ResolveOwner(ctx, obj)
	if err != nil {
		return errors.Wrap(err, "failed to get owner")
	}

	if owner == nil {
		return nil
	}

	ownerRef := ownerutil.MakeOwnerReference(owner)

	obj.SetOwnerReferences(ownerutil.EnsureOwnerRef(obj.GetOwnerReferences(), ownerRef))
	log.V(Info).Info(
		"Set owner reference",
		"ownerGvk", owner.GetObjectKind().GroupVersionKind(),
		"ownerName", owner.GetName())
	err = r.CommitUpdate(ctx, log, obj)

	if err != nil {
		return errors.Wrap(err, "update owner references failed")
	}

	return nil
}

// ClaimResource adds a finalizer and ensures that the owner reference is set
func (r *ARMOwnedResourceReconcilerCommon) ClaimResource(ctx context.Context, log logr.Logger, obj genruntime.ARMMetaObject) (ctrl.Result, error) {
	log.V(Info).Info("applying ownership")
	isOwnerReady, err := r.IsOwnerReady(ctx, log, obj)
	if err != nil {
		return ctrl.Result{}, err
	}

	if !isOwnerReady {
		err = errors.Errorf("Owner %q cannot be found. Progress is blocked until the owner is created.", obj.Owner().String())
		err = conditions.NewReadyConditionImpactingError(err, conditions.ConditionSeverityWarning, conditions.ReasonWaitingForOwner)
		return ctrl.Result{}, err
	}

	// Adding the finalizer should happen in a reconcile loop prior to the PUT being sent to Azure to avoid situations where
	// we issue a PUT to Azure but the commit of the resource into etcd fails, causing us to have an unset
	// finalizer and have started resource creation in Azure.
	log.V(Info).Info("adding finalizer")
	controllerutil.AddFinalizer(obj, GenericControllerFinalizer)

	// Short circuit here if there's no owner management to do
	if obj.Owner() == nil {
		err = r.CommitUpdate(ctx, log, obj)
		err = client.IgnoreNotFound(err)
		if err != nil {
			return ctrl.Result{}, errors.Wrap(err, "updating resource error")
		}

		return ctrl.Result{Requeue: true}, nil
	}

	err = r.ApplyOwnership(ctx, log, obj)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Fast requeue as we're moving to the next stage
	return ctrl.Result{Requeue: true}, nil
}

func (r *ARMOwnedResourceReconcilerCommon) NeedToClaimResource(obj genruntime.ARMOwnedMetaObject) bool {
	owner := obj.Owner()
	unresolvedOwner := owner != nil && len(obj.GetOwnerReferences()) == 0
	unsetFinalizer := !controllerutil.ContainsFinalizer(obj, GenericControllerFinalizer)

	return unresolvedOwner || unsetFinalizer
}

type ReconcilerCommon struct {
	KubeClient         kubeclient.Client
	PositiveConditions *conditions.PositiveConditionBuilder
}

func (r *ReconcilerCommon) CommitUpdate(ctx context.Context, log logr.Logger, obj genruntime.MetaObject) error {
	err := r.KubeClient.CommitObject(ctx, obj)
	if err != nil {
		return err
	}
	LogObj(log, "updated resource in etcd", obj)
	return nil
}

func (r *ReconcilerCommon) WriteReadyConditionError(ctx context.Context, obj genruntime.MetaObject, err *conditions.ReadyConditionImpactingError) error {
	conditions.SetCondition(obj, r.PositiveConditions.Ready.ReadyCondition(
		err.Severity,
		obj.GetGeneration(),
		err.Reason,
		err.Error()))
	commitErr := client.IgnoreNotFound(r.KubeClient.CommitObject(ctx, obj))
	if commitErr != nil {
		return errors.Wrap(commitErr, "updating resource error")
	}

	if err.Severity == conditions.ConditionSeverityError {
		// This is a bit weird, but fatal errors shouldn't trigger a fresh reconcile, so
		// returning nil results in reconcile "succeeding" meaning an event won't be
		// queued to reconcile again.
		return nil
	}

	return err
}
