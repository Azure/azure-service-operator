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
	v1 "k8s.io/api/core/v1"
)

// SecretUsernameKey is the username key in secret
const SecretUsernameKey = "username"

// SecretPasswordKey is the password key in secret
const SecretPasswordKey = "password"

type AzureSqlUserManager struct {
	Log logr.Logger
}

// Grants roles to a user for a given database
func (m *AzureSqlUserManager) GrantUserRoles(ctx context.Context, user string, roles []string, db *sql.DB) error {
	var errorStrings []string
	for _, role := range roles {
		tsql := fmt.Sprintf("sp_addrolemember \"%s\", \"%s\"", role, user)
		_, err := db.ExecContext(ctx, tsql)
		if err != nil {
			m.Log.Info("Error executing add role", "err", err.Error())
			errorStrings = append(errorStrings, err.Error())
		}
	}

	if len(errorStrings) != 0 {
		return fmt.Errorf(strings.Join(errorStrings, "\n"))
	}
	return nil
}

// Creates user with secret credentials
func (m *AzureSqlUserManager) CreateUser(ctx context.Context, secret *v1.Secret, db *sql.DB) (string, error) {
	newUser := string(secret.Data[SecretUsernameKey])
	newPassword := string(secret.Data[SecretPasswordKey])
	tsql := fmt.Sprintf("CREATE USER \"%s\" WITH PASSWORD='%s'", newUser, newPassword)
	_, err := db.ExecContext(ctx, tsql)

	// TODO: Have db lib do string interpolation
	//tsql := fmt.Sprintf(`CREATE USER @User WITH PASSWORD='@Password'`)
	//_, err := db.ExecContext(ctx, tsql, sql.Named("User", newUser), sql.Named("Password", newPassword))

	if err != nil {
		m.Log.Info("Error executing", "err", err.Error())
		return newUser, err
	}
	return newUser, nil
}

// ContainsUser checks if db contains user
func (m *AzureSqlUserManager) ContainsUser(ctx context.Context, db *sql.DB, username string) (bool, error) {
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
