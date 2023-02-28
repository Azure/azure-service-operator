/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Package v1api contains API Schema definitions for serviceoperator APIs
// +kubebuilder:object:generate=true
// All object properties are optional by default, this will be overridden when needed:
// +kubebuilder:validation:Optional
// +groupName=serviceoperator.azure.com
package v1api

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: "serviceoperator.azure.com", Version: "v1api"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)
