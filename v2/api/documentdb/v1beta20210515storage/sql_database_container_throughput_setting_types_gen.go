// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=documentdb.azure.com,resources=sqldatabasecontainerthroughputsettings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=documentdb.azure.com,resources={sqldatabasecontainerthroughputsettings/status,sqldatabasecontainerthroughputsettings/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20210515.SqlDatabaseContainerThroughputSetting
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/throughputSettings/default
type SqlDatabaseContainerThroughputSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccounts_SqlDatabases_Containers_ThroughputSetting_Spec   `json:"spec,omitempty"`
	Status            DatabaseAccounts_SqlDatabases_Containers_ThroughputSetting_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseContainerThroughputSetting{}

// GetConditions returns the conditions of the resource
func (setting *SqlDatabaseContainerThroughputSetting) GetConditions() conditions.Conditions {
	return setting.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (setting *SqlDatabaseContainerThroughputSetting) SetConditions(conditions conditions.Conditions) {
	setting.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &SqlDatabaseContainerThroughputSetting{}

// AzureName returns the Azure name of the resource (always "default")
func (setting *SqlDatabaseContainerThroughputSetting) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (setting SqlDatabaseContainerThroughputSetting) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (setting *SqlDatabaseContainerThroughputSetting) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (setting *SqlDatabaseContainerThroughputSetting) GetSpec() genruntime.ConvertibleSpec {
	return &setting.Spec
}

// GetStatus returns the status of this resource
func (setting *SqlDatabaseContainerThroughputSetting) GetStatus() genruntime.ConvertibleStatus {
	return &setting.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/throughputSettings"
func (setting *SqlDatabaseContainerThroughputSetting) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/throughputSettings"
}

// NewEmptyStatus returns a new empty (blank) status
func (setting *SqlDatabaseContainerThroughputSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DatabaseAccounts_SqlDatabases_Containers_ThroughputSetting_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (setting *SqlDatabaseContainerThroughputSetting) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(setting.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  setting.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (setting *SqlDatabaseContainerThroughputSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DatabaseAccounts_SqlDatabases_Containers_ThroughputSetting_STATUS); ok {
		setting.Status = *st
		return nil
	}

	// Convert status to required version
	var st DatabaseAccounts_SqlDatabases_Containers_ThroughputSetting_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	setting.Status = st
	return nil
}

// Hub marks that this SqlDatabaseContainerThroughputSetting is the hub type for conversion
func (setting *SqlDatabaseContainerThroughputSetting) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (setting *SqlDatabaseContainerThroughputSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: setting.Spec.OriginalVersion,
		Kind:    "SqlDatabaseContainerThroughputSetting",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20210515.SqlDatabaseContainerThroughputSetting
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/throughputSettings/default
type SqlDatabaseContainerThroughputSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseContainerThroughputSetting `json:"items"`
}

// Storage version of v1beta20210515.DatabaseAccounts_SqlDatabases_Containers_ThroughputSetting_Spec
type DatabaseAccounts_SqlDatabases_Containers_ThroughputSetting_Spec struct {
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/SqlDatabaseContainer resource
	Owner       *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"SqlDatabaseContainer"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Resource    *ThroughputSettingsResource        `json:"resource,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccounts_SqlDatabases_Containers_ThroughputSetting_Spec{}

// ConvertSpecFrom populates our DatabaseAccounts_SqlDatabases_Containers_ThroughputSetting_Spec from the provided source
func (setting *DatabaseAccounts_SqlDatabases_Containers_ThroughputSetting_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == setting {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(setting)
}

// ConvertSpecTo populates the provided destination from our DatabaseAccounts_SqlDatabases_Containers_ThroughputSetting_Spec
func (setting *DatabaseAccounts_SqlDatabases_Containers_ThroughputSetting_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == setting {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(setting)
}

// Storage version of v1beta20210515.DatabaseAccounts_SqlDatabases_Containers_ThroughputSetting_STATUS
type DatabaseAccounts_SqlDatabases_Containers_ThroughputSetting_STATUS struct {
	Conditions  []conditions.Condition                           `json:"conditions,omitempty"`
	Id          *string                                          `json:"id,omitempty"`
	Location    *string                                          `json:"location,omitempty"`
	Name        *string                                          `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag                           `json:"$propertyBag,omitempty"`
	Resource    *ThroughputSettingsGetProperties_Resource_STATUS `json:"resource,omitempty"`
	Tags        map[string]string                                `json:"tags,omitempty"`
	Type        *string                                          `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DatabaseAccounts_SqlDatabases_Containers_ThroughputSetting_STATUS{}

// ConvertStatusFrom populates our DatabaseAccounts_SqlDatabases_Containers_ThroughputSetting_STATUS from the provided source
func (setting *DatabaseAccounts_SqlDatabases_Containers_ThroughputSetting_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == setting {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(setting)
}

// ConvertStatusTo populates the provided destination from our DatabaseAccounts_SqlDatabases_Containers_ThroughputSetting_STATUS
func (setting *DatabaseAccounts_SqlDatabases_Containers_ThroughputSetting_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == setting {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(setting)
}

func init() {
	SchemeBuilder.Register(&SqlDatabaseContainerThroughputSetting{}, &SqlDatabaseContainerThroughputSettingList{})
}
