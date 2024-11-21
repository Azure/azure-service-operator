// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import (
	"fmt"
	arm "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/arm"
	storage "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/rotisserie/eris"
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
	Spec              ServersDatabasesTransparentDataEncryption_Spec   `json:"spec,omitempty"`
	Status            ServersDatabasesTransparentDataEncryption_STATUS `json:"status,omitempty"`
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

var _ configmaps.Exporter = &ServersDatabasesTransparentDataEncryption{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (encryption *ServersDatabasesTransparentDataEncryption) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if encryption.Spec.OperatorSpec == nil {
		return nil
	}
	return encryption.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &ServersDatabasesTransparentDataEncryption{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (encryption *ServersDatabasesTransparentDataEncryption) SecretDestinationExpressions() []*core.DestinationExpression {
	if encryption.Spec.OperatorSpec == nil {
		return nil
	}
	return encryption.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.ImportableResource = &ServersDatabasesTransparentDataEncryption{}

// InitializeSpec initializes the spec for this resource from the given status
func (encryption *ServersDatabasesTransparentDataEncryption) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*ServersDatabasesTransparentDataEncryption_STATUS); ok {
		return encryption.Spec.Initialize_From_ServersDatabasesTransparentDataEncryption_STATUS(s)
	}

	return fmt.Errorf("expected Status of type ServersDatabasesTransparentDataEncryption_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &ServersDatabasesTransparentDataEncryption{}

// AzureName returns the Azure name of the resource (always "current")
func (encryption *ServersDatabasesTransparentDataEncryption) AzureName() string {
	return "current"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (encryption ServersDatabasesTransparentDataEncryption) GetAPIVersion() string {
	return "2021-11-01"
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
	return &ServersDatabasesTransparentDataEncryption_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (encryption *ServersDatabasesTransparentDataEncryption) Owner() *genruntime.ResourceReference {
	if encryption.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(encryption.Spec)
	return encryption.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (encryption *ServersDatabasesTransparentDataEncryption) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ServersDatabasesTransparentDataEncryption_STATUS); ok {
		encryption.Status = *st
		return nil
	}

	// Convert status to required version
	var st ServersDatabasesTransparentDataEncryption_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
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
	return []func() (admission.Warnings, error){encryption.validateResourceReferences, encryption.validateOwnerReference, encryption.validateSecretDestinations, encryption.validateConfigMapDestinations}
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
		func(old runtime.Object) (admission.Warnings, error) {
			return encryption.validateSecretDestinations()
		},
		func(old runtime.Object) (admission.Warnings, error) {
			return encryption.validateConfigMapDestinations()
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (encryption *ServersDatabasesTransparentDataEncryption) validateConfigMapDestinations() (admission.Warnings, error) {
	if encryption.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(encryption, nil, encryption.Spec.OperatorSpec.ConfigMapExpressions)
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

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (encryption *ServersDatabasesTransparentDataEncryption) validateSecretDestinations() (admission.Warnings, error) {
	if encryption.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(encryption, nil, encryption.Spec.OperatorSpec.SecretExpressions)
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
	var spec ServersDatabasesTransparentDataEncryption_Spec
	err := spec.AssignProperties_From_ServersDatabasesTransparentDataEncryption_Spec(&source.Spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_ServersDatabasesTransparentDataEncryption_Spec() to populate field Spec")
	}
	encryption.Spec = spec

	// Status
	var status ServersDatabasesTransparentDataEncryption_STATUS
	err = status.AssignProperties_From_ServersDatabasesTransparentDataEncryption_STATUS(&source.Status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_ServersDatabasesTransparentDataEncryption_STATUS() to populate field Status")
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
	var spec storage.ServersDatabasesTransparentDataEncryption_Spec
	err := encryption.Spec.AssignProperties_To_ServersDatabasesTransparentDataEncryption_Spec(&spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_ServersDatabasesTransparentDataEncryption_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.ServersDatabasesTransparentDataEncryption_STATUS
	err = encryption.Status.AssignProperties_To_ServersDatabasesTransparentDataEncryption_STATUS(&status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_ServersDatabasesTransparentDataEncryption_STATUS() to populate field Status")
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

type ServersDatabasesTransparentDataEncryption_Spec struct {
	// OperatorSpec: The specification for configuring operator behavior. This field is interpreted by the operator and not
	// passed directly to Azure
	OperatorSpec *ServersDatabasesTransparentDataEncryptionOperatorSpec `json:"operatorSpec,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a sql.azure.com/ServersDatabase resource
	Owner *genruntime.KnownResourceReference `group:"sql.azure.com" json:"owner,omitempty" kind:"ServersDatabase"`

	// +kubebuilder:validation:Required
	// State: Specifies the state of the transparent data encryption.
	State *TransparentDataEncryptionProperties_State `json:"state,omitempty"`
}

var _ genruntime.ARMTransformer = &ServersDatabasesTransparentDataEncryption_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (encryption *ServersDatabasesTransparentDataEncryption_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if encryption == nil {
		return nil, nil
	}
	result := &arm.ServersDatabasesTransparentDataEncryption_Spec{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if encryption.State != nil {
		result.Properties = &arm.TransparentDataEncryptionProperties{}
	}
	if encryption.State != nil {
		var temp string
		temp = string(*encryption.State)
		state := arm.TransparentDataEncryptionProperties_State(temp)
		result.Properties.State = &state
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (encryption *ServersDatabasesTransparentDataEncryption_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.ServersDatabasesTransparentDataEncryption_Spec{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (encryption *ServersDatabasesTransparentDataEncryption_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.ServersDatabasesTransparentDataEncryption_Spec)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.ServersDatabasesTransparentDataEncryption_Spec, got %T", armInput)
	}

	// no assignment for property "OperatorSpec"

	// Set property "Owner":
	encryption.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "State":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.State != nil {
			var temp string
			temp = string(*typedInput.Properties.State)
			state := TransparentDataEncryptionProperties_State(temp)
			encryption.State = &state
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &ServersDatabasesTransparentDataEncryption_Spec{}

// ConvertSpecFrom populates our ServersDatabasesTransparentDataEncryption_Spec from the provided source
func (encryption *ServersDatabasesTransparentDataEncryption_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.ServersDatabasesTransparentDataEncryption_Spec)
	if ok {
		// Populate our instance from source
		return encryption.AssignProperties_From_ServersDatabasesTransparentDataEncryption_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.ServersDatabasesTransparentDataEncryption_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = encryption.AssignProperties_From_ServersDatabasesTransparentDataEncryption_Spec(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our ServersDatabasesTransparentDataEncryption_Spec
func (encryption *ServersDatabasesTransparentDataEncryption_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.ServersDatabasesTransparentDataEncryption_Spec)
	if ok {
		// Populate destination from our instance
		return encryption.AssignProperties_To_ServersDatabasesTransparentDataEncryption_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.ServersDatabasesTransparentDataEncryption_Spec{}
	err := encryption.AssignProperties_To_ServersDatabasesTransparentDataEncryption_Spec(dst)
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

// AssignProperties_From_ServersDatabasesTransparentDataEncryption_Spec populates our ServersDatabasesTransparentDataEncryption_Spec from the provided source ServersDatabasesTransparentDataEncryption_Spec
func (encryption *ServersDatabasesTransparentDataEncryption_Spec) AssignProperties_From_ServersDatabasesTransparentDataEncryption_Spec(source *storage.ServersDatabasesTransparentDataEncryption_Spec) error {

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec ServersDatabasesTransparentDataEncryptionOperatorSpec
		err := operatorSpec.AssignProperties_From_ServersDatabasesTransparentDataEncryptionOperatorSpec(source.OperatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_ServersDatabasesTransparentDataEncryptionOperatorSpec() to populate field OperatorSpec")
		}
		encryption.OperatorSpec = &operatorSpec
	} else {
		encryption.OperatorSpec = nil
	}

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

// AssignProperties_To_ServersDatabasesTransparentDataEncryption_Spec populates the provided destination ServersDatabasesTransparentDataEncryption_Spec from our ServersDatabasesTransparentDataEncryption_Spec
func (encryption *ServersDatabasesTransparentDataEncryption_Spec) AssignProperties_To_ServersDatabasesTransparentDataEncryption_Spec(destination *storage.ServersDatabasesTransparentDataEncryption_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// OperatorSpec
	if encryption.OperatorSpec != nil {
		var operatorSpec storage.ServersDatabasesTransparentDataEncryptionOperatorSpec
		err := encryption.OperatorSpec.AssignProperties_To_ServersDatabasesTransparentDataEncryptionOperatorSpec(&operatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_ServersDatabasesTransparentDataEncryptionOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

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

// Initialize_From_ServersDatabasesTransparentDataEncryption_STATUS populates our ServersDatabasesTransparentDataEncryption_Spec from the provided source ServersDatabasesTransparentDataEncryption_STATUS
func (encryption *ServersDatabasesTransparentDataEncryption_Spec) Initialize_From_ServersDatabasesTransparentDataEncryption_STATUS(source *ServersDatabasesTransparentDataEncryption_STATUS) error {

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
func (encryption *ServersDatabasesTransparentDataEncryption_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

type ServersDatabasesTransparentDataEncryption_STATUS struct {
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

var _ genruntime.ConvertibleStatus = &ServersDatabasesTransparentDataEncryption_STATUS{}

// ConvertStatusFrom populates our ServersDatabasesTransparentDataEncryption_STATUS from the provided source
func (encryption *ServersDatabasesTransparentDataEncryption_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.ServersDatabasesTransparentDataEncryption_STATUS)
	if ok {
		// Populate our instance from source
		return encryption.AssignProperties_From_ServersDatabasesTransparentDataEncryption_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.ServersDatabasesTransparentDataEncryption_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = encryption.AssignProperties_From_ServersDatabasesTransparentDataEncryption_STATUS(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our ServersDatabasesTransparentDataEncryption_STATUS
func (encryption *ServersDatabasesTransparentDataEncryption_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.ServersDatabasesTransparentDataEncryption_STATUS)
	if ok {
		// Populate destination from our instance
		return encryption.AssignProperties_To_ServersDatabasesTransparentDataEncryption_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.ServersDatabasesTransparentDataEncryption_STATUS{}
	err := encryption.AssignProperties_To_ServersDatabasesTransparentDataEncryption_STATUS(dst)
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

var _ genruntime.FromARMConverter = &ServersDatabasesTransparentDataEncryption_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (encryption *ServersDatabasesTransparentDataEncryption_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.ServersDatabasesTransparentDataEncryption_STATUS{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (encryption *ServersDatabasesTransparentDataEncryption_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.ServersDatabasesTransparentDataEncryption_STATUS)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.ServersDatabasesTransparentDataEncryption_STATUS, got %T", armInput)
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
			var temp string
			temp = string(*typedInput.Properties.State)
			state := TransparentDataEncryptionProperties_State_STATUS(temp)
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

// AssignProperties_From_ServersDatabasesTransparentDataEncryption_STATUS populates our ServersDatabasesTransparentDataEncryption_STATUS from the provided source ServersDatabasesTransparentDataEncryption_STATUS
func (encryption *ServersDatabasesTransparentDataEncryption_STATUS) AssignProperties_From_ServersDatabasesTransparentDataEncryption_STATUS(source *storage.ServersDatabasesTransparentDataEncryption_STATUS) error {

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

// AssignProperties_To_ServersDatabasesTransparentDataEncryption_STATUS populates the provided destination ServersDatabasesTransparentDataEncryption_STATUS from our ServersDatabasesTransparentDataEncryption_STATUS
func (encryption *ServersDatabasesTransparentDataEncryption_STATUS) AssignProperties_To_ServersDatabasesTransparentDataEncryption_STATUS(destination *storage.ServersDatabasesTransparentDataEncryption_STATUS) error {
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

// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type ServersDatabasesTransparentDataEncryptionOperatorSpec struct {
	// ConfigMapExpressions: configures where to place operator written dynamic ConfigMaps (created with CEL expressions).
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`

	// SecretExpressions: configures where to place operator written dynamic secrets (created with CEL expressions).
	SecretExpressions []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_ServersDatabasesTransparentDataEncryptionOperatorSpec populates our ServersDatabasesTransparentDataEncryptionOperatorSpec from the provided source ServersDatabasesTransparentDataEncryptionOperatorSpec
func (operator *ServersDatabasesTransparentDataEncryptionOperatorSpec) AssignProperties_From_ServersDatabasesTransparentDataEncryptionOperatorSpec(source *storage.ServersDatabasesTransparentDataEncryptionOperatorSpec) error {

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

	// No error
	return nil
}

// AssignProperties_To_ServersDatabasesTransparentDataEncryptionOperatorSpec populates the provided destination ServersDatabasesTransparentDataEncryptionOperatorSpec from our ServersDatabasesTransparentDataEncryptionOperatorSpec
func (operator *ServersDatabasesTransparentDataEncryptionOperatorSpec) AssignProperties_To_ServersDatabasesTransparentDataEncryptionOperatorSpec(destination *storage.ServersDatabasesTransparentDataEncryptionOperatorSpec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

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
