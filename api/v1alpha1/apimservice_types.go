// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ApimServiceSpec defines the desired state of ApimService
type ApimServiceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Tier                     string `json:"tier,omitempty"`
	Location                 string `json:"location"`
	ResourceGroup            string `json:"resourceGroup"`
	PublisherName            string `json:"publisherName"`
	PublisherEmail           string `json:"publisherEmail"`
	VnetType                 string `json:"vnetType,omitempty"`
	VnetResourceGroup        string `json:"vnetResourceGroup,omitempty"`
	VnetName                 string `json:"vnetName,omitempty"`
	VnetSubnetName           string `json:"vnetSubnetName,omitempty"`
	AppInsightsResourceGroup string `json:"appInsightsResourceGroup,omitempty"`
	AppInsightsName          string `json:"appInsightsName,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ApimService is the Schema for the apimservices API
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type ApimService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ApimServiceSpec `json:"spec,omitempty"`
	Status ASOStatus       `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApimServiceList contains a list of ApimService
type ApimServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApimService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApimService{}, &ApimServiceList{})
}
