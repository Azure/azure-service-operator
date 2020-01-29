/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	microsoftresourcesv1 "github.com/Azure/k8s-infra/apis/microsoft.resources/v1"
	"github.com/Azure/k8s-infra/pkg/zips"
)

const (
	finalizerName string = "infra.azure.com/finalizer"
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
func reconcileFn(req ctrl.Request, obj runtime.Object, kubeclient client.Client, applier zips.Applier) (ctrl.Result, error) {
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

	// Get the Azure resource representation of the Kubernetes object
	azObj := resourcer.ToResource()
	bits, err := json.Marshal(azObj)
	if err != nil {
		ctrlog.Error(err, "Failed marshalling")
		return ctrl.Result{}, err
	}

	ctrlog.Info("###### Processing: " + string(bits))

	// TODO(ace): consider a union of runtime.Object and metav1.Object
	// ref: https://github.com/kubernetes-sigs/controller-runtime/issues/594
	metaObj, err := meta.Accessor(obj)
	if err != nil {
		ctrlog.Error(err, "Failed at meta.Accessor")
		return ctrl.Result{}, err
	}

	// reconcile delete
	if !metaObj.GetDeletionTimestamp().IsZero() {
		if HasFinalizer(metaObj, finalizerName) {
			azObj.ProvisioningState = "Deleting"
			obj = resourcer.FromResource(azObj)
			if err := kubeclient.Status().Update(ctx, obj); err != nil {
				return ctrl.Result{}, fmt.Errorf("failed updating provisioning status to Deleting with %w", err)
			}

			if err := applier.Delete(ctx, azObj); err != nil {
				err = fmt.Errorf("failed trying to delete %q with %w", req.NamespacedName, err)
				ctrlog.Error(err, "Delete error")
				return ctrl.Result{}, err
			}

			obj = resourcer.FromResource(azObj)
			RemoveFinalizer(metaObj, finalizerName)
			if err := kubeclient.Update(ctx, obj); err != nil {
				return ctrl.Result{}, fmt.Errorf("failed trying to remove finalizer with %w", err)
			}

		}
		return ctrl.Result{}, nil
	}

	// add finalizer if not present
	if !HasFinalizer(metaObj, finalizerName) {
		AddFinalizer(metaObj, finalizerName)
		if err := kubeclient.Update(ctx, obj); err != nil {
			ctrlog.Error(err, fmt.Sprintf("Failed trying to update the resource with a finalizer: %+v ### %+v", metaObj, obj))
			return ctrl.Result{}, err
		}
	}

	azObj.ProvisioningState = "Applying"
	obj = resourcer.FromResource(azObj)
	if err := kubeclient.Status().Update(ctx, obj); err != nil {
		return ctrl.Result{}, fmt.Errorf("failed updating provisioning status to Applying with %w", err)
	}

	// apply normally;
	//
	// This should call to the Applier where the current state of the applied resource should be compared to the cached
	// state of the Azure resource. If the two states differ, the Applier should then apply that state to Azure.
	res, err := applier.Apply(ctx, azObj)
	if err != nil {
		// failed applying the result
		ctrlog.Error(err, "Failed applying the result")
		return ctrl.Result{}, err
	}

	obj = resourcer.FromResource(res)

	if err := kubeclient.Status().Update(ctx, obj); err != nil {
		return ctrl.Result{}, fmt.Errorf("failed updating final status on resource with %w", err)
	}

	return ctrl.Result{}, nil
}

// AddFinalizer accepts a metav1 object and adds the provided finalizer if not present.
func AddFinalizer(o metav1.Object, finalizer string) {
	f := o.GetFinalizers()
	for _, e := range f {
		if e == finalizer {
			return
		}
	}
	o.SetFinalizers(append(f, finalizer))
}

// RemoveFinalizer accepts a metav1 object and removes the finalizer
func RemoveFinalizer(o metav1.Object, finalizer string) {
	f := o.GetFinalizers()
	var finalizers []string
	for _, e := range f {
		if e != finalizer {
			finalizers = append(finalizers, f...)
		}
	}
	o.SetFinalizers(finalizers)
}

// HasFinalizer accepts a metav1 object and returns true if the the object has the provided finalizer.
func HasFinalizer(o metav1.Object, finalizer string) bool {
	return ContainsString(o.GetFinalizers(), finalizer)
}

func ContainsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}
