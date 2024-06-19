// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import (
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/storage"
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
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/TransparentDataEncryptions.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/transparentDataEncryption/{tdeName}
type ServersDatabasesTransparentDataEncryption struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Servers_Databases_TransparentDataEncryption_Spec   `json:"spec,omitempty"`
	Status            Servers_Databases_TransparentDataEncryption_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ServersDatabasesTransparentDataEncryption{}

// GetConditions returns the conditions of the resource
func (encryption *ServersDatabasesTransparentDataEncryption) GetConditions() conditions.Conditions {
	return encryption.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (encryption *ServersDatabasesTransparentDataEncryption) SetConditions(conditions conditions.Conditions) {
	encryption.Status.Conditions = conditions
}

var _ conversion.Convertible = &ServersDatabasesTransparentDataEncryption{}

// ConvertFrom populates our ServersDatabasesTransparentDataEncryption from the provided hub ServersDatabasesTransparentDataEncryption
func (encryption *ServersDatabasesTransparentDataEncryption) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.ServersDatabasesTransparentDataEncryption)
	if !ok {
		return fmt.Errorf("expected sql/v1api20211101/storage/ServersDatabasesTransparentDataEncryption but received %T instead", hub)
	}

	return encryption.AssignProperties_From_ServersDatabasesTransparentDataEncryption(source)
}

// ConvertTo populates the provided hub ServersDatabasesTransparentDataEncryption from our ServersDatabasesTransparentDataEncryption
func (encryption *ServersDatabasesTransparentDataEncryption) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.ServersDatabasesTransparentDataEncryption)
	if !ok {
		return fmt.Errorf("expected sql/v1api20211101/storage/ServersDatabasesTransparentDataEncryption but received %T instead", hub)
	}

	return encryption.AssignProperties_To_ServersDatabasesTransparentDataEncryption(destination)
}

// +kubebuilder:webhook:path=/mutate-sql-azure-com-v1api20211101-serversdatabasestransparentdataencryption,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversdatabasestransparentdataencryptions,verbs=create;update,versions=v1api20211101,name=default.v1api20211101.serversdatabasestransparentdataencryptions.sql.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &ServersDatabasesTransparentDataEncryption{}

// Default applies defaults to the ServersDatabasesTransparentDataEncryption resource
func (encryption *ServersDatabasesTransparentDataEncryption) Default() {
	encryption.defaultImpl()
	var temp any = encryption
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the ServersDatabasesTransparentDataEncryption resource
func (encryption *ServersDatabasesTransparentDataEncryption) defaultImpl() {}

var _ genruntime.ImportableResource = &ServersDatabasesTransparentDataEncryption{}

// InitializeSpec initializes the spec for this resource from the given status
func (encryption *ServersDatabasesTransparentDataEncryption) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*Servers_Databases_TransparentDataEncryption_STATUS); ok {
		return encryption.Spec.Initialize_From_Servers_Databases_TransparentDataEncryption_STATUS(s)
	}

	return fmt.Errorf("expected Status of type Servers_Databases_TransparentDataEncryption_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &ServersDatabasesTransparentDataEncryption{}

// AzureName returns the Azure name of the resource (always "current")
func (encryption *ServersDatabasesTransparentDataEncryption) AzureName() string {
	return "current"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (encryption ServersDatabasesTransparentDataEncryption) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (encryption *ServersDatabasesTransparentDataEncryption) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (encryption *ServersDatabasesTransparentDataEncryption) GetSpec() genruntime.ConvertibleSpec {
	return &encryption.Spec
}

// GetStatus returns the status of this resource
func (encryption *ServersDatabasesTransparentDataEncryption) GetStatus() genruntime.ConvertibleStatus {
	return &encryption.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (encryption *ServersDatabasesTransparentDataEncryption) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/databases/transparentDataEncryption"
func (encryption *ServersDatabasesTransparentDataEncryption) GetType() string {
	return "Microsoft.Sql/servers/databases/transparentDataEncryption"
}

// NewEmptyStatus returns a new empty (blank) status
func (encryption *ServersDatabasesTransparentDataEncryption) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Servers_Databases_TransparentDataEncryption_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (encryption *ServersDatabasesTransparentDataEncryption) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(encryption.Spec)
	return encryption.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (encryption *ServersDatabasesTransparentDataEncryption) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Servers_Databases_TransparentDataEncryption_STATUS); ok {
		encryption.Status = *st
		return nil
	}

	// Convert status to required version
	var st Servers_Databases_TransparentDataEncryption_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	encryption.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-sql-azure-com-v1api20211101-serversdatabasestransparentdataencryption,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversdatabasestransparentdataencryptions,verbs=create;update,versions=v1api20211101,name=validate.v1api20211101.serversdatabasestransparentdataencryptions.sql.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &ServersDatabasesTransparentDataEncryption{}

// ValidateCreate validates the creation of the resource
func (encryption *ServersDatabasesTransparentDataEncryption) ValidateCreate() (admission.Warnings, error) {
	validations := encryption.createValidations()
	var temp any = encryption
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (encryption *ServersDatabasesTransparentDataEncryption) ValidateDelete() (admission.Warnings, error) {
	validations := encryption.deleteValidations()
	var temp any = encryption
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (encryption *ServersDatabasesTransparentDataEncryption) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := encryption.updateValidations()
	var temp any = encryption
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (encryption *ServersDatabasesTransparentDataEncryption) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){encryption.validateResourceReferences, encryption.validateOwnerReference}
}

// deleteValidations validates the deletion of the resource
func (encryption *ServersDatabasesTransparentDataEncryption) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (encryption *ServersDatabasesTransparentDataEncryption) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return encryption.validateResourceReferences()
		},
		encryption.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return encryption.validateOwnerReference()
		},
	}
}

// validateOwnerReference validates the owner field
func (encryption *ServersDatabasesTransparentDataEncryption) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(encryption)
}

// validateResourceReferences validates all resource references
func (encryption *ServersDatabasesTransparentDataEncryption) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&encryption.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (encryption *ServersDatabasesTransparentDataEncryption) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*ServersDatabasesTransparentDataEncryption)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, encryption)
}

// AssignProperties_From_ServersDatabasesTransparentDataEncryption populates our ServersDatabasesTransparentDataEncryption from the provided source ServersDatabasesTransparentDataEncryption
func (encryption *ServersDatabasesTransparentDataEncryption) AssignProperties_From_ServersDatabasesTransparentDataEncryption(source *storage.ServersDatabasesTransparentDataEncryption) error {

	// ObjectMeta
	encryption.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Servers_Databases_TransparentDataEncryption_Spec
	err := spec.AssignProperties_From_Servers_Databases_TransparentDataEncryption_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_Databases_TransparentDataEncryption_Spec() to populate field Spec")
	}
	encryption.Spec = spec

	// Status
	var status Servers_Databases_TransparentDataEncryption_STATUS
	err = status.AssignProperties_From_Servers_Databases_TransparentDataEncryption_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_Databases_TransparentDataEncryption_STATUS() to populate field Status")
	}
	encryption.Status = status

	// No error
	return nil
}

// AssignProperties_To_ServersDatabasesTransparentDataEncryption populates the provided destination ServersDatabasesTransparentDataEncryption from our ServersDatabasesTransparentDataEncryption
func (encryption *ServersDatabasesTransparentDataEncryption) AssignProperties_To_ServersDatabasesTransparentDataEncryption(destination *storage.ServersDatabasesTransparentDataEncryption) error {

	// ObjectMeta
	destination.ObjectMeta = *encryption.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.Servers_Databases_TransparentDataEncryption_Spec
	err := encryption.Spec.AssignProperties_To_Servers_Databases_TransparentDataEncryption_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_Databases_TransparentDataEncryption_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.Servers_Databases_TransparentDataEncryption_STATUS
	err = encryption.Status.AssignProperties_To_Servers_Databases_TransparentDataEncryption_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_Databases_TransparentDataEncryption_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (encryption *ServersDatabasesTransparentDataEncryption) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: encryption.Spec.OriginalVersion(),
		Kind:    "ServersDatabasesTransparentDataEncryption",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/TransparentDataEncryptions.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/transparentDataEncryption/{tdeName}
type ServersDatabasesTransparentDataEncryptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServersDatabasesTransparentDataEncryption `json:"items"`
}

type Servers_Databases_TransparentDataEncryption_Spec struct {
	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a sql.azure.com/ServersDatabase resource
	Owner *genruntime.KnownResourceReference `group:"sql.azure.com" json:"owner,omitempty" kind:"ServersDatabase"`

	// +kubebuilder:validation:Required
	// State: Specifies the state of the transparent data encryption.
	State *TransparentDataEncryptionProperties_State `json:"state,omitempty"`
}

var _ genruntime.ARMTransformer = &Servers_Databases_TransparentDataEncryption_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (encryption *Servers_Databases_TransparentDataEncryption_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if encryption == nil {
		return nil, nil
	}
	result := &Servers_Databases_TransparentDataEncryption_Spec_ARM{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if encryption.State != nil {
		result.Properties = &TransparentDataEncryptionProperties_ARM{}
	}
	if encryption.State != nil {
		state := *encryption.State
		result.Properties.State = &state
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (encryption *Servers_Databases_TransparentDataEncryption_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_Databases_TransparentDataEncryption_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (encryption *Servers_Databases_TransparentDataEncryption_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_Databases_TransparentDataEncryption_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_Databases_TransparentDataEncryption_Spec_ARM, got %T", armInput)
	}

	// Set property "Owner":
	encryption.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "State":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.State != nil {
			state := *typedInput.Properties.State
			encryption.State = &state
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Servers_Databases_TransparentDataEncryption_Spec{}

// ConvertSpecFrom populates our Servers_Databases_TransparentDataEncryption_Spec from the provided source
func (encryption *Servers_Databases_TransparentDataEncryption_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.Servers_Databases_TransparentDataEncryption_Spec)
	if ok {
		// Populate our instance from source
		return encryption.AssignProperties_From_Servers_Databases_TransparentDataEncryption_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.Servers_Databases_TransparentDataEncryption_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = encryption.AssignProperties_From_Servers_Databases_TransparentDataEncryption_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Servers_Databases_TransparentDataEncryption_Spec
func (encryption *Servers_Databases_TransparentDataEncryption_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.Servers_Databases_TransparentDataEncryption_Spec)
	if ok {
		// Populate destination from our instance
		return encryption.AssignProperties_To_Servers_Databases_TransparentDataEncryption_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.Servers_Databases_TransparentDataEncryption_Spec{}
	err := encryption.AssignProperties_To_Servers_Databases_TransparentDataEncryption_Spec(dst)
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

// AssignProperties_From_Servers_Databases_TransparentDataEncryption_Spec populates our Servers_Databases_TransparentDataEncryption_Spec from the provided source Servers_Databases_TransparentDataEncryption_Spec
func (encryption *Servers_Databases_TransparentDataEncryption_Spec) AssignProperties_From_Servers_Databases_TransparentDataEncryption_Spec(source *storage.Servers_Databases_TransparentDataEncryption_Spec) error {

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		encryption.Owner = &owner
	} else {
		encryption.Owner = nil
	}

	// State
	if source.State != nil {
		state := *source.State
		stateTemp := genruntime.ToEnum(state, transparentDataEncryptionProperties_State_Values)
		encryption.State = &stateTemp
	} else {
		encryption.State = nil
	}

	// No error
	return nil
}

// AssignProperties_To_Servers_Databases_TransparentDataEncryption_Spec populates the provided destination Servers_Databases_TransparentDataEncryption_Spec from our Servers_Databases_TransparentDataEncryption_Spec
func (encryption *Servers_Databases_TransparentDataEncryption_Spec) AssignProperties_To_Servers_Databases_TransparentDataEncryption_Spec(destination *storage.Servers_Databases_TransparentDataEncryption_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// OriginalVersion
	destination.OriginalVersion = encryption.OriginalVersion()

	// Owner
	if encryption.Owner != nil {
		owner := encryption.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// State
	if encryption.State != nil {
		state := string(*encryption.State)
		destination.State = &state
	} else {
		destination.State = nil
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

// Initialize_From_Servers_Databases_TransparentDataEncryption_STATUS populates our Servers_Databases_TransparentDataEncryption_Spec from the provided source Servers_Databases_TransparentDataEncryption_STATUS
func (encryption *Servers_Databases_TransparentDataEncryption_Spec) Initialize_From_Servers_Databases_TransparentDataEncryption_STATUS(source *Servers_Databases_TransparentDataEncryption_STATUS) error {

	// State
	if source.State != nil {
		state := genruntime.ToEnum(string(*source.State), transparentDataEncryptionProperties_State_Values)
		encryption.State = &state
	} else {
		encryption.State = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (encryption *Servers_Databases_TransparentDataEncryption_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

type Servers_Databases_TransparentDataEncryption_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// State: Specifies the state of the transparent data encryption.
	State *TransparentDataEncryptionProperties_State_STATUS `json:"state,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Servers_Databases_TransparentDataEncryption_STATUS{}

// ConvertStatusFrom populates our Servers_Databases_TransparentDataEncryption_STATUS from the provided source
func (encryption *Servers_Databases_TransparentDataEncryption_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.Servers_Databases_TransparentDataEncryption_STATUS)
	if ok {
		// Populate our instance from source
		return encryption.AssignProperties_From_Servers_Databases_TransparentDataEncryption_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.Servers_Databases_TransparentDataEncryption_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = encryption.AssignProperties_From_Servers_Databases_TransparentDataEncryption_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Servers_Databases_TransparentDataEncryption_STATUS
func (encryption *Servers_Databases_TransparentDataEncryption_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.Servers_Databases_TransparentDataEncryption_STATUS)
	if ok {
		// Populate destination from our instance
		return encryption.AssignProperties_To_Servers_Databases_TransparentDataEncryption_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.Servers_Databases_TransparentDataEncryption_STATUS{}
	err := encryption.AssignProperties_To_Servers_Databases_TransparentDataEncryption_STATUS(dst)
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

var _ genruntime.FromARMConverter = &Servers_Databases_TransparentDataEncryption_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (encryption *Servers_Databases_TransparentDataEncryption_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_Databases_TransparentDataEncryption_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (encryption *Servers_Databases_TransparentDataEncryption_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_Databases_TransparentDataEncryption_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_Databases_TransparentDataEncryption_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		encryption.Id = &id
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		encryption.Name = &name
	}

	// Set property "State":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.State != nil {
			state := *typedInput.Properties.State
			encryption.State = &state
		}
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		encryption.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_Servers_Databases_TransparentDataEncryption_STATUS populates our Servers_Databases_TransparentDataEncryption_STATUS from the provided source Servers_Databases_TransparentDataEncryption_STATUS
func (encryption *Servers_Databases_TransparentDataEncryption_STATUS) AssignProperties_From_Servers_Databases_TransparentDataEncryption_STATUS(source *storage.Servers_Databases_TransparentDataEncryption_STATUS) error {

	// Conditions
	encryption.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	encryption.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	encryption.Name = genruntime.ClonePointerToString(source.Name)

	// State
	if source.State != nil {
		state := *source.State
		stateTemp := genruntime.ToEnum(state, transparentDataEncryptionProperties_State_STATUS_Values)
		encryption.State = &stateTemp
	} else {
		encryption.State = nil
	}

	// Type
	encryption.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_Servers_Databases_TransparentDataEncryption_STATUS populates the provided destination Servers_Databases_TransparentDataEncryption_STATUS from our Servers_Databases_TransparentDataEncryption_STATUS
func (encryption *Servers_Databases_TransparentDataEncryption_STATUS) AssignProperties_To_Servers_Databases_TransparentDataEncryption_STATUS(destination *storage.Servers_Databases_TransparentDataEncryption_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(encryption.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(encryption.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(encryption.Name)

	// State
	if encryption.State != nil {
		state := string(*encryption.State)
		destination.State = &state
	} else {
		destination.State = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(encryption.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type TransparentDataEncryptionProperties_State string

const (
	TransparentDataEncryptionProperties_State_Disabled = TransparentDataEncryptionProperties_State("Disabled")
	TransparentDataEncryptionProperties_State_Enabled  = TransparentDataEncryptionProperties_State("Enabled")
)

// Mapping from string to TransparentDataEncryptionProperties_State
var transparentDataEncryptionProperties_State_Values = map[string]TransparentDataEncryptionProperties_State{
	"disabled": TransparentDataEncryptionProperties_State_Disabled,
	"enabled":  TransparentDataEncryptionProperties_State_Enabled,
}

type TransparentDataEncryptionProperties_State_STATUS string

const (
	TransparentDataEncryptionProperties_State_STATUS_Disabled = TransparentDataEncryptionProperties_State_STATUS("Disabled")
	TransparentDataEncryptionProperties_State_STATUS_Enabled  = TransparentDataEncryptionProperties_State_STATUS("Enabled")
)

// Mapping from string to TransparentDataEncryptionProperties_State_STATUS
var transparentDataEncryptionProperties_State_STATUS_Values = map[string]TransparentDataEncryptionProperties_State_STATUS{
	"disabled": TransparentDataEncryptionProperties_State_STATUS_Disabled,
	"enabled":  TransparentDataEncryptionProperties_State_STATUS_Enabled,
}

func init() {
	SchemeBuilder.Register(&ServersDatabasesTransparentDataEncryption{}, &ServersDatabasesTransparentDataEncryptionList{})
}
