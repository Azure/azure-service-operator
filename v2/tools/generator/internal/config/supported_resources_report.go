/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"path/filepath"
)

// SupportedResourcesReport is configuration for the report that lists all the supported resources.
type SupportedResourcesReport struct {
	cfg *Configuration // Back reference to global configuration

	OutputPath   string `yaml:"outputPath,omitempty"`   // Destination filepath for the report, relative to DestinationGoModuleFile
	Introduction string `yaml:"introduction,omitempty"` // Introduction to the report
	// ResourceUrlTemplate is a template for URL to the API docs for a resource
	// It may use the placeholders {group} {version} and {kind}
	ResourceUrlTemplate string `yaml:"resourceUrlTemplate"`
	// ResourcePathTemplate is a template used for generating a file path for checking whether docs for a resource have been generated
	// specified relative to the directory of outputPath
	ResourcePathTemplate string `yaml:"resourcePathTemplate"`
}

// NewSupportedResourcesReport creates a new SupportedResourcesReport.
func NewSupportedResourcesReport(cfg *Configuration) *SupportedResourcesReport {
	return &SupportedResourcesReport{
		cfg: cfg,
	}
}

// FullOutputPath returns the fully qualified path to the output file.
func (srr *SupportedResourcesReport) FullOutputPath() string {
	return filepath.Join(
		filepath.Dir(srr.cfg.DestinationGoModuleFile),
		srr.OutputPath)
}
