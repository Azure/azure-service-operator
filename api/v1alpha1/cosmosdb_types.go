// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

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

	// +kubebuilder:validation:MinLength=0

	Location               string                        `json:"location,omitempty"`
	ResourceGroup          string                        `json:"resourceGroup"`
	Kind                   CosmosDBKind                  `json:"kind,omitempty"`
	Properties             CosmosDBProperties            `json:"properties,omitempty"`
	VirtualNetworkRules    *[]CosmosDBVirtualNetworkRule `json:"virtualNetworkRules,omitempty"`
	KeyVaultToStoreSecrets string                        `json:"keyVaultToStoreSecrets,omitempty"`
	IPRules                *[]string                     `json:"ipRules,omitempty"`
}

// CosmosDBKind enumerates the values for kind.
// Only one of the following kinds may be specified.
// If none of the following kinds is specified, the default one
// is GlobalDocumentDBKind.
// +kubebuilder:validation:Enum=GlobalDocumentDB;MongoDB
type CosmosDBKind string

const (
	// CosmosDBKindGlobalDocumentDB string constant describing global document database
	CosmosDBKindGlobalDocumentDB CosmosDBKind = "GlobalDocumentDB"
	// CosmosDBKindMongoDB string constant describing mongo database
	CosmosDBKindMongoDB CosmosDBKind = "MongoDB"
)

// CosmosDBProperties the CosmosDBProperties of CosmosDB.
type CosmosDBProperties struct {
	// DatabaseAccountOfferType - The offer type for the Cosmos DB database account.
	DatabaseAccountOfferType CosmosDBDatabaseAccountOfferType `json:"databaseAccountOfferType,omitempty"`
	// IsVirtualNetworkFilterEnabled - Flag to indicate whether to enable/disable Virtual Network ACL rules.
	IsVirtualNetworkFilterEnabled bool   `json:"isVirtualNetworkFilterEnabled,omitempty"`
	EnableMultipleWriteLocations  bool   `json:"enableMultipleWriteLocations,omitempty"`
	MongoDBVersion                string `json:"mongoDBVersion,omitempty"`
}

// +kubebuilder:validation:Enum=Standard
type CosmosDBDatabaseAccountOfferType string

const (
	// CosmosDBDatabaseAccountOfferTypeStandard string constant describing standard account offer type
	CosmosDBDatabaseAccountOfferTypeStandard CosmosDBDatabaseAccountOfferType = "Standard"
)

/*
type CosmosDBLocation struct {
	FailoverPriority int    `json:"failoverPriority,omitempty"`
	LocationName     string `json:"locationName,omitempty"`
	IsZoneRedundant  bool   `json:"isZoneRedundant,omitempty"`
}
*/

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// CosmosDB is the Schema for the cosmosdbs API
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type CosmosDB struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CosmosDBSpec `json:"spec,omitempty"`
	Status ASOStatus    `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// CosmosDBList contains a list of CosmosDB
type CosmosDBList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CosmosDB `json:"items"`
}

//CosmosDBVirtualNetworkRule virtual Network ACL Rule object
type CosmosDBVirtualNetworkRule struct {
	// ID - Resource ID of a subnet, for example: /subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}.
	SubnetID *string `json:"subnetID,omitempty"`
	// IgnoreMissingVNetServiceEndpoint - Create firewall rule before the virtual network has vnet service endpoint enabled.
	IgnoreMissingVNetServiceEndpoint *bool `json:"ignoreMissingVNetServiceEndpoint,omitempty"`
}

func init() {
	SchemeBuilder.Register(&CosmosDB{}, &CosmosDBList{})
}

// IsSubmitted function to determine if CosmosDB is provisioning or provisioned
func (cosmosDB *CosmosDB) IsSubmitted() bool {
	return cosmosDB.Status.Provisioning || cosmosDB.Status.Provisioned
}
