// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

import (
	"fmt"
	v20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101storage"
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
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2020-11-01/routeTable.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes/{routeName}
type RouteTablesRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteTablesRoute_Spec   `json:"spec,omitempty"`
	Status            RouteTablesRoute_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RouteTablesRoute{}

// GetConditions returns the conditions of the resource
func (route *RouteTablesRoute) GetConditions() conditions.Conditions {
	return route.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (route *RouteTablesRoute) SetConditions(conditions conditions.Conditions) {
	route.Status.Conditions = conditions
}

var _ conversion.Convertible = &RouteTablesRoute{}

// ConvertFrom populates our RouteTablesRoute from the provided hub RouteTablesRoute
func (route *RouteTablesRoute) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20201101s.RouteTablesRoute)
	if !ok {
		return fmt.Errorf("expected network/v1beta20201101storage/RouteTablesRoute but received %T instead", hub)
	}

	return route.AssignPropertiesFromRouteTablesRoute(source)
}

// ConvertTo populates the provided hub RouteTablesRoute from our RouteTablesRoute
func (route *RouteTablesRoute) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20201101s.RouteTablesRoute)
	if !ok {
		return fmt.Errorf("expected network/v1beta20201101storage/RouteTablesRoute but received %T instead", hub)
	}

	return route.AssignPropertiesToRouteTablesRoute(destination)
}

// +kubebuilder:webhook:path=/mutate-network-azure-com-v1beta20201101-routetablesroute,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=routetablesroutes,verbs=create;update,versions=v1beta20201101,name=default.v1beta20201101.routetablesroutes.network.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &RouteTablesRoute{}

// Default applies defaults to the RouteTablesRoute resource
func (route *RouteTablesRoute) Default() {
	route.defaultImpl()
	var temp interface{} = route
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (route *RouteTablesRoute) defaultAzureName() {
	if route.Spec.AzureName == "" {
		route.Spec.AzureName = route.Name
	}
}

// defaultImpl applies the code generated defaults to the RouteTablesRoute resource
func (route *RouteTablesRoute) defaultImpl() { route.defaultAzureName() }

var _ genruntime.KubernetesResource = &RouteTablesRoute{}

// AzureName returns the Azure name of the resource
func (route *RouteTablesRoute) AzureName() string {
	return route.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (route RouteTablesRoute) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceKind returns the kind of the resource
func (route *RouteTablesRoute) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (route *RouteTablesRoute) GetSpec() genruntime.ConvertibleSpec {
	return &route.Spec
}

// GetStatus returns the status of this resource
func (route *RouteTablesRoute) GetStatus() genruntime.ConvertibleStatus {
	return &route.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/routeTables/routes"
func (route *RouteTablesRoute) GetType() string {
	return "Microsoft.Network/routeTables/routes"
}

// NewEmptyStatus returns a new empty (blank) status
func (route *RouteTablesRoute) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RouteTablesRoute_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (route *RouteTablesRoute) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(route.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  route.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (route *RouteTablesRoute) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RouteTablesRoute_STATUS); ok {
		route.Status = *st
		return nil
	}

	// Convert status to required version
	var st RouteTablesRoute_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	route.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-network-azure-com-v1beta20201101-routetablesroute,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=routetablesroutes,verbs=create;update,versions=v1beta20201101,name=validate.v1beta20201101.routetablesroutes.network.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &RouteTablesRoute{}

// ValidateCreate validates the creation of the resource
func (route *RouteTablesRoute) ValidateCreate() error {
	validations := route.createValidations()
	var temp interface{} = route
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
func (route *RouteTablesRoute) ValidateDelete() error {
	validations := route.deleteValidations()
	var temp interface{} = route
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
func (route *RouteTablesRoute) ValidateUpdate(old runtime.Object) error {
	validations := route.updateValidations()
	var temp interface{} = route
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
func (route *RouteTablesRoute) createValidations() []func() error {
	return []func() error{route.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (route *RouteTablesRoute) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (route *RouteTablesRoute) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return route.validateResourceReferences()
		},
		route.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (route *RouteTablesRoute) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&route.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (route *RouteTablesRoute) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*RouteTablesRoute)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, route)
}

// AssignPropertiesFromRouteTablesRoute populates our RouteTablesRoute from the provided source RouteTablesRoute
func (route *RouteTablesRoute) AssignPropertiesFromRouteTablesRoute(source *v20201101s.RouteTablesRoute) error {

	// ObjectMeta
	route.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RouteTablesRoute_Spec
	err := spec.AssignPropertiesFromRouteTablesRoute_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRouteTablesRoute_Spec() to populate field Spec")
	}
	route.Spec = spec

	// Status
	var status RouteTablesRoute_STATUS
	err = status.AssignPropertiesFromRouteTablesRoute_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRouteTablesRoute_STATUS() to populate field Status")
	}
	route.Status = status

	// No error
	return nil
}

// AssignPropertiesToRouteTablesRoute populates the provided destination RouteTablesRoute from our RouteTablesRoute
func (route *RouteTablesRoute) AssignPropertiesToRouteTablesRoute(destination *v20201101s.RouteTablesRoute) error {

	// ObjectMeta
	destination.ObjectMeta = *route.ObjectMeta.DeepCopy()

	// Spec
	var spec v20201101s.RouteTablesRoute_Spec
	err := route.Spec.AssignPropertiesToRouteTablesRoute_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRouteTablesRoute_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20201101s.RouteTablesRoute_STATUS
	err = route.Status.AssignPropertiesToRouteTablesRoute_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRouteTablesRoute_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (route *RouteTablesRoute) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: route.Spec.OriginalVersion(),
		Kind:    "RouteTablesRoute",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2020-11-01/routeTable.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes/{routeName}
type RouteTablesRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTablesRoute `json:"items"`
}

type RouteTablesRoute_STATUS struct {
	// AddressPrefix: The destination CIDR to which the route applies.
	AddressPrefix *string `json:"addressPrefix,omitempty"`

	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// HasBgpOverride: A value indicating whether this route overrides overlapping BGP routes regardless of LPM.
	HasBgpOverride *bool `json:"hasBgpOverride,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// NextHopIpAddress: The IP address packets should be forwarded to. Next hop values are only allowed in routes where the
	// next hop type is VirtualAppliance.
	NextHopIpAddress *string `json:"nextHopIpAddress,omitempty"`

	// NextHopType: The type of Azure hop the packet should be sent to.
	NextHopType *RouteNextHopType_STATUS `json:"nextHopType,omitempty"`

	// ProvisioningState: The provisioning state of the route resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// Type: The type of the resource.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RouteTablesRoute_STATUS{}

// ConvertStatusFrom populates our RouteTablesRoute_STATUS from the provided source
func (route *RouteTablesRoute_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20201101s.RouteTablesRoute_STATUS)
	if ok {
		// Populate our instance from source
		return route.AssignPropertiesFromRouteTablesRoute_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20201101s.RouteTablesRoute_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = route.AssignPropertiesFromRouteTablesRoute_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our RouteTablesRoute_STATUS
func (route *RouteTablesRoute_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20201101s.RouteTablesRoute_STATUS)
	if ok {
		// Populate destination from our instance
		return route.AssignPropertiesToRouteTablesRoute_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20201101s.RouteTablesRoute_STATUS{}
	err := route.AssignPropertiesToRouteTablesRoute_STATUS(dst)
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

var _ genruntime.FromARMConverter = &RouteTablesRoute_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (route *RouteTablesRoute_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RouteTablesRoute_STATUSARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (route *RouteTablesRoute_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RouteTablesRoute_STATUSARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RouteTablesRoute_STATUSARM, got %T", armInput)
	}

	// Set property ‘AddressPrefix’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.AddressPrefix != nil {
			addressPrefix := *typedInput.Properties.AddressPrefix
			route.AddressPrefix = &addressPrefix
		}
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Etag’:
	if typedInput.Etag != nil {
		etag := *typedInput.Etag
		route.Etag = &etag
	}

	// Set property ‘HasBgpOverride’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.HasBgpOverride != nil {
			hasBgpOverride := *typedInput.Properties.HasBgpOverride
			route.HasBgpOverride = &hasBgpOverride
		}
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		route.Id = &id
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		route.Name = &name
	}

	// Set property ‘NextHopIpAddress’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.NextHopIpAddress != nil {
			nextHopIpAddress := *typedInput.Properties.NextHopIpAddress
			route.NextHopIpAddress = &nextHopIpAddress
		}
	}

	// Set property ‘NextHopType’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.NextHopType != nil {
			nextHopType := *typedInput.Properties.NextHopType
			route.NextHopType = &nextHopType
		}
	}

	// Set property ‘ProvisioningState’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ProvisioningState != nil {
			provisioningState := *typedInput.Properties.ProvisioningState
			route.ProvisioningState = &provisioningState
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		route.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromRouteTablesRoute_STATUS populates our RouteTablesRoute_STATUS from the provided source RouteTablesRoute_STATUS
func (route *RouteTablesRoute_STATUS) AssignPropertiesFromRouteTablesRoute_STATUS(source *v20201101s.RouteTablesRoute_STATUS) error {

	// AddressPrefix
	route.AddressPrefix = genruntime.ClonePointerToString(source.AddressPrefix)

	// Conditions
	route.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Etag
	route.Etag = genruntime.ClonePointerToString(source.Etag)

	// HasBgpOverride
	if source.HasBgpOverride != nil {
		hasBgpOverride := *source.HasBgpOverride
		route.HasBgpOverride = &hasBgpOverride
	} else {
		route.HasBgpOverride = nil
	}

	// Id
	route.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	route.Name = genruntime.ClonePointerToString(source.Name)

	// NextHopIpAddress
	route.NextHopIpAddress = genruntime.ClonePointerToString(source.NextHopIpAddress)

	// NextHopType
	if source.NextHopType != nil {
		nextHopType := RouteNextHopType_STATUS(*source.NextHopType)
		route.NextHopType = &nextHopType
	} else {
		route.NextHopType = nil
	}

	// ProvisioningState
	if source.ProvisioningState != nil {
		provisioningState := ProvisioningState_STATUS(*source.ProvisioningState)
		route.ProvisioningState = &provisioningState
	} else {
		route.ProvisioningState = nil
	}

	// Type
	route.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToRouteTablesRoute_STATUS populates the provided destination RouteTablesRoute_STATUS from our RouteTablesRoute_STATUS
func (route *RouteTablesRoute_STATUS) AssignPropertiesToRouteTablesRoute_STATUS(destination *v20201101s.RouteTablesRoute_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AddressPrefix
	destination.AddressPrefix = genruntime.ClonePointerToString(route.AddressPrefix)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(route.Conditions)

	// Etag
	destination.Etag = genruntime.ClonePointerToString(route.Etag)

	// HasBgpOverride
	if route.HasBgpOverride != nil {
		hasBgpOverride := *route.HasBgpOverride
		destination.HasBgpOverride = &hasBgpOverride
	} else {
		destination.HasBgpOverride = nil
	}

	// Id
	destination.Id = genruntime.ClonePointerToString(route.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(route.Name)

	// NextHopIpAddress
	destination.NextHopIpAddress = genruntime.ClonePointerToString(route.NextHopIpAddress)

	// NextHopType
	if route.NextHopType != nil {
		nextHopType := string(*route.NextHopType)
		destination.NextHopType = &nextHopType
	} else {
		destination.NextHopType = nil
	}

	// ProvisioningState
	if route.ProvisioningState != nil {
		provisioningState := string(*route.ProvisioningState)
		destination.ProvisioningState = &provisioningState
	} else {
		destination.ProvisioningState = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(route.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type RouteTablesRoute_Spec struct {
	// AddressPrefix: The destination CIDR to which the route applies.
	AddressPrefix *string `json:"addressPrefix,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// HasBgpOverride: A value indicating whether this route overrides overlapping BGP routes regardless of LPM.
	HasBgpOverride *bool `json:"hasBgpOverride,omitempty"`

	// NextHopIpAddress: The IP address packets should be forwarded to. Next hop values are only allowed in routes where the
	// next hop type is VirtualAppliance.
	NextHopIpAddress *string `json:"nextHopIpAddress,omitempty"`

	// +kubebuilder:validation:Required
	// NextHopType: The type of Azure hop the packet should be sent to.
	NextHopType *RouteNextHopType `json:"nextHopType,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a network.azure.com/RouteTable resource
	Owner *genruntime.KnownResourceReference `group:"network.azure.com" json:"owner,omitempty" kind:"RouteTable"`

	// Reference: Resource ID.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`

	// Type: The type of the resource.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ARMTransformer = &RouteTablesRoute_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (route *RouteTablesRoute_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if route == nil {
		return nil, nil
	}
	result := &RouteTablesRoute_SpecARM{}

	// Set property ‘AzureName’:
	result.AzureName = route.AzureName

	// Set property ‘Id’:
	if route.Reference != nil {
		referenceARMID, err := resolved.ResolvedReferences.ARMIDOrErr(*route.Reference)
		if err != nil {
			return nil, err
		}
		reference := referenceARMID
		result.Id = &reference
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if route.AddressPrefix != nil ||
		route.HasBgpOverride != nil ||
		route.NextHopIpAddress != nil ||
		route.NextHopType != nil {
		result.Properties = &RoutePropertiesFormatARM{}
	}
	if route.AddressPrefix != nil {
		addressPrefix := *route.AddressPrefix
		result.Properties.AddressPrefix = &addressPrefix
	}
	if route.HasBgpOverride != nil {
		hasBgpOverride := *route.HasBgpOverride
		result.Properties.HasBgpOverride = &hasBgpOverride
	}
	if route.NextHopIpAddress != nil {
		nextHopIpAddress := *route.NextHopIpAddress
		result.Properties.NextHopIpAddress = &nextHopIpAddress
	}
	if route.NextHopType != nil {
		nextHopType := *route.NextHopType
		result.Properties.NextHopType = &nextHopType
	}

	// Set property ‘Type’:
	if route.Type != nil {
		typeVar := *route.Type
		result.Type = &typeVar
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (route *RouteTablesRoute_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RouteTablesRoute_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (route *RouteTablesRoute_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RouteTablesRoute_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RouteTablesRoute_SpecARM, got %T", armInput)
	}

	// Set property ‘AddressPrefix’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.AddressPrefix != nil {
			addressPrefix := *typedInput.Properties.AddressPrefix
			route.AddressPrefix = &addressPrefix
		}
	}

	// Set property ‘AzureName’:
	route.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘HasBgpOverride’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.HasBgpOverride != nil {
			hasBgpOverride := *typedInput.Properties.HasBgpOverride
			route.HasBgpOverride = &hasBgpOverride
		}
	}

	// Set property ‘NextHopIpAddress’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.NextHopIpAddress != nil {
			nextHopIpAddress := *typedInput.Properties.NextHopIpAddress
			route.NextHopIpAddress = &nextHopIpAddress
		}
	}

	// Set property ‘NextHopType’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.NextHopType != nil {
			nextHopType := *typedInput.Properties.NextHopType
			route.NextHopType = &nextHopType
		}
	}

	// Set property ‘Owner’:
	route.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// no assignment for property ‘Reference’

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		route.Type = &typeVar
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &RouteTablesRoute_Spec{}

// ConvertSpecFrom populates our RouteTablesRoute_Spec from the provided source
func (route *RouteTablesRoute_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20201101s.RouteTablesRoute_Spec)
	if ok {
		// Populate our instance from source
		return route.AssignPropertiesFromRouteTablesRoute_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20201101s.RouteTablesRoute_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = route.AssignPropertiesFromRouteTablesRoute_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RouteTablesRoute_Spec
func (route *RouteTablesRoute_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20201101s.RouteTablesRoute_Spec)
	if ok {
		// Populate destination from our instance
		return route.AssignPropertiesToRouteTablesRoute_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20201101s.RouteTablesRoute_Spec{}
	err := route.AssignPropertiesToRouteTablesRoute_Spec(dst)
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

// AssignPropertiesFromRouteTablesRoute_Spec populates our RouteTablesRoute_Spec from the provided source RouteTablesRoute_Spec
func (route *RouteTablesRoute_Spec) AssignPropertiesFromRouteTablesRoute_Spec(source *v20201101s.RouteTablesRoute_Spec) error {

	// AddressPrefix
	route.AddressPrefix = genruntime.ClonePointerToString(source.AddressPrefix)

	// AzureName
	route.AzureName = source.AzureName

	// HasBgpOverride
	if source.HasBgpOverride != nil {
		hasBgpOverride := *source.HasBgpOverride
		route.HasBgpOverride = &hasBgpOverride
	} else {
		route.HasBgpOverride = nil
	}

	// NextHopIpAddress
	route.NextHopIpAddress = genruntime.ClonePointerToString(source.NextHopIpAddress)

	// NextHopType
	if source.NextHopType != nil {
		nextHopType := RouteNextHopType(*source.NextHopType)
		route.NextHopType = &nextHopType
	} else {
		route.NextHopType = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		route.Owner = &owner
	} else {
		route.Owner = nil
	}

	// Reference
	if source.Reference != nil {
		reference := source.Reference.Copy()
		route.Reference = &reference
	} else {
		route.Reference = nil
	}

	// Type
	route.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToRouteTablesRoute_Spec populates the provided destination RouteTablesRoute_Spec from our RouteTablesRoute_Spec
func (route *RouteTablesRoute_Spec) AssignPropertiesToRouteTablesRoute_Spec(destination *v20201101s.RouteTablesRoute_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AddressPrefix
	destination.AddressPrefix = genruntime.ClonePointerToString(route.AddressPrefix)

	// AzureName
	destination.AzureName = route.AzureName

	// HasBgpOverride
	if route.HasBgpOverride != nil {
		hasBgpOverride := *route.HasBgpOverride
		destination.HasBgpOverride = &hasBgpOverride
	} else {
		destination.HasBgpOverride = nil
	}

	// NextHopIpAddress
	destination.NextHopIpAddress = genruntime.ClonePointerToString(route.NextHopIpAddress)

	// NextHopType
	if route.NextHopType != nil {
		nextHopType := string(*route.NextHopType)
		destination.NextHopType = &nextHopType
	} else {
		destination.NextHopType = nil
	}

	// OriginalVersion
	destination.OriginalVersion = route.OriginalVersion()

	// Owner
	if route.Owner != nil {
		owner := route.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Reference
	if route.Reference != nil {
		reference := route.Reference.Copy()
		destination.Reference = &reference
	} else {
		destination.Reference = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(route.Type)

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
func (route *RouteTablesRoute_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (route *RouteTablesRoute_Spec) SetAzureName(azureName string) { route.AzureName = azureName }

// +kubebuilder:validation:Enum={"Internet","None","VirtualAppliance","VirtualNetworkGateway","VnetLocal"}
type RouteNextHopType string

const (
	RouteNextHopType_Internet              = RouteNextHopType("Internet")
	RouteNextHopType_None                  = RouteNextHopType("None")
	RouteNextHopType_VirtualAppliance      = RouteNextHopType("VirtualAppliance")
	RouteNextHopType_VirtualNetworkGateway = RouteNextHopType("VirtualNetworkGateway")
	RouteNextHopType_VnetLocal             = RouteNextHopType("VnetLocal")
)

type RouteNextHopType_STATUS string

const (
	RouteNextHopType_Internet_STATUS              = RouteNextHopType_STATUS("Internet")
	RouteNextHopType_None_STATUS                  = RouteNextHopType_STATUS("None")
	RouteNextHopType_VirtualAppliance_STATUS      = RouteNextHopType_STATUS("VirtualAppliance")
	RouteNextHopType_VirtualNetworkGateway_STATUS = RouteNextHopType_STATUS("VirtualNetworkGateway")
	RouteNextHopType_VnetLocal_STATUS             = RouteNextHopType_STATUS("VnetLocal")
)

func init() {
	SchemeBuilder.Register(&RouteTablesRoute{}, &RouteTablesRouteList{})
}
