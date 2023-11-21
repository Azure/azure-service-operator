// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20210501 "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20210501"
	v20210501s "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20210501/storage"
	v20230201 "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230201"
	v20230201s "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230201/storage"
	v20230202p "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230202preview"
	v20230202ps "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230202preview/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type ManagedClusterExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *ManagedClusterExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20210501.ManagedCluster{},
		&v20210501s.ManagedCluster{},
		&v20230201.ManagedCluster{},
		&v20230201s.ManagedCluster{},
		&v20230202p.ManagedCluster{},
		&v20230202ps.ManagedCluster{}}
}
