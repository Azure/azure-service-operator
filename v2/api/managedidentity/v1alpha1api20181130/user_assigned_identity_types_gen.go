// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20181130

import (
	"fmt"
	alpha20181130s "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1alpha1api20181130storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Deprecated version of UserAssignedIdentity. Use v1beta20181130.UserAssignedIdentity instead
type UserAssignedIdentity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserAssignedIdentity_Spec   `json:"spec,omitempty"`
	Status            UserAssignedIdentity_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &UserAssignedIdentity{}

// GetConditions returns the conditions of the resource
func (identity *UserAssignedIdentity) GetConditions() conditions.Conditions {
	return identity.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (identity *UserAssignedIdentity) SetConditions(conditions conditions.Conditions) {
	identity.Status.Conditions = conditions
}

var _ conversion.Convertible = &UserAssignedIdentity{}

// ConvertFrom populates our UserAssignedIdentity from the provided hub UserAssignedIdentity
func (identity *UserAssignedIdentity) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source alpha20181130s.UserAssignedIdentity

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = identity.AssignPropertiesFromUserAssignedIdentity(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to identity")
	}

	return nil
}

// ConvertTo populates the provided hub UserAssignedIdentity from our UserAssignedIdentity
func (identity *UserAssignedIdentity) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination alpha20181130s.UserAssignedIdentity
	err := identity.AssignPropertiesToUserAssignedIdentity(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from identity")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-managedidentity-azure-com-v1alpha1api20181130-userassignedidentity,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=managedidentity.azure.com,resources=userassignedidentities,verbs=create;update,versions=v1alpha1api20181130,name=default.v1alpha1api20181130.userassignedidentities.managedidentity.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &UserAssignedIdentity{}

// Default applies defaults to the UserAssignedIdentity resource
func (identity *UserAssignedIdentity) Default() {
	identity.defaultImpl()
	var temp interface{} = identity
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (identity *UserAssignedIdentity) defaultAzureName() {
	if identity.Spec.AzureName == "" {
		identity.Spec.AzureName = identity.Name
	}
}

// defaultImpl applies the code generated defaults to the UserAssignedIdentity resource
func (identity *UserAssignedIdentity) defaultImpl() { identity.defaultAzureName() }

var _ genruntime.KubernetesResource = &UserAssignedIdentity{}

// AzureName returns the Azure name of the resource
func (identity *UserAssignedIdentity) AzureName() string {
	return identity.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-11-30"
func (identity UserAssignedIdentity) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (identity *UserAssignedIdentity) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (identity *UserAssignedIdentity) GetSpec() genruntime.ConvertibleSpec {
	return &identity.Spec
}

// GetStatus returns the status of this resource
func (identity *UserAssignedIdentity) GetStatus() genruntime.ConvertibleStatus {
	return &identity.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ManagedIdentity/userAssignedIdentities"
func (identity *UserAssignedIdentity) GetType() string {
	return "Microsoft.ManagedIdentity/userAssignedIdentities"
}

// NewEmptyStatus returns a new empty (blank) status
func (identity *UserAssignedIdentity) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &UserAssignedIdentity_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (identity *UserAssignedIdentity) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(identity.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  identity.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (identity *UserAssignedIdentity) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*UserAssignedIdentity_STATUS); ok {
		identity.Status = *st
		return nil
	}

	// Convert status to required version
	var st UserAssignedIdentity_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	identity.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-managedidentity-azure-com-v1alpha1api20181130-userassignedidentity,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=managedidentity.azure.com,resources=userassignedidentities,verbs=create;update,versions=v1alpha1api20181130,name=validate.v1alpha1api20181130.userassignedidentities.managedidentity.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &UserAssignedIdentity{}

// ValidateCreate validates the creation of the resource
func (identity *UserAssignedIdentity) ValidateCreate() error {
	validations := identity.createValidations()
	var temp interface{} = identity
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateDelete validates the deletion of the resource
func (identity *UserAssignedIdentity) ValidateDelete() error {
	validations := identity.deleteValidations()
	var temp interface{} = identity
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateUpdate validates an update of the resource
func (identity *UserAssignedIdentity) ValidateUpdate(old runtime.Object) error {
	validations := identity.updateValidations()
	var temp interface{} = identity
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation(old)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// createValidations validates the creation of the resource
func (identity *UserAssignedIdentity) createValidations() []func() error {
	return []func() error{identity.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (identity *UserAssignedIdentity) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (identity *UserAssignedIdentity) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return identity.validateResourceReferences()
		},
		identity.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (identity *UserAssignedIdentity) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&identity.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (identity *UserAssignedIdentity) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*UserAssignedIdentity)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, identity)
}

// AssignPropertiesFromUserAssignedIdentity populates our UserAssignedIdentity from the provided source UserAssignedIdentity
func (identity *UserAssignedIdentity) AssignPropertiesFromUserAssignedIdentity(source *alpha20181130s.UserAssignedIdentity) error {

	// ObjectMeta
	identity.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec UserAssignedIdentity_Spec
	err := spec.AssignPropertiesFromUserAssignedIdentity_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromUserAssignedIdentity_Spec() to populate field Spec")
	}
	identity.Spec = spec

	// Status
	var status UserAssignedIdentity_STATUS
	err = status.AssignPropertiesFromUserAssignedIdentity_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromUserAssignedIdentity_STATUS() to populate field Status")
	}
	identity.Status = status

	// No error
	return nil
}

// AssignPropertiesToUserAssignedIdentity populates the provided destination UserAssignedIdentity from our UserAssignedIdentity
func (identity *UserAssignedIdentity) AssignPropertiesToUserAssignedIdentity(destination *alpha20181130s.UserAssignedIdentity) error {

	// ObjectMeta
	destination.ObjectMeta = *identity.ObjectMeta.DeepCopy()

	// Spec
	var spec alpha20181130s.UserAssignedIdentity_Spec
	err := identity.Spec.AssignPropertiesToUserAssignedIdentity_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToUserAssignedIdentity_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status alpha20181130s.UserAssignedIdentity_STATUS
	err = identity.Status.AssignPropertiesToUserAssignedIdentity_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToUserAssignedIdentity_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (identity *UserAssignedIdentity) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: identity.Spec.OriginalVersion(),
		Kind:    "UserAssignedIdentity",
	}
}

// +kubebuilder:object:root=true
// Deprecated version of UserAssignedIdentity. Use v1beta20181130.UserAssignedIdentity instead
type UserAssignedIdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserAssignedIdentity `json:"items"`
}

// Deprecated version of APIVersion. Use v1beta20181130.APIVersion instead
// +kubebuilder:validation:Enum={"2018-11-30"}
type APIVersion string

const APIVersion_Value = APIVersion("2018-11-30")

type UserAssignedIdentity_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:Required
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	Tags  map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &UserAssignedIdentity_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (identity *UserAssignedIdentity_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if identity == nil {
		return nil, nil
	}
	result := &UserAssignedIdentity_SpecARM{}

	// Set property ‘AzureName’:
	result.AzureName = identity.AzureName

	// Set property ‘Location’:
	if identity.Location != nil {
		location := *identity.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Tags’:
	if identity.Tags != nil {
		result.Tags = make(map[string]string, len(identity.Tags))
		for key, value := range identity.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (identity *UserAssignedIdentity_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &UserAssignedIdentity_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (identity *UserAssignedIdentity_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(UserAssignedIdentity_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected UserAssignedIdentity_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	identity.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		identity.Location = &location
	}

	// Set property ‘Owner’:
	identity.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		identity.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			identity.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &UserAssignedIdentity_Spec{}

// ConvertSpecFrom populates our UserAssignedIdentity_Spec from the provided source
func (identity *UserAssignedIdentity_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*alpha20181130s.UserAssignedIdentity_Spec)
	if ok {
		// Populate our instance from source
		return identity.AssignPropertiesFromUserAssignedIdentity_Spec(src)
	}

	// Convert to an intermediate form
	src = &alpha20181130s.UserAssignedIdentity_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = identity.AssignPropertiesFromUserAssignedIdentity_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our UserAssignedIdentity_Spec
func (identity *UserAssignedIdentity_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*alpha20181130s.UserAssignedIdentity_Spec)
	if ok {
		// Populate destination from our instance
		return identity.AssignPropertiesToUserAssignedIdentity_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20181130s.UserAssignedIdentity_Spec{}
	err := identity.AssignPropertiesToUserAssignedIdentity_Spec(dst)
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

// AssignPropertiesFromUserAssignedIdentity_Spec populates our UserAssignedIdentity_Spec from the provided source UserAssignedIdentity_Spec
func (identity *UserAssignedIdentity_Spec) AssignPropertiesFromUserAssignedIdentity_Spec(source *alpha20181130s.UserAssignedIdentity_Spec) error {

	// AzureName
	identity.AzureName = source.AzureName

	// Location
	identity.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		identity.Owner = &owner
	} else {
		identity.Owner = nil
	}

	// Tags
	identity.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToUserAssignedIdentity_Spec populates the provided destination UserAssignedIdentity_Spec from our UserAssignedIdentity_Spec
func (identity *UserAssignedIdentity_Spec) AssignPropertiesToUserAssignedIdentity_Spec(destination *alpha20181130s.UserAssignedIdentity_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = identity.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(identity.Location)

	// OriginalVersion
	destination.OriginalVersion = identity.OriginalVersion()

	// Owner
	if identity.Owner != nil {
		owner := identity.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(identity.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (identity *UserAssignedIdentity_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (identity *UserAssignedIdentity_Spec) SetAzureName(azureName string) {
	identity.AzureName = azureName
}

// Deprecated version of UserAssignedIdentity_STATUS. Use v1beta20181130.UserAssignedIdentity_STATUS instead
type UserAssignedIdentity_STATUS struct {
	ClientId *string `json:"clientId,omitempty"`

	// Conditions: The observed state of the resource
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Location    *string                `json:"location,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	Tags        map[string]string      `json:"tags,omitempty"`
	TenantId    *string                `json:"tenantId,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &UserAssignedIdentity_STATUS{}

// ConvertStatusFrom populates our UserAssignedIdentity_STATUS from the provided source
func (identity *UserAssignedIdentity_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*alpha20181130s.UserAssignedIdentity_STATUS)
	if ok {
		// Populate our instance from source
		return identity.AssignPropertiesFromUserAssignedIdentity_STATUS(src)
	}

	// Convert to an intermediate form
	src = &alpha20181130s.UserAssignedIdentity_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = identity.AssignPropertiesFromUserAssignedIdentity_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our UserAssignedIdentity_STATUS
func (identity *UserAssignedIdentity_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*alpha20181130s.UserAssignedIdentity_STATUS)
	if ok {
		// Populate destination from our instance
		return identity.AssignPropertiesToUserAssignedIdentity_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20181130s.UserAssignedIdentity_STATUS{}
	err := identity.AssignPropertiesToUserAssignedIdentity_STATUS(dst)
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

var _ genruntime.FromARMConverter = &UserAssignedIdentity_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (identity *UserAssignedIdentity_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &UserAssignedIdentity_STATUSARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (identity *UserAssignedIdentity_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(UserAssignedIdentity_STATUSARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected UserAssignedIdentity_STATUSARM, got %T", armInput)
	}

	// Set property ‘ClientId’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ClientId != nil {
			clientId := *typedInput.Properties.ClientId
			identity.ClientId = &clientId
		}
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		identity.Id = &id
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		identity.Location = &location
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		identity.Name = &name
	}

	// Set property ‘PrincipalId’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.PrincipalId != nil {
			principalId := *typedInput.Properties.PrincipalId
			identity.PrincipalId = &principalId
		}
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		identity.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			identity.Tags[key] = value
		}
	}

	// Set property ‘TenantId’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.TenantId != nil {
			tenantId := *typedInput.Properties.TenantId
			identity.TenantId = &tenantId
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		identity.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromUserAssignedIdentity_STATUS populates our UserAssignedIdentity_STATUS from the provided source UserAssignedIdentity_STATUS
func (identity *UserAssignedIdentity_STATUS) AssignPropertiesFromUserAssignedIdentity_STATUS(source *alpha20181130s.UserAssignedIdentity_STATUS) error {

	// ClientId
	identity.ClientId = genruntime.ClonePointerToString(source.ClientId)

	// Conditions
	identity.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	identity.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	identity.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	identity.Name = genruntime.ClonePointerToString(source.Name)

	// PrincipalId
	identity.PrincipalId = genruntime.ClonePointerToString(source.PrincipalId)

	// Tags
	identity.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// TenantId
	identity.TenantId = genruntime.ClonePointerToString(source.TenantId)

	// Type
	identity.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToUserAssignedIdentity_STATUS populates the provided destination UserAssignedIdentity_STATUS from our UserAssignedIdentity_STATUS
func (identity *UserAssignedIdentity_STATUS) AssignPropertiesToUserAssignedIdentity_STATUS(destination *alpha20181130s.UserAssignedIdentity_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// ClientId
	destination.ClientId = genruntime.ClonePointerToString(identity.ClientId)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(identity.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(identity.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(identity.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(identity.Name)

	// PrincipalId
	destination.PrincipalId = genruntime.ClonePointerToString(identity.PrincipalId)

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(identity.Tags)

	// TenantId
	destination.TenantId = genruntime.ClonePointerToString(identity.TenantId)

	// Type
	destination.Type = genruntime.ClonePointerToString(identity.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&UserAssignedIdentity{}, &UserAssignedIdentityList{})
}
