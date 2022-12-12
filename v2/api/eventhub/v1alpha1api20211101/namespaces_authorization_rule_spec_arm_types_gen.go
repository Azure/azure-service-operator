// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of Namespaces_AuthorizationRule_Spec. Use v1beta20211101.Namespaces_AuthorizationRule_Spec instead
type Namespaces_AuthorizationRule_Spec_ARM struct {
	Name       string                                            `json:"name,omitempty"`
	Properties *Namespaces_AuthorizationRule_Properties_Spec_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Namespaces_AuthorizationRule_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (rule Namespaces_AuthorizationRule_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (rule *Namespaces_AuthorizationRule_Spec_ARM) GetName() string {
	return rule.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventHub/namespaces/authorizationRules"
func (rule *Namespaces_AuthorizationRule_Spec_ARM) GetType() string {
	return "Microsoft.EventHub/namespaces/authorizationRules"
}

// Deprecated version of Namespaces_AuthorizationRule_Properties_Spec. Use v1beta20211101.Namespaces_AuthorizationRule_Properties_Spec instead
type Namespaces_AuthorizationRule_Properties_Spec_ARM struct {
	Rights []Namespaces_AuthorizationRule_Properties_Rights_Spec `json:"rights,omitempty"`
}
