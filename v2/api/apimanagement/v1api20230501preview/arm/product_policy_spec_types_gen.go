// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type ProductPolicy_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: Properties of the Policy.
	Properties *PolicyContractProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &ProductPolicy_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01-preview"
func (policy ProductPolicy_Spec) GetAPIVersion() string {
	return "2023-05-01-preview"
}

// GetName returns the Name of the resource
func (policy *ProductPolicy_Spec) GetName() string {
	return policy.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/products/policies"
func (policy *ProductPolicy_Spec) GetType() string {
	return "Microsoft.ApiManagement/service/products/policies"
}
