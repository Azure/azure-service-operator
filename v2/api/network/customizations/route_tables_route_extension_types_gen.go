// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v1api20201101 "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	v1api20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101storage"
	v20201101 "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101"
	v20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type RouteTablesRouteExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *RouteTablesRouteExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v1api20201101.RouteTablesRoute{},
		&v1api20201101s.RouteTablesRoute{},
		&v20201101.RouteTablesRoute{},
		&v20201101s.RouteTablesRoute{}}
}
