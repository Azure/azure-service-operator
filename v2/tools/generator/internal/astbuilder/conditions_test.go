/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import "testing"

func Test_IfExprOk_WhenCalled_ReturnsExpectedCode(t *testing.T) {
	t.Parallel()

	actual := IfExprOk(
		"monster",
		"ok",
		CallQualifiedFunc("doctor", "NewMonster", StringLiteral("godzilla")),
		CallQualifiedFuncAsStmt("monster", "Roar"),
	)

	assertStmtExpected(t, actual)
}
