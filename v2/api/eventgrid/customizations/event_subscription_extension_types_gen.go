// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v1api20200601 "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1api20200601"
	v1api20200601s "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1api20200601storage"
	v20200601 "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1beta20200601"
	v20200601s "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1beta20200601storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type EventSubscriptionExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *EventSubscriptionExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v1api20200601.EventSubscription{},
		&v1api20200601s.EventSubscription{},
		&v20200601.EventSubscription{},
		&v20200601s.EventSubscription{}}
}
