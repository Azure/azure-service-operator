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
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// StorageSpec defines the desired state of Storage
type StorageSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:MinLength=0

	Location          string `json:"location,omitempty"`
	ResourceGroupName string `json:"resourceGroup"`

	Sku StorageSku `json:"sku,omitempty"`

	Kind StorageKind `json:"kind,omitempty"`

	AccessTier StorageAccessTier `json:"accessTier,omitempty"`

	EnableHTTPSTrafficOnly *bool `json:"supportsHttpsTrafficOnly,omitempty"`
}

// Sku the SKU of the storage account.
type StorageSku struct {
	// Name - The SKU name. Required for account creation; optional for update.
	// Possible values include: 'StandardLRS', 'StandardGRS', 'StandardRAGRS', 'StandardZRS', 'PremiumLRS', 'PremiumZRS', 'StandardGZRS', 'StandardRAGZRS'
	Name StorageSkuName `json:"name,omitempty"`
}

// StorageSkuName enumerates the values for sku name.
// Only one of the following sku names may be specified.
// If none of the following sku names is specified, the default one
// is StorageV2.
// +kubebuilder:validation:Enum=Premium_LRS;Premium_ZRS;Standard_GRS;Standard_GZRS;Standard_LRS;Standard_RAGRS;Standard_RAGZRS;Standard_ZRS
type StorageSkuName string

// StorageKind enumerates the values for kind.
// Only one of the following kinds may be specified.
// If none of the following kinds is specified, the default one
// is StorageV2.
// +kubebuilder:validation:Enum=BlobStorage;BlockBlobStorage;FileStorage;Storage;StorageV2
type StorageKind string

// AccessTier enumerates the values for access tier.
// Only one of the following access tiers may be specified.
// If none of the following access tiers is specified, the default one
// is Hot.
// +kubebuilder:validation:Enum=Cool;Hot
type StorageAccessTier string

// StorageStatus defines the observed state of Storage
type StorageStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// DeploymentName    string `json:"deploymentName,omitempty"`
	// ProvisioningState string `json:"provisioningState,omitempty"`
	// Generation        int64  `json:"generation,omitempty"`
	Provisioning bool `json:"provisioning,omitempty"`
	Provisioned  bool `json:"provisioned,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Storage is the Schema for the storages API
type Storage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec                StorageSpec                `json:"spec,omitempty"`
	Status              StorageStatus              `json:"status,omitempty"`
	Output              StorageOutput              `json:"output,omitempty"`
	AdditionalResources StorageAdditionalResources `json:"additionalResources,omitempty"`
}

type StorageOutput struct {
	StorageAccountName string `json:"storageAccountName,omitempty"`
	Key1               string `json:"key1,omitempty"`
	Key2               string `json:"key2,omitempty"`
	ConnectionString1  string `json:"connectionString1,omitempty"`
	ConnectionString2  string `json:"connectionString2,omitempty"`
}

// StorageAdditionalResources holds the additional resources
type StorageAdditionalResources struct {
	Secrets []string `json:"secrets,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// StorageList contains a list of Storage
type StorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Storage `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Storage{}, &StorageList{})
}

func (storage *Storage) IsSubmitted() bool {
	return storage.Status.Provisioning || storage.Status.Provisioned
}
