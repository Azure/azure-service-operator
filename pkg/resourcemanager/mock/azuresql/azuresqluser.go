// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package azuresql

import (
	"context"
	"errors"

	"database/sql"
	dbsql "database/sql"

	azuresqluser "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqluser"
)

type MockSqlUserManager struct {
	sqlUsers []MockSqlUserResource
}

type MockSqlUserResource struct {
	username string
	roles    []string
	sqlDB    *dbsql.DB
}

func NewMockAzureSqlUserManager() *MockSqlUserManager {
	return &MockSqlUserManager{}
}

func findSqlUser(res []MockSqlUserResource, predicate func(MockSqlUserResource) bool) (int, MockSqlUserResource) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, MockSqlUserResource{}
}

func (manager *MockSqlUserManager) CreateUser(ctx context.Context, secret map[string][]byte, db *dbsql.DB) (string, error) {
	newUser := string(secret[azuresqluser.SecretUsernameKey])

	q := MockSqlUserResource{
		username: newUser,
		roles:    []string{},
		sqlDB:    db,
	}

	manager.sqlUsers = append(manager.sqlUsers, q)

	return newUser, nil
}

func (manager *MockSqlUserManager) DropUser(ctx context.Context, db *sql.DB, user string) error {
	sqlUsers := manager.sqlUsers

	index, _ := findSqlUser(manager.sqlUsers, func(s MockSqlUserResource) bool {
		return s.username == user
	})

	if index == -1 {
		return errors.New("Sql User Not Found")
	}

	manager.sqlUsers = append(sqlUsers[:index], sqlUsers[index+1:]...)

	return nil
}

func (manager *MockSqlUserManager) UserExists(ctx context.Context, db *sql.DB, username string) (bool, error) {

	index, _ := findSqlUser(manager.sqlUsers, func(s MockSqlUserResource) bool {
		return s.username == username
	})

	if index == -1 {
		return false, errors.New("Sql User Not Found")
	}

	return true, nil
}

func (manager *MockSqlUserManager) GrantUserRoles(ctx context.Context, user string, roles []string, db *sql.DB) error {

	index, _ := findSqlUser(manager.sqlUsers, func(s MockSqlUserResource) bool {
		return s.username == user
	})

	if index == -1 {
		return errors.New("Sql User Not Found")
	}

	manager.sqlUsers[index].roles = roles

	return nil
}
