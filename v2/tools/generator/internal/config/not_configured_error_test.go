package config

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestNewNotConfiguredError(t *testing.T) {
	g := NewGomegaWithT(t)

	msg := "That's not configured"
	err := NewNotConfiguredError(msg)
	g.Expect(err.Error()).To(Equal(msg))
}

func TestNotConfiguredError_WithOptions(t *testing.T) {
	g := NewGomegaWithT(t)
	options := []string{
		"Alpha",
		"Beta",
		"Gamma",
	}

	err := NewNotConfiguredError("options").WithOptions("letters", options)
	g.Expect(err.Error()).To(ContainSubstring("letters"))
	for _, o := range options {
		g.Expect(err.Error()).To(ContainSubstring(o))
	}
}
