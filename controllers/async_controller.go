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
	successMsg    string        = "successfully provisioned"
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

// Reconcile reconciles the update issued to the operator
func (r *AsyncReconciler) Reconcile(req ctrl.Request, local runtime.Object) (result ctrl.Result, errRet error) {
	ctx := context.Background()

	if err := r.Get(ctx, req.NamespacedName, local); err != nil {
		r.Telemetry.LogWarning("ignorable error", "error during fetch from api server")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// log operator start
	r.Telemetry.LogStart()

	res, err := meta.Accessor(local)
	if err != nil {
		r.Telemetry.LogFailureWithError("failure occured during reconcilliation", err)
		return ctrl.Result{}, err
	}

	if res.GetDeletionTimestamp().IsZero() {
		if !HasFinalizer(res, finalizerName) {
			AddFinalizer(res, finalizerName)
			r.Recorder.Event(local, corev1.EventTypeNormal, "Added", "Object finalizer is added")
			r.Telemetry.LogSuccess()
			return ctrl.Result{}, r.Update(ctx, local)
		}
	} else {
		if HasFinalizer(res, finalizerName) {
			found, deleteErr := r.AzureClient.Delete(ctx, local)
			final := multierror.Append(deleteErr)
			if err := final.ErrorOrNil(); err != nil {
				r.Recorder.Event(local, corev1.EventTypeWarning, "FailedDelete", fmt.Sprintf("Failed to delete resource: %s", err.Error()))
				r.Telemetry.LogFailureWithError("failed to delete resource", err)
				return ctrl.Result{}, err
			}
			if !found {
				RemoveFinalizer(res, finalizerName)
				r.Recorder.Event(local, corev1.EventTypeNormal, "Deleted", "Successfully deleted")
				r.Telemetry.LogSuccessWithInfo("deleted", "successfully deleted")
				return ctrl.Result{}, r.Update(ctx, local)
			}
			r.Telemetry.LogFailure()
			return ctrl.Result{RequeueAfter: requeDuration}, nil
		}
		r.Telemetry.LogSuccess()
		return ctrl.Result{}, nil
	}

	// loop through parents until one is successfully referenced
	parents, err := r.AzureClient.GetParents(local)
	for _, p := range parents {
		//r.Telemetry.LogInfo("status", "handling parent "+p.Key.Name)

		if err := r.Get(ctx, p.Key, p.Target); err == nil {
			//r.Telemetry.LogInfo("status", "handling parent get for "+reflect.TypeOf(p.Target).String())

			if pAccessor, err := meta.Accessor(p.Target); err == nil {
				if err := controllerutil.SetControllerReference(pAccessor, res, r.Scheme); err == nil {
					r.Telemetry.LogInfo("status", "setting parent reference to object: "+pAccessor.GetName())
					err := r.Update(ctx, local)
					if err != nil {
						r.Telemetry.LogWarning("update", "failed to update parent instance: "+err.Error())
					}
					break
				}
			}
		}
	}

	done, ensureErr := r.AzureClient.Ensure(ctx, local)

	// update the status of the resource in kubernetes
	// implementations of Ensure() tend to set their outcomes in local.Status
	err = r.Status().Update(ctx, local)
	if err != nil {
		r.Telemetry.LogWarning("update", "failed updating status")
	}

	final := multierror.Append(ensureErr, r.Update(ctx, local))
	err = final.ErrorOrNil()
	result = ctrl.Result{}
	if done {
		if err != nil {

			// done, but an error occurred - make sure to log the error and mark as a failure,
			// 	however return nil as an error to end further reconcilliation
			r.Recorder.Event(local, corev1.EventTypeWarning, "FailedReconcile", fmt.Sprintf("Failed to reconcile resource: %s", err.Error()))
			r.Telemetry.LogFailureWithError("failed to reconcile resource", err)
			err = nil
		} else {

			// done, no error - success!
			r.Recorder.Event(local, corev1.EventTypeNormal, "Reconciled", "Successfully reconciled")
			r.Telemetry.LogSuccess()
		}
	} else {
		if err != nil {

			// not done but an error occurred - log the error and mark as a failure
			r.Telemetry.LogFailureWithError("failed to reconcile resource, requeueing", err)
		} else {

			// not done, but no error - mark as a failure
			r.Telemetry.LogFailure()
		}
		result.RequeueAfter = requeDuration
	}

	return result, err
}
