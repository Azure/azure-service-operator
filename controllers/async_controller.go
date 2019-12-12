/*
Copyright 2019 Alexander Eldeib.
*/

package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	telemetry "github.com/Azure/azure-service-operator/pkg/telemetry"
	multierror "github.com/hashicorp/go-multierror"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

const (
	finalizerName string        = "azure.microsoft.com/finalizer"
	requeDuration time.Duration = time.Second * 20
)

// AsyncReconciler is a generic reconciler for Azure resources.
// It reconciles Kubernets objects which require long running operations in Azure.
type AsyncReconciler struct {
	client.Client
	AzureClient resourcemanager.ARMClient
	Telemetry   telemetry.PrometheusTelemetry
	Recorder    record.EventRecorder
	Scheme      *runtime.Scheme
}

func (r *AsyncReconciler) Reconcile(req ctrl.Request, local runtime.Object) (result ctrl.Result, err error) {
	ctx := context.Background()

  if err := r.Get(ctx, req.NamespacedName, local); err != nil {
		r.Telemetry.LogInfo("ignorable error", "error during fetch from api server")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

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

	res, err := meta.Accessor(local)
	if err != nil {
		r.Telemetry.LogError("accessor fail", err)
		return ctrl.Result{}, err
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
				r.Telemetry.LogError("error deleting object", err)
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

	// loop through parents until one is successfully referenced
	r.Telemetry.LogInfo("status", "handling parent reference for object")

	parents, err := r.AzureClient.GetParents(local)
	for _, p := range parents {
		//r.Telemetry.LogInfo("status", "handling parent "+p.Key.Name)

		if err := r.Get(ctx, p.Key, p.Target); err == nil {
			//r.Telemetry.LogInfo("status", "handling parent get for "+reflect.TypeOf(p.Target).String())

			if pAccessor, err := meta.Accessor(p.Target); err == nil {
				if err := controllerutil.SetControllerReference(pAccessor, res, r.Scheme); err == nil {
					//r.Telemetry.LogInfo("status", "handling parent reference for object "+pAccessor.GetName())
					return ctrl.Result{}, r.Update(ctx, local)

				}
			}
		}
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

	r.Telemetry.LogInfo("status", "exiting reconciliation")

	return result, err
}
