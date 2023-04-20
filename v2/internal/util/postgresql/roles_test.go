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

func TestDiffCurrentAndExpectedSQLRoles(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name                string
		currentRoles        set.Set[string]
		expectedRoles       set.Set[string]
		expectedRoleAdds    set.Set[string]
		expectedRoleDeletes set.Set[string]
	}{{
		name:                "Current and expected equal",
		currentRoles:        set.Set[string]{"azure_pg_admin": {}},
		expectedRoles:       set.Set[string]{"azure_pg_admin": {}},
		expectedRoleAdds:    set.Make[string](),
		expectedRoleDeletes: set.Make[string](),
	}, {
		name:                "Expected has single role more than current",
		currentRoles:        set.Set[string]{"azure_pg_admin": {}},
		expectedRoles:       set.Set[string]{"azure_pg_admin": {}, "azure_superuser": {}},
		expectedRoleAdds:    set.Set[string]{"azure_superuser": {}},
		expectedRoleDeletes: set.Make[string](),
	}, {
		name:                "Expected has single role less than current",
		currentRoles:        set.Set[string]{"azure_pg_admin": {}, "azure_superuser": {}},
		expectedRoles:       set.Set[string]{"azure_pg_admin": {}},
		expectedRoleAdds:    set.Make[string](),
		expectedRoleDeletes: set.Set[string]{"azure_superuser": {}},
	}, {
		name:                "Expected has many roles less than current",
		currentRoles:        set.Set[string]{"azure_pg_admin": {}, "azure_superuser": {}, "pg_monitor": {}},
		expectedRoles:       set.Set[string]{"azure_pg_admin": {}},
		expectedRoleAdds:    set.Make[string](),
		expectedRoleDeletes: set.Set[string]{"azure_superuser": {}, "pg_monitor": {}},
	}, {
		name:                "Expected has many roles more than current",
		currentRoles:        set.Set[string]{"azure_pg_admin": {}},
		expectedRoles:       set.Set[string]{"azure_pg_admin": {}, "azure_superuser": {}, "pg_monitor": {}},
		expectedRoleAdds:    set.Set[string]{"azure_superuser": {}, "pg_monitor": {}},
		expectedRoleDeletes: set.Make[string](),
	}}

	// There's no test for handling ALL in current roles because we
	// don't see that in the database - it gets expanded into all
	// permissions.

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			result := DiffCurrentAndExpectedSQLRoles(c.currentRoles, c.expectedRoles)
			g.Expect(result.AddedRoles).To(Equal(c.expectedRoleAdds))
			g.Expect(result.DeletedRoles).To(Equal(c.expectedRoleDeletes))
		})
	}
}
