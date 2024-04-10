/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"testing"

	"github.com/dave/dst"
)

func TestCallFunc_GivenArguments_GeneratesExpectedCode(t *testing.T) {
	t.Parallel()

	alpha := dst.NewIdent("alpha")
	beta := dst.NewIdent("beta")
	gamma := dst.NewIdent("gamma")
	delta := dst.NewIdent("delta")

	fooFn := CallFunc("foo", alpha, beta)
	barFn := CallFunc("bar", gamma, delta)

	cases := map[string]struct {
		arguments []dst.Expr
	}{
		"NoArguments": {},
		"OneArgument": {
			arguments: []dst.Expr{alpha},
		},
		"TwoArguments": {
			arguments: []dst.Expr{alpha, beta},
		},
		"ThreeArguments": {
			arguments: []dst.Expr{alpha, beta, gamma},
		},
		"FourArguments": {
			arguments: []dst.Expr{alpha, beta, gamma, delta},
		},
		"OneNestedArgument": {
			arguments: []dst.Expr{fooFn},
		},
		"TwoNestedArguments": {
			arguments: []dst.Expr{fooFn, barFn},
		},
	}

	for name, c := range cases {
		c := c
		name := name
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			call := CallFunc(name, c.arguments...)
			assertExprExpected(t, call)
		})
	}
}
