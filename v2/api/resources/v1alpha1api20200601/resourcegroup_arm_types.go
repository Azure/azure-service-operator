// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type ResourceGroupStatusARM struct {
	ID *string `json:"id,omitempty"`

	Name     *string `json:"name,omitempty"`
	Location *string `json:"location,omitempty"`

	// ManagedBy is the management group responsible for managing this group
	ManagedBy *string `json:"managedBy,omitempty"`

	// Tags are user defined key value pairs
	Tags map[string]string `json:"tags,omitempty"`

	Properties *ResourceGroupStatusPropertiesARM `json:"properties,omitempty"`
}

type ResourceGroupStatusPropertiesARM struct {
	ProvisioningState *string `json:"provisioningState,omitempty"`
}

type ResourceGroupSpecARM struct {

	//Name: Name of the resource
	Name string `json:"name,omitempty"`

	// Location is the Azure location for the group (eg westus2, southcentralus, etc...)
	Location *string `json:"location,omitempty"`

	// ManagedBy is the management group responsible for managing this group
	ManagedBy *string `json:"managedBy,omitempty"` // TODO: ??

	// Tags are user defined key value pairs
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &ResourceGroupSpecARM{}

// GetAPIVersion returns the APIVersion of the resource
func (spec ResourceGroupSpecARM) GetAPIVersion() string {
	return "2020-06-01"
}

// GetName returns the Name of the resource
func (spec ResourceGroupSpecARM) GetName() string {
	return spec.Name
}

// GetType returns the Type of the resource
func (spec ResourceGroupSpecARM) GetType() string {
	return string(ResourceGroupTypeResourceGroup)
}

type ResourceGroupType string

const ResourceGroupTypeResourceGroup = ResourceGroupType("Microsoft.Resources/resourceGroups")
