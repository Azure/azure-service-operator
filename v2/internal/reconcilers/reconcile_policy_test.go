/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package reconcilers

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
)

func TestParseReconcilePolicy(t *testing.T) {
	testingMap := map[string]annotations.ReconcilePolicyValue{
		"manage":           annotations.ReconcilePolicyManage,
		"skip":             annotations.ReconcilePolicySkip,
		"detach-on-delete": annotations.ReconcilePolicyDetachOnDelete,
	}

	t.Parallel()
	g := NewGomegaWithT(t)

	// test nominal case
	for policyString, policyValue := range testingMap {
		returnedPolicy, err := ParseReconcilePolicy(policyString, policyValue)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(returnedPolicy).Should(Equal(policyValue))
	}

	// test default value
	returnedPolicy, err := ParseReconcilePolicy("", annotations.ReconcilePolicySkip)
	g.Expect(err).ToNot((HaveOccurred()))
	g.Expect(returnedPolicy).Should(Equal(annotations.ReconcilePolicySkip))

	// test error in case of any other value
	returnedPolicy, err = ParseReconcilePolicy("whatever", annotations.ReconcilePolicySkip)
	g.Expect(err).Should(HaveOccurred())
	g.Expect(returnedPolicy).Should(Equal(annotations.ReconcilePolicySkip))
}

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
