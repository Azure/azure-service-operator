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
//Storage version of v1alpha1api20210515.MongodbDatabaseCollectionThroughputSetting
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_mongodbDatabases_collections_throughputSettings
type MongodbDatabaseCollectionThroughputSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_Spec `json:"spec,omitempty"`
	Status            ThroughputSettingsGetResults_Status                                `json:"status,omitempty"`
}

var _ conditions.Conditioner = &MongodbDatabaseCollectionThroughputSetting{}

// GetConditions returns the conditions of the resource
func (mongodbDatabaseCollectionThroughputSetting *MongodbDatabaseCollectionThroughputSetting) GetConditions() conditions.Conditions {
	return mongodbDatabaseCollectionThroughputSetting.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (mongodbDatabaseCollectionThroughputSetting *MongodbDatabaseCollectionThroughputSetting) SetConditions(conditions conditions.Conditions) {
	mongodbDatabaseCollectionThroughputSetting.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &MongodbDatabaseCollectionThroughputSetting{}

// AzureName returns the Azure name of the resource (always "default")
func (mongodbDatabaseCollectionThroughputSetting *MongodbDatabaseCollectionThroughputSetting) AzureName() string {
	return "default"
}

// GetResourceKind returns the kind of the resource
func (mongodbDatabaseCollectionThroughputSetting *MongodbDatabaseCollectionThroughputSetting) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (mongodbDatabaseCollectionThroughputSetting *MongodbDatabaseCollectionThroughputSetting) GetSpec() genruntime.ConvertibleSpec {
	return &mongodbDatabaseCollectionThroughputSetting.Spec
}

// GetStatus returns the status of this resource
func (mongodbDatabaseCollectionThroughputSetting *MongodbDatabaseCollectionThroughputSetting) GetStatus() genruntime.ConvertibleStatus {
	return &mongodbDatabaseCollectionThroughputSetting.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/collections/throughputSettings"
func (mongodbDatabaseCollectionThroughputSetting *MongodbDatabaseCollectionThroughputSetting) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/collections/throughputSettings"
}

// NewEmptyStatus returns a new empty (blank) status
func (mongodbDatabaseCollectionThroughputSetting *MongodbDatabaseCollectionThroughputSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ThroughputSettingsGetResults_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (mongodbDatabaseCollectionThroughputSetting *MongodbDatabaseCollectionThroughputSetting) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(mongodbDatabaseCollectionThroughputSetting.Spec)
	return &genruntime.ResourceReference{
		Group:     group,
		Kind:      kind,
		Namespace: mongodbDatabaseCollectionThroughputSetting.Namespace,
		Name:      mongodbDatabaseCollectionThroughputSetting.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (mongodbDatabaseCollectionThroughputSetting *MongodbDatabaseCollectionThroughputSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ThroughputSettingsGetResults_Status); ok {
		mongodbDatabaseCollectionThroughputSetting.Status = *st
		return nil
	}

	// Convert status to required version
	var st ThroughputSettingsGetResults_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	mongodbDatabaseCollectionThroughputSetting.Status = st
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (mongodbDatabaseCollectionThroughputSetting *MongodbDatabaseCollectionThroughputSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: mongodbDatabaseCollectionThroughputSetting.Spec.OriginalVersion,
		Kind:    "MongodbDatabaseCollectionThroughputSetting",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210515.MongodbDatabaseCollectionThroughputSetting
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_mongodbDatabases_collections_throughputSettings
type MongodbDatabaseCollectionThroughputSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongodbDatabaseCollectionThroughputSetting `json:"items"`
}

//Storage version of v1alpha1api20210515.DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_Spec
type DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_Spec struct {
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"microsoft.documentdb.azure.com" json:"owner" kind:"MongodbDatabaseCollection"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Resource    *ThroughputSettingsResource       `json:"resource,omitempty"`
	Tags        map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_Spec from the provided source
func (databaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpec *DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == databaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(databaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpec)
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_Spec
func (databaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpec *DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == databaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(databaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpec)
}

//Storage version of v1alpha1api20210515.ThroughputSettingsGetResults_Status
//Generated from:
type ThroughputSettingsGetResults_Status struct {
	Conditions  []conditions.Condition                           `json:"conditions,omitempty"`
	Id          *string                                          `json:"id,omitempty"`
	Location    *string                                          `json:"location,omitempty"`
	Name        *string                                          `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag                           `json:"$propertyBag,omitempty"`
	Resource    *ThroughputSettingsGetProperties_Status_Resource `json:"resource,omitempty"`
	Tags        map[string]string                                `json:"tags,omitempty"`
	Type        *string                                          `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &ThroughputSettingsGetResults_Status{}

// ConvertStatusFrom populates our ThroughputSettingsGetResults_Status from the provided source
func (throughputSettingsGetResultsStatus *ThroughputSettingsGetResults_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == throughputSettingsGetResultsStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(throughputSettingsGetResultsStatus)
}

// ConvertStatusTo populates the provided destination from our ThroughputSettingsGetResults_Status
func (throughputSettingsGetResultsStatus *ThroughputSettingsGetResults_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == throughputSettingsGetResultsStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(throughputSettingsGetResultsStatus)
}

//Storage version of v1alpha1api20210515.ThroughputSettingsGetProperties_Status_Resource
type ThroughputSettingsGetProperties_Status_Resource struct {
	AutoscaleSettings   *AutoscaleSettingsResource_Status `json:"autoscaleSettings,omitempty"`
	Etag                *string                           `json:"_etag,omitempty"`
	MinimumThroughput   *string                           `json:"minimumThroughput,omitempty"`
	OfferReplacePending *string                           `json:"offerReplacePending,omitempty"`
	PropertyBag         genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Rid                 *string                           `json:"_rid,omitempty"`
	Throughput          *int                              `json:"throughput,omitempty"`
	Ts                  *float64                          `json:"_ts,omitempty"`
}

//Storage version of v1alpha1api20210515.ThroughputSettingsResource
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ThroughputSettingsResource
type ThroughputSettingsResource struct {
	AutoscaleSettings *AutoscaleSettingsResource `json:"autoscaleSettings,omitempty"`
	PropertyBag       genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	Throughput        *int                       `json:"throughput,omitempty"`
}

//Storage version of v1alpha1api20210515.AutoscaleSettingsResource
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/AutoscaleSettingsResource
type AutoscaleSettingsResource struct {
	AutoUpgradePolicy *AutoUpgradePolicyResource `json:"autoUpgradePolicy,omitempty"`
	MaxThroughput     *int                       `json:"maxThroughput,omitempty"`
	PropertyBag       genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.AutoscaleSettingsResource_Status
//Generated from:
type AutoscaleSettingsResource_Status struct {
	AutoUpgradePolicy   *AutoUpgradePolicyResource_Status `json:"autoUpgradePolicy,omitempty"`
	MaxThroughput       *int                              `json:"maxThroughput,omitempty"`
	PropertyBag         genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	TargetMaxThroughput *int                              `json:"targetMaxThroughput,omitempty"`
}

//Storage version of v1alpha1api20210515.AutoUpgradePolicyResource
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/AutoUpgradePolicyResource
type AutoUpgradePolicyResource struct {
	PropertyBag      genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
	ThroughputPolicy *ThroughputPolicyResource `json:"throughputPolicy,omitempty"`
}

//Storage version of v1alpha1api20210515.AutoUpgradePolicyResource_Status
//Generated from:
type AutoUpgradePolicyResource_Status struct {
	PropertyBag      genruntime.PropertyBag           `json:"$propertyBag,omitempty"`
	ThroughputPolicy *ThroughputPolicyResource_Status `json:"throughputPolicy,omitempty"`
}

//Storage version of v1alpha1api20210515.ThroughputPolicyResource
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ThroughputPolicyResource
type ThroughputPolicyResource struct {
	IncrementPercent *int                   `json:"incrementPercent,omitempty"`
	IsEnabled        *bool                  `json:"isEnabled,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.ThroughputPolicyResource_Status
//Generated from:
type ThroughputPolicyResource_Status struct {
	IncrementPercent *int                   `json:"incrementPercent,omitempty"`
	IsEnabled        *bool                  `json:"isEnabled,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&MongodbDatabaseCollectionThroughputSetting{}, &MongodbDatabaseCollectionThroughputSettingList{})
}
