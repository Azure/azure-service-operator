// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-service-operator/pkg/helpers"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type DBEdition byte

const (
	SqlDatabaseEditionBasic          = DBEdition(0)
	SqlDatabaseEditionDataWarehouse  = DBEdition(3)
	SqlDatabaseEditionFree           = DBEdition(4)
	SqlDatabaseEditionGeneralPurpose = DBEdition(5)
	SqlDatabaseEditionHyperscale     = DBEdition(6)
	SqlDatabaseEditionPremium        = DBEdition(7)
	SqlDatabaseEditionStandard       = DBEdition(9)
	SqlDatabaseEditionStretch        = DBEdition(10)
)

type SqlDatabaseSku struct {
	// Name - The name of the SKU, typically, a letter + Number code, e.g. P3.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`
	// optional
	// Tier - The tier or edition of the particular SKU, e.g. Basic, Premium.
	Tier string `json:"tier,omitempty"`
	// Size - Size of the particular SKU
	Size string `json:"size,omitempty"`
	// Family - If the service has different generations of hardware, for the same SKU, then that can be captured here.
	Family string `json:"family,omitempty"`
	// Capacity - Capacity of the particular SKU.
	Capacity *int32 `json:"capacity,omitempty"`
}

type SQLDatabaseShortTermRetentionPolicy struct {
	// RetentionDays is the backup retention period in days. This is how many days
	// Point-in-Time Restore will be supported.
	// +kubebuilder:validation:Required
	RetentionDays int32 `json:"retentionDays"`
}

// AzureSqlDatabaseSpec defines the desired state of AzureSqlDatabase
type AzureSqlDatabaseSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Location string `json:"location"`

	// +kubebuilder:validation:Pattern=^[-\w\._\(\)]+$
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	ResourceGroup string `json:"resourceGroup"`

	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Server string `json:"server"`

	// +kubebuilder:validation:Optional
	SubscriptionID           string                               `json:"subscriptionId,omitempty"`
	Edition                  DBEdition                            `json:"edition"`       // TODO: Remove this in v1beta2
	Sku                      *SqlDatabaseSku                      `json:"sku,omitempty"` // TODO: make this required in v1beta2
	MaxSize                  *resource.Quantity                   `json:"maxSize,omitempty"`
	DbName                   string                               `json:"dbName,omitempty"`
	WeeklyRetention          string                               `json:"weeklyRetention,omitempty"`
	MonthlyRetention         string                               `json:"monthlyRetention,omitempty"`
	YearlyRetention          string                               `json:"yearlyRetention,omitempty"`
	WeekOfYear               int32                                `json:"weekOfYear,omitempty"`
	ShortTermRetentionPolicy *SQLDatabaseShortTermRetentionPolicy `json:"shortTermRetentionPolicy,omitempty"`
	ElasticPoolID            string                               `json:"elasticPoolId,omitempty"`
}

// AzureSqlDatabase is the Schema for the azuresqldatabases API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:shortName=asqldb
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

func init() {
	SchemeBuilder.Register(&AzureSqlDatabase{}, &AzureSqlDatabaseList{})
}

func (s *AzureSqlDatabase) IsSubmitted() bool {
	return s.Status.Provisioned
}

func (s *AzureSqlDatabase) HasFinalizer(finalizerName string) bool {
	return helpers.ContainsString(s.ObjectMeta.Finalizers, finalizerName)
}
