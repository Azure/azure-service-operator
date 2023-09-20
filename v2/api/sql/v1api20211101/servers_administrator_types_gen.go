// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import (
	"fmt"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/ServerAzureADAdministrators.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/administrators/{administratorName}
type ServersAdministrator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Servers_Administrator_Spec   `json:"spec,omitempty"`
	Status            Servers_Administrator_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ServersAdministrator{}

// GetConditions returns the conditions of the resource
func (administrator *ServersAdministrator) GetConditions() conditions.Conditions {
	return administrator.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (administrator *ServersAdministrator) SetConditions(conditions conditions.Conditions) {
	administrator.Status.Conditions = conditions
}

var _ conversion.Convertible = &ServersAdministrator{}

// ConvertFrom populates our ServersAdministrator from the provided hub ServersAdministrator
func (administrator *ServersAdministrator) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20211101s.ServersAdministrator)
	if !ok {
		return fmt.Errorf("expected sql/v1api20211101storage/ServersAdministrator but received %T instead", hub)
	}

	return administrator.AssignProperties_From_ServersAdministrator(source)
}

// ConvertTo populates the provided hub ServersAdministrator from our ServersAdministrator
func (administrator *ServersAdministrator) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20211101s.ServersAdministrator)
	if !ok {
		return fmt.Errorf("expected sql/v1api20211101storage/ServersAdministrator but received %T instead", hub)
	}

	return administrator.AssignProperties_To_ServersAdministrator(destination)
}

// +kubebuilder:webhook:path=/mutate-sql-azure-com-v1api20211101-serversadministrator,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversadministrators,verbs=create;update,versions=v1api20211101,name=default.v1api20211101.serversadministrators.sql.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &ServersAdministrator{}

// Default applies defaults to the ServersAdministrator resource
func (administrator *ServersAdministrator) Default() {
	administrator.defaultImpl()
	var temp any = administrator
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the ServersAdministrator resource
func (administrator *ServersAdministrator) defaultImpl() {}

var _ genruntime.ImportableResource = &ServersAdministrator{}

// InitializeSpec initializes the spec for this resource from the given status
func (administrator *ServersAdministrator) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*Servers_Administrator_STATUS); ok {
		return administrator.Spec.Initialize_From_Servers_Administrator_STATUS(s)
	}

	return fmt.Errorf("expected Status of type Servers_Administrator_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &ServersAdministrator{}

// AzureName returns the Azure name of the resource (always "ActiveDirectory")
func (administrator *ServersAdministrator) AzureName() string {
	return "ActiveDirectory"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (administrator ServersAdministrator) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (administrator *ServersAdministrator) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (administrator *ServersAdministrator) GetSpec() genruntime.ConvertibleSpec {
	return &administrator.Spec
}

// GetStatus returns the status of this resource
func (administrator *ServersAdministrator) GetStatus() genruntime.ConvertibleStatus {
	return &administrator.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/administrators"
func (administrator *ServersAdministrator) GetType() string {
	return "Microsoft.Sql/servers/administrators"
}

// NewEmptyStatus returns a new empty (blank) status
func (administrator *ServersAdministrator) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Servers_Administrator_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (administrator *ServersAdministrator) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(administrator.Spec)
	return administrator.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (administrator *ServersAdministrator) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Servers_Administrator_STATUS); ok {
		administrator.Status = *st
		return nil
	}

	// Convert status to required version
	var st Servers_Administrator_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	administrator.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-sql-azure-com-v1api20211101-serversadministrator,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversadministrators,verbs=create;update,versions=v1api20211101,name=validate.v1api20211101.serversadministrators.sql.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &ServersAdministrator{}

// ValidateCreate validates the creation of the resource
func (administrator *ServersAdministrator) ValidateCreate() (admission.Warnings, error) {
	validations := administrator.createValidations()
	var temp any = administrator
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (administrator *ServersAdministrator) ValidateDelete() (admission.Warnings, error) {
	validations := administrator.deleteValidations()
	var temp any = administrator
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (administrator *ServersAdministrator) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := administrator.updateValidations()
	var temp any = administrator
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (administrator *ServersAdministrator) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){administrator.validateResourceReferences, administrator.validateOwnerReference, administrator.validateOptionalConfigMapReferences}
}

// deleteValidations validates the deletion of the resource
func (administrator *ServersAdministrator) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (administrator *ServersAdministrator) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return administrator.validateResourceReferences()
		},
		administrator.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return administrator.validateOwnerReference()
		},
		func(old runtime.Object) (admission.Warnings, error) {
			return administrator.validateOptionalConfigMapReferences()
		},
	}
}

// validateOptionalConfigMapReferences validates all optional configmap reference pairs to ensure that at most 1 is set
func (administrator *ServersAdministrator) validateOptionalConfigMapReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindOptionalConfigMapReferences(&administrator.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateOptionalConfigMapReferences(refs)
}

// validateOwnerReference validates the owner field
func (administrator *ServersAdministrator) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(administrator)
}

// validateResourceReferences validates all resource references
func (administrator *ServersAdministrator) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&administrator.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (administrator *ServersAdministrator) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*ServersAdministrator)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, administrator)
}

// AssignProperties_From_ServersAdministrator populates our ServersAdministrator from the provided source ServersAdministrator
func (administrator *ServersAdministrator) AssignProperties_From_ServersAdministrator(source *v20211101s.ServersAdministrator) error {

	// ObjectMeta
	administrator.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Servers_Administrator_Spec
	err := spec.AssignProperties_From_Servers_Administrator_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_Administrator_Spec() to populate field Spec")
	}
	administrator.Spec = spec

	// Status
	var status Servers_Administrator_STATUS
	err = status.AssignProperties_From_Servers_Administrator_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_Administrator_STATUS() to populate field Status")
	}
	administrator.Status = status

	// No error
	return nil
}

// AssignProperties_To_ServersAdministrator populates the provided destination ServersAdministrator from our ServersAdministrator
func (administrator *ServersAdministrator) AssignProperties_To_ServersAdministrator(destination *v20211101s.ServersAdministrator) error {

	// ObjectMeta
	destination.ObjectMeta = *administrator.ObjectMeta.DeepCopy()

	// Spec
	var spec v20211101s.Servers_Administrator_Spec
	err := administrator.Spec.AssignProperties_To_Servers_Administrator_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_Administrator_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20211101s.Servers_Administrator_STATUS
	err = administrator.Status.AssignProperties_To_Servers_Administrator_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_Administrator_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (administrator *ServersAdministrator) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: administrator.Spec.OriginalVersion(),
		Kind:    "ServersAdministrator",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/ServerAzureADAdministrators.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/administrators/{administratorName}
type ServersAdministratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServersAdministrator `json:"items"`
}

type Servers_Administrator_Spec struct {
	// +kubebuilder:validation:Required
	// AdministratorType: Type of the sever administrator.
	AdministratorType *AdministratorProperties_AdministratorType `json:"administratorType,omitempty"`

	// +kubebuilder:validation:Required
	// Login: Login name of the server administrator.
	Login *string `json:"login,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a sql.azure.com/Server resource
	Owner *genruntime.KnownResourceReference `group:"sql.azure.com" json:"owner,omitempty" kind:"Server"`

	// +kubebuilder:validation:Pattern="^[0-9a-fA-F]{8}(-[0-9a-fA-F]{4}){3}-[0-9a-fA-F]{12}$"
	// Sid: SID (object ID) of the server administrator.
	Sid *string `json:"sid,omitempty" optionalConfigMapPair:"Sid"`

	// SidFromConfig: SID (object ID) of the server administrator.
	SidFromConfig *genruntime.ConfigMapReference `json:"sidFromConfig,omitempty" optionalConfigMapPair:"Sid"`

	// +kubebuilder:validation:Pattern="^[0-9a-fA-F]{8}(-[0-9a-fA-F]{4}){3}-[0-9a-fA-F]{12}$"
	// TenantId: Tenant ID of the administrator.
	TenantId *string `json:"tenantId,omitempty" optionalConfigMapPair:"TenantId"`

	// TenantIdFromConfig: Tenant ID of the administrator.
	TenantIdFromConfig *genruntime.ConfigMapReference `json:"tenantIdFromConfig,omitempty" optionalConfigMapPair:"TenantId"`
}

var _ genruntime.ARMTransformer = &Servers_Administrator_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (administrator *Servers_Administrator_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if administrator == nil {
		return nil, nil
	}
	result := &Servers_Administrator_Spec_ARM{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if administrator.AdministratorType != nil ||
		administrator.Login != nil ||
		administrator.Sid != nil ||
		administrator.SidFromConfig != nil ||
		administrator.TenantId != nil ||
		administrator.TenantIdFromConfig != nil {
		result.Properties = &AdministratorProperties_ARM{}
	}
	if administrator.AdministratorType != nil {
		administratorType := *administrator.AdministratorType
		result.Properties.AdministratorType = &administratorType
	}
	if administrator.Login != nil {
		login := *administrator.Login
		result.Properties.Login = &login
	}
	if administrator.Sid != nil {
		sid := *administrator.Sid
		result.Properties.Sid = &sid
	}
	if administrator.SidFromConfig != nil {
		sidValue, err := resolved.ResolvedConfigMaps.Lookup(*administrator.SidFromConfig)
		if err != nil {
			return nil, errors.Wrap(err, "looking up configmap for property Sid")
		}
		sid := sidValue
		result.Properties.Sid = &sid
	}
	if administrator.TenantId != nil {
		tenantId := *administrator.TenantId
		result.Properties.TenantId = &tenantId
	}
	if administrator.TenantIdFromConfig != nil {
		tenantIdValue, err := resolved.ResolvedConfigMaps.Lookup(*administrator.TenantIdFromConfig)
		if err != nil {
			return nil, errors.Wrap(err, "looking up configmap for property TenantId")
		}
		tenantId := tenantIdValue
		result.Properties.TenantId = &tenantId
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (administrator *Servers_Administrator_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_Administrator_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (administrator *Servers_Administrator_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_Administrator_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_Administrator_Spec_ARM, got %T", armInput)
	}

	// Set property "AdministratorType":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.AdministratorType != nil {
			administratorType := *typedInput.Properties.AdministratorType
			administrator.AdministratorType = &administratorType
		}
	}

	// Set property "Login":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Login != nil {
			login := *typedInput.Properties.Login
			administrator.Login = &login
		}
	}

	// Set property "Owner":
	administrator.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "Sid":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Sid != nil {
			sid := *typedInput.Properties.Sid
			administrator.Sid = &sid
		}
	}

	// no assignment for property "SidFromConfig"

	// Set property "TenantId":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.TenantId != nil {
			tenantId := *typedInput.Properties.TenantId
			administrator.TenantId = &tenantId
		}
	}

	// no assignment for property "TenantIdFromConfig"

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Servers_Administrator_Spec{}

// ConvertSpecFrom populates our Servers_Administrator_Spec from the provided source
func (administrator *Servers_Administrator_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20211101s.Servers_Administrator_Spec)
	if ok {
		// Populate our instance from source
		return administrator.AssignProperties_From_Servers_Administrator_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.Servers_Administrator_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = administrator.AssignProperties_From_Servers_Administrator_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Servers_Administrator_Spec
func (administrator *Servers_Administrator_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20211101s.Servers_Administrator_Spec)
	if ok {
		// Populate destination from our instance
		return administrator.AssignProperties_To_Servers_Administrator_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.Servers_Administrator_Spec{}
	err := administrator.AssignProperties_To_Servers_Administrator_Spec(dst)
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

// AssignProperties_From_Servers_Administrator_Spec populates our Servers_Administrator_Spec from the provided source Servers_Administrator_Spec
func (administrator *Servers_Administrator_Spec) AssignProperties_From_Servers_Administrator_Spec(source *v20211101s.Servers_Administrator_Spec) error {

	// AdministratorType
	if source.AdministratorType != nil {
		administratorType := AdministratorProperties_AdministratorType(*source.AdministratorType)
		administrator.AdministratorType = &administratorType
	} else {
		administrator.AdministratorType = nil
	}

	// Login
	administrator.Login = genruntime.ClonePointerToString(source.Login)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		administrator.Owner = &owner
	} else {
		administrator.Owner = nil
	}

	// Sid
	if source.Sid != nil {
		sid := *source.Sid
		administrator.Sid = &sid
	} else {
		administrator.Sid = nil
	}

	// SidFromConfig
	if source.SidFromConfig != nil {
		sidFromConfig := source.SidFromConfig.Copy()
		administrator.SidFromConfig = &sidFromConfig
	} else {
		administrator.SidFromConfig = nil
	}

	// TenantId
	if source.TenantId != nil {
		tenantId := *source.TenantId
		administrator.TenantId = &tenantId
	} else {
		administrator.TenantId = nil
	}

	// TenantIdFromConfig
	if source.TenantIdFromConfig != nil {
		tenantIdFromConfig := source.TenantIdFromConfig.Copy()
		administrator.TenantIdFromConfig = &tenantIdFromConfig
	} else {
		administrator.TenantIdFromConfig = nil
	}

	// No error
	return nil
}

// AssignProperties_To_Servers_Administrator_Spec populates the provided destination Servers_Administrator_Spec from our Servers_Administrator_Spec
func (administrator *Servers_Administrator_Spec) AssignProperties_To_Servers_Administrator_Spec(destination *v20211101s.Servers_Administrator_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AdministratorType
	if administrator.AdministratorType != nil {
		administratorType := string(*administrator.AdministratorType)
		destination.AdministratorType = &administratorType
	} else {
		destination.AdministratorType = nil
	}

	// Login
	destination.Login = genruntime.ClonePointerToString(administrator.Login)

	// OriginalVersion
	destination.OriginalVersion = administrator.OriginalVersion()

	// Owner
	if administrator.Owner != nil {
		owner := administrator.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Sid
	if administrator.Sid != nil {
		sid := *administrator.Sid
		destination.Sid = &sid
	} else {
		destination.Sid = nil
	}

	// SidFromConfig
	if administrator.SidFromConfig != nil {
		sidFromConfig := administrator.SidFromConfig.Copy()
		destination.SidFromConfig = &sidFromConfig
	} else {
		destination.SidFromConfig = nil
	}

	// TenantId
	if administrator.TenantId != nil {
		tenantId := *administrator.TenantId
		destination.TenantId = &tenantId
	} else {
		destination.TenantId = nil
	}

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

	// No error
	return nil
}

// Initialize_From_Servers_Administrator_STATUS populates our Servers_Administrator_Spec from the provided source Servers_Administrator_STATUS
func (administrator *Servers_Administrator_Spec) Initialize_From_Servers_Administrator_STATUS(source *Servers_Administrator_STATUS) error {

	// AdministratorType
	if source.AdministratorType != nil {
		administratorType := AdministratorProperties_AdministratorType(*source.AdministratorType)
		administrator.AdministratorType = &administratorType
	} else {
		administrator.AdministratorType = nil
	}

	// Login
	administrator.Login = genruntime.ClonePointerToString(source.Login)

	// Sid
	if source.Sid != nil {
		sid := *source.Sid
		administrator.Sid = &sid
	} else {
		administrator.Sid = nil
	}

	// TenantId
	if source.TenantId != nil {
		tenantId := *source.TenantId
		administrator.TenantId = &tenantId
	} else {
		administrator.TenantId = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (administrator *Servers_Administrator_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

type Servers_Administrator_STATUS struct {
	// AdministratorType: Type of the sever administrator.
	AdministratorType *AdministratorProperties_AdministratorType_STATUS `json:"administratorType,omitempty"`

	// AzureADOnlyAuthentication: Azure Active Directory only Authentication enabled.
	AzureADOnlyAuthentication *bool `json:"azureADOnlyAuthentication,omitempty"`

	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Login: Login name of the server administrator.
	Login *string `json:"login,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Sid: SID (object ID) of the server administrator.
	Sid *string `json:"sid,omitempty"`

	// TenantId: Tenant ID of the administrator.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Servers_Administrator_STATUS{}

// ConvertStatusFrom populates our Servers_Administrator_STATUS from the provided source
func (administrator *Servers_Administrator_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20211101s.Servers_Administrator_STATUS)
	if ok {
		// Populate our instance from source
		return administrator.AssignProperties_From_Servers_Administrator_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.Servers_Administrator_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = administrator.AssignProperties_From_Servers_Administrator_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Servers_Administrator_STATUS
func (administrator *Servers_Administrator_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20211101s.Servers_Administrator_STATUS)
	if ok {
		// Populate destination from our instance
		return administrator.AssignProperties_To_Servers_Administrator_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.Servers_Administrator_STATUS{}
	err := administrator.AssignProperties_To_Servers_Administrator_STATUS(dst)
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

var _ genruntime.FromARMConverter = &Servers_Administrator_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (administrator *Servers_Administrator_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_Administrator_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (administrator *Servers_Administrator_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_Administrator_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_Administrator_STATUS_ARM, got %T", armInput)
	}

	// Set property "AdministratorType":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.AdministratorType != nil {
			administratorType := *typedInput.Properties.AdministratorType
			administrator.AdministratorType = &administratorType
		}
	}

	// Set property "AzureADOnlyAuthentication":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.AzureADOnlyAuthentication != nil {
			azureADOnlyAuthentication := *typedInput.Properties.AzureADOnlyAuthentication
			administrator.AzureADOnlyAuthentication = &azureADOnlyAuthentication
		}
	}

	// no assignment for property "Conditions"

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		administrator.Id = &id
	}

	// Set property "Login":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Login != nil {
			login := *typedInput.Properties.Login
			administrator.Login = &login
		}
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		administrator.Name = &name
	}

	// Set property "Sid":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Sid != nil {
			sid := *typedInput.Properties.Sid
			administrator.Sid = &sid
		}
	}

	// Set property "TenantId":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.TenantId != nil {
			tenantId := *typedInput.Properties.TenantId
			administrator.TenantId = &tenantId
		}
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		administrator.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_Servers_Administrator_STATUS populates our Servers_Administrator_STATUS from the provided source Servers_Administrator_STATUS
func (administrator *Servers_Administrator_STATUS) AssignProperties_From_Servers_Administrator_STATUS(source *v20211101s.Servers_Administrator_STATUS) error {

	// AdministratorType
	if source.AdministratorType != nil {
		administratorType := AdministratorProperties_AdministratorType_STATUS(*source.AdministratorType)
		administrator.AdministratorType = &administratorType
	} else {
		administrator.AdministratorType = nil
	}

	// AzureADOnlyAuthentication
	if source.AzureADOnlyAuthentication != nil {
		azureADOnlyAuthentication := *source.AzureADOnlyAuthentication
		administrator.AzureADOnlyAuthentication = &azureADOnlyAuthentication
	} else {
		administrator.AzureADOnlyAuthentication = nil
	}

	// Conditions
	administrator.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	administrator.Id = genruntime.ClonePointerToString(source.Id)

	// Login
	administrator.Login = genruntime.ClonePointerToString(source.Login)

	// Name
	administrator.Name = genruntime.ClonePointerToString(source.Name)

	// Sid
	administrator.Sid = genruntime.ClonePointerToString(source.Sid)

	// TenantId
	administrator.TenantId = genruntime.ClonePointerToString(source.TenantId)

	// Type
	administrator.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_Servers_Administrator_STATUS populates the provided destination Servers_Administrator_STATUS from our Servers_Administrator_STATUS
func (administrator *Servers_Administrator_STATUS) AssignProperties_To_Servers_Administrator_STATUS(destination *v20211101s.Servers_Administrator_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AdministratorType
	if administrator.AdministratorType != nil {
		administratorType := string(*administrator.AdministratorType)
		destination.AdministratorType = &administratorType
	} else {
		destination.AdministratorType = nil
	}

	// AzureADOnlyAuthentication
	if administrator.AzureADOnlyAuthentication != nil {
		azureADOnlyAuthentication := *administrator.AzureADOnlyAuthentication
		destination.AzureADOnlyAuthentication = &azureADOnlyAuthentication
	} else {
		destination.AzureADOnlyAuthentication = nil
	}

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(administrator.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(administrator.Id)

	// Login
	destination.Login = genruntime.ClonePointerToString(administrator.Login)

	// Name
	destination.Name = genruntime.ClonePointerToString(administrator.Name)

	// Sid
	destination.Sid = genruntime.ClonePointerToString(administrator.Sid)

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

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"ActiveDirectory"}
type AdministratorProperties_AdministratorType string

const AdministratorProperties_AdministratorType_ActiveDirectory = AdministratorProperties_AdministratorType("ActiveDirectory")

type AdministratorProperties_AdministratorType_STATUS string

const AdministratorProperties_AdministratorType_STATUS_ActiveDirectory = AdministratorProperties_AdministratorType_STATUS("ActiveDirectory")

func init() {
	SchemeBuilder.Register(&ServersAdministrator{}, &ServersAdministratorList{})
}
