// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	storage "github.com/Azure/azure-service-operator/v2/api/redhatopenshift/v1api20251223preview/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// augmentConversionForKmsEncryptionProfile implementation
// In v20240610preview, VaultName lives on KmsKey (inside ActiveKey).
// In v20251223preview (hub), VaultName moved to KmsEncryptionProfile level.
// The generated conversion puts hub's VaultName into PropertyBag,
// but we need to remap it to/from ActiveKey.VaultName.

var _ augmentConversionForKmsEncryptionProfile = &KmsEncryptionProfile{}

// AssignPropertiesFrom (hub → spoke): pull VaultName from PropertyBag and set on ActiveKey.VaultName
func (profile *KmsEncryptionProfile) AssignPropertiesFrom(src *storage.KmsEncryptionProfile) error {
	if src.VaultName != nil && profile.ActiveKey != nil {
		vaultName := *src.VaultName
		profile.ActiveKey.VaultName = &vaultName

		// Remove from PropertyBag since we've placed it on ActiveKey
		if profile.PropertyBag != nil {
			propertyBag := genruntime.NewPropertyBag(profile.PropertyBag)
			propertyBag.Remove("VaultName")
			if len(propertyBag) > 0 {
				profile.PropertyBag = propertyBag
			} else {
				profile.PropertyBag = nil
			}
		}
	}

	return nil
}

// AssignPropertiesTo (spoke → hub): copy ActiveKey.VaultName to hub's KmsEncryptionProfile.VaultName
func (profile *KmsEncryptionProfile) AssignPropertiesTo(dst *storage.KmsEncryptionProfile) error {
	if profile.ActiveKey != nil && profile.ActiveKey.VaultName != nil {
		vaultName := *profile.ActiveKey.VaultName
		dst.VaultName = &vaultName

		// Remove VaultName from hub KmsKey's PropertyBag (generated KmsKey conversion put it there)
		if dst.ActiveKey != nil && dst.ActiveKey.PropertyBag != nil {
			propertyBag := genruntime.NewPropertyBag(dst.ActiveKey.PropertyBag)
			propertyBag.Remove("VaultName")
			if len(propertyBag) > 0 {
				dst.ActiveKey.PropertyBag = propertyBag
			} else {
				dst.ActiveKey.PropertyBag = nil
			}
		}
	}

	return nil
}

// augmentConversionForKmsEncryptionProfile_STATUS implementation
var _ augmentConversionForKmsEncryptionProfile_STATUS = &KmsEncryptionProfile_STATUS{}

// AssignPropertiesFrom (hub → spoke): pull VaultName from PropertyBag and set on ActiveKey.VaultName
func (profile *KmsEncryptionProfile_STATUS) AssignPropertiesFrom(src *storage.KmsEncryptionProfile_STATUS) error {
	if src.VaultName != nil && profile.ActiveKey != nil {
		vaultName := *src.VaultName
		profile.ActiveKey.VaultName = &vaultName

		if profile.PropertyBag != nil {
			propertyBag := genruntime.NewPropertyBag(profile.PropertyBag)
			propertyBag.Remove("VaultName")
			if len(propertyBag) > 0 {
				profile.PropertyBag = propertyBag
			} else {
				profile.PropertyBag = nil
			}
		}
	}

	return nil
}

// AssignPropertiesTo (spoke → hub): copy ActiveKey.VaultName to hub's KmsEncryptionProfile.VaultName
func (profile *KmsEncryptionProfile_STATUS) AssignPropertiesTo(dst *storage.KmsEncryptionProfile_STATUS) error {
	if profile.ActiveKey != nil && profile.ActiveKey.VaultName != nil {
		vaultName := *profile.ActiveKey.VaultName
		dst.VaultName = &vaultName

		if dst.ActiveKey != nil && dst.ActiveKey.PropertyBag != nil {
			propertyBag := genruntime.NewPropertyBag(dst.ActiveKey.PropertyBag)
			propertyBag.Remove("VaultName")
			if len(propertyBag) > 0 {
				dst.ActiveKey.PropertyBag = propertyBag
			} else {
				dst.ActiveKey.PropertyBag = nil
			}
		}
	}

	return nil
}
