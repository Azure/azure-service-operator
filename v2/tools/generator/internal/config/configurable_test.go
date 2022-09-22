/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestConfigurable_WhenNoConfiguredValueToRead_ReturnsNoValue(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	var s configurable[bool]

	actual, ok := s.read()
	g.Expect(actual).To(BeFalse())
	g.Expect(ok).To(BeFalse())
}

func TestConfigurable_WhenNoConfiguredValueToRead_DoesNotReportUnconsumed(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	var s configurable[bool]
	g.Expect(s.isUnconsumed()).To(BeFalse())
}

func TestConfigurable_WhenConfiguredValue_ReportsUnconsumed(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	const expected = true
	var s configurable[bool]
	s.write(expected)
	g.Expect(s.isUnconsumed()).To(BeTrue())
}

func TestConfigurable_WhenConfiguredValue_ReadMarksAsConsumed(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	const expected = true
	var s configurable[bool]
	s.write(expected)

	actual, ok := s.read()
	g.Expect(actual).To(Equal(expected))
	g.Expect(ok).To(BeTrue())
	g.Expect(s.isUnconsumed()).To(BeFalse())
}

func TestConfigurable_WhenValueWritten_MarkedAsUnconsumed(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	const v = true
	var s configurable[bool]
	s.write(v)

	g.Expect(s.isUnconsumed()).To(BeTrue())
}
