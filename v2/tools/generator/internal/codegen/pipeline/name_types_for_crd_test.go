/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"testing"

	"github.com/onsi/gomega"
)

func TestAppendPreservingSuffixPreservesSuffix(t *testing.T) {
	t.Parallel()

	g := gomega.NewWithT(t)
	str := "something3456"
	result := appendPreservingSuffix(str, "12", "3456")
	g.Expect(result).To(gomega.Equal("something123456"))
}

func TestAppendPreservingSuffixWithoutSuffix(t *testing.T) {
	t.Parallel()

	g := gomega.NewWithT(t)
	str := "something_suffix"
	result := appendPreservingSuffix(str, "_new", "_elsewise")
	g.Expect(result).To(gomega.Equal("something_suffix_new"))
}
