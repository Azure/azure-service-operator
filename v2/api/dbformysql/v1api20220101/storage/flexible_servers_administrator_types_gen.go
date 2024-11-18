// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1api20230630/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/rotisserie/eris"
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
// Storage version of v1api20220101.FlexibleServersAdministrator
// Generator information:
// - Generated from: /mysql/resource-manager/Microsoft.DBforMySQL/AAD/stable/2022-01-01/AzureADAdministrator.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/flexibleServers/{serverName}/administrators/{administratorName}
type FlexibleServersAdministrator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServersAdministrator_Spec   `json:"spec,omitempty"`
	Status            FlexibleServersAdministrator_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &FlexibleServersAdministrator{}

// GetConditions returns the conditions of the resource
func (administrator *FlexibleServersAdministrator) GetConditions() conditions.Conditions {
	return administrator.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (administrator *FlexibleServersAdministrator) SetConditions(conditions conditions.Conditions) {
	administrator.Status.Conditions = conditions
}

var _ conversion.Convertible = &FlexibleServersAdministrator{}

// ConvertFrom populates our FlexibleServersAdministrator from the provided hub FlexibleServersAdministrator
func (administrator *FlexibleServersAdministrator) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.FlexibleServersAdministrator)
	if !ok {
		return fmt.Errorf("expected dbformysql/v1api20230630/storage/FlexibleServersAdministrator but received %T instead", hub)
	}

	return administrator.AssignProperties_From_FlexibleServersAdministrator(source)
}

// ConvertTo populates the provided hub FlexibleServersAdministrator from our FlexibleServersAdministrator
func (administrator *FlexibleServersAdministrator) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.FlexibleServersAdministrator)
	if !ok {
		return fmt.Errorf("expected dbformysql/v1api20230630/storage/FlexibleServersAdministrator but received %T instead", hub)
	}

	return administrator.AssignProperties_To_FlexibleServersAdministrator(destination)
}

var _ configmaps.Exporter = &FlexibleServersAdministrator{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (administrator *FlexibleServersAdministrator) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if administrator.Spec.OperatorSpec == nil {
		return nil
	}
	return administrator.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &FlexibleServersAdministrator{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (administrator *FlexibleServersAdministrator) SecretDestinationExpressions() []*core.DestinationExpression {
	if administrator.Spec.OperatorSpec == nil {
		return nil
	}
	return administrator.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &FlexibleServersAdministrator{}

// AzureName returns the Azure name of the resource (always "ActiveDirectory")
func (administrator *FlexibleServersAdministrator) AzureName() string {
	return "ActiveDirectory"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-01-01"
func (administrator FlexibleServersAdministrator) GetAPIVersion() string {
	return "2022-01-01"
}

// GetResourceScope returns the scope of the resource
func (administrator *FlexibleServersAdministrator) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (administrator *FlexibleServersAdministrator) GetSpec() genruntime.ConvertibleSpec {
	return &administrator.Spec
}

// GetStatus returns the status of this resource
func (administrator *FlexibleServersAdministrator) GetStatus() genruntime.ConvertibleStatus {
	return &administrator.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (administrator *FlexibleServersAdministrator) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMySQL/flexibleServers/administrators"
func (administrator *FlexibleServersAdministrator) GetType() string {
	return "Microsoft.DBforMySQL/flexibleServers/administrators"
}

// NewEmptyStatus returns a new empty (blank) status
func (administrator *FlexibleServersAdministrator) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &FlexibleServersAdministrator_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (administrator *FlexibleServersAdministrator) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(administrator.Spec)
	return administrator.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (administrator *FlexibleServersAdministrator) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*FlexibleServersAdministrator_STATUS); ok {
		administrator.Status = *st
		return nil
	}

	// Convert status to required version
	var st FlexibleServersAdministrator_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	administrator.Status = st
	return nil
}

// AssignProperties_From_FlexibleServersAdministrator populates our FlexibleServersAdministrator from the provided source FlexibleServersAdministrator
func (administrator *FlexibleServersAdministrator) AssignProperties_From_FlexibleServersAdministrator(source *storage.FlexibleServersAdministrator) error {

	// ObjectMeta
	administrator.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec FlexibleServersAdministrator_Spec
	err := spec.AssignProperties_From_FlexibleServersAdministrator_Spec(&source.Spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_FlexibleServersAdministrator_Spec() to populate field Spec")
	}
	administrator.Spec = spec

	// Status
	var status FlexibleServersAdministrator_STATUS
	err = status.AssignProperties_From_FlexibleServersAdministrator_STATUS(&source.Status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_FlexibleServersAdministrator_STATUS() to populate field Status")
	}
	administrator.Status = status

	// Invoke the augmentConversionForFlexibleServersAdministrator interface (if implemented) to customize the conversion
	var administratorAsAny any = administrator
	if augmentedAdministrator, ok := administratorAsAny.(augmentConversionForFlexibleServersAdministrator); ok {
		err := augmentedAdministrator.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FlexibleServersAdministrator populates the provided destination FlexibleServersAdministrator from our FlexibleServersAdministrator
func (administrator *FlexibleServersAdministrator) AssignProperties_To_FlexibleServersAdministrator(destination *storage.FlexibleServersAdministrator) error {

	// ObjectMeta
	destination.ObjectMeta = *administrator.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.FlexibleServersAdministrator_Spec
	err := administrator.Spec.AssignProperties_To_FlexibleServersAdministrator_Spec(&spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_FlexibleServersAdministrator_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.FlexibleServersAdministrator_STATUS
	err = administrator.Status.AssignProperties_To_FlexibleServersAdministrator_STATUS(&status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_FlexibleServersAdministrator_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForFlexibleServersAdministrator interface (if implemented) to customize the conversion
	var administratorAsAny any = administrator
	if augmentedAdministrator, ok := administratorAsAny.(augmentConversionForFlexibleServersAdministrator); ok {
		err := augmentedAdministrator.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (administrator *FlexibleServersAdministrator) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: administrator.Spec.OriginalVersion,
		Kind:    "FlexibleServersAdministrator",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220101.FlexibleServersAdministrator
// Generator information:
// - Generated from: /mysql/resource-manager/Microsoft.DBforMySQL/AAD/stable/2022-01-01/AzureADAdministrator.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/flexibleServers/{serverName}/administrators/{administratorName}
type FlexibleServersAdministratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServersAdministrator `json:"items"`
}

// Storage version of v1api20220101.APIVersion
// +kubebuilder:validation:Enum={"2022-01-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2022-01-01")

type augmentConversionForFlexibleServersAdministrator interface {
	AssignPropertiesFrom(src *storage.FlexibleServersAdministrator) error
	AssignPropertiesTo(dst *storage.FlexibleServersAdministrator) error
}

// Storage version of v1api20220101.FlexibleServersAdministrator_Spec
type FlexibleServersAdministrator_Spec struct {
	AdministratorType *string `json:"administratorType,omitempty"`

	// IdentityResourceReference: The resource id of the identity used for AAD Authentication.
	IdentityResourceReference *genruntime.ResourceReference             `armReference:"IdentityResourceId" json:"identityResourceReference,omitempty"`
	Login                     *string                                   `json:"login,omitempty"`
	OperatorSpec              *FlexibleServersAdministratorOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion           string                                    `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a dbformysql.azure.com/FlexibleServer resource
	Owner              *genruntime.KnownResourceReference `group:"dbformysql.azure.com" json:"owner,omitempty" kind:"FlexibleServer"`
	PropertyBag        genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Sid                *string                            `json:"sid,omitempty" optionalConfigMapPair:"Sid"`
	SidFromConfig      *genruntime.ConfigMapReference     `json:"sidFromConfig,omitempty" optionalConfigMapPair:"Sid"`
	TenantId           *string                            `json:"tenantId,omitempty" optionalConfigMapPair:"TenantId"`
	TenantIdFromConfig *genruntime.ConfigMapReference     `json:"tenantIdFromConfig,omitempty" optionalConfigMapPair:"TenantId"`
}

var _ genruntime.ConvertibleSpec = &FlexibleServersAdministrator_Spec{}

// ConvertSpecFrom populates our FlexibleServersAdministrator_Spec from the provided source
func (administrator *FlexibleServersAdministrator_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.FlexibleServersAdministrator_Spec)
	if ok {
		// Populate our instance from source
		return administrator.AssignProperties_From_FlexibleServersAdministrator_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.FlexibleServersAdministrator_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = administrator.AssignProperties_From_FlexibleServersAdministrator_Spec(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our FlexibleServersAdministrator_Spec
func (administrator *FlexibleServersAdministrator_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.FlexibleServersAdministrator_Spec)
	if ok {
		// Populate destination from our instance
		return administrator.AssignProperties_To_FlexibleServersAdministrator_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.FlexibleServersAdministrator_Spec{}
	err := administrator.AssignProperties_To_FlexibleServersAdministrator_Spec(dst)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_FlexibleServersAdministrator_Spec populates our FlexibleServersAdministrator_Spec from the provided source FlexibleServersAdministrator_Spec
func (administrator *FlexibleServersAdministrator_Spec) AssignProperties_From_FlexibleServersAdministrator_Spec(source *storage.FlexibleServersAdministrator_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AdministratorType
	administrator.AdministratorType = genruntime.ClonePointerToString(source.AdministratorType)

	// IdentityResourceReference
	if source.IdentityResourceReference != nil {
		identityResourceReference := source.IdentityResourceReference.Copy()
		administrator.IdentityResourceReference = &identityResourceReference
	} else {
		administrator.IdentityResourceReference = nil
	}

	// Login
	administrator.Login = genruntime.ClonePointerToString(source.Login)

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec FlexibleServersAdministratorOperatorSpec
		err := operatorSpec.AssignProperties_From_FlexibleServersAdministratorOperatorSpec(source.OperatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_FlexibleServersAdministratorOperatorSpec() to populate field OperatorSpec")
		}
		administrator.OperatorSpec = &operatorSpec
	} else {
		administrator.OperatorSpec = nil
	}

	// OriginalVersion
	administrator.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		administrator.Owner = &owner
	} else {
		administrator.Owner = nil
	}

	// Sid
	administrator.Sid = genruntime.ClonePointerToString(source.Sid)

	// SidFromConfig
	if source.SidFromConfig != nil {
		sidFromConfig := source.SidFromConfig.Copy()
		administrator.SidFromConfig = &sidFromConfig
	} else {
		administrator.SidFromConfig = nil
	}

	// TenantId
	administrator.TenantId = genruntime.ClonePointerToString(source.TenantId)

	// TenantIdFromConfig
	if source.TenantIdFromConfig != nil {
		tenantIdFromConfig := source.TenantIdFromConfig.Copy()
		administrator.TenantIdFromConfig = &tenantIdFromConfig
	} else {
		administrator.TenantIdFromConfig = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		administrator.PropertyBag = propertyBag
	} else {
		administrator.PropertyBag = nil
	}

	// Invoke the augmentConversionForFlexibleServersAdministrator_Spec interface (if implemented) to customize the conversion
	var administratorAsAny any = administrator
	if augmentedAdministrator, ok := administratorAsAny.(augmentConversionForFlexibleServersAdministrator_Spec); ok {
		err := augmentedAdministrator.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FlexibleServersAdministrator_Spec populates the provided destination FlexibleServersAdministrator_Spec from our FlexibleServersAdministrator_Spec
func (administrator *FlexibleServersAdministrator_Spec) AssignProperties_To_FlexibleServersAdministrator_Spec(destination *storage.FlexibleServersAdministrator_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(administrator.PropertyBag)

	// AdministratorType
	destination.AdministratorType = genruntime.ClonePointerToString(administrator.AdministratorType)

	// IdentityResourceReference
	if administrator.IdentityResourceReference != nil {
		identityResourceReference := administrator.IdentityResourceReference.Copy()
		destination.IdentityResourceReference = &identityResourceReference
	} else {
		destination.IdentityResourceReference = nil
	}

	// Login
	destination.Login = genruntime.ClonePointerToString(administrator.Login)

	// OperatorSpec
	if administrator.OperatorSpec != nil {
		var operatorSpec storage.FlexibleServersAdministratorOperatorSpec
		err := administrator.OperatorSpec.AssignProperties_To_FlexibleServersAdministratorOperatorSpec(&operatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_FlexibleServersAdministratorOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

	// OriginalVersion
	destination.OriginalVersion = administrator.OriginalVersion

	// Owner
	if administrator.Owner != nil {
		owner := administrator.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Sid
	destination.Sid = genruntime.ClonePointerToString(administrator.Sid)

	// SidFromConfig
	if administrator.SidFromConfig != nil {
		sidFromConfig := administrator.SidFromConfig.Copy()
		destination.SidFromConfig = &sidFromConfig
	} else {
		destination.SidFromConfig = nil
	}

	// TenantId
	destination.TenantId = genruntime.ClonePointerToString(administrator.TenantId)

	// TenantIdFromConfig
	if administrator.TenantIdFromConfig != nil {
		tenantIdFromConfig := administrator.TenantIdFromConfig.Copy()
		destination.TenantIdFromConfig = &tenantIdFromConfig
	} else {
		destination.TenantIdFromConfig = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForFlexibleServersAdministrator_Spec interface (if implemented) to customize the conversion
	var administratorAsAny any = administrator
	if augmentedAdministrator, ok := administratorAsAny.(augmentConversionForFlexibleServersAdministrator_Spec); ok {
		err := augmentedAdministrator.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20220101.FlexibleServersAdministrator_STATUS
type FlexibleServersAdministrator_STATUS struct {
	AdministratorType  *string                `json:"administratorType,omitempty"`
	Conditions         []conditions.Condition `json:"conditions,omitempty"`
	Id                 *string                `json:"id,omitempty"`
	IdentityResourceId *string                `json:"identityResourceId,omitempty"`
	Login              *string                `json:"login,omitempty"`
	Name               *string                `json:"name,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Sid                *string                `json:"sid,omitempty"`
	SystemData         *SystemData_STATUS     `json:"systemData,omitempty"`
	TenantId           *string                `json:"tenantId,omitempty"`
	Type               *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &FlexibleServersAdministrator_STATUS{}

// ConvertStatusFrom populates our FlexibleServersAdministrator_STATUS from the provided source
func (administrator *FlexibleServersAdministrator_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.FlexibleServersAdministrator_STATUS)
	if ok {
		// Populate our instance from source
		return administrator.AssignProperties_From_FlexibleServersAdministrator_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.FlexibleServersAdministrator_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = administrator.AssignProperties_From_FlexibleServersAdministrator_STATUS(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our FlexibleServersAdministrator_STATUS
func (administrator *FlexibleServersAdministrator_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.FlexibleServersAdministrator_STATUS)
	if ok {
		// Populate destination from our instance
		return administrator.AssignProperties_To_FlexibleServersAdministrator_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.FlexibleServersAdministrator_STATUS{}
	err := administrator.AssignProperties_To_FlexibleServersAdministrator_STATUS(dst)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

// AssignProperties_From_FlexibleServersAdministrator_STATUS populates our FlexibleServersAdministrator_STATUS from the provided source FlexibleServersAdministrator_STATUS
func (administrator *FlexibleServersAdministrator_STATUS) AssignProperties_From_FlexibleServersAdministrator_STATUS(source *storage.FlexibleServersAdministrator_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AdministratorType
	administrator.AdministratorType = genruntime.ClonePointerToString(source.AdministratorType)

	// Conditions
	administrator.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	administrator.Id = genruntime.ClonePointerToString(source.Id)

	// IdentityResourceId
	administrator.IdentityResourceId = genruntime.ClonePointerToString(source.IdentityResourceId)

	// Login
	administrator.Login = genruntime.ClonePointerToString(source.Login)

	// Name
	administrator.Name = genruntime.ClonePointerToString(source.Name)

	// Sid
	administrator.Sid = genruntime.ClonePointerToString(source.Sid)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
		}
		administrator.SystemData = &systemDatum
	} else {
		administrator.SystemData = nil
	}

	// TenantId
	administrator.TenantId = genruntime.ClonePointerToString(source.TenantId)

	// Type
	administrator.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		administrator.PropertyBag = propertyBag
	} else {
		administrator.PropertyBag = nil
	}

	// Invoke the augmentConversionForFlexibleServersAdministrator_STATUS interface (if implemented) to customize the conversion
	var administratorAsAny any = administrator
	if augmentedAdministrator, ok := administratorAsAny.(augmentConversionForFlexibleServersAdministrator_STATUS); ok {
		err := augmentedAdministrator.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FlexibleServersAdministrator_STATUS populates the provided destination FlexibleServersAdministrator_STATUS from our FlexibleServersAdministrator_STATUS
func (administrator *FlexibleServersAdministrator_STATUS) AssignProperties_To_FlexibleServersAdministrator_STATUS(destination *storage.FlexibleServersAdministrator_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(administrator.PropertyBag)

	// AdministratorType
	destination.AdministratorType = genruntime.ClonePointerToString(administrator.AdministratorType)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(administrator.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(administrator.Id)

	// IdentityResourceId
	destination.IdentityResourceId = genruntime.ClonePointerToString(administrator.IdentityResourceId)

	// Login
	destination.Login = genruntime.ClonePointerToString(administrator.Login)

	// Name
	destination.Name = genruntime.ClonePointerToString(administrator.Name)

	// Sid
	destination.Sid = genruntime.ClonePointerToString(administrator.Sid)

	// SystemData
	if administrator.SystemData != nil {
		var systemDatum storage.SystemData_STATUS
		err := administrator.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// TenantId
	destination.TenantId = genruntime.ClonePointerToString(administrator.TenantId)

	// Type
	destination.Type = genruntime.ClonePointerToString(administrator.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForFlexibleServersAdministrator_STATUS interface (if implemented) to customize the conversion
	var administratorAsAny any = administrator
	if augmentedAdministrator, ok := administratorAsAny.(augmentConversionForFlexibleServersAdministrator_STATUS); ok {
		err := augmentedAdministrator.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForFlexibleServersAdministrator_Spec interface {
	AssignPropertiesFrom(src *storage.FlexibleServersAdministrator_Spec) error
	AssignPropertiesTo(dst *storage.FlexibleServersAdministrator_Spec) error
}

type augmentConversionForFlexibleServersAdministrator_STATUS interface {
	AssignPropertiesFrom(src *storage.FlexibleServersAdministrator_STATUS) error
	AssignPropertiesTo(dst *storage.FlexibleServersAdministrator_STATUS) error
}

// Storage version of v1api20220101.FlexibleServersAdministratorOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type FlexibleServersAdministratorOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_FlexibleServersAdministratorOperatorSpec populates our FlexibleServersAdministratorOperatorSpec from the provided source FlexibleServersAdministratorOperatorSpec
func (operator *FlexibleServersAdministratorOperatorSpec) AssignProperties_From_FlexibleServersAdministratorOperatorSpec(source *storage.FlexibleServersAdministratorOperatorSpec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// ConfigMapExpressions
	if source.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(source.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range source.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		operator.ConfigMapExpressions = configMapExpressionList
	} else {
		operator.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if source.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(source.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range source.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		operator.SecretExpressions = secretExpressionList
	} else {
		operator.SecretExpressions = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		operator.PropertyBag = propertyBag
	} else {
		operator.PropertyBag = nil
	}

	// Invoke the augmentConversionForFlexibleServersAdministratorOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForFlexibleServersAdministratorOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FlexibleServersAdministratorOperatorSpec populates the provided destination FlexibleServersAdministratorOperatorSpec from our FlexibleServersAdministratorOperatorSpec
func (operator *FlexibleServersAdministratorOperatorSpec) AssignProperties_To_FlexibleServersAdministratorOperatorSpec(destination *storage.FlexibleServersAdministratorOperatorSpec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(operator.PropertyBag)

	// ConfigMapExpressions
	if operator.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(operator.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range operator.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		destination.ConfigMapExpressions = configMapExpressionList
	} else {
		destination.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if operator.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(operator.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range operator.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		destination.SecretExpressions = secretExpressionList
	} else {
		destination.SecretExpressions = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForFlexibleServersAdministratorOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForFlexibleServersAdministratorOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20220101.SystemData_STATUS
// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS struct {
	CreatedAt          *string                `json:"createdAt,omitempty"`
	CreatedBy          *string                `json:"createdBy,omitempty"`
	CreatedByType      *string                `json:"createdByType,omitempty"`
	LastModifiedAt     *string                `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *string                `json:"lastModifiedByType,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_SystemData_STATUS populates our SystemData_STATUS from the provided source SystemData_STATUS
func (data *SystemData_STATUS) AssignProperties_From_SystemData_STATUS(source *storage.SystemData_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// CreatedAt
	data.CreatedAt = genruntime.ClonePointerToString(source.CreatedAt)

	// CreatedBy
	data.CreatedBy = genruntime.ClonePointerToString(source.CreatedBy)

	// CreatedByType
	data.CreatedByType = genruntime.ClonePointerToString(source.CreatedByType)

	// LastModifiedAt
	data.LastModifiedAt = genruntime.ClonePointerToString(source.LastModifiedAt)

	// LastModifiedBy
	data.LastModifiedBy = genruntime.ClonePointerToString(source.LastModifiedBy)

	// LastModifiedByType
	data.LastModifiedByType = genruntime.ClonePointerToString(source.LastModifiedByType)

	// Update the property bag
	if len(propertyBag) > 0 {
		data.PropertyBag = propertyBag
	} else {
		data.PropertyBag = nil
	}

	// Invoke the augmentConversionForSystemData_STATUS interface (if implemented) to customize the conversion
	var dataAsAny any = data
	if augmentedData, ok := dataAsAny.(augmentConversionForSystemData_STATUS); ok {
		err := augmentedData.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SystemData_STATUS populates the provided destination SystemData_STATUS from our SystemData_STATUS
func (data *SystemData_STATUS) AssignProperties_To_SystemData_STATUS(destination *storage.SystemData_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(data.PropertyBag)

	// CreatedAt
	destination.CreatedAt = genruntime.ClonePointerToString(data.CreatedAt)

	// CreatedBy
	destination.CreatedBy = genruntime.ClonePointerToString(data.CreatedBy)

	// CreatedByType
	destination.CreatedByType = genruntime.ClonePointerToString(data.CreatedByType)

	// LastModifiedAt
	destination.LastModifiedAt = genruntime.ClonePointerToString(data.LastModifiedAt)

	// LastModifiedBy
	destination.LastModifiedBy = genruntime.ClonePointerToString(data.LastModifiedBy)

	// LastModifiedByType
	destination.LastModifiedByType = genruntime.ClonePointerToString(data.LastModifiedByType)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForSystemData_STATUS interface (if implemented) to customize the conversion
	var dataAsAny any = data
	if augmentedData, ok := dataAsAny.(augmentConversionForSystemData_STATUS); ok {
		err := augmentedData.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForFlexibleServersAdministratorOperatorSpec interface {
	AssignPropertiesFrom(src *storage.FlexibleServersAdministratorOperatorSpec) error
	AssignPropertiesTo(dst *storage.FlexibleServersAdministratorOperatorSpec) error
}

type augmentConversionForSystemData_STATUS interface {
	AssignPropertiesFrom(src *storage.SystemData_STATUS) error
	AssignPropertiesTo(dst *storage.SystemData_STATUS) error
}

func init() {
	SchemeBuilder.Register(&FlexibleServersAdministrator{}, &FlexibleServersAdministratorList{})
}
