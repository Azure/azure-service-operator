// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

type ServersFailoverGroup_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *FailoverGroupProperties_STATUS_ARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// Properties of a failover group.
type FailoverGroupProperties_STATUS_ARM struct {
	// Databases: List of databases in the failover group.
	Databases []string `json:"databases,omitempty"`

	// PartnerServers: List of partner server information for the failover group.
	PartnerServers []PartnerInfo_STATUS_ARM `json:"partnerServers,omitempty"`

	// ReadOnlyEndpoint: Read-only endpoint of the failover group instance.
	ReadOnlyEndpoint *FailoverGroupReadOnlyEndpoint_STATUS_ARM `json:"readOnlyEndpoint,omitempty"`

	// ReadWriteEndpoint: Read-write endpoint of the failover group instance.
	ReadWriteEndpoint *FailoverGroupReadWriteEndpoint_STATUS_ARM `json:"readWriteEndpoint,omitempty"`

	// ReplicationRole: Local replication role of the failover group instance.
	ReplicationRole *FailoverGroupProperties_ReplicationRole_STATUS_ARM `json:"replicationRole,omitempty"`

	// ReplicationState: Replication state of the failover group instance.
	ReplicationState *string `json:"replicationState,omitempty"`
}

type FailoverGroupProperties_ReplicationRole_STATUS_ARM string

const (
	FailoverGroupProperties_ReplicationRole_STATUS_ARM_Primary   = FailoverGroupProperties_ReplicationRole_STATUS_ARM("Primary")
	FailoverGroupProperties_ReplicationRole_STATUS_ARM_Secondary = FailoverGroupProperties_ReplicationRole_STATUS_ARM("Secondary")
)

// Mapping from string to FailoverGroupProperties_ReplicationRole_STATUS_ARM
var failoverGroupProperties_ReplicationRole_STATUS_ARM_Values = map[string]FailoverGroupProperties_ReplicationRole_STATUS_ARM{
	"primary":   FailoverGroupProperties_ReplicationRole_STATUS_ARM_Primary,
	"secondary": FailoverGroupProperties_ReplicationRole_STATUS_ARM_Secondary,
}

// Read-only endpoint of the failover group instance.
type FailoverGroupReadOnlyEndpoint_STATUS_ARM struct {
	// FailoverPolicy: Failover policy of the read-only endpoint for the failover group.
	FailoverPolicy *FailoverGroupReadOnlyEndpoint_FailoverPolicy_STATUS_ARM `json:"failoverPolicy,omitempty"`
}

// Read-write endpoint of the failover group instance.
type FailoverGroupReadWriteEndpoint_STATUS_ARM struct {
	// FailoverPolicy: Failover policy of the read-write endpoint for the failover group. If failoverPolicy is Automatic then
	// failoverWithDataLossGracePeriodMinutes is required.
	FailoverPolicy *FailoverGroupReadWriteEndpoint_FailoverPolicy_STATUS_ARM `json:"failoverPolicy,omitempty"`

	// FailoverWithDataLossGracePeriodMinutes: Grace period before failover with data loss is attempted for the read-write
	// endpoint. If failoverPolicy is Automatic then failoverWithDataLossGracePeriodMinutes is required.
	FailoverWithDataLossGracePeriodMinutes *int `json:"failoverWithDataLossGracePeriodMinutes,omitempty"`
}

// Partner server information for the failover group.
type PartnerInfo_STATUS_ARM struct {
	// Id: Resource identifier of the partner server.
	Id *string `json:"id,omitempty"`

	// Location: Geo location of the partner server.
	Location *string `json:"location,omitempty"`

	// ReplicationRole: Replication role of the partner server.
	ReplicationRole *PartnerInfo_ReplicationRole_STATUS_ARM `json:"replicationRole,omitempty"`
}

type FailoverGroupReadOnlyEndpoint_FailoverPolicy_STATUS_ARM string

const (
	FailoverGroupReadOnlyEndpoint_FailoverPolicy_STATUS_ARM_Disabled = FailoverGroupReadOnlyEndpoint_FailoverPolicy_STATUS_ARM("Disabled")
	FailoverGroupReadOnlyEndpoint_FailoverPolicy_STATUS_ARM_Enabled  = FailoverGroupReadOnlyEndpoint_FailoverPolicy_STATUS_ARM("Enabled")
)

// Mapping from string to FailoverGroupReadOnlyEndpoint_FailoverPolicy_STATUS_ARM
var failoverGroupReadOnlyEndpoint_FailoverPolicy_STATUS_ARM_Values = map[string]FailoverGroupReadOnlyEndpoint_FailoverPolicy_STATUS_ARM{
	"disabled": FailoverGroupReadOnlyEndpoint_FailoverPolicy_STATUS_ARM_Disabled,
	"enabled":  FailoverGroupReadOnlyEndpoint_FailoverPolicy_STATUS_ARM_Enabled,
}

type FailoverGroupReadWriteEndpoint_FailoverPolicy_STATUS_ARM string

const (
	FailoverGroupReadWriteEndpoint_FailoverPolicy_STATUS_ARM_Automatic = FailoverGroupReadWriteEndpoint_FailoverPolicy_STATUS_ARM("Automatic")
	FailoverGroupReadWriteEndpoint_FailoverPolicy_STATUS_ARM_Manual    = FailoverGroupReadWriteEndpoint_FailoverPolicy_STATUS_ARM("Manual")
)

// Mapping from string to FailoverGroupReadWriteEndpoint_FailoverPolicy_STATUS_ARM
var failoverGroupReadWriteEndpoint_FailoverPolicy_STATUS_ARM_Values = map[string]FailoverGroupReadWriteEndpoint_FailoverPolicy_STATUS_ARM{
	"automatic": FailoverGroupReadWriteEndpoint_FailoverPolicy_STATUS_ARM_Automatic,
	"manual":    FailoverGroupReadWriteEndpoint_FailoverPolicy_STATUS_ARM_Manual,
}

type PartnerInfo_ReplicationRole_STATUS_ARM string

const (
	PartnerInfo_ReplicationRole_STATUS_ARM_Primary   = PartnerInfo_ReplicationRole_STATUS_ARM("Primary")
	PartnerInfo_ReplicationRole_STATUS_ARM_Secondary = PartnerInfo_ReplicationRole_STATUS_ARM("Secondary")
)

// Mapping from string to PartnerInfo_ReplicationRole_STATUS_ARM
var partnerInfo_ReplicationRole_STATUS_ARM_Values = map[string]PartnerInfo_ReplicationRole_STATUS_ARM{
	"primary":   PartnerInfo_ReplicationRole_STATUS_ARM_Primary,
	"secondary": PartnerInfo_ReplicationRole_STATUS_ARM_Secondary,
}
