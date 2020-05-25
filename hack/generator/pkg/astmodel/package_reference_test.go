/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

// If `localPathPrefix` does not have a trailing slash, `stripLocalPackagePrefix()` and `NewLocalPackageReference()` will misbehave
func TestLocalPackageReferenceMustEndWithSlash(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(localPathPrefix).To(HaveSuffix("/"))
}
