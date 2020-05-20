// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	"github.com/Azure/azure-service-operator/api/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (src *AzureSqlFailoverGroup) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.AzureSqlFailoverGroup)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.ResourceGroup = src.Spec.ResourceGroup
	dst.Spec.Location = src.Spec.Location
	dst.Spec.Server = src.Spec.Server
	dst.Spec.FailoverGracePeriod = src.Spec.FailoverGracePeriod
	dst.Spec.FailoverPolicy = v1beta1.ReadWriteEndpointFailoverPolicy(src.Spec.FailoverPolicy)
	dst.Spec.SecondaryServer = src.Spec.SecondaryServer
	dst.Spec.SecondaryServerResourceGroup = src.Spec.SecondaryServerResourceGroup
	dst.Spec.DatabaseList = src.Spec.DatabaseList
	dst.Spec.KeyVaultToStoreSecrets = src.Spec.KeyVaultToStoreSecrets

	// Status
	dst.Status = v1beta1.ASOStatus(src.Status)

	return nil
}

func (dst *AzureSqlFailoverGroup) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.AzureSqlFailoverGroup)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.ResourceGroup = src.Spec.ResourceGroup
	dst.Spec.Location = src.Spec.Location
	dst.Spec.Server = src.Spec.Server
	dst.Spec.FailoverGracePeriod = src.Spec.FailoverGracePeriod
	dst.Spec.FailoverPolicy = ReadWriteEndpointFailoverPolicy(src.Spec.FailoverPolicy)
	dst.Spec.SecondaryServer = src.Spec.SecondaryServer
	dst.Spec.SecondaryServerResourceGroup = src.Spec.SecondaryServerResourceGroup
	dst.Spec.DatabaseList = src.Spec.DatabaseList
	dst.Spec.KeyVaultToStoreSecrets = src.Spec.KeyVaultToStoreSecrets

	// Status
	dst.Status = ASOStatus(src.Status)

	return nil

}
