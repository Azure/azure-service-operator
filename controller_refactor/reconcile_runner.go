package controller_refactor

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
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
	*GenericController
	*ThisResourceDefinitions
	*DependencyDefinitions
	types.NamespacedName
	objectMeta      metav1.Object
	typeMeta        metav1.Type
	provisionState  azurev1alpha1.ProvisionState
	req             ctrl.Request
	requeueAfter    time.Duration
	log             logr.Logger
	instanceUpdater *customResourceUpdater
}

type reconcileFinalizer struct {
	reconcileRunner
}

//runs a single reconcile on the
func (r *reconcileRunner) run(ctx context.Context) (ctrl.Result, error) {

	// Verify that all dependencies are present in the cluster, and they are
	owner := r.Owner
	allDeps := append([]*CustomResourceDetails{owner}, r.Dependencies...)
	provisionState := r.Details.Status.ProvisionState

	// jump out and requeue if any of the dependencies are missing
	for _, dep := range allDeps {
		if dep != nil && !dep.Status.ProvisionState.IsSucceeded() {
			r.log.Info("one of the dependencies is not in 'Succeeded' state, requeuing")
			return ctrl.Result{Requeue: true, RequeueAfter: r.requeueAfter}, nil
		}
	}

	// set the owner reference if owner is present and references have not been set
	// currently we only have single object ownership, but it is poosible to have multiple owners
	if owner != nil && len(r.objectMeta.GetOwnerReferences()) == 0 {
		return r.setOwner(ctx, owner)
	}

	// now verify the resource state on Azure
	if provisionState.IsVerifying() || provisionState.IsPending() {
		return r.verify(ctx)
	}

	// dependencies are now satisfied, can now reconcile the manfest and create or update the resource
	if provisionState.IsCreating() || provisionState.IsUpdating() {
		return r.ensure(ctx)
	}

	// re-verify if in succeeded state
	if provisionState.IsSucceeded() && r.PostProvisionHandler != nil {
		result, err := r.verify(ctx)
		// if still succeeded run the post provision handler (set secrets in k8s etc)...
		if r.provisionState.IsSucceeded() && err == nil {
			return r.runPostProvisionHandler()
		} else {
			return result, err
		}
	}

	return ctrl.Result{}, nil
}

func (r *reconcileRunner) setOwner(ctx context.Context, owner *CustomResourceDetails) (ctrl.Result, error) {
	//set owner reference if it exists
	r.instanceUpdater.setOwnerReferences([]*CustomResourceDetails{owner})
	if err := r.updateAndLog(ctx, corev1.EventTypeNormal, "OwnerReferences", "setting OwnerReferences for "+r.Name); err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

func (r *reconcileRunner) verify(ctx context.Context) (ctrl.Result, error) {
	updater := r.instanceUpdater
	details := r.Details
	requeueAfter := r.requeueAfter

	verifyResult, err := r.verifyExternal(ctx)
	if err != nil {
		// verification should not return an error - if this happens it's a terminal failure
		r.instanceUpdater.setProvisionState(azurev1alpha1.Failed)
		_ = r.updateAndLog(ctx, corev1.EventTypeWarning, "Verification", "verification failed for "+r.Name)
		return ctrl.Result{}, fmt.Errorf("error verifying resource in azure: %v", err)
	}
	// Success case - the resource is is provisioned on Azure, and is ready
	if verifyResult.ready() {
		if !r.provisionState.IsSucceeded() {
			var ppError error = nil
			if r.PostProvisionHandler != nil {
				ppError = r.PostProvisionHandler(details)
			}
			if ppError != nil {
				updater.setProvisionState(azurev1alpha1.Failed)
				_ = r.updateAndLog(ctx, corev1.EventTypeWarning, "PostProvisionHandler", "PostProvisionHandler failed to execute successfully for "+r.Name)
			} else {
				updater.setProvisionState(azurev1alpha1.Succeeded)
				if err := r.updateAndLog(ctx, corev1.EventTypeNormal, "Succeeded", r.Name+" provisioned"); err != nil {
					return ctrl.Result{}, err
				}
			}
			return ctrl.Result{}, err
		}
	}
	// Missing case - we can now create the resource
	if verifyResult.missing() {
		updater.setProvisionState(azurev1alpha1.Creating)
		err := r.updateAndLog(ctx, corev1.EventTypeNormal, "Creating", r.Name+" ready for creation")
		if err != nil {
			return ctrl.Result{}, err
		}
	}
	// Update case - the resource exists in Azure, is invalid but updateable, so doesn't need to be recreated
	if verifyResult.updateRequired() {
		updater.setProvisionState(azurev1alpha1.Updating)

		if err := r.updateAndLog(ctx, corev1.EventTypeNormal, "Updating", r.Name+" flagged for update"); err != nil {
			return ctrl.Result{}, err
		}
	}
	// Recreate case - the resource exists in Azure, is invalid and needs to be created
	if verifyResult.recreateRequired() {
		deleteResult, err := r.ResourceManagerClient.Delete(ctx, details.Instance)
		if err != nil || deleteResult == DeleteError {
			updater.setProvisionState(azurev1alpha1.Failed)
			return ctrl.Result{}, err
		}

		if deleteResult.alreadyDeleted() || deleteResult.succeed() {
			updater.setProvisionState(azurev1alpha1.Creating)
		}

		// set it back to pending and let it go through the whole process again
		if deleteResult.awaitingVerification() {
			updater.setProvisionState(azurev1alpha1.Pending)
		}

		if err := r.updateAndLog(ctx, corev1.EventTypeNormal, "Recreating", r.Name+" delete and recreate started"); err != nil {
			return ctrl.Result{}, err
		}
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, err
	}
	// if resource is deleting, requeue the reconcile loop
	if verifyResult.deleting() {
		r.log.Info("Retrying verification", "type", "resource awaiting deletion before recreation can begin, requeuing reconcile loop")
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}
	// if still is in progress with provisioning or if it is busy deleting, requeue the reconcile loop
	if verifyResult.provisioning() {
		r.log.Info("Retrying verification", "type", "verification not complete, requeuing reconcile loop")
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}
	return ctrl.Result{}, nil
}

func (r *reconcileRunner) verifyExternal(ctx context.Context) (VerifyResult, error) {
	instance := r.Details.Instance

	r.Recorder.Event(instance, corev1.EventTypeNormal, "Checking", "instance is ready")
	verifyResult, err := r.ResourceManagerClient.Verify(ctx, instance)

	if err != nil {
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "couldn't verify resource in azure")
		return verifyResult, errhelp.NewAzureError(err)
	}
	return verifyResult, nil
}

func (r *reconcileRunner) ensure(ctx context.Context) (ctrl.Result, error) {
	updater := r.instanceUpdater
	details := r.Details
	instance := details.Instance
	requeueAfter := r.requeueAfter

	// TODO: keep this line ?
	r.Recorder.Event(details.Instance, corev1.EventTypeNormal, "Ensure", "ready to create or update resource")
	nextState, ensureErr := r.ensureExternal(ctx)
	// we set the state even if there is an error
	updater.setProvisionState(nextState)
	updateErr := r.updateAndLog(ctx, corev1.EventTypeNormal, "Ensure", r.Name+" state set to "+string(nextState))
	if ensureErr != nil {
		return ctrl.Result{}, fmt.Errorf("error ensuring resource in azure: %v", ensureErr)
	}
	if updateErr != nil {
		return ctrl.Result{}, fmt.Errorf("error updating resource in azure: %v", updateErr)
	}
	if nextState.IsSucceeded() {
		r.Recorder.Event(instance, corev1.EventTypeNormal, "Updated",
			fmt.Sprintf("%s resource '%s' provisioned and ready.", r.ResourceKind, r.Name))
		return ctrl.Result{}, nil
	}
	// give azure some time to catch up
	r.log.Info("waiting for provision to take effect")
	return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
}

func (r *reconcileRunner) ensureExternal(ctx context.Context) (azurev1alpha1.ProvisionState, error) {

	var err error

	details := r.Details
	resourceName := r.Name
	instance := details.Instance

	// ensure that the resource is created or updated in Azure (though it won't necessarily be ready, it still needs to be verified)
	var ensureResult EnsureResult
	if r.provisionState.IsCreating() {
		ensureResult, err = r.ResourceManagerClient.Create(ctx, instance)
	} else {
		ensureResult, err = r.ResourceManagerClient.Update(ctx, instance)
	}
	if err != nil || ensureResult.failed() || ensureResult.invalidRequest() {
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "couldn't create or update resource in azure")
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

func (r *reconcileRunner) runPostProvisionHandler() (ctrl.Result, error) {
	if err := r.PostProvisionHandler(r.Details); err != nil {
		r.Log.Info("Error", "PostProvisionHandler", fmt.Sprintf("PostProvisionHandler failed: %s", err.Error()))
	}
	return ctrl.Result{}, nil
}

func (r *reconcileRunner) updateInstance(ctx context.Context) error {
	return r.tryUpdateInstance(ctx, 2)
}

// this is to get rid of the pesky errors
// "Operation cannot be fulfilled on xxx:the object has been modified; please apply your changes to the latest version and try again"
func (r *reconcileRunner) tryUpdateInstance(ctx context.Context, count int) error {
	// refetch the instance and apply the updates to it
	thisDefs, err := r.DefinitionManager.GetThis(ctx, r.req)
	if err != nil {
		if apierrors.IsNotFound(err) {
			r.Log.Info("unable to update deleted resource. it may have already been finalized. this error is ignorable. Resource: " + r.Name)
			return nil
		} else {
			r.Log.Info("unable to retrieve resource. falling back to prior instance: "+r.Name, "err", err.Error())
			thisDefs = r.ThisResourceDefinitions
		}
	}
	instance := thisDefs.Details.Instance
	err = r.instanceUpdater.applyUpdates(thisDefs.Details)
	if err != nil {
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Update", "unable to convert Object to resource.")
		r.instanceUpdater.clear()
		return err
	}
	err = r.KubeClient.Update(ctx, instance)
	if err != nil {
		if count == 0 {
			r.Recorder.Event(instance, corev1.EventTypeWarning, "Update", "failed to update CRD instance on K8s cluster.")
			r.instanceUpdater.clear()
			return err
		}
		r.Log.Info(fmt.Sprintf("failed to update CRD instance on K8s cluster. retries left=%d", count))
		time.Sleep(2 * time.Second)
		return r.tryUpdateInstance(ctx, count-1)
	} else {
		r.instanceUpdater.clear()
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
