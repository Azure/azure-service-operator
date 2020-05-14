// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha2

import (
	s "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BlobContainerSpec defines the desired state of BlobContainer
type BlobContainerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Location      string         `json:"location"`
	ResourceGroup string         `json:"resourceGroup,omitempty"`
	AccountName   string         `json:"accountName,omitempty"`
	AccessLevel   s.PublicAccess `json:"accessLevel,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BlobContainer is the Schema for the blobcontainers API
// +kubebuilder:resource:shortName=bc,path=blobcontainer
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type BlobContainer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BlobContainerSpec `json:"spec,omitempty"`
	Status ASOStatus         `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BlobContainerList contains a list of BlobContainer
type BlobContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BlobContainer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BlobContainer{}, &BlobContainerList{})
}

func (bc *BlobContainer) IsSubmitted() bool {
	return bc.Status.Provisioned || bc.Status.Provisioning
}

func (bc *BlobContainer) IsProvisioned() bool {
	return bc.Status.Provisioned
}

func (bc *BlobContainer) HasFinalizer(finalizerName string) bool {
	return helpers.ContainsString(bc.ObjectMeta.Finalizers, finalizerName)
}
