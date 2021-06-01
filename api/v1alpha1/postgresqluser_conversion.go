// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/conversion"

	"github.com/Azure/azure-service-operator/api/v1alpha2"
)

var _ conversion.Convertible = &PostgreSQLUser{}

// convert a v1alpha1 to a v1alpha2
func (src *PostgreSQLUser) ConvertTo(dstRaw conversion.Hub) error {
	fmt.Sprintln("PostgreSQLUser v1alpha1 to a v1alpha2")
	dst := dstRaw.(*v1alpha2.PostgreSQLUser)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.ResourceGroup = src.Spec.ResourceGroup
	dst.Spec.Server = src.Spec.Server
	dst.Spec.AdminSecret = src.Spec.AdminSecret
	dst.Spec.AdminSecretKeyVault = src.Spec.AdminSecretKeyVault
	dst.Spec.Username = src.Spec.Username
	dst.Spec.KeyVaultToStoreSecrets = src.Spec.KeyVaultToStoreSecrets
	dst.Spec.Roles = src.Spec.Roles

	if dst.Spec.OwnedDatabases == nil {
		dst.Spec.OwnedDatabases = []string{}
	}

	// Use the DbName attribute as MaintenanceDatabase, as we can be sure
	// this database exists.
	dst.Spec.MaintenanceDatabase = src.Spec.DbName

	// Status
	dst.Status = v1alpha2.ASOStatus(src.Status)

	return nil
}

// convert from a v1alpha2 to a v1alpha1
func (dst *PostgreSQLUser) ConvertFrom(srcRaw conversion.Hub) error {
	fmt.Sprintln("PostgreSQLUser v1alpha2 to a v1alpha1")
	src := srcRaw.(*v1alpha2.PostgreSQLUser)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.ResourceGroup = src.Spec.ResourceGroup
	dst.Spec.Server = src.Spec.Server
	dst.Spec.AdminSecret = src.Spec.AdminSecret
	dst.Spec.AdminSecretKeyVault = src.Spec.AdminSecretKeyVault
	dst.Spec.Username = src.Spec.Username
	dst.Spec.KeyVaultToStoreSecrets = src.Spec.KeyVaultToStoreSecrets

	dst.Spec.Roles = src.Spec.Roles

	if len(src.Spec.MaintenanceDatabase) == 0 {
		// fallback to the default maintenance database.
		dst.Spec.DbName = "postgres"
	} else {
		dst.Spec.DbName = src.Spec.MaintenanceDatabase
	}

	// Status
	dst.Status = ASOStatus(src.Status)
	return nil
}
