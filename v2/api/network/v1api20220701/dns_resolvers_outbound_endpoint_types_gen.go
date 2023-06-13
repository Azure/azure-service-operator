// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701

import (
	"fmt"
	v1api20220701s "github.com/Azure/azure-service-operator/v2/api/network/v1api20220701storage"
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
// Generator information:
// - Generated from: /dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/dnsresolver.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsResolvers/{dnsResolverName}/outboundEndpoints/{outboundEndpointName}
type DnsResolversOutboundEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsResolvers_OutboundEndpoint_Spec   `json:"spec,omitempty"`
	Status            DnsResolvers_OutboundEndpoint_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &DnsResolversOutboundEndpoint{}

// GetConditions returns the conditions of the resource
func (endpoint *DnsResolversOutboundEndpoint) GetConditions() conditions.Conditions {
	return endpoint.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (endpoint *DnsResolversOutboundEndpoint) SetConditions(conditions conditions.Conditions) {
	endpoint.Status.Conditions = conditions
}

var _ conversion.Convertible = &DnsResolversOutboundEndpoint{}

// ConvertFrom populates our DnsResolversOutboundEndpoint from the provided hub DnsResolversOutboundEndpoint
func (endpoint *DnsResolversOutboundEndpoint) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1api20220701s.DnsResolversOutboundEndpoint)
	if !ok {
		return fmt.Errorf("expected network/v1api20220701storage/DnsResolversOutboundEndpoint but received %T instead", hub)
	}

	return endpoint.AssignProperties_From_DnsResolversOutboundEndpoint(source)
}

// ConvertTo populates the provided hub DnsResolversOutboundEndpoint from our DnsResolversOutboundEndpoint
func (endpoint *DnsResolversOutboundEndpoint) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1api20220701s.DnsResolversOutboundEndpoint)
	if !ok {
		return fmt.Errorf("expected network/v1api20220701storage/DnsResolversOutboundEndpoint but received %T instead", hub)
	}

	return endpoint.AssignProperties_To_DnsResolversOutboundEndpoint(destination)
}

// +kubebuilder:webhook:path=/mutate-network-azure-com-v1api20220701-dnsresolversoutboundendpoint,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=dnsresolversoutboundendpoints,verbs=create;update,versions=v1api20220701,name=default.v1api20220701.dnsresolversoutboundendpoints.network.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &DnsResolversOutboundEndpoint{}

// Default applies defaults to the DnsResolversOutboundEndpoint resource
func (endpoint *DnsResolversOutboundEndpoint) Default() {
	endpoint.defaultImpl()
	var temp any = endpoint
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (endpoint *DnsResolversOutboundEndpoint) defaultAzureName() {
	if endpoint.Spec.AzureName == "" {
		endpoint.Spec.AzureName = endpoint.Name
	}
}

// defaultImpl applies the code generated defaults to the DnsResolversOutboundEndpoint resource
func (endpoint *DnsResolversOutboundEndpoint) defaultImpl() { endpoint.defaultAzureName() }

var _ genruntime.ImportableResource = &DnsResolversOutboundEndpoint{}

// InitializeSpec initializes the spec for this resource from the given status
func (endpoint *DnsResolversOutboundEndpoint) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*DnsResolvers_OutboundEndpoint_STATUS); ok {
		return endpoint.Spec.Initialize_From_DnsResolvers_OutboundEndpoint_STATUS(s)
	}

	return fmt.Errorf("expected Status of type DnsResolvers_OutboundEndpoint_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &DnsResolversOutboundEndpoint{}

// AzureName returns the Azure name of the resource
func (endpoint *DnsResolversOutboundEndpoint) AzureName() string {
	return endpoint.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-07-01"
func (endpoint DnsResolversOutboundEndpoint) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (endpoint *DnsResolversOutboundEndpoint) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (endpoint *DnsResolversOutboundEndpoint) GetSpec() genruntime.ConvertibleSpec {
	return &endpoint.Spec
}

// GetStatus returns the status of this resource
func (endpoint *DnsResolversOutboundEndpoint) GetStatus() genruntime.ConvertibleStatus {
	return &endpoint.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/dnsResolvers/outboundEndpoints"
func (endpoint *DnsResolversOutboundEndpoint) GetType() string {
	return "Microsoft.Network/dnsResolvers/outboundEndpoints"
}

// NewEmptyStatus returns a new empty (blank) status
func (endpoint *DnsResolversOutboundEndpoint) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DnsResolvers_OutboundEndpoint_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (endpoint *DnsResolversOutboundEndpoint) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(endpoint.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  endpoint.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (endpoint *DnsResolversOutboundEndpoint) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DnsResolvers_OutboundEndpoint_STATUS); ok {
		endpoint.Status = *st
		return nil
	}

	// Convert status to required version
	var st DnsResolvers_OutboundEndpoint_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	endpoint.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-network-azure-com-v1api20220701-dnsresolversoutboundendpoint,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=dnsresolversoutboundendpoints,verbs=create;update,versions=v1api20220701,name=validate.v1api20220701.dnsresolversoutboundendpoints.network.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &DnsResolversOutboundEndpoint{}

// ValidateCreate validates the creation of the resource
func (endpoint *DnsResolversOutboundEndpoint) ValidateCreate() error {
	validations := endpoint.createValidations()
	var temp any = endpoint
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
func (endpoint *DnsResolversOutboundEndpoint) ValidateDelete() error {
	validations := endpoint.deleteValidations()
	var temp any = endpoint
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
func (endpoint *DnsResolversOutboundEndpoint) ValidateUpdate(old runtime.Object) error {
	validations := endpoint.updateValidations()
	var temp any = endpoint
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
func (endpoint *DnsResolversOutboundEndpoint) createValidations() []func() error {
	return []func() error{endpoint.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (endpoint *DnsResolversOutboundEndpoint) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (endpoint *DnsResolversOutboundEndpoint) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return endpoint.validateResourceReferences()
		},
		endpoint.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (endpoint *DnsResolversOutboundEndpoint) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&endpoint.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (endpoint *DnsResolversOutboundEndpoint) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*DnsResolversOutboundEndpoint)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, endpoint)
}

// AssignProperties_From_DnsResolversOutboundEndpoint populates our DnsResolversOutboundEndpoint from the provided source DnsResolversOutboundEndpoint
func (endpoint *DnsResolversOutboundEndpoint) AssignProperties_From_DnsResolversOutboundEndpoint(source *v1api20220701s.DnsResolversOutboundEndpoint) error {

	// ObjectMeta
	endpoint.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DnsResolvers_OutboundEndpoint_Spec
	err := spec.AssignProperties_From_DnsResolvers_OutboundEndpoint_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_DnsResolvers_OutboundEndpoint_Spec() to populate field Spec")
	}
	endpoint.Spec = spec

	// Status
	var status DnsResolvers_OutboundEndpoint_STATUS
	err = status.AssignProperties_From_DnsResolvers_OutboundEndpoint_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_DnsResolvers_OutboundEndpoint_STATUS() to populate field Status")
	}
	endpoint.Status = status

	// No error
	return nil
}

// AssignProperties_To_DnsResolversOutboundEndpoint populates the provided destination DnsResolversOutboundEndpoint from our DnsResolversOutboundEndpoint
func (endpoint *DnsResolversOutboundEndpoint) AssignProperties_To_DnsResolversOutboundEndpoint(destination *v1api20220701s.DnsResolversOutboundEndpoint) error {

	// ObjectMeta
	destination.ObjectMeta = *endpoint.ObjectMeta.DeepCopy()

	// Spec
	var spec v1api20220701s.DnsResolvers_OutboundEndpoint_Spec
	err := endpoint.Spec.AssignProperties_To_DnsResolvers_OutboundEndpoint_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_DnsResolvers_OutboundEndpoint_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1api20220701s.DnsResolvers_OutboundEndpoint_STATUS
	err = endpoint.Status.AssignProperties_To_DnsResolvers_OutboundEndpoint_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_DnsResolvers_OutboundEndpoint_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (endpoint *DnsResolversOutboundEndpoint) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: endpoint.Spec.OriginalVersion(),
		Kind:    "DnsResolversOutboundEndpoint",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/dnsresolver.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsResolvers/{dnsResolverName}/outboundEndpoints/{outboundEndpointName}
type DnsResolversOutboundEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DnsResolversOutboundEndpoint `json:"items"`
}

type DnsResolvers_OutboundEndpoint_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:Required
	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a network.azure.com/DnsResolver resource
	Owner *genruntime.KnownResourceReference `group:"network.azure.com" json:"owner,omitempty" kind:"DnsResolver"`

	// +kubebuilder:validation:Required
	// Subnet: The reference to the subnet used for the outbound endpoint.
	Subnet *DnsresolverSubResource `json:"subnet,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &DnsResolvers_OutboundEndpoint_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (endpoint *DnsResolvers_OutboundEndpoint_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if endpoint == nil {
		return nil, nil
	}
	result := &DnsResolvers_OutboundEndpoint_Spec_ARM{}

	// Set property ‘Location’:
	if endpoint.Location != nil {
		location := *endpoint.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if endpoint.Subnet != nil {
		result.Properties = &OutboundEndpointProperties_ARM{}
	}
	if endpoint.Subnet != nil {
		subnet_ARM, err := (*endpoint.Subnet).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		subnet := *subnet_ARM.(*DnsresolverSubResource_ARM)
		result.Properties.Subnet = &subnet
	}

	// Set property ‘Tags’:
	if endpoint.Tags != nil {
		result.Tags = make(map[string]string, len(endpoint.Tags))
		for key, value := range endpoint.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (endpoint *DnsResolvers_OutboundEndpoint_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DnsResolvers_OutboundEndpoint_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (endpoint *DnsResolvers_OutboundEndpoint_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DnsResolvers_OutboundEndpoint_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DnsResolvers_OutboundEndpoint_Spec_ARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	endpoint.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		endpoint.Location = &location
	}

	// Set property ‘Owner’:
	endpoint.Owner = &genruntime.KnownResourceReference{Name: owner.Name}

	// Set property ‘Subnet’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Subnet != nil {
			var subnet1 DnsresolverSubResource
			err := subnet1.PopulateFromARM(owner, *typedInput.Properties.Subnet)
			if err != nil {
				return err
			}
			subnet := subnet1
			endpoint.Subnet = &subnet
		}
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		endpoint.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			endpoint.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &DnsResolvers_OutboundEndpoint_Spec{}

// ConvertSpecFrom populates our DnsResolvers_OutboundEndpoint_Spec from the provided source
func (endpoint *DnsResolvers_OutboundEndpoint_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1api20220701s.DnsResolvers_OutboundEndpoint_Spec)
	if ok {
		// Populate our instance from source
		return endpoint.AssignProperties_From_DnsResolvers_OutboundEndpoint_Spec(src)
	}

	// Convert to an intermediate form
	src = &v1api20220701s.DnsResolvers_OutboundEndpoint_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = endpoint.AssignProperties_From_DnsResolvers_OutboundEndpoint_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DnsResolvers_OutboundEndpoint_Spec
func (endpoint *DnsResolvers_OutboundEndpoint_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1api20220701s.DnsResolvers_OutboundEndpoint_Spec)
	if ok {
		// Populate destination from our instance
		return endpoint.AssignProperties_To_DnsResolvers_OutboundEndpoint_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v1api20220701s.DnsResolvers_OutboundEndpoint_Spec{}
	err := endpoint.AssignProperties_To_DnsResolvers_OutboundEndpoint_Spec(dst)
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

// AssignProperties_From_DnsResolvers_OutboundEndpoint_Spec populates our DnsResolvers_OutboundEndpoint_Spec from the provided source DnsResolvers_OutboundEndpoint_Spec
func (endpoint *DnsResolvers_OutboundEndpoint_Spec) AssignProperties_From_DnsResolvers_OutboundEndpoint_Spec(source *v1api20220701s.DnsResolvers_OutboundEndpoint_Spec) error {

	// AzureName
	endpoint.AzureName = source.AzureName

	// Location
	endpoint.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		endpoint.Owner = &owner
	} else {
		endpoint.Owner = nil
	}

	// Subnet
	if source.Subnet != nil {
		var subnet DnsresolverSubResource
		err := subnet.AssignProperties_From_DnsresolverSubResource(source.Subnet)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_DnsresolverSubResource() to populate field Subnet")
		}
		endpoint.Subnet = &subnet
	} else {
		endpoint.Subnet = nil
	}

	// Tags
	endpoint.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignProperties_To_DnsResolvers_OutboundEndpoint_Spec populates the provided destination DnsResolvers_OutboundEndpoint_Spec from our DnsResolvers_OutboundEndpoint_Spec
func (endpoint *DnsResolvers_OutboundEndpoint_Spec) AssignProperties_To_DnsResolvers_OutboundEndpoint_Spec(destination *v1api20220701s.DnsResolvers_OutboundEndpoint_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = endpoint.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(endpoint.Location)

	// OriginalVersion
	destination.OriginalVersion = endpoint.OriginalVersion()

	// Owner
	if endpoint.Owner != nil {
		owner := endpoint.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Subnet
	if endpoint.Subnet != nil {
		var subnet v1api20220701s.DnsresolverSubResource
		err := endpoint.Subnet.AssignProperties_To_DnsresolverSubResource(&subnet)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_DnsresolverSubResource() to populate field Subnet")
		}
		destination.Subnet = &subnet
	} else {
		destination.Subnet = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(endpoint.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Initialize_From_DnsResolvers_OutboundEndpoint_STATUS populates our DnsResolvers_OutboundEndpoint_Spec from the provided source DnsResolvers_OutboundEndpoint_STATUS
func (endpoint *DnsResolvers_OutboundEndpoint_Spec) Initialize_From_DnsResolvers_OutboundEndpoint_STATUS(source *DnsResolvers_OutboundEndpoint_STATUS) error {

	// Location
	endpoint.Location = genruntime.ClonePointerToString(source.Location)

	// Subnet
	if source.Subnet != nil {
		var subnet DnsresolverSubResource
		err := subnet.Initialize_From_DnsresolverSubResource_STATUS(source.Subnet)
		if err != nil {
			return errors.Wrap(err, "calling Initialize_From_DnsresolverSubResource_STATUS() to populate field Subnet")
		}
		endpoint.Subnet = &subnet
	} else {
		endpoint.Subnet = nil
	}

	// Tags
	endpoint.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (endpoint *DnsResolvers_OutboundEndpoint_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (endpoint *DnsResolvers_OutboundEndpoint_Spec) SetAzureName(azureName string) {
	endpoint.AzureName = azureName
}

type DnsResolvers_OutboundEndpoint_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Etag: ETag of the outbound endpoint.
	Etag *string `json:"etag,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// ProvisioningState: The current provisioning state of the outbound endpoint. This is a read-only property and any attempt
	// to set this value will be ignored.
	ProvisioningState *DnsresolverProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// ResourceGuid: The resourceGuid property of the outbound endpoint resource.
	ResourceGuid *string `json:"resourceGuid,omitempty"`

	// Subnet: The reference to the subnet used for the outbound endpoint.
	Subnet *DnsresolverSubResource_STATUS `json:"subnet,omitempty"`

	// SystemData: Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DnsResolvers_OutboundEndpoint_STATUS{}

// ConvertStatusFrom populates our DnsResolvers_OutboundEndpoint_STATUS from the provided source
func (endpoint *DnsResolvers_OutboundEndpoint_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1api20220701s.DnsResolvers_OutboundEndpoint_STATUS)
	if ok {
		// Populate our instance from source
		return endpoint.AssignProperties_From_DnsResolvers_OutboundEndpoint_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v1api20220701s.DnsResolvers_OutboundEndpoint_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = endpoint.AssignProperties_From_DnsResolvers_OutboundEndpoint_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our DnsResolvers_OutboundEndpoint_STATUS
func (endpoint *DnsResolvers_OutboundEndpoint_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1api20220701s.DnsResolvers_OutboundEndpoint_STATUS)
	if ok {
		// Populate destination from our instance
		return endpoint.AssignProperties_To_DnsResolvers_OutboundEndpoint_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v1api20220701s.DnsResolvers_OutboundEndpoint_STATUS{}
	err := endpoint.AssignProperties_To_DnsResolvers_OutboundEndpoint_STATUS(dst)
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

var _ genruntime.FromARMConverter = &DnsResolvers_OutboundEndpoint_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (endpoint *DnsResolvers_OutboundEndpoint_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DnsResolvers_OutboundEndpoint_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (endpoint *DnsResolvers_OutboundEndpoint_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DnsResolvers_OutboundEndpoint_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DnsResolvers_OutboundEndpoint_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Etag’:
	if typedInput.Etag != nil {
		etag := *typedInput.Etag
		endpoint.Etag = &etag
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		endpoint.Id = &id
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		endpoint.Location = &location
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		endpoint.Name = &name
	}

	// Set property ‘ProvisioningState’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ProvisioningState != nil {
			provisioningState := *typedInput.Properties.ProvisioningState
			endpoint.ProvisioningState = &provisioningState
		}
	}

	// Set property ‘ResourceGuid’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ResourceGuid != nil {
			resourceGuid := *typedInput.Properties.ResourceGuid
			endpoint.ResourceGuid = &resourceGuid
		}
	}

	// Set property ‘Subnet’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Subnet != nil {
			var subnet1 DnsresolverSubResource_STATUS
			err := subnet1.PopulateFromARM(owner, *typedInput.Properties.Subnet)
			if err != nil {
				return err
			}
			subnet := subnet1
			endpoint.Subnet = &subnet
		}
	}

	// Set property ‘SystemData’:
	if typedInput.SystemData != nil {
		var systemData1 SystemData_STATUS
		err := systemData1.PopulateFromARM(owner, *typedInput.SystemData)
		if err != nil {
			return err
		}
		systemData := systemData1
		endpoint.SystemData = &systemData
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		endpoint.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			endpoint.Tags[key] = value
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		endpoint.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_DnsResolvers_OutboundEndpoint_STATUS populates our DnsResolvers_OutboundEndpoint_STATUS from the provided source DnsResolvers_OutboundEndpoint_STATUS
func (endpoint *DnsResolvers_OutboundEndpoint_STATUS) AssignProperties_From_DnsResolvers_OutboundEndpoint_STATUS(source *v1api20220701s.DnsResolvers_OutboundEndpoint_STATUS) error {

	// Conditions
	endpoint.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Etag
	endpoint.Etag = genruntime.ClonePointerToString(source.Etag)

	// Id
	endpoint.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	endpoint.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	endpoint.Name = genruntime.ClonePointerToString(source.Name)

	// ProvisioningState
	if source.ProvisioningState != nil {
		provisioningState := DnsresolverProvisioningState_STATUS(*source.ProvisioningState)
		endpoint.ProvisioningState = &provisioningState
	} else {
		endpoint.ProvisioningState = nil
	}

	// ResourceGuid
	endpoint.ResourceGuid = genruntime.ClonePointerToString(source.ResourceGuid)

	// Subnet
	if source.Subnet != nil {
		var subnet DnsresolverSubResource_STATUS
		err := subnet.AssignProperties_From_DnsresolverSubResource_STATUS(source.Subnet)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_DnsresolverSubResource_STATUS() to populate field Subnet")
		}
		endpoint.Subnet = &subnet
	} else {
		endpoint.Subnet = nil
	}

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
		}
		endpoint.SystemData = &systemDatum
	} else {
		endpoint.SystemData = nil
	}

	// Tags
	endpoint.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	endpoint.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_DnsResolvers_OutboundEndpoint_STATUS populates the provided destination DnsResolvers_OutboundEndpoint_STATUS from our DnsResolvers_OutboundEndpoint_STATUS
func (endpoint *DnsResolvers_OutboundEndpoint_STATUS) AssignProperties_To_DnsResolvers_OutboundEndpoint_STATUS(destination *v1api20220701s.DnsResolvers_OutboundEndpoint_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(endpoint.Conditions)

	// Etag
	destination.Etag = genruntime.ClonePointerToString(endpoint.Etag)

	// Id
	destination.Id = genruntime.ClonePointerToString(endpoint.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(endpoint.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(endpoint.Name)

	// ProvisioningState
	if endpoint.ProvisioningState != nil {
		provisioningState := string(*endpoint.ProvisioningState)
		destination.ProvisioningState = &provisioningState
	} else {
		destination.ProvisioningState = nil
	}

	// ResourceGuid
	destination.ResourceGuid = genruntime.ClonePointerToString(endpoint.ResourceGuid)

	// Subnet
	if endpoint.Subnet != nil {
		var subnet v1api20220701s.DnsresolverSubResource_STATUS
		err := endpoint.Subnet.AssignProperties_To_DnsresolverSubResource_STATUS(&subnet)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_DnsresolverSubResource_STATUS() to populate field Subnet")
		}
		destination.Subnet = &subnet
	} else {
		destination.Subnet = nil
	}

	// SystemData
	if endpoint.SystemData != nil {
		var systemDatum v1api20220701s.SystemData_STATUS
		err := endpoint.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(endpoint.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(endpoint.Type)

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
	SchemeBuilder.Register(&DnsResolversOutboundEndpoint{}, &DnsResolversOutboundEndpointList{})
}
