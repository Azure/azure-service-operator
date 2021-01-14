/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

func TestErroredTypeProperties(t *testing.T) {
	g := gopter.NewProperties(nil)

	g.Property("wrapping an ErroredType with another merges its errors and warnings",
		prop.ForAll(
			func(e1 []string, e2 []string, w1 []string, w2 []string) bool {
				t := MakeErroredType(MakeErroredType(StringType, e1, w1), e2, w2)

				return stringSlicesEqual(t.errors, append(e2, e1...)) &&
					stringSlicesEqual(t.warnings, append(w2, w1...)) &&
					t.inner == StringType
			},
			gen.SliceOf(gen.AnyString()),
			gen.SliceOf(gen.AnyString()),
			gen.SliceOf(gen.AnyString()),
			gen.SliceOf(gen.AnyString())))

	g.TestingRun(t)
}
