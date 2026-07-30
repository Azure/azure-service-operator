/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

func TestResourceParameterIdent_GivenReceiver_ReturnsExpectedIdent(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name          string
		receiverIdent string
		expected      string
	}{
		{"receiver leaves resource free", "person", "resource"},
		{"receiver of a kind ending in Resource takes it", "resource", "resourceObj"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			g.Expect(resourceParameterIdent(c.receiverIdent, astmodel.NewIdentifierFactory())).To(Equal(c.expected))
		})
	}
}
