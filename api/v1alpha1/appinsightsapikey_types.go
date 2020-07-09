// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AppInsightsApiKeySpec defines the desired state of AppInsightsApiKey
type AppInsightsApiKeySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	ResourceGroup         string `json:"resourceGroup"`
	AppInsights           string `json:"appInsights"`
	ReadTelemetry         bool   `json:"readTelemetry,omitempty"`
	WriteAnnotations      bool   `json:"writeAnnotations,omitempty"`
	AuthSDKControlChannel bool   `json:"authSDKControlChannel,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
// AppInsightsApiKey is the Schema for the appinsightsapikeys API
type AppInsightsApiKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppInsightsApiKeySpec `json:"spec,omitempty"`
	Status ASOStatus             `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppInsightsApiKeyList contains a list of AppInsightsApiKey
type AppInsightsApiKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppInsightsApiKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AppInsightsApiKey{}, &AppInsightsApiKeyList{})
}
