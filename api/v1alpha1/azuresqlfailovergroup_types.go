// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AzureSqlFailoverGroupSpec defines the desired state of AzureSqlFailoverGroup
type AzureSqlFailoverGroupSpec struct {
	// Important: Run "make" to regenerate code after modifying this file
	Location                     string                          `json:"location"`
	ResourceGroup                string                          `json:"resourcegroup,omitempty"`
	Server                       string                          `json:"server"`
	FailoverPolicy               ReadWriteEndpointFailoverPolicy `json:"failoverpolicy"`
	FailoverGracePeriod          int32                           `json:"failovergraceperiod"`
	SecondaryServerName          string                          `json:"secondaryserver"`
	SecondaryServerResourceGroup string                          `json:"secondaryserverresourcegroup"`
	DatabaseList                 []string                        `json:"databaselist"`
}

// AzureSqlFailoverGroupStatus defines the observed state of AzureSqlFailoverGroup
type AzureSqlFailoverGroupStatus struct {
	// Important: Run "make" to regenerate code after modifying this file
	Provisioning bool   `json:"provisioning,omitempty"`
	Provisioned  bool   `json:"provisioned,omitempty"`
	State        string `json:"state,omitempty"`
	Message      string `json:"message,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// AzureSqlFailoverGroup is the Schema for the azuresqlfailovergroups API
type AzureSqlFailoverGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureSqlFailoverGroupSpec   `json:"spec,omitempty"`
	Status AzureSqlFailoverGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureSqlFailoverGroupList contains a list of AzureSqlFailoverGroup
type AzureSqlFailoverGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureSqlFailoverGroup `json:"items"`
}

// ReadWriteEndpointFailoverPolicy - wraps https://godoc.org/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql#ReadWriteEndpointFailoverPolicy
type ReadWriteEndpointFailoverPolicy string

const (
	// Automatic ...
	FailoverPolicyAutomatic ReadWriteEndpointFailoverPolicy = "Automatic"
	// Manual ...
	FailoverPolicyManual ReadWriteEndpointFailoverPolicy = "Manual"
)

func init() {
	SchemeBuilder.Register(&AzureSqlFailoverGroup{}, &AzureSqlFailoverGroupList{})
}

func (s *AzureSqlFailoverGroup) IsSubmitted() bool {
	return s.Status.Provisioned || s.Status.Provisioning
}
