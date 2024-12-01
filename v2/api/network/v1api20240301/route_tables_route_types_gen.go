// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20240301

import (
	"fmt"
	arm "github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/arm"
	storage "github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/rotisserie/eris"
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
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2024-03-01/routeTable.json
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
	source, ok := hub.(*storage.RouteTablesRoute)
	if !ok {
		return fmt.Errorf("expected network/v1api20240301/storage/RouteTablesRoute but received %T instead", hub)
	}

	return route.AssignProperties_From_RouteTablesRoute(source)
}

// ConvertTo populates the provided hub RouteTablesRoute from our RouteTablesRoute
func (route *RouteTablesRoute) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.RouteTablesRoute)
	if !ok {
		return fmt.Errorf("expected network/v1api20240301/storage/RouteTablesRoute but received %T instead", hub)
	}

	return route.AssignProperties_To_RouteTablesRoute(destination)
}

// +kubebuilder:webhook:path=/mutate-network-azure-com-v1api20240301-routetablesroute,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=routetablesroutes,verbs=create;update,versions=v1api20240301,name=default.v1api20240301.routetablesroutes.network.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &RouteTablesRoute{}

// Default applies defaults to the RouteTablesRoute resource
func (route *RouteTablesRoute) Default() {
	route.defaultImpl()
	var temp any = route
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

var _ configmaps.Exporter = &RouteTablesRoute{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (route *RouteTablesRoute) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if route.Spec.OperatorSpec == nil {
		return nil
	}
	return route.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &RouteTablesRoute{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (route *RouteTablesRoute) SecretDestinationExpressions() []*core.DestinationExpression {
	if route.Spec.OperatorSpec == nil {
		return nil
	}
	return route.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.ImportableResource = &RouteTablesRoute{}

// InitializeSpec initializes the spec for this resource from the given status
func (route *RouteTablesRoute) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*RouteTablesRoute_STATUS); ok {
		return route.Spec.Initialize_From_RouteTablesRoute_STATUS(s)
	}

	return fmt.Errorf("expected Status of type RouteTablesRoute_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &RouteTablesRoute{}

// AzureName returns the Azure name of the resource
func (route *RouteTablesRoute) AzureName() string {
	return route.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-03-01"
func (route RouteTablesRoute) GetAPIVersion() string {
	return "2024-03-01"
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

// GetSupportedOperations returns the operations supported by the resource
func (route *RouteTablesRoute) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/routeTables/routes"
func (route *RouteTablesRoute) GetType() string {
	return "Microsoft.Network/routeTables/routes"
}

// NewEmptyStatus returns a new empty (blank) status
func (route *RouteTablesRoute) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RouteTablesRoute_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (route *RouteTablesRoute) Owner() *genruntime.ResourceReference {
	if route.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(route.Spec)
	return route.Spec.Owner.AsResourceReference(group, kind)
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
		return eris.Wrap(err, "failed to convert status")
	}

	route.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-network-azure-com-v1api20240301-routetablesroute,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=network.azure.com,resources=routetablesroutes,verbs=create;update,versions=v1api20240301,name=validate.v1api20240301.routetablesroutes.network.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &RouteTablesRoute{}

// ValidateCreate validates the creation of the resource
func (route *RouteTablesRoute) ValidateCreate() (admission.Warnings, error) {
	validations := route.createValidations()
	var temp any = route
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (route *RouteTablesRoute) ValidateDelete() (admission.Warnings, error) {
	validations := route.deleteValidations()
	var temp any = route
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (route *RouteTablesRoute) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := route.updateValidations()
	var temp any = route
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (route *RouteTablesRoute) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){route.validateResourceReferences, route.validateOwnerReference, route.validateSecretDestinations, route.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (route *RouteTablesRoute) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (route *RouteTablesRoute) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return route.validateResourceReferences()
		},
		route.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return route.validateOwnerReference()
		},
		func(old runtime.Object) (admission.Warnings, error) {
			return route.validateSecretDestinations()
		},
		func(old runtime.Object) (admission.Warnings, error) {
			return route.validateConfigMapDestinations()
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (route *RouteTablesRoute) validateConfigMapDestinations() (admission.Warnings, error) {
	if route.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(route, nil, route.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (route *RouteTablesRoute) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(route)
}

// validateResourceReferences validates all resource references
func (route *RouteTablesRoute) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&route.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (route *RouteTablesRoute) validateSecretDestinations() (admission.Warnings, error) {
	if route.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(route, nil, route.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (route *RouteTablesRoute) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*RouteTablesRoute)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, route)
}

// AssignProperties_From_RouteTablesRoute populates our RouteTablesRoute from the provided source RouteTablesRoute
func (route *RouteTablesRoute) AssignProperties_From_RouteTablesRoute(source *storage.RouteTablesRoute) error {

	// ObjectMeta
	route.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RouteTablesRoute_Spec
	err := spec.AssignProperties_From_RouteTablesRoute_Spec(&source.Spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_RouteTablesRoute_Spec() to populate field Spec")
	}
	route.Spec = spec

	// Status
	var status RouteTablesRoute_STATUS
	err = status.AssignProperties_From_RouteTablesRoute_STATUS(&source.Status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_RouteTablesRoute_STATUS() to populate field Status")
	}
	route.Status = status

	// No error
	return nil
}

// AssignProperties_To_RouteTablesRoute populates the provided destination RouteTablesRoute from our RouteTablesRoute
func (route *RouteTablesRoute) AssignProperties_To_RouteTablesRoute(destination *storage.RouteTablesRoute) error {

	// ObjectMeta
	destination.ObjectMeta = *route.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.RouteTablesRoute_Spec
	err := route.Spec.AssignProperties_To_RouteTablesRoute_Spec(&spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_RouteTablesRoute_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.RouteTablesRoute_STATUS
	err = route.Status.AssignProperties_To_RouteTablesRoute_STATUS(&status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_RouteTablesRoute_STATUS() to populate field Status")
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
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2024-03-01/routeTable.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeTables/{routeTableName}/routes/{routeName}
type RouteTablesRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTablesRoute `json:"items"`
}

type RouteTablesRoute_Spec struct {
	// AddressPrefix: The destination CIDR to which the route applies.
	AddressPrefix *string `json:"addressPrefix,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// NextHopIpAddress: The IP address packets should be forwarded to. Next hop values are only allowed in routes where the
	// next hop type is VirtualAppliance.
	NextHopIpAddress *string `json:"nextHopIpAddress,omitempty"`

	// +kubebuilder:validation:Required
	// NextHopType: The type of Azure hop the packet should be sent to.
	NextHopType *RouteNextHopType `json:"nextHopType,omitempty"`

	// OperatorSpec: The specification for configuring operator behavior. This field is interpreted by the operator and not
	// passed directly to Azure
	OperatorSpec *RouteTablesRouteOperatorSpec `json:"operatorSpec,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a network.azure.com/RouteTable resource
	Owner *genruntime.KnownResourceReference `group:"network.azure.com" json:"owner,omitempty" kind:"RouteTable"`
}

var _ genruntime.ARMTransformer = &RouteTablesRoute_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (route *RouteTablesRoute_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if route == nil {
		return nil, nil
	}
	result := &arm.RouteTablesRoute_Spec{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if route.AddressPrefix != nil ||
		route.NextHopIpAddress != nil ||
		route.NextHopType != nil {
		result.Properties = &arm.RoutePropertiesFormat{}
	}
	if route.AddressPrefix != nil {
		addressPrefix := *route.AddressPrefix
		result.Properties.AddressPrefix = &addressPrefix
	}
	if route.NextHopIpAddress != nil {
		nextHopIpAddress := *route.NextHopIpAddress
		result.Properties.NextHopIpAddress = &nextHopIpAddress
	}
	if route.NextHopType != nil {
		var temp string
		temp = string(*route.NextHopType)
		nextHopType := arm.RouteNextHopType(temp)
		result.Properties.NextHopType = &nextHopType
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (route *RouteTablesRoute_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.RouteTablesRoute_Spec{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (route *RouteTablesRoute_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.RouteTablesRoute_Spec)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.RouteTablesRoute_Spec, got %T", armInput)
	}

	// Set property "AddressPrefix":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.AddressPrefix != nil {
			addressPrefix := *typedInput.Properties.AddressPrefix
			route.AddressPrefix = &addressPrefix
		}
	}

	// Set property "AzureName":
	route.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property "NextHopIpAddress":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.NextHopIpAddress != nil {
			nextHopIpAddress := *typedInput.Properties.NextHopIpAddress
			route.NextHopIpAddress = &nextHopIpAddress
		}
	}

	// Set property "NextHopType":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.NextHopType != nil {
			var temp string
			temp = string(*typedInput.Properties.NextHopType)
			nextHopType := RouteNextHopType(temp)
			route.NextHopType = &nextHopType
		}
	}

	// no assignment for property "OperatorSpec"

	// Set property "Owner":
	route.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &RouteTablesRoute_Spec{}

// ConvertSpecFrom populates our RouteTablesRoute_Spec from the provided source
func (route *RouteTablesRoute_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.RouteTablesRoute_Spec)
	if ok {
		// Populate our instance from source
		return route.AssignProperties_From_RouteTablesRoute_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.RouteTablesRoute_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = route.AssignProperties_From_RouteTablesRoute_Spec(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RouteTablesRoute_Spec
func (route *RouteTablesRoute_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.RouteTablesRoute_Spec)
	if ok {
		// Populate destination from our instance
		return route.AssignProperties_To_RouteTablesRoute_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.RouteTablesRoute_Spec{}
	err := route.AssignProperties_To_RouteTablesRoute_Spec(dst)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_RouteTablesRoute_Spec populates our RouteTablesRoute_Spec from the provided source RouteTablesRoute_Spec
func (route *RouteTablesRoute_Spec) AssignProperties_From_RouteTablesRoute_Spec(source *storage.RouteTablesRoute_Spec) error {

	// AddressPrefix
	route.AddressPrefix = genruntime.ClonePointerToString(source.AddressPrefix)

	// AzureName
	route.AzureName = source.AzureName

	// NextHopIpAddress
	route.NextHopIpAddress = genruntime.ClonePointerToString(source.NextHopIpAddress)

	// NextHopType
	if source.NextHopType != nil {
		nextHopType := *source.NextHopType
		nextHopTypeTemp := genruntime.ToEnum(nextHopType, routeNextHopType_Values)
		route.NextHopType = &nextHopTypeTemp
	} else {
		route.NextHopType = nil
	}

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec RouteTablesRouteOperatorSpec
		err := operatorSpec.AssignProperties_From_RouteTablesRouteOperatorSpec(source.OperatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_RouteTablesRouteOperatorSpec() to populate field OperatorSpec")
		}
		route.OperatorSpec = &operatorSpec
	} else {
		route.OperatorSpec = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		route.Owner = &owner
	} else {
		route.Owner = nil
	}

	// No error
	return nil
}

// AssignProperties_To_RouteTablesRoute_Spec populates the provided destination RouteTablesRoute_Spec from our RouteTablesRoute_Spec
func (route *RouteTablesRoute_Spec) AssignProperties_To_RouteTablesRoute_Spec(destination *storage.RouteTablesRoute_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AddressPrefix
	destination.AddressPrefix = genruntime.ClonePointerToString(route.AddressPrefix)

	// AzureName
	destination.AzureName = route.AzureName

	// NextHopIpAddress
	destination.NextHopIpAddress = genruntime.ClonePointerToString(route.NextHopIpAddress)

	// NextHopType
	if route.NextHopType != nil {
		nextHopType := string(*route.NextHopType)
		destination.NextHopType = &nextHopType
	} else {
		destination.NextHopType = nil
	}

	// OperatorSpec
	if route.OperatorSpec != nil {
		var operatorSpec storage.RouteTablesRouteOperatorSpec
		err := route.OperatorSpec.AssignProperties_To_RouteTablesRouteOperatorSpec(&operatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_RouteTablesRouteOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
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

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Initialize_From_RouteTablesRoute_STATUS populates our RouteTablesRoute_Spec from the provided source RouteTablesRoute_STATUS
func (route *RouteTablesRoute_Spec) Initialize_From_RouteTablesRoute_STATUS(source *RouteTablesRoute_STATUS) error {

	// AddressPrefix
	route.AddressPrefix = genruntime.ClonePointerToString(source.AddressPrefix)

	// NextHopIpAddress
	route.NextHopIpAddress = genruntime.ClonePointerToString(source.NextHopIpAddress)

	// NextHopType
	if source.NextHopType != nil {
		nextHopType := genruntime.ToEnum(string(*source.NextHopType), routeNextHopType_Values)
		route.NextHopType = &nextHopType
	} else {
		route.NextHopType = nil
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
	src, ok := source.(*storage.RouteTablesRoute_STATUS)
	if ok {
		// Populate our instance from source
		return route.AssignProperties_From_RouteTablesRoute_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.RouteTablesRoute_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = route.AssignProperties_From_RouteTablesRoute_STATUS(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our RouteTablesRoute_STATUS
func (route *RouteTablesRoute_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.RouteTablesRoute_STATUS)
	if ok {
		// Populate destination from our instance
		return route.AssignProperties_To_RouteTablesRoute_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.RouteTablesRoute_STATUS{}
	err := route.AssignProperties_To_RouteTablesRoute_STATUS(dst)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

var _ genruntime.FromARMConverter = &RouteTablesRoute_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (route *RouteTablesRoute_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.RouteTablesRoute_STATUS{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (route *RouteTablesRoute_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.RouteTablesRoute_STATUS)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.RouteTablesRoute_STATUS, got %T", armInput)
	}

	// Set property "AddressPrefix":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.AddressPrefix != nil {
			addressPrefix := *typedInput.Properties.AddressPrefix
			route.AddressPrefix = &addressPrefix
		}
	}

	// no assignment for property "Conditions"

	// Set property "Etag":
	if typedInput.Etag != nil {
		etag := *typedInput.Etag
		route.Etag = &etag
	}

	// Set property "HasBgpOverride":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.HasBgpOverride != nil {
			hasBgpOverride := *typedInput.Properties.HasBgpOverride
			route.HasBgpOverride = &hasBgpOverride
		}
	}

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		route.Id = &id
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		route.Name = &name
	}

	// Set property "NextHopIpAddress":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.NextHopIpAddress != nil {
			nextHopIpAddress := *typedInput.Properties.NextHopIpAddress
			route.NextHopIpAddress = &nextHopIpAddress
		}
	}

	// Set property "NextHopType":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.NextHopType != nil {
			var temp string
			temp = string(*typedInput.Properties.NextHopType)
			nextHopType := RouteNextHopType_STATUS(temp)
			route.NextHopType = &nextHopType
		}
	}

	// Set property "ProvisioningState":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ProvisioningState != nil {
			var temp string
			temp = string(*typedInput.Properties.ProvisioningState)
			provisioningState := ProvisioningState_STATUS(temp)
			route.ProvisioningState = &provisioningState
		}
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		route.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_RouteTablesRoute_STATUS populates our RouteTablesRoute_STATUS from the provided source RouteTablesRoute_STATUS
func (route *RouteTablesRoute_STATUS) AssignProperties_From_RouteTablesRoute_STATUS(source *storage.RouteTablesRoute_STATUS) error {

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
		nextHopType := *source.NextHopType
		nextHopTypeTemp := genruntime.ToEnum(nextHopType, routeNextHopType_STATUS_Values)
		route.NextHopType = &nextHopTypeTemp
	} else {
		route.NextHopType = nil
	}

	// ProvisioningState
	if source.ProvisioningState != nil {
		provisioningState := *source.ProvisioningState
		provisioningStateTemp := genruntime.ToEnum(provisioningState, provisioningState_STATUS_Values)
		route.ProvisioningState = &provisioningStateTemp
	} else {
		route.ProvisioningState = nil
	}

	// Type
	route.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_RouteTablesRoute_STATUS populates the provided destination RouteTablesRoute_STATUS from our RouteTablesRoute_STATUS
func (route *RouteTablesRoute_STATUS) AssignProperties_To_RouteTablesRoute_STATUS(destination *storage.RouteTablesRoute_STATUS) error {
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

// The type of Azure hop the packet should be sent to.
// +kubebuilder:validation:Enum={"Internet","None","VirtualAppliance","VirtualNetworkGateway","VnetLocal"}
type RouteNextHopType string

const (
	RouteNextHopType_Internet              = RouteNextHopType("Internet")
	RouteNextHopType_None                  = RouteNextHopType("None")
	RouteNextHopType_VirtualAppliance      = RouteNextHopType("VirtualAppliance")
	RouteNextHopType_VirtualNetworkGateway = RouteNextHopType("VirtualNetworkGateway")
	RouteNextHopType_VnetLocal             = RouteNextHopType("VnetLocal")
)

// Mapping from string to RouteNextHopType
var routeNextHopType_Values = map[string]RouteNextHopType{
	"internet":              RouteNextHopType_Internet,
	"none":                  RouteNextHopType_None,
	"virtualappliance":      RouteNextHopType_VirtualAppliance,
	"virtualnetworkgateway": RouteNextHopType_VirtualNetworkGateway,
	"vnetlocal":             RouteNextHopType_VnetLocal,
}

// The type of Azure hop the packet should be sent to.
type RouteNextHopType_STATUS string

const (
	RouteNextHopType_STATUS_Internet              = RouteNextHopType_STATUS("Internet")
	RouteNextHopType_STATUS_None                  = RouteNextHopType_STATUS("None")
	RouteNextHopType_STATUS_VirtualAppliance      = RouteNextHopType_STATUS("VirtualAppliance")
	RouteNextHopType_STATUS_VirtualNetworkGateway = RouteNextHopType_STATUS("VirtualNetworkGateway")
	RouteNextHopType_STATUS_VnetLocal             = RouteNextHopType_STATUS("VnetLocal")
)

// Mapping from string to RouteNextHopType_STATUS
var routeNextHopType_STATUS_Values = map[string]RouteNextHopType_STATUS{
	"internet":              RouteNextHopType_STATUS_Internet,
	"none":                  RouteNextHopType_STATUS_None,
	"virtualappliance":      RouteNextHopType_STATUS_VirtualAppliance,
	"virtualnetworkgateway": RouteNextHopType_STATUS_VirtualNetworkGateway,
	"vnetlocal":             RouteNextHopType_STATUS_VnetLocal,
}

// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type RouteTablesRouteOperatorSpec struct {
	// ConfigMapExpressions: configures where to place operator written dynamic ConfigMaps (created with CEL expressions).
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`

	// SecretExpressions: configures where to place operator written dynamic secrets (created with CEL expressions).
	SecretExpressions []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_RouteTablesRouteOperatorSpec populates our RouteTablesRouteOperatorSpec from the provided source RouteTablesRouteOperatorSpec
func (operator *RouteTablesRouteOperatorSpec) AssignProperties_From_RouteTablesRouteOperatorSpec(source *storage.RouteTablesRouteOperatorSpec) error {

	// ConfigMapExpressions
	if source.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(source.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range source.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		operator.ConfigMapExpressions = configMapExpressionList
	} else {
		operator.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if source.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(source.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range source.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		operator.SecretExpressions = secretExpressionList
	} else {
		operator.SecretExpressions = nil
	}

	// No error
	return nil
}

// AssignProperties_To_RouteTablesRouteOperatorSpec populates the provided destination RouteTablesRouteOperatorSpec from our RouteTablesRouteOperatorSpec
func (operator *RouteTablesRouteOperatorSpec) AssignProperties_To_RouteTablesRouteOperatorSpec(destination *storage.RouteTablesRouteOperatorSpec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// ConfigMapExpressions
	if operator.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(operator.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range operator.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		destination.ConfigMapExpressions = configMapExpressionList
	} else {
		destination.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if operator.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(operator.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range operator.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		destination.SecretExpressions = secretExpressionList
	} else {
		destination.SecretExpressions = nil
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
	SchemeBuilder.Register(&RouteTablesRoute{}, &RouteTablesRouteList{})
}
