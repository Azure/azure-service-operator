/*
Copyright 2017 The Kubernetes Authors.
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

// This package is taken from https://github.com/kubernetes-sigs/cluster-api/tree/master/util/patch with slight modifications
// Those modifications are:
//  - Removing stuff associated with conditions, which we're not currently using
//  - Making patch of status and spec happen sequentially and return error if one fails as opposed to returning an aggregate
//    error of both. This just makes it a bit easier to determine why there was a failure as you can use the standard
//    apierrors methods.
// TODO: Ensure we add attribution in License file once we have one

package patch

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Helper is a utility for ensuring the proper patching of objects.
type Helper struct {
	client       client.Client
	beforeObject runtime.Object
	before       *unstructured.Unstructured
	after        *unstructured.Unstructured
	changes      map[string]bool
}

// NewHelper returns an initialized Helper
func NewHelper(obj runtime.Object, crClient client.Client) (*Helper, error) {
	// Return early if the object is nil.
	// If you're wondering why we need reflection to do this check, see https://golang.org/doc/faq#nil_error.
	if obj == nil || (reflect.ValueOf(obj).IsValid() && reflect.ValueOf(obj).IsNil()) {
		return nil, errors.Errorf("expected non-nil object")
	}

	// Convert the object to unstructured to compare against our before copy.
	unstructuredObj, err := toUnstructured(obj)
	if err != nil {
		return nil, err
	}

	return &Helper{
		client:       crClient,
		before:       unstructuredObj,
		beforeObject: obj.DeepCopyObject(),
	}, nil
}

// Patch will attempt to patch the given object, including its status.
func (h *Helper) Patch(ctx context.Context, obj runtime.Object) error {
	if obj == nil {
		return errors.Errorf("expected non-nil object")
	}

	// Convert the object to unstructured to compare against our before copy.
	var err error
	h.after, err = toUnstructured(obj)
	if err != nil {
		return err
	}

	// Calculate and store the top-level field changes (e.g. "metadata", "spec", "status") we have before/after.
	h.changes, err = h.calculateChanges(obj)
	if err != nil {
		return err
	}

	if err := h.patch(ctx, obj); err != nil {
		return err
	}

	if err := h.patchStatus(ctx, obj); err != nil {
		return err
	}

	return nil
}

// patch issues a patch for metadata and spec.
func (h *Helper) patch(ctx context.Context, obj runtime.Object) error {
	if !h.shouldPatch("metadata") && !h.shouldPatch("spec") {
		return nil
	}
	beforeObject, afterObject, err := h.calculatePatch(obj, specPatch)
	if err != nil {
		return err
	}
	return h.client.Patch(ctx, afterObject, client.MergeFrom(beforeObject))
}

// patchStatus issues a patch if the status has changed.
func (h *Helper) patchStatus(ctx context.Context, obj runtime.Object) error {
	if !h.shouldPatch("status") {
		return nil
	}
	beforeObject, afterObject, err := h.calculatePatch(obj, statusPatch)
	if err != nil {
		return err
	}
	return h.client.Status().Patch(ctx, afterObject, client.MergeFrom(beforeObject))
}

// calculatePatch returns the before/after objects to be given in a controller-runtime patch, scoped down to the absolute necessary.
func (h *Helper) calculatePatch(afterObj runtime.Object, focus patchType) (runtime.Object, runtime.Object, error) {
	// Make a copy of the unstructured objects first.
	before := h.before.DeepCopy()
	after := h.after.DeepCopy()

	// Let's loop on the copies of our before/after and remove all the keys we don't need.
	for _, v := range []*unstructured.Unstructured{before, after} {
		// Ranges over the keys of the unstructured object, think of this as the very top level of an object
		// when submitting a yaml to kubectl or a client.
		//
		// These would be keys like `apiVersion`, `kind`, `metadata`, `spec`, `status`, etc.
		for key := range v.Object {
			// If the current key isn't something we absolutetly need (see the map for reference),
			// and it's not our current focus, then we should remove it.
			if key != focus.Key() && !preserveUnstructuredKeys[key] {
				unstructured.RemoveNestedField(v.Object, key)
				continue
			}
		}
	}

	// We've now applied all modifications to local unstructured objects,
	// make copies of the original objects and convert them back.
	beforeObj := h.beforeObject.DeepCopyObject()
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(before.Object, beforeObj); err != nil {
		return nil, nil, err
	}
	afterObj = afterObj.DeepCopyObject()
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(after.Object, afterObj); err != nil {
		return nil, nil, err
	}

	return beforeObj, afterObj, nil
}

func (h *Helper) shouldPatch(in string) bool {
	return h.changes[in]
}

// calculate changes tries to build a patch from the before/after objects we have
// and store in a map which top-level fields (e.g. `metadata`, `spec`, `status`, etc.) have changed.
func (h *Helper) calculateChanges(after runtime.Object) (map[string]bool, error) {
	// Calculate patch data.
	patch := client.MergeFrom(h.beforeObject)
	diff, err := patch.Data(after)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to calculate patch data")
	}

	// Unmarshal patch data into a local map.
	patchDiff := map[string]interface{}{}
	if err := json.Unmarshal(diff, &patchDiff); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal patch data into a map")
	}

	// Return the map.
	res := make(map[string]bool, len(patchDiff))
	for key := range patchDiff {
		res[key] = true
	}
	return res, nil
}
