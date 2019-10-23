package controller_refactor

import (
	"context"
	apimeta "k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
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
type PostProvisionHandler func(r runtime.Object) error

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
	thisDefs := ac.DefinitionManager.GetDefinition(ctx, req.NamespacedName)

	err := ac.KubeClient.Get(ctx, req.NamespacedName, thisDefs.InitialInstance)
	if err != nil {
		log.Info("unable to retrieve resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	instance := thisDefs.InitialInstance
	status, err := thisDefs.StatusGetter(instance)

	requeueAfter := getRequeueAfter(ac.Parameters.RequeueAfterSeconds)
	metaObject, _ := apimeta.Accessor(instance)

	instanceUpdater := customResourceUpdater{
		StatusUpdater: thisDefs.StatusUpdater,
	}

	// get dependency details
	dependencies, err := ac.DefinitionManager.GetDependencies(ctx, instance)
	// if any of the dependencies are not found, we jump out.
	if err != nil || dependencies == nil { // note that dependencies should be an empty array
		if apierrors.IsNotFound(err) {
			log.Info("dependency not found for " + req.Name + ". requeuing request.")
		} else {
			log.Info("unable to retrieve dependency for "+req.Name, "err", err.Error())
		}
		return ctrl.Result{Requeue: true, RequeueAfter: requeueAfter}, client.IgnoreNotFound(err)
	}

	// create a reconcile runner object. this runs a single cycle of the reconcile loop
	reconcileRunner := reconcileRunner{
		GenericController:     ac,
		ResourceDefinition:    thisDefs,
		DependencyDefinitions: dependencies,
		NamespacedName:        req.NamespacedName,
		instance:              instance,
		objectMeta:            metaObject,
		status:                status,
		req:                   req,
		requeueAfter:          requeueAfter,
		log:                   log,
		instanceUpdater:       &instanceUpdater,
	}

	reconcileFinalizer := reconcileFinalizer{
		reconcileRunner: reconcileRunner,
	}

	// if no finalizers have been defined, do that and requeue
	if !reconcileFinalizer.isDefined() {
		return reconcileFinalizer.add(ctx)
	}

	// if it's being deleted go straight to the finalizer step
	isBeingDeleted := !metaObject.GetDeletionTimestamp().IsZero()
	if isBeingDeleted {
		return reconcileFinalizer.handle()
	}

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
