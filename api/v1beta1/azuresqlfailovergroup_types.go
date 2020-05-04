// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1beta1

import (
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AzureSqlFailoverGroupSpec defines the desired state of AzureSqlFailoverGroup
type AzureSqlFailoverGroupSpec struct {
	// Important: Run "make" to regenerate code after modifying this file
	Location                     string                                   `json:"location"`
	ResourceGroup                string                                   `json:"resourcegroup,omitempty"`
	Server                       string                                   `json:"server"`
	FailoverPolicy               v1alpha1.ReadWriteEndpointFailoverPolicy `json:"failoverPolicy"`
	FailoverGracePeriod          int32                                    `json:"failoverGracePeriod"`
	SecondaryServer              string                                   `json:"secondaryserver"`
	SecondaryServerResourceGroup string                                   `json:"secondaryserverresourcegroup"`
	DatabaseList                 []string                                 `json:"databaselist"`
	KeyVaultToStoreSecrets       string                                   `json:"keyVaultToStoreSecrets,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AzureSqlFailoverGroup is the Schema for the azuresqlfailovergroups API
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type AzureSqlFailoverGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureSqlFailoverGroupSpec `json:"spec,omitempty"`
	Status v1alpha1.ASOStatus        `json:"status,omitempty"`
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
