/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestConfigurableStringSlice_WhenNoConfiguredValueToRead_ReturnsNoValue(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	var s configurableStringSlice

	actual, ok := s.read()
	g.Expect(actual).To(BeNil())
	g.Expect(ok).To(BeFalse())

}

func TestConfigurableStringSlice_WhenNoConfiguredValueToRead_DoesNotReportUnconsumed(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	var s configurableStringSlice

	g.Expect(s.isUnconsumed()).To(BeFalse())
}

func TestConfigurableStringSlice_WhenConfiguredValue_ReportsUnconsumed(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	expected := []string{
		"test",
	}
	var s configurableStringSlice
	s.write(expected)

	g.Expect(s.isUnconsumed()).To(BeTrue())
}

func TestConfigurableStringSlice_WhenConfiguredValue_ReadMarksAsConsumed(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	expected := []string{
		"test",
	}
	var s configurableStringSlice
	s.write(expected)

	actual, ok := s.read()
	g.Expect(actual).To(Equal(expected))
	g.Expect(ok).To(BeTrue())
	g.Expect(s.isUnconsumed()).To(BeFalse())
}

func TestConfigurableStringSlice_WhenValueWritten_MarkedAsUnconsumed(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	v := []string{
		"test",
	}
	var s configurableStringSlice
	s.write(v)

	g.Expect(s.isUnconsumed()).To(BeTrue())
}
