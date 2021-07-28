/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genruntime

import (
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
)

func NewObjectFromObject(obj client.Object, scheme *runtime.Scheme) (client.Object, error) {
	gvk, err := apiutil.GVKForObject(obj, scheme)
	if err != nil {
		return nil, err
	}

	// Create a fresh destination to deserialize to
	newObj, err := scheme.New(gvk)
	if err != nil {
		return nil, err
	}
	return newObj.(client.Object), nil
}
