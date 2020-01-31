/*
Copyright 2019 Alexander Eldeib.
*/

package controllers

import (
	"context"
	"reflect"
	"strings"
	"testing"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/go-logr/logr"
	"github.com/stretchr/testify/assert"
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
		log.Error(err, "unable to fetch object")
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
	if hasString(m.GetFinalizers(), finalizer) {
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
	if !hasString(m.GetFinalizers(), finalizer) {
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
	newFinalizers := addString(o.GetFinalizers(), finalizer)
	o.SetFinalizers(newFinalizers)
}

// AddFinalizerIfPossible tries to convert a runtime object to a metav1 object and add the provided finalizer.
// It returns an error if the provided object cannot provide an accessor.
func AddFinalizerIfPossible(o runtime.Object, finalizer string) error {
	m, err := meta.Accessor(o)
	if err != nil {
		return err
	}
	AddFinalizer(m, finalizer)
	return nil
}

// RemoveFinalizer accepts a metav1 object and removes the provided finalizer if present.
func RemoveFinalizer(o metav1.Object, finalizer string) {
	newFinalizers := removeString(o.GetFinalizers(), finalizer)
	o.SetFinalizers(newFinalizers)
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

// RemoveFinalizerIfPossible tries to convert a runtime object to a metav1 object and remove the provided finalizer.
// It returns an error if the provided object cannot provide an accessor.
func RemoveFinalizerIfPossible(o runtime.Object, finalizer string) error {
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

// hasString returns true if a given slice has the provided string s.
func hasString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

// addString returns a []string with s appended if it is not already found in the provided slice.
func addString(slice []string, s string) []string {
	for _, item := range slice {
		if item == s {
			return slice
		}
	}
	return append(slice, s)
}

// removeString returns a newly created []string that contains all items from slice that are not equal to s.
func removeString(slice []string, s string) []string {
	new := make([]string, 0)
	for _, item := range slice {
		if item == s {
			continue
		}
		new = append(new, item)
	}
	return new
}

// EnsureInstance creates the instance and waits for it to exist or timeout
func EnsureInstance(ctx context.Context, t *testing.T, tc testContext, instance runtime.Object) {
	assert := assert.New(t)
	err := tc.k8sClient.Create(ctx, instance)
	assert.Equal(nil, err, "create server2 in k8s")

	res, err := meta.Accessor(instance)
	assert.Equal(nil, err, "not a metav1 object")

	names := types.NamespacedName{Name: res.GetName(), Namespace: res.GetNamespace()}

	// Wait for first sql server to resolve
	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, names, instance)
		return helpers.HasFinalizer(res, AzureSQLServerFinalizerName)
	}, tc.timeoutFast, tc.retry, "wait for server to have finalizer")

	assert.Eventually(func() bool {
		var status v1alpha1.ASOStatus
		_ = tc.k8sClient.Get(ctx, names, instance)
		runtime.Field(reflect.ValueOf(instance), "Status", status)
		return strings.Contains(status.Message, successMsg)
	}, tc.timeout, tc.retry, "wait for server to provision")

}
