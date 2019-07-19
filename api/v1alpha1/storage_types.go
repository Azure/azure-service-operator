/*
MIT License

Copyright (c) Microsoft Corporation. All rights reserved.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	storagesdk "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// StorageSpec defines the desired state of Storage
type StorageSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Location               string                `json:"location,omitempty"`
	AccountName            string                `json:"accountName,omitempty"`
	Sku                    Sku                   `json:"sku,omitempty"`
	Kind                   storagesdk.Kind       `json:"kind,omitempty"`
	AccessTier             storagesdk.AccessTier `json:"accessTier,omitempty"`
	EnableHTTPSTrafficOnly *bool                 `json:"supportsHttpsTrafficOnly,omitempty"`
}

// Sku the SKU of the storage account.
type Sku struct {
	// Name - The SKU name. Required for account creation; optional for update. Possible values include: 'StandardLRS', 'StandardGRS', 'StandardRAGRS', 'StandardZRS', 'PremiumLRS', 'PremiumZRS', 'StandardGZRS', 'StandardRAGZRS'
	Name storagesdk.SkuName `json:"name,omitempty"`
}

// StorageStatus defines the observed state of Storage
type StorageStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ResourceGroupName string `json:"resourceGroupName,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Storage is the Schema for the storages API
type Storage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StorageSpec   `json:"spec,omitempty"`
	Status StorageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StorageList contains a list of Storage
type StorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Storage `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Storage{}, &StorageList{})
}
