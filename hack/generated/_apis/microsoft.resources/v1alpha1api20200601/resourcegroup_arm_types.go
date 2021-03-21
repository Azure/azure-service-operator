// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

import "github.com/Azure/k8s-infra/hack/generated/pkg/genruntime"

type ResourceGroupStatusArm struct {
	ID string `json:"id,omitempty"`

	Name     string `json:"name,omitempty"`
	Location string `json:"location,omitempty"`

	// ManagedBy is the management group responsible for managing this group
	ManagedBy string `json:"managedBy,omitempty"`

	// Tags are user defined key value pairs
	Tags map[string]string `json:"tags,omitempty"`

	Properties *ResourceGroupStatusPropertiesArm `json:"properties,omitempty"` // TODO: Is this required or optional?
}

type ResourceGroupStatusPropertiesArm struct {
	ProvisioningState string `json:"provisioningState,omitempty"` // TODO: Wrong, needs to be in properties
}

type ResourceGroupSpecArm struct {

	//ApiVersion: API Version of the resource type, optional when apiProfile is used
	//on the template
	ApiVersion string `json:"apiVersion"`

	//Name: Name of the resource
	Name string `json:"name"`

	// Location is the Azure location for the group (eg westus2, southcentralus, etc...)
	Location string `json:"location"`

	// ManagedBy is the management group responsible for managing this group
	ManagedBy string `json:"managedBy,omitempty"` // TODO: ??

	// Tags are user defined key value pairs
	Tags map[string]string `json:"tags,omitempty"`

	//Type: Resource type
	Type ResourceGroupType `json:"type"`
}

var _ genruntime.ArmResourceSpec = &ResourceGroupSpecArm{}

// GetApiVersion returns the ApiVersion of the resource
func (spec ResourceGroupSpecArm) GetApiVersion() string {
	return string(spec.ApiVersion)
}

// GetName returns the Name of the resource
func (spec ResourceGroupSpecArm) GetName() string {
	return spec.Name
}

// GetType returns the Type of the resource
func (spec ResourceGroupSpecArm) GetType() string {
	return string(spec.Type)
}

type ResourceGroupType string

const ResourceGroupTypeResourceGroup = ResourceGroupType("Microsoft.Resources/resourceGroups")
