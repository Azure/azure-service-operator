// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	"github.com/Azure/azure-service-operator/api/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (src *AzureSqlDatabase) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.AzureSqlDatabase)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.ResourceGroup = src.Spec.ResourceGroup
	dst.Spec.Location = src.Spec.Location
	dst.Spec.Server = src.Spec.Server
	dst.Spec.Edition = v1beta1.DBEdition(src.Spec.Edition)
	dst.Spec.DbName = src.Spec.DbName

	// Status
	dst.Status = v1beta1.ASOStatus(src.Status)

	return nil
}

func (dst *AzureSqlDatabase) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.AzureSqlDatabase)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.ResourceGroup = src.Spec.ResourceGroup
	dst.Spec.Location = src.Spec.Location
	dst.Spec.Server = src.Spec.Server
	dst.Spec.Edition = DBEdition(src.Spec.Edition)
	dst.Spec.DbName = src.Spec.DbName

	// Status
	dst.Status = ASOStatus(src.Status)

	return nil

}
