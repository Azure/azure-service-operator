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
//Storage version of v1alpha1api20210515.MongodbDatabaseCollection
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_mongodbDatabases_collections
type MongodbDatabaseCollection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsMongodbDatabasesCollections_Spec `json:"spec,omitempty"`
	Status            MongoDBCollectionGetResults_Status               `json:"status,omitempty"`
}

var _ conditions.Conditioner = &MongodbDatabaseCollection{}

// GetConditions returns the conditions of the resource
func (mongodbDatabaseCollection *MongodbDatabaseCollection) GetConditions() conditions.Conditions {
	return mongodbDatabaseCollection.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (mongodbDatabaseCollection *MongodbDatabaseCollection) SetConditions(conditions conditions.Conditions) {
	mongodbDatabaseCollection.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &MongodbDatabaseCollection{}

// AzureName returns the Azure name of the resource
func (mongodbDatabaseCollection *MongodbDatabaseCollection) AzureName() string {
	return mongodbDatabaseCollection.Spec.AzureName
}

// GetResourceKind returns the kind of the resource
func (mongodbDatabaseCollection *MongodbDatabaseCollection) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (mongodbDatabaseCollection *MongodbDatabaseCollection) GetSpec() genruntime.ConvertibleSpec {
	return &mongodbDatabaseCollection.Spec
}

// GetStatus returns the status of this resource
func (mongodbDatabaseCollection *MongodbDatabaseCollection) GetStatus() genruntime.ConvertibleStatus {
	return &mongodbDatabaseCollection.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/collections"
func (mongodbDatabaseCollection *MongodbDatabaseCollection) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/collections"
}

// NewEmptyStatus returns a new empty (blank) status
func (mongodbDatabaseCollection *MongodbDatabaseCollection) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &MongoDBCollectionGetResults_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (mongodbDatabaseCollection *MongodbDatabaseCollection) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(mongodbDatabaseCollection.Spec)
	return &genruntime.ResourceReference{
		Group:     group,
		Kind:      kind,
		Namespace: mongodbDatabaseCollection.Namespace,
		Name:      mongodbDatabaseCollection.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (mongodbDatabaseCollection *MongodbDatabaseCollection) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*MongoDBCollectionGetResults_Status); ok {
		mongodbDatabaseCollection.Status = *st
		return nil
	}

	// Convert status to required version
	var st MongoDBCollectionGetResults_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	mongodbDatabaseCollection.Status = st
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (mongodbDatabaseCollection *MongodbDatabaseCollection) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: mongodbDatabaseCollection.Spec.OriginalVersion,
		Kind:    "MongodbDatabaseCollection",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210515.MongodbDatabaseCollection
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_mongodbDatabases_collections
type MongodbDatabaseCollectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongodbDatabaseCollection `json:"items"`
}

//Storage version of v1alpha1api20210515.DatabaseAccountsMongodbDatabasesCollections_Spec
type DatabaseAccountsMongodbDatabasesCollections_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName       string               `json:"azureName"`
	Location        *string              `json:"location,omitempty"`
	Options         *CreateUpdateOptions `json:"options,omitempty"`
	OriginalVersion string               `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"microsoft.documentdb.azure.com" json:"owner" kind:"MongodbDatabase"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Resource    *MongoDBCollectionResource        `json:"resource,omitempty"`
	Tags        map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsMongodbDatabasesCollections_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsMongodbDatabasesCollections_Spec from the provided source
func (databaseAccountsMongodbDatabasesCollectionsSpec *DatabaseAccountsMongodbDatabasesCollections_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == databaseAccountsMongodbDatabasesCollectionsSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(databaseAccountsMongodbDatabasesCollectionsSpec)
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsMongodbDatabasesCollections_Spec
func (databaseAccountsMongodbDatabasesCollectionsSpec *DatabaseAccountsMongodbDatabasesCollections_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == databaseAccountsMongodbDatabasesCollectionsSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(databaseAccountsMongodbDatabasesCollectionsSpec)
}

//Storage version of v1alpha1api20210515.MongoDBCollectionGetResults_Status
//Generated from:
type MongoDBCollectionGetResults_Status struct {
	Conditions  []conditions.Condition                          `json:"conditions,omitempty"`
	Id          *string                                         `json:"id,omitempty"`
	Location    *string                                         `json:"location,omitempty"`
	Name        *string                                         `json:"name,omitempty"`
	Options     *OptionsResource_Status                         `json:"options,omitempty"`
	PropertyBag genruntime.PropertyBag                          `json:"$propertyBag,omitempty"`
	Resource    *MongoDBCollectionGetProperties_Status_Resource `json:"resource,omitempty"`
	Tags        map[string]string                               `json:"tags,omitempty"`
	Type        *string                                         `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &MongoDBCollectionGetResults_Status{}

// ConvertStatusFrom populates our MongoDBCollectionGetResults_Status from the provided source
func (mongoDBCollectionGetResultsStatus *MongoDBCollectionGetResults_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == mongoDBCollectionGetResultsStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(mongoDBCollectionGetResultsStatus)
}

// ConvertStatusTo populates the provided destination from our MongoDBCollectionGetResults_Status
func (mongoDBCollectionGetResultsStatus *MongoDBCollectionGetResults_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == mongoDBCollectionGetResultsStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(mongoDBCollectionGetResultsStatus)
}

//Storage version of v1alpha1api20210515.MongoDBCollectionGetProperties_Status_Resource
type MongoDBCollectionGetProperties_Status_Resource struct {
	AnalyticalStorageTtl *int                   `json:"analyticalStorageTtl,omitempty"`
	Etag                 *string                `json:"_etag,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	Indexes              []MongoIndex_Status    `json:"indexes,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rid                  *string                `json:"_rid,omitempty"`
	ShardKey             map[string]string      `json:"shardKey,omitempty"`
	Ts                   *float64               `json:"_ts,omitempty"`
}

//Storage version of v1alpha1api20210515.MongoDBCollectionResource
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoDBCollectionResource
type MongoDBCollectionResource struct {
	AnalyticalStorageTtl *int                   `json:"analyticalStorageTtl,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	Indexes              []MongoIndex           `json:"indexes,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ShardKey             map[string]string      `json:"shardKey,omitempty"`
}

//Storage version of v1alpha1api20210515.MongoIndex
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoIndex
type MongoIndex struct {
	Key         *MongoIndexKeys        `json:"key,omitempty"`
	Options     *MongoIndexOptions     `json:"options,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.MongoIndex_Status
//Generated from:
type MongoIndex_Status struct {
	Key         *MongoIndexKeys_Status    `json:"key,omitempty"`
	Options     *MongoIndexOptions_Status `json:"options,omitempty"`
	PropertyBag genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.MongoIndexKeys
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoIndexKeys
type MongoIndexKeys struct {
	Keys        []string               `json:"keys,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.MongoIndexKeys_Status
//Generated from:
type MongoIndexKeys_Status struct {
	Keys        []string               `json:"keys,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.MongoIndexOptions
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/MongoIndexOptions
type MongoIndexOptions struct {
	ExpireAfterSeconds *int                   `json:"expireAfterSeconds,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Unique             *bool                  `json:"unique,omitempty"`
}

//Storage version of v1alpha1api20210515.MongoIndexOptions_Status
//Generated from:
type MongoIndexOptions_Status struct {
	ExpireAfterSeconds *int                   `json:"expireAfterSeconds,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Unique             *bool                  `json:"unique,omitempty"`
}

func init() {
	SchemeBuilder.Register(&MongodbDatabaseCollection{}, &MongodbDatabaseCollectionList{})
}
