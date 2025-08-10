/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestConfigurable_Merge_String(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		setupBase   func(*configurable[string])
		setupOther  func(*configurable[string])
		expectError bool
		validate    func(*WithT, *configurable[string])
	}{
		"merges when base is empty": {
			setupBase:   func(c *configurable[string]) {},
			setupOther:  func(c *configurable[string]) { c.Set("other_value") },
			expectError: false,
			validate: func(g *WithT, c *configurable[string]) {
				val, ok := c.Lookup()
				g.Expect(ok).To(BeTrue())
				g.Expect(val).To(Equal("other_value"))
			},
		},
		"merges when base is empty string": {
			setupBase:   func(c *configurable[string]) { c.Set("") },
			setupOther:  func(c *configurable[string]) { c.Set("other_value") },
			expectError: false,
			validate: func(g *WithT, c *configurable[string]) {
				val, ok := c.Lookup()
				g.Expect(ok).To(BeTrue())
				g.Expect(val).To(Equal("other_value"))
			},
		},
		"preserves base when other is empty": {
			setupBase:   func(c *configurable[string]) { c.Set("base_value") },
			setupOther:  func(c *configurable[string]) { c.Set("") },
			expectError: false,
			validate: func(g *WithT, c *configurable[string]) {
				val, ok := c.Lookup()
				g.Expect(ok).To(BeTrue())
				g.Expect(val).To(Equal("base_value"))
			},
		},
		"preserves base when other is nil": {
			setupBase:   func(c *configurable[string]) { c.Set("base_value") },
			setupOther:  func(c *configurable[string]) {},
			expectError: false,
			validate: func(g *WithT, c *configurable[string]) {
				val, ok := c.Lookup()
				g.Expect(ok).To(BeTrue())
				g.Expect(val).To(Equal("base_value"))
			},
		},
		"succeeds when both values are the same": {
			setupBase:   func(c *configurable[string]) { c.Set("same_value") },
			setupOther:  func(c *configurable[string]) { c.Set("same_value") },
			expectError: false,
			validate: func(g *WithT, c *configurable[string]) {
				val, ok := c.Lookup()
				g.Expect(ok).To(BeTrue())
				g.Expect(val).To(Equal("same_value"))
			},
		},
		"errors when attempting to overwrite": {
			setupBase:   func(c *configurable[string]) { c.Set("base_value") },
			setupOther:  func(c *configurable[string]) { c.Set("other_value") },
			expectError: true,
			validate:    nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			base := makeConfigurable[string]("test", "test")
			other := makeConfigurable[string]("test", "test")

			if test.setupBase != nil {
				test.setupBase(&base)
			}
			if test.setupOther != nil {
				test.setupOther(&other)
			}

			err := base.Merge(&other)

			if test.expectError {
				g.Expect(err).To(HaveOccurred())
			} else {
				g.Expect(err).To(Succeed())
				if test.validate != nil {
					test.validate(g, &base)
				}
			}
		})
	}
}

func TestConfigurable_Merge_Bool(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		setupBase   func(*configurable[bool])
		setupOther  func(*configurable[bool])
		expectError bool
		validate    func(*WithT, *configurable[bool])
	}{
		"merges when base is false": {
			setupBase:   func(c *configurable[bool]) { c.Set(false) },
			setupOther:  func(c *configurable[bool]) { c.Set(true) },
			expectError: false,
			validate: func(g *WithT, c *configurable[bool]) {
				val, ok := c.Lookup()
				g.Expect(ok).To(BeTrue())
				g.Expect(val).To(BeTrue())
			},
		},
		"preserves when other is false": {
			setupBase:   func(c *configurable[bool]) { c.Set(true) },
			setupOther:  func(c *configurable[bool]) { c.Set(false) },
			expectError: false,
			validate: func(g *WithT, c *configurable[bool]) {
				val, ok := c.Lookup()
				g.Expect(ok).To(BeTrue())
				g.Expect(val).To(BeTrue())
			},
		},
		"errors when both are true": {
			setupBase:   func(c *configurable[bool]) { c.Set(true) },
			setupOther:  func(c *configurable[bool]) { c.Set(true) },
			expectError: true,
			validate:    nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			base := makeConfigurable[bool]("test", "test")
			other := makeConfigurable[bool]("test", "test")

			if test.setupBase != nil {
				test.setupBase(&base)
			}
			if test.setupOther != nil {
				test.setupOther(&other)
			}

			err := base.Merge(&other)

			if test.expectError {
				g.Expect(err).To(HaveOccurred())
			} else {
				g.Expect(err).To(Succeed())
				if test.validate != nil {
					test.validate(g, &base)
				}
			}
		})
	}
}

func TestConfigurable_Merge_StringSlice(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		setupBase   func(*configurable[[]string])
		setupOther  func(*configurable[[]string])
		expectError bool
		validate    func(*WithT, *configurable[[]string])
	}{
		"appends to existing slice": {
			setupBase:   func(c *configurable[[]string]) { c.Set([]string{"base1", "base2"}) },
			setupOther:  func(c *configurable[[]string]) { c.Set([]string{"other1", "other2"}) },
			expectError: false,
			validate: func(g *WithT, c *configurable[[]string]) {
				val, ok := c.Lookup()
				g.Expect(ok).To(BeTrue())
				g.Expect(val).To(Equal([]string{"base1", "base2", "other1", "other2"}))
			},
		},
		"sets slice when base is empty": {
			setupBase:   func(c *configurable[[]string]) {},
			setupOther:  func(c *configurable[[]string]) { c.Set([]string{"other1", "other2"}) },
			expectError: false,
			validate: func(g *WithT, c *configurable[[]string]) {
				val, ok := c.Lookup()
				g.Expect(ok).To(BeTrue())
				g.Expect(val).To(Equal([]string{"other1", "other2"}))
			},
		},
		"preserves base when other is empty": {
			setupBase:   func(c *configurable[[]string]) { c.Set([]string{"base1", "base2"}) },
			setupOther:  func(c *configurable[[]string]) { c.Set([]string{}) },
			expectError: false,
			validate: func(g *WithT, c *configurable[[]string]) {
				val, ok := c.Lookup()
				g.Expect(ok).To(BeTrue())
				g.Expect(val).To(Equal([]string{"base1", "base2"}))
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			base := makeConfigurable[[]string]("test", "test")
			other := makeConfigurable[[]string]("test", "test")

			if test.setupBase != nil {
				test.setupBase(&base)
			}
			if test.setupOther != nil {
				test.setupOther(&other)
			}

			err := base.Merge(&other)

			if test.expectError {
				g.Expect(err).To(HaveOccurred())
			} else {
				g.Expect(err).To(Succeed())
				if test.validate != nil {
					test.validate(g, &base)
				}
			}
		})
	}
}

func TestConfigurable_Merge_StringMap(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		setupBase   func(*configurable[map[string]string])
		setupOther  func(*configurable[map[string]string])
		expectError bool
		validate    func(*WithT, *configurable[map[string]string])
	}{
		"merges new keys": {
			setupBase:   func(c *configurable[map[string]string]) { c.Set(map[string]string{"key1": "base1"}) },
			setupOther:  func(c *configurable[map[string]string]) { c.Set(map[string]string{"key2": "other1"}) },
			expectError: false,
			validate: func(g *WithT, c *configurable[map[string]string]) {
				val, ok := c.Lookup()
				g.Expect(ok).To(BeTrue())
				g.Expect(val).To(Equal(map[string]string{"key1": "base1", "key2": "other1"}))
			},
		},
		"sets map when base is empty": {
			setupBase:   func(c *configurable[map[string]string]) {},
			setupOther:  func(c *configurable[map[string]string]) { c.Set(map[string]string{"key1": "other1"}) },
			expectError: false,
			validate: func(g *WithT, c *configurable[map[string]string]) {
				val, ok := c.Lookup()
				g.Expect(ok).To(BeTrue())
				g.Expect(val).To(Equal(map[string]string{"key1": "other1"}))
			},
		},
		"errors on key conflict with different values": {
			setupBase:   func(c *configurable[map[string]string]) { c.Set(map[string]string{"key1": "base_value"}) },
			setupOther:  func(c *configurable[map[string]string]) { c.Set(map[string]string{"key1": "other_value"}) },
			expectError: true,
			validate:    nil,
		},
		"succeeds on key conflict with same values": {
			setupBase:   func(c *configurable[map[string]string]) { c.Set(map[string]string{"key1": "same_value"}) },
			setupOther:  func(c *configurable[map[string]string]) { c.Set(map[string]string{"key1": "same_value"}) },
			expectError: false,
			validate: func(g *WithT, c *configurable[map[string]string]) {
				val, ok := c.Lookup()
				g.Expect(ok).To(BeTrue())
				g.Expect(val).To(Equal(map[string]string{"key1": "same_value"}))
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			base := makeConfigurable[map[string]string]("test", "test")
			other := makeConfigurable[map[string]string]("test", "test")

			if test.setupBase != nil {
				test.setupBase(&base)
			}
			if test.setupOther != nil {
				test.setupOther(&other)
			}

			err := base.Merge(&other)

			if test.expectError {
				g.Expect(err).To(HaveOccurred())
			} else {
				g.Expect(err).To(Succeed())
				if test.validate != nil {
					test.validate(g, &base)
				}
			}
		})
	}
}

func TestConfigurable_Merge_Enum(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		setupBase   func(*configurable[PayloadType])
		setupOther  func(*configurable[PayloadType])
		expectError bool
		validate    func(*WithT, *configurable[PayloadType])
	}{
		"merges when base is empty": {
			setupBase:   func(c *configurable[PayloadType]) {},
			setupOther:  func(c *configurable[PayloadType]) { c.Set(ExplicitCollections) },
			expectError: false,
			validate: func(g *WithT, c *configurable[PayloadType]) {
				val, ok := c.Lookup()
				g.Expect(ok).To(BeTrue())
				g.Expect(val).To(Equal(ExplicitCollections))
			},
		},
		"succeeds when both values are the same": {
			setupBase:   func(c *configurable[PayloadType]) { c.Set(ExplicitCollections) },
			setupOther:  func(c *configurable[PayloadType]) { c.Set(ExplicitCollections) },
			expectError: false,
			validate: func(g *WithT, c *configurable[PayloadType]) {
				val, ok := c.Lookup()
				g.Expect(ok).To(BeTrue())
				g.Expect(val).To(Equal(ExplicitCollections))
			},
		},
		"errors when values differ": {
			setupBase:   func(c *configurable[PayloadType]) { c.Set(ExplicitCollections) },
			setupOther:  func(c *configurable[PayloadType]) { c.Set(ExplicitProperties) },
			expectError: true,
			validate:    nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			base := makeConfigurable[PayloadType]("test", "test")
			other := makeConfigurable[PayloadType]("test", "test")

			if test.setupBase != nil {
				test.setupBase(&base)
			}
			if test.setupOther != nil {
				test.setupOther(&other)
			}

			err := base.Merge(&other)

			if test.expectError {
				g.Expect(err).To(HaveOccurred())
			} else {
				g.Expect(err).To(Succeed())
				if test.validate != nil {
					test.validate(g, &base)
				}
			}
		})
	}
}