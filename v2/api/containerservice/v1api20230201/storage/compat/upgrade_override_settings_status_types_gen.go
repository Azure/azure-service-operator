// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package compat

import (
	v20230701s "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230701/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/pkg/errors"
)

// Storage version of v1api20230202preview.UpgradeOverrideSettings_STATUS
// Settings for overrides when upgrading a cluster.
type UpgradeOverrideSettings_STATUS struct {
	ControlPlaneOverrides []string               `json:"controlPlaneOverrides,omitempty"`
	PropertyBag           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Until                 *string                `json:"until,omitempty"`
}

// AssignProperties_From_UpgradeOverrideSettings_STATUS populates our UpgradeOverrideSettings_STATUS from the provided source UpgradeOverrideSettings_STATUS
func (settings *UpgradeOverrideSettings_STATUS) AssignProperties_From_UpgradeOverrideSettings_STATUS(source *v20230701s.UpgradeOverrideSettings_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// ControlPlaneOverrides
	if propertyBag.Contains("ControlPlaneOverrides") {
		var controlPlaneOverride []string
		err := propertyBag.Pull("ControlPlaneOverrides", &controlPlaneOverride)
		if err != nil {
			return errors.Wrap(err, "pulling 'ControlPlaneOverrides' from propertyBag")
		}

		settings.ControlPlaneOverrides = controlPlaneOverride
	} else {
		settings.ControlPlaneOverrides = nil
	}

	// ForceUpgrade
	if source.ForceUpgrade != nil {
		propertyBag.Add("ForceUpgrade", *source.ForceUpgrade)
	} else {
		propertyBag.Remove("ForceUpgrade")
	}

	// Until
	settings.Until = genruntime.ClonePointerToString(source.Until)

	// Update the property bag
	if len(propertyBag) > 0 {
		settings.PropertyBag = propertyBag
	} else {
		settings.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignProperties_To_UpgradeOverrideSettings_STATUS populates the provided destination UpgradeOverrideSettings_STATUS from our UpgradeOverrideSettings_STATUS
func (settings *UpgradeOverrideSettings_STATUS) AssignProperties_To_UpgradeOverrideSettings_STATUS(destination *v20230701s.UpgradeOverrideSettings_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(settings.PropertyBag)

	// ControlPlaneOverrides
	if len(settings.ControlPlaneOverrides) > 0 {
		propertyBag.Add("ControlPlaneOverrides", settings.ControlPlaneOverrides)
	} else {
		propertyBag.Remove("ControlPlaneOverrides")
	}

	// ForceUpgrade
	if propertyBag.Contains("ForceUpgrade") {
		var forceUpgrade bool
		err := propertyBag.Pull("ForceUpgrade", &forceUpgrade)
		if err != nil {
			return errors.Wrap(err, "pulling 'ForceUpgrade' from propertyBag")
		}

		destination.ForceUpgrade = &forceUpgrade
	} else {
		destination.ForceUpgrade = nil
	}

	// Until
	destination.Until = genruntime.ClonePointerToString(settings.Until)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}
