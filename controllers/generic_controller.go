/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"github.com/Azure/k8s-infra/apis"
	microsoftnetworkv1 "github.com/Azure/k8s-infra/apis/microsoft.network/v1"
	microsoftresourcesv1 "github.com/Azure/k8s-infra/apis/microsoft.resources/v1"
	"github.com/Azure/k8s-infra/pkg/util/patch"
	"github.com/Azure/k8s-infra/pkg/zips"
)

const (
	// ResourceSigAnnotationKey is an annotation key which holds the value of the hash of the spec
	ResourceSigAnnotationKey = "resource-sig.infra.azure.com"
)

var (

	// KnownTypes defines an array of runtime.Objects to be reconciled, where each
	// object in the array will generate a controller. If the concrete type
	// implements the owner interface, the generated controller will inject the
	// owned types supplied by the Owns() method of the CRD type. Each controller
	// may directly reconcile a single object, but may indirectly watch
	// and reconcile many Owned objects. The singular type is necessary to generically
	// produce a reconcile function aware of concrete types, as a closure.
	KnownTypes = []runtime.Object{
		new(microsoftresourcesv1.ResourceGroup),
		new(microsoftnetworkv1.VirtualNetwork),
	}
)

type (
	owner interface {
		Owns() []runtime.Object
	}

	// +kubebuilder:rbac:groups=core,resources=events,verbs=get;list;watch;create;patch
	// +kubebuilder:rbac:groups=microsoft.resources.infra.azure.com,resources=resourcegroups,verbs=get;list;watch;create;update;patch;delete
	// +kubebuilder:rbac:groups=microsoft.resources.infra.azure.com,resources=resourcegroups/status,verbs=get;update;patch
	// +kubebuilder:rbac:groups=microsoft.network.infra.azure.com,resources=virtualnetworks,verbs=get;list;watch;create;update;patch;delete
	// +kubebuilder:rbac:groups=microsoft.network.infra.azure.com,resources=virtualnetworks/status,verbs=get;update;patch

	// GenericReconciler reconciles a Resourcer object
	GenericReconciler struct {
		Client     client.Client
		Log        logr.Logger
		Owns       []runtime.Object
		Applier    zips.Applier
		Scheme     *runtime.Scheme
		Recorder   record.EventRecorder
		Name       string
		GVK        schema.GroupVersionKind
		Controller controller.Controller
	}
)

func RegisterAll(mgr ctrl.Manager, applier zips.Applier, objs []runtime.Object, log logr.Logger, options controller.Options) []error {
	var errs []error
	for _, obj := range objs {
		mgr := mgr
		applier := applier
		obj := obj
		if err := register(mgr, applier, obj, log, options); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

// register takes a manager and a struct describing how to instantiate
// controllers for various types using a generic reconciler function. Only one
// type (using For) may be directly watched by each controller, but zero, one or
// many Owned types are acceptable. This setup allows reconcileFn to have access
// to the concrete type defined as part of a closure, while allowing for
// independent controllers per GVK (== better parallelism, vs 1 controller
// managing many, many List/Watches)
func register(mgr ctrl.Manager, applier zips.Applier, obj runtime.Object, log logr.Logger, options controller.Options) error {
	v, err := conversion.EnforcePtr(obj)
	if err != nil {
		return err
	}

	t := v.Type()
	controllerName := fmt.Sprintf("%sCtrl", t.Name())

	// Use the provided GVK to construct a new runtime object of the desired concrete type.
	gvk, err := apiutil.GVKForObject(obj, mgr.GetScheme())
	if err != nil {
		return err
	}

	reconciler := &GenericReconciler{
		Client:   mgr.GetClient(),
		Applier:  applier,
		Scheme:   mgr.GetScheme(),
		Name:     t.Name(),
		Log:      log.WithName(controllerName),
		Recorder: mgr.GetEventRecorderFor(controllerName),
		GVK:      gvk,
	}

	ctrlBuilder := ctrl.NewControllerManagedBy(mgr).
		For(obj).
		WithOptions(options)

	ownerType, ok := obj.(owner)
	if ok {
		reconciler.Owns = ownerType.Owns()
		for _, own := range ownerType.Owns() {
			ctrlBuilder.Owns(own)
		}
	}

	c, err := ctrlBuilder.Build(reconciler)
	if err != nil {
		return fmt.Errorf("unable to build controller / reconciler with: %w", err)
	}

	reconciler.Controller = c
	return nil
}

// Reconcile will take state in K8s and apply it to Azure
func (gr *GenericReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := gr.Log.WithValues("name", req.Name, "namespace", req.Namespace)

	obj, err := gr.Scheme.New(gr.GVK)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("unable to create object from gvk %+v with: %w", gr.GVK, err)
	}

	if err := gr.Client.Get(ctx, req.NamespacedName, obj); err != nil {
		if apierrors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}

		return ctrl.Result{}, err
	}

	// The Go type for the Kubernetes object must understand how to
	// convert itself to/from the corresponding Azure types.
	resourcer, ok := obj.(zips.Resourcer)
	if !ok {
		return ctrl.Result{}, fmt.Errorf("object: %+v is not a zips.Resourcer", obj)
	}

	// reconcile delete
	if !resourcer.GetDeletionTimestamp().IsZero() {
		log.Info("reconcile delete start")
		result, err := gr.reconcileDelete(ctx, resourcer)
		if err != nil {
			gr.Recorder.Event(resourcer, v1.EventTypeWarning, "ReconcileDeleteError", err.Error())
			log.Error(err, "reconcile delete error")
			return result, err
		}

		log.Info("reconcile delete complete")
		return result, err
	}

	if grouped, ok := obj.(apis.Grouped); ok {
		ready, err := gr.isResourceGroupReady(ctx, grouped, log)
		if err != nil {
			log.Error(err, "failed checking if resource group was ready")
			gr.Recorder.Event(resourcer, v1.EventTypeWarning, "GroupReadyError", fmt.Sprintf("isResourceGroupReady failed with: %s", err))
			return ctrl.Result{}, err
		}

		if !ready {
			requeueTime := 30 * time.Second
			msg := fmt.Sprintf("resource group %q is not ready or not created yet; will try again in about %s", grouped.GetResourceGroupObjectRef().Name, requeueTime)
			gr.Recorder.Event(resourcer, v1.EventTypeNormal, "ResourceGroupNotReady", msg)
			return ctrl.Result{
				RequeueAfter: requeueTime,
			}, nil
		}
	}

	log.Info("reconcile apply start")
	result, err := gr.reconcileApply(ctx, resourcer, log)
	if err != nil {
		log.Error(err, "reconcile apply error")
		gr.Recorder.Event(resourcer, v1.EventTypeWarning, "ReconcileError", err.Error())
		return result, err
	}

	log.Info("reconcile apply complete")
	return result, err
}

// reconcileApply will determine what, if anything, has changed on the resource, and apply that state to Azure.
// The Az infra finalizer will be applied.
//
// There are 3 possible state transitions.
// *  New resource (hasChanged = true) --> Start deploying the resource and move provisioning state from "" to "{Accepted || Succeeded || Failed}"
// *  Existing Resource (hasChanged = true) --> Start deploying the resource and move state from "{terminal state} to "{Accepted || Succeeded || Failed}"
// *  Existing Resource (hasChanged = false) --> Probably a status change. Don't do anything as of now.
func (gr *GenericReconciler) reconcileApply(ctx context.Context, resourcer zips.Resourcer, log logr.Logger) (ctrl.Result, error) {
	// check if the hash on the resource has changed
	hasChanged, err := hasResourceHashAnnotationChanged(resourcer)
	if err != nil {
		err = fmt.Errorf("failed comparing resource hash with: %w", err)
		gr.Recorder.Event(resourcer, v1.EventTypeWarning, "AnnotationError", err.Error())
		return ctrl.Result{}, err
	}

	azObj, err := resourcer.ToResource()
	if err != nil {
		err := fmt.Errorf("unable to transform to resource with: %w", err)
		gr.Recorder.Event(resourcer, v1.EventTypeWarning, "ToResourceError", err.Error())
		return ctrl.Result{}, err
	}

	// if the resource hash (spec) has not changed, don't apply again
	if !hasChanged && zips.IsTerminalProvisioningState(azObj.ProvisioningState) {
		msg := fmt.Sprintf("resource in state %q and spec has not changed", azObj.ProvisioningState)
		gr.Recorder.Event(resourcer, v1.EventTypeNormal, "ResourceHasNotChanged", msg)
		return ctrl.Result{}, nil
	}

	switch {
	case hasChanged:
		msg := fmt.Sprintf("resource in state %q has changed and spec will be applied to Azure", azObj.ProvisioningState)
		gr.Recorder.Event(resourcer, v1.EventTypeNormal, "ResourceHasChanged", msg)
		return gr.applySpecChange(ctx, resourcer)
	case !zips.IsTerminalProvisioningState(azObj.ProvisioningState):
		msg := fmt.Sprintf("resource in state %q, asking Azure for updated state", azObj.ProvisioningState)
		gr.Recorder.Event(resourcer, v1.EventTypeNormal, "ResourceStateNonTerminal", msg)
		return gr.updateFromNonTerminalApplyState(ctx, resourcer)
	default:
		msg := fmt.Sprintf("resource in state %q and spec has not changed", azObj.ProvisioningState)
		gr.Recorder.Event(resourcer, v1.EventTypeNormal, "ResourceNoopReconcile", msg)
		return ctrl.Result{}, nil // all is good and there are no changes to deal with
	}
}

func (gr *GenericReconciler) isResourceGroupReady(ctx context.Context, grouped apis.Grouped, log logr.Logger) (bool, error) {
	// has a resource group, so check if the resource group is already provisioned
	groupRef := grouped.GetResourceGroupObjectRef()
	if groupRef == nil {
		return false, fmt.Errorf("grouped resources must have a resource group")
	}

	key := client.ObjectKey{
		Name:      groupRef.Name,
		Namespace: groupRef.Namespace,
	}

	// get the storage version of the resource group regardless of the referenced version
	var rg microsoftresourcesv1.ResourceGroup
	if err := gr.Client.Get(ctx, key, &rg); err != nil {
		if apierrors.IsNotFound(err) {
			// not able to find the resource, but that's ok. It might not exist yet
			return false, nil
		}
		return false, fmt.Errorf("error GETing rg resource with: %w", err)
	}

	res, err := rg.ToResource()
	if err != nil {
		return false, fmt.Errorf("unable to transform to resource with: %w", err)
	}

	return res.ProvisioningState == zips.SucceededProvisioningState, nil
}

// reconcileDelete will begin and follow a delete operation of a resource. The finalizer will only be removed upon the
// resource actually being deleted in Azure.
//
// There are 2 possible state transitions.
// *  obj.ProvisioningState == \*\ --> Start deleting in Azure and mark state as "Deleting"
// *  obj.ProvisioningState == "Deleting" --> http HEAD to see if resource still exists in Azure. If so, requeue, else, remove finalizer.
func (gr *GenericReconciler) reconcileDelete(ctx context.Context, resourcer zips.Resourcer) (ctrl.Result, error) {
	azObj, err := resourcer.ToResource()
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("unable to transform to resource with: %w", err)
	}

	switch azObj.ProvisioningState {
	case zips.DeletingProvisioningState:
		msg := fmt.Sprintf("deleting... checking for updated state")
		gr.Recorder.Event(resourcer, v1.EventTypeNormal, "ResourceDeleteInProgress", msg)
		return gr.updateFromNonTerminalDeleteState(ctx, resourcer)
	default:
		msg := fmt.Sprintf("start deleting resource in state %q", azObj.ProvisioningState)
		gr.Recorder.Event(resourcer, v1.EventTypeNormal, "ResourceDeleteStart", msg)
		return gr.startDeleteOfResource(ctx, resourcer)
	}
}

// startDeleteOfResource will begin the delete of a resource by telling Azure to start deleting it. The resource will be
// marked with the provisioning state of "Deleting".
func (gr *GenericReconciler) startDeleteOfResource(ctx context.Context, resourcer zips.Resourcer) (ctrl.Result, error) {
	azObj, err := resourcer.ToResource()
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("unable to transform to resource with: %w", err)
	}

	if err := patcher(ctx, gr.Client, resourcer, func(res zips.Resourcer) error {
		if azObj, err = gr.Applier.BeginDelete(ctx, azObj); err != nil {
			return fmt.Errorf("failed trying to delete with %w", err)
		}

		azObj.ProvisioningState = zips.DeletingProvisioningState

		if err := res.FromResource(azObj); err != nil {
			return fmt.Errorf("error res.FromResource with: %w", err)
		}

		return nil
	}); err != nil {
		return ctrl.Result{}, fmt.Errorf("failed to patch after starting delete with: %w", err)
	}

	// delete has started, check back to seen when the finalizer can be removed
	return ctrl.Result{
		RequeueAfter: 5 * time.Second,
	}, nil
}

// updateFromNonTerminalDeleteState will call Azure to check if the resource still exists. If so, it will requeue, else,
// the finalizer will be removed.
func (gr *GenericReconciler) updateFromNonTerminalDeleteState(ctx context.Context, resourcer zips.Resourcer) (ctrl.Result, error) {
	azObj, err := resourcer.ToResource()
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("unable to transform to resource with: %w", err)
	}

	// already deleting, just check to see if it still exists and if it's gone, remove finalizer
	found, err := gr.Applier.HeadResource(ctx, azObj)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("failed to head resource with: %w", err)
	}

	if found {
		return ctrl.Result{RequeueAfter: 10 * time.Second}, nil
	}

	err = patcher(ctx, gr.Client, resourcer, func(res zips.Resourcer) error {
		controllerutil.RemoveFinalizer(res, apis.AzureInfraFinalizer)
		return nil
	})

	// patcher will try to fetch the object after patching, so ignore not found errors
	if apierrors.IsNotFound(err) {
		return ctrl.Result{}, nil
	}
	return ctrl.Result{}, err
}

// updatedFromNonTerminalApplyState will ask Azure for the updated status of the deployment. If the object is in a
// non terminal state, it will requeue, else, status will be updated.
func (gr *GenericReconciler) updateFromNonTerminalApplyState(ctx context.Context, resourcer zips.Resourcer) (ctrl.Result, error) {
	azObj, err := resourcer.ToResource()
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("unable to transform to resource with: %w", err)
	}

	if err := patcher(ctx, gr.Client, resourcer, func(res zips.Resourcer) error {
		// update with latest information about the apply
		azObj, err = gr.Applier.Apply(ctx, azObj)
		if err != nil {
			return fmt.Errorf("failed to apply state to Azure with %w", err)
		}

		if err := res.FromResource(azObj); err != nil {
			return fmt.Errorf("failed FromResource with: %w", err)
		}

		if err := addResourceHashAnnotation(res); err != nil {
			return fmt.Errorf("failed to addResourceHashAnnotation with: %w", err)
		}

		return nil
	}); err != nil {
		return ctrl.Result{}, fmt.Errorf("failed to patch with: %w", err)
	}

	result := ctrl.Result{}
	if !zips.IsTerminalProvisioningState(azObj.ProvisioningState) {
		result = ctrl.Result{
			RequeueAfter: 5 * time.Second,
		}
	}
	return result, err
}

// applySpecChange will apply the new spec state to an Azure resource. The resource should then enter
// into a non terminal state and will then be requeued for polling.
func (gr *GenericReconciler) applySpecChange(ctx context.Context, resourcer zips.Resourcer) (ctrl.Result, error) {
	azObj, err := resourcer.ToResource()
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("unable to transform to resource with: %w", err)
	}

	if err = patcher(ctx, gr.Client, resourcer, func(res zips.Resourcer) error {
		controllerutil.AddFinalizer(res, apis.AzureInfraFinalizer)

		azObj.ProvisioningState = ""
		azObj, err = gr.Applier.Apply(ctx, azObj)
		if err != nil {
			return fmt.Errorf("failed to apply state to Azure with %w", err)
		}

		if err := res.FromResource(azObj); err != nil {
			return err
		}

		if err := addResourceHashAnnotation(res); err != nil {
			return fmt.Errorf("failed to addResourceHashAnnotation with: %w", err)
		}

		return nil
	}); err != nil {
		return ctrl.Result{}, fmt.Errorf("failed to patch with: %w", err)
	}

	result := ctrl.Result{}
	if !zips.IsTerminalProvisioningState(azObj.ProvisioningState) {
		result = ctrl.Result{
			RequeueAfter: 5 * time.Second,
		}
	}
	return result, err
}

func hasResourceHashAnnotationChanged(resourcer zips.Resourcer) (bool, error) {
	oldSig, exists := resourcer.GetAnnotations()[ResourceSigAnnotationKey]
	if !exists {
		// signature does not exist, so yes, it has changed
		return true, nil
	}

	newSig, err := zips.SpecSignature(resourcer)
	if err != nil {
		return false, err
	}
	// check if the last signature matches the new signature
	return oldSig != newSig, nil
}

func addResourceHashAnnotation(resourcer zips.Resourcer) error {
	sig, err := zips.SpecSignature(resourcer)
	if err != nil {
		return err
	}

	annotations := resourcer.GetAnnotations()
	if annotations == nil {
		annotations = map[string]string{}
	}

	annotations[ResourceSigAnnotationKey] = sig
	resourcer.SetAnnotations(annotations)
	return nil
}

func patcher(ctx context.Context, c client.Client, resourcer zips.Resourcer, mutator func(zips.Resourcer) error) error {
	patchHelper, err := patch.NewHelper(resourcer, c)
	if err != nil {
		return err
	}

	if err := mutator(resourcer); err != nil {
		return err
	}

	if err := patchHelper.Patch(ctx, resourcer); err != nil {
		return err
	}

	// fill resourcer with patched updates since patch will copy resourcer
	return c.Get(ctx, client.ObjectKey{
		Namespace: resourcer.GetNamespace(),
		Name:      resourcer.GetName(),
	}, resourcer)
}
