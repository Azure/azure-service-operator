package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

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
		return db, fmt.Errorf("error ping the mysql db:  %v", err)
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
		return db, fmt.Errorf("error ping the mysql db:  %v", err)
	}

	return db, err
}

func GrantUserRoles(ctx context.Context, user string, database string, roles []string, db *sql.DB) error {
	var errorStrings []string
	if err := helpers.FindBadChars(user); err != nil {
		return fmt.Errorf("problem found with username: %v", err)
	}

	for _, role := range roles {

		if err := helpers.FindBadChars(role); err != nil {
			return fmt.Errorf("problem found with role: %v", err)
		}

		// Due to how go-mysql-driver performs parameter replacement, it always wraps
		// string parameters in ''. That doesn't work for this query because some of
		// our parameters are actually SQL keywords or identifiers (backticks). Admittedly
		// protecting against SQL injection here is probably pointless as we're giving the caller
		// permission to create users, which means there's nothing stopping them from creating
		// an administrator user and then doing whatever they want without SQL injection.
		// See https://github.com/go-sql-driver/mysql/blob/3b935426341bc5d229eafd936e4f4240da027ccd/connection.go#L198
		// for specifics of what go-mysql-driver supports.
		tsql := fmt.Sprintf("GRANT %s ON `%s`.* TO ?", role, database)
		_, err := db.ExecContext(ctx, tsql, user)
		if err != nil {
			errorStrings = append(errorStrings, err.Error())
		}
	}

	if len(errorStrings) != 0 {
		return fmt.Errorf(strings.Join(errorStrings, "\n"))
	}
	return nil
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
		return fmt.Errorf("problem found with username: %v", err)
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
