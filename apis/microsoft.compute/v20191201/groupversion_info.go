/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Package v20191201 contains API Schema definitions for the microsoft.compute v20191201 API group
// +kubebuilder:object:generate=true
// +groupName=microsoft.compute.infra.azure.com
package v20191201

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: "microsoft.compute.infra.azure.com", Version: "v20191201"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme

	localSchemeBuilder = SchemeBuilder.SchemeBuilder
)
