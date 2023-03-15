// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// Note: We do NOT include the kubebuilder rbac annotations for installedresourcedefinitions here because
// we want to define a special role for this resource.
// We need a special role because this resource is different than the other CRDs. See the config/rbac/crd_manager_role.yaml
// and config/rbac/crd_manager_role_binding.yaml

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// +kubebuilder:storageversion
// +kubebuilder:resource:singular=installedresourcedefinitions
// InstalledResourceDefinitions is used to configure what resources ASO supports
type InstalledResourceDefinitions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstalledResourceDefinitionsSpec   `json:"spec,omitempty"`
	Status            InstalledResourceDefinitionsStatus `json:"status,omitempty"`
}

var _ conditions.Conditioner = &InstalledResourceDefinitions{}

// GetConditions returns the conditions of the resource
func (c *InstalledResourceDefinitions) GetConditions() conditions.Conditions {
	return c.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (c *InstalledResourceDefinitions) SetConditions(conditions conditions.Conditions) {
	c.Status.Conditions = conditions
}

// +kubebuilder:object:root=true
type InstalledResourceDefinitionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstalledResourceDefinitions `json:"items"`
}

type InstalledResourceDefinitionsSpec struct {
	// +kubebuilder:validation:Required
	// Patterns defines the set of CRDs which should be installed.
	// Currently only "*" (all resources) is supported.
	// Future expansion will allow "<group>.*" to install all resources from a group (e.g. "network.*")
	// as well as "<group>.<kind> to install a specific resource only (e.g. "network.virtualnetworks")
	Patterns []string `json:"patterns,omitempty"`
}

type InstalledResourceDefinitionsStatus struct {
	//Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`
}

func init() {
	SchemeBuilder.Register(&InstalledResourceDefinitions{}, &InstalledResourceDefinitionsList{})
}
