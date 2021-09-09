// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiffCurrentAndExpectedSQLRoles(t *testing.T) {
	cases := []struct {
		name                string
		currentRoles        map[string]struct{}
		expectedRoles       map[string]struct{}
		expectedRoleCreates map[string]struct{}
		expectedRoleDeletes map[string]struct{}
	}{{
		name:                "Current and expected equal",
		currentRoles:        map[string]struct{}{"USAGE": {}},
		expectedRoles:       map[string]struct{}{"USAGE": {}},
		expectedRoleCreates: make(map[string]struct{}),
		expectedRoleDeletes: make(map[string]struct{}),
	}, {
		name:                "Expected has single role more than current",
		currentRoles:        map[string]struct{}{"USAGE": {}},
		expectedRoles:       map[string]struct{}{"USAGE": {}, "SELECT": {}},
		expectedRoleCreates: map[string]struct{}{"SELECT": {}},
		expectedRoleDeletes: make(map[string]struct{}),
	}, {
		name:                "Expected has single role less than current",
		currentRoles:        map[string]struct{}{"USAGE": {}, "SELECT": {}},
		expectedRoles:       map[string]struct{}{"USAGE": {}},
		expectedRoleCreates: make(map[string]struct{}),
		expectedRoleDeletes: map[string]struct{}{"SELECT": {}},
	}, {
		name:                "Expected has many roles less than current",
		currentRoles:        map[string]struct{}{"SELECT": {}, "INSERT": {}, "UPDATE": {}, "DELETE": {}, "CREATE": {}, "DROP": {}, "RELOAD": {}},
		expectedRoles:       map[string]struct{}{"SELECT": {}, "INSERT": {}},
		expectedRoleCreates: make(map[string]struct{}),
		expectedRoleDeletes: map[string]struct{}{"UPDATE": {}, "DELETE": {}, "CREATE": {}, "DROP": {}, "RELOAD": {}},
	}, {
		name:                "Expected has many roles more than current",
		currentRoles:        map[string]struct{}{"SELECT": {}, "INSERT": {}},
		expectedRoles:       map[string]struct{}{"SELECT": {}, "INSERT": {}, "UPDATE": {}, "DELETE": {}, "CREATE": {}, "DROP": {}, "RELOAD": {}},
		expectedRoleCreates: map[string]struct{}{"UPDATE": {}, "DELETE": {}, "CREATE": {}, "DROP": {}, "RELOAD": {}},
		expectedRoleDeletes: make(map[string]struct{}),
	}, {
		name:                "Including ALL sets added to ALL and doesn't delete any",
		currentRoles:        map[string]struct{}{"SELECT": {}, "INSERT": {}},
		expectedRoles:       map[string]struct{}{"ALL": {}, "UPDATE": {}},
		expectedRoleCreates: map[string]struct{}{"ALL": {}},
		expectedRoleDeletes: make(map[string]struct{}),
	}, {
		name:                "ALL is case insensitive",
		currentRoles:        map[string]struct{}{"SELECT": {}, "INSERT": {}},
		expectedRoles:       map[string]struct{}{"aLl": {}, "UPDATE": {}},
		expectedRoleCreates: map[string]struct{}{"ALL": {}},
		expectedRoleDeletes: make(map[string]struct{}),
	}}
	// There's no test for handling ALL in current roles because we
	// don't see that in the database - it gets expanded into all
	// permissions.

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			assert := assert.New(t)
			result := DiffCurrentAndExpectedSQLRoles(c.currentRoles, c.expectedRoles)
			assert.Equal(c.expectedRoleCreates, result.AddedRoles)
			assert.Equal(c.expectedRoleDeletes, result.DeletedRoles)
		})
	}

}
