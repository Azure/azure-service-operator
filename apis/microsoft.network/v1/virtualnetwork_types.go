/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v1

import (
	"encoding/json"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/k8s-infra/pkg/zips"
)

type (
	// DHCPOptionsSpec contains an array of DNS servers available to VMs deployed in the virtual network
	DHCPOptionsSpec struct {
		// DNSServers a list of DNS servers IP addresses
		DNSServers []string `json:"dnsServers,omitempty"`
	}

	// AddressSpaceSpec contains an array of IP address ranges that can be used by subnets
	AddressSpaceSpec struct {
		// AddressPrefixes are a list of address blocks reserved for this virtual network in CIDR notation
		AddressPrefixes []string `json:"addressPrefixes,omitempty"`
	}

	// BGPCommunitiesSpec are BGP communities sent over ExpressRoute with each route corresponding to a prefix in this VNET
	BGPCommunitiesSpec struct {
		// RegionalCommunity is a BGP community associated with the region of the virtual network
		RegionalCommunity string `json:"regionalCommunity,omitempty"`

		// VirtualNetworkCommunity is the BGP community associated with the virtual network
		VirtualNetworkCommunity string `json:"virtualNetworkCommunity,omitempty"`
	}

	// SubnetProperties are the properties of the subnet
	SubnetProperties struct {
		// AddressPrefix for the subnet, eg. 10.0.0.0/24
		AddressPrefix string `json:"addressPrefix,omitempty"`

		// AddressPrefixes are a list of address prefixes for a subnet
		AddressPrefixes []string `json:"addressPrefixes,omitempty"`
	}

	// SubnetSpec is a subnet in a Virtual Network
	// TODO: (dj) I think this should probably be a slice of corev1.ObjectReference
	SubnetSpec struct {
		// ID of the subnet resource
		ID string `json:"id,omitempty"`

		// Name of the subnet
		Name string `json:"name,omitempty"`

		// Properties of the subnet
		Properties SubnetProperties `json:"properties,omitempty"`
	}

	// VirtualNetworkSpecProperties are the property bodies to be applied
	VirtualNetworkSpecProperties struct {
		// AddressSpace contains an array of IP address ranges that can be used by subnets
		AddressSpace AddressSpaceSpec `json:"addressSpace,omitempty"`

		// BGPCommunities are BGP communities sent over ExpressRoute with each route corresponding to a prefix in this VNET
		// +optional
		BGPCommunities *BGPCommunitiesSpec `json:"bgpCommunities,omitempty"`

		// DHCPOptions contains an array of DNS servers available to VMs deployed in the virtual network
		// +optional
		DHCPOptions *DHCPOptionsSpec `json:"dhcpOptions,omitempty"`

		// Subnets is a list of subnets in the VNET
		// +optional
		Subnets []SubnetSpec `json:"subnets,omitempty"`

		// EnableVMProtection indicates if VM protection is enabled for all the subnets in the virtual network
		// +optional
		EnableVMProtection bool `json:"enableVMProtection,omitempty"`
	}

	// VirtualNetworkSpec defines the desired state of VirtualNetwork
	VirtualNetworkSpec struct {
		// ResourceGroup is the Azure Resource Group the VirtualNetwork resides within
		ResourceGroup *corev1.ObjectReference `json:"group"`
		APIVersion    string                  `json:"apiVersion,omitempty"`
		Location      string                  `json:"location,omitempty"`
		Tags          map[string]string       `json:"tags,omitempty"`

		// Properties of the Virtual Network
		Properties VirtualNetworkSpecProperties `json:"properties,omitempty"`
	}

	// VirtualNetworkStatus defines the observed state of VirtualNetwork
	VirtualNetworkStatus struct {
		ID                string `json:"id,omitempty"`
		DeploymentID      string `json:"deploymentId,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty"`
	}

	// +kubebuilder:object:root=true
	// +kubebuilder:subresource:status
	// +kubebuilder:storageversion

	// VirtualNetwork is the Schema for the virtualnetworks API
	VirtualNetwork struct {
		metav1.TypeMeta   `json:",inline"`
		metav1.ObjectMeta `json:"metadata,omitempty"`

		Spec   VirtualNetworkSpec   `json:"spec,omitempty"`
		Status VirtualNetworkStatus `json:"status,omitempty"`
	}

	// +kubebuilder:object:root=true

	// VirtualNetworkList contains a list of VirtualNetwork
	VirtualNetworkList struct {
		metav1.TypeMeta `json:",inline"`
		metav1.ListMeta `json:"metadata,omitempty"`
		Items           []VirtualNetwork `json:"items"`
	}
)

func (*VirtualNetwork) Hub() {
	fmt.Println("Hub was called!")
}

func (vnet *VirtualNetwork) GetResourceGroupObjectRef() *corev1.ObjectReference {
	return vnet.Spec.ResourceGroup
}

func (vnet *VirtualNetwork) ToResource() (zips.Resource, error) {
	rgName := ""
	if vnet.Spec.ResourceGroup != nil {
		rgName = vnet.Spec.ResourceGroup.Name
	}

	res := zips.Resource{
		ID:                vnet.Status.ID,
		DeploymentID:      vnet.Status.DeploymentID,
		Type:              "Microsoft.Network/virtualNetworks",
		ResourceGroup:     rgName,
		Name:              vnet.Name,
		APIVersion:        vnet.Spec.APIVersion,
		Location:          vnet.Spec.Location,
		Tags:              vnet.Spec.Tags,
		ProvisioningState: zips.ProvisioningState(vnet.Status.ProvisioningState),
	}

	bits, err := json.Marshal(vnet.Spec.Properties)
	if err != nil {
		return res, err
	}
	res.Properties = bits

	return *res.SetAnnotations(vnet.Annotations), nil
}

func (vnet *VirtualNetwork) FromResource(res zips.Resource) error {
	vnet.Status.ID = res.ID
	vnet.Status.DeploymentID = res.DeploymentID
	vnet.Status.ProvisioningState = string(res.ProvisioningState)
	vnet.Spec.Tags = res.Tags

	var props VirtualNetworkSpecProperties
	if err := json.Unmarshal(res.Properties, &props); err != nil {
		return err
	}
	vnet.Spec.Properties = props
	return nil
}

func init() {
	SchemeBuilder.Register(&VirtualNetwork{}, &VirtualNetworkList{})
}
