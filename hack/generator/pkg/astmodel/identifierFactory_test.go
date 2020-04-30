/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_CreateIdentifier_GivenName_ReturnsExpectedIdentifier(t *testing.T) {
	cases := []struct {
		name     string
		expected string
	}{
		{"name", "Name"},
		{"Name", "Name"},
		{"$schema", "Schema"},
		{"my_important_name", "MyImportantName"},
		{"MediaServices_liveEvents_liveOutputs_childResource", "MediaServicesLiveEventsLiveOutputsChildResource"},
	}

	idfactory := NewIdentifierFactory()

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)
			identifier := idfactory.CreateIdentifier(c.name)
			g.Expect(identifier).To(Equal(c.expected))
		})
	}
}
