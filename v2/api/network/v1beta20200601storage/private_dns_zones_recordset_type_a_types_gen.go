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

// +kubebuilder:rbac:groups=network.azure.com,resources=privatednszonesrecordsettypeas,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={privatednszonesrecordsettypeas/status,privatednszonesrecordsettypeas/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20200601.PrivateDnsZonesRecordsetTypeA
// Generator information:
// - Generated from: /privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/privatedns.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/A/{relativeRecordSetName}
type PrivateDnsZonesRecordsetTypeA struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateDnsZones_A_Spec   `json:"spec,omitempty"`
	Status            PrivateDnsZones_A_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &PrivateDnsZonesRecordsetTypeA{}

// GetConditions returns the conditions of the resource
func (typeA *PrivateDnsZonesRecordsetTypeA) GetConditions() conditions.Conditions {
	return typeA.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (typeA *PrivateDnsZonesRecordsetTypeA) SetConditions(conditions conditions.Conditions) {
	typeA.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &PrivateDnsZonesRecordsetTypeA{}

// AzureName returns the Azure name of the resource
func (typeA *PrivateDnsZonesRecordsetTypeA) AzureName() string {
	return typeA.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (typeA PrivateDnsZonesRecordsetTypeA) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (typeA *PrivateDnsZonesRecordsetTypeA) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (typeA *PrivateDnsZonesRecordsetTypeA) GetSpec() genruntime.ConvertibleSpec {
	return &typeA.Spec
}

// GetStatus returns the status of this resource
func (typeA *PrivateDnsZonesRecordsetTypeA) GetStatus() genruntime.ConvertibleStatus {
	return &typeA.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/privateDnsZones/A"
func (typeA *PrivateDnsZonesRecordsetTypeA) GetType() string {
	return "Microsoft.Network/privateDnsZones/A"
}

// NewEmptyStatus returns a new empty (blank) status
func (typeA *PrivateDnsZonesRecordsetTypeA) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &PrivateDnsZones_A_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (typeA *PrivateDnsZonesRecordsetTypeA) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(typeA.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  typeA.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (typeA *PrivateDnsZonesRecordsetTypeA) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*PrivateDnsZones_A_STATUS); ok {
		typeA.Status = *st
		return nil
	}

	// Convert status to required version
	var st PrivateDnsZones_A_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	typeA.Status = st
	return nil
}

// Hub marks that this PrivateDnsZonesRecordsetTypeA is the hub type for conversion
func (typeA *PrivateDnsZonesRecordsetTypeA) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (typeA *PrivateDnsZonesRecordsetTypeA) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: typeA.Spec.OriginalVersion,
		Kind:    "PrivateDnsZonesRecordsetTypeA",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20200601.PrivateDnsZonesRecordsetTypeA
// Generator information:
// - Generated from: /privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/privatedns.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/A/{relativeRecordSetName}
type PrivateDnsZonesRecordsetTypeAList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateDnsZonesRecordsetTypeA `json:"items"`
}

// Storage version of v1beta20200601.APIVersion
// +kubebuilder:validation:Enum={"2020-06-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2020-06-01")

// Storage version of v1beta20200601.PrivateDnsZones_A_Spec
type PrivateDnsZones_A_Spec struct {
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

var _ genruntime.ConvertibleSpec = &PrivateDnsZones_A_Spec{}

// ConvertSpecFrom populates our PrivateDnsZones_A_Spec from the provided source
func (zonesA *PrivateDnsZones_A_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == zonesA {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(zonesA)
}

// ConvertSpecTo populates the provided destination from our PrivateDnsZones_A_Spec
func (zonesA *PrivateDnsZones_A_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == zonesA {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(zonesA)
}

// Storage version of v1beta20200601.PrivateDnsZones_A_STATUS
type PrivateDnsZones_A_STATUS struct {
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

var _ genruntime.ConvertibleStatus = &PrivateDnsZones_A_STATUS{}

// ConvertStatusFrom populates our PrivateDnsZones_A_STATUS from the provided source
func (zonesA *PrivateDnsZones_A_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == zonesA {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(zonesA)
}

// ConvertStatusTo populates the provided destination from our PrivateDnsZones_A_STATUS
func (zonesA *PrivateDnsZones_A_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == zonesA {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(zonesA)
}

// Storage version of v1beta20200601.AaaaRecord
// An AAAA record.
type AaaaRecord struct {
	Ipv6Address *string                `json:"ipv6Address,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20200601.AaaaRecord_STATUS
// An AAAA record.
type AaaaRecord_STATUS struct {
	Ipv6Address *string                `json:"ipv6Address,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20200601.ARecord
// An A record.
type ARecord struct {
	Ipv4Address *string                `json:"ipv4Address,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20200601.ARecord_STATUS
// An A record.
type ARecord_STATUS struct {
	Ipv4Address *string                `json:"ipv4Address,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20200601.CnameRecord
// A CNAME record.
type CnameRecord struct {
	Cname       *string                `json:"cname,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20200601.CnameRecord_STATUS
// A CNAME record.
type CnameRecord_STATUS struct {
	Cname       *string                `json:"cname,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20200601.MxRecord
// An MX record.
type MxRecord struct {
	Exchange    *string                `json:"exchange,omitempty"`
	Preference  *int                   `json:"preference,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20200601.MxRecord_STATUS
// An MX record.
type MxRecord_STATUS struct {
	Exchange    *string                `json:"exchange,omitempty"`
	Preference  *int                   `json:"preference,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20200601.PtrRecord
// A PTR record.
type PtrRecord struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Ptrdname    *string                `json:"ptrdname,omitempty"`
}

// Storage version of v1beta20200601.PtrRecord_STATUS
// A PTR record.
type PtrRecord_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Ptrdname    *string                `json:"ptrdname,omitempty"`
}

// Storage version of v1beta20200601.SoaRecord
// An SOA record.
type SoaRecord struct {
	Email        *string                `json:"email,omitempty"`
	ExpireTime   *int                   `json:"expireTime,omitempty"`
	Host         *string                `json:"host,omitempty"`
	MinimumTtl   *int                   `json:"minimumTtl,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RefreshTime  *int                   `json:"refreshTime,omitempty"`
	RetryTime    *int                   `json:"retryTime,omitempty"`
	SerialNumber *int                   `json:"serialNumber,omitempty"`
}

// Storage version of v1beta20200601.SoaRecord_STATUS
// An SOA record.
type SoaRecord_STATUS struct {
	Email        *string                `json:"email,omitempty"`
	ExpireTime   *int                   `json:"expireTime,omitempty"`
	Host         *string                `json:"host,omitempty"`
	MinimumTtl   *int                   `json:"minimumTtl,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RefreshTime  *int                   `json:"refreshTime,omitempty"`
	RetryTime    *int                   `json:"retryTime,omitempty"`
	SerialNumber *int                   `json:"serialNumber,omitempty"`
}

// Storage version of v1beta20200601.SrvRecord
// An SRV record.
type SrvRecord struct {
	Port        *int                   `json:"port,omitempty"`
	Priority    *int                   `json:"priority,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Target      *string                `json:"target,omitempty"`
	Weight      *int                   `json:"weight,omitempty"`
}

// Storage version of v1beta20200601.SrvRecord_STATUS
// An SRV record.
type SrvRecord_STATUS struct {
	Port        *int                   `json:"port,omitempty"`
	Priority    *int                   `json:"priority,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Target      *string                `json:"target,omitempty"`
	Weight      *int                   `json:"weight,omitempty"`
}

// Storage version of v1beta20200601.TxtRecord
// A TXT record.
type TxtRecord struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       []string               `json:"value,omitempty"`
}

// Storage version of v1beta20200601.TxtRecord_STATUS
// A TXT record.
type TxtRecord_STATUS struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       []string               `json:"value,omitempty"`
}

func init() {
	SchemeBuilder.Register(&PrivateDnsZonesRecordsetTypeA{}, &PrivateDnsZonesRecordsetTypeAList{})
}
