// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type NamespacesAuthorizationRule_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: AuthorizationRule properties.
	Properties *Namespaces_AuthorizationRule_Properties_Spec_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NamespacesAuthorizationRule_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (rule NamespacesAuthorizationRule_Spec_ARM) GetAPIVersion() string {
	return "2021-11-01"
}

// GetName returns the Name of the resource
func (rule *NamespacesAuthorizationRule_Spec_ARM) GetName() string {
	return rule.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ServiceBus/namespaces/AuthorizationRules"
func (rule *NamespacesAuthorizationRule_Spec_ARM) GetType() string {
	return "Microsoft.ServiceBus/namespaces/AuthorizationRules"
}

type Namespaces_AuthorizationRule_Properties_Spec_ARM struct {
	// Rights: The rights associated with the rule.
	Rights []Namespaces_AuthorizationRule_Properties_Rights_Spec_ARM `json:"rights,omitempty"`
}

// +kubebuilder:validation:Enum={"Listen","Manage","Send"}
type Namespaces_AuthorizationRule_Properties_Rights_Spec_ARM string

const (
	Namespaces_AuthorizationRule_Properties_Rights_Spec_ARM_Listen = Namespaces_AuthorizationRule_Properties_Rights_Spec_ARM("Listen")
	Namespaces_AuthorizationRule_Properties_Rights_Spec_ARM_Manage = Namespaces_AuthorizationRule_Properties_Rights_Spec_ARM("Manage")
	Namespaces_AuthorizationRule_Properties_Rights_Spec_ARM_Send   = Namespaces_AuthorizationRule_Properties_Rights_Spec_ARM("Send")
)

// Mapping from string to Namespaces_AuthorizationRule_Properties_Rights_Spec_ARM
var namespaces_AuthorizationRule_Properties_Rights_Spec_ARM_Values = map[string]Namespaces_AuthorizationRule_Properties_Rights_Spec_ARM{
	"listen": Namespaces_AuthorizationRule_Properties_Rights_Spec_ARM_Listen,
	"manage": Namespaces_AuthorizationRule_Properties_Rights_Spec_ARM_Manage,
	"send":   Namespaces_AuthorizationRule_Properties_Rights_Spec_ARM_Send,
}
