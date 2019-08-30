/*
Copyright 2019 Alexander Eldeib.
*/

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Fetch retrieves an object by namespaced name from the API server and puts the contents in the runtime.Object parameter.
// TODO(ace): refactor onto base reconciler struct
func Fetch(ctx context.Context, client client.Client, namespacedName types.NamespacedName, obj runtime.Object, log logr.Logger) error {
	if err := client.Get(ctx, namespacedName, obj); err != nil {
		// dont't requeue not found
		if apierrs.IsNotFound(err) {
			return nil
		}
		log.Error(err, "unable to fetch secret")
		return err
	}
	return nil
}

// AddFinalizerAndUpdate removes a finalizer from a runtime object and attempts to update that object in the API server.
// It returns an error if either operation failed.
func AddFinalizerAndUpdate(ctx context.Context, client client.Client, finalizer string, o runtime.Object) error {
	m, err := meta.Accessor(o)
	if err != nil {
		return err
	}
	if contains(m.GetFinalizers(), finalizer) {
		return nil
	}
	AddFinalizer(m, finalizer)
	if err := client.Update(ctx, o); err != nil {
		return err
	}
	return nil
}

// RemoveFinalizerAndUpdate removes a finalizer from a runtime object and attempts to update that object in the API server.
// It returns an error if either operation failed.
func RemoveFinalizerAndUpdate(ctx context.Context, client client.Client, finalizer string, o runtime.Object) error {
	m, err := meta.Accessor(o)
	if err != nil {
		return err
	}
	if !contains(m.GetFinalizers(), finalizer) {
		return nil
	}
	RemoveFinalizer(m, finalizer)
	if err := client.Update(ctx, o); err != nil {
		return err
	}
	return nil
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

// AddFinalizerWithError tries to convert a runtime object to a metav1 object and add the provided finalizer.
// It returns an error if the provided object cannot provide an accessor.
func AddFinalizerWithError(o runtime.Object, finalizer string) error {
	m, err := meta.Accessor(o)
	if err != nil {
		return err
	}
	AddFinalizer(m, finalizer)
	return nil
}

// RemoveFinalizer accepts a metav1 object and removes the provided finalizer if present.
func RemoveFinalizer(o metav1.Object, finalizer string) {
	f := o.GetFinalizers()
	for i, e := range f {
		if e == finalizer {
			f = append(f[:i], f[i+1:]...)
			o.SetFinalizers(f)
			return
		}
	}
}

// HasFinalizer accepts a metav1 object and returns true if the the object has the provided finalizer.
func HasFinalizer(o metav1.Object, finalizer string) bool {
	f := o.GetFinalizers()
	for _, e := range f {
		if e == finalizer {
			return true
		}
	}
	return false
}

// RemoveFinalizerWithError tries to convert a runtime object to a metav1 object and remove the provided finalizer.
// It returns an error if the provided object cannot provide an accessor.
func RemoveFinalizerWithError(o runtime.Object, finalizer string) error {
	m, err := meta.Accessor(o)
	if err != nil {
		return err
	}
	RemoveFinalizer(m, finalizer)
	return nil
}

func DeleteIfFound(ctx context.Context, client client.Client, obj runtime.Object) error {
	if err := client.Delete(ctx, obj); err != nil && !apierrs.IsNotFound(err) {
		return err
	}
	return nil
}

func contains(vals []string, val string) bool {
	for _, v := range vals {
		if v == val {
			return true
		}
	}
	return false
}

// func remove(vals []string, val string) []string {
// 	for i, v := range vals {
// 		if v == val {
// 			return append(vals[:i], vals[i+1:]...)
// 		}
// 	}
// 	return vals
// }

func containsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func removeString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}
