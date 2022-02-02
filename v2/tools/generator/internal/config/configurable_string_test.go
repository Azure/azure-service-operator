/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestConfigurableString_WhenNoConfiguredValueToRead_ReturnsNoValue(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	var s configurableString

	actual, ok := s.read()
	g.Expect(actual).To(Equal(""))
	g.Expect(ok).To(BeFalse())

}

func TestConfigurableString_WhenNoConfiguredValueToRead_DoesNotReportUnconsumed(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	var s configurableString

	g.Expect(s.isUnconsumed()).To(BeFalse())
}

func TestConfigurableString_WhenConfiguredValue_ReportsUnconsumed(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	const expected = "Demo"
	var s configurableString
	s.write(expected)

	g.Expect(s.isUnconsumed()).To(BeTrue())
}

func TestConfigurableString_WhenConfiguredValue_ReadMarksAsConsumed(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	const expected = "Demo"
	var s configurableString
	s.write(expected)

	actual, ok := s.read()
	g.Expect(actual).To(Equal(expected))
	g.Expect(ok).To(BeTrue())
	g.Expect(s.isUnconsumed()).To(BeFalse())
}

func TestConfigurableString_WhenValueWritten_MarkedAsUnConsumed(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	const v = "Demo"
	var s configurableString
	s.write(v)

	g.Expect(s.isUnconsumed()).To(BeTrue())
}
