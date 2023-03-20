// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200601storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=network.azure.com,resources=privatednszonesrecordsettypesrvs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={privatednszonesrecordsettypesrvs/status,privatednszonesrecordsettypesrvs/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20200601.PrivateDnsZonesRecordsetTypeSRV
// Generator information:
// - Generated from: /privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/privatedns.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/SRV/{relativeRecordSetName}
type PrivateDnsZonesRecordsetTypeSRV struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateDnsZones_SRV_Spec   `json:"spec,omitempty"`
	Status            PrivateDnsZones_SRV_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &PrivateDnsZonesRecordsetTypeSRV{}

// GetConditions returns the conditions of the resource
func (typeSRV *PrivateDnsZonesRecordsetTypeSRV) GetConditions() conditions.Conditions {
	return typeSRV.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (typeSRV *PrivateDnsZonesRecordsetTypeSRV) SetConditions(conditions conditions.Conditions) {
	typeSRV.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &PrivateDnsZonesRecordsetTypeSRV{}

// AzureName returns the Azure name of the resource
func (typeSRV *PrivateDnsZonesRecordsetTypeSRV) AzureName() string {
	return typeSRV.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (typeSRV PrivateDnsZonesRecordsetTypeSRV) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (typeSRV *PrivateDnsZonesRecordsetTypeSRV) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (typeSRV *PrivateDnsZonesRecordsetTypeSRV) GetSpec() genruntime.ConvertibleSpec {
	return &typeSRV.Spec
}

// GetStatus returns the status of this resource
func (typeSRV *PrivateDnsZonesRecordsetTypeSRV) GetStatus() genruntime.ConvertibleStatus {
	return &typeSRV.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/privateDnsZones/SRV"
func (typeSRV *PrivateDnsZonesRecordsetTypeSRV) GetType() string {
	return "Microsoft.Network/privateDnsZones/SRV"
}

// NewEmptyStatus returns a new empty (blank) status
func (typeSRV *PrivateDnsZonesRecordsetTypeSRV) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &PrivateDnsZones_SRV_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (typeSRV *PrivateDnsZonesRecordsetTypeSRV) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(typeSRV.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  typeSRV.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (typeSRV *PrivateDnsZonesRecordsetTypeSRV) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*PrivateDnsZones_SRV_STATUS); ok {
		typeSRV.Status = *st
		return nil
	}

	// Convert status to required version
	var st PrivateDnsZones_SRV_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	typeSRV.Status = st
	return nil
}

// Hub marks that this PrivateDnsZonesRecordsetTypeSRV is the hub type for conversion
func (typeSRV *PrivateDnsZonesRecordsetTypeSRV) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (typeSRV *PrivateDnsZonesRecordsetTypeSRV) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: typeSRV.Spec.OriginalVersion,
		Kind:    "PrivateDnsZonesRecordsetTypeSRV",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20200601.PrivateDnsZonesRecordsetTypeSRV
// Generator information:
// - Generated from: /privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/privatedns.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/SRV/{relativeRecordSetName}
type PrivateDnsZonesRecordsetTypeSRVList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateDnsZonesRecordsetTypeSRV `json:"items"`
}

// Storage version of v1beta20200601.PrivateDnsZones_SRV_Spec
type PrivateDnsZones_SRV_Spec struct {
	ARecords    []ARecord    `json:"aRecords,omitempty"`
	AaaaRecords []AaaaRecord `json:"aaaaRecords,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string            `json:"azureName,omitempty"`
	CnameRecord     *CnameRecord      `json:"cnameRecord,omitempty"`
	Etag            *string           `json:"etag,omitempty"`
	Metadata        map[string]string `json:"metadata,omitempty"`
	MxRecords       []MxRecord        `json:"mxRecords,omitempty"`
	OriginalVersion string            `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a network.azure.com/PrivateDnsZone resource
	Owner       *genruntime.KnownResourceReference `group:"network.azure.com" json:"owner,omitempty" kind:"PrivateDnsZone"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	PtrRecords  []PtrRecord                        `json:"ptrRecords,omitempty"`
	SoaRecord   *SoaRecord                         `json:"soaRecord,omitempty"`
	SrvRecords  []SrvRecord                        `json:"srvRecords,omitempty"`
	Ttl         *int                               `json:"ttl,omitempty"`
	TxtRecords  []TxtRecord                        `json:"txtRecords,omitempty"`
}

var _ genruntime.ConvertibleSpec = &PrivateDnsZones_SRV_Spec{}

// ConvertSpecFrom populates our PrivateDnsZones_SRV_Spec from the provided source
func (zonesSRV *PrivateDnsZones_SRV_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == zonesSRV {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(zonesSRV)
}

// ConvertSpecTo populates the provided destination from our PrivateDnsZones_SRV_Spec
func (zonesSRV *PrivateDnsZones_SRV_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == zonesSRV {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(zonesSRV)
}

// Storage version of v1beta20200601.PrivateDnsZones_SRV_STATUS
type PrivateDnsZones_SRV_STATUS struct {
	ARecords         []ARecord_STATUS       `json:"aRecords,omitempty"`
	AaaaRecords      []AaaaRecord_STATUS    `json:"aaaaRecords,omitempty"`
	CnameRecord      *CnameRecord_STATUS    `json:"cnameRecord,omitempty"`
	Conditions       []conditions.Condition `json:"conditions,omitempty"`
	Etag             *string                `json:"etag,omitempty"`
	Fqdn             *string                `json:"fqdn,omitempty"`
	Id               *string                `json:"id,omitempty"`
	IsAutoRegistered *bool                  `json:"isAutoRegistered,omitempty"`
	Metadata         map[string]string      `json:"metadata,omitempty"`
	MxRecords        []MxRecord_STATUS      `json:"mxRecords,omitempty"`
	Name             *string                `json:"name,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	PtrRecords       []PtrRecord_STATUS     `json:"ptrRecords,omitempty"`
	SoaRecord        *SoaRecord_STATUS      `json:"soaRecord,omitempty"`
	SrvRecords       []SrvRecord_STATUS     `json:"srvRecords,omitempty"`
	Ttl              *int                   `json:"ttl,omitempty"`
	TxtRecords       []TxtRecord_STATUS     `json:"txtRecords,omitempty"`
	Type             *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &PrivateDnsZones_SRV_STATUS{}

// ConvertStatusFrom populates our PrivateDnsZones_SRV_STATUS from the provided source
func (zonesSRV *PrivateDnsZones_SRV_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == zonesSRV {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(zonesSRV)
}

// ConvertStatusTo populates the provided destination from our PrivateDnsZones_SRV_STATUS
func (zonesSRV *PrivateDnsZones_SRV_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == zonesSRV {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(zonesSRV)
}

func init() {
	SchemeBuilder.Register(&PrivateDnsZonesRecordsetTypeSRV{}, &PrivateDnsZonesRecordsetTypeSRVList{})
}
