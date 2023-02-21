// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v1api20210515 "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20210515"
	v1api20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20210515storage"
	v20210515 "github.com/Azure/azure-service-operator/v2/api/documentdb/v1beta20210515"
	v20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1beta20210515storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type SqlDatabaseExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *SqlDatabaseExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v1api20210515.SqlDatabase{},
		&v1api20210515s.SqlDatabase{},
		&v20210515.SqlDatabase{},
		&v20210515s.SqlDatabase{}}
}
