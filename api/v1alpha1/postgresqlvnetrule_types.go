// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PotgreSQLVNetRuleSpec defines the desired state of PostgreSQLVNetRule
type PostgreSQLVNetRuleSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ResourceGroup                string `json:"resourceGroup"`
	Server                       string `json:"server"`
	VNetResourceGroup            string `json:"vNetResourceGroup"`
	VNetName                     string `json:"vNetName"`
	SubnetName                   string `json:"subnetName"`
	IgnoreMissingServiceEndpoint bool   `json:"ignoreMissingServiceEndpoint,omitempty"`
}

// +kubebuilder:object:root=true

// PostgreSQLVNetRule is the Schema for the PostgreSQLVNetRules API
// +kubebuilder:resource:shortName=psqlvn,path=postgresqlvnetrule
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type PostgreSQLVNetRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PostgreSQLVNetRuleSpec `json:"spec,omitempty"`
	Status ASOStatus              `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PostgreSQLVNetRuleList contains a list of PostgreSQLVNetRule
type PostgreSQLVNetRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PostgreSQLVNetRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PostgreSQLVNetRule{}, &PostgreSQLVNetRuleList{})
}
