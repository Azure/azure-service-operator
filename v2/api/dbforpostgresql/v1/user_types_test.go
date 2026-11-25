/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v1

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestUserSpec_DatabasePrivileges_Serialization(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name               string
		databasePrivileges map[string][]DatabasePrivilege
		expectField        bool
	}{
		{
			name:               "Omitted map is left out",
			databasePrivileges: nil,
			expectField:        false,
		},
		{
			name:               "Empty map is kept",
			databasePrivileges: map[string][]DatabasePrivilege{},
			expectField:        true,
		},
		{
			name:               "Populated map is kept",
			databasePrivileges: map[string][]DatabasePrivilege{"app": {"CONNECT"}},
			expectField:        true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			data, err := json.Marshal(UserSpec{DatabasePrivileges: c.databasePrivileges})
			g.Expect(err).ToNot(HaveOccurred())

			var fields map[string]json.RawMessage
			g.Expect(json.Unmarshal(data, &fields)).To(Succeed())
			_, found := fields["databasePrivileges"]
			g.Expect(found).To(Equal(c.expectField))

			// The round trip keeps an empty map distinct from an omitted one
			var roundTripped UserSpec
			g.Expect(json.Unmarshal(data, &roundTripped)).To(Succeed())
			g.Expect(roundTripped.DatabasePrivileges == nil).To(Equal(c.databasePrivileges == nil))
		})
	}
}
