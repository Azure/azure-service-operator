/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"testing"

	"github.com/dave/dst"
)

func TestGolden_MakeMap_GivenKeyAndValueTypes_GeneratesExpectedCode(t *testing.T) {
	t.Parallel()

	cases := []struct {
		key string
		val string
	}{
		{"string", "int"},
		{"int", "string"},
		{"string", "string"},
		{"int", "int"},
	}

	for _, c := range cases {
		c := c
		name := c.key + "_" + c.val
		t.Run(
			name,
			func(t *testing.T) {
				t.Parallel()

				key := dst.NewIdent(c.key)
				val := dst.NewIdent(c.val)

				call := MakeMap(key, val)
				assertExprExpected(t, call)
			})
	}
}
