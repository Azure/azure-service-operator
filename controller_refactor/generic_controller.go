package controller_refactor

import (
	"context"
	"time"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/go-logr/logr"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// GenericController reconciles a ResourceGroup object
type GenericController struct {
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
func (ac *GenericController) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&azurev1alpha1.ResourceGroup{}).
		Complete(ac)
}

func (ac *GenericController) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.TODO()
	log := ac.Log.WithValues("NamespacedName", req.NamespacedName)

	// fetch the manifest object
	thisDefs, err := ac.DefinitionManager.GetThis(ctx, req)
	if err != nil {
		log.Info("unable to retrieve resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	details := thisDefs.Details
	requeueAfter := getRequeueAfter(ac.Parameters.RequeueAfterSeconds)

	// create a reconcile runner object. this runs a single cycle of the reconcile loop
	baseDef := thisDefs.Details.BaseDefinition
	reconcileRunner := reconcileRunner{
		GenericController:       ac,
		ThisResourceDefinitions: thisDefs,
		DependencyDefinitions:   nil,
		NamespacedName:          req.NamespacedName,
		provisionState:          baseDef.Status.ProvisionState,
		req:                     req,
		requeueAfter:            requeueAfter,
		log:                     log,
	}

	reconcileFinalizer := reconcileFinalizer{
		reconcileRunner: reconcileRunner,
	}

	// if no finalizers have been defined, do that and requeue
	if !reconcileFinalizer.exists() {
		return reconcileFinalizer.add(ctx)
	}

	// if it's being deleted go straight to the finalizer step
	isBeingDeleted := !details.BaseDefinition.ObjectMeta.DeletionTimestamp.IsZero()
	if isBeingDeleted {
		return reconcileFinalizer.handle()
	}

	// get dependency details
	dependencies, err := ac.DefinitionManager.GetDependencies(ctx, details.Instance)
	if err != nil || dependencies == nil { // note that dependencies should be an empty array
		if apierrors.IsNotFound(err) {
			log.Info("dependency not found for " + details.BaseDefinition.Name + ". requeuing request.")
		} else {
			log.Info("unable to retrieve dependency for "+details.BaseDefinition.Name, "err", err.Error())
		}
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, client.IgnoreNotFound(err)
	}
	reconcileRunner.DependencyDefinitions = dependencies

	// run a single cycle of the reconcile loop
	return reconcileRunner.run(ctx)
}

func getRequeueAfter(requeueSeconds int) time.Duration {
	if requeueSeconds == 0 {
		requeueSeconds = 10
	}
	requeueAfter := time.Duration(requeueSeconds) * time.Second
	return requeueAfter
}
