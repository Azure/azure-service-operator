/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v1

import (
	azcorev1 "github.com/Azure/k8s-infra/apis/core/v1"
)

func (vm *VirtualMachine) GetResourceGroupObjectRef() *azcorev1.KnownTypeReference {
	return vm.Spec.ResourceGroupRef
}

func (vmss *VirtualMachineScaleSet) GetResourceGroupObjectRef() *azcorev1.KnownTypeReference {
	return vmss.Spec.ResourceGroupRef
}

func (*VirtualMachine) ResourceType() string {
	return "Microsoft.Compute/virtualMachines"
}

func (*VirtualMachineScaleSet) ResourceType() string {
	return "Microsoft.Compute/virtualMachineScaleSets"
}
