// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiffCurrentAndExpectedMySQLRoles(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		name                string
		currentRoles        map[string]bool
		expectedRoles       map[string]bool
		expectedRoleCreates map[string]bool
		expectedRoleDeletes map[string]bool
		expectedErr         bool
	}{
		{
			name:                "Current and expected equal",
			currentRoles:        map[string]bool{"INSERT": false},
			expectedRoles:       map[string]bool{"INSERT": false},
			expectedRoleCreates: make(map[string]bool),
			expectedRoleDeletes: make(map[string]bool),
			expectedErr:         false,
		},
		{
			name:                "Expected has single role more than current",
			currentRoles:        map[string]bool{"INSERT": false},
			expectedRoles:       map[string]bool{"INSERT": false, "SELECT": false},
			expectedRoleCreates: map[string]bool{"SELECT": false},
			expectedRoleDeletes: make(map[string]bool),
			expectedErr:         false,
		},
		{
			name:                "Expected has single role less than current",
			currentRoles:        map[string]bool{"INSERT": false, "SELECT": false},
			expectedRoles:       map[string]bool{"INSERT": false},
			expectedRoleCreates: make(map[string]bool),
			expectedRoleDeletes: map[string]bool{"SELECT": false},
			expectedErr:         false,
		},
		{
			name:                "Expected has many roles less than current",
			currentRoles:        map[string]bool{"SELECT": true, "INSERT": true, "UPDATE": true, "DELETE": true, "CREATE": true, "DROP": true, "RELOAD": true},
			expectedRoles:       map[string]bool{"SELECT": true, "INSERT": true},
			expectedRoleCreates: make(map[string]bool),
			expectedRoleDeletes: map[string]bool{"UPDATE": true, "DELETE": true, "CREATE": true, "DROP": true, "RELOAD": true},
			expectedErr:         false,
		},
		{
			name:                "Expected has many roles more than current",
			currentRoles:        map[string]bool{"SELECT": true, "INSERT": true},
			expectedRoles:       map[string]bool{"SELECT": true, "INSERT": true, "UPDATE": true, "DELETE": true, "CREATE": true, "DROP": true, "RELOAD": true},
			expectedRoleCreates: map[string]bool{"UPDATE": true, "DELETE": true, "CREATE": true, "DROP": true, "RELOAD": true},
			expectedRoleDeletes: make(map[string]bool),
			expectedErr:         false,
		},
		{
			name:                "Expected has server level roles but at the database level",
			currentRoles:        map[string]bool{"SELECT": false, "INSERT": false},
			expectedRoles:       map[string]bool{"RELOAD": false},
			expectedRoleCreates: make(map[string]bool),
			expectedRoleDeletes: make(map[string]bool),
			expectedErr:         true,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result, err := DiffCurrentAndExpectedMySQLRoles(c.currentRoles, c.expectedRoles)
			if err != nil {
				if !c.expectedErr {
					t.Errorf("unexepected error %s", err)
				}
			} else {
				assert.Equal(c.expectedRoleCreates, result.AddedRoles)
				assert.Equal(c.expectedRoleDeletes, result.DeletedRoles)
			}
		})
	}

}
