// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// StorageAccountSpec defines the desired state of Storage
type StorageAccountSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:MinLength=0

	Location      string `json:"location,omitempty"`
	ResourceGroup string `json:"resourceGroup"`

	Sku StorageAccountSku `json:"sku,omitempty"`

	Kind StorageAccountKind `json:"kind,omitempty"`

	AccessTier StorageAccountAccessTier `json:"accessTier,omitempty"`

	EnableHTTPSTrafficOnly *bool `json:"supportsHttpsTrafficOnly,omitempty"`

	DataLakeEnabled *bool `json:"dataLakeEnabled,omitempty"`

	NetworkRule *StorageNetworkRuleSet `json:"networkRule,omitempty"`
}

// StorageAccountSku the SKU of the storage account.
type StorageAccountSku struct {
	// Name - The SKU name. Required for account creation; optional for update.
	// Possible values include: 'StandardLRS', 'StandardGRS', 'StandardRAGRS', 'StandardZRS', 'PremiumLRS', 'PremiumZRS', 'StandardGZRS', 'StandardRAGZRS'
	Name StorageAccountSkuName `json:"name,omitempty"`
}

// StorageAccountSkuName enumerates the values for sku name.
// Only one of the following sku names may be specified.
// If none of the following sku names is specified, the default one
// is StorageV2.
// +kubebuilder:validation:Enum=Premium_LRS;Premium_ZRS;Standard_GRS;Standard_GZRS;Standard_LRS;Standard_RAGRS;Standard_RAGZRS;Standard_ZRS
type StorageAccountSkuName string

// StorageAccountKind enumerates the values for kind.
// Only one of the following kinds may be specified.
// If none of the following kinds is specified, the default one
// is StorageV2.
// +kubebuilder:validation:Enum=BlobStorage;BlockBlobStorage;FileStorage;Storage;StorageV2
type StorageAccountKind string

// StorageAccountAccessTier enumerates the values for access tier.
// Only one of the following access tiers may be specified.
// If none of the following access tiers is specified, the default one
// is Hot.
// +kubebuilder:validation:Enum=Cool;Hot
type StorageAccountAccessTier string

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// StorageAccount is the Schema for the storages API
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type StorageAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec                StorageAccountSpec                `json:"spec,omitempty"`
	Status              ASOStatus                         `json:"status,omitempty"`
	Output              StorageAccountOutput              `json:"output,omitempty"`
	AdditionalResources StorageAccountAdditionalResources `json:"additionalResources,omitempty"`
}

// StorageAccountOutput is the object that contains the output from creating a Storage Account object
type StorageAccountOutput struct {
	StorageAccountName string `json:"storageAccountName,omitempty"`
	Key1               string `json:"key1,omitempty"`
	Key2               string `json:"key2,omitempty"`
	ConnectionString1  string `json:"connectionString1,omitempty"`
	ConnectionString2  string `json:"connectionString2,omitempty"`
}

// StorageAccountAdditionalResources holds the additional resources
type StorageAccountAdditionalResources struct {
	Secrets []string `json:"secrets,omitempty"`
}

// +kubebuilder:object:root=true

// StorageAccountList contains a list of Storage
type StorageAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccount `json:"items"`
}

type Bypass string

type StorageNetworkRuleSet struct {
	// Bypass - Specifies whether traffic is bypassed for Logging/Metrics/AzureServices.
	//Possible values are any combination of Logging|Metrics|AzureServices (For example, "Logging, Metrics"), or None to bypass none of those traffics.
	//Possible values include: 'None', 'Logging', 'Metrics', 'AzureServices'
	Bypass Bypass `json:"bypass,omitempty"`
	// VirtualNetworkRules - Sets the virtual network rules
	VirtualNetworkRules *[]VirtualNetworkRule `json:"virtualNetworkRules,omitempty"`
	// IPRules - Sets the IP ACL rules
	IPRules *[]IPRule `json:"ipRules,omitempty"`
	// DefaultAction - Specifies the default action of allow or deny when no other rules match. Possible values include: 'DefaultActionAllow', 'DefaultActionDeny'
	DefaultAction string `json:"defaultAction,omitempty"`
}

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
	// SubnetId - Resource ID of a subnet, for example: /subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}.
	SubnetId *string `json:"subnetId,omitempty"`
}

type IPRule struct {
	// IPAddressOrRange - Specifies the IP or IP range in CIDR format. Only IPV4 address is allowed.
	IPAddressOrRange *string `json:"ipAddressOrRange,omitempty"`
}

func init() {
	SchemeBuilder.Register(&StorageAccount{}, &StorageAccountList{})
}

func (storageAccount *StorageAccount) IsSubmitted() bool {
	return storageAccount.Status.Provisioning || storageAccount.Status.Provisioned
}
