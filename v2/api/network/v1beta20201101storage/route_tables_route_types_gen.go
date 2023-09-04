// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101storage

import (
	"fmt"
	v20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20201101.RouteTablesRoute
// Deprecated version of RouteTablesRoute. Use v1api20201101.RouteTablesRoute instead
type RouteTablesRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteTables_Route_Spec   `json:"spec,omitempty"`
	Status            RouteTables_Route_STATUS `json:"status,omitempty"`
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
		return fmt.Errorf("expected network/v1api20201101storage/RouteTablesRoute but received %T instead", hub)
	}

	return route.AssignProperties_From_RouteTablesRoute(source)
}

// ConvertTo populates the provided hub RouteTablesRoute from our RouteTablesRoute
func (route *RouteTablesRoute) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20201101s.RouteTablesRoute)
	if !ok {
		return fmt.Errorf("expected network/v1api20201101storage/RouteTablesRoute but received %T instead", hub)
	}

	return route.AssignProperties_To_RouteTablesRoute(destination)
}

var _ genruntime.KubernetesResource = &RouteTablesRoute{}

// AzureName returns the Azure name of the resource
func (route *RouteTablesRoute) AzureName() string {
	return route.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (route RouteTablesRoute) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (route *RouteTablesRoute) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
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
	return &RouteTables_Route_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (route *RouteTablesRoute) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(route.Spec)
	return route.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (route *RouteTablesRoute) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RouteTables_Route_STATUS); ok {
		route.Status = *st
		return nil
	}

	// Convert status to required version
	var st RouteTables_Route_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	route.Status = st
	return nil
}

// AssignProperties_From_RouteTablesRoute populates our RouteTablesRoute from the provided source RouteTablesRoute
func (route *RouteTablesRoute) AssignProperties_From_RouteTablesRoute(source *v20201101s.RouteTablesRoute) error {

	// ObjectMeta
	route.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RouteTables_Route_Spec
	err := spec.AssignProperties_From_RouteTables_Route_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_RouteTables_Route_Spec() to populate field Spec")
	}
	route.Spec = spec

	// Status
	var status RouteTables_Route_STATUS
	err = status.AssignProperties_From_RouteTables_Route_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_RouteTables_Route_STATUS() to populate field Status")
	}
	route.Status = status

	// Invoke the augmentConversionForRouteTablesRoute interface (if implemented) to customize the conversion
	var routeAsAny any = route
	if augmentedRoute, ok := routeAsAny.(augmentConversionForRouteTablesRoute); ok {
		err := augmentedRoute.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_RouteTablesRoute populates the provided destination RouteTablesRoute from our RouteTablesRoute
func (route *RouteTablesRoute) AssignProperties_To_RouteTablesRoute(destination *v20201101s.RouteTablesRoute) error {

	// ObjectMeta
	destination.ObjectMeta = *route.ObjectMeta.DeepCopy()

	// Spec
	var spec v20201101s.RouteTables_Route_Spec
	err := route.Spec.AssignProperties_To_RouteTables_Route_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_RouteTables_Route_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20201101s.RouteTables_Route_STATUS
	err = route.Status.AssignProperties_To_RouteTables_Route_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_RouteTables_Route_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForRouteTablesRoute interface (if implemented) to customize the conversion
	var routeAsAny any = route
	if augmentedRoute, ok := routeAsAny.(augmentConversionForRouteTablesRoute); ok {
		err := augmentedRoute.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (route *RouteTablesRoute) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: route.Spec.OriginalVersion,
		Kind:    "RouteTablesRoute",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20201101.RouteTablesRoute
// Deprecated version of RouteTablesRoute. Use v1api20201101.RouteTablesRoute instead
type RouteTablesRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTablesRoute `json:"items"`
}

type augmentConversionForRouteTablesRoute interface {
	AssignPropertiesFrom(src *v20201101s.RouteTablesRoute) error
	AssignPropertiesTo(dst *v20201101s.RouteTablesRoute) error
}

// Storage version of v1beta20201101.RouteTables_Route_Spec
type RouteTables_Route_Spec struct {
	AddressPrefix *string `json:"addressPrefix,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName        string  `json:"azureName,omitempty"`
	HasBgpOverride   *bool   `json:"hasBgpOverride,omitempty"`
	NextHopIpAddress *string `json:"nextHopIpAddress,omitempty"`
	NextHopType      *string `json:"nextHopType,omitempty"`
	OriginalVersion  string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a network.azure.com/RouteTable resource
	Owner       *genruntime.KnownResourceReference `group:"network.azure.com" json:"owner,omitempty" kind:"RouteTable"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

var _ genruntime.ConvertibleSpec = &RouteTables_Route_Spec{}

// ConvertSpecFrom populates our RouteTables_Route_Spec from the provided source
func (route *RouteTables_Route_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20201101s.RouteTables_Route_Spec)
	if ok {
		// Populate our instance from source
		return route.AssignProperties_From_RouteTables_Route_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20201101s.RouteTables_Route_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = route.AssignProperties_From_RouteTables_Route_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RouteTables_Route_Spec
func (route *RouteTables_Route_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20201101s.RouteTables_Route_Spec)
	if ok {
		// Populate destination from our instance
		return route.AssignProperties_To_RouteTables_Route_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20201101s.RouteTables_Route_Spec{}
	err := route.AssignProperties_To_RouteTables_Route_Spec(dst)
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

// AssignProperties_From_RouteTables_Route_Spec populates our RouteTables_Route_Spec from the provided source RouteTables_Route_Spec
func (route *RouteTables_Route_Spec) AssignProperties_From_RouteTables_Route_Spec(source *v20201101s.RouteTables_Route_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

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
	route.NextHopType = genruntime.ClonePointerToString(source.NextHopType)

	// OriginalVersion
	route.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		route.Owner = &owner
	} else {
		route.Owner = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		route.PropertyBag = propertyBag
	} else {
		route.PropertyBag = nil
	}

	// Invoke the augmentConversionForRouteTables_Route_Spec interface (if implemented) to customize the conversion
	var routeAsAny any = route
	if augmentedRoute, ok := routeAsAny.(augmentConversionForRouteTables_Route_Spec); ok {
		err := augmentedRoute.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_RouteTables_Route_Spec populates the provided destination RouteTables_Route_Spec from our RouteTables_Route_Spec
func (route *RouteTables_Route_Spec) AssignProperties_To_RouteTables_Route_Spec(destination *v20201101s.RouteTables_Route_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(route.PropertyBag)

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
	destination.NextHopType = genruntime.ClonePointerToString(route.NextHopType)

	// OriginalVersion
	destination.OriginalVersion = route.OriginalVersion

	// Owner
	if route.Owner != nil {
		owner := route.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForRouteTables_Route_Spec interface (if implemented) to customize the conversion
	var routeAsAny any = route
	if augmentedRoute, ok := routeAsAny.(augmentConversionForRouteTables_Route_Spec); ok {
		err := augmentedRoute.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1beta20201101.RouteTables_Route_STATUS
// Deprecated version of RouteTables_Route_STATUS. Use v1api20201101.RouteTables_Route_STATUS instead
type RouteTables_Route_STATUS struct {
	AddressPrefix     *string                `json:"addressPrefix,omitempty"`
	Conditions        []conditions.Condition `json:"conditions,omitempty"`
	Etag              *string                `json:"etag,omitempty"`
	HasBgpOverride    *bool                  `json:"hasBgpOverride,omitempty"`
	Id                *string                `json:"id,omitempty"`
	Name              *string                `json:"name,omitempty"`
	NextHopIpAddress  *string                `json:"nextHopIpAddress,omitempty"`
	NextHopType       *string                `json:"nextHopType,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	Type              *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RouteTables_Route_STATUS{}

// ConvertStatusFrom populates our RouteTables_Route_STATUS from the provided source
func (route *RouteTables_Route_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20201101s.RouteTables_Route_STATUS)
	if ok {
		// Populate our instance from source
		return route.AssignProperties_From_RouteTables_Route_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20201101s.RouteTables_Route_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = route.AssignProperties_From_RouteTables_Route_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our RouteTables_Route_STATUS
func (route *RouteTables_Route_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20201101s.RouteTables_Route_STATUS)
	if ok {
		// Populate destination from our instance
		return route.AssignProperties_To_RouteTables_Route_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20201101s.RouteTables_Route_STATUS{}
	err := route.AssignProperties_To_RouteTables_Route_STATUS(dst)
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

// AssignProperties_From_RouteTables_Route_STATUS populates our RouteTables_Route_STATUS from the provided source RouteTables_Route_STATUS
func (route *RouteTables_Route_STATUS) AssignProperties_From_RouteTables_Route_STATUS(source *v20201101s.RouteTables_Route_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

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
	route.NextHopType = genruntime.ClonePointerToString(source.NextHopType)

	// ProvisioningState
	route.ProvisioningState = genruntime.ClonePointerToString(source.ProvisioningState)

	// Type
	route.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		route.PropertyBag = propertyBag
	} else {
		route.PropertyBag = nil
	}

	// Invoke the augmentConversionForRouteTables_Route_STATUS interface (if implemented) to customize the conversion
	var routeAsAny any = route
	if augmentedRoute, ok := routeAsAny.(augmentConversionForRouteTables_Route_STATUS); ok {
		err := augmentedRoute.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_RouteTables_Route_STATUS populates the provided destination RouteTables_Route_STATUS from our RouteTables_Route_STATUS
func (route *RouteTables_Route_STATUS) AssignProperties_To_RouteTables_Route_STATUS(destination *v20201101s.RouteTables_Route_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(route.PropertyBag)

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
	destination.NextHopType = genruntime.ClonePointerToString(route.NextHopType)

	// ProvisioningState
	destination.ProvisioningState = genruntime.ClonePointerToString(route.ProvisioningState)

	// Type
	destination.Type = genruntime.ClonePointerToString(route.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForRouteTables_Route_STATUS interface (if implemented) to customize the conversion
	var routeAsAny any = route
	if augmentedRoute, ok := routeAsAny.(augmentConversionForRouteTables_Route_STATUS); ok {
		err := augmentedRoute.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForRouteTables_Route_Spec interface {
	AssignPropertiesFrom(src *v20201101s.RouteTables_Route_Spec) error
	AssignPropertiesTo(dst *v20201101s.RouteTables_Route_Spec) error
}

type augmentConversionForRouteTables_Route_STATUS interface {
	AssignPropertiesFrom(src *v20201101s.RouteTables_Route_STATUS) error
	AssignPropertiesTo(dst *v20201101s.RouteTables_Route_STATUS) error
}

func init() {
	SchemeBuilder.Register(&RouteTablesRoute{}, &RouteTablesRouteList{})
}
