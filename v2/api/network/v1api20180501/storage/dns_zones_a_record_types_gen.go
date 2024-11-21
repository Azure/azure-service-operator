// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/rotisserie/eris"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=network.azure.com,resources=dnszonesarecords,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={dnszonesarecords/status,dnszonesarecords/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20180501.DnsZonesARecord
// Generator information:
// - Generated from: /dns/resource-manager/Microsoft.Network/stable/2018-05-01/dns.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/A/{relativeRecordSetName}
type DnsZonesARecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsZonesARecord_Spec   `json:"spec,omitempty"`
	Status            DnsZonesARecord_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &DnsZonesARecord{}

// GetConditions returns the conditions of the resource
func (record *DnsZonesARecord) GetConditions() conditions.Conditions {
	return record.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (record *DnsZonesARecord) SetConditions(conditions conditions.Conditions) {
	record.Status.Conditions = conditions
}

var _ configmaps.Exporter = &DnsZonesARecord{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (record *DnsZonesARecord) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if record.Spec.OperatorSpec == nil {
		return nil
	}
	return record.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &DnsZonesARecord{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (record *DnsZonesARecord) SecretDestinationExpressions() []*core.DestinationExpression {
	if record.Spec.OperatorSpec == nil {
		return nil
	}
	return record.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &DnsZonesARecord{}

// AzureName returns the Azure name of the resource
func (record *DnsZonesARecord) AzureName() string {
	return record.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-05-01"
func (record DnsZonesARecord) GetAPIVersion() string {
	return "2018-05-01"
}

// GetResourceScope returns the scope of the resource
func (record *DnsZonesARecord) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (record *DnsZonesARecord) GetSpec() genruntime.ConvertibleSpec {
	return &record.Spec
}

// GetStatus returns the status of this resource
func (record *DnsZonesARecord) GetStatus() genruntime.ConvertibleStatus {
	return &record.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (record *DnsZonesARecord) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/dnsZones/A"
func (record *DnsZonesARecord) GetType() string {
	return "Microsoft.Network/dnsZones/A"
}

// NewEmptyStatus returns a new empty (blank) status
func (record *DnsZonesARecord) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DnsZonesARecord_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (record *DnsZonesARecord) Owner() *genruntime.ResourceReference {
	if record.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(record.Spec)
	return record.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (record *DnsZonesARecord) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DnsZonesARecord_STATUS); ok {
		record.Status = *st
		return nil
	}

	// Convert status to required version
	var st DnsZonesARecord_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	record.Status = st
	return nil
}

// Hub marks that this DnsZonesARecord is the hub type for conversion
func (record *DnsZonesARecord) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (record *DnsZonesARecord) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: record.Spec.OriginalVersion,
		Kind:    "DnsZonesARecord",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20180501.DnsZonesARecord
// Generator information:
// - Generated from: /dns/resource-manager/Microsoft.Network/stable/2018-05-01/dns.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/A/{relativeRecordSetName}
type DnsZonesARecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DnsZonesARecord `json:"items"`
}

// Storage version of v1api20180501.DnsZonesARecord_Spec
type DnsZonesARecord_Spec struct {
	AAAARecords []AaaaRecord `json:"AAAARecords,omitempty"`
	ARecords    []ARecord    `json:"ARecords,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                       `json:"azureName,omitempty"`
	CNAMERecord     *CnameRecord                 `json:"CNAMERecord,omitempty"`
	CaaRecords      []CaaRecord                  `json:"caaRecords,omitempty"`
	Etag            *string                      `json:"etag,omitempty"`
	MXRecords       []MxRecord                   `json:"MXRecords,omitempty"`
	Metadata        map[string]string            `json:"metadata,omitempty"`
	NSRecords       []NsRecord                   `json:"NSRecords,omitempty"`
	OperatorSpec    *DnsZonesARecordOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                       `json:"originalVersion,omitempty"`

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

var _ genruntime.ConvertibleSpec = &DnsZonesARecord_Spec{}

// ConvertSpecFrom populates our DnsZonesARecord_Spec from the provided source
func (record *DnsZonesARecord_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == record {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(record)
}

// ConvertSpecTo populates the provided destination from our DnsZonesARecord_Spec
func (record *DnsZonesARecord_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == record {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(record)
}

// Storage version of v1api20180501.DnsZonesARecord_STATUS
type DnsZonesARecord_STATUS struct {
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

var _ genruntime.ConvertibleStatus = &DnsZonesARecord_STATUS{}

// ConvertStatusFrom populates our DnsZonesARecord_STATUS from the provided source
func (record *DnsZonesARecord_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == record {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(record)
}

// ConvertStatusTo populates the provided destination from our DnsZonesARecord_STATUS
func (record *DnsZonesARecord_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == record {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(record)
}

// Storage version of v1api20180501.DnsZonesARecordOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type DnsZonesARecordOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

func init() {
	SchemeBuilder.Register(&DnsZonesARecord{}, &DnsZonesARecordList{})
}
