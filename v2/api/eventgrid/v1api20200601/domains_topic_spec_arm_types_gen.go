// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20200601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DomainsTopic_Spec_ARM struct {
	Name string `json:"name,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DomainsTopic_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (topic DomainsTopic_Spec_ARM) GetAPIVersion() string {
	return "2020-06-01"
}

// GetName returns the Name of the resource
func (topic *DomainsTopic_Spec_ARM) GetName() string {
	return topic.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/domains/topics"
func (topic *DomainsTopic_Spec_ARM) GetType() string {
	return "Microsoft.EventGrid/domains/topics"
}
