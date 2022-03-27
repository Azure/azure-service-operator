/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genruntime

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// GroupVersionKindAware is implemented by resources that are aware of which version of the resource was originally
// specified. This allows us to interace with ARM using an API version specified by an end user.
type GroupVersionKindAware interface {
	// OriginalGVK returns the GroupVersionKind originally used to create the resource (regardless of any conversions)
	OriginalGVK() *schema.GroupVersionKind
}

// GetOriginalGVK gets the GVK the original GVK the object was created with.
func GetOriginalGVK(obj MetaObject) schema.GroupVersionKind {
	// If our current resource is aware of its original GVK, use that for our result
	aware, ok := obj.(GroupVersionKindAware)
	if ok {
		return *aware.OriginalGVK()
	}

	// The GVK of our current object
	return obj.GetObjectKind().GroupVersionKind()
}
