/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v20191101

import (
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/conversion"

	v1 "github.com/Azure/k8s-infra/apis/microsoft.network/v1"
)

const (
	apiVersion = "2019-11-01"
)

func (src *BackendAddressPool) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.BackendAddressPool)

	if err := Convert_v20191101_BackendAddressPool_To_v1_BackendAddressPool(src, dst, nil); err != nil {
		return err
	}

	dst.Spec.APIVersion = apiVersion
	return nil
}

func (dst *BackendAddressPool) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.BackendAddressPool)

	if err := Convert_v1_BackendAddressPool_To_v20191101_BackendAddressPool(src, dst, nil); err != nil {
		return err
	}

	return nil
}

func (src *FrontendIPConfiguration) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.FrontendIPConfiguration)

	if err := Convert_v20191101_FrontendIPConfiguration_To_v1_FrontendIPConfiguration(src, dst, nil); err != nil {
		return err
	}

	dst.Spec.APIVersion = apiVersion
	return nil
}

func (dst *FrontendIPConfiguration) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.FrontendIPConfiguration)

	if err := Convert_v1_FrontendIPConfiguration_To_v20191101_FrontendIPConfiguration(src, dst, nil); err != nil {
		return err
	}

	return nil
}

func (src *InboundNatRule) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.InboundNatRule)

	if err := Convert_v20191101_InboundNatRule_To_v1_InboundNatRule(src, dst, nil); err != nil {
		return err
	}

	dst.Spec.APIVersion = apiVersion
	return nil
}

func (dst *InboundNatRule) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.InboundNatRule)

	if err := Convert_v1_InboundNatRule_To_v20191101_InboundNatRule(src, dst, nil); err != nil {
		return err
	}

	return nil
}

func (src *LoadBalancer) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.LoadBalancer)

	if err := Convert_v20191101_LoadBalancer_To_v1_LoadBalancer(src, dst, nil); err != nil {
		return err
	}

	dst.Spec.APIVersion = apiVersion
	return nil
}

func (dst *LoadBalancer) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.LoadBalancer)

	if err := Convert_v1_LoadBalancer_To_v20191101_LoadBalancer(src, dst, nil); err != nil {
		return err
	}

	return nil
}

func (src *LoadBalancingRule) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.LoadBalancingRule)

	if err := Convert_v20191101_LoadBalancingRule_To_v1_LoadBalancingRule(src, dst, nil); err != nil {
		return err
	}

	dst.Spec.APIVersion = apiVersion
	return nil
}

func (dst *LoadBalancingRule) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.LoadBalancingRule)

	if err := Convert_v1_LoadBalancingRule_To_v20191101_LoadBalancingRule(src, dst, nil); err != nil {
		return err
	}

	return nil
}

func (src *NetworkInterfaceIPConfiguration) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.NetworkInterfaceIPConfiguration)

	if err := Convert_v20191101_NetworkInterfaceIPConfiguration_To_v1_NetworkInterfaceIPConfiguration(src, dst, nil); err != nil {
		return err
	}

	dst.Spec.APIVersion = apiVersion
	return nil
}

func (dst *NetworkInterfaceIPConfiguration) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.NetworkInterfaceIPConfiguration)

	if err := Convert_v1_NetworkInterfaceIPConfiguration_To_v20191101_NetworkInterfaceIPConfiguration(src, dst, nil); err != nil {
		return err
	}

	return nil
}

func (src *NetworkSecurityGroup) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.NetworkSecurityGroup)

	if err := Convert_v20191101_NetworkSecurityGroup_To_v1_NetworkSecurityGroup(src, dst, nil); err != nil {
		return err
	}

	dst.Spec.APIVersion = apiVersion
	return nil
}

func (dst *NetworkSecurityGroup) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.NetworkSecurityGroup)

	if err := Convert_v1_NetworkSecurityGroup_To_v20191101_NetworkSecurityGroup(src, dst, nil); err != nil {
		return err
	}

	return nil
}

func (src *OutboundRule) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.OutboundRule)

	if err := Convert_v20191101_OutboundRule_To_v1_OutboundRule(src, dst, nil); err != nil {
		return err
	}

	dst.Spec.APIVersion = apiVersion
	return nil
}

func (dst *OutboundRule) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.OutboundRule)

	if err := Convert_v1_OutboundRule_To_v20191101_OutboundRule(src, dst, nil); err != nil {
		return err
	}

	return nil
}

func (src *Route) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.Route)

	if err := Convert_v20191101_Route_To_v1_Route(src, dst, nil); err != nil {
		return err
	}

	dst.Spec.APIVersion = apiVersion
	return nil
}

func (dst *Route) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.Route)

	if err := Convert_v1_Route_To_v20191101_Route(src, dst, nil); err != nil {
		return err
	}

	return nil
}

func (src *RouteTable) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.RouteTable)

	if err := Convert_v20191101_RouteTable_To_v1_RouteTable(src, dst, nil); err != nil {
		return err
	}

	dst.Spec.APIVersion = apiVersion
	return nil
}

func (dst *RouteTable) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.RouteTable)

	if err := Convert_v1_RouteTable_To_v20191101_RouteTable(src, dst, nil); err != nil {
		return err
	}

	return nil
}

func (src *SecurityRule) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.SecurityRule)

	if err := Convert_v20191101_SecurityRule_To_v1_SecurityRule(src, dst, nil); err != nil {
		return err
	}

	dst.Spec.APIVersion = apiVersion
	return nil
}

func (dst *SecurityRule) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.SecurityRule)

	if err := Convert_v1_SecurityRule_To_v20191101_SecurityRule(src, dst, nil); err != nil {
		return err
	}

	return nil
}

func (src *VirtualNetwork) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.VirtualNetwork)

	if err := Convert_v20191101_VirtualNetwork_To_v1_VirtualNetwork(src, dst, nil); err != nil {
		fmt.Println(err)
		return err
	}

	dst.Spec.APIVersion = apiVersion
	return nil
}

func (dst *VirtualNetwork) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.VirtualNetwork)

	if err := Convert_v1_VirtualNetwork_To_v20191101_VirtualNetwork(src, dst, nil); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
