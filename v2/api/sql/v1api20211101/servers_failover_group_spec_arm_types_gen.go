// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Servers_FailoverGroup_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *FailoverGroupProperties_ARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Servers_FailoverGroup_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (group Servers_FailoverGroup_Spec_ARM) GetAPIVersion() string {
	return "2021-11-01"
}

// GetName returns the Name of the resource
func (group *Servers_FailoverGroup_Spec_ARM) GetName() string {
	return group.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/failoverGroups"
func (group *Servers_FailoverGroup_Spec_ARM) GetType() string {
	return "Microsoft.Sql/servers/failoverGroups"
}

// Properties of a failover group.
type FailoverGroupProperties_ARM struct {
	Databases []string `json:"databases,omitempty"`

	// PartnerServers: List of partner server information for the failover group.
	PartnerServers []PartnerInfo_ARM `json:"partnerServers,omitempty"`

	// ReadOnlyEndpoint: Read-only endpoint of the failover group instance.
	ReadOnlyEndpoint *FailoverGroupReadOnlyEndpoint_ARM `json:"readOnlyEndpoint,omitempty"`

	// ReadWriteEndpoint: Read-write endpoint of the failover group instance.
	ReadWriteEndpoint *FailoverGroupReadWriteEndpoint_ARM `json:"readWriteEndpoint,omitempty"`
}

// Read-only endpoint of the failover group instance.
type FailoverGroupReadOnlyEndpoint_ARM struct {
	// FailoverPolicy: Failover policy of the read-only endpoint for the failover group.
	FailoverPolicy *FailoverGroupReadOnlyEndpoint_FailoverPolicy_ARM `json:"failoverPolicy,omitempty"`
}

// Read-write endpoint of the failover group instance.
type FailoverGroupReadWriteEndpoint_ARM struct {
	// FailoverPolicy: Failover policy of the read-write endpoint for the failover group. If failoverPolicy is Automatic then
	// failoverWithDataLossGracePeriodMinutes is required.
	FailoverPolicy *FailoverGroupReadWriteEndpoint_FailoverPolicy_ARM `json:"failoverPolicy,omitempty"`

	// FailoverWithDataLossGracePeriodMinutes: Grace period before failover with data loss is attempted for the read-write
	// endpoint. If failoverPolicy is Automatic then failoverWithDataLossGracePeriodMinutes is required.
	FailoverWithDataLossGracePeriodMinutes *int `json:"failoverWithDataLossGracePeriodMinutes,omitempty"`
}

// Partner server information for the failover group.
type PartnerInfo_ARM struct {
	Id *string `json:"id,omitempty"`
}

// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type FailoverGroupReadOnlyEndpoint_FailoverPolicy_ARM string

const (
	FailoverGroupReadOnlyEndpoint_FailoverPolicy_ARM_Disabled = FailoverGroupReadOnlyEndpoint_FailoverPolicy_ARM("Disabled")
	FailoverGroupReadOnlyEndpoint_FailoverPolicy_ARM_Enabled  = FailoverGroupReadOnlyEndpoint_FailoverPolicy_ARM("Enabled")
)

// Mapping from string to FailoverGroupReadOnlyEndpoint_FailoverPolicy_ARM
var failoverGroupReadOnlyEndpoint_FailoverPolicy_ARM_Values = map[string]FailoverGroupReadOnlyEndpoint_FailoverPolicy_ARM{
	"disabled": FailoverGroupReadOnlyEndpoint_FailoverPolicy_ARM_Disabled,
	"enabled":  FailoverGroupReadOnlyEndpoint_FailoverPolicy_ARM_Enabled,
}

// +kubebuilder:validation:Enum={"Automatic","Manual"}
type FailoverGroupReadWriteEndpoint_FailoverPolicy_ARM string

const (
	FailoverGroupReadWriteEndpoint_FailoverPolicy_ARM_Automatic = FailoverGroupReadWriteEndpoint_FailoverPolicy_ARM("Automatic")
	FailoverGroupReadWriteEndpoint_FailoverPolicy_ARM_Manual    = FailoverGroupReadWriteEndpoint_FailoverPolicy_ARM("Manual")
)

// Mapping from string to FailoverGroupReadWriteEndpoint_FailoverPolicy_ARM
var failoverGroupReadWriteEndpoint_FailoverPolicy_ARM_Values = map[string]FailoverGroupReadWriteEndpoint_FailoverPolicy_ARM{
	"automatic": FailoverGroupReadWriteEndpoint_FailoverPolicy_ARM_Automatic,
	"manual":    FailoverGroupReadWriteEndpoint_FailoverPolicy_ARM_Manual,
}
