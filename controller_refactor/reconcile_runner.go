package controller_refactor

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
)

// Contains all the state involved in running a single reconcile event in the reconcile loo[
type reconcileRunner struct {
	*AzureController
	*ThisResourceDefinitions
	*DependencyDefinitions
	req          ctrl.Request
	requeueAfter time.Duration
	log          logr.Logger
}

//runs a single reconcile on the
func (r *reconcileRunner) run(ctx context.Context) (ctrl.Result, error) {

	// Verify that all dependencies are present in the cluster, and they are
	owner := r.Owner
	details := r.Details
	updater := r.Updater
	provisionState := details.ProvisionState
	allDeps := append([]*CustomResourceDetails{owner}, r.Dependencies...)

	// jump out and requeue if any of the dependencies are missing
	for _, dep := range allDeps {
		if dep != nil && !dep.ProvisionState.IsSucceeded() {
			r.log.Info("One of the dependencies is not in Succeeded state, requeuing")
			return ctrl.Result{Requeue: true, RequeueAfter: r.requeueAfter}, nil
		}
	}

	// set the owner reference if owner is present and references have not been set
	// currently we only have single object ownership, but it is poosible to have multiple owners
	if owner != nil && len(details.BaseDefinition.ObjectMeta.GetOwnerReferences()) == 0 {
		return r.setOwner(ctx, updater, owner, details)
	}

	// now verify the resource state on Azure
	if provisionState.IsVerifying() || provisionState.IsPending() {
		return r.verify(ctx)
	}

	// dependencies are now satisfied, can now reconcile the manfest and create or update the resource
	if details.ProvisionState.IsCreating() || details.ProvisionState.IsUpdating() {
		return r.ensure(ctx)
	}

	// reverify if in succeeded state
	if details.ProvisionState.IsSucceeded() && r.PostProvisionHandler != nil {
		result, err := r.verify(ctx)
		// if still succeeded run the post provision handler (set secrets in k8s etc)...
		if details.BaseDefinition.Status.ProvisionState.IsSucceeded() && err == nil {
			return r.runPostProvisionHandler()
		} else {
			return result, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *reconcileRunner) setOwner(ctx context.Context, updater *CustomResourceUpdater, owner *CustomResourceDetails, details *CustomResourceDetails) (ctrl.Result, error) {
	//set owner reference if it exists
	updater.SetOwnerReferences([]*CustomResourceDetails{owner})
	if err := r.updateAndLog(ctx, corev1.EventTypeNormal, "OwnerReferences", "Setting OwnerReferences for "+details.Name); err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

func (r *reconcileRunner) verify(ctx context.Context) (ctrl.Result, error) {
	updater := r.Updater
	details := r.Details
	requeueAfter := r.requeueAfter

	verifyResult, err := r.verifyExternal(ctx)
	if err != nil {
		// verification should not return an error - if this happens it's a terminal failure
		updater.SetProvisionState(azurev1alpha1.Failed)
		_ = r.updateAndLog(ctx, corev1.EventTypeWarning, "Verification", "Verification failed for "+details.Name)
		return ctrl.Result{}, fmt.Errorf("error verifying resource in azure: %v", err)
	}
	// Success case - the resource is is provisioned on Azure, and is ready
	if verifyResult.ready() {
		if !details.ProvisionState.IsSucceeded() {
			var ppError error = nil
			if r.PostProvisionHandler != nil {
				ppError = r.PostProvisionHandler(details)
			}
			if ppError != nil {
				updater.SetProvisionState(azurev1alpha1.Failed)
				_ = r.updateAndLog(ctx, corev1.EventTypeWarning, "PostProvisionHandler", "PostProvisionHandler failed to execute successfully for "+details.Name)
			} else {
				updater.SetProvisionState(azurev1alpha1.Succeeded)
				if err := r.updateAndLog(ctx, corev1.EventTypeNormal, "Succeeded", details.Name+" provisioned"); err != nil {
					return ctrl.Result{}, err
				}
			}
			return ctrl.Result{}, err
		}
	}
	// Missing case - we can now create the resource
	if verifyResult.missing() {
		updater.SetProvisionState(azurev1alpha1.Creating)
		err := r.updateAndLog(ctx, corev1.EventTypeNormal, "Creating", details.Name+" ready for creation")
		if err != nil {
			return ctrl.Result{}, err
		}
	}
	// Update case - the resource exists in Azure, is invalid but updateable, so doesn't need to be recreated
	if verifyResult.updateRequired() {
		updater.SetProvisionState(azurev1alpha1.Updating)

		if err := r.updateAndLog(ctx, corev1.EventTypeNormal, "Updating", details.Name+" flagged for update"); err != nil {
			return ctrl.Result{}, err
		}
	}
	// Recreate case - the resource exists in Azure, is invalid and needs to be created
	if verifyResult.recreateRequired() {
		deleteResult, err := r.ResourceManagerClient.Delete(ctx, details.Instance)
		if err != nil || deleteResult == DeleteError {
			updater.SetProvisionState(azurev1alpha1.Failed)
			return ctrl.Result{}, err
		}

		if deleteResult.alreadyDeleted() || deleteResult.succeed() {
			updater.SetProvisionState(azurev1alpha1.Creating)
		}

		// set it back to pending and let it go through the whole process again
		if deleteResult.awaitingVerification() {
			updater.SetProvisionState(azurev1alpha1.Pending)
		}

		if err := r.updateAndLog(ctx, corev1.EventTypeNormal, "Recreating", details.Name+" delete and recreate started"); err != nil {
			return ctrl.Result{}, err
		}
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, err
	}
	// if resource is deleting, requeue the reconcile loop
	if verifyResult.deleting() {
		r.log.Info("Retrying verification", "type", "Resource awaiting deletion before recreation can begin, requeuing reconcile loop")
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}
	// if still is in progress with provisioning or if it is busy deleting, requeue the reconcile loop
	if verifyResult.provisioning() {
		r.log.Info("Retrying verification", "type", "Verification not complete, requeuing reconcile loop")
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}
	return ctrl.Result{}, nil
}

func (r *reconcileRunner) verifyExternal(ctx context.Context) (VerifyResult, error) {
	instance := r.Details.Instance

	r.Recorder.Event(instance, corev1.EventTypeNormal, "Checking", "instance is ready")
	verifyResult, err := r.ResourceManagerClient.Verify(ctx, instance)

	if err != nil {
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Couldn't verify resource in azure")
		return verifyResult, errhelp.NewAzureError(err)
	}
	return verifyResult, nil
}

func (r *reconcileRunner) ensure(ctx context.Context) (ctrl.Result, error) {
	updater := r.Updater
	details := r.Details
	instance := details.Instance
	requeueAfter := r.requeueAfter

	r.Recorder.Event(details.Instance, corev1.EventTypeNormal, "Submitting", "starting resource reconciliation")
	// TODO: Add error handling for cases where username or password are invalid:
	// https://docs.microsoft.com/en-us/rest/api/sql/servers/createorupdate#response
	nextState, ensureErr := r.ensureExternal(ctx)
	// we set the state even if there is an error
	updater.SetProvisionState(nextState)
	updateErr := r.updateAndLog(ctx, corev1.EventTypeNormal, "Ensure", details.Name+" state set to "+string(nextState))
	if ensureErr != nil {
		return ctrl.Result{}, fmt.Errorf("error ensuring resource in azure: %v", ensureErr)
	}
	if updateErr != nil {
		return ctrl.Result{}, fmt.Errorf("error updating resource in azure: %v", updateErr)
	}
	if nextState.IsSucceeded() {
		r.Recorder.Event(instance, corev1.EventTypeNormal, "Updated",
			fmt.Sprintf("%s resource '%s' provisioned and ready.", details.BaseDefinition.Kind, details.Name))
		return ctrl.Result{}, nil
	}
	// give azure some time to catch up
	r.log.Info("waiting for provision to take effect")
	return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
}

func (r *reconcileRunner) ensureExternal(ctx context.Context) (azurev1alpha1.ProvisionState, error) {

	var err error

	details := r.Details
	resourceName := details.Name
	instance := details.Instance

	// ensure that the resource is created or updated in Azure (though it won't necessarily be ready, it still needs to be verified)
	var ensureResult EnsureResult
	if details.ProvisionState.IsCreating() {
		ensureResult, err = r.ResourceManagerClient.Create(ctx, instance)
	} else {
		ensureResult, err = r.ResourceManagerClient.Update(ctx, instance)
	}
	if err != nil || ensureResult.failed() || ensureResult.invalidRequest() {
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Couldn't create or update resource in azure")
		return azurev1alpha1.Failed, err
	}

	// if successful, set it to succeeded, or await verification
	var nextState azurev1alpha1.ProvisionState
	if ensureResult.awaitingVerification() {
		nextState = azurev1alpha1.Verifying
	} else if ensureResult.succeeded() {
		nextState = azurev1alpha1.Succeeded
	} else {
		return azurev1alpha1.Failed, errhelp.NewAzureError(fmt.Errorf("invalid response from Create for resource '%s'", resourceName))
	}
	return nextState, nil
}

func (r *reconcileRunner) addFinalizers(ctx context.Context) error {
	instance := r.Details.Instance
	updater := r.Updater

	updater.AddFinalizer(r.FinalizerName)
	updater.SetProvisionState(azurev1alpha1.Pending)
	if err := r.updateAndLog(ctx, corev1.EventTypeNormal, "Updated", "finalizers added"); err != nil {
		return err
	}
	r.Recorder.Event(instance, corev1.EventTypeNormal, "Updated", "finalizers added")
	return nil
}

func (r *reconcileRunner) handleFinalizer() (ctrl.Result, error) {
	details := r.Details
	instance := details.Instance
	updater := r.Updater
	ctx := context.Background()
	removeFinalizer := false
	requeue := false

	isTerminating := details.ProvisionState.IsTerminating()

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
			// it's going to be stuck if it ever gets to here, because we are not able to remove the finalizer.
			return ctrl.Result{Requeue: true, RequeueAfter: r.requeueAfter}, err
		}
		if !isTerminating {
			r.Recorder.Event(instance, corev1.EventTypeNormal, "Finalizer", "Setting state to terminating for "+details.Name)
		}
		if removeFinalizer {
			r.Recorder.Event(instance, corev1.EventTypeNormal, "Finalizer", "Removing finalizer for "+details.Name)
		}
	}

	if requeue {
		return ctrl.Result{Requeue: true, RequeueAfter: r.requeueAfter}, nil
	} else {
		r.Recorder.Event(instance, corev1.EventTypeNormal, "Finalizer", details.Name+" finalizer complete")
		return ctrl.Result{}, nil
	}
}

func (r *reconcileRunner) runPostProvisionHandler() (ctrl.Result, error) {
	if err := r.PostProvisionHandler(r.Details); err != nil {
		r.Log.Info("Error", "PostProvisionHandler", fmt.Sprintf("PostProvisionHandler failed: %s", err.Error()))
	}
	return ctrl.Result{}, nil
}

func (r *reconcileRunner) updateInstance(ctx context.Context) error {
	return r.tryUpdateInstance(ctx , 2)
}

// this is to get rid of the pesky errors
// "Operation cannot be fulfilled on xxx:the object has been modified; please apply your changes to the latest version and try again"
func (r *reconcileRunner) tryUpdateInstance(ctx context.Context, count int) error {
	// refetch the instance and apply the updates to it
	thisDefs, err := r.DefinitionManager.GetThis(ctx, r.req)
	if err != nil {
		if apierrors.IsNotFound(err) {
			r.Log.Info("Unable to update deleted resource. It may have already been finalized. This error is ignored. Resource: " + r.Details.Name)
			return nil
		} else {
			r.Log.Info("Unable to retrieve resource. Falling back to prior instance: "+r.Details.Name, "err", err.Error())
			thisDefs = r.ThisResourceDefinitions
		}
	}
	r.Updater.ApplyUpdates(thisDefs.Details.BaseDefinition)
	instance := thisDefs.Details.Instance
	err = r.KubeClient.Update(ctx, instance)
	if err != nil {
		if count == 0 {
			r.Recorder.Event(instance, corev1.EventTypeWarning, "Update", "failed to update CRD instance on K8s cluster.")
			r.Updater.Clear()
			return err
		}
		r.Log.Info(fmt.Sprintf("failed to update CRD instance on K8s cluster. Retries left=%d", count))
		time.Sleep(2 * time.Second)
		return r.tryUpdateInstance(ctx, count - 1)
	} else {
		r.Updater.Clear()
		return nil
	}
}

func (r *reconcileRunner) updateAndLog(ctx context.Context, eventType string, reason string, message string) error {
	instance := r.Details.Instance
	if err := r.updateInstance(ctx); err != nil {
		return err
	}
	r.Recorder.Event(instance, eventType, reason, message)
	return nil
}
