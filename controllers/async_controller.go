/*
Copyright 2019 Alexander Eldeib.
*/

package controllers

import (
	"context"
	"fmt"
	"time"

	telemetry "github.com/Azure/azure-service-operator/pkg/telemetry"
	multierror "github.com/hashicorp/go-multierror"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	finalizerName string        = "azure.microsoft.com/finalizer"
	requeDuration time.Duration = time.Second * 20
)

type AsyncClient interface {
	Ensure(context.Context, runtime.Object) (bool, error)
	Delete(context.Context, runtime.Object) (bool, error)
}

// AsyncReconciler is a generic reconciler for Azure resources.
// It reconciles object which require long running operations.
type AsyncReconciler struct {
	client.Client
	AzureClient AsyncClient
	Telemetry   telemetry.PrometheusTelemetry
	Recorder    record.EventRecorder
}

func (r *AsyncReconciler) Reconcile(req ctrl.Request, local runtime.Object) (result ctrl.Result, err error) {
	ctx := context.Background()

	// log operator start
	r.Telemetry.LogStart()

	// log failure / success
	defer func() {
		if err != nil {
			r.Telemetry.LogError(
				"Failure occured during reconcilliation",
				err)
			r.Telemetry.LogFailure()
		} else if result.Requeue {
			r.Telemetry.LogFailure()
		} else {
			r.Telemetry.LogSuccess()
		}
	}()

	if err := r.Get(ctx, req.NamespacedName, local); err != nil {
		r.Telemetry.LogInfo("ignorable error", "error during fetch from api server")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	res, convertErr := meta.Accessor(local)
	if convertErr != nil {
		r.Telemetry.LogInfo("ignorable error", "accessor fail")
		return ctrl.Result{}, convertErr
	}

	if res.GetDeletionTimestamp().IsZero() {

		if !HasFinalizer(res, finalizerName) {
			AddFinalizer(res, finalizerName)
			r.Recorder.Event(local, corev1.EventTypeNormal, "Added", "Object finalizer is added")
			return ctrl.Result{}, r.Update(ctx, local)
		}
	} else {
		if HasFinalizer(res, finalizerName) {
			found, deleteErr := r.AzureClient.Delete(ctx, local)
			final := multierror.Append(deleteErr, r.Status().Update(ctx, local))
			if err := final.ErrorOrNil(); err != nil {
				r.Telemetry.LogInfo("ignorable error", "error deleting object: "+err.Error())
				r.Recorder.Event(local, corev1.EventTypeWarning, "FailedDelete", fmt.Sprintf("Failed to delete resource: %s", err.Error()))
				return ctrl.Result{}, err
			}
			if !found {
				r.Recorder.Event(local, corev1.EventTypeNormal, "Deleted", "Successfully deleted")
				RemoveFinalizer(res, finalizerName)
				return ctrl.Result{}, r.Update(ctx, local)
			}
			r.Telemetry.LogInfo("requeuing", "deletion unfinished")
			return ctrl.Result{RequeueAfter: requeDuration}, nil
		}
		return ctrl.Result{}, nil
	}

	r.Telemetry.LogInfo("status", "reconciling object")
	done, ensureErr := r.AzureClient.Ensure(ctx, local)
	if ensureErr != nil {

		r.Telemetry.LogError("ensure err", ensureErr)
	}

	final := multierror.Append(ensureErr, r.Status().Update(ctx, local))
	err = final.ErrorOrNil()
	if err != nil {
		r.Recorder.Event(local, corev1.EventTypeWarning, "FailedReconcile", fmt.Sprintf("Failed to reconcile resource: %s", err.Error()))
	} else if done {
		r.Recorder.Event(local, corev1.EventTypeNormal, "Reconciled", "Successfully reconciled")
	}

	result = ctrl.Result{Requeue: !done}
	if !done {
		result.RequeueAfter = requeDuration
	}

	return result, err
}
