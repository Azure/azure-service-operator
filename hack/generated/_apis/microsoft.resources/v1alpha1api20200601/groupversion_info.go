/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Package v1alpha1api20200601 contains API Schema definitions for the microsoft.resources v20200601 API group

// +kubebuilder:object:generate=true
// All object properties are optional by default, this will be overridden when needed:
// +kubebuilder:validation:Optional
// +groupName=microsoft.resources.infra.azure.com
package v1alpha1api20200601

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: "microsoft.resources.infra.azure.com", Version: "v1alpha1api20200601"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)
