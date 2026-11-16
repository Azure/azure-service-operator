/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/rotisserie/eris"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/internal/set"
)

// databasePrivileges are the privileges PostgreSQL can grant on a database.
// See https://www.postgresql.org/docs/current/ddl-priv.html.
var databasePrivileges = []string{"CONNECT", "CREATE", "TEMPORARY"}

type SQLDatabasePrivilegeDelta struct {
	AddedPrivileges   set.Set[string]
	DeletedPrivileges set.Set[string]
}

// NormalizeDatabasePrivileges upper-cases the given privileges, expands ALL (and ALL PRIVILEGES)
// into the individual database privileges, maps TEMP to TEMPORARY, and returns an error for anything
// that is not a database privilege. PostgreSQL reports the expanded form when reading privileges back,
// so comparing normalized sets avoids granting or revoking on every reconcile.
func NormalizeDatabasePrivileges(privileges []string) (set.Set[string], error) {
	result := set.Make[string]()
	for _, p := range privileges {
		switch upper := strings.ToUpper(strings.TrimSpace(p)); upper {
		case "ALL", "ALL PRIVILEGES":
			for _, priv := range databasePrivileges {
				result.Add(priv)
			}
		case "TEMP":
			result.Add("TEMPORARY")
		case "CONNECT", "CREATE", "TEMPORARY":
			result.Add(upper)
		default:
			return nil, eris.Errorf("%q is not a database privilege (expected CONNECT, CREATE, TEMPORARY or ALL)", p)
		}
	}

	return result, nil
}

// DiffCurrentAndExpectedDatabasePrivileges returns the privileges to grant and to revoke so that
// the current privileges become the expected ones. Both sets are expected to be normalized.
func DiffCurrentAndExpectedDatabasePrivileges(currentPrivileges set.Set[string], expectedPrivileges set.Set[string]) SQLDatabasePrivilegeDelta {
	result := SQLDatabasePrivilegeDelta{
		AddedPrivileges:   set.Make[string](),
		DeletedPrivileges: set.Make[string](),
	}

	for priv := range expectedPrivileges {
		// If an expected privilege isn't in the current privilege set, we need to add it
		if !currentPrivileges.Contains(priv) {
			result.AddedPrivileges.Add(priv)
		}
	}

	for priv := range currentPrivileges {
		// If a current privilege isn't in the expected set, we need to remove it
		if !expectedPrivileges.Contains(priv) {
			result.DeletedPrivileges.Add(priv)
		}
	}

	return result
}

// GetUserDatabasePrivileges gets the privileges granted directly to the user on each database,
// keyed by database name. Databases the user owns are omitted: an owner implicitly holds every
// privilege on its database.
func GetUserDatabasePrivileges(ctx context.Context, db *sql.DB, user string) (map[string]set.Set[string], error) {
	rows, err := db.QueryContext(
		ctx,
		`SELECT d.datname, a.privilege_type
FROM pg_database d
CROSS JOIN LATERAL aclexplode(d.datacl) a
JOIN pg_roles r ON r.oid = a.grantee
WHERE r.rolname = $1 AND d.datdba <> r.oid`,
		user,
	)
	if err != nil {
		return nil, eris.Wrapf(err, "listing database privileges for user %s", user)
	}
	defer rows.Close()

	result := make(map[string]set.Set[string])
	for rows.Next() {
		var database, privilege string
		err = rows.Scan(&database, &privilege)
		if err != nil {
			return nil, eris.Wrap(err, "extracting database privilege")
		}
		if _, ok := result[database]; !ok {
			result[database] = set.Make[string]()
		}
		result[database].Add(privilege)
	}

	if err = rows.Err(); err != nil {
		return nil, eris.Wrap(err, "iterating database privileges")
	}

	return result, nil
}

// ReconcileUserDatabasePrivileges grants and revokes database privileges so that the user holds
// exactly the given privileges on each database (keyed by database name), and none on any other
// database. If there's an error applying privileges for one database it will still continue to
// apply privileges for subsequent databases (before reporting all errors).
func ReconcileUserDatabasePrivileges(ctx context.Context, db *sql.DB, user string, dbPrivs map[string][]string) error {
	desiredPrivs := make(map[string]set.Set[string], len(dbPrivs))
	for database, privs := range dbPrivs {
		normalized, err := NormalizeDatabasePrivileges(privs)
		if err != nil {
			return eris.Wrapf(err, "database %s", database)
		}
		desiredPrivs[database] = normalized
	}

	currentPrivs, err := GetUserDatabasePrivileges(ctx, db, user)
	if err != nil {
		return eris.Wrapf(err, "couldn't get existing database privileges for user %s", user)
	}

	allDatabases := make(set.Set[string])
	for database := range desiredPrivs {
		allDatabases.Add(database)
	}
	for database := range currentPrivs {
		allDatabases.Add(database)
	}

	var dbErrors []error
	for database := range allDatabases {
		current := currentPrivs[database]
		if current == nil {
			current = set.Make[string]()
		}
		expected := desiredPrivs[database]
		if expected == nil {
			expected = set.Make[string]()
		}

		delta := DiffCurrentAndExpectedDatabasePrivileges(current, expected)

		err = grantDatabasePrivileges(ctx, db, database, user, delta.AddedPrivileges)
		if err != nil {
			dbErrors = append(dbErrors, eris.Wrap(err, database))
		}
		err = revokeDatabasePrivileges(ctx, db, database, user, delta.DeletedPrivileges)
		if err != nil {
			dbErrors = append(dbErrors, eris.Wrap(err, database))
		}
	}

	return kerrors.NewAggregate(dbErrors)
}

// RevokeAllDatabasePrivileges revokes every privilege the user holds on any database.
// PostgreSQL refuses to drop a role that still holds privileges, so this must run before DropUser.
func RevokeAllDatabasePrivileges(ctx context.Context, db *sql.DB, user string) error {
	return ReconcileUserDatabasePrivileges(ctx, db, user, nil)
}

func grantDatabasePrivileges(ctx context.Context, db *sql.DB, database string, user string, privileges set.Set[string]) error {
	if len(privileges) == 0 {
		// Nothing to do
		return nil
	}

	// Privilege names come from the fixed list in NormalizeDatabasePrivileges (or from the catalog);
	// identifiers are escaped.
	_, err := db.ExecContext(
		ctx,
		fmt.Sprintf(
			"GRANT %s ON DATABASE %s TO %s",
			strings.Join(set.AsSortedSlice(privileges), ", "),
			escapeIdentifier(database),
			escapeIdentifier(user),
		),
	)
	if err != nil {
		return eris.Wrapf(err, "granting database privileges to user %s", user)
	}

	return nil
}

func revokeDatabasePrivileges(ctx context.Context, db *sql.DB, database string, user string, privileges set.Set[string]) error {
	if len(privileges) == 0 {
		// Nothing to do
		return nil
	}

	_, err := db.ExecContext(
		ctx,
		fmt.Sprintf(
			"REVOKE %s ON DATABASE %s FROM %s",
			strings.Join(set.AsSortedSlice(privileges), ", "),
			escapeIdentifier(database),
			escapeIdentifier(user),
		),
	)
	if err != nil {
		return eris.Wrapf(err, "revoking database privileges from user %s", user)
	}

	return nil
}
