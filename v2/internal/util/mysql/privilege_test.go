/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package mysql

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/internal/set"
)

func TestDiffCurrentAndExpectedSQLRoles(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name                     string
		currentPrivileges        set.Set[string]
		expectedPrivileges       set.Set[string]
		expectedPrivilegeAdds    set.Set[string]
		expectedPrivilegeDeletes set.Set[string]
	}{{
		name:                     "Current and expected equal",
		currentPrivileges:        set.Set[string]{"USAGE": {}},
		expectedPrivileges:       set.Set[string]{"USAGE": {}},
		expectedPrivilegeAdds:    set.Make[string](),
		expectedPrivilegeDeletes: set.Make[string](),
	}, {
		name:                     "Expected has single role more than current",
		currentPrivileges:        set.Set[string]{"USAGE": {}},
		expectedPrivileges:       set.Set[string]{"USAGE": {}, "SELECT": {}},
		expectedPrivilegeAdds:    set.Set[string]{"SELECT": {}},
		expectedPrivilegeDeletes: set.Make[string](),
	}, {
		name:                     "Expected has single role less than current",
		currentPrivileges:        set.Set[string]{"USAGE": {}, "SELECT": {}},
		expectedPrivileges:       set.Set[string]{"USAGE": {}},
		expectedPrivilegeAdds:    set.Make[string](),
		expectedPrivilegeDeletes: set.Set[string]{"SELECT": {}},
	}, {
		name:                     "Expected has many roles less than current",
		currentPrivileges:        set.Set[string]{"SELECT": {}, "INSERT": {}, "UPDATE": {}, "DELETE": {}, "CREATE": {}, "DROP": {}, "RELOAD": {}},
		expectedPrivileges:       set.Set[string]{"SELECT": {}, "INSERT": {}},
		expectedPrivilegeAdds:    set.Make[string](),
		expectedPrivilegeDeletes: set.Set[string]{"UPDATE": {}, "DELETE": {}, "CREATE": {}, "DROP": {}, "RELOAD": {}},
	}, {
		name:                     "Expected has many roles more than current",
		currentPrivileges:        set.Set[string]{"SELECT": {}, "INSERT": {}},
		expectedPrivileges:       set.Set[string]{"SELECT": {}, "INSERT": {}, "UPDATE": {}, "DELETE": {}, "CREATE": {}, "DROP": {}, "RELOAD": {}},
		expectedPrivilegeAdds:    set.Set[string]{"UPDATE": {}, "DELETE": {}, "CREATE": {}, "DROP": {}, "RELOAD": {}},
		expectedPrivilegeDeletes: set.Make[string](),
	}, {
		name:                     "Including ALL sets added to ALL and doesn't delete any",
		currentPrivileges:        set.Set[string]{"SELECT": {}, "INSERT": {}},
		expectedPrivileges:       set.Set[string]{"ALL": {}, "UPDATE": {}},
		expectedPrivilegeAdds:    set.Set[string]{"ALL": {}},
		expectedPrivilegeDeletes: set.Make[string](),
	}, {
		name:                     "ALL is case insensitive",
		currentPrivileges:        set.Set[string]{"SELECT": {}, "INSERT": {}},
		expectedPrivileges:       set.Set[string]{"aLl": {}, "UPDATE": {}},
		expectedPrivilegeAdds:    set.Set[string]{"ALL": {}},
		expectedPrivilegeDeletes: set.Make[string](),
	}}

	// There's no test for handling ALL in current roles because we
	// don't see that in the database - it gets expanded into all
	// permissions.

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			result := DiffCurrentAndExpectedSQLPrivileges(c.currentPrivileges, c.expectedPrivileges)
			g.Expect(result.AddedPrivileges).To(Equal(c.expectedPrivilegeAdds))
			g.Expect(result.DeletedPrivileges).To(Equal(c.expectedPrivilegeDeletes))
		})
	}
}
