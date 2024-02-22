// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=compute.azure.com,resources=virtualmachinesextensions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=compute.azure.com,resources={virtualmachinesextensions/status,virtualmachinesextensions/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220301.VirtualMachinesExtension
// Generator information:
// - Generated from: /compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/virtualMachine.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{vmExtensionName}
type VirtualMachinesExtension struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualMachines_Extension_Spec   `json:"spec,omitempty"`
	Status            VirtualMachines_Extension_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &VirtualMachinesExtension{}

// GetConditions returns the conditions of the resource
func (extension *VirtualMachinesExtension) GetConditions() conditions.Conditions {
	return extension.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (extension *VirtualMachinesExtension) SetConditions(conditions conditions.Conditions) {
	extension.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &VirtualMachinesExtension{}

// AzureName returns the Azure name of the resource
func (extension *VirtualMachinesExtension) AzureName() string {
	return extension.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-03-01"
func (extension VirtualMachinesExtension) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (extension *VirtualMachinesExtension) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (extension *VirtualMachinesExtension) GetSpec() genruntime.ConvertibleSpec {
	return &extension.Spec
}

// GetStatus returns the status of this resource
func (extension *VirtualMachinesExtension) GetStatus() genruntime.ConvertibleStatus {
	return &extension.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (extension *VirtualMachinesExtension) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Compute/virtualMachines/extensions"
func (extension *VirtualMachinesExtension) GetType() string {
	return "Microsoft.Compute/virtualMachines/extensions"
}

// NewEmptyStatus returns a new empty (blank) status
func (extension *VirtualMachinesExtension) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &VirtualMachines_Extension_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (extension *VirtualMachinesExtension) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(extension.Spec)
	return extension.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (extension *VirtualMachinesExtension) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*VirtualMachines_Extension_STATUS); ok {
		extension.Status = *st
		return nil
	}

	// Convert status to required version
	var st VirtualMachines_Extension_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	extension.Status = st
	return nil
}

// Hub marks that this VirtualMachinesExtension is the hub type for conversion
func (extension *VirtualMachinesExtension) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (extension *VirtualMachinesExtension) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: extension.Spec.OriginalVersion,
		Kind:    "VirtualMachinesExtension",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220301.VirtualMachinesExtension
// Generator information:
// - Generated from: /compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/virtualMachine.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}/extensions/{vmExtensionName}
type VirtualMachinesExtensionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualMachinesExtension `json:"items"`
}

// Storage version of v1api20220301.VirtualMachines_Extension_Spec
type VirtualMachines_Extension_Spec struct {
	AutoUpgradeMinorVersion *bool `json:"autoUpgradeMinorVersion,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName              string                               `json:"azureName,omitempty"`
	EnableAutomaticUpgrade *bool                                `json:"enableAutomaticUpgrade,omitempty"`
	ForceUpdateTag         *string                              `json:"forceUpdateTag,omitempty"`
	InstanceView           *VirtualMachineExtensionInstanceView `json:"instanceView,omitempty"`
	Location               *string                              `json:"location,omitempty"`
	OriginalVersion        string                               `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a compute.azure.com/VirtualMachine resource
	Owner                         *genruntime.KnownResourceReference `group:"compute.azure.com" json:"owner,omitempty" kind:"VirtualMachine"`
	PropertyBag                   genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	ProtectedSettings             *genruntime.SecretMapReference     `json:"protectedSettings,omitempty"`
	ProtectedSettingsFromKeyVault *KeyVaultSecretReference           `json:"protectedSettingsFromKeyVault,omitempty"`
	Publisher                     *string                            `json:"publisher,omitempty"`
	Settings                      map[string]v1.JSON                 `json:"settings,omitempty"`
	SuppressFailures              *bool                              `json:"suppressFailures,omitempty"`
	Tags                          map[string]string                  `json:"tags,omitempty"`
	Type                          *string                            `json:"type,omitempty"`
	TypeHandlerVersion            *string                            `json:"typeHandlerVersion,omitempty"`
}

var _ genruntime.ConvertibleSpec = &VirtualMachines_Extension_Spec{}

// ConvertSpecFrom populates our VirtualMachines_Extension_Spec from the provided source
func (extension *VirtualMachines_Extension_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == extension {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(extension)
}

// ConvertSpecTo populates the provided destination from our VirtualMachines_Extension_Spec
func (extension *VirtualMachines_Extension_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == extension {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(extension)
}

// Storage version of v1api20220301.VirtualMachines_Extension_STATUS
type VirtualMachines_Extension_STATUS struct {
	AutoUpgradeMinorVersion       *bool                                       `json:"autoUpgradeMinorVersion,omitempty"`
	Conditions                    []conditions.Condition                      `json:"conditions,omitempty"`
	EnableAutomaticUpgrade        *bool                                       `json:"enableAutomaticUpgrade,omitempty"`
	ForceUpdateTag                *string                                     `json:"forceUpdateTag,omitempty"`
	Id                            *string                                     `json:"id,omitempty"`
	InstanceView                  *VirtualMachineExtensionInstanceView_STATUS `json:"instanceView,omitempty"`
	Location                      *string                                     `json:"location,omitempty"`
	Name                          *string                                     `json:"name,omitempty"`
	PropertiesType                *string                                     `json:"properties_type,omitempty"`
	PropertyBag                   genruntime.PropertyBag                      `json:"$propertyBag,omitempty"`
	ProtectedSettingsFromKeyVault *KeyVaultSecretReference_STATUS             `json:"protectedSettingsFromKeyVault,omitempty"`
	ProvisioningState             *string                                     `json:"provisioningState,omitempty"`
	Publisher                     *string                                     `json:"publisher,omitempty"`
	Settings                      map[string]v1.JSON                          `json:"settings,omitempty"`
	SuppressFailures              *bool                                       `json:"suppressFailures,omitempty"`
	Tags                          map[string]string                           `json:"tags,omitempty"`
	Type                          *string                                     `json:"type,omitempty"`
	TypeHandlerVersion            *string                                     `json:"typeHandlerVersion,omitempty"`
}

var _ genruntime.ConvertibleStatus = &VirtualMachines_Extension_STATUS{}

// ConvertStatusFrom populates our VirtualMachines_Extension_STATUS from the provided source
func (extension *VirtualMachines_Extension_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == extension {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(extension)
}

// ConvertStatusTo populates the provided destination from our VirtualMachines_Extension_STATUS
func (extension *VirtualMachines_Extension_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == extension {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(extension)
}

// Storage version of v1api20220301.VirtualMachineExtensionInstanceView
// The instance view of a virtual machine extension.
type VirtualMachineExtensionInstanceView struct {
	Name               *string                `json:"name,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Statuses           []InstanceViewStatus   `json:"statuses,omitempty"`
	Substatuses        []InstanceViewStatus   `json:"substatuses,omitempty"`
	Type               *string                `json:"type,omitempty"`
	TypeHandlerVersion *string                `json:"typeHandlerVersion,omitempty"`
}

// Storage version of v1api20220301.VirtualMachineExtensionInstanceView_STATUS
// The instance view of a virtual machine extension.
type VirtualMachineExtensionInstanceView_STATUS struct {
	Name               *string                     `json:"name,omitempty"`
	PropertyBag        genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	Statuses           []InstanceViewStatus_STATUS `json:"statuses,omitempty"`
	Substatuses        []InstanceViewStatus_STATUS `json:"substatuses,omitempty"`
	Type               *string                     `json:"type,omitempty"`
	TypeHandlerVersion *string                     `json:"typeHandlerVersion,omitempty"`
}

// Storage version of v1api20220301.InstanceViewStatus
// Instance view status.
type InstanceViewStatus struct {
	Code          *string                `json:"code,omitempty"`
	DisplayStatus *string                `json:"displayStatus,omitempty"`
	Level         *string                `json:"level,omitempty"`
	Message       *string                `json:"message,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Time          *string                `json:"time,omitempty"`
}

func init() {
	SchemeBuilder.Register(&VirtualMachinesExtension{}, &VirtualMachinesExtensionList{})
}
