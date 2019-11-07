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
	"fmt"

	"github.com/go-logr/logr"
	apimeta "k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"

	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// GenericController is a generic implementation of a Kubebuilder controller
type GenericController struct {
	Parameters         ReconcileParameters
	ResourceKind       string
	KubeClient         client.Client
	Log                logr.Logger
	Recorder           record.EventRecorder
	Scheme             *runtime.Scheme
	ResourceManager    ResourceManager
	DefinitionManager  DefinitionManager
	FinalizerName      string
	AnnotationBaseName string
	CompletionRunner   func(*GenericController) CompletionRunner
}

// A handler that is invoked after the resource has been successfully created
// and it has been verified to be ready for consumption (ReconcileState=Success)
// This is typically used for example to create secrets with authentication information
type CompletionRunner interface {
	Run(ctx context.Context, r runtime.Object) error
}

type ReconcileParameters struct {
	RequeueAfter        int
	RequeueAfterSuccess int
	RequeueAfterFailure int
}

func CreateGenericController(
	parameters ReconcileParameters,
	resourceKind string,
	kubeClient client.Client,
	logger logr.Logger,
	recorder record.EventRecorder,
	scheme *runtime.Scheme,
	resourceManager ResourceManager,
	defMgr DefinitionManager,
	finalizerName string,
	annotationBaseName string,
	completionRunner func(*GenericController) CompletionRunner) (*GenericController, error) {
	gc := &GenericController{
		Parameters:         parameters,
		ResourceKind:       resourceKind,
		KubeClient:         kubeClient,
		Log:                logger,
		Recorder:           recorder,
		Scheme:             scheme,
		ResourceManager:    resourceManager,
		DefinitionManager:  defMgr,
		FinalizerName:      finalizerName,
		AnnotationBaseName: annotationBaseName,
		CompletionRunner:   completionRunner,
	}
	if err := gc.validate(); err != nil {
		return nil, err
	}
	return gc, nil
}

func (gc *GenericController) validate() error {
	if gc.ResourceKind == "" {
		return fmt.Errorf("resource Kind must be defined for GenericController")
	}
	kind := gc.ResourceKind
	if gc.Scheme == nil {
		return fmt.Errorf("no Scheme defined for controller for %s", kind)
	}
	if gc.ResourceManager == nil {
		return fmt.Errorf("no ResourceManager defined for controller for %s", kind)
	}
	if gc.DefinitionManager == nil {
		return fmt.Errorf("no DefinitionManager defined for controller for %s", kind)
	}
	if gc.FinalizerName == "" {
		return fmt.Errorf("no FinalizerName set for controller for %s", kind)
	}
	return nil
}

func (gc *GenericController) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.TODO()
	log := gc.Log.WithValues("Name", req.NamespacedName)

	// fetch the manifest object
	thisDefs := gc.DefinitionManager.GetDefinition(ctx, req.NamespacedName)

	err := gc.KubeClient.Get(ctx, req.NamespacedName, thisDefs.InitialInstance)
	if err != nil {
		log.Info("Unable to retrieve resource", "err", err.Error())
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	instance := thisDefs.InitialInstance
	status, err := thisDefs.StatusAccessor(instance)
	metaObject, _ := apimeta.Accessor(instance)

	instanceUpdater := instanceUpdater{
		StatusUpdater: thisDefs.StatusUpdater,
	}

	// get dependency details
	dependencies, err := gc.DefinitionManager.GetDependencies(ctx, instance)
	// this is only the names and initial values, if we can't fetch these it's terminal
	if err != nil {
		log.Info("Unable to retrieve dependencies for resource ", "err", err.Error())
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// create a reconcile runner object. this runs a single cycle of the reconcile loop
	reconcileRunner := reconcileRunner{
		GenericController:     gc,
		ResourceDefinition:    thisDefs,
		DependencyDefinitions: dependencies,
		NamespacedName:        req.NamespacedName,
		instance:              instance,
		objectMeta:            metaObject,
		status:                status,
		req:                   req,
		log:                   log,
		instanceUpdater:       &instanceUpdater,
	}

	// handle finalization first
	reconcileFinalizer := reconcileFinalizer{
		reconcileRunner: reconcileRunner,
	}

	// if it's being deleted go straight to the finalizer step
	isBeingDeleted := !metaObject.GetDeletionTimestamp().IsZero()
	if isBeingDeleted {
		return reconcileFinalizer.handle()
	}

	// if no finalizers have been defined, do that and requeue
	if !reconcileFinalizer.isDefined() {
		return reconcileFinalizer.add(ctx)
	}

	// run a single cycle of the reconcile loop
	return reconcileRunner.run(ctx)
}
