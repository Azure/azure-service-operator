// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20210601 "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20210601"
	v20210601s "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20210601/storage"
	v20220120p "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20220120preview"
	v20220120ps "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20220120preview/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type FlexibleServersConfigurationExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *FlexibleServersConfigurationExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20210601.FlexibleServersConfiguration{},
		&v20210601s.FlexibleServersConfiguration{},
		&v20220120p.FlexibleServersConfiguration{},
		&v20220120ps.FlexibleServersConfiguration{}}
}
