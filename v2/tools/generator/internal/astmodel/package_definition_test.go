/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_PropertyDefinitionWithTag_GivenTag_DoesNotModifyOriginal(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	original := NewPropertyDefinition("FullName", "fullName", StringType)
	original = original.WithTag("a", "b")
	field := original.WithTag("c", "d")

	g.Expect(field.tags).NotTo(Equal(original.tags))
}
