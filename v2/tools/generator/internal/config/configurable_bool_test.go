/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestConfigurableBool_WhenNoConfiguredValueToRead_ReturnsNoValue(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	var s configurableBool

	actual, ok := s.read()
	g.Expect(actual).To(BeFalse())
	g.Expect(ok).To(BeFalse())
}

func TestConfigurableBool_WhenNoConfiguredValueToRead_DoesNotReportUnconsumed(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	var s configurableBool
	g.Expect(s.isUnconsumed()).To(BeFalse())
}

func TestConfigurableBool_WhenConfiguredValue_ReportsUnconsumed(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	const expected = true
	var s configurableBool
	s.write(expected)
	g.Expect(s.isUnconsumed()).To(BeTrue())
}

func TestConfigurableBool_WhenConfiguredValue_ReadMarksAsConsumed(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	const expected = true
	var s configurableBool
	s.write(expected)

	actual, ok := s.read()
	g.Expect(actual).To(Equal(expected))
	g.Expect(ok).To(BeTrue())
	g.Expect(s.isUnconsumed()).To(BeFalse())
}

func TestConfigurableBool_WhenValueWritten_MarkedAsUnConsumed(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	const v = true
	var s configurableBool
	s.write(v)

	g.Expect(s.isUnconsumed()).To(BeTrue())
}
