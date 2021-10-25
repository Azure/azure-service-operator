// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/api/microsoft.network/v1alpha1api20201101storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:rbac:groups=microsoft.network.azure.com,resources=routetablesroutes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=microsoft.network.azure.com,resources={routetablesroutes/status,routetablesroutes/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/routeTables_routes
type RouteTablesRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteTablesRoutes_Spec `json:"spec,omitempty"`
	Status            Route_Status           `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RouteTablesRoute{}

// GetConditions returns the conditions of the resource
func (routeTablesRoute *RouteTablesRoute) GetConditions() conditions.Conditions {
	return routeTablesRoute.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (routeTablesRoute *RouteTablesRoute) SetConditions(conditions conditions.Conditions) {
	routeTablesRoute.Status.Conditions = conditions
}

// +kubebuilder:webhook:path=/mutate-microsoft-network-azure-com-v1alpha1api20201101-routetablesroute,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=microsoft.network.azure.com,resources=routetablesroutes,verbs=create;update,versions=v1alpha1api20201101,name=default.v1alpha1api20201101.routetablesroutes.microsoft.network.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &RouteTablesRoute{}

// Default applies defaults to the RouteTablesRoute resource
func (routeTablesRoute *RouteTablesRoute) Default() {
	routeTablesRoute.defaultImpl()
	var temp interface{} = routeTablesRoute
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (routeTablesRoute *RouteTablesRoute) defaultAzureName() {
	if routeTablesRoute.Spec.AzureName == "" {
		routeTablesRoute.Spec.AzureName = routeTablesRoute.Name
	}
}

// defaultImpl applies the code generated defaults to the RouteTablesRoute resource
func (routeTablesRoute *RouteTablesRoute) defaultImpl() { routeTablesRoute.defaultAzureName() }

var _ genruntime.KubernetesResource = &RouteTablesRoute{}

// AzureName returns the Azure name of the resource
func (routeTablesRoute *RouteTablesRoute) AzureName() string {
	return routeTablesRoute.Spec.AzureName
}

// GetResourceKind returns the kind of the resource
func (routeTablesRoute *RouteTablesRoute) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (routeTablesRoute *RouteTablesRoute) GetSpec() genruntime.ConvertibleSpec {
	return &routeTablesRoute.Spec
}

// GetStatus returns the status of this resource
func (routeTablesRoute *RouteTablesRoute) GetStatus() genruntime.ConvertibleStatus {
	return &routeTablesRoute.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/routeTables/routes"
func (routeTablesRoute *RouteTablesRoute) GetType() string {
	return "Microsoft.Network/routeTables/routes"
}

// NewEmptyStatus returns a new empty (blank) status
func (routeTablesRoute *RouteTablesRoute) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Route_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (routeTablesRoute *RouteTablesRoute) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(routeTablesRoute.Spec)
	return &genruntime.ResourceReference{
		Group:     group,
		Kind:      kind,
		Namespace: routeTablesRoute.Namespace,
		Name:      routeTablesRoute.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (routeTablesRoute *RouteTablesRoute) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Route_Status); ok {
		routeTablesRoute.Status = *st
		return nil
	}

	// Convert status to required version
	var st Route_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	routeTablesRoute.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-microsoft-network-azure-com-v1alpha1api20201101-routetablesroute,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=microsoft.network.azure.com,resources=routetablesroutes,verbs=create;update,versions=v1alpha1api20201101,name=validate.v1alpha1api20201101.routetablesroutes.microsoft.network.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &RouteTablesRoute{}

// ValidateCreate validates the creation of the resource
func (routeTablesRoute *RouteTablesRoute) ValidateCreate() error {
	validations := routeTablesRoute.createValidations()
	var temp interface{} = routeTablesRoute
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
func (routeTablesRoute *RouteTablesRoute) ValidateDelete() error {
	validations := routeTablesRoute.deleteValidations()
	var temp interface{} = routeTablesRoute
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
func (routeTablesRoute *RouteTablesRoute) ValidateUpdate(old runtime.Object) error {
	validations := routeTablesRoute.updateValidations()
	var temp interface{} = routeTablesRoute
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
func (routeTablesRoute *RouteTablesRoute) createValidations() []func() error {
	return []func() error{routeTablesRoute.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (routeTablesRoute *RouteTablesRoute) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (routeTablesRoute *RouteTablesRoute) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return routeTablesRoute.validateResourceReferences()
		},
	}
}

// validateResourceReferences validates all resource references
func (routeTablesRoute *RouteTablesRoute) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&routeTablesRoute.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromRouteTablesRoute populates our RouteTablesRoute from the provided source RouteTablesRoute
func (routeTablesRoute *RouteTablesRoute) AssignPropertiesFromRouteTablesRoute(source *v1alpha1api20201101storage.RouteTablesRoute) error {

	// Spec
	var spec RouteTablesRoutes_Spec
	err := spec.AssignPropertiesFromRouteTablesRoutesSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "populating Spec from Spec, calling AssignPropertiesFromRouteTablesRoutesSpec()")
	}
	routeTablesRoute.Spec = spec

	// Status
	var status Route_Status
	err = status.AssignPropertiesFromRouteStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "populating Status from Status, calling AssignPropertiesFromRouteStatus()")
	}
	routeTablesRoute.Status = status

	// No error
	return nil
}

// AssignPropertiesToRouteTablesRoute populates the provided destination RouteTablesRoute from our RouteTablesRoute
func (routeTablesRoute *RouteTablesRoute) AssignPropertiesToRouteTablesRoute(destination *v1alpha1api20201101storage.RouteTablesRoute) error {

	// Spec
	var spec v1alpha1api20201101storage.RouteTablesRoutes_Spec
	err := routeTablesRoute.Spec.AssignPropertiesToRouteTablesRoutesSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "populating Spec from Spec, calling AssignPropertiesToRouteTablesRoutesSpec()")
	}
	destination.Spec = spec

	// Status
	var status v1alpha1api20201101storage.Route_Status
	err = routeTablesRoute.Status.AssignPropertiesToRouteStatus(&status)
	if err != nil {
		return errors.Wrap(err, "populating Status from Status, calling AssignPropertiesToRouteStatus()")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (routeTablesRoute *RouteTablesRoute) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: routeTablesRoute.Spec.OriginalVersion(),
		Kind:    "RouteTablesRoute",
	}
}

// +kubebuilder:object:root=true
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/routeTables_routes
type RouteTablesRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTablesRoute `json:"items"`
}

// +kubebuilder:validation:Enum={"2020-11-01"}
type RouteTablesRoutesSpecAPIVersion string

const RouteTablesRoutesSpecAPIVersion20201101 = RouteTablesRoutesSpecAPIVersion("2020-11-01")

type RouteTablesRoutes_Spec struct {
	// +kubebuilder:validation:Required
	//AddressPrefix: The destination CIDR to which the route applies.
	AddressPrefix string `json:"addressPrefix"`

	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName string `json:"azureName"`

	//HasBgpOverride: A value indicating whether this route overrides overlapping BGP
	//routes regardless of LPM.
	HasBgpOverride *bool `json:"hasBgpOverride,omitempty"`

	//Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	//NextHopIpAddress: The IP address packets should be forwarded to. Next hop values
	//are only allowed in routes where the next hop type is VirtualAppliance.
	NextHopIpAddress *string `json:"nextHopIpAddress,omitempty"`

	// +kubebuilder:validation:Required
	//NextHopType: The type of Azure hop the packet should be sent to.
	NextHopType RoutePropertiesFormatNextHopType `json:"nextHopType"`

	// +kubebuilder:validation:Required
	Owner genruntime.KnownResourceReference `group:"microsoft.network.azure.com" json:"owner" kind:"RouteTable"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &RouteTablesRoutes_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (routeTablesRoutesSpec *RouteTablesRoutes_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if routeTablesRoutesSpec == nil {
		return nil, nil
	}
	var result RouteTablesRoutes_SpecARM

	// Set property ‘Location’:
	if routeTablesRoutesSpec.Location != nil {
		location := *routeTablesRoutesSpec.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	result.Properties.AddressPrefix = routeTablesRoutesSpec.AddressPrefix
	if routeTablesRoutesSpec.HasBgpOverride != nil {
		hasBgpOverride := *routeTablesRoutesSpec.HasBgpOverride
		result.Properties.HasBgpOverride = &hasBgpOverride
	}
	if routeTablesRoutesSpec.NextHopIpAddress != nil {
		nextHopIpAddress := *routeTablesRoutesSpec.NextHopIpAddress
		result.Properties.NextHopIpAddress = &nextHopIpAddress
	}
	result.Properties.NextHopType = routeTablesRoutesSpec.NextHopType

	// Set property ‘Tags’:
	if routeTablesRoutesSpec.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range routeTablesRoutesSpec.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (routeTablesRoutesSpec *RouteTablesRoutes_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RouteTablesRoutes_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (routeTablesRoutesSpec *RouteTablesRoutes_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RouteTablesRoutes_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RouteTablesRoutes_SpecARM, got %T", armInput)
	}

	// Set property ‘AddressPrefix’:
	// copying flattened property:
	routeTablesRoutesSpec.AddressPrefix = typedInput.Properties.AddressPrefix

	// Set property ‘AzureName’:
	routeTablesRoutesSpec.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘HasBgpOverride’:
	// copying flattened property:
	if typedInput.Properties.HasBgpOverride != nil {
		hasBgpOverride := *typedInput.Properties.HasBgpOverride
		routeTablesRoutesSpec.HasBgpOverride = &hasBgpOverride
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		routeTablesRoutesSpec.Location = &location
	}

	// Set property ‘NextHopIpAddress’:
	// copying flattened property:
	if typedInput.Properties.NextHopIpAddress != nil {
		nextHopIpAddress := *typedInput.Properties.NextHopIpAddress
		routeTablesRoutesSpec.NextHopIpAddress = &nextHopIpAddress
	}

	// Set property ‘NextHopType’:
	// copying flattened property:
	routeTablesRoutesSpec.NextHopType = typedInput.Properties.NextHopType

	// Set property ‘Owner’:
	routeTablesRoutesSpec.Owner = genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		routeTablesRoutesSpec.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			routeTablesRoutesSpec.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &RouteTablesRoutes_Spec{}

// ConvertSpecFrom populates our RouteTablesRoutes_Spec from the provided source
func (routeTablesRoutesSpec *RouteTablesRoutes_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1alpha1api20201101storage.RouteTablesRoutes_Spec)
	if ok {
		// Populate our instance from source
		return routeTablesRoutesSpec.AssignPropertiesFromRouteTablesRoutesSpec(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20201101storage.RouteTablesRoutes_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = routeTablesRoutesSpec.AssignPropertiesFromRouteTablesRoutesSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RouteTablesRoutes_Spec
func (routeTablesRoutesSpec *RouteTablesRoutes_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1alpha1api20201101storage.RouteTablesRoutes_Spec)
	if ok {
		// Populate destination from our instance
		return routeTablesRoutesSpec.AssignPropertiesToRouteTablesRoutesSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20201101storage.RouteTablesRoutes_Spec{}
	err := routeTablesRoutesSpec.AssignPropertiesToRouteTablesRoutesSpec(dst)
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

// AssignPropertiesFromRouteTablesRoutesSpec populates our RouteTablesRoutes_Spec from the provided source RouteTablesRoutes_Spec
func (routeTablesRoutesSpec *RouteTablesRoutes_Spec) AssignPropertiesFromRouteTablesRoutesSpec(source *v1alpha1api20201101storage.RouteTablesRoutes_Spec) error {

	// AddressPrefix
	routeTablesRoutesSpec.AddressPrefix = genruntime.GetOptionalStringValue(source.AddressPrefix)

	// AzureName
	routeTablesRoutesSpec.AzureName = source.AzureName

	// HasBgpOverride
	if source.HasBgpOverride != nil {
		hasBgpOverride := *source.HasBgpOverride
		routeTablesRoutesSpec.HasBgpOverride = &hasBgpOverride
	} else {
		routeTablesRoutesSpec.HasBgpOverride = nil
	}

	// Location
	routeTablesRoutesSpec.Location = genruntime.ClonePointerToString(source.Location)

	// NextHopIpAddress
	routeTablesRoutesSpec.NextHopIpAddress = genruntime.ClonePointerToString(source.NextHopIpAddress)

	// NextHopType
	if source.NextHopType != nil {
		routeTablesRoutesSpec.NextHopType = RoutePropertiesFormatNextHopType(*source.NextHopType)
	} else {
		routeTablesRoutesSpec.NextHopType = ""
	}

	// Owner
	routeTablesRoutesSpec.Owner = source.Owner.Copy()

	// Tags
	routeTablesRoutesSpec.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToRouteTablesRoutesSpec populates the provided destination RouteTablesRoutes_Spec from our RouteTablesRoutes_Spec
func (routeTablesRoutesSpec *RouteTablesRoutes_Spec) AssignPropertiesToRouteTablesRoutesSpec(destination *v1alpha1api20201101storage.RouteTablesRoutes_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AddressPrefix
	addressPrefix := routeTablesRoutesSpec.AddressPrefix
	destination.AddressPrefix = &addressPrefix

	// AzureName
	destination.AzureName = routeTablesRoutesSpec.AzureName

	// HasBgpOverride
	if routeTablesRoutesSpec.HasBgpOverride != nil {
		hasBgpOverride := *routeTablesRoutesSpec.HasBgpOverride
		destination.HasBgpOverride = &hasBgpOverride
	} else {
		destination.HasBgpOverride = nil
	}

	// Location
	destination.Location = genruntime.ClonePointerToString(routeTablesRoutesSpec.Location)

	// NextHopIpAddress
	destination.NextHopIpAddress = genruntime.ClonePointerToString(routeTablesRoutesSpec.NextHopIpAddress)

	// NextHopType
	nextHopType := string(routeTablesRoutesSpec.NextHopType)
	destination.NextHopType = &nextHopType

	// OriginalVersion
	destination.OriginalVersion = routeTablesRoutesSpec.OriginalVersion()

	// Owner
	destination.Owner = routeTablesRoutesSpec.Owner.Copy()

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(routeTablesRoutesSpec.Tags)

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
func (routeTablesRoutesSpec *RouteTablesRoutes_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (routeTablesRoutesSpec *RouteTablesRoutes_Spec) SetAzureName(azureName string) {
	routeTablesRoutesSpec.AzureName = azureName
}

//Generated from:
type Route_Status struct {
	//AddressPrefix: The destination CIDR to which the route applies.
	AddressPrefix *string `json:"addressPrefix,omitempty"`

	//Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	//Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	//HasBgpOverride: A value indicating whether this route overrides overlapping BGP
	//routes regardless of LPM.
	HasBgpOverride *bool `json:"hasBgpOverride,omitempty"`

	//Id: Resource ID.
	Id *string `json:"id,omitempty"`

	//Name: The name of the resource that is unique within a resource group. This name
	//can be used to access the resource.
	Name *string `json:"name,omitempty"`

	//NextHopIpAddress: The IP address packets should be forwarded to. Next hop values
	//are only allowed in routes where the next hop type is VirtualAppliance.
	NextHopIpAddress *string `json:"nextHopIpAddress,omitempty"`

	//NextHopType: The type of Azure hop the packet should be sent to.
	NextHopType *RouteNextHopType_Status `json:"nextHopType,omitempty"`

	//ProvisioningState: The provisioning state of the route resource.
	ProvisioningState *ProvisioningState_Status `json:"provisioningState,omitempty"`

	//Type: The type of the resource.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Route_Status{}

// ConvertStatusFrom populates our Route_Status from the provided source
func (routeStatus *Route_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1alpha1api20201101storage.Route_Status)
	if ok {
		// Populate our instance from source
		return routeStatus.AssignPropertiesFromRouteStatus(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20201101storage.Route_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = routeStatus.AssignPropertiesFromRouteStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Route_Status
func (routeStatus *Route_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1alpha1api20201101storage.Route_Status)
	if ok {
		// Populate destination from our instance
		return routeStatus.AssignPropertiesToRouteStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20201101storage.Route_Status{}
	err := routeStatus.AssignPropertiesToRouteStatus(dst)
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

var _ genruntime.FromARMConverter = &Route_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (routeStatus *Route_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Route_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (routeStatus *Route_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Route_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Route_StatusARM, got %T", armInput)
	}

	// Set property ‘AddressPrefix’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.AddressPrefix != nil {
			addressPrefix := *typedInput.Properties.AddressPrefix
			routeStatus.AddressPrefix = &addressPrefix
		}
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Etag’:
	if typedInput.Etag != nil {
		etag := *typedInput.Etag
		routeStatus.Etag = &etag
	}

	// Set property ‘HasBgpOverride’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.HasBgpOverride != nil {
			hasBgpOverride := *typedInput.Properties.HasBgpOverride
			routeStatus.HasBgpOverride = &hasBgpOverride
		}
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		routeStatus.Id = &id
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		routeStatus.Name = &name
	}

	// Set property ‘NextHopIpAddress’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.NextHopIpAddress != nil {
			nextHopIpAddress := *typedInput.Properties.NextHopIpAddress
			routeStatus.NextHopIpAddress = &nextHopIpAddress
		}
	}

	// Set property ‘NextHopType’:
	// copying flattened property:
	if typedInput.Properties != nil {
		routeStatus.NextHopType = &typedInput.Properties.NextHopType
	}

	// Set property ‘ProvisioningState’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ProvisioningState != nil {
			provisioningState := *typedInput.Properties.ProvisioningState
			routeStatus.ProvisioningState = &provisioningState
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		routeStatus.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromRouteStatus populates our Route_Status from the provided source Route_Status
func (routeStatus *Route_Status) AssignPropertiesFromRouteStatus(source *v1alpha1api20201101storage.Route_Status) error {

	// AddressPrefix
	routeStatus.AddressPrefix = genruntime.ClonePointerToString(source.AddressPrefix)

	// Conditions
	routeStatus.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Etag
	routeStatus.Etag = genruntime.ClonePointerToString(source.Etag)

	// HasBgpOverride
	if source.HasBgpOverride != nil {
		hasBgpOverride := *source.HasBgpOverride
		routeStatus.HasBgpOverride = &hasBgpOverride
	} else {
		routeStatus.HasBgpOverride = nil
	}

	// Id
	routeStatus.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	routeStatus.Name = genruntime.ClonePointerToString(source.Name)

	// NextHopIpAddress
	routeStatus.NextHopIpAddress = genruntime.ClonePointerToString(source.NextHopIpAddress)

	// NextHopType
	if source.NextHopType != nil {
		nextHopType := RouteNextHopType_Status(*source.NextHopType)
		routeStatus.NextHopType = &nextHopType
	} else {
		routeStatus.NextHopType = nil
	}

	// ProvisioningState
	if source.ProvisioningState != nil {
		provisioningState := ProvisioningState_Status(*source.ProvisioningState)
		routeStatus.ProvisioningState = &provisioningState
	} else {
		routeStatus.ProvisioningState = nil
	}

	// Type
	routeStatus.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToRouteStatus populates the provided destination Route_Status from our Route_Status
func (routeStatus *Route_Status) AssignPropertiesToRouteStatus(destination *v1alpha1api20201101storage.Route_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AddressPrefix
	destination.AddressPrefix = genruntime.ClonePointerToString(routeStatus.AddressPrefix)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(routeStatus.Conditions)

	// Etag
	destination.Etag = genruntime.ClonePointerToString(routeStatus.Etag)

	// HasBgpOverride
	if routeStatus.HasBgpOverride != nil {
		hasBgpOverride := *routeStatus.HasBgpOverride
		destination.HasBgpOverride = &hasBgpOverride
	} else {
		destination.HasBgpOverride = nil
	}

	// Id
	destination.Id = genruntime.ClonePointerToString(routeStatus.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(routeStatus.Name)

	// NextHopIpAddress
	destination.NextHopIpAddress = genruntime.ClonePointerToString(routeStatus.NextHopIpAddress)

	// NextHopType
	if routeStatus.NextHopType != nil {
		nextHopType := string(*routeStatus.NextHopType)
		destination.NextHopType = &nextHopType
	} else {
		destination.NextHopType = nil
	}

	// ProvisioningState
	if routeStatus.ProvisioningState != nil {
		provisioningState := string(*routeStatus.ProvisioningState)
		destination.ProvisioningState = &provisioningState
	} else {
		destination.ProvisioningState = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(routeStatus.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

//Generated from:
type RouteNextHopType_Status string

const (
	RouteNextHopType_StatusInternet              = RouteNextHopType_Status("Internet")
	RouteNextHopType_StatusNone                  = RouteNextHopType_Status("None")
	RouteNextHopType_StatusVirtualAppliance      = RouteNextHopType_Status("VirtualAppliance")
	RouteNextHopType_StatusVirtualNetworkGateway = RouteNextHopType_Status("VirtualNetworkGateway")
	RouteNextHopType_StatusVnetLocal             = RouteNextHopType_Status("VnetLocal")
)

// +kubebuilder:validation:Enum={"Internet","None","VirtualAppliance","VirtualNetworkGateway","VnetLocal"}
type RoutePropertiesFormatNextHopType string

const (
	RoutePropertiesFormatNextHopTypeInternet              = RoutePropertiesFormatNextHopType("Internet")
	RoutePropertiesFormatNextHopTypeNone                  = RoutePropertiesFormatNextHopType("None")
	RoutePropertiesFormatNextHopTypeVirtualAppliance      = RoutePropertiesFormatNextHopType("VirtualAppliance")
	RoutePropertiesFormatNextHopTypeVirtualNetworkGateway = RoutePropertiesFormatNextHopType("VirtualNetworkGateway")
	RoutePropertiesFormatNextHopTypeVnetLocal             = RoutePropertiesFormatNextHopType("VnetLocal")
)

func init() {
	SchemeBuilder.Register(&RouteTablesRoute{}, &RouteTablesRouteList{})
}
