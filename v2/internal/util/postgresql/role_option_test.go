/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package postgresql

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/set"
)

func TestDiffCurrentAndExpectedSQLRoleOptions(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name                       string
		currentRoleOptions         RoleOptionsSpec
		expectedRoleOptions        RoleOptionsSpec
		expectedChangedRoleOptions set.Set[RoleOption]
	}{{
		name:                       "Current and expected equal",
		currentRoleOptions:         RoleOptionsSpec{Login: true, CreateRole: false, CreateDb: false, Replication: false},
		expectedRoleOptions:        RoleOptionsSpec{Login: true},
		expectedChangedRoleOptions: set.Make[RoleOption](),
	}, {
		name:                       "Expected has single option more than current",
		currentRoleOptions:         RoleOptionsSpec{Login: true, CreateRole: false, CreateDb: false, Replication: false},
		expectedRoleOptions:        RoleOptionsSpec{Login: true, CreateDb: true},
		expectedChangedRoleOptions: set.Set[RoleOption]{CreateDb: {}},
	}, {
		name:                       "Expected all new values are set",
		currentRoleOptions:         RoleOptionsSpec{Login: true, CreateRole: false, CreateDb: false, Replication: false},
		expectedRoleOptions:        RoleOptionsSpec{Login: false, CreateRole: true, CreateDb: true, Replication: true},
		expectedChangedRoleOptions: set.Set[RoleOption]{NoLogin: {}, CreateRole: {}, CreateDb: {}, Replication: {}},
	},
		{
			name:                       "Expected all new values are set (non defaults)",
			currentRoleOptions:         RoleOptionsSpec{Login: false, CreateRole: true, CreateDb: true, Replication: true},
			expectedRoleOptions:        RoleOptionsSpec{Login: true, CreateRole: false, CreateDb: false, Replication: false},
			expectedChangedRoleOptions: set.Set[RoleOption]{Login: {}, NoCreateRole: {}, NoCreateDb: {}, NoReplication: {}},
		},
		{
			name:                       "Nothing changed if current is equal",
			currentRoleOptions:         RoleOptionsSpec{Login: false, CreateRole: true, CreateDb: true, Replication: true},
			expectedRoleOptions:        RoleOptionsSpec{Login: false, CreateRole: true, CreateDb: true, Replication: true},
			expectedChangedRoleOptions: set.Set[RoleOption]{},
		},
		{
			name:                       "Test Login changed",
			currentRoleOptions:         RoleOptionsSpec{},
			expectedRoleOptions:        RoleOptionsSpec{Login: true},
			expectedChangedRoleOptions: set.Set[RoleOption]{Login: {}},
		}, {
			name:                       "Test CreateRole changed",
			currentRoleOptions:         RoleOptionsSpec{},
			expectedRoleOptions:        RoleOptionsSpec{CreateRole: true},
			expectedChangedRoleOptions: set.Set[RoleOption]{CreateRole: {}},
		},
		{
			name:                       "Test CreateDb changed",
			currentRoleOptions:         RoleOptionsSpec{},
			expectedRoleOptions:        RoleOptionsSpec{CreateDb: true},
			expectedChangedRoleOptions: set.Set[RoleOption]{CreateDb: {}},
		},
		{
			name:                       "Test Replication changed",
			currentRoleOptions:         RoleOptionsSpec{Login: true, CreateRole: false, CreateDb: false, Replication: false},
			expectedRoleOptions:        RoleOptionsSpec{Login: true, Replication: true},
			expectedChangedRoleOptions: set.Set[RoleOption]{Replication: {}},
		},
	}

	// There's no test for handling ALL in current roles because we
	// don't see that in the database - it gets expanded into all
	// permissions.

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			result := DiffCurrentAndExpectedSQLRoleOptions(c.currentRoleOptions, c.expectedRoleOptions)
			g.Expect(result.ChangedRoleOptions).To(Equal(c.expectedChangedRoleOptions))
		})
	}
}
