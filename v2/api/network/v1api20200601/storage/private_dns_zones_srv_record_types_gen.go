// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=network.azure.com,resources=privatednszonessrvrecords,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={privatednszonessrvrecords/status,privatednszonessrvrecords/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20200601.PrivateDnsZonesSRVRecord
// Generator information:
// - Generated from: /privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/privatedns.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/SRV/{relativeRecordSetName}
type PrivateDnsZonesSRVRecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateDnsZones_SRV_Spec   `json:"spec,omitempty"`
	Status            PrivateDnsZones_SRV_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &PrivateDnsZonesSRVRecord{}

// GetConditions returns the conditions of the resource
func (record *PrivateDnsZonesSRVRecord) GetConditions() conditions.Conditions {
	return record.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (record *PrivateDnsZonesSRVRecord) SetConditions(conditions conditions.Conditions) {
	record.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &PrivateDnsZonesSRVRecord{}

// AzureName returns the Azure name of the resource
func (record *PrivateDnsZonesSRVRecord) AzureName() string {
	return record.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (record PrivateDnsZonesSRVRecord) GetAPIVersion() string {
	return "2020-06-01"
}

// GetResourceScope returns the scope of the resource
func (record *PrivateDnsZonesSRVRecord) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (record *PrivateDnsZonesSRVRecord) GetSpec() genruntime.ConvertibleSpec {
	return &record.Spec
}

// GetStatus returns the status of this resource
func (record *PrivateDnsZonesSRVRecord) GetStatus() genruntime.ConvertibleStatus {
	return &record.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (record *PrivateDnsZonesSRVRecord) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/privateDnsZones/SRV"
func (record *PrivateDnsZonesSRVRecord) GetType() string {
	return "Microsoft.Network/privateDnsZones/SRV"
}

// NewEmptyStatus returns a new empty (blank) status
func (record *PrivateDnsZonesSRVRecord) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &PrivateDnsZones_SRV_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (record *PrivateDnsZonesSRVRecord) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(record.Spec)
	return record.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (record *PrivateDnsZonesSRVRecord) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*PrivateDnsZones_SRV_STATUS); ok {
		record.Status = *st
		return nil
	}

	// Convert status to required version
	var st PrivateDnsZones_SRV_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	record.Status = st
	return nil
}

// Hub marks that this PrivateDnsZonesSRVRecord is the hub type for conversion
func (record *PrivateDnsZonesSRVRecord) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (record *PrivateDnsZonesSRVRecord) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: record.Spec.OriginalVersion,
		Kind:    "PrivateDnsZonesSRVRecord",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20200601.PrivateDnsZonesSRVRecord
// Generator information:
// - Generated from: /privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/privatedns.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/SRV/{relativeRecordSetName}
type PrivateDnsZonesSRVRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateDnsZonesSRVRecord `json:"items"`
}

// Storage version of v1api20200601.PrivateDnsZones_SRV_Spec
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

// Storage version of v1api20200601.PrivateDnsZones_SRV_STATUS
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
	SchemeBuilder.Register(&PrivateDnsZonesSRVRecord{}, &PrivateDnsZonesSRVRecordList{})
}
