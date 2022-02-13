// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	authorization "github.com/Azure/azure-service-operator/v2/api/authorization/v1alpha1api20200801preview"
	"github.com/Azure/azure-service-operator/v2/api/authorization/v1alpha1api20200801previewstorage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type RoleAssignmentExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *RoleAssignmentExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&authorization.RoleAssignment{},
		&v1alpha1api20200801previewstorage.RoleAssignment{}}
}
