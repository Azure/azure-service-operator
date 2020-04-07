// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

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

	Location      string `json:"location,omitempty"`
	ResourceGroup string `json:"resourceGroup"`

	Sku StorageSku `json:"sku,omitempty"`

	Kind StorageKind `json:"kind,omitempty"`

	AccessTier StorageAccessTier `json:"accessTier,omitempty"`

	EnableHTTPSTrafficOnly *bool `json:"supportsHttpsTrafficOnly,omitempty"`

	DataLakeEnabled *bool `json:"dataLakeEnabled,omitempty"`

	NetworkRule *StorageNetworkRuleSet `json:"networkRule,omitempty"`
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

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Storage is the Schema for the storages API
type Storage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec                StorageSpec                `json:"spec,omitempty"`
	Status              ASOStatus                  `json:"status,omitempty"`
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

type Bypass string
type Action string

type StorageNetworkRuleSet struct {
	// Bypass - Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Possible values are any combination of Logging|Metrics|AzureServices (For example, "Logging, Metrics"), or None to bypass none of those traffics. Possible values include: 'None', 'Logging', 'Metrics', 'AzureServices'
	Bypass Bypass `json:"bypass,omitempty"`
	// VirtualNetworkRules - Sets the virtual network rules
	VirtualNetworkRules *[]VirtualNetworkRule `json:"virtualNetworkRules,omitempty"`
	// IPRules - Sets the IP ACL rules
	IPRule *[]IPRule `json:"ipRules,omitempty"`
	// DefaultAction - Specifies the default action of allow or deny when no other rules match. Possible values include: 'DefaultActionAllow', 'DefaultActionDeny'
	DefaultAction DefaultAction `json:"defaultAction,omitempty"`
}

// DefaultAction defined as a string

const (

	// AzureServices ...
	AzureServices Bypass = "AzureServices"
	// Logging ...
	Logging Bypass = "Logging"
	// Metrics ...
	Metrics Bypass = "Metrics"
	// None ...
	None Bypass = "None"
)

type VirtualNetworkRule struct {
	// VirtualNetworkResourceID - Resource ID of a subnet, for example: /subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}.
	VirtualNetworkResourceID *string `json:"id,omitempty"`
	Action                   Action  `json:"action,omitempty"`
}

type IPRule struct {
	// IPAddressOrRange - Specifies the IP or IP range in CIDR format. Only IPV4 address is allowed.
	IPAddressOrRange *string `json:"value,omitempty"`
	Action           Action  `json:"action,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Storage{}, &StorageList{})
}

func (storage *Storage) IsSubmitted() bool {
	return storage.Status.Provisioning || storage.Status.Provisioned
}
