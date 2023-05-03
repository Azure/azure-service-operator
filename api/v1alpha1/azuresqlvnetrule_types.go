// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
// AzureSQLVNetRuleSpec defines the desired state of AzureSQLVNetRule
type AzureSQLVNetRuleSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// +kubebuilder:validation:Pattern=^[-\w\._\(\)]+$
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	ResourceGroup string `json:"resourceGroup"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Server                       string `json:"server"`
	ServerSubscriptionID         string `json:"serverSubscriptionID,omitempty"`
	VNetResourceGroup            string `json:"vNetResourceGroup"`
	VNetName                     string `json:"vNetName"`
	SubnetName                   string `json:"subnetName"`
	VNetSubscriptionID           string `json:"vNetSubscriptionID,omitempty"`
	IgnoreMissingServiceEndpoint bool   `json:"ignoreMissingServiceEndpoint,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// AzureSQLVNetRule is the Schema for the azuresqlvnetrules API
// +kubebuilder:resource:shortName=asqlvnr
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type AzureSQLVNetRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureSQLVNetRuleSpec `json:"spec,omitempty"`
	Status ASOStatus            `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureSQLVNetRuleList contains a list of AzureSQLVNetRule
type AzureSQLVNetRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureSQLVNetRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzureSQLVNetRule{}, &AzureSQLVNetRuleList{})
}
