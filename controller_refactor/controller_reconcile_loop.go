package controller_refactor

import (
	"fmt"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	v1 "k8s.io/api/core/v1"
	"time"

	"context"

	"github.com/go-logr/logr"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// AzureController reconciles a ResourceGroup object
type AzureController struct {
	ResourceKind         string
	KubeClient           client.Client
	Log                  logr.Logger
	Recorder             record.EventRecorder
	ResourceClient       ResourceManagerClient
	DefinitionManager    DefinitionManager
	FinalizerName        string
	PostProvisionHandler PostProvisionHandler
}

// Reconcile function does the main reconciliation loop of the operator
func (r *AzureController) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("resourcegroup", req.NamespacedName)

	thisDefs, err := r.DefinitionManager.GetThis(ctx, r.KubeClient, req)
	if err != nil {
		log.Info("Unable to retrieve resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	definition := thisDefs.CRDInfo
	updater := thisDefs.CRDUpdater

	if definition.IsBeingDeleted {
		result, err := r.handleFinalizer(definition, updater, r.FinalizerName)
		if err != nil {
			return result, fmt.Errorf("error when handling finalizer: %v", err)
		}
		return result, nil
	}

	if !updater.HasFinalizer(r.FinalizerName) {
		err := r.addFinalizer(definition, updater, r.FinalizerName)
		if err != nil {
			return ctrl.Result{}, fmt.Errorf("error when removing finalizer: %v", err)
		}
		return ctrl.Result{}, nil
	}

	// verify status of dependencies
	dependencyInfo, err := r.DefinitionManager.GetDependencies(ctx, r.KubeClient, req)
	if err != nil {
		log.Info("Unable to retrieve resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// get parameters
	parameters := definition.Parameters
	requeueSeconds := parameters.RequeueAfterSeconds
	if requeueSeconds == 0 {
		requeueSeconds = 30
	}
	requeueAfter := time.Duration(requeueSeconds) * time.Second

	// Verify that all dependencies are present in the cluster, and they are
	owner := dependencyInfo.Owner
	allDeps := append([]*CRDInfo{owner}, dependencyInfo.Dependencies...)

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

	// dependencies are now satisfied, can
	if definition.ProvisionState.IsPending() {
		r.Recorder.Event(definition.CRDInstance, v1.EventTypeNormal, "Submitting", "starting resource reconciliation")
		// TODO: Add error handling for cases where username or password are invalid:
		// https://docs.microsoft.com/en-us/rest/api/sql/servers/createorupdate#response
		if err := r.reconcileExternal(definition, updater); err != nil {
			catch := []string{
				errhelp.ParentNotFoundErrorCode,
				errhelp.ResourceGroupNotFoundErrorCode,
				errhelp.NotFoundErrorCode,
				errhelp.AsyncOpIncompleteError,
				errhelp.InvalidServerName,
			}
			if azerr, ok := err.(*errhelp.AzureError); ok {
				if helpers.ContainsString(catch, azerr.Type) {
					if azerr.Type == errhelp.InvalidServerName {
						r.Recorder.Event(definition.CRDInstance, v1.EventTypeWarning, "Failed", "Invalid Server Name")
						return ctrl.Result{Requeue: false}, nil
					}
					log.Info("Got ignorable error", "type", azerr.Type)
					return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
				}
			}
			return ctrl.Result{}, fmt.Errorf("error reconciling sql server in azure: %v", err)
		}
		// give azure some time to catch up
		log.Info("waiting for provision to take effect")
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
	}

	if definition.ProvisionState.IsVerifying() {
		verifyResult, err := r.verifyExternal(definition, updater)
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
			return ctrl.Result{}, fmt.Errorf("error verifying sql server in azure: %v", err)
		}
		if verifyResult.IsProvisioning() {
			log.Info("Retrying verification", "type", "verification not complete")
			return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, nil
		}

		return ctrl.Result{}, nil
	}

	if definition.ProvisionState.IsSucceeded() && r.PostProvisionHandler != nil {
		if err := r.PostProvisionHandler(definition); err != nil {
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

func (r *AzureController) reconcileExternal(definition *CRDInfo, updater *CRDUpdater) error {

	ctx := context.Background()
	var err error

	resourceName := definition.Name
	instance := definition.CRDInstance

	// write information back to instance
	updater.SetState(azurev1alpha1.Provisioning)
	err = r.KubeClient.Update(ctx, instance)
	if err != nil {
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
	}

	err = r.ResourceClient.Ensure(ctx, instance)
	if err != nil {
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Couldn't create resource in azure")
		updater.SetState(azurev1alpha1.Failed)
		errUpdate := r.KubeClient.Update(ctx, instance)
		if errUpdate != nil {
			r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
		}
		return err
	}

	// write information back to instance
	updater.SetState(azurev1alpha1.Verifying)
	err = r.KubeClient.Update(ctx, instance)
	if err != nil {
		//log error and kill it
		r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
	}

	r.Recorder.Event(instance, v1.EventTypeNormal, "Updated", resourceName+" provisioned")

	return nil
}

func (r *AzureController) verifyExternal(definition *CRDInfo, updater *CRDUpdater) (VerifyResult, error) {
	ctx := context.Background()
	instance := definition.CRDInstance
	resourceName := definition.Name

	r.Recorder.Event(instance, v1.EventTypeNormal, "Checking", "instance is ready")
	verifyResult, err := r.ResourceClient.Verify(ctx, instance)

	if err != nil {
		r.Recorder.Event(definition.CRDInstance, v1.EventTypeWarning, "Failed", "Couldn't validate resource in azure")
		return verifyResult, errhelp.NewAzureError(err)
	}
	if verifyResult.IsReady() {
		updater.SetState(azurev1alpha1.Succeeded)
		err = r.KubeClient.Update(ctx, instance)
		if err != nil {
			//log error and kill it
			r.Recorder.Event(instance, v1.EventTypeWarning, "Failed", "Unable to update instance")
		}

		r.Recorder.Event(instance, v1.EventTypeNormal, "Updated", resourceName+" provisioned")
	}
	return verifyResult, nil
}

func (r *AzureController) addFinalizer(definition *CRDInfo, updater *CRDUpdater, finalizerName string) error {
	updater.AddFinalizer(finalizerName)
	updater.SetState(azurev1alpha1.Pending)
	if updateerr := r.KubeClient.Update(context.Background(), definition.CRDInstance); updateerr != nil {
		r.Recorder.Event(definition.CRDInstance, v1.EventTypeWarning, "Failed", "Failed to update finalizer")
	}
	r.Recorder.Event(definition.CRDInstance, v1.EventTypeNormal, "Updated", fmt.Sprintf("finalizer %s added", finalizerName))
	return nil
}

func (r *AzureController) handleFinalizer(definition *CRDInfo, updater *CRDUpdater, finalizerName string) (ctrl.Result, error) {
	if updater.HasFinalizer(finalizerName) {
		ctx := context.Background()
		if err := r.ResourceClient.Delete(ctx, definition.CRDInstance); err != nil {
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
		if err := r.KubeClient.Update(ctx, definition.CRDInstance); err != nil {
			return ctrl.Result{}, err
		}
	}

	// Our finalizer has finished, so the reconciler can do nothing.
	return ctrl.Result{}, nil
}
