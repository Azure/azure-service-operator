/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package genericarmclient

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestMakeTenantScopeARMID_ReturnsExpectedResult(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	id, err := MakeTenantScopeARMID("prov", "res", "a")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(id).To(Equal("/providers/prov/res/a"))
}

func TestMakeTenantScopeARMID_MismatchedParams_ReturnsError(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	_, err := MakeTenantScopeARMID("prov", "res")
	g.Expect(err).To(HaveOccurred())
}
