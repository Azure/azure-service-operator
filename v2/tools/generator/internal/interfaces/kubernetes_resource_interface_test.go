/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package interfaces

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

func Test_createAzureNameFunctionHandlersForPrimitiveType_givenType_returnsExpectedResult(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		Type                    *astmodel.PrimitiveType
		expectedErrorSubstrings []string
	}{
		"string is accepted": {
			Type: astmodel.StringType,
		},
		"int is not accepted": {
			Type: astmodel.IntType,
			expectedErrorSubstrings: []string{
				"cannot use type",
				"int",
				astmodel.AzureNameProperty,
			},
		},
	}

	for n, c := range cases {
		t.Run(n, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			_, err := createAzureNameFunctionHandlersForPrimitiveType(c.Type)
			if len(c.expectedErrorSubstrings) == 0 {
				g.Expect(err).To(BeNil())
			} else {
				g.Expect(err).To(HaveOccurred())
				for _, substr := range c.expectedErrorSubstrings {
					g.Expect(err).To(MatchError(ContainSubstring(substr)))
				}
			}
		})
	}
}
