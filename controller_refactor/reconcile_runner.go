package controller_refactor

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
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
	*ResourceDefinition
	*DependencyDefinitions
	types.NamespacedName
	instance        runtime.Object
	objectMeta      metav1.Object
	status          *azurev1alpha1.ASOStatus
	req             ctrl.Request
	requeueAfter    time.Duration
	log             logr.Logger
	instanceUpdater *instanceUpdater
}

type reconcileFinalizer struct {
	reconcileRunner
}

//runs a single reconcile on the
func (r *reconcileRunner) run(ctx context.Context) (ctrl.Result, error) {
	result, err := r.runExecute(ctx)
	//updateErr := r.updateInstance(ctx)
	//if updateErr != nil {
	//	return ctrl.Result{}, err
	//}
	return result, err
}

func (r *reconcileRunner) runExecute(ctx context.Context) (ctrl.Result, error) {

	// Verify that all dependencies are present in the cluster, and they are
	owner := r.Owner
	var allDeps []*Dependency
	if owner != nil {
		allDeps = append([]*Dependency{owner}, r.Dependencies...)
	} else {
		allDeps = r.Dependencies
	}
	status := r.status

	// jump out and requeue if any of the dependencies are missing
	for i, dep := range allDeps {
		instance := dep.InitialInstance
		err := r.KubeClient.Get(ctx, dep.NamespacedName, instance)

		// if any of the dependencies are not found, we jump out.
		if err != nil { // note that dependencies should be an empty array
			if apierrors.IsNotFound(err) {
				r.logInfo("dependency not found for " + dep.NamespacedName.Name + ". requeuing request.")
			} else {
				r.logInfo(fmt.Sprintf("unable to retrieve dependency for %s: %v", dep.NamespacedName.Name, err.Error()))
			}
			return ctrl.Result{Requeue: true, RequeueAfter: r.requeueAfter}, client.IgnoreNotFound(err)
		}

		// set the owner reference if owner is present and references have not been set
		// currently we only have single object ownership, but it is poosible to have multiple owners
		if i == 0 && owner != nil && len(r.objectMeta.GetOwnerReferences()) == 0 {
			return r.setOwner(ctx, instance)
		}

		status, err := dep.StatusAccessor(instance)
		if err != nil {
			r.logInfo(fmt.Sprintf("cannot get status for %s. terminal failure.", dep.NamespacedName.Name))
			// TODO fail - this is terminal
			return ctrl.Result{Requeue: true, RequeueAfter: r.requeueAfter}, nil
		}

		if !status.IsSucceeded() {
			r.logInfo("one of the dependencies is not in 'Succeeded' state, requeuing")
			return ctrl.Result{Requeue: true, RequeueAfter: r.requeueAfter}, nil
		}
	}

	// now verify the resource state on Azure
	if status.IsVerifying() || status.IsPending() || status.IsSucceeded() {
		return r.verify(ctx)
	}

	// dependencies are now satisfied, can now reconcile the manfest and create or update the resource
	if status.IsCreating() || status.IsUpdating() {
		return r.ensure(ctx)
	}

	// has created or updated, post provisioning
	if status.IsPostProvisioning() {
		return r.runPostProvision(ctx)
	}

	return ctrl.Result{}, nil
}

func (r *reconcileRunner) setOwner(ctx context.Context, owner runtime.Object) (ctrl.Result, error) {
	//set owner reference if it exists
	r.instanceUpdater.setOwnerReferences([]runtime.Object{owner})
	if err := r.updateAndLog(ctx, corev1.EventTypeNormal, "OwnerReferences", "setting OwnerReferences for "+r.Name); err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

func (r *reconcileRunner) verify(ctx context.Context) (ctrl.Result, error) {
	r.logInfo("verifying state of resource in Azure")
	nextState, ensureErr := r.verifyExecute(ctx)
	return r.applyTransition(ctx, "Verify", nextState, ensureErr)
}

func (r *reconcileRunner) verifyExecute(ctx context.Context) (azurev1alpha1.ProvisionState, error) {
	status := r.status
	instance := r.instance
	currentState := status.ProvisionState()

	r.logInfo("checking state of resource on Azure")
	verifyResult, err := r.ResourceManagerClient.Verify(ctx, instance)

	if err != nil {
		// verification should not return an error - if this happens it's a terminal failure
		return azurev1alpha1.Failed, err
	}
	// Success case - the resource is provisioned on Azure, post provisioning can take place if necessary
	if verifyResult.ready() {
		return r.succeedOrPostProvision(), nil
	}
	// Missing case - we can now create the resource
	if verifyResult.missing() {
		return azurev1alpha1.Creating, nil
	}
	// if resource is deleting, requeue the reconcile loop
	if verifyResult.deleting() {
		r.logInfo("retrying verification: resource awaiting deletion before recreation can begin, requeuing reconcile loop")
		return currentState, nil
	}
	// if still is in progress with provisioning or if it is busy deleting, requeue the reconcile loop
	if verifyResult.provisioning() {
		r.logInfo("Retrying verification: verification of provisioning not complete, requeuing reconcile loop")
		return currentState, nil
	}

	// Update case - the resource exists in Azure, is invalid but updateable, so doesn't need to be recreated
	if verifyResult.updateRequired() {
		return azurev1alpha1.Updating, nil
	}
	// Recreate case - the resource exists in Azure, is invalid and needs to be created
	if verifyResult.recreateRequired() {
		deleteResult, err := r.ResourceManagerClient.Delete(ctx, instance)
		if err != nil || deleteResult == DeleteError {
			// TODO: add log here
			return azurev1alpha1.Failed, err
		}

		// set it back to pending and let it go through the whole process again
		if deleteResult.awaitingVerification() {
			// TODO: make azurev1alpha1.Recreating
			return azurev1alpha1.Pending, err
		}

		if deleteResult.alreadyDeleted() || deleteResult.succeed() {
			return azurev1alpha1.Creating, err
		}

		return azurev1alpha1.Failed, errhelp.NewAzureError(fmt.Errorf("invalid DeleteResult for %s %s in Verify", r.ResourceKind, r.Name))
	}

	return azurev1alpha1.Failed, errhelp.NewAzureError(fmt.Errorf("invalid VerifyResult for %s %s in Verify", r.ResourceKind, r.Name))
}

func (r *reconcileRunner) ensure(ctx context.Context) (ctrl.Result, error) {
	r.logInfo("ready to create or update resource in Azure")
	nextState, ensureErr := r.ensureExecute(ctx)
	return r.applyTransition(ctx, "Ensure", nextState, ensureErr)
}

func (r *reconcileRunner) ensureExecute(ctx context.Context) (azurev1alpha1.ProvisionState, error) {

	resourceName := r.Name
	instance := r.instance
	status := r.status

	// ensure that the resource is created or updated in Azure (though it won't necessarily be ready, it still needs to be verified)
	var err error
	var ensureResult EnsureResult
	if status.IsCreating() {
		ensureResult, err = r.ResourceManagerClient.Create(ctx, instance)
	} else {
		ensureResult, err = r.ResourceManagerClient.Update(ctx, instance)
	}
	if err != nil || ensureResult.failed() || ensureResult.invalidRequest() {
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "couldn't create or update resource in azure")
		return azurev1alpha1.Failed, err
	}

	// if successful, set it to succeeded, post provisioning (if there is a PostProvisioning handler), or await verification
	var nextState azurev1alpha1.ProvisionState
	if ensureResult.awaitingVerification() {
		nextState = azurev1alpha1.Verifying
	} else if ensureResult.succeeded() {
		if r.PostProvisionFactory == nil {
			nextState = azurev1alpha1.Succeeded
		} else {
			nextState = azurev1alpha1.PostProvisioning
		}
	} else {
		return azurev1alpha1.Failed, errhelp.NewAzureError(fmt.Errorf("invalid response from Create for resource '%s'", resourceName))
	}
	return nextState, nil
}

func (r *reconcileRunner) succeedOrPostProvision() azurev1alpha1.ProvisionState {
	if r.PostProvisionFactory == nil {
		return azurev1alpha1.Succeeded
	} else {
		return azurev1alpha1.PostProvisioning
	}
}

func (r *reconcileRunner) runPostProvision(ctx context.Context) (ctrl.Result, error) {
	var ppError error = nil
	if r.PostProvisionFactory != nil {
		if handler := r.PostProvisionFactory(r.GenericController); handler != nil {
			ppError = handler.Run(ctx, r.instance)
		}
	}
	if ppError != nil {
		return r.applyTransition(ctx, "PostProvision", azurev1alpha1.Failed, ppError)
	} else {
		return r.applyTransition(ctx, "PostProvision", azurev1alpha1.Succeeded, nil)
	}
}

func (r *reconcileRunner) updateInstance(ctx context.Context) error {
	if !r.instanceUpdater.hasUpdates() {
		return nil
	}
	return r.tryUpdateInstance(ctx, 2)
}

// this is to get rid of the pesky errors
// "Operation cannot be fulfilled on xxx:the object has been modified; please apply your changes to the latest version and try again"
func (r *reconcileRunner) tryUpdateInstance(ctx context.Context, count int) error {
	// refetch the instance and apply the updates to it
	baseInstance := r.instance
	instance := baseInstance.DeepCopyObject()
	err := r.KubeClient.Get(ctx, r.NamespacedName, baseInstance)
	if err != nil {
		if apierrors.IsNotFound(err) {
			r.logInfo("unable to update deleted resource. it may have already been finalized. this error is ignorable. Resource: " + r.Name)
			return nil
		} else {
			r.logInfo("unable to retrieve resource. falling back to prior instance: " + r.Name + ": err " + err.Error())
		}
	}
	status, _ := r.StatusAccessor(instance)
	err = r.instanceUpdater.applyUpdates(instance, status)
	if err != nil {
		r.Log.Info("unable to convert Object to resource")
		r.instanceUpdater.clear()
		return err
	}
	err = r.KubeClient.Update(ctx, instance)
	if err != nil {
		if count == 0 {
			r.Recorder.Event(instance, corev1.EventTypeWarning, "Update", fmt.Sprintf("failed to update  %s instance %s on K8s cluster.", r.ResourceKind, r.Name))
			r.instanceUpdater.clear()
			return err
		}
		r.logInfo(fmt.Sprintf("failed to update CRD instance on K8s cluster. retries left=%d", count))
		time.Sleep(2 * time.Second)
		return r.tryUpdateInstance(ctx, count-1)
	} else {
		r.instanceUpdater.clear()
		return nil
	}
}

func (r *reconcileRunner) updateAndLog(ctx context.Context, eventType string, reason string, message string) error {
	instance := r.instance
	if err := r.updateInstance(ctx); err != nil {
		r.logInfo(fmt.Sprintf("K8s update failure: %v", err))
		r.Recorder.Event(instance, corev1.EventTypeWarning, reason, fmt.Sprintf("failed to update instance of %s %s in kubernetes cluster", r.ResourceKind, r.Name))
		return err
	}
	r.Recorder.Event(instance, eventType, reason, message)
	return nil
}

func (r *reconcileRunner) getTransitionDetails(nextState azurev1alpha1.ProvisionState) (ctrl.Result, string) {
	requeueResult := ctrl.Result{Requeue: true, RequeueAfter: r.requeueAfter}
	switch nextState {
	case azurev1alpha1.Pending:
		return ctrl.Result{}, fmt.Sprintf("%s %s in pending state.", r.ResourceKind, r.Name)
	case azurev1alpha1.Creating:
		return ctrl.Result{}, fmt.Sprintf("%s %s ready for creation.", r.ResourceKind, r.Name)
	case azurev1alpha1.Updating:
		return ctrl.Result{}, fmt.Sprintf("%s %s ready to be updated.", r.ResourceKind, r.Name)
	case azurev1alpha1.Verifying:
		return requeueResult, fmt.Sprintf("%s %s verification in progress.", r.ResourceKind, r.Name)
	case azurev1alpha1.Succeeded:
		return ctrl.Result{}, fmt.Sprintf("%s %s successfully provisioned and ready for use.", r.ResourceKind, r.Name)
	case azurev1alpha1.PostProvisioning:
		return ctrl.Result{}, fmt.Sprintf("%s %s provisioning succeeded and ready for post-provisioning step", r.ResourceKind, r.Name)
	case azurev1alpha1.Failed:
		return ctrl.Result{}, fmt.Sprintf("%s %s failed.", r.ResourceKind, r.Name)
	case azurev1alpha1.Terminating:
		return ctrl.Result{}, fmt.Sprintf("%s %s termination in progress.", r.ResourceKind, r.Name)
	}
	return ctrl.Result{}, fmt.Sprintf("%s %s set to state %s", r.ResourceKind, r.Name, nextState)
}

func (r *reconcileRunner) applyTransition(ctx context.Context, reason string, nextState azurev1alpha1.ProvisionState, transitionErr error) (ctrl.Result, error) {
	eventType := corev1.EventTypeNormal
	if nextState == azurev1alpha1.Failed {
		eventType = corev1.EventTypeWarning
	}
	if nextState != r.status.ProvisionState() {
		r.instanceUpdater.setProvisionState(nextState)
	}
	result, message := r.getTransitionDetails(nextState)
	updateErr := r.updateAndLog(ctx, eventType, reason, message)
	if transitionErr != nil {
		if updateErr != nil {
			// TODO: is the transition error is more important?
			// we don't requeue if there is an update error
			return ctrl.Result{}, transitionErr
		} else {
			return result, updateErr
		}
	}
	if updateErr != nil {
		return ctrl.Result{}, updateErr
	}
	return result, nil
}

func (r *reconcileRunner) logInfo(message string) {
	r.log.Info(message, "Kind", r.ResourceKind, "Name", r.Name)
}
