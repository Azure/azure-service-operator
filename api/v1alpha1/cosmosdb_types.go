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

// CosmosDBSpec defines the desired state of CosmosDB
type CosmosDBSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:MinLength=0

	Location   string             `json:"location,omitempty"`
	Kind       CosmosDBKind       `json:"kind,omitempty"`
	Properties CosmosDBProperties `json:"properties,omitempty"`
}

// CosmosDBKind enumerates the values for kind.
// Only one of the following kinds may be specified.
// If none of the following kinds is specified, the default one
// is GlobalDocumentDBKind.
// +kubebuilder:validation:Enum=GlobalDocumentDB;MongoDB
type CosmosDBKind string

const (
	CosmosDBKindGlobalDocumentDB CosmosDBKind = "GlobalDocumentDB"
	CosmosDBKindMongoDB          CosmosDBKind = "MongoDB"
)

// CosmosDBProperties the CosmosDBProperties of CosmosDB.
type CosmosDBProperties struct {
	// CosmosDBDatabaseAccountOfferType - The offer type for the Cosmos DB database account.
	DatabaseAccountOfferType CosmosDBDatabaseAccountOfferType `json:"databaseAccountOfferType,omitempty"`
	//Locations                []CosmosDBLocation               `json:"locations,omitempty"`
}

// +kubebuilder:validation:Enum=Standard
type CosmosDBDatabaseAccountOfferType string

const (
	CosmosDBDatabaseAccountOfferTypeStandard CosmosDBDatabaseAccountOfferType = "Standard"
)

/*
type CosmosDBLocation struct {
	FailoverPriority int    `json:"failoverPriority,omitempty"`
	LocationName     string `json:"locationName,omitempty"`
	IsZoneRedundant  bool   `json:"isZoneRedundant,omitempty"`
}
*/

// CosmosDBStatus defines the observed state of CosmosDB
type CosmosDBStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	DeploymentName    string `json:"deploymentName,omitempty"`
	ProvisioningState string `json:"provisioningState,omitempty"`
	Generation        int64  `json:"generation,omitempty"`
}

type CosmosDBOutput struct {
	CosmosDBName     string `json:"cosmosDBName,omitempty"`
	PrimaryMasterKey string `json:"primaryMasterKey,omitempty"`
	//SecondaryMasterKey         string `json:"secondaryMasterKey,omitempty"`
	//PrimaryReadonlyMasterKey   string `json:"primaryReadonlyMasterKey,omitempty"`
	//SecondaryReadonlyMasterKey string `json:"secondaryReadonlyMasterKey,omitempty"`
}

// CosmosDBAdditionalResources holds the additional resources
type CosmosDBAdditionalResources struct {
	Secrets []string `json:"secrets,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// CosmosDB is the Schema for the cosmosdbs API
type CosmosDB struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec                CosmosDBSpec                `json:"spec,omitempty"`
	Status              CosmosDBStatus              `json:"status,omitempty"`
	Output              CosmosDBOutput              `json:"output,omitempty"`
	AdditionalResources CosmosDBAdditionalResources `json:"additionalResources,omitempty"`
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
