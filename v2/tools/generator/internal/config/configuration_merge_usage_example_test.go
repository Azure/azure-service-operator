/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

// TestConfiguration_Merge_UsageExample demonstrates how to use the merge functionality
func TestConfiguration_Merge_UsageExample(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Create a base configuration
	base := NewConfiguration()
	base.SchemaRoot = "schemas/"
	base.Pipeline = GenerationPipelineAzure
	base.AnyTypePackages = []string{"base-package1", "base-package2"}
	base.EmitDocFiles = false

	// Create another configuration to merge
	other := NewConfiguration()
	other.TypesOutputPath = "api/types" // This was empty in base, so it will be set
	other.AnyTypePackages = []string{"other-package1", "other-package2"} // These will be appended
	other.SamplesPath = "samples"

	// Create configuration that demonstrates nested merging
	other.SupportedResourcesReport = NewSupportedResourcesReport(other)
	other.SupportedResourcesReport.OutputFolder = "reports"
	other.SupportedResourcesReport.CurrentRelease = "v2.1.0"

	// Perform the merge
	err := base.Merge(other)
	g.Expect(err).To(Succeed())

	// Verify the results
	g.Expect(base.SchemaRoot).To(Equal("schemas/")) // Preserved from base
	g.Expect(base.Pipeline).To(Equal(GenerationPipelineAzure)) // Preserved from base
	g.Expect(base.TypesOutputPath).To(Equal("api/types")) // Set from other
	g.Expect(base.SamplesPath).To(Equal("samples")) // Set from other
	g.Expect(base.EmitDocFiles).To(BeFalse()) // Preserved from base (false)
	
	// Verify slice was appended
	g.Expect(base.AnyTypePackages).To(Equal([]string{
		"base-package1", "base-package2", 
		"other-package1", "other-package2",
	}))

	// Verify nested configuration was merged
	g.Expect(base.SupportedResourcesReport).ToNot(BeNil())
	g.Expect(base.SupportedResourcesReport.OutputFolder).To(Equal("reports"))
	g.Expect(base.SupportedResourcesReport.CurrentRelease).To(Equal("v2.1.0"))

	fmt.Printf("Merge successful! Final configuration:\n")
	fmt.Printf("  SchemaRoot: %s\n", base.SchemaRoot)
	fmt.Printf("  TypesOutputPath: %s\n", base.TypesOutputPath)
	fmt.Printf("  SamplesPath: %s\n", base.SamplesPath)
	fmt.Printf("  AnyTypePackages: %v\n", base.AnyTypePackages)
	fmt.Printf("  ReportsOutputFolder: %s\n", base.SupportedResourcesReport.OutputFolder)
}

// TestConfiguration_Merge_ConflictExample demonstrates conflict detection
func TestConfiguration_Merge_ConflictExample(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	// Create a base configuration
	base := NewConfiguration()
	base.SchemaRoot = "base/schemas/"

	// Create another configuration with conflicting value
	other := NewConfiguration()
	other.SchemaRoot = "other/schemas/"

	// Attempt to merge - this should fail with a conflict
	err := base.Merge(other)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("conflict"))
	g.Expect(err.Error()).To(ContainSubstring("SchemaRoot"))
	g.Expect(err.Error()).To(ContainSubstring("base/schemas/"))
	g.Expect(err.Error()).To(ContainSubstring("other/schemas/"))

	fmt.Printf("Conflict detected as expected: %s\n", err.Error())
}