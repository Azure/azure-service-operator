package helpers

import (
	"fmt"
)

type MySQLRoleDelta struct {
	AddedRoles   map[string]bool
	DeletedRoles map[string]bool
}

// Boolean value is true for grants which can only be applied at the server level.
var roles = map[string]bool{
	"SELECT":                  false,
	"INSERT":                  false,
	"UPDATE":                  false,
	"DELETE":                  false,
	"CREATE":                  false,
	"DROP":                    false,
	"RELOAD":                  true,
	"PROCESS":                 true,
	"REFERENCES":              false,
	"INDEX":                   false,
	"ALTER":                   false,
	"SHOW DATABASES":          true,
	"CREATE TEMPORARY TABLES": false,
	"LOCK TABLES":             false,
	"EXECUTE":                 false,
	"REPLICATION SLAVE":       true,
	"REPLICATION CLIENT":      true,
	"CREATE VIEW":             false,
	"SHOW VIEW":               false,
	"CREATE ROUTINE":          false,
	"ALTER ROUTINE":           false,
	"CREATE USER":             true,
	"EVENT":                   false,
	"TRIGGER":                 false,
}

func DiffCurrentAndExpectedMySQLRoles(currentRoles map[string]bool, expectedRoles map[string]bool) (MySQLRoleDelta, error) {
	result := MySQLRoleDelta{
		AddedRoles:   make(map[string]bool),
		DeletedRoles: make(map[string]bool),
	}

	for role, expectedServer := range expectedRoles {
		// If an expected role isn't in the current role set, we need to add it
		if _, ok := currentRoles[role]; !ok {
			if roles[role] && !expectedServer {
				return result, fmt.Errorf("%s is a server level role and cannot be applied to a database", role)
			} else {
				result.AddedRoles[role] = expectedServer
			}
		}
	}

	for role, currentServer := range currentRoles {
		// If a current role isn't in the expected set, we need to remove it
		if _, ok := expectedRoles[role]; !ok {
			result.DeletedRoles[role] = currentServer
		}
	}

	return result, nil
}
