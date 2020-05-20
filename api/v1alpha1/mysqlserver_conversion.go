// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	"github.com/Azure/azure-service-operator/api/v1alpha2"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (src *MySQLServer) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha2.MySQLServer)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.ResourceGroup = src.Spec.ResourceGroup
	dst.Spec.Location = src.Spec.Location
	dst.Spec.Sku = v1alpha2.AzureDBsSQLSku{
		Name:     src.Spec.Sku.Name,
		Tier:     v1alpha2.SkuTier(src.Spec.Sku.Tier),
		Capacity: src.Spec.Sku.Capacity,
		Size:     src.Spec.Sku.Size,
		Family:   src.Spec.Sku.Family,
	}

	dst.Spec.ServerVersion = v1alpha2.ServerVersion(src.Spec.ServerVersion)
	dst.Spec.SSLEnforcement = v1alpha2.SslEnforcementEnum(src.Spec.SSLEnforcement)
	dst.Spec.CreateMode = src.Spec.CreateMode
	dst.Spec.ReplicaProperties = v1alpha2.ReplicaProperties(src.Spec.ReplicaProperties)
	dst.Spec.KeyVaultToStoreSecrets = src.Spec.KeyVaultToStoreSecrets

	// New Spec
	dst.Spec.StorageProfile = nil

	// Status
	dst.Status = v1alpha2.ASOStatus(src.Status)

	return nil
}

func (dst *MySQLServer) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha2.MySQLServer)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.ResourceGroup = src.Spec.ResourceGroup
	dst.Spec.Location = src.Spec.Location
	dst.Spec.Sku = AzureDBsSQLSku{
		Name:     src.Spec.Sku.Name,
		Tier:     SkuTier(src.Spec.Sku.Tier),
		Capacity: src.Spec.Sku.Capacity,
		Size:     src.Spec.Sku.Size,
		Family:   src.Spec.Sku.Family,
	}
	dst.Spec.ServerVersion = ServerVersion(src.Spec.ServerVersion)
	dst.Spec.SSLEnforcement = SslEnforcementEnum(src.Spec.SSLEnforcement)
	dst.Spec.CreateMode = src.Spec.CreateMode
	dst.Spec.ReplicaProperties = ReplicaProperties(src.Spec.ReplicaProperties)
	dst.Spec.KeyVaultToStoreSecrets = src.Spec.KeyVaultToStoreSecrets

	// Status
	dst.Status = ASOStatus(src.Status)

	return nil

}
