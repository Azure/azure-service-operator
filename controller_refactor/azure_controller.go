package controller_refactor

import (
	"context"
	"fmt"
	"time"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/go-logr/logr"

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

// A handler that is invoked after the resource has been successfully created
// and it has been verified to be ready for consumption (ProvisionState=Success)
// This is typically used for example to create secrets with authentication information
type PostProvisionHandler func(definition *CustomResourceDetails) error

// SetupWithManager function sets up the functions with the controller
func (r *AzureController) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.ResourceGroup{}).
		Complete(r)
}

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
	parameters := details.Parameters
	requeueAfter := getRequeueAfter(parameters)

	// create a reconcile cycle object
	reconcileCycle := reconcileRunner{
		AzureController:         r,
		ThisResourceDefinitions: thisDefs,
		requeueAfter:            requeueAfter,
		ctx:                     ctx,
		log:                     log,
	}

	// if no finalizers have been defined, do that and requeue
	if len(details.BaseDefinition.Finalizers) == 0 {
		err := reconcileCycle.addFinalizers()
		if err != nil {
			return ctrl.Result{}, fmt.Errorf("error adding finalizer: %v", err)
		}
		return ctrl.Result{}, nil
	}

	// if it's being deleted go straight to the finalizer step
	if details.IsBeingDeleted {
		result, err := reconcileCycle.handleFinalizer()
		if err != nil {
			return result, fmt.Errorf("error when handling finalizer: %v", err)
		}
		return result, nil
	}

	// get dependency details
	reconcileCycle.DependencyDefinitions, err = r.DefinitionManager.GetDependencies(ctx, req)
	if err != nil {
		log.Info("Unable to retrieve resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// get parameters
	return reconcileCycle.run()
}

func getRequeueAfter(parameters azurev1alpha1.Parameters) time.Duration {
	requeueSeconds := parameters.RequeueAfterSeconds
	if requeueSeconds == 0 {
		requeueSeconds = 30
	}
	requeueAfter := time.Duration(requeueSeconds) * time.Second
	return requeueAfter
}

