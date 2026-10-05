/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package reconcilers

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestHasReconcilePolicyAnnotationChanged(t *testing.T) {
	old := "detach-on-delete"
	new := "skip"

	t.Parallel()
	g := NewGomegaWithT(t)

	result := HasReconcilePolicyAnnotationChanged(&old, &new)
	g.Expect(result).Should(Equal(true))

	old = "skip"
	result = HasReconcilePolicyAnnotationChanged(&old, &new)
	g.Expect(result).Should(Equal(false))
}
