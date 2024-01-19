/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package azuresql

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
	}{
		{
			name:                "Current and expected equal",
			currentRoles:        set.Set[string]{"db_datareader": {}},
			expectedRoles:       set.Set[string]{"db_datareader": {}},
			expectedRoleAdds:    set.Make[string](),
			expectedRoleDeletes: set.Make[string](),
		},
		{
			name:                "Expected has single role more than current",
			currentRoles:        set.Set[string]{"db_datareader": {}},
			expectedRoles:       set.Set[string]{"db_datareader": {}, "db_accessadmin": {}},
			expectedRoleAdds:    set.Set[string]{"db_accessadmin": {}},
			expectedRoleDeletes: set.Make[string](),
		},
		{
			name:                "Expected has single role less than current",
			currentRoles:        set.Set[string]{"db_datareader": {}, "db_accessadmin": {}},
			expectedRoles:       set.Set[string]{"db_datareader": {}},
			expectedRoleAdds:    set.Make[string](),
			expectedRoleDeletes: set.Set[string]{"db_accessadmin": {}},
		},
		{
			name:                "Expected has many roles less than current",
			currentRoles:        set.Set[string]{"db_datareader": {}, "db_datawriter": {}, "db_accessadmin": {}, "db_securityadmin": {}, "db_ddladmin": {}},
			expectedRoles:       set.Set[string]{"db_datareader": {}, "db_securityadmin": {}},
			expectedRoleAdds:    set.Make[string](),
			expectedRoleDeletes: set.Set[string]{"db_datawriter": {}, "db_accessadmin": {}, "db_ddladmin": {}},
		},
		{

			name:                "Expected has many roles more than current",
			currentRoles:        set.Set[string]{"db_datareader": {}, "db_securityadmin": {}},
			expectedRoles:       set.Set[string]{"db_datareader": {}, "db_datawriter": {}, "db_accessadmin": {}, "db_securityadmin": {}, "db_ddladmin": {}},
			expectedRoleAdds:    set.Set[string]{"db_datawriter": {}, "db_accessadmin": {}, "db_ddladmin": {}},
			expectedRoleDeletes: set.Make[string](),
		},
	}

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
