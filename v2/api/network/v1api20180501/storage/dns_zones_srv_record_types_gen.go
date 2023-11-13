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

// +kubebuilder:rbac:groups=network.azure.com,resources=dnszonessrvrecords,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={dnszonessrvrecords/status,dnszonessrvrecords/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20180501.DnsZonesSRVRecord
// Generator information:
// - Generated from: /dns/resource-manager/Microsoft.Network/stable/2018-05-01/dns.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/SRV/{relativeRecordSetName}
type DnsZonesSRVRecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsZones_SRV_Spec   `json:"spec,omitempty"`
	Status            DnsZones_SRV_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &DnsZonesSRVRecord{}

// GetConditions returns the conditions of the resource
func (record *DnsZonesSRVRecord) GetConditions() conditions.Conditions {
	return record.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (record *DnsZonesSRVRecord) SetConditions(conditions conditions.Conditions) {
	record.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &DnsZonesSRVRecord{}

// AzureName returns the Azure name of the resource
func (record *DnsZonesSRVRecord) AzureName() string {
	return record.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-05-01"
func (record DnsZonesSRVRecord) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (record *DnsZonesSRVRecord) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (record *DnsZonesSRVRecord) GetSpec() genruntime.ConvertibleSpec {
	return &record.Spec
}

// GetStatus returns the status of this resource
func (record *DnsZonesSRVRecord) GetStatus() genruntime.ConvertibleStatus {
	return &record.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (record *DnsZonesSRVRecord) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/dnsZones/SRV"
func (record *DnsZonesSRVRecord) GetType() string {
	return "Microsoft.Network/dnsZones/SRV"
}

// NewEmptyStatus returns a new empty (blank) status
func (record *DnsZonesSRVRecord) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DnsZones_SRV_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (record *DnsZonesSRVRecord) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(record.Spec)
	return record.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (record *DnsZonesSRVRecord) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DnsZones_SRV_STATUS); ok {
		record.Status = *st
		return nil
	}

	// Convert status to required version
	var st DnsZones_SRV_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	record.Status = st
	return nil
}

// Hub marks that this DnsZonesSRVRecord is the hub type for conversion
func (record *DnsZonesSRVRecord) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (record *DnsZonesSRVRecord) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: record.Spec.OriginalVersion,
		Kind:    "DnsZonesSRVRecord",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20180501.DnsZonesSRVRecord
// Generator information:
// - Generated from: /dns/resource-manager/Microsoft.Network/stable/2018-05-01/dns.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/SRV/{relativeRecordSetName}
type DnsZonesSRVRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DnsZonesSRVRecord `json:"items"`
}

// Storage version of v1api20180501.DnsZones_SRV_Spec
type DnsZones_SRV_Spec struct {
	AAAARecords []AaaaRecord `json:"AAAARecords,omitempty"`
	ARecords    []ARecord    `json:"ARecords,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string            `json:"azureName,omitempty"`
	CNAMERecord     *CnameRecord      `json:"CNAMERecord,omitempty"`
	CaaRecords      []CaaRecord       `json:"caaRecords,omitempty"`
	MXRecords       []MxRecord        `json:"MXRecords,omitempty"`
	Metadata        map[string]string `json:"metadata,omitempty"`
	NSRecords       []NsRecord        `json:"NSRecords,omitempty"`
	OriginalVersion string            `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a network.azure.com/DnsZone resource
	Owner          *genruntime.KnownResourceReference `group:"network.azure.com" json:"owner,omitempty" kind:"DnsZone"`
	PTRRecords     []PtrRecord                        `json:"PTRRecords,omitempty"`
	PropertyBag    genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	SOARecord      *SoaRecord                         `json:"SOARecord,omitempty"`
	SRVRecords     []SrvRecord                        `json:"SRVRecords,omitempty"`
	TTL            *int                               `json:"TTL,omitempty"`
	TXTRecords     []TxtRecord                        `json:"TXTRecords,omitempty"`
	TargetResource *SubResource                       `json:"targetResource,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DnsZones_SRV_Spec{}

// ConvertSpecFrom populates our DnsZones_SRV_Spec from the provided source
func (zonesSRV *DnsZones_SRV_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == zonesSRV {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(zonesSRV)
}

// ConvertSpecTo populates the provided destination from our DnsZones_SRV_Spec
func (zonesSRV *DnsZones_SRV_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == zonesSRV {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(zonesSRV)
}

// Storage version of v1api20180501.DnsZones_SRV_STATUS
type DnsZones_SRV_STATUS struct {
	AAAARecords       []AaaaRecord_STATUS    `json:"AAAARecords,omitempty"`
	ARecords          []ARecord_STATUS       `json:"ARecords,omitempty"`
	CNAMERecord       *CnameRecord_STATUS    `json:"CNAMERecord,omitempty"`
	CaaRecords        []CaaRecord_STATUS     `json:"caaRecords,omitempty"`
	Conditions        []conditions.Condition `json:"conditions,omitempty"`
	Etag              *string                `json:"etag,omitempty"`
	Fqdn              *string                `json:"fqdn,omitempty"`
	Id                *string                `json:"id,omitempty"`
	MXRecords         []MxRecord_STATUS      `json:"MXRecords,omitempty"`
	Metadata          map[string]string      `json:"metadata,omitempty"`
	NSRecords         []NsRecord_STATUS      `json:"NSRecords,omitempty"`
	Name              *string                `json:"name,omitempty"`
	PTRRecords        []PtrRecord_STATUS     `json:"PTRRecords,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	SOARecord         *SoaRecord_STATUS      `json:"SOARecord,omitempty"`
	SRVRecords        []SrvRecord_STATUS     `json:"SRVRecords,omitempty"`
	TTL               *int                   `json:"TTL,omitempty"`
	TXTRecords        []TxtRecord_STATUS     `json:"TXTRecords,omitempty"`
	TargetResource    *SubResource_STATUS    `json:"targetResource,omitempty"`
	Type              *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DnsZones_SRV_STATUS{}

// ConvertStatusFrom populates our DnsZones_SRV_STATUS from the provided source
func (zonesSRV *DnsZones_SRV_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == zonesSRV {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(zonesSRV)
}

// ConvertStatusTo populates the provided destination from our DnsZones_SRV_STATUS
func (zonesSRV *DnsZones_SRV_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == zonesSRV {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(zonesSRV)
}

func init() {
	SchemeBuilder.Register(&DnsZonesSRVRecord{}, &DnsZonesSRVRecordList{})
}
