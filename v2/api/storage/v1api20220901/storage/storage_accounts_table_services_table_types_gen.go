// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20230101/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220901.StorageAccountsTableServicesTable
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2022-09-01/table.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices/default/tables/{tableName}
type StorageAccountsTableServicesTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccounts_TableServices_Table_Spec   `json:"spec,omitempty"`
	Status            StorageAccounts_TableServices_Table_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &StorageAccountsTableServicesTable{}

// GetConditions returns the conditions of the resource
func (table *StorageAccountsTableServicesTable) GetConditions() conditions.Conditions {
	return table.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (table *StorageAccountsTableServicesTable) SetConditions(conditions conditions.Conditions) {
	table.Status.Conditions = conditions
}

var _ conversion.Convertible = &StorageAccountsTableServicesTable{}

// ConvertFrom populates our StorageAccountsTableServicesTable from the provided hub StorageAccountsTableServicesTable
func (table *StorageAccountsTableServicesTable) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.StorageAccountsTableServicesTable)
	if !ok {
		return fmt.Errorf("expected storage/v1api20230101/storage/StorageAccountsTableServicesTable but received %T instead", hub)
	}

	return table.AssignProperties_From_StorageAccountsTableServicesTable(source)
}

// ConvertTo populates the provided hub StorageAccountsTableServicesTable from our StorageAccountsTableServicesTable
func (table *StorageAccountsTableServicesTable) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.StorageAccountsTableServicesTable)
	if !ok {
		return fmt.Errorf("expected storage/v1api20230101/storage/StorageAccountsTableServicesTable but received %T instead", hub)
	}

	return table.AssignProperties_To_StorageAccountsTableServicesTable(destination)
}

var _ genruntime.KubernetesResource = &StorageAccountsTableServicesTable{}

// AzureName returns the Azure name of the resource
func (table *StorageAccountsTableServicesTable) AzureName() string {
	return table.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-09-01"
func (table StorageAccountsTableServicesTable) GetAPIVersion() string {
	return "2022-09-01"
}

// GetResourceScope returns the scope of the resource
func (table *StorageAccountsTableServicesTable) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (table *StorageAccountsTableServicesTable) GetSpec() genruntime.ConvertibleSpec {
	return &table.Spec
}

// GetStatus returns the status of this resource
func (table *StorageAccountsTableServicesTable) GetStatus() genruntime.ConvertibleStatus {
	return &table.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (table *StorageAccountsTableServicesTable) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/tableServices/tables"
func (table *StorageAccountsTableServicesTable) GetType() string {
	return "Microsoft.Storage/storageAccounts/tableServices/tables"
}

// NewEmptyStatus returns a new empty (blank) status
func (table *StorageAccountsTableServicesTable) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &StorageAccounts_TableServices_Table_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (table *StorageAccountsTableServicesTable) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(table.Spec)
	return table.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (table *StorageAccountsTableServicesTable) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*StorageAccounts_TableServices_Table_STATUS); ok {
		table.Status = *st
		return nil
	}

	// Convert status to required version
	var st StorageAccounts_TableServices_Table_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	table.Status = st
	return nil
}

// AssignProperties_From_StorageAccountsTableServicesTable populates our StorageAccountsTableServicesTable from the provided source StorageAccountsTableServicesTable
func (table *StorageAccountsTableServicesTable) AssignProperties_From_StorageAccountsTableServicesTable(source *storage.StorageAccountsTableServicesTable) error {

	// ObjectMeta
	table.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec StorageAccounts_TableServices_Table_Spec
	err := spec.AssignProperties_From_StorageAccounts_TableServices_Table_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_StorageAccounts_TableServices_Table_Spec() to populate field Spec")
	}
	table.Spec = spec

	// Status
	var status StorageAccounts_TableServices_Table_STATUS
	err = status.AssignProperties_From_StorageAccounts_TableServices_Table_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_StorageAccounts_TableServices_Table_STATUS() to populate field Status")
	}
	table.Status = status

	// Invoke the augmentConversionForStorageAccountsTableServicesTable interface (if implemented) to customize the conversion
	var tableAsAny any = table
	if augmentedTable, ok := tableAsAny.(augmentConversionForStorageAccountsTableServicesTable); ok {
		err := augmentedTable.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_StorageAccountsTableServicesTable populates the provided destination StorageAccountsTableServicesTable from our StorageAccountsTableServicesTable
func (table *StorageAccountsTableServicesTable) AssignProperties_To_StorageAccountsTableServicesTable(destination *storage.StorageAccountsTableServicesTable) error {

	// ObjectMeta
	destination.ObjectMeta = *table.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.StorageAccounts_TableServices_Table_Spec
	err := table.Spec.AssignProperties_To_StorageAccounts_TableServices_Table_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_StorageAccounts_TableServices_Table_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.StorageAccounts_TableServices_Table_STATUS
	err = table.Status.AssignProperties_To_StorageAccounts_TableServices_Table_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_StorageAccounts_TableServices_Table_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForStorageAccountsTableServicesTable interface (if implemented) to customize the conversion
	var tableAsAny any = table
	if augmentedTable, ok := tableAsAny.(augmentConversionForStorageAccountsTableServicesTable); ok {
		err := augmentedTable.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (table *StorageAccountsTableServicesTable) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: table.Spec.OriginalVersion,
		Kind:    "StorageAccountsTableServicesTable",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220901.StorageAccountsTableServicesTable
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2022-09-01/table.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices/default/tables/{tableName}
type StorageAccountsTableServicesTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccountsTableServicesTable `json:"items"`
}

type augmentConversionForStorageAccountsTableServicesTable interface {
	AssignPropertiesFrom(src *storage.StorageAccountsTableServicesTable) error
	AssignPropertiesTo(dst *storage.StorageAccountsTableServicesTable) error
}

// Storage version of v1api20220901.StorageAccounts_TableServices_Table_Spec
type StorageAccounts_TableServices_Table_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string `json:"azureName,omitempty"`
	OriginalVersion string `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a storage.azure.com/StorageAccountsTableService resource
	Owner             *genruntime.KnownResourceReference `group:"storage.azure.com" json:"owner,omitempty" kind:"StorageAccountsTableService"`
	PropertyBag       genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	SignedIdentifiers []TableSignedIdentifier            `json:"signedIdentifiers,omitempty"`
}

var _ genruntime.ConvertibleSpec = &StorageAccounts_TableServices_Table_Spec{}

// ConvertSpecFrom populates our StorageAccounts_TableServices_Table_Spec from the provided source
func (table *StorageAccounts_TableServices_Table_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.StorageAccounts_TableServices_Table_Spec)
	if ok {
		// Populate our instance from source
		return table.AssignProperties_From_StorageAccounts_TableServices_Table_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.StorageAccounts_TableServices_Table_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = table.AssignProperties_From_StorageAccounts_TableServices_Table_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our StorageAccounts_TableServices_Table_Spec
func (table *StorageAccounts_TableServices_Table_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.StorageAccounts_TableServices_Table_Spec)
	if ok {
		// Populate destination from our instance
		return table.AssignProperties_To_StorageAccounts_TableServices_Table_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.StorageAccounts_TableServices_Table_Spec{}
	err := table.AssignProperties_To_StorageAccounts_TableServices_Table_Spec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_StorageAccounts_TableServices_Table_Spec populates our StorageAccounts_TableServices_Table_Spec from the provided source StorageAccounts_TableServices_Table_Spec
func (table *StorageAccounts_TableServices_Table_Spec) AssignProperties_From_StorageAccounts_TableServices_Table_Spec(source *storage.StorageAccounts_TableServices_Table_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	table.AzureName = source.AzureName

	// OriginalVersion
	table.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		table.Owner = &owner
	} else {
		table.Owner = nil
	}

	// SignedIdentifiers
	if source.SignedIdentifiers != nil {
		signedIdentifierList := make([]TableSignedIdentifier, len(source.SignedIdentifiers))
		for signedIdentifierIndex, signedIdentifierItem := range source.SignedIdentifiers {
			// Shadow the loop variable to avoid aliasing
			signedIdentifierItem := signedIdentifierItem
			var signedIdentifier TableSignedIdentifier
			err := signedIdentifier.AssignProperties_From_TableSignedIdentifier(&signedIdentifierItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignProperties_From_TableSignedIdentifier() to populate field SignedIdentifiers")
			}
			signedIdentifierList[signedIdentifierIndex] = signedIdentifier
		}
		table.SignedIdentifiers = signedIdentifierList
	} else {
		table.SignedIdentifiers = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		table.PropertyBag = propertyBag
	} else {
		table.PropertyBag = nil
	}

	// Invoke the augmentConversionForStorageAccounts_TableServices_Table_Spec interface (if implemented) to customize the conversion
	var tableAsAny any = table
	if augmentedTable, ok := tableAsAny.(augmentConversionForStorageAccounts_TableServices_Table_Spec); ok {
		err := augmentedTable.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_StorageAccounts_TableServices_Table_Spec populates the provided destination StorageAccounts_TableServices_Table_Spec from our StorageAccounts_TableServices_Table_Spec
func (table *StorageAccounts_TableServices_Table_Spec) AssignProperties_To_StorageAccounts_TableServices_Table_Spec(destination *storage.StorageAccounts_TableServices_Table_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(table.PropertyBag)

	// AzureName
	destination.AzureName = table.AzureName

	// OriginalVersion
	destination.OriginalVersion = table.OriginalVersion

	// Owner
	if table.Owner != nil {
		owner := table.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// SignedIdentifiers
	if table.SignedIdentifiers != nil {
		signedIdentifierList := make([]storage.TableSignedIdentifier, len(table.SignedIdentifiers))
		for signedIdentifierIndex, signedIdentifierItem := range table.SignedIdentifiers {
			// Shadow the loop variable to avoid aliasing
			signedIdentifierItem := signedIdentifierItem
			var signedIdentifier storage.TableSignedIdentifier
			err := signedIdentifierItem.AssignProperties_To_TableSignedIdentifier(&signedIdentifier)
			if err != nil {
				return errors.Wrap(err, "calling AssignProperties_To_TableSignedIdentifier() to populate field SignedIdentifiers")
			}
			signedIdentifierList[signedIdentifierIndex] = signedIdentifier
		}
		destination.SignedIdentifiers = signedIdentifierList
	} else {
		destination.SignedIdentifiers = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForStorageAccounts_TableServices_Table_Spec interface (if implemented) to customize the conversion
	var tableAsAny any = table
	if augmentedTable, ok := tableAsAny.(augmentConversionForStorageAccounts_TableServices_Table_Spec); ok {
		err := augmentedTable.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20220901.StorageAccounts_TableServices_Table_STATUS
type StorageAccounts_TableServices_Table_STATUS struct {
	Conditions        []conditions.Condition         `json:"conditions,omitempty"`
	Id                *string                        `json:"id,omitempty"`
	Name              *string                        `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
	SignedIdentifiers []TableSignedIdentifier_STATUS `json:"signedIdentifiers,omitempty"`
	TableName         *string                        `json:"tableName,omitempty"`
	Type              *string                        `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &StorageAccounts_TableServices_Table_STATUS{}

// ConvertStatusFrom populates our StorageAccounts_TableServices_Table_STATUS from the provided source
func (table *StorageAccounts_TableServices_Table_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.StorageAccounts_TableServices_Table_STATUS)
	if ok {
		// Populate our instance from source
		return table.AssignProperties_From_StorageAccounts_TableServices_Table_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.StorageAccounts_TableServices_Table_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = table.AssignProperties_From_StorageAccounts_TableServices_Table_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our StorageAccounts_TableServices_Table_STATUS
func (table *StorageAccounts_TableServices_Table_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.StorageAccounts_TableServices_Table_STATUS)
	if ok {
		// Populate destination from our instance
		return table.AssignProperties_To_StorageAccounts_TableServices_Table_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.StorageAccounts_TableServices_Table_STATUS{}
	err := table.AssignProperties_To_StorageAccounts_TableServices_Table_STATUS(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

// AssignProperties_From_StorageAccounts_TableServices_Table_STATUS populates our StorageAccounts_TableServices_Table_STATUS from the provided source StorageAccounts_TableServices_Table_STATUS
func (table *StorageAccounts_TableServices_Table_STATUS) AssignProperties_From_StorageAccounts_TableServices_Table_STATUS(source *storage.StorageAccounts_TableServices_Table_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	table.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	table.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	table.Name = genruntime.ClonePointerToString(source.Name)

	// SignedIdentifiers
	if source.SignedIdentifiers != nil {
		signedIdentifierList := make([]TableSignedIdentifier_STATUS, len(source.SignedIdentifiers))
		for signedIdentifierIndex, signedIdentifierItem := range source.SignedIdentifiers {
			// Shadow the loop variable to avoid aliasing
			signedIdentifierItem := signedIdentifierItem
			var signedIdentifier TableSignedIdentifier_STATUS
			err := signedIdentifier.AssignProperties_From_TableSignedIdentifier_STATUS(&signedIdentifierItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignProperties_From_TableSignedIdentifier_STATUS() to populate field SignedIdentifiers")
			}
			signedIdentifierList[signedIdentifierIndex] = signedIdentifier
		}
		table.SignedIdentifiers = signedIdentifierList
	} else {
		table.SignedIdentifiers = nil
	}

	// TableName
	table.TableName = genruntime.ClonePointerToString(source.TableName)

	// Type
	table.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		table.PropertyBag = propertyBag
	} else {
		table.PropertyBag = nil
	}

	// Invoke the augmentConversionForStorageAccounts_TableServices_Table_STATUS interface (if implemented) to customize the conversion
	var tableAsAny any = table
	if augmentedTable, ok := tableAsAny.(augmentConversionForStorageAccounts_TableServices_Table_STATUS); ok {
		err := augmentedTable.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_StorageAccounts_TableServices_Table_STATUS populates the provided destination StorageAccounts_TableServices_Table_STATUS from our StorageAccounts_TableServices_Table_STATUS
func (table *StorageAccounts_TableServices_Table_STATUS) AssignProperties_To_StorageAccounts_TableServices_Table_STATUS(destination *storage.StorageAccounts_TableServices_Table_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(table.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(table.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(table.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(table.Name)

	// SignedIdentifiers
	if table.SignedIdentifiers != nil {
		signedIdentifierList := make([]storage.TableSignedIdentifier_STATUS, len(table.SignedIdentifiers))
		for signedIdentifierIndex, signedIdentifierItem := range table.SignedIdentifiers {
			// Shadow the loop variable to avoid aliasing
			signedIdentifierItem := signedIdentifierItem
			var signedIdentifier storage.TableSignedIdentifier_STATUS
			err := signedIdentifierItem.AssignProperties_To_TableSignedIdentifier_STATUS(&signedIdentifier)
			if err != nil {
				return errors.Wrap(err, "calling AssignProperties_To_TableSignedIdentifier_STATUS() to populate field SignedIdentifiers")
			}
			signedIdentifierList[signedIdentifierIndex] = signedIdentifier
		}
		destination.SignedIdentifiers = signedIdentifierList
	} else {
		destination.SignedIdentifiers = nil
	}

	// TableName
	destination.TableName = genruntime.ClonePointerToString(table.TableName)

	// Type
	destination.Type = genruntime.ClonePointerToString(table.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForStorageAccounts_TableServices_Table_STATUS interface (if implemented) to customize the conversion
	var tableAsAny any = table
	if augmentedTable, ok := tableAsAny.(augmentConversionForStorageAccounts_TableServices_Table_STATUS); ok {
		err := augmentedTable.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForStorageAccounts_TableServices_Table_Spec interface {
	AssignPropertiesFrom(src *storage.StorageAccounts_TableServices_Table_Spec) error
	AssignPropertiesTo(dst *storage.StorageAccounts_TableServices_Table_Spec) error
}

type augmentConversionForStorageAccounts_TableServices_Table_STATUS interface {
	AssignPropertiesFrom(src *storage.StorageAccounts_TableServices_Table_STATUS) error
	AssignPropertiesTo(dst *storage.StorageAccounts_TableServices_Table_STATUS) error
}

// Storage version of v1api20220901.TableSignedIdentifier
// Object to set Table Access Policy.
type TableSignedIdentifier struct {
	AccessPolicy *TableAccessPolicy     `json:"accessPolicy,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// +kubebuilder:validation:Required
	// Reference: unique-64-character-value of the stored access policy.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// AssignProperties_From_TableSignedIdentifier populates our TableSignedIdentifier from the provided source TableSignedIdentifier
func (identifier *TableSignedIdentifier) AssignProperties_From_TableSignedIdentifier(source *storage.TableSignedIdentifier) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AccessPolicy
	if source.AccessPolicy != nil {
		var accessPolicy TableAccessPolicy
		err := accessPolicy.AssignProperties_From_TableAccessPolicy(source.AccessPolicy)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_TableAccessPolicy() to populate field AccessPolicy")
		}
		identifier.AccessPolicy = &accessPolicy
	} else {
		identifier.AccessPolicy = nil
	}

	// Reference
	if source.Reference != nil {
		reference := source.Reference.Copy()
		identifier.Reference = &reference
	} else {
		identifier.Reference = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		identifier.PropertyBag = propertyBag
	} else {
		identifier.PropertyBag = nil
	}

	// Invoke the augmentConversionForTableSignedIdentifier interface (if implemented) to customize the conversion
	var identifierAsAny any = identifier
	if augmentedIdentifier, ok := identifierAsAny.(augmentConversionForTableSignedIdentifier); ok {
		err := augmentedIdentifier.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_TableSignedIdentifier populates the provided destination TableSignedIdentifier from our TableSignedIdentifier
func (identifier *TableSignedIdentifier) AssignProperties_To_TableSignedIdentifier(destination *storage.TableSignedIdentifier) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(identifier.PropertyBag)

	// AccessPolicy
	if identifier.AccessPolicy != nil {
		var accessPolicy storage.TableAccessPolicy
		err := identifier.AccessPolicy.AssignProperties_To_TableAccessPolicy(&accessPolicy)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_TableAccessPolicy() to populate field AccessPolicy")
		}
		destination.AccessPolicy = &accessPolicy
	} else {
		destination.AccessPolicy = nil
	}

	// Reference
	if identifier.Reference != nil {
		reference := identifier.Reference.Copy()
		destination.Reference = &reference
	} else {
		destination.Reference = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForTableSignedIdentifier interface (if implemented) to customize the conversion
	var identifierAsAny any = identifier
	if augmentedIdentifier, ok := identifierAsAny.(augmentConversionForTableSignedIdentifier); ok {
		err := augmentedIdentifier.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20220901.TableSignedIdentifier_STATUS
// Object to set Table Access Policy.
type TableSignedIdentifier_STATUS struct {
	AccessPolicy *TableAccessPolicy_STATUS `json:"accessPolicy,omitempty"`
	Id           *string                   `json:"id,omitempty"`
	PropertyBag  genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_TableSignedIdentifier_STATUS populates our TableSignedIdentifier_STATUS from the provided source TableSignedIdentifier_STATUS
func (identifier *TableSignedIdentifier_STATUS) AssignProperties_From_TableSignedIdentifier_STATUS(source *storage.TableSignedIdentifier_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AccessPolicy
	if source.AccessPolicy != nil {
		var accessPolicy TableAccessPolicy_STATUS
		err := accessPolicy.AssignProperties_From_TableAccessPolicy_STATUS(source.AccessPolicy)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_TableAccessPolicy_STATUS() to populate field AccessPolicy")
		}
		identifier.AccessPolicy = &accessPolicy
	} else {
		identifier.AccessPolicy = nil
	}

	// Id
	identifier.Id = genruntime.ClonePointerToString(source.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		identifier.PropertyBag = propertyBag
	} else {
		identifier.PropertyBag = nil
	}

	// Invoke the augmentConversionForTableSignedIdentifier_STATUS interface (if implemented) to customize the conversion
	var identifierAsAny any = identifier
	if augmentedIdentifier, ok := identifierAsAny.(augmentConversionForTableSignedIdentifier_STATUS); ok {
		err := augmentedIdentifier.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_TableSignedIdentifier_STATUS populates the provided destination TableSignedIdentifier_STATUS from our TableSignedIdentifier_STATUS
func (identifier *TableSignedIdentifier_STATUS) AssignProperties_To_TableSignedIdentifier_STATUS(destination *storage.TableSignedIdentifier_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(identifier.PropertyBag)

	// AccessPolicy
	if identifier.AccessPolicy != nil {
		var accessPolicy storage.TableAccessPolicy_STATUS
		err := identifier.AccessPolicy.AssignProperties_To_TableAccessPolicy_STATUS(&accessPolicy)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_TableAccessPolicy_STATUS() to populate field AccessPolicy")
		}
		destination.AccessPolicy = &accessPolicy
	} else {
		destination.AccessPolicy = nil
	}

	// Id
	destination.Id = genruntime.ClonePointerToString(identifier.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForTableSignedIdentifier_STATUS interface (if implemented) to customize the conversion
	var identifierAsAny any = identifier
	if augmentedIdentifier, ok := identifierAsAny.(augmentConversionForTableSignedIdentifier_STATUS); ok {
		err := augmentedIdentifier.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForTableSignedIdentifier interface {
	AssignPropertiesFrom(src *storage.TableSignedIdentifier) error
	AssignPropertiesTo(dst *storage.TableSignedIdentifier) error
}

type augmentConversionForTableSignedIdentifier_STATUS interface {
	AssignPropertiesFrom(src *storage.TableSignedIdentifier_STATUS) error
	AssignPropertiesTo(dst *storage.TableSignedIdentifier_STATUS) error
}

// Storage version of v1api20220901.TableAccessPolicy
// Table Access Policy Properties Object.
type TableAccessPolicy struct {
	ExpiryTime  *string                `json:"expiryTime,omitempty"`
	Permission  *string                `json:"permission,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartTime   *string                `json:"startTime,omitempty"`
}

// AssignProperties_From_TableAccessPolicy populates our TableAccessPolicy from the provided source TableAccessPolicy
func (policy *TableAccessPolicy) AssignProperties_From_TableAccessPolicy(source *storage.TableAccessPolicy) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// ExpiryTime
	policy.ExpiryTime = genruntime.ClonePointerToString(source.ExpiryTime)

	// Permission
	policy.Permission = genruntime.ClonePointerToString(source.Permission)

	// StartTime
	policy.StartTime = genruntime.ClonePointerToString(source.StartTime)

	// Update the property bag
	if len(propertyBag) > 0 {
		policy.PropertyBag = propertyBag
	} else {
		policy.PropertyBag = nil
	}

	// Invoke the augmentConversionForTableAccessPolicy interface (if implemented) to customize the conversion
	var policyAsAny any = policy
	if augmentedPolicy, ok := policyAsAny.(augmentConversionForTableAccessPolicy); ok {
		err := augmentedPolicy.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_TableAccessPolicy populates the provided destination TableAccessPolicy from our TableAccessPolicy
func (policy *TableAccessPolicy) AssignProperties_To_TableAccessPolicy(destination *storage.TableAccessPolicy) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(policy.PropertyBag)

	// ExpiryTime
	destination.ExpiryTime = genruntime.ClonePointerToString(policy.ExpiryTime)

	// Permission
	destination.Permission = genruntime.ClonePointerToString(policy.Permission)

	// StartTime
	destination.StartTime = genruntime.ClonePointerToString(policy.StartTime)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForTableAccessPolicy interface (if implemented) to customize the conversion
	var policyAsAny any = policy
	if augmentedPolicy, ok := policyAsAny.(augmentConversionForTableAccessPolicy); ok {
		err := augmentedPolicy.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20220901.TableAccessPolicy_STATUS
// Table Access Policy Properties Object.
type TableAccessPolicy_STATUS struct {
	ExpiryTime  *string                `json:"expiryTime,omitempty"`
	Permission  *string                `json:"permission,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartTime   *string                `json:"startTime,omitempty"`
}

// AssignProperties_From_TableAccessPolicy_STATUS populates our TableAccessPolicy_STATUS from the provided source TableAccessPolicy_STATUS
func (policy *TableAccessPolicy_STATUS) AssignProperties_From_TableAccessPolicy_STATUS(source *storage.TableAccessPolicy_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// ExpiryTime
	policy.ExpiryTime = genruntime.ClonePointerToString(source.ExpiryTime)

	// Permission
	policy.Permission = genruntime.ClonePointerToString(source.Permission)

	// StartTime
	policy.StartTime = genruntime.ClonePointerToString(source.StartTime)

	// Update the property bag
	if len(propertyBag) > 0 {
		policy.PropertyBag = propertyBag
	} else {
		policy.PropertyBag = nil
	}

	// Invoke the augmentConversionForTableAccessPolicy_STATUS interface (if implemented) to customize the conversion
	var policyAsAny any = policy
	if augmentedPolicy, ok := policyAsAny.(augmentConversionForTableAccessPolicy_STATUS); ok {
		err := augmentedPolicy.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_TableAccessPolicy_STATUS populates the provided destination TableAccessPolicy_STATUS from our TableAccessPolicy_STATUS
func (policy *TableAccessPolicy_STATUS) AssignProperties_To_TableAccessPolicy_STATUS(destination *storage.TableAccessPolicy_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(policy.PropertyBag)

	// ExpiryTime
	destination.ExpiryTime = genruntime.ClonePointerToString(policy.ExpiryTime)

	// Permission
	destination.Permission = genruntime.ClonePointerToString(policy.Permission)

	// StartTime
	destination.StartTime = genruntime.ClonePointerToString(policy.StartTime)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForTableAccessPolicy_STATUS interface (if implemented) to customize the conversion
	var policyAsAny any = policy
	if augmentedPolicy, ok := policyAsAny.(augmentConversionForTableAccessPolicy_STATUS); ok {
		err := augmentedPolicy.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForTableAccessPolicy interface {
	AssignPropertiesFrom(src *storage.TableAccessPolicy) error
	AssignPropertiesTo(dst *storage.TableAccessPolicy) error
}

type augmentConversionForTableAccessPolicy_STATUS interface {
	AssignPropertiesFrom(src *storage.TableAccessPolicy_STATUS) error
	AssignPropertiesTo(dst *storage.TableAccessPolicy_STATUS) error
}

func init() {
	SchemeBuilder.Register(&StorageAccountsTableServicesTable{}, &StorageAccountsTableServicesTableList{})
}
