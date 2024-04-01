/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"testing"

	"github.com/dave/dst"
)

func Test_MapLiteral_AsExpr_GeneratesExpectedDeclaration(t *testing.T) {
	t.Parallel()

	stringType := dst.NewIdent("string")
	intType := dst.NewIdent("int")

	content := []string{
		"alpha",
		"bravo",
		"charlie",
		"delta",
	}

	mapLiteral := NewMapLiteral(stringType, intType)
	for _, c := range content {
		mapLiteral.Add(StringLiteral(c), IntLiteral(len(c)))
	}

	assertExprExpected(t, mapLiteral.AsExpr())
}
