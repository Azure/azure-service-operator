/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"maps"
	"slices"
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

// maxIdentifierLength is PostgreSQL's default max_identifier_length in bytes. PostgreSQL silently truncates
// longer identifiers, even quoted ones, so a longer database name could grant on a different database.
const maxIdentifierLength = 63

// NormalizeDatabasePrivileges expands ALL (and ALL PRIVILEGES) into the individual database privileges,
// maps TEMP to TEMPORARY, and returns an error for anything that is not an upper case database privilege. PostgreSQL reports the expanded form when reading privileges back,
// so comparing normalized sets avoids granting or revoking on every reconcile.
func NormalizeDatabasePrivileges(privileges []string) (set.Set[string], error) {
	result := set.Make[string]()
	for _, p := range privileges {
		switch p {
		case "ALL", "ALL PRIVILEGES":
			for _, priv := range databasePrivileges {
				result.Add(priv)
			}
		case "TEMP":
			result.Add("TEMPORARY")
		case "CONNECT", "CREATE", "TEMPORARY":
			result.Add(p)
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

// ValidateDatabasePrivileges returns an error if any of the given privileges is not a database privilege,
// or any database name is longer than PostgreSQL's identifier limit.
// It runs no SQL, so callers can reject a bad spec before changing anything on the server.
func ValidateDatabasePrivileges(dbPrivs map[string][]string) error {
	_, err := normalizeDatabasePrivilegeMap(dbPrivs)
	return err
}

func normalizeDatabasePrivilegeMap(dbPrivs map[string][]string) (map[string]set.Set[string], error) {
	result := make(map[string]set.Set[string], len(dbPrivs))
	var errs []error
	for _, database := range slices.Sorted(maps.Keys(dbPrivs)) {
		// Reject rather than truncate: PostgreSQL would silently target a different database.
		// The limit is in bytes, not characters.
		if len(database) > maxIdentifierLength {
			errs = append(errs, eris.Errorf("database name %q is %d bytes, longer than the PostgreSQL limit of %d", database, len(database), maxIdentifierLength))
			continue
		}
		normalized, err := NormalizeDatabasePrivileges(dbPrivs[database])
		if err != nil {
			errs = append(errs, eris.Wrapf(err, "database %s", database))
			continue
		}
		result[database] = normalized
	}

	if len(errs) > 0 {
		return nil, kerrors.NewAggregate(errs)
	}

	return result, nil
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

// getUserOwnedDatabases gets the names of the databases owned by the user.
func getUserOwnedDatabases(ctx context.Context, db *sql.DB, user string) (set.Set[string], error) {
	rows, err := db.QueryContext(
		ctx,
		`SELECT d.datname
FROM pg_database d
JOIN pg_roles r ON r.oid = d.datdba
WHERE r.rolname = $1`,
		user,
	)
	if err != nil {
		return nil, eris.Wrapf(err, "listing databases owned by user %s", user)
	}
	defer rows.Close()

	result := set.Make[string]()
	for rows.Next() {
		var database string
		err = rows.Scan(&database)
		if err != nil {
			return nil, eris.Wrap(err, "extracting database name")
		}
		result.Add(database)
	}

	if err = rows.Err(); err != nil {
		return nil, eris.Wrap(err, "iterating owned databases")
	}

	return result, nil
}

// PlanDatabasePrivileges returns, for each database whose privileges need to change, the privileges to grant
// and to revoke so that the current privileges become the desired ones. Databases in owned are skipped: an owner
// implicitly holds every privilege on its database, so its privileges are never managed.
func PlanDatabasePrivileges(
	current map[string]set.Set[string],
	desired map[string]set.Set[string],
	owned set.Set[string],
) map[string]SQLDatabasePrivilegeDelta {
	result := make(map[string]SQLDatabasePrivilegeDelta)
	for database := range unionKeys(current, desired) {
		if owned.Contains(database) {
			continue
		}

		delta := DiffCurrentAndExpectedDatabasePrivileges(current[database], desired[database])
		if len(delta.AddedPrivileges) > 0 || len(delta.DeletedPrivileges) > 0 {
			result[database] = delta
		}
	}

	return result
}

// ReconcileUserDatabasePrivileges grants and revokes database privileges so that the user holds
// exactly the given privileges on each database (keyed by database name), and none on any other
// database. Databases the user owns are left alone. If there's an error applying privileges for one
// database it will still continue to apply privileges for subsequent databases (before reporting all errors).
//
// PostgreSQL only warns, rather than failing, when the administrator lacks the grant option for some of
// the privileges, so the privileges are read back afterwards and any difference is reported as an error.
func ReconcileUserDatabasePrivileges(ctx context.Context, db *sql.DB, user string, dbPrivs map[string][]string) error {
	desiredPrivs, err := normalizeDatabasePrivilegeMap(dbPrivs)
	if err != nil {
		return err
	}

	currentPrivs, err := GetUserDatabasePrivileges(ctx, db, user)
	if err != nil {
		return eris.Wrapf(err, "couldn't get existing database privileges for user %s", user)
	}

	owned, err := getUserOwnedDatabases(ctx, db, user)
	if err != nil {
		return err
	}

	plan := PlanDatabasePrivileges(currentPrivs, desiredPrivs, owned)
	if len(plan) == 0 {
		// Nothing to do
		return nil
	}

	var dbErrors []error
	for _, database := range slices.Sorted(maps.Keys(plan)) {
		delta := plan[database]
		err = grantDatabasePrivileges(ctx, db, database, user, delta.AddedPrivileges)
		if err != nil {
			dbErrors = append(dbErrors, eris.Wrap(err, database))
		}
		err = revokeDatabasePrivileges(ctx, db, database, user, delta.DeletedPrivileges)
		if err != nil {
			dbErrors = append(dbErrors, eris.Wrap(err, database))
		}
	}

	if len(dbErrors) > 0 {
		return kerrors.NewAggregate(dbErrors)
	}

	// Verify that the changes took effect
	appliedPrivs, err := GetUserDatabasePrivileges(ctx, db, user)
	if err != nil {
		return eris.Wrapf(err, "couldn't verify database privileges for user %s", user)
	}

	return DatabasePrivilegeMismatches(appliedPrivs, desiredPrivs, owned)
}

// DatabasePrivilegeMismatches returns an error naming every database (other than those in owned) where the
// actual privileges differ from the desired ones, or nil if they all match.
func DatabasePrivilegeMismatches(
	actual map[string]set.Set[string],
	desired map[string]set.Set[string],
	owned set.Set[string],
) error {
	plan := PlanDatabasePrivileges(actual, desired, owned)

	errs := make([]error, 0, len(plan))
	for _, database := range slices.Sorted(maps.Keys(plan)) {
		delta := plan[database]
		errs = append(
			errs,
			eris.Errorf(
				"database %s: privileges for the user were not applied (missing: [%s], unexpected: [%s]); check that the administrator holds these privileges WITH GRANT OPTION and that no other role granted them",
				database,
				strings.Join(set.AsSortedSlice(delta.AddedPrivileges), ", "),
				strings.Join(set.AsSortedSlice(delta.DeletedPrivileges), ", "),
			))
	}

	return kerrors.NewAggregate(errs)
}

// RevokeAllDatabasePrivileges revokes every privilege the user holds on any database.
// PostgreSQL refuses to drop a role that still holds privileges, so this must run before DropUser.
func RevokeAllDatabasePrivileges(ctx context.Context, db *sql.DB, user string) error {
	return ReconcileUserDatabasePrivileges(ctx, db, user, nil)
}

func unionKeys[V any](maps ...map[string]V) set.Set[string] {
	result := set.Make[string]()
	for _, m := range maps {
		for k := range m {
			result.Add(k)
		}
	}

	return result
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
