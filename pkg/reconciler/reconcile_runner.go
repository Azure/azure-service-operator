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

package reconciler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/go-logr/logr"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
)

const LastAppliedAnnotation = "/last-applied-spec"

// Contains all the state involved in running a single reconcile event in the reconcile loo[
type reconcileRunner struct {
	*GenericController
	*ResourceDefinition
	*DependencyDefinitions
	types.NamespacedName
	instance        runtime.Object
	objectMeta      metav1.Object
	status          *Status
	req             ctrl.Request
	log             logr.Logger
	instanceUpdater *instanceUpdater
	owner           runtime.Object
	dependencies    map[types.NamespacedName]runtime.Object
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
	r.dependencies = map[types.NamespacedName]runtime.Object{}

	// jump out and requeue if any of the dependencies are missing
	for i, dep := range allDeps {
		instance := dep.InitialInstance
		err := r.KubeClient.Get(ctx, dep.NamespacedName, instance)
		log := r.log.WithValues("Dependency", dep.NamespacedName)

		// if any of the dependencies are not found, we jump out.
		if err != nil { // note that dependencies should be an empty array
			if apierrors.IsNotFound(err) {
				log.Info("Dependency not found for " + dep.NamespacedName.Name + ". Requeuing request.")
			} else {
				log.Info(fmt.Sprintf("Unable to retrieve dependency for %s: %v", dep.NamespacedName.Name, err.Error()))
			}
			return r.applyTransition(ctx, "Dependency", Pending, client.IgnoreNotFound(err))
		}

		// set the owner reference if owner is present and references have not been set
		// currently we only have single object ownership, but it is poosible to have multiple owners
		if owner != nil && i == 0 {
			if len(r.objectMeta.GetOwnerReferences()) == 0 {
				return r.setOwner(ctx, instance)
			}
			r.owner = instance
		}
		r.dependencies[dep.NamespacedName] = instance

		succeeded, err := dep.SucceededAccessor(instance)
		if err != nil {
			log.Info(fmt.Sprintf("Cannot get success state for %s. terminal failure.", dep.NamespacedName.Name))
			// Fail if cannot get status accessor for dependency
			return r.applyTransition(ctx, "Dependency", Failed, err)
		}

		if !succeeded {
			log.Info("One of the dependencies is not in 'Succeeded' state, requeuing")
			return r.applyTransition(ctx, "Dependency", Pending, nil)
		}
	}

	// **** Pending
	// **** Verifying
	// **** Succeeded
	// **** Recreating
	// **** Failed
	// Verify the resource state on Azure
	if status.IsVerifying() || status.IsPending() || status.IsSucceeded() || status.IsRecreating() || status.IsFailed() {
		return r.verify(ctx)
	}

	// **** Creating
	// **** Updating
	// dependencies are now satisfied, can now reconcile the manifest and create or update the resource
	if status.IsCreating() || status.IsUpdating() {
		return r.apply(ctx)
	}

	// **** Completing
	// created or updated successful, final completion step
	if status.IsCompleting() {
		return r.runCompletion(ctx)
	}

	// **** Terminating
	if status.IsCompleting() {
		r.log.Info("unexpected condition. Terminating state should be handled in finalizer")
		return ctrl.Result{}, nil
	}

	// if has no status, set to pending
	return r.applyTransition(ctx, "run", Pending, nil)
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
	nextState, ensureErr := r.verifyExecute(ctx)
	return r.applyTransition(ctx, "Verify", nextState, ensureErr)
}

const rejectCreateManagedResource = "permission to create Azure resource is not set. Annotation '*/access-permissions' is present, but the flag 'C' is not set"
const rejectUpdateManagedResource = "permission to update Azure resource is not set. Annotation '*/access-permissions' is present, but the flag 'U' is not set"
const rejectDeleteManagedResource = "permission to delete or recreate Azure resource is not set. Annotation '*/access-permissions' is present, but the flag 'D' is not set"

func (r *reconcileRunner) verifyExecute(ctx context.Context) (ReconcileState, error) {
	status := r.status
	currentState := status.State

	r.log.Info("Verifying state of resource on Azure")
	verifyResponse, err := r.ResourceManager.Verify(ctx, r.resourceSpec())
	verifyResult := verifyResponse.Result
	permissions := r.getAccessPermissions()

	// **** Error
	// either an error was returned from the SDK, or the error failed
	if err != nil {
		return Failed, err
	}

	// **** Deleting
	// if the Azure resource is any state where it exists, but the K8s resource is recreating
	// we assume that the resource is being deleted asynchronously in Azure, but there is no way to distinguish
	// from the SDK that is deleting - it is either present or not. so we requeue the loop and wait for it to become `missing`
	if verifyResult.deleting() || verifyResult.exists() && status.IsRecreating() {
		r.log.Info("Retrying verification: resource awaiting deletion before recreation can begin, requeuing reconcile loop")
		return currentState, nil
	}

	// **** Ready
	// The resource is provisioned on Azure, set to succeeded or conduct post provisioning completion step
	if verifyResult.ready() {
		// set the status payload if there is any
		r.instanceUpdater.setStatusPayload(verifyResponse.Status)
		return r.succeedOrCompletion(), nil
	}

	// **** Missing
	// We can now create the resource
	if verifyResult.missing() {
		if status.IsFailed() && r.Parameters.RequeueAfterFailure == 0 {
			// if not requeing failure, leave as failed
			return Failed, nil
		}
		if !permissions.create() {
			// fail if permission to create is not present
			return Failed, fmt.Errorf(rejectCreateManagedResource)
		}
		return Creating, nil
	}

	// **** In progress
	// if still is in progress with provisioning requeue the reconcile loop
	if verifyResult.inProgress() {
		r.log.Info("Retrying verification: provisioning in progress, requeuing reconcile loop")
		return currentState, nil
	}

	// **** UpdateRequired
	// The resource exists in Azure, is invalid but updateable, so doesn't need to be recreated
	if verifyResult.updateRequired() {
		if !permissions.update() {
			// fail if permission to update is not present
			return Failed, fmt.Errorf(rejectUpdateManagedResource)
		}
		return Updating, nil
	}

	// **** RecreateRequired
	// The resource exists in Azure, is invalid and needs to be created
	if verifyResult.recreateRequired() {
		if !permissions.delete() {
			// fail if permission to delete is not present
			return Failed, fmt.Errorf(rejectDeleteManagedResource)
		}
		deleteResult, err := r.ResourceManager.Delete(ctx, r.resourceSpec())
		if err != nil || deleteResult == DeleteError {
			return Failed, err
		}

		// set it back to pending and let it go through the whole process again
		if deleteResult.awaitingVerification() {
			return Recreating, err
		}

		if deleteResult.alreadyDeleted() || deleteResult.succeeded() {
			return Creating, err
		}

		return Failed, fmt.Errorf("invalid DeleteResult for %s %s in Verify", r.ResourceKind, r.Name)
	}

	// **** Error
	return Failed, fmt.Errorf("invalid VerifyResult for %s %s in Verify, and no error was specified", r.ResourceKind, r.Name)
}

func (r *reconcileRunner) apply(ctx context.Context) (ctrl.Result, error) {
	r.log.Info("Ready to create or update resource in Azure")
	nextState, ensureErr := r.applyExecute(ctx)
	return r.applyTransition(ctx, "Ensure", nextState, ensureErr)
}

func (r *reconcileRunner) applyExecute(ctx context.Context) (ReconcileState, error) {

	resourceName := r.Name
	instance := r.instance
	status := r.status
	lastAppliedAnnotation := r.AnnotationBaseName + LastAppliedAnnotation
	permissions := r.getAccessPermissions()

	// apply that the resource is created or updated in Azure (though it won't necessarily be ready, it still needs to be verified)
	var err error
	var applyResponse ApplyResponse
	if status.IsCreating() {
		if !permissions.create() {
			// this should never be the case - this is more of an assertion (as the state Verify or Create should never have been set in the first place)
			return Failed, fmt.Errorf(rejectCreateManagedResource)
		}
		applyResponse, err = r.ResourceManager.Create(ctx, r.resourceSpec())
	} else {
		if !permissions.update() {
			// this should never be the case - this is more of an assertion (as the state Verify or Create should never have been set in the first place)
			return Failed, fmt.Errorf(rejectCreateManagedResource)
		}
		applyResponse, err = r.ResourceManager.Update(ctx, r.resourceSpec())
	}
	applyResult := applyResponse.Result
	if applyResult == "" || err != nil || applyResult.failed() {
		// clear last update annotation
		r.instanceUpdater.setAnnotation(lastAppliedAnnotation, "")
		errMsg := "Undetermined error"
		if err != nil {
			errMsg = err.Error()
		}
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", fmt.Sprintf("Couldn't create or update resource in azure: %v", errMsg))
		return Failed, err
	}

	// if successful
	// save the last updated spec as a metadata annotation
	r.instanceUpdater.setAnnotation(lastAppliedAnnotation, r.getJsonSpec())

	// set it to succeeded, conduct the completion step (if there is a Completing handler), or await verification
	if applyResult.awaitingVerification() {
		r.instanceUpdater.setStatusPayload(applyResponse.Status)
		return Verifying, nil
	} else if applyResult.succeeded() {
		r.instanceUpdater.setStatusPayload(applyResponse.Status)
		return r.succeedOrCompletion(), nil
	} else {
		return Failed, fmt.Errorf("invalid response from Create for resource '%s'", resourceName)
	}
}

func (r *reconcileRunner) succeedOrCompletion() ReconcileState {
	if r.CompletionRunner == nil || r.status.IsSucceeded() {
		return Succeeded
	} else {
		return Completing
	}
}

func (r *reconcileRunner) runCompletion(ctx context.Context) (ctrl.Result, error) {
	var ppError error = nil
	if r.CompletionRunner != nil {
		if handler := r.CompletionRunner(r.GenericController); handler != nil {
			ppError = handler.Run(ctx, r.instance)
		}
	}
	if ppError != nil {
		return r.applyTransition(ctx, "Completion", Failed, ppError)
	} else {
		return r.applyTransition(ctx, "Completion", Succeeded, nil)
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
			r.log.Info("Unable to update deleted resource. it may have already been finalized. this error is ignorable. Resource: " + r.Name)
			return nil
		} else {
			r.log.Info("Unable to retrieve resource. falling back to prior instance: " + r.Name + ": err " + err.Error())
		}
	}
	status, _ := r.StatusAccessor(instance)
	err = r.instanceUpdater.applyUpdates(instance, status)
	if err != nil {
		r.log.Info("Unable to convert Object to resource")
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
		r.log.Info(fmt.Sprintf("Failed to update CRD instance on K8s cluster. retries left=%d", count))
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
		return nil
	}
	if err := r.updateInstance(ctx); err != nil {
		r.log.Info(fmt.Sprintf("K8s update failure: %v", err))
		r.Recorder.Event(instance, corev1.EventTypeWarning, reason, fmt.Sprintf("failed to update instance of %s %s in kubernetes cluster", r.ResourceKind, r.Name))
		return err
	}
	r.Recorder.Event(instance, eventType, reason, message)
	return nil
}

func (r *reconcileRunner) getTransitionDetails(nextState ReconcileState) (ctrl.Result, string) {
	requeueAfter := r.getRequeueAfter(nextState)
	requeueResult := ctrl.Result{Requeue: requeueAfter > 0, RequeueAfter: requeueAfter}
	message := ""
	switch nextState {
	case Pending:
		message = fmt.Sprintf("%s %s in pending state.", r.ResourceKind, r.Name)
	case Creating:
		message = fmt.Sprintf("%s %s ready for creation.", r.ResourceKind, r.Name)
	case Updating:
		message = fmt.Sprintf("%s %s ready to be updated.", r.ResourceKind, r.Name)
	case Verifying:
		message = fmt.Sprintf("%s %s verification in progress.", r.ResourceKind, r.Name)
	case Completing:
		message = fmt.Sprintf("%s %s provisioning succeeded and ready for completion step.", r.ResourceKind, r.Name)
	case Succeeded:
		message = fmt.Sprintf("%s %s successfully provisioned and ready for use.", r.ResourceKind, r.Name)
	case Recreating:
		message = fmt.Sprintf("%s %s deleting and recreating in progress.", r.ResourceKind, r.Name)
	case Failed:
		message = fmt.Sprintf("%s %s failed.", r.ResourceKind, r.Name)
	case Terminating:
		message = fmt.Sprintf("%s %s termination in progress.", r.ResourceKind, r.Name)
	default:
		message = fmt.Sprintf("%s %s set to state %s", r.ResourceKind, r.Name, nextState)
	}
	return requeueResult, message
}

func (r *reconcileRunner) applyTransition(ctx context.Context, reason string, nextState ReconcileState, transitionErr error) (ctrl.Result, error) {
	eventType := corev1.EventTypeNormal
	if nextState == Failed {
		eventType = corev1.EventTypeWarning
	}
	errorMsg := ""
	if transitionErr != nil {
		errorMsg = transitionErr.Error()
	}
	if nextState != r.status.State {
		r.instanceUpdater.setReconcileState(nextState, errorMsg)
	}
	result, transitionMsg := r.getTransitionDetails(nextState)
	updateErr := r.updateAndLog(ctx, eventType, reason, transitionMsg)
	if transitionErr != nil {
		if updateErr != nil {
			// TODO: is the transition error is more important?
			// we don't requeue if there is an update error
			return ctrl.Result{}, transitionErr
		} else {
			return result, nil
		}
	}
	if updateErr != nil {
		return ctrl.Result{}, updateErr
	}
	return result, nil
}

func (r *reconcileRunner) getRequeueAfter(transitionState ReconcileState) time.Duration {
	parameters := r.Parameters
	requeueAfterDuration := func(requeueSeconds int) time.Duration {
		requeueAfter := time.Duration(requeueSeconds) * time.Second
		return requeueAfter
	}

	if transitionState == Pending ||
		transitionState == Verifying ||
		transitionState == Recreating {
		// must by default have a non zero requeue for these states
		requeueSeconds := parameters.RequeueAfter
		if requeueSeconds == 0 {
			requeueSeconds = 30
		}
		return requeueAfterDuration(requeueSeconds)
	} else if transitionState == Failed {
		return requeueAfterDuration(parameters.RequeueAfterFailure)
	} else if transitionState == Succeeded {
		return requeueAfterDuration(parameters.RequeueAfterSuccess)
	}
	return 0
}

func (r *reconcileRunner) getJsonSpec() string {
	fetch := func() (string, error) {
		b, err := json.Marshal(r.instance)
		if err != nil {
			return "", err
		}
		var asMap map[string]interface{}
		err = json.Unmarshal(b, &asMap)
		if err != nil {
			return "", err
		}
		spec := asMap["spec"]
		b, err = json.Marshal(spec)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}

	jsonSpec, err := fetch()
	if err != nil {
		r.log.Info("Error fetching Json for instance spec")
		return ""
	}
	return jsonSpec
}

func (r *reconcileRunner) resourceSpec() ResourceSpec {
	return ResourceSpec{Instance: r.instance, Dependencies: r.dependencies}
}

func (r *reconcileRunner) getAccessPermissions() AccessPermissions {
	annotations := r.objectMeta.GetAnnotations()
	return AccessPermissions(annotations[r.AnnotationBaseName+AccessPermissionAnnotation])
}
