// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1beta1

import (
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AzureSqlServerSpec defines the desired state of AzureSqlServer
type AzureSqlServerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Location               string `json:"location"`
	ResourceGroup          string `json:"resourceGroup,omitempty"`
	KeyVaultToStoreSecrets string `json:"keyVaultToStoreSecrets,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AzureSqlServer is the Schema for the azuresqlservers API
type AzureSqlServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureSqlServerSpec `json:"spec,omitempty"`
	Status v1alpha1.ASOStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureSqlServerList contains a list of AzureSqlServer
type AzureSqlServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureSqlServer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzureSqlServer{}, &AzureSqlServerList{})
}

func (s *AzureSqlServer) IsSubmitted() bool {
	return s.Status.Provisioned || s.Status.Provisioning
}

// NewAzureSQLServer returns a simple server struct filled with passed in values
func NewAzureSQLServer(names types.NamespacedName, resourceGroup, region string) *AzureSqlServer {
	return &AzureSqlServer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      names.Name,
			Namespace: names.Namespace,
		},
		Spec: AzureSqlServerSpec{
			Location:      region,
			ResourceGroup: resourceGroup,
		},
	}
}
