// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CosmosDBSQLDatabaseSpec defines the desired state of the CosmosDBSQLDatabase
type CosmosDBSQLDatabaseSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Pattern=^[-\w\._\(\)]+$
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	// ResourceGroup is the resource group the CosmosDBSQLDatabase will be created in.
	ResourceGroup string `json:"resourceGroup"`

	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	// CosmosDBAccount is the account that the SQL database will be created in.
	CosmosDBAccount string `json:"cosmosDBAccount"`

	// +kubebuilder:validation:Min=400
	// Throughput is the user specified manual throughput (RU/s) for the database expressed in units of 100 request
	// units per second. The minimum is 400 up to 1,000,000 (or higher by requesting a limit increase).
	// This must not be specified if autoscale is specified.
	Throughput *int32 `json:"throughput,omitempty"`

	// AutoscaleSettings contains the user specified autoscale configuration.
	AutoscaleSettings *AutoscaleSettings `json:"autoscaleSettings,omitempty"`

	// Tags are key-value pairs associated with the resource.
	Tags map[string]string `json:"tags,omitempty"`
}

type AutoscaleSettings struct {
	// +kubebuilder:validation:Min=0
	// MaxThroughput is the autoscale max RU/s of the database. This must not be specified if Throughput is specified.
	MaxThroughput *int32 `json:"maxThroughput,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// CosmosDBSQLDatabase is the Schema for the cosmosdbsql API
// +kubebuilder:resource:shortName=cdbsql
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type CosmosDBSQLDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CosmosDBSQLDatabaseSpec `json:"spec,omitempty"`
	Status ASOStatus               `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// CosmosDBSQLList contains a list of CosmosDBSQLDatabase
type CosmosDBSQLDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CosmosDBSQLDatabase `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CosmosDBSQLDatabase{}, &CosmosDBSQLDatabaseList{})
}
