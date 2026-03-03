/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"path/filepath"
)

// UpgradableResourcesReport is configuration for the report that identifies resources with available upgrades.
type UpgradableResourcesReport struct {
	cfg *Configuration // Back reference to global configuration
	// OutputFolder is the destination folder for the report, relative to DestinationGoModuleFile
	OutputFolder string `yaml:"outputFolder,omitempty"`
	// StableVersionsExpiry is the number of months after which a stable upgrade is recommended
	StableVersionsExpiry int `yaml:"stableVersionsExpiry"`
	// PreviewVersionsExpiry is the number of months after which a preview upgrade is recommended
	PreviewVersionsExpiry int `yaml:"previewVersionsExpiry"`
	// LatestVersionThreshold is the number of months after which we always recommend upgrading to the latest version
	LatestVersionThreshold int `yaml:"latestVersionThreshold"`
}

// NewUpgradableResourcesReport creates a new UpgradableResourcesReport with default thresholds.
func NewUpgradableResourcesReport(cfg *Configuration) *UpgradableResourcesReport {
	return &UpgradableResourcesReport{
		cfg:                    cfg,
		StableVersionsExpiry:   12,
		PreviewVersionsExpiry:  6,
		LatestVersionThreshold: 24,
	}
}

// FullOutputPath returns the fully qualified path to the output file.
func (urr *UpgradableResourcesReport) FullOutputPath() string {
	return filepath.Join(
		filepath.Dir(urr.cfg.DestinationGoModuleFile),
		urr.OutputFolder,
		"upgradable_resources.md")
}
