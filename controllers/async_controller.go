/*
Copyright 2019 Alexander Eldeib.
*/

package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
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
	ForSubscription(context.Context, runtime.Object) error
	Ensure(context.Context, runtime.Object) (bool, error)
	Delete(context.Context, runtime.Object) (bool, error)
}

// AsyncReconciler is a generic reconciler for Azure resources.
// It reconciles object which require long running operations.
type AsyncReconciler struct {
	client.Client
	Az       AsyncClient
	Log      logr.Logger
	Recorder record.EventRecorder
}

func (r *AsyncReconciler) Reconcile(req ctrl.Request, local runtime.Object) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("type", local.GetObjectKind().GroupVersionKind().String(), "namespace", req.Namespace, "name", req.Name)

	if err := r.Get(ctx, req.NamespacedName, local); err != nil {
		log.Info("error during fetch from api server")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// if err := r.Az.ForSubscription(ctx, local); err != nil {
	// 	return ctrl.Result{}, err
	// }

	res, convertErr := meta.Accessor(local)
	if convertErr != nil {
		log.Info("accessor fail")
		return ctrl.Result{}, convertErr
	}

	if res.GetDeletionTimestamp().IsZero() {

		if !HasFinalizer(res, finalizerName) {
			AddFinalizer(res, finalizerName)
			r.Recorder.Event(local, "Normal", "Added", "Object finalizer is added")
			return ctrl.Result{}, r.Update(ctx, local)
		}
	} else {
		if HasFinalizer(res, finalizerName) {
			found, deleteErr := r.Az.Delete(ctx, local)
			final := multierror.Append(deleteErr, r.Status().Update(ctx, local))
			if err := final.ErrorOrNil(); err != nil {
				log.Info("error deleting object", "error", err.Error())
				r.Recorder.Event(local, "Warning", "FailedDelete", fmt.Sprintf("Failed to delete resource: %s", err.Error()))
				return ctrl.Result{}, err
			}
			if !found {
				r.Recorder.Event(local, corev1.EventTypeNormal, "Deleted", "Successfully deleted")
				RemoveFinalizer(res, finalizerName)
				return ctrl.Result{}, r.Update(ctx, local)
			}
			log.Info("requeuing, deletion unfinished")
			return ctrl.Result{RequeueAfter: requeDuration}, nil
		}
		return ctrl.Result{}, nil
	}

	log.Info("reconciling object")
	done, ensureErr := r.Az.Ensure(ctx, local)
	if ensureErr != nil {

		log.Error(ensureErr, "ensure err")
	}

	final := multierror.Append(ensureErr, r.Update(ctx, local))
	err := final.ErrorOrNil()
	if err != nil {
		r.Recorder.Event(local, "Warning", "FailedReconcile", fmt.Sprintf("Failed to reconcile resource: %s", err.Error()))
	} else if done {
		r.Recorder.Event(local, "Normal", "Reconciled", "Successfully reconciled")
	}

	result := ctrl.Result{Requeue: !done}
	if !done {
		result.RequeueAfter = requeDuration
	}

	return result, err
}
