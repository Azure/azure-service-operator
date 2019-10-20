package controller_refactor

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"time"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/go-logr/logr"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// AzureController reconciles a ResourceGroup object
type AzureController struct {
	ResourceKind          string
	KubeClient            client.Client
	Log                   logr.Logger
	Recorder              record.EventRecorder
	ResourceManagerClient ResourceManagerClient
	DefinitionManager     DefinitionManager
	FinalizerName         string
	PostProvisionHandler  PostProvisionHandler
}

// SetupWithManager function sets up the functions with the controller
func (r *AzureController) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.ResourceGroup{}).
		Complete(r)
}

// Reconcile function does the main reconciliation loop of the operator
func (r *AzureController) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("resourcegroup", req.NamespacedName)

	// fetch the manifest object
	thisDefs, err := r.DefinitionManager.GetThis(ctx, req)
	if err != nil {
		log.Info("Unable to retrieve resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	details := thisDefs.Details
	updater := thisDefs.Updater

	//
	if details.IsBeingDeleted {
		result, err := r.handleFinalizer(details, updater)
		if err != nil {
			return result, fmt.Errorf("error when handling finalizer: %v", err)
		}
		return result, nil
	}

	if len(details.BaseDefinition.Finalizers) == 0 {
		err := r.addFinalizers(details, updater)
		if err != nil {
			return ctrl.Result{}, fmt.Errorf("error adding finalizer: %v", err)
		}
		return ctrl.Result{}, nil
	}

	// verify status of dependencies
	dependencyInfo, err := r.DefinitionManager.GetDependencies(ctx, req)
	if err != nil {
		log.Info("Unable to retrieve resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// get parameters
	parameters := details.Parameters
	requeueSeconds := parameters.RequeueAfterSeconds
	requeueAfter := r.getRequeueAfter(requeueSeconds)

	// Verify that all dependencies are present in the cluster, and they are
	owner := dependencyInfo.Owner
	allDeps := append([]*CustomResourceDetails{owner}, dependencyInfo.Dependencies...)

	for _, dep := range allDeps {
		if dep != nil && !dep.ProvisionState.IsSucceeded() {
			log.Info("One of the dependencies is not in Succeeded state, requeuing")
			return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
		}
	}

	instance := details.Instance
	// set the owner reference if owner is present and references have not been set
	// currently we only have single object ownership, but it is poosible to have multiple owners
	if owner != nil && len(details.BaseDefinition.ObjectMeta.GetOwnerReferences()) == 0 {
		//set owner reference if it exists
		updater.SetOwnerReferences([]*CustomResourceDetails{owner})
		if err := r.updateAndLog(ctx, instance, corev1.EventTypeNormal, "OwnerReferences", "Setting OwnerReferences for "+details.Name); err != nil {
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, nil
	}

	// now verify the resource state on Azure
	if details.ProvisionState.IsVerifying() || details.ProvisionState.IsPending() || details.ProvisionState.IsSucceeded() {
		verifyResult, err := r.verifyExternal(details, updater)
		if err != nil {
			// verification should not return an error - if this happens it's a terminal failure
			updater.SetProvisionState(azurev1alpha1.Failed)
			_ = r.updateAndLog(ctx, instance, corev1.EventTypeWarning, "Verification", "Verification failed for "+details.Name)
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
					_ = r.updateAndLog(ctx, instance, corev1.EventTypeWarning, "PostProvisionHandler", "PostProvisionHandler failed to execute successfully for "+details.Name)
				} else {
					updater.SetProvisionState(azurev1alpha1.Succeeded)
					if err := r.updateAndLog(ctx, instance, corev1.EventTypeNormal, "Succeeded", details.Name+" provisioned"); err != nil {
						return ctrl.Result{}, err
					}
				}
				return ctrl.Result{}, err
			}
		}

		// Missing case - we can now create the resource
		if verifyResult.missing() {
			updater.SetProvisionState(azurev1alpha1.Creating)
			err := r.updateAndLog(ctx, instance, corev1.EventTypeNormal, "Creating", details.Name+" ready for creation")
			if err != nil {
				return ctrl.Result{}, err
			}
		}

		// Update case - the resource exists in Azure, is invalid but updateable, so doesn't need to be recreated
		if verifyResult.updateRequired() {
			updater.SetProvisionState(azurev1alpha1.Updating)

			if err := r.updateAndLog(ctx, instance, corev1.EventTypeNormal, "Updating", details.Name+" flagged for update"); err != nil {
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

			if err := r.updateAndLog(ctx, instance, corev1.EventTypeNormal, "Recreating", details.Name+" delete and recreate started"); err != nil {
				return ctrl.Result{}, err
			}
			return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, err
		}

		// if resource is deleting, requeue the reconcile loop
		if verifyResult.deleting() {
			log.Info("Retrying verification", "type", "Resource awaiting deletion before recreation can begin, requeuing reconcile loop")
			return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
		}

		// if still is in progress with provisioning or if it is busy deleting, requeue the reconcile loop
		if verifyResult.provisioning() {
			log.Info("Retrying verification", "type", "Verification not complete, requeuing reconcile loop")
			return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
		}

		return ctrl.Result{}, nil
	}

	// dependencies are now satisfied, can now reconcile the manfest and create or update the resource
	if details.ProvisionState.IsCreating() || details.ProvisionState.IsUpdating() {
		r.Recorder.Event(details.Instance, corev1.EventTypeNormal, "Submitting", "starting resource reconciliation")
		// TODO: Add error handling for cases where username or password are invalid:
		// https://docs.microsoft.com/en-us/rest/api/sql/servers/createorupdate#response

		nextState, ensureErr := r.ensureExternal(details, updater)

		// we set the state even if there is an error
		updater.SetProvisionState(nextState)
		updateErr := r.updateAndLog(ctx, instance, corev1.EventTypeNormal, "Ensure", details.Name+" state set to "+string(nextState))

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
		log.Info("waiting for provision to take effect")
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}

	if details.ProvisionState.IsSucceeded() && r.PostProvisionHandler != nil {
		if err := r.PostProvisionHandler(details); err != nil {
			r.Log.Info("Error", "PostProvisionHandler", fmt.Sprintf("PostProvisionHandler failed: %s", err.Error()))
		}
	}

	return ctrl.Result{}, nil
}

func (r *AzureController) ensureExternal(details *CustomResourceDetails, updater *CustomResourceUpdater) (azurev1alpha1.ProvisionState, error) {

	ctx := context.Background()
	var err error

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

func (r *AzureController) verifyExternal(details *CustomResourceDetails, updater *CustomResourceUpdater) (VerifyResult, error) {
	ctx := context.Background()
	instance := details.Instance

	r.Recorder.Event(instance, corev1.EventTypeNormal, "Checking", "instance is ready")
	verifyResult, err := r.ResourceManagerClient.Verify(ctx, instance)

	if err != nil {
		r.Recorder.Event(details.Instance, corev1.EventTypeWarning, "Failed", "Couldn't verify resource in azure")
		return verifyResult, errhelp.NewAzureError(err)
	}
	return verifyResult, nil
}

func (r *AzureController) addFinalizers(details *CustomResourceDetails, updater *CustomResourceUpdater) error {
	updater.AddFinalizer(r.FinalizerName)
	updater.SetProvisionState(azurev1alpha1.Pending)
	if err := r.updateAndLog(context.Background(), details.Instance, corev1.EventTypeNormal, "Updated", "finalizers added"); err != nil {
		return err
	}
	r.Recorder.Event(details.Instance, corev1.EventTypeNormal, "Updated", "finalizers added")
	return nil
}

func (r *AzureController) handleFinalizer(details *CustomResourceDetails, updater *CustomResourceUpdater) (ctrl.Result, error) {
	instance := details.Instance
	ctx := context.Background()
	removeFinalizer := false
	requeue := false
	requeueAfter := r.getRequeueAfter(details.Parameters.RequeueAfterSeconds)
	isTerminating := details.ProvisionState.IsTerminating()

	if details.BaseDefinition.HasFinalizer(r.FinalizerName) {
		verifyResult, err := r.ResourceManagerClient.Verify(ctx, instance)

		if verifyResult.error() || err != nil {
			// TODO: log error (this should not happen, but we carry on allowing the result to delete
			removeFinalizer = true
		} else if verifyResult.missing() {
			removeFinalizer = true
		} else if verifyResult.deleting() {
			requeue = true
		} else if !isTerminating { // and one of verifyResult.ready() || verifyResult.recreateRequired() || verifyResult.updateRequired()
			deleteResult, err := r.ResourceManagerClient.Delete(ctx, instance)

			if err != nil || deleteResult.error() {
				// TODO: log error (this should not happen, but we carry on allowing the result to delete
				removeFinalizer = true
			} else if deleteResult.alreadyDeleted() || deleteResult.succeed() {
				removeFinalizer = true
			} else if deleteResult.awaitingVerification() {
				// set it back to pending and let it go through the whole process again
				requeue = true
			} else {
				// assert no more cases
				removeFinalizer = true
			}
		} else {
			// i.e. verify
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
		if err := r.updateInstance(ctx, details.Instance); err != nil {
			// it's going to be stuck if it ever gets to here
			return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, err
		}
		if !isTerminating {
			r.Recorder.Event(instance, corev1.EventTypeNormal, "Finalizer", "Setting state to terminating for "+details.Name)
		}
		if removeFinalizer {
			r.Recorder.Event(instance, corev1.EventTypeNormal, "Finalizer", "Removing finalizer for "+details.Name)
		}
	}

	if requeue {
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	} else {
		r.Recorder.Event(instance, corev1.EventTypeNormal, "Finalizer", details.Name+" finalizer complete")
		return ctrl.Result{}, nil
	}
}

// not sure why this doesn't work on the Cluster, seems to work fine on tests
func (r *AzureController) updateStatus(ctx context.Context, instance runtime.Object) error {
	err := r.KubeClient.Status().Update(ctx, instance)
	if err != nil {
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update CRD instance")
	}
	return err
}

func (r *AzureController) updateInstance(ctx context.Context, instance runtime.Object) error {
	err := r.KubeClient.Update(ctx, instance)
	if err != nil {
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update CRD instance")
	}
	return err
}
func (r *AzureController) updateAndLog(ctx context.Context, instance runtime.Object, eventType string, reason string, message string) error {
	if err := r.updateInstance(ctx, instance); err != nil {
		return err
	}
	r.Recorder.Event(instance, eventType, reason, message)
	return nil
}

func (r *AzureController) getRequeueAfter(requeueSeconds int) time.Duration {
	if requeueSeconds == 0 {
		requeueSeconds = 30
	}
	requeueAfter := time.Duration(requeueSeconds) * time.Second
	return requeueAfter
}
