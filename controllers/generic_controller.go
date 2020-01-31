/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers

import (
	"context"
	"fmt"
	"strings"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/Azure/k8s-infra/apis"
	microsoftresourcesv1 "github.com/Azure/k8s-infra/apis/microsoft.resources/v1"
	"github.com/Azure/k8s-infra/pkg/util/patch"
	"github.com/Azure/k8s-infra/pkg/zips"
)

// +kubebuilder:rbac:groups=microsoft.resources.infra.azure.com,resources=resourcegroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=microsoft.resources.infra.azure.com,resources=resourcegroups/status,verbs=get;update;patch

// KnownTypes defines an array of runtime.Objects to be reconciled, where each
// object in the array will generate a controller. If the concrete type
// implements the owner interface, the generated controller will inject the
// owned types supplied by the Owns() method of the CRD type. Each controller
// may directly reconcile a single object, but may indirectly watch
// and reconcile many Owned objects. The singular type is necessary to generically
// produce a reconcile function aware of concrete types, as a closure.
var KnownTypes = []runtime.Object{
	new(microsoftresourcesv1.ResourceGroup),
}

var (
	ctrlog = ctrl.Log.WithName("resources-controller")
)

type owner interface {
	Owns() []runtime.Object
}

func RegisterAll(mgr ctrl.Manager, applier zips.Applier, objs []runtime.Object) []error {
	var errs []error
	for _, obj := range objs {
		mgr := mgr
		applier := applier
		obj := obj
		if err := register(mgr, applier, obj); err != nil {
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
func register(mgr ctrl.Manager, applier zips.Applier, obj runtime.Object) error {
	controller := ctrl.NewControllerManagedBy(mgr).For(obj)
	ownerType, ok := obj.(owner)
	if ok {
		for _, ownedType := range ownerType.Owns() {
			controller.Owns(ownedType)
		}
	}

	v, err := conversion.EnforcePtr(obj)
	if err != nil {
		return err
	}

	t := v.Type()
	controller.Named(strings.ReplaceAll(t.String(), ".", "_"))
	reconciler := getReconciler(obj, mgr, mgr.GetClient(), applier)
	return controller.Complete(reconciler)
}

// getReconciler is a simple helper to directly produce a reconcile Function. Useful to test against.
func getReconciler(obj runtime.Object, mgr ctrl.Manager, kubeclient client.Client, applier zips.Applier) reconcile.Func {
	return func(req ctrl.Request) (ctrl.Result, error) {
		// Use the provided GVK to construct a new runtime object of the desired concrete type.
		gvk, err := apiutil.GVKForObject(obj, mgr.GetScheme())
		if err != nil {
			return ctrl.Result{}, err
		}

		obj, err = mgr.GetScheme().New(gvk)
		if err != nil {
			return ctrl.Result{}, err
		}

		return reconcileFn(req, obj, kubeclient, applier)
	}
}

// reconcileFn is the "real" reconciler function, created as a closure above at manager startup to have access to the gvk.
func reconcileFn(req ctrl.Request, obj runtime.Object, kubeclient client.Client, applier zips.Applier) (_ ctrl.Result, reterr error) {
	ctx := context.Background()

	if err := kubeclient.Get(ctx, req.NamespacedName, obj); err != nil {
		if apierrors.IsNotFound(err) {
			ctrlog.Info("object not found, requeue")
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return ctrl.Result{}, nil
		}

		ctrlog.Error(err, "error reading object")
		// Error reading the object - requeue the request.
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
		return ctrl.Result{}, reconcileDelete(ctx, kubeclient, applier, resourcer)
	}

	return ctrl.Result{}, reconcileApply(ctx, kubeclient, applier, resourcer)
}

func reconcileApply(ctx context.Context, c client.Client, applier zips.Applier, resourcer zips.Resourcer) error {
	if err := patcher(ctx, c, resourcer, func(res zips.Resourcer) error {
		azobj := res.ToResource()
		azobj.ProvisioningState = "Applying"
		res.FromResource(azobj)
		return nil
	}); err != nil {
		return err
	}

	// apply normally;
	//
	// This should call to the Applier where the current state of the applied resource should be compared to the cached
	// state of the Azure resource. If the two states differ, the Applier should then apply that state to Azure.

	return patcher(ctx, c, resourcer, func(res zips.Resourcer) error {
		controllerutil.AddFinalizer(resourcer, apis.AzureInfraFinalizer)
		resource := res.ToResource()
		appliedResource, err := applier.Apply(ctx, resource)
		if err != nil {
			return fmt.Errorf("failed to apply state to Azure with %w", err)
		}

		// TODO (dj): this should be set from the provisioning state of the resource in Azure. For now, I think we can assume this is "Succeeded" if no error
		appliedResource.ProvisioningState = "Succeeded"
		res.FromResource(appliedResource)
		return nil
	})
}

func reconcileDelete(ctx context.Context, c client.Client, applier zips.Applier, resourcer zips.Resourcer) error {
	if err := patcher(ctx, c, resourcer, func(res zips.Resourcer) error {
		azobj := res.ToResource()
		azobj.ProvisioningState = "Deleting"
		res.FromResource(azobj)
		return nil
	}); err != nil {
		return err
	}

	err := patcher(ctx, c, resourcer, func(res zips.Resourcer) error {
		if err := applier.Delete(ctx, res.ToResource()); err != nil {
			return fmt.Errorf("failed trying to delete with %w", err)
		}

		controllerutil.RemoveFinalizer(res, apis.AzureInfraFinalizer)
		return nil
	})

	// patcher will try to fetch the object after patching, so ignore not found errors
	if apierrors.IsNotFound(err) {
		return nil
	}
	return err
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
