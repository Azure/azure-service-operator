/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by azure-service-operator-codegen. DO NOT EDIT.

// Package v1beta20211001storage contains API Schema definitions for the containerinstance v1beta20211001storage API group
// +kubebuilder:object:generate=true
// All object properties are optional by default, this will be overridden when needed:
// +kubebuilder:validation:Optional
// +groupName=containerinstance.azure.com
package v1beta20211001storage

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: "containerinstance.azure.com", Version: "v1beta20211001storage"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme

	localSchemeBuilder = SchemeBuilder.SchemeBuilder
)
