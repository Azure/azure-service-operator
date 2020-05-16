// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MySQLVNetRuleSpec defines the desired state of MySQLVNetRule
type MySQLVNetRuleSpec struct {
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

// MySQLVNetRule is the Schema for the mysqlvnetrules API
// +kubebuilder:resource:shortName=mysqlvn,path=mysqlvnetrule
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type MySQLVNetRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MySQLVNetRuleSpec `json:"spec,omitempty"`
	Status ASOStatus         `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MySQLVNetRuleList contains a list of MySQLVNetRule
type MySQLVNetRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MySQLVNetRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MySQLVNetRule{}, &MySQLVNetRuleList{})
}
