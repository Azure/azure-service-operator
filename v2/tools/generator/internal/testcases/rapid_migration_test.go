/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package testcases

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestRapidMigration_SelectedSmallGroupsUseRapid(t *testing.T) {
	t.Parallel()

	groups := []string{
		"monitor",
		"network.frontdoor",
		"quota",
		"redhatopenshift",
		"resources",
	}

	for _, group := range groups {
		group := group
		t.Run(group, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			g.Expect(UseRapidForGroup(group)).To(BeTrue())
			g.Expect(UseGopterForGroup(group)).To(BeFalse())
		})
	}
}
