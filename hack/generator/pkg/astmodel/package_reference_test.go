/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"github.com/onsi/gomega"
	"testing"
)

// If `localPathPrefix` does not have a trailing slash, `stripLocalPackagePrefix()` and `NewLocalPackageReference()` will misbehave
func TestLocalPackageReferenceMustEndWithSlash(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	g.Expect(localPathPrefix).To(gomega.HaveSuffix("/"))
}


