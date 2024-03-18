/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astbuilder

import (
	"testing"
)

func Test_NewVariableAssignment_WhenCalled_ReturnsExpectedCode(t *testing.T) {
	t.Parallel()

	helloWorld := StringLiteral("Hello, World!")

	actual := NewVariableAssignment(
		"myVar",
		helloWorld,
	)

	assertDeclExpected(t, actual)
}
