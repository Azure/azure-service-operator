/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package config

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
)

func TestNewNotConfiguredError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	msg := "That's not configured"
	err := NewNotConfiguredError(msg)
	g.Expect(err.Error()).To(Equal(msg))
}

func TestNotConfiguredError_WithOptions(t *testing.T) {
	t.Parallel()
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

func TestIsNotConfiguredError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	nce := NewNotConfiguredError("Boom!")
	err := errors.New("Bang!")

	g.Expect(IsNotConfiguredError(nce)).To(BeTrue())
	g.Expect(IsNotConfiguredError(err)).To(BeFalse())
}
