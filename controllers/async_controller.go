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
	"github.com/go-logr/logr"
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
	Log         logr.Logger
}

func (r *AsyncReconciler) Reconcile(req ctrl.Request, local runtime.Object) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("instance", req.String())

	if err := r.Get(ctx, req.NamespacedName, local); err != nil {
		log.Info("error during fetch from api server")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	status, err := r.AzureClient.GetStatus(local)
	if err != nil {
		return ctrl.Result{}, err
	}

	// // log operator start
	r.Telemetry.LogStart(status)

	res, err := meta.Accessor(local)
	if err != nil {
		log.Info("failed getting meta accessor", "err", err)

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
				log.Info("error deleting object", "err", err)

				r.Recorder.Event(local, corev1.EventTypeWarning, "FailedDelete", fmt.Sprintf("Failed to delete resource: %s", err.Error()))
				return ctrl.Result{}, err
			}
			if !found {
				r.Recorder.Event(local, corev1.EventTypeNormal, "Deleted", "Successfully deleted")
				RemoveFinalizer(res, finalizerName)
				return ctrl.Result{}, r.Update(ctx, local)
			}
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
					log.Info("setting parent reference", "target", pAccessor.GetName())
					err := r.Update(ctx, local)
					if err != nil {
						log.Info("failed to reference parent", "err", err)
					}
					break
				}
			}
		}
	}

	log.Info("reconciling object")
	done, ensureErr := r.AzureClient.Ensure(ctx, local)
	if ensureErr != nil {
		log.Info("error from Ensure", "err", err)
	}

	if done && ensureErr == nil {
		r.Telemetry.LogSuccess(status)
	}

	// update the status of the resource in kubernetes
	// Implementations of Ensure() tend to set their outcomes in local.Status
	err = r.Status().Update(ctx, local)
	if err != nil {
		log.Info("failed to update status", "err", err)
	}

	final := multierror.Append(ensureErr, r.Update(ctx, local))
	err = final.ErrorOrNil()
	if err != nil {
		r.Recorder.Event(local, corev1.EventTypeWarning, "FailedReconcile", fmt.Sprintf("Unexpected error during reconciliation: %s", err.Error()))
	} else if done {
		r.Recorder.Event(local, corev1.EventTypeNormal, "Reconciled", "Reconciliation ended with no errors")
	}

	result := ctrl.Result{}
	if !done {
		log.Info("reconciling object not finished, re-queueing")
		result.RequeueAfter = requeDuration
	} else {
		log.Info("reconciling object complete")
	}

	log.Info("message from operator", "message", status.Message)
	log.Info("exiting reconciliation")

	return result, err
}
