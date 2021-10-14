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
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210515.SqlDatabaseContainerUserDefinedFunction
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_containers_userDefinedFunctions
type SqlDatabaseContainerUserDefinedFunction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec `json:"spec,omitempty"`
	Status            SqlUserDefinedFunctionGetResults_Status                         `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseContainerUserDefinedFunction{}

// GetConditions returns the conditions of the resource
func (sqlDatabaseContainerUserDefinedFunction *SqlDatabaseContainerUserDefinedFunction) GetConditions() conditions.Conditions {
	return sqlDatabaseContainerUserDefinedFunction.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (sqlDatabaseContainerUserDefinedFunction *SqlDatabaseContainerUserDefinedFunction) SetConditions(conditions conditions.Conditions) {
	sqlDatabaseContainerUserDefinedFunction.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &SqlDatabaseContainerUserDefinedFunction{}

// AzureName returns the Azure name of the resource
func (sqlDatabaseContainerUserDefinedFunction *SqlDatabaseContainerUserDefinedFunction) AzureName() string {
	return sqlDatabaseContainerUserDefinedFunction.Spec.AzureName
}

// GetResourceKind returns the kind of the resource
func (sqlDatabaseContainerUserDefinedFunction *SqlDatabaseContainerUserDefinedFunction) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (sqlDatabaseContainerUserDefinedFunction *SqlDatabaseContainerUserDefinedFunction) GetSpec() genruntime.ConvertibleSpec {
	return &sqlDatabaseContainerUserDefinedFunction.Spec
}

// GetStatus returns the status of this resource
func (sqlDatabaseContainerUserDefinedFunction *SqlDatabaseContainerUserDefinedFunction) GetStatus() genruntime.ConvertibleStatus {
	return &sqlDatabaseContainerUserDefinedFunction.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/userDefinedFunctions"
func (sqlDatabaseContainerUserDefinedFunction *SqlDatabaseContainerUserDefinedFunction) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/userDefinedFunctions"
}

// NewEmptyStatus returns a new empty (blank) status
func (sqlDatabaseContainerUserDefinedFunction *SqlDatabaseContainerUserDefinedFunction) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SqlUserDefinedFunctionGetResults_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (sqlDatabaseContainerUserDefinedFunction *SqlDatabaseContainerUserDefinedFunction) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(sqlDatabaseContainerUserDefinedFunction.Spec)
	return &genruntime.ResourceReference{
		Group:     group,
		Kind:      kind,
		Namespace: sqlDatabaseContainerUserDefinedFunction.Namespace,
		Name:      sqlDatabaseContainerUserDefinedFunction.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (sqlDatabaseContainerUserDefinedFunction *SqlDatabaseContainerUserDefinedFunction) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SqlUserDefinedFunctionGetResults_Status); ok {
		sqlDatabaseContainerUserDefinedFunction.Status = *st
		return nil
	}

	// Convert status to required version
	var st SqlUserDefinedFunctionGetResults_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	sqlDatabaseContainerUserDefinedFunction.Status = st
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (sqlDatabaseContainerUserDefinedFunction *SqlDatabaseContainerUserDefinedFunction) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: sqlDatabaseContainerUserDefinedFunction.Spec.OriginalVersion,
		Kind:    "SqlDatabaseContainerUserDefinedFunction",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210515.SqlDatabaseContainerUserDefinedFunction
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_containers_userDefinedFunctions
type SqlDatabaseContainerUserDefinedFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseContainerUserDefinedFunction `json:"items"`
}

//Storage version of v1alpha1api20210515.DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec
type DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName       string               `json:"azureName"`
	Location        *string              `json:"location,omitempty"`
	Options         *CreateUpdateOptions `json:"options,omitempty"`
	OriginalVersion string               `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"microsoft.documentdb.azure.com" json:"owner" kind:"SqlDatabaseContainer"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Resource    *SqlUserDefinedFunctionResource   `json:"resource,omitempty"`
	Tags        map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec from the provided source
func (databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec *DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec)
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec
func (databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec *DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec)
}

//Storage version of v1alpha1api20210515.SqlUserDefinedFunctionGetResults_Status
//Generated from:
type SqlUserDefinedFunctionGetResults_Status struct {
	Conditions  []conditions.Condition                               `json:"conditions,omitempty"`
	Id          *string                                              `json:"id,omitempty"`
	Location    *string                                              `json:"location,omitempty"`
	Name        *string                                              `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag                               `json:"$propertyBag,omitempty"`
	Resource    *SqlUserDefinedFunctionGetProperties_Status_Resource `json:"resource,omitempty"`
	Tags        map[string]string                                    `json:"tags,omitempty"`
	Type        *string                                              `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SqlUserDefinedFunctionGetResults_Status{}

// ConvertStatusFrom populates our SqlUserDefinedFunctionGetResults_Status from the provided source
func (sqlUserDefinedFunctionGetResultsStatus *SqlUserDefinedFunctionGetResults_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == sqlUserDefinedFunctionGetResultsStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(sqlUserDefinedFunctionGetResultsStatus)
}

// ConvertStatusTo populates the provided destination from our SqlUserDefinedFunctionGetResults_Status
func (sqlUserDefinedFunctionGetResultsStatus *SqlUserDefinedFunctionGetResults_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == sqlUserDefinedFunctionGetResultsStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(sqlUserDefinedFunctionGetResultsStatus)
}

//Storage version of v1alpha1api20210515.SqlUserDefinedFunctionGetProperties_Status_Resource
type SqlUserDefinedFunctionGetProperties_Status_Resource struct {
	Body        *string                `json:"body,omitempty"`
	Etag        *string                `json:"_etag,omitempty"`
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rid         *string                `json:"_rid,omitempty"`
	Ts          *float64               `json:"_ts,omitempty"`
}

//Storage version of v1alpha1api20210515.SqlUserDefinedFunctionResource
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlUserDefinedFunctionResource
type SqlUserDefinedFunctionResource struct {
	Body        *string                `json:"body,omitempty"`
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&SqlDatabaseContainerUserDefinedFunction{}, &SqlDatabaseContainerUserDefinedFunctionList{})
}
