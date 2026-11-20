/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package postgresql

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestVerifyAADPrincipal(t *testing.T) {
	t.Parallel()

	const (
		objectID      = "00000000-0000-0000-0000-000000000001"
		otherObjectID = "00000000-0000-0000-0000-000000000002"
	)

	cases := []struct {
		name          string
		actual        *AADPrincipal
		objectID      string
		principalType string
		expectedText  string
	}{{
		name:         "Local role is rejected",
		actual:       nil,
		expectedText: "is not bound to a Microsoft Entra principal",
	}, {
		name:   "Entra role is accepted when bound by display name",
		actual: &AADPrincipal{ObjectID: objectID, PrincipalType: "service"},
	}, {
		name:          "Entra role bound to the expected object is accepted",
		actual:        &AADPrincipal{ObjectID: objectID, PrincipalType: "service"},
		objectID:      objectID,
		principalType: "service",
	}, {
		name:          "Object ID and principal type compare ignoring case",
		actual:        &AADPrincipal{ObjectID: "ABCDEF00-0000-0000-0000-000000000001", PrincipalType: "Service"},
		objectID:      "abcdef00-0000-0000-0000-000000000001",
		principalType: "service",
	}, {
		name:          "Entra role bound to another object is rejected",
		actual:        &AADPrincipal{ObjectID: otherObjectID, PrincipalType: "service"},
		objectID:      objectID,
		principalType: "service",
		expectedText:  "is bound to Microsoft Entra object " + otherObjectID,
	}, {
		name:          "Entra role bound with another principal type is rejected",
		actual:        &AADPrincipal{ObjectID: objectID, PrincipalType: "user"},
		objectID:      objectID,
		principalType: "service",
		expectedText:  "of type user, not service",
	}}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			err := VerifyAADPrincipal("app-identity", c.actual, c.objectID, c.principalType)
			if c.expectedText == "" {
				g.Expect(err).ToNot(HaveOccurred())
				return
			}
			g.Expect(err).To(MatchError(ContainSubstring(c.expectedText)))
		})
	}
}
