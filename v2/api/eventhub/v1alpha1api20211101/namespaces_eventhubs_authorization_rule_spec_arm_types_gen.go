// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of Namespaces_Eventhubs_AuthorizationRule_Spec. Use v1beta20211101.Namespaces_Eventhubs_AuthorizationRule_Spec instead
type Namespaces_Eventhubs_AuthorizationRule_Spec_ARM struct {
	Location   *string                          `json:"location,omitempty"`
	Name       string                           `json:"name,omitempty"`
	Properties *AuthorizationRuleProperties_ARM `json:"properties,omitempty"`
	Tags       map[string]string                `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Namespaces_Eventhubs_AuthorizationRule_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (rule Namespaces_Eventhubs_AuthorizationRule_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (rule *Namespaces_Eventhubs_AuthorizationRule_Spec_ARM) GetName() string {
	return rule.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventHub/namespaces/eventhubs/authorizationRules"
func (rule *Namespaces_Eventhubs_AuthorizationRule_Spec_ARM) GetType() string {
	return "Microsoft.EventHub/namespaces/eventhubs/authorizationRules"
}
