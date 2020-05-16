/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v1

import (
	azcorev1 "github.com/Azure/k8s-infra/apis/core/v1"
)

func (vnet *VirtualNetwork) GetResourceGroupObjectRef() *azcorev1.KnownTypeReference {
	return vnet.Spec.ResourceGroupRef
}

func (lb *LoadBalancer) GetResourceGroupObjectRef() *azcorev1.KnownTypeReference {
	return lb.Spec.ResourceGroupRef
}

func (rt *RouteTable) GetResourceGroupObjectRef() *azcorev1.KnownTypeReference {
	return rt.Spec.ResourceGroupRef
}

func (nsg *NetworkSecurityGroup) GetResourceGroupObjectRef() *azcorev1.KnownTypeReference {
	return nsg.Spec.ResourceGroupRef
}

func (*VirtualNetwork) ResourceType() string {
	return "Microsoft.Network/virtualNetworks"
}

func (*LoadBalancer) ResourceType() string {
	return "Microsoft.Network/loadBalancers"
}

func (*BackendAddressPool) ResourceType() string {
	return "Microsoft.Network/loadBalancers/backendAddressPools"
}

func (*FrontendIPConfiguration) ResourceType() string {
	return "Microsoft.Network/loadBalancers/frontendIPConfigurations"
}

func (*InboundNatRule) ResourceType() string {
	return "Microsoft.Network/loadBalancers/inboundNatRules"
}

func (*OutboundRule) ResourceType() string {
	return "Microsoft.Network/loadBalancers/outboundRules"
}

func (*LoadBalancingRule) ResourceType() string {
	return "Microsoft.Network/loadBalancers/loadBalancingRules"
}

func (*RouteTable) ResourceType() string {
	return "Microsoft.Network/routeTables"
}

func (*Route) ResourceType() string {
	return "Microsoft.Network/routeTables/routes"
}

func (*NetworkSecurityGroup) ResourceType() string {
	return "Microsoft.Network/networkSecurityGroups"
}

func (*SecurityRule) ResourceType() string {
	return "Microsoft.Network/networkSecurityGroups/securityRules"
}

func (*Subnet) ResourceType() string {
	return "Microsoft.Network/virtualNetworks/subnets"
}
