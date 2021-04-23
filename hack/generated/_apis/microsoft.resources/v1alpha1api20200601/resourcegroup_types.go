// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"github.com/Azure/k8s-infra/hack/generated/pkg/genruntime"
)

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

	Properties *ResourceGroupStatusProperties `json:"properties,omitempty"` // TODO: Is this required or optional?
}

var _ genruntime.ARMTransformer = &ResourceGroupStatus{}

func (status *ResourceGroupStatus) CreateEmptyARMValue() interface{} {
	return ResourceGroupStatusArm{}
}

// ConvertToArm converts from a Kubernetes CRD object to an ARM object
func (status *ResourceGroupStatus) ConvertToARM(name string) (interface{}, error) {
	if status == nil {
		return nil, nil
	}
	result := ResourceGroupStatusArm{}
	result.ID = status.ID
	result.Location = status.Location
	result.ManagedBy = status.ManagedBy
	result.Name = status.Name
	result.Tags = status.Tags
	if status.Properties != nil {
		properties, err := status.Properties.ConvertToARM(name)
		if err != nil {
			return nil, err
		}
		propertiesTyped := properties.(ResourceGroupStatusPropertiesArm)
		result.Properties = &propertiesTyped
	}
	return result, nil
}

// PopulateFromArm populates a Kubernetes CRD object from an Azure ARM object
func (status *ResourceGroupStatus) PopulateFromARM(owner genruntime.KnownResourceReference, armInput interface{}) error {
	typedInput, ok := armInput.(ResourceGroupStatusArm)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromArm() function. Expected ResourceGroupStatusArm, got %T", armInput)
	}
	status.ID = typedInput.ID
	status.Location = typedInput.Location
	status.ManagedBy = typedInput.ManagedBy
	status.Name = typedInput.Name
	status.Tags = typedInput.Tags
	var err error
	if typedInput.Properties != nil {
		properties := ResourceGroupStatusProperties{}
		err = properties.PopulateFromARM(owner, *typedInput.Properties)
		if err != nil {
			return err
		}
		status.Properties = &properties
	}
	return nil
}

type ResourceGroupStatusProperties struct {
	ProvisioningState string `json:"provisioningState,omitempty"` // TODO: Wrong, needs to be in properties
}

var _ genruntime.ARMTransformer = &ResourceGroupStatusProperties{}

func (p *ResourceGroupStatusProperties) CreateEmptyARMValue() interface{} {
	return ResourceGroupStatusPropertiesArm{}
}

// ConvertToArm converts from a Kubernetes CRD object to an ARM object
func (p *ResourceGroupStatusProperties) ConvertToARM(name string) (interface{}, error) {
	if p == nil {
		return nil, nil
	}
	result := ResourceGroupStatusPropertiesArm{}
	result.ProvisioningState = p.ProvisioningState
	return result, nil
}

// PopulateFromArm populates a Kubernetes CRD object from an Azure ARM object
func (p *ResourceGroupStatusProperties) PopulateFromARM(owner genruntime.KnownResourceReference, armInput interface{}) error {
	typedInput, ok := armInput.(ResourceGroupStatusPropertiesArm)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromArm() function. Expected ResourceGroupStatusPropertiesArm, got %T", armInput)
	}
	p.ProvisioningState = typedInput.ProvisioningState
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
	return ResourceGroupSpecArm{}
}

// ConvertToArm converts from a Kubernetes CRD object to an ARM object
func (spec *ResourceGroupSpec) ConvertToARM(name string) (interface{}, error) {
	if spec == nil {
		return nil, nil
	}
	result := ResourceGroupSpecArm{}
	result.ApiVersion = "2020-06-01" // TODO: Update this to match what the codegenerated resources do with APIVersion eventually
	result.Location = spec.Location
	result.Name = name
	result.ManagedBy = spec.ManagedBy
	result.Tags = spec.Tags
	result.Type = ResourceGroupTypeResourceGroup
	return result, nil
}

// PopulateFromArm populates a Kubernetes CRD object from an Azure ARM object
func (spec *ResourceGroupSpec) PopulateFromARM(owner genruntime.KnownResourceReference, armInput interface{}) error {
	typedInput, ok := armInput.(ResourceGroupSpecArm)
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
