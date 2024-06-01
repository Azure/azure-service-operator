// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20180901

import (
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/network/v1api20180901/storage"
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
// - Generated from: /privatedns/resource-manager/Microsoft.Network/stable/2018-09-01/privatedns.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}
type PrivateDnsZone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateDnsZone_Spec   `json:"spec,omitempty"`
	Status            PrivateDnsZone_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &PrivateDnsZone{}

// GetConditions returns the conditions of the resource
func (zone *PrivateDnsZone) GetConditions() conditions.Conditions {
	return zone.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (zone *PrivateDnsZone) SetConditions(conditions conditions.Conditions) {
	zone.Status.Conditions = conditions
}

var _ conversion.Convertible = &PrivateDnsZone{}

// ConvertFrom populates our PrivateDnsZone from the provided hub PrivateDnsZone
func (zone *PrivateDnsZone) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.PrivateDnsZone)
	if !ok {
		return fmt.Errorf("expected network/v1api20180901/storage/PrivateDnsZone but received %T instead", hub)
	}

	return zone.AssignProperties_From_PrivateDnsZone(source)
}

// ConvertTo populates the provided hub PrivateDnsZone from our PrivateDnsZone
func (zone *PrivateDnsZone) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.PrivateDnsZone)
	if !ok {
		return fmt.Errorf("expected network/v1api20180901/storage/PrivateDnsZone but received %T instead", hub)
	}

	return zone.AssignProperties_To_PrivateDnsZone(destination)
}

// +kubebuilder:webhook:path=/mutate-network-azure-com-v1api20180901-privatednszone,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=privatednszones,verbs=create;update,versions=v1api20180901,name=default.v1api20180901.privatednszones.network.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &PrivateDnsZone{}

// Default applies defaults to the PrivateDnsZone resource
func (zone *PrivateDnsZone) Default() {
	zone.defaultImpl()
	var temp any = zone
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (zone *PrivateDnsZone) defaultAzureName() {
	if zone.Spec.AzureName == "" {
		zone.Spec.AzureName = zone.Name
	}
}

// defaultImpl applies the code generated defaults to the PrivateDnsZone resource
func (zone *PrivateDnsZone) defaultImpl() { zone.defaultAzureName() }

var _ genruntime.ImportableResource = &PrivateDnsZone{}

// InitializeSpec initializes the spec for this resource from the given status
func (zone *PrivateDnsZone) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*PrivateDnsZone_STATUS); ok {
		return zone.Spec.Initialize_From_PrivateDnsZone_STATUS(s)
	}

	return fmt.Errorf("expected Status of type PrivateDnsZone_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &PrivateDnsZone{}

// AzureName returns the Azure name of the resource
func (zone *PrivateDnsZone) AzureName() string {
	return zone.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-09-01"
func (zone PrivateDnsZone) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (zone *PrivateDnsZone) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (zone *PrivateDnsZone) GetSpec() genruntime.ConvertibleSpec {
	return &zone.Spec
}

// GetStatus returns the status of this resource
func (zone *PrivateDnsZone) GetStatus() genruntime.ConvertibleStatus {
	return &zone.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (zone *PrivateDnsZone) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/privateDnsZones"
func (zone *PrivateDnsZone) GetType() string {
	return "Microsoft.Network/privateDnsZones"
}

// NewEmptyStatus returns a new empty (blank) status
func (zone *PrivateDnsZone) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &PrivateDnsZone_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (zone *PrivateDnsZone) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(zone.Spec)
	return zone.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (zone *PrivateDnsZone) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*PrivateDnsZone_STATUS); ok {
		zone.Status = *st
		return nil
	}

	// Convert status to required version
	var st PrivateDnsZone_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	zone.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-network-azure-com-v1api20180901-privatednszone,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=privatednszones,verbs=create;update,versions=v1api20180901,name=validate.v1api20180901.privatednszones.network.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &PrivateDnsZone{}

// ValidateCreate validates the creation of the resource
func (zone *PrivateDnsZone) ValidateCreate() (admission.Warnings, error) {
	validations := zone.createValidations()
	var temp any = zone
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (zone *PrivateDnsZone) ValidateDelete() (admission.Warnings, error) {
	validations := zone.deleteValidations()
	var temp any = zone
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (zone *PrivateDnsZone) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := zone.updateValidations()
	var temp any = zone
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (zone *PrivateDnsZone) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){zone.validateResourceReferences, zone.validateOwnerReference}
}

// deleteValidations validates the deletion of the resource
func (zone *PrivateDnsZone) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (zone *PrivateDnsZone) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return zone.validateResourceReferences()
		},
		zone.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return zone.validateOwnerReference()
		},
	}
}

// validateOwnerReference validates the owner field
func (zone *PrivateDnsZone) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(zone)
}

// validateResourceReferences validates all resource references
func (zone *PrivateDnsZone) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&zone.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (zone *PrivateDnsZone) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*PrivateDnsZone)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, zone)
}

// AssignProperties_From_PrivateDnsZone populates our PrivateDnsZone from the provided source PrivateDnsZone
func (zone *PrivateDnsZone) AssignProperties_From_PrivateDnsZone(source *storage.PrivateDnsZone) error {

	// ObjectMeta
	zone.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec PrivateDnsZone_Spec
	err := spec.AssignProperties_From_PrivateDnsZone_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_PrivateDnsZone_Spec() to populate field Spec")
	}
	zone.Spec = spec

	// Status
	var status PrivateDnsZone_STATUS
	err = status.AssignProperties_From_PrivateDnsZone_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_PrivateDnsZone_STATUS() to populate field Status")
	}
	zone.Status = status

	// No error
	return nil
}

// AssignProperties_To_PrivateDnsZone populates the provided destination PrivateDnsZone from our PrivateDnsZone
func (zone *PrivateDnsZone) AssignProperties_To_PrivateDnsZone(destination *storage.PrivateDnsZone) error {

	// ObjectMeta
	destination.ObjectMeta = *zone.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.PrivateDnsZone_Spec
	err := zone.Spec.AssignProperties_To_PrivateDnsZone_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_PrivateDnsZone_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.PrivateDnsZone_STATUS
	err = zone.Status.AssignProperties_To_PrivateDnsZone_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_PrivateDnsZone_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (zone *PrivateDnsZone) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: zone.Spec.OriginalVersion(),
		Kind:    "PrivateDnsZone",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /privatedns/resource-manager/Microsoft.Network/stable/2018-09-01/privatedns.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}
type PrivateDnsZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateDnsZone `json:"items"`
}

// +kubebuilder:validation:Enum={"2018-09-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2018-09-01")

type PrivateDnsZone_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// Etag: The ETag of the zone.
	Etag *string `json:"etag,omitempty"`

	// Location: The Azure Region where the resource lives
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &PrivateDnsZone_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (zone *PrivateDnsZone_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if zone == nil {
		return nil, nil
	}
	result := &PrivateDnsZone_Spec_ARM{}

	// Set property "Etag":
	if zone.Etag != nil {
		etag := *zone.Etag
		result.Etag = &etag
	}

	// Set property "Location":
	if zone.Location != nil {
		location := *zone.Location
		result.Location = &location
	}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Tags":
	if zone.Tags != nil {
		result.Tags = make(map[string]string, len(zone.Tags))
		for key, value := range zone.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (zone *PrivateDnsZone_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &PrivateDnsZone_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (zone *PrivateDnsZone_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(PrivateDnsZone_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected PrivateDnsZone_Spec_ARM, got %T", armInput)
	}

	// Set property "AzureName":
	zone.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property "Etag":
	if typedInput.Etag != nil {
		etag := *typedInput.Etag
		zone.Etag = &etag
	}

	// Set property "Location":
	if typedInput.Location != nil {
		location := *typedInput.Location
		zone.Location = &location
	}

	// Set property "Owner":
	zone.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "Tags":
	if typedInput.Tags != nil {
		zone.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			zone.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &PrivateDnsZone_Spec{}

// ConvertSpecFrom populates our PrivateDnsZone_Spec from the provided source
func (zone *PrivateDnsZone_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.PrivateDnsZone_Spec)
	if ok {
		// Populate our instance from source
		return zone.AssignProperties_From_PrivateDnsZone_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.PrivateDnsZone_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = zone.AssignProperties_From_PrivateDnsZone_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our PrivateDnsZone_Spec
func (zone *PrivateDnsZone_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.PrivateDnsZone_Spec)
	if ok {
		// Populate destination from our instance
		return zone.AssignProperties_To_PrivateDnsZone_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.PrivateDnsZone_Spec{}
	err := zone.AssignProperties_To_PrivateDnsZone_Spec(dst)
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

// AssignProperties_From_PrivateDnsZone_Spec populates our PrivateDnsZone_Spec from the provided source PrivateDnsZone_Spec
func (zone *PrivateDnsZone_Spec) AssignProperties_From_PrivateDnsZone_Spec(source *storage.PrivateDnsZone_Spec) error {

	// AzureName
	zone.AzureName = source.AzureName

	// Etag
	zone.Etag = genruntime.ClonePointerToString(source.Etag)

	// Location
	zone.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		zone.Owner = &owner
	} else {
		zone.Owner = nil
	}

	// Tags
	zone.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignProperties_To_PrivateDnsZone_Spec populates the provided destination PrivateDnsZone_Spec from our PrivateDnsZone_Spec
func (zone *PrivateDnsZone_Spec) AssignProperties_To_PrivateDnsZone_Spec(destination *storage.PrivateDnsZone_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = zone.AzureName

	// Etag
	destination.Etag = genruntime.ClonePointerToString(zone.Etag)

	// Location
	destination.Location = genruntime.ClonePointerToString(zone.Location)

	// OriginalVersion
	destination.OriginalVersion = zone.OriginalVersion()

	// Owner
	if zone.Owner != nil {
		owner := zone.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(zone.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Initialize_From_PrivateDnsZone_STATUS populates our PrivateDnsZone_Spec from the provided source PrivateDnsZone_STATUS
func (zone *PrivateDnsZone_Spec) Initialize_From_PrivateDnsZone_STATUS(source *PrivateDnsZone_STATUS) error {

	// Etag
	zone.Etag = genruntime.ClonePointerToString(source.Etag)

	// Location
	zone.Location = genruntime.ClonePointerToString(source.Location)

	// Tags
	zone.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (zone *PrivateDnsZone_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (zone *PrivateDnsZone_Spec) SetAzureName(azureName string) { zone.AzureName = azureName }

type PrivateDnsZone_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Etag: The ETag of the zone.
	Etag *string `json:"etag,omitempty"`

	// Id: Fully qualified resource Id for the resource. Example -
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateDnsZoneName}'.
	Id *string `json:"id,omitempty"`

	// Location: The Azure Region where the resource lives
	Location *string `json:"location,omitempty"`

	// MaxNumberOfRecordSets: The maximum number of record sets that can be created in this Private DNS zone. This is a
	// read-only property and any attempt to set this value will be ignored.
	MaxNumberOfRecordSets *int `json:"maxNumberOfRecordSets,omitempty"`

	// MaxNumberOfVirtualNetworkLinks: The maximum number of virtual networks that can be linked to this Private DNS zone. This
	// is a read-only property and any attempt to set this value will be ignored.
	MaxNumberOfVirtualNetworkLinks *int `json:"maxNumberOfVirtualNetworkLinks,omitempty"`

	// MaxNumberOfVirtualNetworkLinksWithRegistration: The maximum number of virtual networks that can be linked to this
	// Private DNS zone with registration enabled. This is a read-only property and any attempt to set this value will be
	// ignored.
	MaxNumberOfVirtualNetworkLinksWithRegistration *int `json:"maxNumberOfVirtualNetworkLinksWithRegistration,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// NumberOfRecordSets: The current number of record sets in this Private DNS zone. This is a read-only property and any
	// attempt to set this value will be ignored.
	NumberOfRecordSets *int `json:"numberOfRecordSets,omitempty"`

	// NumberOfVirtualNetworkLinks: The current number of virtual networks that are linked to this Private DNS zone. This is a
	// read-only property and any attempt to set this value will be ignored.
	NumberOfVirtualNetworkLinks *int `json:"numberOfVirtualNetworkLinks,omitempty"`

	// NumberOfVirtualNetworkLinksWithRegistration: The current number of virtual networks that are linked to this Private DNS
	// zone with registration enabled. This is a read-only property and any attempt to set this value will be ignored.
	NumberOfVirtualNetworkLinksWithRegistration *int `json:"numberOfVirtualNetworkLinksWithRegistration,omitempty"`

	// ProvisioningState: The provisioning state of the resource. This is a read-only property and any attempt to set this
	// value will be ignored.
	ProvisioningState *PrivateZoneProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. Example - 'Microsoft.Network/privateDnsZones'.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &PrivateDnsZone_STATUS{}

// ConvertStatusFrom populates our PrivateDnsZone_STATUS from the provided source
func (zone *PrivateDnsZone_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.PrivateDnsZone_STATUS)
	if ok {
		// Populate our instance from source
		return zone.AssignProperties_From_PrivateDnsZone_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.PrivateDnsZone_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = zone.AssignProperties_From_PrivateDnsZone_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our PrivateDnsZone_STATUS
func (zone *PrivateDnsZone_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.PrivateDnsZone_STATUS)
	if ok {
		// Populate destination from our instance
		return zone.AssignProperties_To_PrivateDnsZone_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.PrivateDnsZone_STATUS{}
	err := zone.AssignProperties_To_PrivateDnsZone_STATUS(dst)
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

var _ genruntime.FromARMConverter = &PrivateDnsZone_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (zone *PrivateDnsZone_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &PrivateDnsZone_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (zone *PrivateDnsZone_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(PrivateDnsZone_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected PrivateDnsZone_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "Etag":
	if typedInput.Etag != nil {
		etag := *typedInput.Etag
		zone.Etag = &etag
	}

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		zone.Id = &id
	}

	// Set property "Location":
	if typedInput.Location != nil {
		location := *typedInput.Location
		zone.Location = &location
	}

	// Set property "MaxNumberOfRecordSets":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.MaxNumberOfRecordSets != nil {
			maxNumberOfRecordSets := *typedInput.Properties.MaxNumberOfRecordSets
			zone.MaxNumberOfRecordSets = &maxNumberOfRecordSets
		}
	}

	// Set property "MaxNumberOfVirtualNetworkLinks":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.MaxNumberOfVirtualNetworkLinks != nil {
			maxNumberOfVirtualNetworkLinks := *typedInput.Properties.MaxNumberOfVirtualNetworkLinks
			zone.MaxNumberOfVirtualNetworkLinks = &maxNumberOfVirtualNetworkLinks
		}
	}

	// Set property "MaxNumberOfVirtualNetworkLinksWithRegistration":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.MaxNumberOfVirtualNetworkLinksWithRegistration != nil {
			maxNumberOfVirtualNetworkLinksWithRegistration := *typedInput.Properties.MaxNumberOfVirtualNetworkLinksWithRegistration
			zone.MaxNumberOfVirtualNetworkLinksWithRegistration = &maxNumberOfVirtualNetworkLinksWithRegistration
		}
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		zone.Name = &name
	}

	// Set property "NumberOfRecordSets":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.NumberOfRecordSets != nil {
			numberOfRecordSets := *typedInput.Properties.NumberOfRecordSets
			zone.NumberOfRecordSets = &numberOfRecordSets
		}
	}

	// Set property "NumberOfVirtualNetworkLinks":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.NumberOfVirtualNetworkLinks != nil {
			numberOfVirtualNetworkLinks := *typedInput.Properties.NumberOfVirtualNetworkLinks
			zone.NumberOfVirtualNetworkLinks = &numberOfVirtualNetworkLinks
		}
	}

	// Set property "NumberOfVirtualNetworkLinksWithRegistration":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.NumberOfVirtualNetworkLinksWithRegistration != nil {
			numberOfVirtualNetworkLinksWithRegistration := *typedInput.Properties.NumberOfVirtualNetworkLinksWithRegistration
			zone.NumberOfVirtualNetworkLinksWithRegistration = &numberOfVirtualNetworkLinksWithRegistration
		}
	}

	// Set property "ProvisioningState":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ProvisioningState != nil {
			provisioningState := *typedInput.Properties.ProvisioningState
			zone.ProvisioningState = &provisioningState
		}
	}

	// Set property "Tags":
	if typedInput.Tags != nil {
		zone.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			zone.Tags[key] = value
		}
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		zone.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_PrivateDnsZone_STATUS populates our PrivateDnsZone_STATUS from the provided source PrivateDnsZone_STATUS
func (zone *PrivateDnsZone_STATUS) AssignProperties_From_PrivateDnsZone_STATUS(source *storage.PrivateDnsZone_STATUS) error {

	// Conditions
	zone.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Etag
	zone.Etag = genruntime.ClonePointerToString(source.Etag)

	// Id
	zone.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	zone.Location = genruntime.ClonePointerToString(source.Location)

	// MaxNumberOfRecordSets
	zone.MaxNumberOfRecordSets = genruntime.ClonePointerToInt(source.MaxNumberOfRecordSets)

	// MaxNumberOfVirtualNetworkLinks
	zone.MaxNumberOfVirtualNetworkLinks = genruntime.ClonePointerToInt(source.MaxNumberOfVirtualNetworkLinks)

	// MaxNumberOfVirtualNetworkLinksWithRegistration
	zone.MaxNumberOfVirtualNetworkLinksWithRegistration = genruntime.ClonePointerToInt(source.MaxNumberOfVirtualNetworkLinksWithRegistration)

	// Name
	zone.Name = genruntime.ClonePointerToString(source.Name)

	// NumberOfRecordSets
	zone.NumberOfRecordSets = genruntime.ClonePointerToInt(source.NumberOfRecordSets)

	// NumberOfVirtualNetworkLinks
	zone.NumberOfVirtualNetworkLinks = genruntime.ClonePointerToInt(source.NumberOfVirtualNetworkLinks)

	// NumberOfVirtualNetworkLinksWithRegistration
	zone.NumberOfVirtualNetworkLinksWithRegistration = genruntime.ClonePointerToInt(source.NumberOfVirtualNetworkLinksWithRegistration)

	// ProvisioningState
	if source.ProvisioningState != nil {
		provisioningState := *source.ProvisioningState
		provisioningStateTemp := genruntime.ToEnum(provisioningState, privateZoneProperties_ProvisioningState_STATUS_Values)
		zone.ProvisioningState = &provisioningStateTemp
	} else {
		zone.ProvisioningState = nil
	}

	// Tags
	zone.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	zone.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_PrivateDnsZone_STATUS populates the provided destination PrivateDnsZone_STATUS from our PrivateDnsZone_STATUS
func (zone *PrivateDnsZone_STATUS) AssignProperties_To_PrivateDnsZone_STATUS(destination *storage.PrivateDnsZone_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(zone.Conditions)

	// Etag
	destination.Etag = genruntime.ClonePointerToString(zone.Etag)

	// Id
	destination.Id = genruntime.ClonePointerToString(zone.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(zone.Location)

	// MaxNumberOfRecordSets
	destination.MaxNumberOfRecordSets = genruntime.ClonePointerToInt(zone.MaxNumberOfRecordSets)

	// MaxNumberOfVirtualNetworkLinks
	destination.MaxNumberOfVirtualNetworkLinks = genruntime.ClonePointerToInt(zone.MaxNumberOfVirtualNetworkLinks)

	// MaxNumberOfVirtualNetworkLinksWithRegistration
	destination.MaxNumberOfVirtualNetworkLinksWithRegistration = genruntime.ClonePointerToInt(zone.MaxNumberOfVirtualNetworkLinksWithRegistration)

	// Name
	destination.Name = genruntime.ClonePointerToString(zone.Name)

	// NumberOfRecordSets
	destination.NumberOfRecordSets = genruntime.ClonePointerToInt(zone.NumberOfRecordSets)

	// NumberOfVirtualNetworkLinks
	destination.NumberOfVirtualNetworkLinks = genruntime.ClonePointerToInt(zone.NumberOfVirtualNetworkLinks)

	// NumberOfVirtualNetworkLinksWithRegistration
	destination.NumberOfVirtualNetworkLinksWithRegistration = genruntime.ClonePointerToInt(zone.NumberOfVirtualNetworkLinksWithRegistration)

	// ProvisioningState
	if zone.ProvisioningState != nil {
		provisioningState := string(*zone.ProvisioningState)
		destination.ProvisioningState = &provisioningState
	} else {
		destination.ProvisioningState = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(zone.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(zone.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type PrivateZoneProperties_ProvisioningState_STATUS string

const (
	PrivateZoneProperties_ProvisioningState_STATUS_Canceled  = PrivateZoneProperties_ProvisioningState_STATUS("Canceled")
	PrivateZoneProperties_ProvisioningState_STATUS_Creating  = PrivateZoneProperties_ProvisioningState_STATUS("Creating")
	PrivateZoneProperties_ProvisioningState_STATUS_Deleting  = PrivateZoneProperties_ProvisioningState_STATUS("Deleting")
	PrivateZoneProperties_ProvisioningState_STATUS_Failed    = PrivateZoneProperties_ProvisioningState_STATUS("Failed")
	PrivateZoneProperties_ProvisioningState_STATUS_Succeeded = PrivateZoneProperties_ProvisioningState_STATUS("Succeeded")
	PrivateZoneProperties_ProvisioningState_STATUS_Updating  = PrivateZoneProperties_ProvisioningState_STATUS("Updating")
)

// Mapping from string to PrivateZoneProperties_ProvisioningState_STATUS
var privateZoneProperties_ProvisioningState_STATUS_Values = map[string]PrivateZoneProperties_ProvisioningState_STATUS{
	"canceled":  PrivateZoneProperties_ProvisioningState_STATUS_Canceled,
	"creating":  PrivateZoneProperties_ProvisioningState_STATUS_Creating,
	"deleting":  PrivateZoneProperties_ProvisioningState_STATUS_Deleting,
	"failed":    PrivateZoneProperties_ProvisioningState_STATUS_Failed,
	"succeeded": PrivateZoneProperties_ProvisioningState_STATUS_Succeeded,
	"updating":  PrivateZoneProperties_ProvisioningState_STATUS_Updating,
}

func init() {
	SchemeBuilder.Register(&PrivateDnsZone{}, &PrivateDnsZoneList{})
}
