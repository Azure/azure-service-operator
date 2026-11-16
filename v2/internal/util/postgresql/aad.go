/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package postgresql

import (
	"context"
	"database/sql"

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
