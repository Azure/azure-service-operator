// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20181130

import (
	"context"
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20181130/storage"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
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
// - Generated from: /msi/resource-manager/Microsoft.ManagedIdentity/stable/2018-11-30/ManagedIdentity.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}
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
	var source storage.UserAssignedIdentity

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = identity.AssignProperties_From_UserAssignedIdentity(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to identity")
	}

	return nil
}

// ConvertTo populates the provided hub UserAssignedIdentity from our UserAssignedIdentity
func (identity *UserAssignedIdentity) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination storage.UserAssignedIdentity
	err := identity.AssignProperties_To_UserAssignedIdentity(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from identity")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-managedidentity-azure-com-v1api20181130-userassignedidentity,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=managedidentity.azure.com,resources=userassignedidentities,verbs=create;update,versions=v1api20181130,name=default.v1api20181130.userassignedidentities.managedidentity.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &UserAssignedIdentity{}

// Default applies defaults to the UserAssignedIdentity resource
func (identity *UserAssignedIdentity) Default() {
	identity.defaultImpl()
	var temp any = identity
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

var _ genruntime.KubernetesExporter = &UserAssignedIdentity{}

// ExportKubernetesResources defines a resource which can create other resources in Kubernetes.
func (identity *UserAssignedIdentity) ExportKubernetesResources(_ context.Context, _ genruntime.MetaObject, _ *genericarmclient.GenericClient, _ logr.Logger) ([]client.Object, error) {
	collector := configmaps.NewCollector(identity.Namespace)
	if identity.Spec.OperatorSpec != nil && identity.Spec.OperatorSpec.ConfigMaps != nil {
		if identity.Status.ClientId != nil {
			collector.AddValue(identity.Spec.OperatorSpec.ConfigMaps.ClientId, *identity.Status.ClientId)
		}
	}
	if identity.Spec.OperatorSpec != nil && identity.Spec.OperatorSpec.ConfigMaps != nil {
		if identity.Status.PrincipalId != nil {
			collector.AddValue(identity.Spec.OperatorSpec.ConfigMaps.PrincipalId, *identity.Status.PrincipalId)
		}
	}
	if identity.Spec.OperatorSpec != nil && identity.Spec.OperatorSpec.ConfigMaps != nil {
		if identity.Status.TenantId != nil {
			collector.AddValue(identity.Spec.OperatorSpec.ConfigMaps.TenantId, *identity.Status.TenantId)
		}
	}
	result, err := collector.Values()
	if err != nil {
		return nil, err
	}
	return configmaps.SliceToClientObjectSlice(result), nil
}

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

// GetSupportedOperations returns the operations supported by the resource
func (identity *UserAssignedIdentity) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ManagedIdentity/userAssignedIdentities"
func (identity *UserAssignedIdentity) GetType() string {
	return "Microsoft.ManagedIdentity/userAssignedIdentities"
}

// NewEmptyStatus returns a new empty (blank) status
func (identity *UserAssignedIdentity) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &UserAssignedIdentity_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (identity *UserAssignedIdentity) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(identity.Spec)
	return identity.Spec.Owner.AsResourceReference(group, kind)
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

// +kubebuilder:webhook:path=/validate-managedidentity-azure-com-v1api20181130-userassignedidentity,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=managedidentity.azure.com,resources=userassignedidentities,verbs=create;update,versions=v1api20181130,name=validate.v1api20181130.userassignedidentities.managedidentity.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &UserAssignedIdentity{}

// ValidateCreate validates the creation of the resource
func (identity *UserAssignedIdentity) ValidateCreate() (admission.Warnings, error) {
	validations := identity.createValidations()
	var temp any = identity
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (identity *UserAssignedIdentity) ValidateDelete() (admission.Warnings, error) {
	validations := identity.deleteValidations()
	var temp any = identity
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (identity *UserAssignedIdentity) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := identity.updateValidations()
	var temp any = identity
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (identity *UserAssignedIdentity) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){identity.validateResourceReferences, identity.validateOwnerReference, identity.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (identity *UserAssignedIdentity) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (identity *UserAssignedIdentity) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return identity.validateResourceReferences()
		},
		identity.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return identity.validateOwnerReference()
		},
		func(old runtime.Object) (admission.Warnings, error) {
			return identity.validateConfigMapDestinations()
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (identity *UserAssignedIdentity) validateConfigMapDestinations() (admission.Warnings, error) {
	if identity.Spec.OperatorSpec == nil {
		return nil, nil
	}
	if identity.Spec.OperatorSpec.ConfigMaps == nil {
		return nil, nil
	}
	toValidate := []*genruntime.ConfigMapDestination{
		identity.Spec.OperatorSpec.ConfigMaps.ClientId,
		identity.Spec.OperatorSpec.ConfigMaps.PrincipalId,
		identity.Spec.OperatorSpec.ConfigMaps.TenantId,
	}
	return configmaps.ValidateDestinations(toValidate)
}

// validateOwnerReference validates the owner field
func (identity *UserAssignedIdentity) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(identity)
}

// validateResourceReferences validates all resource references
func (identity *UserAssignedIdentity) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&identity.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (identity *UserAssignedIdentity) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*UserAssignedIdentity)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, identity)
}

// AssignProperties_From_UserAssignedIdentity populates our UserAssignedIdentity from the provided source UserAssignedIdentity
func (identity *UserAssignedIdentity) AssignProperties_From_UserAssignedIdentity(source *storage.UserAssignedIdentity) error {

	// ObjectMeta
	identity.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec UserAssignedIdentity_Spec
	err := spec.AssignProperties_From_UserAssignedIdentity_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_UserAssignedIdentity_Spec() to populate field Spec")
	}
	identity.Spec = spec

	// Status
	var status UserAssignedIdentity_STATUS
	err = status.AssignProperties_From_UserAssignedIdentity_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_UserAssignedIdentity_STATUS() to populate field Status")
	}
	identity.Status = status

	// No error
	return nil
}

// AssignProperties_To_UserAssignedIdentity populates the provided destination UserAssignedIdentity from our UserAssignedIdentity
func (identity *UserAssignedIdentity) AssignProperties_To_UserAssignedIdentity(destination *storage.UserAssignedIdentity) error {

	// ObjectMeta
	destination.ObjectMeta = *identity.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.UserAssignedIdentity_Spec
	err := identity.Spec.AssignProperties_To_UserAssignedIdentity_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_UserAssignedIdentity_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.UserAssignedIdentity_STATUS
	err = identity.Status.AssignProperties_To_UserAssignedIdentity_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_UserAssignedIdentity_STATUS() to populate field Status")
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
// Generator information:
// - Generated from: /msi/resource-manager/Microsoft.ManagedIdentity/stable/2018-11-30/ManagedIdentity.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}
type UserAssignedIdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserAssignedIdentity `json:"items"`
}

// +kubebuilder:validation:Enum={"2018-11-30"}
type APIVersion string

const APIVersion_Value = APIVersion("2018-11-30")

type UserAssignedIdentity_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:Required
	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// OperatorSpec: The specification for configuring operator behavior. This field is interpreted by the operator and not
	// passed directly to Azure
	OperatorSpec *UserAssignedIdentityOperatorSpec `json:"operatorSpec,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &UserAssignedIdentity_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (identity *UserAssignedIdentity_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if identity == nil {
		return nil, nil
	}
	result := &UserAssignedIdentity_Spec_ARM{}

	// Set property "Location":
	if identity.Location != nil {
		location := *identity.Location
		result.Location = &location
	}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Tags":
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
	return &UserAssignedIdentity_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (identity *UserAssignedIdentity_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(UserAssignedIdentity_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected UserAssignedIdentity_Spec_ARM, got %T", armInput)
	}

	// Set property "AzureName":
	identity.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property "Location":
	if typedInput.Location != nil {
		location := *typedInput.Location
		identity.Location = &location
	}

	// no assignment for property "OperatorSpec"

	// Set property "Owner":
	identity.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "Tags":
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
	src, ok := source.(*storage.UserAssignedIdentity_Spec)
	if ok {
		// Populate our instance from source
		return identity.AssignProperties_From_UserAssignedIdentity_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.UserAssignedIdentity_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = identity.AssignProperties_From_UserAssignedIdentity_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our UserAssignedIdentity_Spec
func (identity *UserAssignedIdentity_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.UserAssignedIdentity_Spec)
	if ok {
		// Populate destination from our instance
		return identity.AssignProperties_To_UserAssignedIdentity_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.UserAssignedIdentity_Spec{}
	err := identity.AssignProperties_To_UserAssignedIdentity_Spec(dst)
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

// AssignProperties_From_UserAssignedIdentity_Spec populates our UserAssignedIdentity_Spec from the provided source UserAssignedIdentity_Spec
func (identity *UserAssignedIdentity_Spec) AssignProperties_From_UserAssignedIdentity_Spec(source *storage.UserAssignedIdentity_Spec) error {

	// AzureName
	identity.AzureName = source.AzureName

	// Location
	identity.Location = genruntime.ClonePointerToString(source.Location)

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec UserAssignedIdentityOperatorSpec
		err := operatorSpec.AssignProperties_From_UserAssignedIdentityOperatorSpec(source.OperatorSpec)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_UserAssignedIdentityOperatorSpec() to populate field OperatorSpec")
		}
		identity.OperatorSpec = &operatorSpec
	} else {
		identity.OperatorSpec = nil
	}

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

// AssignProperties_To_UserAssignedIdentity_Spec populates the provided destination UserAssignedIdentity_Spec from our UserAssignedIdentity_Spec
func (identity *UserAssignedIdentity_Spec) AssignProperties_To_UserAssignedIdentity_Spec(destination *storage.UserAssignedIdentity_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = identity.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(identity.Location)

	// OperatorSpec
	if identity.OperatorSpec != nil {
		var operatorSpec storage.UserAssignedIdentityOperatorSpec
		err := identity.OperatorSpec.AssignProperties_To_UserAssignedIdentityOperatorSpec(&operatorSpec)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_UserAssignedIdentityOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

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

type UserAssignedIdentity_STATUS struct {
	// ClientId: The id of the app associated with the identity. This is a random generated UUID by MSI.
	ClientId *string `json:"clientId,omitempty"`

	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// PrincipalId: The id of the service principal object associated with the created identity.
	PrincipalId *string `json:"principalId,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// TenantId: The id of the tenant which the identity belongs to.
	TenantId *string `json:"tenantId,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &UserAssignedIdentity_STATUS{}

// ConvertStatusFrom populates our UserAssignedIdentity_STATUS from the provided source
func (identity *UserAssignedIdentity_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.UserAssignedIdentity_STATUS)
	if ok {
		// Populate our instance from source
		return identity.AssignProperties_From_UserAssignedIdentity_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.UserAssignedIdentity_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = identity.AssignProperties_From_UserAssignedIdentity_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our UserAssignedIdentity_STATUS
func (identity *UserAssignedIdentity_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.UserAssignedIdentity_STATUS)
	if ok {
		// Populate destination from our instance
		return identity.AssignProperties_To_UserAssignedIdentity_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.UserAssignedIdentity_STATUS{}
	err := identity.AssignProperties_To_UserAssignedIdentity_STATUS(dst)
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
	return &UserAssignedIdentity_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (identity *UserAssignedIdentity_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(UserAssignedIdentity_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected UserAssignedIdentity_STATUS_ARM, got %T", armInput)
	}

	// Set property "ClientId":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ClientId != nil {
			clientId := *typedInput.Properties.ClientId
			identity.ClientId = &clientId
		}
	}

	// no assignment for property "Conditions"

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		identity.Id = &id
	}

	// Set property "Location":
	if typedInput.Location != nil {
		location := *typedInput.Location
		identity.Location = &location
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		identity.Name = &name
	}

	// Set property "PrincipalId":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.PrincipalId != nil {
			principalId := *typedInput.Properties.PrincipalId
			identity.PrincipalId = &principalId
		}
	}

	// Set property "Tags":
	if typedInput.Tags != nil {
		identity.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			identity.Tags[key] = value
		}
	}

	// Set property "TenantId":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.TenantId != nil {
			tenantId := *typedInput.Properties.TenantId
			identity.TenantId = &tenantId
		}
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		identity.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_UserAssignedIdentity_STATUS populates our UserAssignedIdentity_STATUS from the provided source UserAssignedIdentity_STATUS
func (identity *UserAssignedIdentity_STATUS) AssignProperties_From_UserAssignedIdentity_STATUS(source *storage.UserAssignedIdentity_STATUS) error {

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

// AssignProperties_To_UserAssignedIdentity_STATUS populates the provided destination UserAssignedIdentity_STATUS from our UserAssignedIdentity_STATUS
func (identity *UserAssignedIdentity_STATUS) AssignProperties_To_UserAssignedIdentity_STATUS(destination *storage.UserAssignedIdentity_STATUS) error {
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

// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type UserAssignedIdentityOperatorSpec struct {
	// ConfigMaps: configures where to place operator written ConfigMaps.
	ConfigMaps *UserAssignedIdentityOperatorConfigMaps `json:"configMaps,omitempty"`
}

// AssignProperties_From_UserAssignedIdentityOperatorSpec populates our UserAssignedIdentityOperatorSpec from the provided source UserAssignedIdentityOperatorSpec
func (operator *UserAssignedIdentityOperatorSpec) AssignProperties_From_UserAssignedIdentityOperatorSpec(source *storage.UserAssignedIdentityOperatorSpec) error {

	// ConfigMaps
	if source.ConfigMaps != nil {
		var configMap UserAssignedIdentityOperatorConfigMaps
		err := configMap.AssignProperties_From_UserAssignedIdentityOperatorConfigMaps(source.ConfigMaps)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_UserAssignedIdentityOperatorConfigMaps() to populate field ConfigMaps")
		}
		operator.ConfigMaps = &configMap
	} else {
		operator.ConfigMaps = nil
	}

	// No error
	return nil
}

// AssignProperties_To_UserAssignedIdentityOperatorSpec populates the provided destination UserAssignedIdentityOperatorSpec from our UserAssignedIdentityOperatorSpec
func (operator *UserAssignedIdentityOperatorSpec) AssignProperties_To_UserAssignedIdentityOperatorSpec(destination *storage.UserAssignedIdentityOperatorSpec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// ConfigMaps
	if operator.ConfigMaps != nil {
		var configMap storage.UserAssignedIdentityOperatorConfigMaps
		err := operator.ConfigMaps.AssignProperties_To_UserAssignedIdentityOperatorConfigMaps(&configMap)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_UserAssignedIdentityOperatorConfigMaps() to populate field ConfigMaps")
		}
		destination.ConfigMaps = &configMap
	} else {
		destination.ConfigMaps = nil
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

type UserAssignedIdentityOperatorConfigMaps struct {
	// ClientId: indicates where the ClientId config map should be placed. If omitted, no config map will be created.
	ClientId *genruntime.ConfigMapDestination `json:"clientId,omitempty"`

	// PrincipalId: indicates where the PrincipalId config map should be placed. If omitted, no config map will be created.
	PrincipalId *genruntime.ConfigMapDestination `json:"principalId,omitempty"`

	// TenantId: indicates where the TenantId config map should be placed. If omitted, no config map will be created.
	TenantId *genruntime.ConfigMapDestination `json:"tenantId,omitempty"`
}

// AssignProperties_From_UserAssignedIdentityOperatorConfigMaps populates our UserAssignedIdentityOperatorConfigMaps from the provided source UserAssignedIdentityOperatorConfigMaps
func (maps *UserAssignedIdentityOperatorConfigMaps) AssignProperties_From_UserAssignedIdentityOperatorConfigMaps(source *storage.UserAssignedIdentityOperatorConfigMaps) error {

	// ClientId
	if source.ClientId != nil {
		clientId := source.ClientId.Copy()
		maps.ClientId = &clientId
	} else {
		maps.ClientId = nil
	}

	// PrincipalId
	if source.PrincipalId != nil {
		principalId := source.PrincipalId.Copy()
		maps.PrincipalId = &principalId
	} else {
		maps.PrincipalId = nil
	}

	// TenantId
	if source.TenantId != nil {
		tenantId := source.TenantId.Copy()
		maps.TenantId = &tenantId
	} else {
		maps.TenantId = nil
	}

	// No error
	return nil
}

// AssignProperties_To_UserAssignedIdentityOperatorConfigMaps populates the provided destination UserAssignedIdentityOperatorConfigMaps from our UserAssignedIdentityOperatorConfigMaps
func (maps *UserAssignedIdentityOperatorConfigMaps) AssignProperties_To_UserAssignedIdentityOperatorConfigMaps(destination *storage.UserAssignedIdentityOperatorConfigMaps) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// ClientId
	if maps.ClientId != nil {
		clientId := maps.ClientId.Copy()
		destination.ClientId = &clientId
	} else {
		destination.ClientId = nil
	}

	// PrincipalId
	if maps.PrincipalId != nil {
		principalId := maps.PrincipalId.Copy()
		destination.PrincipalId = &principalId
	} else {
		destination.PrincipalId = nil
	}

	// TenantId
	if maps.TenantId != nil {
		tenantId := maps.TenantId.Copy()
		destination.TenantId = &tenantId
	} else {
		destination.TenantId = nil
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

func init() {
	SchemeBuilder.Register(&UserAssignedIdentity{}, &UserAssignedIdentityList{})
}
