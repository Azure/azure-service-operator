/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v20191201

import (
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	v1 "github.com/Azure/k8s-infra/apis/microsoft.compute/v1"
)

const (
	apiVersion = "2019-12-01"
)

func (src *VirtualMachine) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.VirtualMachine)

	if err := Convert_v20191201_VirtualMachine_To_v1_VirtualMachine(src, dst, nil); err != nil {
		return err
	}

	dst.Spec.APIVersion = apiVersion
	return nil
}

func (dst *VirtualMachine) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.VirtualMachine)

	if err := Convert_v1_VirtualMachine_To_v20191201_VirtualMachine(src, dst, nil); err != nil {
		return err
	}

	return nil
}

func (src *VirtualMachineScaleSet) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.VirtualMachineScaleSet)

	if err := Convert_v20191201_VirtualMachineScaleSet_To_v1_VirtualMachineScaleSet(src, dst, nil); err != nil {
		return err
	}

	dst.Spec.APIVersion = apiVersion
	return nil
}

func (dst *VirtualMachineScaleSet) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.VirtualMachineScaleSet)

	if err := Convert_v1_VirtualMachineScaleSet_To_v20191201_VirtualMachineScaleSet(src, dst, nil); err != nil {
		return err
	}

	return nil
}
