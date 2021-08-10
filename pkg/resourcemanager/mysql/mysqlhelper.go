/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
)

// ServerPort is the default server port for sql server
const ServerPort = 3306

// DriverName is driver name for psqldb connection
const DriverName = "mysql"

// SystemDatabase is the name of the system database in a MySQL server
// where users and privileges are stored (and which we can always
// assume will exist).
const SystemDatabase = "mysql"

func GetFullSQLServerName(serverName string) string {
	return serverName + "." + config.Environment().MySQLDatabaseDNSSuffix
}

func GetFullyQualifiedUserName(userName string, serverName string) string {
	return fmt.Sprintf("%s@%s", userName, serverName)
}

// ConnectToSqlDb connects to the SQL db using the given credentials
func ConnectToSqlDB(ctx context.Context, driverName string, fullServer string, database string, port int, user string, password string) (*sql.DB, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?tls=true&interpolateParams=true", user, password, fullServer, port, database)

	db, err := sql.Open(driverName, connString)
	if err != nil {
		return db, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		return db, errors.Wrapf(err, "error pinging the mysql db (%s:%d/%s)", fullServer, port, database)
	}

	return db, err
}

// ConnectToSQLDBAsCurrentUser connects to the SQL DB using the specified MSI ClientID
func ConnectToSQLDBAsCurrentUser(
	ctx context.Context,
	driverName string,
	fullServer string,
	database string,
	port int,
	user string,
	clientID string) (*sql.DB, error) {

	tokenProvider, err := iam.GetMSITokenProviderForResourceByClientID(
		config.Environment().ResourceIdentifiers.OSSRDBMS,
		clientID)
	if err != nil {
		return nil, err
	}

	// In our case we can't pass the provider directly so we just invoke it and get the token and use that
	token, err := tokenProvider()
	if err != nil {
		return nil, err
	}

	// See https://docs.microsoft.com/en-us/azure/mysql/howto-connect-with-managed-identity
	// As noted here https://docs.microsoft.com/en-us/azure/mysql/howto-configure-sign-in-azure-ad-authentication#compatibility-with-application-drivers
	// we must specify allowCleartextPasswords to pass a token
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?tls=true&allowCleartextPasswords=true&interpolateParams=true", user, token, fullServer, port, database)

	db, err := sql.Open(driverName, connString)
	if err != nil {
		return db, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		return db, errors.Wrapf(err, "error pinging the mysql db (%s:%d/%s) as %s", fullServer, port, database, user)
	}

	return db, err
}

func formatUser(user string) string {
	// Wrap the user name in the weird formatting MySQL uses.
	return fmt.Sprintf("'%s'@'%%'", user)
}

// ExtractUserDatabaseRoles extracts the per-database roles that the
// user has. The user can have different permissions to each
// database. The details of access are returned in the map, keyed by
// database name.
func ExtractUserDatabaseRoles(ctx context.Context, db *sql.DB, user string) (map[string]StringSet, error) {
	// Note: This works because we only assign permissions at the DB level, not at the table, column, etc levels -- if we assigned
	// permissions at more levels we would need to do something else here such as join multiple tables or
	// parse SHOW GRANTS with a regex.
	rows, err := db.QueryContext(
		ctx,
		"SELECT TABLE_SCHEMA, PRIVILEGE_TYPE FROM INFORMATION_SCHEMA.SCHEMA_PRIVILEGES WHERE GRANTEE = ?",
		formatUser(user),
	)
	if err != nil {
		return nil, errors.Wrapf(err, "listing database grants for user %s", user)
	}
	defer rows.Close()

	results := make(map[string]StringSet)
	for rows.Next() {
		var database, privilege string
		err := rows.Scan(&database, &privilege)
		if err != nil {
			return nil, errors.Wrapf(err, "extracting privilege row")
		}

		var privileges StringSet
		if existingPrivileges, found := results[database]; found {
			privileges = existingPrivileges
		} else {
			privileges = make(StringSet)
			results[database] = privileges
		}
		privileges.Add(privilege)
	}

	if rows.Err() != nil {
		return nil, errors.Wrapf(rows.Err(), "iterating database privileges")
	}

	return results, nil
}

// ExtractUserServerRoles extracts the server-level privileges the user has as a set.
func ExtractUserServerRoles(ctx context.Context, db *sql.DB, user string) (StringSet, error) {
	// Note: This works because we only assign permissions at the DB level, not at the table, column, etc levels -- if we assigned
	// permissions at more levels we would need to do something else here such as join multiple tables or
	// parse SHOW GRANTS with a regex.
	// Remove "USAGE" as it's special and we never grant or remove it.
	rows, err := db.QueryContext(
		ctx,
		"SELECT PRIVILEGE_TYPE FROM INFORMATION_SCHEMA.USER_PRIVILEGES WHERE GRANTEE = ? AND PRIVILEGE_TYPE != 'USAGE'",
		formatUser(user),
	)
	if err != nil {
		return nil, errors.Wrapf(err, "listing grants for user %s", user)
	}
	defer rows.Close()

	result := make(StringSet)
	for rows.Next() {
		var row string
		err := rows.Scan(&row)
		if err != nil {
			return nil, errors.Wrapf(err, "extracting privilege field")
		}

		result.Add(row)
	}
	if rows.Err() != nil {
		return nil, errors.Wrapf(rows.Err(), "iterating privileges")
	}

	return result, nil
}

type StringSet map[string]struct{}

func SliceToSet(values []string) StringSet {
	result := make(StringSet)
	for _, value := range values {
		result[value] = struct{}{}
	}
	return result
}

func (s StringSet) Add(value string) {
	s[value] = struct{}{}
}

// EnsureUserServerRoles revokes and grants server-level roles as
// needed so the roles for the user match those passed in.
func EnsureUserServerRoles(ctx context.Context, db *sql.DB, user string, roles []string) error {
	var errorStrings []string
	if err := helpers.FindBadChars(user); err != nil {
		return errors.Wrap(err, "problem found with username")
	}

	desiredRoles := SliceToSet(roles)

	currentRoles, err := ExtractUserServerRoles(ctx, db, user)
	if err != nil {
		return errors.Wrapf(err, "couldn't get existing roles for user %s", user)
	}

	rolesDiff := helpers.DiffCurrentAndExpectedSQLRoles(currentRoles, desiredRoles)
	err = addRoles(ctx, db, "", user, rolesDiff.AddedRoles)
	if err != nil {
		errorStrings = append(errorStrings, err.Error())
	}
	err = deleteRoles(ctx, db, "", user, rolesDiff.DeletedRoles)
	if err != nil {
		errorStrings = append(errorStrings, err.Error())
	}

	if len(errorStrings) != 0 {
		return fmt.Errorf(strings.Join(errorStrings, "\n"))
	}
	return nil
}

// EnsureUserDatabaseRoles revokes and grants database roles as needed
// so they match the ones passed in. If there's an error applying
// privileges for one database it will still continue to apply
// privileges for subsequent databases (before reporting all errors).
func EnsureUserDatabaseRoles(ctx context.Context, conn *sql.DB, user string, dbRoles map[string][]string) error {
	if err := helpers.FindBadChars(user); err != nil {
		return errors.Errorf("problem found with username: %s", err)
	}

	desiredRoles := make(map[string]StringSet)
	for database, roles := range dbRoles {
		desiredRoles[database] = SliceToSet(roles)
	}

	currentRoles, err := ExtractUserDatabaseRoles(ctx, conn, user)
	if err != nil {
		return errors.Wrapf(err, "couldn't get existing database roles for user %s", user)
	}

	allDatabases := make(StringSet)
	for db := range desiredRoles {
		allDatabases.Add(db)
	}
	for db := range currentRoles {
		allDatabases.Add(db)
	}

	var dbErrors error
	for db := range allDatabases {
		rolesDiff := helpers.DiffCurrentAndExpectedSQLRoles(
			currentRoles[db],
			desiredRoles[db],
		)

		err = addRoles(ctx, conn, db, user, rolesDiff.AddedRoles)
		if err != nil {
			dbErrors = multierror.Append(dbErrors, errors.Wrap(err, db))
		}
		err = deleteRoles(ctx, conn, db, user, rolesDiff.DeletedRoles)
		if err != nil {
			dbErrors = multierror.Append(dbErrors, errors.Wrap(err, db))
		}
	}

	return dbErrors
}

func addRoles(ctx context.Context, db *sql.DB, database string, user string, roles StringSet) error {
	if len(roles) == 0 {
		// Nothing to do
		return nil
	}

	var rolesSlice []string
	for elem := range roles {
		rolesSlice = append(rolesSlice, elem)
	}

	toAdd := strings.Join(rolesSlice, ",")
	tsql := fmt.Sprintf("GRANT %s ON %s TO ?", toAdd, asGrantTarget(database))
	_, err := db.ExecContext(ctx, tsql, user)

	return err
}

func deleteRoles(ctx context.Context, db *sql.DB, database string, user string, roles StringSet) error {
	if len(roles) == 0 {
		// Nothing to do
		return nil
	}

	var rolesSlice []string
	for elem := range roles {
		rolesSlice = append(rolesSlice, elem)
	}

	toDelete := strings.Join(rolesSlice, ",")
	tsql := fmt.Sprintf("REVOKE %s ON %s FROM ?", toDelete, asGrantTarget(database))
	_, err := db.ExecContext(ctx, tsql, user)

	return err
}

// asGrantTarget formats the database name as a target suitable for a
// grant or revoke statement. If database is empty it returns "*.*"
// for server-level privileges.
func asGrantTarget(database string) string {
	if database == "" {
		return "*.*"
	}
	return fmt.Sprintf("`%s`.*", database)
}

// UserExists checks if db contains user
func UserExists(ctx context.Context, db *sql.DB, username string) (bool, error) {

	err := db.QueryRowContext(ctx, "SELECT * FROM mysql.user WHERE User = $1", username)

	if err != nil {
		return false, nil
	}
	return true, nil
}

// DropUser drops a user from db
func DropUser(ctx context.Context, db *sql.DB, user string) error {

	if err := helpers.FindBadChars(user); err != nil {
		return errors.Wrap(err, "problem found with username")
	}
	_, err := db.ExecContext(ctx, "DROP USER IF EXISTS ?", user)
	return err
}

// TODO: This is probably more generic than MySQL
func IsErrorResourceNotFound(err error) bool {
	requeueErrors := []string{
		errhelp.ResourceNotFound,
		errhelp.ParentNotFoundErrorCode,
		errhelp.ResourceGroupNotFoundErrorCode,
	}
	azerr := errhelp.NewAzureError(err)
	return helpers.ContainsString(requeueErrors, azerr.Type)
}

func IgnoreResourceNotFound(err error) error {
	if IsErrorResourceNotFound(err) {
		return nil
	}

	return err
}

func IsErrorDatabaseBusy(err error) bool {
	return strings.Contains(err.Error(), "Please retry the connection later")
}

func IgnoreDatabaseBusy(err error) error {
	if IsErrorDatabaseBusy(err) {
		return nil
	}

	return err
}
