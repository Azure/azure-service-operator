package controller_refactor

import (
	"context"
	"fmt"
	"time"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/go-logr/logr"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// AzureController reconciles a ResourceGroup object
type AzureController struct {
	Parameters            Parameters
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

type Parameters struct {
	RequeueAfterSeconds int
}

// SetupWithManager function sets up the functions with the controller
func (ac *AzureController) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.ResourceGroup{}).
		Complete(ac)
}

func (ac *AzureController) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.TODO()
	log := ac.Log.WithValues("resourcegroup", req.NamespacedName)

	// fetch the manifest object
	thisDefs, err := ac.DefinitionManager.GetThis(ctx, req)
	if err != nil {
		log.Info("Unable to retrieve resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	details := thisDefs.Details
	requeueAfter := getRequeueAfter(ac.Parameters.RequeueAfterSeconds)

	// create a reconcile cycle object
	reconcileCycle := reconcileRunner{
		AzureController:         ac,
		ThisResourceDefinitions: thisDefs,
		requeueAfter:            requeueAfter,
		log:                     log,
	}

	// if no finalizers have been defined, do that and requeue
	if len(details.BaseDefinition.Finalizers) == 0 {
		err := reconcileCycle.addFinalizers(ctx)
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
	dependencies, err := ac.DefinitionManager.GetDependencies(ctx, details.Instance)
	if err != nil || dependencies == nil {
		if apierrors.IsNotFound(err) {
			log.Info("Dependency not found for " + details.Name)
		} else {
			log.Info("Unable to retrieve dependency for "+details.Name, "err", err.Error())
		}
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, client.IgnoreNotFound(err)
	}
	reconcileCycle.DependencyDefinitions = dependencies

	// get parameters
	result, err := reconcileCycle.run(ctx)

	if result.Requeue == false {
		return ctrl.Result{RequeueAfter: getRequeueAfter(1)}, err
	}
	return result, err
}

func getRequeueAfter(requeueSeconds int) time.Duration {
	if requeueSeconds == 0 {
		requeueSeconds = 10
	}
	requeueAfter := time.Duration(requeueSeconds) * time.Second
	return requeueAfter
}
