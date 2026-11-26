/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package postgresql

import (
	"context"
	"database/sql"
	"strings"

	"github.com/rotisserie/eris"

	"github.com/Azure/azure-service-operator/v2/internal/set"
)

// Scope is the token scope Azure Database for PostgreSQL flexible server accepts for Microsoft Entra authentication.
const Scope = "https://ossrdbms-aad.database.windows.net/.default"

// PrincipalTypes are the principal types pgaadauth_create_principal_with_oid accepts.
var PrincipalTypes = set.Make("user", "group", "service")

// ConnectToDBAAD connects as a Microsoft Entra principal: the user is the principal's role name and the password is an
// access token for Scope.
func ConnectToDBAAD(ctx context.Context, fullServerName string, database string, port int, user string, token string) (*sql.DB, error) {
	return ConnectToDB(ctx, fullServerName, database, port, user, token)
}

// CreateAADUser creates a role bound to a Microsoft Entra principal with the pgaadauth extension of Azure Database for
// PostgreSQL flexible server. Without an objectID the role is bound to the principal whose display name is roleName
// (pgaadauth_create_principal); with one, to that object under any role name (pgaadauth_create_principal_with_oid).
// The role is created with LOGIN, not as an administrator and without MFA.
// See https://learn.microsoft.com/azure/postgresql/flexible-server/how-to-manage-azure-ad-users.
func CreateAADUser(ctx context.Context, db *sql.DB, roleName string, objectID string, principalType string) (*SQLUser, error) {
	if err := FindBadChars(roleName); err != nil {
		return nil, eris.Wrap(err, "problem found with username")
	}

	var err error
	if objectID == "" {
		_, err = db.ExecContext(ctx, "SELECT * FROM pgaadauth_create_principal($1, false, false)", roleName)
	} else {
		if !PrincipalTypes.Contains(principalType) {
			return nil, eris.Errorf("principal type %q must be one of user, group or service", principalType)
		}
		_, err = db.ExecContext(ctx, "SELECT * FROM pgaadauth_create_principal_with_oid($1, $2, $3, false, false)", roleName, objectID, principalType)
	}
	if err != nil {
		return nil, eris.Wrapf(err, "failed to create Microsoft Entra principal %s", roleName)
	}

	return &SQLUser{Name: roleName}, nil
}

// AADPrincipal is the Microsoft Entra principal a role is bound to.
type AADPrincipal struct {
	ObjectID      string
	PrincipalType string
}

// GetAADPrincipal returns the Microsoft Entra principal the role is bound to, or nil if the role is not bound to one
// (for example a local role created with a password).
// See https://learn.microsoft.com/azure/postgresql/flexible-server/how-to-manage-azure-ad-users.
func GetAADPrincipal(ctx context.Context, db *sql.DB, roleName string) (*AADPrincipal, error) {
	// pgaadauth_list_principals(false) lists every role bound to a Microsoft Entra principal, administrators included.
	// The live server returns "rolname", despite the documentation calling it "rolename".
	// Columns are matched by name, ignoring case, so the query doesn't depend on the extension's column order.
	// See https://learn.microsoft.com/azure/postgresql/security/security-manage-entra-users#list-microsoft-entra-roles-using-sql.
	rows, err := db.QueryContext(ctx, "SELECT * FROM pgaadauth_list_principals(false)")
	if err != nil {
		return nil, eris.Wrap(err, "listing Microsoft Entra principals")
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, eris.Wrap(err, "reading Microsoft Entra principal columns")
	}

	var result *AADPrincipal
	for rows.Next() {
		values := make([]sql.NullString, len(columns))
		dest := make([]any, len(columns))
		for i := range values {
			dest[i] = &values[i]
		}
		err = rows.Scan(dest...)
		if err != nil {
			return nil, eris.Wrap(err, "extracting Microsoft Entra principal")
		}

		principal := make(map[string]string, len(columns))
		for i, column := range columns {
			principal[strings.ToLower(column)] = values[i].String
		}
		if principal["rolname"] == roleName {
			result = &AADPrincipal{
				ObjectID:      principal["objectid"],
				PrincipalType: principal["principaltype"],
			}
		}
	}

	if err = rows.Err(); err != nil {
		return nil, eris.Wrap(err, "iterating Microsoft Entra principals")
	}

	return result, nil
}

// VerifyAADPrincipal returns an error unless an existing role is bound to the Microsoft Entra principal the spec asks
// for: any Entra principal when objectID is empty (the role was bound by display name, which can't be checked here),
// otherwise that object ID and, when given, that principal type.
func VerifyAADPrincipal(roleName string, actual *AADPrincipal, objectID string, principalType string) error {
	if actual == nil {
		return eris.Errorf("role %s already exists and is not bound to a Microsoft Entra principal", roleName)
	}

	// A role bound by display name has no object ID in the spec, and the name-to-object mapping can't be checked from the
	// database, so being bound to some Entra principal is all there is to verify.
	if objectID == "" {
		return nil
	}

	if !strings.EqualFold(actual.ObjectID, objectID) {
		return eris.Errorf(
			"role %s already exists and is bound to Microsoft Entra object %s, not %s",
			roleName,
			actual.ObjectID,
			objectID,
		)
	}

	if principalType != "" && !strings.EqualFold(actual.PrincipalType, principalType) {
		return eris.Errorf(
			"role %s already exists and is bound to a Microsoft Entra principal of type %s, not %s",
			roleName,
			actual.PrincipalType,
			principalType,
		)
	}

	return nil
}
