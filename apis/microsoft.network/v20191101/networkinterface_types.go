/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v20191101

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	azcorev1 "github.com/Azure/k8s-infra/apis/core/v1"
)

type (
	InterfaceDNSSettings struct {
		// DNSServers - List of DNS servers IP addresses. Use 'AzureProvidedDNS' to switch to azure provided DNS resolution. 'AzureProvidedDNS' value cannot be combined with other IPs, it must be the only value in dnsServers collection.
		DNSServers *[]string `json:"dnsServers,omitempty"`
		// InternalDNSNameLabel - Relative DNS name for this NIC used for internal communications between VMs in the same virtual network.
		InternalDNSNameLabel *string `json:"internalDnsNameLabel,omitempty"`
	}

	NetworkInterfaceIPConfigurationSpecProperties struct {
		Primary          bool   `json:"primary,omitempty"`
		PrivateIPAddress string `json:"privateIPAddress,omitempty"`
		// +kubebuilder:validation:Enum=IPv4;IPv6
		PrivateIPAddressVersion string `json:"privateIPAddressVersion,omitempty"`
		// +kubebuilder:validation:Enum=Dynamic;Static
		PrivateIPAllocationMethod string                       `json:"privateIPAllocationMethod,omitempty"`
		PublicIPAddressRef        *azcorev1.KnownTypeReference `json:"publicIPAddressRef,omitempty"`
		SubnetRef                 *azcorev1.KnownTypeReference `json:"subnetRef,omitempty"`
	}

	// NetworkInterfaceIPConfigurationSpec defines the desired state of NetworkInterfaceIPConfiguration
	NetworkInterfaceIPConfigurationSpec struct {
		Name       string                                         `json:"name"`
		Properties *NetworkInterfaceIPConfigurationSpecProperties `json:"properties,omitempty"`
	}

	NetworkInterfaceSpecProperties struct {
		// NetworkSecurityGroup - The reference to the NetworkSecurityGroup resource.
		NetworkSecurityGroupRef *azcorev1.KnownTypeReference `json:"networkSecurityGroupRef,omitempty"`
		// IPConfigurations - A list of IPConfigurations of the network interface.
		IPConfigurations []NetworkInterfaceIPConfigurationSpec `json:"ipConfigurations,omitempty"`
		// DNSSettings - The DNS settings in network interface.
		DNSSettings *InterfaceDNSSettings `json:"dnsSettings,omitempty"`
		// EnableAcceleratedNetworking - If the network interface is accelerated networking enabled.
		EnableAcceleratedNetworking *bool `json:"enableAcceleratedNetworking,omitempty"`
		// EnableIPForwarding - Indicates whether IP forwarding is enabled on this network interface.
		EnableIPForwarding *bool `json:"enableIPForwarding,omitempty"`
	}

	// NetworkInterfaceSpec defines the desired state of NetworkInterface
	NetworkInterfaceSpec struct {
		// ResourceGroupRef is the Azure Resource Group the VirtualNetwork resides within
		// +kubebuilder:validation:Required
		ResourceGroupRef *azcorev1.KnownTypeReference `json:"resourceGroupRef" group:"microsoft.resources.infra.azure.com" kind:"ResourceGroup"`

		// Location of the VNET in Azure
		// +kubebuilder:validation:Required
		Location string `json:"location"`

		// Tags are user defined key value pairs
		// +optional
		Tags map[string]string `json:"tags,omitempty"`

		// Properties of the Virtual Network
		Properties *NetworkInterfaceSpecProperties `json:"properties,omitempty"`
	}

	// NetworkInterfaceStatus defines the observed state of NetworkInterface
	NetworkInterfaceStatus struct {
		ID                string `json:"id,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty"`
	}

	// +kubebuilder:object:root=true

	// NetworkInterface is the Schema for the networkinterfaces API
	NetworkInterface struct {
		metav1.TypeMeta   `json:",inline"`
		metav1.ObjectMeta `json:"metadata,omitempty"`

		Spec   NetworkInterfaceSpec   `json:"spec,omitempty"`
		Status NetworkInterfaceStatus `json:"status,omitempty"`
	}

	// +kubebuilder:object:root=true

	// NetworkInterfaceList contains a list of NetworkInterface
	NetworkInterfaceList struct {
		metav1.TypeMeta `json:",inline"`
		metav1.ListMeta `json:"metadata,omitempty"`
		Items           []NetworkInterface `json:"items"`
	}
)

func init() {
	SchemeBuilder.Register(&NetworkInterface{}, &NetworkInterfaceList{})
}
