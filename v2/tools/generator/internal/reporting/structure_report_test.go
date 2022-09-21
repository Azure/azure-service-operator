/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package reporting

import (
	"bytes"
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/sebdah/goldie/v2"
)

func TestDebugReport_GeneratesExpectedOutput(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)
	gold := goldie.New(t)

	r := NewStructureReport("Hello World")
	addNodes(r, "root", 7)

	var buf bytes.Buffer
	g.Expect(r.SaveTo(&buf)).To(Succeed())

	gold.Assert(t, t.Name(), []byte(buf.String()))
}

func addNodes(r *StructureReport, name string, count int) {
	if count <= 0 {
		return // base case
	}

	for i := 0; i < count; i++ {
		n := fmt.Sprintf("%s.%d", name, i+1)
		nested := r.Addf(n)
		addNodes(nested, n, count-i-1)
	}
}
