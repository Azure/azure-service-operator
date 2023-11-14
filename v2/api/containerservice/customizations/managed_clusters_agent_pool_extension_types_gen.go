// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20210501 "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20210501"
	v20210501s "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20210501storage"
	v20230201 "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230201"
	v20230201s "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230201storage"
	v20230202p "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230202preview"
	v20230202ps "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230202previewstorage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type ManagedClustersAgentPoolExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *ManagedClustersAgentPoolExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20210501.ManagedClustersAgentPool{},
		&v20210501s.ManagedClustersAgentPool{},
		&v20230201.ManagedClustersAgentPool{},
		&v20230201s.ManagedClustersAgentPool{},
		&v20230202p.ManagedClustersAgentPool{},
		&v20230202ps.ManagedClustersAgentPool{}}
}
