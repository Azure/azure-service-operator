// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package sqlclient

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/go-logr/logr"
	"github.com/prometheus/common/log"
)

// SecretUsernameKey is the username key in secret
const SecretUsernameKey = "username"

// SecretPasswordKey is the password key in secret
const SecretPasswordKey = "password"

type AzureSqlUserManager struct {
	Log logr.Logger
}

// ConnectToSqlDb connects to the SQL db using the given credentials
func (m *AzureSqlUserManager) ConnectToSqlDb(ctx context.Context, drivername string, server string, database string, port int, user string, password string) (*sql.DB, error) {

	fullServerAddress := fmt.Sprintf("%s.database.windows.net", server)
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;Initial Catalog=%s;Persist Security Info=False;Pooling=False;MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=False;Connection Timeout=30", fullServerAddress, user, password, port, database)

	m.Log.Info("ConnectToSqlDb:", "user:", user)
	m.Log.Info("ConnectToSqlDb:", "password:", password)
	m.Log.Info("ConnectToSqlDb:", "conn string:", connString)
	db, err := sql.Open(drivername, connString)
	if err != nil {
		m.Log.Info("ConnectToSqlDb", "error from sql.Open is:", err.Error())
		return db, err
	}

	err = db.Ping()
	if err != nil {
		m.Log.Info("ConnectToSqlDb", "error from db.Ping is:", err.Error(), "full error", err)
		return db, err
	}

	return db, err
}

// Grants roles to a user for a given database
func (m *AzureSqlUserManager) GrantUserRoles(ctx context.Context, user string, roles []string, db *sql.DB) error {
	var errorStrings []string
	for _, role := range roles {
		tsql := fmt.Sprintf("sp_addrolemember \"%s\", \"%s\"", role, user)
		_, err := db.ExecContext(ctx, tsql)
		if err != nil {
			m.Log.Info("GrantUserRoles:", "Error executing add role:", err.Error())
			errorStrings = append(errorStrings, err.Error())
		}
	}

	if len(errorStrings) != 0 {
		return fmt.Errorf(strings.Join(errorStrings, "\n"))
	}
	return nil
}

// Creates user with secret credentials
func (m *AzureSqlUserManager) CreateUser(ctx context.Context, secret map[string][]byte, db *sql.DB) (string, error) {
	newUser := string(secret[SecretUsernameKey])
	newPassword := string(secret[SecretPasswordKey])
	tsql := fmt.Sprintf("CREATE USER \"%s\" WITH PASSWORD='%s'", newUser, newPassword)
	_, err := db.ExecContext(ctx, tsql)

	// TODO: Have db lib do string interpolation
	//tsql := fmt.Sprintf(`CREATE USER @User WITH PASSWORD='@Password'`)
	//_, err := db.ExecContext(ctx, tsql, sql.Named("User", newUser), sql.Named("Password", newPassword))

	if err != nil {
		log.Error("Error executing", "err", err.Error())
		return newUser, err
	}
	return newUser, nil
}

// UserExists checks if db contains user
func (m *AzureSqlUserManager) UserExists(ctx context.Context, db *sql.DB, username string) (bool, error) {
	res, err := db.ExecContext(ctx, fmt.Sprintf("SELECT * FROM sysusers WHERE NAME='%s'", username))
	if err != nil {
		return false, err
	}
	rows, err := res.RowsAffected()
	return rows > 0, err
}

// Drops user from db
func (m *AzureSqlUserManager) DropUser(ctx context.Context, db *sql.DB, user string) error {
	tsql := fmt.Sprintf("DROP USER \"%s\"", user)
	_, err := db.ExecContext(ctx, tsql)
	return err
}
