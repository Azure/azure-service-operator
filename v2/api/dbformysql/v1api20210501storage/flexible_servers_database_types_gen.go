// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210501storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=dbformysql.azure.com,resources=flexibleserversdatabases,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=dbformysql.azure.com,resources={flexibleserversdatabases/status,flexibleserversdatabases/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20210501.FlexibleServersDatabase
// Generator information:
// - Generated from: /mysql/resource-manager/Microsoft.DBforMySQL/legacy/stable/2021-05-01/mysql.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/flexibleServers/{serverName}/databases/{databaseName}
type FlexibleServersDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServers_Database_Spec   `json:"spec,omitempty"`
	Status            FlexibleServers_Database_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &FlexibleServersDatabase{}

// GetConditions returns the conditions of the resource
func (database *FlexibleServersDatabase) GetConditions() conditions.Conditions {
	return database.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (database *FlexibleServersDatabase) SetConditions(conditions conditions.Conditions) {
	database.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &FlexibleServersDatabase{}

// AzureName returns the Azure name of the resource
func (database *FlexibleServersDatabase) AzureName() string {
	return database.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-01"
func (database FlexibleServersDatabase) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (database *FlexibleServersDatabase) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (database *FlexibleServersDatabase) GetSpec() genruntime.ConvertibleSpec {
	return &database.Spec
}

// GetStatus returns the status of this resource
func (database *FlexibleServersDatabase) GetStatus() genruntime.ConvertibleStatus {
	return &database.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (database *FlexibleServersDatabase) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMySQL/flexibleServers/databases"
func (database *FlexibleServersDatabase) GetType() string {
	return "Microsoft.DBforMySQL/flexibleServers/databases"
}

// NewEmptyStatus returns a new empty (blank) status
func (database *FlexibleServersDatabase) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &FlexibleServers_Database_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (database *FlexibleServersDatabase) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(database.Spec)
	return database.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (database *FlexibleServersDatabase) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*FlexibleServers_Database_STATUS); ok {
		database.Status = *st
		return nil
	}

	// Convert status to required version
	var st FlexibleServers_Database_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	database.Status = st
	return nil
}

// Hub marks that this FlexibleServersDatabase is the hub type for conversion
func (database *FlexibleServersDatabase) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (database *FlexibleServersDatabase) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: database.Spec.OriginalVersion,
		Kind:    "FlexibleServersDatabase",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20210501.FlexibleServersDatabase
// Generator information:
// - Generated from: /mysql/resource-manager/Microsoft.DBforMySQL/legacy/stable/2021-05-01/mysql.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/flexibleServers/{serverName}/databases/{databaseName}
type FlexibleServersDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServersDatabase `json:"items"`
}

// Storage version of v1api20210501.FlexibleServers_Database_Spec
type FlexibleServers_Database_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string  `json:"azureName,omitempty"`
	Charset         *string `json:"charset,omitempty"`
	Collation       *string `json:"collation,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a dbformysql.azure.com/FlexibleServer resource
	Owner       *genruntime.KnownResourceReference `group:"dbformysql.azure.com" json:"owner,omitempty" kind:"FlexibleServer"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

var _ genruntime.ConvertibleSpec = &FlexibleServers_Database_Spec{}

// ConvertSpecFrom populates our FlexibleServers_Database_Spec from the provided source
func (database *FlexibleServers_Database_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == database {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(database)
}

// ConvertSpecTo populates the provided destination from our FlexibleServers_Database_Spec
func (database *FlexibleServers_Database_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == database {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(database)
}

// Storage version of v1api20210501.FlexibleServers_Database_STATUS
type FlexibleServers_Database_STATUS struct {
	Charset     *string                `json:"charset,omitempty"`
	Collation   *string                `json:"collation,omitempty"`
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SystemData  *SystemData_STATUS     `json:"systemData,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &FlexibleServers_Database_STATUS{}

// ConvertStatusFrom populates our FlexibleServers_Database_STATUS from the provided source
func (database *FlexibleServers_Database_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == database {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(database)
}

// ConvertStatusTo populates the provided destination from our FlexibleServers_Database_STATUS
func (database *FlexibleServers_Database_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == database {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(database)
}

func init() {
	SchemeBuilder.Register(&FlexibleServersDatabase{}, &FlexibleServersDatabaseList{})
}
