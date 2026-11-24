/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package postgresql

import (
	"strings"
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
	}{
		{
			name:       "Empty",
			privileges: nil,
			expected:   set.Make[string](),
		},
		{
			name:       "Single privilege",
			privileges: []string{"CONNECT"},
			expected:   set.Make("CONNECT"),
		},
		{
			name:       "Lower case privilege is rejected",
			privileges: []string{"connect"},
			expectErr:  true,
		},
		{
			name:       "TEMP is TEMPORARY",
			privileges: []string{"TEMP", "TEMPORARY"},
			expected:   set.Make("TEMPORARY"),
		},
		{
			name:       "ALL expands to every database privilege",
			privileges: []string{"ALL"},
			expected:   set.Make("CONNECT", "CREATE", "TEMPORARY"),
		},
		{
			name:       "ALL PRIVILEGES expands to every database privilege",
			privileges: []string{"ALL PRIVILEGES", "CONNECT"},
			expected:   set.Make("CONNECT", "CREATE", "TEMPORARY"),
		},
		{
			name:       "Table privilege is rejected",
			privileges: []string{"CONNECT", "SELECT"},
			expectErr:  true,
		},
		{
			name:       "Injection attempt is rejected",
			privileges: []string{"CONNECT ON DATABASE postgres TO public; --"},
			expectErr:  true,
		},
	}

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
	}{
		{
			name:                     "Current and expected equal",
			currentPrivileges:        set.Make("CONNECT"),
			expectedPrivileges:       set.Make("CONNECT"),
			expectedPrivilegeAdds:    set.Make[string](),
			expectedPrivilegeDeletes: set.Make[string](),
		},
		{
			name:                     "Expected has privileges current lacks",
			currentPrivileges:        set.Make("CONNECT"),
			expectedPrivileges:       set.Make("CONNECT", "CREATE", "TEMPORARY"),
			expectedPrivilegeAdds:    set.Make("CREATE", "TEMPORARY"),
			expectedPrivilegeDeletes: set.Make[string](),
		},
		{
			name:                     "Current has privileges expected lacks",
			currentPrivileges:        set.Make("CONNECT", "CREATE"),
			expectedPrivileges:       set.Make("CONNECT"),
			expectedPrivilegeAdds:    set.Make[string](),
			expectedPrivilegeDeletes: set.Make("CREATE"),
		},
		{
			name:                     "Nothing expected revokes everything",
			currentPrivileges:        set.Make("CONNECT", "TEMPORARY"),
			expectedPrivileges:       set.Make[string](),
			expectedPrivilegeAdds:    set.Make[string](),
			expectedPrivilegeDeletes: set.Make("CONNECT", "TEMPORARY"),
		},
		{
			name:                     "Adds and deletes together",
			currentPrivileges:        set.Make("TEMPORARY"),
			expectedPrivileges:       set.Make("CONNECT"),
			expectedPrivilegeAdds:    set.Make("CONNECT"),
			expectedPrivilegeDeletes: set.Make("TEMPORARY"),
		},
	}

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

func TestValidateDatabasePrivileges(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name      string
		dbPrivs   map[string][]string
		expectErr bool
	}{
		{
			name:    "Nil",
			dbPrivs: nil,
		},
		{
			name:    "Empty",
			dbPrivs: map[string][]string{},
		},
		{
			name:    "Valid privileges",
			dbPrivs: map[string][]string{"app": {"CONNECT", "TEMP"}, "reports": {"ALL"}},
		},
		{
			name:    "Database name of 63 bytes is accepted",
			dbPrivs: map[string][]string{strings.Repeat("a", 63): {"CONNECT"}},
		},
		{
			name:      "Database name that extends a 63 byte name is rejected",
			dbPrivs:   map[string][]string{strings.Repeat("a", 64): {"CONNECT"}},
			expectErr: true,
		},
		{
			name:      "Multibyte database name is measured in bytes",
			dbPrivs:   map[string][]string{strings.Repeat("é", 32): {"CONNECT"}},
			expectErr: true,
		},
		{
			name:      "Table privilege on one database is rejected",
			dbPrivs:   map[string][]string{"app": {"CONNECT"}, "reports": {"SELECT"}},
			expectErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			err := ValidateDatabasePrivileges(c.dbPrivs)
			if c.expectErr {
				g.Expect(err).To(HaveOccurred())
				return
			}
			g.Expect(err).ToNot(HaveOccurred())
		})
	}
}

func TestPlanDatabasePrivileges(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		current  map[string]set.Set[string]
		desired  map[string]set.Set[string]
		owned    set.Set[string]
		expected map[string]SQLDatabasePrivilegeDelta
	}{
		{
			name:     "Matching privileges need no changes",
			current:  map[string]set.Set[string]{"app": set.Make("CONNECT")},
			desired:  map[string]set.Set[string]{"app": set.Make("CONNECT")},
			owned:    set.Make[string](),
			expected: map[string]SQLDatabasePrivilegeDelta{},
		},
		{
			name:    "Unlisted database is revoked and listed database is granted",
			current: map[string]set.Set[string]{"old": set.Make("CONNECT", "TEMPORARY")},
			desired: map[string]set.Set[string]{"app": set.Make("CONNECT")},
			owned:   set.Make[string](),
			expected: map[string]SQLDatabasePrivilegeDelta{
				"old": {AddedPrivileges: set.Make[string](), DeletedPrivileges: set.Make("CONNECT", "TEMPORARY")},
				"app": {AddedPrivileges: set.Make("CONNECT"), DeletedPrivileges: set.Make[string]()},
			},
		},
		{
			name:     "Owned database listed in the spec is never granted",
			current:  map[string]set.Set[string]{},
			desired:  map[string]set.Set[string]{"mine": set.Make("CONNECT", "CREATE", "TEMPORARY")},
			owned:    set.Make("mine"),
			expected: map[string]SQLDatabasePrivilegeDelta{},
		},
		{
			name:     "Owned database not listed in the spec is never revoked",
			current:  map[string]set.Set[string]{"mine": set.Make("CONNECT")},
			desired:  map[string]set.Set[string]{},
			owned:    set.Make("mine"),
			expected: map[string]SQLDatabasePrivilegeDelta{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			g.Expect(PlanDatabasePrivileges(c.current, c.desired, c.owned)).To(Equal(c.expected))
		})
	}
}

func TestDatabasePrivilegeMismatches(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name          string
		actual        map[string]set.Set[string]
		desired       map[string]set.Set[string]
		owned         set.Set[string]
		expectedTexts []string
	}{
		{
			name:    "Applied privileges match",
			actual:  map[string]set.Set[string]{"app": set.Make("CONNECT")},
			desired: map[string]set.Set[string]{"app": set.Make("CONNECT")},
			owned:   set.Make[string](),
		},
		{
			name:          "Grant that only warned is reported",
			actual:        map[string]set.Set[string]{"app": set.Make("CONNECT")},
			desired:       map[string]set.Set[string]{"app": set.Make("CONNECT", "CREATE")},
			owned:         set.Make[string](),
			expectedTexts: []string{"database app", "missing: [CREATE]"},
		},
		{
			name:          "Revoke that only warned is reported",
			actual:        map[string]set.Set[string]{"other": set.Make("TEMPORARY")},
			desired:       map[string]set.Set[string]{},
			owned:         set.Make[string](),
			expectedTexts: []string{"database other", "unexpected: [TEMPORARY]"},
		},
		{
			name:    "Owned database is not reported",
			actual:  map[string]set.Set[string]{},
			desired: map[string]set.Set[string]{"mine": set.Make("CONNECT")},
			owned:   set.Make("mine"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			err := DatabasePrivilegeMismatches(c.actual, c.desired, c.owned)
			if len(c.expectedTexts) == 0 {
				g.Expect(err).ToNot(HaveOccurred())
				return
			}
			g.Expect(err).To(HaveOccurred())
			for _, text := range c.expectedTexts {
				g.Expect(err.Error()).To(ContainSubstring(text))
			}
		})
	}
}
