// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib" // the pgx lib
	"github.com/rotisserie/eris"
)

// PSqlServerPort is the default server port for sql server
const PSqlServerPort = 5432

// PDriverName is driver name for psqldb connection
const PDriverName = "pgx"

// DefaultMaintanenceDatabase is the name of the database in a postgresql server
// where users and roles are stored (and which we can always
// assume will exist).
const DefaultMaintanenceDatabase = "postgres"

// ConnectToDB connects to the PostgreSQL db using the given credentials
func ConnectToDB(ctx context.Context, fullservername string, database string, port int, user string, password string) (*sql.DB, error) {
	connString := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=require connect_timeout=30", fullservername, user, password, port, database)

	db, err := sql.Open(PDriverName, connString)
	if err != nil {
		return db, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		return db, err
	}

	return db, err
}

func CreateUser(ctx context.Context, db *sql.DB, username string, password string) (*SQLUser, error) {
	if err := FindBadChars(username); err != nil {
		return nil, eris.Wrap(err, "problem found with username")
	}
	_, err := db.ExecContext(ctx, fmt.Sprintf("CREATE USER %s WITH PASSWORD %s", EscapeIdentifier(username), EscapeStringLiteral(password)))
	if err != nil {
		return nil, eris.Wrapf(err, "failed to create user %s", username)
	}
	return &SQLUser{Name: username}, nil
}

func UpdateUser(ctx context.Context, db *sql.DB, user SQLUser, password string) error {
	_, err := db.ExecContext(ctx, fmt.Sprintf("ALTER USER %s WITH PASSWORD %s", EscapeIdentifier(user.Name), EscapeStringLiteral(password)))
	if err != nil {
		return eris.Wrapf(err, "failed to alter user %s", user.Name)
	}
	return nil
}

func FindUserIfExist(ctx context.Context, db *sql.DB, username string) (*SQLUser, error) {
	res, err := db.ExecContext(ctx, "SELECT usename FROM pg_user WHERE usename = $1", username)
	if err != nil {
		return nil, err
	}
	rows, err := res.RowsAffected()
	if rows > 0 {
		return &SQLUser{Name: username}, err
	} else {
		return nil, err
	}
}

// DoesUserExist checks if db contains user
func DoesUserExist(ctx context.Context, db *sql.DB, username string) (bool, error) {
	res, err := db.ExecContext(ctx, "SELECT usename FROM pg_user WHERE usename = $1", username)
	if err != nil {
		return false, err
	}
	rows, err := res.RowsAffected()
	return rows > 0, err
}

// DropUser drops a user from db
func DropUser(ctx context.Context, db *sql.DB, user string) error {
	if err := FindBadChars(user); err != nil {
		return eris.Wrap(err, "problem found with username")
	}

	_, err := db.ExecContext(ctx, fmt.Sprintf("DROP USER IF EXISTS %s", EscapeIdentifier(user)))
	return err
}

// DatabaseExists checks if a database exists
func DatabaseExists(ctx context.Context, db *sql.DB, dbName string) (bool, error) {
	res, err := db.ExecContext(ctx, "SELECT datname FROM pg_catalog.pg_database WHERE datname = $1", dbName)
	if err != nil {
		return false, err
	}
	rows, err := res.RowsAffected()
	return rows > 0, err
}

// RoleExists checks if db contains role
func RoleExists(ctx context.Context, db *sql.DB, roleName string) (bool, error) {
	res, err := db.ExecContext(ctx, "SELECT * FROM pg_roles WHERE rolname = $1", roleName)
	if err != nil {
		return false, err
	}
	rows, err := res.RowsAffected()
	return rows > 0, err
}

// CreateRoleWithPermissions creates a role.
// Note that this is not currently used except for test
func CreateRoleWithPermissions(ctx context.Context, db *sql.DB, roleName string, permissions ...string) error {
	if err := FindBadChars(roleName); err != nil {
		return eris.Wrap(err, "problem found with roleName")
	}

	permissionString := strings.Join(permissions, " ")
	if err := FindBadChars(permissionString); err != nil {
		return eris.Wrap(err, "problem found with permissions")
	}

	_, err := db.ExecContext(ctx, fmt.Sprintf("CREATE ROLE %s WITH %s", EscapeIdentifier(roleName), permissionString))
	if err != nil {
		return eris.Wrap(err, "failed to create role")
	}

	return nil
}

// Use this type only for user, which are already checked
type SQLUser struct {
	Name string
}

// EscapeStringLiteral escapes a string for use as a PostgreSQL string literal.
// It wraps the value in single quotes and escapes any internal single quotes by doubling them.
// This is safe against SQL injection for string literals.
func EscapeStringLiteral(value string) string {
	escaped := strings.ReplaceAll(value, "'", "''")
	return "'" + escaped + "'"
}

// EscapeIdentifier escapes a string for use as a PostgreSQL identifier (e.g., username, role name).
// It wraps the value in double quotes and escapes any internal double quotes by doubling them.
// This is safe against SQL injection for identifiers.
func EscapeIdentifier(value string) string {
	escaped := strings.ReplaceAll(value, "\"", "\"\"")
	return "\"" + escaped + "\""
}

// FindBadChars checks for potentially dangerous character sequences in identifiers like usernames and role names.
// Note: This is still used for identifiers (usernames, role names) as an additional safety measure,
// but is no longer used for passwords since proper escaping via EscapeStringLiteral handles all characters safely.
func FindBadChars(stack string) error {
	badChars := []string{
		"'",
		"\"",
		";",
		"--",
		"/*",
	}

	for _, s := range badChars {
		if idx := strings.Index(stack, s); idx > -1 {
			return fmt.Errorf("potentially dangerous character sequence found: '%s' at pos: %d", s, idx)
		}
	}
	return nil
}
