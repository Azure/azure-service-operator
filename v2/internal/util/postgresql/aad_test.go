/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package postgresql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"io"
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

func TestGetAADPrincipal(t *testing.T) {
	t.Parallel()

	// The documented result of pgaadauth_list_principals:
	// TABLE(rolename name, principalType text, objectId text, tenantId text, isMfa integer, isAdmin integer).
	// See https://learn.microsoft.com/azure/postgresql/security/security-manage-entra-users#list-microsoft-entra-roles-using-sql.
	columns := []string{"rolename", "principalType", "objectId", "tenantId", "isMfa", "isAdmin"}
	rows := [][]driver.Value{
		{"aad-admin", "service", "00000000-0000-0000-0000-00000000000a", "00000000-0000-0000-0000-0000000000ff", int64(0), int64(1)},
		{"app-identity", "service", "00000000-0000-0000-0000-000000000001", "00000000-0000-0000-0000-0000000000ff", int64(0), int64(0)},
		{"dba-group", "group", "00000000-0000-0000-0000-000000000002", "00000000-0000-0000-0000-0000000000ff", int64(1), int64(0)},
	}

	cases := []struct {
		name     string
		roleName string
		expected *AADPrincipal
	}{{
		name:     "Entra service principal is found",
		roleName: "app-identity",
		expected: &AADPrincipal{ObjectID: "00000000-0000-0000-0000-000000000001", PrincipalType: "service"},
	}, {
		name:     "Entra group is found",
		roleName: "dba-group",
		expected: &AADPrincipal{ObjectID: "00000000-0000-0000-0000-000000000002", PrincipalType: "group"},
	}, {
		name:     "Role that isn't an Entra principal is not found",
		roleName: "local-user",
		expected: nil,
	}}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			g := NewGomegaWithT(t)

			db := sql.OpenDB(&fakeConnector{columns: columns, rows: rows})
			defer db.Close()

			principal, err := GetAADPrincipal(context.Background(), db, c.roleName)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(principal).To(Equal(c.expected))
		})
	}
}

// fakeConnector is a database/sql connector whose every query returns the same rows.
type fakeConnector struct {
	columns []string
	rows    [][]driver.Value
}

var _ driver.Connector = &fakeConnector{}

func (c *fakeConnector) Connect(context.Context) (driver.Conn, error) {
	return &fakeConn{connector: c}, nil
}

func (c *fakeConnector) Driver() driver.Driver {
	panic("not implemented")
}

type fakeConn struct {
	connector *fakeConnector
}

var _ driver.QueryerContext = &fakeConn{}

func (c *fakeConn) QueryContext(context.Context, string, []driver.NamedValue) (driver.Rows, error) {
	return &fakeRows{columns: c.connector.columns, rows: c.connector.rows}, nil
}

func (c *fakeConn) Prepare(string) (driver.Stmt, error) {
	return nil, errors.New("not implemented")
}

func (c *fakeConn) Close() error {
	return nil
}

func (c *fakeConn) Begin() (driver.Tx, error) {
	return nil, errors.New("not implemented")
}

type fakeRows struct {
	columns []string
	rows    [][]driver.Value
	next    int
}

func (r *fakeRows) Columns() []string {
	return r.columns
}

func (r *fakeRows) Close() error {
	return nil
}

func (r *fakeRows) Next(dest []driver.Value) error {
	if r.next >= len(r.rows) {
		return io.EOF
	}

	copy(dest, r.rows[r.next])
	r.next++
	return nil
}
