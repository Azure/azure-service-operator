// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=network.azure.com,resources=routetables,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={routetables/status,routetables/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20201101.RouteTable
// Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/routeTables
type RouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteTable_Spec   `json:"spec,omitempty"`
	Status            RouteTable_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RouteTable{}

// GetConditions returns the conditions of the resource
func (table *RouteTable) GetConditions() conditions.Conditions {
	return table.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (table *RouteTable) SetConditions(conditions conditions.Conditions) {
	table.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &RouteTable{}

// AzureName returns the Azure name of the resource
func (table *RouteTable) AzureName() string {
	return table.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (table RouteTable) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (table *RouteTable) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (table *RouteTable) GetSpec() genruntime.ConvertibleSpec {
	return &table.Spec
}

// GetStatus returns the status of this resource
func (table *RouteTable) GetStatus() genruntime.ConvertibleStatus {
	return &table.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/routeTables"
func (table *RouteTable) GetType() string {
	return "Microsoft.Network/routeTables"
}

// NewEmptyStatus returns a new empty (blank) status
func (table *RouteTable) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RouteTable_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (table *RouteTable) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(table.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  table.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (table *RouteTable) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RouteTable_STATUS); ok {
		table.Status = *st
		return nil
	}

	// Convert status to required version
	var st RouteTable_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	table.Status = st
	return nil
}

// Hub marks that this RouteTable is the hub type for conversion
func (table *RouteTable) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (table *RouteTable) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: table.Spec.OriginalVersion,
		Kind:    "RouteTable",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20201101.RouteTable
// Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/routeTables
type RouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTable `json:"items"`
}

// Storage version of v1beta20201101.RouteTable_Spec
type RouteTable_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                  string  `json:"azureName,omitempty"`
	DisableBgpRoutePropagation *bool   `json:"disableBgpRoutePropagation,omitempty"`
	Location                   *string `json:"location,omitempty"`
	OriginalVersion            string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &RouteTable_Spec{}

// ConvertSpecFrom populates our RouteTable_Spec from the provided source
func (table *RouteTable_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == table {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(table)
}

// ConvertSpecTo populates the provided destination from our RouteTable_Spec
func (table *RouteTable_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == table {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(table)
}

// Storage version of v1beta20201101.RouteTable_STATUS
type RouteTable_STATUS struct {
	Conditions                 []conditions.Condition `json:"conditions,omitempty"`
	DisableBgpRoutePropagation *bool                  `json:"disableBgpRoutePropagation,omitempty"`
	Etag                       *string                `json:"etag,omitempty"`
	Id                         *string                `json:"id,omitempty"`
	Location                   *string                `json:"location,omitempty"`
	Name                       *string                `json:"name,omitempty"`
	PropertyBag                genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState          *string                `json:"provisioningState,omitempty"`
	ResourceGuid               *string                `json:"resourceGuid,omitempty"`
	Tags                       map[string]string      `json:"tags,omitempty"`
	Type                       *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RouteTable_STATUS{}

// ConvertStatusFrom populates our RouteTable_STATUS from the provided source
func (table *RouteTable_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == table {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(table)
}

// ConvertStatusTo populates the provided destination from our RouteTable_STATUS
func (table *RouteTable_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == table {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(table)
}

func init() {
	SchemeBuilder.Register(&RouteTable{}, &RouteTableList{})
}
