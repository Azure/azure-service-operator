/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v20150101

import (
	cvt "k8s.io/apimachinery/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	v1 "github.com/Azure/k8s-infra/apis/microsoft.resources/v1"
)

func (src *ResourceGroup) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.ResourceGroup)

	if err := Convert_v20150101_ResourceGroup_To_v1_ResourceGroup(src, dst, nil); err != nil {
		return err
	}

	dst.Spec.APIVersion = "2015-01-01"
	return nil
}

func (dst *ResourceGroup) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.ResourceGroup)

	if err := Convert_v1_ResourceGroup_To_v20150101_ResourceGroup(src, dst, nil); err != nil {
		return err
	}

	return nil
}

// Convert_v1_ResourceGroupSpec_To_v20150101_ResourceGroupSpec is required because we are unable to project MangedBy
// into this version
// noinspection GoSnakeCaseUsage
func Convert_v1_ResourceGroupSpec_To_v20150101_ResourceGroupSpec(in *v1.ResourceGroupSpec, out *ResourceGroupSpec, s cvt.Scope) error {
	return autoConvert_v1_ResourceGroupSpec_To_v20150101_ResourceGroupSpec(in, out, s)
}
