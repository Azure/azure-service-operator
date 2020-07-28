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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
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
	azcorev1 "github.com/Azure/k8s-infra/apis/core/v1"
	microsoftcomputev1 "github.com/Azure/k8s-infra/apis/microsoft.compute/v1"
	microsoftnetworkv1 "github.com/Azure/k8s-infra/apis/microsoft.network/v1"
	microsoftresourcesv1 "github.com/Azure/k8s-infra/apis/microsoft.resources/v1"
	"github.com/Azure/k8s-infra/pkg/util/patch"
	"github.com/Azure/k8s-infra/pkg/util/statusutil"
	"github.com/Azure/k8s-infra/pkg/xform"
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
		new(microsoftnetworkv1.BackendAddressPool),
		new(microsoftnetworkv1.FrontendIPConfiguration),
		new(microsoftnetworkv1.InboundNatRule),
		new(microsoftnetworkv1.LoadBalancer),
		new(microsoftnetworkv1.LoadBalancingRule),
		new(microsoftnetworkv1.NetworkSecurityGroup),
		new(microsoftnetworkv1.OutboundRule),
		new(microsoftnetworkv1.Route),
		new(microsoftnetworkv1.RouteTable),
		new(microsoftnetworkv1.SecurityRule),
		new(microsoftnetworkv1.VirtualNetwork),
		new(microsoftnetworkv1.Subnet),
		new(microsoftnetworkv1.NetworkInterface),
		new(microsoftcomputev1.VirtualMachine),
		new(microsoftcomputev1.VirtualMachineScaleSet),
	}
)

// +kubebuilder:rbac:groups=core,resources=events,verbs=get;list;watch;create;patch
// +kubebuilder:rbac:groups=microsoft.resources.infra.azure.com,resources=resourcegroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=microsoft.resources.infra.azure.com,resources=resourcegroups/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=microsoft.network.infra.azure.com,resources=backendaddresspools;frontendipconfigurations;inboundnatrules;loadbalancers;loadbalancingrules;networkinterfaces;networksecuritygroups;outboundrules;routes;routetables;securityrules;subnets;virtualnetworks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=microsoft.network.infra.azure.com,resources=backendaddresspools/status;frontendipconfigurations/status;inboundnatrules/status;loadbalancers/status;loadbalancingrules/status;networkinterfaces/status;networksecuritygroups/status;outboundrules/status;routes/status;routetables/status;securityrules/status;subnets/status;virtualnetworks/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=microsoft.compute.infra.azure.com,resources=virtualmachines;virtualmachinescalesets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=microsoft.compute.infra.azure.com,resources=virtualmachines/status;virtualmachinescalesets/status,verbs=get;update;patch

type (
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
		Converter  *xform.ARMConverter
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

	if err := mgr.GetFieldIndexer().IndexField(obj, "status.id", func(obj runtime.Object) []string {
		unObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
		if err != nil {
			return []string{}
		}

		id, ok, err := unstructured.NestedString(unObj, "status", "id")
		if err != nil || !ok {
			return []string{}
		}

		return []string{id}
	}); err != nil {
		return fmt.Errorf("unable to setup field indexer for status.id of %v with: %w", gvk, err)
	}

	reconciler := &GenericReconciler{
		Client:    mgr.GetClient(),
		Applier:   applier,
		Scheme:    mgr.GetScheme(),
		Name:      t.Name(),
		Log:       log.WithName(controllerName),
		Recorder:  mgr.GetEventRecorderFor(controllerName),
		GVK:       gvk,
		Converter: xform.NewARMConverter(mgr.GetClient(), mgr.GetScheme()),
	}

	ctrlBuilder := ctrl.NewControllerManagedBy(mgr).
		For(obj).
		WithOptions(options)

	if metaObj, ok := obj.(azcorev1.MetaObject); ok {
		trls, err := xform.GetTypeReferenceData(metaObj)
		if err != nil {
			return fmt.Errorf("unable get type reference data for obj %v with: %w", metaObj, err)
		}

		for _, trl := range trls {
			if trl.IsOwned {
				gvk := schema.GroupVersionKind{
					Group:   trl.Group,
					Version: "v1",
					Kind:    trl.Kind,
				}
				ownedObj, err := mgr.GetScheme().New(gvk)
				if err != nil {
					return fmt.Errorf("unable to create GVK %v with: %w", gvk, err)
				}

				reconciler.Owns = append(reconciler.Owns, ownedObj)
				ctrlBuilder.Owns(ownedObj)
			}
		}
	}

	c, err := ctrlBuilder.Build(reconciler)
	if err != nil {
		return fmt.Errorf("unable to build controller / reconciler with: %w", err)
	}

	reconciler.Controller = c

	return ctrl.NewWebhookManagedBy(mgr).
		For(obj).
		Complete()
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
	metaObj, ok := obj.(azcorev1.MetaObject)
	if !ok {
		return ctrl.Result{}, fmt.Errorf("object is not a azcorev1.MetaObject: %+v", obj)
	}

	// reconcile delete
	if !metaObj.GetDeletionTimestamp().IsZero() {
		log.Info("reconcile delete start")
		result, err := gr.reconcileDelete(ctx, metaObj)
		if err != nil {
			gr.Recorder.Event(metaObj, v1.EventTypeWarning, "ReconcileDeleteError", err.Error())
			log.Error(err, "reconcile delete error")
			return result, err
		}

		log.Info("reconcile delete complete")
		return result, err
	}

	if grouped, ok := obj.(azcorev1.Grouped); ok {
		ready, err := gr.isResourceGroupReady(ctx, grouped)
		if err != nil {
			log.Error(err, "failed checking if resource group was ready")
			gr.Recorder.Event(metaObj, v1.EventTypeWarning, "GroupReadyError", fmt.Sprintf("isResourceGroupReady failed with: %s", err))
			return ctrl.Result{}, err
		}

		if !ready {
			requeueTime := 30 * time.Second
			msg := fmt.Sprintf("resource group %q is not ready or not created yet; will try again in about %s", grouped.GetResourceGroupObjectRef().Name, requeueTime)
			gr.Recorder.Event(metaObj, v1.EventTypeNormal, "ResourceGroupNotReady", msg)
			return ctrl.Result{
				RequeueAfter: requeueTime,
			}, nil
		}
	}

	ownersReady, err := gr.Converter.AreOwnersReady(ctx, metaObj)
	if err != nil {
		log.Error(err, "failed checking owner references are ready")
		gr.Recorder.Event(metaObj, v1.EventTypeWarning, "OwnerReferenceCheckError", err.Error())
		return ctrl.Result{}, err
	}

	if !ownersReady {
		gr.Recorder.Event(metaObj, v1.EventTypeNormal, "OwnerReferencesNotReady", "owner reference are not ready; retrying in 30s")
		return ctrl.Result{
			RequeueAfter: 30 * time.Second,
		}, nil
	}

	log.Info("reconcile apply start")
	result, err := gr.reconcileApply(ctx, metaObj, log)
	if err != nil {
		log.Error(err, "reconcile apply error")
		gr.Recorder.Event(metaObj, v1.EventTypeWarning, "ReconcileError", err.Error())
		return result, err
	}

	allApplied, err := gr.Converter.ApplyOwnership(ctx, metaObj)
	if err != nil {
		log.Error(err, "failed applying ownership to owned references")
		gr.Recorder.Event(metaObj, v1.EventTypeWarning, "OwnerReferencesFailedApply", "owner reference are not ready; retrying in 30s")
		return ctrl.Result{}, fmt.Errorf("failed applying ownership to owned references with: %w", err)
	}

	if !allApplied {
		log.Info("not all owned objects were applied; will requeue for 30 seconds")
		result.RequeueAfter = 30 * time.Second
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
func (gr *GenericReconciler) reconcileApply(ctx context.Context, metaObj azcorev1.MetaObject, log logr.Logger) (ctrl.Result, error) {
	// check if the hash on the resource has changed
	hasChanged, err := hasResourceHashAnnotationChanged(metaObj)
	if err != nil {
		err = fmt.Errorf("failed comparing resource hash with: %w", err)
		gr.Recorder.Event(metaObj, v1.EventTypeWarning, "AnnotationError", err.Error())
		return ctrl.Result{}, err
	}

	resource, err := gr.Converter.ToResource(ctx, metaObj)
	if err != nil {
		err := fmt.Errorf("unable to transform to resource with: %w", err)
		gr.Recorder.Event(metaObj, v1.EventTypeWarning, "ToResourceError", err.Error())
		return ctrl.Result{}, err
	}

	// if the resource hash (spec) has not changed, don't apply again
	if !hasChanged && zips.IsTerminalProvisioningState(resource.ProvisioningState) {
		msg := fmt.Sprintf("resource in state %q and spec has not changed", resource.ProvisioningState)
		gr.Recorder.Event(metaObj, v1.EventTypeNormal, "ResourceHasNotChanged", msg)
		return ctrl.Result{}, nil
	}

	switch {
	case hasChanged:
		msg := fmt.Sprintf("resource in state %q has changed and spec will be applied to Azure", resource.ProvisioningState)
		gr.Recorder.Event(metaObj, v1.EventTypeNormal, "ResourceHasChanged", msg)
		return gr.applySpecChange(ctx, metaObj)
	case !zips.IsTerminalProvisioningState(resource.ProvisioningState):
		msg := fmt.Sprintf("resource in state %q, asking Azure for updated state", resource.ProvisioningState)
		gr.Recorder.Event(metaObj, v1.EventTypeNormal, "ResourceStateNonTerminal", msg)
		return gr.updateFromNonTerminalApplyState(ctx, metaObj, log)
	default:
		msg := fmt.Sprintf("resource in state %q and spec has not changed", resource.ProvisioningState)
		gr.Recorder.Event(metaObj, v1.EventTypeNormal, "ResourceNoopReconcile", msg)
		return ctrl.Result{}, nil // all is good and there are no changes to deal with
	}
}

func (gr *GenericReconciler) isResourceGroupReady(ctx context.Context, grouped azcorev1.Grouped) (bool, error) {
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

	status, err := statusutil.GetProvisioningState(&rg)
	if err != nil {
		return false, fmt.Errorf("unable to get provisioning state for resource group with: %w", err)
	}

	return zips.ProvisioningState(status) == zips.SucceededProvisioningState, nil
}

// reconcileDelete will begin and follow a delete operation of a resource. The finalizer will only be removed upon the
// resource actually being deleted in Azure.
//
// There are 2 possible state transitions.
// *  obj.ProvisioningState == \*\ --> Start deleting in Azure and mark state as "Deleting"
// *  obj.ProvisioningState == "Deleting" --> http HEAD to see if resource still exists in Azure. If so, requeue, else, remove finalizer.
func (gr *GenericReconciler) reconcileDelete(ctx context.Context, metaObj azcorev1.MetaObject) (ctrl.Result, error) {
	resource, err := gr.Converter.ToResource(ctx, metaObj)
	// if error IsOwnerNotFound, then carry on. Perhaps, the owner has already been deleted.
	if err != nil && !xform.IsOwnerNotFound(err) {
		return ctrl.Result{}, fmt.Errorf("unable to transform to resource with: %w", err)
	}

	switch resource.ProvisioningState {
	case zips.DeletingProvisioningState:
		msg := "deleting... checking for updated state"
		gr.Recorder.Event(metaObj, v1.EventTypeNormal, "ResourceDeleteInProgress", msg)
		return gr.updateFromNonTerminalDeleteState(ctx, resource, metaObj)
	default:
		msg := fmt.Sprintf("start deleting resource in state %q", resource.ProvisioningState)
		gr.Recorder.Event(metaObj, v1.EventTypeNormal, "ResourceDeleteStart", msg)
		return gr.startDeleteOfResource(ctx, resource, metaObj)
	}
}

// startDeleteOfResource will begin the delete of a resource by telling Azure to start deleting it. The resource will be
// marked with the provisioning state of "Deleting".
func (gr *GenericReconciler) startDeleteOfResource(ctx context.Context, resource *zips.Resource, metaObj azcorev1.MetaObject) (ctrl.Result, error) {
	if err := patcher(ctx, gr.Client, metaObj, func(mutMetaObject azcorev1.MetaObject) error {
		if resource.ID != "" {
			if _, err := gr.Applier.BeginDelete(ctx, resource); err != nil {
				return fmt.Errorf("failed trying to delete with %w", err)
			}

			resource.ProvisioningState = zips.DeletingProvisioningState
		} else {
			controllerutil.RemoveFinalizer(mutMetaObject, apis.AzureInfraFinalizer)
		}

		if err := gr.Converter.FromResource(resource, mutMetaObject); err != nil {
			return fmt.Errorf("error gr.Converter.FromResource with: %w", err)
		}

		return nil
	}); err != nil && !apierrors.IsNotFound(err) {
		return ctrl.Result{}, fmt.Errorf("failed to patch after starting delete with: %w", err)
	}

	// delete has started, check back to seen when the finalizer can be removed
	return ctrl.Result{
		RequeueAfter: 5 * time.Second,
	}, nil
}

// updateFromNonTerminalDeleteState will call Azure to check if the resource still exists. If so, it will requeue, else,
// the finalizer will be removed.
func (gr *GenericReconciler) updateFromNonTerminalDeleteState(ctx context.Context, resource *zips.Resource, metaObj azcorev1.MetaObject) (ctrl.Result, error) {
	// already deleting, just check to see if it still exists and if it's gone, remove finalizer
	found, err := gr.Applier.HeadResource(ctx, resource)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("failed to head resource with: %w", err)
	}

	if found {
		return ctrl.Result{RequeueAfter: 10 * time.Second}, nil
	}

	err = patcher(ctx, gr.Client, metaObj, func(mutMetaObject azcorev1.MetaObject) error {
		controllerutil.RemoveFinalizer(mutMetaObject, apis.AzureInfraFinalizer)
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
func (gr *GenericReconciler) updateFromNonTerminalApplyState(ctx context.Context, metaObj azcorev1.MetaObject, log logr.Logger) (ctrl.Result, error) {
	resource, err := gr.Converter.ToResource(ctx, metaObj)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("unable to transform to resource with: %w", err)
	}

	if err := patcher(ctx, gr.Client, metaObj, func(mutMetaObj azcorev1.MetaObject) error {
		// update with latest information about the apply
		resource, err = gr.Applier.Apply(ctx, resource)
		if err != nil {
			return fmt.Errorf("failed to apply state to Azure with %w", err)
		}

		if err := gr.Converter.FromResource(resource, mutMetaObj); err != nil {
			return fmt.Errorf("failed FromResource with: %w", err)
		}

		if err := addResourceHashAnnotation(mutMetaObj); err != nil {
			return fmt.Errorf("failed to addResourceHashAnnotation with: %w", err)
		}

		return nil
	}); err != nil {
		return ctrl.Result{}, fmt.Errorf("failed to patch with: %w", err)
	}

	result := ctrl.Result{}
	if !zips.IsTerminalProvisioningState(resource.ProvisioningState) {
		log.Info("requeuing in 20 seconds", "res.ID", resource.ID, "res.State", resource.ProvisioningState, "res.deploymentID", resource.DeploymentID, "metaObj", metaObj)
		result = ctrl.Result{
			RequeueAfter: 20 * time.Second,
		}
	}
	return result, err
}

// applySpecChange will apply the new spec state to an Azure resource. The resource should then enter
// into a non terminal state and will then be requeued for polling.
func (gr *GenericReconciler) applySpecChange(ctx context.Context, metaObj azcorev1.MetaObject) (ctrl.Result, error) {
	resource, err := gr.Converter.ToResource(ctx, metaObj)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("unable to transform to resource with: %w", err)
	}

	if err := patcher(ctx, gr.Client, metaObj, func(mutObj azcorev1.MetaObject) error {
		controllerutil.AddFinalizer(mutObj, apis.AzureInfraFinalizer)
		resource.ProvisioningState = ""
		resource, err = gr.Applier.Apply(ctx, resource)
		if err != nil {
			return fmt.Errorf("failed to apply state to Azure with %w", err)
		}

		if err := gr.Converter.FromResource(resource, mutObj); err != nil {
			return err
		}

		if err := addResourceHashAnnotation(mutObj); err != nil {
			return fmt.Errorf("failed to addResourceHashAnnotation with: %w", err)
		}

		return nil
	}); err != nil {
		return ctrl.Result{}, fmt.Errorf("failed to patch with: %w", err)
	}

	result := ctrl.Result{}
	if !zips.IsTerminalProvisioningState(resource.ProvisioningState) {
		result = ctrl.Result{
			RequeueAfter: 5 * time.Second,
		}
	}
	return result, err
}

func hasResourceHashAnnotationChanged(metaObj azcorev1.MetaObject) (bool, error) {
	oldSig, exists := metaObj.GetAnnotations()[ResourceSigAnnotationKey]
	if !exists {
		// signature does not exist, so yes, it has changed
		return true, nil
	}

	newSig, err := azcorev1.SpecSignature(metaObj)
	if err != nil {
		return false, err
	}
	// check if the last signature matches the new signature
	return oldSig != newSig, nil
}

func addResourceHashAnnotation(metaObj azcorev1.MetaObject) error {
	sig, err := azcorev1.SpecSignature(metaObj)
	if err != nil {
		return err
	}

	annotations := metaObj.GetAnnotations()
	if annotations == nil {
		annotations = map[string]string{}
	}

	annotations[ResourceSigAnnotationKey] = sig
	metaObj.SetAnnotations(annotations)
	return nil
}

func patcher(ctx context.Context, c client.Client, metaObj azcorev1.MetaObject, mutator func(azcorev1.MetaObject) error) error {
	patchHelper, err := patch.NewHelper(metaObj, c)
	if err != nil {
		return err
	}

	if err := mutator(metaObj); err != nil {
		return err
	}

	if err := patchHelper.Patch(ctx, metaObj); err != nil {
		return err
	}

	// fill resourcer with patched updates since patch will copy resourcer
	return c.Get(ctx, client.ObjectKey{
		Namespace: metaObj.GetNamespace(),
		Name:      metaObj.GetName(),
	}, metaObj)
}
