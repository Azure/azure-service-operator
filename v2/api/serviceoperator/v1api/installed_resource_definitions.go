// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// +kubebuilder:rbac:groups=serviceoperator.azure.com,resources=installedresourcedefinitions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=serviceoperator.azure.com,resources={installedresourcedefinitions/status,installedresourcedefinitions/finalizers},verbs=get;update;patch

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

// TODO: I think we just watch a magically named one of these? Similar to "aso-controller-settings"?
// TODO: How does the Helm update experience work with this resource? If the user has modified it away from its default settings, won't Helm
// TODO: overwrite it and cause problems? We may need a Helm flag that indicates what to do here
type InstalledResourceDefinitionsSpec struct {
	// +kubebuilder:validation:Required
	// Pattern defines the set of CRDs which should be installed.
	// Currently only "*" (all resources) is supported.
	// Future expansion will allow "<group>.*" to install all resources from a group (e.g. "network.*")
	// as well as "<group>.<kind> to install a specific resource only (e.g. "network.virtualnetworks")
	Pattern []string `json:"pattern,omitempty"`
}

type InstalledResourceDefinitionsStatus struct {
	//Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`
}

func init() {
	SchemeBuilder.Register(&InstalledResourceDefinitions{}, &InstalledResourceDefinitionsList{})
}
