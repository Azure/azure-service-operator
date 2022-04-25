// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	alpha20211101 "github.com/Azure/azure-service-operator/v2/api/eventhub/v1alpha1api20211101"
	alpha20211101s "github.com/Azure/azure-service-operator/v2/api/eventhub/v1alpha1api20211101storage"
	v20211101 "github.com/Azure/azure-service-operator/v2/api/eventhub/v1beta20211101"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/eventhub/v1beta20211101storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type NamespaceExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *NamespaceExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&alpha20211101.Namespace{},
		&alpha20211101s.Namespace{},
		&v20211101.Namespace{},
		&v20211101s.Namespace{}}
}
