/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by azure-service-operator-codegen. DO NOT EDIT.

// Package v1alpha1api20211001 contains API Schema definitions for the microsoft.signalrservice v1alpha1api20211001 API group
// +kubebuilder:object:generate=true
// All object properties are optional by default, this will be overridden when needed:
// +kubebuilder:validation:Optional
// +groupName=microsoft.signalrservice.azure.com
package v1alpha1api20211001

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: "microsoft.signalrservice.azure.com", Version: "v1alpha1api20211001"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme

	localSchemeBuilder = SchemeBuilder.SchemeBuilder
)
