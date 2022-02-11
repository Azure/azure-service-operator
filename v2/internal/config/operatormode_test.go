// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package config_test

import (
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/config"
)

func TestIncludesWebhooks(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	g.Expect(config.OperatorModeWebhooks.IncludesWebhooks()).To(BeTrue())
	g.Expect(config.OperatorModeBoth.IncludesWebhooks()).To(BeTrue())
	g.Expect(config.OperatorModeWatchers.IncludesWebhooks()).To(BeFalse())
}

func TestIncludesWatchers(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	g.Expect(config.OperatorModeWebhooks.IncludesWatchers()).To(BeFalse())
	g.Expect(config.OperatorModeBoth.IncludesWatchers()).To(BeTrue())
	g.Expect(config.OperatorModeWatchers.IncludesWatchers()).To(BeTrue())
}

func TestRoundtripString(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	values := []string{
		"watchers-and-webhooks", "webhooks", "watchers",
		// Check for case insensitivity too.
		"Watchers-And-Webhooks", "wEbHoOkS", "WATCHERS",
	}
	for _, value := range values {
		mode, err := config.ParseOperatorMode(value)
		g.Expect(err).ToNot(HaveOccurred(), value)
		g.Expect(mode.String()).To(Equal(strings.ToLower(value)))
	}
}

func TestRoundtripValue(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	modes := []config.OperatorMode{
		config.OperatorModeWatchers,
		config.OperatorModeWebhooks,
		config.OperatorModeBoth,
	}
	for _, mode := range modes {
		parsed, err := config.ParseOperatorMode(mode.String())
		g.Expect(err).ToNot(HaveOccurred(), mode.String())
		g.Expect(parsed).To(Equal(mode))
	}
}

func TestParseInvalid(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	_, err := config.ParseOperatorMode("dimmer")
	g.Expect(err.Error()).To(ContainSubstring(`operator mode value must be one of "watchers-and-webhooks", "webhooks" or "watchers" but was "dimmer"`))
}

func TestStringInvalid(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	g.Expect(func() { _ = config.OperatorMode(0).String() }).To(Panic())
}
