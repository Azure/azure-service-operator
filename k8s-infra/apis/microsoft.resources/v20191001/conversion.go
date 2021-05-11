/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v20191001

import (
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	v1 "github.com/Azure/k8s-infra/apis/microsoft.resources/v1"
)

func (src *ResourceGroup) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.ResourceGroup)

	if err := Convert_v20191001_ResourceGroup_To_v1_ResourceGroup(src, dst, nil); err != nil {
		return err
	}

	dst.Spec.APIVersion = "2019-10-01"
	return nil
}

func (dst *ResourceGroup) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.ResourceGroup)

	if err := Convert_v1_ResourceGroup_To_v20191001_ResourceGroup(src, dst, nil); err != nil {
		return err
	}

	return nil
}
