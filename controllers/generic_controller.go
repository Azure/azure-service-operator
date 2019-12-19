/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	protov1alpha1 "github.com/Azure/k8s-infra/api/v1alpha1"
	"github.com/Azure/k8s-infra/pkg/zips"
)

// GenericReconciler reconciles any generic Azure object
type (
	GenericReconciler struct {
		client.Client
		Log     logr.Logger
		Scheme  *runtime.Scheme
		Applier zips.Applier
	}
)

const (
	finalizerName string = "infra.azure.com/finalizer"
)

// +kubebuilder:rbac:groups=proto.infra.azure.com,resources=virtualnetwork;resourcegroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=proto.infra.azure.com,resources=virtualnetwork/status;resourcegroups/status,verbs=get;update;patch
func (gr *GenericReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.TODO()
	_ = gr.Log.WithValues("genericReconciler", req.NamespacedName)

	// start from unstructured returning both the unstructured object and the strongly typed unmarshalled object
	unstruct, runObj, err := gr.getUnstructured(ctx, req)
	if err != nil {
		return ctrl.Result{}, err
	}

	// cast into a type we can use to apply or delete
	var blank interface{} = runObj
	resourcer, ok := blank.(zips.Resourcer)
	if !ok {
		return ctrl.Result{}, fmt.Errorf("object: %+v is not a zips.Resourcer", runObj)
	}

	// get the resource representation of the K8s object
	zRes, err := resourcer.ToResource()
	if err != nil {
		return ctrl.Result{}, err
	}

	// delete me...
	//
	// This would call the applier to delete the object. This should probably be in a finalizer, but it is just used
	// here for illustrative purposes.
	if unstruct.GetDeletionTimestamp().IsZero() {
		if !HasFinalizer(unstruct, finalizerName) {
			AddFinalizer(unstruct, finalizerName)
			return ctrl.Result{}, gr.Update(ctx, unstruct)
		}
	} else {
		if HasFinalizer(unstruct, finalizerName) {
			// delete me
			if err := gr.Applier.Delete(ctx, zRes.ID); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}

	// apply normally;
	//
	// This should call to the Applier where the current state of the applied resource should be compared to the cached
	// state of the Azure resource. If the two states differ, the Applier should then apply that state to Azure.
	if err := gr.Applier.Apply(ctx, zRes); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (gr *GenericReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(new(protov1alpha1.VirtualNetwork)).
		For(new(protov1alpha1.ResourceGroup)).
		Complete(gr)
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

// getUnstructured is a _terrible_ method and I don't like it at all. I would like to know the specific GVK here, so
// I could do a GET efficiently. I don't know a better way to do this right now and would love to hear suggestions.
//
// One possibility would be to wrap all infrastructure objects with a outer wrapper which contains a reference, perhaps
// a raw extension, to a registered GVK. That way, we could unmarshal the wrapper, then do a generic unmarshal of the
// internal object.
//
// This is just here to start a conversation about how to better model this.
func (gr *GenericReconciler) getUnstructured(ctx context.Context, req ctrl.Request) (*unstructured.Unstructured, runtime.Object, error) {
	var lastErr error
	for _, gkv := range gr.searchGVKs() {
		var obj unstructured.Unstructured
		obj.SetGroupVersionKind(gkv)
		if err := gr.Get(ctx, req.NamespacedName, &obj); err != nil {
			lastErr = err
			continue
		}

		val := reflect.New(gr.Scheme.AllKnownTypes()[gkv])
		iface := val.Interface()
		bits, err := obj.MarshalJSON()
		if err != nil {
			return &obj, nil, err
		}

		if err := json.Unmarshal(bits, iface); err != nil {
			return &obj, nil, err
		}

		return &obj, iface.(runtime.Object), nil
	}
	return nil, nil, lastErr
}

func (gr *GenericReconciler) searchGVKs() []schema.GroupVersionKind {
	return []schema.GroupVersionKind{
		protov1alpha1.GroupVersion.WithKind("ResourceGroup"),
		protov1alpha1.GroupVersion.WithKind("VirtualNetwork"),
	}
}

