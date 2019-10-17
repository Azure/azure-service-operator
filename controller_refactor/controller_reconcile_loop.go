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

	// set the owner reference if owner is present
	if owner != nil {
		//set owner reference if it exists
		updater.SetOwnerReference(owner)
	}

	// dependencies are now satisfied, can now reconcile the manfest and create or update the resource
	if details.ProvisionState.IsPending() {
		r.Recorder.Event(details.Instance, corev1.EventTypeNormal, "Submitting", "starting resource reconciliation")
		// TODO: Add error handling for cases where username or password are invalid:
		// https://docs.microsoft.com/en-us/rest/api/sql/servers/createorupdate#response

		nextState, reconErr := r.reconcileExternal(details, updater)
		instance := details.Instance

		updater.SetProvisionState(nextState)
		updateErr := r.updateInstance(ctx, instance)

		if reconErr != nil {
			return ctrl.Result{}, fmt.Errorf("error reconciling resource in azure: %v", reconErr)
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

	// now verify the resource has been created
	if details.ProvisionState.IsVerifying() {
		verifyResult, err := r.verifyExternal(details, updater)
		if err != nil {
			catch := []string{
				errhelp.ParentNotFoundErrorCode,
				errhelp.NotFoundErrorCode,
				errhelp.ResourceNotFound,
				errhelp.AsyncOpIncompleteError,
			}
			if azerr, ok := err.(*errhelp.AzureError); ok {
				if helpers.ContainsString(catch, azerr.Type) {
					log.Info("Got ignorable error", "type", azerr.Type)
					return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
				}
			}
			return ctrl.Result{}, fmt.Errorf("error verifying resource in azure: %v", err)
		}

		// if still is in progress with provisioning, requeue the recocile loop
		if verifyResult.IsProvisioning() {
			log.Info("Retrying verification", "type", "Verification not complete, requeuing reconcile loop")
			return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
		}

		return ctrl.Result{}, nil
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

func (r *AzureController) reconcileExternal(crDetails *CustomResourceDetails, updater *CustomResourceUpdater) (azurev1alpha1.ProvisionState, error) {

	ctx := context.Background()
	var err error

	resourceName := crDetails.Name
	instance := crDetails.Instance

	// ensure that the resource is created or updated in Azure (though it won't necessarily be ready, it still needs to be verified)
	ensureResult, err := r.ResourceManagerClient.Ensure(ctx, instance)
	if err != nil || ensureResult.Failed() || ensureResult.InvalidRequest() {
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Couldn't create or update resource in azure")
		return azurev1alpha1.Failed, err
	}

	// if successful, set it to succeeded, or await verification
	var nextState azurev1alpha1.ProvisionState
	if ensureResult.AwaitingVerification() {
		nextState = azurev1alpha1.Verifying
	} else if ensureResult.Succeeded() {
		nextState = azurev1alpha1.Succeeded
	} else {
		return azurev1alpha1.Failed, errhelp.NewAzureError(fmt.Errorf("invalid response from Ensure for resource '%s'", resourceName))
	}
	return nextState, nil
}

func (r *AzureController) verifyExternal(crDetails *CustomResourceDetails, updater *CustomResourceUpdater) (VerifyResult, error) {
	ctx := context.Background()
	instance := crDetails.Instance
	resourceName := crDetails.Name

	r.Recorder.Event(instance, corev1.EventTypeNormal, "Checking", "instance is ready")
	verifyResult, err := r.ResourceManagerClient.Verify(ctx, instance)

	if err != nil {
		r.Recorder.Event(crDetails.Instance, corev1.EventTypeWarning, "Failed", "Couldn't validate resource in azure")
		return verifyResult, errhelp.NewAzureError(err)
	}
	if verifyResult.IsReady() {
		updater.SetProvisionState(azurev1alpha1.Succeeded)

		if err := r.updateInstance(ctx, instance); err != nil {
			return VerifyError, err
		}

		r.Recorder.Event(instance, corev1.EventTypeNormal, "Updated", resourceName+" provisioned")
	}
	return verifyResult, nil
}

func (r *AzureController) addFinalizer(crDetails *CustomResourceDetails, updater *CustomResourceUpdater, finalizerName string) error {
	updater.AddFinalizer(finalizerName)
	updater.SetProvisionState(azurev1alpha1.Pending)
	if err := r.updateInstance(context.Background(), crDetails.Instance); err != nil {
		return err
	}
	r.Recorder.Event(crDetails.Instance, corev1.EventTypeNormal, "Updated", fmt.Sprintf("finalizer %s added", finalizerName))
	return nil
}

func (r *AzureController) handleFinalizer(crDetails *CustomResourceDetails, updater *CustomResourceUpdater, finalizerName string) (ctrl.Result, error) {
	if crDetails.BaseDefinition.HasFinalizer(finalizerName) {
		ctx := context.Background()
		if err := r.ResourceManagerClient.Delete(ctx, crDetails.Instance); err != nil {
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
		if err := r.updateInstance(ctx, crDetails.Instance); err != nil {
			return ctrl.Result{}, err
		}
	}

	// Our finalizer has finished, so the reconciler can do nothing.
	return ctrl.Result{}, nil
}

func (r *AzureController) updateInstance(ctx context.Context, instance runtime.Object) error {
	err := r.KubeClient.Update(ctx, instance)
	if err != nil {
		r.Recorder.Event(instance, corev1.EventTypeWarning, "Failed", "Unable to update CRD instance")
	}
	return err
}
