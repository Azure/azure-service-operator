/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"testing"

	. "github.com/onsi/gomega"
)

// Azure names the connection after the link and a GUID of its own, and that name is the only thing tying
// the two together, so a link must not mistake a connection belonging to another for its own.
func Test_ConnectionOpenedBy_GivenConnectionName_ReportsWhetherItBelongsToTheLink(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		connectionName string
		linkName       string
		expected       bool
	}{
		"The link's own connection": {
			connectionName: "spl-6f8f3c1e-1f3a-4a2b-9c1d-2e5f7a9b0c3d",
			linkName:       "spl",
			expected:       true,
		},
		"A connection belonging to a link this one merely prefixes": {
			connectionName: "spl-extra-6f8f3c1e-1f3a-4a2b-9c1d-2e5f7a9b0c3d",
			linkName:       "spl",
			expected:       false,
		},
		"A link named after a longer one takes only its own": {
			connectionName: "spl-extra-6f8f3c1e-1f3a-4a2b-9c1d-2e5f7a9b0c3d",
			linkName:       "spl-extra",
			expected:       true,
		},
		"A connection opened by something other than a shared private link": {
			connectionName: "some-private-endpoint",
			linkName:       "spl",
			expected:       false,
		},
		"The link's name with nothing after it": {
			connectionName: "spl",
			linkName:       "spl",
			expected:       false,
		},
		"The link's name with an empty GUID after it": {
			connectionName: "spl-",
			linkName:       "spl",
			expected:       false,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			g.Expect(connectionOpenedBy(c.connectionName, c.linkName)).To(Equal(c.expected))
		})
	}
}
