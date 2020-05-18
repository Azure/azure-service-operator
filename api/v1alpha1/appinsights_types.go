// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AppInsightsSpec defines the desired state of AppInsights
type AppInsightsSpec struct {
	Kind            string `json:"kind"`
	Location        string `json:"location"`
	ApplicationType string `json:"applicationType"` // Possible values include 'web' or 'other'
	ResourceGroup   string `json:"resourceGroup,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AppInsights is the Schema for the appinsights API
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type AppInsights struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppInsightsSpec `json:"spec,omitempty"`
	Status ASOStatus       `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AppInsightsList contains a list of AppInsights
type AppInsightsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppInsights `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AppInsights{}, &AppInsightsList{})
}
