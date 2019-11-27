// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package sqlclient

import (
	"context"
	"database/sql"

	"github.com/go-logr/logr"
	v1 "k8s.io/api/core/v1"
)

func NewAzureSqlUserManager(log logr.Logger) *AzureSqlUserManager {
	return &AzureSqlUserManager{Log: log}
}

type SqlUserManager interface {
	GrantUserRoles(ctx context.Context, user string, roles []string, db *sql.DB) error
	CreateUser(ctx context.Context, secret *v1.Secret, db *sql.DB) (string, error)
	ContainsUser(ctx context.Context, db *sql.DB, username string) (bool, error)
	DropUser(ctx context.Context, db *sql.DB, user string) error
}
