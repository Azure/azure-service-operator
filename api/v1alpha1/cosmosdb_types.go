/*
.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CosmosDBSpec defines the desired state of CosmosDB
type CosmosDBSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// CosmosDBStatus defines the observed state of CosmosDB
type CosmosDBStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// CosmosDB is the Schema for the cosmosdbs API
type CosmosDB struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CosmosDBSpec   `json:"spec,omitempty"`
	Status CosmosDBStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CosmosDBList contains a list of CosmosDB
type CosmosDBList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CosmosDB `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CosmosDB{}, &CosmosDBList{})
}
