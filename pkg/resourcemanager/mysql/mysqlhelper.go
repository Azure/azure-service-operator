package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
)

// MSqlServerPort is the default server port for sql server
const MySQLServerPort = 3306

// MDriverName is driver name for psqldb connection
const MySQLDriverName = "mysql"

func GetMySQLDatabaseDNSSuffix() string {
	// TODO: We need an environment specific way of getting the DNS suffix
	// TODO: which the Go SDK doesn't seem to have.
	// TODO: see: https://github.com/Azure/azure-sdk-for-go/issues/13749
	return "mysql.database.azure.com"
}

func GetFullSQLServerName(serverName string) string {
	return serverName + "." + GetMySQLDatabaseDNSSuffix()
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
		return db, fmt.Errorf("error pinging the mysql db (%s:%d/%s): %v", fullServer, port, database, err)
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

	tokenProvider, err := iam.GetMSITokenProviderForResourceByClientID("https://ossrdbms-aad.database.windows.net", clientID)
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
		return db, fmt.Errorf("error pinging the mysql db (%s:%d/%s) as %s: %v", fullServer, port, database, user, err)
	}

	return db, err
}

// GetAllGrants fetches grants for all DBs.
func GetAllGrants(ctx context.Context, db *sql.DB, user string) (map[string]string, error) {

	result := make(map[string]string)
	// Get server wide grants.
	rows, err := db.QueryContext(ctx, "SHOW GRANTS FOR ?", user)
	if err != nil {
		return nil, errors.Wrapf(err, "listing grants for user %s", user)
	}
	defer rows.Close()

	// Boolean value is true for grants which can only be applied at the server level.
	for rows.Next() {
		var row string
		err := rows.Scan(&row)
		if err != nil {
			return result, errors.Wrapf(err, "iterating returned rows")
		}
		r := regexp.MustCompile(`GRANT (.*) ON (.*) TO`)
		match := r.FindStringSubmatch(row)
		if match != nil {
			result[match[2]] = match[1]
		}
	}

	return result, nil
}

// ExtractUserRoles extracts the roles the user has. The result is the set of roles the user has for the requested database
func ExtractUserRoles(ctx context.Context, db *sql.DB, user string, database string) (map[string]bool, error) {
	if err := helpers.FindBadChars(user); err != nil {
		return nil, errors.Wrapf(err, "problem found with username")
	}

	isServerLevel := false
	if database == "" {
		database = "*.*"
		isServerLevel = true
	}

	result := make(map[string]bool)
	// Get server wide grants.
	rows, err := db.QueryContext(ctx, "SHOW GRANTS FOR ?", user)
	if err != nil {
		return nil, errors.Wrapf(err, "listing grants for user %s", user)
	}
	defer rows.Close()

	// Boolean value is true for grants which can only be applied at the server level.
	for rows.Next() {
		var row string
		err := rows.Scan(&row)
		if err != nil {
			return result, errors.Wrapf(err, "iterating returned rows")
		}
		r := regexp.MustCompile(`GRANT (.*) ON (.*) TO`)
		match := r.FindStringSubmatch(row)
		if match != nil {
			// Determine if grant statement is for server level (i.e. *.*), or the specific database of concern.
			if match[2] == database {
				// Add list of grants to result map.
				for _, grant := range strings.Split(match[1], ", ") {
					result[grant] = isServerLevel
				}
			}
		}
	}

	return result, nil
}

func ProcessRows(ctx context.Context, rows *sql.Rows, database string, result map[string]bool) error {
	return nil
}

func GrantUserRoles(ctx context.Context, user string, database string, roles []string, db *sql.DB) error {
	var errorStrings []string
	if err := helpers.FindBadChars(user); err != nil {
		return fmt.Errorf("problem found with username: %v", err)
	}

	// Boolean value is true for grants which can only be applied at the server level.
	isServerLevel := false
	if database == "" {
		isServerLevel = true
	}
	rolesMap := make(map[string]bool)
	for _, role := range roles {
		rolesMap[role] = isServerLevel
	}

	// Get the current roles
	currentRoles, err := ExtractUserRoles(ctx, db, user, database)
	if err != nil {
		return errors.Wrapf(err, "couldn't get existing roles for user %s", user)
	}
	// Remove "USAGE" as it's special and we never grant or remove it
	delete(currentRoles, "USAGE")

	rolesDiff, err := helpers.DiffCurrentAndExpectedMySQLRoles(currentRoles, rolesMap)
	if err != nil {
		return errors.Wrapf(err, "problem with requested roles")
	}

	// Due to how go-mysql-driver performs parameter replacement, it always wraps
	// string parameters in ''. That doesn't work for these queries because some of
	// our parameters are actually SQL keywords or identifiers (requiring backticks). Admittedly
	// protecting against SQL injection here is probably pointless as we're giving the caller
	// permission to create users, which means there's nothing stopping them from creating
	// an administrator user and then doing whatever they want without SQL injection.
	// See https://github.com/go-sql-driver/mysql/blob/3b935426341bc5d229eafd936e4f4240da027ccd/connection.go#L198
	// for specifics of what go-mysql-driver supports.
	err = addRoles(ctx, db, database, user, rolesDiff.AddedRoles)
	if err != nil {
		errorStrings = append(errorStrings, err.Error())
	}
	err = deleteRoles(ctx, db, database, user, rolesDiff.DeletedRoles)
	if err != nil {
		errorStrings = append(errorStrings, err.Error())
	}

	if len(errorStrings) != 0 {
		return fmt.Errorf(strings.Join(errorStrings, "\n"))
	}
	return nil
}

func addRoles(ctx context.Context, db *sql.DB, database string, user string, roles map[string]bool) error {
	if len(roles) == 0 {
		// Nothing to do
		return nil
	}

	if database == "" {
		database = "*.*"
	} else {
		database = fmt.Sprintf("`%s`.*", database)
	}

	var rolesSlice []string
	for elem := range roles {
		rolesSlice = append(rolesSlice, elem)
	}

	toAdd := strings.Join(rolesSlice, ",")
	tsql := fmt.Sprintf("GRANT %s ON %s TO ?", toAdd, database)
	_, err := db.ExecContext(ctx, tsql, user)

	return err
}

func deleteRoles(ctx context.Context, db *sql.DB, database string, user string, roles map[string]bool) error {
	if len(roles) == 0 {
		// Nothing to do
		return nil
	}

	if database == "" {
		database = "*.*"
	} else {
		database = fmt.Sprintf("`%s`.*", database)
	}

	var rolesSlice []string
	for elem := range roles {
		rolesSlice = append(rolesSlice, elem)
	}

	toDelete := strings.Join(rolesSlice, ",")
	tsql := fmt.Sprintf("REVOKE %s ON %s FROM ?", toDelete, database)
	_, err := db.ExecContext(ctx, tsql, user)

	return err
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
func DropUser(ctx context.Context, db *sql.DB, database string, user string) error {

	if err := helpers.FindBadChars(user); err != nil {
		return fmt.Errorf("problem found with username: %v", err)
	}

	if database == "" {
		database = "*.*"
	} else {
		database = fmt.Sprintf("`%s`.*", database)
	}

	grants, err := GetAllGrants(ctx, db, user)
	if err != nil {
		return err
	}
	// Nothing to do If the user has no grants on the database.
	if _, ok := grants[database]; !ok {
		return nil
	}

	tsql := fmt.Sprintf("REVOKE ALL ON %s FROM ?", database)
	_, err = db.ExecContext(ctx, tsql, user)

	// Put back the grants on specific databases if the revoked privileges were server-wide.
	if database == "*.*" {
		delete(grants, "*.*")
		for database, roles := range grants {
			tsql := fmt.Sprintf("GRANT %s ON %s TO ?", roles, database)
			_, err := db.ExecContext(ctx, tsql, user)
			if err != nil {
				return err
			}
		}
	}
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
