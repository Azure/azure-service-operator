/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package compat

import (
	"testing"

	"github.com/Azure/azure-service-operator/v2/internal/util/to"

	. "github.com/onsi/gomega"

	v20231001s "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231001/storage"
)

func Test_UpgradeOverrideSettingsAssignPropertiesTo_WhenControlPlaneOverridesEmpty_ForceUpgradeIsFalse(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	upgradeOverrideSettings := UpgradeOverrideSettings{}

	// Act
	dest := &v20231001s.UpgradeOverrideSettings{}
	g.Expect(upgradeOverrideSettings.AssignPropertiesTo(dest)).To(Succeed())

	// Assert
	g.Expect(*dest.ForceUpgrade).To(BeFalse())
}

func Test_UpgradeOverrideSettingsAssignPropertiesTo_WhenControlPlaneOverridesSpecifiesIgnoreKubernetesDeprecations_ForceUpgradeIsTrue(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	upgradeOverrideSettings := UpgradeOverrideSettings{
		ControlPlaneOverrides: []string{ignoreKubernetesDeprecations},
	}

	// Act
	dest := &v20231001s.UpgradeOverrideSettings{}
	g.Expect(upgradeOverrideSettings.AssignPropertiesTo(dest)).To(Succeed())

	// Assert
	g.Expect(*dest.ForceUpgrade).To(BeTrue())
}

func Test_UpgradeOverrideSettingsAssignPropertiesFrom_WhenForceUpgradeIsNil_ControlePlaneOverridesIsEmpty(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	src := &v20231001s.UpgradeOverrideSettings{}

	// Act
	upgradeOverrideSettings := UpgradeOverrideSettings{}
	g.Expect(upgradeOverrideSettings.AssignPropertiesFrom(src)).To(Succeed())

	// Assert
	g.Expect(upgradeOverrideSettings.ControlPlaneOverrides).To(BeEmpty())
}

func Test_UpgradeOverrideSettingsAssignPropertiesFrom_WhenForceUpgradeIsFalse_ControlePlaneOverridesIsEmpty(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	src := &v20231001s.UpgradeOverrideSettings{
		ForceUpgrade: to.Ptr(false),
	}

	// Act
	upgradeOverrideSettings := UpgradeOverrideSettings{}
	g.Expect(upgradeOverrideSettings.AssignPropertiesFrom(src)).To(Succeed())

	// Assert
	g.Expect(upgradeOverrideSettings.ControlPlaneOverrides).To(BeEmpty())
}

func Test_UpgradeOverrideSettingsAssignPropertiesFrom_WhenForceUpgradeIsTrue_ControlePlaneOverridesContainsIgnoreKubernetesDeprecations(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Arrange
	src := &v20231001s.UpgradeOverrideSettings{
		ForceUpgrade: to.Ptr(true),
	}

	// Act
	upgradeOverrideSettings := UpgradeOverrideSettings{}
	g.Expect(upgradeOverrideSettings.AssignPropertiesFrom(src)).To(Succeed())

	// Assert
	g.Expect(upgradeOverrideSettings.ControlPlaneOverrides).To(ContainElement(ignoreKubernetesDeprecations))
}
