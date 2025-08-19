/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestConfiguration_Merge_PrimitiveFields(t *testing.T) {
	t.Parallel()
	
	tests := map[string]struct {
		setupBase   func(*Configuration)
		setupOther  func(*Configuration)
		expectError bool
		validate    func(*WithT, *Configuration)
	}{
		"string field merges when base is empty": {
			setupBase:   func(c *Configuration) { c.SchemaRoot = "" },
			setupOther:  func(c *Configuration) { c.SchemaRoot = "other/path" },
			expectError: false,
			validate: func(g *WithT, c *Configuration) {
				g.Expect(c.SchemaRoot).To(Equal("other/path"))
			},
		},
		"string field errors when both have values": {
			setupBase:   func(c *Configuration) { c.SchemaRoot = "base/path" },
			setupOther:  func(c *Configuration) { c.SchemaRoot = "other/path" },
			expectError: true,
			validate:    nil,
		},
		"string field keeps base value when other is empty": {
			setupBase:   func(c *Configuration) { c.SchemaRoot = "base/path" },
			setupOther:  func(c *Configuration) { c.SchemaRoot = "" },
			expectError: false,
			validate: func(g *WithT, c *Configuration) {
				g.Expect(c.SchemaRoot).To(Equal("base/path"))
			},
		},
		"bool field merges when base is false": {
			setupBase:   func(c *Configuration) { c.EmitDocFiles = false },
			setupOther:  func(c *Configuration) { c.EmitDocFiles = true },
			expectError: false,
			validate: func(g *WithT, c *Configuration) {
				g.Expect(c.EmitDocFiles).To(BeTrue())
			},
		},
		"bool field errors when both are true": {
			setupBase:   func(c *Configuration) { c.EmitDocFiles = true },
			setupOther:  func(c *Configuration) { c.EmitDocFiles = true },
			expectError: true,
			validate:    nil,
		},
		"enum field merges when base is empty": {
			setupBase:   func(c *Configuration) { c.Pipeline = "" },
			setupOther:  func(c *Configuration) { c.Pipeline = GenerationPipelineCrossplane },
			expectError: false,
			validate: func(g *WithT, c *Configuration) {
				g.Expect(c.Pipeline).To(Equal(GenerationPipelineCrossplane))
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			base := NewConfiguration()
			other := NewConfiguration()
			
			if test.setupBase != nil {
				test.setupBase(base)
			}
			if test.setupOther != nil {
				test.setupOther(other)
			}

			err := base.Merge(other)

			if test.expectError {
				g.Expect(err).To(HaveOccurred())
			} else {
				g.Expect(err).To(Succeed())
				if test.validate != nil {
					test.validate(g, base)
				}
			}
		})
	}
}

func TestConfiguration_Merge_SliceFields(t *testing.T) {
	t.Parallel()
	
	tests := map[string]struct {
		setupBase   func(*Configuration)
		setupOther  func(*Configuration)
		expectError bool
		validate    func(*WithT, *Configuration)
	}{
		"slice appends to existing values": {
			setupBase: func(c *Configuration) {
				c.AnyTypePackages = []string{"base1", "base2"}
			},
			setupOther: func(c *Configuration) {
				c.AnyTypePackages = []string{"other1", "other2"}
			},
			expectError: false,
			validate: func(g *WithT, c *Configuration) {
				g.Expect(c.AnyTypePackages).To(Equal([]string{"base1", "base2", "other1", "other2"}))
			},
		},
		"slice adds to empty base": {
			setupBase: func(c *Configuration) {
				c.AnyTypePackages = nil
			},
			setupOther: func(c *Configuration) {
				c.AnyTypePackages = []string{"other1", "other2"}
			},
			expectError: false,
			validate: func(g *WithT, c *Configuration) {
				g.Expect(c.AnyTypePackages).To(Equal([]string{"other1", "other2"}))
			},
		},
		"slice preserves base when other is empty": {
			setupBase: func(c *Configuration) {
				c.AnyTypePackages = []string{"base1", "base2"}
			},
			setupOther: func(c *Configuration) {
				c.AnyTypePackages = nil
			},
			expectError: false,
			validate: func(g *WithT, c *Configuration) {
				g.Expect(c.AnyTypePackages).To(Equal([]string{"base1", "base2"}))
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			base := NewConfiguration()
			other := NewConfiguration()
			
			if test.setupBase != nil {
				test.setupBase(base)
			}
			if test.setupOther != nil {
				test.setupOther(other)
			}

			err := base.Merge(other)

			if test.expectError {
				g.Expect(err).To(HaveOccurred())
			} else {
				g.Expect(err).To(Succeed())
				if test.validate != nil {
					test.validate(g, base)
				}
			}
		})
	}
}

func TestConfiguration_Merge_ComplexFields(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	base := NewConfiguration()
	other := NewConfiguration()

	// Setup complex fields that should be merged
	other.ObjectModelConfiguration = NewObjectModelConfiguration()
	
	err := base.Merge(other)
	g.Expect(err).To(Succeed())
	
	// The complex types should be merged recursively
	g.Expect(base.ObjectModelConfiguration).ToNot(BeNil())
}

func TestConfiguration_Merge_NestedConfiguration(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	base := NewConfiguration()
	other := NewConfiguration()

	// Setup a configuration with nested values
	other.SupportedResourcesReport = NewSupportedResourcesReport(other)
	other.SupportedResourcesReport.OutputFolder = "output"
	other.SupportedResourcesReport.CurrentRelease = "v2.0.0"
	
	err := base.Merge(other)
	g.Expect(err).To(Succeed())
	
	g.Expect(base.SupportedResourcesReport).ToNot(BeNil())
	g.Expect(base.SupportedResourcesReport.OutputFolder).To(Equal("output"))
	g.Expect(base.SupportedResourcesReport.CurrentRelease).To(Equal("v2.0.0"))
}

func TestConfiguration_Merge_NestedConfigurationConflict(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	base := NewConfiguration()
	base.SupportedResourcesReport = NewSupportedResourcesReport(base)
	base.SupportedResourcesReport.OutputFolder = "base_output"
	
	other := NewConfiguration()
	other.SupportedResourcesReport = NewSupportedResourcesReport(other)
	other.SupportedResourcesReport.OutputFolder = "other_output"
	
	err := base.Merge(other)
	g.Expect(err).To(HaveOccurred())
	g.Expect(err.Error()).To(ContainSubstring("conflict"))
}

func TestConfiguration_Merge_NilOther(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	base := NewConfiguration()
	base.SchemaRoot = "test/path"
	
	err := base.Merge(nil)
	g.Expect(err).To(Succeed())
	g.Expect(base.SchemaRoot).To(Equal("test/path"))
}