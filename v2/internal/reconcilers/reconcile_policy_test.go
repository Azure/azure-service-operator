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

func TestParseReconcilePolicyNoDefault(t *testing.T) {
	testingMap := map[string]annotations.ReconcilePolicyValue{
		"manage":           annotations.ReconcilePolicyManage,
		"skip":             annotations.ReconcilePolicySkip,
		"detach-on-delete": annotations.ReconcilePolicyDetachOnDelete,
	}

	t.Parallel()
	g := NewGomegaWithT(t)

	// test nominal case
	for policyString, policyValue := range testingMap {
		returnedPolicy, err := ParseReconcilePolicyNoDefault(policyString)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(returnedPolicy).Should(Equal(policyValue))

		returnedPolicy, err = ParseReconcilePolicy(policyString, policyValue)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(returnedPolicy).Should(Equal(policyValue))
	}

	// ParseReconcilePolicyNoDefault should raise an error when an empty reconcile policy is provided
	// and should return default manage policy
	returnedPolicy, err := ParseReconcilePolicyNoDefault("")
	g.Expect(err).Should((HaveOccurred()))
	g.Expect(returnedPolicy).Should(Equal(annotations.ReconcilePolicyManage))

	// ParseReconcilePolicyNoDefault should raise an error when a wrong reconcile policy is provided
	// and should return default manage policy
	returnedPolicy, err = ParseReconcilePolicyNoDefault("whatever")
	g.Expect(err).Should(HaveOccurred())
	g.Expect(returnedPolicy).Should(Equal(annotations.ReconcilePolicyManage))
}

func TestParseReconcilePolicy(t *testing.T) {
	testingMap := map[string]annotations.ReconcilePolicyValue{
		"manage":           annotations.ReconcilePolicyManage,
		"skip":             annotations.ReconcilePolicySkip,
		"detach-on-delete": annotations.ReconcilePolicyDetachOnDelete,
	}

	t.Parallel()
	g := NewGomegaWithT(t)

	// ParseReconcilePolicy should raise an error when an empty reconcile policy is provided
	// and should return the wanted default policy provided as input
	for policyString, policyValue := range testingMap {
		returnedPolicy, err := ParseReconcilePolicy("", policyValue)
		g.Expect(err).Should(HaveOccurred())
		g.Expect(returnedPolicy).Should(Equal(policyValue))

		returnedPolicy, err = ParseReconcilePolicy(policyString, annotations.ReconcilePolicySkip)
		g.Expect(err).ShouldNot(HaveOccurred())
		g.Expect(returnedPolicy).Should(Equal(policyValue))
	}
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
