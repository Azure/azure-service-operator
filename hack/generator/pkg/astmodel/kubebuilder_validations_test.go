/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_ValidateRequired(t *testing.T) {
	g := NewGomegaWithT(t)

	validation := ValidateRequired()
	comment := GenerateKubebuilderComment(validation)

	g.Expect(comment).To(Equal("// +kubebuilder:validation:Required"))
}

func Test_ValidateEnum(t *testing.T) {
	g := NewGomegaWithT(t)

	validation := ValidateEnum([]interface{}{1, true, "hello"})
	comment := GenerateKubebuilderComment(validation)

	g.Expect(comment).To(Equal("// +kubebuilder:validation:Enum={1,true,hello}"))
}
