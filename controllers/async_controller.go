/*
Copyright 2019 Alexander Eldeib.
*/

package controllers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	telemetry "github.com/Azure/azure-service-operator/pkg/telemetry"
	multierror "github.com/hashicorp/go-multierror"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	Telemetry   telemetry.TelemetryClient
	Recorder    record.EventRecorder
	Scheme      *runtime.Scheme
}

// Reconcile reconciles the change request
func (r *AsyncReconciler) Reconcile(req ctrl.Request, local runtime.Object) (result ctrl.Result, err error) {
	ctx := context.Background()
	r.Telemetry.SetInstance(req.String())

	if err := r.Get(ctx, req.NamespacedName, local); err != nil {
		r.Telemetry.LogInfo("requeueing", "error during fetch from api server")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// get the ASOStatus struct
	status, err := r.AzureClient.GetStatus(local)
	if err != nil {
		return ctrl.Result{}, err
	}

	// record the time that this request was requested at
	if !status.Provisioning && !status.Provisioned {
		status.Provisioning = true
		timeNow := metav1.NewTime(time.Now())
		status.RequestedAt = &timeNow
	}

	res, err := meta.Accessor(local)
	if err != nil {
		r.Telemetry.LogInfo("requeuing", fmt.Sprintf("failed getting meta accessor: %s", err.Error()))
		return ctrl.Result{}, err
	}

	// Check to see if the skipreconcile annotation is on
	var skipReconcile bool
	annotations := res.GetAnnotations()
	if val, ok := annotations["skipreconcile"]; ok {
		if strings.ToLower(val) == "true" {
			skipReconcile = true
		}
	}

	if skipReconcile {
		// if this is a delete we should delete the finalizer to allow the kube instance to be deleted
		if !res.GetDeletionTimestamp().IsZero() {
			if HasFinalizer(res, finalizerName) {
				RemoveFinalizer(res, finalizerName)
			}
		}
		r.Recorder.Event(local, corev1.EventTypeNormal, "Skipping", "Skipping reconcile based on provided annotation")
		return ctrl.Result{}, r.Update(ctx, local)
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
			final := multierror.Append(deleteErr)
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
	parents, err := r.AzureClient.GetParents(local)
	for _, p := range parents {
		if err := r.Get(ctx, p.Key, p.Target); err == nil {
			if pAccessor, err := meta.Accessor(p.Target); err == nil {
				if err := controllerutil.SetControllerReference(pAccessor, res, r.Scheme); err == nil {
					r.Telemetry.LogInfo("setting parent reference", pAccessor.GetName())
					err := r.Update(ctx, local)
					if err != nil {
						r.Telemetry.LogError("failed to reference parent", err)
					}
					break
				}
			}
		}
	}

	r.Telemetry.LogTrace("reconciling", "reconciling object")
	done, ensureErr := r.AzureClient.Ensure(ctx, local)
	if ensureErr != nil {
		r.Telemetry.LogError("error from Ensure", err)
	}

	// update the status of the resource in kubernetes
	// Implementations of Ensure() tend to set their outcomes in local.Status
	err = r.Status().Update(ctx, local)
	if err != nil {
		r.Telemetry.LogError("failed to update status", err)
	}

	final := multierror.Append(ensureErr, r.Update(ctx, local))
	err = final.ErrorOrNil()
	if err != nil {
		r.Recorder.Event(local, corev1.EventTypeWarning, "FailedReconcile", fmt.Sprintf("Unexpected error during reconciliation: %s", err.Error()))
	} else if done {
		r.Recorder.Event(local, corev1.EventTypeNormal, "Reconciled", "Reconciliation ended with no errors")
	}

	result = ctrl.Result{}
	if !done {
		r.Telemetry.LogInfo("requeueing", "reconciling object not finished, re-queueing")
		result.RequeueAfter = requeDuration
	} else {
		r.Telemetry.LogInfo("reconciling", "success")

		// record the duration of the request
		if status.CompletedAt == nil || status.CompletedAt.IsZero() {
			compTime := metav1.Now()
			status.CompletedAt = &compTime
			durationInSecs := status.CompletedAt.Sub(status.RequestedAt.Time).Seconds()
			r.Telemetry.LogDuration(durationInSecs)
		}
	}

	r.Telemetry.LogInfo("operator", fmt.Sprintf("message from operator: %s", status.Message))
	r.Telemetry.LogTrace("operator", "exiting reconciliation")

	return result, err
}
