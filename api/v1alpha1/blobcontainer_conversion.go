// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	"github.com/Azure/azure-service-operator/api/v1alpha2"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (src *BlobContainer) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha2.BlobContainer)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.ResourceGroup = src.Spec.ResourceGroup
	dst.Spec.Location = src.Spec.Location
	dst.Spec.AccountName = src.Spec.AccountName
	dst.Spec.AccessLevel = src.Spec.AccessLevel

	// Status
	dst.Status = v1alpha2.ASOStatus(src.Status)

	return nil
}

func (dst *BlobContainer) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha2.BlobContainer)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.ResourceGroup = src.Spec.ResourceGroup
	dst.Spec.Location = src.Spec.Location
	dst.Spec.AccountName = src.Spec.AccountName
	dst.Spec.AccessLevel = src.Spec.AccessLevel

	// Status
	dst.Status = ASOStatus(src.Status)

	return nil

}
