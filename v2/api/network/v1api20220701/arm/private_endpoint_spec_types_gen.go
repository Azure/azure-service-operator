// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type PrivateEndpoint_Spec struct {
	// ExtendedLocation: The extended location of the load balancer.
	ExtendedLocation *ExtendedLocation `json:"extendedLocation,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties of the private endpoint.
	Properties *PrivateEndpointProperties `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &PrivateEndpoint_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-07-01"
func (endpoint PrivateEndpoint_Spec) GetAPIVersion() string {
	return "2022-07-01"
}

// GetName returns the Name of the resource
func (endpoint *PrivateEndpoint_Spec) GetName() string {
	return endpoint.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/privateEndpoints"
func (endpoint *PrivateEndpoint_Spec) GetType() string {
	return "Microsoft.Network/privateEndpoints"
}

// ExtendedLocation complex type.
type ExtendedLocation struct {
	// Name: The name of the extended location.
	Name *string `json:"name,omitempty"`

	// Type: The type of the extended location.
	Type *ExtendedLocationType `json:"type,omitempty"`
}

// Properties of the private endpoint.
type PrivateEndpointProperties struct {
	// ApplicationSecurityGroups: Application security groups in which the private endpoint IP configuration is included.
	ApplicationSecurityGroups []ApplicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbedded `json:"applicationSecurityGroups,omitempty"`

	// CustomNetworkInterfaceName: The custom name of the network interface attached to the private endpoint.
	CustomNetworkInterfaceName *string `json:"customNetworkInterfaceName,omitempty"`

	// IpConfigurations: A list of IP configurations of the private endpoint. This will be used to map to the First Party
	// Service's endpoints.
	IpConfigurations []PrivateEndpointIPConfiguration `json:"ipConfigurations,omitempty"`

	// ManualPrivateLinkServiceConnections: A grouping of information about the connection to the remote resource. Used when
	// the network admin does not have access to approve connections to the remote resource.
	ManualPrivateLinkServiceConnections []PrivateLinkServiceConnection `json:"manualPrivateLinkServiceConnections,omitempty"`

	// PrivateLinkServiceConnections: A grouping of information about the connection to the remote resource.
	PrivateLinkServiceConnections []PrivateLinkServiceConnection `json:"privateLinkServiceConnections,omitempty"`

	// Subnet: The ID of the subnet from which the private IP will be allocated.
	Subnet *Subnet_PrivateEndpoint_SubResourceEmbedded `json:"subnet,omitempty"`
}

// An application security group in a resource group.
type ApplicationSecurityGroupSpec_PrivateEndpoint_SubResourceEmbedded struct {
	Id *string `json:"id,omitempty"`
}

// The supported ExtendedLocation types. Currently only EdgeZone is supported in Microsoft.Network resources.
// +kubebuilder:validation:Enum={"EdgeZone"}
type ExtendedLocationType string

const ExtendedLocationType_EdgeZone = ExtendedLocationType("EdgeZone")

// Mapping from string to ExtendedLocationType
var extendedLocationType_Values = map[string]ExtendedLocationType{
	"edgezone": ExtendedLocationType_EdgeZone,
}

// An IP Configuration of the private endpoint.
type PrivateEndpointIPConfiguration struct {
	// Name: The name of the resource that is unique within a resource group.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of private endpoint IP configurations.
	Properties *PrivateEndpointIPConfigurationProperties `json:"properties,omitempty"`
}

// PrivateLinkServiceConnection resource.
type PrivateLinkServiceConnection struct {
	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the private link service connection.
	Properties *PrivateLinkServiceConnectionProperties `json:"properties,omitempty"`
}

// Subnet in a virtual network resource.
type Subnet_PrivateEndpoint_SubResourceEmbedded struct {
	Id *string `json:"id,omitempty"`
}

// Properties of an IP Configuration of the private endpoint.
type PrivateEndpointIPConfigurationProperties struct {
	// GroupId: The ID of a group obtained from the remote resource that this private endpoint should connect to.
	GroupId *string `json:"groupId,omitempty"`

	// MemberName: The member name of a group obtained from the remote resource that this private endpoint should connect to.
	MemberName *string `json:"memberName,omitempty"`

	// PrivateIPAddress: A private ip address obtained from the private endpoint's subnet.
	PrivateIPAddress *string `json:"privateIPAddress,omitempty"`
}

// Properties of the PrivateLinkServiceConnection.
type PrivateLinkServiceConnectionProperties struct {
	// GroupIds: The ID(s) of the group(s) obtained from the remote resource that this private endpoint should connect to.
	GroupIds []string `json:"groupIds,omitempty"`

	// PrivateLinkServiceConnectionState: A collection of read-only information about the state of the connection to the remote
	// resource.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `json:"privateLinkServiceConnectionState,omitempty"`
	PrivateLinkServiceId              *string                            `json:"privateLinkServiceId,omitempty"`

	// RequestMessage: A message passed to the owner of the remote resource with this connection request. Restricted to 140
	// chars.
	RequestMessage *string `json:"requestMessage,omitempty"`
}

// A collection of information about the state of the connection between service consumer and provider.
type PrivateLinkServiceConnectionState struct {
	// ActionsRequired: A message indicating if changes on the service provider require any updates on the consumer.
	ActionsRequired *string `json:"actionsRequired,omitempty"`

	// Description: The reason for approval/rejection of the connection.
	Description *string `json:"description,omitempty"`

	// Status: Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
	Status *string `json:"status,omitempty"`
}
