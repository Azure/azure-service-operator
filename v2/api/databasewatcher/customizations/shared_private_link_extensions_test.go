/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package customizations

import (
	"testing"

	. "github.com/onsi/gomega"
)

const managedEndpointPath = "/subscriptions/s/resourceGroups/managed/providers/Microsoft.Network/privateEndpoints/"

// The managed private endpoint behind a connection carries the link's name, and that name is the only thing
// tying the two together, so a link must not mistake a connection belonging to another for its own.
func Test_ConnectionOpenedBy_GivenConnection_ReportsWhetherItBelongsToTheLink(t *testing.T) {
	t.Parallel()

	cases := map[string]struct {
		endpointID string
		linkName   string
		expected   bool
	}{
		"The link's own connection": {
			endpointID: managedEndpointPath + "spl",
			linkName:   "spl",
			expected:   true,
		},
		"A connection belonging to a link this one merely prefixes": {
			endpointID: managedEndpointPath + "spl-extra",
			linkName:   "spl",
			expected:   false,
		},
		"A link named after a longer one takes only its own": {
			endpointID: managedEndpointPath + "spl-extra",
			linkName:   "spl-extra",
			expected:   true,
		},
		"A connection opened by something other than a shared private link": {
			endpointID: "/subscriptions/s/resourceGroups/rg/providers/Microsoft.Network/privateEndpoints/pe",
			linkName:   "spl",
			expected:   false,
		},
		"A connection reporting no private endpoint at all": {
			endpointID: "",
			linkName:   "spl",
			expected:   false,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			var connection privateEndpointConnection
			connection.Properties.PrivateEndpoint.ID = c.endpointID

			g.Expect(connectionOpenedBy(&connection, c.linkName)).To(Equal(c.expected))
		})
	}
}
