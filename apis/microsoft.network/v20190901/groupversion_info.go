/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Package v20190901 contains API Schema definitions for the microsoftnetwork v20190901 API group
// +kubebuilder:object:generate=true
// +groupName=microsoft.network.infra.azure.com
package v20190901

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: "microsoft.network.infra.azure.com", Version: "v20190901"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)
