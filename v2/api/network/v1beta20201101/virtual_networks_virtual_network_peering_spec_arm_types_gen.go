// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type VirtualNetworks_VirtualNetworkPeering_Spec_ARM struct {
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name string `json:"name,omitempty"`

	// Properties: Properties of the virtual network peering.
	Properties *VirtualNetworkPeeringPropertiesFormat_ARM `json:"properties,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ARMResourceSpec = &VirtualNetworks_VirtualNetworkPeering_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (peering VirtualNetworks_VirtualNetworkPeering_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (peering *VirtualNetworks_VirtualNetworkPeering_Spec_ARM) GetName() string {
	return peering.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/virtualNetworks/virtualNetworkPeerings"
func (peering *VirtualNetworks_VirtualNetworkPeering_Spec_ARM) GetType() string {
	return "Microsoft.Network/virtualNetworks/virtualNetworkPeerings"
}

type VirtualNetworkPeeringPropertiesFormat_ARM struct {
	// AllowForwardedTraffic: Whether the forwarded traffic from the VMs in the local virtual network will be
	// allowed/disallowed in remote virtual network.
	AllowForwardedTraffic *bool `json:"allowForwardedTraffic,omitempty"`

	// AllowGatewayTransit: If gateway links can be used in remote virtual networking to link to this virtual network.
	AllowGatewayTransit *bool `json:"allowGatewayTransit,omitempty"`

	// AllowVirtualNetworkAccess: Whether the VMs in the local virtual network space would be able to access the VMs in remote
	// virtual network space.
	AllowVirtualNetworkAccess *bool `json:"allowVirtualNetworkAccess,omitempty"`

	// DoNotVerifyRemoteGateways: If we need to verify the provisioning state of the remote gateway.
	DoNotVerifyRemoteGateways *bool `json:"doNotVerifyRemoteGateways,omitempty"`

	// PeeringState: The status of the virtual network peering.
	PeeringState *VirtualNetworkPeeringPropertiesFormat_PeeringState `json:"peeringState,omitempty"`

	// RemoteAddressSpace: The reference to the remote virtual network address space.
	RemoteAddressSpace *AddressSpace_ARM `json:"remoteAddressSpace,omitempty"`

	// RemoteBgpCommunities: The reference to the remote virtual network's Bgp Communities.
	RemoteBgpCommunities *VirtualNetworkBgpCommunities_ARM `json:"remoteBgpCommunities,omitempty"`

	// RemoteVirtualNetwork: The reference to the remote virtual network. The remote virtual network can be in the same or
	// different region (preview). See here to register for the preview and learn more
	// (https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-create-peering).
	RemoteVirtualNetwork *SubResource_ARM `json:"remoteVirtualNetwork,omitempty"`

	// UseRemoteGateways: If remote gateways can be used on this virtual network. If the flag is set to true, and
	// allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for
	// transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a
	// gateway.
	UseRemoteGateways *bool `json:"useRemoteGateways,omitempty"`
}
