package controller_refactor

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"time"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
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

// Reconcile function does the main reconciliation loop of the operator
func (r *AzureController) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("resourcegroup", req.NamespacedName)

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

	if details.IsBeingDeleted {
		result, err := r.handleFinalizer(details, updater, r.FinalizerName)
		if err != nil {
			return result, fmt.Errorf("error when handling finalizer: %v", err)
		}
		return result, nil
	}

	if !details.BaseDefinition.HasFinalizer(r.FinalizerName) {
		err := r.addFinalizer(details, updater, r.FinalizerName)
		if err != nil {
			return ctrl.Result{}, fmt.Errorf("error when removing finalizer: %v", err)
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
	if requeueSeconds == 0 {
		requeueSeconds = 30
	}
	requeueAfter := time.Duration(requeueSeconds) * time.Second

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
		if err := r.updateInstance(ctx, instance); err != nil {
			return ctrl.Result{}, err
		}
		return ctrl.Result{}, nil
	}

	// now verify the resource state on Azure
	if details.ProvisionState.IsVerifying() || details.ProvisionState.IsPending() {
		verifyResult, err := r.verifyExternal(details, updater)
		if err != nil {
			// verification should not return an error - if this happens it's a terminal failure
			updater.SetProvisionState(azurev1alpha1.Failed)
			if err := r.updateInstance(ctx, instance); err != nil {
				return ctrl.Result{}, err
			}
			return ctrl.Result{}, fmt.Errorf("error verifying resource in azure: %v", err)
		}
		
		// Success case - the resource is is provisioned on Azure, and is ready
		if verifyResult.ready() {
			updater.SetProvisionState(azurev1alpha1.Succeeded)

			if err := r.updateInstance(ctx, instance); err != nil {
				return ctrl.Result{}, err
			}
			r.Recorder.Event(instance, corev1.EventTypeNormal, "Updated", details.Name + " provisioned")
		}

		// Missing case - we can now create the resource
		if verifyResult.missing() {
			updater.SetProvisionState(azurev1alpha1.Creating)

			if err := r.updateInstance(ctx, instance); err != nil {
				return ctrl.Result{}, err
			}
			r.Recorder.Event(instance, corev1.EventTypeNormal, "Creating", details.Name + " ready for creation")
		}

		// Update case - the resource exists in Azure, is invalid but updateable, so doesn't need to be recreated
		if verifyResult.updateRequired() {
			updater.SetProvisionState(azurev1alpha1.Updating)

			if err := r.updateInstance(ctx, instance); err != nil {
				return ctrl.Result{}, err
			}
			r.Recorder.Event(instance, corev1.EventTypeNormal, "Updated", details.Name + " provisioned")
		}

		// Recreate case - the resource exists in Azure, is invalid and needs to be created
		if verifyResult.recreateRequired() {
			if err = r.ResourceManagerClient.Delete(ctx, details.Instance); err != nil {
				// TODO: proper error handling for delete
				updater.SetProvisionState(azurev1alpha1.Failed)
				return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, err
			}

			// set it back to pending and let it go through the whole process again
			updater.SetProvisionState(azurev1alpha1.Pending)
			if err := r.updateInstance(ctx, instance); err != nil {
				return ctrl.Result{}, err
			}
			r.Recorder.Event(instance, corev1.EventTypeNormal, "Updated", details.Name + " provisioned")
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
		updateErr := r.updateInstance(ctx, instance)

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

// SetupWithManager function sets up the functions with the controller
func (r *AzureController) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.ResourceGroup{}).
		Complete(r)
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

func (r *AzureController) addFinalizer(details *CustomResourceDetails, updater *CustomResourceUpdater, finalizerName string) error {
	updater.AddFinalizer(finalizerName)
	updater.SetProvisionState(azurev1alpha1.Pending)
	if err := r.updateInstance(context.Background(), details.Instance); err != nil {
		return err
	}
	r.Recorder.Event(details.Instance, corev1.EventTypeNormal, "Updated", fmt.Sprintf("finalizer %s added", finalizerName))
	return nil
}

func (r *AzureController) handleFinalizer(details *CustomResourceDetails, updater *CustomResourceUpdater, finalizerName string) (ctrl.Result, error) {
	if details.BaseDefinition.HasFinalizer(finalizerName) {
		ctx := context.Background()
		if err := r.ResourceManagerClient.Delete(ctx, details.Instance); err != nil {
			catch := []string{
				errhelp.AsyncOpIncompleteError,
			}
			if azerr, ok := err.(*errhelp.AzureError); ok {
				if helpers.ContainsString(catch, azerr.Type) {
					r.Log.Info("Got ignorable error", "type", azerr.Type)
					return ctrl.Result{Requeue: true, RequeueAfter: 30 * time.Second}, nil
				}
			}
			r.Log.Info("Delete AzureSqlServer failed with ", "error", err.Error())

			return ctrl.Result{}, err
		}

		updater.RemoveFinalizer(r.FinalizerName)
		if err := r.updateInstance(ctx, details.Instance); err != nil {
			return ctrl.Result{}, err
		}
	}

	// Our finalizer has finished, so the reconciler can do nothing.
	return ctrl.Result{}, nil
}
//
//func (r *AzureController) updateStatus(ctx context.Context, instance runtime.Object) error {
//	err := r.KubeClient.Status().Update(ctx, instance)
//	if err != nil {
//		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update CRD instance")
//	}
//	return err
//}

func (r *AzureController) updateInstance(ctx context.Context, instance runtime.Object) error {
	err := r.KubeClient.Update(ctx, instance)
	if err != nil {
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update CRD instance")
	}
	return err
}
