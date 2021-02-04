// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	"github.com/Azure/azure-service-operator/api/v1alpha2"
)

var _ conversion.Convertible = &MySQLAADUser{}

func (src *MySQLUser) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1alpha2.MySQLUser)

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.ResourceGroup = src.Spec.ResourceGroup
	dst.Spec.Server = src.Spec.Server
	dst.Spec.AdminSecret = src.Spec.AdminSecret
	dst.Spec.AdminSecretKeyVault = src.Spec.AdminSecretKeyVault
	dst.Spec.Username = src.Spec.Username
	dst.Spec.KeyVaultToStoreSecrets = src.Spec.KeyVaultToStoreSecrets

	// v1alpha1 doesn't support server-level roles, only
	// database-level ones, so we move them into the new field.
	dst.Spec.Roles = []string{}
	dst.Spec.DatabaseRoles = make(map[string][]string)
	dst.Spec.DatabaseRoles[src.Spec.DbName] = append([]string(nil), src.Spec.Roles...)

	// Status
	dst.Status = v1alpha2.ASOStatus(src.Status)

	return nil
}

func (dst *MySQLUser) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1alpha2.MySQLUser)

	// Converting a v1alpha2 user into a v1alpha1 one is only allowed
	// if it has exactly one database...
	if len(src.Spec.DatabaseRoles) != 1 {
		return errors.Errorf("can't convert user %q to %T because it has privileges in %d databases", src.ObjectMeta.Name, dst, len(src.Spec.DatabaseRoles))
	}
	// ...and no server-level roles.
	if len(src.Spec.Roles) != 0 {
		return errors.Errorf("can't convert user %q to %T because it has server-level roles", src.ObjectMeta.Name, dst)
	}

	// ObjectMeta
	dst.ObjectMeta = src.ObjectMeta

	// Spec
	dst.Spec.ResourceGroup = src.Spec.ResourceGroup
	dst.Spec.Server = src.Spec.Server
	dst.Spec.AdminSecret = src.Spec.AdminSecret
	dst.Spec.AdminSecretKeyVault = src.Spec.AdminSecretKeyVault
	dst.Spec.Username = src.Spec.Username
	dst.Spec.KeyVaultToStoreSecrets = src.Spec.KeyVaultToStoreSecrets

	for dbName, roles := range src.Spec.DatabaseRoles {
		dst.Spec.DbName = dbName
		dst.Spec.Roles = append(dst.Spec.Roles, roles...)
		break
	}

	// Status
	dst.Status = ASOStatus(src.Status)

	return nil

}
