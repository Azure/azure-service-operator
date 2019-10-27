/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
	log             logr.Logger
	instanceUpdater *instanceUpdater
}

type reconcileFinalizer struct {
	reconcileRunner
}

//runs a single reconcile on the
func (r *reconcileRunner) run(ctx context.Context) (ctrl.Result, error) {

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
		pendingResult := ctrl.Result{Requeue: true, RequeueAfter: r.getRequeueAfter(azurev1alpha1.Pending)}

		// if any of the dependencies are not found, we jump out.
		if err != nil { // note that dependencies should be an empty array
			if apierrors.IsNotFound(err) {
				r.logInfo("Dependency not found for " + dep.NamespacedName.Name + ". Requeuing request.")
			} else {
				r.logInfo(fmt.Sprintf("Unable to retrieve dependency for %s: %v", dep.NamespacedName.Name, err.Error()))
			}
			return pendingResult, client.IgnoreNotFound(err)
		}

		// set the owner reference if owner is present and references have not been set
		// currently we only have single object ownership, but it is poosible to have multiple owners
		if i == 0 && owner != nil && len(r.objectMeta.GetOwnerReferences()) == 0 {
			return r.setOwner(ctx, instance)
		}

		status, err := dep.StatusAccessor(instance)
		if err != nil {
			r.logInfo(fmt.Sprintf("Cannot get status for %s. terminal failure.", dep.NamespacedName.Name))
			// TODO fail - this is terminal
			return ctrl.Result{}, err
		}

		if !status.IsSucceeded() {
			r.logInfo("One of the dependencies is not in 'Succeeded' state, requeuing")
			return pendingResult, nil
		}
	}

	// now verify the resource state on Azure
	if status.IsVerifying() || status.IsPending() || status.IsSucceeded() || status.IsRecreating() {
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

	// if has no status, set to pending
	if status.IsPostProvisioning() {
		return r.applyTransition(ctx, "run", azurev1alpha1.Pending, nil)
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
	r.logInfo("Verifying state of resource in Azure")
	nextState, ensureErr := r.verifyExecute(ctx)
	return r.applyTransition(ctx, "Verify", nextState, ensureErr)
}

func (r *reconcileRunner) verifyExecute(ctx context.Context) (azurev1alpha1.ProvisionState, error) {
	status := r.status
	instance := r.instance
	currentState := status.ProvisionState()

	r.logInfo("Checking state of resource on Azure")
	verifyResult, err := r.ResourceManagerClient.Verify(ctx, instance)

	if err != nil {
		// verification should not return an error - if this happens it's a terminal failure
		return azurev1alpha1.Failed, err
	}
	// Success case - the resource is provisioned on Azure, post provisioning can take place if necessary
	if verifyResult.ready() {
		// if the Azure resource is in ready state, but the K8s resource is recreating
		// we assume that the resource is being deleted asynchronously in Azure, but there is no way to distinguish
		// from the SDK that is deleting - it is either present or not. so we requeue the loop and wait for it to become `missing`
		if status.IsRecreating() {
			r.logInfo("Retrying verification: resource awaiting deletion before recreation can begin, requeuing reconcile loop")
			return currentState, nil
		}
		return r.succeedOrPostProvision(), nil
	}
	// Missing case - we can now create the resource
	if verifyResult.missing() {
		return azurev1alpha1.Creating, nil
	}
	// if resource is deleting, requeue the reconcile loop
	if verifyResult.deleting() {
		r.logInfo("Retrying verification: resource awaiting deletion before recreation can begin, requeuing reconcile loop")
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
			return azurev1alpha1.Recreating, err
		}

		if deleteResult.alreadyDeleted() || deleteResult.succeed() {
			return azurev1alpha1.Creating, err
		}

		return azurev1alpha1.Failed, errhelp.NewAzureError(fmt.Errorf("Invalid DeleteResult for %s %s in Verify", r.ResourceKind, r.Name))
	}

	return azurev1alpha1.Failed, errhelp.NewAzureError(fmt.Errorf("Invalid VerifyResult for %s %s in Verify", r.ResourceKind, r.Name))
}

func (r *reconcileRunner) ensure(ctx context.Context) (ctrl.Result, error) {
	r.logInfo("Ready to create or update resource in Azure")
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
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Couldn't create or update resource in azure")
		return azurev1alpha1.Failed, err
	}

	// if successful, set it to succeeded, post provisioning (if there is a PostProvisioning handler), or await verification
	if ensureResult.awaitingVerification() {
		return azurev1alpha1.Verifying, nil
	} else if ensureResult.succeeded() {
		return r.succeedOrPostProvision(), nil
	} else {
		return azurev1alpha1.Failed, errhelp.NewAzureError(fmt.Errorf("Invalid response from Create for resource '%s'", resourceName))
	}
}

func (r *reconcileRunner) succeedOrPostProvision() azurev1alpha1.ProvisionState {
	if r.PostProvisionFactory == nil || r.status.IsSucceeded() {
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
	return r.tryUpdateInstance(ctx, 5)
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
			r.logInfo("Unable to update deleted resource. it may have already been finalized. this error is ignorable. Resource: " + r.Name)
			return nil
		} else {
			r.logInfo("Unable to retrieve resource. falling back to prior instance: " + r.Name + ": err " + err.Error())
		}
	}
	status, _ := r.StatusAccessor(instance)
	err = r.instanceUpdater.applyUpdates(instance, status)
	if err != nil {
		r.Log.Info("Unable to convert Object to resource")
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
		r.logInfo(fmt.Sprintf("Failed to update CRD instance on K8s cluster. retries left=%d", count))
		time.Sleep(2 * time.Second)
		return r.tryUpdateInstance(ctx, count-1)
	} else {
		r.instanceUpdater.clear()
		return nil
	}
}

func (r *reconcileRunner) updateAndLog(ctx context.Context, eventType string, reason string, message string) error {
	instance := r.instance
	if !r.instanceUpdater.hasUpdates() {
		r.logInfo(fmt.Sprintf("Resource up to date. no further update necessary."))
		return nil
	}
	if err := r.updateInstance(ctx); err != nil {
		r.logInfo(fmt.Sprintf("K8s update failure: %v", err))
		r.Recorder.Event(instance, corev1.EventTypeWarning, reason, fmt.Sprintf("failed to update instance of %s %s in kubernetes cluster", r.ResourceKind, r.Name))
		return err
	}
	r.Recorder.Event(instance, eventType, reason, message)
	return nil
}

func (r *reconcileRunner) getTransitionDetails(nextState azurev1alpha1.ProvisionState) (ctrl.Result, string) {
	requeueAfter := r.getRequeueAfter(nextState)
	requeueResult := ctrl.Result{Requeue: requeueAfter > 0, RequeueAfter: requeueAfter}
	message := ""
	switch nextState {
	case azurev1alpha1.Pending:
		message = fmt.Sprintf("%s %s in pending state.", r.ResourceKind, r.Name)
	case azurev1alpha1.Creating:
		message = fmt.Sprintf("%s %s ready for creation.", r.ResourceKind, r.Name)
	case azurev1alpha1.Updating:
		message = fmt.Sprintf("%s %s ready to be updated.", r.ResourceKind, r.Name)
	case azurev1alpha1.Verifying:
		message = fmt.Sprintf("%s %s verification in progress.", r.ResourceKind, r.Name)
	case azurev1alpha1.PostProvisioning:
		message = fmt.Sprintf("%s %s provisioning succeeded and ready for post-provisioning step", r.ResourceKind, r.Name)
	case azurev1alpha1.Succeeded:
		message = fmt.Sprintf("%s %s successfully provisioned and ready for use.", r.ResourceKind, r.Name)
	case azurev1alpha1.Recreating:
		message = fmt.Sprintf("%s %s deleting and recreating in progress.", r.ResourceKind, r.Name)
	case azurev1alpha1.Failed:
		message = fmt.Sprintf("%s %s failed.", r.ResourceKind, r.Name)
	case azurev1alpha1.Terminating:
		message = fmt.Sprintf("%s %s termination in progress.", r.ResourceKind, r.Name)
	default:
		message = fmt.Sprintf("%s %s set to state %s", r.ResourceKind, r.Name, nextState)
	}
	return requeueResult, message
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
	r.log.Info(message, "Kind", r.ResourceKind)
}

func (r *reconcileRunner) getRequeueAfter(transitionState azurev1alpha1.ProvisionState) time.Duration {
	parameters := r.Parameters
	requeueAfterDuration := func(requeueSeconds int) time.Duration {
		requeueAfter := time.Duration(requeueSeconds) * time.Second
		return requeueAfter
	}

	if transitionState == azurev1alpha1.Pending ||
		transitionState == azurev1alpha1.Verifying ||
		transitionState == azurev1alpha1.Recreating {
		// must by default have a non zero requeue for these states
		requeueSeconds := parameters.RequeueAfter
		if requeueSeconds == 0 {
			requeueSeconds = 10
		}
		return requeueAfterDuration(requeueSeconds)
	} else if transitionState == azurev1alpha1.Failed {
		return requeueAfterDuration(parameters.RequeueAfterFailure)
	} else if transitionState == azurev1alpha1.Succeeded {
		return requeueAfterDuration(parameters.RequeueAfterSuccess)
	}
	return 0
}
