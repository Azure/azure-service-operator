// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha2

import (
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (src *MySQLServer) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha1.MySQLServer)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.ResourceGroup = src.Spec.ResourceGroup
	dst.Spec.Location = src.Spec.Location
	dst.Spec.Sku = src.Spec.Sku
	dst.Spec.ServerVersion = src.Spec.ServerVersion
	dst.Spec.SSLEnforcement = src.Spec.SSLEnforcement
	dst.Spec.CreateMode = src.Spec.CreateMode
	dst.Spec.ReplicaProperties = src.Spec.ReplicaProperties
	dst.Spec.KeyVaultToStoreSecrets = src.Spec.KeyVaultToStoreSecrets

	// New Spec
	//dst.Spec.StorageProfile = nil

	// Status
	dst.Status = src.Status

	return nil
}

func (dst *MySQLServer) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha1.MySQLServer)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.ResourceGroup = src.Spec.ResourceGroup
	dst.Spec.Location = src.Spec.Location
	dst.Spec.Sku = src.Spec.Sku
	dst.Spec.ServerVersion = src.Spec.ServerVersion
	dst.Spec.SSLEnforcement = src.Spec.SSLEnforcement
	dst.Spec.CreateMode = src.Spec.CreateMode
	dst.Spec.ReplicaProperties = src.Spec.ReplicaProperties
	dst.Spec.KeyVaultToStoreSecrets = src.Spec.KeyVaultToStoreSecrets

	// Status
	dst.Status = src.Status

	return nil

}
