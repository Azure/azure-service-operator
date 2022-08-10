/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql" //mysql drive link
	"github.com/pkg/errors"
)

// ServerPort is the default server port for sql server
const ServerPort = 3306

// DriverName is driver name for psqldb connection
const DriverName = "mysql"

// SystemDatabase is the name of the system database in a MySQL server
// where users and privileges are stored (and which we can always
// assume will exist).
const SystemDatabase = "mysql"

func ConnectToDB(ctx context.Context, serverAddress string, database string, port int, user string, password string) (*sql.DB, error) {
	c := mysql.NewConfig()
	c.Addr = fmt.Sprintf("%s:%d", serverAddress, port)
	c.DBName = database
	c.User = user
	c.Passwd = password
	c.TLSConfig = "true"

	// Set other options
	c.InterpolateParams = true
	c.Net = "tcp"

	db, err := sql.Open(DriverName, c.FormatDSN())
	if err != nil {
		return db, err
	}
	db.SetConnMaxLifetime(1 * time.Minute)
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	// We ping here to ensure that the connection is actually viable, as per
	// https://github.com/go-sql-driver/mysql/wiki/Examples#a-word-on-sqlopen
	err = db.PingContext(ctx)
	if err != nil {
		return db, errors.Wrapf(err, "error pinging the mysql db (%s:%d/%s)", serverAddress, port, database)
	}

	return db, err
}

func HostnameOrDefault(hostname string) string {
	if hostname == "" {
		hostname = "%"
	}

	return hostname
}

func CreateOrUpdateUser(ctx context.Context, db *sql.DB, username string, hostname string, password string) error {
	hostname = HostnameOrDefault(hostname)

	// we call both CREATE and ALTER here so achieve an idempotent operation that also updates the password seamlessly
	// if it has changed

	// TODO: Support aliasing: https://github.com/Azure/azure-service-operator/issues/1402
	statement := "CREATE USER IF NOT EXISTS ?@? IDENTIFIED BY ?"
	_, err := db.ExecContext(ctx, statement, username, hostname, password)
	if err != nil {
		return errors.Wrapf(err, "failed to create user %s", username)
	}

	// TODO: Support aliasing: https://github.com/Azure/azure-service-operator/issues/1402
	statement = "ALTER USER IF EXISTS ?@? IDENTIFIED BY ?"
	_, err = db.ExecContext(ctx, statement, username, hostname, password)
	if err != nil {
		return errors.Wrapf(err, "failed to alter user %s", username)
	}

	return nil
}

// DoesUserExist checks if db contains user
func DoesUserExist(ctx context.Context, db *sql.DB, username string) (bool, error) {
	row := db.QueryRowContext(ctx, "SELECT User FROM mysql.user WHERE User = ?", username)
	var name string
	err := row.Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			// User doesn't exist
			return false, nil
		}

		// Something else went wrong
		return false, err
	}

	return true, nil
}

// DropUser drops a user from db
func DropUser(ctx context.Context, db *sql.DB, username string) error {
	_, err := db.ExecContext(ctx, "DROP USER IF EXISTS ?", username)
	return err
}
