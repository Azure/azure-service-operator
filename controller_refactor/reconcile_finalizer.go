package controller_refactor

import (
	"context"
	"fmt"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *reconcileFinalizer) exists() bool {
	return r.Details.BaseDefinition.HasFinalizer(r.FinalizerName)
}

func (r *reconcileFinalizer) add(ctx context.Context) (ctrl.Result, error) {
	instance := r.Details.Instance
	updater := r.Updater

	updater.AddFinalizer(r.FinalizerName)
	updater.SetProvisionState(azurev1alpha1.Pending)
	if err := r.updateAndLog(ctx, corev1.EventTypeNormal, "Updated", "finalizers added"); err != nil {
		return ctrl.Result{}, fmt.Errorf("error adding finalizer: %v", err)
	}
	r.Recorder.Event(instance, corev1.EventTypeNormal, "Updated", "finalizers added")
	return ctrl.Result{}, nil
}

func (r *reconcileFinalizer) handle() (ctrl.Result, error) {
	details := r.Details
	instance := details.Instance
	updater := r.Updater
	ctx := context.Background()
	removeFinalizer := false
	requeue := false

	isTerminating := r.provisionState.IsTerminating()

	if details.BaseDefinition.HasFinalizer(r.FinalizerName) {
		// Even before we cal ResourceManagerClient.Delete, we verify the state of the resource
		// If it has not been created, we don't need to delete anything.
		verifyResult, err := r.ResourceManagerClient.Verify(ctx, instance)

		if verifyResult.error() || err != nil {
			// TODO: log error (this should not happen, but we carry on allowing the result to delete)
			// TODO: maybe should rather retry a certain number of times before failing
			removeFinalizer = true
		} else if verifyResult.missing() {
			removeFinalizer = true
		} else if verifyResult.deleting() {
			requeue = true
		} else if !isTerminating { // and one of verifyResult.ready() || verifyResult.recreateRequired() || verifyResult.updateRequired()
			// This block of code should only ever get called once.
			deleteResult, err := r.ResourceManagerClient.Delete(ctx, instance)
			if err != nil || deleteResult.error() {
				// TODO: log error (this should not happen, but we carry on allowing the result to delete)
				// Neither ResourceManagerClient.Verify nor ResourceManagerClient.Delete should error under usual conditions.
				// This is an unexpected error.
				removeFinalizer = true
			} else if deleteResult.alreadyDeleted() || deleteResult.succeed() {
				removeFinalizer = true
			} else if deleteResult.awaitingVerification() {
				requeue = true
			} else {
				// assert no more cases
				removeFinalizer = true
			}
		} else {
			// this should never be called, as the first time r.ResourceManagerClient.Delete is called isTerminating should be false
			// this implies that r.ResourceManagerClient.Delete didn't throw an error, but didn't do anything either
			removeFinalizer = true
		}
	}

	if !isTerminating {
		updater.SetProvisionState(azurev1alpha1.Terminating)
	}
	if removeFinalizer {
		updater.RemoveFinalizer(r.FinalizerName)
	}

	if removeFinalizer || !isTerminating {
		if err := r.updateInstance(ctx); err != nil {
			// if we can't update we have to requeue and hopefully it will remove the finalizer next time
			return ctrl.Result{Requeue: true, RequeueAfter: r.requeueAfter}, fmt.Errorf("error removing finalizer: %v", err)
		}
		if !isTerminating {
			r.Recorder.Event(instance, corev1.EventTypeNormal, "Finalizer", "setting state to terminating for " + r.Name)
		}
		if removeFinalizer {
			r.Recorder.Event(instance, corev1.EventTypeNormal, "Finalizer", "removing finalizer for " + r.Name)
		}
	}

	if requeue {
		return ctrl.Result{Requeue: true, RequeueAfter: r.requeueAfter}, nil
	} else {
		r.Recorder.Event(instance, corev1.EventTypeNormal, "Finalizer", r.Name+" finalizer complete")
		return ctrl.Result{}, nil
	}
}

