/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	protov1alpha1 "github.com/Azure/k8s-infra/api/v1alpha1"
	"github.com/Azure/k8s-infra/pkg/zips"
)

const (
	finalizerName string = "infra.azure.com/finalizer"
)

// +kubebuilder:rbac:groups=proto.infra.azure.com,resources=virtualnetwork;resourcegroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=proto.infra.azure.com,resources=virtualnetwork/status;resourcegroups/status,verbs=get;update;patch

// KnownTypes defines an array of runtime.Objects to be reconciled, where each
// object in the array will generate a controller. If the concrete type
// implements the owner interface, the generated controller will inject the
// owned types supplied by the Owns() method of the CRD type. Each controller
// may directly reconcile a single object, but may indirectly watch
// and reconcile many Owned objects. The singular type is necessary to generically
// produce a reconcile function aware of concrete types, as a closure.
var KnownTypes = []runtime.Object{
	&protov1alpha1.ResourceGroup{},
	&protov1alpha1.VirtualNetwork{},
}

type owner interface {
	Owns() []runtime.Object
}

func RegisterAll(mgr ctrl.Manager, applier zips.Applier, objs []runtime.Object) []error {
	var errs []error
	for _, obj := range objs {
		if err := register(mgr, applier, obj); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

// register takes a manager and a struct describing how to instantiate
// controllers for various types usng a generic reconciler function. Only one
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
	gvk := obj.GetObjectKind().GroupVersionKind()
	reconciler := getReconciler(gvk, mgr.GetScheme(), mgr.GetClient(), applier)
	return controller.Complete(reconciler)
}

// getReconciler is a simple helper to directly produce a reconcile Function. Useful to test against.
func getReconciler(gvk schema.GroupVersionKind, scheme *runtime.Scheme, kubeclient client.Client, applier zips.Applier) reconcile.Func {
	return func(req ctrl.Request) (ctrl.Result, error) {
		return reconcileFn(req, gvk, scheme, kubeclient, applier)
	}
}

// reconcileFn is the "real" reconciler function, created as a closure above at manager startup to have access to the gvk.
func reconcileFn(req ctrl.Request, gvk schema.GroupVersionKind, scheme *runtime.Scheme, kubeclient client.Client, applier zips.Applier) (ctrl.Result, error) {
	ctx := context.Background()

	// Use the provided GVK to construct a new runtime object of the desired concrete type.
	runObj, err := scheme.New(gvk)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Fetch the object by key + type.
	if err := kubeclient.Get(ctx, req.NamespacedName, runObj); err != nil {
		return ctrl.Result{}, err
	}

	// The Go type for the Kubernetes object must understand how to
	// convert itself to/from the corresponding Azure types.
	resourcer, ok := runObj.(zips.Resourcer)
	if !ok {
		return ctrl.Result{}, fmt.Errorf("object: %+v is not a zips.Resourcer", runObj)
	}

	// Get the Azure resource representation of the Kubernetes object
	azureObj, err := resourcer.ToResource()
	if err != nil {
		return ctrl.Result{}, err
	}

	// TODO(ace): consider a union of runtime.Object and metav1.Object
	// ref: https://github.com/kubernetes-sigs/controller-runtime/issues/594
	metaObj, err := meta.Accessor(runObj)
	if err != nil {
		return ctrl.Result{}, err
	}

	if metaObj.GetDeletionTimestamp().IsZero() {
		if !HasFinalizer(metaObj, finalizerName) {
			AddFinalizer(metaObj, finalizerName)
			return ctrl.Result{}, kubeclient.Update(ctx, runObj)
		}
	} else {
		if HasFinalizer(metaObj, finalizerName) {
			if err := applier.Delete(ctx, azureObj.ID); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	// apply normally;
	//
	// This should call to the Applier where the current state of the applied resource should be compared to the cached
	// state of the Azure resource. If the two states differ, the Applier should then apply that state to Azure.
	if _, err = applier.Apply(ctx, azureObj); err != nil {
		// failed applying the result
		return ctrl.Result{}, err
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
