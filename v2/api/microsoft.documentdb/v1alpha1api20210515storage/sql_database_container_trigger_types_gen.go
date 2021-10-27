// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210515.SqlDatabaseContainerTrigger
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_containers_triggers
type SqlDatabaseContainerTrigger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsSqlDatabasesContainersTriggers_Spec `json:"spec,omitempty"`
	Status            SqlTriggerGetResults_Status                         `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseContainerTrigger{}

// GetConditions returns the conditions of the resource
func (sqlDatabaseContainerTrigger *SqlDatabaseContainerTrigger) GetConditions() conditions.Conditions {
	return sqlDatabaseContainerTrigger.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (sqlDatabaseContainerTrigger *SqlDatabaseContainerTrigger) SetConditions(conditions conditions.Conditions) {
	sqlDatabaseContainerTrigger.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &SqlDatabaseContainerTrigger{}

// AzureName returns the Azure name of the resource
func (sqlDatabaseContainerTrigger *SqlDatabaseContainerTrigger) AzureName() string {
	return sqlDatabaseContainerTrigger.Spec.AzureName
}

// GetResourceKind returns the kind of the resource
func (sqlDatabaseContainerTrigger *SqlDatabaseContainerTrigger) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (sqlDatabaseContainerTrigger *SqlDatabaseContainerTrigger) GetSpec() genruntime.ConvertibleSpec {
	return &sqlDatabaseContainerTrigger.Spec
}

// GetStatus returns the status of this resource
func (sqlDatabaseContainerTrigger *SqlDatabaseContainerTrigger) GetStatus() genruntime.ConvertibleStatus {
	return &sqlDatabaseContainerTrigger.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/triggers"
func (sqlDatabaseContainerTrigger *SqlDatabaseContainerTrigger) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/triggers"
}

// NewEmptyStatus returns a new empty (blank) status
func (sqlDatabaseContainerTrigger *SqlDatabaseContainerTrigger) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SqlTriggerGetResults_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (sqlDatabaseContainerTrigger *SqlDatabaseContainerTrigger) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(sqlDatabaseContainerTrigger.Spec)
	return &genruntime.ResourceReference{
		Group:     group,
		Kind:      kind,
		Namespace: sqlDatabaseContainerTrigger.Namespace,
		Name:      sqlDatabaseContainerTrigger.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (sqlDatabaseContainerTrigger *SqlDatabaseContainerTrigger) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SqlTriggerGetResults_Status); ok {
		sqlDatabaseContainerTrigger.Status = *st
		return nil
	}

	// Convert status to required version
	var st SqlTriggerGetResults_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	sqlDatabaseContainerTrigger.Status = st
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (sqlDatabaseContainerTrigger *SqlDatabaseContainerTrigger) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: sqlDatabaseContainerTrigger.Spec.OriginalVersion,
		Kind:    "SqlDatabaseContainerTrigger",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210515.SqlDatabaseContainerTrigger
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_containers_triggers
type SqlDatabaseContainerTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseContainerTrigger `json:"items"`
}

//Storage version of v1alpha1api20210515.DatabaseAccountsSqlDatabasesContainersTriggers_Spec
type DatabaseAccountsSqlDatabasesContainersTriggers_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName       string               `json:"azureName"`
	Location        *string              `json:"location,omitempty"`
	Options         *CreateUpdateOptions `json:"options,omitempty"`
	OriginalVersion string               `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"microsoft.documentdb.azure.com" json:"owner" kind:"SqlDatabaseContainer"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Resource    *SqlTriggerResource               `json:"resource,omitempty"`
	Tags        map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsSqlDatabasesContainersTriggers_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsSqlDatabasesContainersTriggers_Spec from the provided source
func (databaseAccountsSqlDatabasesContainersTriggersSpec *DatabaseAccountsSqlDatabasesContainersTriggers_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == databaseAccountsSqlDatabasesContainersTriggersSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(databaseAccountsSqlDatabasesContainersTriggersSpec)
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsSqlDatabasesContainersTriggers_Spec
func (databaseAccountsSqlDatabasesContainersTriggersSpec *DatabaseAccountsSqlDatabasesContainersTriggers_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == databaseAccountsSqlDatabasesContainersTriggersSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(databaseAccountsSqlDatabasesContainersTriggersSpec)
}

//Storage version of v1alpha1api20210515.SqlTriggerGetResults_Status
//Generated from:
type SqlTriggerGetResults_Status struct {
	Conditions  []conditions.Condition                   `json:"conditions,omitempty"`
	Id          *string                                  `json:"id,omitempty"`
	Location    *string                                  `json:"location,omitempty"`
	Name        *string                                  `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag                   `json:"$propertyBag,omitempty"`
	Resource    *SqlTriggerGetProperties_Status_Resource `json:"resource,omitempty"`
	Tags        map[string]string                        `json:"tags,omitempty"`
	Type        *string                                  `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SqlTriggerGetResults_Status{}

// ConvertStatusFrom populates our SqlTriggerGetResults_Status from the provided source
func (sqlTriggerGetResultsStatus *SqlTriggerGetResults_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == sqlTriggerGetResultsStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(sqlTriggerGetResultsStatus)
}

// ConvertStatusTo populates the provided destination from our SqlTriggerGetResults_Status
func (sqlTriggerGetResultsStatus *SqlTriggerGetResults_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == sqlTriggerGetResultsStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(sqlTriggerGetResultsStatus)
}

//Storage version of v1alpha1api20210515.SqlTriggerGetProperties_Status_Resource
type SqlTriggerGetProperties_Status_Resource struct {
	Body             *string                `json:"body,omitempty"`
	Etag             *string                `json:"_etag,omitempty"`
	Id               *string                `json:"id,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rid              *string                `json:"_rid,omitempty"`
	TriggerOperation *string                `json:"triggerOperation,omitempty"`
	TriggerType      *string                `json:"triggerType,omitempty"`
	Ts               *float64               `json:"_ts,omitempty"`
}

//Storage version of v1alpha1api20210515.SqlTriggerResource
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlTriggerResource
type SqlTriggerResource struct {
	Body             *string                `json:"body,omitempty"`
	Id               *string                `json:"id,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TriggerOperation *string                `json:"triggerOperation,omitempty"`
	TriggerType      *string                `json:"triggerType,omitempty"`
}

func init() {
	SchemeBuilder.Register(&SqlDatabaseContainerTrigger{}, &SqlDatabaseContainerTriggerList{})
}
