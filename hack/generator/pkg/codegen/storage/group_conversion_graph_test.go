/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package storage

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/hack/generator/pkg/test"
)

func TestGroupConversionGraph_GivenLink_ReturnsLink(t *testing.T) {
	g := NewGomegaWithT(t)

	start := test.MakeLocalPackageReference("demo", "v1")
	finish := test.MakeLocalPackageReference("demo", "v2")

	graph := NewGroupConversionGraph("demo")
	graph.AddLink(start, finish)

	actual, ok := graph.LookupTransition(start)
	g.Expect(ok).To(BeTrue())
	g.Expect(actual).To(Equal(finish))
}
