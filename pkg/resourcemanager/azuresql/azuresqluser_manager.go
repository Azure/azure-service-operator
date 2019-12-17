// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package azuresql

import (
	"context"
	"database/sql"

	"github.com/go-logr/logr"
)

func NewAzureSqlUserManager(log logr.Logger) *AzureSqlUserManager {
	return &AzureSqlUserManager{Log: log}
}

type SqlUserManager interface {
	GrantUserRoles(ctx context.Context, user string, roles []string, db *sql.DB) error
	CreateUser(ctx context.Context, secret map[string][]byte, db *sql.DB) (string, error)
	UserExists(ctx context.Context, db *sql.DB, username string) (bool, error)
	DropUser(ctx context.Context, db *sql.DB, user string) error
}
