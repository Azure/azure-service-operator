// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v1api20201201 "github.com/Azure/azure-service-operator/v2/api/compute/v1api20201201"
	v1api20201201s "github.com/Azure/azure-service-operator/v2/api/compute/v1api20201201storage"
	v1api20220301 "github.com/Azure/azure-service-operator/v2/api/compute/v1api20220301"
	v1api20220301s "github.com/Azure/azure-service-operator/v2/api/compute/v1api20220301storage"
	v20201201 "github.com/Azure/azure-service-operator/v2/api/compute/v1beta20201201"
	v20201201s "github.com/Azure/azure-service-operator/v2/api/compute/v1beta20201201storage"
	v20220301 "github.com/Azure/azure-service-operator/v2/api/compute/v1beta20220301"
	v20220301s "github.com/Azure/azure-service-operator/v2/api/compute/v1beta20220301storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type VirtualMachineScaleSetExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *VirtualMachineScaleSetExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v1api20201201.VirtualMachineScaleSet{},
		&v1api20201201s.VirtualMachineScaleSet{},
		&v1api20220301.VirtualMachineScaleSet{},
		&v1api20220301s.VirtualMachineScaleSet{},
		&v20201201.VirtualMachineScaleSet{},
		&v20201201s.VirtualMachineScaleSet{},
		&v20220301.VirtualMachineScaleSet{},
		&v20220301s.VirtualMachineScaleSet{}}
}