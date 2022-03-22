// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

import (
	"fmt"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
type ResourceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceGroupSpec   `json:"spec,omitempty"`
	Status            ResourceGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:webhook:path=/mutate-resources-azure-com-v1alpha1api20200601-resourcegroup,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=resources.azure.com,resources=resourcegroups,verbs=create;update,versions=v1alpha1api20200601,name=default.v1alpha1api20200601.resourcegroups.resources.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &ResourceGroup{}

// Default defaults the Azure name of the resource to the Kubernetes name
func (rg *ResourceGroup) Default() {
	if rg.Spec.AzureName == "" {
		rg.Spec.AzureName = rg.Name
	}
}

var _ conditions.Conditioner = &ResourceGroup{}

// GetConditions returns the conditions of the resource
func (rg *ResourceGroup) GetConditions() conditions.Conditions {
	return rg.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (rg *ResourceGroup) SetConditions(conditions conditions.Conditions) {
	rg.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &ResourceGroup{}

// AzureName returns the Azure name of the resource
func (rg *ResourceGroup) AzureName() string {
	return rg.Spec.AzureName
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (rg *ResourceGroup) Owner() *genruntime.ResourceReference {
	return nil
}

// GetSpec returns the specification of this resource
func (rg *ResourceGroup) GetSpec() genruntime.ConvertibleSpec {
	return &rg.Spec
}

// GetStatus returns the status of this resource
func (rg *ResourceGroup) GetStatus() genruntime.ConvertibleStatus {
	return &rg.Status
}

// NewEmptyStatus returns a new empty (blank) status
func (rg *ResourceGroup) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ResourceGroupStatus{}
}

// GetAPIVersion returns the API version of the resource. This is always "2020-06-01"
func (rg *ResourceGroup) GetAPIVersion() string {
	return "2020-06-01"
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Resources/resourceGroups"
func (rg *ResourceGroup) GetType() string { return "Microsoft.Resources/resourceGroups" }

// SetStatus sets the status of this resource
func (rg *ResourceGroup) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ResourceGroupStatus); ok {
		rg.Status = *st
		return nil
	}

	// Convert status to required version
	var st ResourceGroupStatus
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rg.Status = st
	return nil
}

// GetResourceKind returns the kind of the resource
func (rg *ResourceGroup) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

var _ genruntime.LocatableResource = &ResourceGroup{}

func (rg *ResourceGroup) Location() string {
	if rg.Spec.Location == nil {
		return ""
	}
	return *rg.Spec.Location
}

var _ conversion.Convertible = &ResourceGroup{}

// ConvertFrom populates our ResourceGroup from the provided hub ResourceGroup
func (rg *ResourceGroup) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1beta20200601.ResourceGroup)
	if !ok {
		return fmt.Errorf("expected compute/v1beta20200601/ResourceGroup but received %T instead", hub)
	}

	return rg.AssignPropertiesFromResourceGroup(source)
}

// ConvertTo populates the provided hub Disk from our Disk
func (rg *ResourceGroup) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1beta20200601.ResourceGroup)
	if !ok {
		return fmt.Errorf("expected compute/v1beta20200601/ResourceGroup but received %T instead", hub)
	}

	return rg.AssignPropertiesToResourceGroup(destination)
}

// AssignPropertiesFromResourceGroup populates our ResourceGroup from the provided source ResourceGroup
func (rg *ResourceGroup) AssignPropertiesFromResourceGroup(source *v1beta20200601.ResourceGroup) error {

	// ObjectMeta
	rg.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec ResourceGroupSpec
	err := spec.ConvertSpecFrom(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromResourceGroupSpec() to populate field Spec")
	}
	rg.Spec = spec

	// Status
	var status ResourceGroupStatus
	err = status.ConvertStatusFrom(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromResourceGroupStatus() to populate field Status")
	}
	rg.Status = status

	// No error
	return nil
}

// AssignPropertiesToResourceGroup populates the provided destination Disk from our Disk
func (rg *ResourceGroup) AssignPropertiesToResourceGroup(destination *v1beta20200601.ResourceGroup) error {

	// ObjectMeta
	destination.ObjectMeta = *rg.ObjectMeta.DeepCopy()

	// Spec
	var spec v1beta20200601.ResourceGroupSpec
	err := rg.Spec.ConvertSpecTo(&spec)
	if err != nil {
		return errors.Wrap(err, "calling ConvertSpecTo() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1beta20200601.ResourceGroupStatus
	err = rg.Status.ConvertStatusTo(&status)
	if err != nil {
		return errors.Wrap(err, "calling ConvertStatusTo() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// +kubebuilder:object:root=true
type ResourceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceGroup `json:"items"`
}

type ResourceGroupStatus struct {
	ID *string `json:"id,omitempty"`

	Name     *string `json:"name,omitempty"`
	Location *string `json:"location,omitempty"`

	// ManagedBy is the management group responsible for managing this group
	ManagedBy *string `json:"managedBy,omitempty"`

	// Tags are user defined key value pairs
	Tags map[string]string `json:"tags,omitempty"`

	ProvisioningState *string `json:"provisioningState,omitempty"`

	// Conditions describe the observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`
}

var (
	_ genruntime.FromARMConverter  = &ResourceGroupStatus{}
	_ genruntime.ConvertibleStatus = &ResourceGroupStatus{}
)

func (status *ResourceGroupStatus) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &ResourceGroupStatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (status *ResourceGroupStatus) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(ResourceGroupStatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromArm() function. Expected ResourceGroupStatusArm, got %T", armInput)
	}
	status.ID = typedInput.ID
	status.Location = typedInput.Location
	status.ManagedBy = typedInput.ManagedBy
	status.Name = typedInput.Name
	status.Tags = typedInput.Tags
	// Set property ‘AccessTier’:
	// copying flattened property:
	if typedInput.Properties != nil {
		status.ProvisioningState = typedInput.Properties.ProvisioningState
	}

	return nil
}

// ConvertStatusTo copies information from the current instance onto the supplied destination
func (status *ResourceGroupStatus) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1beta20200601.ResourceGroupStatus)
	if !ok {
		return errors.Errorf(
			"cannot convert ResourceGroupStatus, expected destination to be a *v1beta20200601.ResourceGroupStatus but received %T",
			destination)
	}

	dst.ID = status.ID
	dst.Name = status.Name
	dst.Location = status.Location
	dst.ManagedBy = status.ManagedBy
	dst.ProvisioningState = status.ProvisioningState

	dst.Tags = make(map[string]string)
	for k, v := range status.Tags {
		dst.Tags[k] = v
	}

	for _, c := range status.Conditions {
		dst.Conditions = append(dst.Conditions, c.Copy())
	}

	return nil
}

// ConvertStatusFrom copies information from the supplied source onto the current instance
func (status *ResourceGroupStatus) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1beta20200601.ResourceGroupStatus)
	if !ok {
		return errors.Errorf(
			"cannot convert ResourceGroupStatus, expected source to be a *v1beta20200601.ResourceGroupStatus but received %T",
			source)
	}

	status.ID = src.ID
	status.Name = src.Name
	status.Location = src.Location
	status.ManagedBy = src.ManagedBy
	status.ProvisioningState = src.ProvisioningState

	status.Tags = make(map[string]string)
	for k, v := range src.Tags {
		status.Tags[k] = v
	}

	for _, c := range src.Conditions {
		status.Conditions = append(status.Conditions, c.Copy())
	}

	return nil
}

type ResourceGroupSpec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name
	// of the resource in Kubernetes but it doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:Required
	// Location is the Azure location for the group (eg westus2, southcentralus, etc...)
	Location *string `json:"location,omitempty"`

	// ManagedBy is the management group responsible for managing this group
	ManagedBy *string `json:"managedBy,omitempty"`

	// Tags are user defined key value pairs
	Tags map[string]string `json:"tags,omitempty"`
}

var (
	_ genruntime.ARMTransformer  = &ResourceGroupSpec{}
	_ genruntime.ConvertibleSpec = &ResourceGroupSpec{}
)

func (spec *ResourceGroupSpec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return ResourceGroupSpecARM{}
}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (spec *ResourceGroupSpec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if spec == nil {
		return nil, nil
	}
	result := ResourceGroupSpecARM{}
	result.Location = spec.Location
	result.Name = resolved.Name
	result.ManagedBy = spec.ManagedBy
	result.Tags = spec.Tags
	return result, nil
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (spec *ResourceGroupSpec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(ResourceGroupSpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromArm() function. Expected ResourceGroupSpecArm, got %T", armInput)
	}
	// spec.ApiVersion = typedInput.ApiVersion
	spec.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))
	spec.Location = typedInput.Location
	spec.ManagedBy = typedInput.ManagedBy
	spec.Tags = typedInput.Tags
	return nil
}

// SetAzureName sets the Azure name of the resource
func (spec *ResourceGroupSpec) SetAzureName(azureName string) { spec.AzureName = azureName }

// ConvertSpecTo copies information from the current instance over to the supplied destination
func (spec *ResourceGroupSpec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1beta20200601.ResourceGroupSpec)
	if !ok {
		return errors.Errorf(
			"cannot convert ResourceGroupSpec, expected destination to be a *v1beta20200601.ResouceGroupSpec but received %T",
			destination)
	}

	dst.AzureName = spec.AzureName
	dst.Location = spec.Location
	dst.ManagedBy = spec.ManagedBy

	dst.Tags = make(map[string]string)
	for k, v := range spec.Tags {
		dst.Tags[k] = v
	}

	return nil
}

// ConvertSpecFrom copies information from the supplied source onto the current instance
func (spec *ResourceGroupSpec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1beta20200601.ResourceGroupSpec)
	if !ok {
		return errors.Errorf(
			"cannot convert ResourceGroupSpec, expected source to be a *v1beta20200601.ResourceGroupSpec but received %T",
			source)
	}

	spec.AzureName = src.AzureName
	spec.Location = src.Location
	spec.ManagedBy = src.ManagedBy

	spec.Tags = make(map[string]string)
	for k, v := range src.Tags {
		spec.Tags[k] = v
	}

	return nil
}

func init() {
	SchemeBuilder.Register(&ResourceGroup{}, &ResourceGroupList{})
}
