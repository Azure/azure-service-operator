// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ReadWriteEndpointFailoverPolicy - wraps https://godoc.org/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql#ReadWriteEndpointFailoverPolicy
// +kubebuilder:validation:Enum=Automatic;Manual
type ReadWriteEndpointFailoverPolicy string

const (
	// Automatic ...
	FailoverPolicyAutomatic ReadWriteEndpointFailoverPolicy = "Automatic"
	// Manual ...
	FailoverPolicyManual ReadWriteEndpointFailoverPolicy = "Manual"
)

// AzureSqlFailoverGroupSpec defines the desired state of AzureSqlFailoverGroup
type AzureSqlFailoverGroupSpec struct {
	// Important: Run "make" to regenerate code after modifying this file
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Location string `json:"location"`
	// +kubebuilder:validation:Pattern=^[-\w\._\(\)]+$
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	ResourceGroup string `json:"resourceGroup"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Server string `json:"server"`

	SubscriptionID string `json:"subscriptionId,omitempty"`

	FailoverPolicy ReadWriteEndpointFailoverPolicy `json:"failoverPolicy"`
	// TODO: This field should be a ptr as it must not be specified if the failover policy is Manual,
	// TODO: but is required when the policy is Automatic
	FailoverGracePeriod           int32    `json:"failoverGracePeriod"`
	SecondaryServer               string   `json:"secondaryServer"`
	SecondaryServerResourceGroup  string   `json:"secondaryServerResourceGroup"`
	SecondaryServerSubscriptionID string   `json:"SecondaryServerSubscriptionId,omitempty"`
	DatabaseList                  []string `json:"databaseList"`
	KeyVaultToStoreSecrets        string   `json:"keyVaultToStoreSecrets,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AzureSqlFailoverGroup is the Schema for the azuresqlfailovergroups API
// +kubebuilder:resource:shortName=asqlfog
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type AzureSqlFailoverGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureSqlFailoverGroupSpec `json:"spec,omitempty"`
	Status ASOStatus                 `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureSqlFailoverGroupList contains a list of AzureSqlFailoverGroup
type AzureSqlFailoverGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureSqlFailoverGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzureSqlFailoverGroup{}, &AzureSqlFailoverGroupList{})
}

func (s *AzureSqlFailoverGroup) IsSubmitted() bool {
	return s.Status.Provisioned || s.Status.Provisioning
}
