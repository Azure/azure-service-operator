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

func TestNormalizeDatabasePrivileges(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name       string
		privileges []string
		expected   set.Set[string]
		expectErr  bool
	}{{
		name:       "Empty",
		privileges: nil,
		expected:   set.Make[string](),
	}, {
		name:       "Single privilege is upper-cased",
		privileges: []string{"connect"},
		expected:   set.Make("CONNECT"),
	}, {
		name:       "TEMP is TEMPORARY",
		privileges: []string{"TEMP", "temporary"},
		expected:   set.Make("TEMPORARY"),
	}, {
		name:       "ALL expands to every database privilege",
		privileges: []string{"ALL"},
		expected:   set.Make("CONNECT", "CREATE", "TEMPORARY"),
	}, {
		name:       "ALL PRIVILEGES expands to every database privilege",
		privileges: []string{"all privileges", "CONNECT"},
		expected:   set.Make("CONNECT", "CREATE", "TEMPORARY"),
	}, {
		name:       "Table privilege is rejected",
		privileges: []string{"CONNECT", "SELECT"},
		expectErr:  true,
	}, {
		name:       "Injection attempt is rejected",
		privileges: []string{"CONNECT ON DATABASE postgres TO public; --"},
		expectErr:  true,
	}}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			result, err := NormalizeDatabasePrivileges(c.privileges)
			if c.expectErr {
				g.Expect(err).To(HaveOccurred())
				return
			}
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(result).To(Equal(c.expected))
		})
	}
}

func TestDiffCurrentAndExpectedDatabasePrivileges(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name                     string
		currentPrivileges        set.Set[string]
		expectedPrivileges       set.Set[string]
		expectedPrivilegeAdds    set.Set[string]
		expectedPrivilegeDeletes set.Set[string]
	}{{
		name:                     "Current and expected equal",
		currentPrivileges:        set.Make("CONNECT"),
		expectedPrivileges:       set.Make("CONNECT"),
		expectedPrivilegeAdds:    set.Make[string](),
		expectedPrivilegeDeletes: set.Make[string](),
	}, {
		name:                     "Expected has privileges current lacks",
		currentPrivileges:        set.Make("CONNECT"),
		expectedPrivileges:       set.Make("CONNECT", "CREATE", "TEMPORARY"),
		expectedPrivilegeAdds:    set.Make("CREATE", "TEMPORARY"),
		expectedPrivilegeDeletes: set.Make[string](),
	}, {
		name:                     "Current has privileges expected lacks",
		currentPrivileges:        set.Make("CONNECT", "CREATE"),
		expectedPrivileges:       set.Make("CONNECT"),
		expectedPrivilegeAdds:    set.Make[string](),
		expectedPrivilegeDeletes: set.Make("CREATE"),
	}, {
		name:                     "Nothing expected revokes everything",
		currentPrivileges:        set.Make("CONNECT", "TEMPORARY"),
		expectedPrivileges:       set.Make[string](),
		expectedPrivilegeAdds:    set.Make[string](),
		expectedPrivilegeDeletes: set.Make("CONNECT", "TEMPORARY"),
	}, {
		name:                     "Adds and deletes together",
		currentPrivileges:        set.Make("TEMPORARY"),
		expectedPrivileges:       set.Make("CONNECT"),
		expectedPrivilegeAdds:    set.Make("CONNECT"),
		expectedPrivilegeDeletes: set.Make("TEMPORARY"),
	}}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			result := DiffCurrentAndExpectedDatabasePrivileges(c.currentPrivileges, c.expectedPrivileges)
			g.Expect(result.AddedPrivileges).To(Equal(c.expectedPrivilegeAdds))
			g.Expect(result.DeletedPrivileges).To(Equal(c.expectedPrivilegeDeletes))
		})
	}
}
