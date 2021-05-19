/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package helpers

type SQLRoleDelta struct {
	AddedRoles   map[string]struct{}
	DeletedRoles map[string]struct{}
}

func DiffCurrentAndExpectedSQLRoles(currentRoles map[string]struct{}, expectedRoles map[string]struct{}) SQLRoleDelta {
	result := SQLRoleDelta{
		AddedRoles:   make(map[string]struct{}),
		DeletedRoles: make(map[string]struct{}),
	}

	for role := range expectedRoles {
		// If an expected role isn't in the current role set, we need to add it
		if _, ok := currentRoles[role]; !ok {
			result.AddedRoles[role] = struct{}{}
		}
	}

	for role := range currentRoles {
		// If a current role isn't in the expected set, we need to remove it
		if _, ok := expectedRoles[role]; !ok {
			result.DeletedRoles[role] = struct{}{}
		}
	}

	return result
}
