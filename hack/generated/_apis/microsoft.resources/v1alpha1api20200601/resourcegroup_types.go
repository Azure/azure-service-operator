// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/conditions"
)

// TODO: it doesn't really matter where these are (as long as they're in _apis, where is where we run controller-gen).
// These are the permissions required by the generic_controller. They're here because they can't go outside the _apis
// directory.

// +kubebuilder:rbac:groups=core,resources=events,verbs=get;list;watch;create;patch
// +kubebuilder:rbac:groups=coordination.k8s.io,resources=leases,verbs=get;list;watch;create;update;patch;delete

// +kubebuilder:rbac:groups=microsoft.resources.infra.azure.com,resources=resourcegroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=microsoft.resources.infra.azure.com,resources={resourcegroups/status,resourcegroups/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
type ResourceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceGroupSpec   `json:"spec,omitempty"`
	Status            ResourceGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:webhook:path=/mutate-microsoft-resources-infra-azure-com-v1alpha1api20200601-resourcegroup,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=microsoft.resources.infra.azure.com,resources=resourcegroups,verbs=create;update,versions=v1alpha1api20200601,name=default.v1alpha1api20200601.resourcegroups.microsoft.resources.infra.azure.com,admissionReviewVersions=v1beta1

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

var _ genruntime.LocatableResource = &ResourceGroup{}

func (rg *ResourceGroup) Location() string {
	return rg.Spec.Location
}

// +kubebuilder:object:root=true
type ResourceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceGroup `json:"items"`
}

type ResourceGroupStatus struct {
	ID string `json:"id,omitempty"`

	Name     string `json:"name,omitempty"`
	Location string `json:"location,omitempty"`

	// ManagedBy is the management group responsible for managing this group
	ManagedBy string `json:"managedBy,omitempty"`

	// Tags are user defined key value pairs
	Tags map[string]string `json:"tags,omitempty"`


	ProvisioningState string `json:"provisioningState,omitempty"`

	// Conditions describe the observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`
}

var _ genruntime.FromARMConverter = &ResourceGroupStatus{}

func (status *ResourceGroupStatus) CreateEmptyARMValue() interface{} {
	return ResourceGroupStatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (status *ResourceGroupStatus) PopulateFromARM(owner genruntime.KnownResourceReference, armInput interface{}) error {
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

type ResourceGroupSpec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName string `json:"azureName"`

	// +kubebuilder:validation:Required
	// Location is the Azure location for the group (eg westus2, southcentralus, etc...)
	Location string `json:"location"`

	// ManagedBy is the management group responsible for managing this group
	ManagedBy string `json:"managedBy,omitempty"`

	// Tags are user defined key value pairs
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &ResourceGroupSpec{}

func (spec *ResourceGroupSpec) CreateEmptyARMValue() interface{} {
	return ResourceGroupSpecARM{}
}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (spec *ResourceGroupSpec) ConvertToARM(name string, resolvedReferences genruntime.ResolvedReferences) (interface{}, error) {
	if spec == nil {
		return nil, nil
	}
	result := ResourceGroupSpecARM{}
	result.APIVersion = "2020-06-01" // TODO: Update this to match what the codegenerated resources do with APIVersion eventually
	result.Location = spec.Location
	result.Name = name
	result.ManagedBy = spec.ManagedBy
	result.Tags = spec.Tags
	result.Type = ResourceGroupTypeResourceGroup
	return result, nil
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (spec *ResourceGroupSpec) PopulateFromARM(owner genruntime.KnownResourceReference, armInput interface{}) error {
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

func init() {
	SchemeBuilder.Register(&ResourceGroup{}, &ResourceGroupList{})
}
