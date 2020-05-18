// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AzureSqlDatabaseSpec defines the desired state of AzureSqlDatabase
type AzureSqlDatabaseSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Location      string    `json:"location"`
	ResourceGroup string    `json:"resourcegroup,omitempty"`
	Server        string    `json:"server"`
	Edition       DBEdition `json:"edition"`
	// optional
	DbName string `json:"dbName,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AzureSqlDatabase is the Schema for the azuresqldatabases API
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type AzureSqlDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureSqlDatabaseSpec `json:"spec,omitempty"`
	Status ASOStatus            `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureSqlDatabaseList contains a list of AzureSqlDatabase
type AzureSqlDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureSqlDatabase `json:"items"`
}

// DBEdition - wraps: https://godoc.org/github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql#DatabaseEdition
type DBEdition byte

const (
	// Basic ...
	SQLEditionBasic DBEdition = 0
	// Business ...
	SQLEditionBusiness DBEdition = 1
	// BusinessCritical ...
	SQLEditionBusinessCritical DBEdition = 2
	// DataWarehouse ...
	SQLEditionDataWarehouse DBEdition = 3
	// Free ...
	SQLEditionFree DBEdition = 4
	// GeneralPurpose ...
	SQLEditionGeneralPurpose DBEdition = 5
	// Hyperscale ...
	SQLEditionHyperscale DBEdition = 6
	// Premium ...
	SQLEditionPremium DBEdition = 7
	// PremiumRS ...
	SQLEditionPremiumRS DBEdition = 8
	// Standard ...
	SQLEditionStandard DBEdition = 9
	// Stretch ...
	SQLEditionStretch DBEdition = 10
	// System ...
	SQLEditionSystem DBEdition = 11
	// System2 ...
	SQLEditionSystem2 DBEdition = 12
	// Web ...
	SQLEditionWeb DBEdition = 13
)

func init() {
	SchemeBuilder.Register(&AzureSqlDatabase{}, &AzureSqlDatabaseList{})
}

func (s *AzureSqlDatabase) IsSubmitted() bool {
	return s.Status.Provisioned
}

func (s *AzureSqlDatabase) HasFinalizer(finalizerName string) bool {
	return helpers.ContainsString(s.ObjectMeta.Finalizers, finalizerName)
}
